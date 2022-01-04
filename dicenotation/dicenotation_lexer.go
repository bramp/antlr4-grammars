// Code generated from DiceNotationLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dicenotation

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 59, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 6, 3, 29, 10, 3, 13, 3, 14, 3, 30, 3, 4, 3, 4, 5, 4, 35,
	10, 4, 3, 5, 3, 5, 5, 5, 39, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 6, 12, 54, 10, 12, 13, 12,
	14, 12, 55, 3, 12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11, 2, 13,
	2, 15, 2, 17, 2, 19, 7, 21, 8, 23, 9, 3, 2, 4, 4, 2, 70, 70, 102, 102,
	4, 2, 11, 12, 15, 15, 2, 58, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 28, 3, 2, 2, 2, 7, 34, 3, 2, 2, 2,
	9, 38, 3, 2, 2, 2, 11, 40, 3, 2, 2, 2, 13, 42, 3, 2, 2, 2, 15, 44, 3, 2,
	2, 2, 17, 46, 3, 2, 2, 2, 19, 48, 3, 2, 2, 2, 21, 50, 3, 2, 2, 2, 23, 53,
	3, 2, 2, 2, 25, 26, 9, 2, 2, 2, 26, 4, 3, 2, 2, 2, 27, 29, 4, 50, 59, 2,
	28, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3,
	2, 2, 2, 31, 6, 3, 2, 2, 2, 32, 35, 5, 11, 6, 2, 33, 35, 5, 13, 7, 2, 34,
	32, 3, 2, 2, 2, 34, 33, 3, 2, 2, 2, 35, 8, 3, 2, 2, 2, 36, 39, 5, 15, 8,
	2, 37, 39, 5, 17, 9, 2, 38, 36, 3, 2, 2, 2, 38, 37, 3, 2, 2, 2, 39, 10,
	3, 2, 2, 2, 40, 41, 7, 45, 2, 2, 41, 12, 3, 2, 2, 2, 42, 43, 7, 47, 2,
	2, 43, 14, 3, 2, 2, 2, 44, 45, 7, 44, 2, 2, 45, 16, 3, 2, 2, 2, 46, 47,
	7, 49, 2, 2, 47, 18, 3, 2, 2, 2, 48, 49, 7, 42, 2, 2, 49, 20, 3, 2, 2,
	2, 50, 51, 7, 43, 2, 2, 51, 22, 3, 2, 2, 2, 52, 54, 9, 3, 2, 2, 53, 52,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2,
	56, 57, 3, 2, 2, 2, 57, 58, 8, 12, 2, 2, 58, 24, 3, 2, 2, 2, 7, 2, 30,
	34, 38, 55, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "DSEPARATOR", "DIGIT", "ADDOPERATOR", "MULTOPERATOR", "LPAREN", "RPAREN",
	"WS",
}

var lexerRuleNames = []string{
	"DSEPARATOR", "DIGIT", "ADDOPERATOR", "MULTOPERATOR", "ADD", "SUB", "MULT",
	"DIV", "LPAREN", "RPAREN", "WS",
}

type DiceNotationLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewDiceNotationLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *DiceNotationLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDiceNotationLexer(input antlr.CharStream) *DiceNotationLexer {
	l := new(DiceNotationLexer)
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
	l.GrammarFileName = "DiceNotationLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DiceNotationLexer tokens.
const (
	DiceNotationLexerDSEPARATOR   = 1
	DiceNotationLexerDIGIT        = 2
	DiceNotationLexerADDOPERATOR  = 3
	DiceNotationLexerMULTOPERATOR = 4
	DiceNotationLexerLPAREN       = 5
	DiceNotationLexerRPAREN       = 6
	DiceNotationLexerWS           = 7
)
