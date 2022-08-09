//go:build !linux
// +build !linux

package memexec

import (
	"os"
	"runtime"
)

func open(b []byte) (*os.File, error) {
	pattern := "go-memexec-"
	if runtime.GOOS == "windows" {
		pattern = "go-memexec-*.exe"
	}
	f, err := os.CreateTemp("", pattern)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = clean(f)
		}
	}()
	if err = os.Chmod(f.Name(), 0o500); err != nil {
		return nil, err
	}
	if _, err = f.Write(b); err != nil {
		return nil, err
	}
	if err = f.Close(); err != nil {
		return nil, err
	}
	return f, nil
}

func clean(f *os.File) error {
	return os.Remove(f.Name())
}
