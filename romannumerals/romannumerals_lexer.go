// Code generated from romannumerals.g4 by ANTLR 4.9.3. DO NOT EDIT.

package romannumerals

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 103,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 6, 21, 98, 10, 21, 13, 21,
	14, 21, 99, 3, 21, 3, 21, 2, 2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 3, 2, 3, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 103, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5, 45, 3, 2, 2, 2, 7, 48,
	3, 2, 2, 2, 9, 50, 3, 2, 2, 2, 11, 53, 3, 2, 2, 2, 13, 55, 3, 2, 2, 2,
	15, 58, 3, 2, 2, 2, 17, 62, 3, 2, 2, 2, 19, 65, 3, 2, 2, 2, 21, 67, 3,
	2, 2, 2, 23, 70, 3, 2, 2, 2, 25, 72, 3, 2, 2, 2, 27, 75, 3, 2, 2, 2, 29,
	79, 3, 2, 2, 2, 31, 82, 3, 2, 2, 2, 33, 84, 3, 2, 2, 2, 35, 87, 3, 2, 2,
	2, 37, 89, 3, 2, 2, 2, 39, 92, 3, 2, 2, 2, 41, 97, 3, 2, 2, 2, 43, 44,
	7, 79, 2, 2, 44, 4, 3, 2, 2, 2, 45, 46, 7, 69, 2, 2, 46, 47, 7, 70, 2,
	2, 47, 6, 3, 2, 2, 2, 48, 49, 7, 70, 2, 2, 49, 8, 3, 2, 2, 2, 50, 51, 7,
	69, 2, 2, 51, 52, 7, 79, 2, 2, 52, 10, 3, 2, 2, 2, 53, 54, 7, 69, 2, 2,
	54, 12, 3, 2, 2, 2, 55, 56, 7, 69, 2, 2, 56, 57, 7, 69, 2, 2, 57, 14, 3,
	2, 2, 2, 58, 59, 7, 69, 2, 2, 59, 60, 7, 69, 2, 2, 60, 61, 7, 69, 2, 2,
	61, 16, 3, 2, 2, 2, 62, 63, 7, 90, 2, 2, 63, 64, 7, 78, 2, 2, 64, 18, 3,
	2, 2, 2, 65, 66, 7, 78, 2, 2, 66, 20, 3, 2, 2, 2, 67, 68, 7, 90, 2, 2,
	68, 69, 7, 69, 2, 2, 69, 22, 3, 2, 2, 2, 70, 71, 7, 90, 2, 2, 71, 24, 3,
	2, 2, 2, 72, 73, 7, 90, 2, 2, 73, 74, 7, 90, 2, 2, 74, 26, 3, 2, 2, 2,
	75, 76, 7, 90, 2, 2, 76, 77, 7, 90, 2, 2, 77, 78, 7, 90, 2, 2, 78, 28,
	3, 2, 2, 2, 79, 80, 7, 75, 2, 2, 80, 81, 7, 88, 2, 2, 81, 30, 3, 2, 2,
	2, 82, 83, 7, 88, 2, 2, 83, 32, 3, 2, 2, 2, 84, 85, 7, 75, 2, 2, 85, 86,
	7, 90, 2, 2, 86, 34, 3, 2, 2, 2, 87, 88, 7, 75, 2, 2, 88, 36, 3, 2, 2,
	2, 89, 90, 7, 75, 2, 2, 90, 91, 7, 75, 2, 2, 91, 38, 3, 2, 2, 2, 92, 93,
	7, 75, 2, 2, 93, 94, 7, 75, 2, 2, 94, 95, 7, 75, 2, 2, 95, 40, 3, 2, 2,
	2, 96, 98, 9, 2, 2, 2, 97, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 97,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 8, 21,
	2, 2, 102, 42, 3, 2, 2, 2, 4, 2, 99, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'M'", "'CD'", "'D'", "'CM'", "'C'", "'CC'", "'CCC'", "'XL'", "'L'",
	"'XC'", "'X'", "'XX'", "'XXX'", "'IV'", "'V'", "'IX'", "'I'", "'II'", "'III'",
}

var lexerSymbolicNames = []string{
	"", "M", "CD", "D", "CM", "C", "CC", "CCC", "XL", "L", "XC", "X", "XX",
	"XXX", "IV", "V", "IX", "I", "II", "III", "WS",
}

var lexerRuleNames = []string{
	"M", "CD", "D", "CM", "C", "CC", "CCC", "XL", "L", "XC", "X", "XX", "XXX",
	"IV", "V", "IX", "I", "II", "III", "WS",
}

type romannumeralsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewromannumeralsLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *romannumeralsLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewromannumeralsLexer(input antlr.CharStream) *romannumeralsLexer {
	l := new(romannumeralsLexer)
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
	l.GrammarFileName = "romannumerals.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// romannumeralsLexer tokens.
const (
	romannumeralsLexerM   = 1
	romannumeralsLexerCD  = 2
	romannumeralsLexerD   = 3
	romannumeralsLexerCM  = 4
	romannumeralsLexerC   = 5
	romannumeralsLexerCC  = 6
	romannumeralsLexerCCC = 7
	romannumeralsLexerXL  = 8
	romannumeralsLexerL   = 9
	romannumeralsLexerXC  = 10
	romannumeralsLexerX   = 11
	romannumeralsLexerXX  = 12
	romannumeralsLexerXXX = 13
	romannumeralsLexerIV  = 14
	romannumeralsLexerV   = 15
	romannumeralsLexerIX  = 16
	romannumeralsLexerI   = 17
	romannumeralsLexerII  = 18
	romannumeralsLexerIII = 19
	romannumeralsLexerWS  = 20
)
