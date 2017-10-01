// Generated from properties.g4 by ANTLR 4.7.

package properties // properties
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasepropertiesVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasepropertiesVisitor) VisitPropertiesFile(ctx *PropertiesFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropertiesVisitor) VisitRow(ctx *RowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropertiesVisitor) VisitDecl(ctx *DeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropertiesVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropertiesVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepropertiesVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}
