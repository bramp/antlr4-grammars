// Code generated from XMLParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package xml // XMLParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 98, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 5, 2, 20, 10, 2, 3, 2, 7, 2, 23, 10, 2, 12,
	2, 14, 2, 26, 11, 2, 3, 2, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11,
	2, 3, 3, 3, 3, 7, 3, 37, 10, 3, 12, 3, 14, 3, 40, 11, 3, 3, 3, 3, 3, 3,
	4, 5, 4, 45, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 52, 10, 4, 3, 4,
	5, 4, 55, 10, 4, 7, 4, 57, 10, 4, 12, 4, 14, 4, 60, 11, 4, 3, 5, 3, 5,
	3, 5, 7, 5, 65, 10, 5, 12, 5, 14, 5, 68, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 80, 10, 5, 12, 5, 14, 5, 83,
	11, 5, 3, 5, 5, 5, 86, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 5, 3,
	2, 6, 7, 4, 2, 8, 8, 11, 11, 5, 2, 3, 3, 8, 8, 20, 20, 2, 103, 2, 19, 3,
	2, 2, 2, 4, 34, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 85, 3, 2, 2, 2, 10, 87,
	3, 2, 2, 2, 12, 89, 3, 2, 2, 2, 14, 93, 3, 2, 2, 2, 16, 95, 3, 2, 2, 2,
	18, 20, 5, 4, 3, 2, 19, 18, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 24, 3,
	2, 2, 2, 21, 23, 5, 16, 9, 2, 22, 21, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24,
	22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 27, 3, 2, 2, 2, 26, 24, 3, 2, 2,
	2, 27, 31, 5, 8, 5, 2, 28, 30, 5, 16, 9, 2, 29, 28, 3, 2, 2, 2, 30, 33,
	3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 3, 3, 2, 2, 2,
	33, 31, 3, 2, 2, 2, 34, 38, 7, 10, 2, 2, 35, 37, 5, 12, 7, 2, 36, 35, 3,
	2, 2, 2, 37, 40, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39,
	41, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 41, 42, 7, 13, 2, 2, 42, 5, 3, 2, 2,
	2, 43, 45, 5, 14, 8, 2, 44, 43, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 58,
	3, 2, 2, 2, 46, 52, 5, 8, 5, 2, 47, 52, 5, 10, 6, 2, 48, 52, 7, 4, 2, 2,
	49, 52, 7, 20, 2, 2, 50, 52, 7, 3, 2, 2, 51, 46, 3, 2, 2, 2, 51, 47, 3,
	2, 2, 2, 51, 48, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52,
	54, 3, 2, 2, 2, 53, 55, 5, 14, 8, 2, 54, 53, 3, 2, 2, 2, 54, 55, 3, 2,
	2, 2, 55, 57, 3, 2, 2, 2, 56, 51, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 7, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2,
	61, 62, 7, 9, 2, 2, 62, 66, 7, 18, 2, 2, 63, 65, 5, 12, 7, 2, 64, 63, 3,
	2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67,
	69, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 70, 7, 12, 2, 2, 70, 71, 5, 6,
	4, 2, 71, 72, 7, 9, 2, 2, 72, 73, 7, 15, 2, 2, 73, 74, 7, 18, 2, 2, 74,
	75, 7, 12, 2, 2, 75, 86, 3, 2, 2, 2, 76, 77, 7, 9, 2, 2, 77, 81, 7, 18,
	2, 2, 78, 80, 5, 12, 7, 2, 79, 78, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81,
	79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2, 83, 81, 3, 2, 2,
	2, 84, 86, 7, 14, 2, 2, 85, 61, 3, 2, 2, 2, 85, 76, 3, 2, 2, 2, 86, 9,
	3, 2, 2, 2, 87, 88, 9, 2, 2, 2, 88, 11, 3, 2, 2, 2, 89, 90, 7, 18, 2, 2,
	90, 91, 7, 16, 2, 2, 91, 92, 7, 17, 2, 2, 92, 13, 3, 2, 2, 2, 93, 94, 9,
	3, 2, 2, 94, 15, 3, 2, 2, 2, 95, 96, 9, 4, 2, 2, 96, 17, 3, 2, 2, 2, 13,
	19, 24, 31, 38, 44, 51, 54, 58, 66, 81, 85,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "'<'", "", "", "'>'", "", "'/>'", "'/'", "'='",
}
var symbolicNames = []string{
	"", "COMMENT", "CDATA", "DTD", "EntityRef", "CharRef", "SEA_WS", "OPEN",
	"XMLDeclOpen", "TEXT", "CLOSE", "SPECIAL_CLOSE", "SLASH_CLOSE", "SLASH",
	"EQUALS", "STRING", "Name", "S", "PI",
}

