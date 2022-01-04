// Code generated from apt.g4 by ANTLR 4.9.3. DO NOT EDIT.

package apt // apt
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 178,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 3,
	2, 5, 2, 58, 10, 2, 3, 2, 3, 2, 3, 2, 5, 2, 63, 10, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 5, 2, 70, 10, 2, 3, 3, 3, 3, 3, 4, 6, 4, 75, 10, 4, 13, 4,
	14, 4, 76, 3, 5, 3, 5, 3, 5, 3, 6, 6, 6, 83, 10, 6, 13, 6, 14, 6, 84, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 7, 9, 96, 10, 9, 12,
	9, 14, 9, 99, 11, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 7, 14, 115, 10, 14, 12, 14,
	14, 14, 118, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 5, 16, 126,
	10, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 6, 19, 134, 10, 19, 13,
	19, 14, 19, 135, 3, 20, 3, 20, 3, 20, 5, 20, 141, 10, 20, 3, 21, 6, 21,
	144, 10, 21, 13, 21, 14, 21, 145, 3, 22, 3, 22, 5, 22, 150, 10, 22, 3,
	23, 6, 23, 153, 10, 23, 13, 23, 14, 23, 154, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 5, 26, 167, 10, 26, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 2, 2, 29, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 2, 4, 4, 2, 6, 6, 17, 17, 6, 2, 3, 6, 9, 9,
	14, 15, 17, 17, 2, 165, 2, 57, 3, 2, 2, 2, 4, 71, 3, 2, 2, 2, 6, 74, 3,
	2, 2, 2, 8, 78, 3, 2, 2, 2, 10, 82, 3, 2, 2, 2, 12, 86, 3, 2, 2, 2, 14,
	91, 3, 2, 2, 2, 16, 97, 3, 2, 2, 2, 18, 100, 3, 2, 2, 2, 20, 103, 3, 2,
	2, 2, 22, 107, 3, 2, 2, 2, 24, 110, 3, 2, 2, 2, 26, 116, 3, 2, 2, 2, 28,
	119, 3, 2, 2, 2, 30, 125, 3, 2, 2, 2, 32, 127, 3, 2, 2, 2, 34, 130, 3,
	2, 2, 2, 36, 133, 3, 2, 2, 2, 38, 140, 3, 2, 2, 2, 40, 143, 3, 2, 2, 2,
	42, 149, 3, 2, 2, 2, 44, 152, 3, 2, 2, 2, 46, 156, 3, 2, 2, 2, 48, 160,
	3, 2, 2, 2, 50, 166, 3, 2, 2, 2, 52, 168, 3, 2, 2, 2, 54, 172, 3, 2, 2,
	2, 56, 58, 5, 32, 17, 2, 57, 56, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59,
	3, 2, 2, 2, 59, 60, 7, 3, 2, 2, 60, 62, 7, 7, 2, 2, 61, 63, 5, 12, 7, 2,
	62, 61, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 65, 5,
	30, 16, 2, 65, 66, 7, 7, 2, 2, 66, 67, 5, 6, 4, 2, 67, 69, 5, 10, 6, 2,
	68, 70, 7, 7, 2, 2, 69, 68, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 3, 3, 2,
	2, 2, 71, 72, 9, 2, 2, 2, 72, 5, 3, 2, 2, 2, 73, 75, 5, 4, 3, 2, 74, 73,
	3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2,
	77, 7, 3, 2, 2, 2, 78, 79, 7, 7, 2, 2, 79, 80, 5, 6, 4, 2, 80, 9, 3, 2,
	2, 2, 81, 83, 5, 8, 5, 2, 82, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 82,
	3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 11, 3, 2, 2, 2, 86, 87, 7, 10, 2, 2,
	87, 88, 5, 14, 8, 2, 88, 89, 7, 11, 2, 2, 89, 90, 7, 7, 2, 2, 90, 13, 3,
	2, 2, 2, 91, 92, 5, 20, 11, 2, 92, 93, 5, 16, 9, 2, 93, 15, 3, 2, 2, 2,
	94, 96, 5, 18, 10, 2, 95, 94, 3, 2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 95, 3,
	2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 17, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100,
	101, 7, 16, 2, 2, 101, 102, 5, 20, 11, 2, 102, 19, 3, 2, 2, 2, 103, 104,
	7, 4, 2, 2, 104, 105, 7, 12, 2, 2, 105, 106, 5, 36, 19, 2, 106, 21, 3,
	2, 2, 2, 107, 108, 7, 14, 2, 2, 108, 109, 7, 6, 2, 2, 109, 23, 3, 2, 2,
	2, 110, 111, 7, 6, 2, 2, 111, 112, 5, 26, 14, 2, 112, 25, 3, 2, 2, 2, 113,
	115, 5, 22, 12, 2, 114, 113, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114,
	3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 27, 3, 2, 2, 2, 118, 116, 3, 2,
	2, 2, 119, 120, 5, 24, 13, 2, 120, 121, 7, 15, 2, 2, 121, 122, 5, 44, 23,
	2, 122, 29, 3, 2, 2, 2, 123, 126, 5, 54, 28, 2, 124, 126, 5, 28, 15, 2,
	125, 123, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 31, 3, 2, 2, 2, 127, 128,
	7, 13, 2, 2, 128, 129, 7, 7, 2, 2, 129, 33, 3, 2, 2, 2, 130, 131, 9, 3,
	2, 2, 131, 35, 3, 2, 2, 2, 132, 134, 5, 34, 18, 2, 133, 132, 3, 2, 2, 2,
	134, 135, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	37, 3, 2, 2, 2, 137, 141, 7, 21, 2, 2, 138, 141, 5, 34, 18, 2, 139, 141,
	7, 12, 2, 2, 140, 137, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 139, 3, 2,
	2, 2, 141, 39, 3, 2, 2, 2, 142, 144, 5, 38, 20, 2, 143, 142, 3, 2, 2, 2,
	144, 145, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146,
	41, 3, 2, 2, 2, 147, 150, 5, 38, 20, 2, 148, 150, 7, 11, 2, 2, 149, 147,
	3, 2, 2, 2, 149, 148, 3, 2, 2, 2, 150, 43, 3, 2, 2, 2, 151, 153, 5, 42,
	22, 2, 152, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2,
	154, 155, 3, 2, 2, 2, 155, 45, 3, 2, 2, 2, 156, 157, 7, 18, 2, 2, 157,
	158, 5, 40, 21, 2, 158, 159, 7, 18, 2, 2, 159, 47, 3, 2, 2, 2, 160, 161,
	7, 19, 2, 2, 161, 162, 5, 40, 21, 2, 162, 163, 7, 19, 2, 2, 163, 49, 3,
	2, 2, 2, 164, 167, 5, 46, 24, 2, 165, 167, 5, 48, 25, 2, 166, 164, 3, 2,
	2, 2, 166, 165, 3, 2, 2, 2, 167, 51, 3, 2, 2, 2, 168, 169, 7, 10, 2, 2,
	169, 170, 5, 50, 26, 2, 170, 171, 7, 11, 2, 2, 171, 53, 3, 2, 2, 2, 172,
	173, 7, 5, 2, 2, 173, 174, 7, 15, 2, 2, 174, 175, 5, 52, 27, 2, 175, 176,
	5, 44, 23, 2, 176, 55, 3, 2, 2, 2, 16, 57, 62, 69, 76, 84, 97, 116, 125,
	135, 140, 145, 149, 154, 166,
}
var literalNames = []string{
	"", "", "", "'cdrom:'", "", "", "", "", "'['", "']'", "'='", "'#'", "'+'",
	"':'", "','", "'-'", "'''", "'\"'",
}
var symbolicNames = []string{
	"", "TypeR", "OptionName", "CdromSchema", "Word", "WSS", "WS", "PunctuationAllowedInOptionValue",
	"OptionsStart", "OptionsEnd", "OptionNameValueSeparator", "CommentMarker",
	"Plus", "Colon", "OptionsSeparator", "Dash", "SingleTick", "DoubleTick",
	"WordChar", "NonWhitespaceNonOptionValueNonSquareRightBracketNonEq",
}

