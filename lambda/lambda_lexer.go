// Code generated from lambda.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lambda

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 34, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 26,
	10, 6, 12, 6, 14, 6, 29, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 2, 2, 8, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 3, 2, 5, 3, 2, 99, 124, 5, 2, 50, 59, 67,
	92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 34, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 3, 15, 3, 2, 2, 2, 5, 17, 3, 2, 2, 2, 7, 19, 3, 2, 2, 2,
	9, 21, 3, 2, 2, 2, 11, 23, 3, 2, 2, 2, 13, 30, 3, 2, 2, 2, 15, 16, 7, 957,
	2, 2, 16, 4, 3, 2, 2, 2, 17, 18, 7, 48, 2, 2, 18, 6, 3, 2, 2, 2, 19, 20,
	7, 42, 2, 2, 20, 8, 3, 2, 2, 2, 21, 22, 7, 43, 2, 2, 22, 10, 3, 2, 2, 2,
	23, 27, 9, 2, 2, 2, 24, 26, 9, 3, 2, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3,
	2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 12, 3, 2, 2, 2, 29,
	27, 3, 2, 2, 2, 30, 31, 9, 4, 2, 2, 31, 32, 3, 2, 2, 2, 32, 33, 8, 7, 2,
	2, 33, 14, 3, 2, 2, 2, 4, 2, 27, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'\u03BB'", "'.'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "VARIABLE", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "VARIABLE", "WS",
}

type lambdaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewlambdaLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *lambdaLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewlambdaLexer(input antlr.CharStream) *lambdaLexer {
	l := new(lambdaLexer)
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
	l.GrammarFileName = "lambda.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// lambdaLexer tokens.
const (
	lambdaLexerT__0     = 1
	lambdaLexerT__1     = 2
	lambdaLexerT__2     = 3
	lambdaLexerT__3     = 4
	lambdaLexerVARIABLE = 5
	lambdaLexerWS       = 6
)
