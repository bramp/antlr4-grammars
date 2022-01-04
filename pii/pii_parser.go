// Code generated from pii.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pii // pii
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 57, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 5, 2, 17, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12,
	2, 3, 3, 2, 5, 6, 2, 51, 2, 16, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 33, 3,
	2, 2, 2, 8, 45, 3, 2, 2, 2, 10, 48, 3, 2, 2, 2, 12, 54, 3, 2, 2, 2, 14,
	17, 5, 4, 3, 2, 15, 17, 5, 6, 4, 2, 16, 14, 3, 2, 2, 2, 16, 15, 3, 2, 2,
	2, 17, 18, 3, 2, 2, 2, 18, 19, 5, 10, 6, 2, 19, 20, 5, 12, 7, 2, 20, 21,
	7, 2, 2, 3, 21, 3, 3, 2, 2, 2, 22, 23, 7, 3, 2, 2, 23, 24, 7, 6, 2, 2,
	24, 25, 7, 6, 2, 2, 25, 26, 7, 6, 2, 2, 26, 27, 7, 6, 2, 2, 27, 28, 7,
	6, 2, 2, 28, 29, 7, 6, 2, 2, 29, 30, 7, 6, 2, 2, 30, 31, 7, 6, 2, 2, 31,
	32, 5, 8, 5, 2, 32, 5, 3, 2, 2, 2, 33, 34, 7, 4, 2, 2, 34, 35, 7, 6, 2,
	2, 35, 36, 7, 6, 2, 2, 36, 37, 7, 6, 2, 2, 37, 38, 7, 6, 2, 2, 38, 39,
	7, 6, 2, 2, 39, 40, 7, 6, 2, 2, 40, 41, 7, 6, 2, 2, 41, 42, 7, 6, 2, 2,
	42, 43, 7, 6, 2, 2, 43, 44, 7, 6, 2, 2, 44, 7, 3, 2, 2, 2, 45, 46, 7, 6,
	2, 2, 46, 47, 7, 6, 2, 2, 47, 9, 3, 2, 2, 2, 48, 49, 7, 6, 2, 2, 49, 50,
	7, 6, 2, 2, 50, 51, 7, 6, 2, 2, 51, 52, 7, 6, 2, 2, 52, 53, 7, 6, 2, 2,
	53, 11, 3, 2, 2, 2, 54, 55, 9, 2, 2, 2, 55, 13, 3, 2, 2, 2, 3, 16,
}
var literalNames = []string{
	"", "'S'", "'B'", "'X'", "", "'('", "')'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "DIGIT", "LPAREN", "RPAREN", "DASH", "WS",
}

var ruleNames = []string{
	"pii", "issn", "isbn", "year", "item", "check",
}

type piiParser struct {
	*antlr.BaseParser
}

// NewpiiParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *piiParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewpiiParser(input antlr.TokenStream) *piiParser {
	this := new(piiParser)
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
	this.GrammarFileName = "pii.g4"

	return this
}

// piiParser tokens.
const (
	piiParserEOF    = antlr.TokenEOF
	piiParserT__0   = 1
	piiParserT__1   = 2
	piiParserT__2   = 3
	piiParserDIGIT  = 4
	piiParserLPAREN = 5
	piiParserRPAREN = 6
	piiParserDASH   = 7
	piiParserWS     = 8
)

// piiParser rules.
const (
	piiParserRULE_pii   = 0
	piiParserRULE_issn  = 1
	piiParserRULE_isbn  = 2
	piiParserRULE_year  = 3
	piiParserRULE_item  = 4
	piiParserRULE_check = 5
)

// IPiiContext is an interface to support dynamic dispatch.
type IPiiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPiiContext differentiates from other interfaces.
	IsPiiContext()
}

type PiiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPiiContext() *PiiContext {
	var p = new(PiiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = piiParserRULE_pii
	return p
}

func (*PiiContext) IsPiiContext() {}

func NewPiiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PiiContext {
	var p = new(PiiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = piiParserRULE_pii

	return p
}

func (s *PiiContext) GetParser() antlr.Parser { return s.parser }

func (s *PiiContext) Item() IItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *PiiContext) Check() ICheckContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheckContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheckContext)
}

func (s *PiiContext) EOF() antlr.TerminalNode {
	return s.GetToken(piiParserEOF, 0)
}

func (s *PiiContext) Issn() IIssnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIssnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIssnContext)
}

func (s *PiiContext) Isbn() IIsbnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsbnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsbnContext)
}

func (s *PiiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PiiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PiiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.EnterPii(s)
	}
}

func (s *PiiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.ExitPii(s)
	}
}

func (p *piiParser) Pii() (localctx IPiiContext) {
	this := p
	_ = this

	localctx = NewPiiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, piiParserRULE_pii)

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
	p.SetState(14)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case piiParserT__0:
		{
			p.SetState(12)
			p.Issn()
		}

	case piiParserT__1:
		{
			p.SetState(13)
			p.Isbn()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(16)
		p.Item()
	}
	{
		p.SetState(17)
		p.Check()
	}
	{
		p.SetState(18)
		p.Match(piiParserEOF)
	}

	return localctx
}

