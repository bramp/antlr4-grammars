// Code generated from fasta.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fasta // fasta
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 32, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 6, 2, 14,
	10, 2, 13, 2, 14, 2, 15, 3, 3, 3, 3, 3, 3, 5, 3, 21, 10, 3, 3, 4, 6, 4,
	24, 10, 4, 13, 4, 14, 4, 25, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 2, 2, 7, 2,
	4, 6, 8, 10, 2, 2, 2, 30, 2, 13, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 23,
	3, 2, 2, 2, 8, 27, 3, 2, 2, 2, 10, 29, 3, 2, 2, 2, 12, 14, 5, 4, 3, 2,
	13, 12, 3, 2, 2, 2, 14, 15, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 16, 3,
	2, 2, 2, 16, 3, 3, 2, 2, 2, 17, 21, 5, 8, 5, 2, 18, 21, 5, 6, 4, 2, 19,
	21, 5, 10, 6, 2, 20, 17, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 19, 3, 2,
	2, 2, 21, 5, 3, 2, 2, 2, 22, 24, 7, 7, 2, 2, 23, 22, 3, 2, 2, 2, 24, 25,
	3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 7, 3, 2, 2, 2,
	27, 28, 7, 4, 2, 2, 28, 9, 3, 2, 2, 2, 29, 30, 7, 3, 2, 2, 30, 11, 3, 2,
	2, 2, 5, 15, 20, 25,
}
var literalNames []string

var symbolicNames = []string{
	"", "COMMENTLINE", "DESCRIPTIONLINE", "TEXT", "EOL", "SEQUENCELINE",
}

var ruleNames = []string{
	"sequence", "section", "sequencelines", "descriptionline", "commentline",
}

type fastaParser struct {
	*antlr.BaseParser
}

// NewfastaParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *fastaParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewfastaParser(input antlr.TokenStream) *fastaParser {
	this := new(fastaParser)
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
	this.GrammarFileName = "fasta.g4"

	return this
}

// fastaParser tokens.
const (
	fastaParserEOF             = antlr.TokenEOF
	fastaParserCOMMENTLINE     = 1
	fastaParserDESCRIPTIONLINE = 2
	fastaParserTEXT            = 3
	fastaParserEOL             = 4
	fastaParserSEQUENCELINE    = 5
)

// fastaParser rules.
const (
	fastaParserRULE_sequence        = 0
	fastaParserRULE_section         = 1
	fastaParserRULE_sequencelines   = 2
	fastaParserRULE_descriptionline = 3
	fastaParserRULE_commentline     = 4
)

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fastaParserRULE_sequence
	return p
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceContext {
	var p = new(SequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fastaParserRULE_sequence

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) AllSection() []ISectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISectionContext)(nil)).Elem())
	var tst = make([]ISectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISectionContext)
		}
	}

	return tst
}

func (s *SequenceContext) Section(i int) ISectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fastaListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fastaListener); ok {
		listenerT.ExitSequence(s)
	}
}

func (p *fastaParser) Sequence() (localctx ISequenceContext) {
	this := p
	_ = this

	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, fastaParserRULE_sequence)
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
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<fastaParserCOMMENTLINE)|(1<<fastaParserDESCRIPTIONLINE)|(1<<fastaParserSEQUENCELINE))) != 0) {
		{
			p.SetState(10)
			p.Section()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISectionContext is an interface to support dynamic dispatch.
type ISectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionContext differentiates from other interfaces.
	IsSectionContext()
}

type SectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionContext() *SectionContext {
	var p = new(SectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fastaParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fastaParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) Descriptionline() IDescriptionlineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionlineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionlineContext)
}

func (s *SectionContext) Sequencelines() ISequencelinesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequencelinesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequencelinesContext)
}

func (s *SectionContext) Commentline() ICommentlineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentlineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentlineContext)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fastaListener); ok {
		listenerT.EnterSection(s)
	}
}

func (s *SectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fastaListener); ok {
		listenerT.ExitSection(s)
	}
}

