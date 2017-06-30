//+build !linux

package memexec

import "os"

func write(t *os.File, b []byte) (*os.File, error) {
	if _, err := t.Write(b); err != nil {
		return nil, err
	}

	if err := t.Close(); err != nil {
		return nil, err
	}
	return t, nil
}

func path(m *mem) string {
	return m.f.Name()
}

func close(m *mem) error {
	return os.Remove(m.f.Name())
}
