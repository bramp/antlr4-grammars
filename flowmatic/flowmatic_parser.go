// Code generated from flowmatic.g4 by ANTLR 4.9.3. DO NOT EDIT.

package flowmatic // flowmatic
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 207,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 7, 2, 56,
	10, 2, 12, 2, 14, 2, 59, 11, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 7, 3, 68, 10, 3, 12, 3, 14, 3, 71, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 95, 10, 5, 12, 5, 14, 5, 98, 11,
	5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 107, 10, 7, 13, 7, 14,
	7, 108, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8, 115, 10, 8, 13, 8, 14, 8, 116, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 189, 10, 21, 12, 21, 14,
	21, 192, 11, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 2, 2, 28, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 52, 2, 4, 4, 2, 31, 31, 35, 35, 4, 2, 28, 29, 33, 33, 2, 201, 2, 57,
	3, 2, 2, 2, 4, 63, 3, 2, 2, 2, 6, 74, 3, 2, 2, 2, 8, 96, 3, 2, 2, 2, 10,
	99, 3, 2, 2, 2, 12, 102, 3, 2, 2, 2, 14, 110, 3, 2, 2, 2, 16, 118, 3, 2,
	2, 2, 18, 122, 3, 2, 2, 2, 20, 125, 3, 2, 2, 2, 22, 128, 3, 2, 2, 2, 24,
	133, 3, 2, 2, 2, 26, 138, 3, 2, 2, 2, 28, 142, 3, 2, 2, 2, 30, 145, 3,
	2, 2, 2, 32, 147, 3, 2, 2, 2, 34, 158, 3, 2, 2, 2, 36, 169, 3, 2, 2, 2,
	38, 175, 3, 2, 2, 2, 40, 183, 3, 2, 2, 2, 42, 193, 3, 2, 2, 2, 44, 195,
	3, 2, 2, 2, 46, 197, 3, 2, 2, 2, 48, 199, 3, 2, 2, 2, 50, 201, 3, 2, 2,
	2, 52, 204, 3, 2, 2, 2, 54, 56, 5, 4, 3, 2, 55, 54, 3, 2, 2, 2, 56, 59,
	3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 60, 3, 2, 2, 2,
	59, 57, 3, 2, 2, 2, 60, 61, 7, 30, 2, 2, 61, 62, 7, 2, 2, 3, 62, 3, 3,
	2, 2, 2, 63, 64, 5, 6, 4, 2, 64, 69, 5, 8, 5, 2, 65, 66, 7, 3, 2, 2, 66,
	68, 5, 8, 5, 2, 67, 65, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2,
	2, 69, 70, 3, 2, 2, 2, 70, 72, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73,
	7, 4, 2, 2, 73, 5, 3, 2, 2, 2, 74, 75, 7, 5, 2, 2, 75, 76, 7, 35, 2, 2,
	76, 77, 7, 6, 2, 2, 77, 7, 3, 2, 2, 2, 78, 95, 5, 40, 21, 2, 79, 95, 5,
	38, 20, 2, 80, 95, 5, 36, 19, 2, 81, 95, 5, 34, 18, 2, 82, 95, 5, 32, 17,
	2, 83, 95, 5, 26, 14, 2, 84, 95, 5, 30, 16, 2, 85, 95, 5, 28, 15, 2, 86,
	95, 5, 22, 12, 2, 87, 95, 5, 20, 11, 2, 88, 95, 5, 16, 9, 2, 89, 95, 5,
	18, 10, 2, 90, 95, 5, 14, 8, 2, 91, 95, 5, 10, 6, 2, 92, 95, 5, 12, 7,
	2, 93, 95, 5, 24, 13, 2, 94, 78, 3, 2, 2, 2, 94, 79, 3, 2, 2, 2, 94, 80,
	3, 2, 2, 2, 94, 81, 3, 2, 2, 2, 94, 82, 3, 2, 2, 2, 94, 83, 3, 2, 2, 2,
	94, 84, 3, 2, 2, 2, 94, 85, 3, 2, 2, 2, 94, 86, 3, 2, 2, 2, 94, 87, 3,
	2, 2, 2, 94, 88, 3, 2, 2, 2, 94, 89, 3, 2, 2, 2, 94, 90, 3, 2, 2, 2, 94,
	91, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2,
	2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 9, 3, 2, 2, 2, 98, 96, 3,
	2, 2, 2, 99, 100, 7, 7, 2, 2, 100, 101, 5, 46, 24, 2, 101, 11, 3, 2, 2,
	2, 102, 106, 7, 8, 2, 2, 103, 104, 5, 44, 23, 2, 104, 105, 5, 46, 24, 2,
	105, 107, 3, 2, 2, 2, 106, 103, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 13, 3, 2, 2, 2, 110, 114, 7,
	9, 2, 2, 111, 112, 5, 44, 23, 2, 112, 113, 5, 46, 24, 2, 113, 115, 3, 2,
	2, 2, 114, 111, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2,
	116, 117, 3, 2, 2, 2, 117, 15, 3, 2, 2, 2, 118, 119, 7, 10, 2, 2, 119,
	120, 7, 11, 2, 2, 120, 121, 5, 50, 26, 2, 121, 17, 3, 2, 2, 2, 122, 123,
	7, 12, 2, 2, 123, 124, 5, 46, 24, 2, 124, 19, 3, 2, 2, 2, 125, 126, 7,
	13, 2, 2, 126, 127, 5, 46, 24, 2, 127, 21, 3, 2, 2, 2, 128, 129, 7, 14,
	2, 2, 129, 130, 5, 46, 24, 2, 130, 131, 7, 11, 2, 2, 131, 132, 5, 46, 24,
	2, 132, 23, 3, 2, 2, 2, 133, 134, 7, 15, 2, 2, 134, 135, 5, 52, 27, 2,
	135, 136, 7, 32, 2, 2, 136, 137, 5, 50, 26, 2, 137, 25, 3, 2, 2, 2, 138,
	139, 7, 16, 2, 2, 139, 140, 7, 32, 2, 2, 140, 141, 5, 50, 26, 2, 141, 27,
	3, 2, 2, 2, 142, 143, 7, 17, 2, 2, 143, 144, 5, 46, 24, 2, 144, 29, 3,
	2, 2, 2, 145, 146, 7, 18, 2, 2, 146, 31, 3, 2, 2, 2, 147, 148, 7, 19, 2,
	2, 148, 149, 5, 48, 25, 2, 149, 150, 7, 5, 2, 2, 150, 151, 5, 46, 24, 2,
	151, 152, 7, 6, 2, 2, 152, 153, 7, 20, 2, 2, 153, 154, 5, 48, 25, 2, 154,
	155, 7, 5, 2, 2, 155, 156, 5, 46, 24, 2, 156, 157, 7, 6, 2, 2, 157, 33,
	3, 2, 2, 2, 158, 159, 7, 21, 2, 2, 159, 160, 5, 48, 25, 2, 160, 161, 7,
	5, 2, 2, 161, 162, 5, 46, 24, 2, 162, 163, 7, 6, 2, 2, 163, 164, 7, 11,
	2, 2, 164, 165, 5, 48, 25, 2, 165, 166, 7, 5, 2, 2, 166, 167, 5, 46, 24,
	2, 167, 168, 7, 6, 2, 2, 168, 35, 3, 2, 2, 2, 169, 170, 7, 22, 2, 2, 170,
	171, 5, 50, 26, 2, 171, 172, 7, 11, 2, 2, 172, 173, 7, 32, 2, 2, 173, 174,
	5, 50, 26, 2, 174, 37, 3, 2, 2, 2, 175, 176, 7, 23, 2, 2, 176, 177, 5,
	48, 25, 2, 177, 178, 7, 5, 2, 2, 178, 179, 5, 46, 24, 2, 179, 180, 7, 6,
	2, 2, 180, 181, 7, 24, 2, 2, 181, 182, 5, 42, 22, 2, 182, 39, 3, 2, 2,
	2, 183, 184, 7, 25, 2, 2, 184, 185, 7, 26, 2, 2, 185, 190, 5, 46, 24, 2,
	186, 187, 7, 3, 2, 2, 187, 189, 5, 46, 24, 2, 188, 186, 3, 2, 2, 2, 189,
	192, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 41, 3,
	2, 2, 2, 192, 190, 3, 2, 2, 2, 193, 194, 9, 2, 2, 2, 194, 43, 3, 2, 2,
	2, 195, 196, 7, 34, 2, 2, 196, 45, 3, 2, 2, 2, 197, 198, 7, 34, 2, 2, 198,
	47, 3, 2, 2, 2, 199, 200, 7, 34, 2, 2, 200, 49, 3, 2, 2, 2, 201, 202, 7,
	27, 2, 2, 202, 203, 7, 35, 2, 2, 203, 51, 3, 2, 2, 2, 204, 205, 9, 3, 2,
	2, 205, 53, 3, 2, 2, 2, 9, 57, 69, 94, 96, 108, 116, 190,
}
var literalNames = []string{
	"", "';'", "'.'", "'('", "')'", "'HSP'", "'OUTPUT'", "'INPUT'", "'JUMP'",
	"'TO'", "'READ-ITEM'", "'WRITE-ITEM'", "'TRANSFER'", "'IF'", "'OTHERWISE'",
	"'REWIND'", "'STOP'", "'COMPARE'", "'WITH'", "'MOVE'", "'SET'", "'TEST'",
	"'AGAINST'", "'CLOSE-OUT'", "'FILES'", "'OPERATION'", "'EQUAL'", "'GREATER'",
	"", "", "'GO TO'", "'END OF DATA'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "END", "ZERO", "GOTO", "EOD", "ID",
	"NUM", "WS",
}

