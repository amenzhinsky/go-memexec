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

	m, err := New(b)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		e := m.Close()
		if err == nil && e != nil {
			t.Fatal(e)
		}
	}()

	for want, args := range map[string][]string{
		"foo bar": {"-n", "foo", "bar"},
		"foo baz": {"-n", "foo", "baz"},
	} {
		t.Run(want, func(t *testing.T) {
			c := m.Command(args...)
			o, err := c.Output()
			if err != nil {
				t.Fatal(err)
			}

			if string(o) != want {
				t.Errorf("Command(%#v...): output = %q, want %q", args, string(o), want)
			}
		})
	}
}
