// Code generated from SUOKIF.g4 by ANTLR 4.9.3. DO NOT EDIT.

package suokif

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 165,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3,
	10, 5, 10, 83, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 89, 10, 11, 3,
	12, 3, 12, 7, 12, 93, 10, 12, 12, 12, 14, 12, 96, 11, 12, 3, 13, 3, 13,
	7, 13, 100, 10, 13, 12, 13, 14, 13, 103, 11, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 111, 10, 14, 3, 15, 5, 15, 114, 10, 15, 3, 15,
	6, 15, 117, 10, 15, 13, 15, 14, 15, 118, 3, 15, 3, 15, 6, 15, 123, 10,
	15, 13, 15, 14, 15, 124, 5, 15, 127, 10, 15, 3, 15, 5, 15, 130, 10, 15,
	3, 16, 3, 16, 5, 16, 134, 10, 16, 3, 16, 6, 16, 137, 10, 16, 13, 16, 14,
	16, 138, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 7, 18, 147, 10, 18,
	12, 18, 14, 18, 150, 11, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 2, 2, 25, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 2, 15, 2, 17, 2, 19, 2, 21, 2, 23, 8, 25,
	9, 27, 10, 29, 11, 31, 2, 33, 12, 35, 13, 37, 14, 39, 15, 41, 16, 43, 17,
	45, 18, 47, 19, 3, 2, 9, 3, 2, 67, 92, 3, 2, 99, 124, 3, 2, 50, 59, 4,
	2, 47, 47, 97, 97, 4, 2, 36, 36, 94, 94, 5, 2, 11, 13, 15, 15, 34, 34,
	4, 2, 12, 12, 15, 15, 2, 173, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 3, 49, 3, 2,
	2, 2, 5, 53, 3, 2, 2, 2, 7, 57, 3, 2, 2, 2, 9, 60, 3, 2, 2, 2, 11, 67,
	3, 2, 2, 2, 13, 74, 3, 2, 2, 2, 15, 76, 3, 2, 2, 2, 17, 78, 3, 2, 2, 2,
	19, 82, 3, 2, 2, 2, 21, 88, 3, 2, 2, 2, 23, 90, 3, 2, 2, 2, 25, 97, 3,
	2, 2, 2, 27, 110, 3, 2, 2, 2, 29, 113, 3, 2, 2, 2, 31, 131, 3, 2, 2, 2,
	33, 140, 3, 2, 2, 2, 35, 144, 3, 2, 2, 2, 37, 153, 3, 2, 2, 2, 39, 155,
	3, 2, 2, 2, 41, 157, 3, 2, 2, 2, 43, 159, 3, 2, 2, 2, 45, 161, 3, 2, 2,
	2, 47, 163, 3, 2, 2, 2, 49, 50, 7, 112, 2, 2, 50, 51, 7, 113, 2, 2, 51,
	52, 7, 118, 2, 2, 52, 4, 3, 2, 2, 2, 53, 54, 7, 99, 2, 2, 54, 55, 7, 112,
	2, 2, 55, 56, 7, 102, 2, 2, 56, 6, 3, 2, 2, 2, 57, 58, 7, 113, 2, 2, 58,
	59, 7, 116, 2, 2, 59, 8, 3, 2, 2, 2, 60, 61, 7, 104, 2, 2, 61, 62, 7, 113,
	2, 2, 62, 63, 7, 116, 2, 2, 63, 64, 7, 99, 2, 2, 64, 65, 7, 110, 2, 2,
	65, 66, 7, 110, 2, 2, 66, 10, 3, 2, 2, 2, 67, 68, 7, 103, 2, 2, 68, 69,
	7, 122, 2, 2, 69, 70, 7, 107, 2, 2, 70, 71, 7, 117, 2, 2, 71, 72, 7, 118,
	2, 2, 72, 73, 7, 117, 2, 2, 73, 12, 3, 2, 2, 2, 74, 75, 9, 2, 2, 2, 75,
	14, 3, 2, 2, 2, 76, 77, 9, 3, 2, 2, 77, 16, 3, 2, 2, 2, 78, 79, 9, 4, 2,
	2, 79, 18, 3, 2, 2, 2, 80, 83, 5, 13, 7, 2, 81, 83, 5, 15, 8, 2, 82, 80,
	3, 2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 20, 3, 2, 2, 2, 84, 89, 5, 13, 7, 2,
	85, 89, 5, 15, 8, 2, 86, 89, 5, 17, 9, 2, 87, 89, 9, 5, 2, 2, 88, 84, 3,
	2, 2, 2, 88, 85, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 87, 3, 2, 2, 2, 89,
	22, 3, 2, 2, 2, 90, 94, 5, 19, 10, 2, 91, 93, 5, 21, 11, 2, 92, 91, 3,
	2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95,
	24, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 101, 7, 36, 2, 2, 98, 100, 10,
	6, 2, 2, 99, 98, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2,
	101, 102, 3, 2, 2, 2, 102, 104, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104,
	105, 7, 36, 2, 2, 105, 26, 3, 2, 2, 2, 106, 107, 7, 65, 2, 2, 107, 111,
	5, 23, 12, 2, 108, 109, 7, 66, 2, 2, 109, 111, 5, 23, 12, 2, 110, 106,
	3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 28, 3, 2, 2, 2, 112, 114, 7, 47,
	2, 2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 116, 3, 2, 2, 2,
	115, 117, 5, 17, 9, 2, 116, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118,
	116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 126, 3, 2, 2, 2, 120, 122,
	7, 48, 2, 2, 121, 123, 5, 17, 9, 2, 122, 121, 3, 2, 2, 2, 123, 124, 3,
	2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2,
	2, 126, 120, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2, 128,
	130, 5, 31, 16, 2, 129, 128, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 30,
	3, 2, 2, 2, 131, 133, 7, 103, 2, 2, 132, 134, 7, 47, 2, 2, 133, 132, 3,
	2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 136, 3, 2, 2, 2, 135, 137, 5, 17, 9,
	2, 136, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138,
	139, 3, 2, 2, 2, 139, 32, 3, 2, 2, 2, 140, 141, 9, 7, 2, 2, 141, 142, 3,
	2, 2, 2, 142, 143, 8, 17, 2, 2, 143, 34, 3, 2, 2, 2, 144, 148, 7, 61, 2,
	2, 145, 147, 10, 8, 2, 2, 146, 145, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148,
	146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 151, 3, 2, 2, 2, 150, 148,
	3, 2, 2, 2, 151, 152, 8, 18, 2, 2, 152, 36, 3, 2, 2, 2, 153, 154, 7, 42,
	2, 2, 154, 38, 3, 2, 2, 2, 155, 156, 7, 43, 2, 2, 156, 40, 3, 2, 2, 2,
	157, 158, 7, 63, 2, 2, 158, 42, 3, 2, 2, 2, 159, 160, 7, 64, 2, 2, 160,
	44, 3, 2, 2, 2, 161, 162, 7, 62, 2, 2, 162, 46, 3, 2, 2, 2, 163, 164, 7,
	65, 2, 2, 164, 48, 3, 2, 2, 2, 16, 2, 82, 88, 94, 101, 110, 113, 118, 124,
	126, 129, 133, 138, 148, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'not'", "'and'", "'or'", "'forall'", "'exists'", "", "", "", "", "",
	"", "'('", "')'", "'='", "'>'", "'<'", "'?'",
}

