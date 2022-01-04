// Code generated from STL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package stl // STL
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 50, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 7, 2, 17, 10, 2, 12, 2, 14, 2, 20, 11, 2, 3, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 45, 10, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 2, 2, 45, 2, 14, 3, 2, 2,
	2, 4, 24, 3, 2, 2, 2, 6, 35, 3, 2, 2, 2, 8, 38, 3, 2, 2, 2, 10, 42, 3,
	2, 2, 2, 12, 46, 3, 2, 2, 2, 14, 18, 5, 10, 6, 2, 15, 17, 5, 4, 3, 2, 16,
	15, 3, 2, 2, 2, 17, 20, 3, 2, 2, 2, 18, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2,
	2, 19, 21, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 21, 22, 5, 12, 7, 2, 22, 23,
	7, 2, 2, 3, 23, 3, 3, 2, 2, 2, 24, 25, 7, 3, 2, 2, 25, 26, 7, 4, 2, 2,
	26, 27, 5, 8, 5, 2, 27, 28, 7, 5, 2, 2, 28, 29, 7, 6, 2, 2, 29, 30, 5,
	6, 4, 2, 30, 31, 5, 6, 4, 2, 31, 32, 5, 6, 4, 2, 32, 33, 7, 7, 2, 2, 33,
	34, 7, 8, 2, 2, 34, 5, 3, 2, 2, 2, 35, 36, 7, 9, 2, 2, 36, 37, 5, 8, 5,
	2, 37, 7, 3, 2, 2, 2, 38, 39, 7, 12, 2, 2, 39, 40, 7, 12, 2, 2, 40, 41,
	7, 12, 2, 2, 41, 9, 3, 2, 2, 2, 42, 44, 7, 10, 2, 2, 43, 45, 7, 13, 2,
	2, 44, 43, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 11, 3, 2, 2, 2, 46, 47,
	7, 11, 2, 2, 47, 48, 7, 13, 2, 2, 48, 13, 3, 2, 2, 2, 4, 18, 44,
}
var literalNames = []string{
	"", "'facet'", "'normal'", "'outer'", "'loop'", "'endloop'", "'endfacet'",
	"'vertex'", "'solid'", "'endsolid'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "FLOAT", "IDENTIFIER", "WS",
}

var ruleNames = []string{
	"file_", "triangle", "vertex", "triple", "header", "footer",
}

type STLParser struct {
	*antlr.BaseParser
}

// NewSTLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *STLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSTLParser(input antlr.TokenStream) *STLParser {
	this := new(STLParser)
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
	this.GrammarFileName = "STL.g4"

	return this
}

// STLParser tokens.
const (
	STLParserEOF        = antlr.TokenEOF
	STLParserT__0       = 1
	STLParserT__1       = 2
	STLParserT__2       = 3
	STLParserT__3       = 4
	STLParserT__4       = 5
	STLParserT__5       = 6
	STLParserT__6       = 7
	STLParserT__7       = 8
	STLParserT__8       = 9
	STLParserFLOAT      = 10
	STLParserIDENTIFIER = 11
	STLParserWS         = 12
)

// STLParser rules.
const (
	STLParserRULE_file_    = 0
	STLParserRULE_triangle = 1
	STLParserRULE_vertex   = 2
	STLParserRULE_triple   = 3
	STLParserRULE_header   = 4
	STLParserRULE_footer   = 5
)

// IFile_Context is an interface to support dynamic dispatch.
type IFile_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFile_Context differentiates from other interfaces.
	IsFile_Context()
}

type File_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_Context() *File_Context {
	var p = new(File_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STLParserRULE_file_
	return p
}

func (*File_Context) IsFile_Context() {}

func NewFile_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_Context {
	var p = new(File_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STLParserRULE_file_

	return p
}

func (s *File_Context) GetParser() antlr.Parser { return s.parser }

func (s *File_Context) Header() IHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *File_Context) Footer() IFooterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFooterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFooterContext)
}

func (s *File_Context) EOF() antlr.TerminalNode {
	return s.GetToken(STLParserEOF, 0)
}

func (s *File_Context) AllTriangle() []ITriangleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITriangleContext)(nil)).Elem())
	var tst = make([]ITriangleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITriangleContext)
		}
	}

	return tst
}

