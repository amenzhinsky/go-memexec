package memexec

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// Cmd is a program prepared to run.
// Equivalent of *exec.Cmd except the need of closing it.
type Cmd struct {
	*exec.Cmd
}

// Close removes temporary file.
func (cmd *Cmd) Close() error {
	return os.Remove(cmd.Path)
}

// TODO: consider using syscall.Mmap or fexecve on unix
// Command returns a Cmd struct prepared to run, like exec.Command.
// It's clients responsibility to close it.
func Command(code []byte, argv ...string) (*Cmd, error) {
	f, err := ioutil.TempFile("", "go-memexec-")
	if err != nil {
		return nil, err
	}

	// we need execution privileges
	if err = os.Chmod(f.Name(), 0700); err != nil {
		f.Close()
		os.Remove(f.Name())
		return nil, err
	}

	if _, err = f.Write(code); err != nil {
		f.Close()
		os.Remove(f.Name())
		return nil, err
	}

	// binary file has to be closed otherwise
	// we'll get the "text file busy" error
	if err = f.Close(); err != nil {
		os.Remove(f.Name())
		return nil, err
	}

	return &Cmd{exec.Command(f.Name(), argv...)}, nil
}
