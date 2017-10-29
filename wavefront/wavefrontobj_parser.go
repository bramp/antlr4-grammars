// Generated from WavefrontOBJ.g4 by ANTLR 4.7.

package wavefront // WavefrontOBJ
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 47, 65, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 6, 2, 20, 10, 2, 13, 2, 14, 2, 21, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 3, 5, 3, 33, 10, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 6, 4, 39, 10, 4, 13, 4, 14, 4, 40, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 50, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6,
	8, 10, 12, 14, 16, 2, 4, 3, 2, 8, 10, 3, 2, 6, 7, 2, 65, 2, 19, 3, 2, 2,
	2, 4, 29, 3, 2, 2, 2, 6, 36, 3, 2, 2, 2, 8, 49, 3, 2, 2, 2, 10, 51, 3,
	2, 2, 2, 12, 55, 3, 2, 2, 2, 14, 58, 3, 2, 2, 2, 16, 61, 3, 2, 2, 2, 18,
	20, 5, 4, 3, 2, 19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 19, 3, 2, 2,
	2, 21, 22, 3, 2, 2, 2, 22, 3, 3, 2, 2, 2, 23, 30, 5, 8, 5, 2, 24, 30, 5,
	6, 4, 2, 25, 30, 5, 10, 6, 2, 26, 30, 5, 12, 7, 2, 27, 30, 5, 14, 8, 2,
	28, 30, 5, 16, 9, 2, 29, 23, 3, 2, 2, 2, 29, 24, 3, 2, 2, 2, 29, 25, 3,
	2, 2, 2, 29, 26, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 28, 3, 2, 2, 2, 30,
	32, 3, 2, 2, 2, 31, 33, 7, 3, 2, 2, 32, 31, 3, 2, 2, 2, 32, 33, 3, 2, 2,
	2, 33, 34, 3, 2, 2, 2, 34, 35, 7, 4, 2, 2, 35, 5, 3, 2, 2, 2, 36, 38, 7,
	18, 2, 2, 37, 39, 7, 43, 2, 2, 38, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2,
	40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 7, 3, 2, 2, 2, 42, 43, 9, 2,
	2, 2, 43, 44, 7, 44, 2, 2, 44, 45, 7, 44, 2, 2, 45, 50, 7, 44, 2, 2, 46,
	47, 7, 11, 2, 2, 47, 48, 7, 44, 2, 2, 48, 50, 7, 44, 2, 2, 49, 42, 3, 2,
	2, 2, 49, 46, 3, 2, 2, 2, 50, 9, 3, 2, 2, 2, 51, 52, 7, 38, 2, 2, 52, 53,
	7, 46, 2, 2, 53, 54, 7, 5, 2, 2, 54, 11, 3, 2, 2, 2, 55, 56, 7, 32, 2,
	2, 56, 57, 7, 46, 2, 2, 57, 13, 3, 2, 2, 2, 58, 59, 7, 37, 2, 2, 59, 60,
	7, 46, 2, 2, 60, 15, 3, 2, 2, 2, 61, 62, 7, 30, 2, 2, 62, 63, 9, 3, 2,
	2, 63, 17, 3, 2, 2, 2, 7, 21, 29, 32, 40, 49,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\r'", "'\n'", "'.mtl'", "'on'", "'off'", "'v'", "'vt'", "'vn'", "'vp'",
	"'deg'", "'bmat'", "'step'", "'cstype'", "'p'", "'l'", "'f'", "'curv'",
	"'curv2'", "'surf'", "'parm'", "'trim'", "'hole'", "'scrv'", "'sp'", "'end'",
	"'con'", "'g'", "'s'", "'mg'", "'o'", "'bevel'", "'c_interp'", "'d_interp'",
	"'lod'", "'usemtl'", "'mtllib'", "'shadow_obj'", "'trace_obj'", "'ctech'",
	"'stech'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "GEOMETRIC_VERTEX", "TEXTURE_VERTEX", "VERTEX_NORMAL",
	"PARAMETER_SPACE_VERTEX", "DEGREE", "BASIS_MATRIX", "STEP_SIZE", "CURVE_SURF_TYPE",
	"POINT", "LINE", "FACE", "CURVE", "CURVE2D", "SURF", "PARAM", "OUTER_TRIMMING_HOLE",
	"INNER_TRIMMING_HOLE", "SPECIAL_CURVE", "SPECIAL_POINT", "END", "CONNECT",
	"GROUP_NAME", "SMOOTHING_GROUP", "MERGING_GROUP", "OBJECT_NAME", "BEVEL_INTERPOLATION",
	"COLOR_INTERPOLATION", "DISSOLVE_INTERPOLATION", "LEVEL_OF_DETAIL", "MATERIAL_NAME",
	"MATERIAL_LIBRARY", "SHADOW_CASTING", "RAY_TRACING", "CURVE_APPROX", "SURF_APPROX",
	"INTEGER", "DECIMAL", "COMMENT", "NAME", "WS",
}

