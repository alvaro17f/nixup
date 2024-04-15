package utils

import (
	"fmt"
	"strings"
)

func TitleMaker(text string) {
	textLen := len(text)
	border := strings.Repeat("*", textLen+4)
	fmt.Printf(
		"\n%s\n%s %s %s\n%s\n",
		Color.BlueStr(border), Color.BlueStr("*"), Color.RedStr(text), Color.BlueStr("*"), Color.BlueStr(border),
	)
}
