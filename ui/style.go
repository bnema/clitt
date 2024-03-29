package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const maxWidth = 80

var (
	red     = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo  = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	blueish = lipgloss.AdaptiveColor{Light: "#6E7FF3", Dark: "#6E7FF3"}
	green   = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
	yellow  = lipgloss.AdaptiveColor{Light: "#FFD700", Dark: "#FFD700"}
	gray    = lipgloss.AdaptiveColor{Light: "#A9A9A9", Dark: "#A9A9A9"}
)

type Styles struct {
	Base, HeaderText, Status, StatusHeader, Highlight, ErrorHeaderText, Help lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().Foreground(blueish).Bold(true).Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(gray).PaddingLeft(1).MarginTop(1)
	s.StatusHeader = lg.NewStyle().Foreground(green).Bold(true)
	s.Highlight = lg.NewStyle().Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.Copy().Foreground(red)
	s.Help = lg.NewStyle().Foreground(lipgloss.Color("240"))
	return &s
}
