package ui

import (
	"github.com/alvaro17f/nixup/internal/errors"
	"github.com/charmbracelet/huh"
)

func Confirm(message string, defaultConfirm ...bool) bool {
	var confirm bool = true

	if len(defaultConfirm) > 0 {
		confirm = defaultConfirm[0]
	}

	err := NewHuh(huh.NewConfirm().
		Title(message).
		Affirmative("Yes").
		Negative("No").
		Value(&confirm)).Run()
	if err != nil {
		errors.ErrorFormatFatal("Error executing command", err)
	}

	return confirm
}
