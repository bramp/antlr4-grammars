// Code generated from sickbay.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sickbay

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 26, 164,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 7, 21, 139, 10, 21, 12, 21, 14, 21, 142, 11,
	21, 3, 22, 3, 22, 7, 22, 146, 10, 22, 12, 22, 14, 22, 149, 11, 22, 3, 22,
	3, 22, 3, 23, 7, 23, 154, 10, 23, 12, 23, 14, 23, 157, 11, 23, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 25, 3, 25, 2, 2, 26, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49,
	26, 3, 2, 10, 5, 2, 44, 45, 47, 47, 49, 49, 3, 2, 67, 92, 5, 2, 39, 39,
	50, 59, 67, 92, 3, 2, 50, 59, 3, 2, 36, 36, 3, 2, 12, 12, 4, 2, 12, 12,
	15, 15, 5, 2, 11, 12, 15, 15, 34, 34, 2, 166, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 3, 51, 3,
	2, 2, 2, 5, 53, 3, 2, 2, 2, 7, 57, 3, 2, 2, 2, 9, 61, 3, 2, 2, 2, 11, 63,
	3, 2, 2, 2, 13, 68, 3, 2, 2, 2, 15, 74, 3, 2, 2, 2, 17, 81, 3, 2, 2, 2,
	19, 85, 3, 2, 2, 2, 21, 93, 3, 2, 2, 2, 23, 102, 3, 2, 2, 2, 25, 106, 3,
	2, 2, 2, 27, 111, 3, 2, 2, 2, 29, 113, 3, 2, 2, 2, 31, 115, 3, 2, 2, 2,
	33, 121, 3, 2, 2, 2, 35, 126, 3, 2, 2, 2, 37, 131, 3, 2, 2, 2, 39, 133,
	3, 2, 2, 2, 41, 136, 3, 2, 2, 2, 43, 143, 3, 2, 2, 2, 45, 155, 3, 2, 2,
	2, 47, 158, 3, 2, 2, 2, 49, 160, 3, 2, 2, 2, 51, 52, 7, 60, 2, 2, 52, 4,
	3, 2, 2, 2, 53, 54, 7, 84, 2, 2, 54, 55, 7, 71, 2, 2, 55, 56, 7, 79, 2,
	2, 56, 6, 3, 2, 2, 2, 57, 58, 7, 78, 2, 2, 58, 59, 7, 71, 2, 2, 59, 60,
	7, 86, 2, 2, 60, 8, 3, 2, 2, 2, 61, 62, 7, 63, 2, 2, 62, 10, 3, 2, 2, 2,
	63, 64, 7, 73, 2, 2, 64, 65, 7, 81, 2, 2, 65, 66, 7, 86, 2, 2, 66, 67,
	7, 81, 2, 2, 67, 12, 3, 2, 2, 2, 68, 69, 7, 73, 2, 2, 69, 70, 7, 81, 2,
	2, 70, 71, 7, 85, 2, 2, 71, 72, 7, 87, 2, 2, 72, 73, 7, 68, 2, 2, 73, 14,
	3, 2, 2, 2, 74, 75, 7, 84, 2, 2, 75, 76, 7, 71, 2, 2, 76, 77, 7, 86, 2,
	2, 77, 78, 7, 87, 2, 2, 78, 79, 7, 84, 2, 2, 79, 80, 7, 80, 2, 2, 80, 16,
	3, 2, 2, 2, 81, 82, 7, 71, 2, 2, 82, 83, 7, 80, 2, 2, 83, 84, 7, 70, 2,
	2, 84, 18, 3, 2, 2, 2, 85, 86, 7, 82, 2, 2, 86, 87, 7, 84, 2, 2, 87, 88,
	7, 81, 2, 2, 88, 89, 7, 78, 2, 2, 89, 90, 7, 81, 2, 2, 90, 91, 7, 80, 2,
	2, 91, 92, 7, 73, 2, 2, 92, 20, 3, 2, 2, 2, 93, 94, 7, 69, 2, 2, 94, 95,
	7, 87, 2, 2, 95, 96, 7, 86, 2, 2, 96, 97, 7, 85, 2, 2, 97, 98, 7, 74, 2,
	2, 98, 99, 7, 81, 2, 2, 99, 100, 7, 84, 2, 2, 100, 101, 7, 86, 2, 2, 101,
	22, 3, 2, 2, 2, 102, 103, 7, 70, 2, 2, 103, 104, 7, 75, 2, 2, 104, 105,
	7, 79, 2, 2, 105, 24, 3, 2, 2, 2, 106, 107, 7, 84, 2, 2, 107, 108, 7, 75,
	2, 2, 108, 109, 7, 80, 2, 2, 109, 110, 7, 73, 2, 2, 110, 26, 3, 2, 2, 2,
	111, 112, 7, 42, 2, 2, 112, 28, 3, 2, 2, 2, 113, 114, 7, 43, 2, 2, 114,
	30, 3, 2, 2, 2, 115, 116, 7, 75, 2, 2, 116, 117, 7, 80, 2, 2, 117, 118,
	7, 82, 2, 2, 118, 119, 7, 87, 2, 2, 119, 120, 7, 86, 2, 2, 120, 32, 3,
	2, 2, 2, 121, 122, 7, 69, 2, 2, 122, 123, 7, 74, 2, 2, 123, 124, 7, 84,
	2, 2, 124, 125, 7, 38, 2, 2, 125, 34, 3, 2, 2, 2, 126, 127, 7, 84, 2, 2,
	127, 128, 7, 80, 2, 2, 128, 129, 7, 70, 2, 2, 129, 130, 7, 39, 2, 2, 130,
	36, 3, 2, 2, 2, 131, 132, 9, 2, 2, 2, 132, 38, 3, 2, 2, 2, 133, 134, 9,
	3, 2, 2, 134, 135, 9, 4, 2, 2, 135, 40, 3, 2, 2, 2, 136, 140, 9, 5, 2,
	2, 137, 139, 9, 5, 2, 2, 138, 137, 3, 2, 2, 2, 139, 142, 3, 2, 2, 2, 140,
	138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 42, 3, 2, 2, 2, 142, 140, 3,
	2, 2, 2, 143, 147, 7, 36, 2, 2, 144, 146, 10, 6, 2, 2, 145, 144, 3, 2,
	2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2,
	148, 150, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 151, 7, 36, 2, 2, 151,
	44, 3, 2, 2, 2, 152, 154, 10, 7, 2, 2, 153, 152, 3, 2, 2, 2, 154, 157,
	3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 46, 3, 2,
	2, 2, 157, 155, 3, 2, 2, 2, 158, 159, 9, 8, 2, 2, 159, 48, 3, 2, 2, 2,
	160, 161, 9, 9, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 8, 25, 2, 2, 163,
	50, 3, 2, 2, 2, 6, 2, 140, 147, 155, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "':'", "'REM'", "'LET'", "'='", "'GOTO'", "'GOSUB'", "'RETURN'", "'END'",
	"'PROLONG'", "'CUTSHORT'", "'DIM'", "'RING'", "'('", "')'", "'INPUT'",
	"'CHR$'", "'RND%'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"INTOP", "IINTID", "INTCONST", "STRCONST", "ARBTEXT", "NL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"INTOP", "IINTID", "INTCONST", "STRCONST", "ARBTEXT", "NL", "WS",
}

type sickbayLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewsickbayLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *sickbayLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewsickbayLexer(input antlr.CharStream) *sickbayLexer {
	l := new(sickbayLexer)
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
	l.GrammarFileName = "sickbay.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// sickbayLexer tokens.
const (
	sickbayLexerT__0     = 1
	sickbayLexerT__1     = 2
	sickbayLexerT__2     = 3
	sickbayLexerT__3     = 4
	sickbayLexerT__4     = 5
	sickbayLexerT__5     = 6
	sickbayLexerT__6     = 7
	sickbayLexerT__7     = 8
	sickbayLexerT__8     = 9
	sickbayLexerT__9     = 10
	sickbayLexerT__10    = 11
	sickbayLexerT__11    = 12
	sickbayLexerT__12    = 13
	sickbayLexerT__13    = 14
	sickbayLexerT__14    = 15
	sickbayLexerT__15    = 16
	sickbayLexerT__16    = 17
	sickbayLexerINTOP    = 18
	sickbayLexerIINTID   = 19
	sickbayLexerINTCONST = 20
	sickbayLexerSTRCONST = 21
	sickbayLexerARBTEXT  = 22
	sickbayLexerNL       = 23
	sickbayLexerWS       = 24
)
