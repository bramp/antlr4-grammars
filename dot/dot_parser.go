// Generated from DOT.g4 by ANTLR 4.7.

package dot // DOT
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 130,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 5, 2, 32, 10, 2, 3, 2, 3, 2, 5,
	2, 36, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 44, 10, 3, 7, 3,
	46, 10, 3, 12, 3, 14, 3, 49, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 59, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 66, 10,
	6, 3, 6, 6, 6, 69, 10, 6, 13, 6, 14, 6, 70, 3, 7, 3, 7, 3, 7, 5, 7, 76,
	10, 7, 3, 7, 5, 7, 79, 10, 7, 6, 7, 81, 10, 7, 13, 7, 14, 7, 82, 3, 8,
	3, 8, 5, 8, 87, 10, 8, 3, 8, 3, 8, 5, 8, 91, 10, 8, 3, 9, 3, 9, 3, 9, 5,
	9, 96, 10, 9, 6, 9, 98, 10, 9, 13, 9, 14, 9, 99, 3, 10, 3, 10, 3, 11, 3,
	11, 5, 11, 106, 10, 11, 3, 12, 3, 12, 5, 12, 110, 10, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 5, 13, 116, 10, 13, 3, 14, 3, 14, 5, 14, 120, 10, 14, 5,
	14, 122, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 2, 2,
	16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 2, 6, 3, 2, 14,
	15, 4, 2, 14, 14, 16, 17, 3, 2, 10, 11, 3, 2, 19, 22, 2, 137, 2, 31, 3,
	2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 58, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2, 10, 68,
	3, 2, 2, 2, 12, 80, 3, 2, 2, 2, 14, 86, 3, 2, 2, 2, 16, 97, 3, 2, 2, 2,
	18, 101, 3, 2, 2, 2, 20, 103, 3, 2, 2, 2, 22, 107, 3, 2, 2, 2, 24, 111,
	3, 2, 2, 2, 26, 121, 3, 2, 2, 2, 28, 127, 3, 2, 2, 2, 30, 32, 7, 13, 2,
	2, 31, 30, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 35,
	9, 2, 2, 2, 34, 36, 5, 28, 15, 2, 35, 34, 3, 2, 2, 2, 35, 36, 3, 2, 2,
	2, 36, 37, 3, 2, 2, 2, 37, 38, 7, 3, 2, 2, 38, 39, 5, 4, 3, 2, 39, 40,
	7, 4, 2, 2, 40, 3, 3, 2, 2, 2, 41, 43, 5, 6, 4, 2, 42, 44, 7, 5, 2, 2,
	43, 42, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 41, 3,
	2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48,
	5, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 59, 5, 20, 11, 2, 51, 59, 5, 14,
	8, 2, 52, 59, 5, 8, 5, 2, 53, 54, 5, 28, 15, 2, 54, 55, 7, 6, 2, 2, 55,
	56, 5, 28, 15, 2, 56, 59, 3, 2, 2, 2, 57, 59, 5, 26, 14, 2, 58, 50, 3,
	2, 2, 2, 58, 51, 3, 2, 2, 2, 58, 52, 3, 2, 2, 2, 58, 53, 3, 2, 2, 2, 58,
	57, 3, 2, 2, 2, 59, 7, 3, 2, 2, 2, 60, 61, 9, 3, 2, 2, 61, 62, 5, 10, 6,
	2, 62, 9, 3, 2, 2, 2, 63, 65, 7, 7, 2, 2, 64, 66, 5, 12, 7, 2, 65, 64,
	3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 69, 7, 8, 2, 2,
	68, 63, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3,
	2, 2, 2, 71, 11, 3, 2, 2, 2, 72, 75, 5, 28, 15, 2, 73, 74, 7, 6, 2, 2,
	74, 76, 5, 28, 15, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78, 3,
	2, 2, 2, 77, 79, 7, 9, 2, 2, 78, 77, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79,
	81, 3, 2, 2, 2, 80, 72, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2,
	2, 82, 83, 3, 2, 2, 2, 83, 13, 3, 2, 2, 2, 84, 87, 5, 22, 12, 2, 85, 87,
	5, 26, 14, 2, 86, 84, 3, 2, 2, 2, 86, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2,
	2, 88, 90, 5, 16, 9, 2, 89, 91, 5, 10, 6, 2, 90, 89, 3, 2, 2, 2, 90, 91,
	3, 2, 2, 2, 91, 15, 3, 2, 2, 2, 92, 95, 5, 18, 10, 2, 93, 96, 5, 22, 12,
	2, 94, 96, 5, 26, 14, 2, 95, 93, 3, 2, 2, 2, 95, 94, 3, 2, 2, 2, 96, 98,
	3, 2, 2, 2, 97, 92, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2,
	99, 100, 3, 2, 2, 2, 100, 17, 3, 2, 2, 2, 101, 102, 9, 4, 2, 2, 102, 19,
	3, 2, 2, 2, 103, 105, 5, 22, 12, 2, 104, 106, 5, 10, 6, 2, 105, 104, 3,
	2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 21, 3, 2, 2, 2, 107, 109, 5, 28, 15,
	2, 108, 110, 5, 24, 13, 2, 109, 108, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2,
	110, 23, 3, 2, 2, 2, 111, 112, 7, 12, 2, 2, 112, 115, 5, 28, 15, 2, 113,
	114, 7, 12, 2, 2, 114, 116, 5, 28, 15, 2, 115, 113, 3, 2, 2, 2, 115, 116,
	3, 2, 2, 2, 116, 25, 3, 2, 2, 2, 117, 119, 7, 18, 2, 2, 118, 120, 5, 28,
	15, 2, 119, 118, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 122, 3, 2, 2, 2,
	121, 117, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	124, 7, 3, 2, 2, 124, 125, 5, 4, 3, 2, 125, 126, 7, 4, 2, 2, 126, 27, 3,
	2, 2, 2, 127, 128, 9, 5, 2, 2, 128, 29, 3, 2, 2, 2, 21, 31, 35, 43, 47,
	58, 65, 70, 75, 78, 82, 86, 90, 95, 99, 105, 109, 115, 119, 121,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "';'", "'='", "'['", "']'", "','", "'->'", "'--'", "':'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "STRICT", "GRAPH", "DIGRAPH",
	"NODE", "EDGE", "SUBGRAPH", "NUMBER", "STRING", "ID", "HTML_STRING", "COMMENT",
	"LINE_COMMENT", "PREPROC", "WS",
}

