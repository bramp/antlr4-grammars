// Code generated from mps.g4 by ANTLR 4.9.3. DO NOT EDIT.

package mps

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 191,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 141, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15,
	7, 15, 147, 10, 15, 12, 15, 14, 15, 150, 11, 15, 3, 16, 3, 16, 7, 16, 154,
	10, 16, 12, 16, 14, 16, 157, 11, 16, 3, 17, 6, 17, 160, 10, 17, 13, 17,
	14, 17, 161, 3, 17, 3, 17, 3, 18, 3, 18, 7, 18, 168, 10, 18, 12, 18, 14,
	18, 171, 11, 18, 3, 18, 5, 18, 174, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 19, 3, 19, 5, 19, 182, 10, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3,
	22, 5, 22, 190, 10, 22, 2, 2, 23, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 2, 39, 2, 41, 2, 43, 2, 3, 2, 9, 6, 2, 71, 71, 73,
	73, 78, 78, 80, 80, 5, 2, 11, 12, 14, 15, 34, 34, 4, 2, 38, 38, 44, 44,
	4, 2, 12, 12, 15, 15, 8, 2, 37, 37, 42, 43, 49, 49, 66, 92, 97, 97, 99,
	124, 4, 2, 45, 48, 50, 59, 4, 2, 70, 71, 102, 103, 2, 202, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 3, 45, 3, 2, 2, 2, 5, 50, 3, 2, 2, 2, 7, 55, 3, 2, 2, 2,
	9, 63, 3, 2, 2, 2, 11, 67, 3, 2, 2, 2, 13, 74, 3, 2, 2, 2, 15, 81, 3, 2,
	2, 2, 17, 88, 3, 2, 2, 2, 19, 97, 3, 2, 2, 2, 21, 106, 3, 2, 2, 2, 23,
	115, 3, 2, 2, 2, 25, 140, 3, 2, 2, 2, 27, 142, 3, 2, 2, 2, 29, 144, 3,
	2, 2, 2, 31, 151, 3, 2, 2, 2, 33, 159, 3, 2, 2, 2, 35, 165, 3, 2, 2, 2,
	37, 181, 3, 2, 2, 2, 39, 183, 3, 2, 2, 2, 41, 185, 3, 2, 2, 2, 43, 189,
	3, 2, 2, 2, 45, 46, 7, 80, 2, 2, 46, 47, 7, 67, 2, 2, 47, 48, 7, 79, 2,
	2, 48, 49, 7, 71, 2, 2, 49, 4, 3, 2, 2, 2, 50, 51, 7, 84, 2, 2, 51, 52,
	7, 81, 2, 2, 52, 53, 7, 89, 2, 2, 53, 54, 7, 85, 2, 2, 54, 6, 3, 2, 2,
	2, 55, 56, 7, 69, 2, 2, 56, 57, 7, 81, 2, 2, 57, 58, 7, 78, 2, 2, 58, 59,
	7, 87, 2, 2, 59, 60, 7, 79, 2, 2, 60, 61, 7, 80, 2, 2, 61, 62, 7, 85, 2,
	2, 62, 8, 3, 2, 2, 2, 63, 64, 7, 84, 2, 2, 64, 65, 7, 74, 2, 2, 65, 66,
	7, 85, 2, 2, 66, 10, 3, 2, 2, 2, 67, 68, 7, 84, 2, 2, 68, 69, 7, 67, 2,
	2, 69, 70, 7, 80, 2, 2, 70, 71, 7, 73, 2, 2, 71, 72, 7, 71, 2, 2, 72, 73,
	7, 85, 2, 2, 73, 12, 3, 2, 2, 2, 74, 75, 7, 68, 2, 2, 75, 76, 7, 81, 2,
	2, 76, 77, 7, 87, 2, 2, 77, 78, 7, 80, 2, 2, 78, 79, 7, 70, 2, 2, 79, 80,
	7, 85, 2, 2, 80, 14, 3, 2, 2, 2, 81, 82, 7, 71, 2, 2, 82, 83, 7, 80, 2,
	2, 83, 84, 7, 70, 2, 2, 84, 85, 7, 67, 2, 2, 85, 86, 7, 86, 2, 2, 86, 87,
	7, 67, 2, 2, 87, 16, 3, 2, 2, 2, 88, 89, 7, 41, 2, 2, 89, 90, 7, 79, 2,
	2, 90, 91, 7, 67, 2, 2, 91, 92, 7, 84, 2, 2, 92, 93, 7, 77, 2, 2, 93, 94,
	7, 71, 2, 2, 94, 95, 7, 84, 2, 2, 95, 96, 7, 41, 2, 2, 96, 18, 3, 2, 2,
	2, 97, 98, 7, 41, 2, 2, 98, 99, 7, 75, 2, 2, 99, 100, 7, 80, 2, 2, 100,
	101, 7, 86, 2, 2, 101, 102, 7, 81, 2, 2, 102, 103, 7, 84, 2, 2, 103, 104,
	7, 73, 2, 2, 104, 105, 7, 41, 2, 2, 105, 20, 3, 2, 2, 2, 106, 107, 7, 41,
	2, 2, 107, 108, 7, 75, 2, 2, 108, 109, 7, 80, 2, 2, 109, 110, 7, 86, 2,
	2, 110, 111, 7, 71, 2, 2, 111, 112, 7, 80, 2, 2, 112, 113, 7, 70, 2, 2,
	113, 114, 7, 41, 2, 2, 114, 22, 3, 2, 2, 2, 115, 116, 7, 72, 2, 2, 116,
	117, 7, 84, 2, 2, 117, 118, 7, 71, 2, 2, 118, 119, 7, 71, 2, 2, 119, 24,
	3, 2, 2, 2, 120, 121, 7, 87, 2, 2, 121, 141, 7, 82, 2, 2, 122, 123, 7,
	78, 2, 2, 123, 141, 7, 81, 2, 2, 124, 125, 7, 72, 2, 2, 125, 141, 7, 90,
	2, 2, 126, 127, 7, 78, 2, 2, 127, 141, 7, 75, 2, 2, 128, 129, 7, 87, 2,
	2, 129, 141, 7, 75, 2, 2, 130, 131, 7, 85, 2, 2, 131, 141, 7, 69, 2, 2,
	132, 133, 7, 72, 2, 2, 133, 141, 7, 84, 2, 2, 134, 135, 7, 68, 2, 2, 135,
	141, 7, 88, 2, 2, 136, 137, 7, 79, 2, 2, 137, 141, 7, 75, 2, 2, 138, 139,
	7, 82, 2, 2, 139, 141, 7, 78, 2, 2, 140, 120, 3, 2, 2, 2, 140, 122, 3,
	2, 2, 2, 140, 124, 3, 2, 2, 2, 140, 126, 3, 2, 2, 2, 140, 128, 3, 2, 2,
	2, 140, 130, 3, 2, 2, 2, 140, 132, 3, 2, 2, 2, 140, 134, 3, 2, 2, 2, 140,
	136, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141, 26, 3, 2, 2, 2, 142, 143, 9,
	2, 2, 2, 143, 28, 3, 2, 2, 2, 144, 148, 5, 39, 20, 2, 145, 147, 5, 37,
	19, 2, 146, 145, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2,
	148, 149, 3, 2, 2, 2, 149, 30, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 155,
	5, 41, 21, 2, 152, 154, 5, 43, 22, 2, 153, 152, 3, 2, 2, 2, 154, 157, 3,
	2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 32, 3, 2, 2,
	2, 157, 155, 3, 2, 2, 2, 158, 160, 9, 3, 2, 2, 159, 158, 3, 2, 2, 2, 160,
	161, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163,
	3, 2, 2, 2, 163, 164, 8, 17, 2, 2, 164, 34, 3, 2, 2, 2, 165, 169, 9, 4,
	2, 2, 166, 168, 10, 5, 2, 2, 167, 166, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2,
	169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2, 171,
	169, 3, 2, 2, 2, 172, 174, 7, 15, 2, 2, 173, 172, 3, 2, 2, 2, 173, 174,
	3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 7, 12, 2, 2, 176, 177, 3, 2,
	2, 2, 177, 178, 8, 18, 2, 2, 178, 36, 3, 2, 2, 2, 179, 182, 5, 39, 20,
	2, 180, 182, 5, 41, 21, 2, 181, 179, 3, 2, 2, 2, 181, 180, 3, 2, 2, 2,
	182, 38, 3, 2, 2, 2, 183, 184, 9, 6, 2, 2, 184, 40, 3, 2, 2, 2, 185, 186,
	9, 7, 2, 2, 186, 42, 3, 2, 2, 2, 187, 190, 5, 41, 21, 2, 188, 190, 9, 8,
	2, 2, 189, 187, 3, 2, 2, 2, 189, 188, 3, 2, 2, 2, 190, 44, 3, 2, 2, 2,
	11, 2, 140, 148, 155, 161, 169, 173, 181, 189, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'NAME'", "'ROWS'", "'COLUMNS'", "'RHS'", "'RANGES'", "'BOUNDS'", "'ENDATA'",
	"''MARKER''", "''INTORG''", "''INTEND''", "'FREE'",
}

