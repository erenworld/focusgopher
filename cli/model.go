package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/erenworld/focusgopher/hosts"
)

// model holds all application state required by Bubble Tea.
type model struct {
	hostsManager            *hosts.Manager	
	commands                []command		
	commandsListSelection   int				
	currentCommand          *command		
	initialised             bool
	fatalErr                error
	domains                 []string	
}

// initResult carries the outcome of loading initial configuration.
type initResult struct {
	err error
}

// NewModel constructs a model with a hosts manager and default commands.
func NewModel() model {
	return model{
		hostsManager: &hosts.Manager{},
		commands:     []command{focusOn, focusOff, configureBlocklist},
	}
}

// Init is called once at program start. It triggers background config loading.
func (m model) Init() tea.Cmd {
	return m.loadInitialConfig
}

// loadInitialConfig initializes the hosts manager and reports success or error.
func (m model) loadInitialConfig() tea.Msg {
	initErr := m.hostsManager.Init()
	return initResult{
		err: initErr,
	}
}

// Update applies messages (events) to the model and returns the next command.
// initResult: marks initialization complete or records a fatal error
// tea.KeyMsg: handles keypresses (navigation, selection, quit)
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
		case "up", "k":
			if m.commandsListSelection > 0 {
				m.commandsListSelection--
			}
		case "down", "j":
			if m.commandsListSelection < len(m.commands)-1 {
				m.commandsListSelection++
			}
		case "enter", " ":
			m.currentCommand = &m.commands[m.commandsListSelection]
			return m, m.currentCommand.Run()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}
