package models

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
