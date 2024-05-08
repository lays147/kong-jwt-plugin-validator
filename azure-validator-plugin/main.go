package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const Version = "1.0.0"
const Priority = 1010
const PluginName = "azure-jwt-validator-plugin"

type AzureAuthConfig struct {
	TenantId string `json:"tenant_id"`
	ClientId string `json:"client_id"`
}

func New() interface{} {
	return &AzureAuthConfig{}
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
}

func main() {
	//nolint:errcheck
	server.StartServer(New, Version, Priority)
}

func (conf AzureAuthConfig) Access(kong *pdk.PDK) {
	bearerToken, err := kong.Request.GetHeader("Authorization")

	if err != nil || bearerToken == "" {
		payload := PreparePayload(PluginName, 400, "", "Missing Authorization header", "")
		kong.Response.Exit(400, payload, JsonHeader())
		return
	}

	isJwtValid := IsJwtValid(bearerToken, conf.TenantId, conf.ClientId)

	if !isJwtValid.IsValid {
		payload := PreparePayload(PluginName, 401, "", isJwtValid.InvalidReason, "")
		kong.Response.Exit(401, payload, JsonHeader())
		return
	}
}
