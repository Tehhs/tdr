package parser

import (
	"errors"
	"fmt"
	"strings"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_go "github.com/tree-sitter/tree-sitter-go/bindings/go"
)

type GoParser struct {
	Parser
}

func (p GoParser) Parse(content *string) (*ParseResult, error) {
	if content == nil {
		return nil, errors.New("content is nil")
	}
	parser := tree_sitter.NewParser()
	defer parser.Close()

	parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_go.Language()))
	tree := parser.Parse([]byte(*content), nil)
	defer tree.Close()

	if tree == nil {
		fmt.Println("Failed to parse the code")
		return nil, nil
	}

	cursor := tree.RootNode().Walk()
	defer cursor.Close()

	parseResult := ParseResult{}

	var commentBlock *CommentBlock = nil
	var lastLine int = -2
	var evalTodo func(content string, line int) = func(content string, line int) {
		if line == lastLine+1  { 
			commentBlock.Lines = append(commentBlock.Lines, content)
			commentBlock.EndLine = line 
			lastLine = line 
			return 
		}
		if commentBlock != nil { 
			parseResult.Comments = append(parseResult.Comments, *commentBlock)
		}
		commentBlock = &CommentBlock{} 
		commentBlock.StartLine = line
		commentBlock.Lines = append(commentBlock.Lines, content)
		lastLine = line 
	}



	var evalComment func(*tree_sitter.Node) = func(node *tree_sitter.Node) {

		stringContent := string(node.Utf8Text([]byte(*content)))
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

		line := node.EndPosition().Row

		evalTodo(stringContent, int(line))
		
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

	return &parseResult, nil

}

func (p GoParser) ShouldParseFile(extension string) bool {
	if extension[len(extension)-2:] == "go" {
		return true
	}
	return false
}
