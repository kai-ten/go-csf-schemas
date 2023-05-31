// https://schema.ocsf.io/objects/fingerprint
package main

import "github.com/go-playground/validator/v10"

type Fingerprint struct {
	Algorithm   string `json:"algorithm" validate:"omitempty"`
	AlgorithmID *uint8 `json:"algorithm_id" validate:"required"`
	Value       string `json:"value" validate:"required"`
}

func ValidateFingerprint(fingerprint *Fingerprint) (*Fingerprint, error) {
	err := validator.New().Struct(fingerprint)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return fingerprint, nil
}
