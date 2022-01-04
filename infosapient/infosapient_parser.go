// Code generated from infosapient.g4 by ANTLR 4.9.3. DO NOT EDIT.

package infosapient // infosapient
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 120,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 39, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 47, 10,
	5, 3, 5, 3, 5, 3, 5, 5, 5, 52, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 57, 10, 5,
	7, 5, 59, 10, 5, 12, 5, 14, 5, 62, 11, 5, 5, 5, 64, 10, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 71, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 76, 10, 7, 3,
	7, 3, 7, 3, 7, 5, 7, 81, 10, 7, 3, 8, 6, 8, 84, 10, 8, 13, 8, 14, 8, 85,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 5, 10, 112, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	2, 7, 3, 2, 6, 7, 3, 2, 8, 9, 3, 2, 23, 24, 3, 2, 26, 27, 3, 2, 34, 46,
	2, 139, 2, 26, 3, 2, 2, 2, 4, 31, 3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8, 63,
	3, 2, 2, 2, 10, 70, 3, 2, 2, 2, 12, 72, 3, 2, 2, 2, 14, 83, 3, 2, 2, 2,
	16, 87, 3, 2, 2, 2, 18, 111, 3, 2, 2, 2, 20, 113, 3, 2, 2, 2, 22, 115,
	3, 2, 2, 2, 24, 117, 3, 2, 2, 2, 26, 27, 7, 47, 2, 2, 27, 28, 5, 4, 3,
	2, 28, 29, 5, 6, 4, 2, 29, 30, 7, 3, 2, 2, 30, 3, 3, 2, 2, 2, 31, 32, 7,
	49, 2, 2, 32, 33, 5, 8, 5, 2, 33, 5, 3, 2, 2, 2, 34, 35, 7, 50, 2, 2, 35,
	38, 5, 12, 7, 2, 36, 37, 7, 51, 2, 2, 37, 39, 5, 12, 7, 2, 38, 36, 3, 2,
	2, 2, 38, 39, 3, 2, 2, 2, 39, 7, 3, 2, 2, 2, 40, 41, 7, 4, 2, 2, 41, 42,
	5, 8, 5, 2, 42, 43, 7, 5, 2, 2, 43, 64, 3, 2, 2, 2, 44, 47, 5, 10, 6, 2,
	45, 47, 5, 12, 7, 2, 46, 44, 3, 2, 2, 2, 46, 45, 3, 2, 2, 2, 47, 48, 3,
	2, 2, 2, 48, 51, 5, 20, 11, 2, 49, 52, 5, 10, 6, 2, 50, 52, 5, 12, 7, 2,
	51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 60, 3, 2, 2, 2, 53, 56, 5,
	20, 11, 2, 54, 57, 5, 10, 6, 2, 55, 57, 5, 12, 7, 2, 56, 54, 3, 2, 2, 2,
	56, 55, 3, 2, 2, 2, 57, 59, 3, 2, 2, 2, 58, 53, 3, 2, 2, 2, 59, 62, 3,
	2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62,
	60, 3, 2, 2, 2, 63, 40, 3, 2, 2, 2, 63, 46, 3, 2, 2, 2, 64, 9, 3, 2, 2,
	2, 65, 66, 7, 4, 2, 2, 66, 67, 5, 10, 6, 2, 67, 68, 7, 5, 2, 2, 68, 71,
	3, 2, 2, 2, 69, 71, 5, 12, 7, 2, 70, 65, 3, 2, 2, 2, 70, 69, 3, 2, 2, 2,
	71, 11, 3, 2, 2, 2, 72, 73, 5, 24, 13, 2, 73, 75, 9, 2, 2, 2, 74, 76, 5,
	14, 8, 2, 75, 74, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 80, 3, 2, 2, 2, 77,
	81, 5, 22, 12, 2, 78, 81, 5, 24, 13, 2, 79, 81, 5, 16, 9, 2, 80, 77, 3,
	2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81, 13, 3, 2, 2, 2, 82,
	84, 5, 18, 10, 2, 83, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 83, 3, 2,
	2, 2, 85, 86, 3, 2, 2, 2, 86, 15, 3, 2, 2, 2, 87, 88, 9, 3, 2, 2, 88, 17,
	3, 2, 2, 2, 89, 112, 7, 10, 2, 2, 90, 112, 7, 11, 2, 2, 91, 112, 7, 12,
	2, 2, 92, 112, 7, 13, 2, 2, 93, 112, 7, 14, 2, 2, 94, 112, 7, 15, 2, 2,
	95, 112, 7, 16, 2, 2, 96, 112, 7, 17, 2, 2, 97, 112, 7, 18, 2, 2, 98, 112,
	7, 19, 2, 2, 99, 112, 7, 20, 2, 2, 100, 112, 7, 21, 2, 2, 101, 112, 7,
	22, 2, 2, 102, 112, 9, 4, 2, 2, 103, 112, 7, 25, 2, 2, 104, 112, 9, 5,
	2, 2, 105, 112, 7, 28, 2, 2, 106, 112, 7, 29, 2, 2, 107, 112, 7, 30, 2,
	2, 108, 112, 7, 31, 2, 2, 109, 112, 7, 32, 2, 2, 110, 112, 7, 33, 2, 2,
	111, 89, 3, 2, 2, 2, 111, 90, 3, 2, 2, 2, 111, 91, 3, 2, 2, 2, 111, 92,
	3, 2, 2, 2, 111, 93, 3, 2, 2, 2, 111, 94, 3, 2, 2, 2, 111, 95, 3, 2, 2,
	2, 111, 96, 3, 2, 2, 2, 111, 97, 3, 2, 2, 2, 111, 98, 3, 2, 2, 2, 111,
	99, 3, 2, 2, 2, 111, 100, 3, 2, 2, 2, 111, 101, 3, 2, 2, 2, 111, 102, 3,
	2, 2, 2, 111, 103, 3, 2, 2, 2, 111, 104, 3, 2, 2, 2, 111, 105, 3, 2, 2,
	2, 111, 106, 3, 2, 2, 2, 111, 107, 3, 2, 2, 2, 111, 108, 3, 2, 2, 2, 111,
	109, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2, 112, 19, 3, 2, 2, 2, 113, 114, 9,
	6, 2, 2, 114, 21, 3, 2, 2, 2, 115, 116, 7, 48, 2, 2, 116, 23, 3, 2, 2,
	2, 117, 118, 7, 47, 2, 2, 118, 25, 3, 2, 2, 2, 13, 38, 46, 51, 56, 60,
	63, 70, 75, 80, 85, 111,
}
var literalNames = []string{
	"", "';'", "'('", "')'", "'s'", "'are'", "'increased'", "'decreased'",
	"'about'", "'above'", "'after'", "'around'", "'before'", "'below'", "'closeTo'",
	"'definitely'", "'extremely'", "'generally'", "'mostly'", "'must'", "'near'",
	"'negative'", "'negatively'", "'not'", "'positive'", "'positively'", "'roughly'",
	"'should'", "'slightly'", "'somewhat'", "'very'", "'inVicinityOf'", "'and'",
	"'boundedAnd'", "'cosineNot'", "'meanAnd'", "'meanOr'", "'or'", "'productAnd'",
	"'productOr'", "'sugenoNot'", "'thresholdNot'", "'yagerAnd'", "'yagerNot'",
	"'yagerOr'", "", "", "'if'", "'then'", "'else'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "IDENTIFIER", "FP_LITERAL", "IF", "THEN",
	"ELSE", "WS",
}

