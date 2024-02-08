package main

import (
	"fmt"
	"os"

	"github.com/bnema/clitt/io/db"
	"github.com/bnema/clitt/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// init the database
	db.InitDB()
	_, err := tea.NewProgram(ui.NewModel()).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}
