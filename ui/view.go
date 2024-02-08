package ui

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)


func (m Model) View() string {
	s := m.styles
	formView := m.form.View()

	// Status (right side)
	var status string
	if m.form.State == huh.StateCompleted {
		status = fmt.Sprintf("Duration: %s\nTask: %s\nCategory: %s", m.tick, m.taskDescription, m.taskCategory)
		// TAKE THE FULL WIDTH

	} else {
		status = "Timer: " + m.tick.String()
	}
	const statusWidth = 30
	statusMarginLeft := m.width - statusWidth - lipgloss.Width(formView) - s.Status.GetMarginRight()
	statusView := s.Status.Copy().
		Width(statusWidth).
		MarginLeft(statusMarginLeft).
		Render(status)

	errors := m.form.Errors()
	header := m.appBoundaryView("Command Line Time Tracker v0.1")
	if len(errors) > 0 {
		header = m.appErrorBoundaryView(m.errorView())
	}
	body := lipgloss.JoinHorizontal(lipgloss.Top, formView, statusView)
	footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))
	if len(errors) > 0 {
		footer = m.appErrorBoundaryView("")
	}

	return s.Base.Render(header + "\n" + body + "\n\n" + footer)
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(gray),
	)
}

func (m Model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(red),
	)
}

func (m Model) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}
