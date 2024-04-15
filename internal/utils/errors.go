package utils

import (
	"fmt"
	"os"
)

func ErrorFormat(text string, err error) {
	fmt.Printf(
		"%s %s %s\n",
		Color.RedStr("⚠"), Color.YellowStr(text+":"), Color.RedStr(err.Error()),
	)
}

func ErrorFormatFatal(text string, err error) {
	fmt.Printf(
		"%s %s %s\n",
		Color.RedStr("⚠"), Color.YellowStr(text+":"), Color.RedStr(err.Error()),
	)
	os.Exit(0)
}
