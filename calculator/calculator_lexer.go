// Code generated from calculator.g4 by ANTLR 4.9.3. DO NOT EDIT.

package calculator

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 191,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 7, 26, 142, 10, 26, 12,
	26, 14, 26, 145, 11, 26, 3, 27, 5, 27, 148, 10, 27, 3, 28, 3, 28, 5, 28,
	152, 10, 28, 3, 29, 3, 29, 3, 29, 5, 29, 157, 10, 29, 3, 29, 5, 29, 160,
	10, 29, 3, 29, 3, 29, 5, 29, 164, 10, 29, 3, 30, 6, 30, 167, 10, 30, 13,
	30, 14, 30, 168, 3, 30, 3, 30, 6, 30, 173, 10, 30, 13, 30, 14, 30, 174,
	5, 30, 177, 10, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 6,
	34, 186, 10, 34, 13, 34, 14, 34, 187, 3, 34, 3, 34, 2, 2, 35, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 2, 55, 2, 57, 28, 59, 2, 61, 2,
	63, 2, 65, 2, 67, 29, 3, 2, 5, 5, 2, 67, 92, 97, 97, 99, 124, 4, 2, 45,
	45, 47, 47, 5, 2, 11, 12, 15, 15, 34, 34, 2, 193, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2,
	2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3,
	2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 3, 69, 3, 2, 2, 2, 5,
	73, 3, 2, 2, 2, 7, 77, 3, 2, 2, 2, 9, 81, 3, 2, 2, 2, 11, 86, 3, 2, 2,
	2, 13, 91, 3, 2, 2, 2, 15, 96, 3, 2, 2, 2, 17, 99, 3, 2, 2, 2, 19, 103,
	3, 2, 2, 2, 21, 108, 3, 2, 2, 2, 23, 110, 3, 2, 2, 2, 25, 112, 3, 2, 2,
	2, 27, 114, 3, 2, 2, 2, 29, 116, 3, 2, 2, 2, 31, 118, 3, 2, 2, 2, 33, 120,
	3, 2, 2, 2, 35, 122, 3, 2, 2, 2, 37, 124, 3, 2, 2, 2, 39, 126, 3, 2, 2,
	2, 41, 128, 3, 2, 2, 2, 43, 130, 3, 2, 2, 2, 45, 132, 3, 2, 2, 2, 47, 135,
	3, 2, 2, 2, 49, 137, 3, 2, 2, 2, 51, 139, 3, 2, 2, 2, 53, 147, 3, 2, 2,
	2, 55, 151, 3, 2, 2, 2, 57, 153, 3, 2, 2, 2, 59, 166, 3, 2, 2, 2, 61, 178,
	3, 2, 2, 2, 63, 180, 3, 2, 2, 2, 65, 182, 3, 2, 2, 2, 67, 185, 3, 2, 2,
	2, 69, 70, 7, 101, 2, 2, 70, 71, 7, 113, 2, 2, 71, 72, 7, 117, 2, 2, 72,
	4, 3, 2, 2, 2, 73, 74, 7, 117, 2, 2, 74, 75, 7, 107, 2, 2, 75, 76, 7, 112,
	2, 2, 76, 6, 3, 2, 2, 2, 77, 78, 7, 118, 2, 2, 78, 79, 7, 99, 2, 2, 79,
	80, 7, 112, 2, 2, 80, 8, 3, 2, 2, 2, 81, 82, 7, 99, 2, 2, 82, 83, 7, 101,
	2, 2, 83, 84, 7, 113, 2, 2, 84, 85, 7, 117, 2, 2, 85, 10, 3, 2, 2, 2, 86,
	87, 7, 99, 2, 2, 87, 88, 7, 117, 2, 2, 88, 89, 7, 107, 2, 2, 89, 90, 7,
	112, 2, 2, 90, 12, 3, 2, 2, 2, 91, 92, 7, 99, 2, 2, 92, 93, 7, 118, 2,
	2, 93, 94, 7, 99, 2, 2, 94, 95, 7, 112, 2, 2, 95, 14, 3, 2, 2, 2, 96, 97,
	7, 110, 2, 2, 97, 98, 7, 112, 2, 2, 98, 16, 3, 2, 2, 2, 99, 100, 7, 110,
	2, 2, 100, 101, 7, 113, 2, 2, 101, 102, 7, 105, 2, 2, 102, 18, 3, 2, 2,
	2, 103, 104, 7, 117, 2, 2, 104, 105, 7, 115, 2, 2, 105, 106, 7, 116, 2,
	2, 106, 107, 7, 118, 2, 2, 107, 20, 3, 2, 2, 2, 108, 109, 7, 42, 2, 2,
	109, 22, 3, 2, 2, 2, 110, 111, 7, 43, 2, 2, 111, 24, 3, 2, 2, 2, 112, 113,
	7, 45, 2, 2, 113, 26, 3, 2, 2, 2, 114, 115, 7, 47, 2, 2, 115, 28, 3, 2,
	2, 2, 116, 117, 7, 44, 2, 2, 117, 30, 3, 2, 2, 2, 118, 119, 7, 49, 2, 2,
	119, 32, 3, 2, 2, 2, 120, 121, 7, 64, 2, 2, 121, 34, 3, 2, 2, 2, 122, 123,
	7, 62, 2, 2, 123, 36, 3, 2, 2, 2, 124, 125, 7, 63, 2, 2, 125, 38, 3, 2,
	2, 2, 126, 127, 7, 46, 2, 2, 127, 40, 3, 2, 2, 2, 128, 129, 7, 48, 2, 2,
	129, 42, 3, 2, 2, 2, 130, 131, 7, 96, 2, 2, 131, 44, 3, 2, 2, 2, 132, 133,
	7, 114, 2, 2, 133, 134, 7, 107, 2, 2, 134, 46, 3, 2, 2, 2, 135, 136, 5,
	63, 32, 2, 136, 48, 3, 2, 2, 2, 137, 138, 7, 107, 2, 2, 138, 50, 3, 2,
	2, 2, 139, 143, 5, 53, 27, 2, 140, 142, 5, 55, 28, 2, 141, 140, 3, 2, 2,
	2, 142, 145, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144,
	52, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 146, 148, 9, 2, 2, 2, 147, 146, 3,
	2, 2, 2, 148, 54, 3, 2, 2, 2, 149, 152, 5, 53, 27, 2, 150, 152, 4, 50,
	59, 2, 151, 149, 3, 2, 2, 2, 151, 150, 3, 2, 2, 2, 152, 56, 3, 2, 2, 2,
	153, 163, 5, 59, 30, 2, 154, 157, 5, 61, 31, 2, 155, 157, 5, 63, 32, 2,
	156, 154, 3, 2, 2, 2, 156, 155, 3, 2, 2, 2, 157, 159, 3, 2, 2, 2, 158,
	160, 5, 65, 33, 2, 159, 158, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161,
	3, 2, 2, 2, 161, 162, 5, 59, 30, 2, 162, 164, 3, 2, 2, 2, 163, 156, 3,
	2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 58, 3, 2, 2, 2, 165, 167, 4, 50, 59,
	2, 166, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 168,
	169, 3, 2, 2, 2, 169, 176, 3, 2, 2, 2, 170, 172, 7, 48, 2, 2, 171, 173,
	4, 50, 59, 2, 172, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 172, 3,
	2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177, 3, 2, 2, 2, 176, 170, 3, 2, 2,
	2, 176, 177, 3, 2, 2, 2, 177, 60, 3, 2, 2, 2, 178, 179, 7, 71, 2, 2, 179,
	62, 3, 2, 2, 2, 180, 181, 7, 103, 2, 2, 181, 64, 3, 2, 2, 2, 182, 183,
	9, 3, 2, 2, 183, 66, 3, 2, 2, 2, 184, 186, 9, 4, 2, 2, 185, 184, 3, 2,
	2, 2, 186, 187, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2,
	188, 189, 3, 2, 2, 2, 189, 190, 8, 34, 2, 2, 190, 68, 3, 2, 2, 2, 13, 2,
	143, 147, 151, 156, 159, 163, 168, 174, 176, 187, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'cos'", "'sin'", "'tan'", "'acos'", "'asin'", "'atan'", "'ln'", "'log'",
	"'sqrt'", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "'>'", "'<'", "'='",
	"','", "'.'", "'^'", "'pi'", "", "'i'",
}

