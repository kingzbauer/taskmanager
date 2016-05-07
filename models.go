package main

import (
	"gopkg.in/yaml.v2"
)

// Task represents a single runnable command
type Task struct {
	Name    string   `yaml:"name"`
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
	Dir     string   `yaml:"dir"`
}

// Project identifies a single running instance of the listed tasks
type Project struct {
	Name       string `yaml:"project"`
	WorkingDir string `yaml:"working_dir"`
	Tasks      []Task `yaml:"tasks"`
}

// NewProjectFromFile creates a new project from the contents of a file
func NewProjectFromFile(fileContents []byte) *Project {
	project := new(Project)
	err := yaml.Unmarshal(fileContents, project)
	handleError(err, true)

	return project
}
