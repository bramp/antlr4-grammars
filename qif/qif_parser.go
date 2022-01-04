// Code generated from qifParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package qif // qifParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 86, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 45, 10, 3,
	12, 3, 14, 3, 48, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 76, 10, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 2, 2, 15, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 2, 2, 82, 2, 31, 3, 2, 2, 2,
	4, 46, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 54, 3, 2, 2, 2, 10, 57, 3, 2,
	2, 2, 12, 60, 3, 2, 2, 2, 14, 63, 3, 2, 2, 2, 16, 66, 3, 2, 2, 2, 18, 69,
	3, 2, 2, 2, 20, 72, 3, 2, 2, 2, 22, 77, 3, 2, 2, 2, 24, 81, 3, 2, 2, 2,
	26, 83, 3, 2, 2, 2, 28, 30, 5, 4, 3, 2, 29, 28, 3, 2, 2, 2, 30, 33, 3,
	2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2, 2, 33,
	31, 3, 2, 2, 2, 34, 35, 7, 2, 2, 3, 35, 3, 3, 2, 2, 2, 36, 45, 5, 6, 4,
	2, 37, 45, 5, 8, 5, 2, 38, 45, 5, 10, 6, 2, 39, 45, 5, 12, 7, 2, 40, 45,
	5, 14, 8, 2, 41, 45, 5, 16, 9, 2, 42, 45, 5, 18, 10, 2, 43, 45, 5, 20,
	11, 2, 44, 36, 3, 2, 2, 2, 44, 37, 3, 2, 2, 2, 44, 38, 3, 2, 2, 2, 44,
	39, 3, 2, 2, 2, 44, 40, 3, 2, 2, 2, 44, 41, 3, 2, 2, 2, 44, 42, 3, 2, 2,
	2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47,
	3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 5, 26, 14,
	2, 50, 5, 3, 2, 2, 2, 51, 52, 7, 3, 2, 2, 52, 53, 7, 16, 2, 2, 53, 7, 3,
	2, 2, 2, 54, 55, 7, 10, 2, 2, 55, 56, 7, 14, 2, 2, 56, 9, 3, 2, 2, 2, 57,
	58, 7, 4, 2, 2, 58, 59, 7, 15, 2, 2, 59, 11, 3, 2, 2, 2, 60, 61, 7, 6,
	2, 2, 61, 62, 7, 15, 2, 2, 62, 13, 3, 2, 2, 2, 63, 64, 7, 5, 2, 2, 64,
	65, 7, 11, 2, 2, 65, 15, 3, 2, 2, 2, 66, 67, 7, 7, 2, 2, 67, 68, 7, 16,
	2, 2, 68, 17, 3, 2, 2, 2, 69, 70, 7, 8, 2, 2, 70, 71, 7, 16, 2, 2, 71,
	19, 3, 2, 2, 2, 72, 75, 7, 9, 2, 2, 73, 76, 5, 22, 12, 2, 74, 76, 5, 24,
	13, 2, 75, 73, 3, 2, 2, 2, 75, 74, 3, 2, 2, 2, 76, 21, 3, 2, 2, 2, 77,
	78, 7, 18, 2, 2, 78, 79, 7, 19, 2, 2, 79, 80, 7, 20, 2, 2, 80, 23, 3, 2,
	2, 2, 81, 82, 7, 19, 2, 2, 82, 25, 3, 2, 2, 2, 83, 84, 7, 12, 2, 2, 84,
	27, 3, 2, 2, 2, 6, 31, 44, 46, 75,
}
var literalNames = []string{
	"", "'!Type:'", "'T'", "'C'", "'N'", "'M'", "'P'", "'L'", "'D'", "", "'^'",
	"", "", "", "", "", "'['", "", "']'",
}
var symbolicNames = []string{
	"", "TYPE", "T", "C", "N", "M", "P", "L", "D", "STATE", "EOR", "WS", "DATE",
	"NUM", "TEXT", "EOL", "LB", "ACCNTCATNAME", "RB", "EOL2",
}

var ruleNames = []string{
	"qif", "record", "recordtype", "date", "total", "check", "state_", "memo",
	"payee", "accountorcategory", "account", "category", "eor",
}

type qifParser struct {
	*antlr.BaseParser
}

// NewqifParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *qifParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewqifParser(input antlr.TokenStream) *qifParser {
	this := new(qifParser)
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
	this.GrammarFileName = "qifParser.g4"

	return this
}

