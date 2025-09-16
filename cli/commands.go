package cli

import tea "github.com/charmbracelet/bubbletea"

type command struct {
	Name string
	Desc string
	Run  func() tea.Cmd
}

var focusOn = command{
	Name: "focus on",
	Desc: "Enable focus window",
	Run: func() tea.Cmd {
		return nil
	},
}

var focusOff = command{
	Name: "focus off",
	Desc: "Disable focus window",
	Run: func() tea.Cmd {
		return nil
	},
}

var configureBlocklist = command{
	Name: "blocklist",
	Desc: "Configure blocklist",
	Run: func() tea.Cmd{
		return nil
	},
}