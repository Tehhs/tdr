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
    "", "'#'", "'todo'", "'('", "')'", "','", "':'", "'='",
  }
  staticData.SymbolicNames = []string{
    "", "HASH", "TODO", "LPAREN", "RPAREN", "COMMA", "COLON", "EQUALS", 
    "IDENTIFIER", "TEXT_LINE", "NEWLINE",
  }
  staticData.RuleNames = []string{
    "HASH", "TODO", "LPAREN", "RPAREN", "COMMA", "COLON", "EQUALS", "IDENTIFIER", 
    "TEXT_LINE", "NEWLINE",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 10, 57, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 
	0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 
	4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 5, 7, 41, 8, 7, 10, 7, 12, 7, 44, 
	9, 7, 1, 8, 4, 8, 47, 8, 8, 11, 8, 12, 8, 48, 1, 8, 1, 8, 1, 9, 4, 9, 54, 
	8, 9, 11, 9, 12, 9, 55, 0, 0, 10, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 
	13, 7, 15, 8, 17, 9, 19, 10, 1, 0, 3, 3, 0, 65, 90, 95, 95, 97, 122, 5, 
	0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 59, 0, 
	1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 
	9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 
	0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 21, 1, 0, 0, 0, 3, 23, 1, 0, 0, 
	0, 5, 28, 1, 0, 0, 0, 7, 30, 1, 0, 0, 0, 9, 32, 1, 0, 0, 0, 11, 34, 1, 
	0, 0, 0, 13, 36, 1, 0, 0, 0, 15, 38, 1, 0, 0, 0, 17, 46, 1, 0, 0, 0, 19, 
	53, 1, 0, 0, 0, 21, 22, 5, 35, 0, 0, 22, 2, 1, 0, 0, 0, 23, 24, 5, 116, 
	0, 0, 24, 25, 5, 111, 0, 0, 25, 26, 5, 100, 0, 0, 26, 27, 5, 111, 0, 0, 
	27, 4, 1, 0, 0, 0, 28, 29, 5, 40, 0, 0, 29, 6, 1, 0, 0, 0, 30, 31, 5, 41, 
	0, 0, 31, 8, 1, 0, 0, 0, 32, 33, 5, 44, 0, 0, 33, 10, 1, 0, 0, 0, 34, 35, 
	5, 58, 0, 0, 35, 12, 1, 0, 0, 0, 36, 37, 5, 61, 0, 0, 37, 14, 1, 0, 0, 
	0, 38, 42, 7, 0, 0, 0, 39, 41, 7, 1, 0, 0, 40, 39, 1, 0, 0, 0, 41, 44, 
	1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 16, 1, 0, 0, 0, 
	44, 42, 1, 0, 0, 0, 45, 47, 8, 2, 0, 0, 46, 45, 1, 0, 0, 0, 47, 48, 1, 
	0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 
	51, 6, 8, 0, 0, 51, 18, 1, 0, 0, 0, 52, 54, 7, 2, 0, 0, 53, 52, 1, 0, 0, 
	0, 54, 55, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 20, 
	1, 0, 0, 0, 4, 0, 42, 48, 55, 1, 6, 0, 0,
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
	tdrlLexerHASH = 1
	tdrlLexerTODO = 2
	tdrlLexerLPAREN = 3
	tdrlLexerRPAREN = 4
	tdrlLexerCOMMA = 5
	tdrlLexerCOLON = 6
	tdrlLexerEQUALS = 7
	tdrlLexerIDENTIFIER = 8
	tdrlLexerTEXT_LINE = 9
	tdrlLexerNEWLINE = 10
)

