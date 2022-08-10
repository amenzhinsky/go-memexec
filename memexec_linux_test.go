//go:build linux
// +build linux

package memexec

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestGenerateDynLink(t *testing.T) {
	e := exec.Command("go", "run", "../cmd/memexec-gen", "/usr/bin/perl")
	e.Dir = "testdata"
	b, err := e.CombinedOutput()
	if err != nil {
		t.Fatal(outputOrErr(b, err))
	}

	e = exec.Command("go", "run", ".")
	e.Dir = "testdata"
	e.Stdin = bytes.NewBufferString("print 42")
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
