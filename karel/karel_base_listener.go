// Code generated from karel.g4 by ANTLR 4.9.3. DO NOT EDIT.

package karel // karel
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasekarelListener is a complete listener for a parse tree produced by karelParser.
type BasekarelListener struct{}

var _ karelListener = &BasekarelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasekarelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasekarelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasekarelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasekarelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKarel is called when production karel is entered.
func (s *BasekarelListener) EnterKarel(ctx *KarelContext) {}

// ExitKarel is called when production karel is exited.
func (s *BasekarelListener) ExitKarel(ctx *KarelContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BasekarelListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BasekarelListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasekarelListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasekarelListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BasekarelListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasekarelListener) ExitBlock(ctx *BlockContext) {}

// EnterIteration is called when production iteration is entered.
func (s *BasekarelListener) EnterIteration(ctx *IterationContext) {}

// ExitIteration is called when production iteration is exited.
func (s *BasekarelListener) ExitIteration(ctx *IterationContext) {}

// EnterLoop is called when production loop is entered.
func (s *BasekarelListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BasekarelListener) ExitLoop(ctx *LoopContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BasekarelListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BasekarelListener) ExitConditional(ctx *ConditionalContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BasekarelListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BasekarelListener) ExitInstruction(ctx *InstructionContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasekarelListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasekarelListener) ExitCondition(ctx *ConditionContext) {}

// EnterNumber is called when production number is entered.
func (s *BasekarelListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasekarelListener) ExitNumber(ctx *NumberContext) {}
