package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/Tehhs/tdr/pkg/comments"
	"github.com/Tehhs/tdr/pkg/util"
	"github.com/Tehhs/tdr/pkg/tdrl"
)

type TodoBlock struct {
	CommentBlock *comments.CommentBlock
	FileName     string
}


type ProcessedTag struct {
}

type ProcessArguments struct {
	List      *bool
	TagFilter *string
	Path      *string
}

func ParseArguments() ProcessArguments {
	var processArguments ProcessArguments = ProcessArguments{
		List:      flag.Bool("l", false, "List all tags"),
		TagFilter: flag.String("t", "", "Tag"),
		Path:      flag.String("f", ".", "File or folder"),
	}

	flag.Parse()

	return processArguments
}

func ProcessFile(path string) []TodoBlock {

	var todoCommentBlocks []TodoBlock = []TodoBlock{}

	var extension *string = nil
	filePathParts := strings.Split(path, ".")
	extension = util.Ptr(filePathParts[len(filePathParts)-1])

	content, err := os.ReadFile(path)
	if err != nil {
		log.Panicf("File '%s' could not be read.", path)
	}

	//todo(refactor): Comments layer should return comment layer errors, and 
	//have a special file extension not supported error to check for here
	//instead of just ingoring.
	parseResult, err := comments.Parse(Ptr(string(content)), *extension)

	if err != nil {
		// log.Panic("Error parsing")
	}

	if parseResult == nil || parseResult.Comments == nil {
		return todoCommentBlocks
	}

	for _, commentBlock := range parseResult.Comments {
		hasTodo := strings.Contains(strings.ToLower(commentBlock.String()), "todo")
		if hasTodo {
			todoCommentBlocks = append(todoCommentBlocks, TodoBlock{
				CommentBlock: commentBlock,
				FileName:     path,
			})
		}
	}

	return todoCommentBlocks

}


func main() {
	tdrlParser := tdrl.NewParser()
	args := ParseArguments()

	fileOrFolderInfo, err := os.Stat(*args.Path)
	if err != nil {
		log.Panicf("File or folder '%s' is not a file or folder.\n", *args.Path)
	}

	var todoCommentBlocks []TodoBlock = []TodoBlock{}

	if !fileOrFolderInfo.IsDir() {
		newBlocks := ProcessFile(*args.Path)
		todoCommentBlocks = append(todoCommentBlocks, newBlocks...)
	} else {
		filepath.Walk(*args.Path, func(path string, info fs.FileInfo, err error) error {

			if info.IsDir() {
				return nil
			}
			newBlocks := ProcessFile(path)
			todoCommentBlocks = append(todoCommentBlocks, newBlocks...)

			return nil
		})
	}

	listedTags := []string{}
	for _, todoBlock := range todoCommentBlocks {

		todoLines := []string{}
		for _, line := range todoBlock.CommentBlock.Lines {
			todoLines = append(todoLines, line.String())
		}
		todoContent := strings.Join(todoLines, " ")

		todoContent = strings.ToLower(todoContent)

		processedTodos := tdrlParser.ProcessTodo(todoContent)

		//i know this is a code smell and a half will refactor later :P
		if *args.List {
			for _, todo := range processedTodos {
				for _, tag := range todo.Tags {
					if slices.Contains(listedTags, strings.TrimSpace(tag)) {
						continue
					}
					listedTags = append(listedTags, strings.TrimSpace(tag))
				}
			}
		} else if len(processedTodos) != 0 {
			for _, tag := range processedTodos[0].Tags {
				if strings.Contains(strings.ToLower(tag), strings.ToLower(*args.TagFilter)) {
					fmt.Printf("\n%s(lines %d to %d):\n", todoBlock.FileName, todoBlock.CommentBlock.StartLine, todoBlock.CommentBlock.EndLine)
					fmt.Printf("\t%v -> %s", processedTodos[0].Tags, processedTodos[0].ProcessedContent)
					break
				}
			}

		}

		//todo(very_important): test this thing

		//todo(important): another test

		//todo(easy,ez): this should be easy and
		//when this breaks into multiple lines
		//it should handle that
	}

	if *args.List {
		fmt.Print("List of all todo tags:\n\n")
		for _, tag := range listedTags {
			fmt.Println(tag)
		}
	}

	fmt.Print("\n")

}

func Ptr[T any](v T) *T {
	return &v
}
