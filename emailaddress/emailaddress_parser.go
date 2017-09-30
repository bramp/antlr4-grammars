// Generated from emailaddress.g4 by ANTLR 4.7.

package emailaddress // emailaddress
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 166,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 5, 2, 45, 10,
	2, 3, 3, 3, 3, 3, 3, 7, 3, 50, 10, 3, 12, 3, 14, 3, 53, 11, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 61, 10, 4, 3, 5, 3, 5, 7, 5, 65, 10, 5,
	12, 5, 14, 5, 68, 11, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 84, 10, 8, 12, 8, 14, 8, 87,
	11, 8, 3, 9, 3, 9, 3, 9, 7, 9, 92, 10, 9, 12, 9, 14, 9, 95, 11, 9, 3, 10,
	3, 10, 5, 10, 99, 10, 10, 3, 11, 3, 11, 3, 12, 6, 12, 104, 10, 12, 13,
	12, 14, 12, 105, 3, 13, 3, 13, 5, 13, 110, 10, 13, 3, 14, 3, 14, 3, 15,
	5, 15, 115, 10, 15, 3, 15, 6, 15, 118, 10, 15, 13, 15, 14, 15, 119, 3,
	16, 3, 16, 3, 16, 5, 16, 125, 10, 16, 3, 17, 6, 17, 128, 10, 17, 13, 17,
	14, 17, 129, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 7, 19, 138, 10,
	19, 12, 19, 14, 19, 141, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 7,
	20, 148, 10, 20, 12, 20, 14, 20, 151, 11, 20, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 7, 21, 159, 10, 21, 12, 21, 14, 21, 162, 11, 21, 3, 21, 3,
	21, 3, 21, 2, 2, 22, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 2, 3, 3, 2, 21, 22, 2, 166, 2, 44, 3, 2, 2, 2,
	4, 46, 3, 2, 2, 2, 6, 60, 3, 2, 2, 2, 8, 62, 3, 2, 2, 2, 10, 72, 3, 2,
	2, 2, 12, 76, 3, 2, 2, 2, 14, 80, 3, 2, 2, 2, 16, 88, 3, 2, 2, 2, 18, 98,
	3, 2, 2, 2, 20, 100, 3, 2, 2, 2, 22, 103, 3, 2, 2, 2, 24, 109, 3, 2, 2,
	2, 26, 111, 3, 2, 2, 2, 28, 117, 3, 2, 2, 2, 30, 124, 3, 2, 2, 2, 32, 127,
	3, 2, 2, 2, 34, 131, 3, 2, 2, 2, 36, 134, 3, 2, 2, 2, 38, 144, 3, 2, 2,
	2, 40, 154, 3, 2, 2, 2, 42, 45, 5, 6, 4, 2, 43, 45, 5, 4, 3, 2, 44, 42,
	3, 2, 2, 2, 44, 43, 3, 2, 2, 2, 45, 3, 3, 2, 2, 2, 46, 47, 5, 22, 12, 2,
	47, 51, 7, 3, 2, 2, 48, 50, 5, 6, 4, 2, 49, 48, 3, 2, 2, 2, 50, 53, 3,
	2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53,
	51, 3, 2, 2, 2, 54, 55, 7, 4, 2, 2, 55, 5, 3, 2, 2, 2, 56, 61, 5, 12, 7,
	2, 57, 58, 5, 22, 12, 2, 58, 59, 5, 8, 5, 2, 59, 61, 3, 2, 2, 2, 60, 56,
	3, 2, 2, 2, 60, 57, 3, 2, 2, 2, 61, 7, 3, 2, 2, 2, 62, 66, 7, 5, 2, 2,
	63, 65, 5, 10, 6, 2, 64, 63, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3,
	2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 69, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69,
	70, 5, 12, 7, 2, 70, 71, 7, 6, 2, 2, 71, 9, 3, 2, 2, 2, 72, 73, 7, 7, 2,
	2, 73, 74, 5, 16, 9, 2, 74, 75, 7, 3, 2, 2, 75, 11, 3, 2, 2, 2, 76, 77,
	5, 14, 8, 2, 77, 78, 7, 7, 2, 2, 78, 79, 5, 16, 9, 2, 79, 13, 3, 2, 2,
	2, 80, 85, 5, 24, 13, 2, 81, 82, 7, 8, 2, 2, 82, 84, 5, 24, 13, 2, 83,
	81, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2,
	2, 86, 15, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 93, 5, 18, 10, 2, 89, 90,
	7, 8, 2, 2, 90, 92, 5, 18, 10, 2, 91, 89, 3, 2, 2, 2, 92, 95, 3, 2, 2,
	2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 17, 3, 2, 2, 2, 95, 93,
	3, 2, 2, 2, 96, 99, 5, 20, 11, 2, 97, 99, 5, 36, 19, 2, 98, 96, 3, 2, 2,
	2, 98, 97, 3, 2, 2, 2, 99, 19, 3, 2, 2, 2, 100, 101, 5, 32, 17, 2, 101,
	21, 3, 2, 2, 2, 102, 104, 5, 24, 13, 2, 103, 102, 3, 2, 2, 2, 104, 105,
	3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 23, 3, 2,
	2, 2, 107, 110, 5, 32, 17, 2, 108, 110, 5, 38, 20, 2, 109, 107, 3, 2, 2,
	2, 109, 108, 3, 2, 2, 2, 110, 25, 3, 2, 2, 2, 111, 112, 9, 2, 2, 2, 112,
	27, 3, 2, 2, 2, 113, 115, 7, 23, 2, 2, 114, 113, 3, 2, 2, 2, 114, 115,
	3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 118, 5, 26, 14, 2, 117, 114, 3,
	2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2,
	2, 120, 29, 3, 2, 2, 2, 121, 125, 7, 24, 2, 2, 122, 125, 5, 28, 15, 2,
	123, 125, 5, 40, 21, 2, 124, 121, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124,
	123, 3, 2, 2, 2, 125, 31, 3, 2, 2, 2, 126, 128, 7, 15, 2, 2, 127, 126,
	3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2,
	2, 2, 130, 33, 3, 2, 2, 2, 131, 132, 7, 9, 2, 2, 132, 133, 7, 15, 2, 2,
	133, 35, 3, 2, 2, 2, 134, 139, 7, 10, 2, 2, 135, 138, 7, 27, 2, 2, 136,
	138, 5, 34, 18, 2, 137, 135, 3, 2, 2, 2, 137, 136, 3, 2, 2, 2, 138, 141,
	3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 142, 3, 2,
	2, 2, 141, 139, 3, 2, 2, 2, 142, 143, 7, 11, 2, 2, 143, 37, 3, 2, 2, 2,
	144, 149, 7, 12, 2, 2, 145, 148, 7, 26, 2, 2, 146, 148, 5, 34, 18, 2, 147,
	145, 3, 2, 2, 2, 147, 146, 3, 2, 2, 2, 148, 151, 3, 2, 2, 2, 149, 147,
	3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151, 149, 3, 2,
	2, 2, 152, 153, 7, 12, 2, 2, 153, 39, 3, 2, 2, 2, 154, 160, 7, 13, 2, 2,
	155, 159, 7, 28, 2, 2, 156, 159, 5, 34, 18, 2, 157, 159, 5, 40, 21, 2,
	158, 155, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 157, 3, 2, 2, 2, 159,
	162, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 163,
	3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 163, 164, 7, 14, 2, 2, 164, 41, 3, 2,
	2, 2, 21, 44, 51, 60, 66, 85, 93, 98, 105, 109, 114, 119, 124, 129, 137,
	139, 147, 149, 158, 160,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "';'", "'<'", "'>'", "'@'", "'.'", "'\\'", "'['", "']'", "'''",
	"'('", "')'", "", "", "", "", "'\n'", "'\r'", "' '", "'\t'", "'\r\n'",
	"", "'\"'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "CHAR", "ALPHA", "DIGIT",
	"CTL", "CR", "LF", "SPACE", "HTAB", "CRLF", "SPECIALS", "QUOTE", "QTEXT",
	"DTEXT", "CTEXT",
}

