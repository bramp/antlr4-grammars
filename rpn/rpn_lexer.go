// Code generated from rpn.g4 by ANTLR 4.9.3. DO NOT EDIT.

package rpn

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 140,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 3, 2, 3, 2, 3, 2, 5, 2, 51, 10, 2, 3, 2, 3, 2, 5, 2, 55, 10, 2,
	3, 3, 6, 3, 58, 10, 3, 13, 3, 14, 3, 59, 3, 3, 3, 3, 6, 3, 64, 10, 3, 13,
	3, 14, 3, 65, 5, 3, 68, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6,
	76, 10, 6, 12, 6, 14, 6, 79, 11, 6, 3, 7, 5, 7, 82, 10, 7, 3, 8, 3, 8,
	5, 8, 86, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 6, 23, 135, 10, 23,
	13, 23, 14, 23, 136, 3, 23, 3, 23, 2, 2, 24, 3, 3, 5, 2, 7, 2, 9, 2, 11,
	4, 13, 2, 15, 2, 17, 5, 19, 6, 21, 7, 23, 8, 25, 9, 27, 10, 29, 11, 31,
	12, 33, 13, 35, 14, 37, 15, 39, 16, 41, 17, 43, 18, 45, 19, 3, 2, 6, 4,
	2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 5, 2, 67, 92, 97, 97, 99, 124,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 142, 2, 3, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2,
	2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3,
	2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39,
	3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 3,
	47, 3, 2, 2, 2, 5, 57, 3, 2, 2, 2, 7, 69, 3, 2, 2, 2, 9, 71, 3, 2, 2, 2,
	11, 73, 3, 2, 2, 2, 13, 81, 3, 2, 2, 2, 15, 85, 3, 2, 2, 2, 17, 87, 3,
	2, 2, 2, 19, 89, 3, 2, 2, 2, 21, 91, 3, 2, 2, 2, 23, 93, 3, 2, 2, 2, 25,
	95, 3, 2, 2, 2, 27, 97, 3, 2, 2, 2, 29, 101, 3, 2, 2, 2, 31, 105, 3, 2,
	2, 2, 33, 109, 3, 2, 2, 2, 35, 114, 3, 2, 2, 2, 37, 119, 3, 2, 2, 2, 39,
	124, 3, 2, 2, 2, 41, 127, 3, 2, 2, 2, 43, 131, 3, 2, 2, 2, 45, 134, 3,
	2, 2, 2, 47, 54, 5, 5, 3, 2, 48, 50, 5, 7, 4, 2, 49, 51, 5, 9, 5, 2, 50,
	49, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53, 5, 5, 3,
	2, 53, 55, 3, 2, 2, 2, 54, 48, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 4, 3,
	2, 2, 2, 56, 58, 4, 50, 59, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2,
	59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 67, 3, 2, 2, 2, 61, 63, 7,
	48, 2, 2, 62, 64, 4, 50, 59, 2, 63, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2,
	65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 2, 67, 61, 3,
	2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 6, 3, 2, 2, 2, 69, 70, 9, 2, 2, 2, 70,
	8, 3, 2, 2, 2, 71, 72, 9, 3, 2, 2, 72, 10, 3, 2, 2, 2, 73, 77, 5, 13, 7,
	2, 74, 76, 5, 15, 8, 2, 75, 74, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75,
	3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 12, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2,
	80, 82, 9, 4, 2, 2, 81, 80, 3, 2, 2, 2, 82, 14, 3, 2, 2, 2, 83, 86, 5,
	13, 7, 2, 84, 86, 4, 50, 59, 2, 85, 83, 3, 2, 2, 2, 85, 84, 3, 2, 2, 2,
	86, 16, 3, 2, 2, 2, 87, 88, 7, 96, 2, 2, 88, 18, 3, 2, 2, 2, 89, 90, 7,
	45, 2, 2, 90, 20, 3, 2, 2, 2, 91, 92, 7, 47, 2, 2, 92, 22, 3, 2, 2, 2,
	93, 94, 7, 44, 2, 2, 94, 24, 3, 2, 2, 2, 95, 96, 7, 49, 2, 2, 96, 26, 3,
	2, 2, 2, 97, 98, 7, 101, 2, 2, 98, 99, 7, 113, 2, 2, 99, 100, 7, 117, 2,
	2, 100, 28, 3, 2, 2, 2, 101, 102, 7, 117, 2, 2, 102, 103, 7, 107, 2, 2,
	103, 104, 7, 112, 2, 2, 104, 30, 3, 2, 2, 2, 105, 106, 7, 118, 2, 2, 106,
	107, 7, 99, 2, 2, 107, 108, 7, 112, 2, 2, 108, 32, 3, 2, 2, 2, 109, 110,
	7, 99, 2, 2, 110, 111, 7, 101, 2, 2, 111, 112, 7, 113, 2, 2, 112, 113,
	7, 117, 2, 2, 113, 34, 3, 2, 2, 2, 114, 115, 7, 99, 2, 2, 115, 116, 7,
	117, 2, 2, 116, 117, 7, 107, 2, 2, 117, 118, 7, 112, 2, 2, 118, 36, 3,
	2, 2, 2, 119, 120, 7, 99, 2, 2, 120, 121, 7, 118, 2, 2, 121, 122, 7, 99,
	2, 2, 122, 123, 7, 112, 2, 2, 123, 38, 3, 2, 2, 2, 124, 125, 7, 110, 2,
	2, 125, 126, 7, 112, 2, 2, 126, 40, 3, 2, 2, 2, 127, 128, 7, 110, 2, 2,
	128, 129, 7, 113, 2, 2, 129, 130, 7, 105, 2, 2, 130, 42, 3, 2, 2, 2, 131,
	132, 7, 48, 2, 2, 132, 44, 3, 2, 2, 2, 133, 135, 9, 5, 2, 2, 134, 133,
	3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2,
	2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 8, 23, 2, 2, 139, 46, 3, 2, 2, 2,
	12, 2, 50, 54, 59, 65, 67, 77, 81, 85, 136, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "'^'", "'+'", "'-'", "'*'", "'/'", "'cos'", "'sin'", "'tan'",
	"'acos'", "'asin'", "'atan'", "'ln'", "'log'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "SCIENTIFIC_NUMBER", "VARIABLE", "POW", "PLUS", "MINUS", "TIMES", "DIV",
	"COS", "SIN", "TAN", "ACOS", "ASIN", "ATAN", "LN", "LOG", "POINT", "WS",
}

