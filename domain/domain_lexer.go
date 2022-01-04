// Code generated from domain.g4 by ANTLR 4.9.3. DO NOT EDIT.

package domain

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 52, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 26,
	10, 4, 12, 4, 14, 4, 29, 11, 4, 3, 4, 7, 4, 32, 10, 4, 12, 4, 14, 4, 35,
	11, 4, 3, 5, 3, 5, 5, 5, 39, 10, 5, 3, 6, 3, 6, 5, 6, 43, 10, 6, 3, 7,
	3, 7, 5, 7, 47, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 2, 2, 10, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 3, 2, 4, 4, 2, 67, 92, 99, 124, 3,
	2, 50, 59, 2, 56, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5, 21, 3, 2, 2, 2, 7, 23, 3, 2,
	2, 2, 9, 36, 3, 2, 2, 2, 11, 42, 3, 2, 2, 2, 13, 46, 3, 2, 2, 2, 15, 48,
	3, 2, 2, 2, 17, 50, 3, 2, 2, 2, 19, 20, 7, 34, 2, 2, 20, 4, 3, 2, 2, 2,
	21, 22, 7, 48, 2, 2, 22, 6, 3, 2, 2, 2, 23, 33, 5, 15, 8, 2, 24, 26, 5,
	9, 5, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27,
	28, 3, 2, 2, 2, 28, 30, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 32, 5, 13,
	7, 2, 31, 27, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34,
	3, 2, 2, 2, 34, 8, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 36, 38, 5, 11, 6, 2,
	37, 39, 5, 9, 5, 2, 38, 37, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 10, 3,
	2, 2, 2, 40, 43, 5, 13, 7, 2, 41, 43, 7, 47, 2, 2, 42, 40, 3, 2, 2, 2,
	42, 41, 3, 2, 2, 2, 43, 12, 3, 2, 2, 2, 44, 47, 5, 15, 8, 2, 45, 47, 5,
	17, 9, 2, 46, 44, 3, 2, 2, 2, 46, 45, 3, 2, 2, 2, 47, 14, 3, 2, 2, 2, 48,
	49, 9, 2, 2, 2, 49, 16, 3, 2, 2, 2, 50, 51, 9, 3, 2, 2, 51, 18, 3, 2, 2,
	2, 8, 2, 27, 33, 38, 42, 46, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "' '", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "LABEL", "LDH_STR", "LET_DIG_HYP", "LET_DIG", "LETTER", "DIGIT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "LABEL", "LDH_STR", "LET_DIG_HYP", "LET_DIG", "LETTER",
	"DIGIT",
}

type domainLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewdomainLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *domainLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdomainLexer(input antlr.CharStream) *domainLexer {
	l := new(domainLexer)
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
	l.GrammarFileName = "domain.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// domainLexer tokens.
const (
	domainLexerT__0        = 1
	domainLexerT__1        = 2
	domainLexerLABEL       = 3
	domainLexerLDH_STR     = 4
	domainLexerLET_DIG_HYP = 5
	domainLexerLET_DIG     = 6
	domainLexerLETTER      = 7
	domainLexerDIGIT       = 8
)
