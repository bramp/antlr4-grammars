// Code generated from vmf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package vmf // vmf
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 48, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 6, 2, 16, 10, 2, 13, 2, 14, 2, 17, 3, 3, 3, 3, 6, 3, 22, 10, 3, 13,
	3, 14, 3, 23, 3, 3, 6, 3, 27, 10, 3, 13, 3, 14, 3, 28, 5, 3, 31, 10, 3,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 7, 7, 41, 10, 7, 12, 7,
	14, 7, 44, 11, 7, 3, 7, 3, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10, 12, 2, 3,
	3, 2, 5, 6, 2, 46, 2, 15, 3, 2, 2, 2, 4, 19, 3, 2, 2, 2, 6, 32, 3, 2, 2,
	2, 8, 34, 3, 2, 2, 2, 10, 36, 3, 2, 2, 2, 12, 38, 3, 2, 2, 2, 14, 16, 5,
	4, 3, 2, 15, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17,
	18, 3, 2, 2, 2, 18, 3, 3, 2, 2, 2, 19, 30, 5, 6, 4, 2, 20, 22, 5, 8, 5,
	2, 21, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24,
	3, 2, 2, 2, 24, 31, 3, 2, 2, 2, 25, 27, 5, 12, 7, 2, 26, 25, 3, 2, 2, 2,
	27, 28, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 31, 3,
	2, 2, 2, 30, 21, 3, 2, 2, 2, 30, 26, 3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32,
	33, 5, 10, 6, 2, 33, 7, 3, 2, 2, 2, 34, 35, 5, 10, 6, 2, 35, 9, 3, 2, 2,
	2, 36, 37, 9, 2, 2, 2, 37, 11, 3, 2, 2, 2, 38, 42, 7, 3, 2, 2, 39, 41,
	5, 4, 3, 2, 40, 39, 3, 2, 2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2,
	42, 43, 3, 2, 2, 2, 43, 45, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 46, 7,
	4, 2, 2, 46, 13, 3, 2, 2, 2, 7, 17, 23, 28, 30, 42,
}
var literalNames = []string{
	"", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "", "", "QUOTEDSTTRING", "STRING", "WS",
}

var ruleNames = []string{
	"vmf", "keyvalue", "key", "atomicvalue", "val", "listvalue",
}

type vmfParser struct {
	*antlr.BaseParser
}

// NewvmfParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *vmfParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewvmfParser(input antlr.TokenStream) *vmfParser {
	this := new(vmfParser)
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
	this.GrammarFileName = "vmf.g4"

	return this
}

// vmfParser tokens.
const (
	vmfParserEOF           = antlr.TokenEOF
	vmfParserT__0          = 1
	vmfParserT__1          = 2
	vmfParserQUOTEDSTTRING = 3
	vmfParserSTRING        = 4
	vmfParserWS            = 5
)

// vmfParser rules.
const (
	vmfParserRULE_vmf         = 0
	vmfParserRULE_keyvalue    = 1
	vmfParserRULE_key         = 2
	vmfParserRULE_atomicvalue = 3
	vmfParserRULE_val         = 4
	vmfParserRULE_listvalue   = 5
)

// IVmfContext is an interface to support dynamic dispatch.
type IVmfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVmfContext differentiates from other interfaces.
	IsVmfContext()
}

type VmfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVmfContext() *VmfContext {
	var p = new(VmfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vmfParserRULE_vmf
	return p
}

func (*VmfContext) IsVmfContext() {}

func NewVmfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VmfContext {
	var p = new(VmfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vmfParserRULE_vmf

	return p
}

func (s *VmfContext) GetParser() antlr.Parser { return s.parser }

func (s *VmfContext) AllKeyvalue() []IKeyvalueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyvalueContext)(nil)).Elem())
	var tst = make([]IKeyvalueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyvalueContext)
		}
	}

	return tst
}

func (s *VmfContext) Keyvalue(i int) IKeyvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyvalueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyvalueContext)
}

func (s *VmfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VmfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VmfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.EnterVmf(s)
	}
}

func (s *VmfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.ExitVmf(s)
	}
}

func (p *vmfParser) Vmf() (localctx IVmfContext) {
	this := p
	_ = this

	localctx = NewVmfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, vmfParserRULE_vmf)
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

	for ok := true; ok; ok = _la == vmfParserQUOTEDSTTRING || _la == vmfParserSTRING {
		{
			p.SetState(12)
			p.Keyvalue()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKeyvalueContext is an interface to support dynamic dispatch.
type IKeyvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyvalueContext differentiates from other interfaces.
	IsKeyvalueContext()
}

type KeyvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyvalueContext() *KeyvalueContext {
	var p = new(KeyvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vmfParserRULE_keyvalue
	return p
}

func (*KeyvalueContext) IsKeyvalueContext() {}

func NewKeyvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyvalueContext {
	var p = new(KeyvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vmfParserRULE_keyvalue

	return p
}

func (s *KeyvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyvalueContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *KeyvalueContext) AllAtomicvalue() []IAtomicvalueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomicvalueContext)(nil)).Elem())
	var tst = make([]IAtomicvalueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomicvalueContext)
		}
	}

	return tst
}

