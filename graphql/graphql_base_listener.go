// Code generated from GraphQL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package graphql // GraphQL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGraphQLListener is a complete listener for a parse tree produced by GraphQLParser.
type BaseGraphQLListener struct{}

var _ GraphQLListener = &BaseGraphQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGraphQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGraphQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGraphQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGraphQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseGraphQLListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseGraphQLListener) ExitDocument(ctx *DocumentContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseGraphQLListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseGraphQLListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterExecutableDefinition is called when production executableDefinition is entered.
func (s *BaseGraphQLListener) EnterExecutableDefinition(ctx *ExecutableDefinitionContext) {}

// ExitExecutableDefinition is called when production executableDefinition is exited.
func (s *BaseGraphQLListener) ExitExecutableDefinition(ctx *ExecutableDefinitionContext) {}

// EnterOperationDefinition is called when production operationDefinition is entered.
func (s *BaseGraphQLListener) EnterOperationDefinition(ctx *OperationDefinitionContext) {}

// ExitOperationDefinition is called when production operationDefinition is exited.
func (s *BaseGraphQLListener) ExitOperationDefinition(ctx *OperationDefinitionContext) {}

// EnterOperationType is called when production operationType is entered.
func (s *BaseGraphQLListener) EnterOperationType(ctx *OperationTypeContext) {}

// ExitOperationType is called when production operationType is exited.
func (s *BaseGraphQLListener) ExitOperationType(ctx *OperationTypeContext) {}

// EnterSelectionSet is called when production selectionSet is entered.
func (s *BaseGraphQLListener) EnterSelectionSet(ctx *SelectionSetContext) {}

// ExitSelectionSet is called when production selectionSet is exited.
func (s *BaseGraphQLListener) ExitSelectionSet(ctx *SelectionSetContext) {}

// EnterSelection is called when production selection is entered.
func (s *BaseGraphQLListener) EnterSelection(ctx *SelectionContext) {}

// ExitSelection is called when production selection is exited.
func (s *BaseGraphQLListener) ExitSelection(ctx *SelectionContext) {}

// EnterField is called when production field is entered.
func (s *BaseGraphQLListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseGraphQLListener) ExitField(ctx *FieldContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGraphQLListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGraphQLListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseGraphQLListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseGraphQLListener) ExitArgument(ctx *ArgumentContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseGraphQLListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseGraphQLListener) ExitAlias(ctx *AliasContext) {}

// EnterFragmentSpread is called when production fragmentSpread is entered.
func (s *BaseGraphQLListener) EnterFragmentSpread(ctx *FragmentSpreadContext) {}

// ExitFragmentSpread is called when production fragmentSpread is exited.
func (s *BaseGraphQLListener) ExitFragmentSpread(ctx *FragmentSpreadContext) {}

// EnterFragmentDefinition is called when production fragmentDefinition is entered.
func (s *BaseGraphQLListener) EnterFragmentDefinition(ctx *FragmentDefinitionContext) {}

// ExitFragmentDefinition is called when production fragmentDefinition is exited.
func (s *BaseGraphQLListener) ExitFragmentDefinition(ctx *FragmentDefinitionContext) {}

// EnterFragmentName is called when production fragmentName is entered.
func (s *BaseGraphQLListener) EnterFragmentName(ctx *FragmentNameContext) {}

// ExitFragmentName is called when production fragmentName is exited.
func (s *BaseGraphQLListener) ExitFragmentName(ctx *FragmentNameContext) {}

// EnterTypeCondition is called when production typeCondition is entered.
func (s *BaseGraphQLListener) EnterTypeCondition(ctx *TypeConditionContext) {}

// ExitTypeCondition is called when production typeCondition is exited.
func (s *BaseGraphQLListener) ExitTypeCondition(ctx *TypeConditionContext) {}

// EnterInlineFragment is called when production inlineFragment is entered.
func (s *BaseGraphQLListener) EnterInlineFragment(ctx *InlineFragmentContext) {}

// ExitInlineFragment is called when production inlineFragment is exited.
func (s *BaseGraphQLListener) ExitInlineFragment(ctx *InlineFragmentContext) {}

// EnterValue is called when production value is entered.
func (s *BaseGraphQLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseGraphQLListener) ExitValue(ctx *ValueContext) {}

// EnterIntValue is called when production intValue is entered.
func (s *BaseGraphQLListener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production intValue is exited.
func (s *BaseGraphQLListener) ExitIntValue(ctx *IntValueContext) {}

// EnterFloatValue is called when production floatValue is entered.
func (s *BaseGraphQLListener) EnterFloatValue(ctx *FloatValueContext) {}

// ExitFloatValue is called when production floatValue is exited.
func (s *BaseGraphQLListener) ExitFloatValue(ctx *FloatValueContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseGraphQLListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseGraphQLListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseGraphQLListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseGraphQLListener) ExitStringValue(ctx *StringValueContext) {}

// EnterNullValue is called when production nullValue is entered.
func (s *BaseGraphQLListener) EnterNullValue(ctx *NullValueContext) {}

// ExitNullValue is called when production nullValue is exited.
func (s *BaseGraphQLListener) ExitNullValue(ctx *NullValueContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseGraphQLListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseGraphQLListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterListValue is called when production listValue is entered.
func (s *BaseGraphQLListener) EnterListValue(ctx *ListValueContext) {}

// ExitListValue is called when production listValue is exited.
func (s *BaseGraphQLListener) ExitListValue(ctx *ListValueContext) {}

// EnterObjectValue is called when production objectValue is entered.
func (s *BaseGraphQLListener) EnterObjectValue(ctx *ObjectValueContext) {}

// ExitObjectValue is called when production objectValue is exited.
func (s *BaseGraphQLListener) ExitObjectValue(ctx *ObjectValueContext) {}

// EnterObjectField is called when production objectField is entered.
func (s *BaseGraphQLListener) EnterObjectField(ctx *ObjectFieldContext) {}

// ExitObjectField is called when production objectField is exited.
func (s *BaseGraphQLListener) ExitObjectField(ctx *ObjectFieldContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseGraphQLListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseGraphQLListener) ExitVariable(ctx *VariableContext) {}

// EnterVariableDefinitions is called when production variableDefinitions is entered.
func (s *BaseGraphQLListener) EnterVariableDefinitions(ctx *VariableDefinitionsContext) {}

// ExitVariableDefinitions is called when production variableDefinitions is exited.
func (s *BaseGraphQLListener) ExitVariableDefinitions(ctx *VariableDefinitionsContext) {}

// EnterVariableDefinition is called when production variableDefinition is entered.
func (s *BaseGraphQLListener) EnterVariableDefinition(ctx *VariableDefinitionContext) {}

// ExitVariableDefinition is called when production variableDefinition is exited.
func (s *BaseGraphQLListener) ExitVariableDefinition(ctx *VariableDefinitionContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *BaseGraphQLListener) EnterDefaultValue(ctx *DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *BaseGraphQLListener) ExitDefaultValue(ctx *DefaultValueContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseGraphQLListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseGraphQLListener) ExitType_(ctx *Type_Context) {}

// EnterNamedType is called when production namedType is entered.
func (s *BaseGraphQLListener) EnterNamedType(ctx *NamedTypeContext) {}

// ExitNamedType is called when production namedType is exited.
func (s *BaseGraphQLListener) ExitNamedType(ctx *NamedTypeContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseGraphQLListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseGraphQLListener) ExitListType(ctx *ListTypeContext) {}

// EnterDirectives is called when production directives is entered.
func (s *BaseGraphQLListener) EnterDirectives(ctx *DirectivesContext) {}

// ExitDirectives is called when production directives is exited.
func (s *BaseGraphQLListener) ExitDirectives(ctx *DirectivesContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseGraphQLListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseGraphQLListener) ExitDirective(ctx *DirectiveContext) {}

// EnterTypeSystemDefinition is called when production typeSystemDefinition is entered.
func (s *BaseGraphQLListener) EnterTypeSystemDefinition(ctx *TypeSystemDefinitionContext) {}

// ExitTypeSystemDefinition is called when production typeSystemDefinition is exited.
func (s *BaseGraphQLListener) ExitTypeSystemDefinition(ctx *TypeSystemDefinitionContext) {}

// EnterTypeSystemExtension is called when production typeSystemExtension is entered.
func (s *BaseGraphQLListener) EnterTypeSystemExtension(ctx *TypeSystemExtensionContext) {}

// ExitTypeSystemExtension is called when production typeSystemExtension is exited.
func (s *BaseGraphQLListener) ExitTypeSystemExtension(ctx *TypeSystemExtensionContext) {}

// EnterSchemaDefinition is called when production schemaDefinition is entered.
func (s *BaseGraphQLListener) EnterSchemaDefinition(ctx *SchemaDefinitionContext) {}

// ExitSchemaDefinition is called when production schemaDefinition is exited.
func (s *BaseGraphQLListener) ExitSchemaDefinition(ctx *SchemaDefinitionContext) {}

// EnterRootOperationTypeDefinition is called when production rootOperationTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterRootOperationTypeDefinition(ctx *RootOperationTypeDefinitionContext) {
}

// ExitRootOperationTypeDefinition is called when production rootOperationTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitRootOperationTypeDefinition(ctx *RootOperationTypeDefinitionContext) {
}

// EnterSchemaExtension is called when production schemaExtension is entered.
func (s *BaseGraphQLListener) EnterSchemaExtension(ctx *SchemaExtensionContext) {}

// ExitSchemaExtension is called when production schemaExtension is exited.
func (s *BaseGraphQLListener) ExitSchemaExtension(ctx *SchemaExtensionContext) {}

// EnterOperationTypeDefinition is called when production operationTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterOperationTypeDefinition(ctx *OperationTypeDefinitionContext) {}

// ExitOperationTypeDefinition is called when production operationTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitOperationTypeDefinition(ctx *OperationTypeDefinitionContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseGraphQLListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseGraphQLListener) ExitDescription(ctx *DescriptionContext) {}

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BaseGraphQLListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BaseGraphQLListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterTypeExtension is called when production typeExtension is entered.
func (s *BaseGraphQLListener) EnterTypeExtension(ctx *TypeExtensionContext) {}

// ExitTypeExtension is called when production typeExtension is exited.
func (s *BaseGraphQLListener) ExitTypeExtension(ctx *TypeExtensionContext) {}

// EnterScalarTypeDefinition is called when production scalarTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterScalarTypeDefinition(ctx *ScalarTypeDefinitionContext) {}

// ExitScalarTypeDefinition is called when production scalarTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitScalarTypeDefinition(ctx *ScalarTypeDefinitionContext) {}

// EnterScalarTypeExtension is called when production scalarTypeExtension is entered.
func (s *BaseGraphQLListener) EnterScalarTypeExtension(ctx *ScalarTypeExtensionContext) {}

// ExitScalarTypeExtension is called when production scalarTypeExtension is exited.
func (s *BaseGraphQLListener) ExitScalarTypeExtension(ctx *ScalarTypeExtensionContext) {}

// EnterObjectTypeDefinition is called when production objectTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterObjectTypeDefinition(ctx *ObjectTypeDefinitionContext) {}

// ExitObjectTypeDefinition is called when production objectTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitObjectTypeDefinition(ctx *ObjectTypeDefinitionContext) {}

// EnterImplementsInterfaces is called when production implementsInterfaces is entered.
func (s *BaseGraphQLListener) EnterImplementsInterfaces(ctx *ImplementsInterfacesContext) {}

// ExitImplementsInterfaces is called when production implementsInterfaces is exited.
func (s *BaseGraphQLListener) ExitImplementsInterfaces(ctx *ImplementsInterfacesContext) {}

// EnterFieldsDefinition is called when production fieldsDefinition is entered.
func (s *BaseGraphQLListener) EnterFieldsDefinition(ctx *FieldsDefinitionContext) {}

// ExitFieldsDefinition is called when production fieldsDefinition is exited.
func (s *BaseGraphQLListener) ExitFieldsDefinition(ctx *FieldsDefinitionContext) {}

// EnterFieldDefinition is called when production fieldDefinition is entered.
func (s *BaseGraphQLListener) EnterFieldDefinition(ctx *FieldDefinitionContext) {}

// ExitFieldDefinition is called when production fieldDefinition is exited.
func (s *BaseGraphQLListener) ExitFieldDefinition(ctx *FieldDefinitionContext) {}

// EnterArgumentsDefinition is called when production argumentsDefinition is entered.
func (s *BaseGraphQLListener) EnterArgumentsDefinition(ctx *ArgumentsDefinitionContext) {}

// ExitArgumentsDefinition is called when production argumentsDefinition is exited.
func (s *BaseGraphQLListener) ExitArgumentsDefinition(ctx *ArgumentsDefinitionContext) {}

// EnterInputValueDefinition is called when production inputValueDefinition is entered.
func (s *BaseGraphQLListener) EnterInputValueDefinition(ctx *InputValueDefinitionContext) {}

// ExitInputValueDefinition is called when production inputValueDefinition is exited.
func (s *BaseGraphQLListener) ExitInputValueDefinition(ctx *InputValueDefinitionContext) {}

// EnterObjectTypeExtension is called when production objectTypeExtension is entered.
func (s *BaseGraphQLListener) EnterObjectTypeExtension(ctx *ObjectTypeExtensionContext) {}

// ExitObjectTypeExtension is called when production objectTypeExtension is exited.
func (s *BaseGraphQLListener) ExitObjectTypeExtension(ctx *ObjectTypeExtensionContext) {}

// EnterInterfaceTypeDefinition is called when production interfaceTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterInterfaceTypeDefinition(ctx *InterfaceTypeDefinitionContext) {}

// ExitInterfaceTypeDefinition is called when production interfaceTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitInterfaceTypeDefinition(ctx *InterfaceTypeDefinitionContext) {}

// EnterInterfaceTypeExtension is called when production interfaceTypeExtension is entered.
func (s *BaseGraphQLListener) EnterInterfaceTypeExtension(ctx *InterfaceTypeExtensionContext) {}

// ExitInterfaceTypeExtension is called when production interfaceTypeExtension is exited.
func (s *BaseGraphQLListener) ExitInterfaceTypeExtension(ctx *InterfaceTypeExtensionContext) {}

// EnterUnionTypeDefinition is called when production unionTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterUnionTypeDefinition(ctx *UnionTypeDefinitionContext) {}

// ExitUnionTypeDefinition is called when production unionTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitUnionTypeDefinition(ctx *UnionTypeDefinitionContext) {}

// EnterUnionMemberTypes is called when production unionMemberTypes is entered.
func (s *BaseGraphQLListener) EnterUnionMemberTypes(ctx *UnionMemberTypesContext) {}

// ExitUnionMemberTypes is called when production unionMemberTypes is exited.
func (s *BaseGraphQLListener) ExitUnionMemberTypes(ctx *UnionMemberTypesContext) {}

// EnterUnionTypeExtension is called when production unionTypeExtension is entered.
func (s *BaseGraphQLListener) EnterUnionTypeExtension(ctx *UnionTypeExtensionContext) {}

// ExitUnionTypeExtension is called when production unionTypeExtension is exited.
func (s *BaseGraphQLListener) ExitUnionTypeExtension(ctx *UnionTypeExtensionContext) {}

// EnterEnumTypeDefinition is called when production enumTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterEnumTypeDefinition(ctx *EnumTypeDefinitionContext) {}

// ExitEnumTypeDefinition is called when production enumTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitEnumTypeDefinition(ctx *EnumTypeDefinitionContext) {}

// EnterEnumValuesDefinition is called when production enumValuesDefinition is entered.
func (s *BaseGraphQLListener) EnterEnumValuesDefinition(ctx *EnumValuesDefinitionContext) {}

// ExitEnumValuesDefinition is called when production enumValuesDefinition is exited.
func (s *BaseGraphQLListener) ExitEnumValuesDefinition(ctx *EnumValuesDefinitionContext) {}

// EnterEnumValueDefinition is called when production enumValueDefinition is entered.
func (s *BaseGraphQLListener) EnterEnumValueDefinition(ctx *EnumValueDefinitionContext) {}

// ExitEnumValueDefinition is called when production enumValueDefinition is exited.
func (s *BaseGraphQLListener) ExitEnumValueDefinition(ctx *EnumValueDefinitionContext) {}

// EnterEnumTypeExtension is called when production enumTypeExtension is entered.
func (s *BaseGraphQLListener) EnterEnumTypeExtension(ctx *EnumTypeExtensionContext) {}

// ExitEnumTypeExtension is called when production enumTypeExtension is exited.
func (s *BaseGraphQLListener) ExitEnumTypeExtension(ctx *EnumTypeExtensionContext) {}

// EnterInputObjectTypeDefinition is called when production inputObjectTypeDefinition is entered.
func (s *BaseGraphQLListener) EnterInputObjectTypeDefinition(ctx *InputObjectTypeDefinitionContext) {}

// ExitInputObjectTypeDefinition is called when production inputObjectTypeDefinition is exited.
func (s *BaseGraphQLListener) ExitInputObjectTypeDefinition(ctx *InputObjectTypeDefinitionContext) {}

// EnterInputFieldsDefinition is called when production inputFieldsDefinition is entered.
func (s *BaseGraphQLListener) EnterInputFieldsDefinition(ctx *InputFieldsDefinitionContext) {}

// ExitInputFieldsDefinition is called when production inputFieldsDefinition is exited.
func (s *BaseGraphQLListener) ExitInputFieldsDefinition(ctx *InputFieldsDefinitionContext) {}

// EnterInputObjectTypeExtension is called when production inputObjectTypeExtension is entered.
func (s *BaseGraphQLListener) EnterInputObjectTypeExtension(ctx *InputObjectTypeExtensionContext) {}

// ExitInputObjectTypeExtension is called when production inputObjectTypeExtension is exited.
func (s *BaseGraphQLListener) ExitInputObjectTypeExtension(ctx *InputObjectTypeExtensionContext) {}

// EnterDirectiveDefinition is called when production directiveDefinition is entered.
func (s *BaseGraphQLListener) EnterDirectiveDefinition(ctx *DirectiveDefinitionContext) {}

// ExitDirectiveDefinition is called when production directiveDefinition is exited.
func (s *BaseGraphQLListener) ExitDirectiveDefinition(ctx *DirectiveDefinitionContext) {}

// EnterDirectiveLocations is called when production directiveLocations is entered.
func (s *BaseGraphQLListener) EnterDirectiveLocations(ctx *DirectiveLocationsContext) {}

// ExitDirectiveLocations is called when production directiveLocations is exited.
func (s *BaseGraphQLListener) ExitDirectiveLocations(ctx *DirectiveLocationsContext) {}

// EnterDirectiveLocation is called when production directiveLocation is entered.
func (s *BaseGraphQLListener) EnterDirectiveLocation(ctx *DirectiveLocationContext) {}

// ExitDirectiveLocation is called when production directiveLocation is exited.
func (s *BaseGraphQLListener) ExitDirectiveLocation(ctx *DirectiveLocationContext) {}

// EnterExecutableDirectiveLocation is called when production executableDirectiveLocation is entered.
func (s *BaseGraphQLListener) EnterExecutableDirectiveLocation(ctx *ExecutableDirectiveLocationContext) {
}

// ExitExecutableDirectiveLocation is called when production executableDirectiveLocation is exited.
func (s *BaseGraphQLListener) ExitExecutableDirectiveLocation(ctx *ExecutableDirectiveLocationContext) {
}

// EnterTypeSystemDirectiveLocation is called when production typeSystemDirectiveLocation is entered.
func (s *BaseGraphQLListener) EnterTypeSystemDirectiveLocation(ctx *TypeSystemDirectiveLocationContext) {
}

// ExitTypeSystemDirectiveLocation is called when production typeSystemDirectiveLocation is exited.
func (s *BaseGraphQLListener) ExitTypeSystemDirectiveLocation(ctx *TypeSystemDirectiveLocationContext) {
}

// EnterName is called when production name is entered.
func (s *BaseGraphQLListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseGraphQLListener) ExitName(ctx *NameContext) {}
