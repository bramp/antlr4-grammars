// Code generated from clf.g4 by ANTLR 4.7.2. DO NOT EDIT.

package clf // clf
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 71, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 5,
	2, 26, 10, 2, 3, 2, 6, 2, 29, 10, 2, 13, 2, 14, 2, 30, 3, 2, 5, 2, 34,
	10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	46, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 2, 2, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2,
	3, 3, 2, 10, 11, 2, 63, 2, 28, 3, 2, 2, 2, 4, 35, 3, 2, 2, 2, 6, 47, 3,
	2, 2, 2, 8, 49, 3, 2, 2, 2, 10, 51, 3, 2, 2, 2, 12, 53, 3, 2, 2, 2, 14,
	60, 3, 2, 2, 2, 16, 62, 3, 2, 2, 2, 18, 64, 3, 2, 2, 2, 20, 66, 3, 2, 2,
	2, 22, 68, 3, 2, 2, 2, 24, 26, 5, 4, 3, 2, 25, 24, 3, 2, 2, 2, 25, 26,
	3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 29, 7, 12, 2, 2, 28, 25, 3, 2, 2, 2,
	29, 30, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 3,
	2, 2, 2, 32, 34, 5, 4, 3, 2, 33, 32, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34,
	3, 3, 2, 2, 2, 35, 36, 5, 6, 4, 2, 36, 37, 5, 8, 5, 2, 37, 38, 5, 10, 6,
	2, 38, 39, 5, 12, 7, 2, 39, 40, 5, 16, 9, 2, 40, 41, 5, 20, 11, 2, 41,
	45, 5, 22, 12, 2, 42, 43, 5, 14, 8, 2, 43, 44, 5, 18, 10, 2, 44, 46, 3,
	2, 2, 2, 45, 42, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 5, 3, 2, 2, 2, 47,
	48, 9, 2, 2, 2, 48, 7, 3, 2, 2, 2, 49, 50, 7, 11, 2, 2, 50, 9, 3, 2, 2,
	2, 51, 52, 7, 11, 2, 2, 52, 11, 3, 2, 2, 2, 53, 54, 7, 3, 2, 2, 54, 55,
	7, 6, 2, 2, 55, 56, 7, 4, 2, 2, 56, 57, 7, 7, 2, 2, 57, 58, 7, 8, 2, 2,
	58, 59, 7, 5, 2, 2, 59, 13, 3, 2, 2, 2, 60, 61, 7, 9, 2, 2, 61, 15, 3,
	2, 2, 2, 62, 63, 7, 9, 2, 2, 63, 17, 3, 2, 2, 2, 64, 65, 7, 9, 2, 2, 65,
	19, 3, 2, 2, 2, 66, 67, 7, 11, 2, 2, 67, 21, 3, 2, 2, 2, 68, 69, 7, 11,
	2, 2, 69, 23, 3, 2, 2, 2, 6, 25, 30, 33, 45,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "':'", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "DATE", "TIME", "TZ", "LITERAL", "IP", "STRING", "EOL",
	"WS",
}

