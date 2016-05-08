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
