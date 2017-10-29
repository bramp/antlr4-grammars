// Generated from propcalc.g4 by ANTLR 4.7.

package propcalc // propcalc
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 67, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 26,
	10, 3, 12, 3, 14, 3, 29, 11, 3, 3, 4, 3, 4, 3, 4, 5, 4, 34, 10, 4, 3, 5,
	3, 5, 3, 5, 7, 5, 39, 10, 5, 12, 5, 14, 5, 42, 11, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 51, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 7, 9, 62, 10, 9, 12, 9, 14, 9, 65, 11, 9, 3, 9,
	2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 3, 3, 2, 4, 5, 2, 65, 2, 18, 3,
	2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 33, 3, 2, 2, 2, 8, 35, 3, 2, 2, 2, 10, 50,
	3, 2, 2, 2, 12, 52, 3, 2, 2, 2, 14, 56, 3, 2, 2, 2, 16, 63, 3, 2, 2, 2,
	18, 19, 5, 4, 3, 2, 19, 20, 7, 9, 2, 2, 20, 21, 5, 4, 3, 2, 21, 3, 3, 2,
	2, 2, 22, 27, 5, 6, 4, 2, 23, 24, 9, 2, 2, 2, 24, 26, 5, 6, 4, 2, 25, 23,
	3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2,
	28, 5, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 34, 5, 10, 6, 2, 31, 34, 5,
	12, 7, 2, 32, 34, 5, 14, 8, 2, 33, 30, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2,
	33, 32, 3, 2, 2, 2, 34, 7, 3, 2, 2, 2, 35, 40, 5, 10, 6, 2, 36, 37, 7,
	3, 2, 2, 37, 39, 5, 10, 6, 2, 38, 36, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40,
	38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 9, 3, 2, 2, 2, 42, 40, 3, 2, 2,
	2, 43, 51, 5, 16, 9, 2, 44, 45, 7, 6, 2, 2, 45, 51, 5, 10, 6, 2, 46, 47,
	7, 11, 2, 2, 47, 48, 5, 4, 3, 2, 48, 49, 7, 12, 2, 2, 49, 51, 3, 2, 2,
	2, 50, 43, 3, 2, 2, 2, 50, 44, 3, 2, 2, 2, 50, 46, 3, 2, 2, 2, 51, 11,
	3, 2, 2, 2, 52, 53, 5, 10, 6, 2, 53, 54, 7, 10, 2, 2, 54, 55, 5, 10, 6,
	2, 55, 13, 3, 2, 2, 2, 56, 57, 5, 10, 6, 2, 57, 58, 7, 8, 2, 2, 58, 59,
	5, 10, 6, 2, 59, 15, 3, 2, 2, 2, 60, 62, 7, 13, 2, 2, 61, 60, 3, 2, 2,
	2, 62, 65, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 17,
	3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 7, 27, 33, 40, 50, 63,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'^'", "'v'", "'!'", "'='", "'->'", "'|-'", "'<->'", "'('",
	"')'",
}
var symbolicNames = []string{
	"", "", "AND", "OR", "NOT", "EQ", "IMPLIES", "THEREFORE", "EQUIV", "LPAREN",
	"RPAREN", "LETTER", "WS",
}

