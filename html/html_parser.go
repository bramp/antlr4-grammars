// Code generated from HTMLParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package html // HTMLParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 130,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	7, 2, 26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 2, 5, 2, 32, 10, 2, 3, 2,
	7, 2, 35, 10, 2, 12, 2, 14, 2, 38, 11, 2, 3, 2, 5, 2, 41, 10, 2, 3, 2,
	7, 2, 44, 10, 2, 12, 2, 14, 2, 47, 11, 2, 3, 2, 7, 2, 50, 10, 2, 12, 2,
	14, 2, 53, 11, 2, 3, 3, 3, 3, 3, 4, 7, 4, 58, 10, 4, 12, 4, 14, 4, 61,
	11, 4, 3, 4, 3, 4, 7, 4, 65, 10, 4, 12, 4, 14, 4, 68, 11, 4, 3, 5, 3, 5,
	3, 5, 7, 5, 73, 10, 5, 12, 5, 14, 5, 76, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 85, 10, 5, 3, 5, 5, 5, 88, 10, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 93, 10, 5, 3, 6, 5, 6, 96, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 101,
	10, 6, 3, 6, 5, 6, 104, 10, 6, 7, 6, 106, 10, 6, 12, 6, 14, 6, 109, 11,
	6, 3, 7, 3, 7, 3, 7, 5, 7, 114, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 120,
	10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 2,
	2, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 7, 3, 2, 8, 9, 4, 2,
	9, 9, 13, 13, 3, 2, 3, 4, 3, 2, 20, 21, 3, 2, 22, 23, 2, 139, 2, 27, 3,
	2, 2, 2, 4, 54, 3, 2, 2, 2, 6, 59, 3, 2, 2, 2, 8, 92, 3, 2, 2, 2, 10, 95,
	3, 2, 2, 2, 12, 110, 3, 2, 2, 2, 14, 115, 3, 2, 2, 2, 16, 119, 3, 2, 2,
	2, 18, 121, 3, 2, 2, 2, 20, 123, 3, 2, 2, 2, 22, 126, 3, 2, 2, 2, 24, 26,
	5, 4, 3, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2,
	27, 28, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 32, 7,
	5, 2, 2, 31, 30, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 36, 3, 2, 2, 2, 33,
	35, 5, 4, 3, 2, 34, 33, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3, 2, 2,
	2, 36, 37, 3, 2, 2, 2, 37, 40, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39, 41,
	7, 7, 2, 2, 40, 39, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 45, 3, 2, 2, 2,
	42, 44, 5, 4, 3, 2, 43, 42, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3,
	2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 51, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48,
	50, 5, 6, 4, 2, 49, 48, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2,
	2, 51, 52, 3, 2, 2, 2, 52, 3, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 55, 9,
	2, 2, 2, 55, 5, 3, 2, 2, 2, 56, 58, 5, 16, 9, 2, 57, 56, 3, 2, 2, 2, 58,
	61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 3, 2, 2,
	2, 61, 59, 3, 2, 2, 2, 62, 66, 5, 8, 5, 2, 63, 65, 5, 16, 9, 2, 64, 63,
	3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2,
	67, 7, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 70, 7, 12, 2, 2, 70, 74, 7,
	18, 2, 2, 71, 73, 5, 12, 7, 2, 72, 71, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2,
	74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 87, 3, 2, 2, 2, 76, 74, 3,
	2, 2, 2, 77, 84, 7, 14, 2, 2, 78, 79, 5, 10, 6, 2, 79, 80, 7, 12, 2, 2,
	80, 81, 7, 16, 2, 2, 81, 82, 7, 18, 2, 2, 82, 83, 7, 14, 2, 2, 83, 85,
	3, 2, 2, 2, 84, 78, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2,
	86, 88, 7, 15, 2, 2, 87, 77, 3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 93, 3,
	2, 2, 2, 89, 93, 7, 8, 2, 2, 90, 93, 5, 20, 11, 2, 91, 93, 5, 22, 12, 2,
	92, 69, 3, 2, 2, 2, 92, 89, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 91, 3,
	2, 2, 2, 93, 9, 3, 2, 2, 2, 94, 96, 5, 14, 8, 2, 95, 94, 3, 2, 2, 2, 95,
	96, 3, 2, 2, 2, 96, 107, 3, 2, 2, 2, 97, 101, 5, 8, 5, 2, 98, 101, 7, 6,
	2, 2, 99, 101, 5, 18, 10, 2, 100, 97, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2,
	100, 99, 3, 2, 2, 2, 101, 103, 3, 2, 2, 2, 102, 104, 5, 14, 8, 2, 103,
	102, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 106, 3, 2, 2, 2, 105, 100,
	3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2,
	2, 2, 108, 11, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 113, 7, 18, 2, 2,
	111, 112, 7, 17, 2, 2, 112, 114, 7, 24, 2, 2, 113, 111, 3, 2, 2, 2, 113,
	114, 3, 2, 2, 2, 114, 13, 3, 2, 2, 2, 115, 116, 9, 3, 2, 2, 116, 15, 3,
	2, 2, 2, 117, 120, 5, 18, 10, 2, 118, 120, 7, 9, 2, 2, 119, 117, 3, 2,
	2, 2, 119, 118, 3, 2, 2, 2, 120, 17, 3, 2, 2, 2, 121, 122, 9, 4, 2, 2,
	122, 19, 3, 2, 2, 2, 123, 124, 7, 10, 2, 2, 124, 125, 9, 5, 2, 2, 125,
	21, 3, 2, 2, 2, 126, 127, 7, 11, 2, 2, 127, 128, 9, 6, 2, 2, 128, 23, 3,
	2, 2, 2, 20, 27, 31, 36, 40, 45, 51, 59, 66, 74, 84, 87, 92, 95, 100, 103,
	107, 113, 119,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "'<'", "", "'>'", "'/>'", "'/'",
	"'='",
}
var symbolicNames = []string{
	"", "HTML_COMMENT", "HTML_CONDITIONAL_COMMENT", "XML", "CDATA", "DTD",
	"SCRIPTLET", "SEA_WS", "SCRIPT_OPEN", "STYLE_OPEN", "TAG_OPEN", "HTML_TEXT",
	"TAG_CLOSE", "TAG_SLASH_CLOSE", "TAG_SLASH", "TAG_EQUALS", "TAG_NAME",
	"TAG_WHITESPACE", "SCRIPT_BODY", "SCRIPT_SHORT_BODY", "STYLE_BODY", "STYLE_SHORT_BODY",
	"ATTVALUE_VALUE", "ATTRIBUTE",
}

