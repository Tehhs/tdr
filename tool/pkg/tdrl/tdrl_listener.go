// Code generated from tool/pkg/tdrl/grammar/tdrl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package tdrl // tdrl
import "github.com/antlr4-go/antlr/v4"


// tdrlListener is a complete listener for a parse tree produced by tdrlParser.
type tdrlListener interface {
	antlr.ParseTreeListener

	// EnterTodoComment is called when entering the todoComment production.
	EnterTodoComment(c *TodoCommentContext)

	// EnterTagBlock is called when entering the tagBlock production.
	EnterTagBlock(c *TagBlockContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// ExitTodoComment is called when exiting the todoComment production.
	ExitTodoComment(c *TodoCommentContext)

	// ExitTagBlock is called when exiting the tagBlock production.
	ExitTagBlock(c *TagBlockContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)
}
