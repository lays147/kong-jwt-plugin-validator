package main

import (
	"testing"

	"github.com/Kong/go-pdk/test"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
)

func TestPluginMissingAuthorizationHeader(t *testing.T) {
	ja := jsonassert.New(t)
	env, err := test.New(t, test.Request{
		Method: "GET",
		Url:    "/",
	})
	assert.NoError(t, err)

	env.DoHttps(&AzureAuthConfig{TenantId: "xxx"})
	assert.Equal(t, 400, env.ClientRes.Status)
	ja.Assertf(string(env.ClientRes.Body), `{"statusCode": 400, "error": "Missing Authorization header", "plugin": "azure-jwt-validator-plugin"}`)
}

func TestPluginWithInvalidTenantId(t *testing.T) {
	ja := jsonassert.New(t)
	env, err := test.New(t, test.Request{
		Method:  "GET",
		Url:     "/",
		Headers: map[string][]string{"Authorization": {"Bearer xxx"}},
	})
	assert.NoError(t, err)

	env.DoHttps(&AzureAuthConfig{TenantId: "xxx"})
	assert.Equal(t, 401, env.ClientRes.Status)
	ja.Assertf(string(env.ClientRes.Body), `{"statusCode": 401, "error": "Invalid TenantId. Unable to retrieve jwks keys for verification of token.", "plugin": "azure-jwt-validator-plugin"}`)
}
