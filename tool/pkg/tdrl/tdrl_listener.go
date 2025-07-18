// Code generated from tool/pkg/tdrl/grammar/tdrl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package tdrl // tdrl
import "github.com/antlr4-go/antlr/v4"


// tdrlListener is a complete listener for a parse tree produced by tdrlParser.
type tdrlListener interface {
	antlr.ParseTreeListener

	// EnterTodo is called when entering the todo production.
	EnterTodo(c *TodoContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// ExitTodo is called when exiting the todo production.
	ExitTodo(c *TodoContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)
}
