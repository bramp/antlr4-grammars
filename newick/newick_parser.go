// Code generated from newick.g4 by ANTLR 4.9.3. DO NOT EDIT.

package newick // newick
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 77, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 32, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 40, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 5, 5, 51, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 65, 10, 8, 3, 9, 3, 9, 3, 9, 3, 10,
	5, 10, 71, 10, 10, 3, 11, 3, 11, 5, 11, 75, 10, 11, 3, 11, 2, 2, 12, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 2, 2, 73, 2, 31, 3, 2, 2, 2, 4, 39,
	3, 2, 2, 2, 6, 41, 3, 2, 2, 2, 8, 50, 3, 2, 2, 2, 10, 52, 3, 2, 2, 2, 12,
	54, 3, 2, 2, 2, 14, 64, 3, 2, 2, 2, 16, 66, 3, 2, 2, 2, 18, 70, 3, 2, 2,
	2, 20, 74, 3, 2, 2, 2, 22, 23, 5, 4, 3, 2, 23, 24, 7, 3, 2, 2, 24, 32,
	3, 2, 2, 2, 25, 26, 5, 6, 4, 2, 26, 27, 7, 3, 2, 2, 27, 32, 3, 2, 2, 2,
	28, 29, 5, 16, 9, 2, 29, 30, 7, 3, 2, 2, 30, 32, 3, 2, 2, 2, 31, 22, 3,
	2, 2, 2, 31, 25, 3, 2, 2, 2, 31, 28, 3, 2, 2, 2, 32, 3, 3, 2, 2, 2, 33,
	40, 5, 18, 10, 2, 34, 35, 7, 4, 2, 2, 35, 36, 5, 16, 9, 2, 36, 37, 7, 5,
	2, 2, 37, 38, 5, 18, 10, 2, 38, 40, 3, 2, 2, 2, 39, 33, 3, 2, 2, 2, 39,
	34, 3, 2, 2, 2, 40, 5, 3, 2, 2, 2, 41, 42, 7, 4, 2, 2, 42, 43, 5, 16, 9,
	2, 43, 44, 7, 6, 2, 2, 44, 45, 5, 14, 8, 2, 45, 46, 7, 5, 2, 2, 46, 47,
	5, 18, 10, 2, 47, 7, 3, 2, 2, 2, 48, 51, 5, 10, 6, 2, 49, 51, 5, 12, 7,
	2, 50, 48, 3, 2, 2, 2, 50, 49, 3, 2, 2, 2, 51, 9, 3, 2, 2, 2, 52, 53, 5,
	18, 10, 2, 53, 11, 3, 2, 2, 2, 54, 55, 7, 4, 2, 2, 55, 56, 5, 14, 8, 2,
	56, 57, 7, 5, 2, 2, 57, 58, 5, 18, 10, 2, 58, 13, 3, 2, 2, 2, 59, 65, 5,
	16, 9, 2, 60, 61, 5, 16, 9, 2, 61, 62, 7, 6, 2, 2, 62, 63, 5, 14, 8, 2,
	63, 65, 3, 2, 2, 2, 64, 59, 3, 2, 2, 2, 64, 60, 3, 2, 2, 2, 65, 15, 3,
	2, 2, 2, 66, 67, 5, 8, 5, 2, 67, 68, 5, 20, 11, 2, 68, 17, 3, 2, 2, 2,
	69, 71, 7, 9, 2, 2, 70, 69, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 19, 3,
	2, 2, 2, 72, 73, 7, 7, 2, 2, 73, 75, 7, 8, 2, 2, 74, 72, 3, 2, 2, 2, 74,
	75, 3, 2, 2, 2, 75, 21, 3, 2, 2, 2, 8, 31, 39, 50, 64, 70, 74,
}
var literalNames = []string{
	"", "';'", "'('", "')'", "','", "':'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "NUMBER", "STRING", "WS",
}

var ruleNames = []string{
	"tree_", "rootLeaf", "rootInternal", "subtree", "leaf", "internal_", "branchSet",
	"branch", "name", "length",
}

