// Code generated from SUOKIF.g4 by ANTLR 4.7.2. DO NOT EDIT.

package suokif // SUOKIF
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSUOKIFListener is a complete listener for a parse tree produced by SUOKIFParser.
type BaseSUOKIFListener struct{}

var _ SUOKIFListener = &BaseSUOKIFListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSUOKIFListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSUOKIFListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSUOKIFListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSUOKIFListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTop_level is called when production top_level is entered.
func (s *BaseSUOKIFListener) EnterTop_level(ctx *Top_levelContext) {}

// ExitTop_level is called when production top_level is exited.
func (s *BaseSUOKIFListener) ExitTop_level(ctx *Top_levelContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSUOKIFListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSUOKIFListener) ExitTerm(ctx *TermContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseSUOKIFListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseSUOKIFListener) ExitArgument(ctx *ArgumentContext) {}

// EnterFunterm is called when production funterm is entered.
func (s *BaseSUOKIFListener) EnterFunterm(ctx *FuntermContext) {}

// ExitFunterm is called when production funterm is exited.
func (s *BaseSUOKIFListener) ExitFunterm(ctx *FuntermContext) {}

// EnterSentence is called when production sentence is entered.
func (s *BaseSUOKIFListener) EnterSentence(ctx *SentenceContext) {}

// ExitSentence is called when production sentence is exited.
func (s *BaseSUOKIFListener) ExitSentence(ctx *SentenceContext) {}

// EnterEquation is called when production equation is entered.
func (s *BaseSUOKIFListener) EnterEquation(ctx *EquationContext) {}

// ExitEquation is called when production equation is exited.
func (s *BaseSUOKIFListener) ExitEquation(ctx *EquationContext) {}

// EnterRelsent is called when production relsent is entered.
func (s *BaseSUOKIFListener) EnterRelsent(ctx *RelsentContext) {}

// ExitRelsent is called when production relsent is exited.
func (s *BaseSUOKIFListener) ExitRelsent(ctx *RelsentContext) {}

// EnterLogsent is called when production logsent is entered.
func (s *BaseSUOKIFListener) EnterLogsent(ctx *LogsentContext) {}

// ExitLogsent is called when production logsent is exited.
func (s *BaseSUOKIFListener) ExitLogsent(ctx *LogsentContext) {}

// EnterQuantsent is called when production quantsent is entered.
func (s *BaseSUOKIFListener) EnterQuantsent(ctx *QuantsentContext) {}

// ExitQuantsent is called when production quantsent is exited.
func (s *BaseSUOKIFListener) ExitQuantsent(ctx *QuantsentContext) {}