var ruleNames = []string{
	"record", "wordWithDashSegment", "wordWithDash", "component", "componentsR",
	"optionsR", "optionsList", "additionalOptions", "additionalOption", "optionR",
	"wordWithPlus", "uriSchema", "restSchemaWords", "genericURI", "uriR", "commenterR",
	"optionValueSegment", "optionValue", "nonSquareBracketStringSegment", "nonSquareBracketString",
	"nonSpaceStringSegment", "nonSpaceString", "singleTickEnclosedString",
	"doubleTickEnclosedString", "tickEnclosedString", "enclosedString", "cdromURI",
}

type aptParser struct {
	*antlr.BaseParser
}

// NewaptParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *aptParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewaptParser(input antlr.TokenStream) *aptParser {
	this := new(aptParser)
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
	this.GrammarFileName = "apt.g4"

	return this
}

// aptParser tokens.
const (
	aptParserEOF                                                   = antlr.TokenEOF
	aptParserTypeR                                                 = 1
	aptParserOptionName                                            = 2
	aptParserCdromSchema                                           = 3
	aptParserWord                                                  = 4
	aptParserWSS                                                   = 5
	aptParserWS                                                    = 6
	aptParserPunctuationAllowedInOptionValue                       = 7
	aptParserOptionsStart                                          = 8
	aptParserOptionsEnd                                            = 9
	aptParserOptionNameValueSeparator                              = 10
	aptParserCommentMarker                                         = 11
	aptParserPlus                                                  = 12
	aptParserColon                                                 = 13
	aptParserOptionsSeparator                                      = 14
	aptParserDash                                                  = 15
	aptParserSingleTick                                            = 16
	aptParserDoubleTick                                            = 17
	aptParserWordChar                                              = 18
	aptParserNonWhitespaceNonOptionValueNonSquareRightBracketNonEq = 19
)

// aptParser rules.
const (
	aptParserRULE_record                        = 0
	aptParserRULE_wordWithDashSegment           = 1
	aptParserRULE_wordWithDash                  = 2
	aptParserRULE_component                     = 3
	aptParserRULE_componentsR                   = 4
	aptParserRULE_optionsR                      = 5
	aptParserRULE_optionsList                   = 6
	aptParserRULE_additionalOptions             = 7
	aptParserRULE_additionalOption              = 8
	aptParserRULE_optionR                       = 9
	aptParserRULE_wordWithPlus                  = 10
	aptParserRULE_uriSchema                     = 11
	aptParserRULE_restSchemaWords               = 12
	aptParserRULE_genericURI                    = 13
	aptParserRULE_uriR                          = 14
	aptParserRULE_commenterR                    = 15
	aptParserRULE_optionValueSegment            = 16
	aptParserRULE_optionValue                   = 17
	aptParserRULE_nonSquareBracketStringSegment = 18
	aptParserRULE_nonSquareBracketString        = 19
	aptParserRULE_nonSpaceStringSegment         = 20
	aptParserRULE_nonSpaceString                = 21
	aptParserRULE_singleTickEnclosedString      = 22
	aptParserRULE_doubleTickEnclosedString      = 23
	aptParserRULE_tickEnclosedString            = 24
	aptParserRULE_enclosedString                = 25
	aptParserRULE_cdromURI                      = 26
)

// IRecordContext is an interface to support dynamic dispatch.
type IRecordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRType returns the rType token.
	GetRType() antlr.Token

	// SetRType sets the rType token.
	SetRType(antlr.Token)

	// GetCommented returns the commented rule contexts.
	GetCommented() ICommenterRContext

	// GetOptions returns the options rule contexts.
	GetOptions() IOptionsRContext

	// GetUri returns the uri rule contexts.
	GetUri() IUriRContext

	// GetDistribution returns the distribution rule contexts.
	GetDistribution() IWordWithDashContext

	// GetComponents returns the components rule contexts.
	GetComponents() IComponentsRContext

	// SetCommented sets the commented rule contexts.
	SetCommented(ICommenterRContext)

	// SetOptions sets the options rule contexts.
	SetOptions(IOptionsRContext)

	// SetUri sets the uri rule contexts.
	SetUri(IUriRContext)

	// SetDistribution sets the distribution rule contexts.
	SetDistribution(IWordWithDashContext)

	// SetComponents sets the components rule contexts.
	SetComponents(IComponentsRContext)

	// IsRecordContext differentiates from other interfaces.
	IsRecordContext()
}

type RecordContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	commented    ICommenterRContext
	rType        antlr.Token
	options      IOptionsRContext
	uri          IUriRContext
	distribution IWordWithDashContext
	components   IComponentsRContext
}

func NewEmptyRecordContext() *RecordContext {
	var p = new(RecordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_record
	return p
}

func (*RecordContext) IsRecordContext() {}

func NewRecordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordContext {
	var p = new(RecordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_record

	return p
}

func (s *RecordContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordContext) GetRType() antlr.Token { return s.rType }

func (s *RecordContext) SetRType(v antlr.Token) { s.rType = v }

func (s *RecordContext) GetCommented() ICommenterRContext { return s.commented }

func (s *RecordContext) GetOptions() IOptionsRContext { return s.options }

func (s *RecordContext) GetUri() IUriRContext { return s.uri }

func (s *RecordContext) GetDistribution() IWordWithDashContext { return s.distribution }

func (s *RecordContext) GetComponents() IComponentsRContext { return s.components }

func (s *RecordContext) SetCommented(v ICommenterRContext) { s.commented = v }

func (s *RecordContext) SetOptions(v IOptionsRContext) { s.options = v }

func (s *RecordContext) SetUri(v IUriRContext) { s.uri = v }

func (s *RecordContext) SetDistribution(v IWordWithDashContext) { s.distribution = v }

func (s *RecordContext) SetComponents(v IComponentsRContext) { s.components = v }

func (s *RecordContext) AllWSS() []antlr.TerminalNode {
	return s.GetTokens(aptParserWSS)
}

func (s *RecordContext) WSS(i int) antlr.TerminalNode {
	return s.GetToken(aptParserWSS, i)
}

func (s *RecordContext) TypeR() antlr.TerminalNode {
	return s.GetToken(aptParserTypeR, 0)
}

func (s *RecordContext) UriR() IUriRContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUriRContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUriRContext)
}

func (s *RecordContext) WordWithDash() IWordWithDashContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordWithDashContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWordWithDashContext)
}

func (s *RecordContext) ComponentsR() IComponentsRContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentsRContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComponentsRContext)
}

func (s *RecordContext) CommenterR() ICommenterRContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommenterRContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommenterRContext)
}

func (s *RecordContext) OptionsR() IOptionsRContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionsRContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionsRContext)
}

func (s *RecordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterRecord(s)
	}
}

func (s *RecordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitRecord(s)
	}
}

func (p *aptParser) Record() (localctx IRecordContext) {
	this := p
	_ = this

	localctx = NewRecordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, aptParserRULE_record)
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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == aptParserCommentMarker {
		{
			p.SetState(54)

			var _x = p.CommenterR()

			localctx.(*RecordContext).commented = _x
		}

	}
	{
		p.SetState(57)

		var _m = p.Match(aptParserTypeR)

		localctx.(*RecordContext).rType = _m
	}
	{
		p.SetState(58)
		p.Match(aptParserWSS)
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == aptParserOptionsStart {
		{
			p.SetState(59)

			var _x = p.OptionsR()

			localctx.(*RecordContext).options = _x
		}

	}
	{
		p.SetState(62)

		var _x = p.UriR()

		localctx.(*RecordContext).uri = _x
	}
	{
		p.SetState(63)
		p.Match(aptParserWSS)
	}
	{
		p.SetState(64)

		var _x = p.WordWithDash()

		localctx.(*RecordContext).distribution = _x
	}
	{
		p.SetState(65)

		var _x = p.ComponentsR()

		localctx.(*RecordContext).components = _x
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == aptParserWSS {
		{
			p.SetState(66)
			p.Match(aptParserWSS)
		}

	}

	return localctx
}

// IWordWithDashSegmentContext is an interface to support dynamic dispatch.
type IWordWithDashSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordWithDashSegmentContext differentiates from other interfaces.
	IsWordWithDashSegmentContext()
}

type WordWithDashSegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordWithDashSegmentContext() *WordWithDashSegmentContext {
	var p = new(WordWithDashSegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_wordWithDashSegment
	return p
}

func (*WordWithDashSegmentContext) IsWordWithDashSegmentContext() {}

func NewWordWithDashSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordWithDashSegmentContext {
	var p = new(WordWithDashSegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_wordWithDashSegment

	return p
}

func (s *WordWithDashSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *WordWithDashSegmentContext) Word() antlr.TerminalNode {
	return s.GetToken(aptParserWord, 0)
}

func (s *WordWithDashSegmentContext) Dash() antlr.TerminalNode {
	return s.GetToken(aptParserDash, 0)
}

func (s *WordWithDashSegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordWithDashSegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordWithDashSegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterWordWithDashSegment(s)
	}
}

func (s *WordWithDashSegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitWordWithDashSegment(s)
	}
}

func (p *aptParser) WordWithDashSegment() (localctx IWordWithDashSegmentContext) {
	this := p
	_ = this

	localctx = NewWordWithDashSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, aptParserRULE_wordWithDashSegment)
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
		p.SetState(69)
		_la = p.GetTokenStream().LA(1)

		if !(_la == aptParserWord || _la == aptParserDash) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IWordWithDashContext is an interface to support dynamic dispatch.
type IWordWithDashContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordWithDashContext differentiates from other interfaces.
	IsWordWithDashContext()
}

type WordWithDashContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordWithDashContext() *WordWithDashContext {
	var p = new(WordWithDashContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_wordWithDash
	return p
}

func (*WordWithDashContext) IsWordWithDashContext() {}

func NewWordWithDashContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordWithDashContext {
	var p = new(WordWithDashContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_wordWithDash

	return p
}

func (s *WordWithDashContext) GetParser() antlr.Parser { return s.parser }

func (s *WordWithDashContext) AllWordWithDashSegment() []IWordWithDashSegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWordWithDashSegmentContext)(nil)).Elem())
	var tst = make([]IWordWithDashSegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWordWithDashSegmentContext)
		}
	}

	return tst
}

func (s *WordWithDashContext) WordWithDashSegment(i int) IWordWithDashSegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordWithDashSegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWordWithDashSegmentContext)
}

func (s *WordWithDashContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordWithDashContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordWithDashContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterWordWithDash(s)
	}
}

func (s *WordWithDashContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitWordWithDash(s)
	}
}

func (p *aptParser) WordWithDash() (localctx IWordWithDashContext) {
	this := p
	_ = this

	localctx = NewWordWithDashContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, aptParserRULE_wordWithDash)
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

	for ok := true; ok; ok = _la == aptParserWord || _la == aptParserDash {
		{
			p.SetState(71)
			p.WordWithDashSegment()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IComponentContext is an interface to support dynamic dispatch.
type IComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCId returns the cId rule contexts.
	GetCId() IWordWithDashContext

	// SetCId sets the cId rule contexts.
	SetCId(IWordWithDashContext)

	// IsComponentContext differentiates from other interfaces.
	IsComponentContext()
}

type ComponentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cId    IWordWithDashContext
}

func NewEmptyComponentContext() *ComponentContext {
	var p = new(ComponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_component
	return p
}

func (*ComponentContext) IsComponentContext() {}

func NewComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentContext {
	var p = new(ComponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_component

	return p
}

func (s *ComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentContext) GetCId() IWordWithDashContext { return s.cId }

func (s *ComponentContext) SetCId(v IWordWithDashContext) { s.cId = v }

func (s *ComponentContext) WSS() antlr.TerminalNode {
	return s.GetToken(aptParserWSS, 0)
}

func (s *ComponentContext) WordWithDash() IWordWithDashContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordWithDashContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWordWithDashContext)
}

func (s *ComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterComponent(s)
	}
}

func (s *ComponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitComponent(s)
	}
}

func (p *aptParser) Component() (localctx IComponentContext) {
	this := p
	_ = this

	localctx = NewComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, aptParserRULE_component)

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
		p.SetState(76)
		p.Match(aptParserWSS)
	}
	{
		p.SetState(77)

		var _x = p.WordWithDash()

		localctx.(*ComponentContext).cId = _x
	}

	return localctx
}

// IComponentsRContext is an interface to support dynamic dispatch.
type IComponentsRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentsRContext differentiates from other interfaces.
	IsComponentsRContext()
}

type ComponentsRContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentsRContext() *ComponentsRContext {
	var p = new(ComponentsRContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_componentsR
	return p
}

func (*ComponentsRContext) IsComponentsRContext() {}

func NewComponentsRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentsRContext {
	var p = new(ComponentsRContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_componentsR

	return p
}

func (s *ComponentsRContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentsRContext) AllComponent() []IComponentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentContext)(nil)).Elem())
	var tst = make([]IComponentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentContext)
		}
	}

	return tst
}

func (s *ComponentsRContext) Component(i int) IComponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentContext)
}

func (s *ComponentsRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentsRContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentsRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterComponentsR(s)
	}
}

func (s *ComponentsRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitComponentsR(s)
	}
}

func (p *aptParser) ComponentsR() (localctx IComponentsRContext) {
	this := p
	_ = this

	localctx = NewComponentsRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, aptParserRULE_componentsR)

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
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(79)
				p.Component()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IOptionsRContext is an interface to support dynamic dispatch.
type IOptionsRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpeningBrace returns the openingBrace token.
	GetOpeningBrace() antlr.Token

	// GetClosingBrace returns the closingBrace token.
	GetClosingBrace() antlr.Token

	// SetOpeningBrace sets the openingBrace token.
	SetOpeningBrace(antlr.Token)

	// SetClosingBrace sets the closingBrace token.
	SetClosingBrace(antlr.Token)

	// GetPairs returns the pairs rule contexts.
	GetPairs() IOptionsListContext

	// SetPairs sets the pairs rule contexts.
	SetPairs(IOptionsListContext)

	// IsOptionsRContext differentiates from other interfaces.
	IsOptionsRContext()
}

type OptionsRContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	openingBrace antlr.Token
	pairs        IOptionsListContext
	closingBrace antlr.Token
}

func NewEmptyOptionsRContext() *OptionsRContext {
	var p = new(OptionsRContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_optionsR
	return p
}

func (*OptionsRContext) IsOptionsRContext() {}

func NewOptionsRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionsRContext {
	var p = new(OptionsRContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_optionsR

	return p
}

func (s *OptionsRContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionsRContext) GetOpeningBrace() antlr.Token { return s.openingBrace }

func (s *OptionsRContext) GetClosingBrace() antlr.Token { return s.closingBrace }

func (s *OptionsRContext) SetOpeningBrace(v antlr.Token) { s.openingBrace = v }

func (s *OptionsRContext) SetClosingBrace(v antlr.Token) { s.closingBrace = v }

func (s *OptionsRContext) GetPairs() IOptionsListContext { return s.pairs }

func (s *OptionsRContext) SetPairs(v IOptionsListContext) { s.pairs = v }

func (s *OptionsRContext) WSS() antlr.TerminalNode {
	return s.GetToken(aptParserWSS, 0)
}

func (s *OptionsRContext) OptionsStart() antlr.TerminalNode {
	return s.GetToken(aptParserOptionsStart, 0)
}

func (s *OptionsRContext) OptionsList() IOptionsListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionsListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionsListContext)
}

func (s *OptionsRContext) OptionsEnd() antlr.TerminalNode {
	return s.GetToken(aptParserOptionsEnd, 0)
}

func (s *OptionsRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionsRContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionsRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterOptionsR(s)
	}
}

func (s *OptionsRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitOptionsR(s)
	}
}

