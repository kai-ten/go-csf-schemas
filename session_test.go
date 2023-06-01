package gcs

import (
	"testing"
	"time"
)

func TestValidSessionObject(t *testing.T) {
	ts := time.Now()
	session := &Session{
		CreatedTimed:              &ts,
		ExpirationTime:            &ts,
		IssuerDetails:             "1",
		MultiFactorAuthentication: true,
		UUID:                      "123",
		UniqueID:                  "234",
		UserCredentialID:          "456",
	}
	_, err := ValidateSession(session)
	if err != nil {
		t.Fatalf("Session object should have been valid: %v", err)
	}
}

func BenchmarkValidSessionObject(b *testing.B) {
	b.ReportAllocs()

	ts := time.Now()
	session := &Session{
		CreatedTimed:              &ts,
		ExpirationTime:            &ts,
		IssuerDetails:             "1",
		MultiFactorAuthentication: true,
		UUID:                      "123",
		UniqueID:                  "234",
		UserCredentialID:          "456",
	}

	for n := 0; n < b.N; n++ {
		_, err := ValidateSession(session)
		if err != nil {
			b.Fatalf("Session object is invalid: %v", err)
		}
	}
}
