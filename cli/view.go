package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

var appHeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00C2FF")).
	Bold(true)
	
// var statusBarStyle = lipgloss.NewStyle().
// 	Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
// 	Background(lipgloss.AdaptiveColor{Light: "#C8F7FF", Dark: "#1E1E1E"})

// var statusStyle = lipgloss.NewStyle().
// 	Inherit(statusBarStyle).
// 	Foreground(lipgloss.Color("#FFFDF5")).
// 	Background(lipgloss.Color("#FF5F87")).
// 	Padding(1, 0)

// var statusText = lipgloss.NewStyle().Inherit(statusBarStyle).Padding(0, 1)

var errorAlertStyle = lipgloss.NewStyle().
Foreground(lipgloss.Color("#FFFDF5")).
Background(lipgloss.Color("#FF5F87")).
Padding(1, 0)

var errorInfoStyle = lipgloss.NewStyle().
Foreground(lipgloss.Color("250")).
Padding(0, 1)


var listStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("81")).MarginRight(1)
var listItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("117")).MarginRight(1)

var faint = lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Faint(true)

// View model
func (m model) View() string {
	style := appHeaderStyle.Render("focusgopher") + faint.Render(" - block distractions") + "\n\n"

	if !m.initialised {
		style += "..." + "\n"
		return style
	}

	if m.fatalErr != nil {
		style += errorAlertStyle.Render("ERROR") + errorInfoStyle.Render(m.fatalErr.Error()+"\n")
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