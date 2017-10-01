// Generated from UCBLogo.g4 by ANTLR 4.7.

package ucb_logo // UCBLogo
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseUCBLogoListener is a complete listener for a parse tree produced by UCBLogoParser.
type BaseUCBLogoListener struct{}

var _ UCBLogoListener = &BaseUCBLogoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUCBLogoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUCBLogoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUCBLogoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUCBLogoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseUCBLogoListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseUCBLogoListener) ExitParse(ctx *ParseContext) {}

// EnterProcedureDefInstruction is called when production procedureDefInstruction is entered.
func (s *BaseUCBLogoListener) EnterProcedureDefInstruction(ctx *ProcedureDefInstructionContext) {}

// ExitProcedureDefInstruction is called when production procedureDefInstruction is exited.
func (s *BaseUCBLogoListener) ExitProcedureDefInstruction(ctx *ProcedureDefInstructionContext) {}

// EnterMacroDefInstruction is called when production macroDefInstruction is entered.
func (s *BaseUCBLogoListener) EnterMacroDefInstruction(ctx *MacroDefInstructionContext) {}

// ExitMacroDefInstruction is called when production macroDefInstruction is exited.
func (s *BaseUCBLogoListener) ExitMacroDefInstruction(ctx *MacroDefInstructionContext) {}

// EnterProcedureCallExtraInputInstruction is called when production procedureCallExtraInputInstruction is entered.
func (s *BaseUCBLogoListener) EnterProcedureCallExtraInputInstruction(ctx *ProcedureCallExtraInputInstructionContext) {
}

// ExitProcedureCallExtraInputInstruction is called when production procedureCallExtraInputInstruction is exited.
func (s *BaseUCBLogoListener) ExitProcedureCallExtraInputInstruction(ctx *ProcedureCallExtraInputInstructionContext) {
}

// EnterProcedureCallInstruction is called when production procedureCallInstruction is entered.
func (s *BaseUCBLogoListener) EnterProcedureCallInstruction(ctx *ProcedureCallInstructionContext) {}

// ExitProcedureCallInstruction is called when production procedureCallInstruction is exited.
func (s *BaseUCBLogoListener) ExitProcedureCallInstruction(ctx *ProcedureCallInstructionContext) {}

// EnterProcedure_def is called when production procedure_def is entered.
func (s *BaseUCBLogoListener) EnterProcedure_def(ctx *Procedure_defContext) {}

// ExitProcedure_def is called when production procedure_def is exited.
func (s *BaseUCBLogoListener) ExitProcedure_def(ctx *Procedure_defContext) {}

// EnterMacro_def is called when production macro_def is entered.
func (s *BaseUCBLogoListener) EnterMacro_def(ctx *Macro_defContext) {}

// ExitMacro_def is called when production macro_def is exited.
func (s *BaseUCBLogoListener) ExitMacro_def(ctx *Macro_defContext) {}

// EnterVariables is called when production variables is entered.
func (s *BaseUCBLogoListener) EnterVariables(ctx *VariablesContext) {}

// ExitVariables is called when production variables is exited.
func (s *BaseUCBLogoListener) ExitVariables(ctx *VariablesContext) {}

// EnterBody_def is called when production body_def is entered.
func (s *BaseUCBLogoListener) EnterBody_def(ctx *Body_defContext) {}

// ExitBody_def is called when production body_def is exited.
func (s *BaseUCBLogoListener) ExitBody_def(ctx *Body_defContext) {}

// EnterBody_instruction is called when production body_instruction is entered.
func (s *BaseUCBLogoListener) EnterBody_instruction(ctx *Body_instructionContext) {}

// ExitBody_instruction is called when production body_instruction is exited.
func (s *BaseUCBLogoListener) ExitBody_instruction(ctx *Body_instructionContext) {}

// EnterProcedure_call_extra_input is called when production procedure_call_extra_input is entered.
func (s *BaseUCBLogoListener) EnterProcedure_call_extra_input(ctx *Procedure_call_extra_inputContext) {
}

// ExitProcedure_call_extra_input is called when production procedure_call_extra_input is exited.
func (s *BaseUCBLogoListener) ExitProcedure_call_extra_input(ctx *Procedure_call_extra_inputContext) {}