func (p *aptParser) OptionsR() (localctx IOptionsRContext) {
	this := p
	_ = this

	localctx = NewOptionsRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, aptParserRULE_optionsR)

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
		p.SetState(84)

		var _m = p.Match(aptParserOptionsStart)

		localctx.(*OptionsRContext).openingBrace = _m
	}
	{
		p.SetState(85)

		var _x = p.OptionsList()

		localctx.(*OptionsRContext).pairs = _x
	}
	{
		p.SetState(86)

		var _m = p.Match(aptParserOptionsEnd)

		localctx.(*OptionsRContext).closingBrace = _m
	}
	{
		p.SetState(87)
		p.Match(aptParserWSS)
	}

	return localctx
}

// IOptionsListContext is an interface to support dynamic dispatch.
type IOptionsListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirstOption returns the firstOption rule contexts.
	GetFirstOption() IOptionRContext

	// GetRestOptions returns the restOptions rule contexts.
	GetRestOptions() IAdditionalOptionsContext

	// SetFirstOption sets the firstOption rule contexts.
	SetFirstOption(IOptionRContext)

	// SetRestOptions sets the restOptions rule contexts.
	SetRestOptions(IAdditionalOptionsContext)

	// IsOptionsListContext differentiates from other interfaces.
	IsOptionsListContext()
}

type OptionsListContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	firstOption IOptionRContext
	restOptions IAdditionalOptionsContext
}

func NewEmptyOptionsListContext() *OptionsListContext {
	var p = new(OptionsListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_optionsList
	return p
}

func (*OptionsListContext) IsOptionsListContext() {}

func NewOptionsListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionsListContext {
	var p = new(OptionsListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_optionsList

	return p
}

func (s *OptionsListContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionsListContext) GetFirstOption() IOptionRContext { return s.firstOption }

func (s *OptionsListContext) GetRestOptions() IAdditionalOptionsContext { return s.restOptions }

func (s *OptionsListContext) SetFirstOption(v IOptionRContext) { s.firstOption = v }

func (s *OptionsListContext) SetRestOptions(v IAdditionalOptionsContext) { s.restOptions = v }

func (s *OptionsListContext) OptionR() IOptionRContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionRContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionRContext)
}

func (s *OptionsListContext) AdditionalOptions() IAdditionalOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditionalOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditionalOptionsContext)
}

func (s *OptionsListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionsListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionsListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterOptionsList(s)
	}
}

func (s *OptionsListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitOptionsList(s)
	}
}

func (p *aptParser) OptionsList() (localctx IOptionsListContext) {
	this := p
	_ = this

	localctx = NewOptionsListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, aptParserRULE_optionsList)

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
		p.SetState(89)

		var _x = p.OptionR()

		localctx.(*OptionsListContext).firstOption = _x
	}
	{
		p.SetState(90)

		var _x = p.AdditionalOptions()

		localctx.(*OptionsListContext).restOptions = _x
	}

	return localctx
}

// IAdditionalOptionsContext is an interface to support dynamic dispatch.
type IAdditionalOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditionalOptionsContext differentiates from other interfaces.
	IsAdditionalOptionsContext()
}

type AdditionalOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditionalOptionsContext() *AdditionalOptionsContext {
	var p = new(AdditionalOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_additionalOptions
	return p
}

func (*AdditionalOptionsContext) IsAdditionalOptionsContext() {}

func NewAdditionalOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditionalOptionsContext {
	var p = new(AdditionalOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_additionalOptions

	return p
}

func (s *AdditionalOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditionalOptionsContext) AllAdditionalOption() []IAdditionalOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdditionalOptionContext)(nil)).Elem())
	var tst = make([]IAdditionalOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdditionalOptionContext)
		}
	}

	return tst
}

func (s *AdditionalOptionsContext) AdditionalOption(i int) IAdditionalOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditionalOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdditionalOptionContext)
}

func (s *AdditionalOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionalOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditionalOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterAdditionalOptions(s)
	}
}

func (s *AdditionalOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitAdditionalOptions(s)
	}
}

func (p *aptParser) AdditionalOptions() (localctx IAdditionalOptionsContext) {
	this := p
	_ = this

	localctx = NewAdditionalOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, aptParserRULE_additionalOptions)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == aptParserOptionsSeparator {
		{
			p.SetState(92)
			p.AdditionalOption()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAdditionalOptionContext is an interface to support dynamic dispatch.
type IAdditionalOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSeparator returns the separator token.
	GetSeparator() antlr.Token

	// SetSeparator sets the separator token.
	SetSeparator(antlr.Token)

	// GetOption returns the option rule contexts.
	GetOption() IOptionRContext

	// SetOption sets the option rule contexts.
	SetOption(IOptionRContext)

	// IsAdditionalOptionContext differentiates from other interfaces.
	IsAdditionalOptionContext()
}

type AdditionalOptionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	separator antlr.Token
	option    IOptionRContext
}

func NewEmptyAdditionalOptionContext() *AdditionalOptionContext {
	var p = new(AdditionalOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_additionalOption
	return p
}

func (*AdditionalOptionContext) IsAdditionalOptionContext() {}

func NewAdditionalOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditionalOptionContext {
	var p = new(AdditionalOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_additionalOption

	return p
}

func (s *AdditionalOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditionalOptionContext) GetSeparator() antlr.Token { return s.separator }

func (s *AdditionalOptionContext) SetSeparator(v antlr.Token) { s.separator = v }

func (s *AdditionalOptionContext) GetOption() IOptionRContext { return s.option }

func (s *AdditionalOptionContext) SetOption(v IOptionRContext) { s.option = v }

func (s *AdditionalOptionContext) OptionsSeparator() antlr.TerminalNode {
	return s.GetToken(aptParserOptionsSeparator, 0)
}

func (s *AdditionalOptionContext) OptionR() IOptionRContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionRContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionRContext)
}

func (s *AdditionalOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionalOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditionalOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterAdditionalOption(s)
	}
}

func (s *AdditionalOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitAdditionalOption(s)
	}
}

func (p *aptParser) AdditionalOption() (localctx IAdditionalOptionContext) {
	this := p
	_ = this

	localctx = NewAdditionalOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, aptParserRULE_additionalOption)

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

		var _m = p.Match(aptParserOptionsSeparator)

		localctx.(*AdditionalOptionContext).separator = _m
	}
	{
		p.SetState(99)

		var _x = p.OptionR()

		localctx.(*AdditionalOptionContext).option = _x
	}

	return localctx
}

// IOptionRContext is an interface to support dynamic dispatch.
type IOptionRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IOptionValueContext

	// SetValue sets the value rule contexts.
	SetValue(IOptionValueContext)

	// IsOptionRContext differentiates from other interfaces.
	IsOptionRContext()
}

type OptionRContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  IOptionValueContext
}

func NewEmptyOptionRContext() *OptionRContext {
	var p = new(OptionRContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_optionR
	return p
}

func (*OptionRContext) IsOptionRContext() {}

func NewOptionRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionRContext {
	var p = new(OptionRContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_optionR

	return p
}

func (s *OptionRContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionRContext) GetKey() antlr.Token { return s.key }

func (s *OptionRContext) SetKey(v antlr.Token) { s.key = v }

func (s *OptionRContext) GetValue() IOptionValueContext { return s.value }

func (s *OptionRContext) SetValue(v IOptionValueContext) { s.value = v }

func (s *OptionRContext) OptionNameValueSeparator() antlr.TerminalNode {
	return s.GetToken(aptParserOptionNameValueSeparator, 0)
}

func (s *OptionRContext) OptionName() antlr.TerminalNode {
	return s.GetToken(aptParserOptionName, 0)
}

func (s *OptionRContext) OptionValue() IOptionValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionValueContext)
}

func (s *OptionRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionRContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterOptionR(s)
	}
}

func (s *OptionRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitOptionR(s)
	}
}

func (p *aptParser) OptionR() (localctx IOptionRContext) {
	this := p
	_ = this

	localctx = NewOptionRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, aptParserRULE_optionR)

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
		p.SetState(101)

		var _m = p.Match(aptParserOptionName)

		localctx.(*OptionRContext).key = _m
	}
	{
		p.SetState(102)
		p.Match(aptParserOptionNameValueSeparator)
	}
	{
		p.SetState(103)

		var _x = p.OptionValue()

		localctx.(*OptionRContext).value = _x
	}

	return localctx
}

