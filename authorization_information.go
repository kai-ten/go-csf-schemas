// https://schema.ocsf.io/objects/authorization
package main

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type AuthorizationInformation struct {
	AuthorizationDecision string  `json:"decision" validate:"omitempty"`
	Policy                *Policy `json:"policy" validate:"omitempty"`
}

func ValidateAuthorizationInformation(authInfo *AuthorizationInformation) (*AuthorizationInformation, error) {
	err := validator.New().Struct(authInfo)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("field: %v, tag: %v, type: %v, value: %v", err.Field(), err.Tag(), err.Type(), err.Value())
		}

		return nil, err
	}

	return authInfo, nil
}
