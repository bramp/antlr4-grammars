// Code generated from guido.g4 by ANTLR 4.7.2. DO NOT EDIT.

package guido // guido
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 219,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 6, 2, 62, 10, 2, 13, 2, 14, 2, 63, 3, 2,
	5, 2, 67, 10, 2, 3, 3, 3, 3, 6, 3, 71, 10, 3, 13, 3, 14, 3, 72, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 7, 4, 80, 10, 4, 12, 4, 14, 4, 83, 11, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 6, 5, 89, 10, 5, 13, 5, 14, 5, 90, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 5, 6, 98, 10, 6, 3, 6, 5, 6, 101, 10, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 113, 10, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 7, 8, 119, 10, 8, 12, 8, 14, 8, 122, 11, 8, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 9, 5, 9, 129, 10, 9, 3, 10, 3, 10, 3, 10, 6, 10, 134, 10, 10,
	13, 10, 14, 10, 135, 3, 11, 3, 11, 3, 11, 6, 11, 141, 10, 11, 13, 11, 14,
	11, 142, 3, 11, 3, 11, 3, 12, 3, 12, 5, 12, 149, 10, 12, 3, 12, 5, 12,
	152, 10, 12, 3, 12, 5, 12, 155, 10, 12, 3, 12, 5, 12, 158, 10, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 7, 13, 164, 10, 13, 12, 13, 14, 13, 167, 11, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 5, 16, 176, 10, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 18, 5, 18, 183, 10, 18, 3, 18, 3, 18, 5, 18, 187,
	10, 18, 3, 19, 5, 19, 190, 10, 19, 3, 19, 3, 19, 3, 20, 6, 20, 195, 10,
	20, 13, 20, 14, 20, 196, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 30, 3, 30, 3, 30, 2, 2, 31, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 2, 6, 3, 2, 32, 33, 4, 2, 31, 31, 33, 33, 3, 2, 13, 14, 3, 2, 15, 16,
	2, 223, 2, 66, 3, 2, 2, 2, 4, 68, 3, 2, 2, 2, 6, 76, 3, 2, 2, 2, 8, 84,
	3, 2, 2, 2, 10, 94, 3, 2, 2, 2, 12, 112, 3, 2, 2, 2, 14, 114, 3, 2, 2,
	2, 16, 128, 3, 2, 2, 2, 18, 130, 3, 2, 2, 2, 20, 137, 3, 2, 2, 2, 22, 146,
	3, 2, 2, 2, 24, 159, 3, 2, 2, 2, 26, 170, 3, 2, 2, 2, 28, 172, 3, 2, 2,
	2, 30, 175, 3, 2, 2, 2, 32, 179, 3, 2, 2, 2, 34, 182, 3, 2, 2, 2, 36, 189,
	3, 2, 2, 2, 38, 194, 3, 2, 2, 2, 40, 198, 3, 2, 2, 2, 42, 200, 3, 2, 2,
	2, 44, 202, 3, 2, 2, 2, 46, 204, 3, 2, 2, 2, 48, 206, 3, 2, 2, 2, 50, 208,
	3, 2, 2, 2, 52, 210, 3, 2, 2, 2, 54, 212, 3, 2, 2, 2, 56, 214, 3, 2, 2,
	2, 58, 216, 3, 2, 2, 2, 60, 62, 5, 4, 3, 2, 61, 60, 3, 2, 2, 2, 62, 63,
	3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2,
	65, 67, 5, 6, 4, 2, 66, 61, 3, 2, 2, 2, 66, 65, 3, 2, 2, 2, 67, 3, 3, 2,
	2, 2, 68, 70, 7, 3, 2, 2, 69, 71, 5, 6, 4, 2, 70, 69, 3, 2, 2, 2, 71, 72,
	3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2,
	74, 75, 7, 4, 2, 2, 75, 5, 3, 2, 2, 2, 76, 81, 5, 8, 5, 2, 77, 78, 7, 5,
	2, 2, 78, 80, 5, 8, 5, 2, 79, 77, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 7, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2,
	84, 88, 7, 6, 2, 2, 85, 89, 5, 10, 6, 2, 86, 89, 5, 22, 12, 2, 87, 89,
	5, 24, 13, 2, 88, 85, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 87, 3, 2, 2,
	2, 89, 90, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 92,
	3, 2, 2, 2, 92, 93, 7, 7, 2, 2, 93, 9, 3, 2, 2, 2, 94, 95, 7, 30, 2, 2,
	95, 97, 5, 12, 7, 2, 96, 98, 5, 14, 8, 2, 97, 96, 3, 2, 2, 2, 97, 98, 3,
	2, 2, 2, 98, 100, 3, 2, 2, 2, 99, 101, 5, 20, 11, 2, 100, 99, 3, 2, 2,
	2, 100, 101, 3, 2, 2, 2, 101, 11, 3, 2, 2, 2, 102, 113, 5, 40, 21, 2, 103,
	113, 5, 42, 22, 2, 104, 113, 5, 44, 23, 2, 105, 113, 5, 46, 24, 2, 106,
	113, 5, 48, 25, 2, 107, 113, 5, 50, 26, 2, 108, 113, 5, 52, 27, 2, 109,
	113, 5, 54, 28, 2, 110, 113, 5, 58, 30, 2, 111, 113, 5, 56, 29, 2, 112,
	102, 3, 2, 2, 2, 112, 103, 3, 2, 2, 2, 112, 104, 3, 2, 2, 2, 112, 105,
	3, 2, 2, 2, 112, 106, 3, 2, 2, 2, 112, 107, 3, 2, 2, 2, 112, 108, 3, 2,
	2, 2, 112, 109, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 111, 3, 2, 2, 2,
	113, 13, 3, 2, 2, 2, 114, 115, 7, 8, 2, 2, 115, 120, 5, 16, 9, 2, 116,
	117, 7, 5, 2, 2, 117, 119, 5, 16, 9, 2, 118, 116, 3, 2, 2, 2, 119, 122,
	3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 123, 3, 2,
	2, 2, 122, 120, 3, 2, 2, 2, 123, 124, 7, 9, 2, 2, 124, 15, 3, 2, 2, 2,
	125, 129, 7, 34, 2, 2, 126, 129, 5, 30, 16, 2, 127, 129, 5, 18, 10, 2,
	128, 125, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129,
	17, 3, 2, 2, 2, 130, 131, 7, 33, 2, 2, 131, 133, 7, 10, 2, 2, 132, 134,
	9, 2, 2, 2, 133, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 133, 3, 2,
	2, 2, 135, 136, 3, 2, 2, 2, 136, 19, 3, 2, 2, 2, 137, 140, 7, 11, 2, 2,
	138, 141, 5, 22, 12, 2, 139, 141, 5, 24, 13, 2, 140, 138, 3, 2, 2, 2, 140,
	139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143,
	3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 7, 12, 2, 2, 145, 21, 3, 2,
	2, 2, 146, 148, 5, 26, 14, 2, 147, 149, 5, 28, 15, 2, 148, 147, 3, 2, 2,
	2, 148, 149, 3, 2, 2, 2, 149, 151, 3, 2, 2, 2, 150, 152, 5, 32, 17, 2,
	151, 150, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 154, 3, 2, 2, 2, 153,
	155, 5, 36, 19, 2, 154, 153, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 157,
	3, 2, 2, 2, 156, 158, 5, 38, 20, 2, 157, 156, 3, 2, 2, 2, 157, 158, 3,
	2, 2, 2, 158, 23, 3, 2, 2, 2, 159, 160, 7, 3, 2, 2, 160, 165, 5, 22, 12,
	2, 161, 162, 7, 5, 2, 2, 162, 164, 5, 22, 12, 2, 163, 161, 3, 2, 2, 2,
	164, 167, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166,
	168, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 168, 169, 7, 4, 2, 2, 169, 25, 3,
	2, 2, 2, 170, 171, 9, 3, 2, 2, 171, 27, 3, 2, 2, 2, 172, 173, 9, 4, 2,
	2, 173, 29, 3, 2, 2, 2, 174, 176, 9, 5, 2, 2, 175, 174, 3, 2, 2, 2, 175,
	176, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 178, 7, 32, 2, 2, 178, 31,
	3, 2, 2, 2, 179, 180, 5, 30, 16, 2, 180, 33, 3, 2, 2, 2, 181, 183, 5, 30,
	16, 2, 182, 181, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2,
	184, 185, 7, 17, 2, 2, 185, 187, 5, 30, 16, 2, 186, 184, 3, 2, 2, 2, 186,
	187, 3, 2, 2, 2, 187, 35, 3, 2, 2, 2, 188, 190, 7, 18, 2, 2, 189, 188,
	3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192, 5, 34,
	18, 2, 192, 37, 3, 2, 2, 2, 193, 195, 7, 19, 2, 2, 194, 193, 3, 2, 2, 2,
	195, 196, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197,
	39, 3, 2, 2, 2, 198, 199, 7, 20, 2, 2, 199, 41, 3, 2, 2, 2, 200, 201, 7,
	21, 2, 2, 201, 43, 3, 2, 2, 2, 202, 203, 7, 22, 2, 2, 203, 45, 3, 2, 2,
	2, 204, 205, 7, 23, 2, 2, 205, 47, 3, 2, 2, 2, 206, 207, 7, 24, 2, 2, 207,
	49, 3, 2, 2, 2, 208, 209, 7, 25, 2, 2, 209, 51, 3, 2, 2, 2, 210, 211, 7,
	26, 2, 2, 211, 53, 3, 2, 2, 2, 212, 213, 7, 27, 2, 2, 213, 55, 3, 2, 2,
	2, 214, 215, 7, 28, 2, 2, 215, 57, 3, 2, 2, 2, 216, 217, 7, 29, 2, 2, 217,
	59, 3, 2, 2, 2, 26, 63, 66, 72, 81, 88, 90, 97, 100, 112, 120, 128, 135,
	140, 142, 148, 151, 154, 157, 165, 175, 182, 186, 189, 196,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "','", "'['", "']'", "'<'", "'>'", "'='", "'('", "')'",
	"'#'", "'&'", "'-'", "'+'", "'/'", "'*'", "'.'", "'title'", "'tempo'",
	"'clef'", "'meter'", "'slur'", "'key'", "'barFormat'", "'staff'", "'repeatEnd'",
	"'t'", "'\\'", "'_'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "TAGSTART", "REST", "NUMBER", "STRING",
	"STRINGLITERAL", "COMMENT", "WS",
}

