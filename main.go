package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var taskfile string

func init() {
	flag.StringVar(&taskfile, "taskfile", "tasks.yml", "path to your tasks file")
}

func main() {
	flag.Parse()
	// read the tasks.yml file contents
	contents, err := ioutil.ReadFile(taskfile)
	handleError(err, true)

	// create a project
	project := NewProjectFromFile(contents)
	fmt.Println(project)
}
