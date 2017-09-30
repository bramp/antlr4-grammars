// Generated from tinyc.g4 by ANTLR 4.7.

package tinyc

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 84, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 6, 15, 72, 10, 15, 13, 15, 14, 15, 73, 3, 16, 6, 16, 77, 10, 16, 13,
	16, 14, 16, 78, 3, 17, 3, 17, 3, 17, 3, 17, 2, 2, 18, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 3, 2, 5, 3, 2, 99, 124, 3, 2, 50, 59, 5, 2,
	11, 12, 15, 15, 34, 34, 2, 85, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 38, 3, 2,
	2, 2, 7, 43, 3, 2, 2, 2, 9, 49, 3, 2, 2, 2, 11, 52, 3, 2, 2, 2, 13, 54,
	3, 2, 2, 2, 15, 56, 3, 2, 2, 2, 17, 58, 3, 2, 2, 2, 19, 60, 3, 2, 2, 2,
	21, 62, 3, 2, 2, 2, 23, 64, 3, 2, 2, 2, 25, 66, 3, 2, 2, 2, 27, 68, 3,
	2, 2, 2, 29, 71, 3, 2, 2, 2, 31, 76, 3, 2, 2, 2, 33, 80, 3, 2, 2, 2, 35,
	36, 7, 107, 2, 2, 36, 37, 7, 104, 2, 2, 37, 4, 3, 2, 2, 2, 38, 39, 7, 103,
	2, 2, 39, 40, 7, 110, 2, 2, 40, 41, 7, 117, 2, 2, 41, 42, 7, 103, 2, 2,
	42, 6, 3, 2, 2, 2, 43, 44, 7, 121, 2, 2, 44, 45, 7, 106, 2, 2, 45, 46,
	7, 107, 2, 2, 46, 47, 7, 110, 2, 2, 47, 48, 7, 103, 2, 2, 48, 8, 3, 2,
	2, 2, 49, 50, 7, 102, 2, 2, 50, 51, 7, 113, 2, 2, 51, 10, 3, 2, 2, 2, 52,
	53, 7, 61, 2, 2, 53, 12, 3, 2, 2, 2, 54, 55, 7, 125, 2, 2, 55, 14, 3, 2,
	2, 2, 56, 57, 7, 127, 2, 2, 57, 16, 3, 2, 2, 2, 58, 59, 7, 42, 2, 2, 59,
	18, 3, 2, 2, 2, 60, 61, 7, 43, 2, 2, 61, 20, 3, 2, 2, 2, 62, 63, 7, 63,
	2, 2, 63, 22, 3, 2, 2, 2, 64, 65, 7, 62, 2, 2, 65, 24, 3, 2, 2, 2, 66,
	67, 7, 45, 2, 2, 67, 26, 3, 2, 2, 2, 68, 69, 7, 47, 2, 2, 69, 28, 3, 2,
	2, 2, 70, 72, 9, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 71,
	3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 30, 3, 2, 2, 2, 75, 77, 9, 3, 2, 2,
	76, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3,
	2, 2, 2, 79, 32, 3, 2, 2, 2, 80, 81, 9, 4, 2, 2, 81, 82, 3, 2, 2, 2, 82,
	83, 8, 17, 2, 2, 83, 34, 3, 2, 2, 2, 5, 2, 73, 78, 3, 8, 2, 2,
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
	"", "'if'", "'else'", "'while'", "'do'", "';'", "'{'", "'}'", "'('", "')'",
	"'='", "'<'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING", "INT",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "STRING", "INT", "WS",
}

type tinycLexer struct {
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

func NewtinycLexer(input antlr.CharStream) *tinycLexer {

	l := new(tinycLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "tinyc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// tinycLexer tokens.
const (
	tinycLexerT__0   = 1
	tinycLexerT__1   = 2
	tinycLexerT__2   = 3
	tinycLexerT__3   = 4
	tinycLexerT__4   = 5
	tinycLexerT__5   = 6
	tinycLexerT__6   = 7
	tinycLexerT__7   = 8
	tinycLexerT__8   = 9
	tinycLexerT__9   = 10
	tinycLexerT__10  = 11
	tinycLexerT__11  = 12
	tinycLexerT__12  = 13
	tinycLexerSTRING = 14
	tinycLexerINT    = 15
	tinycLexerWS     = 16
)