var ruleNames = []string{
	"prog", "segment", "sequencelist", "sequence", "tag", "tagname", "parameters",
	"parameter", "kvpair", "notes", "note", "chord", "notename", "accidental",
	"number", "octave", "fraction", "duration", "dotting", "title", "tempo",
	"clef", "meter", "slur", "key", "barformat", "staff", "repeatEnd", "t",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type guidoParser struct {
	*antlr.BaseParser
}

func NewguidoParser(input antlr.TokenStream) *guidoParser {
	this := new(guidoParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "guido.g4"

	return this
}

// guidoParser tokens.
const (
	guidoParserEOF           = antlr.TokenEOF
	guidoParserT__0          = 1
	guidoParserT__1          = 2
	guidoParserT__2          = 3
	guidoParserT__3          = 4
	guidoParserT__4          = 5
	guidoParserT__5          = 6
	guidoParserT__6          = 7
	guidoParserT__7          = 8
	guidoParserT__8          = 9
	guidoParserT__9          = 10
	guidoParserT__10         = 11
	guidoParserT__11         = 12
	guidoParserT__12         = 13
	guidoParserT__13         = 14
	guidoParserT__14         = 15
	guidoParserT__15         = 16
	guidoParserT__16         = 17
	guidoParserT__17         = 18
	guidoParserT__18         = 19
	guidoParserT__19         = 20
	guidoParserT__20         = 21
	guidoParserT__21         = 22
	guidoParserT__22         = 23
	guidoParserT__23         = 24
	guidoParserT__24         = 25
	guidoParserT__25         = 26
	guidoParserT__26         = 27
	guidoParserTAGSTART      = 28
	guidoParserREST          = 29
	guidoParserNUMBER        = 30
	guidoParserSTRING        = 31
	guidoParserSTRINGLITERAL = 32
	guidoParserCOMMENT       = 33
	guidoParserWS            = 34
)

// guidoParser rules.
const (
	guidoParserRULE_prog         = 0
	guidoParserRULE_segment      = 1
	guidoParserRULE_sequencelist = 2
	guidoParserRULE_sequence     = 3
	guidoParserRULE_tag          = 4
	guidoParserRULE_tagname      = 5
	guidoParserRULE_parameters   = 6
	guidoParserRULE_parameter    = 7
	guidoParserRULE_kvpair       = 8
	guidoParserRULE_notes        = 9
	guidoParserRULE_note         = 10
	guidoParserRULE_chord        = 11
	guidoParserRULE_notename     = 12
	guidoParserRULE_accidental   = 13
	guidoParserRULE_number       = 14
	guidoParserRULE_octave       = 15
	guidoParserRULE_fraction     = 16
	guidoParserRULE_duration     = 17
	guidoParserRULE_dotting      = 18
	guidoParserRULE_title        = 19
	guidoParserRULE_tempo        = 20
	guidoParserRULE_clef         = 21
	guidoParserRULE_meter        = 22
	guidoParserRULE_slur         = 23
	guidoParserRULE_key          = 24
	guidoParserRULE_barformat    = 25
	guidoParserRULE_staff        = 26
	guidoParserRULE_repeatEnd    = 27
	guidoParserRULE_t            = 28
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllSegment() []ISegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISegmentContext)(nil)).Elem())
	var tst = make([]ISegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISegmentContext)
		}
	}

	return tst
}