var ruleNames = []string{
	"flowmatic", "line", "label", "statement", "hsp_statement", "output_statement",
	"input_statement", "jumpto_statement", "readitem_statement", "writeitem_statement",
	"transfer_statement", "if_statement", "otherwise_statement", "rewind_statement",
	"stop_statement", "compare_statement", "move_statement", "set_statement",
	"test_statement", "closeout_statement", "num", "filename", "fileletter",
	"fieldname", "operation", "op",
}

type flowmaticParser struct {
	*antlr.BaseParser
}

// NewflowmaticParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *flowmaticParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewflowmaticParser(input antlr.TokenStream) *flowmaticParser {
	this := new(flowmaticParser)
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
	this.GrammarFileName = "flowmatic.g4"

	return this
}

// flowmaticParser tokens.
const (
	flowmaticParserEOF   = antlr.TokenEOF
	flowmaticParserT__0  = 1
	flowmaticParserT__1  = 2
	flowmaticParserT__2  = 3
	flowmaticParserT__3  = 4
	flowmaticParserT__4  = 5
	flowmaticParserT__5  = 6
	flowmaticParserT__6  = 7
	flowmaticParserT__7  = 8
	flowmaticParserT__8  = 9
	flowmaticParserT__9  = 10
	flowmaticParserT__10 = 11
	flowmaticParserT__11 = 12
	flowmaticParserT__12 = 13
	flowmaticParserT__13 = 14
	flowmaticParserT__14 = 15
	flowmaticParserT__15 = 16
	flowmaticParserT__16 = 17
	flowmaticParserT__17 = 18
	flowmaticParserT__18 = 19
	flowmaticParserT__19 = 20
	flowmaticParserT__20 = 21
	flowmaticParserT__21 = 22
	flowmaticParserT__22 = 23
	flowmaticParserT__23 = 24
	flowmaticParserT__24 = 25
	flowmaticParserT__25 = 26
	flowmaticParserT__26 = 27
	flowmaticParserEND   = 28
	flowmaticParserZERO  = 29
	flowmaticParserGOTO  = 30
	flowmaticParserEOD   = 31
	flowmaticParserID    = 32
	flowmaticParserNUM   = 33
	flowmaticParserWS    = 34
)

