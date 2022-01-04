// Code generated from WatParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wat // WatParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWatParserListener is a complete listener for a parse tree produced by WatParser.
type BaseWatParserListener struct{}

var _ WatParserListener = &BaseWatParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWatParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWatParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWatParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWatParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterValue is called when production value is entered.
func (s *BaseWatParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseWatParserListener) ExitValue(ctx *ValueContext) {}

// EnterName is called when production name is entered.
func (s *BaseWatParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseWatParserListener) ExitName(ctx *NameContext) {}

// EnterValue_type is called when production value_type is entered.
func (s *BaseWatParserListener) EnterValue_type(ctx *Value_typeContext) {}

// ExitValue_type is called when production value_type is exited.
func (s *BaseWatParserListener) ExitValue_type(ctx *Value_typeContext) {}

// EnterElem_type is called when production elem_type is entered.
func (s *BaseWatParserListener) EnterElem_type(ctx *Elem_typeContext) {}

// ExitElem_type is called when production elem_type is exited.
func (s *BaseWatParserListener) ExitElem_type(ctx *Elem_typeContext) {}

// EnterGlobal_type is called when production global_type is entered.
func (s *BaseWatParserListener) EnterGlobal_type(ctx *Global_typeContext) {}

// ExitGlobal_type is called when production global_type is exited.
func (s *BaseWatParserListener) ExitGlobal_type(ctx *Global_typeContext) {}

// EnterDef_type is called when production def_type is entered.
func (s *BaseWatParserListener) EnterDef_type(ctx *Def_typeContext) {}

// ExitDef_type is called when production def_type is exited.
func (s *BaseWatParserListener) ExitDef_type(ctx *Def_typeContext) {}

// EnterFunc_type is called when production func_type is entered.
func (s *BaseWatParserListener) EnterFunc_type(ctx *Func_typeContext) {}

// ExitFunc_type is called when production func_type is exited.
func (s *BaseWatParserListener) ExitFunc_type(ctx *Func_typeContext) {}

// EnterTable_type is called when production table_type is entered.
func (s *BaseWatParserListener) EnterTable_type(ctx *Table_typeContext) {}

// ExitTable_type is called when production table_type is exited.
func (s *BaseWatParserListener) ExitTable_type(ctx *Table_typeContext) {}

// EnterMemory_type is called when production memory_type is entered.
func (s *BaseWatParserListener) EnterMemory_type(ctx *Memory_typeContext) {}

// ExitMemory_type is called when production memory_type is exited.
func (s *BaseWatParserListener) ExitMemory_type(ctx *Memory_typeContext) {}

// EnterType_use is called when production type_use is entered.
func (s *BaseWatParserListener) EnterType_use(ctx *Type_useContext) {}

