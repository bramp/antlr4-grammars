// Code generated from CtoParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cto // CtoParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCtoParserListener is a complete listener for a parse tree produced by CtoParser.
type BaseCtoParserListener struct{}

var _ CtoParserListener = &BaseCtoParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCtoParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCtoParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCtoParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCtoParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModelUnit is called when production modelUnit is entered.
func (s *BaseCtoParserListener) EnterModelUnit(ctx *ModelUnitContext) {}

// ExitModelUnit is called when production modelUnit is exited.
func (s *BaseCtoParserListener) ExitModelUnit(ctx *ModelUnitContext) {}

// EnterNamespaceDeclaration is called when production namespaceDeclaration is entered.
func (s *BaseCtoParserListener) EnterNamespaceDeclaration(ctx *NamespaceDeclarationContext) {}

// ExitNamespaceDeclaration is called when production namespaceDeclaration is exited.
func (s *BaseCtoParserListener) ExitNamespaceDeclaration(ctx *NamespaceDeclarationContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseCtoParserListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseCtoParserListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseCtoParserListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseCtoParserListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterClassModifier is called when production classModifier is entered.
func (s *BaseCtoParserListener) EnterClassModifier(ctx *ClassModifierContext) {}

// ExitClassModifier is called when production classModifier is exited.
func (s *BaseCtoParserListener) ExitClassModifier(ctx *ClassModifierContext) {}

// EnterAssetDeclaration is called when production assetDeclaration is entered.
func (s *BaseCtoParserListener) EnterAssetDeclaration(ctx *AssetDeclarationContext) {}

// ExitAssetDeclaration is called when production assetDeclaration is exited.
func (s *BaseCtoParserListener) ExitAssetDeclaration(ctx *AssetDeclarationContext) {}

// EnterConceptDeclaration is called when production conceptDeclaration is entered.
func (s *BaseCtoParserListener) EnterConceptDeclaration(ctx *ConceptDeclarationContext) {}

// ExitConceptDeclaration is called when production conceptDeclaration is exited.
func (s *BaseCtoParserListener) ExitConceptDeclaration(ctx *ConceptDeclarationContext) {}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseCtoParserListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseCtoParserListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumConstant is called when production enumConstant is entered.
func (s *BaseCtoParserListener) EnterEnumConstant(ctx *EnumConstantContext) {}

// ExitEnumConstant is called when production enumConstant is exited.
func (s *BaseCtoParserListener) ExitEnumConstant(ctx *EnumConstantContext) {}

// EnterEventDeclaration is called when production eventDeclaration is entered.
func (s *BaseCtoParserListener) EnterEventDeclaration(ctx *EventDeclarationContext) {}

// ExitEventDeclaration is called when production eventDeclaration is exited.
func (s *BaseCtoParserListener) ExitEventDeclaration(ctx *EventDeclarationContext) {}

// EnterParticipantDeclaration is called when production participantDeclaration is entered.
func (s *BaseCtoParserListener) EnterParticipantDeclaration(ctx *ParticipantDeclarationContext) {}

// ExitParticipantDeclaration is called when production participantDeclaration is exited.
func (s *BaseCtoParserListener) ExitParticipantDeclaration(ctx *ParticipantDeclarationContext) {}

// EnterTransactionDeclaration is called when production transactionDeclaration is entered.
func (s *BaseCtoParserListener) EnterTransactionDeclaration(ctx *TransactionDeclarationContext) {}

// ExitTransactionDeclaration is called when production transactionDeclaration is exited.
func (s *BaseCtoParserListener) ExitTransactionDeclaration(ctx *TransactionDeclarationContext) {}

// EnterExtendsOrIdentified is called when production extendsOrIdentified is entered.
func (s *BaseCtoParserListener) EnterExtendsOrIdentified(ctx *ExtendsOrIdentifiedContext) {}

// ExitExtendsOrIdentified is called when production extendsOrIdentified is exited.
func (s *BaseCtoParserListener) ExitExtendsOrIdentified(ctx *ExtendsOrIdentifiedContext) {}

// EnterIdentified is called when production identified is entered.
func (s *BaseCtoParserListener) EnterIdentified(ctx *IdentifiedContext) {}

// ExitIdentified is called when production identified is exited.
func (s *BaseCtoParserListener) ExitIdentified(ctx *IdentifiedContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseCtoParserListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseCtoParserListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BaseCtoParserListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BaseCtoParserListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseCtoParserListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseCtoParserListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterIdentifierField is called when production identifierField is entered.
func (s *BaseCtoParserListener) EnterIdentifierField(ctx *IdentifierFieldContext) {}

// ExitIdentifierField is called when production identifierField is exited.
func (s *BaseCtoParserListener) ExitIdentifierField(ctx *IdentifierFieldContext) {}

// EnterNumericField is called when production numericField is entered.
func (s *BaseCtoParserListener) EnterNumericField(ctx *NumericFieldContext) {}

// ExitNumericField is called when production numericField is exited.
func (s *BaseCtoParserListener) ExitNumericField(ctx *NumericFieldContext) {}

// EnterNumericPrimitive is called when production numericPrimitive is entered.
func (s *BaseCtoParserListener) EnterNumericPrimitive(ctx *NumericPrimitiveContext) {}

// ExitNumericPrimitive is called when production numericPrimitive is exited.
func (s *BaseCtoParserListener) ExitNumericPrimitive(ctx *NumericPrimitiveContext) {}

// EnterBooleanField is called when production booleanField is entered.
func (s *BaseCtoParserListener) EnterBooleanField(ctx *BooleanFieldContext) {}

// ExitBooleanField is called when production booleanField is exited.
func (s *BaseCtoParserListener) ExitBooleanField(ctx *BooleanFieldContext) {}

// EnterDateField is called when production dateField is entered.
func (s *BaseCtoParserListener) EnterDateField(ctx *DateFieldContext) {}

// ExitDateField is called when production dateField is exited.
func (s *BaseCtoParserListener) ExitDateField(ctx *DateFieldContext) {}

// EnterDefaultDate is called when production defaultDate is entered.
func (s *BaseCtoParserListener) EnterDefaultDate(ctx *DefaultDateContext) {}

// ExitDefaultDate is called when production defaultDate is exited.
func (s *BaseCtoParserListener) ExitDefaultDate(ctx *DefaultDateContext) {}

// EnterRegexDeclaration is called when production regexDeclaration is entered.
func (s *BaseCtoParserListener) EnterRegexDeclaration(ctx *RegexDeclarationContext) {}

// ExitRegexDeclaration is called when production regexDeclaration is exited.
func (s *BaseCtoParserListener) ExitRegexDeclaration(ctx *RegexDeclarationContext) {}

// EnterStringField is called when production stringField is entered.
func (s *BaseCtoParserListener) EnterStringField(ctx *StringFieldContext) {}

// ExitStringField is called when production stringField is exited.
func (s *BaseCtoParserListener) ExitStringField(ctx *StringFieldContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseCtoParserListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseCtoParserListener) ExitReference(ctx *ReferenceContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseCtoParserListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseCtoParserListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterRangeValidation is called when production rangeValidation is entered.
func (s *BaseCtoParserListener) EnterRangeValidation(ctx *RangeValidationContext) {}

// ExitRangeValidation is called when production rangeValidation is exited.
func (s *BaseCtoParserListener) ExitRangeValidation(ctx *RangeValidationContext) {}

// EnterRangeDeclaration is called when production rangeDeclaration is entered.
func (s *BaseCtoParserListener) EnterRangeDeclaration(ctx *RangeDeclarationContext) {}

// ExitRangeDeclaration is called when production rangeDeclaration is exited.
func (s *BaseCtoParserListener) ExitRangeDeclaration(ctx *RangeDeclarationContext) {}

// EnterDefaultBoolean is called when production defaultBoolean is entered.
func (s *BaseCtoParserListener) EnterDefaultBoolean(ctx *DefaultBooleanContext) {}

// ExitDefaultBoolean is called when production defaultBoolean is exited.
func (s *BaseCtoParserListener) ExitDefaultBoolean(ctx *DefaultBooleanContext) {}

// EnterDefaultString is called when production defaultString is entered.
func (s *BaseCtoParserListener) EnterDefaultString(ctx *DefaultStringContext) {}

// ExitDefaultString is called when production defaultString is exited.
func (s *BaseCtoParserListener) ExitDefaultString(ctx *DefaultStringContext) {}

// EnterDefaultNumber is called when production defaultNumber is entered.
func (s *BaseCtoParserListener) EnterDefaultNumber(ctx *DefaultNumberContext) {}

// ExitDefaultNumber is called when production defaultNumber is exited.
func (s *BaseCtoParserListener) ExitDefaultNumber(ctx *DefaultNumberContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseCtoParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseCtoParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCtoParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCtoParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseCtoParserListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseCtoParserListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseCtoParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseCtoParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseCtoParserListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseCtoParserListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseCtoParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseCtoParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterDecorator is called when production decorator is entered.
func (s *BaseCtoParserListener) EnterDecorator(ctx *DecoratorContext) {}

// ExitDecorator is called when production decorator is exited.
func (s *BaseCtoParserListener) ExitDecorator(ctx *DecoratorContext) {}

// EnterElementValuePair is called when production elementValuePair is entered.
func (s *BaseCtoParserListener) EnterElementValuePair(ctx *ElementValuePairContext) {}

// ExitElementValuePair is called when production elementValuePair is exited.
func (s *BaseCtoParserListener) ExitElementValuePair(ctx *ElementValuePairContext) {}

// EnterSquare is called when production square is entered.
func (s *BaseCtoParserListener) EnterSquare(ctx *SquareContext) {}

// ExitSquare is called when production square is exited.
func (s *BaseCtoParserListener) ExitSquare(ctx *SquareContext) {}