var ruleNames = []string{
	"emailaddress", "group", "mailbox", "routeaddr", "route", "addrspec", "localpart",
	"domain", "subdomain", "domainref", "phrase", "word", "lwspchar", "lwsp",
	"delimeters", "atom", "quotedpair", "domainliteral", "quotedstring", "comment",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type emailaddressParser struct {
	*antlr.BaseParser
}

func NewemailaddressParser(input antlr.TokenStream) *emailaddressParser {
	this := new(emailaddressParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "emailaddress.g4"

	return this
}

// emailaddressParser tokens.
const (
	emailaddressParserEOF      = antlr.TokenEOF
	emailaddressParserT__0     = 1
	emailaddressParserT__1     = 2
	emailaddressParserT__2     = 3
	emailaddressParserT__3     = 4
	emailaddressParserT__4     = 5
	emailaddressParserT__5     = 6
	emailaddressParserT__6     = 7
	emailaddressParserT__7     = 8
	emailaddressParserT__8     = 9
	emailaddressParserT__9     = 10
	emailaddressParserT__10    = 11
	emailaddressParserT__11    = 12
	emailaddressParserCHAR     = 13
	emailaddressParserALPHA    = 14
	emailaddressParserDIGIT    = 15
	emailaddressParserCTL      = 16
	emailaddressParserCR       = 17
	emailaddressParserLF       = 18
	emailaddressParserSPACE    = 19
	emailaddressParserHTAB     = 20
	emailaddressParserCRLF     = 21
	emailaddressParserSPECIALS = 22
	emailaddressParserQUOTE    = 23
	emailaddressParserQTEXT    = 24
	emailaddressParserDTEXT    = 25
	emailaddressParserCTEXT    = 26
)

// emailaddressParser rules.
const (
	emailaddressParserRULE_emailaddress  = 0
	emailaddressParserRULE_group         = 1
	emailaddressParserRULE_mailbox       = 2
	emailaddressParserRULE_routeaddr     = 3
	emailaddressParserRULE_route         = 4
	emailaddressParserRULE_addrspec      = 5
	emailaddressParserRULE_localpart     = 6
	emailaddressParserRULE_domain        = 7
	emailaddressParserRULE_subdomain     = 8
	emailaddressParserRULE_domainref     = 9
	emailaddressParserRULE_phrase        = 10
	emailaddressParserRULE_word          = 11
	emailaddressParserRULE_lwspchar      = 12
	emailaddressParserRULE_lwsp          = 13
	emailaddressParserRULE_delimeters    = 14
	emailaddressParserRULE_atom          = 15
	emailaddressParserRULE_quotedpair    = 16
	emailaddressParserRULE_domainliteral = 17
	emailaddressParserRULE_quotedstring  = 18
	emailaddressParserRULE_comment       = 19
)

// IEmailaddressContext is an interface to support dynamic dispatch.
type IEmailaddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmailaddressContext differentiates from other interfaces.
	IsEmailaddressContext()
}

type EmailaddressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmailaddressContext() *EmailaddressContext {
	var p = new(EmailaddressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_emailaddress
	return p
}

func (*EmailaddressContext) IsEmailaddressContext() {}

func NewEmailaddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmailaddressContext {
	var p = new(EmailaddressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_emailaddress

	return p
}

func (s *EmailaddressContext) GetParser() antlr.Parser { return s.parser }

func (s *EmailaddressContext) Mailbox() IMailboxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMailboxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMailboxContext)
}

func (s *EmailaddressContext) Group() IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *EmailaddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmailaddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmailaddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterEmailaddress(s)
	}
}