// IWordWithPlusContext is an interface to support dynamic dispatch.
type IWordWithPlusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWord returns the word token.
	GetWord() antlr.Token

	// SetWord sets the word token.
	SetWord(antlr.Token)

	// IsWordWithPlusContext differentiates from other interfaces.
	IsWordWithPlusContext()
}

type WordWithPlusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	word   antlr.Token
}

func NewEmptyWordWithPlusContext() *WordWithPlusContext {
	var p = new(WordWithPlusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_wordWithPlus
	return p
}

func (*WordWithPlusContext) IsWordWithPlusContext() {}

func NewWordWithPlusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordWithPlusContext {
	var p = new(WordWithPlusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_wordWithPlus

	return p
}

func (s *WordWithPlusContext) GetParser() antlr.Parser { return s.parser }

func (s *WordWithPlusContext) GetWord() antlr.Token { return s.word }

func (s *WordWithPlusContext) SetWord(v antlr.Token) { s.word = v }

func (s *WordWithPlusContext) Plus() antlr.TerminalNode {
	return s.GetToken(aptParserPlus, 0)
}

func (s *WordWithPlusContext) Word() antlr.TerminalNode {
	return s.GetToken(aptParserWord, 0)
}

func (s *WordWithPlusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordWithPlusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordWithPlusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterWordWithPlus(s)
	}
}

func (s *WordWithPlusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitWordWithPlus(s)
	}
}

func (p *aptParser) WordWithPlus() (localctx IWordWithPlusContext) {
	this := p
	_ = this

	localctx = NewWordWithPlusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, aptParserRULE_wordWithPlus)

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
		p.SetState(105)
		p.Match(aptParserPlus)
	}
	{
		p.SetState(106)

		var _m = p.Match(aptParserWord)

		localctx.(*WordWithPlusContext).word = _m
	}

	return localctx
}

// IUriSchemaContext is an interface to support dynamic dispatch.
type IUriSchemaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWord returns the word token.
	GetWord() antlr.Token

	// SetWord sets the word token.
	SetWord(antlr.Token)

	// GetRestWords returns the restWords rule contexts.
	GetRestWords() IRestSchemaWordsContext

	// SetRestWords sets the restWords rule contexts.
	SetRestWords(IRestSchemaWordsContext)

	// IsUriSchemaContext differentiates from other interfaces.
	IsUriSchemaContext()
}

type UriSchemaContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	word      antlr.Token
	restWords IRestSchemaWordsContext
}

func NewEmptyUriSchemaContext() *UriSchemaContext {
	var p = new(UriSchemaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_uriSchema
	return p
}

func (*UriSchemaContext) IsUriSchemaContext() {}

func NewUriSchemaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UriSchemaContext {
	var p = new(UriSchemaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_uriSchema

	return p
}

func (s *UriSchemaContext) GetParser() antlr.Parser { return s.parser }

func (s *UriSchemaContext) GetWord() antlr.Token { return s.word }

func (s *UriSchemaContext) SetWord(v antlr.Token) { s.word = v }

func (s *UriSchemaContext) GetRestWords() IRestSchemaWordsContext { return s.restWords }

func (s *UriSchemaContext) SetRestWords(v IRestSchemaWordsContext) { s.restWords = v }

func (s *UriSchemaContext) Word() antlr.TerminalNode {
	return s.GetToken(aptParserWord, 0)
}

func (s *UriSchemaContext) RestSchemaWords() IRestSchemaWordsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRestSchemaWordsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRestSchemaWordsContext)
}

func (s *UriSchemaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UriSchemaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UriSchemaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterUriSchema(s)
	}
}

func (s *UriSchemaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitUriSchema(s)
	}
}

func (p *aptParser) UriSchema() (localctx IUriSchemaContext) {
	this := p
	_ = this

	localctx = NewUriSchemaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, aptParserRULE_uriSchema)

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
		p.SetState(108)

		var _m = p.Match(aptParserWord)

		localctx.(*UriSchemaContext).word = _m
	}
	{
		p.SetState(109)

		var _x = p.RestSchemaWords()

		localctx.(*UriSchemaContext).restWords = _x
	}

	return localctx
}

// IRestSchemaWordsContext is an interface to support dynamic dispatch.
type IRestSchemaWordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRestSchemaWordsContext differentiates from other interfaces.
	IsRestSchemaWordsContext()
}

type RestSchemaWordsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRestSchemaWordsContext() *RestSchemaWordsContext {
	var p = new(RestSchemaWordsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_restSchemaWords
	return p
}

func (*RestSchemaWordsContext) IsRestSchemaWordsContext() {}

func NewRestSchemaWordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RestSchemaWordsContext {
	var p = new(RestSchemaWordsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_restSchemaWords

	return p
}

func (s *RestSchemaWordsContext) GetParser() antlr.Parser { return s.parser }

func (s *RestSchemaWordsContext) AllWordWithPlus() []IWordWithPlusContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWordWithPlusContext)(nil)).Elem())
	var tst = make([]IWordWithPlusContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWordWithPlusContext)
		}
	}

	return tst
}

func (s *RestSchemaWordsContext) WordWithPlus(i int) IWordWithPlusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWordWithPlusContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWordWithPlusContext)
}

func (s *RestSchemaWordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RestSchemaWordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RestSchemaWordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterRestSchemaWords(s)
	}
}

func (s *RestSchemaWordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitRestSchemaWords(s)
	}
}

func (p *aptParser) RestSchemaWords() (localctx IRestSchemaWordsContext) {
	this := p
	_ = this

	localctx = NewRestSchemaWordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, aptParserRULE_restSchemaWords)
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
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == aptParserPlus {
		{
			p.SetState(111)
			p.WordWithPlus()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGenericURIContext is an interface to support dynamic dispatch.
type IGenericURIContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSchema returns the schema rule contexts.
	GetSchema() IUriSchemaContext

	// GetRestOfURI returns the restOfURI rule contexts.
	GetRestOfURI() INonSpaceStringContext

	// SetSchema sets the schema rule contexts.
	SetSchema(IUriSchemaContext)

	// SetRestOfURI sets the restOfURI rule contexts.
	SetRestOfURI(INonSpaceStringContext)

	// IsGenericURIContext differentiates from other interfaces.
	IsGenericURIContext()
}

type GenericURIContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	schema    IUriSchemaContext
	restOfURI INonSpaceStringContext
}

func NewEmptyGenericURIContext() *GenericURIContext {
	var p = new(GenericURIContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_genericURI
	return p
}

func (*GenericURIContext) IsGenericURIContext() {}

func NewGenericURIContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericURIContext {
	var p = new(GenericURIContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_genericURI

	return p
}

func (s *GenericURIContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericURIContext) GetSchema() IUriSchemaContext { return s.schema }

func (s *GenericURIContext) GetRestOfURI() INonSpaceStringContext { return s.restOfURI }

func (s *GenericURIContext) SetSchema(v IUriSchemaContext) { s.schema = v }

func (s *GenericURIContext) SetRestOfURI(v INonSpaceStringContext) { s.restOfURI = v }

func (s *GenericURIContext) Colon() antlr.TerminalNode {
	return s.GetToken(aptParserColon, 0)
}

func (s *GenericURIContext) UriSchema() IUriSchemaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUriSchemaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUriSchemaContext)
}

func (s *GenericURIContext) NonSpaceString() INonSpaceStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonSpaceStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonSpaceStringContext)
}

func (s *GenericURIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericURIContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericURIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterGenericURI(s)
	}
}

func (s *GenericURIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitGenericURI(s)
	}
}

func (p *aptParser) GenericURI() (localctx IGenericURIContext) {
	this := p
	_ = this

	localctx = NewGenericURIContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, aptParserRULE_genericURI)

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
		p.SetState(117)

		var _x = p.UriSchema()

		localctx.(*GenericURIContext).schema = _x
	}
	{
		p.SetState(118)
		p.Match(aptParserColon)
	}
	{
		p.SetState(119)

		var _x = p.NonSpaceString()

		localctx.(*GenericURIContext).restOfURI = _x
	}

	return localctx
}

