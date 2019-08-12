// Code generated from clu.g4 by ANTLR 4.7.2. DO NOT EDIT.

package clu // clu
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasecluListener is a complete listener for a parse tree produced by cluParser.
type BasecluListener struct{}

var _ cluListener = &BasecluListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecluListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecluListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecluListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecluListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModule is called when production module is entered.
func (s *BasecluListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BasecluListener) ExitModule(ctx *ModuleContext) {}

// EnterProcedure is called when production procedure is entered.
func (s *BasecluListener) EnterProcedure(ctx *ProcedureContext) {}

// ExitProcedure is called when production procedure is exited.
func (s *BasecluListener) ExitProcedure(ctx *ProcedureContext) {}

// EnterIterator is called when production iterator is entered.
func (s *BasecluListener) EnterIterator(ctx *IteratorContext) {}

// ExitIterator is called when production iterator is exited.
func (s *BasecluListener) ExitIterator(ctx *IteratorContext) {}

// EnterCluster is called when production cluster is entered.
func (s *BasecluListener) EnterCluster(ctx *ClusterContext) {}

// ExitCluster is called when production cluster is exited.
func (s *BasecluListener) ExitCluster(ctx *ClusterContext) {}

// EnterParms is called when production parms is entered.
func (s *BasecluListener) EnterParms(ctx *ParmsContext) {}

// ExitParms is called when production parms is exited.
func (s *BasecluListener) ExitParms(ctx *ParmsContext) {}

// EnterParam is called when production param is entered.
func (s *BasecluListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasecluListener) ExitParam(ctx *ParamContext) {}

// EnterArgs is called when production args is entered.
func (s *BasecluListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BasecluListener) ExitArgs(ctx *ArgsContext) {}

// EnterDecl_list is called when production decl_list is entered.
func (s *BasecluListener) EnterDecl_list(ctx *Decl_listContext) {}

// ExitDecl_list is called when production decl_list is exited.
func (s *BasecluListener) ExitDecl_list(ctx *Decl_listContext) {}

// EnterDecl is called when production decl is entered.
func (s *BasecluListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BasecluListener) ExitDecl(ctx *DeclContext) {}

// EnterReturnz is called when production returnz is entered.
func (s *BasecluListener) EnterReturnz(ctx *ReturnzContext) {}

// ExitReturnz is called when production returnz is exited.
func (s *BasecluListener) ExitReturnz(ctx *ReturnzContext) {}

// EnterYields is called when production yields is entered.
func (s *BasecluListener) EnterYields(ctx *YieldsContext) {}

// ExitYields is called when production yields is exited.
func (s *BasecluListener) ExitYields(ctx *YieldsContext) {}

// EnterSignals is called when production signals is entered.
func (s *BasecluListener) EnterSignals(ctx *SignalsContext) {}

// ExitSignals is called when production signals is exited.
func (s *BasecluListener) ExitSignals(ctx *SignalsContext) {}

// EnterException is called when production exception is entered.
func (s *BasecluListener) EnterException(ctx *ExceptionContext) {}

// ExitException is called when production exception is exited.
func (s *BasecluListener) ExitException(ctx *ExceptionContext) {}

// EnterType_spec_list is called when production type_spec_list is entered.
func (s *BasecluListener) EnterType_spec_list(ctx *Type_spec_listContext) {}

// ExitType_spec_list is called when production type_spec_list is exited.
func (s *BasecluListener) ExitType_spec_list(ctx *Type_spec_listContext) {}

// EnterWhere is called when production where is entered.
func (s *BasecluListener) EnterWhere(ctx *WhereContext) {}

// ExitWhere is called when production where is exited.
func (s *BasecluListener) ExitWhere(ctx *WhereContext) {}

// EnterRestriction is called when production restriction is entered.
func (s *BasecluListener) EnterRestriction(ctx *RestrictionContext) {}

// ExitRestriction is called when production restriction is exited.
func (s *BasecluListener) ExitRestriction(ctx *RestrictionContext) {}

// EnterType_set is called when production type_set is entered.
func (s *BasecluListener) EnterType_set(ctx *Type_setContext) {}

// ExitType_set is called when production type_set is exited.
func (s *BasecluListener) ExitType_set(ctx *Type_setContext) {}

// EnterOper_decl_list is called when production oper_decl_list is entered.
func (s *BasecluListener) EnterOper_decl_list(ctx *Oper_decl_listContext) {}

// ExitOper_decl_list is called when production oper_decl_list is exited.
func (s *BasecluListener) ExitOper_decl_list(ctx *Oper_decl_listContext) {}

// EnterOper_decl is called when production oper_decl is entered.
func (s *BasecluListener) EnterOper_decl(ctx *Oper_declContext) {}

// ExitOper_decl is called when production oper_decl is exited.
func (s *BasecluListener) ExitOper_decl(ctx *Oper_declContext) {}

// EnterOp_name_list is called when production op_name_list is entered.
func (s *BasecluListener) EnterOp_name_list(ctx *Op_name_listContext) {}

// ExitOp_name_list is called when production op_name_list is exited.
func (s *BasecluListener) ExitOp_name_list(ctx *Op_name_listContext) {}

// EnterOp_name is called when production op_name is entered.
func (s *BasecluListener) EnterOp_name(ctx *Op_nameContext) {}

// ExitOp_name is called when production op_name is exited.
func (s *BasecluListener) ExitOp_name(ctx *Op_nameContext) {}

// EnterConstant_list is called when production constant_list is entered.
func (s *BasecluListener) EnterConstant_list(ctx *Constant_listContext) {}

// ExitConstant_list is called when production constant_list is exited.
func (s *BasecluListener) ExitConstant_list(ctx *Constant_listContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasecluListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasecluListener) ExitConstant(ctx *ConstantContext) {}

