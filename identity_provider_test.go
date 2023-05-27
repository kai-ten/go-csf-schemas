package main

import (
	"testing"
)

func TestValidIdentityProviderObject(t *testing.T) {
	idp := &IdentityProvider{
		Name:     "test_name",
		UniqueID: "123",
	}

	idp, err := ValidateIdentityProvider(idp)
	if err != nil {
		t.Fatalf("Unable to create and validate Group: %v", err)
	}

	t.Logf("Name = %v", idp.Name)
}

func BenchmarkIdentityProviderValidation(b *testing.B) {
	b.ReportAllocs()

	idp := &IdentityProvider{
		Name:     "test_name",
		UniqueID: "123",
	}

	ValidateIdentityProvider(idp)
}
