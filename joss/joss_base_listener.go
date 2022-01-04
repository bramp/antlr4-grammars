// Code generated from joss.g4 by ANTLR 4.9.3. DO NOT EDIT.

package joss // joss
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasejossListener is a complete listener for a parse tree produced by jossParser.
type BasejossListener struct{}

var _ jossListener = &BasejossListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasejossListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasejossListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasejossListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasejossListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasejossListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasejossListener) ExitProg(ctx *ProgContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasejossListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasejossListener) ExitStatement(ctx *StatementContext) {}

// EnterDirect is called when production direct is entered.
func (s *BasejossListener) EnterDirect(ctx *DirectContext) {}

// ExitDirect is called when production direct is exited.
func (s *BasejossListener) ExitDirect(ctx *DirectContext) {}

// EnterIndirect_ is called when production indirect_ is entered.
func (s *BasejossListener) EnterIndirect_(ctx *Indirect_Context) {}

// ExitIndirect_ is called when production indirect_ is exited.
func (s *BasejossListener) ExitIndirect_(ctx *Indirect_Context) {}

// EnterIndirectCMD is called when production indirectCMD is entered.
func (s *BasejossListener) EnterIndirectCMD(ctx *IndirectCMDContext) {}

// ExitIndirectCMD is called when production indirectCMD is exited.
func (s *BasejossListener) ExitIndirectCMD(ctx *IndirectCMDContext) {}

// EnterArbitraryCMD is called when production arbitraryCMD is entered.
func (s *BasejossListener) EnterArbitraryCMD(ctx *ArbitraryCMDContext) {}

// ExitArbitraryCMD is called when production arbitraryCMD is exited.
func (s *BasejossListener) ExitArbitraryCMD(ctx *ArbitraryCMDContext) {}

// EnterGoCMD is called when production goCMD is entered.
func (s *BasejossListener) EnterGoCMD(ctx *GoCMDContext) {}

// ExitGoCMD is called when production goCMD is exited.
func (s *BasejossListener) ExitGoCMD(ctx *GoCMDContext) {}

// EnterCancelCMD is called when production cancelCMD is entered.
func (s *BasejossListener) EnterCancelCMD(ctx *CancelCMDContext) {}

// ExitCancelCMD is called when production cancelCMD is exited.
func (s *BasejossListener) ExitCancelCMD(ctx *CancelCMDContext) {}

// EnterDeleteCmdCMD is called when production deleteCmdCMD is entered.
func (s *BasejossListener) EnterDeleteCmdCMD(ctx *DeleteCmdCMDContext) {}

// ExitDeleteCmdCMD is called when production deleteCmdCMD is exited.
func (s *BasejossListener) ExitDeleteCmdCMD(ctx *DeleteCmdCMDContext) {}

// EnterDelCmdSel is called when production delCmdSel is entered.
func (s *BasejossListener) EnterDelCmdSel(ctx *DelCmdSelContext) {}

// ExitDelCmdSel is called when production delCmdSel is exited.
func (s *BasejossListener) ExitDelCmdSel(ctx *DelCmdSelContext) {}

// EnterDelCmdAllSel is called when production delCmdAllSel is entered.
func (s *BasejossListener) EnterDelCmdAllSel(ctx *DelCmdAllSelContext) {}

// ExitDelCmdAllSel is called when production delCmdAllSel is exited.
func (s *BasejossListener) ExitDelCmdAllSel(ctx *DelCmdAllSelContext) {}

// EnterDelCmdAllType is called when production delCmdAllType is entered.
func (s *BasejossListener) EnterDelCmdAllType(ctx *DelCmdAllTypeContext) {}

// ExitDelCmdAllType is called when production delCmdAllType is exited.
func (s *BasejossListener) ExitDelCmdAllType(ctx *DelCmdAllTypeContext) {}

// EnterDoneCMD is called when production doneCMD is entered.
func (s *BasejossListener) EnterDoneCMD(ctx *DoneCMDContext) {}

// ExitDoneCMD is called when production doneCMD is exited.
func (s *BasejossListener) ExitDoneCMD(ctx *DoneCMDContext) {}

// EnterStopCMD is called when production stopCMD is entered.
func (s *BasejossListener) EnterStopCMD(ctx *StopCMDContext) {}

// ExitStopCMD is called when production stopCMD is exited.
func (s *BasejossListener) ExitStopCMD(ctx *StopCMDContext) {}

// EnterDemandCMD is called when production demandCMD is entered.
func (s *BasejossListener) EnterDemandCMD(ctx *DemandCMDContext) {}

// ExitDemandCMD is called when production demandCMD is exited.
func (s *BasejossListener) ExitDemandCMD(ctx *DemandCMDContext) {}

