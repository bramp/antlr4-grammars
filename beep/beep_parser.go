// Code generated from beep.g4 by ANTLR 4.9.3. DO NOT EDIT.

package beep // beep
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 95, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 45, 10, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3,
	17, 5, 17, 93, 10, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 2, 3, 3, 2, 3, 4, 2, 83, 2, 34, 3, 2, 2, 2,
	4, 36, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 46, 3, 2, 2, 2, 10, 50, 3, 2,
	2, 2, 12, 54, 3, 2, 2, 2, 14, 60, 3, 2, 2, 2, 16, 64, 3, 2, 2, 2, 18, 68,
	3, 2, 2, 2, 20, 78, 3, 2, 2, 2, 22, 80, 3, 2, 2, 2, 24, 82, 3, 2, 2, 2,
	26, 84, 3, 2, 2, 2, 28, 86, 3, 2, 2, 2, 30, 88, 3, 2, 2, 2, 32, 90, 3,
	2, 2, 2, 34, 35, 5, 4, 3, 2, 35, 3, 3, 2, 2, 2, 36, 37, 5, 6, 4, 2, 37,
	38, 5, 32, 17, 2, 38, 5, 3, 2, 2, 2, 39, 45, 5, 8, 5, 2, 40, 45, 5, 10,
	6, 2, 41, 45, 5, 14, 8, 2, 42, 45, 5, 12, 7, 2, 43, 45, 5, 16, 9, 2, 44,
	39, 3, 2, 2, 2, 44, 40, 3, 2, 2, 2, 44, 41, 3, 2, 2, 2, 44, 42, 3, 2, 2,
	2, 44, 43, 3, 2, 2, 2, 45, 7, 3, 2, 2, 2, 46, 47, 7, 9, 2, 2, 47, 48, 7,
	11, 2, 2, 48, 49, 5, 18, 10, 2, 49, 9, 3, 2, 2, 2, 50, 51, 7, 8, 2, 2,
	51, 52, 7, 11, 2, 2, 52, 53, 5, 18, 10, 2, 53, 11, 3, 2, 2, 2, 54, 55,
	7, 7, 2, 2, 55, 56, 7, 11, 2, 2, 56, 57, 5, 18, 10, 2, 57, 58, 7, 11, 2,
	2, 58, 59, 5, 30, 16, 2, 59, 13, 3, 2, 2, 2, 60, 61, 7, 6, 2, 2, 61, 62,
	7, 11, 2, 2, 62, 63, 5, 18, 10, 2, 63, 15, 3, 2, 2, 2, 64, 65, 7, 5, 2,
	2, 65, 66, 7, 11, 2, 2, 66, 67, 5, 18, 10, 2, 67, 17, 3, 2, 2, 2, 68, 69,
	5, 20, 11, 2, 69, 70, 7, 11, 2, 2, 70, 71, 5, 22, 12, 2, 71, 72, 7, 11,
	2, 2, 72, 73, 5, 24, 13, 2, 73, 74, 7, 11, 2, 2, 74, 75, 5, 26, 14, 2,
	75, 76, 7, 11, 2, 2, 76, 77, 5, 28, 15, 2, 77, 19, 3, 2, 2, 2, 78, 79,
	7, 10, 2, 2, 79, 21, 3, 2, 2, 2, 80, 81, 7, 10, 2, 2, 81, 23, 3, 2, 2,
	2, 82, 83, 9, 2, 2, 2, 83, 25, 3, 2, 2, 2, 84, 85, 7, 10, 2, 2, 85, 27,
	3, 2, 2, 2, 86, 87, 7, 10, 2, 2, 87, 29, 3, 2, 2, 2, 88, 89, 7, 10, 2,
	2, 89, 31, 3, 2, 2, 2, 90, 92, 7, 13, 2, 2, 91, 93, 7, 12, 2, 2, 92, 91,
	3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 33, 3, 2, 2, 2, 4, 44, 92,
}
var literalNames = []string{
	"", "'.'", "'*'", "'NUL'", "'ERR'", "'ANS'", "'RPY'", "'MSG'", "", "' '",
}
var symbolicNames = []string{
	"", "DOT", "STAR", "NUL", "ERR", "ANS", "RPY", "MSG", "NUMBER", "SP", "CRLF",
	"PAYLOAD_TRAILER",
}

