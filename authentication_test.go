package gcs

import (
	"testing"
	"time"
)

func TestValidAuthenticationObject(t *testing.T) {
	ts := time.Now()
	auth := &Authentication{
		CategoryID: UInteger8(0), // handle 0..... go makes me sad.
		ClassID:    UInteger(0),
		EventTime:  &ts,
		TypeID:     UInteger(2),
		User:       &User{EmailAddress: "test@goscf.com"},
	}
	_, err := ValidateAuthentication(auth)
	if err != nil {
		t.Fatalf("Authentication should have been valid: %v", err)
	}
}
