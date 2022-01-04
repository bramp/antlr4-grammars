// Code generated from Teal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package teal // Teal
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTealListener is a complete listener for a parse tree produced by TealParser.
type BaseTealListener struct{}

var _ TealListener = &BaseTealListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTealListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTealListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTealListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTealListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterChunk is called when production chunk is entered.
func (s *BaseTealListener) EnterChunk(ctx *ChunkContext) {}

// ExitChunk is called when production chunk is exited.
func (s *BaseTealListener) ExitChunk(ctx *ChunkContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseTealListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseTealListener) ExitBlock(ctx *BlockContext) {}

// EnterSemiStat is called when production SemiStat is entered.
func (s *BaseTealListener) EnterSemiStat(ctx *SemiStatContext) {}

// ExitSemiStat is called when production SemiStat is exited.
func (s *BaseTealListener) ExitSemiStat(ctx *SemiStatContext) {}

// EnterAssignStat is called when production AssignStat is entered.
func (s *BaseTealListener) EnterAssignStat(ctx *AssignStatContext) {}

// ExitAssignStat is called when production AssignStat is exited.
func (s *BaseTealListener) ExitAssignStat(ctx *AssignStatContext) {}

// EnterFuncCallStat is called when production FuncCallStat is entered.
func (s *BaseTealListener) EnterFuncCallStat(ctx *FuncCallStatContext) {}

// ExitFuncCallStat is called when production FuncCallStat is exited.
func (s *BaseTealListener) ExitFuncCallStat(ctx *FuncCallStatContext) {}

// EnterLabelStat is called when production LabelStat is entered.
func (s *BaseTealListener) EnterLabelStat(ctx *LabelStatContext) {}

// ExitLabelStat is called when production LabelStat is exited.
func (s *BaseTealListener) ExitLabelStat(ctx *LabelStatContext) {}

// EnterBreakStat is called when production BreakStat is entered.
func (s *BaseTealListener) EnterBreakStat(ctx *BreakStatContext) {}

// ExitBreakStat is called when production BreakStat is exited.
func (s *BaseTealListener) ExitBreakStat(ctx *BreakStatContext) {}

// EnterGotoStat is called when production GotoStat is entered.
func (s *BaseTealListener) EnterGotoStat(ctx *GotoStatContext) {}

// ExitGotoStat is called when production GotoStat is exited.
func (s *BaseTealListener) ExitGotoStat(ctx *GotoStatContext) {}

// EnterDoStat is called when production DoStat is entered.
func (s *BaseTealListener) EnterDoStat(ctx *DoStatContext) {}

// ExitDoStat is called when production DoStat is exited.
func (s *BaseTealListener) ExitDoStat(ctx *DoStatContext) {}

// EnterWhileStat is called when production WhileStat is entered.
func (s *BaseTealListener) EnterWhileStat(ctx *WhileStatContext) {}

// ExitWhileStat is called when production WhileStat is exited.
func (s *BaseTealListener) ExitWhileStat(ctx *WhileStatContext) {}

// EnterRepeatStat is called when production RepeatStat is entered.
func (s *BaseTealListener) EnterRepeatStat(ctx *RepeatStatContext) {}

// ExitRepeatStat is called when production RepeatStat is exited.
func (s *BaseTealListener) ExitRepeatStat(ctx *RepeatStatContext) {}

// EnterIfStat is called when production IfStat is entered.
func (s *BaseTealListener) EnterIfStat(ctx *IfStatContext) {}

// ExitIfStat is called when production IfStat is exited.
func (s *BaseTealListener) ExitIfStat(ctx *IfStatContext) {}

// EnterForStat is called when production ForStat is entered.
func (s *BaseTealListener) EnterForStat(ctx *ForStatContext) {}

// ExitForStat is called when production ForStat is exited.
func (s *BaseTealListener) ExitForStat(ctx *ForStatContext) {}

// EnterForInStat is called when production ForInStat is entered.
func (s *BaseTealListener) EnterForInStat(ctx *ForInStatContext) {}

// ExitForInStat is called when production ForInStat is exited.
func (s *BaseTealListener) ExitForInStat(ctx *ForInStatContext) {}

