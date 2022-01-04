// Code generated from cayenne.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cayenne // cayenne
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 24, 161,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 46, 10, 2, 12, 2,
	14, 2, 49, 11, 2, 3, 2, 3, 2, 7, 2, 53, 10, 2, 12, 2, 14, 2, 56, 11, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 66, 10, 2, 12, 2,
	14, 2, 69, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 76, 10, 2, 12, 2,
	14, 2, 79, 11, 2, 3, 2, 3, 2, 7, 2, 83, 10, 2, 12, 2, 14, 2, 86, 11, 2,
	3, 2, 3, 2, 5, 2, 90, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 97, 10,
	2, 12, 2, 14, 2, 100, 11, 2, 3, 3, 3, 3, 3, 3, 7, 3, 105, 10, 3, 12, 3,
	14, 3, 108, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 120, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 134, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 147, 10, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 2, 3,
	2, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 3, 3, 2, 21, 22, 2, 168,
	2, 89, 3, 2, 2, 2, 4, 119, 3, 2, 2, 2, 6, 133, 3, 2, 2, 2, 8, 135, 3, 2,
	2, 2, 10, 146, 3, 2, 2, 2, 12, 148, 3, 2, 2, 2, 14, 150, 3, 2, 2, 2, 16,
	152, 3, 2, 2, 2, 18, 154, 3, 2, 2, 2, 20, 156, 3, 2, 2, 2, 22, 158, 3,
	2, 2, 2, 24, 25, 8, 2, 1, 2, 25, 26, 7, 3, 2, 2, 26, 27, 5, 16, 9, 2, 27,
	28, 7, 4, 2, 2, 28, 29, 5, 14, 8, 2, 29, 30, 7, 5, 2, 2, 30, 31, 7, 6,
	2, 2, 31, 32, 5, 2, 2, 13, 32, 90, 3, 2, 2, 2, 33, 34, 7, 7, 2, 2, 34,
	35, 7, 3, 2, 2, 35, 36, 5, 16, 9, 2, 36, 37, 7, 4, 2, 2, 37, 38, 5, 14,
	8, 2, 38, 39, 7, 5, 2, 2, 39, 40, 7, 6, 2, 2, 40, 41, 5, 2, 2, 12, 41,
	90, 3, 2, 2, 2, 42, 54, 7, 8, 2, 2, 43, 47, 5, 18, 10, 2, 44, 46, 5, 14,
	8, 2, 45, 44, 3, 2, 2, 2, 46, 49, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48,
	3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 51, 7, 9, 2, 2,
	51, 53, 3, 2, 2, 2, 52, 43, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52, 3,
	2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 90, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 57,
	58, 5, 18, 10, 2, 58, 59, 7, 10, 2, 2, 59, 60, 5, 14, 8, 2, 60, 90, 3,
	2, 2, 2, 61, 62, 7, 11, 2, 2, 62, 63, 5, 16, 9, 2, 63, 67, 7, 12, 2, 2,
	64, 66, 5, 4, 3, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3,
	2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70,
	71, 7, 4, 2, 2, 71, 72, 5, 14, 8, 2, 72, 90, 3, 2, 2, 2, 73, 77, 7, 13,
	2, 2, 74, 76, 5, 6, 4, 2, 75, 74, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75,
	3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 90, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2,
	80, 84, 7, 14, 2, 2, 81, 83, 5, 8, 5, 2, 82, 81, 3, 2, 2, 2, 83, 86, 3,
	2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 90, 3, 2, 2, 2, 86,
	84, 3, 2, 2, 2, 87, 90, 5, 22, 12, 2, 88, 90, 7, 16, 2, 2, 89, 24, 3, 2,
	2, 2, 89, 33, 3, 2, 2, 2, 89, 42, 3, 2, 2, 2, 89, 57, 3, 2, 2, 2, 89, 61,
	3, 2, 2, 2, 89, 73, 3, 2, 2, 2, 89, 80, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2,
	89, 88, 3, 2, 2, 2, 90, 98, 3, 2, 2, 2, 91, 92, 12, 11, 2, 2, 92, 97, 5,
	2, 2, 12, 93, 94, 12, 5, 2, 2, 94, 95, 7, 15, 2, 2, 95, 97, 5, 20, 11,
	2, 96, 91, 3, 2, 2, 2, 96, 93, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98, 96,
	3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 3, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2,
	101, 102, 7, 3, 2, 2, 102, 106, 5, 18, 10, 2, 103, 105, 5, 16, 9, 2, 104,
	103, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107,
	3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 110, 7, 5,
	2, 2, 110, 111, 7, 6, 2, 2, 111, 112, 5, 2, 2, 2, 112, 113, 7, 17, 2, 2,
	113, 120, 3, 2, 2, 2, 114, 115, 5, 16, 9, 2, 115, 116, 7, 6, 2, 2, 116,
	117, 5, 2, 2, 2, 117, 118, 7, 17, 2, 2, 118, 120, 3, 2, 2, 2, 119, 101,
	3, 2, 2, 2, 119, 114, 3, 2, 2, 2, 120, 5, 3, 2, 2, 2, 121, 122, 5, 20,
	11, 2, 122, 123, 7, 4, 2, 2, 123, 124, 5, 14, 8, 2, 124, 125, 7, 17, 2,
	2, 125, 134, 3, 2, 2, 2, 126, 127, 5, 20, 11, 2, 127, 128, 7, 4, 2, 2,
	128, 129, 5, 14, 8, 2, 129, 130, 7, 18, 2, 2, 130, 131, 5, 2, 2, 2, 131,
	132, 7, 17, 2, 2, 132, 134, 3, 2, 2, 2, 133, 121, 3, 2, 2, 2, 133, 126,
	3, 2, 2, 2, 134, 7, 3, 2, 2, 2, 135, 136, 5, 10, 6, 2, 136, 137, 5, 20,
	11, 2, 137, 138, 7, 4, 2, 2, 138, 139, 5, 14, 8, 2, 139, 140, 7, 18, 2,
	2, 140, 141, 5, 2, 2, 2, 141, 142, 7, 17, 2, 2, 142, 9, 3, 2, 2, 2, 143,
	147, 7, 19, 2, 2, 144, 145, 7, 20, 2, 2, 145, 147, 5, 12, 7, 2, 146, 143,
	3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 147, 11, 3, 2, 2, 2, 148, 149, 9, 2,
	2, 2, 149, 13, 3, 2, 2, 2, 150, 151, 5, 2, 2, 2, 151, 15, 3, 2, 2, 2, 152,
	153, 5, 22, 12, 2, 153, 17, 3, 2, 2, 2, 154, 155, 5, 22, 12, 2, 155, 19,
	3, 2, 2, 2, 156, 157, 5, 22, 12, 2, 157, 21, 3, 2, 2, 2, 158, 159, 7, 23,
	2, 2, 159, 23, 3, 2, 2, 2, 14, 47, 54, 67, 77, 84, 89, 96, 98, 106, 119,
	133, 146,
}
var literalNames = []string{
	"", "'('", "'::'", "')'", "'->'", "'\\'", "'data'", "'|'", "'@'", "'case'",
	"'of'", "'sig'", "'struct'", "'.'", "'#'", "';'", "'='", "'private'", "'public'",
	"'abstract'", "'concrete'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "ID", "WS",
}

