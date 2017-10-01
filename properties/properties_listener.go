// Generated from properties.g4 by ANTLR 4.7.

package properties // properties
import "github.com/antlr/antlr4/runtime/Go/antlr"

// propertiesListener is a complete listener for a parse tree produced by propertiesParser.
type propertiesListener interface {
	antlr.ParseTreeListener

	// EnterPropertiesFile is called when entering the propertiesFile production.
	EnterPropertiesFile(c *PropertiesFileContext)

	// EnterRow is called when entering the row production.
	EnterRow(c *RowContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitPropertiesFile is called when exiting the propertiesFile production.
	ExitPropertiesFile(c *PropertiesFileContext)

	// ExitRow is called when exiting the row production.
	ExitRow(c *RowContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
