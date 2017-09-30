// Generated from tiny.g4 by ANTLR 4.7.

package tiny // tiny
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by tinyParser.
type tinyVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tinyParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by tinyParser#stmt_list.
	VisitStmt_list(ctx *Stmt_listContext) interface{}

	// Visit a parse tree produced by tinyParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by tinyParser#assign_stmt.
	VisitAssign_stmt(ctx *Assign_stmtContext) interface{}

	// Visit a parse tree produced by tinyParser#read_stmt.
	VisitRead_stmt(ctx *Read_stmtContext) interface{}

	// Visit a parse tree produced by tinyParser#write_stmt.
	VisitWrite_stmt(ctx *Write_stmtContext) interface{}

	// Visit a parse tree produced by tinyParser#id_list.
	VisitId_list(ctx *Id_listContext) interface{}

	// Visit a parse tree produced by tinyParser#expr_list.
	VisitExpr_list(ctx *Expr_listContext) interface{}

	// Visit a parse tree produced by tinyParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by tinyParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by tinyParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by tinyParser#op.
	VisitOp(ctx *OpContext) interface{}

	// Visit a parse tree produced by tinyParser#ident.
	VisitIdent(ctx *IdentContext) interface{}
}