var ruleNames = []string{
	"graph", "stmt_list", "stmt", "attr_stmt", "attr_list", "a_list", "edge_stmt",
	"edgeRHS", "edgeop", "node_stmt", "node_id", "port", "subgraph", "id",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DOTParser struct {
	*antlr.BaseParser
}

func NewDOTParser(input antlr.TokenStream) *DOTParser {
	this := new(DOTParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DOT.g4"

	return this
}

// DOTParser tokens.
const (
	DOTParserEOF          = antlr.TokenEOF
	DOTParserT__0         = 1
	DOTParserT__1         = 2
	DOTParserT__2         = 3
	DOTParserT__3         = 4
	DOTParserT__4         = 5
	DOTParserT__5         = 6
	DOTParserT__6         = 7
	DOTParserT__7         = 8
	DOTParserT__8         = 9
	DOTParserT__9         = 10
	DOTParserSTRICT       = 11
	DOTParserGRAPH        = 12
	DOTParserDIGRAPH      = 13
	DOTParserNODE         = 14
	DOTParserEDGE         = 15
	DOTParserSUBGRAPH     = 16
	DOTParserNUMBER       = 17
	DOTParserSTRING       = 18
	DOTParserID           = 19
	DOTParserHTML_STRING  = 20
	DOTParserCOMMENT      = 21
	DOTParserLINE_COMMENT = 22
	DOTParserPREPROC      = 23
	DOTParserWS           = 24
)

// DOTParser rules.
const (
	DOTParserRULE_graph     = 0
	DOTParserRULE_stmt_list = 1
	DOTParserRULE_stmt      = 2
	DOTParserRULE_attr_stmt = 3
	DOTParserRULE_attr_list = 4
	DOTParserRULE_a_list    = 5
	DOTParserRULE_edge_stmt = 6
	DOTParserRULE_edgeRHS   = 7
	DOTParserRULE_edgeop    = 8
	DOTParserRULE_node_stmt = 9
	DOTParserRULE_node_id   = 10
	DOTParserRULE_port      = 11
	DOTParserRULE_subgraph  = 12
	DOTParserRULE_id        = 13
)

// IGraphContext is an interface to support dynamic dispatch.
type IGraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGraphContext differentiates from other interfaces.
	IsGraphContext()
}

type GraphContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphContext() *GraphContext {
	var p = new(GraphContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_graph
	return p
}

func (*GraphContext) IsGraphContext() {}

func NewGraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphContext {
	var p = new(GraphContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_graph

	return p
}

func (s *GraphContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphContext) Stmt_list() IStmt_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmt_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmt_listContext)
}

func (s *GraphContext) GRAPH() antlr.TerminalNode {
	return s.GetToken(DOTParserGRAPH, 0)
}

func (s *GraphContext) DIGRAPH() antlr.TerminalNode {
	return s.GetToken(DOTParserDIGRAPH, 0)
}

func (s *GraphContext) STRICT() antlr.TerminalNode {
	return s.GetToken(DOTParserSTRICT, 0)
}

func (s *GraphContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *GraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterGraph(s)
	}
}

