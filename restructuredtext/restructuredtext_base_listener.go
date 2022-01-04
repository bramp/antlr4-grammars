// Code generated from ReStructuredText.g4 by ANTLR 4.9.3. DO NOT EDIT.

package restructuredtext // ReStructuredText
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseReStructuredTextListener is a complete listener for a parse tree produced by ReStructuredTextParser.
type BaseReStructuredTextListener struct{}

var _ ReStructuredTextListener = &BaseReStructuredTextListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseReStructuredTextListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseReStructuredTextListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseReStructuredTextListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseReStructuredTextListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseReStructuredTextListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseReStructuredTextListener) ExitParse(ctx *ParseContext) {}

// EnterElement is called when production element is entered.
func (s *BaseReStructuredTextListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseReStructuredTextListener) ExitElement(ctx *ElementContext) {}

// EnterSectionElement is called when production sectionElement is entered.
func (s *BaseReStructuredTextListener) EnterSectionElement(ctx *SectionElementContext) {}

// ExitSectionElement is called when production sectionElement is exited.
func (s *BaseReStructuredTextListener) ExitSectionElement(ctx *SectionElementContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseReStructuredTextListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseReStructuredTextListener) ExitComment(ctx *CommentContext) {}

// EnterCommentParagraphs is called when production commentParagraphs is entered.
func (s *BaseReStructuredTextListener) EnterCommentParagraphs(ctx *CommentParagraphsContext) {}

// ExitCommentParagraphs is called when production commentParagraphs is exited.
func (s *BaseReStructuredTextListener) ExitCommentParagraphs(ctx *CommentParagraphsContext) {}

// EnterCommentRest is called when production commentRest is entered.
func (s *BaseReStructuredTextListener) EnterCommentRest(ctx *CommentRestContext) {}

// ExitCommentRest is called when production commentRest is exited.
func (s *BaseReStructuredTextListener) ExitCommentRest(ctx *CommentRestContext) {}

// EnterCommentParagraph is called when production commentParagraph is entered.
func (s *BaseReStructuredTextListener) EnterCommentParagraph(ctx *CommentParagraphContext) {}

// ExitCommentParagraph is called when production commentParagraph is exited.
func (s *BaseReStructuredTextListener) ExitCommentParagraph(ctx *CommentParagraphContext) {}

// EnterCommentLineNoBreak is called when production commentLineNoBreak is entered.
func (s *BaseReStructuredTextListener) EnterCommentLineNoBreak(ctx *CommentLineNoBreakContext) {}

// ExitCommentLineNoBreak is called when production commentLineNoBreak is exited.
func (s *BaseReStructuredTextListener) ExitCommentLineNoBreak(ctx *CommentLineNoBreakContext) {}

// EnterCommentLine is called when production commentLine is entered.
func (s *BaseReStructuredTextListener) EnterCommentLine(ctx *CommentLineContext) {}

// ExitCommentLine is called when production commentLine is exited.
func (s *BaseReStructuredTextListener) ExitCommentLine(ctx *CommentLineContext) {}

// EnterCommentLineAtoms is called when production commentLineAtoms is entered.
func (s *BaseReStructuredTextListener) EnterCommentLineAtoms(ctx *CommentLineAtomsContext) {}

// ExitCommentLineAtoms is called when production commentLineAtoms is exited.
func (s *BaseReStructuredTextListener) ExitCommentLineAtoms(ctx *CommentLineAtomsContext) {}

// EnterParagraph is called when production paragraph is entered.
func (s *BaseReStructuredTextListener) EnterParagraph(ctx *ParagraphContext) {}

// ExitParagraph is called when production paragraph is exited.
func (s *BaseReStructuredTextListener) ExitParagraph(ctx *ParagraphContext) {}

// EnterSection is called when production section is entered.
func (s *BaseReStructuredTextListener) EnterSection(ctx *SectionContext) {}

