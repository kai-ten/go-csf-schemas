// https://schema.ocsf.io/objects/product
package gcs

import "github.com/go-playground/validator/v10"

// WARN: ParentFolder and Path should be of type "Path Name" according to spec, does not exist rn

type Product struct {
	Feature        *Feature `json:"feature" validate:"omitempty"`
	Language       string   `json:"lang" validate:"omitempty"`
	ProductID      string   `json:"uid" validate:"omitempty"`
	ProductName    string   `json:"name" validate:"required"`
	ProductPath    string   `json:"path" validate:"omitempty"`
	ProductVersion string   `json:"version" validate:"omitempty"`
	VendorName     string   `json:"vendor_name" validate:"required"`
}

func ValidateProduct(product *Product) (*Product, error) {
	err := validator.New().Struct(product)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return product, nil
}
