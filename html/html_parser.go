// Generated from HTMLParser.g4 by ANTLR 4.7.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 174,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 7, 2, 39, 10, 2, 12, 2, 14, 2, 42, 11, 2, 3, 2, 5, 2, 45,
	10, 2, 3, 2, 3, 2, 7, 2, 49, 10, 2, 12, 2, 14, 2, 52, 11, 2, 3, 2, 5, 2,
	55, 10, 2, 3, 2, 3, 2, 7, 2, 59, 10, 2, 12, 2, 14, 2, 62, 11, 2, 3, 2,
	7, 2, 65, 10, 2, 12, 2, 14, 2, 68, 11, 2, 3, 3, 7, 3, 71, 10, 3, 12, 3,
	14, 3, 74, 11, 3, 3, 3, 3, 3, 7, 3, 78, 10, 3, 12, 3, 14, 3, 81, 11, 3,
	3, 4, 3, 4, 3, 4, 7, 4, 86, 10, 4, 12, 4, 14, 4, 89, 11, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 101, 10, 4, 12, 4,
	14, 4, 104, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 111, 10, 4, 12,
	4, 14, 4, 114, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 121, 10, 4, 3,
	5, 5, 5, 124, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 129, 10, 5, 3, 5, 5, 5, 132,
	10, 5, 7, 5, 134, 10, 5, 12, 5, 14, 5, 137, 11, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 5, 6, 144, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 5, 11, 156, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 2, 2, 19, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 2, 6, 4, 2, 9, 9, 13, 13, 3, 2, 3, 4, 3, 2, 20, 21, 3, 2, 22,
	23, 2, 182, 2, 40, 3, 2, 2, 2, 4, 72, 3, 2, 2, 2, 6, 120, 3, 2, 2, 2, 8,
	123, 3, 2, 2, 2, 10, 143, 3, 2, 2, 2, 12, 145, 3, 2, 2, 2, 14, 147, 3,
	2, 2, 2, 16, 149, 3, 2, 2, 2, 18, 151, 3, 2, 2, 2, 20, 155, 3, 2, 2, 2,
	22, 157, 3, 2, 2, 2, 24, 159, 3, 2, 2, 2, 26, 161, 3, 2, 2, 2, 28, 163,
	3, 2, 2, 2, 30, 165, 3, 2, 2, 2, 32, 167, 3, 2, 2, 2, 34, 170, 3, 2, 2,
	2, 36, 39, 5, 30, 16, 2, 37, 39, 7, 9, 2, 2, 38, 36, 3, 2, 2, 2, 38, 37,
	3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2,
	41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 45, 5, 28, 15, 2, 44, 43, 3,
	2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 50, 3, 2, 2, 2, 46, 49, 5, 30, 16, 2,
	47, 49, 7, 9, 2, 2, 48, 46, 3, 2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 52, 3,
	2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52,
	50, 3, 2, 2, 2, 53, 55, 5, 26, 14, 2, 54, 53, 3, 2, 2, 2, 54, 55, 3, 2,
	2, 2, 55, 60, 3, 2, 2, 2, 56, 59, 5, 30, 16, 2, 57, 59, 7, 9, 2, 2, 58,
	56, 3, 2, 2, 2, 58, 57, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2,
	2, 60, 61, 3, 2, 2, 2, 61, 66, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 65,
	5, 4, 3, 2, 64, 63, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2,
	66, 67, 3, 2, 2, 2, 67, 3, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 71, 5, 20,
	11, 2, 70, 69, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72,
	73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 79, 5, 6, 4,
	2, 76, 78, 5, 20, 11, 2, 77, 76, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77,
	3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 5, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2,
	82, 83, 7, 12, 2, 2, 83, 87, 5, 16, 9, 2, 84, 86, 5, 10, 6, 2, 85, 84,
	3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2,
	88, 90, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90, 91, 7, 14, 2, 2, 91, 92, 5,
	8, 5, 2, 92, 93, 7, 12, 2, 2, 93, 94, 7, 16, 2, 2, 94, 95, 5, 16, 9, 2,
	95, 96, 7, 14, 2, 2, 96, 121, 3, 2, 2, 2, 97, 98, 7, 12, 2, 2, 98, 102,
	5, 16, 9, 2, 99, 101, 5, 10, 6, 2, 100, 99, 3, 2, 2, 2, 101, 104, 3, 2,
	2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 105, 3, 2, 2, 2,
	104, 102, 3, 2, 2, 2, 105, 106, 7, 15, 2, 2, 106, 121, 3, 2, 2, 2, 107,
	108, 7, 12, 2, 2, 108, 112, 5, 16, 9, 2, 109, 111, 5, 10, 6, 2, 110, 109,
	3, 2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2,
	2, 2, 113, 115, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 116, 7, 14, 2, 2,
	116, 121, 3, 2, 2, 2, 117, 121, 5, 30, 16, 2, 118, 121, 5, 32, 17, 2, 119,
	121, 5, 34, 18, 2, 120, 82, 3, 2, 2, 2, 120, 97, 3, 2, 2, 2, 120, 107,
	3, 2, 2, 2, 120, 117, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 119, 3, 2,
	2, 2, 121, 7, 3, 2, 2, 2, 122, 124, 5, 18, 10, 2, 123, 122, 3, 2, 2, 2,
	123, 124, 3, 2, 2, 2, 124, 135, 3, 2, 2, 2, 125, 129, 5, 6, 4, 2, 126,
	129, 5, 24, 13, 2, 127, 129, 5, 22, 12, 2, 128, 125, 3, 2, 2, 2, 128, 126,
	3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 132, 5, 18,
	10, 2, 131, 130, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 134, 3, 2, 2, 2,
	133, 128, 3, 2, 2, 2, 134, 137, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135,
	136, 3, 2, 2, 2, 136, 9, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 138, 139, 5,
	12, 7, 2, 139, 140, 7, 17, 2, 2, 140, 141, 5, 14, 8, 2, 141, 144, 3, 2,
	2, 2, 142, 144, 5, 12, 7, 2, 143, 138, 3, 2, 2, 2, 143, 142, 3, 2, 2, 2,
	144, 11, 3, 2, 2, 2, 145, 146, 7, 18, 2, 2, 146, 13, 3, 2, 2, 2, 147, 148,
	7, 24, 2, 2, 148, 15, 3, 2, 2, 2, 149, 150, 7, 18, 2, 2, 150, 17, 3, 2,
	2, 2, 151, 152, 9, 2, 2, 2, 152, 19, 3, 2, 2, 2, 153, 156, 5, 22, 12, 2,
	154, 156, 7, 9, 2, 2, 155, 153, 3, 2, 2, 2, 155, 154, 3, 2, 2, 2, 156,
	21, 3, 2, 2, 2, 157, 158, 9, 3, 2, 2, 158, 23, 3, 2, 2, 2, 159, 160, 7,
	6, 2, 2, 160, 25, 3, 2, 2, 2, 161, 162, 7, 7, 2, 2, 162, 27, 3, 2, 2, 2,
	163, 164, 7, 5, 2, 2, 164, 29, 3, 2, 2, 2, 165, 166, 7, 8, 2, 2, 166, 31,
	3, 2, 2, 2, 167, 168, 7, 10, 2, 2, 168, 169, 9, 4, 2, 2, 169, 33, 3, 2,
	2, 2, 170, 171, 7, 11, 2, 2, 171, 172, 9, 5, 2, 2, 172, 35, 3, 2, 2, 2,
	23, 38, 40, 44, 48, 50, 54, 58, 60, 66, 72, 79, 87, 102, 112, 120, 123,
	128, 131, 135, 143, 155,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "'<'", "", "'>'", "'/>'", "'/'",
	"'='",
}
var symbolicNames = []string{
	"", "HTML_COMMENT", "HTML_CONDITIONAL_COMMENT", "XML_DECLARATION", "CDATA",
	"DTD", "SCRIPTLET", "SEA_WS", "SCRIPT_OPEN", "STYLE_OPEN", "TAG_OPEN",
	"HTML_TEXT", "TAG_CLOSE", "TAG_SLASH_CLOSE", "TAG_SLASH", "TAG_EQUALS",
	"TAG_NAME", "TAG_WHITESPACE", "SCRIPT_BODY", "SCRIPT_SHORT_BODY", "STYLE_BODY",
	"STYLE_SHORT_BODY", "ATTVALUE_VALUE", "ATTRIBUTE",
}

