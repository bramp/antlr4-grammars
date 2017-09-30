// Generated from arithmetic.g4 by ANTLR 4.7.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 16, 110,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 7, 2, 44, 10, 2, 12,
	2, 14, 2, 47, 11, 2, 3, 3, 5, 3, 50, 10, 3, 3, 4, 3, 4, 5, 4, 54, 10, 4,
	3, 5, 3, 5, 3, 5, 5, 5, 59, 10, 5, 3, 5, 3, 5, 5, 5, 63, 10, 5, 3, 6, 6,
	6, 66, 10, 6, 13, 6, 14, 6, 67, 3, 6, 3, 6, 6, 6, 72, 10, 6, 13, 6, 14,
	6, 73, 5, 6, 76, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 6, 20, 105,
	10, 20, 13, 20, 14, 20, 106, 3, 20, 3, 20, 2, 2, 21, 3, 3, 5, 2, 7, 2,
	9, 4, 11, 2, 13, 2, 15, 2, 17, 5, 19, 6, 21, 7, 23, 8, 25, 9, 27, 10, 29,
	11, 31, 12, 33, 13, 35, 14, 37, 15, 39, 16, 3, 2, 6, 5, 2, 67, 92, 97,
	97, 99, 124, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 5, 2, 11, 12,
	15, 15, 34, 34, 2, 112, 2, 3, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 17, 3, 2,
	2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3,
	2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33,
	3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 3,
	41, 3, 2, 2, 2, 5, 49, 3, 2, 2, 2, 7, 53, 3, 2, 2, 2, 9, 55, 3, 2, 2, 2,
	11, 65, 3, 2, 2, 2, 13, 77, 3, 2, 2, 2, 15, 79, 3, 2, 2, 2, 17, 81, 3,
	2, 2, 2, 19, 83, 3, 2, 2, 2, 21, 85, 3, 2, 2, 2, 23, 87, 3, 2, 2, 2, 25,
	89, 3, 2, 2, 2, 27, 91, 3, 2, 2, 2, 29, 93, 3, 2, 2, 2, 31, 95, 3, 2, 2,
	2, 33, 97, 3, 2, 2, 2, 35, 99, 3, 2, 2, 2, 37, 101, 3, 2, 2, 2, 39, 104,
	3, 2, 2, 2, 41, 45, 5, 5, 3, 2, 42, 44, 5, 7, 4, 2, 43, 42, 3, 2, 2, 2,
	44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 4, 3, 2,
	2, 2, 47, 45, 3, 2, 2, 2, 48, 50, 9, 2, 2, 2, 49, 48, 3, 2, 2, 2, 50, 6,
	3, 2, 2, 2, 51, 54, 5, 5, 3, 2, 52, 54, 4, 50, 59, 2, 53, 51, 3, 2, 2,
	2, 53, 52, 3, 2, 2, 2, 54, 8, 3, 2, 2, 2, 55, 62, 5, 11, 6, 2, 56, 58,
	5, 13, 7, 2, 57, 59, 5, 15, 8, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3, 2, 2,
	2, 59, 60, 3, 2, 2, 2, 60, 61, 5, 11, 6, 2, 61, 63, 3, 2, 2, 2, 62, 56,
	3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 10, 3, 2, 2, 2, 64, 66, 4, 50, 59,
	2, 65, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68,
	3, 2, 2, 2, 68, 75, 3, 2, 2, 2, 69, 71, 7, 48, 2, 2, 70, 72, 4, 50, 59,
	2, 71, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74,
	3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 69, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2,
	76, 12, 3, 2, 2, 2, 77, 78, 9, 3, 2, 2, 78, 14, 3, 2, 2, 2, 79, 80, 9,
	4, 2, 2, 80, 16, 3, 2, 2, 2, 81, 82, 7, 42, 2, 2, 82, 18, 3, 2, 2, 2, 83,
	84, 7, 43, 2, 2, 84, 20, 3, 2, 2, 2, 85, 86, 7, 45, 2, 2, 86, 22, 3, 2,
	2, 2, 87, 88, 7, 47, 2, 2, 88, 24, 3, 2, 2, 2, 89, 90, 7, 44, 2, 2, 90,
	26, 3, 2, 2, 2, 91, 92, 7, 49, 2, 2, 92, 28, 3, 2, 2, 2, 93, 94, 7, 64,
	2, 2, 94, 30, 3, 2, 2, 2, 95, 96, 7, 62, 2, 2, 96, 32, 3, 2, 2, 2, 97,
	98, 7, 63, 2, 2, 98, 34, 3, 2, 2, 2, 99, 100, 7, 48, 2, 2, 100, 36, 3,
	2, 2, 2, 101, 102, 7, 96, 2, 2, 102, 38, 3, 2, 2, 2, 103, 105, 9, 5, 2,
	2, 104, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106,
	107, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 8, 20, 2, 2, 109, 40,
	3, 2, 2, 2, 12, 2, 45, 49, 53, 58, 62, 67, 73, 75, 106, 3, 8, 2, 2,
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
	"E", "SIGN", "LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "GT",
	"LT", "EQ", "POINT", "POW", "WS",
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
