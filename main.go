package main

import (
	"time"

	"github.com/alvaro17f/nixup/cmd"
	"github.com/carlmjohnson/versioninfo"
)

func main() {

	cmd.SetVersionInfo(versioninfo.Version, versioninfo.Revision, versioninfo.LastCommit.Format(time.DateOnly))
	cmd.Execute()
}