var ruleNames = []string{
	"htmlDocument", "htmlElements", "htmlElement", "htmlContent", "htmlAttribute",
	"htmlAttributeName", "htmlAttributeValue", "htmlTagName", "htmlChardata",
	"htmlMisc", "htmlComment", "xhtmlCDATA", "dtd", "xml", "scriptlet", "script",
	"style",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type HTMLParser struct {
	*antlr.BaseParser
}

func NewHTMLParser(input antlr.TokenStream) *HTMLParser {
	this := new(HTMLParser)

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
	HTMLParserXML_DECLARATION          = 3
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
	HTMLParserRULE_htmlDocument       = 0
	HTMLParserRULE_htmlElements       = 1
	HTMLParserRULE_htmlElement        = 2
	HTMLParserRULE_htmlContent        = 3
	HTMLParserRULE_htmlAttribute      = 4
	HTMLParserRULE_htmlAttributeName  = 5
	HTMLParserRULE_htmlAttributeValue = 6
	HTMLParserRULE_htmlTagName        = 7
	HTMLParserRULE_htmlChardata       = 8
	HTMLParserRULE_htmlMisc           = 9
	HTMLParserRULE_htmlComment        = 10
	HTMLParserRULE_xhtmlCDATA         = 11
	HTMLParserRULE_dtd                = 12
	HTMLParserRULE_xml                = 13
	HTMLParserRULE_scriptlet          = 14
	HTMLParserRULE_script             = 15
	HTMLParserRULE_style              = 16
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

func (s *HtmlDocumentContext) AllScriptlet() []IScriptletContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScriptletContext)(nil)).Elem())
	var tst = make([]IScriptletContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScriptletContext)
		}
	}

	return tst
}

