package gcs

import (
	"testing"
)

func TestValidAuthorizationInformationObject(t *testing.T) {
	authInfo := &AuthorizationInformation{
		AuthorizationDecision: "allow",
		Policy: &Policy{
			Description: "test desc",
			Name:        "test_name",
			UniqueID:    "123",
			Version:     "1",
			Group:       &Group{Name: "test_name"},
		},
	}

	authInfo, err := ValidateAuthorizationInformation(authInfo)
	if err != nil {
		t.Fatalf("Unable to create and validate Group: %v", err)
	}

	t.Logf("Name = %v", authInfo.AuthorizationDecision)
}

func TestErrorInvalidAuthorizationInformationObject(t *testing.T) {
	authInfo := &AuthorizationInformation{
		AuthorizationDecision: "allow",
		Policy: &Policy{
			Description: "test desc",
			Name:        "",
			UniqueID:    "123",
			Version:     "1",
			Group:       &Group{Name: "test_name"},
		},
	}

	_, err := ValidateAuthorizationInformation(authInfo)
	if err == nil {
		t.Fatalf("Validator should have flagged this as an error: %v", err)
	}
}

func BenchmarkAuthorizationInformationValidation(b *testing.B) {
	b.ReportAllocs()

	authInfo := &AuthorizationInformation{
		AuthorizationDecision: "allow",
		Policy: &Policy{
			Description: "test desc",
			Name:        "test_name",
			UniqueID:    "123",
			Version:     "1",
			Group:       &Group{Name: "test_name"},
		},
	}

	for n := 0; n < b.N; n++ {
		_, err := ValidateAuthorizationInformation(authInfo)
		if err != nil {
			b.Fatalf("AuthorizationInformation object is invalid: %v", err)
		}
	}
}