func (s *EmailaddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitEmailaddress(s)
	}
}

func (s *EmailaddressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitEmailaddress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Emailaddress() (localctx IEmailaddressContext) {
	localctx = NewEmailaddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, emailaddressParserRULE_emailaddress)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Mailbox()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.Group()
		}

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
	p.RuleIndex = emailaddressParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) Phrase() IPhraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPhraseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPhraseContext)
}

func (s *GroupContext) AllMailbox() []IMailboxContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMailboxContext)(nil)).Elem())
	var tst = make([]IMailboxContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMailboxContext)
		}
	}

	return tst
}

func (s *GroupContext) Mailbox(i int) IMailboxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMailboxContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMailboxContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, emailaddressParserRULE_group)
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
		p.SetState(44)
		p.Phrase()
	}
	{
		p.SetState(45)
		p.Match(emailaddressParserT__0)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == emailaddressParserT__9 || _la == emailaddressParserCHAR {
		{
			p.SetState(46)
			p.Mailbox()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
		p.Match(emailaddressParserT__1)
	}

	return localctx
}

// IMailboxContext is an interface to support dynamic dispatch.
type IMailboxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMailboxContext differentiates from other interfaces.
	IsMailboxContext()
}

type MailboxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMailboxContext() *MailboxContext {
	var p = new(MailboxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_mailbox
	return p
}

func (*MailboxContext) IsMailboxContext() {}

func NewMailboxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MailboxContext {
	var p = new(MailboxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_mailbox

	return p
}

func (s *MailboxContext) GetParser() antlr.Parser { return s.parser }

func (s *MailboxContext) Addrspec() IAddrspecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddrspecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddrspecContext)
}

