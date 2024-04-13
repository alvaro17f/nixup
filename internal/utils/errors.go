package utils

import (
	"fmt"
	"os"
)

func ErrorFormat(text string, err error) {
	fmt.Printf(
		"%s %s %s\n",
		Red("⚠"), Yellow(text+":"), Red(err),
	)
}

func ErrorFormatFatal(text string, err error) {
	fmt.Printf(
		"%s %s %s\n",
		Red("⚠"), Yellow(text+":"), Red(err),
	)
	os.Exit(0)
}
