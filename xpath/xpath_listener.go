// Generated from xpath.g4 by ANTLR 4.7.

package xpath // xpath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// xpathListener is a complete listener for a parse tree produced by xpathParser.
type xpathListener interface {
	antlr.ParseTreeListener

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterLocationPath is called when entering the locationPath production.
	EnterLocationPath(c *LocationPathContext)

	// EnterAbsoluteLocationPathNoroot is called when entering the absoluteLocationPathNoroot production.
	EnterAbsoluteLocationPathNoroot(c *AbsoluteLocationPathNorootContext)

	// EnterRelativeLocationPath is called when entering the relativeLocationPath production.
	EnterRelativeLocationPath(c *RelativeLocationPathContext)

	// EnterStep is called when entering the step production.
	EnterStep(c *StepContext)

	// EnterAxisSpecifier is called when entering the axisSpecifier production.
	EnterAxisSpecifier(c *AxisSpecifierContext)

	// EnterNodeTest is called when entering the nodeTest production.
	EnterNodeTest(c *NodeTestContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterAbbreviatedStep is called when entering the abbreviatedStep production.
	EnterAbbreviatedStep(c *AbbreviatedStepContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterUnionExprNoRoot is called when entering the unionExprNoRoot production.
	EnterUnionExprNoRoot(c *UnionExprNoRootContext)

	// EnterPathExprNoRoot is called when entering the pathExprNoRoot production.
	EnterPathExprNoRoot(c *PathExprNoRootContext)

	// EnterFilterExpr is called when entering the filterExpr production.
	EnterFilterExpr(c *FilterExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterEqualityExpr is called when entering the equalityExpr production.
	EnterEqualityExpr(c *EqualityExprContext)

	// EnterRelationalExpr is called when entering the relationalExpr production.
	EnterRelationalExpr(c *RelationalExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterMultiplicativeExpr is called when entering the multiplicativeExpr production.
	EnterMultiplicativeExpr(c *MultiplicativeExprContext)

	// EnterUnaryExprNoRoot is called when entering the unaryExprNoRoot production.
	EnterUnaryExprNoRoot(c *UnaryExprNoRootContext)

	// EnterQName is called when entering the qName production.
	EnterQName(c *QNameContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterVariableReference is called when entering the variableReference production.
	EnterVariableReference(c *VariableReferenceContext)

	// EnterNameTest is called when entering the nameTest production.
	EnterNameTest(c *NameTestContext)

	// EnterNCName is called when entering the nCName production.
	EnterNCName(c *NCNameContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitLocationPath is called when exiting the locationPath production.
	ExitLocationPath(c *LocationPathContext)

	// ExitAbsoluteLocationPathNoroot is called when exiting the absoluteLocationPathNoroot production.
	ExitAbsoluteLocationPathNoroot(c *AbsoluteLocationPathNorootContext)

	// ExitRelativeLocationPath is called when exiting the relativeLocationPath production.
	ExitRelativeLocationPath(c *RelativeLocationPathContext)

	// ExitStep is called when exiting the step production.
	ExitStep(c *StepContext)

	// ExitAxisSpecifier is called when exiting the axisSpecifier production.
	ExitAxisSpecifier(c *AxisSpecifierContext)

	// ExitNodeTest is called when exiting the nodeTest production.
	ExitNodeTest(c *NodeTestContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitAbbreviatedStep is called when exiting the abbreviatedStep production.
	ExitAbbreviatedStep(c *AbbreviatedStepContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitUnionExprNoRoot is called when exiting the unionExprNoRoot production.
	ExitUnionExprNoRoot(c *UnionExprNoRootContext)

	// ExitPathExprNoRoot is called when exiting the pathExprNoRoot production.
	ExitPathExprNoRoot(c *PathExprNoRootContext)

	// ExitFilterExpr is called when exiting the filterExpr production.
	ExitFilterExpr(c *FilterExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitEqualityExpr is called when exiting the equalityExpr production.
	ExitEqualityExpr(c *EqualityExprContext)

	// ExitRelationalExpr is called when exiting the relationalExpr production.
	ExitRelationalExpr(c *RelationalExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitMultiplicativeExpr is called when exiting the multiplicativeExpr production.
	ExitMultiplicativeExpr(c *MultiplicativeExprContext)

	// ExitUnaryExprNoRoot is called when exiting the unaryExprNoRoot production.
	ExitUnaryExprNoRoot(c *UnaryExprNoRootContext)

	// ExitQName is called when exiting the qName production.
	ExitQName(c *QNameContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitVariableReference is called when exiting the variableReference production.
	ExitVariableReference(c *VariableReferenceContext)

	// ExitNameTest is called when exiting the nameTest production.
	ExitNameTest(c *NameTestContext)

	// ExitNCName is called when exiting the nCName production.
	ExitNCName(c *NCNameContext)
}
