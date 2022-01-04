// Code generated from bibcode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bibcode // bibcode
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 99, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4,
	38, 10, 4, 3, 4, 3, 4, 3, 4, 5, 4, 43, 10, 4, 3, 4, 3, 4, 3, 4, 5, 4, 48,
	10, 4, 3, 4, 3, 4, 3, 4, 5, 4, 53, 10, 4, 3, 4, 3, 4, 3, 4, 5, 4, 58, 10,
	4, 3, 5, 3, 5, 3, 5, 5, 5, 63, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 68, 10, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 73, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 78, 10, 5, 3,
	6, 5, 6, 81, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 6, 8, 89, 10, 8,
	13, 8, 14, 8, 90, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 2, 2,
	12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 3, 3, 2, 4, 5, 2, 109, 2, 22,
	3, 2, 2, 2, 4, 29, 3, 2, 2, 2, 6, 37, 3, 2, 2, 2, 8, 62, 3, 2, 2, 2, 10,
	80, 3, 2, 2, 2, 12, 84, 3, 2, 2, 2, 14, 88, 3, 2, 2, 2, 16, 92, 3, 2, 2,
	2, 18, 94, 3, 2, 2, 2, 20, 96, 3, 2, 2, 2, 22, 23, 5, 4, 3, 2, 23, 24,
	5, 6, 4, 2, 24, 25, 5, 8, 5, 2, 25, 26, 5, 10, 6, 2, 26, 27, 5, 16, 9,
	2, 27, 28, 7, 2, 2, 3, 28, 3, 3, 2, 2, 2, 29, 30, 7, 6, 2, 2, 30, 31, 7,
	6, 2, 2, 31, 32, 7, 6, 2, 2, 32, 33, 7, 6, 2, 2, 33, 5, 3, 2, 2, 2, 34,
	38, 5, 18, 10, 2, 35, 38, 5, 20, 11, 2, 36, 38, 7, 3, 2, 2, 37, 34, 3,
	2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 42, 3, 2, 2, 2, 39,
	43, 5, 18, 10, 2, 40, 43, 5, 20, 11, 2, 41, 43, 7, 3, 2, 2, 42, 39, 3,
	2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 47, 3, 2, 2, 2, 44,
	48, 5, 18, 10, 2, 45, 48, 5, 20, 11, 2, 46, 48, 7, 3, 2, 2, 47, 44, 3,
	2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 52, 3, 2, 2, 2, 49,
	53, 5, 18, 10, 2, 50, 53, 5, 20, 11, 2, 51, 53, 7, 3, 2, 2, 52, 49, 3,
	2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 57, 3, 2, 2, 2, 54,
	58, 5, 18, 10, 2, 55, 58, 5, 20, 11, 2, 56, 58, 7, 3, 2, 2, 57, 54, 3,
	2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 7, 3, 2, 2, 2, 59,
	63, 5, 18, 10, 2, 60, 63, 5, 20, 11, 2, 61, 63, 7, 3, 2, 2, 62, 59, 3,
	2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 67, 3, 2, 2, 2, 64,
	68, 5, 18, 10, 2, 65, 68, 5, 20, 11, 2, 66, 68, 7, 3, 2, 2, 67, 64, 3,
	2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 66, 3, 2, 2, 2, 68, 72, 3, 2, 2, 2, 69,
	73, 5, 18, 10, 2, 70, 73, 5, 20, 11, 2, 71, 73, 7, 3, 2, 2, 72, 69, 3,
	2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 71, 3, 2, 2, 2, 73, 77, 3, 2, 2, 2, 74,
	78, 5, 18, 10, 2, 75, 78, 5, 20, 11, 2, 76, 78, 7, 3, 2, 2, 77, 74, 3,
	2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 76, 3, 2, 2, 2, 78, 9, 3, 2, 2, 2, 79,
	81, 5, 12, 7, 2, 80, 79, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 82, 3, 2,
	2, 2, 82, 83, 5, 14, 8, 2, 83, 11, 3, 2, 2, 2, 84, 85, 5, 18, 10, 2, 85,
	13, 3, 2, 2, 2, 86, 89, 5, 20, 11, 2, 87, 89, 7, 3, 2, 2, 88, 86, 3, 2,
	2, 2, 88, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91,
	3, 2, 2, 2, 91, 15, 3, 2, 2, 2, 92, 93, 5, 18, 10, 2, 93, 17, 3, 2, 2,
	2, 94, 95, 9, 2, 2, 2, 95, 19, 3, 2, 2, 2, 96, 97, 7, 6, 2, 2, 97, 21,
	3, 2, 2, 2, 14, 37, 42, 47, 52, 57, 62, 67, 72, 77, 80, 88, 90,
}
var literalNames = []string{
	"", "'.'",
}
var symbolicNames = []string{
	"", "DOT", "UPPERLETTER", "LOWERLETTER", "DIGIT", "WS",
}

