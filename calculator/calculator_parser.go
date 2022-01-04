// Code generated from calculator.g4 by ANTLR 4.9.3. DO NOT EDIT.

package calculator // calculator
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 94, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 34, 10, 3, 12, 3, 14,
	3, 37, 11, 3, 3, 4, 3, 4, 3, 4, 7, 4, 42, 10, 4, 12, 4, 14, 4, 45, 11,
	4, 3, 5, 3, 5, 3, 5, 7, 5, 50, 10, 5, 12, 5, 14, 5, 53, 11, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 61, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 5, 7, 70, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 83, 10, 11, 12, 11, 14, 11, 86,
	11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 2, 2, 14, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 7, 3, 2, 14, 15, 3, 2, 16, 17,
	3, 2, 24, 26, 3, 2, 3, 11, 3, 2, 18, 20, 2, 91, 2, 26, 3, 2, 2, 2, 4, 30,
	3, 2, 2, 2, 6, 38, 3, 2, 2, 2, 8, 46, 3, 2, 2, 2, 10, 60, 3, 2, 2, 2, 12,
	69, 3, 2, 2, 2, 14, 71, 3, 2, 2, 2, 16, 73, 3, 2, 2, 2, 18, 75, 3, 2, 2,
	2, 20, 77, 3, 2, 2, 2, 22, 89, 3, 2, 2, 2, 24, 91, 3, 2, 2, 2, 26, 27,
	5, 4, 3, 2, 27, 28, 5, 24, 13, 2, 28, 29, 5, 4, 3, 2, 29, 3, 3, 2, 2, 2,
	30, 35, 5, 6, 4, 2, 31, 32, 9, 2, 2, 2, 32, 34, 5, 6, 4, 2, 33, 31, 3,
	2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36,
	5, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 43, 5, 8, 5, 2, 39, 40, 9, 3, 2,
	2, 40, 42, 5, 8, 5, 2, 41, 39, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41,
	3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 7, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2,
	46, 51, 5, 10, 6, 2, 47, 48, 7, 23, 2, 2, 48, 50, 5, 10, 6, 2, 49, 47,
	3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2,
	52, 9, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 55, 7, 14, 2, 2, 55, 61, 5,
	10, 6, 2, 56, 57, 7, 15, 2, 2, 57, 61, 5, 10, 6, 2, 58, 61, 5, 20, 11,
	2, 59, 61, 5, 12, 7, 2, 60, 54, 3, 2, 2, 2, 60, 56, 3, 2, 2, 2, 60, 58,
	3, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 11, 3, 2, 2, 2, 62, 70, 5, 14, 8, 2,
	63, 70, 5, 18, 10, 2, 64, 70, 5, 16, 9, 2, 65, 66, 7, 12, 2, 2, 66, 67,
	5, 4, 3, 2, 67, 68, 7, 13, 2, 2, 68, 70, 3, 2, 2, 2, 69, 62, 3, 2, 2, 2,
	69, 63, 3, 2, 2, 2, 69, 64, 3, 2, 2, 2, 69, 65, 3, 2, 2, 2, 70, 13, 3,
	2, 2, 2, 71, 72, 7, 28, 2, 2, 72, 15, 3, 2, 2, 2, 73, 74, 9, 4, 2, 2, 74,
	17, 3, 2, 2, 2, 75, 76, 7, 27, 2, 2, 76, 19, 3, 2, 2, 2, 77, 78, 5, 22,
	12, 2, 78, 79, 7, 12, 2, 2, 79, 84, 5, 4, 3, 2, 80, 81, 7, 21, 2, 2, 81,
	83, 5, 4, 3, 2, 82, 80, 3, 2, 2, 2, 83, 86, 3, 2, 2, 2, 84, 82, 3, 2, 2,
	2, 84, 85, 3, 2, 2, 2, 85, 87, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 87, 88,
	7, 13, 2, 2, 88, 21, 3, 2, 2, 2, 89, 90, 9, 5, 2, 2, 90, 23, 3, 2, 2, 2,
	91, 92, 9, 6, 2, 2, 92, 25, 3, 2, 2, 2, 8, 35, 43, 51, 60, 69, 84,
}
var literalNames = []string{
	"", "'cos'", "'sin'", "'tan'", "'acos'", "'asin'", "'atan'", "'ln'", "'log'",
	"'sqrt'", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "'>'", "'<'", "'='",
	"','", "'.'", "'^'", "'pi'", "", "'i'",
}
var symbolicNames = []string{
	"", "COS", "SIN", "TAN", "ACOS", "ASIN", "ATAN", "LN", "LOG", "SQRT", "LPAREN",
	"RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "GT", "LT", "EQ", "COMMA", "POINT",
	"POW", "PI", "EULER", "I", "VARIABLE", "SCIENTIFIC_NUMBER", "WS",
}

