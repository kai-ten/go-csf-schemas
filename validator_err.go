package main

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

func ValidatorErrLog(err error) {
	for _, err := range err.(validator.ValidationErrors) {
		log.Errorf("field: %v, tag: %v, type: %v, value: %v", err.StructField(), err.Tag(), err.Type(), err.Value())
	}
}
