// Generated from postalcode.g4 by ANTLR 4.7.

package postalcode

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 5, 17, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 2, 2, 5, 3, 3, 5, 4, 7, 5, 3, 2, 3, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 16, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 3,
	9, 3, 2, 2, 2, 5, 11, 3, 2, 2, 2, 7, 13, 3, 2, 2, 2, 9, 10, 4, 50, 59,
	2, 10, 4, 3, 2, 2, 2, 11, 12, 4, 67, 92, 2, 12, 6, 3, 2, 2, 2, 13, 14,
	9, 2, 2, 2, 14, 15, 3, 2, 2, 2, 15, 16, 8, 4, 2, 2, 16, 8, 3, 2, 2, 2,
	3, 2, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames []string

var lexerSymbolicNames = []string{
	"", "DIGIT", "LETTER", "WS",
}

var lexerRuleNames = []string{
	"DIGIT", "LETTER", "WS",
}

type postalcodeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewpostalcodeLexer(input antlr.CharStream) *postalcodeLexer {

	l := new(postalcodeLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "postalcode.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// postalcodeLexer tokens.
const (
	postalcodeLexerDIGIT  = 1
	postalcodeLexerLETTER = 2
	postalcodeLexerWS     = 3
)
