package cli

import tea "github.com/charmbracelet/bubbletea"

// command represents a single action in the CLI menu.
// Each has a name, description, and a function to execute.
type command struct {
	Name string
	Desc string
	Run  func() tea.Cmd
}

// focusOn enables the focus window (blocks distractions).
var focusOn = command{
	Name: "focus on",
	Desc: "Enable focus window",
	Run: func() tea.Cmd {
		return nil
	},
}

// focusOff disables the focus window (unblocks distractions).
var focusOff = command{
	Name: "focus off",
	Desc: "Disable focus window",
	Run: func() tea.Cmd {
		return nil
	},
}

// configureBlocklist allows editing of the blocked domains list.
var configureBlocklist = command{
	Name: "blocklist",
	Desc: "Configure blocklist",
	Run: func() tea.Cmd{
		return nil
	},
}