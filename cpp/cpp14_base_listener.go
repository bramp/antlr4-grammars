// Generated from CPP14.g4 by ANTLR 4.7.

package cpp // CPP14
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCPP14Listener is a complete listener for a parse tree produced by CPP14Parser.
type BaseCPP14Listener struct{}

var _ CPP14Listener = &BaseCPP14Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCPP14Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCPP14Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCPP14Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCPP14Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTranslationunit is called when production translationunit is entered.
func (s *BaseCPP14Listener) EnterTranslationunit(ctx *TranslationunitContext) {}

// ExitTranslationunit is called when production translationunit is exited.
func (s *BaseCPP14Listener) ExitTranslationunit(ctx *TranslationunitContext) {}

// EnterPrimaryexpression is called when production primaryexpression is entered.
func (s *BaseCPP14Listener) EnterPrimaryexpression(ctx *PrimaryexpressionContext) {}

// ExitPrimaryexpression is called when production primaryexpression is exited.
func (s *BaseCPP14Listener) ExitPrimaryexpression(ctx *PrimaryexpressionContext) {}

// EnterIdexpression is called when production idexpression is entered.
func (s *BaseCPP14Listener) EnterIdexpression(ctx *IdexpressionContext) {}

// ExitIdexpression is called when production idexpression is exited.
func (s *BaseCPP14Listener) ExitIdexpression(ctx *IdexpressionContext) {}

// EnterUnqualifiedid is called when production unqualifiedid is entered.
func (s *BaseCPP14Listener) EnterUnqualifiedid(ctx *UnqualifiedidContext) {}

// ExitUnqualifiedid is called when production unqualifiedid is exited.
func (s *BaseCPP14Listener) ExitUnqualifiedid(ctx *UnqualifiedidContext) {}

// EnterQualifiedid is called when production qualifiedid is entered.
func (s *BaseCPP14Listener) EnterQualifiedid(ctx *QualifiedidContext) {}

// ExitQualifiedid is called when production qualifiedid is exited.
func (s *BaseCPP14Listener) ExitQualifiedid(ctx *QualifiedidContext) {}

// EnterNestednamespecifier is called when production nestednamespecifier is entered.
func (s *BaseCPP14Listener) EnterNestednamespecifier(ctx *NestednamespecifierContext) {}

// ExitNestednamespecifier is called when production nestednamespecifier is exited.
func (s *BaseCPP14Listener) ExitNestednamespecifier(ctx *NestednamespecifierContext) {}

// EnterLambdaexpression is called when production lambdaexpression is entered.
func (s *BaseCPP14Listener) EnterLambdaexpression(ctx *LambdaexpressionContext) {}

// ExitLambdaexpression is called when production lambdaexpression is exited.
func (s *BaseCPP14Listener) ExitLambdaexpression(ctx *LambdaexpressionContext) {}

// EnterLambdaintroducer is called when production lambdaintroducer is entered.
func (s *BaseCPP14Listener) EnterLambdaintroducer(ctx *LambdaintroducerContext) {}

// ExitLambdaintroducer is called when production lambdaintroducer is exited.
func (s *BaseCPP14Listener) ExitLambdaintroducer(ctx *LambdaintroducerContext) {}

// EnterLambdacapture is called when production lambdacapture is entered.
func (s *BaseCPP14Listener) EnterLambdacapture(ctx *LambdacaptureContext) {}

// ExitLambdacapture is called when production lambdacapture is exited.
func (s *BaseCPP14Listener) ExitLambdacapture(ctx *LambdacaptureContext) {}

// EnterCapturedefault is called when production capturedefault is entered.
func (s *BaseCPP14Listener) EnterCapturedefault(ctx *CapturedefaultContext) {}

// ExitCapturedefault is called when production capturedefault is exited.
func (s *BaseCPP14Listener) ExitCapturedefault(ctx *CapturedefaultContext) {}

// EnterCapturelist is called when production capturelist is entered.
func (s *BaseCPP14Listener) EnterCapturelist(ctx *CapturelistContext) {}

// ExitCapturelist is called when production capturelist is exited.
func (s *BaseCPP14Listener) ExitCapturelist(ctx *CapturelistContext) {}

// EnterCapture is called when production capture is entered.
func (s *BaseCPP14Listener) EnterCapture(ctx *CaptureContext) {}

// ExitCapture is called when production capture is exited.
func (s *BaseCPP14Listener) ExitCapture(ctx *CaptureContext) {}

// EnterSimplecapture is called when production simplecapture is entered.
func (s *BaseCPP14Listener) EnterSimplecapture(ctx *SimplecaptureContext) {}

// ExitSimplecapture is called when production simplecapture is exited.
func (s *BaseCPP14Listener) ExitSimplecapture(ctx *SimplecaptureContext) {}

// EnterInitcapture is called when production initcapture is entered.
func (s *BaseCPP14Listener) EnterInitcapture(ctx *InitcaptureContext) {}

// ExitInitcapture is called when production initcapture is exited.
func (s *BaseCPP14Listener) ExitInitcapture(ctx *InitcaptureContext) {}

// EnterLambdadeclarator is called when production lambdadeclarator is entered.
func (s *BaseCPP14Listener) EnterLambdadeclarator(ctx *LambdadeclaratorContext) {}

// ExitLambdadeclarator is called when production lambdadeclarator is exited.
func (s *BaseCPP14Listener) ExitLambdadeclarator(ctx *LambdadeclaratorContext) {}

// EnterPostfixexpression is called when production postfixexpression is entered.
func (s *BaseCPP14Listener) EnterPostfixexpression(ctx *PostfixexpressionContext) {}

// ExitPostfixexpression is called when production postfixexpression is exited.
func (s *BaseCPP14Listener) ExitPostfixexpression(ctx *PostfixexpressionContext) {}

// EnterExpressionlist is called when production expressionlist is entered.
func (s *BaseCPP14Listener) EnterExpressionlist(ctx *ExpressionlistContext) {}

// ExitExpressionlist is called when production expressionlist is exited.
func (s *BaseCPP14Listener) ExitExpressionlist(ctx *ExpressionlistContext) {}

// EnterPseudodestructorname is called when production pseudodestructorname is entered.
func (s *BaseCPP14Listener) EnterPseudodestructorname(ctx *PseudodestructornameContext) {}

// ExitPseudodestructorname is called when production pseudodestructorname is exited.
func (s *BaseCPP14Listener) ExitPseudodestructorname(ctx *PseudodestructornameContext) {}

// EnterUnaryexpression is called when production unaryexpression is entered.
func (s *BaseCPP14Listener) EnterUnaryexpression(ctx *UnaryexpressionContext) {}

// ExitUnaryexpression is called when production unaryexpression is exited.
func (s *BaseCPP14Listener) ExitUnaryexpression(ctx *UnaryexpressionContext) {}

// EnterUnaryoperator is called when production unaryoperator is entered.
func (s *BaseCPP14Listener) EnterUnaryoperator(ctx *UnaryoperatorContext) {}

// ExitUnaryoperator is called when production unaryoperator is exited.
func (s *BaseCPP14Listener) ExitUnaryoperator(ctx *UnaryoperatorContext) {}

// EnterNewexpression is called when production newexpression is entered.
func (s *BaseCPP14Listener) EnterNewexpression(ctx *NewexpressionContext) {}

// ExitNewexpression is called when production newexpression is exited.
func (s *BaseCPP14Listener) ExitNewexpression(ctx *NewexpressionContext) {}

// EnterNewplacement is called when production newplacement is entered.
func (s *BaseCPP14Listener) EnterNewplacement(ctx *NewplacementContext) {}

// ExitNewplacement is called when production newplacement is exited.
func (s *BaseCPP14Listener) ExitNewplacement(ctx *NewplacementContext) {}

// EnterNewtypeid is called when production newtypeid is entered.
func (s *BaseCPP14Listener) EnterNewtypeid(ctx *NewtypeidContext) {}

// ExitNewtypeid is called when production newtypeid is exited.
func (s *BaseCPP14Listener) ExitNewtypeid(ctx *NewtypeidContext) {}

// EnterNewdeclarator is called when production newdeclarator is entered.
func (s *BaseCPP14Listener) EnterNewdeclarator(ctx *NewdeclaratorContext) {}

// ExitNewdeclarator is called when production newdeclarator is exited.
func (s *BaseCPP14Listener) ExitNewdeclarator(ctx *NewdeclaratorContext) {}

// EnterNoptrnewdeclarator is called when production noptrnewdeclarator is entered.
func (s *BaseCPP14Listener) EnterNoptrnewdeclarator(ctx *NoptrnewdeclaratorContext) {}