// EnterToCMD is called when production toCMD is entered.
func (s *BasejossListener) EnterToCMD(ctx *ToCMDContext) {}

// ExitToCMD is called when production toCMD is exited.
func (s *BasejossListener) ExitToCMD(ctx *ToCMDContext) {}

// EnterToSel is called when production toSel is entered.
func (s *BasejossListener) EnterToSel(ctx *ToSelContext) {}

// ExitToSel is called when production toSel is exited.
func (s *BasejossListener) ExitToSel(ctx *ToSelContext) {}

// EnterSetCMD is called when production setCMD is entered.
func (s *BasejossListener) EnterSetCMD(ctx *SetCMDContext) {}

// ExitSetCMD is called when production setCMD is exited.
func (s *BasejossListener) ExitSetCMD(ctx *SetCMDContext) {}

// EnterDoCMD is called when production doCMD is entered.
func (s *BasejossListener) EnterDoCMD(ctx *DoCMDContext) {}

// ExitDoCMD is called when production doCMD is exited.
func (s *BasejossListener) ExitDoCMD(ctx *DoCMDContext) {}

// EnterDoSel is called when production doSel is entered.
func (s *BasejossListener) EnterDoSel(ctx *DoSelContext) {}

// ExitDoSel is called when production doSel is exited.
func (s *BasejossListener) ExitDoSel(ctx *DoSelContext) {}

// EnterTypeCMD is called when production typeCMD is entered.
func (s *BasejossListener) EnterTypeCMD(ctx *TypeCMDContext) {}

// ExitTypeCMD is called when production typeCMD is exited.
func (s *BasejossListener) ExitTypeCMD(ctx *TypeCMDContext) {}

// EnterTypeContent is called when production typeContent is entered.
func (s *BasejossListener) EnterTypeContent(ctx *TypeContentContext) {}

// ExitTypeContent is called when production typeContent is exited.
func (s *BasejossListener) ExitTypeContent(ctx *TypeContentContext) {}

// EnterTypeSys is called when production typeSys is entered.
func (s *BasejossListener) EnterTypeSys(ctx *TypeSysContext) {}

// ExitTypeSys is called when production typeSys is exited.
func (s *BasejossListener) ExitTypeSys(ctx *TypeSysContext) {}

// EnterTypeForm is called when production typeForm is entered.
func (s *BasejossListener) EnterTypeForm(ctx *TypeFormContext) {}

// ExitTypeForm is called when production typeForm is exited.
func (s *BasejossListener) ExitTypeForm(ctx *TypeFormContext) {}

// EnterTypeData is called when production typeData is entered.
func (s *BasejossListener) EnterTypeData(ctx *TypeDataContext) {}

// ExitTypeData is called when production typeData is exited.
func (s *BasejossListener) ExitTypeData(ctx *TypeDataContext) {}

// EnterTypeElem is called when production typeElem is entered.
func (s *BasejossListener) EnterTypeElem(ctx *TypeElemContext) {}

// ExitTypeElem is called when production typeElem is exited.
func (s *BasejossListener) ExitTypeElem(ctx *TypeElemContext) {}

// EnterTypeSel is called when production typeSel is entered.
func (s *BasejossListener) EnterTypeSel(ctx *TypeSelContext) {}

// ExitTypeSel is called when production typeSel is exited.
func (s *BasejossListener) ExitTypeSel(ctx *TypeSelContext) {}

// EnterTypeAllSel is called when production typeAllSel is entered.
func (s *BasejossListener) EnterTypeAllSel(ctx *TypeAllSelContext) {}

// ExitTypeAllSel is called when production typeAllSel is exited.
func (s *BasejossListener) ExitTypeAllSel(ctx *TypeAllSelContext) {}

// EnterTypeAllType is called when production typeAllType is entered.
func (s *BasejossListener) EnterTypeAllType(ctx *TypeAllTypeContext) {}

// ExitTypeAllType is called when production typeAllType is exited.
func (s *BasejossListener) ExitTypeAllType(ctx *TypeAllTypeContext) {}

// EnterDeleteValCMD is called when production deleteValCMD is entered.
func (s *BasejossListener) EnterDeleteValCMD(ctx *DeleteValCMDContext) {}

// ExitDeleteValCMD is called when production deleteValCMD is exited.
func (s *BasejossListener) ExitDeleteValCMD(ctx *DeleteValCMDContext) {}

// EnterFormCMD is called when production formCMD is entered.
func (s *BasejossListener) EnterFormCMD(ctx *FormCMDContext) {}

// ExitFormCMD is called when production formCMD is exited.
func (s *BasejossListener) ExitFormCMD(ctx *FormCMDContext) {}

// EnterFormContent is called when production formContent is entered.
func (s *BasejossListener) EnterFormContent(ctx *FormContentContext) {}