var ruleNames = []string{
	"proposition", "expression", "relExpression", "atoms", "atom", "equiv",
	"implies", "variable",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type propcalcParser struct {
	*antlr.BaseParser
}

func NewpropcalcParser(input antlr.TokenStream) *propcalcParser {
	this := new(propcalcParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "propcalc.g4"

	return this
}

// propcalcParser tokens.
const (
	propcalcParserEOF       = antlr.TokenEOF
	propcalcParserT__0      = 1
	propcalcParserAND       = 2
	propcalcParserOR        = 3
	propcalcParserNOT       = 4
	propcalcParserEQ        = 5
	propcalcParserIMPLIES   = 6
	propcalcParserTHEREFORE = 7
	propcalcParserEQUIV     = 8
	propcalcParserLPAREN    = 9
	propcalcParserRPAREN    = 10
	propcalcParserLETTER    = 11
	propcalcParserWS        = 12
)

// propcalcParser rules.
const (
	propcalcParserRULE_proposition   = 0
	propcalcParserRULE_expression    = 1
	propcalcParserRULE_relExpression = 2
	propcalcParserRULE_atoms         = 3
	propcalcParserRULE_atom          = 4
	propcalcParserRULE_equiv         = 5
	propcalcParserRULE_implies       = 6
	propcalcParserRULE_variable      = 7
)

// IPropositionContext is an interface to support dynamic dispatch.
type IPropositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropositionContext differentiates from other interfaces.
	IsPropositionContext()
}

type PropositionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropositionContext() *PropositionContext {
	var p = new(PropositionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = propcalcParserRULE_proposition
	return p
}

func (*PropositionContext) IsPropositionContext() {}

func NewPropositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropositionContext {
	var p = new(PropositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = propcalcParserRULE_proposition

	return p
}

func (s *PropositionContext) GetParser() antlr.Parser { return s.parser }

func (s *PropositionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *PropositionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PropositionContext) THEREFORE() antlr.TerminalNode {
	return s.GetToken(propcalcParserTHEREFORE, 0)
}

func (s *PropositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.EnterProposition(s)
	}
}

func (s *PropositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.ExitProposition(s)
	}
}

func (p *propcalcParser) Proposition() (localctx IPropositionContext) {
	localctx = NewPropositionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, propcalcParserRULE_proposition)

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
		p.SetState(16)
		p.Expression()
	}
	{
		p.SetState(17)
		p.Match(propcalcParserTHEREFORE)
	}
	{
		p.SetState(18)
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
	p.RuleIndex = propcalcParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = propcalcParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllRelExpression() []IRelExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelExpressionContext)(nil)).Elem())
	var tst = make([]IRelExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) RelExpression(i int) IRelExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelExpressionContext)
}

func (s *ExpressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(propcalcParserAND)
}

func (s *ExpressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(propcalcParserAND, i)
}

func (s *ExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(propcalcParserOR)
}

func (s *ExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(propcalcParserOR, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *propcalcParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, propcalcParserRULE_expression)
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
		p.SetState(20)
		p.RelExpression()
	}
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == propcalcParserAND || _la == propcalcParserOR {
		p.SetState(21)
		_la = p.GetTokenStream().LA(1)

		if !(_la == propcalcParserAND || _la == propcalcParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(22)
			p.RelExpression()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRelExpressionContext is an interface to support dynamic dispatch.
type IRelExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelExpressionContext differentiates from other interfaces.
	IsRelExpressionContext()
}

type RelExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelExpressionContext() *RelExpressionContext {
	var p = new(RelExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = propcalcParserRULE_relExpression
	return p
}

func (*RelExpressionContext) IsRelExpressionContext() {}

func NewRelExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelExpressionContext {
	var p = new(RelExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = propcalcParserRULE_relExpression

	return p
}

func (s *RelExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelExpressionContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *RelExpressionContext) Equiv() IEquivContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEquivContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEquivContext)
}

func (s *RelExpressionContext) Implies() IImpliesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpliesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpliesContext)
}

func (s *RelExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.EnterRelExpression(s)
	}
}

func (s *RelExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.ExitRelExpression(s)
	}
}

func (p *propcalcParser) RelExpression() (localctx IRelExpressionContext) {
	localctx = NewRelExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, propcalcParserRULE_relExpression)

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

	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Atom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.Equiv()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
			p.Implies()
		}

	}

	return localctx
}

// IAtomsContext is an interface to support dynamic dispatch.
type IAtomsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomsContext differentiates from other interfaces.
	IsAtomsContext()
}

type AtomsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomsContext() *AtomsContext {
	var p = new(AtomsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = propcalcParserRULE_atoms
	return p
}

func (*AtomsContext) IsAtomsContext() {}

func NewAtomsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomsContext {
	var p = new(AtomsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = propcalcParserRULE_atoms

	return p
}

func (s *AtomsContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomsContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *AtomsContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.EnterAtoms(s)
	}
}

func (s *AtomsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.ExitAtoms(s)
	}
}

