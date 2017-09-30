// Generated from postalcode.g4 by ANTLR 4.7.

package postalcode // postalcode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by postalcodeParser.
type postalcodeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by postalcodeParser#postalcode.
	VisitPostalcode(ctx *PostalcodeContext) interface{}
}
