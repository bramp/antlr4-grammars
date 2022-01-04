// Code generated from lolcode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lolcode // lolcode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// lolcodeListener is a complete listener for a parse tree produced by lolcodeParser.
type lolcodeListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterCode_block is called when entering the code_block production.
	EnterCode_block(c *Code_blockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterPrint_block is called when entering the print_block production.
	EnterPrint_block(c *Print_blockContext)

	// EnterIf_block is called when entering the if_block production.
	EnterIf_block(c *If_blockContext)

	// EnterElse_if_block is called when entering the else_if_block production.
	EnterElse_if_block(c *Else_if_blockContext)

	// EnterInput_block is called when entering the input_block production.
	EnterInput_block(c *Input_blockContext)

	// EnterFunc_decl is called when entering the func_decl production.
	EnterFunc_decl(c *Func_declContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEquals is called when entering the equals production.
	EnterEquals(c *EqualsContext)

	// EnterNot_equals is called when entering the not_equals production.
	EnterNot_equals(c *Not_equalsContext)

	// EnterBoth is called when entering the both production.
	EnterBoth(c *BothContext)

	// EnterEither is called when entering the either production.
	EnterEither(c *EitherContext)

	// EnterGreater is called when entering the greater production.
	EnterGreater(c *GreaterContext)

	// EnterLess is called when entering the less production.
	EnterLess(c *LessContext)

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

	// EnterSub is called when entering the sub production.
	EnterSub(c *SubContext)

	// EnterMul is called when entering the mul production.
	EnterMul(c *MulContext)

	// EnterDiv is called when entering the div production.
	EnterDiv(c *DivContext)

	// EnterMod is called when entering the mod production.
	EnterMod(c *ModContext)

	// EnterCast is called when entering the cast production.
	EnterCast(c *CastContext)

	// EnterAll_ is called when entering the all_ production.
	EnterAll_(c *All_Context)

	// EnterAny_ is called when entering the any_ production.
	EnterAny_(c *Any_Context)

	// EnterNot_ is called when entering the not_ production.
	EnterNot_(c *Not_Context)

	// EnterFunc_ is called when entering the func_ production.
	EnterFunc_(c *Func_Context)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitCode_block is called when exiting the code_block production.
	ExitCode_block(c *Code_blockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitPrint_block is called when exiting the print_block production.
	ExitPrint_block(c *Print_blockContext)

	// ExitIf_block is called when exiting the if_block production.
	ExitIf_block(c *If_blockContext)

	// ExitElse_if_block is called when exiting the else_if_block production.
	ExitElse_if_block(c *Else_if_blockContext)

	// ExitInput_block is called when exiting the input_block production.
	ExitInput_block(c *Input_blockContext)

	// ExitFunc_decl is called when exiting the func_decl production.
	ExitFunc_decl(c *Func_declContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEquals is called when exiting the equals production.
	ExitEquals(c *EqualsContext)

	// ExitNot_equals is called when exiting the not_equals production.
	ExitNot_equals(c *Not_equalsContext)

	// ExitBoth is called when exiting the both production.
	ExitBoth(c *BothContext)

	// ExitEither is called when exiting the either production.
	ExitEither(c *EitherContext)

	// ExitGreater is called when exiting the greater production.
	ExitGreater(c *GreaterContext)

	// ExitLess is called when exiting the less production.
	ExitLess(c *LessContext)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

	// ExitSub is called when exiting the sub production.
	ExitSub(c *SubContext)

	// ExitMul is called when exiting the mul production.
	ExitMul(c *MulContext)

	// ExitDiv is called when exiting the div production.
	ExitDiv(c *DivContext)

	// ExitMod is called when exiting the mod production.
	ExitMod(c *ModContext)

	// ExitCast is called when exiting the cast production.
	ExitCast(c *CastContext)

	// ExitAll_ is called when exiting the all_ production.
	ExitAll_(c *All_Context)

	// ExitAny_ is called when exiting the any_ production.
	ExitAny_(c *Any_Context)

	// ExitNot_ is called when exiting the not_ production.
	ExitNot_(c *Not_Context)

	// ExitFunc_ is called when exiting the func_ production.
	ExitFunc_(c *Func_Context)
}
