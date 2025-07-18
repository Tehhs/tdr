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
    "", "'todo'", "':'",
  }
  staticData.SymbolicNames = []string{
    "", "TODO", "COLON", "WS", "NL",
  }
  staticData.RuleNames = []string{
    "todo", "message",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 4, 16, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 
	5, 1, 11, 8, 1, 10, 1, 12, 1, 14, 9, 1, 1, 1, 0, 0, 2, 0, 2, 0, 1, 1, 0, 
	4, 4, 14, 0, 4, 1, 0, 0, 0, 2, 12, 1, 0, 0, 0, 4, 5, 5, 1, 0, 0, 5, 6, 
	5, 2, 0, 0, 6, 7, 3, 2, 1, 0, 7, 8, 5, 0, 0, 1, 8, 1, 1, 0, 0, 0, 9, 11, 
	8, 0, 0, 0, 10, 9, 1, 0, 0, 0, 11, 14, 1, 0, 0, 0, 12, 10, 1, 0, 0, 0, 
	12, 13, 1, 0, 0, 0, 13, 3, 1, 0, 0, 0, 14, 12, 1, 0, 0, 0, 1, 12,
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
	tdrlParserCOLON = 2
	tdrlParserWS = 3
	tdrlParserNL = 4
)

// tdrlParser rules.
const (
	tdrlParserRULE_todo = 0
	tdrlParserRULE_message = 1
)

// ITodoContext is an interface to support dynamic dispatch.
type ITodoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TODO() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Message() IMessageContext
	EOF() antlr.TerminalNode

	// IsTodoContext differentiates from other interfaces.
	IsTodoContext()
}

type TodoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTodoContext() *TodoContext {
	var p = new(TodoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_todo
	return p
}

func InitEmptyTodoContext(p *TodoContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tdrlParserRULE_todo
}

func (*TodoContext) IsTodoContext() {}

func NewTodoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TodoContext {
	var p = new(TodoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tdrlParserRULE_todo

	return p
}

func (s *TodoContext) GetParser() antlr.Parser { return s.parser }

func (s *TodoContext) TODO() antlr.TerminalNode {
	return s.GetToken(tdrlParserTODO, 0)
}

func (s *TodoContext) COLON() antlr.TerminalNode {
	return s.GetToken(tdrlParserCOLON, 0)
}

func (s *TodoContext) Message() IMessageContext {
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

func (s *TodoContext) EOF() antlr.TerminalNode {
	return s.GetToken(tdrlParserEOF, 0)
}

func (s *TodoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TodoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TodoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.EnterTodo(s)
	}
}

func (s *TodoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tdrlListener); ok {
		listenerT.ExitTodo(s)
	}
}




func (p *tdrlParser) Todo() (localctx ITodoContext) {
	localctx = NewTodoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tdrlParserRULE_todo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Match(tdrlParserTODO)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(5)
		p.Match(tdrlParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(6)
		p.Message()
	}
	{
		p.SetState(7)
		p.Match(tdrlParserEOF)
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


// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

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

func (s *MessageContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(tdrlParserNL)
}

func (s *MessageContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(tdrlParserNL, i)
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
	p.EnterRule(localctx, 2, tdrlParserRULE_message)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 14) != 0) {
		{
			p.SetState(9)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == tdrlParserNL  {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


		p.SetState(14)
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


