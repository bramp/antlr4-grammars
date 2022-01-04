// Code generated from http.g4 by ANTLR 4.9.3. DO NOT EDIT.

package http // http
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 171,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 55, 10, 2, 12,
	2, 14, 2, 58, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 5, 7, 78, 10,
	7, 3, 8, 3, 8, 6, 8, 82, 10, 8, 13, 8, 14, 8, 83, 3, 9, 7, 9, 87, 10, 9,
	12, 9, 14, 9, 90, 11, 9, 3, 10, 3, 10, 3, 10, 7, 10, 95, 10, 10, 12, 10,
	14, 10, 98, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 7, 13, 110, 10, 13, 12, 13, 14, 13, 113, 11, 13, 3, 13,
	3, 13, 7, 13, 117, 10, 13, 12, 13, 14, 13, 120, 11, 13, 3, 14, 3, 14, 3,
	15, 6, 15, 125, 10, 15, 13, 15, 14, 15, 126, 3, 16, 3, 16, 6, 16, 131,
	10, 16, 13, 16, 14, 16, 132, 3, 17, 3, 17, 6, 17, 137, 10, 17, 13, 17,
	14, 17, 138, 3, 17, 5, 17, 142, 10, 17, 3, 18, 3, 18, 5, 18, 146, 10, 18,
	3, 19, 3, 19, 3, 20, 3, 20, 6, 20, 152, 10, 20, 13, 20, 14, 20, 153, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 161, 10, 21, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 2, 2, 26, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	2, 8, 3, 2, 3, 10, 4, 2, 12, 12, 43, 43, 4, 2, 13, 14, 23, 26, 7, 2, 17,
	21, 29, 29, 32, 32, 34, 34, 36, 38, 6, 2, 13, 14, 23, 26, 29, 29, 32, 41,
	4, 2, 13, 14, 44, 44, 2, 166, 2, 50, 3, 2, 2, 2, 4, 61, 3, 2, 2, 2, 6,
	63, 3, 2, 2, 2, 8, 70, 3, 2, 2, 2, 10, 72, 3, 2, 2, 2, 12, 74, 3, 2, 2,
	2, 14, 81, 3, 2, 2, 2, 16, 88, 3, 2, 2, 2, 18, 96, 3, 2, 2, 2, 20, 99,
	3, 2, 2, 2, 22, 104, 3, 2, 2, 2, 24, 106, 3, 2, 2, 2, 26, 121, 3, 2, 2,
	2, 28, 124, 3, 2, 2, 2, 30, 130, 3, 2, 2, 2, 32, 134, 3, 2, 2, 2, 34, 145,
	3, 2, 2, 2, 36, 147, 3, 2, 2, 2, 38, 149, 3, 2, 2, 2, 40, 160, 3, 2, 2,
	2, 42, 162, 3, 2, 2, 2, 44, 164, 3, 2, 2, 2, 46, 166, 3, 2, 2, 2, 48, 168,
	3, 2, 2, 2, 50, 56, 5, 4, 3, 2, 51, 52, 5, 24, 13, 2, 52, 53, 7, 22, 2,
	2, 53, 55, 3, 2, 2, 2, 54, 51, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54,
	3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 59, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2,
	59, 60, 7, 22, 2, 2, 60, 3, 3, 2, 2, 2, 61, 62, 5, 6, 4, 2, 62, 5, 3, 2,
	2, 2, 63, 64, 5, 8, 5, 2, 64, 65, 7, 12, 2, 2, 65, 66, 5, 10, 6, 2, 66,
	67, 7, 12, 2, 2, 67, 68, 5, 20, 11, 2, 68, 69, 7, 22, 2, 2, 69, 7, 3, 2,
	2, 2, 70, 71, 9, 2, 2, 2, 71, 9, 3, 2, 2, 2, 72, 73, 5, 12, 7, 2, 73, 11,
	3, 2, 2, 2, 74, 77, 5, 14, 8, 2, 75, 76, 7, 27, 2, 2, 76, 78, 5, 18, 10,
	2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 13, 3, 2, 2, 2, 79, 80,
	7, 28, 2, 2, 80, 82, 5, 16, 9, 2, 81, 79, 3, 2, 2, 2, 82, 83, 3, 2, 2,
	2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 15, 3, 2, 2, 2, 85, 87,
	5, 40, 21, 2, 86, 85, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 3, 2, 2,
	2, 88, 89, 3, 2, 2, 2, 89, 17, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91, 95,
	5, 40, 21, 2, 92, 95, 7, 28, 2, 2, 93, 95, 7, 27, 2, 2, 94, 91, 3, 2, 2,
	2, 94, 92, 3, 2, 2, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94,
	3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 19, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2,
	99, 100, 5, 22, 12, 2, 100, 101, 7, 14, 2, 2, 101, 102, 7, 24, 2, 2, 102,
	103, 7, 14, 2, 2, 103, 21, 3, 2, 2, 2, 104, 105, 7, 11, 2, 2, 105, 23,
	3, 2, 2, 2, 106, 107, 5, 26, 14, 2, 107, 111, 7, 30, 2, 2, 108, 110, 7,
	42, 2, 2, 109, 108, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2,
	2, 111, 112, 3, 2, 2, 2, 112, 114, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114,
	118, 5, 30, 16, 2, 115, 117, 7, 42, 2, 2, 116, 115, 3, 2, 2, 2, 117, 120,
	3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 25, 3, 2,
	2, 2, 120, 118, 3, 2, 2, 2, 121, 122, 5, 28, 15, 2, 122, 27, 3, 2, 2, 2,
	123, 125, 5, 46, 24, 2, 124, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126,
	124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 29, 3, 2, 2, 2, 128, 131, 5,
	32, 17, 2, 129, 131, 5, 38, 20, 2, 130, 128, 3, 2, 2, 2, 130, 129, 3, 2,
	2, 2, 131, 132, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2,
	133, 31, 3, 2, 2, 2, 134, 141, 5, 34, 18, 2, 135, 137, 9, 3, 2, 2, 136,
	135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138, 139,
	3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 142, 5, 34, 18, 2, 141, 136, 3,
	2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 33, 3, 2, 2, 2, 143, 146, 5, 48, 25,
	2, 144, 146, 5, 36, 19, 2, 145, 143, 3, 2, 2, 2, 145, 144, 3, 2, 2, 2,
	146, 35, 3, 2, 2, 2, 147, 148, 7, 45, 2, 2, 148, 37, 3, 2, 2, 2, 149, 151,
	7, 22, 2, 2, 150, 152, 9, 3, 2, 2, 151, 150, 3, 2, 2, 2, 152, 153, 3, 2,
	2, 2, 153, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 39, 3, 2, 2, 2,
	155, 161, 5, 42, 22, 2, 156, 161, 7, 15, 2, 2, 157, 161, 5, 44, 23, 2,
	158, 161, 7, 30, 2, 2, 159, 161, 7, 31, 2, 2, 160, 155, 3, 2, 2, 2, 160,
	156, 3, 2, 2, 2, 160, 157, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 159,
	3, 2, 2, 2, 161, 41, 3, 2, 2, 2, 162, 163, 9, 4, 2, 2, 163, 43, 3, 2, 2,
	2, 164, 165, 9, 5, 2, 2, 165, 45, 3, 2, 2, 2, 166, 167, 9, 6, 2, 2, 167,
	47, 3, 2, 2, 2, 168, 169, 9, 7, 2, 2, 169, 49, 3, 2, 2, 2, 18, 56, 77,
	83, 88, 94, 96, 111, 118, 126, 130, 132, 138, 141, 145, 153, 160,
}
var literalNames = []string{
	"", "'GET'", "'HEAD'", "'POST'", "'PUT'", "'DELETE'", "'CONNECT'", "'OPTIONS'",
	"'TRACE'", "'HTTP/'", "' '", "", "", "", "", "'('", "')'", "';'", "'='",
	"','", "'\n'", "'-'", "'.'", "'_'", "'~'", "'?'", "'/'", "'!'", "':'",
	"'@'", "'$'", "'#'", "'&'", "'%'", "'''", "'*'", "'+'", "'^'", "'`'", "'|'",
	"", "'\t'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "SP", "ALPHA", "DIGIT", "Pct_encoded",
	"HEXDIG", "LColumn", "RColumn", "SemiColon", "Equals", "Period", "CRLF",
	"Minus", "Dot", "Underscore", "Tilde", "QuestionMark", "Slash", "ExclamationMark",
	"Colon", "At", "DollarSign", "Hashtag", "Ampersand", "Percent", "SQuote",
	"Star", "Plus", "Caret", "BackQuote", "VBar", "OWS", "HTAB", "VCHAR", "OBS_TEXT",
}

