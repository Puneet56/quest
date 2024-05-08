package main

import "github.com/charmbracelet/lipgloss"

const (
	COLOR_TEXT                  = "15"
	COLOR_SELECTED              = "10"
	COLOR_SECTION_BORDER        = "7"
	COLOR_ACTIVE_SECTION_BORDER = "2"
	COLOR_FILE_INFO             = "12"
	COLOR_HELP_TEXT             = "1"
)

type Styles struct {
	width  int
	height int
}

func NewStyles(width, height int) *Styles {
	return &Styles{width, height}
}

func (s *Styles) Heading() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(COLOR_SELECTED)).
		Bold(true)

}

func (s *Styles) Text() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(COLOR_TEXT))
}

func (s *Styles) ActiveText() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(lipgloss.Color(COLOR_SELECTED)).
		Foreground(lipgloss.Color("#000")).
		Bold(true)
}

func (s *Styles) SectionStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#fff")).
		BorderLeft(true).
		BorderRight(true).
		BorderTop(true).
		BorderBottom(true).
		Width(s.width / 3).
		Height(s.height - 4)
}
