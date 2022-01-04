// Code generated from gedcom.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gedcom

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 89, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18,
	9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 20, 6, 20, 82, 10, 20, 13, 20, 14, 20, 83, 3, 21, 3, 21, 3, 21, 3, 21,
	2, 2, 22, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 3, 2, 6, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4,
	2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 2, 89, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3, 2,
	2, 2, 5, 45, 3, 2, 2, 2, 7, 47, 3, 2, 2, 2, 9, 50, 3, 2, 2, 2, 11, 52,
	3, 2, 2, 2, 13, 54, 3, 2, 2, 2, 15, 56, 3, 2, 2, 2, 17, 58, 3, 2, 2, 2,
	19, 60, 3, 2, 2, 2, 21, 62, 3, 2, 2, 2, 23, 64, 3, 2, 2, 2, 25, 66, 3,
	2, 2, 2, 27, 68, 3, 2, 2, 2, 29, 70, 3, 2, 2, 2, 31, 72, 3, 2, 2, 2, 33,
	74, 3, 2, 2, 2, 35, 76, 3, 2, 2, 2, 37, 78, 3, 2, 2, 2, 39, 81, 3, 2, 2,
	2, 41, 85, 3, 2, 2, 2, 43, 44, 7, 66, 2, 2, 44, 4, 3, 2, 2, 2, 45, 46,
	7, 37, 2, 2, 46, 6, 3, 2, 2, 2, 47, 48, 7, 66, 2, 2, 48, 49, 7, 66, 2,
	2, 49, 8, 3, 2, 2, 2, 50, 51, 7, 35, 2, 2, 51, 10, 3, 2, 2, 2, 52, 53,
	7, 36, 2, 2, 53, 12, 3, 2, 2, 2, 54, 55, 7, 38, 2, 2, 55, 14, 3, 2, 2,
	2, 56, 57, 7, 40, 2, 2, 57, 16, 3, 2, 2, 2, 58, 59, 7, 41, 2, 2, 59, 18,
	3, 2, 2, 2, 60, 61, 7, 42, 2, 2, 61, 20, 3, 2, 2, 2, 62, 63, 7, 43, 2,
	2, 63, 22, 3, 2, 2, 2, 64, 65, 7, 44, 2, 2, 65, 24, 3, 2, 2, 2, 66, 67,
	7, 45, 2, 2, 67, 26, 3, 2, 2, 2, 68, 69, 7, 47, 2, 2, 69, 28, 3, 2, 2,
	2, 70, 71, 7, 46, 2, 2, 71, 30, 3, 2, 2, 2, 72, 73, 7, 48, 2, 2, 73, 32,
	3, 2, 2, 2, 74, 75, 7, 49, 2, 2, 75, 34, 3, 2, 2, 2, 76, 77, 9, 2, 2, 2,
	77, 36, 3, 2, 2, 2, 78, 79, 9, 3, 2, 2, 79, 38, 3, 2, 2, 2, 80, 82, 9,
	4, 2, 2, 81, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83,
	84, 3, 2, 2, 2, 84, 40, 3, 2, 2, 2, 85, 86, 9, 5, 2, 2, 86, 87, 3, 2, 2,
	2, 87, 88, 8, 21, 2, 2, 88, 42, 3, 2, 2, 2, 4, 2, 83, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'@'", "'#'", "'@@'", "'!'", "'\"'", "'$'", "'&'", "'''", "'('", "')'",
	"'*'", "'+'", "'-'", "','", "'.'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ALPHA",
	"DIGIT", "EOL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "ALPHA",
	"DIGIT", "EOL", "WS",
}

type gedcomLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewgedcomLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *gedcomLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewgedcomLexer(input antlr.CharStream) *gedcomLexer {
	l := new(gedcomLexer)
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
	l.GrammarFileName = "gedcom.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gedcomLexer tokens.
const (
	gedcomLexerT__0  = 1
	gedcomLexerT__1  = 2
	gedcomLexerT__2  = 3
	gedcomLexerT__3  = 4
	gedcomLexerT__4  = 5
	gedcomLexerT__5  = 6
	gedcomLexerT__6  = 7
	gedcomLexerT__7  = 8
	gedcomLexerT__8  = 9
	gedcomLexerT__9  = 10
	gedcomLexerT__10 = 11
	gedcomLexerT__11 = 12
	gedcomLexerT__12 = 13
	gedcomLexerT__13 = 14
	gedcomLexerT__14 = 15
	gedcomLexerT__15 = 16
	gedcomLexerALPHA = 17
	gedcomLexerDIGIT = 18
	gedcomLexerEOL   = 19
	gedcomLexerWS    = 20
)
