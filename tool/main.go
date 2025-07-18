package main

import (
	// "flag"
	// "fmt"
	// "io/fs"
	// "log"
	// "os"
	"comments"
	"log"

	// "path/filepath"
	// "strings"
	"tdrl"

	antlr_v4 "github.com/antlr4-go/antlr/v4"
)

type TodoBlock struct {
	CommentBlock *comments.CommentBlock
	FileName     string
}

// Custom listener to print parsed todo comment details
type PrintTodoListener struct {
	antlr_v4.BaseParseTreeListener
}

// EnterTodoComment is called when entering the todoComment rule
func (l *PrintTodoListener) EnterTodoComment(ctx *tdrl.TodoContext) {
	log.Println("Parsed a todo comment:")
	log.Println(ctx.GetText())
}

func main() {
	// Example input that matches your grammar
	input := "todo:dsdfga"

	charStream := antlr_v4.NewInputStream(input)
	lexer := tdrl.NewtdrlLexer(charStream)
	tokens := antlr_v4.NewCommonTokenStream(lexer, antlr_v4.TokenDefaultChannel)
	parser := tdrl.NewtdrlParser(tokens)

	tree := parser.Todo()

	listener := &PrintTodoListener{}
	antlr_v4.ParseTreeWalkerDefault.Walk(listener, tree)

	return

	// ...existing code...

	// fileOrFolder := flag.String("f", ".", "File or folder")
	// if fileOrFolder == nil {
	// 	log.Panic("Missing file or folder. Use -f to specify file or folder")
	// }

	// fileOrFolderInfo, err := os.Stat(*fileOrFolder)
	// if err != nil {
	// 	log.Panicf("File or folder '%s' is not a file or folder.\n", *fileOrFolder)
	// }

	// var todoBlocks []TodoBlock = []TodoBlock{}

	// var DoFile func(string) = func(file string) {
	// 	if file[len(file)-2:] != "go" {

	// 		return
	// 	}

	// 	content, err := os.ReadFile(file)
	// 	if err != nil {
	// 		log.Panicf("File '%s' could not be read.", file)
	// 	}
	// 	//todo yes please
	// 	//haha
	// 	parseResult, err := comments.Parse(Ptr(string(content)), "go")

	// 	if err != nil {
	// 		log.Panic("Error parsing")
	// 	}

	// 	for _, commentBlock := range parseResult.Comments {
	// 		hasTodo := strings.Contains(strings.ToLower(commentBlock.String()), "todo")
	// 		if hasTodo {
	// 			todoBlocks = append(todoBlocks, TodoBlock{
	// 				CommentBlock: commentBlock,
	// 				FileName:     file,
	// 			})
	// 		}
	// 	}
	// }

	// if !fileOrFolderInfo.IsDir() {
	// 	// log.Println("DOING FILE", *fileOrFolder)
	// 	// todo todo
	// 	DoFile(*fileOrFolder)
	// } else {
	// 	filepath.Walk(*fileOrFolder, func(path string, info fs.FileInfo, err error) error {

	// 		// log.Println("DOING FILE", path) todo
	// 		if info.IsDir() {
	// 			return nil
	// 		}
	// 		DoFile(path)
	// 		return nil
	// 	})
	// }

	// fmt.Printf("Got %d comment blocks\n", len(todoBlocks))

	// // log.Println("DOING FILE", *fileOrFolder)
	// // todo this is a test todo

	// for _, todoBlock := range todoBlocks {
	// 	fmt.Printf("\n%s(lines %d to %d):\n", todoBlock.FileName, todoBlock.CommentBlock.StartLine, todoBlock.CommentBlock.EndLine)
	// 	for _, line := range todoBlock.CommentBlock.Lines {
	// 		fmt.Printf("\t%s\n", line)
	// 	}
	// 	fmt.Printf("\n")
	// }

	//test todo
}

func Ptr[T any](v T) *T {
	return &v
}
