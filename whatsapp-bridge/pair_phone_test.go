package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWritePairCodeIsOwnerOnlyAndReplacesThePreviousCode(t *testing.T) {
	path := filepath.Join(t.TempDir(), "store", "pair_code.txt")
	if err := writePairCode(path, "AAAA-BBBB"); err != nil {
		t.Fatal(err)
	}
	if err := writePairCode(path, "CCCC-DDDD"); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "CCCC-DDDD\n" {
		t.Fatalf("code = %q", got)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("mode = %v, want 0600", info.Mode().Perm())
	}
}

func TestPairPhoneFromEnvIsEmptyWhenUnset(t *testing.T) {
	t.Setenv(pairPhoneEnv, "  ")
	if got := pairPhoneFromEnv(); got != "" {
		t.Fatalf("got %q", got)
	}
}
