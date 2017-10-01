// Generated from ZParser.g4 by ANTLR 4.7.

package z // ZParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ZParserListener is a complete listener for a parse tree produced by ZParser.
type ZParserListener interface {
	antlr.ParseTreeListener

	// EnterSpecification is called when entering the specification production.
	EnterSpecification(c *SpecificationContext)

	// EnterInheritingSection is called when entering the InheritingSection production.
	EnterInheritingSection(c *InheritingSectionContext)

	// EnterBaseSection is called when entering the BaseSection production.
	EnterBaseSection(c *BaseSectionContext)

	// EnterGivenTypesParagraph is called when entering the GivenTypesParagraph production.
	EnterGivenTypesParagraph(c *GivenTypesParagraphContext)

	// EnterAxiomaticDescriptionParagraph is called when entering the AxiomaticDescriptionParagraph production.
	EnterAxiomaticDescriptionParagraph(c *AxiomaticDescriptionParagraphContext)

	// EnterSchemaDefinitionParagraph is called when entering the SchemaDefinitionParagraph production.
	EnterSchemaDefinitionParagraph(c *SchemaDefinitionParagraphContext)

	// EnterGenericAxiomaticDescriptionParagraph is called when entering the GenericAxiomaticDescriptionParagraph production.
	EnterGenericAxiomaticDescriptionParagraph(c *GenericAxiomaticDescriptionParagraphContext)

	// EnterGenericSchemaDefinitionParagraph is called when entering the GenericSchemaDefinitionParagraph production.
	EnterGenericSchemaDefinitionParagraph(c *GenericSchemaDefinitionParagraphContext)

	// EnterHorizontalDefinitionParagraph is called when entering the HorizontalDefinitionParagraph production.
	EnterHorizontalDefinitionParagraph(c *HorizontalDefinitionParagraphContext)

	// EnterGenericHorizontalDefinitionParagraph is called when entering the GenericHorizontalDefinitionParagraph production.
	EnterGenericHorizontalDefinitionParagraph(c *GenericHorizontalDefinitionParagraphContext)

	// EnterGenericOperatorDefinitionParagraph is called when entering the GenericOperatorDefinitionParagraph production.
	EnterGenericOperatorDefinitionParagraph(c *GenericOperatorDefinitionParagraphContext)

	// EnterFreeTypesParagraph is called when entering the FreeTypesParagraph production.
	EnterFreeTypesParagraph(c *FreeTypesParagraphContext)

	// EnterConjectureParagraph is called when entering the ConjectureParagraph production.
	EnterConjectureParagraph(c *ConjectureParagraphContext)

	// EnterGenericConjectureParagraph is called when entering the GenericConjectureParagraph production.
	EnterGenericConjectureParagraph(c *GenericConjectureParagraphContext)

	// EnterOperatorTemplateParagraph is called when entering the OperatorTemplateParagraph production.
	EnterOperatorTemplateParagraph(c *OperatorTemplateParagraphContext)

	// EnterFreetype is called when entering the freetype production.
	EnterFreetype(c *FreetypeContext)

	// EnterBranch is called when entering the branch production.
	EnterBranch(c *BranchContext)

	// EnterFormals is called when entering the formals production.
	EnterFormals(c *FormalsContext)

	// EnterExistentialQuantificationPredicate is called when entering the ExistentialQuantificationPredicate production.
	EnterExistentialQuantificationPredicate(c *ExistentialQuantificationPredicateContext)

	// EnterConjunctionPredicate is called when entering the ConjunctionPredicate production.
	EnterConjunctionPredicate(c *ConjunctionPredicateContext)

	// EnterEquivalencePredicate is called when entering the EquivalencePredicate production.
	EnterEquivalencePredicate(c *EquivalencePredicateContext)

	// EnterNewlineConjunctionPredicate is called when entering the NewlineConjunctionPredicate production.
	EnterNewlineConjunctionPredicate(c *NewlineConjunctionPredicateContext)

	// EnterImplicationPredicate is called when entering the ImplicationPredicate production.
	EnterImplicationPredicate(c *ImplicationPredicateContext)

	// EnterRelationOperatorApplicationPredicate is called when entering the RelationOperatorApplicationPredicate production.
	EnterRelationOperatorApplicationPredicate(c *RelationOperatorApplicationPredicateContext)

	// EnterUniversalQuantificationPredicate is called when entering the UniversalQuantificationPredicate production.
	EnterUniversalQuantificationPredicate(c *UniversalQuantificationPredicateContext)

	// EnterTruthPredicate is called when entering the TruthPredicate production.
	EnterTruthPredicate(c *TruthPredicateContext)

	// EnterFalsityPredicate is called when entering the FalsityPredicate production.
	EnterFalsityPredicate(c *FalsityPredicateContext)

	// EnterParenthesizedPredicate is called when entering the ParenthesizedPredicate production.
	EnterParenthesizedPredicate(c *ParenthesizedPredicateContext)

	// EnterDisjunctionPredicate is called when entering the DisjunctionPredicate production.
	EnterDisjunctionPredicate(c *DisjunctionPredicateContext)

	// EnterSemicolonConjunctionPredicate is called when entering the SemicolonConjunctionPredicate production.
	EnterSemicolonConjunctionPredicate(c *SemicolonConjunctionPredicateContext)

	// EnterSchemaPredicatePredicate is called when entering the SchemaPredicatePredicate production.
	EnterSchemaPredicatePredicate(c *SchemaPredicatePredicateContext)

	// EnterUniqueExistentialQuantificationPredicate is called when entering the UniqueExistentialQuantificationPredicate production.
	EnterUniqueExistentialQuantificationPredicate(c *UniqueExistentialQuantificationPredicateContext)

	// EnterNegationPredicate is called when entering the NegationPredicate production.
	EnterNegationPredicate(c *NegationPredicateContext)

	// EnterSetComprehensionExpression is called when entering the SetComprehensionExpression production.
	EnterSetComprehensionExpression(c *SetComprehensionExpressionContext)

	// EnterSchemaEquivalenceExpression is called when entering the SchemaEquivalenceExpression production.
	EnterSchemaEquivalenceExpression(c *SchemaEquivalenceExpressionContext)

	// EnterNofixApplicationExpression is called when entering the NofixApplicationExpression production.
	EnterNofixApplicationExpression(c *NofixApplicationExpressionContext)

	// EnterSchemaConstructionExpression is called when entering the SchemaConstructionExpression production.
	EnterSchemaConstructionExpression(c *SchemaConstructionExpressionContext)

	// EnterGenericPostfixApplicationExpression is called when entering the GenericPostfixApplicationExpression production.
	EnterGenericPostfixApplicationExpression(c *GenericPostfixApplicationExpressionContext)

	// EnterGenericPrefixApplicationExpression is called when entering the GenericPrefixApplicationExpression production.
	EnterGenericPrefixApplicationExpression(c *GenericPrefixApplicationExpressionContext)

	// EnterCharacteristicDefiniteDescriptionExpression is called when entering the CharacteristicDefiniteDescriptionExpression production.
	EnterCharacteristicDefiniteDescriptionExpression(c *CharacteristicDefiniteDescriptionExpressionContext)

	// EnterSchemaUniversalQuantificationExpression is called when entering the SchemaUniversalQuantificationExpression production.
	EnterSchemaUniversalQuantificationExpression(c *SchemaUniversalQuantificationExpressionContext)

	// EnterSetExtensionExpression is called when entering the SetExtensionExpression production.
	EnterSetExtensionExpression(c *SetExtensionExpressionContext)

	// EnterFunctionConstructionExpression is called when entering the FunctionConstructionExpression production.
	EnterFunctionConstructionExpression(c *FunctionConstructionExpressionContext)

	// EnterSchemaNegationExpression is called when entering the SchemaNegationExpression production.
	EnterSchemaNegationExpression(c *SchemaNegationExpressionContext)

	// EnterTupleExtensionExpression is called when entering the TupleExtensionExpression production.
	EnterTupleExtensionExpression(c *TupleExtensionExpressionContext)

	// EnterSchemaCompositionExpression is called when entering the SchemaCompositionExpression production.
	EnterSchemaCompositionExpression(c *SchemaCompositionExpressionContext)

	// EnterPowersetExpression is called when entering the PowersetExpression production.
	EnterPowersetExpression(c *PowersetExpressionContext)

	// EnterSchemaDisjunctionExpression is called when entering the SchemaDisjunctionExpression production.
	EnterSchemaDisjunctionExpression(c *SchemaDisjunctionExpressionContext)

	// EnterSubstitutionExpressionExpression is called when entering the SubstitutionExpressionExpression production.
	EnterSubstitutionExpressionExpression(c *SubstitutionExpressionExpressionContext)

	// EnterSchemaHidingExpression is called when entering the SchemaHidingExpression production.
	EnterSchemaHidingExpression(c *SchemaHidingExpressionContext)

	// EnterPostfixApplicationExpression is called when entering the PostfixApplicationExpression production.
	EnterPostfixApplicationExpression(c *PostfixApplicationExpressionContext)

	// EnterCartesianProductExpression is called when entering the CartesianProductExpression production.
	EnterCartesianProductExpression(c *CartesianProductExpressionContext)

	// EnterBindingSelectionExpression is called when entering the BindingSelectionExpression production.
	EnterBindingSelectionExpression(c *BindingSelectionExpressionContext)

	// EnterGenericInstantiationExpression is called when entering the GenericInstantiationExpression production.
	EnterGenericInstantiationExpression(c *GenericInstantiationExpressionContext)

	// EnterSchemaExistentialQuantificationExpression is called when entering the SchemaExistentialQuantificationExpression production.
	EnterSchemaExistentialQuantificationExpression(c *SchemaExistentialQuantificationExpressionContext)

	// EnterNumberLiteralExpression is called when entering the NumberLiteralExpression production.
	EnterNumberLiteralExpression(c *NumberLiteralExpressionContext)

	// EnterParenthesizedExpression is called when entering the ParenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterInfixLeftApplicationExpression is called when entering the InfixLeftApplicationExpression production.
	EnterInfixLeftApplicationExpression(c *InfixLeftApplicationExpressionContext)

	// EnterApplicationExpression is called when entering the ApplicationExpression production.
	EnterApplicationExpression(c *ApplicationExpressionContext)

	// EnterSchemaRenamingExpression is called when entering the SchemaRenamingExpression production.
	EnterSchemaRenamingExpression(c *SchemaRenamingExpressionContext)

	// EnterSchemaConjunctionExpression is called when entering the SchemaConjunctionExpression production.
	EnterSchemaConjunctionExpression(c *SchemaConjunctionExpressionContext)

	// EnterConditionalExpression is called when entering the ConditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterSchemaProjectionExpression is called when entering the SchemaProjectionExpression production.
	EnterSchemaProjectionExpression(c *SchemaProjectionExpressionContext)

	// EnterInfixRightApplicationExpression is called when entering the InfixRightApplicationExpression production.
	EnterInfixRightApplicationExpression(c *InfixRightApplicationExpressionContext)

	// EnterCharacteristicSetComprehensionExpression is called when entering the CharacteristicSetComprehensionExpression production.
	EnterCharacteristicSetComprehensionExpression(c *CharacteristicSetComprehensionExpressionContext)

	// EnterSchemaPreconditionExpression is called when entering the SchemaPreconditionExpression production.
	EnterSchemaPreconditionExpression(c *SchemaPreconditionExpressionContext)

	// EnterPrefixApplicationExpression is called when entering the PrefixApplicationExpression production.
	EnterPrefixApplicationExpression(c *PrefixApplicationExpressionContext)

	// EnterBindingConstructionExpression is called when entering the BindingConstructionExpression production.
	EnterBindingConstructionExpression(c *BindingConstructionExpressionContext)

	// EnterSchemaPipingExpression is called when entering the SchemaPipingExpression production.
	EnterSchemaPipingExpression(c *SchemaPipingExpressionContext)

	// EnterSchemaImplicationExpression is called when entering the SchemaImplicationExpression production.
	EnterSchemaImplicationExpression(c *SchemaImplicationExpressionContext)

	// EnterBindingExtensionExpression is called when entering the BindingExtensionExpression production.
	EnterBindingExtensionExpression(c *BindingExtensionExpressionContext)

	// EnterSchemaDecorationExpression is called when entering the SchemaDecorationExpression production.
	EnterSchemaDecorationExpression(c *SchemaDecorationExpressionContext)

	// EnterSchemaUniqueExistentialQuantificationExpression is called when entering the SchemaUniqueExistentialQuantificationExpression production.
	EnterSchemaUniqueExistentialQuantificationExpression(c *SchemaUniqueExistentialQuantificationExpressionContext)

	// EnterDefiniteDescriptionExpression is called when entering the DefiniteDescriptionExpression production.
	EnterDefiniteDescriptionExpression(c *DefiniteDescriptionExpressionContext)

	// EnterReferenceExpression is called when entering the ReferenceExpression production.
	EnterReferenceExpression(c *ReferenceExpressionContext)

	// EnterGenericInfixApplicationExpression is called when entering the GenericInfixApplicationExpression production.
	EnterGenericInfixApplicationExpression(c *GenericInfixApplicationExpressionContext)

	// EnterTupleSelectionExpression is called when entering the TupleSelectionExpression production.
	EnterTupleSelectionExpression(c *TupleSelectionExpressionContext)

	// EnterSchemaText is called when entering the schemaText production.
	EnterSchemaText(c *SchemaTextContext)

	// EnterDeclPart is called when entering the declPart production.
	EnterDeclPart(c *DeclPartContext)

	// EnterDeclNameExpression is called when entering the declNameExpression production.
	EnterDeclNameExpression(c *DeclNameExpressionContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterOperatorTemplate is called when entering the operatorTemplate production.
	EnterOperatorTemplate(c *OperatorTemplateContext)

	// EnterCategoryTemplate is called when entering the categoryTemplate production.
	EnterCategoryTemplate(c *CategoryTemplateContext)

	// EnterPrec is called when entering the prec production.
	EnterPrec(c *PrecContext)

	// EnterAssoc is called when entering the assoc production.
	EnterAssoc(c *AssocContext)

	// EnterTemplate is called when entering the template production.
	EnterTemplate(c *TemplateContext)

	// EnterPrefixTemplate is called when entering the prefixTemplate production.
	EnterPrefixTemplate(c *PrefixTemplateContext)

	// EnterPostfixTemplate is called when entering the postfixTemplate production.
	EnterPostfixTemplate(c *PostfixTemplateContext)

	// EnterInfixTemplate is called when entering the infixTemplate production.
	EnterInfixTemplate(c *InfixTemplateContext)

	// EnterNofixTemplate is called when entering the nofixTemplate production.
	EnterNofixTemplate(c *NofixTemplateContext)

	// EnterDeclName is called when entering the declName production.
	EnterDeclName(c *DeclNameContext)

	// EnterRefName is called when entering the refName production.
	EnterRefName(c *RefNameContext)

	// EnterOpName is called when entering the opName production.
	EnterOpName(c *OpNameContext)

	// EnterPrefixName is called when entering the prefixName production.
	EnterPrefixName(c *PrefixNameContext)

	// EnterPostfixName is called when entering the postfixName production.
	EnterPostfixName(c *PostfixNameContext)

	// EnterInfixName is called when entering the infixName production.
	EnterInfixName(c *InfixNameContext)

	// EnterNofixName is called when entering the nofixName production.
	EnterNofixName(c *NofixNameContext)

	// EnterGenName is called when entering the genName production.
	EnterGenName(c *GenNameContext)

	// EnterPrefixGenName is called when entering the prefixGenName production.
	EnterPrefixGenName(c *PrefixGenNameContext)

	// EnterPostfixGenName is called when entering the postfixGenName production.
	EnterPostfixGenName(c *PostfixGenNameContext)

	// EnterInfixGenName is called when entering the infixGenName production.
	EnterInfixGenName(c *InfixGenNameContext)

	// EnterNofixGenName is called when entering the nofixGenName production.
	EnterNofixGenName(c *NofixGenNameContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterPrefixRel is called when entering the prefixRel production.
	EnterPrefixRel(c *PrefixRelContext)

	// EnterPostfixRel is called when entering the postfixRel production.
	EnterPostfixRel(c *PostfixRelContext)

	// EnterInfixRel is called when entering the infixRel production.
	EnterInfixRel(c *InfixRelContext)

	// EnterNofixRel is called when entering the nofixRel production.
	EnterNofixRel(c *NofixRelContext)

	// EnterApplication is called when entering the application production.
	EnterApplication(c *ApplicationContext)

	// EnterPrefixApp is called when entering the prefixApp production.
	EnterPrefixApp(c *PrefixAppContext)

	// EnterPostfixApp is called when entering the postfixApp production.
	EnterPostfixApp(c *PostfixAppContext)

	// EnterInfixApp is called when entering the infixApp production.
	EnterInfixApp(c *InfixAppContext)

	// EnterNofixApp is called when entering the nofixApp production.
	EnterNofixApp(c *NofixAppContext)

	// EnterExpSep is called when entering the expSep production.
	EnterExpSep(c *ExpSepContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// ExitSpecification is called when exiting the specification production.
	ExitSpecification(c *SpecificationContext)

	// ExitInheritingSection is called when exiting the InheritingSection production.
	ExitInheritingSection(c *InheritingSectionContext)

	// ExitBaseSection is called when exiting the BaseSection production.
	ExitBaseSection(c *BaseSectionContext)

	// ExitGivenTypesParagraph is called when exiting the GivenTypesParagraph production.
	ExitGivenTypesParagraph(c *GivenTypesParagraphContext)

	// ExitAxiomaticDescriptionParagraph is called when exiting the AxiomaticDescriptionParagraph production.
	ExitAxiomaticDescriptionParagraph(c *AxiomaticDescriptionParagraphContext)

	// ExitSchemaDefinitionParagraph is called when exiting the SchemaDefinitionParagraph production.
	ExitSchemaDefinitionParagraph(c *SchemaDefinitionParagraphContext)

	// ExitGenericAxiomaticDescriptionParagraph is called when exiting the GenericAxiomaticDescriptionParagraph production.
	ExitGenericAxiomaticDescriptionParagraph(c *GenericAxiomaticDescriptionParagraphContext)

	// ExitGenericSchemaDefinitionParagraph is called when exiting the GenericSchemaDefinitionParagraph production.
	ExitGenericSchemaDefinitionParagraph(c *GenericSchemaDefinitionParagraphContext)

	// ExitHorizontalDefinitionParagraph is called when exiting the HorizontalDefinitionParagraph production.
	ExitHorizontalDefinitionParagraph(c *HorizontalDefinitionParagraphContext)

	// ExitGenericHorizontalDefinitionParagraph is called when exiting the GenericHorizontalDefinitionParagraph production.
	ExitGenericHorizontalDefinitionParagraph(c *GenericHorizontalDefinitionParagraphContext)

	// ExitGenericOperatorDefinitionParagraph is called when exiting the GenericOperatorDefinitionParagraph production.
	ExitGenericOperatorDefinitionParagraph(c *GenericOperatorDefinitionParagraphContext)

	// ExitFreeTypesParagraph is called when exiting the FreeTypesParagraph production.
	ExitFreeTypesParagraph(c *FreeTypesParagraphContext)

	// ExitConjectureParagraph is called when exiting the ConjectureParagraph production.
	ExitConjectureParagraph(c *ConjectureParagraphContext)

	// ExitGenericConjectureParagraph is called when exiting the GenericConjectureParagraph production.
	ExitGenericConjectureParagraph(c *GenericConjectureParagraphContext)

	// ExitOperatorTemplateParagraph is called when exiting the OperatorTemplateParagraph production.
	ExitOperatorTemplateParagraph(c *OperatorTemplateParagraphContext)

	// ExitFreetype is called when exiting the freetype production.
	ExitFreetype(c *FreetypeContext)

	// ExitBranch is called when exiting the branch production.
	ExitBranch(c *BranchContext)

	// ExitFormals is called when exiting the formals production.
	ExitFormals(c *FormalsContext)

	// ExitExistentialQuantificationPredicate is called when exiting the ExistentialQuantificationPredicate production.
	ExitExistentialQuantificationPredicate(c *ExistentialQuantificationPredicateContext)

	// ExitConjunctionPredicate is called when exiting the ConjunctionPredicate production.
	ExitConjunctionPredicate(c *ConjunctionPredicateContext)

	// ExitEquivalencePredicate is called when exiting the EquivalencePredicate production.
	ExitEquivalencePredicate(c *EquivalencePredicateContext)

	// ExitNewlineConjunctionPredicate is called when exiting the NewlineConjunctionPredicate production.
	ExitNewlineConjunctionPredicate(c *NewlineConjunctionPredicateContext)

	// ExitImplicationPredicate is called when exiting the ImplicationPredicate production.
	ExitImplicationPredicate(c *ImplicationPredicateContext)

	// ExitRelationOperatorApplicationPredicate is called when exiting the RelationOperatorApplicationPredicate production.
	ExitRelationOperatorApplicationPredicate(c *RelationOperatorApplicationPredicateContext)

	// ExitUniversalQuantificationPredicate is called when exiting the UniversalQuantificationPredicate production.
	ExitUniversalQuantificationPredicate(c *UniversalQuantificationPredicateContext)

	// ExitTruthPredicate is called when exiting the TruthPredicate production.
	ExitTruthPredicate(c *TruthPredicateContext)

	// ExitFalsityPredicate is called when exiting the FalsityPredicate production.
	ExitFalsityPredicate(c *FalsityPredicateContext)

	// ExitParenthesizedPredicate is called when exiting the ParenthesizedPredicate production.
	ExitParenthesizedPredicate(c *ParenthesizedPredicateContext)

	// ExitDisjunctionPredicate is called when exiting the DisjunctionPredicate production.
	ExitDisjunctionPredicate(c *DisjunctionPredicateContext)

	// ExitSemicolonConjunctionPredicate is called when exiting the SemicolonConjunctionPredicate production.
	ExitSemicolonConjunctionPredicate(c *SemicolonConjunctionPredicateContext)

	// ExitSchemaPredicatePredicate is called when exiting the SchemaPredicatePredicate production.
	ExitSchemaPredicatePredicate(c *SchemaPredicatePredicateContext)

	// ExitUniqueExistentialQuantificationPredicate is called when exiting the UniqueExistentialQuantificationPredicate production.
	ExitUniqueExistentialQuantificationPredicate(c *UniqueExistentialQuantificationPredicateContext)

	// ExitNegationPredicate is called when exiting the NegationPredicate production.
	ExitNegationPredicate(c *NegationPredicateContext)

	// ExitSetComprehensionExpression is called when exiting the SetComprehensionExpression production.
	ExitSetComprehensionExpression(c *SetComprehensionExpressionContext)

	// ExitSchemaEquivalenceExpression is called when exiting the SchemaEquivalenceExpression production.
	ExitSchemaEquivalenceExpression(c *SchemaEquivalenceExpressionContext)

	// ExitNofixApplicationExpression is called when exiting the NofixApplicationExpression production.
	ExitNofixApplicationExpression(c *NofixApplicationExpressionContext)

	// ExitSchemaConstructionExpression is called when exiting the SchemaConstructionExpression production.
	ExitSchemaConstructionExpression(c *SchemaConstructionExpressionContext)

	// ExitGenericPostfixApplicationExpression is called when exiting the GenericPostfixApplicationExpression production.
	ExitGenericPostfixApplicationExpression(c *GenericPostfixApplicationExpressionContext)

	// ExitGenericPrefixApplicationExpression is called when exiting the GenericPrefixApplicationExpression production.
	ExitGenericPrefixApplicationExpression(c *GenericPrefixApplicationExpressionContext)

	// ExitCharacteristicDefiniteDescriptionExpression is called when exiting the CharacteristicDefiniteDescriptionExpression production.
	ExitCharacteristicDefiniteDescriptionExpression(c *CharacteristicDefiniteDescriptionExpressionContext)

	// ExitSchemaUniversalQuantificationExpression is called when exiting the SchemaUniversalQuantificationExpression production.
	ExitSchemaUniversalQuantificationExpression(c *SchemaUniversalQuantificationExpressionContext)

	// ExitSetExtensionExpression is called when exiting the SetExtensionExpression production.
	ExitSetExtensionExpression(c *SetExtensionExpressionContext)

	// ExitFunctionConstructionExpression is called when exiting the FunctionConstructionExpression production.
	ExitFunctionConstructionExpression(c *FunctionConstructionExpressionContext)

	// ExitSchemaNegationExpression is called when exiting the SchemaNegationExpression production.
	ExitSchemaNegationExpression(c *SchemaNegationExpressionContext)

	// ExitTupleExtensionExpression is called when exiting the TupleExtensionExpression production.
	ExitTupleExtensionExpression(c *TupleExtensionExpressionContext)

	// ExitSchemaCompositionExpression is called when exiting the SchemaCompositionExpression production.
	ExitSchemaCompositionExpression(c *SchemaCompositionExpressionContext)

	// ExitPowersetExpression is called when exiting the PowersetExpression production.
	ExitPowersetExpression(c *PowersetExpressionContext)

	// ExitSchemaDisjunctionExpression is called when exiting the SchemaDisjunctionExpression production.
	ExitSchemaDisjunctionExpression(c *SchemaDisjunctionExpressionContext)

	// ExitSubstitutionExpressionExpression is called when exiting the SubstitutionExpressionExpression production.
	ExitSubstitutionExpressionExpression(c *SubstitutionExpressionExpressionContext)

	// ExitSchemaHidingExpression is called when exiting the SchemaHidingExpression production.
	ExitSchemaHidingExpression(c *SchemaHidingExpressionContext)

	// ExitPostfixApplicationExpression is called when exiting the PostfixApplicationExpression production.
	ExitPostfixApplicationExpression(c *PostfixApplicationExpressionContext)

	// ExitCartesianProductExpression is called when exiting the CartesianProductExpression production.
	ExitCartesianProductExpression(c *CartesianProductExpressionContext)

	// ExitBindingSelectionExpression is called when exiting the BindingSelectionExpression production.
	ExitBindingSelectionExpression(c *BindingSelectionExpressionContext)

	// ExitGenericInstantiationExpression is called when exiting the GenericInstantiationExpression production.
	ExitGenericInstantiationExpression(c *GenericInstantiationExpressionContext)

	// ExitSchemaExistentialQuantificationExpression is called when exiting the SchemaExistentialQuantificationExpression production.
	ExitSchemaExistentialQuantificationExpression(c *SchemaExistentialQuantificationExpressionContext)

	// ExitNumberLiteralExpression is called when exiting the NumberLiteralExpression production.
	ExitNumberLiteralExpression(c *NumberLiteralExpressionContext)

	// ExitParenthesizedExpression is called when exiting the ParenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitInfixLeftApplicationExpression is called when exiting the InfixLeftApplicationExpression production.
	ExitInfixLeftApplicationExpression(c *InfixLeftApplicationExpressionContext)

	// ExitApplicationExpression is called when exiting the ApplicationExpression production.
	ExitApplicationExpression(c *ApplicationExpressionContext)

	// ExitSchemaRenamingExpression is called when exiting the SchemaRenamingExpression production.
	ExitSchemaRenamingExpression(c *SchemaRenamingExpressionContext)

	// ExitSchemaConjunctionExpression is called when exiting the SchemaConjunctionExpression production.
	ExitSchemaConjunctionExpression(c *SchemaConjunctionExpressionContext)

	// ExitConditionalExpression is called when exiting the ConditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitSchemaProjectionExpression is called when exiting the SchemaProjectionExpression production.
	ExitSchemaProjectionExpression(c *SchemaProjectionExpressionContext)

	// ExitInfixRightApplicationExpression is called when exiting the InfixRightApplicationExpression production.
	ExitInfixRightApplicationExpression(c *InfixRightApplicationExpressionContext)

	// ExitCharacteristicSetComprehensionExpression is called when exiting the CharacteristicSetComprehensionExpression production.
	ExitCharacteristicSetComprehensionExpression(c *CharacteristicSetComprehensionExpressionContext)

	// ExitSchemaPreconditionExpression is called when exiting the SchemaPreconditionExpression production.
	ExitSchemaPreconditionExpression(c *SchemaPreconditionExpressionContext)

	// ExitPrefixApplicationExpression is called when exiting the PrefixApplicationExpression production.
	ExitPrefixApplicationExpression(c *PrefixApplicationExpressionContext)

	// ExitBindingConstructionExpression is called when exiting the BindingConstructionExpression production.
	ExitBindingConstructionExpression(c *BindingConstructionExpressionContext)

	// ExitSchemaPipingExpression is called when exiting the SchemaPipingExpression production.
	ExitSchemaPipingExpression(c *SchemaPipingExpressionContext)

	// ExitSchemaImplicationExpression is called when exiting the SchemaImplicationExpression production.
	ExitSchemaImplicationExpression(c *SchemaImplicationExpressionContext)

	// ExitBindingExtensionExpression is called when exiting the BindingExtensionExpression production.
	ExitBindingExtensionExpression(c *BindingExtensionExpressionContext)

	// ExitSchemaDecorationExpression is called when exiting the SchemaDecorationExpression production.
	ExitSchemaDecorationExpression(c *SchemaDecorationExpressionContext)

	// ExitSchemaUniqueExistentialQuantificationExpression is called when exiting the SchemaUniqueExistentialQuantificationExpression production.
	ExitSchemaUniqueExistentialQuantificationExpression(c *SchemaUniqueExistentialQuantificationExpressionContext)

	// ExitDefiniteDescriptionExpression is called when exiting the DefiniteDescriptionExpression production.
	ExitDefiniteDescriptionExpression(c *DefiniteDescriptionExpressionContext)

	// ExitReferenceExpression is called when exiting the ReferenceExpression production.
	ExitReferenceExpression(c *ReferenceExpressionContext)

	// ExitGenericInfixApplicationExpression is called when exiting the GenericInfixApplicationExpression production.
	ExitGenericInfixApplicationExpression(c *GenericInfixApplicationExpressionContext)

	// ExitTupleSelectionExpression is called when exiting the TupleSelectionExpression production.
	ExitTupleSelectionExpression(c *TupleSelectionExpressionContext)

	// ExitSchemaText is called when exiting the schemaText production.
	ExitSchemaText(c *SchemaTextContext)

	// ExitDeclPart is called when exiting the declPart production.
	ExitDeclPart(c *DeclPartContext)

	// ExitDeclNameExpression is called when exiting the declNameExpression production.
	ExitDeclNameExpression(c *DeclNameExpressionContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitOperatorTemplate is called when exiting the operatorTemplate production.
	ExitOperatorTemplate(c *OperatorTemplateContext)

	// ExitCategoryTemplate is called when exiting the categoryTemplate production.
	ExitCategoryTemplate(c *CategoryTemplateContext)

	// ExitPrec is called when exiting the prec production.
	ExitPrec(c *PrecContext)

	// ExitAssoc is called when exiting the assoc production.
	ExitAssoc(c *AssocContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitPrefixTemplate is called when exiting the prefixTemplate production.
	ExitPrefixTemplate(c *PrefixTemplateContext)

	// ExitPostfixTemplate is called when exiting the postfixTemplate production.
	ExitPostfixTemplate(c *PostfixTemplateContext)

	// ExitInfixTemplate is called when exiting the infixTemplate production.
	ExitInfixTemplate(c *InfixTemplateContext)

	// ExitNofixTemplate is called when exiting the nofixTemplate production.
	ExitNofixTemplate(c *NofixTemplateContext)

	// ExitDeclName is called when exiting the declName production.
	ExitDeclName(c *DeclNameContext)

	// ExitRefName is called when exiting the refName production.
	ExitRefName(c *RefNameContext)

	// ExitOpName is called when exiting the opName production.
	ExitOpName(c *OpNameContext)

	// ExitPrefixName is called when exiting the prefixName production.
	ExitPrefixName(c *PrefixNameContext)

	// ExitPostfixName is called when exiting the postfixName production.
	ExitPostfixName(c *PostfixNameContext)

	// ExitInfixName is called when exiting the infixName production.
	ExitInfixName(c *InfixNameContext)

	// ExitNofixName is called when exiting the nofixName production.
	ExitNofixName(c *NofixNameContext)

	// ExitGenName is called when exiting the genName production.
	ExitGenName(c *GenNameContext)

	// ExitPrefixGenName is called when exiting the prefixGenName production.
	ExitPrefixGenName(c *PrefixGenNameContext)

	// ExitPostfixGenName is called when exiting the postfixGenName production.
	ExitPostfixGenName(c *PostfixGenNameContext)

	// ExitInfixGenName is called when exiting the infixGenName production.
	ExitInfixGenName(c *InfixGenNameContext)

	// ExitNofixGenName is called when exiting the nofixGenName production.
	ExitNofixGenName(c *NofixGenNameContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitPrefixRel is called when exiting the prefixRel production.
	ExitPrefixRel(c *PrefixRelContext)

	// ExitPostfixRel is called when exiting the postfixRel production.
	ExitPostfixRel(c *PostfixRelContext)

	// ExitInfixRel is called when exiting the infixRel production.
	ExitInfixRel(c *InfixRelContext)

	// ExitNofixRel is called when exiting the nofixRel production.
	ExitNofixRel(c *NofixRelContext)

	// ExitApplication is called when exiting the application production.
	ExitApplication(c *ApplicationContext)

	// ExitPrefixApp is called when exiting the prefixApp production.
	ExitPrefixApp(c *PrefixAppContext)

	// ExitPostfixApp is called when exiting the postfixApp production.
	ExitPostfixApp(c *PostfixAppContext)

	// ExitInfixApp is called when exiting the infixApp production.
	ExitInfixApp(c *InfixAppContext)

	// ExitNofixApp is called when exiting the nofixApp production.
	ExitNofixApp(c *NofixAppContext)

	// ExitExpSep is called when exiting the expSep production.
	ExitExpSep(c *ExpSepContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)
}
