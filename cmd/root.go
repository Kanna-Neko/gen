package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gen = &cobra.Command{
	Use:   "gen generateFileName solutionFileName num",
	Short: "gen is a simple tests generator",
	Long:  "A simple tests generator build with love by jaxleof in go and Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("helloWorld")
	},
}

func Execute() {
	gen.Execute()
}
