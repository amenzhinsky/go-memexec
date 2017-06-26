# go-memexec

Small library that executes code from memory.

## Usage

Let's say need to embed ruby interpreter (or any other binary) into your code using tools like `go-bindata` and have ability to execute it:

```go
func RubyExec(file string) (b []byte, err error) {
	b, err := Asset("/bin/ruby")
	if err != nil {
		return
	}

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

	return m.Command(file).Output()
}
```
