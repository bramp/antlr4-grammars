// Code generated from mumath.g4 by ANTLR 4.9.3. DO NOT EDIT.

package mumath

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 40, 237,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 7, 16, 159, 10, 16, 12, 16, 14, 16, 162, 11, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3,
	35, 3, 35, 3, 36, 3, 36, 7, 36, 212, 10, 36, 12, 36, 14, 36, 215, 11, 36,
	3, 36, 5, 36, 218, 10, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 7,
	38, 226, 10, 38, 12, 38, 14, 38, 229, 11, 38, 3, 38, 3, 38, 3, 39, 6, 39,
	234, 10, 39, 13, 39, 14, 39, 235, 2, 2, 40, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16,
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25,
	49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34,
	67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 3, 2, 7, 5, 2, 11, 12,
	15, 15, 34, 34, 4, 2, 12, 12, 39, 39, 5, 2, 37, 37, 66, 92, 125, 125, 6,
	2, 37, 37, 50, 59, 67, 92, 127, 127, 3, 2, 36, 36, 2, 242, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2,
	2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2,
	2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3,
	2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73,
	3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 3, 79, 3, 2, 2, 2, 5,
	85, 3, 2, 2, 2, 7, 94, 3, 2, 2, 2, 9, 103, 3, 2, 2, 2, 11, 110, 3, 2, 2,
	2, 13, 113, 3, 2, 2, 2, 15, 118, 3, 2, 2, 2, 17, 126, 3, 2, 2, 2, 19, 131,
	3, 2, 2, 2, 21, 136, 3, 2, 2, 2, 23, 139, 3, 2, 2, 2, 25, 143, 3, 2, 2,
	2, 27, 147, 3, 2, 2, 2, 29, 151, 3, 2, 2, 2, 31, 155, 3, 2, 2, 2, 33, 167,
	3, 2, 2, 2, 35, 170, 3, 2, 2, 2, 37, 172, 3, 2, 2, 2, 39, 174, 3, 2, 2,
	2, 41, 176, 3, 2, 2, 2, 43, 178, 3, 2, 2, 2, 45, 180, 3, 2, 2, 2, 47, 182,
	3, 2, 2, 2, 49, 184, 3, 2, 2, 2, 51, 186, 3, 2, 2, 2, 53, 188, 3, 2, 2,
	2, 55, 190, 3, 2, 2, 2, 57, 193, 3, 2, 2, 2, 59, 195, 3, 2, 2, 2, 61, 198,
	3, 2, 2, 2, 63, 201, 3, 2, 2, 2, 65, 203, 3, 2, 2, 2, 67, 205, 3, 2, 2,
	2, 69, 207, 3, 2, 2, 2, 71, 209, 3, 2, 2, 2, 73, 219, 3, 2, 2, 2, 75, 223,
	3, 2, 2, 2, 77, 233, 3, 2, 2, 2, 79, 80, 7, 68, 2, 2, 80, 81, 7, 78, 2,
	2, 81, 82, 7, 81, 2, 2, 82, 83, 7, 69, 2, 2, 83, 84, 7, 77, 2, 2, 84, 4,
	3, 2, 2, 2, 85, 86, 7, 71, 2, 2, 86, 87, 7, 80, 2, 2, 87, 88, 7, 70, 2,
	2, 88, 89, 7, 68, 2, 2, 89, 90, 7, 78, 2, 2, 90, 91, 7, 81, 2, 2, 91, 92,
	7, 69, 2, 2, 92, 93, 7, 77, 2, 2, 93, 6, 3, 2, 2, 2, 94, 95, 7, 72, 2,
	2, 95, 96, 7, 87, 2, 2, 96, 97, 7, 80, 2, 2, 97, 98, 7, 69, 2, 2, 98, 99,
	7, 86, 2, 2, 99, 100, 7, 75, 2, 2, 100, 101, 7, 81, 2, 2, 101, 102, 7,
	80, 2, 2, 102, 8, 3, 2, 2, 2, 103, 104, 7, 71, 2, 2, 104, 105, 7, 80, 2,
	2, 105, 106, 7, 70, 2, 2, 106, 107, 7, 72, 2, 2, 107, 108, 7, 87, 2, 2,
	108, 109, 7, 80, 2, 2, 109, 10, 3, 2, 2, 2, 110, 111, 7, 71, 2, 2, 111,
	112, 7, 83, 2, 2, 112, 12, 3, 2, 2, 2, 113, 114, 7, 78, 2, 2, 114, 115,
	7, 81, 2, 2, 115, 116, 7, 81, 2, 2, 116, 117, 7, 82, 2, 2, 117, 14, 3,
	2, 2, 2, 118, 119, 7, 71, 2, 2, 119, 120, 7, 80, 2, 2, 120, 121, 7, 70,
	2, 2, 121, 122, 7, 78, 2, 2, 122, 123, 7, 81, 2, 2, 123, 124, 7, 81, 2,
	2, 124, 125, 7, 82, 2, 2, 125, 16, 3, 2, 2, 2, 126, 127, 7, 89, 2, 2, 127,
	128, 7, 74, 2, 2, 128, 129, 7, 71, 2, 2, 129, 130, 7, 80, 2, 2, 130, 18,
	3, 2, 2, 2, 131, 132, 7, 71, 2, 2, 132, 133, 7, 90, 2, 2, 133, 134, 7,
	75, 2, 2, 134, 135, 7, 86, 2, 2, 135, 20, 3, 2, 2, 2, 136, 137, 7, 81,
	2, 2, 137, 138, 7, 84, 2, 2, 138, 22, 3, 2, 2, 2, 139, 140, 7, 67, 2, 2,
	140, 141, 7, 80, 2, 2, 141, 142, 7, 70, 2, 2, 142, 24, 3, 2, 2, 2, 143,
	144, 7, 80, 2, 2, 144, 145, 7, 81, 2, 2, 145, 146, 7, 86, 2, 2, 146, 26,
	3, 2, 2, 2, 147, 148, 7, 111, 2, 2, 148, 149, 7, 113, 2, 2, 149, 150, 7,
	102, 2, 2, 150, 28, 3, 2, 2, 2, 151, 152, 9, 2, 2, 2, 152, 153, 3, 2, 2,
	2, 153, 154, 8, 15, 2, 2, 154, 30, 3, 2, 2, 2, 155, 160, 7, 39, 2, 2, 156,
	159, 7, 12, 2, 2, 157, 159, 10, 3, 2, 2, 158, 156, 3, 2, 2, 2, 158, 157,
	3, 2, 2, 2, 159, 162, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2,
	2, 2, 161, 163, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 163, 164, 7, 39, 2, 2,
	164, 165, 3, 2, 2, 2, 165, 166, 8, 16, 2, 2, 166, 32, 3, 2, 2, 2, 167,
	168, 7, 63, 2, 2, 168, 169, 7, 63, 2, 2, 169, 34, 3, 2, 2, 2, 170, 171,
	7, 41, 2, 2, 171, 36, 3, 2, 2, 2, 172, 173, 7, 45, 2, 2, 173, 38, 3, 2,
	2, 2, 174, 175, 7, 47, 2, 2, 175, 40, 3, 2, 2, 2, 176, 177, 7, 44, 2, 2,
	177, 42, 3, 2, 2, 2, 178, 179, 7, 49, 2, 2, 179, 44, 3, 2, 2, 2, 180, 181,
	7, 46, 2, 2, 181, 46, 3, 2, 2, 2, 182, 183, 7, 61, 2, 2, 183, 48, 3, 2,
	2, 2, 184, 185, 7, 38, 2, 2, 185, 50, 3, 2, 2, 2, 186, 187, 7, 60, 2, 2,
	187, 52, 3, 2, 2, 2, 188, 189, 7, 63, 2, 2, 189, 54, 3, 2, 2, 2, 190, 191,
	7, 62, 2, 2, 191, 192, 7, 64, 2, 2, 192, 56, 3, 2, 2, 2, 193, 194, 7, 62,
	2, 2, 194, 58, 3, 2, 2, 2, 195, 196, 7, 62, 2, 2, 196, 197, 7, 63, 2, 2,
	197, 60, 3, 2, 2, 2, 198, 199, 7, 64, 2, 2, 199, 200, 7, 63, 2, 2, 200,
	62, 3, 2, 2, 2, 201, 202, 7, 64, 2, 2, 202, 64, 3, 2, 2, 2, 203, 204, 7,
	42, 2, 2, 204, 66, 3, 2, 2, 2, 205, 206, 7, 43, 2, 2, 206, 68, 3, 2, 2,
	2, 207, 208, 7, 96, 2, 2, 208, 70, 3, 2, 2, 2, 209, 213, 9, 4, 2, 2, 210,
	212, 9, 5, 2, 2, 211, 210, 3, 2, 2, 2, 212, 215, 3, 2, 2, 2, 213, 211,
	3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 217, 3, 2, 2, 2, 215, 213, 3, 2,
	2, 2, 216, 218, 5, 73, 37, 2, 217, 216, 3, 2, 2, 2, 217, 218, 3, 2, 2,
	2, 218, 72, 3, 2, 2, 2, 219, 220, 7, 93, 2, 2, 220, 221, 5, 77, 39, 2,
	221, 222, 7, 95, 2, 2, 222, 74, 3, 2, 2, 2, 223, 227, 7, 36, 2, 2, 224,
	226, 10, 6, 2, 2, 225, 224, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 225,
	3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 230, 3, 2, 2, 2, 229, 227, 3, 2,
	2, 2, 230, 231, 7, 36, 2, 2, 231, 76, 3, 2, 2, 2, 232, 234, 4, 50, 59,
	2, 233, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 235,
	236, 3, 2, 2, 2, 236, 78, 3, 2, 2, 2, 9, 2, 158, 160, 213, 217, 227, 235,
	3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'BLOCK'", "'ENDBLOCK'", "'FUNCTION'", "'ENDFUN'", "'EQ'", "'LOOP'",
	"'ENDLOOP'", "'WHEN'", "'EXIT'", "'OR'", "'AND'", "'NOT'", "'mod'", "",
	"", "'=='", "'''", "'+'", "'-'", "'*'", "'/'", "','", "';'", "'$'", "':'",
	"'='", "'<>'", "'<'", "'<='", "'>='", "'>'", "'('", "')'", "'^'",
}

