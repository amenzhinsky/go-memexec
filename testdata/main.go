package main

import (
	"fmt"
	"os"

	"github.com/amenzhinsky/go-memexec/testdata/perl"
)

func main() {
	exe, err := perl.New()
	if err != nil {
		panic(err)
	}
	c := exe.Command()
	c.Stdin = os.Stdin
	b, err := c.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
