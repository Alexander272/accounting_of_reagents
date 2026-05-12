package services

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/events"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository/postgres"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/auth"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/logger"
	"github.com/Nerzal/gocloak/v13"
)

type UserService struct {
	repo      repository.User
	keycloak  *auth.KeycloakClient
	userRealm UserRealm
	tm        TransactionManager
	eventBus  *events.PolicyEventManager
}

type UserDeps struct {
	Repo      repository.User
	TM        TransactionManager
	Keycloak  *auth.KeycloakClient
	UserRealm UserRealm
	EventBus  *events.PolicyEventManager
}

func NewUserService(deps *UserDeps) *UserService {
	return &UserService{
		repo:      deps.Repo,
		keycloak:  deps.Keycloak,
		userRealm: deps.UserRealm,
		tm:        deps.TM,
		eventBus:  deps.EventBus,
	}
}

type Users interface {
	LoadPolicy(ctx context.Context) ([]*models.UserRealm, error)
	SyncFromKeycloak(ctx context.Context) error
	GetByUsername(ctx context.Context, username string) (*models.UserData, error)
	GetById(ctx context.Context, id string) (*models.UserData, error)
	GetAll(ctx context.Context) ([]*models.UserData, error)
	UpdateRole(ctx context.Context, dto *models.UserDataDTO) error
}

func (s *UserService) LoadPolicy(ctx context.Context) ([]*models.UserRealm, error) {
	tmp, err := s.userRealm.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load user policy: %w", err)
	}

	data := []*models.UserRealm{}
	for _, ur := range tmp {
		if ur.IsActive {
			data = append(data, ur)
		}
	}

	return data, nil
}

func (s *UserService) SyncFromKeycloak(ctx context.Context) error {
	logger.Info("Sync users started")

	token, err := s.keycloak.Login(ctx)
	if err != nil {
		return fmt.Errorf("failed to login admin to keycloak: %w", err)
	}

	// 1. Поиск ID группы (например, "reagents")
	groupName := "reagents"
	group, err := s.keycloak.Client.GetGroupByPath(ctx, token.AccessToken, s.keycloak.Realm, "/"+groupName)
	if err != nil {
		return fmt.Errorf("failed to get group by path: %w", err)
	}

	// 2. Собираем ID всех подгрупп рекурсивно из поля SubGroups
	allGroupIDs := []string{*group.ID}
	if group.SubGroups != nil {
		s.collectSubGroupIDs(*group.SubGroups, &allGroupIDs)
	}

	// 3. Получаем пользователей из группы и всех подгрупп
	userMap := make(map[string]gocloak.User)
	for _, gid := range allGroupIDs {
		members, err := s.keycloak.Client.GetGroupMembers(ctx, token.AccessToken, s.keycloak.Realm, gid, gocloak.GetGroupsParams{Max: gocloak.IntP(1000)})
		if err != nil {
			return fmt.Errorf("failed to get group members for group %s: %w", gid, err)
		}
		for _, m := range members {
			if m.ID != nil {
				userMap[*m.ID] = *m
			}
		}
	}

	if len(userMap) == 0 {
		logger.Info("group 'reagents' and sub-groups are empty")
		return nil
	}

	keycloakUsers := make([]gocloak.User, 0, len(userMap))
	for _, u := range userMap {
		keycloakUsers = append(keycloakUsers, u)
	}

	// Конвертация в KeycloakUser
	kcData := make([]*models.KeycloakUser, 0, len(keycloakUsers))
	for _, u := range keycloakUsers {
		if u.Enabled != nil && !*u.Enabled {
			continue
		}
		kcData = append(kcData, &models.KeycloakUser{
			Id:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		})
	}

	// 4. Синхронизация с БД
	if err := s.repo.SyncFromKeycloak(ctx, kcData); err != nil {
		return fmt.Errorf("failed to sync users to database: %w", err)
	}

	s.eventBus.Notify(events.PolicyEvent{})

	logger.Info("Sync finished", slog.Int("synced_users", len(kcData)))
	return nil
}

// collectSubGroupIDs рекурсивно собирает ID всех подгрупп из поля SubGroups
func (s *UserService) collectSubGroupIDs(subGroups []gocloak.Group, allIDs *[]string) {
	for _, sg := range subGroups {
		if sg.ID == nil {
			continue
		}
		*allIDs = append(*allIDs, *sg.ID)
		if sg.SubGroups != nil {
			s.collectSubGroupIDs(*sg.SubGroups, allIDs)
		}
	}
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*models.UserData, error) {
	data, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return data, nil
}

func (s *UserService) GetById(ctx context.Context, id string) (*models.UserData, error) {
	data, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return data, nil
}

func (s *UserService) GetAll(ctx context.Context) ([]*models.UserData, error) {
	data, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	return data, nil
}

func (s *UserService) UpdateRole(ctx context.Context, dto *models.UserDataDTO) error {
	created := []*models.UserRealmDTO{}
	updated := []*models.UserRealmDTO{}
	deleted := []*models.UserRealmDTO{}

	for _, r := range dto.Realms {
		if r.CreatedAt == "" {
			created = append(created, r)
		} else if r.RoleId != "" {
			updated = append(updated, r)
		} else {
			deleted = append(deleted, r)
		}
	}

	return s.tm.WithinTransaction(ctx, func(tx postgres.Tx) error {
		if err := s.userRealm.CreateSeveral(ctx, tx, created); err != nil {
			return fmt.Errorf("failed to create user realms: %w", err)
		}
		if err := s.userRealm.UpdateSeveral(ctx, tx, updated); err != nil {
			return fmt.Errorf("failed to update user realms: %w", err)
		}
		if err := s.userRealm.DeleteSeveral(ctx, tx, deleted); err != nil {
			return fmt.Errorf("failed to delete user realms: %w", err)
		}

		s.eventBus.Notify(events.PolicyEvent{})
		return nil
	})
}
