package auth

import (
	"context"
	"log/slog"

	"github.com/Nerzal/gocloak/v13"
)

type KeycloakClient struct {
	Client       *gocloak.GoCloak // keycloak client
	ClientId     string           // clientId specified in Keycloak
	ClientSecret string           // client secret specified in Keycloak
	Realm        string           // realm specified in Keycloak
}

type Deps struct {
	Url       string
	ClientId  string
	Realm     string
	AdminName string
	AdminPass string
}

func NewKeycloakClient(deps Deps) *KeycloakClient {
	client := gocloak.NewClient(deps.Url)

	ctx := context.Background()

	token, err := client.LoginAdmin(ctx, deps.AdminName, deps.AdminPass, deps.Realm)
	if err != nil {
		slog.Error("failed to login admin to keycloak.", slog.String("error", err.Error()))
	}

	// store, err := client.GetKeyStoreConfig()
	// store.ActiveKeys.RS256

	clients, err := client.GetClients(ctx, token.AccessToken, deps.Realm, gocloak.GetClientsParams{ClientID: &deps.ClientId})
	if err != nil {
		slog.Error("failed to get clients to keycloak.", slog.String("error", err.Error()))
	}
	//logger.Debug(clients)

	secret := *clients[0].Secret

	return &KeycloakClient{
		Client:       client,
		ClientId:     deps.ClientId,
		ClientSecret: secret,
		Realm:        deps.Realm,
	}
}
