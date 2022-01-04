// Code generated from domain.g4 by ANTLR 4.9.3. DO NOT EDIT.

package domain // domain
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 19, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 5, 2, 9, 10, 2, 3, 3, 3, 3, 3, 3, 7, 3,
	14, 10, 3, 12, 3, 14, 3, 17, 11, 3, 3, 3, 2, 2, 4, 2, 4, 2, 2, 2, 18, 2,
	8, 3, 2, 2, 2, 4, 10, 3, 2, 2, 2, 6, 9, 5, 4, 3, 2, 7, 9, 7, 3, 2, 2, 8,
	6, 3, 2, 2, 2, 8, 7, 3, 2, 2, 2, 9, 3, 3, 2, 2, 2, 10, 15, 7, 5, 2, 2,
	11, 12, 7, 4, 2, 2, 12, 14, 7, 5, 2, 2, 13, 11, 3, 2, 2, 2, 14, 17, 3,
	2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 5, 3, 2, 2, 2, 17,
	15, 3, 2, 2, 2, 4, 8, 15,
}
var literalNames = []string{
	"", "' '", "'.'",
}
var symbolicNames = []string{
	"", "", "", "LABEL", "LDH_STR", "LET_DIG_HYP", "LET_DIG", "LETTER", "DIGIT",
}

var ruleNames = []string{
	"domain", "subdomain",
}

type domainParser struct {
	*antlr.BaseParser
}

// NewdomainParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *domainParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewdomainParser(input antlr.TokenStream) *domainParser {
	this := new(domainParser)
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
	this.GrammarFileName = "domain.g4"

	return this
}

// domainParser tokens.
const (
	domainParserEOF         = antlr.TokenEOF
	domainParserT__0        = 1
	domainParserT__1        = 2
	domainParserLABEL       = 3
	domainParserLDH_STR     = 4
	domainParserLET_DIG_HYP = 5
	domainParserLET_DIG     = 6
	domainParserLETTER      = 7
	domainParserDIGIT       = 8
)

// domainParser rules.
const (
	domainParserRULE_domain    = 0
	domainParserRULE_subdomain = 1
)

// IDomainContext is an interface to support dynamic dispatch.
type IDomainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDomainContext differentiates from other interfaces.
	IsDomainContext()
}

type DomainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDomainContext() *DomainContext {
	var p = new(DomainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = domainParserRULE_domain
	return p
}

func (*DomainContext) IsDomainContext() {}

func NewDomainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainContext {
	var p = new(DomainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = domainParserRULE_domain

	return p
}

func (s *DomainContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainContext) Subdomain() ISubdomainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubdomainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubdomainContext)
}

func (s *DomainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DomainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DomainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(domainListener); ok {
		listenerT.EnterDomain(s)
	}
}

func (s *DomainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(domainListener); ok {
		listenerT.ExitDomain(s)
	}
}

func (p *domainParser) Domain() (localctx IDomainContext) {
	this := p
	_ = this

	localctx = NewDomainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, domainParserRULE_domain)

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

	p.SetState(6)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case domainParserLABEL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(4)
			p.Subdomain()
		}

	case domainParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(5)
			p.Match(domainParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISubdomainContext is an interface to support dynamic dispatch.
type ISubdomainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubdomainContext differentiates from other interfaces.
	IsSubdomainContext()
}

type SubdomainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubdomainContext() *SubdomainContext {
	var p = new(SubdomainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = domainParserRULE_subdomain
	return p
}

func (*SubdomainContext) IsSubdomainContext() {}

func NewSubdomainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubdomainContext {
	var p = new(SubdomainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = domainParserRULE_subdomain

	return p
}

func (s *SubdomainContext) GetParser() antlr.Parser { return s.parser }

func (s *SubdomainContext) AllLABEL() []antlr.TerminalNode {
	return s.GetTokens(domainParserLABEL)
}

func (s *SubdomainContext) LABEL(i int) antlr.TerminalNode {
	return s.GetToken(domainParserLABEL, i)
}

func (s *SubdomainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubdomainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubdomainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(domainListener); ok {
		listenerT.EnterSubdomain(s)
	}
}

func (s *SubdomainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(domainListener); ok {
		listenerT.ExitSubdomain(s)
	}
}

func (p *domainParser) Subdomain() (localctx ISubdomainContext) {
	this := p
	_ = this

	localctx = NewSubdomainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, domainParserRULE_subdomain)
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
		p.SetState(8)
		p.Match(domainParserLABEL)
	}
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == domainParserT__1 {
		{
			p.SetState(9)
			p.Match(domainParserT__1)
		}
		{
			p.SetState(10)
			p.Match(domainParserLABEL)
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
