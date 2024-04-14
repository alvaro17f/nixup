package cmd

import (
	"os"

	"github.com/alvaro17f/nixup/internal/utils"
	"github.com/spf13/cobra"
)

const (
	RepoFlag     = "repo"
	HostnameFlag = "hostname"
	DiffFlag     = "diff"
	KeepFlag     = "keep"
	UpdateFlag   = "update"
)

var rootCmd = &cobra.Command{
	Use:   "nixup",
	Short: "Update your nixos system with a single command",
	Long: `
**********
* NIXUP  *
**********
a tool to update your nixos system with a single command.
`,
	Run: func(cmd *cobra.Command, args []string) {
		Nixup(cmd, args)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP(DiffFlag, "d", false, "Show the diff of the last generation")
	rootCmd.PersistentFlags().StringP(HostnameFlag, "n", utils.GetHostname(), "Set the hostname")
	rootCmd.PersistentFlags().IntP(KeepFlag, "k", 10, "Keep last generations")
	rootCmd.PersistentFlags().StringP(RepoFlag, "r", "~/.dotfiles", "Path to the git repository")
	rootCmd.PersistentFlags().BoolP(UpdateFlag, "u", false, "Update the system")
}
