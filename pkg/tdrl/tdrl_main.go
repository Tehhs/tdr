package tdrl

import (
	"fmt"
	"strings"

	antlr_v4 "github.com/antlr4-go/antlr/v4"
)

type TDRLTag struct {
}
type TDRLTodo struct {
	Tags             []string
	ProcessedContent string
}

func (t TDRLTodo) HasTag(checkTag string) bool {
	for _, tag := range t.Tags {
		if strings.ToLower(tag) == strings.ToLower(checkTag) {
			return true
		}
	}
	return false
}

type TDRListener struct {
	BasetdrlListener
	OnFindTodo *func(tags []string, content string)
}

func (s *TDRListener) EnterTodoRule(ctx *TodoRuleContext) {
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

type TDRLanguageParser struct {

}

func NewParser() TDRLanguageParser { 
	return TDRLanguageParser{}
}

func (p TDRLanguageParser) ProcessTodo(content string) []TDRLTodo {
	charStream := antlr_v4.NewInputStream(content)
	lexer := NewtdrlLexer(charStream)
	lexer.RemoveErrorListeners()

	tokens := antlr_v4.NewCommonTokenStream(lexer, antlr_v4.TokenDefaultChannel)
	parser := NewtdrlParser(tokens)
	parser.RemoveErrorListeners()

	tree := parser.Main()

	processedTodos := []TDRLTodo{}

	listener := &TDRListener{}
	onFindTodo := func(tags []string, content string) {
		processedTodos = append(processedTodos, TDRLTodo{
			Tags:             tags,
			ProcessedContent: content,
		})
	}
	listener.OnFindTodo = &onFindTodo
	antlr_v4.ParseTreeWalkerDefault.Walk(listener, tree)

	return processedTodos
}
