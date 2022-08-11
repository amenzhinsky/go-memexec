package memexec

import (
	"context"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	exe := newEchoExec(t)
	defer func() {
		if err := exe.Close(); err != nil {
			t.Fatalf("close error: %s", err)
		}
	}()

	c := exe.Command("test")
	if have := runCommand(t, c); !strings.HasPrefix(have, "test") {
		t.Errorf("command output = %q doesn't contain %q", have, "test")
	}
}

func TestCommandContext(t *testing.T) {
	// the test is failing on windows, probably due to missing sleep command,
	// unfortunately I have no windows machines around to fix this
	if runtime.GOOS == "windows" {
		return
	}

	exe := newSleepExec(t)
	defer func() {
		if err := exe.Close(); err != nil {
			t.Fatalf("close error: %s", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	c := exe.CommandContext(ctx, "1")

	start := time.Now()
	if err := c.Start(); err != nil {
		t.Fatalf("start failed: %s", err)
	}
	if err := c.Wait(); err != nil && err.Error() != "signal: killed" {
		t.Fatalf("command execution failed: %s", err)
	}
	stop := time.Now()
	if delta := stop.Sub(start); delta > time.Millisecond*600 || delta < time.Millisecond*500 {
		t.Fatalf("unexpected command execution time: delta=%s", delta)
	}
}

func BenchmarkCommand(b *testing.B) {
	exe := newEchoExec(b)
	defer func() {
		if err := exe.Close(); err != nil {
			b.Fatalf("close error: %s", err)
		}
	}()

	for i := 0; i < b.N; i++ {
		cmd := exe.Command("test")
		if _, err := cmd.Output(); err != nil {
			b.Fatal(err)
		}
	}
}

func newExec(t testing.TB, name string) *Exec {
	path, err := exec.LookPath(name)
	if err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	exe, err := New(b)
	if err != nil {
		t.Fatal(err)
	}
	return exe
}

func newEchoExec(t testing.TB) *Exec {
	// lookup echo binary that is provided on all unix systems
	// and it's not a built-in opposed to `ls` and `type`
	return newExec(t, "echo")
}

func newSleepExec(t testing.TB) *Exec {
	return newExec(t, "sleep")
}

func runCommand(t *testing.T, cmd *exec.Cmd) string {
	t.Helper()
	b, err := cmd.CombinedOutput()
	if err != nil {
		if b != nil {
			t.Fatal(string(b))
		}
		t.Fatal(err)
	}
	return string(b)
}
