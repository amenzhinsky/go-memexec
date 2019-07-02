# go-memexec

Small library that executes code from the memory.

## Usage

The following example executes ruby executable embedded into the binary: 

```go
// Asset function provided by a library such as go-bindata
b, err := Asset("/bin/ruby")
if err != nil {
	return err
}

exe, err := memexec.New(b)
if err != nil {
	return err
}
defer exe.Close()

cmd := exe.Command(argv...)
cmd.Output() // cmd is a `*exec.Cmd` from the standard library
```
