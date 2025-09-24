package comments

import (
	"errors"
	"fmt"
	"log"
	"strings"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

type IsLanguageFileFunction func(ext string) bool
type CommentStripperFunction func(comment string) string

type TreeSitterParser struct {
	Parser
	Language       *tree_sitter.Language
	IsLanguageFile IsLanguageFileFunction
	Strip          CommentStripperFunction
}

func Extensions(exts ...string) IsLanguageFileFunction {
	return func(ext string) bool {
		ext = strings.ToLower(ext)
		ext = strings.TrimSpace(ext)
		for _, e := range exts {
			e = strings.ToLower(e)
			e = strings.TrimSpace(e)
			if e == ext {
				return true
			}
		}
		return false
	}
}

func StripPrescedingSlashes() CommentStripperFunction {
	return StripPrescedingCharacters("//")
	// return func(comment string) string {
	// 	comment = strings.TrimSpace(comment)
	// 	if comment[0:2] == "//" {
	// 		comment = comment[2:]
	// 	} else {
	// 		comment = comment[2:]
	// 		comment = comment[:len(comment)-2]
	// 	}
	// 	return comment
	// }
}

func StripPrescedingCharacters(prescedingCharacters string) CommentStripperFunction {
	return func(comment string) string {
		comment = strings.TrimSpace(comment)
		if comment[0:len(prescedingCharacters)] == prescedingCharacters {
			comment = comment[len(prescedingCharacters):]
		} else {
			comment = comment[len(prescedingCharacters):]
			comment = comment[:len(comment)-len(prescedingCharacters)]
		}
		return comment
	}
}

func (p TreeSitterParser) Parse(content *string) (*ParseResult, error) {

	if p.Language == nil {
		log.Panic("no language set up for this tree sitter parser")
	}
	var stripFunction CommentStripperFunction = nil
	if p.Strip != nil {
		stripFunction = p.Strip
	} else {
		// stripFunction = StripPrescedingSlashes
		panic("does not have strip function")
	}

	if content == nil {
		return nil, errors.New("content is nil")
	}
	parser := tree_sitter.NewParser()
	defer parser.Close()

	parser.SetLanguage(p.Language)
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

		if line == lastLine+1 {
			commentBlock.Lines = append(commentBlock.Lines, Line{
				Content:    content,
				LineNumber: line + 1,
			})
			commentBlock.EndLine = line + 1
			lastLine = line
			return
		}

		commentBlock = &CommentBlock{}
		parseResult.Comments = append(parseResult.Comments, commentBlock)
		commentBlock.StartLine = line + 1
		commentBlock.EndLine = line + 1
		commentBlock.Lines = append(commentBlock.Lines, Line{
			Content:    content,
			LineNumber: line + 1,
		})
		lastLine = line

	}

	var evalComment func(*tree_sitter.Node) = func(node *tree_sitter.Node) {

		stringContent := string(node.Utf8Text([]byte(*content)))
		stringContent = strings.TrimSpace(stringContent)

		stringContent = stripFunction(stringContent)

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

func (p TreeSitterParser) ShouldParseFile(extension string) bool {
	if p.IsLanguageFile == nil {
		return false
	}
	return p.IsLanguageFile(extension)
}