func (s *GraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitGraph(s)
	}
}

func (s *GraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitGraph(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Graph() (localctx IGraphContext) {
	localctx = NewGraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DOTParserRULE_graph)
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

	if _la == DOTParserSTRICT {
		{
			p.SetState(28)
			p.Match(DOTParserSTRICT)
		}

	}
	p.SetState(31)
	_la = p.GetTokenStream().LA(1)

	if !(_la == DOTParserGRAPH || _la == DOTParserDIGRAPH) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DOTParserNUMBER)|(1<<DOTParserSTRING)|(1<<DOTParserID)|(1<<DOTParserHTML_STRING))) != 0 {
		{
			p.SetState(32)
			p.Id()
		}

	}
	{
		p.SetState(35)
		p.Match(DOTParserT__0)
	}
	{
		p.SetState(36)
		p.Stmt_list()
	}
	{
		p.SetState(37)
		p.Match(DOTParserT__1)
	}

	return localctx
}

// IStmt_listContext is an interface to support dynamic dispatch.
type IStmt_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmt_listContext differentiates from other interfaces.
	IsStmt_listContext()
}

type Stmt_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmt_listContext() *Stmt_listContext {
	var p = new(Stmt_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_stmt_list
	return p
}

func (*Stmt_listContext) IsStmt_listContext() {}

func NewStmt_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stmt_listContext {
	var p = new(Stmt_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_stmt_list

	return p
}

func (s *Stmt_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Stmt_listContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *Stmt_listContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *Stmt_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stmt_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterStmt_list(s)
	}
}

func (s *Stmt_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitStmt_list(s)
	}
}

func (s *Stmt_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitStmt_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Stmt_list() (localctx IStmt_listContext) {
	localctx = NewStmt_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DOTParserRULE_stmt_list)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DOTParserT__0)|(1<<DOTParserGRAPH)|(1<<DOTParserNODE)|(1<<DOTParserEDGE)|(1<<DOTParserSUBGRAPH)|(1<<DOTParserNUMBER)|(1<<DOTParserSTRING)|(1<<DOTParserID)|(1<<DOTParserHTML_STRING))) != 0 {
		{
			p.SetState(39)
			p.Stmt()
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DOTParserT__2 {
			{
				p.SetState(40)
				p.Match(DOTParserT__2)
			}

		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = DOTParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Node_stmt() INode_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INode_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INode_stmtContext)
}

func (s *StmtContext) Edge_stmt() IEdge_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdge_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdge_stmtContext)
}

