// Code generated from dgol.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dgol

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 25, 148,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 7, 22,
	135, 10, 22, 12, 22, 14, 22, 138, 11, 22, 3, 23, 6, 23, 141, 10, 23, 13,
	23, 14, 23, 142, 3, 24, 3, 24, 3, 24, 3, 24, 2, 2, 25, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 3, 2, 6, 4, 2, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124,
	4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 15, 15, 34, 34, 2, 149, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 3, 49, 3, 2, 2,
	2, 5, 53, 3, 2, 2, 2, 7, 64, 3, 2, 2, 2, 9, 66, 3, 2, 2, 2, 11, 68, 3,
	2, 2, 2, 13, 70, 3, 2, 2, 2, 15, 74, 3, 2, 2, 2, 17, 82, 3, 2, 2, 2, 19,
	90, 3, 2, 2, 2, 21, 92, 3, 2, 2, 2, 23, 96, 3, 2, 2, 2, 25, 98, 3, 2, 2,
	2, 27, 100, 3, 2, 2, 2, 29, 102, 3, 2, 2, 2, 31, 107, 3, 2, 2, 2, 33, 110,
	3, 2, 2, 2, 35, 113, 3, 2, 2, 2, 37, 118, 3, 2, 2, 2, 39, 120, 3, 2, 2,
	2, 41, 127, 3, 2, 2, 2, 43, 132, 3, 2, 2, 2, 45, 140, 3, 2, 2, 2, 47, 144,
	3, 2, 2, 2, 49, 50, 7, 87, 2, 2, 50, 51, 7, 85, 2, 2, 51, 52, 7, 71, 2,
	2, 52, 4, 3, 2, 2, 2, 53, 54, 7, 85, 2, 2, 54, 55, 7, 87, 2, 2, 55, 56,
	7, 68, 2, 2, 56, 57, 7, 84, 2, 2, 57, 58, 7, 81, 2, 2, 58, 59, 7, 87, 2,
	2, 59, 60, 7, 86, 2, 2, 60, 61, 7, 75, 2, 2, 61, 62, 7, 80, 2, 2, 62, 63,
	7, 71, 2, 2, 63, 6, 3, 2, 2, 2, 64, 65, 7, 42, 2, 2, 65, 8, 3, 2, 2, 2,
	66, 67, 7, 46, 2, 2, 67, 10, 3, 2, 2, 2, 68, 69, 7, 43, 2, 2, 69, 12, 3,
	2, 2, 2, 70, 71, 7, 71, 2, 2, 71, 72, 7, 80, 2, 2, 72, 73, 7, 70, 2, 2,
	73, 14, 3, 2, 2, 2, 74, 75, 7, 82, 2, 2, 75, 76, 7, 84, 2, 2, 76, 77, 7,
	81, 2, 2, 77, 78, 7, 73, 2, 2, 78, 79, 7, 84, 2, 2, 79, 80, 7, 67, 2, 2,
	80, 81, 7, 79, 2, 2, 81, 16, 3, 2, 2, 2, 82, 83, 7, 78, 2, 2, 83, 84, 7,
	75, 2, 2, 84, 85, 7, 68, 2, 2, 85, 86, 7, 84, 2, 2, 86, 87, 7, 67, 2, 2,
	87, 88, 7, 84, 2, 2, 88, 89, 7, 91, 2, 2, 89, 18, 3, 2, 2, 2, 90, 91, 7,
	50, 2, 2, 91, 20, 3, 2, 2, 2, 92, 93, 7, 78, 2, 2, 93, 94, 7, 71, 2, 2,
	94, 95, 7, 86, 2, 2, 95, 22, 3, 2, 2, 2, 96, 97, 7, 63, 2, 2, 97, 24, 3,
	2, 2, 2, 98, 99, 7, 62, 2, 2, 99, 26, 3, 2, 2, 2, 100, 101, 7, 64, 2, 2,
	101, 28, 3, 2, 2, 2, 102, 103, 7, 71, 2, 2, 103, 104, 7, 78, 2, 2, 104,
	105, 7, 85, 2, 2, 105, 106, 7, 71, 2, 2, 106, 30, 3, 2, 2, 2, 107, 108,
	7, 75, 2, 2, 108, 109, 7, 72, 2, 2, 109, 32, 3, 2, 2, 2, 110, 111, 7, 70,
	2, 2, 111, 112, 7, 81, 2, 2, 112, 34, 3, 2, 2, 2, 113, 114, 7, 69, 2, 2,
	114, 115, 7, 67, 2, 2, 115, 116, 7, 78, 2, 2, 116, 117, 7, 78, 2, 2, 117,
	36, 3, 2, 2, 2, 118, 119, 7, 48, 2, 2, 119, 38, 3, 2, 2, 2, 120, 121, 7,
	84, 2, 2, 121, 122, 7, 71, 2, 2, 122, 123, 7, 86, 2, 2, 123, 124, 7, 87,
	2, 2, 124, 125, 7, 84, 2, 2, 125, 126, 7, 80, 2, 2, 126, 40, 3, 2, 2, 2,
	127, 128, 7, 71, 2, 2, 128, 129, 7, 90, 2, 2, 129, 130, 7, 75, 2, 2, 130,
	131, 7, 86, 2, 2, 131, 42, 3, 2, 2, 2, 132, 136, 9, 2, 2, 2, 133, 135,
	9, 3, 2, 2, 134, 133, 3, 2, 2, 2, 135, 138, 3, 2, 2, 2, 136, 134, 3, 2,
	2, 2, 136, 137, 3, 2, 2, 2, 137, 44, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2,
	139, 141, 9, 4, 2, 2, 140, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142,
	140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 46, 3, 2, 2, 2, 144, 145, 9,
	5, 2, 2, 145, 146, 3, 2, 2, 2, 146, 147, 8, 24, 2, 2, 147, 48, 3, 2, 2,
	2, 5, 2, 136, 142, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'USE'", "'SUBROUTINE'", "'('", "','", "')'", "'END'", "'PROGRAM'",
	"'LIBRARY'", "'0'", "'LET'", "'='", "'<'", "'>'", "'ELSE'", "'IF'", "'DO'",
	"'CALL'", "'.'", "'RETURN'", "'EXIT'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "IDENTIFER", "NL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "IDENTIFER", "NL", "WS",
}

type dgolLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewdgolLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *dgolLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdgolLexer(input antlr.CharStream) *dgolLexer {
	l := new(dgolLexer)
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
	l.GrammarFileName = "dgol.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// dgolLexer tokens.
const (
	dgolLexerT__0      = 1
	dgolLexerT__1      = 2
	dgolLexerT__2      = 3
	dgolLexerT__3      = 4
	dgolLexerT__4      = 5
	dgolLexerT__5      = 6
	dgolLexerT__6      = 7
	dgolLexerT__7      = 8
	dgolLexerT__8      = 9
	dgolLexerT__9      = 10
	dgolLexerT__10     = 11
	dgolLexerT__11     = 12
	dgolLexerT__12     = 13
	dgolLexerT__13     = 14
	dgolLexerT__14     = 15
	dgolLexerT__15     = 16
	dgolLexerT__16     = 17
	dgolLexerT__17     = 18
	dgolLexerT__18     = 19
	dgolLexerT__19     = 20
	dgolLexerIDENTIFER = 21
	dgolLexerNL        = 22
	dgolLexerWS        = 23
)
