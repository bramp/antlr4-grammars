// Generated from JavadocParser.g4 by ANTLR 4.7.

package javadoc // JavadocParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// JavadocParserListener is a complete listener for a parse tree produced by JavadocParser.
type JavadocParserListener interface {
	antlr.ParseTreeListener

	// EnterDocumentation is called when entering the documentation production.
	EnterDocumentation(c *DocumentationContext)

	// EnterDocumentationContent is called when entering the documentationContent production.
	EnterDocumentationContent(c *DocumentationContentContext)

	// EnterSkipWhitespace is called when entering the skipWhitespace production.
	EnterSkipWhitespace(c *SkipWhitespaceContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterDescriptionLine is called when entering the descriptionLine production.
	EnterDescriptionLine(c *DescriptionLineContext)

	// EnterDescriptionLineStart is called when entering the descriptionLineStart production.
	EnterDescriptionLineStart(c *DescriptionLineStartContext)

	// EnterDescriptionLineNoSpaceNoAt is called when entering the descriptionLineNoSpaceNoAt production.
	EnterDescriptionLineNoSpaceNoAt(c *DescriptionLineNoSpaceNoAtContext)

	// EnterDescriptionLineElement is called when entering the descriptionLineElement production.
	EnterDescriptionLineElement(c *DescriptionLineElementContext)

	// EnterDescriptionLineText is called when entering the descriptionLineText production.
	EnterDescriptionLineText(c *DescriptionLineTextContext)

	// EnterDescriptionNewline is called when entering the descriptionNewline production.
	EnterDescriptionNewline(c *DescriptionNewlineContext)

	// EnterTagSection is called when entering the tagSection production.
	EnterTagSection(c *TagSectionContext)

	// EnterBlockTag is called when entering the blockTag production.
	EnterBlockTag(c *BlockTagContext)

	// EnterBlockTagName is called when entering the blockTagName production.
	EnterBlockTagName(c *BlockTagNameContext)

	// EnterBlockTagContent is called when entering the blockTagContent production.
	EnterBlockTagContent(c *BlockTagContentContext)

	// EnterBlockTagText is called when entering the blockTagText production.
	EnterBlockTagText(c *BlockTagTextContext)

	// EnterBlockTagTextElement is called when entering the blockTagTextElement production.
	EnterBlockTagTextElement(c *BlockTagTextElementContext)

	// EnterInlineTag is called when entering the inlineTag production.
	EnterInlineTag(c *InlineTagContext)

	// EnterInlineTagName is called when entering the inlineTagName production.
	EnterInlineTagName(c *InlineTagNameContext)

	// EnterInlineTagContent is called when entering the inlineTagContent production.
	EnterInlineTagContent(c *InlineTagContentContext)

	// EnterBraceExpression is called when entering the braceExpression production.
	EnterBraceExpression(c *BraceExpressionContext)

	// EnterBraceContent is called when entering the braceContent production.
	EnterBraceContent(c *BraceContentContext)

	// EnterBraceText is called when entering the braceText production.
	EnterBraceText(c *BraceTextContext)

	// ExitDocumentation is called when exiting the documentation production.
	ExitDocumentation(c *DocumentationContext)

	// ExitDocumentationContent is called when exiting the documentationContent production.
	ExitDocumentationContent(c *DocumentationContentContext)

	// ExitSkipWhitespace is called when exiting the skipWhitespace production.
	ExitSkipWhitespace(c *SkipWhitespaceContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitDescriptionLine is called when exiting the descriptionLine production.
	ExitDescriptionLine(c *DescriptionLineContext)

	// ExitDescriptionLineStart is called when exiting the descriptionLineStart production.
	ExitDescriptionLineStart(c *DescriptionLineStartContext)

	// ExitDescriptionLineNoSpaceNoAt is called when exiting the descriptionLineNoSpaceNoAt production.
	ExitDescriptionLineNoSpaceNoAt(c *DescriptionLineNoSpaceNoAtContext)

	// ExitDescriptionLineElement is called when exiting the descriptionLineElement production.
	ExitDescriptionLineElement(c *DescriptionLineElementContext)

	// ExitDescriptionLineText is called when exiting the descriptionLineText production.
	ExitDescriptionLineText(c *DescriptionLineTextContext)

	// ExitDescriptionNewline is called when exiting the descriptionNewline production.
	ExitDescriptionNewline(c *DescriptionNewlineContext)

	// ExitTagSection is called when exiting the tagSection production.
	ExitTagSection(c *TagSectionContext)

	// ExitBlockTag is called when exiting the blockTag production.
	ExitBlockTag(c *BlockTagContext)

	// ExitBlockTagName is called when exiting the blockTagName production.
	ExitBlockTagName(c *BlockTagNameContext)

	// ExitBlockTagContent is called when exiting the blockTagContent production.
	ExitBlockTagContent(c *BlockTagContentContext)

	// ExitBlockTagText is called when exiting the blockTagText production.
	ExitBlockTagText(c *BlockTagTextContext)

	// ExitBlockTagTextElement is called when exiting the blockTagTextElement production.
	ExitBlockTagTextElement(c *BlockTagTextElementContext)

	// ExitInlineTag is called when exiting the inlineTag production.
	ExitInlineTag(c *InlineTagContext)

	// ExitInlineTagName is called when exiting the inlineTagName production.
	ExitInlineTagName(c *InlineTagNameContext)

	// ExitInlineTagContent is called when exiting the inlineTagContent production.
	ExitInlineTagContent(c *InlineTagContentContext)

	// ExitBraceExpression is called when exiting the braceExpression production.
	ExitBraceExpression(c *BraceExpressionContext)

	// ExitBraceContent is called when exiting the braceContent production.
	ExitBraceContent(c *BraceContentContext)

	// ExitBraceText is called when exiting the braceText production.
	ExitBraceText(c *BraceTextContext)
}
