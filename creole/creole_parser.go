// Generated from creole.g4 by ANTLR 4.7.

package creole // creole
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 181,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 5, 2, 32, 10, 2, 3, 2, 7, 2, 35,
	10, 2, 12, 2, 14, 2, 38, 11, 2, 3, 3, 6, 3, 41, 10, 3, 13, 3, 14, 3, 42,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	56, 10, 4, 3, 5, 6, 5, 59, 10, 5, 13, 5, 14, 5, 60, 3, 5, 3, 5, 7, 5, 65,
	10, 5, 12, 5, 14, 5, 68, 11, 5, 3, 6, 3, 6, 6, 6, 72, 10, 6, 13, 6, 14,
	6, 73, 3, 6, 5, 6, 77, 10, 6, 3, 7, 3, 7, 3, 7, 6, 7, 82, 10, 7, 13, 7,
	14, 7, 83, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8, 93, 10, 8, 13,
	8, 14, 8, 94, 5, 8, 97, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8,
	105, 10, 8, 13, 8, 14, 8, 106, 3, 8, 3, 8, 5, 8, 111, 10, 8, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 6, 11, 120, 10, 11, 13, 11, 14, 11, 121,
	3, 11, 3, 11, 6, 11, 126, 10, 11, 13, 11, 14, 11, 127, 3, 11, 5, 11, 131,
	10, 11, 3, 12, 3, 12, 6, 12, 135, 10, 12, 13, 12, 14, 12, 136, 6, 12, 139,
	10, 12, 13, 12, 14, 12, 140, 3, 12, 3, 12, 7, 12, 145, 10, 12, 12, 12,
	14, 12, 148, 11, 12, 3, 13, 3, 13, 6, 13, 152, 10, 13, 13, 13, 14, 13,
	153, 6, 13, 156, 10, 13, 13, 13, 14, 13, 157, 3, 13, 3, 13, 7, 13, 162,
	10, 13, 12, 13, 14, 13, 165, 11, 13, 3, 14, 6, 14, 168, 10, 14, 13, 14,
	14, 14, 169, 3, 14, 3, 14, 7, 14, 174, 10, 14, 12, 14, 14, 14, 177, 11,
	14, 3, 15, 3, 15, 3, 15, 2, 2, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 2, 3, 4, 2, 15, 15, 19, 19, 2, 199, 2, 36, 3, 2, 2, 2,
	4, 40, 3, 2, 2, 2, 6, 55, 3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 69, 3, 2,
	2, 2, 12, 78, 3, 2, 2, 2, 14, 110, 3, 2, 2, 2, 16, 112, 3, 2, 2, 2, 18,
	116, 3, 2, 2, 2, 20, 130, 3, 2, 2, 2, 22, 138, 3, 2, 2, 2, 24, 155, 3,
	2, 2, 2, 26, 167, 3, 2, 2, 2, 28, 178, 3, 2, 2, 2, 30, 32, 5, 4, 3, 2,
	31, 30, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 35, 7,
	17, 2, 2, 34, 31, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36,
	37, 3, 2, 2, 2, 37, 3, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39, 41, 5, 6, 4,
	2, 40, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43,
	3, 2, 2, 2, 43, 5, 3, 2, 2, 2, 44, 56, 5, 10, 6, 2, 45, 56, 5, 12, 7, 2,
	46, 56, 5, 14, 8, 2, 47, 56, 5, 26, 14, 2, 48, 56, 5, 18, 10, 2, 49, 56,
	5, 8, 5, 2, 50, 56, 5, 20, 11, 2, 51, 56, 5, 16, 9, 2, 52, 56, 5, 24, 13,
	2, 53, 56, 5, 22, 12, 2, 54, 56, 5, 28, 15, 2, 55, 44, 3, 2, 2, 2, 55,
	45, 3, 2, 2, 2, 55, 46, 3, 2, 2, 2, 55, 47, 3, 2, 2, 2, 55, 48, 3, 2, 2,
	2, 55, 49, 3, 2, 2, 2, 55, 50, 3, 2, 2, 2, 55, 51, 3, 2, 2, 2, 55, 52,
	3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 54, 3, 2, 2, 2, 56, 7, 3, 2, 2, 2,
	57, 59, 9, 2, 2, 2, 58, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3,
	2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 66, 3, 2, 2, 2, 62, 63, 7, 3, 2, 2, 63,
	65, 5, 8, 5, 2, 64, 62, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2,
	2, 66, 67, 3, 2, 2, 2, 67, 9, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 71, 7,
	4, 2, 2, 70, 72, 5, 6, 4, 2, 71, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73,
	71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 77, 7, 4, 2,
	2, 76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 11, 3, 2, 2, 2, 78, 79,
	7, 19, 2, 2, 79, 81, 7, 19, 2, 2, 80, 82, 5, 6, 4, 2, 81, 80, 3, 2, 2,
	2, 82, 83, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85,
	3, 2, 2, 2, 85, 86, 7, 19, 2, 2, 86, 87, 7, 19, 2, 2, 87, 13, 3, 2, 2,
	2, 88, 89, 7, 11, 2, 2, 89, 96, 5, 8, 5, 2, 90, 92, 7, 5, 2, 2, 91, 93,
	5, 6, 4, 2, 92, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2,
	94, 95, 3, 2, 2, 2, 95, 97, 3, 2, 2, 2, 96, 90, 3, 2, 2, 2, 96, 97, 3,
	2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 7, 12, 2, 2, 99, 111, 3, 2, 2, 2,
	100, 101, 7, 13, 2, 2, 101, 102, 5, 8, 5, 2, 102, 104, 7, 5, 2, 2, 103,
	105, 5, 6, 4, 2, 104, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 104,
	3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 7, 14,
	2, 2, 109, 111, 3, 2, 2, 2, 110, 88, 3, 2, 2, 2, 110, 100, 3, 2, 2, 2,
	111, 15, 3, 2, 2, 2, 112, 113, 7, 13, 2, 2, 113, 114, 5, 8, 5, 2, 114,
	115, 7, 14, 2, 2, 115, 17, 3, 2, 2, 2, 116, 117, 7, 6, 2, 2, 117, 19, 3,
	2, 2, 2, 118, 120, 7, 7, 2, 2, 119, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2,
	2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	131, 5, 6, 4, 2, 124, 126, 7, 10, 2, 2, 125, 124, 3, 2, 2, 2, 126, 127,
	3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 3, 2,
	2, 2, 129, 131, 5, 6, 4, 2, 130, 119, 3, 2, 2, 2, 130, 125, 3, 2, 2, 2,
	131, 21, 3, 2, 2, 2, 132, 134, 7, 8, 2, 2, 133, 135, 5, 6, 4, 2, 134, 133,
	3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2,
	2, 2, 137, 139, 3, 2, 2, 2, 138, 132, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2,
	140, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142,
	146, 7, 5, 2, 2, 143, 145, 7, 16, 2, 2, 144, 143, 3, 2, 2, 2, 145, 148,
	3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 23, 3, 2,
	2, 2, 148, 146, 3, 2, 2, 2, 149, 151, 7, 5, 2, 2, 150, 152, 5, 6, 4, 2,
	151, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153,
	154, 3, 2, 2, 2, 154, 156, 3, 2, 2, 2, 155, 149, 3, 2, 2, 2, 156, 157,
	3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 3, 2,
	2, 2, 159, 163, 7, 5, 2, 2, 160, 162, 7, 16, 2, 2, 161, 160, 3, 2, 2, 2,
	162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164,
	25, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166, 168, 7, 9, 2, 2, 167, 166, 3,
	2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2,
	2, 170, 171, 3, 2, 2, 2, 171, 175, 5, 6, 4, 2, 172, 174, 7, 9, 2, 2, 173,
	172, 3, 2, 2, 2, 174, 177, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176,
	3, 2, 2, 2, 176, 27, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 178, 179, 7, 18,
	2, 2, 179, 29, 3, 2, 2, 2, 26, 31, 36, 42, 55, 60, 66, 73, 76, 83, 94,
	96, 106, 110, 121, 127, 130, 136, 140, 146, 153, 157, 163, 169, 175,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\\\\'", "'**'", "'|'", "'----'", "'*'", "'|='", "'='", "'#'", "'[['",
	"']]'", "'{{'", "'}}'", "", "", "", "", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "HASH", "LBRACKET", "RBRACKET", "LBRACE",
	"RBRACE", "TEXT", "WS", "CR", "NOWIKI", "RSLASH",
}