var ruleNames = []string{
	"bibcode", "year", "publish", "volume", "pagesection", "section", "page",
	"author", "letter", "digit",
}

type bibcodeParser struct {
	*antlr.BaseParser
}

// NewbibcodeParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *bibcodeParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewbibcodeParser(input antlr.TokenStream) *bibcodeParser {
	this := new(bibcodeParser)
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
	this.GrammarFileName = "bibcode.g4"

	return this
}

// bibcodeParser tokens.
const (
	bibcodeParserEOF         = antlr.TokenEOF
	bibcodeParserDOT         = 1
	bibcodeParserUPPERLETTER = 2
	bibcodeParserLOWERLETTER = 3
	bibcodeParserDIGIT       = 4
	bibcodeParserWS          = 5
)

// bibcodeParser rules.
const (
	bibcodeParserRULE_bibcode     = 0
	bibcodeParserRULE_year        = 1
	bibcodeParserRULE_publish     = 2
	bibcodeParserRULE_volume      = 3
	bibcodeParserRULE_pagesection = 4
	bibcodeParserRULE_section     = 5
	bibcodeParserRULE_page        = 6
	bibcodeParserRULE_author      = 7
	bibcodeParserRULE_letter      = 8
	bibcodeParserRULE_digit       = 9
)

// IBibcodeContext is an interface to support dynamic dispatch.
type IBibcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBibcodeContext differentiates from other interfaces.
	IsBibcodeContext()
}

type BibcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBibcodeContext() *BibcodeContext {
	var p = new(BibcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bibcodeParserRULE_bibcode
	return p
}

func (*BibcodeContext) IsBibcodeContext() {}

func NewBibcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BibcodeContext {
	var p = new(BibcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bibcodeParserRULE_bibcode

	return p
}

func (s *BibcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *BibcodeContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *BibcodeContext) Publish() IPublishContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPublishContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPublishContext)
}

func (s *BibcodeContext) Volume() IVolumeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVolumeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVolumeContext)
}

func (s *BibcodeContext) Pagesection() IPagesectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPagesectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPagesectionContext)
}

func (s *BibcodeContext) Author() IAuthorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAuthorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAuthorContext)
}

func (s *BibcodeContext) EOF() antlr.TerminalNode {
	return s.GetToken(bibcodeParserEOF, 0)
}

func (s *BibcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BibcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BibcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.EnterBibcode(s)
	}
}

func (s *BibcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.ExitBibcode(s)
	}
}

func (p *bibcodeParser) Bibcode() (localctx IBibcodeContext) {
	this := p
	_ = this

	localctx = NewBibcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, bibcodeParserRULE_bibcode)

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
		p.Year()
	}
	{
		p.SetState(21)
		p.Publish()
	}
	{
		p.SetState(22)
		p.Volume()
	}
	{
		p.SetState(23)
		p.Pagesection()
	}
	{
		p.SetState(24)
		p.Author()
	}
	{
		p.SetState(25)
		p.Match(bibcodeParserEOF)
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
	p.RuleIndex = bibcodeParserRULE_year
	return p
}

func (*YearContext) IsYearContext() {}

func NewYearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YearContext {
	var p = new(YearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bibcodeParserRULE_year

	return p
}

func (s *YearContext) GetParser() antlr.Parser { return s.parser }

func (s *YearContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(bibcodeParserDIGIT)
}

func (s *YearContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(bibcodeParserDIGIT, i)
}

func (s *YearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.EnterYear(s)
	}
}

func (s *YearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.ExitYear(s)
	}
}

