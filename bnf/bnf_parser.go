// Generated from bnf.g4 by ANTLR 4.7.

package bnf // bnf
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 86, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	7, 6, 48, 10, 6, 12, 6, 14, 6, 51, 11, 6, 3, 7, 7, 7, 54, 10, 7, 12, 7,
	14, 7, 57, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 64, 10, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 2, 2, 15,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 2, 2, 79, 2, 31, 3,
	2, 2, 2, 4, 36, 3, 2, 2, 2, 6, 40, 3, 2, 2, 2, 8, 42, 3, 2, 2, 2, 10, 44,
	3, 2, 2, 2, 12, 55, 3, 2, 2, 2, 14, 63, 3, 2, 2, 2, 16, 65, 3, 2, 2, 2,
	18, 69, 3, 2, 2, 2, 20, 73, 3, 2, 2, 2, 22, 77, 3, 2, 2, 2, 24, 79, 3,
	2, 2, 2, 26, 83, 3, 2, 2, 2, 28, 30, 5, 4, 3, 2, 29, 28, 3, 2, 2, 2, 30,
	33, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2,
	2, 33, 31, 3, 2, 2, 2, 34, 35, 7, 2, 2, 3, 35, 3, 3, 2, 2, 2, 36, 37, 5,
	6, 4, 2, 37, 38, 7, 3, 2, 2, 38, 39, 5, 8, 5, 2, 39, 5, 3, 2, 2, 2, 40,
	41, 5, 24, 13, 2, 41, 7, 3, 2, 2, 2, 42, 43, 5, 10, 6, 2, 43, 9, 3, 2,
	2, 2, 44, 49, 5, 12, 7, 2, 45, 46, 7, 10, 2, 2, 46, 48, 5, 12, 7, 2, 47,
	45, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2,
	2, 50, 11, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 54, 5, 14, 8, 2, 53, 52,
	3, 2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2,
	56, 13, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 64, 5, 16, 9, 2, 59, 64, 5,
	18, 10, 2, 60, 64, 5, 20, 11, 2, 61, 64, 5, 22, 12, 2, 62, 64, 5, 24, 13,
	2, 63, 58, 3, 2, 2, 2, 63, 59, 3, 2, 2, 2, 63, 60, 3, 2, 2, 2, 63, 61,
	3, 2, 2, 2, 63, 62, 3, 2, 2, 2, 64, 15, 3, 2, 2, 2, 65, 66, 7, 9, 2, 2,
	66, 67, 5, 10, 6, 2, 67, 68, 7, 8, 2, 2, 68, 17, 3, 2, 2, 2, 69, 70, 7,
	7, 2, 2, 70, 71, 5, 10, 6, 2, 71, 72, 7, 6, 2, 2, 72, 19, 3, 2, 2, 2, 73,
	74, 7, 5, 2, 2, 74, 75, 5, 10, 6, 2, 75, 76, 7, 4, 2, 2, 76, 21, 3, 2,
	2, 2, 77, 78, 7, 13, 2, 2, 78, 23, 3, 2, 2, 2, 79, 80, 7, 12, 2, 2, 80,
	81, 5, 26, 14, 2, 81, 82, 7, 11, 2, 2, 82, 25, 3, 2, 2, 2, 83, 84, 7, 13,
	2, 2, 84, 27, 3, 2, 2, 2, 6, 31, 49, 55, 63,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'::='", "')'", "'('", "'}'", "'{'", "']'", "'['", "'|'", "'>'", "'<'",
}
var symbolicNames = []string{
	"", "ASSIGN", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LEND", "REND", "BAR",
	"GT", "LT", "ID", "WS",
}