var ruleNames = []string{
	"document", "line", "markup", "text", "bold", "italics", "href", "image",
	"hline", "listitem", "tableheader", "tablerow", "title", "nowiki",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type creoleParser struct {
	*antlr.BaseParser
}

func NewcreoleParser(input antlr.TokenStream) *creoleParser {
	this := new(creoleParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "creole.g4"

	return this
}

// creoleParser tokens.
const (
	creoleParserEOF      = antlr.TokenEOF
	creoleParserT__0     = 1
	creoleParserT__1     = 2
	creoleParserT__2     = 3
	creoleParserT__3     = 4
	creoleParserT__4     = 5
	creoleParserT__5     = 6
	creoleParserT__6     = 7
	creoleParserHASH     = 8
	creoleParserLBRACKET = 9
	creoleParserRBRACKET = 10
	creoleParserLBRACE   = 11
	creoleParserRBRACE   = 12
	creoleParserTEXT     = 13
	creoleParserWS       = 14
	creoleParserCR       = 15
	creoleParserNOWIKI   = 16
	creoleParserRSLASH   = 17
)

// creoleParser rules.
const (
	creoleParserRULE_document    = 0
	creoleParserRULE_line        = 1
	creoleParserRULE_markup      = 2
	creoleParserRULE_text        = 3
	creoleParserRULE_bold        = 4
	creoleParserRULE_italics     = 5
	creoleParserRULE_href        = 6
	creoleParserRULE_image       = 7
	creoleParserRULE_hline       = 8
	creoleParserRULE_listitem    = 9
	creoleParserRULE_tableheader = 10
	creoleParserRULE_tablerow    = 11
	creoleParserRULE_title       = 12
	creoleParserRULE_nowiki      = 13
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllCR() []antlr.TerminalNode {
	return s.GetTokens(creoleParserCR)
}

func (s *DocumentContext) CR(i int) antlr.TerminalNode {
	return s.GetToken(creoleParserCR, i)
}

func (s *DocumentContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *DocumentContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *creoleParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, creoleParserRULE_document)
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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<creoleParserT__1)|(1<<creoleParserT__2)|(1<<creoleParserT__3)|(1<<creoleParserT__4)|(1<<creoleParserT__5)|(1<<creoleParserT__6)|(1<<creoleParserHASH)|(1<<creoleParserLBRACKET)|(1<<creoleParserLBRACE)|(1<<creoleParserTEXT)|(1<<creoleParserCR)|(1<<creoleParserNOWIKI)|(1<<creoleParserRSLASH))) != 0 {
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<creoleParserT__1)|(1<<creoleParserT__2)|(1<<creoleParserT__3)|(1<<creoleParserT__4)|(1<<creoleParserT__5)|(1<<creoleParserT__6)|(1<<creoleParserHASH)|(1<<creoleParserLBRACKET)|(1<<creoleParserLBRACE)|(1<<creoleParserTEXT)|(1<<creoleParserNOWIKI)|(1<<creoleParserRSLASH))) != 0 {
			{
				p.SetState(28)
				p.Line()
			}

		}
		{
			p.SetState(31)
			p.Match(creoleParserCR)
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) AllMarkup() []IMarkupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMarkupContext)(nil)).Elem())
	var tst = make([]IMarkupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMarkupContext)
		}
	}

	return tst
}

