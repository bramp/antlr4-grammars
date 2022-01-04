// Code generated from fasta.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fasta

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 64, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 7, 2, 22, 10, 2, 12, 2, 14, 2, 25,
	11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 33, 10, 3, 12, 3, 14,
	3, 36, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 6, 4, 43, 10, 4, 13, 4, 14,
	4, 44, 3, 5, 5, 5, 48, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 6, 9, 59, 10, 9, 13, 9, 14, 9, 60, 3, 9, 3, 9, 3, 23, 2, 10,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 2, 13, 2, 15, 2, 17, 7, 3, 2, 5, 3, 2, 50,
	59, 4, 2, 67, 92, 99, 124, 9, 2, 34, 34, 40, 43, 45, 49, 60, 60, 93, 93,
	95, 95, 97, 97, 2, 67, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5, 28, 3,
	2, 2, 2, 7, 42, 3, 2, 2, 2, 9, 47, 3, 2, 2, 2, 11, 51, 3, 2, 2, 2, 13,
	53, 3, 2, 2, 2, 15, 55, 3, 2, 2, 2, 17, 58, 3, 2, 2, 2, 19, 23, 7, 61,
	2, 2, 20, 22, 11, 2, 2, 2, 21, 20, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23,
	24, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 24, 26, 3, 2, 2, 2, 25, 23, 3, 2, 2,
	2, 26, 27, 5, 9, 5, 2, 27, 4, 3, 2, 2, 2, 28, 29, 7, 64, 2, 2, 29, 34,
	5, 7, 4, 2, 30, 31, 7, 126, 2, 2, 31, 33, 5, 7, 4, 2, 32, 30, 3, 2, 2,
	2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 37,
	3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 38, 5, 9, 5, 2, 38, 6, 3, 2, 2, 2,
	39, 43, 5, 11, 6, 2, 40, 43, 5, 13, 7, 2, 41, 43, 5, 15, 8, 2, 42, 39,
	3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2,
	44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 8, 3, 2, 2, 2, 46, 48, 7, 15,
	2, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50,
	7, 12, 2, 2, 50, 10, 3, 2, 2, 2, 51, 52, 9, 2, 2, 2, 52, 12, 3, 2, 2, 2,
	53, 54, 9, 3, 2, 2, 54, 14, 3, 2, 2, 2, 55, 56, 9, 4, 2, 2, 56, 16, 3,
	2, 2, 2, 57, 59, 5, 13, 7, 2, 58, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60,
	58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 5, 9, 5,
	2, 63, 18, 3, 2, 2, 2, 9, 2, 23, 34, 42, 44, 47, 60, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames []string

var lexerSymbolicNames = []string{
	"", "COMMENTLINE", "DESCRIPTIONLINE", "TEXT", "EOL", "SEQUENCELINE",
}

var lexerRuleNames = []string{
	"COMMENTLINE", "DESCRIPTIONLINE", "TEXT", "EOL", "DIGIT", "LETTER", "SYMBOL",
	"SEQUENCELINE",
}

type fastaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewfastaLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *fastaLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewfastaLexer(input antlr.CharStream) *fastaLexer {
	l := new(fastaLexer)
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
	l.GrammarFileName = "fasta.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// fastaLexer tokens.
const (
	fastaLexerCOMMENTLINE     = 1
	fastaLexerDESCRIPTIONLINE = 2
	fastaLexerTEXT            = 3
	fastaLexerEOL             = 4
	fastaLexerSEQUENCELINE    = 5
)
