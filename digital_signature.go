// https://schema.ocsf.io/objects/digital_signature
package main

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type DigitalSignature struct {
	CompanyName  string         `json:"company_name" validate:"required"`
	CreatedTimed *time.Time     `json:"created_time" validate:"omitempty"`
	DeveloperUID string         `json:"developer_uid" validate:"omitempty"`
	Fingerprints *[]Fingerprint `json:"fingerprints" validate:"omitempty"`
	IssuerName   string         `json:"issuer_name" validate:"omitempty"`
	SerialNumber string         `json:"serial_number" validate:"omitempty"`
}

func ValidateDigitalSignature(ds *DigitalSignature) (*DigitalSignature, error) {
	err := validator.New().Struct(ds)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return ds, nil
}