func (s *StmtContext) Attr_stmt() IAttr_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttr_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttr_stmtContext)
}

func (s *StmtContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *StmtContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *StmtContext) Subgraph() ISubgraphContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubgraphContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubgraphContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DOTParserRULE_stmt)

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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Node_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Edge_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Attr_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.Id()
		}
		{
			p.SetState(52)
			p.Match(DOTParserT__3)
		}
		{
			p.SetState(53)
			p.Id()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(55)
			p.Subgraph()
		}

	}

	return localctx
}

// IAttr_stmtContext is an interface to support dynamic dispatch.
type IAttr_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttr_stmtContext differentiates from other interfaces.
	IsAttr_stmtContext()
}

type Attr_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttr_stmtContext() *Attr_stmtContext {
	var p = new(Attr_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_attr_stmt
	return p
}

func (*Attr_stmtContext) IsAttr_stmtContext() {}

func NewAttr_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Attr_stmtContext {
	var p = new(Attr_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_attr_stmt

	return p
}

func (s *Attr_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Attr_stmtContext) Attr_list() IAttr_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttr_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttr_listContext)
}

func (s *Attr_stmtContext) GRAPH() antlr.TerminalNode {
	return s.GetToken(DOTParserGRAPH, 0)
}

func (s *Attr_stmtContext) NODE() antlr.TerminalNode {
	return s.GetToken(DOTParserNODE, 0)
}

func (s *Attr_stmtContext) EDGE() antlr.TerminalNode {
	return s.GetToken(DOTParserEDGE, 0)
}

func (s *Attr_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Attr_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Attr_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterAttr_stmt(s)
	}
}

func (s *Attr_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitAttr_stmt(s)
	}
}

func (s *Attr_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitAttr_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Attr_stmt() (localctx IAttr_stmtContext) {
	localctx = NewAttr_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DOTParserRULE_attr_stmt)
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
	p.SetState(58)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DOTParserGRAPH)|(1<<DOTParserNODE)|(1<<DOTParserEDGE))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(59)
		p.Attr_list()
	}

	return localctx
}

// IAttr_listContext is an interface to support dynamic dispatch.
type IAttr_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttr_listContext differentiates from other interfaces.
	IsAttr_listContext()
}

type Attr_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttr_listContext() *Attr_listContext {
	var p = new(Attr_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_attr_list
	return p
}

func (*Attr_listContext) IsAttr_listContext() {}

func NewAttr_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Attr_listContext {
	var p = new(Attr_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_attr_list

	return p
}

func (s *Attr_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Attr_listContext) AllA_list() []IA_listContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IA_listContext)(nil)).Elem())
	var tst = make([]IA_listContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IA_listContext)
		}
	}

	return tst
}

func (s *Attr_listContext) A_list(i int) IA_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IA_listContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IA_listContext)
}

func (s *Attr_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Attr_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Attr_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterAttr_list(s)
	}
}

func (s *Attr_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitAttr_list(s)
	}
}

func (s *Attr_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitAttr_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Attr_list() (localctx IAttr_listContext) {
	localctx = NewAttr_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DOTParserRULE_attr_list)
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
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DOTParserT__4 {
		{
			p.SetState(61)
			p.Match(DOTParserT__4)
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DOTParserNUMBER)|(1<<DOTParserSTRING)|(1<<DOTParserID)|(1<<DOTParserHTML_STRING))) != 0 {
			{
				p.SetState(62)
				p.A_list()
			}

		}
		{
			p.SetState(65)
			p.Match(DOTParserT__5)
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IA_listContext is an interface to support dynamic dispatch.
type IA_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsA_listContext differentiates from other interfaces.
	IsA_listContext()
}

