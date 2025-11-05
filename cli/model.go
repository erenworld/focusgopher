package cli

import (
	"focusgopher/hosts"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	hostsManager		 *hosts.Manager
	commands			 []command
	commandListSelection int
	currentCommand	     *command
	initialised			 bool
	fatalErr			 error
	domains				 []string
	status				 hosts.FocusStatus
}

type initResult struct {
	err error
}

func NewModel() model {
	return model{
		hostsManager: &hosts.Manager{},
		commands: []command{},
	}
}

func (m model) Init() tea.Cmd {
	return m.loadInitialConfig
}

func (m model) loadInitialConfig() tea.Msg {
	initErr := m.hostsManager.Init()	

	return initResult{
		err: initErr,
	}
}

// Update model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case initResult:
		m.initialised = true

		if msg.err != nil {
			m.fatalErr = msg.err
			return m, tea.Quit
		}

		if m.status == hosts.StatusFocusOn {
			m.commands = []command{CommandFocusOff, ConfigureBlacklist}
		} else {
			m.commands = []command{CommandFocusOn, ConfigureBlacklist}
		}
		if len(m.domains) == 0 {
			// todos: return edit
		}

		m.domains = m.hostsManager.Domains
		m.status = m.hostsManager.Status


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