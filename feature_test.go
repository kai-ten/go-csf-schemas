package gcs

import "testing"

func TestValidFeatureObject(t *testing.T) {
	feature := &Feature{
		FeatureID:      "123",
		FeatureName:    "name",
		FeatureVersion: "1",
	}
	_, err := ValidateFeature(feature)
	if err != nil {
		t.Fatalf("Feature should have been valid: %v", err)
	}
}

func BenchmarkValidFeatureObject(b *testing.B) {
	b.ReportAllocs()

	feature := &Feature{
		FeatureID:      "123",
		FeatureName:    "name",
		FeatureVersion: "1",
	}
	ValidateFeature(feature)
}