// IUriRContext is an interface to support dynamic dispatch.
type IUriRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUriRContext differentiates from other interfaces.
	IsUriRContext()
}

type UriRContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUriRContext() *UriRContext {
	var p = new(UriRContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_uriR
	return p
}

func (*UriRContext) IsUriRContext() {}

func NewUriRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UriRContext {
	var p = new(UriRContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_uriR

	return p
}

func (s *UriRContext) GetParser() antlr.Parser { return s.parser }

func (s *UriRContext) CdromURI() ICdromURIContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICdromURIContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICdromURIContext)
}

func (s *UriRContext) GenericURI() IGenericURIContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGenericURIContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGenericURIContext)
}

func (s *UriRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UriRContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UriRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterUriR(s)
	}
}

func (s *UriRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitUriR(s)
	}
}

func (p *aptParser) UriR() (localctx IUriRContext) {
	this := p
	_ = this

	localctx = NewUriRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, aptParserRULE_uriR)

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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aptParserCdromSchema:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.CdromURI()
		}

	case aptParserWord:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.GenericURI()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommenterRContext is an interface to support dynamic dispatch.
type ICommenterRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommenterRContext differentiates from other interfaces.
	IsCommenterRContext()
}

type CommenterRContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommenterRContext() *CommenterRContext {
	var p = new(CommenterRContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_commenterR
	return p
}

func (*CommenterRContext) IsCommenterRContext() {}

func NewCommenterRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommenterRContext {
	var p = new(CommenterRContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_commenterR

	return p
}

func (s *CommenterRContext) GetParser() antlr.Parser { return s.parser }

func (s *CommenterRContext) CommentMarker() antlr.TerminalNode {
	return s.GetToken(aptParserCommentMarker, 0)
}

func (s *CommenterRContext) WSS() antlr.TerminalNode {
	return s.GetToken(aptParserWSS, 0)
}

func (s *CommenterRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommenterRContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommenterRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterCommenterR(s)
	}
}

func (s *CommenterRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitCommenterR(s)
	}
}

func (p *aptParser) CommenterR() (localctx ICommenterRContext) {
	this := p
	_ = this

	localctx = NewCommenterRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, aptParserRULE_commenterR)

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
		p.SetState(125)
		p.Match(aptParserCommentMarker)
	}
	{
		p.SetState(126)
		p.Match(aptParserWSS)
	}

	return localctx
}

// IOptionValueSegmentContext is an interface to support dynamic dispatch.
type IOptionValueSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionValueSegmentContext differentiates from other interfaces.
	IsOptionValueSegmentContext()
}

type OptionValueSegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionValueSegmentContext() *OptionValueSegmentContext {
	var p = new(OptionValueSegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_optionValueSegment
	return p
}

func (*OptionValueSegmentContext) IsOptionValueSegmentContext() {}

func NewOptionValueSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionValueSegmentContext {
	var p = new(OptionValueSegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_optionValueSegment

	return p
}

func (s *OptionValueSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionValueSegmentContext) Word() antlr.TerminalNode {
	return s.GetToken(aptParserWord, 0)
}

func (s *OptionValueSegmentContext) PunctuationAllowedInOptionValue() antlr.TerminalNode {
	return s.GetToken(aptParserPunctuationAllowedInOptionValue, 0)
}

func (s *OptionValueSegmentContext) Dash() antlr.TerminalNode {
	return s.GetToken(aptParserDash, 0)
}

func (s *OptionValueSegmentContext) OptionName() antlr.TerminalNode {
	return s.GetToken(aptParserOptionName, 0)
}

func (s *OptionValueSegmentContext) CdromSchema() antlr.TerminalNode {
	return s.GetToken(aptParserCdromSchema, 0)
}

func (s *OptionValueSegmentContext) TypeR() antlr.TerminalNode {
	return s.GetToken(aptParserTypeR, 0)
}

func (s *OptionValueSegmentContext) Plus() antlr.TerminalNode {
	return s.GetToken(aptParserPlus, 0)
}

func (s *OptionValueSegmentContext) Colon() antlr.TerminalNode {
	return s.GetToken(aptParserColon, 0)
}

func (s *OptionValueSegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionValueSegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionValueSegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterOptionValueSegment(s)
	}
}

func (s *OptionValueSegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitOptionValueSegment(s)
	}
}

func (p *aptParser) OptionValueSegment() (localctx IOptionValueSegmentContext) {
	this := p
	_ = this

	localctx = NewOptionValueSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, aptParserRULE_optionValueSegment)
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
		p.SetState(128)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<aptParserTypeR)|(1<<aptParserOptionName)|(1<<aptParserCdromSchema)|(1<<aptParserWord)|(1<<aptParserPunctuationAllowedInOptionValue)|(1<<aptParserPlus)|(1<<aptParserColon)|(1<<aptParserDash))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOptionValueContext is an interface to support dynamic dispatch.
type IOptionValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionValueContext differentiates from other interfaces.
	IsOptionValueContext()
}

type OptionValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionValueContext() *OptionValueContext {
	var p = new(OptionValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_optionValue
	return p
}

func (*OptionValueContext) IsOptionValueContext() {}

func NewOptionValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionValueContext {
	var p = new(OptionValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_optionValue

	return p
}

func (s *OptionValueContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionValueContext) AllOptionValueSegment() []IOptionValueSegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionValueSegmentContext)(nil)).Elem())
	var tst = make([]IOptionValueSegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionValueSegmentContext)
		}
	}

	return tst
}

func (s *OptionValueContext) OptionValueSegment(i int) IOptionValueSegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionValueSegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionValueSegmentContext)
}

func (s *OptionValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterOptionValue(s)
	}
}

func (s *OptionValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitOptionValue(s)
	}
}

func (p *aptParser) OptionValue() (localctx IOptionValueContext) {
	this := p
	_ = this

	localctx = NewOptionValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, aptParserRULE_optionValue)
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
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<aptParserTypeR)|(1<<aptParserOptionName)|(1<<aptParserCdromSchema)|(1<<aptParserWord)|(1<<aptParserPunctuationAllowedInOptionValue)|(1<<aptParserPlus)|(1<<aptParserColon)|(1<<aptParserDash))) != 0) {
		{
			p.SetState(130)
			p.OptionValueSegment()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INonSquareBracketStringSegmentContext is an interface to support dynamic dispatch.
type INonSquareBracketStringSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonSquareBracketStringSegmentContext differentiates from other interfaces.
	IsNonSquareBracketStringSegmentContext()
}

type NonSquareBracketStringSegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonSquareBracketStringSegmentContext() *NonSquareBracketStringSegmentContext {
	var p = new(NonSquareBracketStringSegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_nonSquareBracketStringSegment
	return p
}

func (*NonSquareBracketStringSegmentContext) IsNonSquareBracketStringSegmentContext() {}

func NewNonSquareBracketStringSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonSquareBracketStringSegmentContext {
	var p = new(NonSquareBracketStringSegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_nonSquareBracketStringSegment

	return p
}

func (s *NonSquareBracketStringSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *NonSquareBracketStringSegmentContext) NonWhitespaceNonOptionValueNonSquareRightBracketNonEq() antlr.TerminalNode {
	return s.GetToken(aptParserNonWhitespaceNonOptionValueNonSquareRightBracketNonEq, 0)
}

func (s *NonSquareBracketStringSegmentContext) OptionValueSegment() IOptionValueSegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionValueSegmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionValueSegmentContext)
}

func (s *NonSquareBracketStringSegmentContext) OptionNameValueSeparator() antlr.TerminalNode {
	return s.GetToken(aptParserOptionNameValueSeparator, 0)
}

func (s *NonSquareBracketStringSegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonSquareBracketStringSegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonSquareBracketStringSegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterNonSquareBracketStringSegment(s)
	}
}

