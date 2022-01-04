// Code generated from vba.g4 by ANTLR 4.9.3. DO NOT EDIT.

package vba // vba
import "github.com/antlr/antlr4/runtime/Go/antlr"

// vbaListener is a complete listener for a parse tree produced by vbaParser.
type vbaListener interface {
	antlr.ParseTreeListener

	// EnterStartRule is called when entering the startRule production.
	EnterStartRule(c *StartRuleContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterModuleHeader is called when entering the moduleHeader production.
	EnterModuleHeader(c *ModuleHeaderContext)

	// EnterModuleConfig is called when entering the moduleConfig production.
	EnterModuleConfig(c *ModuleConfigContext)

	// EnterModuleConfigElement is called when entering the moduleConfigElement production.
	EnterModuleConfigElement(c *ModuleConfigElementContext)

	// EnterModuleAttributes is called when entering the moduleAttributes production.
	EnterModuleAttributes(c *ModuleAttributesContext)

	// EnterModuleDeclarations is called when entering the moduleDeclarations production.
	EnterModuleDeclarations(c *ModuleDeclarationsContext)

	// EnterOptionBaseStmt is called when entering the optionBaseStmt production.
	EnterOptionBaseStmt(c *OptionBaseStmtContext)

	// EnterOptionCompareStmt is called when entering the optionCompareStmt production.
	EnterOptionCompareStmt(c *OptionCompareStmtContext)

	// EnterOptionExplicitStmt is called when entering the optionExplicitStmt production.
	EnterOptionExplicitStmt(c *OptionExplicitStmtContext)

	// EnterOptionPrivateModuleStmt is called when entering the optionPrivateModuleStmt production.
	EnterOptionPrivateModuleStmt(c *OptionPrivateModuleStmtContext)

	// EnterModuleDeclarationsElement is called when entering the moduleDeclarationsElement production.
	EnterModuleDeclarationsElement(c *ModuleDeclarationsElementContext)

	// EnterMacroStmt is called when entering the macroStmt production.
	EnterMacroStmt(c *MacroStmtContext)

	// EnterModuleBody is called when entering the moduleBody production.
	EnterModuleBody(c *ModuleBodyContext)

	// EnterModuleBodyElement is called when entering the moduleBodyElement production.
	EnterModuleBodyElement(c *ModuleBodyElementContext)

	// EnterAttributeStmt is called when entering the attributeStmt production.
	EnterAttributeStmt(c *AttributeStmtContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStmt is called when entering the blockStmt production.
	EnterBlockStmt(c *BlockStmtContext)

	// EnterAppactivateStmt is called when entering the appactivateStmt production.
	EnterAppactivateStmt(c *AppactivateStmtContext)

	// EnterBeepStmt is called when entering the beepStmt production.
	EnterBeepStmt(c *BeepStmtContext)

	// EnterChdirStmt is called when entering the chdirStmt production.
	EnterChdirStmt(c *ChdirStmtContext)

	// EnterChdriveStmt is called when entering the chdriveStmt production.
	EnterChdriveStmt(c *ChdriveStmtContext)

	// EnterCloseStmt is called when entering the closeStmt production.
	EnterCloseStmt(c *CloseStmtContext)

	// EnterConstStmt is called when entering the constStmt production.
	EnterConstStmt(c *ConstStmtContext)

	// EnterConstSubStmt is called when entering the constSubStmt production.
	EnterConstSubStmt(c *ConstSubStmtContext)

	// EnterDateStmt is called when entering the dateStmt production.
	EnterDateStmt(c *DateStmtContext)

	// EnterDeclareStmt is called when entering the declareStmt production.
	EnterDeclareStmt(c *DeclareStmtContext)

	// EnterDeftypeStmt is called when entering the deftypeStmt production.
	EnterDeftypeStmt(c *DeftypeStmtContext)

	// EnterDeleteSettingStmt is called when entering the deleteSettingStmt production.
	EnterDeleteSettingStmt(c *DeleteSettingStmtContext)

	// EnterDoLoopStmt is called when entering the doLoopStmt production.
	EnterDoLoopStmt(c *DoLoopStmtContext)

	// EnterEndStmt is called when entering the endStmt production.
	EnterEndStmt(c *EndStmtContext)

	// EnterEnumerationStmt is called when entering the enumerationStmt production.
	EnterEnumerationStmt(c *EnumerationStmtContext)

	// EnterEnumerationStmt_Constant is called when entering the enumerationStmt_Constant production.
	EnterEnumerationStmt_Constant(c *EnumerationStmt_ConstantContext)

	// EnterEraseStmt is called when entering the eraseStmt production.
	EnterEraseStmt(c *EraseStmtContext)

	// EnterErrorStmt is called when entering the errorStmt production.
	EnterErrorStmt(c *ErrorStmtContext)

	// EnterEventStmt is called when entering the eventStmt production.
	EnterEventStmt(c *EventStmtContext)

	// EnterExitStmt is called when entering the exitStmt production.
	EnterExitStmt(c *ExitStmtContext)

	// EnterFilecopyStmt is called when entering the filecopyStmt production.
	EnterFilecopyStmt(c *FilecopyStmtContext)

	// EnterForEachStmt is called when entering the forEachStmt production.
	EnterForEachStmt(c *ForEachStmtContext)

	// EnterForNextStmt is called when entering the forNextStmt production.
	EnterForNextStmt(c *ForNextStmtContext)

	// EnterFunctionStmt is called when entering the functionStmt production.
	EnterFunctionStmt(c *FunctionStmtContext)

	// EnterGetStmt is called when entering the getStmt production.
	EnterGetStmt(c *GetStmtContext)

	// EnterGoSubStmt is called when entering the goSubStmt production.
	EnterGoSubStmt(c *GoSubStmtContext)

	// EnterGoToStmt is called when entering the goToStmt production.
	EnterGoToStmt(c *GoToStmtContext)

	// EnterInlineIfThenElse is called when entering the inlineIfThenElse production.
	EnterInlineIfThenElse(c *InlineIfThenElseContext)

	// EnterBlockIfThenElse is called when entering the blockIfThenElse production.
	EnterBlockIfThenElse(c *BlockIfThenElseContext)

	// EnterIfBlockStmt is called when entering the ifBlockStmt production.
	EnterIfBlockStmt(c *IfBlockStmtContext)

	// EnterIfConditionStmt is called when entering the ifConditionStmt production.
	EnterIfConditionStmt(c *IfConditionStmtContext)

	// EnterIfElseIfBlockStmt is called when entering the ifElseIfBlockStmt production.
	EnterIfElseIfBlockStmt(c *IfElseIfBlockStmtContext)

	// EnterIfElseBlockStmt is called when entering the ifElseBlockStmt production.
	EnterIfElseBlockStmt(c *IfElseBlockStmtContext)

	// EnterImplementsStmt is called when entering the implementsStmt production.
	EnterImplementsStmt(c *ImplementsStmtContext)

	// EnterInputStmt is called when entering the inputStmt production.
	EnterInputStmt(c *InputStmtContext)

	// EnterKillStmt is called when entering the killStmt production.
	EnterKillStmt(c *KillStmtContext)

	// EnterLetStmt is called when entering the letStmt production.
	EnterLetStmt(c *LetStmtContext)

	// EnterLineInputStmt is called when entering the lineInputStmt production.
	EnterLineInputStmt(c *LineInputStmtContext)

	// EnterLoadStmt is called when entering the loadStmt production.
	EnterLoadStmt(c *LoadStmtContext)

	// EnterLockStmt is called when entering the lockStmt production.
	EnterLockStmt(c *LockStmtContext)

	// EnterLsetStmt is called when entering the lsetStmt production.
	EnterLsetStmt(c *LsetStmtContext)

	// EnterMacroConstStmt is called when entering the macroConstStmt production.
	EnterMacroConstStmt(c *MacroConstStmtContext)

	// EnterMacroIfThenElseStmt is called when entering the macroIfThenElseStmt production.
	EnterMacroIfThenElseStmt(c *MacroIfThenElseStmtContext)

	// EnterMacroIfBlockStmt is called when entering the macroIfBlockStmt production.
	EnterMacroIfBlockStmt(c *MacroIfBlockStmtContext)

	// EnterMacroElseIfBlockStmt is called when entering the macroElseIfBlockStmt production.
	EnterMacroElseIfBlockStmt(c *MacroElseIfBlockStmtContext)

	// EnterMacroElseBlockStmt is called when entering the macroElseBlockStmt production.
	EnterMacroElseBlockStmt(c *MacroElseBlockStmtContext)

	// EnterMidStmt is called when entering the midStmt production.
	EnterMidStmt(c *MidStmtContext)

	// EnterMkdirStmt is called when entering the mkdirStmt production.
	EnterMkdirStmt(c *MkdirStmtContext)

	// EnterNameStmt is called when entering the nameStmt production.
	EnterNameStmt(c *NameStmtContext)

	// EnterOnErrorStmt is called when entering the onErrorStmt production.
	EnterOnErrorStmt(c *OnErrorStmtContext)

	// EnterOnGoToStmt is called when entering the onGoToStmt production.
	EnterOnGoToStmt(c *OnGoToStmtContext)

	// EnterOnGoSubStmt is called when entering the onGoSubStmt production.
	EnterOnGoSubStmt(c *OnGoSubStmtContext)

	// EnterOpenStmt is called when entering the openStmt production.
	EnterOpenStmt(c *OpenStmtContext)

	// EnterOutputList is called when entering the outputList production.
	EnterOutputList(c *OutputListContext)

	// EnterOutputList_Expression is called when entering the outputList_Expression production.
	EnterOutputList_Expression(c *OutputList_ExpressionContext)

	// EnterPrintStmt is called when entering the printStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterPropertyGetStmt is called when entering the propertyGetStmt production.
	EnterPropertyGetStmt(c *PropertyGetStmtContext)

	// EnterPropertySetStmt is called when entering the propertySetStmt production.
	EnterPropertySetStmt(c *PropertySetStmtContext)

	// EnterPropertyLetStmt is called when entering the propertyLetStmt production.
	EnterPropertyLetStmt(c *PropertyLetStmtContext)

	// EnterPutStmt is called when entering the putStmt production.
	EnterPutStmt(c *PutStmtContext)

	// EnterRaiseEventStmt is called when entering the raiseEventStmt production.
	EnterRaiseEventStmt(c *RaiseEventStmtContext)

	// EnterRandomizeStmt is called when entering the randomizeStmt production.
	EnterRandomizeStmt(c *RandomizeStmtContext)

	// EnterRedimStmt is called when entering the redimStmt production.
	EnterRedimStmt(c *RedimStmtContext)

	// EnterRedimSubStmt is called when entering the redimSubStmt production.
	EnterRedimSubStmt(c *RedimSubStmtContext)

	// EnterResetStmt is called when entering the resetStmt production.
	EnterResetStmt(c *ResetStmtContext)

	// EnterResumeStmt is called when entering the resumeStmt production.
	EnterResumeStmt(c *ResumeStmtContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterRmdirStmt is called when entering the rmdirStmt production.
	EnterRmdirStmt(c *RmdirStmtContext)

	// EnterRsetStmt is called when entering the rsetStmt production.
	EnterRsetStmt(c *RsetStmtContext)

	// EnterSavepictureStmt is called when entering the savepictureStmt production.
	EnterSavepictureStmt(c *SavepictureStmtContext)

	// EnterSaveSettingStmt is called when entering the saveSettingStmt production.
	EnterSaveSettingStmt(c *SaveSettingStmtContext)

	// EnterSeekStmt is called when entering the seekStmt production.
	EnterSeekStmt(c *SeekStmtContext)

	// EnterSelectCaseStmt is called when entering the selectCaseStmt production.
	EnterSelectCaseStmt(c *SelectCaseStmtContext)

	// EnterCaseCondIs is called when entering the caseCondIs production.
	EnterCaseCondIs(c *CaseCondIsContext)

	// EnterCaseCondTo is called when entering the caseCondTo production.
	EnterCaseCondTo(c *CaseCondToContext)

	// EnterCaseCondValue is called when entering the caseCondValue production.
	EnterCaseCondValue(c *CaseCondValueContext)

	// EnterSC_Case is called when entering the sC_Case production.
	EnterSC_Case(c *SC_CaseContext)

	// EnterCaseCondElse is called when entering the caseCondElse production.
	EnterCaseCondElse(c *CaseCondElseContext)

	// EnterCaseCondSelection is called when entering the caseCondSelection production.
	EnterCaseCondSelection(c *CaseCondSelectionContext)

	// EnterSendkeysStmt is called when entering the sendkeysStmt production.
	EnterSendkeysStmt(c *SendkeysStmtContext)

	// EnterSetattrStmt is called when entering the setattrStmt production.
	EnterSetattrStmt(c *SetattrStmtContext)

	// EnterSetStmt is called when entering the setStmt production.
	EnterSetStmt(c *SetStmtContext)

	// EnterStopStmt is called when entering the stopStmt production.
	EnterStopStmt(c *StopStmtContext)

	// EnterSubStmt is called when entering the subStmt production.
	EnterSubStmt(c *SubStmtContext)

	// EnterTimeStmt is called when entering the timeStmt production.
	EnterTimeStmt(c *TimeStmtContext)

	// EnterTypeStmt is called when entering the typeStmt production.
	EnterTypeStmt(c *TypeStmtContext)

	// EnterTypeStmt_Element is called when entering the typeStmt_Element production.
	EnterTypeStmt_Element(c *TypeStmt_ElementContext)

	// EnterTypeOfStmt is called when entering the typeOfStmt production.
	EnterTypeOfStmt(c *TypeOfStmtContext)

	// EnterUnloadStmt is called when entering the unloadStmt production.
	EnterUnloadStmt(c *UnloadStmtContext)

	// EnterUnlockStmt is called when entering the unlockStmt production.
	EnterUnlockStmt(c *UnlockStmtContext)

	// EnterVsStruct is called when entering the vsStruct production.
	EnterVsStruct(c *VsStructContext)

	// EnterVsAdd is called when entering the vsAdd production.
	EnterVsAdd(c *VsAddContext)

	// EnterVsLt is called when entering the vsLt production.
	EnterVsLt(c *VsLtContext)

	// EnterVsAddressOf is called when entering the vsAddressOf production.
	EnterVsAddressOf(c *VsAddressOfContext)

	// EnterVsNew is called when entering the vsNew production.
	EnterVsNew(c *VsNewContext)

	// EnterVsMult is called when entering the vsMult production.
	EnterVsMult(c *VsMultContext)

	// EnterVsNegation is called when entering the vsNegation production.
	EnterVsNegation(c *VsNegationContext)

	// EnterVsAssign is called when entering the vsAssign production.
	EnterVsAssign(c *VsAssignContext)

	// EnterVsLike is called when entering the vsLike production.
	EnterVsLike(c *VsLikeContext)

	// EnterVsDiv is called when entering the vsDiv production.
	EnterVsDiv(c *VsDivContext)

	// EnterVsPlus is called when entering the vsPlus production.
	EnterVsPlus(c *VsPlusContext)

	// EnterVsNot is called when entering the vsNot production.
	EnterVsNot(c *VsNotContext)

	// EnterVsGeq is called when entering the vsGeq production.
	EnterVsGeq(c *VsGeqContext)

	// EnterVsTypeOf is called when entering the vsTypeOf production.
	EnterVsTypeOf(c *VsTypeOfContext)

	// EnterVsICS is called when entering the vsICS production.
	EnterVsICS(c *VsICSContext)

	// EnterVsNeq is called when entering the vsNeq production.
	EnterVsNeq(c *VsNeqContext)

	// EnterVsXor is called when entering the vsXor production.
	EnterVsXor(c *VsXorContext)

	// EnterVsAnd is called when entering the vsAnd production.
	EnterVsAnd(c *VsAndContext)

	// EnterVsLeq is called when entering the vsLeq production.
	EnterVsLeq(c *VsLeqContext)

	// EnterVsPow is called when entering the vsPow production.
	EnterVsPow(c *VsPowContext)

	// EnterVsIs is called when entering the vsIs production.
	EnterVsIs(c *VsIsContext)

	// EnterVsMod is called when entering the vsMod production.
	EnterVsMod(c *VsModContext)

	// EnterVsAmp is called when entering the vsAmp production.
	EnterVsAmp(c *VsAmpContext)

	// EnterVsOr is called when entering the vsOr production.
	EnterVsOr(c *VsOrContext)

	// EnterVsMinus is called when entering the vsMinus production.
	EnterVsMinus(c *VsMinusContext)

	// EnterVsLiteral is called when entering the vsLiteral production.
	EnterVsLiteral(c *VsLiteralContext)

	// EnterVsEqv is called when entering the vsEqv production.
	EnterVsEqv(c *VsEqvContext)

	// EnterVsImp is called when entering the vsImp production.
	EnterVsImp(c *VsImpContext)

	// EnterVsGt is called when entering the vsGt production.
	EnterVsGt(c *VsGtContext)

	// EnterVsEq is called when entering the vsEq production.
	EnterVsEq(c *VsEqContext)

	// EnterVsMid is called when entering the vsMid production.
	EnterVsMid(c *VsMidContext)

	// EnterVariableStmt is called when entering the variableStmt production.
	EnterVariableStmt(c *VariableStmtContext)

	// EnterVariableListStmt is called when entering the variableListStmt production.
	EnterVariableListStmt(c *VariableListStmtContext)

	// EnterVariableSubStmt is called when entering the variableSubStmt production.
	EnterVariableSubStmt(c *VariableSubStmtContext)

	// EnterWhileWendStmt is called when entering the whileWendStmt production.
	EnterWhileWendStmt(c *WhileWendStmtContext)

	// EnterWidthStmt is called when entering the widthStmt production.
	EnterWidthStmt(c *WidthStmtContext)

	// EnterWithStmt is called when entering the withStmt production.
	EnterWithStmt(c *WithStmtContext)

	// EnterWriteStmt is called when entering the writeStmt production.
	EnterWriteStmt(c *WriteStmtContext)

	// EnterFileNumber is called when entering the fileNumber production.
	EnterFileNumber(c *FileNumberContext)

	// EnterExplicitCallStmt is called when entering the explicitCallStmt production.
	EnterExplicitCallStmt(c *ExplicitCallStmtContext)

	// EnterECS_ProcedureCall is called when entering the eCS_ProcedureCall production.
	EnterECS_ProcedureCall(c *ECS_ProcedureCallContext)

	// EnterECS_MemberProcedureCall is called when entering the eCS_MemberProcedureCall production.
	EnterECS_MemberProcedureCall(c *ECS_MemberProcedureCallContext)

	// EnterImplicitCallStmt_InBlock is called when entering the implicitCallStmt_InBlock production.
	EnterImplicitCallStmt_InBlock(c *ImplicitCallStmt_InBlockContext)

	// EnterICS_B_MemberProcedureCall is called when entering the iCS_B_MemberProcedureCall production.
	EnterICS_B_MemberProcedureCall(c *ICS_B_MemberProcedureCallContext)

	// EnterICS_B_ProcedureCall is called when entering the iCS_B_ProcedureCall production.
	EnterICS_B_ProcedureCall(c *ICS_B_ProcedureCallContext)

	// EnterImplicitCallStmt_InStmt is called when entering the implicitCallStmt_InStmt production.
	EnterImplicitCallStmt_InStmt(c *ImplicitCallStmt_InStmtContext)

	// EnterICS_S_VariableOrProcedureCall is called when entering the iCS_S_VariableOrProcedureCall production.
	EnterICS_S_VariableOrProcedureCall(c *ICS_S_VariableOrProcedureCallContext)

	// EnterICS_S_ProcedureOrArrayCall is called when entering the iCS_S_ProcedureOrArrayCall production.
	EnterICS_S_ProcedureOrArrayCall(c *ICS_S_ProcedureOrArrayCallContext)

	// EnterICS_S_MembersCall is called when entering the iCS_S_MembersCall production.
	EnterICS_S_MembersCall(c *ICS_S_MembersCallContext)

	// EnterICS_S_MemberCall is called when entering the iCS_S_MemberCall production.
	EnterICS_S_MemberCall(c *ICS_S_MemberCallContext)

	// EnterICS_S_DictionaryCall is called when entering the iCS_S_DictionaryCall production.
	EnterICS_S_DictionaryCall(c *ICS_S_DictionaryCallContext)

	// EnterArgsCall is called when entering the argsCall production.
	EnterArgsCall(c *ArgsCallContext)

	// EnterArgCall is called when entering the argCall production.
	EnterArgCall(c *ArgCallContext)

	// EnterDictionaryCallStmt is called when entering the dictionaryCallStmt production.
	EnterDictionaryCallStmt(c *DictionaryCallStmtContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterArgDefaultValue is called when entering the argDefaultValue production.
	EnterArgDefaultValue(c *ArgDefaultValueContext)

	// EnterSubscripts is called when entering the subscripts production.
	EnterSubscripts(c *SubscriptsContext)

	// EnterSubscript_ is called when entering the subscript_ production.
	EnterSubscript_(c *Subscript_Context)

	// EnterAmbiguousIdentifier is called when entering the ambiguousIdentifier production.
	EnterAmbiguousIdentifier(c *AmbiguousIdentifierContext)

	// EnterAsTypeClause is called when entering the asTypeClause production.
	EnterAsTypeClause(c *AsTypeClauseContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// EnterCertainIdentifier is called when entering the certainIdentifier production.
	EnterCertainIdentifier(c *CertainIdentifierContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterComplexType is called when entering the complexType production.
	EnterComplexType(c *ComplexTypeContext)

	// EnterFieldLength is called when entering the fieldLength production.
	EnterFieldLength(c *FieldLengthContext)

	// EnterLetterrange is called when entering the letterrange production.
	EnterLetterrange(c *LetterrangeContext)

	// EnterLineLabel is called when entering the lineLabel production.
	EnterLineLabel(c *LineLabelContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterTypeHint is called when entering the typeHint production.
	EnterTypeHint(c *TypeHintContext)

	// EnterVisibility is called when entering the visibility production.
	EnterVisibility(c *VisibilityContext)

	// EnterAmbiguousKeyword is called when entering the ambiguousKeyword production.
	EnterAmbiguousKeyword(c *AmbiguousKeywordContext)

	// EnterRemComment is called when entering the remComment production.
	EnterRemComment(c *RemCommentContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterEndOfLine is called when entering the endOfLine production.
	EnterEndOfLine(c *EndOfLineContext)

	// EnterEndOfStatement is called when entering the endOfStatement production.
	EnterEndOfStatement(c *EndOfStatementContext)

	// ExitStartRule is called when exiting the startRule production.
	ExitStartRule(c *StartRuleContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitModuleHeader is called when exiting the moduleHeader production.
	ExitModuleHeader(c *ModuleHeaderContext)

	// ExitModuleConfig is called when exiting the moduleConfig production.
	ExitModuleConfig(c *ModuleConfigContext)

	// ExitModuleConfigElement is called when exiting the moduleConfigElement production.
	ExitModuleConfigElement(c *ModuleConfigElementContext)

	// ExitModuleAttributes is called when exiting the moduleAttributes production.
	ExitModuleAttributes(c *ModuleAttributesContext)

	// ExitModuleDeclarations is called when exiting the moduleDeclarations production.
	ExitModuleDeclarations(c *ModuleDeclarationsContext)

	// ExitOptionBaseStmt is called when exiting the optionBaseStmt production.
	ExitOptionBaseStmt(c *OptionBaseStmtContext)

	// ExitOptionCompareStmt is called when exiting the optionCompareStmt production.
	ExitOptionCompareStmt(c *OptionCompareStmtContext)

	// ExitOptionExplicitStmt is called when exiting the optionExplicitStmt production.
	ExitOptionExplicitStmt(c *OptionExplicitStmtContext)

	// ExitOptionPrivateModuleStmt is called when exiting the optionPrivateModuleStmt production.
	ExitOptionPrivateModuleStmt(c *OptionPrivateModuleStmtContext)

	// ExitModuleDeclarationsElement is called when exiting the moduleDeclarationsElement production.
	ExitModuleDeclarationsElement(c *ModuleDeclarationsElementContext)

	// ExitMacroStmt is called when exiting the macroStmt production.
	ExitMacroStmt(c *MacroStmtContext)

	// ExitModuleBody is called when exiting the moduleBody production.
	ExitModuleBody(c *ModuleBodyContext)

	// ExitModuleBodyElement is called when exiting the moduleBodyElement production.
	ExitModuleBodyElement(c *ModuleBodyElementContext)

	// ExitAttributeStmt is called when exiting the attributeStmt production.
	ExitAttributeStmt(c *AttributeStmtContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStmt is called when exiting the blockStmt production.
	ExitBlockStmt(c *BlockStmtContext)

	// ExitAppactivateStmt is called when exiting the appactivateStmt production.
	ExitAppactivateStmt(c *AppactivateStmtContext)

	// ExitBeepStmt is called when exiting the beepStmt production.
	ExitBeepStmt(c *BeepStmtContext)

	// ExitChdirStmt is called when exiting the chdirStmt production.
	ExitChdirStmt(c *ChdirStmtContext)

	// ExitChdriveStmt is called when exiting the chdriveStmt production.
	ExitChdriveStmt(c *ChdriveStmtContext)

	// ExitCloseStmt is called when exiting the closeStmt production.
	ExitCloseStmt(c *CloseStmtContext)

	// ExitConstStmt is called when exiting the constStmt production.
	ExitConstStmt(c *ConstStmtContext)

	// ExitConstSubStmt is called when exiting the constSubStmt production.
	ExitConstSubStmt(c *ConstSubStmtContext)

	// ExitDateStmt is called when exiting the dateStmt production.
	ExitDateStmt(c *DateStmtContext)

	// ExitDeclareStmt is called when exiting the declareStmt production.
	ExitDeclareStmt(c *DeclareStmtContext)

	// ExitDeftypeStmt is called when exiting the deftypeStmt production.
	ExitDeftypeStmt(c *DeftypeStmtContext)

	// ExitDeleteSettingStmt is called when exiting the deleteSettingStmt production.
	ExitDeleteSettingStmt(c *DeleteSettingStmtContext)

	// ExitDoLoopStmt is called when exiting the doLoopStmt production.
	ExitDoLoopStmt(c *DoLoopStmtContext)

	// ExitEndStmt is called when exiting the endStmt production.
	ExitEndStmt(c *EndStmtContext)

	// ExitEnumerationStmt is called when exiting the enumerationStmt production.
	ExitEnumerationStmt(c *EnumerationStmtContext)

	// ExitEnumerationStmt_Constant is called when exiting the enumerationStmt_Constant production.
	ExitEnumerationStmt_Constant(c *EnumerationStmt_ConstantContext)

	// ExitEraseStmt is called when exiting the eraseStmt production.
	ExitEraseStmt(c *EraseStmtContext)

	// ExitErrorStmt is called when exiting the errorStmt production.
	ExitErrorStmt(c *ErrorStmtContext)

	// ExitEventStmt is called when exiting the eventStmt production.
	ExitEventStmt(c *EventStmtContext)

	// ExitExitStmt is called when exiting the exitStmt production.
	ExitExitStmt(c *ExitStmtContext)

	// ExitFilecopyStmt is called when exiting the filecopyStmt production.
	ExitFilecopyStmt(c *FilecopyStmtContext)

	// ExitForEachStmt is called when exiting the forEachStmt production.
	ExitForEachStmt(c *ForEachStmtContext)

	// ExitForNextStmt is called when exiting the forNextStmt production.
	ExitForNextStmt(c *ForNextStmtContext)

	// ExitFunctionStmt is called when exiting the functionStmt production.
	ExitFunctionStmt(c *FunctionStmtContext)

	// ExitGetStmt is called when exiting the getStmt production.
	ExitGetStmt(c *GetStmtContext)

	// ExitGoSubStmt is called when exiting the goSubStmt production.
	ExitGoSubStmt(c *GoSubStmtContext)

	// ExitGoToStmt is called when exiting the goToStmt production.
	ExitGoToStmt(c *GoToStmtContext)

	// ExitInlineIfThenElse is called when exiting the inlineIfThenElse production.
	ExitInlineIfThenElse(c *InlineIfThenElseContext)

	// ExitBlockIfThenElse is called when exiting the blockIfThenElse production.
	ExitBlockIfThenElse(c *BlockIfThenElseContext)

	// ExitIfBlockStmt is called when exiting the ifBlockStmt production.
	ExitIfBlockStmt(c *IfBlockStmtContext)

	// ExitIfConditionStmt is called when exiting the ifConditionStmt production.
	ExitIfConditionStmt(c *IfConditionStmtContext)

	// ExitIfElseIfBlockStmt is called when exiting the ifElseIfBlockStmt production.
	ExitIfElseIfBlockStmt(c *IfElseIfBlockStmtContext)

	// ExitIfElseBlockStmt is called when exiting the ifElseBlockStmt production.
	ExitIfElseBlockStmt(c *IfElseBlockStmtContext)

	// ExitImplementsStmt is called when exiting the implementsStmt production.
	ExitImplementsStmt(c *ImplementsStmtContext)

	// ExitInputStmt is called when exiting the inputStmt production.
	ExitInputStmt(c *InputStmtContext)

	// ExitKillStmt is called when exiting the killStmt production.
	ExitKillStmt(c *KillStmtContext)

	// ExitLetStmt is called when exiting the letStmt production.
	ExitLetStmt(c *LetStmtContext)

	// ExitLineInputStmt is called when exiting the lineInputStmt production.
	ExitLineInputStmt(c *LineInputStmtContext)

	// ExitLoadStmt is called when exiting the loadStmt production.
	ExitLoadStmt(c *LoadStmtContext)

	// ExitLockStmt is called when exiting the lockStmt production.
	ExitLockStmt(c *LockStmtContext)

	// ExitLsetStmt is called when exiting the lsetStmt production.
	ExitLsetStmt(c *LsetStmtContext)

	// ExitMacroConstStmt is called when exiting the macroConstStmt production.
	ExitMacroConstStmt(c *MacroConstStmtContext)

	// ExitMacroIfThenElseStmt is called when exiting the macroIfThenElseStmt production.
	ExitMacroIfThenElseStmt(c *MacroIfThenElseStmtContext)

	// ExitMacroIfBlockStmt is called when exiting the macroIfBlockStmt production.
	ExitMacroIfBlockStmt(c *MacroIfBlockStmtContext)

	// ExitMacroElseIfBlockStmt is called when exiting the macroElseIfBlockStmt production.
	ExitMacroElseIfBlockStmt(c *MacroElseIfBlockStmtContext)

	// ExitMacroElseBlockStmt is called when exiting the macroElseBlockStmt production.
	ExitMacroElseBlockStmt(c *MacroElseBlockStmtContext)

	// ExitMidStmt is called when exiting the midStmt production.
	ExitMidStmt(c *MidStmtContext)

	// ExitMkdirStmt is called when exiting the mkdirStmt production.
	ExitMkdirStmt(c *MkdirStmtContext)

	// ExitNameStmt is called when exiting the nameStmt production.
	ExitNameStmt(c *NameStmtContext)

	// ExitOnErrorStmt is called when exiting the onErrorStmt production.
	ExitOnErrorStmt(c *OnErrorStmtContext)

	// ExitOnGoToStmt is called when exiting the onGoToStmt production.
	ExitOnGoToStmt(c *OnGoToStmtContext)

	// ExitOnGoSubStmt is called when exiting the onGoSubStmt production.
	ExitOnGoSubStmt(c *OnGoSubStmtContext)

	// ExitOpenStmt is called when exiting the openStmt production.
	ExitOpenStmt(c *OpenStmtContext)

	// ExitOutputList is called when exiting the outputList production.
	ExitOutputList(c *OutputListContext)

	// ExitOutputList_Expression is called when exiting the outputList_Expression production.
	ExitOutputList_Expression(c *OutputList_ExpressionContext)

	// ExitPrintStmt is called when exiting the printStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitPropertyGetStmt is called when exiting the propertyGetStmt production.
	ExitPropertyGetStmt(c *PropertyGetStmtContext)

	// ExitPropertySetStmt is called when exiting the propertySetStmt production.
	ExitPropertySetStmt(c *PropertySetStmtContext)

	// ExitPropertyLetStmt is called when exiting the propertyLetStmt production.
	ExitPropertyLetStmt(c *PropertyLetStmtContext)

	// ExitPutStmt is called when exiting the putStmt production.
	ExitPutStmt(c *PutStmtContext)

	// ExitRaiseEventStmt is called when exiting the raiseEventStmt production.
	ExitRaiseEventStmt(c *RaiseEventStmtContext)

	// ExitRandomizeStmt is called when exiting the randomizeStmt production.
	ExitRandomizeStmt(c *RandomizeStmtContext)

	// ExitRedimStmt is called when exiting the redimStmt production.
	ExitRedimStmt(c *RedimStmtContext)

	// ExitRedimSubStmt is called when exiting the redimSubStmt production.
	ExitRedimSubStmt(c *RedimSubStmtContext)

	// ExitResetStmt is called when exiting the resetStmt production.
	ExitResetStmt(c *ResetStmtContext)

	// ExitResumeStmt is called when exiting the resumeStmt production.
	ExitResumeStmt(c *ResumeStmtContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitRmdirStmt is called when exiting the rmdirStmt production.
	ExitRmdirStmt(c *RmdirStmtContext)

	// ExitRsetStmt is called when exiting the rsetStmt production.
	ExitRsetStmt(c *RsetStmtContext)

	// ExitSavepictureStmt is called when exiting the savepictureStmt production.
	ExitSavepictureStmt(c *SavepictureStmtContext)

	// ExitSaveSettingStmt is called when exiting the saveSettingStmt production.
	ExitSaveSettingStmt(c *SaveSettingStmtContext)

	// ExitSeekStmt is called when exiting the seekStmt production.
	ExitSeekStmt(c *SeekStmtContext)

	// ExitSelectCaseStmt is called when exiting the selectCaseStmt production.
	ExitSelectCaseStmt(c *SelectCaseStmtContext)

	// ExitCaseCondIs is called when exiting the caseCondIs production.
	ExitCaseCondIs(c *CaseCondIsContext)

	// ExitCaseCondTo is called when exiting the caseCondTo production.
	ExitCaseCondTo(c *CaseCondToContext)

	// ExitCaseCondValue is called when exiting the caseCondValue production.
	ExitCaseCondValue(c *CaseCondValueContext)

	// ExitSC_Case is called when exiting the sC_Case production.
	ExitSC_Case(c *SC_CaseContext)

	// ExitCaseCondElse is called when exiting the caseCondElse production.
	ExitCaseCondElse(c *CaseCondElseContext)

	// ExitCaseCondSelection is called when exiting the caseCondSelection production.
	ExitCaseCondSelection(c *CaseCondSelectionContext)

	// ExitSendkeysStmt is called when exiting the sendkeysStmt production.
	ExitSendkeysStmt(c *SendkeysStmtContext)

	// ExitSetattrStmt is called when exiting the setattrStmt production.
	ExitSetattrStmt(c *SetattrStmtContext)

	// ExitSetStmt is called when exiting the setStmt production.
	ExitSetStmt(c *SetStmtContext)

	// ExitStopStmt is called when exiting the stopStmt production.
	ExitStopStmt(c *StopStmtContext)

	// ExitSubStmt is called when exiting the subStmt production.
	ExitSubStmt(c *SubStmtContext)

	// ExitTimeStmt is called when exiting the timeStmt production.
	ExitTimeStmt(c *TimeStmtContext)

	// ExitTypeStmt is called when exiting the typeStmt production.
	ExitTypeStmt(c *TypeStmtContext)

	// ExitTypeStmt_Element is called when exiting the typeStmt_Element production.
	ExitTypeStmt_Element(c *TypeStmt_ElementContext)

	// ExitTypeOfStmt is called when exiting the typeOfStmt production.
	ExitTypeOfStmt(c *TypeOfStmtContext)

	// ExitUnloadStmt is called when exiting the unloadStmt production.
	ExitUnloadStmt(c *UnloadStmtContext)

	// ExitUnlockStmt is called when exiting the unlockStmt production.
	ExitUnlockStmt(c *UnlockStmtContext)

	// ExitVsStruct is called when exiting the vsStruct production.
	ExitVsStruct(c *VsStructContext)

	// ExitVsAdd is called when exiting the vsAdd production.
	ExitVsAdd(c *VsAddContext)

	// ExitVsLt is called when exiting the vsLt production.
	ExitVsLt(c *VsLtContext)

	// ExitVsAddressOf is called when exiting the vsAddressOf production.
	ExitVsAddressOf(c *VsAddressOfContext)

	// ExitVsNew is called when exiting the vsNew production.
	ExitVsNew(c *VsNewContext)

	// ExitVsMult is called when exiting the vsMult production.
	ExitVsMult(c *VsMultContext)

	// ExitVsNegation is called when exiting the vsNegation production.
	ExitVsNegation(c *VsNegationContext)

	// ExitVsAssign is called when exiting the vsAssign production.
	ExitVsAssign(c *VsAssignContext)

	// ExitVsLike is called when exiting the vsLike production.
	ExitVsLike(c *VsLikeContext)

	// ExitVsDiv is called when exiting the vsDiv production.
	ExitVsDiv(c *VsDivContext)

	// ExitVsPlus is called when exiting the vsPlus production.
	ExitVsPlus(c *VsPlusContext)

	// ExitVsNot is called when exiting the vsNot production.
	ExitVsNot(c *VsNotContext)

	// ExitVsGeq is called when exiting the vsGeq production.
	ExitVsGeq(c *VsGeqContext)

	// ExitVsTypeOf is called when exiting the vsTypeOf production.
	ExitVsTypeOf(c *VsTypeOfContext)

	// ExitVsICS is called when exiting the vsICS production.
	ExitVsICS(c *VsICSContext)

	// ExitVsNeq is called when exiting the vsNeq production.
	ExitVsNeq(c *VsNeqContext)

	// ExitVsXor is called when exiting the vsXor production.
	ExitVsXor(c *VsXorContext)

	// ExitVsAnd is called when exiting the vsAnd production.
	ExitVsAnd(c *VsAndContext)

	// ExitVsLeq is called when exiting the vsLeq production.
	ExitVsLeq(c *VsLeqContext)

	// ExitVsPow is called when exiting the vsPow production.
	ExitVsPow(c *VsPowContext)

	// ExitVsIs is called when exiting the vsIs production.
	ExitVsIs(c *VsIsContext)

	// ExitVsMod is called when exiting the vsMod production.
	ExitVsMod(c *VsModContext)

	// ExitVsAmp is called when exiting the vsAmp production.
	ExitVsAmp(c *VsAmpContext)

	// ExitVsOr is called when exiting the vsOr production.
	ExitVsOr(c *VsOrContext)

	// ExitVsMinus is called when exiting the vsMinus production.
	ExitVsMinus(c *VsMinusContext)

	// ExitVsLiteral is called when exiting the vsLiteral production.
	ExitVsLiteral(c *VsLiteralContext)

	// ExitVsEqv is called when exiting the vsEqv production.
	ExitVsEqv(c *VsEqvContext)

	// ExitVsImp is called when exiting the vsImp production.
	ExitVsImp(c *VsImpContext)

	// ExitVsGt is called when exiting the vsGt production.
	ExitVsGt(c *VsGtContext)

	// ExitVsEq is called when exiting the vsEq production.
	ExitVsEq(c *VsEqContext)

	// ExitVsMid is called when exiting the vsMid production.
	ExitVsMid(c *VsMidContext)

	// ExitVariableStmt is called when exiting the variableStmt production.
	ExitVariableStmt(c *VariableStmtContext)

	// ExitVariableListStmt is called when exiting the variableListStmt production.
	ExitVariableListStmt(c *VariableListStmtContext)

	// ExitVariableSubStmt is called when exiting the variableSubStmt production.
	ExitVariableSubStmt(c *VariableSubStmtContext)

	// ExitWhileWendStmt is called when exiting the whileWendStmt production.
	ExitWhileWendStmt(c *WhileWendStmtContext)

	// ExitWidthStmt is called when exiting the widthStmt production.
	ExitWidthStmt(c *WidthStmtContext)

	// ExitWithStmt is called when exiting the withStmt production.
	ExitWithStmt(c *WithStmtContext)

	// ExitWriteStmt is called when exiting the writeStmt production.
	ExitWriteStmt(c *WriteStmtContext)

	// ExitFileNumber is called when exiting the fileNumber production.
	ExitFileNumber(c *FileNumberContext)

	// ExitExplicitCallStmt is called when exiting the explicitCallStmt production.
	ExitExplicitCallStmt(c *ExplicitCallStmtContext)

	// ExitECS_ProcedureCall is called when exiting the eCS_ProcedureCall production.
	ExitECS_ProcedureCall(c *ECS_ProcedureCallContext)

	// ExitECS_MemberProcedureCall is called when exiting the eCS_MemberProcedureCall production.
	ExitECS_MemberProcedureCall(c *ECS_MemberProcedureCallContext)

	// ExitImplicitCallStmt_InBlock is called when exiting the implicitCallStmt_InBlock production.
	ExitImplicitCallStmt_InBlock(c *ImplicitCallStmt_InBlockContext)

	// ExitICS_B_MemberProcedureCall is called when exiting the iCS_B_MemberProcedureCall production.
	ExitICS_B_MemberProcedureCall(c *ICS_B_MemberProcedureCallContext)

	// ExitICS_B_ProcedureCall is called when exiting the iCS_B_ProcedureCall production.
	ExitICS_B_ProcedureCall(c *ICS_B_ProcedureCallContext)

	// ExitImplicitCallStmt_InStmt is called when exiting the implicitCallStmt_InStmt production.
	ExitImplicitCallStmt_InStmt(c *ImplicitCallStmt_InStmtContext)

	// ExitICS_S_VariableOrProcedureCall is called when exiting the iCS_S_VariableOrProcedureCall production.
	ExitICS_S_VariableOrProcedureCall(c *ICS_S_VariableOrProcedureCallContext)

	// ExitICS_S_ProcedureOrArrayCall is called when exiting the iCS_S_ProcedureOrArrayCall production.
	ExitICS_S_ProcedureOrArrayCall(c *ICS_S_ProcedureOrArrayCallContext)

	// ExitICS_S_MembersCall is called when exiting the iCS_S_MembersCall production.
	ExitICS_S_MembersCall(c *ICS_S_MembersCallContext)

	// ExitICS_S_MemberCall is called when exiting the iCS_S_MemberCall production.
	ExitICS_S_MemberCall(c *ICS_S_MemberCallContext)

	// ExitICS_S_DictionaryCall is called when exiting the iCS_S_DictionaryCall production.
	ExitICS_S_DictionaryCall(c *ICS_S_DictionaryCallContext)

	// ExitArgsCall is called when exiting the argsCall production.
	ExitArgsCall(c *ArgsCallContext)

	// ExitArgCall is called when exiting the argCall production.
	ExitArgCall(c *ArgCallContext)

	// ExitDictionaryCallStmt is called when exiting the dictionaryCallStmt production.
	ExitDictionaryCallStmt(c *DictionaryCallStmtContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitArgDefaultValue is called when exiting the argDefaultValue production.
	ExitArgDefaultValue(c *ArgDefaultValueContext)

	// ExitSubscripts is called when exiting the subscripts production.
	ExitSubscripts(c *SubscriptsContext)

	// ExitSubscript_ is called when exiting the subscript_ production.
	ExitSubscript_(c *Subscript_Context)

	// ExitAmbiguousIdentifier is called when exiting the ambiguousIdentifier production.
	ExitAmbiguousIdentifier(c *AmbiguousIdentifierContext)

	// ExitAsTypeClause is called when exiting the asTypeClause production.
	ExitAsTypeClause(c *AsTypeClauseContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)

	// ExitCertainIdentifier is called when exiting the certainIdentifier production.
	ExitCertainIdentifier(c *CertainIdentifierContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitComplexType is called when exiting the complexType production.
	ExitComplexType(c *ComplexTypeContext)

	// ExitFieldLength is called when exiting the fieldLength production.
	ExitFieldLength(c *FieldLengthContext)

	// ExitLetterrange is called when exiting the letterrange production.
	ExitLetterrange(c *LetterrangeContext)

	// ExitLineLabel is called when exiting the lineLabel production.
	ExitLineLabel(c *LineLabelContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitTypeHint is called when exiting the typeHint production.
	ExitTypeHint(c *TypeHintContext)

	// ExitVisibility is called when exiting the visibility production.
	ExitVisibility(c *VisibilityContext)

	// ExitAmbiguousKeyword is called when exiting the ambiguousKeyword production.
	ExitAmbiguousKeyword(c *AmbiguousKeywordContext)

	// ExitRemComment is called when exiting the remComment production.
	ExitRemComment(c *RemCommentContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitEndOfLine is called when exiting the endOfLine production.
	ExitEndOfLine(c *EndOfLineContext)

	// ExitEndOfStatement is called when exiting the endOfStatement production.
	ExitEndOfStatement(c *EndOfStatementContext)
}
