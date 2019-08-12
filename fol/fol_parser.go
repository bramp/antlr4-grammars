// Code generated from fol.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 134,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 7, 2, 24, 10, 2,
	12, 2, 14, 2, 27, 11, 2, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 63, 10, 3, 12, 3, 14, 3, 66, 11, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 74, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	7, 3, 80, 10, 3, 12, 3, 14, 3, 83, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 7, 4, 93, 10, 4, 12, 4, 14, 4, 96, 11, 4, 3, 4, 3, 4,
	5, 4, 100, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 106, 10, 6, 12, 6, 14,
	6, 109, 11, 6, 3, 7, 3, 7, 7, 7, 113, 10, 7, 12, 7, 14, 7, 116, 11, 7,
	3, 8, 3, 8, 7, 8, 120, 10, 8, 12, 8, 14, 8, 123, 11, 8, 3, 9, 3, 9, 7,
	9, 127, 10, 9, 12, 9, 14, 9, 130, 11, 9, 3, 10, 3, 10, 3, 10, 2, 3, 4,
	11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 3, 3, 2, 15, 18, 2, 140, 2, 20,
	3, 2, 2, 2, 4, 73, 3, 2, 2, 2, 6, 99, 3, 2, 2, 2, 8, 101, 3, 2, 2, 2, 10,
	103, 3, 2, 2, 2, 12, 110, 3, 2, 2, 2, 14, 117, 3, 2, 2, 2, 16, 124, 3,
	2, 2, 2, 18, 131, 3, 2, 2, 2, 20, 25, 5, 4, 3, 2, 21, 22, 7, 19, 2, 2,
	22, 24, 5, 4, 3, 2, 23, 21, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3,
	2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 31, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28,
	30, 7, 19, 2, 2, 29, 28, 3, 2, 2, 2, 30, 33, 3, 2, 2, 2, 31, 29, 3, 2,
	2, 2, 31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 34, 35,
	7, 2, 2, 3, 35, 3, 3, 2, 2, 2, 36, 37, 8, 3, 1, 2, 37, 38, 7, 11, 2, 2,
	38, 39, 5, 4, 3, 2, 39, 40, 5, 8, 5, 2, 40, 41, 5, 4, 3, 8, 41, 74, 3,
	2, 2, 2, 42, 43, 7, 11, 2, 2, 43, 74, 5, 4, 3, 7, 44, 45, 7, 12, 2, 2,
	45, 46, 7, 8, 2, 2, 46, 47, 5, 10, 6, 2, 47, 48, 7, 9, 2, 2, 48, 49, 5,
	4, 3, 6, 49, 74, 3, 2, 2, 2, 50, 51, 7, 13, 2, 2, 51, 52, 7, 8, 2, 2, 52,
	53, 5, 10, 6, 2, 53, 54, 7, 9, 2, 2, 54, 55, 5, 4, 3, 5, 55, 74, 3, 2,
	2, 2, 56, 57, 5, 12, 7, 2, 57, 58, 7, 8, 2, 2, 58, 64, 5, 6, 4, 2, 59,
	60, 5, 18, 10, 2, 60, 61, 5, 6, 4, 2, 61, 63, 3, 2, 2, 2, 62, 59, 3, 2,
	2, 2, 63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67,
	3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 68, 7, 9, 2, 2, 68, 74, 3, 2, 2, 2,
	69, 70, 5, 6, 4, 2, 70, 71, 7, 10, 2, 2, 71, 72, 5, 6, 4, 2, 72, 74, 3,
	2, 2, 2, 73, 36, 3, 2, 2, 2, 73, 42, 3, 2, 2, 2, 73, 44, 3, 2, 2, 2, 73,
	50, 3, 2, 2, 2, 73, 56, 3, 2, 2, 2, 73, 69, 3, 2, 2, 2, 74, 81, 3, 2, 2,
	2, 75, 76, 12, 9, 2, 2, 76, 77, 5, 8, 5, 2, 77, 78, 5, 4, 3, 10, 78, 80,
	3, 2, 2, 2, 79, 75, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2,
	81, 82, 3, 2, 2, 2, 82, 5, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 100, 5,
	14, 8, 2, 85, 100, 5, 10, 6, 2, 86, 87, 5, 16, 9, 2, 87, 88, 7, 8, 2, 2,
	88, 94, 5, 6, 4, 2, 89, 90, 5, 18, 10, 2, 90, 91, 5, 6, 4, 2, 91, 93, 3,
	2, 2, 2, 92, 89, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94,
	95, 3, 2, 2, 2, 95, 97, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 98, 7, 9, 2,
	2, 98, 100, 3, 2, 2, 2, 99, 84, 3, 2, 2, 2, 99, 85, 3, 2, 2, 2, 99, 86,
	3, 2, 2, 2, 100, 7, 3, 2, 2, 2, 101, 102, 9, 2, 2, 2, 102, 9, 3, 2, 2,
	2, 103, 107, 7, 3, 2, 2, 104, 106, 7, 14, 2, 2, 105, 104, 3, 2, 2, 2, 106,
	109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 11, 3,
	2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 114, 7, 4, 2, 2, 111, 113, 7, 14, 2,
	2, 112, 111, 3, 2, 2, 2, 113, 116, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114,
	115, 3, 2, 2, 2, 115, 13, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 117, 121, 7,
	5, 2, 2, 118, 120, 7, 14, 2, 2, 119, 118, 3, 2, 2, 2, 120, 123, 3, 2, 2,
	2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 15, 3, 2, 2, 2, 123,
	121, 3, 2, 2, 2, 124, 128, 7, 6, 2, 2, 125, 127, 7, 14, 2, 2, 126, 125,
	3, 2, 2, 2, 127, 130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2,
	2, 2, 129, 17, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132, 7, 7, 2, 2,
	132, 19, 3, 2, 2, 2, 13, 25, 31, 64, 73, 81, 94, 99, 107, 114, 121, 128,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'?'", "'_'", "'#'", "'.'", "','", "'('", "')'", "'='", "'!'", "'Forall'",
	"'Exists'", "", "'\\/'", "'^'", "'->'", "'<->'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "LPAREN", "RPAREN", "EQUAL", "NOT", "FORALL", "EXISTS",
	"CHARACTER", "CONJ", "DISJ", "IMPL", "BICOND", "ENDLINE", "WHITESPACE",
}

