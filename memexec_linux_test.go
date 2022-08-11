//go:build linux

package memexec

import (
	"os/exec"
	"testing"
)

func TestGenLinux(t *testing.T) {
	e := exec.Command("go", "run", "../cmd/memexec-gen", "/usr/bin/python3")
	e.Dir = "testdata"
	_ = runCommand(t, e)

	e = exec.Command("go", "run", ".", "-c", "print(42)")
	e.Dir = "testdata"
	if have := runCommand(t, e); have != "42" {
		t.Fatalf("output mismatch:\n\thave: %s\n\twant: %s", have, "42")
	}
}
