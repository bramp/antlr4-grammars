// Generated from JavadocParser.g4 by ANTLR 4.7.

package javadoc // JavadocParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseJavadocParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJavadocParserVisitor) VisitDocumentation(ctx *DocumentationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitDocumentationContent(ctx *DocumentationContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitSkipWhitespace(ctx *SkipWhitespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitDescription(ctx *DescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitDescriptionLine(ctx *DescriptionLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitDescriptionLineStart(ctx *DescriptionLineStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitDescriptionLineNoSpaceNoAt(ctx *DescriptionLineNoSpaceNoAtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitDescriptionLineElement(ctx *DescriptionLineElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitDescriptionLineText(ctx *DescriptionLineTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitDescriptionNewline(ctx *DescriptionNewlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitTagSection(ctx *TagSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitBlockTag(ctx *BlockTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitBlockTagName(ctx *BlockTagNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitBlockTagContent(ctx *BlockTagContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitBlockTagText(ctx *BlockTagTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitBlockTagTextElement(ctx *BlockTagTextElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitInlineTag(ctx *InlineTagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitInlineTagName(ctx *InlineTagNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitInlineTagContent(ctx *InlineTagContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitBraceExpression(ctx *BraceExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitBraceContent(ctx *BraceContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJavadocParserVisitor) VisitBraceText(ctx *BraceTextContext) interface{} {
	return v.VisitChildren(ctx)
}
