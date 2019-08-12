// Code generated from sexpression.g4 by ANTLR 4.7.2. DO NOT EDIT.

package sexpression // sexpression
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 40, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 7, 2, 12, 10, 2, 12,
	2, 14, 2, 15, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 27, 10, 3, 3, 4, 3, 4, 7, 4, 31, 10, 4, 12, 4, 14, 4, 34,
	11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 3, 5, 2, 3,
	3, 5, 6, 9, 9, 2, 39, 2, 13, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 28, 3, 2,
	2, 2, 8, 37, 3, 2, 2, 2, 10, 12, 5, 4, 3, 2, 11, 10, 3, 2, 2, 2, 12, 15,
	3, 2, 2, 2, 13, 11, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14, 16, 3, 2, 2, 2,
	15, 13, 3, 2, 2, 2, 16, 17, 7, 2, 2, 3, 17, 3, 3, 2, 2, 2, 18, 27, 5, 8,
	5, 2, 19, 27, 5, 6, 4, 2, 20, 21, 7, 7, 2, 2, 21, 22, 5, 4, 3, 2, 22, 23,
	7, 9, 2, 2, 23, 24, 5, 4, 3, 2, 24, 25, 7, 8, 2, 2, 25, 27, 3, 2, 2, 2,
	26, 18, 3, 2, 2, 2, 26, 19, 3, 2, 2, 2, 26, 20, 3, 2, 2, 2, 27, 5, 3, 2,
	2, 2, 28, 32, 7, 7, 2, 2, 29, 31, 5, 4, 3, 2, 30, 29, 3, 2, 2, 2, 31, 34,
	3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 35, 3, 2, 2, 2,
	34, 32, 3, 2, 2, 2, 35, 36, 7, 8, 2, 2, 36, 7, 3, 2, 2, 2, 37, 38, 9, 2,
	2, 2, 38, 9, 3, 2, 2, 2, 5, 13, 26, 32,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'('", "')'", "'.'",
}
var symbolicNames = []string{
	"", "STRING", "WHITESPACE", "NUMBER", "SYMBOL", "LPAREN", "RPAREN", "DOT",
}

var ruleNames = []string{
	"sexpr", "item", "list", "atom",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type sexpressionParser struct {
	*antlr.BaseParser
}

func NewsexpressionParser(input antlr.TokenStream) *sexpressionParser {
	this := new(sexpressionParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "sexpression.g4"

	return this
}

// sexpressionParser tokens.
const (
	sexpressionParserEOF        = antlr.TokenEOF
	sexpressionParserSTRING     = 1
	sexpressionParserWHITESPACE = 2
	sexpressionParserNUMBER     = 3
	sexpressionParserSYMBOL     = 4
	sexpressionParserLPAREN     = 5
	sexpressionParserRPAREN     = 6
	sexpressionParserDOT        = 7
)

// sexpressionParser rules.
const (
	sexpressionParserRULE_sexpr = 0
	sexpressionParserRULE_item  = 1
	sexpressionParserRULE_list  = 2
	sexpressionParserRULE_atom  = 3
)

// ISexprContext is an interface to support dynamic dispatch.
type ISexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSexprContext differentiates from other interfaces.
	IsSexprContext()
}

type SexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySexprContext() *SexprContext {
	var p = new(SexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sexpressionParserRULE_sexpr
	return p
}

func (*SexprContext) IsSexprContext() {}

func NewSexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SexprContext {
	var p = new(SexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sexpressionParserRULE_sexpr

	return p
}

func (s *SexprContext) GetParser() antlr.Parser { return s.parser }

func (s *SexprContext) EOF() antlr.TerminalNode {
	return s.GetToken(sexpressionParserEOF, 0)
}

func (s *SexprContext) AllItem() []IItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IItemContext)(nil)).Elem())
	var tst = make([]IItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IItemContext)
		}
	}

	return tst
}