func (s *ProgContext) Segment(i int) ISegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISegmentContext)
}

func (s *ProgContext) Sequencelist() ISequencelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequencelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequencelistContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *guidoParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, guidoParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(64)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case guidoParserT__0:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == guidoParserT__0 {
			{
				p.SetState(58)
				p.Segment()
			}

			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case guidoParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Sequencelist()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISegmentContext is an interface to support dynamic dispatch.
type ISegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentContext differentiates from other interfaces.
	IsSegmentContext()
}

type SegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentContext() *SegmentContext {
	var p = new(SegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_segment
	return p
}

func (*SegmentContext) IsSegmentContext() {}

func NewSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentContext {
	var p = new(SegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_segment

	return p
}

func (s *SegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentContext) AllSequencelist() []ISequencelistContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequencelistContext)(nil)).Elem())
	var tst = make([]ISequencelistContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequencelistContext)
		}
	}

	return tst
}

func (s *SegmentContext) Sequencelist(i int) ISequencelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequencelistContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequencelistContext)
}

func (s *SegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterSegment(s)
	}
}

func (s *SegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitSegment(s)
	}
}

func (p *guidoParser) Segment() (localctx ISegmentContext) {
	localctx = NewSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, guidoParserRULE_segment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(guidoParserT__0)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == guidoParserT__3 {
		{
			p.SetState(67)
			p.Sequencelist()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(guidoParserT__1)
	}

	return localctx
}

// ISequencelistContext is an interface to support dynamic dispatch.
type ISequencelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequencelistContext differentiates from other interfaces.
	IsSequencelistContext()
}

type SequencelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequencelistContext() *SequencelistContext {
	var p = new(SequencelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_sequencelist
	return p
}

func (*SequencelistContext) IsSequencelistContext() {}

func NewSequencelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequencelistContext {
	var p = new(SequencelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_sequencelist

	return p
}

func (s *SequencelistContext) GetParser() antlr.Parser { return s.parser }

func (s *SequencelistContext) AllSequence() []ISequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequenceContext)(nil)).Elem())
	var tst = make([]ISequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequenceContext)
		}
	}

	return tst
}

func (s *SequencelistContext) Sequence(i int) ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *SequencelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequencelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequencelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterSequencelist(s)
	}
}

func (s *SequencelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitSequencelist(s)
	}
}

func (p *guidoParser) Sequencelist() (localctx ISequencelistContext) {
	localctx = NewSequencelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, guidoParserRULE_sequencelist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Sequence()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == guidoParserT__2 {
		{
			p.SetState(75)
			p.Match(guidoParserT__2)
		}
		{
			p.SetState(76)
			p.Sequence()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_sequence
	return p
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceContext {
	var p = new(SequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_sequence

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) AllTag() []ITagContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagContext)(nil)).Elem())
	var tst = make([]ITagContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagContext)
		}
	}

	return tst
}

func (s *SequenceContext) Tag(i int) ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *SequenceContext) AllNote() []INoteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INoteContext)(nil)).Elem())
	var tst = make([]INoteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INoteContext)
		}
	}

	return tst
}

func (s *SequenceContext) Note(i int) INoteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INoteContext)
}

func (s *SequenceContext) AllChord() []IChordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IChordContext)(nil)).Elem())
	var tst = make([]IChordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IChordContext)
		}
	}

	return tst
}

func (s *SequenceContext) Chord(i int) IChordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IChordContext)
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitSequence(s)
	}
}

func (p *guidoParser) Sequence() (localctx ISequenceContext) {
	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, guidoParserRULE_sequence)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(guidoParserT__3)
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<guidoParserT__0)|(1<<guidoParserTAGSTART)|(1<<guidoParserREST)|(1<<guidoParserSTRING))) != 0) {
		p.SetState(86)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case guidoParserTAGSTART:
			{
				p.SetState(83)
				p.Tag()
			}

		case guidoParserREST, guidoParserSTRING:
			{
				p.SetState(84)
				p.Note()
			}

		case guidoParserT__0:
			{
				p.SetState(85)
				p.Chord()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(guidoParserT__4)
	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) TAGSTART() antlr.TerminalNode {
	return s.GetToken(guidoParserTAGSTART, 0)
}

func (s *TagContext) Tagname() ITagnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagnameContext)
}

func (s *TagContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *TagContext) Notes() INotesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotesContext)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterTag(s)
	}
}

func (s *TagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitTag(s)
	}
}

func (p *guidoParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, guidoParserRULE_tag)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(guidoParserTAGSTART)
	}
	{
		p.SetState(93)
		p.Tagname()
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == guidoParserT__5 {
		{
			p.SetState(94)
			p.Parameters()
		}

	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == guidoParserT__8 {
		{
			p.SetState(97)
			p.Notes()
		}

	}

	return localctx
}

// ITagnameContext is an interface to support dynamic dispatch.
type ITagnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagnameContext differentiates from other interfaces.
	IsTagnameContext()
}

type TagnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagnameContext() *TagnameContext {
	var p = new(TagnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_tagname
	return p
}

func (*TagnameContext) IsTagnameContext() {}

func NewTagnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagnameContext {
	var p = new(TagnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_tagname

	return p
}

func (s *TagnameContext) GetParser() antlr.Parser { return s.parser }

func (s *TagnameContext) Title() ITitleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITitleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITitleContext)
}

func (s *TagnameContext) Tempo() ITempoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITempoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITempoContext)
}

func (s *TagnameContext) Clef() IClefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClefContext)
}

func (s *TagnameContext) Meter() IMeterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeterContext)
}

func (s *TagnameContext) Slur() ISlurContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISlurContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISlurContext)
}

func (s *TagnameContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *TagnameContext) Barformat() IBarformatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBarformatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBarformatContext)
}

func (s *TagnameContext) Staff() IStaffContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStaffContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStaffContext)
}

func (s *TagnameContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *TagnameContext) RepeatEnd() IRepeatEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepeatEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRepeatEndContext)
}

func (s *TagnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterTagname(s)
	}
}

func (s *TagnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitTagname(s)
	}
}