var ruleNames = []string{
	"htmlDocument", "scriptletOrSeaWs", "htmlElements", "htmlElement", "htmlContent",
	"htmlAttribute", "htmlChardata", "htmlMisc", "htmlComment", "script", "style",
}

type HTMLParser struct {
	*antlr.BaseParser
}

// NewHTMLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *HTMLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewHTMLParser(input antlr.TokenStream) *HTMLParser {
	this := new(HTMLParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "HTMLParser.g4"

	return this
}

// HTMLParser tokens.
const (
	HTMLParserEOF                      = antlr.TokenEOF
	HTMLParserHTML_COMMENT             = 1
	HTMLParserHTML_CONDITIONAL_COMMENT = 2
	HTMLParserXML                      = 3
	HTMLParserCDATA                    = 4
	HTMLParserDTD                      = 5
	HTMLParserSCRIPTLET                = 6
	HTMLParserSEA_WS                   = 7
	HTMLParserSCRIPT_OPEN              = 8
	HTMLParserSTYLE_OPEN               = 9
	HTMLParserTAG_OPEN                 = 10
	HTMLParserHTML_TEXT                = 11
	HTMLParserTAG_CLOSE                = 12
	HTMLParserTAG_SLASH_CLOSE          = 13
	HTMLParserTAG_SLASH                = 14
	HTMLParserTAG_EQUALS               = 15
	HTMLParserTAG_NAME                 = 16
	HTMLParserTAG_WHITESPACE           = 17
	HTMLParserSCRIPT_BODY              = 18
	HTMLParserSCRIPT_SHORT_BODY        = 19
	HTMLParserSTYLE_BODY               = 20
	HTMLParserSTYLE_SHORT_BODY         = 21
	HTMLParserATTVALUE_VALUE           = 22
	HTMLParserATTRIBUTE                = 23
)

