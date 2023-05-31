// https://schema.ocsf.io/objects/authorization
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type AuthorizationInformation struct {
	AuthorizationDecision string  `json:"decision" validate:"omitempty"`
	Policy                *Policy `json:"policy" validate:"omitempty"`
}

func ValidateAuthorizationInformation(authInfo *AuthorizationInformation) (*AuthorizationInformation, error) {
	err := validator.New().Struct(authInfo)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return authInfo, nil
}
