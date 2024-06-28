package main

import (
	"fmt"
	"os"

	"github.com/thebadams/coffee-cli/internal/program"
)

func main() {
	p := program.CreateProgram()
	if _, err := p.Run(); err != nil {
		fmt.Printf("There was an error on start up: %v", err)
		os.Exit(1)
	}
}
