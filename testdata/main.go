package main

import (
	"fmt"
	"os"

	"github.com/amenzhinsky/go-memexec/testdata/python3"
)

func main() {
	exe, err := python3.New()
	if err != nil {
		panic(err)
	}
	b, err := exe.Command(os.Args[1:]...).CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
