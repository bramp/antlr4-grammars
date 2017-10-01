// Generated from ZParser.g4 by ANTLR 4.7.

package z // ZParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseZParserListener is a complete listener for a parse tree produced by ZParser.
type BaseZParserListener struct{}

var _ ZParserListener = &BaseZParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSpecification is called when production specification is entered.
func (s *BaseZParserListener) EnterSpecification(ctx *SpecificationContext) {}

// ExitSpecification is called when production specification is exited.
func (s *BaseZParserListener) ExitSpecification(ctx *SpecificationContext) {}

// EnterInheritingSection is called when production InheritingSection is entered.
func (s *BaseZParserListener) EnterInheritingSection(ctx *InheritingSectionContext) {}

// ExitInheritingSection is called when production InheritingSection is exited.
func (s *BaseZParserListener) ExitInheritingSection(ctx *InheritingSectionContext) {}

// EnterBaseSection is called when production BaseSection is entered.
func (s *BaseZParserListener) EnterBaseSection(ctx *BaseSectionContext) {}

// ExitBaseSection is called when production BaseSection is exited.
func (s *BaseZParserListener) ExitBaseSection(ctx *BaseSectionContext) {}

// EnterGivenTypesParagraph is called when production GivenTypesParagraph is entered.
func (s *BaseZParserListener) EnterGivenTypesParagraph(ctx *GivenTypesParagraphContext) {}

// ExitGivenTypesParagraph is called when production GivenTypesParagraph is exited.
func (s *BaseZParserListener) ExitGivenTypesParagraph(ctx *GivenTypesParagraphContext) {}

// EnterAxiomaticDescriptionParagraph is called when production AxiomaticDescriptionParagraph is entered.
func (s *BaseZParserListener) EnterAxiomaticDescriptionParagraph(ctx *AxiomaticDescriptionParagraphContext) {
}

// ExitAxiomaticDescriptionParagraph is called when production AxiomaticDescriptionParagraph is exited.
func (s *BaseZParserListener) ExitAxiomaticDescriptionParagraph(ctx *AxiomaticDescriptionParagraphContext) {
}

// EnterSchemaDefinitionParagraph is called when production SchemaDefinitionParagraph is entered.
func (s *BaseZParserListener) EnterSchemaDefinitionParagraph(ctx *SchemaDefinitionParagraphContext) {}

// ExitSchemaDefinitionParagraph is called when production SchemaDefinitionParagraph is exited.
func (s *BaseZParserListener) ExitSchemaDefinitionParagraph(ctx *SchemaDefinitionParagraphContext) {}

// EnterGenericAxiomaticDescriptionParagraph is called when production GenericAxiomaticDescriptionParagraph is entered.
func (s *BaseZParserListener) EnterGenericAxiomaticDescriptionParagraph(ctx *GenericAxiomaticDescriptionParagraphContext) {
}

// ExitGenericAxiomaticDescriptionParagraph is called when production GenericAxiomaticDescriptionParagraph is exited.
func (s *BaseZParserListener) ExitGenericAxiomaticDescriptionParagraph(ctx *GenericAxiomaticDescriptionParagraphContext) {
}

// EnterGenericSchemaDefinitionParagraph is called when production GenericSchemaDefinitionParagraph is entered.
func (s *BaseZParserListener) EnterGenericSchemaDefinitionParagraph(ctx *GenericSchemaDefinitionParagraphContext) {
}

// ExitGenericSchemaDefinitionParagraph is called when production GenericSchemaDefinitionParagraph is exited.
func (s *BaseZParserListener) ExitGenericSchemaDefinitionParagraph(ctx *GenericSchemaDefinitionParagraphContext) {
}

// EnterHorizontalDefinitionParagraph is called when production HorizontalDefinitionParagraph is entered.
func (s *BaseZParserListener) EnterHorizontalDefinitionParagraph(ctx *HorizontalDefinitionParagraphContext) {
}

// ExitHorizontalDefinitionParagraph is called when production HorizontalDefinitionParagraph is exited.
func (s *BaseZParserListener) ExitHorizontalDefinitionParagraph(ctx *HorizontalDefinitionParagraphContext) {
}

// EnterGenericHorizontalDefinitionParagraph is called when production GenericHorizontalDefinitionParagraph is entered.
func (s *BaseZParserListener) EnterGenericHorizontalDefinitionParagraph(ctx *GenericHorizontalDefinitionParagraphContext) {
}

// ExitGenericHorizontalDefinitionParagraph is called when production GenericHorizontalDefinitionParagraph is exited.
func (s *BaseZParserListener) ExitGenericHorizontalDefinitionParagraph(ctx *GenericHorizontalDefinitionParagraphContext) {
}

// EnterGenericOperatorDefinitionParagraph is called when production GenericOperatorDefinitionParagraph is entered.
func (s *BaseZParserListener) EnterGenericOperatorDefinitionParagraph(ctx *GenericOperatorDefinitionParagraphContext) {
}

// ExitGenericOperatorDefinitionParagraph is called when production GenericOperatorDefinitionParagraph is exited.
func (s *BaseZParserListener) ExitGenericOperatorDefinitionParagraph(ctx *GenericOperatorDefinitionParagraphContext) {
}

// EnterFreeTypesParagraph is called when production FreeTypesParagraph is entered.
func (s *BaseZParserListener) EnterFreeTypesParagraph(ctx *FreeTypesParagraphContext) {}

