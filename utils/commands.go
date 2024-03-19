package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/alvaro17f/nixup/flags"
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

func TitleMaker(text string) {
	textLen := len(text)
	borders := strings.Repeat("*", textLen+4)
	fmt.Printf(
		"\n%s\n* %s *\n%s\n",
		borders, text, borders,
	)
}

func GitPull() {
	out, err := ExecuteCommand(gitPullCmd, *flags.Repo)
	fmt.Print(out)
	if err != nil {
		log.Printf("Error executing git pull: %v", err)
	}
}
func GitDiff() bool {
	_, err := ExecuteCommand(gitDiffCmd, *flags.Repo)
	return err != nil
}

func GitStatus() string {
	out, err := ExecuteCommand(gitStatusCmd, *flags.Repo)
	if err != nil {
		log.Printf("Error executing git status: %v", err)
	}
	return out
}

func GitAdd() {
	out, err := ExecuteCommand(gitAddCmd, *flags.Repo)
	fmt.Print(out)
	if err != nil {
		log.Printf("Error executing git add: %v", err)
	}
}

func NixUpdate() {
	out, err := ExecuteCommand(nixUpdateCmd, *flags.Repo)
	fmt.Print(out)
	if err != nil {
		log.Printf("Error executing nix flake update: %v", err)
	}
}

func NixRebuild(hostname string) {
	out, err := ExecuteCommand(nixRebuildCmd, *flags.Repo, hostname)
	fmt.Print(out)
	if err != nil {
		log.Printf("Error executing nixos rebuild: %v", err)
	}
}

func NixKeep() {
	out, err := ExecuteCommand(nixKeepCmd, *flags.Keep)
	fmt.Print(out)
	if err != nil {
		log.Printf("Error executing deleting older generations: %v", err)
	}
}

func NixDiff() {
	out, err := ExecuteCommand(nixDiffCmd)
	fmt.Print(out)
	if err != nil {
		log.Printf("Error showing nix diff: %v", err)
	}
}
