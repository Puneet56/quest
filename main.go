package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// text - 15
// selectSectionBorder 7
// Active section border 2
// Active entry 10
// File info at bottom 12
// Help text

const (
	COLOR_TEXT                  = "15"
	COLOR_SELECTED              = "10"
	COLOR_SECTION_BORDER        = "7"
	COLOR_ACTIVE_SECTION_BORDER = "2"
	COLOR_FILE_INFO             = "12"
	COLOR_HELP_TEXT             = "1"
)

type DemoDir struct{}

func (d *DemoDir) Name() string {
	return "file"
}

func (d *DemoDir) IsDir() bool {
	return false
}

type Pos struct {
	X int
	Y int
}

type Section struct {
	entries []DemoDir
}

type model struct {
	sections []Section
	cursor   Pos
}

func initialModel(dir string) model {
	return model{
		sections: getDummyData(),
		cursor:   Pos{0, 0},
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
			if m.cursor.Y > 0 {
				m.cursor.Y--
			}

		case "down", "j":
			s := m.sections[m.cursor.X]

			if m.cursor.Y < len(s.entries)-1 {
				m.cursor.Y++
			}

		case "left", "h":
			if m.cursor.X > 0 {
				m.cursor.X--
			}

		case "right", "l":
			if m.cursor.X < len(m.sections)-1 {
				m.cursor.X--
			}
		}
	}

	return m, nil
}

var text = lipgloss.NewStyle().
	Foreground(lipgloss.Color(COLOR_TEXT))

var activeText = text.Foreground(lipgloss.Color(COLOR_SELECTED))

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
	h := headingStyle.Render("Quest")

	sections := []string{}

	s := ""
	for i, entry := range m.sections[0].entries {
		if i == m.cursor.Y {
			s += activeText.Render(entry.Name())
		} else {
			s += text.Render(entry.Name())
		}

		if entry.IsDir() {
			s += "/"
		}

		s += "\n"
	}

	sections = append(sections, sectionStyle.Render(s))
	s = ""

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

// func getDirEntries(dir string) []fs.DirEntry {
// 	f, err := os.ReadDir(dir)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return f
// }

func getDummyData() []Section {
	out := make([]Section, 0)

	section1 := Section{
		entries: make([]DemoDir, 10),
	}

	out = append(out, section1)

	section2 := Section{
		entries: make([]DemoDir, 10),
	}

	out = append(out, section2)

	section3 := Section{
		entries: make([]DemoDir, 10),
	}

	out = append(out, section3)

	return out
}
