// Code generated from VisualBasic6Parser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package visualbasic6 // VisualBasic6Parser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVisualBasic6ParserListener is a complete listener for a parse tree produced by VisualBasic6Parser.
type BaseVisualBasic6ParserListener struct{}

var _ VisualBasic6ParserListener = &BaseVisualBasic6ParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVisualBasic6ParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVisualBasic6ParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVisualBasic6ParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVisualBasic6ParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStartRule is called when production startRule is entered.
func (s *BaseVisualBasic6ParserListener) EnterStartRule(ctx *StartRuleContext) {}

// ExitStartRule is called when production startRule is exited.
func (s *BaseVisualBasic6ParserListener) ExitStartRule(ctx *StartRuleContext) {}

// EnterModule is called when production module is entered.
func (s *BaseVisualBasic6ParserListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseVisualBasic6ParserListener) ExitModule(ctx *ModuleContext) {}

// EnterModuleReferences is called when production moduleReferences is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleReferences(ctx *ModuleReferencesContext) {}

// ExitModuleReferences is called when production moduleReferences is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleReferences(ctx *ModuleReferencesContext) {}

// EnterModuleReference is called when production moduleReference is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleReference(ctx *ModuleReferenceContext) {}

// ExitModuleReference is called when production moduleReference is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleReference(ctx *ModuleReferenceContext) {}

// EnterModuleReferenceValue is called when production moduleReferenceValue is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleReferenceValue(ctx *ModuleReferenceValueContext) {
}

// ExitModuleReferenceValue is called when production moduleReferenceValue is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleReferenceValue(ctx *ModuleReferenceValueContext) {}

// EnterModuleReferenceComponent is called when production moduleReferenceComponent is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleReferenceComponent(ctx *ModuleReferenceComponentContext) {
}

// ExitModuleReferenceComponent is called when production moduleReferenceComponent is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleReferenceComponent(ctx *ModuleReferenceComponentContext) {
}

// EnterModuleHeader is called when production moduleHeader is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleHeader(ctx *ModuleHeaderContext) {}

// ExitModuleHeader is called when production moduleHeader is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleHeader(ctx *ModuleHeaderContext) {}

// EnterModuleConfig is called when production moduleConfig is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleConfig(ctx *ModuleConfigContext) {}

// ExitModuleConfig is called when production moduleConfig is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleConfig(ctx *ModuleConfigContext) {}

// EnterModuleConfigElement is called when production moduleConfigElement is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleConfigElement(ctx *ModuleConfigElementContext) {}

// ExitModuleConfigElement is called when production moduleConfigElement is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleConfigElement(ctx *ModuleConfigElementContext) {}

// EnterModuleAttributes is called when production moduleAttributes is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleAttributes(ctx *ModuleAttributesContext) {}

// ExitModuleAttributes is called when production moduleAttributes is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleAttributes(ctx *ModuleAttributesContext) {}

// EnterModuleOptions is called when production moduleOptions is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleOptions(ctx *ModuleOptionsContext) {}

// ExitModuleOptions is called when production moduleOptions is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleOptions(ctx *ModuleOptionsContext) {}

// EnterOptionBaseStmt is called when production optionBaseStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterOptionBaseStmt(ctx *OptionBaseStmtContext) {}

// ExitOptionBaseStmt is called when production optionBaseStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitOptionBaseStmt(ctx *OptionBaseStmtContext) {}

// EnterOptionCompareStmt is called when production optionCompareStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterOptionCompareStmt(ctx *OptionCompareStmtContext) {}

// ExitOptionCompareStmt is called when production optionCompareStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitOptionCompareStmt(ctx *OptionCompareStmtContext) {}

// EnterOptionExplicitStmt is called when production optionExplicitStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterOptionExplicitStmt(ctx *OptionExplicitStmtContext) {}

// ExitOptionExplicitStmt is called when production optionExplicitStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitOptionExplicitStmt(ctx *OptionExplicitStmtContext) {}

// EnterOptionPrivateModuleStmt is called when production optionPrivateModuleStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterOptionPrivateModuleStmt(ctx *OptionPrivateModuleStmtContext) {
}

// ExitOptionPrivateModuleStmt is called when production optionPrivateModuleStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitOptionPrivateModuleStmt(ctx *OptionPrivateModuleStmtContext) {
}

// EnterModuleBody is called when production moduleBody is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleBody(ctx *ModuleBodyContext) {}

// ExitModuleBody is called when production moduleBody is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleBody(ctx *ModuleBodyContext) {}

// EnterModuleBodyElement is called when production moduleBodyElement is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleBodyElement(ctx *ModuleBodyElementContext) {}

// ExitModuleBodyElement is called when production moduleBodyElement is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleBodyElement(ctx *ModuleBodyElementContext) {}

// EnterControlProperties is called when production controlProperties is entered.
func (s *BaseVisualBasic6ParserListener) EnterControlProperties(ctx *ControlPropertiesContext) {}

// ExitControlProperties is called when production controlProperties is exited.
func (s *BaseVisualBasic6ParserListener) ExitControlProperties(ctx *ControlPropertiesContext) {}

// EnterCp_Properties is called when production cp_Properties is entered.
func (s *BaseVisualBasic6ParserListener) EnterCp_Properties(ctx *Cp_PropertiesContext) {}

// ExitCp_Properties is called when production cp_Properties is exited.
func (s *BaseVisualBasic6ParserListener) ExitCp_Properties(ctx *Cp_PropertiesContext) {}

// EnterCp_SingleProperty is called when production cp_SingleProperty is entered.
func (s *BaseVisualBasic6ParserListener) EnterCp_SingleProperty(ctx *Cp_SinglePropertyContext) {}

// ExitCp_SingleProperty is called when production cp_SingleProperty is exited.
func (s *BaseVisualBasic6ParserListener) ExitCp_SingleProperty(ctx *Cp_SinglePropertyContext) {}

// EnterCp_PropertyName is called when production cp_PropertyName is entered.
func (s *BaseVisualBasic6ParserListener) EnterCp_PropertyName(ctx *Cp_PropertyNameContext) {}

