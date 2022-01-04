// Code generated from vba.g4 by ANTLR 4.9.3. DO NOT EDIT.

package vba // vba
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasevbaListener is a complete listener for a parse tree produced by vbaParser.
type BasevbaListener struct{}

var _ vbaListener = &BasevbaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasevbaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasevbaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasevbaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasevbaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStartRule is called when production startRule is entered.
func (s *BasevbaListener) EnterStartRule(ctx *StartRuleContext) {}

// ExitStartRule is called when production startRule is exited.
func (s *BasevbaListener) ExitStartRule(ctx *StartRuleContext) {}

// EnterModule is called when production module is entered.
func (s *BasevbaListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BasevbaListener) ExitModule(ctx *ModuleContext) {}

// EnterModuleHeader is called when production moduleHeader is entered.
func (s *BasevbaListener) EnterModuleHeader(ctx *ModuleHeaderContext) {}

// ExitModuleHeader is called when production moduleHeader is exited.
func (s *BasevbaListener) ExitModuleHeader(ctx *ModuleHeaderContext) {}

// EnterModuleConfig is called when production moduleConfig is entered.
func (s *BasevbaListener) EnterModuleConfig(ctx *ModuleConfigContext) {}

// ExitModuleConfig is called when production moduleConfig is exited.
func (s *BasevbaListener) ExitModuleConfig(ctx *ModuleConfigContext) {}

// EnterModuleConfigElement is called when production moduleConfigElement is entered.
func (s *BasevbaListener) EnterModuleConfigElement(ctx *ModuleConfigElementContext) {}

// ExitModuleConfigElement is called when production moduleConfigElement is exited.
func (s *BasevbaListener) ExitModuleConfigElement(ctx *ModuleConfigElementContext) {}

// EnterModuleAttributes is called when production moduleAttributes is entered.
func (s *BasevbaListener) EnterModuleAttributes(ctx *ModuleAttributesContext) {}

// ExitModuleAttributes is called when production moduleAttributes is exited.
func (s *BasevbaListener) ExitModuleAttributes(ctx *ModuleAttributesContext) {}

// EnterModuleDeclarations is called when production moduleDeclarations is entered.
func (s *BasevbaListener) EnterModuleDeclarations(ctx *ModuleDeclarationsContext) {}

// ExitModuleDeclarations is called when production moduleDeclarations is exited.
func (s *BasevbaListener) ExitModuleDeclarations(ctx *ModuleDeclarationsContext) {}

// EnterOptionBaseStmt is called when production optionBaseStmt is entered.
func (s *BasevbaListener) EnterOptionBaseStmt(ctx *OptionBaseStmtContext) {}

// ExitOptionBaseStmt is called when production optionBaseStmt is exited.
func (s *BasevbaListener) ExitOptionBaseStmt(ctx *OptionBaseStmtContext) {}

// EnterOptionCompareStmt is called when production optionCompareStmt is entered.
func (s *BasevbaListener) EnterOptionCompareStmt(ctx *OptionCompareStmtContext) {}

// ExitOptionCompareStmt is called when production optionCompareStmt is exited.
func (s *BasevbaListener) ExitOptionCompareStmt(ctx *OptionCompareStmtContext) {}

// EnterOptionExplicitStmt is called when production optionExplicitStmt is entered.
func (s *BasevbaListener) EnterOptionExplicitStmt(ctx *OptionExplicitStmtContext) {}

// ExitOptionExplicitStmt is called when production optionExplicitStmt is exited.
func (s *BasevbaListener) ExitOptionExplicitStmt(ctx *OptionExplicitStmtContext) {}

// EnterOptionPrivateModuleStmt is called when production optionPrivateModuleStmt is entered.
func (s *BasevbaListener) EnterOptionPrivateModuleStmt(ctx *OptionPrivateModuleStmtContext) {}

// ExitOptionPrivateModuleStmt is called when production optionPrivateModuleStmt is exited.
func (s *BasevbaListener) ExitOptionPrivateModuleStmt(ctx *OptionPrivateModuleStmtContext) {}

// EnterModuleDeclarationsElement is called when production moduleDeclarationsElement is entered.
func (s *BasevbaListener) EnterModuleDeclarationsElement(ctx *ModuleDeclarationsElementContext) {}

// ExitModuleDeclarationsElement is called when production moduleDeclarationsElement is exited.
func (s *BasevbaListener) ExitModuleDeclarationsElement(ctx *ModuleDeclarationsElementContext) {}

// EnterMacroStmt is called when production macroStmt is entered.
func (s *BasevbaListener) EnterMacroStmt(ctx *MacroStmtContext) {}

// ExitMacroStmt is called when production macroStmt is exited.
func (s *BasevbaListener) ExitMacroStmt(ctx *MacroStmtContext) {}

// EnterModuleBody is called when production moduleBody is entered.
func (s *BasevbaListener) EnterModuleBody(ctx *ModuleBodyContext) {}

// ExitModuleBody is called when production moduleBody is exited.
func (s *BasevbaListener) ExitModuleBody(ctx *ModuleBodyContext) {}

// EnterModuleBodyElement is called when production moduleBodyElement is entered.
func (s *BasevbaListener) EnterModuleBodyElement(ctx *ModuleBodyElementContext) {}

// ExitModuleBodyElement is called when production moduleBodyElement is exited.
func (s *BasevbaListener) ExitModuleBodyElement(ctx *ModuleBodyElementContext) {}

// EnterAttributeStmt is called when production attributeStmt is entered.
func (s *BasevbaListener) EnterAttributeStmt(ctx *AttributeStmtContext) {}

// ExitAttributeStmt is called when production attributeStmt is exited.
func (s *BasevbaListener) ExitAttributeStmt(ctx *AttributeStmtContext) {}

