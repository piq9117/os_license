package os_license

import (
	"github.com/piq9117/os_license/pkg/os_license"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves an OSS license",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		licenseInput := args[0]
		fileOutputFlag, _ := cmd.Flags().GetString("file-output")

		os_license.Get(licenseInput, fileOutputFlag)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().String("file-output", "", "file-output")
}
