package comments

import (
	"errors"
	"log"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_go "github.com/tree-sitter/tree-sitter-go/bindings/go"
	tree_sitter_js "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
)

//todo(refactor): most tree sitter parsers are all the same you should make a util func for that instead of copy pasta
type Line struct {
	Content    string
	LineNumber int
}

func (l Line) String() string {
	return l.Content
}

type CommentBlock struct {
	Lines     []Line
	StartLine int
	EndLine   int
}

func (b CommentBlock) String() string {
	var output string = ""
	for _, l := range b.Lines {
		output += l.String() + "\n"
	}
	return output
}

type ParseResult struct {
	Comments []*CommentBlock
}

type Parser interface {
	Parse(content *string) (*ParseResult, error)
	ShouldParseFile(extension string) bool
}

func ParserFromTreeSitter() { 
	
}

var parsers []Parser = []Parser{
	TreeSitterParser{
		Language: tree_sitter.NewLanguage(tree_sitter_go.Language()),
		IsLanguageFile: Extensions("go"),
	},
	TreeSitterParser{
		Language: tree_sitter.NewLanguage(tree_sitter_js.Language()),
		IsLanguageFile: Extensions("js"),
	},
}

func Parse(content *string, extension string) (*ParseResult, error) {
	for _, parser := range parsers {
		if !parser.ShouldParseFile(extension) {
			continue
		}
		parseResult, err := parser.Parse(content)
		if err != nil {
			log.Panic("Parsing failed", err)
		}
		return parseResult, nil
	}
	return nil, errors.New("no supported parser found")
}


