// Code generated from CtoParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cto // CtoParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CtoParserListener is a complete listener for a parse tree produced by CtoParser.
type CtoParserListener interface {
	antlr.ParseTreeListener

	// EnterModelUnit is called when entering the modelUnit production.
	EnterModelUnit(c *ModelUnitContext)

	// EnterNamespaceDeclaration is called when entering the namespaceDeclaration production.
	EnterNamespaceDeclaration(c *NamespaceDeclarationContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterClassModifier is called when entering the classModifier production.
	EnterClassModifier(c *ClassModifierContext)

	// EnterAssetDeclaration is called when entering the assetDeclaration production.
	EnterAssetDeclaration(c *AssetDeclarationContext)

	// EnterConceptDeclaration is called when entering the conceptDeclaration production.
	EnterConceptDeclaration(c *ConceptDeclarationContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterEnumConstant is called when entering the enumConstant production.
	EnterEnumConstant(c *EnumConstantContext)

	// EnterEventDeclaration is called when entering the eventDeclaration production.
	EnterEventDeclaration(c *EventDeclarationContext)

	// EnterParticipantDeclaration is called when entering the participantDeclaration production.
	EnterParticipantDeclaration(c *ParticipantDeclarationContext)

	// EnterTransactionDeclaration is called when entering the transactionDeclaration production.
	EnterTransactionDeclaration(c *TransactionDeclarationContext)

	// EnterExtendsOrIdentified is called when entering the extendsOrIdentified production.
	EnterExtendsOrIdentified(c *ExtendsOrIdentifiedContext)

	// EnterIdentified is called when entering the identified production.
	EnterIdentified(c *IdentifiedContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterClassBodyDeclaration is called when entering the classBodyDeclaration production.
	EnterClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterIdentifierField is called when entering the identifierField production.
	EnterIdentifierField(c *IdentifierFieldContext)

	// EnterNumericField is called when entering the numericField production.
	EnterNumericField(c *NumericFieldContext)

	// EnterNumericPrimitive is called when entering the numericPrimitive production.
	EnterNumericPrimitive(c *NumericPrimitiveContext)

	// EnterBooleanField is called when entering the booleanField production.
	EnterBooleanField(c *BooleanFieldContext)

	// EnterDateField is called when entering the dateField production.
	EnterDateField(c *DateFieldContext)

	// EnterDefaultDate is called when entering the defaultDate production.
	EnterDefaultDate(c *DefaultDateContext)

	// EnterRegexDeclaration is called when entering the regexDeclaration production.
	EnterRegexDeclaration(c *RegexDeclarationContext)

	// EnterStringField is called when entering the stringField production.
	EnterStringField(c *StringFieldContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterRangeValidation is called when entering the rangeValidation production.
	EnterRangeValidation(c *RangeValidationContext)

	// EnterRangeDeclaration is called when entering the rangeDeclaration production.
	EnterRangeDeclaration(c *RangeDeclarationContext)

	// EnterDefaultBoolean is called when entering the defaultBoolean production.
	EnterDefaultBoolean(c *DefaultBooleanContext)

	// EnterDefaultString is called when entering the defaultString production.
	EnterDefaultString(c *DefaultStringContext)

	// EnterDefaultNumber is called when entering the defaultNumber production.
	EnterDefaultNumber(c *DefaultNumberContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterDecorator is called when entering the decorator production.
	EnterDecorator(c *DecoratorContext)

	// EnterElementValuePair is called when entering the elementValuePair production.
	EnterElementValuePair(c *ElementValuePairContext)

	// EnterSquare is called when entering the square production.
	EnterSquare(c *SquareContext)

	// ExitModelUnit is called when exiting the modelUnit production.
	ExitModelUnit(c *ModelUnitContext)

	// ExitNamespaceDeclaration is called when exiting the namespaceDeclaration production.
	ExitNamespaceDeclaration(c *NamespaceDeclarationContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitClassModifier is called when exiting the classModifier production.
	ExitClassModifier(c *ClassModifierContext)

	// ExitAssetDeclaration is called when exiting the assetDeclaration production.
	ExitAssetDeclaration(c *AssetDeclarationContext)

	// ExitConceptDeclaration is called when exiting the conceptDeclaration production.
	ExitConceptDeclaration(c *ConceptDeclarationContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitEnumConstant is called when exiting the enumConstant production.
	ExitEnumConstant(c *EnumConstantContext)

	// ExitEventDeclaration is called when exiting the eventDeclaration production.
	ExitEventDeclaration(c *EventDeclarationContext)

	// ExitParticipantDeclaration is called when exiting the participantDeclaration production.
	ExitParticipantDeclaration(c *ParticipantDeclarationContext)

	// ExitTransactionDeclaration is called when exiting the transactionDeclaration production.
	ExitTransactionDeclaration(c *TransactionDeclarationContext)

	// ExitExtendsOrIdentified is called when exiting the extendsOrIdentified production.
	ExitExtendsOrIdentified(c *ExtendsOrIdentifiedContext)

	// ExitIdentified is called when exiting the identified production.
	ExitIdentified(c *IdentifiedContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitClassBodyDeclaration is called when exiting the classBodyDeclaration production.
	ExitClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitIdentifierField is called when exiting the identifierField production.
	ExitIdentifierField(c *IdentifierFieldContext)

	// ExitNumericField is called when exiting the numericField production.
	ExitNumericField(c *NumericFieldContext)

	// ExitNumericPrimitive is called when exiting the numericPrimitive production.
	ExitNumericPrimitive(c *NumericPrimitiveContext)

	// ExitBooleanField is called when exiting the booleanField production.
	ExitBooleanField(c *BooleanFieldContext)

	// ExitDateField is called when exiting the dateField production.
	ExitDateField(c *DateFieldContext)

	// ExitDefaultDate is called when exiting the defaultDate production.
	ExitDefaultDate(c *DefaultDateContext)

	// ExitRegexDeclaration is called when exiting the regexDeclaration production.
	ExitRegexDeclaration(c *RegexDeclarationContext)

	// ExitStringField is called when exiting the stringField production.
	ExitStringField(c *StringFieldContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitRangeValidation is called when exiting the rangeValidation production.
	ExitRangeValidation(c *RangeValidationContext)

	// ExitRangeDeclaration is called when exiting the rangeDeclaration production.
	ExitRangeDeclaration(c *RangeDeclarationContext)

	// ExitDefaultBoolean is called when exiting the defaultBoolean production.
	ExitDefaultBoolean(c *DefaultBooleanContext)

	// ExitDefaultString is called when exiting the defaultString production.
	ExitDefaultString(c *DefaultStringContext)

	// ExitDefaultNumber is called when exiting the defaultNumber production.
	ExitDefaultNumber(c *DefaultNumberContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitDecorator is called when exiting the decorator production.
	ExitDecorator(c *DecoratorContext)

	// ExitElementValuePair is called when exiting the elementValuePair production.
	ExitElementValuePair(c *ElementValuePairContext)

	// ExitSquare is called when exiting the square production.
	ExitSquare(c *SquareContext)
}
