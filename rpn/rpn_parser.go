// Generated from rpn.g4 by ANTLR 4.7.

package rpn // rpn
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 40, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 7, 2, 17, 10, 2, 12, 2, 14, 2, 20, 11, 2, 3, 3, 3, 3, 5, 3, 24,
	10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 34, 10, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 3, 3, 2,
	5, 17, 2, 38, 2, 14, 3, 2, 2, 2, 4, 23, 3, 2, 2, 2, 6, 25, 3, 2, 2, 2,
	8, 33, 3, 2, 2, 2, 10, 35, 3, 2, 2, 2, 12, 37, 3, 2, 2, 2, 14, 18, 5, 8,
	5, 2, 15, 17, 5, 4, 3, 2, 16, 15, 3, 2, 2, 2, 17, 20, 3, 2, 2, 2, 18, 16,
	3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 3, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2,
	21, 24, 5, 8, 5, 2, 22, 24, 5, 6, 4, 2, 23, 21, 3, 2, 2, 2, 23, 22, 3,
	2, 2, 2, 24, 5, 3, 2, 2, 2, 25, 26, 9, 2, 2, 2, 26, 7, 3, 2, 2, 2, 27,
	28, 7, 6, 2, 2, 28, 34, 5, 8, 5, 2, 29, 30, 7, 7, 2, 2, 30, 34, 5, 8, 5,
	2, 31, 34, 5, 12, 7, 2, 32, 34, 5, 10, 6, 2, 33, 27, 3, 2, 2, 2, 33, 29,
	3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 32, 3, 2, 2, 2, 34, 9, 3, 2, 2, 2,
	35, 36, 7, 4, 2, 2, 36, 11, 3, 2, 2, 2, 37, 38, 7, 3, 2, 2, 38, 13, 3,
	2, 2, 2, 5, 18, 23, 33,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "'^'", "'+'", "'-'", "'*'", "'/'", "'cos'", "'sin'", "'tan'",
	"'acos'", "'asin'", "'atan'", "'ln'", "'log'", "'.'",
}
var symbolicNames = []string{
	"", "SCIENTIFIC_NUMBER", "VARIABLE", "POW", "PLUS", "MINUS", "TIMES", "DIV",
	"COS", "SIN", "TAN", "ACOS", "ASIN", "ATAN", "LN", "LOG", "POINT", "WS",
}

var ruleNames = []string{
	"expression", "term", "oper", "signedAtom", "variable", "scientific",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type rpnParser struct {
	*antlr.BaseParser
}

func NewrpnParser(input antlr.TokenStream) *rpnParser {
	this := new(rpnParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "rpn.g4"

	return this
}

// rpnParser tokens.
const (
	rpnParserEOF               = antlr.TokenEOF
	rpnParserSCIENTIFIC_NUMBER = 1
	rpnParserVARIABLE          = 2
	rpnParserPOW               = 3
	rpnParserPLUS              = 4
	rpnParserMINUS             = 5
	rpnParserTIMES             = 6
	rpnParserDIV               = 7
	rpnParserCOS               = 8
	rpnParserSIN               = 9
	rpnParserTAN               = 10
	rpnParserACOS              = 11
	rpnParserASIN              = 12
	rpnParserATAN              = 13
	rpnParserLN                = 14
	rpnParserLOG               = 15
	rpnParserPOINT             = 16
	rpnParserWS                = 17
)

// rpnParser rules.
const (
	rpnParserRULE_expression = 0
	rpnParserRULE_term       = 1
	rpnParserRULE_oper       = 2
	rpnParserRULE_signedAtom = 3
	rpnParserRULE_variable   = 4
	rpnParserRULE_scientific = 5
)

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
	p.RuleIndex = rpnParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rpnParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) SignedAtom() ISignedAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignedAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignedAtomContext)
}

func (s *ExpressionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case rpnVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *rpnParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, rpnParserRULE_expression)
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
		p.SetState(12)
		p.SignedAtom()
	}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rpnParserSCIENTIFIC_NUMBER)|(1<<rpnParserVARIABLE)|(1<<rpnParserPOW)|(1<<rpnParserPLUS)|(1<<rpnParserMINUS)|(1<<rpnParserTIMES)|(1<<rpnParserDIV)|(1<<rpnParserCOS)|(1<<rpnParserSIN)|(1<<rpnParserTAN)|(1<<rpnParserACOS)|(1<<rpnParserASIN)|(1<<rpnParserATAN)|(1<<rpnParserLN)|(1<<rpnParserLOG))) != 0 {
		{
			p.SetState(13)
			p.Term()
		}

		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = rpnParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rpnParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) SignedAtom() ISignedAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignedAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignedAtomContext)
}

func (s *TermContext) Oper() IOperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case rpnVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *rpnParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, rpnParserRULE_term)

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

	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.SignedAtom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(20)
			p.Oper()
		}

	}

	return localctx
}

// IOperContext is an interface to support dynamic dispatch.
type IOperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperContext differentiates from other interfaces.
	IsOperContext()
}

type OperContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperContext() *OperContext {
	var p = new(OperContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rpnParserRULE_oper
	return p
}

func (*OperContext) IsOperContext() {}

func NewOperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperContext {
	var p = new(OperContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rpnParserRULE_oper

	return p
}

func (s *OperContext) GetParser() antlr.Parser { return s.parser }

func (s *OperContext) POW() antlr.TerminalNode {
	return s.GetToken(rpnParserPOW, 0)
}

func (s *OperContext) PLUS() antlr.TerminalNode {
	return s.GetToken(rpnParserPLUS, 0)
}

func (s *OperContext) MINUS() antlr.TerminalNode {
	return s.GetToken(rpnParserMINUS, 0)
}

func (s *OperContext) TIMES() antlr.TerminalNode {
	return s.GetToken(rpnParserTIMES, 0)
}

func (s *OperContext) DIV() antlr.TerminalNode {
	return s.GetToken(rpnParserDIV, 0)
}

func (s *OperContext) COS() antlr.TerminalNode {
	return s.GetToken(rpnParserCOS, 0)
}

func (s *OperContext) TAN() antlr.TerminalNode {
	return s.GetToken(rpnParserTAN, 0)
}

func (s *OperContext) SIN() antlr.TerminalNode {
	return s.GetToken(rpnParserSIN, 0)
}

func (s *OperContext) ACOS() antlr.TerminalNode {
	return s.GetToken(rpnParserACOS, 0)
}

func (s *OperContext) ATAN() antlr.TerminalNode {
	return s.GetToken(rpnParserATAN, 0)
}

func (s *OperContext) ASIN() antlr.TerminalNode {
	return s.GetToken(rpnParserASIN, 0)
}

func (s *OperContext) LOG() antlr.TerminalNode {
	return s.GetToken(rpnParserLOG, 0)
}

func (s *OperContext) LN() antlr.TerminalNode {
	return s.GetToken(rpnParserLN, 0)
}

func (s *OperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.EnterOper(s)
	}
}

func (s *OperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.ExitOper(s)
	}
}

func (s *OperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case rpnVisitor:
		return t.VisitOper(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *rpnParser) Oper() (localctx IOperContext) {
	localctx = NewOperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, rpnParserRULE_oper)
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
	p.SetState(23)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rpnParserPOW)|(1<<rpnParserPLUS)|(1<<rpnParserMINUS)|(1<<rpnParserTIMES)|(1<<rpnParserDIV)|(1<<rpnParserCOS)|(1<<rpnParserSIN)|(1<<rpnParserTAN)|(1<<rpnParserACOS)|(1<<rpnParserASIN)|(1<<rpnParserATAN)|(1<<rpnParserLN)|(1<<rpnParserLOG))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ISignedAtomContext is an interface to support dynamic dispatch.
type ISignedAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignedAtomContext differentiates from other interfaces.
	IsSignedAtomContext()
}

type SignedAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignedAtomContext() *SignedAtomContext {
	var p = new(SignedAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rpnParserRULE_signedAtom
	return p
}

func (*SignedAtomContext) IsSignedAtomContext() {}

func NewSignedAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignedAtomContext {
	var p = new(SignedAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rpnParserRULE_signedAtom

	return p
}

func (s *SignedAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *SignedAtomContext) PLUS() antlr.TerminalNode {
	return s.GetToken(rpnParserPLUS, 0)
}

func (s *SignedAtomContext) SignedAtom() ISignedAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignedAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignedAtomContext)
}

func (s *SignedAtomContext) MINUS() antlr.TerminalNode {
	return s.GetToken(rpnParserMINUS, 0)
}

func (s *SignedAtomContext) Scientific() IScientificContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScientificContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScientificContext)
}

func (s *SignedAtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *SignedAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignedAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignedAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.EnterSignedAtom(s)
	}
}

func (s *SignedAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.ExitSignedAtom(s)
	}
}

func (s *SignedAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case rpnVisitor:
		return t.VisitSignedAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *rpnParser) SignedAtom() (localctx ISignedAtomContext) {
	localctx = NewSignedAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, rpnParserRULE_signedAtom)

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

	switch p.GetTokenStream().LA(1) {
	case rpnParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)
			p.Match(rpnParserPLUS)
		}
		{
			p.SetState(26)
			p.SignedAtom()
		}

	case rpnParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.Match(rpnParserMINUS)
		}
		{
			p.SetState(28)
			p.SignedAtom()
		}

	case rpnParserSCIENTIFIC_NUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(29)
			p.Scientific()
		}

	case rpnParserVARIABLE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(30)
			p.Variable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = rpnParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rpnParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(rpnParserVARIABLE, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case rpnVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *rpnParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, rpnParserRULE_variable)

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
		p.Match(rpnParserVARIABLE)
	}

	return localctx
}

// IScientificContext is an interface to support dynamic dispatch.
type IScientificContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScientificContext differentiates from other interfaces.
	IsScientificContext()
}

type ScientificContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScientificContext() *ScientificContext {
	var p = new(ScientificContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rpnParserRULE_scientific
	return p
}

func (*ScientificContext) IsScientificContext() {}

func NewScientificContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScientificContext {
	var p = new(ScientificContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rpnParserRULE_scientific

	return p
}

func (s *ScientificContext) GetParser() antlr.Parser { return s.parser }

func (s *ScientificContext) SCIENTIFIC_NUMBER() antlr.TerminalNode {
	return s.GetToken(rpnParserSCIENTIFIC_NUMBER, 0)
}

func (s *ScientificContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScientificContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScientificContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.EnterScientific(s)
	}
}

func (s *ScientificContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rpnListener); ok {
		listenerT.ExitScientific(s)
	}
}

func (s *ScientificContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case rpnVisitor:
		return t.VisitScientific(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *rpnParser) Scientific() (localctx IScientificContext) {
	localctx = NewScientificContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, rpnParserRULE_scientific)

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
		p.Match(rpnParserSCIENTIFIC_NUMBER)
	}

	return localctx
}
