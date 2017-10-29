// Generated from tnt.g4 by ANTLR 4.7.

package tnt // tnt
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 81, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 23, 10, 3, 3, 4, 7,
	4, 26, 10, 4, 12, 4, 14, 4, 29, 11, 4, 3, 4, 3, 4, 3, 5, 7, 5, 34, 10,
	5, 12, 5, 14, 5, 37, 11, 5, 3, 5, 3, 5, 7, 5, 41, 10, 5, 12, 5, 14, 5,
	44, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 60, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 7, 6, 68, 10, 6, 12, 6, 14, 6, 71, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 2, 3, 10, 9, 2, 4, 6, 8, 10, 12, 14, 2, 3,
	3, 2, 12, 16, 2, 83, 2, 16, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 27, 3, 2,
	2, 2, 8, 35, 3, 2, 2, 2, 10, 59, 3, 2, 2, 2, 12, 72, 3, 2, 2, 2, 14, 76,
	3, 2, 2, 2, 16, 17, 5, 10, 6, 2, 17, 18, 7, 3, 2, 2, 18, 19, 5, 10, 6,
	2, 19, 3, 3, 2, 2, 2, 20, 23, 5, 6, 4, 2, 21, 23, 5, 8, 5, 2, 22, 20, 3,
	2, 2, 2, 22, 21, 3, 2, 2, 2, 23, 5, 3, 2, 2, 2, 24, 26, 7, 11, 2, 2, 25,
	24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2,
	2, 28, 30, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 31, 7, 10, 2, 2, 31, 7,
	3, 2, 2, 2, 32, 34, 7, 11, 2, 2, 33, 32, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2,
	35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 38, 3, 2, 2, 2, 37, 35, 3,
	2, 2, 2, 38, 42, 9, 2, 2, 2, 39, 41, 7, 17, 2, 2, 40, 39, 3, 2, 2, 2, 41,
	44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 9, 3, 2, 2,
	2, 44, 42, 3, 2, 2, 2, 45, 46, 8, 6, 1, 2, 46, 60, 5, 4, 3, 2, 47, 48,
	7, 6, 2, 2, 48, 49, 5, 10, 6, 2, 49, 50, 7, 7, 2, 2, 50, 60, 3, 2, 2, 2,
	51, 52, 7, 8, 2, 2, 52, 60, 5, 10, 6, 5, 53, 54, 5, 12, 7, 2, 54, 55, 5,
	10, 6, 4, 55, 60, 3, 2, 2, 2, 56, 57, 5, 14, 8, 2, 57, 58, 5, 10, 6, 3,
	58, 60, 3, 2, 2, 2, 59, 45, 3, 2, 2, 2, 59, 47, 3, 2, 2, 2, 59, 51, 3,
	2, 2, 2, 59, 53, 3, 2, 2, 2, 59, 56, 3, 2, 2, 2, 60, 69, 3, 2, 2, 2, 61,
	62, 12, 8, 2, 2, 62, 63, 7, 4, 2, 2, 63, 68, 5, 10, 6, 9, 64, 65, 12, 7,
	2, 2, 65, 66, 7, 5, 2, 2, 66, 68, 5, 10, 6, 8, 67, 61, 3, 2, 2, 2, 67,
	64, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2,
	2, 70, 11, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73, 7, 18, 2, 2, 73, 74,
	5, 8, 5, 2, 74, 75, 7, 9, 2, 2, 75, 13, 3, 2, 2, 2, 76, 77, 7, 19, 2, 2,
	77, 78, 5, 8, 5, 2, 78, 79, 7, 9, 2, 2, 79, 15, 3, 2, 2, 2, 9, 22, 27,
	35, 42, 59, 67, 69,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'+'", "'*'", "'('", "')'", "'~'", "':'", "'0'", "'S'", "'a'",
	"'b'", "'c'", "'d'", "'e'", "'''", "'A'", "'E'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "ZERO", "SUCCESSOR", "A", "B", "C", "D",
	"E", "PRIME", "FOREVERY", "EXISTS", "WS",
}

