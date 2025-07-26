package main

import (
	// "flag"
	// "fmt"
	// "io/fs"
	// "log"
	// "os"
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

	// "path/filepath"
	// "strings"
	"github.com/Tehhs/tdr/pkg/tdrl"

	antlr_v4 "github.com/antlr4-go/antlr/v4"
)

type TodoBlock struct {
	CommentBlock *comments.CommentBlock
	FileName     string
}

type TodoListener struct {
	tdrl.BasetdrlListener
	OnFindTodo *func(tags []string, content string)
}

func (s *TodoListener) EnterTodoRule(ctx *tdrl.TodoRuleContext) {
	var amountOfTags int = 0
	if ctx.TagRule() != nil && ctx.TagRule().TagList().GetChildCount() != 0 {
		amountOfTags = (ctx.TagRule().TagList().GetChildCount() + 1) / 2
	}

	tags := []string{}
	for i := range amountOfTags {
		tags = append(tags, ctx.TagRule().TagList().TAG_ID(i).GetText())

	}

	var contentParts []string = []string{}
	if ctx.MessageContent() != nil {
		for partIndex := range ctx.MessageContent().GetChildCount() {
			contentParts = append(contentParts, fmt.Sprint(ctx.MessageContent().GetChild(partIndex)))
		}

	}

	if s.OnFindTodo != nil {
		var f func(tags []string, content string)
		f = *s.OnFindTodo
		f(tags, strings.Join(contentParts, " "))
	}
}

type ProcessedTag struct {
}
type ProcessedTodo struct {
	Tags             []string
	ProcessedContent string
}

func (t ProcessedTodo) HasTag(checkTag string) bool {
	for _, tag := range t.Tags {
		if strings.ToLower(tag) == strings.ToLower(checkTag) {
			return true
		}
	}
	return false
}

func ProcessTodo(content string) []ProcessedTodo {
	charStream := antlr_v4.NewInputStream(content)
	lexer := tdrl.NewtdrlLexer(charStream)
	lexer.RemoveErrorListeners()

	tokens := antlr_v4.NewCommonTokenStream(lexer, antlr_v4.TokenDefaultChannel)
	parser := tdrl.NewtdrlParser(tokens)
	parser.RemoveErrorListeners()

	tree := parser.Main()

	processedTodos := []ProcessedTodo{}

	listener := &TodoListener{}
	onFindTodo := func(tags []string, content string) {
		processedTodos = append(processedTodos, ProcessedTodo{
			Tags:             tags,
			ProcessedContent: content,
		})
	}
	listener.OnFindTodo = &onFindTodo
	antlr_v4.ParseTreeWalkerDefault.Walk(listener, tree)

	return processedTodos
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

		processedTodos := ProcessTodo(todoContent)

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