var ruleNames = []string{
	"http_message", "start_line", "request_line", "method", "request_target",
	"origin_form", "absolute_path", "segment", "query", "http_version", "http_name",
	"header_field", "field_name", "token", "field_value", "field_content",
	"field_vchar", "obs_text", "obs_fold", "pchar", "unreserved", "sub_delims",
	"tchar", "vCHAR",
}

type httpParser struct {
	*antlr.BaseParser
}

// NewhttpParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *httpParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewhttpParser(input antlr.TokenStream) *httpParser {
	this := new(httpParser)
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
	this.GrammarFileName = "http.g4"

	return this
}

// httpParser tokens.
const (
	httpParserEOF             = antlr.TokenEOF
	httpParserT__0            = 1
	httpParserT__1            = 2
	httpParserT__2            = 3
	httpParserT__3            = 4
	httpParserT__4            = 5
	httpParserT__5            = 6
	httpParserT__6            = 7
	httpParserT__7            = 8
	httpParserT__8            = 9
	httpParserSP              = 10
	httpParserALPHA           = 11
	httpParserDIGIT           = 12
	httpParserPct_encoded     = 13
	httpParserHEXDIG          = 14
	httpParserLColumn         = 15
	httpParserRColumn         = 16
	httpParserSemiColon       = 17
	httpParserEquals          = 18
	httpParserPeriod          = 19
	httpParserCRLF            = 20
	httpParserMinus           = 21
	httpParserDot             = 22
	httpParserUnderscore      = 23
	httpParserTilde           = 24
	httpParserQuestionMark    = 25
	httpParserSlash           = 26
	httpParserExclamationMark = 27
	httpParserColon           = 28
	httpParserAt              = 29
	httpParserDollarSign      = 30
	httpParserHashtag         = 31
	httpParserAmpersand       = 32
	httpParserPercent         = 33
	httpParserSQuote          = 34
	httpParserStar            = 35
	httpParserPlus            = 36
	httpParserCaret           = 37
	httpParserBackQuote       = 38
	httpParserVBar            = 39
	httpParserOWS             = 40
	httpParserHTAB            = 41
	httpParserVCHAR           = 42
	httpParserOBS_TEXT        = 43
)

// httpParser rules.
const (
	httpParserRULE_http_message   = 0
	httpParserRULE_start_line     = 1
	httpParserRULE_request_line   = 2
	httpParserRULE_method         = 3
	httpParserRULE_request_target = 4
	httpParserRULE_origin_form    = 5
	httpParserRULE_absolute_path  = 6
	httpParserRULE_segment        = 7
	httpParserRULE_query          = 8
	httpParserRULE_http_version   = 9
	httpParserRULE_http_name      = 10
	httpParserRULE_header_field   = 11
	httpParserRULE_field_name     = 12
	httpParserRULE_token          = 13
	httpParserRULE_field_value    = 14
	httpParserRULE_field_content  = 15
	httpParserRULE_field_vchar    = 16
	httpParserRULE_obs_text       = 17
	httpParserRULE_obs_fold       = 18
	httpParserRULE_pchar          = 19
	httpParserRULE_unreserved     = 20
	httpParserRULE_sub_delims     = 21
	httpParserRULE_tchar          = 22
	httpParserRULE_vCHAR          = 23
)

// IHttp_messageContext is an interface to support dynamic dispatch.
type IHttp_messageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHttp_messageContext differentiates from other interfaces.
	IsHttp_messageContext()
}

type Http_messageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttp_messageContext() *Http_messageContext {
	var p = new(Http_messageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_http_message
	return p
}

func (*Http_messageContext) IsHttp_messageContext() {}

func NewHttp_messageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Http_messageContext {
	var p = new(Http_messageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_http_message

	return p
}

func (s *Http_messageContext) GetParser() antlr.Parser { return s.parser }

func (s *Http_messageContext) Start_line() IStart_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStart_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStart_lineContext)
}

func (s *Http_messageContext) AllCRLF() []antlr.TerminalNode {
	return s.GetTokens(httpParserCRLF)
}

func (s *Http_messageContext) CRLF(i int) antlr.TerminalNode {
	return s.GetToken(httpParserCRLF, i)
}