var lexerSymbolicNames = []string{
	"", "COS", "SIN", "TAN", "ACOS", "ASIN", "ATAN", "LN", "LOG", "SQRT", "LPAREN",
	"RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "GT", "LT", "EQ", "COMMA", "POINT",
	"POW", "PI", "EULER", "I", "VARIABLE", "SCIENTIFIC_NUMBER", "WS",
}

var lexerRuleNames = []string{
	"COS", "SIN", "TAN", "ACOS", "ASIN", "ATAN", "LN", "LOG", "SQRT", "LPAREN",
	"RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "GT", "LT", "EQ", "COMMA", "POINT",
	"POW", "PI", "EULER", "I", "VARIABLE", "VALID_ID_START", "VALID_ID_CHAR",
	"SCIENTIFIC_NUMBER", "NUMBER", "E1", "E2", "SIGN", "WS",
}

type calculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewcalculatorLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *calculatorLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewcalculatorLexer(input antlr.CharStream) *calculatorLexer {
	l := new(calculatorLexer)
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
	l.GrammarFileName = "calculator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// calculatorLexer tokens.
const (
	calculatorLexerCOS               = 1
	calculatorLexerSIN               = 2
	calculatorLexerTAN               = 3
	calculatorLexerACOS              = 4
	calculatorLexerASIN              = 5
	calculatorLexerATAN              = 6
	calculatorLexerLN                = 7
	calculatorLexerLOG               = 8
	calculatorLexerSQRT              = 9
	calculatorLexerLPAREN            = 10
	calculatorLexerRPAREN            = 11
	calculatorLexerPLUS              = 12
	calculatorLexerMINUS             = 13
	calculatorLexerTIMES             = 14
	calculatorLexerDIV               = 15
	calculatorLexerGT                = 16
	calculatorLexerLT                = 17
	calculatorLexerEQ                = 18
	calculatorLexerCOMMA             = 19
	calculatorLexerPOINT             = 20
	calculatorLexerPOW               = 21
	calculatorLexerPI                = 22
	calculatorLexerEULER             = 23
	calculatorLexerI                 = 24
	calculatorLexerVARIABLE          = 25
	calculatorLexerSCIENTIFIC_NUMBER = 26
	calculatorLexerWS                = 27
)
