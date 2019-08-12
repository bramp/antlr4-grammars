// Code generated from ASN.g4 by ANTLR 4.7.2. DO NOT EDIT.

package asn // ASN
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseASNListener is a complete listener for a parse tree produced by ASNParser.
type BaseASNListener struct{}

var _ ASNListener = &BaseASNListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseASNListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseASNListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseASNListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseASNListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModules is called when production modules is entered.
func (s *BaseASNListener) EnterModules(ctx *ModulesContext) {}

// ExitModules is called when production modules is exited.
func (s *BaseASNListener) ExitModules(ctx *ModulesContext) {}

// EnterModuleDefinition is called when production moduleDefinition is entered.
func (s *BaseASNListener) EnterModuleDefinition(ctx *ModuleDefinitionContext) {}

// ExitModuleDefinition is called when production moduleDefinition is exited.
func (s *BaseASNListener) ExitModuleDefinition(ctx *ModuleDefinitionContext) {}

// EnterTagDefault is called when production tagDefault is entered.
func (s *BaseASNListener) EnterTagDefault(ctx *TagDefaultContext) {}

// ExitTagDefault is called when production tagDefault is exited.
func (s *BaseASNListener) ExitTagDefault(ctx *TagDefaultContext) {}

// EnterExtensionDefault is called when production extensionDefault is entered.
func (s *BaseASNListener) EnterExtensionDefault(ctx *ExtensionDefaultContext) {}

// ExitExtensionDefault is called when production extensionDefault is exited.
func (s *BaseASNListener) ExitExtensionDefault(ctx *ExtensionDefaultContext) {}

// EnterModuleBody is called when production moduleBody is entered.
func (s *BaseASNListener) EnterModuleBody(ctx *ModuleBodyContext) {}

// ExitModuleBody is called when production moduleBody is exited.
func (s *BaseASNListener) ExitModuleBody(ctx *ModuleBodyContext) {}

// EnterExports is called when production exports is entered.
func (s *BaseASNListener) EnterExports(ctx *ExportsContext) {}

// ExitExports is called when production exports is exited.
func (s *BaseASNListener) ExitExports(ctx *ExportsContext) {}

// EnterSymbolsExported is called when production symbolsExported is entered.
func (s *BaseASNListener) EnterSymbolsExported(ctx *SymbolsExportedContext) {}

// ExitSymbolsExported is called when production symbolsExported is exited.
func (s *BaseASNListener) ExitSymbolsExported(ctx *SymbolsExportedContext) {}

// EnterImports is called when production imports is entered.
func (s *BaseASNListener) EnterImports(ctx *ImportsContext) {}

// ExitImports is called when production imports is exited.
func (s *BaseASNListener) ExitImports(ctx *ImportsContext) {}

// EnterSymbolsImported is called when production symbolsImported is entered.
func (s *BaseASNListener) EnterSymbolsImported(ctx *SymbolsImportedContext) {}

// ExitSymbolsImported is called when production symbolsImported is exited.
func (s *BaseASNListener) ExitSymbolsImported(ctx *SymbolsImportedContext) {}

// EnterSymbolsFromModuleList is called when production symbolsFromModuleList is entered.
func (s *BaseASNListener) EnterSymbolsFromModuleList(ctx *SymbolsFromModuleListContext) {}

// ExitSymbolsFromModuleList is called when production symbolsFromModuleList is exited.
func (s *BaseASNListener) ExitSymbolsFromModuleList(ctx *SymbolsFromModuleListContext) {}

// EnterSymbolsFromModule is called when production symbolsFromModule is entered.
func (s *BaseASNListener) EnterSymbolsFromModule(ctx *SymbolsFromModuleContext) {}

// ExitSymbolsFromModule is called when production symbolsFromModule is exited.
func (s *BaseASNListener) ExitSymbolsFromModule(ctx *SymbolsFromModuleContext) {}

// EnterGlobalModuleReference is called when production globalModuleReference is entered.
func (s *BaseASNListener) EnterGlobalModuleReference(ctx *GlobalModuleReferenceContext) {}

// ExitGlobalModuleReference is called when production globalModuleReference is exited.
func (s *BaseASNListener) ExitGlobalModuleReference(ctx *GlobalModuleReferenceContext) {}

// EnterAssignedIdentifier is called when production assignedIdentifier is entered.
func (s *BaseASNListener) EnterAssignedIdentifier(ctx *AssignedIdentifierContext) {}

// ExitAssignedIdentifier is called when production assignedIdentifier is exited.
func (s *BaseASNListener) ExitAssignedIdentifier(ctx *AssignedIdentifierContext) {}

// EnterSymbolList is called when production symbolList is entered.
func (s *BaseASNListener) EnterSymbolList(ctx *SymbolListContext) {}

