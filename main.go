package main

import (
	"fmt"
	"os"

	"focusgopher/cli"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := cli.NewModel()
	program := tea.NewProgram(model)
	if _, err := program.Run(); err != nil {
		fmt.Printf("Unable to launch focusgopher: %v", err)
		os.Exit(1)
	}
}
