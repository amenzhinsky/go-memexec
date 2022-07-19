package memexec

import (
	"context"
	"io/ioutil"
	"os/exec"
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
	o, err := c.Output()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(o), "test") {
		t.Errorf("command output = %q doesn't contain %q", string(o), "test")
	}
}

func TestCommandContext(t *testing.T) {
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
	go c.Run()
	c.Wait()
	stop := time.Now()
	delta := stop.Sub(start)

	if delta > time.Millisecond*600 || delta < time.Millisecond*500 {
		t.Errorf("unexpected command execution time, delta=%s", delta)
	}
}

func BenchmarkCommand(b *testing.B) {
	exe := newEchoExec(b)
	defer exe.Close()
	for i := 0; i < b.N; i++ {
		cmd := exe.Command("test")
		if _, err := cmd.Output(); err != nil {
			b.Fatal(err)
		}
	}
}

func newExec(name string, t testing.TB) *Exec {
	path, err := exec.LookPath(name)
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadFile(path)
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
	return newExec("echo", t)
}

func newSleepExec(t testing.TB) *Exec {
	return newExec("sleep", t)
}