func (s *MailboxContext) Phrase() IPhraseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPhraseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPhraseContext)
}

func (s *MailboxContext) Routeaddr() IRouteaddrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRouteaddrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRouteaddrContext)
}

func (s *MailboxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MailboxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MailboxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterMailbox(s)
	}
}

func (s *MailboxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitMailbox(s)
	}
}

func (s *MailboxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitMailbox(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Mailbox() (localctx IMailboxContext) {
	localctx = NewMailboxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, emailaddressParserRULE_mailbox)

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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Addrspec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Phrase()
		}
		{
			p.SetState(56)
			p.Routeaddr()
		}

	}

	return localctx
}

// IRouteaddrContext is an interface to support dynamic dispatch.
type IRouteaddrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRouteaddrContext differentiates from other interfaces.
	IsRouteaddrContext()
}

type RouteaddrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRouteaddrContext() *RouteaddrContext {
	var p = new(RouteaddrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_routeaddr
	return p
}

func (*RouteaddrContext) IsRouteaddrContext() {}

func NewRouteaddrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouteaddrContext {
	var p = new(RouteaddrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_routeaddr

	return p
}

func (s *RouteaddrContext) GetParser() antlr.Parser { return s.parser }

func (s *RouteaddrContext) Addrspec() IAddrspecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddrspecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddrspecContext)
}

func (s *RouteaddrContext) AllRoute() []IRouteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRouteContext)(nil)).Elem())
	var tst = make([]IRouteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRouteContext)
		}
	}

	return tst
}

func (s *RouteaddrContext) Route(i int) IRouteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRouteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRouteContext)
}

func (s *RouteaddrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RouteaddrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RouteaddrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterRouteaddr(s)
	}
}

func (s *RouteaddrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitRouteaddr(s)
	}
}

func (s *RouteaddrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitRouteaddr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Routeaddr() (localctx IRouteaddrContext) {
	localctx = NewRouteaddrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, emailaddressParserRULE_routeaddr)
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
		p.SetState(60)
		p.Match(emailaddressParserT__2)
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == emailaddressParserT__4 {
		{
			p.SetState(61)
			p.Route()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Addrspec()
	}
	{
		p.SetState(68)
		p.Match(emailaddressParserT__3)
	}

	return localctx
}

// IRouteContext is an interface to support dynamic dispatch.
type IRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRouteContext differentiates from other interfaces.
	IsRouteContext()
}

type RouteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRouteContext() *RouteContext {
	var p = new(RouteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_route
	return p
}

func (*RouteContext) IsRouteContext() {}

func NewRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouteContext {
	var p = new(RouteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_route

	return p
}

func (s *RouteContext) GetParser() antlr.Parser { return s.parser }

func (s *RouteContext) Domain() IDomainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDomainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDomainContext)
}

func (s *RouteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RouteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterRoute(s)
	}
}

func (s *RouteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitRoute(s)
	}
}

func (s *RouteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitRoute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Route() (localctx IRouteContext) {
	localctx = NewRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, emailaddressParserRULE_route)

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
		p.Match(emailaddressParserT__4)
	}
	{
		p.SetState(71)
		p.Domain()
	}
	{
		p.SetState(72)
		p.Match(emailaddressParserT__0)
	}

	return localctx
}

// IAddrspecContext is an interface to support dynamic dispatch.
type IAddrspecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddrspecContext differentiates from other interfaces.
	IsAddrspecContext()
}

type AddrspecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddrspecContext() *AddrspecContext {
	var p = new(AddrspecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_addrspec
	return p
}

func (*AddrspecContext) IsAddrspecContext() {}

func NewAddrspecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddrspecContext {
	var p = new(AddrspecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_addrspec

	return p
}

func (s *AddrspecContext) GetParser() antlr.Parser { return s.parser }

func (s *AddrspecContext) Localpart() ILocalpartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalpartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalpartContext)
}

func (s *AddrspecContext) Domain() IDomainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDomainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDomainContext)
}

func (s *AddrspecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddrspecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddrspecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterAddrspec(s)
	}
}

func (s *AddrspecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitAddrspec(s)
	}
}