func (p *fastaParser) Section() (localctx ISectionContext) {
	this := p
	_ = this

	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, fastaParserRULE_section)

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

	p.SetState(18)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case fastaParserDESCRIPTIONLINE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(15)
			p.Descriptionline()
		}

	case fastaParserSEQUENCELINE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(16)
			p.Sequencelines()
		}

	case fastaParserCOMMENTLINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(17)
			p.Commentline()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISequencelinesContext is an interface to support dynamic dispatch.
type ISequencelinesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequencelinesContext differentiates from other interfaces.
	IsSequencelinesContext()
}

type SequencelinesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequencelinesContext() *SequencelinesContext {
	var p = new(SequencelinesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fastaParserRULE_sequencelines
	return p
}

func (*SequencelinesContext) IsSequencelinesContext() {}

func NewSequencelinesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequencelinesContext {
	var p = new(SequencelinesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fastaParserRULE_sequencelines

	return p
}

func (s *SequencelinesContext) GetParser() antlr.Parser { return s.parser }

func (s *SequencelinesContext) AllSEQUENCELINE() []antlr.TerminalNode {
	return s.GetTokens(fastaParserSEQUENCELINE)
}

func (s *SequencelinesContext) SEQUENCELINE(i int) antlr.TerminalNode {
	return s.GetToken(fastaParserSEQUENCELINE, i)
}

func (s *SequencelinesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequencelinesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequencelinesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fastaListener); ok {
		listenerT.EnterSequencelines(s)
	}
}

func (s *SequencelinesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fastaListener); ok {
		listenerT.ExitSequencelines(s)
	}
}

func (p *fastaParser) Sequencelines() (localctx ISequencelinesContext) {
	this := p
	_ = this

	localctx = NewSequencelinesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, fastaParserRULE_sequencelines)

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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(20)
				p.Match(fastaParserSEQUENCELINE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IDescriptionlineContext is an interface to support dynamic dispatch.
type IDescriptionlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionlineContext differentiates from other interfaces.
	IsDescriptionlineContext()
}

type DescriptionlineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionlineContext() *DescriptionlineContext {
	var p = new(DescriptionlineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fastaParserRULE_descriptionline
	return p
}

func (*DescriptionlineContext) IsDescriptionlineContext() {}

func NewDescriptionlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionlineContext {
	var p = new(DescriptionlineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fastaParserRULE_descriptionline

	return p
}

func (s *DescriptionlineContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionlineContext) DESCRIPTIONLINE() antlr.TerminalNode {
	return s.GetToken(fastaParserDESCRIPTIONLINE, 0)
}

func (s *DescriptionlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fastaListener); ok {
		listenerT.EnterDescriptionline(s)
	}
}

func (s *DescriptionlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fastaListener); ok {
		listenerT.ExitDescriptionline(s)
	}
}

func (p *fastaParser) Descriptionline() (localctx IDescriptionlineContext) {
	this := p
	_ = this

	localctx = NewDescriptionlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, fastaParserRULE_descriptionline)

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
		p.SetState(25)
		p.Match(fastaParserDESCRIPTIONLINE)
	}

	return localctx
}

// ICommentlineContext is an interface to support dynamic dispatch.
type ICommentlineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentlineContext differentiates from other interfaces.
	IsCommentlineContext()
}

type CommentlineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentlineContext() *CommentlineContext {
	var p = new(CommentlineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = fastaParserRULE_commentline
	return p
}

func (*CommentlineContext) IsCommentlineContext() {}

func NewCommentlineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentlineContext {
	var p = new(CommentlineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = fastaParserRULE_commentline

	return p
}

func (s *CommentlineContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentlineContext) COMMENTLINE() antlr.TerminalNode {
	return s.GetToken(fastaParserCOMMENTLINE, 0)
}

func (s *CommentlineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentlineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentlineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fastaListener); ok {
		listenerT.EnterCommentline(s)
	}
}

func (s *CommentlineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(fastaListener); ok {
		listenerT.ExitCommentline(s)
	}
}

func (p *fastaParser) Commentline() (localctx ICommentlineContext) {
	this := p
	_ = this

	localctx = NewCommentlineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, fastaParserRULE_commentline)

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
		p.SetState(27)
		p.Match(fastaParserCOMMENTLINE)
	}

	return localctx
}
