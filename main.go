package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

var (
	taskfile string
	tags     string
)

func init() {
	flag.StringVar(&taskfile, "taskfile", "tasks.yml", "path to your tasks file")
	flag.StringVar(&tags, "tags", "", "a list of tags separated with commans")
}

func main() {
	flag.Parse()
	// create a channel to receive sigterm signals for when to terminate
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// read the tasks.yml file contents
	contents, err := ioutil.ReadFile(taskfile)
	handleError(err, true)

	// Retrieve and process tags if any
	processedTags := strings.Split(
		regexp.MustCompile("\\s+").ReplaceAllString(tags, ""), ",")
	if len(tags) == 0 {
		processedTags = []string{}
	}

	// create a project
	project := NewProjectFromFile(contents, processedTags)
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