var ruleNames = []string{
	"conditionalRule", "premise", "consequent", "clause", "expr", "attribClause",
	"hedgeCollection", "restrictionHedge", "hedge", "operator_", "nLiteral",
	"id_",
}

type infosapientParser struct {
	*antlr.BaseParser
}

// NewinfosapientParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *infosapientParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewinfosapientParser(input antlr.TokenStream) *infosapientParser {
	this := new(infosapientParser)
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
	this.GrammarFileName = "infosapient.g4"

	return this
}

// infosapientParser tokens.
const (
	infosapientParserEOF        = antlr.TokenEOF
	infosapientParserT__0       = 1
	infosapientParserT__1       = 2
	infosapientParserT__2       = 3
	infosapientParserT__3       = 4
	infosapientParserT__4       = 5
	infosapientParserT__5       = 6
	infosapientParserT__6       = 7
	infosapientParserT__7       = 8
	infosapientParserT__8       = 9
	infosapientParserT__9       = 10
	infosapientParserT__10      = 11
	infosapientParserT__11      = 12
	infosapientParserT__12      = 13
	infosapientParserT__13      = 14
	infosapientParserT__14      = 15
	infosapientParserT__15      = 16
	infosapientParserT__16      = 17
	infosapientParserT__17      = 18
	infosapientParserT__18      = 19
	infosapientParserT__19      = 20
	infosapientParserT__20      = 21
	infosapientParserT__21      = 22
	infosapientParserT__22      = 23
	infosapientParserT__23      = 24
	infosapientParserT__24      = 25
	infosapientParserT__25      = 26
	infosapientParserT__26      = 27
	infosapientParserT__27      = 28
	infosapientParserT__28      = 29
	infosapientParserT__29      = 30
	infosapientParserT__30      = 31
	infosapientParserT__31      = 32
	infosapientParserT__32      = 33
	infosapientParserT__33      = 34
	infosapientParserT__34      = 35
	infosapientParserT__35      = 36
	infosapientParserT__36      = 37
	infosapientParserT__37      = 38
	infosapientParserT__38      = 39
	infosapientParserT__39      = 40
	infosapientParserT__40      = 41
	infosapientParserT__41      = 42
	infosapientParserT__42      = 43
	infosapientParserT__43      = 44
	infosapientParserIDENTIFIER = 45
	infosapientParserFP_LITERAL = 46
	infosapientParserIF         = 47
	infosapientParserTHEN       = 48
	infosapientParserELSE       = 49
	infosapientParserWS         = 50
)