// ExitSection is called when production section is exited.
func (s *BaseReStructuredTextListener) ExitSection(ctx *SectionContext) {}

// EnterTitle is called when production title is entered.
func (s *BaseReStructuredTextListener) EnterTitle(ctx *TitleContext) {}

// ExitTitle is called when production title is exited.
func (s *BaseReStructuredTextListener) ExitTitle(ctx *TitleContext) {}

// EnterLineBlock is called when production lineBlock is entered.
func (s *BaseReStructuredTextListener) EnterLineBlock(ctx *LineBlockContext) {}

// ExitLineBlock is called when production lineBlock is exited.
func (s *BaseReStructuredTextListener) ExitLineBlock(ctx *LineBlockContext) {}

// EnterLineBlockLine is called when production lineBlockLine is entered.
func (s *BaseReStructuredTextListener) EnterLineBlockLine(ctx *LineBlockLineContext) {}

// ExitLineBlockLine is called when production lineBlockLine is exited.
func (s *BaseReStructuredTextListener) ExitLineBlockLine(ctx *LineBlockLineContext) {}

// EnterListItemBullet is called when production listItemBullet is entered.
func (s *BaseReStructuredTextListener) EnterListItemBullet(ctx *ListItemBulletContext) {}

// ExitListItemBullet is called when production listItemBullet is exited.
func (s *BaseReStructuredTextListener) ExitListItemBullet(ctx *ListItemBulletContext) {}

// EnterBulletCrossLine is called when production bulletCrossLine is entered.
func (s *BaseReStructuredTextListener) EnterBulletCrossLine(ctx *BulletCrossLineContext) {}

// ExitBulletCrossLine is called when production bulletCrossLine is exited.
func (s *BaseReStructuredTextListener) ExitBulletCrossLine(ctx *BulletCrossLineContext) {}

// EnterBulletSimple is called when production bulletSimple is entered.
func (s *BaseReStructuredTextListener) EnterBulletSimple(ctx *BulletSimpleContext) {}

// ExitBulletSimple is called when production bulletSimple is exited.
func (s *BaseReStructuredTextListener) ExitBulletSimple(ctx *BulletSimpleContext) {}

// EnterBullet is called when production bullet is entered.
func (s *BaseReStructuredTextListener) EnterBullet(ctx *BulletContext) {}

// ExitBullet is called when production bullet is exited.
func (s *BaseReStructuredTextListener) ExitBullet(ctx *BulletContext) {}

// EnterListItemEnumerated is called when production listItemEnumerated is entered.
func (s *BaseReStructuredTextListener) EnterListItemEnumerated(ctx *ListItemEnumeratedContext) {}

// ExitListItemEnumerated is called when production listItemEnumerated is exited.
func (s *BaseReStructuredTextListener) ExitListItemEnumerated(ctx *ListItemEnumeratedContext) {}

// EnterParagraphNoBreak is called when production paragraphNoBreak is entered.
func (s *BaseReStructuredTextListener) EnterParagraphNoBreak(ctx *ParagraphNoBreakContext) {}

// ExitParagraphNoBreak is called when production paragraphNoBreak is exited.
func (s *BaseReStructuredTextListener) ExitParagraphNoBreak(ctx *ParagraphNoBreakContext) {}

// EnterLineNoBreak is called when production lineNoBreak is entered.
func (s *BaseReStructuredTextListener) EnterLineNoBreak(ctx *LineNoBreakContext) {}

// ExitLineNoBreak is called when production lineNoBreak is exited.
func (s *BaseReStructuredTextListener) ExitLineNoBreak(ctx *LineNoBreakContext) {}

// EnterLines is called when production lines is entered.
func (s *BaseReStructuredTextListener) EnterLines(ctx *LinesContext) {}

// ExitLines is called when production lines is exited.
func (s *BaseReStructuredTextListener) ExitLines(ctx *LinesContext) {}