func (p *guidoParser) Tagname() (localctx ITagnameContext) {
	localctx = NewTagnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, guidoParserRULE_tagname)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case guidoParserT__17:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Title()
		}

	case guidoParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Tempo()
		}

	case guidoParserT__19:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)
			p.Clef()
		}

	case guidoParserT__20:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.Meter()
		}

	case guidoParserT__21:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(104)
			p.Slur()
		}

	case guidoParserT__22:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(105)
			p.Key()
		}

	case guidoParserT__23:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(106)
			p.Barformat()
		}

	case guidoParserT__24:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(107)
			p.Staff()
		}

	case guidoParserT__26:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(108)
			p.T()
		}

	case guidoParserT__25:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(109)
			p.RepeatEnd()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParametersContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *guidoParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, guidoParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(guidoParserT__5)
	}
	{
		p.SetState(113)
		p.Parameter()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == guidoParserT__2 {
		{
			p.SetState(114)
			p.Match(guidoParserT__2)
		}
		{
			p.SetState(115)
			p.Parameter()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(guidoParserT__6)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) STRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(guidoParserSTRINGLITERAL, 0)
}

func (s *ParameterContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ParameterContext) Kvpair() IKvpairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKvpairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKvpairContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *guidoParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, guidoParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case guidoParserSTRINGLITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(guidoParserSTRINGLITERAL)
		}

	case guidoParserT__12, guidoParserT__13, guidoParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Number()
		}

	case guidoParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Kvpair()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKvpairContext is an interface to support dynamic dispatch.
type IKvpairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKvpairContext differentiates from other interfaces.
	IsKvpairContext()
}

type KvpairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKvpairContext() *KvpairContext {
	var p = new(KvpairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_kvpair
	return p
}

func (*KvpairContext) IsKvpairContext() {}

func NewKvpairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KvpairContext {
	var p = new(KvpairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_kvpair

	return p
}

func (s *KvpairContext) GetParser() antlr.Parser { return s.parser }

func (s *KvpairContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(guidoParserSTRING)
}

func (s *KvpairContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(guidoParserSTRING, i)
}

func (s *KvpairContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(guidoParserNUMBER)
}

func (s *KvpairContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(guidoParserNUMBER, i)
}

func (s *KvpairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KvpairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KvpairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterKvpair(s)
	}
}

func (s *KvpairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitKvpair(s)
	}
}

func (p *guidoParser) Kvpair() (localctx IKvpairContext) {
	localctx = NewKvpairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, guidoParserRULE_kvpair)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(guidoParserSTRING)
	}
	{
		p.SetState(129)
		p.Match(guidoParserT__7)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == guidoParserNUMBER || _la == guidoParserSTRING {
		{
			p.SetState(130)
			_la = p.GetTokenStream().LA(1)

			if !(_la == guidoParserNUMBER || _la == guidoParserSTRING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INotesContext is an interface to support dynamic dispatch.
type INotesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotesContext differentiates from other interfaces.
	IsNotesContext()
}

type NotesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotesContext() *NotesContext {
	var p = new(NotesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_notes
	return p
}

func (*NotesContext) IsNotesContext() {}

func NewNotesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotesContext {
	var p = new(NotesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_notes

	return p
}

func (s *NotesContext) GetParser() antlr.Parser { return s.parser }

func (s *NotesContext) AllNote() []INoteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INoteContext)(nil)).Elem())
	var tst = make([]INoteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INoteContext)
		}
	}

	return tst
}

func (s *NotesContext) Note(i int) INoteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INoteContext)
}

func (s *NotesContext) AllChord() []IChordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IChordContext)(nil)).Elem())
	var tst = make([]IChordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IChordContext)
		}
	}

	return tst
}

func (s *NotesContext) Chord(i int) IChordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IChordContext)
}

func (s *NotesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterNotes(s)
	}
}

func (s *NotesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitNotes(s)
	}
}

func (p *guidoParser) Notes() (localctx INotesContext) {
	localctx = NewNotesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, guidoParserRULE_notes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(guidoParserT__8)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<guidoParserT__0)|(1<<guidoParserREST)|(1<<guidoParserSTRING))) != 0) {
		p.SetState(138)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case guidoParserREST, guidoParserSTRING:
			{
				p.SetState(136)
				p.Note()
			}

		case guidoParserT__0:
			{
				p.SetState(137)
				p.Chord()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(142)
		p.Match(guidoParserT__9)
	}

	return localctx
}

// INoteContext is an interface to support dynamic dispatch.
type INoteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoteContext differentiates from other interfaces.
	IsNoteContext()
}

type NoteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoteContext() *NoteContext {
	var p = new(NoteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_note
	return p
}

func (*NoteContext) IsNoteContext() {}

func NewNoteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoteContext {
	var p = new(NoteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_note

	return p
}

func (s *NoteContext) GetParser() antlr.Parser { return s.parser }

func (s *NoteContext) Notename() INotenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotenameContext)
}

func (s *NoteContext) Accidental() IAccidentalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccidentalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccidentalContext)
}

func (s *NoteContext) Octave() IOctaveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOctaveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOctaveContext)
}

func (s *NoteContext) Duration() IDurationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDurationContext)
}

