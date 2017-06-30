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

// New creates new memory execution object that can be
// used for executing commands on a memory based binary.
func New(b []byte) (*mem, error) {
	f, err := ioutil.TempFile("", "go-memexec-")
	if err != nil {
		return nil, err
	}

	// close and remove the temporary file
	// if something goes wrong, we need a reference
	// because f can be overwritten
	defer func(f *os.File) {
		if f != nil && err != nil {
			f.Close()
			os.Remove(f.Name())
		}
	}(f)

	// we need only read and execution privileges
	// ioutil.TempFile creates files with 0600 perms
	if err = os.Chmod(f.Name(), 0500); err != nil {
		return nil, err
	}

	f, err = write(f, b)
	if err != nil {
		return nil, err
	}

	return &mem{f: f}, nil
}

// Command is an equivalent of exec.Command except that path
// to binary must be omitted.
func (m *mem) Command(arg ...string) *exec.Cmd {
	return exec.Command(path(m), arg...)
}

// Close closes mem object.
// Any further command will fail, it's client's responsibility
// to control the flow by using synchronization algorithms.
func (m *mem) Close() error {
	return close(m)
}