func (p *propcalcParser) Atoms() (localctx IAtomsContext) {
	localctx = NewAtomsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, propcalcParserRULE_atoms)
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
		p.SetState(33)
		p.Atom()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == propcalcParserT__0 {
		{
			p.SetState(34)
			p.Match(propcalcParserT__0)
		}
		{
			p.SetState(35)
			p.Atom()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = propcalcParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = propcalcParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AtomContext) NOT() antlr.TerminalNode {
	return s.GetToken(propcalcParserNOT, 0)
}

func (s *AtomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(propcalcParserLPAREN, 0)
}

func (s *AtomContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(propcalcParserRPAREN, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *propcalcParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, propcalcParserRULE_atom)

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

	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case propcalcParserEOF, propcalcParserT__0, propcalcParserAND, propcalcParserOR, propcalcParserIMPLIES, propcalcParserTHEREFORE, propcalcParserEQUIV, propcalcParserRPAREN, propcalcParserLETTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Variable()
		}

	case propcalcParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(propcalcParserNOT)
		}
		{
			p.SetState(43)
			p.Atom()
		}

	case propcalcParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Match(propcalcParserLPAREN)
		}
		{
			p.SetState(45)
			p.Expression()
		}
		{
			p.SetState(46)
			p.Match(propcalcParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEquivContext is an interface to support dynamic dispatch.
type IEquivContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEquivContext differentiates from other interfaces.
	IsEquivContext()
}

type EquivContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquivContext() *EquivContext {
	var p = new(EquivContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = propcalcParserRULE_equiv
	return p
}

func (*EquivContext) IsEquivContext() {}

func NewEquivContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquivContext {
	var p = new(EquivContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = propcalcParserRULE_equiv

	return p
}

func (s *EquivContext) GetParser() antlr.Parser { return s.parser }

func (s *EquivContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *EquivContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *EquivContext) EQUIV() antlr.TerminalNode {
	return s.GetToken(propcalcParserEQUIV, 0)
}

func (s *EquivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquivContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EquivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.EnterEquiv(s)
	}
}

func (s *EquivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.ExitEquiv(s)
	}
}

func (p *propcalcParser) Equiv() (localctx IEquivContext) {
	localctx = NewEquivContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, propcalcParserRULE_equiv)

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
		p.SetState(50)
		p.Atom()
	}
	{
		p.SetState(51)
		p.Match(propcalcParserEQUIV)
	}
	{
		p.SetState(52)
		p.Atom()
	}

	return localctx
}

// IImpliesContext is an interface to support dynamic dispatch.
type IImpliesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImpliesContext differentiates from other interfaces.
	IsImpliesContext()
}

type ImpliesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImpliesContext() *ImpliesContext {
	var p = new(ImpliesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = propcalcParserRULE_implies
	return p
}

func (*ImpliesContext) IsImpliesContext() {}

func NewImpliesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpliesContext {
	var p = new(ImpliesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = propcalcParserRULE_implies

	return p
}

func (s *ImpliesContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpliesContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *ImpliesContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ImpliesContext) IMPLIES() antlr.TerminalNode {
	return s.GetToken(propcalcParserIMPLIES, 0)
}

func (s *ImpliesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpliesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpliesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.EnterImplies(s)
	}
}

func (s *ImpliesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.ExitImplies(s)
	}
}

func (p *propcalcParser) Implies() (localctx IImpliesContext) {
	localctx = NewImpliesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, propcalcParserRULE_implies)

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
		p.SetState(54)
		p.Atom()
	}
	{
		p.SetState(55)
		p.Match(propcalcParserIMPLIES)
	}
	{
		p.SetState(56)
		p.Atom()
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
	p.RuleIndex = propcalcParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = propcalcParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) AllLETTER() []antlr.TerminalNode {
	return s.GetTokens(propcalcParserLETTER)
}

func (s *VariableContext) LETTER(i int) antlr.TerminalNode {
	return s.GetToken(propcalcParserLETTER, i)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(propcalcListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *propcalcParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, propcalcParserRULE_variable)
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == propcalcParserLETTER {
		{
			p.SetState(58)
			p.Match(propcalcParserLETTER)
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
