// Code generated from wkt.g4 by ANTLR 4.7.2. DO NOT EDIT.

package wkt // wkt
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 146,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 36,
	10, 2, 13, 2, 14, 2, 37, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 44, 10, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 51, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 64, 10, 6, 12, 6, 14, 6, 67,
	11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 76, 10, 7, 12, 7,
	14, 7, 79, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 88, 10,
	8, 12, 8, 14, 8, 91, 11, 8, 3, 8, 3, 8, 3, 8, 5, 8, 96, 10, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 7, 9, 103, 10, 9, 12, 9, 14, 9, 106, 11, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 115, 10, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 7, 11, 121, 10, 11, 12, 11, 14, 11, 124, 11, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 132, 10, 12, 12, 12, 14, 12,
	135, 11, 12, 3, 12, 3, 12, 3, 13, 6, 13, 140, 10, 13, 13, 13, 14, 13, 141,
	3, 14, 3, 14, 3, 14, 2, 2, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 2, 2, 2, 150, 2, 35, 3, 2, 2, 2, 4, 41, 3, 2, 2, 2, 6, 52, 3, 2,
	2, 2, 8, 55, 3, 2, 2, 2, 10, 58, 3, 2, 2, 2, 12, 70, 3, 2, 2, 2, 14, 82,
	3, 2, 2, 2, 16, 97, 3, 2, 2, 2, 18, 114, 3, 2, 2, 2, 20, 116, 3, 2, 2,
	2, 22, 127, 3, 2, 2, 2, 24, 139, 3, 2, 2, 2, 26, 143, 3, 2, 2, 2, 28, 36,
	5, 8, 5, 2, 29, 36, 5, 6, 4, 2, 30, 36, 5, 4, 3, 2, 31, 36, 5, 10, 6, 2,
	32, 36, 5, 12, 7, 2, 33, 36, 5, 14, 8, 2, 34, 36, 5, 16, 9, 2, 35, 28,
	3, 2, 2, 2, 35, 29, 3, 2, 2, 2, 35, 30, 3, 2, 2, 2, 35, 31, 3, 2, 2, 2,
	35, 32, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 37, 3,
	2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39,
	40, 7, 2, 2, 3, 40, 3, 3, 2, 2, 2, 41, 50, 7, 9, 2, 2, 42, 44, 5, 26, 14,
	2, 43, 42, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46,
	7, 7, 2, 2, 46, 47, 5, 24, 13, 2, 47, 48, 7, 8, 2, 2, 48, 51, 3, 2, 2,
	2, 49, 51, 7, 16, 2, 2, 50, 43, 3, 2, 2, 2, 50, 49, 3, 2, 2, 2, 51, 5,
	3, 2, 2, 2, 52, 53, 7, 10, 2, 2, 53, 54, 5, 22, 12, 2, 54, 7, 3, 2, 2,
	2, 55, 56, 7, 11, 2, 2, 56, 57, 5, 20, 11, 2, 57, 9, 3, 2, 2, 2, 58, 59,
	7, 12, 2, 2, 59, 60, 7, 7, 2, 2, 60, 65, 5, 18, 10, 2, 61, 62, 7, 6, 2,
	2, 62, 64, 5, 18, 10, 2, 63, 61, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2, 65, 63,
	3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2,
	68, 69, 7, 8, 2, 2, 69, 11, 3, 2, 2, 2, 70, 71, 7, 13, 2, 2, 71, 72, 7,
	7, 2, 2, 72, 77, 5, 22, 12, 2, 73, 74, 7, 6, 2, 2, 74, 76, 5, 22, 12, 2,
	75, 73, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3,
	2, 2, 2, 78, 80, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81, 7, 8, 2, 2, 81,
	13, 3, 2, 2, 2, 82, 95, 7, 14, 2, 2, 83, 84, 7, 7, 2, 2, 84, 89, 5, 20,
	11, 2, 85, 86, 7, 6, 2, 2, 86, 88, 5, 20, 11, 2, 87, 85, 3, 2, 2, 2, 88,
	91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 92, 3, 2, 2,
	2, 91, 89, 3, 2, 2, 2, 92, 93, 7, 8, 2, 2, 93, 96, 3, 2, 2, 2, 94, 96,
	7, 16, 2, 2, 95, 83, 3, 2, 2, 2, 95, 94, 3, 2, 2, 2, 96, 15, 3, 2, 2, 2,
	97, 98, 7, 17, 2, 2, 98, 99, 7, 7, 2, 2, 99, 104, 5, 24, 13, 2, 100, 101,
	7, 6, 2, 2, 101, 103, 5, 24, 13, 2, 102, 100, 3, 2, 2, 2, 103, 106, 3,
	2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2,
	2, 106, 104, 3, 2, 2, 2, 107, 108, 7, 8, 2, 2, 108, 17, 3, 2, 2, 2, 109,
	115, 5, 24, 13, 2, 110, 111, 7, 7, 2, 2, 111, 112, 5, 24, 13, 2, 112, 113,
	7, 8, 2, 2, 113, 115, 3, 2, 2, 2, 114, 109, 3, 2, 2, 2, 114, 110, 3, 2,
	2, 2, 115, 19, 3, 2, 2, 2, 116, 117, 7, 7, 2, 2, 117, 122, 5, 22, 12, 2,
	118, 119, 7, 6, 2, 2, 119, 121, 5, 22, 12, 2, 120, 118, 3, 2, 2, 2, 121,
	124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 125,
	3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125, 126, 7, 8, 2, 2, 126, 21, 3, 2,
	2, 2, 127, 128, 7, 7, 2, 2, 128, 133, 5, 24, 13, 2, 129, 130, 7, 6, 2,
	2, 130, 132, 5, 24, 13, 2, 131, 129, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2,
	133, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 136, 3, 2, 2, 2, 135,
	133, 3, 2, 2, 2, 136, 137, 7, 8, 2, 2, 137, 23, 3, 2, 2, 2, 138, 140, 7,
	3, 2, 2, 139, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 139, 3, 2, 2,
	2, 141, 142, 3, 2, 2, 2, 142, 25, 3, 2, 2, 2, 143, 144, 7, 24, 2, 2, 144,
	27, 3, 2, 2, 2, 15, 35, 37, 43, 50, 65, 77, 89, 95, 104, 114, 122, 133,
	141,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "','", "'('", "')'",
}
var symbolicNames = []string{
	"", "DECIMAL", "INTEGERPART", "DECIMALPART", "COMMA", "LPAR", "RPAR", "POINT",
	"LINESTRING", "POLYGON", "MULTIPOINT", "MULTILINESTRING", "MULTIPOLYGON",
	"GEOMETRYCOLLECTION", "EMPTY", "CIRCULARSTRING", "COMPOUNDCURVE", "CURVEPOLYGON",
	"MULTICURVE", "TRIANGLE", "TIN", "POLYHEDRALSURFACE", "STRING", "WS",
}

