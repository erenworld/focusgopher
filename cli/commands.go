package cli

import (
	"strings"

	"focusgopher/hosts"
) 

type command struct {
	CommandName string
	Description string
	RunFunc     func(m model) model
}

var CommandFocusOn =  command{
	CommandName: "focus on",
	Description: "Start focus window.",
	RunFunc: func(m model) model {
		if err := hosts.WriteDomainsToHosts(m.domains, hosts.FocusStatusOn);
		err != nil {
			m.fatalErr = err 
			return m
		}
		m.status = hosts.FocusStatusOn
		return m
	},
}

var CommandFocusOff =  command{
	CommandName: "focus off",
	Description: "Stop focus window.",
	RunFunc: func(m model) model {
		if err := hosts.WriteDomainsToHosts(m.domains, hosts.FocusStatusOff);
		err != nil {
			m.fatalErr = err 
			return m 
		}
		m.status = hosts.FocusStatusOff
		return m
	},
}

var CommandConfigureBlacklist =  command{
	CommandName: "blacklist",
	Description: "Configure blacklist.",
	RunFunc: func(m model) model {
		m.state = blacklistView
		m.textarea.SetValue(strings.Join(m.domains, "\n"))
		m.textarea.Focus()
		m.textarea.CursorEnd()
		return m
	},
}