var lexerSymbolicNames = []string{
	"", "NAMEINDICATORCARD", "ROWINDICATORCARD", "COLUMNINDICATORCARD", "RHSINDICATORCARD",
	"RANGESINDICATORCARD", "BOUNDSINDICATORCARD", "ENDATAINDICATORCARD", "KEYWORDMARKER",
	"STARTMARKER", "ENDMARKER", "KEYWORDFREE", "BOUNDKEY", "ROWTYPE", "IDENTIFIER",
	"NUMERICALVALUE", "WS", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"NAMEINDICATORCARD", "ROWINDICATORCARD", "COLUMNINDICATORCARD", "RHSINDICATORCARD",
	"RANGESINDICATORCARD", "BOUNDSINDICATORCARD", "ENDATAINDICATORCARD", "KEYWORDMARKER",
	"STARTMARKER", "ENDMARKER", "KEYWORDFREE", "BOUNDKEY", "ROWTYPE", "IDENTIFIER",
	"NUMERICALVALUE", "WS", "LINE_COMMENT", "CHARACTER", "LETTER", "DIGIT",
	"DIGITS",
}

type mpsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewmpsLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *mpsLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewmpsLexer(input antlr.CharStream) *mpsLexer {
	l := new(mpsLexer)
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
	l.GrammarFileName = "mps.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// mpsLexer tokens.
const (
	mpsLexerNAMEINDICATORCARD   = 1
	mpsLexerROWINDICATORCARD    = 2
	mpsLexerCOLUMNINDICATORCARD = 3
	mpsLexerRHSINDICATORCARD    = 4
	mpsLexerRANGESINDICATORCARD = 5
	mpsLexerBOUNDSINDICATORCARD = 6
	mpsLexerENDATAINDICATORCARD = 7
	mpsLexerKEYWORDMARKER       = 8
	mpsLexerSTARTMARKER         = 9
	mpsLexerENDMARKER           = 10
	mpsLexerKEYWORDFREE         = 11
	mpsLexerBOUNDKEY            = 12
	mpsLexerROWTYPE             = 13
	mpsLexerIDENTIFIER          = 14
	mpsLexerNUMERICALVALUE      = 15
	mpsLexerWS                  = 16
	mpsLexerLINE_COMMENT        = 17
)
