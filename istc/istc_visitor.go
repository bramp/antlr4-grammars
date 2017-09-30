// Generated from istc.g4 by ANTLR 4.7.

package istc // istc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by istcParser.
type istcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by istcParser#istc.
	VisitIstc(ctx *IstcContext) interface{}

	// Visit a parse tree produced by istcParser#registration.
	VisitRegistration(ctx *RegistrationContext) interface{}

	// Visit a parse tree produced by istcParser#year.
	VisitYear(ctx *YearContext) interface{}

	// Visit a parse tree produced by istcParser#work.
	VisitWork(ctx *WorkContext) interface{}

	// Visit a parse tree produced by istcParser#check.
	VisitCheck(ctx *CheckContext) interface{}
}
