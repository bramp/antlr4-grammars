// Code generated from filter.g4 by ANTLR 4.9.3. DO NOT EDIT.

package filter

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 58, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 6, 13, 55, 10, 13, 13, 13, 14, 13,
	56, 2, 2, 14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 3, 2, 3, 7, 2, 34, 34, 48, 48, 50, 59, 67,
	92, 99, 124, 2, 58, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2, 7, 31,
	3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 35, 3, 2, 2, 2, 13, 37, 3, 2, 2, 2,
	15, 40, 3, 2, 2, 2, 17, 42, 3, 2, 2, 2, 19, 44, 3, 2, 2, 2, 21, 47, 3,
	2, 2, 2, 23, 50, 3, 2, 2, 2, 25, 54, 3, 2, 2, 2, 27, 28, 7, 42, 2, 2, 28,
	4, 3, 2, 2, 2, 29, 30, 7, 43, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 40, 2,
	2, 32, 8, 3, 2, 2, 2, 33, 34, 7, 126, 2, 2, 34, 10, 3, 2, 2, 2, 35, 36,
	7, 35, 2, 2, 36, 12, 3, 2, 2, 2, 37, 38, 7, 63, 2, 2, 38, 39, 7, 44, 2,
	2, 39, 14, 3, 2, 2, 2, 40, 41, 7, 44, 2, 2, 41, 16, 3, 2, 2, 2, 42, 43,
	7, 63, 2, 2, 43, 18, 3, 2, 2, 2, 44, 45, 7, 128, 2, 2, 45, 46, 7, 63, 2,
	2, 46, 20, 3, 2, 2, 2, 47, 48, 7, 64, 2, 2, 48, 49, 7, 63, 2, 2, 49, 22,
	3, 2, 2, 2, 50, 51, 7, 62, 2, 2, 51, 52, 7, 63, 2, 2, 52, 24, 3, 2, 2,
	2, 53, 55, 9, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 54,
	3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 26, 3, 2, 2, 2, 4, 2, 56, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'&'", "'|'", "'!'", "'=*'", "'*'", "'='", "'~='", "'>='",
	"'<='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "EQUAL", "APPROX", "GREATER", "LESS", "OCTETSTRING",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "EQUAL", "APPROX",
	"GREATER", "LESS", "OCTETSTRING",
}

type filterLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewfilterLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *filterLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewfilterLexer(input antlr.CharStream) *filterLexer {
	l := new(filterLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "filter.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// filterLexer tokens.
const (
	filterLexerT__0        = 1
	filterLexerT__1        = 2
	filterLexerT__2        = 3
	filterLexerT__3        = 4
	filterLexerT__4        = 5
	filterLexerT__5        = 6
	filterLexerT__6        = 7
	filterLexerEQUAL       = 8
	filterLexerAPPROX      = 9
	filterLexerGREATER     = 10
	filterLexerLESS        = 11
	filterLexerOCTETSTRING = 12
)
