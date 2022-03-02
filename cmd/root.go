package cmd

import (
	"bytes"
	"errors"
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
	inputData        [][]byte
)

func init() {
	gen.PersistentFlags().StringVarP(&prefix, "prefix", "p", "test", "add a prefix to all fileName")
	gen.PersistentFlags().StringVarP(&inputSuffix, "inputSuffix", "i", "in", "add a suffix to all inputFile")
	gen.PersistentFlags().StringVarP(&outputSuffix, "outputSuffix", "o", "out", "add a suffix to all outputFile")
	gen.PersistentFlags().IntVarP(&start, "start", "s", 1, "set a starting sequence number before all files")
	wd, _ = os.Getwd()
	log.SetFlags(log.Ldate | log.LstdFlags | log.Lshortfile)
}

var gen = &cobra.Command{
	Use:   "gen generateFileName [solutionFileName] num",
	Short: "gen is a simple tests generator",
	Long:  "A simple tests generator build with love by jaxleof in go and Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(generateFileName, solutionFileName, num) test the input is valid
		// fmt.Println(prefix,inputSuffix, outputSuffix) test the prefix, suffix

		//generate the binary file of generatorFile and generate the inputFile
		var genFileExtension = path.Ext(generateFileName)
		if genFileExtension == ".cpp" {
			cppGen()
		} else {
			log.Fatalln("this version doesn't support this file type: " + genFileExtension)
		}

		// if there is no output File
		if len(args) == 2 {
			return
		}
		// generate the solution binary program
		var solFileExtension = path.Ext(solutionFileName)
		if solFileExtension == ".cpp" {
			cppSol()
		} else {
			log.Fatalln("this version doesn't support this file type: " + solFileExtension)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("accepts least 2 args, receive " + strconv.Itoa(len(args)))
		} else if len(args) > 3 {
			return errors.New("accept most 3 args, receive " + strconv.Itoa(len(args)))
		} else if len(args) == 3 {
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
		} else {
			// only two args [generator file, num]
			if _, err := os.Stat(args[0]); err != nil {
				return err
			} else {
				generateFileName = args[0]
			}
			if tmpNum, err := strconv.Atoi(args[1]); err != nil {
				return err
			} else {
				num = tmpNum
			}
		}
		return nil
	},
}

func Execute() {
	gen.Execute()
}

func cppGen() {
	cmd := exec.Command("g++", generateFileName, "-o", "generator")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	for i := start; i < start+num; i++ {
		var input bytes.Buffer
		generator := exec.Command(wd + "/generator")
		generator.Stdin = os.Stdin
		generator.Stdout = &input
		if err := generator.Run(); err != nil {
			log.Fatal(err)
		}
		var err error
		var tmpData []byte
		tmpData, err = ioutil.ReadAll(&input)
		fmt.Println(tmpData)
		if err != nil {
			log.Fatal(err)
		}
		inputData = append(inputData, tmpData)
		if err := ioutil.WriteFile(prefix+strconv.Itoa(i)+"."+inputSuffix, inputData[i-start], 0666); err != nil {
			log.Fatal(err)
		}
	}
	defer os.Remove("generator")
}

func cppSol() {
	cmd := exec.Command("g++", solutionFileName, "-o", "solution")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	// generate the outPut File
	for i := start; i < start+num; i++ {
		var input = bytes.NewBuffer(inputData[i-start])
		var output bytes.Buffer
		cmd := exec.Command(wd + "/solution")
		cmd.Stdin = input
		cmd.Stdout = &output
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		var outputData, err = ioutil.ReadAll(&output)
		if err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(prefix+strconv.Itoa(i)+"."+outputSuffix, outputData, 0666); err != nil {
			log.Fatal(err)
		}
	}
	defer os.Remove("solution")

}
