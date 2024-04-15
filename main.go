package main

import (
	"os"

	"github.com/alvaro17f/nixup/cmd"
)

func main() {

	version := os.Getenv("VERSION")

	if version == "" {
		version = "dev"
	}

	cmd.SetVersionInfo(version)
	cmd.Execute()
}