// flowmaticParser rules.
const (
	flowmaticParserRULE_flowmatic           = 0
	flowmaticParserRULE_line                = 1
	flowmaticParserRULE_label               = 2
	flowmaticParserRULE_statement           = 3
	flowmaticParserRULE_hsp_statement       = 4
	flowmaticParserRULE_output_statement    = 5
	flowmaticParserRULE_input_statement     = 6
	flowmaticParserRULE_jumpto_statement    = 7
	flowmaticParserRULE_readitem_statement  = 8
	flowmaticParserRULE_writeitem_statement = 9
	flowmaticParserRULE_transfer_statement  = 10
	flowmaticParserRULE_if_statement        = 11
	flowmaticParserRULE_otherwise_statement = 12
	flowmaticParserRULE_rewind_statement    = 13
	flowmaticParserRULE_stop_statement      = 14
	flowmaticParserRULE_compare_statement   = 15
	flowmaticParserRULE_move_statement      = 16
	flowmaticParserRULE_set_statement       = 17
	flowmaticParserRULE_test_statement      = 18
	flowmaticParserRULE_closeout_statement  = 19
	flowmaticParserRULE_num                 = 20
	flowmaticParserRULE_filename            = 21
	flowmaticParserRULE_fileletter          = 22
	flowmaticParserRULE_fieldname           = 23
	flowmaticParserRULE_operation           = 24
	flowmaticParserRULE_op                  = 25
)

// IFlowmaticContext is an interface to support dynamic dispatch.
type IFlowmaticContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlowmaticContext differentiates from other interfaces.
	IsFlowmaticContext()
}

type FlowmaticContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlowmaticContext() *FlowmaticContext {
	var p = new(FlowmaticContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_flowmatic
	return p
}

func (*FlowmaticContext) IsFlowmaticContext() {}

func NewFlowmaticContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlowmaticContext {
	var p = new(FlowmaticContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_flowmatic

	return p
}

func (s *FlowmaticContext) GetParser() antlr.Parser { return s.parser }

func (s *FlowmaticContext) END() antlr.TerminalNode {
	return s.GetToken(flowmaticParserEND, 0)
}

func (s *FlowmaticContext) EOF() antlr.TerminalNode {
	return s.GetToken(flowmaticParserEOF, 0)
}

func (s *FlowmaticContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *FlowmaticContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *FlowmaticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowmaticContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlowmaticContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterFlowmatic(s)
	}
}

func (s *FlowmaticContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitFlowmatic(s)
	}
}

func (p *flowmaticParser) Flowmatic() (localctx IFlowmaticContext) {
	this := p
	_ = this

	localctx = NewFlowmaticContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, flowmaticParserRULE_flowmatic)
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

	for _la == flowmaticParserT__2 {
		{
			p.SetState(52)
			p.Line()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(flowmaticParserEND)
	}
	{
		p.SetState(59)
		p.Match(flowmaticParserEOF)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LineContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *LineContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *flowmaticParser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, flowmaticParserRULE_line)
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
		p.SetState(61)
		p.Label()
	}
	{
		p.SetState(62)
		p.Statement()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == flowmaticParserT__0 {
		{
			p.SetState(63)
			p.Match(flowmaticParserT__0)
		}
		{
			p.SetState(64)
			p.Statement()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(flowmaticParserT__1)
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) NUM() antlr.TerminalNode {
	return s.GetToken(flowmaticParserNUM, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *flowmaticParser) Label() (localctx ILabelContext) {
	this := p
	_ = this

	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, flowmaticParserRULE_label)

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
		p.Match(flowmaticParserT__2)
	}
	{
		p.SetState(73)
		p.Match(flowmaticParserNUM)
	}
	{
		p.SetState(74)
		p.Match(flowmaticParserT__3)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) AllCloseout_statement() []ICloseout_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICloseout_statementContext)(nil)).Elem())
	var tst = make([]ICloseout_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICloseout_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Closeout_statement(i int) ICloseout_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICloseout_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICloseout_statementContext)
}

func (s *StatementContext) AllTest_statement() []ITest_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITest_statementContext)(nil)).Elem())
	var tst = make([]ITest_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITest_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Test_statement(i int) ITest_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITest_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITest_statementContext)
}

func (s *StatementContext) AllSet_statement() []ISet_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISet_statementContext)(nil)).Elem())
	var tst = make([]ISet_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISet_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Set_statement(i int) ISet_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISet_statementContext)
}

func (s *StatementContext) AllMove_statement() []IMove_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMove_statementContext)(nil)).Elem())
	var tst = make([]IMove_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMove_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Move_statement(i int) IMove_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMove_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMove_statementContext)
}

func (s *StatementContext) AllCompare_statement() []ICompare_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICompare_statementContext)(nil)).Elem())
	var tst = make([]ICompare_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICompare_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Compare_statement(i int) ICompare_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompare_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICompare_statementContext)
}

func (s *StatementContext) AllOtherwise_statement() []IOtherwise_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOtherwise_statementContext)(nil)).Elem())
	var tst = make([]IOtherwise_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOtherwise_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Otherwise_statement(i int) IOtherwise_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOtherwise_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOtherwise_statementContext)
}

func (s *StatementContext) AllStop_statement() []IStop_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStop_statementContext)(nil)).Elem())
	var tst = make([]IStop_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStop_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Stop_statement(i int) IStop_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStop_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStop_statementContext)
}

func (s *StatementContext) AllRewind_statement() []IRewind_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRewind_statementContext)(nil)).Elem())
	var tst = make([]IRewind_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRewind_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Rewind_statement(i int) IRewind_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRewind_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRewind_statementContext)
}

func (s *StatementContext) AllTransfer_statement() []ITransfer_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITransfer_statementContext)(nil)).Elem())
	var tst = make([]ITransfer_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITransfer_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Transfer_statement(i int) ITransfer_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransfer_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITransfer_statementContext)
}

func (s *StatementContext) AllWriteitem_statement() []IWriteitem_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWriteitem_statementContext)(nil)).Elem())
	var tst = make([]IWriteitem_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWriteitem_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Writeitem_statement(i int) IWriteitem_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWriteitem_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWriteitem_statementContext)
}