var ruleNames = []string{
	"document", "prolog", "content", "element", "reference", "attribute", "chardata",
	"misc",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type XMLParser struct {
	*antlr.BaseParser
}

func NewXMLParser(input antlr.TokenStream) *XMLParser {
	this := new(XMLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "XMLParser.g4"

	return this
}

// XMLParser tokens.
const (
	XMLParserEOF           = antlr.TokenEOF
	XMLParserCOMMENT       = 1
	XMLParserCDATA         = 2
	XMLParserDTD           = 3
	XMLParserEntityRef     = 4
	XMLParserCharRef       = 5
	XMLParserSEA_WS        = 6
	XMLParserOPEN          = 7
	XMLParserXMLDeclOpen   = 8
	XMLParserTEXT          = 9
	XMLParserCLOSE         = 10
	XMLParserSPECIAL_CLOSE = 11
	XMLParserSLASH_CLOSE   = 12
	XMLParserSLASH         = 13
	XMLParserEQUALS        = 14
	XMLParserSTRING        = 15
	XMLParserName          = 16
	XMLParserS             = 17
	XMLParserPI            = 18
)

// XMLParser rules.
const (
	XMLParserRULE_document  = 0
	XMLParserRULE_prolog    = 1
	XMLParserRULE_content   = 2
	XMLParserRULE_element   = 3
	XMLParserRULE_reference = 4
	XMLParserRULE_attribute = 5
	XMLParserRULE_chardata  = 6
	XMLParserRULE_misc      = 7
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
	p.RuleIndex = XMLParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XMLParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *DocumentContext) Prolog() IPrologContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrologContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrologContext)
}

func (s *DocumentContext) AllMisc() []IMiscContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMiscContext)(nil)).Elem())
	var tst = make([]IMiscContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMiscContext)
		}
	}

	return tst
}

func (s *DocumentContext) Misc(i int) IMiscContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMiscContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMiscContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *XMLParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, XMLParserRULE_document)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == XMLParserXMLDeclOpen {
		{
			p.SetState(16)
			p.Prolog()
		}

	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<XMLParserCOMMENT)|(1<<XMLParserSEA_WS)|(1<<XMLParserPI))) != 0 {
		{
			p.SetState(19)
			p.Misc()
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.Element()
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<XMLParserCOMMENT)|(1<<XMLParserSEA_WS)|(1<<XMLParserPI))) != 0 {
		{
			p.SetState(26)
			p.Misc()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPrologContext is an interface to support dynamic dispatch.
type IPrologContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrologContext differentiates from other interfaces.
	IsPrologContext()
}

type PrologContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrologContext() *PrologContext {
	var p = new(PrologContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XMLParserRULE_prolog
	return p
}

func (*PrologContext) IsPrologContext() {}

func NewPrologContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrologContext {
	var p = new(PrologContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XMLParserRULE_prolog

	return p
}

func (s *PrologContext) GetParser() antlr.Parser { return s.parser }

func (s *PrologContext) XMLDeclOpen() antlr.TerminalNode {
	return s.GetToken(XMLParserXMLDeclOpen, 0)
}

func (s *PrologContext) SPECIAL_CLOSE() antlr.TerminalNode {
	return s.GetToken(XMLParserSPECIAL_CLOSE, 0)
}

func (s *PrologContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *PrologContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *PrologContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrologContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrologContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.EnterProlog(s)
	}
}

func (s *PrologContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.ExitProlog(s)
	}
}

func (p *XMLParser) Prolog() (localctx IPrologContext) {
	localctx = NewPrologContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, XMLParserRULE_prolog)
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
		p.SetState(32)
		p.Match(XMLParserXMLDeclOpen)
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == XMLParserName {
		{
			p.SetState(33)
			p.Attribute()
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(XMLParserSPECIAL_CLOSE)
	}

	return localctx
}

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XMLParserRULE_content
	return p
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XMLParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) AllChardata() []IChardataContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IChardataContext)(nil)).Elem())
	var tst = make([]IChardataContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IChardataContext)
		}
	}

	return tst
}