func (s *Http_messageContext) AllHeader_field() []IHeader_fieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHeader_fieldContext)(nil)).Elem())
	var tst = make([]IHeader_fieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHeader_fieldContext)
		}
	}

	return tst
}

func (s *Http_messageContext) Header_field(i int) IHeader_fieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeader_fieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHeader_fieldContext)
}

func (s *Http_messageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Http_messageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Http_messageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterHttp_message(s)
	}
}

func (s *Http_messageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitHttp_message(s)
	}
}

func (p *httpParser) Http_message() (localctx IHttp_messageContext) {
	this := p
	_ = this

	localctx = NewHttp_messageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, httpParserRULE_http_message)
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
		p.SetState(48)
		p.Start_line()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(httpParserALPHA-11))|(1<<(httpParserDIGIT-11))|(1<<(httpParserMinus-11))|(1<<(httpParserDot-11))|(1<<(httpParserUnderscore-11))|(1<<(httpParserTilde-11))|(1<<(httpParserExclamationMark-11))|(1<<(httpParserDollarSign-11))|(1<<(httpParserHashtag-11))|(1<<(httpParserAmpersand-11))|(1<<(httpParserPercent-11))|(1<<(httpParserSQuote-11))|(1<<(httpParserStar-11))|(1<<(httpParserPlus-11))|(1<<(httpParserCaret-11))|(1<<(httpParserBackQuote-11))|(1<<(httpParserVBar-11)))) != 0 {
		{
			p.SetState(49)
			p.Header_field()
		}
		{
			p.SetState(50)
			p.Match(httpParserCRLF)
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(httpParserCRLF)
	}

	return localctx
}

// IStart_lineContext is an interface to support dynamic dispatch.
type IStart_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStart_lineContext differentiates from other interfaces.
	IsStart_lineContext()
}

type Start_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStart_lineContext() *Start_lineContext {
	var p = new(Start_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_start_line
	return p
}

func (*Start_lineContext) IsStart_lineContext() {}

func NewStart_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Start_lineContext {
	var p = new(Start_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_start_line

	return p
}

func (s *Start_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Start_lineContext) Request_line() IRequest_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequest_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequest_lineContext)
}

func (s *Start_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Start_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Start_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterStart_line(s)
	}
}

func (s *Start_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitStart_line(s)
	}
}

func (p *httpParser) Start_line() (localctx IStart_lineContext) {
	this := p
	_ = this

	localctx = NewStart_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, httpParserRULE_start_line)

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
		p.Request_line()
	}

	return localctx
}

// IRequest_lineContext is an interface to support dynamic dispatch.
type IRequest_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequest_lineContext differentiates from other interfaces.
	IsRequest_lineContext()
}

type Request_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequest_lineContext() *Request_lineContext {
	var p = new(Request_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_request_line
	return p
}

func (*Request_lineContext) IsRequest_lineContext() {}

func NewRequest_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Request_lineContext {
	var p = new(Request_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_request_line

	return p
}

func (s *Request_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Request_lineContext) Method() IMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *Request_lineContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(httpParserSP)
}

func (s *Request_lineContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(httpParserSP, i)
}

func (s *Request_lineContext) Request_target() IRequest_targetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequest_targetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequest_targetContext)
}

func (s *Request_lineContext) Http_version() IHttp_versionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHttp_versionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHttp_versionContext)
}

func (s *Request_lineContext) CRLF() antlr.TerminalNode {
	return s.GetToken(httpParserCRLF, 0)
}

func (s *Request_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Request_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Request_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterRequest_line(s)
	}
}

func (s *Request_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitRequest_line(s)
	}
}

func (p *httpParser) Request_line() (localctx IRequest_lineContext) {
	this := p
	_ = this

	localctx = NewRequest_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, httpParserRULE_request_line)

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
		p.SetState(61)
		p.Method()
	}
	{
		p.SetState(62)
		p.Match(httpParserSP)
	}
	{
		p.SetState(63)
		p.Request_target()
	}
	{
		p.SetState(64)
		p.Match(httpParserSP)
	}
	{
		p.SetState(65)
		p.Http_version()
	}
	{
		p.SetState(66)
		p.Match(httpParserCRLF)
	}

	return localctx
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_method
	return p
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }
func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *httpParser) Method() (localctx IMethodContext) {
	this := p
	_ = this

	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, httpParserRULE_method)
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
		p.SetState(68)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpParserT__0)|(1<<httpParserT__1)|(1<<httpParserT__2)|(1<<httpParserT__3)|(1<<httpParserT__4)|(1<<httpParserT__5)|(1<<httpParserT__6)|(1<<httpParserT__7))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRequest_targetContext is an interface to support dynamic dispatch.
type IRequest_targetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequest_targetContext differentiates from other interfaces.
	IsRequest_targetContext()
}

type Request_targetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequest_targetContext() *Request_targetContext {
	var p = new(Request_targetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_request_target
	return p
}

func (*Request_targetContext) IsRequest_targetContext() {}

func NewRequest_targetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Request_targetContext {
	var p = new(Request_targetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_request_target

	return p
}

func (s *Request_targetContext) GetParser() antlr.Parser { return s.parser }

func (s *Request_targetContext) Origin_form() IOrigin_formContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrigin_formContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrigin_formContext)
}

func (s *Request_targetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Request_targetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Request_targetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterRequest_target(s)
	}
}

func (s *Request_targetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitRequest_target(s)
	}
}

func (p *httpParser) Request_target() (localctx IRequest_targetContext) {
	this := p
	_ = this

	localctx = NewRequest_targetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, httpParserRULE_request_target)

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
		p.Origin_form()
	}

	return localctx
}

// IOrigin_formContext is an interface to support dynamic dispatch.
type IOrigin_formContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrigin_formContext differentiates from other interfaces.
	IsOrigin_formContext()
}

type Origin_formContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrigin_formContext() *Origin_formContext {
	var p = new(Origin_formContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_origin_form
	return p
}

func (*Origin_formContext) IsOrigin_formContext() {}

func NewOrigin_formContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Origin_formContext {
	var p = new(Origin_formContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_origin_form

	return p
}

func (s *Origin_formContext) GetParser() antlr.Parser { return s.parser }

func (s *Origin_formContext) Absolute_path() IAbsolute_pathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAbsolute_pathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAbsolute_pathContext)
}

func (s *Origin_formContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(httpParserQuestionMark, 0)
}