// ExitFreeTypesParagraph is called when production FreeTypesParagraph is exited.
func (s *BaseZParserListener) ExitFreeTypesParagraph(ctx *FreeTypesParagraphContext) {}

// EnterConjectureParagraph is called when production ConjectureParagraph is entered.
func (s *BaseZParserListener) EnterConjectureParagraph(ctx *ConjectureParagraphContext) {}

// ExitConjectureParagraph is called when production ConjectureParagraph is exited.
func (s *BaseZParserListener) ExitConjectureParagraph(ctx *ConjectureParagraphContext) {}

// EnterGenericConjectureParagraph is called when production GenericConjectureParagraph is entered.
func (s *BaseZParserListener) EnterGenericConjectureParagraph(ctx *GenericConjectureParagraphContext) {
}

// ExitGenericConjectureParagraph is called when production GenericConjectureParagraph is exited.
func (s *BaseZParserListener) ExitGenericConjectureParagraph(ctx *GenericConjectureParagraphContext) {}

// EnterOperatorTemplateParagraph is called when production OperatorTemplateParagraph is entered.
func (s *BaseZParserListener) EnterOperatorTemplateParagraph(ctx *OperatorTemplateParagraphContext) {}

// ExitOperatorTemplateParagraph is called when production OperatorTemplateParagraph is exited.
func (s *BaseZParserListener) ExitOperatorTemplateParagraph(ctx *OperatorTemplateParagraphContext) {}

// EnterFreetype is called when production freetype is entered.
func (s *BaseZParserListener) EnterFreetype(ctx *FreetypeContext) {}

// ExitFreetype is called when production freetype is exited.
func (s *BaseZParserListener) ExitFreetype(ctx *FreetypeContext) {}

// EnterBranch is called when production branch is entered.
func (s *BaseZParserListener) EnterBranch(ctx *BranchContext) {}

// ExitBranch is called when production branch is exited.
func (s *BaseZParserListener) ExitBranch(ctx *BranchContext) {}

// EnterFormals is called when production formals is entered.
func (s *BaseZParserListener) EnterFormals(ctx *FormalsContext) {}

// ExitFormals is called when production formals is exited.
func (s *BaseZParserListener) ExitFormals(ctx *FormalsContext) {}

// EnterExistentialQuantificationPredicate is called when production ExistentialQuantificationPredicate is entered.
func (s *BaseZParserListener) EnterExistentialQuantificationPredicate(ctx *ExistentialQuantificationPredicateContext) {
}

// ExitExistentialQuantificationPredicate is called when production ExistentialQuantificationPredicate is exited.
func (s *BaseZParserListener) ExitExistentialQuantificationPredicate(ctx *ExistentialQuantificationPredicateContext) {
}

// EnterConjunctionPredicate is called when production ConjunctionPredicate is entered.
func (s *BaseZParserListener) EnterConjunctionPredicate(ctx *ConjunctionPredicateContext) {}

// ExitConjunctionPredicate is called when production ConjunctionPredicate is exited.
func (s *BaseZParserListener) ExitConjunctionPredicate(ctx *ConjunctionPredicateContext) {}

// EnterEquivalencePredicate is called when production EquivalencePredicate is entered.
func (s *BaseZParserListener) EnterEquivalencePredicate(ctx *EquivalencePredicateContext) {}

// ExitEquivalencePredicate is called when production EquivalencePredicate is exited.
func (s *BaseZParserListener) ExitEquivalencePredicate(ctx *EquivalencePredicateContext) {}

// EnterNewlineConjunctionPredicate is called when production NewlineConjunctionPredicate is entered.
func (s *BaseZParserListener) EnterNewlineConjunctionPredicate(ctx *NewlineConjunctionPredicateContext) {
}

// ExitNewlineConjunctionPredicate is called when production NewlineConjunctionPredicate is exited.
func (s *BaseZParserListener) ExitNewlineConjunctionPredicate(ctx *NewlineConjunctionPredicateContext) {
}

// EnterImplicationPredicate is called when production ImplicationPredicate is entered.
func (s *BaseZParserListener) EnterImplicationPredicate(ctx *ImplicationPredicateContext) {}

// ExitImplicationPredicate is called when production ImplicationPredicate is exited.
func (s *BaseZParserListener) ExitImplicationPredicate(ctx *ImplicationPredicateContext) {}

// EnterRelationOperatorApplicationPredicate is called when production RelationOperatorApplicationPredicate is entered.
func (s *BaseZParserListener) EnterRelationOperatorApplicationPredicate(ctx *RelationOperatorApplicationPredicateContext) {
}

// ExitRelationOperatorApplicationPredicate is called when production RelationOperatorApplicationPredicate is exited.
func (s *BaseZParserListener) ExitRelationOperatorApplicationPredicate(ctx *RelationOperatorApplicationPredicateContext) {
}

// EnterUniversalQuantificationPredicate is called when production UniversalQuantificationPredicate is entered.
func (s *BaseZParserListener) EnterUniversalQuantificationPredicate(ctx *UniversalQuantificationPredicateContext) {
}

// ExitUniversalQuantificationPredicate is called when production UniversalQuantificationPredicate is exited.
func (s *BaseZParserListener) ExitUniversalQuantificationPredicate(ctx *UniversalQuantificationPredicateContext) {
}

// EnterTruthPredicate is called when production TruthPredicate is entered.
func (s *BaseZParserListener) EnterTruthPredicate(ctx *TruthPredicateContext) {}

