package cli

import tea "github.com/charmbracelet/bubbletea"

type command struct {
	CommandName string
	Description string
	RunFunc     func() tea.Cmd
}

var CommandFocusOn =  command{
	CommandName: "focus on",
	Description: "Start focus window.",
	RunFunc: func() tea.Cmd { return nil },
}

var CommandFocusOff =  command{
	CommandName: "focus off",
	Description: "Stop focus window.",
	RunFunc: func() tea.Cmd { return nil },
}

var ConfigureBlacklist =  command{
	CommandName: "blacklist",
	Description: "Configure blacklist.",
	RunFunc: func() tea.Cmd { return nil },
}