// EnterLinesNormal is called when production linesNormal is entered.
func (s *BaseReStructuredTextListener) EnterLinesNormal(ctx *LinesNormalContext) {}

// ExitLinesNormal is called when production linesNormal is exited.
func (s *BaseReStructuredTextListener) ExitLinesNormal(ctx *LinesNormalContext) {}

// EnterLinesStar is called when production linesStar is entered.
func (s *BaseReStructuredTextListener) EnterLinesStar(ctx *LinesStarContext) {}

// ExitLinesStar is called when production linesStar is exited.
func (s *BaseReStructuredTextListener) ExitLinesStar(ctx *LinesStarContext) {}

// EnterLineNormal is called when production lineNormal is entered.
func (s *BaseReStructuredTextListener) EnterLineNormal(ctx *LineNormalContext) {}

// ExitLineNormal is called when production lineNormal is exited.
func (s *BaseReStructuredTextListener) ExitLineNormal(ctx *LineNormalContext) {}

// EnterLineStar is called when production lineStar is entered.
func (s *BaseReStructuredTextListener) EnterLineStar(ctx *LineStarContext) {}

// ExitLineStar is called when production lineStar is exited.
func (s *BaseReStructuredTextListener) ExitLineStar(ctx *LineStarContext) {}

// EnterLineSpecial is called when production lineSpecial is entered.
func (s *BaseReStructuredTextListener) EnterLineSpecial(ctx *LineSpecialContext) {}

// ExitLineSpecial is called when production lineSpecial is exited.
func (s *BaseReStructuredTextListener) ExitLineSpecial(ctx *LineSpecialContext) {}

// EnterEmpty_line is called when production empty_line is entered.
func (s *BaseReStructuredTextListener) EnterEmpty_line(ctx *Empty_lineContext) {}

// ExitEmpty_line is called when production empty_line is exited.
func (s *BaseReStructuredTextListener) ExitEmpty_line(ctx *Empty_lineContext) {}

// EnterIndentation is called when production indentation is entered.
func (s *BaseReStructuredTextListener) EnterIndentation(ctx *IndentationContext) {}

// ExitIndentation is called when production indentation is exited.
func (s *BaseReStructuredTextListener) ExitIndentation(ctx *IndentationContext) {}

// EnterSpanLineStartNoStar is called when production spanLineStartNoStar is entered.
func (s *BaseReStructuredTextListener) EnterSpanLineStartNoStar(ctx *SpanLineStartNoStarContext) {}

// ExitSpanLineStartNoStar is called when production spanLineStartNoStar is exited.
func (s *BaseReStructuredTextListener) ExitSpanLineStartNoStar(ctx *SpanLineStartNoStarContext) {}

// EnterTextLineStart is called when production textLineStart is entered.
func (s *BaseReStructuredTextListener) EnterTextLineStart(ctx *TextLineStartContext) {}

// ExitTextLineStart is called when production textLineStart is exited.
func (s *BaseReStructuredTextListener) ExitTextLineStart(ctx *TextLineStartContext) {}

// EnterLineStart_fragment is called when production lineStart_fragment is entered.
func (s *BaseReStructuredTextListener) EnterLineStart_fragment(ctx *LineStart_fragmentContext) {}

// ExitLineStart_fragment is called when production lineStart_fragment is exited.
func (s *BaseReStructuredTextListener) ExitLineStart_fragment(ctx *LineStart_fragmentContext) {}

// EnterText is called when production text is entered.
func (s *BaseReStructuredTextListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseReStructuredTextListener) ExitText(ctx *TextContext) {}

// EnterTextStart is called when production textStart is entered.
func (s *BaseReStructuredTextListener) EnterTextStart(ctx *TextStartContext) {}

// ExitTextStart is called when production textStart is exited.
func (s *BaseReStructuredTextListener) ExitTextStart(ctx *TextStartContext) {}

