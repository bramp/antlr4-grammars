// Code generated from tcp.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tcp

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 4, 16, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 3, 6, 3, 11, 10, 3, 13, 3, 14,
	3, 12, 3, 3, 3, 3, 2, 2, 4, 3, 3, 5, 4, 3, 2, 3, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 16, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 3, 7, 3, 2, 2, 2, 5,
	10, 3, 2, 2, 2, 7, 8, 4, 2, 257, 2, 8, 4, 3, 2, 2, 2, 9, 11, 9, 2, 2, 2,
	10, 9, 3, 2, 2, 2, 11, 12, 3, 2, 2, 2, 12, 10, 3, 2, 2, 2, 12, 13, 3, 2,
	2, 2, 13, 14, 3, 2, 2, 2, 14, 15, 8, 3, 2, 2, 15, 6, 3, 2, 2, 2, 4, 2,
	12, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames []string

var lexerSymbolicNames = []string{
	"", "BYTE", "WS",
}

var lexerRuleNames = []string{
	"BYTE", "WS",
}

type tcpLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewtcpLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *tcpLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewtcpLexer(input antlr.CharStream) *tcpLexer {
	l := new(tcpLexer)
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
	l.GrammarFileName = "tcp.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// tcpLexer tokens.
const (
	tcpLexerBYTE = 1
	tcpLexerWS   = 2
)
