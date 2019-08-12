// Code generated from MuParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package mu // MuParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 85, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 7, 2, 12, 10, 2, 12,
	2, 14, 2, 15, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 30, 10, 3, 12, 3, 14, 3, 33, 11, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 41, 10, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7,
	3, 70, 10, 3, 12, 3, 14, 3, 73, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 83, 10, 4, 3, 4, 2, 3, 4, 5, 2, 4, 6, 2, 10, 3, 2,
	7, 11, 3, 2, 22, 23, 3, 2, 20, 21, 4, 2, 14, 15, 18, 19, 3, 2, 16, 17,
	3, 2, 30, 31, 3, 2, 32, 33, 3, 2, 34, 35, 2, 99, 2, 8, 3, 2, 2, 2, 4, 40,
	3, 2, 2, 2, 6, 82, 3, 2, 2, 2, 8, 13, 5, 4, 3, 2, 9, 10, 7, 3, 2, 2, 10,
	12, 5, 4, 3, 2, 11, 9, 3, 2, 2, 2, 12, 15, 3, 2, 2, 2, 13, 11, 3, 2, 2,
	2, 13, 14, 3, 2, 2, 2, 14, 3, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 16, 17, 8,
	3, 1, 2, 17, 18, 7, 21, 2, 2, 18, 41, 5, 4, 3, 14, 19, 20, 7, 5, 2, 2,
	20, 21, 7, 28, 2, 2, 21, 22, 5, 4, 3, 2, 22, 23, 7, 29, 2, 2, 23, 41, 3,
	2, 2, 2, 24, 25, 7, 6, 2, 2, 25, 26, 7, 28, 2, 2, 26, 31, 5, 4, 3, 2, 27,
	28, 7, 4, 2, 2, 28, 30, 5, 4, 3, 2, 29, 27, 3, 2, 2, 2, 30, 33, 3, 2, 2,
	2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2, 2, 33, 31,
	3, 2, 2, 2, 34, 35, 7, 29, 2, 2, 35, 41, 3, 2, 2, 2, 36, 41, 5, 6, 4, 2,
	37, 38, 7, 36, 2, 2, 38, 39, 9, 2, 2, 2, 39, 41, 5, 4, 3, 3, 40, 16, 3,
	2, 2, 2, 40, 19, 3, 2, 2, 2, 40, 24, 3, 2, 2, 2, 40, 36, 3, 2, 2, 2, 40,
	37, 3, 2, 2, 2, 41, 71, 3, 2, 2, 2, 42, 43, 12, 15, 2, 2, 43, 44, 7, 24,
	2, 2, 44, 70, 5, 4, 3, 15, 45, 46, 12, 13, 2, 2, 46, 47, 9, 3, 2, 2, 47,
	70, 5, 4, 3, 14, 48, 49, 12, 12, 2, 2, 49, 50, 9, 4, 2, 2, 50, 70, 5, 4,
	3, 13, 51, 52, 12, 11, 2, 2, 52, 53, 9, 5, 2, 2, 53, 70, 5, 4, 3, 12, 54,
	55, 12, 10, 2, 2, 55, 56, 9, 6, 2, 2, 56, 70, 5, 4, 3, 11, 57, 58, 12,
	9, 2, 2, 58, 59, 7, 12, 2, 2, 59, 70, 5, 4, 3, 10, 60, 61, 12, 8, 2, 2,
	61, 62, 7, 13, 2, 2, 62, 70, 5, 4, 3, 9, 63, 64, 12, 7, 2, 2, 64, 65, 7,
	26, 2, 2, 65, 66, 5, 4, 3, 2, 66, 67, 7, 27, 2, 2, 67, 68, 5, 4, 3, 8,
	68, 70, 3, 2, 2, 2, 69, 42, 3, 2, 2, 2, 69, 45, 3, 2, 2, 2, 69, 48, 3,
	2, 2, 2, 69, 51, 3, 2, 2, 2, 69, 54, 3, 2, 2, 2, 69, 57, 3, 2, 2, 2, 69,
	60, 3, 2, 2, 2, 69, 63, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2,
	2, 71, 72, 3, 2, 2, 2, 72, 5, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 75, 7,
	28, 2, 2, 75, 76, 5, 4, 3, 2, 76, 77, 7, 29, 2, 2, 77, 83, 3, 2, 2, 2,
	78, 83, 9, 7, 2, 2, 79, 83, 9, 8, 2, 2, 80, 83, 9, 9, 2, 2, 81, 83, 7,
	36, 2, 2, 82, 74, 3, 2, 2, 2, 82, 78, 3, 2, 2, 2, 82, 79, 3, 2, 2, 2, 82,
	80, 3, 2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 7, 3, 2, 2, 2, 8, 13, 31, 40, 69,
	71, 82,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\n'", "','", "", "", "'='", "'+='", "'-='", "'*='", "'/='", "'&&'",
	"'||'", "'<='", "'>='", "'!='", "'=='", "'<'", "'>'", "'+'", "'-'", "'*'",
	"'/'", "'^'", "'!'", "'?'", "':'", "'('", "')'", "", "", "'true'", "'false'",
	"'_e'", "'_pi'",
}
var symbolicNames = []string{
	"", "", "", "FUNCTION", "FUNCTIONMULTI", "ASSIGN", "ASSIGNADD", "ASSIGNSUB",
	"ASSIGNMUL", "ASSIGNDIV", "AND", "OR", "LTEQ", "GTEQ", "NEQ", "EQ", "LT",
	"GT", "ADD", "SUB", "MUL", "DIV", "POW", "NOT", "QUESTION", "COLON", "OPAR",
	"CPAR", "INT", "FLOAT", "TRUE", "FALSE", "E", "PI", "ID", "SPACE",
}

