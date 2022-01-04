// Code generated from http.g4 by ANTLR 4.9.3. DO NOT EDIT.

package http

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 45, 247,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 155, 10, 15,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 5,
	41, 209, 10, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 244, 10, 43, 3, 44, 3, 44, 2, 2, 45,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59,
	31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77,
	40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 3, 2, 5, 4, 2, 67, 92, 99,
	124, 3, 2, 50, 59, 4, 2, 125, 125, 127, 127, 2, 278, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2,
	51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2,
	2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2,
	2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2,
	2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3,
	2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 3, 89,
	3, 2, 2, 2, 5, 93, 3, 2, 2, 2, 7, 98, 3, 2, 2, 2, 9, 103, 3, 2, 2, 2, 11,
	107, 3, 2, 2, 2, 13, 114, 3, 2, 2, 2, 15, 122, 3, 2, 2, 2, 17, 130, 3,
	2, 2, 2, 19, 136, 3, 2, 2, 2, 21, 142, 3, 2, 2, 2, 23, 144, 3, 2, 2, 2,
	25, 146, 3, 2, 2, 2, 27, 148, 3, 2, 2, 2, 29, 154, 3, 2, 2, 2, 31, 156,
	3, 2, 2, 2, 33, 158, 3, 2, 2, 2, 35, 160, 3, 2, 2, 2, 37, 162, 3, 2, 2,
	2, 39, 164, 3, 2, 2, 2, 41, 166, 3, 2, 2, 2, 43, 168, 3, 2, 2, 2, 45, 170,
	3, 2, 2, 2, 47, 172, 3, 2, 2, 2, 49, 174, 3, 2, 2, 2, 51, 176, 3, 2, 2,
	2, 53, 178, 3, 2, 2, 2, 55, 180, 3, 2, 2, 2, 57, 182, 3, 2, 2, 2, 59, 184,
	3, 2, 2, 2, 61, 186, 3, 2, 2, 2, 63, 188, 3, 2, 2, 2, 65, 190, 3, 2, 2,
	2, 67, 192, 3, 2, 2, 2, 69, 194, 3, 2, 2, 2, 71, 196, 3, 2, 2, 2, 73, 198,
	3, 2, 2, 2, 75, 200, 3, 2, 2, 2, 77, 202, 3, 2, 2, 2, 79, 204, 3, 2, 2,
	2, 81, 208, 3, 2, 2, 2, 83, 210, 3, 2, 2, 2, 85, 243, 3, 2, 2, 2, 87, 245,
	3, 2, 2, 2, 89, 90, 7, 73, 2, 2, 90, 91, 7, 71, 2, 2, 91, 92, 7, 86, 2,
	2, 92, 4, 3, 2, 2, 2, 93, 94, 7, 74, 2, 2, 94, 95, 7, 71, 2, 2, 95, 96,
	7, 67, 2, 2, 96, 97, 7, 70, 2, 2, 97, 6, 3, 2, 2, 2, 98, 99, 7, 82, 2,
	2, 99, 100, 7, 81, 2, 2, 100, 101, 7, 85, 2, 2, 101, 102, 7, 86, 2, 2,
	102, 8, 3, 2, 2, 2, 103, 104, 7, 82, 2, 2, 104, 105, 7, 87, 2, 2, 105,
	106, 7, 86, 2, 2, 106, 10, 3, 2, 2, 2, 107, 108, 7, 70, 2, 2, 108, 109,
	7, 71, 2, 2, 109, 110, 7, 78, 2, 2, 110, 111, 7, 71, 2, 2, 111, 112, 7,
	86, 2, 2, 112, 113, 7, 71, 2, 2, 113, 12, 3, 2, 2, 2, 114, 115, 7, 69,
	2, 2, 115, 116, 7, 81, 2, 2, 116, 117, 7, 80, 2, 2, 117, 118, 7, 80, 2,
	2, 118, 119, 7, 71, 2, 2, 119, 120, 7, 69, 2, 2, 120, 121, 7, 86, 2, 2,
	121, 14, 3, 2, 2, 2, 122, 123, 7, 81, 2, 2, 123, 124, 7, 82, 2, 2, 124,
	125, 7, 86, 2, 2, 125, 126, 7, 75, 2, 2, 126, 127, 7, 81, 2, 2, 127, 128,
	7, 80, 2, 2, 128, 129, 7, 85, 2, 2, 129, 16, 3, 2, 2, 2, 130, 131, 7, 86,
	2, 2, 131, 132, 7, 84, 2, 2, 132, 133, 7, 67, 2, 2, 133, 134, 7, 69, 2,
	2, 134, 135, 7, 71, 2, 2, 135, 18, 3, 2, 2, 2, 136, 137, 7, 74, 2, 2, 137,
	138, 7, 86, 2, 2, 138, 139, 7, 86, 2, 2, 139, 140, 7, 82, 2, 2, 140, 141,
	7, 49, 2, 2, 141, 20, 3, 2, 2, 2, 142, 143, 7, 34, 2, 2, 143, 22, 3, 2,
	2, 2, 144, 145, 9, 2, 2, 2, 145, 24, 3, 2, 2, 2, 146, 147, 9, 3, 2, 2,
	147, 26, 3, 2, 2, 2, 148, 149, 5, 67, 34, 2, 149, 150, 5, 29, 15, 2, 150,
	151, 5, 29, 15, 2, 151, 28, 3, 2, 2, 2, 152, 155, 5, 25, 13, 2, 153, 155,
	4, 67, 72, 2, 154, 152, 3, 2, 2, 2, 154, 153, 3, 2, 2, 2, 155, 30, 3, 2,
	2, 2, 156, 157, 7, 42, 2, 2, 157, 32, 3, 2, 2, 2, 158, 159, 7, 43, 2, 2,
	159, 34, 3, 2, 2, 2, 160, 161, 7, 61, 2, 2, 161, 36, 3, 2, 2, 2, 162, 163,
	7, 63, 2, 2, 163, 38, 3, 2, 2, 2, 164, 165, 7, 46, 2, 2, 165, 40, 3, 2,
	2, 2, 166, 167, 7, 12, 2, 2, 167, 42, 3, 2, 2, 2, 168, 169, 7, 47, 2, 2,
	169, 44, 3, 2, 2, 2, 170, 171, 7, 48, 2, 2, 171, 46, 3, 2, 2, 2, 172, 173,
	7, 97, 2, 2, 173, 48, 3, 2, 2, 2, 174, 175, 7, 128, 2, 2, 175, 50, 3, 2,
	2, 2, 176, 177, 7, 65, 2, 2, 177, 52, 3, 2, 2, 2, 178, 179, 7, 49, 2, 2,
	179, 54, 3, 2, 2, 2, 180, 181, 7, 35, 2, 2, 181, 56, 3, 2, 2, 2, 182, 183,
	7, 60, 2, 2, 183, 58, 3, 2, 2, 2, 184, 185, 7, 66, 2, 2, 185, 60, 3, 2,
	2, 2, 186, 187, 7, 38, 2, 2, 187, 62, 3, 2, 2, 2, 188, 189, 7, 37, 2, 2,
	189, 64, 3, 2, 2, 2, 190, 191, 7, 40, 2, 2, 191, 66, 3, 2, 2, 2, 192, 193,
	7, 39, 2, 2, 193, 68, 3, 2, 2, 2, 194, 195, 7, 41, 2, 2, 195, 70, 3, 2,
	2, 2, 196, 197, 7, 44, 2, 2, 197, 72, 3, 2, 2, 2, 198, 199, 7, 45, 2, 2,
	199, 74, 3, 2, 2, 2, 200, 201, 7, 96, 2, 2, 201, 76, 3, 2, 2, 2, 202, 203,
	7, 98, 2, 2, 203, 78, 3, 2, 2, 2, 204, 205, 7, 126, 2, 2, 205, 80, 3, 2,
	2, 2, 206, 209, 5, 21, 11, 2, 207, 209, 5, 83, 42, 2, 208, 206, 3, 2, 2,
	2, 208, 207, 3, 2, 2, 2, 209, 82, 3, 2, 2, 2, 210, 211, 7, 11, 2, 2, 211,
	84, 3, 2, 2, 2, 212, 244, 5, 55, 28, 2, 213, 244, 7, 36, 2, 2, 214, 244,
	5, 63, 32, 2, 215, 244, 5, 61, 31, 2, 216, 244, 5, 67, 34, 2, 217, 244,
	5, 65, 33, 2, 218, 244, 5, 69, 35, 2, 219, 244, 5, 31, 16, 2, 220, 244,
	5, 33, 17, 2, 221, 244, 5, 33, 17, 2, 222, 244, 5, 71, 36, 2, 223, 244,
	5, 73, 37, 2, 224, 244, 5, 39, 20, 2, 225, 244, 5, 43, 22, 2, 226, 244,
	5, 45, 23, 2, 227, 244, 5, 53, 27, 2, 228, 244, 5, 57, 29, 2, 229, 244,
	5, 35, 18, 2, 230, 244, 7, 62, 2, 2, 231, 244, 5, 37, 19, 2, 232, 244,
	7, 64, 2, 2, 233, 244, 5, 51, 26, 2, 234, 244, 5, 59, 30, 2, 235, 244,
	4, 93, 94, 2, 236, 244, 5, 75, 38, 2, 237, 244, 5, 47, 24, 2, 238, 244,
	7, 95, 2, 2, 239, 244, 5, 77, 39, 2, 240, 244, 9, 4, 2, 2, 241, 244, 5,
	79, 40, 2, 242, 244, 5, 49, 25, 2, 243, 212, 3, 2, 2, 2, 243, 213, 3, 2,
	2, 2, 243, 214, 3, 2, 2, 2, 243, 215, 3, 2, 2, 2, 243, 216, 3, 2, 2, 2,
	243, 217, 3, 2, 2, 2, 243, 218, 3, 2, 2, 2, 243, 219, 3, 2, 2, 2, 243,
	220, 3, 2, 2, 2, 243, 221, 3, 2, 2, 2, 243, 222, 3, 2, 2, 2, 243, 223,
	3, 2, 2, 2, 243, 224, 3, 2, 2, 2, 243, 225, 3, 2, 2, 2, 243, 226, 3, 2,
	2, 2, 243, 227, 3, 2, 2, 2, 243, 228, 3, 2, 2, 2, 243, 229, 3, 2, 2, 2,
	243, 230, 3, 2, 2, 2, 243, 231, 3, 2, 2, 2, 243, 232, 3, 2, 2, 2, 243,
	233, 3, 2, 2, 2, 243, 234, 3, 2, 2, 2, 243, 235, 3, 2, 2, 2, 243, 236,
	3, 2, 2, 2, 243, 237, 3, 2, 2, 2, 243, 238, 3, 2, 2, 2, 243, 239, 3, 2,
	2, 2, 243, 240, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 242, 3, 2, 2, 2,
	244, 86, 3, 2, 2, 2, 245, 246, 4, 130, 257, 2, 246, 88, 3, 2, 2, 2, 6,
	2, 154, 208, 243, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'GET'", "'HEAD'", "'POST'", "'PUT'", "'DELETE'", "'CONNECT'", "'OPTIONS'",
	"'TRACE'", "'HTTP/'", "' '", "", "", "", "", "'('", "')'", "';'", "'='",
	"','", "'\n'", "'-'", "'.'", "'_'", "'~'", "'?'", "'/'", "'!'", "':'",
	"'@'", "'$'", "'#'", "'&'", "'%'", "'''", "'*'", "'+'", "'^'", "'`'", "'|'",
	"", "'\t'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "SP", "ALPHA", "DIGIT", "Pct_encoded",
	"HEXDIG", "LColumn", "RColumn", "SemiColon", "Equals", "Period", "CRLF",
	"Minus", "Dot", "Underscore", "Tilde", "QuestionMark", "Slash", "ExclamationMark",
	"Colon", "At", "DollarSign", "Hashtag", "Ampersand", "Percent", "SQuote",
	"Star", "Plus", "Caret", "BackQuote", "VBar", "OWS", "HTAB", "VCHAR", "OBS_TEXT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"SP", "ALPHA", "DIGIT", "Pct_encoded", "HEXDIG", "LColumn", "RColumn",
	"SemiColon", "Equals", "Period", "CRLF", "Minus", "Dot", "Underscore",
	"Tilde", "QuestionMark", "Slash", "ExclamationMark", "Colon", "At", "DollarSign",
	"Hashtag", "Ampersand", "Percent", "SQuote", "Star", "Plus", "Caret", "BackQuote",
	"VBar", "OWS", "HTAB", "VCHAR", "OBS_TEXT",
}

type httpLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewhttpLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *httpLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewhttpLexer(input antlr.CharStream) *httpLexer {
	l := new(httpLexer)
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
	l.GrammarFileName = "http.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// httpLexer tokens.
const (
	httpLexerT__0            = 1
	httpLexerT__1            = 2
	httpLexerT__2            = 3
	httpLexerT__3            = 4
	httpLexerT__4            = 5
	httpLexerT__5            = 6
	httpLexerT__6            = 7
	httpLexerT__7            = 8
	httpLexerT__8            = 9
	httpLexerSP              = 10
	httpLexerALPHA           = 11
	httpLexerDIGIT           = 12
	httpLexerPct_encoded     = 13
	httpLexerHEXDIG          = 14
	httpLexerLColumn         = 15
	httpLexerRColumn         = 16
	httpLexerSemiColon       = 17
	httpLexerEquals          = 18
	httpLexerPeriod          = 19
	httpLexerCRLF            = 20
	httpLexerMinus           = 21
	httpLexerDot             = 22
	httpLexerUnderscore      = 23
	httpLexerTilde           = 24
	httpLexerQuestionMark    = 25
	httpLexerSlash           = 26
	httpLexerExclamationMark = 27
	httpLexerColon           = 28
	httpLexerAt              = 29
	httpLexerDollarSign      = 30
	httpLexerHashtag         = 31
	httpLexerAmpersand       = 32
	httpLexerPercent         = 33
	httpLexerSQuote          = 34
	httpLexerStar            = 35
	httpLexerPlus            = 36
	httpLexerCaret           = 37
	httpLexerBackQuote       = 38
	httpLexerVBar            = 39
	httpLexerOWS             = 40
	httpLexerHTAB            = 41
	httpLexerVCHAR           = 42
	httpLexerOBS_TEXT        = 43
)
