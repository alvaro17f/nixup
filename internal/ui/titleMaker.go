package ui

import (
	"fmt"
	"strings"

	"github.com/alvaro17f/nixup/internal/colors"
)

func TitleMaker(text string) {
	textLen := len(text)
	border := strings.Repeat("*", textLen+4)
	fmt.Printf(
		"\n%s\n%s %s %s\n%s\n",
		colors.Color.BlueStr(border), colors.Color.BlueStr("*"), colors.Color.RedStr(text), colors.Color.BlueStr("*"), colors.Color.BlueStr(border),
	)
}
