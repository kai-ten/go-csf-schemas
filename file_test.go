package gcs

import (
	"testing"
	"time"
)

func TestValidFileObject(t *testing.T) {
	// contains many more optional fields
	file := &File{
		Name:   "name",
		TypeID: UInteger8(0),
	}

	_, err := ValidateFile(file)
	if err != nil {
		t.Fatalf("File should have been valid: %v", err)
	}
}

func TestInvalidNameFileObject(t *testing.T) {
	time := time.Now()
	// contains many more optional fields
	file := &File{
		AccessedTime: &time,
		Name:         "",
		TypeID:       UInteger8(1),
	}

	_, err := ValidateFile(file)
	if err == nil {
		t.Fatalf("File Name should have been invalid: %v", err)
	}
}

func TestInvalidTypeIDFileObject(t *testing.T) {
	time := time.Now()
	// contains many more optional fields
	file := &File{
		AccessedTime: &time,
		Name:         "name",
	}

	_, err := ValidateFile(file)
	if err == nil {
		t.Fatalf("File TypeID should have been invalid: %v", err)
	}
}

func BenchmarkValidFileObject(b *testing.B) {
	b.ReportAllocs()

	// contains many more optional fields
	file := &File{
		Name:   "name",
		TypeID: UInteger8(2),
	}

	for n := 0; n < b.N; n++ {
		_, err := ValidateFile(file)
		if err != nil {
			b.Fatalf("File object is invalid: %v", err)
		}
	}
}
