// Code generated from creole.g4 by ANTLR 4.9.3. DO NOT EDIT.

package creole

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 116,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 6, 14, 82, 10, 14, 13, 14, 14, 14, 83, 3, 15, 3, 15, 3, 16, 5, 16,
	89, 10, 16, 3, 16, 3, 16, 5, 16, 93, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 7, 17, 100, 10, 17, 12, 17, 14, 17, 103, 11, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 101,
	2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 2, 39,
	2, 41, 2, 3, 2, 6, 4, 2, 11, 11, 34, 34, 4, 2, 67, 92, 99, 124, 3, 2, 50,
	59, 8, 2, 36, 36, 41, 43, 45, 48, 60, 61, 94, 94, 128, 128, 2, 119, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2,
	2, 2, 2, 35, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5, 46, 3, 2, 2, 2, 7, 49, 3,
	2, 2, 2, 9, 51, 3, 2, 2, 2, 11, 56, 3, 2, 2, 2, 13, 58, 3, 2, 2, 2, 15,
	61, 3, 2, 2, 2, 17, 63, 3, 2, 2, 2, 19, 65, 3, 2, 2, 2, 21, 68, 3, 2, 2,
	2, 23, 71, 3, 2, 2, 2, 25, 74, 3, 2, 2, 2, 27, 81, 3, 2, 2, 2, 29, 85,
	3, 2, 2, 2, 31, 92, 3, 2, 2, 2, 33, 94, 3, 2, 2, 2, 35, 108, 3, 2, 2, 2,
	37, 110, 3, 2, 2, 2, 39, 112, 3, 2, 2, 2, 41, 114, 3, 2, 2, 2, 43, 44,
	7, 94, 2, 2, 44, 45, 7, 94, 2, 2, 45, 4, 3, 2, 2, 2, 46, 47, 7, 44, 2,
	2, 47, 48, 7, 44, 2, 2, 48, 6, 3, 2, 2, 2, 49, 50, 7, 126, 2, 2, 50, 8,
	3, 2, 2, 2, 51, 52, 7, 47, 2, 2, 52, 53, 7, 47, 2, 2, 53, 54, 7, 47, 2,
	2, 54, 55, 7, 47, 2, 2, 55, 10, 3, 2, 2, 2, 56, 57, 7, 44, 2, 2, 57, 12,
	3, 2, 2, 2, 58, 59, 7, 126, 2, 2, 59, 60, 7, 63, 2, 2, 60, 14, 3, 2, 2,
	2, 61, 62, 7, 63, 2, 2, 62, 16, 3, 2, 2, 2, 63, 64, 7, 37, 2, 2, 64, 18,
	3, 2, 2, 2, 65, 66, 7, 93, 2, 2, 66, 67, 7, 93, 2, 2, 67, 20, 3, 2, 2,
	2, 68, 69, 7, 95, 2, 2, 69, 70, 7, 95, 2, 2, 70, 22, 3, 2, 2, 2, 71, 72,
	7, 125, 2, 2, 72, 73, 7, 125, 2, 2, 73, 24, 3, 2, 2, 2, 74, 75, 7, 127,
	2, 2, 75, 76, 7, 127, 2, 2, 76, 26, 3, 2, 2, 2, 77, 82, 5, 37, 19, 2, 78,
	82, 5, 39, 20, 2, 79, 82, 5, 41, 21, 2, 80, 82, 5, 29, 15, 2, 81, 77, 3,
	2, 2, 2, 81, 78, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82,
	83, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 28, 3, 2, 2,
	2, 85, 86, 9, 2, 2, 2, 86, 30, 3, 2, 2, 2, 87, 89, 7, 15, 2, 2, 88, 87,
	3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 93, 7, 12, 2, 2,
	91, 93, 7, 2, 2, 3, 92, 88, 3, 2, 2, 2, 92, 91, 3, 2, 2, 2, 93, 32, 3,
	2, 2, 2, 94, 95, 7, 125, 2, 2, 95, 96, 7, 125, 2, 2, 96, 97, 7, 125, 2,
	2, 97, 101, 3, 2, 2, 2, 98, 100, 11, 2, 2, 2, 99, 98, 3, 2, 2, 2, 100,
	103, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 102, 104, 3,
	2, 2, 2, 103, 101, 3, 2, 2, 2, 104, 105, 7, 127, 2, 2, 105, 106, 7, 127,
	2, 2, 106, 107, 7, 127, 2, 2, 107, 34, 3, 2, 2, 2, 108, 109, 7, 49, 2,
	2, 109, 36, 3, 2, 2, 2, 110, 111, 9, 3, 2, 2, 111, 38, 3, 2, 2, 2, 112,
	113, 9, 4, 2, 2, 113, 40, 3, 2, 2, 2, 114, 115, 9, 5, 2, 2, 115, 42, 3,
	2, 2, 2, 8, 2, 81, 83, 88, 92, 101, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'\\\\'", "'**'", "'|'", "'----'", "'*'", "'|='", "'='", "'#'", "'[['",
	"']]'", "'{{'", "'}}'", "", "", "", "", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "HASH", "LBRACKET", "RBRACKET", "LBRACE",
	"RBRACE", "TEXT", "WS", "CR", "NOWIKI", "RSLASH",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "HASH", "LBRACKET",
	"RBRACKET", "LBRACE", "RBRACE", "TEXT", "WS", "CR", "NOWIKI", "RSLASH",
	"LETTERS", "DIGITS", "SYMBOL",
}

type creoleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewcreoleLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *creoleLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewcreoleLexer(input antlr.CharStream) *creoleLexer {
	l := new(creoleLexer)
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
	l.GrammarFileName = "creole.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// creoleLexer tokens.
const (
	creoleLexerT__0     = 1
	creoleLexerT__1     = 2
	creoleLexerT__2     = 3
	creoleLexerT__3     = 4
	creoleLexerT__4     = 5
	creoleLexerT__5     = 6
	creoleLexerT__6     = 7
	creoleLexerHASH     = 8
	creoleLexerLBRACKET = 9
	creoleLexerRBRACKET = 10
	creoleLexerLBRACE   = 11
	creoleLexerRBRACE   = 12
	creoleLexerTEXT     = 13
	creoleLexerWS       = 14
	creoleLexerCR       = 15
	creoleLexerNOWIKI   = 16
	creoleLexerRSLASH   = 17
)
