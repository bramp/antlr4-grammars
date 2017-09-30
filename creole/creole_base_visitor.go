// Generated from creole.g4 by ANTLR 4.7.

package creole // creole
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasecreoleVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasecreoleVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitMarkup(ctx *MarkupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitText(ctx *TextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitBold(ctx *BoldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitItalics(ctx *ItalicsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitHref(ctx *HrefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitImage(ctx *ImageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitHline(ctx *HlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitListitem(ctx *ListitemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitTableheader(ctx *TableheaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitTablerow(ctx *TablerowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitTitle(ctx *TitleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecreoleVisitor) VisitNowiki(ctx *NowikiContext) interface{} {
	return v.VisitChildren(ctx)
}
