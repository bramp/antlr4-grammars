// Generated from ZParser.g4 by ANTLR 4.7.

package z // ZParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ZParser.
type ZParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ZParser#specification.
	VisitSpecification(ctx *SpecificationContext) interface{}

	// Visit a parse tree produced by ZParser#InheritingSection.
	VisitInheritingSection(ctx *InheritingSectionContext) interface{}

	// Visit a parse tree produced by ZParser#BaseSection.
	VisitBaseSection(ctx *BaseSectionContext) interface{}

	// Visit a parse tree produced by ZParser#GivenTypesParagraph.
	VisitGivenTypesParagraph(ctx *GivenTypesParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#AxiomaticDescriptionParagraph.
	VisitAxiomaticDescriptionParagraph(ctx *AxiomaticDescriptionParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaDefinitionParagraph.
	VisitSchemaDefinitionParagraph(ctx *SchemaDefinitionParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#GenericAxiomaticDescriptionParagraph.
	VisitGenericAxiomaticDescriptionParagraph(ctx *GenericAxiomaticDescriptionParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#GenericSchemaDefinitionParagraph.
	VisitGenericSchemaDefinitionParagraph(ctx *GenericSchemaDefinitionParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#HorizontalDefinitionParagraph.
	VisitHorizontalDefinitionParagraph(ctx *HorizontalDefinitionParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#GenericHorizontalDefinitionParagraph.
	VisitGenericHorizontalDefinitionParagraph(ctx *GenericHorizontalDefinitionParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#GenericOperatorDefinitionParagraph.
	VisitGenericOperatorDefinitionParagraph(ctx *GenericOperatorDefinitionParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#FreeTypesParagraph.
	VisitFreeTypesParagraph(ctx *FreeTypesParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#ConjectureParagraph.
	VisitConjectureParagraph(ctx *ConjectureParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#GenericConjectureParagraph.
	VisitGenericConjectureParagraph(ctx *GenericConjectureParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#OperatorTemplateParagraph.
	VisitOperatorTemplateParagraph(ctx *OperatorTemplateParagraphContext) interface{}

	// Visit a parse tree produced by ZParser#freetype.
	VisitFreetype(ctx *FreetypeContext) interface{}

	// Visit a parse tree produced by ZParser#branch.
	VisitBranch(ctx *BranchContext) interface{}

	// Visit a parse tree produced by ZParser#formals.
	VisitFormals(ctx *FormalsContext) interface{}

	// Visit a parse tree produced by ZParser#ExistentialQuantificationPredicate.
	VisitExistentialQuantificationPredicate(ctx *ExistentialQuantificationPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#ConjunctionPredicate.
	VisitConjunctionPredicate(ctx *ConjunctionPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#EquivalencePredicate.
	VisitEquivalencePredicate(ctx *EquivalencePredicateContext) interface{}

	// Visit a parse tree produced by ZParser#NewlineConjunctionPredicate.
	VisitNewlineConjunctionPredicate(ctx *NewlineConjunctionPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#ImplicationPredicate.
	VisitImplicationPredicate(ctx *ImplicationPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#RelationOperatorApplicationPredicate.
	VisitRelationOperatorApplicationPredicate(ctx *RelationOperatorApplicationPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#UniversalQuantificationPredicate.
	VisitUniversalQuantificationPredicate(ctx *UniversalQuantificationPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#TruthPredicate.
	VisitTruthPredicate(ctx *TruthPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#FalsityPredicate.
	VisitFalsityPredicate(ctx *FalsityPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#ParenthesizedPredicate.
	VisitParenthesizedPredicate(ctx *ParenthesizedPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#DisjunctionPredicate.
	VisitDisjunctionPredicate(ctx *DisjunctionPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#SemicolonConjunctionPredicate.
	VisitSemicolonConjunctionPredicate(ctx *SemicolonConjunctionPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaPredicatePredicate.
	VisitSchemaPredicatePredicate(ctx *SchemaPredicatePredicateContext) interface{}

	// Visit a parse tree produced by ZParser#UniqueExistentialQuantificationPredicate.
	VisitUniqueExistentialQuantificationPredicate(ctx *UniqueExistentialQuantificationPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#NegationPredicate.
	VisitNegationPredicate(ctx *NegationPredicateContext) interface{}

	// Visit a parse tree produced by ZParser#SetComprehensionExpression.
	VisitSetComprehensionExpression(ctx *SetComprehensionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaEquivalenceExpression.
	VisitSchemaEquivalenceExpression(ctx *SchemaEquivalenceExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#NofixApplicationExpression.
	VisitNofixApplicationExpression(ctx *NofixApplicationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaConstructionExpression.
	VisitSchemaConstructionExpression(ctx *SchemaConstructionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#GenericPostfixApplicationExpression.
	VisitGenericPostfixApplicationExpression(ctx *GenericPostfixApplicationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#GenericPrefixApplicationExpression.
	VisitGenericPrefixApplicationExpression(ctx *GenericPrefixApplicationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#CharacteristicDefiniteDescriptionExpression.
	VisitCharacteristicDefiniteDescriptionExpression(ctx *CharacteristicDefiniteDescriptionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaUniversalQuantificationExpression.
	VisitSchemaUniversalQuantificationExpression(ctx *SchemaUniversalQuantificationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SetExtensionExpression.
	VisitSetExtensionExpression(ctx *SetExtensionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#FunctionConstructionExpression.
	VisitFunctionConstructionExpression(ctx *FunctionConstructionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaNegationExpression.
	VisitSchemaNegationExpression(ctx *SchemaNegationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#TupleExtensionExpression.
	VisitTupleExtensionExpression(ctx *TupleExtensionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaCompositionExpression.
	VisitSchemaCompositionExpression(ctx *SchemaCompositionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#PowersetExpression.
	VisitPowersetExpression(ctx *PowersetExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaDisjunctionExpression.
	VisitSchemaDisjunctionExpression(ctx *SchemaDisjunctionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SubstitutionExpressionExpression.
	VisitSubstitutionExpressionExpression(ctx *SubstitutionExpressionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaHidingExpression.
	VisitSchemaHidingExpression(ctx *SchemaHidingExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#PostfixApplicationExpression.
	VisitPostfixApplicationExpression(ctx *PostfixApplicationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#CartesianProductExpression.
	VisitCartesianProductExpression(ctx *CartesianProductExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#BindingSelectionExpression.
	VisitBindingSelectionExpression(ctx *BindingSelectionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#GenericInstantiationExpression.
	VisitGenericInstantiationExpression(ctx *GenericInstantiationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaExistentialQuantificationExpression.
	VisitSchemaExistentialQuantificationExpression(ctx *SchemaExistentialQuantificationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#NumberLiteralExpression.
	VisitNumberLiteralExpression(ctx *NumberLiteralExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#ParenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#InfixLeftApplicationExpression.
	VisitInfixLeftApplicationExpression(ctx *InfixLeftApplicationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#ApplicationExpression.
	VisitApplicationExpression(ctx *ApplicationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaRenamingExpression.
	VisitSchemaRenamingExpression(ctx *SchemaRenamingExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaConjunctionExpression.
	VisitSchemaConjunctionExpression(ctx *SchemaConjunctionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#ConditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaProjectionExpression.
	VisitSchemaProjectionExpression(ctx *SchemaProjectionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#InfixRightApplicationExpression.
	VisitInfixRightApplicationExpression(ctx *InfixRightApplicationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#CharacteristicSetComprehensionExpression.
	VisitCharacteristicSetComprehensionExpression(ctx *CharacteristicSetComprehensionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaPreconditionExpression.
	VisitSchemaPreconditionExpression(ctx *SchemaPreconditionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#PrefixApplicationExpression.
	VisitPrefixApplicationExpression(ctx *PrefixApplicationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#BindingConstructionExpression.
	VisitBindingConstructionExpression(ctx *BindingConstructionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaPipingExpression.
	VisitSchemaPipingExpression(ctx *SchemaPipingExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaImplicationExpression.
	VisitSchemaImplicationExpression(ctx *SchemaImplicationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#BindingExtensionExpression.
	VisitBindingExtensionExpression(ctx *BindingExtensionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaDecorationExpression.
	VisitSchemaDecorationExpression(ctx *SchemaDecorationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#SchemaUniqueExistentialQuantificationExpression.
	VisitSchemaUniqueExistentialQuantificationExpression(ctx *SchemaUniqueExistentialQuantificationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#DefiniteDescriptionExpression.
	VisitDefiniteDescriptionExpression(ctx *DefiniteDescriptionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#ReferenceExpression.
	VisitReferenceExpression(ctx *ReferenceExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#GenericInfixApplicationExpression.
	VisitGenericInfixApplicationExpression(ctx *GenericInfixApplicationExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#TupleSelectionExpression.
	VisitTupleSelectionExpression(ctx *TupleSelectionExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#schemaText.
	VisitSchemaText(ctx *SchemaTextContext) interface{}

	// Visit a parse tree produced by ZParser#declPart.
	VisitDeclPart(ctx *DeclPartContext) interface{}

	// Visit a parse tree produced by ZParser#declNameExpression.
	VisitDeclNameExpression(ctx *DeclNameExpressionContext) interface{}

	// Visit a parse tree produced by ZParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ZParser#operatorTemplate.
	VisitOperatorTemplate(ctx *OperatorTemplateContext) interface{}

	// Visit a parse tree produced by ZParser#categoryTemplate.
	VisitCategoryTemplate(ctx *CategoryTemplateContext) interface{}

	// Visit a parse tree produced by ZParser#prec.
	VisitPrec(ctx *PrecContext) interface{}

	// Visit a parse tree produced by ZParser#assoc.
	VisitAssoc(ctx *AssocContext) interface{}

	// Visit a parse tree produced by ZParser#template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by ZParser#prefixTemplate.
	VisitPrefixTemplate(ctx *PrefixTemplateContext) interface{}

	// Visit a parse tree produced by ZParser#postfixTemplate.
	VisitPostfixTemplate(ctx *PostfixTemplateContext) interface{}

	// Visit a parse tree produced by ZParser#infixTemplate.
	VisitInfixTemplate(ctx *InfixTemplateContext) interface{}

	// Visit a parse tree produced by ZParser#nofixTemplate.
	VisitNofixTemplate(ctx *NofixTemplateContext) interface{}

	// Visit a parse tree produced by ZParser#declName.
	VisitDeclName(ctx *DeclNameContext) interface{}

	// Visit a parse tree produced by ZParser#refName.
	VisitRefName(ctx *RefNameContext) interface{}

	// Visit a parse tree produced by ZParser#opName.
	VisitOpName(ctx *OpNameContext) interface{}

	// Visit a parse tree produced by ZParser#prefixName.
	VisitPrefixName(ctx *PrefixNameContext) interface{}

	// Visit a parse tree produced by ZParser#postfixName.
	VisitPostfixName(ctx *PostfixNameContext) interface{}

	// Visit a parse tree produced by ZParser#infixName.
	VisitInfixName(ctx *InfixNameContext) interface{}

	// Visit a parse tree produced by ZParser#nofixName.
	VisitNofixName(ctx *NofixNameContext) interface{}

	// Visit a parse tree produced by ZParser#genName.
	VisitGenName(ctx *GenNameContext) interface{}

	// Visit a parse tree produced by ZParser#prefixGenName.
	VisitPrefixGenName(ctx *PrefixGenNameContext) interface{}

	// Visit a parse tree produced by ZParser#postfixGenName.
	VisitPostfixGenName(ctx *PostfixGenNameContext) interface{}

	// Visit a parse tree produced by ZParser#infixGenName.
	VisitInfixGenName(ctx *InfixGenNameContext) interface{}

	// Visit a parse tree produced by ZParser#nofixGenName.
	VisitNofixGenName(ctx *NofixGenNameContext) interface{}

	// Visit a parse tree produced by ZParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by ZParser#prefixRel.
	VisitPrefixRel(ctx *PrefixRelContext) interface{}

	// Visit a parse tree produced by ZParser#postfixRel.
	VisitPostfixRel(ctx *PostfixRelContext) interface{}

	// Visit a parse tree produced by ZParser#infixRel.
	VisitInfixRel(ctx *InfixRelContext) interface{}

	// Visit a parse tree produced by ZParser#nofixRel.
	VisitNofixRel(ctx *NofixRelContext) interface{}

	// Visit a parse tree produced by ZParser#application.
	VisitApplication(ctx *ApplicationContext) interface{}

	// Visit a parse tree produced by ZParser#prefixApp.
	VisitPrefixApp(ctx *PrefixAppContext) interface{}

	// Visit a parse tree produced by ZParser#postfixApp.
	VisitPostfixApp(ctx *PostfixAppContext) interface{}

	// Visit a parse tree produced by ZParser#infixApp.
	VisitInfixApp(ctx *InfixAppContext) interface{}

	// Visit a parse tree produced by ZParser#nofixApp.
	VisitNofixApp(ctx *NofixAppContext) interface{}

	// Visit a parse tree produced by ZParser#expSep.
	VisitExpSep(ctx *ExpSepContext) interface{}

	// Visit a parse tree produced by ZParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}
}
