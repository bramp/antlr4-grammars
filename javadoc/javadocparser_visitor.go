// Generated from JavadocParser.g4 by ANTLR 4.7.

package javadoc // JavadocParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by JavadocParser.
type JavadocParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JavadocParser#documentation.
	VisitDocumentation(ctx *DocumentationContext) interface{}

	// Visit a parse tree produced by JavadocParser#documentationContent.
	VisitDocumentationContent(ctx *DocumentationContentContext) interface{}

	// Visit a parse tree produced by JavadocParser#skipWhitespace.
	VisitSkipWhitespace(ctx *SkipWhitespaceContext) interface{}

	// Visit a parse tree produced by JavadocParser#description.
	VisitDescription(ctx *DescriptionContext) interface{}

	// Visit a parse tree produced by JavadocParser#descriptionLine.
	VisitDescriptionLine(ctx *DescriptionLineContext) interface{}

	// Visit a parse tree produced by JavadocParser#descriptionLineStart.
	VisitDescriptionLineStart(ctx *DescriptionLineStartContext) interface{}

	// Visit a parse tree produced by JavadocParser#descriptionLineNoSpaceNoAt.
	VisitDescriptionLineNoSpaceNoAt(ctx *DescriptionLineNoSpaceNoAtContext) interface{}

	// Visit a parse tree produced by JavadocParser#descriptionLineElement.
	VisitDescriptionLineElement(ctx *DescriptionLineElementContext) interface{}

	// Visit a parse tree produced by JavadocParser#descriptionLineText.
	VisitDescriptionLineText(ctx *DescriptionLineTextContext) interface{}

	// Visit a parse tree produced by JavadocParser#descriptionNewline.
	VisitDescriptionNewline(ctx *DescriptionNewlineContext) interface{}

	// Visit a parse tree produced by JavadocParser#tagSection.
	VisitTagSection(ctx *TagSectionContext) interface{}

	// Visit a parse tree produced by JavadocParser#blockTag.
	VisitBlockTag(ctx *BlockTagContext) interface{}

	// Visit a parse tree produced by JavadocParser#blockTagName.
	VisitBlockTagName(ctx *BlockTagNameContext) interface{}

	// Visit a parse tree produced by JavadocParser#blockTagContent.
	VisitBlockTagContent(ctx *BlockTagContentContext) interface{}

	// Visit a parse tree produced by JavadocParser#blockTagText.
	VisitBlockTagText(ctx *BlockTagTextContext) interface{}

	// Visit a parse tree produced by JavadocParser#blockTagTextElement.
	VisitBlockTagTextElement(ctx *BlockTagTextElementContext) interface{}

	// Visit a parse tree produced by JavadocParser#inlineTag.
	VisitInlineTag(ctx *InlineTagContext) interface{}

	// Visit a parse tree produced by JavadocParser#inlineTagName.
	VisitInlineTagName(ctx *InlineTagNameContext) interface{}

	// Visit a parse tree produced by JavadocParser#inlineTagContent.
	VisitInlineTagContent(ctx *InlineTagContentContext) interface{}

	// Visit a parse tree produced by JavadocParser#braceExpression.
	VisitBraceExpression(ctx *BraceExpressionContext) interface{}

	// Visit a parse tree produced by JavadocParser#braceContent.
	VisitBraceContent(ctx *BraceContentContext) interface{}

	// Visit a parse tree produced by JavadocParser#braceText.
	VisitBraceText(ctx *BraceTextContext) interface{}
}