var ruleNames = []string{
	"expr", "arm", "sign", "defn", "vis", "abs_", "type_", "varid", "conid",
	"lblid", "id_",
}

type cayenneParser struct {
	*antlr.BaseParser
}

// NewcayenneParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *cayenneParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewcayenneParser(input antlr.TokenStream) *cayenneParser {
	this := new(cayenneParser)
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
	this.GrammarFileName = "cayenne.g4"

	return this
}

// cayenneParser tokens.
const (
	cayenneParserEOF   = antlr.TokenEOF
	cayenneParserT__0  = 1
	cayenneParserT__1  = 2
	cayenneParserT__2  = 3
	cayenneParserT__3  = 4
	cayenneParserT__4  = 5
	cayenneParserT__5  = 6
	cayenneParserT__6  = 7
	cayenneParserT__7  = 8
	cayenneParserT__8  = 9
	cayenneParserT__9  = 10
	cayenneParserT__10 = 11
	cayenneParserT__11 = 12
	cayenneParserT__12 = 13
	cayenneParserT__13 = 14
	cayenneParserT__14 = 15
	cayenneParserT__15 = 16
	cayenneParserT__16 = 17
	cayenneParserT__17 = 18
	cayenneParserT__18 = 19
	cayenneParserT__19 = 20
	cayenneParserID    = 21
	cayenneParserWS    = 22
)

