package memexec

import (
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestMemexec(t *testing.T) {
	t.Parallel()

	p, err := exec.LookPath("echo")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}

	cmd, err := Command(b, "foo", "bar")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = cmd.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	o, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}

	if string(o) != "foo bar\n" {
		t.Errorf("output = %q, want %q", string(o), "foo bar\n")
	}
}
