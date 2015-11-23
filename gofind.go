package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/docker/docker/pkg/mflag"
)

var (
	h    bool
	path string
	name string
)

func init() {
	// Check https://github.com/docker/docker/blob/master/pkg/mflag/example/example.go
	flag.StringVar(&path, []string{"p", "#pathhidden", "-path"}, "", "path to traverse")
	flag.StringVar(&path, []string{"n", "#namehidden", "-name"}, "", "name to traverse")
	flag.BoolVar(&h, []string{"h", "#help", "-help"}, false, "display the help")
	flag.Parse()
}

func finder(path string, name string) error {

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stdout, "Usage: gofind [OPTIONS] COMMAND [arg...]\ngofind [ --help | -v | --version ]\n")
		flag.CommandLine.SetOutput(os.Stdout)
		flag.PrintDefaults()
	}
	if h {
		flag.Usage()
	} else {
		fmt.Printf("Path: %s\n", path)
		fmt.Printf("Name: %s\n", name)
		fmt.Printf("ARGS: %v\n", flag.Args())
		err := finder(path, name)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
