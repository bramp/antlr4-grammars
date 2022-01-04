// Code generated from SwiftFinLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package swiftfin

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 151,
	8, 1, 8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5,
	4, 6, 9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9,
	11, 4, 12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16,
	4, 17, 9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4,
	22, 9, 22, 4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27,
	9, 27, 4, 28, 9, 28, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 13, 5, 13, 107, 10, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 2, 2, 29, 6, 3, 8, 4, 10, 5,
	12, 2, 14, 6, 16, 7, 18, 8, 20, 9, 22, 10, 24, 11, 26, 12, 28, 2, 30, 2,
	32, 2, 34, 2, 36, 2, 38, 2, 40, 2, 42, 2, 44, 13, 46, 14, 48, 15, 50, 16,
	52, 17, 54, 18, 56, 19, 58, 20, 6, 2, 3, 4, 5, 3, 3, 2, 60, 60, 2, 141,
	2, 6, 3, 2, 2, 2, 2, 8, 3, 2, 2, 2, 2, 10, 3, 2, 2, 2, 2, 14, 3, 2, 2,
	2, 2, 16, 3, 2, 2, 2, 2, 18, 3, 2, 2, 2, 2, 20, 3, 2, 2, 2, 2, 22, 3, 2,
	2, 2, 2, 24, 3, 2, 2, 2, 2, 26, 3, 2, 2, 2, 3, 40, 3, 2, 2, 2, 3, 42, 3,
	2, 2, 2, 3, 44, 3, 2, 2, 2, 3, 46, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 4, 50,
	3, 2, 2, 2, 4, 52, 3, 2, 2, 2, 4, 54, 3, 2, 2, 2, 5, 56, 3, 2, 2, 2, 5,
	58, 3, 2, 2, 2, 6, 60, 3, 2, 2, 2, 8, 66, 3, 2, 2, 2, 10, 72, 3, 2, 2,
	2, 12, 78, 3, 2, 2, 2, 14, 82, 3, 2, 2, 2, 16, 87, 3, 2, 2, 2, 18, 91,
	3, 2, 2, 2, 20, 97, 3, 2, 2, 2, 22, 99, 3, 2, 2, 2, 24, 101, 3, 2, 2, 2,
	26, 103, 3, 2, 2, 2, 28, 106, 3, 2, 2, 2, 30, 110, 3, 2, 2, 2, 32, 112,
	3, 2, 2, 2, 34, 114, 3, 2, 2, 2, 36, 116, 3, 2, 2, 2, 38, 118, 3, 2, 2,
	2, 40, 120, 3, 2, 2, 2, 42, 125, 3, 2, 2, 2, 44, 130, 3, 2, 2, 2, 46, 132,
	3, 2, 2, 2, 48, 134, 3, 2, 2, 2, 50, 139, 3, 2, 2, 2, 52, 141, 3, 2, 2,
	2, 54, 143, 3, 2, 2, 2, 56, 145, 3, 2, 2, 2, 58, 149, 3, 2, 2, 2, 60, 61,
	5, 30, 14, 2, 61, 62, 7, 51, 2, 2, 62, 63, 5, 34, 16, 2, 63, 64, 3, 2,
	2, 2, 64, 65, 8, 2, 2, 2, 65, 7, 3, 2, 2, 2, 66, 67, 5, 30, 14, 2, 67,
	68, 7, 52, 2, 2, 68, 69, 5, 34, 16, 2, 69, 70, 3, 2, 2, 2, 70, 71, 8, 3,
	2, 2, 71, 9, 3, 2, 2, 2, 72, 73, 5, 30, 14, 2, 73, 74, 7, 53, 2, 2, 74,
	75, 5, 34, 16, 2, 75, 76, 3, 2, 2, 2, 76, 77, 8, 4, 3, 2, 77, 11, 3, 2,
	2, 2, 78, 79, 5, 30, 14, 2, 79, 80, 7, 54, 2, 2, 80, 81, 5, 34, 16, 2,
	81, 13, 3, 2, 2, 2, 82, 83, 5, 12, 5, 2, 83, 84, 5, 28, 13, 2, 84, 85,
	3, 2, 2, 2, 85, 86, 8, 6, 4, 2, 86, 15, 3, 2, 2, 2, 87, 88, 5, 12, 5, 2,
	88, 89, 3, 2, 2, 2, 89, 90, 8, 7, 3, 2, 90, 17, 3, 2, 2, 2, 91, 92, 5,
	30, 14, 2, 92, 93, 7, 55, 2, 2, 93, 94, 5, 34, 16, 2, 94, 95, 3, 2, 2,
	2, 95, 96, 8, 8, 3, 2, 96, 19, 3, 2, 2, 2, 97, 98, 5, 30, 14, 2, 98, 21,
	3, 2, 2, 2, 99, 100, 5, 32, 15, 2, 100, 23, 3, 2, 2, 2, 101, 102, 5, 34,
	16, 2, 102, 25, 3, 2, 2, 2, 103, 104, 5, 28, 13, 2, 104, 27, 3, 2, 2, 2,
	105, 107, 7, 15, 2, 2, 106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107,
	108, 3, 2, 2, 2, 108, 109, 7, 12, 2, 2, 109, 29, 3, 2, 2, 2, 110, 111,
	7, 125, 2, 2, 111, 31, 3, 2, 2, 2, 112, 113, 7, 127, 2, 2, 113, 33, 3,
	2, 2, 2, 114, 115, 7, 60, 2, 2, 115, 35, 3, 2, 2, 2, 116, 117, 7, 47, 2,
	2, 117, 37, 3, 2, 2, 2, 118, 119, 11, 2, 2, 2, 119, 39, 3, 2, 2, 2, 120,
	121, 5, 30, 14, 2, 121, 122, 3, 2, 2, 2, 122, 123, 8, 19, 5, 2, 123, 124,
	8, 19, 3, 2, 124, 41, 3, 2, 2, 2, 125, 126, 5, 32, 15, 2, 126, 127, 3,
	2, 2, 2, 127, 128, 8, 20, 6, 2, 128, 129, 8, 20, 7, 2, 129, 43, 3, 2, 2,
	2, 130, 131, 5, 34, 16, 2, 131, 45, 3, 2, 2, 2, 132, 133, 10, 2, 2, 2,
	133, 47, 3, 2, 2, 2, 134, 135, 5, 36, 17, 2, 135, 136, 5, 32, 15, 2, 136,
	137, 3, 2, 2, 2, 137, 138, 8, 23, 7, 2, 138, 49, 3, 2, 2, 2, 139, 140,
	5, 34, 16, 2, 140, 51, 3, 2, 2, 2, 141, 142, 5, 28, 13, 2, 142, 53, 3,
	2, 2, 2, 143, 144, 10, 2, 2, 2, 144, 55, 3, 2, 2, 2, 145, 146, 5, 32, 15,
	2, 146, 147, 3, 2, 2, 2, 147, 148, 8, 27, 7, 2, 148, 57, 3, 2, 2, 2, 149,
	150, 5, 38, 18, 2, 150, 59, 3, 2, 2, 2, 7, 2, 3, 4, 5, 106, 8, 7, 5, 2,
	7, 3, 2, 7, 4, 2, 9, 9, 2, 9, 10, 2, 6, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "InsideMaps", "InsideB4", "InsideValue",
}