func (s *ContentContext) Chardata(i int) IChardataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChardataContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IChardataContext)
}

func (s *ContentContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *ContentContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ContentContext) AllReference() []IReferenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReferenceContext)(nil)).Elem())
	var tst = make([]IReferenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReferenceContext)
		}
	}

	return tst
}

func (s *ContentContext) Reference(i int) IReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *ContentContext) AllCDATA() []antlr.TerminalNode {
	return s.GetTokens(XMLParserCDATA)
}

func (s *ContentContext) CDATA(i int) antlr.TerminalNode {
	return s.GetToken(XMLParserCDATA, i)
}

func (s *ContentContext) AllPI() []antlr.TerminalNode {
	return s.GetTokens(XMLParserPI)
}

func (s *ContentContext) PI(i int) antlr.TerminalNode {
	return s.GetToken(XMLParserPI, i)
}

func (s *ContentContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(XMLParserCOMMENT)
}

func (s *ContentContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(XMLParserCOMMENT, i)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.ExitContent(s)
	}
}

func (p *XMLParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, XMLParserRULE_content)
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == XMLParserSEA_WS || _la == XMLParserTEXT {
		{
			p.SetState(41)
			p.Chardata()
		}

	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(49)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case XMLParserOPEN:
				{
					p.SetState(44)
					p.Element()
				}

			case XMLParserEntityRef, XMLParserCharRef:
				{
					p.SetState(45)
					p.Reference()
				}

			case XMLParserCDATA:
				{
					p.SetState(46)
					p.Match(XMLParserCDATA)
				}

			case XMLParserPI:
				{
					p.SetState(47)
					p.Match(XMLParserPI)
				}

			case XMLParserCOMMENT:
				{
					p.SetState(48)
					p.Match(XMLParserCOMMENT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == XMLParserSEA_WS || _la == XMLParserTEXT {
				{
					p.SetState(51)
					p.Chardata()
				}

			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XMLParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XMLParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) AllOPEN() []antlr.TerminalNode {
	return s.GetTokens(XMLParserOPEN)
}

func (s *ElementContext) OPEN(i int) antlr.TerminalNode {
	return s.GetToken(XMLParserOPEN, i)
}

func (s *ElementContext) AllName() []antlr.TerminalNode {
	return s.GetTokens(XMLParserName)
}

func (s *ElementContext) Name(i int) antlr.TerminalNode {
	return s.GetToken(XMLParserName, i)
}

func (s *ElementContext) AllCLOSE() []antlr.TerminalNode {
	return s.GetTokens(XMLParserCLOSE)
}

func (s *ElementContext) CLOSE(i int) antlr.TerminalNode {
	return s.GetToken(XMLParserCLOSE, i)
}

func (s *ElementContext) Content() IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *ElementContext) SLASH() antlr.TerminalNode {
	return s.GetToken(XMLParserSLASH, 0)
}

func (s *ElementContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *ElementContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *ElementContext) SLASH_CLOSE() antlr.TerminalNode {
	return s.GetToken(XMLParserSLASH_CLOSE, 0)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.ExitElement(s)
	}
}

func (p *XMLParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, XMLParserRULE_element)
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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Match(XMLParserOPEN)
		}
		{
			p.SetState(60)
			p.Match(XMLParserName)
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == XMLParserName {
			{
				p.SetState(61)
				p.Attribute()
			}

			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(67)
			p.Match(XMLParserCLOSE)
		}
		{
			p.SetState(68)
			p.Content()
		}
		{
			p.SetState(69)
			p.Match(XMLParserOPEN)
		}
		{
			p.SetState(70)
			p.Match(XMLParserSLASH)
		}
		{
			p.SetState(71)
			p.Match(XMLParserName)
		}
		{
			p.SetState(72)
			p.Match(XMLParserCLOSE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Match(XMLParserOPEN)
		}
		{
			p.SetState(75)
			p.Match(XMLParserName)
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == XMLParserName {
			{
				p.SetState(76)
				p.Attribute()
			}

			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(82)
			p.Match(XMLParserSLASH_CLOSE)
		}

	}

	return localctx
}

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XMLParserRULE_reference
	return p
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XMLParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) EntityRef() antlr.TerminalNode {
	return s.GetToken(XMLParserEntityRef, 0)
}