func (s *NonSquareBracketStringSegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitNonSquareBracketStringSegment(s)
	}
}

func (p *aptParser) NonSquareBracketStringSegment() (localctx INonSquareBracketStringSegmentContext) {
	this := p
	_ = this

	localctx = NewNonSquareBracketStringSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, aptParserRULE_nonSquareBracketStringSegment)

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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aptParserNonWhitespaceNonOptionValueNonSquareRightBracketNonEq:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(aptParserNonWhitespaceNonOptionValueNonSquareRightBracketNonEq)
		}

	case aptParserTypeR, aptParserOptionName, aptParserCdromSchema, aptParserWord, aptParserPunctuationAllowedInOptionValue, aptParserPlus, aptParserColon, aptParserDash:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.OptionValueSegment()
		}

	case aptParserOptionNameValueSeparator:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.Match(aptParserOptionNameValueSeparator)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INonSquareBracketStringContext is an interface to support dynamic dispatch.
type INonSquareBracketStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonSquareBracketStringContext differentiates from other interfaces.
	IsNonSquareBracketStringContext()
}

type NonSquareBracketStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonSquareBracketStringContext() *NonSquareBracketStringContext {
	var p = new(NonSquareBracketStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_nonSquareBracketString
	return p
}

func (*NonSquareBracketStringContext) IsNonSquareBracketStringContext() {}

func NewNonSquareBracketStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonSquareBracketStringContext {
	var p = new(NonSquareBracketStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_nonSquareBracketString

	return p
}

func (s *NonSquareBracketStringContext) GetParser() antlr.Parser { return s.parser }

func (s *NonSquareBracketStringContext) AllNonSquareBracketStringSegment() []INonSquareBracketStringSegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INonSquareBracketStringSegmentContext)(nil)).Elem())
	var tst = make([]INonSquareBracketStringSegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INonSquareBracketStringSegmentContext)
		}
	}

	return tst
}

func (s *NonSquareBracketStringContext) NonSquareBracketStringSegment(i int) INonSquareBracketStringSegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonSquareBracketStringSegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INonSquareBracketStringSegmentContext)
}

func (s *NonSquareBracketStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonSquareBracketStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonSquareBracketStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterNonSquareBracketString(s)
	}
}

func (s *NonSquareBracketStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitNonSquareBracketString(s)
	}
}

func (p *aptParser) NonSquareBracketString() (localctx INonSquareBracketStringContext) {
	this := p
	_ = this

	localctx = NewNonSquareBracketStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, aptParserRULE_nonSquareBracketString)
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
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<aptParserTypeR)|(1<<aptParserOptionName)|(1<<aptParserCdromSchema)|(1<<aptParserWord)|(1<<aptParserPunctuationAllowedInOptionValue)|(1<<aptParserOptionNameValueSeparator)|(1<<aptParserPlus)|(1<<aptParserColon)|(1<<aptParserDash)|(1<<aptParserNonWhitespaceNonOptionValueNonSquareRightBracketNonEq))) != 0) {
		{
			p.SetState(140)
			p.NonSquareBracketStringSegment()
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INonSpaceStringSegmentContext is an interface to support dynamic dispatch.
type INonSpaceStringSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonSpaceStringSegmentContext differentiates from other interfaces.
	IsNonSpaceStringSegmentContext()
}

type NonSpaceStringSegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonSpaceStringSegmentContext() *NonSpaceStringSegmentContext {
	var p = new(NonSpaceStringSegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_nonSpaceStringSegment
	return p
}

func (*NonSpaceStringSegmentContext) IsNonSpaceStringSegmentContext() {}

func NewNonSpaceStringSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonSpaceStringSegmentContext {
	var p = new(NonSpaceStringSegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_nonSpaceStringSegment

	return p
}

func (s *NonSpaceStringSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *NonSpaceStringSegmentContext) NonSquareBracketStringSegment() INonSquareBracketStringSegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonSquareBracketStringSegmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonSquareBracketStringSegmentContext)
}

func (s *NonSpaceStringSegmentContext) OptionsEnd() antlr.TerminalNode {
	return s.GetToken(aptParserOptionsEnd, 0)
}

func (s *NonSpaceStringSegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonSpaceStringSegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonSpaceStringSegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterNonSpaceStringSegment(s)
	}
}

func (s *NonSpaceStringSegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitNonSpaceStringSegment(s)
	}
}

func (p *aptParser) NonSpaceStringSegment() (localctx INonSpaceStringSegmentContext) {
	this := p
	_ = this

	localctx = NewNonSpaceStringSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, aptParserRULE_nonSpaceStringSegment)

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

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aptParserTypeR, aptParserOptionName, aptParserCdromSchema, aptParserWord, aptParserPunctuationAllowedInOptionValue, aptParserOptionNameValueSeparator, aptParserPlus, aptParserColon, aptParserDash, aptParserNonWhitespaceNonOptionValueNonSquareRightBracketNonEq:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.NonSquareBracketStringSegment()
		}

	case aptParserOptionsEnd:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Match(aptParserOptionsEnd)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INonSpaceStringContext is an interface to support dynamic dispatch.
type INonSpaceStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonSpaceStringContext differentiates from other interfaces.
	IsNonSpaceStringContext()
}

type NonSpaceStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonSpaceStringContext() *NonSpaceStringContext {
	var p = new(NonSpaceStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_nonSpaceString
	return p
}

func (*NonSpaceStringContext) IsNonSpaceStringContext() {}

func NewNonSpaceStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonSpaceStringContext {
	var p = new(NonSpaceStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_nonSpaceString

	return p
}

func (s *NonSpaceStringContext) GetParser() antlr.Parser { return s.parser }

func (s *NonSpaceStringContext) AllNonSpaceStringSegment() []INonSpaceStringSegmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INonSpaceStringSegmentContext)(nil)).Elem())
	var tst = make([]INonSpaceStringSegmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INonSpaceStringSegmentContext)
		}
	}

	return tst
}

func (s *NonSpaceStringContext) NonSpaceStringSegment(i int) INonSpaceStringSegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonSpaceStringSegmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INonSpaceStringSegmentContext)
}

func (s *NonSpaceStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonSpaceStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonSpaceStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterNonSpaceString(s)
	}
}

func (s *NonSpaceStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitNonSpaceString(s)
	}
}

func (p *aptParser) NonSpaceString() (localctx INonSpaceStringContext) {
	this := p
	_ = this

	localctx = NewNonSpaceStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, aptParserRULE_nonSpaceString)
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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<aptParserTypeR)|(1<<aptParserOptionName)|(1<<aptParserCdromSchema)|(1<<aptParserWord)|(1<<aptParserPunctuationAllowedInOptionValue)|(1<<aptParserOptionsEnd)|(1<<aptParserOptionNameValueSeparator)|(1<<aptParserPlus)|(1<<aptParserColon)|(1<<aptParserDash)|(1<<aptParserNonWhitespaceNonOptionValueNonSquareRightBracketNonEq))) != 0) {
		{
			p.SetState(149)
			p.NonSpaceStringSegment()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISingleTickEnclosedStringContext is an interface to support dynamic dispatch.
type ISingleTickEnclosedStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleTickEnclosedStringContext differentiates from other interfaces.
	IsSingleTickEnclosedStringContext()
}

type SingleTickEnclosedStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleTickEnclosedStringContext() *SingleTickEnclosedStringContext {
	var p = new(SingleTickEnclosedStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_singleTickEnclosedString
	return p
}

func (*SingleTickEnclosedStringContext) IsSingleTickEnclosedStringContext() {}

func NewSingleTickEnclosedStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleTickEnclosedStringContext {
	var p = new(SingleTickEnclosedStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_singleTickEnclosedString

	return p
}

func (s *SingleTickEnclosedStringContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleTickEnclosedStringContext) AllSingleTick() []antlr.TerminalNode {
	return s.GetTokens(aptParserSingleTick)
}

func (s *SingleTickEnclosedStringContext) SingleTick(i int) antlr.TerminalNode {
	return s.GetToken(aptParserSingleTick, i)
}

func (s *SingleTickEnclosedStringContext) NonSquareBracketString() INonSquareBracketStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonSquareBracketStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonSquareBracketStringContext)
}

func (s *SingleTickEnclosedStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTickEnclosedStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleTickEnclosedStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterSingleTickEnclosedString(s)
	}
}

