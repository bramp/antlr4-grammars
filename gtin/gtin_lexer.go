// Code generated from gtin.g4 by ANTLR 4.7.2. DO NOT EDIT.

package gtin

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 58, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 13, 6, 13, 53, 10, 13, 13, 13, 14, 13, 54, 3, 13,
	3, 13, 2, 2, 14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 3, 2, 3, 5, 2, 11, 12, 15, 15, 34, 34, 2, 58,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 3, 27, 3, 2, 2, 2, 5, 29, 3, 2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 33, 3,
	2, 2, 2, 11, 35, 3, 2, 2, 2, 13, 37, 3, 2, 2, 2, 15, 39, 3, 2, 2, 2, 17,
	41, 3, 2, 2, 2, 19, 43, 3, 2, 2, 2, 21, 45, 3, 2, 2, 2, 23, 47, 3, 2, 2,
	2, 25, 52, 3, 2, 2, 2, 27, 28, 7, 50, 2, 2, 28, 4, 3, 2, 2, 2, 29, 30,
	7, 51, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 52, 2, 2, 32, 8, 3, 2, 2, 2,
	33, 34, 7, 53, 2, 2, 34, 10, 3, 2, 2, 2, 35, 36, 7, 54, 2, 2, 36, 12, 3,
	2, 2, 2, 37, 38, 7, 55, 2, 2, 38, 14, 3, 2, 2, 2, 39, 40, 7, 56, 2, 2,
	40, 16, 3, 2, 2, 2, 41, 42, 7, 57, 2, 2, 42, 18, 3, 2, 2, 2, 43, 44, 7,
	58, 2, 2, 44, 20, 3, 2, 2, 2, 45, 46, 7, 59, 2, 2, 46, 22, 3, 2, 2, 2,
	47, 48, 7, 47, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50, 8, 12, 2, 2, 50, 24, 3,
	2, 2, 2, 51, 53, 9, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54,
	52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57, 8, 13,
	2, 2, 57, 26, 3, 2, 2, 2, 4, 2, 54, 3, 8, 2, 2,
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
	"", "'0'", "'1'", "'2'", "'3'", "'4'", "'5'", "'6'", "'7'", "'8'", "'9'",
	"'-'",
}

var lexerSymbolicNames = []string{
	"", "DIGIT_0", "DIGIT_1", "DIGIT_2", "DIGIT_3", "DIGIT_4", "DIGIT_5", "DIGIT_6",
	"DIGIT_7", "DIGIT_8", "DIGIT_9", "HYPHEN", "WS",
}

var lexerRuleNames = []string{
	"DIGIT_0", "DIGIT_1", "DIGIT_2", "DIGIT_3", "DIGIT_4", "DIGIT_5", "DIGIT_6",
	"DIGIT_7", "DIGIT_8", "DIGIT_9", "HYPHEN", "WS",
}

type gtinLexer struct {
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

func NewgtinLexer(input antlr.CharStream) *gtinLexer {

	l := new(gtinLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "gtin.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gtinLexer tokens.
const (
	gtinLexerDIGIT_0 = 1
	gtinLexerDIGIT_1 = 2
	gtinLexerDIGIT_2 = 3
	gtinLexerDIGIT_3 = 4
	gtinLexerDIGIT_4 = 5
	gtinLexerDIGIT_5 = 6
	gtinLexerDIGIT_6 = 7
	gtinLexerDIGIT_7 = 8
	gtinLexerDIGIT_8 = 9
	gtinLexerDIGIT_9 = 10
	gtinLexerHYPHEN  = 11
	gtinLexerWS      = 12
)