type newickParser struct {
	*antlr.BaseParser
}

// NewnewickParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *newickParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewnewickParser(input antlr.TokenStream) *newickParser {
	this := new(newickParser)
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
	this.GrammarFileName = "newick.g4"

	return this
}

// newickParser tokens.
const (
	newickParserEOF    = antlr.TokenEOF
	newickParserT__0   = 1
	newickParserT__1   = 2
	newickParserT__2   = 3
	newickParserT__3   = 4
	newickParserT__4   = 5
	newickParserNUMBER = 6
	newickParserSTRING = 7
	newickParserWS     = 8
)

// newickParser rules.
const (
	newickParserRULE_tree_        = 0
	newickParserRULE_rootLeaf     = 1
	newickParserRULE_rootInternal = 2
	newickParserRULE_subtree      = 3
	newickParserRULE_leaf         = 4
	newickParserRULE_internal_    = 5
	newickParserRULE_branchSet    = 6
	newickParserRULE_branch       = 7
	newickParserRULE_name         = 8
	newickParserRULE_length       = 9
)

// ITree_Context is an interface to support dynamic dispatch.
type ITree_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTree_Context differentiates from other interfaces.
	IsTree_Context()
}

type Tree_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTree_Context() *Tree_Context {
	var p = new(Tree_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = newickParserRULE_tree_
	return p
}

func (*Tree_Context) IsTree_Context() {}

func NewTree_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tree_Context {
	var p = new(Tree_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = newickParserRULE_tree_

	return p
}

func (s *Tree_Context) GetParser() antlr.Parser { return s.parser }

func (s *Tree_Context) RootLeaf() IRootLeafContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRootLeafContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRootLeafContext)
}

func (s *Tree_Context) RootInternal() IRootInternalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRootInternalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRootInternalContext)
}

func (s *Tree_Context) Branch() IBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *Tree_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tree_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tree_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.EnterTree_(s)
	}
}

func (s *Tree_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.ExitTree_(s)
	}
}

func (p *newickParser) Tree_() (localctx ITree_Context) {
	this := p
	_ = this

	localctx = NewTree_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, newickParserRULE_tree_)

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

	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.RootLeaf()
		}
		{
			p.SetState(21)
			p.Match(newickParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.RootInternal()
		}
		{
			p.SetState(24)
			p.Match(newickParserT__0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Branch()
		}
		{
			p.SetState(27)
			p.Match(newickParserT__0)
		}

	}

	return localctx
}

// IRootLeafContext is an interface to support dynamic dispatch.
type IRootLeafContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootLeafContext differentiates from other interfaces.
	IsRootLeafContext()
}

type RootLeafContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootLeafContext() *RootLeafContext {
	var p = new(RootLeafContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = newickParserRULE_rootLeaf
	return p
}

func (*RootLeafContext) IsRootLeafContext() {}

func NewRootLeafContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootLeafContext {
	var p = new(RootLeafContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = newickParserRULE_rootLeaf

	return p
}

func (s *RootLeafContext) GetParser() antlr.Parser { return s.parser }

func (s *RootLeafContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RootLeafContext) Branch() IBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *RootLeafContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootLeafContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootLeafContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.EnterRootLeaf(s)
	}
}

func (s *RootLeafContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.ExitRootLeaf(s)
	}
}

func (p *newickParser) RootLeaf() (localctx IRootLeafContext) {
	this := p
	_ = this

	localctx = NewRootLeafContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, newickParserRULE_rootLeaf)

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

	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case newickParserT__0, newickParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Name()
		}

	case newickParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Match(newickParserT__1)
		}
		{
			p.SetState(33)
			p.Branch()
		}
		{
			p.SetState(34)
			p.Match(newickParserT__2)
		}
		{
			p.SetState(35)
			p.Name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRootInternalContext is an interface to support dynamic dispatch.
type IRootInternalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootInternalContext differentiates from other interfaces.
	IsRootInternalContext()
}

type RootInternalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootInternalContext() *RootInternalContext {
	var p = new(RootInternalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = newickParserRULE_rootInternal
	return p
}

func (*RootInternalContext) IsRootInternalContext() {}

func NewRootInternalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootInternalContext {
	var p = new(RootInternalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = newickParserRULE_rootInternal

	return p
}

func (s *RootInternalContext) GetParser() antlr.Parser { return s.parser }

func (s *RootInternalContext) Branch() IBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *RootInternalContext) BranchSet() IBranchSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchSetContext)
}

func (s *RootInternalContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RootInternalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootInternalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootInternalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.EnterRootInternal(s)
	}
}

func (s *RootInternalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.ExitRootInternal(s)
	}
}

func (p *newickParser) RootInternal() (localctx IRootInternalContext) {
	this := p
	_ = this

	localctx = NewRootInternalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, newickParserRULE_rootInternal)

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
		p.Match(newickParserT__1)
	}
	{
		p.SetState(40)
		p.Branch()
	}
	{
		p.SetState(41)
		p.Match(newickParserT__3)
	}
	{
		p.SetState(42)
		p.BranchSet()
	}
	{
		p.SetState(43)
		p.Match(newickParserT__2)
	}
	{
		p.SetState(44)
		p.Name()
	}

	return localctx
}

// ISubtreeContext is an interface to support dynamic dispatch.
type ISubtreeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubtreeContext differentiates from other interfaces.
	IsSubtreeContext()
}

type SubtreeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubtreeContext() *SubtreeContext {
	var p = new(SubtreeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = newickParserRULE_subtree
	return p
}

func (*SubtreeContext) IsSubtreeContext() {}

func NewSubtreeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubtreeContext {
	var p = new(SubtreeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = newickParserRULE_subtree

	return p
}

func (s *SubtreeContext) GetParser() antlr.Parser { return s.parser }

func (s *SubtreeContext) Leaf() ILeafContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeafContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeafContext)
}

func (s *SubtreeContext) Internal_() IInternal_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInternal_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInternal_Context)
}

func (s *SubtreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtreeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubtreeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.EnterSubtree(s)
	}
}

func (s *SubtreeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.ExitSubtree(s)
	}
}

func (p *newickParser) Subtree() (localctx ISubtreeContext) {
	this := p
	_ = this

	localctx = NewSubtreeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, newickParserRULE_subtree)

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

	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case newickParserT__0, newickParserT__2, newickParserT__3, newickParserT__4, newickParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Leaf()
		}

	case newickParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Internal_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILeafContext is an interface to support dynamic dispatch.
type ILeafContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeafContext differentiates from other interfaces.
	IsLeafContext()
}

type LeafContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeafContext() *LeafContext {
	var p = new(LeafContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = newickParserRULE_leaf
	return p
}

func (*LeafContext) IsLeafContext() {}

func NewLeafContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeafContext {
	var p = new(LeafContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = newickParserRULE_leaf

	return p
}

func (s *LeafContext) GetParser() antlr.Parser { return s.parser }

func (s *LeafContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LeafContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeafContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeafContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.EnterLeaf(s)
	}
}

func (s *LeafContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.ExitLeaf(s)
	}
}

func (p *newickParser) Leaf() (localctx ILeafContext) {
	this := p
	_ = this

	localctx = NewLeafContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, newickParserRULE_leaf)

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
		p.SetState(50)
		p.Name()
	}

	return localctx
}

// IInternal_Context is an interface to support dynamic dispatch.
type IInternal_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInternal_Context differentiates from other interfaces.
	IsInternal_Context()
}

type Internal_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInternal_Context() *Internal_Context {
	var p = new(Internal_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = newickParserRULE_internal_
	return p
}

func (*Internal_Context) IsInternal_Context() {}

func NewInternal_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Internal_Context {
	var p = new(Internal_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = newickParserRULE_internal_

	return p
}

func (s *Internal_Context) GetParser() antlr.Parser { return s.parser }

func (s *Internal_Context) BranchSet() IBranchSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchSetContext)
}

