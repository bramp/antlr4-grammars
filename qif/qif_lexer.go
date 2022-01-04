// Code generated from qifLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package qif

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 21, 152,
	8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6,
	9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4,
	12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17,
	9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	12, 6, 12, 84, 10, 12, 13, 12, 14, 12, 85, 3, 12, 3, 12, 3, 13, 6, 13,
	91, 10, 13, 13, 13, 14, 13, 92, 3, 13, 3, 13, 6, 13, 97, 10, 13, 13, 13,
	14, 13, 98, 3, 13, 3, 13, 6, 13, 103, 10, 13, 13, 13, 14, 13, 104, 3, 14,
	5, 14, 108, 10, 14, 3, 14, 6, 14, 111, 10, 14, 13, 14, 14, 14, 112, 3,
	14, 3, 14, 6, 14, 117, 10, 14, 13, 14, 14, 14, 118, 5, 14, 121, 10, 14,
	3, 15, 6, 15, 124, 10, 15, 13, 15, 14, 15, 125, 3, 16, 6, 16, 129, 10,
	16, 13, 16, 14, 16, 130, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 6, 18,
	139, 10, 18, 13, 18, 14, 18, 140, 3, 19, 3, 19, 3, 20, 6, 20, 146, 10,
	20, 13, 20, 14, 20, 147, 3, 20, 3, 20, 3, 20, 2, 2, 21, 5, 3, 7, 4, 9,
	5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
	29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 5, 2, 3, 4, 8,
	5, 2, 44, 44, 90, 90, 122, 122, 5, 2, 11, 12, 15, 15, 34, 34, 3, 2, 50,
	59, 4, 2, 46, 46, 50, 59, 4, 2, 12, 12, 15, 15, 6, 2, 34, 34, 50, 59, 67,
	92, 99, 124, 2, 161, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2,
	2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3,
	2, 2, 2, 3, 27, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 3, 31, 3, 2, 2, 2, 3, 33,
	3, 2, 2, 2, 4, 35, 3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 4, 39, 3, 2, 2, 2, 4,
	41, 3, 2, 2, 2, 5, 43, 3, 2, 2, 2, 7, 52, 3, 2, 2, 2, 9, 56, 3, 2, 2, 2,
	11, 58, 3, 2, 2, 2, 13, 62, 3, 2, 2, 2, 15, 66, 3, 2, 2, 2, 17, 70, 3,
	2, 2, 2, 19, 74, 3, 2, 2, 2, 21, 78, 3, 2, 2, 2, 23, 80, 3, 2, 2, 2, 25,
	83, 3, 2, 2, 2, 27, 90, 3, 2, 2, 2, 29, 107, 3, 2, 2, 2, 31, 123, 3, 2,
	2, 2, 33, 128, 3, 2, 2, 2, 35, 135, 3, 2, 2, 2, 37, 138, 3, 2, 2, 2, 39,
	142, 3, 2, 2, 2, 41, 145, 3, 2, 2, 2, 43, 44, 7, 35, 2, 2, 44, 45, 7, 86,
	2, 2, 45, 46, 7, 123, 2, 2, 46, 47, 7, 114, 2, 2, 47, 48, 7, 103, 2, 2,
	48, 49, 7, 60, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 8, 2, 2, 2, 51, 6, 3,
	2, 2, 2, 52, 53, 7, 86, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 8, 3, 2, 2, 55,
	8, 3, 2, 2, 2, 56, 57, 7, 69, 2, 2, 57, 10, 3, 2, 2, 2, 58, 59, 7, 80,
	2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 8, 5, 2, 2, 61, 12, 3, 2, 2, 2, 62, 63,
	7, 79, 2, 2, 63, 64, 3, 2, 2, 2, 64, 65, 8, 6, 2, 2, 65, 14, 3, 2, 2, 2,
	66, 67, 7, 82, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 8, 7, 2, 2, 69, 16, 3,
	2, 2, 2, 70, 71, 7, 78, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 8, 8, 3, 2, 73,
	18, 3, 2, 2, 2, 74, 75, 7, 70, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 8, 9,
	2, 2, 77, 20, 3, 2, 2, 2, 78, 79, 9, 2, 2, 2, 79, 22, 3, 2, 2, 2, 80, 81,
	7, 96, 2, 2, 81, 24, 3, 2, 2, 2, 82, 84, 9, 3, 2, 2, 83, 82, 3, 2, 2, 2,
	84, 85, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87, 3,
	2, 2, 2, 87, 88, 8, 12, 4, 2, 88, 26, 3, 2, 2, 2, 89, 91, 9, 4, 2, 2, 90,
	89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2,
	2, 93, 94, 3, 2, 2, 2, 94, 96, 7, 49, 2, 2, 95, 97, 9, 4, 2, 2, 96, 95,
	3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2,
	99, 100, 3, 2, 2, 2, 100, 102, 7, 49, 2, 2, 101, 103, 9, 4, 2, 2, 102,
	101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105,
	3, 2, 2, 2, 105, 28, 3, 2, 2, 2, 106, 108, 7, 47, 2, 2, 107, 106, 3, 2,
	2, 2, 107, 108, 3, 2, 2, 2, 108, 110, 3, 2, 2, 2, 109, 111, 9, 5, 2, 2,
	110, 109, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112,
	113, 3, 2, 2, 2, 113, 120, 3, 2, 2, 2, 114, 116, 7, 48, 2, 2, 115, 117,
	9, 4, 2, 2, 116, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 116, 3, 2,
	2, 2, 118, 119, 3, 2, 2, 2, 119, 121, 3, 2, 2, 2, 120, 114, 3, 2, 2, 2,
	120, 121, 3, 2, 2, 2, 121, 30, 3, 2, 2, 2, 122, 124, 10, 6, 2, 2, 123,
	122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126,
	3, 2, 2, 2, 126, 32, 3, 2, 2, 2, 127, 129, 9, 6, 2, 2, 128, 127, 3, 2,
	2, 2, 129, 130, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2,
	131, 132, 3, 2, 2, 2, 132, 133, 8, 16, 4, 2, 133, 134, 8, 16, 5, 2, 134,
	34, 3, 2, 2, 2, 135, 136, 7, 93, 2, 2, 136, 36, 3, 2, 2, 2, 137, 139, 9,
	7, 2, 2, 138, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 138, 3, 2, 2,
	2, 140, 141, 3, 2, 2, 2, 141, 38, 3, 2, 2, 2, 142, 143, 7, 95, 2, 2, 143,
	40, 3, 2, 2, 2, 144, 146, 9, 6, 2, 2, 145, 144, 3, 2, 2, 2, 146, 147, 3,
	2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 149, 3, 2, 2,
	2, 149, 150, 8, 20, 4, 2, 150, 151, 8, 20, 5, 2, 151, 42, 3, 2, 2, 2, 17,
	2, 3, 4, 85, 92, 98, 104, 107, 112, 118, 120, 125, 130, 140, 147, 6, 7,
	3, 2, 7, 4, 2, 8, 2, 2, 6, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "LINETEXT", "ACCTCATE",
}

