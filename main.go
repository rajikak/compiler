package compiler

import (
	"fmt"
	"github.com/rajikak/compiler/repl"
	"os"
	"os/user"
)

func main() {
	_, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("monkey v0.0.1")
	repl.Start(os.Stdin, os.Stdout)
}
