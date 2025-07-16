package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"parser"
	"path/filepath"
	"strings"
)

type TodoBlock struct {
	CommentBlock *parser.CommentBlock
	FileName     string
}

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

	var todoBlocks []TodoBlock = []TodoBlock{}

	var DoFile func(string) = func(file string) {
		if file[len(file)-2:] != "go" {
			// log.Printf("Cannot do file '%s' as it is not a go file", file)
			return
		}

		content, err := os.ReadFile(file)
		if err != nil {
			log.Panicf("File '%s' could not be read.", file)
		}
		//todo yes please
		//haha
		parseResult, err := parser.Parse(Ptr(string(content)), "go")

		if err != nil {
			log.Panic("Error parsing")
		}

		for _, commentBlock := range parseResult.Comments {
			hasTodo := strings.Contains(strings.ToLower(commentBlock.String()), "todo")
			if hasTodo {
				todoBlocks = append(todoBlocks, TodoBlock{
					CommentBlock: commentBlock,
					FileName:     file,
				})
			}
		}
	}

	if !fileOrFolderInfo.IsDir() {
		// log.Println("DOING FILE", *fileOrFolder)
		// todo kill someone
		DoFile(*fileOrFolder)
	} else {
		filepath.Walk(*fileOrFolder, func(path string, info fs.FileInfo, err error) error {

			// log.Println("DOING FILE", path) todo
			if info.IsDir() {
				return nil
			}
			DoFile(path)
			return nil
		})
	}

	fmt.Printf("Got %d comment blocks\n", len(todoBlocks))

	// log.Println("DOING FILE", *fileOrFolder)
	// todo kill someone 22
	// THIS IS NOT BEING PICKED UP WTF

	for _, todoBlock := range todoBlocks {
		fmt.Printf("\n%s: ", todoBlock.FileName)
		for _, line := range todoBlock.CommentBlock.Lines {
			fmt.Printf("%s ", line)
		}
		fmt.Printf("\n")
	}


	//test todo
}

func Ptr[T any](v T) *T {
	return &v
}
