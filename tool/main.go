package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"parser"
	"path/filepath"
)

//this should all
//be on the same
//comment block

//but not this

func main() {
	fmt.Println()

	fileOrFolder := flag.String("f", ".", "File or folder")
	if fileOrFolder == nil {
		log.Panic("Missing file or folder. Use -f to specify file or folder")
	}

	fileOrFolderInfo, err := os.Stat(*fileOrFolder)
	if err != nil {
		log.Panicf("File or folder '%s' is not a file or folder.\n", *fileOrFolder)
	}

	var DoFile func(string) = func(file string) {
		if file[len(file)-2:] != "go" {
			// log.Printf("Cannot do file '%s' as it is not a go file", file)
			return
		}

		content, err := os.ReadFile(file)
		if err != nil {
			log.Panicf("File '%s' could not be read.", file)
		}

		parseResult, err := parser.Parse(Ptr(string(content)), "go")

		if err != nil {
			log.Panic("Error parsing")
		}

		for _, commentBlock := range parseResult.Comments {
			log.Printf("Found comment block: ")
			for _, line := range commentBlock.Lines {
				log.Printf(line + ", ")
			}
			log.Print("\n")
		}
	}

	if !fileOrFolderInfo.IsDir() {
		// log.Println("DOING FILE", *fileOrFolder)
		DoFile(*fileOrFolder)
	} else {
		filepath.Walk(*fileOrFolder, func(path string, info fs.FileInfo, err error) error {

			// log.Println("DOING FILE", path)
			if info.IsDir() {
				return nil
			}
			DoFile(path)
			return nil
		})
	}

	fmt.Println()

}

func Ptr[T any](v T) *T {
	return &v
}
