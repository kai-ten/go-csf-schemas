// https://schema.ocsf.io/objects/file
package main

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// WARN: ParentFolder and Path should be of type "Path Name" according to spec, does not exist rn

type File struct {
	AccessedTime       *time.Time             `json:"accessed_time" validate:"omitempty"`
	Accessor           *User                  `json:"accessor" validate:"omitempty"`
	Attributes         uint                   `json:"attributes" validate:"omitempty"`
	CompanyName        string                 `json:"company_name" validate:"omitempty"`
	Confidentiality    string                 `json:"confidentiality" validate:"omitempty"`
	ConfidentialityID  uint8                  `json:"confidentiality_id" validate:"omitempty"`
	CreatedTime        *time.Time             `json:"created_time" validate:"omitempty"`
	Creator            *User                  `json:"creator" validate:"omitempty"`
	Description        string                 `json:"desc" validate:"omitempty"`
	DigitalSignature   *DigitalSignature      `json:"signature" validate:"omitempty"`
	ExtendedAttributes map[string]interface{} `json:"xattributes" validate:"omitempty"`
	Fingerprints       *[]Fingerprint         `json:"fingerprints" validate:"omitempty"`
	MIMEType           string                 `json:"mime_type" validate:"omitempty"`
	Modifier           *User                  `json:"modifier" validate:"omitempty"`
	Name               string                 `json:"name" validate:"required"`
	Owner              *User                  `json:"owner" validate:"omitempty"`
	ParentFolder       string                 `json:"parent_folder" validate:"omitempty"`
	Path               string                 `json:"path" validate:"omitempty"`
	Product            *Product               `json:"product" validate:"omitempty"`
	SecurityDescriptor string                 `json:"security_descriptor" validate:"omitempty"`
	Size               int64                  `json:"size" validate:"omitempty"`
	System             bool                   `json:"is_system" validate:"omitempty"`
	Type               string                 `json:"type" validate:"omitempty"`
	TypeID             uint8                  `json:"type_id" validate:"required"`
	UniqueID           string                 `json:"uid" validate:"omitempty"`
	Version            string                 `json:"version" validate:"omitempty"`
}

func ValidateFile(file *File) (*File, error) {
	err := validator.New().Struct(file)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return file, nil
}
