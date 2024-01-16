package os_license

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "os_license",
	Short: "os_license - A CLI tool that retreives OSS license",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing CLI '%s'", err)
		os.Exit(1)
	}
}