var ruleNames = []string{
	"geometry", "pointGeometry", "lineStringGeometry", "polygonGeometry", "multiPointGeometry",
	"multiLineStringGeometry", "multiPolygonGeometry", "circularStringGeometry",
	"pointOrClosedPoint", "polygon", "lineString", "point", "name",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type wktParser struct {
	*antlr.BaseParser
}

func NewwktParser(input antlr.TokenStream) *wktParser {
	this := new(wktParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "wkt.g4"

	return this
}

// wktParser tokens.
const (
	wktParserEOF                = antlr.TokenEOF
	wktParserDECIMAL            = 1
	wktParserINTEGERPART        = 2
	wktParserDECIMALPART        = 3
	wktParserCOMMA              = 4
	wktParserLPAR               = 5
	wktParserRPAR               = 6
	wktParserPOINT              = 7
	wktParserLINESTRING         = 8
	wktParserPOLYGON            = 9
	wktParserMULTIPOINT         = 10
	wktParserMULTILINESTRING    = 11
	wktParserMULTIPOLYGON       = 12
	wktParserGEOMETRYCOLLECTION = 13
	wktParserEMPTY              = 14
	wktParserCIRCULARSTRING     = 15
	wktParserCOMPOUNDCURVE      = 16
	wktParserCURVEPOLYGON       = 17
	wktParserMULTICURVE         = 18
	wktParserTRIANGLE           = 19
	wktParserTIN                = 20
	wktParserPOLYHEDRALSURFACE  = 21
	wktParserSTRING             = 22
	wktParserWS                 = 23
)

// wktParser rules.
const (
	wktParserRULE_geometry                = 0
	wktParserRULE_pointGeometry           = 1
	wktParserRULE_lineStringGeometry      = 2
	wktParserRULE_polygonGeometry         = 3
	wktParserRULE_multiPointGeometry      = 4
	wktParserRULE_multiLineStringGeometry = 5
	wktParserRULE_multiPolygonGeometry    = 6
	wktParserRULE_circularStringGeometry  = 7
	wktParserRULE_pointOrClosedPoint      = 8
	wktParserRULE_polygon                 = 9
	wktParserRULE_lineString              = 10
	wktParserRULE_point                   = 11
	wktParserRULE_name                    = 12
)

// IGeometryContext is an interface to support dynamic dispatch.
type IGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeometryContext differentiates from other interfaces.
	IsGeometryContext()
}

type GeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeometryContext() *GeometryContext {
	var p = new(GeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_geometry
	return p
}

func (*GeometryContext) IsGeometryContext() {}

func NewGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeometryContext {
	var p = new(GeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_geometry

	return p
}

func (s *GeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *GeometryContext) EOF() antlr.TerminalNode {
	return s.GetToken(wktParserEOF, 0)
}

func (s *GeometryContext) AllPolygonGeometry() []IPolygonGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPolygonGeometryContext)(nil)).Elem())
	var tst = make([]IPolygonGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPolygonGeometryContext)
		}
	}

	return tst
}

