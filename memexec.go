package memexec

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// mem is in-memory executable code unit.
type mem struct {
	f *os.File
}

// New creates new memory execution object.
func New(b []byte) (*mem, error) {
	t, err := ioutil.TempFile("", "go-memexec-")
	if err != nil {
		return nil, err
	}

	if _, err = t.Write(b); err != nil {
		t.Close()
		os.Remove(t.Name())
		return nil, err
	}

	// we need only read and execution privileges
	// ioutil.TempFile creates files with 0600 perms
	if err = os.Chmod(t.Name(), 0500); err != nil {
		t.Close()
		os.Remove(t.Name())
		return nil, err
	}

	// binary file has to be closed otherwise
	// we'll get the "text file busy" error
	if err = t.Close(); err != nil {
		os.Remove(t.Name())
		return nil, err
	}

	return &mem{f: t}, nil
}

// Command is an equivalent of exec.Command except name must be omitted.
func (m *mem) Command(arg ...string) *exec.Cmd {
	return exec.Command(m.f.Name(), arg...)
}

// Close removes underlying file.
func (m *mem) Close() error {
	return os.Remove(m.f.Name())
}
