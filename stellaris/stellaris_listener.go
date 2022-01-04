// Code generated from stellaris.g4 by ANTLR 4.9.3. DO NOT EDIT.

package stellaris // stellaris
import "github.com/antlr/antlr4/runtime/Go/antlr"

// stellarisListener is a complete listener for a parse tree produced by stellarisParser.
type stellarisListener interface {
	antlr.ParseTreeListener

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterKeyval is called when entering the keyval production.
	EnterKeyval(c *KeyvalContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterVal is called when entering the val production.
	EnterVal(c *ValContext)

	// EnterAttrib is called when entering the attrib production.
	EnterAttrib(c *AttribContext)

	// EnterAccessor is called when entering the accessor production.
	EnterAccessor(c *AccessorContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterId_ is called when entering the id_ production.
	EnterId_(c *Id_Context)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitKeyval is called when exiting the keyval production.
	ExitKeyval(c *KeyvalContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitVal is called when exiting the val production.
	ExitVal(c *ValContext)

	// ExitAttrib is called when exiting the attrib production.
	ExitAttrib(c *AttribContext)

	// ExitAccessor is called when exiting the accessor production.
	ExitAccessor(c *AccessorContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitId_ is called when exiting the id_ production.
	ExitId_(c *Id_Context)
}
