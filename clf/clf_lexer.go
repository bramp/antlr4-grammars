// Code generated from clf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package clf

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 113,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 6, 5, 33, 10, 5, 13, 5, 14, 5,
	34, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 41, 10, 5, 13, 5, 14, 5, 42, 3, 6, 6,
	6, 46, 10, 6, 13, 6, 14, 6, 47, 3, 6, 3, 6, 6, 6, 52, 10, 6, 13, 6, 14,
	6, 53, 3, 6, 3, 6, 6, 6, 58, 10, 6, 13, 6, 14, 6, 59, 3, 7, 3, 7, 6, 7,
	64, 10, 7, 13, 7, 14, 7, 65, 3, 8, 3, 8, 7, 8, 70, 10, 8, 12, 8, 14, 8,
	73, 11, 8, 3, 8, 3, 8, 3, 9, 6, 9, 78, 10, 9, 13, 9, 14, 9, 79, 3, 9, 3,
	9, 6, 9, 84, 10, 9, 13, 9, 14, 9, 85, 3, 9, 3, 9, 6, 9, 90, 10, 9, 13,
	9, 14, 9, 91, 3, 9, 3, 9, 6, 9, 96, 10, 9, 13, 9, 14, 9, 97, 3, 10, 6,
	10, 101, 10, 10, 13, 10, 14, 10, 102, 3, 11, 5, 11, 106, 10, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 3, 2, 6, 3, 2, 50, 59,
	3, 2, 36, 36, 9, 2, 42, 43, 47, 48, 50, 59, 61, 61, 67, 92, 97, 97, 99,
	124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 125, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 27, 3, 2, 2, 2, 7,
	29, 3, 2, 2, 2, 9, 32, 3, 2, 2, 2, 11, 45, 3, 2, 2, 2, 13, 61, 3, 2, 2,
	2, 15, 67, 3, 2, 2, 2, 17, 77, 3, 2, 2, 2, 19, 100, 3, 2, 2, 2, 21, 105,
	3, 2, 2, 2, 23, 109, 3, 2, 2, 2, 25, 26, 7, 93, 2, 2, 26, 4, 3, 2, 2, 2,
	27, 28, 7, 60, 2, 2, 28, 6, 3, 2, 2, 2, 29, 30, 7, 95, 2, 2, 30, 8, 3,
	2, 2, 2, 31, 33, 9, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34,
	32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 37, 7, 49,
	2, 2, 37, 38, 5, 19, 10, 2, 38, 40, 7, 49, 2, 2, 39, 41, 9, 2, 2, 2, 40,
	39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2,
	2, 43, 10, 3, 2, 2, 2, 44, 46, 9, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 47,
	3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2,
	49, 51, 7, 60, 2, 2, 50, 52, 9, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 53, 3,
	2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55,
	57, 7, 60, 2, 2, 56, 58, 9, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2,
	2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 12, 3, 2, 2, 2, 61, 63,
	7, 47, 2, 2, 62, 64, 9, 2, 2, 2, 63, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2,
	65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 14, 3, 2, 2, 2, 67, 71, 7,
	36, 2, 2, 68, 70, 10, 3, 2, 2, 69, 68, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2,
	71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 74, 3, 2, 2, 2, 73, 71, 3,
	2, 2, 2, 74, 75, 7, 36, 2, 2, 75, 16, 3, 2, 2, 2, 76, 78, 9, 2, 2, 2, 77,
	76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2,
	2, 80, 81, 3, 2, 2, 2, 81, 83, 7, 48, 2, 2, 82, 84, 9, 2, 2, 2, 83, 82,
	3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2,
	86, 87, 3, 2, 2, 2, 87, 89, 7, 48, 2, 2, 88, 90, 9, 2, 2, 2, 89, 88, 3,
	2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92,
	93, 3, 2, 2, 2, 93, 95, 7, 48, 2, 2, 94, 96, 9, 2, 2, 2, 95, 94, 3, 2,
	2, 2, 96, 97, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 18,
	3, 2, 2, 2, 99, 101, 9, 4, 2, 2, 100, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2,
	2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 20, 3, 2, 2, 2, 104,
	106, 7, 15, 2, 2, 105, 104, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107,
	3, 2, 2, 2, 107, 108, 7, 12, 2, 2, 108, 22, 3, 2, 2, 2, 109, 110, 9, 5,
	2, 2, 110, 111, 3, 2, 2, 2, 111, 112, 8, 12, 2, 2, 112, 24, 3, 2, 2, 2,
	16, 2, 34, 42, 47, 53, 59, 65, 71, 79, 85, 91, 97, 102, 105, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "':'", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "DATE", "TIME", "TZ", "LITERAL", "IP", "STRING", "EOL",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "DATE", "TIME", "TZ", "LITERAL", "IP", "STRING",
	"EOL", "WS",
}

type clfLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewclfLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *clfLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewclfLexer(input antlr.CharStream) *clfLexer {
	l := new(clfLexer)
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
	l.GrammarFileName = "clf.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// clfLexer tokens.
const (
	clfLexerT__0    = 1
	clfLexerT__1    = 2
	clfLexerT__2    = 3
	clfLexerDATE    = 4
	clfLexerTIME    = 5
	clfLexerTZ      = 6
	clfLexerLITERAL = 7
	clfLexerIP      = 8
	clfLexerSTRING  = 9
	clfLexerEOL     = 10
	clfLexerWS      = 11
)
