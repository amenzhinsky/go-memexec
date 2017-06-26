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
	
	cmd, err := memexec.Command(b, file)
	if err != nil {
		return
	}
	defer func() {
		cerr := cmd.Close() 
		if err == nil {
			err = cerr
		}
	}() // important!
	
	// cmd embeds *exec.Cmd so you can operate it the same way
	return cmd.Output()
}
```
