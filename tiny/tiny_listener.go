// Generated from tiny.g4 by ANTLR 4.7.

package tiny // tiny
import "github.com/antlr/antlr4/runtime/Go/antlr"

// tinyListener is a complete listener for a parse tree produced by tinyParser.
type tinyListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStmt_list is called when entering the stmt_list production.
	EnterStmt_list(c *Stmt_listContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssign_stmt is called when entering the assign_stmt production.
	EnterAssign_stmt(c *Assign_stmtContext)

	// EnterRead_stmt is called when entering the read_stmt production.
	EnterRead_stmt(c *Read_stmtContext)

	// EnterWrite_stmt is called when entering the write_stmt production.
	EnterWrite_stmt(c *Write_stmtContext)

	// EnterId_list is called when entering the id_list production.
	EnterId_list(c *Id_listContext)

	// EnterExpr_list is called when entering the expr_list production.
	EnterExpr_list(c *Expr_listContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStmt_list is called when exiting the stmt_list production.
	ExitStmt_list(c *Stmt_listContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssign_stmt is called when exiting the assign_stmt production.
	ExitAssign_stmt(c *Assign_stmtContext)

	// ExitRead_stmt is called when exiting the read_stmt production.
	ExitRead_stmt(c *Read_stmtContext)

	// ExitWrite_stmt is called when exiting the write_stmt production.
	ExitWrite_stmt(c *Write_stmtContext)

	// ExitId_list is called when exiting the id_list production.
	ExitId_list(c *Id_listContext)

	// ExitExpr_list is called when exiting the expr_list production.
	ExitExpr_list(c *Expr_listContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)
}
