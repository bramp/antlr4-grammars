// Code generated from callable_.g4 by ANTLR 4.9.3. DO NOT EDIT.

package callable_ // callable_
import "github.com/antlr/antlr4/runtime/Go/antlr"

// callable_Listener is a complete listener for a parse tree produced by callable_Parser.
type callable_Listener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterF_inner is called when entering the f_inner production.
	EnterF_inner(c *F_innerContext)

	// EnterF_arg is called when entering the f_arg production.
	EnterF_arg(c *F_argContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitF_inner is called when exiting the f_inner production.
	ExitF_inner(c *F_innerContext)

	// ExitF_arg is called when exiting the f_arg production.
	ExitF_arg(c *F_argContext)
}
