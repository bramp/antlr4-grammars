// Code generated from gff3.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gff3

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
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6,
	50, 10, 6, 12, 6, 14, 6, 53, 11, 6, 3, 6, 3, 6, 3, 7, 5, 7, 58, 10, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 6, 8, 65, 10, 8, 13, 8, 14, 8, 66, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 51, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 2, 19, 2, 21, 2, 3, 2, 5, 4, 2, 67, 92, 99, 124,
	3, 2, 50, 59, 9, 2, 34, 35, 38, 39, 44, 48, 60, 60, 65, 66, 96, 97, 126,
	126, 2, 75, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3,
	23, 3, 2, 2, 2, 5, 25, 3, 2, 2, 2, 7, 27, 3, 2, 2, 2, 9, 29, 3, 2, 2, 2,
	11, 47, 3, 2, 2, 2, 13, 57, 3, 2, 2, 2, 15, 64, 3, 2, 2, 2, 17, 68, 3,
	2, 2, 2, 19, 70, 3, 2, 2, 2, 21, 72, 3, 2, 2, 2, 23, 24, 7, 11, 2, 2, 24,
	4, 3, 2, 2, 2, 25, 26, 7, 61, 2, 2, 26, 6, 3, 2, 2, 2, 27, 28, 7, 63, 2,
	2, 28, 8, 3, 2, 2, 2, 29, 30, 7, 37, 2, 2, 30, 31, 7, 37, 2, 2, 31, 32,
	7, 105, 2, 2, 32, 33, 7, 104, 2, 2, 33, 34, 7, 104, 2, 2, 34, 35, 7, 47,
	2, 2, 35, 36, 7, 120, 2, 2, 36, 37, 7, 103, 2, 2, 37, 38, 7, 116, 2, 2,
	38, 39, 7, 117, 2, 2, 39, 40, 7, 107, 2, 2, 40, 41, 7, 113, 2, 2, 41, 42,
	7, 112, 2, 2, 42, 43, 7, 34, 2, 2, 43, 44, 7, 53, 2, 2, 44, 45, 3, 2, 2,
	2, 45, 46, 5, 13, 7, 2, 46, 10, 3, 2, 2, 2, 47, 51, 7, 37, 2, 2, 48, 50,
	11, 2, 2, 2, 49, 48, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2,
	51, 49, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 55, 5,
	13, 7, 2, 55, 12, 3, 2, 2, 2, 56, 58, 7, 15, 2, 2, 57, 56, 3, 2, 2, 2,
	57, 58, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 60, 7, 12, 2, 2, 60, 14, 3,
	2, 2, 2, 61, 65, 5, 17, 9, 2, 62, 65, 5, 21, 11, 2, 63, 65, 5, 19, 10,
	2, 64, 61, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 63, 3, 2, 2, 2, 65, 66,
	3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 16, 3, 2, 2, 2,
	68, 69, 9, 2, 2, 2, 69, 18, 3, 2, 2, 2, 70, 71, 9, 3, 2, 2, 71, 20, 3,
	2, 2, 2, 72, 73, 9, 4, 2, 2, 73, 22, 3, 2, 2, 2, 7, 2, 51, 57, 64, 66,
	2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'\t'", "';'", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "HEADER", "COMMENTLINE", "EOL", "TEXT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "HEADER", "COMMENTLINE", "EOL", "TEXT", "CHAR",
	"DIGIT", "SYMBOL",
}

type gff3Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// Newgff3Lexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *gff3Lexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newgff3Lexer(input antlr.CharStream) *gff3Lexer {
	l := new(gff3Lexer)
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
	l.GrammarFileName = "gff3.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gff3Lexer tokens.
const (
	gff3LexerT__0        = 1
	gff3LexerT__1        = 2
	gff3LexerT__2        = 3
	gff3LexerHEADER      = 4
	gff3LexerCOMMENTLINE = 5
	gff3LexerEOL         = 6
	gff3LexerTEXT        = 7
)
