// Code generated from css3.g4 by ANTLR 4.9.3. DO NOT EDIT.

package css3 // css3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Basecss3Listener is a complete listener for a parse tree produced by css3Parser.
type Basecss3Listener struct{}

var _ css3Listener = &Basecss3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basecss3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basecss3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basecss3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basecss3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStylesheet is called when production stylesheet is entered.
func (s *Basecss3Listener) EnterStylesheet(ctx *StylesheetContext) {}

// ExitStylesheet is called when production stylesheet is exited.
func (s *Basecss3Listener) ExitStylesheet(ctx *StylesheetContext) {}

// EnterGoodCharset is called when production goodCharset is entered.
func (s *Basecss3Listener) EnterGoodCharset(ctx *GoodCharsetContext) {}

// ExitGoodCharset is called when production goodCharset is exited.
func (s *Basecss3Listener) ExitGoodCharset(ctx *GoodCharsetContext) {}

// EnterBadCharset is called when production badCharset is entered.
func (s *Basecss3Listener) EnterBadCharset(ctx *BadCharsetContext) {}

// ExitBadCharset is called when production badCharset is exited.
func (s *Basecss3Listener) ExitBadCharset(ctx *BadCharsetContext) {}

// EnterGoodImport is called when production goodImport is entered.
func (s *Basecss3Listener) EnterGoodImport(ctx *GoodImportContext) {}

// ExitGoodImport is called when production goodImport is exited.
func (s *Basecss3Listener) ExitGoodImport(ctx *GoodImportContext) {}

// EnterBadImport is called when production badImport is entered.
func (s *Basecss3Listener) EnterBadImport(ctx *BadImportContext) {}

// ExitBadImport is called when production badImport is exited.
func (s *Basecss3Listener) ExitBadImport(ctx *BadImportContext) {}

// EnterGoodNamespace is called when production goodNamespace is entered.
func (s *Basecss3Listener) EnterGoodNamespace(ctx *GoodNamespaceContext) {}

// ExitGoodNamespace is called when production goodNamespace is exited.
func (s *Basecss3Listener) ExitGoodNamespace(ctx *GoodNamespaceContext) {}

// EnterBadNamespace is called when production badNamespace is entered.
func (s *Basecss3Listener) EnterBadNamespace(ctx *BadNamespaceContext) {}

// ExitBadNamespace is called when production badNamespace is exited.
func (s *Basecss3Listener) ExitBadNamespace(ctx *BadNamespaceContext) {}

// EnterNamespacePrefix is called when production namespacePrefix is entered.
func (s *Basecss3Listener) EnterNamespacePrefix(ctx *NamespacePrefixContext) {}

// ExitNamespacePrefix is called when production namespacePrefix is exited.
func (s *Basecss3Listener) ExitNamespacePrefix(ctx *NamespacePrefixContext) {}

// EnterMedia is called when production media is entered.
func (s *Basecss3Listener) EnterMedia(ctx *MediaContext) {}

// ExitMedia is called when production media is exited.
func (s *Basecss3Listener) ExitMedia(ctx *MediaContext) {}

// EnterMediaQueryList is called when production mediaQueryList is entered.
func (s *Basecss3Listener) EnterMediaQueryList(ctx *MediaQueryListContext) {}

// ExitMediaQueryList is called when production mediaQueryList is exited.
func (s *Basecss3Listener) ExitMediaQueryList(ctx *MediaQueryListContext) {}

// EnterMediaQuery is called when production mediaQuery is entered.
func (s *Basecss3Listener) EnterMediaQuery(ctx *MediaQueryContext) {}

// ExitMediaQuery is called when production mediaQuery is exited.
func (s *Basecss3Listener) ExitMediaQuery(ctx *MediaQueryContext) {}

// EnterMediaType is called when production mediaType is entered.
func (s *Basecss3Listener) EnterMediaType(ctx *MediaTypeContext) {}

// ExitMediaType is called when production mediaType is exited.
func (s *Basecss3Listener) ExitMediaType(ctx *MediaTypeContext) {}

