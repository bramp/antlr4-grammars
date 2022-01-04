// Code generated from PromQLParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package promql // PromQLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PromQLParserListener is a complete listener for a parse tree produced by PromQLParser.
type PromQLParserListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVectorOperation is called when entering the vectorOperation production.
	EnterVectorOperation(c *VectorOperationContext)

	// EnterUnaryOp is called when entering the unaryOp production.
	EnterUnaryOp(c *UnaryOpContext)

	// EnterPowOp is called when entering the powOp production.
	EnterPowOp(c *PowOpContext)

	// EnterMultOp is called when entering the multOp production.
	EnterMultOp(c *MultOpContext)

	// EnterAddOp is called when entering the addOp production.
	EnterAddOp(c *AddOpContext)

	// EnterCompareOp is called when entering the compareOp production.
	EnterCompareOp(c *CompareOpContext)

	// EnterAndUnlessOp is called when entering the andUnlessOp production.
	EnterAndUnlessOp(c *AndUnlessOpContext)

	// EnterOrOp is called when entering the orOp production.
	EnterOrOp(c *OrOpContext)

	// EnterVector is called when entering the vector production.
	EnterVector(c *VectorContext)

	// EnterParens is called when entering the parens production.
	EnterParens(c *ParensContext)

	// EnterInstantSelector is called when entering the instantSelector production.
	EnterInstantSelector(c *InstantSelectorContext)

	// EnterLabelMatcher is called when entering the labelMatcher production.
	EnterLabelMatcher(c *LabelMatcherContext)

	// EnterLabelMatcherOperator is called when entering the labelMatcherOperator production.
	EnterLabelMatcherOperator(c *LabelMatcherOperatorContext)

	// EnterLabelMatcherList is called when entering the labelMatcherList production.
	EnterLabelMatcherList(c *LabelMatcherListContext)

	// EnterMatrixSelector is called when entering the matrixSelector production.
	EnterMatrixSelector(c *MatrixSelectorContext)

	// EnterOffset is called when entering the offset production.
	EnterOffset(c *OffsetContext)

	// EnterFunction_ is called when entering the function_ production.
	EnterFunction_(c *Function_Context)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterAggregation is called when entering the aggregation production.
	EnterAggregation(c *AggregationContext)

	// EnterBy is called when entering the by production.
	EnterBy(c *ByContext)

	// EnterWithout is called when entering the without production.
	EnterWithout(c *WithoutContext)

	// EnterGrouping is called when entering the grouping production.
	EnterGrouping(c *GroupingContext)

	// EnterOn_ is called when entering the on_ production.
	EnterOn_(c *On_Context)

	// EnterIgnoring is called when entering the ignoring production.
	EnterIgnoring(c *IgnoringContext)

	// EnterGroupLeft is called when entering the groupLeft production.
	EnterGroupLeft(c *GroupLeftContext)

	// EnterGroupRight is called when entering the groupRight production.
	EnterGroupRight(c *GroupRightContext)

	// EnterLabelName is called when entering the labelName production.
	EnterLabelName(c *LabelNameContext)

	// EnterLabelNameList is called when entering the labelNameList production.
	EnterLabelNameList(c *LabelNameListContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVectorOperation is called when exiting the vectorOperation production.
	ExitVectorOperation(c *VectorOperationContext)

	// ExitUnaryOp is called when exiting the unaryOp production.
	ExitUnaryOp(c *UnaryOpContext)

	// ExitPowOp is called when exiting the powOp production.
	ExitPowOp(c *PowOpContext)

	// ExitMultOp is called when exiting the multOp production.
	ExitMultOp(c *MultOpContext)

	// ExitAddOp is called when exiting the addOp production.
	ExitAddOp(c *AddOpContext)

	// ExitCompareOp is called when exiting the compareOp production.
	ExitCompareOp(c *CompareOpContext)

	// ExitAndUnlessOp is called when exiting the andUnlessOp production.
	ExitAndUnlessOp(c *AndUnlessOpContext)

	// ExitOrOp is called when exiting the orOp production.
	ExitOrOp(c *OrOpContext)

	// ExitVector is called when exiting the vector production.
	ExitVector(c *VectorContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitInstantSelector is called when exiting the instantSelector production.
	ExitInstantSelector(c *InstantSelectorContext)

	// ExitLabelMatcher is called when exiting the labelMatcher production.
	ExitLabelMatcher(c *LabelMatcherContext)

	// ExitLabelMatcherOperator is called when exiting the labelMatcherOperator production.
	ExitLabelMatcherOperator(c *LabelMatcherOperatorContext)

	// ExitLabelMatcherList is called when exiting the labelMatcherList production.
	ExitLabelMatcherList(c *LabelMatcherListContext)

	// ExitMatrixSelector is called when exiting the matrixSelector production.
	ExitMatrixSelector(c *MatrixSelectorContext)

	// ExitOffset is called when exiting the offset production.
	ExitOffset(c *OffsetContext)

	// ExitFunction_ is called when exiting the function_ production.
	ExitFunction_(c *Function_Context)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitAggregation is called when exiting the aggregation production.
	ExitAggregation(c *AggregationContext)

	// ExitBy is called when exiting the by production.
	ExitBy(c *ByContext)

	// ExitWithout is called when exiting the without production.
	ExitWithout(c *WithoutContext)

	// ExitGrouping is called when exiting the grouping production.
	ExitGrouping(c *GroupingContext)

	// ExitOn_ is called when exiting the on_ production.
	ExitOn_(c *On_Context)

	// ExitIgnoring is called when exiting the ignoring production.
	ExitIgnoring(c *IgnoringContext)

	// ExitGroupLeft is called when exiting the groupLeft production.
	ExitGroupLeft(c *GroupLeftContext)

	// ExitGroupRight is called when exiting the groupRight production.
	ExitGroupRight(c *GroupRightContext)

	// ExitLabelName is called when exiting the labelName production.
	ExitLabelName(c *LabelNameContext)

	// ExitLabelNameList is called when exiting the labelNameList production.
	ExitLabelNameList(c *LabelNameListContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
