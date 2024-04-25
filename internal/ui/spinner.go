package ui

import (
	"fmt"

	"github.com/alvaro17f/nixup/internal/colors"
	"github.com/charmbracelet/huh/spinner"
)

func Spinner(msg string) {
	spinnerColor := colors.Color.Style(colors.Color.Red)
	titleColor := colors.Color.Style(colors.Color.White)

	err := spinner.New().Type(spinner.MiniDot).Style(spinnerColor).TitleStyle(titleColor).Title(" " + msg).Run()
	if err != nil {
		fmt.Println("Error executing command")
	}
}