// ExitCp_PropertyName is called when production cp_PropertyName is exited.
func (s *BaseVisualBasic6ParserListener) ExitCp_PropertyName(ctx *Cp_PropertyNameContext) {}

// EnterCp_PropertyValue is called when production cp_PropertyValue is entered.
func (s *BaseVisualBasic6ParserListener) EnterCp_PropertyValue(ctx *Cp_PropertyValueContext) {}

// ExitCp_PropertyValue is called when production cp_PropertyValue is exited.
func (s *BaseVisualBasic6ParserListener) ExitCp_PropertyValue(ctx *Cp_PropertyValueContext) {}

// EnterCp_NestedProperty is called when production cp_NestedProperty is entered.
func (s *BaseVisualBasic6ParserListener) EnterCp_NestedProperty(ctx *Cp_NestedPropertyContext) {}

// ExitCp_NestedProperty is called when production cp_NestedProperty is exited.
func (s *BaseVisualBasic6ParserListener) ExitCp_NestedProperty(ctx *Cp_NestedPropertyContext) {}

// EnterCp_ControlType is called when production cp_ControlType is entered.
func (s *BaseVisualBasic6ParserListener) EnterCp_ControlType(ctx *Cp_ControlTypeContext) {}

// ExitCp_ControlType is called when production cp_ControlType is exited.
func (s *BaseVisualBasic6ParserListener) ExitCp_ControlType(ctx *Cp_ControlTypeContext) {}

// EnterCp_ControlIdentifier is called when production cp_ControlIdentifier is entered.
func (s *BaseVisualBasic6ParserListener) EnterCp_ControlIdentifier(ctx *Cp_ControlIdentifierContext) {
}

// ExitCp_ControlIdentifier is called when production cp_ControlIdentifier is exited.
func (s *BaseVisualBasic6ParserListener) ExitCp_ControlIdentifier(ctx *Cp_ControlIdentifierContext) {}

// EnterModuleBlock is called when production moduleBlock is entered.
func (s *BaseVisualBasic6ParserListener) EnterModuleBlock(ctx *ModuleBlockContext) {}

// ExitModuleBlock is called when production moduleBlock is exited.
func (s *BaseVisualBasic6ParserListener) ExitModuleBlock(ctx *ModuleBlockContext) {}

// EnterAttributeStmt is called when production attributeStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterAttributeStmt(ctx *AttributeStmtContext) {}

// ExitAttributeStmt is called when production attributeStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitAttributeStmt(ctx *AttributeStmtContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseVisualBasic6ParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseVisualBasic6ParserListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStmt is called when production blockStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterBlockStmt(ctx *BlockStmtContext) {}

// ExitBlockStmt is called when production blockStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitBlockStmt(ctx *BlockStmtContext) {}

// EnterAppActivateStmt is called when production appActivateStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterAppActivateStmt(ctx *AppActivateStmtContext) {}

// ExitAppActivateStmt is called when production appActivateStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitAppActivateStmt(ctx *AppActivateStmtContext) {}

// EnterBeepStmt is called when production beepStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterBeepStmt(ctx *BeepStmtContext) {}

// ExitBeepStmt is called when production beepStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitBeepStmt(ctx *BeepStmtContext) {}

// EnterChDirStmt is called when production chDirStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterChDirStmt(ctx *ChDirStmtContext) {}

// ExitChDirStmt is called when production chDirStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitChDirStmt(ctx *ChDirStmtContext) {}

// EnterChDriveStmt is called when production chDriveStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterChDriveStmt(ctx *ChDriveStmtContext) {}

// ExitChDriveStmt is called when production chDriveStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitChDriveStmt(ctx *ChDriveStmtContext) {}

// EnterCloseStmt is called when production closeStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterCloseStmt(ctx *CloseStmtContext) {}

// ExitCloseStmt is called when production closeStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitCloseStmt(ctx *CloseStmtContext) {}

// EnterConstStmt is called when production constStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterConstStmt(ctx *ConstStmtContext) {}

// ExitConstStmt is called when production constStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitConstStmt(ctx *ConstStmtContext) {}

// EnterConstSubStmt is called when production constSubStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterConstSubStmt(ctx *ConstSubStmtContext) {}

// ExitConstSubStmt is called when production constSubStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitConstSubStmt(ctx *ConstSubStmtContext) {}

// EnterDateStmt is called when production dateStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterDateStmt(ctx *DateStmtContext) {}

// ExitDateStmt is called when production dateStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitDateStmt(ctx *DateStmtContext) {}

// EnterDeclareStmt is called when production declareStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterDeclareStmt(ctx *DeclareStmtContext) {}

// ExitDeclareStmt is called when production declareStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitDeclareStmt(ctx *DeclareStmtContext) {}

// EnterDeftypeStmt is called when production deftypeStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterDeftypeStmt(ctx *DeftypeStmtContext) {}

// ExitDeftypeStmt is called when production deftypeStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitDeftypeStmt(ctx *DeftypeStmtContext) {}

// EnterDeleteSettingStmt is called when production deleteSettingStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterDeleteSettingStmt(ctx *DeleteSettingStmtContext) {}

// ExitDeleteSettingStmt is called when production deleteSettingStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitDeleteSettingStmt(ctx *DeleteSettingStmtContext) {}

// EnterDoLoopStmt is called when production doLoopStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterDoLoopStmt(ctx *DoLoopStmtContext) {}

// ExitDoLoopStmt is called when production doLoopStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitDoLoopStmt(ctx *DoLoopStmtContext) {}

// EnterEndStmt is called when production endStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterEndStmt(ctx *EndStmtContext) {}

// ExitEndStmt is called when production endStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitEndStmt(ctx *EndStmtContext) {}

// EnterEnumerationStmt is called when production enumerationStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterEnumerationStmt(ctx *EnumerationStmtContext) {}

// ExitEnumerationStmt is called when production enumerationStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitEnumerationStmt(ctx *EnumerationStmtContext) {}

// EnterEnumerationStmt_Constant is called when production enumerationStmt_Constant is entered.
func (s *BaseVisualBasic6ParserListener) EnterEnumerationStmt_Constant(ctx *EnumerationStmt_ConstantContext) {
}

