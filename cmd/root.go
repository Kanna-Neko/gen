package cmd

import (
	"bytes"
	"errors"
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
	gen.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "add a prefix to all fileName")
	gen.PersistentFlags().StringVarP(&inputSuffix, "inputSuffix", "i", "in", "add a suffix to all inputFile")
	gen.PersistentFlags().StringVarP(&outputSuffix, "outputSuffix", "o", "out", "add a suffix to all outputFile")
	gen.PersistentFlags().IntVarP(&start, "start", "s", 1, "set a starting sequence number before all files")
	wd, _ = os.Getwd()
}

var gen = &cobra.Command{
	Use:   "gen generateFileName solutionFileName num",
	Short: "gen is a simple tests generator",
	Long:  "A simple tests generator build with love by jaxleof in go and Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(generateFileName, solutionFileName, num) test the input is valid
		// fmt.Println(prefix,inputSuffix, outputSuffix) test the prefix, suffix

		//generate the binary file of generatorFile and solution File
		var genFileExtension = path.Ext(generateFileName)
		if genFileExtension == ".cpp" {
			cmd := exec.Command("g++", generateFileName, "-o", "generator")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatalln("this version doesn't support this file type: " + genFileExtension)
		}
		var solFileExtension = path.Ext(solutionFileName)
		if solFileExtension == ".cpp" {
			cmd := exec.Command("g++", solutionFileName, "-o", "solution")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatalln("this version doesn't support this file type: " + solFileExtension)
		}

		// generate the inputFile
		for i := start; i < start+num; i++ {
			var input bytes.Buffer
			cmd := exec.Command(wd + "/generator")
			cmd.Stdout = &input
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
			var inputData, err = ioutil.ReadAll(&input)
			if err != nil {
				log.Fatal(err)
			}
			ioutil.WriteFile(prefix+strconv.Itoa(i)+"."+inputSuffix, inputData, 0666)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return errors.New("accepts 3 args, receive " + strconv.Itoa(len(args)))
		} else {
			if _, err := os.Stat(args[0]); err != nil {
				return err
			} else {
				generateFileName = args[0]
			}
			if _, err := os.Stat(args[1]); err != nil {
				return err
			} else {
				solutionFileName = args[1]
			}
			if tmpNum, err := strconv.Atoi(args[2]); err != nil {
				return err
			} else {
				num = tmpNum
			}
			return nil
		}
	},
}

func Execute() {
	gen.Execute()
}
