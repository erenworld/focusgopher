package cli

import (
	"strings"

	"focusgopher/hosts"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState uint

const (
	menuView sessionState = iota
	blacklistView
)

type model struct {
	commands			 []command
	commandListSelection int
	fatalErr			 error
	domains				 []string
	status				 hosts.FocusStatus
	textarea			 textarea.Model
	state				 sessionState
}

type initResult struct {
	err error
}

func NewModel() model {
	domains, status, err := hosts.ExtractDomainsFromHosts()

	state := menuView
	text := textarea.New()
	if len(domains) == 0 {
		state = blacklistView
		text.SetValue(strings.Join(hosts.DefaultDomains, "\n"))
		text.Focus()
		text.CursorEnd()
	} else {
		text.Blur()
	}

	return model{
		textarea: 	text,
		domains: 	domains,
		state:		state,
		status:		status,
		fatalErr:   err,
	}
}

func (m model) Init() tea.Cmd {
	return m.loadInitialConfig
}

func (m model) loadInitialConfig() tea.Msg {
	if m.fatalErr != nil {
		return tea.Quit()
	}

	return nil
}

func (m *model) getCommandsList() []command {
	if m.status == hosts.FocusStatusOn {
		return []command{CommandFocusOff, CommandConfigureBlacklist}
	}

	return []command{CommandFocusOn, CommandConfigureBlacklist}
}

// Update model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	
	switch msg := msg.(type) {
		
	case tea.KeyMsg:
		commands := m.getCommandsList()
		switch msg.String() {
		case "up", "k":
			if m.state == menuView && m.commandListSelection > 0 {
				m.commandListSelection--
			}
		case "down", "j":
			if m.state == menuView && m.commandListSelection < len(commands)-1 {
				m.commandListSelection++
			}
		case "enter", " ":
			if m.state == menuView {
				m = commands[m.commandListSelection].RunFunc(m)
				if m.fatalErr != nil {
					return m, tea.Quit
				}
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.state == blacklistView {
				domains := strings.Split(m.textarea.Value() + "\n")
				domains = hosts.CleanDomainsList(domains)

				if err := hosts.WriteDomainsToHostsFile(domains, m.status); err != nil {
					m.fatalErr = err 
					return m, tea.Quit
				}

				m.domains = domains
				m.state = menuView
				m.textarea.Blur()
			}
		}
	}
	
	return m, tea.Batch(cmds...)
}