func (p *bibcodeParser) Year() (localctx IYearContext) {
	this := p
	_ = this

	localctx = NewYearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, bibcodeParserRULE_year)

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
		p.SetState(27)
		p.Match(bibcodeParserDIGIT)
	}
	{
		p.SetState(28)
		p.Match(bibcodeParserDIGIT)
	}
	{
		p.SetState(29)
		p.Match(bibcodeParserDIGIT)
	}
	{
		p.SetState(30)
		p.Match(bibcodeParserDIGIT)
	}

	return localctx
}

// IPublishContext is an interface to support dynamic dispatch.
type IPublishContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPublishContext differentiates from other interfaces.
	IsPublishContext()
}

type PublishContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPublishContext() *PublishContext {
	var p = new(PublishContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bibcodeParserRULE_publish
	return p
}

func (*PublishContext) IsPublishContext() {}

func NewPublishContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PublishContext {
	var p = new(PublishContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bibcodeParserRULE_publish

	return p
}

func (s *PublishContext) GetParser() antlr.Parser { return s.parser }

func (s *PublishContext) AllLetter() []ILetterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILetterContext)(nil)).Elem())
	var tst = make([]ILetterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILetterContext)
		}
	}

	return tst
}

func (s *PublishContext) Letter(i int) ILetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILetterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILetterContext)
}

func (s *PublishContext) AllDigit() []IDigitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDigitContext)(nil)).Elem())
	var tst = make([]IDigitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDigitContext)
		}
	}

	return tst
}

func (s *PublishContext) Digit(i int) IDigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDigitContext)
}

func (s *PublishContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(bibcodeParserDOT)
}

func (s *PublishContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(bibcodeParserDOT, i)
}

func (s *PublishContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PublishContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PublishContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.EnterPublish(s)
	}
}

func (s *PublishContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.ExitPublish(s)
	}
}

func (p *bibcodeParser) Publish() (localctx IPublishContext) {
	this := p
	_ = this

	localctx = NewPublishContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, bibcodeParserRULE_publish)

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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bibcodeParserUPPERLETTER, bibcodeParserLOWERLETTER:
		{
			p.SetState(32)
			p.Letter()
		}

	case bibcodeParserDIGIT:
		{
			p.SetState(33)
			p.Digit()
		}

	case bibcodeParserDOT:
		{
			p.SetState(34)
			p.Match(bibcodeParserDOT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bibcodeParserUPPERLETTER, bibcodeParserLOWERLETTER:
		{
			p.SetState(37)
			p.Letter()
		}

	case bibcodeParserDIGIT:
		{
			p.SetState(38)
			p.Digit()
		}

	case bibcodeParserDOT:
		{
			p.SetState(39)
			p.Match(bibcodeParserDOT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bibcodeParserUPPERLETTER, bibcodeParserLOWERLETTER:
		{
			p.SetState(42)
			p.Letter()
		}

	case bibcodeParserDIGIT:
		{
			p.SetState(43)
			p.Digit()
		}

	case bibcodeParserDOT:
		{
			p.SetState(44)
			p.Match(bibcodeParserDOT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bibcodeParserUPPERLETTER, bibcodeParserLOWERLETTER:
		{
			p.SetState(47)
			p.Letter()
		}

	case bibcodeParserDIGIT:
		{
			p.SetState(48)
			p.Digit()
		}

	case bibcodeParserDOT:
		{
			p.SetState(49)
			p.Match(bibcodeParserDOT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bibcodeParserUPPERLETTER, bibcodeParserLOWERLETTER:
		{
			p.SetState(52)
			p.Letter()
		}

	case bibcodeParserDIGIT:
		{
			p.SetState(53)
			p.Digit()
		}

	case bibcodeParserDOT:
		{
			p.SetState(54)
			p.Match(bibcodeParserDOT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVolumeContext is an interface to support dynamic dispatch.
type IVolumeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVolumeContext differentiates from other interfaces.
	IsVolumeContext()
}

type VolumeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVolumeContext() *VolumeContext {
	var p = new(VolumeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bibcodeParserRULE_volume
	return p
}

func (*VolumeContext) IsVolumeContext() {}

func NewVolumeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VolumeContext {
	var p = new(VolumeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bibcodeParserRULE_volume

	return p
}

func (s *VolumeContext) GetParser() antlr.Parser { return s.parser }

func (s *VolumeContext) AllLetter() []ILetterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILetterContext)(nil)).Elem())
	var tst = make([]ILetterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILetterContext)
		}
	}

	return tst
}

func (s *VolumeContext) Letter(i int) ILetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILetterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILetterContext)
}

func (s *VolumeContext) AllDigit() []IDigitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDigitContext)(nil)).Elem())
	var tst = make([]IDigitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDigitContext)
		}
	}

	return tst
}

