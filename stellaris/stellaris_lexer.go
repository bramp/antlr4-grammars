// Code generated from stellaris.g4 by ANTLR 4.9.3. DO NOT EDIT.

package stellaris

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 16, 102,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 7, 10, 56, 10, 10, 12,
	10, 14, 10, 59, 11, 10, 3, 11, 5, 11, 62, 10, 11, 3, 11, 3, 11, 3, 12,
	6, 12, 67, 10, 12, 13, 12, 14, 12, 68, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14,
	75, 10, 14, 3, 15, 3, 15, 7, 15, 79, 10, 15, 12, 15, 14, 15, 82, 11, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 7, 16, 88, 10, 16, 12, 16, 14, 16, 91, 11,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	2, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 2, 25, 2, 27, 2, 29, 13, 31, 14, 33, 15, 35, 16, 3, 2, 9, 4,
	2, 45, 45, 47, 47, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 4, 2, 50, 59, 97,
	97, 5, 2, 12, 12, 15, 15, 36, 36, 4, 2, 12, 12, 15, 15, 5, 2, 11, 11, 14,
	14, 34, 34, 2, 104, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 29, 3,
	2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 3, 37,
	3, 2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 41, 3, 2, 2, 2, 9, 43, 3, 2, 2, 2, 11,
	45, 3, 2, 2, 2, 13, 47, 3, 2, 2, 2, 15, 49, 3, 2, 2, 2, 17, 51, 3, 2, 2,
	2, 19, 53, 3, 2, 2, 2, 21, 61, 3, 2, 2, 2, 23, 66, 3, 2, 2, 2, 25, 70,
	3, 2, 2, 2, 27, 74, 3, 2, 2, 2, 29, 76, 3, 2, 2, 2, 31, 85, 3, 2, 2, 2,
	33, 94, 3, 2, 2, 2, 35, 98, 3, 2, 2, 2, 37, 38, 7, 63, 2, 2, 38, 4, 3,
	2, 2, 2, 39, 40, 7, 64, 2, 2, 40, 6, 3, 2, 2, 2, 41, 42, 7, 62, 2, 2, 42,
	8, 3, 2, 2, 2, 43, 44, 7, 48, 2, 2, 44, 10, 3, 2, 2, 2, 45, 46, 7, 66,
	2, 2, 46, 12, 3, 2, 2, 2, 47, 48, 7, 60, 2, 2, 48, 14, 3, 2, 2, 2, 49,
	50, 7, 125, 2, 2, 50, 16, 3, 2, 2, 2, 51, 52, 7, 127, 2, 2, 52, 18, 3,
	2, 2, 2, 53, 57, 5, 25, 13, 2, 54, 56, 5, 27, 14, 2, 55, 54, 3, 2, 2, 2,
	56, 59, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 20, 3,
	2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 62, 9, 2, 2, 2, 61, 60, 3, 2, 2, 2, 61,
	62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 5, 23, 12, 2, 64, 22, 3, 2,
	2, 2, 65, 67, 9, 3, 2, 2, 66, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 66,
	3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 24, 3, 2, 2, 2, 70, 71, 9, 4, 2, 2,
	71, 26, 3, 2, 2, 2, 72, 75, 5, 25, 13, 2, 73, 75, 9, 5, 2, 2, 74, 72, 3,
	2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 28, 3, 2, 2, 2, 76, 80, 7, 36, 2, 2, 77,
	79, 10, 6, 2, 2, 78, 77, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2,
	2, 2, 80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 84,
	7, 36, 2, 2, 84, 30, 3, 2, 2, 2, 85, 89, 7, 37, 2, 2, 86, 88, 10, 7, 2,
	2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90,
	3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 8, 16, 2, 2,
	93, 32, 3, 2, 2, 2, 94, 95, 9, 8, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 8,
	17, 2, 2, 97, 34, 3, 2, 2, 2, 98, 99, 9, 7, 2, 2, 99, 100, 3, 2, 2, 2,
	100, 101, 8, 18, 2, 2, 101, 36, 3, 2, 2, 2, 9, 2, 57, 61, 68, 74, 80, 89,
	3, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'>'", "'<'", "'.'", "'@'", "':'", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "IDENTIFIER", "INTEGER", "STRING",
	"COMMENT", "SPACE", "NL",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "IDENTIFIER",
	"INTEGER", "INTEGERFRAG", "IDENITIFIERHEAD", "IDENITIFIERBODY", "STRING",
	"COMMENT", "SPACE", "NL",
}

type stellarisLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewstellarisLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *stellarisLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewstellarisLexer(input antlr.CharStream) *stellarisLexer {
	l := new(stellarisLexer)
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
	l.GrammarFileName = "stellaris.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// stellarisLexer tokens.
const (
	stellarisLexerT__0       = 1
	stellarisLexerT__1       = 2
	stellarisLexerT__2       = 3
	stellarisLexerT__3       = 4
	stellarisLexerT__4       = 5
	stellarisLexerT__5       = 6
	stellarisLexerT__6       = 7
	stellarisLexerT__7       = 8
	stellarisLexerIDENTIFIER = 9
	stellarisLexerINTEGER    = 10
	stellarisLexerSTRING     = 11
	stellarisLexerCOMMENT    = 12
	stellarisLexerSPACE      = 13
	stellarisLexerNL         = 14
)
