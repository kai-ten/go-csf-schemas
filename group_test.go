package gcs

import (
	"testing"
)

func TestValidNewGroup(t *testing.T) {
	group := &Group{
		AccountType: "test_type",
		Description: "test desc",
		Name:        "test_name",
		Privileges:  []string{"test privs"},
		UniqueID:    "123",
	}

	group, err := ValidateGroup(group)
	if err != nil {
		t.Fatalf("Unable to create and validate Group: %v", err)
	}

	t.Logf("Name = %v", group.Name)
}

func TestErrorInvalidNewGroup(t *testing.T) {
	group := &Group{
		AccountType: "test_type",
		Description: "test desc",
		Name:        "",
		Privileges:  []string{"test privs"},
		UniqueID:    "123",
	}

	_, err := ValidateGroup(group)
	if err == nil {
		t.Fatalf("Validator should have flagged this as an error: %v", err)
	}
}

func BenchmarkGroupValidation(b *testing.B) {
	b.ReportAllocs()

	group := &Group{
		AccountType: "test_type",
		Description: "test desc",
		Name:        "test_name",
		Privileges:  []string{"test privs"},
		UniqueID:    "123",
	}

	for n := 0; n < b.N; n++ {
		_, err := ValidateGroup(group)
		if err != nil {
			b.Fatalf("Group object is invalid: %v", err)
		}
	}
}
