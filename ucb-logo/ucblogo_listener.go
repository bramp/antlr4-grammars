// Generated from UCBLogo.g4 by ANTLR 4.7.

package ucb-logo // UCBLogo
import "github.com/antlr/antlr4/runtime/Go/antlr"

// UCBLogoListener is a complete listener for a parse tree produced by UCBLogoParser.
type UCBLogoListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterProcedureDefInstruction is called when entering the procedureDefInstruction production.
	EnterProcedureDefInstruction(c *ProcedureDefInstructionContext)

	// EnterMacroDefInstruction is called when entering the macroDefInstruction production.
	EnterMacroDefInstruction(c *MacroDefInstructionContext)

	// EnterProcedureCallExtraInputInstruction is called when entering the procedureCallExtraInputInstruction production.
	EnterProcedureCallExtraInputInstruction(c *ProcedureCallExtraInputInstructionContext)

	// EnterProcedureCallInstruction is called when entering the procedureCallInstruction production.
	EnterProcedureCallInstruction(c *ProcedureCallInstructionContext)

	// EnterProcedure_def is called when entering the procedure_def production.
	EnterProcedure_def(c *Procedure_defContext)

	// EnterMacro_def is called when entering the macro_def production.
	EnterMacro_def(c *Macro_defContext)

	// EnterVariables is called when entering the variables production.
	EnterVariables(c *VariablesContext)

	// EnterBody_def is called when entering the body_def production.
	EnterBody_def(c *Body_defContext)

	// EnterBody_instruction is called when entering the body_instruction production.
	EnterBody_instruction(c *Body_instructionContext)

	// EnterProcedure_call_extra_input is called when entering the procedure_call_extra_input production.
	EnterProcedure_call_extra_input(c *Procedure_call_extra_inputContext)

	// EnterProcedure_call is called when entering the procedure_call production.
	EnterProcedure_call(c *Procedure_callContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterNotEqualsExpressionExpression is called when entering the notEqualsExpressionExpression production.
	EnterNotEqualsExpressionExpression(c *NotEqualsExpressionExpressionContext)

	// EnterArrayExpression is called when entering the arrayExpression production.
	EnterArrayExpression(c *ArrayExpressionContext)

	// EnterAdditionExpression is called when entering the additionExpression production.
	EnterAdditionExpression(c *AdditionExpressionContext)

	// EnterWordExpression is called when entering the wordExpression production.
	EnterWordExpression(c *WordExpressionContext)

	// EnterNumberExpression is called when entering the numberExpression production.
	EnterNumberExpression(c *NumberExpressionContext)

	// EnterParensExpression is called when entering the parensExpression production.
	EnterParensExpression(c *ParensExpressionContext)

	// EnterMultiplyExpression is called when entering the multiplyExpression production.
	EnterMultiplyExpression(c *MultiplyExpressionContext)

	// EnterGreaterThanExpression is called when entering the greaterThanExpression production.
	EnterGreaterThanExpression(c *GreaterThanExpressionContext)

	// EnterVariableExpression is called when entering the variableExpression production.
	EnterVariableExpression(c *VariableExpressionContext)

	// EnterDivideExpression is called when entering the divideExpression production.
	EnterDivideExpression(c *DivideExpressionContext)

	// EnterLessThanEqualsExpression is called when entering the lessThanEqualsExpression production.
	EnterLessThanEqualsExpression(c *LessThanEqualsExpressionContext)

	// EnterGreaterThanEqualsExpression is called when entering the greaterThanEqualsExpression production.
	EnterGreaterThanEqualsExpression(c *GreaterThanEqualsExpressionContext)

	// EnterUnaryMinusExpression is called when entering the unaryMinusExpression production.
	EnterUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// EnterQuotedWordExpression is called when entering the quotedWordExpression production.
	EnterQuotedWordExpression(c *QuotedWordExpressionContext)

	// EnterEqualsExpression is called when entering the equalsExpression production.
	EnterEqualsExpression(c *EqualsExpressionContext)

	// EnterSubtractionExpression is called when entering the subtractionExpression production.
	EnterSubtractionExpression(c *SubtractionExpressionContext)

	// EnterProcedureCallExpression is called when entering the procedureCallExpression production.
	EnterProcedureCallExpression(c *ProcedureCallExpressionContext)

	// EnterLessThanExpression is called when entering the lessThanExpression production.
	EnterLessThanExpression(c *LessThanExpressionContext)

	// EnterProcedureCallExtraInput is called when entering the procedureCallExtraInput production.
	EnterProcedureCallExtraInput(c *ProcedureCallExtraInputContext)

	// EnterListExpression is called when entering the listExpression production.
	EnterListExpression(c *ListExpressionContext)

	// EnterNameExpression is called when entering the nameExpression production.
	EnterNameExpression(c *NameExpressionContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitProcedureDefInstruction is called when exiting the procedureDefInstruction production.
	ExitProcedureDefInstruction(c *ProcedureDefInstructionContext)

	// ExitMacroDefInstruction is called when exiting the macroDefInstruction production.
	ExitMacroDefInstruction(c *MacroDefInstructionContext)

	// ExitProcedureCallExtraInputInstruction is called when exiting the procedureCallExtraInputInstruction production.
	ExitProcedureCallExtraInputInstruction(c *ProcedureCallExtraInputInstructionContext)

	// ExitProcedureCallInstruction is called when exiting the procedureCallInstruction production.
	ExitProcedureCallInstruction(c *ProcedureCallInstructionContext)

	// ExitProcedure_def is called when exiting the procedure_def production.
	ExitProcedure_def(c *Procedure_defContext)

	// ExitMacro_def is called when exiting the macro_def production.
	ExitMacro_def(c *Macro_defContext)

	// ExitVariables is called when exiting the variables production.
	ExitVariables(c *VariablesContext)

	// ExitBody_def is called when exiting the body_def production.
	ExitBody_def(c *Body_defContext)

	// ExitBody_instruction is called when exiting the body_instruction production.
	ExitBody_instruction(c *Body_instructionContext)

	// ExitProcedure_call_extra_input is called when exiting the procedure_call_extra_input production.
	ExitProcedure_call_extra_input(c *Procedure_call_extra_inputContext)

	// ExitProcedure_call is called when exiting the procedure_call production.
	ExitProcedure_call(c *Procedure_callContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitNotEqualsExpressionExpression is called when exiting the notEqualsExpressionExpression production.
	ExitNotEqualsExpressionExpression(c *NotEqualsExpressionExpressionContext)

	// ExitArrayExpression is called when exiting the arrayExpression production.
	ExitArrayExpression(c *ArrayExpressionContext)

	// ExitAdditionExpression is called when exiting the additionExpression production.
	ExitAdditionExpression(c *AdditionExpressionContext)

	// ExitWordExpression is called when exiting the wordExpression production.
	ExitWordExpression(c *WordExpressionContext)

	// ExitNumberExpression is called when exiting the numberExpression production.
	ExitNumberExpression(c *NumberExpressionContext)

	// ExitParensExpression is called when exiting the parensExpression production.
	ExitParensExpression(c *ParensExpressionContext)

	// ExitMultiplyExpression is called when exiting the multiplyExpression production.
	ExitMultiplyExpression(c *MultiplyExpressionContext)

	// ExitGreaterThanExpression is called when exiting the greaterThanExpression production.
	ExitGreaterThanExpression(c *GreaterThanExpressionContext)

	// ExitVariableExpression is called when exiting the variableExpression production.
	ExitVariableExpression(c *VariableExpressionContext)

	// ExitDivideExpression is called when exiting the divideExpression production.
	ExitDivideExpression(c *DivideExpressionContext)

	// ExitLessThanEqualsExpression is called when exiting the lessThanEqualsExpression production.
	ExitLessThanEqualsExpression(c *LessThanEqualsExpressionContext)

	// ExitGreaterThanEqualsExpression is called when exiting the greaterThanEqualsExpression production.
	ExitGreaterThanEqualsExpression(c *GreaterThanEqualsExpressionContext)

	// ExitUnaryMinusExpression is called when exiting the unaryMinusExpression production.
	ExitUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// ExitQuotedWordExpression is called when exiting the quotedWordExpression production.
	ExitQuotedWordExpression(c *QuotedWordExpressionContext)

	// ExitEqualsExpression is called when exiting the equalsExpression production.
	ExitEqualsExpression(c *EqualsExpressionContext)

	// ExitSubtractionExpression is called when exiting the subtractionExpression production.
	ExitSubtractionExpression(c *SubtractionExpressionContext)

	// ExitProcedureCallExpression is called when exiting the procedureCallExpression production.
	ExitProcedureCallExpression(c *ProcedureCallExpressionContext)

	// ExitLessThanExpression is called when exiting the lessThanExpression production.
	ExitLessThanExpression(c *LessThanExpressionContext)

	// ExitProcedureCallExtraInput is called when exiting the procedureCallExtraInput production.
	ExitProcedureCallExtraInput(c *ProcedureCallExtraInputContext)

	// ExitListExpression is called when exiting the listExpression production.
	ExitListExpression(c *ListExpressionContext)

	// ExitNameExpression is called when exiting the nameExpression production.
	ExitNameExpression(c *NameExpressionContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)
}
