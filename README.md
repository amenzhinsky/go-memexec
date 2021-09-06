# go-memexec

Small library that executes code from the memory.

## Usage

The following example executes executable embedded into the binary: 

```go
import _ "embed"

// go:embed path-to-binary
var mybinary []byte

exe, err := memexec.New(mybinary)
if err != nil {
	return err
}
defer exe.Close()

cmd := exe.Command(argv...)
cmd.Output() // cmd is a `*exec.Cmd` from the standard library
```
