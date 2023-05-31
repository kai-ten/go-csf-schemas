// https://schema.ocsf.io/objects/session
package gcs

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Session struct {
	CreatedTimed              *time.Time `json:"created_time" validate:"omitempty"`
	ExpirationTime            *time.Time `json:"expiration_time" validate:"omitempty"`
	IssuerDetails             string     `json:"issuer" validate:"omitempty"`
	MultiFactorAuthentication bool       `json:"mfa" validate:"omitempty"`
	UUID                      string     `json:"uuid" validate:"omitempty"`
	UniqueID                  string     `json:"uid" validate:"omitempty"`
	UserCredentialID          string     `json:"credential_uid" validate:"omitempty"`
}

func ValidateSession(session *Session) (*Session, error) {
	err := validator.New().Struct(session)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return session, nil
}