func (s *SexprContext) Item(i int) IItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *SexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sexpressionListener); ok {
		listenerT.EnterSexpr(s)
	}
}

func (s *SexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sexpressionListener); ok {
		listenerT.ExitSexpr(s)
	}
}

func (p *sexpressionParser) Sexpr() (localctx ISexprContext) {
	localctx = NewSexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, sexpressionParserRULE_sexpr)
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
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<sexpressionParserSTRING)|(1<<sexpressionParserNUMBER)|(1<<sexpressionParserSYMBOL)|(1<<sexpressionParserLPAREN)|(1<<sexpressionParserDOT))) != 0 {
		{
			p.SetState(8)
			p.Item()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(14)
		p.Match(sexpressionParserEOF)
	}

	return localctx
}

// IItemContext is an interface to support dynamic dispatch.
type IItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemContext differentiates from other interfaces.
	IsItemContext()
}

type ItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemContext() *ItemContext {
	var p = new(ItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sexpressionParserRULE_item
	return p
}

func (*ItemContext) IsItemContext() {}

func NewItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemContext {
	var p = new(ItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sexpressionParserRULE_item

	return p
}

func (s *ItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ItemContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ItemContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(sexpressionParserLPAREN, 0)
}

func (s *ItemContext) AllItem() []IItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IItemContext)(nil)).Elem())
	var tst = make([]IItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IItemContext)
		}
	}

	return tst
}

func (s *ItemContext) Item(i int) IItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *ItemContext) DOT() antlr.TerminalNode {
	return s.GetToken(sexpressionParserDOT, 0)
}

func (s *ItemContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(sexpressionParserRPAREN, 0)
}

func (s *ItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sexpressionListener); ok {
		listenerT.EnterItem(s)
	}
}

func (s *ItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sexpressionListener); ok {
		listenerT.ExitItem(s)
	}
}

func (p *sexpressionParser) Item() (localctx IItemContext) {
	localctx = NewItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, sexpressionParserRULE_item)

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

	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Atom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)
			p.List()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(18)
			p.Match(sexpressionParserLPAREN)
		}
		{
			p.SetState(19)
			p.Item()
		}
		{
			p.SetState(20)
			p.Match(sexpressionParserDOT)
		}
		{
			p.SetState(21)
			p.Item()
		}
		{
			p.SetState(22)
			p.Match(sexpressionParserRPAREN)
		}

	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sexpressionParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sexpressionParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(sexpressionParserLPAREN, 0)
}

func (s *ListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(sexpressionParserRPAREN, 0)
}

func (s *ListContext) AllItem() []IItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IItemContext)(nil)).Elem())
	var tst = make([]IItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IItemContext)
		}
	}

	return tst
}

func (s *ListContext) Item(i int) IItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sexpressionListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sexpressionListener); ok {
		listenerT.ExitList(s)
	}
}

func (p *sexpressionParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, sexpressionParserRULE_list)
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
		p.SetState(26)
		p.Match(sexpressionParserLPAREN)
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<sexpressionParserSTRING)|(1<<sexpressionParserNUMBER)|(1<<sexpressionParserSYMBOL)|(1<<sexpressionParserLPAREN)|(1<<sexpressionParserDOT))) != 0 {
		{
			p.SetState(27)
			p.Item()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Match(sexpressionParserRPAREN)
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
	p.RuleIndex = sexpressionParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sexpressionParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(sexpressionParserSTRING, 0)
}

func (s *AtomContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(sexpressionParserSYMBOL, 0)
}

func (s *AtomContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(sexpressionParserNUMBER, 0)
}

func (s *AtomContext) DOT() antlr.TerminalNode {
	return s.GetToken(sexpressionParserDOT, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sexpressionListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sexpressionListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *sexpressionParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, sexpressionParserRULE_atom)
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
		p.SetState(35)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<sexpressionParserSTRING)|(1<<sexpressionParserNUMBER)|(1<<sexpressionParserSYMBOL)|(1<<sexpressionParserDOT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
