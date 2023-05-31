package gcs

import (
	"testing"
	"time"
)

func TestValidDigitalSignatureObject(t *testing.T) {
	time := time.Now()
	dg := &DigitalSignature{
		CompanyName:  "Microsoft Corporation",
		CreatedTimed: &time,
		DeveloperUID: "123",
		Fingerprints: nil,
		IssuerName:   "LibreSSL",
		SerialNumber: "345",
	}
	_, err := ValidateDigitalSignature(dg)
	if err != nil {
		t.Fatalf("Digital Signature should have been valid: %v", err)
	}
}

func TestInvalidDigitalSignatureObject(t *testing.T) {
	time := time.Now()
	dg := &DigitalSignature{
		CompanyName:  "",
		CreatedTimed: &time,
		DeveloperUID: "123",
		Fingerprints: nil,
		IssuerName:   "LibreSSL",
		SerialNumber: "345",
	}
	_, err := ValidateDigitalSignature(dg)
	if err == nil {
		t.Fatalf("Digital Signature should have been invalid: %v", err)
	}
}

func BenchmarkValidDigitalSignatureObject(b *testing.B) {
	b.ReportAllocs()

	time := time.Now()
	dg := &DigitalSignature{
		CompanyName:  "Microsoft Corporation",
		CreatedTimed: &time,
		DeveloperUID: "123",
		Fingerprints: nil,
		IssuerName:   "LibreSSL",
		SerialNumber: "345",
	}

	ValidateDigitalSignature(dg)
}
