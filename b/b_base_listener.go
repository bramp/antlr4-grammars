// Code generated from b.g4 by ANTLR 4.9.3. DO NOT EDIT.

package b // b
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasebListener is a complete listener for a parse tree produced by bParser.
type BasebListener struct{}

var _ bListener = &BasebListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasebListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasebListener) ExitProgram(ctx *ProgramContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BasebListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BasebListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterIval is called when production ival is entered.
func (s *BasebListener) EnterIval(ctx *IvalContext) {}

// ExitIval is called when production ival is exited.
func (s *BasebListener) ExitIval(ctx *IvalContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasebListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasebListener) ExitStatement(ctx *StatementContext) {}

// EnterNullstmt is called when production nullstmt is entered.
func (s *BasebListener) EnterNullstmt(ctx *NullstmtContext) {}

// ExitNullstmt is called when production nullstmt is exited.
func (s *BasebListener) ExitNullstmt(ctx *NullstmtContext) {}

// EnterExpressionstmt is called when production expressionstmt is entered.
func (s *BasebListener) EnterExpressionstmt(ctx *ExpressionstmtContext) {}

// ExitExpressionstmt is called when production expressionstmt is exited.
func (s *BasebListener) ExitExpressionstmt(ctx *ExpressionstmtContext) {}

// EnterBlockstmt is called when production blockstmt is entered.
func (s *BasebListener) EnterBlockstmt(ctx *BlockstmtContext) {}

// ExitBlockstmt is called when production blockstmt is exited.
func (s *BasebListener) ExitBlockstmt(ctx *BlockstmtContext) {}

// EnterReturnstmt is called when production returnstmt is entered.
func (s *BasebListener) EnterReturnstmt(ctx *ReturnstmtContext) {}

// ExitReturnstmt is called when production returnstmt is exited.
func (s *BasebListener) ExitReturnstmt(ctx *ReturnstmtContext) {}

// EnterGotostmt is called when production gotostmt is entered.
func (s *BasebListener) EnterGotostmt(ctx *GotostmtContext) {}

// ExitGotostmt is called when production gotostmt is exited.
func (s *BasebListener) ExitGotostmt(ctx *GotostmtContext) {}

// EnterSwitchstmt is called when production switchstmt is entered.
func (s *BasebListener) EnterSwitchstmt(ctx *SwitchstmtContext) {}

// ExitSwitchstmt is called when production switchstmt is exited.
func (s *BasebListener) ExitSwitchstmt(ctx *SwitchstmtContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *BasebListener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *BasebListener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BasebListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BasebListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterCasestmt is called when production casestmt is entered.
func (s *BasebListener) EnterCasestmt(ctx *CasestmtContext) {}

// ExitCasestmt is called when production casestmt is exited.
func (s *BasebListener) ExitCasestmt(ctx *CasestmtContext) {}

// EnterExternsmt is called when production externsmt is entered.
func (s *BasebListener) EnterExternsmt(ctx *ExternsmtContext) {}

// ExitExternsmt is called when production externsmt is exited.
func (s *BasebListener) ExitExternsmt(ctx *ExternsmtContext) {}

// EnterAutosmt is called when production autosmt is entered.
func (s *BasebListener) EnterAutosmt(ctx *AutosmtContext) {}

// ExitAutosmt is called when production autosmt is exited.
func (s *BasebListener) ExitAutosmt(ctx *AutosmtContext) {}

// EnterRvalue is called when production rvalue is entered.
func (s *BasebListener) EnterRvalue(ctx *RvalueContext) {}

// ExitRvalue is called when production rvalue is exited.
func (s *BasebListener) ExitRvalue(ctx *RvalueContext) {}

// EnterTernary is called when production ternary is entered.
func (s *BasebListener) EnterTernary(ctx *TernaryContext) {}

// ExitTernary is called when production ternary is exited.
func (s *BasebListener) ExitTernary(ctx *TernaryContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BasebListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BasebListener) ExitComparison(ctx *ComparisonContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasebListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasebListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasebListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasebListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFunctioninvocation is called when production functioninvocation is entered.
func (s *BasebListener) EnterFunctioninvocation(ctx *FunctioninvocationContext) {}

// ExitFunctioninvocation is called when production functioninvocation is exited.
func (s *BasebListener) ExitFunctioninvocation(ctx *FunctioninvocationContext) {}

// EnterFunctionparameters is called when production functionparameters is entered.
func (s *BasebListener) EnterFunctionparameters(ctx *FunctionparametersContext) {}

// ExitFunctionparameters is called when production functionparameters is exited.
func (s *BasebListener) ExitFunctionparameters(ctx *FunctionparametersContext) {}

// EnterAssign is called when production assign is entered.
func (s *BasebListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasebListener) ExitAssign(ctx *AssignContext) {}

// EnterIncdec is called when production incdec is entered.
func (s *BasebListener) EnterIncdec(ctx *IncdecContext) {}

// ExitIncdec is called when production incdec is exited.
func (s *BasebListener) ExitIncdec(ctx *IncdecContext) {}

// EnterUnary is called when production unary is entered.
func (s *BasebListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production unary is exited.
func (s *BasebListener) ExitUnary(ctx *UnaryContext) {}

// EnterBinary is called when production binary is entered.
func (s *BasebListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BasebListener) ExitBinary(ctx *BinaryContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BasebListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BasebListener) ExitLvalue(ctx *LvalueContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasebListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasebListener) ExitConstant(ctx *ConstantContext) {}

// EnterName is called when production name is entered.
func (s *BasebListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasebListener) ExitName(ctx *NameContext) {}
