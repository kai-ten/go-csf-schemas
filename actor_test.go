package gcs

import "testing"

func TestValidActorObject(t *testing.T) {
	actor := &Actor{
		AuthorizationInformation: &[]AuthorizationInformation{},
		IdentityProvider:         &IdentityProvider{},
		InvokedBy:                "",
		Process:                  &Process{},
		User:                     &User{EmailAddress: "test@email.com"},
		UserSession:              &Session{},
	}
	_, err := ValidateActor(actor)
	if err != nil {
		t.Fatalf("Actor object should have been valid: %v", err)
	}
}

func BenchmarkValidActorObject(b *testing.B) {
	b.ReportAllocs()

	actor := &Actor{
		AuthorizationInformation: &[]AuthorizationInformation{},
		IdentityProvider:         &IdentityProvider{},
		InvokedBy:                "",
		Process:                  &Process{},
		User:                     &User{},
		UserSession:              &Session{},
	}
	ValidateActor(actor)
}
