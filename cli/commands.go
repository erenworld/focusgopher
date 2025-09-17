package cli

import tea "github.com/charmbracelet/bubbletea"

// command represents a single action in the CLI menu.
// Each has a name, description, and a function to execute.
type command struct {
	Name string
	Desc string
	Run  func() tea.Cmd
}

// commandFocusOn enables the focus window (blocks distractions).
var commandFocusOn = command{
	Name: "focus on",
	Desc: "Enable focus window",
	Run: func() tea.Cmd {
		return nil
	},
}

// commandFocusOff disables the focus window (unblocks distractions).
var commandFocusOff = command{
	Name: "focus off",
	Desc: "Disable focus window",
	Run: func() tea.Cmd {
		return nil
	},
}

// commandConfigureBlocklist allows editing of the blocked domains list.
var commandConfigureBlocklist = command{
	Name: "blocklist",
	Desc: "Configure blocklist",
	Run: func() tea.Cmd{
		return nil
	},
}