// Generated from fol.g4 by ANTLR 4.7.

package fol // fol
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 94, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 5, 3, 28, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 35, 10, 4, 12,
	4, 14, 4, 38, 11, 4, 3, 5, 3, 5, 3, 5, 7, 5, 43, 10, 5, 12, 5, 14, 5, 46,
	11, 5, 3, 6, 5, 6, 49, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 56, 10,
	6, 3, 7, 3, 7, 3, 7, 5, 7, 61, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 67,
	10, 8, 12, 8, 14, 8, 70, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 76, 10, 9,
	3, 10, 3, 10, 3, 10, 5, 10, 81, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7,
	11, 87, 10, 11, 12, 11, 14, 11, 90, 11, 11, 3, 11, 3, 11, 3, 11, 2, 2,
	12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 3, 2, 9, 10, 3, 2, 11, 12,
	2, 93, 2, 22, 3, 2, 2, 2, 4, 27, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8, 39,
	3, 2, 2, 2, 10, 48, 3, 2, 2, 2, 12, 60, 3, 2, 2, 2, 14, 62, 3, 2, 2, 2,
	16, 75, 3, 2, 2, 2, 18, 80, 3, 2, 2, 2, 20, 82, 3, 2, 2, 2, 22, 23, 5,
	4, 3, 2, 23, 24, 7, 2, 2, 3, 24, 3, 3, 2, 2, 2, 25, 26, 9, 2, 2, 2, 26,
	28, 7, 11, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 29, 3, 2,
	2, 2, 29, 30, 5, 6, 4, 2, 30, 5, 3, 2, 2, 2, 31, 36, 5, 8, 5, 2, 32, 33,
	7, 7, 2, 2, 33, 35, 5, 8, 5, 2, 34, 32, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2,
	36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 7, 3, 2, 2, 2, 38, 36, 3, 2,
	2, 2, 39, 44, 5, 10, 6, 2, 40, 41, 7, 6, 2, 2, 41, 43, 5, 10, 6, 2, 42,
	40, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2,
	2, 45, 9, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 49, 7, 8, 2, 2, 48, 47, 3,
	2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 55, 3, 2, 2, 2, 50, 56, 5, 12, 7, 2, 51,
	52, 7, 4, 2, 2, 52, 53, 5, 4, 3, 2, 53, 54, 7, 5, 2, 2, 54, 56, 3, 2, 2,
	2, 55, 50, 3, 2, 2, 2, 55, 51, 3, 2, 2, 2, 56, 11, 3, 2, 2, 2, 57, 58,
	7, 13, 2, 2, 58, 61, 5, 14, 8, 2, 59, 61, 7, 13, 2, 2, 60, 57, 3, 2, 2,
	2, 60, 59, 3, 2, 2, 2, 61, 13, 3, 2, 2, 2, 62, 63, 7, 4, 2, 2, 63, 68,
	5, 16, 9, 2, 64, 65, 7, 3, 2, 2, 65, 67, 5, 16, 9, 2, 66, 64, 3, 2, 2,
	2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 71,
	3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 72, 7, 5, 2, 2, 72, 15, 3, 2, 2, 2,
	73, 76, 5, 18, 10, 2, 74, 76, 7, 11, 2, 2, 75, 73, 3, 2, 2, 2, 75, 74,
	3, 2, 2, 2, 76, 17, 3, 2, 2, 2, 77, 78, 7, 12, 2, 2, 78, 81, 5, 20, 11,
	2, 79, 81, 7, 12, 2, 2, 80, 77, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81, 19,
	3, 2, 2, 2, 82, 83, 7, 4, 2, 2, 83, 88, 9, 3, 2, 2, 84, 85, 7, 3, 2, 2,
	85, 87, 9, 3, 2, 2, 86, 84, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 3,
	2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 91, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91,
	92, 7, 5, 2, 2, 92, 21, 3, 2, 2, 2, 12, 27, 36, 44, 48, 55, 60, 68, 75,
	80, 88,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'('", "')'", "'&'", "'|'", "'!'", "'Forall'", "'Exists'",
}
var symbolicNames = []string{
	"", "", "LPAREN", "RPAREN", "AND", "OR", "NOT", "FORALL", "EXISTS", "VARIABLE",
	"CONSTANT", "PREPOSITION", "WS",
}