// EnterForcedText is called when production forcedText is entered.
func (s *BaseReStructuredTextListener) EnterForcedText(ctx *ForcedTextContext) {}

// ExitForcedText is called when production forcedText is exited.
func (s *BaseReStructuredTextListener) ExitForcedText(ctx *ForcedTextContext) {}

// EnterSpanNoStar is called when production spanNoStar is entered.
func (s *BaseReStructuredTextListener) EnterSpanNoStar(ctx *SpanNoStarContext) {}

// ExitSpanNoStar is called when production spanNoStar is exited.
func (s *BaseReStructuredTextListener) ExitSpanNoStar(ctx *SpanNoStarContext) {}

// EnterSpan is called when production span is entered.
func (s *BaseReStructuredTextListener) EnterSpan(ctx *SpanContext) {}

// ExitSpan is called when production span is exited.
func (s *BaseReStructuredTextListener) ExitSpan(ctx *SpanContext) {}

// EnterQuotedLiteral is called when production quotedLiteral is entered.
func (s *BaseReStructuredTextListener) EnterQuotedLiteral(ctx *QuotedLiteralContext) {}

// ExitQuotedLiteral is called when production quotedLiteral is exited.
func (s *BaseReStructuredTextListener) ExitQuotedLiteral(ctx *QuotedLiteralContext) {}

// EnterText_fragment_start is called when production text_fragment_start is entered.
func (s *BaseReStructuredTextListener) EnterText_fragment_start(ctx *Text_fragment_startContext) {}

// ExitText_fragment_start is called when production text_fragment_start is exited.
func (s *BaseReStructuredTextListener) ExitText_fragment_start(ctx *Text_fragment_startContext) {}

// EnterText_fragment is called when production text_fragment is entered.
func (s *BaseReStructuredTextListener) EnterText_fragment(ctx *Text_fragmentContext) {}

// ExitText_fragment is called when production text_fragment is exited.
func (s *BaseReStructuredTextListener) ExitText_fragment(ctx *Text_fragmentContext) {}

// EnterStarText is called when production starText is entered.
func (s *BaseReStructuredTextListener) EnterStarText(ctx *StarTextContext) {}

// ExitStarText is called when production starText is exited.
func (s *BaseReStructuredTextListener) ExitStarText(ctx *StarTextContext) {}

// EnterStarAtoms is called when production starAtoms is entered.
func (s *BaseReStructuredTextListener) EnterStarAtoms(ctx *StarAtomsContext) {}

// ExitStarAtoms is called when production starAtoms is exited.
func (s *BaseReStructuredTextListener) ExitStarAtoms(ctx *StarAtomsContext) {}

// EnterStarNoSpace is called when production starNoSpace is entered.
func (s *BaseReStructuredTextListener) EnterStarNoSpace(ctx *StarNoSpaceContext) {}

// ExitStarNoSpace is called when production starNoSpace is exited.
func (s *BaseReStructuredTextListener) ExitStarNoSpace(ctx *StarNoSpaceContext) {}

// EnterStarAtom is called when production starAtom is entered.
func (s *BaseReStructuredTextListener) EnterStarAtom(ctx *StarAtomContext) {}

// ExitStarAtom is called when production starAtom is exited.
func (s *BaseReStructuredTextListener) ExitStarAtom(ctx *StarAtomContext) {}

// EnterBackTickText is called when production backTickText is entered.
func (s *BaseReStructuredTextListener) EnterBackTickText(ctx *BackTickTextContext) {}

// ExitBackTickText is called when production backTickText is exited.
func (s *BaseReStructuredTextListener) ExitBackTickText(ctx *BackTickTextContext) {}

// EnterBody is called when production body is entered.
func (s *BaseReStructuredTextListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseReStructuredTextListener) ExitBody(ctx *BodyContext) {}

// EnterBackTickAtoms is called when production backTickAtoms is entered.
func (s *BaseReStructuredTextListener) EnterBackTickAtoms(ctx *BackTickAtomsContext) {}

