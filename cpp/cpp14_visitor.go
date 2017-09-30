// Generated from CPP14.g4 by ANTLR 4.7.

package cpp // CPP14
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CPP14Parser.
type CPP14Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CPP14Parser#translationunit.
	VisitTranslationunit(ctx *TranslationunitContext) interface{}

	// Visit a parse tree produced by CPP14Parser#primaryexpression.
	VisitPrimaryexpression(ctx *PrimaryexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#idexpression.
	VisitIdexpression(ctx *IdexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#unqualifiedid.
	VisitUnqualifiedid(ctx *UnqualifiedidContext) interface{}

	// Visit a parse tree produced by CPP14Parser#qualifiedid.
	VisitQualifiedid(ctx *QualifiedidContext) interface{}

	// Visit a parse tree produced by CPP14Parser#nestednamespecifier.
	VisitNestednamespecifier(ctx *NestednamespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#lambdaexpression.
	VisitLambdaexpression(ctx *LambdaexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#lambdaintroducer.
	VisitLambdaintroducer(ctx *LambdaintroducerContext) interface{}

	// Visit a parse tree produced by CPP14Parser#lambdacapture.
	VisitLambdacapture(ctx *LambdacaptureContext) interface{}

	// Visit a parse tree produced by CPP14Parser#capturedefault.
	VisitCapturedefault(ctx *CapturedefaultContext) interface{}

	// Visit a parse tree produced by CPP14Parser#capturelist.
	VisitCapturelist(ctx *CapturelistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#capture.
	VisitCapture(ctx *CaptureContext) interface{}

	// Visit a parse tree produced by CPP14Parser#simplecapture.
	VisitSimplecapture(ctx *SimplecaptureContext) interface{}

	// Visit a parse tree produced by CPP14Parser#initcapture.
	VisitInitcapture(ctx *InitcaptureContext) interface{}

	// Visit a parse tree produced by CPP14Parser#lambdadeclarator.
	VisitLambdadeclarator(ctx *LambdadeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#postfixexpression.
	VisitPostfixexpression(ctx *PostfixexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#expressionlist.
	VisitExpressionlist(ctx *ExpressionlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#pseudodestructorname.
	VisitPseudodestructorname(ctx *PseudodestructornameContext) interface{}

	// Visit a parse tree produced by CPP14Parser#unaryexpression.
	VisitUnaryexpression(ctx *UnaryexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#unaryoperator.
	VisitUnaryoperator(ctx *UnaryoperatorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#newexpression.
	VisitNewexpression(ctx *NewexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#newplacement.
	VisitNewplacement(ctx *NewplacementContext) interface{}

	// Visit a parse tree produced by CPP14Parser#newtypeid.
	VisitNewtypeid(ctx *NewtypeidContext) interface{}

	// Visit a parse tree produced by CPP14Parser#newdeclarator.
	VisitNewdeclarator(ctx *NewdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#noptrnewdeclarator.
	VisitNoptrnewdeclarator(ctx *NoptrnewdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#newinitializer.
	VisitNewinitializer(ctx *NewinitializerContext) interface{}

	// Visit a parse tree produced by CPP14Parser#deleteexpression.
	VisitDeleteexpression(ctx *DeleteexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#noexceptexpression.
	VisitNoexceptexpression(ctx *NoexceptexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#castexpression.
	VisitCastexpression(ctx *CastexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#pmexpression.
	VisitPmexpression(ctx *PmexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#multiplicativeexpression.
	VisitMultiplicativeexpression(ctx *MultiplicativeexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#additiveexpression.
	VisitAdditiveexpression(ctx *AdditiveexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#shiftexpression.
	VisitShiftexpression(ctx *ShiftexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#relationalexpression.
	VisitRelationalexpression(ctx *RelationalexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#equalityexpression.
	VisitEqualityexpression(ctx *EqualityexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#andexpression.
	VisitAndexpression(ctx *AndexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#exclusiveorexpression.
	VisitExclusiveorexpression(ctx *ExclusiveorexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#inclusiveorexpression.
	VisitInclusiveorexpression(ctx *InclusiveorexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#logicalandexpression.
	VisitLogicalandexpression(ctx *LogicalandexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#logicalorexpression.
	VisitLogicalorexpression(ctx *LogicalorexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#conditionalexpression.
	VisitConditionalexpression(ctx *ConditionalexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#assignmentexpression.
	VisitAssignmentexpression(ctx *AssignmentexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#assignmentoperator.
	VisitAssignmentoperator(ctx *AssignmentoperatorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#constantexpression.
	VisitConstantexpression(ctx *ConstantexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CPP14Parser#labeledstatement.
	VisitLabeledstatement(ctx *LabeledstatementContext) interface{}

	// Visit a parse tree produced by CPP14Parser#expressionstatement.
	VisitExpressionstatement(ctx *ExpressionstatementContext) interface{}

	// Visit a parse tree produced by CPP14Parser#compoundstatement.
	VisitCompoundstatement(ctx *CompoundstatementContext) interface{}

	// Visit a parse tree produced by CPP14Parser#statementseq.
	VisitStatementseq(ctx *StatementseqContext) interface{}

	// Visit a parse tree produced by CPP14Parser#selectionstatement.
	VisitSelectionstatement(ctx *SelectionstatementContext) interface{}

	// Visit a parse tree produced by CPP14Parser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#iterationstatement.
	VisitIterationstatement(ctx *IterationstatementContext) interface{}

	// Visit a parse tree produced by CPP14Parser#forinitstatement.
	VisitForinitstatement(ctx *ForinitstatementContext) interface{}

	// Visit a parse tree produced by CPP14Parser#forrangedeclaration.
	VisitForrangedeclaration(ctx *ForrangedeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#forrangeinitializer.
	VisitForrangeinitializer(ctx *ForrangeinitializerContext) interface{}

	// Visit a parse tree produced by CPP14Parser#jumpstatement.
	VisitJumpstatement(ctx *JumpstatementContext) interface{}

	// Visit a parse tree produced by CPP14Parser#declarationstatement.
	VisitDeclarationstatement(ctx *DeclarationstatementContext) interface{}

	// Visit a parse tree produced by CPP14Parser#declarationseq.
	VisitDeclarationseq(ctx *DeclarationseqContext) interface{}

	// Visit a parse tree produced by CPP14Parser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#blockdeclaration.
	VisitBlockdeclaration(ctx *BlockdeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#aliasdeclaration.
	VisitAliasdeclaration(ctx *AliasdeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#simpledeclaration.
	VisitSimpledeclaration(ctx *SimpledeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#static_assertdeclaration.
	VisitStatic_assertdeclaration(ctx *Static_assertdeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#emptydeclaration.
	VisitEmptydeclaration(ctx *EmptydeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#attributedeclaration.
	VisitAttributedeclaration(ctx *AttributedeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#declspecifier.
	VisitDeclspecifier(ctx *DeclspecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#declspecifierseq.
	VisitDeclspecifierseq(ctx *DeclspecifierseqContext) interface{}

	// Visit a parse tree produced by CPP14Parser#storageclassspecifier.
	VisitStorageclassspecifier(ctx *StorageclassspecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#functionspecifier.
	VisitFunctionspecifier(ctx *FunctionspecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#typedefname.
	VisitTypedefname(ctx *TypedefnameContext) interface{}

	// Visit a parse tree produced by CPP14Parser#typespecifier.
	VisitTypespecifier(ctx *TypespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#trailingtypespecifier.
	VisitTrailingtypespecifier(ctx *TrailingtypespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#typespecifierseq.
	VisitTypespecifierseq(ctx *TypespecifierseqContext) interface{}

	// Visit a parse tree produced by CPP14Parser#trailingtypespecifierseq.
	VisitTrailingtypespecifierseq(ctx *TrailingtypespecifierseqContext) interface{}

	// Visit a parse tree produced by CPP14Parser#simpletypespecifier.
	VisitSimpletypespecifier(ctx *SimpletypespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#thetypename.
	VisitThetypename(ctx *ThetypenameContext) interface{}

	// Visit a parse tree produced by CPP14Parser#decltypespecifier.
	VisitDecltypespecifier(ctx *DecltypespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#elaboratedtypespecifier.
	VisitElaboratedtypespecifier(ctx *ElaboratedtypespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#enumname.
	VisitEnumname(ctx *EnumnameContext) interface{}

	// Visit a parse tree produced by CPP14Parser#enumspecifier.
	VisitEnumspecifier(ctx *EnumspecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#enumhead.
	VisitEnumhead(ctx *EnumheadContext) interface{}

	// Visit a parse tree produced by CPP14Parser#opaqueenumdeclaration.
	VisitOpaqueenumdeclaration(ctx *OpaqueenumdeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#enumkey.
	VisitEnumkey(ctx *EnumkeyContext) interface{}

	// Visit a parse tree produced by CPP14Parser#enumbase.
	VisitEnumbase(ctx *EnumbaseContext) interface{}

	// Visit a parse tree produced by CPP14Parser#enumeratorlist.
	VisitEnumeratorlist(ctx *EnumeratorlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#enumeratordefinition.
	VisitEnumeratordefinition(ctx *EnumeratordefinitionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#enumerator.
	VisitEnumerator(ctx *EnumeratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#namespacename.
	VisitNamespacename(ctx *NamespacenameContext) interface{}

	// Visit a parse tree produced by CPP14Parser#originalnamespacename.
	VisitOriginalnamespacename(ctx *OriginalnamespacenameContext) interface{}

	// Visit a parse tree produced by CPP14Parser#namespacedefinition.
	VisitNamespacedefinition(ctx *NamespacedefinitionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#namednamespacedefinition.
	VisitNamednamespacedefinition(ctx *NamednamespacedefinitionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#originalnamespacedefinition.
	VisitOriginalnamespacedefinition(ctx *OriginalnamespacedefinitionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#extensionnamespacedefinition.
	VisitExtensionnamespacedefinition(ctx *ExtensionnamespacedefinitionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#unnamednamespacedefinition.
	VisitUnnamednamespacedefinition(ctx *UnnamednamespacedefinitionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#namespacebody.
	VisitNamespacebody(ctx *NamespacebodyContext) interface{}

	// Visit a parse tree produced by CPP14Parser#namespacealias.
	VisitNamespacealias(ctx *NamespacealiasContext) interface{}

	// Visit a parse tree produced by CPP14Parser#namespacealiasdefinition.
	VisitNamespacealiasdefinition(ctx *NamespacealiasdefinitionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#qualifiednamespacespecifier.
	VisitQualifiednamespacespecifier(ctx *QualifiednamespacespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#usingdeclaration.
	VisitUsingdeclaration(ctx *UsingdeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#usingdirective.
	VisitUsingdirective(ctx *UsingdirectiveContext) interface{}

	// Visit a parse tree produced by CPP14Parser#asmdefinition.
	VisitAsmdefinition(ctx *AsmdefinitionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#linkagespecification.
	VisitLinkagespecification(ctx *LinkagespecificationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#attributespecifierseq.
	VisitAttributespecifierseq(ctx *AttributespecifierseqContext) interface{}

	// Visit a parse tree produced by CPP14Parser#attributespecifier.
	VisitAttributespecifier(ctx *AttributespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#alignmentspecifier.
	VisitAlignmentspecifier(ctx *AlignmentspecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#attributelist.
	VisitAttributelist(ctx *AttributelistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by CPP14Parser#attributetoken.
	VisitAttributetoken(ctx *AttributetokenContext) interface{}

	// Visit a parse tree produced by CPP14Parser#attributescopedtoken.
	VisitAttributescopedtoken(ctx *AttributescopedtokenContext) interface{}

	// Visit a parse tree produced by CPP14Parser#attributenamespace.
	VisitAttributenamespace(ctx *AttributenamespaceContext) interface{}

	// Visit a parse tree produced by CPP14Parser#attributeargumentclause.
	VisitAttributeargumentclause(ctx *AttributeargumentclauseContext) interface{}

	// Visit a parse tree produced by CPP14Parser#balancedtokenseq.
	VisitBalancedtokenseq(ctx *BalancedtokenseqContext) interface{}

	// Visit a parse tree produced by CPP14Parser#balancedtoken.
	VisitBalancedtoken(ctx *BalancedtokenContext) interface{}

	// Visit a parse tree produced by CPP14Parser#initdeclaratorlist.
	VisitInitdeclaratorlist(ctx *InitdeclaratorlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#initdeclarator.
	VisitInitdeclarator(ctx *InitdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#declarator.
	VisitDeclarator(ctx *DeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#ptrdeclarator.
	VisitPtrdeclarator(ctx *PtrdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#noptrdeclarator.
	VisitNoptrdeclarator(ctx *NoptrdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#parametersandqualifiers.
	VisitParametersandqualifiers(ctx *ParametersandqualifiersContext) interface{}

	// Visit a parse tree produced by CPP14Parser#trailingreturntype.
	VisitTrailingreturntype(ctx *TrailingreturntypeContext) interface{}

	// Visit a parse tree produced by CPP14Parser#ptroperator.
	VisitPtroperator(ctx *PtroperatorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#cvqualifierseq.
	VisitCvqualifierseq(ctx *CvqualifierseqContext) interface{}

	// Visit a parse tree produced by CPP14Parser#cvqualifier.
	VisitCvqualifier(ctx *CvqualifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#refqualifier.
	VisitRefqualifier(ctx *RefqualifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#declaratorid.
	VisitDeclaratorid(ctx *DeclaratoridContext) interface{}

	// Visit a parse tree produced by CPP14Parser#thetypeid.
	VisitThetypeid(ctx *ThetypeidContext) interface{}

	// Visit a parse tree produced by CPP14Parser#abstractdeclarator.
	VisitAbstractdeclarator(ctx *AbstractdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#ptrabstractdeclarator.
	VisitPtrabstractdeclarator(ctx *PtrabstractdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#noptrabstractdeclarator.
	VisitNoptrabstractdeclarator(ctx *NoptrabstractdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#abstractpackdeclarator.
	VisitAbstractpackdeclarator(ctx *AbstractpackdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#noptrabstractpackdeclarator.
	VisitNoptrabstractpackdeclarator(ctx *NoptrabstractpackdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#parameterdeclarationclause.
	VisitParameterdeclarationclause(ctx *ParameterdeclarationclauseContext) interface{}

	// Visit a parse tree produced by CPP14Parser#parameterdeclarationlist.
	VisitParameterdeclarationlist(ctx *ParameterdeclarationlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#parameterdeclaration.
	VisitParameterdeclaration(ctx *ParameterdeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#functiondefinition.
	VisitFunctiondefinition(ctx *FunctiondefinitionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#functionbody.
	VisitFunctionbody(ctx *FunctionbodyContext) interface{}

	// Visit a parse tree produced by CPP14Parser#initializer.
	VisitInitializer(ctx *InitializerContext) interface{}

	// Visit a parse tree produced by CPP14Parser#braceorequalinitializer.
	VisitBraceorequalinitializer(ctx *BraceorequalinitializerContext) interface{}

	// Visit a parse tree produced by CPP14Parser#initializerclause.
	VisitInitializerclause(ctx *InitializerclauseContext) interface{}

	// Visit a parse tree produced by CPP14Parser#initializerlist.
	VisitInitializerlist(ctx *InitializerlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#bracedinitlist.
	VisitBracedinitlist(ctx *BracedinitlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#classname.
	VisitClassname(ctx *ClassnameContext) interface{}

	// Visit a parse tree produced by CPP14Parser#classspecifier.
	VisitClassspecifier(ctx *ClassspecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#classhead.
	VisitClasshead(ctx *ClassheadContext) interface{}

	// Visit a parse tree produced by CPP14Parser#classheadname.
	VisitClassheadname(ctx *ClassheadnameContext) interface{}

	// Visit a parse tree produced by CPP14Parser#classvirtspecifier.
	VisitClassvirtspecifier(ctx *ClassvirtspecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#classkey.
	VisitClasskey(ctx *ClasskeyContext) interface{}

	// Visit a parse tree produced by CPP14Parser#memberspecification.
	VisitMemberspecification(ctx *MemberspecificationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#memberdeclaration.
	VisitMemberdeclaration(ctx *MemberdeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#memberdeclaratorlist.
	VisitMemberdeclaratorlist(ctx *MemberdeclaratorlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#memberdeclarator.
	VisitMemberdeclarator(ctx *MemberdeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#virtspecifierseq.
	VisitVirtspecifierseq(ctx *VirtspecifierseqContext) interface{}

	// Visit a parse tree produced by CPP14Parser#virtspecifier.
	VisitVirtspecifier(ctx *VirtspecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#purespecifier.
	VisitPurespecifier(ctx *PurespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#baseclause.
	VisitBaseclause(ctx *BaseclauseContext) interface{}

	// Visit a parse tree produced by CPP14Parser#basespecifierlist.
	VisitBasespecifierlist(ctx *BasespecifierlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#basespecifier.
	VisitBasespecifier(ctx *BasespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#classordecltype.
	VisitClassordecltype(ctx *ClassordecltypeContext) interface{}

	// Visit a parse tree produced by CPP14Parser#basetypespecifier.
	VisitBasetypespecifier(ctx *BasetypespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#accessspecifier.
	VisitAccessspecifier(ctx *AccessspecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#conversionfunctionid.
	VisitConversionfunctionid(ctx *ConversionfunctionidContext) interface{}

	// Visit a parse tree produced by CPP14Parser#conversiontypeid.
	VisitConversiontypeid(ctx *ConversiontypeidContext) interface{}

	// Visit a parse tree produced by CPP14Parser#conversiondeclarator.
	VisitConversiondeclarator(ctx *ConversiondeclaratorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#ctorinitializer.
	VisitCtorinitializer(ctx *CtorinitializerContext) interface{}

	// Visit a parse tree produced by CPP14Parser#meminitializerlist.
	VisitMeminitializerlist(ctx *MeminitializerlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#meminitializer.
	VisitMeminitializer(ctx *MeminitializerContext) interface{}

	// Visit a parse tree produced by CPP14Parser#meminitializerid.
	VisitMeminitializerid(ctx *MeminitializeridContext) interface{}

	// Visit a parse tree produced by CPP14Parser#operatorfunctionid.
	VisitOperatorfunctionid(ctx *OperatorfunctionidContext) interface{}

	// Visit a parse tree produced by CPP14Parser#literaloperatorid.
	VisitLiteraloperatorid(ctx *LiteraloperatoridContext) interface{}

	// Visit a parse tree produced by CPP14Parser#templatedeclaration.
	VisitTemplatedeclaration(ctx *TemplatedeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#templateparameterlist.
	VisitTemplateparameterlist(ctx *TemplateparameterlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#templateparameter.
	VisitTemplateparameter(ctx *TemplateparameterContext) interface{}

	// Visit a parse tree produced by CPP14Parser#typeparameter.
	VisitTypeparameter(ctx *TypeparameterContext) interface{}

	// Visit a parse tree produced by CPP14Parser#simpletemplateid.
	VisitSimpletemplateid(ctx *SimpletemplateidContext) interface{}

	// Visit a parse tree produced by CPP14Parser#templateid.
	VisitTemplateid(ctx *TemplateidContext) interface{}

	// Visit a parse tree produced by CPP14Parser#templatename.
	VisitTemplatename(ctx *TemplatenameContext) interface{}

	// Visit a parse tree produced by CPP14Parser#templateargumentlist.
	VisitTemplateargumentlist(ctx *TemplateargumentlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#templateargument.
	VisitTemplateargument(ctx *TemplateargumentContext) interface{}

	// Visit a parse tree produced by CPP14Parser#typenamespecifier.
	VisitTypenamespecifier(ctx *TypenamespecifierContext) interface{}

	// Visit a parse tree produced by CPP14Parser#explicitinstantiation.
	VisitExplicitinstantiation(ctx *ExplicitinstantiationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#explicitspecialization.
	VisitExplicitspecialization(ctx *ExplicitspecializationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#tryblock.
	VisitTryblock(ctx *TryblockContext) interface{}

	// Visit a parse tree produced by CPP14Parser#functiontryblock.
	VisitFunctiontryblock(ctx *FunctiontryblockContext) interface{}

	// Visit a parse tree produced by CPP14Parser#handlerseq.
	VisitHandlerseq(ctx *HandlerseqContext) interface{}

	// Visit a parse tree produced by CPP14Parser#handler.
	VisitHandler(ctx *HandlerContext) interface{}

	// Visit a parse tree produced by CPP14Parser#exceptiondeclaration.
	VisitExceptiondeclaration(ctx *ExceptiondeclarationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#throwexpression.
	VisitThrowexpression(ctx *ThrowexpressionContext) interface{}

	// Visit a parse tree produced by CPP14Parser#exceptionspecification.
	VisitExceptionspecification(ctx *ExceptionspecificationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#dynamicexceptionspecification.
	VisitDynamicexceptionspecification(ctx *DynamicexceptionspecificationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#typeidlist.
	VisitTypeidlist(ctx *TypeidlistContext) interface{}

	// Visit a parse tree produced by CPP14Parser#noexceptspecification.
	VisitNoexceptspecification(ctx *NoexceptspecificationContext) interface{}

	// Visit a parse tree produced by CPP14Parser#rightShift.
	VisitRightShift(ctx *RightShiftContext) interface{}

	// Visit a parse tree produced by CPP14Parser#rightShiftAssign.
	VisitRightShiftAssign(ctx *RightShiftAssignContext) interface{}

	// Visit a parse tree produced by CPP14Parser#theoperator.
	VisitTheoperator(ctx *TheoperatorContext) interface{}

	// Visit a parse tree produced by CPP14Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by CPP14Parser#booleanliteral.
	VisitBooleanliteral(ctx *BooleanliteralContext) interface{}

	// Visit a parse tree produced by CPP14Parser#pointerliteral.
	VisitPointerliteral(ctx *PointerliteralContext) interface{}

	// Visit a parse tree produced by CPP14Parser#userdefinedliteral.
	VisitUserdefinedliteral(ctx *UserdefinedliteralContext) interface{}
}
