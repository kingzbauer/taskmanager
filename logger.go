package main

import (
	"os"

	"fmt"
	"github.com/fatih/color"
	"sync"
)

var redPrint func(string, ...interface{}) string
var greenPrint func(string, ...interface{}) string
var lock sync.Mutex

func init() {
	redPrint = color.New(color.FgRed).SprintfFunc()
	greenPrint = color.New(color.FgGreen).SprintfFunc()
	lock = sync.Mutex{}
}

type writerFunc func([]byte) (int, error)

func (w writerFunc) Write(data []byte) (int, error) {
	return w(data)
}

func stderr(data []byte) (n int, err error) {
	defer lock.Unlock()

	msg := redPrint("%s\n", string(data))
	lock.Lock()

	return fmt.Fprint(os.Stderr, msg)
}

func errString(str string) {
	stderr([]byte(str))
}

func stdout(data []byte) (n int, err error) {
	defer lock.Unlock()

	msg := greenPrint("%s\n", string(data))
	lock.Lock()

	return fmt.Fprint(os.Stdout, msg)
}

func outputString(str string) {
	stdout([]byte(str))
}
