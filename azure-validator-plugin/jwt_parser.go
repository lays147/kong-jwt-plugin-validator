package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type Validation struct {
	IsValid       bool
	InvalidReason string
	RawJwt        string
}

func extractBearerToken(token string) (string, error) {
	if len(token) < 7 || !strings.EqualFold(token[:7], "Bearer ") {
		return "", errors.New("invalid bearer token")
	}
	return token[7:], nil
}

func IsJwtValid(tokenString string, tenantId string, cliendId string) Validation {
	keyFunc, err := GenerateKeyFunc(tenantId)

	if err != nil {
		log.Error().Err(err).Msg("Error creating keyfunc.")
		return Validation{
			IsValid:       false,
			InvalidReason: "Invalid TenantId. Unable to retrieve jwks keys for verification of token.",
		}
	}

	extractToken, err := extractBearerToken(tokenString)
	if err != nil {
		return Validation{
			IsValid:       false,
			InvalidReason: err.Error(),
		}
	}

	token, err := jwt.Parse(extractToken, keyFunc.Keyfunc, jwt.WithAudience(fmt.Sprintf("api://%s", cliendId)))

	switch {
	case token.Valid:
		return Validation{IsValid: true, InvalidReason: "", RawJwt: token.Raw}

	case errors.Is(err, jwt.ErrTokenInvalidAudience):
		return Validation{IsValid: false, InvalidReason: "Invalid audience. Check your client id value and see if matches the environment that you are targeting."}
	case errors.Is(err, jwt.ErrTokenMalformed):
		return Validation{IsValid: false, InvalidReason: "Malformed token."}
	case errors.Is(err, jwt.ErrTokenExpired):
		return Validation{IsValid: false, InvalidReason: "Token expired."}
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return Validation{IsValid: false, InvalidReason: "Invalid signature."}

	default:
		{
			log.Warn().Err(err).Msg("Unable to verify token.")
			return Validation{IsValid: false, InvalidReason: "Unable to verify token"}
		}
	}
}
