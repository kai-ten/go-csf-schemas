package main

import (
	"testing"
)

func TestValidNewGroup(t *testing.T) {
	group, err := NewGroup("test_type", "test desc", "test_name", []string{"test privs"}, "123")
	if err != nil {
		t.Fatalf("Unable to create and validate Group: %v", err)
	}

	t.Logf("Name = %v", group.Name)
}

func TestErrorInvalidNewGroup(t *testing.T) {
	_, err := NewGroup("test_type", "test desc", "", []string{"test privs"}, "123")
	if err == nil {
		t.Fatalf("Validator should have flagged this as an error: %v", err)
	}
}
