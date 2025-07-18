// Code generated from tool/pkg/tdrl/grammar/tdrl.g4 by ANTLR 4.13.1. DO NOT EDIT.

package tdrl
import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type tdrlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var TdrlLexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func tdrllexerLexerInit() {
  staticData := &TdrlLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.LiteralNames = []string{
    "", "'todo'", "':'",
  }
  staticData.SymbolicNames = []string{
    "", "TODO", "COLON", "WS", "NL",
  }
  staticData.RuleNames = []string{
    "TODO", "COLON", "WS", "NL",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 4, 30, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 
	0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 4, 2, 18, 8, 2, 11, 2, 12, 
	2, 19, 1, 2, 1, 2, 1, 3, 3, 3, 25, 8, 3, 1, 3, 1, 3, 3, 3, 29, 8, 3, 0, 
	0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 1, 3, 0, 9, 10, 13, 13, 32, 32, 32, 
	0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 
	1, 9, 1, 0, 0, 0, 3, 14, 1, 0, 0, 0, 5, 17, 1, 0, 0, 0, 7, 28, 1, 0, 0, 
	0, 9, 10, 5, 116, 0, 0, 10, 11, 5, 111, 0, 0, 11, 12, 5, 100, 0, 0, 12, 
	13, 5, 111, 0, 0, 13, 2, 1, 0, 0, 0, 14, 15, 5, 58, 0, 0, 15, 4, 1, 0, 
	0, 0, 16, 18, 7, 0, 0, 0, 17, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 17, 
	1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22, 6, 2, 0, 0, 
	22, 6, 1, 0, 0, 0, 23, 25, 5, 13, 0, 0, 24, 23, 1, 0, 0, 0, 24, 25, 1, 
	0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 29, 5, 10, 0, 0, 27, 29, 5, 13, 0, 0, 
	28, 24, 1, 0, 0, 0, 28, 27, 1, 0, 0, 0, 29, 8, 1, 0, 0, 0, 4, 0, 19, 24, 
	28, 1, 6, 0, 0,
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

// tdrlLexerInit initializes any static state used to implement tdrlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewtdrlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TdrlLexerInit() {
  staticData := &TdrlLexerLexerStaticData
  staticData.once.Do(tdrllexerLexerInit)
}

// NewtdrlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewtdrlLexer(input antlr.CharStream) *tdrlLexer {
  TdrlLexerInit()
	l := new(tdrlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &TdrlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "tdrl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// tdrlLexer tokens.
const (
	tdrlLexerTODO = 1
	tdrlLexerCOLON = 2
	tdrlLexerWS = 3
	tdrlLexerNL = 4
)

