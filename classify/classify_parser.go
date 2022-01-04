// Code generated from classify.g4 by ANTLR 4.9.3. DO NOT EDIT.

package classify // classify
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 7, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 2, 2, 3, 2, 2, 3, 3, 2, 3, 33, 2, 5, 2, 4, 3,
	2, 2, 2, 4, 5, 9, 2, 2, 2, 5, 3, 3, 2, 2, 2, 2,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "'\u2028'", "'\u2029'",
}
var symbolicNames = []string{
	"", "CLASSIFY___", "CLASSIFY_Cc", "CLASSIFY_Cf", "CLASSIFY_Cn", "CLASSIFY_Co",
	"CLASSIFY_Cs", "CLASSIFY_Ll", "CLASSIFY_Lm", "CLASSIFY_Lo", "CLASSIFY_Lt",
	"CLASSIFY_Lu", "CLASSIFY_Mc", "CLASSIFY_Me", "CLASSIFY_Mn", "CLASSIFY_Nd",
	"CLASSIFY_Nl", "CLASSIFY_No", "CLASSIFY_Pc", "CLASSIFY_Pd", "CLASSIFY_Pe",
	"CLASSIFY_Pf", "CLASSIFY_Pi", "CLASSIFY_Po", "CLASSIFY_Ps", "CLASSIFY_Sc",
	"CLASSIFY_Sk", "CLASSIFY_Sm", "CLASSIFY_So", "CLASSIFY_Zl", "CLASSIFY_Zp",
	"CLASSIFY_Zs", "CLASSIFY_C", "CLASSIFY_LC", "CLASSIFY_M", "CLASSIFY_L",
	"CLASSIFY_N", "CLASSIFY_P", "CLASSIFY_S", "CLASSIFY_Z", "CLASSIFY_WS",
	"CLASSIFY_ID0", "CLASSIFY_ID1", "ID",
}

var ruleNames = []string{
	"codepoint",
}

type classifyParser struct {
	*antlr.BaseParser
}

// NewclassifyParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *classifyParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewclassifyParser(input antlr.TokenStream) *classifyParser {
	this := new(classifyParser)
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
	this.GrammarFileName = "classify.g4"

	return this
}

// classifyParser tokens.
const (
	classifyParserEOF          = antlr.TokenEOF
	classifyParserCLASSIFY___  = 1
	classifyParserCLASSIFY_Cc  = 2
	classifyParserCLASSIFY_Cf  = 3
	classifyParserCLASSIFY_Cn  = 4
	classifyParserCLASSIFY_Co  = 5
	classifyParserCLASSIFY_Cs  = 6
	classifyParserCLASSIFY_Ll  = 7
	classifyParserCLASSIFY_Lm  = 8
	classifyParserCLASSIFY_Lo  = 9
	classifyParserCLASSIFY_Lt  = 10
	classifyParserCLASSIFY_Lu  = 11
	classifyParserCLASSIFY_Mc  = 12
	classifyParserCLASSIFY_Me  = 13
	classifyParserCLASSIFY_Mn  = 14
	classifyParserCLASSIFY_Nd  = 15
	classifyParserCLASSIFY_Nl  = 16
	classifyParserCLASSIFY_No  = 17
	classifyParserCLASSIFY_Pc  = 18
	classifyParserCLASSIFY_Pd  = 19
	classifyParserCLASSIFY_Pe  = 20
	classifyParserCLASSIFY_Pf  = 21
	classifyParserCLASSIFY_Pi  = 22
	classifyParserCLASSIFY_Po  = 23
	classifyParserCLASSIFY_Ps  = 24
	classifyParserCLASSIFY_Sc  = 25
	classifyParserCLASSIFY_Sk  = 26
	classifyParserCLASSIFY_Sm  = 27
	classifyParserCLASSIFY_So  = 28
	classifyParserCLASSIFY_Zl  = 29
	classifyParserCLASSIFY_Zp  = 30
	classifyParserCLASSIFY_Zs  = 31
	classifyParserCLASSIFY_C   = 32
	classifyParserCLASSIFY_LC  = 33
	classifyParserCLASSIFY_M   = 34
	classifyParserCLASSIFY_L   = 35
	classifyParserCLASSIFY_N   = 36
	classifyParserCLASSIFY_P   = 37
	classifyParserCLASSIFY_S   = 38
	classifyParserCLASSIFY_Z   = 39
	classifyParserCLASSIFY_WS  = 40
	classifyParserCLASSIFY_ID0 = 41
	classifyParserCLASSIFY_ID1 = 42
	classifyParserID           = 43
)

// classifyParserRULE_codepoint is the classifyParser rule.
const classifyParserRULE_codepoint = 0

// ICodepointContext is an interface to support dynamic dispatch.
type ICodepointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCodepointContext differentiates from other interfaces.
	IsCodepointContext()
}

type CodepointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodepointContext() *CodepointContext {
	var p = new(CodepointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = classifyParserRULE_codepoint
	return p
}

func (*CodepointContext) IsCodepointContext() {}

func NewCodepointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodepointContext {
	var p = new(CodepointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = classifyParserRULE_codepoint

	return p
}

func (s *CodepointContext) GetParser() antlr.Parser { return s.parser }

func (s *CodepointContext) CLASSIFY___() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY___, 0)
}

func (s *CodepointContext) CLASSIFY_Cc() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Cc, 0)
}

func (s *CodepointContext) CLASSIFY_Cf() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Cf, 0)
}

func (s *CodepointContext) CLASSIFY_Cn() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Cn, 0)
}

func (s *CodepointContext) CLASSIFY_Co() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Co, 0)
}

func (s *CodepointContext) CLASSIFY_Cs() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Cs, 0)
}

func (s *CodepointContext) CLASSIFY_Ll() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Ll, 0)
}

func (s *CodepointContext) CLASSIFY_Lm() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Lm, 0)
}

func (s *CodepointContext) CLASSIFY_Lo() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Lo, 0)
}

func (s *CodepointContext) CLASSIFY_Lt() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Lt, 0)
}

func (s *CodepointContext) CLASSIFY_Lu() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Lu, 0)
}

func (s *CodepointContext) CLASSIFY_Mc() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Mc, 0)
}

func (s *CodepointContext) CLASSIFY_Me() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Me, 0)
}

func (s *CodepointContext) CLASSIFY_Mn() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Mn, 0)
}

func (s *CodepointContext) CLASSIFY_Nd() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Nd, 0)
}

func (s *CodepointContext) CLASSIFY_Nl() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Nl, 0)
}

func (s *CodepointContext) CLASSIFY_No() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_No, 0)
}

func (s *CodepointContext) CLASSIFY_Pc() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Pc, 0)
}

func (s *CodepointContext) CLASSIFY_Pd() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Pd, 0)
}

func (s *CodepointContext) CLASSIFY_Pe() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Pe, 0)
}

func (s *CodepointContext) CLASSIFY_Pf() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Pf, 0)
}

func (s *CodepointContext) CLASSIFY_Pi() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Pi, 0)
}

func (s *CodepointContext) CLASSIFY_Po() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Po, 0)
}

func (s *CodepointContext) CLASSIFY_Ps() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Ps, 0)
}

func (s *CodepointContext) CLASSIFY_Sc() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Sc, 0)
}

func (s *CodepointContext) CLASSIFY_Sk() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Sk, 0)
}

func (s *CodepointContext) CLASSIFY_Sm() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Sm, 0)
}

func (s *CodepointContext) CLASSIFY_So() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_So, 0)
}

func (s *CodepointContext) CLASSIFY_Zl() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Zl, 0)
}

func (s *CodepointContext) CLASSIFY_Zp() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Zp, 0)
}

func (s *CodepointContext) CLASSIFY_Zs() antlr.TerminalNode {
	return s.GetToken(classifyParserCLASSIFY_Zs, 0)
}

func (s *CodepointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodepointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodepointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(classifyListener); ok {
		listenerT.EnterCodepoint(s)
	}
}

func (s *CodepointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(classifyListener); ok {
		listenerT.ExitCodepoint(s)
	}
}

func (p *classifyParser) Codepoint() (localctx ICodepointContext) {
	this := p
	_ = this

	localctx = NewCodepointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, classifyParserRULE_codepoint)
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
		p.SetState(2)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<classifyParserCLASSIFY___)|(1<<classifyParserCLASSIFY_Cc)|(1<<classifyParserCLASSIFY_Cf)|(1<<classifyParserCLASSIFY_Cn)|(1<<classifyParserCLASSIFY_Co)|(1<<classifyParserCLASSIFY_Cs)|(1<<classifyParserCLASSIFY_Ll)|(1<<classifyParserCLASSIFY_Lm)|(1<<classifyParserCLASSIFY_Lo)|(1<<classifyParserCLASSIFY_Lt)|(1<<classifyParserCLASSIFY_Lu)|(1<<classifyParserCLASSIFY_Mc)|(1<<classifyParserCLASSIFY_Me)|(1<<classifyParserCLASSIFY_Mn)|(1<<classifyParserCLASSIFY_Nd)|(1<<classifyParserCLASSIFY_Nl)|(1<<classifyParserCLASSIFY_No)|(1<<classifyParserCLASSIFY_Pc)|(1<<classifyParserCLASSIFY_Pd)|(1<<classifyParserCLASSIFY_Pe)|(1<<classifyParserCLASSIFY_Pf)|(1<<classifyParserCLASSIFY_Pi)|(1<<classifyParserCLASSIFY_Po)|(1<<classifyParserCLASSIFY_Ps)|(1<<classifyParserCLASSIFY_Sc)|(1<<classifyParserCLASSIFY_Sk)|(1<<classifyParserCLASSIFY_Sm)|(1<<classifyParserCLASSIFY_So)|(1<<classifyParserCLASSIFY_Zl)|(1<<classifyParserCLASSIFY_Zp)|(1<<classifyParserCLASSIFY_Zs))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
