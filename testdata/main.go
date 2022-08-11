package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/amenzhinsky/go-memexec/testdata/python3"
)

func main() {
	exe, err := python3.New()
	if err != nil {
		panic(err)
	}
	c := exe.Command(os.Args[1:]...)
	c.Stdin = os.Stdin
	b, err := c.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Print(strings.TrimSpace(string(b)))
}
