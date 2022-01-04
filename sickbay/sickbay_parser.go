// Code generated from sickbay.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sickbay // sickbay
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 81, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 7, 2, 14,
	10, 2, 12, 2, 14, 2, 17, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 23, 10, 3,
	12, 3, 14, 3, 26, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 55, 10, 4, 5, 4, 57, 10,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 5, 5, 72, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 79, 10, 6,
	3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 3, 3, 2, 9, 10, 2, 90, 2, 15, 3, 2, 2,
	2, 4, 18, 3, 2, 2, 2, 6, 56, 3, 2, 2, 2, 8, 71, 3, 2, 2, 2, 10, 73, 3,
	2, 2, 2, 12, 14, 5, 4, 3, 2, 13, 12, 3, 2, 2, 2, 14, 17, 3, 2, 2, 2, 15,
	13, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 3, 3, 2, 2, 2, 17, 15, 3, 2, 2,
	2, 18, 19, 5, 8, 5, 2, 19, 24, 5, 6, 4, 2, 20, 21, 7, 3, 2, 2, 21, 23,
	5, 6, 4, 2, 22, 20, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 22, 3, 2, 2, 2,
	24, 25, 3, 2, 2, 2, 25, 27, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 27, 28, 7,
	25, 2, 2, 28, 5, 3, 2, 2, 2, 29, 30, 7, 4, 2, 2, 30, 57, 7, 24, 2, 2, 31,
	32, 7, 5, 2, 2, 32, 33, 5, 10, 6, 2, 33, 34, 7, 6, 2, 2, 34, 35, 5, 8,
	5, 2, 35, 57, 3, 2, 2, 2, 36, 37, 7, 7, 2, 2, 37, 57, 7, 22, 2, 2, 38,
	39, 7, 8, 2, 2, 39, 57, 7, 22, 2, 2, 40, 57, 9, 2, 2, 2, 41, 42, 7, 11,
	2, 2, 42, 57, 7, 22, 2, 2, 43, 57, 7, 12, 2, 2, 44, 45, 7, 13, 2, 2, 45,
	46, 7, 14, 2, 2, 46, 47, 7, 15, 2, 2, 47, 48, 5, 8, 5, 2, 48, 49, 7, 16,
	2, 2, 49, 57, 3, 2, 2, 2, 50, 54, 7, 17, 2, 2, 51, 55, 5, 10, 6, 2, 52,
	53, 7, 18, 2, 2, 53, 55, 5, 10, 6, 2, 54, 51, 3, 2, 2, 2, 54, 52, 3, 2,
	2, 2, 55, 57, 3, 2, 2, 2, 56, 29, 3, 2, 2, 2, 56, 31, 3, 2, 2, 2, 56, 36,
	3, 2, 2, 2, 56, 38, 3, 2, 2, 2, 56, 40, 3, 2, 2, 2, 56, 41, 3, 2, 2, 2,
	56, 43, 3, 2, 2, 2, 56, 44, 3, 2, 2, 2, 56, 50, 3, 2, 2, 2, 57, 7, 3, 2,
	2, 2, 58, 72, 5, 10, 6, 2, 59, 72, 7, 22, 2, 2, 60, 61, 7, 19, 2, 2, 61,
	62, 7, 15, 2, 2, 62, 63, 5, 8, 5, 2, 63, 64, 7, 16, 2, 2, 64, 72, 3, 2,
	2, 2, 65, 66, 7, 15, 2, 2, 66, 67, 5, 8, 5, 2, 67, 68, 7, 20, 2, 2, 68,
	69, 5, 8, 5, 2, 69, 70, 7, 16, 2, 2, 70, 72, 3, 2, 2, 2, 71, 58, 3, 2,
	2, 2, 71, 59, 3, 2, 2, 2, 71, 60, 3, 2, 2, 2, 71, 65, 3, 2, 2, 2, 72, 9,
	3, 2, 2, 2, 73, 78, 7, 21, 2, 2, 74, 75, 7, 15, 2, 2, 75, 76, 5, 8, 5,
	2, 76, 77, 7, 16, 2, 2, 77, 79, 3, 2, 2, 2, 78, 74, 3, 2, 2, 2, 78, 79,
	3, 2, 2, 2, 79, 11, 3, 2, 2, 2, 8, 15, 24, 54, 56, 71, 78,
}
var literalNames = []string{
	"", "':'", "'REM'", "'LET'", "'='", "'GOTO'", "'GOSUB'", "'RETURN'", "'END'",
	"'PROLONG'", "'CUTSHORT'", "'DIM'", "'RING'", "'('", "')'", "'INPUT'",
	"'CHR$'", "'RND%'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"INTOP", "IINTID", "INTCONST", "STRCONST", "ARBTEXT", "NL", "WS",
}

var ruleNames = []string{
	"sickbay", "line", "stmt", "intExpr", "intVar",
}

