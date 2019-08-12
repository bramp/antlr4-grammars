// Code generated from istc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package istc // istc
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 44, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2, 38,
	2, 12, 3, 2, 2, 2, 4, 23, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 32, 3, 2, 2,
	2, 10, 41, 3, 2, 2, 2, 12, 13, 7, 3, 2, 2, 13, 14, 7, 4, 2, 2, 14, 15,
	5, 4, 3, 2, 15, 16, 7, 5, 2, 2, 16, 17, 5, 6, 4, 2, 17, 18, 7, 5, 2, 2,
	18, 19, 5, 8, 5, 2, 19, 20, 7, 5, 2, 2, 20, 21, 5, 10, 6, 2, 21, 22, 7,
	2, 2, 3, 22, 3, 3, 2, 2, 2, 23, 24, 7, 6, 2, 2, 24, 25, 7, 6, 2, 2, 25,
	26, 7, 6, 2, 2, 26, 5, 3, 2, 2, 2, 27, 28, 7, 6, 2, 2, 28, 29, 7, 6, 2,
	2, 29, 30, 7, 6, 2, 2, 30, 31, 7, 6, 2, 2, 31, 7, 3, 2, 2, 2, 32, 33, 7,
	6, 2, 2, 33, 34, 7, 6, 2, 2, 34, 35, 7, 6, 2, 2, 35, 36, 7, 6, 2, 2, 36,
	37, 7, 6, 2, 2, 37, 38, 7, 6, 2, 2, 38, 39, 7, 6, 2, 2, 39, 40, 7, 6, 2,
	2, 40, 9, 3, 2, 2, 2, 41, 42, 7, 6, 2, 2, 42, 11, 3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'ISTC'", "' '",
}
var symbolicNames = []string{
	"", "", "", "SEP", "CHAR", "WS",
}

var ruleNames = []string{
	"istc", "registration", "year", "work", "check",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type istcParser struct {
	*antlr.BaseParser
}

func NewistcParser(input antlr.TokenStream) *istcParser {
	this := new(istcParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "istc.g4"

	return this
}

// istcParser tokens.
const (
	istcParserEOF  = antlr.TokenEOF
	istcParserT__0 = 1
	istcParserT__1 = 2
	istcParserSEP  = 3
	istcParserCHAR = 4
	istcParserWS   = 5
)

// istcParser rules.
const (
	istcParserRULE_istc         = 0
	istcParserRULE_registration = 1
	istcParserRULE_year         = 2
	istcParserRULE_work         = 3
	istcParserRULE_check        = 4
)

// IIstcContext is an interface to support dynamic dispatch.
type IIstcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIstcContext differentiates from other interfaces.
	IsIstcContext()
}

type IstcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIstcContext() *IstcContext {
	var p = new(IstcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = istcParserRULE_istc
	return p
}

func (*IstcContext) IsIstcContext() {}

func NewIstcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IstcContext {
	var p = new(IstcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = istcParserRULE_istc

	return p
}

func (s *IstcContext) GetParser() antlr.Parser { return s.parser }

func (s *IstcContext) Registration() IRegistrationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegistrationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegistrationContext)
}

func (s *IstcContext) AllSEP() []antlr.TerminalNode {
	return s.GetTokens(istcParserSEP)
}

func (s *IstcContext) SEP(i int) antlr.TerminalNode {
	return s.GetToken(istcParserSEP, i)
}

func (s *IstcContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *IstcContext) Work() IWorkContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWorkContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWorkContext)
}

func (s *IstcContext) Check() ICheckContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheckContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheckContext)
}

func (s *IstcContext) EOF() antlr.TerminalNode {
	return s.GetToken(istcParserEOF, 0)
}

func (s *IstcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IstcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IstcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(istcListener); ok {
		listenerT.EnterIstc(s)
	}
}

func (s *IstcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(istcListener); ok {
		listenerT.ExitIstc(s)
	}
}

func (p *istcParser) Istc() (localctx IIstcContext) {
	localctx = NewIstcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, istcParserRULE_istc)

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
		p.SetState(10)
		p.Match(istcParserT__0)
	}
	{
		p.SetState(11)
		p.Match(istcParserT__1)
	}
	{
		p.SetState(12)
		p.Registration()
	}
	{
		p.SetState(13)
		p.Match(istcParserSEP)
	}
	{
		p.SetState(14)
		p.Year()
	}
	{
		p.SetState(15)
		p.Match(istcParserSEP)
	}
	{
		p.SetState(16)
		p.Work()
	}
	{
		p.SetState(17)
		p.Match(istcParserSEP)
	}
	{
		p.SetState(18)
		p.Check()
	}
	{
		p.SetState(19)
		p.Match(istcParserEOF)
	}

	return localctx
}