// infosapientParser rules.
const (
	infosapientParserRULE_conditionalRule  = 0
	infosapientParserRULE_premise          = 1
	infosapientParserRULE_consequent       = 2
	infosapientParserRULE_clause           = 3
	infosapientParserRULE_expr             = 4
	infosapientParserRULE_attribClause     = 5
	infosapientParserRULE_hedgeCollection  = 6
	infosapientParserRULE_restrictionHedge = 7
	infosapientParserRULE_hedge            = 8
	infosapientParserRULE_operator_        = 9
	infosapientParserRULE_nLiteral         = 10
	infosapientParserRULE_id_              = 11
)

// IConditionalRuleContext is an interface to support dynamic dispatch.
type IConditionalRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalRuleContext differentiates from other interfaces.
	IsConditionalRuleContext()
}

type ConditionalRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalRuleContext() *ConditionalRuleContext {
	var p = new(ConditionalRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = infosapientParserRULE_conditionalRule
	return p
}

func (*ConditionalRuleContext) IsConditionalRuleContext() {}

func NewConditionalRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalRuleContext {
	var p = new(ConditionalRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_conditionalRule

	return p
}

func (s *ConditionalRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(infosapientParserIDENTIFIER, 0)
}

func (s *ConditionalRuleContext) Premise() IPremiseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPremiseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPremiseContext)
}

func (s *ConditionalRuleContext) Consequent() IConsequentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConsequentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConsequentContext)
}

func (s *ConditionalRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterConditionalRule(s)
	}
}

func (s *ConditionalRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitConditionalRule(s)
	}
}