type A_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_listContext() *A_listContext {
	var p = new(A_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_a_list
	return p
}

func (*A_listContext) IsA_listContext() {}

func NewA_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_listContext {
	var p = new(A_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_a_list

	return p
}

func (s *A_listContext) GetParser() antlr.Parser { return s.parser }

func (s *A_listContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *A_listContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *A_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterA_list(s)
	}
}

func (s *A_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitA_list(s)
	}
}

func (s *A_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitA_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) A_list() (localctx IA_listContext) {
	localctx = NewA_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DOTParserRULE_a_list)
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
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DOTParserNUMBER)|(1<<DOTParserSTRING)|(1<<DOTParserID)|(1<<DOTParserHTML_STRING))) != 0) {
		{
			p.SetState(70)
			p.Id()
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DOTParserT__3 {
			{
				p.SetState(71)
				p.Match(DOTParserT__3)
			}
			{
				p.SetState(72)
				p.Id()
			}

		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DOTParserT__6 {
			{
				p.SetState(75)
				p.Match(DOTParserT__6)
			}

		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEdge_stmtContext is an interface to support dynamic dispatch.
type IEdge_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdge_stmtContext differentiates from other interfaces.
	IsEdge_stmtContext()
}

type Edge_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdge_stmtContext() *Edge_stmtContext {
	var p = new(Edge_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_edge_stmt
	return p
}

func (*Edge_stmtContext) IsEdge_stmtContext() {}

func NewEdge_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Edge_stmtContext {
	var p = new(Edge_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_edge_stmt

	return p
}

func (s *Edge_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Edge_stmtContext) EdgeRHS() IEdgeRHSContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdgeRHSContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEdgeRHSContext)
}

func (s *Edge_stmtContext) Node_id() INode_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INode_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INode_idContext)
}

func (s *Edge_stmtContext) Subgraph() ISubgraphContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubgraphContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubgraphContext)
}

func (s *Edge_stmtContext) Attr_list() IAttr_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttr_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttr_listContext)
}

func (s *Edge_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Edge_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Edge_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterEdge_stmt(s)
	}
}

func (s *Edge_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitEdge_stmt(s)
	}
}

func (s *Edge_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitEdge_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Edge_stmt() (localctx IEdge_stmtContext) {
	localctx = NewEdge_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DOTParserRULE_edge_stmt)
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
	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DOTParserNUMBER, DOTParserSTRING, DOTParserID, DOTParserHTML_STRING:
		{
			p.SetState(82)
			p.Node_id()
		}

	case DOTParserT__0, DOTParserSUBGRAPH:
		{
			p.SetState(83)
			p.Subgraph()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(86)
		p.EdgeRHS()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DOTParserT__4 {
		{
			p.SetState(87)
			p.Attr_list()
		}

	}

	return localctx
}

// IEdgeRHSContext is an interface to support dynamic dispatch.
type IEdgeRHSContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgeRHSContext differentiates from other interfaces.
	IsEdgeRHSContext()
}

type EdgeRHSContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgeRHSContext() *EdgeRHSContext {
	var p = new(EdgeRHSContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_edgeRHS
	return p
}

func (*EdgeRHSContext) IsEdgeRHSContext() {}

func NewEdgeRHSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgeRHSContext {
	var p = new(EdgeRHSContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_edgeRHS

	return p
}

func (s *EdgeRHSContext) GetParser() antlr.Parser { return s.parser }

func (s *EdgeRHSContext) AllEdgeop() []IEdgeopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEdgeopContext)(nil)).Elem())
	var tst = make([]IEdgeopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEdgeopContext)
		}
	}

	return tst
}

func (s *EdgeRHSContext) Edgeop(i int) IEdgeopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEdgeopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEdgeopContext)
}

func (s *EdgeRHSContext) AllNode_id() []INode_idContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INode_idContext)(nil)).Elem())
	var tst = make([]INode_idContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INode_idContext)
		}
	}

	return tst
}