// IRegistrationContext is an interface to support dynamic dispatch.
type IRegistrationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegistrationContext differentiates from other interfaces.
	IsRegistrationContext()
}

type RegistrationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegistrationContext() *RegistrationContext {
	var p = new(RegistrationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = istcParserRULE_registration
	return p
}

func (*RegistrationContext) IsRegistrationContext() {}

func NewRegistrationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegistrationContext {
	var p = new(RegistrationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = istcParserRULE_registration

	return p
}

func (s *RegistrationContext) GetParser() antlr.Parser { return s.parser }

func (s *RegistrationContext) AllCHAR() []antlr.TerminalNode {
	return s.GetTokens(istcParserCHAR)
}

func (s *RegistrationContext) CHAR(i int) antlr.TerminalNode {
	return s.GetToken(istcParserCHAR, i)
}

func (s *RegistrationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegistrationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegistrationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(istcListener); ok {
		listenerT.EnterRegistration(s)
	}
}

func (s *RegistrationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(istcListener); ok {
		listenerT.ExitRegistration(s)
	}
}

func (p *istcParser) Registration() (localctx IRegistrationContext) {
	localctx = NewRegistrationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, istcParserRULE_registration)

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
		p.SetState(21)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(22)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(23)
		p.Match(istcParserCHAR)
	}

	return localctx
}

// IYearContext is an interface to support dynamic dispatch.
type IYearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYearContext differentiates from other interfaces.
	IsYearContext()
}

type YearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYearContext() *YearContext {
	var p = new(YearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = istcParserRULE_year
	return p
}

func (*YearContext) IsYearContext() {}

func NewYearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YearContext {
	var p = new(YearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = istcParserRULE_year

	return p
}

func (s *YearContext) GetParser() antlr.Parser { return s.parser }

func (s *YearContext) AllCHAR() []antlr.TerminalNode {
	return s.GetTokens(istcParserCHAR)
}

func (s *YearContext) CHAR(i int) antlr.TerminalNode {
	return s.GetToken(istcParserCHAR, i)
}

func (s *YearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(istcListener); ok {
		listenerT.EnterYear(s)
	}
}

func (s *YearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(istcListener); ok {
		listenerT.ExitYear(s)
	}
}

func (p *istcParser) Year() (localctx IYearContext) {
	localctx = NewYearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, istcParserRULE_year)

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
		p.SetState(25)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(26)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(27)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(28)
		p.Match(istcParserCHAR)
	}

	return localctx
}

// IWorkContext is an interface to support dynamic dispatch.
type IWorkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWorkContext differentiates from other interfaces.
	IsWorkContext()
}

type WorkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWorkContext() *WorkContext {
	var p = new(WorkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = istcParserRULE_work
	return p
}

func (*WorkContext) IsWorkContext() {}

func NewWorkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WorkContext {
	var p = new(WorkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = istcParserRULE_work

	return p
}

func (s *WorkContext) GetParser() antlr.Parser { return s.parser }

func (s *WorkContext) AllCHAR() []antlr.TerminalNode {
	return s.GetTokens(istcParserCHAR)
}

func (s *WorkContext) CHAR(i int) antlr.TerminalNode {
	return s.GetToken(istcParserCHAR, i)
}

func (s *WorkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WorkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(istcListener); ok {
		listenerT.EnterWork(s)
	}
}

func (s *WorkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(istcListener); ok {
		listenerT.ExitWork(s)
	}
}

func (p *istcParser) Work() (localctx IWorkContext) {
	localctx = NewWorkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, istcParserRULE_work)

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
		p.SetState(30)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(31)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(32)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(33)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(34)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(35)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(36)
		p.Match(istcParserCHAR)
	}
	{
		p.SetState(37)
		p.Match(istcParserCHAR)
	}

	return localctx
}

// ICheckContext is an interface to support dynamic dispatch.
type ICheckContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCheckContext differentiates from other interfaces.
	IsCheckContext()
}

type CheckContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCheckContext() *CheckContext {
	var p = new(CheckContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = istcParserRULE_check
	return p
}

func (*CheckContext) IsCheckContext() {}

func NewCheckContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CheckContext {
	var p = new(CheckContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = istcParserRULE_check

	return p
}

func (s *CheckContext) GetParser() antlr.Parser { return s.parser }

func (s *CheckContext) CHAR() antlr.TerminalNode {
	return s.GetToken(istcParserCHAR, 0)
}

func (s *CheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(istcListener); ok {
		listenerT.EnterCheck(s)
	}
}

func (s *CheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(istcListener); ok {
		listenerT.ExitCheck(s)
	}
}

func (p *istcParser) Check() (localctx ICheckContext) {
	localctx = NewCheckContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, istcParserRULE_check)

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
		p.SetState(39)
		p.Match(istcParserCHAR)
	}

	return localctx
}