func (p *infosapientParser) ConditionalRule() (localctx IConditionalRuleContext) {
	this := p
	_ = this

	localctx = NewConditionalRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, infosapientParserRULE_conditionalRule)

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
		p.SetState(24)
		p.Match(infosapientParserIDENTIFIER)
	}
	{
		p.SetState(25)
		p.Premise()
	}
	{
		p.SetState(26)
		p.Consequent()
	}
	{
		p.SetState(27)
		p.Match(infosapientParserT__0)
	}

	return localctx
}

// IPremiseContext is an interface to support dynamic dispatch.
type IPremiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPremiseContext differentiates from other interfaces.
	IsPremiseContext()
}

type PremiseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPremiseContext() *PremiseContext {
	var p = new(PremiseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = infosapientParserRULE_premise
	return p
}

func (*PremiseContext) IsPremiseContext() {}

func NewPremiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PremiseContext {
	var p = new(PremiseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_premise

	return p
}

func (s *PremiseContext) GetParser() antlr.Parser { return s.parser }

func (s *PremiseContext) IF() antlr.TerminalNode {
	return s.GetToken(infosapientParserIF, 0)
}

func (s *PremiseContext) Clause() IClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClauseContext)
}

func (s *PremiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PremiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PremiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterPremise(s)
	}
}

func (s *PremiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitPremise(s)
	}
}

func (p *infosapientParser) Premise() (localctx IPremiseContext) {
	this := p
	_ = this

	localctx = NewPremiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, infosapientParserRULE_premise)

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
		p.SetState(29)
		p.Match(infosapientParserIF)
	}
	{
		p.SetState(30)
		p.Clause()
	}

	return localctx
}

// IConsequentContext is an interface to support dynamic dispatch.
type IConsequentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConsequentContext differentiates from other interfaces.
	IsConsequentContext()
}

type ConsequentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConsequentContext() *ConsequentContext {
	var p = new(ConsequentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = infosapientParserRULE_consequent
	return p
}

func (*ConsequentContext) IsConsequentContext() {}

func NewConsequentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConsequentContext {
	var p = new(ConsequentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_consequent

	return p
}

func (s *ConsequentContext) GetParser() antlr.Parser { return s.parser }

func (s *ConsequentContext) THEN() antlr.TerminalNode {
	return s.GetToken(infosapientParserTHEN, 0)
}

func (s *ConsequentContext) AllAttribClause() []IAttribClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttribClauseContext)(nil)).Elem())
	var tst = make([]IAttribClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttribClauseContext)
		}
	}

	return tst
}

func (s *ConsequentContext) AttribClause(i int) IAttribClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttribClauseContext)
}

func (s *ConsequentContext) ELSE() antlr.TerminalNode {
	return s.GetToken(infosapientParserELSE, 0)
}

func (s *ConsequentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsequentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConsequentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterConsequent(s)
	}
}

func (s *ConsequentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitConsequent(s)
	}
}

func (p *infosapientParser) Consequent() (localctx IConsequentContext) {
	this := p
	_ = this

	localctx = NewConsequentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, infosapientParserRULE_consequent)
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
		p.SetState(32)
		p.Match(infosapientParserTHEN)
	}
	{
		p.SetState(33)
		p.AttribClause()
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == infosapientParserELSE {
		{
			p.SetState(34)
			p.Match(infosapientParserELSE)
		}
		{
			p.SetState(35)
			p.AttribClause()
		}

	}

	return localctx
}

// IClauseContext is an interface to support dynamic dispatch.
type IClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseContext differentiates from other interfaces.
	IsClauseContext()
}

type ClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = infosapientParserRULE_clause
	return p
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_clause

	return p
}

func (s *ClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseContext) Clause() IClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClauseContext)
}

func (s *ClauseContext) AllOperator_() []IOperator_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperator_Context)(nil)).Elem())
	var tst = make([]IOperator_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperator_Context)
		}
	}

	return tst
}

