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
    "", "'todo'", "", "','", "'('", "')'", "':'",
  }
  staticData.SymbolicNames = []string{
    "", "TODO", "TAG_ID", "COMMA", "LPARAN", "RPARAN", "COLON", "WS",
  }
  staticData.RuleNames = []string{
    "main", "todoRule", "tagRule", "messageContent", "tagList",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 7, 39, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4, 
	1, 0, 3, 0, 12, 8, 0, 1, 1, 1, 1, 3, 1, 16, 8, 1, 1, 1, 1, 1, 1, 1, 1, 
	2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 4, 3, 27, 8, 3, 11, 3, 12, 3, 28, 1, 4, 
	1, 4, 1, 4, 5, 4, 34, 8, 4, 10, 4, 12, 4, 37, 9, 4, 1, 4, 1, 28, 0, 5, 
	0, 2, 4, 6, 8, 0, 0, 38, 0, 11, 1, 0, 0, 0, 2, 13, 1, 0, 0, 0, 4, 20, 1, 
	0, 0, 0, 6, 26, 1, 0, 0, 0, 8, 30, 1, 0, 0, 0, 10, 12, 3, 2, 1, 0, 11, 
	10, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 12, 1, 1, 0, 0, 0, 13, 15, 5, 1, 0, 
	0, 14, 16, 3, 4, 2, 0, 15, 14, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 17, 
	1, 0, 0, 0, 17, 18, 5, 6, 0, 0, 18, 19, 3, 6, 3, 0, 19, 3, 1, 0, 0, 0, 
	20, 21, 5, 4, 0, 0, 21, 22, 3, 8, 4, 0, 22, 23, 5, 5, 0, 0, 23, 5, 1, 0, 
	0, 0, 24, 27, 5, 7, 0, 0, 25, 27, 9, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 25, 
	1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 
	29, 7, 1, 0, 0, 0, 30, 35, 5, 2, 0, 0, 31, 32, 5, 3, 0, 0, 32, 34, 5, 2, 
	0, 0, 33, 31, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 
	1, 0, 0, 0, 36, 9, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 5, 11, 15, 26, 28, 35,
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
	tdrlParserTODO = 1
	tdrlParserTAG_ID = 2
	tdrlParserCOMMA = 3
	tdrlParserLPARAN = 4
	tdrlParserRPARAN = 5
	tdrlParserCOLON = 6
	tdrlParserWS = 7
)

// tdrlParser rules.
const (
	tdrlParserRULE_main = 0
	tdrlParserRULE_todoRule = 1
	tdrlParserRULE_tagRule = 2
	tdrlParserRULE_messageContent = 3
	tdrlParserRULE_tagList = 4
)

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TodoRule() ITodoRuleContext

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_main
	return p
}

func InitEmptyMainContext(p *MainContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_main
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tdrlParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) TodoRule() ITodoRuleContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITodoRuleContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITodoRuleContext)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.ExitMain(s)
	}
}




func (p *tdrlParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tdrlParserRULE_main)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == tdrlParserTODO {
		{
			p.SetState(10)
			p.TodoRule()
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


// ITodoRuleContext is an interface to support dynamic dispatch.
type ITodoRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TODO() antlr.TerminalNode
	COLON() antlr.TerminalNode
	MessageContent() IMessageContentContext
	TagRule() ITagRuleContext

	// IsTodoRuleContext differentiates from other interfaces.
	IsTodoRuleContext()
}

type TodoRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTodoRuleContext() *TodoRuleContext {
	var p = new(TodoRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_todoRule
	return p
}

func InitEmptyTodoRuleContext(p *TodoRuleContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_todoRule
}

func (*TodoRuleContext) IsTodoRuleContext() {}

func NewTodoRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TodoRuleContext {
	var p = new(TodoRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tdrlParserRULE_todoRule

	return p
}

func (s *TodoRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TodoRuleContext) TODO() antlr.TerminalNode {
	return s.GetToken(tdrlParserTODO, 0)
}

func (s *TodoRuleContext) COLON() antlr.TerminalNode {
	return s.GetToken(tdrlParserCOLON, 0)
}

func (s *TodoRuleContext) MessageContent() IMessageContentContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContentContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContentContext)
}

func (s *TodoRuleContext) TagRule() ITagRuleContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagRuleContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagRuleContext)
}

func (s *TodoRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TodoRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TodoRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.EnterTodoRule(s)
	}
}

func (s *TodoRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.ExitTodoRule(s)
	}
}