// ExitFormContent is called when production formContent is exited.
func (s *BasejossListener) ExitFormContent(ctx *FormContentContext) {}

// EnterFormObject is called when production formObject is entered.
func (s *BasejossListener) EnterFormObject(ctx *FormObjectContext) {}

// ExitFormObject is called when production formObject is exited.
func (s *BasejossListener) ExitFormObject(ctx *FormObjectContext) {}

// EnterFormPH is called when production formPH is entered.
func (s *BasejossListener) EnterFormPH(ctx *FormPHContext) {}

// ExitFormPH is called when production formPH is exited.
func (s *BasejossListener) ExitFormPH(ctx *FormPHContext) {}

// EnterSciNotation is called when production sciNotation is entered.
func (s *BasejossListener) EnterSciNotation(ctx *SciNotationContext) {}

// ExitSciNotation is called when production sciNotation is exited.
func (s *BasejossListener) ExitSciNotation(ctx *SciNotationContext) {}

// EnterDot is called when production dot is entered.
func (s *BasejossListener) EnterDot(ctx *DotContext) {}

// ExitDot is called when production dot is exited.
func (s *BasejossListener) ExitDot(ctx *DotContext) {}

// EnterFixedNotation is called when production fixedNotation is entered.
func (s *BasejossListener) EnterFixedNotation(ctx *FixedNotationContext) {}

// ExitFixedNotation is called when production fixedNotation is exited.
func (s *BasejossListener) ExitFixedNotation(ctx *FixedNotationContext) {}

// EnterUScore is called when production uScore is entered.
func (s *BasejossListener) EnterUScore(ctx *UScoreContext) {}

// ExitUScore is called when production uScore is exited.
func (s *BasejossListener) ExitUScore(ctx *UScoreContext) {}

// EnterStepSel is called when production stepSel is entered.
func (s *BasejossListener) EnterStepSel(ctx *StepSelContext) {}

// ExitStepSel is called when production stepSel is exited.
func (s *BasejossListener) ExitStepSel(ctx *StepSelContext) {}

// EnterPartSel is called when production partSel is entered.
func (s *BasejossListener) EnterPartSel(ctx *PartSelContext) {}

// ExitPartSel is called when production partSel is exited.
func (s *BasejossListener) ExitPartSel(ctx *PartSelContext) {}

// EnterFormSel is called when production formSel is entered.
func (s *BasejossListener) EnterFormSel(ctx *FormSelContext) {}

// ExitFormSel is called when production formSel is exited.
func (s *BasejossListener) ExitFormSel(ctx *FormSelContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BasejossListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BasejossListener) ExitConditional(ctx *ConditionalContext) {}

// EnterIf_ is called when production if_ is entered.
func (s *BasejossListener) EnterIf_(ctx *If_Context) {}

// ExitIf_ is called when production if_ is exited.
func (s *BasejossListener) ExitIf_(ctx *If_Context) {}

// EnterBoolExp is called when production boolExp is entered.
func (s *BasejossListener) EnterBoolExp(ctx *BoolExpContext) {}

// ExitBoolExp is called when production boolExp is exited.
func (s *BasejossListener) ExitBoolExp(ctx *BoolExpContext) {}

// EnterBoolOp is called when production boolOp is entered.
func (s *BasejossListener) EnterBoolOp(ctx *BoolOpContext) {}

// ExitBoolOp is called when production boolOp is exited.
func (s *BasejossListener) ExitBoolOp(ctx *BoolOpContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BasejossListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BasejossListener) ExitComparison(ctx *ComparisonContext) {}

// EnterBoolComp is called when production boolComp is entered.
func (s *BasejossListener) EnterBoolComp(ctx *BoolCompContext) {}

// ExitBoolComp is called when production boolComp is exited.
func (s *BasejossListener) ExitBoolComp(ctx *BoolCompContext) {}

// EnterFor_ is called when production for_ is entered.
func (s *BasejossListener) EnterFor_(ctx *For_Context) {}

// ExitFor_ is called when production for_ is exited.
func (s *BasejossListener) ExitFor_(ctx *For_Context) {}

// EnterRange_ is called when production range_ is entered.
func (s *BasejossListener) EnterRange_(ctx *Range_Context) {}

// ExitRange_ is called when production range_ is exited.
func (s *BasejossListener) ExitRange_(ctx *Range_Context) {}

// EnterRangeExp is called when production rangeExp is entered.
func (s *BasejossListener) EnterRangeExp(ctx *RangeExpContext) {}

// ExitRangeExp is called when production rangeExp is exited.
func (s *BasejossListener) ExitRangeExp(ctx *RangeExpContext) {}

// EnterRangeVal is called when production rangeVal is entered.
func (s *BasejossListener) EnterRangeVal(ctx *RangeValContext) {}

