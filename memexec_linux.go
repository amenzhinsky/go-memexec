//+build linux

package memexec

import (
	"fmt"
	"os"
)

// on linux we can keep a read only fd of the temp file and remove
// kernel buffers its content in memory until all fds are closed.
func write(t *os.File, b []byte) (f *os.File, err error) {
	f, err = os.OpenFile(t.Name(), os.O_RDONLY, 0500)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			f.Close()
		}
	}()

	// check if /proc is mounted
	p := fmt.Sprintf("/proc/self/fd/%d", int(f.Fd()))
	if _, err := os.Lstat(p); err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("%s dosn't exist, probably /proc is not mounted", p)
		}
		return nil, err
	}

	// write code, remove and close the original temporary file
	// otherwise  we'll get the "text file busy" error
	if _, err = t.Write(b); err != nil {
		return nil, err
	}

	if err = os.Remove(t.Name()); err != nil {
		return nil, err
	}
	return f, t.Close()
}

func path(m *mem) string {
	return fmt.Sprintf("/proc/self/fd/%d", int(m.f.Fd()))
}

func close(m *mem) error {
	return m.f.Close()
}
