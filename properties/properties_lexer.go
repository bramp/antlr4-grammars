// Code generated from properties.g4 by ANTLR 4.7.2. DO NOT EDIT.

package properties

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 45, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 6, 3, 17, 10, 3, 13, 3, 14, 3, 18, 3, 4, 3, 4, 3, 4, 3, 4, 7,
	4, 25, 10, 4, 12, 4, 14, 4, 28, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 34,
	10, 5, 12, 5, 14, 5, 37, 11, 5, 3, 6, 6, 6, 40, 10, 6, 13, 6, 14, 6, 41,
	3, 6, 3, 6, 2, 2, 7, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 3, 2, 5, 9, 2, 34,
	34, 39, 39, 46, 60, 66, 92, 97, 97, 99, 125, 127, 127, 3, 2, 36, 36, 4,
	2, 12, 12, 15, 15, 2, 49, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 16,
	3, 2, 2, 2, 7, 20, 3, 2, 2, 2, 9, 31, 3, 2, 2, 2, 11, 39, 3, 2, 2, 2, 13,
	14, 7, 63, 2, 2, 14, 4, 3, 2, 2, 2, 15, 17, 9, 2, 2, 2, 16, 15, 3, 2, 2,
	2, 17, 18, 3, 2, 2, 2, 18, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 6, 3,
	2, 2, 2, 20, 26, 7, 36, 2, 2, 21, 22, 7, 36, 2, 2, 22, 25, 7, 36, 2, 2,
	23, 25, 10, 3, 2, 2, 24, 21, 3, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 28, 3,
	2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 29, 3, 2, 2, 2, 28,
	26, 3, 2, 2, 2, 29, 30, 7, 36, 2, 2, 30, 8, 3, 2, 2, 2, 31, 35, 7, 37,
	2, 2, 32, 34, 10, 4, 2, 2, 33, 32, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2, 35,
	33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 10, 3, 2, 2, 2, 37, 35, 3, 2, 2,
	2, 38, 40, 9, 4, 2, 2, 39, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 39,
	3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 44, 8, 6, 2, 2,
	44, 12, 3, 2, 2, 2, 8, 2, 18, 24, 26, 35, 41, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "TEXT", "STRING", "COMMENT", "TERMINATOR",
}

var lexerRuleNames = []string{
	"T__0", "TEXT", "STRING", "COMMENT", "TERMINATOR",
}

type propertiesLexer struct {
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

func NewpropertiesLexer(input antlr.CharStream) *propertiesLexer {

	l := new(propertiesLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "properties.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// propertiesLexer tokens.
const (
	propertiesLexerT__0       = 1
	propertiesLexerTEXT       = 2
	propertiesLexerSTRING     = 3
	propertiesLexerCOMMENT    = 4
	propertiesLexerTERMINATOR = 5
)
