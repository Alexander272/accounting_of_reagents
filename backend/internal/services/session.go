package services

type SessionService struct {
	// keycloak *auth.KeycloakClient
	// role     Role
	// filter   DefaultFilter
}

func NewSessionService() *SessionService {
	return &SessionService{
		// keycloak: keycloak,
		// role:     role,
		// filter:   filter,
	}
}

type Session interface {
	// SignIn(ctx context.Context, u models.SignIn) (*models.User, error)
	// SignOut(ctx context.Context, refreshToken string) error
	// Refresh(ctx context.Context, refreshToken string) (*models.User, error)
	// DecodeAccessToken(ctx context.Context, token string) (*models.User, error)
}