var ruleNames = []string{
	"equation", "expression", "multiplyingExpression", "powExpression", "signedAtom",
	"atom", "scientific", "constant", "variable", "func_", "funcname", "relop",
}

type calculatorParser struct {
	*antlr.BaseParser
}

// NewcalculatorParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *calculatorParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewcalculatorParser(input antlr.TokenStream) *calculatorParser {
	this := new(calculatorParser)
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
	this.GrammarFileName = "calculator.g4"

	return this
}

// calculatorParser tokens.
const (
	calculatorParserEOF               = antlr.TokenEOF
	calculatorParserCOS               = 1
	calculatorParserSIN               = 2
	calculatorParserTAN               = 3
	calculatorParserACOS              = 4
	calculatorParserASIN              = 5
	calculatorParserATAN              = 6
	calculatorParserLN                = 7
	calculatorParserLOG               = 8
	calculatorParserSQRT              = 9
	calculatorParserLPAREN            = 10
	calculatorParserRPAREN            = 11
	calculatorParserPLUS              = 12
	calculatorParserMINUS             = 13
	calculatorParserTIMES             = 14
	calculatorParserDIV               = 15
	calculatorParserGT                = 16
	calculatorParserLT                = 17
	calculatorParserEQ                = 18
	calculatorParserCOMMA             = 19
	calculatorParserPOINT             = 20
	calculatorParserPOW               = 21
	calculatorParserPI                = 22
	calculatorParserEULER             = 23
	calculatorParserI                 = 24
	calculatorParserVARIABLE          = 25
	calculatorParserSCIENTIFIC_NUMBER = 26
	calculatorParserWS                = 27
)

// calculatorParser rules.
const (
	calculatorParserRULE_equation              = 0
	calculatorParserRULE_expression            = 1
	calculatorParserRULE_multiplyingExpression = 2
	calculatorParserRULE_powExpression         = 3
	calculatorParserRULE_signedAtom            = 4
	calculatorParserRULE_atom                  = 5
	calculatorParserRULE_scientific            = 6
	calculatorParserRULE_constant              = 7
	calculatorParserRULE_variable              = 8
	calculatorParserRULE_func_                 = 9
	calculatorParserRULE_funcname              = 10
	calculatorParserRULE_relop                 = 11
)

// IEquationContext is an interface to support dynamic dispatch.
type IEquationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEquationContext differentiates from other interfaces.
	IsEquationContext()
}

type EquationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquationContext() *EquationContext {
	var p = new(EquationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_equation
	return p
}

func (*EquationContext) IsEquationContext() {}

func NewEquationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquationContext {
	var p = new(EquationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_equation

	return p
}

func (s *EquationContext) GetParser() antlr.Parser { return s.parser }

func (s *EquationContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EquationContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EquationContext) Relop() IRelopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelopContext)
}

func (s *EquationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EquationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterEquation(s)
	}
}

func (s *EquationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitEquation(s)
	}
}

func (p *calculatorParser) Equation() (localctx IEquationContext) {
	this := p
	_ = this

	localctx = NewEquationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, calculatorParserRULE_equation)

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
		p.Expression()
	}
	{
		p.SetState(25)
		p.Relop()
	}
	{
		p.SetState(26)
		p.Expression()
	}

	return localctx
}

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
	p.RuleIndex = calculatorParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllMultiplyingExpression() []IMultiplyingExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplyingExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplyingExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) MultiplyingExpression(i int) IMultiplyingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplyingExpressionContext)
}

func (s *ExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(calculatorParserPLUS)
}

func (s *ExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(calculatorParserPLUS, i)
}

func (s *ExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(calculatorParserMINUS)
}

