package main

import (
	"os"
)

func handleError(err error, doPanic bool) {
	if err != nil {
		stderr([]byte(err.Error()))
	}

	if err != nil && doPanic {
		os.Exit(1)
	}
}

func filterTasksOnTags(tasks []Task, tags []string) []Task {
	var filteredTasks []Task
	for _, task := range tasks {
		if task.containsTag(tags) {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks
}