// HTMLParser rules.
const (
	HTMLParserRULE_htmlDocument     = 0
	HTMLParserRULE_scriptletOrSeaWs = 1
	HTMLParserRULE_htmlElements     = 2
	HTMLParserRULE_htmlElement      = 3
	HTMLParserRULE_htmlContent      = 4
	HTMLParserRULE_htmlAttribute    = 5
	HTMLParserRULE_htmlChardata     = 6
	HTMLParserRULE_htmlMisc         = 7
	HTMLParserRULE_htmlComment      = 8
	HTMLParserRULE_script           = 9
	HTMLParserRULE_style            = 10
)

// IHtmlDocumentContext is an interface to support dynamic dispatch.
type IHtmlDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlDocumentContext differentiates from other interfaces.
	IsHtmlDocumentContext()
}

type HtmlDocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlDocumentContext() *HtmlDocumentContext {
	var p = new(HtmlDocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlDocument
	return p
}

func (*HtmlDocumentContext) IsHtmlDocumentContext() {}

func NewHtmlDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlDocumentContext {
	var p = new(HtmlDocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlDocument

	return p
}

func (s *HtmlDocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlDocumentContext) AllScriptletOrSeaWs() []IScriptletOrSeaWsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScriptletOrSeaWsContext)(nil)).Elem())
	var tst = make([]IScriptletOrSeaWsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScriptletOrSeaWsContext)
		}
	}

	return tst
}

func (s *HtmlDocumentContext) ScriptletOrSeaWs(i int) IScriptletOrSeaWsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScriptletOrSeaWsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScriptletOrSeaWsContext)
}

func (s *HtmlDocumentContext) XML() antlr.TerminalNode {
	return s.GetToken(HTMLParserXML, 0)
}

func (s *HtmlDocumentContext) DTD() antlr.TerminalNode {
	return s.GetToken(HTMLParserDTD, 0)
}

func (s *HtmlDocumentContext) AllHtmlElements() []IHtmlElementsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHtmlElementsContext)(nil)).Elem())
	var tst = make([]IHtmlElementsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHtmlElementsContext)
		}
	}

	return tst
}

func (s *HtmlDocumentContext) HtmlElements(i int) IHtmlElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlElementsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHtmlElementsContext)
}

func (s *HtmlDocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlDocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlDocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlDocument(s)
	}
}

func (s *HtmlDocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlDocument(s)
	}
}

func (p *HTMLParser) HtmlDocument() (localctx IHtmlDocumentContext) {
	this := p
	_ = this

	localctx = NewHtmlDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HTMLParserRULE_htmlDocument)
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(22)
				p.ScriptletOrSeaWs()
			}

		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == HTMLParserXML {
		{
			p.SetState(28)
			p.Match(HTMLParserXML)
		}

	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(31)
				p.ScriptletOrSeaWs()
			}

		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == HTMLParserDTD {
		{
			p.SetState(37)
			p.Match(HTMLParserDTD)
		}

	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(40)
				p.ScriptletOrSeaWs()
			}

		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HTMLParserHTML_COMMENT)|(1<<HTMLParserHTML_CONDITIONAL_COMMENT)|(1<<HTMLParserSCRIPTLET)|(1<<HTMLParserSEA_WS)|(1<<HTMLParserSCRIPT_OPEN)|(1<<HTMLParserSTYLE_OPEN)|(1<<HTMLParserTAG_OPEN))) != 0 {
		{
			p.SetState(46)
			p.HtmlElements()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IScriptletOrSeaWsContext is an interface to support dynamic dispatch.
type IScriptletOrSeaWsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptletOrSeaWsContext differentiates from other interfaces.
	IsScriptletOrSeaWsContext()
}

type ScriptletOrSeaWsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptletOrSeaWsContext() *ScriptletOrSeaWsContext {
	var p = new(ScriptletOrSeaWsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_scriptletOrSeaWs
	return p
}

func (*ScriptletOrSeaWsContext) IsScriptletOrSeaWsContext() {}

func NewScriptletOrSeaWsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptletOrSeaWsContext {
	var p = new(ScriptletOrSeaWsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_scriptletOrSeaWs

	return p
}

func (s *ScriptletOrSeaWsContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptletOrSeaWsContext) SCRIPTLET() antlr.TerminalNode {
	return s.GetToken(HTMLParserSCRIPTLET, 0)
}

func (s *ScriptletOrSeaWsContext) SEA_WS() antlr.TerminalNode {
	return s.GetToken(HTMLParserSEA_WS, 0)
}

func (s *ScriptletOrSeaWsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptletOrSeaWsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptletOrSeaWsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterScriptletOrSeaWs(s)
	}
}

func (s *ScriptletOrSeaWsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitScriptletOrSeaWs(s)
	}
}

func (p *HTMLParser) ScriptletOrSeaWs() (localctx IScriptletOrSeaWsContext) {
	this := p
	_ = this

	localctx = NewScriptletOrSeaWsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HTMLParserRULE_scriptletOrSeaWs)
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
		p.SetState(52)
		_la = p.GetTokenStream().LA(1)

		if !(_la == HTMLParserSCRIPTLET || _la == HTMLParserSEA_WS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IHtmlElementsContext is an interface to support dynamic dispatch.
type IHtmlElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlElementsContext differentiates from other interfaces.
	IsHtmlElementsContext()
}

type HtmlElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlElementsContext() *HtmlElementsContext {
	var p = new(HtmlElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlElements
	return p
}

func (*HtmlElementsContext) IsHtmlElementsContext() {}

func NewHtmlElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlElementsContext {
	var p = new(HtmlElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlElements

	return p
}

func (s *HtmlElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlElementsContext) HtmlElement() IHtmlElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHtmlElementContext)
}

func (s *HtmlElementsContext) AllHtmlMisc() []IHtmlMiscContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHtmlMiscContext)(nil)).Elem())
	var tst = make([]IHtmlMiscContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHtmlMiscContext)
		}
	}

	return tst
}