var ruleNames = []string{
	"condition", "formula", "disjunction", "conjunction", "negation", "predicate",
	"predicateTuple", "term", "function", "functionTuple",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type folParser struct {
	*antlr.BaseParser
}

func NewfolParser(input antlr.TokenStream) *folParser {
	this := new(folParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "fol.g4"

	return this
}

// folParser tokens.
const (
	folParserEOF         = antlr.TokenEOF
	folParserT__0        = 1
	folParserLPAREN      = 2
	folParserRPAREN      = 3
	folParserAND         = 4
	folParserOR          = 5
	folParserNOT         = 6
	folParserFORALL      = 7
	folParserEXISTS      = 8
	folParserVARIABLE    = 9
	folParserCONSTANT    = 10
	folParserPREPOSITION = 11
	folParserWS          = 12
)

// folParser rules.
const (
	folParserRULE_condition      = 0
	folParserRULE_formula        = 1
	folParserRULE_disjunction    = 2
	folParserRULE_conjunction    = 3
	folParserRULE_negation       = 4
	folParserRULE_predicate      = 5
	folParserRULE_predicateTuple = 6
	folParserRULE_term           = 7
	folParserRULE_function       = 8
	folParserRULE_functionTuple  = 9
)

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Formula() IFormulaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormulaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormulaContext)
}

func (s *ConditionContext) EOF() antlr.TerminalNode {
	return s.GetToken(folParserEOF, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *folParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, folParserRULE_condition)

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
		p.Formula()
	}
	{
		p.SetState(21)
		p.Match(folParserEOF)
	}

	return localctx
}

// IFormulaContext is an interface to support dynamic dispatch.
type IFormulaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormulaContext differentiates from other interfaces.
	IsFormulaContext()
}

type FormulaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormulaContext() *FormulaContext {
	var p = new(FormulaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_formula
	return p
}

func (*FormulaContext) IsFormulaContext() {}

func NewFormulaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormulaContext {
	var p = new(FormulaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_formula

	return p
}

func (s *FormulaContext) GetParser() antlr.Parser { return s.parser }

func (s *FormulaContext) Disjunction() IDisjunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisjunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisjunctionContext)
}

func (s *FormulaContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(folParserVARIABLE, 0)
}

func (s *FormulaContext) FORALL() antlr.TerminalNode {
	return s.GetToken(folParserFORALL, 0)
}

func (s *FormulaContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(folParserEXISTS, 0)
}

func (s *FormulaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormulaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormulaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterFormula(s)
	}
}

func (s *FormulaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitFormula(s)
	}
}

func (p *folParser) Formula() (localctx IFormulaContext) {
	localctx = NewFormulaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, folParserRULE_formula)
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

	if _la == folParserFORALL || _la == folParserEXISTS {
		p.SetState(23)
		_la = p.GetTokenStream().LA(1)

		if !(_la == folParserFORALL || _la == folParserEXISTS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(24)
			p.Match(folParserVARIABLE)
		}

	}
	{
		p.SetState(27)
		p.Disjunction()
	}

	return localctx
}

// IDisjunctionContext is an interface to support dynamic dispatch.
type IDisjunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisjunctionContext differentiates from other interfaces.
	IsDisjunctionContext()
}

type DisjunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisjunctionContext() *DisjunctionContext {
	var p = new(DisjunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_disjunction
	return p
}

func (*DisjunctionContext) IsDisjunctionContext() {}

func NewDisjunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisjunctionContext {
	var p = new(DisjunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_disjunction

	return p
}

func (s *DisjunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *DisjunctionContext) AllConjunction() []IConjunctionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConjunctionContext)(nil)).Elem())
	var tst = make([]IConjunctionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConjunctionContext)
		}
	}

	return tst
}