var ruleNames = []string{
	"start", "line", "face", "vertex", "mtllib", "object", "use_material",
	"group",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type WavefrontOBJParser struct {
	*antlr.BaseParser
}

func NewWavefrontOBJParser(input antlr.TokenStream) *WavefrontOBJParser {
	this := new(WavefrontOBJParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "WavefrontOBJ.g4"

	return this
}

// WavefrontOBJParser tokens.
const (
	WavefrontOBJParserEOF                    = antlr.TokenEOF
	WavefrontOBJParserT__0                   = 1
	WavefrontOBJParserT__1                   = 2
	WavefrontOBJParserT__2                   = 3
	WavefrontOBJParserT__3                   = 4
	WavefrontOBJParserT__4                   = 5
	WavefrontOBJParserGEOMETRIC_VERTEX       = 6
	WavefrontOBJParserTEXTURE_VERTEX         = 7
	WavefrontOBJParserVERTEX_NORMAL          = 8
	WavefrontOBJParserPARAMETER_SPACE_VERTEX = 9
	WavefrontOBJParserDEGREE                 = 10
	WavefrontOBJParserBASIS_MATRIX           = 11
	WavefrontOBJParserSTEP_SIZE              = 12
	WavefrontOBJParserCURVE_SURF_TYPE        = 13
	WavefrontOBJParserPOINT                  = 14
	WavefrontOBJParserLINE                   = 15
	WavefrontOBJParserFACE                   = 16
	WavefrontOBJParserCURVE                  = 17
	WavefrontOBJParserCURVE2D                = 18
	WavefrontOBJParserSURF                   = 19
	WavefrontOBJParserPARAM                  = 20
	WavefrontOBJParserOUTER_TRIMMING_HOLE    = 21
	WavefrontOBJParserINNER_TRIMMING_HOLE    = 22
	WavefrontOBJParserSPECIAL_CURVE          = 23
	WavefrontOBJParserSPECIAL_POINT          = 24
	WavefrontOBJParserEND                    = 25
	WavefrontOBJParserCONNECT                = 26
	WavefrontOBJParserGROUP_NAME             = 27
	WavefrontOBJParserSMOOTHING_GROUP        = 28
	WavefrontOBJParserMERGING_GROUP          = 29
	WavefrontOBJParserOBJECT_NAME            = 30
	WavefrontOBJParserBEVEL_INTERPOLATION    = 31
	WavefrontOBJParserCOLOR_INTERPOLATION    = 32
	WavefrontOBJParserDISSOLVE_INTERPOLATION = 33
	WavefrontOBJParserLEVEL_OF_DETAIL        = 34
	WavefrontOBJParserMATERIAL_NAME          = 35
	WavefrontOBJParserMATERIAL_LIBRARY       = 36
	WavefrontOBJParserSHADOW_CASTING         = 37
	WavefrontOBJParserRAY_TRACING            = 38
	WavefrontOBJParserCURVE_APPROX           = 39
	WavefrontOBJParserSURF_APPROX            = 40
	WavefrontOBJParserINTEGER                = 41
	WavefrontOBJParserDECIMAL                = 42
	WavefrontOBJParserCOMMENT                = 43
	WavefrontOBJParserNAME                   = 44
	WavefrontOBJParserWS                     = 45
)

// WavefrontOBJParser rules.
const (
	WavefrontOBJParserRULE_start        = 0
	WavefrontOBJParserRULE_line         = 1
	WavefrontOBJParserRULE_face         = 2
	WavefrontOBJParserRULE_vertex       = 3
	WavefrontOBJParserRULE_mtllib       = 4
	WavefrontOBJParserRULE_object       = 5
	WavefrontOBJParserRULE_use_material = 6
	WavefrontOBJParserRULE_group        = 7
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *StartContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *WavefrontOBJParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WavefrontOBJParserRULE_start)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(WavefrontOBJParserGEOMETRIC_VERTEX-6))|(1<<(WavefrontOBJParserTEXTURE_VERTEX-6))|(1<<(WavefrontOBJParserVERTEX_NORMAL-6))|(1<<(WavefrontOBJParserPARAMETER_SPACE_VERTEX-6))|(1<<(WavefrontOBJParserFACE-6))|(1<<(WavefrontOBJParserSMOOTHING_GROUP-6))|(1<<(WavefrontOBJParserOBJECT_NAME-6))|(1<<(WavefrontOBJParserMATERIAL_NAME-6))|(1<<(WavefrontOBJParserMATERIAL_LIBRARY-6)))) != 0) {
		{
			p.SetState(16)
			p.Line()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = WavefrontOBJParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Vertex() IVertexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertexContext)
}