// ExitNoptrnewdeclarator is called when production noptrnewdeclarator is exited.
func (s *BaseCPP14Listener) ExitNoptrnewdeclarator(ctx *NoptrnewdeclaratorContext) {}

// EnterNewinitializer is called when production newinitializer is entered.
func (s *BaseCPP14Listener) EnterNewinitializer(ctx *NewinitializerContext) {}

// ExitNewinitializer is called when production newinitializer is exited.
func (s *BaseCPP14Listener) ExitNewinitializer(ctx *NewinitializerContext) {}

// EnterDeleteexpression is called when production deleteexpression is entered.
func (s *BaseCPP14Listener) EnterDeleteexpression(ctx *DeleteexpressionContext) {}

// ExitDeleteexpression is called when production deleteexpression is exited.
func (s *BaseCPP14Listener) ExitDeleteexpression(ctx *DeleteexpressionContext) {}

// EnterNoexceptexpression is called when production noexceptexpression is entered.
func (s *BaseCPP14Listener) EnterNoexceptexpression(ctx *NoexceptexpressionContext) {}

// ExitNoexceptexpression is called when production noexceptexpression is exited.
func (s *BaseCPP14Listener) ExitNoexceptexpression(ctx *NoexceptexpressionContext) {}

// EnterCastexpression is called when production castexpression is entered.
func (s *BaseCPP14Listener) EnterCastexpression(ctx *CastexpressionContext) {}

// ExitCastexpression is called when production castexpression is exited.
func (s *BaseCPP14Listener) ExitCastexpression(ctx *CastexpressionContext) {}

// EnterPmexpression is called when production pmexpression is entered.
func (s *BaseCPP14Listener) EnterPmexpression(ctx *PmexpressionContext) {}

// ExitPmexpression is called when production pmexpression is exited.
func (s *BaseCPP14Listener) ExitPmexpression(ctx *PmexpressionContext) {}

// EnterMultiplicativeexpression is called when production multiplicativeexpression is entered.
func (s *BaseCPP14Listener) EnterMultiplicativeexpression(ctx *MultiplicativeexpressionContext) {}

// ExitMultiplicativeexpression is called when production multiplicativeexpression is exited.
func (s *BaseCPP14Listener) ExitMultiplicativeexpression(ctx *MultiplicativeexpressionContext) {}

// EnterAdditiveexpression is called when production additiveexpression is entered.
func (s *BaseCPP14Listener) EnterAdditiveexpression(ctx *AdditiveexpressionContext) {}

// ExitAdditiveexpression is called when production additiveexpression is exited.
func (s *BaseCPP14Listener) ExitAdditiveexpression(ctx *AdditiveexpressionContext) {}

// EnterShiftexpression is called when production shiftexpression is entered.
func (s *BaseCPP14Listener) EnterShiftexpression(ctx *ShiftexpressionContext) {}

// ExitShiftexpression is called when production shiftexpression is exited.
func (s *BaseCPP14Listener) ExitShiftexpression(ctx *ShiftexpressionContext) {}

// EnterRelationalexpression is called when production relationalexpression is entered.
func (s *BaseCPP14Listener) EnterRelationalexpression(ctx *RelationalexpressionContext) {}

// ExitRelationalexpression is called when production relationalexpression is exited.
func (s *BaseCPP14Listener) ExitRelationalexpression(ctx *RelationalexpressionContext) {}

// EnterEqualityexpression is called when production equalityexpression is entered.
func (s *BaseCPP14Listener) EnterEqualityexpression(ctx *EqualityexpressionContext) {}

// ExitEqualityexpression is called when production equalityexpression is exited.
func (s *BaseCPP14Listener) ExitEqualityexpression(ctx *EqualityexpressionContext) {}

// EnterAndexpression is called when production andexpression is entered.
func (s *BaseCPP14Listener) EnterAndexpression(ctx *AndexpressionContext) {}

// ExitAndexpression is called when production andexpression is exited.
func (s *BaseCPP14Listener) ExitAndexpression(ctx *AndexpressionContext) {}

// EnterExclusiveorexpression is called when production exclusiveorexpression is entered.
func (s *BaseCPP14Listener) EnterExclusiveorexpression(ctx *ExclusiveorexpressionContext) {}

// ExitExclusiveorexpression is called when production exclusiveorexpression is exited.
func (s *BaseCPP14Listener) ExitExclusiveorexpression(ctx *ExclusiveorexpressionContext) {}

// EnterInclusiveorexpression is called when production inclusiveorexpression is entered.
func (s *BaseCPP14Listener) EnterInclusiveorexpression(ctx *InclusiveorexpressionContext) {}

// ExitInclusiveorexpression is called when production inclusiveorexpression is exited.
func (s *BaseCPP14Listener) ExitInclusiveorexpression(ctx *InclusiveorexpressionContext) {}

// EnterLogicalandexpression is called when production logicalandexpression is entered.
func (s *BaseCPP14Listener) EnterLogicalandexpression(ctx *LogicalandexpressionContext) {}

// ExitLogicalandexpression is called when production logicalandexpression is exited.
func (s *BaseCPP14Listener) ExitLogicalandexpression(ctx *LogicalandexpressionContext) {}

// EnterLogicalorexpression is called when production logicalorexpression is entered.
func (s *BaseCPP14Listener) EnterLogicalorexpression(ctx *LogicalorexpressionContext) {}

// ExitLogicalorexpression is called when production logicalorexpression is exited.
func (s *BaseCPP14Listener) ExitLogicalorexpression(ctx *LogicalorexpressionContext) {}

// EnterConditionalexpression is called when production conditionalexpression is entered.
func (s *BaseCPP14Listener) EnterConditionalexpression(ctx *ConditionalexpressionContext) {}

// ExitConditionalexpression is called when production conditionalexpression is exited.
func (s *BaseCPP14Listener) ExitConditionalexpression(ctx *ConditionalexpressionContext) {}

// EnterAssignmentexpression is called when production assignmentexpression is entered.
func (s *BaseCPP14Listener) EnterAssignmentexpression(ctx *AssignmentexpressionContext) {}

// ExitAssignmentexpression is called when production assignmentexpression is exited.
func (s *BaseCPP14Listener) ExitAssignmentexpression(ctx *AssignmentexpressionContext) {}

// EnterAssignmentoperator is called when production assignmentoperator is entered.
func (s *BaseCPP14Listener) EnterAssignmentoperator(ctx *AssignmentoperatorContext) {}