func (s *SingleTickEnclosedStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitSingleTickEnclosedString(s)
	}
}

func (p *aptParser) SingleTickEnclosedString() (localctx ISingleTickEnclosedStringContext) {
	this := p
	_ = this

	localctx = NewSingleTickEnclosedStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, aptParserRULE_singleTickEnclosedString)

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
		p.SetState(154)
		p.Match(aptParserSingleTick)
	}
	{
		p.SetState(155)
		p.NonSquareBracketString()
	}
	{
		p.SetState(156)
		p.Match(aptParserSingleTick)
	}

	return localctx
}

// IDoubleTickEnclosedStringContext is an interface to support dynamic dispatch.
type IDoubleTickEnclosedStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoubleTickEnclosedStringContext differentiates from other interfaces.
	IsDoubleTickEnclosedStringContext()
}

type DoubleTickEnclosedStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoubleTickEnclosedStringContext() *DoubleTickEnclosedStringContext {
	var p = new(DoubleTickEnclosedStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_doubleTickEnclosedString
	return p
}

func (*DoubleTickEnclosedStringContext) IsDoubleTickEnclosedStringContext() {}

func NewDoubleTickEnclosedStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoubleTickEnclosedStringContext {
	var p = new(DoubleTickEnclosedStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_doubleTickEnclosedString

	return p
}

func (s *DoubleTickEnclosedStringContext) GetParser() antlr.Parser { return s.parser }

func (s *DoubleTickEnclosedStringContext) AllDoubleTick() []antlr.TerminalNode {
	return s.GetTokens(aptParserDoubleTick)
}

func (s *DoubleTickEnclosedStringContext) DoubleTick(i int) antlr.TerminalNode {
	return s.GetToken(aptParserDoubleTick, i)
}

func (s *DoubleTickEnclosedStringContext) NonSquareBracketString() INonSquareBracketStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonSquareBracketStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonSquareBracketStringContext)
}

func (s *DoubleTickEnclosedStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleTickEnclosedStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoubleTickEnclosedStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterDoubleTickEnclosedString(s)
	}
}

func (s *DoubleTickEnclosedStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitDoubleTickEnclosedString(s)
	}
}

func (p *aptParser) DoubleTickEnclosedString() (localctx IDoubleTickEnclosedStringContext) {
	this := p
	_ = this

	localctx = NewDoubleTickEnclosedStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, aptParserRULE_doubleTickEnclosedString)

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
		p.SetState(158)
		p.Match(aptParserDoubleTick)
	}
	{
		p.SetState(159)
		p.NonSquareBracketString()
	}
	{
		p.SetState(160)
		p.Match(aptParserDoubleTick)
	}

	return localctx
}

// ITickEnclosedStringContext is an interface to support dynamic dispatch.
type ITickEnclosedStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTickEnclosedStringContext differentiates from other interfaces.
	IsTickEnclosedStringContext()
}

type TickEnclosedStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTickEnclosedStringContext() *TickEnclosedStringContext {
	var p = new(TickEnclosedStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_tickEnclosedString
	return p
}

func (*TickEnclosedStringContext) IsTickEnclosedStringContext() {}

func NewTickEnclosedStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TickEnclosedStringContext {
	var p = new(TickEnclosedStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_tickEnclosedString

	return p
}

func (s *TickEnclosedStringContext) GetParser() antlr.Parser { return s.parser }

func (s *TickEnclosedStringContext) SingleTickEnclosedString() ISingleTickEnclosedStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleTickEnclosedStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleTickEnclosedStringContext)
}

func (s *TickEnclosedStringContext) DoubleTickEnclosedString() IDoubleTickEnclosedStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoubleTickEnclosedStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoubleTickEnclosedStringContext)
}

func (s *TickEnclosedStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TickEnclosedStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TickEnclosedStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterTickEnclosedString(s)
	}
}

func (s *TickEnclosedStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitTickEnclosedString(s)
	}
}

func (p *aptParser) TickEnclosedString() (localctx ITickEnclosedStringContext) {
	this := p
	_ = this

	localctx = NewTickEnclosedStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, aptParserRULE_tickEnclosedString)

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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aptParserSingleTick:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.SingleTickEnclosedString()
		}

	case aptParserDoubleTick:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.DoubleTickEnclosedString()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnclosedStringContext is an interface to support dynamic dispatch.
type IEnclosedStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnclosedStringContext differentiates from other interfaces.
	IsEnclosedStringContext()
}

type EnclosedStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnclosedStringContext() *EnclosedStringContext {
	var p = new(EnclosedStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_enclosedString
	return p
}

func (*EnclosedStringContext) IsEnclosedStringContext() {}

func NewEnclosedStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnclosedStringContext {
	var p = new(EnclosedStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_enclosedString

	return p
}

func (s *EnclosedStringContext) GetParser() antlr.Parser { return s.parser }

func (s *EnclosedStringContext) OptionsStart() antlr.TerminalNode {
	return s.GetToken(aptParserOptionsStart, 0)
}

func (s *EnclosedStringContext) TickEnclosedString() ITickEnclosedStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITickEnclosedStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITickEnclosedStringContext)
}

func (s *EnclosedStringContext) OptionsEnd() antlr.TerminalNode {
	return s.GetToken(aptParserOptionsEnd, 0)
}

func (s *EnclosedStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnclosedStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnclosedStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterEnclosedString(s)
	}
}

func (s *EnclosedStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitEnclosedString(s)
	}
}

func (p *aptParser) EnclosedString() (localctx IEnclosedStringContext) {
	this := p
	_ = this

	localctx = NewEnclosedStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, aptParserRULE_enclosedString)

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
		p.Match(aptParserOptionsStart)
	}
	{
		p.SetState(167)
		p.TickEnclosedString()
	}
	{
		p.SetState(168)
		p.Match(aptParserOptionsEnd)
	}

	return localctx
}

// ICdromURIContext is an interface to support dynamic dispatch.
type ICdromURIContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCdromURIContext differentiates from other interfaces.
	IsCdromURIContext()
}

type CdromURIContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCdromURIContext() *CdromURIContext {
	var p = new(CdromURIContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aptParserRULE_cdromURI
	return p
}

func (*CdromURIContext) IsCdromURIContext() {}

func NewCdromURIContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CdromURIContext {
	var p = new(CdromURIContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aptParserRULE_cdromURI

	return p
}

func (s *CdromURIContext) GetParser() antlr.Parser { return s.parser }

func (s *CdromURIContext) CdromSchema() antlr.TerminalNode {
	return s.GetToken(aptParserCdromSchema, 0)
}

func (s *CdromURIContext) Colon() antlr.TerminalNode {
	return s.GetToken(aptParserColon, 0)
}

func (s *CdromURIContext) EnclosedString() IEnclosedStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnclosedStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnclosedStringContext)
}

func (s *CdromURIContext) NonSpaceString() INonSpaceStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonSpaceStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonSpaceStringContext)
}

func (s *CdromURIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CdromURIContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CdromURIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.EnterCdromURI(s)
	}
}

func (s *CdromURIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aptListener); ok {
		listenerT.ExitCdromURI(s)
	}
}

func (p *aptParser) CdromURI() (localctx ICdromURIContext) {
	this := p
	_ = this

	localctx = NewCdromURIContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, aptParserRULE_cdromURI)

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
		p.SetState(170)
		p.Match(aptParserCdromSchema)
	}
	{
		p.SetState(171)
		p.Match(aptParserColon)
	}
	{
		p.SetState(172)
		p.EnclosedString()
	}
	{
		p.SetState(173)
		p.NonSpaceString()
	}

	return localctx
}