var ruleNames = []string{
	"prog", "expr", "atom",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MuParserParser struct {
	*antlr.BaseParser
}

func NewMuParserParser(input antlr.TokenStream) *MuParserParser {
	this := new(MuParserParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "MuParser.g4"

	return this
}

// MuParserParser tokens.
const (
	MuParserParserEOF           = antlr.TokenEOF
	MuParserParserT__0          = 1
	MuParserParserT__1          = 2
	MuParserParserFUNCTION      = 3
	MuParserParserFUNCTIONMULTI = 4
	MuParserParserASSIGN        = 5
	MuParserParserASSIGNADD     = 6
	MuParserParserASSIGNSUB     = 7
	MuParserParserASSIGNMUL     = 8
	MuParserParserASSIGNDIV     = 9
	MuParserParserAND           = 10
	MuParserParserOR            = 11
	MuParserParserLTEQ          = 12
	MuParserParserGTEQ          = 13
	MuParserParserNEQ           = 14
	MuParserParserEQ            = 15
	MuParserParserLT            = 16
	MuParserParserGT            = 17
	MuParserParserADD           = 18
	MuParserParserSUB           = 19
	MuParserParserMUL           = 20
	MuParserParserDIV           = 21
	MuParserParserPOW           = 22
	MuParserParserNOT           = 23
	MuParserParserQUESTION      = 24
	MuParserParserCOLON         = 25
	MuParserParserOPAR          = 26
	MuParserParserCPAR          = 27
	MuParserParserINT           = 28
	MuParserParserFLOAT         = 29
	MuParserParserTRUE          = 30
	MuParserParserFALSE         = 31
	MuParserParserE             = 32
	MuParserParserPI            = 33
	MuParserParserID            = 34
	MuParserParserSPACE         = 35
)

// MuParserParser rules.
const (
	MuParserParserRULE_prog = 0
	MuParserParserRULE_expr = 1
	MuParserParserRULE_atom = 2
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MuParserParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MuParserParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) CopyFrom(ctx *ProgContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ProgExprContext struct {
	*ProgContext
}

func NewProgExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProgExprContext {
	var p = new(ProgExprContext)

	p.ProgContext = NewEmptyProgContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ProgContext))

	return p
}

func (s *ProgExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ProgExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterProgExpr(s)
	}
}

func (s *ProgExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitProgExpr(s)
	}
}

