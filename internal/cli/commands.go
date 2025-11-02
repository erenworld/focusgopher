package cli

import tea "github.com/charmbracelet/bubbletea"

type command struct {
	CommandName string
	Description string
	RunFunc     func() tea.Cmd
}