func (s *File_Context) Triangle(i int) ITriangleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITriangleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITriangleContext)
}

func (s *File_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.EnterFile_(s)
	}
}

func (s *File_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.ExitFile_(s)
	}
}

func (p *STLParser) File_() (localctx IFile_Context) {
	this := p
	_ = this

	localctx = NewFile_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, STLParserRULE_file_)
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
		p.SetState(12)
		p.Header()
	}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == STLParserT__0 {
		{
			p.SetState(13)
			p.Triangle()
		}

		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(19)
		p.Footer()
	}
	{
		p.SetState(20)
		p.Match(STLParserEOF)
	}

	return localctx
}

// ITriangleContext is an interface to support dynamic dispatch.
type ITriangleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n rule contexts.
	GetN() ITripleContext

	// SetN sets the n rule contexts.
	SetN(ITripleContext)

	// IsTriangleContext differentiates from other interfaces.
	IsTriangleContext()
}

type TriangleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	n      ITripleContext
}

func NewEmptyTriangleContext() *TriangleContext {
	var p = new(TriangleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STLParserRULE_triangle
	return p
}

func (*TriangleContext) IsTriangleContext() {}

func NewTriangleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriangleContext {
	var p = new(TriangleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STLParserRULE_triangle

	return p
}

func (s *TriangleContext) GetParser() antlr.Parser { return s.parser }

func (s *TriangleContext) GetN() ITripleContext { return s.n }

func (s *TriangleContext) SetN(v ITripleContext) { s.n = v }

func (s *TriangleContext) AllVertex() []IVertexContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVertexContext)(nil)).Elem())
	var tst = make([]IVertexContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVertexContext)
		}
	}

	return tst
}

func (s *TriangleContext) Vertex(i int) IVertexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertexContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVertexContext)
}

func (s *TriangleContext) Triple() ITripleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITripleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITripleContext)
}

func (s *TriangleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TriangleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TriangleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.EnterTriangle(s)
	}
}

func (s *TriangleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.ExitTriangle(s)
	}
}

func (p *STLParser) Triangle() (localctx ITriangleContext) {
	this := p
	_ = this

	localctx = NewTriangleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, STLParserRULE_triangle)

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
		p.SetState(22)
		p.Match(STLParserT__0)
	}
	{
		p.SetState(23)
		p.Match(STLParserT__1)
	}
	{
		p.SetState(24)

		var _x = p.Triple()

		localctx.(*TriangleContext).n = _x
	}
	{
		p.SetState(25)
		p.Match(STLParserT__2)
	}
	{
		p.SetState(26)
		p.Match(STLParserT__3)
	}
	{
		p.SetState(27)
		p.Vertex()
	}
	{
		p.SetState(28)
		p.Vertex()
	}
	{
		p.SetState(29)
		p.Vertex()
	}
	{
		p.SetState(30)
		p.Match(STLParserT__4)
	}
	{
		p.SetState(31)
		p.Match(STLParserT__5)
	}

	return localctx
}

// IVertexContext is an interface to support dynamic dispatch.
type IVertexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVertexContext differentiates from other interfaces.
	IsVertexContext()
}

type VertexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVertexContext() *VertexContext {
	var p = new(VertexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STLParserRULE_vertex
	return p
}

func (*VertexContext) IsVertexContext() {}

func NewVertexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VertexContext {
	var p = new(VertexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STLParserRULE_vertex

	return p
}

func (s *VertexContext) GetParser() antlr.Parser { return s.parser }

func (s *VertexContext) Triple() ITripleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITripleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITripleContext)
}

func (s *VertexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VertexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VertexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.EnterVertex(s)
	}
}

func (s *VertexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.ExitVertex(s)
	}
}

func (p *STLParser) Vertex() (localctx IVertexContext) {
	this := p
	_ = this

	localctx = NewVertexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, STLParserRULE_vertex)

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
		p.Match(STLParserT__6)
	}
	{
		p.SetState(34)
		p.Triple()
	}

	return localctx
}