func (s *LineContext) Markup(i int) IMarkupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarkupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMarkupContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *creoleParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, creoleParserRULE_line)
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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<creoleParserT__1)|(1<<creoleParserT__2)|(1<<creoleParserT__3)|(1<<creoleParserT__4)|(1<<creoleParserT__5)|(1<<creoleParserT__6)|(1<<creoleParserHASH)|(1<<creoleParserLBRACKET)|(1<<creoleParserLBRACE)|(1<<creoleParserTEXT)|(1<<creoleParserNOWIKI)|(1<<creoleParserRSLASH))) != 0) {
		{
			p.SetState(37)
			p.Markup()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMarkupContext is an interface to support dynamic dispatch.
type IMarkupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMarkupContext differentiates from other interfaces.
	IsMarkupContext()
}

type MarkupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMarkupContext() *MarkupContext {
	var p = new(MarkupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_markup
	return p
}

func (*MarkupContext) IsMarkupContext() {}

func NewMarkupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MarkupContext {
	var p = new(MarkupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_markup

	return p
}

func (s *MarkupContext) GetParser() antlr.Parser { return s.parser }

func (s *MarkupContext) Bold() IBoldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoldContext)
}

func (s *MarkupContext) Italics() IItalicsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItalicsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItalicsContext)
}

func (s *MarkupContext) Href() IHrefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHrefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHrefContext)
}