func (s *GeometryContext) PolygonGeometry(i int) IPolygonGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPolygonGeometryContext)
}

func (s *GeometryContext) AllLineStringGeometry() []ILineStringGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineStringGeometryContext)(nil)).Elem())
	var tst = make([]ILineStringGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineStringGeometryContext)
		}
	}

	return tst
}

func (s *GeometryContext) LineStringGeometry(i int) ILineStringGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineStringGeometryContext)
}

func (s *GeometryContext) AllPointGeometry() []IPointGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointGeometryContext)(nil)).Elem())
	var tst = make([]IPointGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointGeometryContext)
		}
	}

	return tst
}

func (s *GeometryContext) PointGeometry(i int) IPointGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointGeometryContext)
}

func (s *GeometryContext) AllMultiPointGeometry() []IMultiPointGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiPointGeometryContext)(nil)).Elem())
	var tst = make([]IMultiPointGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiPointGeometryContext)
		}
	}

	return tst
}

func (s *GeometryContext) MultiPointGeometry(i int) IMultiPointGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPointGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiPointGeometryContext)
}

func (s *GeometryContext) AllMultiLineStringGeometry() []IMultiLineStringGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiLineStringGeometryContext)(nil)).Elem())
	var tst = make([]IMultiLineStringGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiLineStringGeometryContext)
		}
	}

	return tst
}

func (s *GeometryContext) MultiLineStringGeometry(i int) IMultiLineStringGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiLineStringGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiLineStringGeometryContext)
}

func (s *GeometryContext) AllMultiPolygonGeometry() []IMultiPolygonGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiPolygonGeometryContext)(nil)).Elem())
	var tst = make([]IMultiPolygonGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiPolygonGeometryContext)
		}
	}

	return tst
}

func (s *GeometryContext) MultiPolygonGeometry(i int) IMultiPolygonGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPolygonGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiPolygonGeometryContext)
}

func (s *GeometryContext) AllCircularStringGeometry() []ICircularStringGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICircularStringGeometryContext)(nil)).Elem())
	var tst = make([]ICircularStringGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICircularStringGeometryContext)
		}
	}

	return tst
}

func (s *GeometryContext) CircularStringGeometry(i int) ICircularStringGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICircularStringGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICircularStringGeometryContext)
}

func (s *GeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterGeometry(s)
	}
}

func (s *GeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitGeometry(s)
	}
}

func (p *wktParser) Geometry() (localctx IGeometryContext) {
	localctx = NewGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, wktParserRULE_geometry)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<wktParserPOINT)|(1<<wktParserLINESTRING)|(1<<wktParserPOLYGON)|(1<<wktParserMULTIPOINT)|(1<<wktParserMULTILINESTRING)|(1<<wktParserMULTIPOLYGON)|(1<<wktParserCIRCULARSTRING))) != 0) {
		p.SetState(33)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case wktParserPOLYGON:
			{
				p.SetState(26)
				p.PolygonGeometry()
			}

		case wktParserLINESTRING:
			{
				p.SetState(27)
				p.LineStringGeometry()
			}

		case wktParserPOINT:
			{
				p.SetState(28)
				p.PointGeometry()
			}

		case wktParserMULTIPOINT:
			{
				p.SetState(29)
				p.MultiPointGeometry()
			}

		case wktParserMULTILINESTRING:
			{
				p.SetState(30)
				p.MultiLineStringGeometry()
			}

		case wktParserMULTIPOLYGON:
			{
				p.SetState(31)
				p.MultiPolygonGeometry()
			}

		case wktParserCIRCULARSTRING:
			{
				p.SetState(32)
				p.CircularStringGeometry()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
		p.Match(wktParserEOF)
	}

	return localctx
}

