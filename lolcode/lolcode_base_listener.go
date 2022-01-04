// Code generated from lolcode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lolcode // lolcode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaselolcodeListener is a complete listener for a parse tree produced by lolcodeParser.
type BaselolcodeListener struct{}

var _ lolcodeListener = &BaselolcodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaselolcodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaselolcodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselolcodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselolcodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaselolcodeListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaselolcodeListener) ExitProgram(ctx *ProgramContext) {}

// EnterCode_block is called when production code_block is entered.
func (s *BaselolcodeListener) EnterCode_block(ctx *Code_blockContext) {}

// ExitCode_block is called when production code_block is exited.
func (s *BaselolcodeListener) ExitCode_block(ctx *Code_blockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaselolcodeListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaselolcodeListener) ExitStatement(ctx *StatementContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaselolcodeListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaselolcodeListener) ExitLoop(ctx *LoopContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaselolcodeListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaselolcodeListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterComment is called when production comment is entered.
func (s *BaselolcodeListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaselolcodeListener) ExitComment(ctx *CommentContext) {}

// EnterPrint_block is called when production print_block is entered.
func (s *BaselolcodeListener) EnterPrint_block(ctx *Print_blockContext) {}

// ExitPrint_block is called when production print_block is exited.
func (s *BaselolcodeListener) ExitPrint_block(ctx *Print_blockContext) {}

// EnterIf_block is called when production if_block is entered.
func (s *BaselolcodeListener) EnterIf_block(ctx *If_blockContext) {}

// ExitIf_block is called when production if_block is exited.
func (s *BaselolcodeListener) ExitIf_block(ctx *If_blockContext) {}

// EnterElse_if_block is called when production else_if_block is entered.
func (s *BaselolcodeListener) EnterElse_if_block(ctx *Else_if_blockContext) {}

// ExitElse_if_block is called when production else_if_block is exited.
func (s *BaselolcodeListener) ExitElse_if_block(ctx *Else_if_blockContext) {}

// EnterInput_block is called when production input_block is entered.
func (s *BaselolcodeListener) EnterInput_block(ctx *Input_blockContext) {}

// ExitInput_block is called when production input_block is exited.
func (s *BaselolcodeListener) ExitInput_block(ctx *Input_blockContext) {}

// EnterFunc_decl is called when production func_decl is entered.
func (s *BaselolcodeListener) EnterFunc_decl(ctx *Func_declContext) {}

// ExitFunc_decl is called when production func_decl is exited.
func (s *BaselolcodeListener) ExitFunc_decl(ctx *Func_declContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaselolcodeListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaselolcodeListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaselolcodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaselolcodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEquals is called when production equals is entered.
func (s *BaselolcodeListener) EnterEquals(ctx *EqualsContext) {}

// ExitEquals is called when production equals is exited.
func (s *BaselolcodeListener) ExitEquals(ctx *EqualsContext) {}

// EnterNot_equals is called when production not_equals is entered.
func (s *BaselolcodeListener) EnterNot_equals(ctx *Not_equalsContext) {}

// ExitNot_equals is called when production not_equals is exited.
func (s *BaselolcodeListener) ExitNot_equals(ctx *Not_equalsContext) {}

// EnterBoth is called when production both is entered.
func (s *BaselolcodeListener) EnterBoth(ctx *BothContext) {}

// ExitBoth is called when production both is exited.
func (s *BaselolcodeListener) ExitBoth(ctx *BothContext) {}

// EnterEither is called when production either is entered.
func (s *BaselolcodeListener) EnterEither(ctx *EitherContext) {}

// ExitEither is called when production either is exited.
func (s *BaselolcodeListener) ExitEither(ctx *EitherContext) {}

// EnterGreater is called when production greater is entered.
func (s *BaselolcodeListener) EnterGreater(ctx *GreaterContext) {}

// ExitGreater is called when production greater is exited.
func (s *BaselolcodeListener) ExitGreater(ctx *GreaterContext) {}

// EnterLess is called when production less is entered.
func (s *BaselolcodeListener) EnterLess(ctx *LessContext) {}

// ExitLess is called when production less is exited.
func (s *BaselolcodeListener) ExitLess(ctx *LessContext) {}

// EnterAdd is called when production add is entered.
func (s *BaselolcodeListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production add is exited.
func (s *BaselolcodeListener) ExitAdd(ctx *AddContext) {}

// EnterSub is called when production sub is entered.
func (s *BaselolcodeListener) EnterSub(ctx *SubContext) {}

// ExitSub is called when production sub is exited.
func (s *BaselolcodeListener) ExitSub(ctx *SubContext) {}

// EnterMul is called when production mul is entered.
func (s *BaselolcodeListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production mul is exited.
func (s *BaselolcodeListener) ExitMul(ctx *MulContext) {}

// EnterDiv is called when production div is entered.
func (s *BaselolcodeListener) EnterDiv(ctx *DivContext) {}

// ExitDiv is called when production div is exited.
func (s *BaselolcodeListener) ExitDiv(ctx *DivContext) {}

// EnterMod is called when production mod is entered.
func (s *BaselolcodeListener) EnterMod(ctx *ModContext) {}

// ExitMod is called when production mod is exited.
func (s *BaselolcodeListener) ExitMod(ctx *ModContext) {}

// EnterCast is called when production cast is entered.
func (s *BaselolcodeListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaselolcodeListener) ExitCast(ctx *CastContext) {}

// EnterAll_ is called when production all_ is entered.
func (s *BaselolcodeListener) EnterAll_(ctx *All_Context) {}

// ExitAll_ is called when production all_ is exited.
func (s *BaselolcodeListener) ExitAll_(ctx *All_Context) {}

// EnterAny_ is called when production any_ is entered.
func (s *BaselolcodeListener) EnterAny_(ctx *Any_Context) {}

// ExitAny_ is called when production any_ is exited.
func (s *BaselolcodeListener) ExitAny_(ctx *Any_Context) {}

// EnterNot_ is called when production not_ is entered.
func (s *BaselolcodeListener) EnterNot_(ctx *Not_Context) {}

// ExitNot_ is called when production not_ is exited.
func (s *BaselolcodeListener) ExitNot_(ctx *Not_Context) {}

// EnterFunc_ is called when production func_ is entered.
func (s *BaselolcodeListener) EnterFunc_(ctx *Func_Context) {}

// ExitFunc_ is called when production func_ is exited.
func (s *BaselolcodeListener) ExitFunc_(ctx *Func_Context) {}