func (s *NoteContext) Dotting() IDottingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDottingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDottingContext)
}

func (s *NoteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterNote(s)
	}
}

func (s *NoteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitNote(s)
	}
}

func (p *guidoParser) Note() (localctx INoteContext) {
	localctx = NewNoteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, guidoParserRULE_note)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Notename()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == guidoParserT__10 || _la == guidoParserT__11 {
		{
			p.SetState(145)
			p.Accidental()
		}

	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(148)
			p.Octave()
		}

	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(151)
			p.Duration()
		}

	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == guidoParserT__16 {
		{
			p.SetState(154)
			p.Dotting()
		}

	}

	return localctx
}

// IChordContext is an interface to support dynamic dispatch.
type IChordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChordContext differentiates from other interfaces.
	IsChordContext()
}

type ChordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChordContext() *ChordContext {
	var p = new(ChordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_chord
	return p
}

func (*ChordContext) IsChordContext() {}

func NewChordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChordContext {
	var p = new(ChordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_chord

	return p
}

func (s *ChordContext) GetParser() antlr.Parser { return s.parser }

func (s *ChordContext) AllNote() []INoteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INoteContext)(nil)).Elem())
	var tst = make([]INoteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INoteContext)
		}
	}

	return tst
}

func (s *ChordContext) Note(i int) INoteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INoteContext)
}

func (s *ChordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterChord(s)
	}
}

func (s *ChordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitChord(s)
	}
}

func (p *guidoParser) Chord() (localctx IChordContext) {
	localctx = NewChordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, guidoParserRULE_chord)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(guidoParserT__0)
	}
	{
		p.SetState(158)
		p.Note()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == guidoParserT__2 {
		{
			p.SetState(159)
			p.Match(guidoParserT__2)
		}
		{
			p.SetState(160)
			p.Note()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.Match(guidoParserT__1)
	}

	return localctx
}

// INotenameContext is an interface to support dynamic dispatch.
type INotenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotenameContext differentiates from other interfaces.
	IsNotenameContext()
}

type NotenameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotenameContext() *NotenameContext {
	var p = new(NotenameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_notename
	return p
}

func (*NotenameContext) IsNotenameContext() {}

func NewNotenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotenameContext {
	var p = new(NotenameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_notename

	return p
}

func (s *NotenameContext) GetParser() antlr.Parser { return s.parser }

func (s *NotenameContext) STRING() antlr.TerminalNode {
	return s.GetToken(guidoParserSTRING, 0)
}

func (s *NotenameContext) REST() antlr.TerminalNode {
	return s.GetToken(guidoParserREST, 0)
}

func (s *NotenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotenameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterNotename(s)
	}
}

func (s *NotenameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitNotename(s)
	}
}

func (p *guidoParser) Notename() (localctx INotenameContext) {
	localctx = NewNotenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, guidoParserRULE_notename)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		_la = p.GetTokenStream().LA(1)

		if !(_la == guidoParserREST || _la == guidoParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAccidentalContext is an interface to support dynamic dispatch.
type IAccidentalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccidentalContext differentiates from other interfaces.
	IsAccidentalContext()
}

type AccidentalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccidentalContext() *AccidentalContext {
	var p = new(AccidentalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_accidental
	return p
}

func (*AccidentalContext) IsAccidentalContext() {}

func NewAccidentalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccidentalContext {
	var p = new(AccidentalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_accidental

	return p
}

func (s *AccidentalContext) GetParser() antlr.Parser { return s.parser }
func (s *AccidentalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccidentalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccidentalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterAccidental(s)
	}
}

func (s *AccidentalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitAccidental(s)
	}
}

func (p *guidoParser) Accidental() (localctx IAccidentalContext) {
	localctx = NewAccidentalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, guidoParserRULE_accidental)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		_la = p.GetTokenStream().LA(1)

		if !(_la == guidoParserT__10 || _la == guidoParserT__11) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(guidoParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *guidoParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, guidoParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == guidoParserT__12 || _la == guidoParserT__13 {
		{
			p.SetState(172)
			_la = p.GetTokenStream().LA(1)

			if !(_la == guidoParserT__12 || _la == guidoParserT__13) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(175)
		p.Match(guidoParserNUMBER)
	}

	return localctx
}

// IOctaveContext is an interface to support dynamic dispatch.
type IOctaveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOctaveContext differentiates from other interfaces.
	IsOctaveContext()
}

type OctaveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOctaveContext() *OctaveContext {
	var p = new(OctaveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_octave
	return p
}

func (*OctaveContext) IsOctaveContext() {}

func NewOctaveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OctaveContext {
	var p = new(OctaveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_octave

	return p
}

func (s *OctaveContext) GetParser() antlr.Parser { return s.parser }

func (s *OctaveContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *OctaveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctaveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OctaveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterOctave(s)
	}
}

func (s *OctaveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitOctave(s)
	}
}

func (p *guidoParser) Octave() (localctx IOctaveContext) {
	localctx = NewOctaveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, guidoParserRULE_octave)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Number()
	}

	return localctx
}

// IFractionContext is an interface to support dynamic dispatch.
type IFractionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFractionContext differentiates from other interfaces.
	IsFractionContext()
}

type FractionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFractionContext() *FractionContext {
	var p = new(FractionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_fraction
	return p
}

func (*FractionContext) IsFractionContext() {}

func NewFractionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FractionContext {
	var p = new(FractionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_fraction

	return p
}

func (s *FractionContext) GetParser() antlr.Parser { return s.parser }

func (s *FractionContext) AllNumber() []INumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberContext)(nil)).Elem())
	var tst = make([]INumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberContext)
		}
	}

	return tst
}

func (s *FractionContext) Number(i int) INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *FractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FractionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterFraction(s)
	}
}

func (s *FractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitFraction(s)
	}
}

func (p *guidoParser) Fraction() (localctx IFractionContext) {
	localctx = NewFractionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, guidoParserRULE_fraction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<guidoParserT__12)|(1<<guidoParserT__13)|(1<<guidoParserNUMBER))) != 0 {
		{
			p.SetState(179)
			p.Number()
		}

	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == guidoParserT__14 {
		{
			p.SetState(182)
			p.Match(guidoParserT__14)
		}
		{
			p.SetState(183)
			p.Number()
		}

	}

	return localctx
}

// IDurationContext is an interface to support dynamic dispatch.
type IDurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDurationContext differentiates from other interfaces.
	IsDurationContext()
}

type DurationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationContext() *DurationContext {
	var p = new(DurationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_duration
	return p
}

func (*DurationContext) IsDurationContext() {}

func NewDurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationContext {
	var p = new(DurationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_duration

	return p
}

func (s *DurationContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationContext) Fraction() IFractionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFractionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFractionContext)
}

func (s *DurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterDuration(s)
	}
}

func (s *DurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitDuration(s)
	}
}

func (p *guidoParser) Duration() (localctx IDurationContext) {
	localctx = NewDurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, guidoParserRULE_duration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == guidoParserT__15 {
		{
			p.SetState(186)
			p.Match(guidoParserT__15)
		}

	}
	{
		p.SetState(189)
		p.Fraction()
	}

	return localctx
}

// IDottingContext is an interface to support dynamic dispatch.
type IDottingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDottingContext differentiates from other interfaces.
	IsDottingContext()
}

type DottingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDottingContext() *DottingContext {
	var p = new(DottingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_dotting
	return p
}

func (*DottingContext) IsDottingContext() {}

func NewDottingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DottingContext {
	var p = new(DottingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_dotting

	return p
}

func (s *DottingContext) GetParser() antlr.Parser { return s.parser }
func (s *DottingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DottingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DottingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterDotting(s)
	}
}

func (s *DottingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitDotting(s)
	}
}

func (p *guidoParser) Dotting() (localctx IDottingContext) {
	localctx = NewDottingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, guidoParserRULE_dotting)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == guidoParserT__16 {
		{
			p.SetState(191)
			p.Match(guidoParserT__16)
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITitleContext is an interface to support dynamic dispatch.
type ITitleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTitleContext differentiates from other interfaces.
	IsTitleContext()
}

type TitleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTitleContext() *TitleContext {
	var p = new(TitleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_title
	return p
}

func (*TitleContext) IsTitleContext() {}

func NewTitleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TitleContext {
	var p = new(TitleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_title

	return p
}

func (s *TitleContext) GetParser() antlr.Parser { return s.parser }
func (s *TitleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TitleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TitleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterTitle(s)
	}
}

func (s *TitleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitTitle(s)
	}
}

func (p *guidoParser) Title() (localctx ITitleContext) {
	localctx = NewTitleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, guidoParserRULE_title)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(guidoParserT__17)
	}

	return localctx
}

// ITempoContext is an interface to support dynamic dispatch.
type ITempoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTempoContext differentiates from other interfaces.
	IsTempoContext()
}

type TempoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTempoContext() *TempoContext {
	var p = new(TempoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_tempo
	return p
}

func (*TempoContext) IsTempoContext() {}

func NewTempoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TempoContext {
	var p = new(TempoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_tempo

	return p
}

func (s *TempoContext) GetParser() antlr.Parser { return s.parser }
func (s *TempoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TempoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TempoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterTempo(s)
	}
}

func (s *TempoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitTempo(s)
	}
}

func (p *guidoParser) Tempo() (localctx ITempoContext) {
	localctx = NewTempoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, guidoParserRULE_tempo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(guidoParserT__18)
	}

	return localctx
}

// IClefContext is an interface to support dynamic dispatch.
type IClefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClefContext differentiates from other interfaces.
	IsClefContext()
}

type ClefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClefContext() *ClefContext {
	var p = new(ClefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_clef
	return p
}

func (*ClefContext) IsClefContext() {}

func NewClefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClefContext {
	var p = new(ClefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_clef

	return p
}

func (s *ClefContext) GetParser() antlr.Parser { return s.parser }
func (s *ClefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterClef(s)
	}
}

func (s *ClefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitClef(s)
	}
}