// qifParser tokens.
const (
	qifParserEOF          = antlr.TokenEOF
	qifParserTYPE         = 1
	qifParserT            = 2
	qifParserC            = 3
	qifParserN            = 4
	qifParserM            = 5
	qifParserP            = 6
	qifParserL            = 7
	qifParserD            = 8
	qifParserSTATE        = 9
	qifParserEOR          = 10
	qifParserWS           = 11
	qifParserDATE         = 12
	qifParserNUM          = 13
	qifParserTEXT         = 14
	qifParserEOL          = 15
	qifParserLB           = 16
	qifParserACCNTCATNAME = 17
	qifParserRB           = 18
	qifParserEOL2         = 19
)

// qifParser rules.
const (
	qifParserRULE_qif               = 0
	qifParserRULE_record            = 1
	qifParserRULE_recordtype        = 2
	qifParserRULE_date              = 3
	qifParserRULE_total             = 4
	qifParserRULE_check             = 5
	qifParserRULE_state_            = 6
	qifParserRULE_memo              = 7
	qifParserRULE_payee             = 8
	qifParserRULE_accountorcategory = 9
	qifParserRULE_account           = 10
	qifParserRULE_category          = 11
	qifParserRULE_eor               = 12
)

// IQifContext is an interface to support dynamic dispatch.
type IQifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQifContext differentiates from other interfaces.
	IsQifContext()
}

type QifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQifContext() *QifContext {
	var p = new(QifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_qif
	return p
}

func (*QifContext) IsQifContext() {}

func NewQifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QifContext {
	var p = new(QifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_qif

	return p
}

func (s *QifContext) GetParser() antlr.Parser { return s.parser }

func (s *QifContext) EOF() antlr.TerminalNode {
	return s.GetToken(qifParserEOF, 0)
}

func (s *QifContext) AllRecord() []IRecordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRecordContext)(nil)).Elem())
	var tst = make([]IRecordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRecordContext)
		}
	}

	return tst
}

func (s *QifContext) Record(i int) IRecordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRecordContext)
}

func (s *QifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterQif(s)
	}
}

func (s *QifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitQif(s)
	}
}

func (p *qifParser) Qif() (localctx IQifContext) {
	this := p
	_ = this

	localctx = NewQifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, qifParserRULE_qif)
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

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qifParserTYPE)|(1<<qifParserT)|(1<<qifParserC)|(1<<qifParserN)|(1<<qifParserM)|(1<<qifParserP)|(1<<qifParserL)|(1<<qifParserD)|(1<<qifParserEOR))) != 0 {
		{
			p.SetState(26)
			p.Record()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(qifParserEOF)
	}

	return localctx
}

// IRecordContext is an interface to support dynamic dispatch.
type IRecordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecordContext differentiates from other interfaces.
	IsRecordContext()
}

type RecordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordContext() *RecordContext {
	var p = new(RecordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_record
	return p
}

func (*RecordContext) IsRecordContext() {}

func NewRecordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordContext {
	var p = new(RecordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_record

	return p
}

func (s *RecordContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordContext) Eor() IEorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEorContext)
}

func (s *RecordContext) AllRecordtype() []IRecordtypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRecordtypeContext)(nil)).Elem())
	var tst = make([]IRecordtypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRecordtypeContext)
		}
	}

	return tst
}

func (s *RecordContext) Recordtype(i int) IRecordtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecordtypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRecordtypeContext)
}

func (s *RecordContext) AllDate() []IDateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateContext)(nil)).Elem())
	var tst = make([]IDateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateContext)
		}
	}

	return tst
}

func (s *RecordContext) Date(i int) IDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *RecordContext) AllTotal() []ITotalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITotalContext)(nil)).Elem())
	var tst = make([]ITotalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITotalContext)
		}
	}

	return tst
}

func (s *RecordContext) Total(i int) ITotalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITotalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITotalContext)
}

func (s *RecordContext) AllCheck() []ICheckContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICheckContext)(nil)).Elem())
	var tst = make([]ICheckContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICheckContext)
		}
	}

	return tst
}

func (s *RecordContext) Check(i int) ICheckContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheckContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICheckContext)
}

func (s *RecordContext) AllState_() []IState_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IState_Context)(nil)).Elem())
	var tst = make([]IState_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IState_Context)
		}
	}

	return tst
}

func (s *RecordContext) State_(i int) IState_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IState_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IState_Context)
}

func (s *RecordContext) AllMemo() []IMemoContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemoContext)(nil)).Elem())
	var tst = make([]IMemoContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemoContext)
		}
	}

	return tst
}

func (s *RecordContext) Memo(i int) IMemoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemoContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemoContext)
}

