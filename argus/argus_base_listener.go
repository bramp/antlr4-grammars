// Code generated from argus.g4 by ANTLR 4.9.3. DO NOT EDIT.

package argus // argus
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseargusListener is a complete listener for a parse tree produced by argusParser.
type BaseargusListener struct{}

var _ argusListener = &BaseargusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseargusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseargusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseargusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseargusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModule is called when production module is entered.
func (s *BaseargusListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseargusListener) ExitModule(ctx *ModuleContext) {}

// EnterEquates_ is called when production equates_ is entered.
func (s *BaseargusListener) EnterEquates_(ctx *Equates_Context) {}

// ExitEquates_ is called when production equates_ is exited.
func (s *BaseargusListener) ExitEquates_(ctx *Equates_Context) {}

// EnterGuardian is called when production guardian is entered.
func (s *BaseargusListener) EnterGuardian(ctx *GuardianContext) {}

// ExitGuardian is called when production guardian is exited.
func (s *BaseargusListener) ExitGuardian(ctx *GuardianContext) {}

// EnterCluster is called when production cluster is entered.
func (s *BaseargusListener) EnterCluster(ctx *ClusterContext) {}

// ExitCluster is called when production cluster is exited.
func (s *BaseargusListener) ExitCluster(ctx *ClusterContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseargusListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseargusListener) ExitOperation(ctx *OperationContext) {}

// EnterRoutine is called when production routine is entered.
func (s *BaseargusListener) EnterRoutine(ctx *RoutineContext) {}

// ExitRoutine is called when production routine is exited.
func (s *BaseargusListener) ExitRoutine(ctx *RoutineContext) {}

// EnterProcedure is called when production procedure is entered.
func (s *BaseargusListener) EnterProcedure(ctx *ProcedureContext) {}

// ExitProcedure is called when production procedure is exited.
func (s *BaseargusListener) ExitProcedure(ctx *ProcedureContext) {}

// EnterIterator is called when production iterator is entered.
func (s *BaseargusListener) EnterIterator(ctx *IteratorContext) {}

// ExitIterator is called when production iterator is exited.
func (s *BaseargusListener) ExitIterator(ctx *IteratorContext) {}

// EnterCreator is called when production creator is entered.
func (s *BaseargusListener) EnterCreator(ctx *CreatorContext) {}

// ExitCreator is called when production creator is exited.
func (s *BaseargusListener) ExitCreator(ctx *CreatorContext) {}

// EnterHandler is called when production handler is entered.
func (s *BaseargusListener) EnterHandler(ctx *HandlerContext) {}

// ExitHandler is called when production handler is exited.
func (s *BaseargusListener) ExitHandler(ctx *HandlerContext) {}

// EnterRoutine_body is called when production routine_body is entered.
func (s *BaseargusListener) EnterRoutine_body(ctx *Routine_bodyContext) {}

// ExitRoutine_body is called when production routine_body is exited.
func (s *BaseargusListener) ExitRoutine_body(ctx *Routine_bodyContext) {}

// EnterParms is called when production parms is entered.
func (s *BaseargusListener) EnterParms(ctx *ParmsContext) {}

// ExitParms is called when production parms is exited.
func (s *BaseargusListener) ExitParms(ctx *ParmsContext) {}

// EnterParm is called when production parm is entered.
func (s *BaseargusListener) EnterParm(ctx *ParmContext) {}

// ExitParm is called when production parm is exited.
func (s *BaseargusListener) ExitParm(ctx *ParmContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseargusListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseargusListener) ExitArgs(ctx *ArgsContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseargusListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseargusListener) ExitDecl(ctx *DeclContext) {}

// EnterReturnz is called when production returnz is entered.
func (s *BaseargusListener) EnterReturnz(ctx *ReturnzContext) {}

// ExitReturnz is called when production returnz is exited.
func (s *BaseargusListener) ExitReturnz(ctx *ReturnzContext) {}

// EnterYields is called when production yields is entered.
func (s *BaseargusListener) EnterYields(ctx *YieldsContext) {}

// ExitYields is called when production yields is exited.
func (s *BaseargusListener) ExitYields(ctx *YieldsContext) {}

// EnterSignals is called when production signals is entered.
func (s *BaseargusListener) EnterSignals(ctx *SignalsContext) {}

// ExitSignals is called when production signals is exited.
func (s *BaseargusListener) ExitSignals(ctx *SignalsContext) {}

// EnterException_ is called when production exception_ is entered.
func (s *BaseargusListener) EnterException_(ctx *Exception_Context) {}

// ExitException_ is called when production exception_ is exited.
func (s *BaseargusListener) ExitException_(ctx *Exception_Context) {}

