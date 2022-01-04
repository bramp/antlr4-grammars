// Code generated from vmf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package vmf

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 37, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 22, 10, 4, 12, 4, 14, 4, 25,
	11, 4, 3, 4, 3, 4, 3, 5, 6, 5, 30, 10, 5, 13, 5, 14, 5, 31, 3, 6, 3, 6,
	3, 6, 3, 6, 2, 2, 7, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 3, 2, 6, 6, 2, 12,
	12, 15, 15, 36, 36, 94, 94, 4, 2, 36, 36, 94, 94, 12, 2, 38, 38, 42, 43,
	47, 47, 49, 59, 62, 62, 64, 64, 66, 93, 95, 95, 97, 97, 99, 124, 5, 2,
	11, 12, 15, 15, 34, 34, 2, 39, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5,
	15, 3, 2, 2, 2, 7, 17, 3, 2, 2, 2, 9, 29, 3, 2, 2, 2, 11, 33, 3, 2, 2,
	2, 13, 14, 7, 125, 2, 2, 14, 4, 3, 2, 2, 2, 15, 16, 7, 127, 2, 2, 16, 6,
	3, 2, 2, 2, 17, 23, 7, 36, 2, 2, 18, 22, 10, 2, 2, 2, 19, 20, 7, 94, 2,
	2, 20, 22, 9, 3, 2, 2, 21, 18, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 22, 25,
	3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 26, 3, 2, 2, 2,
	25, 23, 3, 2, 2, 2, 26, 27, 7, 36, 2, 2, 27, 8, 3, 2, 2, 2, 28, 30, 9,
	4, 2, 2, 29, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31,
	32, 3, 2, 2, 2, 32, 10, 3, 2, 2, 2, 33, 34, 9, 5, 2, 2, 34, 35, 3, 2, 2,
	2, 35, 36, 8, 6, 2, 2, 36, 12, 3, 2, 2, 2, 6, 2, 21, 23, 31, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "", "", "QUOTEDSTTRING", "STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "QUOTEDSTTRING", "STRING", "WS",
}

type vmfLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewvmfLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *vmfLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewvmfLexer(input antlr.CharStream) *vmfLexer {
	l := new(vmfLexer)
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
	l.GrammarFileName = "vmf.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// vmfLexer tokens.
const (
	vmfLexerT__0          = 1
	vmfLexerT__1          = 2
	vmfLexerQUOTEDSTTRING = 3
	vmfLexerSTRING        = 4
	vmfLexerWS            = 5
)