// EnterFuncStat is called when production FuncStat is entered.
func (s *BaseTealListener) EnterFuncStat(ctx *FuncStatContext) {}

// ExitFuncStat is called when production FuncStat is exited.
func (s *BaseTealListener) ExitFuncStat(ctx *FuncStatContext) {}

// EnterLocalFuncStat is called when production LocalFuncStat is entered.
func (s *BaseTealListener) EnterLocalFuncStat(ctx *LocalFuncStatContext) {}

// ExitLocalFuncStat is called when production LocalFuncStat is exited.
func (s *BaseTealListener) ExitLocalFuncStat(ctx *LocalFuncStatContext) {}

// EnterLocalAttrAssignStat is called when production LocalAttrAssignStat is entered.
func (s *BaseTealListener) EnterLocalAttrAssignStat(ctx *LocalAttrAssignStatContext) {}

// ExitLocalAttrAssignStat is called when production LocalAttrAssignStat is exited.
func (s *BaseTealListener) ExitLocalAttrAssignStat(ctx *LocalAttrAssignStatContext) {}

// EnterLocalNewTypeStat is called when production LocalNewTypeStat is entered.
func (s *BaseTealListener) EnterLocalNewTypeStat(ctx *LocalNewTypeStatContext) {}

// ExitLocalNewTypeStat is called when production LocalNewTypeStat is exited.
func (s *BaseTealListener) ExitLocalNewTypeStat(ctx *LocalNewTypeStatContext) {}

// EnterGlobalFuncStat is called when production GlobalFuncStat is entered.
func (s *BaseTealListener) EnterGlobalFuncStat(ctx *GlobalFuncStatContext) {}

// ExitGlobalFuncStat is called when production GlobalFuncStat is exited.
func (s *BaseTealListener) ExitGlobalFuncStat(ctx *GlobalFuncStatContext) {}

// EnterGlobalAttrStat is called when production GlobalAttrStat is entered.
func (s *BaseTealListener) EnterGlobalAttrStat(ctx *GlobalAttrStatContext) {}

// ExitGlobalAttrStat is called when production GlobalAttrStat is exited.
func (s *BaseTealListener) ExitGlobalAttrStat(ctx *GlobalAttrStatContext) {}

// EnterGlobalAttrAssignStat is called when production GlobalAttrAssignStat is entered.
func (s *BaseTealListener) EnterGlobalAttrAssignStat(ctx *GlobalAttrAssignStatContext) {}

// ExitGlobalAttrAssignStat is called when production GlobalAttrAssignStat is exited.
func (s *BaseTealListener) ExitGlobalAttrAssignStat(ctx *GlobalAttrAssignStatContext) {}

// EnterGlobalAssignStat is called when production GlobalAssignStat is entered.
func (s *BaseTealListener) EnterGlobalAssignStat(ctx *GlobalAssignStatContext) {}

// ExitGlobalAssignStat is called when production GlobalAssignStat is exited.
func (s *BaseTealListener) ExitGlobalAssignStat(ctx *GlobalAssignStatContext) {}

// EnterAttnamelist is called when production attnamelist is entered.
func (s *BaseTealListener) EnterAttnamelist(ctx *AttnamelistContext) {}

// ExitAttnamelist is called when production attnamelist is exited.
func (s *BaseTealListener) ExitAttnamelist(ctx *AttnamelistContext) {}

// EnterAttrib is called when production attrib is entered.
func (s *BaseTealListener) EnterAttrib(ctx *AttribContext) {}

// ExitAttrib is called when production attrib is exited.
func (s *BaseTealListener) ExitAttrib(ctx *AttribContext) {}

// EnterTyp is called when production typ is entered.
func (s *BaseTealListener) EnterTyp(ctx *TypContext) {}

// ExitTyp is called when production typ is exited.
func (s *BaseTealListener) ExitTyp(ctx *TypContext) {}

// EnterBasetype is called when production basetype is entered.
func (s *BaseTealListener) EnterBasetype(ctx *BasetypeContext) {}

// ExitBasetype is called when production basetype is exited.
func (s *BaseTealListener) ExitBasetype(ctx *BasetypeContext) {}

// EnterTypelist is called when production typelist is entered.
func (s *BaseTealListener) EnterTypelist(ctx *TypelistContext) {}

