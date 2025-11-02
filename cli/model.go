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

