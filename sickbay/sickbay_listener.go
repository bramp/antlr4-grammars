// Code generated from sickbay.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sickbay // sickbay
import "github.com/antlr/antlr4/runtime/Go/antlr"

// sickbayListener is a complete listener for a parse tree produced by sickbayParser.
type sickbayListener interface {
	antlr.ParseTreeListener

	// EnterSickbay is called when entering the sickbay production.
	EnterSickbay(c *SickbayContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterIntExpr is called when entering the intExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterIntVar is called when entering the intVar production.
	EnterIntVar(c *IntVarContext)

	// ExitSickbay is called when exiting the sickbay production.
	ExitSickbay(c *SickbayContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitIntExpr is called when exiting the intExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitIntVar is called when exiting the intVar production.
	ExitIntVar(c *IntVarContext)
}
