// Code generated from gff3.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gff3 // gff3
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 92, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 6, 2, 33, 10, 2, 13, 2, 14,
	2, 34, 3, 3, 3, 3, 5, 3, 39, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	58, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 7, 5, 65, 10, 5, 12, 5, 14, 5,
	68, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 15, 2, 2, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 2, 2, 2, 81, 2, 30, 3, 2, 2, 2, 4, 38, 3, 2, 2, 2, 6, 40, 3, 2,
	2, 2, 8, 61, 3, 2, 2, 2, 10, 69, 3, 2, 2, 2, 12, 73, 3, 2, 2, 2, 14, 75,
	3, 2, 2, 2, 16, 77, 3, 2, 2, 2, 18, 79, 3, 2, 2, 2, 20, 81, 3, 2, 2, 2,
	22, 83, 3, 2, 2, 2, 24, 85, 3, 2, 2, 2, 26, 87, 3, 2, 2, 2, 28, 89, 3,
	2, 2, 2, 30, 32, 7, 6, 2, 2, 31, 33, 5, 4, 3, 2, 32, 31, 3, 2, 2, 2, 33,
	34, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 3, 3, 2, 2,
	2, 36, 39, 5, 28, 15, 2, 37, 39, 5, 6, 4, 2, 38, 36, 3, 2, 2, 2, 38, 37,
	3, 2, 2, 2, 39, 5, 3, 2, 2, 2, 40, 41, 5, 12, 7, 2, 41, 42, 7, 3, 2, 2,
	42, 43, 5, 14, 8, 2, 43, 44, 7, 3, 2, 2, 44, 45, 5, 16, 9, 2, 45, 46, 7,
	3, 2, 2, 46, 47, 5, 18, 10, 2, 47, 48, 7, 3, 2, 2, 48, 49, 5, 20, 11, 2,
	49, 50, 7, 3, 2, 2, 50, 51, 5, 24, 13, 2, 51, 52, 7, 3, 2, 2, 52, 53, 5,
	22, 12, 2, 53, 54, 7, 3, 2, 2, 54, 55, 5, 26, 14, 2, 55, 57, 7, 3, 2, 2,
	56, 58, 5, 8, 5, 2, 57, 56, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 3,
	2, 2, 2, 59, 60, 7, 8, 2, 2, 60, 7, 3, 2, 2, 2, 61, 66, 5, 10, 6, 2, 62,
	63, 7, 4, 2, 2, 63, 65, 5, 10, 6, 2, 64, 62, 3, 2, 2, 2, 65, 68, 3, 2,
	2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 9, 3, 2, 2, 2, 68, 66,
	3, 2, 2, 2, 69, 70, 7, 9, 2, 2, 70, 71, 7, 5, 2, 2, 71, 72, 7, 9, 2, 2,
	72, 11, 3, 2, 2, 2, 73, 74, 7, 9, 2, 2, 74, 13, 3, 2, 2, 2, 75, 76, 7,
	9, 2, 2, 76, 15, 3, 2, 2, 2, 77, 78, 7, 9, 2, 2, 78, 17, 3, 2, 2, 2, 79,
	80, 7, 9, 2, 2, 80, 19, 3, 2, 2, 2, 81, 82, 7, 9, 2, 2, 82, 21, 3, 2, 2,
	2, 83, 84, 7, 9, 2, 2, 84, 23, 3, 2, 2, 2, 85, 86, 7, 9, 2, 2, 86, 25,
	3, 2, 2, 2, 87, 88, 7, 9, 2, 2, 88, 27, 3, 2, 2, 2, 89, 90, 7, 7, 2, 2,
	90, 29, 3, 2, 2, 2, 6, 34, 38, 57, 66,
}
var literalNames = []string{
	"", "'\t'", "';'", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "HEADER", "COMMENTLINE", "EOL", "TEXT",
}

var ruleNames = []string{
	"document", "line", "dataline", "attributes", "attribute", "seqid", "source",
	"type_", "start", "end", "strand", "score", "phase", "commentline",
}

type gff3Parser struct {
	*antlr.BaseParser
}

// Newgff3Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *gff3Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newgff3Parser(input antlr.TokenStream) *gff3Parser {
	this := new(gff3Parser)
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
	this.GrammarFileName = "gff3.g4"

	return this
}

