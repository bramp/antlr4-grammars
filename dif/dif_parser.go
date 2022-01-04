// Code generated from dif.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dif // dif
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 75, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 8, 6, 8, 53, 10, 8, 13, 8, 14, 8, 54, 3, 9, 3, 9, 3, 9, 5,
	9, 60, 10, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 2, 3, 3, 2, 7, 8, 2, 65, 2, 26, 3, 2, 2, 2, 4, 30,
	3, 2, 2, 2, 6, 35, 3, 2, 2, 2, 8, 39, 3, 2, 2, 2, 10, 43, 3, 2, 2, 2, 12,
	47, 3, 2, 2, 2, 14, 52, 3, 2, 2, 2, 16, 59, 3, 2, 2, 2, 18, 61, 3, 2, 2,
	2, 20, 64, 3, 2, 2, 2, 22, 67, 3, 2, 2, 2, 24, 70, 3, 2, 2, 2, 26, 27,
	5, 4, 3, 2, 27, 28, 5, 14, 8, 2, 28, 29, 7, 2, 2, 3, 29, 3, 3, 2, 2, 2,
	30, 31, 5, 6, 4, 2, 31, 32, 5, 8, 5, 2, 32, 33, 5, 10, 6, 2, 33, 34, 5,
	12, 7, 2, 34, 5, 3, 2, 2, 2, 35, 36, 7, 3, 2, 2, 36, 37, 5, 24, 13, 2,
	37, 38, 7, 12, 2, 2, 38, 7, 3, 2, 2, 2, 39, 40, 7, 4, 2, 2, 40, 41, 5,
	24, 13, 2, 41, 42, 7, 12, 2, 2, 42, 9, 3, 2, 2, 2, 43, 44, 7, 5, 2, 2,
	44, 45, 5, 24, 13, 2, 45, 46, 7, 12, 2, 2, 46, 11, 3, 2, 2, 2, 47, 48,
	7, 6, 2, 2, 48, 49, 5, 24, 13, 2, 49, 50, 7, 12, 2, 2, 50, 13, 3, 2, 2,
	2, 51, 53, 5, 16, 9, 2, 52, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 52,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 15, 3, 2, 2, 2, 56, 60, 5, 18, 10,
	2, 57, 60, 5, 20, 11, 2, 58, 60, 5, 22, 12, 2, 59, 56, 3, 2, 2, 2, 59,
	57, 3, 2, 2, 2, 59, 58, 3, 2, 2, 2, 60, 17, 3, 2, 2, 2, 61, 62, 5, 24,
	13, 2, 62, 63, 9, 2, 2, 2, 63, 19, 3, 2, 2, 2, 64, 65, 5, 24, 13, 2, 65,
	66, 7, 12, 2, 2, 66, 21, 3, 2, 2, 2, 67, 68, 5, 24, 13, 2, 68, 69, 7, 9,
	2, 2, 69, 23, 3, 2, 2, 2, 70, 71, 7, 11, 2, 2, 71, 72, 7, 10, 2, 2, 72,
	73, 7, 11, 2, 2, 73, 25, 3, 2, 2, 2, 4, 54, 59,
}
var literalNames = []string{
	"", "'TABLE'", "'VECTORS'", "'TUPLES'", "'DATA'", "'BOT'", "'EOD'", "'V'",
	"','",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "NUM", "STRING", "WS",
}

var ruleNames = []string{
	"dif", "header", "tableheader", "vectorsheader", "tuplesheader", "dataheader",
	"data", "datum", "directive", "string_", "numeric", "pair",
}

type difParser struct {
	*antlr.BaseParser
}

// NewdifParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *difParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdifParser(input antlr.TokenStream) *difParser {
	this := new(difParser)
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
	this.GrammarFileName = "dif.g4"

	return this
}

// difParser tokens.
const (
	difParserEOF    = antlr.TokenEOF
	difParserT__0   = 1
	difParserT__1   = 2
	difParserT__2   = 3
	difParserT__3   = 4
	difParserT__4   = 5
	difParserT__5   = 6
	difParserT__6   = 7
	difParserT__7   = 8
	difParserNUM    = 9
	difParserSTRING = 10
	difParserWS     = 11
)

// difParser rules.
const (
	difParserRULE_dif           = 0
	difParserRULE_header        = 1
	difParserRULE_tableheader   = 2
	difParserRULE_vectorsheader = 3
	difParserRULE_tuplesheader  = 4
	difParserRULE_dataheader    = 5
	difParserRULE_data          = 6
	difParserRULE_datum         = 7
	difParserRULE_directive     = 8
	difParserRULE_string_       = 9
	difParserRULE_numeric       = 10
	difParserRULE_pair          = 11
)

// IDifContext is an interface to support dynamic dispatch.
type IDifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDifContext differentiates from other interfaces.
	IsDifContext()
}

type DifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDifContext() *DifContext {
	var p = new(DifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_dif
	return p
}

func (*DifContext) IsDifContext() {}

func NewDifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DifContext {
	var p = new(DifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_dif

	return p
}

func (s *DifContext) GetParser() antlr.Parser { return s.parser }

func (s *DifContext) Header() IHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *DifContext) Data() IDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *DifContext) EOF() antlr.TerminalNode {
	return s.GetToken(difParserEOF, 0)
}

func (s *DifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterDif(s)
	}
}

func (s *DifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitDif(s)
	}
}

func (p *difParser) Dif() (localctx IDifContext) {
	this := p
	_ = this

	localctx = NewDifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, difParserRULE_dif)

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
		p.SetState(24)
		p.Header()
	}
	{
		p.SetState(25)
		p.Data()
	}
	{
		p.SetState(26)
		p.Match(difParserEOF)
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) Tableheader() ITableheaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableheaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableheaderContext)
}

func (s *HeaderContext) Vectorsheader() IVectorsheaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVectorsheaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVectorsheaderContext)
}

func (s *HeaderContext) Tuplesheader() ITuplesheaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITuplesheaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITuplesheaderContext)
}

func (s *HeaderContext) Dataheader() IDataheaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataheaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataheaderContext)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *difParser) Header() (localctx IHeaderContext) {
	this := p
	_ = this

	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, difParserRULE_header)

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
		p.SetState(28)
		p.Tableheader()
	}
	{
		p.SetState(29)
		p.Vectorsheader()
	}
	{
		p.SetState(30)
		p.Tuplesheader()
	}
	{
		p.SetState(31)
		p.Dataheader()
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
	p.RuleIndex = difParserRULE_tableheader
	return p
}

func (*TableheaderContext) IsTableheaderContext() {}

func NewTableheaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableheaderContext {
	var p = new(TableheaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_tableheader

	return p
}

func (s *TableheaderContext) GetParser() antlr.Parser { return s.parser }

func (s *TableheaderContext) Pair() IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *TableheaderContext) STRING() antlr.TerminalNode {
	return s.GetToken(difParserSTRING, 0)
}

func (s *TableheaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableheaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableheaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterTableheader(s)
	}
}

func (s *TableheaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitTableheader(s)
	}
}

func (p *difParser) Tableheader() (localctx ITableheaderContext) {
	this := p
	_ = this

	localctx = NewTableheaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, difParserRULE_tableheader)

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
		p.SetState(33)
		p.Match(difParserT__0)
	}
	{
		p.SetState(34)
		p.Pair()
	}
	{
		p.SetState(35)
		p.Match(difParserSTRING)
	}

	return localctx
}

// IVectorsheaderContext is an interface to support dynamic dispatch.
type IVectorsheaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVectorsheaderContext differentiates from other interfaces.
	IsVectorsheaderContext()
}

type VectorsheaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorsheaderContext() *VectorsheaderContext {
	var p = new(VectorsheaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_vectorsheader
	return p
}

func (*VectorsheaderContext) IsVectorsheaderContext() {}

func NewVectorsheaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorsheaderContext {
	var p = new(VectorsheaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_vectorsheader

	return p
}

func (s *VectorsheaderContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorsheaderContext) Pair() IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *VectorsheaderContext) STRING() antlr.TerminalNode {
	return s.GetToken(difParserSTRING, 0)
}

func (s *VectorsheaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorsheaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorsheaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterVectorsheader(s)
	}
}

func (s *VectorsheaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitVectorsheader(s)
	}
}

func (p *difParser) Vectorsheader() (localctx IVectorsheaderContext) {
	this := p
	_ = this

	localctx = NewVectorsheaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, difParserRULE_vectorsheader)

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
		p.SetState(37)
		p.Match(difParserT__1)
	}
	{
		p.SetState(38)
		p.Pair()
	}
	{
		p.SetState(39)
		p.Match(difParserSTRING)
	}

	return localctx
}

// ITuplesheaderContext is an interface to support dynamic dispatch.
type ITuplesheaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTuplesheaderContext differentiates from other interfaces.
	IsTuplesheaderContext()
}

type TuplesheaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTuplesheaderContext() *TuplesheaderContext {
	var p = new(TuplesheaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_tuplesheader
	return p
}

func (*TuplesheaderContext) IsTuplesheaderContext() {}

func NewTuplesheaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TuplesheaderContext {
	var p = new(TuplesheaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_tuplesheader

	return p
}

func (s *TuplesheaderContext) GetParser() antlr.Parser { return s.parser }

func (s *TuplesheaderContext) Pair() IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *TuplesheaderContext) STRING() antlr.TerminalNode {
	return s.GetToken(difParserSTRING, 0)
}

func (s *TuplesheaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TuplesheaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TuplesheaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterTuplesheader(s)
	}
}

func (s *TuplesheaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitTuplesheader(s)
	}
}

