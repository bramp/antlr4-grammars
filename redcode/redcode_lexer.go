// Code generated from redcode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package redcode

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 38, 194,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	34, 6, 34, 175, 10, 34, 13, 34, 14, 34, 176, 3, 35, 3, 35, 7, 35, 181,
	10, 35, 12, 35, 14, 35, 184, 11, 35, 3, 36, 6, 36, 187, 10, 36, 13, 36,
	14, 36, 188, 3, 37, 3, 37, 3, 37, 3, 37, 2, 2, 38, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15,
	29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24,
	47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33,
	65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 3, 2, 5, 3, 2, 50, 59, 4, 2, 12,
	12, 15, 15, 4, 2, 11, 11, 34, 34, 2, 196, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2,
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2,
	2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2,
	2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3,
	2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67,
	3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 3,
	75, 3, 2, 2, 2, 5, 77, 3, 2, 2, 2, 7, 79, 3, 2, 2, 2, 9, 81, 3, 2, 2, 2,
	11, 83, 3, 2, 2, 2, 13, 85, 3, 2, 2, 2, 15, 87, 3, 2, 2, 2, 17, 89, 3,
	2, 2, 2, 19, 91, 3, 2, 2, 2, 21, 93, 3, 2, 2, 2, 23, 95, 3, 2, 2, 2, 25,
	97, 3, 2, 2, 2, 27, 100, 3, 2, 2, 2, 29, 103, 3, 2, 2, 2, 31, 105, 3, 2,
	2, 2, 33, 107, 3, 2, 2, 2, 35, 109, 3, 2, 2, 2, 37, 113, 3, 2, 2, 2, 39,
	117, 3, 2, 2, 2, 41, 121, 3, 2, 2, 2, 43, 125, 3, 2, 2, 2, 45, 129, 3,
	2, 2, 2, 47, 133, 3, 2, 2, 2, 49, 137, 3, 2, 2, 2, 51, 141, 3, 2, 2, 2,
	53, 145, 3, 2, 2, 2, 55, 149, 3, 2, 2, 2, 57, 153, 3, 2, 2, 2, 59, 157,
	3, 2, 2, 2, 61, 161, 3, 2, 2, 2, 63, 165, 3, 2, 2, 2, 65, 169, 3, 2, 2,
	2, 67, 174, 3, 2, 2, 2, 69, 178, 3, 2, 2, 2, 71, 186, 3, 2, 2, 2, 73, 190,
	3, 2, 2, 2, 75, 76, 7, 48, 2, 2, 76, 4, 3, 2, 2, 2, 77, 78, 7, 46, 2, 2,
	78, 6, 3, 2, 2, 2, 79, 80, 7, 37, 2, 2, 80, 8, 3, 2, 2, 2, 81, 82, 7, 38,
	2, 2, 82, 10, 3, 2, 2, 2, 83, 84, 7, 66, 2, 2, 84, 12, 3, 2, 2, 2, 85,
	86, 7, 62, 2, 2, 86, 14, 3, 2, 2, 2, 87, 88, 7, 64, 2, 2, 88, 16, 3, 2,
	2, 2, 89, 90, 7, 45, 2, 2, 90, 18, 3, 2, 2, 2, 91, 92, 7, 47, 2, 2, 92,
	20, 3, 2, 2, 2, 93, 94, 7, 67, 2, 2, 94, 22, 3, 2, 2, 2, 95, 96, 7, 68,
	2, 2, 96, 24, 3, 2, 2, 2, 97, 98, 7, 67, 2, 2, 98, 99, 7, 68, 2, 2, 99,
	26, 3, 2, 2, 2, 100, 101, 7, 68, 2, 2, 101, 102, 7, 67, 2, 2, 102, 28,
	3, 2, 2, 2, 103, 104, 7, 72, 2, 2, 104, 30, 3, 2, 2, 2, 105, 106, 7, 90,
	2, 2, 106, 32, 3, 2, 2, 2, 107, 108, 7, 75, 2, 2, 108, 34, 3, 2, 2, 2,
	109, 110, 7, 70, 2, 2, 110, 111, 7, 67, 2, 2, 111, 112, 7, 86, 2, 2, 112,
	36, 3, 2, 2, 2, 113, 114, 7, 79, 2, 2, 114, 115, 7, 81, 2, 2, 115, 116,
	7, 88, 2, 2, 116, 38, 3, 2, 2, 2, 117, 118, 7, 67, 2, 2, 118, 119, 7, 70,
	2, 2, 119, 120, 7, 70, 2, 2, 120, 40, 3, 2, 2, 2, 121, 122, 7, 85, 2, 2,
	122, 123, 7, 87, 2, 2, 123, 124, 7, 68, 2, 2, 124, 42, 3, 2, 2, 2, 125,
	126, 7, 79, 2, 2, 126, 127, 7, 87, 2, 2, 127, 128, 7, 78, 2, 2, 128, 44,
	3, 2, 2, 2, 129, 130, 7, 70, 2, 2, 130, 131, 7, 75, 2, 2, 131, 132, 7,
	88, 2, 2, 132, 46, 3, 2, 2, 2, 133, 134, 7, 79, 2, 2, 134, 135, 7, 81,
	2, 2, 135, 136, 7, 70, 2, 2, 136, 48, 3, 2, 2, 2, 137, 138, 7, 76, 2, 2,
	138, 139, 7, 79, 2, 2, 139, 140, 7, 82, 2, 2, 140, 50, 3, 2, 2, 2, 141,
	142, 7, 76, 2, 2, 142, 143, 7, 79, 2, 2, 143, 144, 7, 92, 2, 2, 144, 52,
	3, 2, 2, 2, 145, 146, 7, 76, 2, 2, 146, 147, 7, 79, 2, 2, 147, 148, 7,
	80, 2, 2, 148, 54, 3, 2, 2, 2, 149, 150, 7, 70, 2, 2, 150, 151, 7, 76,
	2, 2, 151, 152, 7, 80, 2, 2, 152, 56, 3, 2, 2, 2, 153, 154, 7, 69, 2, 2,
	154, 155, 7, 79, 2, 2, 155, 156, 7, 82, 2, 2, 156, 58, 3, 2, 2, 2, 157,
	158, 7, 85, 2, 2, 158, 159, 7, 78, 2, 2, 159, 160, 7, 86, 2, 2, 160, 60,
	3, 2, 2, 2, 161, 162, 7, 70, 2, 2, 162, 163, 7, 76, 2, 2, 163, 164, 7,
	92, 2, 2, 164, 62, 3, 2, 2, 2, 165, 166, 7, 85, 2, 2, 166, 167, 7, 82,
	2, 2, 167, 168, 7, 78, 2, 2, 168, 64, 3, 2, 2, 2, 169, 170, 7, 81, 2, 2,
	170, 171, 7, 84, 2, 2, 171, 172, 7, 73, 2, 2, 172, 66, 3, 2, 2, 2, 173,
	175, 9, 2, 2, 2, 174, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 174,
	3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 68, 3, 2, 2, 2, 178, 182, 7, 61,
	2, 2, 179, 181, 10, 3, 2, 2, 180, 179, 3, 2, 2, 2, 181, 184, 3, 2, 2, 2,
	182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 70, 3, 2, 2, 2, 184, 182,
	3, 2, 2, 2, 185, 187, 9, 3, 2, 2, 186, 185, 3, 2, 2, 2, 187, 188, 3, 2,
	2, 2, 188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 72, 3, 2, 2, 2,
	190, 191, 9, 4, 2, 2, 191, 192, 3, 2, 2, 2, 192, 193, 8, 37, 2, 2, 193,
	74, 3, 2, 2, 2, 6, 2, 176, 182, 188, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'.'", "','", "'#'", "'$'", "'@'", "'<'", "'>'", "'+'", "'-'", "'A'",
	"'B'", "'AB'", "'BA'", "'F'", "'X'", "'I'", "'DAT'", "'MOV'", "'ADD'",
	"'SUB'", "'MUL'", "'DIV'", "'MOD'", "'JMP'", "'JMZ'", "'JMN'", "'DJN'",
	"'CMP'", "'SLT'", "'DJZ'", "'SPL'", "'ORG'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "A", "B", "AB", "BA", "F", "X",
	"I", "DAT", "MOV", "ADD", "SUB", "MUL", "DIV", "MOD", "JMP", "JMZ", "JMN",
	"DJN", "CMP", "SLT", "DJZ", "SPL", "ORG", "NUMBER", "COMMENT", "EOL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"A", "B", "AB", "BA", "F", "X", "I", "DAT", "MOV", "ADD", "SUB", "MUL",
	"DIV", "MOD", "JMP", "JMZ", "JMN", "DJN", "CMP", "SLT", "DJZ", "SPL", "ORG",
	"NUMBER", "COMMENT", "EOL", "WS",
}

type redcodeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewredcodeLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *redcodeLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewredcodeLexer(input antlr.CharStream) *redcodeLexer {
	l := new(redcodeLexer)
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
	l.GrammarFileName = "redcode.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// redcodeLexer tokens.
const (
	redcodeLexerT__0    = 1
	redcodeLexerT__1    = 2
	redcodeLexerT__2    = 3
	redcodeLexerT__3    = 4
	redcodeLexerT__4    = 5
	redcodeLexerT__5    = 6
	redcodeLexerT__6    = 7
	redcodeLexerT__7    = 8
	redcodeLexerT__8    = 9
	redcodeLexerA       = 10
	redcodeLexerB       = 11
	redcodeLexerAB      = 12
	redcodeLexerBA      = 13
	redcodeLexerF       = 14
	redcodeLexerX       = 15
	redcodeLexerI       = 16
	redcodeLexerDAT     = 17
	redcodeLexerMOV     = 18
	redcodeLexerADD     = 19
	redcodeLexerSUB     = 20
	redcodeLexerMUL     = 21
	redcodeLexerDIV     = 22
	redcodeLexerMOD     = 23
	redcodeLexerJMP     = 24
	redcodeLexerJMZ     = 25
	redcodeLexerJMN     = 26
	redcodeLexerDJN     = 27
	redcodeLexerCMP     = 28
	redcodeLexerSLT     = 29
	redcodeLexerDJZ     = 30
	redcodeLexerSPL     = 31
	redcodeLexerORG     = 32
	redcodeLexerNUMBER  = 33
	redcodeLexerCOMMENT = 34
	redcodeLexerEOL     = 35
	redcodeLexerWS      = 36
)
