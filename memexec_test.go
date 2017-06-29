package memexec

import (
	"io/ioutil"
	"os/exec"
	"strings"
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
		if err != nil {
			return
		}

		if err = m.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	for want, args := range map[string][]string{
		"foo bar": {"-n", "foo", "bar"},
		"foo baz": {"-n", "foo", "baz"},
	} {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
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
