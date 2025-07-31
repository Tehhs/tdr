package comments

import (
	"errors"
	"log"
)

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

var parsers []Parser = []Parser{
	GoParser{},
	JavascriptParser{},
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
