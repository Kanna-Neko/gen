package cmd

import (
	"fmt"
	"gen/tool"
	"io/ioutil"
	"log"
	"path"

	"github.com/spf13/cobra"
)

var (
	vsSolutionFileName string
	vsSolutionFileExt  string
	vsOutputSuffix     string
)

func init() {
	vs.PersistentFlags().StringVarP(&vsOutputSuffix, "outputSuffix", "o", "out", "add a suffix to all outputFile")
	gen.AddCommand(vs)
}

var vs = &cobra.Command{
	Use:   "vs solutionFileName [inputFile...]",
	Short: "vs is a simple output generator by your solutionFile, which input is your inputFile",
	Long:  "A simple tests generator build with love by jaxleof in go and Cobra",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vsSolutionFileName = args[0]
		vsSolutionFileExt = path.Ext(vsSolutionFileName)
		for i := 1; i < len(args); i++ {
			args[i] = path.Base(args[i])
		}
		switch vsSolutionFileExt {
		case ".cpp":
			for i := 1; i < len(args); i++ {
				input, err := ioutil.ReadFile(args[i])
				var outputFileName = path.Base(args[i])[:len(args[i])-len(path.Ext(args[i]))] + "." + vsOutputSuffix
				if err != nil {
					log.Fatal(err)
				}
				output, err := tool.RunCpp(vsSolutionFileName, input)
				if err != nil {
					log.Fatal(err)
				}
				ioutil.WriteFile(outputFileName, output, 0666)
			}
		default:
			fmt.Println("The format of file is not supported")
		}
	},
}
