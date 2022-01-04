// Code generated from tcp.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tcp // tcp
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 4, 62, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2,
	2, 2, 49, 2, 26, 3, 2, 2, 2, 4, 35, 3, 2, 2, 2, 6, 37, 3, 2, 2, 2, 8, 39,
	3, 2, 2, 2, 10, 41, 3, 2, 2, 2, 12, 43, 3, 2, 2, 2, 14, 45, 3, 2, 2, 2,
	16, 47, 3, 2, 2, 2, 18, 49, 3, 2, 2, 2, 20, 51, 3, 2, 2, 2, 22, 56, 3,
	2, 2, 2, 24, 59, 3, 2, 2, 2, 26, 27, 5, 4, 3, 2, 27, 28, 5, 6, 4, 2, 28,
	29, 5, 8, 5, 2, 29, 30, 5, 10, 6, 2, 30, 31, 5, 12, 7, 2, 31, 32, 5, 14,
	8, 2, 32, 33, 5, 16, 9, 2, 33, 34, 5, 18, 10, 2, 34, 3, 3, 2, 2, 2, 35,
	36, 5, 22, 12, 2, 36, 5, 3, 2, 2, 2, 37, 38, 5, 22, 12, 2, 38, 7, 3, 2,
	2, 2, 39, 40, 5, 20, 11, 2, 40, 9, 3, 2, 2, 2, 41, 42, 5, 20, 11, 2, 42,
	11, 3, 2, 2, 2, 43, 44, 5, 22, 12, 2, 44, 13, 3, 2, 2, 2, 45, 46, 5, 22,
	12, 2, 46, 15, 3, 2, 2, 2, 47, 48, 5, 22, 12, 2, 48, 17, 3, 2, 2, 2, 49,
	50, 5, 22, 12, 2, 50, 19, 3, 2, 2, 2, 51, 52, 7, 3, 2, 2, 52, 53, 7, 3,
	2, 2, 53, 54, 7, 3, 2, 2, 54, 55, 7, 3, 2, 2, 55, 21, 3, 2, 2, 2, 56, 57,
	7, 3, 2, 2, 57, 58, 7, 3, 2, 2, 58, 23, 3, 2, 2, 2, 59, 60, 7, 3, 2, 2,
	60, 25, 3, 2, 2, 2, 2,
}
var literalNames []string

var symbolicNames = []string{
	"", "BYTE", "WS",
}

var ruleNames = []string{
	"segmentheader", "sourceport", "destinationport", "sequencenumber", "acknumber",
	"flags", "windowsize", "checksum", "urgent", "dword_", "word_", "byte_",
}

type tcpParser struct {
	*antlr.BaseParser
}

// NewtcpParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *tcpParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewtcpParser(input antlr.TokenStream) *tcpParser {
	this := new(tcpParser)
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
	this.GrammarFileName = "tcp.g4"

	return this
}

// tcpParser tokens.
const (
	tcpParserEOF  = antlr.TokenEOF
	tcpParserBYTE = 1
	tcpParserWS   = 2
)

// tcpParser rules.
const (
	tcpParserRULE_segmentheader   = 0
	tcpParserRULE_sourceport      = 1
	tcpParserRULE_destinationport = 2
	tcpParserRULE_sequencenumber  = 3
	tcpParserRULE_acknumber       = 4
	tcpParserRULE_flags           = 5
	tcpParserRULE_windowsize      = 6
	tcpParserRULE_checksum        = 7
	tcpParserRULE_urgent          = 8
	tcpParserRULE_dword_          = 9
	tcpParserRULE_word_           = 10
	tcpParserRULE_byte_           = 11
)

// ISegmentheaderContext is an interface to support dynamic dispatch.
type ISegmentheaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentheaderContext differentiates from other interfaces.
	IsSegmentheaderContext()
}

type SegmentheaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentheaderContext() *SegmentheaderContext {
	var p = new(SegmentheaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_segmentheader
	return p
}

func (*SegmentheaderContext) IsSegmentheaderContext() {}

func NewSegmentheaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentheaderContext {
	var p = new(SegmentheaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_segmentheader

	return p
}

func (s *SegmentheaderContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentheaderContext) Sourceport() ISourceportContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceportContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceportContext)
}

