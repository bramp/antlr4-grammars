// Code generated from karel.g4 by ANTLR 4.9.3. DO NOT EDIT.

package karel // karel
import "github.com/antlr/antlr4/runtime/Go/antlr"

// karelListener is a complete listener for a parse tree produced by karelParser.
type karelListener interface {
	antlr.ParseTreeListener

	// EnterKarel is called when entering the karel production.
	EnterKarel(c *KarelContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterIteration is called when entering the iteration production.
	EnterIteration(c *IterationContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitKarel is called when exiting the karel production.
	ExitKarel(c *KarelContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitIteration is called when exiting the iteration production.
	ExitIteration(c *IterationContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