// EnterProcedure_call is called when production procedure_call is entered.
func (s *BaseUCBLogoListener) EnterProcedure_call(ctx *Procedure_callContext) {}

// ExitProcedure_call is called when production procedure_call is exited.
func (s *BaseUCBLogoListener) ExitProcedure_call(ctx *Procedure_callContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseUCBLogoListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseUCBLogoListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterNotEqualsExpressionExpression is called when production notEqualsExpressionExpression is entered.
func (s *BaseUCBLogoListener) EnterNotEqualsExpressionExpression(ctx *NotEqualsExpressionExpressionContext) {
}

// ExitNotEqualsExpressionExpression is called when production notEqualsExpressionExpression is exited.
func (s *BaseUCBLogoListener) ExitNotEqualsExpressionExpression(ctx *NotEqualsExpressionExpressionContext) {
}

// EnterArrayExpression is called when production arrayExpression is entered.
func (s *BaseUCBLogoListener) EnterArrayExpression(ctx *ArrayExpressionContext) {}

// ExitArrayExpression is called when production arrayExpression is exited.
func (s *BaseUCBLogoListener) ExitArrayExpression(ctx *ArrayExpressionContext) {}

// EnterAdditionExpression is called when production additionExpression is entered.
func (s *BaseUCBLogoListener) EnterAdditionExpression(ctx *AdditionExpressionContext) {}

// ExitAdditionExpression is called when production additionExpression is exited.
func (s *BaseUCBLogoListener) ExitAdditionExpression(ctx *AdditionExpressionContext) {}

// EnterWordExpression is called when production wordExpression is entered.
func (s *BaseUCBLogoListener) EnterWordExpression(ctx *WordExpressionContext) {}

// ExitWordExpression is called when production wordExpression is exited.
func (s *BaseUCBLogoListener) ExitWordExpression(ctx *WordExpressionContext) {}

// EnterNumberExpression is called when production numberExpression is entered.
func (s *BaseUCBLogoListener) EnterNumberExpression(ctx *NumberExpressionContext) {}

// ExitNumberExpression is called when production numberExpression is exited.
func (s *BaseUCBLogoListener) ExitNumberExpression(ctx *NumberExpressionContext) {}

// EnterParensExpression is called when production parensExpression is entered.
func (s *BaseUCBLogoListener) EnterParensExpression(ctx *ParensExpressionContext) {}

// ExitParensExpression is called when production parensExpression is exited.
func (s *BaseUCBLogoListener) ExitParensExpression(ctx *ParensExpressionContext) {}

// EnterMultiplyExpression is called when production multiplyExpression is entered.
func (s *BaseUCBLogoListener) EnterMultiplyExpression(ctx *MultiplyExpressionContext) {}

// ExitMultiplyExpression is called when production multiplyExpression is exited.
func (s *BaseUCBLogoListener) ExitMultiplyExpression(ctx *MultiplyExpressionContext) {}

// EnterGreaterThanExpression is called when production greaterThanExpression is entered.
func (s *BaseUCBLogoListener) EnterGreaterThanExpression(ctx *GreaterThanExpressionContext) {}

// ExitGreaterThanExpression is called when production greaterThanExpression is exited.
func (s *BaseUCBLogoListener) ExitGreaterThanExpression(ctx *GreaterThanExpressionContext) {}

// EnterVariableExpression is called when production variableExpression is entered.
func (s *BaseUCBLogoListener) EnterVariableExpression(ctx *VariableExpressionContext) {}

// ExitVariableExpression is called when production variableExpression is exited.
func (s *BaseUCBLogoListener) ExitVariableExpression(ctx *VariableExpressionContext) {}

// EnterDivideExpression is called when production divideExpression is entered.
func (s *BaseUCBLogoListener) EnterDivideExpression(ctx *DivideExpressionContext) {}

// ExitDivideExpression is called when production divideExpression is exited.
func (s *BaseUCBLogoListener) ExitDivideExpression(ctx *DivideExpressionContext) {}

// EnterLessThanEqualsExpression is called when production lessThanEqualsExpression is entered.
func (s *BaseUCBLogoListener) EnterLessThanEqualsExpression(ctx *LessThanEqualsExpressionContext) {}

// ExitLessThanEqualsExpression is called when production lessThanEqualsExpression is exited.
func (s *BaseUCBLogoListener) ExitLessThanEqualsExpression(ctx *LessThanEqualsExpressionContext) {}

// EnterGreaterThanEqualsExpression is called when production greaterThanEqualsExpression is entered.
func (s *BaseUCBLogoListener) EnterGreaterThanEqualsExpression(ctx *GreaterThanEqualsExpressionContext) {
}

// ExitGreaterThanEqualsExpression is called when production greaterThanEqualsExpression is exited.
func (s *BaseUCBLogoListener) ExitGreaterThanEqualsExpression(ctx *GreaterThanEqualsExpressionContext) {
}

// EnterUnaryMinusExpression is called when production unaryMinusExpression is entered.
func (s *BaseUCBLogoListener) EnterUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// ExitUnaryMinusExpression is called when production unaryMinusExpression is exited.
func (s *BaseUCBLogoListener) ExitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// EnterQuotedWordExpression is called when production quotedWordExpression is entered.
func (s *BaseUCBLogoListener) EnterQuotedWordExpression(ctx *QuotedWordExpressionContext) {}

// ExitQuotedWordExpression is called when production quotedWordExpression is exited.
func (s *BaseUCBLogoListener) ExitQuotedWordExpression(ctx *QuotedWordExpressionContext) {}

// EnterEqualsExpression is called when production equalsExpression is entered.
func (s *BaseUCBLogoListener) EnterEqualsExpression(ctx *EqualsExpressionContext) {}

// ExitEqualsExpression is called when production equalsExpression is exited.
func (s *BaseUCBLogoListener) ExitEqualsExpression(ctx *EqualsExpressionContext) {}

// EnterSubtractionExpression is called when production subtractionExpression is entered.
func (s *BaseUCBLogoListener) EnterSubtractionExpression(ctx *SubtractionExpressionContext) {}

// ExitSubtractionExpression is called when production subtractionExpression is exited.
func (s *BaseUCBLogoListener) ExitSubtractionExpression(ctx *SubtractionExpressionContext) {}

// EnterProcedureCallExpression is called when production procedureCallExpression is entered.
func (s *BaseUCBLogoListener) EnterProcedureCallExpression(ctx *ProcedureCallExpressionContext) {}

// ExitProcedureCallExpression is called when production procedureCallExpression is exited.
func (s *BaseUCBLogoListener) ExitProcedureCallExpression(ctx *ProcedureCallExpressionContext) {}

// EnterLessThanExpression is called when production lessThanExpression is entered.
func (s *BaseUCBLogoListener) EnterLessThanExpression(ctx *LessThanExpressionContext) {}

// ExitLessThanExpression is called when production lessThanExpression is exited.
func (s *BaseUCBLogoListener) ExitLessThanExpression(ctx *LessThanExpressionContext) {}

// EnterProcedureCallExtraInput is called when production procedureCallExtraInput is entered.
func (s *BaseUCBLogoListener) EnterProcedureCallExtraInput(ctx *ProcedureCallExtraInputContext) {}

// ExitProcedureCallExtraInput is called when production procedureCallExtraInput is exited.
func (s *BaseUCBLogoListener) ExitProcedureCallExtraInput(ctx *ProcedureCallExtraInputContext) {}

// EnterListExpression is called when production listExpression is entered.
func (s *BaseUCBLogoListener) EnterListExpression(ctx *ListExpressionContext) {}

// ExitListExpression is called when production listExpression is exited.
func (s *BaseUCBLogoListener) ExitListExpression(ctx *ListExpressionContext) {}

// EnterNameExpression is called when production nameExpression is entered.
func (s *BaseUCBLogoListener) EnterNameExpression(ctx *NameExpressionContext) {}

// ExitNameExpression is called when production nameExpression is exited.
func (s *BaseUCBLogoListener) ExitNameExpression(ctx *NameExpressionContext) {}

// EnterArray is called when production array is entered.
func (s *BaseUCBLogoListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseUCBLogoListener) ExitArray(ctx *ArrayContext) {}

// EnterList is called when production list is entered.
func (s *BaseUCBLogoListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseUCBLogoListener) ExitList(ctx *ListContext) {}