func (s *DisjunctionContext) Conjunction(i int) IConjunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConjunctionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConjunctionContext)
}

func (s *DisjunctionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(folParserOR)
}

func (s *DisjunctionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(folParserOR, i)
}

func (s *DisjunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisjunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisjunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterDisjunction(s)
	}
}

func (s *DisjunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitDisjunction(s)
	}
}

func (p *folParser) Disjunction() (localctx IDisjunctionContext) {
	localctx = NewDisjunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, folParserRULE_disjunction)
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
		p.SetState(29)
		p.Conjunction()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == folParserOR {
		{
			p.SetState(30)
			p.Match(folParserOR)
		}
		{
			p.SetState(31)
			p.Conjunction()
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConjunctionContext is an interface to support dynamic dispatch.
type IConjunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConjunctionContext differentiates from other interfaces.
	IsConjunctionContext()
}

type ConjunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConjunctionContext() *ConjunctionContext {
	var p = new(ConjunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_conjunction
	return p
}

func (*ConjunctionContext) IsConjunctionContext() {}

func NewConjunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConjunctionContext {
	var p = new(ConjunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_conjunction

	return p
}

func (s *ConjunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConjunctionContext) AllNegation() []INegationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INegationContext)(nil)).Elem())
	var tst = make([]INegationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INegationContext)
		}
	}

	return tst
}

func (s *ConjunctionContext) Negation(i int) INegationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INegationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INegationContext)
}

func (s *ConjunctionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(folParserAND)
}

func (s *ConjunctionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(folParserAND, i)
}

func (s *ConjunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConjunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConjunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterConjunction(s)
	}
}

func (s *ConjunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitConjunction(s)
	}
}

func (p *folParser) Conjunction() (localctx IConjunctionContext) {
	localctx = NewConjunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, folParserRULE_conjunction)
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
		p.SetState(37)
		p.Negation()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == folParserAND {
		{
			p.SetState(38)
			p.Match(folParserAND)
		}
		{
			p.SetState(39)
			p.Negation()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INegationContext is an interface to support dynamic dispatch.
type INegationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNegationContext differentiates from other interfaces.
	IsNegationContext()
}

type NegationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNegationContext() *NegationContext {
	var p = new(NegationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_negation
	return p
}

func (*NegationContext) IsNegationContext() {}

func NewNegationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NegationContext {
	var p = new(NegationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_negation

	return p
}

func (s *NegationContext) GetParser() antlr.Parser { return s.parser }

func (s *NegationContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *NegationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(folParserLPAREN, 0)
}

func (s *NegationContext) Formula() IFormulaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormulaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormulaContext)
}

func (s *NegationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(folParserRPAREN, 0)
}

func (s *NegationContext) NOT() antlr.TerminalNode {
	return s.GetToken(folParserNOT, 0)
}

func (s *NegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterNegation(s)
	}
}

func (s *NegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitNegation(s)
	}
}

func (p *folParser) Negation() (localctx INegationContext) {
	localctx = NewNegationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, folParserRULE_negation)
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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == folParserNOT {
		{
			p.SetState(45)
			p.Match(folParserNOT)
		}

	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case folParserPREPOSITION:
		{
			p.SetState(48)
			p.Predicate()
		}

	case folParserLPAREN:
		{
			p.SetState(49)
			p.Match(folParserLPAREN)
		}
		{
			p.SetState(50)
			p.Formula()
		}
		{
			p.SetState(51)
			p.Match(folParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) PREPOSITION() antlr.TerminalNode {
	return s.GetToken(folParserPREPOSITION, 0)
}

func (s *PredicateContext) PredicateTuple() IPredicateTupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateTupleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateTupleContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *folParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, folParserRULE_predicate)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Match(folParserPREPOSITION)
		}
		{
			p.SetState(56)
			p.PredicateTuple()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(folParserPREPOSITION)
		}

	}

	return localctx
}