var ruleNames = []string{
	"condition", "formula", "term", "bin_connective", "variable", "pred_constant",
	"ind_constant", "func_constant", "separator",
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
	folParserEOF        = antlr.TokenEOF
	folParserT__0       = 1
	folParserT__1       = 2
	folParserT__2       = 3
	folParserT__3       = 4
	folParserT__4       = 5
	folParserLPAREN     = 6
	folParserRPAREN     = 7
	folParserEQUAL      = 8
	folParserNOT        = 9
	folParserFORALL     = 10
	folParserEXISTS     = 11
	folParserCHARACTER  = 12
	folParserCONJ       = 13
	folParserDISJ       = 14
	folParserIMPL       = 15
	folParserBICOND     = 16
	folParserENDLINE    = 17
	folParserWHITESPACE = 18
)

// folParser rules.
const (
	folParserRULE_condition      = 0
	folParserRULE_formula        = 1
	folParserRULE_term           = 2
	folParserRULE_bin_connective = 3
	folParserRULE_variable       = 4
	folParserRULE_pred_constant  = 5
	folParserRULE_ind_constant   = 6
	folParserRULE_func_constant  = 7
	folParserRULE_separator      = 8
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

func (s *ConditionContext) AllFormula() []IFormulaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormulaContext)(nil)).Elem())
	var tst = make([]IFormulaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormulaContext)
		}
	}

	return tst
}

func (s *ConditionContext) Formula(i int) IFormulaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormulaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormulaContext)
}

func (s *ConditionContext) EOF() antlr.TerminalNode {
	return s.GetToken(folParserEOF, 0)
}

