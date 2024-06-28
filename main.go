package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/thebadams/coffee-cli/internal/program"
)

func main() {
	p := tea.NewProgram(program.DefaultProgram())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been an error: %v", err)
		os.Exit(1)
	}
}
