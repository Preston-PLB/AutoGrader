package main

import (
	"io/ioutil"
	"encoding/json"
)

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