func (p *guidoParser) Clef() (localctx IClefContext) {
	localctx = NewClefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, guidoParserRULE_clef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(guidoParserT__19)
	}

	return localctx
}

// IMeterContext is an interface to support dynamic dispatch.
type IMeterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeterContext differentiates from other interfaces.
	IsMeterContext()
}

type MeterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeterContext() *MeterContext {
	var p = new(MeterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_meter
	return p
}

func (*MeterContext) IsMeterContext() {}

func NewMeterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MeterContext {
	var p = new(MeterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_meter

	return p
}

func (s *MeterContext) GetParser() antlr.Parser { return s.parser }
func (s *MeterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MeterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterMeter(s)
	}
}

func (s *MeterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitMeter(s)
	}
}

func (p *guidoParser) Meter() (localctx IMeterContext) {
	localctx = NewMeterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, guidoParserRULE_meter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(guidoParserT__20)
	}

	return localctx
}

// ISlurContext is an interface to support dynamic dispatch.
type ISlurContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSlurContext differentiates from other interfaces.
	IsSlurContext()
}

type SlurContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlurContext() *SlurContext {
	var p = new(SlurContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_slur
	return p
}

func (*SlurContext) IsSlurContext() {}

func NewSlurContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SlurContext {
	var p = new(SlurContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_slur

	return p
}

func (s *SlurContext) GetParser() antlr.Parser { return s.parser }
func (s *SlurContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlurContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SlurContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterSlur(s)
	}
}

func (s *SlurContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitSlur(s)
	}
}

func (p *guidoParser) Slur() (localctx ISlurContext) {
	localctx = NewSlurContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, guidoParserRULE_slur)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(guidoParserT__21)
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }
func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *guidoParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, guidoParserRULE_key)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(guidoParserT__22)
	}

	return localctx
}

// IBarformatContext is an interface to support dynamic dispatch.
type IBarformatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBarformatContext differentiates from other interfaces.
	IsBarformatContext()
}

type BarformatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBarformatContext() *BarformatContext {
	var p = new(BarformatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_barformat
	return p
}

func (*BarformatContext) IsBarformatContext() {}

func NewBarformatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BarformatContext {
	var p = new(BarformatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_barformat

	return p
}

func (s *BarformatContext) GetParser() antlr.Parser { return s.parser }
func (s *BarformatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BarformatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BarformatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterBarformat(s)
	}
}

func (s *BarformatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitBarformat(s)
	}
}

func (p *guidoParser) Barformat() (localctx IBarformatContext) {
	localctx = NewBarformatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, guidoParserRULE_barformat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(guidoParserT__23)
	}

	return localctx
}

// IStaffContext is an interface to support dynamic dispatch.
type IStaffContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStaffContext differentiates from other interfaces.
	IsStaffContext()
}

type StaffContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStaffContext() *StaffContext {
	var p = new(StaffContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_staff
	return p
}

func (*StaffContext) IsStaffContext() {}

func NewStaffContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StaffContext {
	var p = new(StaffContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_staff

	return p
}

func (s *StaffContext) GetParser() antlr.Parser { return s.parser }
func (s *StaffContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StaffContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StaffContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterStaff(s)
	}
}

func (s *StaffContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitStaff(s)
	}
}

func (p *guidoParser) Staff() (localctx IStaffContext) {
	localctx = NewStaffContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, guidoParserRULE_staff)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(guidoParserT__24)
	}

	return localctx
}

// IRepeatEndContext is an interface to support dynamic dispatch.
type IRepeatEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRepeatEndContext differentiates from other interfaces.
	IsRepeatEndContext()
}

type RepeatEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatEndContext() *RepeatEndContext {
	var p = new(RepeatEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_repeatEnd
	return p
}

func (*RepeatEndContext) IsRepeatEndContext() {}

func NewRepeatEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatEndContext {
	var p = new(RepeatEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_repeatEnd

	return p
}

func (s *RepeatEndContext) GetParser() antlr.Parser { return s.parser }
func (s *RepeatEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterRepeatEnd(s)
	}
}

func (s *RepeatEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitRepeatEnd(s)
	}
}

func (p *guidoParser) RepeatEnd() (localctx IRepeatEndContext) {
	localctx = NewRepeatEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, guidoParserRULE_repeatEnd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(guidoParserT__25)
	}

	return localctx
}

// ITContext is an interface to support dynamic dispatch.
type ITContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTContext differentiates from other interfaces.
	IsTContext()
}

type TContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTContext() *TContext {
	var p = new(TContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = guidoParserRULE_t
	return p
}

func (*TContext) IsTContext() {}

func NewTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TContext {
	var p = new(TContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = guidoParserRULE_t

	return p
}

func (s *TContext) GetParser() antlr.Parser { return s.parser }
func (s *TContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.EnterT(s)
	}
}

func (s *TContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(guidoListener); ok {
		listenerT.ExitT(s)
	}
}

func (p *guidoParser) T() (localctx ITContext) {
	localctx = NewTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, guidoParserRULE_t)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(guidoParserT__26)
	}

	return localctx
}
