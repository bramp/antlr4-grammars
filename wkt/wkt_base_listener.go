// Code generated from wkt.g4 by ANTLR 4.7.2. DO NOT EDIT.

package wkt // wkt
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasewktListener is a complete listener for a parse tree produced by wktParser.
type BasewktListener struct{}

var _ wktListener = &BasewktListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasewktListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasewktListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasewktListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasewktListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGeometry is called when production geometry is entered.
func (s *BasewktListener) EnterGeometry(ctx *GeometryContext) {}

// ExitGeometry is called when production geometry is exited.
func (s *BasewktListener) ExitGeometry(ctx *GeometryContext) {}

// EnterPointGeometry is called when production pointGeometry is entered.
func (s *BasewktListener) EnterPointGeometry(ctx *PointGeometryContext) {}

// ExitPointGeometry is called when production pointGeometry is exited.
func (s *BasewktListener) ExitPointGeometry(ctx *PointGeometryContext) {}

// EnterLineStringGeometry is called when production lineStringGeometry is entered.
func (s *BasewktListener) EnterLineStringGeometry(ctx *LineStringGeometryContext) {}

// ExitLineStringGeometry is called when production lineStringGeometry is exited.
func (s *BasewktListener) ExitLineStringGeometry(ctx *LineStringGeometryContext) {}

// EnterPolygonGeometry is called when production polygonGeometry is entered.
func (s *BasewktListener) EnterPolygonGeometry(ctx *PolygonGeometryContext) {}

// ExitPolygonGeometry is called when production polygonGeometry is exited.
func (s *BasewktListener) ExitPolygonGeometry(ctx *PolygonGeometryContext) {}

// EnterMultiPointGeometry is called when production multiPointGeometry is entered.
func (s *BasewktListener) EnterMultiPointGeometry(ctx *MultiPointGeometryContext) {}

// ExitMultiPointGeometry is called when production multiPointGeometry is exited.
func (s *BasewktListener) ExitMultiPointGeometry(ctx *MultiPointGeometryContext) {}

// EnterMultiLineStringGeometry is called when production multiLineStringGeometry is entered.
func (s *BasewktListener) EnterMultiLineStringGeometry(ctx *MultiLineStringGeometryContext) {}

// ExitMultiLineStringGeometry is called when production multiLineStringGeometry is exited.
func (s *BasewktListener) ExitMultiLineStringGeometry(ctx *MultiLineStringGeometryContext) {}

// EnterMultiPolygonGeometry is called when production multiPolygonGeometry is entered.
func (s *BasewktListener) EnterMultiPolygonGeometry(ctx *MultiPolygonGeometryContext) {}

// ExitMultiPolygonGeometry is called when production multiPolygonGeometry is exited.
func (s *BasewktListener) ExitMultiPolygonGeometry(ctx *MultiPolygonGeometryContext) {}

// EnterCircularStringGeometry is called when production circularStringGeometry is entered.
func (s *BasewktListener) EnterCircularStringGeometry(ctx *CircularStringGeometryContext) {}

// ExitCircularStringGeometry is called when production circularStringGeometry is exited.
func (s *BasewktListener) ExitCircularStringGeometry(ctx *CircularStringGeometryContext) {}

// EnterPointOrClosedPoint is called when production pointOrClosedPoint is entered.
func (s *BasewktListener) EnterPointOrClosedPoint(ctx *PointOrClosedPointContext) {}

// ExitPointOrClosedPoint is called when production pointOrClosedPoint is exited.
func (s *BasewktListener) ExitPointOrClosedPoint(ctx *PointOrClosedPointContext) {}

// EnterPolygon is called when production polygon is entered.
func (s *BasewktListener) EnterPolygon(ctx *PolygonContext) {}

// ExitPolygon is called when production polygon is exited.
func (s *BasewktListener) ExitPolygon(ctx *PolygonContext) {}

// EnterLineString is called when production lineString is entered.
func (s *BasewktListener) EnterLineString(ctx *LineStringContext) {}

// ExitLineString is called when production lineString is exited.
func (s *BasewktListener) ExitLineString(ctx *LineStringContext) {}

// EnterPoint is called when production point is entered.
func (s *BasewktListener) EnterPoint(ctx *PointContext) {}

// ExitPoint is called when production point is exited.
func (s *BasewktListener) ExitPoint(ctx *PointContext) {}

// EnterName is called when production name is entered.
func (s *BasewktListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasewktListener) ExitName(ctx *NameContext) {}