// gff3Parser tokens.
const (
	gff3ParserEOF         = antlr.TokenEOF
	gff3ParserT__0        = 1
	gff3ParserT__1        = 2
	gff3ParserT__2        = 3
	gff3ParserHEADER      = 4
	gff3ParserCOMMENTLINE = 5
	gff3ParserEOL         = 6
	gff3ParserTEXT        = 7
)

// gff3Parser rules.
const (
	gff3ParserRULE_document    = 0
	gff3ParserRULE_line        = 1
	gff3ParserRULE_dataline    = 2
	gff3ParserRULE_attributes  = 3
	gff3ParserRULE_attribute   = 4
	gff3ParserRULE_seqid       = 5
	gff3ParserRULE_source      = 6
	gff3ParserRULE_type_       = 7
	gff3ParserRULE_start       = 8
	gff3ParserRULE_end         = 9
	gff3ParserRULE_strand      = 10
	gff3ParserRULE_score       = 11
	gff3ParserRULE_phase       = 12
	gff3ParserRULE_commentline = 13
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
	p.RuleIndex = gff3ParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) HEADER() antlr.TerminalNode {
	return s.GetToken(gff3ParserHEADER, 0)
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
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *gff3Parser) Document() (localctx IDocumentContext) {
	this := p
	_ = this

	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gff3ParserRULE_document)
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
		p.SetState(28)
		p.Match(gff3ParserHEADER)
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gff3ParserCOMMENTLINE || _la == gff3ParserTEXT {
		{
			p.SetState(29)
			p.Line()
		}

		p.SetState(32)
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
	p.RuleIndex = gff3ParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Commentline() ICommentlineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentlineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentlineContext)
}

func (s *LineContext) Dataline() IDatalineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatalineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatalineContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *gff3Parser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gff3ParserRULE_line)

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

	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gff3ParserCOMMENTLINE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Commentline()
		}

	case gff3ParserTEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Dataline()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDatalineContext is an interface to support dynamic dispatch.
type IDatalineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatalineContext differentiates from other interfaces.
	IsDatalineContext()
}

type DatalineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatalineContext() *DatalineContext {
	var p = new(DatalineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_dataline
	return p
}

func (*DatalineContext) IsDatalineContext() {}

func NewDatalineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatalineContext {
	var p = new(DatalineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_dataline

	return p
}

func (s *DatalineContext) GetParser() antlr.Parser { return s.parser }

func (s *DatalineContext) Seqid() ISeqidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeqidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeqidContext)
}

func (s *DatalineContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *DatalineContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *DatalineContext) Start() IStartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStartContext)
}

func (s *DatalineContext) End() IEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndContext)
}

func (s *DatalineContext) Score() IScoreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScoreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScoreContext)
}

func (s *DatalineContext) Strand() IStrandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrandContext)
}

func (s *DatalineContext) Phase() IPhaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPhaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPhaseContext)
}

func (s *DatalineContext) EOL() antlr.TerminalNode {
	return s.GetToken(gff3ParserEOL, 0)
}

func (s *DatalineContext) Attributes() IAttributesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributesContext)
}

func (s *DatalineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatalineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatalineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterDataline(s)
	}
}

func (s *DatalineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitDataline(s)
	}
}

func (p *gff3Parser) Dataline() (localctx IDatalineContext) {
	this := p
	_ = this

	localctx = NewDatalineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gff3ParserRULE_dataline)
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
		p.SetState(38)
		p.Seqid()
	}
	{
		p.SetState(39)
		p.Match(gff3ParserT__0)
	}
	{
		p.SetState(40)
		p.Source()
	}
	{
		p.SetState(41)
		p.Match(gff3ParserT__0)
	}
	{
		p.SetState(42)
		p.Type_()
	}
	{
		p.SetState(43)
		p.Match(gff3ParserT__0)
	}
	{
		p.SetState(44)
		p.Start()
	}
	{
		p.SetState(45)
		p.Match(gff3ParserT__0)
	}
	{
		p.SetState(46)
		p.End()
	}
	{
		p.SetState(47)
		p.Match(gff3ParserT__0)
	}
	{
		p.SetState(48)
		p.Score()
	}
	{
		p.SetState(49)
		p.Match(gff3ParserT__0)
	}
	{
		p.SetState(50)
		p.Strand()
	}
	{
		p.SetState(51)
		p.Match(gff3ParserT__0)
	}
	{
		p.SetState(52)
		p.Phase()
	}
	{
		p.SetState(53)
		p.Match(gff3ParserT__0)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gff3ParserTEXT {
		{
			p.SetState(54)
			p.Attributes()
		}

	}
	{
		p.SetState(57)
		p.Match(gff3ParserEOL)
	}

	return localctx
}

