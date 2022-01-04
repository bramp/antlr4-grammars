// Code generated from gedcom.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gedcom // gedcom
import "github.com/antlr/antlr4/runtime/Go/antlr"

// gedcomListener is a complete listener for a parse tree produced by gedcomParser.
type gedcomListener interface {
	antlr.ParseTreeListener

	// EnterGedcom is called when entering the gedcom production.
	EnterGedcom(c *GedcomContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterLevel is called when entering the level production.
	EnterLevel(c *LevelContext)

	// EnterOpt_xref_id is called when entering the opt_xref_id production.
	EnterOpt_xref_id(c *Opt_xref_idContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterLine_value is called when entering the line_value production.
	EnterLine_value(c *Line_valueContext)

	// EnterLine_item is called when entering the line_item production.
	EnterLine_item(c *Line_itemContext)

	// EnterEscape is called when entering the escape production.
	EnterEscape(c *EscapeContext)

	// EnterNon_at is called when entering the non_at production.
	EnterNon_at(c *Non_atContext)

	// EnterEscape_text is called when entering the escape_text production.
	EnterEscape_text(c *Escape_textContext)

	// EnterPointer is called when entering the pointer production.
	EnterPointer(c *PointerContext)

	// EnterPointer_string is called when entering the pointer_string production.
	EnterPointer_string(c *Pointer_stringContext)

	// EnterPointer_char is called when entering the pointer_char production.
	EnterPointer_char(c *Pointer_charContext)

	// EnterAlphanum is called when entering the alphanum production.
	EnterAlphanum(c *AlphanumContext)

	// EnterAnychar is called when entering the anychar production.
	EnterAnychar(c *AnycharContext)

	// EnterOtherchar is called when entering the otherchar production.
	EnterOtherchar(c *OthercharContext)

	// ExitGedcom is called when exiting the gedcom production.
	ExitGedcom(c *GedcomContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitLevel is called when exiting the level production.
	ExitLevel(c *LevelContext)

	// ExitOpt_xref_id is called when exiting the opt_xref_id production.
	ExitOpt_xref_id(c *Opt_xref_idContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitLine_value is called when exiting the line_value production.
	ExitLine_value(c *Line_valueContext)

	// ExitLine_item is called when exiting the line_item production.
	ExitLine_item(c *Line_itemContext)

	// ExitEscape is called when exiting the escape production.
	ExitEscape(c *EscapeContext)

	// ExitNon_at is called when exiting the non_at production.
	ExitNon_at(c *Non_atContext)

	// ExitEscape_text is called when exiting the escape_text production.
	ExitEscape_text(c *Escape_textContext)

	// ExitPointer is called when exiting the pointer production.
	ExitPointer(c *PointerContext)

	// ExitPointer_string is called when exiting the pointer_string production.
	ExitPointer_string(c *Pointer_stringContext)

	// ExitPointer_char is called when exiting the pointer_char production.
	ExitPointer_char(c *Pointer_charContext)

	// ExitAlphanum is called when exiting the alphanum production.
	ExitAlphanum(c *AlphanumContext)

	// ExitAnychar is called when exiting the anychar production.
	ExitAnychar(c *AnycharContext)

	// ExitOtherchar is called when exiting the otherchar production.
	ExitOtherchar(c *OthercharContext)
}
