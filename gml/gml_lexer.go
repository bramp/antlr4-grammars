// Code generated from gml.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gml

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 69, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 7, 4, 28, 10, 4, 12, 4, 14, 4, 31, 11, 4, 3, 4, 3, 4, 3, 5, 5, 5,
	36, 10, 5, 3, 5, 7, 5, 39, 10, 5, 12, 5, 14, 5, 42, 11, 5, 3, 5, 3, 5,
	6, 5, 46, 10, 5, 13, 5, 14, 5, 47, 3, 5, 5, 5, 51, 10, 5, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 6, 9, 62, 10, 9, 13, 9, 14, 9, 63,
	3, 10, 3, 10, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 3, 2, 7, 3, 2, 36, 36, 4, 2, 45, 45, 47, 47,
	3, 2, 50, 59, 5, 2, 50, 59, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34,
	34, 2, 74, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2,
	7, 25, 3, 2, 2, 2, 9, 35, 3, 2, 2, 2, 11, 52, 3, 2, 2, 2, 13, 54, 3, 2,
	2, 2, 15, 56, 3, 2, 2, 2, 17, 61, 3, 2, 2, 2, 19, 65, 3, 2, 2, 2, 21, 22,
	7, 93, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7, 95, 2, 2, 24, 6, 3, 2, 2, 2,
	25, 29, 7, 36, 2, 2, 26, 28, 10, 2, 2, 2, 27, 26, 3, 2, 2, 2, 28, 31, 3,
	2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 32, 3, 2, 2, 2, 31,
	29, 3, 2, 2, 2, 32, 33, 7, 36, 2, 2, 33, 8, 3, 2, 2, 2, 34, 36, 5, 11,
	6, 2, 35, 34, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 40, 3, 2, 2, 2, 37, 39,
	5, 13, 7, 2, 38, 37, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2,
	40, 41, 3, 2, 2, 2, 41, 43, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 45, 7,
	48, 2, 2, 44, 46, 5, 13, 7, 2, 45, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2,
	47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 51, 5,
	15, 8, 2, 50, 49, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 10, 3, 2, 2, 2, 52,
	53, 9, 3, 2, 2, 53, 12, 3, 2, 2, 2, 54, 55, 9, 4, 2, 2, 55, 14, 3, 2, 2,
	2, 56, 57, 7, 71, 2, 2, 57, 58, 5, 11, 6, 2, 58, 59, 5, 13, 7, 2, 59, 16,
	3, 2, 2, 2, 60, 62, 9, 5, 2, 2, 61, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2,
	63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 18, 3, 2, 2, 2, 65, 66, 9,
	6, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 8, 10, 2, 2, 68, 20, 3, 2, 2, 2, 9,
	2, 29, 35, 40, 47, 50, 63, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "STRINGLITERAL", "REAL", "SIGN", "DIGIT", "MANTISSA", "VALUE",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "STRINGLITERAL", "REAL", "SIGN", "DIGIT", "MANTISSA", "VALUE",
	"WS",
}

type gmlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewgmlLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *gmlLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewgmlLexer(input antlr.CharStream) *gmlLexer {
	l := new(gmlLexer)
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
	l.GrammarFileName = "gml.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gmlLexer tokens.
const (
	gmlLexerT__0          = 1
	gmlLexerT__1          = 2
	gmlLexerSTRINGLITERAL = 3
	gmlLexerREAL          = 4
	gmlLexerSIGN          = 5
	gmlLexerDIGIT         = 6
	gmlLexerMANTISSA      = 7
	gmlLexerVALUE         = 8
	gmlLexerWS            = 9
)
