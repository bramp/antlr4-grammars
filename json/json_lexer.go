// Code generated from JSON.g4 by ANTLR 4.9.3. DO NOT EDIT.

package json

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 130,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	7, 11, 71, 10, 11, 12, 11, 14, 11, 74, 11, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 5, 12, 81, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 5, 16, 94, 10, 16, 3, 16, 3, 16, 3,
	16, 6, 16, 99, 10, 16, 13, 16, 14, 16, 100, 5, 16, 103, 10, 16, 3, 16,
	5, 16, 106, 10, 16, 3, 17, 3, 17, 3, 17, 7, 17, 111, 10, 17, 12, 17, 14,
	17, 114, 11, 17, 5, 17, 116, 10, 17, 3, 18, 3, 18, 5, 18, 120, 10, 18,
	3, 18, 3, 18, 3, 19, 6, 19, 125, 10, 19, 13, 19, 14, 19, 126, 3, 19, 3,
	19, 2, 2, 20, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 2, 25, 2, 27, 2, 29, 2, 31, 13, 33, 2, 35, 2, 37, 14, 3,
	2, 10, 10, 2, 36, 36, 49, 49, 94, 94, 100, 100, 104, 104, 112, 112, 116,
	116, 118, 118, 5, 2, 50, 59, 67, 72, 99, 104, 5, 2, 2, 33, 36, 36, 94,
	94, 3, 2, 50, 59, 3, 2, 51, 59, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47,
	47, 5, 2, 11, 12, 15, 15, 34, 34, 2, 134, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3, 39, 3, 2, 2, 2, 5,
	41, 3, 2, 2, 2, 7, 43, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 47, 3, 2, 2,
	2, 13, 49, 3, 2, 2, 2, 15, 51, 3, 2, 2, 2, 17, 56, 3, 2, 2, 2, 19, 62,
	3, 2, 2, 2, 21, 67, 3, 2, 2, 2, 23, 77, 3, 2, 2, 2, 25, 82, 3, 2, 2, 2,
	27, 88, 3, 2, 2, 2, 29, 90, 3, 2, 2, 2, 31, 93, 3, 2, 2, 2, 33, 115, 3,
	2, 2, 2, 35, 117, 3, 2, 2, 2, 37, 124, 3, 2, 2, 2, 39, 40, 7, 125, 2, 2,
	40, 4, 3, 2, 2, 2, 41, 42, 7, 46, 2, 2, 42, 6, 3, 2, 2, 2, 43, 44, 7, 127,
	2, 2, 44, 8, 3, 2, 2, 2, 45, 46, 7, 60, 2, 2, 46, 10, 3, 2, 2, 2, 47, 48,
	7, 93, 2, 2, 48, 12, 3, 2, 2, 2, 49, 50, 7, 95, 2, 2, 50, 14, 3, 2, 2,
	2, 51, 52, 7, 118, 2, 2, 52, 53, 7, 116, 2, 2, 53, 54, 7, 119, 2, 2, 54,
	55, 7, 103, 2, 2, 55, 16, 3, 2, 2, 2, 56, 57, 7, 104, 2, 2, 57, 58, 7,
	99, 2, 2, 58, 59, 7, 110, 2, 2, 59, 60, 7, 117, 2, 2, 60, 61, 7, 103, 2,
	2, 61, 18, 3, 2, 2, 2, 62, 63, 7, 112, 2, 2, 63, 64, 7, 119, 2, 2, 64,
	65, 7, 110, 2, 2, 65, 66, 7, 110, 2, 2, 66, 20, 3, 2, 2, 2, 67, 72, 7,
	36, 2, 2, 68, 71, 5, 23, 12, 2, 69, 71, 5, 29, 15, 2, 70, 68, 3, 2, 2,
	2, 70, 69, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73,
	3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 76, 7, 36, 2, 2,
	76, 22, 3, 2, 2, 2, 77, 80, 7, 94, 2, 2, 78, 81, 9, 2, 2, 2, 79, 81, 5,
	25, 13, 2, 80, 78, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81, 24, 3, 2, 2, 2,
	82, 83, 7, 119, 2, 2, 83, 84, 5, 27, 14, 2, 84, 85, 5, 27, 14, 2, 85, 86,
	5, 27, 14, 2, 86, 87, 5, 27, 14, 2, 87, 26, 3, 2, 2, 2, 88, 89, 9, 3, 2,
	2, 89, 28, 3, 2, 2, 2, 90, 91, 10, 4, 2, 2, 91, 30, 3, 2, 2, 2, 92, 94,
	7, 47, 2, 2, 93, 92, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2,
	95, 102, 5, 33, 17, 2, 96, 98, 7, 48, 2, 2, 97, 99, 9, 5, 2, 2, 98, 97,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 103, 3, 2, 2, 2, 102, 96, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103,
	105, 3, 2, 2, 2, 104, 106, 5, 35, 18, 2, 105, 104, 3, 2, 2, 2, 105, 106,
	3, 2, 2, 2, 106, 32, 3, 2, 2, 2, 107, 116, 7, 50, 2, 2, 108, 112, 9, 6,
	2, 2, 109, 111, 9, 5, 2, 2, 110, 109, 3, 2, 2, 2, 111, 114, 3, 2, 2, 2,
	112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 116, 3, 2, 2, 2, 114,
	112, 3, 2, 2, 2, 115, 107, 3, 2, 2, 2, 115, 108, 3, 2, 2, 2, 116, 34, 3,
	2, 2, 2, 117, 119, 9, 7, 2, 2, 118, 120, 9, 8, 2, 2, 119, 118, 3, 2, 2,
	2, 119, 120, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 122, 5, 33, 17, 2,
	122, 36, 3, 2, 2, 2, 123, 125, 9, 9, 2, 2, 124, 123, 3, 2, 2, 2, 125, 126,
	3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 3, 2,
	2, 2, 128, 129, 8, 19, 2, 2, 129, 38, 3, 2, 2, 2, 14, 2, 70, 72, 80, 93,
	100, 102, 105, 112, 115, 119, 126, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'{'", "','", "'}'", "':'", "'['", "']'", "'true'", "'false'", "'null'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"STRING", "ESC", "UNICODE", "HEX", "SAFECODEPOINT", "NUMBER", "INT", "EXP",
	"WS",
}

type JSONLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewJSONLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *JSONLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewJSONLexer(input antlr.CharStream) *JSONLexer {
	l := new(JSONLexer)
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
	l.GrammarFileName = "JSON.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JSONLexer tokens.
const (
	JSONLexerT__0   = 1
	JSONLexerT__1   = 2
	JSONLexerT__2   = 3
	JSONLexerT__3   = 4
	JSONLexerT__4   = 5
	JSONLexerT__5   = 6
	JSONLexerT__6   = 7
	JSONLexerT__7   = 8
	JSONLexerT__8   = 9
	JSONLexerSTRING = 10
	JSONLexerNUMBER = 11
	JSONLexerWS     = 12
)
