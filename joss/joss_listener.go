// Code generated from joss.g4 by ANTLR 4.9.3. DO NOT EDIT.

package joss // joss
import "github.com/antlr/antlr4/runtime/Go/antlr"

// jossListener is a complete listener for a parse tree produced by jossParser.
type jossListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDirect is called when entering the direct production.
	EnterDirect(c *DirectContext)

	// EnterIndirect_ is called when entering the indirect_ production.
	EnterIndirect_(c *Indirect_Context)

	// EnterIndirectCMD is called when entering the indirectCMD production.
	EnterIndirectCMD(c *IndirectCMDContext)

	// EnterArbitraryCMD is called when entering the arbitraryCMD production.
	EnterArbitraryCMD(c *ArbitraryCMDContext)

	// EnterGoCMD is called when entering the goCMD production.
	EnterGoCMD(c *GoCMDContext)

	// EnterCancelCMD is called when entering the cancelCMD production.
	EnterCancelCMD(c *CancelCMDContext)

	// EnterDeleteCmdCMD is called when entering the deleteCmdCMD production.
	EnterDeleteCmdCMD(c *DeleteCmdCMDContext)

	// EnterDelCmdSel is called when entering the delCmdSel production.
	EnterDelCmdSel(c *DelCmdSelContext)

	// EnterDelCmdAllSel is called when entering the delCmdAllSel production.
	EnterDelCmdAllSel(c *DelCmdAllSelContext)

	// EnterDelCmdAllType is called when entering the delCmdAllType production.
	EnterDelCmdAllType(c *DelCmdAllTypeContext)

	// EnterDoneCMD is called when entering the doneCMD production.
	EnterDoneCMD(c *DoneCMDContext)

	// EnterStopCMD is called when entering the stopCMD production.
	EnterStopCMD(c *StopCMDContext)

	// EnterDemandCMD is called when entering the demandCMD production.
	EnterDemandCMD(c *DemandCMDContext)

	// EnterToCMD is called when entering the toCMD production.
	EnterToCMD(c *ToCMDContext)

	// EnterToSel is called when entering the toSel production.
	EnterToSel(c *ToSelContext)

	// EnterSetCMD is called when entering the setCMD production.
	EnterSetCMD(c *SetCMDContext)

	// EnterDoCMD is called when entering the doCMD production.
	EnterDoCMD(c *DoCMDContext)

	// EnterDoSel is called when entering the doSel production.
	EnterDoSel(c *DoSelContext)

	// EnterTypeCMD is called when entering the typeCMD production.
	EnterTypeCMD(c *TypeCMDContext)

	// EnterTypeContent is called when entering the typeContent production.
	EnterTypeContent(c *TypeContentContext)

	// EnterTypeSys is called when entering the typeSys production.
	EnterTypeSys(c *TypeSysContext)

	// EnterTypeForm is called when entering the typeForm production.
	EnterTypeForm(c *TypeFormContext)

	// EnterTypeData is called when entering the typeData production.
	EnterTypeData(c *TypeDataContext)

	// EnterTypeElem is called when entering the typeElem production.
	EnterTypeElem(c *TypeElemContext)

	// EnterTypeSel is called when entering the typeSel production.
	EnterTypeSel(c *TypeSelContext)

	// EnterTypeAllSel is called when entering the typeAllSel production.
	EnterTypeAllSel(c *TypeAllSelContext)

	// EnterTypeAllType is called when entering the typeAllType production.
	EnterTypeAllType(c *TypeAllTypeContext)

	// EnterDeleteValCMD is called when entering the deleteValCMD production.
	EnterDeleteValCMD(c *DeleteValCMDContext)

	// EnterFormCMD is called when entering the formCMD production.
	EnterFormCMD(c *FormCMDContext)

	// EnterFormContent is called when entering the formContent production.
	EnterFormContent(c *FormContentContext)

	// EnterFormObject is called when entering the formObject production.
	EnterFormObject(c *FormObjectContext)

	// EnterFormPH is called when entering the formPH production.
	EnterFormPH(c *FormPHContext)

	// EnterSciNotation is called when entering the sciNotation production.
	EnterSciNotation(c *SciNotationContext)

	// EnterDot is called when entering the dot production.
	EnterDot(c *DotContext)

	// EnterFixedNotation is called when entering the fixedNotation production.
	EnterFixedNotation(c *FixedNotationContext)

	// EnterUScore is called when entering the uScore production.
	EnterUScore(c *UScoreContext)

	// EnterStepSel is called when entering the stepSel production.
	EnterStepSel(c *StepSelContext)

	// EnterPartSel is called when entering the partSel production.
	EnterPartSel(c *PartSelContext)

	// EnterFormSel is called when entering the formSel production.
	EnterFormSel(c *FormSelContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterIf_ is called when entering the if_ production.
	EnterIf_(c *If_Context)

	// EnterBoolExp is called when entering the boolExp production.
	EnterBoolExp(c *BoolExpContext)

	// EnterBoolOp is called when entering the boolOp production.
	EnterBoolOp(c *BoolOpContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterBoolComp is called when entering the boolComp production.
	EnterBoolComp(c *BoolCompContext)

	// EnterFor_ is called when entering the for_ production.
	EnterFor_(c *For_Context)

	// EnterRange_ is called when entering the range_ production.
	EnterRange_(c *Range_Context)

	// EnterRangeExp is called when entering the rangeExp production.
	EnterRangeExp(c *RangeExpContext)

	// EnterRangeVal is called when entering the rangeVal production.
	EnterRangeVal(c *RangeValContext)

	// EnterMathExp is called when entering the mathExp production.
	EnterMathExp(c *MathExpContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterFunction_ is called when entering the function_ production.
	EnterFunction_(c *Function_Context)

	// EnterFuncSqrt is called when entering the funcSqrt production.
	EnterFuncSqrt(c *FuncSqrtContext)

	// EnterFuncLog is called when entering the funcLog production.
	EnterFuncLog(c *FuncLogContext)

	// EnterFuncExp is called when entering the funcExp production.
	EnterFuncExp(c *FuncExpContext)

	// EnterFuncSin is called when entering the funcSin production.
	EnterFuncSin(c *FuncSinContext)

	// EnterFuncCos is called when entering the funcCos production.
	EnterFuncCos(c *FuncCosContext)

	// EnterFuncIp is called when entering the funcIp production.
	EnterFuncIp(c *FuncIpContext)

	// EnterFuncFp is called when entering the funcFp production.
	EnterFuncFp(c *FuncFpContext)

	// EnterFuncDp is called when entering the funcDp production.
	EnterFuncDp(c *FuncDpContext)

	// EnterFuncXp is called when entering the funcXp production.
	EnterFuncXp(c *FuncXpContext)

	// EnterFuncSgn is called when entering the funcSgn production.
	EnterFuncSgn(c *FuncSgnContext)

	// EnterFuncMax is called when entering the funcMax production.
	EnterFuncMax(c *FuncMaxContext)

	// EnterFuncMin is called when entering the funcMin production.
	EnterFuncMin(c *FuncMinContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterVarList is called when entering the varList production.
	EnterVarList(c *VarListContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDirect is called when exiting the direct production.
	ExitDirect(c *DirectContext)

	// ExitIndirect_ is called when exiting the indirect_ production.
	ExitIndirect_(c *Indirect_Context)

	// ExitIndirectCMD is called when exiting the indirectCMD production.
	ExitIndirectCMD(c *IndirectCMDContext)

	// ExitArbitraryCMD is called when exiting the arbitraryCMD production.
	ExitArbitraryCMD(c *ArbitraryCMDContext)

	// ExitGoCMD is called when exiting the goCMD production.
	ExitGoCMD(c *GoCMDContext)

	// ExitCancelCMD is called when exiting the cancelCMD production.
	ExitCancelCMD(c *CancelCMDContext)

	// ExitDeleteCmdCMD is called when exiting the deleteCmdCMD production.
	ExitDeleteCmdCMD(c *DeleteCmdCMDContext)

	// ExitDelCmdSel is called when exiting the delCmdSel production.
	ExitDelCmdSel(c *DelCmdSelContext)

	// ExitDelCmdAllSel is called when exiting the delCmdAllSel production.
	ExitDelCmdAllSel(c *DelCmdAllSelContext)

	// ExitDelCmdAllType is called when exiting the delCmdAllType production.
	ExitDelCmdAllType(c *DelCmdAllTypeContext)

	// ExitDoneCMD is called when exiting the doneCMD production.
	ExitDoneCMD(c *DoneCMDContext)

	// ExitStopCMD is called when exiting the stopCMD production.
	ExitStopCMD(c *StopCMDContext)

	// ExitDemandCMD is called when exiting the demandCMD production.
	ExitDemandCMD(c *DemandCMDContext)

	// ExitToCMD is called when exiting the toCMD production.
	ExitToCMD(c *ToCMDContext)

	// ExitToSel is called when exiting the toSel production.
	ExitToSel(c *ToSelContext)

	// ExitSetCMD is called when exiting the setCMD production.
	ExitSetCMD(c *SetCMDContext)

	// ExitDoCMD is called when exiting the doCMD production.
	ExitDoCMD(c *DoCMDContext)

	// ExitDoSel is called when exiting the doSel production.
	ExitDoSel(c *DoSelContext)

	// ExitTypeCMD is called when exiting the typeCMD production.
	ExitTypeCMD(c *TypeCMDContext)

	// ExitTypeContent is called when exiting the typeContent production.
	ExitTypeContent(c *TypeContentContext)

	// ExitTypeSys is called when exiting the typeSys production.
	ExitTypeSys(c *TypeSysContext)

	// ExitTypeForm is called when exiting the typeForm production.
	ExitTypeForm(c *TypeFormContext)

	// ExitTypeData is called when exiting the typeData production.
	ExitTypeData(c *TypeDataContext)

	// ExitTypeElem is called when exiting the typeElem production.
	ExitTypeElem(c *TypeElemContext)

	// ExitTypeSel is called when exiting the typeSel production.
	ExitTypeSel(c *TypeSelContext)

	// ExitTypeAllSel is called when exiting the typeAllSel production.
	ExitTypeAllSel(c *TypeAllSelContext)

	// ExitTypeAllType is called when exiting the typeAllType production.
	ExitTypeAllType(c *TypeAllTypeContext)

	// ExitDeleteValCMD is called when exiting the deleteValCMD production.
	ExitDeleteValCMD(c *DeleteValCMDContext)

	// ExitFormCMD is called when exiting the formCMD production.
	ExitFormCMD(c *FormCMDContext)

	// ExitFormContent is called when exiting the formContent production.
	ExitFormContent(c *FormContentContext)

	// ExitFormObject is called when exiting the formObject production.
	ExitFormObject(c *FormObjectContext)

	// ExitFormPH is called when exiting the formPH production.
	ExitFormPH(c *FormPHContext)

	// ExitSciNotation is called when exiting the sciNotation production.
	ExitSciNotation(c *SciNotationContext)

	// ExitDot is called when exiting the dot production.
	ExitDot(c *DotContext)

	// ExitFixedNotation is called when exiting the fixedNotation production.
	ExitFixedNotation(c *FixedNotationContext)

	// ExitUScore is called when exiting the uScore production.
	ExitUScore(c *UScoreContext)

	// ExitStepSel is called when exiting the stepSel production.
	ExitStepSel(c *StepSelContext)

	// ExitPartSel is called when exiting the partSel production.
	ExitPartSel(c *PartSelContext)

	// ExitFormSel is called when exiting the formSel production.
	ExitFormSel(c *FormSelContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitIf_ is called when exiting the if_ production.
	ExitIf_(c *If_Context)

	// ExitBoolExp is called when exiting the boolExp production.
	ExitBoolExp(c *BoolExpContext)

	// ExitBoolOp is called when exiting the boolOp production.
	ExitBoolOp(c *BoolOpContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitBoolComp is called when exiting the boolComp production.
	ExitBoolComp(c *BoolCompContext)

	// ExitFor_ is called when exiting the for_ production.
	ExitFor_(c *For_Context)

	// ExitRange_ is called when exiting the range_ production.
	ExitRange_(c *Range_Context)

	// ExitRangeExp is called when exiting the rangeExp production.
	ExitRangeExp(c *RangeExpContext)

	// ExitRangeVal is called when exiting the rangeVal production.
	ExitRangeVal(c *RangeValContext)

	// ExitMathExp is called when exiting the mathExp production.
	ExitMathExp(c *MathExpContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitFunction_ is called when exiting the function_ production.
	ExitFunction_(c *Function_Context)

	// ExitFuncSqrt is called when exiting the funcSqrt production.
	ExitFuncSqrt(c *FuncSqrtContext)

	// ExitFuncLog is called when exiting the funcLog production.
	ExitFuncLog(c *FuncLogContext)

	// ExitFuncExp is called when exiting the funcExp production.
	ExitFuncExp(c *FuncExpContext)

	// ExitFuncSin is called when exiting the funcSin production.
	ExitFuncSin(c *FuncSinContext)

	// ExitFuncCos is called when exiting the funcCos production.
	ExitFuncCos(c *FuncCosContext)

	// ExitFuncIp is called when exiting the funcIp production.
	ExitFuncIp(c *FuncIpContext)

	// ExitFuncFp is called when exiting the funcFp production.
	ExitFuncFp(c *FuncFpContext)

	// ExitFuncDp is called when exiting the funcDp production.
	ExitFuncDp(c *FuncDpContext)

	// ExitFuncXp is called when exiting the funcXp production.
	ExitFuncXp(c *FuncXpContext)

	// ExitFuncSgn is called when exiting the funcSgn production.
	ExitFuncSgn(c *FuncSgnContext)

	// ExitFuncMax is called when exiting the funcMax production.
	ExitFuncMax(c *FuncMaxContext)

	// ExitFuncMin is called when exiting the funcMin production.
	ExitFuncMin(c *FuncMinContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitVarList is called when exiting the varList production.
	ExitVarList(c *VarListContext)
}
