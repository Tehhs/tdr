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

func DoContent(content []byte, fileName *string) (*Todo, error) {
	if content == nil {
		return nil, errors.New("content is nil")
	}

	//todo unfuck this pls 
	//#TODO (NOW, IMPORTANT): todo unfuck this pls [#2w340857]
	parser := tree_sitter.NewParser()
	defer parser.Close()
	parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_go.Language()))

	tree := parser.Parse(content, nil)
	defer tree.Close()

	if tree == nil {
		fmt.Println("Failed to parse the code")
		return nil, nil
	}

	cursor := tree.RootNode().Walk()
	defer cursor.Close()

	// root := tree.RootNode()

	// fmt.Println(root.ToSexp())

	var evalTodo func(content string, line int) = func(content string, line int) {
		// fmt.Println("Found todo", content)
		fmt.Printf("%s, Line %d: %s\n", *fileName, line, content)
	}

	var evalComment func(*tree_sitter.Node) = func(node *tree_sitter.Node) {

		stringContent := string(node.Utf8Text(content))
		stringContent = strings.TrimSpace(stringContent)

		if stringContent[0:2] == "//" {
			stringContent = stringContent[2:]
		} else {
			stringContent = stringContent[2:]
			stringContent = stringContent[:len(stringContent)-2]
		}
		// fmt.Println(node.)
		// fmt.Println(string(node.Utf8Text(code)))
		// fmt.Println(node.EndPosition().Row)

		stringContent = strings.TrimSpace(stringContent)
		line := node.EndPosition().Row

		if strings.Contains(strings.ToLower(stringContent), "todo") {
			evalTodo(stringContent, int(line))
		}
	}

	var evalNode func(*tree_sitter.Node) = func(node *tree_sitter.Node) {
		if node.GrammarName() != "comment" {
			return
		}
		evalComment(node)
	}

	for {

		//eval element here
		evalNode(cursor.Node())

		if cursor.GotoFirstChild() {
			continue
		}

		if cursor.GotoNextSibling() {
			continue
		}

		//Infinitly go backwards and try to find next ancestor sibling
		var shouldBreak bool = false
		for {
			if cursor.GotoParent() {
				if cursor.GotoNextSibling() {
					break
				}
			} else {
				shouldBreak = true
				break
			}
		}
		if shouldBreak {
			break
		}

	}

	return nil, nil

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
		// todo todo 
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
	// todo this is a test todo

	for _, todoBlock := range todoBlocks {
		fmt.Printf("\n%s(lines %d to %d):\n", todoBlock.FileName, todoBlock.CommentBlock.StartLine, todoBlock.CommentBlock.EndLine)
		for _, line := range todoBlock.CommentBlock.Lines {
			fmt.Printf("\t%s\n", line)
		}
		fmt.Printf("\n")
	}


	//test todo
}

func Ptr[T any](v T) *T {
	return &v
}