// IIssnContext is an interface to support dynamic dispatch.
type IIssnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIssnContext differentiates from other interfaces.
	IsIssnContext()
}

type IssnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIssnContext() *IssnContext {
	var p = new(IssnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = piiParserRULE_issn
	return p
}

func (*IssnContext) IsIssnContext() {}

func NewIssnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IssnContext {
	var p = new(IssnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = piiParserRULE_issn

	return p
}

func (s *IssnContext) GetParser() antlr.Parser { return s.parser }

func (s *IssnContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(piiParserDIGIT)
}

func (s *IssnContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(piiParserDIGIT, i)
}

func (s *IssnContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *IssnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IssnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IssnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.EnterIssn(s)
	}
}

func (s *IssnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.ExitIssn(s)
	}
}

func (p *piiParser) Issn() (localctx IIssnContext) {
	this := p
	_ = this

	localctx = NewIssnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, piiParserRULE_issn)

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
		p.SetState(20)
		p.Match(piiParserT__0)
	}
	{
		p.SetState(21)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(22)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(23)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(24)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(25)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(26)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(27)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(28)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(29)
		p.Year()
	}

	return localctx
}

// IIsbnContext is an interface to support dynamic dispatch.
type IIsbnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsbnContext differentiates from other interfaces.
	IsIsbnContext()
}

type IsbnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsbnContext() *IsbnContext {
	var p = new(IsbnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = piiParserRULE_isbn
	return p
}

func (*IsbnContext) IsIsbnContext() {}

func NewIsbnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsbnContext {
	var p = new(IsbnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = piiParserRULE_isbn

	return p
}

func (s *IsbnContext) GetParser() antlr.Parser { return s.parser }

func (s *IsbnContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(piiParserDIGIT)
}

func (s *IsbnContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(piiParserDIGIT, i)
}

func (s *IsbnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsbnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsbnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.EnterIsbn(s)
	}
}

func (s *IsbnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.ExitIsbn(s)
	}
}

func (p *piiParser) Isbn() (localctx IIsbnContext) {
	this := p
	_ = this

	localctx = NewIsbnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, piiParserRULE_isbn)

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
		p.SetState(31)
		p.Match(piiParserT__1)
	}
	{
		p.SetState(32)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(33)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(34)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(35)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(36)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(37)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(38)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(39)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(40)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(41)
		p.Match(piiParserDIGIT)
	}

	return localctx
}

// IYearContext is an interface to support dynamic dispatch.
type IYearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYearContext differentiates from other interfaces.
	IsYearContext()
}

type YearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYearContext() *YearContext {
	var p = new(YearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = piiParserRULE_year
	return p
}

func (*YearContext) IsYearContext() {}

func NewYearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YearContext {
	var p = new(YearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = piiParserRULE_year

	return p
}

func (s *YearContext) GetParser() antlr.Parser { return s.parser }

func (s *YearContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(piiParserDIGIT)
}

func (s *YearContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(piiParserDIGIT, i)
}

func (s *YearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.EnterYear(s)
	}
}

func (s *YearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.ExitYear(s)
	}
}

func (p *piiParser) Year() (localctx IYearContext) {
	this := p
	_ = this

	localctx = NewYearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, piiParserRULE_year)

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
		p.SetState(43)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(44)
		p.Match(piiParserDIGIT)
	}

	return localctx
}

// IItemContext is an interface to support dynamic dispatch.
type IItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemContext differentiates from other interfaces.
	IsItemContext()
}

type ItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemContext() *ItemContext {
	var p = new(ItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = piiParserRULE_item
	return p
}

func (*ItemContext) IsItemContext() {}

func NewItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemContext {
	var p = new(ItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = piiParserRULE_item

	return p
}

func (s *ItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(piiParserDIGIT)
}

func (s *ItemContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(piiParserDIGIT, i)
}

func (s *ItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.EnterItem(s)
	}
}

func (s *ItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.ExitItem(s)
	}
}

func (p *piiParser) Item() (localctx IItemContext) {
	this := p
	_ = this

	localctx = NewItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, piiParserRULE_item)

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
		p.SetState(46)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(47)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(48)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(49)
		p.Match(piiParserDIGIT)
	}
	{
		p.SetState(50)
		p.Match(piiParserDIGIT)
	}

	return localctx
}

// ICheckContext is an interface to support dynamic dispatch.
type ICheckContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCheckContext differentiates from other interfaces.
	IsCheckContext()
}

type CheckContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCheckContext() *CheckContext {
	var p = new(CheckContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = piiParserRULE_check
	return p
}

func (*CheckContext) IsCheckContext() {}

func NewCheckContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CheckContext {
	var p = new(CheckContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = piiParserRULE_check

	return p
}

func (s *CheckContext) GetParser() antlr.Parser { return s.parser }

func (s *CheckContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(piiParserDIGIT, 0)
}

func (s *CheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.EnterCheck(s)
	}
}

func (s *CheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(piiListener); ok {
		listenerT.ExitCheck(s)
	}
}

func (p *piiParser) Check() (localctx ICheckContext) {
	this := p
	_ = this

	localctx = NewCheckContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, piiParserRULE_check)
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

		if !(_la == piiParserT__2 || _la == piiParserDIGIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
