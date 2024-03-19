package ui

import (
	"github.com/charmbracelet/huh/spinner"
)

func Spinner(msg string) {
	_ = spinner.New().Title(msg).Run()
	// fmt.Println("Order up!")
}
