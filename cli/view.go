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

var listIconStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
var listItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).MarginRight(1)

var faint = lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Faint(true)
