package cli

import tea "github.com/charmbracelet/bubbletea"

type command struct {
	Name string
	Desc string
	Run  func() tea.Cmd
}

var commandFocusOn = command{
	Name: "focus on",
	Desc: "Enable focus window",
	Run: func() tea.Cmd {
		return nil
	},
}

var commandFocusOff = command{
	Name: "focus off",
	Desc: "Disable focus window",
	Run: func() tea.Cmd {
		return nil
	},
}


