package cli

import tea "github.com/charmbracelet/bubbletea"

type command struct {
	Name string
	Desc string
	Run  func() tea.Cmd
}