func (s *SegmentheaderContext) Destinationport() IDestinationportContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestinationportContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestinationportContext)
}

func (s *SegmentheaderContext) Sequencenumber() ISequencenumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequencenumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequencenumberContext)
}

func (s *SegmentheaderContext) Acknumber() IAcknumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAcknumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAcknumberContext)
}

func (s *SegmentheaderContext) Flags() IFlagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlagsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlagsContext)
}

func (s *SegmentheaderContext) Windowsize() IWindowsizeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWindowsizeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWindowsizeContext)
}

func (s *SegmentheaderContext) Checksum() IChecksumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChecksumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChecksumContext)
}

func (s *SegmentheaderContext) Urgent() IUrgentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUrgentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUrgentContext)
}

func (s *SegmentheaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentheaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentheaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterSegmentheader(s)
	}
}

func (s *SegmentheaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitSegmentheader(s)
	}
}

func (p *tcpParser) Segmentheader() (localctx ISegmentheaderContext) {
	this := p
	_ = this

	localctx = NewSegmentheaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tcpParserRULE_segmentheader)

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
		p.Sourceport()
	}
	{
		p.SetState(25)
		p.Destinationport()
	}
	{
		p.SetState(26)
		p.Sequencenumber()
	}
	{
		p.SetState(27)
		p.Acknumber()
	}
	{
		p.SetState(28)
		p.Flags()
	}
	{
		p.SetState(29)
		p.Windowsize()
	}
	{
		p.SetState(30)
		p.Checksum()
	}
	{
		p.SetState(31)
		p.Urgent()
	}

	return localctx
}

// ISourceportContext is an interface to support dynamic dispatch.
type ISourceportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceportContext differentiates from other interfaces.
	IsSourceportContext()
}

type SourceportContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceportContext() *SourceportContext {
	var p = new(SourceportContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_sourceport
	return p
}

func (*SourceportContext) IsSourceportContext() {}

func NewSourceportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceportContext {
	var p = new(SourceportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_sourceport

	return p
}

func (s *SourceportContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceportContext) Word_() IWord_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWord_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWord_Context)
}

func (s *SourceportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterSourceport(s)
	}
}

func (s *SourceportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitSourceport(s)
	}
}

func (p *tcpParser) Sourceport() (localctx ISourceportContext) {
	this := p
	_ = this

	localctx = NewSourceportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tcpParserRULE_sourceport)

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
		p.Word_()
	}

	return localctx
}

// IDestinationportContext is an interface to support dynamic dispatch.
type IDestinationportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDestinationportContext differentiates from other interfaces.
	IsDestinationportContext()
}

type DestinationportContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestinationportContext() *DestinationportContext {
	var p = new(DestinationportContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_destinationport
	return p
}

func (*DestinationportContext) IsDestinationportContext() {}

func NewDestinationportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinationportContext {
	var p = new(DestinationportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_destinationport

	return p
}

func (s *DestinationportContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinationportContext) Word_() IWord_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWord_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWord_Context)
}

func (s *DestinationportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestinationportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterDestinationport(s)
	}
}

func (s *DestinationportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitDestinationport(s)
	}
}

func (p *tcpParser) Destinationport() (localctx IDestinationportContext) {
	this := p
	_ = this

	localctx = NewDestinationportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tcpParserRULE_destinationport)

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
		p.Word_()
	}

	return localctx
}

// ISequencenumberContext is an interface to support dynamic dispatch.
type ISequencenumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequencenumberContext differentiates from other interfaces.
	IsSequencenumberContext()
}

type SequencenumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequencenumberContext() *SequencenumberContext {
	var p = new(SequencenumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_sequencenumber
	return p
}

func (*SequencenumberContext) IsSequencenumberContext() {}

func NewSequencenumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequencenumberContext {
	var p = new(SequencenumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_sequencenumber

	return p
}

func (s *SequencenumberContext) GetParser() antlr.Parser { return s.parser }

func (s *SequencenumberContext) Dword_() IDword_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDword_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDword_Context)
}

func (s *SequencenumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequencenumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequencenumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterSequencenumber(s)
	}
}

func (s *SequencenumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitSequencenumber(s)
	}
}

func (p *tcpParser) Sequencenumber() (localctx ISequencenumberContext) {
	this := p
	_ = this

	localctx = NewSequencenumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tcpParserRULE_sequencenumber)

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
		p.SetState(37)
		p.Dword_()
	}

	return localctx
}

