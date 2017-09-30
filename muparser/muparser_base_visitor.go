// Generated from MuParser.g4 by ANTLR 4.7.

package muparser // MuParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMuParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMuParserVisitor) VisitProgExpr(ctx *ProgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitFunctionMultiExpr(ctx *FunctionMultiExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitAtomExpr(ctx *AtomExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitRelationalExpr(ctx *RelationalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitFunctionExpr(ctx *FunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitUnaryMinusExpr(ctx *UnaryMinusExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitPowExpr(ctx *PowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitAssignExpr(ctx *AssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitMulDivExpr(ctx *MulDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitIteExpr(ctx *IteExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitNumberAtom(ctx *NumberAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitBooleanAtom(ctx *BooleanAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitPredefinedConstantAtom(ctx *PredefinedConstantAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMuParserVisitor) VisitIdAtom(ctx *IdAtomContext) interface{} {
	return v.VisitChildren(ctx)
}
