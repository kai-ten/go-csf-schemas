// https://schema.ocsf.io/objects/feature
package main

import "github.com/go-playground/validator/v10"

type Feature struct {
	FeatureID      string `json:"uid" validate:"omitempty"`
	FeatureName    string `json:"name" validate:"omitempty"`
	FeatureVersion string `json:"version" validate:"omitempty"`
}

func ValidateFeature(feature *Feature) (*Feature, error) {
	err := validator.New().Struct(feature)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return feature, nil
}
