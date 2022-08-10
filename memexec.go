package memexec

import (
	"context"
	"os"
	"os/exec"
)

type Option func(e *Exec)

// Exec is an in-memory executable code unit.
type Exec struct {
	f     *os.File
	opts  []func(cmd *exec.Cmd)
	clean func() error
}

// WithPrepare configures cmd with default values such as Env, Dir, etc.
func WithPrepare(fn func(cmd *exec.Cmd)) Option {
	return func(e *Exec) {
		e.opts = append(e.opts, fn)
	}
}

// WithCleanup is executed right after Exec.Close.
func WithCleanup(fn func() error) Option {
	return func(e *Exec) {
		e.clean = fn
	}
}

// New creates new memory execution object that can be
// used for executing commands on a memory based binary.
func New(b []byte, opts ...Option) (*Exec, error) {
	f, err := open(b)
	if err != nil {
		return nil, err
	}
	e := &Exec{f: f}
	for _, opt := range opts {
		opt(e)
	}
	return e, nil
}

// Command is an equivalent of `exec.Command`,
// except that the path to the executable is being omitted.
func (m *Exec) Command(args ...string) *exec.Cmd {
	return m.CommandContext(context.Background(), args...)
}

// CommandContext is an equivalent of `exec.CommandContext`,
// except that the path to the executable is being omitted.
func (m *Exec) CommandContext(ctx context.Context, args ...string) *exec.Cmd {
	exe := exec.CommandContext(ctx, m.f.Name(), args...)
	for _, opt := range m.opts {
		opt(exe)
	}
	return exe
}

// Close closes Exec object.
//
// Any further command will fail, it's client's responsibility
// to control the flow by using synchronization algorithms.
func (m *Exec) Close() error {
	if err := clean(m.f); err != nil {
		if m.clean != nil {
			_ = m.clean()
		}
		return err
	}
	if m.clean == nil {
		return nil
	}
	return m.clean()
}