func (s *KeyvalueContext) Atomicvalue(i int) IAtomicvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomicvalueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomicvalueContext)
}

func (s *KeyvalueContext) AllListvalue() []IListvalueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IListvalueContext)(nil)).Elem())
	var tst = make([]IListvalueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IListvalueContext)
		}
	}

	return tst
}

func (s *KeyvalueContext) Listvalue(i int) IListvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListvalueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IListvalueContext)
}

func (s *KeyvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.EnterKeyvalue(s)
	}
}

func (s *KeyvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.ExitKeyvalue(s)
	}
}

func (p *vmfParser) Keyvalue() (localctx IKeyvalueContext) {
	this := p
	_ = this

	localctx = NewKeyvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, vmfParserRULE_keyvalue)
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
	{
		p.SetState(17)
		p.Key()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case vmfParserQUOTEDSTTRING, vmfParserSTRING:
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(18)
					p.Atomicvalue()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(21)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}

	case vmfParserT__0:
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == vmfParserT__0 {
			{
				p.SetState(23)
				p.Listvalue()
			}

			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vmfParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vmfParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *vmfParser) Key() (localctx IKeyContext) {
	this := p
	_ = this

	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, vmfParserRULE_key)

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
		p.SetState(30)
		p.Val()
	}

	return localctx
}

// IAtomicvalueContext is an interface to support dynamic dispatch.
type IAtomicvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomicvalueContext differentiates from other interfaces.
	IsAtomicvalueContext()
}

type AtomicvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomicvalueContext() *AtomicvalueContext {
	var p = new(AtomicvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vmfParserRULE_atomicvalue
	return p
}

func (*AtomicvalueContext) IsAtomicvalueContext() {}

func NewAtomicvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomicvalueContext {
	var p = new(AtomicvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vmfParserRULE_atomicvalue

	return p
}

func (s *AtomicvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomicvalueContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *AtomicvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomicvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.EnterAtomicvalue(s)
	}
}

func (s *AtomicvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.ExitAtomicvalue(s)
	}
}

func (p *vmfParser) Atomicvalue() (localctx IAtomicvalueContext) {
	this := p
	_ = this

	localctx = NewAtomicvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, vmfParserRULE_atomicvalue)

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
		p.Val()
	}

	return localctx
}

// IValContext is an interface to support dynamic dispatch.
type IValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValContext differentiates from other interfaces.
	IsValContext()
}

type ValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValContext() *ValContext {
	var p = new(ValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vmfParserRULE_val
	return p
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vmfParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) QUOTEDSTTRING() antlr.TerminalNode {
	return s.GetToken(vmfParserQUOTEDSTTRING, 0)
}

func (s *ValContext) STRING() antlr.TerminalNode {
	return s.GetToken(vmfParserSTRING, 0)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.EnterVal(s)
	}
}

func (s *ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.ExitVal(s)
	}
}

func (p *vmfParser) Val() (localctx IValContext) {
	this := p
	_ = this

	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, vmfParserRULE_val)
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
		p.SetState(34)
		_la = p.GetTokenStream().LA(1)

		if !(_la == vmfParserQUOTEDSTTRING || _la == vmfParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IListvalueContext is an interface to support dynamic dispatch.
type IListvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListvalueContext differentiates from other interfaces.
	IsListvalueContext()
}

type ListvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListvalueContext() *ListvalueContext {
	var p = new(ListvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vmfParserRULE_listvalue
	return p
}

func (*ListvalueContext) IsListvalueContext() {}

func NewListvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListvalueContext {
	var p = new(ListvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vmfParserRULE_listvalue

	return p
}

func (s *ListvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *ListvalueContext) AllKeyvalue() []IKeyvalueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyvalueContext)(nil)).Elem())
	var tst = make([]IKeyvalueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyvalueContext)
		}
	}

	return tst
}

func (s *ListvalueContext) Keyvalue(i int) IKeyvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyvalueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyvalueContext)
}

func (s *ListvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.EnterListvalue(s)
	}
}

func (s *ListvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vmfListener); ok {
		listenerT.ExitListvalue(s)
	}
}

func (p *vmfParser) Listvalue() (localctx IListvalueContext) {
	this := p
	_ = this

	localctx = NewListvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, vmfParserRULE_listvalue)
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
		p.SetState(36)
		p.Match(vmfParserT__0)
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == vmfParserQUOTEDSTTRING || _la == vmfParserSTRING {
		{
			p.SetState(37)
			p.Keyvalue()
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(43)
		p.Match(vmfParserT__1)
	}

	return localctx
}
