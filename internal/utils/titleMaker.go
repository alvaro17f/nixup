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
		Blue(border), Blue("*"), Red(text), Blue("*"), Blue(border),
	)
}