func (s *HtmlDocumentContext) Scriptlet(i int) IScriptletContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScriptletContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScriptletContext)
}

func (s *HtmlDocumentContext) AllSEA_WS() []antlr.TerminalNode {
	return s.GetTokens(HTMLParserSEA_WS)
}

func (s *HtmlDocumentContext) SEA_WS(i int) antlr.TerminalNode {
	return s.GetToken(HTMLParserSEA_WS, i)
}

func (s *HtmlDocumentContext) Xml() IXmlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXmlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXmlContext)
}

func (s *HtmlDocumentContext) Dtd() IDtdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDtdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDtdContext)
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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(36)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case HTMLParserSCRIPTLET:
				{
					p.SetState(34)
					p.Scriptlet()
				}

			case HTMLParserSEA_WS:
				{
					p.SetState(35)
					p.Match(HTMLParserSEA_WS)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == HTMLParserXML_DECLARATION {
		{
			p.SetState(41)
			p.Xml()
		}

	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(46)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case HTMLParserSCRIPTLET:
				{
					p.SetState(44)
					p.Scriptlet()
				}

			case HTMLParserSEA_WS:
				{
					p.SetState(45)
					p.Match(HTMLParserSEA_WS)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == HTMLParserDTD {
		{
			p.SetState(51)
			p.Dtd()
		}

	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(56)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case HTMLParserSCRIPTLET:
				{
					p.SetState(54)
					p.Scriptlet()
				}

			case HTMLParserSEA_WS:
				{
					p.SetState(55)
					p.Match(HTMLParserSEA_WS)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HTMLParserHTML_COMMENT)|(1<<HTMLParserHTML_CONDITIONAL_COMMENT)|(1<<HTMLParserSCRIPTLET)|(1<<HTMLParserSEA_WS)|(1<<HTMLParserSCRIPT_OPEN)|(1<<HTMLParserSTYLE_OPEN)|(1<<HTMLParserTAG_OPEN))) != 0 {
		{
			p.SetState(61)
			p.HtmlElements()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	localctx = NewHtmlElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HTMLParserRULE_htmlElements)
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
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HTMLParserHTML_COMMENT)|(1<<HTMLParserHTML_CONDITIONAL_COMMENT)|(1<<HTMLParserSEA_WS))) != 0 {
		{
			p.SetState(67)
			p.HtmlMisc()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(73)
		p.HtmlElement()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(74)
				p.HtmlMisc()
			}

		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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

func (s *HtmlElementContext) AllHtmlTagName() []IHtmlTagNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHtmlTagNameContext)(nil)).Elem())
	var tst = make([]IHtmlTagNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHtmlTagNameContext)
		}
	}

	return tst
}

func (s *HtmlElementContext) HtmlTagName(i int) IHtmlTagNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlTagNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHtmlTagNameContext)
}

func (s *HtmlElementContext) AllTAG_CLOSE() []antlr.TerminalNode {
	return s.GetTokens(HTMLParserTAG_CLOSE)
}

func (s *HtmlElementContext) TAG_CLOSE(i int) antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_CLOSE, i)
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

func (s *HtmlElementContext) TAG_SLASH_CLOSE() antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_SLASH_CLOSE, 0)
}

