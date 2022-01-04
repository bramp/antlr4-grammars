// Code generated from metamath.g4 by ANTLR 4.9.3. DO NOT EDIT.

package metamath

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 23, 131,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 18, 6, 18, 94, 10, 18, 13, 18, 14, 18, 95, 3, 19, 3, 19, 6, 19,
	100, 10, 19, 13, 19, 14, 19, 101, 3, 20, 3, 20, 3, 21, 6, 21, 107, 10,
	21, 13, 21, 14, 21, 108, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 115, 10, 22,
	12, 22, 14, 22, 118, 11, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23,
	6, 23, 126, 10, 23, 13, 23, 14, 23, 127, 3, 23, 3, 23, 3, 116, 2, 24, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 2, 41, 21,
	43, 22, 45, 23, 3, 2, 7, 3, 2, 35, 128, 4, 2, 47, 48, 97, 97, 5, 2, 50,
	59, 67, 92, 99, 124, 4, 2, 65, 65, 67, 92, 5, 2, 11, 12, 14, 15, 34, 34,
	2, 135, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 3, 47, 3, 2, 2, 2, 5, 50, 3, 2,
	2, 2, 7, 53, 3, 2, 2, 2, 9, 56, 3, 2, 2, 2, 11, 59, 3, 2, 2, 2, 13, 62,
	3, 2, 2, 2, 15, 65, 3, 2, 2, 2, 17, 68, 3, 2, 2, 2, 19, 71, 3, 2, 2, 2,
	21, 74, 3, 2, 2, 2, 23, 77, 3, 2, 2, 2, 25, 80, 3, 2, 2, 2, 27, 83, 3,
	2, 2, 2, 29, 86, 3, 2, 2, 2, 31, 88, 3, 2, 2, 2, 33, 90, 3, 2, 2, 2, 35,
	93, 3, 2, 2, 2, 37, 99, 3, 2, 2, 2, 39, 103, 3, 2, 2, 2, 41, 106, 3, 2,
	2, 2, 43, 110, 3, 2, 2, 2, 45, 125, 3, 2, 2, 2, 47, 48, 7, 38, 2, 2, 48,
	49, 7, 93, 2, 2, 49, 4, 3, 2, 2, 2, 50, 51, 7, 38, 2, 2, 51, 52, 7, 95,
	2, 2, 52, 6, 3, 2, 2, 2, 53, 54, 7, 38, 2, 2, 54, 55, 7, 101, 2, 2, 55,
	8, 3, 2, 2, 2, 56, 57, 7, 38, 2, 2, 57, 58, 7, 48, 2, 2, 58, 10, 3, 2,
	2, 2, 59, 60, 7, 38, 2, 2, 60, 61, 7, 125, 2, 2, 61, 12, 3, 2, 2, 2, 62,
	63, 7, 38, 2, 2, 63, 64, 7, 127, 2, 2, 64, 14, 3, 2, 2, 2, 65, 66, 7, 38,
	2, 2, 66, 67, 7, 120, 2, 2, 67, 16, 3, 2, 2, 2, 68, 69, 7, 38, 2, 2, 69,
	70, 7, 102, 2, 2, 70, 18, 3, 2, 2, 2, 71, 72, 7, 38, 2, 2, 72, 73, 7, 104,
	2, 2, 73, 20, 3, 2, 2, 2, 74, 75, 7, 38, 2, 2, 75, 76, 7, 103, 2, 2, 76,
	22, 3, 2, 2, 2, 77, 78, 7, 38, 2, 2, 78, 79, 7, 99, 2, 2, 79, 24, 3, 2,
	2, 2, 80, 81, 7, 38, 2, 2, 81, 82, 7, 114, 2, 2, 82, 26, 3, 2, 2, 2, 83,
	84, 7, 38, 2, 2, 84, 85, 7, 63, 2, 2, 85, 28, 3, 2, 2, 2, 86, 87, 7, 65,
	2, 2, 87, 30, 3, 2, 2, 2, 88, 89, 7, 42, 2, 2, 89, 32, 3, 2, 2, 2, 90,
	91, 7, 43, 2, 2, 91, 34, 3, 2, 2, 2, 92, 94, 9, 2, 2, 2, 93, 92, 3, 2,
	2, 2, 94, 95, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 36,
	3, 2, 2, 2, 97, 100, 5, 39, 20, 2, 98, 100, 9, 3, 2, 2, 99, 97, 3, 2, 2,
	2, 99, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101,
	102, 3, 2, 2, 2, 102, 38, 3, 2, 2, 2, 103, 104, 9, 4, 2, 2, 104, 40, 3,
	2, 2, 2, 105, 107, 9, 5, 2, 2, 106, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2,
	2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 42, 3, 2, 2, 2, 110,
	111, 7, 38, 2, 2, 111, 112, 7, 42, 2, 2, 112, 116, 3, 2, 2, 2, 113, 115,
	11, 2, 2, 2, 114, 113, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 117, 3, 2,
	2, 2, 116, 114, 3, 2, 2, 2, 117, 119, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2,
	119, 120, 7, 38, 2, 2, 120, 121, 7, 43, 2, 2, 121, 122, 3, 2, 2, 2, 122,
	123, 8, 22, 2, 2, 123, 44, 3, 2, 2, 2, 124, 126, 9, 6, 2, 2, 125, 124,
	3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2,
	2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 8, 23, 2, 2, 130, 46, 3, 2, 2, 2,
	10, 2, 95, 99, 101, 106, 108, 116, 127, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'$['", "'$]'", "'$c'", "'$.'", "'${'", "'$}'", "'$v'", "'$d'", "'$f'",
	"'$e'", "'$a'", "'$p'", "'$='", "'?'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LPAREN", "RPAREN",
	"PRINTABLECHARACTER", "LABEL", "COMPRESSEDPROOFBLOCK", "BLOCK_COMMENT",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "LPAREN", "RPAREN", "PRINTABLECHARACTER",
	"LABEL", "LETTERORDIGIT", "COMPRESSEDPROOFBLOCK", "BLOCK_COMMENT", "WS",
}

type metamathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewmetamathLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *metamathLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewmetamathLexer(input antlr.CharStream) *metamathLexer {
	l := new(metamathLexer)
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
	l.GrammarFileName = "metamath.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// metamathLexer tokens.
const (
	metamathLexerT__0                 = 1
	metamathLexerT__1                 = 2
	metamathLexerT__2                 = 3
	metamathLexerT__3                 = 4
	metamathLexerT__4                 = 5
	metamathLexerT__5                 = 6
	metamathLexerT__6                 = 7
	metamathLexerT__7                 = 8
	metamathLexerT__8                 = 9
	metamathLexerT__9                 = 10
	metamathLexerT__10                = 11
	metamathLexerT__11                = 12
	metamathLexerT__12                = 13
	metamathLexerT__13                = 14
	metamathLexerLPAREN               = 15
	metamathLexerRPAREN               = 16
	metamathLexerPRINTABLECHARACTER   = 17
	metamathLexerLABEL                = 18
	metamathLexerCOMPRESSEDPROOFBLOCK = 19
	metamathLexerBLOCK_COMMENT        = 20
	metamathLexerWS                   = 21
)
