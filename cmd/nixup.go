package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alvaro17f/nixup/internal/features"
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

	ui.TitleMaker("Nixup Configuration:")
	features.Configuration(repo, hostname, update, keep, diff)

	proceed = ui.Confirm(fmt.Sprintf("Hi %s, Do you want to update your system?", utils.GetUser().Name))
	if !proceed {
		os.Exit(0)
	}
	ui.TitleMaker("Git Pull:")
	ui.Spinner("pulling changes...")
	features.GitPull(cmd.Flag("repo").Value.String())

	if cmd.Flag("update").Changed {
		ui.TitleMaker("Nix Update:")
		ui.Spinner("updating nixos")
		features.NixUpdate(repo)
	}

	if features.GitDiff(repo) {
		ui.TitleMaker("Git Changes:")
		ui.Spinner("checking status...")
		output := features.GitStatus(repo)

		if output != "" {
			fmt.Println(output)
			proceed = ui.Confirm("Do you want to add these changes to the stage?")
			if proceed {
				ui.Spinner("adding changes...")
				features.GitAdd(repo)
			}
		}
	}

	ui.TitleMaker("Nix Rebuild:")
	ui.Spinner("nixos rebuild...")
	features.NixRebuild(repo, hostname)

	ui.Spinner("deleting older generations...")
	keepInt, _ := strconv.Atoi(keep)
	features.NixKeep(keepInt)

	if cmd.Flag("diff").Changed {
		ui.TitleMaker("Nix Diff:")
		ui.Spinner("nix diff...")
		features.NixDiff()
	}
}
