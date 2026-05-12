package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Dandarprox/tuiference/internal/tui"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "update":
			if err := update(); err != nil {
				fmt.Fprintf(os.Stderr, "tuiference update: %v\n", err)
				os.Exit(1)
			}
			return
		case "help", "--help", "-h":
			printHelp()
			return
		}
	}

	program := tea.NewProgram(tui.New(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "tuiference: %v\n", err)
		os.Exit(1)
	}
}

func update() error {
	if _, err := exec.LookPath("go"); err != nil {
		return fmt.Errorf("go is required for self-update; install manually from https://github.com/Dandarprox/tuiference/releases/latest or run the install script from the README")
	}

	cmd := exec.Command("go", "install", "github.com/Dandarprox/tuiference@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Updating tuiference with go install github.com/Dandarprox/tuiference@latest")
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("tuiference is up to date.")
	return nil
}

func printHelp() {
	fmt.Println(`tuiference - minimal terminal WordReference lookup

Usage:
  tuiference          Start the TUI
  tuiference update   Update via go install
  tuiference help     Show this help`)
}