var ruleNames = []string{
	"rulelist", "rule_", "lhs", "rhs", "alternatives", "alternative", "element",
	"optional", "zeroormore", "oneormore", "text", "id", "ruleid",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type bnfParser struct {
	*antlr.BaseParser
}

func NewbnfParser(input antlr.TokenStream) *bnfParser {
	this := new(bnfParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "bnf.g4"

	return this
}

// bnfParser tokens.
const (
	bnfParserEOF    = antlr.TokenEOF
	bnfParserASSIGN = 1
	bnfParserLPAREN = 2
	bnfParserRPAREN = 3
	bnfParserLBRACE = 4
	bnfParserRBRACE = 5
	bnfParserLEND   = 6
	bnfParserREND   = 7
	bnfParserBAR    = 8
	bnfParserGT     = 9
	bnfParserLT     = 10
	bnfParserID     = 11
	bnfParserWS     = 12
)

// bnfParser rules.
const (
	bnfParserRULE_rulelist     = 0
	bnfParserRULE_rule_        = 1
	bnfParserRULE_lhs          = 2
	bnfParserRULE_rhs          = 3
	bnfParserRULE_alternatives = 4
	bnfParserRULE_alternative  = 5
	bnfParserRULE_element      = 6
	bnfParserRULE_optional     = 7
	bnfParserRULE_zeroormore   = 8
	bnfParserRULE_oneormore    = 9
	bnfParserRULE_text         = 10
	bnfParserRULE_id           = 11
	bnfParserRULE_ruleid       = 12
)

// IRulelistContext is an interface to support dynamic dispatch.
type IRulelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulelistContext differentiates from other interfaces.
	IsRulelistContext()
}

type RulelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulelistContext() *RulelistContext {
	var p = new(RulelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_rulelist
	return p
}

func (*RulelistContext) IsRulelistContext() {}

func NewRulelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulelistContext {
	var p = new(RulelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_rulelist

	return p
}

func (s *RulelistContext) GetParser() antlr.Parser { return s.parser }

func (s *RulelistContext) EOF() antlr.TerminalNode {
	return s.GetToken(bnfParserEOF, 0)
}

func (s *RulelistContext) AllRule_() []IRule_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRule_Context)(nil)).Elem())
	var tst = make([]IRule_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRule_Context)
		}
	}

	return tst
}

func (s *RulelistContext) Rule_(i int) IRule_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRule_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRule_Context)
}

func (s *RulelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterRulelist(s)
	}
}

func (s *RulelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitRulelist(s)
	}
}

func (p *bnfParser) Rulelist() (localctx IRulelistContext) {
	localctx = NewRulelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, bnfParserRULE_rulelist)
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
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == bnfParserLT {
		{
			p.SetState(26)
			p.Rule_()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(bnfParserEOF)
	}

	return localctx
}

// IRule_Context is an interface to support dynamic dispatch.
type IRule_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRule_Context differentiates from other interfaces.
	IsRule_Context()
}

type Rule_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_Context() *Rule_Context {
	var p = new(Rule_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_rule_
	return p
}

func (*Rule_Context) IsRule_Context() {}

func NewRule_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_Context {
	var p = new(Rule_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_rule_

	return p
}

func (s *Rule_Context) GetParser() antlr.Parser { return s.parser }

func (s *Rule_Context) Lhs() ILhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILhsContext)
}

func (s *Rule_Context) ASSIGN() antlr.TerminalNode {
	return s.GetToken(bnfParserASSIGN, 0)
}

func (s *Rule_Context) Rhs() IRhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRhsContext)
}

func (s *Rule_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterRule_(s)
	}
}

func (s *Rule_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitRule_(s)
	}
}

func (p *bnfParser) Rule_() (localctx IRule_Context) {
	localctx = NewRule_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, bnfParserRULE_rule_)

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
		p.SetState(34)
		p.Lhs()
	}
	{
		p.SetState(35)
		p.Match(bnfParserASSIGN)
	}
	{
		p.SetState(36)
		p.Rhs()
	}

	return localctx
}

// ILhsContext is an interface to support dynamic dispatch.
type ILhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLhsContext differentiates from other interfaces.
	IsLhsContext()
}

type LhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLhsContext() *LhsContext {
	var p = new(LhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_lhs
	return p
}

func (*LhsContext) IsLhsContext() {}

func NewLhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LhsContext {
	var p = new(LhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_lhs

	return p
}

func (s *LhsContext) GetParser() antlr.Parser { return s.parser }

func (s *LhsContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *LhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterLhs(s)
	}
}

func (s *LhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitLhs(s)
	}
}

func (p *bnfParser) Lhs() (localctx ILhsContext) {
	localctx = NewLhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, bnfParserRULE_lhs)

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
		p.Id()
	}

	return localctx
}

// IRhsContext is an interface to support dynamic dispatch.
type IRhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRhsContext differentiates from other interfaces.
	IsRhsContext()
}

type RhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhsContext() *RhsContext {
	var p = new(RhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_rhs
	return p
}

func (*RhsContext) IsRhsContext() {}

func NewRhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhsContext {
	var p = new(RhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_rhs

	return p
}

func (s *RhsContext) GetParser() antlr.Parser { return s.parser }

func (s *RhsContext) Alternatives() IAlternativesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlternativesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlternativesContext)
}

func (s *RhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterRhs(s)
	}
}

