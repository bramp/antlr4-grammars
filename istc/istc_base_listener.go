// Generated from istc.g4 by ANTLR 4.7.

package istc // istc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseistcListener is a complete listener for a parse tree produced by istcParser.
type BaseistcListener struct{}

var _ istcListener = &BaseistcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseistcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseistcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseistcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseistcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterIstc is called when production istc is entered.
func (s *BaseistcListener) EnterIstc(ctx *IstcContext) {}

// ExitIstc is called when production istc is exited.
func (s *BaseistcListener) ExitIstc(ctx *IstcContext) {}

// EnterRegistration is called when production registration is entered.
func (s *BaseistcListener) EnterRegistration(ctx *RegistrationContext) {}

// ExitRegistration is called when production registration is exited.
func (s *BaseistcListener) ExitRegistration(ctx *RegistrationContext) {}

// EnterYear is called when production year is entered.
func (s *BaseistcListener) EnterYear(ctx *YearContext) {}

// ExitYear is called when production year is exited.
func (s *BaseistcListener) ExitYear(ctx *YearContext) {}

// EnterWork is called when production work is entered.
func (s *BaseistcListener) EnterWork(ctx *WorkContext) {}

// ExitWork is called when production work is exited.
func (s *BaseistcListener) ExitWork(ctx *WorkContext) {}

// EnterCheck is called when production check is entered.
func (s *BaseistcListener) EnterCheck(ctx *CheckContext) {}

// ExitCheck is called when production check is exited.
func (s *BaseistcListener) ExitCheck(ctx *CheckContext) {}
