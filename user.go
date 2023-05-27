// https://schema.ocsf.io/objects/user
package main

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	AccountType      string   `json:"account_type" validate:"omitempty"`
	AccountTypeID    uint8    `json:"account_type_id" validate:"omitempty"`
	AccountUID       string   `json:"account_uid" validate:"omitempty"`
	Domain           string   `json:"domain" validate:"omitempty"`
	EmailAddress     string   `json:"email_addr" validate:"email,omitempty"`
	Groups           *[]Group `json:"groups" validate:"omitempty"`
	Name             string   `json:"name" validate:"omitempty"`
	OrgID            string   `json:"org_uid" validate:"omitempty"`
	Type             string   `json:"type" validate:"omitempty"`
	TypeID           uint8    `json:"type_id" validate:"omitempty"`
	UniqueUserID     string   `json:"uuid" validate:"omitempty"`
	UserCredentialID string   `json:"credential_uid" validate:"omitempty"`
	UserID           string   `json:"uid" validate:"omitempty"`
}

func ValidateUser(user *User) (*User, error) {
	err := validator.New().Struct(user)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return user, nil
}