func (s *StatementContext) AllJumpto_statement() []IJumpto_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJumpto_statementContext)(nil)).Elem())
	var tst = make([]IJumpto_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJumpto_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Jumpto_statement(i int) IJumpto_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJumpto_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJumpto_statementContext)
}

func (s *StatementContext) AllReaditem_statement() []IReaditem_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReaditem_statementContext)(nil)).Elem())
	var tst = make([]IReaditem_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReaditem_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Readitem_statement(i int) IReaditem_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReaditem_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReaditem_statementContext)
}

func (s *StatementContext) AllInput_statement() []IInput_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInput_statementContext)(nil)).Elem())
	var tst = make([]IInput_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInput_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Input_statement(i int) IInput_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInput_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInput_statementContext)
}

func (s *StatementContext) AllHsp_statement() []IHsp_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHsp_statementContext)(nil)).Elem())
	var tst = make([]IHsp_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHsp_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Hsp_statement(i int) IHsp_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHsp_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHsp_statementContext)
}

func (s *StatementContext) AllOutput_statement() []IOutput_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOutput_statementContext)(nil)).Elem())
	var tst = make([]IOutput_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOutput_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) Output_statement(i int) IOutput_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutput_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOutput_statementContext)
}

func (s *StatementContext) AllIf_statement() []IIf_statementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIf_statementContext)(nil)).Elem())
	var tst = make([]IIf_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIf_statementContext)
		}
	}

	return tst
}

func (s *StatementContext) If_statement(i int) IIf_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *flowmaticParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, flowmaticParserRULE_statement)
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

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<flowmaticParserT__4)|(1<<flowmaticParserT__5)|(1<<flowmaticParserT__6)|(1<<flowmaticParserT__7)|(1<<flowmaticParserT__9)|(1<<flowmaticParserT__10)|(1<<flowmaticParserT__11)|(1<<flowmaticParserT__12)|(1<<flowmaticParserT__13)|(1<<flowmaticParserT__14)|(1<<flowmaticParserT__15)|(1<<flowmaticParserT__16)|(1<<flowmaticParserT__18)|(1<<flowmaticParserT__19)|(1<<flowmaticParserT__20)|(1<<flowmaticParserT__22))) != 0 {
		p.SetState(92)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case flowmaticParserT__22:
			{
				p.SetState(76)
				p.Closeout_statement()
			}

		case flowmaticParserT__20:
			{
				p.SetState(77)
				p.Test_statement()
			}

		case flowmaticParserT__19:
			{
				p.SetState(78)
				p.Set_statement()
			}

		case flowmaticParserT__18:
			{
				p.SetState(79)
				p.Move_statement()
			}

		case flowmaticParserT__16:
			{
				p.SetState(80)
				p.Compare_statement()
			}

		case flowmaticParserT__13:
			{
				p.SetState(81)
				p.Otherwise_statement()
			}

		case flowmaticParserT__15:
			{
				p.SetState(82)
				p.Stop_statement()
			}

		case flowmaticParserT__14:
			{
				p.SetState(83)
				p.Rewind_statement()
			}

		case flowmaticParserT__11:
			{
				p.SetState(84)
				p.Transfer_statement()
			}

		case flowmaticParserT__10:
			{
				p.SetState(85)
				p.Writeitem_statement()
			}

		case flowmaticParserT__7:
			{
				p.SetState(86)
				p.Jumpto_statement()
			}

		case flowmaticParserT__9:
			{
				p.SetState(87)
				p.Readitem_statement()
			}

		case flowmaticParserT__6:
			{
				p.SetState(88)
				p.Input_statement()
			}

		case flowmaticParserT__4:
			{
				p.SetState(89)
				p.Hsp_statement()
			}

		case flowmaticParserT__5:
			{
				p.SetState(90)
				p.Output_statement()
			}

		case flowmaticParserT__12:
			{
				p.SetState(91)
				p.If_statement()
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

// IHsp_statementContext is an interface to support dynamic dispatch.
type IHsp_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHsp_statementContext differentiates from other interfaces.
	IsHsp_statementContext()
}

type Hsp_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHsp_statementContext() *Hsp_statementContext {
	var p = new(Hsp_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_hsp_statement
	return p
}

func (*Hsp_statementContext) IsHsp_statementContext() {}

func NewHsp_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Hsp_statementContext {
	var p = new(Hsp_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_hsp_statement

	return p
}

func (s *Hsp_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Hsp_statementContext) Fileletter() IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Hsp_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Hsp_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Hsp_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterHsp_statement(s)
	}
}

func (s *Hsp_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitHsp_statement(s)
	}
}

func (p *flowmaticParser) Hsp_statement() (localctx IHsp_statementContext) {
	this := p
	_ = this

	localctx = NewHsp_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, flowmaticParserRULE_hsp_statement)

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
		p.Match(flowmaticParserT__4)
	}
	{
		p.SetState(98)
		p.Fileletter()
	}

	return localctx
}

// IOutput_statementContext is an interface to support dynamic dispatch.
type IOutput_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutput_statementContext differentiates from other interfaces.
	IsOutput_statementContext()
}

type Output_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutput_statementContext() *Output_statementContext {
	var p = new(Output_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_output_statement
	return p
}

func (*Output_statementContext) IsOutput_statementContext() {}

func NewOutput_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Output_statementContext {
	var p = new(Output_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_output_statement

	return p
}

func (s *Output_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Output_statementContext) AllFilename() []IFilenameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilenameContext)(nil)).Elem())
	var tst = make([]IFilenameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilenameContext)
		}
	}

	return tst
}

func (s *Output_statementContext) Filename(i int) IFilenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilenameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilenameContext)
}

func (s *Output_statementContext) AllFileletter() []IFileletterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFileletterContext)(nil)).Elem())
	var tst = make([]IFileletterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFileletterContext)
		}
	}

	return tst
}

func (s *Output_statementContext) Fileletter(i int) IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Output_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Output_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Output_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterOutput_statement(s)
	}
}

func (s *Output_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitOutput_statement(s)
	}
}

