// Code generated from useragent.g4 by ANTLR 4.7.2. DO NOT EDIT.

package useragent // useragent
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 37, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 5,
	2, 15, 10, 2, 6, 2, 17, 10, 2, 13, 2, 14, 2, 18, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 7, 5, 30, 10, 5, 12, 5, 14, 5, 33, 11, 5,
	3, 6, 3, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2, 34, 2, 16, 3, 2, 2,
	2, 4, 20, 3, 2, 2, 2, 6, 24, 3, 2, 2, 2, 8, 26, 3, 2, 2, 2, 10, 34, 3,
	2, 2, 2, 12, 14, 5, 4, 3, 2, 13, 15, 5, 10, 6, 2, 14, 13, 3, 2, 2, 2, 14,
	15, 3, 2, 2, 2, 15, 17, 3, 2, 2, 2, 16, 12, 3, 2, 2, 2, 17, 18, 3, 2, 2,
	2, 18, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 3, 3, 2, 2, 2, 20, 21, 5,
	6, 4, 2, 21, 22, 7, 3, 2, 2, 22, 23, 5, 8, 5, 2, 23, 5, 3, 2, 2, 2, 24,
	25, 7, 6, 2, 2, 25, 7, 3, 2, 2, 2, 26, 31, 7, 6, 2, 2, 27, 28, 7, 4, 2,
	2, 28, 30, 7, 6, 2, 2, 29, 27, 3, 2, 2, 2, 30, 33, 3, 2, 2, 2, 31, 29,
	3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 9, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2,
	34, 35, 7, 5, 2, 2, 35, 11, 3, 2, 2, 2, 5, 14, 18, 31,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'/'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "COMMENT", "STRING", "WS",
}

var ruleNames = []string{
	"prog", "product", "name", "version", "comment",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type useragentParser struct {
	*antlr.BaseParser
}

func NewuseragentParser(input antlr.TokenStream) *useragentParser {
	this := new(useragentParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "useragent.g4"

	return this
}

// useragentParser tokens.
const (
	useragentParserEOF     = antlr.TokenEOF
	useragentParserT__0    = 1
	useragentParserT__1    = 2
	useragentParserCOMMENT = 3
	useragentParserSTRING  = 4
	useragentParserWS      = 5
)

// useragentParser rules.
const (
	useragentParserRULE_prog    = 0
	useragentParserRULE_product = 1
	useragentParserRULE_name    = 2
	useragentParserRULE_version = 3
	useragentParserRULE_comment = 4
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = useragentParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = useragentParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllProduct() []IProductContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProductContext)(nil)).Elem())
	var tst = make([]IProductContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProductContext)
		}
	}

	return tst
}

func (s *ProgContext) Product(i int) IProductContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProductContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProductContext)
}

func (s *ProgContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *ProgContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(useragentListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(useragentListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *useragentParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, useragentParserRULE_prog)
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
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == useragentParserSTRING {
		{
			p.SetState(10)
			p.Product()
		}
		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == useragentParserCOMMENT {
			{
				p.SetState(11)
				p.Comment()
			}

		}

		p.SetState(16)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IProductContext is an interface to support dynamic dispatch.
type IProductContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProductContext differentiates from other interfaces.
	IsProductContext()
}

type ProductContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProductContext() *ProductContext {
	var p = new(ProductContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = useragentParserRULE_product
	return p
}

func (*ProductContext) IsProductContext() {}

func NewProductContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProductContext {
	var p = new(ProductContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = useragentParserRULE_product

	return p
}

func (s *ProductContext) GetParser() antlr.Parser { return s.parser }

func (s *ProductContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ProductContext) Version() IVersionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVersionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *ProductContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProductContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProductContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(useragentListener); ok {
		listenerT.EnterProduct(s)
	}
}

func (s *ProductContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(useragentListener); ok {
		listenerT.ExitProduct(s)
	}
}

func (p *useragentParser) Product() (localctx IProductContext) {
	localctx = NewProductContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, useragentParserRULE_product)

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
		p.SetState(18)
		p.Name()
	}
	{
		p.SetState(19)
		p.Match(useragentParserT__0)
	}
	{
		p.SetState(20)
		p.Version()
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = useragentParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = useragentParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(useragentParserSTRING, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(useragentListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(useragentListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *useragentParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, useragentParserRULE_name)

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
		p.Match(useragentParserSTRING)
	}

	return localctx
}

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = useragentParserRULE_version
	return p
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = useragentParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(useragentParserSTRING)
}

func (s *VersionContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(useragentParserSTRING, i)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(useragentListener); ok {
		listenerT.EnterVersion(s)
	}
}

func (s *VersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(useragentListener); ok {
		listenerT.ExitVersion(s)
	}
}

func (p *useragentParser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, useragentParserRULE_version)
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
		p.SetState(24)
		p.Match(useragentParserSTRING)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == useragentParserT__1 {
		{
			p.SetState(25)
			p.Match(useragentParserT__1)
		}
		{
			p.SetState(26)
			p.Match(useragentParserSTRING)
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = useragentParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = useragentParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(useragentParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(useragentListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(useragentListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *useragentParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, useragentParserRULE_comment)

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
		p.Match(useragentParserCOMMENT)
	}

	return localctx
}
