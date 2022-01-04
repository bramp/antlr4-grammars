// Code generated from Teal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package teal // Teal
import "github.com/antlr/antlr4/runtime/Go/antlr"

// TealListener is a complete listener for a parse tree produced by TealParser.
type TealListener interface {
	antlr.ParseTreeListener

	// EnterChunk is called when entering the chunk production.
	EnterChunk(c *ChunkContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterSemiStat is called when entering the SemiStat production.
	EnterSemiStat(c *SemiStatContext)

	// EnterAssignStat is called when entering the AssignStat production.
	EnterAssignStat(c *AssignStatContext)

	// EnterFuncCallStat is called when entering the FuncCallStat production.
	EnterFuncCallStat(c *FuncCallStatContext)

	// EnterLabelStat is called when entering the LabelStat production.
	EnterLabelStat(c *LabelStatContext)

	// EnterBreakStat is called when entering the BreakStat production.
	EnterBreakStat(c *BreakStatContext)

	// EnterGotoStat is called when entering the GotoStat production.
	EnterGotoStat(c *GotoStatContext)

	// EnterDoStat is called when entering the DoStat production.
	EnterDoStat(c *DoStatContext)

	// EnterWhileStat is called when entering the WhileStat production.
	EnterWhileStat(c *WhileStatContext)

	// EnterRepeatStat is called when entering the RepeatStat production.
	EnterRepeatStat(c *RepeatStatContext)

	// EnterIfStat is called when entering the IfStat production.
	EnterIfStat(c *IfStatContext)

	// EnterForStat is called when entering the ForStat production.
	EnterForStat(c *ForStatContext)

	// EnterForInStat is called when entering the ForInStat production.
	EnterForInStat(c *ForInStatContext)

	// EnterFuncStat is called when entering the FuncStat production.
	EnterFuncStat(c *FuncStatContext)

	// EnterLocalFuncStat is called when entering the LocalFuncStat production.
	EnterLocalFuncStat(c *LocalFuncStatContext)

	// EnterLocalAttrAssignStat is called when entering the LocalAttrAssignStat production.
	EnterLocalAttrAssignStat(c *LocalAttrAssignStatContext)

	// EnterLocalNewTypeStat is called when entering the LocalNewTypeStat production.
	EnterLocalNewTypeStat(c *LocalNewTypeStatContext)

	// EnterGlobalFuncStat is called when entering the GlobalFuncStat production.
	EnterGlobalFuncStat(c *GlobalFuncStatContext)

	// EnterGlobalAttrStat is called when entering the GlobalAttrStat production.
	EnterGlobalAttrStat(c *GlobalAttrStatContext)

	// EnterGlobalAttrAssignStat is called when entering the GlobalAttrAssignStat production.
	EnterGlobalAttrAssignStat(c *GlobalAttrAssignStatContext)

	// EnterGlobalAssignStat is called when entering the GlobalAssignStat production.
	EnterGlobalAssignStat(c *GlobalAssignStatContext)

	// EnterAttnamelist is called when entering the attnamelist production.
	EnterAttnamelist(c *AttnamelistContext)

	// EnterAttrib is called when entering the attrib production.
	EnterAttrib(c *AttribContext)

	// EnterTyp is called when entering the typ production.
	EnterTyp(c *TypContext)

	// EnterBasetype is called when entering the basetype production.
	EnterBasetype(c *BasetypeContext)

	// EnterTypelist is called when entering the typelist production.
	EnterTypelist(c *TypelistContext)

	// EnterRetlist is called when entering the retlist production.
	EnterRetlist(c *RetlistContext)

	// EnterTypeargs is called when entering the typeargs production.
	EnterTypeargs(c *TypeargsContext)

	// EnterRecordNewType is called when entering the RecordNewType production.
	EnterRecordNewType(c *RecordNewTypeContext)

	// EnterEnumNewType is called when entering the EnumNewType production.
	EnterEnumNewType(c *EnumNewTypeContext)

	// EnterFuncNewType is called when entering the FuncNewType production.
	EnterFuncNewType(c *FuncNewTypeContext)

	// EnterFunctiontype is called when entering the functiontype production.
	EnterFunctiontype(c *FunctiontypeContext)

	// EnterPartypelist is called when entering the partypelist production.
	EnterPartypelist(c *PartypelistContext)

	// EnterPartype is called when entering the partype production.
	EnterPartype(c *PartypeContext)

	// EnterParnamelist is called when entering the parnamelist production.
	EnterParnamelist(c *ParnamelistContext)

	// EnterParname is called when entering the parname production.
	EnterParname(c *ParnameContext)

	// EnterReturnStat is called when entering the ReturnStat production.
	EnterReturnStat(c *ReturnStatContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterFuncname is called when entering the funcname production.
	EnterFuncname(c *FuncnameContext)

	// EnterVarlist is called when entering the varlist production.
	EnterVarlist(c *VarlistContext)

	// EnterNamelist is called when entering the namelist production.
	EnterNamelist(c *NamelistContext)

	// EnterExplist is called when entering the explist production.
	EnterExplist(c *ExplistContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterPrefixexp is called when entering the prefixexp production.
	EnterPrefixexp(c *PrefixexpContext)

	// EnterFunctioncall is called when entering the functioncall production.
	EnterFunctioncall(c *FunctioncallContext)

	// EnterVarOrExp is called when entering the varOrExp production.
	EnterVarOrExp(c *VarOrExpContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterVarSuffix is called when entering the varSuffix production.
	EnterVarSuffix(c *VarSuffixContext)

	// EnterNameAndArgs is called when entering the nameAndArgs production.
	EnterNameAndArgs(c *NameAndArgsContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterFunctiondef is called when entering the functiondef production.
	EnterFunctiondef(c *FunctiondefContext)

	// EnterFuncbody is called when entering the funcbody production.
	EnterFuncbody(c *FuncbodyContext)

	// EnterParlist is called when entering the parlist production.
	EnterParlist(c *ParlistContext)

	// EnterTableconstructor is called when entering the tableconstructor production.
	EnterTableconstructor(c *TableconstructorContext)

	// EnterFieldlist is called when entering the fieldlist production.
	EnterFieldlist(c *FieldlistContext)

	// EnterBracketAssginField is called when entering the BracketAssginField production.
	EnterBracketAssginField(c *BracketAssginFieldContext)

	// EnterAssignField is called when entering the AssignField production.
	EnterAssignField(c *AssignFieldContext)

	// EnterAssignNewTypeField is called when entering the AssignNewTypeField production.
	EnterAssignNewTypeField(c *AssignNewTypeFieldContext)

	// EnterExprField is called when entering the ExprField production.
	EnterExprField(c *ExprFieldContext)

	// EnterFieldsep is called when entering the fieldsep production.
	EnterFieldsep(c *FieldsepContext)

	// EnterOperatorOr is called when entering the operatorOr production.
	EnterOperatorOr(c *OperatorOrContext)

	// EnterOperatorAnd is called when entering the operatorAnd production.
	EnterOperatorAnd(c *OperatorAndContext)

	// EnterOperatorComparison is called when entering the operatorComparison production.
	EnterOperatorComparison(c *OperatorComparisonContext)

	// EnterOperatorStrcat is called when entering the operatorStrcat production.
	EnterOperatorStrcat(c *OperatorStrcatContext)

	// EnterOperatorAddSub is called when entering the operatorAddSub production.
	EnterOperatorAddSub(c *OperatorAddSubContext)

	// EnterOperatorMulDivMod is called when entering the operatorMulDivMod production.
	EnterOperatorMulDivMod(c *OperatorMulDivModContext)

	// EnterOperatorBitwise is called when entering the operatorBitwise production.
	EnterOperatorBitwise(c *OperatorBitwiseContext)

	// EnterOperatorUnary is called when entering the operatorUnary production.
	EnterOperatorUnary(c *OperatorUnaryContext)

	// EnterOperatorPower is called when entering the operatorPower production.
	EnterOperatorPower(c *OperatorPowerContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// ExitChunk is called when exiting the chunk production.
	ExitChunk(c *ChunkContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitSemiStat is called when exiting the SemiStat production.
	ExitSemiStat(c *SemiStatContext)

	// ExitAssignStat is called when exiting the AssignStat production.
	ExitAssignStat(c *AssignStatContext)

	// ExitFuncCallStat is called when exiting the FuncCallStat production.
	ExitFuncCallStat(c *FuncCallStatContext)

	// ExitLabelStat is called when exiting the LabelStat production.
	ExitLabelStat(c *LabelStatContext)

	// ExitBreakStat is called when exiting the BreakStat production.
	ExitBreakStat(c *BreakStatContext)

	// ExitGotoStat is called when exiting the GotoStat production.
	ExitGotoStat(c *GotoStatContext)

	// ExitDoStat is called when exiting the DoStat production.
	ExitDoStat(c *DoStatContext)

	// ExitWhileStat is called when exiting the WhileStat production.
	ExitWhileStat(c *WhileStatContext)

	// ExitRepeatStat is called when exiting the RepeatStat production.
	ExitRepeatStat(c *RepeatStatContext)

	// ExitIfStat is called when exiting the IfStat production.
	ExitIfStat(c *IfStatContext)

	// ExitForStat is called when exiting the ForStat production.
	ExitForStat(c *ForStatContext)

	// ExitForInStat is called when exiting the ForInStat production.
	ExitForInStat(c *ForInStatContext)

	// ExitFuncStat is called when exiting the FuncStat production.
	ExitFuncStat(c *FuncStatContext)

	// ExitLocalFuncStat is called when exiting the LocalFuncStat production.
	ExitLocalFuncStat(c *LocalFuncStatContext)

	// ExitLocalAttrAssignStat is called when exiting the LocalAttrAssignStat production.
	ExitLocalAttrAssignStat(c *LocalAttrAssignStatContext)

	// ExitLocalNewTypeStat is called when exiting the LocalNewTypeStat production.
	ExitLocalNewTypeStat(c *LocalNewTypeStatContext)

	// ExitGlobalFuncStat is called when exiting the GlobalFuncStat production.
	ExitGlobalFuncStat(c *GlobalFuncStatContext)

	// ExitGlobalAttrStat is called when exiting the GlobalAttrStat production.
	ExitGlobalAttrStat(c *GlobalAttrStatContext)

	// ExitGlobalAttrAssignStat is called when exiting the GlobalAttrAssignStat production.
	ExitGlobalAttrAssignStat(c *GlobalAttrAssignStatContext)

	// ExitGlobalAssignStat is called when exiting the GlobalAssignStat production.
	ExitGlobalAssignStat(c *GlobalAssignStatContext)

	// ExitAttnamelist is called when exiting the attnamelist production.
	ExitAttnamelist(c *AttnamelistContext)

	// ExitAttrib is called when exiting the attrib production.
	ExitAttrib(c *AttribContext)

	// ExitTyp is called when exiting the typ production.
	ExitTyp(c *TypContext)

	// ExitBasetype is called when exiting the basetype production.
	ExitBasetype(c *BasetypeContext)

	// ExitTypelist is called when exiting the typelist production.
	ExitTypelist(c *TypelistContext)

	// ExitRetlist is called when exiting the retlist production.
	ExitRetlist(c *RetlistContext)

	// ExitTypeargs is called when exiting the typeargs production.
	ExitTypeargs(c *TypeargsContext)

	// ExitRecordNewType is called when exiting the RecordNewType production.
	ExitRecordNewType(c *RecordNewTypeContext)

	// ExitEnumNewType is called when exiting the EnumNewType production.
	ExitEnumNewType(c *EnumNewTypeContext)

	// ExitFuncNewType is called when exiting the FuncNewType production.
	ExitFuncNewType(c *FuncNewTypeContext)

	// ExitFunctiontype is called when exiting the functiontype production.
	ExitFunctiontype(c *FunctiontypeContext)

	// ExitPartypelist is called when exiting the partypelist production.
	ExitPartypelist(c *PartypelistContext)

	// ExitPartype is called when exiting the partype production.
	ExitPartype(c *PartypeContext)

	// ExitParnamelist is called when exiting the parnamelist production.
	ExitParnamelist(c *ParnamelistContext)

	// ExitParname is called when exiting the parname production.
	ExitParname(c *ParnameContext)

	// ExitReturnStat is called when exiting the ReturnStat production.
	ExitReturnStat(c *ReturnStatContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitFuncname is called when exiting the funcname production.
	ExitFuncname(c *FuncnameContext)

	// ExitVarlist is called when exiting the varlist production.
	ExitVarlist(c *VarlistContext)

	// ExitNamelist is called when exiting the namelist production.
	ExitNamelist(c *NamelistContext)

	// ExitExplist is called when exiting the explist production.
	ExitExplist(c *ExplistContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitPrefixexp is called when exiting the prefixexp production.
	ExitPrefixexp(c *PrefixexpContext)

	// ExitFunctioncall is called when exiting the functioncall production.
	ExitFunctioncall(c *FunctioncallContext)

	// ExitVarOrExp is called when exiting the varOrExp production.
	ExitVarOrExp(c *VarOrExpContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitVarSuffix is called when exiting the varSuffix production.
	ExitVarSuffix(c *VarSuffixContext)

	// ExitNameAndArgs is called when exiting the nameAndArgs production.
	ExitNameAndArgs(c *NameAndArgsContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitFunctiondef is called when exiting the functiondef production.
	ExitFunctiondef(c *FunctiondefContext)

	// ExitFuncbody is called when exiting the funcbody production.
	ExitFuncbody(c *FuncbodyContext)

	// ExitParlist is called when exiting the parlist production.
	ExitParlist(c *ParlistContext)

	// ExitTableconstructor is called when exiting the tableconstructor production.
	ExitTableconstructor(c *TableconstructorContext)

	// ExitFieldlist is called when exiting the fieldlist production.
	ExitFieldlist(c *FieldlistContext)

	// ExitBracketAssginField is called when exiting the BracketAssginField production.
	ExitBracketAssginField(c *BracketAssginFieldContext)

	// ExitAssignField is called when exiting the AssignField production.
	ExitAssignField(c *AssignFieldContext)

	// ExitAssignNewTypeField is called when exiting the AssignNewTypeField production.
	ExitAssignNewTypeField(c *AssignNewTypeFieldContext)

	// ExitExprField is called when exiting the ExprField production.
	ExitExprField(c *ExprFieldContext)

	// ExitFieldsep is called when exiting the fieldsep production.
	ExitFieldsep(c *FieldsepContext)

	// ExitOperatorOr is called when exiting the operatorOr production.
	ExitOperatorOr(c *OperatorOrContext)

	// ExitOperatorAnd is called when exiting the operatorAnd production.
	ExitOperatorAnd(c *OperatorAndContext)

	// ExitOperatorComparison is called when exiting the operatorComparison production.
	ExitOperatorComparison(c *OperatorComparisonContext)

	// ExitOperatorStrcat is called when exiting the operatorStrcat production.
	ExitOperatorStrcat(c *OperatorStrcatContext)

	// ExitOperatorAddSub is called when exiting the operatorAddSub production.
	ExitOperatorAddSub(c *OperatorAddSubContext)

	// ExitOperatorMulDivMod is called when exiting the operatorMulDivMod production.
	ExitOperatorMulDivMod(c *OperatorMulDivModContext)

	// ExitOperatorBitwise is called when exiting the operatorBitwise production.
	ExitOperatorBitwise(c *OperatorBitwiseContext)

	// ExitOperatorUnary is called when exiting the operatorUnary production.
	ExitOperatorUnary(c *OperatorUnaryContext)

	// ExitOperatorPower is called when exiting the operatorPower production.
	ExitOperatorPower(c *OperatorPowerContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)
}
