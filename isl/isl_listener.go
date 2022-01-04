// Code generated from ISL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package isl // ISL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ISLListener is a complete listener for a parse tree produced by ISLParser.
type ISLListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDefOrExpr is called when entering the defOrExpr production.
	EnterDefOrExpr(c *DefOrExprContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterQuoted is called when entering the quoted production.
	EnterQuoted(c *QuotedContext)

	// EnterQuasiQuoted is called when entering the quasiQuoted production.
	EnterQuasiQuoted(c *QuasiQuotedContext)

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

	// ExitQuoted is called when exiting the quoted production.
	ExitQuoted(c *QuotedContext)

	// ExitQuasiQuoted is called when exiting the quasiQuoted production.
	ExitQuasiQuoted(c *QuasiQuotedContext)

	// ExitTestCase is called when exiting the testCase production.
	ExitTestCase(c *TestCaseContext)

	// ExitLibraryRequire is called when exiting the libraryRequire production.
	ExitLibraryRequire(c *LibraryRequireContext)

	// ExitPkg is called when exiting the pkg production.
	ExitPkg(c *PkgContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)
}