func (p *difParser) Tuplesheader() (localctx ITuplesheaderContext) {
	this := p
	_ = this

	localctx = NewTuplesheaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, difParserRULE_tuplesheader)

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
		p.SetState(41)
		p.Match(difParserT__2)
	}
	{
		p.SetState(42)
		p.Pair()
	}
	{
		p.SetState(43)
		p.Match(difParserSTRING)
	}

	return localctx
}

// IDataheaderContext is an interface to support dynamic dispatch.
type IDataheaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataheaderContext differentiates from other interfaces.
	IsDataheaderContext()
}

type DataheaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataheaderContext() *DataheaderContext {
	var p = new(DataheaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_dataheader
	return p
}

func (*DataheaderContext) IsDataheaderContext() {}

func NewDataheaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataheaderContext {
	var p = new(DataheaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_dataheader

	return p
}

func (s *DataheaderContext) GetParser() antlr.Parser { return s.parser }

func (s *DataheaderContext) Pair() IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *DataheaderContext) STRING() antlr.TerminalNode {
	return s.GetToken(difParserSTRING, 0)
}

func (s *DataheaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataheaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataheaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterDataheader(s)
	}
}

func (s *DataheaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitDataheader(s)
	}
}

func (p *difParser) Dataheader() (localctx IDataheaderContext) {
	this := p
	_ = this

	localctx = NewDataheaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, difParserRULE_dataheader)

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
		p.SetState(45)
		p.Match(difParserT__3)
	}
	{
		p.SetState(46)
		p.Pair()
	}
	{
		p.SetState(47)
		p.Match(difParserSTRING)
	}

	return localctx
}

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_data
	return p
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) AllDatum() []IDatumContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatumContext)(nil)).Elem())
	var tst = make([]IDatumContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatumContext)
		}
	}

	return tst
}

func (s *DataContext) Datum(i int) IDatumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatumContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatumContext)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterData(s)
	}
}

func (s *DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitData(s)
	}
}

func (p *difParser) Data() (localctx IDataContext) {
	this := p
	_ = this

	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, difParserRULE_data)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == difParserNUM {
		{
			p.SetState(49)
			p.Datum()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDatumContext is an interface to support dynamic dispatch.
type IDatumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatumContext differentiates from other interfaces.
	IsDatumContext()
}

type DatumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatumContext() *DatumContext {
	var p = new(DatumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_datum
	return p
}

func (*DatumContext) IsDatumContext() {}

func NewDatumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatumContext {
	var p = new(DatumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_datum

	return p
}

func (s *DatumContext) GetParser() antlr.Parser { return s.parser }

func (s *DatumContext) Directive() IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *DatumContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *DatumContext) Numeric() INumericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericContext)
}

func (s *DatumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterDatum(s)
	}
}

func (s *DatumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitDatum(s)
	}
}

func (p *difParser) Datum() (localctx IDatumContext) {
	this := p
	_ = this

	localctx = NewDatumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, difParserRULE_datum)

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

	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Directive()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.String_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Numeric()
		}

	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Pair() IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *difParser) Directive() (localctx IDirectiveContext) {
	this := p
	_ = this

	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, difParserRULE_directive)
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
		p.SetState(59)
		p.Pair()
	}
	{
		p.SetState(60)
		_la = p.GetTokenStream().LA(1)

		if !(_la == difParserT__4 || _la == difParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) Pair() IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *String_Context) STRING() antlr.TerminalNode {
	return s.GetToken(difParserSTRING, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitString_(s)
	}
}

func (p *difParser) String_() (localctx IString_Context) {
	this := p
	_ = this

	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, difParserRULE_string_)

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
		p.SetState(62)
		p.Pair()
	}
	{
		p.SetState(63)
		p.Match(difParserSTRING)
	}

	return localctx
}

// INumericContext is an interface to support dynamic dispatch.
type INumericContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericContext differentiates from other interfaces.
	IsNumericContext()
}

type NumericContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericContext() *NumericContext {
	var p = new(NumericContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_numeric
	return p
}

func (*NumericContext) IsNumericContext() {}

func NewNumericContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericContext {
	var p = new(NumericContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_numeric

	return p
}

func (s *NumericContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericContext) Pair() IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *NumericContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterNumeric(s)
	}
}

func (s *NumericContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitNumeric(s)
	}
}

func (p *difParser) Numeric() (localctx INumericContext) {
	this := p
	_ = this

	localctx = NewNumericContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, difParserRULE_numeric)

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
		p.SetState(65)
		p.Pair()
	}
	{
		p.SetState(66)
		p.Match(difParserT__6)
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = difParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = difParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(difParserNUM)
}

func (s *PairContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(difParserNUM, i)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(difListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *difParser) Pair() (localctx IPairContext) {
	this := p
	_ = this

	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, difParserRULE_pair)

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
		p.SetState(68)
		p.Match(difParserNUM)
	}
	{
		p.SetState(69)
		p.Match(difParserT__7)
	}
	{
		p.SetState(70)
		p.Match(difParserNUM)
	}

	return localctx
}
