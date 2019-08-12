// Code generated from ReStructuredText.g4 by ANTLR 4.7.2. DO NOT EDIT.

package restructuredtext

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 32, 180,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 74, 10, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 80, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 6, 3, 86, 10, 3, 13, 3, 14,
	3, 87, 3, 4, 3, 4, 3, 4, 7, 4, 93, 10, 4, 12, 4, 14, 4, 96, 11, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 106, 10, 5, 3, 6, 6, 6,
	109, 10, 6, 13, 6, 14, 6, 110, 3, 7, 6, 7, 114, 10, 7, 13, 7, 14, 7, 115,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 158, 10,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 164, 10, 25, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 5, 30, 175, 10, 30, 3, 30, 3,
	30, 3, 31, 3, 31, 2, 2, 32, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18,
	35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27,
	53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 3, 2, 5, 4, 2, 67, 92, 99, 124,
	3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 2, 196, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 3, 63, 3, 2, 2, 2, 5, 73, 3, 2, 2, 2, 7,
	89, 3, 2, 2, 2, 9, 105, 3, 2, 2, 2, 11, 108, 3, 2, 2, 2, 13, 113, 3, 2,
	2, 2, 15, 117, 3, 2, 2, 2, 17, 120, 3, 2, 2, 2, 19, 122, 3, 2, 2, 2, 21,
	124, 3, 2, 2, 2, 23, 126, 3, 2, 2, 2, 25, 128, 3, 2, 2, 2, 27, 130, 3,
	2, 2, 2, 29, 132, 3, 2, 2, 2, 31, 134, 3, 2, 2, 2, 33, 136, 3, 2, 2, 2,
	35, 138, 3, 2, 2, 2, 37, 140, 3, 2, 2, 2, 39, 142, 3, 2, 2, 2, 41, 144,
	3, 2, 2, 2, 43, 146, 3, 2, 2, 2, 45, 148, 3, 2, 2, 2, 47, 150, 3, 2, 2,
	2, 49, 163, 3, 2, 2, 2, 51, 165, 3, 2, 2, 2, 53, 167, 3, 2, 2, 2, 55, 169,
	3, 2, 2, 2, 57, 171, 3, 2, 2, 2, 59, 174, 3, 2, 2, 2, 61, 178, 3, 2, 2,
	2, 63, 64, 7, 60, 2, 2, 64, 65, 7, 102, 2, 2, 65, 66, 7, 113, 2, 2, 66,
	67, 7, 101, 2, 2, 67, 68, 7, 60, 2, 2, 68, 4, 3, 2, 2, 2, 69, 74, 5, 45,
	23, 2, 70, 74, 5, 41, 21, 2, 71, 74, 5, 43, 22, 2, 72, 74, 5, 29, 15, 2,
	73, 69, 3, 2, 2, 2, 73, 70, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 72, 3,
	2, 2, 2, 74, 79, 3, 2, 2, 2, 75, 80, 5, 45, 23, 2, 76, 80, 5, 41, 21, 2,
	77, 80, 5, 43, 22, 2, 78, 80, 5, 29, 15, 2, 79, 75, 3, 2, 2, 2, 79, 76,
	3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 78, 3, 2, 2, 2, 80, 85, 3, 2, 2, 2,
	81, 86, 5, 45, 23, 2, 82, 86, 5, 41, 21, 2, 83, 86, 5, 43, 22, 2, 84, 86,
	5, 29, 15, 2, 85, 81, 3, 2, 2, 2, 85, 82, 3, 2, 2, 2, 85, 83, 3, 2, 2,
	2, 85, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88,
	3, 2, 2, 2, 88, 6, 3, 2, 2, 2, 89, 90, 5, 39, 20, 2, 90, 94, 5, 59, 30,
	2, 91, 93, 5, 59, 30, 2, 92, 91, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92,
	3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 97, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2,
	97, 98, 5, 39, 20, 2, 98, 99, 5, 39, 20, 2, 99, 8, 3, 2, 2, 2, 100, 101,
	5, 13, 7, 2, 101, 102, 5, 55, 28, 2, 102, 106, 3, 2, 2, 2, 103, 104, 7,
	122, 2, 2, 104, 106, 5, 55, 28, 2, 105, 100, 3, 2, 2, 2, 105, 103, 3, 2,
	2, 2, 106, 10, 3, 2, 2, 2, 107, 109, 9, 2, 2, 2, 108, 107, 3, 2, 2, 2,
	109, 110, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111,
	12, 3, 2, 2, 2, 112, 114, 9, 3, 2, 2, 113, 112, 3, 2, 2, 2, 114, 115, 3,
	2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 14, 3, 2, 2,
	2, 117, 118, 5, 39, 20, 2, 118, 119, 5, 39, 20, 2, 119, 16, 3, 2, 2, 2,
	120, 121, 7, 93, 2, 2, 121, 18, 3, 2, 2, 2, 122, 123, 7, 95, 2, 2, 123,
	20, 3, 2, 2, 2, 124, 125, 7, 42, 2, 2, 125, 22, 3, 2, 2, 2, 126, 127, 7,
	43, 2, 2, 127, 24, 3, 2, 2, 2, 128, 129, 7, 62, 2, 2, 129, 26, 3, 2, 2,
	2, 130, 131, 7, 64, 2, 2, 131, 28, 3, 2, 2, 2, 132, 133, 7, 96, 2, 2, 133,
	30, 3, 2, 2, 2, 134, 135, 7, 36, 2, 2, 135, 32, 3, 2, 2, 2, 136, 137, 7,
	41, 2, 2, 137, 34, 3, 2, 2, 2, 138, 139, 7, 48, 2, 2, 139, 36, 3, 2, 2,
	2, 140, 141, 7, 61, 2, 2, 141, 38, 3, 2, 2, 2, 142, 143, 7, 60, 2, 2, 143,
	40, 3, 2, 2, 2, 144, 145, 7, 63, 2, 2, 145, 42, 3, 2, 2, 2, 146, 147, 7,
	45, 2, 2, 147, 44, 3, 2, 2, 2, 148, 149, 7, 47, 2, 2, 149, 46, 3, 2, 2,
	2, 150, 151, 7, 126, 2, 2, 151, 48, 3, 2, 2, 2, 152, 153, 7, 48, 2, 2,
	153, 154, 7, 48, 2, 2, 154, 155, 7, 34, 2, 2, 155, 157, 3, 2, 2, 2, 156,
	158, 5, 59, 30, 2, 157, 156, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 164,
	3, 2, 2, 2, 159, 160, 7, 48, 2, 2, 160, 161, 7, 48, 2, 2, 161, 162, 3,
	2, 2, 2, 162, 164, 5, 59, 30, 2, 163, 152, 3, 2, 2, 2, 163, 159, 3, 2,
	2, 2, 164, 50, 3, 2, 2, 2, 165, 166, 7, 97, 2, 2, 166, 52, 3, 2, 2, 2,
	167, 168, 7, 98, 2, 2, 168, 54, 3, 2, 2, 2, 169, 170, 7, 44, 2, 2, 170,
	56, 3, 2, 2, 2, 171, 172, 9, 4, 2, 2, 172, 58, 3, 2, 2, 2, 173, 175, 7,
	15, 2, 2, 174, 173, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 3, 2, 2,
	2, 176, 177, 7, 12, 2, 2, 177, 60, 3, 2, 2, 2, 178, 179, 11, 2, 2, 2, 179,
	62, 3, 2, 2, 2, 14, 2, 73, 79, 85, 87, 94, 105, 110, 115, 157, 163, 174,
	2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "':doc:'", "", "", "", "", "", "", "'['", "']'", "'('", "')'", "'<'",
	"'>'", "'^'", "'\"'", "'''", "'.'", "';'", "':'", "'='", "'+'", "'-'",
	"'|'", "", "'_'", "'`'", "'*'",
}

