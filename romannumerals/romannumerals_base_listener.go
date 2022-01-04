// Code generated from romannumerals.g4 by ANTLR 4.9.3. DO NOT EDIT.

package romannumerals // romannumerals
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseromannumeralsListener is a complete listener for a parse tree produced by romannumeralsParser.
type BaseromannumeralsListener struct{}

var _ romannumeralsListener = &BaseromannumeralsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseromannumeralsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseromannumeralsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseromannumeralsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseromannumeralsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseromannumeralsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseromannumeralsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterThousands is called when production thousands is entered.
func (s *BaseromannumeralsListener) EnterThousands(ctx *ThousandsContext) {}

// ExitThousands is called when production thousands is exited.
func (s *BaseromannumeralsListener) ExitThousands(ctx *ThousandsContext) {}

// EnterThous_part is called when production thous_part is entered.
func (s *BaseromannumeralsListener) EnterThous_part(ctx *Thous_partContext) {}

// ExitThous_part is called when production thous_part is exited.
func (s *BaseromannumeralsListener) ExitThous_part(ctx *Thous_partContext) {}

// EnterHundreds is called when production hundreds is entered.
func (s *BaseromannumeralsListener) EnterHundreds(ctx *HundredsContext) {}

// ExitHundreds is called when production hundreds is exited.
func (s *BaseromannumeralsListener) ExitHundreds(ctx *HundredsContext) {}

// EnterHun_part is called when production hun_part is entered.
func (s *BaseromannumeralsListener) EnterHun_part(ctx *Hun_partContext) {}

// ExitHun_part is called when production hun_part is exited.
func (s *BaseromannumeralsListener) ExitHun_part(ctx *Hun_partContext) {}

// EnterHun_rep is called when production hun_rep is entered.
func (s *BaseromannumeralsListener) EnterHun_rep(ctx *Hun_repContext) {}

// ExitHun_rep is called when production hun_rep is exited.
func (s *BaseromannumeralsListener) ExitHun_rep(ctx *Hun_repContext) {}

// EnterTens is called when production tens is entered.
func (s *BaseromannumeralsListener) EnterTens(ctx *TensContext) {}

// ExitTens is called when production tens is exited.
func (s *BaseromannumeralsListener) ExitTens(ctx *TensContext) {}

// EnterTens_part is called when production tens_part is entered.
func (s *BaseromannumeralsListener) EnterTens_part(ctx *Tens_partContext) {}

// ExitTens_part is called when production tens_part is exited.
func (s *BaseromannumeralsListener) ExitTens_part(ctx *Tens_partContext) {}

// EnterTens_rep is called when production tens_rep is entered.
func (s *BaseromannumeralsListener) EnterTens_rep(ctx *Tens_repContext) {}

// ExitTens_rep is called when production tens_rep is exited.
func (s *BaseromannumeralsListener) ExitTens_rep(ctx *Tens_repContext) {}

// EnterOnes is called when production ones is entered.
func (s *BaseromannumeralsListener) EnterOnes(ctx *OnesContext) {}

// ExitOnes is called when production ones is exited.
func (s *BaseromannumeralsListener) ExitOnes(ctx *OnesContext) {}

// EnterOnes_rep is called when production ones_rep is entered.
func (s *BaseromannumeralsListener) EnterOnes_rep(ctx *Ones_repContext) {}

// ExitOnes_rep is called when production ones_rep is exited.
func (s *BaseromannumeralsListener) ExitOnes_rep(ctx *Ones_repContext) {}
