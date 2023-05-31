// https://schema.ocsf.io/objects/process
package gcs

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Process struct {
	CommandLine        string                 `json:"cmd_line" validate:"omitempty"`
	CreatedTimed       *time.Time             `json:"created_time" validate:"omitempty"`
	ExtendedAttributes map[string]interface{} `json:"xattributes" validate:"omitempty"`
	File               *File                  `json:"file" validate:"omitempty"`
	Integrity          string                 `json:"integrity" validate:"omitempty"`
	IntegrityLevel     uint8                  `json:"integrity_id" validate:"omitempty"`
	Lineage            []string               `json:"lineage" validate:"omitempty"`
	LoadedModules      []string               `json:"loaded_modules" validate:"omitempty"`
	Name               string                 `json:"name" validate:"omitempty"`
	ParentProcess      *Process               `json:"parent_process" validate:"omitempty"`
	ProcessID          uint                   `json:"pid" validate:"omitempty"`
	ProcessUID         string                 `json:"uid" validate:"omitempty"`
	Sandbox            string                 `json:"sandbox" validate:"omitempty"`
	TerminatedTime     *time.Time             `json:"terminated_time" validate:"omitempty"`
	ThreadID           uint                   `json:"tid" validate:"omitempty"`
	User               *User                  `json:"user" validate:"omitempty"`
	UserSession        *Session               `json:"session" validae:"omitempty"`
}

func ValidateProcess(process *Process) (*Process, error) {
	err := validator.New().Struct(process)
	if err != nil {
		ValidatorErrLog(err)
		return nil, err
	}

	return process, nil
}