func (s *Origin_formContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *Origin_formContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Origin_formContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Origin_formContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterOrigin_form(s)
	}
}

func (s *Origin_formContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitOrigin_form(s)
	}
}

func (p *httpParser) Origin_form() (localctx IOrigin_formContext) {
	this := p
	_ = this

	localctx = NewOrigin_formContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, httpParserRULE_origin_form)
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
		p.SetState(72)
		p.Absolute_path()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpParserQuestionMark {
		{
			p.SetState(73)
			p.Match(httpParserQuestionMark)
		}
		{
			p.SetState(74)
			p.Query()
		}

	}

	return localctx
}

// IAbsolute_pathContext is an interface to support dynamic dispatch.
type IAbsolute_pathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbsolute_pathContext differentiates from other interfaces.
	IsAbsolute_pathContext()
}

type Absolute_pathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsolute_pathContext() *Absolute_pathContext {
	var p = new(Absolute_pathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_absolute_path
	return p
}

func (*Absolute_pathContext) IsAbsolute_pathContext() {}

func NewAbsolute_pathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Absolute_pathContext {
	var p = new(Absolute_pathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_absolute_path

	return p
}

func (s *Absolute_pathContext) GetParser() antlr.Parser { return s.parser }

func (s *Absolute_pathContext) AllSlash() []antlr.TerminalNode {
	return s.GetTokens(httpParserSlash)
}

func (s *Absolute_pathContext) Slash(i int) antlr.TerminalNode {
	return s.GetToken(httpParserSlash, i)
}

func (s *Absolute_pathContext) AllSegment() []ISegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISegmentContext)(nil)).Elem())
	var tst = make([]ISegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISegmentContext)
		}
	}

	return tst
}

func (s *Absolute_pathContext) Segment(i int) ISegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISegmentContext)
}

func (s *Absolute_pathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Absolute_pathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Absolute_pathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterAbsolute_path(s)
	}
}

func (s *Absolute_pathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitAbsolute_path(s)
	}
}

func (p *httpParser) Absolute_path() (localctx IAbsolute_pathContext) {
	this := p
	_ = this

	localctx = NewAbsolute_pathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, httpParserRULE_absolute_path)
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpParserSlash {
		{
			p.SetState(77)
			p.Match(httpParserSlash)
		}
		{
			p.SetState(78)
			p.Segment()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISegmentContext is an interface to support dynamic dispatch.
type ISegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentContext differentiates from other interfaces.
	IsSegmentContext()
}

type SegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentContext() *SegmentContext {
	var p = new(SegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_segment
	return p
}

func (*SegmentContext) IsSegmentContext() {}

func NewSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentContext {
	var p = new(SegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_segment

	return p
}

func (s *SegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentContext) AllPchar() []IPcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPcharContext)(nil)).Elem())
	var tst = make([]IPcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPcharContext)
		}
	}

	return tst
}

func (s *SegmentContext) Pchar(i int) IPcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPcharContext)
}

func (s *SegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterSegment(s)
	}
}

func (s *SegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitSegment(s)
	}
}

func (p *httpParser) Segment() (localctx ISegmentContext) {
	this := p
	_ = this

	localctx = NewSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, httpParserRULE_segment)
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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(httpParserALPHA-11))|(1<<(httpParserDIGIT-11))|(1<<(httpParserPct_encoded-11))|(1<<(httpParserLColumn-11))|(1<<(httpParserRColumn-11))|(1<<(httpParserSemiColon-11))|(1<<(httpParserEquals-11))|(1<<(httpParserPeriod-11))|(1<<(httpParserMinus-11))|(1<<(httpParserDot-11))|(1<<(httpParserUnderscore-11))|(1<<(httpParserTilde-11))|(1<<(httpParserExclamationMark-11))|(1<<(httpParserColon-11))|(1<<(httpParserAt-11))|(1<<(httpParserDollarSign-11))|(1<<(httpParserAmpersand-11))|(1<<(httpParserSQuote-11))|(1<<(httpParserStar-11))|(1<<(httpParserPlus-11)))) != 0 {
		{
			p.SetState(83)
			p.Pchar()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) AllPchar() []IPcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPcharContext)(nil)).Elem())
	var tst = make([]IPcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPcharContext)
		}
	}

	return tst
}

func (s *QueryContext) Pchar(i int) IPcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPcharContext)
}

func (s *QueryContext) AllSlash() []antlr.TerminalNode {
	return s.GetTokens(httpParserSlash)
}

func (s *QueryContext) Slash(i int) antlr.TerminalNode {
	return s.GetToken(httpParserSlash, i)
}

func (s *QueryContext) AllQuestionMark() []antlr.TerminalNode {
	return s.GetTokens(httpParserQuestionMark)
}

func (s *QueryContext) QuestionMark(i int) antlr.TerminalNode {
	return s.GetToken(httpParserQuestionMark, i)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *httpParser) Query() (localctx IQueryContext) {
	this := p
	_ = this

	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, httpParserRULE_query)
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
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(httpParserALPHA-11))|(1<<(httpParserDIGIT-11))|(1<<(httpParserPct_encoded-11))|(1<<(httpParserLColumn-11))|(1<<(httpParserRColumn-11))|(1<<(httpParserSemiColon-11))|(1<<(httpParserEquals-11))|(1<<(httpParserPeriod-11))|(1<<(httpParserMinus-11))|(1<<(httpParserDot-11))|(1<<(httpParserUnderscore-11))|(1<<(httpParserTilde-11))|(1<<(httpParserQuestionMark-11))|(1<<(httpParserSlash-11))|(1<<(httpParserExclamationMark-11))|(1<<(httpParserColon-11))|(1<<(httpParserAt-11))|(1<<(httpParserDollarSign-11))|(1<<(httpParserAmpersand-11))|(1<<(httpParserSQuote-11))|(1<<(httpParserStar-11))|(1<<(httpParserPlus-11)))) != 0 {
		p.SetState(92)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case httpParserALPHA, httpParserDIGIT, httpParserPct_encoded, httpParserLColumn, httpParserRColumn, httpParserSemiColon, httpParserEquals, httpParserPeriod, httpParserMinus, httpParserDot, httpParserUnderscore, httpParserTilde, httpParserExclamationMark, httpParserColon, httpParserAt, httpParserDollarSign, httpParserAmpersand, httpParserSQuote, httpParserStar, httpParserPlus:
			{
				p.SetState(89)
				p.Pchar()
			}

		case httpParserSlash:
			{
				p.SetState(90)
				p.Match(httpParserSlash)
			}

		case httpParserQuestionMark:
			{
				p.SetState(91)
				p.Match(httpParserQuestionMark)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHttp_versionContext is an interface to support dynamic dispatch.
type IHttp_versionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHttp_versionContext differentiates from other interfaces.
	IsHttp_versionContext()
}

type Http_versionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttp_versionContext() *Http_versionContext {
	var p = new(Http_versionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_http_version
	return p
}

func (*Http_versionContext) IsHttp_versionContext() {}

func NewHttp_versionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Http_versionContext {
	var p = new(Http_versionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_http_version

	return p
}

func (s *Http_versionContext) GetParser() antlr.Parser { return s.parser }

func (s *Http_versionContext) Http_name() IHttp_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHttp_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHttp_nameContext)
}

func (s *Http_versionContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(httpParserDIGIT)
}

func (s *Http_versionContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(httpParserDIGIT, i)
}

func (s *Http_versionContext) Dot() antlr.TerminalNode {
	return s.GetToken(httpParserDot, 0)
}

func (s *Http_versionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Http_versionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Http_versionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterHttp_version(s)
	}
}

