package main

import (
	"os"

	valid "github.com/asaskevich/govalidator"
	"gopkg.in/yaml.v2"
	"os/exec"
)

// Task represents a single runnable command
type Task struct {
	Name    string    `yaml:"name"`
	Command string    `yaml:"command" valid:"required"`
	Args    []string  `yaml:"args"`
	Dir     string    `yaml:"dir"`
	Cmd     *exec.Cmd `yaml:"-"`
}

// Project identifies a single running instance of the listed tasks
type Project struct {
	Name       string `yaml:"project"`
	WorkingDir string `yaml:"working_dir"`
	Tasks      []Task `yaml:"tasks" valid:"required"`
}

// NewProjectFromFile creates a new project from the contents of a file
func NewProjectFromFile(fileContents []byte) *Project {
	project := new(Project)
	err := yaml.Unmarshal(fileContents, project)
	handleError(err, true)

	return project
}

// Validate checks whether the model is valid
func (project Project) Validate() (bool, error) {
	return valid.ValidateStruct(project)
}

func (project *Project) init() {
	// update the working directory if provided
	if len(project.WorkingDir) > 0 {
		os.Chdir(project.WorkingDir)
	}

	// initialize the tasks
	// TODO: do LookPath for each of the tasks
	for i := range project.Tasks {
		project.Tasks[i].init()
	}
}

// Validate checks whether the model is valid
func (task Task) Validate() (bool, error) {
	return valid.ValidateStruct(task)
}

func (task *Task) init() {
	// initialiaze the command struct
	task.Cmd = exec.Command(task.Command, task.Args...)
	// set the stdout and stderr
	task.Cmd.Stderr = writerFunc(stderr)
	task.Cmd.Stdout = writerFunc(stdout)
	// set the working directory of the command
	task.Cmd.Dir = task.Dir
}