func (p *MuParserParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MuParserParserRULE_prog)
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

	localctx = NewProgExprContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.expr(0)
	}
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MuParserParserT__0 {
		{
			p.SetState(7)
			p.Match(MuParserParserT__0)
		}
		{
			p.SetState(8)
			p.expr(0)
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MuParserParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MuParserParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionMultiExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewFunctionMultiExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionMultiExprContext {
	var p = new(FunctionMultiExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FunctionMultiExprContext) GetOp() antlr.Token { return s.op }

func (s *FunctionMultiExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *FunctionMultiExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionMultiExprContext) OPAR() antlr.TerminalNode {
	return s.GetToken(MuParserParserOPAR, 0)
}

func (s *FunctionMultiExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *FunctionMultiExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionMultiExprContext) CPAR() antlr.TerminalNode {
	return s.GetToken(MuParserParserCPAR, 0)
}

func (s *FunctionMultiExprContext) FUNCTIONMULTI() antlr.TerminalNode {
	return s.GetToken(MuParserParserFUNCTIONMULTI, 0)
}

func (s *FunctionMultiExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterFunctionMultiExpr(s)
	}
}

func (s *FunctionMultiExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitFunctionMultiExpr(s)
	}
}

type AddSubExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubExprContext) GetOp() antlr.Token { return s.op }

func (s *AddSubExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(MuParserParserADD, 0)
}

func (s *AddSubExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(MuParserParserSUB, 0)
}

func (s *AddSubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterAddSubExpr(s)
	}
}

func (s *AddSubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitAddSubExpr(s)
	}
}

type AtomExprContext struct {
	*ExprContext
}

func NewAtomExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomExprContext {
	var p = new(AtomExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AtomExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExprContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterAtomExpr(s)
	}
}

func (s *AtomExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitAtomExpr(s)
	}
}

type OrExprContext struct {
	*ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(MuParserParserOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

type RelationalExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewRelationalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExprContext {
	var p = new(RelationalExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RelationalExprContext) GetOp() antlr.Token { return s.op }

func (s *RelationalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RelationalExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RelationalExprContext) LTEQ() antlr.TerminalNode {
	return s.GetToken(MuParserParserLTEQ, 0)
}

func (s *RelationalExprContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(MuParserParserGTEQ, 0)
}

func (s *RelationalExprContext) LT() antlr.TerminalNode {
	return s.GetToken(MuParserParserLT, 0)
}

func (s *RelationalExprContext) GT() antlr.TerminalNode {
	return s.GetToken(MuParserParserGT, 0)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

type FunctionExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewFunctionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExprContext {
	var p = new(FunctionExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FunctionExprContext) GetOp() antlr.Token { return s.op }

func (s *FunctionExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *FunctionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExprContext) OPAR() antlr.TerminalNode {
	return s.GetToken(MuParserParserOPAR, 0)
}

func (s *FunctionExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionExprContext) CPAR() antlr.TerminalNode {
	return s.GetToken(MuParserParserCPAR, 0)
}

func (s *FunctionExprContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(MuParserParserFUNCTION, 0)
}

func (s *FunctionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterFunctionExpr(s)
	}
}

func (s *FunctionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitFunctionExpr(s)
	}
}

type UnaryMinusExprContext struct {
	*ExprContext
}

func NewUnaryMinusExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExprContext {
	var p = new(UnaryMinusExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnaryMinusExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(MuParserParserSUB, 0)
}

func (s *UnaryMinusExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryMinusExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterUnaryMinusExpr(s)
	}
}

func (s *UnaryMinusExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitUnaryMinusExpr(s)
	}
}

type PowExprContext struct {
	*ExprContext
}

func NewPowExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowExprContext {
	var p = new(PowExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PowExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PowExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PowExprContext) POW() antlr.TerminalNode {
	return s.GetToken(MuParserParserPOW, 0)
}

func (s *PowExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterPowExpr(s)
	}
}

func (s *PowExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitPowExpr(s)
	}
}

type AssignExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExprContext {
	var p = new(AssignExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AssignExprContext) GetOp() antlr.Token { return s.op }

func (s *AssignExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) ID() antlr.TerminalNode {
	return s.GetToken(MuParserParserID, 0)
}

func (s *AssignExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignExprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MuParserParserASSIGN, 0)
}

func (s *AssignExprContext) ASSIGNADD() antlr.TerminalNode {
	return s.GetToken(MuParserParserASSIGNADD, 0)
}

func (s *AssignExprContext) ASSIGNSUB() antlr.TerminalNode {
	return s.GetToken(MuParserParserASSIGNSUB, 0)
}

func (s *AssignExprContext) ASSIGNMUL() antlr.TerminalNode {
	return s.GetToken(MuParserParserASSIGNMUL, 0)
}

func (s *AssignExprContext) ASSIGNDIV() antlr.TerminalNode {
	return s.GetToken(MuParserParserASSIGNDIV, 0)
}

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

type MulDivExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivExprContext {
	var p = new(MulDivExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivExprContext) GetOp() antlr.Token { return s.op }

func (s *MulDivExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulDivExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(MuParserParserMUL, 0)
}

func (s *MulDivExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(MuParserParserDIV, 0)
}

func (s *MulDivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterMulDivExpr(s)
	}
}

func (s *MulDivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitMulDivExpr(s)
	}
}

type EqualityExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewEqualityExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExprContext {
	var p = new(EqualityExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualityExprContext) GetOp() antlr.Token { return s.op }

func (s *EqualityExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualityExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualityExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(MuParserParserEQ, 0)
}

func (s *EqualityExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(MuParserParserNEQ, 0)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

type AndExprContext struct {
	*ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(MuParserParserAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

type IteExprContext struct {
	*ExprContext
}

func NewIteExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IteExprContext {
	var p = new(IteExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IteExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IteExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IteExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IteExprContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(MuParserParserQUESTION, 0)
}

func (s *IteExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(MuParserParserCOLON, 0)
}

func (s *IteExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterIteExpr(s)
	}
}

func (s *IteExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitIteExpr(s)
	}
}

func (p *MuParserParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *MuParserParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, MuParserParserRULE_expr, _p)
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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnaryMinusExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(15)
			p.Match(MuParserParserSUB)
		}
		{
			p.SetState(16)
			p.expr(12)
		}

	case 2:
		localctx = NewFunctionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)

			var _m = p.Match(MuParserParserFUNCTION)

			localctx.(*FunctionExprContext).op = _m
		}
		{
			p.SetState(18)
			p.Match(MuParserParserOPAR)
		}
		{
			p.SetState(19)
			p.expr(0)
		}
		{
			p.SetState(20)
			p.Match(MuParserParserCPAR)
		}

	case 3:
		localctx = NewFunctionMultiExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)

			var _m = p.Match(MuParserParserFUNCTIONMULTI)

			localctx.(*FunctionMultiExprContext).op = _m
		}
		{
			p.SetState(23)
			p.Match(MuParserParserOPAR)
		}
		{
			p.SetState(24)
			p.expr(0)
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MuParserParserT__1 {
			{
				p.SetState(25)
				p.Match(MuParserParserT__1)
			}
			{
				p.SetState(26)
				p.expr(0)
			}

			p.SetState(31)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(32)
			p.Match(MuParserParserCPAR)
		}

	case 4:
		localctx = NewAtomExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			p.Atom()
		}

	case 5:
		localctx = NewAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(35)
			p.Match(MuParserParserID)
		}
		{
			p.SetState(36)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AssignExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MuParserParserASSIGN)|(1<<MuParserParserASSIGNADD)|(1<<MuParserParserASSIGNSUB)|(1<<MuParserParserASSIGNMUL)|(1<<MuParserParserASSIGNDIV))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AssignExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(37)
			p.expr(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MuParserParserRULE_expr)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(41)
					p.Match(MuParserParserPOW)
				}
				{
					p.SetState(42)
					p.expr(13)
				}

			case 2:
				localctx = NewMulDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MuParserParserRULE_expr)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(44)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MuParserParserMUL || _la == MuParserParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(45)
					p.expr(12)
				}

			case 3:
				localctx = NewAddSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MuParserParserRULE_expr)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(47)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MuParserParserADD || _la == MuParserParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(48)
					p.expr(11)
				}

			case 4:
				localctx = NewRelationalExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MuParserParserRULE_expr)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(50)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MuParserParserLTEQ)|(1<<MuParserParserGTEQ)|(1<<MuParserParserLT)|(1<<MuParserParserGT))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(51)
					p.expr(10)
				}

			case 5:
				localctx = NewEqualityExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MuParserParserRULE_expr)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(53)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MuParserParserNEQ || _la == MuParserParserEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(54)
					p.expr(9)
				}

			case 6:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MuParserParserRULE_expr)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(56)
					p.Match(MuParserParserAND)
				}
				{
					p.SetState(57)
					p.expr(8)
				}

			case 7:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MuParserParserRULE_expr)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(59)
					p.Match(MuParserParserOR)
				}
				{
					p.SetState(60)
					p.expr(7)
				}

			case 8:
				localctx = NewIteExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MuParserParserRULE_expr)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(62)
					p.Match(MuParserParserQUESTION)
				}
				{
					p.SetState(63)
					p.expr(0)
				}
				{
					p.SetState(64)
					p.Match(MuParserParserCOLON)
				}
				{
					p.SetState(65)
					p.expr(6)
				}

			}

		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
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
	p.RuleIndex = MuParserParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MuParserParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParExprContext struct {
	*AtomContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) OPAR() antlr.TerminalNode {
	return s.GetToken(MuParserParserOPAR, 0)
}

