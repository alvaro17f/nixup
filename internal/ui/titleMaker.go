package ui

import (
	"fmt"
	"strings"

	"github.com/alvaro17f/nixup/internal/colors"
)

func TitleMaker(text string) {
	const borderPadding = 4

	textLen := len(text)
	border := strings.Repeat("*", textLen+borderPadding)
	fmt.Printf(
		"\n%s\n%s %s %s\n%s\n",
		colors.Color.BlueStr(border),
		colors.Color.BlueStr("*"),
		colors.Color.RedStr(text),
		colors.Color.BlueStr("*"),
		colors.Color.BlueStr(border),
	)
}
