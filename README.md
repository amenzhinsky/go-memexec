# go-memexec

Small library that runs embedded binaries.

## Usage

Let's say need to embed ruby interpreter (or any other binary) into your code using tools like `go-bindata` and have ability to execute it:

```go
b, err := Asset("/bin/ruby")
if err != nil {
	panic(err)
}

cmd, err := memexec.Command(b, "-W2", "generate.rb")
if err != nil {
	panic(err)
}
defer cmd.Close() // important!

// cmd embeds *exec.Cmd so you can operate it the same way
o, err := cmd.Output()
if err != nil {
	panic(err)
}

fmt.Println(string(o))
```
