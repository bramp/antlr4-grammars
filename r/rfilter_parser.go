// Generated from RFilter.g4 by ANTLR 4.7.

package r // RFilter
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 65, 149, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 
	3, 2, 7, 2, 16, 10, 2, 12, 2, 14, 2, 19, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 
	6, 3, 25, 10, 3, 13, 3, 14, 3, 26, 3, 4, 3, 4, 5, 4, 31, 10, 4, 3, 4, 3, 
	4, 3, 4, 5, 4, 36, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 42, 10, 4, 12, 
	4, 14, 4, 45, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 52, 10, 4, 12, 
	4, 14, 4, 55, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 61, 10, 4, 12, 4, 14, 
	4, 64, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 70, 10, 4, 12, 4, 14, 4, 73, 
	11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 79, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 
	84, 10, 4, 12, 4, 14, 4, 87, 11, 4, 3, 4, 3, 4, 5, 4, 91, 10, 4, 3, 4, 
	3, 4, 5, 4, 95, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 100, 10, 4, 12, 4, 14, 4, 
	103, 11, 4, 3, 4, 3, 4, 5, 4, 107, 10, 4, 3, 4, 3, 4, 5, 4, 111, 10, 4, 
	3, 4, 3, 4, 3, 4, 7, 4, 116, 10, 4, 12, 4, 14, 4, 119, 11, 4, 3, 4, 3, 
	4, 5, 4, 123, 10, 4, 3, 4, 3, 4, 5, 4, 127, 10, 4, 3, 4, 3, 4, 3, 4, 7, 
	4, 132, 10, 4, 12, 4, 14, 4, 135, 11, 4, 3, 4, 3, 4, 5, 4, 139, 10, 4, 
	3, 4, 3, 4, 5, 4, 143, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 2, 2, 7, 2, 
	4, 6, 8, 10, 2, 4, 4, 2, 47, 54, 57, 62, 8, 2, 7, 29, 31, 32, 43, 43, 45, 
	46, 55, 56, 63, 63, 2, 184, 2, 17, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 142, 
	3, 2, 2, 2, 8, 144, 3, 2, 2, 2, 10, 146, 3, 2, 2, 2, 12, 16, 5, 6, 4, 2, 
	13, 16, 7, 64, 2, 2, 14, 16, 7, 3, 2, 2, 15, 12, 3, 2, 2, 2, 15, 13, 3, 
	2, 2, 2, 15, 14, 3, 2, 2, 2, 16, 19, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 
	18, 3, 2, 2, 2, 18, 20, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 20, 21, 7, 2, 2, 
	3, 21, 3, 3, 2, 2, 2, 22, 23, 7, 64, 2, 2, 23, 25, 8, 3, 1, 2, 24, 22, 
	3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 
	27, 5, 3, 2, 2, 2, 28, 30, 5, 10, 6, 2, 29, 31, 5, 4, 3, 2, 30, 29, 3, 
	2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 143, 3, 2, 2, 2, 32, 143, 5, 8, 5, 2, 
	33, 35, 7, 38, 2, 2, 34, 36, 5, 4, 3, 2, 35, 34, 3, 2, 2, 2, 35, 36, 3, 
	2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 43, 8, 4, 1, 2, 38, 42, 5, 6, 4, 2, 39, 
	42, 7, 64, 2, 2, 40, 42, 7, 3, 2, 2, 41, 38, 3, 2, 2, 2, 41, 39, 3, 2, 
	2, 2, 41, 40, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 
	3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 47, 8, 4, 1, 2, 
	47, 143, 7, 39, 2, 2, 48, 53, 7, 36, 2, 2, 49, 52, 5, 6, 4, 2, 50, 52, 
	5, 4, 3, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 
	53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 53, 3, 
	2, 2, 2, 56, 143, 7, 37, 2, 2, 57, 62, 7, 6, 2, 2, 58, 61, 5, 6, 4, 2, 
	59, 61, 5, 4, 3, 2, 60, 58, 3, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 64, 3, 
	2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 
	62, 3, 2, 2, 2, 65, 143, 7, 5, 2, 2, 66, 71, 7, 4, 2, 2, 67, 70, 5, 6, 
	4, 2, 68, 70, 5, 4, 3, 2, 69, 67, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 73, 
	3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 74, 3, 2, 2, 2, 
	73, 71, 3, 2, 2, 2, 74, 75, 7, 5, 2, 2, 75, 143, 7, 5, 2, 2, 76, 78, 7, 
	35, 2, 2, 77, 79, 5, 4, 3, 2, 78, 77, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 
	80, 3, 2, 2, 2, 80, 85, 7, 36, 2, 2, 81, 84, 5, 6, 4, 2, 82, 84, 5, 4, 
	3, 2, 83, 81, 3, 2, 2, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 
	3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 88, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 
	88, 90, 7, 37, 2, 2, 89, 91, 5, 4, 3, 2, 90, 89, 3, 2, 2, 2, 90, 91, 3, 
	2, 2, 2, 91, 143, 3, 2, 2, 2, 92, 94, 7, 42, 2, 2, 93, 95, 5, 4, 3, 2, 
	94, 93, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 101, 7, 
	36, 2, 2, 97, 100, 5, 6, 4, 2, 98, 100, 5, 4, 3, 2, 99, 97, 3, 2, 2, 2, 
	99, 98, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 
	3, 2, 2, 2, 102, 104, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104, 106, 7, 37, 
	2, 2, 105, 107, 5, 4, 3, 2, 106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 
	107, 143, 3, 2, 2, 2, 108, 110, 7, 44, 2, 2, 109, 111, 5, 4, 3, 2, 110, 
	109, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 117, 
	7, 36, 2, 2, 113, 116, 5, 6, 4, 2, 114, 116, 5, 4, 3, 2, 115, 113, 3, 2, 
	2, 2, 115, 114, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 
	117, 118, 3, 2, 2, 2, 118, 120, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 
	122, 7, 37, 2, 2, 121, 123, 5, 4, 3, 2, 122, 121, 3, 2, 2, 2, 122, 123, 
	3, 2, 2, 2, 123, 143, 3, 2, 2, 2, 124, 126, 7, 40, 2, 2, 125, 127, 5, 4, 
	3, 2, 126, 125, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 
	128, 133, 7, 36, 2, 2, 129, 132, 5, 6, 4, 2, 130, 132, 5, 4, 3, 2, 131, 
	129, 3, 2, 2, 2, 131, 130, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 
	3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 136, 3, 2, 2, 2, 135, 133, 3, 2, 
	2, 2, 136, 138, 7, 37, 2, 2, 137, 139, 5, 4, 3, 2, 138, 137, 3, 2, 2, 2, 
	138, 139, 3, 2, 2, 2, 139, 143, 3, 2, 2, 2, 140, 141, 7, 41, 2, 2, 141, 
	143, 8, 4, 1, 2, 142, 28, 3, 2, 2, 2, 142, 32, 3, 2, 2, 2, 142, 33, 3, 
	2, 2, 2, 142, 48, 3, 2, 2, 2, 142, 57, 3, 2, 2, 2, 142, 66, 3, 2, 2, 2, 
	142, 76, 3, 2, 2, 2, 142, 92, 3, 2, 2, 2, 142, 108, 3, 2, 2, 2, 142, 124, 
	3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 143, 7, 3, 2, 2, 2, 144, 145, 9, 2, 2, 
	2, 145, 9, 3, 2, 2, 2, 146, 147, 9, 3, 2, 2, 147, 11, 3, 2, 2, 2, 32, 15, 
	17, 26, 30, 35, 41, 43, 51, 53, 60, 62, 69, 71, 78, 83, 85, 90, 94, 99, 
	101, 106, 110, 115, 117, 122, 126, 131, 133, 138, 142,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'[['", "']'", "'['", "'::'", "':::'", "'$'", "'@'", "'^'", 
	"'-'", "'+'", "':'", "'*'", "'/'", "'>'", "'>='", "'<'", "'<='", "'=='", 
	"'!='", "'!'", "'&'", "'&&'", "'|'", "'||'", "'~'", "'<-'", "'<<-'", "'='", 
	"'->'", "'->>'", "':='", "'function'", "'('", "')'", "'{'", "'}'", "'if'", 
	"'else'", "'for'", "'in'", "'while'", "'repeat'", "'?'", "'next'", "'break'", 
	"'NULL'", "'NA'", "'Inf'", "'NaN'", "'TRUE'", "'FALSE'", "','", "'...'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "HEX", "INT", "FLOAT", "COMPLEX", "STRING", "ID", "USER_OP", "NL", 
	"WS",
}