var lexerSymbolicNames = []string{
	"", "NOT", "AND", "OR", "FORALL", "EXISTS", "WORD", "STRING", "VARIABLE",
	"NUMBER", "WHITE", "COMMENT", "LPAREN", "RPAREN", "ASSIGN", "GT", "LT",
	"QUESTION",
}

var lexerRuleNames = []string{
	"NOT", "AND", "OR", "FORALL", "EXISTS", "UPPER", "LOWER", "DIGIT", "INITIALCHAR",
	"WORDCHAR", "WORD", "STRING", "VARIABLE", "NUMBER", "EXPONENT", "WHITE",
	"COMMENT", "LPAREN", "RPAREN", "ASSIGN", "GT", "LT", "QUESTION",
}

type SUOKIFLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewSUOKIFLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *SUOKIFLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSUOKIFLexer(input antlr.CharStream) *SUOKIFLexer {
	l := new(SUOKIFLexer)
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
	l.GrammarFileName = "SUOKIF.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SUOKIFLexer tokens.
const (
	SUOKIFLexerNOT      = 1
	SUOKIFLexerAND      = 2
	SUOKIFLexerOR       = 3
	SUOKIFLexerFORALL   = 4
	SUOKIFLexerEXISTS   = 5
	SUOKIFLexerWORD     = 6
	SUOKIFLexerSTRING   = 7
	SUOKIFLexerVARIABLE = 8
	SUOKIFLexerNUMBER   = 9
	SUOKIFLexerWHITE    = 10
	SUOKIFLexerCOMMENT  = 11
	SUOKIFLexerLPAREN   = 12
	SUOKIFLexerRPAREN   = 13
	SUOKIFLexerASSIGN   = 14
	SUOKIFLexerGT       = 15
	SUOKIFLexerLT       = 16
	SUOKIFLexerQUESTION = 17
)
