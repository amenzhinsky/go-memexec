# go-memexec [![CircleCI](https://circleci.com/gh/amenzhinsky/go-memexec.svg?style=svg)](https://circleci.com/gh/amenzhinsky/go-memexec)

Small library that executes code from memory.

## Usage

Let's say you have ruby binary embedded into you code by [go-bindata](https://github.com/jteeuwen/go-bindata) and you need to have ability to use it:

```go
func RubyExec(argv ...string) (b []byte, err error) {
	// Asset function provided by go-bindata
	b, err := Asset("/bin/ruby")
	if err != nil {
		return
	}

	// m can be cached to avoid extra copying
	// when it's needed exec the same code multiple times
	m, err := memexec.New(b)
	if err != nil {
		return
	}
	defer func() {
		cerr := m.Close() 
		if err == nil {
			err = cerr
		}
	}()

	return m.Command(argv...).Output()
}
```