func (s *HtmlElementContext) Scriptlet() IScriptletContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScriptletContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScriptletContext)
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
	localctx = NewHtmlElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, HTMLParserRULE_htmlElement)
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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Match(HTMLParserTAG_OPEN)
		}
		{
			p.SetState(81)
			p.HtmlTagName()
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == HTMLParserTAG_NAME {
			{
				p.SetState(82)
				p.HtmlAttribute()
			}

			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(88)
			p.Match(HTMLParserTAG_CLOSE)
		}
		{
			p.SetState(89)
			p.HtmlContent()
		}
		{
			p.SetState(90)
			p.Match(HTMLParserTAG_OPEN)
		}
		{
			p.SetState(91)
			p.Match(HTMLParserTAG_SLASH)
		}
		{
			p.SetState(92)
			p.HtmlTagName()
		}
		{
			p.SetState(93)
			p.Match(HTMLParserTAG_CLOSE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Match(HTMLParserTAG_OPEN)
		}
		{
			p.SetState(96)
			p.HtmlTagName()
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == HTMLParserTAG_NAME {
			{
				p.SetState(97)
				p.HtmlAttribute()
			}

			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(103)
			p.Match(HTMLParserTAG_SLASH_CLOSE)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Match(HTMLParserTAG_OPEN)
		}
		{
			p.SetState(106)
			p.HtmlTagName()
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == HTMLParserTAG_NAME {
			{
				p.SetState(107)
				p.HtmlAttribute()
			}

			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(113)
			p.Match(HTMLParserTAG_CLOSE)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(115)
			p.Scriptlet()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(116)
			p.Script()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(117)
			p.Style()
		}

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

func (s *HtmlContentContext) AllXhtmlCDATA() []IXhtmlCDATAContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IXhtmlCDATAContext)(nil)).Elem())
	var tst = make([]IXhtmlCDATAContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IXhtmlCDATAContext)
		}
	}

	return tst
}

func (s *HtmlContentContext) XhtmlCDATA(i int) IXhtmlCDATAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXhtmlCDATAContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IXhtmlCDATAContext)
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
	localctx = NewHtmlContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HTMLParserRULE_htmlContent)
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
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == HTMLParserSEA_WS || _la == HTMLParserHTML_TEXT {
		{
			p.SetState(120)
			p.HtmlChardata()
		}

	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(126)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case HTMLParserSCRIPTLET, HTMLParserSCRIPT_OPEN, HTMLParserSTYLE_OPEN, HTMLParserTAG_OPEN:
				{
					p.SetState(123)
					p.HtmlElement()
				}

			case HTMLParserCDATA:
				{
					p.SetState(124)
					p.XhtmlCDATA()
				}

			case HTMLParserHTML_COMMENT, HTMLParserHTML_CONDITIONAL_COMMENT:
				{
					p.SetState(125)
					p.HtmlComment()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == HTMLParserSEA_WS || _la == HTMLParserHTML_TEXT {
				{
					p.SetState(128)
					p.HtmlChardata()
				}

			}

		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
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

func (s *HtmlAttributeContext) HtmlAttributeName() IHtmlAttributeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlAttributeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHtmlAttributeNameContext)
}

func (s *HtmlAttributeContext) TAG_EQUALS() antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_EQUALS, 0)
}

func (s *HtmlAttributeContext) HtmlAttributeValue() IHtmlAttributeValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHtmlAttributeValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHtmlAttributeValueContext)
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
	localctx = NewHtmlAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, HTMLParserRULE_htmlAttribute)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.HtmlAttributeName()
		}
		{
			p.SetState(137)
			p.Match(HTMLParserTAG_EQUALS)
		}
		{
			p.SetState(138)
			p.HtmlAttributeValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.HtmlAttributeName()
		}

	}

	return localctx
}

// IHtmlAttributeNameContext is an interface to support dynamic dispatch.
type IHtmlAttributeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlAttributeNameContext differentiates from other interfaces.
	IsHtmlAttributeNameContext()
}

type HtmlAttributeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlAttributeNameContext() *HtmlAttributeNameContext {
	var p = new(HtmlAttributeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlAttributeName
	return p
}

func (*HtmlAttributeNameContext) IsHtmlAttributeNameContext() {}

func NewHtmlAttributeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlAttributeNameContext {
	var p = new(HtmlAttributeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlAttributeName

	return p
}

func (s *HtmlAttributeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlAttributeNameContext) TAG_NAME() antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_NAME, 0)
}

