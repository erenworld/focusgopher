package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"focusgopher/hosts"
)

type model struct {
	hostsManager		 *hosts.Manager
	commands			 []command
	commandListSelection int
	currentCommand	     *command
	initialised			 bool
	fatalErr			 error
	domains				 []string
}

func NewModel() model {
	return model{
		hostsManager: &hosts.Manager,
		commands: []command{CommandFocusOn, CommandFocusOff, ConfigureBlacklist},
	}
}

