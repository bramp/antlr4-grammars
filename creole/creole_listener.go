// Code generated from creole.g4 by ANTLR 4.9.3. DO NOT EDIT.

package creole // creole
import "github.com/antlr/antlr4/runtime/Go/antlr"

// creoleListener is a complete listener for a parse tree produced by creoleParser.
type creoleListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterMarkup is called when entering the markup production.
	EnterMarkup(c *MarkupContext)

	// EnterText_ is called when entering the text_ production.
	EnterText_(c *Text_Context)

	// EnterBold is called when entering the bold production.
	EnterBold(c *BoldContext)

	// EnterItalics is called when entering the italics production.
	EnterItalics(c *ItalicsContext)

	// EnterHref is called when entering the href production.
	EnterHref(c *HrefContext)

	// EnterImage is called when entering the image production.
	EnterImage(c *ImageContext)

	// EnterHline is called when entering the hline production.
	EnterHline(c *HlineContext)

	// EnterListitem is called when entering the listitem production.
	EnterListitem(c *ListitemContext)

	// EnterTableheader is called when entering the tableheader production.
	EnterTableheader(c *TableheaderContext)

	// EnterTablerow is called when entering the tablerow production.
	EnterTablerow(c *TablerowContext)

	// EnterTitle is called when entering the title production.
	EnterTitle(c *TitleContext)

	// EnterNowiki is called when entering the nowiki production.
	EnterNowiki(c *NowikiContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitMarkup is called when exiting the markup production.
	ExitMarkup(c *MarkupContext)

	// ExitText_ is called when exiting the text_ production.
	ExitText_(c *Text_Context)

	// ExitBold is called when exiting the bold production.
	ExitBold(c *BoldContext)

	// ExitItalics is called when exiting the italics production.
	ExitItalics(c *ItalicsContext)

	// ExitHref is called when exiting the href production.
	ExitHref(c *HrefContext)

	// ExitImage is called when exiting the image production.
	ExitImage(c *ImageContext)

	// ExitHline is called when exiting the hline production.
	ExitHline(c *HlineContext)

	// ExitListitem is called when exiting the listitem production.
	ExitListitem(c *ListitemContext)

	// ExitTableheader is called when exiting the tableheader production.
	ExitTableheader(c *TableheaderContext)

	// ExitTablerow is called when exiting the tablerow production.
	ExitTablerow(c *TablerowContext)

	// ExitTitle is called when exiting the title production.
	ExitTitle(c *TitleContext)

	// ExitNowiki is called when exiting the nowiki production.
	ExitNowiki(c *NowikiContext)
}