func (s *AddrspecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitAddrspec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Addrspec() (localctx IAddrspecContext) {
	localctx = NewAddrspecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, emailaddressParserRULE_addrspec)

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
		p.SetState(74)
		p.Localpart()
	}
	{
		p.SetState(75)
		p.Match(emailaddressParserT__4)
	}
	{
		p.SetState(76)
		p.Domain()
	}

	return localctx
}

// ILocalpartContext is an interface to support dynamic dispatch.
type ILocalpartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalpartContext differentiates from other interfaces.
	IsLocalpartContext()
}

type LocalpartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalpartContext() *LocalpartContext {
	var p = new(LocalpartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_localpart
	return p
}

func (*LocalpartContext) IsLocalpartContext() {}

func NewLocalpartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalpartContext {
	var p = new(LocalpartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_localpart

	return p
}

func (s *LocalpartContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalpartContext) AllWord() []IWordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWordContext)(nil)).Elem())
	var tst = make([]IWordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWordContext)
		}
	}

	return tst
}

func (s *LocalpartContext) Word(i int) IWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *LocalpartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalpartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalpartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterLocalpart(s)
	}
}

func (s *LocalpartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitLocalpart(s)
	}
}

func (s *LocalpartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitLocalpart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Localpart() (localctx ILocalpartContext) {
	localctx = NewLocalpartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, emailaddressParserRULE_localpart)
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
		p.SetState(78)
		p.Word()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == emailaddressParserT__5 {
		{
			p.SetState(79)
			p.Match(emailaddressParserT__5)
		}
		{
			p.SetState(80)
			p.Word()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

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
	p.RuleIndex = emailaddressParserRULE_domain
	return p
}

func (*DomainContext) IsDomainContext() {}

func NewDomainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainContext {
	var p = new(DomainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_domain

	return p
}

func (s *DomainContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainContext) AllSubdomain() []ISubdomainContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubdomainContext)(nil)).Elem())
	var tst = make([]ISubdomainContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubdomainContext)
		}
	}

	return tst
}

func (s *DomainContext) Subdomain(i int) ISubdomainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubdomainContext)(nil)).Elem(), i)

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
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterDomain(s)
	}
}

func (s *DomainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitDomain(s)
	}
}

func (s *DomainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitDomain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Domain() (localctx IDomainContext) {
	localctx = NewDomainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, emailaddressParserRULE_domain)
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
		p.SetState(86)
		p.Subdomain()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == emailaddressParserT__5 {
		{
			p.SetState(87)
			p.Match(emailaddressParserT__5)
		}
		{
			p.SetState(88)
			p.Subdomain()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = emailaddressParserRULE_subdomain
	return p
}

func (*SubdomainContext) IsSubdomainContext() {}

func NewSubdomainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubdomainContext {
	var p = new(SubdomainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_subdomain

	return p
}

func (s *SubdomainContext) GetParser() antlr.Parser { return s.parser }

func (s *SubdomainContext) Domainref() IDomainrefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDomainrefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDomainrefContext)
}

func (s *SubdomainContext) Domainliteral() IDomainliteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDomainliteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDomainliteralContext)
}

func (s *SubdomainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubdomainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubdomainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterSubdomain(s)
	}
}

func (s *SubdomainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitSubdomain(s)
	}
}

func (s *SubdomainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitSubdomain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Subdomain() (localctx ISubdomainContext) {
	localctx = NewSubdomainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, emailaddressParserRULE_subdomain)

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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case emailaddressParserCHAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Domainref()
		}

	case emailaddressParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Domainliteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDomainrefContext is an interface to support dynamic dispatch.
type IDomainrefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDomainrefContext differentiates from other interfaces.
	IsDomainrefContext()
}

type DomainrefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDomainrefContext() *DomainrefContext {
	var p = new(DomainrefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_domainref
	return p
}

func (*DomainrefContext) IsDomainrefContext() {}

func NewDomainrefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainrefContext {
	var p = new(DomainrefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_domainref

	return p
}

func (s *DomainrefContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainrefContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *DomainrefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DomainrefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DomainrefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterDomainref(s)
	}
}

func (s *DomainrefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitDomainref(s)
	}
}

func (s *DomainrefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitDomainref(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Domainref() (localctx IDomainrefContext) {
	localctx = NewDomainrefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, emailaddressParserRULE_domainref)

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
		p.SetState(98)
		p.Atom()
	}

	return localctx
}

