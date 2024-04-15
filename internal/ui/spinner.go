package ui

import (
	"fmt"

	"github.com/alvaro17f/nixup/internal/utils"
	"github.com/charmbracelet/huh/spinner"
)

func Spinner(msg string) {
	spinnerColor := utils.Color.Style(utils.Color.Red)
	titleColor := utils.Color.Style(utils.Color.White)
	err := spinner.New().Type(spinner.MiniDot).Style(spinnerColor).TitleStyle(titleColor).Title(" " + msg).Run()
	if err != nil {
		fmt.Println("Error executing command")
	}
}
