// Generated from sexpression.g4 by ANTLR 4.7.

package sexpression

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 74, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2,
	26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 2, 3, 2, 3, 3, 6, 3, 34, 10, 3,
	13, 3, 14, 3, 35, 3, 3, 3, 3, 3, 4, 5, 4, 41, 10, 4, 3, 4, 6, 4, 44, 10,
	4, 13, 4, 14, 4, 45, 3, 4, 3, 4, 6, 4, 50, 10, 4, 13, 4, 14, 4, 51, 5,
	4, 54, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 59, 10, 5, 12, 5, 14, 5, 62, 11,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 5, 9, 71, 10, 9, 3, 10, 3,
	10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 2, 19, 2,
	3, 2, 6, 4, 2, 36, 36, 94, 94, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 45,
	45, 47, 47, 6, 2, 44, 45, 47, 49, 67, 92, 99, 124, 2, 80, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 33, 3, 2,
	2, 2, 7, 40, 3, 2, 2, 2, 9, 55, 3, 2, 2, 2, 11, 63, 3, 2, 2, 2, 13, 65,
	3, 2, 2, 2, 15, 67, 3, 2, 2, 2, 17, 70, 3, 2, 2, 2, 19, 72, 3, 2, 2, 2,
	21, 27, 7, 36, 2, 2, 22, 23, 7, 94, 2, 2, 23, 26, 11, 2, 2, 2, 24, 26,
	10, 2, 2, 2, 25, 22, 3, 2, 2, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2,
	27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 30, 3, 2, 2, 2, 29, 27, 3,
	2, 2, 2, 30, 31, 7, 36, 2, 2, 31, 4, 3, 2, 2, 2, 32, 34, 9, 3, 2, 2, 33,
	32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2,
	2, 36, 37, 3, 2, 2, 2, 37, 38, 8, 3, 2, 2, 38, 6, 3, 2, 2, 2, 39, 41, 9,
	4, 2, 2, 40, 39, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 43, 3, 2, 2, 2, 42,
	44, 5, 19, 10, 2, 43, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 43, 3, 2,
	2, 2, 45, 46, 3, 2, 2, 2, 46, 53, 3, 2, 2, 2, 47, 49, 7, 48, 2, 2, 48,
	50, 5, 19, 10, 2, 49, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 49, 3, 2,
	2, 2, 51, 52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 47, 3, 2, 2, 2, 53, 54,
	3, 2, 2, 2, 54, 8, 3, 2, 2, 2, 55, 60, 5, 17, 9, 2, 56, 59, 5, 17, 9, 2,
	57, 59, 5, 19, 10, 2, 58, 56, 3, 2, 2, 2, 58, 57, 3, 2, 2, 2, 59, 62, 3,
	2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 10, 3, 2, 2, 2, 62,
	60, 3, 2, 2, 2, 63, 64, 7, 42, 2, 2, 64, 12, 3, 2, 2, 2, 65, 66, 7, 43,
	2, 2, 66, 14, 3, 2, 2, 2, 67, 68, 7, 48, 2, 2, 68, 16, 3, 2, 2, 2, 69,
	71, 9, 5, 2, 2, 70, 69, 3, 2, 2, 2, 71, 18, 3, 2, 2, 2, 72, 73, 4, 50,
	59, 2, 73, 20, 3, 2, 2, 2, 13, 2, 25, 27, 35, 40, 45, 51, 53, 58, 60, 70,
	3, 8, 2, 2,
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
	"", "", "", "", "", "'('", "')'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "STRING", "WHITESPACE", "NUMBER", "SYMBOL", "LPAREN", "RPAREN", "DOT",
}

var lexerRuleNames = []string{
	"STRING", "WHITESPACE", "NUMBER", "SYMBOL", "LPAREN", "RPAREN", "DOT",
	"SYMBOL_START", "DIGIT",
}

type sexpressionLexer struct {
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

func NewsexpressionLexer(input antlr.CharStream) *sexpressionLexer {

	l := new(sexpressionLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "sexpression.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// sexpressionLexer tokens.
const (
	sexpressionLexerSTRING     = 1
	sexpressionLexerWHITESPACE = 2
	sexpressionLexerNUMBER     = 3
	sexpressionLexerSYMBOL     = 4
	sexpressionLexerLPAREN     = 5
	sexpressionLexerRPAREN     = 6
	sexpressionLexerDOT        = 7
)
