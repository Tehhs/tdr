// Code generated from pkg/tdrl/grammar/tdrl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package tdrl // tdrl
import "github.com/antlr4-go/antlr/v4"


// tdrlListener is a complete listener for a parse tree produced by tdrlParser.
type tdrlListener interface {
	antlr.ParseTreeListener

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterTodoRule is called when entering the todoRule production.
	EnterTodoRule(c *TodoRuleContext)

	// EnterTagRule is called when entering the tagRule production.
	EnterTagRule(c *TagRuleContext)

	// EnterMessageContent is called when entering the messageContent production.
	EnterMessageContent(c *MessageContentContext)

	// EnterTagList is called when entering the tagList production.
	EnterTagList(c *TagListContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitTodoRule is called when exiting the todoRule production.
	ExitTodoRule(c *TodoRuleContext)

	// ExitTagRule is called when exiting the tagRule production.
	ExitTagRule(c *TagRuleContext)

	// ExitMessageContent is called when exiting the messageContent production.
	ExitMessageContent(c *MessageContentContext)

	// ExitTagList is called when exiting the tagList production.
	ExitTagList(c *TagListContext)
}