func (s *Http_versionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitHttp_version(s)
	}
}

func (p *httpParser) Http_version() (localctx IHttp_versionContext) {
	this := p
	_ = this

	localctx = NewHttp_versionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, httpParserRULE_http_version)

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
		p.SetState(97)
		p.Http_name()
	}
	{
		p.SetState(98)
		p.Match(httpParserDIGIT)
	}
	{
		p.SetState(99)
		p.Match(httpParserDot)
	}
	{
		p.SetState(100)
		p.Match(httpParserDIGIT)
	}

	return localctx
}

// IHttp_nameContext is an interface to support dynamic dispatch.
type IHttp_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHttp_nameContext differentiates from other interfaces.
	IsHttp_nameContext()
}

type Http_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHttp_nameContext() *Http_nameContext {
	var p = new(Http_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_http_name
	return p
}

func (*Http_nameContext) IsHttp_nameContext() {}

func NewHttp_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Http_nameContext {
	var p = new(Http_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_http_name

	return p
}

func (s *Http_nameContext) GetParser() antlr.Parser { return s.parser }
func (s *Http_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Http_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Http_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterHttp_name(s)
	}
}

func (s *Http_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitHttp_name(s)
	}
}

func (p *httpParser) Http_name() (localctx IHttp_nameContext) {
	this := p
	_ = this

	localctx = NewHttp_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, httpParserRULE_http_name)

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
		p.SetState(102)
		p.Match(httpParserT__8)
	}

	return localctx
}

// IHeader_fieldContext is an interface to support dynamic dispatch.
type IHeader_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeader_fieldContext differentiates from other interfaces.
	IsHeader_fieldContext()
}

type Header_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeader_fieldContext() *Header_fieldContext {
	var p = new(Header_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_header_field
	return p
}

func (*Header_fieldContext) IsHeader_fieldContext() {}

func NewHeader_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Header_fieldContext {
	var p = new(Header_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_header_field

	return p
}

func (s *Header_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Header_fieldContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Header_fieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(httpParserColon, 0)
}

func (s *Header_fieldContext) Field_value() IField_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_valueContext)
}

func (s *Header_fieldContext) AllOWS() []antlr.TerminalNode {
	return s.GetTokens(httpParserOWS)
}

func (s *Header_fieldContext) OWS(i int) antlr.TerminalNode {
	return s.GetToken(httpParserOWS, i)
}

func (s *Header_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Header_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Header_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterHeader_field(s)
	}
}

func (s *Header_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitHeader_field(s)
	}
}

func (p *httpParser) Header_field() (localctx IHeader_fieldContext) {
	this := p
	_ = this

	localctx = NewHeader_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, httpParserRULE_header_field)
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
		p.SetState(104)
		p.Field_name()
	}
	{
		p.SetState(105)
		p.Match(httpParserColon)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpParserOWS {
		{
			p.SetState(106)
			p.Match(httpParserOWS)
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(112)
		p.Field_value()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == httpParserOWS {
		{
			p.SetState(113)
			p.Match(httpParserOWS)
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IField_nameContext is an interface to support dynamic dispatch.
type IField_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_nameContext differentiates from other interfaces.
	IsField_nameContext()
}

type Field_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_nameContext() *Field_nameContext {
	var p = new(Field_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_field_name
	return p
}

func (*Field_nameContext) IsField_nameContext() {}

func NewField_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_nameContext {
	var p = new(Field_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_field_name

	return p
}

func (s *Field_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_nameContext) Token() ITokenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITokenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITokenContext)
}

func (s *Field_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterField_name(s)
	}
}

func (s *Field_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitField_name(s)
	}
}

func (p *httpParser) Field_name() (localctx IField_nameContext) {
	this := p
	_ = this

	localctx = NewField_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, httpParserRULE_field_name)

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
		p.SetState(119)
		p.Token()
	}

	return localctx
}

// ITokenContext is an interface to support dynamic dispatch.
type ITokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTokenContext differentiates from other interfaces.
	IsTokenContext()
}

type TokenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTokenContext() *TokenContext {
	var p = new(TokenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_token
	return p
}

func (*TokenContext) IsTokenContext() {}

func NewTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TokenContext {
	var p = new(TokenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_token

	return p
}

func (s *TokenContext) GetParser() antlr.Parser { return s.parser }

func (s *TokenContext) AllTchar() []ITcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITcharContext)(nil)).Elem())
	var tst = make([]ITcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITcharContext)
		}
	}

	return tst
}

func (s *TokenContext) Tchar(i int) ITcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITcharContext)
}

func (s *TokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterToken(s)
	}
}

func (s *TokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitToken(s)
	}
}

