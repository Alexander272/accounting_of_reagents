package services

import (
	"github.com/Alexander272/accounting_of_reagents/backend/internal/config"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/events"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/repository"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/auth"
)

type Services struct {
	Permissions
	Roles
	Realm
	UserRealm
	Users
	Session
	AccessPolices

	AmountType
	ReagentType
	Reagent
	Spending
	Extending
	Note

	Scheduler
}

type Deps struct {
	Repos    *repository.Repository
	Keycloak *auth.KeycloakClient
	Conf     *config.Config
}

func NewServices(deps Deps) *Services {
	transaction := NewTransactionManager(deps.Repos.Transaction)

	updatePolicyEvent := &events.PolicyEventManager{}

	permissions := NewPermissionService(deps.Repos.Permissions, transaction, updatePolicyEvent)
	roleHierarchy := NewRoleHierarchyService(deps.Repos.RoleHierarchy)
	role := NewRoleService(&RoleDeps{
		Repo:        deps.Repos.Roles,
		Hierarchy:   roleHierarchy,
		Permissions: permissions,
		EventBus:    updatePolicyEvent,
		TM:          transaction,
	})

	realm := NewRealmService(deps.Repos.Realm)
	userRealm := NewUserRealmService(deps.Repos.UserRealm, realm, role, updatePolicyEvent)
	user := NewUserService(&UserDeps{
		Repo:      deps.Repos.User,
		TM:        transaction,
		Keycloak:  deps.Keycloak,
		UserRealm: userRealm,
		EventBus:  updatePolicyEvent,
	})

	cacheSvc := NewSessionCacheService(deps.Repos.SessionCache)

	adapter := NewAdapter(&AdapterDeps{Permissions: permissions, Users: user, RoleHierarchy: roleHierarchy})
	policies := NewAccessPoliciesService(&PoliciesDeps{
		Conf:     deps.Conf.Casbin,
		Adapter:  adapter,
		EventBus: updatePolicyEvent,
		Cache:    cacheSvc,
	})

	session := NewSessionService(deps.Keycloak, policies, userRealm, user, cacheSvc)

	most := NewMostService(deps.Conf.Bot.Url, deps.Conf.Bot.ChannelId)

	amountType := NewAmountTypeService(deps.Repos.AmountType)
	reagentType := NewReagentTypeService(deps.Repos.ReagentType)
	reagent := NewReagentService(deps.Repos.Reagent, reagentType, most)
	spending := NewSpendingService(deps.Repos.Spending, reagent, most)
	extending := NewExtendingService(deps.Repos.Extending, reagent)
	note := NewNoteService(deps.Repos.Notes)

	scheduler := NewSchedulerService(reagent)

	return &Services{
		Permissions: permissions,
		Roles:       role,
		Session:     session,
		Realm:       realm,
		UserRealm:   userRealm,
		Users:       user,

		AccessPolices: policies,

		AmountType:  amountType,
		ReagentType: reagentType,
		Reagent:     reagent,
		Spending:    spending,
		Extending:   extending,
		Note:        note,

		Scheduler: scheduler,
	}
}
