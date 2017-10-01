// Generated from properties.g4 by ANTLR 4.7.

package properties // properties
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by propertiesParser.
type propertiesVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by propertiesParser#propertiesFile.
	VisitPropertiesFile(ctx *PropertiesFileContext) interface{}

	// Visit a parse tree produced by propertiesParser#row.
	VisitRow(ctx *RowContext) interface{}

	// Visit a parse tree produced by propertiesParser#decl.
	VisitDecl(ctx *DeclContext) interface{}

	// Visit a parse tree produced by propertiesParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by propertiesParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by propertiesParser#comment.
	VisitComment(ctx *CommentContext) interface{}
}