// IPhraseContext is an interface to support dynamic dispatch.
type IPhraseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhraseContext differentiates from other interfaces.
	IsPhraseContext()
}

type PhraseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhraseContext() *PhraseContext {
	var p = new(PhraseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_phrase
	return p
}

func (*PhraseContext) IsPhraseContext() {}

func NewPhraseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhraseContext {
	var p = new(PhraseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_phrase

	return p
}

func (s *PhraseContext) GetParser() antlr.Parser { return s.parser }

func (s *PhraseContext) AllWord() []IWordContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWordContext)(nil)).Elem())
	var tst = make([]IWordContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWordContext)
		}
	}

	return tst
}

func (s *PhraseContext) Word(i int) IWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWordContext)
}

func (s *PhraseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhraseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhraseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterPhrase(s)
	}
}

func (s *PhraseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitPhrase(s)
	}
}

func (s *PhraseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitPhrase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Phrase() (localctx IPhraseContext) {
	localctx = NewPhraseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, emailaddressParserRULE_phrase)
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
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == emailaddressParserT__9 || _la == emailaddressParserCHAR {
		{
			p.SetState(100)
			p.Word()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IWordContext is an interface to support dynamic dispatch.
type IWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordContext differentiates from other interfaces.
	IsWordContext()
}

type WordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordContext() *WordContext {
	var p = new(WordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_word
	return p
}

func (*WordContext) IsWordContext() {}

func NewWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordContext {
	var p = new(WordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_word

	return p
}

func (s *WordContext) GetParser() antlr.Parser { return s.parser }

func (s *WordContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *WordContext) Quotedstring() IQuotedstringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedstringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedstringContext)
}

func (s *WordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterWord(s)
	}
}

func (s *WordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitWord(s)
	}
}

func (s *WordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitWord(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Word() (localctx IWordContext) {
	localctx = NewWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, emailaddressParserRULE_word)

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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case emailaddressParserCHAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Atom()
		}

	case emailaddressParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.Quotedstring()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILwspcharContext is an interface to support dynamic dispatch.
type ILwspcharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLwspcharContext differentiates from other interfaces.
	IsLwspcharContext()
}

type LwspcharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLwspcharContext() *LwspcharContext {
	var p = new(LwspcharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_lwspchar
	return p
}

func (*LwspcharContext) IsLwspcharContext() {}

func NewLwspcharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LwspcharContext {
	var p = new(LwspcharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_lwspchar

	return p
}

func (s *LwspcharContext) GetParser() antlr.Parser { return s.parser }

func (s *LwspcharContext) SPACE() antlr.TerminalNode {
	return s.GetToken(emailaddressParserSPACE, 0)
}

func (s *LwspcharContext) HTAB() antlr.TerminalNode {
	return s.GetToken(emailaddressParserHTAB, 0)
}

func (s *LwspcharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LwspcharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LwspcharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterLwspchar(s)
	}
}

func (s *LwspcharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitLwspchar(s)
	}
}

func (s *LwspcharContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitLwspchar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Lwspchar() (localctx ILwspcharContext) {
	localctx = NewLwspcharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, emailaddressParserRULE_lwspchar)
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
	p.SetState(109)
	_la = p.GetTokenStream().LA(1)

	if !(_la == emailaddressParserSPACE || _la == emailaddressParserHTAB) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ILwspContext is an interface to support dynamic dispatch.
type ILwspContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLwspContext differentiates from other interfaces.
	IsLwspContext()
}

type LwspContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLwspContext() *LwspContext {
	var p = new(LwspContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_lwsp
	return p
}

func (*LwspContext) IsLwspContext() {}

func NewLwspContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LwspContext {
	var p = new(LwspContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_lwsp

	return p
}

func (s *LwspContext) GetParser() antlr.Parser { return s.parser }

func (s *LwspContext) AllLwspchar() []ILwspcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILwspcharContext)(nil)).Elem())
	var tst = make([]ILwspcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILwspcharContext)
		}
	}

	return tst
}

func (s *LwspContext) Lwspchar(i int) ILwspcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILwspcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILwspcharContext)
}

func (s *LwspContext) AllCRLF() []antlr.TerminalNode {
	return s.GetTokens(emailaddressParserCRLF)
}

func (s *LwspContext) CRLF(i int) antlr.TerminalNode {
	return s.GetToken(emailaddressParserCRLF, i)
}

