package ui

import (
	"github.com/alvaro17f/nixup/internal/utils"
	"github.com/charmbracelet/huh"
)

func Confirm(message string, defaultConfirm ...bool) bool {
	var confirm bool = true

	if len(defaultConfirm) > 0 {
		confirm = defaultConfirm[0]
	}

	err := huh.NewConfirm().
		Title(message).
		Affirmative("Yes").
		Negative("No").
		Value(&confirm).
		Run()
	if err != nil {
		utils.ErrorFormatFatal("Error executing command", err)
	}
	return confirm
}