func (s *RecordContext) AllPayee() []IPayeeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPayeeContext)(nil)).Elem())
	var tst = make([]IPayeeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPayeeContext)
		}
	}

	return tst
}

func (s *RecordContext) Payee(i int) IPayeeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPayeeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPayeeContext)
}

func (s *RecordContext) AllAccountorcategory() []IAccountorcategoryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAccountorcategoryContext)(nil)).Elem())
	var tst = make([]IAccountorcategoryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAccountorcategoryContext)
		}
	}

	return tst
}

func (s *RecordContext) Accountorcategory(i int) IAccountorcategoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountorcategoryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAccountorcategoryContext)
}

func (s *RecordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterRecord(s)
	}
}

func (s *RecordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitRecord(s)
	}
}

func (p *qifParser) Record() (localctx IRecordContext) {
	this := p
	_ = this

	localctx = NewRecordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, qifParserRULE_record)
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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<qifParserTYPE)|(1<<qifParserT)|(1<<qifParserC)|(1<<qifParserN)|(1<<qifParserM)|(1<<qifParserP)|(1<<qifParserL)|(1<<qifParserD))) != 0 {
		p.SetState(42)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case qifParserTYPE:
			{
				p.SetState(34)
				p.Recordtype()
			}

		case qifParserD:
			{
				p.SetState(35)
				p.Date()
			}

		case qifParserT:
			{
				p.SetState(36)
				p.Total()
			}

		case qifParserN:
			{
				p.SetState(37)
				p.Check()
			}

		case qifParserC:
			{
				p.SetState(38)
				p.State_()
			}

		case qifParserM:
			{
				p.SetState(39)
				p.Memo()
			}

		case qifParserP:
			{
				p.SetState(40)
				p.Payee()
			}

		case qifParserL:
			{
				p.SetState(41)
				p.Accountorcategory()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Eor()
	}

	return localctx
}

// IRecordtypeContext is an interface to support dynamic dispatch.
type IRecordtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecordtypeContext differentiates from other interfaces.
	IsRecordtypeContext()
}

type RecordtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordtypeContext() *RecordtypeContext {
	var p = new(RecordtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_recordtype
	return p
}

func (*RecordtypeContext) IsRecordtypeContext() {}

func NewRecordtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordtypeContext {
	var p = new(RecordtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_recordtype

	return p
}

func (s *RecordtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordtypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(qifParserTYPE, 0)
}

func (s *RecordtypeContext) TEXT() antlr.TerminalNode {
	return s.GetToken(qifParserTEXT, 0)
}

func (s *RecordtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterRecordtype(s)
	}
}

func (s *RecordtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitRecordtype(s)
	}
}

func (p *qifParser) Recordtype() (localctx IRecordtypeContext) {
	this := p
	_ = this

	localctx = NewRecordtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, qifParserRULE_recordtype)

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
		p.SetState(49)
		p.Match(qifParserTYPE)
	}
	{
		p.SetState(50)
		p.Match(qifParserTEXT)
	}

	return localctx
}

// IDateContext is an interface to support dynamic dispatch.
type IDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateContext differentiates from other interfaces.
	IsDateContext()
}

type DateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateContext() *DateContext {
	var p = new(DateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_date
	return p
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) D() antlr.TerminalNode {
	return s.GetToken(qifParserD, 0)
}

func (s *DateContext) DATE() antlr.TerminalNode {
	return s.GetToken(qifParserDATE, 0)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitDate(s)
	}
}

func (p *qifParser) Date() (localctx IDateContext) {
	this := p
	_ = this

	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, qifParserRULE_date)

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
		p.SetState(52)
		p.Match(qifParserD)
	}
	{
		p.SetState(53)
		p.Match(qifParserDATE)
	}

	return localctx
}

// ITotalContext is an interface to support dynamic dispatch.
type ITotalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTotalContext differentiates from other interfaces.
	IsTotalContext()
}

type TotalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTotalContext() *TotalContext {
	var p = new(TotalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_total
	return p
}

func (*TotalContext) IsTotalContext() {}

func NewTotalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TotalContext {
	var p = new(TotalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_total

	return p
}

func (s *TotalContext) GetParser() antlr.Parser { return s.parser }

func (s *TotalContext) T() antlr.TerminalNode {
	return s.GetToken(qifParserT, 0)
}

func (s *TotalContext) NUM() antlr.TerminalNode {
	return s.GetToken(qifParserNUM, 0)
}

func (s *TotalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TotalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TotalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterTotal(s)
	}
}

func (s *TotalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitTotal(s)
	}
}

