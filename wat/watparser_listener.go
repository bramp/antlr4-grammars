// Code generated from WatParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wat // WatParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// WatParserListener is a complete listener for a parse tree produced by WatParser.
type WatParserListener interface {
	antlr.ParseTreeListener

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterValue_type is called when entering the value_type production.
	EnterValue_type(c *Value_typeContext)

	// EnterElem_type is called when entering the elem_type production.
	EnterElem_type(c *Elem_typeContext)

	// EnterGlobal_type is called when entering the global_type production.
	EnterGlobal_type(c *Global_typeContext)

	// EnterDef_type is called when entering the def_type production.
	EnterDef_type(c *Def_typeContext)

	// EnterFunc_type is called when entering the func_type production.
	EnterFunc_type(c *Func_typeContext)

	// EnterTable_type is called when entering the table_type production.
	EnterTable_type(c *Table_typeContext)

	// EnterMemory_type is called when entering the memory_type production.
	EnterMemory_type(c *Memory_typeContext)

	// EnterType_use is called when entering the type_use production.
	EnterType_use(c *Type_useContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterVar_ is called when entering the var_ production.
	EnterVar_(c *Var_Context)

	// EnterBind_var is called when entering the bind_var production.
	EnterBind_var(c *Bind_varContext)

	// EnterInstr is called when entering the instr production.
	EnterInstr(c *InstrContext)

	// EnterPlain_instr is called when entering the plain_instr production.
	EnterPlain_instr(c *Plain_instrContext)

	// EnterCall_instr is called when entering the call_instr production.
	EnterCall_instr(c *Call_instrContext)

	// EnterCall_instr_params is called when entering the call_instr_params production.
	EnterCall_instr_params(c *Call_instr_paramsContext)

	// EnterCall_instr_instr is called when entering the call_instr_instr production.
	EnterCall_instr_instr(c *Call_instr_instrContext)

	// EnterCall_instr_params_instr is called when entering the call_instr_params_instr production.
	EnterCall_instr_params_instr(c *Call_instr_params_instrContext)

	// EnterCall_instr_results_instr is called when entering the call_instr_results_instr production.
	EnterCall_instr_results_instr(c *Call_instr_results_instrContext)

	// EnterBlock_instr is called when entering the block_instr production.
	EnterBlock_instr(c *Block_instrContext)

	// EnterBlock_type is called when entering the block_type production.
	EnterBlock_type(c *Block_typeContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExpr1 is called when entering the expr1 production.
	EnterExpr1(c *Expr1Context)

	// EnterCall_expr_type is called when entering the call_expr_type production.
	EnterCall_expr_type(c *Call_expr_typeContext)

	// EnterCall_expr_params is called when entering the call_expr_params production.
	EnterCall_expr_params(c *Call_expr_paramsContext)

	// EnterCall_expr_results is called when entering the call_expr_results production.
	EnterCall_expr_results(c *Call_expr_resultsContext)

	// EnterIf_block is called when entering the if_block production.
	EnterIf_block(c *If_blockContext)

	// EnterInstr_list is called when entering the instr_list production.
	EnterInstr_list(c *Instr_listContext)

	// EnterConst_expr is called when entering the const_expr production.
	EnterConst_expr(c *Const_exprContext)

	// EnterFunc_ is called when entering the func_ production.
	EnterFunc_(c *Func_Context)

	// EnterFunc_fields is called when entering the func_fields production.
	EnterFunc_fields(c *Func_fieldsContext)

	// EnterFunc_fields_import is called when entering the func_fields_import production.
	EnterFunc_fields_import(c *Func_fields_importContext)

	// EnterFunc_fields_import_result is called when entering the func_fields_import_result production.
	EnterFunc_fields_import_result(c *Func_fields_import_resultContext)

	// EnterFunc_fields_body is called when entering the func_fields_body production.
	EnterFunc_fields_body(c *Func_fields_bodyContext)

	// EnterFunc_result_body is called when entering the func_result_body production.
	EnterFunc_result_body(c *Func_result_bodyContext)

	// EnterFunc_body is called when entering the func_body production.
	EnterFunc_body(c *Func_bodyContext)

	// EnterOffset is called when entering the offset production.
	EnterOffset(c *OffsetContext)

	// EnterElem is called when entering the elem production.
	EnterElem(c *ElemContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterTable_fields is called when entering the table_fields production.
	EnterTable_fields(c *Table_fieldsContext)

	// EnterData is called when entering the data production.
	EnterData(c *DataContext)

	// EnterMemory is called when entering the memory production.
	EnterMemory(c *MemoryContext)

	// EnterMemory_fields is called when entering the memory_fields production.
	EnterMemory_fields(c *Memory_fieldsContext)

	// EnterSglobal is called when entering the sglobal production.
	EnterSglobal(c *SglobalContext)

	// EnterGlobal_fields is called when entering the global_fields production.
	EnterGlobal_fields(c *Global_fieldsContext)

	// EnterImport_desc is called when entering the import_desc production.
	EnterImport_desc(c *Import_descContext)

	// EnterSimport is called when entering the simport production.
	EnterSimport(c *SimportContext)

	// EnterInline_import is called when entering the inline_import production.
	EnterInline_import(c *Inline_importContext)

	// EnterExport_desc is called when entering the export_desc production.
	EnterExport_desc(c *Export_descContext)

	// EnterExport_ is called when entering the export_ production.
	EnterExport_(c *Export_Context)

	// EnterInline_export is called when entering the inline_export production.
	EnterInline_export(c *Inline_exportContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterType_def is called when entering the type_def production.
	EnterType_def(c *Type_defContext)

	// EnterStart_ is called when entering the start_ production.
	EnterStart_(c *Start_Context)

	// EnterModule_field is called when entering the module_field production.
	EnterModule_field(c *Module_fieldContext)

	// EnterModule_ is called when entering the module_ production.
	EnterModule_(c *Module_Context)

	// EnterScript_module is called when entering the script_module production.
	EnterScript_module(c *Script_moduleContext)

	// EnterAction_ is called when entering the action_ production.
	EnterAction_(c *Action_Context)

	// EnterAssertion is called when entering the assertion production.
	EnterAssertion(c *AssertionContext)

	// EnterCmd is called when entering the cmd production.
	EnterCmd(c *CmdContext)

	// EnterMeta is called when entering the meta production.
	EnterMeta(c *MetaContext)

	// EnterWconst is called when entering the wconst production.
	EnterWconst(c *WconstContext)

	// EnterConst_list is called when entering the const_list production.
	EnterConst_list(c *Const_listContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitValue_type is called when exiting the value_type production.
	ExitValue_type(c *Value_typeContext)

	// ExitElem_type is called when exiting the elem_type production.
	ExitElem_type(c *Elem_typeContext)

	// ExitGlobal_type is called when exiting the global_type production.
	ExitGlobal_type(c *Global_typeContext)

	// ExitDef_type is called when exiting the def_type production.
	ExitDef_type(c *Def_typeContext)

	// ExitFunc_type is called when exiting the func_type production.
	ExitFunc_type(c *Func_typeContext)

	// ExitTable_type is called when exiting the table_type production.
	ExitTable_type(c *Table_typeContext)

	// ExitMemory_type is called when exiting the memory_type production.
	ExitMemory_type(c *Memory_typeContext)

	// ExitType_use is called when exiting the type_use production.
	ExitType_use(c *Type_useContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitVar_ is called when exiting the var_ production.
	ExitVar_(c *Var_Context)

	// ExitBind_var is called when exiting the bind_var production.
	ExitBind_var(c *Bind_varContext)

	// ExitInstr is called when exiting the instr production.
	ExitInstr(c *InstrContext)

	// ExitPlain_instr is called when exiting the plain_instr production.
	ExitPlain_instr(c *Plain_instrContext)

	// ExitCall_instr is called when exiting the call_instr production.
	ExitCall_instr(c *Call_instrContext)

	// ExitCall_instr_params is called when exiting the call_instr_params production.
	ExitCall_instr_params(c *Call_instr_paramsContext)

	// ExitCall_instr_instr is called when exiting the call_instr_instr production.
	ExitCall_instr_instr(c *Call_instr_instrContext)

	// ExitCall_instr_params_instr is called when exiting the call_instr_params_instr production.
	ExitCall_instr_params_instr(c *Call_instr_params_instrContext)

	// ExitCall_instr_results_instr is called when exiting the call_instr_results_instr production.
	ExitCall_instr_results_instr(c *Call_instr_results_instrContext)

	// ExitBlock_instr is called when exiting the block_instr production.
	ExitBlock_instr(c *Block_instrContext)

	// ExitBlock_type is called when exiting the block_type production.
	ExitBlock_type(c *Block_typeContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExpr1 is called when exiting the expr1 production.
	ExitExpr1(c *Expr1Context)

	// ExitCall_expr_type is called when exiting the call_expr_type production.
	ExitCall_expr_type(c *Call_expr_typeContext)

	// ExitCall_expr_params is called when exiting the call_expr_params production.
	ExitCall_expr_params(c *Call_expr_paramsContext)

	// ExitCall_expr_results is called when exiting the call_expr_results production.
	ExitCall_expr_results(c *Call_expr_resultsContext)

	// ExitIf_block is called when exiting the if_block production.
	ExitIf_block(c *If_blockContext)

	// ExitInstr_list is called when exiting the instr_list production.
	ExitInstr_list(c *Instr_listContext)

	// ExitConst_expr is called when exiting the const_expr production.
	ExitConst_expr(c *Const_exprContext)

	// ExitFunc_ is called when exiting the func_ production.
	ExitFunc_(c *Func_Context)

	// ExitFunc_fields is called when exiting the func_fields production.
	ExitFunc_fields(c *Func_fieldsContext)

	// ExitFunc_fields_import is called when exiting the func_fields_import production.
	ExitFunc_fields_import(c *Func_fields_importContext)

	// ExitFunc_fields_import_result is called when exiting the func_fields_import_result production.
	ExitFunc_fields_import_result(c *Func_fields_import_resultContext)

	// ExitFunc_fields_body is called when exiting the func_fields_body production.
	ExitFunc_fields_body(c *Func_fields_bodyContext)

	// ExitFunc_result_body is called when exiting the func_result_body production.
	ExitFunc_result_body(c *Func_result_bodyContext)

	// ExitFunc_body is called when exiting the func_body production.
	ExitFunc_body(c *Func_bodyContext)

	// ExitOffset is called when exiting the offset production.
	ExitOffset(c *OffsetContext)

	// ExitElem is called when exiting the elem production.
	ExitElem(c *ElemContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitTable_fields is called when exiting the table_fields production.
	ExitTable_fields(c *Table_fieldsContext)

	// ExitData is called when exiting the data production.
	ExitData(c *DataContext)

	// ExitMemory is called when exiting the memory production.
	ExitMemory(c *MemoryContext)

	// ExitMemory_fields is called when exiting the memory_fields production.
	ExitMemory_fields(c *Memory_fieldsContext)

	// ExitSglobal is called when exiting the sglobal production.
	ExitSglobal(c *SglobalContext)

	// ExitGlobal_fields is called when exiting the global_fields production.
	ExitGlobal_fields(c *Global_fieldsContext)

	// ExitImport_desc is called when exiting the import_desc production.
	ExitImport_desc(c *Import_descContext)

	// ExitSimport is called when exiting the simport production.
	ExitSimport(c *SimportContext)

	// ExitInline_import is called when exiting the inline_import production.
	ExitInline_import(c *Inline_importContext)

	// ExitExport_desc is called when exiting the export_desc production.
	ExitExport_desc(c *Export_descContext)

	// ExitExport_ is called when exiting the export_ production.
	ExitExport_(c *Export_Context)

	// ExitInline_export is called when exiting the inline_export production.
	ExitInline_export(c *Inline_exportContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitType_def is called when exiting the type_def production.
	ExitType_def(c *Type_defContext)

	// ExitStart_ is called when exiting the start_ production.
	ExitStart_(c *Start_Context)

	// ExitModule_field is called when exiting the module_field production.
	ExitModule_field(c *Module_fieldContext)

	// ExitModule_ is called when exiting the module_ production.
	ExitModule_(c *Module_Context)

	// ExitScript_module is called when exiting the script_module production.
	ExitScript_module(c *Script_moduleContext)

	// ExitAction_ is called when exiting the action_ production.
	ExitAction_(c *Action_Context)

	// ExitAssertion is called when exiting the assertion production.
	ExitAssertion(c *AssertionContext)

	// ExitCmd is called when exiting the cmd production.
	ExitCmd(c *CmdContext)

	// ExitMeta is called when exiting the meta production.
	ExitMeta(c *MetaContext)

	// ExitWconst is called when exiting the wconst production.
	ExitWconst(c *WconstContext)

	// ExitConst_list is called when exiting the const_list production.
	ExitConst_list(c *Const_listContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)
}