// EnterMediaExpression is called when production mediaExpression is entered.
func (s *Basecss3Listener) EnterMediaExpression(ctx *MediaExpressionContext) {}

// ExitMediaExpression is called when production mediaExpression is exited.
func (s *Basecss3Listener) ExitMediaExpression(ctx *MediaExpressionContext) {}

// EnterMediaFeature is called when production mediaFeature is entered.
func (s *Basecss3Listener) EnterMediaFeature(ctx *MediaFeatureContext) {}

// ExitMediaFeature is called when production mediaFeature is exited.
func (s *Basecss3Listener) ExitMediaFeature(ctx *MediaFeatureContext) {}

// EnterPage is called when production page is entered.
func (s *Basecss3Listener) EnterPage(ctx *PageContext) {}

// ExitPage is called when production page is exited.
func (s *Basecss3Listener) ExitPage(ctx *PageContext) {}

// EnterPseudoPage is called when production pseudoPage is entered.
func (s *Basecss3Listener) EnterPseudoPage(ctx *PseudoPageContext) {}

// ExitPseudoPage is called when production pseudoPage is exited.
func (s *Basecss3Listener) ExitPseudoPage(ctx *PseudoPageContext) {}

// EnterSelectorGroup is called when production selectorGroup is entered.
func (s *Basecss3Listener) EnterSelectorGroup(ctx *SelectorGroupContext) {}

// ExitSelectorGroup is called when production selectorGroup is exited.
func (s *Basecss3Listener) ExitSelectorGroup(ctx *SelectorGroupContext) {}

// EnterSelector is called when production selector is entered.
func (s *Basecss3Listener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *Basecss3Listener) ExitSelector(ctx *SelectorContext) {}

// EnterCombinator is called when production combinator is entered.
func (s *Basecss3Listener) EnterCombinator(ctx *CombinatorContext) {}

// ExitCombinator is called when production combinator is exited.
func (s *Basecss3Listener) ExitCombinator(ctx *CombinatorContext) {}

// EnterSimpleSelectorSequence is called when production simpleSelectorSequence is entered.
func (s *Basecss3Listener) EnterSimpleSelectorSequence(ctx *SimpleSelectorSequenceContext) {}

// ExitSimpleSelectorSequence is called when production simpleSelectorSequence is exited.
func (s *Basecss3Listener) ExitSimpleSelectorSequence(ctx *SimpleSelectorSequenceContext) {}

// EnterTypeSelector is called when production typeSelector is entered.
func (s *Basecss3Listener) EnterTypeSelector(ctx *TypeSelectorContext) {}

// ExitTypeSelector is called when production typeSelector is exited.
func (s *Basecss3Listener) ExitTypeSelector(ctx *TypeSelectorContext) {}

// EnterTypeNamespacePrefix is called when production typeNamespacePrefix is entered.
func (s *Basecss3Listener) EnterTypeNamespacePrefix(ctx *TypeNamespacePrefixContext) {}

// ExitTypeNamespacePrefix is called when production typeNamespacePrefix is exited.
func (s *Basecss3Listener) ExitTypeNamespacePrefix(ctx *TypeNamespacePrefixContext) {}

// EnterElementName is called when production elementName is entered.
func (s *Basecss3Listener) EnterElementName(ctx *ElementNameContext) {}

// ExitElementName is called when production elementName is exited.
func (s *Basecss3Listener) ExitElementName(ctx *ElementNameContext) {}

// EnterUniversal is called when production universal is entered.
func (s *Basecss3Listener) EnterUniversal(ctx *UniversalContext) {}

// ExitUniversal is called when production universal is exited.
func (s *Basecss3Listener) ExitUniversal(ctx *UniversalContext) {}

// EnterClassName is called when production className is entered.
func (s *Basecss3Listener) EnterClassName(ctx *ClassNameContext) {}

// ExitClassName is called when production className is exited.
func (s *Basecss3Listener) ExitClassName(ctx *ClassNameContext) {}