func (s *MarkupContext) Title() ITitleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITitleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITitleContext)
}

func (s *MarkupContext) Hline() IHlineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHlineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHlineContext)
}

func (s *MarkupContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *MarkupContext) Listitem() IListitemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListitemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListitemContext)
}

func (s *MarkupContext) Image() IImageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImageContext)
}

func (s *MarkupContext) Tablerow() ITablerowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITablerowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITablerowContext)
}

func (s *MarkupContext) Tableheader() ITableheaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableheaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableheaderContext)
}

func (s *MarkupContext) Nowiki() INowikiContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INowikiContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INowikiContext)
}

func (s *MarkupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarkupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarkupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterMarkup(s)
	}
}

func (s *MarkupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitMarkup(s)
	}
}

func (p *creoleParser) Markup() (localctx IMarkupContext) {
	localctx = NewMarkupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, creoleParserRULE_markup)

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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Bold()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Italics()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Href()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.Title()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(46)
			p.Hline()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(47)
			p.Text()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(48)
			p.Listitem()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(49)
			p.Image()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(50)
			p.Tablerow()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(51)
			p.Tableheader()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(52)
			p.Nowiki()
		}

	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) AllText() []ITextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITextContext)(nil)).Elem())
	var tst = make([]ITextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITextContext)
		}
	}

	return tst
}

func (s *TextContext) Text(i int) ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *TextContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(creoleParserTEXT)
}

func (s *TextContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(creoleParserTEXT, i)
}

func (s *TextContext) AllRSLASH() []antlr.TerminalNode {
	return s.GetTokens(creoleParserRSLASH)
}

func (s *TextContext) RSLASH(i int) antlr.TerminalNode {
	return s.GetToken(creoleParserRSLASH, i)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitText(s)
	}
}

func (p *creoleParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, creoleParserRULE_text)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(55)
			_la = p.GetTokenStream().LA(1)

			if !(_la == creoleParserTEXT || _la == creoleParserRSLASH) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(60)
				p.Match(creoleParserT__0)
			}
			{
				p.SetState(61)
				p.Text()
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IBoldContext is an interface to support dynamic dispatch.
type IBoldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoldContext differentiates from other interfaces.
	IsBoldContext()
}

type BoldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoldContext() *BoldContext {
	var p = new(BoldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_bold
	return p
}

func (*BoldContext) IsBoldContext() {}

func NewBoldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoldContext {
	var p = new(BoldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_bold

	return p
}

func (s *BoldContext) GetParser() antlr.Parser { return s.parser }

func (s *BoldContext) AllMarkup() []IMarkupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMarkupContext)(nil)).Elem())
	var tst = make([]IMarkupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMarkupContext)
		}
	}

	return tst
}

func (s *BoldContext) Markup(i int) IMarkupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarkupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMarkupContext)
}

func (s *BoldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterBold(s)
	}
}

func (s *BoldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitBold(s)
	}
}

func (p *creoleParser) Bold() (localctx IBoldContext) {
	localctx = NewBoldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, creoleParserRULE_bold)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(creoleParserT__1)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(68)
				p.Markup()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(73)
			p.Match(creoleParserT__1)
		}

	}

	return localctx
}

// IItalicsContext is an interface to support dynamic dispatch.
type IItalicsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItalicsContext differentiates from other interfaces.
	IsItalicsContext()
}

type ItalicsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItalicsContext() *ItalicsContext {
	var p = new(ItalicsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_italics
	return p
}

func (*ItalicsContext) IsItalicsContext() {}

func NewItalicsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItalicsContext {
	var p = new(ItalicsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_italics

	return p
}

func (s *ItalicsContext) GetParser() antlr.Parser { return s.parser }

func (s *ItalicsContext) AllRSLASH() []antlr.TerminalNode {
	return s.GetTokens(creoleParserRSLASH)
}

func (s *ItalicsContext) RSLASH(i int) antlr.TerminalNode {
	return s.GetToken(creoleParserRSLASH, i)
}

func (s *ItalicsContext) AllMarkup() []IMarkupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMarkupContext)(nil)).Elem())
	var tst = make([]IMarkupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMarkupContext)
		}
	}

	return tst
}

func (s *ItalicsContext) Markup(i int) IMarkupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarkupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMarkupContext)
}

func (s *ItalicsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItalicsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItalicsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterItalics(s)
	}
}

func (s *ItalicsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitItalics(s)
	}
}

func (p *creoleParser) Italics() (localctx IItalicsContext) {
	localctx = NewItalicsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, creoleParserRULE_italics)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(creoleParserRSLASH)
	}
	{
		p.SetState(77)
		p.Match(creoleParserRSLASH)
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(78)
				p.Markup()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	{
		p.SetState(83)
		p.Match(creoleParserRSLASH)
	}
	{
		p.SetState(84)
		p.Match(creoleParserRSLASH)
	}

	return localctx
}

// IHrefContext is an interface to support dynamic dispatch.
type IHrefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHrefContext differentiates from other interfaces.
	IsHrefContext()
}

type HrefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHrefContext() *HrefContext {
	var p = new(HrefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_href
	return p
}

func (*HrefContext) IsHrefContext() {}

func NewHrefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HrefContext {
	var p = new(HrefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_href

	return p
}

func (s *HrefContext) GetParser() antlr.Parser { return s.parser }

func (s *HrefContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(creoleParserLBRACKET, 0)
}

func (s *HrefContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *HrefContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(creoleParserRBRACKET, 0)
}

func (s *HrefContext) AllMarkup() []IMarkupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMarkupContext)(nil)).Elem())
	var tst = make([]IMarkupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMarkupContext)
		}
	}

	return tst
}

func (s *HrefContext) Markup(i int) IMarkupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarkupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMarkupContext)
}

func (s *HrefContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(creoleParserLBRACE, 0)
}

func (s *HrefContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(creoleParserRBRACE, 0)
}

func (s *HrefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HrefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HrefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterHref(s)
	}
}

func (s *HrefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitHref(s)
	}
}

func (p *creoleParser) Href() (localctx IHrefContext) {
	localctx = NewHrefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, creoleParserRULE_href)
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

	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case creoleParserLBRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(creoleParserLBRACKET)
		}
		{
			p.SetState(87)
			p.Text()
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == creoleParserT__2 {
			{
				p.SetState(88)
				p.Match(creoleParserT__2)
			}
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<creoleParserT__1)|(1<<creoleParserT__2)|(1<<creoleParserT__3)|(1<<creoleParserT__4)|(1<<creoleParserT__5)|(1<<creoleParserT__6)|(1<<creoleParserHASH)|(1<<creoleParserLBRACKET)|(1<<creoleParserLBRACE)|(1<<creoleParserTEXT)|(1<<creoleParserNOWIKI)|(1<<creoleParserRSLASH))) != 0) {
				{
					p.SetState(89)
					p.Markup()
				}

				p.SetState(92)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(96)
			p.Match(creoleParserRBRACKET)
		}

	case creoleParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Match(creoleParserLBRACE)
		}
		{
			p.SetState(99)
			p.Text()
		}
		{
			p.SetState(100)
			p.Match(creoleParserT__2)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<creoleParserT__1)|(1<<creoleParserT__2)|(1<<creoleParserT__3)|(1<<creoleParserT__4)|(1<<creoleParserT__5)|(1<<creoleParserT__6)|(1<<creoleParserHASH)|(1<<creoleParserLBRACKET)|(1<<creoleParserLBRACE)|(1<<creoleParserTEXT)|(1<<creoleParserNOWIKI)|(1<<creoleParserRSLASH))) != 0) {
			{
				p.SetState(101)
				p.Markup()
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(106)
			p.Match(creoleParserRBRACE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImageContext is an interface to support dynamic dispatch.
type IImageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImageContext differentiates from other interfaces.
	IsImageContext()
}

type ImageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImageContext() *ImageContext {
	var p = new(ImageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_image
	return p
}

func (*ImageContext) IsImageContext() {}

func NewImageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImageContext {
	var p = new(ImageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_image

	return p
}

func (s *ImageContext) GetParser() antlr.Parser { return s.parser }

func (s *ImageContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(creoleParserLBRACE, 0)
}

func (s *ImageContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *ImageContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(creoleParserRBRACE, 0)
}

func (s *ImageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterImage(s)
	}
}

func (s *ImageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitImage(s)
	}
}

func (p *creoleParser) Image() (localctx IImageContext) {
	localctx = NewImageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, creoleParserRULE_image)

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
		p.SetState(110)
		p.Match(creoleParserLBRACE)
	}
	{
		p.SetState(111)
		p.Text()
	}
	{
		p.SetState(112)
		p.Match(creoleParserRBRACE)
	}

	return localctx
}