type sickbayParser struct {
	*antlr.BaseParser
}

// NewsickbayParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *sickbayParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewsickbayParser(input antlr.TokenStream) *sickbayParser {
	this := new(sickbayParser)
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
	this.GrammarFileName = "sickbay.g4"

	return this
}

// sickbayParser tokens.
const (
	sickbayParserEOF      = antlr.TokenEOF
	sickbayParserT__0     = 1
	sickbayParserT__1     = 2
	sickbayParserT__2     = 3
	sickbayParserT__3     = 4
	sickbayParserT__4     = 5
	sickbayParserT__5     = 6
	sickbayParserT__6     = 7
	sickbayParserT__7     = 8
	sickbayParserT__8     = 9
	sickbayParserT__9     = 10
	sickbayParserT__10    = 11
	sickbayParserT__11    = 12
	sickbayParserT__12    = 13
	sickbayParserT__13    = 14
	sickbayParserT__14    = 15
	sickbayParserT__15    = 16
	sickbayParserT__16    = 17
	sickbayParserINTOP    = 18
	sickbayParserIINTID   = 19
	sickbayParserINTCONST = 20
	sickbayParserSTRCONST = 21
	sickbayParserARBTEXT  = 22
	sickbayParserNL       = 23
	sickbayParserWS       = 24
)

// sickbayParser rules.
const (
	sickbayParserRULE_sickbay = 0
	sickbayParserRULE_line    = 1
	sickbayParserRULE_stmt    = 2
	sickbayParserRULE_intExpr = 3
	sickbayParserRULE_intVar  = 4
)

// ISickbayContext is an interface to support dynamic dispatch.
type ISickbayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSickbayContext differentiates from other interfaces.
	IsSickbayContext()
}

type SickbayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySickbayContext() *SickbayContext {
	var p = new(SickbayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sickbayParserRULE_sickbay
	return p
}

func (*SickbayContext) IsSickbayContext() {}

func NewSickbayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SickbayContext {
	var p = new(SickbayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sickbayParserRULE_sickbay

	return p
}

func (s *SickbayContext) GetParser() antlr.Parser { return s.parser }

func (s *SickbayContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *SickbayContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *SickbayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SickbayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SickbayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sickbayListener); ok {
		listenerT.EnterSickbay(s)
	}
}

func (s *SickbayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sickbayListener); ok {
		listenerT.ExitSickbay(s)
	}
}

func (p *sickbayParser) Sickbay() (localctx ISickbayContext) {
	this := p
	_ = this

	localctx = NewSickbayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, sickbayParserRULE_sickbay)
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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<sickbayParserT__12)|(1<<sickbayParserT__16)|(1<<sickbayParserIINTID)|(1<<sickbayParserINTCONST))) != 0 {
		{
			p.SetState(10)
			p.Line()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sickbayParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sickbayParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) IntExpr() IIntExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntExprContext)
}

func (s *LineContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *LineContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *LineContext) NL() antlr.TerminalNode {
	return s.GetToken(sickbayParserNL, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sickbayListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sickbayListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *sickbayParser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, sickbayParserRULE_line)
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
		p.SetState(16)
		p.IntExpr()
	}
	{
		p.SetState(17)
		p.Stmt()
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == sickbayParserT__0 {
		{
			p.SetState(18)
			p.Match(sickbayParserT__0)
		}
		{
			p.SetState(19)
			p.Stmt()
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.Match(sickbayParserNL)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sickbayParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sickbayParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) ARBTEXT() antlr.TerminalNode {
	return s.GetToken(sickbayParserARBTEXT, 0)
}

func (s *StmtContext) IntVar() IIntVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntVarContext)
}

func (s *StmtContext) IntExpr() IIntExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntExprContext)
}