// cayenneParser rules.
const (
	cayenneParserRULE_expr  = 0
	cayenneParserRULE_arm   = 1
	cayenneParserRULE_sign  = 2
	cayenneParserRULE_defn  = 3
	cayenneParserRULE_vis   = 4
	cayenneParserRULE_abs_  = 5
	cayenneParserRULE_type_ = 6
	cayenneParserRULE_varid = 7
	cayenneParserRULE_conid = 8
	cayenneParserRULE_lblid = 9
	cayenneParserRULE_id_   = 10
)

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
	p.RuleIndex = cayenneParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Varid() IVaridContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVaridContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVaridContext)
}

func (s *ExprContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *ExprContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) AllConid() []IConidContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConidContext)(nil)).Elem())
	var tst = make([]IConidContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConidContext)
		}
	}

	return tst
}

func (s *ExprContext) Conid(i int) IConidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConidContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConidContext)
}

func (s *ExprContext) AllArm() []IArmContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArmContext)(nil)).Elem())
	var tst = make([]IArmContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArmContext)
		}
	}

	return tst
}

func (s *ExprContext) Arm(i int) IArmContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArmContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArmContext)
}

func (s *ExprContext) AllSign() []ISignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISignContext)(nil)).Elem())
	var tst = make([]ISignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISignContext)
		}
	}

	return tst
}

func (s *ExprContext) Sign(i int) ISignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISignContext)
}

func (s *ExprContext) AllDefn() []IDefnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefnContext)(nil)).Elem())
	var tst = make([]IDefnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefnContext)
		}
	}

	return tst
}

func (s *ExprContext) Defn(i int) IDefnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefnContext)
}

func (s *ExprContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *ExprContext) Lblid() ILblidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILblidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILblidContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *cayenneParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *cayenneParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, cayenneParserRULE_expr, _p)
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(23)
			p.Match(cayenneParserT__0)
		}
		{
			p.SetState(24)
			p.Varid()
		}
		{
			p.SetState(25)
			p.Match(cayenneParserT__1)
		}
		{
			p.SetState(26)
			p.Type_()
		}
		{
			p.SetState(27)
			p.Match(cayenneParserT__2)
		}
		{
			p.SetState(28)
			p.Match(cayenneParserT__3)
		}
		{
			p.SetState(29)
			p.expr(11)
		}

	case 2:
		{
			p.SetState(31)
			p.Match(cayenneParserT__4)
		}
		{
			p.SetState(32)
			p.Match(cayenneParserT__0)
		}
		{
			p.SetState(33)
			p.Varid()
		}
		{
			p.SetState(34)
			p.Match(cayenneParserT__1)
		}
		{
			p.SetState(35)
			p.Type_()
		}
		{
			p.SetState(36)
			p.Match(cayenneParserT__2)
		}
		{
			p.SetState(37)
			p.Match(cayenneParserT__3)
		}
		{
			p.SetState(38)
			p.expr(10)
		}

	case 3:
		{
			p.SetState(40)
			p.Match(cayenneParserT__5)
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(41)
					p.Conid()
				}
				p.SetState(45)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<cayenneParserT__0)|(1<<cayenneParserT__4)|(1<<cayenneParserT__5)|(1<<cayenneParserT__8)|(1<<cayenneParserT__10)|(1<<cayenneParserT__11)|(1<<cayenneParserT__13)|(1<<cayenneParserID))) != 0 {
					{
						p.SetState(42)
						p.Type_()
					}

					p.SetState(47)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(48)
					p.Match(cayenneParserT__6)
				}

			}
			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}

	case 4:
		{
			p.SetState(55)
			p.Conid()
		}
		{
			p.SetState(56)
			p.Match(cayenneParserT__7)
		}
		{
			p.SetState(57)
			p.Type_()
		}

	case 5:
		{
			p.SetState(59)
			p.Match(cayenneParserT__8)
		}
		{
			p.SetState(60)
			p.Varid()
		}
		{
			p.SetState(61)
			p.Match(cayenneParserT__9)
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == cayenneParserT__0 || _la == cayenneParserID {
			{
				p.SetState(62)
				p.Arm()
			}

			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(68)
			p.Match(cayenneParserT__1)
		}
		{
			p.SetState(69)
			p.Type_()
		}

	case 6:
		{
			p.SetState(71)
			p.Match(cayenneParserT__10)
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(72)
					p.Sign()
				}

			}
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}

	case 7:
		{
			p.SetState(78)
			p.Match(cayenneParserT__11)
		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(79)
					p.Defn()
				}

			}
			p.SetState(84)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}

	case 8:
		{
			p.SetState(85)
			p.Id_()
		}

	case 9:
		{
			p.SetState(86)
			p.Match(cayenneParserT__13)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cayenneParserRULE_expr)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(90)
					p.expr(10)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cayenneParserRULE_expr)
				p.SetState(91)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(92)
					p.Match(cayenneParserT__12)
				}
				{
					p.SetState(93)
					p.Lblid()
				}

			}

		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IArmContext is an interface to support dynamic dispatch.
type IArmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArmContext differentiates from other interfaces.
	IsArmContext()
}

type ArmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArmContext() *ArmContext {
	var p = new(ArmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cayenneParserRULE_arm
	return p
}

func (*ArmContext) IsArmContext() {}

func NewArmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArmContext {
	var p = new(ArmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_arm

	return p
}

func (s *ArmContext) GetParser() antlr.Parser { return s.parser }

func (s *ArmContext) Conid() IConidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConidContext)
}

func (s *ArmContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArmContext) AllVarid() []IVaridContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVaridContext)(nil)).Elem())
	var tst = make([]IVaridContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVaridContext)
		}
	}

	return tst
}

func (s *ArmContext) Varid(i int) IVaridContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVaridContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVaridContext)
}

func (s *ArmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterArm(s)
	}
}

func (s *ArmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitArm(s)
	}
}

func (p *cayenneParser) Arm() (localctx IArmContext) {
	this := p
	_ = this

	localctx = NewArmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, cayenneParserRULE_arm)
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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cayenneParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Match(cayenneParserT__0)
		}
		{
			p.SetState(100)
			p.Conid()
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == cayenneParserID {
			{
				p.SetState(101)
				p.Varid()
			}

			p.SetState(106)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(107)
			p.Match(cayenneParserT__2)
		}
		{
			p.SetState(108)
			p.Match(cayenneParserT__3)
		}
		{
			p.SetState(109)
			p.expr(0)
		}
		{
			p.SetState(110)
			p.Match(cayenneParserT__14)
		}

	case cayenneParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Varid()
		}
		{
			p.SetState(113)
			p.Match(cayenneParserT__3)
		}
		{
			p.SetState(114)
			p.expr(0)
		}
		{
			p.SetState(115)
			p.Match(cayenneParserT__14)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISignContext is an interface to support dynamic dispatch.
type ISignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignContext differentiates from other interfaces.
	IsSignContext()
}

type SignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignContext() *SignContext {
	var p = new(SignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cayenneParserRULE_sign
	return p
}

func (*SignContext) IsSignContext() {}

func NewSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignContext {
	var p = new(SignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_sign

	return p
}

func (s *SignContext) GetParser() antlr.Parser { return s.parser }

func (s *SignContext) Lblid() ILblidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILblidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILblidContext)
}

func (s *SignContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *SignContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterSign(s)
	}
}

