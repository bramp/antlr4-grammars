// Code generated from arithmetic.g4 by ANTLR 4.7.2. DO NOT EDIT.

package arithmetic

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 16, 117,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 7, 2,
	46, 10, 2, 12, 2, 14, 2, 49, 11, 2, 3, 3, 5, 3, 52, 10, 3, 3, 4, 3, 4,
	5, 4, 56, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 61, 10, 5, 3, 5, 3, 5, 5, 5, 65,
	10, 5, 3, 6, 6, 6, 68, 10, 6, 13, 6, 14, 6, 69, 3, 6, 3, 6, 6, 6, 74, 10,
	6, 13, 6, 14, 6, 75, 5, 6, 78, 10, 6, 3, 7, 6, 7, 81, 10, 7, 13, 7, 14,
	7, 82, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 6, 21, 112, 10, 21, 13, 21,
	14, 21, 113, 3, 21, 3, 21, 2, 2, 22, 3, 3, 5, 2, 7, 2, 9, 4, 11, 2, 13,
	2, 15, 2, 17, 2, 19, 5, 21, 6, 23, 7, 25, 8, 27, 9, 29, 10, 31, 11, 33,
	12, 35, 13, 37, 14, 39, 15, 41, 16, 3, 2, 6, 5, 2, 67, 92, 97, 97, 99,
	124, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 119, 2, 3, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2,
	2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 3, 43, 3,
	2, 2, 2, 5, 51, 3, 2, 2, 2, 7, 55, 3, 2, 2, 2, 9, 57, 3, 2, 2, 2, 11, 67,
	3, 2, 2, 2, 13, 80, 3, 2, 2, 2, 15, 84, 3, 2, 2, 2, 17, 86, 3, 2, 2, 2,
	19, 88, 3, 2, 2, 2, 21, 90, 3, 2, 2, 2, 23, 92, 3, 2, 2, 2, 25, 94, 3,
	2, 2, 2, 27, 96, 3, 2, 2, 2, 29, 98, 3, 2, 2, 2, 31, 100, 3, 2, 2, 2, 33,
	102, 3, 2, 2, 2, 35, 104, 3, 2, 2, 2, 37, 106, 3, 2, 2, 2, 39, 108, 3,
	2, 2, 2, 41, 111, 3, 2, 2, 2, 43, 47, 5, 5, 3, 2, 44, 46, 5, 7, 4, 2, 45,
	44, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2,
	2, 48, 4, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 52, 9, 2, 2, 2, 51, 50, 3,
	2, 2, 2, 52, 6, 3, 2, 2, 2, 53, 56, 5, 5, 3, 2, 54, 56, 4, 50, 59, 2, 55,
	53, 3, 2, 2, 2, 55, 54, 3, 2, 2, 2, 56, 8, 3, 2, 2, 2, 57, 64, 5, 11, 6,
	2, 58, 60, 5, 15, 8, 2, 59, 61, 5, 17, 9, 2, 60, 59, 3, 2, 2, 2, 60, 61,
	3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 5, 13, 7, 2, 63, 65, 3, 2, 2, 2,
	64, 58, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 10, 3, 2, 2, 2, 66, 68, 4,
	50, 59, 2, 67, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2,
	69, 70, 3, 2, 2, 2, 70, 77, 3, 2, 2, 2, 71, 73, 7, 48, 2, 2, 72, 74, 4,
	50, 59, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2,
	75, 76, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77, 71, 3, 2, 2, 2, 77, 78, 3,
	2, 2, 2, 78, 12, 3, 2, 2, 2, 79, 81, 4, 50, 59, 2, 80, 79, 3, 2, 2, 2,
	81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 14, 3,
	2, 2, 2, 84, 85, 9, 3, 2, 2, 85, 16, 3, 2, 2, 2, 86, 87, 9, 4, 2, 2, 87,
	18, 3, 2, 2, 2, 88, 89, 7, 42, 2, 2, 89, 20, 3, 2, 2, 2, 90, 91, 7, 43,
	2, 2, 91, 22, 3, 2, 2, 2, 92, 93, 7, 45, 2, 2, 93, 24, 3, 2, 2, 2, 94,
	95, 7, 47, 2, 2, 95, 26, 3, 2, 2, 2, 96, 97, 7, 44, 2, 2, 97, 28, 3, 2,
	2, 2, 98, 99, 7, 49, 2, 2, 99, 30, 3, 2, 2, 2, 100, 101, 7, 64, 2, 2, 101,
	32, 3, 2, 2, 2, 102, 103, 7, 62, 2, 2, 103, 34, 3, 2, 2, 2, 104, 105, 7,
	63, 2, 2, 105, 36, 3, 2, 2, 2, 106, 107, 7, 48, 2, 2, 107, 38, 3, 2, 2,
	2, 108, 109, 7, 96, 2, 2, 109, 40, 3, 2, 2, 2, 110, 112, 9, 5, 2, 2, 111,
	110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114,
	3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 116, 8, 21, 2, 2, 116, 42, 3, 2,
	2, 2, 13, 2, 47, 51, 55, 60, 64, 69, 75, 77, 82, 113, 3, 8, 2, 2,
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
	"", "", "", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "'>'", "'<'", "'='",
	"'.'", "'^'",
}

var lexerSymbolicNames = []string{
	"", "VARIABLE", "SCIENTIFIC_NUMBER", "LPAREN", "RPAREN", "PLUS", "MINUS",
	"TIMES", "DIV", "GT", "LT", "EQ", "POINT", "POW", "WS",
}

var lexerRuleNames = []string{
	"VARIABLE", "VALID_ID_START", "VALID_ID_CHAR", "SCIENTIFIC_NUMBER", "NUMBER",
	"UNSIGNED_INTEGER", "E", "SIGN", "LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES",
	"DIV", "GT", "LT", "EQ", "POINT", "POW", "WS",
}

type arithmeticLexer struct {
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

func NewarithmeticLexer(input antlr.CharStream) *arithmeticLexer {

	l := new(arithmeticLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "arithmetic.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// arithmeticLexer tokens.
const (
	arithmeticLexerVARIABLE          = 1
	arithmeticLexerSCIENTIFIC_NUMBER = 2
	arithmeticLexerLPAREN            = 3
	arithmeticLexerRPAREN            = 4
	arithmeticLexerPLUS              = 5
	arithmeticLexerMINUS             = 6
	arithmeticLexerTIMES             = 7
	arithmeticLexerDIV               = 8
	arithmeticLexerGT                = 9
	arithmeticLexerLT                = 10
	arithmeticLexerEQ                = 11
	arithmeticLexerPOINT             = 12
	arithmeticLexerPOW               = 13
	arithmeticLexerWS                = 14
)
