// Code generated from scotty.g4 by ANTLR 4.9.3. DO NOT EDIT.

package scotty // scotty
import "github.com/antlr/antlr4/runtime/Go/antlr"

// scottyListener is a complete listener for a parse tree produced by scottyParser.
type scottyListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterProgram_lines is called when entering the program_lines production.
	EnterProgram_lines(c *Program_linesContext)

	// EnterVar_assign is called when entering the var_assign production.
	EnterVar_assign(c *Var_assignContext)

	// EnterFn_def is called when entering the fn_def production.
	EnterFn_def(c *Fn_defContext)

	// EnterPrefix_exp is called when entering the prefix_exp production.
	EnterPrefix_exp(c *Prefix_expContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterId_tail is called when entering the id_tail production.
	EnterId_tail(c *Id_tailContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterDigits is called when entering the digits production.
	EnterDigits(c *DigitsContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitProgram_lines is called when exiting the program_lines production.
	ExitProgram_lines(c *Program_linesContext)

	// ExitVar_assign is called when exiting the var_assign production.
	ExitVar_assign(c *Var_assignContext)

	// ExitFn_def is called when exiting the fn_def production.
	ExitFn_def(c *Fn_defContext)

	// ExitPrefix_exp is called when exiting the prefix_exp production.
	ExitPrefix_exp(c *Prefix_expContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitId_tail is called when exiting the id_tail production.
	ExitId_tail(c *Id_tailContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitDigits is called when exiting the digits production.
	ExitDigits(c *DigitsContext)
}
