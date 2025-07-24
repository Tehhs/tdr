// Code generated from pkg/tdrl/grammar/tdrl.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
    "", "'todo'", "", "','", "'('", "')'", "':'",
  }
  staticData.SymbolicNames = []string{
    "", "TODO", "TAG_ID", "COMMA", "LPARAN", "RPARAN", "COLON", "WS",
  }
  staticData.RuleNames = []string{
    "TODO", "TAG_ID", "COMMA", "LPARAN", "RPARAN", "COLON", "WS",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 7, 40, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 
	1, 22, 8, 1, 11, 1, 12, 1, 23, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 
	1, 5, 1, 6, 4, 6, 35, 8, 6, 11, 6, 12, 6, 36, 1, 6, 1, 6, 0, 0, 7, 1, 1, 
	3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 1, 0, 2, 4, 0, 48, 57, 65, 90, 95, 
	95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 41, 0, 1, 1, 0, 0, 0, 0, 3, 1, 
	0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 
	0, 0, 0, 0, 13, 1, 0, 0, 0, 1, 15, 1, 0, 0, 0, 3, 21, 1, 0, 0, 0, 5, 25, 
	1, 0, 0, 0, 7, 27, 1, 0, 0, 0, 9, 29, 1, 0, 0, 0, 11, 31, 1, 0, 0, 0, 13, 
	34, 1, 0, 0, 0, 15, 16, 5, 116, 0, 0, 16, 17, 5, 111, 0, 0, 17, 18, 5, 
	100, 0, 0, 18, 19, 5, 111, 0, 0, 19, 2, 1, 0, 0, 0, 20, 22, 7, 0, 0, 0, 
	21, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 
	0, 0, 0, 24, 4, 1, 0, 0, 0, 25, 26, 5, 44, 0, 0, 26, 6, 1, 0, 0, 0, 27, 
	28, 5, 40, 0, 0, 28, 8, 1, 0, 0, 0, 29, 30, 5, 41, 0, 0, 30, 10, 1, 0, 
	0, 0, 31, 32, 5, 58, 0, 0, 32, 12, 1, 0, 0, 0, 33, 35, 7, 1, 0, 0, 34, 
	33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 
	0, 37, 38, 1, 0, 0, 0, 38, 39, 6, 6, 0, 0, 39, 14, 1, 0, 0, 0, 3, 0, 23, 
	36, 1, 0, 1, 0,
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
	tdrlLexerTAG_ID = 2
	tdrlLexerCOMMA = 3
	tdrlLexerLPARAN = 4
	tdrlLexerRPARAN = 5
	tdrlLexerCOLON = 6
	tdrlLexerWS = 7
)

