// Code generated from DiceNotationParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dicenotation // DiceNotationParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 58, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 5, 2, 18, 10, 2, 3, 3, 3, 3, 3, 3, 7, 3, 23, 10, 3, 12,
	3, 14, 3, 26, 11, 3, 3, 4, 3, 4, 3, 4, 7, 4, 31, 10, 4, 12, 4, 14, 4, 34,
	11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 42, 10, 5, 3, 6, 5, 6,
	45, 10, 6, 3, 6, 5, 6, 48, 10, 6, 3, 6, 3, 6, 3, 6, 3, 7, 5, 7, 54, 10,
	7, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 2, 2, 60, 2, 17, 3,
	2, 2, 2, 4, 19, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 41, 3, 2, 2, 2, 10, 44,
	3, 2, 2, 2, 12, 53, 3, 2, 2, 2, 14, 18, 5, 10, 6, 2, 15, 18, 5, 12, 7,
	2, 16, 18, 5, 4, 3, 2, 17, 14, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 16,
	3, 2, 2, 2, 18, 3, 3, 2, 2, 2, 19, 24, 5, 6, 4, 2, 20, 21, 7, 5, 2, 2,
	21, 23, 5, 6, 4, 2, 22, 20, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 22, 3,
	2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 5, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 27,
	32, 5, 8, 5, 2, 28, 29, 7, 6, 2, 2, 29, 31, 5, 8, 5, 2, 30, 28, 3, 2, 2,
	2, 31, 34, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 7, 3,
	2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 42, 5, 10, 6, 2, 36, 42, 5, 12, 7, 2,
	37, 38, 7, 7, 2, 2, 38, 39, 5, 2, 2, 2, 39, 40, 7, 8, 2, 2, 40, 42, 3,
	2, 2, 2, 41, 35, 3, 2, 2, 2, 41, 36, 3, 2, 2, 2, 41, 37, 3, 2, 2, 2, 42,
	9, 3, 2, 2, 2, 43, 45, 7, 5, 2, 2, 44, 43, 3, 2, 2, 2, 44, 45, 3, 2, 2,
	2, 45, 47, 3, 2, 2, 2, 46, 48, 7, 4, 2, 2, 47, 46, 3, 2, 2, 2, 47, 48,
	3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50, 7, 3, 2, 2, 50, 51, 7, 4, 2, 2,
	51, 11, 3, 2, 2, 2, 52, 54, 7, 5, 2, 2, 53, 52, 3, 2, 2, 2, 53, 54, 3,
	2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 7, 4, 2, 2, 56, 13, 3, 2, 2, 2, 9,
	17, 24, 32, 41, 44, 47, 53,
}
var literalNames = []string{
	"", "", "", "", "", "'('", "')'",
}
var symbolicNames = []string{
	"", "DSEPARATOR", "DIGIT", "ADDOPERATOR", "MULTOPERATOR", "LPAREN", "RPAREN",
	"WS",
}

var ruleNames = []string{
	"notation", "addOp", "multOp", "operand", "dice", "number",
}

type DiceNotationParser struct {
	*antlr.BaseParser
}

// NewDiceNotationParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DiceNotationParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDiceNotationParser(input antlr.TokenStream) *DiceNotationParser {
	this := new(DiceNotationParser)
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
	this.GrammarFileName = "DiceNotationParser.g4"

	return this
}

// DiceNotationParser tokens.
const (
	DiceNotationParserEOF          = antlr.TokenEOF
	DiceNotationParserDSEPARATOR   = 1
	DiceNotationParserDIGIT        = 2
	DiceNotationParserADDOPERATOR  = 3
	DiceNotationParserMULTOPERATOR = 4
	DiceNotationParserLPAREN       = 5
	DiceNotationParserRPAREN       = 6
	DiceNotationParserWS           = 7
)

// DiceNotationParser rules.
const (
	DiceNotationParserRULE_notation = 0
	DiceNotationParserRULE_addOp    = 1
	DiceNotationParserRULE_multOp   = 2
	DiceNotationParserRULE_operand  = 3
	DiceNotationParserRULE_dice     = 4
	DiceNotationParserRULE_number   = 5
)

// INotationContext is an interface to support dynamic dispatch.
type INotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotationContext differentiates from other interfaces.
	IsNotationContext()
}

type NotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotationContext() *NotationContext {
	var p = new(NotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceNotationParserRULE_notation
	return p
}

func (*NotationContext) IsNotationContext() {}

func NewNotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotationContext {
	var p = new(NotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceNotationParserRULE_notation

	return p
}

func (s *NotationContext) GetParser() antlr.Parser { return s.parser }

func (s *NotationContext) Dice() IDiceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDiceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDiceContext)
}

func (s *NotationContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *NotationContext) AddOp() IAddOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddOpContext)
}

func (s *NotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.EnterNotation(s)
	}
}

func (s *NotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.ExitNotation(s)
	}
}

func (p *DiceNotationParser) Notation() (localctx INotationContext) {
	this := p
	_ = this

	localctx = NewNotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiceNotationParserRULE_notation)

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

	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(12)
			p.Dice()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.Number()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(14)
			p.AddOp()
		}

	}

	return localctx
}

// IAddOpContext is an interface to support dynamic dispatch.
type IAddOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddOpContext differentiates from other interfaces.
	IsAddOpContext()
}

type AddOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddOpContext() *AddOpContext {
	var p = new(AddOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceNotationParserRULE_addOp
	return p
}

func (*AddOpContext) IsAddOpContext() {}

func NewAddOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOpContext {
	var p = new(AddOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceNotationParserRULE_addOp

	return p
}

func (s *AddOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOpContext) AllMultOp() []IMultOpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultOpContext)(nil)).Elem())
	var tst = make([]IMultOpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultOpContext)
		}
	}

	return tst
}

