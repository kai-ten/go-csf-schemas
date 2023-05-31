// https://schema.ocsf.io/classes/authentication?extensions=
package main

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// DestinationEndpoint *NetworkEndpoint `json:"dst_endpoint" validate:"required"`
// Enrichments *[]Enrichment `json:"enrichments" valiadte:"omitempty"`
// HTTPRequest *HTTPRequest `json:"http_request" validate:"omitempty"`
// Metadata *Metadata `json:"metadata" validate:"omitempty"`
// Observables *[]Observable `json:"observables" validate:"omitempty"`
// SourceEndpoint NetworkEndpoint `json:"network_endpoint" validate:"omitempty"`

type Authentication struct {
	Activity                  string                 `json:"activity_name" validate:"omitempty"`
	ActivityID                uint8                  `json:"activity_id" validate:"omitempty"`
	Actor                     *Actor                 `json:"actor" validate:"omitempty"`
	AuthProtocol              string                 `json:"auth_protocol" validate:"omitempty"`
	AuthProtocolID            uint8                  `json:"auth_protocol_id" validate:"omitempty"`
	Category                  string                 `json:"category_name" validate:"omitempty"`
	CategoryID                *uint8                 `json:"category_uid" validate:"required"`
	Class                     string                 `json:"class_name" validate:"omitempty"`
	ClassID                   *uint                  `json:"class_uid" validate:"required"`
	CleartextCredentials      bool                   `json:"is_cleartext" validate:"omitempty"`
	Confidence                uint8                  `json:"confidence" validate:"omitempty"`
	Count                     uint                   `json:"count" validate:"omitempty"`
	Data                      map[string]interface{} `json:"data" validate:"omitempty"`
	Duration                  uint                   `json:"duration" validate:"omitempty"`
	EndTime                   *time.Time             `json:"end_time" validate:"omitempty"`
	EventTime                 *time.Time             `json:"time" validate:"required"`
	LogonProcess              *Process               `json:"logon_process" validate:"omitempty"`
	LogonTypeID               uint8                  `json:"logon_type_id" validate:"omitempty"`
	Message                   string                 `json:"message" validate:"omitempty"`
	MultiFactorAuthentication bool                   `json:"mfa" validate:"omitempty"`
	RawData                   string                 `json:"raw_data" validate:"omitempty"`
	Remote                    bool                   `json:"remote" validate:"omitempty"`
	Severity                  string                 `json:"severity" validate:"omitempty"`
	SeverityID                uint8                  `json:"severity_id" validate:"omitempty"`
	StartTime                 *time.Time             `json:"start_time" validate:"omitempty"`
	Status                    string                 `json:"status" validate:"omitempty"`
	StatusCode                string                 `json:"status_code" validate:"omitempty"`
	StatusDetails             string                 `json:"status_detail" validate:"omitempty"`
	StatusID                  uint8                  `json:"status_id" validate:"omitempty"`
	TimezoneOffset            int                    `json:"timezone_offset" validate:"omitempty"`
	TypeID                    *uint                  `json:"type_uid" validate:"required"`
	TypeName                  string                 `json:"type_name" validate:"omitempty"`
	UnmappedData              map[string]interface{} `json:"unmapped" validate:"omitempty"`
	User                      *User                  `json:"user" validate:"required"`
	UserSession               *Session               `json:"session" validate:"omitempty"`
}

func ValidateAuthentication(auth *Authentication) (*Authentication, error) {
	err := validator.New().Struct(auth)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return auth, nil
}
