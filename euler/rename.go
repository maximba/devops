package main

import "fmt"
import "os/exec"
import "strings"
import "path/filepath"

type Worker struct {
	Command   string
	Args      []string
}

func (cmd *Worker) Run() {
	out, err := exec.Command(cmd.Command, cmd.Args...).Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", out)
}

func main() {
	folder := "./problem*.go"
	files, errorcode := filepath.Glob(folder)
	if errorcode != nil {
		panic(errorcode)
	}

	for _, infile := range files {
		prefix := "problem"
		outfile := "p" + strings.TrimLeft(infile, prefix)
		args := []string{infile, outfile}
		w := &Worker{Command: "cp", Args: args}
		w.Run()
	}
}