// ExitType_use is called when production type_use is exited.
func (s *BaseWatParserListener) ExitType_use(ctx *Type_useContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseWatParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseWatParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterVar_ is called when production var_ is entered.
func (s *BaseWatParserListener) EnterVar_(ctx *Var_Context) {}

// ExitVar_ is called when production var_ is exited.
func (s *BaseWatParserListener) ExitVar_(ctx *Var_Context) {}

// EnterBind_var is called when production bind_var is entered.
func (s *BaseWatParserListener) EnterBind_var(ctx *Bind_varContext) {}

// ExitBind_var is called when production bind_var is exited.
func (s *BaseWatParserListener) ExitBind_var(ctx *Bind_varContext) {}

// EnterInstr is called when production instr is entered.
func (s *BaseWatParserListener) EnterInstr(ctx *InstrContext) {}

// ExitInstr is called when production instr is exited.
func (s *BaseWatParserListener) ExitInstr(ctx *InstrContext) {}

// EnterPlain_instr is called when production plain_instr is entered.
func (s *BaseWatParserListener) EnterPlain_instr(ctx *Plain_instrContext) {}

// ExitPlain_instr is called when production plain_instr is exited.
func (s *BaseWatParserListener) ExitPlain_instr(ctx *Plain_instrContext) {}

// EnterCall_instr is called when production call_instr is entered.
func (s *BaseWatParserListener) EnterCall_instr(ctx *Call_instrContext) {}

// ExitCall_instr is called when production call_instr is exited.
func (s *BaseWatParserListener) ExitCall_instr(ctx *Call_instrContext) {}

// EnterCall_instr_params is called when production call_instr_params is entered.
func (s *BaseWatParserListener) EnterCall_instr_params(ctx *Call_instr_paramsContext) {}

// ExitCall_instr_params is called when production call_instr_params is exited.
func (s *BaseWatParserListener) ExitCall_instr_params(ctx *Call_instr_paramsContext) {}

// EnterCall_instr_instr is called when production call_instr_instr is entered.
func (s *BaseWatParserListener) EnterCall_instr_instr(ctx *Call_instr_instrContext) {}

// ExitCall_instr_instr is called when production call_instr_instr is exited.
func (s *BaseWatParserListener) ExitCall_instr_instr(ctx *Call_instr_instrContext) {}

// EnterCall_instr_params_instr is called when production call_instr_params_instr is entered.
func (s *BaseWatParserListener) EnterCall_instr_params_instr(ctx *Call_instr_params_instrContext) {}

// ExitCall_instr_params_instr is called when production call_instr_params_instr is exited.
func (s *BaseWatParserListener) ExitCall_instr_params_instr(ctx *Call_instr_params_instrContext) {}

// EnterCall_instr_results_instr is called when production call_instr_results_instr is entered.
func (s *BaseWatParserListener) EnterCall_instr_results_instr(ctx *Call_instr_results_instrContext) {}

// ExitCall_instr_results_instr is called when production call_instr_results_instr is exited.
func (s *BaseWatParserListener) ExitCall_instr_results_instr(ctx *Call_instr_results_instrContext) {}

// EnterBlock_instr is called when production block_instr is entered.
func (s *BaseWatParserListener) EnterBlock_instr(ctx *Block_instrContext) {}

// ExitBlock_instr is called when production block_instr is exited.
func (s *BaseWatParserListener) ExitBlock_instr(ctx *Block_instrContext) {}

// EnterBlock_type is called when production block_type is entered.
func (s *BaseWatParserListener) EnterBlock_type(ctx *Block_typeContext) {}

// ExitBlock_type is called when production block_type is exited.
func (s *BaseWatParserListener) ExitBlock_type(ctx *Block_typeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseWatParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseWatParserListener) ExitBlock(ctx *BlockContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseWatParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseWatParserListener) ExitExpr(ctx *ExprContext) {}

// EnterExpr1 is called when production expr1 is entered.
func (s *BaseWatParserListener) EnterExpr1(ctx *Expr1Context) {}

// ExitExpr1 is called when production expr1 is exited.
func (s *BaseWatParserListener) ExitExpr1(ctx *Expr1Context) {}

// EnterCall_expr_type is called when production call_expr_type is entered.
func (s *BaseWatParserListener) EnterCall_expr_type(ctx *Call_expr_typeContext) {}

// ExitCall_expr_type is called when production call_expr_type is exited.
func (s *BaseWatParserListener) ExitCall_expr_type(ctx *Call_expr_typeContext) {}

// EnterCall_expr_params is called when production call_expr_params is entered.
func (s *BaseWatParserListener) EnterCall_expr_params(ctx *Call_expr_paramsContext) {}

// ExitCall_expr_params is called when production call_expr_params is exited.
func (s *BaseWatParserListener) ExitCall_expr_params(ctx *Call_expr_paramsContext) {}

// EnterCall_expr_results is called when production call_expr_results is entered.
func (s *BaseWatParserListener) EnterCall_expr_results(ctx *Call_expr_resultsContext) {}

// ExitCall_expr_results is called when production call_expr_results is exited.
func (s *BaseWatParserListener) ExitCall_expr_results(ctx *Call_expr_resultsContext) {}

// EnterIf_block is called when production if_block is entered.
func (s *BaseWatParserListener) EnterIf_block(ctx *If_blockContext) {}

// ExitIf_block is called when production if_block is exited.
func (s *BaseWatParserListener) ExitIf_block(ctx *If_blockContext) {}

// EnterInstr_list is called when production instr_list is entered.
func (s *BaseWatParserListener) EnterInstr_list(ctx *Instr_listContext) {}

// ExitInstr_list is called when production instr_list is exited.
func (s *BaseWatParserListener) ExitInstr_list(ctx *Instr_listContext) {}

// EnterConst_expr is called when production const_expr is entered.
func (s *BaseWatParserListener) EnterConst_expr(ctx *Const_exprContext) {}

// ExitConst_expr is called when production const_expr is exited.
func (s *BaseWatParserListener) ExitConst_expr(ctx *Const_exprContext) {}

// EnterFunc_ is called when production func_ is entered.
func (s *BaseWatParserListener) EnterFunc_(ctx *Func_Context) {}

// ExitFunc_ is called when production func_ is exited.
func (s *BaseWatParserListener) ExitFunc_(ctx *Func_Context) {}

// EnterFunc_fields is called when production func_fields is entered.
func (s *BaseWatParserListener) EnterFunc_fields(ctx *Func_fieldsContext) {}

// ExitFunc_fields is called when production func_fields is exited.
func (s *BaseWatParserListener) ExitFunc_fields(ctx *Func_fieldsContext) {}

// EnterFunc_fields_import is called when production func_fields_import is entered.
func (s *BaseWatParserListener) EnterFunc_fields_import(ctx *Func_fields_importContext) {}

// ExitFunc_fields_import is called when production func_fields_import is exited.
func (s *BaseWatParserListener) ExitFunc_fields_import(ctx *Func_fields_importContext) {}

// EnterFunc_fields_import_result is called when production func_fields_import_result is entered.
func (s *BaseWatParserListener) EnterFunc_fields_import_result(ctx *Func_fields_import_resultContext) {
}

// ExitFunc_fields_import_result is called when production func_fields_import_result is exited.
func (s *BaseWatParserListener) ExitFunc_fields_import_result(ctx *Func_fields_import_resultContext) {
}

// EnterFunc_fields_body is called when production func_fields_body is entered.
func (s *BaseWatParserListener) EnterFunc_fields_body(ctx *Func_fields_bodyContext) {}

// ExitFunc_fields_body is called when production func_fields_body is exited.
func (s *BaseWatParserListener) ExitFunc_fields_body(ctx *Func_fields_bodyContext) {}

// EnterFunc_result_body is called when production func_result_body is entered.
func (s *BaseWatParserListener) EnterFunc_result_body(ctx *Func_result_bodyContext) {}

// ExitFunc_result_body is called when production func_result_body is exited.
func (s *BaseWatParserListener) ExitFunc_result_body(ctx *Func_result_bodyContext) {}

// EnterFunc_body is called when production func_body is entered.
func (s *BaseWatParserListener) EnterFunc_body(ctx *Func_bodyContext) {}

// ExitFunc_body is called when production func_body is exited.
func (s *BaseWatParserListener) ExitFunc_body(ctx *Func_bodyContext) {}

// EnterOffset is called when production offset is entered.
func (s *BaseWatParserListener) EnterOffset(ctx *OffsetContext) {}

// ExitOffset is called when production offset is exited.
func (s *BaseWatParserListener) ExitOffset(ctx *OffsetContext) {}

// EnterElem is called when production elem is entered.
func (s *BaseWatParserListener) EnterElem(ctx *ElemContext) {}

// ExitElem is called when production elem is exited.
func (s *BaseWatParserListener) ExitElem(ctx *ElemContext) {}

// EnterTable is called when production table is entered.
func (s *BaseWatParserListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseWatParserListener) ExitTable(ctx *TableContext) {}

// EnterTable_fields is called when production table_fields is entered.
func (s *BaseWatParserListener) EnterTable_fields(ctx *Table_fieldsContext) {}

// ExitTable_fields is called when production table_fields is exited.
func (s *BaseWatParserListener) ExitTable_fields(ctx *Table_fieldsContext) {}

// EnterData is called when production data is entered.
func (s *BaseWatParserListener) EnterData(ctx *DataContext) {}

// ExitData is called when production data is exited.
func (s *BaseWatParserListener) ExitData(ctx *DataContext) {}

// EnterMemory is called when production memory is entered.
func (s *BaseWatParserListener) EnterMemory(ctx *MemoryContext) {}

// ExitMemory is called when production memory is exited.
func (s *BaseWatParserListener) ExitMemory(ctx *MemoryContext) {}

// EnterMemory_fields is called when production memory_fields is entered.
func (s *BaseWatParserListener) EnterMemory_fields(ctx *Memory_fieldsContext) {}

// ExitMemory_fields is called when production memory_fields is exited.
func (s *BaseWatParserListener) ExitMemory_fields(ctx *Memory_fieldsContext) {}

// EnterSglobal is called when production sglobal is entered.
func (s *BaseWatParserListener) EnterSglobal(ctx *SglobalContext) {}

// ExitSglobal is called when production sglobal is exited.
func (s *BaseWatParserListener) ExitSglobal(ctx *SglobalContext) {}

// EnterGlobal_fields is called when production global_fields is entered.
func (s *BaseWatParserListener) EnterGlobal_fields(ctx *Global_fieldsContext) {}

// ExitGlobal_fields is called when production global_fields is exited.
func (s *BaseWatParserListener) ExitGlobal_fields(ctx *Global_fieldsContext) {}

// EnterImport_desc is called when production import_desc is entered.
func (s *BaseWatParserListener) EnterImport_desc(ctx *Import_descContext) {}

// ExitImport_desc is called when production import_desc is exited.
func (s *BaseWatParserListener) ExitImport_desc(ctx *Import_descContext) {}

// EnterSimport is called when production simport is entered.
func (s *BaseWatParserListener) EnterSimport(ctx *SimportContext) {}

// ExitSimport is called when production simport is exited.
func (s *BaseWatParserListener) ExitSimport(ctx *SimportContext) {}

// EnterInline_import is called when production inline_import is entered.
func (s *BaseWatParserListener) EnterInline_import(ctx *Inline_importContext) {}

// ExitInline_import is called when production inline_import is exited.
func (s *BaseWatParserListener) ExitInline_import(ctx *Inline_importContext) {}

// EnterExport_desc is called when production export_desc is entered.
func (s *BaseWatParserListener) EnterExport_desc(ctx *Export_descContext) {}

// ExitExport_desc is called when production export_desc is exited.
func (s *BaseWatParserListener) ExitExport_desc(ctx *Export_descContext) {}

// EnterExport_ is called when production export_ is entered.
func (s *BaseWatParserListener) EnterExport_(ctx *Export_Context) {}

// ExitExport_ is called when production export_ is exited.
func (s *BaseWatParserListener) ExitExport_(ctx *Export_Context) {}

// EnterInline_export is called when production inline_export is entered.
func (s *BaseWatParserListener) EnterInline_export(ctx *Inline_exportContext) {}

// ExitInline_export is called when production inline_export is exited.
func (s *BaseWatParserListener) ExitInline_export(ctx *Inline_exportContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseWatParserListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseWatParserListener) ExitType_(ctx *Type_Context) {}

// EnterType_def is called when production type_def is entered.
func (s *BaseWatParserListener) EnterType_def(ctx *Type_defContext) {}

// ExitType_def is called when production type_def is exited.
func (s *BaseWatParserListener) ExitType_def(ctx *Type_defContext) {}

// EnterStart_ is called when production start_ is entered.
func (s *BaseWatParserListener) EnterStart_(ctx *Start_Context) {}

// ExitStart_ is called when production start_ is exited.
func (s *BaseWatParserListener) ExitStart_(ctx *Start_Context) {}

// EnterModule_field is called when production module_field is entered.
func (s *BaseWatParserListener) EnterModule_field(ctx *Module_fieldContext) {}

// ExitModule_field is called when production module_field is exited.
func (s *BaseWatParserListener) ExitModule_field(ctx *Module_fieldContext) {}

// EnterModule_ is called when production module_ is entered.
func (s *BaseWatParserListener) EnterModule_(ctx *Module_Context) {}

// ExitModule_ is called when production module_ is exited.
func (s *BaseWatParserListener) ExitModule_(ctx *Module_Context) {}

// EnterScript_module is called when production script_module is entered.
func (s *BaseWatParserListener) EnterScript_module(ctx *Script_moduleContext) {}

// ExitScript_module is called when production script_module is exited.
func (s *BaseWatParserListener) ExitScript_module(ctx *Script_moduleContext) {}

// EnterAction_ is called when production action_ is entered.
func (s *BaseWatParserListener) EnterAction_(ctx *Action_Context) {}

// ExitAction_ is called when production action_ is exited.
func (s *BaseWatParserListener) ExitAction_(ctx *Action_Context) {}

// EnterAssertion is called when production assertion is entered.
func (s *BaseWatParserListener) EnterAssertion(ctx *AssertionContext) {}

// ExitAssertion is called when production assertion is exited.
func (s *BaseWatParserListener) ExitAssertion(ctx *AssertionContext) {}

// EnterCmd is called when production cmd is entered.
func (s *BaseWatParserListener) EnterCmd(ctx *CmdContext) {}

// ExitCmd is called when production cmd is exited.
func (s *BaseWatParserListener) ExitCmd(ctx *CmdContext) {}

// EnterMeta is called when production meta is entered.
func (s *BaseWatParserListener) EnterMeta(ctx *MetaContext) {}

// ExitMeta is called when production meta is exited.
func (s *BaseWatParserListener) ExitMeta(ctx *MetaContext) {}

// EnterWconst is called when production wconst is entered.
func (s *BaseWatParserListener) EnterWconst(ctx *WconstContext) {}

// ExitWconst is called when production wconst is exited.
func (s *BaseWatParserListener) ExitWconst(ctx *WconstContext) {}

// EnterConst_list is called when production const_list is entered.
func (s *BaseWatParserListener) EnterConst_list(ctx *Const_listContext) {}

// ExitConst_list is called when production const_list is exited.
func (s *BaseWatParserListener) ExitConst_list(ctx *Const_listContext) {}

// EnterScript is called when production script is entered.
func (s *BaseWatParserListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseWatParserListener) ExitScript(ctx *ScriptContext) {}

// EnterModule is called when production module is entered.
func (s *BaseWatParserListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseWatParserListener) ExitModule(ctx *ModuleContext) {}
