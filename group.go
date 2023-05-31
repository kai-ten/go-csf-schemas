// https://schema.ocsf.io/objects/group
package gcs

import (
	"github.com/go-playground/validator/v10"
)

type Group struct {
	AccountType string   `json:"type" validate:"omitempty"`
	Description string   `json:"desc" validate:"omitempty"`
	Name        string   `json:"name" validate:"required"`
	Privileges  []string `json:"privileges" validate:"omitempty"`
	UniqueID    string   `json:"unique_id" validate:"omitempty"`
}

func ValidateGroup(group *Group) (*Group, error) {
	err := validator.New().Struct(group)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return group, nil
}
