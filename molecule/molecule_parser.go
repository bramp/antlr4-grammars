// Code generated from molecule.g4 by ANTLR 4.9.3. DO NOT EDIT.

package molecule // molecule
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 64, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 7, 2, 22, 10, 2, 12, 2, 14, 2, 25,
	11, 2, 3, 3, 5, 3, 28, 10, 3, 3, 3, 6, 3, 31, 10, 3, 13, 3, 14, 3, 32,
	3, 4, 3, 4, 5, 4, 37, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 42, 10, 5, 3, 6, 3,
	6, 6, 6, 46, 10, 6, 13, 6, 14, 6, 47, 3, 6, 3, 6, 3, 7, 3, 7, 6, 7, 54,
	10, 7, 13, 7, 14, 7, 55, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 2, 2,
	10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 2, 2, 63, 2, 18, 3, 2, 2, 2, 4, 30,
	3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8, 41, 3, 2, 2, 2, 10, 43, 3, 2, 2, 2, 12,
	51, 3, 2, 2, 2, 14, 59, 3, 2, 2, 2, 16, 61, 3, 2, 2, 2, 18, 23, 5, 4, 3,
	2, 19, 20, 7, 3, 2, 2, 20, 22, 5, 4, 3, 2, 21, 19, 3, 2, 2, 2, 22, 25,
	3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 3, 3, 2, 2, 2,
	25, 23, 3, 2, 2, 2, 26, 28, 5, 16, 9, 2, 27, 26, 3, 2, 2, 2, 27, 28, 3,
	2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 31, 5, 6, 4, 2, 30, 27, 3, 2, 2, 2, 31,
	32, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 5, 3, 2, 2,
	2, 34, 36, 5, 8, 5, 2, 35, 37, 5, 16, 9, 2, 36, 35, 3, 2, 2, 2, 36, 37,
	3, 2, 2, 2, 37, 7, 3, 2, 2, 2, 38, 42, 5, 14, 8, 2, 39, 42, 5, 10, 6, 2,
	40, 42, 5, 12, 7, 2, 41, 38, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 40, 3,
	2, 2, 2, 42, 9, 3, 2, 2, 2, 43, 45, 7, 4, 2, 2, 44, 46, 5, 6, 4, 2, 45,
	44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2,
	2, 48, 49, 3, 2, 2, 2, 49, 50, 7, 5, 2, 2, 50, 11, 3, 2, 2, 2, 51, 53,
	7, 6, 2, 2, 52, 54, 5, 6, 4, 2, 53, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2,
	55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 58, 7,
	7, 2, 2, 58, 13, 3, 2, 2, 2, 59, 60, 7, 8, 2, 2, 60, 15, 3, 2, 2, 2, 61,
	62, 7, 9, 2, 2, 62, 17, 3, 2, 2, 2, 9, 23, 27, 32, 36, 41, 47, 55,
}
var literalNames = []string{
	"", "'\u00B7'", "'('", "')'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "ELEMENT", "NUMBER", "WS",
}

var ruleNames = []string{
	"molecule", "part_", "structure", "symbol", "group", "ion", "element",
	"count",
}

type moleculeParser struct {
	*antlr.BaseParser
}

// NewmoleculeParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *moleculeParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewmoleculeParser(input antlr.TokenStream) *moleculeParser {
	this := new(moleculeParser)
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
	this.GrammarFileName = "molecule.g4"

	return this
}

// moleculeParser tokens.
const (
	moleculeParserEOF     = antlr.TokenEOF
	moleculeParserT__0    = 1
	moleculeParserT__1    = 2
	moleculeParserT__2    = 3
	moleculeParserT__3    = 4
	moleculeParserT__4    = 5
	moleculeParserELEMENT = 6
	moleculeParserNUMBER  = 7
	moleculeParserWS      = 8
)