// ExitBackTickAtoms is called when production backTickAtoms is exited.
func (s *BaseReStructuredTextListener) ExitBackTickAtoms(ctx *BackTickAtomsContext) {}

// EnterBackTickNoSpace is called when production backTickNoSpace is entered.
func (s *BaseReStructuredTextListener) EnterBackTickNoSpace(ctx *BackTickNoSpaceContext) {}

// ExitBackTickNoSpace is called when production backTickNoSpace is exited.
func (s *BaseReStructuredTextListener) ExitBackTickNoSpace(ctx *BackTickNoSpaceContext) {}

// EnterBackTickAtom is called when production backTickAtom is entered.
func (s *BaseReStructuredTextListener) EnterBackTickAtom(ctx *BackTickAtomContext) {}

// ExitBackTickAtom is called when production backTickAtom is exited.
func (s *BaseReStructuredTextListener) ExitBackTickAtom(ctx *BackTickAtomContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseReStructuredTextListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseReStructuredTextListener) ExitReference(ctx *ReferenceContext) {}

// EnterReferenceIn is called when production referenceIn is entered.
func (s *BaseReStructuredTextListener) EnterReferenceIn(ctx *ReferenceInContext) {}

// ExitReferenceIn is called when production referenceIn is exited.
func (s *BaseReStructuredTextListener) ExitReferenceIn(ctx *ReferenceInContext) {}

// EnterHyperlinkTarget is called when production hyperlinkTarget is entered.
func (s *BaseReStructuredTextListener) EnterHyperlinkTarget(ctx *HyperlinkTargetContext) {}

// ExitHyperlinkTarget is called when production hyperlinkTarget is exited.
func (s *BaseReStructuredTextListener) ExitHyperlinkTarget(ctx *HyperlinkTargetContext) {}

// EnterHyperlink is called when production hyperlink is entered.
func (s *BaseReStructuredTextListener) EnterHyperlink(ctx *HyperlinkContext) {}

// ExitHyperlink is called when production hyperlink is exited.
func (s *BaseReStructuredTextListener) ExitHyperlink(ctx *HyperlinkContext) {}

// EnterHyperlinkDoc is called when production hyperlinkDoc is entered.
func (s *BaseReStructuredTextListener) EnterHyperlinkDoc(ctx *HyperlinkDocContext) {}

// ExitHyperlinkDoc is called when production hyperlinkDoc is exited.
func (s *BaseReStructuredTextListener) ExitHyperlinkDoc(ctx *HyperlinkDocContext) {}

// EnterUrl is called when production url is entered.
func (s *BaseReStructuredTextListener) EnterUrl(ctx *UrlContext) {}

// ExitUrl is called when production url is exited.
func (s *BaseReStructuredTextListener) ExitUrl(ctx *UrlContext) {}

// EnterUrlAtom is called when production urlAtom is entered.
func (s *BaseReStructuredTextListener) EnterUrlAtom(ctx *UrlAtomContext) {}

// ExitUrlAtom is called when production urlAtom is exited.
func (s *BaseReStructuredTextListener) ExitUrlAtom(ctx *UrlAtomContext) {}

// EnterHyperlinkAtom is called when production hyperlinkAtom is entered.
func (s *BaseReStructuredTextListener) EnterHyperlinkAtom(ctx *HyperlinkAtomContext) {}

// ExitHyperlinkAtom is called when production hyperlinkAtom is exited.
func (s *BaseReStructuredTextListener) ExitHyperlinkAtom(ctx *HyperlinkAtomContext) {}

// EnterSeparator is called when production separator is entered.
func (s *BaseReStructuredTextListener) EnterSeparator(ctx *SeparatorContext) {}

// ExitSeparator is called when production separator is exited.
func (s *BaseReStructuredTextListener) ExitSeparator(ctx *SeparatorContext) {}