func (s *ReferenceContext) CharRef() antlr.TerminalNode {
	return s.GetToken(XMLParserCharRef, 0)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.ExitReference(s)
	}
}

func (p *XMLParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, XMLParserRULE_reference)
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
		p.SetState(85)
		_la = p.GetTokenStream().LA(1)

		if !(_la == XMLParserEntityRef || _la == XMLParserCharRef) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XMLParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XMLParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) Name() antlr.TerminalNode {
	return s.GetToken(XMLParserName, 0)
}

func (s *AttributeContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(XMLParserEQUALS, 0)
}

func (s *AttributeContext) STRING() antlr.TerminalNode {
	return s.GetToken(XMLParserSTRING, 0)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *XMLParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, XMLParserRULE_attribute)

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
		p.SetState(87)
		p.Match(XMLParserName)
	}
	{
		p.SetState(88)
		p.Match(XMLParserEQUALS)
	}
	{
		p.SetState(89)
		p.Match(XMLParserSTRING)
	}

	return localctx
}

// IChardataContext is an interface to support dynamic dispatch.
type IChardataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChardataContext differentiates from other interfaces.
	IsChardataContext()
}

type ChardataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChardataContext() *ChardataContext {
	var p = new(ChardataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XMLParserRULE_chardata
	return p
}

func (*ChardataContext) IsChardataContext() {}

func NewChardataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChardataContext {
	var p = new(ChardataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XMLParserRULE_chardata

	return p
}

func (s *ChardataContext) GetParser() antlr.Parser { return s.parser }

func (s *ChardataContext) TEXT() antlr.TerminalNode {
	return s.GetToken(XMLParserTEXT, 0)
}

func (s *ChardataContext) SEA_WS() antlr.TerminalNode {
	return s.GetToken(XMLParserSEA_WS, 0)
}

func (s *ChardataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChardataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChardataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.EnterChardata(s)
	}
}

func (s *ChardataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.ExitChardata(s)
	}
}

func (p *XMLParser) Chardata() (localctx IChardataContext) {
	localctx = NewChardataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, XMLParserRULE_chardata)
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
		p.SetState(91)
		_la = p.GetTokenStream().LA(1)

		if !(_la == XMLParserSEA_WS || _la == XMLParserTEXT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMiscContext is an interface to support dynamic dispatch.
type IMiscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMiscContext differentiates from other interfaces.
	IsMiscContext()
}

type MiscContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMiscContext() *MiscContext {
	var p = new(MiscContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = XMLParserRULE_misc
	return p
}

func (*MiscContext) IsMiscContext() {}

func NewMiscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MiscContext {
	var p = new(MiscContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = XMLParserRULE_misc

	return p
}

func (s *MiscContext) GetParser() antlr.Parser { return s.parser }

func (s *MiscContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(XMLParserCOMMENT, 0)
}

func (s *MiscContext) PI() antlr.TerminalNode {
	return s.GetToken(XMLParserPI, 0)
}

func (s *MiscContext) SEA_WS() antlr.TerminalNode {
	return s.GetToken(XMLParserSEA_WS, 0)
}

func (s *MiscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MiscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MiscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.EnterMisc(s)
	}
}

func (s *MiscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(XMLParserListener); ok {
		listenerT.ExitMisc(s)
	}
}

func (p *XMLParser) Misc() (localctx IMiscContext) {
	localctx = NewMiscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, XMLParserRULE_misc)
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
		p.SetState(93)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<XMLParserCOMMENT)|(1<<XMLParserSEA_WS)|(1<<XMLParserPI))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
