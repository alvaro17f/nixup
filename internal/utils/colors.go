package utils

import (
	"github.com/charmbracelet/lipgloss"
)

type colors struct {
	Blue   string
	Red    string
	Yellow string
	Green  string
	White  string
}

var Color = colors{
	Blue:   "#89b4fa",
	Red:    "#f38ba8",
	Yellow: "#f9e2af",
	Green:  "#a6e3a1",
	White:  "#cdd6f4",
}

func (c *colors) Style(fg string) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(fg))
}

func (c *colors) BlueStr(text string) string {
	return c.Style(c.Blue).Render(text)
}

func (c *colors) YellowStr(text string) string {
	return c.Style(c.Yellow).Render(text)
}

func (c *colors) GreenStr(text string) string {
	return c.Style(c.Green).Render(text)
}

func (c *colors) RedStr(text string) string {
	return c.Style(c.Red).Render(text)
}
