// Code generated from focal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package focal // focal
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasefocalListener is a complete listener for a parse tree produced by focalParser.
type BasefocalListener struct{}

var _ focalListener = &BasefocalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasefocalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasefocalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasefocalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasefocalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasefocalListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasefocalListener) ExitProg(ctx *ProgContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasefocalListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasefocalListener) ExitStatement(ctx *StatementContext) {}

// EnterGrpnum is called when production grpnum is entered.
func (s *BasefocalListener) EnterGrpnum(ctx *GrpnumContext) {}

// ExitGrpnum is called when production grpnum is exited.
func (s *BasefocalListener) ExitGrpnum(ctx *GrpnumContext) {}

// EnterLinenum is called when production linenum is entered.
func (s *BasefocalListener) EnterLinenum(ctx *LinenumContext) {}

// ExitLinenum is called when production linenum is exited.
func (s *BasefocalListener) ExitLinenum(ctx *LinenumContext) {}

// EnterCommand is called when production command is entered.
func (s *BasefocalListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BasefocalListener) ExitCommand(ctx *CommandContext) {}

// EnterAsk is called when production ask is entered.
func (s *BasefocalListener) EnterAsk(ctx *AskContext) {}

// ExitAsk is called when production ask is exited.
func (s *BasefocalListener) ExitAsk(ctx *AskContext) {}

// EnterAskpair is called when production askpair is entered.
func (s *BasefocalListener) EnterAskpair(ctx *AskpairContext) {}

// ExitAskpair is called when production askpair is exited.
func (s *BasefocalListener) ExitAskpair(ctx *AskpairContext) {}

// EnterDo_ is called when production do_ is entered.
func (s *BasefocalListener) EnterDo_(ctx *Do_Context) {}

// ExitDo_ is called when production do_ is exited.
func (s *BasefocalListener) ExitDo_(ctx *Do_Context) {}

// EnterFor_ is called when production for_ is entered.
func (s *BasefocalListener) EnterFor_(ctx *For_Context) {}

// ExitFor_ is called when production for_ is exited.
func (s *BasefocalListener) ExitFor_(ctx *For_Context) {}

// EnterQuit is called when production quit is entered.
func (s *BasefocalListener) EnterQuit(ctx *QuitContext) {}

// ExitQuit is called when production quit is exited.
func (s *BasefocalListener) ExitQuit(ctx *QuitContext) {}

// EnterSet_ is called when production set_ is entered.
func (s *BasefocalListener) EnterSet_(ctx *Set_Context) {}

// ExitSet_ is called when production set_ is exited.
func (s *BasefocalListener) ExitSet_(ctx *Set_Context) {}

// EnterGoto_ is called when production goto_ is entered.
func (s *BasefocalListener) EnterGoto_(ctx *Goto_Context) {}

// ExitGoto_ is called when production goto_ is exited.
func (s *BasefocalListener) ExitGoto_(ctx *Goto_Context) {}

// EnterIf_ is called when production if_ is entered.
func (s *BasefocalListener) EnterIf_(ctx *If_Context) {}

// ExitIf_ is called when production if_ is exited.
func (s *BasefocalListener) ExitIf_(ctx *If_Context) {}

// EnterType_ is called when production type_ is entered.
func (s *BasefocalListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BasefocalListener) ExitType_(ctx *Type_Context) {}

// EnterTypeexpression is called when production typeexpression is entered.
func (s *BasefocalListener) EnterTypeexpression(ctx *TypeexpressionContext) {}

// ExitTypeexpression is called when production typeexpression is exited.
func (s *BasefocalListener) ExitTypeexpression(ctx *TypeexpressionContext) {}

// EnterReturn_ is called when production return_ is entered.
func (s *BasefocalListener) EnterReturn_(ctx *Return_Context) {}

// ExitReturn_ is called when production return_ is exited.
func (s *BasefocalListener) ExitReturn_(ctx *Return_Context) {}

// EnterWrite_ is called when production write_ is entered.
func (s *BasefocalListener) EnterWrite_(ctx *Write_Context) {}

// ExitWrite_ is called when production write_ is exited.
func (s *BasefocalListener) ExitWrite_(ctx *Write_Context) {}

// EnterComment is called when production comment is entered.
func (s *BasefocalListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BasefocalListener) ExitComment(ctx *CommentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasefocalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasefocalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasefocalListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasefocalListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterTerm is called when production term is entered.
func (s *BasefocalListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasefocalListener) ExitTerm(ctx *TermContext) {}

// EnterNumber is called when production number is entered.
func (s *BasefocalListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasefocalListener) ExitNumber(ctx *NumberContext) {}

// EnterMantissa is called when production mantissa is entered.
func (s *BasefocalListener) EnterMantissa(ctx *MantissaContext) {}

// ExitMantissa is called when production mantissa is exited.
func (s *BasefocalListener) ExitMantissa(ctx *MantissaContext) {}

// EnterSigned_ is called when production signed_ is entered.
func (s *BasefocalListener) EnterSigned_(ctx *Signed_Context) {}

// ExitSigned_ is called when production signed_ is exited.
func (s *BasefocalListener) ExitSigned_(ctx *Signed_Context) {}
