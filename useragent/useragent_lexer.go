// Code generated from useragent.g4 by ANTLR 4.9.3. DO NOT EDIT.

package useragent

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 38, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 20, 10, 4, 12, 4, 14, 4, 23, 11, 4, 3,
	4, 3, 4, 3, 5, 6, 5, 28, 10, 5, 13, 5, 14, 5, 29, 3, 6, 6, 6, 33, 10, 6,
	13, 6, 14, 6, 34, 3, 6, 3, 6, 2, 2, 7, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 3,
	2, 5, 3, 2, 43, 43, 5, 2, 50, 59, 67, 92, 99, 124, 5, 2, 12, 12, 15, 15,
	34, 34, 2, 40, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 15, 3, 2, 2, 2,
	7, 17, 3, 2, 2, 2, 9, 27, 3, 2, 2, 2, 11, 32, 3, 2, 2, 2, 13, 14, 7, 49,
	2, 2, 14, 4, 3, 2, 2, 2, 15, 16, 7, 48, 2, 2, 16, 6, 3, 2, 2, 2, 17, 21,
	7, 42, 2, 2, 18, 20, 10, 2, 2, 2, 19, 18, 3, 2, 2, 2, 20, 23, 3, 2, 2,
	2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 24, 3, 2, 2, 2, 23, 21,
	3, 2, 2, 2, 24, 25, 7, 43, 2, 2, 25, 8, 3, 2, 2, 2, 26, 28, 9, 3, 2, 2,
	27, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3,
	2, 2, 2, 30, 10, 3, 2, 2, 2, 31, 33, 9, 4, 2, 2, 32, 31, 3, 2, 2, 2, 33,
	34, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 36, 3, 2, 2,
	2, 36, 37, 8, 6, 2, 2, 37, 12, 3, 2, 2, 2, 6, 2, 21, 29, 34, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'/'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "COMMENT", "STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "COMMENT", "STRING", "WS",
}

type useragentLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewuseragentLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *useragentLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewuseragentLexer(input antlr.CharStream) *useragentLexer {
	l := new(useragentLexer)
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
	l.GrammarFileName = "useragent.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// useragentLexer tokens.
const (
	useragentLexerT__0    = 1
	useragentLexerT__1    = 2
	useragentLexerCOMMENT = 3
	useragentLexerSTRING  = 4
	useragentLexerWS      = 5
)
