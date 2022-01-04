// Code generated from iso8601.g4 by ANTLR 4.9.3. DO NOT EDIT.

package iso8601

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 77, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18,
	9, 18, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 18, 3, 18, 6, 18, 74, 10, 18, 13, 18, 14, 18, 75, 2, 2, 19, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 3, 2, 15, 4, 2, 12, 12,
	15, 15, 4, 2, 86, 86, 118, 118, 4, 2, 92, 92, 124, 124, 4, 2, 89, 89, 121,
	121, 4, 2, 82, 82, 114, 114, 4, 2, 91, 91, 123, 123, 4, 2, 79, 79, 111,
	111, 4, 2, 70, 70, 102, 102, 4, 2, 74, 74, 106, 106, 4, 2, 85, 85, 117,
	117, 4, 2, 84, 84, 116, 116, 3, 2, 50, 59, 4, 2, 46, 46, 48, 48, 2, 77,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 2, 35, 3, 2, 2, 2, 3, 37, 3, 2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 41,
	3, 2, 2, 2, 9, 43, 3, 2, 2, 2, 11, 45, 3, 2, 2, 2, 13, 49, 3, 2, 2, 2,
	15, 51, 3, 2, 2, 2, 17, 53, 3, 2, 2, 2, 19, 55, 3, 2, 2, 2, 21, 57, 3,
	2, 2, 2, 23, 59, 3, 2, 2, 2, 25, 61, 3, 2, 2, 2, 27, 63, 3, 2, 2, 2, 29,
	65, 3, 2, 2, 2, 31, 67, 3, 2, 2, 2, 33, 69, 3, 2, 2, 2, 35, 71, 3, 2, 2,
	2, 37, 38, 7, 45, 2, 2, 38, 4, 3, 2, 2, 2, 39, 40, 7, 47, 2, 2, 40, 6,
	3, 2, 2, 2, 41, 42, 7, 60, 2, 2, 42, 8, 3, 2, 2, 2, 43, 44, 7, 49, 2, 2,
	44, 10, 3, 2, 2, 2, 45, 46, 9, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 48, 8,
	6, 2, 2, 48, 12, 3, 2, 2, 2, 49, 50, 9, 3, 2, 2, 50, 14, 3, 2, 2, 2, 51,
	52, 9, 4, 2, 2, 52, 16, 3, 2, 2, 2, 53, 54, 9, 5, 2, 2, 54, 18, 3, 2, 2,
	2, 55, 56, 9, 6, 2, 2, 56, 20, 3, 2, 2, 2, 57, 58, 9, 7, 2, 2, 58, 22,
	3, 2, 2, 2, 59, 60, 9, 8, 2, 2, 60, 24, 3, 2, 2, 2, 61, 62, 9, 9, 2, 2,
	62, 26, 3, 2, 2, 2, 63, 64, 9, 10, 2, 2, 64, 28, 3, 2, 2, 2, 65, 66, 9,
	11, 2, 2, 66, 30, 3, 2, 2, 2, 67, 68, 9, 12, 2, 2, 68, 32, 3, 2, 2, 2,
	69, 70, 9, 13, 2, 2, 70, 34, 3, 2, 2, 2, 71, 73, 9, 14, 2, 2, 72, 74, 5,
	33, 17, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2,
	75, 76, 3, 2, 2, 2, 76, 36, 3, 2, 2, 2, 4, 2, 75, 3, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'+'", "'-'", "':'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "Newline", "T", "Z", "W", "P", "Y", "M", "D", "H",
	"S", "R", "Digit", "Fraction",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "Newline", "T", "Z", "W", "P", "Y", "M",
	"D", "H", "S", "R", "Digit", "Fraction",
}

type iso8601Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// Newiso8601Lexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *iso8601Lexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newiso8601Lexer(input antlr.CharStream) *iso8601Lexer {
	l := new(iso8601Lexer)
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
	l.GrammarFileName = "iso8601.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// iso8601Lexer tokens.
const (
	iso8601LexerT__0     = 1
	iso8601LexerT__1     = 2
	iso8601LexerT__2     = 3
	iso8601LexerT__3     = 4
	iso8601LexerNewline  = 5
	iso8601LexerT        = 6
	iso8601LexerZ        = 7
	iso8601LexerW        = 8
	iso8601LexerP        = 9
	iso8601LexerY        = 10
	iso8601LexerM        = 11
	iso8601LexerD        = 12
	iso8601LexerH        = 13
	iso8601LexerS        = 14
	iso8601LexerR        = 15
	iso8601LexerDigit    = 16
	iso8601LexerFraction = 17
)