// moleculeParser rules.
const (
	moleculeParserRULE_molecule  = 0
	moleculeParserRULE_part_     = 1
	moleculeParserRULE_structure = 2
	moleculeParserRULE_symbol    = 3
	moleculeParserRULE_group     = 4
	moleculeParserRULE_ion       = 5
	moleculeParserRULE_element   = 6
	moleculeParserRULE_count     = 7
)

// IMoleculeContext is an interface to support dynamic dispatch.
type IMoleculeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoleculeContext differentiates from other interfaces.
	IsMoleculeContext()
}

type MoleculeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoleculeContext() *MoleculeContext {
	var p = new(MoleculeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = moleculeParserRULE_molecule
	return p
}

func (*MoleculeContext) IsMoleculeContext() {}

func NewMoleculeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoleculeContext {
	var p = new(MoleculeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = moleculeParserRULE_molecule

	return p
}

func (s *MoleculeContext) GetParser() antlr.Parser { return s.parser }

func (s *MoleculeContext) AllPart_() []IPart_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPart_Context)(nil)).Elem())
	var tst = make([]IPart_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPart_Context)
		}
	}

	return tst
}

func (s *MoleculeContext) Part_(i int) IPart_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPart_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPart_Context)
}

func (s *MoleculeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoleculeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoleculeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.EnterMolecule(s)
	}
}

func (s *MoleculeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.ExitMolecule(s)
	}
}

func (p *moleculeParser) Molecule() (localctx IMoleculeContext) {
	this := p
	_ = this

	localctx = NewMoleculeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, moleculeParserRULE_molecule)
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
		p.Part_()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == moleculeParserT__0 {
		{
			p.SetState(17)
			p.Match(moleculeParserT__0)
		}
		{
			p.SetState(18)
			p.Part_()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPart_Context is an interface to support dynamic dispatch.
type IPart_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPart_Context differentiates from other interfaces.
	IsPart_Context()
}

type Part_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPart_Context() *Part_Context {
	var p = new(Part_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = moleculeParserRULE_part_
	return p
}

func (*Part_Context) IsPart_Context() {}

func NewPart_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Part_Context {
	var p = new(Part_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = moleculeParserRULE_part_

	return p
}

func (s *Part_Context) GetParser() antlr.Parser { return s.parser }

func (s *Part_Context) AllStructure() []IStructureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructureContext)(nil)).Elem())
	var tst = make([]IStructureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructureContext)
		}
	}

	return tst
}

func (s *Part_Context) Structure(i int) IStructureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructureContext)
}

func (s *Part_Context) AllCount() []ICountContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICountContext)(nil)).Elem())
	var tst = make([]ICountContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICountContext)
		}
	}

	return tst
}

func (s *Part_Context) Count(i int) ICountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICountContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICountContext)
}

func (s *Part_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Part_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Part_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.EnterPart_(s)
	}
}

func (s *Part_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.ExitPart_(s)
	}
}

func (p *moleculeParser) Part_() (localctx IPart_Context) {
	this := p
	_ = this

	localctx = NewPart_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, moleculeParserRULE_part_)
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
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<moleculeParserT__1)|(1<<moleculeParserT__3)|(1<<moleculeParserELEMENT)|(1<<moleculeParserNUMBER))) != 0) {
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == moleculeParserNUMBER {
			{
				p.SetState(24)
				p.Count()
			}

		}
		{
			p.SetState(27)
			p.Structure()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStructureContext is an interface to support dynamic dispatch.
type IStructureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructureContext differentiates from other interfaces.
	IsStructureContext()
}

type StructureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructureContext() *StructureContext {
	var p = new(StructureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = moleculeParserRULE_structure
	return p
}

func (*StructureContext) IsStructureContext() {}

func NewStructureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructureContext {
	var p = new(StructureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = moleculeParserRULE_structure

	return p
}

func (s *StructureContext) GetParser() antlr.Parser { return s.parser }

func (s *StructureContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *StructureContext) Count() ICountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICountContext)
}

func (s *StructureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.EnterStructure(s)
	}
}

func (s *StructureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.ExitStructure(s)
	}
}