var ruleNames = []string{
	"stream", "eat", "elem", "atom", "op",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RFilter struct {
	*antlr.BaseParser
}

func NewRFilter(input antlr.TokenStream) *RFilter {
	this := new(RFilter)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "RFilter.g4"

	return this
}


protected int curlies = 0;



// RFilter tokens.
const (
	RFilterEOF = antlr.TokenEOF
	RFilterT__0 = 1
	RFilterT__1 = 2
	RFilterT__2 = 3
	RFilterT__3 = 4
	RFilterT__4 = 5
	RFilterT__5 = 6
	RFilterT__6 = 7
	RFilterT__7 = 8
	RFilterT__8 = 9
	RFilterT__9 = 10
	RFilterT__10 = 11
	RFilterT__11 = 12
	RFilterT__12 = 13
	RFilterT__13 = 14
	RFilterT__14 = 15
	RFilterT__15 = 16
	RFilterT__16 = 17
	RFilterT__17 = 18
	RFilterT__18 = 19
	RFilterT__19 = 20
	RFilterT__20 = 21
	RFilterT__21 = 22
	RFilterT__22 = 23
	RFilterT__23 = 24
	RFilterT__24 = 25
	RFilterT__25 = 26
	RFilterT__26 = 27
	RFilterT__27 = 28
	RFilterT__28 = 29
	RFilterT__29 = 30
	RFilterT__30 = 31
	RFilterT__31 = 32
	RFilterT__32 = 33
	RFilterT__33 = 34
	RFilterT__34 = 35
	RFilterT__35 = 36
	RFilterT__36 = 37
	RFilterT__37 = 38
	RFilterT__38 = 39
	RFilterT__39 = 40
	RFilterT__40 = 41
	RFilterT__41 = 42
	RFilterT__42 = 43
	RFilterT__43 = 44
	RFilterT__44 = 45
	RFilterT__45 = 46
	RFilterT__46 = 47
	RFilterT__47 = 48
	RFilterT__48 = 49
	RFilterT__49 = 50
	RFilterT__50 = 51
	RFilterT__51 = 52
	RFilterT__52 = 53
	RFilterT__53 = 54
	RFilterHEX = 55
	RFilterINT = 56
	RFilterFLOAT = 57
	RFilterCOMPLEX = 58
	RFilterSTRING = 59
	RFilterID = 60
	RFilterUSER_OP = 61
	RFilterNL = 62
	RFilterWS = 63
)

// RFilter rules.
const (
	RFilterRULE_stream = 0
	RFilterRULE_eat = 1
	RFilterRULE_elem = 2
	RFilterRULE_atom = 3
	RFilterRULE_op = 4
)

// IStreamContext is an interface to support dynamic dispatch.
type IStreamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStreamContext differentiates from other interfaces.
	IsStreamContext()
}

type StreamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreamContext() *StreamContext {
	var p = new(StreamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RFilterRULE_stream
	return p
}

func (*StreamContext) IsStreamContext() {}

func NewStreamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreamContext {
	var p = new(StreamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RFilterRULE_stream

	return p
}

func (s *StreamContext) GetParser() antlr.Parser { return s.parser }

func (s *StreamContext) EOF() antlr.TerminalNode {
	return s.GetToken(RFilterEOF, 0)
}

func (s *StreamContext) AllElem() []IElemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElemContext)(nil)).Elem())
	var tst = make([]IElemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElemContext)
		}
	}

	return tst
}

func (s *StreamContext) Elem(i int) IElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElemContext)
}

func (s *StreamContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RFilterNL)
}

func (s *StreamContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RFilterNL, i)
}

func (s *StreamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StreamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RFilterListener); ok {
		listenerT.EnterStream(s)
	}
}

func (s *StreamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RFilterListener); ok {
		listenerT.ExitStream(s)
	}
}

