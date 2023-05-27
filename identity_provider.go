// https://schema.ocsf.io/objects/idp
package main

import (
	"github.com/go-playground/validator/v10"
)

type IdentityProvider struct {
	Name     string `json:"name" validate:"omitempty"`
	UniqueID string `json:"uid" validate:"omitempty"`
}

func ValidateIdentityProvider(idp *IdentityProvider) (*IdentityProvider, error) {
	err := validator.New().Struct(idp)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return idp, nil
}
