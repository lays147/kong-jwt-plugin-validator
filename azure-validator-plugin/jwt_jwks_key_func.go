package main

import (
	"encoding/json"
	"net/http"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/rs/zerolog/log"
)

var jwksCache = NewJWKSCache()

// GenerateKeyFunc returns a key func compatible with golang-jwt parser
// to be able to validate the signature of a JWT token based on the jwks information
// provided by AzureAD OpenId Information.
func GenerateKeyFunc(tenantId string) (keyfunc.Keyfunc, error) {
	jwksUri := retrieveJwksUri(tenantId)
	return keyfunc.NewDefault([]string{jwksUri})
}

// retrieveJwksUri retrieves the jwks uri from AzureAD OpenId Information
// it caches the uri in memory, to avoid making web requests on every authorization request.
func retrieveJwksUri(tenantId string) string {
	discoveryUrl := "https://login.microsoftonline.com/" + tenantId + "/v2.0/.well-known/openid-configuration"
	const key = "jkws_uri"

	if uri, found := jwksCache.Get(key); found {
		return uri.(string)
	}
	//nolint:gosec
	resp, err := http.Get(discoveryUrl)
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving jwks uri.")
	}

	defer resp.Body.Close()

	var body struct {
		JwksUri string `json:"jwks_uri"`
	}

	//nolint:errcheck
	json.NewDecoder(resp.Body).Decode(&body)
	jwksCache.Set(key, body.JwksUri)

	return body.JwksUri
}