var ruleNames = []string{
	"log", "line", "host", "logname", "username", "datetimetz", "referer",
	"request", "useragent", "statuscode", "bytes",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type clfParser struct {
	*antlr.BaseParser
}

func NewclfParser(input antlr.TokenStream) *clfParser {
	this := new(clfParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "clf.g4"

	return this
}

// clfParser tokens.
const (
	clfParserEOF     = antlr.TokenEOF
	clfParserT__0    = 1
	clfParserT__1    = 2
	clfParserT__2    = 3
	clfParserDATE    = 4
	clfParserTIME    = 5
	clfParserTZ      = 6
	clfParserLITERAL = 7
	clfParserIP      = 8
	clfParserSTRING  = 9
	clfParserEOL     = 10
	clfParserWS      = 11
)

// clfParser rules.
const (
	clfParserRULE_log        = 0
	clfParserRULE_line       = 1
	clfParserRULE_host       = 2
	clfParserRULE_logname    = 3
	clfParserRULE_username   = 4
	clfParserRULE_datetimetz = 5
	clfParserRULE_referer    = 6
	clfParserRULE_request    = 7
	clfParserRULE_useragent  = 8
	clfParserRULE_statuscode = 9
	clfParserRULE_bytes      = 10
)

// ILogContext is an interface to support dynamic dispatch.
type ILogContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogContext differentiates from other interfaces.
	IsLogContext()
}

type LogContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogContext() *LogContext {
	var p = new(LogContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = clfParserRULE_log
	return p
}

func (*LogContext) IsLogContext() {}

func NewLogContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogContext {
	var p = new(LogContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_log

	return p
}

func (s *LogContext) GetParser() antlr.Parser { return s.parser }

func (s *LogContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(clfParserEOL)
}

func (s *LogContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(clfParserEOL, i)
}

func (s *LogContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *LogContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *LogContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterLog(s)
	}
}

func (s *LogContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitLog(s)
	}
}

func (p *clfParser) Log() (localctx ILogContext) {
	localctx = NewLogContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, clfParserRULE_log)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(23)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == clfParserIP || _la == clfParserSTRING {
				{
					p.SetState(22)
					p.Line()
				}

			}
			{
				p.SetState(25)
				p.Match(clfParserEOL)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == clfParserIP || _la == clfParserSTRING {
		{
			p.SetState(30)
			p.Line()
		}

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
	p.RuleIndex = clfParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Host() IHostContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHostContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHostContext)
}

func (s *LineContext) Logname() ILognameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILognameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILognameContext)
}

func (s *LineContext) Username() IUsernameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUsernameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUsernameContext)
}

func (s *LineContext) Datetimetz() IDatetimetzContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimetzContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimetzContext)
}

func (s *LineContext) Request() IRequestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequestContext)
}

func (s *LineContext) Statuscode() IStatuscodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatuscodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatuscodeContext)
}

func (s *LineContext) Bytes() IBytesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBytesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBytesContext)
}

func (s *LineContext) Referer() IRefererContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRefererContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRefererContext)
}

func (s *LineContext) Useragent() IUseragentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUseragentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUseragentContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *clfParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, clfParserRULE_line)
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
		p.SetState(33)
		p.Host()
	}
	{
		p.SetState(34)
		p.Logname()
	}
	{
		p.SetState(35)
		p.Username()
	}
	{
		p.SetState(36)
		p.Datetimetz()
	}
	{
		p.SetState(37)
		p.Request()
	}
	{
		p.SetState(38)
		p.Statuscode()
	}
	{
		p.SetState(39)
		p.Bytes()
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == clfParserLITERAL {
		{
			p.SetState(40)
			p.Referer()
		}
		{
			p.SetState(41)
			p.Useragent()
		}

	}

	return localctx
}

// IHostContext is an interface to support dynamic dispatch.
type IHostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHostContext differentiates from other interfaces.
	IsHostContext()
}

type HostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHostContext() *HostContext {
	var p = new(HostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = clfParserRULE_host
	return p
}

func (*HostContext) IsHostContext() {}

func NewHostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HostContext {
	var p = new(HostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_host

	return p
}

func (s *HostContext) GetParser() antlr.Parser { return s.parser }

func (s *HostContext) STRING() antlr.TerminalNode {
	return s.GetToken(clfParserSTRING, 0)
}

func (s *HostContext) IP() antlr.TerminalNode {
	return s.GetToken(clfParserIP, 0)
}

func (s *HostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterHost(s)
	}
}

func (s *HostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitHost(s)
	}
}

func (p *clfParser) Host() (localctx IHostContext) {
	localctx = NewHostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, clfParserRULE_host)
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
		p.SetState(45)
		_la = p.GetTokenStream().LA(1)

		if !(_la == clfParserIP || _la == clfParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILognameContext is an interface to support dynamic dispatch.
type ILognameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLognameContext differentiates from other interfaces.
	IsLognameContext()
}

type LognameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLognameContext() *LognameContext {
	var p = new(LognameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = clfParserRULE_logname
	return p
}

func (*LognameContext) IsLognameContext() {}

func NewLognameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LognameContext {
	var p = new(LognameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_logname

	return p
}

func (s *LognameContext) GetParser() antlr.Parser { return s.parser }

func (s *LognameContext) STRING() antlr.TerminalNode {
	return s.GetToken(clfParserSTRING, 0)
}

func (s *LognameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LognameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LognameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterLogname(s)
	}
}

func (s *LognameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitLogname(s)
	}
}

func (p *clfParser) Logname() (localctx ILognameContext) {
	localctx = NewLognameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, clfParserRULE_logname)

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
		p.SetState(47)
		p.Match(clfParserSTRING)
	}

	return localctx
}

// IUsernameContext is an interface to support dynamic dispatch.
type IUsernameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUsernameContext differentiates from other interfaces.
	IsUsernameContext()
}

type UsernameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUsernameContext() *UsernameContext {
	var p = new(UsernameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = clfParserRULE_username
	return p
}

func (*UsernameContext) IsUsernameContext() {}

func NewUsernameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UsernameContext {
	var p = new(UsernameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_username

	return p
}

func (s *UsernameContext) GetParser() antlr.Parser { return s.parser }

func (s *UsernameContext) STRING() antlr.TerminalNode {
	return s.GetToken(clfParserSTRING, 0)
}

func (s *UsernameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UsernameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UsernameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterUsername(s)
	}
}

func (s *UsernameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitUsername(s)
	}
}

func (p *clfParser) Username() (localctx IUsernameContext) {
	localctx = NewUsernameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, clfParserRULE_username)

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
		p.Match(clfParserSTRING)
	}

	return localctx
}

// IDatetimetzContext is an interface to support dynamic dispatch.
type IDatetimetzContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetimetzContext differentiates from other interfaces.
	IsDatetimetzContext()
}

type DatetimetzContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimetzContext() *DatetimetzContext {
	var p = new(DatetimetzContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = clfParserRULE_datetimetz
	return p
}

func (*DatetimetzContext) IsDatetimetzContext() {}

func NewDatetimetzContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimetzContext {
	var p = new(DatetimetzContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_datetimetz

	return p
}

func (s *DatetimetzContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimetzContext) DATE() antlr.TerminalNode {
	return s.GetToken(clfParserDATE, 0)
}

func (s *DatetimetzContext) TIME() antlr.TerminalNode {
	return s.GetToken(clfParserTIME, 0)
}

func (s *DatetimetzContext) TZ() antlr.TerminalNode {
	return s.GetToken(clfParserTZ, 0)
}

func (s *DatetimetzContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimetzContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimetzContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterDatetimetz(s)
	}
}

func (s *DatetimetzContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitDatetimetz(s)
	}
}

func (p *clfParser) Datetimetz() (localctx IDatetimetzContext) {
	localctx = NewDatetimetzContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, clfParserRULE_datetimetz)

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
		p.SetState(51)
		p.Match(clfParserT__0)
	}
	{
		p.SetState(52)
		p.Match(clfParserDATE)
	}
	{
		p.SetState(53)
		p.Match(clfParserT__1)
	}
	{
		p.SetState(54)
		p.Match(clfParserTIME)
	}
	{
		p.SetState(55)
		p.Match(clfParserTZ)
	}
	{
		p.SetState(56)
		p.Match(clfParserT__2)
	}

	return localctx
}

// IRefererContext is an interface to support dynamic dispatch.
type IRefererContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRefererContext differentiates from other interfaces.
	IsRefererContext()
}

type RefererContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRefererContext() *RefererContext {
	var p = new(RefererContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = clfParserRULE_referer
	return p
}

func (*RefererContext) IsRefererContext() {}

func NewRefererContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RefererContext {
	var p = new(RefererContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_referer

	return p
}

func (s *RefererContext) GetParser() antlr.Parser { return s.parser }

func (s *RefererContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(clfParserLITERAL, 0)
}

func (s *RefererContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefererContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RefererContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterReferer(s)
	}
}

func (s *RefererContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitReferer(s)
	}
}