// ExitSymbolList is called when production symbolList is exited.
func (s *BaseASNListener) ExitSymbolList(ctx *SymbolListContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseASNListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseASNListener) ExitSymbol(ctx *SymbolContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseASNListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseASNListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseASNListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseASNListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterSequenceType is called when production sequenceType is entered.
func (s *BaseASNListener) EnterSequenceType(ctx *SequenceTypeContext) {}

// ExitSequenceType is called when production sequenceType is exited.
func (s *BaseASNListener) ExitSequenceType(ctx *SequenceTypeContext) {}

// EnterExtensionAndException is called when production extensionAndException is entered.
func (s *BaseASNListener) EnterExtensionAndException(ctx *ExtensionAndExceptionContext) {}

// ExitExtensionAndException is called when production extensionAndException is exited.
func (s *BaseASNListener) ExitExtensionAndException(ctx *ExtensionAndExceptionContext) {}

// EnterOptionalExtensionMarker is called when production optionalExtensionMarker is entered.
func (s *BaseASNListener) EnterOptionalExtensionMarker(ctx *OptionalExtensionMarkerContext) {}

// ExitOptionalExtensionMarker is called when production optionalExtensionMarker is exited.
func (s *BaseASNListener) ExitOptionalExtensionMarker(ctx *OptionalExtensionMarkerContext) {}

// EnterComponentTypeLists is called when production componentTypeLists is entered.
func (s *BaseASNListener) EnterComponentTypeLists(ctx *ComponentTypeListsContext) {}

// ExitComponentTypeLists is called when production componentTypeLists is exited.
func (s *BaseASNListener) ExitComponentTypeLists(ctx *ComponentTypeListsContext) {}

// EnterRootComponentTypeList is called when production rootComponentTypeList is entered.
func (s *BaseASNListener) EnterRootComponentTypeList(ctx *RootComponentTypeListContext) {}

// ExitRootComponentTypeList is called when production rootComponentTypeList is exited.
func (s *BaseASNListener) ExitRootComponentTypeList(ctx *RootComponentTypeListContext) {}

// EnterComponentTypeList is called when production componentTypeList is entered.
func (s *BaseASNListener) EnterComponentTypeList(ctx *ComponentTypeListContext) {}

// ExitComponentTypeList is called when production componentTypeList is exited.
func (s *BaseASNListener) ExitComponentTypeList(ctx *ComponentTypeListContext) {}

// EnterComponentType is called when production componentType is entered.
func (s *BaseASNListener) EnterComponentType(ctx *ComponentTypeContext) {}

// ExitComponentType is called when production componentType is exited.
func (s *BaseASNListener) ExitComponentType(ctx *ComponentTypeContext) {}

// EnterExtensionAdditions is called when production extensionAdditions is entered.
func (s *BaseASNListener) EnterExtensionAdditions(ctx *ExtensionAdditionsContext) {}

// ExitExtensionAdditions is called when production extensionAdditions is exited.
func (s *BaseASNListener) ExitExtensionAdditions(ctx *ExtensionAdditionsContext) {}

// EnterExtensionAdditionList is called when production extensionAdditionList is entered.
func (s *BaseASNListener) EnterExtensionAdditionList(ctx *ExtensionAdditionListContext) {}

// ExitExtensionAdditionList is called when production extensionAdditionList is exited.
func (s *BaseASNListener) ExitExtensionAdditionList(ctx *ExtensionAdditionListContext) {}

// EnterExtensionAddition is called when production extensionAddition is entered.
func (s *BaseASNListener) EnterExtensionAddition(ctx *ExtensionAdditionContext) {}

// ExitExtensionAddition is called when production extensionAddition is exited.
func (s *BaseASNListener) ExitExtensionAddition(ctx *ExtensionAdditionContext) {}

// EnterExtensionAdditionGroup is called when production extensionAdditionGroup is entered.
func (s *BaseASNListener) EnterExtensionAdditionGroup(ctx *ExtensionAdditionGroupContext) {}

// ExitExtensionAdditionGroup is called when production extensionAdditionGroup is exited.
func (s *BaseASNListener) ExitExtensionAdditionGroup(ctx *ExtensionAdditionGroupContext) {}

// EnterVersionNumber is called when production versionNumber is entered.
func (s *BaseASNListener) EnterVersionNumber(ctx *VersionNumberContext) {}

// ExitVersionNumber is called when production versionNumber is exited.
func (s *BaseASNListener) ExitVersionNumber(ctx *VersionNumberContext) {}

// EnterSequenceOfType is called when production sequenceOfType is entered.
func (s *BaseASNListener) EnterSequenceOfType(ctx *SequenceOfTypeContext) {}

// ExitSequenceOfType is called when production sequenceOfType is exited.
func (s *BaseASNListener) ExitSequenceOfType(ctx *SequenceOfTypeContext) {}

// EnterSizeConstraint is called when production sizeConstraint is entered.
func (s *BaseASNListener) EnterSizeConstraint(ctx *SizeConstraintContext) {}

// ExitSizeConstraint is called when production sizeConstraint is exited.
func (s *BaseASNListener) ExitSizeConstraint(ctx *SizeConstraintContext) {}

// EnterParameterizedAssignment is called when production parameterizedAssignment is entered.
func (s *BaseASNListener) EnterParameterizedAssignment(ctx *ParameterizedAssignmentContext) {}

// ExitParameterizedAssignment is called when production parameterizedAssignment is exited.
func (s *BaseASNListener) ExitParameterizedAssignment(ctx *ParameterizedAssignmentContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseASNListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseASNListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseASNListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseASNListener) ExitParameter(ctx *ParameterContext) {}

// EnterParamGovernor is called when production paramGovernor is entered.
func (s *BaseASNListener) EnterParamGovernor(ctx *ParamGovernorContext) {}

// ExitParamGovernor is called when production paramGovernor is exited.
func (s *BaseASNListener) ExitParamGovernor(ctx *ParamGovernorContext) {}

// EnterGovernor is called when production governor is entered.
func (s *BaseASNListener) EnterGovernor(ctx *GovernorContext) {}

// ExitGovernor is called when production governor is exited.
func (s *BaseASNListener) ExitGovernor(ctx *GovernorContext) {}

// EnterObjectClassAssignment is called when production objectClassAssignment is entered.
func (s *BaseASNListener) EnterObjectClassAssignment(ctx *ObjectClassAssignmentContext) {}

// ExitObjectClassAssignment is called when production objectClassAssignment is exited.
func (s *BaseASNListener) ExitObjectClassAssignment(ctx *ObjectClassAssignmentContext) {}

// EnterObjectClass is called when production objectClass is entered.
func (s *BaseASNListener) EnterObjectClass(ctx *ObjectClassContext) {}

// ExitObjectClass is called when production objectClass is exited.
func (s *BaseASNListener) ExitObjectClass(ctx *ObjectClassContext) {}

// EnterDefinedObjectClass is called when production definedObjectClass is entered.
func (s *BaseASNListener) EnterDefinedObjectClass(ctx *DefinedObjectClassContext) {}

// ExitDefinedObjectClass is called when production definedObjectClass is exited.
func (s *BaseASNListener) ExitDefinedObjectClass(ctx *DefinedObjectClassContext) {}

// EnterUsefulObjectClassReference is called when production usefulObjectClassReference is entered.
func (s *BaseASNListener) EnterUsefulObjectClassReference(ctx *UsefulObjectClassReferenceContext) {}

// ExitUsefulObjectClassReference is called when production usefulObjectClassReference is exited.
func (s *BaseASNListener) ExitUsefulObjectClassReference(ctx *UsefulObjectClassReferenceContext) {}

// EnterExternalObjectClassReference is called when production externalObjectClassReference is entered.
func (s *BaseASNListener) EnterExternalObjectClassReference(ctx *ExternalObjectClassReferenceContext) {
}

// ExitExternalObjectClassReference is called when production externalObjectClassReference is exited.
func (s *BaseASNListener) ExitExternalObjectClassReference(ctx *ExternalObjectClassReferenceContext) {}

// EnterObjectClassDefn is called when production objectClassDefn is entered.
func (s *BaseASNListener) EnterObjectClassDefn(ctx *ObjectClassDefnContext) {}

// ExitObjectClassDefn is called when production objectClassDefn is exited.
func (s *BaseASNListener) ExitObjectClassDefn(ctx *ObjectClassDefnContext) {}

// EnterWithSyntaxSpec is called when production withSyntaxSpec is entered.
func (s *BaseASNListener) EnterWithSyntaxSpec(ctx *WithSyntaxSpecContext) {}

// ExitWithSyntaxSpec is called when production withSyntaxSpec is exited.
func (s *BaseASNListener) ExitWithSyntaxSpec(ctx *WithSyntaxSpecContext) {}

// EnterSyntaxList is called when production syntaxList is entered.
func (s *BaseASNListener) EnterSyntaxList(ctx *SyntaxListContext) {}

// ExitSyntaxList is called when production syntaxList is exited.
func (s *BaseASNListener) ExitSyntaxList(ctx *SyntaxListContext) {}

// EnterTokenOrGroupSpec is called when production tokenOrGroupSpec is entered.
func (s *BaseASNListener) EnterTokenOrGroupSpec(ctx *TokenOrGroupSpecContext) {}

// ExitTokenOrGroupSpec is called when production tokenOrGroupSpec is exited.
func (s *BaseASNListener) ExitTokenOrGroupSpec(ctx *TokenOrGroupSpecContext) {}

// EnterOptionalGroup is called when production optionalGroup is entered.
func (s *BaseASNListener) EnterOptionalGroup(ctx *OptionalGroupContext) {}

// ExitOptionalGroup is called when production optionalGroup is exited.
func (s *BaseASNListener) ExitOptionalGroup(ctx *OptionalGroupContext) {}

// EnterRequiredToken is called when production requiredToken is entered.
func (s *BaseASNListener) EnterRequiredToken(ctx *RequiredTokenContext) {}

// ExitRequiredToken is called when production requiredToken is exited.
func (s *BaseASNListener) ExitRequiredToken(ctx *RequiredTokenContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseASNListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseASNListener) ExitLiteral(ctx *LiteralContext) {}

// EnterPrimitiveFieldName is called when production primitiveFieldName is entered.
func (s *BaseASNListener) EnterPrimitiveFieldName(ctx *PrimitiveFieldNameContext) {}

// ExitPrimitiveFieldName is called when production primitiveFieldName is exited.
func (s *BaseASNListener) ExitPrimitiveFieldName(ctx *PrimitiveFieldNameContext) {}

// EnterFieldSpec is called when production fieldSpec is entered.
func (s *BaseASNListener) EnterFieldSpec(ctx *FieldSpecContext) {}

// ExitFieldSpec is called when production fieldSpec is exited.
func (s *BaseASNListener) ExitFieldSpec(ctx *FieldSpecContext) {}

// EnterTypeFieldSpec is called when production typeFieldSpec is entered.
func (s *BaseASNListener) EnterTypeFieldSpec(ctx *TypeFieldSpecContext) {}

// ExitTypeFieldSpec is called when production typeFieldSpec is exited.
func (s *BaseASNListener) ExitTypeFieldSpec(ctx *TypeFieldSpecContext) {}

// EnterTypeOptionalitySpec is called when production typeOptionalitySpec is entered.
func (s *BaseASNListener) EnterTypeOptionalitySpec(ctx *TypeOptionalitySpecContext) {}

// ExitTypeOptionalitySpec is called when production typeOptionalitySpec is exited.
func (s *BaseASNListener) ExitTypeOptionalitySpec(ctx *TypeOptionalitySpecContext) {}

// EnterFixedTypeValueFieldSpec is called when production fixedTypeValueFieldSpec is entered.
func (s *BaseASNListener) EnterFixedTypeValueFieldSpec(ctx *FixedTypeValueFieldSpecContext) {}

// ExitFixedTypeValueFieldSpec is called when production fixedTypeValueFieldSpec is exited.
func (s *BaseASNListener) ExitFixedTypeValueFieldSpec(ctx *FixedTypeValueFieldSpecContext) {}

// EnterValueOptionalitySpec is called when production valueOptionalitySpec is entered.
func (s *BaseASNListener) EnterValueOptionalitySpec(ctx *ValueOptionalitySpecContext) {}

// ExitValueOptionalitySpec is called when production valueOptionalitySpec is exited.
func (s *BaseASNListener) ExitValueOptionalitySpec(ctx *ValueOptionalitySpecContext) {}

// EnterVariableTypeValueFieldSpec is called when production variableTypeValueFieldSpec is entered.
func (s *BaseASNListener) EnterVariableTypeValueFieldSpec(ctx *VariableTypeValueFieldSpecContext) {}

// ExitVariableTypeValueFieldSpec is called when production variableTypeValueFieldSpec is exited.
func (s *BaseASNListener) ExitVariableTypeValueFieldSpec(ctx *VariableTypeValueFieldSpecContext) {}

// EnterFixedTypeValueSetFieldSpec is called when production fixedTypeValueSetFieldSpec is entered.
func (s *BaseASNListener) EnterFixedTypeValueSetFieldSpec(ctx *FixedTypeValueSetFieldSpecContext) {}

// ExitFixedTypeValueSetFieldSpec is called when production fixedTypeValueSetFieldSpec is exited.
func (s *BaseASNListener) ExitFixedTypeValueSetFieldSpec(ctx *FixedTypeValueSetFieldSpecContext) {}

// EnterValueSetOptionalitySpec is called when production valueSetOptionalitySpec is entered.
func (s *BaseASNListener) EnterValueSetOptionalitySpec(ctx *ValueSetOptionalitySpecContext) {}

// ExitValueSetOptionalitySpec is called when production valueSetOptionalitySpec is exited.
func (s *BaseASNListener) ExitValueSetOptionalitySpec(ctx *ValueSetOptionalitySpecContext) {}

// EnterObject is called when production object is entered.
func (s *BaseASNListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseASNListener) ExitObject(ctx *ObjectContext) {}

// EnterParameterizedObject is called when production parameterizedObject is entered.
func (s *BaseASNListener) EnterParameterizedObject(ctx *ParameterizedObjectContext) {}

// ExitParameterizedObject is called when production parameterizedObject is exited.
func (s *BaseASNListener) ExitParameterizedObject(ctx *ParameterizedObjectContext) {}

// EnterDefinedObject is called when production definedObject is entered.
func (s *BaseASNListener) EnterDefinedObject(ctx *DefinedObjectContext) {}

// ExitDefinedObject is called when production definedObject is exited.
func (s *BaseASNListener) ExitDefinedObject(ctx *DefinedObjectContext) {}

// EnterObjectSet is called when production objectSet is entered.
func (s *BaseASNListener) EnterObjectSet(ctx *ObjectSetContext) {}

// ExitObjectSet is called when production objectSet is exited.
func (s *BaseASNListener) ExitObjectSet(ctx *ObjectSetContext) {}

// EnterObjectSetSpec is called when production objectSetSpec is entered.
func (s *BaseASNListener) EnterObjectSetSpec(ctx *ObjectSetSpecContext) {}

// ExitObjectSetSpec is called when production objectSetSpec is exited.
func (s *BaseASNListener) ExitObjectSetSpec(ctx *ObjectSetSpecContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseASNListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseASNListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterValueSet is called when production valueSet is entered.
func (s *BaseASNListener) EnterValueSet(ctx *ValueSetContext) {}

// ExitValueSet is called when production valueSet is exited.
func (s *BaseASNListener) ExitValueSet(ctx *ValueSetContext) {}

// EnterElementSetSpecs is called when production elementSetSpecs is entered.
func (s *BaseASNListener) EnterElementSetSpecs(ctx *ElementSetSpecsContext) {}

// ExitElementSetSpecs is called when production elementSetSpecs is exited.
func (s *BaseASNListener) ExitElementSetSpecs(ctx *ElementSetSpecsContext) {}

// EnterRootElementSetSpec is called when production rootElementSetSpec is entered.
func (s *BaseASNListener) EnterRootElementSetSpec(ctx *RootElementSetSpecContext) {}

// ExitRootElementSetSpec is called when production rootElementSetSpec is exited.
func (s *BaseASNListener) ExitRootElementSetSpec(ctx *RootElementSetSpecContext) {}

// EnterAdditionalElementSetSpec is called when production additionalElementSetSpec is entered.
func (s *BaseASNListener) EnterAdditionalElementSetSpec(ctx *AdditionalElementSetSpecContext) {}

// ExitAdditionalElementSetSpec is called when production additionalElementSetSpec is exited.
func (s *BaseASNListener) ExitAdditionalElementSetSpec(ctx *AdditionalElementSetSpecContext) {}

// EnterElementSetSpec is called when production elementSetSpec is entered.
func (s *BaseASNListener) EnterElementSetSpec(ctx *ElementSetSpecContext) {}

// ExitElementSetSpec is called when production elementSetSpec is exited.
func (s *BaseASNListener) ExitElementSetSpec(ctx *ElementSetSpecContext) {}

// EnterUnions is called when production unions is entered.
func (s *BaseASNListener) EnterUnions(ctx *UnionsContext) {}

// ExitUnions is called when production unions is exited.
func (s *BaseASNListener) ExitUnions(ctx *UnionsContext) {}

// EnterExclusions is called when production exclusions is entered.
func (s *BaseASNListener) EnterExclusions(ctx *ExclusionsContext) {}

// ExitExclusions is called when production exclusions is exited.
func (s *BaseASNListener) ExitExclusions(ctx *ExclusionsContext) {}

// EnterIntersections is called when production intersections is entered.
func (s *BaseASNListener) EnterIntersections(ctx *IntersectionsContext) {}

// ExitIntersections is called when production intersections is exited.
func (s *BaseASNListener) ExitIntersections(ctx *IntersectionsContext) {}

// EnterUnionMark is called when production unionMark is entered.
func (s *BaseASNListener) EnterUnionMark(ctx *UnionMarkContext) {}

// ExitUnionMark is called when production unionMark is exited.
func (s *BaseASNListener) ExitUnionMark(ctx *UnionMarkContext) {}

// EnterIntersectionMark is called when production intersectionMark is entered.
func (s *BaseASNListener) EnterIntersectionMark(ctx *IntersectionMarkContext) {}

// ExitIntersectionMark is called when production intersectionMark is exited.
func (s *BaseASNListener) ExitIntersectionMark(ctx *IntersectionMarkContext) {}

// EnterElements is called when production elements is entered.
func (s *BaseASNListener) EnterElements(ctx *ElementsContext) {}

// ExitElements is called when production elements is exited.
func (s *BaseASNListener) ExitElements(ctx *ElementsContext) {}

// EnterObjectSetElements is called when production objectSetElements is entered.
func (s *BaseASNListener) EnterObjectSetElements(ctx *ObjectSetElementsContext) {}

// ExitObjectSetElements is called when production objectSetElements is exited.
func (s *BaseASNListener) ExitObjectSetElements(ctx *ObjectSetElementsContext) {}

// EnterIntersectionElements is called when production intersectionElements is entered.
func (s *BaseASNListener) EnterIntersectionElements(ctx *IntersectionElementsContext) {}

// ExitIntersectionElements is called when production intersectionElements is exited.
func (s *BaseASNListener) ExitIntersectionElements(ctx *IntersectionElementsContext) {}

// EnterSubtypeElements is called when production subtypeElements is entered.
func (s *BaseASNListener) EnterSubtypeElements(ctx *SubtypeElementsContext) {}

// ExitSubtypeElements is called when production subtypeElements is exited.
func (s *BaseASNListener) ExitSubtypeElements(ctx *SubtypeElementsContext) {}

// EnterVariableTypeValueSetFieldSpec is called when production variableTypeValueSetFieldSpec is entered.
func (s *BaseASNListener) EnterVariableTypeValueSetFieldSpec(ctx *VariableTypeValueSetFieldSpecContext) {
}

// ExitVariableTypeValueSetFieldSpec is called when production variableTypeValueSetFieldSpec is exited.
func (s *BaseASNListener) ExitVariableTypeValueSetFieldSpec(ctx *VariableTypeValueSetFieldSpecContext) {
}

// EnterObjectFieldSpec is called when production objectFieldSpec is entered.
func (s *BaseASNListener) EnterObjectFieldSpec(ctx *ObjectFieldSpecContext) {}

// ExitObjectFieldSpec is called when production objectFieldSpec is exited.
func (s *BaseASNListener) ExitObjectFieldSpec(ctx *ObjectFieldSpecContext) {}

// EnterObjectOptionalitySpec is called when production objectOptionalitySpec is entered.
func (s *BaseASNListener) EnterObjectOptionalitySpec(ctx *ObjectOptionalitySpecContext) {}

// ExitObjectOptionalitySpec is called when production objectOptionalitySpec is exited.
func (s *BaseASNListener) ExitObjectOptionalitySpec(ctx *ObjectOptionalitySpecContext) {}

// EnterObjectSetFieldSpec is called when production objectSetFieldSpec is entered.
func (s *BaseASNListener) EnterObjectSetFieldSpec(ctx *ObjectSetFieldSpecContext) {}

// ExitObjectSetFieldSpec is called when production objectSetFieldSpec is exited.
func (s *BaseASNListener) ExitObjectSetFieldSpec(ctx *ObjectSetFieldSpecContext) {}

// EnterObjectSetOptionalitySpec is called when production objectSetOptionalitySpec is entered.
func (s *BaseASNListener) EnterObjectSetOptionalitySpec(ctx *ObjectSetOptionalitySpecContext) {}

// ExitObjectSetOptionalitySpec is called when production objectSetOptionalitySpec is exited.
func (s *BaseASNListener) ExitObjectSetOptionalitySpec(ctx *ObjectSetOptionalitySpecContext) {}

// EnterTypeAssignment is called when production typeAssignment is entered.
func (s *BaseASNListener) EnterTypeAssignment(ctx *TypeAssignmentContext) {}

// ExitTypeAssignment is called when production typeAssignment is exited.
func (s *BaseASNListener) ExitTypeAssignment(ctx *TypeAssignmentContext) {}

// EnterValueAssignment is called when production valueAssignment is entered.
func (s *BaseASNListener) EnterValueAssignment(ctx *ValueAssignmentContext) {}

// ExitValueAssignment is called when production valueAssignment is exited.
func (s *BaseASNListener) ExitValueAssignment(ctx *ValueAssignmentContext) {}

// EnterAsnType is called when production asnType is entered.
func (s *BaseASNListener) EnterAsnType(ctx *AsnTypeContext) {}

// ExitAsnType is called when production asnType is exited.
func (s *BaseASNListener) ExitAsnType(ctx *AsnTypeContext) {}

// EnterBuiltinType is called when production builtinType is entered.
func (s *BaseASNListener) EnterBuiltinType(ctx *BuiltinTypeContext) {}

// ExitBuiltinType is called when production builtinType is exited.
func (s *BaseASNListener) ExitBuiltinType(ctx *BuiltinTypeContext) {}

// EnterObjectClassFieldType is called when production objectClassFieldType is entered.
func (s *BaseASNListener) EnterObjectClassFieldType(ctx *ObjectClassFieldTypeContext) {}

// ExitObjectClassFieldType is called when production objectClassFieldType is exited.
func (s *BaseASNListener) ExitObjectClassFieldType(ctx *ObjectClassFieldTypeContext) {}

// EnterSetType is called when production setType is entered.
func (s *BaseASNListener) EnterSetType(ctx *SetTypeContext) {}

// ExitSetType is called when production setType is exited.
func (s *BaseASNListener) ExitSetType(ctx *SetTypeContext) {}

// EnterSetOfType is called when production setOfType is entered.
func (s *BaseASNListener) EnterSetOfType(ctx *SetOfTypeContext) {}

// ExitSetOfType is called when production setOfType is exited.
func (s *BaseASNListener) ExitSetOfType(ctx *SetOfTypeContext) {}

// EnterReferencedType is called when production referencedType is entered.
func (s *BaseASNListener) EnterReferencedType(ctx *ReferencedTypeContext) {}

// ExitReferencedType is called when production referencedType is exited.
func (s *BaseASNListener) ExitReferencedType(ctx *ReferencedTypeContext) {}

// EnterDefinedType is called when production definedType is entered.
func (s *BaseASNListener) EnterDefinedType(ctx *DefinedTypeContext) {}

// ExitDefinedType is called when production definedType is exited.
func (s *BaseASNListener) ExitDefinedType(ctx *DefinedTypeContext) {}

// EnterConstraint is called when production constraint is entered.
func (s *BaseASNListener) EnterConstraint(ctx *ConstraintContext) {}

// ExitConstraint is called when production constraint is exited.
func (s *BaseASNListener) ExitConstraint(ctx *ConstraintContext) {}

// EnterConstraintSpec is called when production constraintSpec is entered.
func (s *BaseASNListener) EnterConstraintSpec(ctx *ConstraintSpecContext) {}

// ExitConstraintSpec is called when production constraintSpec is exited.
func (s *BaseASNListener) ExitConstraintSpec(ctx *ConstraintSpecContext) {}

// EnterUserDefinedConstraint is called when production userDefinedConstraint is entered.
func (s *BaseASNListener) EnterUserDefinedConstraint(ctx *UserDefinedConstraintContext) {}

// ExitUserDefinedConstraint is called when production userDefinedConstraint is exited.
func (s *BaseASNListener) ExitUserDefinedConstraint(ctx *UserDefinedConstraintContext) {}

// EnterGeneralConstraint is called when production generalConstraint is entered.
func (s *BaseASNListener) EnterGeneralConstraint(ctx *GeneralConstraintContext) {}

// ExitGeneralConstraint is called when production generalConstraint is exited.
func (s *BaseASNListener) ExitGeneralConstraint(ctx *GeneralConstraintContext) {}

// EnterUserDefinedConstraintParameter is called when production userDefinedConstraintParameter is entered.
func (s *BaseASNListener) EnterUserDefinedConstraintParameter(ctx *UserDefinedConstraintParameterContext) {
}

// ExitUserDefinedConstraintParameter is called when production userDefinedConstraintParameter is exited.
func (s *BaseASNListener) ExitUserDefinedConstraintParameter(ctx *UserDefinedConstraintParameterContext) {
}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseASNListener) EnterTableConstraint(ctx *TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseASNListener) ExitTableConstraint(ctx *TableConstraintContext) {}

// EnterSimpleTableConstraint is called when production simpleTableConstraint is entered.
func (s *BaseASNListener) EnterSimpleTableConstraint(ctx *SimpleTableConstraintContext) {}

// ExitSimpleTableConstraint is called when production simpleTableConstraint is exited.
func (s *BaseASNListener) ExitSimpleTableConstraint(ctx *SimpleTableConstraintContext) {}

// EnterContentsConstraint is called when production contentsConstraint is entered.
func (s *BaseASNListener) EnterContentsConstraint(ctx *ContentsConstraintContext) {}

// ExitContentsConstraint is called when production contentsConstraint is exited.
func (s *BaseASNListener) ExitContentsConstraint(ctx *ContentsConstraintContext) {}

// EnterComponentPresenceLists is called when production componentPresenceLists is entered.
func (s *BaseASNListener) EnterComponentPresenceLists(ctx *ComponentPresenceListsContext) {}

// ExitComponentPresenceLists is called when production componentPresenceLists is exited.
func (s *BaseASNListener) ExitComponentPresenceLists(ctx *ComponentPresenceListsContext) {}

// EnterComponentPresenceList is called when production componentPresenceList is entered.
func (s *BaseASNListener) EnterComponentPresenceList(ctx *ComponentPresenceListContext) {}

// ExitComponentPresenceList is called when production componentPresenceList is exited.
func (s *BaseASNListener) ExitComponentPresenceList(ctx *ComponentPresenceListContext) {}

// EnterComponentPresence is called when production componentPresence is entered.
func (s *BaseASNListener) EnterComponentPresence(ctx *ComponentPresenceContext) {}

// ExitComponentPresence is called when production componentPresence is exited.
func (s *BaseASNListener) ExitComponentPresence(ctx *ComponentPresenceContext) {}

// EnterSubtypeConstraint is called when production subtypeConstraint is entered.
func (s *BaseASNListener) EnterSubtypeConstraint(ctx *SubtypeConstraintContext) {}

// ExitSubtypeConstraint is called when production subtypeConstraint is exited.
func (s *BaseASNListener) ExitSubtypeConstraint(ctx *SubtypeConstraintContext) {}

// EnterValue is called when production value is entered.
func (s *BaseASNListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseASNListener) ExitValue(ctx *ValueContext) {}

// EnterBuiltinValue is called when production builtinValue is entered.
func (s *BaseASNListener) EnterBuiltinValue(ctx *BuiltinValueContext) {}

// ExitBuiltinValue is called when production builtinValue is exited.
func (s *BaseASNListener) ExitBuiltinValue(ctx *BuiltinValueContext) {}

// EnterObjectIdentifierValue is called when production objectIdentifierValue is entered.
func (s *BaseASNListener) EnterObjectIdentifierValue(ctx *ObjectIdentifierValueContext) {}

// ExitObjectIdentifierValue is called when production objectIdentifierValue is exited.
func (s *BaseASNListener) ExitObjectIdentifierValue(ctx *ObjectIdentifierValueContext) {}

// EnterObjIdComponentsList is called when production objIdComponentsList is entered.
func (s *BaseASNListener) EnterObjIdComponentsList(ctx *ObjIdComponentsListContext) {}

// ExitObjIdComponentsList is called when production objIdComponentsList is exited.
func (s *BaseASNListener) ExitObjIdComponentsList(ctx *ObjIdComponentsListContext) {}

// EnterObjIdComponents is called when production objIdComponents is entered.
func (s *BaseASNListener) EnterObjIdComponents(ctx *ObjIdComponentsContext) {}

// ExitObjIdComponents is called when production objIdComponents is exited.
func (s *BaseASNListener) ExitObjIdComponents(ctx *ObjIdComponentsContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseASNListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseASNListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterChoiceValue is called when production choiceValue is entered.
func (s *BaseASNListener) EnterChoiceValue(ctx *ChoiceValueContext) {}

// ExitChoiceValue is called when production choiceValue is exited.
func (s *BaseASNListener) ExitChoiceValue(ctx *ChoiceValueContext) {}

// EnterEnumeratedValue is called when production enumeratedValue is entered.
func (s *BaseASNListener) EnterEnumeratedValue(ctx *EnumeratedValueContext) {}

// ExitEnumeratedValue is called when production enumeratedValue is exited.
func (s *BaseASNListener) ExitEnumeratedValue(ctx *EnumeratedValueContext) {}

// EnterSignedNumber is called when production signedNumber is entered.
func (s *BaseASNListener) EnterSignedNumber(ctx *SignedNumberContext) {}

// ExitSignedNumber is called when production signedNumber is exited.
func (s *BaseASNListener) ExitSignedNumber(ctx *SignedNumberContext) {}

// EnterChoiceType is called when production choiceType is entered.
func (s *BaseASNListener) EnterChoiceType(ctx *ChoiceTypeContext) {}

// ExitChoiceType is called when production choiceType is exited.
func (s *BaseASNListener) ExitChoiceType(ctx *ChoiceTypeContext) {}

// EnterAlternativeTypeLists is called when production alternativeTypeLists is entered.
func (s *BaseASNListener) EnterAlternativeTypeLists(ctx *AlternativeTypeListsContext) {}

// ExitAlternativeTypeLists is called when production alternativeTypeLists is exited.
func (s *BaseASNListener) ExitAlternativeTypeLists(ctx *AlternativeTypeListsContext) {}

// EnterExtensionAdditionAlternatives is called when production extensionAdditionAlternatives is entered.
func (s *BaseASNListener) EnterExtensionAdditionAlternatives(ctx *ExtensionAdditionAlternativesContext) {
}

// ExitExtensionAdditionAlternatives is called when production extensionAdditionAlternatives is exited.
func (s *BaseASNListener) ExitExtensionAdditionAlternatives(ctx *ExtensionAdditionAlternativesContext) {
}

// EnterExtensionAdditionAlternativesList is called when production extensionAdditionAlternativesList is entered.
func (s *BaseASNListener) EnterExtensionAdditionAlternativesList(ctx *ExtensionAdditionAlternativesListContext) {
}

// ExitExtensionAdditionAlternativesList is called when production extensionAdditionAlternativesList is exited.
func (s *BaseASNListener) ExitExtensionAdditionAlternativesList(ctx *ExtensionAdditionAlternativesListContext) {
}

// EnterExtensionAdditionAlternative is called when production extensionAdditionAlternative is entered.
func (s *BaseASNListener) EnterExtensionAdditionAlternative(ctx *ExtensionAdditionAlternativeContext) {
}

// ExitExtensionAdditionAlternative is called when production extensionAdditionAlternative is exited.
func (s *BaseASNListener) ExitExtensionAdditionAlternative(ctx *ExtensionAdditionAlternativeContext) {}

// EnterExtensionAdditionAlternativesGroup is called when production extensionAdditionAlternativesGroup is entered.
func (s *BaseASNListener) EnterExtensionAdditionAlternativesGroup(ctx *ExtensionAdditionAlternativesGroupContext) {
}

// ExitExtensionAdditionAlternativesGroup is called when production extensionAdditionAlternativesGroup is exited.
func (s *BaseASNListener) ExitExtensionAdditionAlternativesGroup(ctx *ExtensionAdditionAlternativesGroupContext) {
}

// EnterRootAlternativeTypeList is called when production rootAlternativeTypeList is entered.
func (s *BaseASNListener) EnterRootAlternativeTypeList(ctx *RootAlternativeTypeListContext) {}

// ExitRootAlternativeTypeList is called when production rootAlternativeTypeList is exited.
func (s *BaseASNListener) ExitRootAlternativeTypeList(ctx *RootAlternativeTypeListContext) {}

// EnterAlternativeTypeList is called when production alternativeTypeList is entered.
func (s *BaseASNListener) EnterAlternativeTypeList(ctx *AlternativeTypeListContext) {}

// ExitAlternativeTypeList is called when production alternativeTypeList is exited.
func (s *BaseASNListener) ExitAlternativeTypeList(ctx *AlternativeTypeListContext) {}

// EnterNamedType is called when production namedType is entered.
func (s *BaseASNListener) EnterNamedType(ctx *NamedTypeContext) {}

// ExitNamedType is called when production namedType is exited.
func (s *BaseASNListener) ExitNamedType(ctx *NamedTypeContext) {}

// EnterEnumeratedType is called when production enumeratedType is entered.
func (s *BaseASNListener) EnterEnumeratedType(ctx *EnumeratedTypeContext) {}

// ExitEnumeratedType is called when production enumeratedType is exited.
func (s *BaseASNListener) ExitEnumeratedType(ctx *EnumeratedTypeContext) {}

// EnterEnumerations is called when production enumerations is entered.
func (s *BaseASNListener) EnterEnumerations(ctx *EnumerationsContext) {}

// ExitEnumerations is called when production enumerations is exited.
func (s *BaseASNListener) ExitEnumerations(ctx *EnumerationsContext) {}

// EnterRootEnumeration is called when production rootEnumeration is entered.
func (s *BaseASNListener) EnterRootEnumeration(ctx *RootEnumerationContext) {}

// ExitRootEnumeration is called when production rootEnumeration is exited.
func (s *BaseASNListener) ExitRootEnumeration(ctx *RootEnumerationContext) {}

// EnterEnumeration is called when production enumeration is entered.
func (s *BaseASNListener) EnterEnumeration(ctx *EnumerationContext) {}

// ExitEnumeration is called when production enumeration is exited.
func (s *BaseASNListener) ExitEnumeration(ctx *EnumerationContext) {}

// EnterEnumerationItem is called when production enumerationItem is entered.
func (s *BaseASNListener) EnterEnumerationItem(ctx *EnumerationItemContext) {}

// ExitEnumerationItem is called when production enumerationItem is exited.
func (s *BaseASNListener) ExitEnumerationItem(ctx *EnumerationItemContext) {}

// EnterNamedNumber is called when production namedNumber is entered.
func (s *BaseASNListener) EnterNamedNumber(ctx *NamedNumberContext) {}

// ExitNamedNumber is called when production namedNumber is exited.
func (s *BaseASNListener) ExitNamedNumber(ctx *NamedNumberContext) {}

// EnterDefinedValue is called when production definedValue is entered.
func (s *BaseASNListener) EnterDefinedValue(ctx *DefinedValueContext) {}

// ExitDefinedValue is called when production definedValue is exited.
func (s *BaseASNListener) ExitDefinedValue(ctx *DefinedValueContext) {}

// EnterParameterizedValue is called when production parameterizedValue is entered.
func (s *BaseASNListener) EnterParameterizedValue(ctx *ParameterizedValueContext) {}

// ExitParameterizedValue is called when production parameterizedValue is exited.
func (s *BaseASNListener) ExitParameterizedValue(ctx *ParameterizedValueContext) {}

// EnterSimpleDefinedValue is called when production simpleDefinedValue is entered.
func (s *BaseASNListener) EnterSimpleDefinedValue(ctx *SimpleDefinedValueContext) {}

// ExitSimpleDefinedValue is called when production simpleDefinedValue is exited.
func (s *BaseASNListener) ExitSimpleDefinedValue(ctx *SimpleDefinedValueContext) {}

// EnterActualParameterList is called when production actualParameterList is entered.
func (s *BaseASNListener) EnterActualParameterList(ctx *ActualParameterListContext) {}

// ExitActualParameterList is called when production actualParameterList is exited.
func (s *BaseASNListener) ExitActualParameterList(ctx *ActualParameterListContext) {}

// EnterActualParameter is called when production actualParameter is entered.
func (s *BaseASNListener) EnterActualParameter(ctx *ActualParameterContext) {}

// ExitActualParameter is called when production actualParameter is exited.
func (s *BaseASNListener) ExitActualParameter(ctx *ActualParameterContext) {}

// EnterExceptionSpec is called when production exceptionSpec is entered.
func (s *BaseASNListener) EnterExceptionSpec(ctx *ExceptionSpecContext) {}

// ExitExceptionSpec is called when production exceptionSpec is exited.
func (s *BaseASNListener) ExitExceptionSpec(ctx *ExceptionSpecContext) {}

// EnterExceptionIdentification is called when production exceptionIdentification is entered.
func (s *BaseASNListener) EnterExceptionIdentification(ctx *ExceptionIdentificationContext) {}

// ExitExceptionIdentification is called when production exceptionIdentification is exited.
func (s *BaseASNListener) ExitExceptionIdentification(ctx *ExceptionIdentificationContext) {}

// EnterAdditionalEnumeration is called when production additionalEnumeration is entered.
func (s *BaseASNListener) EnterAdditionalEnumeration(ctx *AdditionalEnumerationContext) {}

// ExitAdditionalEnumeration is called when production additionalEnumeration is exited.
func (s *BaseASNListener) ExitAdditionalEnumeration(ctx *AdditionalEnumerationContext) {}

// EnterIntegerType is called when production integerType is entered.
func (s *BaseASNListener) EnterIntegerType(ctx *IntegerTypeContext) {}

// ExitIntegerType is called when production integerType is exited.
func (s *BaseASNListener) ExitIntegerType(ctx *IntegerTypeContext) {}

// EnterNamedNumberList is called when production namedNumberList is entered.
func (s *BaseASNListener) EnterNamedNumberList(ctx *NamedNumberListContext) {}

// ExitNamedNumberList is called when production namedNumberList is exited.
func (s *BaseASNListener) ExitNamedNumberList(ctx *NamedNumberListContext) {}

// EnterObjectidentifiertype is called when production objectidentifiertype is entered.
func (s *BaseASNListener) EnterObjectidentifiertype(ctx *ObjectidentifiertypeContext) {}

// ExitObjectidentifiertype is called when production objectidentifiertype is exited.
func (s *BaseASNListener) ExitObjectidentifiertype(ctx *ObjectidentifiertypeContext) {}

// EnterComponentRelationConstraint is called when production componentRelationConstraint is entered.
func (s *BaseASNListener) EnterComponentRelationConstraint(ctx *ComponentRelationConstraintContext) {}

// ExitComponentRelationConstraint is called when production componentRelationConstraint is exited.
func (s *BaseASNListener) ExitComponentRelationConstraint(ctx *ComponentRelationConstraintContext) {}

// EnterAtNotation is called when production atNotation is entered.
func (s *BaseASNListener) EnterAtNotation(ctx *AtNotationContext) {}

// ExitAtNotation is called when production atNotation is exited.
func (s *BaseASNListener) ExitAtNotation(ctx *AtNotationContext) {}

// EnterLevel is called when production level is entered.
func (s *BaseASNListener) EnterLevel(ctx *LevelContext) {}

// ExitLevel is called when production level is exited.
func (s *BaseASNListener) ExitLevel(ctx *LevelContext) {}

// EnterComponentIdList is called when production componentIdList is entered.
func (s *BaseASNListener) EnterComponentIdList(ctx *ComponentIdListContext) {}

// ExitComponentIdList is called when production componentIdList is exited.
func (s *BaseASNListener) ExitComponentIdList(ctx *ComponentIdListContext) {}

// EnterOctetStringType is called when production octetStringType is entered.
func (s *BaseASNListener) EnterOctetStringType(ctx *OctetStringTypeContext) {}

// ExitOctetStringType is called when production octetStringType is exited.
func (s *BaseASNListener) ExitOctetStringType(ctx *OctetStringTypeContext) {}

// EnterBitStringType is called when production bitStringType is entered.
func (s *BaseASNListener) EnterBitStringType(ctx *BitStringTypeContext) {}

// ExitBitStringType is called when production bitStringType is exited.
func (s *BaseASNListener) ExitBitStringType(ctx *BitStringTypeContext) {}

// EnterNamedBitList is called when production namedBitList is entered.
func (s *BaseASNListener) EnterNamedBitList(ctx *NamedBitListContext) {}

// ExitNamedBitList is called when production namedBitList is exited.
func (s *BaseASNListener) ExitNamedBitList(ctx *NamedBitListContext) {}

// EnterNamedBit is called when production namedBit is entered.
func (s *BaseASNListener) EnterNamedBit(ctx *NamedBitContext) {}

// ExitNamedBit is called when production namedBit is exited.
func (s *BaseASNListener) ExitNamedBit(ctx *NamedBitContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseASNListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseASNListener) ExitBooleanValue(ctx *BooleanValueContext) {}