// IAttributesContext is an interface to support dynamic dispatch.
type IAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributesContext differentiates from other interfaces.
	IsAttributesContext()
}

type AttributesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributesContext() *AttributesContext {
	var p = new(AttributesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_attributes
	return p
}

func (*AttributesContext) IsAttributesContext() {}

func NewAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributesContext {
	var p = new(AttributesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_attributes

	return p
}

func (s *AttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributesContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *AttributesContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *AttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterAttributes(s)
	}
}

func (s *AttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitAttributes(s)
	}
}

func (p *gff3Parser) Attributes() (localctx IAttributesContext) {
	this := p
	_ = this

	localctx = NewAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gff3ParserRULE_attributes)
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
		p.Attribute()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gff3ParserT__1 {
		{
			p.SetState(60)
			p.Match(gff3ParserT__1)
		}
		{
			p.SetState(61)
			p.Attribute()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = gff3ParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(gff3ParserTEXT)
}

func (s *AttributeContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(gff3ParserTEXT, i)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (p *gff3Parser) Attribute() (localctx IAttributeContext) {
	this := p
	_ = this

	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gff3ParserRULE_attribute)

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
		p.SetState(67)
		p.Match(gff3ParserTEXT)
	}
	{
		p.SetState(68)
		p.Match(gff3ParserT__2)
	}
	{
		p.SetState(69)
		p.Match(gff3ParserTEXT)
	}

	return localctx
}

// ISeqidContext is an interface to support dynamic dispatch.
type ISeqidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSeqidContext differentiates from other interfaces.
	IsSeqidContext()
}

type SeqidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeqidContext() *SeqidContext {
	var p = new(SeqidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_seqid
	return p
}

func (*SeqidContext) IsSeqidContext() {}

func NewSeqidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeqidContext {
	var p = new(SeqidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_seqid

	return p
}

func (s *SeqidContext) GetParser() antlr.Parser { return s.parser }

func (s *SeqidContext) TEXT() antlr.TerminalNode {
	return s.GetToken(gff3ParserTEXT, 0)
}

func (s *SeqidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeqidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeqidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterSeqid(s)
	}
}

func (s *SeqidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitSeqid(s)
	}
}

func (p *gff3Parser) Seqid() (localctx ISeqidContext) {
	this := p
	_ = this

	localctx = NewSeqidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gff3ParserRULE_seqid)

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
		p.SetState(71)
		p.Match(gff3ParserTEXT)
	}

	return localctx
}

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	var p = new(SourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_source
	return p
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	var p = new(SourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) TEXT() antlr.TerminalNode {
	return s.GetToken(gff3ParserTEXT, 0)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterSource(s)
	}
}

func (s *SourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitSource(s)
	}
}

func (p *gff3Parser) Source() (localctx ISourceContext) {
	this := p
	_ = this

	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gff3ParserRULE_source)

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
		p.SetState(73)
		p.Match(gff3ParserTEXT)
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) TEXT() antlr.TerminalNode {
	return s.GetToken(gff3ParserTEXT, 0)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *gff3Parser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gff3ParserRULE_type_)

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
		p.SetState(75)
		p.Match(gff3ParserTEXT)
	}

	return localctx
}

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) TEXT() antlr.TerminalNode {
	return s.GetToken(gff3ParserTEXT, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *gff3Parser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gff3ParserRULE_start)

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
		p.SetState(77)
		p.Match(gff3ParserTEXT)
	}

	return localctx
}

// IEndContext is an interface to support dynamic dispatch.
type IEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndContext differentiates from other interfaces.
	IsEndContext()
}

type EndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndContext() *EndContext {
	var p = new(EndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_end
	return p
}

func (*EndContext) IsEndContext() {}

func NewEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndContext {
	var p = new(EndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_end

	return p
}

func (s *EndContext) GetParser() antlr.Parser { return s.parser }

func (s *EndContext) TEXT() antlr.TerminalNode {
	return s.GetToken(gff3ParserTEXT, 0)
}

func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterEnd(s)
	}
}