func (p *moleculeParser) Structure() (localctx IStructureContext) {
	this := p
	_ = this

	localctx = NewStructureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, moleculeParserRULE_structure)

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
		p.Symbol()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(33)
			p.Count()
		}

	}

	return localctx
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = moleculeParserRULE_symbol
	return p
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = moleculeParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *SymbolContext) Group() IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *SymbolContext) Ion() IIonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIonContext)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (p *moleculeParser) Symbol() (localctx ISymbolContext) {
	this := p
	_ = this

	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, moleculeParserRULE_symbol)

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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case moleculeParserELEMENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Element()
		}

	case moleculeParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Group()
		}

	case moleculeParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.Ion()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = moleculeParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = moleculeParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) AllStructure() []IStructureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructureContext)(nil)).Elem())
	var tst = make([]IStructureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructureContext)
		}
	}

	return tst
}

func (s *GroupContext) Structure(i int) IStructureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructureContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *moleculeParser) Group() (localctx IGroupContext) {
	this := p
	_ = this

	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, moleculeParserRULE_group)
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
		p.SetState(41)
		p.Match(moleculeParserT__1)
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<moleculeParserT__1)|(1<<moleculeParserT__3)|(1<<moleculeParserELEMENT))) != 0) {
		{
			p.SetState(42)
			p.Structure()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(moleculeParserT__2)
	}

	return localctx
}

// IIonContext is an interface to support dynamic dispatch.
type IIonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIonContext differentiates from other interfaces.
	IsIonContext()
}

type IonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIonContext() *IonContext {
	var p = new(IonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = moleculeParserRULE_ion
	return p
}

func (*IonContext) IsIonContext() {}

func NewIonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IonContext {
	var p = new(IonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = moleculeParserRULE_ion

	return p
}

func (s *IonContext) GetParser() antlr.Parser { return s.parser }

func (s *IonContext) AllStructure() []IStructureContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructureContext)(nil)).Elem())
	var tst = make([]IStructureContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructureContext)
		}
	}

	return tst
}

func (s *IonContext) Structure(i int) IStructureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructureContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructureContext)
}

func (s *IonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.EnterIon(s)
	}
}

func (s *IonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.ExitIon(s)
	}
}

func (p *moleculeParser) Ion() (localctx IIonContext) {
	this := p
	_ = this

	localctx = NewIonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, moleculeParserRULE_ion)
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
		p.SetState(49)
		p.Match(moleculeParserT__3)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<moleculeParserT__1)|(1<<moleculeParserT__3)|(1<<moleculeParserELEMENT))) != 0) {
		{
			p.SetState(50)
			p.Structure()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Match(moleculeParserT__4)
	}

	return localctx
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = moleculeParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = moleculeParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) ELEMENT() antlr.TerminalNode {
	return s.GetToken(moleculeParserELEMENT, 0)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.ExitElement(s)
	}
}

func (p *moleculeParser) Element() (localctx IElementContext) {
	this := p
	_ = this

	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, moleculeParserRULE_element)

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
		p.Match(moleculeParserELEMENT)
	}

	return localctx
}

// ICountContext is an interface to support dynamic dispatch.
type ICountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCountContext differentiates from other interfaces.
	IsCountContext()
}

type CountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountContext() *CountContext {
	var p = new(CountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = moleculeParserRULE_count
	return p
}

func (*CountContext) IsCountContext() {}

func NewCountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountContext {
	var p = new(CountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = moleculeParserRULE_count

	return p
}

func (s *CountContext) GetParser() antlr.Parser { return s.parser }

func (s *CountContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(moleculeParserNUMBER, 0)
}

func (s *CountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.EnterCount(s)
	}
}

func (s *CountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(moleculeListener); ok {
		listenerT.ExitCount(s)
	}
}

func (p *moleculeParser) Count() (localctx ICountContext) {
	this := p
	_ = this

	localctx = NewCountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, moleculeParserRULE_count)

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
		p.SetState(59)
		p.Match(moleculeParserNUMBER)
	}

	return localctx
}
