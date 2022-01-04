// Code generated from BSL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bsl // BSL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BSLListener is a complete listener for a parse tree produced by BSLParser.
type BSLListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDefOrExpr is called when entering the defOrExpr production.
	EnterDefOrExpr(c *DefOrExprContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterTestCase is called when entering the testCase production.
	EnterTestCase(c *TestCaseContext)

	// EnterLibraryRequire is called when entering the libraryRequire production.
	EnterLibraryRequire(c *LibraryRequireContext)

	// EnterPkg is called when entering the pkg production.
	EnterPkg(c *PkgContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDefOrExpr is called when exiting the defOrExpr production.
	ExitDefOrExpr(c *DefOrExprContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitTestCase is called when exiting the testCase production.
	ExitTestCase(c *TestCaseContext)

	// ExitLibraryRequire is called when exiting the libraryRequire production.
	ExitLibraryRequire(c *LibraryRequireContext)

	// ExitPkg is called when exiting the pkg production.
	ExitPkg(c *PkgContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)
}
