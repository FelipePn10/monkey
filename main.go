package main

import (
	"fmt"
	"interpretador/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programmin language!\n",
		user.Username)
	fmt.Printf("Feel free to typer")
	repl.Start(os.Stdin, os.Stdout)
}