func (s *LineContext) Face() IFaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFaceContext)
}

func (s *LineContext) Mtllib() IMtllibContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMtllibContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMtllibContext)
}

func (s *LineContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *LineContext) Use_material() IUse_materialContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUse_materialContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUse_materialContext)
}

func (s *LineContext) Group() IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *WavefrontOBJParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WavefrontOBJParserRULE_line)
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WavefrontOBJParserGEOMETRIC_VERTEX, WavefrontOBJParserTEXTURE_VERTEX, WavefrontOBJParserVERTEX_NORMAL, WavefrontOBJParserPARAMETER_SPACE_VERTEX:
		{
			p.SetState(21)
			p.Vertex()
		}

	case WavefrontOBJParserFACE:
		{
			p.SetState(22)
			p.Face()
		}

	case WavefrontOBJParserMATERIAL_LIBRARY:
		{
			p.SetState(23)
			p.Mtllib()
		}

	case WavefrontOBJParserOBJECT_NAME:
		{
			p.SetState(24)
			p.Object()
		}

	case WavefrontOBJParserMATERIAL_NAME:
		{
			p.SetState(25)
			p.Use_material()
		}

	case WavefrontOBJParserSMOOTHING_GROUP:
		{
			p.SetState(26)
			p.Group()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == WavefrontOBJParserT__0 {
		{
			p.SetState(29)
			p.Match(WavefrontOBJParserT__0)
		}

	}
	{
		p.SetState(32)
		p.Match(WavefrontOBJParserT__1)
	}

	return localctx
}

// IFaceContext is an interface to support dynamic dispatch.
type IFaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFaceContext differentiates from other interfaces.
	IsFaceContext()
}

type FaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFaceContext() *FaceContext {
	var p = new(FaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_face
	return p
}

func (*FaceContext) IsFaceContext() {}

func NewFaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FaceContext {
	var p = new(FaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_face

	return p
}

func (s *FaceContext) GetParser() antlr.Parser { return s.parser }

func (s *FaceContext) FACE() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserFACE, 0)
}

func (s *FaceContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserINTEGER)
}

func (s *FaceContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserINTEGER, i)
}

func (s *FaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterFace(s)
	}
}

func (s *FaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitFace(s)
	}
}

func (p *WavefrontOBJParser) Face() (localctx IFaceContext) {
	localctx = NewFaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WavefrontOBJParserRULE_face)
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
		p.SetState(34)
		p.Match(WavefrontOBJParserFACE)
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == WavefrontOBJParserINTEGER {
		{
			p.SetState(35)
			p.Match(WavefrontOBJParserINTEGER)
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVertexContext is an interface to support dynamic dispatch.
type IVertexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVertexContext differentiates from other interfaces.
	IsVertexContext()
}

type VertexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVertexContext() *VertexContext {
	var p = new(VertexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_vertex
	return p
}

func (*VertexContext) IsVertexContext() {}

func NewVertexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VertexContext {
	var p = new(VertexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_vertex

	return p
}

func (s *VertexContext) GetParser() antlr.Parser { return s.parser }

func (s *VertexContext) AllDECIMAL() []antlr.TerminalNode {
	return s.GetTokens(WavefrontOBJParserDECIMAL)
}

func (s *VertexContext) DECIMAL(i int) antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserDECIMAL, i)
}

func (s *VertexContext) GEOMETRIC_VERTEX() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserGEOMETRIC_VERTEX, 0)
}

func (s *VertexContext) TEXTURE_VERTEX() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserTEXTURE_VERTEX, 0)
}

func (s *VertexContext) VERTEX_NORMAL() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserVERTEX_NORMAL, 0)
}

func (s *VertexContext) PARAMETER_SPACE_VERTEX() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserPARAMETER_SPACE_VERTEX, 0)
}

func (s *VertexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VertexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VertexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterVertex(s)
	}
}

func (s *VertexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitVertex(s)
	}
}

