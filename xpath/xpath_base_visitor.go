// Generated from xpath.g4 by ANTLR 4.7.

package xpath // xpath
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasexpathVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasexpathVisitor) VisitMain(ctx *MainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitLocationPath(ctx *LocationPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitAbsoluteLocationPathNoroot(ctx *AbsoluteLocationPathNorootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitRelativeLocationPath(ctx *RelativeLocationPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitStep(ctx *StepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitAxisSpecifier(ctx *AxisSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitNodeTest(ctx *NodeTestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitAbbreviatedStep(ctx *AbbreviatedStepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitUnionExprNoRoot(ctx *UnionExprNoRootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitPathExprNoRoot(ctx *PathExprNoRootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitFilterExpr(ctx *FilterExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitRelationalExpr(ctx *RelationalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitUnaryExprNoRoot(ctx *UnaryExprNoRootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitQName(ctx *QNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitVariableReference(ctx *VariableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitNameTest(ctx *NameTestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexpathVisitor) VisitNCName(ctx *NCNameContext) interface{} {
	return v.VisitChildren(ctx)
}