// IHlineContext is an interface to support dynamic dispatch.
type IHlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHlineContext differentiates from other interfaces.
	IsHlineContext()
}

type HlineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHlineContext() *HlineContext {
	var p = new(HlineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_hline
	return p
}

func (*HlineContext) IsHlineContext() {}

func NewHlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HlineContext {
	var p = new(HlineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_hline

	return p
}

func (s *HlineContext) GetParser() antlr.Parser { return s.parser }
func (s *HlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterHline(s)
	}
}

func (s *HlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitHline(s)
	}
}

func (p *creoleParser) Hline() (localctx IHlineContext) {
	localctx = NewHlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, creoleParserRULE_hline)

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
		p.SetState(114)
		p.Match(creoleParserT__3)
	}

	return localctx
}

// IListitemContext is an interface to support dynamic dispatch.
type IListitemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListitemContext differentiates from other interfaces.
	IsListitemContext()
}

type ListitemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListitemContext() *ListitemContext {
	var p = new(ListitemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_listitem
	return p
}

func (*ListitemContext) IsListitemContext() {}

func NewListitemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListitemContext {
	var p = new(ListitemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_listitem

	return p
}

func (s *ListitemContext) GetParser() antlr.Parser { return s.parser }

func (s *ListitemContext) Markup() IMarkupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarkupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMarkupContext)
}

func (s *ListitemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListitemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListitemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterListitem(s)
	}
}

func (s *ListitemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitListitem(s)
	}
}

func (p *creoleParser) Listitem() (localctx IListitemContext) {
	localctx = NewListitemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, creoleParserRULE_listitem)

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

	var _alt int

	p.SetState(128)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case creoleParserT__4:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(116)
					p.Match(creoleParserT__4)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
		}
		{
			p.SetState(121)
			p.Markup()
		}

	case creoleParserHASH:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(122)
					p.Match(creoleParserHASH)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
		}
		{
			p.SetState(127)
			p.Markup()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITableheaderContext is an interface to support dynamic dispatch.
type ITableheaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableheaderContext differentiates from other interfaces.
	IsTableheaderContext()
}

type TableheaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableheaderContext() *TableheaderContext {
	var p = new(TableheaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_tableheader
	return p
}

func (*TableheaderContext) IsTableheaderContext() {}

func NewTableheaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableheaderContext {
	var p = new(TableheaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_tableheader

	return p
}

func (s *TableheaderContext) GetParser() antlr.Parser { return s.parser }

func (s *TableheaderContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(creoleParserWS)
}

func (s *TableheaderContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(creoleParserWS, i)
}

func (s *TableheaderContext) AllMarkup() []IMarkupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMarkupContext)(nil)).Elem())
	var tst = make([]IMarkupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMarkupContext)
		}
	}

	return tst
}