// ExitTypelist is called when production typelist is exited.
func (s *BaseTealListener) ExitTypelist(ctx *TypelistContext) {}

// EnterRetlist is called when production retlist is entered.
func (s *BaseTealListener) EnterRetlist(ctx *RetlistContext) {}

// ExitRetlist is called when production retlist is exited.
func (s *BaseTealListener) ExitRetlist(ctx *RetlistContext) {}

// EnterTypeargs is called when production typeargs is entered.
func (s *BaseTealListener) EnterTypeargs(ctx *TypeargsContext) {}

// ExitTypeargs is called when production typeargs is exited.
func (s *BaseTealListener) ExitTypeargs(ctx *TypeargsContext) {}

// EnterRecordNewType is called when production RecordNewType is entered.
func (s *BaseTealListener) EnterRecordNewType(ctx *RecordNewTypeContext) {}

// ExitRecordNewType is called when production RecordNewType is exited.
func (s *BaseTealListener) ExitRecordNewType(ctx *RecordNewTypeContext) {}

// EnterEnumNewType is called when production EnumNewType is entered.
func (s *BaseTealListener) EnterEnumNewType(ctx *EnumNewTypeContext) {}

// ExitEnumNewType is called when production EnumNewType is exited.
func (s *BaseTealListener) ExitEnumNewType(ctx *EnumNewTypeContext) {}

// EnterFuncNewType is called when production FuncNewType is entered.
func (s *BaseTealListener) EnterFuncNewType(ctx *FuncNewTypeContext) {}

// ExitFuncNewType is called when production FuncNewType is exited.
func (s *BaseTealListener) ExitFuncNewType(ctx *FuncNewTypeContext) {}

// EnterFunctiontype is called when production functiontype is entered.
func (s *BaseTealListener) EnterFunctiontype(ctx *FunctiontypeContext) {}

// ExitFunctiontype is called when production functiontype is exited.
func (s *BaseTealListener) ExitFunctiontype(ctx *FunctiontypeContext) {}

// EnterPartypelist is called when production partypelist is entered.
func (s *BaseTealListener) EnterPartypelist(ctx *PartypelistContext) {}

// ExitPartypelist is called when production partypelist is exited.
func (s *BaseTealListener) ExitPartypelist(ctx *PartypelistContext) {}

// EnterPartype is called when production partype is entered.
func (s *BaseTealListener) EnterPartype(ctx *PartypeContext) {}

// ExitPartype is called when production partype is exited.
func (s *BaseTealListener) ExitPartype(ctx *PartypeContext) {}

// EnterParnamelist is called when production parnamelist is entered.
func (s *BaseTealListener) EnterParnamelist(ctx *ParnamelistContext) {}

// ExitParnamelist is called when production parnamelist is exited.
func (s *BaseTealListener) ExitParnamelist(ctx *ParnamelistContext) {}

// EnterParname is called when production parname is entered.
func (s *BaseTealListener) EnterParname(ctx *ParnameContext) {}

// ExitParname is called when production parname is exited.
func (s *BaseTealListener) ExitParname(ctx *ParnameContext) {}

// EnterReturnStat is called when production ReturnStat is entered.
func (s *BaseTealListener) EnterReturnStat(ctx *ReturnStatContext) {}

// ExitReturnStat is called when production ReturnStat is exited.
func (s *BaseTealListener) ExitReturnStat(ctx *ReturnStatContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseTealListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseTealListener) ExitLabel(ctx *LabelContext) {}

// EnterFuncname is called when production funcname is entered.
func (s *BaseTealListener) EnterFuncname(ctx *FuncnameContext) {}

// ExitFuncname is called when production funcname is exited.
func (s *BaseTealListener) ExitFuncname(ctx *FuncnameContext) {}

// EnterVarlist is called when production varlist is entered.
func (s *BaseTealListener) EnterVarlist(ctx *VarlistContext) {}

// ExitVarlist is called when production varlist is exited.
func (s *BaseTealListener) ExitVarlist(ctx *VarlistContext) {}

// EnterNamelist is called when production namelist is entered.
func (s *BaseTealListener) EnterNamelist(ctx *NamelistContext) {}

// ExitNamelist is called when production namelist is exited.
func (s *BaseTealListener) ExitNamelist(ctx *NamelistContext) {}

