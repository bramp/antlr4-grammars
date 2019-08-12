// Code generated from CMake.g4 by ANTLR 4.7.2. DO NOT EDIT.

package cmake // CMake
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CMakeListener is a complete listener for a parse tree produced by CMakeParser.
type CMakeListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterCommand_invocation is called when entering the command_invocation production.
	EnterCommand_invocation(c *Command_invocationContext)

	// EnterSingle_argument is called when entering the single_argument production.
	EnterSingle_argument(c *Single_argumentContext)

	// EnterCompound_argument is called when entering the compound_argument production.
	EnterCompound_argument(c *Compound_argumentContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitCommand_invocation is called when exiting the command_invocation production.
	ExitCommand_invocation(c *Command_invocationContext)

	// ExitSingle_argument is called when exiting the single_argument production.
	ExitSingle_argument(c *Single_argumentContext)

	// ExitCompound_argument is called when exiting the compound_argument production.
	ExitCompound_argument(c *Compound_argumentContext)
}