func (p *httpParser) Token() (localctx ITokenContext) {
	this := p
	_ = this

	localctx = NewTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, httpParserRULE_token)
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
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(httpParserALPHA-11))|(1<<(httpParserDIGIT-11))|(1<<(httpParserMinus-11))|(1<<(httpParserDot-11))|(1<<(httpParserUnderscore-11))|(1<<(httpParserTilde-11))|(1<<(httpParserExclamationMark-11))|(1<<(httpParserDollarSign-11))|(1<<(httpParserHashtag-11))|(1<<(httpParserAmpersand-11))|(1<<(httpParserPercent-11))|(1<<(httpParserSQuote-11))|(1<<(httpParserStar-11))|(1<<(httpParserPlus-11))|(1<<(httpParserCaret-11))|(1<<(httpParserBackQuote-11))|(1<<(httpParserVBar-11)))) != 0) {
		{
			p.SetState(121)
			p.Tchar()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IField_valueContext is an interface to support dynamic dispatch.
type IField_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_valueContext differentiates from other interfaces.
	IsField_valueContext()
}

type Field_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_valueContext() *Field_valueContext {
	var p = new(Field_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_field_value
	return p
}

func (*Field_valueContext) IsField_valueContext() {}

func NewField_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_valueContext {
	var p = new(Field_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_field_value

	return p
}

func (s *Field_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_valueContext) AllField_content() []IField_contentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IField_contentContext)(nil)).Elem())
	var tst = make([]IField_contentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IField_contentContext)
		}
	}

	return tst
}

func (s *Field_valueContext) Field_content(i int) IField_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_contentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IField_contentContext)
}

func (s *Field_valueContext) AllObs_fold() []IObs_foldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObs_foldContext)(nil)).Elem())
	var tst = make([]IObs_foldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObs_foldContext)
		}
	}

	return tst
}

func (s *Field_valueContext) Obs_fold(i int) IObs_foldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObs_foldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObs_foldContext)
}

func (s *Field_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterField_value(s)
	}
}

func (s *Field_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitField_value(s)
	}
}

func (p *httpParser) Field_value() (localctx IField_valueContext) {
	this := p
	_ = this

	localctx = NewField_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, httpParserRULE_field_value)

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
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(128)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case httpParserALPHA, httpParserDIGIT, httpParserVCHAR, httpParserOBS_TEXT:
				{
					p.SetState(126)
					p.Field_content()
				}

			case httpParserCRLF:
				{
					p.SetState(127)
					p.Obs_fold()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IField_contentContext is an interface to support dynamic dispatch.
type IField_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_contentContext differentiates from other interfaces.
	IsField_contentContext()
}

type Field_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_contentContext() *Field_contentContext {
	var p = new(Field_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_field_content
	return p
}

func (*Field_contentContext) IsField_contentContext() {}

func NewField_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_contentContext {
	var p = new(Field_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_field_content

	return p
}

func (s *Field_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_contentContext) AllField_vchar() []IField_vcharContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IField_vcharContext)(nil)).Elem())
	var tst = make([]IField_vcharContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IField_vcharContext)
		}
	}

	return tst
}

func (s *Field_contentContext) Field_vchar(i int) IField_vcharContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_vcharContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IField_vcharContext)
}

func (s *Field_contentContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(httpParserSP)
}

func (s *Field_contentContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(httpParserSP, i)
}

func (s *Field_contentContext) AllHTAB() []antlr.TerminalNode {
	return s.GetTokens(httpParserHTAB)
}

func (s *Field_contentContext) HTAB(i int) antlr.TerminalNode {
	return s.GetToken(httpParserHTAB, i)
}

func (s *Field_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterField_content(s)
	}
}

func (s *Field_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitField_content(s)
	}
}

func (p *httpParser) Field_content() (localctx IField_contentContext) {
	this := p
	_ = this

	localctx = NewField_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, httpParserRULE_field_content)
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
		p.Field_vchar()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == httpParserSP || _la == httpParserHTAB {
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == httpParserSP || _la == httpParserHTAB {
			{
				p.SetState(133)
				_la = p.GetTokenStream().LA(1)

				if !(_la == httpParserSP || _la == httpParserHTAB) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(138)
			p.Field_vchar()
		}

	}

	return localctx
}

// IField_vcharContext is an interface to support dynamic dispatch.
type IField_vcharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_vcharContext differentiates from other interfaces.
	IsField_vcharContext()
}

type Field_vcharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_vcharContext() *Field_vcharContext {
	var p = new(Field_vcharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_field_vchar
	return p
}

func (*Field_vcharContext) IsField_vcharContext() {}

func NewField_vcharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_vcharContext {
	var p = new(Field_vcharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_field_vchar

	return p
}

func (s *Field_vcharContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_vcharContext) VCHAR() IVCHARContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVCHARContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVCHARContext)
}

func (s *Field_vcharContext) Obs_text() IObs_textContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObs_textContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObs_textContext)
}

func (s *Field_vcharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_vcharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_vcharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterField_vchar(s)
	}
}

func (s *Field_vcharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitField_vchar(s)
	}
}

func (p *httpParser) Field_vchar() (localctx IField_vcharContext) {
	this := p
	_ = this

	localctx = NewField_vcharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, httpParserRULE_field_vchar)

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

	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case httpParserALPHA, httpParserDIGIT, httpParserVCHAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.VCHAR()
		}

	case httpParserOBS_TEXT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Obs_text()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObs_textContext is an interface to support dynamic dispatch.
type IObs_textContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObs_textContext differentiates from other interfaces.
	IsObs_textContext()
}

type Obs_textContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObs_textContext() *Obs_textContext {
	var p = new(Obs_textContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_obs_text
	return p
}

func (*Obs_textContext) IsObs_textContext() {}

func NewObs_textContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Obs_textContext {
	var p = new(Obs_textContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_obs_text

	return p
}

func (s *Obs_textContext) GetParser() antlr.Parser { return s.parser }

func (s *Obs_textContext) OBS_TEXT() antlr.TerminalNode {
	return s.GetToken(httpParserOBS_TEXT, 0)
}

func (s *Obs_textContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Obs_textContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Obs_textContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterObs_text(s)
	}
}

func (s *Obs_textContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitObs_text(s)
	}
}

func (p *httpParser) Obs_text() (localctx IObs_textContext) {
	this := p
	_ = this

	localctx = NewObs_textContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, httpParserRULE_obs_text)

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
		p.SetState(145)
		p.Match(httpParserOBS_TEXT)
	}

	return localctx
}

// IObs_foldContext is an interface to support dynamic dispatch.
type IObs_foldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObs_foldContext differentiates from other interfaces.
	IsObs_foldContext()
}

