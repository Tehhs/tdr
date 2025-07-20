package main

import (
	// "flag"
	// "fmt"
	// "io/fs"
	// "log"
	// "os"
	"comments"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	// "path/filepath"
	// "strings"
	"tdrl"

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

func main() {

	fileOrFolder := flag.String("f", ".", "File or folder")
	if fileOrFolder == nil {
		log.Panic("Missing file or folder. Use -f to specify file or folder")
	}

	filterTag := flag.String("t", "", "Tag")

	flag.Parse()

	fileOrFolderInfo, err := os.Stat(*fileOrFolder)
	if err != nil {
		log.Panicf("File or folder '%s' is not a file or folder.\n", *fileOrFolder)
	}

	var todoBlocks []TodoBlock = []TodoBlock{}

	var DoFile func(string) = func(file string) {
		if file[len(file)-2:] != "go" {

			return
		}

		content, err := os.ReadFile(file)
		if err != nil {
			log.Panicf("File '%s' could not be read.", file)
		}
		//todo yes please
		//haha
		parseResult, err := comments.Parse(Ptr(string(content)), "go")

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
		DoFile(*fileOrFolder)
	} else {
		filepath.Walk(*fileOrFolder, func(path string, info fs.FileInfo, err error) error {

			if info.IsDir() {
				return nil
			}
			DoFile(path)
			return nil
		})
	}


	// log.Println("DOING FILE", *fileOrFolder)
	// todo this is a test todo

	for _, todoBlock := range todoBlocks {

		todoLines := []string{}
		for _, line := range todoBlock.CommentBlock.Lines {
			todoLines = append(todoLines, line.String())
		}
		todoContent := strings.Join(todoLines, " ")

		todoContent = strings.ToLower(todoContent)

		processedTodos := ProcessTodo(todoContent)
		if len(processedTodos) != 0 {
			for _, tag := range processedTodos[0].Tags {
				if strings.Contains(strings.ToLower(tag), strings.ToLower(*filterTag)) { 
					fmt.Printf("\n%s(lines %d to %d):\n", todoBlock.FileName, todoBlock.CommentBlock.StartLine, todoBlock.CommentBlock.EndLine)
					fmt.Printf("\t%v -> %s", processedTodos[0].Tags, processedTodos[0].ProcessedContent)
					break; 
				}
			}
			
		}

		//todo(very_important): test this thing

		//todo(important): another test 

		//todo(easy,ez): this should be easy and 
		//when this breaks into multiple lines
		//it should handle that 
	}

	fmt.Print("\n")

}

func Ptr[T any](v T) *T {
	return &v
}