var lexerSymbolicNames = []string{
	"", "", "SectionSeparator", "Literal", "TimeStar", "Alphabet", "Numbers",
	"Quote", "SquareLeft", "SquareRight", "RoundLeft", "RoundRight", "AngleLeft",
	"AngleRight", "Hat", "QuotationDouble", "QuotationSingle", "Dot", "SemiColon",
	"Colon", "Equal", "Plus", "Minus", "Block", "Comment", "UnderScore", "BackTick",
	"Star", "Space", "LineBreak", "Any",
}

var lexerRuleNames = []string{
	"T__0", "SectionSeparator", "Literal", "TimeStar", "Alphabet", "Numbers",
	"Quote", "SquareLeft", "SquareRight", "RoundLeft", "RoundRight", "AngleLeft",
	"AngleRight", "Hat", "QuotationDouble", "QuotationSingle", "Dot", "SemiColon",
	"Colon", "Equal", "Plus", "Minus", "Block", "Comment", "UnderScore", "BackTick",
	"Star", "Space", "LineBreak", "Any",
}

type ReStructuredTextLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewReStructuredTextLexer(input antlr.CharStream) *ReStructuredTextLexer {

	l := new(ReStructuredTextLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ReStructuredText.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ReStructuredTextLexer tokens.
const (
	ReStructuredTextLexerT__0             = 1
	ReStructuredTextLexerSectionSeparator = 2
	ReStructuredTextLexerLiteral          = 3
	ReStructuredTextLexerTimeStar         = 4
	ReStructuredTextLexerAlphabet         = 5
	ReStructuredTextLexerNumbers          = 6
	ReStructuredTextLexerQuote            = 7
	ReStructuredTextLexerSquareLeft       = 8
	ReStructuredTextLexerSquareRight      = 9
	ReStructuredTextLexerRoundLeft        = 10
	ReStructuredTextLexerRoundRight       = 11
	ReStructuredTextLexerAngleLeft        = 12
	ReStructuredTextLexerAngleRight       = 13
	ReStructuredTextLexerHat              = 14
	ReStructuredTextLexerQuotationDouble  = 15
	ReStructuredTextLexerQuotationSingle  = 16
	ReStructuredTextLexerDot              = 17
	ReStructuredTextLexerSemiColon        = 18
	ReStructuredTextLexerColon            = 19
	ReStructuredTextLexerEqual            = 20
	ReStructuredTextLexerPlus             = 21
	ReStructuredTextLexerMinus            = 22
	ReStructuredTextLexerBlock            = 23
	ReStructuredTextLexerComment          = 24
	ReStructuredTextLexerUnderScore       = 25
	ReStructuredTextLexerBackTick         = 26
	ReStructuredTextLexerStar             = 27
	ReStructuredTextLexerSpace            = 28
	ReStructuredTextLexerLineBreak        = 29
	ReStructuredTextLexerAny              = 30
)