func (s *ExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(calculatorParserMINUS, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *calculatorParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, calculatorParserRULE_expression)
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
		p.MultiplyingExpression()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == calculatorParserPLUS || _la == calculatorParserMINUS {
		{
			p.SetState(29)
			_la = p.GetTokenStream().LA(1)

			if !(_la == calculatorParserPLUS || _la == calculatorParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(30)
			p.MultiplyingExpression()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplyingExpressionContext is an interface to support dynamic dispatch.
type IMultiplyingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplyingExpressionContext differentiates from other interfaces.
	IsMultiplyingExpressionContext()
}

type MultiplyingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyingExpressionContext() *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_multiplyingExpression
	return p
}

func (*MultiplyingExpressionContext) IsMultiplyingExpressionContext() {}

func NewMultiplyingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_multiplyingExpression

	return p
}

func (s *MultiplyingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyingExpressionContext) AllPowExpression() []IPowExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPowExpressionContext)(nil)).Elem())
	var tst = make([]IPowExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPowExpressionContext)
		}
	}

	return tst
}

func (s *MultiplyingExpressionContext) PowExpression(i int) IPowExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPowExpressionContext)
}

func (s *MultiplyingExpressionContext) AllTIMES() []antlr.TerminalNode {
	return s.GetTokens(calculatorParserTIMES)
}

func (s *MultiplyingExpressionContext) TIMES(i int) antlr.TerminalNode {
	return s.GetToken(calculatorParserTIMES, i)
}

func (s *MultiplyingExpressionContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(calculatorParserDIV)
}

func (s *MultiplyingExpressionContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(calculatorParserDIV, i)
}

func (s *MultiplyingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterMultiplyingExpression(s)
	}
}

func (s *MultiplyingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitMultiplyingExpression(s)
	}
}

func (p *calculatorParser) MultiplyingExpression() (localctx IMultiplyingExpressionContext) {
	this := p
	_ = this

	localctx = NewMultiplyingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, calculatorParserRULE_multiplyingExpression)
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
		p.SetState(36)
		p.PowExpression()
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == calculatorParserTIMES || _la == calculatorParserDIV {
		{
			p.SetState(37)
			_la = p.GetTokenStream().LA(1)

			if !(_la == calculatorParserTIMES || _la == calculatorParserDIV) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(38)
			p.PowExpression()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPowExpressionContext is an interface to support dynamic dispatch.
type IPowExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPowExpressionContext differentiates from other interfaces.
	IsPowExpressionContext()
}

type PowExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowExpressionContext() *PowExpressionContext {
	var p = new(PowExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_powExpression
	return p
}

func (*PowExpressionContext) IsPowExpressionContext() {}

func NewPowExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowExpressionContext {
	var p = new(PowExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_powExpression

	return p
}

func (s *PowExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PowExpressionContext) AllSignedAtom() []ISignedAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISignedAtomContext)(nil)).Elem())
	var tst = make([]ISignedAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISignedAtomContext)
		}
	}

	return tst
}

func (s *PowExpressionContext) SignedAtom(i int) ISignedAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignedAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISignedAtomContext)
}

func (s *PowExpressionContext) AllPOW() []antlr.TerminalNode {
	return s.GetTokens(calculatorParserPOW)
}

func (s *PowExpressionContext) POW(i int) antlr.TerminalNode {
	return s.GetToken(calculatorParserPOW, i)
}

func (s *PowExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterPowExpression(s)
	}
}

func (s *PowExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitPowExpression(s)
	}
}

func (p *calculatorParser) PowExpression() (localctx IPowExpressionContext) {
	this := p
	_ = this

	localctx = NewPowExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, calculatorParserRULE_powExpression)
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
		p.SetState(44)
		p.SignedAtom()
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == calculatorParserPOW {
		{
			p.SetState(45)
			p.Match(calculatorParserPOW)
		}
		{
			p.SetState(46)
			p.SignedAtom()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISignedAtomContext is an interface to support dynamic dispatch.
type ISignedAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignedAtomContext differentiates from other interfaces.
	IsSignedAtomContext()
}

type SignedAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignedAtomContext() *SignedAtomContext {
	var p = new(SignedAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_signedAtom
	return p
}

func (*SignedAtomContext) IsSignedAtomContext() {}

func NewSignedAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignedAtomContext {
	var p = new(SignedAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_signedAtom

	return p
}

func (s *SignedAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *SignedAtomContext) PLUS() antlr.TerminalNode {
	return s.GetToken(calculatorParserPLUS, 0)
}

func (s *SignedAtomContext) SignedAtom() ISignedAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignedAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignedAtomContext)
}

func (s *SignedAtomContext) MINUS() antlr.TerminalNode {
	return s.GetToken(calculatorParserMINUS, 0)
}

func (s *SignedAtomContext) Func_() IFunc_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_Context)
}

