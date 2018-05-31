package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strings"
	"os"
)

type Entry struct{
	Name string
	Points float64
	ProblemDirectoryPrefix string
	OutputDirectory string
	TemplateFile string
	TestFile string
}

func main(){
	var entries []Entry

	file, err := ioutil.ReadFile("config.json")
	check(err)

	json.Unmarshal(file, &entries)

	fmt.Printf("%+v\n", entries)

	template, err := ioutil.ReadFile(entries[0].ProblemDirectoryPrefix+entries[0].TemplateFile)
	check(err)

	templateString := string(template)
	templateString = strings.Replace(templateString, "%data%", "Hello World", 1)

	fmt.Println(templateString)

	outErr := ioutil.WriteFile(entries[0].ProblemDirectoryPrefix+entries[0].OutputDirectory+entries[0].TemplateFile, []byte(templateString), os.ModeDevice)
	check(outErr)
}

func check(err error){
	if err != nil {
		panic(err)
	}
}