func (s *ConditionContext) AllENDLINE() []antlr.TerminalNode {
	return s.GetTokens(folParserENDLINE)
}

func (s *ConditionContext) ENDLINE(i int) antlr.TerminalNode {
	return s.GetToken(folParserENDLINE, i)
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
	{
		p.SetState(18)
		p.formula(0)
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(19)
				p.Match(folParserENDLINE)
			}
			{
				p.SetState(20)
				p.formula(0)
			}

		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == folParserENDLINE {
		{
			p.SetState(26)
			p.Match(folParserENDLINE)
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
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

func (s *FormulaContext) NOT() antlr.TerminalNode {
	return s.GetToken(folParserNOT, 0)
}

func (s *FormulaContext) AllFormula() []IFormulaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormulaContext)(nil)).Elem())
	var tst = make([]IFormulaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormulaContext)
		}
	}

	return tst
}

func (s *FormulaContext) Formula(i int) IFormulaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormulaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormulaContext)
}

func (s *FormulaContext) Bin_connective() IBin_connectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBin_connectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBin_connectiveContext)
}

func (s *FormulaContext) FORALL() antlr.TerminalNode {
	return s.GetToken(folParserFORALL, 0)
}

func (s *FormulaContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(folParserLPAREN, 0)
}

func (s *FormulaContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FormulaContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(folParserRPAREN, 0)
}

func (s *FormulaContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(folParserEXISTS, 0)
}

func (s *FormulaContext) Pred_constant() IPred_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPred_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPred_constantContext)
}

func (s *FormulaContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *FormulaContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *FormulaContext) AllSeparator() []ISeparatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISeparatorContext)(nil)).Elem())
	var tst = make([]ISeparatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISeparatorContext)
		}
	}

	return tst
}

func (s *FormulaContext) Separator(i int) ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *FormulaContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(folParserEQUAL, 0)
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
	return p.formula(0)
}

func (p *folParser) formula(_p int) (localctx IFormulaContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFormulaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFormulaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, folParserRULE_formula, _p)
	var _la int

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
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(35)
			p.Match(folParserNOT)
		}
		{
			p.SetState(36)
			p.formula(0)
		}
		{
			p.SetState(37)
			p.Bin_connective()
		}
		{
			p.SetState(38)
			p.formula(6)
		}

	case 2:
		{
			p.SetState(40)
			p.Match(folParserNOT)
		}
		{
			p.SetState(41)
			p.formula(5)
		}

	case 3:
		{
			p.SetState(42)
			p.Match(folParserFORALL)
		}
		{
			p.SetState(43)
			p.Match(folParserLPAREN)
		}
		{
			p.SetState(44)
			p.Variable()
		}
		{
			p.SetState(45)
			p.Match(folParserRPAREN)
		}
		{
			p.SetState(46)
			p.formula(4)
		}

	case 4:
		{
			p.SetState(48)
			p.Match(folParserEXISTS)
		}
		{
			p.SetState(49)
			p.Match(folParserLPAREN)
		}
		{
			p.SetState(50)
			p.Variable()
		}
		{
			p.SetState(51)
			p.Match(folParserRPAREN)
		}
		{
			p.SetState(52)
			p.formula(3)
		}

	case 5:
		{
			p.SetState(54)
			p.Pred_constant()
		}
		{
			p.SetState(55)
			p.Match(folParserLPAREN)
		}
		{
			p.SetState(56)
			p.Term()
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == folParserT__4 {
			{
				p.SetState(57)
				p.Separator()
			}
			{
				p.SetState(58)
				p.Term()
			}

			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(65)
			p.Match(folParserRPAREN)
		}

	case 6:
		{
			p.SetState(67)
			p.Term()
		}
		{
			p.SetState(68)
			p.Match(folParserEQUAL)
		}
		{
			p.SetState(69)
			p.Term()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFormulaContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, folParserRULE_formula)
			p.SetState(73)

			if !(p.Precpred(p.GetParserRuleContext(), 7)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
			}
			{
				p.SetState(74)
				p.Bin_connective()
			}
			{
				p.SetState(75)
				p.formula(8)
			}

		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
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

func (s *TermContext) Ind_constant() IInd_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInd_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInd_constantContext)
}

