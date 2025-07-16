package parser

import (
	"errors"
	"log"
)



type CommentBlock struct { 
	Lines []string
	StartLine int
	EndLine int 
}

func(b CommentBlock) String() string { 
	var output string = ""
	for _, l := range b.Lines { 
		output += l + "\n"
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
