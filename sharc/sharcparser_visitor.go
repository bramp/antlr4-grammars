// Generated from SHARCParser.g4 by ANTLR 4.7.

package sharc // SHARCParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SHARCParser.
type SHARCParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SHARCParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by SHARCParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by SHARCParser#stmt_atom.
	VisitStmt_atom(ctx *Stmt_atomContext) interface{}

	// Visit a parse tree produced by SHARCParser#sec.
	VisitSec(ctx *SecContext) interface{}

	// Visit a parse tree produced by SHARCParser#seg.
	VisitSeg(ctx *SegContext) interface{}

	// Visit a parse tree produced by SHARCParser#end_seg.
	VisitEnd_seg(ctx *End_segContext) interface{}

	// Visit a parse tree produced by SHARCParser#seg_qualifier.
	VisitSeg_qualifier(ctx *Seg_qualifierContext) interface{}

	// Visit a parse tree produced by SHARCParser#seg_qualifier1.
	VisitSeg_qualifier1(ctx *Seg_qualifier1Context) interface{}

	// Visit a parse tree produced by SHARCParser#seg_qualifier2.
	VisitSeg_qualifier2(ctx *Seg_qualifier2Context) interface{}

	// Visit a parse tree produced by SHARCParser#seg_qualifier3.
	VisitSeg_qualifier3(ctx *Seg_qualifier3Context) interface{}

	// Visit a parse tree produced by SHARCParser#seg_qualifier_1.
	VisitSeg_qualifier_1(ctx *Seg_qualifier_1Context) interface{}

	// Visit a parse tree produced by SHARCParser#seg_qualifier_2.
	VisitSeg_qualifier_2(ctx *Seg_qualifier_2Context) interface{}

	// Visit a parse tree produced by SHARCParser#seg_qualifier_3.
	VisitSeg_qualifier_3(ctx *Seg_qualifier_3Context) interface{}

	// Visit a parse tree produced by SHARCParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by SHARCParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by SHARCParser#declaration_exp1.
	VisitDeclaration_exp1(ctx *Declaration_exp1Context) interface{}

	// Visit a parse tree produced by SHARCParser#declaration_exp2.
	VisitDeclaration_exp2(ctx *Declaration_exp2Context) interface{}

	// Visit a parse tree produced by SHARCParser#declaration_exp3.
	VisitDeclaration_exp3(ctx *Declaration_exp3Context) interface{}

	// Visit a parse tree produced by SHARCParser#declaration_exp4.
	VisitDeclaration_exp4(ctx *Declaration_exp4Context) interface{}

	// Visit a parse tree produced by SHARCParser#declaration_exp5.
	VisitDeclaration_exp5(ctx *Declaration_exp5Context) interface{}

	// Visit a parse tree produced by SHARCParser#declaration_exp_f1.
	VisitDeclaration_exp_f1(ctx *Declaration_exp_f1Context) interface{}

	// Visit a parse tree produced by SHARCParser#declaration_exp_f2.
	VisitDeclaration_exp_f2(ctx *Declaration_exp_f2Context) interface{}

	// Visit a parse tree produced by SHARCParser#initExpression.
	VisitInitExpression(ctx *InitExpressionContext) interface{}

	// Visit a parse tree produced by SHARCParser#var_addr.
	VisitVar_addr(ctx *Var_addrContext) interface{}

	// Visit a parse tree produced by SHARCParser#value_exp.
	VisitValue_exp(ctx *Value_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#value_exp2.
	VisitValue_exp2(ctx *Value_exp2Context) interface{}

	// Visit a parse tree produced by SHARCParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by SHARCParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by SHARCParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by SHARCParser#compute.
	VisitCompute(ctx *ComputeContext) interface{}

	// Visit a parse tree produced by SHARCParser#if_compute_mov.
	VisitIf_compute_mov(ctx *If_compute_movContext) interface{}

	// Visit a parse tree produced by SHARCParser#if_compute_mov_exp.
	VisitIf_compute_mov_exp(ctx *If_compute_mov_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#compute_mov_exp.
	VisitCompute_mov_exp(ctx *Compute_mov_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_1.
	VisitMov_exp_1(ctx *Mov_exp_1Context) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_1_1.
	VisitMov_exp_1_1(ctx *Mov_exp_1_1Context) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_1_2.
	VisitMov_exp_1_2(ctx *Mov_exp_1_2Context) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_3a.
	VisitMov_exp_3a(ctx *Mov_exp_3aContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_3b.
	VisitMov_exp_3b(ctx *Mov_exp_3bContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_3c.
	VisitMov_exp_3c(ctx *Mov_exp_3cContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_3d.
	VisitMov_exp_3d(ctx *Mov_exp_3dContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_4a.
	VisitMov_exp_4a(ctx *Mov_exp_4aContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_4b.
	VisitMov_exp_4b(ctx *Mov_exp_4bContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_4c.
	VisitMov_exp_4c(ctx *Mov_exp_4cContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_4d.
	VisitMov_exp_4d(ctx *Mov_exp_4dContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_5.
	VisitMov_exp_5(ctx *Mov_exp_5Context) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_6a.
	VisitMov_exp_6a(ctx *Mov_exp_6aContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_6b.
	VisitMov_exp_6b(ctx *Mov_exp_6bContext) interface{}

	// Visit a parse tree produced by SHARCParser#mov_exp_7.
	VisitMov_exp_7(ctx *Mov_exp_7Context) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_ia_mb.
	VisitMem_addr_ia_mb(ctx *Mem_addr_ia_mbContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_ic_md.
	VisitMem_addr_ic_md(ctx *Mem_addr_ic_mdContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_md_ic.
	VisitMem_addr_md_ic(ctx *Mem_addr_md_icContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_mb_ia.
	VisitMem_addr_mb_ia(ctx *Mem_addr_mb_iaContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_ia_int.
	VisitMem_addr_ia_int(ctx *Mem_addr_ia_intContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_ic_int.
	VisitMem_addr_ic_int(ctx *Mem_addr_ic_intContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_int_ia.
	VisitMem_addr_int_ia(ctx *Mem_addr_int_iaContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_int_ic.
	VisitMem_addr_int_ic(ctx *Mem_addr_int_icContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_int.
	VisitMem_addr_int(ctx *Mem_addr_intContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_int_.
	VisitMem_addr_int_(ctx *Mem_addr_int_Context) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_dm_ia_mb.
	VisitMem_addr_dm_ia_mb(ctx *Mem_addr_dm_ia_mbContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_pm_ic_md.
	VisitMem_addr_pm_ic_md(ctx *Mem_addr_pm_ic_mdContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_dm_mb_ia.
	VisitMem_addr_dm_mb_ia(ctx *Mem_addr_dm_mb_iaContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_pm_md_ic.
	VisitMem_addr_pm_md_ic(ctx *Mem_addr_pm_md_icContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_dm_ia_int.
	VisitMem_addr_dm_ia_int(ctx *Mem_addr_dm_ia_intContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_pm_ic_int.
	VisitMem_addr_pm_ic_int(ctx *Mem_addr_pm_ic_intContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_dm_int_ia.
	VisitMem_addr_dm_int_ia(ctx *Mem_addr_dm_int_iaContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_pm_int_ic.
	VisitMem_addr_pm_int_ic(ctx *Mem_addr_pm_int_icContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_dm_int.
	VisitMem_addr_dm_int(ctx *Mem_addr_dm_intContext) interface{}

	// Visit a parse tree produced by SHARCParser#mem_addr_pm_int.
	VisitMem_addr_pm_int(ctx *Mem_addr_pm_intContext) interface{}

	// Visit a parse tree produced by SHARCParser#fixpoint_alu_op.
	VisitFixpoint_alu_op(ctx *Fixpoint_alu_opContext) interface{}

	// Visit a parse tree produced by SHARCParser#r_exp.
	VisitR_exp(ctx *R_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#floating_point_alu_op.
	VisitFloating_point_alu_op(ctx *Floating_point_alu_opContext) interface{}

	// Visit a parse tree produced by SHARCParser#f_exp.
	VisitF_exp(ctx *F_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#multi_op.
	VisitMulti_op(ctx *Multi_opContext) interface{}

	// Visit a parse tree produced by SHARCParser#multi_r.
	VisitMulti_r(ctx *Multi_rContext) interface{}

	// Visit a parse tree produced by SHARCParser#multi_exp_r.
	VisitMulti_exp_r(ctx *Multi_exp_rContext) interface{}

	// Visit a parse tree produced by SHARCParser#multi_exp_mrf.
	VisitMulti_exp_mrf(ctx *Multi_exp_mrfContext) interface{}

	// Visit a parse tree produced by SHARCParser#multi_exp_mrb.
	VisitMulti_exp_mrb(ctx *Multi_exp_mrbContext) interface{}

	// Visit a parse tree produced by SHARCParser#mr.
	VisitMr(ctx *MrContext) interface{}

	// Visit a parse tree produced by SHARCParser#shifter_op.
	VisitShifter_op(ctx *Shifter_opContext) interface{}

	// Visit a parse tree produced by SHARCParser#shifter_exp.
	VisitShifter_exp(ctx *Shifter_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#sec_op.
	VisitSec_op(ctx *Sec_opContext) interface{}

	// Visit a parse tree produced by SHARCParser#sec_op2.
	VisitSec_op2(ctx *Sec_op2Context) interface{}

	// Visit a parse tree produced by SHARCParser#bit_data.
	VisitBit_data(ctx *Bit_dataContext) interface{}

	// Visit a parse tree produced by SHARCParser#add_or_sub.
	VisitAdd_or_sub(ctx *Add_or_subContext) interface{}

	// Visit a parse tree produced by SHARCParser#dual_op.
	VisitDual_op(ctx *Dual_opContext) interface{}

	// Visit a parse tree produced by SHARCParser#dual_add_r.
	VisitDual_add_r(ctx *Dual_add_rContext) interface{}

	// Visit a parse tree produced by SHARCParser#parallel_multi.
	VisitParallel_multi(ctx *Parallel_multiContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_exp.
	VisitFlow_control_exp(ctx *Flow_control_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_contorl_8.
	VisitFlow_contorl_8(ctx *Flow_contorl_8Context) interface{}

	// Visit a parse tree produced by SHARCParser#flow_contorl_8_exp.
	VisitFlow_contorl_8_exp(ctx *Flow_contorl_8_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_9_and_11.
	VisitFlow_control_9_and_11(ctx *Flow_control_9_and_11Context) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_9_and_11_exp.
	VisitFlow_control_9_and_11_exp(ctx *Flow_control_9_and_11_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_10.
	VisitFlow_control_10(ctx *Flow_control_10Context) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_10_frag.
	VisitFlow_control_10_frag(ctx *Flow_control_10_fragContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_12.
	VisitFlow_control_12(ctx *Flow_control_12Context) interface{}

	// Visit a parse tree produced by SHARCParser#lcntr_v.
	VisitLcntr_v(ctx *Lcntr_vContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_13.
	VisitFlow_control_13(ctx *Flow_control_13Context) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_8a.
	VisitFlow_control_8a(ctx *Flow_control_8aContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_8b.
	VisitFlow_control_8b(ctx *Flow_control_8bContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_9a.
	VisitFlow_control_9a(ctx *Flow_control_9aContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_9b.
	VisitFlow_control_9b(ctx *Flow_control_9bContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_11a.
	VisitFlow_control_11a(ctx *Flow_control_11aContext) interface{}

	// Visit a parse tree produced by SHARCParser#flow_control_11b.
	VisitFlow_control_11b(ctx *Flow_control_11bContext) interface{}

	// Visit a parse tree produced by SHARCParser#jump_addr_int_or_pc.
	VisitJump_addr_int_or_pc(ctx *Jump_addr_int_or_pcContext) interface{}

	// Visit a parse tree produced by SHARCParser#jump_addr_md_or_pc.
	VisitJump_addr_md_or_pc(ctx *Jump_addr_md_or_pcContext) interface{}

	// Visit a parse tree produced by SHARCParser#jump_addr_pc.
	VisitJump_addr_pc(ctx *Jump_addr_pcContext) interface{}

	// Visit a parse tree produced by SHARCParser#jump_addr_int.
	VisitJump_addr_int(ctx *Jump_addr_intContext) interface{}

	// Visit a parse tree produced by SHARCParser#jump_modifier.
	VisitJump_modifier(ctx *Jump_modifierContext) interface{}

	// Visit a parse tree produced by SHARCParser#jump_modifier_.
	VisitJump_modifier_(ctx *Jump_modifier_Context) interface{}

	// Visit a parse tree produced by SHARCParser#jump_modifier_1.
	VisitJump_modifier_1(ctx *Jump_modifier_1Context) interface{}

	// Visit a parse tree produced by SHARCParser#jump_modifier2.
	VisitJump_modifier2(ctx *Jump_modifier2Context) interface{}

	// Visit a parse tree produced by SHARCParser#jump_modifier3.
	VisitJump_modifier3(ctx *Jump_modifier3Context) interface{}

	// Visit a parse tree produced by SHARCParser#jump_modifier3_.
	VisitJump_modifier3_(ctx *Jump_modifier3_Context) interface{}

	// Visit a parse tree produced by SHARCParser#jump_modifier3_1.
	VisitJump_modifier3_1(ctx *Jump_modifier3_1Context) interface{}

	// Visit a parse tree produced by SHARCParser#imm_mov_exp.
	VisitImm_mov_exp(ctx *Imm_mov_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#imm_mov_14a.
	VisitImm_mov_14a(ctx *Imm_mov_14aContext) interface{}

	// Visit a parse tree produced by SHARCParser#imm_mov_15a.
	VisitImm_mov_15a(ctx *Imm_mov_15aContext) interface{}

	// Visit a parse tree produced by SHARCParser#imm_mov_14b.
	VisitImm_mov_14b(ctx *Imm_mov_14bContext) interface{}

	// Visit a parse tree produced by SHARCParser#imm_mov_15b.
	VisitImm_mov_15b(ctx *Imm_mov_15bContext) interface{}

	// Visit a parse tree produced by SHARCParser#imm_mov_16.
	VisitImm_mov_16(ctx *Imm_mov_16Context) interface{}

	// Visit a parse tree produced by SHARCParser#imm_mov_17.
	VisitImm_mov_17(ctx *Imm_mov_17Context) interface{}

	// Visit a parse tree produced by SHARCParser#u_reg2.
	VisitU_reg2(ctx *U_reg2Context) interface{}

	// Visit a parse tree produced by SHARCParser#misc_exp.
	VisitMisc_exp(ctx *Misc_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#misc_20.
	VisitMisc_20(ctx *Misc_20Context) interface{}

	// Visit a parse tree produced by SHARCParser#directive_exp.
	VisitDirective_exp(ctx *Directive_expContext) interface{}

	// Visit a parse tree produced by SHARCParser#b_reg.
	VisitB_reg(ctx *B_regContext) interface{}

	// Visit a parse tree produced by SHARCParser#l_reg.
	VisitL_reg(ctx *L_regContext) interface{}

	// Visit a parse tree produced by SHARCParser#r_reg.
	VisitR_reg(ctx *R_regContext) interface{}

	// Visit a parse tree produced by SHARCParser#f_reg.
	VisitF_reg(ctx *F_regContext) interface{}

	// Visit a parse tree produced by SHARCParser#s_reg.
	VisitS_reg(ctx *S_regContext) interface{}

	// Visit a parse tree produced by SHARCParser#ia.
	VisitIa(ctx *IaContext) interface{}

	// Visit a parse tree produced by SHARCParser#mb.
	VisitMb(ctx *MbContext) interface{}

	// Visit a parse tree produced by SHARCParser#ic.
	VisitIc(ctx *IcContext) interface{}

	// Visit a parse tree produced by SHARCParser#md.
	VisitMd(ctx *MdContext) interface{}

	// Visit a parse tree produced by SHARCParser#i_reg.
	VisitI_reg(ctx *I_regContext) interface{}

	// Visit a parse tree produced by SHARCParser#m_reg.
	VisitM_reg(ctx *M_regContext) interface{}

	// Visit a parse tree produced by SHARCParser#dag_reg.
	VisitDag_reg(ctx *Dag_regContext) interface{}

	// Visit a parse tree produced by SHARCParser#d_reg.
	VisitD_reg(ctx *D_regContext) interface{}

	// Visit a parse tree produced by SHARCParser#u_reg.
	VisitU_reg(ctx *U_regContext) interface{}

	// Visit a parse tree produced by SHARCParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by SHARCParser#ccondition.
	VisitCcondition(ctx *CconditionContext) interface{}

	// Visit a parse tree produced by SHARCParser#multi_mod1.
	VisitMulti_mod1(ctx *Multi_mod1Context) interface{}

	// Visit a parse tree produced by SHARCParser#multi_mod1_.
	VisitMulti_mod1_(ctx *Multi_mod1_Context) interface{}

	// Visit a parse tree produced by SHARCParser#multi_mod2.
	VisitMulti_mod2(ctx *Multi_mod2Context) interface{}

	// Visit a parse tree produced by SHARCParser#multi_mod2_.
	VisitMulti_mod2_(ctx *Multi_mod2_Context) interface{}

	// Visit a parse tree produced by SHARCParser#r3_0.
	VisitR3_0(ctx *R3_0Context) interface{}

	// Visit a parse tree produced by SHARCParser#r7_4.
	VisitR7_4(ctx *R7_4Context) interface{}

	// Visit a parse tree produced by SHARCParser#r11_8.
	VisitR11_8(ctx *R11_8Context) interface{}

	// Visit a parse tree produced by SHARCParser#r15_12.
	VisitR15_12(ctx *R15_12Context) interface{}

	// Visit a parse tree produced by SHARCParser#f3_0.
	VisitF3_0(ctx *F3_0Context) interface{}

	// Visit a parse tree produced by SHARCParser#f7_4.
	VisitF7_4(ctx *F7_4Context) interface{}

	// Visit a parse tree produced by SHARCParser#f11_8.
	VisitF11_8(ctx *F11_8Context) interface{}

	// Visit a parse tree produced by SHARCParser#f15_12.
	VisitF15_12(ctx *F15_12Context) interface{}

	// Visit a parse tree produced by SHARCParser#addr.
	VisitAddr(ctx *AddrContext) interface{}

	// Visit a parse tree produced by SHARCParser#mrf.
	VisitMrf(ctx *MrfContext) interface{}

	// Visit a parse tree produced by SHARCParser#mrb.
	VisitMrb(ctx *MrbContext) interface{}
}