func (s *RhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitRhs(s)
	}
}

func (p *bnfParser) Rhs() (localctx IRhsContext) {
	localctx = NewRhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, bnfParserRULE_rhs)

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
		p.Alternatives()
	}

	return localctx
}

// IAlternativesContext is an interface to support dynamic dispatch.
type IAlternativesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlternativesContext differentiates from other interfaces.
	IsAlternativesContext()
}

type AlternativesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlternativesContext() *AlternativesContext {
	var p = new(AlternativesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_alternatives
	return p
}

func (*AlternativesContext) IsAlternativesContext() {}

func NewAlternativesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlternativesContext {
	var p = new(AlternativesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_alternatives

	return p
}

func (s *AlternativesContext) GetParser() antlr.Parser { return s.parser }

func (s *AlternativesContext) AllAlternative() []IAlternativeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAlternativeContext)(nil)).Elem())
	var tst = make([]IAlternativeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAlternativeContext)
		}
	}

	return tst
}

func (s *AlternativesContext) Alternative(i int) IAlternativeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlternativeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAlternativeContext)
}

func (s *AlternativesContext) AllBAR() []antlr.TerminalNode {
	return s.GetTokens(bnfParserBAR)
}

func (s *AlternativesContext) BAR(i int) antlr.TerminalNode {
	return s.GetToken(bnfParserBAR, i)
}

func (s *AlternativesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlternativesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlternativesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterAlternatives(s)
	}
}

func (s *AlternativesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitAlternatives(s)
	}
}

func (p *bnfParser) Alternatives() (localctx IAlternativesContext) {
	localctx = NewAlternativesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, bnfParserRULE_alternatives)
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
		p.SetState(42)
		p.Alternative()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == bnfParserBAR {
		{
			p.SetState(43)
			p.Match(bnfParserBAR)
		}
		{
			p.SetState(44)
			p.Alternative()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAlternativeContext is an interface to support dynamic dispatch.
type IAlternativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlternativeContext differentiates from other interfaces.
	IsAlternativeContext()
}

type AlternativeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlternativeContext() *AlternativeContext {
	var p = new(AlternativeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_alternative
	return p
}

func (*AlternativeContext) IsAlternativeContext() {}

func NewAlternativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlternativeContext {
	var p = new(AlternativeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_alternative

	return p
}

func (s *AlternativeContext) GetParser() antlr.Parser { return s.parser }

func (s *AlternativeContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *AlternativeContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *AlternativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlternativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlternativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterAlternative(s)
	}
}

func (s *AlternativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitAlternative(s)
	}
}

func (p *bnfParser) Alternative() (localctx IAlternativeContext) {
	localctx = NewAlternativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, bnfParserRULE_alternative)

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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(50)
				p.Element()
			}

		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) Optional() IOptionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalContext)
}

func (s *ElementContext) Zeroormore() IZeroormoreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IZeroormoreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IZeroormoreContext)
}

func (s *ElementContext) Oneormore() IOneormoreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneormoreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneormoreContext)
}

func (s *ElementContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *ElementContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitElement(s)
	}
}

func (p *bnfParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, bnfParserRULE_element)

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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case bnfParserREND:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Optional()
		}

	case bnfParserRBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Zeroormore()
		}

	case bnfParserRPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.Oneormore()
		}

	case bnfParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(59)
			p.Text()
		}

	case bnfParserLT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(60)
			p.Id()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOptionalContext is an interface to support dynamic dispatch.
type IOptionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionalContext differentiates from other interfaces.
	IsOptionalContext()
}

type OptionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalContext() *OptionalContext {
	var p = new(OptionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_optional
	return p
}

func (*OptionalContext) IsOptionalContext() {}

func NewOptionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalContext {
	var p = new(OptionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_optional

	return p
}

func (s *OptionalContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalContext) REND() antlr.TerminalNode {
	return s.GetToken(bnfParserREND, 0)
}

func (s *OptionalContext) Alternatives() IAlternativesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlternativesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlternativesContext)
}

func (s *OptionalContext) LEND() antlr.TerminalNode {
	return s.GetToken(bnfParserLEND, 0)
}

func (s *OptionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterOptional(s)
	}
}

func (s *OptionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitOptional(s)
	}
}

func (p *bnfParser) Optional() (localctx IOptionalContext) {
	localctx = NewOptionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, bnfParserRULE_optional)

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
		p.SetState(63)
		p.Match(bnfParserREND)
	}
	{
		p.SetState(64)
		p.Alternatives()
	}
	{
		p.SetState(65)
		p.Match(bnfParserLEND)
	}

	return localctx
}

