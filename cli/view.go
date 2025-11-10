package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

var appHeaderStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#F9A825")). 
	Bold(true)

var statusStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFF8E1")).
	Background(lipgloss.Color("#33691E")). 
	Padding(0, 1)

var errorAlertStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFE1")). 
	Background(lipgloss.Color("#B71C1C")).
	Padding(1, 0)

var errorInfoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#D7CCC8")). 
	Padding(0, 1)

var listStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4DB6AC")). 
	MarginRight(1)

var listItemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#81D4FA")). 
	MarginRight(1)

var faint = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#A1887F")). 
	Faint(true)


// View model
func (m model) View() string {
	if m.fatalErr != nil {
		return errorAlertStyle.Render("ERROR") + errorInfoStyle.Render(m.fatalErr.Error()) + "\n"
	}

	style := appHeaderStyle.Render("focusgopher") + faint.Render(" - block distractions") + "\n\n"
	style += statusStyle.Render("STATUS") + errorInfoStyle.Render(string(m.status)) + "\n\n"

	if m.state == blacklistView {
		style += "Edit/add domains:\n\n" + m.textarea.View() + "\n\n"
		style += "press Esc to save.\n"
	}

	if m.state == menuView {
		commands := m.getCommandsList()

		l := list.New().Enumerator(func(items list.Items, i int) string {
			if i == m.commandListSelection {
				return "→"
			}
			return " "
		}).
		EnumeratorStyle(listStyle).
		ItemStyle(listItemStyle)
		
		for _, c := range commands {
			l.Item(c.CommandName + faint.Render(" - "+c.Description))
		}
	
		style += l.String() + "\n\n"
	}

	return style
}