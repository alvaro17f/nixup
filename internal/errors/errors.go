package errors

import (
	"fmt"
	"os"

	"github.com/alvaro17f/nixup/internal/colors"
)

func ErrorFormat(text string, err error) {
	if err != nil {
		fmt.Printf(
			"%s %s %s\n",
			colors.Color.RedStr("⚠"), colors.Color.YellowStr(text+":"), colors.Color.RedStr(err.Error()),
		)
	} else {
		fmt.Printf(
			"%s %s\n",
			colors.Color.RedStr("⚠"), colors.Color.YellowStr(text),
		)
	}
}

func ErrorFormatFatal(text string, err error) {
	if err != nil {
		fmt.Printf(
			"%s %s %s\n",
			colors.Color.RedStr("⚠"), colors.Color.YellowStr(text+":"), colors.Color.RedStr(err.Error()),
		)
	} else {
		fmt.Printf(
			"%s %s\n",
			colors.Color.RedStr("⚠"), colors.Color.YellowStr(text),
		)
	}
	os.Exit(0)
}