// ExitTruthPredicate is called when production TruthPredicate is exited.
func (s *BaseZParserListener) ExitTruthPredicate(ctx *TruthPredicateContext) {}

// EnterFalsityPredicate is called when production FalsityPredicate is entered.
func (s *BaseZParserListener) EnterFalsityPredicate(ctx *FalsityPredicateContext) {}

// ExitFalsityPredicate is called when production FalsityPredicate is exited.
func (s *BaseZParserListener) ExitFalsityPredicate(ctx *FalsityPredicateContext) {}

// EnterParenthesizedPredicate is called when production ParenthesizedPredicate is entered.
func (s *BaseZParserListener) EnterParenthesizedPredicate(ctx *ParenthesizedPredicateContext) {}

// ExitParenthesizedPredicate is called when production ParenthesizedPredicate is exited.
func (s *BaseZParserListener) ExitParenthesizedPredicate(ctx *ParenthesizedPredicateContext) {}

// EnterDisjunctionPredicate is called when production DisjunctionPredicate is entered.
func (s *BaseZParserListener) EnterDisjunctionPredicate(ctx *DisjunctionPredicateContext) {}

// ExitDisjunctionPredicate is called when production DisjunctionPredicate is exited.
func (s *BaseZParserListener) ExitDisjunctionPredicate(ctx *DisjunctionPredicateContext) {}

// EnterSemicolonConjunctionPredicate is called when production SemicolonConjunctionPredicate is entered.
func (s *BaseZParserListener) EnterSemicolonConjunctionPredicate(ctx *SemicolonConjunctionPredicateContext) {
}

// ExitSemicolonConjunctionPredicate is called when production SemicolonConjunctionPredicate is exited.
func (s *BaseZParserListener) ExitSemicolonConjunctionPredicate(ctx *SemicolonConjunctionPredicateContext) {
}

// EnterSchemaPredicatePredicate is called when production SchemaPredicatePredicate is entered.
func (s *BaseZParserListener) EnterSchemaPredicatePredicate(ctx *SchemaPredicatePredicateContext) {}

// ExitSchemaPredicatePredicate is called when production SchemaPredicatePredicate is exited.
func (s *BaseZParserListener) ExitSchemaPredicatePredicate(ctx *SchemaPredicatePredicateContext) {}

// EnterUniqueExistentialQuantificationPredicate is called when production UniqueExistentialQuantificationPredicate is entered.
func (s *BaseZParserListener) EnterUniqueExistentialQuantificationPredicate(ctx *UniqueExistentialQuantificationPredicateContext) {
}

// ExitUniqueExistentialQuantificationPredicate is called when production UniqueExistentialQuantificationPredicate is exited.
func (s *BaseZParserListener) ExitUniqueExistentialQuantificationPredicate(ctx *UniqueExistentialQuantificationPredicateContext) {
}

// EnterNegationPredicate is called when production NegationPredicate is entered.
func (s *BaseZParserListener) EnterNegationPredicate(ctx *NegationPredicateContext) {}

// ExitNegationPredicate is called when production NegationPredicate is exited.
func (s *BaseZParserListener) ExitNegationPredicate(ctx *NegationPredicateContext) {}

// EnterSetComprehensionExpression is called when production SetComprehensionExpression is entered.
func (s *BaseZParserListener) EnterSetComprehensionExpression(ctx *SetComprehensionExpressionContext) {
}

// ExitSetComprehensionExpression is called when production SetComprehensionExpression is exited.
func (s *BaseZParserListener) ExitSetComprehensionExpression(ctx *SetComprehensionExpressionContext) {}

// EnterSchemaEquivalenceExpression is called when production SchemaEquivalenceExpression is entered.
func (s *BaseZParserListener) EnterSchemaEquivalenceExpression(ctx *SchemaEquivalenceExpressionContext) {
}

// ExitSchemaEquivalenceExpression is called when production SchemaEquivalenceExpression is exited.
func (s *BaseZParserListener) ExitSchemaEquivalenceExpression(ctx *SchemaEquivalenceExpressionContext) {
}

// EnterNofixApplicationExpression is called when production NofixApplicationExpression is entered.
func (s *BaseZParserListener) EnterNofixApplicationExpression(ctx *NofixApplicationExpressionContext) {
}

// ExitNofixApplicationExpression is called when production NofixApplicationExpression is exited.
func (s *BaseZParserListener) ExitNofixApplicationExpression(ctx *NofixApplicationExpressionContext) {}

// EnterSchemaConstructionExpression is called when production SchemaConstructionExpression is entered.
func (s *BaseZParserListener) EnterSchemaConstructionExpression(ctx *SchemaConstructionExpressionContext) {
}

// ExitSchemaConstructionExpression is called when production SchemaConstructionExpression is exited.
func (s *BaseZParserListener) ExitSchemaConstructionExpression(ctx *SchemaConstructionExpressionContext) {
}

// EnterGenericPostfixApplicationExpression is called when production GenericPostfixApplicationExpression is entered.
func (s *BaseZParserListener) EnterGenericPostfixApplicationExpression(ctx *GenericPostfixApplicationExpressionContext) {
}

// ExitGenericPostfixApplicationExpression is called when production GenericPostfixApplicationExpression is exited.
func (s *BaseZParserListener) ExitGenericPostfixApplicationExpression(ctx *GenericPostfixApplicationExpressionContext) {
}