func (p *clfParser) Referer() (localctx IRefererContext) {
	localctx = NewRefererContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, clfParserRULE_referer)

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
		p.Match(clfParserLITERAL)
	}

	return localctx
}

// IRequestContext is an interface to support dynamic dispatch.
type IRequestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequestContext differentiates from other interfaces.
	IsRequestContext()
}

type RequestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestContext() *RequestContext {
	var p = new(RequestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = clfParserRULE_request
	return p
}

func (*RequestContext) IsRequestContext() {}

func NewRequestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestContext {
	var p = new(RequestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_request

	return p
}

func (s *RequestContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(clfParserLITERAL, 0)
}

func (s *RequestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterRequest(s)
	}
}

func (s *RequestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitRequest(s)
	}
}

func (p *clfParser) Request() (localctx IRequestContext) {
	localctx = NewRequestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, clfParserRULE_request)

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
		p.SetState(60)
		p.Match(clfParserLITERAL)
	}

	return localctx
}

// IUseragentContext is an interface to support dynamic dispatch.
type IUseragentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUseragentContext differentiates from other interfaces.
	IsUseragentContext()
}

type UseragentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseragentContext() *UseragentContext {
	var p = new(UseragentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = clfParserRULE_useragent
	return p
}

func (*UseragentContext) IsUseragentContext() {}

func NewUseragentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseragentContext {
	var p = new(UseragentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_useragent

	return p
}

func (s *UseragentContext) GetParser() antlr.Parser { return s.parser }

func (s *UseragentContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(clfParserLITERAL, 0)
}

func (s *UseragentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseragentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseragentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterUseragent(s)
	}
}

func (s *UseragentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitUseragent(s)
	}
}

func (p *clfParser) Useragent() (localctx IUseragentContext) {
	localctx = NewUseragentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, clfParserRULE_useragent)

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
		p.Match(clfParserLITERAL)
	}

	return localctx
}

// IStatuscodeContext is an interface to support dynamic dispatch.
type IStatuscodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatuscodeContext differentiates from other interfaces.
	IsStatuscodeContext()
}

type StatuscodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatuscodeContext() *StatuscodeContext {
	var p = new(StatuscodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = clfParserRULE_statuscode
	return p
}

func (*StatuscodeContext) IsStatuscodeContext() {}

func NewStatuscodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatuscodeContext {
	var p = new(StatuscodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_statuscode

	return p
}

func (s *StatuscodeContext) GetParser() antlr.Parser { return s.parser }

func (s *StatuscodeContext) STRING() antlr.TerminalNode {
	return s.GetToken(clfParserSTRING, 0)
}

func (s *StatuscodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatuscodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatuscodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterStatuscode(s)
	}
}

func (s *StatuscodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitStatuscode(s)
	}
}

func (p *clfParser) Statuscode() (localctx IStatuscodeContext) {
	localctx = NewStatuscodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, clfParserRULE_statuscode)

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
		p.Match(clfParserSTRING)
	}

	return localctx
}

// IBytesContext is an interface to support dynamic dispatch.
type IBytesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBytesContext differentiates from other interfaces.
	IsBytesContext()
}

type BytesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBytesContext() *BytesContext {
	var p = new(BytesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = clfParserRULE_bytes
	return p
}

func (*BytesContext) IsBytesContext() {}

func NewBytesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BytesContext {
	var p = new(BytesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = clfParserRULE_bytes

	return p
}

func (s *BytesContext) GetParser() antlr.Parser { return s.parser }

func (s *BytesContext) STRING() antlr.TerminalNode {
	return s.GetToken(clfParserSTRING, 0)
}

func (s *BytesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BytesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BytesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.EnterBytes(s)
	}
}

func (s *BytesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(clfListener); ok {
		listenerT.ExitBytes(s)
	}
}

func (p *clfParser) Bytes() (localctx IBytesContext) {
	localctx = NewBytesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, clfParserRULE_bytes)

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
		p.Match(clfParserSTRING)
	}

	return localctx
}