func (s *HtmlElementsContext) HtmlMisc(i int) IHtmlMiscContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlMiscContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHtmlMiscContext)
}

func (s *HtmlElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlElements(s)
	}
}

func (s *HtmlElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlElements(s)
	}
}

func (p *HTMLParser) HtmlElements() (localctx IHtmlElementsContext) {
	this := p
	_ = this

	localctx = NewHtmlElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, HTMLParserRULE_htmlElements)
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HTMLParserHTML_COMMENT)|(1<<HTMLParserHTML_CONDITIONAL_COMMENT)|(1<<HTMLParserSEA_WS))) != 0 {
		{
			p.SetState(54)
			p.HtmlMisc()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
		p.HtmlElement()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(61)
				p.HtmlMisc()
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IHtmlElementContext is an interface to support dynamic dispatch.
type IHtmlElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlElementContext differentiates from other interfaces.
	IsHtmlElementContext()
}

type HtmlElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlElementContext() *HtmlElementContext {
	var p = new(HtmlElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlElement
	return p
}

func (*HtmlElementContext) IsHtmlElementContext() {}

func NewHtmlElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlElementContext {
	var p = new(HtmlElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlElement

	return p
}

func (s *HtmlElementContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlElementContext) AllTAG_OPEN() []antlr.TerminalNode {
	return s.GetTokens(HTMLParserTAG_OPEN)
}

func (s *HtmlElementContext) TAG_OPEN(i int) antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_OPEN, i)
}

func (s *HtmlElementContext) AllTAG_NAME() []antlr.TerminalNode {
	return s.GetTokens(HTMLParserTAG_NAME)
}

func (s *HtmlElementContext) TAG_NAME(i int) antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_NAME, i)
}

func (s *HtmlElementContext) AllTAG_CLOSE() []antlr.TerminalNode {
	return s.GetTokens(HTMLParserTAG_CLOSE)
}

func (s *HtmlElementContext) TAG_CLOSE(i int) antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_CLOSE, i)
}

func (s *HtmlElementContext) TAG_SLASH_CLOSE() antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_SLASH_CLOSE, 0)
}

func (s *HtmlElementContext) AllHtmlAttribute() []IHtmlAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHtmlAttributeContext)(nil)).Elem())
	var tst = make([]IHtmlAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHtmlAttributeContext)
		}
	}

	return tst
}

func (s *HtmlElementContext) HtmlAttribute(i int) IHtmlAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHtmlAttributeContext)
}

func (s *HtmlElementContext) HtmlContent() IHtmlContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHtmlContentContext)
}

func (s *HtmlElementContext) TAG_SLASH() antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_SLASH, 0)
}