// EnterGenericPrefixApplicationExpression is called when production GenericPrefixApplicationExpression is entered.
func (s *BaseZParserListener) EnterGenericPrefixApplicationExpression(ctx *GenericPrefixApplicationExpressionContext) {
}

// ExitGenericPrefixApplicationExpression is called when production GenericPrefixApplicationExpression is exited.
func (s *BaseZParserListener) ExitGenericPrefixApplicationExpression(ctx *GenericPrefixApplicationExpressionContext) {
}

// EnterCharacteristicDefiniteDescriptionExpression is called when production CharacteristicDefiniteDescriptionExpression is entered.
func (s *BaseZParserListener) EnterCharacteristicDefiniteDescriptionExpression(ctx *CharacteristicDefiniteDescriptionExpressionContext) {
}

// ExitCharacteristicDefiniteDescriptionExpression is called when production CharacteristicDefiniteDescriptionExpression is exited.
func (s *BaseZParserListener) ExitCharacteristicDefiniteDescriptionExpression(ctx *CharacteristicDefiniteDescriptionExpressionContext) {
}

// EnterSchemaUniversalQuantificationExpression is called when production SchemaUniversalQuantificationExpression is entered.
func (s *BaseZParserListener) EnterSchemaUniversalQuantificationExpression(ctx *SchemaUniversalQuantificationExpressionContext) {
}

// ExitSchemaUniversalQuantificationExpression is called when production SchemaUniversalQuantificationExpression is exited.
func (s *BaseZParserListener) ExitSchemaUniversalQuantificationExpression(ctx *SchemaUniversalQuantificationExpressionContext) {
}

// EnterSetExtensionExpression is called when production SetExtensionExpression is entered.
func (s *BaseZParserListener) EnterSetExtensionExpression(ctx *SetExtensionExpressionContext) {}

// ExitSetExtensionExpression is called when production SetExtensionExpression is exited.
func (s *BaseZParserListener) ExitSetExtensionExpression(ctx *SetExtensionExpressionContext) {}

// EnterFunctionConstructionExpression is called when production FunctionConstructionExpression is entered.
func (s *BaseZParserListener) EnterFunctionConstructionExpression(ctx *FunctionConstructionExpressionContext) {
}

// ExitFunctionConstructionExpression is called when production FunctionConstructionExpression is exited.
func (s *BaseZParserListener) ExitFunctionConstructionExpression(ctx *FunctionConstructionExpressionContext) {
}

// EnterSchemaNegationExpression is called when production SchemaNegationExpression is entered.
func (s *BaseZParserListener) EnterSchemaNegationExpression(ctx *SchemaNegationExpressionContext) {}

// ExitSchemaNegationExpression is called when production SchemaNegationExpression is exited.
func (s *BaseZParserListener) ExitSchemaNegationExpression(ctx *SchemaNegationExpressionContext) {}

// EnterTupleExtensionExpression is called when production TupleExtensionExpression is entered.
func (s *BaseZParserListener) EnterTupleExtensionExpression(ctx *TupleExtensionExpressionContext) {}

// ExitTupleExtensionExpression is called when production TupleExtensionExpression is exited.
func (s *BaseZParserListener) ExitTupleExtensionExpression(ctx *TupleExtensionExpressionContext) {}

// EnterSchemaCompositionExpression is called when production SchemaCompositionExpression is entered.
func (s *BaseZParserListener) EnterSchemaCompositionExpression(ctx *SchemaCompositionExpressionContext) {
}

// ExitSchemaCompositionExpression is called when production SchemaCompositionExpression is exited.
func (s *BaseZParserListener) ExitSchemaCompositionExpression(ctx *SchemaCompositionExpressionContext) {
}

// EnterPowersetExpression is called when production PowersetExpression is entered.
func (s *BaseZParserListener) EnterPowersetExpression(ctx *PowersetExpressionContext) {}

// ExitPowersetExpression is called when production PowersetExpression is exited.
func (s *BaseZParserListener) ExitPowersetExpression(ctx *PowersetExpressionContext) {}

// EnterSchemaDisjunctionExpression is called when production SchemaDisjunctionExpression is entered.
func (s *BaseZParserListener) EnterSchemaDisjunctionExpression(ctx *SchemaDisjunctionExpressionContext) {
}

// ExitSchemaDisjunctionExpression is called when production SchemaDisjunctionExpression is exited.
func (s *BaseZParserListener) ExitSchemaDisjunctionExpression(ctx *SchemaDisjunctionExpressionContext) {
}

// EnterSubstitutionExpressionExpression is called when production SubstitutionExpressionExpression is entered.
func (s *BaseZParserListener) EnterSubstitutionExpressionExpression(ctx *SubstitutionExpressionExpressionContext) {
}

// ExitSubstitutionExpressionExpression is called when production SubstitutionExpressionExpression is exited.
func (s *BaseZParserListener) ExitSubstitutionExpressionExpression(ctx *SubstitutionExpressionExpressionContext) {
}

// EnterSchemaHidingExpression is called when production SchemaHidingExpression is entered.
func (s *BaseZParserListener) EnterSchemaHidingExpression(ctx *SchemaHidingExpressionContext) {}

// ExitSchemaHidingExpression is called when production SchemaHidingExpression is exited.
func (s *BaseZParserListener) ExitSchemaHidingExpression(ctx *SchemaHidingExpressionContext) {}

// EnterPostfixApplicationExpression is called when production PostfixApplicationExpression is entered.
func (s *BaseZParserListener) EnterPostfixApplicationExpression(ctx *PostfixApplicationExpressionContext) {
}