var lexerLiteralNames []string

var lexerSymbolicNames = []string{
	"", "BLOCK1", "BLOCK2", "BLOCK3", "BLOCK4_A", "BLOCK4_B", "BLOCK5", "LBRACE",
	"RBRACE", "COLON", "CRLF", "M_COLON", "M_VALUE", "B4_END", "B4_COLON",
	"B4_CRLF", "B4_VALUE", "V_END", "V_VALUE",
}

var lexerRuleNames = []string{
	"BLOCK1", "BLOCK2", "BLOCK3", "Block4", "BLOCK4_A", "BLOCK4_B", "BLOCK5",
	"LBRACE", "RBRACE", "COLON", "CRLF", "Crlf", "LBrace", "RBrace", "Colon",
	"Minus", "Any", "M_LBRACE", "M_RBRACE", "M_COLON", "M_VALUE", "B4_END",
	"B4_COLON", "B4_CRLF", "B4_VALUE", "V_END", "V_VALUE",
}

type SwiftFinLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewSwiftFinLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *SwiftFinLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSwiftFinLexer(input antlr.CharStream) *SwiftFinLexer {
	l := new(SwiftFinLexer)
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
	l.GrammarFileName = "SwiftFinLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SwiftFinLexer tokens.
const (
	SwiftFinLexerBLOCK1   = 1
	SwiftFinLexerBLOCK2   = 2
	SwiftFinLexerBLOCK3   = 3
	SwiftFinLexerBLOCK4_A = 4
	SwiftFinLexerBLOCK4_B = 5
	SwiftFinLexerBLOCK5   = 6
	SwiftFinLexerLBRACE   = 7
	SwiftFinLexerRBRACE   = 8
	SwiftFinLexerCOLON    = 9
	SwiftFinLexerCRLF     = 10
	SwiftFinLexerM_COLON  = 11
	SwiftFinLexerM_VALUE  = 12
	SwiftFinLexerB4_END   = 13
	SwiftFinLexerB4_COLON = 14
	SwiftFinLexerB4_CRLF  = 15
	SwiftFinLexerB4_VALUE = 16
	SwiftFinLexerV_END    = 17
	SwiftFinLexerV_VALUE  = 18
)

// SwiftFinLexer modes.
const (
	SwiftFinLexerInsideMaps = iota + 1
	SwiftFinLexerInsideB4
	SwiftFinLexerInsideValue
)
