// Generated from metric.g4 by ANTLR 4.7.

package metric // metric
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by metricParser.
type metricVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by metricParser#uom.
	VisitUom(ctx *UomContext) interface{}

	// Visit a parse tree produced by metricParser#measure.
	VisitMeasure(ctx *MeasureContext) interface{}

	// Visit a parse tree produced by metricParser#exponent.
	VisitExponent(ctx *ExponentContext) interface{}

	// Visit a parse tree produced by metricParser#prefix.
	VisitPrefix(ctx *PrefixContext) interface{}

	// Visit a parse tree produced by metricParser#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by metricParser#baseunit.
	VisitBaseunit(ctx *BaseunitContext) interface{}

	// Visit a parse tree produced by metricParser#derivedunit.
	VisitDerivedunit(ctx *DerivedunitContext) interface{}
}