func (p *flowmaticParser) Output_statement() (localctx IOutput_statementContext) {
	this := p
	_ = this

	localctx = NewOutput_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, flowmaticParserRULE_output_statement)
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
		p.SetState(100)
		p.Match(flowmaticParserT__5)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == flowmaticParserID {
		{
			p.SetState(101)
			p.Filename()
		}
		{
			p.SetState(102)
			p.Fileletter()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInput_statementContext is an interface to support dynamic dispatch.
type IInput_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInput_statementContext differentiates from other interfaces.
	IsInput_statementContext()
}

type Input_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInput_statementContext() *Input_statementContext {
	var p = new(Input_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_input_statement
	return p
}

func (*Input_statementContext) IsInput_statementContext() {}

func NewInput_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Input_statementContext {
	var p = new(Input_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_input_statement

	return p
}

func (s *Input_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Input_statementContext) AllFilename() []IFilenameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilenameContext)(nil)).Elem())
	var tst = make([]IFilenameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilenameContext)
		}
	}

	return tst
}

func (s *Input_statementContext) Filename(i int) IFilenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilenameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilenameContext)
}

func (s *Input_statementContext) AllFileletter() []IFileletterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFileletterContext)(nil)).Elem())
	var tst = make([]IFileletterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFileletterContext)
		}
	}

	return tst
}

func (s *Input_statementContext) Fileletter(i int) IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Input_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Input_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Input_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterInput_statement(s)
	}
}

func (s *Input_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitInput_statement(s)
	}
}

func (p *flowmaticParser) Input_statement() (localctx IInput_statementContext) {
	this := p
	_ = this

	localctx = NewInput_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, flowmaticParserRULE_input_statement)
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
		p.SetState(108)
		p.Match(flowmaticParserT__6)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == flowmaticParserID {
		{
			p.SetState(109)
			p.Filename()
		}
		{
			p.SetState(110)
			p.Fileletter()
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IJumpto_statementContext is an interface to support dynamic dispatch.
type IJumpto_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJumpto_statementContext differentiates from other interfaces.
	IsJumpto_statementContext()
}

type Jumpto_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJumpto_statementContext() *Jumpto_statementContext {
	var p = new(Jumpto_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_jumpto_statement
	return p
}

func (*Jumpto_statementContext) IsJumpto_statementContext() {}

func NewJumpto_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Jumpto_statementContext {
	var p = new(Jumpto_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_jumpto_statement

	return p
}

func (s *Jumpto_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Jumpto_statementContext) Operation() IOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *Jumpto_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Jumpto_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Jumpto_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterJumpto_statement(s)
	}
}

func (s *Jumpto_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitJumpto_statement(s)
	}
}

func (p *flowmaticParser) Jumpto_statement() (localctx IJumpto_statementContext) {
	this := p
	_ = this

	localctx = NewJumpto_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, flowmaticParserRULE_jumpto_statement)

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
		p.SetState(116)
		p.Match(flowmaticParserT__7)
	}
	{
		p.SetState(117)
		p.Match(flowmaticParserT__8)
	}
	{
		p.SetState(118)
		p.Operation()
	}

	return localctx
}

// IReaditem_statementContext is an interface to support dynamic dispatch.
type IReaditem_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReaditem_statementContext differentiates from other interfaces.
	IsReaditem_statementContext()
}

type Readitem_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReaditem_statementContext() *Readitem_statementContext {
	var p = new(Readitem_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_readitem_statement
	return p
}

func (*Readitem_statementContext) IsReaditem_statementContext() {}

func NewReaditem_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Readitem_statementContext {
	var p = new(Readitem_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_readitem_statement

	return p
}

func (s *Readitem_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Readitem_statementContext) Fileletter() IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Readitem_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Readitem_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Readitem_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterReaditem_statement(s)
	}
}

func (s *Readitem_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitReaditem_statement(s)
	}
}

func (p *flowmaticParser) Readitem_statement() (localctx IReaditem_statementContext) {
	this := p
	_ = this

	localctx = NewReaditem_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, flowmaticParserRULE_readitem_statement)

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
		p.SetState(120)
		p.Match(flowmaticParserT__9)
	}
	{
		p.SetState(121)
		p.Fileletter()
	}

	return localctx
}

// IWriteitem_statementContext is an interface to support dynamic dispatch.
type IWriteitem_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWriteitem_statementContext differentiates from other interfaces.
	IsWriteitem_statementContext()
}

type Writeitem_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWriteitem_statementContext() *Writeitem_statementContext {
	var p = new(Writeitem_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_writeitem_statement
	return p
}

func (*Writeitem_statementContext) IsWriteitem_statementContext() {}

func NewWriteitem_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Writeitem_statementContext {
	var p = new(Writeitem_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_writeitem_statement

	return p
}

func (s *Writeitem_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Writeitem_statementContext) Fileletter() IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Writeitem_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Writeitem_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Writeitem_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterWriteitem_statement(s)
	}
}

func (s *Writeitem_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitWriteitem_statement(s)
	}
}

func (p *flowmaticParser) Writeitem_statement() (localctx IWriteitem_statementContext) {
	this := p
	_ = this

	localctx = NewWriteitem_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, flowmaticParserRULE_writeitem_statement)

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
		p.SetState(123)
		p.Match(flowmaticParserT__10)
	}
	{
		p.SetState(124)
		p.Fileletter()
	}

	return localctx
}

// ITransfer_statementContext is an interface to support dynamic dispatch.
type ITransfer_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransfer_statementContext differentiates from other interfaces.
	IsTransfer_statementContext()
}

type Transfer_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransfer_statementContext() *Transfer_statementContext {
	var p = new(Transfer_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_transfer_statement
	return p
}

func (*Transfer_statementContext) IsTransfer_statementContext() {}

func NewTransfer_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Transfer_statementContext {
	var p = new(Transfer_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_transfer_statement

	return p
}

func (s *Transfer_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Transfer_statementContext) AllFileletter() []IFileletterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFileletterContext)(nil)).Elem())
	var tst = make([]IFileletterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFileletterContext)
		}
	}

	return tst
}

func (s *Transfer_statementContext) Fileletter(i int) IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Transfer_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Transfer_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Transfer_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterTransfer_statement(s)
	}
}

func (s *Transfer_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitTransfer_statement(s)
	}
}