// ITripleContext is an interface to support dynamic dispatch.
type ITripleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetI returns the i token.
	GetI() antlr.Token

	// GetJ returns the j token.
	GetJ() antlr.Token

	// GetK returns the k token.
	GetK() antlr.Token

	// SetI sets the i token.
	SetI(antlr.Token)

	// SetJ sets the j token.
	SetJ(antlr.Token)

	// SetK sets the k token.
	SetK(antlr.Token)

	// IsTripleContext differentiates from other interfaces.
	IsTripleContext()
}

type TripleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	i      antlr.Token
	j      antlr.Token
	k      antlr.Token
}

func NewEmptyTripleContext() *TripleContext {
	var p = new(TripleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STLParserRULE_triple
	return p
}

func (*TripleContext) IsTripleContext() {}

func NewTripleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TripleContext {
	var p = new(TripleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STLParserRULE_triple

	return p
}

func (s *TripleContext) GetParser() antlr.Parser { return s.parser }

func (s *TripleContext) GetI() antlr.Token { return s.i }

func (s *TripleContext) GetJ() antlr.Token { return s.j }

func (s *TripleContext) GetK() antlr.Token { return s.k }

func (s *TripleContext) SetI(v antlr.Token) { s.i = v }

func (s *TripleContext) SetJ(v antlr.Token) { s.j = v }

func (s *TripleContext) SetK(v antlr.Token) { s.k = v }

func (s *TripleContext) AllFLOAT() []antlr.TerminalNode {
	return s.GetTokens(STLParserFLOAT)
}

func (s *TripleContext) FLOAT(i int) antlr.TerminalNode {
	return s.GetToken(STLParserFLOAT, i)
}

func (s *TripleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TripleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TripleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.EnterTriple(s)
	}
}

func (s *TripleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.ExitTriple(s)
	}
}

func (p *STLParser) Triple() (localctx ITripleContext) {
	this := p
	_ = this

	localctx = NewTripleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, STLParserRULE_triple)

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
		p.SetState(36)

		var _m = p.Match(STLParserFLOAT)

		localctx.(*TripleContext).i = _m
	}
	{
		p.SetState(37)

		var _m = p.Match(STLParserFLOAT)

		localctx.(*TripleContext).j = _m
	}
	{
		p.SetState(38)

		var _m = p.Match(STLParserFLOAT)

		localctx.(*TripleContext).k = _m
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STLParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STLParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) GetName() antlr.Token { return s.name }

func (s *HeaderContext) SetName(v antlr.Token) { s.name = v }

func (s *HeaderContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STLParserIDENTIFIER, 0)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *STLParser) Header() (localctx IHeaderContext) {
	this := p
	_ = this

	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, STLParserRULE_header)
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
		p.SetState(40)
		p.Match(STLParserT__7)
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == STLParserIDENTIFIER {
		{
			p.SetState(41)

			var _m = p.Match(STLParserIDENTIFIER)

			localctx.(*HeaderContext).name = _m
		}

	}

	return localctx
}

// IFooterContext is an interface to support dynamic dispatch.
type IFooterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsFooterContext differentiates from other interfaces.
	IsFooterContext()
}

type FooterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyFooterContext() *FooterContext {
	var p = new(FooterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STLParserRULE_footer
	return p
}

func (*FooterContext) IsFooterContext() {}

func NewFooterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FooterContext {
	var p = new(FooterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STLParserRULE_footer

	return p
}

func (s *FooterContext) GetParser() antlr.Parser { return s.parser }

func (s *FooterContext) GetName() antlr.Token { return s.name }

func (s *FooterContext) SetName(v antlr.Token) { s.name = v }

func (s *FooterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(STLParserIDENTIFIER, 0)
}

func (s *FooterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FooterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FooterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.EnterFooter(s)
	}
}

func (s *FooterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STLListener); ok {
		listenerT.ExitFooter(s)
	}
}

func (p *STLParser) Footer() (localctx IFooterContext) {
	this := p
	_ = this

	localctx = NewFooterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, STLParserRULE_footer)

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
		p.SetState(44)
		p.Match(STLParserT__8)
	}
	{
		p.SetState(45)

		var _m = p.Match(STLParserIDENTIFIER)

		localctx.(*FooterContext).name = _m
	}

	return localctx
}