func (s *LwspContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LwspContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LwspContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterLwsp(s)
	}
}

func (s *LwspContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitLwsp(s)
	}
}

func (s *LwspContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitLwsp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Lwsp() (localctx ILwspContext) {
	localctx = NewLwspContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, emailaddressParserRULE_lwsp)
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
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<emailaddressParserSPACE)|(1<<emailaddressParserHTAB)|(1<<emailaddressParserCRLF))) != 0) {
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == emailaddressParserCRLF {
			{
				p.SetState(111)
				p.Match(emailaddressParserCRLF)
			}

		}
		{
			p.SetState(114)
			p.Lwspchar()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDelimetersContext is an interface to support dynamic dispatch.
type IDelimetersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelimetersContext differentiates from other interfaces.
	IsDelimetersContext()
}

type DelimetersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelimetersContext() *DelimetersContext {
	var p = new(DelimetersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_delimeters
	return p
}

func (*DelimetersContext) IsDelimetersContext() {}

func NewDelimetersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelimetersContext {
	var p = new(DelimetersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_delimeters

	return p
}

func (s *DelimetersContext) GetParser() antlr.Parser { return s.parser }

func (s *DelimetersContext) SPECIALS() antlr.TerminalNode {
	return s.GetToken(emailaddressParserSPECIALS, 0)
}

func (s *DelimetersContext) Lwsp() ILwspContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILwspContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILwspContext)
}

func (s *DelimetersContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *DelimetersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelimetersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelimetersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterDelimeters(s)
	}
}

func (s *DelimetersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitDelimeters(s)
	}
}

func (s *DelimetersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitDelimeters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Delimeters() (localctx IDelimetersContext) {
	localctx = NewDelimetersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, emailaddressParserRULE_delimeters)

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

	p.SetState(122)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case emailaddressParserSPECIALS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(emailaddressParserSPECIALS)
		}

	case emailaddressParserSPACE, emailaddressParserHTAB, emailaddressParserCRLF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Lwsp()
		}

	case emailaddressParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.Comment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) AllCHAR() []antlr.TerminalNode {
	return s.GetTokens(emailaddressParserCHAR)
}

func (s *AtomContext) CHAR(i int) antlr.TerminalNode {
	return s.GetToken(emailaddressParserCHAR, i)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, emailaddressParserRULE_atom)

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
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(124)
				p.Match(emailaddressParserCHAR)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IQuotedpairContext is an interface to support dynamic dispatch.
type IQuotedpairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedpairContext differentiates from other interfaces.
	IsQuotedpairContext()
}

type QuotedpairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedpairContext() *QuotedpairContext {
	var p = new(QuotedpairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_quotedpair
	return p
}

func (*QuotedpairContext) IsQuotedpairContext() {}

func NewQuotedpairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedpairContext {
	var p = new(QuotedpairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_quotedpair

	return p
}

func (s *QuotedpairContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedpairContext) CHAR() antlr.TerminalNode {
	return s.GetToken(emailaddressParserCHAR, 0)
}

func (s *QuotedpairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedpairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedpairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterQuotedpair(s)
	}
}

func (s *QuotedpairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitQuotedpair(s)
	}
}

func (s *QuotedpairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitQuotedpair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Quotedpair() (localctx IQuotedpairContext) {
	localctx = NewQuotedpairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, emailaddressParserRULE_quotedpair)

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
		p.SetState(129)
		p.Match(emailaddressParserT__6)
	}
	{
		p.SetState(130)
		p.Match(emailaddressParserCHAR)
	}

	return localctx
}

// IDomainliteralContext is an interface to support dynamic dispatch.
type IDomainliteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDomainliteralContext differentiates from other interfaces.
	IsDomainliteralContext()
}

type DomainliteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDomainliteralContext() *DomainliteralContext {
	var p = new(DomainliteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_domainliteral
	return p
}

func (*DomainliteralContext) IsDomainliteralContext() {}

func NewDomainliteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainliteralContext {
	var p = new(DomainliteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_domainliteral

	return p
}

func (s *DomainliteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainliteralContext) AllDTEXT() []antlr.TerminalNode {
	return s.GetTokens(emailaddressParserDTEXT)
}

func (s *DomainliteralContext) DTEXT(i int) antlr.TerminalNode {
	return s.GetToken(emailaddressParserDTEXT, i)
}

func (s *DomainliteralContext) AllQuotedpair() []IQuotedpairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuotedpairContext)(nil)).Elem())
	var tst = make([]IQuotedpairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuotedpairContext)
		}
	}

	return tst
}

