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

type initResult struct {
	err error
}

func NewModel() model {
	return model{
		hostsManager: &hosts.Manager{},
		commands: []command{CommandFocusOn, CommandFocusOff, ConfigureBlacklist},
	}
}

func (m model) Init() tea.Cmd {
	return m.loadInitialConfig
}

func (m model) loadInitialConfig() tea.Msg {
	if err := m.hostsManager.Init(); err != nil {
		return initResult{
			err: err,
		}
	}

	return initResult{}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case initResult:
		m.initialised = true
		if msg.err != nil {
			m.fatalErr = msg.err
			return m, tea.Quit
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.commandListSelection > 0 {
				m.commandListSelection--
			}
		case "down":
			if m.commandListSelection < len(m.commands) - 1 {
				m.commandListSelection++
			}
		case "enter", " ":
			m.currentCommand = &m.commands[m.commandListSelection]
			return m, m.currentCommand.RunFunc()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}