// ExitEnumerationStmt_Constant is called when production enumerationStmt_Constant is exited.
func (s *BaseVisualBasic6ParserListener) ExitEnumerationStmt_Constant(ctx *EnumerationStmt_ConstantContext) {
}

// EnterEraseStmt is called when production eraseStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterEraseStmt(ctx *EraseStmtContext) {}

// ExitEraseStmt is called when production eraseStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitEraseStmt(ctx *EraseStmtContext) {}

// EnterErrorStmt is called when production errorStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterErrorStmt(ctx *ErrorStmtContext) {}

// ExitErrorStmt is called when production errorStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitErrorStmt(ctx *ErrorStmtContext) {}

// EnterEventStmt is called when production eventStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterEventStmt(ctx *EventStmtContext) {}

// ExitEventStmt is called when production eventStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitEventStmt(ctx *EventStmtContext) {}

// EnterExitStmt is called when production exitStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterExitStmt(ctx *ExitStmtContext) {}

// ExitExitStmt is called when production exitStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitExitStmt(ctx *ExitStmtContext) {}

// EnterFilecopyStmt is called when production filecopyStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterFilecopyStmt(ctx *FilecopyStmtContext) {}

// ExitFilecopyStmt is called when production filecopyStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitFilecopyStmt(ctx *FilecopyStmtContext) {}

// EnterForEachStmt is called when production forEachStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterForEachStmt(ctx *ForEachStmtContext) {}

// ExitForEachStmt is called when production forEachStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitForEachStmt(ctx *ForEachStmtContext) {}

// EnterForNextStmt is called when production forNextStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterForNextStmt(ctx *ForNextStmtContext) {}

// ExitForNextStmt is called when production forNextStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitForNextStmt(ctx *ForNextStmtContext) {}

// EnterFunctionStmt is called when production functionStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterFunctionStmt(ctx *FunctionStmtContext) {}

// ExitFunctionStmt is called when production functionStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitFunctionStmt(ctx *FunctionStmtContext) {}

// EnterGetStmt is called when production getStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterGetStmt(ctx *GetStmtContext) {}

// ExitGetStmt is called when production getStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitGetStmt(ctx *GetStmtContext) {}

// EnterGoSubStmt is called when production goSubStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterGoSubStmt(ctx *GoSubStmtContext) {}

// ExitGoSubStmt is called when production goSubStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitGoSubStmt(ctx *GoSubStmtContext) {}

// EnterGoToStmt is called when production goToStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterGoToStmt(ctx *GoToStmtContext) {}

// ExitGoToStmt is called when production goToStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitGoToStmt(ctx *GoToStmtContext) {}

// EnterInlineIfThenElse is called when production inlineIfThenElse is entered.
func (s *BaseVisualBasic6ParserListener) EnterInlineIfThenElse(ctx *InlineIfThenElseContext) {}

// ExitInlineIfThenElse is called when production inlineIfThenElse is exited.
func (s *BaseVisualBasic6ParserListener) ExitInlineIfThenElse(ctx *InlineIfThenElseContext) {}

// EnterBlockIfThenElse is called when production blockIfThenElse is entered.
func (s *BaseVisualBasic6ParserListener) EnterBlockIfThenElse(ctx *BlockIfThenElseContext) {}

// ExitBlockIfThenElse is called when production blockIfThenElse is exited.
func (s *BaseVisualBasic6ParserListener) ExitBlockIfThenElse(ctx *BlockIfThenElseContext) {}

// EnterIfBlockStmt is called when production ifBlockStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterIfBlockStmt(ctx *IfBlockStmtContext) {}

// ExitIfBlockStmt is called when production ifBlockStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitIfBlockStmt(ctx *IfBlockStmtContext) {}

// EnterIfConditionStmt is called when production ifConditionStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterIfConditionStmt(ctx *IfConditionStmtContext) {}

// ExitIfConditionStmt is called when production ifConditionStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitIfConditionStmt(ctx *IfConditionStmtContext) {}

// EnterIfElseIfBlockStmt is called when production ifElseIfBlockStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterIfElseIfBlockStmt(ctx *IfElseIfBlockStmtContext) {}

// ExitIfElseIfBlockStmt is called when production ifElseIfBlockStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitIfElseIfBlockStmt(ctx *IfElseIfBlockStmtContext) {}

// EnterIfElseBlockStmt is called when production ifElseBlockStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterIfElseBlockStmt(ctx *IfElseBlockStmtContext) {}

// ExitIfElseBlockStmt is called when production ifElseBlockStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitIfElseBlockStmt(ctx *IfElseBlockStmtContext) {}

// EnterImplementsStmt is called when production implementsStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterImplementsStmt(ctx *ImplementsStmtContext) {}

// ExitImplementsStmt is called when production implementsStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitImplementsStmt(ctx *ImplementsStmtContext) {}

// EnterInputStmt is called when production inputStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterInputStmt(ctx *InputStmtContext) {}

// ExitInputStmt is called when production inputStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitInputStmt(ctx *InputStmtContext) {}

// EnterKillStmt is called when production killStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterKillStmt(ctx *KillStmtContext) {}

// ExitKillStmt is called when production killStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitKillStmt(ctx *KillStmtContext) {}

// EnterLetStmt is called when production letStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterLetStmt(ctx *LetStmtContext) {}

// ExitLetStmt is called when production letStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitLetStmt(ctx *LetStmtContext) {}

// EnterLineInputStmt is called when production lineInputStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterLineInputStmt(ctx *LineInputStmtContext) {}

// ExitLineInputStmt is called when production lineInputStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitLineInputStmt(ctx *LineInputStmtContext) {}

// EnterLoadStmt is called when production loadStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterLoadStmt(ctx *LoadStmtContext) {}

// ExitLoadStmt is called when production loadStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitLoadStmt(ctx *LoadStmtContext) {}

// EnterLockStmt is called when production lockStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterLockStmt(ctx *LockStmtContext) {}

// ExitLockStmt is called when production lockStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitLockStmt(ctx *LockStmtContext) {}

