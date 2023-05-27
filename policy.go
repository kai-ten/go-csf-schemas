// https://schema.ocsf.io/objects/policy
package main

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type Policy struct {
	Description string `json:"desc"`
	Group       *Group `json:"group" validate:"omitempty"`
	Name        string `json:"name" validate:"required"`
	UniqueID    string `json:"uid"`
	Version     string `json:"version"`
}

func ValidatePolicy(policy *Policy) (*Policy, error) {
	err := validator.New().Struct(policy)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("field: %v, tag: %v, type: %v, value: %v", err.Field(), err.Tag(), err.Type(), err.Value())
		}

		return nil, err
	}

	return policy, nil
}