func (s *StreamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RFilterVisitor:
		return t.VisitStream(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RFilter) Stream() (localctx IStreamContext) {
	localctx = NewStreamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RFilterRULE_stream)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RFilterT__0) | (1 << RFilterT__1) | (1 << RFilterT__3) | (1 << RFilterT__4) | (1 << RFilterT__5) | (1 << RFilterT__6) | (1 << RFilterT__7) | (1 << RFilterT__8) | (1 << RFilterT__9) | (1 << RFilterT__10) | (1 << RFilterT__11) | (1 << RFilterT__12) | (1 << RFilterT__13) | (1 << RFilterT__14) | (1 << RFilterT__15) | (1 << RFilterT__16) | (1 << RFilterT__17) | (1 << RFilterT__18) | (1 << RFilterT__19) | (1 << RFilterT__20) | (1 << RFilterT__21) | (1 << RFilterT__22) | (1 << RFilterT__23) | (1 << RFilterT__24) | (1 << RFilterT__25) | (1 << RFilterT__26) | (1 << RFilterT__28) | (1 << RFilterT__29))) != 0) || ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (RFilterT__32 - 33)) | (1 << (RFilterT__33 - 33)) | (1 << (RFilterT__35 - 33)) | (1 << (RFilterT__37 - 33)) | (1 << (RFilterT__38 - 33)) | (1 << (RFilterT__39 - 33)) | (1 << (RFilterT__40 - 33)) | (1 << (RFilterT__41 - 33)) | (1 << (RFilterT__42 - 33)) | (1 << (RFilterT__43 - 33)) | (1 << (RFilterT__44 - 33)) | (1 << (RFilterT__45 - 33)) | (1 << (RFilterT__46 - 33)) | (1 << (RFilterT__47 - 33)) | (1 << (RFilterT__48 - 33)) | (1 << (RFilterT__49 - 33)) | (1 << (RFilterT__50 - 33)) | (1 << (RFilterT__51 - 33)) | (1 << (RFilterT__52 - 33)) | (1 << (RFilterT__53 - 33)) | (1 << (RFilterHEX - 33)) | (1 << (RFilterINT - 33)) | (1 << (RFilterFLOAT - 33)) | (1 << (RFilterCOMPLEX - 33)) | (1 << (RFilterSTRING - 33)) | (1 << (RFilterID - 33)) | (1 << (RFilterUSER_OP - 33)) | (1 << (RFilterNL - 33)))) != 0) {
		p.SetState(13)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case RFilterT__1, RFilterT__3, RFilterT__4, RFilterT__5, RFilterT__6, RFilterT__7, RFilterT__8, RFilterT__9, RFilterT__10, RFilterT__11, RFilterT__12, RFilterT__13, RFilterT__14, RFilterT__15, RFilterT__16, RFilterT__17, RFilterT__18, RFilterT__19, RFilterT__20, RFilterT__21, RFilterT__22, RFilterT__23, RFilterT__24, RFilterT__25, RFilterT__26, RFilterT__28, RFilterT__29, RFilterT__32, RFilterT__33, RFilterT__35, RFilterT__37, RFilterT__38, RFilterT__39, RFilterT__40, RFilterT__41, RFilterT__42, RFilterT__43, RFilterT__44, RFilterT__45, RFilterT__46, RFilterT__47, RFilterT__48, RFilterT__49, RFilterT__50, RFilterT__51, RFilterT__52, RFilterT__53, RFilterHEX, RFilterINT, RFilterFLOAT, RFilterCOMPLEX, RFilterSTRING, RFilterID, RFilterUSER_OP:
			{
				p.SetState(10)
				p.Elem()
			}


		case RFilterNL:
			{
				p.SetState(11)
				p.Match(RFilterNL)
			}


		case RFilterT__0:
			{
				p.SetState(12)
				p.Match(RFilterT__0)
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
		p.Match(RFilterEOF)
	}



	return localctx
}


// IEatContext is an interface to support dynamic dispatch.
type IEatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NL returns the _NL token.
	Get_NL() antlr.Token 


	// Set_NL sets the _NL token.
	Set_NL(antlr.Token) 


	// IsEatContext differentiates from other interfaces.
	IsEatContext()
}

type EatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_NL antlr.Token
}

func NewEmptyEatContext() *EatContext {
	var p = new(EatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RFilterRULE_eat
	return p
}

func (*EatContext) IsEatContext() {}

func NewEatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EatContext {
	var p = new(EatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RFilterRULE_eat

	return p
}

func (s *EatContext) GetParser() antlr.Parser { return s.parser }

func (s *EatContext) Get_NL() antlr.Token { return s._NL }


func (s *EatContext) Set_NL(v antlr.Token) { s._NL = v }


func (s *EatContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RFilterNL)
}

func (s *EatContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RFilterNL, i)
}

func (s *EatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RFilterListener); ok {
		listenerT.EnterEat(s)
	}
}

func (s *EatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RFilterListener); ok {
		listenerT.ExitEat(s)
	}
}

func (s *EatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RFilterVisitor:
		return t.VisitEat(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RFilter) Eat() (localctx IEatContext) {
	localctx = NewEatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RFilterRULE_eat)

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
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(20)

					var _m = p.Match(RFilterNL)

					localctx.(*EatContext)._NL = _m
				}
				((WritableToken)localctx.(*EatContext).Get_NL()).setChannel(Token.HIDDEN_CHANNEL);




		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}



	return localctx
}


// IElemContext is an interface to support dynamic dispatch.
type IElemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElemContext differentiates from other interfaces.
	IsElemContext()
}

type ElemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElemContext() *ElemContext {
	var p = new(ElemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RFilterRULE_elem
	return p
}

func (*ElemContext) IsElemContext() {}

func NewElemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElemContext {
	var p = new(ElemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RFilterRULE_elem

	return p
}

func (s *ElemContext) GetParser() antlr.Parser { return s.parser }

func (s *ElemContext) Op() IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *ElemContext) AllEat() []IEatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEatContext)(nil)).Elem())
	var tst = make([]IEatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEatContext)
		}
	}

	return tst
}

func (s *ElemContext) Eat(i int) IEatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEatContext)
}

func (s *ElemContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ElemContext) AllElem() []IElemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElemContext)(nil)).Elem())
	var tst = make([]IElemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElemContext)
		}
	}

	return tst
}

func (s *ElemContext) Elem(i int) IElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElemContext)
}

func (s *ElemContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RFilterNL)
}

func (s *ElemContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RFilterNL, i)
}

func (s *ElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RFilterListener); ok {
		listenerT.EnterElem(s)
	}
}

func (s *ElemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RFilterListener); ok {
		listenerT.ExitElem(s)
	}
}

