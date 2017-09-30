// Generated from creole.g4 by ANTLR 4.7.

package creole // creole
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasecreoleListener is a complete listener for a parse tree produced by creoleParser.
type BasecreoleListener struct{}

var _ creoleListener = &BasecreoleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecreoleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecreoleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecreoleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecreoleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BasecreoleListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BasecreoleListener) ExitDocument(ctx *DocumentContext) {}

// EnterLine is called when production line is entered.
func (s *BasecreoleListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BasecreoleListener) ExitLine(ctx *LineContext) {}

// EnterMarkup is called when production markup is entered.
func (s *BasecreoleListener) EnterMarkup(ctx *MarkupContext) {}

// ExitMarkup is called when production markup is exited.
func (s *BasecreoleListener) ExitMarkup(ctx *MarkupContext) {}

// EnterText is called when production text is entered.
func (s *BasecreoleListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BasecreoleListener) ExitText(ctx *TextContext) {}

// EnterBold is called when production bold is entered.
func (s *BasecreoleListener) EnterBold(ctx *BoldContext) {}

// ExitBold is called when production bold is exited.
func (s *BasecreoleListener) ExitBold(ctx *BoldContext) {}

// EnterItalics is called when production italics is entered.
func (s *BasecreoleListener) EnterItalics(ctx *ItalicsContext) {}

// ExitItalics is called when production italics is exited.
func (s *BasecreoleListener) ExitItalics(ctx *ItalicsContext) {}

// EnterHref is called when production href is entered.
func (s *BasecreoleListener) EnterHref(ctx *HrefContext) {}

// ExitHref is called when production href is exited.
func (s *BasecreoleListener) ExitHref(ctx *HrefContext) {}

// EnterImage is called when production image is entered.
func (s *BasecreoleListener) EnterImage(ctx *ImageContext) {}

// ExitImage is called when production image is exited.
func (s *BasecreoleListener) ExitImage(ctx *ImageContext) {}

// EnterHline is called when production hline is entered.
func (s *BasecreoleListener) EnterHline(ctx *HlineContext) {}

// ExitHline is called when production hline is exited.
func (s *BasecreoleListener) ExitHline(ctx *HlineContext) {}

// EnterListitem is called when production listitem is entered.
func (s *BasecreoleListener) EnterListitem(ctx *ListitemContext) {}

// ExitListitem is called when production listitem is exited.
func (s *BasecreoleListener) ExitListitem(ctx *ListitemContext) {}

// EnterTableheader is called when production tableheader is entered.
func (s *BasecreoleListener) EnterTableheader(ctx *TableheaderContext) {}

// ExitTableheader is called when production tableheader is exited.
func (s *BasecreoleListener) ExitTableheader(ctx *TableheaderContext) {}

// EnterTablerow is called when production tablerow is entered.
func (s *BasecreoleListener) EnterTablerow(ctx *TablerowContext) {}

// ExitTablerow is called when production tablerow is exited.
func (s *BasecreoleListener) ExitTablerow(ctx *TablerowContext) {}

// EnterTitle is called when production title is entered.
func (s *BasecreoleListener) EnterTitle(ctx *TitleContext) {}

// ExitTitle is called when production title is exited.
func (s *BasecreoleListener) ExitTitle(ctx *TitleContext) {}

// EnterNowiki is called when production nowiki is entered.
func (s *BasecreoleListener) EnterNowiki(ctx *NowikiContext) {}

// ExitNowiki is called when production nowiki is exited.
func (s *BasecreoleListener) ExitNowiki(ctx *NowikiContext) {}
