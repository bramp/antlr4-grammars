// Generated from xpath.g4 by ANTLR 4.7.

package xpath // xpath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasexpathListener is a complete listener for a parse tree produced by xpathParser.
type BasexpathListener struct{}

var _ xpathListener = &BasexpathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasexpathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasexpathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasexpathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasexpathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMain is called when production main is entered.
func (s *BasexpathListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BasexpathListener) ExitMain(ctx *MainContext) {}

// EnterLocationPath is called when production locationPath is entered.
func (s *BasexpathListener) EnterLocationPath(ctx *LocationPathContext) {}

// ExitLocationPath is called when production locationPath is exited.
func (s *BasexpathListener) ExitLocationPath(ctx *LocationPathContext) {}

// EnterAbsoluteLocationPathNoroot is called when production absoluteLocationPathNoroot is entered.
func (s *BasexpathListener) EnterAbsoluteLocationPathNoroot(ctx *AbsoluteLocationPathNorootContext) {}

// ExitAbsoluteLocationPathNoroot is called when production absoluteLocationPathNoroot is exited.
func (s *BasexpathListener) ExitAbsoluteLocationPathNoroot(ctx *AbsoluteLocationPathNorootContext) {}

// EnterRelativeLocationPath is called when production relativeLocationPath is entered.
func (s *BasexpathListener) EnterRelativeLocationPath(ctx *RelativeLocationPathContext) {}

// ExitRelativeLocationPath is called when production relativeLocationPath is exited.
func (s *BasexpathListener) ExitRelativeLocationPath(ctx *RelativeLocationPathContext) {}

// EnterStep is called when production step is entered.
func (s *BasexpathListener) EnterStep(ctx *StepContext) {}

// ExitStep is called when production step is exited.
func (s *BasexpathListener) ExitStep(ctx *StepContext) {}

// EnterAxisSpecifier is called when production axisSpecifier is entered.
func (s *BasexpathListener) EnterAxisSpecifier(ctx *AxisSpecifierContext) {}

// ExitAxisSpecifier is called when production axisSpecifier is exited.
func (s *BasexpathListener) ExitAxisSpecifier(ctx *AxisSpecifierContext) {}

// EnterNodeTest is called when production nodeTest is entered.
func (s *BasexpathListener) EnterNodeTest(ctx *NodeTestContext) {}

// ExitNodeTest is called when production nodeTest is exited.
func (s *BasexpathListener) ExitNodeTest(ctx *NodeTestContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BasexpathListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BasexpathListener) ExitPredicate(ctx *PredicateContext) {}

// EnterAbbreviatedStep is called when production abbreviatedStep is entered.
func (s *BasexpathListener) EnterAbbreviatedStep(ctx *AbbreviatedStepContext) {}

// ExitAbbreviatedStep is called when production abbreviatedStep is exited.
func (s *BasexpathListener) ExitAbbreviatedStep(ctx *AbbreviatedStepContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasexpathListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasexpathListener) ExitExpr(ctx *ExprContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BasexpathListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BasexpathListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasexpathListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasexpathListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterUnionExprNoRoot is called when production unionExprNoRoot is entered.
func (s *BasexpathListener) EnterUnionExprNoRoot(ctx *UnionExprNoRootContext) {}

// ExitUnionExprNoRoot is called when production unionExprNoRoot is exited.
func (s *BasexpathListener) ExitUnionExprNoRoot(ctx *UnionExprNoRootContext) {}

// EnterPathExprNoRoot is called when production pathExprNoRoot is entered.
func (s *BasexpathListener) EnterPathExprNoRoot(ctx *PathExprNoRootContext) {}

// ExitPathExprNoRoot is called when production pathExprNoRoot is exited.
func (s *BasexpathListener) ExitPathExprNoRoot(ctx *PathExprNoRootContext) {}

// EnterFilterExpr is called when production filterExpr is entered.
func (s *BasexpathListener) EnterFilterExpr(ctx *FilterExprContext) {}

// ExitFilterExpr is called when production filterExpr is exited.
func (s *BasexpathListener) ExitFilterExpr(ctx *FilterExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BasexpathListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BasexpathListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BasexpathListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BasexpathListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterEqualityExpr is called when production equalityExpr is entered.
func (s *BasexpathListener) EnterEqualityExpr(ctx *EqualityExprContext) {}

// ExitEqualityExpr is called when production equalityExpr is exited.
func (s *BasexpathListener) ExitEqualityExpr(ctx *EqualityExprContext) {}

// EnterRelationalExpr is called when production relationalExpr is entered.
func (s *BasexpathListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production relationalExpr is exited.
func (s *BasexpathListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BasexpathListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BasexpathListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterMultiplicativeExpr is called when production multiplicativeExpr is entered.
func (s *BasexpathListener) EnterMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// ExitMultiplicativeExpr is called when production multiplicativeExpr is exited.
func (s *BasexpathListener) ExitMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// EnterUnaryExprNoRoot is called when production unaryExprNoRoot is entered.
func (s *BasexpathListener) EnterUnaryExprNoRoot(ctx *UnaryExprNoRootContext) {}

// ExitUnaryExprNoRoot is called when production unaryExprNoRoot is exited.
func (s *BasexpathListener) ExitUnaryExprNoRoot(ctx *UnaryExprNoRootContext) {}

// EnterQName is called when production qName is entered.
func (s *BasexpathListener) EnterQName(ctx *QNameContext) {}

// ExitQName is called when production qName is exited.
func (s *BasexpathListener) ExitQName(ctx *QNameContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BasexpathListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BasexpathListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterVariableReference is called when production variableReference is entered.
func (s *BasexpathListener) EnterVariableReference(ctx *VariableReferenceContext) {}

// ExitVariableReference is called when production variableReference is exited.
func (s *BasexpathListener) ExitVariableReference(ctx *VariableReferenceContext) {}

// EnterNameTest is called when production nameTest is entered.
func (s *BasexpathListener) EnterNameTest(ctx *NameTestContext) {}

// ExitNameTest is called when production nameTest is exited.
func (s *BasexpathListener) ExitNameTest(ctx *NameTestContext) {}

// EnterNCName is called when production nCName is entered.
func (s *BasexpathListener) EnterNCName(ctx *NCNameContext) {}

// ExitNCName is called when production nCName is exited.
func (s *BasexpathListener) ExitNCName(ctx *NCNameContext) {}