// ExitPostfixApplicationExpression is called when production PostfixApplicationExpression is exited.
func (s *BaseZParserListener) ExitPostfixApplicationExpression(ctx *PostfixApplicationExpressionContext) {
}

// EnterCartesianProductExpression is called when production CartesianProductExpression is entered.
func (s *BaseZParserListener) EnterCartesianProductExpression(ctx *CartesianProductExpressionContext) {
}

// ExitCartesianProductExpression is called when production CartesianProductExpression is exited.
func (s *BaseZParserListener) ExitCartesianProductExpression(ctx *CartesianProductExpressionContext) {}

// EnterBindingSelectionExpression is called when production BindingSelectionExpression is entered.
func (s *BaseZParserListener) EnterBindingSelectionExpression(ctx *BindingSelectionExpressionContext) {
}

// ExitBindingSelectionExpression is called when production BindingSelectionExpression is exited.
func (s *BaseZParserListener) ExitBindingSelectionExpression(ctx *BindingSelectionExpressionContext) {}

// EnterGenericInstantiationExpression is called when production GenericInstantiationExpression is entered.
func (s *BaseZParserListener) EnterGenericInstantiationExpression(ctx *GenericInstantiationExpressionContext) {
}

// ExitGenericInstantiationExpression is called when production GenericInstantiationExpression is exited.
func (s *BaseZParserListener) ExitGenericInstantiationExpression(ctx *GenericInstantiationExpressionContext) {
}

// EnterSchemaExistentialQuantificationExpression is called when production SchemaExistentialQuantificationExpression is entered.
func (s *BaseZParserListener) EnterSchemaExistentialQuantificationExpression(ctx *SchemaExistentialQuantificationExpressionContext) {
}

// ExitSchemaExistentialQuantificationExpression is called when production SchemaExistentialQuantificationExpression is exited.
func (s *BaseZParserListener) ExitSchemaExistentialQuantificationExpression(ctx *SchemaExistentialQuantificationExpressionContext) {
}

// EnterNumberLiteralExpression is called when production NumberLiteralExpression is entered.
func (s *BaseZParserListener) EnterNumberLiteralExpression(ctx *NumberLiteralExpressionContext) {}

// ExitNumberLiteralExpression is called when production NumberLiteralExpression is exited.
func (s *BaseZParserListener) ExitNumberLiteralExpression(ctx *NumberLiteralExpressionContext) {}