var ruleNames = []string{
	"frame", "data", "header", "msg", "rpy", "ans", "err", "nul", "common",
	"channel", "msgno", "more", "seqno", "size", "ansno", "payload_trailer",
}

type beepParser struct {
	*antlr.BaseParser
}

// NewbeepParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *beepParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewbeepParser(input antlr.TokenStream) *beepParser {
	this := new(beepParser)
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
	this.GrammarFileName = "beep.g4"

	return this
}

// beepParser tokens.
const (
	beepParserEOF             = antlr.TokenEOF
	beepParserDOT             = 1
	beepParserSTAR            = 2
	beepParserNUL             = 3
	beepParserERR             = 4
	beepParserANS             = 5
	beepParserRPY             = 6
	beepParserMSG             = 7
	beepParserNUMBER          = 8
	beepParserSP              = 9
	beepParserCRLF            = 10
	beepParserPAYLOAD_TRAILER = 11
)

// beepParser rules.
const (
	beepParserRULE_frame           = 0
	beepParserRULE_data            = 1
	beepParserRULE_header          = 2
	beepParserRULE_msg             = 3
	beepParserRULE_rpy             = 4
	beepParserRULE_ans             = 5
	beepParserRULE_err             = 6
	beepParserRULE_nul             = 7
	beepParserRULE_common          = 8
	beepParserRULE_channel         = 9
	beepParserRULE_msgno           = 10
	beepParserRULE_more            = 11
	beepParserRULE_seqno           = 12
	beepParserRULE_size            = 13
	beepParserRULE_ansno           = 14
	beepParserRULE_payload_trailer = 15
)

// IFrameContext is an interface to support dynamic dispatch.
type IFrameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFrameContext differentiates from other interfaces.
	IsFrameContext()
}

type FrameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrameContext() *FrameContext {
	var p = new(FrameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_frame
	return p
}

func (*FrameContext) IsFrameContext() {}

func NewFrameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FrameContext {
	var p = new(FrameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_frame

	return p
}

func (s *FrameContext) GetParser() antlr.Parser { return s.parser }

func (s *FrameContext) Data() IDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *FrameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FrameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FrameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterFrame(s)
	}
}

func (s *FrameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitFrame(s)
	}
}

func (p *beepParser) Frame() (localctx IFrameContext) {
	this := p
	_ = this

	localctx = NewFrameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, beepParserRULE_frame)

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
		p.Data()
	}

	return localctx
}

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_data
	return p
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) Header() IHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *DataContext) Payload_trailer() IPayload_trailerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPayload_trailerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPayload_trailerContext)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterData(s)
	}
}

func (s *DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitData(s)
	}
}

func (p *beepParser) Data() (localctx IDataContext) {
	this := p
	_ = this

	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, beepParserRULE_data)

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
		p.SetState(34)
		p.Header()
	}
	{
		p.SetState(35)
		p.Payload_trailer()
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) Msg() IMsgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMsgContext)
}

func (s *HeaderContext) Rpy() IRpyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRpyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRpyContext)
}

func (s *HeaderContext) Err() IErrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IErrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IErrContext)
}

func (s *HeaderContext) Ans() IAnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnsContext)
}

func (s *HeaderContext) Nul() INulContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INulContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INulContext)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *beepParser) Header() (localctx IHeaderContext) {
	this := p
	_ = this

	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, beepParserRULE_header)

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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case beepParserMSG:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Msg()
		}

	case beepParserRPY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Rpy()
		}

	case beepParserERR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.Err()
		}

	case beepParserANS:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.Ans()
		}

	case beepParserNUL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(41)
			p.Nul()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMsgContext is an interface to support dynamic dispatch.
type IMsgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsgContext differentiates from other interfaces.
	IsMsgContext()
}

type MsgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgContext() *MsgContext {
	var p = new(MsgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_msg
	return p
}

func (*MsgContext) IsMsgContext() {}

func NewMsgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgContext {
	var p = new(MsgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_msg

	return p
}

func (s *MsgContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgContext) MSG() antlr.TerminalNode {
	return s.GetToken(beepParserMSG, 0)
}

func (s *MsgContext) SP() antlr.TerminalNode {
	return s.GetToken(beepParserSP, 0)
}

func (s *MsgContext) Common() ICommonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommonContext)
}

func (s *MsgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterMsg(s)
	}
}

func (s *MsgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitMsg(s)
	}
}

func (p *beepParser) Msg() (localctx IMsgContext) {
	this := p
	_ = this

	localctx = NewMsgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, beepParserRULE_msg)

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
		p.SetState(44)
		p.Match(beepParserMSG)
	}
	{
		p.SetState(45)
		p.Match(beepParserSP)
	}
	{
		p.SetState(46)
		p.Common()
	}

	return localctx
}

// IRpyContext is an interface to support dynamic dispatch.
type IRpyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpyContext differentiates from other interfaces.
	IsRpyContext()
}

type RpyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpyContext() *RpyContext {
	var p = new(RpyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_rpy
	return p
}

func (*RpyContext) IsRpyContext() {}

func NewRpyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpyContext {
	var p = new(RpyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_rpy

	return p
}

func (s *RpyContext) GetParser() antlr.Parser { return s.parser }

func (s *RpyContext) RPY() antlr.TerminalNode {
	return s.GetToken(beepParserRPY, 0)
}

func (s *RpyContext) SP() antlr.TerminalNode {
	return s.GetToken(beepParserSP, 0)
}

func (s *RpyContext) Common() ICommonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommonContext)
}

func (s *RpyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterRpy(s)
	}
}

func (s *RpyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitRpy(s)
	}
}

func (p *beepParser) Rpy() (localctx IRpyContext) {
	this := p
	_ = this

	localctx = NewRpyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, beepParserRULE_rpy)

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
		p.SetState(48)
		p.Match(beepParserRPY)
	}
	{
		p.SetState(49)
		p.Match(beepParserSP)
	}
	{
		p.SetState(50)
		p.Common()
	}

	return localctx
}

// IAnsContext is an interface to support dynamic dispatch.
type IAnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnsContext differentiates from other interfaces.
	IsAnsContext()
}

type AnsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnsContext() *AnsContext {
	var p = new(AnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_ans
	return p
}

func (*AnsContext) IsAnsContext() {}

func NewAnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnsContext {
	var p = new(AnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_ans

	return p
}

func (s *AnsContext) GetParser() antlr.Parser { return s.parser }

func (s *AnsContext) ANS() antlr.TerminalNode {
	return s.GetToken(beepParserANS, 0)
}

func (s *AnsContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(beepParserSP)
}

func (s *AnsContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(beepParserSP, i)
}

func (s *AnsContext) Common() ICommonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommonContext)
}

func (s *AnsContext) Ansno() IAnsnoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnsnoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnsnoContext)
}

func (s *AnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterAns(s)
	}
}

func (s *AnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitAns(s)
	}
}

func (p *beepParser) Ans() (localctx IAnsContext) {
	this := p
	_ = this

	localctx = NewAnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, beepParserRULE_ans)

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
		p.Match(beepParserANS)
	}
	{
		p.SetState(53)
		p.Match(beepParserSP)
	}
	{
		p.SetState(54)
		p.Common()
	}
	{
		p.SetState(55)
		p.Match(beepParserSP)
	}
	{
		p.SetState(56)
		p.Ansno()
	}

	return localctx
}

