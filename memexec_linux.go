//+build linux

package memexec

import (
	"fmt"
	"os"
)

// keep read only fd of the temp file and remove it as soon as possible
// kernel buffers its content in memory until all fds are closed.
func write(t *os.File, b []byte) (*os.File, error) {
	f, err := os.OpenFile(t.Name(), os.O_RDONLY, 0500)
	if err != nil {
		return nil, err
	}

	if err = os.Remove(t.Name()); err != nil {
		f.Close()
		return nil, err
	}

	// write content
	if _, err = t.Write(b); err != nil {
		f.Close()
		return nil, err
	}

	// binary file has to be closed otherwise
	// we'll get the "text file busy" error
	if err = t.Close(); err != nil {
		f.Close()
		return nil, err
	}
	return f, nil
}

func path(m *mem) string {
	return fmt.Sprintf("/proc/self/fd/%d", int(m.f.Fd()))
}

func close(m *mem) error {
	return m.f.Close()
}
