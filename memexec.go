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
	f, err := ioutil.TempFile("", "go-memexec-")
	if err != nil {
		return nil, err
	}

	// we need only read and execution privileges
	// ioutil.TempFile creates files with 0600 perms
	if err = os.Chmod(f.Name(), 0500); err != nil {
		f.Close()
		os.Remove(f.Name())
		return nil, err
	}

	f, err = write(f, b)
	if err != nil {
		f.Close()
		os.Remove(f.Name())
		return nil, err
	}

	return &mem{f: f}, nil
}

// Command is an equivalent of exec.Command except name must be omitted.
func (m *mem) Command(arg ...string) *exec.Cmd {
	return exec.Command(path(m), arg...)
}

// Close removes underlying file.
func (m *mem) Close() error {
	return close(m)
}