// IErrContext is an interface to support dynamic dispatch.
type IErrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsErrContext differentiates from other interfaces.
	IsErrContext()
}

type ErrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErrContext() *ErrContext {
	var p = new(ErrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_err
	return p
}

func (*ErrContext) IsErrContext() {}

func NewErrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrContext {
	var p = new(ErrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_err

	return p
}

func (s *ErrContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrContext) ERR() antlr.TerminalNode {
	return s.GetToken(beepParserERR, 0)
}

func (s *ErrContext) SP() antlr.TerminalNode {
	return s.GetToken(beepParserSP, 0)
}

func (s *ErrContext) Common() ICommonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommonContext)
}

func (s *ErrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterErr(s)
	}
}

func (s *ErrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitErr(s)
	}
}

func (p *beepParser) Err() (localctx IErrContext) {
	this := p
	_ = this

	localctx = NewErrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, beepParserRULE_err)

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
		p.Match(beepParserERR)
	}
	{
		p.SetState(59)
		p.Match(beepParserSP)
	}
	{
		p.SetState(60)
		p.Common()
	}

	return localctx
}

// INulContext is an interface to support dynamic dispatch.
type INulContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNulContext differentiates from other interfaces.
	IsNulContext()
}

type NulContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNulContext() *NulContext {
	var p = new(NulContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_nul
	return p
}

func (*NulContext) IsNulContext() {}

func NewNulContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NulContext {
	var p = new(NulContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_nul

	return p
}

func (s *NulContext) GetParser() antlr.Parser { return s.parser }

func (s *NulContext) NUL() antlr.TerminalNode {
	return s.GetToken(beepParserNUL, 0)
}

func (s *NulContext) SP() antlr.TerminalNode {
	return s.GetToken(beepParserSP, 0)
}

func (s *NulContext) Common() ICommonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommonContext)
}

func (s *NulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NulContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterNul(s)
	}
}

func (s *NulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitNul(s)
	}
}

func (p *beepParser) Nul() (localctx INulContext) {
	this := p
	_ = this

	localctx = NewNulContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, beepParserRULE_nul)

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
		p.SetState(62)
		p.Match(beepParserNUL)
	}
	{
		p.SetState(63)
		p.Match(beepParserSP)
	}
	{
		p.SetState(64)
		p.Common()
	}

	return localctx
}

// ICommonContext is an interface to support dynamic dispatch.
type ICommonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommonContext differentiates from other interfaces.
	IsCommonContext()
}

type CommonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommonContext() *CommonContext {
	var p = new(CommonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_common
	return p
}

func (*CommonContext) IsCommonContext() {}

func NewCommonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommonContext {
	var p = new(CommonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_common

	return p
}

func (s *CommonContext) GetParser() antlr.Parser { return s.parser }

func (s *CommonContext) Channel() IChannelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChannelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChannelContext)
}

func (s *CommonContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(beepParserSP)
}

func (s *CommonContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(beepParserSP, i)
}

func (s *CommonContext) Msgno() IMsgnoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMsgnoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMsgnoContext)
}

func (s *CommonContext) More() IMoreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMoreContext)
}

func (s *CommonContext) Seqno() ISeqnoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISeqnoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISeqnoContext)
}

func (s *CommonContext) Size() ISizeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISizeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISizeContext)
}

func (s *CommonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterCommon(s)
	}
}

func (s *CommonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitCommon(s)
	}
}

func (p *beepParser) Common() (localctx ICommonContext) {
	this := p
	_ = this

	localctx = NewCommonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, beepParserRULE_common)

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
		p.SetState(66)
		p.Channel()
	}
	{
		p.SetState(67)
		p.Match(beepParserSP)
	}
	{
		p.SetState(68)
		p.Msgno()
	}
	{
		p.SetState(69)
		p.Match(beepParserSP)
	}
	{
		p.SetState(70)
		p.More()
	}
	{
		p.SetState(71)
		p.Match(beepParserSP)
	}
	{
		p.SetState(72)
		p.Seqno()
	}
	{
		p.SetState(73)
		p.Match(beepParserSP)
	}
	{
		p.SetState(74)
		p.Size()
	}

	return localctx
}