// IPointGeometryContext is an interface to support dynamic dispatch.
type IPointGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointGeometryContext differentiates from other interfaces.
	IsPointGeometryContext()
}

type PointGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointGeometryContext() *PointGeometryContext {
	var p = new(PointGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_pointGeometry
	return p
}

func (*PointGeometryContext) IsPointGeometryContext() {}

func NewPointGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointGeometryContext {
	var p = new(PointGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_pointGeometry

	return p
}

func (s *PointGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *PointGeometryContext) POINT() antlr.TerminalNode {
	return s.GetToken(wktParserPOINT, 0)
}

func (s *PointGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *PointGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *PointGeometryContext) Point() IPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *PointGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *PointGeometryContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *PointGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterPointGeometry(s)
	}
}

func (s *PointGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitPointGeometry(s)
	}
}

func (p *wktParser) PointGeometry() (localctx IPointGeometryContext) {
	localctx = NewPointGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, wktParserRULE_pointGeometry)
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
		p.SetState(39)
		p.Match(wktParserPOINT)
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR, wktParserSTRING:
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == wktParserSTRING {
			{
				p.SetState(40)
				p.Name()
			}

		}
		{
			p.SetState(43)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(44)
			p.Point()
		}
		{
			p.SetState(45)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(47)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILineStringGeometryContext is an interface to support dynamic dispatch.
type ILineStringGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineStringGeometryContext differentiates from other interfaces.
	IsLineStringGeometryContext()
}

type LineStringGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineStringGeometryContext() *LineStringGeometryContext {
	var p = new(LineStringGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_lineStringGeometry
	return p
}

func (*LineStringGeometryContext) IsLineStringGeometryContext() {}

func NewLineStringGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineStringGeometryContext {
	var p = new(LineStringGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_lineStringGeometry

	return p
}

func (s *LineStringGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *LineStringGeometryContext) LINESTRING() antlr.TerminalNode {
	return s.GetToken(wktParserLINESTRING, 0)
}

func (s *LineStringGeometryContext) LineString() ILineStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineStringContext)
}

func (s *LineStringGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineStringGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineStringGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterLineStringGeometry(s)
	}
}

func (s *LineStringGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitLineStringGeometry(s)
	}
}

func (p *wktParser) LineStringGeometry() (localctx ILineStringGeometryContext) {
	localctx = NewLineStringGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, wktParserRULE_lineStringGeometry)

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
		p.Match(wktParserLINESTRING)
	}
	{
		p.SetState(51)
		p.LineString()
	}

	return localctx
}

// IPolygonGeometryContext is an interface to support dynamic dispatch.
type IPolygonGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolygonGeometryContext differentiates from other interfaces.
	IsPolygonGeometryContext()
}

type PolygonGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolygonGeometryContext() *PolygonGeometryContext {
	var p = new(PolygonGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_polygonGeometry
	return p
}

func (*PolygonGeometryContext) IsPolygonGeometryContext() {}

func NewPolygonGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonGeometryContext {
	var p = new(PolygonGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_polygonGeometry

	return p
}

func (s *PolygonGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonGeometryContext) POLYGON() antlr.TerminalNode {
	return s.GetToken(wktParserPOLYGON, 0)
}

func (s *PolygonGeometryContext) Polygon() IPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPolygonContext)
}

func (s *PolygonGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterPolygonGeometry(s)
	}
}

func (s *PolygonGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitPolygonGeometry(s)
	}
}

func (p *wktParser) PolygonGeometry() (localctx IPolygonGeometryContext) {
	localctx = NewPolygonGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, wktParserRULE_polygonGeometry)

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
		p.Match(wktParserPOLYGON)
	}
	{
		p.SetState(54)
		p.Polygon()
	}

	return localctx
}

// IMultiPointGeometryContext is an interface to support dynamic dispatch.
type IMultiPointGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiPointGeometryContext differentiates from other interfaces.
	IsMultiPointGeometryContext()
}

type MultiPointGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiPointGeometryContext() *MultiPointGeometryContext {
	var p = new(MultiPointGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_multiPointGeometry
	return p
}

func (*MultiPointGeometryContext) IsMultiPointGeometryContext() {}

func NewMultiPointGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPointGeometryContext {
	var p = new(MultiPointGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_multiPointGeometry

	return p
}

func (s *MultiPointGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPointGeometryContext) MULTIPOINT() antlr.TerminalNode {
	return s.GetToken(wktParserMULTIPOINT, 0)
}

func (s *MultiPointGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *MultiPointGeometryContext) AllPointOrClosedPoint() []IPointOrClosedPointContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointOrClosedPointContext)(nil)).Elem())
	var tst = make([]IPointOrClosedPointContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointOrClosedPointContext)
		}
	}

	return tst
}

func (s *MultiPointGeometryContext) PointOrClosedPoint(i int) IPointOrClosedPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointOrClosedPointContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointOrClosedPointContext)
}

func (s *MultiPointGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *MultiPointGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *MultiPointGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *MultiPointGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPointGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPointGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterMultiPointGeometry(s)
	}
}

func (s *MultiPointGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitMultiPointGeometry(s)
	}
}

func (p *wktParser) MultiPointGeometry() (localctx IMultiPointGeometryContext) {
	localctx = NewMultiPointGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, wktParserRULE_multiPointGeometry)
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
		p.SetState(56)
		p.Match(wktParserMULTIPOINT)
	}
	{
		p.SetState(57)
		p.Match(wktParserLPAR)
	}
	{
		p.SetState(58)
		p.PointOrClosedPoint()
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == wktParserCOMMA {
		{
			p.SetState(59)
			p.Match(wktParserCOMMA)
		}
		{
			p.SetState(60)
			p.PointOrClosedPoint()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(wktParserRPAR)
	}

	return localctx
}

// IMultiLineStringGeometryContext is an interface to support dynamic dispatch.
type IMultiLineStringGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiLineStringGeometryContext differentiates from other interfaces.
	IsMultiLineStringGeometryContext()
}

type MultiLineStringGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiLineStringGeometryContext() *MultiLineStringGeometryContext {
	var p = new(MultiLineStringGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_multiLineStringGeometry
	return p
}

func (*MultiLineStringGeometryContext) IsMultiLineStringGeometryContext() {}

func NewMultiLineStringGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiLineStringGeometryContext {
	var p = new(MultiLineStringGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_multiLineStringGeometry

	return p
}

func (s *MultiLineStringGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiLineStringGeometryContext) MULTILINESTRING() antlr.TerminalNode {
	return s.GetToken(wktParserMULTILINESTRING, 0)
}

func (s *MultiLineStringGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *MultiLineStringGeometryContext) AllLineString() []ILineStringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineStringContext)(nil)).Elem())
	var tst = make([]ILineStringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineStringContext)
		}
	}

	return tst
}

func (s *MultiLineStringGeometryContext) LineString(i int) ILineStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineStringContext)
}

func (s *MultiLineStringGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *MultiLineStringGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *MultiLineStringGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *MultiLineStringGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiLineStringGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiLineStringGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterMultiLineStringGeometry(s)
	}
}

func (s *MultiLineStringGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitMultiLineStringGeometry(s)
	}
}

func (p *wktParser) MultiLineStringGeometry() (localctx IMultiLineStringGeometryContext) {
	localctx = NewMultiLineStringGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, wktParserRULE_multiLineStringGeometry)
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
		p.Match(wktParserMULTILINESTRING)
	}
	{
		p.SetState(69)
		p.Match(wktParserLPAR)
	}
	{
		p.SetState(70)
		p.LineString()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == wktParserCOMMA {
		{
			p.SetState(71)
			p.Match(wktParserCOMMA)
		}
		{
			p.SetState(72)
			p.LineString()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(78)
		p.Match(wktParserRPAR)
	}

	return localctx
}

// IMultiPolygonGeometryContext is an interface to support dynamic dispatch.
type IMultiPolygonGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiPolygonGeometryContext differentiates from other interfaces.
	IsMultiPolygonGeometryContext()
}

type MultiPolygonGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiPolygonGeometryContext() *MultiPolygonGeometryContext {
	var p = new(MultiPolygonGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_multiPolygonGeometry
	return p
}

func (*MultiPolygonGeometryContext) IsMultiPolygonGeometryContext() {}

func NewMultiPolygonGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPolygonGeometryContext {
	var p = new(MultiPolygonGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_multiPolygonGeometry

	return p
}

func (s *MultiPolygonGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPolygonGeometryContext) MULTIPOLYGON() antlr.TerminalNode {
	return s.GetToken(wktParserMULTIPOLYGON, 0)
}

func (s *MultiPolygonGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *MultiPolygonGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *MultiPolygonGeometryContext) AllPolygon() []IPolygonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPolygonContext)(nil)).Elem())
	var tst = make([]IPolygonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPolygonContext)
		}
	}

	return tst
}

func (s *MultiPolygonGeometryContext) Polygon(i int) IPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPolygonContext)
}

func (s *MultiPolygonGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *MultiPolygonGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *MultiPolygonGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *MultiPolygonGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPolygonGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPolygonGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterMultiPolygonGeometry(s)
	}
}

func (s *MultiPolygonGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitMultiPolygonGeometry(s)
	}
}

func (p *wktParser) MultiPolygonGeometry() (localctx IMultiPolygonGeometryContext) {
	localctx = NewMultiPolygonGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, wktParserRULE_multiPolygonGeometry)
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
		p.SetState(80)
		p.Match(wktParserMULTIPOLYGON)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(81)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(82)
			p.Polygon()
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(83)
				p.Match(wktParserCOMMA)
			}
			{
				p.SetState(84)
				p.Polygon()
			}

			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(90)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(92)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICircularStringGeometryContext is an interface to support dynamic dispatch.
type ICircularStringGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCircularStringGeometryContext differentiates from other interfaces.
	IsCircularStringGeometryContext()
}

type CircularStringGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCircularStringGeometryContext() *CircularStringGeometryContext {
	var p = new(CircularStringGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_circularStringGeometry
	return p
}

func (*CircularStringGeometryContext) IsCircularStringGeometryContext() {}

func NewCircularStringGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CircularStringGeometryContext {
	var p = new(CircularStringGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_circularStringGeometry

	return p
}

func (s *CircularStringGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *CircularStringGeometryContext) CIRCULARSTRING() antlr.TerminalNode {
	return s.GetToken(wktParserCIRCULARSTRING, 0)
}

func (s *CircularStringGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *CircularStringGeometryContext) AllPoint() []IPointContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointContext)(nil)).Elem())
	var tst = make([]IPointContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointContext)
		}
	}

	return tst
}

func (s *CircularStringGeometryContext) Point(i int) IPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *CircularStringGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *CircularStringGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *CircularStringGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *CircularStringGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CircularStringGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CircularStringGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterCircularStringGeometry(s)
	}
}

func (s *CircularStringGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitCircularStringGeometry(s)
	}
}

func (p *wktParser) CircularStringGeometry() (localctx ICircularStringGeometryContext) {
	localctx = NewCircularStringGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, wktParserRULE_circularStringGeometry)
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
		p.SetState(95)
		p.Match(wktParserCIRCULARSTRING)
	}
	{
		p.SetState(96)
		p.Match(wktParserLPAR)
	}
	{
		p.SetState(97)
		p.Point()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == wktParserCOMMA {
		{
			p.SetState(98)
			p.Match(wktParserCOMMA)
		}
		{
			p.SetState(99)
			p.Point()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
		p.Match(wktParserRPAR)
	}

	return localctx
}

// IPointOrClosedPointContext is an interface to support dynamic dispatch.
type IPointOrClosedPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointOrClosedPointContext differentiates from other interfaces.
	IsPointOrClosedPointContext()
}

type PointOrClosedPointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointOrClosedPointContext() *PointOrClosedPointContext {
	var p = new(PointOrClosedPointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_pointOrClosedPoint
	return p
}

func (*PointOrClosedPointContext) IsPointOrClosedPointContext() {}

func NewPointOrClosedPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointOrClosedPointContext {
	var p = new(PointOrClosedPointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_pointOrClosedPoint

	return p
}

func (s *PointOrClosedPointContext) GetParser() antlr.Parser { return s.parser }

func (s *PointOrClosedPointContext) Point() IPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *PointOrClosedPointContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *PointOrClosedPointContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *PointOrClosedPointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointOrClosedPointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointOrClosedPointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterPointOrClosedPoint(s)
	}
}