func (p *flowmaticParser) Transfer_statement() (localctx ITransfer_statementContext) {
	this := p
	_ = this

	localctx = NewTransfer_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, flowmaticParserRULE_transfer_statement)

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
		p.SetState(126)
		p.Match(flowmaticParserT__11)
	}
	{
		p.SetState(127)
		p.Fileletter()
	}
	{
		p.SetState(128)
		p.Match(flowmaticParserT__8)
	}
	{
		p.SetState(129)
		p.Fileletter()
	}

	return localctx
}

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_if_statement
	return p
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) Op() IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *If_statementContext) GOTO() antlr.TerminalNode {
	return s.GetToken(flowmaticParserGOTO, 0)
}

func (s *If_statementContext) Operation() IOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (p *flowmaticParser) If_statement() (localctx IIf_statementContext) {
	this := p
	_ = this

	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, flowmaticParserRULE_if_statement)

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
		p.SetState(131)
		p.Match(flowmaticParserT__12)
	}
	{
		p.SetState(132)
		p.Op()
	}
	{
		p.SetState(133)
		p.Match(flowmaticParserGOTO)
	}
	{
		p.SetState(134)
		p.Operation()
	}

	return localctx
}

// IOtherwise_statementContext is an interface to support dynamic dispatch.
type IOtherwise_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOtherwise_statementContext differentiates from other interfaces.
	IsOtherwise_statementContext()
}

type Otherwise_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOtherwise_statementContext() *Otherwise_statementContext {
	var p = new(Otherwise_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_otherwise_statement
	return p
}

func (*Otherwise_statementContext) IsOtherwise_statementContext() {}

func NewOtherwise_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Otherwise_statementContext {
	var p = new(Otherwise_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_otherwise_statement

	return p
}

func (s *Otherwise_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Otherwise_statementContext) GOTO() antlr.TerminalNode {
	return s.GetToken(flowmaticParserGOTO, 0)
}

func (s *Otherwise_statementContext) Operation() IOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *Otherwise_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Otherwise_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Otherwise_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterOtherwise_statement(s)
	}
}

func (s *Otherwise_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitOtherwise_statement(s)
	}
}

func (p *flowmaticParser) Otherwise_statement() (localctx IOtherwise_statementContext) {
	this := p
	_ = this

	localctx = NewOtherwise_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, flowmaticParserRULE_otherwise_statement)

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
		p.SetState(136)
		p.Match(flowmaticParserT__13)
	}
	{
		p.SetState(137)
		p.Match(flowmaticParserGOTO)
	}
	{
		p.SetState(138)
		p.Operation()
	}

	return localctx
}

// IRewind_statementContext is an interface to support dynamic dispatch.
type IRewind_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRewind_statementContext differentiates from other interfaces.
	IsRewind_statementContext()
}

type Rewind_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRewind_statementContext() *Rewind_statementContext {
	var p = new(Rewind_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_rewind_statement
	return p
}

func (*Rewind_statementContext) IsRewind_statementContext() {}

func NewRewind_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rewind_statementContext {
	var p = new(Rewind_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_rewind_statement

	return p
}

func (s *Rewind_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Rewind_statementContext) Fileletter() IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Rewind_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rewind_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rewind_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterRewind_statement(s)
	}
}

func (s *Rewind_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitRewind_statement(s)
	}
}

func (p *flowmaticParser) Rewind_statement() (localctx IRewind_statementContext) {
	this := p
	_ = this

	localctx = NewRewind_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, flowmaticParserRULE_rewind_statement)

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
		p.SetState(140)
		p.Match(flowmaticParserT__14)
	}
	{
		p.SetState(141)
		p.Fileletter()
	}

	return localctx
}

// IStop_statementContext is an interface to support dynamic dispatch.
type IStop_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStop_statementContext differentiates from other interfaces.
	IsStop_statementContext()
}

type Stop_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStop_statementContext() *Stop_statementContext {
	var p = new(Stop_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_stop_statement
	return p
}

func (*Stop_statementContext) IsStop_statementContext() {}

func NewStop_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stop_statementContext {
	var p = new(Stop_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_stop_statement

	return p
}

func (s *Stop_statementContext) GetParser() antlr.Parser { return s.parser }
func (s *Stop_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stop_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stop_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterStop_statement(s)
	}
}

func (s *Stop_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitStop_statement(s)
	}
}

func (p *flowmaticParser) Stop_statement() (localctx IStop_statementContext) {
	this := p
	_ = this

	localctx = NewStop_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, flowmaticParserRULE_stop_statement)

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
		p.SetState(143)
		p.Match(flowmaticParserT__15)
	}

	return localctx
}

// ICompare_statementContext is an interface to support dynamic dispatch.
type ICompare_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompare_statementContext differentiates from other interfaces.
	IsCompare_statementContext()
}

type Compare_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompare_statementContext() *Compare_statementContext {
	var p = new(Compare_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_compare_statement
	return p
}

func (*Compare_statementContext) IsCompare_statementContext() {}

func NewCompare_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compare_statementContext {
	var p = new(Compare_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_compare_statement

	return p
}

func (s *Compare_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Compare_statementContext) AllFieldname() []IFieldnameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldnameContext)(nil)).Elem())
	var tst = make([]IFieldnameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldnameContext)
		}
	}

	return tst
}

func (s *Compare_statementContext) Fieldname(i int) IFieldnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldnameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldnameContext)
}

func (s *Compare_statementContext) AllFileletter() []IFileletterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFileletterContext)(nil)).Elem())
	var tst = make([]IFileletterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFileletterContext)
		}
	}

	return tst
}

func (s *Compare_statementContext) Fileletter(i int) IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Compare_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compare_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Compare_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterCompare_statement(s)
	}
}

func (s *Compare_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitCompare_statement(s)
	}
}

