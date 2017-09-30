// Generated from MuParser.g4 by ANTLR 4.7.

package muparser // MuParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by MuParserParser.
type MuParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MuParserParser#progExpr.
	VisitProgExpr(ctx *ProgExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#functionMultiExpr.
	VisitFunctionMultiExpr(ctx *FunctionMultiExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#addSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#atomExpr.
	VisitAtomExpr(ctx *AtomExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#relationalExpr.
	VisitRelationalExpr(ctx *RelationalExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#functionExpr.
	VisitFunctionExpr(ctx *FunctionExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#unaryMinusExpr.
	VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#powExpr.
	VisitPowExpr(ctx *PowExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#assignExpr.
	VisitAssignExpr(ctx *AssignExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#mulDivExpr.
	VisitMulDivExpr(ctx *MulDivExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#equalityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#iteExpr.
	VisitIteExpr(ctx *IteExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#parExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by MuParserParser#numberAtom.
	VisitNumberAtom(ctx *NumberAtomContext) interface{}

	// Visit a parse tree produced by MuParserParser#booleanAtom.
	VisitBooleanAtom(ctx *BooleanAtomContext) interface{}

	// Visit a parse tree produced by MuParserParser#predefinedConstantAtom.
	VisitPredefinedConstantAtom(ctx *PredefinedConstantAtomContext) interface{}

	// Visit a parse tree produced by MuParserParser#idAtom.
	VisitIdAtom(ctx *IdAtomContext) interface{}
}
