// Code generated from janus.g4 by ANTLR 4.9.3. DO NOT EDIT.

package janus

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 31, 189,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 163, 10, 26, 3,
	27, 3, 27, 7, 27, 167, 10, 27, 12, 27, 14, 27, 170, 11, 27, 3, 28, 3, 28,
	3, 29, 3, 29, 7, 29, 176, 10, 29, 12, 29, 14, 29, 179, 11, 29, 3, 29, 3,
	29, 3, 30, 6, 30, 184, 10, 30, 13, 30, 14, 30, 185, 3, 30, 3, 30, 2, 2,
	31, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12,
	23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21,
	41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30,
	59, 31, 3, 2, 9, 9, 2, 35, 35, 37, 37, 40, 40, 45, 45, 47, 47, 62, 64,
	126, 126, 5, 2, 44, 44, 49, 49, 94, 94, 4, 2, 67, 92, 99, 124, 5, 2, 50,
	59, 67, 92, 99, 124, 3, 2, 50, 59, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12,
	15, 15, 34, 34, 2, 194, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2,
	31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2,
	2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2,
	2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2,
	2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 3, 61, 3,
	2, 2, 2, 5, 63, 3, 2, 2, 2, 7, 65, 3, 2, 2, 2, 9, 75, 3, 2, 2, 2, 11, 78,
	3, 2, 2, 2, 13, 83, 3, 2, 2, 2, 15, 88, 3, 2, 2, 2, 17, 91, 3, 2, 2, 2,
	19, 96, 3, 2, 2, 2, 21, 99, 3, 2, 2, 2, 23, 104, 3, 2, 2, 2, 25, 110, 3,
	2, 2, 2, 27, 115, 3, 2, 2, 2, 29, 122, 3, 2, 2, 2, 31, 127, 3, 2, 2, 2,
	33, 133, 3, 2, 2, 2, 35, 136, 3, 2, 2, 2, 37, 139, 3, 2, 2, 2, 39, 142,
	3, 2, 2, 2, 41, 146, 3, 2, 2, 2, 43, 148, 3, 2, 2, 2, 45, 150, 3, 2, 2,
	2, 47, 152, 3, 2, 2, 2, 49, 154, 3, 2, 2, 2, 51, 162, 3, 2, 2, 2, 53, 164,
	3, 2, 2, 2, 55, 171, 3, 2, 2, 2, 57, 173, 3, 2, 2, 2, 59, 183, 3, 2, 2,
	2, 61, 62, 7, 93, 2, 2, 62, 4, 3, 2, 2, 2, 63, 64, 7, 95, 2, 2, 64, 6,
	3, 2, 2, 2, 65, 66, 7, 82, 2, 2, 66, 67, 7, 84, 2, 2, 67, 68, 7, 81, 2,
	2, 68, 69, 7, 69, 2, 2, 69, 70, 7, 71, 2, 2, 70, 71, 7, 70, 2, 2, 71, 72,
	7, 87, 2, 2, 72, 73, 7, 84, 2, 2, 73, 74, 7, 71, 2, 2, 74, 8, 3, 2, 2,
	2, 75, 76, 7, 75, 2, 2, 76, 77, 7, 72, 2, 2, 77, 10, 3, 2, 2, 2, 78, 79,
	7, 86, 2, 2, 79, 80, 7, 74, 2, 2, 80, 81, 7, 71, 2, 2, 81, 82, 7, 80, 2,
	2, 82, 12, 3, 2, 2, 2, 83, 84, 7, 71, 2, 2, 84, 85, 7, 78, 2, 2, 85, 86,
	7, 85, 2, 2, 86, 87, 7, 71, 2, 2, 87, 14, 3, 2, 2, 2, 88, 89, 7, 72, 2,
	2, 89, 90, 7, 75, 2, 2, 90, 16, 3, 2, 2, 2, 91, 92, 7, 72, 2, 2, 92, 93,
	7, 84, 2, 2, 93, 94, 7, 81, 2, 2, 94, 95, 7, 79, 2, 2, 95, 18, 3, 2, 2,
	2, 96, 97, 7, 70, 2, 2, 97, 98, 7, 81, 2, 2, 98, 20, 3, 2, 2, 2, 99, 100,
	7, 78, 2, 2, 100, 101, 7, 81, 2, 2, 101, 102, 7, 81, 2, 2, 102, 103, 7,
	82, 2, 2, 103, 22, 3, 2, 2, 2, 104, 105, 7, 87, 2, 2, 105, 106, 7, 80,
	2, 2, 106, 107, 7, 86, 2, 2, 107, 108, 7, 75, 2, 2, 108, 109, 7, 78, 2,
	2, 109, 24, 3, 2, 2, 2, 110, 111, 7, 69, 2, 2, 111, 112, 7, 67, 2, 2, 112,
	113, 7, 78, 2, 2, 113, 114, 7, 78, 2, 2, 114, 26, 3, 2, 2, 2, 115, 116,
	7, 87, 2, 2, 116, 117, 7, 80, 2, 2, 117, 118, 7, 69, 2, 2, 118, 119, 7,
	67, 2, 2, 119, 120, 7, 78, 2, 2, 120, 121, 7, 78, 2, 2, 121, 28, 3, 2,
	2, 2, 122, 123, 7, 84, 2, 2, 123, 124, 7, 71, 2, 2, 124, 125, 7, 67, 2,
	2, 125, 126, 7, 70, 2, 2, 126, 30, 3, 2, 2, 2, 127, 128, 7, 89, 2, 2, 128,
	129, 7, 84, 2, 2, 129, 130, 7, 75, 2, 2, 130, 131, 7, 86, 2, 2, 131, 132,
	7, 71, 2, 2, 132, 32, 3, 2, 2, 2, 133, 134, 7, 45, 2, 2, 134, 135, 7, 63,
	2, 2, 135, 34, 3, 2, 2, 2, 136, 137, 7, 47, 2, 2, 137, 138, 7, 63, 2, 2,
	138, 36, 3, 2, 2, 2, 139, 140, 7, 35, 2, 2, 140, 141, 7, 63, 2, 2, 141,
	38, 3, 2, 2, 2, 142, 143, 7, 62, 2, 2, 143, 144, 7, 63, 2, 2, 144, 145,
	7, 64, 2, 2, 145, 40, 3, 2, 2, 2, 146, 147, 7, 60, 2, 2, 147, 42, 3, 2,
	2, 2, 148, 149, 7, 42, 2, 2, 149, 44, 3, 2, 2, 2, 150, 151, 7, 43, 2, 2,
	151, 46, 3, 2, 2, 2, 152, 153, 7, 47, 2, 2, 153, 48, 3, 2, 2, 2, 154, 155,
	7, 128, 2, 2, 155, 50, 3, 2, 2, 2, 156, 163, 9, 2, 2, 2, 157, 158, 7, 62,
	2, 2, 158, 163, 7, 63, 2, 2, 159, 160, 7, 64, 2, 2, 160, 163, 7, 63, 2,
	2, 161, 163, 9, 3, 2, 2, 162, 156, 3, 2, 2, 2, 162, 157, 3, 2, 2, 2, 162,
	159, 3, 2, 2, 2, 162, 161, 3, 2, 2, 2, 163, 52, 3, 2, 2, 2, 164, 168, 9,
	4, 2, 2, 165, 167, 9, 5, 2, 2, 166, 165, 3, 2, 2, 2, 167, 170, 3, 2, 2,
	2, 168, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 54, 3, 2, 2, 2, 170,
	168, 3, 2, 2, 2, 171, 172, 9, 6, 2, 2, 172, 56, 3, 2, 2, 2, 173, 177, 7,
	61, 2, 2, 174, 176, 10, 7, 2, 2, 175, 174, 3, 2, 2, 2, 176, 179, 3, 2,
	2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 180, 3, 2, 2, 2,
	179, 177, 3, 2, 2, 2, 180, 181, 8, 29, 2, 2, 181, 58, 3, 2, 2, 2, 182,
	184, 9, 8, 2, 2, 183, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 183,
	3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 188, 8, 30,
	2, 2, 188, 60, 3, 2, 2, 2, 7, 2, 162, 168, 177, 185, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "']'", "'PROCEDURE'", "'IF'", "'THEN'", "'ELSE'", "'FI'", "'FROM'",
	"'DO'", "'LOOP'", "'UNTIL'", "'CALL'", "'UNCALL'", "'READ'", "'WRITE'",
	"'+='", "'-='", "'!='", "'<=>'", "':'", "'('", "')'", "'-'", "'~'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "BINOP", "IDENT", "NUM", "COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "BINOP",
	"IDENT", "NUM", "COMMENT", "WS",
}

type janusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewjanusLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *janusLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewjanusLexer(input antlr.CharStream) *janusLexer {
	l := new(janusLexer)
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
	l.GrammarFileName = "janus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// janusLexer tokens.
const (
	janusLexerT__0    = 1
	janusLexerT__1    = 2
	janusLexerT__2    = 3
	janusLexerT__3    = 4
	janusLexerT__4    = 5
	janusLexerT__5    = 6
	janusLexerT__6    = 7
	janusLexerT__7    = 8
	janusLexerT__8    = 9
	janusLexerT__9    = 10
	janusLexerT__10   = 11
	janusLexerT__11   = 12
	janusLexerT__12   = 13
	janusLexerT__13   = 14
	janusLexerT__14   = 15
	janusLexerT__15   = 16
	janusLexerT__16   = 17
	janusLexerT__17   = 18
	janusLexerT__18   = 19
	janusLexerT__19   = 20
	janusLexerT__20   = 21
	janusLexerT__21   = 22
	janusLexerT__22   = 23
	janusLexerT__23   = 24
	janusLexerBINOP   = 25
	janusLexerIDENT   = 26
	janusLexerNUM     = 27
	janusLexerCOMMENT = 28
	janusLexerWS      = 29
)