func (s *SignedAtomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *SignedAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignedAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignedAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterSignedAtom(s)
	}
}

func (s *SignedAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitSignedAtom(s)
	}
}

func (p *calculatorParser) SignedAtom() (localctx ISignedAtomContext) {
	this := p
	_ = this

	localctx = NewSignedAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, calculatorParserRULE_signedAtom)

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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case calculatorParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(calculatorParserPLUS)
		}
		{
			p.SetState(53)
			p.SignedAtom()
		}

	case calculatorParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(calculatorParserMINUS)
		}
		{
			p.SetState(55)
			p.SignedAtom()
		}

	case calculatorParserCOS, calculatorParserSIN, calculatorParserTAN, calculatorParserACOS, calculatorParserASIN, calculatorParserATAN, calculatorParserLN, calculatorParserLOG, calculatorParserSQRT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Func_()
		}

	case calculatorParserLPAREN, calculatorParserPI, calculatorParserEULER, calculatorParserI, calculatorParserVARIABLE, calculatorParserSCIENTIFIC_NUMBER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(57)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Scientific() IScientificContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScientificContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScientificContext)
}

func (s *AtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *AtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(calculatorParserLPAREN, 0)
}

func (s *AtomContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(calculatorParserRPAREN, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *calculatorParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, calculatorParserRULE_atom)

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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case calculatorParserSCIENTIFIC_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Scientific()
		}

	case calculatorParserVARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Variable()
		}

	case calculatorParserPI, calculatorParserEULER, calculatorParserI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Constant()
		}

	case calculatorParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(63)
			p.Match(calculatorParserLPAREN)
		}
		{
			p.SetState(64)
			p.Expression()
		}
		{
			p.SetState(65)
			p.Match(calculatorParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IScientificContext is an interface to support dynamic dispatch.
type IScientificContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScientificContext differentiates from other interfaces.
	IsScientificContext()
}

type ScientificContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScientificContext() *ScientificContext {
	var p = new(ScientificContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_scientific
	return p
}

func (*ScientificContext) IsScientificContext() {}

func NewScientificContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScientificContext {
	var p = new(ScientificContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_scientific

	return p
}

func (s *ScientificContext) GetParser() antlr.Parser { return s.parser }

func (s *ScientificContext) SCIENTIFIC_NUMBER() antlr.TerminalNode {
	return s.GetToken(calculatorParserSCIENTIFIC_NUMBER, 0)
}

func (s *ScientificContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScientificContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScientificContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterScientific(s)
	}
}

func (s *ScientificContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitScientific(s)
	}
}

func (p *calculatorParser) Scientific() (localctx IScientificContext) {
	this := p
	_ = this

	localctx = NewScientificContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, calculatorParserRULE_scientific)

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
		p.SetState(69)
		p.Match(calculatorParserSCIENTIFIC_NUMBER)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) PI() antlr.TerminalNode {
	return s.GetToken(calculatorParserPI, 0)
}

func (s *ConstantContext) EULER() antlr.TerminalNode {
	return s.GetToken(calculatorParserEULER, 0)
}

func (s *ConstantContext) I() antlr.TerminalNode {
	return s.GetToken(calculatorParserI, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *calculatorParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, calculatorParserRULE_constant)
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
		p.SetState(71)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<calculatorParserPI)|(1<<calculatorParserEULER)|(1<<calculatorParserI))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(calculatorParserVARIABLE, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *calculatorParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, calculatorParserRULE_variable)

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
		p.Match(calculatorParserVARIABLE)
	}

	return localctx
}

// IFunc_Context is an interface to support dynamic dispatch.
type IFunc_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_Context differentiates from other interfaces.
	IsFunc_Context()
}

type Func_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_Context() *Func_Context {
	var p = new(Func_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_func_
	return p
}

func (*Func_Context) IsFunc_Context() {}

func NewFunc_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_Context {
	var p = new(Func_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_func_

	return p
}

func (s *Func_Context) GetParser() antlr.Parser { return s.parser }

func (s *Func_Context) Funcname() IFuncnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncnameContext)
}

func (s *Func_Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(calculatorParserLPAREN, 0)
}

func (s *Func_Context) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Func_Context) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Func_Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(calculatorParserRPAREN, 0)
}