func (s *DomainliteralContext) Quotedpair(i int) IQuotedpairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedpairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuotedpairContext)
}

func (s *DomainliteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DomainliteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DomainliteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterDomainliteral(s)
	}
}

func (s *DomainliteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitDomainliteral(s)
	}
}

func (s *DomainliteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitDomainliteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Domainliteral() (localctx IDomainliteralContext) {
	localctx = NewDomainliteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, emailaddressParserRULE_domainliteral)
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
		p.SetState(132)
		p.Match(emailaddressParserT__7)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == emailaddressParserT__6 || _la == emailaddressParserDTEXT {
		p.SetState(135)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case emailaddressParserDTEXT:
			{
				p.SetState(133)
				p.Match(emailaddressParserDTEXT)
			}

		case emailaddressParserT__6:
			{
				p.SetState(134)
				p.Quotedpair()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
		p.Match(emailaddressParserT__8)
	}

	return localctx
}

// IQuotedstringContext is an interface to support dynamic dispatch.
type IQuotedstringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedstringContext differentiates from other interfaces.
	IsQuotedstringContext()
}

type QuotedstringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedstringContext() *QuotedstringContext {
	var p = new(QuotedstringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_quotedstring
	return p
}

func (*QuotedstringContext) IsQuotedstringContext() {}

func NewQuotedstringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedstringContext {
	var p = new(QuotedstringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_quotedstring

	return p
}

func (s *QuotedstringContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedstringContext) AllQTEXT() []antlr.TerminalNode {
	return s.GetTokens(emailaddressParserQTEXT)
}

func (s *QuotedstringContext) QTEXT(i int) antlr.TerminalNode {
	return s.GetToken(emailaddressParserQTEXT, i)
}

func (s *QuotedstringContext) AllQuotedpair() []IQuotedpairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuotedpairContext)(nil)).Elem())
	var tst = make([]IQuotedpairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuotedpairContext)
		}
	}

	return tst
}

func (s *QuotedstringContext) Quotedpair(i int) IQuotedpairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedpairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuotedpairContext)
}

func (s *QuotedstringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedstringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedstringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterQuotedstring(s)
	}
}

func (s *QuotedstringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitQuotedstring(s)
	}
}

func (s *QuotedstringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitQuotedstring(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Quotedstring() (localctx IQuotedstringContext) {
	localctx = NewQuotedstringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, emailaddressParserRULE_quotedstring)
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
		p.SetState(142)
		p.Match(emailaddressParserT__9)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == emailaddressParserT__6 || _la == emailaddressParserQTEXT {
		p.SetState(145)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case emailaddressParserQTEXT:
			{
				p.SetState(143)
				p.Match(emailaddressParserQTEXT)
			}

		case emailaddressParserT__6:
			{
				p.SetState(144)
				p.Quotedpair()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(emailaddressParserT__9)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = emailaddressParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = emailaddressParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) AllCTEXT() []antlr.TerminalNode {
	return s.GetTokens(emailaddressParserCTEXT)
}

func (s *CommentContext) CTEXT(i int) antlr.TerminalNode {
	return s.GetToken(emailaddressParserCTEXT, i)
}

func (s *CommentContext) AllQuotedpair() []IQuotedpairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuotedpairContext)(nil)).Elem())
	var tst = make([]IQuotedpairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuotedpairContext)
		}
	}

	return tst
}

func (s *CommentContext) Quotedpair(i int) IQuotedpairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedpairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuotedpairContext)
}

func (s *CommentContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *CommentContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(emailaddressListener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case emailaddressVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *emailaddressParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, emailaddressParserRULE_comment)
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
		p.SetState(152)
		p.Match(emailaddressParserT__10)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<emailaddressParserT__6)|(1<<emailaddressParserT__10)|(1<<emailaddressParserCTEXT))) != 0 {
		p.SetState(156)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case emailaddressParserCTEXT:
			{
				p.SetState(153)
				p.Match(emailaddressParserCTEXT)
			}

		case emailaddressParserT__6:
			{
				p.SetState(154)
				p.Quotedpair()
			}

		case emailaddressParserT__10:
			{
				p.SetState(155)
				p.Comment()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(161)
		p.Match(emailaddressParserT__11)
	}

	return localctx
}
