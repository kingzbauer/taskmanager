package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

var taskfile string

func init() {
	flag.StringVar(&taskfile, "taskfile", "tasks.yml", "path to your tasks file")
}

func main() {
	flag.Parse()
	// create a channel to receive sigterm signals for when to terminate
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// read the tasks.yml file contents
	contents, err := ioutil.ReadFile(taskfile)
	handleError(err, true)

	// create a project
	project := NewProjectFromFile(contents)
	// validate the models
	if valid, err := project.Validate(); !valid {
		handleError(err, true)
	}

	go func() {
		_ = <-signals
		project.Exit()
	}()

	// start the project
	project.init()
	project.Start()
}
