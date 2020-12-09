package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// get current project's root path
// return path not contain the exec file
func GetProjectRoot() string {
	var (
		path string
		err  error
	)
	defer func() {
		if err != nil {
			panic(fmt.Sprintf("GetProjectRoot error :%+v", err))
		}
	}()
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return path
}

func main() {
	log.Println(GetProjectRoot())
}
