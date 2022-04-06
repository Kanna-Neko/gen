package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
	wd               string
)

func init() {
	gen.PersistentFlags().StringVarP(&prefix, "prefix", "p", "test", "add a prefix to all fileName")
	gen.PersistentFlags().StringVarP(&inputSuffix, "inputSuffix", "i", "in", "add a suffix to all inputFile")
	gen.PersistentFlags().StringVarP(&outputSuffix, "outputSuffix", "o", "out", "add a suffix to all outputFile")
	gen.PersistentFlags().IntVarP(&start, "start", "s", 1, "set a starting sequence number before all files")
	gen.PersistentFlags().IntVarP(&num, "num", "n", 10, "The number of input file and output file")
	wd, _ = os.Getwd()
	log.SetFlags(log.Ldate | log.LstdFlags | log.Lshortfile)
}

var gen = &cobra.Command{
	Use:   "gen generateFileName [solutionFileName] num",
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
				if inputData, err = runCpp(generateFileName, nil); err != nil {
					log.Fatal(err)
				}
			default:
				fmt.Println("this file format is not supported")
			}
			ioutil.WriteFile(inputName, inputData, 0666)
			switch solutionExt {
			case ".cpp":
				var err error
				if outputData, err = runCpp(solutionFileName, inputData); err != nil {
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

func runCpp(fileName string, input []byte) ([]byte, error) {
	cmd := exec.Command("g++", fileName, "-o", "catt")
	defer os.Remove("catt")
	cmd.Run()
	cmd = exec.Command(wd + "/catt")
	if input != nil {
		cmd.Stdin = bytes.NewBuffer(input)
	}
	var output = new(bytes.Buffer)
	cmd.Stdout = output
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return ioutil.ReadAll(output)
}