func (p *flowmaticParser) Compare_statement() (localctx ICompare_statementContext) {
	this := p
	_ = this

	localctx = NewCompare_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, flowmaticParserRULE_compare_statement)

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
		p.Match(flowmaticParserT__16)
	}
	{
		p.SetState(146)
		p.Fieldname()
	}
	{
		p.SetState(147)
		p.Match(flowmaticParserT__2)
	}
	{
		p.SetState(148)
		p.Fileletter()
	}
	{
		p.SetState(149)
		p.Match(flowmaticParserT__3)
	}
	{
		p.SetState(150)
		p.Match(flowmaticParserT__17)
	}
	{
		p.SetState(151)
		p.Fieldname()
	}
	{
		p.SetState(152)
		p.Match(flowmaticParserT__2)
	}
	{
		p.SetState(153)
		p.Fileletter()
	}
	{
		p.SetState(154)
		p.Match(flowmaticParserT__3)
	}

	return localctx
}

// IMove_statementContext is an interface to support dynamic dispatch.
type IMove_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMove_statementContext differentiates from other interfaces.
	IsMove_statementContext()
}

type Move_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMove_statementContext() *Move_statementContext {
	var p = new(Move_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_move_statement
	return p
}

func (*Move_statementContext) IsMove_statementContext() {}

func NewMove_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Move_statementContext {
	var p = new(Move_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_move_statement

	return p
}

func (s *Move_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Move_statementContext) AllFieldname() []IFieldnameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldnameContext)(nil)).Elem())
	var tst = make([]IFieldnameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldnameContext)
		}
	}

	return tst
}

func (s *Move_statementContext) Fieldname(i int) IFieldnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldnameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldnameContext)
}

func (s *Move_statementContext) AllFileletter() []IFileletterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFileletterContext)(nil)).Elem())
	var tst = make([]IFileletterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFileletterContext)
		}
	}

	return tst
}

func (s *Move_statementContext) Fileletter(i int) IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Move_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Move_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Move_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterMove_statement(s)
	}
}

func (s *Move_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitMove_statement(s)
	}
}

func (p *flowmaticParser) Move_statement() (localctx IMove_statementContext) {
	this := p
	_ = this

	localctx = NewMove_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, flowmaticParserRULE_move_statement)

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
		p.SetState(156)
		p.Match(flowmaticParserT__18)
	}
	{
		p.SetState(157)
		p.Fieldname()
	}
	{
		p.SetState(158)
		p.Match(flowmaticParserT__2)
	}
	{
		p.SetState(159)
		p.Fileletter()
	}
	{
		p.SetState(160)
		p.Match(flowmaticParserT__3)
	}
	{
		p.SetState(161)
		p.Match(flowmaticParserT__8)
	}
	{
		p.SetState(162)
		p.Fieldname()
	}
	{
		p.SetState(163)
		p.Match(flowmaticParserT__2)
	}
	{
		p.SetState(164)
		p.Fileletter()
	}
	{
		p.SetState(165)
		p.Match(flowmaticParserT__3)
	}

	return localctx
}

// ISet_statementContext is an interface to support dynamic dispatch.
type ISet_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_statementContext differentiates from other interfaces.
	IsSet_statementContext()
}

type Set_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_statementContext() *Set_statementContext {
	var p = new(Set_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_set_statement
	return p
}

func (*Set_statementContext) IsSet_statementContext() {}

func NewSet_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_statementContext {
	var p = new(Set_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_set_statement

	return p
}

func (s *Set_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_statementContext) AllOperation() []IOperationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperationContext)(nil)).Elem())
	var tst = make([]IOperationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperationContext)
		}
	}

	return tst
}

func (s *Set_statementContext) Operation(i int) IOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *Set_statementContext) GOTO() antlr.TerminalNode {
	return s.GetToken(flowmaticParserGOTO, 0)
}

func (s *Set_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterSet_statement(s)
	}
}

func (s *Set_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitSet_statement(s)
	}
}

func (p *flowmaticParser) Set_statement() (localctx ISet_statementContext) {
	this := p
	_ = this

	localctx = NewSet_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, flowmaticParserRULE_set_statement)

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
		p.SetState(167)
		p.Match(flowmaticParserT__19)
	}
	{
		p.SetState(168)
		p.Operation()
	}
	{
		p.SetState(169)
		p.Match(flowmaticParserT__8)
	}
	{
		p.SetState(170)
		p.Match(flowmaticParserGOTO)
	}
	{
		p.SetState(171)
		p.Operation()
	}

	return localctx
}

// ITest_statementContext is an interface to support dynamic dispatch.
type ITest_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTest_statementContext differentiates from other interfaces.
	IsTest_statementContext()
}

type Test_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTest_statementContext() *Test_statementContext {
	var p = new(Test_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_test_statement
	return p
}

func (*Test_statementContext) IsTest_statementContext() {}

func NewTest_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Test_statementContext {
	var p = new(Test_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_test_statement

	return p
}

func (s *Test_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Test_statementContext) Fieldname() IFieldnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldnameContext)
}

func (s *Test_statementContext) Fileletter() IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Test_statementContext) Num() INumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *Test_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Test_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Test_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterTest_statement(s)
	}
}

func (s *Test_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitTest_statement(s)
	}
}

func (p *flowmaticParser) Test_statement() (localctx ITest_statementContext) {
	this := p
	_ = this

	localctx = NewTest_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, flowmaticParserRULE_test_statement)

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
		p.SetState(173)
		p.Match(flowmaticParserT__20)
	}
	{
		p.SetState(174)
		p.Fieldname()
	}
	{
		p.SetState(175)
		p.Match(flowmaticParserT__2)
	}
	{
		p.SetState(176)
		p.Fileletter()
	}
	{
		p.SetState(177)
		p.Match(flowmaticParserT__3)
	}
	{
		p.SetState(178)
		p.Match(flowmaticParserT__21)
	}
	{
		p.SetState(179)
		p.Num()
	}

	return localctx
}

// ICloseout_statementContext is an interface to support dynamic dispatch.
type ICloseout_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCloseout_statementContext differentiates from other interfaces.
	IsCloseout_statementContext()
}

type Closeout_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCloseout_statementContext() *Closeout_statementContext {
	var p = new(Closeout_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_closeout_statement
	return p
}

func (*Closeout_statementContext) IsCloseout_statementContext() {}

func NewCloseout_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Closeout_statementContext {
	var p = new(Closeout_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_closeout_statement

	return p
}

func (s *Closeout_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Closeout_statementContext) AllFileletter() []IFileletterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFileletterContext)(nil)).Elem())
	var tst = make([]IFileletterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFileletterContext)
		}
	}

	return tst
}

