package cmd

import (
	"os"

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
	Run: func(cmd *cobra.Command, _ []string) {
		Nixup(cmd)
	},
}

func flags() {
	rootCmd.PersistentFlags().BoolP(diff.getFlagDetails())
	rootCmd.PersistentFlags().StringP(hostname.getFlagDetails())
	rootCmd.PersistentFlags().IntP(keep.getFlagDetails())
	rootCmd.PersistentFlags().StringP(repo.getFlagDetails())
	rootCmd.PersistentFlags().BoolP(update.getFlagDetails())
}

func Execute() {
	flags()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func SetVersionInfo(version string) {
	rootCmd.Version = version
}