func (s *EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitEnd(s)
	}
}

func (p *gff3Parser) End() (localctx IEndContext) {
	this := p
	_ = this

	localctx = NewEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, gff3ParserRULE_end)

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
		p.SetState(79)
		p.Match(gff3ParserTEXT)
	}

	return localctx
}

// IStrandContext is an interface to support dynamic dispatch.
type IStrandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrandContext differentiates from other interfaces.
	IsStrandContext()
}

type StrandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrandContext() *StrandContext {
	var p = new(StrandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_strand
	return p
}

func (*StrandContext) IsStrandContext() {}

func NewStrandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrandContext {
	var p = new(StrandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_strand

	return p
}

func (s *StrandContext) GetParser() antlr.Parser { return s.parser }

func (s *StrandContext) TEXT() antlr.TerminalNode {
	return s.GetToken(gff3ParserTEXT, 0)
}

func (s *StrandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterStrand(s)
	}
}

func (s *StrandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitStrand(s)
	}
}

func (p *gff3Parser) Strand() (localctx IStrandContext) {
	this := p
	_ = this

	localctx = NewStrandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, gff3ParserRULE_strand)

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
		p.SetState(81)
		p.Match(gff3ParserTEXT)
	}

	return localctx
}

// IScoreContext is an interface to support dynamic dispatch.
type IScoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScoreContext differentiates from other interfaces.
	IsScoreContext()
}

type ScoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScoreContext() *ScoreContext {
	var p = new(ScoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_score
	return p
}

func (*ScoreContext) IsScoreContext() {}

func NewScoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScoreContext {
	var p = new(ScoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_score

	return p
}

func (s *ScoreContext) GetParser() antlr.Parser { return s.parser }

func (s *ScoreContext) TEXT() antlr.TerminalNode {
	return s.GetToken(gff3ParserTEXT, 0)
}

func (s *ScoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterScore(s)
	}
}

func (s *ScoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitScore(s)
	}
}

func (p *gff3Parser) Score() (localctx IScoreContext) {
	this := p
	_ = this

	localctx = NewScoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gff3ParserRULE_score)

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
		p.SetState(83)
		p.Match(gff3ParserTEXT)
	}

	return localctx
}

// IPhaseContext is an interface to support dynamic dispatch.
type IPhaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhaseContext differentiates from other interfaces.
	IsPhaseContext()
}

type PhaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhaseContext() *PhaseContext {
	var p = new(PhaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_phase
	return p
}

func (*PhaseContext) IsPhaseContext() {}

func NewPhaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhaseContext {
	var p = new(PhaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_phase

	return p
}

func (s *PhaseContext) GetParser() antlr.Parser { return s.parser }

func (s *PhaseContext) TEXT() antlr.TerminalNode {
	return s.GetToken(gff3ParserTEXT, 0)
}

func (s *PhaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterPhase(s)
	}
}

func (s *PhaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitPhase(s)
	}
}

func (p *gff3Parser) Phase() (localctx IPhaseContext) {
	this := p
	_ = this

	localctx = NewPhaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gff3ParserRULE_phase)

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
		p.Match(gff3ParserTEXT)
	}

	return localctx
}

// ICommentlineContext is an interface to support dynamic dispatch.
type ICommentlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentlineContext differentiates from other interfaces.
	IsCommentlineContext()
}

type CommentlineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentlineContext() *CommentlineContext {
	var p = new(CommentlineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gff3ParserRULE_commentline
	return p
}

func (*CommentlineContext) IsCommentlineContext() {}

func NewCommentlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentlineContext {
	var p = new(CommentlineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gff3ParserRULE_commentline

	return p
}

func (s *CommentlineContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentlineContext) COMMENTLINE() antlr.TerminalNode {
	return s.GetToken(gff3ParserCOMMENTLINE, 0)
}

func (s *CommentlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.EnterCommentline(s)
	}
}

func (s *CommentlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gff3Listener); ok {
		listenerT.ExitCommentline(s)
	}
}

func (p *gff3Parser) Commentline() (localctx ICommentlineContext) {
	this := p
	_ = this

	localctx = NewCommentlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gff3ParserRULE_commentline)

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
		p.Match(gff3ParserCOMMENTLINE)
	}

	return localctx
}
