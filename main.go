package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/PyMarcus/lioninterpreter/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome, %s, to Lion Programming Language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
