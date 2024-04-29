package main

import (
	"github.com/alvaro17f/nixup/cmd"
)

var version = "dev"

func main() {
	cmd.SetVersionInfo(version)
	cmd.Execute()
}
