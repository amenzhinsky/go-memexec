package memexec

import (
	"context"
	"os"
	"os/exec"
)

// Exec is an in-memory executable code unit.
type Exec struct {
	f *os.File
}

// New creates new memory execution object that can be
// used for executing commands on a memory based binary.
func New(b []byte) (*Exec, error) {
	f, err := open(b)
	if err != nil {
		return nil, err
	}
	return &Exec{f: f}, nil
}

// Command is an equivalent of `exec.Command`,
// except that the path to the executable is being omitted.
func (m *Exec) Command(arg ...string) *exec.Cmd {
	return exec.Command(m.f.Name(), arg...)
}

// CommandContext is an equivalent of `exec.CommandContext`,
// except that the path to the executable is being omitted.
func (m *Exec) CommandContext(ctx context.Context, arg ...string) *exec.Cmd {
	return exec.CommandContext(ctx, m.f.Name(), arg...)
}

// Close closes Exec object.
//
// Any further command will fail, it's client's responsibility
// to control the flow by using synchronization algorithms.
func (m *Exec) Close() error {
	return clean(m.f)
}