func (s *ClauseContext) Operator_(i int) IOperator_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperator_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperator_Context)
}

func (s *ClauseContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ClauseContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ClauseContext) AllAttribClause() []IAttribClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttribClauseContext)(nil)).Elem())
	var tst = make([]IAttribClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttribClauseContext)
		}
	}

	return tst
}

func (s *ClauseContext) AttribClause(i int) IAttribClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttribClauseContext)
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitClause(s)
	}
}

func (p *infosapientParser) Clause() (localctx IClauseContext) {
	this := p
	_ = this

	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, infosapientParserRULE_clause)
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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(infosapientParserT__1)
		}
		{
			p.SetState(39)
			p.Clause()
		}
		{
			p.SetState(40)
			p.Match(infosapientParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(42)
				p.Expr()
			}

		case 2:
			{
				p.SetState(43)
				p.AttribClause()
			}

		}
		{
			p.SetState(46)
			p.Operator_()
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(47)
				p.Expr()
			}

		case 2:
			{
				p.SetState(48)
				p.AttribClause()
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(infosapientParserT__31-32))|(1<<(infosapientParserT__32-32))|(1<<(infosapientParserT__33-32))|(1<<(infosapientParserT__34-32))|(1<<(infosapientParserT__35-32))|(1<<(infosapientParserT__36-32))|(1<<(infosapientParserT__37-32))|(1<<(infosapientParserT__38-32))|(1<<(infosapientParserT__39-32))|(1<<(infosapientParserT__40-32))|(1<<(infosapientParserT__41-32))|(1<<(infosapientParserT__42-32))|(1<<(infosapientParserT__43-32)))) != 0 {
			{
				p.SetState(51)
				p.Operator_()
			}
			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(52)
					p.Expr()
				}

			case 2:
				{
					p.SetState(53)
					p.AttribClause()
				}

			}

			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

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
	p.RuleIndex = infosapientParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) AttribClause() IAttribClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribClauseContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *infosapientParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, infosapientParserRULE_expr)

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

	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case infosapientParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Match(infosapientParserT__1)
		}
		{
			p.SetState(64)
			p.Expr()
		}
		{
			p.SetState(65)
			p.Match(infosapientParserT__2)
		}

	case infosapientParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.AttribClause()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAttribClauseContext is an interface to support dynamic dispatch.
type IAttribClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribClauseContext differentiates from other interfaces.
	IsAttribClauseContext()
}

type AttribClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribClauseContext() *AttribClauseContext {
	var p = new(AttribClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = infosapientParserRULE_attribClause
	return p
}

func (*AttribClauseContext) IsAttribClauseContext() {}

func NewAttribClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribClauseContext {
	var p = new(AttribClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_attribClause

	return p
}

func (s *AttribClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribClauseContext) AllId_() []IId_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IId_Context)(nil)).Elem())
	var tst = make([]IId_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IId_Context)
		}
	}

	return tst
}

func (s *AttribClauseContext) Id_(i int) IId_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IId_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IId_Context)
}

func (s *AttribClauseContext) NLiteral() INLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INLiteralContext)
}

func (s *AttribClauseContext) RestrictionHedge() IRestrictionHedgeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRestrictionHedgeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRestrictionHedgeContext)
}

func (s *AttribClauseContext) HedgeCollection() IHedgeCollectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHedgeCollectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHedgeCollectionContext)
}

func (s *AttribClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterAttribClause(s)
	}
}

func (s *AttribClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitAttribClause(s)
	}
}