var ruleNames = []string{
	"equation", "atom", "number", "variable", "expression", "forevery", "exists",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type tntParser struct {
	*antlr.BaseParser
}

func NewtntParser(input antlr.TokenStream) *tntParser {
	this := new(tntParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "tnt.g4"

	return this
}

// tntParser tokens.
const (
	tntParserEOF       = antlr.TokenEOF
	tntParserT__0      = 1
	tntParserT__1      = 2
	tntParserT__2      = 3
	tntParserT__3      = 4
	tntParserT__4      = 5
	tntParserT__5      = 6
	tntParserT__6      = 7
	tntParserZERO      = 8
	tntParserSUCCESSOR = 9
	tntParserA         = 10
	tntParserB         = 11
	tntParserC         = 12
	tntParserD         = 13
	tntParserE         = 14
	tntParserPRIME     = 15
	tntParserFOREVERY  = 16
	tntParserEXISTS    = 17
	tntParserWS        = 18
)

// tntParser rules.
const (
	tntParserRULE_equation   = 0
	tntParserRULE_atom       = 1
	tntParserRULE_number     = 2
	tntParserRULE_variable   = 3
	tntParserRULE_expression = 4
	tntParserRULE_forevery   = 5
	tntParserRULE_exists     = 6
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
	p.RuleIndex = tntParserRULE_equation
	return p
}

func (*EquationContext) IsEquationContext() {}

func NewEquationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquationContext {
	var p = new(EquationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tntParserRULE_equation

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

func (s *EquationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EquationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.EnterEquation(s)
	}
}

func (s *EquationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.ExitEquation(s)
	}
}

func (p *tntParser) Equation() (localctx IEquationContext) {
	localctx = NewEquationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tntParserRULE_equation)

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
		p.SetState(14)
		p.expression(0)
	}
	{
		p.SetState(15)
		p.Match(tntParserT__0)
	}
	{
		p.SetState(16)
		p.expression(0)
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
	p.RuleIndex = tntParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tntParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *AtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *tntParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tntParserRULE_atom)

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

	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.Variable()
		}

	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tntParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tntParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) ZERO() antlr.TerminalNode {
	return s.GetToken(tntParserZERO, 0)
}

func (s *NumberContext) AllSUCCESSOR() []antlr.TerminalNode {
	return s.GetTokens(tntParserSUCCESSOR)
}

func (s *NumberContext) SUCCESSOR(i int) antlr.TerminalNode {
	return s.GetToken(tntParserSUCCESSOR, i)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *tntParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tntParserRULE_number)
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tntParserSUCCESSOR {
		{
			p.SetState(22)
			p.Match(tntParserSUCCESSOR)
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(28)
		p.Match(tntParserZERO)
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
	p.RuleIndex = tntParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tntParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) A() antlr.TerminalNode {
	return s.GetToken(tntParserA, 0)
}

func (s *VariableContext) B() antlr.TerminalNode {
	return s.GetToken(tntParserB, 0)
}

func (s *VariableContext) C() antlr.TerminalNode {
	return s.GetToken(tntParserC, 0)
}

func (s *VariableContext) D() antlr.TerminalNode {
	return s.GetToken(tntParserD, 0)
}

func (s *VariableContext) E() antlr.TerminalNode {
	return s.GetToken(tntParserE, 0)
}

func (s *VariableContext) AllSUCCESSOR() []antlr.TerminalNode {
	return s.GetTokens(tntParserSUCCESSOR)
}

func (s *VariableContext) SUCCESSOR(i int) antlr.TerminalNode {
	return s.GetToken(tntParserSUCCESSOR, i)
}

func (s *VariableContext) AllPRIME() []antlr.TerminalNode {
	return s.GetTokens(tntParserPRIME)
}

func (s *VariableContext) PRIME(i int) antlr.TerminalNode {
	return s.GetToken(tntParserPRIME, i)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *tntParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tntParserRULE_variable)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tntParserSUCCESSOR {
		{
			p.SetState(30)
			p.Match(tntParserSUCCESSOR)
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(36)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tntParserA)|(1<<tntParserB)|(1<<tntParserC)|(1<<tntParserD)|(1<<tntParserE))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(37)
				p.Match(tntParserPRIME)
			}

		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.RuleIndex = tntParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tntParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) Forevery() IForeveryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForeveryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForeveryContext)
}

