// Generated from useragent.g4 by ANTLR 4.7.

package useragent // useragent
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseuseragentVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseuseragentVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseuseragentVisitor) VisitProduct(ctx *ProductContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseuseragentVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseuseragentVisitor) VisitVersion(ctx *VersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseuseragentVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}