func (s *Internal_Context) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Internal_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Internal_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Internal_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.EnterInternal_(s)
	}
}

func (s *Internal_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.ExitInternal_(s)
	}
}

func (p *newickParser) Internal_() (localctx IInternal_Context) {
	this := p
	_ = this

	localctx = NewInternal_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, newickParserRULE_internal_)

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
		p.Match(newickParserT__1)
	}
	{
		p.SetState(53)
		p.BranchSet()
	}
	{
		p.SetState(54)
		p.Match(newickParserT__2)
	}
	{
		p.SetState(55)
		p.Name()
	}

	return localctx
}

// IBranchSetContext is an interface to support dynamic dispatch.
type IBranchSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBranchSetContext differentiates from other interfaces.
	IsBranchSetContext()
}

type BranchSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchSetContext() *BranchSetContext {
	var p = new(BranchSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = newickParserRULE_branchSet
	return p
}

func (*BranchSetContext) IsBranchSetContext() {}

func NewBranchSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchSetContext {
	var p = new(BranchSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = newickParserRULE_branchSet

	return p
}

func (s *BranchSetContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchSetContext) Branch() IBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *BranchSetContext) BranchSet() IBranchSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranchSetContext)
}

func (s *BranchSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BranchSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.EnterBranchSet(s)
	}
}

func (s *BranchSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.ExitBranchSet(s)
	}
}

func (p *newickParser) BranchSet() (localctx IBranchSetContext) {
	this := p
	_ = this

	localctx = NewBranchSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, newickParserRULE_branchSet)

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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Branch()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Branch()
		}
		{
			p.SetState(59)
			p.Match(newickParserT__3)
		}
		{
			p.SetState(60)
			p.BranchSet()
		}

	}

	return localctx
}

// IBranchContext is an interface to support dynamic dispatch.
type IBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBranchContext differentiates from other interfaces.
	IsBranchContext()
}

type BranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchContext() *BranchContext {
	var p = new(BranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = newickParserRULE_branch
	return p
}

func (*BranchContext) IsBranchContext() {}

func NewBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchContext {
	var p = new(BranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = newickParserRULE_branch

	return p
}

func (s *BranchContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchContext) Subtree() ISubtreeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubtreeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubtreeContext)
}

func (s *BranchContext) Length() ILengthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthContext)
}

func (s *BranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.EnterBranch(s)
	}
}

func (s *BranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.ExitBranch(s)
	}
}

func (p *newickParser) Branch() (localctx IBranchContext) {
	this := p
	_ = this

	localctx = NewBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, newickParserRULE_branch)

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
		p.Subtree()
	}
	{
		p.SetState(65)
		p.Length()
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = newickParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = newickParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(newickParserSTRING, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *newickParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, newickParserRULE_name)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == newickParserSTRING {
		{
			p.SetState(67)
			p.Match(newickParserSTRING)
		}

	}

	return localctx
}

// ILengthContext is an interface to support dynamic dispatch.
type ILengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthContext differentiates from other interfaces.
	IsLengthContext()
}

type LengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthContext() *LengthContext {
	var p = new(LengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = newickParserRULE_length
	return p
}

func (*LengthContext) IsLengthContext() {}

func NewLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthContext {
	var p = new(LengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = newickParserRULE_length

	return p
}

func (s *LengthContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(newickParserNUMBER, 0)
}

func (s *LengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.EnterLength(s)
	}
}

func (s *LengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(newickListener); ok {
		listenerT.ExitLength(s)
	}
}

func (p *newickParser) Length() (localctx ILengthContext) {
	this := p
	_ = this

	localctx = NewLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, newickParserRULE_length)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == newickParserT__4 {
		{
			p.SetState(70)
			p.Match(newickParserT__4)
		}
		{
			p.SetState(71)
			p.Match(newickParserNUMBER)
		}

	}

	return localctx
}