// IPredicateTupleContext is an interface to support dynamic dispatch.
type IPredicateTupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateTupleContext differentiates from other interfaces.
	IsPredicateTupleContext()
}

type PredicateTupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateTupleContext() *PredicateTupleContext {
	var p = new(PredicateTupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_predicateTuple
	return p
}

func (*PredicateTupleContext) IsPredicateTupleContext() {}

func NewPredicateTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateTupleContext {
	var p = new(PredicateTupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_predicateTuple

	return p
}

func (s *PredicateTupleContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateTupleContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(folParserLPAREN, 0)
}

func (s *PredicateTupleContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *PredicateTupleContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *PredicateTupleContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(folParserRPAREN, 0)
}

func (s *PredicateTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateTupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateTupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterPredicateTuple(s)
	}
}

func (s *PredicateTupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitPredicateTuple(s)
	}
}

func (p *folParser) PredicateTuple() (localctx IPredicateTupleContext) {
	localctx = NewPredicateTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, folParserRULE_predicateTuple)
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
		p.SetState(60)
		p.Match(folParserLPAREN)
	}
	{
		p.SetState(61)
		p.Term()
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == folParserT__0 {
		{
			p.SetState(62)
			p.Match(folParserT__0)
		}
		{
			p.SetState(63)
			p.Term()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(69)
		p.Match(folParserRPAREN)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *TermContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(folParserVARIABLE, 0)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *folParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, folParserRULE_term)

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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case folParserCONSTANT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Function()
		}

	case folParserVARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(folParserVARIABLE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) CONSTANT() antlr.TerminalNode {
	return s.GetToken(folParserCONSTANT, 0)
}

func (s *FunctionContext) FunctionTuple() IFunctionTupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionTupleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionTupleContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *folParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, folParserRULE_function)

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

	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Match(folParserCONSTANT)
		}
		{
			p.SetState(76)
			p.FunctionTuple()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Match(folParserCONSTANT)
		}

	}

	return localctx
}

// IFunctionTupleContext is an interface to support dynamic dispatch.
type IFunctionTupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionTupleContext differentiates from other interfaces.
	IsFunctionTupleContext()
}

type FunctionTupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionTupleContext() *FunctionTupleContext {
	var p = new(FunctionTupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_functionTuple
	return p
}

func (*FunctionTupleContext) IsFunctionTupleContext() {}

func NewFunctionTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTupleContext {
	var p = new(FunctionTupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_functionTuple

	return p
}

func (s *FunctionTupleContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTupleContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(folParserLPAREN, 0)
}

func (s *FunctionTupleContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(folParserRPAREN, 0)
}

func (s *FunctionTupleContext) AllCONSTANT() []antlr.TerminalNode {
	return s.GetTokens(folParserCONSTANT)
}

func (s *FunctionTupleContext) CONSTANT(i int) antlr.TerminalNode {
	return s.GetToken(folParserCONSTANT, i)
}

func (s *FunctionTupleContext) AllVARIABLE() []antlr.TerminalNode {
	return s.GetTokens(folParserVARIABLE)
}

func (s *FunctionTupleContext) VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(folParserVARIABLE, i)
}

func (s *FunctionTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionTupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterFunctionTuple(s)
	}
}

func (s *FunctionTupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitFunctionTuple(s)
	}
}

func (p *folParser) FunctionTuple() (localctx IFunctionTupleContext) {
	localctx = NewFunctionTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, folParserRULE_functionTuple)
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
		p.SetState(80)
		p.Match(folParserLPAREN)
	}
	p.SetState(81)
	_la = p.GetTokenStream().LA(1)

	if !(_la == folParserVARIABLE || _la == folParserCONSTANT) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == folParserT__0 {
		{
			p.SetState(82)
			p.Match(folParserT__0)
		}
		p.SetState(83)
		_la = p.GetTokenStream().LA(1)

		if !(_la == folParserVARIABLE || _la == folParserCONSTANT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
		p.Match(folParserRPAREN)
	}

	return localctx
}
