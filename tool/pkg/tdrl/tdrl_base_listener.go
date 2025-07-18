// Code generated from tool/pkg/tdrl/grammar/tdrl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package tdrl // tdrl
import "github.com/antlr4-go/antlr/v4"

// BasetdrlListener is a complete listener for a parse tree produced by tdrlParser.
type BasetdrlListener struct{}

var _ tdrlListener = &BasetdrlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetdrlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetdrlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetdrlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetdrlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTodoComment is called when production todoComment is entered.
func (s *BasetdrlListener) EnterTodoComment(ctx *TodoCommentContext) {}

// ExitTodoComment is called when production todoComment is exited.
func (s *BasetdrlListener) ExitTodoComment(ctx *TodoCommentContext) {}

// EnterTagBlock is called when production tagBlock is entered.
func (s *BasetdrlListener) EnterTagBlock(ctx *TagBlockContext) {}

// ExitTagBlock is called when production tagBlock is exited.
func (s *BasetdrlListener) ExitTagBlock(ctx *TagBlockContext) {}

// EnterTag is called when production tag is entered.
func (s *BasetdrlListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BasetdrlListener) ExitTag(ctx *TagContext) {}

// EnterMessage is called when production message is entered.
func (s *BasetdrlListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BasetdrlListener) ExitMessage(ctx *MessageContext) {}
