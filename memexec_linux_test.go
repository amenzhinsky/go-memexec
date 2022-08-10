//go:build linux
// +build linux

package memexec

import (
	"os/exec"
	"strings"
	"testing"
)

func TestGenerateDynLink(t *testing.T) {
	_ = execute(t, "go", "run", "../cmd/memexec-gen", "/usr/bin/python3")
	if have := execute(t, "go", "run", ".", "-c", "print(42)"); have != "42" {
		t.Fatalf("output mismatch:\n\thave: %s\n\twant: %s", have, "42")
	}
}

func execute(t *testing.T, cmd string, args ...string) string {
	exe := exec.Command(cmd, args...)
	exe.Dir = "testdata"
	b, err := exe.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(b))
}
