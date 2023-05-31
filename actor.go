// https://schema.ocsf.io/objects/actor
package gcs

import "github.com/go-playground/validator/v10"

type Actor struct {
	AuthorizationInformation *[]AuthorizationInformation `json:"authorizations" validate:"omitempty"`
	IdentityProvider         *IdentityProvider           `json:"idp" validate:"omitempty"`
	InvokedBy                string                      `json:"invoked_by" validate:"omitempty"`
	Process                  *Process                    `json:"process" validate:"omitempty"`
	User                     *User                       `json:"user" validate:"omitempty"`
	UserSession              *Session                    `json:"session" validate:"omitempty"`
}

func ValidateActor(actor *Actor) (*Actor, error) {
	err := validator.New().Struct(actor)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return actor, nil
}
