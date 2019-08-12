// Code generated from tsv.g4 by ANTLR 4.7.2. DO NOT EDIT.

package tsv

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 34, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 3, 6,
	3, 15, 10, 3, 13, 3, 14, 3, 16, 3, 4, 6, 4, 20, 10, 4, 13, 4, 14, 4, 21,
	3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 28, 10, 5, 12, 5, 14, 5, 31, 11, 5, 3, 5,
	3, 5, 2, 2, 6, 3, 3, 5, 4, 7, 5, 9, 6, 3, 2, 5, 4, 2, 12, 12, 15, 15, 6,
	2, 12, 12, 15, 15, 36, 36, 46, 46, 3, 2, 36, 36, 2, 37, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 3, 11, 3, 2, 2,
	2, 5, 14, 3, 2, 2, 2, 7, 19, 3, 2, 2, 2, 9, 23, 3, 2, 2, 2, 11, 12, 7,
	11, 2, 2, 12, 4, 3, 2, 2, 2, 13, 15, 9, 2, 2, 2, 14, 13, 3, 2, 2, 2, 15,
	16, 3, 2, 2, 2, 16, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 6, 3, 2, 2,
	2, 18, 20, 10, 3, 2, 2, 19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 19,
	3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 8, 3, 2, 2, 2, 23, 29, 7, 36, 2, 2,
	24, 25, 7, 36, 2, 2, 25, 28, 7, 36, 2, 2, 26, 28, 10, 4, 2, 2, 27, 24,
	3, 2, 2, 2, 27, 26, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2,
	29, 30, 3, 2, 2, 2, 30, 32, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 32, 33, 7,
	36, 2, 2, 33, 10, 3, 2, 2, 2, 7, 2, 16, 21, 27, 29, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'\t'",
}

var lexerSymbolicNames = []string{
	"", "TAB", "EOL", "TEXT", "STRING",
}

var lexerRuleNames = []string{
	"TAB", "EOL", "TEXT", "STRING",
}

type tsvLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewtsvLexer(input antlr.CharStream) *tsvLexer {

	l := new(tsvLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "tsv.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// tsvLexer tokens.
const (
	tsvLexerTAB    = 1
	tsvLexerEOL    = 2
	tsvLexerTEXT   = 3
	tsvLexerSTRING = 4
)
