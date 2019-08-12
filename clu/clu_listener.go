// Code generated from clu.g4 by ANTLR 4.7.2. DO NOT EDIT.

package clu // clu
import "github.com/antlr/antlr4/runtime/Go/antlr"

// cluListener is a complete listener for a parse tree produced by cluParser.
type cluListener interface {
	antlr.ParseTreeListener

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterProcedure is called when entering the procedure production.
	EnterProcedure(c *ProcedureContext)

	// EnterIterator is called when entering the iterator production.
	EnterIterator(c *IteratorContext)

	// EnterCluster is called when entering the cluster production.
	EnterCluster(c *ClusterContext)

	// EnterParms is called when entering the parms production.
	EnterParms(c *ParmsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterDecl_list is called when entering the decl_list production.
	EnterDecl_list(c *Decl_listContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterReturnz is called when entering the returnz production.
	EnterReturnz(c *ReturnzContext)

	// EnterYields is called when entering the yields production.
	EnterYields(c *YieldsContext)

	// EnterSignals is called when entering the signals production.
	EnterSignals(c *SignalsContext)

	// EnterException is called when entering the exception production.
	EnterException(c *ExceptionContext)

	// EnterType_spec_list is called when entering the type_spec_list production.
	EnterType_spec_list(c *Type_spec_listContext)

	// EnterWhere is called when entering the where production.
	EnterWhere(c *WhereContext)

	// EnterRestriction is called when entering the restriction production.
	EnterRestriction(c *RestrictionContext)

	// EnterType_set is called when entering the type_set production.
	EnterType_set(c *Type_setContext)

	// EnterOper_decl_list is called when entering the oper_decl_list production.
	EnterOper_decl_list(c *Oper_decl_listContext)

	// EnterOper_decl is called when entering the oper_decl production.
	EnterOper_decl(c *Oper_declContext)

	// EnterOp_name_list is called when entering the op_name_list production.
	EnterOp_name_list(c *Op_name_listContext)

	// EnterOp_name is called when entering the op_name production.
	EnterOp_name(c *Op_nameContext)

	// EnterConstant_list is called when entering the constant_list production.
	EnterConstant_list(c *Constant_listContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterRoutine_body is called when entering the routine_body production.
	EnterRoutine_body(c *Routine_bodyContext)

	// EnterCluster_body is called when entering the cluster_body production.
	EnterCluster_body(c *Cluster_bodyContext)

	// EnterRoutine is called when entering the routine production.
	EnterRoutine(c *RoutineContext)

	// EnterEquate is called when entering the equate production.
	EnterEquate(c *EquateContext)

	// EnterOwn_var is called when entering the own_var production.
	EnterOwn_var(c *Own_varContext)

	// EnterType_spec is called when entering the type_spec production.
	EnterType_spec(c *Type_specContext)

	// EnterField_spec_list is called when entering the field_spec_list production.
	EnterField_spec_list(c *Field_spec_listContext)

	// EnterField_spec is called when entering the field_spec production.
	EnterField_spec(c *Field_specContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterTag_arm is called when entering the tag_arm production.
	EnterTag_arm(c *Tag_armContext)

	// EnterWhen_handler is called when entering the when_handler production.
	EnterWhen_handler(c *When_handlerContext)

	// EnterOthers_handler is called when entering the others_handler production.
	EnterOthers_handler(c *Others_handlerContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterInvocation is called when entering the invocation production.
	EnterInvocation(c *InvocationContext)

	// EnterField_list is called when entering the field_list production.
	EnterField_list(c *Field_listContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterIdn_list is called when entering the idn_list production.
	EnterIdn_list(c *Idn_listContext)

	// EnterIdn is called when entering the idn production.
	EnterIdn(c *IdnContext)

	// EnterName_list is called when entering the name_list production.
	EnterName_list(c *Name_listContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterInt_literal is called when entering the int_literal production.
	EnterInt_literal(c *Int_literalContext)

	// EnterReal_literal is called when entering the real_literal production.
	EnterReal_literal(c *Real_literalContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitProcedure is called when exiting the procedure production.
	ExitProcedure(c *ProcedureContext)

	// ExitIterator is called when exiting the iterator production.
	ExitIterator(c *IteratorContext)

	// ExitCluster is called when exiting the cluster production.
	ExitCluster(c *ClusterContext)

	// ExitParms is called when exiting the parms production.
	ExitParms(c *ParmsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitDecl_list is called when exiting the decl_list production.
	ExitDecl_list(c *Decl_listContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitReturnz is called when exiting the returnz production.
	ExitReturnz(c *ReturnzContext)

	// ExitYields is called when exiting the yields production.
	ExitYields(c *YieldsContext)

	// ExitSignals is called when exiting the signals production.
	ExitSignals(c *SignalsContext)

	// ExitException is called when exiting the exception production.
	ExitException(c *ExceptionContext)

	// ExitType_spec_list is called when exiting the type_spec_list production.
	ExitType_spec_list(c *Type_spec_listContext)

	// ExitWhere is called when exiting the where production.
	ExitWhere(c *WhereContext)

	// ExitRestriction is called when exiting the restriction production.
	ExitRestriction(c *RestrictionContext)

	// ExitType_set is called when exiting the type_set production.
	ExitType_set(c *Type_setContext)

	// ExitOper_decl_list is called when exiting the oper_decl_list production.
	ExitOper_decl_list(c *Oper_decl_listContext)

	// ExitOper_decl is called when exiting the oper_decl production.
	ExitOper_decl(c *Oper_declContext)

	// ExitOp_name_list is called when exiting the op_name_list production.
	ExitOp_name_list(c *Op_name_listContext)

	// ExitOp_name is called when exiting the op_name production.
	ExitOp_name(c *Op_nameContext)

	// ExitConstant_list is called when exiting the constant_list production.
	ExitConstant_list(c *Constant_listContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitRoutine_body is called when exiting the routine_body production.
	ExitRoutine_body(c *Routine_bodyContext)

	// ExitCluster_body is called when exiting the cluster_body production.
	ExitCluster_body(c *Cluster_bodyContext)

	// ExitRoutine is called when exiting the routine production.
	ExitRoutine(c *RoutineContext)

	// ExitEquate is called when exiting the equate production.
	ExitEquate(c *EquateContext)

	// ExitOwn_var is called when exiting the own_var production.
	ExitOwn_var(c *Own_varContext)

	// ExitType_spec is called when exiting the type_spec production.
	ExitType_spec(c *Type_specContext)

	// ExitField_spec_list is called when exiting the field_spec_list production.
	ExitField_spec_list(c *Field_spec_listContext)

	// ExitField_spec is called when exiting the field_spec production.
	ExitField_spec(c *Field_specContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitTag_arm is called when exiting the tag_arm production.
	ExitTag_arm(c *Tag_armContext)

	// ExitWhen_handler is called when exiting the when_handler production.
	ExitWhen_handler(c *When_handlerContext)

	// ExitOthers_handler is called when exiting the others_handler production.
	ExitOthers_handler(c *Others_handlerContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitInvocation is called when exiting the invocation production.
	ExitInvocation(c *InvocationContext)

	// ExitField_list is called when exiting the field_list production.
	ExitField_list(c *Field_listContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitIdn_list is called when exiting the idn_list production.
	ExitIdn_list(c *Idn_listContext)

	// ExitIdn is called when exiting the idn production.
	ExitIdn(c *IdnContext)

	// ExitName_list is called when exiting the name_list production.
	ExitName_list(c *Name_listContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitInt_literal is called when exiting the int_literal production.
	ExitInt_literal(c *Int_literalContext)

	// ExitReal_literal is called when exiting the real_literal production.
	ExitReal_literal(c *Real_literalContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)
}
