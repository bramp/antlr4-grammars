// Code generated from callable_.g4 by ANTLR 4.9.3. DO NOT EDIT.

package callable_

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 56, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 6, 5, 31, 10, 5, 13, 5, 14, 5, 32, 3, 6, 3, 6,
	7, 6, 37, 10, 6, 12, 6, 14, 6, 40, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 10, 6, 10, 51, 10, 10, 13, 10, 14, 10, 52, 3, 10,
	3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 3, 2, 7, 3, 2, 36, 36, 3, 2, 67, 92, 3, 2, 99, 124, 4, 2, 12, 12, 15,
	15, 4, 2, 11, 11, 34, 34, 2, 60, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2, 2,
	2, 5, 23, 3, 2, 2, 2, 7, 25, 3, 2, 2, 2, 9, 30, 3, 2, 2, 2, 11, 34, 3,
	2, 2, 2, 13, 43, 3, 2, 2, 2, 15, 45, 3, 2, 2, 2, 17, 47, 3, 2, 2, 2, 19,
	50, 3, 2, 2, 2, 21, 22, 7, 42, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7, 43,
	2, 2, 24, 6, 3, 2, 2, 2, 25, 26, 7, 46, 2, 2, 26, 8, 3, 2, 2, 2, 27, 31,
	5, 13, 7, 2, 28, 31, 5, 15, 8, 2, 29, 31, 7, 47, 2, 2, 30, 27, 3, 2, 2,
	2, 30, 28, 3, 2, 2, 2, 30, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 30,
	3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 10, 3, 2, 2, 2, 34, 38, 7, 36, 2, 2,
	35, 37, 10, 2, 2, 2, 36, 35, 3, 2, 2, 2, 37, 40, 3, 2, 2, 2, 38, 36, 3,
	2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 41, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 41,
	42, 7, 36, 2, 2, 42, 12, 3, 2, 2, 2, 43, 44, 9, 3, 2, 2, 44, 14, 3, 2,
	2, 2, 45, 46, 9, 4, 2, 2, 46, 16, 3, 2, 2, 2, 47, 48, 9, 5, 2, 2, 48, 18,
	3, 2, 2, 2, 49, 51, 9, 6, 2, 2, 50, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2,
	52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 8,
	10, 2, 2, 55, 20, 3, 2, 2, 2, 7, 2, 30, 32, 38, 52, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "','",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "ID", "STRING", "LETTER_UPPER", "LETTER_LOWER", "EOL",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "ID", "STRING", "LETTER_UPPER", "LETTER_LOWER",
	"EOL", "WS",
}

type callable_Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// Newcallable_Lexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *callable_Lexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newcallable_Lexer(input antlr.CharStream) *callable_Lexer {
	l := new(callable_Lexer)
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
	l.GrammarFileName = "callable_.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// callable_Lexer tokens.
const (
	callable_LexerT__0         = 1
	callable_LexerT__1         = 2
	callable_LexerT__2         = 3
	callable_LexerID           = 4
	callable_LexerSTRING       = 5
	callable_LexerLETTER_UPPER = 6
	callable_LexerLETTER_LOWER = 7
	callable_LexerEOL          = 8
	callable_LexerWS           = 9
)
