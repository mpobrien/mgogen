package main

import (
	".."
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"

	"fmt"
	"log"
)

var (
	file     = flag.String("file", "", "help message for file")
	typeName = flag.String("type", "", "typename to process")
	output   = flag.String("output", "", "output file name; default srcdir/<type>_gen.go")
)

func main() {

	flag.Parse()

	g := mgogen.NewGenerator(*typeName, *file)
	err := g.Parse()
	if err != nil {
		log.Fatalf("Parsing file: %v", err)
	}
	src, err := g.Generate()
	if err != nil {
		log.Fatalf("Generating output: %v", err)
	}

	// Write to file.
	outputName := *output
	if outputName == "" {
		baseName := fmt.Sprintf("%s_gen.go", *typeName)
		outputName = filepath.Join(filepath.Dir(*file), strings.ToLower(baseName))
	}
	err = ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