func (s *HtmlElementContext) SCRIPTLET() antlr.TerminalNode {
	return s.GetToken(HTMLParserSCRIPTLET, 0)
}

func (s *HtmlElementContext) Script() IScriptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScriptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScriptContext)
}

func (s *HtmlElementContext) Style() IStyleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStyleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStyleContext)
}

func (s *HtmlElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlElement(s)
	}
}

func (s *HtmlElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlElement(s)
	}
}

func (p *HTMLParser) HtmlElement() (localctx IHtmlElementContext) {
	this := p
	_ = this

	localctx = NewHtmlElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HTMLParserRULE_htmlElement)
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

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HTMLParserTAG_OPEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Match(HTMLParserTAG_OPEN)
		}
		{
			p.SetState(68)
			p.Match(HTMLParserTAG_NAME)
		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == HTMLParserTAG_NAME {
			{
				p.SetState(69)
				p.HtmlAttribute()
			}

			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case HTMLParserTAG_CLOSE:
			{
				p.SetState(75)
				p.Match(HTMLParserTAG_CLOSE)
			}
			p.SetState(82)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(76)
					p.HtmlContent()
				}
				{
					p.SetState(77)
					p.Match(HTMLParserTAG_OPEN)
				}
				{
					p.SetState(78)
					p.Match(HTMLParserTAG_SLASH)
				}
				{
					p.SetState(79)
					p.Match(HTMLParserTAG_NAME)
				}
				{
					p.SetState(80)
					p.Match(HTMLParserTAG_CLOSE)
				}

			}

		case HTMLParserTAG_SLASH_CLOSE:
			{
				p.SetState(84)
				p.Match(HTMLParserTAG_SLASH_CLOSE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case HTMLParserSCRIPTLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Match(HTMLParserSCRIPTLET)
		}

	case HTMLParserSCRIPT_OPEN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Script()
		}

	case HTMLParserSTYLE_OPEN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Style()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IHtmlContentContext is an interface to support dynamic dispatch.
type IHtmlContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlContentContext differentiates from other interfaces.
	IsHtmlContentContext()
}

type HtmlContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlContentContext() *HtmlContentContext {
	var p = new(HtmlContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlContent
	return p
}

func (*HtmlContentContext) IsHtmlContentContext() {}

func NewHtmlContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlContentContext {
	var p = new(HtmlContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlContent

	return p
}

func (s *HtmlContentContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlContentContext) AllHtmlChardata() []IHtmlChardataContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHtmlChardataContext)(nil)).Elem())
	var tst = make([]IHtmlChardataContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHtmlChardataContext)
		}
	}

	return tst
}

func (s *HtmlContentContext) HtmlChardata(i int) IHtmlChardataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlChardataContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHtmlChardataContext)
}

func (s *HtmlContentContext) AllHtmlElement() []IHtmlElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHtmlElementContext)(nil)).Elem())
	var tst = make([]IHtmlElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHtmlElementContext)
		}
	}

	return tst
}

func (s *HtmlContentContext) HtmlElement(i int) IHtmlElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHtmlElementContext)
}

func (s *HtmlContentContext) AllCDATA() []antlr.TerminalNode {
	return s.GetTokens(HTMLParserCDATA)
}

func (s *HtmlContentContext) CDATA(i int) antlr.TerminalNode {
	return s.GetToken(HTMLParserCDATA, i)
}

func (s *HtmlContentContext) AllHtmlComment() []IHtmlCommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHtmlCommentContext)(nil)).Elem())
	var tst = make([]IHtmlCommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHtmlCommentContext)
		}
	}

	return tst
}

func (s *HtmlContentContext) HtmlComment(i int) IHtmlCommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlCommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHtmlCommentContext)
}

func (s *HtmlContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlContent(s)
	}
}

func (s *HtmlContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlContent(s)
	}
}

