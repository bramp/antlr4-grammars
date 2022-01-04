// Code generated from beep.g4 by ANTLR 4.9.3. DO NOT EDIT.

package beep

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 72, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	9, 6, 9, 51, 10, 9, 13, 9, 14, 9, 52, 3, 10, 3, 10, 3, 11, 6, 11, 58, 10,
	11, 13, 11, 14, 11, 59, 3, 12, 3, 12, 7, 12, 64, 10, 12, 12, 12, 14, 12,
	67, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 65, 2, 13, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 3, 2, 4, 3,
	2, 50, 59, 4, 2, 12, 12, 15, 15, 2, 74, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 27, 3, 2, 2, 2, 7, 29,
	3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 37, 3, 2, 2, 2, 13, 41, 3, 2, 2, 2,
	15, 45, 3, 2, 2, 2, 17, 50, 3, 2, 2, 2, 19, 54, 3, 2, 2, 2, 21, 57, 3,
	2, 2, 2, 23, 61, 3, 2, 2, 2, 25, 26, 7, 48, 2, 2, 26, 4, 3, 2, 2, 2, 27,
	28, 7, 44, 2, 2, 28, 6, 3, 2, 2, 2, 29, 30, 7, 80, 2, 2, 30, 31, 7, 87,
	2, 2, 31, 32, 7, 78, 2, 2, 32, 8, 3, 2, 2, 2, 33, 34, 7, 71, 2, 2, 34,
	35, 7, 84, 2, 2, 35, 36, 7, 84, 2, 2, 36, 10, 3, 2, 2, 2, 37, 38, 7, 67,
	2, 2, 38, 39, 7, 80, 2, 2, 39, 40, 7, 85, 2, 2, 40, 12, 3, 2, 2, 2, 41,
	42, 7, 84, 2, 2, 42, 43, 7, 82, 2, 2, 43, 44, 7, 91, 2, 2, 44, 14, 3, 2,
	2, 2, 45, 46, 7, 79, 2, 2, 46, 47, 7, 85, 2, 2, 47, 48, 7, 73, 2, 2, 48,
	16, 3, 2, 2, 2, 49, 51, 9, 2, 2, 2, 50, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2,
	2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 18, 3, 2, 2, 2, 54, 55,
	7, 34, 2, 2, 55, 20, 3, 2, 2, 2, 56, 58, 9, 3, 2, 2, 57, 56, 3, 2, 2, 2,
	58, 59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 22, 3,
	2, 2, 2, 61, 65, 5, 21, 11, 2, 62, 64, 11, 2, 2, 2, 63, 62, 3, 2, 2, 2,
	64, 67, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 68, 3,
	2, 2, 2, 67, 65, 3, 2, 2, 2, 68, 69, 7, 71, 2, 2, 69, 70, 7, 80, 2, 2,
	70, 71, 7, 70, 2, 2, 71, 24, 3, 2, 2, 2, 6, 2, 52, 59, 65, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'.'", "'*'", "'NUL'", "'ERR'", "'ANS'", "'RPY'", "'MSG'", "", "' '",
}

var lexerSymbolicNames = []string{
	"", "DOT", "STAR", "NUL", "ERR", "ANS", "RPY", "MSG", "NUMBER", "SP", "CRLF",
	"PAYLOAD_TRAILER",
}

var lexerRuleNames = []string{
	"DOT", "STAR", "NUL", "ERR", "ANS", "RPY", "MSG", "NUMBER", "SP", "CRLF",
	"PAYLOAD_TRAILER",
}

type beepLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewbeepLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *beepLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewbeepLexer(input antlr.CharStream) *beepLexer {
	l := new(beepLexer)
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
	l.GrammarFileName = "beep.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// beepLexer tokens.
const (
	beepLexerDOT             = 1
	beepLexerSTAR            = 2
	beepLexerNUL             = 3
	beepLexerERR             = 4
	beepLexerANS             = 5
	beepLexerRPY             = 6
	beepLexerMSG             = 7
	beepLexerNUMBER          = 8
	beepLexerSP              = 9
	beepLexerCRLF            = 10
	beepLexerPAYLOAD_TRAILER = 11
)
