//+build !linux

package memexec

import "os"

type executor struct {
	f *os.File
}

func (e *executor) prepare(t *os.File) error {
	e.f = t
	return nil
}

func (e *executor) path() string {
	return e.f.Name()
}

func (e *executor) close() error {
	return os.Remove(e.f.Name())
}