func (s *ParExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExprContext) CPAR() antlr.TerminalNode {
	return s.GetToken(MuParserParserCPAR, 0)
}

func (s *ParExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterParExpr(s)
	}
}

func (s *ParExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitParExpr(s)
	}
}

type BooleanAtomContext struct {
	*AtomContext
}

func NewBooleanAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanAtomContext {
	var p = new(BooleanAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *BooleanAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanAtomContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MuParserParserTRUE, 0)
}

func (s *BooleanAtomContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MuParserParserFALSE, 0)
}

func (s *BooleanAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterBooleanAtom(s)
	}
}

func (s *BooleanAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitBooleanAtom(s)
	}
}

type IdAtomContext struct {
	*AtomContext
}

func NewIdAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdAtomContext {
	var p = new(IdAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *IdAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdAtomContext) ID() antlr.TerminalNode {
	return s.GetToken(MuParserParserID, 0)
}

func (s *IdAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterIdAtom(s)
	}
}

func (s *IdAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitIdAtom(s)
	}
}

type PredefinedConstantAtomContext struct {
	*AtomContext
}

func NewPredefinedConstantAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PredefinedConstantAtomContext {
	var p = new(PredefinedConstantAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *PredefinedConstantAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredefinedConstantAtomContext) E() antlr.TerminalNode {
	return s.GetToken(MuParserParserE, 0)
}

func (s *PredefinedConstantAtomContext) PI() antlr.TerminalNode {
	return s.GetToken(MuParserParserPI, 0)
}

func (s *PredefinedConstantAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterPredefinedConstantAtom(s)
	}
}

func (s *PredefinedConstantAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitPredefinedConstantAtom(s)
	}
}

type NumberAtomContext struct {
	*AtomContext
}

func NewNumberAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberAtomContext {
	var p = new(NumberAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *NumberAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberAtomContext) INT() antlr.TerminalNode {
	return s.GetToken(MuParserParserINT, 0)
}

func (s *NumberAtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(MuParserParserFLOAT, 0)
}

func (s *NumberAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.EnterNumberAtom(s)
	}
}

func (s *NumberAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MuParserListener); ok {
		listenerT.ExitNumberAtom(s)
	}
}

func (p *MuParserParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MuParserParserRULE_atom)
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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MuParserParserOPAR:
		localctx = NewParExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Match(MuParserParserOPAR)
		}
		{
			p.SetState(73)
			p.expr(0)
		}
		{
			p.SetState(74)
			p.Match(MuParserParserCPAR)
		}

	case MuParserParserINT, MuParserParserFLOAT:
		localctx = NewNumberAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MuParserParserINT || _la == MuParserParserFLOAT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MuParserParserTRUE, MuParserParserFALSE:
		localctx = NewBooleanAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MuParserParserTRUE || _la == MuParserParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MuParserParserE, MuParserParserPI:
		localctx = NewPredefinedConstantAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(78)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MuParserParserE || _la == MuParserParserPI) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MuParserParserID:
		localctx = NewIdAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(79)
			p.Match(MuParserParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *MuParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MuParserParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
