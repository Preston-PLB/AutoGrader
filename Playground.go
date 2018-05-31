package main

import (
	"io/ioutil"
	"strings"
	"os"
	"encoding/json"
	"os/exec"
	"fmt"
)

func main(){
	problems := Parse("config.json")

	for i := 0; i< len(problems); i++{
		templateFile, fileReadError := ioutil.ReadFile(problems[i].ProblemDirectoryPrefix+problems[i].TemplateFile)
		check(fileReadError)

		templateString := string(templateFile)
		templateString = strings.Replace(templateString, "%data%", "Hellow World, BJ has da BIG GAY", 1)

		 fileWriteError := ioutil.WriteFile(problems[i].ProblemDirectoryPrefix+problems[0].OutputDirectory+problems[i].TemplateFile, []byte(templateString), os.ModeDevice)
		 check(fileWriteError)

		 out, commandErr := exec.Command("echo Hello!").Output()
		 check(commandErr)

		 fmt.Println(out)
	}
}

type Problem struct{
	Name string
	Points float64
	ProblemDirectoryPrefix string
	OutputDirectory string
	TemplateFile string
	TestFile string
}

func Parse(configFilePath string) []Problem{
	var out []Problem

	configFileData, fileReadError := ioutil.ReadFile(configFilePath)
	check(fileReadError)

	jsonParseError := json.Unmarshal(configFileData, &out)
	check(jsonParseError)

	return out
}

func check(err error){
	if err != nil {
		panic(err)
	}
}