// EnterRoutine_body is called when production routine_body is entered.
func (s *BasecluListener) EnterRoutine_body(ctx *Routine_bodyContext) {}

// ExitRoutine_body is called when production routine_body is exited.
func (s *BasecluListener) ExitRoutine_body(ctx *Routine_bodyContext) {}

// EnterCluster_body is called when production cluster_body is entered.
func (s *BasecluListener) EnterCluster_body(ctx *Cluster_bodyContext) {}

// ExitCluster_body is called when production cluster_body is exited.
func (s *BasecluListener) ExitCluster_body(ctx *Cluster_bodyContext) {}

// EnterRoutine is called when production routine is entered.
func (s *BasecluListener) EnterRoutine(ctx *RoutineContext) {}

// ExitRoutine is called when production routine is exited.
func (s *BasecluListener) ExitRoutine(ctx *RoutineContext) {}

// EnterEquate is called when production equate is entered.
func (s *BasecluListener) EnterEquate(ctx *EquateContext) {}

// ExitEquate is called when production equate is exited.
func (s *BasecluListener) ExitEquate(ctx *EquateContext) {}

// EnterOwn_var is called when production own_var is entered.
func (s *BasecluListener) EnterOwn_var(ctx *Own_varContext) {}

// ExitOwn_var is called when production own_var is exited.
func (s *BasecluListener) ExitOwn_var(ctx *Own_varContext) {}

// EnterType_spec is called when production type_spec is entered.
func (s *BasecluListener) EnterType_spec(ctx *Type_specContext) {}

// ExitType_spec is called when production type_spec is exited.
func (s *BasecluListener) ExitType_spec(ctx *Type_specContext) {}

// EnterField_spec_list is called when production field_spec_list is entered.
func (s *BasecluListener) EnterField_spec_list(ctx *Field_spec_listContext) {}

// ExitField_spec_list is called when production field_spec_list is exited.
func (s *BasecluListener) ExitField_spec_list(ctx *Field_spec_listContext) {}

// EnterField_spec is called when production field_spec is entered.
func (s *BasecluListener) EnterField_spec(ctx *Field_specContext) {}

// ExitField_spec is called when production field_spec is exited.
func (s *BasecluListener) ExitField_spec(ctx *Field_specContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasecluListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasecluListener) ExitStatement(ctx *StatementContext) {}

// EnterTag_arm is called when production tag_arm is entered.
func (s *BasecluListener) EnterTag_arm(ctx *Tag_armContext) {}

// ExitTag_arm is called when production tag_arm is exited.
func (s *BasecluListener) ExitTag_arm(ctx *Tag_armContext) {}

// EnterWhen_handler is called when production when_handler is entered.
func (s *BasecluListener) EnterWhen_handler(ctx *When_handlerContext) {}

// ExitWhen_handler is called when production when_handler is exited.
func (s *BasecluListener) ExitWhen_handler(ctx *When_handlerContext) {}

// EnterOthers_handler is called when production others_handler is entered.
func (s *BasecluListener) EnterOthers_handler(ctx *Others_handlerContext) {}

// ExitOthers_handler is called when production others_handler is exited.
func (s *BasecluListener) ExitOthers_handler(ctx *Others_handlerContext) {}

// EnterBody is called when production body is entered.
func (s *BasecluListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BasecluListener) ExitBody(ctx *BodyContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BasecluListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BasecluListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasecluListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasecluListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasecluListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasecluListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterInvocation is called when production invocation is entered.
func (s *BasecluListener) EnterInvocation(ctx *InvocationContext) {}

// ExitInvocation is called when production invocation is exited.
func (s *BasecluListener) ExitInvocation(ctx *InvocationContext) {}

// EnterField_list is called when production field_list is entered.
func (s *BasecluListener) EnterField_list(ctx *Field_listContext) {}

// ExitField_list is called when production field_list is exited.
func (s *BasecluListener) ExitField_list(ctx *Field_listContext) {}

// EnterField is called when production field is entered.
func (s *BasecluListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BasecluListener) ExitField(ctx *FieldContext) {}

// EnterIdn_list is called when production idn_list is entered.
func (s *BasecluListener) EnterIdn_list(ctx *Idn_listContext) {}

// ExitIdn_list is called when production idn_list is exited.
func (s *BasecluListener) ExitIdn_list(ctx *Idn_listContext) {}

// EnterIdn is called when production idn is entered.
func (s *BasecluListener) EnterIdn(ctx *IdnContext) {}

// ExitIdn is called when production idn is exited.
func (s *BasecluListener) ExitIdn(ctx *IdnContext) {}

// EnterName_list is called when production name_list is entered.
func (s *BasecluListener) EnterName_list(ctx *Name_listContext) {}

// ExitName_list is called when production name_list is exited.
func (s *BasecluListener) ExitName_list(ctx *Name_listContext) {}

// EnterName is called when production name is entered.
func (s *BasecluListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasecluListener) ExitName(ctx *NameContext) {}

// EnterInt_literal is called when production int_literal is entered.
func (s *BasecluListener) EnterInt_literal(ctx *Int_literalContext) {}

// ExitInt_literal is called when production int_literal is exited.
func (s *BasecluListener) ExitInt_literal(ctx *Int_literalContext) {}

// EnterReal_literal is called when production real_literal is entered.
func (s *BasecluListener) EnterReal_literal(ctx *Real_literalContext) {}

// ExitReal_literal is called when production real_literal is exited.
func (s *BasecluListener) ExitReal_literal(ctx *Real_literalContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BasecluListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BasecluListener) ExitString_literal(ctx *String_literalContext) {}
