package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	gen.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gen",
	Long:  `All software has versions. This is gen's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen version v1.0")
	},
}