// IChannelContext is an interface to support dynamic dispatch.
type IChannelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChannelContext differentiates from other interfaces.
	IsChannelContext()
}

type ChannelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChannelContext() *ChannelContext {
	var p = new(ChannelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_channel
	return p
}

func (*ChannelContext) IsChannelContext() {}

func NewChannelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChannelContext {
	var p = new(ChannelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_channel

	return p
}

func (s *ChannelContext) GetParser() antlr.Parser { return s.parser }

func (s *ChannelContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(beepParserNUMBER, 0)
}

func (s *ChannelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChannelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChannelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterChannel(s)
	}
}

func (s *ChannelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitChannel(s)
	}
}

func (p *beepParser) Channel() (localctx IChannelContext) {
	this := p
	_ = this

	localctx = NewChannelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, beepParserRULE_channel)

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
		p.SetState(76)
		p.Match(beepParserNUMBER)
	}

	return localctx
}

// IMsgnoContext is an interface to support dynamic dispatch.
type IMsgnoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMsgnoContext differentiates from other interfaces.
	IsMsgnoContext()
}

type MsgnoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMsgnoContext() *MsgnoContext {
	var p = new(MsgnoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_msgno
	return p
}

func (*MsgnoContext) IsMsgnoContext() {}

func NewMsgnoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MsgnoContext {
	var p = new(MsgnoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_msgno

	return p
}

func (s *MsgnoContext) GetParser() antlr.Parser { return s.parser }

func (s *MsgnoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(beepParserNUMBER, 0)
}

func (s *MsgnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MsgnoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MsgnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterMsgno(s)
	}
}

func (s *MsgnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitMsgno(s)
	}
}

func (p *beepParser) Msgno() (localctx IMsgnoContext) {
	this := p
	_ = this

	localctx = NewMsgnoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, beepParserRULE_msgno)

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
		p.SetState(78)
		p.Match(beepParserNUMBER)
	}

	return localctx
}

// IMoreContext is an interface to support dynamic dispatch.
type IMoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoreContext differentiates from other interfaces.
	IsMoreContext()
}

type MoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoreContext() *MoreContext {
	var p = new(MoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_more
	return p
}

func (*MoreContext) IsMoreContext() {}

func NewMoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoreContext {
	var p = new(MoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_more

	return p
}

func (s *MoreContext) GetParser() antlr.Parser { return s.parser }

func (s *MoreContext) DOT() antlr.TerminalNode {
	return s.GetToken(beepParserDOT, 0)
}

func (s *MoreContext) STAR() antlr.TerminalNode {
	return s.GetToken(beepParserSTAR, 0)
}

func (s *MoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterMore(s)
	}
}

func (s *MoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitMore(s)
	}
}

func (p *beepParser) More() (localctx IMoreContext) {
	this := p
	_ = this

	localctx = NewMoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, beepParserRULE_more)
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
		p.SetState(80)
		_la = p.GetTokenStream().LA(1)

		if !(_la == beepParserDOT || _la == beepParserSTAR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISeqnoContext is an interface to support dynamic dispatch.
type ISeqnoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSeqnoContext differentiates from other interfaces.
	IsSeqnoContext()
}

type SeqnoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeqnoContext() *SeqnoContext {
	var p = new(SeqnoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_seqno
	return p
}

func (*SeqnoContext) IsSeqnoContext() {}

func NewSeqnoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeqnoContext {
	var p = new(SeqnoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_seqno

	return p
}

func (s *SeqnoContext) GetParser() antlr.Parser { return s.parser }

func (s *SeqnoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(beepParserNUMBER, 0)
}

func (s *SeqnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeqnoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeqnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterSeqno(s)
	}
}

func (s *SeqnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitSeqno(s)
	}
}

