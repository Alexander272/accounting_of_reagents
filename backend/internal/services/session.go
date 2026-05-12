package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/models"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/auth"
)

type SessionService struct {
	keycloak  *auth.KeycloakClient
	userRealm UserRealm
	user      Users
	policies  AccessPolices
}

func NewSessionService(keycloak *auth.KeycloakClient, policies AccessPolices, userRealm UserRealm, user Users) *SessionService {
	return &SessionService{
		keycloak:  keycloak,
		policies:  policies,
		userRealm: userRealm,
		user:      user,
	}
}

type Session interface {
	SignIn(ctx context.Context, u models.SignIn) (*models.User, error)
	SignOut(ctx context.Context, refreshToken string) error
	Refresh(ctx context.Context, refreshToken, realm string) (*models.User, error)
	DecodeAccessToken(ctx context.Context, token string) (*models.User, error)
}

func (s *SessionService) SignIn(ctx context.Context, u models.SignIn) (*models.User, error) {
	res, err := s.keycloak.Client.Login(ctx, s.keycloak.ClientId, s.keycloak.ClientSecret, s.keycloak.Realm, u.Username, u.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to login to keycloak. error: %w", err)
	}

	user, err := s.DecodeAccessToken(ctx, res.AccessToken)
	if err != nil {
		return nil, err
	}

	// access, err := s.policies.GetPolicies(user.Id, "")
	// if err != nil {
	// 	return nil, err
	// }

	userRealms, err := s.userRealm.GetByUserId(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	user.Realms = userRealms

	user.Permissions = map[string][]string{}
	for _, r := range userRealms {
		access, err := s.policies.GetPolicies(user.Id, r.RealmId)
		if err != nil {
			return nil, err
		}

		user.Permissions[r.RealmId] = access.Perms
	}
	// user.Role = access.Role
	// user.Permissions = map[string][]string{
	// 	access.Domain: access.Perms,
	// }

	user.AccessToken = res.AccessToken
	user.RefreshToken = res.RefreshToken

	return user, nil
}

func (s *SessionService) SignOut(ctx context.Context, refreshToken string) error {
	err := s.keycloak.Client.Logout(ctx, s.keycloak.ClientId, s.keycloak.ClientSecret, s.keycloak.Realm, refreshToken)
	if err != nil {
		return fmt.Errorf("failed to logout to keycloak. error: %w", err)
	}
	return nil
}

func (s *SessionService) Refresh(ctx context.Context, refreshToken, realm string) (*models.User, error) {
	res, err := s.keycloak.Client.RefreshToken(ctx, refreshToken, s.keycloak.ClientId, s.keycloak.ClientSecret, s.keycloak.Realm)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token in keycloak. error: %w", err)
	}

	user, err := s.DecodeAccessToken(ctx, res.AccessToken)
	if err != nil {
		return nil, err
	}

	userRealms, err := s.userRealm.GetByUserId(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	user.Realms = userRealms

	user.Permissions = map[string][]string{}
	for _, r := range userRealms {
		access, err := s.policies.GetPolicies(user.Id, r.RealmId)
		if err != nil {
			return nil, err
		}

		user.Permissions[r.RealmId] = access.Perms
	}
	// access, err := s.policies.GetPolicies(user.Id, realm)
	// if err != nil {
	// 	return nil, err
	// }
	// user.Role = access.Role
	// user.Permissions = map[string][]string{
	// 	access.Domain: access.Perms,
	// }

	user.AccessToken = res.AccessToken
	user.RefreshToken = res.RefreshToken

	return user, nil
}

func (s *SessionService) DecodeAccessToken(ctx context.Context, token string) (*models.User, error) {
	//TODO расшифровку токена тоже лучше делать здесь, а в keycloak
	_, claims, err := s.keycloak.Client.DecodeAccessToken(ctx, token, s.keycloak.Realm)
	if err != nil {
		return nil, fmt.Errorf("failed to decode access token. error: %w", err)
	}

	serviceName := os.Getenv("SERVICE_ID")

	user := &models.User{}
	var username, userId string
	c := *claims
	access, ok := c["realm_access"]
	if ok {
		a := access.(map[string]interface{})["roles"]
		roles := a.([]interface{})
		for _, r := range roles {
			if strings.Contains(r.(string), serviceName) {
				break
			}
		}
	}

	u, ok := c["preferred_username"]
	if ok {
		username = u.(string)
	}

	uId, ok := c["sub"]
	if ok {
		userId = uId.(string)
	}

	user.Id = userId
	user.Name = username

	return user, nil
}