type Obs_foldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObs_foldContext() *Obs_foldContext {
	var p = new(Obs_foldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_obs_fold
	return p
}

func (*Obs_foldContext) IsObs_foldContext() {}

func NewObs_foldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Obs_foldContext {
	var p = new(Obs_foldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_obs_fold

	return p
}

func (s *Obs_foldContext) GetParser() antlr.Parser { return s.parser }

func (s *Obs_foldContext) CRLF() antlr.TerminalNode {
	return s.GetToken(httpParserCRLF, 0)
}

func (s *Obs_foldContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(httpParserSP)
}

func (s *Obs_foldContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(httpParserSP, i)
}

func (s *Obs_foldContext) AllHTAB() []antlr.TerminalNode {
	return s.GetTokens(httpParserHTAB)
}

func (s *Obs_foldContext) HTAB(i int) antlr.TerminalNode {
	return s.GetToken(httpParserHTAB, i)
}

func (s *Obs_foldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Obs_foldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Obs_foldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterObs_fold(s)
	}
}

func (s *Obs_foldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitObs_fold(s)
	}
}

func (p *httpParser) Obs_fold() (localctx IObs_foldContext) {
	this := p
	_ = this

	localctx = NewObs_foldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, httpParserRULE_obs_fold)
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
		p.SetState(147)
		p.Match(httpParserCRLF)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == httpParserSP || _la == httpParserHTAB {
		{
			p.SetState(148)
			_la = p.GetTokenStream().LA(1)

			if !(_la == httpParserSP || _la == httpParserHTAB) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPcharContext is an interface to support dynamic dispatch.
type IPcharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPcharContext differentiates from other interfaces.
	IsPcharContext()
}

type PcharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPcharContext() *PcharContext {
	var p = new(PcharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_pchar
	return p
}

func (*PcharContext) IsPcharContext() {}

func NewPcharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PcharContext {
	var p = new(PcharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_pchar

	return p
}

func (s *PcharContext) GetParser() antlr.Parser { return s.parser }

func (s *PcharContext) Unreserved() IUnreservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnreservedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnreservedContext)
}

func (s *PcharContext) Pct_encoded() antlr.TerminalNode {
	return s.GetToken(httpParserPct_encoded, 0)
}

func (s *PcharContext) Sub_delims() ISub_delimsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISub_delimsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISub_delimsContext)
}

func (s *PcharContext) Colon() antlr.TerminalNode {
	return s.GetToken(httpParserColon, 0)
}

func (s *PcharContext) At() antlr.TerminalNode {
	return s.GetToken(httpParserAt, 0)
}

func (s *PcharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PcharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PcharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterPchar(s)
	}
}

func (s *PcharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitPchar(s)
	}
}

func (p *httpParser) Pchar() (localctx IPcharContext) {
	this := p
	_ = this

	localctx = NewPcharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, httpParserRULE_pchar)

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

	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case httpParserALPHA, httpParserDIGIT, httpParserMinus, httpParserDot, httpParserUnderscore, httpParserTilde:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Unreserved()
		}

	case httpParserPct_encoded:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Match(httpParserPct_encoded)
		}

	case httpParserLColumn, httpParserRColumn, httpParserSemiColon, httpParserEquals, httpParserPeriod, httpParserExclamationMark, httpParserDollarSign, httpParserAmpersand, httpParserSQuote, httpParserStar, httpParserPlus:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.Sub_delims()
		}

	case httpParserColon:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(156)
			p.Match(httpParserColon)
		}

	case httpParserAt:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(157)
			p.Match(httpParserAt)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnreservedContext is an interface to support dynamic dispatch.
type IUnreservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnreservedContext differentiates from other interfaces.
	IsUnreservedContext()
}

type UnreservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnreservedContext() *UnreservedContext {
	var p = new(UnreservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_unreserved
	return p
}

func (*UnreservedContext) IsUnreservedContext() {}

func NewUnreservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnreservedContext {
	var p = new(UnreservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_unreserved

	return p
}

func (s *UnreservedContext) GetParser() antlr.Parser { return s.parser }

func (s *UnreservedContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(httpParserALPHA, 0)
}

func (s *UnreservedContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(httpParserDIGIT, 0)
}

func (s *UnreservedContext) Minus() antlr.TerminalNode {
	return s.GetToken(httpParserMinus, 0)
}

func (s *UnreservedContext) Dot() antlr.TerminalNode {
	return s.GetToken(httpParserDot, 0)
}

func (s *UnreservedContext) Underscore() antlr.TerminalNode {
	return s.GetToken(httpParserUnderscore, 0)
}

func (s *UnreservedContext) Tilde() antlr.TerminalNode {
	return s.GetToken(httpParserTilde, 0)
}

func (s *UnreservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnreservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnreservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterUnreserved(s)
	}
}

func (s *UnreservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitUnreserved(s)
	}
}

func (p *httpParser) Unreserved() (localctx IUnreservedContext) {
	this := p
	_ = this

	localctx = NewUnreservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, httpParserRULE_unreserved)
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
		p.SetState(160)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<httpParserALPHA)|(1<<httpParserDIGIT)|(1<<httpParserMinus)|(1<<httpParserDot)|(1<<httpParserUnderscore)|(1<<httpParserTilde))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISub_delimsContext is an interface to support dynamic dispatch.
type ISub_delimsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSub_delimsContext differentiates from other interfaces.
	IsSub_delimsContext()
}

type Sub_delimsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySub_delimsContext() *Sub_delimsContext {
	var p = new(Sub_delimsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_sub_delims
	return p
}

func (*Sub_delimsContext) IsSub_delimsContext() {}

func NewSub_delimsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sub_delimsContext {
	var p = new(Sub_delimsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_sub_delims

	return p
}

func (s *Sub_delimsContext) GetParser() antlr.Parser { return s.parser }

func (s *Sub_delimsContext) ExclamationMark() antlr.TerminalNode {
	return s.GetToken(httpParserExclamationMark, 0)
}

func (s *Sub_delimsContext) DollarSign() antlr.TerminalNode {
	return s.GetToken(httpParserDollarSign, 0)
}

func (s *Sub_delimsContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(httpParserAmpersand, 0)
}

func (s *Sub_delimsContext) SQuote() antlr.TerminalNode {
	return s.GetToken(httpParserSQuote, 0)
}

func (s *Sub_delimsContext) LColumn() antlr.TerminalNode {
	return s.GetToken(httpParserLColumn, 0)
}

func (s *Sub_delimsContext) RColumn() antlr.TerminalNode {
	return s.GetToken(httpParserRColumn, 0)
}

func (s *Sub_delimsContext) Star() antlr.TerminalNode {
	return s.GetToken(httpParserStar, 0)
}

func (s *Sub_delimsContext) Plus() antlr.TerminalNode {
	return s.GetToken(httpParserPlus, 0)
}

func (s *Sub_delimsContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(httpParserSemiColon, 0)
}

func (s *Sub_delimsContext) Period() antlr.TerminalNode {
	return s.GetToken(httpParserPeriod, 0)
}

func (s *Sub_delimsContext) Equals() antlr.TerminalNode {
	return s.GetToken(httpParserEquals, 0)
}

func (s *Sub_delimsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sub_delimsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sub_delimsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterSub_delims(s)
	}
}

func (s *Sub_delimsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitSub_delims(s)
	}
}