func (p *beepParser) Seqno() (localctx ISeqnoContext) {
	this := p
	_ = this

	localctx = NewSeqnoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, beepParserRULE_seqno)

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
		p.SetState(82)
		p.Match(beepParserNUMBER)
	}

	return localctx
}

// ISizeContext is an interface to support dynamic dispatch.
type ISizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSizeContext differentiates from other interfaces.
	IsSizeContext()
}

type SizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySizeContext() *SizeContext {
	var p = new(SizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_size
	return p
}

func (*SizeContext) IsSizeContext() {}

func NewSizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SizeContext {
	var p = new(SizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_size

	return p
}

func (s *SizeContext) GetParser() antlr.Parser { return s.parser }

func (s *SizeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(beepParserNUMBER, 0)
}

func (s *SizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterSize(s)
	}
}

func (s *SizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitSize(s)
	}
}

func (p *beepParser) Size() (localctx ISizeContext) {
	this := p
	_ = this

	localctx = NewSizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, beepParserRULE_size)

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
		p.SetState(84)
		p.Match(beepParserNUMBER)
	}

	return localctx
}

// IAnsnoContext is an interface to support dynamic dispatch.
type IAnsnoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnsnoContext differentiates from other interfaces.
	IsAnsnoContext()
}

type AnsnoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnsnoContext() *AnsnoContext {
	var p = new(AnsnoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_ansno
	return p
}

func (*AnsnoContext) IsAnsnoContext() {}

func NewAnsnoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnsnoContext {
	var p = new(AnsnoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_ansno

	return p
}

func (s *AnsnoContext) GetParser() antlr.Parser { return s.parser }

func (s *AnsnoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(beepParserNUMBER, 0)
}

func (s *AnsnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnsnoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnsnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterAnsno(s)
	}
}

func (s *AnsnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitAnsno(s)
	}
}

func (p *beepParser) Ansno() (localctx IAnsnoContext) {
	this := p
	_ = this

	localctx = NewAnsnoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, beepParserRULE_ansno)

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
		p.SetState(86)
		p.Match(beepParserNUMBER)
	}

	return localctx
}

// IPayload_trailerContext is an interface to support dynamic dispatch.
type IPayload_trailerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPayload_trailerContext differentiates from other interfaces.
	IsPayload_trailerContext()
}

type Payload_trailerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPayload_trailerContext() *Payload_trailerContext {
	var p = new(Payload_trailerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = beepParserRULE_payload_trailer
	return p
}

func (*Payload_trailerContext) IsPayload_trailerContext() {}

func NewPayload_trailerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Payload_trailerContext {
	var p = new(Payload_trailerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = beepParserRULE_payload_trailer

	return p
}

func (s *Payload_trailerContext) GetParser() antlr.Parser { return s.parser }

func (s *Payload_trailerContext) PAYLOAD_TRAILER() antlr.TerminalNode {
	return s.GetToken(beepParserPAYLOAD_TRAILER, 0)
}

func (s *Payload_trailerContext) CRLF() antlr.TerminalNode {
	return s.GetToken(beepParserCRLF, 0)
}

func (s *Payload_trailerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Payload_trailerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Payload_trailerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.EnterPayload_trailer(s)
	}
}

func (s *Payload_trailerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(beepListener); ok {
		listenerT.ExitPayload_trailer(s)
	}
}

func (p *beepParser) Payload_trailer() (localctx IPayload_trailerContext) {
	this := p
	_ = this

	localctx = NewPayload_trailerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, beepParserRULE_payload_trailer)
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
		p.SetState(88)
		p.Match(beepParserPAYLOAD_TRAILER)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == beepParserCRLF {
		{
			p.SetState(89)
			p.Match(beepParserCRLF)
		}

	}

	return localctx
}