func (p *HTMLParser) HtmlContent() (localctx IHtmlContentContext) {
	this := p
	_ = this

	localctx = NewHtmlContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, HTMLParserRULE_htmlContent)
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
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == HTMLParserSEA_WS || _la == HTMLParserHTML_TEXT {
		{
			p.SetState(92)
			p.HtmlChardata()
		}

	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(98)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case HTMLParserSCRIPTLET, HTMLParserSCRIPT_OPEN, HTMLParserSTYLE_OPEN, HTMLParserTAG_OPEN:
				{
					p.SetState(95)
					p.HtmlElement()
				}

			case HTMLParserCDATA:
				{
					p.SetState(96)
					p.Match(HTMLParserCDATA)
				}

			case HTMLParserHTML_COMMENT, HTMLParserHTML_CONDITIONAL_COMMENT:
				{
					p.SetState(97)
					p.HtmlComment()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == HTMLParserSEA_WS || _la == HTMLParserHTML_TEXT {
				{
					p.SetState(100)
					p.HtmlChardata()
				}

			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IHtmlAttributeContext is an interface to support dynamic dispatch.
type IHtmlAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlAttributeContext differentiates from other interfaces.
	IsHtmlAttributeContext()
}

type HtmlAttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlAttributeContext() *HtmlAttributeContext {
	var p = new(HtmlAttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlAttribute
	return p
}

func (*HtmlAttributeContext) IsHtmlAttributeContext() {}

func NewHtmlAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlAttributeContext {
	var p = new(HtmlAttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlAttribute

	return p
}

func (s *HtmlAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlAttributeContext) TAG_NAME() antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_NAME, 0)
}

func (s *HtmlAttributeContext) TAG_EQUALS() antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_EQUALS, 0)
}

func (s *HtmlAttributeContext) ATTVALUE_VALUE() antlr.TerminalNode {
	return s.GetToken(HTMLParserATTVALUE_VALUE, 0)
}

func (s *HtmlAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlAttribute(s)
	}
}

func (s *HtmlAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlAttribute(s)
	}
}

func (p *HTMLParser) HtmlAttribute() (localctx IHtmlAttributeContext) {
	this := p
	_ = this

	localctx = NewHtmlAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, HTMLParserRULE_htmlAttribute)
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
		p.SetState(108)
		p.Match(HTMLParserTAG_NAME)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == HTMLParserTAG_EQUALS {
		{
			p.SetState(109)
			p.Match(HTMLParserTAG_EQUALS)
		}
		{
			p.SetState(110)
			p.Match(HTMLParserATTVALUE_VALUE)
		}

	}

	return localctx
}

// IHtmlChardataContext is an interface to support dynamic dispatch.
type IHtmlChardataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlChardataContext differentiates from other interfaces.
	IsHtmlChardataContext()
}

type HtmlChardataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlChardataContext() *HtmlChardataContext {
	var p = new(HtmlChardataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlChardata
	return p
}

func (*HtmlChardataContext) IsHtmlChardataContext() {}

func NewHtmlChardataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlChardataContext {
	var p = new(HtmlChardataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlChardata

	return p
}

func (s *HtmlChardataContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlChardataContext) HTML_TEXT() antlr.TerminalNode {
	return s.GetToken(HTMLParserHTML_TEXT, 0)
}

func (s *HtmlChardataContext) SEA_WS() antlr.TerminalNode {
	return s.GetToken(HTMLParserSEA_WS, 0)
}

func (s *HtmlChardataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlChardataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlChardataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlChardata(s)
	}
}

func (s *HtmlChardataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlChardata(s)
	}
}

func (p *HTMLParser) HtmlChardata() (localctx IHtmlChardataContext) {
	this := p
	_ = this

	localctx = NewHtmlChardataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, HTMLParserRULE_htmlChardata)
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
		p.SetState(113)
		_la = p.GetTokenStream().LA(1)

		if !(_la == HTMLParserSEA_WS || _la == HTMLParserHTML_TEXT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IHtmlMiscContext is an interface to support dynamic dispatch.
type IHtmlMiscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlMiscContext differentiates from other interfaces.
	IsHtmlMiscContext()
}

type HtmlMiscContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlMiscContext() *HtmlMiscContext {
	var p = new(HtmlMiscContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlMisc
	return p
}

func (*HtmlMiscContext) IsHtmlMiscContext() {}

func NewHtmlMiscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlMiscContext {
	var p = new(HtmlMiscContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlMisc

	return p
}

func (s *HtmlMiscContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlMiscContext) HtmlComment() IHtmlCommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlCommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHtmlCommentContext)
}

func (s *HtmlMiscContext) SEA_WS() antlr.TerminalNode {
	return s.GetToken(HTMLParserSEA_WS, 0)
}

func (s *HtmlMiscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlMiscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlMiscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlMisc(s)
	}
}

