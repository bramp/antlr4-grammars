// Code generated from cayenne.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cayenne // cayenne
import "github.com/antlr/antlr4/runtime/Go/antlr"

// cayenneListener is a complete listener for a parse tree produced by cayenneParser.
type cayenneListener interface {
	antlr.ParseTreeListener

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterArm is called when entering the arm production.
	EnterArm(c *ArmContext)

	// EnterSign is called when entering the sign production.
	EnterSign(c *SignContext)

	// EnterDefn is called when entering the defn production.
	EnterDefn(c *DefnContext)

	// EnterVis is called when entering the vis production.
	EnterVis(c *VisContext)

	// EnterAbs_ is called when entering the abs_ production.
	EnterAbs_(c *Abs_Context)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterVarid is called when entering the varid production.
	EnterVarid(c *VaridContext)

	// EnterConid is called when entering the conid production.
	EnterConid(c *ConidContext)

	// EnterLblid is called when entering the lblid production.
	EnterLblid(c *LblidContext)

	// EnterId_ is called when entering the id_ production.
	EnterId_(c *Id_Context)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitArm is called when exiting the arm production.
	ExitArm(c *ArmContext)

	// ExitSign is called when exiting the sign production.
	ExitSign(c *SignContext)

	// ExitDefn is called when exiting the defn production.
	ExitDefn(c *DefnContext)

	// ExitVis is called when exiting the vis production.
	ExitVis(c *VisContext)

	// ExitAbs_ is called when exiting the abs_ production.
	ExitAbs_(c *Abs_Context)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitVarid is called when exiting the varid production.
	ExitVarid(c *VaridContext)

	// ExitConid is called when exiting the conid production.
	ExitConid(c *ConidContext)

	// ExitLblid is called when exiting the lblid production.
	ExitLblid(c *LblidContext)

	// ExitId_ is called when exiting the id_ production.
	ExitId_(c *Id_Context)
}
