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

func commands() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use:                   "completion [bash|zsh|fish|powershell]",
			Short:                 "Generate completion script",
			Long:                  "To load completions",
			DisableFlagsInUseLine: true,
			ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
			Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
			Run: func(cmd *cobra.Command, args []string) {
				switch args[0] {
				case "bash":
					if err := cmd.Root().GenBashCompletion(os.Stdout); err != nil {
						cmd.PrintErr(err)
					}
				case "zsh":
					if err := cmd.Root().GenZshCompletion(os.Stdout); err != nil {
						cmd.PrintErr(err)
					}
				case "fish":
					if err := cmd.Root().GenFishCompletion(os.Stdout, true); err != nil {
						cmd.PrintErr(err)
					}
				case "powershell":
					if err := cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout); err != nil {
						cmd.PrintErr(err)
					}
				}
			},
		},
	)
}

func flags() {
	rootCmd.PersistentFlags().BoolP(diff.getFlagDetails())
	rootCmd.PersistentFlags().StringP(hostname.getFlagDetails())
	rootCmd.PersistentFlags().IntP(keep.getFlagDetails())
	rootCmd.PersistentFlags().StringP(repo.getFlagDetails())
	rootCmd.PersistentFlags().BoolP(update.getFlagDetails())
}

func Execute() {
	commands()
	flags()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func SetVersionInfo(version string) {
	rootCmd.Version = version
}