func (s *TableheaderContext) Markup(i int) IMarkupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarkupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMarkupContext)
}

func (s *TableheaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableheaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableheaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterTableheader(s)
	}
}

func (s *TableheaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitTableheader(s)
	}
}

func (p *creoleParser) Tableheader() (localctx ITableheaderContext) {
	localctx = NewTableheaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, creoleParserRULE_tableheader)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == creoleParserT__5 {
		{
			p.SetState(130)
			p.Match(creoleParserT__5)
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(131)
					p.Markup()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
		p.Match(creoleParserT__2)
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == creoleParserWS {
		{
			p.SetState(141)
			p.Match(creoleParserWS)
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITablerowContext is an interface to support dynamic dispatch.
type ITablerowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTablerowContext differentiates from other interfaces.
	IsTablerowContext()
}

type TablerowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTablerowContext() *TablerowContext {
	var p = new(TablerowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_tablerow
	return p
}

func (*TablerowContext) IsTablerowContext() {}

func NewTablerowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TablerowContext {
	var p = new(TablerowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_tablerow

	return p
}

func (s *TablerowContext) GetParser() antlr.Parser { return s.parser }

func (s *TablerowContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(creoleParserWS)
}

func (s *TablerowContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(creoleParserWS, i)
}

func (s *TablerowContext) AllMarkup() []IMarkupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMarkupContext)(nil)).Elem())
	var tst = make([]IMarkupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMarkupContext)
		}
	}

	return tst
}

func (s *TablerowContext) Markup(i int) IMarkupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarkupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMarkupContext)
}

func (s *TablerowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TablerowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TablerowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterTablerow(s)
	}
}

func (s *TablerowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitTablerow(s)
	}
}

func (p *creoleParser) Tablerow() (localctx ITablerowContext) {
	localctx = NewTablerowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, creoleParserRULE_tablerow)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(147)
				p.Match(creoleParserT__2)
			}
			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(148)
						p.Markup()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(151)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	{
		p.SetState(157)
		p.Match(creoleParserT__2)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == creoleParserWS {
		{
			p.SetState(158)
			p.Match(creoleParserWS)
		}

		p.SetState(163)
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
	p.RuleIndex = creoleParserRULE_title
	return p
}

func (*TitleContext) IsTitleContext() {}

func NewTitleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TitleContext {
	var p = new(TitleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_title

	return p
}

func (s *TitleContext) GetParser() antlr.Parser { return s.parser }

func (s *TitleContext) Markup() IMarkupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarkupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMarkupContext)
}

func (s *TitleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TitleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TitleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterTitle(s)
	}
}

func (s *TitleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitTitle(s)
	}
}

func (p *creoleParser) Title() (localctx ITitleContext) {
	localctx = NewTitleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, creoleParserRULE_title)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(164)
				p.Match(creoleParserT__6)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}
	{
		p.SetState(169)
		p.Markup()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(170)
				p.Match(creoleParserT__6)
			}

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// INowikiContext is an interface to support dynamic dispatch.
type INowikiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNowikiContext differentiates from other interfaces.
	IsNowikiContext()
}

type NowikiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNowikiContext() *NowikiContext {
	var p = new(NowikiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = creoleParserRULE_nowiki
	return p
}

func (*NowikiContext) IsNowikiContext() {}

func NewNowikiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NowikiContext {
	var p = new(NowikiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = creoleParserRULE_nowiki

	return p
}

func (s *NowikiContext) GetParser() antlr.Parser { return s.parser }

func (s *NowikiContext) NOWIKI() antlr.TerminalNode {
	return s.GetToken(creoleParserNOWIKI, 0)
}

func (s *NowikiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NowikiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NowikiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.EnterNowiki(s)
	}
}

func (s *NowikiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(creoleListener); ok {
		listenerT.ExitNowiki(s)
	}
}

func (p *creoleParser) Nowiki() (localctx INowikiContext) {
	localctx = NewNowikiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, creoleParserRULE_nowiki)

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
		p.SetState(176)
		p.Match(creoleParserNOWIKI)
	}

	return localctx
}