// EnterLsetStmt is called when production lsetStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterLsetStmt(ctx *LsetStmtContext) {}

// ExitLsetStmt is called when production lsetStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitLsetStmt(ctx *LsetStmtContext) {}

// EnterMacroIfThenElseStmt is called when production macroIfThenElseStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterMacroIfThenElseStmt(ctx *MacroIfThenElseStmtContext) {}

// ExitMacroIfThenElseStmt is called when production macroIfThenElseStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitMacroIfThenElseStmt(ctx *MacroIfThenElseStmtContext) {}

// EnterMacroIfBlockStmt is called when production macroIfBlockStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterMacroIfBlockStmt(ctx *MacroIfBlockStmtContext) {}

// ExitMacroIfBlockStmt is called when production macroIfBlockStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitMacroIfBlockStmt(ctx *MacroIfBlockStmtContext) {}

// EnterMacroElseIfBlockStmt is called when production macroElseIfBlockStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterMacroElseIfBlockStmt(ctx *MacroElseIfBlockStmtContext) {
}

// ExitMacroElseIfBlockStmt is called when production macroElseIfBlockStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitMacroElseIfBlockStmt(ctx *MacroElseIfBlockStmtContext) {}

// EnterMacroElseBlockStmt is called when production macroElseBlockStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterMacroElseBlockStmt(ctx *MacroElseBlockStmtContext) {}

// ExitMacroElseBlockStmt is called when production macroElseBlockStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitMacroElseBlockStmt(ctx *MacroElseBlockStmtContext) {}

// EnterMidStmt is called when production midStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterMidStmt(ctx *MidStmtContext) {}

// ExitMidStmt is called when production midStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitMidStmt(ctx *MidStmtContext) {}

// EnterMkdirStmt is called when production mkdirStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterMkdirStmt(ctx *MkdirStmtContext) {}

// ExitMkdirStmt is called when production mkdirStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitMkdirStmt(ctx *MkdirStmtContext) {}

// EnterNameStmt is called when production nameStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterNameStmt(ctx *NameStmtContext) {}

// ExitNameStmt is called when production nameStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitNameStmt(ctx *NameStmtContext) {}

// EnterOnErrorStmt is called when production onErrorStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterOnErrorStmt(ctx *OnErrorStmtContext) {}

// ExitOnErrorStmt is called when production onErrorStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitOnErrorStmt(ctx *OnErrorStmtContext) {}

// EnterOnGoToStmt is called when production onGoToStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterOnGoToStmt(ctx *OnGoToStmtContext) {}

// ExitOnGoToStmt is called when production onGoToStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitOnGoToStmt(ctx *OnGoToStmtContext) {}

// EnterOnGoSubStmt is called when production onGoSubStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterOnGoSubStmt(ctx *OnGoSubStmtContext) {}

// ExitOnGoSubStmt is called when production onGoSubStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitOnGoSubStmt(ctx *OnGoSubStmtContext) {}

// EnterOpenStmt is called when production openStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterOpenStmt(ctx *OpenStmtContext) {}

// ExitOpenStmt is called when production openStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitOpenStmt(ctx *OpenStmtContext) {}

// EnterOutputList is called when production outputList is entered.
func (s *BaseVisualBasic6ParserListener) EnterOutputList(ctx *OutputListContext) {}

// ExitOutputList is called when production outputList is exited.
func (s *BaseVisualBasic6ParserListener) ExitOutputList(ctx *OutputListContext) {}

// EnterOutputList_Expression is called when production outputList_Expression is entered.
func (s *BaseVisualBasic6ParserListener) EnterOutputList_Expression(ctx *OutputList_ExpressionContext) {
}

// ExitOutputList_Expression is called when production outputList_Expression is exited.
func (s *BaseVisualBasic6ParserListener) ExitOutputList_Expression(ctx *OutputList_ExpressionContext) {
}

// EnterPrintStmt is called when production printStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production printStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterPropertyGetStmt is called when production propertyGetStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterPropertyGetStmt(ctx *PropertyGetStmtContext) {}

// ExitPropertyGetStmt is called when production propertyGetStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitPropertyGetStmt(ctx *PropertyGetStmtContext) {}

// EnterPropertySetStmt is called when production propertySetStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterPropertySetStmt(ctx *PropertySetStmtContext) {}

// ExitPropertySetStmt is called when production propertySetStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitPropertySetStmt(ctx *PropertySetStmtContext) {}

// EnterPropertyLetStmt is called when production propertyLetStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterPropertyLetStmt(ctx *PropertyLetStmtContext) {}

// ExitPropertyLetStmt is called when production propertyLetStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitPropertyLetStmt(ctx *PropertyLetStmtContext) {}

// EnterPutStmt is called when production putStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterPutStmt(ctx *PutStmtContext) {}

// ExitPutStmt is called when production putStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitPutStmt(ctx *PutStmtContext) {}

// EnterRaiseEventStmt is called when production raiseEventStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterRaiseEventStmt(ctx *RaiseEventStmtContext) {}

// ExitRaiseEventStmt is called when production raiseEventStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitRaiseEventStmt(ctx *RaiseEventStmtContext) {}

// EnterRandomizeStmt is called when production randomizeStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterRandomizeStmt(ctx *RandomizeStmtContext) {}

// ExitRandomizeStmt is called when production randomizeStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitRandomizeStmt(ctx *RandomizeStmtContext) {}

// EnterRedimStmt is called when production redimStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterRedimStmt(ctx *RedimStmtContext) {}

// ExitRedimStmt is called when production redimStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitRedimStmt(ctx *RedimStmtContext) {}

// EnterRedimSubStmt is called when production redimSubStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterRedimSubStmt(ctx *RedimSubStmtContext) {}

// ExitRedimSubStmt is called when production redimSubStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitRedimSubStmt(ctx *RedimSubStmtContext) {}

// EnterResetStmt is called when production resetStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterResetStmt(ctx *ResetStmtContext) {}

// ExitResetStmt is called when production resetStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitResetStmt(ctx *ResetStmtContext) {}

// EnterResumeStmt is called when production resumeStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterResumeStmt(ctx *ResumeStmtContext) {}

