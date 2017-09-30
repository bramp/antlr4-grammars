// Generated from JavadocParser.g4 by ANTLR 4.7.

package javadoc // JavadocParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJavadocParserListener is a complete listener for a parse tree produced by JavadocParser.
type BaseJavadocParserListener struct{}

var _ JavadocParserListener = &BaseJavadocParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJavadocParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJavadocParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJavadocParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJavadocParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocumentation is called when production documentation is entered.
func (s *BaseJavadocParserListener) EnterDocumentation(ctx *DocumentationContext) {}

// ExitDocumentation is called when production documentation is exited.
func (s *BaseJavadocParserListener) ExitDocumentation(ctx *DocumentationContext) {}

// EnterDocumentationContent is called when production documentationContent is entered.
func (s *BaseJavadocParserListener) EnterDocumentationContent(ctx *DocumentationContentContext) {}

// ExitDocumentationContent is called when production documentationContent is exited.
func (s *BaseJavadocParserListener) ExitDocumentationContent(ctx *DocumentationContentContext) {}

// EnterSkipWhitespace is called when production skipWhitespace is entered.
func (s *BaseJavadocParserListener) EnterSkipWhitespace(ctx *SkipWhitespaceContext) {}

// ExitSkipWhitespace is called when production skipWhitespace is exited.
func (s *BaseJavadocParserListener) ExitSkipWhitespace(ctx *SkipWhitespaceContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseJavadocParserListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseJavadocParserListener) ExitDescription(ctx *DescriptionContext) {}

// EnterDescriptionLine is called when production descriptionLine is entered.
func (s *BaseJavadocParserListener) EnterDescriptionLine(ctx *DescriptionLineContext) {}

// ExitDescriptionLine is called when production descriptionLine is exited.
func (s *BaseJavadocParserListener) ExitDescriptionLine(ctx *DescriptionLineContext) {}

// EnterDescriptionLineStart is called when production descriptionLineStart is entered.
func (s *BaseJavadocParserListener) EnterDescriptionLineStart(ctx *DescriptionLineStartContext) {}

// ExitDescriptionLineStart is called when production descriptionLineStart is exited.
func (s *BaseJavadocParserListener) ExitDescriptionLineStart(ctx *DescriptionLineStartContext) {}

// EnterDescriptionLineNoSpaceNoAt is called when production descriptionLineNoSpaceNoAt is entered.
func (s *BaseJavadocParserListener) EnterDescriptionLineNoSpaceNoAt(ctx *DescriptionLineNoSpaceNoAtContext) {
}

// ExitDescriptionLineNoSpaceNoAt is called when production descriptionLineNoSpaceNoAt is exited.
func (s *BaseJavadocParserListener) ExitDescriptionLineNoSpaceNoAt(ctx *DescriptionLineNoSpaceNoAtContext) {
}

// EnterDescriptionLineElement is called when production descriptionLineElement is entered.
func (s *BaseJavadocParserListener) EnterDescriptionLineElement(ctx *DescriptionLineElementContext) {}

// ExitDescriptionLineElement is called when production descriptionLineElement is exited.
func (s *BaseJavadocParserListener) ExitDescriptionLineElement(ctx *DescriptionLineElementContext) {}

// EnterDescriptionLineText is called when production descriptionLineText is entered.
func (s *BaseJavadocParserListener) EnterDescriptionLineText(ctx *DescriptionLineTextContext) {}

// ExitDescriptionLineText is called when production descriptionLineText is exited.
func (s *BaseJavadocParserListener) ExitDescriptionLineText(ctx *DescriptionLineTextContext) {}

// EnterDescriptionNewline is called when production descriptionNewline is entered.
func (s *BaseJavadocParserListener) EnterDescriptionNewline(ctx *DescriptionNewlineContext) {}

// ExitDescriptionNewline is called when production descriptionNewline is exited.
func (s *BaseJavadocParserListener) ExitDescriptionNewline(ctx *DescriptionNewlineContext) {}