// EnterExplist is called when production explist is entered.
func (s *BaseTealListener) EnterExplist(ctx *ExplistContext) {}

// ExitExplist is called when production explist is exited.
func (s *BaseTealListener) ExitExplist(ctx *ExplistContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseTealListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseTealListener) ExitExp(ctx *ExpContext) {}

// EnterPrefixexp is called when production prefixexp is entered.
func (s *BaseTealListener) EnterPrefixexp(ctx *PrefixexpContext) {}

// ExitPrefixexp is called when production prefixexp is exited.
func (s *BaseTealListener) ExitPrefixexp(ctx *PrefixexpContext) {}

// EnterFunctioncall is called when production functioncall is entered.
func (s *BaseTealListener) EnterFunctioncall(ctx *FunctioncallContext) {}

// ExitFunctioncall is called when production functioncall is exited.
func (s *BaseTealListener) ExitFunctioncall(ctx *FunctioncallContext) {}

// EnterVarOrExp is called when production varOrExp is entered.
func (s *BaseTealListener) EnterVarOrExp(ctx *VarOrExpContext) {}

// ExitVarOrExp is called when production varOrExp is exited.
func (s *BaseTealListener) ExitVarOrExp(ctx *VarOrExpContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseTealListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseTealListener) ExitVariable(ctx *VariableContext) {}

// EnterVarSuffix is called when production varSuffix is entered.
func (s *BaseTealListener) EnterVarSuffix(ctx *VarSuffixContext) {}

// ExitVarSuffix is called when production varSuffix is exited.
func (s *BaseTealListener) ExitVarSuffix(ctx *VarSuffixContext) {}

// EnterNameAndArgs is called when production nameAndArgs is entered.
func (s *BaseTealListener) EnterNameAndArgs(ctx *NameAndArgsContext) {}

// ExitNameAndArgs is called when production nameAndArgs is exited.
func (s *BaseTealListener) ExitNameAndArgs(ctx *NameAndArgsContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseTealListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseTealListener) ExitArgs(ctx *ArgsContext) {}

// EnterFunctiondef is called when production functiondef is entered.
func (s *BaseTealListener) EnterFunctiondef(ctx *FunctiondefContext) {}

// ExitFunctiondef is called when production functiondef is exited.
func (s *BaseTealListener) ExitFunctiondef(ctx *FunctiondefContext) {}

// EnterFuncbody is called when production funcbody is entered.
func (s *BaseTealListener) EnterFuncbody(ctx *FuncbodyContext) {}

// ExitFuncbody is called when production funcbody is exited.
func (s *BaseTealListener) ExitFuncbody(ctx *FuncbodyContext) {}

// EnterParlist is called when production parlist is entered.
func (s *BaseTealListener) EnterParlist(ctx *ParlistContext) {}

// ExitParlist is called when production parlist is exited.
func (s *BaseTealListener) ExitParlist(ctx *ParlistContext) {}

// EnterTableconstructor is called when production tableconstructor is entered.
func (s *BaseTealListener) EnterTableconstructor(ctx *TableconstructorContext) {}

// ExitTableconstructor is called when production tableconstructor is exited.
func (s *BaseTealListener) ExitTableconstructor(ctx *TableconstructorContext) {}

// EnterFieldlist is called when production fieldlist is entered.
func (s *BaseTealListener) EnterFieldlist(ctx *FieldlistContext) {}

// ExitFieldlist is called when production fieldlist is exited.
func (s *BaseTealListener) ExitFieldlist(ctx *FieldlistContext) {}

// EnterBracketAssginField is called when production BracketAssginField is entered.
func (s *BaseTealListener) EnterBracketAssginField(ctx *BracketAssginFieldContext) {}

// ExitBracketAssginField is called when production BracketAssginField is exited.
func (s *BaseTealListener) ExitBracketAssginField(ctx *BracketAssginFieldContext) {}

// EnterAssignField is called when production AssignField is entered.
func (s *BaseTealListener) EnterAssignField(ctx *AssignFieldContext) {}

// ExitAssignField is called when production AssignField is exited.
func (s *BaseTealListener) ExitAssignField(ctx *AssignFieldContext) {}

