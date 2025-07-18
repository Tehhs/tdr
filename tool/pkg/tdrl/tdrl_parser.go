// Code generated from tool/pkg/tdrl/grammar/tdrl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package tdrl // tdrl
import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type tdrlParser struct {
	*antlr.BaseParser
}

var TdrlParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func tdrlParserInit() {
  staticData := &TdrlParserStaticData
  staticData.LiteralNames = []string{
    "", "'#'", "'todo'", "'('", "')'", "','", "':'", "'='",
  }
  staticData.SymbolicNames = []string{
    "", "HASH", "TODO", "LPAREN", "RPAREN", "COMMA", "COLON", "EQUALS", 
    "IDENTIFIER", "TEXT_LINE", "NEWLINE",
  }
  staticData.RuleNames = []string{
    "todoComment", "tagBlock", "tag", "message",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 10, 40, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 3, 
	0, 10, 8, 0, 1, 0, 1, 0, 3, 0, 14, 8, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 
	1, 1, 1, 1, 5, 1, 23, 8, 1, 10, 1, 12, 1, 26, 9, 1, 1, 1, 1, 1, 1, 2, 1, 
	2, 1, 2, 3, 2, 33, 8, 2, 1, 3, 4, 3, 36, 8, 3, 11, 3, 12, 3, 37, 1, 3, 
	0, 0, 4, 0, 2, 4, 6, 0, 0, 40, 0, 9, 1, 0, 0, 0, 2, 18, 1, 0, 0, 0, 4, 
	29, 1, 0, 0, 0, 6, 35, 1, 0, 0, 0, 8, 10, 5, 1, 0, 0, 9, 8, 1, 0, 0, 0, 
	9, 10, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11, 13, 5, 2, 0, 0, 12, 14, 3, 2, 
	1, 0, 13, 12, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 16, 
	5, 6, 0, 0, 16, 17, 3, 6, 3, 0, 17, 1, 1, 0, 0, 0, 18, 19, 5, 3, 0, 0, 
	19, 24, 3, 4, 2, 0, 20, 21, 5, 5, 0, 0, 21, 23, 3, 4, 2, 0, 22, 20, 1, 
	0, 0, 0, 23, 26, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 
	27, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 27, 28, 5, 4, 0, 0, 28, 3, 1, 0, 0, 
	0, 29, 32, 5, 8, 0, 0, 30, 31, 5, 7, 0, 0, 31, 33, 5, 8, 0, 0, 32, 30, 
	1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 5, 1, 0, 0, 0, 34, 36, 5, 9, 0, 0, 
	35, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 
	0, 0, 0, 38, 7, 1, 0, 0, 0, 5, 9, 13, 24, 32, 37,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// tdrlParserInit initializes any static state used to implement tdrlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewtdrlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TdrlParserInit() {
  staticData := &TdrlParserStaticData
  staticData.once.Do(tdrlParserInit)
}

// NewtdrlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewtdrlParser(input antlr.TokenStream) *tdrlParser {
	TdrlParserInit()
	this := new(tdrlParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &TdrlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "tdrl.g4"

	return this
}


// tdrlParser tokens.
const (
	tdrlParserEOF = antlr.TokenEOF
	tdrlParserHASH = 1
	tdrlParserTODO = 2
	tdrlParserLPAREN = 3
	tdrlParserRPAREN = 4
	tdrlParserCOMMA = 5
	tdrlParserCOLON = 6
	tdrlParserEQUALS = 7
	tdrlParserIDENTIFIER = 8
	tdrlParserTEXT_LINE = 9
	tdrlParserNEWLINE = 10
)

// tdrlParser rules.
const (
	tdrlParserRULE_todoComment = 0
	tdrlParserRULE_tagBlock = 1
	tdrlParserRULE_tag = 2
	tdrlParserRULE_message = 3
)

// ITodoCommentContext is an interface to support dynamic dispatch.
type ITodoCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TODO() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Message() IMessageContext
	HASH() antlr.TerminalNode
	TagBlock() ITagBlockContext

	// IsTodoCommentContext differentiates from other interfaces.
	IsTodoCommentContext()
}

type TodoCommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTodoCommentContext() *TodoCommentContext {
	var p = new(TodoCommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_todoComment
	return p
}

func InitEmptyTodoCommentContext(p *TodoCommentContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_todoComment
}

func (*TodoCommentContext) IsTodoCommentContext() {}

func NewTodoCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TodoCommentContext {
	var p = new(TodoCommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tdrlParserRULE_todoComment

	return p
}

func (s *TodoCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *TodoCommentContext) TODO() antlr.TerminalNode {
	return s.GetToken(tdrlParserTODO, 0)
}

func (s *TodoCommentContext) COLON() antlr.TerminalNode {
	return s.GetToken(tdrlParserCOLON, 0)
}

func (s *TodoCommentContext) Message() IMessageContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *TodoCommentContext) HASH() antlr.TerminalNode {
	return s.GetToken(tdrlParserHASH, 0)
}

func (s *TodoCommentContext) TagBlock() ITagBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagBlockContext)
}

func (s *TodoCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TodoCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TodoCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.EnterTodoComment(s)
	}
}

func (s *TodoCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.ExitTodoComment(s)
	}
}




func (p *tdrlParser) TodoComment() (localctx ITodoCommentContext) {
	localctx = NewTodoCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tdrlParserRULE_todoComment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == tdrlParserHASH {
		{
			p.SetState(8)
			p.Match(tdrlParserHASH)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(11)
		p.Match(tdrlParserTODO)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == tdrlParserLPAREN {
		{
			p.SetState(12)
			p.TagBlock()
		}

	}
	{
		p.SetState(15)
		p.Match(tdrlParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(16)
		p.Message()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITagBlockContext is an interface to support dynamic dispatch.
type ITagBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllTag() []ITagContext
	Tag(i int) ITagContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsTagBlockContext differentiates from other interfaces.
	IsTagBlockContext()
}

type TagBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagBlockContext() *TagBlockContext {
	var p = new(TagBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_tagBlock
	return p
}

func InitEmptyTagBlockContext(p *TagBlockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_tagBlock
}

func (*TagBlockContext) IsTagBlockContext() {}

func NewTagBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagBlockContext {
	var p = new(TagBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tdrlParserRULE_tagBlock

	return p
}

func (s *TagBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *TagBlockContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(tdrlParserLPAREN, 0)
}

func (s *TagBlockContext) AllTag() []ITagContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITagContext); ok {
			len++
		}
	}

	tst := make([]ITagContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITagContext); ok {
			tst[i] = t.(ITagContext)
			i++
		}
	}

	return tst
}

func (s *TagBlockContext) Tag(i int) ITagContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *TagBlockContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(tdrlParserRPAREN, 0)
}

func (s *TagBlockContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tdrlParserCOMMA)
}

func (s *TagBlockContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tdrlParserCOMMA, i)
}

func (s *TagBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TagBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.EnterTagBlock(s)
	}
}

func (s *TagBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.ExitTagBlock(s)
	}
}




func (p *tdrlParser) TagBlock() (localctx ITagBlockContext) {
	localctx = NewTagBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tdrlParserRULE_tagBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Match(tdrlParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(19)
		p.Tag()
	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == tdrlParserCOMMA {
		{
			p.SetState(20)
			p.Match(tdrlParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(21)
			p.Tag()
		}


		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(27)
		p.Match(tdrlParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	EQUALS() antlr.TerminalNode

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_tag
	return p
}

func InitEmptyTagContext(p *TagContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_tag
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tdrlParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(tdrlParserIDENTIFIER)
}

func (s *TagContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(tdrlParserIDENTIFIER, i)
}

func (s *TagContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(tdrlParserEQUALS, 0)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.ExitTag(s)
	}
}




func (p *tdrlParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tdrlParserRULE_tag)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(tdrlParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == tdrlParserEQUALS {
		{
			p.SetState(30)
			p.Match(tdrlParserEQUALS)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(31)
			p.Match(tdrlParserIDENTIFIER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTEXT_LINE() []antlr.TerminalNode
	TEXT_LINE(i int) antlr.TerminalNode

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_message
	return p
}

func InitEmptyMessageContext(p *MessageContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_message
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tdrlParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) AllTEXT_LINE() []antlr.TerminalNode {
	return s.GetTokens(tdrlParserTEXT_LINE)
}

func (s *MessageContext) TEXT_LINE(i int) antlr.TerminalNode {
	return s.GetToken(tdrlParserTEXT_LINE, i)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.ExitMessage(s)
	}
}




func (p *tdrlParser) Message() (localctx IMessageContext) {
	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tdrlParserRULE_message)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == tdrlParserTEXT_LINE {
		{
			p.SetState(34)
			p.Match(tdrlParserTEXT_LINE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


