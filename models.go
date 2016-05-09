package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	valid "github.com/asaskevich/govalidator"
	"gopkg.in/yaml.v2"
)

// Task represents a single runnable command
type Task struct {
	Name    string    `yaml:"name"`
	Command string    `yaml:"command" valid:"required"`
	Args    []string  `yaml:"args"`
	Dir     string    `yaml:"dir"`
	Tags    []string  `yaml:"tags"`
	Cmd     *exec.Cmd `yaml:"-"`
	project *Project
}

// Project identifies a single running instance of the listed tasks
type Project struct {
	Name       string `yaml:"project"`
	WorkingDir string `yaml:"working_dir"`
	Tasks      []Task `yaml:"tasks" valid:"required"`
	wg         *sync.WaitGroup
	tags       []string
}

// NewProjectFromFile creates a new project from the contents of a file
func NewProjectFromFile(fileContents []byte, tags []string) *Project {
	project := new(Project)
	err := yaml.Unmarshal(fileContents, project)
	handleError(err, true)

	project.tags = tags

	// if any tags have been specified, filter out tasks with the given tags
	if len(tags) > 0 {
		project.Tasks = filterTasksOnTags(project.Tasks, tags)
	}
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
	// initialize the waitgroup
	project.wg = new(sync.WaitGroup)

	// initialize the tasks
	// TODO: do LookPath for each of the tasks
	for i := range project.Tasks {
		project.Tasks[i].init(project)
	}
}

// Start starts the tasks under the project
func (project *Project) Start() {
	outputString("Starting tasks...")
	for i := range project.Tasks {
		project.wg.Add(1)
		go project.Tasks[i].Start()
	}

	// wait for all the tasks to exit
	project.wg.Wait()
	outputString("Bye.")
}

// Exit causes all the tasks running under this project to exit
func (project *Project) Exit() {
	for i := range project.Tasks {
		task := project.Tasks[i]
		if task.Cmd != nil {
			task.Cmd.Process.Kill()
		}
	}
}

// Validate checks whether the model is valid
func (task Task) Validate() (bool, error) {
	return valid.ValidateStruct(task)
}

func (task *Task) init(project *Project) {
	task.project = project
	// initialiaze the command struct
	task.Cmd = exec.Command(task.Command, task.Args...)
	// set the stdout and stderr
	task.Cmd.Stderr = os.Stderr
	task.Cmd.Stdout = os.Stdout
	// set the working directory of the command
	task.Cmd.Dir = task.Dir
}

// Start starts the task with the given command
func (task *Task) Start() {
	defer task.project.wg.Done()

	outputString(fmt.Sprintf("Starting task [%s]", task.Name))
	if err := task.Cmd.Start(); err != nil {
		handleError(err, false)
		return
	}
	outputString(fmt.Sprintf("Started task [%s] successfully", task.Name))

	// wait for the task to finish
	if err := task.Cmd.Wait(); err != nil {
		handleError(err, false)
		return
	}

	outputString(fmt.Sprintf("Task [%s] has exited", task.Name))
}

func (task *Task) hasTag(tag string) bool {
	for _, ownTag := range task.Tags {
		if ownTag == tag {
			return true
		}
	}

	return false
}

func (task *Task) containsTag(tags []string) bool {
	for _, tag := range tags {
		if task.hasTag(tag) {
			return true
		}
	}

	return false
}
