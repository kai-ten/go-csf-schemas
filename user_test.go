package gcs

import "testing"

func TestValidUserObject(t *testing.T) {
	user := &User{
		AccountType:      "test",
		AccountTypeID:    2,
		AccountUID:       "123",
		Domain:           "Active Directory",
		EmailAddress:     "name@gocsf.com",
		Groups:           nil,
		Name:             "name",
		OrgID:            "345",
		Type:             "AD User",
		TypeID:           1,
		UniqueUserID:     "456",
		UserCredentialID: "",
		UserID:           "456",
	}

	user, err := ValidateUser(user)
	if err != nil {
		t.Fatalf("User should have been valid: %v", err)
	}

	t.Logf("User email: %v", user.EmailAddress)
}

func TestInvalidUserEmail(t *testing.T) {
	user := &User{
		EmailAddress: "invalid",
	}

	user, err := ValidateUser(user)
	if err == nil {
		t.Fatalf("User validation should have failed: %v", user.EmailAddress)
	}
}

func BenchmarkValidUserObject(b *testing.B) {
	b.ReportAllocs()

	user := &User{
		AccountType:      "test",
		AccountTypeID:    2,
		AccountUID:       "123",
		Domain:           "Active Directory",
		EmailAddress:     "name@gocsf.com",
		Groups:           nil,
		Name:             "name",
		OrgID:            "345",
		Type:             "AD User",
		TypeID:           1,
		UniqueUserID:     "456",
		UserCredentialID: "",
		UserID:           "456",
	}
	ValidateUser(user)
}
