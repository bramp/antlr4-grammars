// Code generated from argus.g4 by ANTLR 4.9.3. DO NOT EDIT.

package argus // argus
import "github.com/antlr/antlr4/runtime/Go/antlr"

// argusListener is a complete listener for a parse tree produced by argusParser.
type argusListener interface {
	antlr.ParseTreeListener

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterEquates_ is called when entering the equates_ production.
	EnterEquates_(c *Equates_Context)

	// EnterGuardian is called when entering the guardian production.
	EnterGuardian(c *GuardianContext)

	// EnterCluster is called when entering the cluster production.
	EnterCluster(c *ClusterContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterRoutine is called when entering the routine production.
	EnterRoutine(c *RoutineContext)

	// EnterProcedure is called when entering the procedure production.
	EnterProcedure(c *ProcedureContext)

	// EnterIterator is called when entering the iterator production.
	EnterIterator(c *IteratorContext)

	// EnterCreator is called when entering the creator production.
	EnterCreator(c *CreatorContext)

	// EnterHandler is called when entering the handler production.
	EnterHandler(c *HandlerContext)

	// EnterRoutine_body is called when entering the routine_body production.
	EnterRoutine_body(c *Routine_bodyContext)

	// EnterParms is called when entering the parms production.
	EnterParms(c *ParmsContext)

	// EnterParm is called when entering the parm production.
	EnterParm(c *ParmContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterReturnz is called when entering the returnz production.
	EnterReturnz(c *ReturnzContext)

	// EnterYields is called when entering the yields production.
	EnterYields(c *YieldsContext)

	// EnterSignals is called when entering the signals production.
	EnterSignals(c *SignalsContext)

	// EnterException_ is called when entering the exception_ production.
	EnterException_(c *Exception_Context)

	// EnterOpidn is called when entering the opidn production.
	EnterOpidn(c *OpidnContext)

	// EnterWhere is called when entering the where production.
	EnterWhere(c *WhereContext)

	// EnterRestriction is called when entering the restriction production.
	EnterRestriction(c *RestrictionContext)

	// EnterType_set is called when entering the type_set production.
	EnterType_set(c *Type_setContext)

	// EnterOper_decl is called when entering the oper_decl production.
	EnterOper_decl(c *Oper_declContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterState_decl is called when entering the state_decl production.
	EnterState_decl(c *State_declContext)

	// EnterEquate is called when entering the equate production.
	EnterEquate(c *EquateContext)

	// EnterOwn_var is called when entering the own_var production.
	EnterOwn_var(c *Own_varContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterEnter_stmt is called when entering the enter_stmt production.
	EnterEnter_stmt(c *Enter_stmtContext)

	// EnterCoarm is called when entering the coarm production.
	EnterCoarm(c *CoarmContext)

	// EnterArmtag is called when entering the armtag production.
	EnterArmtag(c *ArmtagContext)

	// EnterFor_stmt is called when entering the for_stmt production.
	EnterFor_stmt(c *For_stmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterTagcase_stmt is called when entering the tagcase_stmt production.
	EnterTagcase_stmt(c *Tagcase_stmtContext)

	// EnterTagtest_stmt is called when entering the tagtest_stmt production.
	EnterTagtest_stmt(c *Tagtest_stmtContext)

	// EnterTagwait_stmt is called when entering the tagwait_stmt production.
	EnterTagwait_stmt(c *Tagwait_stmtContext)

	// EnterTag_arm is called when entering the tag_arm production.
	EnterTag_arm(c *Tag_armContext)

	// EnterAtag_arm is called when entering the atag_arm production.
	EnterAtag_arm(c *Atag_armContext)

	// EnterTag_kind is called when entering the tag_kind production.
	EnterTag_kind(c *Tag_kindContext)

	// EnterWhen_handler is called when entering the when_handler production.
	EnterWhen_handler(c *When_handlerContext)

	// EnterOthers_handler is called when entering the others_handler production.
	EnterOthers_handler(c *Others_handlerContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterType_spec is called when entering the type_spec production.
	EnterType_spec(c *Type_specContext)

	// EnterField_spec is called when entering the field_spec production.
	EnterField_spec(c *Field_specContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterActual_parm is called when entering the actual_parm production.
	EnterActual_parm(c *Actual_parmContext)

	// EnterType_actual is called when entering the type_actual production.
	EnterType_actual(c *Type_actualContext)

	// EnterOpbinding is called when entering the opbinding production.
	EnterOpbinding(c *OpbindingContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimaries is called when entering the primaries production.
	EnterPrimaries(c *PrimariesContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterEntities is called when entering the entities production.
	EnterEntities(c *EntitiesContext)

	// EnterEntity is called when entering the entity production.
	EnterEntity(c *EntityContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterBind_arg is called when entering the bind_arg production.
	EnterBind_arg(c *Bind_argContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterIdn is called when entering the idn production.
	EnterIdn(c *IdnContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitEquates_ is called when exiting the equates_ production.
	ExitEquates_(c *Equates_Context)

	// ExitGuardian is called when exiting the guardian production.
	ExitGuardian(c *GuardianContext)

	// ExitCluster is called when exiting the cluster production.
	ExitCluster(c *ClusterContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitRoutine is called when exiting the routine production.
	ExitRoutine(c *RoutineContext)

	// ExitProcedure is called when exiting the procedure production.
	ExitProcedure(c *ProcedureContext)

	// ExitIterator is called when exiting the iterator production.
	ExitIterator(c *IteratorContext)

	// ExitCreator is called when exiting the creator production.
	ExitCreator(c *CreatorContext)

	// ExitHandler is called when exiting the handler production.
	ExitHandler(c *HandlerContext)

	// ExitRoutine_body is called when exiting the routine_body production.
	ExitRoutine_body(c *Routine_bodyContext)

	// ExitParms is called when exiting the parms production.
	ExitParms(c *ParmsContext)

	// ExitParm is called when exiting the parm production.
	ExitParm(c *ParmContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitReturnz is called when exiting the returnz production.
	ExitReturnz(c *ReturnzContext)

	// ExitYields is called when exiting the yields production.
	ExitYields(c *YieldsContext)

	// ExitSignals is called when exiting the signals production.
	ExitSignals(c *SignalsContext)

	// ExitException_ is called when exiting the exception_ production.
	ExitException_(c *Exception_Context)

	// ExitOpidn is called when exiting the opidn production.
	ExitOpidn(c *OpidnContext)

	// ExitWhere is called when exiting the where production.
	ExitWhere(c *WhereContext)

	// ExitRestriction is called when exiting the restriction production.
	ExitRestriction(c *RestrictionContext)

	// ExitType_set is called when exiting the type_set production.
	ExitType_set(c *Type_setContext)

	// ExitOper_decl is called when exiting the oper_decl production.
	ExitOper_decl(c *Oper_declContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitState_decl is called when exiting the state_decl production.
	ExitState_decl(c *State_declContext)

	// ExitEquate is called when exiting the equate production.
	ExitEquate(c *EquateContext)

	// ExitOwn_var is called when exiting the own_var production.
	ExitOwn_var(c *Own_varContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitEnter_stmt is called when exiting the enter_stmt production.
	ExitEnter_stmt(c *Enter_stmtContext)

	// ExitCoarm is called when exiting the coarm production.
	ExitCoarm(c *CoarmContext)

	// ExitArmtag is called when exiting the armtag production.
	ExitArmtag(c *ArmtagContext)

	// ExitFor_stmt is called when exiting the for_stmt production.
	ExitFor_stmt(c *For_stmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitTagcase_stmt is called when exiting the tagcase_stmt production.
	ExitTagcase_stmt(c *Tagcase_stmtContext)

	// ExitTagtest_stmt is called when exiting the tagtest_stmt production.
	ExitTagtest_stmt(c *Tagtest_stmtContext)

	// ExitTagwait_stmt is called when exiting the tagwait_stmt production.
	ExitTagwait_stmt(c *Tagwait_stmtContext)

	// ExitTag_arm is called when exiting the tag_arm production.
	ExitTag_arm(c *Tag_armContext)

	// ExitAtag_arm is called when exiting the atag_arm production.
	ExitAtag_arm(c *Atag_armContext)

	// ExitTag_kind is called when exiting the tag_kind production.
	ExitTag_kind(c *Tag_kindContext)

	// ExitWhen_handler is called when exiting the when_handler production.
	ExitWhen_handler(c *When_handlerContext)

	// ExitOthers_handler is called when exiting the others_handler production.
	ExitOthers_handler(c *Others_handlerContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitType_spec is called when exiting the type_spec production.
	ExitType_spec(c *Type_specContext)

	// ExitField_spec is called when exiting the field_spec production.
	ExitField_spec(c *Field_specContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitActual_parm is called when exiting the actual_parm production.
	ExitActual_parm(c *Actual_parmContext)

	// ExitType_actual is called when exiting the type_actual production.
	ExitType_actual(c *Type_actualContext)

	// ExitOpbinding is called when exiting the opbinding production.
	ExitOpbinding(c *OpbindingContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimaries is called when exiting the primaries production.
	ExitPrimaries(c *PrimariesContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitEntities is called when exiting the entities production.
	ExitEntities(c *EntitiesContext)

	// ExitEntity is called when exiting the entity production.
	ExitEntity(c *EntityContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitBind_arg is called when exiting the bind_arg production.
	ExitBind_arg(c *Bind_argContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitIdn is called when exiting the idn production.
	ExitIdn(c *IdnContext)
}