// EnterOpidn is called when production opidn is entered.
func (s *BaseargusListener) EnterOpidn(ctx *OpidnContext) {}

// ExitOpidn is called when production opidn is exited.
func (s *BaseargusListener) ExitOpidn(ctx *OpidnContext) {}

// EnterWhere is called when production where is entered.
func (s *BaseargusListener) EnterWhere(ctx *WhereContext) {}

// ExitWhere is called when production where is exited.
func (s *BaseargusListener) ExitWhere(ctx *WhereContext) {}

// EnterRestriction is called when production restriction is entered.
func (s *BaseargusListener) EnterRestriction(ctx *RestrictionContext) {}

// ExitRestriction is called when production restriction is exited.
func (s *BaseargusListener) ExitRestriction(ctx *RestrictionContext) {}

// EnterType_set is called when production type_set is entered.
func (s *BaseargusListener) EnterType_set(ctx *Type_setContext) {}

// ExitType_set is called when production type_set is exited.
func (s *BaseargusListener) ExitType_set(ctx *Type_setContext) {}

// EnterOper_decl is called when production oper_decl is entered.
func (s *BaseargusListener) EnterOper_decl(ctx *Oper_declContext) {}

// ExitOper_decl is called when production oper_decl is exited.
func (s *BaseargusListener) ExitOper_decl(ctx *Oper_declContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseargusListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseargusListener) ExitConstant(ctx *ConstantContext) {}

// EnterState_decl is called when production state_decl is entered.
func (s *BaseargusListener) EnterState_decl(ctx *State_declContext) {}

// ExitState_decl is called when production state_decl is exited.
func (s *BaseargusListener) ExitState_decl(ctx *State_declContext) {}

// EnterEquate is called when production equate is entered.
func (s *BaseargusListener) EnterEquate(ctx *EquateContext) {}

// ExitEquate is called when production equate is exited.
func (s *BaseargusListener) ExitEquate(ctx *EquateContext) {}

// EnterOwn_var is called when production own_var is entered.
func (s *BaseargusListener) EnterOwn_var(ctx *Own_varContext) {}

// ExitOwn_var is called when production own_var is exited.
func (s *BaseargusListener) ExitOwn_var(ctx *Own_varContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseargusListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseargusListener) ExitStatement(ctx *StatementContext) {}

// EnterEnter_stmt is called when production enter_stmt is entered.
func (s *BaseargusListener) EnterEnter_stmt(ctx *Enter_stmtContext) {}

// ExitEnter_stmt is called when production enter_stmt is exited.
func (s *BaseargusListener) ExitEnter_stmt(ctx *Enter_stmtContext) {}

// EnterCoarm is called when production coarm is entered.
func (s *BaseargusListener) EnterCoarm(ctx *CoarmContext) {}

// ExitCoarm is called when production coarm is exited.
func (s *BaseargusListener) ExitCoarm(ctx *CoarmContext) {}

// EnterArmtag is called when production armtag is entered.
func (s *BaseargusListener) EnterArmtag(ctx *ArmtagContext) {}

