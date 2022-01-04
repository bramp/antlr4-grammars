// Code generated from b.g4 by ANTLR 4.9.3. DO NOT EDIT.

package b // b
import "github.com/antlr/antlr4/runtime/Go/antlr"

// bListener is a complete listener for a parse tree produced by bParser.
type bListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterIval is called when entering the ival production.
	EnterIval(c *IvalContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterNullstmt is called when entering the nullstmt production.
	EnterNullstmt(c *NullstmtContext)

	// EnterExpressionstmt is called when entering the expressionstmt production.
	EnterExpressionstmt(c *ExpressionstmtContext)

	// EnterBlockstmt is called when entering the blockstmt production.
	EnterBlockstmt(c *BlockstmtContext)

	// EnterReturnstmt is called when entering the returnstmt production.
	EnterReturnstmt(c *ReturnstmtContext)

	// EnterGotostmt is called when entering the gotostmt production.
	EnterGotostmt(c *GotostmtContext)

	// EnterSwitchstmt is called when entering the switchstmt production.
	EnterSwitchstmt(c *SwitchstmtContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterCasestmt is called when entering the casestmt production.
	EnterCasestmt(c *CasestmtContext)

	// EnterExternsmt is called when entering the externsmt production.
	EnterExternsmt(c *ExternsmtContext)

	// EnterAutosmt is called when entering the autosmt production.
	EnterAutosmt(c *AutosmtContext)

	// EnterRvalue is called when entering the rvalue production.
	EnterRvalue(c *RvalueContext)

	// EnterTernary is called when entering the ternary production.
	EnterTernary(c *TernaryContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFunctioninvocation is called when entering the functioninvocation production.
	EnterFunctioninvocation(c *FunctioninvocationContext)

	// EnterFunctionparameters is called when entering the functionparameters production.
	EnterFunctionparameters(c *FunctionparametersContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterIncdec is called when entering the incdec production.
	EnterIncdec(c *IncdecContext)

	// EnterUnary is called when entering the unary production.
	EnterUnary(c *UnaryContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterLvalue is called when entering the lvalue production.
	EnterLvalue(c *LvalueContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitIval is called when exiting the ival production.
	ExitIval(c *IvalContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitNullstmt is called when exiting the nullstmt production.
	ExitNullstmt(c *NullstmtContext)

	// ExitExpressionstmt is called when exiting the expressionstmt production.
	ExitExpressionstmt(c *ExpressionstmtContext)

	// ExitBlockstmt is called when exiting the blockstmt production.
	ExitBlockstmt(c *BlockstmtContext)

	// ExitReturnstmt is called when exiting the returnstmt production.
	ExitReturnstmt(c *ReturnstmtContext)

	// ExitGotostmt is called when exiting the gotostmt production.
	ExitGotostmt(c *GotostmtContext)

	// ExitSwitchstmt is called when exiting the switchstmt production.
	ExitSwitchstmt(c *SwitchstmtContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitCasestmt is called when exiting the casestmt production.
	ExitCasestmt(c *CasestmtContext)

	// ExitExternsmt is called when exiting the externsmt production.
	ExitExternsmt(c *ExternsmtContext)

	// ExitAutosmt is called when exiting the autosmt production.
	ExitAutosmt(c *AutosmtContext)

	// ExitRvalue is called when exiting the rvalue production.
	ExitRvalue(c *RvalueContext)

	// ExitTernary is called when exiting the ternary production.
	ExitTernary(c *TernaryContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFunctioninvocation is called when exiting the functioninvocation production.
	ExitFunctioninvocation(c *FunctioninvocationContext)

	// ExitFunctionparameters is called when exiting the functionparameters production.
	ExitFunctionparameters(c *FunctionparametersContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitIncdec is called when exiting the incdec production.
	ExitIncdec(c *IncdecContext)

	// ExitUnary is called when exiting the unary production.
	ExitUnary(c *UnaryContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitLvalue is called when exiting the lvalue production.
	ExitLvalue(c *LvalueContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)
}
