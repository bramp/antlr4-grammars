// Code generated from css3.g4 by ANTLR 4.9.3. DO NOT EDIT.

package css3 // css3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// css3Listener is a complete listener for a parse tree produced by css3Parser.
type css3Listener interface {
	antlr.ParseTreeListener

	// EnterStylesheet is called when entering the stylesheet production.
	EnterStylesheet(c *StylesheetContext)

	// EnterGoodCharset is called when entering the goodCharset production.
	EnterGoodCharset(c *GoodCharsetContext)

	// EnterBadCharset is called when entering the badCharset production.
	EnterBadCharset(c *BadCharsetContext)

	// EnterGoodImport is called when entering the goodImport production.
	EnterGoodImport(c *GoodImportContext)

	// EnterBadImport is called when entering the badImport production.
	EnterBadImport(c *BadImportContext)

	// EnterGoodNamespace is called when entering the goodNamespace production.
	EnterGoodNamespace(c *GoodNamespaceContext)

	// EnterBadNamespace is called when entering the badNamespace production.
	EnterBadNamespace(c *BadNamespaceContext)

	// EnterNamespacePrefix is called when entering the namespacePrefix production.
	EnterNamespacePrefix(c *NamespacePrefixContext)

	// EnterMedia is called when entering the media production.
	EnterMedia(c *MediaContext)

	// EnterMediaQueryList is called when entering the mediaQueryList production.
	EnterMediaQueryList(c *MediaQueryListContext)

	// EnterMediaQuery is called when entering the mediaQuery production.
	EnterMediaQuery(c *MediaQueryContext)

	// EnterMediaType is called when entering the mediaType production.
	EnterMediaType(c *MediaTypeContext)

	// EnterMediaExpression is called when entering the mediaExpression production.
	EnterMediaExpression(c *MediaExpressionContext)

	// EnterMediaFeature is called when entering the mediaFeature production.
	EnterMediaFeature(c *MediaFeatureContext)

	// EnterPage is called when entering the page production.
	EnterPage(c *PageContext)

	// EnterPseudoPage is called when entering the pseudoPage production.
	EnterPseudoPage(c *PseudoPageContext)

	// EnterSelectorGroup is called when entering the selectorGroup production.
	EnterSelectorGroup(c *SelectorGroupContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterCombinator is called when entering the combinator production.
	EnterCombinator(c *CombinatorContext)

	// EnterSimpleSelectorSequence is called when entering the simpleSelectorSequence production.
	EnterSimpleSelectorSequence(c *SimpleSelectorSequenceContext)

	// EnterTypeSelector is called when entering the typeSelector production.
	EnterTypeSelector(c *TypeSelectorContext)

	// EnterTypeNamespacePrefix is called when entering the typeNamespacePrefix production.
	EnterTypeNamespacePrefix(c *TypeNamespacePrefixContext)

	// EnterElementName is called when entering the elementName production.
	EnterElementName(c *ElementNameContext)

	// EnterUniversal is called when entering the universal production.
	EnterUniversal(c *UniversalContext)

	// EnterClassName is called when entering the className production.
	EnterClassName(c *ClassNameContext)

	// EnterAttrib is called when entering the attrib production.
	EnterAttrib(c *AttribContext)

	// EnterPseudo is called when entering the pseudo production.
	EnterPseudo(c *PseudoContext)

	// EnterFunctionalPseudo is called when entering the functionalPseudo production.
	EnterFunctionalPseudo(c *FunctionalPseudoContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNegation is called when entering the negation production.
	EnterNegation(c *NegationContext)

	// EnterNegationArg is called when entering the negationArg production.
	EnterNegationArg(c *NegationArgContext)

	// EnterGoodOperator is called when entering the goodOperator production.
	EnterGoodOperator(c *GoodOperatorContext)

	// EnterBadOperator is called when entering the badOperator production.
	EnterBadOperator(c *BadOperatorContext)

	// EnterGoodProperty is called when entering the goodProperty production.
	EnterGoodProperty(c *GoodPropertyContext)

	// EnterBadProperty is called when entering the badProperty production.
	EnterBadProperty(c *BadPropertyContext)

	// EnterKnownRuleset is called when entering the knownRuleset production.
	EnterKnownRuleset(c *KnownRulesetContext)

	// EnterUnknownRuleset is called when entering the unknownRuleset production.
	EnterUnknownRuleset(c *UnknownRulesetContext)

	// EnterDeclarationList is called when entering the declarationList production.
	EnterDeclarationList(c *DeclarationListContext)

	// EnterKnownDeclaration is called when entering the knownDeclaration production.
	EnterKnownDeclaration(c *KnownDeclarationContext)

	// EnterUnknownDeclaration is called when entering the unknownDeclaration production.
	EnterUnknownDeclaration(c *UnknownDeclarationContext)

	// EnterPrio is called when entering the prio production.
	EnterPrio(c *PrioContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterKnownTerm is called when entering the knownTerm production.
	EnterKnownTerm(c *KnownTermContext)

	// EnterUnknownTerm is called when entering the unknownTerm production.
	EnterUnknownTerm(c *UnknownTermContext)

	// EnterBadTerm is called when entering the badTerm production.
	EnterBadTerm(c *BadTermContext)

	// EnterFunction_ is called when entering the function_ production.
	EnterFunction_(c *Function_Context)

	// EnterDxImageTransform is called when entering the dxImageTransform production.
	EnterDxImageTransform(c *DxImageTransformContext)

	// EnterHexcolor is called when entering the hexcolor production.
	EnterHexcolor(c *HexcolorContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterPercentage is called when entering the percentage production.
	EnterPercentage(c *PercentageContext)

	// EnterDimension is called when entering the dimension production.
	EnterDimension(c *DimensionContext)

	// EnterUnknownDimension is called when entering the unknownDimension production.
	EnterUnknownDimension(c *UnknownDimensionContext)

	// EnterAny_ is called when entering the any_ production.
	EnterAny_(c *Any_Context)

	// EnterUnknownAtRule is called when entering the unknownAtRule production.
	EnterUnknownAtRule(c *UnknownAtRuleContext)

	// EnterAtKeyword is called when entering the atKeyword production.
	EnterAtKeyword(c *AtKeywordContext)

	// EnterUnused is called when entering the unused production.
	EnterUnused(c *UnusedContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterNestedStatement is called when entering the nestedStatement production.
	EnterNestedStatement(c *NestedStatementContext)

	// EnterGroupRuleBody is called when entering the groupRuleBody production.
	EnterGroupRuleBody(c *GroupRuleBodyContext)

	// EnterSupportsRule is called when entering the supportsRule production.
	EnterSupportsRule(c *SupportsRuleContext)

	// EnterSupportsCondition is called when entering the supportsCondition production.
	EnterSupportsCondition(c *SupportsConditionContext)

	// EnterSupportsConditionInParens is called when entering the supportsConditionInParens production.
	EnterSupportsConditionInParens(c *SupportsConditionInParensContext)

	// EnterSupportsNegation is called when entering the supportsNegation production.
	EnterSupportsNegation(c *SupportsNegationContext)

	// EnterSupportsConjunction is called when entering the supportsConjunction production.
	EnterSupportsConjunction(c *SupportsConjunctionContext)

	// EnterSupportsDisjunction is called when entering the supportsDisjunction production.
	EnterSupportsDisjunction(c *SupportsDisjunctionContext)

	// EnterSupportsDeclarationCondition is called when entering the supportsDeclarationCondition production.
	EnterSupportsDeclarationCondition(c *SupportsDeclarationConditionContext)

	// EnterGeneralEnclosed is called when entering the generalEnclosed production.
	EnterGeneralEnclosed(c *GeneralEnclosedContext)

	// EnterVar_ is called when entering the var_ production.
	EnterVar_(c *Var_Context)

	// EnterCalc is called when entering the calc production.
	EnterCalc(c *CalcContext)

	// EnterCalcSum is called when entering the calcSum production.
	EnterCalcSum(c *CalcSumContext)

	// EnterCalcProduct is called when entering the calcProduct production.
	EnterCalcProduct(c *CalcProductContext)

	// EnterCalcValue is called when entering the calcValue production.
	EnterCalcValue(c *CalcValueContext)

	// EnterFontFaceRule is called when entering the fontFaceRule production.
	EnterFontFaceRule(c *FontFaceRuleContext)

	// EnterKnownFontFaceDeclaration is called when entering the knownFontFaceDeclaration production.
	EnterKnownFontFaceDeclaration(c *KnownFontFaceDeclarationContext)

	// EnterUnknownFontFaceDeclaration is called when entering the unknownFontFaceDeclaration production.
	EnterUnknownFontFaceDeclaration(c *UnknownFontFaceDeclarationContext)

	// EnterKeyframesRule is called when entering the keyframesRule production.
	EnterKeyframesRule(c *KeyframesRuleContext)

	// EnterKeyframesBlocks is called when entering the keyframesBlocks production.
	EnterKeyframesBlocks(c *KeyframesBlocksContext)

	// EnterKeyframeSelector is called when entering the keyframeSelector production.
	EnterKeyframeSelector(c *KeyframeSelectorContext)

	// EnterViewport is called when entering the viewport production.
	EnterViewport(c *ViewportContext)

	// EnterCounterStyle is called when entering the counterStyle production.
	EnterCounterStyle(c *CounterStyleContext)

	// EnterFontFeatureValuesRule is called when entering the fontFeatureValuesRule production.
	EnterFontFeatureValuesRule(c *FontFeatureValuesRuleContext)

	// EnterFontFamilyNameList is called when entering the fontFamilyNameList production.
	EnterFontFamilyNameList(c *FontFamilyNameListContext)

	// EnterFontFamilyName is called when entering the fontFamilyName production.
	EnterFontFamilyName(c *FontFamilyNameContext)

	// EnterFeatureValueBlock is called when entering the featureValueBlock production.
	EnterFeatureValueBlock(c *FeatureValueBlockContext)

	// EnterFeatureType is called when entering the featureType production.
	EnterFeatureType(c *FeatureTypeContext)

	// EnterFeatureValueDefinition is called when entering the featureValueDefinition production.
	EnterFeatureValueDefinition(c *FeatureValueDefinitionContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterWs is called when entering the ws production.
	EnterWs(c *WsContext)

	// ExitStylesheet is called when exiting the stylesheet production.
	ExitStylesheet(c *StylesheetContext)

	// ExitGoodCharset is called when exiting the goodCharset production.
	ExitGoodCharset(c *GoodCharsetContext)

	// ExitBadCharset is called when exiting the badCharset production.
	ExitBadCharset(c *BadCharsetContext)

	// ExitGoodImport is called when exiting the goodImport production.
	ExitGoodImport(c *GoodImportContext)

	// ExitBadImport is called when exiting the badImport production.
	ExitBadImport(c *BadImportContext)

	// ExitGoodNamespace is called when exiting the goodNamespace production.
	ExitGoodNamespace(c *GoodNamespaceContext)

	// ExitBadNamespace is called when exiting the badNamespace production.
	ExitBadNamespace(c *BadNamespaceContext)

	// ExitNamespacePrefix is called when exiting the namespacePrefix production.
	ExitNamespacePrefix(c *NamespacePrefixContext)

	// ExitMedia is called when exiting the media production.
	ExitMedia(c *MediaContext)

	// ExitMediaQueryList is called when exiting the mediaQueryList production.
	ExitMediaQueryList(c *MediaQueryListContext)

	// ExitMediaQuery is called when exiting the mediaQuery production.
	ExitMediaQuery(c *MediaQueryContext)

	// ExitMediaType is called when exiting the mediaType production.
	ExitMediaType(c *MediaTypeContext)

	// ExitMediaExpression is called when exiting the mediaExpression production.
	ExitMediaExpression(c *MediaExpressionContext)

	// ExitMediaFeature is called when exiting the mediaFeature production.
	ExitMediaFeature(c *MediaFeatureContext)

	// ExitPage is called when exiting the page production.
	ExitPage(c *PageContext)

	// ExitPseudoPage is called when exiting the pseudoPage production.
	ExitPseudoPage(c *PseudoPageContext)

	// ExitSelectorGroup is called when exiting the selectorGroup production.
	ExitSelectorGroup(c *SelectorGroupContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitCombinator is called when exiting the combinator production.
	ExitCombinator(c *CombinatorContext)

	// ExitSimpleSelectorSequence is called when exiting the simpleSelectorSequence production.
	ExitSimpleSelectorSequence(c *SimpleSelectorSequenceContext)

	// ExitTypeSelector is called when exiting the typeSelector production.
	ExitTypeSelector(c *TypeSelectorContext)

	// ExitTypeNamespacePrefix is called when exiting the typeNamespacePrefix production.
	ExitTypeNamespacePrefix(c *TypeNamespacePrefixContext)

	// ExitElementName is called when exiting the elementName production.
	ExitElementName(c *ElementNameContext)

	// ExitUniversal is called when exiting the universal production.
	ExitUniversal(c *UniversalContext)

	// ExitClassName is called when exiting the className production.
	ExitClassName(c *ClassNameContext)

	// ExitAttrib is called when exiting the attrib production.
	ExitAttrib(c *AttribContext)

	// ExitPseudo is called when exiting the pseudo production.
	ExitPseudo(c *PseudoContext)

	// ExitFunctionalPseudo is called when exiting the functionalPseudo production.
	ExitFunctionalPseudo(c *FunctionalPseudoContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNegation is called when exiting the negation production.
	ExitNegation(c *NegationContext)

	// ExitNegationArg is called when exiting the negationArg production.
	ExitNegationArg(c *NegationArgContext)

	// ExitGoodOperator is called when exiting the goodOperator production.
	ExitGoodOperator(c *GoodOperatorContext)

	// ExitBadOperator is called when exiting the badOperator production.
	ExitBadOperator(c *BadOperatorContext)

	// ExitGoodProperty is called when exiting the goodProperty production.
	ExitGoodProperty(c *GoodPropertyContext)

	// ExitBadProperty is called when exiting the badProperty production.
	ExitBadProperty(c *BadPropertyContext)

	// ExitKnownRuleset is called when exiting the knownRuleset production.
	ExitKnownRuleset(c *KnownRulesetContext)

	// ExitUnknownRuleset is called when exiting the unknownRuleset production.
	ExitUnknownRuleset(c *UnknownRulesetContext)

	// ExitDeclarationList is called when exiting the declarationList production.
	ExitDeclarationList(c *DeclarationListContext)

	// ExitKnownDeclaration is called when exiting the knownDeclaration production.
	ExitKnownDeclaration(c *KnownDeclarationContext)

	// ExitUnknownDeclaration is called when exiting the unknownDeclaration production.
	ExitUnknownDeclaration(c *UnknownDeclarationContext)

	// ExitPrio is called when exiting the prio production.
	ExitPrio(c *PrioContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitKnownTerm is called when exiting the knownTerm production.
	ExitKnownTerm(c *KnownTermContext)

	// ExitUnknownTerm is called when exiting the unknownTerm production.
	ExitUnknownTerm(c *UnknownTermContext)

	// ExitBadTerm is called when exiting the badTerm production.
	ExitBadTerm(c *BadTermContext)

	// ExitFunction_ is called when exiting the function_ production.
	ExitFunction_(c *Function_Context)

	// ExitDxImageTransform is called when exiting the dxImageTransform production.
	ExitDxImageTransform(c *DxImageTransformContext)

	// ExitHexcolor is called when exiting the hexcolor production.
	ExitHexcolor(c *HexcolorContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitPercentage is called when exiting the percentage production.
	ExitPercentage(c *PercentageContext)

	// ExitDimension is called when exiting the dimension production.
	ExitDimension(c *DimensionContext)

	// ExitUnknownDimension is called when exiting the unknownDimension production.
	ExitUnknownDimension(c *UnknownDimensionContext)

	// ExitAny_ is called when exiting the any_ production.
	ExitAny_(c *Any_Context)

	// ExitUnknownAtRule is called when exiting the unknownAtRule production.
	ExitUnknownAtRule(c *UnknownAtRuleContext)

	// ExitAtKeyword is called when exiting the atKeyword production.
	ExitAtKeyword(c *AtKeywordContext)

	// ExitUnused is called when exiting the unused production.
	ExitUnused(c *UnusedContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitNestedStatement is called when exiting the nestedStatement production.
	ExitNestedStatement(c *NestedStatementContext)

	// ExitGroupRuleBody is called when exiting the groupRuleBody production.
	ExitGroupRuleBody(c *GroupRuleBodyContext)

	// ExitSupportsRule is called when exiting the supportsRule production.
	ExitSupportsRule(c *SupportsRuleContext)

	// ExitSupportsCondition is called when exiting the supportsCondition production.
	ExitSupportsCondition(c *SupportsConditionContext)

	// ExitSupportsConditionInParens is called when exiting the supportsConditionInParens production.
	ExitSupportsConditionInParens(c *SupportsConditionInParensContext)

	// ExitSupportsNegation is called when exiting the supportsNegation production.
	ExitSupportsNegation(c *SupportsNegationContext)

	// ExitSupportsConjunction is called when exiting the supportsConjunction production.
	ExitSupportsConjunction(c *SupportsConjunctionContext)

	// ExitSupportsDisjunction is called when exiting the supportsDisjunction production.
	ExitSupportsDisjunction(c *SupportsDisjunctionContext)

	// ExitSupportsDeclarationCondition is called when exiting the supportsDeclarationCondition production.
	ExitSupportsDeclarationCondition(c *SupportsDeclarationConditionContext)

	// ExitGeneralEnclosed is called when exiting the generalEnclosed production.
	ExitGeneralEnclosed(c *GeneralEnclosedContext)

	// ExitVar_ is called when exiting the var_ production.
	ExitVar_(c *Var_Context)

	// ExitCalc is called when exiting the calc production.
	ExitCalc(c *CalcContext)

	// ExitCalcSum is called when exiting the calcSum production.
	ExitCalcSum(c *CalcSumContext)

	// ExitCalcProduct is called when exiting the calcProduct production.
	ExitCalcProduct(c *CalcProductContext)

	// ExitCalcValue is called when exiting the calcValue production.
	ExitCalcValue(c *CalcValueContext)

	// ExitFontFaceRule is called when exiting the fontFaceRule production.
	ExitFontFaceRule(c *FontFaceRuleContext)

	// ExitKnownFontFaceDeclaration is called when exiting the knownFontFaceDeclaration production.
	ExitKnownFontFaceDeclaration(c *KnownFontFaceDeclarationContext)

	// ExitUnknownFontFaceDeclaration is called when exiting the unknownFontFaceDeclaration production.
	ExitUnknownFontFaceDeclaration(c *UnknownFontFaceDeclarationContext)

	// ExitKeyframesRule is called when exiting the keyframesRule production.
	ExitKeyframesRule(c *KeyframesRuleContext)

	// ExitKeyframesBlocks is called when exiting the keyframesBlocks production.
	ExitKeyframesBlocks(c *KeyframesBlocksContext)

	// ExitKeyframeSelector is called when exiting the keyframeSelector production.
	ExitKeyframeSelector(c *KeyframeSelectorContext)

	// ExitViewport is called when exiting the viewport production.
	ExitViewport(c *ViewportContext)

	// ExitCounterStyle is called when exiting the counterStyle production.
	ExitCounterStyle(c *CounterStyleContext)

	// ExitFontFeatureValuesRule is called when exiting the fontFeatureValuesRule production.
	ExitFontFeatureValuesRule(c *FontFeatureValuesRuleContext)

	// ExitFontFamilyNameList is called when exiting the fontFamilyNameList production.
	ExitFontFamilyNameList(c *FontFamilyNameListContext)

	// ExitFontFamilyName is called when exiting the fontFamilyName production.
	ExitFontFamilyName(c *FontFamilyNameContext)

	// ExitFeatureValueBlock is called when exiting the featureValueBlock production.
	ExitFeatureValueBlock(c *FeatureValueBlockContext)

	// ExitFeatureType is called when exiting the featureType production.
	ExitFeatureType(c *FeatureTypeContext)

	// ExitFeatureValueDefinition is called when exiting the featureValueDefinition production.
	ExitFeatureValueDefinition(c *FeatureValueDefinitionContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitWs is called when exiting the ws production.
	ExitWs(c *WsContext)
}