func (s *Closeout_statementContext) Fileletter(i int) IFileletterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileletterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFileletterContext)
}

func (s *Closeout_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Closeout_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Closeout_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterCloseout_statement(s)
	}
}

func (s *Closeout_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitCloseout_statement(s)
	}
}

func (p *flowmaticParser) Closeout_statement() (localctx ICloseout_statementContext) {
	this := p
	_ = this

	localctx = NewCloseout_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, flowmaticParserRULE_closeout_statement)

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
		p.SetState(181)
		p.Match(flowmaticParserT__22)
	}
	{
		p.SetState(182)
		p.Match(flowmaticParserT__23)
	}
	{
		p.SetState(183)
		p.Fileletter()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(184)
				p.Match(flowmaticParserT__0)
			}
			{
				p.SetState(185)
				p.Fileletter()
			}

		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_num
	return p
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) NUM() antlr.TerminalNode {
	return s.GetToken(flowmaticParserNUM, 0)
}

func (s *NumContext) ZERO() antlr.TerminalNode {
	return s.GetToken(flowmaticParserZERO, 0)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitNum(s)
	}
}

func (p *flowmaticParser) Num() (localctx INumContext) {
	this := p
	_ = this

	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, flowmaticParserRULE_num)
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
		p.SetState(191)
		_la = p.GetTokenStream().LA(1)

		if !(_la == flowmaticParserZERO || _la == flowmaticParserNUM) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFilenameContext is an interface to support dynamic dispatch.
type IFilenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilenameContext differentiates from other interfaces.
	IsFilenameContext()
}

type FilenameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilenameContext() *FilenameContext {
	var p = new(FilenameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_filename
	return p
}

func (*FilenameContext) IsFilenameContext() {}

func NewFilenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilenameContext {
	var p = new(FilenameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_filename

	return p
}

func (s *FilenameContext) GetParser() antlr.Parser { return s.parser }

func (s *FilenameContext) ID() antlr.TerminalNode {
	return s.GetToken(flowmaticParserID, 0)
}

func (s *FilenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilenameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterFilename(s)
	}
}

func (s *FilenameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitFilename(s)
	}
}

func (p *flowmaticParser) Filename() (localctx IFilenameContext) {
	this := p
	_ = this

	localctx = NewFilenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, flowmaticParserRULE_filename)

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
		p.SetState(193)
		p.Match(flowmaticParserID)
	}

	return localctx
}

// IFileletterContext is an interface to support dynamic dispatch.
type IFileletterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileletterContext differentiates from other interfaces.
	IsFileletterContext()
}

type FileletterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileletterContext() *FileletterContext {
	var p = new(FileletterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_fileletter
	return p
}

func (*FileletterContext) IsFileletterContext() {}

func NewFileletterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileletterContext {
	var p = new(FileletterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_fileletter

	return p
}

func (s *FileletterContext) GetParser() antlr.Parser { return s.parser }

func (s *FileletterContext) ID() antlr.TerminalNode {
	return s.GetToken(flowmaticParserID, 0)
}

func (s *FileletterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileletterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileletterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterFileletter(s)
	}
}

func (s *FileletterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitFileletter(s)
	}
}

func (p *flowmaticParser) Fileletter() (localctx IFileletterContext) {
	this := p
	_ = this

	localctx = NewFileletterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, flowmaticParserRULE_fileletter)

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
		p.SetState(195)
		p.Match(flowmaticParserID)
	}

	return localctx
}

// IFieldnameContext is an interface to support dynamic dispatch.
type IFieldnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldnameContext differentiates from other interfaces.
	IsFieldnameContext()
}

type FieldnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldnameContext() *FieldnameContext {
	var p = new(FieldnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_fieldname
	return p
}

func (*FieldnameContext) IsFieldnameContext() {}

func NewFieldnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldnameContext {
	var p = new(FieldnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_fieldname

	return p
}

func (s *FieldnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldnameContext) ID() antlr.TerminalNode {
	return s.GetToken(flowmaticParserID, 0)
}

func (s *FieldnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterFieldname(s)
	}
}

func (s *FieldnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitFieldname(s)
	}
}

func (p *flowmaticParser) Fieldname() (localctx IFieldnameContext) {
	this := p
	_ = this

	localctx = NewFieldnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, flowmaticParserRULE_fieldname)

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
		p.SetState(197)
		p.Match(flowmaticParserID)
	}

	return localctx
}

// IOperationContext is an interface to support dynamic dispatch.
type IOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationContext differentiates from other interfaces.
	IsOperationContext()
}

type OperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationContext() *OperationContext {
	var p = new(OperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = flowmaticParserRULE_operation
	return p
}

func (*OperationContext) IsOperationContext() {}

func NewOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationContext {
	var p = new(OperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_operation

	return p
}

func (s *OperationContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationContext) NUM() antlr.TerminalNode {
	return s.GetToken(flowmaticParserNUM, 0)
}

func (s *OperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterOperation(s)
	}
}

func (s *OperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitOperation(s)
	}
}

func (p *flowmaticParser) Operation() (localctx IOperationContext) {
	this := p
	_ = this

	localctx = NewOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, flowmaticParserRULE_operation)

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
		p.SetState(199)
		p.Match(flowmaticParserT__24)
	}
	{
		p.SetState(200)
		p.Match(flowmaticParserNUM)
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
	p.RuleIndex = flowmaticParserRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = flowmaticParserRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) EOD() antlr.TerminalNode {
	return s.GetToken(flowmaticParserEOD, 0)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(flowmaticListener); ok {
		listenerT.ExitOp(s)
	}
}

func (p *flowmaticParser) Op() (localctx IOpContext) {
	this := p
	_ = this

	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, flowmaticParserRULE_op)
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
		p.SetState(202)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<flowmaticParserT__25)|(1<<flowmaticParserT__26)|(1<<flowmaticParserEOD))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
