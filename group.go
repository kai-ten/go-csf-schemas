// https://schema.ocsf.io/objects/group
package main

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
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
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("field: %v, tag: %v, type: %v, value: %v", err.Field(), err.Tag(), err.Type(), err.Value())
		}

		return nil, err
	}

	return group, nil
}
