package main

import (
	"github.com/alvaro17f/nixup/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(version)
	cmd.Execute()
}