func (s *Func_Context) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(calculatorParserCOMMA)
}

func (s *Func_Context) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(calculatorParserCOMMA, i)
}

func (s *Func_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterFunc_(s)
	}
}

func (s *Func_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitFunc_(s)
	}
}

func (p *calculatorParser) Func_() (localctx IFunc_Context) {
	this := p
	_ = this

	localctx = NewFunc_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, calculatorParserRULE_func_)
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
		p.SetState(75)
		p.Funcname()
	}
	{
		p.SetState(76)
		p.Match(calculatorParserLPAREN)
	}
	{
		p.SetState(77)
		p.Expression()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == calculatorParserCOMMA {
		{
			p.SetState(78)
			p.Match(calculatorParserCOMMA)
		}
		{
			p.SetState(79)
			p.Expression()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(85)
		p.Match(calculatorParserRPAREN)
	}

	return localctx
}

// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncnameContext differentiates from other interfaces.
	IsFuncnameContext()
}

type FuncnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncnameContext() *FuncnameContext {
	var p = new(FuncnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_funcname
	return p
}

func (*FuncnameContext) IsFuncnameContext() {}

func NewFuncnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncnameContext {
	var p = new(FuncnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_funcname

	return p
}

func (s *FuncnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncnameContext) COS() antlr.TerminalNode {
	return s.GetToken(calculatorParserCOS, 0)
}

func (s *FuncnameContext) TAN() antlr.TerminalNode {
	return s.GetToken(calculatorParserTAN, 0)
}

func (s *FuncnameContext) SIN() antlr.TerminalNode {
	return s.GetToken(calculatorParserSIN, 0)
}

func (s *FuncnameContext) ACOS() antlr.TerminalNode {
	return s.GetToken(calculatorParserACOS, 0)
}

func (s *FuncnameContext) ATAN() antlr.TerminalNode {
	return s.GetToken(calculatorParserATAN, 0)
}

func (s *FuncnameContext) ASIN() antlr.TerminalNode {
	return s.GetToken(calculatorParserASIN, 0)
}

func (s *FuncnameContext) LOG() antlr.TerminalNode {
	return s.GetToken(calculatorParserLOG, 0)
}

func (s *FuncnameContext) LN() antlr.TerminalNode {
	return s.GetToken(calculatorParserLN, 0)
}

func (s *FuncnameContext) SQRT() antlr.TerminalNode {
	return s.GetToken(calculatorParserSQRT, 0)
}

func (s *FuncnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterFuncname(s)
	}
}

func (s *FuncnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitFuncname(s)
	}
}

func (p *calculatorParser) Funcname() (localctx IFuncnameContext) {
	this := p
	_ = this

	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, calculatorParserRULE_funcname)
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
		p.SetState(87)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<calculatorParserCOS)|(1<<calculatorParserSIN)|(1<<calculatorParserTAN)|(1<<calculatorParserACOS)|(1<<calculatorParserASIN)|(1<<calculatorParserATAN)|(1<<calculatorParserLN)|(1<<calculatorParserLOG)|(1<<calculatorParserSQRT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRelopContext is an interface to support dynamic dispatch.
type IRelopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelopContext differentiates from other interfaces.
	IsRelopContext()
}

type RelopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelopContext() *RelopContext {
	var p = new(RelopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = calculatorParserRULE_relop
	return p
}

func (*RelopContext) IsRelopContext() {}

func NewRelopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelopContext {
	var p = new(RelopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = calculatorParserRULE_relop

	return p
}

func (s *RelopContext) GetParser() antlr.Parser { return s.parser }

func (s *RelopContext) EQ() antlr.TerminalNode {
	return s.GetToken(calculatorParserEQ, 0)
}

func (s *RelopContext) GT() antlr.TerminalNode {
	return s.GetToken(calculatorParserGT, 0)
}

func (s *RelopContext) LT() antlr.TerminalNode {
	return s.GetToken(calculatorParserLT, 0)
}

func (s *RelopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.EnterRelop(s)
	}
}

func (s *RelopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calculatorListener); ok {
		listenerT.ExitRelop(s)
	}
}

func (p *calculatorParser) Relop() (localctx IRelopContext) {
	this := p
	_ = this

	localctx = NewRelopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, calculatorParserRULE_relop)
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
		p.SetState(89)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<calculatorParserGT)|(1<<calculatorParserLT)|(1<<calculatorParserEQ))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