func (p *qifParser) Total() (localctx ITotalContext) {
	this := p
	_ = this

	localctx = NewTotalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, qifParserRULE_total)

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
		p.SetState(55)
		p.Match(qifParserT)
	}
	{
		p.SetState(56)
		p.Match(qifParserNUM)
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
	p.RuleIndex = qifParserRULE_check
	return p
}

func (*CheckContext) IsCheckContext() {}

func NewCheckContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CheckContext {
	var p = new(CheckContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_check

	return p
}

func (s *CheckContext) GetParser() antlr.Parser { return s.parser }

func (s *CheckContext) N() antlr.TerminalNode {
	return s.GetToken(qifParserN, 0)
}

func (s *CheckContext) NUM() antlr.TerminalNode {
	return s.GetToken(qifParserNUM, 0)
}

func (s *CheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterCheck(s)
	}
}

func (s *CheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitCheck(s)
	}
}

func (p *qifParser) Check() (localctx ICheckContext) {
	this := p
	_ = this

	localctx = NewCheckContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, qifParserRULE_check)

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
		p.SetState(58)
		p.Match(qifParserN)
	}
	{
		p.SetState(59)
		p.Match(qifParserNUM)
	}

	return localctx
}

// IState_Context is an interface to support dynamic dispatch.
type IState_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsState_Context differentiates from other interfaces.
	IsState_Context()
}

type State_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyState_Context() *State_Context {
	var p = new(State_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_state_
	return p
}

func (*State_Context) IsState_Context() {}

func NewState_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *State_Context {
	var p = new(State_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_state_

	return p
}

func (s *State_Context) GetParser() antlr.Parser { return s.parser }

func (s *State_Context) C() antlr.TerminalNode {
	return s.GetToken(qifParserC, 0)
}

func (s *State_Context) STATE() antlr.TerminalNode {
	return s.GetToken(qifParserSTATE, 0)
}

func (s *State_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *State_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *State_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterState_(s)
	}
}

func (s *State_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitState_(s)
	}
}

func (p *qifParser) State_() (localctx IState_Context) {
	this := p
	_ = this

	localctx = NewState_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, qifParserRULE_state_)

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
		p.SetState(61)
		p.Match(qifParserC)
	}
	{
		p.SetState(62)
		p.Match(qifParserSTATE)
	}

	return localctx
}

// IMemoContext is an interface to support dynamic dispatch.
type IMemoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemoContext differentiates from other interfaces.
	IsMemoContext()
}

type MemoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemoContext() *MemoContext {
	var p = new(MemoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_memo
	return p
}

func (*MemoContext) IsMemoContext() {}

func NewMemoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemoContext {
	var p = new(MemoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_memo

	return p
}

func (s *MemoContext) GetParser() antlr.Parser { return s.parser }

func (s *MemoContext) M() antlr.TerminalNode {
	return s.GetToken(qifParserM, 0)
}

func (s *MemoContext) TEXT() antlr.TerminalNode {
	return s.GetToken(qifParserTEXT, 0)
}

func (s *MemoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterMemo(s)
	}
}

func (s *MemoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitMemo(s)
	}
}

func (p *qifParser) Memo() (localctx IMemoContext) {
	this := p
	_ = this

	localctx = NewMemoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, qifParserRULE_memo)

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
		p.SetState(64)
		p.Match(qifParserM)
	}
	{
		p.SetState(65)
		p.Match(qifParserTEXT)
	}

	return localctx
}

// IPayeeContext is an interface to support dynamic dispatch.
type IPayeeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPayeeContext differentiates from other interfaces.
	IsPayeeContext()
}

type PayeeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPayeeContext() *PayeeContext {
	var p = new(PayeeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_payee
	return p
}

func (*PayeeContext) IsPayeeContext() {}

func NewPayeeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PayeeContext {
	var p = new(PayeeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_payee

	return p
}

func (s *PayeeContext) GetParser() antlr.Parser { return s.parser }

func (s *PayeeContext) P() antlr.TerminalNode {
	return s.GetToken(qifParserP, 0)
}

func (s *PayeeContext) TEXT() antlr.TerminalNode {
	return s.GetToken(qifParserTEXT, 0)
}

func (s *PayeeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PayeeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PayeeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterPayee(s)
	}
}

func (s *PayeeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitPayee(s)
	}
}

func (p *qifParser) Payee() (localctx IPayeeContext) {
	this := p
	_ = this

	localctx = NewPayeeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, qifParserRULE_payee)

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
		p.Match(qifParserP)
	}
	{
		p.SetState(68)
		p.Match(qifParserTEXT)
	}

	return localctx
}

// IAccountorcategoryContext is an interface to support dynamic dispatch.
type IAccountorcategoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccountorcategoryContext differentiates from other interfaces.
	IsAccountorcategoryContext()
}

type AccountorcategoryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccountorcategoryContext() *AccountorcategoryContext {
	var p = new(AccountorcategoryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_accountorcategory
	return p
}

func (*AccountorcategoryContext) IsAccountorcategoryContext() {}

func NewAccountorcategoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccountorcategoryContext {
	var p = new(AccountorcategoryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_accountorcategory

	return p
}

func (s *AccountorcategoryContext) GetParser() antlr.Parser { return s.parser }

func (s *AccountorcategoryContext) L() antlr.TerminalNode {
	return s.GetToken(qifParserL, 0)
}

func (s *AccountorcategoryContext) Account() IAccountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccountContext)
}

func (s *AccountorcategoryContext) Category() ICategoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICategoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICategoryContext)
}

func (s *AccountorcategoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccountorcategoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccountorcategoryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterAccountorcategory(s)
	}
}

func (s *AccountorcategoryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitAccountorcategory(s)
	}
}

func (p *qifParser) Accountorcategory() (localctx IAccountorcategoryContext) {
	this := p
	_ = this

	localctx = NewAccountorcategoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, qifParserRULE_accountorcategory)

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
		p.Match(qifParserL)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case qifParserLB:
		{
			p.SetState(71)
			p.Account()
		}

	case qifParserACCNTCATNAME:
		{
			p.SetState(72)
			p.Category()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAccountContext is an interface to support dynamic dispatch.
type IAccountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccountContext differentiates from other interfaces.
	IsAccountContext()
}

type AccountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccountContext() *AccountContext {
	var p = new(AccountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_account
	return p
}

func (*AccountContext) IsAccountContext() {}

func NewAccountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccountContext {
	var p = new(AccountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_account

	return p
}

func (s *AccountContext) GetParser() antlr.Parser { return s.parser }

func (s *AccountContext) LB() antlr.TerminalNode {
	return s.GetToken(qifParserLB, 0)
}

func (s *AccountContext) ACCNTCATNAME() antlr.TerminalNode {
	return s.GetToken(qifParserACCNTCATNAME, 0)
}

func (s *AccountContext) RB() antlr.TerminalNode {
	return s.GetToken(qifParserRB, 0)
}

func (s *AccountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterAccount(s)
	}
}

func (s *AccountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitAccount(s)
	}
}

func (p *qifParser) Account() (localctx IAccountContext) {
	this := p
	_ = this

	localctx = NewAccountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, qifParserRULE_account)

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
		p.Match(qifParserLB)
	}
	{
		p.SetState(76)
		p.Match(qifParserACCNTCATNAME)
	}
	{
		p.SetState(77)
		p.Match(qifParserRB)
	}

	return localctx
}

// ICategoryContext is an interface to support dynamic dispatch.
type ICategoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCategoryContext differentiates from other interfaces.
	IsCategoryContext()
}

type CategoryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCategoryContext() *CategoryContext {
	var p = new(CategoryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_category
	return p
}

func (*CategoryContext) IsCategoryContext() {}

func NewCategoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CategoryContext {
	var p = new(CategoryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_category

	return p
}

func (s *CategoryContext) GetParser() antlr.Parser { return s.parser }

func (s *CategoryContext) ACCNTCATNAME() antlr.TerminalNode {
	return s.GetToken(qifParserACCNTCATNAME, 0)
}

func (s *CategoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CategoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CategoryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterCategory(s)
	}
}

func (s *CategoryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitCategory(s)
	}
}

func (p *qifParser) Category() (localctx ICategoryContext) {
	this := p
	_ = this

	localctx = NewCategoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, qifParserRULE_category)

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
		p.SetState(79)
		p.Match(qifParserACCNTCATNAME)
	}

	return localctx
}

// IEorContext is an interface to support dynamic dispatch.
type IEorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEorContext differentiates from other interfaces.
	IsEorContext()
}

type EorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEorContext() *EorContext {
	var p = new(EorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = qifParserRULE_eor
	return p
}

func (*EorContext) IsEorContext() {}

func NewEorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EorContext {
	var p = new(EorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = qifParserRULE_eor

	return p
}

func (s *EorContext) GetParser() antlr.Parser { return s.parser }

func (s *EorContext) EOR() antlr.TerminalNode {
	return s.GetToken(qifParserEOR, 0)
}

func (s *EorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.EnterEor(s)
	}
}

func (s *EorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(qifParserListener); ok {
		listenerT.ExitEor(s)
	}
}

func (p *qifParser) Eor() (localctx IEorContext) {
	this := p
	_ = this

	localctx = NewEorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, qifParserRULE_eor)

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
		p.Match(qifParserEOR)
	}

	return localctx
}