func (p *WavefrontOBJParser) Vertex() (localctx IVertexContext) {
	localctx = NewVertexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, WavefrontOBJParserRULE_vertex)
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

	p.SetState(47)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WavefrontOBJParserGEOMETRIC_VERTEX, WavefrontOBJParserTEXTURE_VERTEX, WavefrontOBJParserVERTEX_NORMAL:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(40)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<WavefrontOBJParserGEOMETRIC_VERTEX)|(1<<WavefrontOBJParserTEXTURE_VERTEX)|(1<<WavefrontOBJParserVERTEX_NORMAL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(41)
			p.Match(WavefrontOBJParserDECIMAL)
		}
		{
			p.SetState(42)
			p.Match(WavefrontOBJParserDECIMAL)
		}
		{
			p.SetState(43)
			p.Match(WavefrontOBJParserDECIMAL)
		}

	case WavefrontOBJParserPARAMETER_SPACE_VERTEX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(WavefrontOBJParserPARAMETER_SPACE_VERTEX)
		}
		{
			p.SetState(45)
			p.Match(WavefrontOBJParserDECIMAL)
		}
		{
			p.SetState(46)
			p.Match(WavefrontOBJParserDECIMAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMtllibContext is an interface to support dynamic dispatch.
type IMtllibContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetA returns the a token.
	GetA() antlr.Token

	// SetA sets the a token.
	SetA(antlr.Token)

	// IsMtllibContext differentiates from other interfaces.
	IsMtllibContext()
}

type MtllibContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	a      antlr.Token
}

func NewEmptyMtllibContext() *MtllibContext {
	var p = new(MtllibContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_mtllib
	return p
}

func (*MtllibContext) IsMtllibContext() {}

func NewMtllibContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MtllibContext {
	var p = new(MtllibContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_mtllib

	return p
}

func (s *MtllibContext) GetParser() antlr.Parser { return s.parser }

func (s *MtllibContext) GetA() antlr.Token { return s.a }

func (s *MtllibContext) SetA(v antlr.Token) { s.a = v }

func (s *MtllibContext) MATERIAL_LIBRARY() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserMATERIAL_LIBRARY, 0)
}

func (s *MtllibContext) NAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNAME, 0)
}

func (s *MtllibContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MtllibContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MtllibContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterMtllib(s)
	}
}

func (s *MtllibContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitMtllib(s)
	}
}

func (p *WavefrontOBJParser) Mtllib() (localctx IMtllibContext) {
	localctx = NewMtllibContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, WavefrontOBJParserRULE_mtllib)

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
		p.Match(WavefrontOBJParserMATERIAL_LIBRARY)
	}
	{
		p.SetState(50)

		var _m = p.Match(WavefrontOBJParserNAME)

		localctx.(*MtllibContext).a = _m
	}
	{
		p.SetState(51)
		p.Match(WavefrontOBJParserT__2)
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) OBJECT_NAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserOBJECT_NAME, 0)
}

func (s *ObjectContext) NAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNAME, 0)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitObject(s)
	}
}

func (p *WavefrontOBJParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, WavefrontOBJParserRULE_object)

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
		p.SetState(53)
		p.Match(WavefrontOBJParserOBJECT_NAME)
	}
	{
		p.SetState(54)
		p.Match(WavefrontOBJParserNAME)
	}

	return localctx
}

// IUse_materialContext is an interface to support dynamic dispatch.
type IUse_materialContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUse_materialContext differentiates from other interfaces.
	IsUse_materialContext()
}

type Use_materialContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUse_materialContext() *Use_materialContext {
	var p = new(Use_materialContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WavefrontOBJParserRULE_use_material
	return p
}

func (*Use_materialContext) IsUse_materialContext() {}

func NewUse_materialContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Use_materialContext {
	var p = new(Use_materialContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_use_material

	return p
}

func (s *Use_materialContext) GetParser() antlr.Parser { return s.parser }

func (s *Use_materialContext) MATERIAL_NAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserMATERIAL_NAME, 0)
}

func (s *Use_materialContext) NAME() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserNAME, 0)
}

func (s *Use_materialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Use_materialContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Use_materialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterUse_material(s)
	}
}

func (s *Use_materialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitUse_material(s)
	}
}

func (p *WavefrontOBJParser) Use_material() (localctx IUse_materialContext) {
	localctx = NewUse_materialContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, WavefrontOBJParserRULE_use_material)

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
		p.SetState(56)
		p.Match(WavefrontOBJParserMATERIAL_NAME)
	}
	{
		p.SetState(57)
		p.Match(WavefrontOBJParserNAME)
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
	p.RuleIndex = WavefrontOBJParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WavefrontOBJParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) SMOOTHING_GROUP() antlr.TerminalNode {
	return s.GetToken(WavefrontOBJParserSMOOTHING_GROUP, 0)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WavefrontOBJListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (p *WavefrontOBJParser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, WavefrontOBJParserRULE_group)
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
		p.SetState(59)
		p.Match(WavefrontOBJParserSMOOTHING_GROUP)
	}
	p.SetState(60)
	_la = p.GetTokenStream().LA(1)

	if !(_la == WavefrontOBJParserT__3 || _la == WavefrontOBJParserT__4) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
