// Code generated from pike.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pike // pike
import "github.com/antlr/antlr4/runtime/Go/antlr"

// pikeListener is a complete listener for a parse tree produced by pikeParser.
type pikeListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterImpo is called when entering the impo production.
	EnterImpo(c *ImpoContext)

	// EnterInheritance is called when entering the inheritance production.
	EnterInheritance(c *InheritanceContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterFunction_definition is called when entering the function_definition production.
	EnterFunction_definition(c *Function_definitionContext)

	// EnterVariables is called when entering the variables production.
	EnterVariables(c *VariablesContext)

	// EnterVariable_names is called when entering the variable_names production.
	EnterVariable_names(c *Variable_namesContext)

	// EnterVariable_name is called when entering the variable_name production.
	EnterVariable_name(c *Variable_nameContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterConstant_names is called when entering the constant_names production.
	EnterConstant_names(c *Constant_namesContext)

	// EnterConstant_name is called when entering the constant_name production.
	EnterConstant_name(c *Constant_nameContext)

	// EnterClass_def is called when entering the class_def production.
	EnterClass_def(c *Class_defContext)

	// EnterClass_implementation is called when entering the class_implementation production.
	EnterClass_implementation(c *Class_implementationContext)

	// EnterModifiers is called when entering the modifiers production.
	EnterModifiers(c *ModifiersContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterCond is called when entering the cond production.
	EnterCond(c *CondContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterDo_while_stmt is called when entering the do_while_stmt production.
	EnterDo_while_stmt(c *Do_while_stmtContext)

	// EnterFor_stmt is called when entering the for_stmt production.
	EnterFor_stmt(c *For_stmtContext)

	// EnterSwitch_stmt is called when entering the switch_stmt production.
	EnterSwitch_stmt(c *Switch_stmtContext)

	// EnterCase_stmt is called when entering the case_stmt production.
	EnterCase_stmt(c *Case_stmtContext)

	// EnterDefault_stmt is called when entering the default_stmt production.
	EnterDefault_stmt(c *Default_stmtContext)

	// EnterForeach_stmt is called when entering the foreach_stmt production.
	EnterForeach_stmt(c *Foreach_stmtContext)

	// EnterBreak_stmt is called when entering the break_stmt production.
	EnterBreak_stmt(c *Break_stmtContext)

	// EnterContinue_stmt is called when entering the continue_stmt production.
	EnterContinue_stmt(c *Continue_stmtContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpression2 is called when entering the expression2 production.
	EnterExpression2(c *Expression2Context)

	// EnterExpression3 is called when entering the expression3 production.
	EnterExpression3(c *Expression3Context)

	// EnterExpression4 is called when entering the expression4 production.
	EnterExpression4(c *Expression4Context)

	// EnterExpression5 is called when entering the expression5 production.
	EnterExpression5(c *Expression5Context)

	// EnterExpression6 is called when entering the expression6 production.
	EnterExpression6(c *Expression6Context)

	// EnterExtension is called when entering the extension production.
	EnterExtension(c *ExtensionContext)

	// EnterCatch_ is called when entering the catch_ production.
	EnterCatch_(c *Catch_Context)

	// EnterGauge is called when entering the gauge production.
	EnterGauge(c *GaugeContext)

	// EnterSscanf is called when entering the sscanf production.
	EnterSscanf(c *SscanfContext)

	// EnterLvalue is called when entering the lvalue production.
	EnterLvalue(c *LvalueContext)

	// EnterLambda is called when entering the lambda production.
	EnterLambda(c *LambdaContext)

	// EnterConstant_identifier is called when entering the constant_identifier production.
	EnterConstant_identifier(c *Constant_identifierContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterMultiset is called when entering the multiset production.
	EnterMultiset(c *MultisetContext)

	// EnterMapping is called when entering the mapping production.
	EnterMapping(c *MappingContext)

	// EnterProgram_specifier is called when entering the program_specifier production.
	EnterProgram_specifier(c *Program_specifierContext)

	// EnterParenthesis is called when entering the parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterSplice_expression is called when entering the splice_expression production.
	EnterSplice_expression(c *Splice_expressionContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterFunction_type is called when entering the function_type production.
	EnterFunction_type(c *Function_typeContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitImpo is called when exiting the impo production.
	ExitImpo(c *ImpoContext)

	// ExitInheritance is called when exiting the inheritance production.
	ExitInheritance(c *InheritanceContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitFunction_definition is called when exiting the function_definition production.
	ExitFunction_definition(c *Function_definitionContext)

	// ExitVariables is called when exiting the variables production.
	ExitVariables(c *VariablesContext)

	// ExitVariable_names is called when exiting the variable_names production.
	ExitVariable_names(c *Variable_namesContext)

	// ExitVariable_name is called when exiting the variable_name production.
	ExitVariable_name(c *Variable_nameContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitConstant_names is called when exiting the constant_names production.
	ExitConstant_names(c *Constant_namesContext)

	// ExitConstant_name is called when exiting the constant_name production.
	ExitConstant_name(c *Constant_nameContext)

	// ExitClass_def is called when exiting the class_def production.
	ExitClass_def(c *Class_defContext)

	// ExitClass_implementation is called when exiting the class_implementation production.
	ExitClass_implementation(c *Class_implementationContext)

	// ExitModifiers is called when exiting the modifiers production.
	ExitModifiers(c *ModifiersContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitCond is called when exiting the cond production.
	ExitCond(c *CondContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitDo_while_stmt is called when exiting the do_while_stmt production.
	ExitDo_while_stmt(c *Do_while_stmtContext)

	// ExitFor_stmt is called when exiting the for_stmt production.
	ExitFor_stmt(c *For_stmtContext)

	// ExitSwitch_stmt is called when exiting the switch_stmt production.
	ExitSwitch_stmt(c *Switch_stmtContext)

	// ExitCase_stmt is called when exiting the case_stmt production.
	ExitCase_stmt(c *Case_stmtContext)

	// ExitDefault_stmt is called when exiting the default_stmt production.
	ExitDefault_stmt(c *Default_stmtContext)

	// ExitForeach_stmt is called when exiting the foreach_stmt production.
	ExitForeach_stmt(c *Foreach_stmtContext)

	// ExitBreak_stmt is called when exiting the break_stmt production.
	ExitBreak_stmt(c *Break_stmtContext)

	// ExitContinue_stmt is called when exiting the continue_stmt production.
	ExitContinue_stmt(c *Continue_stmtContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpression2 is called when exiting the expression2 production.
	ExitExpression2(c *Expression2Context)

	// ExitExpression3 is called when exiting the expression3 production.
	ExitExpression3(c *Expression3Context)

	// ExitExpression4 is called when exiting the expression4 production.
	ExitExpression4(c *Expression4Context)

	// ExitExpression5 is called when exiting the expression5 production.
	ExitExpression5(c *Expression5Context)

	// ExitExpression6 is called when exiting the expression6 production.
	ExitExpression6(c *Expression6Context)

	// ExitExtension is called when exiting the extension production.
	ExitExtension(c *ExtensionContext)

	// ExitCatch_ is called when exiting the catch_ production.
	ExitCatch_(c *Catch_Context)

	// ExitGauge is called when exiting the gauge production.
	ExitGauge(c *GaugeContext)

	// ExitSscanf is called when exiting the sscanf production.
	ExitSscanf(c *SscanfContext)

	// ExitLvalue is called when exiting the lvalue production.
	ExitLvalue(c *LvalueContext)

	// ExitLambda is called when exiting the lambda production.
	ExitLambda(c *LambdaContext)

	// ExitConstant_identifier is called when exiting the constant_identifier production.
	ExitConstant_identifier(c *Constant_identifierContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitMultiset is called when exiting the multiset production.
	ExitMultiset(c *MultisetContext)

	// ExitMapping is called when exiting the mapping production.
	ExitMapping(c *MappingContext)

	// ExitProgram_specifier is called when exiting the program_specifier production.
	ExitProgram_specifier(c *Program_specifierContext)

	// ExitParenthesis is called when exiting the parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitSplice_expression is called when exiting the splice_expression production.
	ExitSplice_expression(c *Splice_expressionContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitFunction_type is called when exiting the function_type production.
	ExitFunction_type(c *Function_typeContext)
}
