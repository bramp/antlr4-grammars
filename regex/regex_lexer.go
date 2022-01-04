// Code generated from regexLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package regex

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 39, 232,
	8, 1, 8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5,
	4, 6, 9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9,
	11, 4, 12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16,
	4, 17, 9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4,
	22, 9, 22, 4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27,
	9, 27, 4, 28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9,
	32, 4, 33, 9, 33, 4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37,
	4, 38, 9, 38, 4, 39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 18, 6, 18, 135, 10, 18, 13, 18, 14, 18, 136, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5,
	21, 152, 10, 21, 3, 22, 3, 22, 5, 22, 156, 10, 22, 3, 23, 3, 23, 5, 23,
	160, 10, 23, 3, 24, 3, 24, 5, 24, 164, 10, 24, 3, 25, 3, 25, 5, 25, 168,
	10, 25, 3, 26, 3, 26, 5, 26, 172, 10, 26, 3, 27, 3, 27, 5, 27, 176, 10,
	27, 3, 28, 3, 28, 5, 28, 180, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 6, 29,
	186, 10, 29, 13, 29, 14, 29, 187, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3,
	37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 2, 2, 43, 6, 3, 8, 4, 10, 5,
	12, 6, 14, 7, 16, 8, 18, 9, 20, 10, 22, 11, 24, 12, 26, 13, 28, 14, 30,
	15, 32, 16, 34, 17, 36, 18, 38, 19, 40, 20, 42, 21, 44, 22, 46, 23, 48,
	24, 50, 25, 52, 26, 54, 27, 56, 28, 58, 29, 60, 30, 62, 31, 64, 32, 66,
	33, 68, 34, 70, 35, 72, 36, 74, 37, 76, 38, 78, 39, 80, 2, 82, 2, 84, 2,
	86, 2, 6, 2, 3, 4, 5, 15, 7, 2, 42, 45, 48, 48, 65, 65, 93, 95, 126, 126,
	3, 2, 50, 59, 5, 2, 110, 111, 113, 113, 118, 119, 5, 2, 101, 101, 103,
	103, 112, 112, 5, 2, 102, 102, 110, 110, 113, 113, 6, 2, 101, 104, 107,
	107, 113, 113, 117, 117, 5, 2, 110, 110, 114, 114, 117, 117, 6, 2, 101,
	101, 109, 109, 111, 111, 113, 113, 5, 2, 101, 101, 104, 104, 112, 113,
	6, 2, 47, 47, 50, 59, 67, 92, 99, 124, 5, 2, 47, 47, 93, 93, 95, 95, 10,
	2, 69, 70, 75, 75, 85, 85, 89, 89, 101, 102, 107, 107, 117, 117, 121, 121,
	10, 2, 42, 45, 47, 48, 65, 65, 93, 96, 112, 112, 116, 116, 118, 118, 125,
	127, 2, 239, 2, 6, 3, 2, 2, 2, 2, 8, 3, 2, 2, 2, 2, 10, 3, 2, 2, 2, 2,
	12, 3, 2, 2, 2, 2, 14, 3, 2, 2, 2, 2, 16, 3, 2, 2, 2, 2, 18, 3, 2, 2, 2,
	2, 20, 3, 2, 2, 2, 2, 22, 3, 2, 2, 2, 2, 24, 3, 2, 2, 2, 2, 26, 3, 2, 2,
	2, 2, 28, 3, 2, 2, 2, 2, 30, 3, 2, 2, 2, 2, 32, 3, 2, 2, 2, 2, 34, 3, 2,
	2, 2, 3, 36, 3, 2, 2, 2, 3, 38, 3, 2, 2, 2, 3, 40, 3, 2, 2, 2, 4, 42, 3,
	2, 2, 2, 4, 44, 3, 2, 2, 2, 4, 46, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 4, 50,
	3, 2, 2, 2, 4, 52, 3, 2, 2, 2, 4, 54, 3, 2, 2, 2, 4, 56, 3, 2, 2, 2, 4,
	58, 3, 2, 2, 2, 4, 60, 3, 2, 2, 2, 5, 62, 3, 2, 2, 2, 5, 64, 3, 2, 2, 2,
	5, 66, 3, 2, 2, 2, 5, 68, 3, 2, 2, 2, 5, 70, 3, 2, 2, 2, 5, 72, 3, 2, 2,
	2, 5, 74, 3, 2, 2, 2, 5, 76, 3, 2, 2, 2, 5, 78, 3, 2, 2, 2, 6, 88, 3, 2,
	2, 2, 8, 90, 3, 2, 2, 2, 10, 92, 3, 2, 2, 2, 12, 94, 3, 2, 2, 2, 14, 96,
	3, 2, 2, 2, 16, 98, 3, 2, 2, 2, 18, 100, 3, 2, 2, 2, 20, 102, 3, 2, 2,
	2, 22, 104, 3, 2, 2, 2, 24, 108, 3, 2, 2, 2, 26, 110, 3, 2, 2, 2, 28, 112,
	3, 2, 2, 2, 30, 116, 3, 2, 2, 2, 32, 120, 3, 2, 2, 2, 34, 125, 3, 2, 2,
	2, 36, 129, 3, 2, 2, 2, 38, 134, 3, 2, 2, 2, 40, 138, 3, 2, 2, 2, 42, 140,
	3, 2, 2, 2, 44, 151, 3, 2, 2, 2, 46, 153, 3, 2, 2, 2, 48, 157, 3, 2, 2,
	2, 50, 161, 3, 2, 2, 2, 52, 165, 3, 2, 2, 2, 54, 169, 3, 2, 2, 2, 56, 173,
	3, 2, 2, 2, 58, 177, 3, 2, 2, 2, 60, 181, 3, 2, 2, 2, 62, 189, 3, 2, 2,
	2, 64, 191, 3, 2, 2, 2, 66, 193, 3, 2, 2, 2, 68, 197, 3, 2, 2, 2, 70, 201,
	3, 2, 2, 2, 72, 206, 3, 2, 2, 2, 74, 210, 3, 2, 2, 2, 76, 214, 3, 2, 2,
	2, 78, 216, 3, 2, 2, 2, 80, 218, 3, 2, 2, 2, 82, 222, 3, 2, 2, 2, 84, 226,
	3, 2, 2, 2, 86, 229, 3, 2, 2, 2, 88, 89, 7, 42, 2, 2, 89, 7, 3, 2, 2, 2,
	90, 91, 7, 43, 2, 2, 91, 9, 3, 2, 2, 2, 92, 93, 7, 126, 2, 2, 93, 11, 3,
	2, 2, 2, 94, 95, 7, 45, 2, 2, 95, 13, 3, 2, 2, 2, 96, 97, 7, 65, 2, 2,
	97, 15, 3, 2, 2, 2, 98, 99, 7, 44, 2, 2, 99, 17, 3, 2, 2, 2, 100, 101,
	7, 48, 2, 2, 101, 19, 3, 2, 2, 2, 102, 103, 10, 2, 2, 2, 103, 21, 3, 2,
	2, 2, 104, 105, 7, 125, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 8, 10, 2,
	2, 107, 23, 3, 2, 2, 2, 108, 109, 5, 86, 42, 2, 109, 25, 3, 2, 2, 2, 110,
	111, 5, 84, 41, 2, 111, 27, 3, 2, 2, 2, 112, 113, 5, 80, 39, 2, 113, 114,
	3, 2, 2, 2, 114, 115, 8, 13, 3, 2, 115, 29, 3, 2, 2, 2, 116, 117, 5, 82,
	40, 2, 117, 118, 3, 2, 2, 2, 118, 119, 8, 14, 3, 2, 119, 31, 3, 2, 2, 2,
	120, 121, 7, 93, 2, 2, 121, 122, 7, 96, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	124, 8, 15, 4, 2, 124, 33, 3, 2, 2, 2, 125, 126, 7, 93, 2, 2, 126, 127,
	3, 2, 2, 2, 127, 128, 8, 16, 4, 2, 128, 35, 3, 2, 2, 2, 129, 130, 7, 127,
	2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 8, 17, 5, 2, 132, 37, 3, 2, 2, 2,
	133, 135, 9, 3, 2, 2, 134, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 39, 3, 2, 2, 2, 138, 139, 7,
	46, 2, 2, 139, 41, 3, 2, 2, 2, 140, 141, 7, 127, 2, 2, 141, 142, 3, 2,
	2, 2, 142, 143, 8, 20, 5, 2, 143, 43, 3, 2, 2, 2, 144, 152, 5, 46, 22,
	2, 145, 152, 5, 48, 23, 2, 146, 152, 5, 50, 24, 2, 147, 152, 5, 52, 25,
	2, 148, 152, 5, 54, 26, 2, 149, 152, 5, 56, 27, 2, 150, 152, 5, 58, 28,
	2, 151, 144, 3, 2, 2, 2, 151, 145, 3, 2, 2, 2, 151, 146, 3, 2, 2, 2, 151,
	147, 3, 2, 2, 2, 151, 148, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 150,
	3, 2, 2, 2, 152, 45, 3, 2, 2, 2, 153, 155, 7, 78, 2, 2, 154, 156, 9, 4,
	2, 2, 155, 154, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 47, 3, 2, 2, 2,
	157, 159, 7, 79, 2, 2, 158, 160, 9, 5, 2, 2, 159, 158, 3, 2, 2, 2, 159,
	160, 3, 2, 2, 2, 160, 49, 3, 2, 2, 2, 161, 163, 7, 80, 2, 2, 162, 164,
	9, 6, 2, 2, 163, 162, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 51, 3, 2,
	2, 2, 165, 167, 7, 82, 2, 2, 166, 168, 9, 7, 2, 2, 167, 166, 3, 2, 2, 2,
	167, 168, 3, 2, 2, 2, 168, 53, 3, 2, 2, 2, 169, 171, 7, 92, 2, 2, 170,
	172, 9, 8, 2, 2, 171, 170, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 55, 3,
	2, 2, 2, 173, 175, 7, 85, 2, 2, 174, 176, 9, 9, 2, 2, 175, 174, 3, 2, 2,
	2, 175, 176, 3, 2, 2, 2, 176, 57, 3, 2, 2, 2, 177, 179, 7, 69, 2, 2, 178,
	180, 9, 10, 2, 2, 179, 178, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 59,
	3, 2, 2, 2, 181, 182, 7, 75, 2, 2, 182, 183, 7, 117, 2, 2, 183, 185, 3,
	2, 2, 2, 184, 186, 9, 11, 2, 2, 185, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2,
	2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 61, 3, 2, 2, 2, 189,
	190, 5, 86, 42, 2, 190, 63, 3, 2, 2, 2, 191, 192, 5, 84, 41, 2, 192, 65,
	3, 2, 2, 2, 193, 194, 5, 80, 39, 2, 194, 195, 3, 2, 2, 2, 195, 196, 8,
	32, 3, 2, 196, 67, 3, 2, 2, 2, 197, 198, 5, 82, 40, 2, 198, 199, 3, 2,
	2, 2, 199, 200, 8, 33, 3, 2, 200, 69, 3, 2, 2, 2, 201, 202, 7, 93, 2, 2,
	202, 203, 7, 96, 2, 2, 203, 204, 3, 2, 2, 2, 204, 205, 8, 34, 4, 2, 205,
	71, 3, 2, 2, 2, 206, 207, 7, 93, 2, 2, 207, 208, 3, 2, 2, 2, 208, 209,
	8, 35, 4, 2, 209, 73, 3, 2, 2, 2, 210, 211, 7, 95, 2, 2, 211, 212, 3, 2,
	2, 2, 212, 213, 8, 36, 5, 2, 213, 75, 3, 2, 2, 2, 214, 215, 7, 47, 2, 2,
	215, 77, 3, 2, 2, 2, 216, 217, 10, 12, 2, 2, 217, 79, 3, 2, 2, 2, 218,
	219, 7, 94, 2, 2, 219, 220, 7, 114, 2, 2, 220, 221, 7, 125, 2, 2, 221,
	81, 3, 2, 2, 2, 222, 223, 7, 94, 2, 2, 223, 224, 7, 82, 2, 2, 224, 225,
	7, 125, 2, 2, 225, 83, 3, 2, 2, 2, 226, 227, 7, 94, 2, 2, 227, 228, 9,
	13, 2, 2, 228, 85, 3, 2, 2, 2, 229, 230, 7, 94, 2, 2, 230, 231, 9, 14,
	2, 2, 231, 87, 3, 2, 2, 2, 17, 2, 3, 4, 5, 136, 151, 155, 159, 163, 167,
	171, 175, 179, 185, 187, 6, 7, 3, 2, 7, 4, 2, 7, 5, 2, 6, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "QUANTITY", "CATEGORY", "CHARGROUP",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'|'", "'+'", "'?'", "'*'", "'.'", "", "'{'", "", "",
	"", "", "", "", "", "", "','", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "']'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "LPAREN", "RPAREN", "PIPE", "PLUS", "QUESTION", "STAR", "WildcardEsc",
	"Char", "StartQuantity", "SingleCharEsc", "MultiCharEsc", "CatEsc", "ComplEsc",
	"NegCharGroup", "PosCharGroup", "EndQuantity", "QuantExact", "COMMA", "EndCategory",
	"IsCategory", "Letters", "Marks", "Numbers", "Punctuation", "Separators",
	"Symbols", "Others", "IsBlock", "NestedSingleCharEsc", "NestedMultiCharEsc",
	"NestedCatEsc", "NestedComplEsc", "NestedNegCharGroup", "NestedPosCharGroup",
	"EndCharGroup", "DASH", "XmlChar",
}