func (s *SignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitSign(s)
	}
}

func (p *cayenneParser) Sign() (localctx ISignContext) {
	this := p
	_ = this

	localctx = NewSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, cayenneParserRULE_sign)

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

	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Lblid()
		}
		{
			p.SetState(120)
			p.Match(cayenneParserT__1)
		}
		{
			p.SetState(121)
			p.Type_()
		}
		{
			p.SetState(122)
			p.Match(cayenneParserT__14)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Lblid()
		}
		{
			p.SetState(125)
			p.Match(cayenneParserT__1)
		}
		{
			p.SetState(126)
			p.Type_()
		}
		{
			p.SetState(127)
			p.Match(cayenneParserT__15)
		}
		{
			p.SetState(128)
			p.expr(0)
		}
		{
			p.SetState(129)
			p.Match(cayenneParserT__14)
		}

	}

	return localctx
}

// IDefnContext is an interface to support dynamic dispatch.
type IDefnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefnContext differentiates from other interfaces.
	IsDefnContext()
}

type DefnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefnContext() *DefnContext {
	var p = new(DefnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cayenneParserRULE_defn
	return p
}

func (*DefnContext) IsDefnContext() {}

func NewDefnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefnContext {
	var p = new(DefnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_defn

	return p
}

func (s *DefnContext) GetParser() antlr.Parser { return s.parser }

func (s *DefnContext) Vis() IVisContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVisContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVisContext)
}

func (s *DefnContext) Lblid() ILblidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILblidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILblidContext)
}

func (s *DefnContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *DefnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DefnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterDefn(s)
	}
}

func (s *DefnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitDefn(s)
	}
}

func (p *cayenneParser) Defn() (localctx IDefnContext) {
	this := p
	_ = this

	localctx = NewDefnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, cayenneParserRULE_defn)

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
		p.SetState(133)
		p.Vis()
	}
	{
		p.SetState(134)
		p.Lblid()
	}
	{
		p.SetState(135)
		p.Match(cayenneParserT__1)
	}
	{
		p.SetState(136)
		p.Type_()
	}
	{
		p.SetState(137)
		p.Match(cayenneParserT__15)
	}
	{
		p.SetState(138)
		p.expr(0)
	}
	{
		p.SetState(139)
		p.Match(cayenneParserT__14)
	}

	return localctx
}

// IVisContext is an interface to support dynamic dispatch.
type IVisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVisContext differentiates from other interfaces.
	IsVisContext()
}

type VisContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVisContext() *VisContext {
	var p = new(VisContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cayenneParserRULE_vis
	return p
}

func (*VisContext) IsVisContext() {}

func NewVisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VisContext {
	var p = new(VisContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_vis

	return p
}

func (s *VisContext) GetParser() antlr.Parser { return s.parser }

func (s *VisContext) Abs_() IAbs_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAbs_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAbs_Context)
}

func (s *VisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterVis(s)
	}
}

func (s *VisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitVis(s)
	}
}

func (p *cayenneParser) Vis() (localctx IVisContext) {
	this := p
	_ = this

	localctx = NewVisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, cayenneParserRULE_vis)

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

	p.SetState(144)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cayenneParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(cayenneParserT__16)
		}

	case cayenneParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Match(cayenneParserT__17)
		}
		{
			p.SetState(143)
			p.Abs_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAbs_Context is an interface to support dynamic dispatch.
type IAbs_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbs_Context differentiates from other interfaces.
	IsAbs_Context()
}

type Abs_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbs_Context() *Abs_Context {
	var p = new(Abs_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cayenneParserRULE_abs_
	return p
}

func (*Abs_Context) IsAbs_Context() {}

func NewAbs_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Abs_Context {
	var p = new(Abs_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_abs_

	return p
}

func (s *Abs_Context) GetParser() antlr.Parser { return s.parser }
func (s *Abs_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Abs_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Abs_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterAbs_(s)
	}
}

func (s *Abs_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitAbs_(s)
	}
}