func (s *HtmlMiscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlMisc(s)
	}
}

func (p *HTMLParser) HtmlMisc() (localctx IHtmlMiscContext) {
	this := p
	_ = this

	localctx = NewHtmlMiscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, HTMLParserRULE_htmlMisc)

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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HTMLParserHTML_COMMENT, HTMLParserHTML_CONDITIONAL_COMMENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.HtmlComment()
		}

	case HTMLParserSEA_WS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Match(HTMLParserSEA_WS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IHtmlCommentContext is an interface to support dynamic dispatch.
type IHtmlCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlCommentContext differentiates from other interfaces.
	IsHtmlCommentContext()
}

type HtmlCommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlCommentContext() *HtmlCommentContext {
	var p = new(HtmlCommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlComment
	return p
}

func (*HtmlCommentContext) IsHtmlCommentContext() {}

func NewHtmlCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlCommentContext {
	var p = new(HtmlCommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlComment

	return p
}

func (s *HtmlCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlCommentContext) HTML_COMMENT() antlr.TerminalNode {
	return s.GetToken(HTMLParserHTML_COMMENT, 0)
}

func (s *HtmlCommentContext) HTML_CONDITIONAL_COMMENT() antlr.TerminalNode {
	return s.GetToken(HTMLParserHTML_CONDITIONAL_COMMENT, 0)
}

func (s *HtmlCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlComment(s)
	}
}

func (s *HtmlCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlComment(s)
	}
}

func (p *HTMLParser) HtmlComment() (localctx IHtmlCommentContext) {
	this := p
	_ = this

	localctx = NewHtmlCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HTMLParserRULE_htmlComment)
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
		p.SetState(119)
		_la = p.GetTokenStream().LA(1)

		if !(_la == HTMLParserHTML_COMMENT || _la == HTMLParserHTML_CONDITIONAL_COMMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) SCRIPT_OPEN() antlr.TerminalNode {
	return s.GetToken(HTMLParserSCRIPT_OPEN, 0)
}

func (s *ScriptContext) SCRIPT_BODY() antlr.TerminalNode {
	return s.GetToken(HTMLParserSCRIPT_BODY, 0)
}

func (s *ScriptContext) SCRIPT_SHORT_BODY() antlr.TerminalNode {
	return s.GetToken(HTMLParserSCRIPT_SHORT_BODY, 0)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *HTMLParser) Script() (localctx IScriptContext) {
	this := p
	_ = this

	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, HTMLParserRULE_script)
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
		p.SetState(121)
		p.Match(HTMLParserSCRIPT_OPEN)
	}
	{
		p.SetState(122)
		_la = p.GetTokenStream().LA(1)

		if !(_la == HTMLParserSCRIPT_BODY || _la == HTMLParserSCRIPT_SHORT_BODY) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStyleContext is an interface to support dynamic dispatch.
type IStyleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStyleContext differentiates from other interfaces.
	IsStyleContext()
}

type StyleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStyleContext() *StyleContext {
	var p = new(StyleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_style
	return p
}

func (*StyleContext) IsStyleContext() {}

func NewStyleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StyleContext {
	var p = new(StyleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_style

	return p
}

func (s *StyleContext) GetParser() antlr.Parser { return s.parser }

func (s *StyleContext) STYLE_OPEN() antlr.TerminalNode {
	return s.GetToken(HTMLParserSTYLE_OPEN, 0)
}

func (s *StyleContext) STYLE_BODY() antlr.TerminalNode {
	return s.GetToken(HTMLParserSTYLE_BODY, 0)
}

func (s *StyleContext) STYLE_SHORT_BODY() antlr.TerminalNode {
	return s.GetToken(HTMLParserSTYLE_SHORT_BODY, 0)
}

func (s *StyleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StyleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StyleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterStyle(s)
	}
}

func (s *StyleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitStyle(s)
	}
}

func (p *HTMLParser) Style() (localctx IStyleContext) {
	this := p
	_ = this

	localctx = NewStyleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, HTMLParserRULE_style)
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
		p.SetState(124)
		p.Match(HTMLParserSTYLE_OPEN)
	}
	{
		p.SetState(125)
		_la = p.GetTokenStream().LA(1)

		if !(_la == HTMLParserSTYLE_BODY || _la == HTMLParserSTYLE_SHORT_BODY) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
