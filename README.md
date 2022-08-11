# go-memexec

Small library that executes binaries from the memory.

## Usage

## Static Binary

Running static binaries is quite simple, it's only needed to embed it into the app and pass its content to `memexec.New`:

```go
import (
	_ "embed"
	
	"github.com/amenzhinsky/go-memexec"
)

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

## Dynamic Binary (Linux only)

With dynamic linked binaries things get more complicated, it's needed to embed all dependencies along with the executable.

At the runtime deps are copied to a temp dir and executable receives the corresponding `LD_LIBRARY_PATH` that forces the dynamic linker to use the copied libraries.

The dynamic linker must be the same on both building and running machines since it's not included in the resulting binary (there's no interoperability between musl and GNU systems).

The following script helps generating packages, say `python3`:

```sh
go install github.com/amenzhinsky/go-memexec/cmd/memexec-gen@latest
PATH=$(go env GOPATH)/bin:$PATH memexec-gen /usr/bin/python3
```

It produces `python3` directory with binaries are to embed and `gen.go` file, that is a go package:

```go
import "mypackagename/python3"

exe, err := python3.New()
if err != nil {
	return err
}
defer exe.Close()

b, err := exe.Command("-c", "print('Hello World')").CombinedOutput()
if err != nil {
	return err
}
```