func (p *tdrlParser) TodoRule() (localctx ITodoRuleContext) {
	localctx = NewTodoRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tdrlParserRULE_todoRule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)
		p.Match(tdrlParserTODO)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == tdrlParserLPARAN {
		{
			p.SetState(14)
			p.TagRule()
		}

	}
	{
		p.SetState(17)
		p.Match(tdrlParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(18)
		p.MessageContent()
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


// ITagRuleContext is an interface to support dynamic dispatch.
type ITagRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPARAN() antlr.TerminalNode
	TagList() ITagListContext
	RPARAN() antlr.TerminalNode

	// IsTagRuleContext differentiates from other interfaces.
	IsTagRuleContext()
}

type TagRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagRuleContext() *TagRuleContext {
	var p = new(TagRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_tagRule
	return p
}

func InitEmptyTagRuleContext(p *TagRuleContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_tagRule
}

func (*TagRuleContext) IsTagRuleContext() {}

func NewTagRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagRuleContext {
	var p = new(TagRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tdrlParserRULE_tagRule

	return p
}

func (s *TagRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TagRuleContext) LPARAN() antlr.TerminalNode {
	return s.GetToken(tdrlParserLPARAN, 0)
}

func (s *TagRuleContext) TagList() ITagListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITagListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITagListContext)
}

func (s *TagRuleContext) RPARAN() antlr.TerminalNode {
	return s.GetToken(tdrlParserRPARAN, 0)
}

func (s *TagRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TagRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.EnterTagRule(s)
	}
}

func (s *TagRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.ExitTagRule(s)
	}
}




func (p *tdrlParser) TagRule() (localctx ITagRuleContext) {
	localctx = NewTagRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tdrlParserRULE_tagRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Match(tdrlParserLPARAN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(21)
		p.TagList()
	}
	{
		p.SetState(22)
		p.Match(tdrlParserRPARAN)
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


// IMessageContentContext is an interface to support dynamic dispatch.
type IMessageContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsMessageContentContext differentiates from other interfaces.
	IsMessageContentContext()
}

type MessageContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContentContext() *MessageContentContext {
	var p = new(MessageContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_messageContent
	return p
}

func InitEmptyMessageContentContext(p *MessageContentContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_messageContent
}

func (*MessageContentContext) IsMessageContentContext() {}

func NewMessageContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContentContext {
	var p = new(MessageContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tdrlParserRULE_messageContent

	return p
}

func (s *MessageContentContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContentContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(tdrlParserWS)
}

func (s *MessageContentContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(tdrlParserWS, i)
}

func (s *MessageContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MessageContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.EnterMessageContent(s)
	}
}

func (s *MessageContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.ExitMessageContent(s)
	}
}




func (p *tdrlParser) MessageContent() (localctx IMessageContentContext) {
	localctx = NewMessageContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tdrlParserRULE_messageContent)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1+1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1+1:
				p.SetState(26)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(24)
						p.Match(tdrlParserWS)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}


				case 2:
					p.SetState(25)
					p.MatchWildcard()


				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
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


// ITagListContext is an interface to support dynamic dispatch.
type ITagListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTAG_ID() []antlr.TerminalNode
	TAG_ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsTagListContext differentiates from other interfaces.
	IsTagListContext()
}

type TagListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagListContext() *TagListContext {
	var p = new(TagListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_tagList
	return p
}

func InitEmptyTagListContext(p *TagListContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_tagList
}

func (*TagListContext) IsTagListContext() {}

func NewTagListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagListContext {
	var p = new(TagListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tdrlParserRULE_tagList

	return p
}

func (s *TagListContext) GetParser() antlr.Parser { return s.parser }

func (s *TagListContext) AllTAG_ID() []antlr.TerminalNode {
	return s.GetTokens(tdrlParserTAG_ID)
}

func (s *TagListContext) TAG_ID(i int) antlr.TerminalNode {
	return s.GetToken(tdrlParserTAG_ID, i)
}

func (s *TagListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tdrlParserCOMMA)
}

func (s *TagListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tdrlParserCOMMA, i)
}

func (s *TagListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TagListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.EnterTagList(s)
	}
}

func (s *TagListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.ExitTagList(s)
	}
}




func (p *tdrlParser) TagList() (localctx ITagListContext) {
	localctx = NewTagListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tdrlParserRULE_tagList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(tdrlParserTAG_ID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == tdrlParserCOMMA {
		{
			p.SetState(31)
			p.Match(tdrlParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(32)
			p.Match(tdrlParserTAG_ID)
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


