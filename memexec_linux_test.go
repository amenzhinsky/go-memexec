//go:build linux

package memexec

import (
	"os/exec"
	"testing"
)

func TestGenLinux(t *testing.T) {
	e := exec.Command("go", "run", "../cmd/memexec-gen", "/usr/bin/python3")
	e.Dir = "testdata"
	b, err := e.CombinedOutput()
	if err != nil {
		t.Fatal(outputOrErr(b, err))
	}

	e = exec.Command("go", "run", ".", "-c", "print(42)")
	e.Dir = "testdata"
	b, err = e.CombinedOutput()
	if err != nil {
		t.Fatal(outputOrErr(b, err))
	}
	if have := string(b); have != "42" {
		t.Fatalf("output mismatch:\n\thave: %s\n\twant: %s", have, "42")
	}
}

func outputOrErr(b []byte, err error) string {
	if b != nil {
		return string(b)
	}
	return err.Error()
}
