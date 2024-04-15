package ui

import "github.com/charmbracelet/huh"

func NewHuh(module huh.Field, theme ...*huh.Theme) *huh.Form {
	selectedTheme := huh.ThemeCatppuccin()

	if len(theme) > 0 {
		selectedTheme = theme[0]
	}

	return huh.NewForm(huh.NewGroup(module)).WithTheme(selectedTheme)
}
