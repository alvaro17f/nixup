package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alvaro17f/nixup/internal/ui"
	"github.com/alvaro17f/nixup/internal/utils"
	"github.com/spf13/cobra"
)

var proceed bool

func Nixup(cmd *cobra.Command, args []string) {
	var (
		repo     = cmd.Flag(RepoFlag).Value.String()
		hostname = cmd.Flag(HostnameFlag).Value.String()
		diff     = cmd.Flag(DiffFlag).Changed
		keep     = cmd.Flag(KeepFlag).Value.String()
		update   = cmd.Flag(UpdateFlag).Changed
	)

	utils.TitleMaker("Nixup Configuration:")
	utils.Configuration(repo, hostname, update, keep, diff)

	proceed = ui.Confirm(fmt.Sprintf("Hi %s, Do you want to update your system?", utils.GetUser().Name))
	if !proceed {
		os.Exit(0)
	}
	utils.TitleMaker("Git Pull:")
	ui.Spinner("pulling changes...")
	utils.GitPull(cmd.Flag("repo").Value.String())

	if cmd.Flag("update").Changed {
		utils.TitleMaker("Nix Update:")
		ui.Spinner("updating nixos")
		utils.NixUpdate(repo)
	}

	if utils.GitDiff(repo) {
		utils.TitleMaker("Git Changes:")
		ui.Spinner("checking status...")
		output := utils.GitStatus(repo)

		if output != "" {
			fmt.Println(output)
			proceed = ui.Confirm("Do you want to add these changes to the stage?")
			if proceed {
				ui.Spinner("adding changes...")
				utils.GitAdd(repo)
			}
		}
	}

	utils.TitleMaker("Nix Rebuild:")
	ui.Spinner("nixos rebuild...")
	utils.NixRebuild(repo, hostname)

	ui.Spinner("deleting older generations...")
	keepInt, _ := strconv.Atoi(keep)
	utils.NixKeep(keepInt)

	if cmd.Flag("diff").Changed {
		utils.TitleMaker("Nix Diff:")
		ui.Spinner("nix diff...")
		utils.NixDiff()
	}
}
