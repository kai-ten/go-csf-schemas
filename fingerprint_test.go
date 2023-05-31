package gcs

import "testing"

func TestValidFingerprintObject(t *testing.T) {
	fingerprint := &Fingerprint{
		Algorithm:   "",
		AlgorithmID: UInteger8(0),
		Value:       "ABC123==",
	}
	fingerprint, err := ValidateFingerprint(fingerprint)
	if err != nil {
		t.Fatalf("Fingerprint object should have been valid: %v", err)
	}

	t.Logf("Value: %v", fingerprint.Value)
}

func TestInvalidAlgorithmIDFingerprintObject(t *testing.T) {
	fingerprint := &Fingerprint{
		Algorithm: "",
		Value:     "ABC123==",
	}
	_, err := ValidateFingerprint(fingerprint)
	if err == nil {
		t.Fatalf("Fingerprint object should have been valid: %v", err)
	}
}

func TestInvalidValueFingerprintObject(t *testing.T) {
	fingerprint := &Fingerprint{
		Algorithm:   "",
		AlgorithmID: UInteger8(1),
		Value:       "",
	}
	_, err := ValidateFingerprint(fingerprint)
	if err == nil {
		t.Fatalf("Fingerprint object should have been valid: %v", err)
	}
}

func BenchmarkValidFingerprintObject(b *testing.B) {
	b.ReportAllocs()

	fingerprint := &Fingerprint{
		Algorithm:   "",
		AlgorithmID: UInteger8(2),
		Value:       "ABC123==",
	}
	ValidateFingerprint(fingerprint)
}