func (s *HtmlAttributeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlAttributeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlAttributeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlAttributeName(s)
	}
}

func (s *HtmlAttributeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlAttributeName(s)
	}
}

func (p *HTMLParser) HtmlAttributeName() (localctx IHtmlAttributeNameContext) {
	localctx = NewHtmlAttributeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, HTMLParserRULE_htmlAttributeName)

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
		p.SetState(143)
		p.Match(HTMLParserTAG_NAME)
	}

	return localctx
}

// IHtmlAttributeValueContext is an interface to support dynamic dispatch.
type IHtmlAttributeValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlAttributeValueContext differentiates from other interfaces.
	IsHtmlAttributeValueContext()
}

type HtmlAttributeValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlAttributeValueContext() *HtmlAttributeValueContext {
	var p = new(HtmlAttributeValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlAttributeValue
	return p
}

func (*HtmlAttributeValueContext) IsHtmlAttributeValueContext() {}

func NewHtmlAttributeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlAttributeValueContext {
	var p = new(HtmlAttributeValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlAttributeValue

	return p
}

func (s *HtmlAttributeValueContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlAttributeValueContext) ATTVALUE_VALUE() antlr.TerminalNode {
	return s.GetToken(HTMLParserATTVALUE_VALUE, 0)
}

func (s *HtmlAttributeValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlAttributeValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlAttributeValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlAttributeValue(s)
	}
}

func (s *HtmlAttributeValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlAttributeValue(s)
	}
}

func (p *HTMLParser) HtmlAttributeValue() (localctx IHtmlAttributeValueContext) {
	localctx = NewHtmlAttributeValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, HTMLParserRULE_htmlAttributeValue)

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
		p.SetState(145)
		p.Match(HTMLParserATTVALUE_VALUE)
	}

	return localctx
}

// IHtmlTagNameContext is an interface to support dynamic dispatch.
type IHtmlTagNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHtmlTagNameContext differentiates from other interfaces.
	IsHtmlTagNameContext()
}

type HtmlTagNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHtmlTagNameContext() *HtmlTagNameContext {
	var p = new(HtmlTagNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_htmlTagName
	return p
}

func (*HtmlTagNameContext) IsHtmlTagNameContext() {}

func NewHtmlTagNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HtmlTagNameContext {
	var p = new(HtmlTagNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_htmlTagName

	return p
}

func (s *HtmlTagNameContext) GetParser() antlr.Parser { return s.parser }

func (s *HtmlTagNameContext) TAG_NAME() antlr.TerminalNode {
	return s.GetToken(HTMLParserTAG_NAME, 0)
}

func (s *HtmlTagNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HtmlTagNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HtmlTagNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterHtmlTagName(s)
	}
}

func (s *HtmlTagNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitHtmlTagName(s)
	}
}

func (p *HTMLParser) HtmlTagName() (localctx IHtmlTagNameContext) {
	localctx = NewHtmlTagNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, HTMLParserRULE_htmlTagName)

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
		p.SetState(147)
		p.Match(HTMLParserTAG_NAME)
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
	localctx = NewHtmlChardataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HTMLParserRULE_htmlChardata)
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
	p.SetState(149)
	_la = p.GetTokenStream().LA(1)

	if !(_la == HTMLParserSEA_WS || _la == HTMLParserHTML_TEXT) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
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
	localctx = NewHtmlMiscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, HTMLParserRULE_htmlMisc)

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

	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HTMLParserHTML_COMMENT, HTMLParserHTML_CONDITIONAL_COMMENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.HtmlComment()
		}

	case HTMLParserSEA_WS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
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
	localctx = NewHtmlCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, HTMLParserRULE_htmlComment)
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
	p.SetState(155)
	_la = p.GetTokenStream().LA(1)

	if !(_la == HTMLParserHTML_COMMENT || _la == HTMLParserHTML_CONDITIONAL_COMMENT) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IXhtmlCDATAContext is an interface to support dynamic dispatch.
type IXhtmlCDATAContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXhtmlCDATAContext differentiates from other interfaces.
	IsXhtmlCDATAContext()
}

type XhtmlCDATAContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXhtmlCDATAContext() *XhtmlCDATAContext {
	var p = new(XhtmlCDATAContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_xhtmlCDATA
	return p
}

func (*XhtmlCDATAContext) IsXhtmlCDATAContext() {}

func NewXhtmlCDATAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XhtmlCDATAContext {
	var p = new(XhtmlCDATAContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_xhtmlCDATA

	return p
}

func (s *XhtmlCDATAContext) GetParser() antlr.Parser { return s.parser }

func (s *XhtmlCDATAContext) CDATA() antlr.TerminalNode {
	return s.GetToken(HTMLParserCDATA, 0)
}

func (s *XhtmlCDATAContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XhtmlCDATAContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XhtmlCDATAContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterXhtmlCDATA(s)
	}
}

func (s *XhtmlCDATAContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitXhtmlCDATA(s)
	}
}