var lexerRuleNames = []string{
	"LPAREN", "RPAREN", "PIPE", "PLUS", "QUESTION", "STAR", "WildcardEsc",
	"Char", "StartQuantity", "SingleCharEsc", "MultiCharEsc", "CatEsc", "ComplEsc",
	"NegCharGroup", "PosCharGroup", "EndQuantity", "QuantExact", "COMMA", "EndCategory",
	"IsCategory", "Letters", "Marks", "Numbers", "Punctuation", "Separators",
	"Symbols", "Others", "IsBlock", "NestedSingleCharEsc", "NestedMultiCharEsc",
	"NestedCatEsc", "NestedComplEsc", "NestedNegCharGroup", "NestedPosCharGroup",
	"EndCharGroup", "DASH", "XmlChar", "CAT_ESC", "COMPL_ESC", "MULTI_ESC",
	"SINGLE_ESC",
}

type regexLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewregexLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *regexLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewregexLexer(input antlr.CharStream) *regexLexer {
	l := new(regexLexer)
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
	l.GrammarFileName = "regexLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// regexLexer tokens.
const (
	regexLexerLPAREN              = 1
	regexLexerRPAREN              = 2
	regexLexerPIPE                = 3
	regexLexerPLUS                = 4
	regexLexerQUESTION            = 5
	regexLexerSTAR                = 6
	regexLexerWildcardEsc         = 7
	regexLexerChar                = 8
	regexLexerStartQuantity       = 9
	regexLexerSingleCharEsc       = 10
	regexLexerMultiCharEsc        = 11
	regexLexerCatEsc              = 12
	regexLexerComplEsc            = 13
	regexLexerNegCharGroup        = 14
	regexLexerPosCharGroup        = 15
	regexLexerEndQuantity         = 16
	regexLexerQuantExact          = 17
	regexLexerCOMMA               = 18
	regexLexerEndCategory         = 19
	regexLexerIsCategory          = 20
	regexLexerLetters             = 21
	regexLexerMarks               = 22
	regexLexerNumbers             = 23
	regexLexerPunctuation         = 24
	regexLexerSeparators          = 25
	regexLexerSymbols             = 26
	regexLexerOthers              = 27
	regexLexerIsBlock             = 28
	regexLexerNestedSingleCharEsc = 29
	regexLexerNestedMultiCharEsc  = 30
	regexLexerNestedCatEsc        = 31
	regexLexerNestedComplEsc      = 32
	regexLexerNestedNegCharGroup  = 33
	regexLexerNestedPosCharGroup  = 34
	regexLexerEndCharGroup        = 35
	regexLexerDASH                = 36
	regexLexerXmlChar             = 37
)

// regexLexer modes.
const (
	regexLexerQUANTITY = iota + 1
	regexLexerCATEGORY
	regexLexerCHARGROUP
)