func (s *EdgeRHSContext) Node_id(i int) INode_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INode_idContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INode_idContext)
}

func (s *EdgeRHSContext) AllSubgraph() []ISubgraphContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubgraphContext)(nil)).Elem())
	var tst = make([]ISubgraphContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubgraphContext)
		}
	}

	return tst
}

func (s *EdgeRHSContext) Subgraph(i int) ISubgraphContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubgraphContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubgraphContext)
}

func (s *EdgeRHSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeRHSContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EdgeRHSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterEdgeRHS(s)
	}
}

func (s *EdgeRHSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitEdgeRHS(s)
	}
}

func (s *EdgeRHSContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitEdgeRHS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) EdgeRHS() (localctx IEdgeRHSContext) {
	localctx = NewEdgeRHSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DOTParserRULE_edgeRHS)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DOTParserT__7 || _la == DOTParserT__8 {
		{
			p.SetState(90)
			p.Edgeop()
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DOTParserNUMBER, DOTParserSTRING, DOTParserID, DOTParserHTML_STRING:
			{
				p.SetState(91)
				p.Node_id()
			}

		case DOTParserT__0, DOTParserSUBGRAPH:
			{
				p.SetState(92)
				p.Subgraph()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEdgeopContext is an interface to support dynamic dispatch.
type IEdgeopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgeopContext differentiates from other interfaces.
	IsEdgeopContext()
}

type EdgeopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgeopContext() *EdgeopContext {
	var p = new(EdgeopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_edgeop
	return p
}

func (*EdgeopContext) IsEdgeopContext() {}

func NewEdgeopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgeopContext {
	var p = new(EdgeopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_edgeop

	return p
}

func (s *EdgeopContext) GetParser() antlr.Parser { return s.parser }
func (s *EdgeopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EdgeopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterEdgeop(s)
	}
}

func (s *EdgeopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitEdgeop(s)
	}
}

func (s *EdgeopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitEdgeop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Edgeop() (localctx IEdgeopContext) {
	localctx = NewEdgeopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DOTParserRULE_edgeop)
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
	p.SetState(99)
	_la = p.GetTokenStream().LA(1)

	if !(_la == DOTParserT__7 || _la == DOTParserT__8) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// INode_stmtContext is an interface to support dynamic dispatch.
type INode_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNode_stmtContext differentiates from other interfaces.
	IsNode_stmtContext()
}

type Node_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNode_stmtContext() *Node_stmtContext {
	var p = new(Node_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_node_stmt
	return p
}

func (*Node_stmtContext) IsNode_stmtContext() {}

func NewNode_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Node_stmtContext {
	var p = new(Node_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_node_stmt

	return p
}

func (s *Node_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Node_stmtContext) Node_id() INode_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INode_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INode_idContext)
}

func (s *Node_stmtContext) Attr_list() IAttr_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttr_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttr_listContext)
}

func (s *Node_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Node_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Node_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterNode_stmt(s)
	}
}

func (s *Node_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitNode_stmt(s)
	}
}

func (s *Node_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitNode_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Node_stmt() (localctx INode_stmtContext) {
	localctx = NewNode_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DOTParserRULE_node_stmt)
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
		p.SetState(101)
		p.Node_id()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DOTParserT__4 {
		{
			p.SetState(102)
			p.Attr_list()
		}

	}

	return localctx
}

// INode_idContext is an interface to support dynamic dispatch.
type INode_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNode_idContext differentiates from other interfaces.
	IsNode_idContext()
}

type Node_idContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNode_idContext() *Node_idContext {
	var p = new(Node_idContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_node_id
	return p
}

func (*Node_idContext) IsNode_idContext() {}

func NewNode_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Node_idContext {
	var p = new(Node_idContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_node_id

	return p
}

func (s *Node_idContext) GetParser() antlr.Parser { return s.parser }

func (s *Node_idContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *Node_idContext) Port() IPortContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortContext)
}

func (s *Node_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Node_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Node_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterNode_id(s)
	}
}

