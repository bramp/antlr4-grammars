// Code generated from fol.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fol

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 97, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18,
	9, 18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 6, 18, 87, 10, 18, 13, 18, 14,
	18, 88, 3, 19, 6, 19, 92, 10, 19, 13, 19, 14, 19, 93, 3, 19, 3, 19, 2,
	2, 20, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 3,
	2, 5, 5, 2, 50, 59, 67, 92, 99, 124, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11,
	34, 34, 2, 98, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3, 39, 3,
	2, 2, 2, 5, 41, 3, 2, 2, 2, 7, 43, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 47,
	3, 2, 2, 2, 13, 49, 3, 2, 2, 2, 15, 51, 3, 2, 2, 2, 17, 53, 3, 2, 2, 2,
	19, 55, 3, 2, 2, 2, 21, 57, 3, 2, 2, 2, 23, 64, 3, 2, 2, 2, 25, 71, 3,
	2, 2, 2, 27, 73, 3, 2, 2, 2, 29, 76, 3, 2, 2, 2, 31, 78, 3, 2, 2, 2, 33,
	81, 3, 2, 2, 2, 35, 86, 3, 2, 2, 2, 37, 91, 3, 2, 2, 2, 39, 40, 7, 65,
	2, 2, 40, 4, 3, 2, 2, 2, 41, 42, 7, 97, 2, 2, 42, 6, 3, 2, 2, 2, 43, 44,
	7, 37, 2, 2, 44, 8, 3, 2, 2, 2, 45, 46, 7, 48, 2, 2, 46, 10, 3, 2, 2, 2,
	47, 48, 7, 46, 2, 2, 48, 12, 3, 2, 2, 2, 49, 50, 7, 42, 2, 2, 50, 14, 3,
	2, 2, 2, 51, 52, 7, 43, 2, 2, 52, 16, 3, 2, 2, 2, 53, 54, 7, 63, 2, 2,
	54, 18, 3, 2, 2, 2, 55, 56, 7, 35, 2, 2, 56, 20, 3, 2, 2, 2, 57, 58, 7,
	72, 2, 2, 58, 59, 7, 113, 2, 2, 59, 60, 7, 116, 2, 2, 60, 61, 7, 99, 2,
	2, 61, 62, 7, 110, 2, 2, 62, 63, 7, 110, 2, 2, 63, 22, 3, 2, 2, 2, 64,
	65, 7, 71, 2, 2, 65, 66, 7, 122, 2, 2, 66, 67, 7, 107, 2, 2, 67, 68, 7,
	117, 2, 2, 68, 69, 7, 118, 2, 2, 69, 70, 7, 117, 2, 2, 70, 24, 3, 2, 2,
	2, 71, 72, 9, 2, 2, 2, 72, 26, 3, 2, 2, 2, 73, 74, 7, 94, 2, 2, 74, 75,
	7, 49, 2, 2, 75, 28, 3, 2, 2, 2, 76, 77, 7, 96, 2, 2, 77, 30, 3, 2, 2,
	2, 78, 79, 7, 47, 2, 2, 79, 80, 7, 64, 2, 2, 80, 32, 3, 2, 2, 2, 81, 82,
	7, 62, 2, 2, 82, 83, 7, 47, 2, 2, 83, 84, 7, 64, 2, 2, 84, 34, 3, 2, 2,
	2, 85, 87, 9, 3, 2, 2, 86, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 86,
	3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 36, 3, 2, 2, 2, 90, 92, 9, 4, 2, 2,
	91, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3,
	2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 96, 8, 19, 2, 2, 96, 38, 3, 2, 2, 2, 5,
	2, 88, 93, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'?'", "'_'", "'#'", "'.'", "','", "'('", "')'", "'='", "'!'", "'Forall'",
	"'Exists'", "", "'\\/'", "'^'", "'->'", "'<->'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "LPAREN", "RPAREN", "EQUAL", "NOT", "FORALL", "EXISTS",
	"CHARACTER", "CONJ", "DISJ", "IMPL", "BICOND", "ENDLINE", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "LPAREN", "RPAREN", "EQUAL", "NOT",
	"FORALL", "EXISTS", "CHARACTER", "CONJ", "DISJ", "IMPL", "BICOND", "ENDLINE",
	"WHITESPACE",
}

type folLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewfolLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *folLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewfolLexer(input antlr.CharStream) *folLexer {
	l := new(folLexer)
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
	l.GrammarFileName = "fol.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// folLexer tokens.
const (
	folLexerT__0       = 1
	folLexerT__1       = 2
	folLexerT__2       = 3
	folLexerT__3       = 4
	folLexerT__4       = 5
	folLexerLPAREN     = 6
	folLexerRPAREN     = 7
	folLexerEQUAL      = 8
	folLexerNOT        = 9
	folLexerFORALL     = 10
	folLexerEXISTS     = 11
	folLexerCHARACTER  = 12
	folLexerCONJ       = 13
	folLexerDISJ       = 14
	folLexerIMPL       = 15
	folLexerBICOND     = 16
	folLexerENDLINE    = 17
	folLexerWHITESPACE = 18
)
