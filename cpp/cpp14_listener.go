// Generated from CPP14.g4 by ANTLR 4.7.

package cpp // CPP14
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CPP14Listener is a complete listener for a parse tree produced by CPP14Parser.
type CPP14Listener interface {
	antlr.ParseTreeListener

	// EnterTranslationunit is called when entering the translationunit production.
	EnterTranslationunit(c *TranslationunitContext)

	// EnterPrimaryexpression is called when entering the primaryexpression production.
	EnterPrimaryexpression(c *PrimaryexpressionContext)

	// EnterIdexpression is called when entering the idexpression production.
	EnterIdexpression(c *IdexpressionContext)

	// EnterUnqualifiedid is called when entering the unqualifiedid production.
	EnterUnqualifiedid(c *UnqualifiedidContext)

	// EnterQualifiedid is called when entering the qualifiedid production.
	EnterQualifiedid(c *QualifiedidContext)

	// EnterNestednamespecifier is called when entering the nestednamespecifier production.
	EnterNestednamespecifier(c *NestednamespecifierContext)

	// EnterLambdaexpression is called when entering the lambdaexpression production.
	EnterLambdaexpression(c *LambdaexpressionContext)

	// EnterLambdaintroducer is called when entering the lambdaintroducer production.
	EnterLambdaintroducer(c *LambdaintroducerContext)

	// EnterLambdacapture is called when entering the lambdacapture production.
	EnterLambdacapture(c *LambdacaptureContext)

	// EnterCapturedefault is called when entering the capturedefault production.
	EnterCapturedefault(c *CapturedefaultContext)

	// EnterCapturelist is called when entering the capturelist production.
	EnterCapturelist(c *CapturelistContext)

	// EnterCapture is called when entering the capture production.
	EnterCapture(c *CaptureContext)

	// EnterSimplecapture is called when entering the simplecapture production.
	EnterSimplecapture(c *SimplecaptureContext)

	// EnterInitcapture is called when entering the initcapture production.
	EnterInitcapture(c *InitcaptureContext)

	// EnterLambdadeclarator is called when entering the lambdadeclarator production.
	EnterLambdadeclarator(c *LambdadeclaratorContext)

	// EnterPostfixexpression is called when entering the postfixexpression production.
	EnterPostfixexpression(c *PostfixexpressionContext)

	// EnterExpressionlist is called when entering the expressionlist production.
	EnterExpressionlist(c *ExpressionlistContext)

	// EnterPseudodestructorname is called when entering the pseudodestructorname production.
	EnterPseudodestructorname(c *PseudodestructornameContext)

	// EnterUnaryexpression is called when entering the unaryexpression production.
	EnterUnaryexpression(c *UnaryexpressionContext)

	// EnterUnaryoperator is called when entering the unaryoperator production.
	EnterUnaryoperator(c *UnaryoperatorContext)

	// EnterNewexpression is called when entering the newexpression production.
	EnterNewexpression(c *NewexpressionContext)

	// EnterNewplacement is called when entering the newplacement production.
	EnterNewplacement(c *NewplacementContext)

	// EnterNewtypeid is called when entering the newtypeid production.
	EnterNewtypeid(c *NewtypeidContext)

	// EnterNewdeclarator is called when entering the newdeclarator production.
	EnterNewdeclarator(c *NewdeclaratorContext)

	// EnterNoptrnewdeclarator is called when entering the noptrnewdeclarator production.
	EnterNoptrnewdeclarator(c *NoptrnewdeclaratorContext)

	// EnterNewinitializer is called when entering the newinitializer production.
	EnterNewinitializer(c *NewinitializerContext)

	// EnterDeleteexpression is called when entering the deleteexpression production.
	EnterDeleteexpression(c *DeleteexpressionContext)

	// EnterNoexceptexpression is called when entering the noexceptexpression production.
	EnterNoexceptexpression(c *NoexceptexpressionContext)

	// EnterCastexpression is called when entering the castexpression production.
	EnterCastexpression(c *CastexpressionContext)

	// EnterPmexpression is called when entering the pmexpression production.
	EnterPmexpression(c *PmexpressionContext)

	// EnterMultiplicativeexpression is called when entering the multiplicativeexpression production.
	EnterMultiplicativeexpression(c *MultiplicativeexpressionContext)

	// EnterAdditiveexpression is called when entering the additiveexpression production.
	EnterAdditiveexpression(c *AdditiveexpressionContext)

	// EnterShiftexpression is called when entering the shiftexpression production.
	EnterShiftexpression(c *ShiftexpressionContext)

	// EnterRelationalexpression is called when entering the relationalexpression production.
	EnterRelationalexpression(c *RelationalexpressionContext)

	// EnterEqualityexpression is called when entering the equalityexpression production.
	EnterEqualityexpression(c *EqualityexpressionContext)

	// EnterAndexpression is called when entering the andexpression production.
	EnterAndexpression(c *AndexpressionContext)

	// EnterExclusiveorexpression is called when entering the exclusiveorexpression production.
	EnterExclusiveorexpression(c *ExclusiveorexpressionContext)

	// EnterInclusiveorexpression is called when entering the inclusiveorexpression production.
	EnterInclusiveorexpression(c *InclusiveorexpressionContext)

	// EnterLogicalandexpression is called when entering the logicalandexpression production.
	EnterLogicalandexpression(c *LogicalandexpressionContext)

	// EnterLogicalorexpression is called when entering the logicalorexpression production.
	EnterLogicalorexpression(c *LogicalorexpressionContext)

	// EnterConditionalexpression is called when entering the conditionalexpression production.
	EnterConditionalexpression(c *ConditionalexpressionContext)

	// EnterAssignmentexpression is called when entering the assignmentexpression production.
	EnterAssignmentexpression(c *AssignmentexpressionContext)

	// EnterAssignmentoperator is called when entering the assignmentoperator production.
	EnterAssignmentoperator(c *AssignmentoperatorContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConstantexpression is called when entering the constantexpression production.
	EnterConstantexpression(c *ConstantexpressionContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLabeledstatement is called when entering the labeledstatement production.
	EnterLabeledstatement(c *LabeledstatementContext)

	// EnterExpressionstatement is called when entering the expressionstatement production.
	EnterExpressionstatement(c *ExpressionstatementContext)

	// EnterCompoundstatement is called when entering the compoundstatement production.
	EnterCompoundstatement(c *CompoundstatementContext)

	// EnterStatementseq is called when entering the statementseq production.
	EnterStatementseq(c *StatementseqContext)

	// EnterSelectionstatement is called when entering the selectionstatement production.
	EnterSelectionstatement(c *SelectionstatementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterIterationstatement is called when entering the iterationstatement production.
	EnterIterationstatement(c *IterationstatementContext)

	// EnterForinitstatement is called when entering the forinitstatement production.
	EnterForinitstatement(c *ForinitstatementContext)

	// EnterForrangedeclaration is called when entering the forrangedeclaration production.
	EnterForrangedeclaration(c *ForrangedeclarationContext)

	// EnterForrangeinitializer is called when entering the forrangeinitializer production.
	EnterForrangeinitializer(c *ForrangeinitializerContext)

	// EnterJumpstatement is called when entering the jumpstatement production.
	EnterJumpstatement(c *JumpstatementContext)

	// EnterDeclarationstatement is called when entering the declarationstatement production.
	EnterDeclarationstatement(c *DeclarationstatementContext)

	// EnterDeclarationseq is called when entering the declarationseq production.
	EnterDeclarationseq(c *DeclarationseqContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterBlockdeclaration is called when entering the blockdeclaration production.
	EnterBlockdeclaration(c *BlockdeclarationContext)

	// EnterAliasdeclaration is called when entering the aliasdeclaration production.
	EnterAliasdeclaration(c *AliasdeclarationContext)

	// EnterSimpledeclaration is called when entering the simpledeclaration production.
	EnterSimpledeclaration(c *SimpledeclarationContext)

	// EnterStatic_assertdeclaration is called when entering the static_assertdeclaration production.
	EnterStatic_assertdeclaration(c *Static_assertdeclarationContext)

	// EnterEmptydeclaration is called when entering the emptydeclaration production.
	EnterEmptydeclaration(c *EmptydeclarationContext)

	// EnterAttributedeclaration is called when entering the attributedeclaration production.
	EnterAttributedeclaration(c *AttributedeclarationContext)

	// EnterDeclspecifier is called when entering the declspecifier production.
	EnterDeclspecifier(c *DeclspecifierContext)

	// EnterDeclspecifierseq is called when entering the declspecifierseq production.
	EnterDeclspecifierseq(c *DeclspecifierseqContext)

	// EnterStorageclassspecifier is called when entering the storageclassspecifier production.
	EnterStorageclassspecifier(c *StorageclassspecifierContext)

	// EnterFunctionspecifier is called when entering the functionspecifier production.
	EnterFunctionspecifier(c *FunctionspecifierContext)

	// EnterTypedefname is called when entering the typedefname production.
	EnterTypedefname(c *TypedefnameContext)

	// EnterTypespecifier is called when entering the typespecifier production.
	EnterTypespecifier(c *TypespecifierContext)

	// EnterTrailingtypespecifier is called when entering the trailingtypespecifier production.
	EnterTrailingtypespecifier(c *TrailingtypespecifierContext)

	// EnterTypespecifierseq is called when entering the typespecifierseq production.
	EnterTypespecifierseq(c *TypespecifierseqContext)

	// EnterTrailingtypespecifierseq is called when entering the trailingtypespecifierseq production.
	EnterTrailingtypespecifierseq(c *TrailingtypespecifierseqContext)

	// EnterSimpletypespecifier is called when entering the simpletypespecifier production.
	EnterSimpletypespecifier(c *SimpletypespecifierContext)

	// EnterThetypename is called when entering the thetypename production.
	EnterThetypename(c *ThetypenameContext)

	// EnterDecltypespecifier is called when entering the decltypespecifier production.
	EnterDecltypespecifier(c *DecltypespecifierContext)

	// EnterElaboratedtypespecifier is called when entering the elaboratedtypespecifier production.
	EnterElaboratedtypespecifier(c *ElaboratedtypespecifierContext)

	// EnterEnumname is called when entering the enumname production.
	EnterEnumname(c *EnumnameContext)

	// EnterEnumspecifier is called when entering the enumspecifier production.
	EnterEnumspecifier(c *EnumspecifierContext)

	// EnterEnumhead is called when entering the enumhead production.
	EnterEnumhead(c *EnumheadContext)

	// EnterOpaqueenumdeclaration is called when entering the opaqueenumdeclaration production.
	EnterOpaqueenumdeclaration(c *OpaqueenumdeclarationContext)

	// EnterEnumkey is called when entering the enumkey production.
	EnterEnumkey(c *EnumkeyContext)

	// EnterEnumbase is called when entering the enumbase production.
	EnterEnumbase(c *EnumbaseContext)

	// EnterEnumeratorlist is called when entering the enumeratorlist production.
	EnterEnumeratorlist(c *EnumeratorlistContext)

	// EnterEnumeratordefinition is called when entering the enumeratordefinition production.
	EnterEnumeratordefinition(c *EnumeratordefinitionContext)

	// EnterEnumerator is called when entering the enumerator production.
	EnterEnumerator(c *EnumeratorContext)

	// EnterNamespacename is called when entering the namespacename production.
	EnterNamespacename(c *NamespacenameContext)

	// EnterOriginalnamespacename is called when entering the originalnamespacename production.
	EnterOriginalnamespacename(c *OriginalnamespacenameContext)

	// EnterNamespacedefinition is called when entering the namespacedefinition production.
	EnterNamespacedefinition(c *NamespacedefinitionContext)

	// EnterNamednamespacedefinition is called when entering the namednamespacedefinition production.
	EnterNamednamespacedefinition(c *NamednamespacedefinitionContext)

	// EnterOriginalnamespacedefinition is called when entering the originalnamespacedefinition production.
	EnterOriginalnamespacedefinition(c *OriginalnamespacedefinitionContext)

	// EnterExtensionnamespacedefinition is called when entering the extensionnamespacedefinition production.
	EnterExtensionnamespacedefinition(c *ExtensionnamespacedefinitionContext)

	// EnterUnnamednamespacedefinition is called when entering the unnamednamespacedefinition production.
	EnterUnnamednamespacedefinition(c *UnnamednamespacedefinitionContext)

	// EnterNamespacebody is called when entering the namespacebody production.
	EnterNamespacebody(c *NamespacebodyContext)

	// EnterNamespacealias is called when entering the namespacealias production.
	EnterNamespacealias(c *NamespacealiasContext)

	// EnterNamespacealiasdefinition is called when entering the namespacealiasdefinition production.
	EnterNamespacealiasdefinition(c *NamespacealiasdefinitionContext)

	// EnterQualifiednamespacespecifier is called when entering the qualifiednamespacespecifier production.
	EnterQualifiednamespacespecifier(c *QualifiednamespacespecifierContext)

	// EnterUsingdeclaration is called when entering the usingdeclaration production.
	EnterUsingdeclaration(c *UsingdeclarationContext)

	// EnterUsingdirective is called when entering the usingdirective production.
	EnterUsingdirective(c *UsingdirectiveContext)

	// EnterAsmdefinition is called when entering the asmdefinition production.
	EnterAsmdefinition(c *AsmdefinitionContext)

	// EnterLinkagespecification is called when entering the linkagespecification production.
	EnterLinkagespecification(c *LinkagespecificationContext)

	// EnterAttributespecifierseq is called when entering the attributespecifierseq production.
	EnterAttributespecifierseq(c *AttributespecifierseqContext)

	// EnterAttributespecifier is called when entering the attributespecifier production.
	EnterAttributespecifier(c *AttributespecifierContext)

	// EnterAlignmentspecifier is called when entering the alignmentspecifier production.
	EnterAlignmentspecifier(c *AlignmentspecifierContext)

	// EnterAttributelist is called when entering the attributelist production.
	EnterAttributelist(c *AttributelistContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterAttributetoken is called when entering the attributetoken production.
	EnterAttributetoken(c *AttributetokenContext)

	// EnterAttributescopedtoken is called when entering the attributescopedtoken production.
	EnterAttributescopedtoken(c *AttributescopedtokenContext)

	// EnterAttributenamespace is called when entering the attributenamespace production.
	EnterAttributenamespace(c *AttributenamespaceContext)

	// EnterAttributeargumentclause is called when entering the attributeargumentclause production.
	EnterAttributeargumentclause(c *AttributeargumentclauseContext)

	// EnterBalancedtokenseq is called when entering the balancedtokenseq production.
	EnterBalancedtokenseq(c *BalancedtokenseqContext)

	// EnterBalancedtoken is called when entering the balancedtoken production.
	EnterBalancedtoken(c *BalancedtokenContext)

	// EnterInitdeclaratorlist is called when entering the initdeclaratorlist production.
	EnterInitdeclaratorlist(c *InitdeclaratorlistContext)

	// EnterInitdeclarator is called when entering the initdeclarator production.
	EnterInitdeclarator(c *InitdeclaratorContext)

	// EnterDeclarator is called when entering the declarator production.
	EnterDeclarator(c *DeclaratorContext)

	// EnterPtrdeclarator is called when entering the ptrdeclarator production.
	EnterPtrdeclarator(c *PtrdeclaratorContext)

	// EnterNoptrdeclarator is called when entering the noptrdeclarator production.
	EnterNoptrdeclarator(c *NoptrdeclaratorContext)

	// EnterParametersandqualifiers is called when entering the parametersandqualifiers production.
	EnterParametersandqualifiers(c *ParametersandqualifiersContext)

	// EnterTrailingreturntype is called when entering the trailingreturntype production.
	EnterTrailingreturntype(c *TrailingreturntypeContext)

	// EnterPtroperator is called when entering the ptroperator production.
	EnterPtroperator(c *PtroperatorContext)

	// EnterCvqualifierseq is called when entering the cvqualifierseq production.
	EnterCvqualifierseq(c *CvqualifierseqContext)

	// EnterCvqualifier is called when entering the cvqualifier production.
	EnterCvqualifier(c *CvqualifierContext)

	// EnterRefqualifier is called when entering the refqualifier production.
	EnterRefqualifier(c *RefqualifierContext)

	// EnterDeclaratorid is called when entering the declaratorid production.
	EnterDeclaratorid(c *DeclaratoridContext)

	// EnterThetypeid is called when entering the thetypeid production.
	EnterThetypeid(c *ThetypeidContext)

	// EnterAbstractdeclarator is called when entering the abstractdeclarator production.
	EnterAbstractdeclarator(c *AbstractdeclaratorContext)

	// EnterPtrabstractdeclarator is called when entering the ptrabstractdeclarator production.
	EnterPtrabstractdeclarator(c *PtrabstractdeclaratorContext)

	// EnterNoptrabstractdeclarator is called when entering the noptrabstractdeclarator production.
	EnterNoptrabstractdeclarator(c *NoptrabstractdeclaratorContext)

	// EnterAbstractpackdeclarator is called when entering the abstractpackdeclarator production.
	EnterAbstractpackdeclarator(c *AbstractpackdeclaratorContext)

	// EnterNoptrabstractpackdeclarator is called when entering the noptrabstractpackdeclarator production.
	EnterNoptrabstractpackdeclarator(c *NoptrabstractpackdeclaratorContext)

	// EnterParameterdeclarationclause is called when entering the parameterdeclarationclause production.
	EnterParameterdeclarationclause(c *ParameterdeclarationclauseContext)

	// EnterParameterdeclarationlist is called when entering the parameterdeclarationlist production.
	EnterParameterdeclarationlist(c *ParameterdeclarationlistContext)

	// EnterParameterdeclaration is called when entering the parameterdeclaration production.
	EnterParameterdeclaration(c *ParameterdeclarationContext)

	// EnterFunctiondefinition is called when entering the functiondefinition production.
	EnterFunctiondefinition(c *FunctiondefinitionContext)

	// EnterFunctionbody is called when entering the functionbody production.
	EnterFunctionbody(c *FunctionbodyContext)

	// EnterInitializer is called when entering the initializer production.
	EnterInitializer(c *InitializerContext)

	// EnterBraceorequalinitializer is called when entering the braceorequalinitializer production.
	EnterBraceorequalinitializer(c *BraceorequalinitializerContext)

	// EnterInitializerclause is called when entering the initializerclause production.
	EnterInitializerclause(c *InitializerclauseContext)

	// EnterInitializerlist is called when entering the initializerlist production.
	EnterInitializerlist(c *InitializerlistContext)

	// EnterBracedinitlist is called when entering the bracedinitlist production.
	EnterBracedinitlist(c *BracedinitlistContext)

	// EnterClassname is called when entering the classname production.
	EnterClassname(c *ClassnameContext)

	// EnterClassspecifier is called when entering the classspecifier production.
	EnterClassspecifier(c *ClassspecifierContext)

	// EnterClasshead is called when entering the classhead production.
	EnterClasshead(c *ClassheadContext)

	// EnterClassheadname is called when entering the classheadname production.
	EnterClassheadname(c *ClassheadnameContext)

	// EnterClassvirtspecifier is called when entering the classvirtspecifier production.
	EnterClassvirtspecifier(c *ClassvirtspecifierContext)

	// EnterClasskey is called when entering the classkey production.
	EnterClasskey(c *ClasskeyContext)

	// EnterMemberspecification is called when entering the memberspecification production.
	EnterMemberspecification(c *MemberspecificationContext)

	// EnterMemberdeclaration is called when entering the memberdeclaration production.
	EnterMemberdeclaration(c *MemberdeclarationContext)

	// EnterMemberdeclaratorlist is called when entering the memberdeclaratorlist production.
	EnterMemberdeclaratorlist(c *MemberdeclaratorlistContext)

	// EnterMemberdeclarator is called when entering the memberdeclarator production.
	EnterMemberdeclarator(c *MemberdeclaratorContext)

	// EnterVirtspecifierseq is called when entering the virtspecifierseq production.
	EnterVirtspecifierseq(c *VirtspecifierseqContext)

	// EnterVirtspecifier is called when entering the virtspecifier production.
	EnterVirtspecifier(c *VirtspecifierContext)

	// EnterPurespecifier is called when entering the purespecifier production.
	EnterPurespecifier(c *PurespecifierContext)

	// EnterBaseclause is called when entering the baseclause production.
	EnterBaseclause(c *BaseclauseContext)

	// EnterBasespecifierlist is called when entering the basespecifierlist production.
	EnterBasespecifierlist(c *BasespecifierlistContext)

	// EnterBasespecifier is called when entering the basespecifier production.
	EnterBasespecifier(c *BasespecifierContext)

	// EnterClassordecltype is called when entering the classordecltype production.
	EnterClassordecltype(c *ClassordecltypeContext)

	// EnterBasetypespecifier is called when entering the basetypespecifier production.
	EnterBasetypespecifier(c *BasetypespecifierContext)

	// EnterAccessspecifier is called when entering the accessspecifier production.
	EnterAccessspecifier(c *AccessspecifierContext)

	// EnterConversionfunctionid is called when entering the conversionfunctionid production.
	EnterConversionfunctionid(c *ConversionfunctionidContext)

	// EnterConversiontypeid is called when entering the conversiontypeid production.
	EnterConversiontypeid(c *ConversiontypeidContext)

	// EnterConversiondeclarator is called when entering the conversiondeclarator production.
	EnterConversiondeclarator(c *ConversiondeclaratorContext)

	// EnterCtorinitializer is called when entering the ctorinitializer production.
	EnterCtorinitializer(c *CtorinitializerContext)

	// EnterMeminitializerlist is called when entering the meminitializerlist production.
	EnterMeminitializerlist(c *MeminitializerlistContext)

	// EnterMeminitializer is called when entering the meminitializer production.
	EnterMeminitializer(c *MeminitializerContext)

	// EnterMeminitializerid is called when entering the meminitializerid production.
	EnterMeminitializerid(c *MeminitializeridContext)

	// EnterOperatorfunctionid is called when entering the operatorfunctionid production.
	EnterOperatorfunctionid(c *OperatorfunctionidContext)

	// EnterLiteraloperatorid is called when entering the literaloperatorid production.
	EnterLiteraloperatorid(c *LiteraloperatoridContext)

	// EnterTemplatedeclaration is called when entering the templatedeclaration production.
	EnterTemplatedeclaration(c *TemplatedeclarationContext)

	// EnterTemplateparameterlist is called when entering the templateparameterlist production.
	EnterTemplateparameterlist(c *TemplateparameterlistContext)

	// EnterTemplateparameter is called when entering the templateparameter production.
	EnterTemplateparameter(c *TemplateparameterContext)

	// EnterTypeparameter is called when entering the typeparameter production.
	EnterTypeparameter(c *TypeparameterContext)

	// EnterSimpletemplateid is called when entering the simpletemplateid production.
	EnterSimpletemplateid(c *SimpletemplateidContext)

	// EnterTemplateid is called when entering the templateid production.
	EnterTemplateid(c *TemplateidContext)

	// EnterTemplatename is called when entering the templatename production.
	EnterTemplatename(c *TemplatenameContext)

	// EnterTemplateargumentlist is called when entering the templateargumentlist production.
	EnterTemplateargumentlist(c *TemplateargumentlistContext)

	// EnterTemplateargument is called when entering the templateargument production.
	EnterTemplateargument(c *TemplateargumentContext)

	// EnterTypenamespecifier is called when entering the typenamespecifier production.
	EnterTypenamespecifier(c *TypenamespecifierContext)

	// EnterExplicitinstantiation is called when entering the explicitinstantiation production.
	EnterExplicitinstantiation(c *ExplicitinstantiationContext)

	// EnterExplicitspecialization is called when entering the explicitspecialization production.
	EnterExplicitspecialization(c *ExplicitspecializationContext)

	// EnterTryblock is called when entering the tryblock production.
	EnterTryblock(c *TryblockContext)

	// EnterFunctiontryblock is called when entering the functiontryblock production.
	EnterFunctiontryblock(c *FunctiontryblockContext)

	// EnterHandlerseq is called when entering the handlerseq production.
	EnterHandlerseq(c *HandlerseqContext)

	// EnterHandler is called when entering the handler production.
	EnterHandler(c *HandlerContext)

	// EnterExceptiondeclaration is called when entering the exceptiondeclaration production.
	EnterExceptiondeclaration(c *ExceptiondeclarationContext)

	// EnterThrowexpression is called when entering the throwexpression production.
	EnterThrowexpression(c *ThrowexpressionContext)

	// EnterExceptionspecification is called when entering the exceptionspecification production.
	EnterExceptionspecification(c *ExceptionspecificationContext)

	// EnterDynamicexceptionspecification is called when entering the dynamicexceptionspecification production.
	EnterDynamicexceptionspecification(c *DynamicexceptionspecificationContext)

	// EnterTypeidlist is called when entering the typeidlist production.
	EnterTypeidlist(c *TypeidlistContext)

	// EnterNoexceptspecification is called when entering the noexceptspecification production.
	EnterNoexceptspecification(c *NoexceptspecificationContext)

	// EnterRightShift is called when entering the rightShift production.
	EnterRightShift(c *RightShiftContext)

	// EnterRightShiftAssign is called when entering the rightShiftAssign production.
	EnterRightShiftAssign(c *RightShiftAssignContext)

	// EnterTheoperator is called when entering the theoperator production.
	EnterTheoperator(c *TheoperatorContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterBooleanliteral is called when entering the booleanliteral production.
	EnterBooleanliteral(c *BooleanliteralContext)

	// EnterPointerliteral is called when entering the pointerliteral production.
	EnterPointerliteral(c *PointerliteralContext)

	// EnterUserdefinedliteral is called when entering the userdefinedliteral production.
	EnterUserdefinedliteral(c *UserdefinedliteralContext)

	// ExitTranslationunit is called when exiting the translationunit production.
	ExitTranslationunit(c *TranslationunitContext)

	// ExitPrimaryexpression is called when exiting the primaryexpression production.
	ExitPrimaryexpression(c *PrimaryexpressionContext)

	// ExitIdexpression is called when exiting the idexpression production.
	ExitIdexpression(c *IdexpressionContext)

	// ExitUnqualifiedid is called when exiting the unqualifiedid production.
	ExitUnqualifiedid(c *UnqualifiedidContext)

	// ExitQualifiedid is called when exiting the qualifiedid production.
	ExitQualifiedid(c *QualifiedidContext)

	// ExitNestednamespecifier is called when exiting the nestednamespecifier production.
	ExitNestednamespecifier(c *NestednamespecifierContext)

	// ExitLambdaexpression is called when exiting the lambdaexpression production.
	ExitLambdaexpression(c *LambdaexpressionContext)

	// ExitLambdaintroducer is called when exiting the lambdaintroducer production.
	ExitLambdaintroducer(c *LambdaintroducerContext)

	// ExitLambdacapture is called when exiting the lambdacapture production.
	ExitLambdacapture(c *LambdacaptureContext)

	// ExitCapturedefault is called when exiting the capturedefault production.
	ExitCapturedefault(c *CapturedefaultContext)

	// ExitCapturelist is called when exiting the capturelist production.
	ExitCapturelist(c *CapturelistContext)

	// ExitCapture is called when exiting the capture production.
	ExitCapture(c *CaptureContext)

	// ExitSimplecapture is called when exiting the simplecapture production.
	ExitSimplecapture(c *SimplecaptureContext)

	// ExitInitcapture is called when exiting the initcapture production.
	ExitInitcapture(c *InitcaptureContext)

	// ExitLambdadeclarator is called when exiting the lambdadeclarator production.
	ExitLambdadeclarator(c *LambdadeclaratorContext)

	// ExitPostfixexpression is called when exiting the postfixexpression production.
	ExitPostfixexpression(c *PostfixexpressionContext)

	// ExitExpressionlist is called when exiting the expressionlist production.
	ExitExpressionlist(c *ExpressionlistContext)

	// ExitPseudodestructorname is called when exiting the pseudodestructorname production.
	ExitPseudodestructorname(c *PseudodestructornameContext)

	// ExitUnaryexpression is called when exiting the unaryexpression production.
	ExitUnaryexpression(c *UnaryexpressionContext)

	// ExitUnaryoperator is called when exiting the unaryoperator production.
	ExitUnaryoperator(c *UnaryoperatorContext)

	// ExitNewexpression is called when exiting the newexpression production.
	ExitNewexpression(c *NewexpressionContext)

	// ExitNewplacement is called when exiting the newplacement production.
	ExitNewplacement(c *NewplacementContext)

	// ExitNewtypeid is called when exiting the newtypeid production.
	ExitNewtypeid(c *NewtypeidContext)

	// ExitNewdeclarator is called when exiting the newdeclarator production.
	ExitNewdeclarator(c *NewdeclaratorContext)

	// ExitNoptrnewdeclarator is called when exiting the noptrnewdeclarator production.
	ExitNoptrnewdeclarator(c *NoptrnewdeclaratorContext)

	// ExitNewinitializer is called when exiting the newinitializer production.
	ExitNewinitializer(c *NewinitializerContext)

	// ExitDeleteexpression is called when exiting the deleteexpression production.
	ExitDeleteexpression(c *DeleteexpressionContext)

	// ExitNoexceptexpression is called when exiting the noexceptexpression production.
	ExitNoexceptexpression(c *NoexceptexpressionContext)

	// ExitCastexpression is called when exiting the castexpression production.
	ExitCastexpression(c *CastexpressionContext)

	// ExitPmexpression is called when exiting the pmexpression production.
	ExitPmexpression(c *PmexpressionContext)

	// ExitMultiplicativeexpression is called when exiting the multiplicativeexpression production.
	ExitMultiplicativeexpression(c *MultiplicativeexpressionContext)

	// ExitAdditiveexpression is called when exiting the additiveexpression production.
	ExitAdditiveexpression(c *AdditiveexpressionContext)

	// ExitShiftexpression is called when exiting the shiftexpression production.
	ExitShiftexpression(c *ShiftexpressionContext)

	// ExitRelationalexpression is called when exiting the relationalexpression production.
	ExitRelationalexpression(c *RelationalexpressionContext)

	// ExitEqualityexpression is called when exiting the equalityexpression production.
	ExitEqualityexpression(c *EqualityexpressionContext)

	// ExitAndexpression is called when exiting the andexpression production.
	ExitAndexpression(c *AndexpressionContext)

	// ExitExclusiveorexpression is called when exiting the exclusiveorexpression production.
	ExitExclusiveorexpression(c *ExclusiveorexpressionContext)

	// ExitInclusiveorexpression is called when exiting the inclusiveorexpression production.
	ExitInclusiveorexpression(c *InclusiveorexpressionContext)

	// ExitLogicalandexpression is called when exiting the logicalandexpression production.
	ExitLogicalandexpression(c *LogicalandexpressionContext)

	// ExitLogicalorexpression is called when exiting the logicalorexpression production.
	ExitLogicalorexpression(c *LogicalorexpressionContext)

	// ExitConditionalexpression is called when exiting the conditionalexpression production.
	ExitConditionalexpression(c *ConditionalexpressionContext)

	// ExitAssignmentexpression is called when exiting the assignmentexpression production.
	ExitAssignmentexpression(c *AssignmentexpressionContext)

	// ExitAssignmentoperator is called when exiting the assignmentoperator production.
	ExitAssignmentoperator(c *AssignmentoperatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConstantexpression is called when exiting the constantexpression production.
	ExitConstantexpression(c *ConstantexpressionContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLabeledstatement is called when exiting the labeledstatement production.
	ExitLabeledstatement(c *LabeledstatementContext)

	// ExitExpressionstatement is called when exiting the expressionstatement production.
	ExitExpressionstatement(c *ExpressionstatementContext)

	// ExitCompoundstatement is called when exiting the compoundstatement production.
	ExitCompoundstatement(c *CompoundstatementContext)

	// ExitStatementseq is called when exiting the statementseq production.
	ExitStatementseq(c *StatementseqContext)

	// ExitSelectionstatement is called when exiting the selectionstatement production.
	ExitSelectionstatement(c *SelectionstatementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitIterationstatement is called when exiting the iterationstatement production.
	ExitIterationstatement(c *IterationstatementContext)

	// ExitForinitstatement is called when exiting the forinitstatement production.
	ExitForinitstatement(c *ForinitstatementContext)

	// ExitForrangedeclaration is called when exiting the forrangedeclaration production.
	ExitForrangedeclaration(c *ForrangedeclarationContext)

	// ExitForrangeinitializer is called when exiting the forrangeinitializer production.
	ExitForrangeinitializer(c *ForrangeinitializerContext)

	// ExitJumpstatement is called when exiting the jumpstatement production.
	ExitJumpstatement(c *JumpstatementContext)

	// ExitDeclarationstatement is called when exiting the declarationstatement production.
	ExitDeclarationstatement(c *DeclarationstatementContext)

	// ExitDeclarationseq is called when exiting the declarationseq production.
	ExitDeclarationseq(c *DeclarationseqContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitBlockdeclaration is called when exiting the blockdeclaration production.
	ExitBlockdeclaration(c *BlockdeclarationContext)

	// ExitAliasdeclaration is called when exiting the aliasdeclaration production.
	ExitAliasdeclaration(c *AliasdeclarationContext)

	// ExitSimpledeclaration is called when exiting the simpledeclaration production.
	ExitSimpledeclaration(c *SimpledeclarationContext)

	// ExitStatic_assertdeclaration is called when exiting the static_assertdeclaration production.
	ExitStatic_assertdeclaration(c *Static_assertdeclarationContext)

	// ExitEmptydeclaration is called when exiting the emptydeclaration production.
	ExitEmptydeclaration(c *EmptydeclarationContext)

	// ExitAttributedeclaration is called when exiting the attributedeclaration production.
	ExitAttributedeclaration(c *AttributedeclarationContext)

	// ExitDeclspecifier is called when exiting the declspecifier production.
	ExitDeclspecifier(c *DeclspecifierContext)

	// ExitDeclspecifierseq is called when exiting the declspecifierseq production.
	ExitDeclspecifierseq(c *DeclspecifierseqContext)

	// ExitStorageclassspecifier is called when exiting the storageclassspecifier production.
	ExitStorageclassspecifier(c *StorageclassspecifierContext)

	// ExitFunctionspecifier is called when exiting the functionspecifier production.
	ExitFunctionspecifier(c *FunctionspecifierContext)

	// ExitTypedefname is called when exiting the typedefname production.
	ExitTypedefname(c *TypedefnameContext)

	// ExitTypespecifier is called when exiting the typespecifier production.
	ExitTypespecifier(c *TypespecifierContext)

	// ExitTrailingtypespecifier is called when exiting the trailingtypespecifier production.
	ExitTrailingtypespecifier(c *TrailingtypespecifierContext)

	// ExitTypespecifierseq is called when exiting the typespecifierseq production.
	ExitTypespecifierseq(c *TypespecifierseqContext)

	// ExitTrailingtypespecifierseq is called when exiting the trailingtypespecifierseq production.
	ExitTrailingtypespecifierseq(c *TrailingtypespecifierseqContext)

	// ExitSimpletypespecifier is called when exiting the simpletypespecifier production.
	ExitSimpletypespecifier(c *SimpletypespecifierContext)

	// ExitThetypename is called when exiting the thetypename production.
	ExitThetypename(c *ThetypenameContext)

	// ExitDecltypespecifier is called when exiting the decltypespecifier production.
	ExitDecltypespecifier(c *DecltypespecifierContext)

	// ExitElaboratedtypespecifier is called when exiting the elaboratedtypespecifier production.
	ExitElaboratedtypespecifier(c *ElaboratedtypespecifierContext)

	// ExitEnumname is called when exiting the enumname production.
	ExitEnumname(c *EnumnameContext)

	// ExitEnumspecifier is called when exiting the enumspecifier production.
	ExitEnumspecifier(c *EnumspecifierContext)

	// ExitEnumhead is called when exiting the enumhead production.
	ExitEnumhead(c *EnumheadContext)

	// ExitOpaqueenumdeclaration is called when exiting the opaqueenumdeclaration production.
	ExitOpaqueenumdeclaration(c *OpaqueenumdeclarationContext)

	// ExitEnumkey is called when exiting the enumkey production.
	ExitEnumkey(c *EnumkeyContext)

	// ExitEnumbase is called when exiting the enumbase production.
	ExitEnumbase(c *EnumbaseContext)

	// ExitEnumeratorlist is called when exiting the enumeratorlist production.
	ExitEnumeratorlist(c *EnumeratorlistContext)

	// ExitEnumeratordefinition is called when exiting the enumeratordefinition production.
	ExitEnumeratordefinition(c *EnumeratordefinitionContext)

	// ExitEnumerator is called when exiting the enumerator production.
	ExitEnumerator(c *EnumeratorContext)

	// ExitNamespacename is called when exiting the namespacename production.
	ExitNamespacename(c *NamespacenameContext)

	// ExitOriginalnamespacename is called when exiting the originalnamespacename production.
	ExitOriginalnamespacename(c *OriginalnamespacenameContext)

	// ExitNamespacedefinition is called when exiting the namespacedefinition production.
	ExitNamespacedefinition(c *NamespacedefinitionContext)

	// ExitNamednamespacedefinition is called when exiting the namednamespacedefinition production.
	ExitNamednamespacedefinition(c *NamednamespacedefinitionContext)

	// ExitOriginalnamespacedefinition is called when exiting the originalnamespacedefinition production.
	ExitOriginalnamespacedefinition(c *OriginalnamespacedefinitionContext)

	// ExitExtensionnamespacedefinition is called when exiting the extensionnamespacedefinition production.
	ExitExtensionnamespacedefinition(c *ExtensionnamespacedefinitionContext)

	// ExitUnnamednamespacedefinition is called when exiting the unnamednamespacedefinition production.
	ExitUnnamednamespacedefinition(c *UnnamednamespacedefinitionContext)

	// ExitNamespacebody is called when exiting the namespacebody production.
	ExitNamespacebody(c *NamespacebodyContext)

	// ExitNamespacealias is called when exiting the namespacealias production.
	ExitNamespacealias(c *NamespacealiasContext)

	// ExitNamespacealiasdefinition is called when exiting the namespacealiasdefinition production.
	ExitNamespacealiasdefinition(c *NamespacealiasdefinitionContext)

	// ExitQualifiednamespacespecifier is called when exiting the qualifiednamespacespecifier production.
	ExitQualifiednamespacespecifier(c *QualifiednamespacespecifierContext)

	// ExitUsingdeclaration is called when exiting the usingdeclaration production.
	ExitUsingdeclaration(c *UsingdeclarationContext)

	// ExitUsingdirective is called when exiting the usingdirective production.
	ExitUsingdirective(c *UsingdirectiveContext)

	// ExitAsmdefinition is called when exiting the asmdefinition production.
	ExitAsmdefinition(c *AsmdefinitionContext)

	// ExitLinkagespecification is called when exiting the linkagespecification production.
	ExitLinkagespecification(c *LinkagespecificationContext)

	// ExitAttributespecifierseq is called when exiting the attributespecifierseq production.
	ExitAttributespecifierseq(c *AttributespecifierseqContext)

	// ExitAttributespecifier is called when exiting the attributespecifier production.
	ExitAttributespecifier(c *AttributespecifierContext)

	// ExitAlignmentspecifier is called when exiting the alignmentspecifier production.
	ExitAlignmentspecifier(c *AlignmentspecifierContext)

	// ExitAttributelist is called when exiting the attributelist production.
	ExitAttributelist(c *AttributelistContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitAttributetoken is called when exiting the attributetoken production.
	ExitAttributetoken(c *AttributetokenContext)

	// ExitAttributescopedtoken is called when exiting the attributescopedtoken production.
	ExitAttributescopedtoken(c *AttributescopedtokenContext)

	// ExitAttributenamespace is called when exiting the attributenamespace production.
	ExitAttributenamespace(c *AttributenamespaceContext)

	// ExitAttributeargumentclause is called when exiting the attributeargumentclause production.
	ExitAttributeargumentclause(c *AttributeargumentclauseContext)

	// ExitBalancedtokenseq is called when exiting the balancedtokenseq production.
	ExitBalancedtokenseq(c *BalancedtokenseqContext)

	// ExitBalancedtoken is called when exiting the balancedtoken production.
	ExitBalancedtoken(c *BalancedtokenContext)

	// ExitInitdeclaratorlist is called when exiting the initdeclaratorlist production.
	ExitInitdeclaratorlist(c *InitdeclaratorlistContext)

	// ExitInitdeclarator is called when exiting the initdeclarator production.
	ExitInitdeclarator(c *InitdeclaratorContext)

	// ExitDeclarator is called when exiting the declarator production.
	ExitDeclarator(c *DeclaratorContext)

	// ExitPtrdeclarator is called when exiting the ptrdeclarator production.
	ExitPtrdeclarator(c *PtrdeclaratorContext)

	// ExitNoptrdeclarator is called when exiting the noptrdeclarator production.
	ExitNoptrdeclarator(c *NoptrdeclaratorContext)

	// ExitParametersandqualifiers is called when exiting the parametersandqualifiers production.
	ExitParametersandqualifiers(c *ParametersandqualifiersContext)

	// ExitTrailingreturntype is called when exiting the trailingreturntype production.
	ExitTrailingreturntype(c *TrailingreturntypeContext)

	// ExitPtroperator is called when exiting the ptroperator production.
	ExitPtroperator(c *PtroperatorContext)

	// ExitCvqualifierseq is called when exiting the cvqualifierseq production.
	ExitCvqualifierseq(c *CvqualifierseqContext)

	// ExitCvqualifier is called when exiting the cvqualifier production.
	ExitCvqualifier(c *CvqualifierContext)

	// ExitRefqualifier is called when exiting the refqualifier production.
	ExitRefqualifier(c *RefqualifierContext)

	// ExitDeclaratorid is called when exiting the declaratorid production.
	ExitDeclaratorid(c *DeclaratoridContext)

	// ExitThetypeid is called when exiting the thetypeid production.
	ExitThetypeid(c *ThetypeidContext)

	// ExitAbstractdeclarator is called when exiting the abstractdeclarator production.
	ExitAbstractdeclarator(c *AbstractdeclaratorContext)

	// ExitPtrabstractdeclarator is called when exiting the ptrabstractdeclarator production.
	ExitPtrabstractdeclarator(c *PtrabstractdeclaratorContext)

	// ExitNoptrabstractdeclarator is called when exiting the noptrabstractdeclarator production.
	ExitNoptrabstractdeclarator(c *NoptrabstractdeclaratorContext)

	// ExitAbstractpackdeclarator is called when exiting the abstractpackdeclarator production.
	ExitAbstractpackdeclarator(c *AbstractpackdeclaratorContext)

	// ExitNoptrabstractpackdeclarator is called when exiting the noptrabstractpackdeclarator production.
	ExitNoptrabstractpackdeclarator(c *NoptrabstractpackdeclaratorContext)

	// ExitParameterdeclarationclause is called when exiting the parameterdeclarationclause production.
	ExitParameterdeclarationclause(c *ParameterdeclarationclauseContext)

	// ExitParameterdeclarationlist is called when exiting the parameterdeclarationlist production.
	ExitParameterdeclarationlist(c *ParameterdeclarationlistContext)

	// ExitParameterdeclaration is called when exiting the parameterdeclaration production.
	ExitParameterdeclaration(c *ParameterdeclarationContext)

	// ExitFunctiondefinition is called when exiting the functiondefinition production.
	ExitFunctiondefinition(c *FunctiondefinitionContext)

	// ExitFunctionbody is called when exiting the functionbody production.
	ExitFunctionbody(c *FunctionbodyContext)

	// ExitInitializer is called when exiting the initializer production.
	ExitInitializer(c *InitializerContext)

	// ExitBraceorequalinitializer is called when exiting the braceorequalinitializer production.
	ExitBraceorequalinitializer(c *BraceorequalinitializerContext)

	// ExitInitializerclause is called when exiting the initializerclause production.
	ExitInitializerclause(c *InitializerclauseContext)

	// ExitInitializerlist is called when exiting the initializerlist production.
	ExitInitializerlist(c *InitializerlistContext)

	// ExitBracedinitlist is called when exiting the bracedinitlist production.
	ExitBracedinitlist(c *BracedinitlistContext)

	// ExitClassname is called when exiting the classname production.
	ExitClassname(c *ClassnameContext)

	// ExitClassspecifier is called when exiting the classspecifier production.
	ExitClassspecifier(c *ClassspecifierContext)

	// ExitClasshead is called when exiting the classhead production.
	ExitClasshead(c *ClassheadContext)

	// ExitClassheadname is called when exiting the classheadname production.
	ExitClassheadname(c *ClassheadnameContext)

	// ExitClassvirtspecifier is called when exiting the classvirtspecifier production.
	ExitClassvirtspecifier(c *ClassvirtspecifierContext)

	// ExitClasskey is called when exiting the classkey production.
	ExitClasskey(c *ClasskeyContext)

	// ExitMemberspecification is called when exiting the memberspecification production.
	ExitMemberspecification(c *MemberspecificationContext)

	// ExitMemberdeclaration is called when exiting the memberdeclaration production.
	ExitMemberdeclaration(c *MemberdeclarationContext)

	// ExitMemberdeclaratorlist is called when exiting the memberdeclaratorlist production.
	ExitMemberdeclaratorlist(c *MemberdeclaratorlistContext)

	// ExitMemberdeclarator is called when exiting the memberdeclarator production.
	ExitMemberdeclarator(c *MemberdeclaratorContext)

	// ExitVirtspecifierseq is called when exiting the virtspecifierseq production.
	ExitVirtspecifierseq(c *VirtspecifierseqContext)

	// ExitVirtspecifier is called when exiting the virtspecifier production.
	ExitVirtspecifier(c *VirtspecifierContext)

	// ExitPurespecifier is called when exiting the purespecifier production.
	ExitPurespecifier(c *PurespecifierContext)

	// ExitBaseclause is called when exiting the baseclause production.
	ExitBaseclause(c *BaseclauseContext)

	// ExitBasespecifierlist is called when exiting the basespecifierlist production.
	ExitBasespecifierlist(c *BasespecifierlistContext)

	// ExitBasespecifier is called when exiting the basespecifier production.
	ExitBasespecifier(c *BasespecifierContext)

	// ExitClassordecltype is called when exiting the classordecltype production.
	ExitClassordecltype(c *ClassordecltypeContext)

	// ExitBasetypespecifier is called when exiting the basetypespecifier production.
	ExitBasetypespecifier(c *BasetypespecifierContext)

	// ExitAccessspecifier is called when exiting the accessspecifier production.
	ExitAccessspecifier(c *AccessspecifierContext)

	// ExitConversionfunctionid is called when exiting the conversionfunctionid production.
	ExitConversionfunctionid(c *ConversionfunctionidContext)

	// ExitConversiontypeid is called when exiting the conversiontypeid production.
	ExitConversiontypeid(c *ConversiontypeidContext)

	// ExitConversiondeclarator is called when exiting the conversiondeclarator production.
	ExitConversiondeclarator(c *ConversiondeclaratorContext)

	// ExitCtorinitializer is called when exiting the ctorinitializer production.
	ExitCtorinitializer(c *CtorinitializerContext)

	// ExitMeminitializerlist is called when exiting the meminitializerlist production.
	ExitMeminitializerlist(c *MeminitializerlistContext)

	// ExitMeminitializer is called when exiting the meminitializer production.
	ExitMeminitializer(c *MeminitializerContext)

	// ExitMeminitializerid is called when exiting the meminitializerid production.
	ExitMeminitializerid(c *MeminitializeridContext)

	// ExitOperatorfunctionid is called when exiting the operatorfunctionid production.
	ExitOperatorfunctionid(c *OperatorfunctionidContext)

	// ExitLiteraloperatorid is called when exiting the literaloperatorid production.
	ExitLiteraloperatorid(c *LiteraloperatoridContext)

	// ExitTemplatedeclaration is called when exiting the templatedeclaration production.
	ExitTemplatedeclaration(c *TemplatedeclarationContext)

	// ExitTemplateparameterlist is called when exiting the templateparameterlist production.
	ExitTemplateparameterlist(c *TemplateparameterlistContext)

	// ExitTemplateparameter is called when exiting the templateparameter production.
	ExitTemplateparameter(c *TemplateparameterContext)

	// ExitTypeparameter is called when exiting the typeparameter production.
	ExitTypeparameter(c *TypeparameterContext)

	// ExitSimpletemplateid is called when exiting the simpletemplateid production.
	ExitSimpletemplateid(c *SimpletemplateidContext)

	// ExitTemplateid is called when exiting the templateid production.
	ExitTemplateid(c *TemplateidContext)

	// ExitTemplatename is called when exiting the templatename production.
	ExitTemplatename(c *TemplatenameContext)

	// ExitTemplateargumentlist is called when exiting the templateargumentlist production.
	ExitTemplateargumentlist(c *TemplateargumentlistContext)

	// ExitTemplateargument is called when exiting the templateargument production.
	ExitTemplateargument(c *TemplateargumentContext)

	// ExitTypenamespecifier is called when exiting the typenamespecifier production.
	ExitTypenamespecifier(c *TypenamespecifierContext)

	// ExitExplicitinstantiation is called when exiting the explicitinstantiation production.
	ExitExplicitinstantiation(c *ExplicitinstantiationContext)

	// ExitExplicitspecialization is called when exiting the explicitspecialization production.
	ExitExplicitspecialization(c *ExplicitspecializationContext)

	// ExitTryblock is called when exiting the tryblock production.
	ExitTryblock(c *TryblockContext)

	// ExitFunctiontryblock is called when exiting the functiontryblock production.
	ExitFunctiontryblock(c *FunctiontryblockContext)

	// ExitHandlerseq is called when exiting the handlerseq production.
	ExitHandlerseq(c *HandlerseqContext)

	// ExitHandler is called when exiting the handler production.
	ExitHandler(c *HandlerContext)

	// ExitExceptiondeclaration is called when exiting the exceptiondeclaration production.
	ExitExceptiondeclaration(c *ExceptiondeclarationContext)

	// ExitThrowexpression is called when exiting the throwexpression production.
	ExitThrowexpression(c *ThrowexpressionContext)

	// ExitExceptionspecification is called when exiting the exceptionspecification production.
	ExitExceptionspecification(c *ExceptionspecificationContext)

	// ExitDynamicexceptionspecification is called when exiting the dynamicexceptionspecification production.
	ExitDynamicexceptionspecification(c *DynamicexceptionspecificationContext)

	// ExitTypeidlist is called when exiting the typeidlist production.
	ExitTypeidlist(c *TypeidlistContext)

	// ExitNoexceptspecification is called when exiting the noexceptspecification production.
	ExitNoexceptspecification(c *NoexceptspecificationContext)

	// ExitRightShift is called when exiting the rightShift production.
	ExitRightShift(c *RightShiftContext)

	// ExitRightShiftAssign is called when exiting the rightShiftAssign production.
	ExitRightShiftAssign(c *RightShiftAssignContext)

	// ExitTheoperator is called when exiting the theoperator production.
	ExitTheoperator(c *TheoperatorContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitBooleanliteral is called when exiting the booleanliteral production.
	ExitBooleanliteral(c *BooleanliteralContext)

	// ExitPointerliteral is called when exiting the pointerliteral production.
	ExitPointerliteral(c *PointerliteralContext)

	// ExitUserdefinedliteral is called when exiting the userdefinedliteral production.
	ExitUserdefinedliteral(c *UserdefinedliteralContext)
}
