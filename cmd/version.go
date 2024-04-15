package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of nixup",
	Long:  `All software has versions. This is nixup's`,
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("git", "describe", "--tags").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("nixup %s\n", out)
	},
}
