// Code generated from tsv.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tsv // tsv
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 32, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 7, 2, 13, 10,
	2, 12, 2, 14, 2, 16, 11, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 23, 10,
	4, 12, 4, 14, 4, 26, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4,
	6, 8, 2, 3, 3, 2, 5, 6, 2, 29, 2, 10, 3, 2, 2, 2, 4, 17, 3, 2, 2, 2, 6,
	19, 3, 2, 2, 2, 8, 29, 3, 2, 2, 2, 10, 14, 5, 4, 3, 2, 11, 13, 5, 6, 4,
	2, 12, 11, 3, 2, 2, 2, 13, 16, 3, 2, 2, 2, 14, 12, 3, 2, 2, 2, 14, 15,
	3, 2, 2, 2, 15, 3, 3, 2, 2, 2, 16, 14, 3, 2, 2, 2, 17, 18, 5, 6, 4, 2,
	18, 5, 3, 2, 2, 2, 19, 24, 5, 8, 5, 2, 20, 21, 7, 3, 2, 2, 21, 23, 5, 8,
	5, 2, 22, 20, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2, 24, 25,
	3, 2, 2, 2, 25, 27, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 27, 28, 7, 4, 2, 2,
	28, 7, 3, 2, 2, 2, 29, 30, 9, 2, 2, 2, 30, 9, 3, 2, 2, 2, 4, 14, 24,
}
var literalNames = []string{
	"", "'\t'",
}
var symbolicNames = []string{
	"", "TAB", "EOL", "TEXT", "STRING",
}

var ruleNames = []string{
	"tsvFile", "hdr", "row", "field",
}

type tsvParser struct {
	*antlr.BaseParser
}

// NewtsvParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *tsvParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewtsvParser(input antlr.TokenStream) *tsvParser {
	this := new(tsvParser)
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
	this.GrammarFileName = "tsv.g4"

	return this
}

// tsvParser tokens.
const (
	tsvParserEOF    = antlr.TokenEOF
	tsvParserTAB    = 1
	tsvParserEOL    = 2
	tsvParserTEXT   = 3
	tsvParserSTRING = 4
)

// tsvParser rules.
const (
	tsvParserRULE_tsvFile = 0
	tsvParserRULE_hdr     = 1
	tsvParserRULE_row     = 2
	tsvParserRULE_field   = 3
)

// ITsvFileContext is an interface to support dynamic dispatch.
type ITsvFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTsvFileContext differentiates from other interfaces.
	IsTsvFileContext()
}

type TsvFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTsvFileContext() *TsvFileContext {
	var p = new(TsvFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tsvParserRULE_tsvFile
	return p
}

func (*TsvFileContext) IsTsvFileContext() {}

func NewTsvFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TsvFileContext {
	var p = new(TsvFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tsvParserRULE_tsvFile

	return p
}

func (s *TsvFileContext) GetParser() antlr.Parser { return s.parser }

func (s *TsvFileContext) Hdr() IHdrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHdrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHdrContext)
}

func (s *TsvFileContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *TsvFileContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *TsvFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TsvFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TsvFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tsvListener); ok {
		listenerT.EnterTsvFile(s)
	}
}

func (s *TsvFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tsvListener); ok {
		listenerT.ExitTsvFile(s)
	}
}

func (p *tsvParser) TsvFile() (localctx ITsvFileContext) {
	this := p
	_ = this

	localctx = NewTsvFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tsvParserRULE_tsvFile)
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
		p.SetState(8)
		p.Hdr()
	}
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tsvParserTEXT || _la == tsvParserSTRING {
		{
			p.SetState(9)
			p.Row()
		}

		p.SetState(14)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHdrContext is an interface to support dynamic dispatch.
type IHdrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHdrContext differentiates from other interfaces.
	IsHdrContext()
}

type HdrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHdrContext() *HdrContext {
	var p = new(HdrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tsvParserRULE_hdr
	return p
}

func (*HdrContext) IsHdrContext() {}

func NewHdrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HdrContext {
	var p = new(HdrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tsvParserRULE_hdr

	return p
}

func (s *HdrContext) GetParser() antlr.Parser { return s.parser }

func (s *HdrContext) Row() IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *HdrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HdrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HdrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tsvListener); ok {
		listenerT.EnterHdr(s)
	}
}

func (s *HdrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tsvListener); ok {
		listenerT.ExitHdr(s)
	}
}

func (p *tsvParser) Hdr() (localctx IHdrContext) {
	this := p
	_ = this

	localctx = NewHdrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tsvParserRULE_hdr)

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
		p.SetState(15)
		p.Row()
	}

	return localctx
}

// IRowContext is an interface to support dynamic dispatch.
type IRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowContext differentiates from other interfaces.
	IsRowContext()
}

type RowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowContext() *RowContext {
	var p = new(RowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tsvParserRULE_row
	return p
}

func (*RowContext) IsRowContext() {}

func NewRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowContext {
	var p = new(RowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tsvParserRULE_row

	return p
}

func (s *RowContext) GetParser() antlr.Parser { return s.parser }

func (s *RowContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *RowContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *RowContext) EOL() antlr.TerminalNode {
	return s.GetToken(tsvParserEOL, 0)
}

func (s *RowContext) AllTAB() []antlr.TerminalNode {
	return s.GetTokens(tsvParserTAB)
}

func (s *RowContext) TAB(i int) antlr.TerminalNode {
	return s.GetToken(tsvParserTAB, i)
}

func (s *RowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tsvListener); ok {
		listenerT.EnterRow(s)
	}
}

func (s *RowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tsvListener); ok {
		listenerT.ExitRow(s)
	}
}

func (p *tsvParser) Row() (localctx IRowContext) {
	this := p
	_ = this

	localctx = NewRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tsvParserRULE_row)
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
		p.SetState(17)
		p.Field()
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tsvParserTAB {
		{
			p.SetState(18)
			p.Match(tsvParserTAB)
		}
		{
			p.SetState(19)
			p.Field()
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.Match(tsvParserEOL)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tsvParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tsvParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) TEXT() antlr.TerminalNode {
	return s.GetToken(tsvParserTEXT, 0)
}

func (s *FieldContext) STRING() antlr.TerminalNode {
	return s.GetToken(tsvParserSTRING, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tsvListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tsvListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *tsvParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tsvParserRULE_field)
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
		p.SetState(27)
		_la = p.GetTokenStream().LA(1)

		if !(_la == tsvParserTEXT || _la == tsvParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
