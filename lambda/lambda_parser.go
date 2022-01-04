// Code generated from lambda.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lambda // lambda
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 28, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 5, 2, 14,
	10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 2, 2, 25, 2, 13, 3, 2, 2, 2, 4, 15,
	3, 2, 2, 2, 6, 20, 3, 2, 2, 2, 8, 25, 3, 2, 2, 2, 10, 14, 7, 7, 2, 2, 11,
	14, 5, 4, 3, 2, 12, 14, 5, 6, 4, 2, 13, 10, 3, 2, 2, 2, 13, 11, 3, 2, 2,
	2, 13, 12, 3, 2, 2, 2, 14, 3, 3, 2, 2, 2, 15, 16, 7, 3, 2, 2, 16, 17, 7,
	7, 2, 2, 17, 18, 7, 4, 2, 2, 18, 19, 5, 8, 5, 2, 19, 5, 3, 2, 2, 2, 20,
	21, 7, 5, 2, 2, 21, 22, 5, 2, 2, 2, 22, 23, 5, 2, 2, 2, 23, 24, 7, 6, 2,
	2, 24, 7, 3, 2, 2, 2, 25, 26, 5, 2, 2, 2, 26, 9, 3, 2, 2, 2, 3, 13,
}
var literalNames = []string{
	"", "'\u03BB'", "'.'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "VARIABLE", "WS",
}

var ruleNames = []string{
	"expression", "function_", "application", "scope",
}

type lambdaParser struct {
	*antlr.BaseParser
}

// NewlambdaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *lambdaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewlambdaParser(input antlr.TokenStream) *lambdaParser {
	this := new(lambdaParser)
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
	this.GrammarFileName = "lambda.g4"

	return this
}

// lambdaParser tokens.
const (
	lambdaParserEOF      = antlr.TokenEOF
	lambdaParserT__0     = 1
	lambdaParserT__1     = 2
	lambdaParserT__2     = 3
	lambdaParserT__3     = 4
	lambdaParserVARIABLE = 5
	lambdaParserWS       = 6
)

// lambdaParser rules.
const (
	lambdaParserRULE_expression  = 0
	lambdaParserRULE_function_   = 1
	lambdaParserRULE_application = 2
	lambdaParserRULE_scope       = 3
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lambdaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lambdaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(lambdaParserVARIABLE, 0)
}

func (s *ExpressionContext) Function_() IFunction_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_Context)
}

func (s *ExpressionContext) Application() IApplicationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IApplicationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IApplicationContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lambdaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lambdaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *lambdaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, lambdaParserRULE_expression)

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

	p.SetState(11)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case lambdaParserVARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(8)
			p.Match(lambdaParserVARIABLE)
		}

	case lambdaParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(9)
			p.Function_()
		}

	case lambdaParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(10)
			p.Application()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunction_Context is an interface to support dynamic dispatch.
type IFunction_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_Context differentiates from other interfaces.
	IsFunction_Context()
}

type Function_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_Context() *Function_Context {
	var p = new(Function_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lambdaParserRULE_function_
	return p
}

func (*Function_Context) IsFunction_Context() {}

func NewFunction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_Context {
	var p = new(Function_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lambdaParserRULE_function_

	return p
}

func (s *Function_Context) GetParser() antlr.Parser { return s.parser }

func (s *Function_Context) VARIABLE() antlr.TerminalNode {
	return s.GetToken(lambdaParserVARIABLE, 0)
}

func (s *Function_Context) Scope() IScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *Function_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lambdaListener); ok {
		listenerT.EnterFunction_(s)
	}
}

func (s *Function_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lambdaListener); ok {
		listenerT.ExitFunction_(s)
	}
}

func (p *lambdaParser) Function_() (localctx IFunction_Context) {
	this := p
	_ = this

	localctx = NewFunction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, lambdaParserRULE_function_)

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
		p.SetState(13)
		p.Match(lambdaParserT__0)
	}
	{
		p.SetState(14)
		p.Match(lambdaParserVARIABLE)
	}
	{
		p.SetState(15)
		p.Match(lambdaParserT__1)
	}
	{
		p.SetState(16)
		p.Scope()
	}

	return localctx
}

// IApplicationContext is an interface to support dynamic dispatch.
type IApplicationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsApplicationContext differentiates from other interfaces.
	IsApplicationContext()
}

type ApplicationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApplicationContext() *ApplicationContext {
	var p = new(ApplicationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lambdaParserRULE_application
	return p
}

func (*ApplicationContext) IsApplicationContext() {}

func NewApplicationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApplicationContext {
	var p = new(ApplicationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lambdaParserRULE_application

	return p
}

func (s *ApplicationContext) GetParser() antlr.Parser { return s.parser }

func (s *ApplicationContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ApplicationContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ApplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplicationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApplicationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lambdaListener); ok {
		listenerT.EnterApplication(s)
	}
}

func (s *ApplicationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lambdaListener); ok {
		listenerT.ExitApplication(s)
	}
}

func (p *lambdaParser) Application() (localctx IApplicationContext) {
	this := p
	_ = this

	localctx = NewApplicationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, lambdaParserRULE_application)

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
		p.Match(lambdaParserT__2)
	}
	{
		p.SetState(19)
		p.Expression()
	}
	{
		p.SetState(20)
		p.Expression()
	}
	{
		p.SetState(21)
		p.Match(lambdaParserT__3)
	}

	return localctx
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = lambdaParserRULE_scope
	return p
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = lambdaParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lambdaListener); ok {
		listenerT.EnterScope(s)
	}
}

func (s *ScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(lambdaListener); ok {
		listenerT.ExitScope(s)
	}
}

func (p *lambdaParser) Scope() (localctx IScopeContext) {
	this := p
	_ = this

	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, lambdaParserRULE_scope)

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
		p.SetState(23)
		p.Expression()
	}

	return localctx
}
