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

// EnterMain is called when production main is entered.
func (s *BasetdrlListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BasetdrlListener) ExitMain(ctx *MainContext) {}

// EnterTodoRule is called when production todoRule is entered.
func (s *BasetdrlListener) EnterTodoRule(ctx *TodoRuleContext) {}

// ExitTodoRule is called when production todoRule is exited.
func (s *BasetdrlListener) ExitTodoRule(ctx *TodoRuleContext) {}

// EnterTagRule is called when production tagRule is entered.
func (s *BasetdrlListener) EnterTagRule(ctx *TagRuleContext) {}

// ExitTagRule is called when production tagRule is exited.
func (s *BasetdrlListener) ExitTagRule(ctx *TagRuleContext) {}

// EnterMessageContent is called when production messageContent is entered.
func (s *BasetdrlListener) EnterMessageContent(ctx *MessageContentContext) {}

// ExitMessageContent is called when production messageContent is exited.
func (s *BasetdrlListener) ExitMessageContent(ctx *MessageContentContext) {}

// EnterTagList is called when production tagList is entered.
func (s *BasetdrlListener) EnterTagList(ctx *TagListContext) {}

// ExitTagList is called when production tagList is exited.
func (s *BasetdrlListener) ExitTagList(ctx *TagListContext) {}