var lexerSymbolicNames = []string{
	"", "BLOCK", "ENDBLOCK", "FUNCTION", "ENDFUN", "EQF", "LOOP", "ENDLOOP",
	"WHEN", "EXIT", "OR", "AND", "NOT", "MOD", "WS", "COMMENT", "EQUATION",
	"QUOTE", "PLUS", "MINUS", "STAR", "SLASH", "COMMA", "SEMI", "DOLLAR", "COLON",
	"EQC", "NOT_EQUAL", "LT", "LE", "GE", "GT", "LPAREN", "RPAREN", "POWER",
	"ID", "ARR", "STRING", "NUMBER",
}

var lexerRuleNames = []string{
	"BLOCK", "ENDBLOCK", "FUNCTION", "ENDFUN", "EQF", "LOOP", "ENDLOOP", "WHEN",
	"EXIT", "OR", "AND", "NOT", "MOD", "WS", "COMMENT", "EQUATION", "QUOTE",
	"PLUS", "MINUS", "STAR", "SLASH", "COMMA", "SEMI", "DOLLAR", "COLON", "EQC",
	"NOT_EQUAL", "LT", "LE", "GE", "GT", "LPAREN", "RPAREN", "POWER", "ID",
	"ARR", "STRING", "NUMBER",
}

type mumathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewmumathLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *mumathLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewmumathLexer(input antlr.CharStream) *mumathLexer {
	l := new(mumathLexer)
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
	l.GrammarFileName = "mumath.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// mumathLexer tokens.
const (
	mumathLexerBLOCK     = 1
	mumathLexerENDBLOCK  = 2
	mumathLexerFUNCTION  = 3
	mumathLexerENDFUN    = 4
	mumathLexerEQF       = 5
	mumathLexerLOOP      = 6
	mumathLexerENDLOOP   = 7
	mumathLexerWHEN      = 8
	mumathLexerEXIT      = 9
	mumathLexerOR        = 10
	mumathLexerAND       = 11
	mumathLexerNOT       = 12
	mumathLexerMOD       = 13
	mumathLexerWS        = 14
	mumathLexerCOMMENT   = 15
	mumathLexerEQUATION  = 16
	mumathLexerQUOTE     = 17
	mumathLexerPLUS      = 18
	mumathLexerMINUS     = 19
	mumathLexerSTAR      = 20
	mumathLexerSLASH     = 21
	mumathLexerCOMMA     = 22
	mumathLexerSEMI      = 23
	mumathLexerDOLLAR    = 24
	mumathLexerCOLON     = 25
	mumathLexerEQC       = 26
	mumathLexerNOT_EQUAL = 27
	mumathLexerLT        = 28
	mumathLexerLE        = 29
	mumathLexerGE        = 30
	mumathLexerGT        = 31
	mumathLexerLPAREN    = 32
	mumathLexerRPAREN    = 33
	mumathLexerPOWER     = 34
	mumathLexerID        = 35
	mumathLexerARR       = 36
	mumathLexerSTRING    = 37
	mumathLexerNUMBER    = 38
)
