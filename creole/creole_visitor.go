// Generated from creole.g4 by ANTLR 4.7.

package creole // creole
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by creoleParser.
type creoleVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by creoleParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by creoleParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by creoleParser#markup.
	VisitMarkup(ctx *MarkupContext) interface{}

	// Visit a parse tree produced by creoleParser#text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by creoleParser#bold.
	VisitBold(ctx *BoldContext) interface{}

	// Visit a parse tree produced by creoleParser#italics.
	VisitItalics(ctx *ItalicsContext) interface{}

	// Visit a parse tree produced by creoleParser#href.
	VisitHref(ctx *HrefContext) interface{}

	// Visit a parse tree produced by creoleParser#image.
	VisitImage(ctx *ImageContext) interface{}

	// Visit a parse tree produced by creoleParser#hline.
	VisitHline(ctx *HlineContext) interface{}

	// Visit a parse tree produced by creoleParser#listitem.
	VisitListitem(ctx *ListitemContext) interface{}

	// Visit a parse tree produced by creoleParser#tableheader.
	VisitTableheader(ctx *TableheaderContext) interface{}

	// Visit a parse tree produced by creoleParser#tablerow.
	VisitTablerow(ctx *TablerowContext) interface{}

	// Visit a parse tree produced by creoleParser#title.
	VisitTitle(ctx *TitleContext) interface{}

	// Visit a parse tree produced by creoleParser#nowiki.
	VisitNowiki(ctx *NowikiContext) interface{}
}
