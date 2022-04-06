package tool

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
)

var wd string

func init() {
	wd, _ = os.Getwd()
}

func RunCpp(fileName string, input []byte) ([]byte, error) {
	cmd := exec.Command("g++", fileName, "-o", "catt")
	cmd.Stderr = os.Stderr
	defer os.Remove("catt")
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
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
