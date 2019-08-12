// Code generated from ReStructuredText.g4 by ANTLR 4.7.2. DO NOT EDIT.

package restructuredtext // ReStructuredText
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ReStructuredTextListener is a complete listener for a parse tree produced by ReStructuredTextParser.
type ReStructuredTextListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterSectionElement is called when entering the sectionElement production.
	EnterSectionElement(c *SectionElementContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterCommentParagraphs is called when entering the commentParagraphs production.
	EnterCommentParagraphs(c *CommentParagraphsContext)

	// EnterCommentRest is called when entering the commentRest production.
	EnterCommentRest(c *CommentRestContext)

	// EnterCommentParagraph is called when entering the commentParagraph production.
	EnterCommentParagraph(c *CommentParagraphContext)

	// EnterCommentLineNoBreak is called when entering the commentLineNoBreak production.
	EnterCommentLineNoBreak(c *CommentLineNoBreakContext)

	// EnterCommentLine is called when entering the commentLine production.
	EnterCommentLine(c *CommentLineContext)

	// EnterCommentLineAtoms is called when entering the commentLineAtoms production.
	EnterCommentLineAtoms(c *CommentLineAtomsContext)

	// EnterParagraph is called when entering the paragraph production.
	EnterParagraph(c *ParagraphContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// EnterTitle is called when entering the title production.
	EnterTitle(c *TitleContext)

	// EnterLineBlock is called when entering the lineBlock production.
	EnterLineBlock(c *LineBlockContext)

	// EnterLineBlockLine is called when entering the lineBlockLine production.
	EnterLineBlockLine(c *LineBlockLineContext)

	// EnterListItemBullet is called when entering the listItemBullet production.
	EnterListItemBullet(c *ListItemBulletContext)

	// EnterBulletCrossLine is called when entering the bulletCrossLine production.
	EnterBulletCrossLine(c *BulletCrossLineContext)

	// EnterBulletSimple is called when entering the bulletSimple production.
	EnterBulletSimple(c *BulletSimpleContext)

	// EnterBullet is called when entering the bullet production.
	EnterBullet(c *BulletContext)

	// EnterListItemEnumerated is called when entering the listItemEnumerated production.
	EnterListItemEnumerated(c *ListItemEnumeratedContext)

	// EnterParagraphNoBreak is called when entering the paragraphNoBreak production.
	EnterParagraphNoBreak(c *ParagraphNoBreakContext)

	// EnterLineNoBreak is called when entering the lineNoBreak production.
	EnterLineNoBreak(c *LineNoBreakContext)

	// EnterLines is called when entering the lines production.
	EnterLines(c *LinesContext)

	// EnterLinesNormal is called when entering the linesNormal production.
	EnterLinesNormal(c *LinesNormalContext)

	// EnterLinesStar is called when entering the linesStar production.
	EnterLinesStar(c *LinesStarContext)

	// EnterLineNormal is called when entering the lineNormal production.
	EnterLineNormal(c *LineNormalContext)

	// EnterLineStar is called when entering the lineStar production.
	EnterLineStar(c *LineStarContext)

	// EnterLineSpecial is called when entering the lineSpecial production.
	EnterLineSpecial(c *LineSpecialContext)

	// EnterEmpty_line is called when entering the empty_line production.
	EnterEmpty_line(c *Empty_lineContext)

	// EnterIndentation is called when entering the indentation production.
	EnterIndentation(c *IndentationContext)

	// EnterSpanLineStartNoStar is called when entering the spanLineStartNoStar production.
	EnterSpanLineStartNoStar(c *SpanLineStartNoStarContext)

	// EnterTextLineStart is called when entering the textLineStart production.
	EnterTextLineStart(c *TextLineStartContext)

	// EnterLineStart_fragment is called when entering the lineStart_fragment production.
	EnterLineStart_fragment(c *LineStart_fragmentContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterTextStart is called when entering the textStart production.
	EnterTextStart(c *TextStartContext)

	// EnterForcedText is called when entering the forcedText production.
	EnterForcedText(c *ForcedTextContext)

	// EnterSpanNoStar is called when entering the spanNoStar production.
	EnterSpanNoStar(c *SpanNoStarContext)

	// EnterSpan is called when entering the span production.
	EnterSpan(c *SpanContext)

	// EnterQuotedLiteral is called when entering the quotedLiteral production.
	EnterQuotedLiteral(c *QuotedLiteralContext)

	// EnterText_fragment_start is called when entering the text_fragment_start production.
	EnterText_fragment_start(c *Text_fragment_startContext)

	// EnterText_fragment is called when entering the text_fragment production.
	EnterText_fragment(c *Text_fragmentContext)

	// EnterStarText is called when entering the starText production.
	EnterStarText(c *StarTextContext)

	// EnterStarAtoms is called when entering the starAtoms production.
	EnterStarAtoms(c *StarAtomsContext)

	// EnterStarNoSpace is called when entering the starNoSpace production.
	EnterStarNoSpace(c *StarNoSpaceContext)

	// EnterStarAtom is called when entering the starAtom production.
	EnterStarAtom(c *StarAtomContext)

	// EnterBackTickText is called when entering the backTickText production.
	EnterBackTickText(c *BackTickTextContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterBackTickAtoms is called when entering the backTickAtoms production.
	EnterBackTickAtoms(c *BackTickAtomsContext)

	// EnterBackTickNoSpace is called when entering the backTickNoSpace production.
	EnterBackTickNoSpace(c *BackTickNoSpaceContext)

	// EnterBackTickAtom is called when entering the backTickAtom production.
	EnterBackTickAtom(c *BackTickAtomContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterReferenceIn is called when entering the referenceIn production.
	EnterReferenceIn(c *ReferenceInContext)

	// EnterHyperlinkTarget is called when entering the hyperlinkTarget production.
	EnterHyperlinkTarget(c *HyperlinkTargetContext)

	// EnterHyperlink is called when entering the hyperlink production.
	EnterHyperlink(c *HyperlinkContext)

	// EnterHyperlinkDoc is called when entering the hyperlinkDoc production.
	EnterHyperlinkDoc(c *HyperlinkDocContext)

	// EnterUrl is called when entering the url production.
	EnterUrl(c *UrlContext)

	// EnterUrlAtom is called when entering the urlAtom production.
	EnterUrlAtom(c *UrlAtomContext)

	// EnterHyperlinkAtom is called when entering the hyperlinkAtom production.
	EnterHyperlinkAtom(c *HyperlinkAtomContext)

	// EnterSeparator is called when entering the separator production.
	EnterSeparator(c *SeparatorContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitSectionElement is called when exiting the sectionElement production.
	ExitSectionElement(c *SectionElementContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitCommentParagraphs is called when exiting the commentParagraphs production.
	ExitCommentParagraphs(c *CommentParagraphsContext)

	// ExitCommentRest is called when exiting the commentRest production.
	ExitCommentRest(c *CommentRestContext)

	// ExitCommentParagraph is called when exiting the commentParagraph production.
	ExitCommentParagraph(c *CommentParagraphContext)

	// ExitCommentLineNoBreak is called when exiting the commentLineNoBreak production.
	ExitCommentLineNoBreak(c *CommentLineNoBreakContext)

	// ExitCommentLine is called when exiting the commentLine production.
	ExitCommentLine(c *CommentLineContext)

	// ExitCommentLineAtoms is called when exiting the commentLineAtoms production.
	ExitCommentLineAtoms(c *CommentLineAtomsContext)

	// ExitParagraph is called when exiting the paragraph production.
	ExitParagraph(c *ParagraphContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitTitle is called when exiting the title production.
	ExitTitle(c *TitleContext)

	// ExitLineBlock is called when exiting the lineBlock production.
	ExitLineBlock(c *LineBlockContext)

	// ExitLineBlockLine is called when exiting the lineBlockLine production.
	ExitLineBlockLine(c *LineBlockLineContext)

	// ExitListItemBullet is called when exiting the listItemBullet production.
	ExitListItemBullet(c *ListItemBulletContext)

	// ExitBulletCrossLine is called when exiting the bulletCrossLine production.
	ExitBulletCrossLine(c *BulletCrossLineContext)

	// ExitBulletSimple is called when exiting the bulletSimple production.
	ExitBulletSimple(c *BulletSimpleContext)

	// ExitBullet is called when exiting the bullet production.
	ExitBullet(c *BulletContext)

	// ExitListItemEnumerated is called when exiting the listItemEnumerated production.
	ExitListItemEnumerated(c *ListItemEnumeratedContext)

	// ExitParagraphNoBreak is called when exiting the paragraphNoBreak production.
	ExitParagraphNoBreak(c *ParagraphNoBreakContext)

	// ExitLineNoBreak is called when exiting the lineNoBreak production.
	ExitLineNoBreak(c *LineNoBreakContext)

	// ExitLines is called when exiting the lines production.
	ExitLines(c *LinesContext)

	// ExitLinesNormal is called when exiting the linesNormal production.
	ExitLinesNormal(c *LinesNormalContext)

	// ExitLinesStar is called when exiting the linesStar production.
	ExitLinesStar(c *LinesStarContext)

	// ExitLineNormal is called when exiting the lineNormal production.
	ExitLineNormal(c *LineNormalContext)

	// ExitLineStar is called when exiting the lineStar production.
	ExitLineStar(c *LineStarContext)

	// ExitLineSpecial is called when exiting the lineSpecial production.
	ExitLineSpecial(c *LineSpecialContext)

	// ExitEmpty_line is called when exiting the empty_line production.
	ExitEmpty_line(c *Empty_lineContext)

	// ExitIndentation is called when exiting the indentation production.
	ExitIndentation(c *IndentationContext)

	// ExitSpanLineStartNoStar is called when exiting the spanLineStartNoStar production.
	ExitSpanLineStartNoStar(c *SpanLineStartNoStarContext)

	// ExitTextLineStart is called when exiting the textLineStart production.
	ExitTextLineStart(c *TextLineStartContext)

	// ExitLineStart_fragment is called when exiting the lineStart_fragment production.
	ExitLineStart_fragment(c *LineStart_fragmentContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitTextStart is called when exiting the textStart production.
	ExitTextStart(c *TextStartContext)

	// ExitForcedText is called when exiting the forcedText production.
	ExitForcedText(c *ForcedTextContext)

	// ExitSpanNoStar is called when exiting the spanNoStar production.
	ExitSpanNoStar(c *SpanNoStarContext)

	// ExitSpan is called when exiting the span production.
	ExitSpan(c *SpanContext)

	// ExitQuotedLiteral is called when exiting the quotedLiteral production.
	ExitQuotedLiteral(c *QuotedLiteralContext)

	// ExitText_fragment_start is called when exiting the text_fragment_start production.
	ExitText_fragment_start(c *Text_fragment_startContext)

	// ExitText_fragment is called when exiting the text_fragment production.
	ExitText_fragment(c *Text_fragmentContext)

	// ExitStarText is called when exiting the starText production.
	ExitStarText(c *StarTextContext)

	// ExitStarAtoms is called when exiting the starAtoms production.
	ExitStarAtoms(c *StarAtomsContext)

	// ExitStarNoSpace is called when exiting the starNoSpace production.
	ExitStarNoSpace(c *StarNoSpaceContext)

	// ExitStarAtom is called when exiting the starAtom production.
	ExitStarAtom(c *StarAtomContext)

	// ExitBackTickText is called when exiting the backTickText production.
	ExitBackTickText(c *BackTickTextContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitBackTickAtoms is called when exiting the backTickAtoms production.
	ExitBackTickAtoms(c *BackTickAtomsContext)

	// ExitBackTickNoSpace is called when exiting the backTickNoSpace production.
	ExitBackTickNoSpace(c *BackTickNoSpaceContext)

	// ExitBackTickAtom is called when exiting the backTickAtom production.
	ExitBackTickAtom(c *BackTickAtomContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitReferenceIn is called when exiting the referenceIn production.
	ExitReferenceIn(c *ReferenceInContext)

	// ExitHyperlinkTarget is called when exiting the hyperlinkTarget production.
	ExitHyperlinkTarget(c *HyperlinkTargetContext)

	// ExitHyperlink is called when exiting the hyperlink production.
	ExitHyperlink(c *HyperlinkContext)

	// ExitHyperlinkDoc is called when exiting the hyperlinkDoc production.
	ExitHyperlinkDoc(c *HyperlinkDocContext)

	// ExitUrl is called when exiting the url production.
	ExitUrl(c *UrlContext)

	// ExitUrlAtom is called when exiting the urlAtom production.
	ExitUrlAtom(c *UrlAtomContext)

	// ExitHyperlinkAtom is called when exiting the hyperlinkAtom production.
	ExitHyperlinkAtom(c *HyperlinkAtomContext)

	// ExitSeparator is called when exiting the separator production.
	ExitSeparator(c *SeparatorContext)
}