func (s *ExpressionContext) Exists() IExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExistsContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *tntParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *tntParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, tntParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case tntParserZERO, tntParserSUCCESSOR, tntParserA, tntParserB, tntParserC, tntParserD, tntParserE:
		{
			p.SetState(44)
			p.Atom()
		}

	case tntParserT__3:
		{
			p.SetState(45)
			p.Match(tntParserT__3)
		}
		{
			p.SetState(46)
			p.expression(0)
		}
		{
			p.SetState(47)
			p.Match(tntParserT__4)
		}

	case tntParserT__5:
		{
			p.SetState(49)
			p.Match(tntParserT__5)
		}
		{
			p.SetState(50)
			p.expression(3)
		}

	case tntParserFOREVERY:
		{
			p.SetState(51)
			p.Forevery()
		}
		{
			p.SetState(52)
			p.expression(2)
		}

	case tntParserEXISTS:
		{
			p.SetState(54)
			p.Exists()
		}
		{
			p.SetState(55)
			p.expression(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tntParserRULE_expression)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(60)
					p.Match(tntParserT__1)
				}
				{
					p.SetState(61)
					p.expression(7)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, tntParserRULE_expression)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(63)
					p.Match(tntParserT__2)
				}
				{
					p.SetState(64)
					p.expression(6)
				}

			}

		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IForeveryContext is an interface to support dynamic dispatch.
type IForeveryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForeveryContext differentiates from other interfaces.
	IsForeveryContext()
}

type ForeveryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeveryContext() *ForeveryContext {
	var p = new(ForeveryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tntParserRULE_forevery
	return p
}

func (*ForeveryContext) IsForeveryContext() {}

func NewForeveryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForeveryContext {
	var p = new(ForeveryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tntParserRULE_forevery

	return p
}

func (s *ForeveryContext) GetParser() antlr.Parser { return s.parser }

func (s *ForeveryContext) FOREVERY() antlr.TerminalNode {
	return s.GetToken(tntParserFOREVERY, 0)
}

func (s *ForeveryContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ForeveryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForeveryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForeveryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.EnterForevery(s)
	}
}

func (s *ForeveryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.ExitForevery(s)
	}
}

func (p *tntParser) Forevery() (localctx IForeveryContext) {
	localctx = NewForeveryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tntParserRULE_forevery)

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
		p.SetState(70)
		p.Match(tntParserFOREVERY)
	}
	{
		p.SetState(71)
		p.Variable()
	}
	{
		p.SetState(72)
		p.Match(tntParserT__6)
	}

	return localctx
}

// IExistsContext is an interface to support dynamic dispatch.
type IExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExistsContext differentiates from other interfaces.
	IsExistsContext()
}

type ExistsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExistsContext() *ExistsContext {
	var p = new(ExistsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tntParserRULE_exists
	return p
}

func (*ExistsContext) IsExistsContext() {}

func NewExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExistsContext {
	var p = new(ExistsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tntParserRULE_exists

	return p
}

func (s *ExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExistsContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(tntParserEXISTS, 0)
}

func (s *ExistsContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExistsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.EnterExists(s)
	}
}

func (s *ExistsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tntListener); ok {
		listenerT.ExitExists(s)
	}
}

func (p *tntParser) Exists() (localctx IExistsContext) {
	localctx = NewExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tntParserRULE_exists)

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
		p.SetState(74)
		p.Match(tntParserEXISTS)
	}
	{
		p.SetState(75)
		p.Variable()
	}
	{
		p.SetState(76)
		p.Match(tntParserT__6)
	}

	return localctx
}

func (p *tntParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *tntParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