func (s *VolumeContext) Digit(i int) IDigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDigitContext)
}

func (s *VolumeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(bibcodeParserDOT)
}

func (s *VolumeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(bibcodeParserDOT, i)
}

func (s *VolumeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VolumeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VolumeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.EnterVolume(s)
	}
}

func (s *VolumeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.ExitVolume(s)
	}
}

func (p *bibcodeParser) Volume() (localctx IVolumeContext) {
	this := p
	_ = this

	localctx = NewVolumeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, bibcodeParserRULE_volume)

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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bibcodeParserUPPERLETTER, bibcodeParserLOWERLETTER:
		{
			p.SetState(57)
			p.Letter()
		}

	case bibcodeParserDIGIT:
		{
			p.SetState(58)
			p.Digit()
		}

	case bibcodeParserDOT:
		{
			p.SetState(59)
			p.Match(bibcodeParserDOT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bibcodeParserUPPERLETTER, bibcodeParserLOWERLETTER:
		{
			p.SetState(62)
			p.Letter()
		}

	case bibcodeParserDIGIT:
		{
			p.SetState(63)
			p.Digit()
		}

	case bibcodeParserDOT:
		{
			p.SetState(64)
			p.Match(bibcodeParserDOT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bibcodeParserUPPERLETTER, bibcodeParserLOWERLETTER:
		{
			p.SetState(67)
			p.Letter()
		}

	case bibcodeParserDIGIT:
		{
			p.SetState(68)
			p.Digit()
		}

	case bibcodeParserDOT:
		{
			p.SetState(69)
			p.Match(bibcodeParserDOT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bibcodeParserUPPERLETTER, bibcodeParserLOWERLETTER:
		{
			p.SetState(72)
			p.Letter()
		}

	case bibcodeParserDIGIT:
		{
			p.SetState(73)
			p.Digit()
		}

	case bibcodeParserDOT:
		{
			p.SetState(74)
			p.Match(bibcodeParserDOT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPagesectionContext is an interface to support dynamic dispatch.
type IPagesectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPagesectionContext differentiates from other interfaces.
	IsPagesectionContext()
}

type PagesectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPagesectionContext() *PagesectionContext {
	var p = new(PagesectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bibcodeParserRULE_pagesection
	return p
}

func (*PagesectionContext) IsPagesectionContext() {}

func NewPagesectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PagesectionContext {
	var p = new(PagesectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bibcodeParserRULE_pagesection

	return p
}

func (s *PagesectionContext) GetParser() antlr.Parser { return s.parser }

func (s *PagesectionContext) Page() IPageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPageContext)
}

func (s *PagesectionContext) Section() ISectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *PagesectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PagesectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PagesectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.EnterPagesection(s)
	}
}

func (s *PagesectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.ExitPagesection(s)
	}
}

func (p *bibcodeParser) Pagesection() (localctx IPagesectionContext) {
	this := p
	_ = this

	localctx = NewPagesectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, bibcodeParserRULE_pagesection)
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
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == bibcodeParserUPPERLETTER || _la == bibcodeParserLOWERLETTER {
		{
			p.SetState(77)
			p.Section()
		}

	}
	{
		p.SetState(80)
		p.Page()
	}

	return localctx
}

// ISectionContext is an interface to support dynamic dispatch.
type ISectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionContext differentiates from other interfaces.
	IsSectionContext()
}

type SectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionContext() *SectionContext {
	var p = new(SectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bibcodeParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bibcodeParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) Letter() ILetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILetterContext)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.EnterSection(s)
	}
}

func (s *SectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.ExitSection(s)
	}
}

func (p *bibcodeParser) Section() (localctx ISectionContext) {
	this := p
	_ = this

	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, bibcodeParserRULE_section)

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
		p.Letter()
	}

	return localctx
}

// IPageContext is an interface to support dynamic dispatch.
type IPageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPageContext differentiates from other interfaces.
	IsPageContext()
}

type PageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPageContext() *PageContext {
	var p = new(PageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bibcodeParserRULE_page
	return p
}

func (*PageContext) IsPageContext() {}

func NewPageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PageContext {
	var p = new(PageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bibcodeParserRULE_page

	return p
}

func (s *PageContext) GetParser() antlr.Parser { return s.parser }

func (s *PageContext) AllDigit() []IDigitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDigitContext)(nil)).Elem())
	var tst = make([]IDigitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDigitContext)
		}
	}

	return tst
}

func (s *PageContext) Digit(i int) IDigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDigitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDigitContext)
}

func (s *PageContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(bibcodeParserDOT)
}

func (s *PageContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(bibcodeParserDOT, i)
}

func (s *PageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.EnterPage(s)
	}
}

func (s *PageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.ExitPage(s)
	}
}

func (p *bibcodeParser) Page() (localctx IPageContext) {
	this := p
	_ = this

	localctx = NewPageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, bibcodeParserRULE_page)
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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == bibcodeParserDOT || _la == bibcodeParserDIGIT {
		p.SetState(86)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case bibcodeParserDIGIT:
			{
				p.SetState(84)
				p.Digit()
			}

		case bibcodeParserDOT:
			{
				p.SetState(85)
				p.Match(bibcodeParserDOT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAuthorContext is an interface to support dynamic dispatch.
type IAuthorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAuthorContext differentiates from other interfaces.
	IsAuthorContext()
}

type AuthorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuthorContext() *AuthorContext {
	var p = new(AuthorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bibcodeParserRULE_author
	return p
}

func (*AuthorContext) IsAuthorContext() {}

func NewAuthorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuthorContext {
	var p = new(AuthorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bibcodeParserRULE_author

	return p
}

func (s *AuthorContext) GetParser() antlr.Parser { return s.parser }

func (s *AuthorContext) Letter() ILetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILetterContext)
}

func (s *AuthorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuthorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuthorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.EnterAuthor(s)
	}
}

func (s *AuthorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.ExitAuthor(s)
	}
}

func (p *bibcodeParser) Author() (localctx IAuthorContext) {
	this := p
	_ = this

	localctx = NewAuthorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, bibcodeParserRULE_author)

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
		p.SetState(90)
		p.Letter()
	}

	return localctx
}

// ILetterContext is an interface to support dynamic dispatch.
type ILetterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLetterContext differentiates from other interfaces.
	IsLetterContext()
}

type LetterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetterContext() *LetterContext {
	var p = new(LetterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bibcodeParserRULE_letter
	return p
}

func (*LetterContext) IsLetterContext() {}

func NewLetterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetterContext {
	var p = new(LetterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bibcodeParserRULE_letter

	return p
}

func (s *LetterContext) GetParser() antlr.Parser { return s.parser }

func (s *LetterContext) UPPERLETTER() antlr.TerminalNode {
	return s.GetToken(bibcodeParserUPPERLETTER, 0)
}

func (s *LetterContext) LOWERLETTER() antlr.TerminalNode {
	return s.GetToken(bibcodeParserLOWERLETTER, 0)
}

func (s *LetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.EnterLetter(s)
	}
}

func (s *LetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.ExitLetter(s)
	}
}

func (p *bibcodeParser) Letter() (localctx ILetterContext) {
	this := p
	_ = this

	localctx = NewLetterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, bibcodeParserRULE_letter)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == bibcodeParserUPPERLETTER || _la == bibcodeParserLOWERLETTER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDigitContext is an interface to support dynamic dispatch.
type IDigitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDigitContext differentiates from other interfaces.
	IsDigitContext()
}

type DigitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDigitContext() *DigitContext {
	var p = new(DigitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bibcodeParserRULE_digit
	return p
}

func (*DigitContext) IsDigitContext() {}

func NewDigitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DigitContext {
	var p = new(DigitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bibcodeParserRULE_digit

	return p
}

func (s *DigitContext) GetParser() antlr.Parser { return s.parser }

func (s *DigitContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(bibcodeParserDIGIT, 0)
}

func (s *DigitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DigitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.EnterDigit(s)
	}
}

func (s *DigitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bibcodeListener); ok {
		listenerT.ExitDigit(s)
	}
}

func (p *bibcodeParser) Digit() (localctx IDigitContext) {
	this := p
	_ = this

	localctx = NewDigitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, bibcodeParserRULE_digit)

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
		p.SetState(94)
		p.Match(bibcodeParserDIGIT)
	}

	return localctx
}
