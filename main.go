package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	Width    int
	Height   int
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
				m.cursor.X++
			}
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}

	return m, nil
}

func (m model) View() string {
	st := NewStyles(m.Width, m.Height)

	h := st.Heading().Render("Quest")

	sections := []string{}

	for x, section := range m.sections {

		s := ""
		for y, entry := range section.entries {
			if y == m.cursor.Y && x == m.cursor.X {
				s += st.ActiveText().Render(entry.Name())
			} else {
				s += st.Text().Render(entry.Name())
			}

			if entry.IsDir() {
				s += "/"
			}

			s += "\n"

		}

		sectionStyle := st.SectionStyle()

		if x > 0 {
			sectionStyle = sectionStyle.BorderLeft(false)
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

	m := initialModel(dir)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", &m)

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
