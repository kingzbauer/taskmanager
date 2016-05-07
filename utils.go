package main

func handleError(err error, doPanic bool) {
	if err != nil && doPanic {
		panic(err)
	} else if err != nil {
		// log the error
	}
}