func (p *cayenneParser) Abs_() (localctx IAbs_Context) {
	this := p
	_ = this

	localctx = NewAbs_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, cayenneParserRULE_abs_)
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
		p.SetState(146)
		_la = p.GetTokenStream().LA(1)

		if !(_la == cayenneParserT__18 || _la == cayenneParserT__19) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cayenneParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *cayenneParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, cayenneParserRULE_type_)

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
		p.SetState(148)
		p.expr(0)
	}

	return localctx
}

// IVaridContext is an interface to support dynamic dispatch.
type IVaridContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVaridContext differentiates from other interfaces.
	IsVaridContext()
}

type VaridContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVaridContext() *VaridContext {
	var p = new(VaridContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cayenneParserRULE_varid
	return p
}

func (*VaridContext) IsVaridContext() {}

func NewVaridContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VaridContext {
	var p = new(VaridContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_varid

	return p
}

func (s *VaridContext) GetParser() antlr.Parser { return s.parser }

func (s *VaridContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *VaridContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VaridContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VaridContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterVarid(s)
	}
}

func (s *VaridContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitVarid(s)
	}
}

func (p *cayenneParser) Varid() (localctx IVaridContext) {
	this := p
	_ = this

	localctx = NewVaridContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, cayenneParserRULE_varid)

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
		p.SetState(150)
		p.Id_()
	}

	return localctx
}

// IConidContext is an interface to support dynamic dispatch.
type IConidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConidContext differentiates from other interfaces.
	IsConidContext()
}

type ConidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConidContext() *ConidContext {
	var p = new(ConidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cayenneParserRULE_conid
	return p
}

func (*ConidContext) IsConidContext() {}

func NewConidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConidContext {
	var p = new(ConidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_conid

	return p
}

func (s *ConidContext) GetParser() antlr.Parser { return s.parser }

func (s *ConidContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *ConidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterConid(s)
	}
}

func (s *ConidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitConid(s)
	}
}

func (p *cayenneParser) Conid() (localctx IConidContext) {
	this := p
	_ = this

	localctx = NewConidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, cayenneParserRULE_conid)

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
		p.SetState(152)
		p.Id_()
	}

	return localctx
}

// ILblidContext is an interface to support dynamic dispatch.
type ILblidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLblidContext differentiates from other interfaces.
	IsLblidContext()
}

type LblidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLblidContext() *LblidContext {
	var p = new(LblidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cayenneParserRULE_lblid
	return p
}

func (*LblidContext) IsLblidContext() {}

func NewLblidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LblidContext {
	var p = new(LblidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_lblid

	return p
}

func (s *LblidContext) GetParser() antlr.Parser { return s.parser }

func (s *LblidContext) Id_() IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *LblidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LblidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LblidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterLblid(s)
	}
}

func (s *LblidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitLblid(s)
	}
}

func (p *cayenneParser) Lblid() (localctx ILblidContext) {
	this := p
	_ = this

	localctx = NewLblidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, cayenneParserRULE_lblid)

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
		p.SetState(154)
		p.Id_()
	}

	return localctx
}

// IId_Context is an interface to support dynamic dispatch.
type IId_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsId_Context differentiates from other interfaces.
	IsId_Context()
}

type Id_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_Context() *Id_Context {
	var p = new(Id_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cayenneParserRULE_id_
	return p
}

func (*Id_Context) IsId_Context() {}

func NewId_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_Context {
	var p = new(Id_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cayenneParserRULE_id_

	return p
}

func (s *Id_Context) GetParser() antlr.Parser { return s.parser }

func (s *Id_Context) ID() antlr.TerminalNode {
	return s.GetToken(cayenneParserID, 0)
}

func (s *Id_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.EnterId_(s)
	}
}

func (s *Id_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cayenneListener); ok {
		listenerT.ExitId_(s)
	}
}

func (p *cayenneParser) Id_() (localctx IId_Context) {
	this := p
	_ = this

	localctx = NewId_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, cayenneParserRULE_id_)

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
		p.SetState(156)
		p.Match(cayenneParserID)
	}

	return localctx
}

func (p *cayenneParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *cayenneParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