// ExitArmtag is called when production armtag is exited.
func (s *BaseargusListener) ExitArmtag(ctx *ArmtagContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BaseargusListener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BaseargusListener) ExitFor_stmt(ctx *For_stmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseargusListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseargusListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterTagcase_stmt is called when production tagcase_stmt is entered.
func (s *BaseargusListener) EnterTagcase_stmt(ctx *Tagcase_stmtContext) {}

// ExitTagcase_stmt is called when production tagcase_stmt is exited.
func (s *BaseargusListener) ExitTagcase_stmt(ctx *Tagcase_stmtContext) {}

// EnterTagtest_stmt is called when production tagtest_stmt is entered.
func (s *BaseargusListener) EnterTagtest_stmt(ctx *Tagtest_stmtContext) {}

// ExitTagtest_stmt is called when production tagtest_stmt is exited.
func (s *BaseargusListener) ExitTagtest_stmt(ctx *Tagtest_stmtContext) {}

// EnterTagwait_stmt is called when production tagwait_stmt is entered.
func (s *BaseargusListener) EnterTagwait_stmt(ctx *Tagwait_stmtContext) {}

// ExitTagwait_stmt is called when production tagwait_stmt is exited.
func (s *BaseargusListener) ExitTagwait_stmt(ctx *Tagwait_stmtContext) {}

// EnterTag_arm is called when production tag_arm is entered.
func (s *BaseargusListener) EnterTag_arm(ctx *Tag_armContext) {}

// ExitTag_arm is called when production tag_arm is exited.
func (s *BaseargusListener) ExitTag_arm(ctx *Tag_armContext) {}

// EnterAtag_arm is called when production atag_arm is entered.
func (s *BaseargusListener) EnterAtag_arm(ctx *Atag_armContext) {}

// ExitAtag_arm is called when production atag_arm is exited.
func (s *BaseargusListener) ExitAtag_arm(ctx *Atag_armContext) {}

// EnterTag_kind is called when production tag_kind is entered.
func (s *BaseargusListener) EnterTag_kind(ctx *Tag_kindContext) {}

// ExitTag_kind is called when production tag_kind is exited.
func (s *BaseargusListener) ExitTag_kind(ctx *Tag_kindContext) {}

// EnterWhen_handler is called when production when_handler is entered.
func (s *BaseargusListener) EnterWhen_handler(ctx *When_handlerContext) {}

// ExitWhen_handler is called when production when_handler is exited.
func (s *BaseargusListener) ExitWhen_handler(ctx *When_handlerContext) {}

// EnterOthers_handler is called when production others_handler is entered.
func (s *BaseargusListener) EnterOthers_handler(ctx *Others_handlerContext) {}

// ExitOthers_handler is called when production others_handler is exited.
func (s *BaseargusListener) ExitOthers_handler(ctx *Others_handlerContext) {}

// EnterBody is called when production body is entered.
func (s *BaseargusListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseargusListener) ExitBody(ctx *BodyContext) {}

// EnterType_spec is called when production type_spec is entered.
func (s *BaseargusListener) EnterType_spec(ctx *Type_specContext) {}

// ExitType_spec is called when production type_spec is exited.
func (s *BaseargusListener) ExitType_spec(ctx *Type_specContext) {}

// EnterField_spec is called when production field_spec is entered.
func (s *BaseargusListener) EnterField_spec(ctx *Field_specContext) {}

// ExitField_spec is called when production field_spec is exited.
func (s *BaseargusListener) ExitField_spec(ctx *Field_specContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseargusListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseargusListener) ExitReference(ctx *ReferenceContext) {}

// EnterActual_parm is called when production actual_parm is entered.
func (s *BaseargusListener) EnterActual_parm(ctx *Actual_parmContext) {}

// ExitActual_parm is called when production actual_parm is exited.
func (s *BaseargusListener) ExitActual_parm(ctx *Actual_parmContext) {}

// EnterType_actual is called when production type_actual is entered.
func (s *BaseargusListener) EnterType_actual(ctx *Type_actualContext) {}

// ExitType_actual is called when production type_actual is exited.
func (s *BaseargusListener) ExitType_actual(ctx *Type_actualContext) {}

// EnterOpbinding is called when production opbinding is entered.
func (s *BaseargusListener) EnterOpbinding(ctx *OpbindingContext) {}

// ExitOpbinding is called when production opbinding is exited.
func (s *BaseargusListener) ExitOpbinding(ctx *OpbindingContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseargusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseargusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaries is called when production primaries is entered.
func (s *BaseargusListener) EnterPrimaries(ctx *PrimariesContext) {}

// ExitPrimaries is called when production primaries is exited.
func (s *BaseargusListener) ExitPrimaries(ctx *PrimariesContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseargusListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseargusListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterCall is called when production call is entered.
func (s *BaseargusListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseargusListener) ExitCall(ctx *CallContext) {}

// EnterEntities is called when production entities is entered.
func (s *BaseargusListener) EnterEntities(ctx *EntitiesContext) {}

// ExitEntities is called when production entities is exited.
func (s *BaseargusListener) ExitEntities(ctx *EntitiesContext) {}

// EnterEntity is called when production entity is entered.
func (s *BaseargusListener) EnterEntity(ctx *EntityContext) {}

// ExitEntity is called when production entity is exited.
func (s *BaseargusListener) ExitEntity(ctx *EntityContext) {}

// EnterField is called when production field is entered.
func (s *BaseargusListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseargusListener) ExitField(ctx *FieldContext) {}

// EnterBind_arg is called when production bind_arg is entered.
func (s *BaseargusListener) EnterBind_arg(ctx *Bind_argContext) {}

// ExitBind_arg is called when production bind_arg is exited.
func (s *BaseargusListener) ExitBind_arg(ctx *Bind_argContext) {}

// EnterName is called when production name is entered.
func (s *BaseargusListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseargusListener) ExitName(ctx *NameContext) {}

// EnterIdn is called when production idn is entered.
func (s *BaseargusListener) EnterIdn(ctx *IdnContext) {}

// ExitIdn is called when production idn is exited.
func (s *BaseargusListener) ExitIdn(ctx *IdnContext) {}
