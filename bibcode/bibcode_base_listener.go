// Code generated from bibcode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bibcode // bibcode
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasebibcodeListener is a complete listener for a parse tree produced by bibcodeParser.
type BasebibcodeListener struct{}

var _ bibcodeListener = &BasebibcodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebibcodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebibcodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebibcodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebibcodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBibcode is called when production bibcode is entered.
func (s *BasebibcodeListener) EnterBibcode(ctx *BibcodeContext) {}

// ExitBibcode is called when production bibcode is exited.
func (s *BasebibcodeListener) ExitBibcode(ctx *BibcodeContext) {}

// EnterYear is called when production year is entered.
func (s *BasebibcodeListener) EnterYear(ctx *YearContext) {}

// ExitYear is called when production year is exited.
func (s *BasebibcodeListener) ExitYear(ctx *YearContext) {}

// EnterPublish is called when production publish is entered.
func (s *BasebibcodeListener) EnterPublish(ctx *PublishContext) {}

// ExitPublish is called when production publish is exited.
func (s *BasebibcodeListener) ExitPublish(ctx *PublishContext) {}

// EnterVolume is called when production volume is entered.
func (s *BasebibcodeListener) EnterVolume(ctx *VolumeContext) {}

// ExitVolume is called when production volume is exited.
func (s *BasebibcodeListener) ExitVolume(ctx *VolumeContext) {}

// EnterPagesection is called when production pagesection is entered.
func (s *BasebibcodeListener) EnterPagesection(ctx *PagesectionContext) {}

// ExitPagesection is called when production pagesection is exited.
func (s *BasebibcodeListener) ExitPagesection(ctx *PagesectionContext) {}

// EnterSection is called when production section is entered.
func (s *BasebibcodeListener) EnterSection(ctx *SectionContext) {}

// ExitSection is called when production section is exited.
func (s *BasebibcodeListener) ExitSection(ctx *SectionContext) {}

// EnterPage is called when production page is entered.
func (s *BasebibcodeListener) EnterPage(ctx *PageContext) {}

// ExitPage is called when production page is exited.
func (s *BasebibcodeListener) ExitPage(ctx *PageContext) {}

// EnterAuthor is called when production author is entered.
func (s *BasebibcodeListener) EnterAuthor(ctx *AuthorContext) {}

// ExitAuthor is called when production author is exited.
func (s *BasebibcodeListener) ExitAuthor(ctx *AuthorContext) {}

// EnterLetter is called when production letter is entered.
func (s *BasebibcodeListener) EnterLetter(ctx *LetterContext) {}

// ExitLetter is called when production letter is exited.
func (s *BasebibcodeListener) ExitLetter(ctx *LetterContext) {}

// EnterDigit is called when production digit is entered.
func (s *BasebibcodeListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production digit is exited.
func (s *BasebibcodeListener) ExitDigit(ctx *DigitContext) {}