// EnterTagSection is called when production tagSection is entered.
func (s *BaseJavadocParserListener) EnterTagSection(ctx *TagSectionContext) {}

// ExitTagSection is called when production tagSection is exited.
func (s *BaseJavadocParserListener) ExitTagSection(ctx *TagSectionContext) {}

// EnterBlockTag is called when production blockTag is entered.
func (s *BaseJavadocParserListener) EnterBlockTag(ctx *BlockTagContext) {}

// ExitBlockTag is called when production blockTag is exited.
func (s *BaseJavadocParserListener) ExitBlockTag(ctx *BlockTagContext) {}

// EnterBlockTagName is called when production blockTagName is entered.
func (s *BaseJavadocParserListener) EnterBlockTagName(ctx *BlockTagNameContext) {}

// ExitBlockTagName is called when production blockTagName is exited.
func (s *BaseJavadocParserListener) ExitBlockTagName(ctx *BlockTagNameContext) {}

// EnterBlockTagContent is called when production blockTagContent is entered.
func (s *BaseJavadocParserListener) EnterBlockTagContent(ctx *BlockTagContentContext) {}

// ExitBlockTagContent is called when production blockTagContent is exited.
func (s *BaseJavadocParserListener) ExitBlockTagContent(ctx *BlockTagContentContext) {}

// EnterBlockTagText is called when production blockTagText is entered.
func (s *BaseJavadocParserListener) EnterBlockTagText(ctx *BlockTagTextContext) {}

// ExitBlockTagText is called when production blockTagText is exited.
func (s *BaseJavadocParserListener) ExitBlockTagText(ctx *BlockTagTextContext) {}

// EnterBlockTagTextElement is called when production blockTagTextElement is entered.
func (s *BaseJavadocParserListener) EnterBlockTagTextElement(ctx *BlockTagTextElementContext) {}

// ExitBlockTagTextElement is called when production blockTagTextElement is exited.
func (s *BaseJavadocParserListener) ExitBlockTagTextElement(ctx *BlockTagTextElementContext) {}

// EnterInlineTag is called when production inlineTag is entered.
func (s *BaseJavadocParserListener) EnterInlineTag(ctx *InlineTagContext) {}

// ExitInlineTag is called when production inlineTag is exited.
func (s *BaseJavadocParserListener) ExitInlineTag(ctx *InlineTagContext) {}

// EnterInlineTagName is called when production inlineTagName is entered.
func (s *BaseJavadocParserListener) EnterInlineTagName(ctx *InlineTagNameContext) {}

// ExitInlineTagName is called when production inlineTagName is exited.
func (s *BaseJavadocParserListener) ExitInlineTagName(ctx *InlineTagNameContext) {}

// EnterInlineTagContent is called when production inlineTagContent is entered.
func (s *BaseJavadocParserListener) EnterInlineTagContent(ctx *InlineTagContentContext) {}

// ExitInlineTagContent is called when production inlineTagContent is exited.
func (s *BaseJavadocParserListener) ExitInlineTagContent(ctx *InlineTagContentContext) {}

// EnterBraceExpression is called when production braceExpression is entered.
func (s *BaseJavadocParserListener) EnterBraceExpression(ctx *BraceExpressionContext) {}

// ExitBraceExpression is called when production braceExpression is exited.
func (s *BaseJavadocParserListener) ExitBraceExpression(ctx *BraceExpressionContext) {}

// EnterBraceContent is called when production braceContent is entered.
func (s *BaseJavadocParserListener) EnterBraceContent(ctx *BraceContentContext) {}

// ExitBraceContent is called when production braceContent is exited.
func (s *BaseJavadocParserListener) ExitBraceContent(ctx *BraceContentContext) {}

// EnterBraceText is called when production braceText is entered.
func (s *BaseJavadocParserListener) EnterBraceText(ctx *BraceTextContext) {}

// ExitBraceText is called when production braceText is exited.
func (s *BaseJavadocParserListener) ExitBraceText(ctx *BraceTextContext) {}
