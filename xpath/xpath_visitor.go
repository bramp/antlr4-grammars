// Generated from xpath.g4 by ANTLR 4.7.

package xpath // xpath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by xpathParser.
type xpathVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by xpathParser#main.
	VisitMain(ctx *MainContext) interface{}

	// Visit a parse tree produced by xpathParser#locationPath.
	VisitLocationPath(ctx *LocationPathContext) interface{}

	// Visit a parse tree produced by xpathParser#absoluteLocationPathNoroot.
	VisitAbsoluteLocationPathNoroot(ctx *AbsoluteLocationPathNorootContext) interface{}

	// Visit a parse tree produced by xpathParser#relativeLocationPath.
	VisitRelativeLocationPath(ctx *RelativeLocationPathContext) interface{}

	// Visit a parse tree produced by xpathParser#step.
	VisitStep(ctx *StepContext) interface{}

	// Visit a parse tree produced by xpathParser#axisSpecifier.
	VisitAxisSpecifier(ctx *AxisSpecifierContext) interface{}

	// Visit a parse tree produced by xpathParser#nodeTest.
	VisitNodeTest(ctx *NodeTestContext) interface{}

	// Visit a parse tree produced by xpathParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by xpathParser#abbreviatedStep.
	VisitAbbreviatedStep(ctx *AbbreviatedStepContext) interface{}

	// Visit a parse tree produced by xpathParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by xpathParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by xpathParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by xpathParser#unionExprNoRoot.
	VisitUnionExprNoRoot(ctx *UnionExprNoRootContext) interface{}

	// Visit a parse tree produced by xpathParser#pathExprNoRoot.
	VisitPathExprNoRoot(ctx *PathExprNoRootContext) interface{}

	// Visit a parse tree produced by xpathParser#filterExpr.
	VisitFilterExpr(ctx *FilterExprContext) interface{}

	// Visit a parse tree produced by xpathParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by xpathParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by xpathParser#equalityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by xpathParser#relationalExpr.
	VisitRelationalExpr(ctx *RelationalExprContext) interface{}

	// Visit a parse tree produced by xpathParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by xpathParser#multiplicativeExpr.
	VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{}

	// Visit a parse tree produced by xpathParser#unaryExprNoRoot.
	VisitUnaryExprNoRoot(ctx *UnaryExprNoRootContext) interface{}

	// Visit a parse tree produced by xpathParser#qName.
	VisitQName(ctx *QNameContext) interface{}

	// Visit a parse tree produced by xpathParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by xpathParser#variableReference.
	VisitVariableReference(ctx *VariableReferenceContext) interface{}

	// Visit a parse tree produced by xpathParser#nameTest.
	VisitNameTest(ctx *NameTestContext) interface{}

	// Visit a parse tree produced by xpathParser#nCName.
	VisitNCName(ctx *NCNameContext) interface{}
}
