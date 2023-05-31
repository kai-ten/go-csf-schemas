package gcs

import (
	"testing"
	"time"
)

func TestValidProcessObject(t *testing.T) {
	ts := time.Now()
	process := &Process{
		CommandLine:    "ssh user@0.0.0.0",
		CreatedTimed:   &ts,
		File:           nil,
		Integrity:      "Windows",
		IntegrityLevel: 99,
		Lineage:        []string{"/usr/sbin/sshd"},
		LoadedModules:  []string{"Docker"},
		Name:           "Docker",
		ParentProcess:  nil,
		ProcessID:      3899,
		ProcessUID:     "Docker",
		Sandbox:        "default_ps",
		TerminatedTime: &ts,
		ThreadID:       123,
		User:           nil,
		UserSession:    nil,
	}

	_, err := ValidateProcess(process)
	if err != nil {
		t.Fatalf("Process object should be valid: %v", err)
	}
}

func BenchmarkValidProcessObject(b *testing.B) {
	b.ReportAllocs()

	ts := time.Now()
	process := &Process{
		CommandLine:    "ssh user@0.0.0.0",
		CreatedTimed:   &ts,
		File:           nil,
		Integrity:      "Windows",
		IntegrityLevel: 99,
		Lineage:        []string{"/usr/sbin/sshd"},
		LoadedModules:  []string{"Docker"},
		Name:           "Docker",
		ParentProcess:  nil,
		ProcessID:      3899,
		ProcessUID:     "Docker",
		Sandbox:        "default_ps",
		TerminatedTime: &ts,
		ThreadID:       123,
		User:           nil,
		UserSession:    nil,
	}

	ValidateProcess(process)
}