// EnterAssignNewTypeField is called when production AssignNewTypeField is entered.
func (s *BaseTealListener) EnterAssignNewTypeField(ctx *AssignNewTypeFieldContext) {}

// ExitAssignNewTypeField is called when production AssignNewTypeField is exited.
func (s *BaseTealListener) ExitAssignNewTypeField(ctx *AssignNewTypeFieldContext) {}

// EnterExprField is called when production ExprField is entered.
func (s *BaseTealListener) EnterExprField(ctx *ExprFieldContext) {}

// ExitExprField is called when production ExprField is exited.
func (s *BaseTealListener) ExitExprField(ctx *ExprFieldContext) {}

// EnterFieldsep is called when production fieldsep is entered.
func (s *BaseTealListener) EnterFieldsep(ctx *FieldsepContext) {}

// ExitFieldsep is called when production fieldsep is exited.
func (s *BaseTealListener) ExitFieldsep(ctx *FieldsepContext) {}

// EnterOperatorOr is called when production operatorOr is entered.
func (s *BaseTealListener) EnterOperatorOr(ctx *OperatorOrContext) {}

// ExitOperatorOr is called when production operatorOr is exited.
func (s *BaseTealListener) ExitOperatorOr(ctx *OperatorOrContext) {}

// EnterOperatorAnd is called when production operatorAnd is entered.
func (s *BaseTealListener) EnterOperatorAnd(ctx *OperatorAndContext) {}

// ExitOperatorAnd is called when production operatorAnd is exited.
func (s *BaseTealListener) ExitOperatorAnd(ctx *OperatorAndContext) {}

// EnterOperatorComparison is called when production operatorComparison is entered.
func (s *BaseTealListener) EnterOperatorComparison(ctx *OperatorComparisonContext) {}

// ExitOperatorComparison is called when production operatorComparison is exited.
func (s *BaseTealListener) ExitOperatorComparison(ctx *OperatorComparisonContext) {}

// EnterOperatorStrcat is called when production operatorStrcat is entered.
func (s *BaseTealListener) EnterOperatorStrcat(ctx *OperatorStrcatContext) {}

// ExitOperatorStrcat is called when production operatorStrcat is exited.
func (s *BaseTealListener) ExitOperatorStrcat(ctx *OperatorStrcatContext) {}

// EnterOperatorAddSub is called when production operatorAddSub is entered.
func (s *BaseTealListener) EnterOperatorAddSub(ctx *OperatorAddSubContext) {}

// ExitOperatorAddSub is called when production operatorAddSub is exited.
func (s *BaseTealListener) ExitOperatorAddSub(ctx *OperatorAddSubContext) {}

// EnterOperatorMulDivMod is called when production operatorMulDivMod is entered.
func (s *BaseTealListener) EnterOperatorMulDivMod(ctx *OperatorMulDivModContext) {}

// ExitOperatorMulDivMod is called when production operatorMulDivMod is exited.
func (s *BaseTealListener) ExitOperatorMulDivMod(ctx *OperatorMulDivModContext) {}

// EnterOperatorBitwise is called when production operatorBitwise is entered.
func (s *BaseTealListener) EnterOperatorBitwise(ctx *OperatorBitwiseContext) {}

// ExitOperatorBitwise is called when production operatorBitwise is exited.
func (s *BaseTealListener) ExitOperatorBitwise(ctx *OperatorBitwiseContext) {}

// EnterOperatorUnary is called when production operatorUnary is entered.
func (s *BaseTealListener) EnterOperatorUnary(ctx *OperatorUnaryContext) {}

// ExitOperatorUnary is called when production operatorUnary is exited.
func (s *BaseTealListener) ExitOperatorUnary(ctx *OperatorUnaryContext) {}

// EnterOperatorPower is called when production operatorPower is entered.
func (s *BaseTealListener) EnterOperatorPower(ctx *OperatorPowerContext) {}

// ExitOperatorPower is called when production operatorPower is exited.
func (s *BaseTealListener) ExitOperatorPower(ctx *OperatorPowerContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseTealListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseTealListener) ExitNumber(ctx *NumberContext) {}

// EnterStr is called when production str is entered.
func (s *BaseTealListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BaseTealListener) ExitStr(ctx *StrContext) {}