// ExitResumeStmt is called when production resumeStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitResumeStmt(ctx *ResumeStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterRmdirStmt is called when production rmdirStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterRmdirStmt(ctx *RmdirStmtContext) {}

// ExitRmdirStmt is called when production rmdirStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitRmdirStmt(ctx *RmdirStmtContext) {}

// EnterRsetStmt is called when production rsetStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterRsetStmt(ctx *RsetStmtContext) {}

// ExitRsetStmt is called when production rsetStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitRsetStmt(ctx *RsetStmtContext) {}

// EnterSavepictureStmt is called when production savepictureStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterSavepictureStmt(ctx *SavepictureStmtContext) {}

// ExitSavepictureStmt is called when production savepictureStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitSavepictureStmt(ctx *SavepictureStmtContext) {}

// EnterSaveSettingStmt is called when production saveSettingStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterSaveSettingStmt(ctx *SaveSettingStmtContext) {}

// ExitSaveSettingStmt is called when production saveSettingStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitSaveSettingStmt(ctx *SaveSettingStmtContext) {}

// EnterSeekStmt is called when production seekStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterSeekStmt(ctx *SeekStmtContext) {}

// ExitSeekStmt is called when production seekStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitSeekStmt(ctx *SeekStmtContext) {}

// EnterSelectCaseStmt is called when production selectCaseStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterSelectCaseStmt(ctx *SelectCaseStmtContext) {}

// ExitSelectCaseStmt is called when production selectCaseStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitSelectCaseStmt(ctx *SelectCaseStmtContext) {}

// EnterSC_Case is called when production sC_Case is entered.
func (s *BaseVisualBasic6ParserListener) EnterSC_Case(ctx *SC_CaseContext) {}

// ExitSC_Case is called when production sC_Case is exited.
func (s *BaseVisualBasic6ParserListener) ExitSC_Case(ctx *SC_CaseContext) {}

// EnterCaseCondElse is called when production caseCondElse is entered.
func (s *BaseVisualBasic6ParserListener) EnterCaseCondElse(ctx *CaseCondElseContext) {}

// ExitCaseCondElse is called when production caseCondElse is exited.
func (s *BaseVisualBasic6ParserListener) ExitCaseCondElse(ctx *CaseCondElseContext) {}

// EnterCaseCondExpr is called when production caseCondExpr is entered.
func (s *BaseVisualBasic6ParserListener) EnterCaseCondExpr(ctx *CaseCondExprContext) {}

// ExitCaseCondExpr is called when production caseCondExpr is exited.
func (s *BaseVisualBasic6ParserListener) ExitCaseCondExpr(ctx *CaseCondExprContext) {}

// EnterCaseCondExprIs is called when production caseCondExprIs is entered.
func (s *BaseVisualBasic6ParserListener) EnterCaseCondExprIs(ctx *CaseCondExprIsContext) {}

// ExitCaseCondExprIs is called when production caseCondExprIs is exited.
func (s *BaseVisualBasic6ParserListener) ExitCaseCondExprIs(ctx *CaseCondExprIsContext) {}

// EnterCaseCondExprValue is called when production caseCondExprValue is entered.
func (s *BaseVisualBasic6ParserListener) EnterCaseCondExprValue(ctx *CaseCondExprValueContext) {}

// ExitCaseCondExprValue is called when production caseCondExprValue is exited.
func (s *BaseVisualBasic6ParserListener) ExitCaseCondExprValue(ctx *CaseCondExprValueContext) {}

// EnterCaseCondExprTo is called when production caseCondExprTo is entered.
func (s *BaseVisualBasic6ParserListener) EnterCaseCondExprTo(ctx *CaseCondExprToContext) {}

// ExitCaseCondExprTo is called when production caseCondExprTo is exited.
func (s *BaseVisualBasic6ParserListener) ExitCaseCondExprTo(ctx *CaseCondExprToContext) {}

// EnterSendkeysStmt is called when production sendkeysStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterSendkeysStmt(ctx *SendkeysStmtContext) {}

// ExitSendkeysStmt is called when production sendkeysStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitSendkeysStmt(ctx *SendkeysStmtContext) {}

// EnterSetattrStmt is called when production setattrStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterSetattrStmt(ctx *SetattrStmtContext) {}

// ExitSetattrStmt is called when production setattrStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitSetattrStmt(ctx *SetattrStmtContext) {}

// EnterSetStmt is called when production setStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterSetStmt(ctx *SetStmtContext) {}

// ExitSetStmt is called when production setStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitSetStmt(ctx *SetStmtContext) {}

// EnterStopStmt is called when production stopStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterStopStmt(ctx *StopStmtContext) {}

// ExitStopStmt is called when production stopStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitStopStmt(ctx *StopStmtContext) {}

// EnterSubStmt is called when production subStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterSubStmt(ctx *SubStmtContext) {}

// ExitSubStmt is called when production subStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitSubStmt(ctx *SubStmtContext) {}

// EnterTimeStmt is called when production timeStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterTimeStmt(ctx *TimeStmtContext) {}

// ExitTimeStmt is called when production timeStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitTimeStmt(ctx *TimeStmtContext) {}

// EnterTypeStmt is called when production typeStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterTypeStmt(ctx *TypeStmtContext) {}

// ExitTypeStmt is called when production typeStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitTypeStmt(ctx *TypeStmtContext) {}

// EnterTypeStmt_Element is called when production typeStmt_Element is entered.
func (s *BaseVisualBasic6ParserListener) EnterTypeStmt_Element(ctx *TypeStmt_ElementContext) {}

// ExitTypeStmt_Element is called when production typeStmt_Element is exited.
func (s *BaseVisualBasic6ParserListener) ExitTypeStmt_Element(ctx *TypeStmt_ElementContext) {}

// EnterTypeOfStmt is called when production typeOfStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterTypeOfStmt(ctx *TypeOfStmtContext) {}

// ExitTypeOfStmt is called when production typeOfStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitTypeOfStmt(ctx *TypeOfStmtContext) {}

// EnterUnloadStmt is called when production unloadStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterUnloadStmt(ctx *UnloadStmtContext) {}