// EnterAttrib is called when production attrib is entered.
func (s *Basecss3Listener) EnterAttrib(ctx *AttribContext) {}

// ExitAttrib is called when production attrib is exited.
func (s *Basecss3Listener) ExitAttrib(ctx *AttribContext) {}

// EnterPseudo is called when production pseudo is entered.
func (s *Basecss3Listener) EnterPseudo(ctx *PseudoContext) {}

// ExitPseudo is called when production pseudo is exited.
func (s *Basecss3Listener) ExitPseudo(ctx *PseudoContext) {}

// EnterFunctionalPseudo is called when production functionalPseudo is entered.
func (s *Basecss3Listener) EnterFunctionalPseudo(ctx *FunctionalPseudoContext) {}

// ExitFunctionalPseudo is called when production functionalPseudo is exited.
func (s *Basecss3Listener) ExitFunctionalPseudo(ctx *FunctionalPseudoContext) {}

// EnterExpression is called when production expression is entered.
func (s *Basecss3Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Basecss3Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterNegation is called when production negation is entered.
func (s *Basecss3Listener) EnterNegation(ctx *NegationContext) {}

// ExitNegation is called when production negation is exited.
func (s *Basecss3Listener) ExitNegation(ctx *NegationContext) {}

// EnterNegationArg is called when production negationArg is entered.
func (s *Basecss3Listener) EnterNegationArg(ctx *NegationArgContext) {}

// ExitNegationArg is called when production negationArg is exited.
func (s *Basecss3Listener) ExitNegationArg(ctx *NegationArgContext) {}

// EnterGoodOperator is called when production goodOperator is entered.
func (s *Basecss3Listener) EnterGoodOperator(ctx *GoodOperatorContext) {}

// ExitGoodOperator is called when production goodOperator is exited.
func (s *Basecss3Listener) ExitGoodOperator(ctx *GoodOperatorContext) {}

// EnterBadOperator is called when production badOperator is entered.
func (s *Basecss3Listener) EnterBadOperator(ctx *BadOperatorContext) {}

// ExitBadOperator is called when production badOperator is exited.
func (s *Basecss3Listener) ExitBadOperator(ctx *BadOperatorContext) {}

// EnterGoodProperty is called when production goodProperty is entered.
func (s *Basecss3Listener) EnterGoodProperty(ctx *GoodPropertyContext) {}

// ExitGoodProperty is called when production goodProperty is exited.
func (s *Basecss3Listener) ExitGoodProperty(ctx *GoodPropertyContext) {}

// EnterBadProperty is called when production badProperty is entered.
func (s *Basecss3Listener) EnterBadProperty(ctx *BadPropertyContext) {}

// ExitBadProperty is called when production badProperty is exited.
func (s *Basecss3Listener) ExitBadProperty(ctx *BadPropertyContext) {}

// EnterKnownRuleset is called when production knownRuleset is entered.
func (s *Basecss3Listener) EnterKnownRuleset(ctx *KnownRulesetContext) {}

// ExitKnownRuleset is called when production knownRuleset is exited.
func (s *Basecss3Listener) ExitKnownRuleset(ctx *KnownRulesetContext) {}

// EnterUnknownRuleset is called when production unknownRuleset is entered.
func (s *Basecss3Listener) EnterUnknownRuleset(ctx *UnknownRulesetContext) {}

// ExitUnknownRuleset is called when production unknownRuleset is exited.
func (s *Basecss3Listener) ExitUnknownRuleset(ctx *UnknownRulesetContext) {}

// EnterDeclarationList is called when production declarationList is entered.
func (s *Basecss3Listener) EnterDeclarationList(ctx *DeclarationListContext) {}

// ExitDeclarationList is called when production declarationList is exited.
func (s *Basecss3Listener) ExitDeclarationList(ctx *DeclarationListContext) {}

// EnterKnownDeclaration is called when production knownDeclaration is entered.
func (s *Basecss3Listener) EnterKnownDeclaration(ctx *KnownDeclarationContext) {}

// ExitKnownDeclaration is called when production knownDeclaration is exited.
func (s *Basecss3Listener) ExitKnownDeclaration(ctx *KnownDeclarationContext) {}

// EnterUnknownDeclaration is called when production unknownDeclaration is entered.
func (s *Basecss3Listener) EnterUnknownDeclaration(ctx *UnknownDeclarationContext) {}

// ExitUnknownDeclaration is called when production unknownDeclaration is exited.
func (s *Basecss3Listener) ExitUnknownDeclaration(ctx *UnknownDeclarationContext) {}

// EnterPrio is called when production prio is entered.
func (s *Basecss3Listener) EnterPrio(ctx *PrioContext) {}

// ExitPrio is called when production prio is exited.
func (s *Basecss3Listener) ExitPrio(ctx *PrioContext) {}

// EnterValue is called when production value is entered.
func (s *Basecss3Listener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *Basecss3Listener) ExitValue(ctx *ValueContext) {}

// EnterExpr is called when production expr is entered.
func (s *Basecss3Listener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *Basecss3Listener) ExitExpr(ctx *ExprContext) {}

// EnterKnownTerm is called when production knownTerm is entered.
func (s *Basecss3Listener) EnterKnownTerm(ctx *KnownTermContext) {}

// ExitKnownTerm is called when production knownTerm is exited.
func (s *Basecss3Listener) ExitKnownTerm(ctx *KnownTermContext) {}

// EnterUnknownTerm is called when production unknownTerm is entered.
func (s *Basecss3Listener) EnterUnknownTerm(ctx *UnknownTermContext) {}

// ExitUnknownTerm is called when production unknownTerm is exited.
func (s *Basecss3Listener) ExitUnknownTerm(ctx *UnknownTermContext) {}

// EnterBadTerm is called when production badTerm is entered.
func (s *Basecss3Listener) EnterBadTerm(ctx *BadTermContext) {}

// ExitBadTerm is called when production badTerm is exited.
func (s *Basecss3Listener) ExitBadTerm(ctx *BadTermContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *Basecss3Listener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *Basecss3Listener) ExitFunction_(ctx *Function_Context) {}

// EnterDxImageTransform is called when production dxImageTransform is entered.
func (s *Basecss3Listener) EnterDxImageTransform(ctx *DxImageTransformContext) {}

// ExitDxImageTransform is called when production dxImageTransform is exited.
func (s *Basecss3Listener) ExitDxImageTransform(ctx *DxImageTransformContext) {}

// EnterHexcolor is called when production hexcolor is entered.
func (s *Basecss3Listener) EnterHexcolor(ctx *HexcolorContext) {}

// ExitHexcolor is called when production hexcolor is exited.
func (s *Basecss3Listener) ExitHexcolor(ctx *HexcolorContext) {}

// EnterNumber is called when production number is entered.
func (s *Basecss3Listener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *Basecss3Listener) ExitNumber(ctx *NumberContext) {}

// EnterPercentage is called when production percentage is entered.
func (s *Basecss3Listener) EnterPercentage(ctx *PercentageContext) {}

// ExitPercentage is called when production percentage is exited.
func (s *Basecss3Listener) ExitPercentage(ctx *PercentageContext) {}

// EnterDimension is called when production dimension is entered.
func (s *Basecss3Listener) EnterDimension(ctx *DimensionContext) {}

// ExitDimension is called when production dimension is exited.
func (s *Basecss3Listener) ExitDimension(ctx *DimensionContext) {}

// EnterUnknownDimension is called when production unknownDimension is entered.
func (s *Basecss3Listener) EnterUnknownDimension(ctx *UnknownDimensionContext) {}

// ExitUnknownDimension is called when production unknownDimension is exited.
func (s *Basecss3Listener) ExitUnknownDimension(ctx *UnknownDimensionContext) {}

// EnterAny_ is called when production any_ is entered.
func (s *Basecss3Listener) EnterAny_(ctx *Any_Context) {}

// ExitAny_ is called when production any_ is exited.
func (s *Basecss3Listener) ExitAny_(ctx *Any_Context) {}

// EnterUnknownAtRule is called when production unknownAtRule is entered.
func (s *Basecss3Listener) EnterUnknownAtRule(ctx *UnknownAtRuleContext) {}

// ExitUnknownAtRule is called when production unknownAtRule is exited.
func (s *Basecss3Listener) ExitUnknownAtRule(ctx *UnknownAtRuleContext) {}

// EnterAtKeyword is called when production atKeyword is entered.
func (s *Basecss3Listener) EnterAtKeyword(ctx *AtKeywordContext) {}

// ExitAtKeyword is called when production atKeyword is exited.
func (s *Basecss3Listener) ExitAtKeyword(ctx *AtKeywordContext) {}

// EnterUnused is called when production unused is entered.
func (s *Basecss3Listener) EnterUnused(ctx *UnusedContext) {}

// ExitUnused is called when production unused is exited.
func (s *Basecss3Listener) ExitUnused(ctx *UnusedContext) {}

// EnterBlock is called when production block is entered.
func (s *Basecss3Listener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *Basecss3Listener) ExitBlock(ctx *BlockContext) {}

// EnterNestedStatement is called when production nestedStatement is entered.
func (s *Basecss3Listener) EnterNestedStatement(ctx *NestedStatementContext) {}

// ExitNestedStatement is called when production nestedStatement is exited.
func (s *Basecss3Listener) ExitNestedStatement(ctx *NestedStatementContext) {}

// EnterGroupRuleBody is called when production groupRuleBody is entered.
func (s *Basecss3Listener) EnterGroupRuleBody(ctx *GroupRuleBodyContext) {}

// ExitGroupRuleBody is called when production groupRuleBody is exited.
func (s *Basecss3Listener) ExitGroupRuleBody(ctx *GroupRuleBodyContext) {}

// EnterSupportsRule is called when production supportsRule is entered.
func (s *Basecss3Listener) EnterSupportsRule(ctx *SupportsRuleContext) {}

// ExitSupportsRule is called when production supportsRule is exited.
func (s *Basecss3Listener) ExitSupportsRule(ctx *SupportsRuleContext) {}

// EnterSupportsCondition is called when production supportsCondition is entered.
func (s *Basecss3Listener) EnterSupportsCondition(ctx *SupportsConditionContext) {}

// ExitSupportsCondition is called when production supportsCondition is exited.
func (s *Basecss3Listener) ExitSupportsCondition(ctx *SupportsConditionContext) {}

// EnterSupportsConditionInParens is called when production supportsConditionInParens is entered.
func (s *Basecss3Listener) EnterSupportsConditionInParens(ctx *SupportsConditionInParensContext) {}

// ExitSupportsConditionInParens is called when production supportsConditionInParens is exited.
func (s *Basecss3Listener) ExitSupportsConditionInParens(ctx *SupportsConditionInParensContext) {}

// EnterSupportsNegation is called when production supportsNegation is entered.
func (s *Basecss3Listener) EnterSupportsNegation(ctx *SupportsNegationContext) {}

// ExitSupportsNegation is called when production supportsNegation is exited.
func (s *Basecss3Listener) ExitSupportsNegation(ctx *SupportsNegationContext) {}

// EnterSupportsConjunction is called when production supportsConjunction is entered.
func (s *Basecss3Listener) EnterSupportsConjunction(ctx *SupportsConjunctionContext) {}

// ExitSupportsConjunction is called when production supportsConjunction is exited.
func (s *Basecss3Listener) ExitSupportsConjunction(ctx *SupportsConjunctionContext) {}

// EnterSupportsDisjunction is called when production supportsDisjunction is entered.
func (s *Basecss3Listener) EnterSupportsDisjunction(ctx *SupportsDisjunctionContext) {}

// ExitSupportsDisjunction is called when production supportsDisjunction is exited.
func (s *Basecss3Listener) ExitSupportsDisjunction(ctx *SupportsDisjunctionContext) {}

// EnterSupportsDeclarationCondition is called when production supportsDeclarationCondition is entered.
func (s *Basecss3Listener) EnterSupportsDeclarationCondition(ctx *SupportsDeclarationConditionContext) {
}

// ExitSupportsDeclarationCondition is called when production supportsDeclarationCondition is exited.
func (s *Basecss3Listener) ExitSupportsDeclarationCondition(ctx *SupportsDeclarationConditionContext) {
}

// EnterGeneralEnclosed is called when production generalEnclosed is entered.
func (s *Basecss3Listener) EnterGeneralEnclosed(ctx *GeneralEnclosedContext) {}

// ExitGeneralEnclosed is called when production generalEnclosed is exited.
func (s *Basecss3Listener) ExitGeneralEnclosed(ctx *GeneralEnclosedContext) {}

// EnterVar_ is called when production var_ is entered.
func (s *Basecss3Listener) EnterVar_(ctx *Var_Context) {}

// ExitVar_ is called when production var_ is exited.
func (s *Basecss3Listener) ExitVar_(ctx *Var_Context) {}

// EnterCalc is called when production calc is entered.
func (s *Basecss3Listener) EnterCalc(ctx *CalcContext) {}

// ExitCalc is called when production calc is exited.
func (s *Basecss3Listener) ExitCalc(ctx *CalcContext) {}

// EnterCalcSum is called when production calcSum is entered.
func (s *Basecss3Listener) EnterCalcSum(ctx *CalcSumContext) {}

// ExitCalcSum is called when production calcSum is exited.
func (s *Basecss3Listener) ExitCalcSum(ctx *CalcSumContext) {}

// EnterCalcProduct is called when production calcProduct is entered.
func (s *Basecss3Listener) EnterCalcProduct(ctx *CalcProductContext) {}

// ExitCalcProduct is called when production calcProduct is exited.
func (s *Basecss3Listener) ExitCalcProduct(ctx *CalcProductContext) {}

// EnterCalcValue is called when production calcValue is entered.
func (s *Basecss3Listener) EnterCalcValue(ctx *CalcValueContext) {}

// ExitCalcValue is called when production calcValue is exited.
func (s *Basecss3Listener) ExitCalcValue(ctx *CalcValueContext) {}

// EnterFontFaceRule is called when production fontFaceRule is entered.
func (s *Basecss3Listener) EnterFontFaceRule(ctx *FontFaceRuleContext) {}

// ExitFontFaceRule is called when production fontFaceRule is exited.
func (s *Basecss3Listener) ExitFontFaceRule(ctx *FontFaceRuleContext) {}

// EnterKnownFontFaceDeclaration is called when production knownFontFaceDeclaration is entered.
func (s *Basecss3Listener) EnterKnownFontFaceDeclaration(ctx *KnownFontFaceDeclarationContext) {}

// ExitKnownFontFaceDeclaration is called when production knownFontFaceDeclaration is exited.
func (s *Basecss3Listener) ExitKnownFontFaceDeclaration(ctx *KnownFontFaceDeclarationContext) {}

// EnterUnknownFontFaceDeclaration is called when production unknownFontFaceDeclaration is entered.
func (s *Basecss3Listener) EnterUnknownFontFaceDeclaration(ctx *UnknownFontFaceDeclarationContext) {}

// ExitUnknownFontFaceDeclaration is called when production unknownFontFaceDeclaration is exited.
func (s *Basecss3Listener) ExitUnknownFontFaceDeclaration(ctx *UnknownFontFaceDeclarationContext) {}

// EnterKeyframesRule is called when production keyframesRule is entered.
func (s *Basecss3Listener) EnterKeyframesRule(ctx *KeyframesRuleContext) {}

// ExitKeyframesRule is called when production keyframesRule is exited.
func (s *Basecss3Listener) ExitKeyframesRule(ctx *KeyframesRuleContext) {}

// EnterKeyframesBlocks is called when production keyframesBlocks is entered.
func (s *Basecss3Listener) EnterKeyframesBlocks(ctx *KeyframesBlocksContext) {}

// ExitKeyframesBlocks is called when production keyframesBlocks is exited.
func (s *Basecss3Listener) ExitKeyframesBlocks(ctx *KeyframesBlocksContext) {}

// EnterKeyframeSelector is called when production keyframeSelector is entered.
func (s *Basecss3Listener) EnterKeyframeSelector(ctx *KeyframeSelectorContext) {}

// ExitKeyframeSelector is called when production keyframeSelector is exited.
func (s *Basecss3Listener) ExitKeyframeSelector(ctx *KeyframeSelectorContext) {}

// EnterViewport is called when production viewport is entered.
func (s *Basecss3Listener) EnterViewport(ctx *ViewportContext) {}

// ExitViewport is called when production viewport is exited.
func (s *Basecss3Listener) ExitViewport(ctx *ViewportContext) {}

// EnterCounterStyle is called when production counterStyle is entered.
func (s *Basecss3Listener) EnterCounterStyle(ctx *CounterStyleContext) {}

// ExitCounterStyle is called when production counterStyle is exited.
func (s *Basecss3Listener) ExitCounterStyle(ctx *CounterStyleContext) {}

// EnterFontFeatureValuesRule is called when production fontFeatureValuesRule is entered.
func (s *Basecss3Listener) EnterFontFeatureValuesRule(ctx *FontFeatureValuesRuleContext) {}

// ExitFontFeatureValuesRule is called when production fontFeatureValuesRule is exited.
func (s *Basecss3Listener) ExitFontFeatureValuesRule(ctx *FontFeatureValuesRuleContext) {}

// EnterFontFamilyNameList is called when production fontFamilyNameList is entered.
func (s *Basecss3Listener) EnterFontFamilyNameList(ctx *FontFamilyNameListContext) {}

// ExitFontFamilyNameList is called when production fontFamilyNameList is exited.
func (s *Basecss3Listener) ExitFontFamilyNameList(ctx *FontFamilyNameListContext) {}

// EnterFontFamilyName is called when production fontFamilyName is entered.
func (s *Basecss3Listener) EnterFontFamilyName(ctx *FontFamilyNameContext) {}

// ExitFontFamilyName is called when production fontFamilyName is exited.
func (s *Basecss3Listener) ExitFontFamilyName(ctx *FontFamilyNameContext) {}

// EnterFeatureValueBlock is called when production featureValueBlock is entered.
func (s *Basecss3Listener) EnterFeatureValueBlock(ctx *FeatureValueBlockContext) {}

// ExitFeatureValueBlock is called when production featureValueBlock is exited.
func (s *Basecss3Listener) ExitFeatureValueBlock(ctx *FeatureValueBlockContext) {}

// EnterFeatureType is called when production featureType is entered.
func (s *Basecss3Listener) EnterFeatureType(ctx *FeatureTypeContext) {}

// ExitFeatureType is called when production featureType is exited.
func (s *Basecss3Listener) ExitFeatureType(ctx *FeatureTypeContext) {}

// EnterFeatureValueDefinition is called when production featureValueDefinition is entered.
func (s *Basecss3Listener) EnterFeatureValueDefinition(ctx *FeatureValueDefinitionContext) {}

// ExitFeatureValueDefinition is called when production featureValueDefinition is exited.
func (s *Basecss3Listener) ExitFeatureValueDefinition(ctx *FeatureValueDefinitionContext) {}

// EnterIdent is called when production ident is entered.
func (s *Basecss3Listener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *Basecss3Listener) ExitIdent(ctx *IdentContext) {}

// EnterWs is called when production ws is entered.
func (s *Basecss3Listener) EnterWs(ctx *WsContext) {}

// ExitWs is called when production ws is exited.
func (s *Basecss3Listener) ExitWs(ctx *WsContext) {}
