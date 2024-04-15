package main

import (
	"os"

	"github.com/alvaro17f/nixup/cmd"
)

func main() {

	version := os.Getenv("VERSION")

	cmd.SetVersionInfo(version)
	cmd.Execute()
}