// IAcknumberContext is an interface to support dynamic dispatch.
type IAcknumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAcknumberContext differentiates from other interfaces.
	IsAcknumberContext()
}

type AcknumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAcknumberContext() *AcknumberContext {
	var p = new(AcknumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_acknumber
	return p
}

func (*AcknumberContext) IsAcknumberContext() {}

func NewAcknumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AcknumberContext {
	var p = new(AcknumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_acknumber

	return p
}

func (s *AcknumberContext) GetParser() antlr.Parser { return s.parser }

func (s *AcknumberContext) Dword_() IDword_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDword_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDword_Context)
}

func (s *AcknumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AcknumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AcknumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterAcknumber(s)
	}
}

func (s *AcknumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitAcknumber(s)
	}
}

func (p *tcpParser) Acknumber() (localctx IAcknumberContext) {
	this := p
	_ = this

	localctx = NewAcknumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tcpParserRULE_acknumber)

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
		p.SetState(39)
		p.Dword_()
	}

	return localctx
}

// IFlagsContext is an interface to support dynamic dispatch.
type IFlagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlagsContext differentiates from other interfaces.
	IsFlagsContext()
}

type FlagsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlagsContext() *FlagsContext {
	var p = new(FlagsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_flags
	return p
}

func (*FlagsContext) IsFlagsContext() {}

func NewFlagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlagsContext {
	var p = new(FlagsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_flags

	return p
}

func (s *FlagsContext) GetParser() antlr.Parser { return s.parser }

func (s *FlagsContext) Word_() IWord_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWord_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWord_Context)
}

func (s *FlagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterFlags(s)
	}
}

func (s *FlagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitFlags(s)
	}
}

func (p *tcpParser) Flags() (localctx IFlagsContext) {
	this := p
	_ = this

	localctx = NewFlagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tcpParserRULE_flags)

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
		p.SetState(41)
		p.Word_()
	}

	return localctx
}

// IWindowsizeContext is an interface to support dynamic dispatch.
type IWindowsizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindowsizeContext differentiates from other interfaces.
	IsWindowsizeContext()
}

type WindowsizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowsizeContext() *WindowsizeContext {
	var p = new(WindowsizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_windowsize
	return p
}

func (*WindowsizeContext) IsWindowsizeContext() {}

func NewWindowsizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowsizeContext {
	var p = new(WindowsizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_windowsize

	return p
}

func (s *WindowsizeContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowsizeContext) Word_() IWord_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWord_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWord_Context)
}

func (s *WindowsizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowsizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowsizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterWindowsize(s)
	}
}

func (s *WindowsizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitWindowsize(s)
	}
}

func (p *tcpParser) Windowsize() (localctx IWindowsizeContext) {
	this := p
	_ = this

	localctx = NewWindowsizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tcpParserRULE_windowsize)

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
		p.SetState(43)
		p.Word_()
	}

	return localctx
}

// IChecksumContext is an interface to support dynamic dispatch.
type IChecksumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChecksumContext differentiates from other interfaces.
	IsChecksumContext()
}

type ChecksumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChecksumContext() *ChecksumContext {
	var p = new(ChecksumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_checksum
	return p
}

func (*ChecksumContext) IsChecksumContext() {}

func NewChecksumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChecksumContext {
	var p = new(ChecksumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_checksum

	return p
}

func (s *ChecksumContext) GetParser() antlr.Parser { return s.parser }

func (s *ChecksumContext) Word_() IWord_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWord_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWord_Context)
}

func (s *ChecksumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChecksumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChecksumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterChecksum(s)
	}
}

func (s *ChecksumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitChecksum(s)
	}
}

func (p *tcpParser) Checksum() (localctx IChecksumContext) {
	this := p
	_ = this

	localctx = NewChecksumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tcpParserRULE_checksum)

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
		p.Word_()
	}

	return localctx
}

// IUrgentContext is an interface to support dynamic dispatch.
type IUrgentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUrgentContext differentiates from other interfaces.
	IsUrgentContext()
}

type UrgentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrgentContext() *UrgentContext {
	var p = new(UrgentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_urgent
	return p
}

func (*UrgentContext) IsUrgentContext() {}

func NewUrgentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UrgentContext {
	var p = new(UrgentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_urgent

	return p
}

func (s *UrgentContext) GetParser() antlr.Parser { return s.parser }

func (s *UrgentContext) Word_() IWord_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWord_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWord_Context)
}

func (s *UrgentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UrgentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UrgentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterUrgent(s)
	}
}

func (s *UrgentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitUrgent(s)
	}
}

func (p *tcpParser) Urgent() (localctx IUrgentContext) {
	this := p
	_ = this

	localctx = NewUrgentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tcpParserRULE_urgent)

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
		p.Word_()
	}

	return localctx
}

// IDword_Context is an interface to support dynamic dispatch.
type IDword_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDword_Context differentiates from other interfaces.
	IsDword_Context()
}

type Dword_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDword_Context() *Dword_Context {
	var p = new(Dword_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_dword_
	return p
}

func (*Dword_Context) IsDword_Context() {}

func NewDword_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dword_Context {
	var p = new(Dword_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_dword_

	return p
}

func (s *Dword_Context) GetParser() antlr.Parser { return s.parser }

func (s *Dword_Context) AllBYTE() []antlr.TerminalNode {
	return s.GetTokens(tcpParserBYTE)
}

func (s *Dword_Context) BYTE(i int) antlr.TerminalNode {
	return s.GetToken(tcpParserBYTE, i)
}

func (s *Dword_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dword_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dword_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterDword_(s)
	}
}

func (s *Dword_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitDword_(s)
	}
}

func (p *tcpParser) Dword_() (localctx IDword_Context) {
	this := p
	_ = this

	localctx = NewDword_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tcpParserRULE_dword_)

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
		p.Match(tcpParserBYTE)
	}
	{
		p.SetState(50)
		p.Match(tcpParserBYTE)
	}
	{
		p.SetState(51)
		p.Match(tcpParserBYTE)
	}
	{
		p.SetState(52)
		p.Match(tcpParserBYTE)
	}

	return localctx
}

// IWord_Context is an interface to support dynamic dispatch.
type IWord_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWord_Context differentiates from other interfaces.
	IsWord_Context()
}

type Word_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWord_Context() *Word_Context {
	var p = new(Word_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_word_
	return p
}

func (*Word_Context) IsWord_Context() {}

func NewWord_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Word_Context {
	var p = new(Word_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_word_

	return p
}

func (s *Word_Context) GetParser() antlr.Parser { return s.parser }

func (s *Word_Context) AllBYTE() []antlr.TerminalNode {
	return s.GetTokens(tcpParserBYTE)
}

func (s *Word_Context) BYTE(i int) antlr.TerminalNode {
	return s.GetToken(tcpParserBYTE, i)
}

func (s *Word_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Word_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Word_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterWord_(s)
	}
}

func (s *Word_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitWord_(s)
	}
}

func (p *tcpParser) Word_() (localctx IWord_Context) {
	this := p
	_ = this

	localctx = NewWord_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tcpParserRULE_word_)

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
		p.SetState(54)
		p.Match(tcpParserBYTE)
	}
	{
		p.SetState(55)
		p.Match(tcpParserBYTE)
	}

	return localctx
}

// IByte_Context is an interface to support dynamic dispatch.
type IByte_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsByte_Context differentiates from other interfaces.
	IsByte_Context()
}

type Byte_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyByte_Context() *Byte_Context {
	var p = new(Byte_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tcpParserRULE_byte_
	return p
}

func (*Byte_Context) IsByte_Context() {}

func NewByte_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Byte_Context {
	var p = new(Byte_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tcpParserRULE_byte_

	return p
}

func (s *Byte_Context) GetParser() antlr.Parser { return s.parser }

func (s *Byte_Context) BYTE() antlr.TerminalNode {
	return s.GetToken(tcpParserBYTE, 0)
}

func (s *Byte_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Byte_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Byte_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.EnterByte_(s)
	}
}

func (s *Byte_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tcpListener); ok {
		listenerT.ExitByte_(s)
	}
}

func (p *tcpParser) Byte_() (localctx IByte_Context) {
	this := p
	_ = this

	localctx = NewByte_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, tcpParserRULE_byte_)

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
		p.SetState(57)
		p.Match(tcpParserBYTE)
	}

	return localctx
}