// ExitUnloadStmt is called when production unloadStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitUnloadStmt(ctx *UnloadStmtContext) {}

// EnterUnlockStmt is called when production unlockStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterUnlockStmt(ctx *UnlockStmtContext) {}

// ExitUnlockStmt is called when production unlockStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitUnlockStmt(ctx *UnlockStmtContext) {}

// EnterVsAssign is called when production vsAssign is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsAssign(ctx *VsAssignContext) {}

// ExitVsAssign is called when production vsAssign is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsAssign(ctx *VsAssignContext) {}

// EnterVsNot is called when production vsNot is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsNot(ctx *VsNotContext) {}

// ExitVsNot is called when production vsNot is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsNot(ctx *VsNotContext) {}

// EnterVsStruct is called when production vsStruct is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsStruct(ctx *VsStructContext) {}

// ExitVsStruct is called when production vsStruct is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsStruct(ctx *VsStructContext) {}

// EnterVsMultDiv is called when production vsMultDiv is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsMultDiv(ctx *VsMultDivContext) {}

// ExitVsMultDiv is called when production vsMultDiv is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsMultDiv(ctx *VsMultDivContext) {}

// EnterVsTypeOf is called when production vsTypeOf is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsTypeOf(ctx *VsTypeOfContext) {}

// ExitVsTypeOf is called when production vsTypeOf is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsTypeOf(ctx *VsTypeOfContext) {}

// EnterVsICS is called when production vsICS is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsICS(ctx *VsICSContext) {}

// ExitVsICS is called when production vsICS is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsICS(ctx *VsICSContext) {}

// EnterVsXor is called when production vsXor is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsXor(ctx *VsXorContext) {}

// ExitVsXor is called when production vsXor is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsXor(ctx *VsXorContext) {}

// EnterVsAddSub is called when production vsAddSub is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsAddSub(ctx *VsAddSubContext) {}

// ExitVsAddSub is called when production vsAddSub is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsAddSub(ctx *VsAddSubContext) {}

// EnterVsAnd is called when production vsAnd is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsAnd(ctx *VsAndContext) {}

// ExitVsAnd is called when production vsAnd is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsAnd(ctx *VsAndContext) {}

// EnterVsPow is called when production vsPow is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsPow(ctx *VsPowContext) {}

// ExitVsPow is called when production vsPow is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsPow(ctx *VsPowContext) {}

// EnterVsMod is called when production vsMod is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsMod(ctx *VsModContext) {}

// ExitVsMod is called when production vsMod is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsMod(ctx *VsModContext) {}

// EnterVsAmp is called when production vsAmp is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsAmp(ctx *VsAmpContext) {}

// ExitVsAmp is called when production vsAmp is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsAmp(ctx *VsAmpContext) {}

// EnterVsAddressOf is called when production vsAddressOf is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsAddressOf(ctx *VsAddressOfContext) {}

// ExitVsAddressOf is called when production vsAddressOf is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsAddressOf(ctx *VsAddressOfContext) {}

// EnterVsNew is called when production vsNew is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsNew(ctx *VsNewContext) {}

// ExitVsNew is called when production vsNew is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsNew(ctx *VsNewContext) {}

// EnterVsPlusMinus is called when production vsPlusMinus is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsPlusMinus(ctx *VsPlusMinusContext) {}

// ExitVsPlusMinus is called when production vsPlusMinus is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsPlusMinus(ctx *VsPlusMinusContext) {}

// EnterVsIDiv is called when production vsIDiv is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsIDiv(ctx *VsIDivContext) {}

// ExitVsIDiv is called when production vsIDiv is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsIDiv(ctx *VsIDivContext) {}

// EnterVsOr is called when production vsOr is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsOr(ctx *VsOrContext) {}

// ExitVsOr is called when production vsOr is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsOr(ctx *VsOrContext) {}

// EnterVsLiteral is called when production vsLiteral is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsLiteral(ctx *VsLiteralContext) {}

// ExitVsLiteral is called when production vsLiteral is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsLiteral(ctx *VsLiteralContext) {}

// EnterVsEqv is called when production vsEqv is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsEqv(ctx *VsEqvContext) {}

// ExitVsEqv is called when production vsEqv is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsEqv(ctx *VsEqvContext) {}

// EnterVsImp is called when production vsImp is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsImp(ctx *VsImpContext) {}

// ExitVsImp is called when production vsImp is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsImp(ctx *VsImpContext) {}

// EnterVsComp is called when production vsComp is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsComp(ctx *VsCompContext) {}

// ExitVsComp is called when production vsComp is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsComp(ctx *VsCompContext) {}

// EnterVsMid is called when production vsMid is entered.
func (s *BaseVisualBasic6ParserListener) EnterVsMid(ctx *VsMidContext) {}

// ExitVsMid is called when production vsMid is exited.
func (s *BaseVisualBasic6ParserListener) ExitVsMid(ctx *VsMidContext) {}

// EnterVariableStmt is called when production variableStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterVariableStmt(ctx *VariableStmtContext) {}

// ExitVariableStmt is called when production variableStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitVariableStmt(ctx *VariableStmtContext) {}

// EnterVariableListStmt is called when production variableListStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterVariableListStmt(ctx *VariableListStmtContext) {}

// ExitVariableListStmt is called when production variableListStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitVariableListStmt(ctx *VariableListStmtContext) {}

// EnterVariableSubStmt is called when production variableSubStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterVariableSubStmt(ctx *VariableSubStmtContext) {}

// ExitVariableSubStmt is called when production variableSubStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitVariableSubStmt(ctx *VariableSubStmtContext) {}

// EnterWhileWendStmt is called when production whileWendStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterWhileWendStmt(ctx *WhileWendStmtContext) {}

// ExitWhileWendStmt is called when production whileWendStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitWhileWendStmt(ctx *WhileWendStmtContext) {}

// EnterWidthStmt is called when production widthStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterWidthStmt(ctx *WidthStmtContext) {}

// ExitWidthStmt is called when production widthStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitWidthStmt(ctx *WidthStmtContext) {}

// EnterWithStmt is called when production withStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterWithStmt(ctx *WithStmtContext) {}