func (s *ElemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RFilterVisitor:
		return t.VisitElem(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RFilter) Elem() (localctx IElemContext) {
	localctx = NewElemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RFilterRULE_elem)
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

	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RFilterT__4, RFilterT__5, RFilterT__6, RFilterT__7, RFilterT__8, RFilterT__9, RFilterT__10, RFilterT__11, RFilterT__12, RFilterT__13, RFilterT__14, RFilterT__15, RFilterT__16, RFilterT__17, RFilterT__18, RFilterT__19, RFilterT__20, RFilterT__21, RFilterT__22, RFilterT__23, RFilterT__24, RFilterT__25, RFilterT__26, RFilterT__28, RFilterT__29, RFilterT__40, RFilterT__42, RFilterT__43, RFilterT__52, RFilterT__53, RFilterUSER_OP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.Op()
		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(27)
				p.Eat()
			}


		}


	case RFilterT__44, RFilterT__45, RFilterT__46, RFilterT__47, RFilterT__48, RFilterT__49, RFilterT__50, RFilterT__51, RFilterHEX, RFilterINT, RFilterFLOAT, RFilterCOMPLEX, RFilterSTRING, RFilterID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Atom()
		}


	case RFilterT__35:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(31)
			p.Match(RFilterT__35)
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(32)
				p.Eat()
			}


		}
		curlies++;
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RFilterT__0) | (1 << RFilterT__1) | (1 << RFilterT__3) | (1 << RFilterT__4) | (1 << RFilterT__5) | (1 << RFilterT__6) | (1 << RFilterT__7) | (1 << RFilterT__8) | (1 << RFilterT__9) | (1 << RFilterT__10) | (1 << RFilterT__11) | (1 << RFilterT__12) | (1 << RFilterT__13) | (1 << RFilterT__14) | (1 << RFilterT__15) | (1 << RFilterT__16) | (1 << RFilterT__17) | (1 << RFilterT__18) | (1 << RFilterT__19) | (1 << RFilterT__20) | (1 << RFilterT__21) | (1 << RFilterT__22) | (1 << RFilterT__23) | (1 << RFilterT__24) | (1 << RFilterT__25) | (1 << RFilterT__26) | (1 << RFilterT__28) | (1 << RFilterT__29))) != 0) || ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (RFilterT__32 - 33)) | (1 << (RFilterT__33 - 33)) | (1 << (RFilterT__35 - 33)) | (1 << (RFilterT__37 - 33)) | (1 << (RFilterT__38 - 33)) | (1 << (RFilterT__39 - 33)) | (1 << (RFilterT__40 - 33)) | (1 << (RFilterT__41 - 33)) | (1 << (RFilterT__42 - 33)) | (1 << (RFilterT__43 - 33)) | (1 << (RFilterT__44 - 33)) | (1 << (RFilterT__45 - 33)) | (1 << (RFilterT__46 - 33)) | (1 << (RFilterT__47 - 33)) | (1 << (RFilterT__48 - 33)) | (1 << (RFilterT__49 - 33)) | (1 << (RFilterT__50 - 33)) | (1 << (RFilterT__51 - 33)) | (1 << (RFilterT__52 - 33)) | (1 << (RFilterT__53 - 33)) | (1 << (RFilterHEX - 33)) | (1 << (RFilterINT - 33)) | (1 << (RFilterFLOAT - 33)) | (1 << (RFilterCOMPLEX - 33)) | (1 << (RFilterSTRING - 33)) | (1 << (RFilterID - 33)) | (1 << (RFilterUSER_OP - 33)) | (1 << (RFilterNL - 33)))) != 0) {
			p.SetState(39)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RFilterT__1, RFilterT__3, RFilterT__4, RFilterT__5, RFilterT__6, RFilterT__7, RFilterT__8, RFilterT__9, RFilterT__10, RFilterT__11, RFilterT__12, RFilterT__13, RFilterT__14, RFilterT__15, RFilterT__16, RFilterT__17, RFilterT__18, RFilterT__19, RFilterT__20, RFilterT__21, RFilterT__22, RFilterT__23, RFilterT__24, RFilterT__25, RFilterT__26, RFilterT__28, RFilterT__29, RFilterT__32, RFilterT__33, RFilterT__35, RFilterT__37, RFilterT__38, RFilterT__39, RFilterT__40, RFilterT__41, RFilterT__42, RFilterT__43, RFilterT__44, RFilterT__45, RFilterT__46, RFilterT__47, RFilterT__48, RFilterT__49, RFilterT__50, RFilterT__51, RFilterT__52, RFilterT__53, RFilterHEX, RFilterINT, RFilterFLOAT, RFilterCOMPLEX, RFilterSTRING, RFilterID, RFilterUSER_OP:
				{
					p.SetState(36)
					p.Elem()
				}


			case RFilterNL:
				{
					p.SetState(37)
					p.Match(RFilterNL)
				}


			case RFilterT__0:
				{
					p.SetState(38)
					p.Match(RFilterT__0)
				}



			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		curlies--;
		{
			p.SetState(45)
			p.Match(RFilterT__36)
		}


	case RFilterT__33:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(46)
			p.Match(RFilterT__33)
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RFilterT__1) | (1 << RFilterT__3) | (1 << RFilterT__4) | (1 << RFilterT__5) | (1 << RFilterT__6) | (1 << RFilterT__7) | (1 << RFilterT__8) | (1 << RFilterT__9) | (1 << RFilterT__10) | (1 << RFilterT__11) | (1 << RFilterT__12) | (1 << RFilterT__13) | (1 << RFilterT__14) | (1 << RFilterT__15) | (1 << RFilterT__16) | (1 << RFilterT__17) | (1 << RFilterT__18) | (1 << RFilterT__19) | (1 << RFilterT__20) | (1 << RFilterT__21) | (1 << RFilterT__22) | (1 << RFilterT__23) | (1 << RFilterT__24) | (1 << RFilterT__25) | (1 << RFilterT__26) | (1 << RFilterT__28) | (1 << RFilterT__29))) != 0) || ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (RFilterT__32 - 33)) | (1 << (RFilterT__33 - 33)) | (1 << (RFilterT__35 - 33)) | (1 << (RFilterT__37 - 33)) | (1 << (RFilterT__38 - 33)) | (1 << (RFilterT__39 - 33)) | (1 << (RFilterT__40 - 33)) | (1 << (RFilterT__41 - 33)) | (1 << (RFilterT__42 - 33)) | (1 << (RFilterT__43 - 33)) | (1 << (RFilterT__44 - 33)) | (1 << (RFilterT__45 - 33)) | (1 << (RFilterT__46 - 33)) | (1 << (RFilterT__47 - 33)) | (1 << (RFilterT__48 - 33)) | (1 << (RFilterT__49 - 33)) | (1 << (RFilterT__50 - 33)) | (1 << (RFilterT__51 - 33)) | (1 << (RFilterT__52 - 33)) | (1 << (RFilterT__53 - 33)) | (1 << (RFilterHEX - 33)) | (1 << (RFilterINT - 33)) | (1 << (RFilterFLOAT - 33)) | (1 << (RFilterCOMPLEX - 33)) | (1 << (RFilterSTRING - 33)) | (1 << (RFilterID - 33)) | (1 << (RFilterUSER_OP - 33)) | (1 << (RFilterNL - 33)))) != 0) {
			p.SetState(49)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RFilterT__1, RFilterT__3, RFilterT__4, RFilterT__5, RFilterT__6, RFilterT__7, RFilterT__8, RFilterT__9, RFilterT__10, RFilterT__11, RFilterT__12, RFilterT__13, RFilterT__14, RFilterT__15, RFilterT__16, RFilterT__17, RFilterT__18, RFilterT__19, RFilterT__20, RFilterT__21, RFilterT__22, RFilterT__23, RFilterT__24, RFilterT__25, RFilterT__26, RFilterT__28, RFilterT__29, RFilterT__32, RFilterT__33, RFilterT__35, RFilterT__37, RFilterT__38, RFilterT__39, RFilterT__40, RFilterT__41, RFilterT__42, RFilterT__43, RFilterT__44, RFilterT__45, RFilterT__46, RFilterT__47, RFilterT__48, RFilterT__49, RFilterT__50, RFilterT__51, RFilterT__52, RFilterT__53, RFilterHEX, RFilterINT, RFilterFLOAT, RFilterCOMPLEX, RFilterSTRING, RFilterID, RFilterUSER_OP:
				{
					p.SetState(47)
					p.Elem()
				}


			case RFilterNL:
				{
					p.SetState(48)
					p.Eat()
				}



			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(53)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(54)
			p.Match(RFilterT__34)
		}


	case RFilterT__3:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(55)
			p.Match(RFilterT__3)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RFilterT__1) | (1 << RFilterT__3) | (1 << RFilterT__4) | (1 << RFilterT__5) | (1 << RFilterT__6) | (1 << RFilterT__7) | (1 << RFilterT__8) | (1 << RFilterT__9) | (1 << RFilterT__10) | (1 << RFilterT__11) | (1 << RFilterT__12) | (1 << RFilterT__13) | (1 << RFilterT__14) | (1 << RFilterT__15) | (1 << RFilterT__16) | (1 << RFilterT__17) | (1 << RFilterT__18) | (1 << RFilterT__19) | (1 << RFilterT__20) | (1 << RFilterT__21) | (1 << RFilterT__22) | (1 << RFilterT__23) | (1 << RFilterT__24) | (1 << RFilterT__25) | (1 << RFilterT__26) | (1 << RFilterT__28) | (1 << RFilterT__29))) != 0) || ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (RFilterT__32 - 33)) | (1 << (RFilterT__33 - 33)) | (1 << (RFilterT__35 - 33)) | (1 << (RFilterT__37 - 33)) | (1 << (RFilterT__38 - 33)) | (1 << (RFilterT__39 - 33)) | (1 << (RFilterT__40 - 33)) | (1 << (RFilterT__41 - 33)) | (1 << (RFilterT__42 - 33)) | (1 << (RFilterT__43 - 33)) | (1 << (RFilterT__44 - 33)) | (1 << (RFilterT__45 - 33)) | (1 << (RFilterT__46 - 33)) | (1 << (RFilterT__47 - 33)) | (1 << (RFilterT__48 - 33)) | (1 << (RFilterT__49 - 33)) | (1 << (RFilterT__50 - 33)) | (1 << (RFilterT__51 - 33)) | (1 << (RFilterT__52 - 33)) | (1 << (RFilterT__53 - 33)) | (1 << (RFilterHEX - 33)) | (1 << (RFilterINT - 33)) | (1 << (RFilterFLOAT - 33)) | (1 << (RFilterCOMPLEX - 33)) | (1 << (RFilterSTRING - 33)) | (1 << (RFilterID - 33)) | (1 << (RFilterUSER_OP - 33)) | (1 << (RFilterNL - 33)))) != 0) {
			p.SetState(58)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RFilterT__1, RFilterT__3, RFilterT__4, RFilterT__5, RFilterT__6, RFilterT__7, RFilterT__8, RFilterT__9, RFilterT__10, RFilterT__11, RFilterT__12, RFilterT__13, RFilterT__14, RFilterT__15, RFilterT__16, RFilterT__17, RFilterT__18, RFilterT__19, RFilterT__20, RFilterT__21, RFilterT__22, RFilterT__23, RFilterT__24, RFilterT__25, RFilterT__26, RFilterT__28, RFilterT__29, RFilterT__32, RFilterT__33, RFilterT__35, RFilterT__37, RFilterT__38, RFilterT__39, RFilterT__40, RFilterT__41, RFilterT__42, RFilterT__43, RFilterT__44, RFilterT__45, RFilterT__46, RFilterT__47, RFilterT__48, RFilterT__49, RFilterT__50, RFilterT__51, RFilterT__52, RFilterT__53, RFilterHEX, RFilterINT, RFilterFLOAT, RFilterCOMPLEX, RFilterSTRING, RFilterID, RFilterUSER_OP:
				{
					p.SetState(56)
					p.Elem()
				}


			case RFilterNL:
				{
					p.SetState(57)
					p.Eat()
				}



			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(63)
			p.Match(RFilterT__2)
		}


	case RFilterT__1:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(64)
			p.Match(RFilterT__1)
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RFilterT__1) | (1 << RFilterT__3) | (1 << RFilterT__4) | (1 << RFilterT__5) | (1 << RFilterT__6) | (1 << RFilterT__7) | (1 << RFilterT__8) | (1 << RFilterT__9) | (1 << RFilterT__10) | (1 << RFilterT__11) | (1 << RFilterT__12) | (1 << RFilterT__13) | (1 << RFilterT__14) | (1 << RFilterT__15) | (1 << RFilterT__16) | (1 << RFilterT__17) | (1 << RFilterT__18) | (1 << RFilterT__19) | (1 << RFilterT__20) | (1 << RFilterT__21) | (1 << RFilterT__22) | (1 << RFilterT__23) | (1 << RFilterT__24) | (1 << RFilterT__25) | (1 << RFilterT__26) | (1 << RFilterT__28) | (1 << RFilterT__29))) != 0) || ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (RFilterT__32 - 33)) | (1 << (RFilterT__33 - 33)) | (1 << (RFilterT__35 - 33)) | (1 << (RFilterT__37 - 33)) | (1 << (RFilterT__38 - 33)) | (1 << (RFilterT__39 - 33)) | (1 << (RFilterT__40 - 33)) | (1 << (RFilterT__41 - 33)) | (1 << (RFilterT__42 - 33)) | (1 << (RFilterT__43 - 33)) | (1 << (RFilterT__44 - 33)) | (1 << (RFilterT__45 - 33)) | (1 << (RFilterT__46 - 33)) | (1 << (RFilterT__47 - 33)) | (1 << (RFilterT__48 - 33)) | (1 << (RFilterT__49 - 33)) | (1 << (RFilterT__50 - 33)) | (1 << (RFilterT__51 - 33)) | (1 << (RFilterT__52 - 33)) | (1 << (RFilterT__53 - 33)) | (1 << (RFilterHEX - 33)) | (1 << (RFilterINT - 33)) | (1 << (RFilterFLOAT - 33)) | (1 << (RFilterCOMPLEX - 33)) | (1 << (RFilterSTRING - 33)) | (1 << (RFilterID - 33)) | (1 << (RFilterUSER_OP - 33)) | (1 << (RFilterNL - 33)))) != 0) {
			p.SetState(67)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RFilterT__1, RFilterT__3, RFilterT__4, RFilterT__5, RFilterT__6, RFilterT__7, RFilterT__8, RFilterT__9, RFilterT__10, RFilterT__11, RFilterT__12, RFilterT__13, RFilterT__14, RFilterT__15, RFilterT__16, RFilterT__17, RFilterT__18, RFilterT__19, RFilterT__20, RFilterT__21, RFilterT__22, RFilterT__23, RFilterT__24, RFilterT__25, RFilterT__26, RFilterT__28, RFilterT__29, RFilterT__32, RFilterT__33, RFilterT__35, RFilterT__37, RFilterT__38, RFilterT__39, RFilterT__40, RFilterT__41, RFilterT__42, RFilterT__43, RFilterT__44, RFilterT__45, RFilterT__46, RFilterT__47, RFilterT__48, RFilterT__49, RFilterT__50, RFilterT__51, RFilterT__52, RFilterT__53, RFilterHEX, RFilterINT, RFilterFLOAT, RFilterCOMPLEX, RFilterSTRING, RFilterID, RFilterUSER_OP:
				{
					p.SetState(65)
					p.Elem()
				}


			case RFilterNL:
				{
					p.SetState(66)
					p.Eat()
				}



			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(72)
			p.Match(RFilterT__2)
		}
		{
			p.SetState(73)
			p.Match(RFilterT__2)
		}


	case RFilterT__32:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(74)
			p.Match(RFilterT__32)
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == RFilterNL {
			{
				p.SetState(75)
				p.Eat()
			}

		}
		{
			p.SetState(78)
			p.Match(RFilterT__33)
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RFilterT__1) | (1 << RFilterT__3) | (1 << RFilterT__4) | (1 << RFilterT__5) | (1 << RFilterT__6) | (1 << RFilterT__7) | (1 << RFilterT__8) | (1 << RFilterT__9) | (1 << RFilterT__10) | (1 << RFilterT__11) | (1 << RFilterT__12) | (1 << RFilterT__13) | (1 << RFilterT__14) | (1 << RFilterT__15) | (1 << RFilterT__16) | (1 << RFilterT__17) | (1 << RFilterT__18) | (1 << RFilterT__19) | (1 << RFilterT__20) | (1 << RFilterT__21) | (1 << RFilterT__22) | (1 << RFilterT__23) | (1 << RFilterT__24) | (1 << RFilterT__25) | (1 << RFilterT__26) | (1 << RFilterT__28) | (1 << RFilterT__29))) != 0) || ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (RFilterT__32 - 33)) | (1 << (RFilterT__33 - 33)) | (1 << (RFilterT__35 - 33)) | (1 << (RFilterT__37 - 33)) | (1 << (RFilterT__38 - 33)) | (1 << (RFilterT__39 - 33)) | (1 << (RFilterT__40 - 33)) | (1 << (RFilterT__41 - 33)) | (1 << (RFilterT__42 - 33)) | (1 << (RFilterT__43 - 33)) | (1 << (RFilterT__44 - 33)) | (1 << (RFilterT__45 - 33)) | (1 << (RFilterT__46 - 33)) | (1 << (RFilterT__47 - 33)) | (1 << (RFilterT__48 - 33)) | (1 << (RFilterT__49 - 33)) | (1 << (RFilterT__50 - 33)) | (1 << (RFilterT__51 - 33)) | (1 << (RFilterT__52 - 33)) | (1 << (RFilterT__53 - 33)) | (1 << (RFilterHEX - 33)) | (1 << (RFilterINT - 33)) | (1 << (RFilterFLOAT - 33)) | (1 << (RFilterCOMPLEX - 33)) | (1 << (RFilterSTRING - 33)) | (1 << (RFilterID - 33)) | (1 << (RFilterUSER_OP - 33)) | (1 << (RFilterNL - 33)))) != 0) {
			p.SetState(81)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RFilterT__1, RFilterT__3, RFilterT__4, RFilterT__5, RFilterT__6, RFilterT__7, RFilterT__8, RFilterT__9, RFilterT__10, RFilterT__11, RFilterT__12, RFilterT__13, RFilterT__14, RFilterT__15, RFilterT__16, RFilterT__17, RFilterT__18, RFilterT__19, RFilterT__20, RFilterT__21, RFilterT__22, RFilterT__23, RFilterT__24, RFilterT__25, RFilterT__26, RFilterT__28, RFilterT__29, RFilterT__32, RFilterT__33, RFilterT__35, RFilterT__37, RFilterT__38, RFilterT__39, RFilterT__40, RFilterT__41, RFilterT__42, RFilterT__43, RFilterT__44, RFilterT__45, RFilterT__46, RFilterT__47, RFilterT__48, RFilterT__49, RFilterT__50, RFilterT__51, RFilterT__52, RFilterT__53, RFilterHEX, RFilterINT, RFilterFLOAT, RFilterCOMPLEX, RFilterSTRING, RFilterID, RFilterUSER_OP:
				{
					p.SetState(79)
					p.Elem()
				}


			case RFilterNL:
				{
					p.SetState(80)
					p.Eat()
				}



			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(86)
			p.Match(RFilterT__34)
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(87)
				p.Eat()
			}


		}


	case RFilterT__39:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(90)
			p.Match(RFilterT__39)
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == RFilterNL {
			{
				p.SetState(91)
				p.Eat()
			}

		}
		{
			p.SetState(94)
			p.Match(RFilterT__33)
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RFilterT__1) | (1 << RFilterT__3) | (1 << RFilterT__4) | (1 << RFilterT__5) | (1 << RFilterT__6) | (1 << RFilterT__7) | (1 << RFilterT__8) | (1 << RFilterT__9) | (1 << RFilterT__10) | (1 << RFilterT__11) | (1 << RFilterT__12) | (1 << RFilterT__13) | (1 << RFilterT__14) | (1 << RFilterT__15) | (1 << RFilterT__16) | (1 << RFilterT__17) | (1 << RFilterT__18) | (1 << RFilterT__19) | (1 << RFilterT__20) | (1 << RFilterT__21) | (1 << RFilterT__22) | (1 << RFilterT__23) | (1 << RFilterT__24) | (1 << RFilterT__25) | (1 << RFilterT__26) | (1 << RFilterT__28) | (1 << RFilterT__29))) != 0) || ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (RFilterT__32 - 33)) | (1 << (RFilterT__33 - 33)) | (1 << (RFilterT__35 - 33)) | (1 << (RFilterT__37 - 33)) | (1 << (RFilterT__38 - 33)) | (1 << (RFilterT__39 - 33)) | (1 << (RFilterT__40 - 33)) | (1 << (RFilterT__41 - 33)) | (1 << (RFilterT__42 - 33)) | (1 << (RFilterT__43 - 33)) | (1 << (RFilterT__44 - 33)) | (1 << (RFilterT__45 - 33)) | (1 << (RFilterT__46 - 33)) | (1 << (RFilterT__47 - 33)) | (1 << (RFilterT__48 - 33)) | (1 << (RFilterT__49 - 33)) | (1 << (RFilterT__50 - 33)) | (1 << (RFilterT__51 - 33)) | (1 << (RFilterT__52 - 33)) | (1 << (RFilterT__53 - 33)) | (1 << (RFilterHEX - 33)) | (1 << (RFilterINT - 33)) | (1 << (RFilterFLOAT - 33)) | (1 << (RFilterCOMPLEX - 33)) | (1 << (RFilterSTRING - 33)) | (1 << (RFilterID - 33)) | (1 << (RFilterUSER_OP - 33)) | (1 << (RFilterNL - 33)))) != 0) {
			p.SetState(97)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RFilterT__1, RFilterT__3, RFilterT__4, RFilterT__5, RFilterT__6, RFilterT__7, RFilterT__8, RFilterT__9, RFilterT__10, RFilterT__11, RFilterT__12, RFilterT__13, RFilterT__14, RFilterT__15, RFilterT__16, RFilterT__17, RFilterT__18, RFilterT__19, RFilterT__20, RFilterT__21, RFilterT__22, RFilterT__23, RFilterT__24, RFilterT__25, RFilterT__26, RFilterT__28, RFilterT__29, RFilterT__32, RFilterT__33, RFilterT__35, RFilterT__37, RFilterT__38, RFilterT__39, RFilterT__40, RFilterT__41, RFilterT__42, RFilterT__43, RFilterT__44, RFilterT__45, RFilterT__46, RFilterT__47, RFilterT__48, RFilterT__49, RFilterT__50, RFilterT__51, RFilterT__52, RFilterT__53, RFilterHEX, RFilterINT, RFilterFLOAT, RFilterCOMPLEX, RFilterSTRING, RFilterID, RFilterUSER_OP:
				{
					p.SetState(95)
					p.Elem()
				}


			case RFilterNL:
				{
					p.SetState(96)
					p.Eat()
				}



			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(102)
			p.Match(RFilterT__34)
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(103)
				p.Eat()
			}


		}


	case RFilterT__41:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(106)
			p.Match(RFilterT__41)
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == RFilterNL {
			{
				p.SetState(107)
				p.Eat()
			}

		}
		{
			p.SetState(110)
			p.Match(RFilterT__33)
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RFilterT__1) | (1 << RFilterT__3) | (1 << RFilterT__4) | (1 << RFilterT__5) | (1 << RFilterT__6) | (1 << RFilterT__7) | (1 << RFilterT__8) | (1 << RFilterT__9) | (1 << RFilterT__10) | (1 << RFilterT__11) | (1 << RFilterT__12) | (1 << RFilterT__13) | (1 << RFilterT__14) | (1 << RFilterT__15) | (1 << RFilterT__16) | (1 << RFilterT__17) | (1 << RFilterT__18) | (1 << RFilterT__19) | (1 << RFilterT__20) | (1 << RFilterT__21) | (1 << RFilterT__22) | (1 << RFilterT__23) | (1 << RFilterT__24) | (1 << RFilterT__25) | (1 << RFilterT__26) | (1 << RFilterT__28) | (1 << RFilterT__29))) != 0) || ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (RFilterT__32 - 33)) | (1 << (RFilterT__33 - 33)) | (1 << (RFilterT__35 - 33)) | (1 << (RFilterT__37 - 33)) | (1 << (RFilterT__38 - 33)) | (1 << (RFilterT__39 - 33)) | (1 << (RFilterT__40 - 33)) | (1 << (RFilterT__41 - 33)) | (1 << (RFilterT__42 - 33)) | (1 << (RFilterT__43 - 33)) | (1 << (RFilterT__44 - 33)) | (1 << (RFilterT__45 - 33)) | (1 << (RFilterT__46 - 33)) | (1 << (RFilterT__47 - 33)) | (1 << (RFilterT__48 - 33)) | (1 << (RFilterT__49 - 33)) | (1 << (RFilterT__50 - 33)) | (1 << (RFilterT__51 - 33)) | (1 << (RFilterT__52 - 33)) | (1 << (RFilterT__53 - 33)) | (1 << (RFilterHEX - 33)) | (1 << (RFilterINT - 33)) | (1 << (RFilterFLOAT - 33)) | (1 << (RFilterCOMPLEX - 33)) | (1 << (RFilterSTRING - 33)) | (1 << (RFilterID - 33)) | (1 << (RFilterUSER_OP - 33)) | (1 << (RFilterNL - 33)))) != 0) {
			p.SetState(113)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RFilterT__1, RFilterT__3, RFilterT__4, RFilterT__5, RFilterT__6, RFilterT__7, RFilterT__8, RFilterT__9, RFilterT__10, RFilterT__11, RFilterT__12, RFilterT__13, RFilterT__14, RFilterT__15, RFilterT__16, RFilterT__17, RFilterT__18, RFilterT__19, RFilterT__20, RFilterT__21, RFilterT__22, RFilterT__23, RFilterT__24, RFilterT__25, RFilterT__26, RFilterT__28, RFilterT__29, RFilterT__32, RFilterT__33, RFilterT__35, RFilterT__37, RFilterT__38, RFilterT__39, RFilterT__40, RFilterT__41, RFilterT__42, RFilterT__43, RFilterT__44, RFilterT__45, RFilterT__46, RFilterT__47, RFilterT__48, RFilterT__49, RFilterT__50, RFilterT__51, RFilterT__52, RFilterT__53, RFilterHEX, RFilterINT, RFilterFLOAT, RFilterCOMPLEX, RFilterSTRING, RFilterID, RFilterUSER_OP:
				{
					p.SetState(111)
					p.Elem()
				}


			case RFilterNL:
				{
					p.SetState(112)
					p.Eat()
				}



			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(118)
			p.Match(RFilterT__34)
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(119)
				p.Eat()
			}


		}


	case RFilterT__37:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(122)
			p.Match(RFilterT__37)
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == RFilterNL {
			{
				p.SetState(123)
				p.Eat()
			}

		}
		{
			p.SetState(126)
			p.Match(RFilterT__33)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RFilterT__1) | (1 << RFilterT__3) | (1 << RFilterT__4) | (1 << RFilterT__5) | (1 << RFilterT__6) | (1 << RFilterT__7) | (1 << RFilterT__8) | (1 << RFilterT__9) | (1 << RFilterT__10) | (1 << RFilterT__11) | (1 << RFilterT__12) | (1 << RFilterT__13) | (1 << RFilterT__14) | (1 << RFilterT__15) | (1 << RFilterT__16) | (1 << RFilterT__17) | (1 << RFilterT__18) | (1 << RFilterT__19) | (1 << RFilterT__20) | (1 << RFilterT__21) | (1 << RFilterT__22) | (1 << RFilterT__23) | (1 << RFilterT__24) | (1 << RFilterT__25) | (1 << RFilterT__26) | (1 << RFilterT__28) | (1 << RFilterT__29))) != 0) || ((((_la - 33)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 33))) & ((1 << (RFilterT__32 - 33)) | (1 << (RFilterT__33 - 33)) | (1 << (RFilterT__35 - 33)) | (1 << (RFilterT__37 - 33)) | (1 << (RFilterT__38 - 33)) | (1 << (RFilterT__39 - 33)) | (1 << (RFilterT__40 - 33)) | (1 << (RFilterT__41 - 33)) | (1 << (RFilterT__42 - 33)) | (1 << (RFilterT__43 - 33)) | (1 << (RFilterT__44 - 33)) | (1 << (RFilterT__45 - 33)) | (1 << (RFilterT__46 - 33)) | (1 << (RFilterT__47 - 33)) | (1 << (RFilterT__48 - 33)) | (1 << (RFilterT__49 - 33)) | (1 << (RFilterT__50 - 33)) | (1 << (RFilterT__51 - 33)) | (1 << (RFilterT__52 - 33)) | (1 << (RFilterT__53 - 33)) | (1 << (RFilterHEX - 33)) | (1 << (RFilterINT - 33)) | (1 << (RFilterFLOAT - 33)) | (1 << (RFilterCOMPLEX - 33)) | (1 << (RFilterSTRING - 33)) | (1 << (RFilterID - 33)) | (1 << (RFilterUSER_OP - 33)) | (1 << (RFilterNL - 33)))) != 0) {
			p.SetState(129)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RFilterT__1, RFilterT__3, RFilterT__4, RFilterT__5, RFilterT__6, RFilterT__7, RFilterT__8, RFilterT__9, RFilterT__10, RFilterT__11, RFilterT__12, RFilterT__13, RFilterT__14, RFilterT__15, RFilterT__16, RFilterT__17, RFilterT__18, RFilterT__19, RFilterT__20, RFilterT__21, RFilterT__22, RFilterT__23, RFilterT__24, RFilterT__25, RFilterT__26, RFilterT__28, RFilterT__29, RFilterT__32, RFilterT__33, RFilterT__35, RFilterT__37, RFilterT__38, RFilterT__39, RFilterT__40, RFilterT__41, RFilterT__42, RFilterT__43, RFilterT__44, RFilterT__45, RFilterT__46, RFilterT__47, RFilterT__48, RFilterT__49, RFilterT__50, RFilterT__51, RFilterT__52, RFilterT__53, RFilterHEX, RFilterINT, RFilterFLOAT, RFilterCOMPLEX, RFilterSTRING, RFilterID, RFilterUSER_OP:
				{
					p.SetState(127)
					p.Elem()
				}


			case RFilterNL:
				{
					p.SetState(128)
					p.Eat()
				}



			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(133)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(134)
			p.Match(RFilterT__34)
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(135)
				p.Eat()
			}


		}


	case RFilterT__38:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(138)
			p.Match(RFilterT__38)
		}

		        // ``inside a compound expression, a newline before else is discarded,
		        // whereas at the outermost level, the newline terminates the if
		        // construction and a subsequent else causes a syntax error.''
		        /*
		        Works here
		            if (1==0) { print(1) } else { print(2) }

		        and correctly gets error here:

		            if (1==0) { print(1) }
		            else { print(2) }

		        this works too:

		            if (1==0) {
		              if (2==0) print(1)
		              else print(2)
		            }
		        */
		        WritableToken tok = (WritableToken)_input.LT(-2);
		        if (curlies>0&&tok.getType()==NL) tok.setChannel(Token.HIDDEN_CHANNEL);
		        



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
	p.RuleIndex = RFilterRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RFilterRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) ID() antlr.TerminalNode {
	return s.GetToken(RFilterID, 0)
}

