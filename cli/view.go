package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

var appHeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00C2FF")).
	Bold(true)
	
var statusBarStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
	Background(lipgloss.AdaptiveColor{Light: "#C8F7FF", Dark: "#1E1E1E"})

var statusStyle = lipgloss.NewStyle().
	Inherit(statusBarStyle).
	Foreground(lipgloss.Color("#0077B6")).
	Background(lipgloss.Color("#0077B6")).
	Padding(0, 1)

var statusText = lipgloss.NewStyle().Inherit(statusBarStyle).Padding(0, 1)

var listStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("81")).MarginRight(1)
var listItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("117")).MarginRight(1)

var faint = lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Faint(true)

// View model
func (m model) View() string {
	style := appHeaderStyle.Render("focusgopher") + faint.Render(" - block distractions") + "\n\n"

	if !m.initialised {
		return style
	}

	if m.fatalErr != nil {
		return style + faint.Render("ERROR: "+m.fatalErr.Error()+"\n")
		// bar := lipgloss.JoinHorizontal(lipgloss.Top, statusStyle.Render("ERROR"),
		// 	statusText.Render(m.fatalErr.Error()))
	
		// style += bar + "\n\n"
	}

	l := list.New().Enumerator(func(items list.Items, i int) string {
		if i == m.commandListSelection {
			return "→"
		}
		return " "
	}).
	EnumeratorStyle(listStyle).
	ItemStyle(listItemStyle)
	
	for _, c := range m.commands {
		l.Item(c.CommandName + faint.Render(" - "+c.Description))
	}

	style += l.String() + "\n\n"

	return style
}