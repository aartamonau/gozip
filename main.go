package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	exitSuccess = 0
	exitFailure = 1
)

var (
	zipPath   string
	paths     []string
	prefix    string
	recursive bool

	name string
	f    *flag.FlagSet
)

func usage(code int) {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintf(os.Stderr, "  %s [options] zipfile [file ...]\n", name)
	fmt.Fprintln(os.Stderr, "Options:")
	f.SetOutput(os.Stderr)
	f.PrintDefaults()

	os.Exit(code)
}

func maybeAddExt(name string) string {
	if filepath.Ext(name) == ".zip" {
		return name
	} else {
		return name + ".zip"
	}
}

func main() {
	name = os.Args[0]

	f = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	f.SetOutput(ioutil.Discard)

	f.BoolVar(&recursive, "recursive", false, "scan directories recursively")
	f.StringVar(&prefix, "prefix", "", "prepend each path with prefix")

	err := f.Parse(os.Args[1:])
	if err == nil {
		if f.NArg() == 0 {
			err = fmt.Errorf("output zip file is not specified")
		} else {
			zipPath = maybeAddExt(f.Arg(0))
			paths = f.Args()[1:]
		}
	}

	if err != nil {
		switch err {
		case flag.ErrHelp:
			usage(exitSuccess)
		default:
			fmt.Fprintln(os.Stderr, err)
			usage(exitFailure)
		}
	}
}