// ExitWithStmt is called when production withStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitWithStmt(ctx *WithStmtContext) {}

// EnterWriteStmt is called when production writeStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterWriteStmt(ctx *WriteStmtContext) {}

// ExitWriteStmt is called when production writeStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitWriteStmt(ctx *WriteStmtContext) {}

// EnterExplicitCallStmt is called when production explicitCallStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterExplicitCallStmt(ctx *ExplicitCallStmtContext) {}

// ExitExplicitCallStmt is called when production explicitCallStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitExplicitCallStmt(ctx *ExplicitCallStmtContext) {}

// EnterECS_ProcedureCall is called when production eCS_ProcedureCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterECS_ProcedureCall(ctx *ECS_ProcedureCallContext) {}

// ExitECS_ProcedureCall is called when production eCS_ProcedureCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitECS_ProcedureCall(ctx *ECS_ProcedureCallContext) {}

// EnterECS_MemberProcedureCall is called when production eCS_MemberProcedureCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterECS_MemberProcedureCall(ctx *ECS_MemberProcedureCallContext) {
}

// ExitECS_MemberProcedureCall is called when production eCS_MemberProcedureCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitECS_MemberProcedureCall(ctx *ECS_MemberProcedureCallContext) {
}

// EnterImplicitCallStmt_InBlock is called when production implicitCallStmt_InBlock is entered.
func (s *BaseVisualBasic6ParserListener) EnterImplicitCallStmt_InBlock(ctx *ImplicitCallStmt_InBlockContext) {
}

// ExitImplicitCallStmt_InBlock is called when production implicitCallStmt_InBlock is exited.
func (s *BaseVisualBasic6ParserListener) ExitImplicitCallStmt_InBlock(ctx *ImplicitCallStmt_InBlockContext) {
}

// EnterICS_B_ProcedureCall is called when production iCS_B_ProcedureCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterICS_B_ProcedureCall(ctx *ICS_B_ProcedureCallContext) {}

// ExitICS_B_ProcedureCall is called when production iCS_B_ProcedureCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitICS_B_ProcedureCall(ctx *ICS_B_ProcedureCallContext) {}

// EnterICS_B_MemberProcedureCall is called when production iCS_B_MemberProcedureCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterICS_B_MemberProcedureCall(ctx *ICS_B_MemberProcedureCallContext) {
}

// ExitICS_B_MemberProcedureCall is called when production iCS_B_MemberProcedureCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitICS_B_MemberProcedureCall(ctx *ICS_B_MemberProcedureCallContext) {
}

// EnterImplicitCallStmt_InStmt is called when production implicitCallStmt_InStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterImplicitCallStmt_InStmt(ctx *ImplicitCallStmt_InStmtContext) {
}

// ExitImplicitCallStmt_InStmt is called when production implicitCallStmt_InStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitImplicitCallStmt_InStmt(ctx *ImplicitCallStmt_InStmtContext) {
}

// EnterICS_S_VariableOrProcedureCall is called when production iCS_S_VariableOrProcedureCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterICS_S_VariableOrProcedureCall(ctx *ICS_S_VariableOrProcedureCallContext) {
}

// ExitICS_S_VariableOrProcedureCall is called when production iCS_S_VariableOrProcedureCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitICS_S_VariableOrProcedureCall(ctx *ICS_S_VariableOrProcedureCallContext) {
}

// EnterICS_S_ProcedureOrArrayCall is called when production iCS_S_ProcedureOrArrayCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterICS_S_ProcedureOrArrayCall(ctx *ICS_S_ProcedureOrArrayCallContext) {
}

// ExitICS_S_ProcedureOrArrayCall is called when production iCS_S_ProcedureOrArrayCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitICS_S_ProcedureOrArrayCall(ctx *ICS_S_ProcedureOrArrayCallContext) {
}

// EnterICS_S_NestedProcedureCall is called when production iCS_S_NestedProcedureCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterICS_S_NestedProcedureCall(ctx *ICS_S_NestedProcedureCallContext) {
}

// ExitICS_S_NestedProcedureCall is called when production iCS_S_NestedProcedureCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitICS_S_NestedProcedureCall(ctx *ICS_S_NestedProcedureCallContext) {
}

// EnterICS_S_MembersCall is called when production iCS_S_MembersCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterICS_S_MembersCall(ctx *ICS_S_MembersCallContext) {}

// ExitICS_S_MembersCall is called when production iCS_S_MembersCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitICS_S_MembersCall(ctx *ICS_S_MembersCallContext) {}

// EnterICS_S_MemberCall is called when production iCS_S_MemberCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterICS_S_MemberCall(ctx *ICS_S_MemberCallContext) {}

// ExitICS_S_MemberCall is called when production iCS_S_MemberCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitICS_S_MemberCall(ctx *ICS_S_MemberCallContext) {}

// EnterICS_S_DictionaryCall is called when production iCS_S_DictionaryCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterICS_S_DictionaryCall(ctx *ICS_S_DictionaryCallContext) {
}

// ExitICS_S_DictionaryCall is called when production iCS_S_DictionaryCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitICS_S_DictionaryCall(ctx *ICS_S_DictionaryCallContext) {}

// EnterArgsCall is called when production argsCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterArgsCall(ctx *ArgsCallContext) {}

// ExitArgsCall is called when production argsCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitArgsCall(ctx *ArgsCallContext) {}

// EnterArgCall is called when production argCall is entered.
func (s *BaseVisualBasic6ParserListener) EnterArgCall(ctx *ArgCallContext) {}

// ExitArgCall is called when production argCall is exited.
func (s *BaseVisualBasic6ParserListener) ExitArgCall(ctx *ArgCallContext) {}

// EnterDictionaryCallStmt is called when production dictionaryCallStmt is entered.
func (s *BaseVisualBasic6ParserListener) EnterDictionaryCallStmt(ctx *DictionaryCallStmtContext) {}

// ExitDictionaryCallStmt is called when production dictionaryCallStmt is exited.
func (s *BaseVisualBasic6ParserListener) ExitDictionaryCallStmt(ctx *DictionaryCallStmtContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseVisualBasic6ParserListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseVisualBasic6ParserListener) ExitArgList(ctx *ArgListContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseVisualBasic6ParserListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseVisualBasic6ParserListener) ExitArg(ctx *ArgContext) {}

