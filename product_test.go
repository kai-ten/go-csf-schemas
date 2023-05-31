package gcs

import "testing"

func TestValidProductObject(t *testing.T) {
	product := &Product{
		Feature:        nil,
		Language:       "en",
		ProductID:      "123",
		ProductName:    "Falcon Sensor",
		ProductPath:    "/usr/local/bin",
		ProductVersion: "1",
		VendorName:     "Crowdstrike",
	}
	_, err := ValidateProduct(product)
	if err != nil {
		t.Fatalf("Product should have been valid: %v", err)
	}
}

func TestInvalidProductNameProductObject(t *testing.T) {
	product := &Product{
		Feature:        nil,
		Language:       "en",
		ProductID:      "123",
		ProductName:    "",
		ProductPath:    "/usr/local/bin",
		ProductVersion: "1",
		VendorName:     "Crowdstrike",
	}
	_, err := ValidateProduct(product)
	if err == nil {
		t.Fatalf("Product ProductName should have been invalid: %v", err)
	}
}

func TestInvalidVendorNameProductObject(t *testing.T) {
	product := &Product{
		Feature:        nil,
		Language:       "en",
		ProductID:      "123",
		ProductName:    "Falcon Sensor",
		ProductPath:    "/usr/local/bin",
		ProductVersion: "1",
		VendorName:     "",
	}
	_, err := ValidateProduct(product)
	if err == nil {
		t.Fatalf("Product VendorName should have been valid: %v", err)
	}
}

func BenchmarkValidProductObject(b *testing.B) {
	b.ReportAllocs()

	product := &Product{
		Feature:        nil,
		Language:       "en",
		ProductID:      "123",
		ProductName:    "Falcon Sensor",
		ProductPath:    "/usr/local/bin",
		ProductVersion: "1",
		VendorName:     "Crowdstrike",
	}
	ValidateProduct(product)
}
