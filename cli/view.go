package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

var appHeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF33")).
	Bold(true)
	
var statusBarStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
	Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

var statusStyle = lipgloss.NewStyle().
	Inherit(statusBarStyle).
	Foreground(lipgloss.Color("#FFFDF5")).
	Background(lipgloss.Color("#FF5F87")).
	Padding(0, 1)

var statusText = lipgloss.NewStyle().Inherit(statusBarStyle).Padding(0, 1)

var listStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
var listItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).MarginRight(1)

var faint = lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Faint(true)

func (m model) View() string {
	style := appHeaderStyle.Render("focusgopher") + faint.Render(" - block distractions") + "\n\n"

	if !m.initialised {
		style += "Loading current configuration...\n\n"
	}

	if m.fatalErr != nil {
		bar := lipgloss.JoinHorizontal(lipgloss.Top, statusStyle.Render("ERROR"), statusText.Render(m.fatalErr.Error()))
		
		style += bar + "\n\n"
	}

	if m.initialised && m.fatalErr == nil {
		l := list.New().Enumerator(func(items list.Items, i int) string {
			if i == m.commandListSelection {
				return "→"
			}
			return " "
		}).
		EnumeratorStyle(listStyle).
		ItemStyle(listItemStyle)

		for _, c := range m.commands {
			l.Item(c.CommandName + faint.Render(" - " + c.Description))
		}
		style += l.String() + "\n\n"
	}

	style += "press q to quit.\n"

	return style
}