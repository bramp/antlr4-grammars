// Code generated from databank.g4 by ANTLR 4.9.3. DO NOT EDIT.

package databank

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 94, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 5, 6, 42, 10, 6, 3, 6, 3, 6, 5, 6, 46, 10, 6, 3, 7,
	6, 7, 49, 10, 7, 13, 7, 14, 7, 50, 3, 7, 3, 7, 7, 7, 55, 10, 7, 12, 7,
	14, 7, 58, 11, 7, 5, 7, 60, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 7, 10, 70, 10, 10, 12, 10, 14, 10, 73, 11, 10, 3, 10, 3,
	10, 7, 10, 77, 10, 10, 12, 10, 14, 10, 80, 11, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 6, 11, 87, 10, 11, 13, 11, 14, 11, 88, 3, 12, 3, 12, 3, 12,
	3, 12, 2, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 2, 15, 2, 17, 2, 19,
	8, 21, 9, 23, 10, 3, 2, 7, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47,
	3, 2, 36, 36, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 15, 15, 34, 34, 2, 98,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 3, 25, 3, 2, 2, 2, 5, 28, 3, 2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 35, 3, 2,
	2, 2, 11, 38, 3, 2, 2, 2, 13, 48, 3, 2, 2, 2, 15, 61, 3, 2, 2, 2, 17, 63,
	3, 2, 2, 2, 19, 65, 3, 2, 2, 2, 21, 86, 3, 2, 2, 2, 23, 90, 3, 2, 2, 2,
	25, 26, 7, 47, 2, 2, 26, 27, 7, 51, 2, 2, 27, 4, 3, 2, 2, 2, 28, 29, 7,
	47, 2, 2, 29, 30, 7, 54, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 47, 2, 2,
	32, 33, 7, 51, 2, 2, 33, 34, 7, 52, 2, 2, 34, 8, 3, 2, 2, 2, 35, 36, 7,
	80, 2, 2, 36, 37, 7, 67, 2, 2, 37, 10, 3, 2, 2, 2, 38, 45, 5, 13, 7, 2,
	39, 41, 5, 15, 8, 2, 40, 42, 5, 17, 9, 2, 41, 40, 3, 2, 2, 2, 41, 42, 3,
	2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 44, 5, 13, 7, 2, 44, 46, 3, 2, 2, 2, 45,
	39, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 12, 3, 2, 2, 2, 47, 49, 4, 50,
	59, 2, 48, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50,
	51, 3, 2, 2, 2, 51, 59, 3, 2, 2, 2, 52, 56, 7, 48, 2, 2, 53, 55, 4, 50,
	59, 2, 54, 53, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56,
	57, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 52, 3, 2, 2,
	2, 59, 60, 3, 2, 2, 2, 60, 14, 3, 2, 2, 2, 61, 62, 9, 2, 2, 2, 62, 16,
	3, 2, 2, 2, 63, 64, 9, 3, 2, 2, 64, 18, 3, 2, 2, 2, 65, 66, 7, 36, 2, 2,
	66, 67, 7, 101, 2, 2, 67, 71, 3, 2, 2, 2, 68, 70, 10, 4, 2, 2, 69, 68,
	3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 74, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 78, 7, 36, 2, 2, 75, 77, 7,
	34, 2, 2, 76, 75, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78,
	79, 3, 2, 2, 2, 79, 81, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 82, 5, 21,
	11, 2, 82, 83, 3, 2, 2, 2, 83, 84, 8, 10, 2, 2, 84, 20, 3, 2, 2, 2, 85,
	87, 9, 5, 2, 2, 86, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 86, 3, 2, 2,
	2, 88, 89, 3, 2, 2, 2, 89, 22, 3, 2, 2, 2, 90, 91, 9, 6, 2, 2, 91, 92,
	3, 2, 2, 2, 92, 93, 8, 12, 2, 2, 93, 24, 3, 2, 2, 2, 11, 2, 41, 45, 50,
	56, 59, 71, 78, 88, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'-1'", "'-4'", "'-12'", "'NA'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "FLOATINGPOINT", "COMMENT", "EOL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "FLOATINGPOINT", "NUMBER", "E", "SIGN",
	"COMMENT", "EOL", "WS",
}

type databankLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewdatabankLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *databankLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdatabankLexer(input antlr.CharStream) *databankLexer {
	l := new(databankLexer)
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
	l.GrammarFileName = "databank.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// databankLexer tokens.
const (
	databankLexerT__0          = 1
	databankLexerT__1          = 2
	databankLexerT__2          = 3
	databankLexerT__3          = 4
	databankLexerFLOATINGPOINT = 5
	databankLexerCOMMENT       = 6
	databankLexerEOL           = 7
	databankLexerWS            = 8
)