func (s *StmtContext) INTCONST() antlr.TerminalNode {
	return s.GetToken(sickbayParserINTCONST, 0)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sickbayListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sickbayListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *sickbayParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, sickbayParserRULE_stmt)
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

	p.SetState(54)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sickbayParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Match(sickbayParserT__1)
		}
		{
			p.SetState(28)
			p.Match(sickbayParserARBTEXT)
		}

	case sickbayParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.Match(sickbayParserT__2)
		}
		{
			p.SetState(30)
			p.IntVar()
		}
		{
			p.SetState(31)
			p.Match(sickbayParserT__3)
		}
		{
			p.SetState(32)
			p.IntExpr()
		}

	case sickbayParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.Match(sickbayParserT__4)
		}
		{
			p.SetState(35)
			p.Match(sickbayParserINTCONST)
		}

	case sickbayParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(36)
			p.Match(sickbayParserT__5)
		}
		{
			p.SetState(37)
			p.Match(sickbayParserINTCONST)
		}

	case sickbayParserT__6, sickbayParserT__7:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(38)
			_la = p.GetTokenStream().LA(1)

			if !(_la == sickbayParserT__6 || _la == sickbayParserT__7) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case sickbayParserT__8:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(39)
			p.Match(sickbayParserT__8)
		}
		{
			p.SetState(40)
			p.Match(sickbayParserINTCONST)
		}

	case sickbayParserT__9:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(41)
			p.Match(sickbayParserT__9)
		}

	case sickbayParserT__10:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(42)
			p.Match(sickbayParserT__10)
		}
		{
			p.SetState(43)
			p.Match(sickbayParserT__11)
		}
		{
			p.SetState(44)
			p.Match(sickbayParserT__12)
		}
		{
			p.SetState(45)
			p.IntExpr()
		}
		{
			p.SetState(46)
			p.Match(sickbayParserT__13)
		}

	case sickbayParserT__14:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(48)
			p.Match(sickbayParserT__14)
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case sickbayParserIINTID:
			{
				p.SetState(49)
				p.IntVar()
			}

		case sickbayParserT__15:
			{
				p.SetState(50)
				p.Match(sickbayParserT__15)
			}
			{
				p.SetState(51)
				p.IntVar()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntExprContext is an interface to support dynamic dispatch.
type IIntExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntExprContext differentiates from other interfaces.
	IsIntExprContext()
}

type IntExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntExprContext() *IntExprContext {
	var p = new(IntExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sickbayParserRULE_intExpr
	return p
}

func (*IntExprContext) IsIntExprContext() {}

func NewIntExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntExprContext {
	var p = new(IntExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sickbayParserRULE_intExpr

	return p
}

func (s *IntExprContext) GetParser() antlr.Parser { return s.parser }

func (s *IntExprContext) IntVar() IIntVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntVarContext)
}

func (s *IntExprContext) INTCONST() antlr.TerminalNode {
	return s.GetToken(sickbayParserINTCONST, 0)
}

func (s *IntExprContext) AllIntExpr() []IIntExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntExprContext)(nil)).Elem())
	var tst = make([]IIntExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntExprContext)
		}
	}

	return tst
}

func (s *IntExprContext) IntExpr(i int) IIntExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntExprContext)
}

func (s *IntExprContext) INTOP() antlr.TerminalNode {
	return s.GetToken(sickbayParserINTOP, 0)
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sickbayListener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sickbayListener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (p *sickbayParser) IntExpr() (localctx IIntExprContext) {
	this := p
	_ = this

	localctx = NewIntExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, sickbayParserRULE_intExpr)

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

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sickbayParserIINTID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.IntVar()
		}

	case sickbayParserINTCONST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(sickbayParserINTCONST)
		}

	case sickbayParserT__16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)
			p.Match(sickbayParserT__16)
		}
		{
			p.SetState(59)
			p.Match(sickbayParserT__12)
		}
		{
			p.SetState(60)
			p.IntExpr()
		}
		{
			p.SetState(61)
			p.Match(sickbayParserT__13)
		}

	case sickbayParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(63)
			p.Match(sickbayParserT__12)
		}
		{
			p.SetState(64)
			p.IntExpr()
		}
		{
			p.SetState(65)
			p.Match(sickbayParserINTOP)
		}
		{
			p.SetState(66)
			p.IntExpr()
		}
		{
			p.SetState(67)
			p.Match(sickbayParserT__13)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntVarContext is an interface to support dynamic dispatch.
type IIntVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntVarContext differentiates from other interfaces.
	IsIntVarContext()
}

type IntVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntVarContext() *IntVarContext {
	var p = new(IntVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sickbayParserRULE_intVar
	return p
}

func (*IntVarContext) IsIntVarContext() {}

func NewIntVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntVarContext {
	var p = new(IntVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sickbayParserRULE_intVar

	return p
}

func (s *IntVarContext) GetParser() antlr.Parser { return s.parser }

func (s *IntVarContext) IINTID() antlr.TerminalNode {
	return s.GetToken(sickbayParserIINTID, 0)
}

func (s *IntVarContext) IntExpr() IIntExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntExprContext)
}

func (s *IntVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sickbayListener); ok {
		listenerT.EnterIntVar(s)
	}
}

func (s *IntVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sickbayListener); ok {
		listenerT.ExitIntVar(s)
	}
}

func (p *sickbayParser) IntVar() (localctx IIntVarContext) {
	this := p
	_ = this

	localctx = NewIntVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, sickbayParserRULE_intVar)
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
		p.Match(sickbayParserIINTID)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == sickbayParserT__12 {
		{
			p.SetState(72)
			p.Match(sickbayParserT__12)
		}
		{
			p.SetState(73)
			p.IntExpr()
		}
		{
			p.SetState(74)
			p.Match(sickbayParserT__13)
		}

	}

	return localctx
}
