package main

import (
	"fmt"
	"io/fs"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	entries []fs.DirEntry
	cursor  int
}

func initialModel(dir string) model {
	return model{
		entries: getDirEntries(dir),
		cursor:  0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.entries)-1 {
				m.cursor++
			}

		}
	}

	return m, nil
}

var style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF"))

var selectedStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF8800"))

var sectionStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#fff")).
	Width(25).
	Height(10)

var headingStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF8800")).
	Width(25)

func (m model) View() string {
	var h = headingStyle.Render("Quest") + "\n\n"

	var sections = []string{}

	var s = ""
	for i, entry := range m.entries {
		if i == m.cursor {
			s += selectedStyle.Render(entry.Name())
		} else {
			s += style.Render(entry.Name())
		}

		if entry.IsDir() {
			s += "/"
		}

		s += "\n"
	}

	sections = append(sections, sectionStyle.Render(s))
	s = ""

	if m.entries[m.cursor].IsDir() {
		entries := getDirEntries(m.entries[m.cursor].Name())

		for _, entry := range entries {
			if entry.IsDir() {
				s += style.Render(entry.Name()) + "/" + "\n"
			} else {
				s += style.Render(entry.Name()) + "\n"
			}
		}

		sections = append(sections, sectionStyle.Render(s))
		s = ""
	}

	return lipgloss.JoinVertical(lipgloss.Top, h, lipgloss.JoinHorizontal(lipgloss.Left, sections...))
}

func main() {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(initialModel(dir), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func getDirEntries(dir string) []fs.DirEntry {
	f, err := os.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	return f
}