func (s *TermContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *TermContext) Func_constant() IFunc_constantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_constantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_constantContext)
}

func (s *TermContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(folParserLPAREN, 0)
}

func (s *TermContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *TermContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(folParserRPAREN, 0)
}

func (s *TermContext) AllSeparator() []ISeparatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISeparatorContext)(nil)).Elem())
	var tst = make([]ISeparatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISeparatorContext)
		}
	}

	return tst
}

func (s *TermContext) Separator(i int) ISeparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeparatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
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
	p.EnterRule(localctx, 4, folParserRULE_term)
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

	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case folParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Ind_constant()
		}

	case folParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Variable()
		}

	case folParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Func_constant()
		}
		{
			p.SetState(85)
			p.Match(folParserLPAREN)
		}
		{
			p.SetState(86)
			p.Term()
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == folParserT__4 {
			{
				p.SetState(87)
				p.Separator()
			}
			{
				p.SetState(88)
				p.Term()
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(95)
			p.Match(folParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBin_connectiveContext is an interface to support dynamic dispatch.
type IBin_connectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBin_connectiveContext differentiates from other interfaces.
	IsBin_connectiveContext()
}

type Bin_connectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBin_connectiveContext() *Bin_connectiveContext {
	var p = new(Bin_connectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_bin_connective
	return p
}

func (*Bin_connectiveContext) IsBin_connectiveContext() {}

func NewBin_connectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bin_connectiveContext {
	var p = new(Bin_connectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_bin_connective

	return p
}

func (s *Bin_connectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Bin_connectiveContext) CONJ() antlr.TerminalNode {
	return s.GetToken(folParserCONJ, 0)
}

func (s *Bin_connectiveContext) DISJ() antlr.TerminalNode {
	return s.GetToken(folParserDISJ, 0)
}

func (s *Bin_connectiveContext) IMPL() antlr.TerminalNode {
	return s.GetToken(folParserIMPL, 0)
}

func (s *Bin_connectiveContext) BICOND() antlr.TerminalNode {
	return s.GetToken(folParserBICOND, 0)
}

func (s *Bin_connectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bin_connectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bin_connectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterBin_connective(s)
	}
}

func (s *Bin_connectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitBin_connective(s)
	}
}

func (p *folParser) Bin_connective() (localctx IBin_connectiveContext) {
	localctx = NewBin_connectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, folParserRULE_bin_connective)
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
		p.SetState(99)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<folParserCONJ)|(1<<folParserDISJ)|(1<<folParserIMPL)|(1<<folParserBICOND))) != 0) {
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
	p.RuleIndex = folParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) AllCHARACTER() []antlr.TerminalNode {
	return s.GetTokens(folParserCHARACTER)
}

func (s *VariableContext) CHARACTER(i int) antlr.TerminalNode {
	return s.GetToken(folParserCHARACTER, i)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *folParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, folParserRULE_variable)

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
	{
		p.SetState(101)
		p.Match(folParserT__0)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(102)
				p.Match(folParserCHARACTER)
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IPred_constantContext is an interface to support dynamic dispatch.
type IPred_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPred_constantContext differentiates from other interfaces.
	IsPred_constantContext()
}

type Pred_constantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPred_constantContext() *Pred_constantContext {
	var p = new(Pred_constantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_pred_constant
	return p
}

func (*Pred_constantContext) IsPred_constantContext() {}

func NewPred_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pred_constantContext {
	var p = new(Pred_constantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_pred_constant

	return p
}

func (s *Pred_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Pred_constantContext) AllCHARACTER() []antlr.TerminalNode {
	return s.GetTokens(folParserCHARACTER)
}

func (s *Pred_constantContext) CHARACTER(i int) antlr.TerminalNode {
	return s.GetToken(folParserCHARACTER, i)
}

func (s *Pred_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pred_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pred_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterPred_constant(s)
	}
}

func (s *Pred_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitPred_constant(s)
	}
}

