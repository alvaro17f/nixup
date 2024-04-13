package cmd

import (
	"os"

	"github.com/alvaro17f/nixup/internal/utils"
	"github.com/spf13/cobra"
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
	rootCmd.PersistentFlags().BoolP("diff", "d", false, "Show the diff of the last generation")
	rootCmd.PersistentFlags().StringP("hostname", "n", utils.GetHostname(), "Set the hostname")
	rootCmd.PersistentFlags().IntP("keep", "k", 10, "Keep last generations")
	rootCmd.PersistentFlags().StringP("repo", "r", "~/.dotfiles", "Path to the git repository")
	rootCmd.PersistentFlags().BoolP("update", "u", false, "Update the system")
}
