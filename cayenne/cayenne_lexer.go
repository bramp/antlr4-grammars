// Code generated from cayenne.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cayenne

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 24, 142,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 7, 22, 131, 10, 22,
	12, 22, 14, 22, 134, 11, 22, 3, 23, 6, 23, 137, 10, 23, 13, 23, 14, 23,
	138, 3, 23, 3, 23, 2, 2, 24, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 3, 2, 5, 4, 2, 67,
	92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34,
	2, 143, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 3, 47, 3, 2,
	2, 2, 5, 49, 3, 2, 2, 2, 7, 52, 3, 2, 2, 2, 9, 54, 3, 2, 2, 2, 11, 57,
	3, 2, 2, 2, 13, 59, 3, 2, 2, 2, 15, 64, 3, 2, 2, 2, 17, 66, 3, 2, 2, 2,
	19, 68, 3, 2, 2, 2, 21, 73, 3, 2, 2, 2, 23, 76, 3, 2, 2, 2, 25, 80, 3,
	2, 2, 2, 27, 87, 3, 2, 2, 2, 29, 89, 3, 2, 2, 2, 31, 91, 3, 2, 2, 2, 33,
	93, 3, 2, 2, 2, 35, 95, 3, 2, 2, 2, 37, 103, 3, 2, 2, 2, 39, 110, 3, 2,
	2, 2, 41, 119, 3, 2, 2, 2, 43, 128, 3, 2, 2, 2, 45, 136, 3, 2, 2, 2, 47,
	48, 7, 42, 2, 2, 48, 4, 3, 2, 2, 2, 49, 50, 7, 60, 2, 2, 50, 51, 7, 60,
	2, 2, 51, 6, 3, 2, 2, 2, 52, 53, 7, 43, 2, 2, 53, 8, 3, 2, 2, 2, 54, 55,
	7, 47, 2, 2, 55, 56, 7, 64, 2, 2, 56, 10, 3, 2, 2, 2, 57, 58, 7, 94, 2,
	2, 58, 12, 3, 2, 2, 2, 59, 60, 7, 102, 2, 2, 60, 61, 7, 99, 2, 2, 61, 62,
	7, 118, 2, 2, 62, 63, 7, 99, 2, 2, 63, 14, 3, 2, 2, 2, 64, 65, 7, 126,
	2, 2, 65, 16, 3, 2, 2, 2, 66, 67, 7, 66, 2, 2, 67, 18, 3, 2, 2, 2, 68,
	69, 7, 101, 2, 2, 69, 70, 7, 99, 2, 2, 70, 71, 7, 117, 2, 2, 71, 72, 7,
	103, 2, 2, 72, 20, 3, 2, 2, 2, 73, 74, 7, 113, 2, 2, 74, 75, 7, 104, 2,
	2, 75, 22, 3, 2, 2, 2, 76, 77, 7, 117, 2, 2, 77, 78, 7, 107, 2, 2, 78,
	79, 7, 105, 2, 2, 79, 24, 3, 2, 2, 2, 80, 81, 7, 117, 2, 2, 81, 82, 7,
	118, 2, 2, 82, 83, 7, 116, 2, 2, 83, 84, 7, 119, 2, 2, 84, 85, 7, 101,
	2, 2, 85, 86, 7, 118, 2, 2, 86, 26, 3, 2, 2, 2, 87, 88, 7, 48, 2, 2, 88,
	28, 3, 2, 2, 2, 89, 90, 7, 37, 2, 2, 90, 30, 3, 2, 2, 2, 91, 92, 7, 61,
	2, 2, 92, 32, 3, 2, 2, 2, 93, 94, 7, 63, 2, 2, 94, 34, 3, 2, 2, 2, 95,
	96, 7, 114, 2, 2, 96, 97, 7, 116, 2, 2, 97, 98, 7, 107, 2, 2, 98, 99, 7,
	120, 2, 2, 99, 100, 7, 99, 2, 2, 100, 101, 7, 118, 2, 2, 101, 102, 7, 103,
	2, 2, 102, 36, 3, 2, 2, 2, 103, 104, 7, 114, 2, 2, 104, 105, 7, 119, 2,
	2, 105, 106, 7, 100, 2, 2, 106, 107, 7, 110, 2, 2, 107, 108, 7, 107, 2,
	2, 108, 109, 7, 101, 2, 2, 109, 38, 3, 2, 2, 2, 110, 111, 7, 99, 2, 2,
	111, 112, 7, 100, 2, 2, 112, 113, 7, 117, 2, 2, 113, 114, 7, 118, 2, 2,
	114, 115, 7, 116, 2, 2, 115, 116, 7, 99, 2, 2, 116, 117, 7, 101, 2, 2,
	117, 118, 7, 118, 2, 2, 118, 40, 3, 2, 2, 2, 119, 120, 7, 101, 2, 2, 120,
	121, 7, 113, 2, 2, 121, 122, 7, 112, 2, 2, 122, 123, 7, 101, 2, 2, 123,
	124, 7, 116, 2, 2, 124, 125, 7, 103, 2, 2, 125, 126, 7, 118, 2, 2, 126,
	127, 7, 103, 2, 2, 127, 42, 3, 2, 2, 2, 128, 132, 9, 2, 2, 2, 129, 131,
	9, 3, 2, 2, 130, 129, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2,
	2, 2, 132, 133, 3, 2, 2, 2, 133, 44, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2,
	135, 137, 9, 4, 2, 2, 136, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138,
	136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141,
	8, 23, 2, 2, 141, 46, 3, 2, 2, 2, 5, 2, 132, 138, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "'::'", "')'", "'->'", "'\\'", "'data'", "'|'", "'@'", "'case'",
	"'of'", "'sig'", "'struct'", "'.'", "'#'", "';'", "'='", "'private'", "'public'",
	"'abstract'", "'concrete'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "ID", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "ID", "WS",
}

type cayenneLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewcayenneLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *cayenneLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewcayenneLexer(input antlr.CharStream) *cayenneLexer {
	l := new(cayenneLexer)
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
	l.GrammarFileName = "cayenne.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// cayenneLexer tokens.
const (
	cayenneLexerT__0  = 1
	cayenneLexerT__1  = 2
	cayenneLexerT__2  = 3
	cayenneLexerT__3  = 4
	cayenneLexerT__4  = 5
	cayenneLexerT__5  = 6
	cayenneLexerT__6  = 7
	cayenneLexerT__7  = 8
	cayenneLexerT__8  = 9
	cayenneLexerT__9  = 10
	cayenneLexerT__10 = 11
	cayenneLexerT__11 = 12
	cayenneLexerT__12 = 13
	cayenneLexerT__13 = 14
	cayenneLexerT__14 = 15
	cayenneLexerT__15 = 16
	cayenneLexerT__16 = 17
	cayenneLexerT__17 = 18
	cayenneLexerT__18 = 19
	cayenneLexerT__19 = 20
	cayenneLexerID    = 21
	cayenneLexerWS    = 22
)