func (s *PointOrClosedPointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitPointOrClosedPoint(s)
	}
}

func (p *wktParser) PointOrClosedPoint() (localctx IPointOrClosedPointContext) {
	localctx = NewPointOrClosedPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, wktParserRULE_pointOrClosedPoint)

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

	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserDECIMAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Point()
		}

	case wktParserLPAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(109)
			p.Point()
		}
		{
			p.SetState(110)
			p.Match(wktParserRPAR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPolygonContext is an interface to support dynamic dispatch.
type IPolygonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolygonContext differentiates from other interfaces.
	IsPolygonContext()
}

type PolygonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolygonContext() *PolygonContext {
	var p = new(PolygonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_polygon
	return p
}

func (*PolygonContext) IsPolygonContext() {}

func NewPolygonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonContext {
	var p = new(PolygonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_polygon

	return p
}

func (s *PolygonContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *PolygonContext) AllLineString() []ILineStringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineStringContext)(nil)).Elem())
	var tst = make([]ILineStringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineStringContext)
		}
	}

	return tst
}

func (s *PolygonContext) LineString(i int) ILineStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineStringContext)
}

func (s *PolygonContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *PolygonContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *PolygonContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *PolygonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterPolygon(s)
	}
}

func (s *PolygonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitPolygon(s)
	}
}

func (p *wktParser) Polygon() (localctx IPolygonContext) {
	localctx = NewPolygonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, wktParserRULE_polygon)
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
		p.SetState(114)
		p.Match(wktParserLPAR)
	}
	{
		p.SetState(115)
		p.LineString()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == wktParserCOMMA {
		{
			p.SetState(116)
			p.Match(wktParserCOMMA)
		}
		{
			p.SetState(117)
			p.LineString()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(123)
		p.Match(wktParserRPAR)
	}

	return localctx
}

// ILineStringContext is an interface to support dynamic dispatch.
type ILineStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineStringContext differentiates from other interfaces.
	IsLineStringContext()
}

type LineStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineStringContext() *LineStringContext {
	var p = new(LineStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_lineString
	return p
}

func (*LineStringContext) IsLineStringContext() {}

func NewLineStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineStringContext {
	var p = new(LineStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_lineString

	return p
}

func (s *LineStringContext) GetParser() antlr.Parser { return s.parser }

func (s *LineStringContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *LineStringContext) AllPoint() []IPointContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointContext)(nil)).Elem())
	var tst = make([]IPointContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointContext)
		}
	}

	return tst
}

func (s *LineStringContext) Point(i int) IPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *LineStringContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *LineStringContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *LineStringContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *LineStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterLineString(s)
	}
}

func (s *LineStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitLineString(s)
	}
}

func (p *wktParser) LineString() (localctx ILineStringContext) {
	localctx = NewLineStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, wktParserRULE_lineString)
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
		p.SetState(125)
		p.Match(wktParserLPAR)
	}
	{
		p.SetState(126)
		p.Point()
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == wktParserCOMMA {
		{
			p.SetState(127)
			p.Match(wktParserCOMMA)
		}
		{
			p.SetState(128)
			p.Point()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(wktParserRPAR)
	}

	return localctx
}

// IPointContext is an interface to support dynamic dispatch.
type IPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointContext differentiates from other interfaces.
	IsPointContext()
}

type PointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointContext() *PointContext {
	var p = new(PointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_point
	return p
}

func (*PointContext) IsPointContext() {}

func NewPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointContext {
	var p = new(PointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_point

	return p
}

func (s *PointContext) GetParser() antlr.Parser { return s.parser }

func (s *PointContext) AllDECIMAL() []antlr.TerminalNode {
	return s.GetTokens(wktParserDECIMAL)
}

func (s *PointContext) DECIMAL(i int) antlr.TerminalNode {
	return s.GetToken(wktParserDECIMAL, i)
}

func (s *PointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterPoint(s)
	}
}

func (s *PointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitPoint(s)
	}
}

func (p *wktParser) Point() (localctx IPointContext) {
	localctx = NewPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, wktParserRULE_point)
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
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == wktParserDECIMAL {
		{
			p.SetState(136)
			p.Match(wktParserDECIMAL)
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = wktParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(wktParserSTRING, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *wktParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, wktParserRULE_name)

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
		p.SetState(141)
		p.Match(wktParserSTRING)
	}

	return localctx
}
