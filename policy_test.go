package main

import (
	"testing"
)

func TestValidNewPolicy(t *testing.T) {
	policy := &Policy{
		Description: "test desc",
		Name:        "test_name",
		UniqueID:    "123",
		Version:     "1",
	}

	policy, err := ValidatePolicy(policy)
	if err != nil {
		t.Fatalf("Unable to create and validate Group: %v", err)
	}

	t.Logf("Name = %v", policy.Name)
}

func TestErrorInvalidNewPolicy(t *testing.T) {
	policy := &Policy{
		Description: "test desc",
		Name:        "",
		UniqueID:    "123",
		Version:     "1",
	}

	_, err := ValidatePolicy(policy)
	if err == nil {
		t.Fatalf("Validator should have flagged this as an error: %v", err)
	}
}

func TestErrorValidGroupInPolicy(t *testing.T) {
	policy := &Policy{
		Description: "test desc",
		Name:        "",
		UniqueID:    "123",
		Version:     "1",
		Group:       &Group{Name: "test_name"},
	}

	_, err := ValidatePolicy(policy)
	if err == nil {
		t.Fatalf("Validator should have flagged this as an error: %v", err)
	}
}

func TestErrorInvalidGroupInPolicy(t *testing.T) {
	policy := &Policy{
		Description: "test desc",
		Name:        "test_name",
		UniqueID:    "123",
		Version:     "1",
		Group:       &Group{},
	}

	_, err := ValidatePolicy(policy)
	if err == nil {
		t.Fatalf("Validator should have flagged this as an error: %v", err)
	}
}

func BenchmarkPolicyValidation(b *testing.B) {
	b.ReportAllocs()

	policy := &Policy{
		Description: "test desc",
		Name:        "test_name",
		UniqueID:    "123",
		Version:     "1",
	}

	ValidatePolicy(policy)
}
