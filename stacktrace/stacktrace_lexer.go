// Code generated from StackTrace.g4 by ANTLR 4.9.3. DO NOT EDIT.

package stacktrace

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 148,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	6, 6, 59, 10, 6, 13, 6, 14, 6, 60, 3, 7, 6, 7, 64, 10, 7, 13, 7, 14, 7,
	65, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 72, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 2, 2, 23, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 2, 17, 9, 19, 10, 21, 11, 23, 12, 25,
	13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
	22, 3, 2, 3, 5, 2, 11, 12, 14, 15, 34, 34, 2, 151, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2,
	2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 3, 45, 3,
	2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 49, 3, 2, 2, 2, 9, 51, 3, 2, 2, 2, 11, 58,
	3, 2, 2, 2, 13, 63, 3, 2, 2, 2, 15, 71, 3, 2, 2, 2, 17, 73, 3, 2, 2, 2,
	19, 75, 3, 2, 2, 2, 21, 78, 3, 2, 2, 2, 23, 89, 3, 2, 2, 2, 25, 94, 3,
	2, 2, 2, 27, 98, 3, 2, 2, 2, 29, 100, 3, 2, 2, 2, 31, 114, 3, 2, 2, 2,
	33, 129, 3, 2, 2, 2, 35, 136, 3, 2, 2, 2, 37, 138, 3, 2, 2, 2, 39, 140,
	3, 2, 2, 2, 41, 142, 3, 2, 2, 2, 43, 144, 3, 2, 2, 2, 45, 46, 7, 42, 2,
	2, 46, 4, 3, 2, 2, 2, 47, 48, 7, 43, 2, 2, 48, 6, 3, 2, 2, 2, 49, 50, 7,
	38, 2, 2, 50, 8, 3, 2, 2, 2, 51, 52, 7, 48, 2, 2, 52, 53, 7, 108, 2, 2,
	53, 54, 7, 99, 2, 2, 54, 55, 7, 120, 2, 2, 55, 56, 7, 99, 2, 2, 56, 10,
	3, 2, 2, 2, 57, 59, 5, 41, 21, 2, 58, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2,
	2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 12, 3, 2, 2, 2, 62, 64,
	5, 15, 8, 2, 63, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2,
	65, 66, 3, 2, 2, 2, 66, 14, 3, 2, 2, 2, 67, 72, 5, 37, 19, 2, 68, 72, 5,
	35, 18, 2, 69, 72, 5, 39, 20, 2, 70, 72, 5, 41, 21, 2, 71, 67, 3, 2, 2,
	2, 71, 68, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 16,
	3, 2, 2, 2, 73, 74, 7, 48, 2, 2, 74, 18, 3, 2, 2, 2, 75, 76, 7, 99, 2,
	2, 76, 77, 7, 118, 2, 2, 77, 20, 3, 2, 2, 2, 78, 79, 7, 69, 2, 2, 79, 80,
	7, 99, 2, 2, 80, 81, 7, 119, 2, 2, 81, 82, 7, 117, 2, 2, 82, 83, 7, 103,
	2, 2, 83, 84, 7, 102, 2, 2, 84, 85, 7, 34, 2, 2, 85, 86, 7, 100, 2, 2,
	86, 87, 7, 123, 2, 2, 87, 88, 7, 60, 2, 2, 88, 22, 3, 2, 2, 2, 89, 90,
	7, 111, 2, 2, 90, 91, 7, 113, 2, 2, 91, 92, 7, 116, 2, 2, 92, 93, 7, 103,
	2, 2, 93, 24, 3, 2, 2, 2, 94, 95, 7, 48, 2, 2, 95, 96, 7, 48, 2, 2, 96,
	97, 7, 48, 2, 2, 97, 26, 3, 2, 2, 2, 98, 99, 7, 60, 2, 2, 99, 28, 3, 2,
	2, 2, 100, 101, 7, 80, 2, 2, 101, 102, 7, 99, 2, 2, 102, 103, 7, 118, 2,
	2, 103, 104, 7, 107, 2, 2, 104, 105, 7, 120, 2, 2, 105, 106, 7, 103, 2,
	2, 106, 107, 7, 34, 2, 2, 107, 108, 7, 79, 2, 2, 108, 109, 7, 103, 2, 2,
	109, 110, 7, 118, 2, 2, 110, 111, 7, 106, 2, 2, 111, 112, 7, 113, 2, 2,
	112, 113, 7, 102, 2, 2, 113, 30, 3, 2, 2, 2, 114, 115, 7, 87, 2, 2, 115,
	116, 7, 112, 2, 2, 116, 117, 7, 109, 2, 2, 117, 118, 7, 112, 2, 2, 118,
	119, 7, 113, 2, 2, 119, 120, 7, 121, 2, 2, 120, 121, 7, 112, 2, 2, 121,
	122, 7, 34, 2, 2, 122, 123, 7, 85, 2, 2, 123, 124, 7, 113, 2, 2, 124, 125,
	7, 119, 2, 2, 125, 126, 7, 116, 2, 2, 126, 127, 7, 101, 2, 2, 127, 128,
	7, 103, 2, 2, 128, 32, 3, 2, 2, 2, 129, 130, 7, 62, 2, 2, 130, 131, 7,
	107, 2, 2, 131, 132, 7, 112, 2, 2, 132, 133, 7, 107, 2, 2, 133, 134, 7,
	118, 2, 2, 134, 135, 7, 64, 2, 2, 135, 34, 3, 2, 2, 2, 136, 137, 4, 99,
	124, 2, 137, 36, 3, 2, 2, 2, 138, 139, 4, 67, 92, 2, 139, 38, 3, 2, 2,
	2, 140, 141, 7, 97, 2, 2, 141, 40, 3, 2, 2, 2, 142, 143, 4, 50, 59, 2,
	143, 42, 3, 2, 2, 2, 144, 145, 9, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 147,
	8, 22, 2, 2, 147, 44, 3, 2, 2, 2, 6, 2, 60, 65, 71, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'$'", "'.java'", "", "", "'.'", "'at'", "'Caused by:'",
	"'more'", "'...'", "':'", "'Native Method'", "'Unknown Source'", "'<init>'",
	"", "", "'_'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "Number", "JavaWord", "DOT", "AT", "CAUSED_BY", "MORE_",
	"ELLIPSIS", "COLON", "NATIVE_METHOD", "UNKNOWN_SOURCE", "INIT", "NonCapitalLetter",
	"CapitalLetter", "Symbol", "Digit", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "Number", "JavaWord", "JavaCharacter",
	"DOT", "AT", "CAUSED_BY", "MORE_", "ELLIPSIS", "COLON", "NATIVE_METHOD",
	"UNKNOWN_SOURCE", "INIT", "NonCapitalLetter", "CapitalLetter", "Symbol",
	"Digit", "WS",
}

type StackTraceLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewStackTraceLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *StackTraceLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewStackTraceLexer(input antlr.CharStream) *StackTraceLexer {
	l := new(StackTraceLexer)
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
	l.GrammarFileName = "StackTrace.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// StackTraceLexer tokens.
const (
	StackTraceLexerT__0             = 1
	StackTraceLexerT__1             = 2
	StackTraceLexerT__2             = 3
	StackTraceLexerT__3             = 4
	StackTraceLexerNumber           = 5
	StackTraceLexerJavaWord         = 6
	StackTraceLexerDOT              = 7
	StackTraceLexerAT               = 8
	StackTraceLexerCAUSED_BY        = 9
	StackTraceLexerMORE_            = 10
	StackTraceLexerELLIPSIS         = 11
	StackTraceLexerCOLON            = 12
	StackTraceLexerNATIVE_METHOD    = 13
	StackTraceLexerUNKNOWN_SOURCE   = 14
	StackTraceLexerINIT             = 15
	StackTraceLexerNonCapitalLetter = 16
	StackTraceLexerCapitalLetter    = 17
	StackTraceLexerSymbol           = 18
	StackTraceLexerDigit            = 19
	StackTraceLexerWS               = 20
)
