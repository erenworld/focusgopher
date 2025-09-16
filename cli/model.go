package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/erenworld/focusgopher/hosts"
)

type model struct {
	hostsManager			*hosts.Manager
	commands				[]command
	commandsListSelection	int
	currentCommand			*command
	initialised				bool
	fatalErr				error
	domains					[]string
}