// EnterArgDefaultValue is called when production argDefaultValue is entered.
func (s *BaseVisualBasic6ParserListener) EnterArgDefaultValue(ctx *ArgDefaultValueContext) {}

// ExitArgDefaultValue is called when production argDefaultValue is exited.
func (s *BaseVisualBasic6ParserListener) ExitArgDefaultValue(ctx *ArgDefaultValueContext) {}

// EnterSubscripts is called when production subscripts is entered.
func (s *BaseVisualBasic6ParserListener) EnterSubscripts(ctx *SubscriptsContext) {}

// ExitSubscripts is called when production subscripts is exited.
func (s *BaseVisualBasic6ParserListener) ExitSubscripts(ctx *SubscriptsContext) {}

// EnterSubscript_ is called when production subscript_ is entered.
func (s *BaseVisualBasic6ParserListener) EnterSubscript_(ctx *Subscript_Context) {}

// ExitSubscript_ is called when production subscript_ is exited.
func (s *BaseVisualBasic6ParserListener) ExitSubscript_(ctx *Subscript_Context) {}

// EnterAmbiguousIdentifier is called when production ambiguousIdentifier is entered.
func (s *BaseVisualBasic6ParserListener) EnterAmbiguousIdentifier(ctx *AmbiguousIdentifierContext) {}

// ExitAmbiguousIdentifier is called when production ambiguousIdentifier is exited.
func (s *BaseVisualBasic6ParserListener) ExitAmbiguousIdentifier(ctx *AmbiguousIdentifierContext) {}

// EnterAsTypeClause is called when production asTypeClause is entered.
func (s *BaseVisualBasic6ParserListener) EnterAsTypeClause(ctx *AsTypeClauseContext) {}

// ExitAsTypeClause is called when production asTypeClause is exited.
func (s *BaseVisualBasic6ParserListener) ExitAsTypeClause(ctx *AsTypeClauseContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseVisualBasic6ParserListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseVisualBasic6ParserListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterCertainIdentifier is called when production certainIdentifier is entered.
func (s *BaseVisualBasic6ParserListener) EnterCertainIdentifier(ctx *CertainIdentifierContext) {}

// ExitCertainIdentifier is called when production certainIdentifier is exited.
func (s *BaseVisualBasic6ParserListener) ExitCertainIdentifier(ctx *CertainIdentifierContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseVisualBasic6ParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseVisualBasic6ParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterComplexType is called when production complexType is entered.
func (s *BaseVisualBasic6ParserListener) EnterComplexType(ctx *ComplexTypeContext) {}

// ExitComplexType is called when production complexType is exited.
func (s *BaseVisualBasic6ParserListener) ExitComplexType(ctx *ComplexTypeContext) {}

// EnterFieldLength is called when production fieldLength is entered.
func (s *BaseVisualBasic6ParserListener) EnterFieldLength(ctx *FieldLengthContext) {}

// ExitFieldLength is called when production fieldLength is exited.
func (s *BaseVisualBasic6ParserListener) ExitFieldLength(ctx *FieldLengthContext) {}

// EnterLetterrange is called when production letterrange is entered.
func (s *BaseVisualBasic6ParserListener) EnterLetterrange(ctx *LetterrangeContext) {}

// ExitLetterrange is called when production letterrange is exited.
func (s *BaseVisualBasic6ParserListener) ExitLetterrange(ctx *LetterrangeContext) {}

// EnterLineLabel is called when production lineLabel is entered.
func (s *BaseVisualBasic6ParserListener) EnterLineLabel(ctx *LineLabelContext) {}

// ExitLineLabel is called when production lineLabel is exited.
func (s *BaseVisualBasic6ParserListener) ExitLineLabel(ctx *LineLabelContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseVisualBasic6ParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseVisualBasic6ParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterPublicPrivateVisibility is called when production publicPrivateVisibility is entered.
func (s *BaseVisualBasic6ParserListener) EnterPublicPrivateVisibility(ctx *PublicPrivateVisibilityContext) {
}

// ExitPublicPrivateVisibility is called when production publicPrivateVisibility is exited.
func (s *BaseVisualBasic6ParserListener) ExitPublicPrivateVisibility(ctx *PublicPrivateVisibilityContext) {
}

// EnterPublicPrivateGlobalVisibility is called when production publicPrivateGlobalVisibility is entered.
func (s *BaseVisualBasic6ParserListener) EnterPublicPrivateGlobalVisibility(ctx *PublicPrivateGlobalVisibilityContext) {
}

// ExitPublicPrivateGlobalVisibility is called when production publicPrivateGlobalVisibility is exited.
func (s *BaseVisualBasic6ParserListener) ExitPublicPrivateGlobalVisibility(ctx *PublicPrivateGlobalVisibilityContext) {
}

// EnterType_ is called when production type_ is entered.
func (s *BaseVisualBasic6ParserListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseVisualBasic6ParserListener) ExitType_(ctx *Type_Context) {}

// EnterTypeHint is called when production typeHint is entered.
func (s *BaseVisualBasic6ParserListener) EnterTypeHint(ctx *TypeHintContext) {}

// ExitTypeHint is called when production typeHint is exited.
func (s *BaseVisualBasic6ParserListener) ExitTypeHint(ctx *TypeHintContext) {}

// EnterVisibility is called when production visibility is entered.
func (s *BaseVisualBasic6ParserListener) EnterVisibility(ctx *VisibilityContext) {}

// ExitVisibility is called when production visibility is exited.
func (s *BaseVisualBasic6ParserListener) ExitVisibility(ctx *VisibilityContext) {}

// EnterAmbiguousKeyword is called when production ambiguousKeyword is entered.
func (s *BaseVisualBasic6ParserListener) EnterAmbiguousKeyword(ctx *AmbiguousKeywordContext) {}

// ExitAmbiguousKeyword is called when production ambiguousKeyword is exited.
func (s *BaseVisualBasic6ParserListener) ExitAmbiguousKeyword(ctx *AmbiguousKeywordContext) {}