// EnterBlock is called when production block is entered.
func (s *BasevbaListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasevbaListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStmt is called when production blockStmt is entered.
func (s *BasevbaListener) EnterBlockStmt(ctx *BlockStmtContext) {}

// ExitBlockStmt is called when production blockStmt is exited.
func (s *BasevbaListener) ExitBlockStmt(ctx *BlockStmtContext) {}

// EnterAppactivateStmt is called when production appactivateStmt is entered.
func (s *BasevbaListener) EnterAppactivateStmt(ctx *AppactivateStmtContext) {}

// ExitAppactivateStmt is called when production appactivateStmt is exited.
func (s *BasevbaListener) ExitAppactivateStmt(ctx *AppactivateStmtContext) {}

// EnterBeepStmt is called when production beepStmt is entered.
func (s *BasevbaListener) EnterBeepStmt(ctx *BeepStmtContext) {}

// ExitBeepStmt is called when production beepStmt is exited.
func (s *BasevbaListener) ExitBeepStmt(ctx *BeepStmtContext) {}

// EnterChdirStmt is called when production chdirStmt is entered.
func (s *BasevbaListener) EnterChdirStmt(ctx *ChdirStmtContext) {}

// ExitChdirStmt is called when production chdirStmt is exited.
func (s *BasevbaListener) ExitChdirStmt(ctx *ChdirStmtContext) {}

// EnterChdriveStmt is called when production chdriveStmt is entered.
func (s *BasevbaListener) EnterChdriveStmt(ctx *ChdriveStmtContext) {}

// ExitChdriveStmt is called when production chdriveStmt is exited.
func (s *BasevbaListener) ExitChdriveStmt(ctx *ChdriveStmtContext) {}

// EnterCloseStmt is called when production closeStmt is entered.
func (s *BasevbaListener) EnterCloseStmt(ctx *CloseStmtContext) {}

// ExitCloseStmt is called when production closeStmt is exited.
func (s *BasevbaListener) ExitCloseStmt(ctx *CloseStmtContext) {}

// EnterConstStmt is called when production constStmt is entered.
func (s *BasevbaListener) EnterConstStmt(ctx *ConstStmtContext) {}

// ExitConstStmt is called when production constStmt is exited.
func (s *BasevbaListener) ExitConstStmt(ctx *ConstStmtContext) {}

// EnterConstSubStmt is called when production constSubStmt is entered.
func (s *BasevbaListener) EnterConstSubStmt(ctx *ConstSubStmtContext) {}

// ExitConstSubStmt is called when production constSubStmt is exited.
func (s *BasevbaListener) ExitConstSubStmt(ctx *ConstSubStmtContext) {}

// EnterDateStmt is called when production dateStmt is entered.
func (s *BasevbaListener) EnterDateStmt(ctx *DateStmtContext) {}

// ExitDateStmt is called when production dateStmt is exited.
func (s *BasevbaListener) ExitDateStmt(ctx *DateStmtContext) {}

// EnterDeclareStmt is called when production declareStmt is entered.
func (s *BasevbaListener) EnterDeclareStmt(ctx *DeclareStmtContext) {}

// ExitDeclareStmt is called when production declareStmt is exited.
func (s *BasevbaListener) ExitDeclareStmt(ctx *DeclareStmtContext) {}

// EnterDeftypeStmt is called when production deftypeStmt is entered.
func (s *BasevbaListener) EnterDeftypeStmt(ctx *DeftypeStmtContext) {}

// ExitDeftypeStmt is called when production deftypeStmt is exited.
func (s *BasevbaListener) ExitDeftypeStmt(ctx *DeftypeStmtContext) {}

// EnterDeleteSettingStmt is called when production deleteSettingStmt is entered.
func (s *BasevbaListener) EnterDeleteSettingStmt(ctx *DeleteSettingStmtContext) {}

// ExitDeleteSettingStmt is called when production deleteSettingStmt is exited.
func (s *BasevbaListener) ExitDeleteSettingStmt(ctx *DeleteSettingStmtContext) {}

// EnterDoLoopStmt is called when production doLoopStmt is entered.
func (s *BasevbaListener) EnterDoLoopStmt(ctx *DoLoopStmtContext) {}

// ExitDoLoopStmt is called when production doLoopStmt is exited.
func (s *BasevbaListener) ExitDoLoopStmt(ctx *DoLoopStmtContext) {}

// EnterEndStmt is called when production endStmt is entered.
func (s *BasevbaListener) EnterEndStmt(ctx *EndStmtContext) {}

// ExitEndStmt is called when production endStmt is exited.
func (s *BasevbaListener) ExitEndStmt(ctx *EndStmtContext) {}

// EnterEnumerationStmt is called when production enumerationStmt is entered.
func (s *BasevbaListener) EnterEnumerationStmt(ctx *EnumerationStmtContext) {}

// ExitEnumerationStmt is called when production enumerationStmt is exited.
func (s *BasevbaListener) ExitEnumerationStmt(ctx *EnumerationStmtContext) {}

// EnterEnumerationStmt_Constant is called when production enumerationStmt_Constant is entered.
func (s *BasevbaListener) EnterEnumerationStmt_Constant(ctx *EnumerationStmt_ConstantContext) {}

// ExitEnumerationStmt_Constant is called when production enumerationStmt_Constant is exited.
func (s *BasevbaListener) ExitEnumerationStmt_Constant(ctx *EnumerationStmt_ConstantContext) {}

// EnterEraseStmt is called when production eraseStmt is entered.
func (s *BasevbaListener) EnterEraseStmt(ctx *EraseStmtContext) {}

// ExitEraseStmt is called when production eraseStmt is exited.
func (s *BasevbaListener) ExitEraseStmt(ctx *EraseStmtContext) {}

// EnterErrorStmt is called when production errorStmt is entered.
func (s *BasevbaListener) EnterErrorStmt(ctx *ErrorStmtContext) {}

// ExitErrorStmt is called when production errorStmt is exited.
func (s *BasevbaListener) ExitErrorStmt(ctx *ErrorStmtContext) {}

// EnterEventStmt is called when production eventStmt is entered.
func (s *BasevbaListener) EnterEventStmt(ctx *EventStmtContext) {}

// ExitEventStmt is called when production eventStmt is exited.
func (s *BasevbaListener) ExitEventStmt(ctx *EventStmtContext) {}

// EnterExitStmt is called when production exitStmt is entered.
func (s *BasevbaListener) EnterExitStmt(ctx *ExitStmtContext) {}

// ExitExitStmt is called when production exitStmt is exited.
func (s *BasevbaListener) ExitExitStmt(ctx *ExitStmtContext) {}

// EnterFilecopyStmt is called when production filecopyStmt is entered.
func (s *BasevbaListener) EnterFilecopyStmt(ctx *FilecopyStmtContext) {}

// ExitFilecopyStmt is called when production filecopyStmt is exited.
func (s *BasevbaListener) ExitFilecopyStmt(ctx *FilecopyStmtContext) {}

// EnterForEachStmt is called when production forEachStmt is entered.
func (s *BasevbaListener) EnterForEachStmt(ctx *ForEachStmtContext) {}

// ExitForEachStmt is called when production forEachStmt is exited.
func (s *BasevbaListener) ExitForEachStmt(ctx *ForEachStmtContext) {}

// EnterForNextStmt is called when production forNextStmt is entered.
func (s *BasevbaListener) EnterForNextStmt(ctx *ForNextStmtContext) {}

// ExitForNextStmt is called when production forNextStmt is exited.
func (s *BasevbaListener) ExitForNextStmt(ctx *ForNextStmtContext) {}

// EnterFunctionStmt is called when production functionStmt is entered.
func (s *BasevbaListener) EnterFunctionStmt(ctx *FunctionStmtContext) {}

// ExitFunctionStmt is called when production functionStmt is exited.
func (s *BasevbaListener) ExitFunctionStmt(ctx *FunctionStmtContext) {}

// EnterGetStmt is called when production getStmt is entered.
func (s *BasevbaListener) EnterGetStmt(ctx *GetStmtContext) {}

// ExitGetStmt is called when production getStmt is exited.
func (s *BasevbaListener) ExitGetStmt(ctx *GetStmtContext) {}

// EnterGoSubStmt is called when production goSubStmt is entered.
func (s *BasevbaListener) EnterGoSubStmt(ctx *GoSubStmtContext) {}

// ExitGoSubStmt is called when production goSubStmt is exited.
func (s *BasevbaListener) ExitGoSubStmt(ctx *GoSubStmtContext) {}

// EnterGoToStmt is called when production goToStmt is entered.
func (s *BasevbaListener) EnterGoToStmt(ctx *GoToStmtContext) {}

// ExitGoToStmt is called when production goToStmt is exited.
func (s *BasevbaListener) ExitGoToStmt(ctx *GoToStmtContext) {}

// EnterInlineIfThenElse is called when production inlineIfThenElse is entered.
func (s *BasevbaListener) EnterInlineIfThenElse(ctx *InlineIfThenElseContext) {}

// ExitInlineIfThenElse is called when production inlineIfThenElse is exited.
func (s *BasevbaListener) ExitInlineIfThenElse(ctx *InlineIfThenElseContext) {}

// EnterBlockIfThenElse is called when production blockIfThenElse is entered.
func (s *BasevbaListener) EnterBlockIfThenElse(ctx *BlockIfThenElseContext) {}

// ExitBlockIfThenElse is called when production blockIfThenElse is exited.
func (s *BasevbaListener) ExitBlockIfThenElse(ctx *BlockIfThenElseContext) {}

// EnterIfBlockStmt is called when production ifBlockStmt is entered.
func (s *BasevbaListener) EnterIfBlockStmt(ctx *IfBlockStmtContext) {}

// ExitIfBlockStmt is called when production ifBlockStmt is exited.
func (s *BasevbaListener) ExitIfBlockStmt(ctx *IfBlockStmtContext) {}

// EnterIfConditionStmt is called when production ifConditionStmt is entered.
func (s *BasevbaListener) EnterIfConditionStmt(ctx *IfConditionStmtContext) {}

// ExitIfConditionStmt is called when production ifConditionStmt is exited.
func (s *BasevbaListener) ExitIfConditionStmt(ctx *IfConditionStmtContext) {}

// EnterIfElseIfBlockStmt is called when production ifElseIfBlockStmt is entered.
func (s *BasevbaListener) EnterIfElseIfBlockStmt(ctx *IfElseIfBlockStmtContext) {}

// ExitIfElseIfBlockStmt is called when production ifElseIfBlockStmt is exited.
func (s *BasevbaListener) ExitIfElseIfBlockStmt(ctx *IfElseIfBlockStmtContext) {}

// EnterIfElseBlockStmt is called when production ifElseBlockStmt is entered.
func (s *BasevbaListener) EnterIfElseBlockStmt(ctx *IfElseBlockStmtContext) {}

// ExitIfElseBlockStmt is called when production ifElseBlockStmt is exited.
func (s *BasevbaListener) ExitIfElseBlockStmt(ctx *IfElseBlockStmtContext) {}

// EnterImplementsStmt is called when production implementsStmt is entered.
func (s *BasevbaListener) EnterImplementsStmt(ctx *ImplementsStmtContext) {}

// ExitImplementsStmt is called when production implementsStmt is exited.
func (s *BasevbaListener) ExitImplementsStmt(ctx *ImplementsStmtContext) {}

// EnterInputStmt is called when production inputStmt is entered.
func (s *BasevbaListener) EnterInputStmt(ctx *InputStmtContext) {}

// ExitInputStmt is called when production inputStmt is exited.
func (s *BasevbaListener) ExitInputStmt(ctx *InputStmtContext) {}

// EnterKillStmt is called when production killStmt is entered.
func (s *BasevbaListener) EnterKillStmt(ctx *KillStmtContext) {}

// ExitKillStmt is called when production killStmt is exited.
func (s *BasevbaListener) ExitKillStmt(ctx *KillStmtContext) {}

// EnterLetStmt is called when production letStmt is entered.
func (s *BasevbaListener) EnterLetStmt(ctx *LetStmtContext) {}

// ExitLetStmt is called when production letStmt is exited.
func (s *BasevbaListener) ExitLetStmt(ctx *LetStmtContext) {}

// EnterLineInputStmt is called when production lineInputStmt is entered.
func (s *BasevbaListener) EnterLineInputStmt(ctx *LineInputStmtContext) {}

// ExitLineInputStmt is called when production lineInputStmt is exited.
func (s *BasevbaListener) ExitLineInputStmt(ctx *LineInputStmtContext) {}

// EnterLoadStmt is called when production loadStmt is entered.
func (s *BasevbaListener) EnterLoadStmt(ctx *LoadStmtContext) {}

// ExitLoadStmt is called when production loadStmt is exited.
func (s *BasevbaListener) ExitLoadStmt(ctx *LoadStmtContext) {}

// EnterLockStmt is called when production lockStmt is entered.
func (s *BasevbaListener) EnterLockStmt(ctx *LockStmtContext) {}

// ExitLockStmt is called when production lockStmt is exited.
func (s *BasevbaListener) ExitLockStmt(ctx *LockStmtContext) {}

// EnterLsetStmt is called when production lsetStmt is entered.
func (s *BasevbaListener) EnterLsetStmt(ctx *LsetStmtContext) {}

// ExitLsetStmt is called when production lsetStmt is exited.
func (s *BasevbaListener) ExitLsetStmt(ctx *LsetStmtContext) {}

// EnterMacroConstStmt is called when production macroConstStmt is entered.
func (s *BasevbaListener) EnterMacroConstStmt(ctx *MacroConstStmtContext) {}

// ExitMacroConstStmt is called when production macroConstStmt is exited.
func (s *BasevbaListener) ExitMacroConstStmt(ctx *MacroConstStmtContext) {}

// EnterMacroIfThenElseStmt is called when production macroIfThenElseStmt is entered.
func (s *BasevbaListener) EnterMacroIfThenElseStmt(ctx *MacroIfThenElseStmtContext) {}

// ExitMacroIfThenElseStmt is called when production macroIfThenElseStmt is exited.
func (s *BasevbaListener) ExitMacroIfThenElseStmt(ctx *MacroIfThenElseStmtContext) {}

// EnterMacroIfBlockStmt is called when production macroIfBlockStmt is entered.
func (s *BasevbaListener) EnterMacroIfBlockStmt(ctx *MacroIfBlockStmtContext) {}

// ExitMacroIfBlockStmt is called when production macroIfBlockStmt is exited.
func (s *BasevbaListener) ExitMacroIfBlockStmt(ctx *MacroIfBlockStmtContext) {}

// EnterMacroElseIfBlockStmt is called when production macroElseIfBlockStmt is entered.
func (s *BasevbaListener) EnterMacroElseIfBlockStmt(ctx *MacroElseIfBlockStmtContext) {}

// ExitMacroElseIfBlockStmt is called when production macroElseIfBlockStmt is exited.
func (s *BasevbaListener) ExitMacroElseIfBlockStmt(ctx *MacroElseIfBlockStmtContext) {}

// EnterMacroElseBlockStmt is called when production macroElseBlockStmt is entered.
func (s *BasevbaListener) EnterMacroElseBlockStmt(ctx *MacroElseBlockStmtContext) {}

// ExitMacroElseBlockStmt is called when production macroElseBlockStmt is exited.
func (s *BasevbaListener) ExitMacroElseBlockStmt(ctx *MacroElseBlockStmtContext) {}

// EnterMidStmt is called when production midStmt is entered.
func (s *BasevbaListener) EnterMidStmt(ctx *MidStmtContext) {}

// ExitMidStmt is called when production midStmt is exited.
func (s *BasevbaListener) ExitMidStmt(ctx *MidStmtContext) {}

// EnterMkdirStmt is called when production mkdirStmt is entered.
func (s *BasevbaListener) EnterMkdirStmt(ctx *MkdirStmtContext) {}

// ExitMkdirStmt is called when production mkdirStmt is exited.
func (s *BasevbaListener) ExitMkdirStmt(ctx *MkdirStmtContext) {}

// EnterNameStmt is called when production nameStmt is entered.
func (s *BasevbaListener) EnterNameStmt(ctx *NameStmtContext) {}

// ExitNameStmt is called when production nameStmt is exited.
func (s *BasevbaListener) ExitNameStmt(ctx *NameStmtContext) {}

// EnterOnErrorStmt is called when production onErrorStmt is entered.
func (s *BasevbaListener) EnterOnErrorStmt(ctx *OnErrorStmtContext) {}

// ExitOnErrorStmt is called when production onErrorStmt is exited.
func (s *BasevbaListener) ExitOnErrorStmt(ctx *OnErrorStmtContext) {}

// EnterOnGoToStmt is called when production onGoToStmt is entered.
func (s *BasevbaListener) EnterOnGoToStmt(ctx *OnGoToStmtContext) {}

// ExitOnGoToStmt is called when production onGoToStmt is exited.
func (s *BasevbaListener) ExitOnGoToStmt(ctx *OnGoToStmtContext) {}

// EnterOnGoSubStmt is called when production onGoSubStmt is entered.
func (s *BasevbaListener) EnterOnGoSubStmt(ctx *OnGoSubStmtContext) {}

// ExitOnGoSubStmt is called when production onGoSubStmt is exited.
func (s *BasevbaListener) ExitOnGoSubStmt(ctx *OnGoSubStmtContext) {}

// EnterOpenStmt is called when production openStmt is entered.
func (s *BasevbaListener) EnterOpenStmt(ctx *OpenStmtContext) {}

// ExitOpenStmt is called when production openStmt is exited.
func (s *BasevbaListener) ExitOpenStmt(ctx *OpenStmtContext) {}

// EnterOutputList is called when production outputList is entered.
func (s *BasevbaListener) EnterOutputList(ctx *OutputListContext) {}

// ExitOutputList is called when production outputList is exited.
func (s *BasevbaListener) ExitOutputList(ctx *OutputListContext) {}

// EnterOutputList_Expression is called when production outputList_Expression is entered.
func (s *BasevbaListener) EnterOutputList_Expression(ctx *OutputList_ExpressionContext) {}

// ExitOutputList_Expression is called when production outputList_Expression is exited.
func (s *BasevbaListener) ExitOutputList_Expression(ctx *OutputList_ExpressionContext) {}

// EnterPrintStmt is called when production printStmt is entered.
func (s *BasevbaListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production printStmt is exited.
func (s *BasevbaListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterPropertyGetStmt is called when production propertyGetStmt is entered.
func (s *BasevbaListener) EnterPropertyGetStmt(ctx *PropertyGetStmtContext) {}

// ExitPropertyGetStmt is called when production propertyGetStmt is exited.
func (s *BasevbaListener) ExitPropertyGetStmt(ctx *PropertyGetStmtContext) {}

// EnterPropertySetStmt is called when production propertySetStmt is entered.
func (s *BasevbaListener) EnterPropertySetStmt(ctx *PropertySetStmtContext) {}

// ExitPropertySetStmt is called when production propertySetStmt is exited.
func (s *BasevbaListener) ExitPropertySetStmt(ctx *PropertySetStmtContext) {}

// EnterPropertyLetStmt is called when production propertyLetStmt is entered.
func (s *BasevbaListener) EnterPropertyLetStmt(ctx *PropertyLetStmtContext) {}

// ExitPropertyLetStmt is called when production propertyLetStmt is exited.
func (s *BasevbaListener) ExitPropertyLetStmt(ctx *PropertyLetStmtContext) {}

// EnterPutStmt is called when production putStmt is entered.
func (s *BasevbaListener) EnterPutStmt(ctx *PutStmtContext) {}

// ExitPutStmt is called when production putStmt is exited.
func (s *BasevbaListener) ExitPutStmt(ctx *PutStmtContext) {}

// EnterRaiseEventStmt is called when production raiseEventStmt is entered.
func (s *BasevbaListener) EnterRaiseEventStmt(ctx *RaiseEventStmtContext) {}

// ExitRaiseEventStmt is called when production raiseEventStmt is exited.
func (s *BasevbaListener) ExitRaiseEventStmt(ctx *RaiseEventStmtContext) {}

// EnterRandomizeStmt is called when production randomizeStmt is entered.
func (s *BasevbaListener) EnterRandomizeStmt(ctx *RandomizeStmtContext) {}

// ExitRandomizeStmt is called when production randomizeStmt is exited.
func (s *BasevbaListener) ExitRandomizeStmt(ctx *RandomizeStmtContext) {}

// EnterRedimStmt is called when production redimStmt is entered.
func (s *BasevbaListener) EnterRedimStmt(ctx *RedimStmtContext) {}

// ExitRedimStmt is called when production redimStmt is exited.
func (s *BasevbaListener) ExitRedimStmt(ctx *RedimStmtContext) {}

// EnterRedimSubStmt is called when production redimSubStmt is entered.
func (s *BasevbaListener) EnterRedimSubStmt(ctx *RedimSubStmtContext) {}

// ExitRedimSubStmt is called when production redimSubStmt is exited.
func (s *BasevbaListener) ExitRedimSubStmt(ctx *RedimSubStmtContext) {}

// EnterResetStmt is called when production resetStmt is entered.
func (s *BasevbaListener) EnterResetStmt(ctx *ResetStmtContext) {}

// ExitResetStmt is called when production resetStmt is exited.
func (s *BasevbaListener) ExitResetStmt(ctx *ResetStmtContext) {}

// EnterResumeStmt is called when production resumeStmt is entered.
func (s *BasevbaListener) EnterResumeStmt(ctx *ResumeStmtContext) {}

// ExitResumeStmt is called when production resumeStmt is exited.
func (s *BasevbaListener) ExitResumeStmt(ctx *ResumeStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BasevbaListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BasevbaListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterRmdirStmt is called when production rmdirStmt is entered.
func (s *BasevbaListener) EnterRmdirStmt(ctx *RmdirStmtContext) {}

// ExitRmdirStmt is called when production rmdirStmt is exited.
func (s *BasevbaListener) ExitRmdirStmt(ctx *RmdirStmtContext) {}

// EnterRsetStmt is called when production rsetStmt is entered.
func (s *BasevbaListener) EnterRsetStmt(ctx *RsetStmtContext) {}

// ExitRsetStmt is called when production rsetStmt is exited.
func (s *BasevbaListener) ExitRsetStmt(ctx *RsetStmtContext) {}

// EnterSavepictureStmt is called when production savepictureStmt is entered.
func (s *BasevbaListener) EnterSavepictureStmt(ctx *SavepictureStmtContext) {}

// ExitSavepictureStmt is called when production savepictureStmt is exited.
func (s *BasevbaListener) ExitSavepictureStmt(ctx *SavepictureStmtContext) {}

// EnterSaveSettingStmt is called when production saveSettingStmt is entered.
func (s *BasevbaListener) EnterSaveSettingStmt(ctx *SaveSettingStmtContext) {}

// ExitSaveSettingStmt is called when production saveSettingStmt is exited.
func (s *BasevbaListener) ExitSaveSettingStmt(ctx *SaveSettingStmtContext) {}

// EnterSeekStmt is called when production seekStmt is entered.
func (s *BasevbaListener) EnterSeekStmt(ctx *SeekStmtContext) {}

// ExitSeekStmt is called when production seekStmt is exited.
func (s *BasevbaListener) ExitSeekStmt(ctx *SeekStmtContext) {}

// EnterSelectCaseStmt is called when production selectCaseStmt is entered.
func (s *BasevbaListener) EnterSelectCaseStmt(ctx *SelectCaseStmtContext) {}

// ExitSelectCaseStmt is called when production selectCaseStmt is exited.
func (s *BasevbaListener) ExitSelectCaseStmt(ctx *SelectCaseStmtContext) {}

// EnterCaseCondIs is called when production caseCondIs is entered.
func (s *BasevbaListener) EnterCaseCondIs(ctx *CaseCondIsContext) {}

// ExitCaseCondIs is called when production caseCondIs is exited.
func (s *BasevbaListener) ExitCaseCondIs(ctx *CaseCondIsContext) {}

// EnterCaseCondTo is called when production caseCondTo is entered.
func (s *BasevbaListener) EnterCaseCondTo(ctx *CaseCondToContext) {}

// ExitCaseCondTo is called when production caseCondTo is exited.
func (s *BasevbaListener) ExitCaseCondTo(ctx *CaseCondToContext) {}

// EnterCaseCondValue is called when production caseCondValue is entered.
func (s *BasevbaListener) EnterCaseCondValue(ctx *CaseCondValueContext) {}

// ExitCaseCondValue is called when production caseCondValue is exited.
func (s *BasevbaListener) ExitCaseCondValue(ctx *CaseCondValueContext) {}

// EnterSC_Case is called when production sC_Case is entered.
func (s *BasevbaListener) EnterSC_Case(ctx *SC_CaseContext) {}

// ExitSC_Case is called when production sC_Case is exited.
func (s *BasevbaListener) ExitSC_Case(ctx *SC_CaseContext) {}

// EnterCaseCondElse is called when production caseCondElse is entered.
func (s *BasevbaListener) EnterCaseCondElse(ctx *CaseCondElseContext) {}

// ExitCaseCondElse is called when production caseCondElse is exited.
func (s *BasevbaListener) ExitCaseCondElse(ctx *CaseCondElseContext) {}

// EnterCaseCondSelection is called when production caseCondSelection is entered.
func (s *BasevbaListener) EnterCaseCondSelection(ctx *CaseCondSelectionContext) {}

// ExitCaseCondSelection is called when production caseCondSelection is exited.
func (s *BasevbaListener) ExitCaseCondSelection(ctx *CaseCondSelectionContext) {}

// EnterSendkeysStmt is called when production sendkeysStmt is entered.
func (s *BasevbaListener) EnterSendkeysStmt(ctx *SendkeysStmtContext) {}

// ExitSendkeysStmt is called when production sendkeysStmt is exited.
func (s *BasevbaListener) ExitSendkeysStmt(ctx *SendkeysStmtContext) {}

// EnterSetattrStmt is called when production setattrStmt is entered.
func (s *BasevbaListener) EnterSetattrStmt(ctx *SetattrStmtContext) {}

// ExitSetattrStmt is called when production setattrStmt is exited.
func (s *BasevbaListener) ExitSetattrStmt(ctx *SetattrStmtContext) {}

// EnterSetStmt is called when production setStmt is entered.
func (s *BasevbaListener) EnterSetStmt(ctx *SetStmtContext) {}

// ExitSetStmt is called when production setStmt is exited.
func (s *BasevbaListener) ExitSetStmt(ctx *SetStmtContext) {}

// EnterStopStmt is called when production stopStmt is entered.
func (s *BasevbaListener) EnterStopStmt(ctx *StopStmtContext) {}

// ExitStopStmt is called when production stopStmt is exited.
func (s *BasevbaListener) ExitStopStmt(ctx *StopStmtContext) {}

// EnterSubStmt is called when production subStmt is entered.
func (s *BasevbaListener) EnterSubStmt(ctx *SubStmtContext) {}

// ExitSubStmt is called when production subStmt is exited.
func (s *BasevbaListener) ExitSubStmt(ctx *SubStmtContext) {}

// EnterTimeStmt is called when production timeStmt is entered.
func (s *BasevbaListener) EnterTimeStmt(ctx *TimeStmtContext) {}

// ExitTimeStmt is called when production timeStmt is exited.
func (s *BasevbaListener) ExitTimeStmt(ctx *TimeStmtContext) {}

// EnterTypeStmt is called when production typeStmt is entered.
func (s *BasevbaListener) EnterTypeStmt(ctx *TypeStmtContext) {}

// ExitTypeStmt is called when production typeStmt is exited.
func (s *BasevbaListener) ExitTypeStmt(ctx *TypeStmtContext) {}

// EnterTypeStmt_Element is called when production typeStmt_Element is entered.
func (s *BasevbaListener) EnterTypeStmt_Element(ctx *TypeStmt_ElementContext) {}

// ExitTypeStmt_Element is called when production typeStmt_Element is exited.
func (s *BasevbaListener) ExitTypeStmt_Element(ctx *TypeStmt_ElementContext) {}

// EnterTypeOfStmt is called when production typeOfStmt is entered.
func (s *BasevbaListener) EnterTypeOfStmt(ctx *TypeOfStmtContext) {}

// ExitTypeOfStmt is called when production typeOfStmt is exited.
func (s *BasevbaListener) ExitTypeOfStmt(ctx *TypeOfStmtContext) {}

// EnterUnloadStmt is called when production unloadStmt is entered.
func (s *BasevbaListener) EnterUnloadStmt(ctx *UnloadStmtContext) {}

// ExitUnloadStmt is called when production unloadStmt is exited.
func (s *BasevbaListener) ExitUnloadStmt(ctx *UnloadStmtContext) {}

// EnterUnlockStmt is called when production unlockStmt is entered.
func (s *BasevbaListener) EnterUnlockStmt(ctx *UnlockStmtContext) {}

// ExitUnlockStmt is called when production unlockStmt is exited.
func (s *BasevbaListener) ExitUnlockStmt(ctx *UnlockStmtContext) {}

// EnterVsStruct is called when production vsStruct is entered.
func (s *BasevbaListener) EnterVsStruct(ctx *VsStructContext) {}

// ExitVsStruct is called when production vsStruct is exited.
func (s *BasevbaListener) ExitVsStruct(ctx *VsStructContext) {}

// EnterVsAdd is called when production vsAdd is entered.
func (s *BasevbaListener) EnterVsAdd(ctx *VsAddContext) {}

// ExitVsAdd is called when production vsAdd is exited.
func (s *BasevbaListener) ExitVsAdd(ctx *VsAddContext) {}

// EnterVsLt is called when production vsLt is entered.
func (s *BasevbaListener) EnterVsLt(ctx *VsLtContext) {}

// ExitVsLt is called when production vsLt is exited.
func (s *BasevbaListener) ExitVsLt(ctx *VsLtContext) {}

// EnterVsAddressOf is called when production vsAddressOf is entered.
func (s *BasevbaListener) EnterVsAddressOf(ctx *VsAddressOfContext) {}

// ExitVsAddressOf is called when production vsAddressOf is exited.
func (s *BasevbaListener) ExitVsAddressOf(ctx *VsAddressOfContext) {}

// EnterVsNew is called when production vsNew is entered.
func (s *BasevbaListener) EnterVsNew(ctx *VsNewContext) {}

// ExitVsNew is called when production vsNew is exited.
func (s *BasevbaListener) ExitVsNew(ctx *VsNewContext) {}

// EnterVsMult is called when production vsMult is entered.
func (s *BasevbaListener) EnterVsMult(ctx *VsMultContext) {}

// ExitVsMult is called when production vsMult is exited.
func (s *BasevbaListener) ExitVsMult(ctx *VsMultContext) {}

// EnterVsNegation is called when production vsNegation is entered.
func (s *BasevbaListener) EnterVsNegation(ctx *VsNegationContext) {}

// ExitVsNegation is called when production vsNegation is exited.
func (s *BasevbaListener) ExitVsNegation(ctx *VsNegationContext) {}

// EnterVsAssign is called when production vsAssign is entered.
func (s *BasevbaListener) EnterVsAssign(ctx *VsAssignContext) {}

// ExitVsAssign is called when production vsAssign is exited.
func (s *BasevbaListener) ExitVsAssign(ctx *VsAssignContext) {}

// EnterVsLike is called when production vsLike is entered.
func (s *BasevbaListener) EnterVsLike(ctx *VsLikeContext) {}

// ExitVsLike is called when production vsLike is exited.
func (s *BasevbaListener) ExitVsLike(ctx *VsLikeContext) {}

// EnterVsDiv is called when production vsDiv is entered.
func (s *BasevbaListener) EnterVsDiv(ctx *VsDivContext) {}

// ExitVsDiv is called when production vsDiv is exited.
func (s *BasevbaListener) ExitVsDiv(ctx *VsDivContext) {}

// EnterVsPlus is called when production vsPlus is entered.
func (s *BasevbaListener) EnterVsPlus(ctx *VsPlusContext) {}

// ExitVsPlus is called when production vsPlus is exited.
func (s *BasevbaListener) ExitVsPlus(ctx *VsPlusContext) {}

// EnterVsNot is called when production vsNot is entered.
func (s *BasevbaListener) EnterVsNot(ctx *VsNotContext) {}

// ExitVsNot is called when production vsNot is exited.
func (s *BasevbaListener) ExitVsNot(ctx *VsNotContext) {}

// EnterVsGeq is called when production vsGeq is entered.
func (s *BasevbaListener) EnterVsGeq(ctx *VsGeqContext) {}

// ExitVsGeq is called when production vsGeq is exited.
func (s *BasevbaListener) ExitVsGeq(ctx *VsGeqContext) {}

// EnterVsTypeOf is called when production vsTypeOf is entered.
func (s *BasevbaListener) EnterVsTypeOf(ctx *VsTypeOfContext) {}

// ExitVsTypeOf is called when production vsTypeOf is exited.
func (s *BasevbaListener) ExitVsTypeOf(ctx *VsTypeOfContext) {}

// EnterVsICS is called when production vsICS is entered.
func (s *BasevbaListener) EnterVsICS(ctx *VsICSContext) {}

// ExitVsICS is called when production vsICS is exited.
func (s *BasevbaListener) ExitVsICS(ctx *VsICSContext) {}

// EnterVsNeq is called when production vsNeq is entered.
func (s *BasevbaListener) EnterVsNeq(ctx *VsNeqContext) {}

// ExitVsNeq is called when production vsNeq is exited.
func (s *BasevbaListener) ExitVsNeq(ctx *VsNeqContext) {}

// EnterVsXor is called when production vsXor is entered.
func (s *BasevbaListener) EnterVsXor(ctx *VsXorContext) {}

// ExitVsXor is called when production vsXor is exited.
func (s *BasevbaListener) ExitVsXor(ctx *VsXorContext) {}

// EnterVsAnd is called when production vsAnd is entered.
func (s *BasevbaListener) EnterVsAnd(ctx *VsAndContext) {}

// ExitVsAnd is called when production vsAnd is exited.
func (s *BasevbaListener) ExitVsAnd(ctx *VsAndContext) {}

// EnterVsLeq is called when production vsLeq is entered.
func (s *BasevbaListener) EnterVsLeq(ctx *VsLeqContext) {}

// ExitVsLeq is called when production vsLeq is exited.
func (s *BasevbaListener) ExitVsLeq(ctx *VsLeqContext) {}

// EnterVsPow is called when production vsPow is entered.
func (s *BasevbaListener) EnterVsPow(ctx *VsPowContext) {}

// ExitVsPow is called when production vsPow is exited.
func (s *BasevbaListener) ExitVsPow(ctx *VsPowContext) {}

// EnterVsIs is called when production vsIs is entered.
func (s *BasevbaListener) EnterVsIs(ctx *VsIsContext) {}

// ExitVsIs is called when production vsIs is exited.
func (s *BasevbaListener) ExitVsIs(ctx *VsIsContext) {}

// EnterVsMod is called when production vsMod is entered.
func (s *BasevbaListener) EnterVsMod(ctx *VsModContext) {}

// ExitVsMod is called when production vsMod is exited.
func (s *BasevbaListener) ExitVsMod(ctx *VsModContext) {}

// EnterVsAmp is called when production vsAmp is entered.
func (s *BasevbaListener) EnterVsAmp(ctx *VsAmpContext) {}

// ExitVsAmp is called when production vsAmp is exited.
func (s *BasevbaListener) ExitVsAmp(ctx *VsAmpContext) {}

// EnterVsOr is called when production vsOr is entered.
func (s *BasevbaListener) EnterVsOr(ctx *VsOrContext) {}

// ExitVsOr is called when production vsOr is exited.
func (s *BasevbaListener) ExitVsOr(ctx *VsOrContext) {}

// EnterVsMinus is called when production vsMinus is entered.
func (s *BasevbaListener) EnterVsMinus(ctx *VsMinusContext) {}

// ExitVsMinus is called when production vsMinus is exited.
func (s *BasevbaListener) ExitVsMinus(ctx *VsMinusContext) {}

// EnterVsLiteral is called when production vsLiteral is entered.
func (s *BasevbaListener) EnterVsLiteral(ctx *VsLiteralContext) {}

// ExitVsLiteral is called when production vsLiteral is exited.
func (s *BasevbaListener) ExitVsLiteral(ctx *VsLiteralContext) {}

// EnterVsEqv is called when production vsEqv is entered.
func (s *BasevbaListener) EnterVsEqv(ctx *VsEqvContext) {}

// ExitVsEqv is called when production vsEqv is exited.
func (s *BasevbaListener) ExitVsEqv(ctx *VsEqvContext) {}

// EnterVsImp is called when production vsImp is entered.
func (s *BasevbaListener) EnterVsImp(ctx *VsImpContext) {}

// ExitVsImp is called when production vsImp is exited.
func (s *BasevbaListener) ExitVsImp(ctx *VsImpContext) {}

// EnterVsGt is called when production vsGt is entered.
func (s *BasevbaListener) EnterVsGt(ctx *VsGtContext) {}

// ExitVsGt is called when production vsGt is exited.
func (s *BasevbaListener) ExitVsGt(ctx *VsGtContext) {}

// EnterVsEq is called when production vsEq is entered.
func (s *BasevbaListener) EnterVsEq(ctx *VsEqContext) {}

// ExitVsEq is called when production vsEq is exited.
func (s *BasevbaListener) ExitVsEq(ctx *VsEqContext) {}

// EnterVsMid is called when production vsMid is entered.
func (s *BasevbaListener) EnterVsMid(ctx *VsMidContext) {}

// ExitVsMid is called when production vsMid is exited.
func (s *BasevbaListener) ExitVsMid(ctx *VsMidContext) {}

// EnterVariableStmt is called when production variableStmt is entered.
func (s *BasevbaListener) EnterVariableStmt(ctx *VariableStmtContext) {}

// ExitVariableStmt is called when production variableStmt is exited.
func (s *BasevbaListener) ExitVariableStmt(ctx *VariableStmtContext) {}

// EnterVariableListStmt is called when production variableListStmt is entered.
func (s *BasevbaListener) EnterVariableListStmt(ctx *VariableListStmtContext) {}

// ExitVariableListStmt is called when production variableListStmt is exited.
func (s *BasevbaListener) ExitVariableListStmt(ctx *VariableListStmtContext) {}

// EnterVariableSubStmt is called when production variableSubStmt is entered.
func (s *BasevbaListener) EnterVariableSubStmt(ctx *VariableSubStmtContext) {}

// ExitVariableSubStmt is called when production variableSubStmt is exited.
func (s *BasevbaListener) ExitVariableSubStmt(ctx *VariableSubStmtContext) {}

// EnterWhileWendStmt is called when production whileWendStmt is entered.
func (s *BasevbaListener) EnterWhileWendStmt(ctx *WhileWendStmtContext) {}

// ExitWhileWendStmt is called when production whileWendStmt is exited.
func (s *BasevbaListener) ExitWhileWendStmt(ctx *WhileWendStmtContext) {}

// EnterWidthStmt is called when production widthStmt is entered.
func (s *BasevbaListener) EnterWidthStmt(ctx *WidthStmtContext) {}

// ExitWidthStmt is called when production widthStmt is exited.
func (s *BasevbaListener) ExitWidthStmt(ctx *WidthStmtContext) {}

// EnterWithStmt is called when production withStmt is entered.
func (s *BasevbaListener) EnterWithStmt(ctx *WithStmtContext) {}

// ExitWithStmt is called when production withStmt is exited.
func (s *BasevbaListener) ExitWithStmt(ctx *WithStmtContext) {}

// EnterWriteStmt is called when production writeStmt is entered.
func (s *BasevbaListener) EnterWriteStmt(ctx *WriteStmtContext) {}

// ExitWriteStmt is called when production writeStmt is exited.
func (s *BasevbaListener) ExitWriteStmt(ctx *WriteStmtContext) {}

// EnterFileNumber is called when production fileNumber is entered.
func (s *BasevbaListener) EnterFileNumber(ctx *FileNumberContext) {}

// ExitFileNumber is called when production fileNumber is exited.
func (s *BasevbaListener) ExitFileNumber(ctx *FileNumberContext) {}

// EnterExplicitCallStmt is called when production explicitCallStmt is entered.
func (s *BasevbaListener) EnterExplicitCallStmt(ctx *ExplicitCallStmtContext) {}

// ExitExplicitCallStmt is called when production explicitCallStmt is exited.
func (s *BasevbaListener) ExitExplicitCallStmt(ctx *ExplicitCallStmtContext) {}

// EnterECS_ProcedureCall is called when production eCS_ProcedureCall is entered.
func (s *BasevbaListener) EnterECS_ProcedureCall(ctx *ECS_ProcedureCallContext) {}

// ExitECS_ProcedureCall is called when production eCS_ProcedureCall is exited.
func (s *BasevbaListener) ExitECS_ProcedureCall(ctx *ECS_ProcedureCallContext) {}

// EnterECS_MemberProcedureCall is called when production eCS_MemberProcedureCall is entered.
func (s *BasevbaListener) EnterECS_MemberProcedureCall(ctx *ECS_MemberProcedureCallContext) {}

// ExitECS_MemberProcedureCall is called when production eCS_MemberProcedureCall is exited.
func (s *BasevbaListener) ExitECS_MemberProcedureCall(ctx *ECS_MemberProcedureCallContext) {}

// EnterImplicitCallStmt_InBlock is called when production implicitCallStmt_InBlock is entered.
func (s *BasevbaListener) EnterImplicitCallStmt_InBlock(ctx *ImplicitCallStmt_InBlockContext) {}

// ExitImplicitCallStmt_InBlock is called when production implicitCallStmt_InBlock is exited.
func (s *BasevbaListener) ExitImplicitCallStmt_InBlock(ctx *ImplicitCallStmt_InBlockContext) {}

// EnterICS_B_MemberProcedureCall is called when production iCS_B_MemberProcedureCall is entered.
func (s *BasevbaListener) EnterICS_B_MemberProcedureCall(ctx *ICS_B_MemberProcedureCallContext) {}

// ExitICS_B_MemberProcedureCall is called when production iCS_B_MemberProcedureCall is exited.
func (s *BasevbaListener) ExitICS_B_MemberProcedureCall(ctx *ICS_B_MemberProcedureCallContext) {}

// EnterICS_B_ProcedureCall is called when production iCS_B_ProcedureCall is entered.
func (s *BasevbaListener) EnterICS_B_ProcedureCall(ctx *ICS_B_ProcedureCallContext) {}

// ExitICS_B_ProcedureCall is called when production iCS_B_ProcedureCall is exited.
func (s *BasevbaListener) ExitICS_B_ProcedureCall(ctx *ICS_B_ProcedureCallContext) {}

// EnterImplicitCallStmt_InStmt is called when production implicitCallStmt_InStmt is entered.
func (s *BasevbaListener) EnterImplicitCallStmt_InStmt(ctx *ImplicitCallStmt_InStmtContext) {}

// ExitImplicitCallStmt_InStmt is called when production implicitCallStmt_InStmt is exited.
func (s *BasevbaListener) ExitImplicitCallStmt_InStmt(ctx *ImplicitCallStmt_InStmtContext) {}

// EnterICS_S_VariableOrProcedureCall is called when production iCS_S_VariableOrProcedureCall is entered.
func (s *BasevbaListener) EnterICS_S_VariableOrProcedureCall(ctx *ICS_S_VariableOrProcedureCallContext) {
}

// ExitICS_S_VariableOrProcedureCall is called when production iCS_S_VariableOrProcedureCall is exited.
func (s *BasevbaListener) ExitICS_S_VariableOrProcedureCall(ctx *ICS_S_VariableOrProcedureCallContext) {
}

// EnterICS_S_ProcedureOrArrayCall is called when production iCS_S_ProcedureOrArrayCall is entered.
func (s *BasevbaListener) EnterICS_S_ProcedureOrArrayCall(ctx *ICS_S_ProcedureOrArrayCallContext) {}

// ExitICS_S_ProcedureOrArrayCall is called when production iCS_S_ProcedureOrArrayCall is exited.
func (s *BasevbaListener) ExitICS_S_ProcedureOrArrayCall(ctx *ICS_S_ProcedureOrArrayCallContext) {}

// EnterICS_S_MembersCall is called when production iCS_S_MembersCall is entered.
func (s *BasevbaListener) EnterICS_S_MembersCall(ctx *ICS_S_MembersCallContext) {}

// ExitICS_S_MembersCall is called when production iCS_S_MembersCall is exited.
func (s *BasevbaListener) ExitICS_S_MembersCall(ctx *ICS_S_MembersCallContext) {}

// EnterICS_S_MemberCall is called when production iCS_S_MemberCall is entered.
func (s *BasevbaListener) EnterICS_S_MemberCall(ctx *ICS_S_MemberCallContext) {}

// ExitICS_S_MemberCall is called when production iCS_S_MemberCall is exited.
func (s *BasevbaListener) ExitICS_S_MemberCall(ctx *ICS_S_MemberCallContext) {}

// EnterICS_S_DictionaryCall is called when production iCS_S_DictionaryCall is entered.
func (s *BasevbaListener) EnterICS_S_DictionaryCall(ctx *ICS_S_DictionaryCallContext) {}

// ExitICS_S_DictionaryCall is called when production iCS_S_DictionaryCall is exited.
func (s *BasevbaListener) ExitICS_S_DictionaryCall(ctx *ICS_S_DictionaryCallContext) {}

// EnterArgsCall is called when production argsCall is entered.
func (s *BasevbaListener) EnterArgsCall(ctx *ArgsCallContext) {}

// ExitArgsCall is called when production argsCall is exited.
func (s *BasevbaListener) ExitArgsCall(ctx *ArgsCallContext) {}

// EnterArgCall is called when production argCall is entered.
func (s *BasevbaListener) EnterArgCall(ctx *ArgCallContext) {}

// ExitArgCall is called when production argCall is exited.
func (s *BasevbaListener) ExitArgCall(ctx *ArgCallContext) {}

// EnterDictionaryCallStmt is called when production dictionaryCallStmt is entered.
func (s *BasevbaListener) EnterDictionaryCallStmt(ctx *DictionaryCallStmtContext) {}

// ExitDictionaryCallStmt is called when production dictionaryCallStmt is exited.
func (s *BasevbaListener) ExitDictionaryCallStmt(ctx *DictionaryCallStmtContext) {}

// EnterArgList is called when production argList is entered.
func (s *BasevbaListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BasevbaListener) ExitArgList(ctx *ArgListContext) {}

// EnterArg is called when production arg is entered.
func (s *BasevbaListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BasevbaListener) ExitArg(ctx *ArgContext) {}

// EnterArgDefaultValue is called when production argDefaultValue is entered.
func (s *BasevbaListener) EnterArgDefaultValue(ctx *ArgDefaultValueContext) {}

// ExitArgDefaultValue is called when production argDefaultValue is exited.
func (s *BasevbaListener) ExitArgDefaultValue(ctx *ArgDefaultValueContext) {}

// EnterSubscripts is called when production subscripts is entered.
func (s *BasevbaListener) EnterSubscripts(ctx *SubscriptsContext) {}

// ExitSubscripts is called when production subscripts is exited.
func (s *BasevbaListener) ExitSubscripts(ctx *SubscriptsContext) {}

// EnterSubscript_ is called when production subscript_ is entered.
func (s *BasevbaListener) EnterSubscript_(ctx *Subscript_Context) {}

// ExitSubscript_ is called when production subscript_ is exited.
func (s *BasevbaListener) ExitSubscript_(ctx *Subscript_Context) {}

// EnterAmbiguousIdentifier is called when production ambiguousIdentifier is entered.
func (s *BasevbaListener) EnterAmbiguousIdentifier(ctx *AmbiguousIdentifierContext) {}

// ExitAmbiguousIdentifier is called when production ambiguousIdentifier is exited.
func (s *BasevbaListener) ExitAmbiguousIdentifier(ctx *AmbiguousIdentifierContext) {}

// EnterAsTypeClause is called when production asTypeClause is entered.
func (s *BasevbaListener) EnterAsTypeClause(ctx *AsTypeClauseContext) {}

// ExitAsTypeClause is called when production asTypeClause is exited.
func (s *BasevbaListener) ExitAsTypeClause(ctx *AsTypeClauseContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BasevbaListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BasevbaListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterCertainIdentifier is called when production certainIdentifier is entered.
func (s *BasevbaListener) EnterCertainIdentifier(ctx *CertainIdentifierContext) {}

// ExitCertainIdentifier is called when production certainIdentifier is exited.
func (s *BasevbaListener) ExitCertainIdentifier(ctx *CertainIdentifierContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BasevbaListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BasevbaListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterComplexType is called when production complexType is entered.
func (s *BasevbaListener) EnterComplexType(ctx *ComplexTypeContext) {}

// ExitComplexType is called when production complexType is exited.
func (s *BasevbaListener) ExitComplexType(ctx *ComplexTypeContext) {}

// EnterFieldLength is called when production fieldLength is entered.
func (s *BasevbaListener) EnterFieldLength(ctx *FieldLengthContext) {}

// ExitFieldLength is called when production fieldLength is exited.
func (s *BasevbaListener) ExitFieldLength(ctx *FieldLengthContext) {}

// EnterLetterrange is called when production letterrange is entered.
func (s *BasevbaListener) EnterLetterrange(ctx *LetterrangeContext) {}

// ExitLetterrange is called when production letterrange is exited.
func (s *BasevbaListener) ExitLetterrange(ctx *LetterrangeContext) {}

// EnterLineLabel is called when production lineLabel is entered.
func (s *BasevbaListener) EnterLineLabel(ctx *LineLabelContext) {}

// ExitLineLabel is called when production lineLabel is exited.
func (s *BasevbaListener) ExitLineLabel(ctx *LineLabelContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasevbaListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasevbaListener) ExitLiteral(ctx *LiteralContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BasevbaListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BasevbaListener) ExitType_(ctx *Type_Context) {}

// EnterTypeHint is called when production typeHint is entered.
func (s *BasevbaListener) EnterTypeHint(ctx *TypeHintContext) {}

// ExitTypeHint is called when production typeHint is exited.
func (s *BasevbaListener) ExitTypeHint(ctx *TypeHintContext) {}

// EnterVisibility is called when production visibility is entered.
func (s *BasevbaListener) EnterVisibility(ctx *VisibilityContext) {}

// ExitVisibility is called when production visibility is exited.
func (s *BasevbaListener) ExitVisibility(ctx *VisibilityContext) {}

// EnterAmbiguousKeyword is called when production ambiguousKeyword is entered.
func (s *BasevbaListener) EnterAmbiguousKeyword(ctx *AmbiguousKeywordContext) {}

// ExitAmbiguousKeyword is called when production ambiguousKeyword is exited.
func (s *BasevbaListener) ExitAmbiguousKeyword(ctx *AmbiguousKeywordContext) {}

// EnterRemComment is called when production remComment is entered.
func (s *BasevbaListener) EnterRemComment(ctx *RemCommentContext) {}

// ExitRemComment is called when production remComment is exited.
func (s *BasevbaListener) ExitRemComment(ctx *RemCommentContext) {}

// EnterComment is called when production comment is entered.
func (s *BasevbaListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BasevbaListener) ExitComment(ctx *CommentContext) {}

// EnterEndOfLine is called when production endOfLine is entered.
func (s *BasevbaListener) EnterEndOfLine(ctx *EndOfLineContext) {}

// ExitEndOfLine is called when production endOfLine is exited.
func (s *BasevbaListener) ExitEndOfLine(ctx *EndOfLineContext) {}

// EnterEndOfStatement is called when production endOfStatement is entered.
func (s *BasevbaListener) EnterEndOfStatement(ctx *EndOfStatementContext) {}

// ExitEndOfStatement is called when production endOfStatement is exited.
func (s *BasevbaListener) ExitEndOfStatement(ctx *EndOfStatementContext) {}