func (s *Node_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitNode_id(s)
	}
}

func (s *Node_idContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitNode_id(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Node_id() (localctx INode_idContext) {
	localctx = NewNode_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DOTParserRULE_node_id)
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
		p.SetState(105)
		p.Id()
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DOTParserT__9 {
		{
			p.SetState(106)
			p.Port()
		}

	}

	return localctx
}

// IPortContext is an interface to support dynamic dispatch.
type IPortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPortContext differentiates from other interfaces.
	IsPortContext()
}

type PortContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortContext() *PortContext {
	var p = new(PortContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_port
	return p
}

func (*PortContext) IsPortContext() {}

func NewPortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortContext {
	var p = new(PortContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_port

	return p
}

func (s *PortContext) GetParser() antlr.Parser { return s.parser }

func (s *PortContext) AllId() []IIdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdContext)(nil)).Elem())
	var tst = make([]IIdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdContext)
		}
	}

	return tst
}

func (s *PortContext) Id(i int) IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *PortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterPort(s)
	}
}

func (s *PortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitPort(s)
	}
}

func (s *PortContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitPort(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Port() (localctx IPortContext) {
	localctx = NewPortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DOTParserRULE_port)
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
		p.SetState(109)
		p.Match(DOTParserT__9)
	}
	{
		p.SetState(110)
		p.Id()
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DOTParserT__9 {
		{
			p.SetState(111)
			p.Match(DOTParserT__9)
		}
		{
			p.SetState(112)
			p.Id()
		}

	}

	return localctx
}

// ISubgraphContext is an interface to support dynamic dispatch.
type ISubgraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubgraphContext differentiates from other interfaces.
	IsSubgraphContext()
}

type SubgraphContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubgraphContext() *SubgraphContext {
	var p = new(SubgraphContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DOTParserRULE_subgraph
	return p
}

func (*SubgraphContext) IsSubgraphContext() {}

func NewSubgraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubgraphContext {
	var p = new(SubgraphContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_subgraph

	return p
}

func (s *SubgraphContext) GetParser() antlr.Parser { return s.parser }

func (s *SubgraphContext) Stmt_list() IStmt_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmt_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmt_listContext)
}

func (s *SubgraphContext) SUBGRAPH() antlr.TerminalNode {
	return s.GetToken(DOTParserSUBGRAPH, 0)
}

func (s *SubgraphContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *SubgraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubgraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubgraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterSubgraph(s)
	}
}

func (s *SubgraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitSubgraph(s)
	}
}

func (s *SubgraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitSubgraph(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Subgraph() (localctx ISubgraphContext) {
	localctx = NewSubgraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DOTParserRULE_subgraph)
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DOTParserSUBGRAPH {
		{
			p.SetState(115)
			p.Match(DOTParserSUBGRAPH)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DOTParserNUMBER)|(1<<DOTParserSTRING)|(1<<DOTParserID)|(1<<DOTParserHTML_STRING))) != 0 {
			{
				p.SetState(116)
				p.Id()
			}

		}

	}
	{
		p.SetState(121)
		p.Match(DOTParserT__0)
	}
	{
		p.SetState(122)
		p.Stmt_list()
	}
	{
		p.SetState(123)
		p.Match(DOTParserT__1)
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
	p.RuleIndex = DOTParserRULE_id
	return p
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DOTParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(DOTParserID, 0)
}

func (s *IdContext) STRING() antlr.TerminalNode {
	return s.GetToken(DOTParserSTRING, 0)
}

func (s *IdContext) HTML_STRING() antlr.TerminalNode {
	return s.GetToken(DOTParserHTML_STRING, 0)
}

func (s *IdContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DOTParserNUMBER, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DOTListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DOTVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DOTParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DOTParserRULE_id)
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
	p.SetState(125)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DOTParserNUMBER)|(1<<DOTParserSTRING)|(1<<DOTParserID)|(1<<DOTParserHTML_STRING))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
