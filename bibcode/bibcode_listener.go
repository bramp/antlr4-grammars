// Code generated from bibcode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bibcode // bibcode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// bibcodeListener is a complete listener for a parse tree produced by bibcodeParser.
type bibcodeListener interface {
	antlr.ParseTreeListener

	// EnterBibcode is called when entering the bibcode production.
	EnterBibcode(c *BibcodeContext)

	// EnterYear is called when entering the year production.
	EnterYear(c *YearContext)

	// EnterPublish is called when entering the publish production.
	EnterPublish(c *PublishContext)

	// EnterVolume is called when entering the volume production.
	EnterVolume(c *VolumeContext)

	// EnterPagesection is called when entering the pagesection production.
	EnterPagesection(c *PagesectionContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// EnterPage is called when entering the page production.
	EnterPage(c *PageContext)

	// EnterAuthor is called when entering the author production.
	EnterAuthor(c *AuthorContext)

	// EnterLetter is called when entering the letter production.
	EnterLetter(c *LetterContext)

	// EnterDigit is called when entering the digit production.
	EnterDigit(c *DigitContext)

	// ExitBibcode is called when exiting the bibcode production.
	ExitBibcode(c *BibcodeContext)

	// ExitYear is called when exiting the year production.
	ExitYear(c *YearContext)

	// ExitPublish is called when exiting the publish production.
	ExitPublish(c *PublishContext)

	// ExitVolume is called when exiting the volume production.
	ExitVolume(c *VolumeContext)

	// ExitPagesection is called when exiting the pagesection production.
	ExitPagesection(c *PagesectionContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitPage is called when exiting the page production.
	ExitPage(c *PageContext)

	// ExitAuthor is called when exiting the author production.
	ExitAuthor(c *AuthorContext)

	// ExitLetter is called when exiting the letter production.
	ExitLetter(c *LetterContext)

	// ExitDigit is called when exiting the digit production.
	ExitDigit(c *DigitContext)
}
