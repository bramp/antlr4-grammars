// Code generated from BSL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bsl // BSL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBSLListener is a complete listener for a parse tree produced by BSLParser.
type BaseBSLListener struct{}

var _ BSLListener = &BaseBSLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBSLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBSLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBSLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBSLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBSLListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBSLListener) ExitProgram(ctx *ProgramContext) {}

// EnterDefOrExpr is called when production defOrExpr is entered.
func (s *BaseBSLListener) EnterDefOrExpr(ctx *DefOrExprContext) {}

// ExitDefOrExpr is called when production defOrExpr is exited.
func (s *BaseBSLListener) ExitDefOrExpr(ctx *DefOrExprContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseBSLListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseBSLListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseBSLListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseBSLListener) ExitExpr(ctx *ExprContext) {}

// EnterTestCase is called when production testCase is entered.
func (s *BaseBSLListener) EnterTestCase(ctx *TestCaseContext) {}

// ExitTestCase is called when production testCase is exited.
func (s *BaseBSLListener) ExitTestCase(ctx *TestCaseContext) {}

// EnterLibraryRequire is called when production libraryRequire is entered.
func (s *BaseBSLListener) EnterLibraryRequire(ctx *LibraryRequireContext) {}

// ExitLibraryRequire is called when production libraryRequire is exited.
func (s *BaseBSLListener) ExitLibraryRequire(ctx *LibraryRequireContext) {}

// EnterPkg is called when production pkg is entered.
func (s *BaseBSLListener) EnterPkg(ctx *PkgContext) {}

// ExitPkg is called when production pkg is exited.
func (s *BaseBSLListener) ExitPkg(ctx *PkgContext) {}

// EnterName is called when production name is entered.
func (s *BaseBSLListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseBSLListener) ExitName(ctx *NameContext) {}
