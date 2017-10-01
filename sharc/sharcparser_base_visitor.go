// Generated from SHARCParser.g4 by ANTLR 4.7.

package sharc // SHARCParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSHARCParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSHARCParserVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitStmt_atom(ctx *Stmt_atomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSec(ctx *SecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSeg(ctx *SegContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitEnd_seg(ctx *End_segContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSeg_qualifier(ctx *Seg_qualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSeg_qualifier1(ctx *Seg_qualifier1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSeg_qualifier2(ctx *Seg_qualifier2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSeg_qualifier3(ctx *Seg_qualifier3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSeg_qualifier_1(ctx *Seg_qualifier_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSeg_qualifier_2(ctx *Seg_qualifier_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSeg_qualifier_3(ctx *Seg_qualifier_3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDeclaration_exp1(ctx *Declaration_exp1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDeclaration_exp2(ctx *Declaration_exp2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDeclaration_exp3(ctx *Declaration_exp3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDeclaration_exp4(ctx *Declaration_exp4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDeclaration_exp5(ctx *Declaration_exp5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDeclaration_exp_f1(ctx *Declaration_exp_f1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDeclaration_exp_f2(ctx *Declaration_exp_f2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitInitExpression(ctx *InitExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitVar_addr(ctx *Var_addrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitValue_exp(ctx *Value_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitValue_exp2(ctx *Value_exp2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitCompute(ctx *ComputeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitIf_compute_mov(ctx *If_compute_movContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitIf_compute_mov_exp(ctx *If_compute_mov_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitCompute_mov_exp(ctx *Compute_mov_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_1(ctx *Mov_exp_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_1_1(ctx *Mov_exp_1_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_1_2(ctx *Mov_exp_1_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_3a(ctx *Mov_exp_3aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_3b(ctx *Mov_exp_3bContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_3c(ctx *Mov_exp_3cContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_3d(ctx *Mov_exp_3dContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_4a(ctx *Mov_exp_4aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_4b(ctx *Mov_exp_4bContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_4c(ctx *Mov_exp_4cContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_4d(ctx *Mov_exp_4dContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_5(ctx *Mov_exp_5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_6a(ctx *Mov_exp_6aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_6b(ctx *Mov_exp_6bContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMov_exp_7(ctx *Mov_exp_7Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_ia_mb(ctx *Mem_addr_ia_mbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_ic_md(ctx *Mem_addr_ic_mdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_md_ic(ctx *Mem_addr_md_icContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_mb_ia(ctx *Mem_addr_mb_iaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_ia_int(ctx *Mem_addr_ia_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_ic_int(ctx *Mem_addr_ic_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_int_ia(ctx *Mem_addr_int_iaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_int_ic(ctx *Mem_addr_int_icContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_int(ctx *Mem_addr_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_int_(ctx *Mem_addr_int_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_dm_ia_mb(ctx *Mem_addr_dm_ia_mbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_pm_ic_md(ctx *Mem_addr_pm_ic_mdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_dm_mb_ia(ctx *Mem_addr_dm_mb_iaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_pm_md_ic(ctx *Mem_addr_pm_md_icContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_dm_ia_int(ctx *Mem_addr_dm_ia_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_pm_ic_int(ctx *Mem_addr_pm_ic_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_dm_int_ia(ctx *Mem_addr_dm_int_iaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_pm_int_ic(ctx *Mem_addr_pm_int_icContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_dm_int(ctx *Mem_addr_dm_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMem_addr_pm_int(ctx *Mem_addr_pm_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFixpoint_alu_op(ctx *Fixpoint_alu_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitR_exp(ctx *R_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFloating_point_alu_op(ctx *Floating_point_alu_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitF_exp(ctx *F_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMulti_op(ctx *Multi_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMulti_r(ctx *Multi_rContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMulti_exp_r(ctx *Multi_exp_rContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMulti_exp_mrf(ctx *Multi_exp_mrfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMulti_exp_mrb(ctx *Multi_exp_mrbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMr(ctx *MrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitShifter_op(ctx *Shifter_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitShifter_exp(ctx *Shifter_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSec_op(ctx *Sec_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitSec_op2(ctx *Sec_op2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitBit_data(ctx *Bit_dataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitAdd_or_sub(ctx *Add_or_subContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDual_op(ctx *Dual_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDual_add_r(ctx *Dual_add_rContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitParallel_multi(ctx *Parallel_multiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_exp(ctx *Flow_control_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_contorl_8(ctx *Flow_contorl_8Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_contorl_8_exp(ctx *Flow_contorl_8_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_9_and_11(ctx *Flow_control_9_and_11Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_9_and_11_exp(ctx *Flow_control_9_and_11_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_10(ctx *Flow_control_10Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_10_frag(ctx *Flow_control_10_fragContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_12(ctx *Flow_control_12Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitLcntr_v(ctx *Lcntr_vContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_13(ctx *Flow_control_13Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_8a(ctx *Flow_control_8aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_8b(ctx *Flow_control_8bContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_9a(ctx *Flow_control_9aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_9b(ctx *Flow_control_9bContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_11a(ctx *Flow_control_11aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitFlow_control_11b(ctx *Flow_control_11bContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_addr_int_or_pc(ctx *Jump_addr_int_or_pcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_addr_md_or_pc(ctx *Jump_addr_md_or_pcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_addr_pc(ctx *Jump_addr_pcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_addr_int(ctx *Jump_addr_intContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_modifier(ctx *Jump_modifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_modifier_(ctx *Jump_modifier_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_modifier_1(ctx *Jump_modifier_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_modifier2(ctx *Jump_modifier2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_modifier3(ctx *Jump_modifier3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_modifier3_(ctx *Jump_modifier3_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitJump_modifier3_1(ctx *Jump_modifier3_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitImm_mov_exp(ctx *Imm_mov_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitImm_mov_14a(ctx *Imm_mov_14aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitImm_mov_15a(ctx *Imm_mov_15aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitImm_mov_14b(ctx *Imm_mov_14bContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitImm_mov_15b(ctx *Imm_mov_15bContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitImm_mov_16(ctx *Imm_mov_16Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitImm_mov_17(ctx *Imm_mov_17Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitU_reg2(ctx *U_reg2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMisc_exp(ctx *Misc_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMisc_20(ctx *Misc_20Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDirective_exp(ctx *Directive_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitB_reg(ctx *B_regContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitL_reg(ctx *L_regContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitR_reg(ctx *R_regContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitF_reg(ctx *F_regContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitS_reg(ctx *S_regContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitIa(ctx *IaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMb(ctx *MbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitIc(ctx *IcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMd(ctx *MdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitI_reg(ctx *I_regContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitM_reg(ctx *M_regContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitDag_reg(ctx *Dag_regContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitD_reg(ctx *D_regContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitU_reg(ctx *U_regContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitCcondition(ctx *CconditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMulti_mod1(ctx *Multi_mod1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMulti_mod1_(ctx *Multi_mod1_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMulti_mod2(ctx *Multi_mod2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMulti_mod2_(ctx *Multi_mod2_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitR3_0(ctx *R3_0Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitR7_4(ctx *R7_4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitR11_8(ctx *R11_8Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitR15_12(ctx *R15_12Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitF3_0(ctx *F3_0Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitF7_4(ctx *F7_4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitF11_8(ctx *F11_8Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitF15_12(ctx *F15_12Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitAddr(ctx *AddrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMrf(ctx *MrfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSHARCParserVisitor) VisitMrb(ctx *MrbContext) interface{} {
	return v.VisitChildren(ctx)
}