var lexerLiteralNames = []string{
	"", "'!Type:'", "'T'", "'C'", "'N'", "'M'", "'P'", "'L'", "'D'", "", "'^'",
	"", "", "", "", "", "'['", "", "']'",
}

var lexerSymbolicNames = []string{
	"", "TYPE", "T", "C", "N", "M", "P", "L", "D", "STATE", "EOR", "WS", "DATE",
	"NUM", "TEXT", "EOL", "LB", "ACCNTCATNAME", "RB", "EOL2",
}

var lexerRuleNames = []string{
	"TYPE", "T", "C", "N", "M", "P", "L", "D", "STATE", "EOR", "WS", "DATE",
	"NUM", "TEXT", "EOL", "LB", "ACCNTCATNAME", "RB", "EOL2",
}

type qifLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewqifLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *qifLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewqifLexer(input antlr.CharStream) *qifLexer {
	l := new(qifLexer)
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
	l.GrammarFileName = "qifLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// qifLexer tokens.
const (
	qifLexerTYPE         = 1
	qifLexerT            = 2
	qifLexerC            = 3
	qifLexerN            = 4
	qifLexerM            = 5
	qifLexerP            = 6
	qifLexerL            = 7
	qifLexerD            = 8
	qifLexerSTATE        = 9
	qifLexerEOR          = 10
	qifLexerWS           = 11
	qifLexerDATE         = 12
	qifLexerNUM          = 13
	qifLexerTEXT         = 14
	qifLexerEOL          = 15
	qifLexerLB           = 16
	qifLexerACCNTCATNAME = 17
	qifLexerRB           = 18
	qifLexerEOL2         = 19
)

// qifLexer modes.
const (
	qifLexerLINETEXT = iota + 1
	qifLexerACCTCATE
)