func (p *HTMLParser) XhtmlCDATA() (localctx IXhtmlCDATAContext) {
	localctx = NewXhtmlCDATAContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, HTMLParserRULE_xhtmlCDATA)

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
		p.Match(HTMLParserCDATA)
	}

	return localctx
}

// IDtdContext is an interface to support dynamic dispatch.
type IDtdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDtdContext differentiates from other interfaces.
	IsDtdContext()
}

type DtdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDtdContext() *DtdContext {
	var p = new(DtdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_dtd
	return p
}

func (*DtdContext) IsDtdContext() {}

func NewDtdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DtdContext {
	var p = new(DtdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_dtd

	return p
}

func (s *DtdContext) GetParser() antlr.Parser { return s.parser }

func (s *DtdContext) DTD() antlr.TerminalNode {
	return s.GetToken(HTMLParserDTD, 0)
}

func (s *DtdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DtdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DtdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterDtd(s)
	}
}

func (s *DtdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitDtd(s)
	}
}

func (p *HTMLParser) Dtd() (localctx IDtdContext) {
	localctx = NewDtdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, HTMLParserRULE_dtd)

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
		p.SetState(159)
		p.Match(HTMLParserDTD)
	}

	return localctx
}

// IXmlContext is an interface to support dynamic dispatch.
type IXmlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXmlContext differentiates from other interfaces.
	IsXmlContext()
}

type XmlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXmlContext() *XmlContext {
	var p = new(XmlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_xml
	return p
}

func (*XmlContext) IsXmlContext() {}

func NewXmlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XmlContext {
	var p = new(XmlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_xml

	return p
}

func (s *XmlContext) GetParser() antlr.Parser { return s.parser }

func (s *XmlContext) XML_DECLARATION() antlr.TerminalNode {
	return s.GetToken(HTMLParserXML_DECLARATION, 0)
}

func (s *XmlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XmlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XmlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterXml(s)
	}
}

func (s *XmlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitXml(s)
	}
}

func (p *HTMLParser) Xml() (localctx IXmlContext) {
	localctx = NewXmlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, HTMLParserRULE_xml)

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
		p.SetState(161)
		p.Match(HTMLParserXML_DECLARATION)
	}

	return localctx
}

// IScriptletContext is an interface to support dynamic dispatch.
type IScriptletContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptletContext differentiates from other interfaces.
	IsScriptletContext()
}

type ScriptletContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptletContext() *ScriptletContext {
	var p = new(ScriptletContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HTMLParserRULE_scriptlet
	return p
}

func (*ScriptletContext) IsScriptletContext() {}

func NewScriptletContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptletContext {
	var p = new(ScriptletContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HTMLParserRULE_scriptlet

	return p
}

func (s *ScriptletContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptletContext) SCRIPTLET() antlr.TerminalNode {
	return s.GetToken(HTMLParserSCRIPTLET, 0)
}

func (s *ScriptletContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptletContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptletContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.EnterScriptlet(s)
	}
}

func (s *ScriptletContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HTMLParserListener); ok {
		listenerT.ExitScriptlet(s)
	}
}

func (p *HTMLParser) Scriptlet() (localctx IScriptletContext) {
	localctx = NewScriptletContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, HTMLParserRULE_scriptlet)

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
		p.SetState(163)
		p.Match(HTMLParserSCRIPTLET)
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
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, HTMLParserRULE_script)
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
		p.SetState(165)
		p.Match(HTMLParserSCRIPT_OPEN)
	}
	p.SetState(166)
	_la = p.GetTokenStream().LA(1)

	if !(_la == HTMLParserSCRIPT_BODY || _la == HTMLParserSCRIPT_SHORT_BODY) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
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
	localctx = NewStyleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, HTMLParserRULE_style)
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
		p.Match(HTMLParserSTYLE_OPEN)
	}
	p.SetState(169)
	_la = p.GetTokenStream().LA(1)

	if !(_la == HTMLParserSTYLE_BODY || _la == HTMLParserSTYLE_SHORT_BODY) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
