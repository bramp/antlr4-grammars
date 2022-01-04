// Code generated from doiurl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package doiurl

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 70, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 5, 7, 45, 10, 7, 3, 8, 3, 8,
	3, 8, 5, 8, 50, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 6, 14, 65, 10, 14, 13, 14, 14, 14,
	66, 3, 14, 3, 14, 2, 2, 15, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 2,
	17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 27, 9, 3, 2, 8, 8, 2, 38, 38, 40, 40,
	45, 46, 60, 61, 63, 63, 66, 66, 7, 2, 35, 35, 41, 44, 47, 48, 97, 97, 128,
	128, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 4, 2, 50, 59, 67, 72, 5, 2, 11,
	12, 15, 15, 34, 34, 2, 68, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 5, 31, 3, 2, 2, 2, 7, 33, 3, 2, 2, 2, 9,
	35, 3, 2, 2, 2, 11, 39, 3, 2, 2, 2, 13, 44, 3, 2, 2, 2, 15, 49, 3, 2, 2,
	2, 17, 51, 3, 2, 2, 2, 19, 55, 3, 2, 2, 2, 21, 57, 3, 2, 2, 2, 23, 59,
	3, 2, 2, 2, 25, 61, 3, 2, 2, 2, 27, 64, 3, 2, 2, 2, 29, 30, 7, 60, 2, 2,
	30, 4, 3, 2, 2, 2, 31, 32, 7, 65, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34, 7, 37,
	2, 2, 34, 8, 3, 2, 2, 2, 35, 36, 7, 102, 2, 2, 36, 37, 7, 113, 2, 2, 37,
	38, 7, 107, 2, 2, 38, 10, 3, 2, 2, 2, 39, 40, 7, 49, 2, 2, 40, 12, 3, 2,
	2, 2, 41, 45, 5, 15, 8, 2, 42, 45, 5, 17, 9, 2, 43, 45, 9, 2, 2, 2, 44,
	41, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 43, 3, 2, 2, 2, 45, 14, 3, 2, 2,
	2, 46, 50, 5, 21, 11, 2, 47, 50, 5, 23, 12, 2, 48, 50, 5, 19, 10, 2, 49,
	46, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 48, 3, 2, 2, 2, 50, 16, 3, 2, 2,
	2, 51, 52, 7, 39, 2, 2, 52, 53, 5, 25, 13, 2, 53, 54, 5, 25, 13, 2, 54,
	18, 3, 2, 2, 2, 55, 56, 9, 3, 2, 2, 56, 20, 3, 2, 2, 2, 57, 58, 9, 4, 2,
	2, 58, 22, 3, 2, 2, 2, 59, 60, 9, 5, 2, 2, 60, 24, 3, 2, 2, 2, 61, 62,
	9, 6, 2, 2, 62, 26, 3, 2, 2, 2, 63, 65, 9, 7, 2, 2, 64, 63, 3, 2, 2, 2,
	65, 66, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3,
	2, 2, 2, 68, 69, 8, 14, 2, 2, 69, 28, 3, 2, 2, 2, 6, 2, 44, 49, 66, 3,
	8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "':'", "'?'", "'#'", "'doi'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "PCHAR", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "PCHAR", "UNRESERVED", "ESCAPED",
	"MARK", "ALPHA", "DIGIT", "HEXDIG", "WS",
}

type doiurlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewdoiurlLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *doiurlLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdoiurlLexer(input antlr.CharStream) *doiurlLexer {
	l := new(doiurlLexer)
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
	l.GrammarFileName = "doiurl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// doiurlLexer tokens.
const (
	doiurlLexerT__0  = 1
	doiurlLexerT__1  = 2
	doiurlLexerT__2  = 3
	doiurlLexerT__3  = 4
	doiurlLexerT__4  = 5
	doiurlLexerPCHAR = 6
	doiurlLexerWS    = 7
)