func (p *folParser) Pred_constant() (localctx IPred_constantContext) {
	localctx = NewPred_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, folParserRULE_pred_constant)
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
		p.SetState(108)
		p.Match(folParserT__1)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == folParserCHARACTER {
		{
			p.SetState(109)
			p.Match(folParserCHARACTER)
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInd_constantContext is an interface to support dynamic dispatch.
type IInd_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInd_constantContext differentiates from other interfaces.
	IsInd_constantContext()
}

type Ind_constantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInd_constantContext() *Ind_constantContext {
	var p = new(Ind_constantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_ind_constant
	return p
}

func (*Ind_constantContext) IsInd_constantContext() {}

func NewInd_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ind_constantContext {
	var p = new(Ind_constantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_ind_constant

	return p
}

func (s *Ind_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Ind_constantContext) AllCHARACTER() []antlr.TerminalNode {
	return s.GetTokens(folParserCHARACTER)
}

func (s *Ind_constantContext) CHARACTER(i int) antlr.TerminalNode {
	return s.GetToken(folParserCHARACTER, i)
}

func (s *Ind_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ind_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ind_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterInd_constant(s)
	}
}

func (s *Ind_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitInd_constant(s)
	}
}

func (p *folParser) Ind_constant() (localctx IInd_constantContext) {
	localctx = NewInd_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, folParserRULE_ind_constant)

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
	{
		p.SetState(115)
		p.Match(folParserT__2)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(116)
				p.Match(folParserCHARACTER)
			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IFunc_constantContext is an interface to support dynamic dispatch.
type IFunc_constantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_constantContext differentiates from other interfaces.
	IsFunc_constantContext()
}

type Func_constantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_constantContext() *Func_constantContext {
	var p = new(Func_constantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_func_constant
	return p
}

func (*Func_constantContext) IsFunc_constantContext() {}

func NewFunc_constantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_constantContext {
	var p = new(Func_constantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_func_constant

	return p
}

func (s *Func_constantContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_constantContext) AllCHARACTER() []antlr.TerminalNode {
	return s.GetTokens(folParserCHARACTER)
}

func (s *Func_constantContext) CHARACTER(i int) antlr.TerminalNode {
	return s.GetToken(folParserCHARACTER, i)
}

func (s *Func_constantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_constantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_constantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterFunc_constant(s)
	}
}

func (s *Func_constantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitFunc_constant(s)
	}
}

func (p *folParser) Func_constant() (localctx IFunc_constantContext) {
	localctx = NewFunc_constantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, folParserRULE_func_constant)
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
		p.SetState(122)
		p.Match(folParserT__3)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == folParserCHARACTER {
		{
			p.SetState(123)
			p.Match(folParserCHARACTER)
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISeparatorContext is an interface to support dynamic dispatch.
type ISeparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSeparatorContext differentiates from other interfaces.
	IsSeparatorContext()
}

type SeparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeparatorContext() *SeparatorContext {
	var p = new(SeparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = folParserRULE_separator
	return p
}

func (*SeparatorContext) IsSeparatorContext() {}

func NewSeparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeparatorContext {
	var p = new(SeparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = folParserRULE_separator

	return p
}

func (s *SeparatorContext) GetParser() antlr.Parser { return s.parser }
func (s *SeparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.EnterSeparator(s)
	}
}

func (s *SeparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(folListener); ok {
		listenerT.ExitSeparator(s)
	}
}

func (p *folParser) Separator() (localctx ISeparatorContext) {
	localctx = NewSeparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, folParserRULE_separator)

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
		p.SetState(129)
		p.Match(folParserT__4)
	}

	return localctx
}

func (p *folParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *FormulaContext = nil
		if localctx != nil {
			t = localctx.(*FormulaContext)
		}
		return p.Formula_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *folParser) Formula_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