// ExitAssignmentoperator is called when production assignmentoperator is exited.
func (s *BaseCPP14Listener) ExitAssignmentoperator(ctx *AssignmentoperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCPP14Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCPP14Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterConstantexpression is called when production constantexpression is entered.
func (s *BaseCPP14Listener) EnterConstantexpression(ctx *ConstantexpressionContext) {}

// ExitConstantexpression is called when production constantexpression is exited.
func (s *BaseCPP14Listener) ExitConstantexpression(ctx *ConstantexpressionContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCPP14Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCPP14Listener) ExitStatement(ctx *StatementContext) {}

// EnterLabeledstatement is called when production labeledstatement is entered.
func (s *BaseCPP14Listener) EnterLabeledstatement(ctx *LabeledstatementContext) {}

// ExitLabeledstatement is called when production labeledstatement is exited.
func (s *BaseCPP14Listener) ExitLabeledstatement(ctx *LabeledstatementContext) {}

// EnterExpressionstatement is called when production expressionstatement is entered.
func (s *BaseCPP14Listener) EnterExpressionstatement(ctx *ExpressionstatementContext) {}

// ExitExpressionstatement is called when production expressionstatement is exited.
func (s *BaseCPP14Listener) ExitExpressionstatement(ctx *ExpressionstatementContext) {}

// EnterCompoundstatement is called when production compoundstatement is entered.
func (s *BaseCPP14Listener) EnterCompoundstatement(ctx *CompoundstatementContext) {}

// ExitCompoundstatement is called when production compoundstatement is exited.
func (s *BaseCPP14Listener) ExitCompoundstatement(ctx *CompoundstatementContext) {}

// EnterStatementseq is called when production statementseq is entered.
func (s *BaseCPP14Listener) EnterStatementseq(ctx *StatementseqContext) {}

// ExitStatementseq is called when production statementseq is exited.
func (s *BaseCPP14Listener) ExitStatementseq(ctx *StatementseqContext) {}

// EnterSelectionstatement is called when production selectionstatement is entered.
func (s *BaseCPP14Listener) EnterSelectionstatement(ctx *SelectionstatementContext) {}

// ExitSelectionstatement is called when production selectionstatement is exited.
func (s *BaseCPP14Listener) ExitSelectionstatement(ctx *SelectionstatementContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseCPP14Listener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseCPP14Listener) ExitCondition(ctx *ConditionContext) {}

// EnterIterationstatement is called when production iterationstatement is entered.
func (s *BaseCPP14Listener) EnterIterationstatement(ctx *IterationstatementContext) {}

// ExitIterationstatement is called when production iterationstatement is exited.
func (s *BaseCPP14Listener) ExitIterationstatement(ctx *IterationstatementContext) {}

// EnterForinitstatement is called when production forinitstatement is entered.
func (s *BaseCPP14Listener) EnterForinitstatement(ctx *ForinitstatementContext) {}

// ExitForinitstatement is called when production forinitstatement is exited.
func (s *BaseCPP14Listener) ExitForinitstatement(ctx *ForinitstatementContext) {}

// EnterForrangedeclaration is called when production forrangedeclaration is entered.
func (s *BaseCPP14Listener) EnterForrangedeclaration(ctx *ForrangedeclarationContext) {}

// ExitForrangedeclaration is called when production forrangedeclaration is exited.
func (s *BaseCPP14Listener) ExitForrangedeclaration(ctx *ForrangedeclarationContext) {}

// EnterForrangeinitializer is called when production forrangeinitializer is entered.
func (s *BaseCPP14Listener) EnterForrangeinitializer(ctx *ForrangeinitializerContext) {}

// ExitForrangeinitializer is called when production forrangeinitializer is exited.
func (s *BaseCPP14Listener) ExitForrangeinitializer(ctx *ForrangeinitializerContext) {}

// EnterJumpstatement is called when production jumpstatement is entered.
func (s *BaseCPP14Listener) EnterJumpstatement(ctx *JumpstatementContext) {}

// ExitJumpstatement is called when production jumpstatement is exited.
func (s *BaseCPP14Listener) ExitJumpstatement(ctx *JumpstatementContext) {}

// EnterDeclarationstatement is called when production declarationstatement is entered.
func (s *BaseCPP14Listener) EnterDeclarationstatement(ctx *DeclarationstatementContext) {}

// ExitDeclarationstatement is called when production declarationstatement is exited.
func (s *BaseCPP14Listener) ExitDeclarationstatement(ctx *DeclarationstatementContext) {}

// EnterDeclarationseq is called when production declarationseq is entered.
func (s *BaseCPP14Listener) EnterDeclarationseq(ctx *DeclarationseqContext) {}

// ExitDeclarationseq is called when production declarationseq is exited.
func (s *BaseCPP14Listener) ExitDeclarationseq(ctx *DeclarationseqContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseCPP14Listener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseCPP14Listener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterBlockdeclaration is called when production blockdeclaration is entered.
func (s *BaseCPP14Listener) EnterBlockdeclaration(ctx *BlockdeclarationContext) {}

// ExitBlockdeclaration is called when production blockdeclaration is exited.
func (s *BaseCPP14Listener) ExitBlockdeclaration(ctx *BlockdeclarationContext) {}

// EnterAliasdeclaration is called when production aliasdeclaration is entered.
func (s *BaseCPP14Listener) EnterAliasdeclaration(ctx *AliasdeclarationContext) {}

// ExitAliasdeclaration is called when production aliasdeclaration is exited.
func (s *BaseCPP14Listener) ExitAliasdeclaration(ctx *AliasdeclarationContext) {}

// EnterSimpledeclaration is called when production simpledeclaration is entered.
func (s *BaseCPP14Listener) EnterSimpledeclaration(ctx *SimpledeclarationContext) {}

// ExitSimpledeclaration is called when production simpledeclaration is exited.
func (s *BaseCPP14Listener) ExitSimpledeclaration(ctx *SimpledeclarationContext) {}

// EnterStatic_assertdeclaration is called when production static_assertdeclaration is entered.
func (s *BaseCPP14Listener) EnterStatic_assertdeclaration(ctx *Static_assertdeclarationContext) {}

// ExitStatic_assertdeclaration is called when production static_assertdeclaration is exited.
func (s *BaseCPP14Listener) ExitStatic_assertdeclaration(ctx *Static_assertdeclarationContext) {}

// EnterEmptydeclaration is called when production emptydeclaration is entered.
func (s *BaseCPP14Listener) EnterEmptydeclaration(ctx *EmptydeclarationContext) {}

// ExitEmptydeclaration is called when production emptydeclaration is exited.
func (s *BaseCPP14Listener) ExitEmptydeclaration(ctx *EmptydeclarationContext) {}

// EnterAttributedeclaration is called when production attributedeclaration is entered.
func (s *BaseCPP14Listener) EnterAttributedeclaration(ctx *AttributedeclarationContext) {}

// ExitAttributedeclaration is called when production attributedeclaration is exited.
func (s *BaseCPP14Listener) ExitAttributedeclaration(ctx *AttributedeclarationContext) {}

// EnterDeclspecifier is called when production declspecifier is entered.
func (s *BaseCPP14Listener) EnterDeclspecifier(ctx *DeclspecifierContext) {}

// ExitDeclspecifier is called when production declspecifier is exited.
func (s *BaseCPP14Listener) ExitDeclspecifier(ctx *DeclspecifierContext) {}

// EnterDeclspecifierseq is called when production declspecifierseq is entered.
func (s *BaseCPP14Listener) EnterDeclspecifierseq(ctx *DeclspecifierseqContext) {}

// ExitDeclspecifierseq is called when production declspecifierseq is exited.
func (s *BaseCPP14Listener) ExitDeclspecifierseq(ctx *DeclspecifierseqContext) {}

// EnterStorageclassspecifier is called when production storageclassspecifier is entered.
func (s *BaseCPP14Listener) EnterStorageclassspecifier(ctx *StorageclassspecifierContext) {}

// ExitStorageclassspecifier is called when production storageclassspecifier is exited.
func (s *BaseCPP14Listener) ExitStorageclassspecifier(ctx *StorageclassspecifierContext) {}

// EnterFunctionspecifier is called when production functionspecifier is entered.
func (s *BaseCPP14Listener) EnterFunctionspecifier(ctx *FunctionspecifierContext) {}

// ExitFunctionspecifier is called when production functionspecifier is exited.
func (s *BaseCPP14Listener) ExitFunctionspecifier(ctx *FunctionspecifierContext) {}

// EnterTypedefname is called when production typedefname is entered.
func (s *BaseCPP14Listener) EnterTypedefname(ctx *TypedefnameContext) {}

// ExitTypedefname is called when production typedefname is exited.
func (s *BaseCPP14Listener) ExitTypedefname(ctx *TypedefnameContext) {}

// EnterTypespecifier is called when production typespecifier is entered.
func (s *BaseCPP14Listener) EnterTypespecifier(ctx *TypespecifierContext) {}

// ExitTypespecifier is called when production typespecifier is exited.
func (s *BaseCPP14Listener) ExitTypespecifier(ctx *TypespecifierContext) {}

// EnterTrailingtypespecifier is called when production trailingtypespecifier is entered.
func (s *BaseCPP14Listener) EnterTrailingtypespecifier(ctx *TrailingtypespecifierContext) {}

// ExitTrailingtypespecifier is called when production trailingtypespecifier is exited.
func (s *BaseCPP14Listener) ExitTrailingtypespecifier(ctx *TrailingtypespecifierContext) {}

// EnterTypespecifierseq is called when production typespecifierseq is entered.
func (s *BaseCPP14Listener) EnterTypespecifierseq(ctx *TypespecifierseqContext) {}

// ExitTypespecifierseq is called when production typespecifierseq is exited.
func (s *BaseCPP14Listener) ExitTypespecifierseq(ctx *TypespecifierseqContext) {}

// EnterTrailingtypespecifierseq is called when production trailingtypespecifierseq is entered.
func (s *BaseCPP14Listener) EnterTrailingtypespecifierseq(ctx *TrailingtypespecifierseqContext) {}

// ExitTrailingtypespecifierseq is called when production trailingtypespecifierseq is exited.
func (s *BaseCPP14Listener) ExitTrailingtypespecifierseq(ctx *TrailingtypespecifierseqContext) {}

// EnterSimpletypespecifier is called when production simpletypespecifier is entered.
func (s *BaseCPP14Listener) EnterSimpletypespecifier(ctx *SimpletypespecifierContext) {}

// ExitSimpletypespecifier is called when production simpletypespecifier is exited.
func (s *BaseCPP14Listener) ExitSimpletypespecifier(ctx *SimpletypespecifierContext) {}

// EnterThetypename is called when production thetypename is entered.
func (s *BaseCPP14Listener) EnterThetypename(ctx *ThetypenameContext) {}

// ExitThetypename is called when production thetypename is exited.
func (s *BaseCPP14Listener) ExitThetypename(ctx *ThetypenameContext) {}

// EnterDecltypespecifier is called when production decltypespecifier is entered.
func (s *BaseCPP14Listener) EnterDecltypespecifier(ctx *DecltypespecifierContext) {}

// ExitDecltypespecifier is called when production decltypespecifier is exited.
func (s *BaseCPP14Listener) ExitDecltypespecifier(ctx *DecltypespecifierContext) {}

// EnterElaboratedtypespecifier is called when production elaboratedtypespecifier is entered.
func (s *BaseCPP14Listener) EnterElaboratedtypespecifier(ctx *ElaboratedtypespecifierContext) {}

// ExitElaboratedtypespecifier is called when production elaboratedtypespecifier is exited.
func (s *BaseCPP14Listener) ExitElaboratedtypespecifier(ctx *ElaboratedtypespecifierContext) {}

// EnterEnumname is called when production enumname is entered.
func (s *BaseCPP14Listener) EnterEnumname(ctx *EnumnameContext) {}

// ExitEnumname is called when production enumname is exited.
func (s *BaseCPP14Listener) ExitEnumname(ctx *EnumnameContext) {}

// EnterEnumspecifier is called when production enumspecifier is entered.
func (s *BaseCPP14Listener) EnterEnumspecifier(ctx *EnumspecifierContext) {}

// ExitEnumspecifier is called when production enumspecifier is exited.
func (s *BaseCPP14Listener) ExitEnumspecifier(ctx *EnumspecifierContext) {}

// EnterEnumhead is called when production enumhead is entered.
func (s *BaseCPP14Listener) EnterEnumhead(ctx *EnumheadContext) {}

// ExitEnumhead is called when production enumhead is exited.
func (s *BaseCPP14Listener) ExitEnumhead(ctx *EnumheadContext) {}

// EnterOpaqueenumdeclaration is called when production opaqueenumdeclaration is entered.
func (s *BaseCPP14Listener) EnterOpaqueenumdeclaration(ctx *OpaqueenumdeclarationContext) {}

// ExitOpaqueenumdeclaration is called when production opaqueenumdeclaration is exited.
func (s *BaseCPP14Listener) ExitOpaqueenumdeclaration(ctx *OpaqueenumdeclarationContext) {}

// EnterEnumkey is called when production enumkey is entered.
func (s *BaseCPP14Listener) EnterEnumkey(ctx *EnumkeyContext) {}

// ExitEnumkey is called when production enumkey is exited.
func (s *BaseCPP14Listener) ExitEnumkey(ctx *EnumkeyContext) {}

// EnterEnumbase is called when production enumbase is entered.
func (s *BaseCPP14Listener) EnterEnumbase(ctx *EnumbaseContext) {}

// ExitEnumbase is called when production enumbase is exited.
func (s *BaseCPP14Listener) ExitEnumbase(ctx *EnumbaseContext) {}

// EnterEnumeratorlist is called when production enumeratorlist is entered.
func (s *BaseCPP14Listener) EnterEnumeratorlist(ctx *EnumeratorlistContext) {}

// ExitEnumeratorlist is called when production enumeratorlist is exited.
func (s *BaseCPP14Listener) ExitEnumeratorlist(ctx *EnumeratorlistContext) {}

// EnterEnumeratordefinition is called when production enumeratordefinition is entered.
func (s *BaseCPP14Listener) EnterEnumeratordefinition(ctx *EnumeratordefinitionContext) {}

// ExitEnumeratordefinition is called when production enumeratordefinition is exited.
func (s *BaseCPP14Listener) ExitEnumeratordefinition(ctx *EnumeratordefinitionContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseCPP14Listener) EnterEnumerator(ctx *EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseCPP14Listener) ExitEnumerator(ctx *EnumeratorContext) {}

// EnterNamespacename is called when production namespacename is entered.
func (s *BaseCPP14Listener) EnterNamespacename(ctx *NamespacenameContext) {}

// ExitNamespacename is called when production namespacename is exited.
func (s *BaseCPP14Listener) ExitNamespacename(ctx *NamespacenameContext) {}

// EnterOriginalnamespacename is called when production originalnamespacename is entered.
func (s *BaseCPP14Listener) EnterOriginalnamespacename(ctx *OriginalnamespacenameContext) {}

// ExitOriginalnamespacename is called when production originalnamespacename is exited.
func (s *BaseCPP14Listener) ExitOriginalnamespacename(ctx *OriginalnamespacenameContext) {}

// EnterNamespacedefinition is called when production namespacedefinition is entered.
func (s *BaseCPP14Listener) EnterNamespacedefinition(ctx *NamespacedefinitionContext) {}

// ExitNamespacedefinition is called when production namespacedefinition is exited.
func (s *BaseCPP14Listener) ExitNamespacedefinition(ctx *NamespacedefinitionContext) {}

// EnterNamednamespacedefinition is called when production namednamespacedefinition is entered.
func (s *BaseCPP14Listener) EnterNamednamespacedefinition(ctx *NamednamespacedefinitionContext) {}

// ExitNamednamespacedefinition is called when production namednamespacedefinition is exited.
func (s *BaseCPP14Listener) ExitNamednamespacedefinition(ctx *NamednamespacedefinitionContext) {}

// EnterOriginalnamespacedefinition is called when production originalnamespacedefinition is entered.
func (s *BaseCPP14Listener) EnterOriginalnamespacedefinition(ctx *OriginalnamespacedefinitionContext) {
}

// ExitOriginalnamespacedefinition is called when production originalnamespacedefinition is exited.
func (s *BaseCPP14Listener) ExitOriginalnamespacedefinition(ctx *OriginalnamespacedefinitionContext) {}

// EnterExtensionnamespacedefinition is called when production extensionnamespacedefinition is entered.
func (s *BaseCPP14Listener) EnterExtensionnamespacedefinition(ctx *ExtensionnamespacedefinitionContext) {
}

// ExitExtensionnamespacedefinition is called when production extensionnamespacedefinition is exited.
func (s *BaseCPP14Listener) ExitExtensionnamespacedefinition(ctx *ExtensionnamespacedefinitionContext) {
}

// EnterUnnamednamespacedefinition is called when production unnamednamespacedefinition is entered.
func (s *BaseCPP14Listener) EnterUnnamednamespacedefinition(ctx *UnnamednamespacedefinitionContext) {}

// ExitUnnamednamespacedefinition is called when production unnamednamespacedefinition is exited.
func (s *BaseCPP14Listener) ExitUnnamednamespacedefinition(ctx *UnnamednamespacedefinitionContext) {}

// EnterNamespacebody is called when production namespacebody is entered.
func (s *BaseCPP14Listener) EnterNamespacebody(ctx *NamespacebodyContext) {}

// ExitNamespacebody is called when production namespacebody is exited.
func (s *BaseCPP14Listener) ExitNamespacebody(ctx *NamespacebodyContext) {}

// EnterNamespacealias is called when production namespacealias is entered.
func (s *BaseCPP14Listener) EnterNamespacealias(ctx *NamespacealiasContext) {}

// ExitNamespacealias is called when production namespacealias is exited.
func (s *BaseCPP14Listener) ExitNamespacealias(ctx *NamespacealiasContext) {}

// EnterNamespacealiasdefinition is called when production namespacealiasdefinition is entered.
func (s *BaseCPP14Listener) EnterNamespacealiasdefinition(ctx *NamespacealiasdefinitionContext) {}

// ExitNamespacealiasdefinition is called when production namespacealiasdefinition is exited.
func (s *BaseCPP14Listener) ExitNamespacealiasdefinition(ctx *NamespacealiasdefinitionContext) {}

// EnterQualifiednamespacespecifier is called when production qualifiednamespacespecifier is entered.
func (s *BaseCPP14Listener) EnterQualifiednamespacespecifier(ctx *QualifiednamespacespecifierContext) {
}

// ExitQualifiednamespacespecifier is called when production qualifiednamespacespecifier is exited.
func (s *BaseCPP14Listener) ExitQualifiednamespacespecifier(ctx *QualifiednamespacespecifierContext) {}

// EnterUsingdeclaration is called when production usingdeclaration is entered.
func (s *BaseCPP14Listener) EnterUsingdeclaration(ctx *UsingdeclarationContext) {}

// ExitUsingdeclaration is called when production usingdeclaration is exited.
func (s *BaseCPP14Listener) ExitUsingdeclaration(ctx *UsingdeclarationContext) {}

// EnterUsingdirective is called when production usingdirective is entered.
func (s *BaseCPP14Listener) EnterUsingdirective(ctx *UsingdirectiveContext) {}

// ExitUsingdirective is called when production usingdirective is exited.
func (s *BaseCPP14Listener) ExitUsingdirective(ctx *UsingdirectiveContext) {}

// EnterAsmdefinition is called when production asmdefinition is entered.
func (s *BaseCPP14Listener) EnterAsmdefinition(ctx *AsmdefinitionContext) {}

// ExitAsmdefinition is called when production asmdefinition is exited.
func (s *BaseCPP14Listener) ExitAsmdefinition(ctx *AsmdefinitionContext) {}

// EnterLinkagespecification is called when production linkagespecification is entered.
func (s *BaseCPP14Listener) EnterLinkagespecification(ctx *LinkagespecificationContext) {}

// ExitLinkagespecification is called when production linkagespecification is exited.
func (s *BaseCPP14Listener) ExitLinkagespecification(ctx *LinkagespecificationContext) {}

// EnterAttributespecifierseq is called when production attributespecifierseq is entered.
func (s *BaseCPP14Listener) EnterAttributespecifierseq(ctx *AttributespecifierseqContext) {}

// ExitAttributespecifierseq is called when production attributespecifierseq is exited.
func (s *BaseCPP14Listener) ExitAttributespecifierseq(ctx *AttributespecifierseqContext) {}

// EnterAttributespecifier is called when production attributespecifier is entered.
func (s *BaseCPP14Listener) EnterAttributespecifier(ctx *AttributespecifierContext) {}

// ExitAttributespecifier is called when production attributespecifier is exited.
func (s *BaseCPP14Listener) ExitAttributespecifier(ctx *AttributespecifierContext) {}

// EnterAlignmentspecifier is called when production alignmentspecifier is entered.
func (s *BaseCPP14Listener) EnterAlignmentspecifier(ctx *AlignmentspecifierContext) {}

// ExitAlignmentspecifier is called when production alignmentspecifier is exited.
func (s *BaseCPP14Listener) ExitAlignmentspecifier(ctx *AlignmentspecifierContext) {}

// EnterAttributelist is called when production attributelist is entered.
func (s *BaseCPP14Listener) EnterAttributelist(ctx *AttributelistContext) {}

// ExitAttributelist is called when production attributelist is exited.
func (s *BaseCPP14Listener) ExitAttributelist(ctx *AttributelistContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseCPP14Listener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseCPP14Listener) ExitAttribute(ctx *AttributeContext) {}

// EnterAttributetoken is called when production attributetoken is entered.
func (s *BaseCPP14Listener) EnterAttributetoken(ctx *AttributetokenContext) {}

// ExitAttributetoken is called when production attributetoken is exited.
func (s *BaseCPP14Listener) ExitAttributetoken(ctx *AttributetokenContext) {}

// EnterAttributescopedtoken is called when production attributescopedtoken is entered.
func (s *BaseCPP14Listener) EnterAttributescopedtoken(ctx *AttributescopedtokenContext) {}

// ExitAttributescopedtoken is called when production attributescopedtoken is exited.
func (s *BaseCPP14Listener) ExitAttributescopedtoken(ctx *AttributescopedtokenContext) {}

// EnterAttributenamespace is called when production attributenamespace is entered.
func (s *BaseCPP14Listener) EnterAttributenamespace(ctx *AttributenamespaceContext) {}

// ExitAttributenamespace is called when production attributenamespace is exited.
func (s *BaseCPP14Listener) ExitAttributenamespace(ctx *AttributenamespaceContext) {}

// EnterAttributeargumentclause is called when production attributeargumentclause is entered.
func (s *BaseCPP14Listener) EnterAttributeargumentclause(ctx *AttributeargumentclauseContext) {}

// ExitAttributeargumentclause is called when production attributeargumentclause is exited.
func (s *BaseCPP14Listener) ExitAttributeargumentclause(ctx *AttributeargumentclauseContext) {}

// EnterBalancedtokenseq is called when production balancedtokenseq is entered.
func (s *BaseCPP14Listener) EnterBalancedtokenseq(ctx *BalancedtokenseqContext) {}

// ExitBalancedtokenseq is called when production balancedtokenseq is exited.
func (s *BaseCPP14Listener) ExitBalancedtokenseq(ctx *BalancedtokenseqContext) {}

// EnterBalancedtoken is called when production balancedtoken is entered.
func (s *BaseCPP14Listener) EnterBalancedtoken(ctx *BalancedtokenContext) {}

// ExitBalancedtoken is called when production balancedtoken is exited.
func (s *BaseCPP14Listener) ExitBalancedtoken(ctx *BalancedtokenContext) {}

// EnterInitdeclaratorlist is called when production initdeclaratorlist is entered.
func (s *BaseCPP14Listener) EnterInitdeclaratorlist(ctx *InitdeclaratorlistContext) {}

// ExitInitdeclaratorlist is called when production initdeclaratorlist is exited.
func (s *BaseCPP14Listener) ExitInitdeclaratorlist(ctx *InitdeclaratorlistContext) {}

// EnterInitdeclarator is called when production initdeclarator is entered.
func (s *BaseCPP14Listener) EnterInitdeclarator(ctx *InitdeclaratorContext) {}

// ExitInitdeclarator is called when production initdeclarator is exited.
func (s *BaseCPP14Listener) ExitInitdeclarator(ctx *InitdeclaratorContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseCPP14Listener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseCPP14Listener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterPtrdeclarator is called when production ptrdeclarator is entered.
func (s *BaseCPP14Listener) EnterPtrdeclarator(ctx *PtrdeclaratorContext) {}

// ExitPtrdeclarator is called when production ptrdeclarator is exited.
func (s *BaseCPP14Listener) ExitPtrdeclarator(ctx *PtrdeclaratorContext) {}

// EnterNoptrdeclarator is called when production noptrdeclarator is entered.
func (s *BaseCPP14Listener) EnterNoptrdeclarator(ctx *NoptrdeclaratorContext) {}

// ExitNoptrdeclarator is called when production noptrdeclarator is exited.
func (s *BaseCPP14Listener) ExitNoptrdeclarator(ctx *NoptrdeclaratorContext) {}

// EnterParametersandqualifiers is called when production parametersandqualifiers is entered.
func (s *BaseCPP14Listener) EnterParametersandqualifiers(ctx *ParametersandqualifiersContext) {}

// ExitParametersandqualifiers is called when production parametersandqualifiers is exited.
func (s *BaseCPP14Listener) ExitParametersandqualifiers(ctx *ParametersandqualifiersContext) {}

// EnterTrailingreturntype is called when production trailingreturntype is entered.
func (s *BaseCPP14Listener) EnterTrailingreturntype(ctx *TrailingreturntypeContext) {}

// ExitTrailingreturntype is called when production trailingreturntype is exited.
func (s *BaseCPP14Listener) ExitTrailingreturntype(ctx *TrailingreturntypeContext) {}

// EnterPtroperator is called when production ptroperator is entered.
func (s *BaseCPP14Listener) EnterPtroperator(ctx *PtroperatorContext) {}

// ExitPtroperator is called when production ptroperator is exited.
func (s *BaseCPP14Listener) ExitPtroperator(ctx *PtroperatorContext) {}

// EnterCvqualifierseq is called when production cvqualifierseq is entered.
func (s *BaseCPP14Listener) EnterCvqualifierseq(ctx *CvqualifierseqContext) {}

// ExitCvqualifierseq is called when production cvqualifierseq is exited.
func (s *BaseCPP14Listener) ExitCvqualifierseq(ctx *CvqualifierseqContext) {}

// EnterCvqualifier is called when production cvqualifier is entered.
func (s *BaseCPP14Listener) EnterCvqualifier(ctx *CvqualifierContext) {}

// ExitCvqualifier is called when production cvqualifier is exited.
func (s *BaseCPP14Listener) ExitCvqualifier(ctx *CvqualifierContext) {}

// EnterRefqualifier is called when production refqualifier is entered.
func (s *BaseCPP14Listener) EnterRefqualifier(ctx *RefqualifierContext) {}

// ExitRefqualifier is called when production refqualifier is exited.
func (s *BaseCPP14Listener) ExitRefqualifier(ctx *RefqualifierContext) {}

// EnterDeclaratorid is called when production declaratorid is entered.
func (s *BaseCPP14Listener) EnterDeclaratorid(ctx *DeclaratoridContext) {}

// ExitDeclaratorid is called when production declaratorid is exited.
func (s *BaseCPP14Listener) ExitDeclaratorid(ctx *DeclaratoridContext) {}

// EnterThetypeid is called when production thetypeid is entered.
func (s *BaseCPP14Listener) EnterThetypeid(ctx *ThetypeidContext) {}

// ExitThetypeid is called when production thetypeid is exited.
func (s *BaseCPP14Listener) ExitThetypeid(ctx *ThetypeidContext) {}

// EnterAbstractdeclarator is called when production abstractdeclarator is entered.
func (s *BaseCPP14Listener) EnterAbstractdeclarator(ctx *AbstractdeclaratorContext) {}

// ExitAbstractdeclarator is called when production abstractdeclarator is exited.
func (s *BaseCPP14Listener) ExitAbstractdeclarator(ctx *AbstractdeclaratorContext) {}

// EnterPtrabstractdeclarator is called when production ptrabstractdeclarator is entered.
func (s *BaseCPP14Listener) EnterPtrabstractdeclarator(ctx *PtrabstractdeclaratorContext) {}

// ExitPtrabstractdeclarator is called when production ptrabstractdeclarator is exited.
func (s *BaseCPP14Listener) ExitPtrabstractdeclarator(ctx *PtrabstractdeclaratorContext) {}

// EnterNoptrabstractdeclarator is called when production noptrabstractdeclarator is entered.
func (s *BaseCPP14Listener) EnterNoptrabstractdeclarator(ctx *NoptrabstractdeclaratorContext) {}

// ExitNoptrabstractdeclarator is called when production noptrabstractdeclarator is exited.
func (s *BaseCPP14Listener) ExitNoptrabstractdeclarator(ctx *NoptrabstractdeclaratorContext) {}

// EnterAbstractpackdeclarator is called when production abstractpackdeclarator is entered.
func (s *BaseCPP14Listener) EnterAbstractpackdeclarator(ctx *AbstractpackdeclaratorContext) {}

// ExitAbstractpackdeclarator is called when production abstractpackdeclarator is exited.
func (s *BaseCPP14Listener) ExitAbstractpackdeclarator(ctx *AbstractpackdeclaratorContext) {}

// EnterNoptrabstractpackdeclarator is called when production noptrabstractpackdeclarator is entered.
func (s *BaseCPP14Listener) EnterNoptrabstractpackdeclarator(ctx *NoptrabstractpackdeclaratorContext) {
}

// ExitNoptrabstractpackdeclarator is called when production noptrabstractpackdeclarator is exited.
func (s *BaseCPP14Listener) ExitNoptrabstractpackdeclarator(ctx *NoptrabstractpackdeclaratorContext) {}

// EnterParameterdeclarationclause is called when production parameterdeclarationclause is entered.
func (s *BaseCPP14Listener) EnterParameterdeclarationclause(ctx *ParameterdeclarationclauseContext) {}

// ExitParameterdeclarationclause is called when production parameterdeclarationclause is exited.
func (s *BaseCPP14Listener) ExitParameterdeclarationclause(ctx *ParameterdeclarationclauseContext) {}

// EnterParameterdeclarationlist is called when production parameterdeclarationlist is entered.
func (s *BaseCPP14Listener) EnterParameterdeclarationlist(ctx *ParameterdeclarationlistContext) {}

// ExitParameterdeclarationlist is called when production parameterdeclarationlist is exited.
func (s *BaseCPP14Listener) ExitParameterdeclarationlist(ctx *ParameterdeclarationlistContext) {}

// EnterParameterdeclaration is called when production parameterdeclaration is entered.
func (s *BaseCPP14Listener) EnterParameterdeclaration(ctx *ParameterdeclarationContext) {}

// ExitParameterdeclaration is called when production parameterdeclaration is exited.
func (s *BaseCPP14Listener) ExitParameterdeclaration(ctx *ParameterdeclarationContext) {}

// EnterFunctiondefinition is called when production functiondefinition is entered.
func (s *BaseCPP14Listener) EnterFunctiondefinition(ctx *FunctiondefinitionContext) {}

// ExitFunctiondefinition is called when production functiondefinition is exited.
func (s *BaseCPP14Listener) ExitFunctiondefinition(ctx *FunctiondefinitionContext) {}

// EnterFunctionbody is called when production functionbody is entered.
func (s *BaseCPP14Listener) EnterFunctionbody(ctx *FunctionbodyContext) {}

// ExitFunctionbody is called when production functionbody is exited.
func (s *BaseCPP14Listener) ExitFunctionbody(ctx *FunctionbodyContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BaseCPP14Listener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BaseCPP14Listener) ExitInitializer(ctx *InitializerContext) {}

// EnterBraceorequalinitializer is called when production braceorequalinitializer is entered.
func (s *BaseCPP14Listener) EnterBraceorequalinitializer(ctx *BraceorequalinitializerContext) {}

// ExitBraceorequalinitializer is called when production braceorequalinitializer is exited.
func (s *BaseCPP14Listener) ExitBraceorequalinitializer(ctx *BraceorequalinitializerContext) {}

// EnterInitializerclause is called when production initializerclause is entered.
func (s *BaseCPP14Listener) EnterInitializerclause(ctx *InitializerclauseContext) {}

// ExitInitializerclause is called when production initializerclause is exited.
func (s *BaseCPP14Listener) ExitInitializerclause(ctx *InitializerclauseContext) {}

// EnterInitializerlist is called when production initializerlist is entered.
func (s *BaseCPP14Listener) EnterInitializerlist(ctx *InitializerlistContext) {}

// ExitInitializerlist is called when production initializerlist is exited.
func (s *BaseCPP14Listener) ExitInitializerlist(ctx *InitializerlistContext) {}

// EnterBracedinitlist is called when production bracedinitlist is entered.
func (s *BaseCPP14Listener) EnterBracedinitlist(ctx *BracedinitlistContext) {}

// ExitBracedinitlist is called when production bracedinitlist is exited.
func (s *BaseCPP14Listener) ExitBracedinitlist(ctx *BracedinitlistContext) {}

// EnterClassname is called when production classname is entered.
func (s *BaseCPP14Listener) EnterClassname(ctx *ClassnameContext) {}

// ExitClassname is called when production classname is exited.
func (s *BaseCPP14Listener) ExitClassname(ctx *ClassnameContext) {}

// EnterClassspecifier is called when production classspecifier is entered.
func (s *BaseCPP14Listener) EnterClassspecifier(ctx *ClassspecifierContext) {}

// ExitClassspecifier is called when production classspecifier is exited.
func (s *BaseCPP14Listener) ExitClassspecifier(ctx *ClassspecifierContext) {}

// EnterClasshead is called when production classhead is entered.
func (s *BaseCPP14Listener) EnterClasshead(ctx *ClassheadContext) {}

// ExitClasshead is called when production classhead is exited.
func (s *BaseCPP14Listener) ExitClasshead(ctx *ClassheadContext) {}

// EnterClassheadname is called when production classheadname is entered.
func (s *BaseCPP14Listener) EnterClassheadname(ctx *ClassheadnameContext) {}

// ExitClassheadname is called when production classheadname is exited.
func (s *BaseCPP14Listener) ExitClassheadname(ctx *ClassheadnameContext) {}

// EnterClassvirtspecifier is called when production classvirtspecifier is entered.
func (s *BaseCPP14Listener) EnterClassvirtspecifier(ctx *ClassvirtspecifierContext) {}

// ExitClassvirtspecifier is called when production classvirtspecifier is exited.
func (s *BaseCPP14Listener) ExitClassvirtspecifier(ctx *ClassvirtspecifierContext) {}

// EnterClasskey is called when production classkey is entered.
func (s *BaseCPP14Listener) EnterClasskey(ctx *ClasskeyContext) {}

// ExitClasskey is called when production classkey is exited.
func (s *BaseCPP14Listener) ExitClasskey(ctx *ClasskeyContext) {}

// EnterMemberspecification is called when production memberspecification is entered.
func (s *BaseCPP14Listener) EnterMemberspecification(ctx *MemberspecificationContext) {}

// ExitMemberspecification is called when production memberspecification is exited.
func (s *BaseCPP14Listener) ExitMemberspecification(ctx *MemberspecificationContext) {}

// EnterMemberdeclaration is called when production memberdeclaration is entered.
func (s *BaseCPP14Listener) EnterMemberdeclaration(ctx *MemberdeclarationContext) {}

// ExitMemberdeclaration is called when production memberdeclaration is exited.
func (s *BaseCPP14Listener) ExitMemberdeclaration(ctx *MemberdeclarationContext) {}

// EnterMemberdeclaratorlist is called when production memberdeclaratorlist is entered.
func (s *BaseCPP14Listener) EnterMemberdeclaratorlist(ctx *MemberdeclaratorlistContext) {}

// ExitMemberdeclaratorlist is called when production memberdeclaratorlist is exited.
func (s *BaseCPP14Listener) ExitMemberdeclaratorlist(ctx *MemberdeclaratorlistContext) {}

// EnterMemberdeclarator is called when production memberdeclarator is entered.
func (s *BaseCPP14Listener) EnterMemberdeclarator(ctx *MemberdeclaratorContext) {}

// ExitMemberdeclarator is called when production memberdeclarator is exited.
func (s *BaseCPP14Listener) ExitMemberdeclarator(ctx *MemberdeclaratorContext) {}

// EnterVirtspecifierseq is called when production virtspecifierseq is entered.
func (s *BaseCPP14Listener) EnterVirtspecifierseq(ctx *VirtspecifierseqContext) {}

// ExitVirtspecifierseq is called when production virtspecifierseq is exited.
func (s *BaseCPP14Listener) ExitVirtspecifierseq(ctx *VirtspecifierseqContext) {}

// EnterVirtspecifier is called when production virtspecifier is entered.
func (s *BaseCPP14Listener) EnterVirtspecifier(ctx *VirtspecifierContext) {}

// ExitVirtspecifier is called when production virtspecifier is exited.
func (s *BaseCPP14Listener) ExitVirtspecifier(ctx *VirtspecifierContext) {}

// EnterPurespecifier is called when production purespecifier is entered.
func (s *BaseCPP14Listener) EnterPurespecifier(ctx *PurespecifierContext) {}

// ExitPurespecifier is called when production purespecifier is exited.
func (s *BaseCPP14Listener) ExitPurespecifier(ctx *PurespecifierContext) {}

// EnterBaseclause is called when production baseclause is entered.
func (s *BaseCPP14Listener) EnterBaseclause(ctx *BaseclauseContext) {}

// ExitBaseclause is called when production baseclause is exited.
func (s *BaseCPP14Listener) ExitBaseclause(ctx *BaseclauseContext) {}

// EnterBasespecifierlist is called when production basespecifierlist is entered.
func (s *BaseCPP14Listener) EnterBasespecifierlist(ctx *BasespecifierlistContext) {}

// ExitBasespecifierlist is called when production basespecifierlist is exited.
func (s *BaseCPP14Listener) ExitBasespecifierlist(ctx *BasespecifierlistContext) {}

// EnterBasespecifier is called when production basespecifier is entered.
func (s *BaseCPP14Listener) EnterBasespecifier(ctx *BasespecifierContext) {}

// ExitBasespecifier is called when production basespecifier is exited.
func (s *BaseCPP14Listener) ExitBasespecifier(ctx *BasespecifierContext) {}

// EnterClassordecltype is called when production classordecltype is entered.
func (s *BaseCPP14Listener) EnterClassordecltype(ctx *ClassordecltypeContext) {}

// ExitClassordecltype is called when production classordecltype is exited.
func (s *BaseCPP14Listener) ExitClassordecltype(ctx *ClassordecltypeContext) {}

// EnterBasetypespecifier is called when production basetypespecifier is entered.
func (s *BaseCPP14Listener) EnterBasetypespecifier(ctx *BasetypespecifierContext) {}

// ExitBasetypespecifier is called when production basetypespecifier is exited.
func (s *BaseCPP14Listener) ExitBasetypespecifier(ctx *BasetypespecifierContext) {}

// EnterAccessspecifier is called when production accessspecifier is entered.
func (s *BaseCPP14Listener) EnterAccessspecifier(ctx *AccessspecifierContext) {}

// ExitAccessspecifier is called when production accessspecifier is exited.
func (s *BaseCPP14Listener) ExitAccessspecifier(ctx *AccessspecifierContext) {}

// EnterConversionfunctionid is called when production conversionfunctionid is entered.
func (s *BaseCPP14Listener) EnterConversionfunctionid(ctx *ConversionfunctionidContext) {}

// ExitConversionfunctionid is called when production conversionfunctionid is exited.
func (s *BaseCPP14Listener) ExitConversionfunctionid(ctx *ConversionfunctionidContext) {}

// EnterConversiontypeid is called when production conversiontypeid is entered.
func (s *BaseCPP14Listener) EnterConversiontypeid(ctx *ConversiontypeidContext) {}

// ExitConversiontypeid is called when production conversiontypeid is exited.
func (s *BaseCPP14Listener) ExitConversiontypeid(ctx *ConversiontypeidContext) {}

// EnterConversiondeclarator is called when production conversiondeclarator is entered.
func (s *BaseCPP14Listener) EnterConversiondeclarator(ctx *ConversiondeclaratorContext) {}

// ExitConversiondeclarator is called when production conversiondeclarator is exited.
func (s *BaseCPP14Listener) ExitConversiondeclarator(ctx *ConversiondeclaratorContext) {}

// EnterCtorinitializer is called when production ctorinitializer is entered.
func (s *BaseCPP14Listener) EnterCtorinitializer(ctx *CtorinitializerContext) {}

// ExitCtorinitializer is called when production ctorinitializer is exited.
func (s *BaseCPP14Listener) ExitCtorinitializer(ctx *CtorinitializerContext) {}

// EnterMeminitializerlist is called when production meminitializerlist is entered.
func (s *BaseCPP14Listener) EnterMeminitializerlist(ctx *MeminitializerlistContext) {}

// ExitMeminitializerlist is called when production meminitializerlist is exited.
func (s *BaseCPP14Listener) ExitMeminitializerlist(ctx *MeminitializerlistContext) {}

// EnterMeminitializer is called when production meminitializer is entered.
func (s *BaseCPP14Listener) EnterMeminitializer(ctx *MeminitializerContext) {}

// ExitMeminitializer is called when production meminitializer is exited.
func (s *BaseCPP14Listener) ExitMeminitializer(ctx *MeminitializerContext) {}

// EnterMeminitializerid is called when production meminitializerid is entered.
func (s *BaseCPP14Listener) EnterMeminitializerid(ctx *MeminitializeridContext) {}

// ExitMeminitializerid is called when production meminitializerid is exited.
func (s *BaseCPP14Listener) ExitMeminitializerid(ctx *MeminitializeridContext) {}

// EnterOperatorfunctionid is called when production operatorfunctionid is entered.
func (s *BaseCPP14Listener) EnterOperatorfunctionid(ctx *OperatorfunctionidContext) {}

// ExitOperatorfunctionid is called when production operatorfunctionid is exited.
func (s *BaseCPP14Listener) ExitOperatorfunctionid(ctx *OperatorfunctionidContext) {}

// EnterLiteraloperatorid is called when production literaloperatorid is entered.
func (s *BaseCPP14Listener) EnterLiteraloperatorid(ctx *LiteraloperatoridContext) {}

// ExitLiteraloperatorid is called when production literaloperatorid is exited.
func (s *BaseCPP14Listener) ExitLiteraloperatorid(ctx *LiteraloperatoridContext) {}

// EnterTemplatedeclaration is called when production templatedeclaration is entered.
func (s *BaseCPP14Listener) EnterTemplatedeclaration(ctx *TemplatedeclarationContext) {}

// ExitTemplatedeclaration is called when production templatedeclaration is exited.
func (s *BaseCPP14Listener) ExitTemplatedeclaration(ctx *TemplatedeclarationContext) {}

// EnterTemplateparameterlist is called when production templateparameterlist is entered.
func (s *BaseCPP14Listener) EnterTemplateparameterlist(ctx *TemplateparameterlistContext) {}

// ExitTemplateparameterlist is called when production templateparameterlist is exited.
func (s *BaseCPP14Listener) ExitTemplateparameterlist(ctx *TemplateparameterlistContext) {}

// EnterTemplateparameter is called when production templateparameter is entered.
func (s *BaseCPP14Listener) EnterTemplateparameter(ctx *TemplateparameterContext) {}

// ExitTemplateparameter is called when production templateparameter is exited.
func (s *BaseCPP14Listener) ExitTemplateparameter(ctx *TemplateparameterContext) {}

// EnterTypeparameter is called when production typeparameter is entered.
func (s *BaseCPP14Listener) EnterTypeparameter(ctx *TypeparameterContext) {}

// ExitTypeparameter is called when production typeparameter is exited.
func (s *BaseCPP14Listener) ExitTypeparameter(ctx *TypeparameterContext) {}

// EnterSimpletemplateid is called when production simpletemplateid is entered.
func (s *BaseCPP14Listener) EnterSimpletemplateid(ctx *SimpletemplateidContext) {}

// ExitSimpletemplateid is called when production simpletemplateid is exited.
func (s *BaseCPP14Listener) ExitSimpletemplateid(ctx *SimpletemplateidContext) {}

// EnterTemplateid is called when production templateid is entered.
func (s *BaseCPP14Listener) EnterTemplateid(ctx *TemplateidContext) {}

// ExitTemplateid is called when production templateid is exited.
func (s *BaseCPP14Listener) ExitTemplateid(ctx *TemplateidContext) {}

// EnterTemplatename is called when production templatename is entered.
func (s *BaseCPP14Listener) EnterTemplatename(ctx *TemplatenameContext) {}

// ExitTemplatename is called when production templatename is exited.
func (s *BaseCPP14Listener) ExitTemplatename(ctx *TemplatenameContext) {}

// EnterTemplateargumentlist is called when production templateargumentlist is entered.
func (s *BaseCPP14Listener) EnterTemplateargumentlist(ctx *TemplateargumentlistContext) {}

// ExitTemplateargumentlist is called when production templateargumentlist is exited.
func (s *BaseCPP14Listener) ExitTemplateargumentlist(ctx *TemplateargumentlistContext) {}

// EnterTemplateargument is called when production templateargument is entered.
func (s *BaseCPP14Listener) EnterTemplateargument(ctx *TemplateargumentContext) {}

// ExitTemplateargument is called when production templateargument is exited.
func (s *BaseCPP14Listener) ExitTemplateargument(ctx *TemplateargumentContext) {}

// EnterTypenamespecifier is called when production typenamespecifier is entered.
func (s *BaseCPP14Listener) EnterTypenamespecifier(ctx *TypenamespecifierContext) {}

// ExitTypenamespecifier is called when production typenamespecifier is exited.
func (s *BaseCPP14Listener) ExitTypenamespecifier(ctx *TypenamespecifierContext) {}

// EnterExplicitinstantiation is called when production explicitinstantiation is entered.
func (s *BaseCPP14Listener) EnterExplicitinstantiation(ctx *ExplicitinstantiationContext) {}

// ExitExplicitinstantiation is called when production explicitinstantiation is exited.
func (s *BaseCPP14Listener) ExitExplicitinstantiation(ctx *ExplicitinstantiationContext) {}

// EnterExplicitspecialization is called when production explicitspecialization is entered.
func (s *BaseCPP14Listener) EnterExplicitspecialization(ctx *ExplicitspecializationContext) {}

// ExitExplicitspecialization is called when production explicitspecialization is exited.
func (s *BaseCPP14Listener) ExitExplicitspecialization(ctx *ExplicitspecializationContext) {}

// EnterTryblock is called when production tryblock is entered.
func (s *BaseCPP14Listener) EnterTryblock(ctx *TryblockContext) {}

// ExitTryblock is called when production tryblock is exited.
func (s *BaseCPP14Listener) ExitTryblock(ctx *TryblockContext) {}

// EnterFunctiontryblock is called when production functiontryblock is entered.
func (s *BaseCPP14Listener) EnterFunctiontryblock(ctx *FunctiontryblockContext) {}

// ExitFunctiontryblock is called when production functiontryblock is exited.
func (s *BaseCPP14Listener) ExitFunctiontryblock(ctx *FunctiontryblockContext) {}

// EnterHandlerseq is called when production handlerseq is entered.
func (s *BaseCPP14Listener) EnterHandlerseq(ctx *HandlerseqContext) {}

// ExitHandlerseq is called when production handlerseq is exited.
func (s *BaseCPP14Listener) ExitHandlerseq(ctx *HandlerseqContext) {}

// EnterHandler is called when production handler is entered.
func (s *BaseCPP14Listener) EnterHandler(ctx *HandlerContext) {}

// ExitHandler is called when production handler is exited.
func (s *BaseCPP14Listener) ExitHandler(ctx *HandlerContext) {}

// EnterExceptiondeclaration is called when production exceptiondeclaration is entered.
func (s *BaseCPP14Listener) EnterExceptiondeclaration(ctx *ExceptiondeclarationContext) {}

// ExitExceptiondeclaration is called when production exceptiondeclaration is exited.
func (s *BaseCPP14Listener) ExitExceptiondeclaration(ctx *ExceptiondeclarationContext) {}

// EnterThrowexpression is called when production throwexpression is entered.
func (s *BaseCPP14Listener) EnterThrowexpression(ctx *ThrowexpressionContext) {}

// ExitThrowexpression is called when production throwexpression is exited.
func (s *BaseCPP14Listener) ExitThrowexpression(ctx *ThrowexpressionContext) {}

// EnterExceptionspecification is called when production exceptionspecification is entered.
func (s *BaseCPP14Listener) EnterExceptionspecification(ctx *ExceptionspecificationContext) {}

// ExitExceptionspecification is called when production exceptionspecification is exited.
func (s *BaseCPP14Listener) ExitExceptionspecification(ctx *ExceptionspecificationContext) {}

// EnterDynamicexceptionspecification is called when production dynamicexceptionspecification is entered.
func (s *BaseCPP14Listener) EnterDynamicexceptionspecification(ctx *DynamicexceptionspecificationContext) {
}

// ExitDynamicexceptionspecification is called when production dynamicexceptionspecification is exited.
func (s *BaseCPP14Listener) ExitDynamicexceptionspecification(ctx *DynamicexceptionspecificationContext) {
}

// EnterTypeidlist is called when production typeidlist is entered.
func (s *BaseCPP14Listener) EnterTypeidlist(ctx *TypeidlistContext) {}

// ExitTypeidlist is called when production typeidlist is exited.
func (s *BaseCPP14Listener) ExitTypeidlist(ctx *TypeidlistContext) {}

// EnterNoexceptspecification is called when production noexceptspecification is entered.
func (s *BaseCPP14Listener) EnterNoexceptspecification(ctx *NoexceptspecificationContext) {}

// ExitNoexceptspecification is called when production noexceptspecification is exited.
func (s *BaseCPP14Listener) ExitNoexceptspecification(ctx *NoexceptspecificationContext) {}

// EnterRightShift is called when production rightShift is entered.
func (s *BaseCPP14Listener) EnterRightShift(ctx *RightShiftContext) {}

// ExitRightShift is called when production rightShift is exited.
func (s *BaseCPP14Listener) ExitRightShift(ctx *RightShiftContext) {}

// EnterRightShiftAssign is called when production rightShiftAssign is entered.
func (s *BaseCPP14Listener) EnterRightShiftAssign(ctx *RightShiftAssignContext) {}

// ExitRightShiftAssign is called when production rightShiftAssign is exited.
func (s *BaseCPP14Listener) ExitRightShiftAssign(ctx *RightShiftAssignContext) {}

// EnterTheoperator is called when production theoperator is entered.
func (s *BaseCPP14Listener) EnterTheoperator(ctx *TheoperatorContext) {}

// ExitTheoperator is called when production theoperator is exited.
func (s *BaseCPP14Listener) ExitTheoperator(ctx *TheoperatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCPP14Listener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCPP14Listener) ExitLiteral(ctx *LiteralContext) {}

// EnterBooleanliteral is called when production booleanliteral is entered.
func (s *BaseCPP14Listener) EnterBooleanliteral(ctx *BooleanliteralContext) {}

// ExitBooleanliteral is called when production booleanliteral is exited.
func (s *BaseCPP14Listener) ExitBooleanliteral(ctx *BooleanliteralContext) {}

// EnterPointerliteral is called when production pointerliteral is entered.
func (s *BaseCPP14Listener) EnterPointerliteral(ctx *PointerliteralContext) {}

// ExitPointerliteral is called when production pointerliteral is exited.
func (s *BaseCPP14Listener) ExitPointerliteral(ctx *PointerliteralContext) {}

// EnterUserdefinedliteral is called when production userdefinedliteral is entered.
func (s *BaseCPP14Listener) EnterUserdefinedliteral(ctx *UserdefinedliteralContext) {}

// ExitUserdefinedliteral is called when production userdefinedliteral is exited.
func (s *BaseCPP14Listener) ExitUserdefinedliteral(ctx *UserdefinedliteralContext) {}