func (s *AtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(RFilterSTRING, 0)
}

func (s *AtomContext) HEX() antlr.TerminalNode {
	return s.GetToken(RFilterHEX, 0)
}

func (s *AtomContext) INT() antlr.TerminalNode {
	return s.GetToken(RFilterINT, 0)
}

func (s *AtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(RFilterFLOAT, 0)
}

func (s *AtomContext) COMPLEX() antlr.TerminalNode {
	return s.GetToken(RFilterCOMPLEX, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RFilterListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RFilterListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RFilterVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RFilter) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RFilterRULE_atom)
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
	p.SetState(142)
	_la = p.GetTokenStream().LA(1)

	if !(((((_la - 45)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 45))) & ((1 << (RFilterT__44 - 45)) | (1 << (RFilterT__45 - 45)) | (1 << (RFilterT__46 - 45)) | (1 << (RFilterT__47 - 45)) | (1 << (RFilterT__48 - 45)) | (1 << (RFilterT__49 - 45)) | (1 << (RFilterT__50 - 45)) | (1 << (RFilterT__51 - 45)) | (1 << (RFilterHEX - 45)) | (1 << (RFilterINT - 45)) | (1 << (RFilterFLOAT - 45)) | (1 << (RFilterCOMPLEX - 45)) | (1 << (RFilterSTRING - 45)) | (1 << (RFilterID - 45)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RFilterRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RFilterRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) USER_OP() antlr.TerminalNode {
	return s.GetToken(RFilterUSER_OP, 0)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RFilterListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RFilterListener); ok {
		listenerT.ExitOp(s)
	}
}

func (s *OpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RFilterVisitor:
		return t.VisitOp(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *RFilter) Op() (localctx IOpContext) {
	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RFilterRULE_op)
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
	p.SetState(144)
	_la = p.GetTokenStream().LA(1)

	if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << RFilterT__4) | (1 << RFilterT__5) | (1 << RFilterT__6) | (1 << RFilterT__7) | (1 << RFilterT__8) | (1 << RFilterT__9) | (1 << RFilterT__10) | (1 << RFilterT__11) | (1 << RFilterT__12) | (1 << RFilterT__13) | (1 << RFilterT__14) | (1 << RFilterT__15) | (1 << RFilterT__16) | (1 << RFilterT__17) | (1 << RFilterT__18) | (1 << RFilterT__19) | (1 << RFilterT__20) | (1 << RFilterT__21) | (1 << RFilterT__22) | (1 << RFilterT__23) | (1 << RFilterT__24) | (1 << RFilterT__25) | (1 << RFilterT__26) | (1 << RFilterT__28) | (1 << RFilterT__29))) != 0) || ((((_la - 41)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 41))) & ((1 << (RFilterT__40 - 41)) | (1 << (RFilterT__42 - 41)) | (1 << (RFilterT__43 - 41)) | (1 << (RFilterT__52 - 41)) | (1 << (RFilterT__53 - 41)) | (1 << (RFilterUSER_OP - 41)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