var lexerRuleNames = []string{
	"SCIENTIFIC_NUMBER", "NUMBER", "E", "SIGN", "VARIABLE", "VALID_ID_START",
	"VALID_ID_CHAR", "POW", "PLUS", "MINUS", "TIMES", "DIV", "COS", "SIN",
	"TAN", "ACOS", "ASIN", "ATAN", "LN", "LOG", "POINT", "WS",
}

type rpnLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewrpnLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *rpnLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewrpnLexer(input antlr.CharStream) *rpnLexer {
	l := new(rpnLexer)
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
	l.GrammarFileName = "rpn.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// rpnLexer tokens.
const (
	rpnLexerSCIENTIFIC_NUMBER = 1
	rpnLexerVARIABLE          = 2
	rpnLexerPOW               = 3
	rpnLexerPLUS              = 4
	rpnLexerMINUS             = 5
	rpnLexerTIMES             = 6
	rpnLexerDIV               = 7
	rpnLexerCOS               = 8
	rpnLexerSIN               = 9
	rpnLexerTAN               = 10
	rpnLexerACOS              = 11
	rpnLexerASIN              = 12
	rpnLexerATAN              = 13
	rpnLexerLN                = 14
	rpnLexerLOG               = 15
	rpnLexerPOINT             = 16
	rpnLexerWS                = 17
)
