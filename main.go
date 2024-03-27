package main

import (
	"fmt"
	"os"

	"github.com/alvaro17f/nixup/flags"
	"github.com/alvaro17f/nixup/ui"
	"github.com/alvaro17f/nixup/utils"
	"github.com/spf13/pflag"
)

var proceed bool

func main() {
	pflag.Parse()

	utils.TitleMaker("Nixup Configuration:")
	utils.Configuration()

	proceed = ui.Confirm("Do you want to update your system?")
	if !proceed {
		os.Exit(0)
	}
	utils.TitleMaker("Git Pull:")
	ui.Spinner("pulling changes...")
	utils.GitPull()

	if *flags.Update {
		utils.TitleMaker("Nix Update:")
		ui.Spinner("updating nixos")
		utils.NixUpdate()
	}

	if utils.GitDiff() {
		utils.TitleMaker("Git Changes:")
		ui.Spinner("checking status...")
		output := utils.GitStatus()

		if output != "" {
			fmt.Println(output)
			proceed = ui.Confirm("Do you want to add these changes to the stage?")
			if proceed {
				ui.Spinner("adding changes...")
				utils.GitAdd()
			}
		}
	}

	utils.TitleMaker("Nix Rebuild:")
	ui.Spinner("nixos rebuild...")
	utils.NixRebuild(utils.GetHostname())

	ui.Spinner("deleting older generations...")
	utils.NixKeep()

	if *flags.Diff {
		utils.TitleMaker("Nix Diff:")
		ui.Spinner("nix diff...")
		utils.NixDiff()
	}
}
