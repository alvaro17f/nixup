package utils

import (
	"fmt"
	"strconv"
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

	fmt.Println(Color.BlueStr("• Repo:"), Color.YellowStr(repo))
	fmt.Println(Color.BlueStr("• Hostname:"), Color.YellowStr(hostname))
	fmt.Println(Color.BlueStr("• Update:"), Color.YellowStr(strconv.FormatBool(update)))
	fmt.Println(Color.BlueStr("• Keep:"), Color.YellowStr(keep))
	fmt.Println(Color.BlueStr("• Diff:"), Color.YellowStr(strconv.FormatBool(diff)))
	fmt.Println("")
}

func GitPull(repo string) {
	out, err := ExecuteCommand(gitPullCmd, repo)
	fmt.Print(out)
	if err != nil {
		ErrorFormat("Error executing git pull", err)
	}
}
func GitDiff(repo string) bool {
	_, err := ExecuteCommand(gitDiffCmd, repo)
	return err != nil
}

func GitStatus(repo string) string {
	out, err := ExecuteCommand(gitStatusCmd, repo)
	if err != nil {
		ErrorFormat("Error executing git status", err)
	}
	return out
}

func GitAdd(repo string) {
	out, err := ExecuteCommand(gitAddCmd, repo)
	fmt.Print(out)
	if err != nil {
		ErrorFormat("Error executing git add", err)
	}
}

func NixUpdate(repo string) {
	out, err := ExecuteCommand(nixUpdateCmd, repo)
	fmt.Print(out)
	if err != nil {
		ErrorFormat("Error executing nix flake update", err)
	}
}

func NixRebuild(repo string, hostname string) {
	out, err := ExecuteCommand(nixRebuildCmd, repo, hostname)
	fmt.Print(out)
	if err != nil {
		ErrorFormat("Error executing nixos rebuild", err)
	}
}

func NixKeep(keep int) {
	out, err := ExecuteCommand(nixKeepCmd, keep)
	fmt.Print(out)
	if err != nil {
		ErrorFormat("Error executing deleting older generations", err)
	}
}

func NixDiff() {
	out, err := ExecuteCommand(nixDiffCmd)
	fmt.Print(out)
	if err != nil {
		ErrorFormat("Error showing nix diff", err)
	}
}