// EnterParenthesizedExpression is called when production ParenthesizedExpression is entered.
func (s *BaseZParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// ExitParenthesizedExpression is called when production ParenthesizedExpression is exited.
func (s *BaseZParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {}

// EnterInfixLeftApplicationExpression is called when production InfixLeftApplicationExpression is entered.
func (s *BaseZParserListener) EnterInfixLeftApplicationExpression(ctx *InfixLeftApplicationExpressionContext) {
}

// ExitInfixLeftApplicationExpression is called when production InfixLeftApplicationExpression is exited.
func (s *BaseZParserListener) ExitInfixLeftApplicationExpression(ctx *InfixLeftApplicationExpressionContext) {
}

// EnterApplicationExpression is called when production ApplicationExpression is entered.
func (s *BaseZParserListener) EnterApplicationExpression(ctx *ApplicationExpressionContext) {}

// ExitApplicationExpression is called when production ApplicationExpression is exited.
func (s *BaseZParserListener) ExitApplicationExpression(ctx *ApplicationExpressionContext) {}

// EnterSchemaRenamingExpression is called when production SchemaRenamingExpression is entered.
func (s *BaseZParserListener) EnterSchemaRenamingExpression(ctx *SchemaRenamingExpressionContext) {}

// ExitSchemaRenamingExpression is called when production SchemaRenamingExpression is exited.
func (s *BaseZParserListener) ExitSchemaRenamingExpression(ctx *SchemaRenamingExpressionContext) {}

// EnterSchemaConjunctionExpression is called when production SchemaConjunctionExpression is entered.
func (s *BaseZParserListener) EnterSchemaConjunctionExpression(ctx *SchemaConjunctionExpressionContext) {
}

// ExitSchemaConjunctionExpression is called when production SchemaConjunctionExpression is exited.
func (s *BaseZParserListener) ExitSchemaConjunctionExpression(ctx *SchemaConjunctionExpressionContext) {
}

// EnterConditionalExpression is called when production ConditionalExpression is entered.
func (s *BaseZParserListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production ConditionalExpression is exited.
func (s *BaseZParserListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterSchemaProjectionExpression is called when production SchemaProjectionExpression is entered.
func (s *BaseZParserListener) EnterSchemaProjectionExpression(ctx *SchemaProjectionExpressionContext) {
}

// ExitSchemaProjectionExpression is called when production SchemaProjectionExpression is exited.
func (s *BaseZParserListener) ExitSchemaProjectionExpression(ctx *SchemaProjectionExpressionContext) {}

// EnterInfixRightApplicationExpression is called when production InfixRightApplicationExpression is entered.
func (s *BaseZParserListener) EnterInfixRightApplicationExpression(ctx *InfixRightApplicationExpressionContext) {
}

// ExitInfixRightApplicationExpression is called when production InfixRightApplicationExpression is exited.
func (s *BaseZParserListener) ExitInfixRightApplicationExpression(ctx *InfixRightApplicationExpressionContext) {
}

// EnterCharacteristicSetComprehensionExpression is called when production CharacteristicSetComprehensionExpression is entered.
func (s *BaseZParserListener) EnterCharacteristicSetComprehensionExpression(ctx *CharacteristicSetComprehensionExpressionContext) {
}

// ExitCharacteristicSetComprehensionExpression is called when production CharacteristicSetComprehensionExpression is exited.
func (s *BaseZParserListener) ExitCharacteristicSetComprehensionExpression(ctx *CharacteristicSetComprehensionExpressionContext) {
}

// EnterSchemaPreconditionExpression is called when production SchemaPreconditionExpression is entered.
func (s *BaseZParserListener) EnterSchemaPreconditionExpression(ctx *SchemaPreconditionExpressionContext) {
}

// ExitSchemaPreconditionExpression is called when production SchemaPreconditionExpression is exited.
func (s *BaseZParserListener) ExitSchemaPreconditionExpression(ctx *SchemaPreconditionExpressionContext) {
}

// EnterPrefixApplicationExpression is called when production PrefixApplicationExpression is entered.
func (s *BaseZParserListener) EnterPrefixApplicationExpression(ctx *PrefixApplicationExpressionContext) {
}

// ExitPrefixApplicationExpression is called when production PrefixApplicationExpression is exited.
func (s *BaseZParserListener) ExitPrefixApplicationExpression(ctx *PrefixApplicationExpressionContext) {
}

// EnterBindingConstructionExpression is called when production BindingConstructionExpression is entered.
func (s *BaseZParserListener) EnterBindingConstructionExpression(ctx *BindingConstructionExpressionContext) {
}

// ExitBindingConstructionExpression is called when production BindingConstructionExpression is exited.
func (s *BaseZParserListener) ExitBindingConstructionExpression(ctx *BindingConstructionExpressionContext) {
}

// EnterSchemaPipingExpression is called when production SchemaPipingExpression is entered.
func (s *BaseZParserListener) EnterSchemaPipingExpression(ctx *SchemaPipingExpressionContext) {}

// ExitSchemaPipingExpression is called when production SchemaPipingExpression is exited.
func (s *BaseZParserListener) ExitSchemaPipingExpression(ctx *SchemaPipingExpressionContext) {}

// EnterSchemaImplicationExpression is called when production SchemaImplicationExpression is entered.
func (s *BaseZParserListener) EnterSchemaImplicationExpression(ctx *SchemaImplicationExpressionContext) {
}

// ExitSchemaImplicationExpression is called when production SchemaImplicationExpression is exited.
func (s *BaseZParserListener) ExitSchemaImplicationExpression(ctx *SchemaImplicationExpressionContext) {
}

// EnterBindingExtensionExpression is called when production BindingExtensionExpression is entered.
func (s *BaseZParserListener) EnterBindingExtensionExpression(ctx *BindingExtensionExpressionContext) {
}

// ExitBindingExtensionExpression is called when production BindingExtensionExpression is exited.
func (s *BaseZParserListener) ExitBindingExtensionExpression(ctx *BindingExtensionExpressionContext) {}

// EnterSchemaDecorationExpression is called when production SchemaDecorationExpression is entered.
func (s *BaseZParserListener) EnterSchemaDecorationExpression(ctx *SchemaDecorationExpressionContext) {
}

// ExitSchemaDecorationExpression is called when production SchemaDecorationExpression is exited.
func (s *BaseZParserListener) ExitSchemaDecorationExpression(ctx *SchemaDecorationExpressionContext) {}

// EnterSchemaUniqueExistentialQuantificationExpression is called when production SchemaUniqueExistentialQuantificationExpression is entered.
func (s *BaseZParserListener) EnterSchemaUniqueExistentialQuantificationExpression(ctx *SchemaUniqueExistentialQuantificationExpressionContext) {
}

// ExitSchemaUniqueExistentialQuantificationExpression is called when production SchemaUniqueExistentialQuantificationExpression is exited.
func (s *BaseZParserListener) ExitSchemaUniqueExistentialQuantificationExpression(ctx *SchemaUniqueExistentialQuantificationExpressionContext) {
}

// EnterDefiniteDescriptionExpression is called when production DefiniteDescriptionExpression is entered.
func (s *BaseZParserListener) EnterDefiniteDescriptionExpression(ctx *DefiniteDescriptionExpressionContext) {
}

// ExitDefiniteDescriptionExpression is called when production DefiniteDescriptionExpression is exited.
func (s *BaseZParserListener) ExitDefiniteDescriptionExpression(ctx *DefiniteDescriptionExpressionContext) {
}

// EnterReferenceExpression is called when production ReferenceExpression is entered.
func (s *BaseZParserListener) EnterReferenceExpression(ctx *ReferenceExpressionContext) {}

// ExitReferenceExpression is called when production ReferenceExpression is exited.
func (s *BaseZParserListener) ExitReferenceExpression(ctx *ReferenceExpressionContext) {}

// EnterGenericInfixApplicationExpression is called when production GenericInfixApplicationExpression is entered.
func (s *BaseZParserListener) EnterGenericInfixApplicationExpression(ctx *GenericInfixApplicationExpressionContext) {
}

// ExitGenericInfixApplicationExpression is called when production GenericInfixApplicationExpression is exited.
func (s *BaseZParserListener) ExitGenericInfixApplicationExpression(ctx *GenericInfixApplicationExpressionContext) {
}

// EnterTupleSelectionExpression is called when production TupleSelectionExpression is entered.
func (s *BaseZParserListener) EnterTupleSelectionExpression(ctx *TupleSelectionExpressionContext) {}

// ExitTupleSelectionExpression is called when production TupleSelectionExpression is exited.
func (s *BaseZParserListener) ExitTupleSelectionExpression(ctx *TupleSelectionExpressionContext) {}

// EnterSchemaText is called when production schemaText is entered.
func (s *BaseZParserListener) EnterSchemaText(ctx *SchemaTextContext) {}

// ExitSchemaText is called when production schemaText is exited.
func (s *BaseZParserListener) ExitSchemaText(ctx *SchemaTextContext) {}

// EnterDeclPart is called when production declPart is entered.
func (s *BaseZParserListener) EnterDeclPart(ctx *DeclPartContext) {}

// ExitDeclPart is called when production declPart is exited.
func (s *BaseZParserListener) ExitDeclPart(ctx *DeclPartContext) {}

// EnterDeclNameExpression is called when production declNameExpression is entered.
func (s *BaseZParserListener) EnterDeclNameExpression(ctx *DeclNameExpressionContext) {}

// ExitDeclNameExpression is called when production declNameExpression is exited.
func (s *BaseZParserListener) ExitDeclNameExpression(ctx *DeclNameExpressionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseZParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseZParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterOperatorTemplate is called when production operatorTemplate is entered.
func (s *BaseZParserListener) EnterOperatorTemplate(ctx *OperatorTemplateContext) {}

// ExitOperatorTemplate is called when production operatorTemplate is exited.
func (s *BaseZParserListener) ExitOperatorTemplate(ctx *OperatorTemplateContext) {}

// EnterCategoryTemplate is called when production categoryTemplate is entered.
func (s *BaseZParserListener) EnterCategoryTemplate(ctx *CategoryTemplateContext) {}

// ExitCategoryTemplate is called when production categoryTemplate is exited.
func (s *BaseZParserListener) ExitCategoryTemplate(ctx *CategoryTemplateContext) {}

// EnterPrec is called when production prec is entered.
func (s *BaseZParserListener) EnterPrec(ctx *PrecContext) {}

// ExitPrec is called when production prec is exited.
func (s *BaseZParserListener) ExitPrec(ctx *PrecContext) {}

// EnterAssoc is called when production assoc is entered.
func (s *BaseZParserListener) EnterAssoc(ctx *AssocContext) {}

// ExitAssoc is called when production assoc is exited.
func (s *BaseZParserListener) ExitAssoc(ctx *AssocContext) {}

// EnterTemplate is called when production template is entered.
func (s *BaseZParserListener) EnterTemplate(ctx *TemplateContext) {}

// ExitTemplate is called when production template is exited.
func (s *BaseZParserListener) ExitTemplate(ctx *TemplateContext) {}

// EnterPrefixTemplate is called when production prefixTemplate is entered.
func (s *BaseZParserListener) EnterPrefixTemplate(ctx *PrefixTemplateContext) {}

// ExitPrefixTemplate is called when production prefixTemplate is exited.
func (s *BaseZParserListener) ExitPrefixTemplate(ctx *PrefixTemplateContext) {}

// EnterPostfixTemplate is called when production postfixTemplate is entered.
func (s *BaseZParserListener) EnterPostfixTemplate(ctx *PostfixTemplateContext) {}

// ExitPostfixTemplate is called when production postfixTemplate is exited.
func (s *BaseZParserListener) ExitPostfixTemplate(ctx *PostfixTemplateContext) {}

// EnterInfixTemplate is called when production infixTemplate is entered.
func (s *BaseZParserListener) EnterInfixTemplate(ctx *InfixTemplateContext) {}

// ExitInfixTemplate is called when production infixTemplate is exited.
func (s *BaseZParserListener) ExitInfixTemplate(ctx *InfixTemplateContext) {}

// EnterNofixTemplate is called when production nofixTemplate is entered.
func (s *BaseZParserListener) EnterNofixTemplate(ctx *NofixTemplateContext) {}

// ExitNofixTemplate is called when production nofixTemplate is exited.
func (s *BaseZParserListener) ExitNofixTemplate(ctx *NofixTemplateContext) {}

// EnterDeclName is called when production declName is entered.
func (s *BaseZParserListener) EnterDeclName(ctx *DeclNameContext) {}

// ExitDeclName is called when production declName is exited.
func (s *BaseZParserListener) ExitDeclName(ctx *DeclNameContext) {}

// EnterRefName is called when production refName is entered.
func (s *BaseZParserListener) EnterRefName(ctx *RefNameContext) {}

// ExitRefName is called when production refName is exited.
func (s *BaseZParserListener) ExitRefName(ctx *RefNameContext) {}

// EnterOpName is called when production opName is entered.
func (s *BaseZParserListener) EnterOpName(ctx *OpNameContext) {}

// ExitOpName is called when production opName is exited.
func (s *BaseZParserListener) ExitOpName(ctx *OpNameContext) {}

// EnterPrefixName is called when production prefixName is entered.
func (s *BaseZParserListener) EnterPrefixName(ctx *PrefixNameContext) {}

// ExitPrefixName is called when production prefixName is exited.
func (s *BaseZParserListener) ExitPrefixName(ctx *PrefixNameContext) {}

// EnterPostfixName is called when production postfixName is entered.
func (s *BaseZParserListener) EnterPostfixName(ctx *PostfixNameContext) {}

// ExitPostfixName is called when production postfixName is exited.
func (s *BaseZParserListener) ExitPostfixName(ctx *PostfixNameContext) {}

// EnterInfixName is called when production infixName is entered.
func (s *BaseZParserListener) EnterInfixName(ctx *InfixNameContext) {}

// ExitInfixName is called when production infixName is exited.
func (s *BaseZParserListener) ExitInfixName(ctx *InfixNameContext) {}

// EnterNofixName is called when production nofixName is entered.
func (s *BaseZParserListener) EnterNofixName(ctx *NofixNameContext) {}

// ExitNofixName is called when production nofixName is exited.
func (s *BaseZParserListener) ExitNofixName(ctx *NofixNameContext) {}

// EnterGenName is called when production genName is entered.
func (s *BaseZParserListener) EnterGenName(ctx *GenNameContext) {}

// ExitGenName is called when production genName is exited.
func (s *BaseZParserListener) ExitGenName(ctx *GenNameContext) {}

// EnterPrefixGenName is called when production prefixGenName is entered.
func (s *BaseZParserListener) EnterPrefixGenName(ctx *PrefixGenNameContext) {}

// ExitPrefixGenName is called when production prefixGenName is exited.
func (s *BaseZParserListener) ExitPrefixGenName(ctx *PrefixGenNameContext) {}

// EnterPostfixGenName is called when production postfixGenName is entered.
func (s *BaseZParserListener) EnterPostfixGenName(ctx *PostfixGenNameContext) {}

// ExitPostfixGenName is called when production postfixGenName is exited.
func (s *BaseZParserListener) ExitPostfixGenName(ctx *PostfixGenNameContext) {}

// EnterInfixGenName is called when production infixGenName is entered.
func (s *BaseZParserListener) EnterInfixGenName(ctx *InfixGenNameContext) {}

// ExitInfixGenName is called when production infixGenName is exited.
func (s *BaseZParserListener) ExitInfixGenName(ctx *InfixGenNameContext) {}

// EnterNofixGenName is called when production nofixGenName is entered.
func (s *BaseZParserListener) EnterNofixGenName(ctx *NofixGenNameContext) {}

// ExitNofixGenName is called when production nofixGenName is exited.
func (s *BaseZParserListener) ExitNofixGenName(ctx *NofixGenNameContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseZParserListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseZParserListener) ExitRelation(ctx *RelationContext) {}

// EnterPrefixRel is called when production prefixRel is entered.
func (s *BaseZParserListener) EnterPrefixRel(ctx *PrefixRelContext) {}

// ExitPrefixRel is called when production prefixRel is exited.
func (s *BaseZParserListener) ExitPrefixRel(ctx *PrefixRelContext) {}

// EnterPostfixRel is called when production postfixRel is entered.
func (s *BaseZParserListener) EnterPostfixRel(ctx *PostfixRelContext) {}

// ExitPostfixRel is called when production postfixRel is exited.
func (s *BaseZParserListener) ExitPostfixRel(ctx *PostfixRelContext) {}

// EnterInfixRel is called when production infixRel is entered.
func (s *BaseZParserListener) EnterInfixRel(ctx *InfixRelContext) {}

// ExitInfixRel is called when production infixRel is exited.
func (s *BaseZParserListener) ExitInfixRel(ctx *InfixRelContext) {}

// EnterNofixRel is called when production nofixRel is entered.
func (s *BaseZParserListener) EnterNofixRel(ctx *NofixRelContext) {}

// ExitNofixRel is called when production nofixRel is exited.
func (s *BaseZParserListener) ExitNofixRel(ctx *NofixRelContext) {}

// EnterApplication is called when production application is entered.
func (s *BaseZParserListener) EnterApplication(ctx *ApplicationContext) {}

// ExitApplication is called when production application is exited.
func (s *BaseZParserListener) ExitApplication(ctx *ApplicationContext) {}

// EnterPrefixApp is called when production prefixApp is entered.
func (s *BaseZParserListener) EnterPrefixApp(ctx *PrefixAppContext) {}

// ExitPrefixApp is called when production prefixApp is exited.
func (s *BaseZParserListener) ExitPrefixApp(ctx *PrefixAppContext) {}

// EnterPostfixApp is called when production postfixApp is entered.
func (s *BaseZParserListener) EnterPostfixApp(ctx *PostfixAppContext) {}

// ExitPostfixApp is called when production postfixApp is exited.
func (s *BaseZParserListener) ExitPostfixApp(ctx *PostfixAppContext) {}

// EnterInfixApp is called when production infixApp is entered.
func (s *BaseZParserListener) EnterInfixApp(ctx *InfixAppContext) {}

// ExitInfixApp is called when production infixApp is exited.
func (s *BaseZParserListener) ExitInfixApp(ctx *InfixAppContext) {}

// EnterNofixApp is called when production nofixApp is entered.
func (s *BaseZParserListener) EnterNofixApp(ctx *NofixAppContext) {}

// ExitNofixApp is called when production nofixApp is exited.
func (s *BaseZParserListener) ExitNofixApp(ctx *NofixAppContext) {}

// EnterExpSep is called when production expSep is entered.
func (s *BaseZParserListener) EnterExpSep(ctx *ExpSepContext) {}

// ExitExpSep is called when production expSep is exited.
func (s *BaseZParserListener) ExitExpSep(ctx *ExpSepContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseZParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseZParserListener) ExitExpressionList(ctx *ExpressionListContext) {}