func (p *infosapientParser) AttribClause() (localctx IAttribClauseContext) {
	this := p
	_ = this

	localctx = NewAttribClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, infosapientParserRULE_attribClause)
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
		p.SetState(70)
		p.Id_()
	}
	{
		p.SetState(71)
		_la = p.GetTokenStream().LA(1)

		if !(_la == infosapientParserT__3 || _la == infosapientParserT__4) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<infosapientParserT__7)|(1<<infosapientParserT__8)|(1<<infosapientParserT__9)|(1<<infosapientParserT__10)|(1<<infosapientParserT__11)|(1<<infosapientParserT__12)|(1<<infosapientParserT__13)|(1<<infosapientParserT__14)|(1<<infosapientParserT__15)|(1<<infosapientParserT__16)|(1<<infosapientParserT__17)|(1<<infosapientParserT__18)|(1<<infosapientParserT__19)|(1<<infosapientParserT__20)|(1<<infosapientParserT__21)|(1<<infosapientParserT__22)|(1<<infosapientParserT__23)|(1<<infosapientParserT__24)|(1<<infosapientParserT__25)|(1<<infosapientParserT__26)|(1<<infosapientParserT__27)|(1<<infosapientParserT__28)|(1<<infosapientParserT__29)|(1<<infosapientParserT__30))) != 0 {
		{
			p.SetState(72)
			p.HedgeCollection()
		}

	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case infosapientParserFP_LITERAL:
		{
			p.SetState(75)
			p.NLiteral()
		}

	case infosapientParserIDENTIFIER:
		{
			p.SetState(76)
			p.Id_()
		}

	case infosapientParserT__5, infosapientParserT__6:
		{
			p.SetState(77)
			p.RestrictionHedge()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IHedgeCollectionContext is an interface to support dynamic dispatch.
type IHedgeCollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHedgeCollectionContext differentiates from other interfaces.
	IsHedgeCollectionContext()
}

type HedgeCollectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHedgeCollectionContext() *HedgeCollectionContext {
	var p = new(HedgeCollectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = infosapientParserRULE_hedgeCollection
	return p
}

func (*HedgeCollectionContext) IsHedgeCollectionContext() {}

func NewHedgeCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HedgeCollectionContext {
	var p = new(HedgeCollectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_hedgeCollection

	return p
}

func (s *HedgeCollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *HedgeCollectionContext) AllHedge() []IHedgeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHedgeContext)(nil)).Elem())
	var tst = make([]IHedgeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHedgeContext)
		}
	}

	return tst
}

func (s *HedgeCollectionContext) Hedge(i int) IHedgeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHedgeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHedgeContext)
}

func (s *HedgeCollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HedgeCollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HedgeCollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterHedgeCollection(s)
	}
}

func (s *HedgeCollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitHedgeCollection(s)
	}
}

func (p *infosapientParser) HedgeCollection() (localctx IHedgeCollectionContext) {
	this := p
	_ = this

	localctx = NewHedgeCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, infosapientParserRULE_hedgeCollection)
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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<infosapientParserT__7)|(1<<infosapientParserT__8)|(1<<infosapientParserT__9)|(1<<infosapientParserT__10)|(1<<infosapientParserT__11)|(1<<infosapientParserT__12)|(1<<infosapientParserT__13)|(1<<infosapientParserT__14)|(1<<infosapientParserT__15)|(1<<infosapientParserT__16)|(1<<infosapientParserT__17)|(1<<infosapientParserT__18)|(1<<infosapientParserT__19)|(1<<infosapientParserT__20)|(1<<infosapientParserT__21)|(1<<infosapientParserT__22)|(1<<infosapientParserT__23)|(1<<infosapientParserT__24)|(1<<infosapientParserT__25)|(1<<infosapientParserT__26)|(1<<infosapientParserT__27)|(1<<infosapientParserT__28)|(1<<infosapientParserT__29)|(1<<infosapientParserT__30))) != 0) {
		{
			p.SetState(80)
			p.Hedge()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRestrictionHedgeContext is an interface to support dynamic dispatch.
type IRestrictionHedgeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRestrictionHedgeContext differentiates from other interfaces.
	IsRestrictionHedgeContext()
}

type RestrictionHedgeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRestrictionHedgeContext() *RestrictionHedgeContext {
	var p = new(RestrictionHedgeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = infosapientParserRULE_restrictionHedge
	return p
}

func (*RestrictionHedgeContext) IsRestrictionHedgeContext() {}

func NewRestrictionHedgeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RestrictionHedgeContext {
	var p = new(RestrictionHedgeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_restrictionHedge

	return p
}

func (s *RestrictionHedgeContext) GetParser() antlr.Parser { return s.parser }
func (s *RestrictionHedgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RestrictionHedgeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RestrictionHedgeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterRestrictionHedge(s)
	}
}

func (s *RestrictionHedgeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitRestrictionHedge(s)
	}
}

func (p *infosapientParser) RestrictionHedge() (localctx IRestrictionHedgeContext) {
	this := p
	_ = this

	localctx = NewRestrictionHedgeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, infosapientParserRULE_restrictionHedge)
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
		p.SetState(85)
		_la = p.GetTokenStream().LA(1)

		if !(_la == infosapientParserT__5 || _la == infosapientParserT__6) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IHedgeContext is an interface to support dynamic dispatch.
type IHedgeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHedgeContext differentiates from other interfaces.
	IsHedgeContext()
}

type HedgeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHedgeContext() *HedgeContext {
	var p = new(HedgeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = infosapientParserRULE_hedge
	return p
}

func (*HedgeContext) IsHedgeContext() {}

func NewHedgeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HedgeContext {
	var p = new(HedgeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_hedge

	return p
}

func (s *HedgeContext) GetParser() antlr.Parser { return s.parser }
func (s *HedgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HedgeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HedgeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterHedge(s)
	}
}

func (s *HedgeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitHedge(s)
	}
}

func (p *infosapientParser) Hedge() (localctx IHedgeContext) {
	this := p
	_ = this

	localctx = NewHedgeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, infosapientParserRULE_hedge)
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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case infosapientParserT__7:
		{
			p.SetState(87)
			p.Match(infosapientParserT__7)
		}

	case infosapientParserT__8:
		{
			p.SetState(88)
			p.Match(infosapientParserT__8)
		}

	case infosapientParserT__9:
		{
			p.SetState(89)
			p.Match(infosapientParserT__9)
		}

	case infosapientParserT__10:
		{
			p.SetState(90)
			p.Match(infosapientParserT__10)
		}

	case infosapientParserT__11:
		{
			p.SetState(91)
			p.Match(infosapientParserT__11)
		}

	case infosapientParserT__12:
		{
			p.SetState(92)
			p.Match(infosapientParserT__12)
		}

	case infosapientParserT__13:
		{
			p.SetState(93)
			p.Match(infosapientParserT__13)
		}

	case infosapientParserT__14:
		{
			p.SetState(94)
			p.Match(infosapientParserT__14)
		}

	case infosapientParserT__15:
		{
			p.SetState(95)
			p.Match(infosapientParserT__15)
		}

	case infosapientParserT__16:
		{
			p.SetState(96)
			p.Match(infosapientParserT__16)
		}

	case infosapientParserT__17:
		{
			p.SetState(97)
			p.Match(infosapientParserT__17)
		}

	case infosapientParserT__18:
		{
			p.SetState(98)
			p.Match(infosapientParserT__18)
		}

	case infosapientParserT__19:
		{
			p.SetState(99)
			p.Match(infosapientParserT__19)
		}

	case infosapientParserT__20, infosapientParserT__21:
		{
			p.SetState(100)
			_la = p.GetTokenStream().LA(1)

			if !(_la == infosapientParserT__20 || _la == infosapientParserT__21) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case infosapientParserT__22:
		{
			p.SetState(101)
			p.Match(infosapientParserT__22)
		}

	case infosapientParserT__23, infosapientParserT__24:
		{
			p.SetState(102)
			_la = p.GetTokenStream().LA(1)

			if !(_la == infosapientParserT__23 || _la == infosapientParserT__24) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case infosapientParserT__25:
		{
			p.SetState(103)
			p.Match(infosapientParserT__25)
		}

	case infosapientParserT__26:
		{
			p.SetState(104)
			p.Match(infosapientParserT__26)
		}

	case infosapientParserT__27:
		{
			p.SetState(105)
			p.Match(infosapientParserT__27)
		}

	case infosapientParserT__28:
		{
			p.SetState(106)
			p.Match(infosapientParserT__28)
		}

	case infosapientParserT__29:
		{
			p.SetState(107)
			p.Match(infosapientParserT__29)
		}

	case infosapientParserT__30:
		{
			p.SetState(108)
			p.Match(infosapientParserT__30)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperator_Context is an interface to support dynamic dispatch.
type IOperator_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperator_Context differentiates from other interfaces.
	IsOperator_Context()
}

type Operator_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperator_Context() *Operator_Context {
	var p = new(Operator_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = infosapientParserRULE_operator_
	return p
}

func (*Operator_Context) IsOperator_Context() {}

func NewOperator_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operator_Context {
	var p = new(Operator_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_operator_

	return p
}

func (s *Operator_Context) GetParser() antlr.Parser { return s.parser }
func (s *Operator_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operator_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operator_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterOperator_(s)
	}
}

func (s *Operator_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitOperator_(s)
	}
}

func (p *infosapientParser) Operator_() (localctx IOperator_Context) {
	this := p
	_ = this

	localctx = NewOperator_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, infosapientParserRULE_operator_)
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
		p.SetState(111)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(infosapientParserT__31-32))|(1<<(infosapientParserT__32-32))|(1<<(infosapientParserT__33-32))|(1<<(infosapientParserT__34-32))|(1<<(infosapientParserT__35-32))|(1<<(infosapientParserT__36-32))|(1<<(infosapientParserT__37-32))|(1<<(infosapientParserT__38-32))|(1<<(infosapientParserT__39-32))|(1<<(infosapientParserT__40-32))|(1<<(infosapientParserT__41-32))|(1<<(infosapientParserT__42-32))|(1<<(infosapientParserT__43-32)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INLiteralContext is an interface to support dynamic dispatch.
type INLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNLiteralContext differentiates from other interfaces.
	IsNLiteralContext()
}

type NLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNLiteralContext() *NLiteralContext {
	var p = new(NLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = infosapientParserRULE_nLiteral
	return p
}

func (*NLiteralContext) IsNLiteralContext() {}

func NewNLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NLiteralContext {
	var p = new(NLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_nLiteral

	return p
}

func (s *NLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NLiteralContext) FP_LITERAL() antlr.TerminalNode {
	return s.GetToken(infosapientParserFP_LITERAL, 0)
}

func (s *NLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterNLiteral(s)
	}
}

func (s *NLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitNLiteral(s)
	}
}

func (p *infosapientParser) NLiteral() (localctx INLiteralContext) {
	this := p
	_ = this

	localctx = NewNLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, infosapientParserRULE_nLiteral)

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
		p.SetState(113)
		p.Match(infosapientParserFP_LITERAL)
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
	p.RuleIndex = infosapientParserRULE_id_
	return p
}

func (*Id_Context) IsId_Context() {}

func NewId_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_Context {
	var p = new(Id_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = infosapientParserRULE_id_

	return p
}

func (s *Id_Context) GetParser() antlr.Parser { return s.parser }

func (s *Id_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(infosapientParserIDENTIFIER, 0)
}

func (s *Id_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.EnterId_(s)
	}
}

func (s *Id_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(infosapientListener); ok {
		listenerT.ExitId_(s)
	}
}

func (p *infosapientParser) Id_() (localctx IId_Context) {
	this := p
	_ = this

	localctx = NewId_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, infosapientParserRULE_id_)

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
		p.SetState(115)
		p.Match(infosapientParserIDENTIFIER)
	}

	return localctx
}
