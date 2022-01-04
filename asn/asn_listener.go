// Code generated from ASN.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asn // ASN
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ASNListener is a complete listener for a parse tree produced by ASNParser.
type ASNListener interface {
	antlr.ParseTreeListener

	// EnterModules is called when entering the modules production.
	EnterModules(c *ModulesContext)

	// EnterModuleDefinition is called when entering the moduleDefinition production.
	EnterModuleDefinition(c *ModuleDefinitionContext)

	// EnterTagDefault is called when entering the tagDefault production.
	EnterTagDefault(c *TagDefaultContext)

	// EnterExtensionDefault is called when entering the extensionDefault production.
	EnterExtensionDefault(c *ExtensionDefaultContext)

	// EnterModuleBody is called when entering the moduleBody production.
	EnterModuleBody(c *ModuleBodyContext)

	// EnterExports is called when entering the exports production.
	EnterExports(c *ExportsContext)

	// EnterSymbolsExported is called when entering the symbolsExported production.
	EnterSymbolsExported(c *SymbolsExportedContext)

	// EnterImports is called when entering the imports production.
	EnterImports(c *ImportsContext)

	// EnterSymbolsImported is called when entering the symbolsImported production.
	EnterSymbolsImported(c *SymbolsImportedContext)

	// EnterSymbolsFromModuleList is called when entering the symbolsFromModuleList production.
	EnterSymbolsFromModuleList(c *SymbolsFromModuleListContext)

	// EnterSymbolsFromModule is called when entering the symbolsFromModule production.
	EnterSymbolsFromModule(c *SymbolsFromModuleContext)

	// EnterGlobalModuleReference is called when entering the globalModuleReference production.
	EnterGlobalModuleReference(c *GlobalModuleReferenceContext)

	// EnterAssignedIdentifier is called when entering the assignedIdentifier production.
	EnterAssignedIdentifier(c *AssignedIdentifierContext)

	// EnterSymbolList is called when entering the symbolList production.
	EnterSymbolList(c *SymbolListContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterAssignmentList is called when entering the assignmentList production.
	EnterAssignmentList(c *AssignmentListContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterSequenceType is called when entering the sequenceType production.
	EnterSequenceType(c *SequenceTypeContext)

	// EnterExtensionAndException is called when entering the extensionAndException production.
	EnterExtensionAndException(c *ExtensionAndExceptionContext)

	// EnterOptionalExtensionMarker is called when entering the optionalExtensionMarker production.
	EnterOptionalExtensionMarker(c *OptionalExtensionMarkerContext)

	// EnterComponentTypeLists is called when entering the componentTypeLists production.
	EnterComponentTypeLists(c *ComponentTypeListsContext)

	// EnterRootComponentTypeList is called when entering the rootComponentTypeList production.
	EnterRootComponentTypeList(c *RootComponentTypeListContext)

	// EnterComponentTypeList is called when entering the componentTypeList production.
	EnterComponentTypeList(c *ComponentTypeListContext)

	// EnterComponentType is called when entering the componentType production.
	EnterComponentType(c *ComponentTypeContext)

	// EnterExtensionAdditions is called when entering the extensionAdditions production.
	EnterExtensionAdditions(c *ExtensionAdditionsContext)

	// EnterExtensionAdditionList is called when entering the extensionAdditionList production.
	EnterExtensionAdditionList(c *ExtensionAdditionListContext)

	// EnterExtensionAddition is called when entering the extensionAddition production.
	EnterExtensionAddition(c *ExtensionAdditionContext)

	// EnterExtensionAdditionGroup is called when entering the extensionAdditionGroup production.
	EnterExtensionAdditionGroup(c *ExtensionAdditionGroupContext)

	// EnterVersionNumber is called when entering the versionNumber production.
	EnterVersionNumber(c *VersionNumberContext)

	// EnterSequenceOfType is called when entering the sequenceOfType production.
	EnterSequenceOfType(c *SequenceOfTypeContext)

	// EnterSizeConstraint is called when entering the sizeConstraint production.
	EnterSizeConstraint(c *SizeConstraintContext)

	// EnterParameterizedAssignment is called when entering the parameterizedAssignment production.
	EnterParameterizedAssignment(c *ParameterizedAssignmentContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterParamGovernor is called when entering the paramGovernor production.
	EnterParamGovernor(c *ParamGovernorContext)

	// EnterGovernor is called when entering the governor production.
	EnterGovernor(c *GovernorContext)

	// EnterObjectClassAssignment is called when entering the objectClassAssignment production.
	EnterObjectClassAssignment(c *ObjectClassAssignmentContext)

	// EnterObjectClass is called when entering the objectClass production.
	EnterObjectClass(c *ObjectClassContext)

	// EnterDefinedObjectClass is called when entering the definedObjectClass production.
	EnterDefinedObjectClass(c *DefinedObjectClassContext)

	// EnterUsefulObjectClassReference is called when entering the usefulObjectClassReference production.
	EnterUsefulObjectClassReference(c *UsefulObjectClassReferenceContext)

	// EnterExternalObjectClassReference is called when entering the externalObjectClassReference production.
	EnterExternalObjectClassReference(c *ExternalObjectClassReferenceContext)

	// EnterObjectClassDefn is called when entering the objectClassDefn production.
	EnterObjectClassDefn(c *ObjectClassDefnContext)

	// EnterWithSyntaxSpec is called when entering the withSyntaxSpec production.
	EnterWithSyntaxSpec(c *WithSyntaxSpecContext)

	// EnterSyntaxList is called when entering the syntaxList production.
	EnterSyntaxList(c *SyntaxListContext)

	// EnterTokenOrGroupSpec is called when entering the tokenOrGroupSpec production.
	EnterTokenOrGroupSpec(c *TokenOrGroupSpecContext)

	// EnterOptionalGroup is called when entering the optionalGroup production.
	EnterOptionalGroup(c *OptionalGroupContext)

	// EnterRequiredToken is called when entering the requiredToken production.
	EnterRequiredToken(c *RequiredTokenContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterPrimitiveFieldName is called when entering the primitiveFieldName production.
	EnterPrimitiveFieldName(c *PrimitiveFieldNameContext)

	// EnterFieldSpec is called when entering the fieldSpec production.
	EnterFieldSpec(c *FieldSpecContext)

	// EnterTypeFieldSpec is called when entering the typeFieldSpec production.
	EnterTypeFieldSpec(c *TypeFieldSpecContext)

	// EnterTypeOptionalitySpec is called when entering the typeOptionalitySpec production.
	EnterTypeOptionalitySpec(c *TypeOptionalitySpecContext)

	// EnterFixedTypeValueFieldSpec is called when entering the fixedTypeValueFieldSpec production.
	EnterFixedTypeValueFieldSpec(c *FixedTypeValueFieldSpecContext)

	// EnterValueOptionalitySpec is called when entering the valueOptionalitySpec production.
	EnterValueOptionalitySpec(c *ValueOptionalitySpecContext)

	// EnterVariableTypeValueFieldSpec is called when entering the variableTypeValueFieldSpec production.
	EnterVariableTypeValueFieldSpec(c *VariableTypeValueFieldSpecContext)

	// EnterFixedTypeValueSetFieldSpec is called when entering the fixedTypeValueSetFieldSpec production.
	EnterFixedTypeValueSetFieldSpec(c *FixedTypeValueSetFieldSpecContext)

	// EnterValueSetOptionalitySpec is called when entering the valueSetOptionalitySpec production.
	EnterValueSetOptionalitySpec(c *ValueSetOptionalitySpecContext)

	// EnterObject_ is called when entering the object_ production.
	EnterObject_(c *Object_Context)

	// EnterParameterizedObject is called when entering the parameterizedObject production.
	EnterParameterizedObject(c *ParameterizedObjectContext)

	// EnterDefinedObject is called when entering the definedObject production.
	EnterDefinedObject(c *DefinedObjectContext)

	// EnterObjectSet is called when entering the objectSet production.
	EnterObjectSet(c *ObjectSetContext)

	// EnterObjectSetSpec is called when entering the objectSetSpec production.
	EnterObjectSetSpec(c *ObjectSetSpecContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterValueSet is called when entering the valueSet production.
	EnterValueSet(c *ValueSetContext)

	// EnterElementSetSpecs is called when entering the elementSetSpecs production.
	EnterElementSetSpecs(c *ElementSetSpecsContext)

	// EnterRootElementSetSpec is called when entering the rootElementSetSpec production.
	EnterRootElementSetSpec(c *RootElementSetSpecContext)

	// EnterAdditionalElementSetSpec is called when entering the additionalElementSetSpec production.
	EnterAdditionalElementSetSpec(c *AdditionalElementSetSpecContext)

	// EnterElementSetSpec is called when entering the elementSetSpec production.
	EnterElementSetSpec(c *ElementSetSpecContext)

	// EnterUnions is called when entering the unions production.
	EnterUnions(c *UnionsContext)

	// EnterExclusions is called when entering the exclusions production.
	EnterExclusions(c *ExclusionsContext)

	// EnterIntersections is called when entering the intersections production.
	EnterIntersections(c *IntersectionsContext)

	// EnterUnionMark is called when entering the unionMark production.
	EnterUnionMark(c *UnionMarkContext)

	// EnterIntersectionMark is called when entering the intersectionMark production.
	EnterIntersectionMark(c *IntersectionMarkContext)

	// EnterElements is called when entering the elements production.
	EnterElements(c *ElementsContext)

	// EnterObjectSetElements is called when entering the objectSetElements production.
	EnterObjectSetElements(c *ObjectSetElementsContext)

	// EnterIntersectionElements is called when entering the intersectionElements production.
	EnterIntersectionElements(c *IntersectionElementsContext)

	// EnterSubtypeElements is called when entering the subtypeElements production.
	EnterSubtypeElements(c *SubtypeElementsContext)

	// EnterVariableTypeValueSetFieldSpec is called when entering the variableTypeValueSetFieldSpec production.
	EnterVariableTypeValueSetFieldSpec(c *VariableTypeValueSetFieldSpecContext)

	// EnterObjectFieldSpec is called when entering the objectFieldSpec production.
	EnterObjectFieldSpec(c *ObjectFieldSpecContext)

	// EnterObjectOptionalitySpec is called when entering the objectOptionalitySpec production.
	EnterObjectOptionalitySpec(c *ObjectOptionalitySpecContext)

	// EnterObjectSetFieldSpec is called when entering the objectSetFieldSpec production.
	EnterObjectSetFieldSpec(c *ObjectSetFieldSpecContext)

	// EnterObjectSetOptionalitySpec is called when entering the objectSetOptionalitySpec production.
	EnterObjectSetOptionalitySpec(c *ObjectSetOptionalitySpecContext)

	// EnterTypeAssignment is called when entering the typeAssignment production.
	EnterTypeAssignment(c *TypeAssignmentContext)

	// EnterValueAssignment is called when entering the valueAssignment production.
	EnterValueAssignment(c *ValueAssignmentContext)

	// EnterAsnType is called when entering the asnType production.
	EnterAsnType(c *AsnTypeContext)

	// EnterBuiltinType is called when entering the builtinType production.
	EnterBuiltinType(c *BuiltinTypeContext)

	// EnterObjectClassFieldType is called when entering the objectClassFieldType production.
	EnterObjectClassFieldType(c *ObjectClassFieldTypeContext)

	// EnterSetType is called when entering the setType production.
	EnterSetType(c *SetTypeContext)

	// EnterSetOfType is called when entering the setOfType production.
	EnterSetOfType(c *SetOfTypeContext)

	// EnterReferencedType is called when entering the referencedType production.
	EnterReferencedType(c *ReferencedTypeContext)

	// EnterDefinedType is called when entering the definedType production.
	EnterDefinedType(c *DefinedTypeContext)

	// EnterConstraint is called when entering the constraint production.
	EnterConstraint(c *ConstraintContext)

	// EnterConstraintSpec is called when entering the constraintSpec production.
	EnterConstraintSpec(c *ConstraintSpecContext)

	// EnterUserDefinedConstraint is called when entering the userDefinedConstraint production.
	EnterUserDefinedConstraint(c *UserDefinedConstraintContext)

	// EnterGeneralConstraint is called when entering the generalConstraint production.
	EnterGeneralConstraint(c *GeneralConstraintContext)

	// EnterUserDefinedConstraintParameter is called when entering the userDefinedConstraintParameter production.
	EnterUserDefinedConstraintParameter(c *UserDefinedConstraintParameterContext)

	// EnterTableConstraint is called when entering the tableConstraint production.
	EnterTableConstraint(c *TableConstraintContext)

	// EnterSimpleTableConstraint is called when entering the simpleTableConstraint production.
	EnterSimpleTableConstraint(c *SimpleTableConstraintContext)

	// EnterContentsConstraint is called when entering the contentsConstraint production.
	EnterContentsConstraint(c *ContentsConstraintContext)

	// EnterComponentPresenceLists is called when entering the componentPresenceLists production.
	EnterComponentPresenceLists(c *ComponentPresenceListsContext)

	// EnterComponentPresenceList is called when entering the componentPresenceList production.
	EnterComponentPresenceList(c *ComponentPresenceListContext)

	// EnterComponentPresence is called when entering the componentPresence production.
	EnterComponentPresence(c *ComponentPresenceContext)

	// EnterSubtypeConstraint is called when entering the subtypeConstraint production.
	EnterSubtypeConstraint(c *SubtypeConstraintContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterBuiltinValue is called when entering the builtinValue production.
	EnterBuiltinValue(c *BuiltinValueContext)

	// EnterObjectIdentifierValue is called when entering the objectIdentifierValue production.
	EnterObjectIdentifierValue(c *ObjectIdentifierValueContext)

	// EnterObjIdComponentsList is called when entering the objIdComponentsList production.
	EnterObjIdComponentsList(c *ObjIdComponentsListContext)

	// EnterObjIdComponents is called when entering the objIdComponents production.
	EnterObjIdComponents(c *ObjIdComponentsContext)

	// EnterIntegerValue is called when entering the integerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterChoiceValue is called when entering the choiceValue production.
	EnterChoiceValue(c *ChoiceValueContext)

	// EnterEnumeratedValue is called when entering the enumeratedValue production.
	EnterEnumeratedValue(c *EnumeratedValueContext)

	// EnterSignedNumber is called when entering the signedNumber production.
	EnterSignedNumber(c *SignedNumberContext)

	// EnterChoiceType is called when entering the choiceType production.
	EnterChoiceType(c *ChoiceTypeContext)

	// EnterAlternativeTypeLists is called when entering the alternativeTypeLists production.
	EnterAlternativeTypeLists(c *AlternativeTypeListsContext)

	// EnterExtensionAdditionAlternatives is called when entering the extensionAdditionAlternatives production.
	EnterExtensionAdditionAlternatives(c *ExtensionAdditionAlternativesContext)

	// EnterExtensionAdditionAlternativesList is called when entering the extensionAdditionAlternativesList production.
	EnterExtensionAdditionAlternativesList(c *ExtensionAdditionAlternativesListContext)

	// EnterExtensionAdditionAlternative is called when entering the extensionAdditionAlternative production.
	EnterExtensionAdditionAlternative(c *ExtensionAdditionAlternativeContext)

	// EnterExtensionAdditionAlternativesGroup is called when entering the extensionAdditionAlternativesGroup production.
	EnterExtensionAdditionAlternativesGroup(c *ExtensionAdditionAlternativesGroupContext)

	// EnterRootAlternativeTypeList is called when entering the rootAlternativeTypeList production.
	EnterRootAlternativeTypeList(c *RootAlternativeTypeListContext)

	// EnterAlternativeTypeList is called when entering the alternativeTypeList production.
	EnterAlternativeTypeList(c *AlternativeTypeListContext)

	// EnterNamedType is called when entering the namedType production.
	EnterNamedType(c *NamedTypeContext)

	// EnterEnumeratedType is called when entering the enumeratedType production.
	EnterEnumeratedType(c *EnumeratedTypeContext)

	// EnterEnumerations is called when entering the enumerations production.
	EnterEnumerations(c *EnumerationsContext)

	// EnterRootEnumeration is called when entering the rootEnumeration production.
	EnterRootEnumeration(c *RootEnumerationContext)

	// EnterEnumeration is called when entering the enumeration production.
	EnterEnumeration(c *EnumerationContext)

	// EnterEnumerationItem is called when entering the enumerationItem production.
	EnterEnumerationItem(c *EnumerationItemContext)

	// EnterNamedNumber is called when entering the namedNumber production.
	EnterNamedNumber(c *NamedNumberContext)

	// EnterDefinedValue is called when entering the definedValue production.
	EnterDefinedValue(c *DefinedValueContext)

	// EnterParameterizedValue is called when entering the parameterizedValue production.
	EnterParameterizedValue(c *ParameterizedValueContext)

	// EnterSimpleDefinedValue is called when entering the simpleDefinedValue production.
	EnterSimpleDefinedValue(c *SimpleDefinedValueContext)

	// EnterActualParameterList is called when entering the actualParameterList production.
	EnterActualParameterList(c *ActualParameterListContext)

	// EnterActualParameter is called when entering the actualParameter production.
	EnterActualParameter(c *ActualParameterContext)

	// EnterExceptionSpec is called when entering the exceptionSpec production.
	EnterExceptionSpec(c *ExceptionSpecContext)

	// EnterExceptionIdentification is called when entering the exceptionIdentification production.
	EnterExceptionIdentification(c *ExceptionIdentificationContext)

	// EnterAdditionalEnumeration is called when entering the additionalEnumeration production.
	EnterAdditionalEnumeration(c *AdditionalEnumerationContext)

	// EnterIntegerType is called when entering the integerType production.
	EnterIntegerType(c *IntegerTypeContext)

	// EnterNamedNumberList is called when entering the namedNumberList production.
	EnterNamedNumberList(c *NamedNumberListContext)

	// EnterObjectidentifiertype is called when entering the objectidentifiertype production.
	EnterObjectidentifiertype(c *ObjectidentifiertypeContext)

	// EnterComponentRelationConstraint is called when entering the componentRelationConstraint production.
	EnterComponentRelationConstraint(c *ComponentRelationConstraintContext)

	// EnterAtNotation is called when entering the atNotation production.
	EnterAtNotation(c *AtNotationContext)

	// EnterLevel is called when entering the level production.
	EnterLevel(c *LevelContext)

	// EnterComponentIdList is called when entering the componentIdList production.
	EnterComponentIdList(c *ComponentIdListContext)

	// EnterOctetStringType is called when entering the octetStringType production.
	EnterOctetStringType(c *OctetStringTypeContext)

	// EnterBitStringType is called when entering the bitStringType production.
	EnterBitStringType(c *BitStringTypeContext)

	// EnterNamedBitList is called when entering the namedBitList production.
	EnterNamedBitList(c *NamedBitListContext)

	// EnterNamedBit is called when entering the namedBit production.
	EnterNamedBit(c *NamedBitContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// ExitModules is called when exiting the modules production.
	ExitModules(c *ModulesContext)

	// ExitModuleDefinition is called when exiting the moduleDefinition production.
	ExitModuleDefinition(c *ModuleDefinitionContext)

	// ExitTagDefault is called when exiting the tagDefault production.
	ExitTagDefault(c *TagDefaultContext)

	// ExitExtensionDefault is called when exiting the extensionDefault production.
	ExitExtensionDefault(c *ExtensionDefaultContext)

	// ExitModuleBody is called when exiting the moduleBody production.
	ExitModuleBody(c *ModuleBodyContext)

	// ExitExports is called when exiting the exports production.
	ExitExports(c *ExportsContext)

	// ExitSymbolsExported is called when exiting the symbolsExported production.
	ExitSymbolsExported(c *SymbolsExportedContext)

	// ExitImports is called when exiting the imports production.
	ExitImports(c *ImportsContext)

	// ExitSymbolsImported is called when exiting the symbolsImported production.
	ExitSymbolsImported(c *SymbolsImportedContext)

	// ExitSymbolsFromModuleList is called when exiting the symbolsFromModuleList production.
	ExitSymbolsFromModuleList(c *SymbolsFromModuleListContext)

	// ExitSymbolsFromModule is called when exiting the symbolsFromModule production.
	ExitSymbolsFromModule(c *SymbolsFromModuleContext)

	// ExitGlobalModuleReference is called when exiting the globalModuleReference production.
	ExitGlobalModuleReference(c *GlobalModuleReferenceContext)

	// ExitAssignedIdentifier is called when exiting the assignedIdentifier production.
	ExitAssignedIdentifier(c *AssignedIdentifierContext)

	// ExitSymbolList is called when exiting the symbolList production.
	ExitSymbolList(c *SymbolListContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitAssignmentList is called when exiting the assignmentList production.
	ExitAssignmentList(c *AssignmentListContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitSequenceType is called when exiting the sequenceType production.
	ExitSequenceType(c *SequenceTypeContext)

	// ExitExtensionAndException is called when exiting the extensionAndException production.
	ExitExtensionAndException(c *ExtensionAndExceptionContext)

	// ExitOptionalExtensionMarker is called when exiting the optionalExtensionMarker production.
	ExitOptionalExtensionMarker(c *OptionalExtensionMarkerContext)

	// ExitComponentTypeLists is called when exiting the componentTypeLists production.
	ExitComponentTypeLists(c *ComponentTypeListsContext)

	// ExitRootComponentTypeList is called when exiting the rootComponentTypeList production.
	ExitRootComponentTypeList(c *RootComponentTypeListContext)

	// ExitComponentTypeList is called when exiting the componentTypeList production.
	ExitComponentTypeList(c *ComponentTypeListContext)

	// ExitComponentType is called when exiting the componentType production.
	ExitComponentType(c *ComponentTypeContext)

	// ExitExtensionAdditions is called when exiting the extensionAdditions production.
	ExitExtensionAdditions(c *ExtensionAdditionsContext)

	// ExitExtensionAdditionList is called when exiting the extensionAdditionList production.
	ExitExtensionAdditionList(c *ExtensionAdditionListContext)

	// ExitExtensionAddition is called when exiting the extensionAddition production.
	ExitExtensionAddition(c *ExtensionAdditionContext)

	// ExitExtensionAdditionGroup is called when exiting the extensionAdditionGroup production.
	ExitExtensionAdditionGroup(c *ExtensionAdditionGroupContext)

	// ExitVersionNumber is called when exiting the versionNumber production.
	ExitVersionNumber(c *VersionNumberContext)

	// ExitSequenceOfType is called when exiting the sequenceOfType production.
	ExitSequenceOfType(c *SequenceOfTypeContext)

	// ExitSizeConstraint is called when exiting the sizeConstraint production.
	ExitSizeConstraint(c *SizeConstraintContext)

	// ExitParameterizedAssignment is called when exiting the parameterizedAssignment production.
	ExitParameterizedAssignment(c *ParameterizedAssignmentContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitParamGovernor is called when exiting the paramGovernor production.
	ExitParamGovernor(c *ParamGovernorContext)

	// ExitGovernor is called when exiting the governor production.
	ExitGovernor(c *GovernorContext)

	// ExitObjectClassAssignment is called when exiting the objectClassAssignment production.
	ExitObjectClassAssignment(c *ObjectClassAssignmentContext)

	// ExitObjectClass is called when exiting the objectClass production.
	ExitObjectClass(c *ObjectClassContext)

	// ExitDefinedObjectClass is called when exiting the definedObjectClass production.
	ExitDefinedObjectClass(c *DefinedObjectClassContext)

	// ExitUsefulObjectClassReference is called when exiting the usefulObjectClassReference production.
	ExitUsefulObjectClassReference(c *UsefulObjectClassReferenceContext)

	// ExitExternalObjectClassReference is called when exiting the externalObjectClassReference production.
	ExitExternalObjectClassReference(c *ExternalObjectClassReferenceContext)

	// ExitObjectClassDefn is called when exiting the objectClassDefn production.
	ExitObjectClassDefn(c *ObjectClassDefnContext)

	// ExitWithSyntaxSpec is called when exiting the withSyntaxSpec production.
	ExitWithSyntaxSpec(c *WithSyntaxSpecContext)

	// ExitSyntaxList is called when exiting the syntaxList production.
	ExitSyntaxList(c *SyntaxListContext)

	// ExitTokenOrGroupSpec is called when exiting the tokenOrGroupSpec production.
	ExitTokenOrGroupSpec(c *TokenOrGroupSpecContext)

	// ExitOptionalGroup is called when exiting the optionalGroup production.
	ExitOptionalGroup(c *OptionalGroupContext)

	// ExitRequiredToken is called when exiting the requiredToken production.
	ExitRequiredToken(c *RequiredTokenContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitPrimitiveFieldName is called when exiting the primitiveFieldName production.
	ExitPrimitiveFieldName(c *PrimitiveFieldNameContext)

	// ExitFieldSpec is called when exiting the fieldSpec production.
	ExitFieldSpec(c *FieldSpecContext)

	// ExitTypeFieldSpec is called when exiting the typeFieldSpec production.
	ExitTypeFieldSpec(c *TypeFieldSpecContext)

	// ExitTypeOptionalitySpec is called when exiting the typeOptionalitySpec production.
	ExitTypeOptionalitySpec(c *TypeOptionalitySpecContext)

	// ExitFixedTypeValueFieldSpec is called when exiting the fixedTypeValueFieldSpec production.
	ExitFixedTypeValueFieldSpec(c *FixedTypeValueFieldSpecContext)

	// ExitValueOptionalitySpec is called when exiting the valueOptionalitySpec production.
	ExitValueOptionalitySpec(c *ValueOptionalitySpecContext)

	// ExitVariableTypeValueFieldSpec is called when exiting the variableTypeValueFieldSpec production.
	ExitVariableTypeValueFieldSpec(c *VariableTypeValueFieldSpecContext)

	// ExitFixedTypeValueSetFieldSpec is called when exiting the fixedTypeValueSetFieldSpec production.
	ExitFixedTypeValueSetFieldSpec(c *FixedTypeValueSetFieldSpecContext)

	// ExitValueSetOptionalitySpec is called when exiting the valueSetOptionalitySpec production.
	ExitValueSetOptionalitySpec(c *ValueSetOptionalitySpecContext)

	// ExitObject_ is called when exiting the object_ production.
	ExitObject_(c *Object_Context)

	// ExitParameterizedObject is called when exiting the parameterizedObject production.
	ExitParameterizedObject(c *ParameterizedObjectContext)

	// ExitDefinedObject is called when exiting the definedObject production.
	ExitDefinedObject(c *DefinedObjectContext)

	// ExitObjectSet is called when exiting the objectSet production.
	ExitObjectSet(c *ObjectSetContext)

	// ExitObjectSetSpec is called when exiting the objectSetSpec production.
	ExitObjectSetSpec(c *ObjectSetSpecContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitValueSet is called when exiting the valueSet production.
	ExitValueSet(c *ValueSetContext)

	// ExitElementSetSpecs is called when exiting the elementSetSpecs production.
	ExitElementSetSpecs(c *ElementSetSpecsContext)

	// ExitRootElementSetSpec is called when exiting the rootElementSetSpec production.
	ExitRootElementSetSpec(c *RootElementSetSpecContext)

	// ExitAdditionalElementSetSpec is called when exiting the additionalElementSetSpec production.
	ExitAdditionalElementSetSpec(c *AdditionalElementSetSpecContext)

	// ExitElementSetSpec is called when exiting the elementSetSpec production.
	ExitElementSetSpec(c *ElementSetSpecContext)

	// ExitUnions is called when exiting the unions production.
	ExitUnions(c *UnionsContext)

	// ExitExclusions is called when exiting the exclusions production.
	ExitExclusions(c *ExclusionsContext)

	// ExitIntersections is called when exiting the intersections production.
	ExitIntersections(c *IntersectionsContext)

	// ExitUnionMark is called when exiting the unionMark production.
	ExitUnionMark(c *UnionMarkContext)

	// ExitIntersectionMark is called when exiting the intersectionMark production.
	ExitIntersectionMark(c *IntersectionMarkContext)

	// ExitElements is called when exiting the elements production.
	ExitElements(c *ElementsContext)

	// ExitObjectSetElements is called when exiting the objectSetElements production.
	ExitObjectSetElements(c *ObjectSetElementsContext)

	// ExitIntersectionElements is called when exiting the intersectionElements production.
	ExitIntersectionElements(c *IntersectionElementsContext)

	// ExitSubtypeElements is called when exiting the subtypeElements production.
	ExitSubtypeElements(c *SubtypeElementsContext)

	// ExitVariableTypeValueSetFieldSpec is called when exiting the variableTypeValueSetFieldSpec production.
	ExitVariableTypeValueSetFieldSpec(c *VariableTypeValueSetFieldSpecContext)

	// ExitObjectFieldSpec is called when exiting the objectFieldSpec production.
	ExitObjectFieldSpec(c *ObjectFieldSpecContext)

	// ExitObjectOptionalitySpec is called when exiting the objectOptionalitySpec production.
	ExitObjectOptionalitySpec(c *ObjectOptionalitySpecContext)

	// ExitObjectSetFieldSpec is called when exiting the objectSetFieldSpec production.
	ExitObjectSetFieldSpec(c *ObjectSetFieldSpecContext)

	// ExitObjectSetOptionalitySpec is called when exiting the objectSetOptionalitySpec production.
	ExitObjectSetOptionalitySpec(c *ObjectSetOptionalitySpecContext)

	// ExitTypeAssignment is called when exiting the typeAssignment production.
	ExitTypeAssignment(c *TypeAssignmentContext)

	// ExitValueAssignment is called when exiting the valueAssignment production.
	ExitValueAssignment(c *ValueAssignmentContext)

	// ExitAsnType is called when exiting the asnType production.
	ExitAsnType(c *AsnTypeContext)

	// ExitBuiltinType is called when exiting the builtinType production.
	ExitBuiltinType(c *BuiltinTypeContext)

	// ExitObjectClassFieldType is called when exiting the objectClassFieldType production.
	ExitObjectClassFieldType(c *ObjectClassFieldTypeContext)

	// ExitSetType is called when exiting the setType production.
	ExitSetType(c *SetTypeContext)

	// ExitSetOfType is called when exiting the setOfType production.
	ExitSetOfType(c *SetOfTypeContext)

	// ExitReferencedType is called when exiting the referencedType production.
	ExitReferencedType(c *ReferencedTypeContext)

	// ExitDefinedType is called when exiting the definedType production.
	ExitDefinedType(c *DefinedTypeContext)

	// ExitConstraint is called when exiting the constraint production.
	ExitConstraint(c *ConstraintContext)

	// ExitConstraintSpec is called when exiting the constraintSpec production.
	ExitConstraintSpec(c *ConstraintSpecContext)

	// ExitUserDefinedConstraint is called when exiting the userDefinedConstraint production.
	ExitUserDefinedConstraint(c *UserDefinedConstraintContext)

	// ExitGeneralConstraint is called when exiting the generalConstraint production.
	ExitGeneralConstraint(c *GeneralConstraintContext)

	// ExitUserDefinedConstraintParameter is called when exiting the userDefinedConstraintParameter production.
	ExitUserDefinedConstraintParameter(c *UserDefinedConstraintParameterContext)

	// ExitTableConstraint is called when exiting the tableConstraint production.
	ExitTableConstraint(c *TableConstraintContext)

	// ExitSimpleTableConstraint is called when exiting the simpleTableConstraint production.
	ExitSimpleTableConstraint(c *SimpleTableConstraintContext)

	// ExitContentsConstraint is called when exiting the contentsConstraint production.
	ExitContentsConstraint(c *ContentsConstraintContext)

	// ExitComponentPresenceLists is called when exiting the componentPresenceLists production.
	ExitComponentPresenceLists(c *ComponentPresenceListsContext)

	// ExitComponentPresenceList is called when exiting the componentPresenceList production.
	ExitComponentPresenceList(c *ComponentPresenceListContext)

	// ExitComponentPresence is called when exiting the componentPresence production.
	ExitComponentPresence(c *ComponentPresenceContext)

	// ExitSubtypeConstraint is called when exiting the subtypeConstraint production.
	ExitSubtypeConstraint(c *SubtypeConstraintContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitBuiltinValue is called when exiting the builtinValue production.
	ExitBuiltinValue(c *BuiltinValueContext)

	// ExitObjectIdentifierValue is called when exiting the objectIdentifierValue production.
	ExitObjectIdentifierValue(c *ObjectIdentifierValueContext)

	// ExitObjIdComponentsList is called when exiting the objIdComponentsList production.
	ExitObjIdComponentsList(c *ObjIdComponentsListContext)

	// ExitObjIdComponents is called when exiting the objIdComponents production.
	ExitObjIdComponents(c *ObjIdComponentsContext)

	// ExitIntegerValue is called when exiting the integerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitChoiceValue is called when exiting the choiceValue production.
	ExitChoiceValue(c *ChoiceValueContext)

	// ExitEnumeratedValue is called when exiting the enumeratedValue production.
	ExitEnumeratedValue(c *EnumeratedValueContext)

	// ExitSignedNumber is called when exiting the signedNumber production.
	ExitSignedNumber(c *SignedNumberContext)

	// ExitChoiceType is called when exiting the choiceType production.
	ExitChoiceType(c *ChoiceTypeContext)

	// ExitAlternativeTypeLists is called when exiting the alternativeTypeLists production.
	ExitAlternativeTypeLists(c *AlternativeTypeListsContext)

	// ExitExtensionAdditionAlternatives is called when exiting the extensionAdditionAlternatives production.
	ExitExtensionAdditionAlternatives(c *ExtensionAdditionAlternativesContext)

	// ExitExtensionAdditionAlternativesList is called when exiting the extensionAdditionAlternativesList production.
	ExitExtensionAdditionAlternativesList(c *ExtensionAdditionAlternativesListContext)

	// ExitExtensionAdditionAlternative is called when exiting the extensionAdditionAlternative production.
	ExitExtensionAdditionAlternative(c *ExtensionAdditionAlternativeContext)

	// ExitExtensionAdditionAlternativesGroup is called when exiting the extensionAdditionAlternativesGroup production.
	ExitExtensionAdditionAlternativesGroup(c *ExtensionAdditionAlternativesGroupContext)

	// ExitRootAlternativeTypeList is called when exiting the rootAlternativeTypeList production.
	ExitRootAlternativeTypeList(c *RootAlternativeTypeListContext)

	// ExitAlternativeTypeList is called when exiting the alternativeTypeList production.
	ExitAlternativeTypeList(c *AlternativeTypeListContext)

	// ExitNamedType is called when exiting the namedType production.
	ExitNamedType(c *NamedTypeContext)

	// ExitEnumeratedType is called when exiting the enumeratedType production.
	ExitEnumeratedType(c *EnumeratedTypeContext)

	// ExitEnumerations is called when exiting the enumerations production.
	ExitEnumerations(c *EnumerationsContext)

	// ExitRootEnumeration is called when exiting the rootEnumeration production.
	ExitRootEnumeration(c *RootEnumerationContext)

	// ExitEnumeration is called when exiting the enumeration production.
	ExitEnumeration(c *EnumerationContext)

	// ExitEnumerationItem is called when exiting the enumerationItem production.
	ExitEnumerationItem(c *EnumerationItemContext)

	// ExitNamedNumber is called when exiting the namedNumber production.
	ExitNamedNumber(c *NamedNumberContext)

	// ExitDefinedValue is called when exiting the definedValue production.
	ExitDefinedValue(c *DefinedValueContext)

	// ExitParameterizedValue is called when exiting the parameterizedValue production.
	ExitParameterizedValue(c *ParameterizedValueContext)

	// ExitSimpleDefinedValue is called when exiting the simpleDefinedValue production.
	ExitSimpleDefinedValue(c *SimpleDefinedValueContext)

	// ExitActualParameterList is called when exiting the actualParameterList production.
	ExitActualParameterList(c *ActualParameterListContext)

	// ExitActualParameter is called when exiting the actualParameter production.
	ExitActualParameter(c *ActualParameterContext)

	// ExitExceptionSpec is called when exiting the exceptionSpec production.
	ExitExceptionSpec(c *ExceptionSpecContext)

	// ExitExceptionIdentification is called when exiting the exceptionIdentification production.
	ExitExceptionIdentification(c *ExceptionIdentificationContext)

	// ExitAdditionalEnumeration is called when exiting the additionalEnumeration production.
	ExitAdditionalEnumeration(c *AdditionalEnumerationContext)

	// ExitIntegerType is called when exiting the integerType production.
	ExitIntegerType(c *IntegerTypeContext)

	// ExitNamedNumberList is called when exiting the namedNumberList production.
	ExitNamedNumberList(c *NamedNumberListContext)

	// ExitObjectidentifiertype is called when exiting the objectidentifiertype production.
	ExitObjectidentifiertype(c *ObjectidentifiertypeContext)

	// ExitComponentRelationConstraint is called when exiting the componentRelationConstraint production.
	ExitComponentRelationConstraint(c *ComponentRelationConstraintContext)

	// ExitAtNotation is called when exiting the atNotation production.
	ExitAtNotation(c *AtNotationContext)

	// ExitLevel is called when exiting the level production.
	ExitLevel(c *LevelContext)

	// ExitComponentIdList is called when exiting the componentIdList production.
	ExitComponentIdList(c *ComponentIdListContext)

	// ExitOctetStringType is called when exiting the octetStringType production.
	ExitOctetStringType(c *OctetStringTypeContext)

	// ExitBitStringType is called when exiting the bitStringType production.
	ExitBitStringType(c *BitStringTypeContext)

	// ExitNamedBitList is called when exiting the namedBitList production.
	ExitNamedBitList(c *NamedBitListContext)

	// ExitNamedBit is called when exiting the namedBit production.
	ExitNamedBit(c *NamedBitContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)
}