// ExitRangeVal is called when production rangeVal is exited.
func (s *BasejossListener) ExitRangeVal(ctx *RangeValContext) {}

// EnterMathExp is called when production mathExp is entered.
func (s *BasejossListener) EnterMathExp(ctx *MathExpContext) {}

// ExitMathExp is called when production mathExp is exited.
func (s *BasejossListener) ExitMathExp(ctx *MathExpContext) {}

// EnterTerm is called when production term is entered.
func (s *BasejossListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasejossListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BasejossListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BasejossListener) ExitFactor(ctx *FactorContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasejossListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasejossListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterValue is called when production value is entered.
func (s *BasejossListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasejossListener) ExitValue(ctx *ValueContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BasejossListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BasejossListener) ExitFunction_(ctx *Function_Context) {}

// EnterFuncSqrt is called when production funcSqrt is entered.
func (s *BasejossListener) EnterFuncSqrt(ctx *FuncSqrtContext) {}

// ExitFuncSqrt is called when production funcSqrt is exited.
func (s *BasejossListener) ExitFuncSqrt(ctx *FuncSqrtContext) {}

// EnterFuncLog is called when production funcLog is entered.
func (s *BasejossListener) EnterFuncLog(ctx *FuncLogContext) {}

// ExitFuncLog is called when production funcLog is exited.
func (s *BasejossListener) ExitFuncLog(ctx *FuncLogContext) {}

// EnterFuncExp is called when production funcExp is entered.
func (s *BasejossListener) EnterFuncExp(ctx *FuncExpContext) {}

// ExitFuncExp is called when production funcExp is exited.
func (s *BasejossListener) ExitFuncExp(ctx *FuncExpContext) {}

// EnterFuncSin is called when production funcSin is entered.
func (s *BasejossListener) EnterFuncSin(ctx *FuncSinContext) {}

// ExitFuncSin is called when production funcSin is exited.
func (s *BasejossListener) ExitFuncSin(ctx *FuncSinContext) {}

// EnterFuncCos is called when production funcCos is entered.
func (s *BasejossListener) EnterFuncCos(ctx *FuncCosContext) {}

// ExitFuncCos is called when production funcCos is exited.
func (s *BasejossListener) ExitFuncCos(ctx *FuncCosContext) {}

// EnterFuncIp is called when production funcIp is entered.
func (s *BasejossListener) EnterFuncIp(ctx *FuncIpContext) {}

// ExitFuncIp is called when production funcIp is exited.
func (s *BasejossListener) ExitFuncIp(ctx *FuncIpContext) {}

// EnterFuncFp is called when production funcFp is entered.
func (s *BasejossListener) EnterFuncFp(ctx *FuncFpContext) {}

// ExitFuncFp is called when production funcFp is exited.
func (s *BasejossListener) ExitFuncFp(ctx *FuncFpContext) {}

// EnterFuncDp is called when production funcDp is entered.
func (s *BasejossListener) EnterFuncDp(ctx *FuncDpContext) {}

// ExitFuncDp is called when production funcDp is exited.
func (s *BasejossListener) ExitFuncDp(ctx *FuncDpContext) {}

// EnterFuncXp is called when production funcXp is entered.
func (s *BasejossListener) EnterFuncXp(ctx *FuncXpContext) {}

// ExitFuncXp is called when production funcXp is exited.
func (s *BasejossListener) ExitFuncXp(ctx *FuncXpContext) {}

// EnterFuncSgn is called when production funcSgn is entered.
func (s *BasejossListener) EnterFuncSgn(ctx *FuncSgnContext) {}

// ExitFuncSgn is called when production funcSgn is exited.
func (s *BasejossListener) ExitFuncSgn(ctx *FuncSgnContext) {}

// EnterFuncMax is called when production funcMax is entered.
func (s *BasejossListener) EnterFuncMax(ctx *FuncMaxContext) {}

// ExitFuncMax is called when production funcMax is exited.
func (s *BasejossListener) ExitFuncMax(ctx *FuncMaxContext) {}

// EnterFuncMin is called when production funcMin is entered.
func (s *BasejossListener) EnterFuncMin(ctx *FuncMinContext) {}

// ExitFuncMin is called when production funcMin is exited.
func (s *BasejossListener) ExitFuncMin(ctx *FuncMinContext) {}

// EnterArgList is called when production argList is entered.
func (s *BasejossListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BasejossListener) ExitArgList(ctx *ArgListContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasejossListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasejossListener) ExitVariable(ctx *VariableContext) {}

// EnterVarList is called when production varList is entered.
func (s *BasejossListener) EnterVarList(ctx *VarListContext) {}

// ExitVarList is called when production varList is exited.
func (s *BasejossListener) ExitVarList(ctx *VarListContext) {}