func (s *AddOpContext) MultOp(i int) IMultOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultOpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultOpContext)
}

func (s *AddOpContext) AllADDOPERATOR() []antlr.TerminalNode {
	return s.GetTokens(DiceNotationParserADDOPERATOR)
}

func (s *AddOpContext) ADDOPERATOR(i int) antlr.TerminalNode {
	return s.GetToken(DiceNotationParserADDOPERATOR, i)
}

func (s *AddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.EnterAddOp(s)
	}
}

func (s *AddOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.ExitAddOp(s)
	}
}

func (p *DiceNotationParser) AddOp() (localctx IAddOpContext) {
	this := p
	_ = this

	localctx = NewAddOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiceNotationParserRULE_addOp)
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
		p.SetState(17)
		p.MultOp()
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DiceNotationParserADDOPERATOR {
		{
			p.SetState(18)
			p.Match(DiceNotationParserADDOPERATOR)
		}
		{
			p.SetState(19)
			p.MultOp()
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultOpContext is an interface to support dynamic dispatch.
type IMultOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultOpContext differentiates from other interfaces.
	IsMultOpContext()
}

type MultOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultOpContext() *MultOpContext {
	var p = new(MultOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceNotationParserRULE_multOp
	return p
}

func (*MultOpContext) IsMultOpContext() {}

func NewMultOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultOpContext {
	var p = new(MultOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceNotationParserRULE_multOp

	return p
}

func (s *MultOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultOpContext) AllOperand() []IOperandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperandContext)(nil)).Elem())
	var tst = make([]IOperandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperandContext)
		}
	}

	return tst
}

func (s *MultOpContext) Operand(i int) IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *MultOpContext) AllMULTOPERATOR() []antlr.TerminalNode {
	return s.GetTokens(DiceNotationParserMULTOPERATOR)
}

func (s *MultOpContext) MULTOPERATOR(i int) antlr.TerminalNode {
	return s.GetToken(DiceNotationParserMULTOPERATOR, i)
}

func (s *MultOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.EnterMultOp(s)
	}
}

func (s *MultOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.ExitMultOp(s)
	}
}

func (p *DiceNotationParser) MultOp() (localctx IMultOpContext) {
	this := p
	_ = this

	localctx = NewMultOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DiceNotationParserRULE_multOp)
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
		p.SetState(25)
		p.Operand()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DiceNotationParserMULTOPERATOR {
		{
			p.SetState(26)
			p.Match(DiceNotationParserMULTOPERATOR)
		}
		{
			p.SetState(27)
			p.Operand()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceNotationParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceNotationParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Dice() IDiceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDiceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDiceContext)
}

func (s *OperandContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *OperandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DiceNotationParserLPAREN, 0)
}

func (s *OperandContext) Notation() INotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotationContext)
}

func (s *OperandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DiceNotationParserRPAREN, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (p *DiceNotationParser) Operand() (localctx IOperandContext) {
	this := p
	_ = this

	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DiceNotationParserRULE_operand)

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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Dice()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Number()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.Match(DiceNotationParserLPAREN)
		}
		{
			p.SetState(36)
			p.Notation()
		}
		{
			p.SetState(37)
			p.Match(DiceNotationParserRPAREN)
		}

	}

	return localctx
}

// IDiceContext is an interface to support dynamic dispatch.
type IDiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiceContext differentiates from other interfaces.
	IsDiceContext()
}

type DiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiceContext() *DiceContext {
	var p = new(DiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiceNotationParserRULE_dice
	return p
}

func (*DiceContext) IsDiceContext() {}

func NewDiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiceContext {
	var p = new(DiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceNotationParserRULE_dice

	return p
}

func (s *DiceContext) GetParser() antlr.Parser { return s.parser }

func (s *DiceContext) DSEPARATOR() antlr.TerminalNode {
	return s.GetToken(DiceNotationParserDSEPARATOR, 0)
}

func (s *DiceContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(DiceNotationParserDIGIT)
}

func (s *DiceContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(DiceNotationParserDIGIT, i)
}

func (s *DiceContext) ADDOPERATOR() antlr.TerminalNode {
	return s.GetToken(DiceNotationParserADDOPERATOR, 0)
}

func (s *DiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.EnterDice(s)
	}
}

func (s *DiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.ExitDice(s)
	}
}

func (p *DiceNotationParser) Dice() (localctx IDiceContext) {
	this := p
	_ = this

	localctx = NewDiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DiceNotationParserRULE_dice)
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
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DiceNotationParserADDOPERATOR {
		{
			p.SetState(41)
			p.Match(DiceNotationParserADDOPERATOR)
		}

	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DiceNotationParserDIGIT {
		{
			p.SetState(44)
			p.Match(DiceNotationParserDIGIT)
		}

	}
	{
		p.SetState(47)
		p.Match(DiceNotationParserDSEPARATOR)
	}
	{
		p.SetState(48)
		p.Match(DiceNotationParserDIGIT)
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
	p.RuleIndex = DiceNotationParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiceNotationParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(DiceNotationParserDIGIT, 0)
}

func (s *NumberContext) ADDOPERATOR() antlr.TerminalNode {
	return s.GetToken(DiceNotationParserADDOPERATOR, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiceNotationParserListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *DiceNotationParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DiceNotationParserRULE_number)
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
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DiceNotationParserADDOPERATOR {
		{
			p.SetState(50)
			p.Match(DiceNotationParserADDOPERATOR)
		}

	}
	{
		p.SetState(53)
		p.Match(DiceNotationParserDIGIT)
	}

	return localctx
}
