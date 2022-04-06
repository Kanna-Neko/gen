package cmd

import (
	"fmt"
	"gen/tool"
	"io/ioutil"
	"log"
	"path"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	generateFileName string
	solutionFileName string
	num              int
	start            int
	prefix           string
	inputSuffix      string
	outputSuffix     string
)

func init() {
	gen.PersistentFlags().StringVarP(&prefix, "prefix", "p", "test", "add a prefix to all fileName")
	gen.PersistentFlags().StringVarP(&inputSuffix, "inputSuffix", "i", "in", "add a suffix to all inputFile")
	gen.PersistentFlags().StringVarP(&outputSuffix, "outputSuffix", "o", "out", "add a suffix to all outputFile")
	gen.PersistentFlags().IntVarP(&start, "start", "s", 1, "set a starting sequence number before all files")
	gen.PersistentFlags().IntVarP(&num, "num", "n", 10, "The number of input file and output file")
	log.SetFlags(log.Ldate | log.LstdFlags | log.Lshortfile)
}

var gen = &cobra.Command{
	Use:   "gen generateFileName solutionFileName",
	Short: "gen is a simple tests generator",
	Long:  "A simple tests generator build with love by jaxleof in go and Cobra",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		generateFileName = args[0]
		solutionFileName = args[1]
		generateExt := path.Ext(generateFileName)
		solutionExt := path.Ext(solutionFileName)
		for i := 0; i < num; i++ {
			now := start + i
			inputName := prefix + strconv.Itoa(now) + "." + inputSuffix
			outputName := prefix + strconv.Itoa(now) + "." + outputSuffix
			var inputData []byte
			var outputData []byte
			switch generateExt {
			case ".cpp":
				var err error
				if inputData, err = tool.RunCpp(generateFileName, nil); err != nil {
					log.Fatal(err)
				}
			default:
				fmt.Println("this file format is not supported")
			}
			ioutil.WriteFile(inputName, inputData, 0666)
			switch solutionExt {
			case ".cpp":
				var err error
				if outputData, err = tool.RunCpp(solutionFileName, inputData); err != nil {
					log.Fatal(err)
				}
			default:
				fmt.Println("this file format is not supported")
			}
			ioutil.WriteFile(outputName, outputData, 0666)
			fmt.Println("sample", i+1, "is finished")
		}
	},
}

func Execute() {
	gen.Execute()
}