func (p *httpParser) Sub_delims() (localctx ISub_delimsContext) {
	this := p
	_ = this

	localctx = NewSub_delimsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, httpParserRULE_sub_delims)
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
		p.SetState(162)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-15)&-(0x1f+1)) == 0 && ((1<<uint((_la-15)))&((1<<(httpParserLColumn-15))|(1<<(httpParserRColumn-15))|(1<<(httpParserSemiColon-15))|(1<<(httpParserEquals-15))|(1<<(httpParserPeriod-15))|(1<<(httpParserExclamationMark-15))|(1<<(httpParserDollarSign-15))|(1<<(httpParserAmpersand-15))|(1<<(httpParserSQuote-15))|(1<<(httpParserStar-15))|(1<<(httpParserPlus-15)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITcharContext is an interface to support dynamic dispatch.
type ITcharContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTcharContext differentiates from other interfaces.
	IsTcharContext()
}

type TcharContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTcharContext() *TcharContext {
	var p = new(TcharContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_tchar
	return p
}

func (*TcharContext) IsTcharContext() {}

func NewTcharContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TcharContext {
	var p = new(TcharContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_tchar

	return p
}

func (s *TcharContext) GetParser() antlr.Parser { return s.parser }

func (s *TcharContext) ExclamationMark() antlr.TerminalNode {
	return s.GetToken(httpParserExclamationMark, 0)
}

func (s *TcharContext) DollarSign() antlr.TerminalNode {
	return s.GetToken(httpParserDollarSign, 0)
}

func (s *TcharContext) Hashtag() antlr.TerminalNode {
	return s.GetToken(httpParserHashtag, 0)
}

func (s *TcharContext) Percent() antlr.TerminalNode {
	return s.GetToken(httpParserPercent, 0)
}

func (s *TcharContext) Ampersand() antlr.TerminalNode {
	return s.GetToken(httpParserAmpersand, 0)
}

func (s *TcharContext) SQuote() antlr.TerminalNode {
	return s.GetToken(httpParserSQuote, 0)
}

func (s *TcharContext) Star() antlr.TerminalNode {
	return s.GetToken(httpParserStar, 0)
}

func (s *TcharContext) Plus() antlr.TerminalNode {
	return s.GetToken(httpParserPlus, 0)
}

func (s *TcharContext) Minus() antlr.TerminalNode {
	return s.GetToken(httpParserMinus, 0)
}

func (s *TcharContext) Dot() antlr.TerminalNode {
	return s.GetToken(httpParserDot, 0)
}

func (s *TcharContext) Caret() antlr.TerminalNode {
	return s.GetToken(httpParserCaret, 0)
}

func (s *TcharContext) Underscore() antlr.TerminalNode {
	return s.GetToken(httpParserUnderscore, 0)
}

func (s *TcharContext) BackQuote() antlr.TerminalNode {
	return s.GetToken(httpParserBackQuote, 0)
}

func (s *TcharContext) VBar() antlr.TerminalNode {
	return s.GetToken(httpParserVBar, 0)
}

func (s *TcharContext) Tilde() antlr.TerminalNode {
	return s.GetToken(httpParserTilde, 0)
}

func (s *TcharContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(httpParserDIGIT, 0)
}

func (s *TcharContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(httpParserALPHA, 0)
}

func (s *TcharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TcharContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TcharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterTchar(s)
	}
}

func (s *TcharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitTchar(s)
	}
}

func (p *httpParser) Tchar() (localctx ITcharContext) {
	this := p
	_ = this

	localctx = NewTcharContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, httpParserRULE_tchar)
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
		p.SetState(164)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(httpParserALPHA-11))|(1<<(httpParserDIGIT-11))|(1<<(httpParserMinus-11))|(1<<(httpParserDot-11))|(1<<(httpParserUnderscore-11))|(1<<(httpParserTilde-11))|(1<<(httpParserExclamationMark-11))|(1<<(httpParserDollarSign-11))|(1<<(httpParserHashtag-11))|(1<<(httpParserAmpersand-11))|(1<<(httpParserPercent-11))|(1<<(httpParserSQuote-11))|(1<<(httpParserStar-11))|(1<<(httpParserPlus-11))|(1<<(httpParserCaret-11))|(1<<(httpParserBackQuote-11))|(1<<(httpParserVBar-11)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IVCHARContext is an interface to support dynamic dispatch.
type IVCHARContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVCHARContext differentiates from other interfaces.
	IsVCHARContext()
}

type VCHARContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVCHARContext() *VCHARContext {
	var p = new(VCHARContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = httpParserRULE_vCHAR
	return p
}

func (*VCHARContext) IsVCHARContext() {}

func NewVCHARContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VCHARContext {
	var p = new(VCHARContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = httpParserRULE_vCHAR

	return p
}

func (s *VCHARContext) GetParser() antlr.Parser { return s.parser }

func (s *VCHARContext) ALPHA() antlr.TerminalNode {
	return s.GetToken(httpParserALPHA, 0)
}

func (s *VCHARContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(httpParserDIGIT, 0)
}

func (s *VCHARContext) VCHAR() antlr.TerminalNode {
	return s.GetToken(httpParserVCHAR, 0)
}

func (s *VCHARContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VCHARContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VCHARContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.EnterVCHAR(s)
	}
}

func (s *VCHARContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(httpListener); ok {
		listenerT.ExitVCHAR(s)
	}
}

func (p *httpParser) VCHAR() (localctx IVCHARContext) {
	this := p
	_ = this

	localctx = NewVCHARContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, httpParserRULE_vCHAR)
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
		p.SetState(166)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(httpParserALPHA-11))|(1<<(httpParserDIGIT-11))|(1<<(httpParserVCHAR-11)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
