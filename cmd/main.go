package main

import (
	"fmt"
	"github.com/bnema/clitt/ui"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	_, err := tea.NewProgram(ui.NewModel()).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

}
