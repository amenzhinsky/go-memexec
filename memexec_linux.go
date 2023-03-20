//go:build linux
// +build linux

package memexec

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func open(b []byte) (*os.File, error) {
	fd, err := unix.MemfdCreate("", unix.MFD_CLOEXEC)
	if err != nil {
		return nil, err
	}
	f := os.NewFile(uintptr(fd), fmt.Sprintf("/proc/%d/fd/%d", os.Getpid(), fd))
	if _, err := f.Write(b); err != nil {
		_ = f.Close()
		return nil, err
	}
	return f, nil
}

func clean(f *os.File) error {
	return f.Close()
}
