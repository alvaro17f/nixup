package features

import (
	"fmt"
	"strconv"

	"github.com/alvaro17f/nixup/internal/colors"
	"github.com/alvaro17f/nixup/internal/errors"
	"github.com/alvaro17f/nixup/internal/utils"
)

const (
	gitPullCmd    = "git -C %s pull"
	gitDiffCmd    = "git -C %s diff --exit-code"
	gitStatusCmd  = "git -C %s status --porcelain"
	gitAddCmd     = "git -C %s add ."
	nixUpdateCmd  = "cd %s && nix flake update"
	nixRebuildCmd = "sudo nixos-rebuild switch --flake %s#%s --show-trace"
	nixKeepCmd    = "sudo nix-env --profile /nix/var/nix/profiles/system --delete-generations +%d"
	nixDiffCmd    = "nix profile diff-closures --profile /nix/var/nix/profiles/system | tac | awk '/Version/{print; exit} 1' | tac"
)

func Configuration(repo string, hostname string, update bool, keep string, diff bool) {

	fmt.Println(colors.Color.BlueStr("• Repo:"), colors.Color.YellowStr(repo))
	fmt.Println(colors.Color.BlueStr("• Hostname:"), colors.Color.YellowStr(hostname))
	fmt.Println(colors.Color.BlueStr("• Update:"), colors.Color.YellowStr(strconv.FormatBool(update)))
	fmt.Println(colors.Color.BlueStr("• Keep:"), colors.Color.YellowStr(keep))
	fmt.Println(colors.Color.BlueStr("• Diff:"), colors.Color.YellowStr(strconv.FormatBool(diff)))
	fmt.Println("")
}

func GitPull(repo string) {
	out, err := utils.RunCommand(gitPullCmd, repo)
	fmt.Print(out)
	if err != nil {
		errors.ErrorFormat("Error executing git pull", err)
	}
}
func GitDiff(repo string) bool {
	_, err := utils.RunCommand(gitDiffCmd, repo)
	return err != nil
}

func GitStatus(repo string) string {
	out, err := utils.RunCommand(gitStatusCmd, repo)
	if err != nil {
		errors.ErrorFormat("Error executing git status", err)
	}
	return out
}

func GitAdd(repo string) {
	out, err := utils.RunCommand(gitAddCmd, repo)
	fmt.Print(out)
	if err != nil {
		errors.ErrorFormat("Error executing git add", err)
	}
}

func NixUpdate(repo string) {
	out, err := utils.RunCommand(nixUpdateCmd, repo)
	fmt.Print(out)
	if err != nil {
		errors.ErrorFormat("Error executing nix flake update", err)
	}
}

func NixRebuild(repo string, hostname string) {
	out, err := utils.RunCommand(nixRebuildCmd, repo, hostname)
	fmt.Print(out)
	if err != nil {
		errors.ErrorFormat("Error executing nixos rebuild", err)
	}
}

func NixKeep(keep int) {
	out, err := utils.RunCommand(nixKeepCmd, keep)
	fmt.Print(out)
	if err != nil {
		errors.ErrorFormat("Error executing deleting older generations", err)
	}
}

func NixDiff() {
	out, err := utils.RunCommand(nixDiffCmd)
	fmt.Print(out)
	if err != nil {
		errors.ErrorFormat("Error showing nix diff", err)
	}
}
