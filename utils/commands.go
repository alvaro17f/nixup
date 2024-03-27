package utils

import (
	"fmt"
	"os"
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
	border := strings.Repeat("*", textLen+4)
	fmt.Printf(
		"\n%s\n%s %s %s\n%s\n",
		Blue(border), Blue("*"), Red(text), Blue("*"), Blue(border),
	)
}

func ErrorFormat(text string, err error) {
	fmt.Printf(
		"%s %s %s\n",
		Red("⚠"), Yellow(text+":"), Red(err),
	)
}

func ErrorFormatFatal(text string, err error) {
	fmt.Printf(
		"%s %s %s\n",
		Red("⚠"), Yellow(text+":"), Red(err),
	)
	os.Exit(0)
}

func Configuration() {
	fmt.Println(Blue("• Repo:"), Yellow(*flags.Repo))
	fmt.Println(Blue("• Update:"), Yellow(*flags.Update))
	fmt.Println(Blue("• Keep:"), Yellow(*flags.Keep))
	fmt.Println(Blue("• Diff:"), Yellow(*flags.Diff))
	fmt.Println("")
}

func GitPull() {
	out, err := ExecuteCommand(gitPullCmd, *flags.Repo)
	fmt.Print(out)
	if err != nil {
		ErrorFormat("Error executing git pull", err)
	}
}
func GitDiff() bool {
	_, err := ExecuteCommand(gitDiffCmd, *flags.Repo)
	return err != nil
}

func GitStatus() string {
	out, err := ExecuteCommand(gitStatusCmd, *flags.Repo)
	if err != nil {
		ErrorFormat("Error executing git status", err)
	}
	return out
}

func GitAdd() {
	out, err := ExecuteCommand(gitAddCmd, *flags.Repo)
	fmt.Print(out)
	if err != nil {
		ErrorFormat("Error executing git add", err)
	}
}

func NixUpdate() {
	out, err := ExecuteCommand(nixUpdateCmd, *flags.Repo)
	fmt.Print(out)
	if err != nil {
		ErrorFormat("Error executing nix flake update", err)
	}
}

func NixRebuild(hostname string) {
	out, err := ExecuteCommand(nixRebuildCmd, *flags.Repo, hostname)
	fmt.Print(out)
	if err != nil {
		ErrorFormat("Error executing nixos rebuild", err)
	}
}

func NixKeep() {
	out, err := ExecuteCommand(nixKeepCmd, *flags.Keep)
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
