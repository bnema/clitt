package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	lg              *lipgloss.Renderer
	styles          *Styles
	form            *huh.Form
	start           time.Time
	tick            time.Duration
	width           int
	taskCategory    string
	taskDescription string
	taskCategories  []string
}

type tickMsg struct{}

func tick() tea.Msg {
	return tickMsg{}
}

func NewModel() Model {
	m := Model{width: maxWidth, start: time.Now()}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)
	m.taskCategories = []string{"Internal Call", "Training", "Technical Service", "Phone Support", "Email Processing", "Development"}

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Key("task").Title("What are you working on?"),
			huh.NewSelect[string]().Key("category").Options(huh.NewOptions(m.taskCategories...)...).Title("Choose a category"),
			huh.NewConfirm().Key("done").Title("All done?").Validate(func(v bool) error {
				if !v {
					return fmt.Errorf("Welp, finish up then")
				}
				return nil
			}).Affirmative("Yep").Negative("Wait, no"),
		),
	).WithWidth(45).WithShowHelp(false).WithShowErrors(false)
	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tick()
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		}
	case tickMsg:
		m.tick = time.Since(m.start).Round(time.Second)
		cmds = append(cmds, tea.Tick(time.Second, func(t time.Time) tea.Msg {
			return tick()
		}))
	}

	// Process the form
	form, cmd := m.form.Update(msg)
	m.form = form.(*huh.Form)
	cmds = append(cmds, cmd)

	if m.form.State == huh.StateCompleted {
		m.taskDescription = m.form.GetString("task")
		m.taskCategory = m.form.GetString("category")
		// Equivalent action for task completion
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}