// IZeroormoreContext is an interface to support dynamic dispatch.
type IZeroormoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZeroormoreContext differentiates from other interfaces.
	IsZeroormoreContext()
}

type ZeroormoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZeroormoreContext() *ZeroormoreContext {
	var p = new(ZeroormoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_zeroormore
	return p
}

func (*ZeroormoreContext) IsZeroormoreContext() {}

func NewZeroormoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZeroormoreContext {
	var p = new(ZeroormoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_zeroormore

	return p
}

func (s *ZeroormoreContext) GetParser() antlr.Parser { return s.parser }

func (s *ZeroormoreContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(bnfParserRBRACE, 0)
}

func (s *ZeroormoreContext) Alternatives() IAlternativesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlternativesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlternativesContext)
}

func (s *ZeroormoreContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(bnfParserLBRACE, 0)
}

func (s *ZeroormoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZeroormoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZeroormoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterZeroormore(s)
	}
}

func (s *ZeroormoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitZeroormore(s)
	}
}

func (p *bnfParser) Zeroormore() (localctx IZeroormoreContext) {
	localctx = NewZeroormoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, bnfParserRULE_zeroormore)

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
		p.Match(bnfParserRBRACE)
	}
	{
		p.SetState(68)
		p.Alternatives()
	}
	{
		p.SetState(69)
		p.Match(bnfParserLBRACE)
	}

	return localctx
}

// IOneormoreContext is an interface to support dynamic dispatch.
type IOneormoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneormoreContext differentiates from other interfaces.
	IsOneormoreContext()
}

type OneormoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneormoreContext() *OneormoreContext {
	var p = new(OneormoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_oneormore
	return p
}

func (*OneormoreContext) IsOneormoreContext() {}

func NewOneormoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneormoreContext {
	var p = new(OneormoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_oneormore

	return p
}

func (s *OneormoreContext) GetParser() antlr.Parser { return s.parser }

func (s *OneormoreContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(bnfParserRPAREN, 0)
}

func (s *OneormoreContext) Alternatives() IAlternativesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlternativesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlternativesContext)
}

func (s *OneormoreContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(bnfParserLPAREN, 0)
}

func (s *OneormoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneormoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneormoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterOneormore(s)
	}
}

func (s *OneormoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitOneormore(s)
	}
}

func (p *bnfParser) Oneormore() (localctx IOneormoreContext) {
	localctx = NewOneormoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, bnfParserRULE_oneormore)

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
		p.Match(bnfParserRPAREN)
	}
	{
		p.SetState(72)
		p.Alternatives()
	}
	{
		p.SetState(73)
		p.Match(bnfParserLPAREN)
	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) ID() antlr.TerminalNode {
	return s.GetToken(bnfParserID, 0)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitText(s)
	}
}

func (p *bnfParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, bnfParserRULE_text)

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
		p.Match(bnfParserID)
	}

	return localctx
}

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_id
	return p
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) LT() antlr.TerminalNode {
	return s.GetToken(bnfParserLT, 0)
}

func (s *IdContext) Ruleid() IRuleidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleidContext)
}

func (s *IdContext) GT() antlr.TerminalNode {
	return s.GetToken(bnfParserGT, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitId(s)
	}
}

func (p *bnfParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, bnfParserRULE_id)

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
		p.Match(bnfParserLT)
	}
	{
		p.SetState(78)
		p.Ruleid()
	}
	{
		p.SetState(79)
		p.Match(bnfParserGT)
	}

	return localctx
}

// IRuleidContext is an interface to support dynamic dispatch.
type IRuleidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleidContext differentiates from other interfaces.
	IsRuleidContext()
}

type RuleidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleidContext() *RuleidContext {
	var p = new(RuleidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bnfParserRULE_ruleid
	return p
}

func (*RuleidContext) IsRuleidContext() {}

func NewRuleidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleidContext {
	var p = new(RuleidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bnfParserRULE_ruleid

	return p
}

func (s *RuleidContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleidContext) ID() antlr.TerminalNode {
	return s.GetToken(bnfParserID, 0)
}

func (s *RuleidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.EnterRuleid(s)
	}
}

func (s *RuleidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bnfListener); ok {
		listenerT.ExitRuleid(s)
	}
}

func (p *bnfParser) Ruleid() (localctx IRuleidContext) {
	localctx = NewRuleidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, bnfParserRULE_ruleid)

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
		p.Match(bnfParserID)
	}

	return localctx
}
