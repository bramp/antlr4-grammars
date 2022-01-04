// Code generated from pike.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pike // pike
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasepikeListener is a complete listener for a parse tree produced by pikeParser.
type BasepikeListener struct{}

var _ pikeListener = &BasepikeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasepikeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasepikeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasepikeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasepikeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasepikeListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasepikeListener) ExitProgram(ctx *ProgramContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BasepikeListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BasepikeListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterImpo is called when production impo is entered.
func (s *BasepikeListener) EnterImpo(ctx *ImpoContext) {}

// ExitImpo is called when production impo is exited.
func (s *BasepikeListener) ExitImpo(ctx *ImpoContext) {}

// EnterInheritance is called when production inheritance is entered.
func (s *BasepikeListener) EnterInheritance(ctx *InheritanceContext) {}

// ExitInheritance is called when production inheritance is exited.
func (s *BasepikeListener) ExitInheritance(ctx *InheritanceContext) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *BasepikeListener) EnterFunction_declaration(ctx *Function_declarationContext) {}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *BasepikeListener) ExitFunction_declaration(ctx *Function_declarationContext) {}

// EnterFunction_definition is called when production function_definition is entered.
func (s *BasepikeListener) EnterFunction_definition(ctx *Function_definitionContext) {}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BasepikeListener) ExitFunction_definition(ctx *Function_definitionContext) {}

// EnterVariables is called when production variables is entered.
func (s *BasepikeListener) EnterVariables(ctx *VariablesContext) {}

// ExitVariables is called when production variables is exited.
func (s *BasepikeListener) ExitVariables(ctx *VariablesContext) {}

// EnterVariable_names is called when production variable_names is entered.
func (s *BasepikeListener) EnterVariable_names(ctx *Variable_namesContext) {}

// ExitVariable_names is called when production variable_names is exited.
func (s *BasepikeListener) ExitVariable_names(ctx *Variable_namesContext) {}

// EnterVariable_name is called when production variable_name is entered.
func (s *BasepikeListener) EnterVariable_name(ctx *Variable_nameContext) {}

// ExitVariable_name is called when production variable_name is exited.
func (s *BasepikeListener) ExitVariable_name(ctx *Variable_nameContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasepikeListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasepikeListener) ExitConstant(ctx *ConstantContext) {}

// EnterConstant_names is called when production constant_names is entered.
func (s *BasepikeListener) EnterConstant_names(ctx *Constant_namesContext) {}

// ExitConstant_names is called when production constant_names is exited.
func (s *BasepikeListener) ExitConstant_names(ctx *Constant_namesContext) {}

// EnterConstant_name is called when production constant_name is entered.
func (s *BasepikeListener) EnterConstant_name(ctx *Constant_nameContext) {}

// ExitConstant_name is called when production constant_name is exited.
func (s *BasepikeListener) ExitConstant_name(ctx *Constant_nameContext) {}

// EnterClass_def is called when production class_def is entered.
func (s *BasepikeListener) EnterClass_def(ctx *Class_defContext) {}

// ExitClass_def is called when production class_def is exited.
func (s *BasepikeListener) ExitClass_def(ctx *Class_defContext) {}

// EnterClass_implementation is called when production class_implementation is entered.
func (s *BasepikeListener) EnterClass_implementation(ctx *Class_implementationContext) {}

// ExitClass_implementation is called when production class_implementation is exited.
func (s *BasepikeListener) ExitClass_implementation(ctx *Class_implementationContext) {}

// EnterModifiers is called when production modifiers is entered.
func (s *BasepikeListener) EnterModifiers(ctx *ModifiersContext) {}

// ExitModifiers is called when production modifiers is exited.
func (s *BasepikeListener) ExitModifiers(ctx *ModifiersContext) {}

// EnterBlock is called when production block is entered.
func (s *BasepikeListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasepikeListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasepikeListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasepikeListener) ExitStatement(ctx *StatementContext) {}

// EnterCond is called when production cond is entered.
func (s *BasepikeListener) EnterCond(ctx *CondContext) {}

// ExitCond is called when production cond is exited.
func (s *BasepikeListener) ExitCond(ctx *CondContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BasepikeListener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BasepikeListener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterDo_while_stmt is called when production do_while_stmt is entered.
func (s *BasepikeListener) EnterDo_while_stmt(ctx *Do_while_stmtContext) {}

// ExitDo_while_stmt is called when production do_while_stmt is exited.
func (s *BasepikeListener) ExitDo_while_stmt(ctx *Do_while_stmtContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BasepikeListener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BasepikeListener) ExitFor_stmt(ctx *For_stmtContext) {}

// EnterSwitch_stmt is called when production switch_stmt is entered.
func (s *BasepikeListener) EnterSwitch_stmt(ctx *Switch_stmtContext) {}

// ExitSwitch_stmt is called when production switch_stmt is exited.
func (s *BasepikeListener) ExitSwitch_stmt(ctx *Switch_stmtContext) {}

// EnterCase_stmt is called when production case_stmt is entered.
func (s *BasepikeListener) EnterCase_stmt(ctx *Case_stmtContext) {}

// ExitCase_stmt is called when production case_stmt is exited.
func (s *BasepikeListener) ExitCase_stmt(ctx *Case_stmtContext) {}

// EnterDefault_stmt is called when production default_stmt is entered.
func (s *BasepikeListener) EnterDefault_stmt(ctx *Default_stmtContext) {}

// ExitDefault_stmt is called when production default_stmt is exited.
func (s *BasepikeListener) ExitDefault_stmt(ctx *Default_stmtContext) {}

// EnterForeach_stmt is called when production foreach_stmt is entered.
func (s *BasepikeListener) EnterForeach_stmt(ctx *Foreach_stmtContext) {}

// ExitForeach_stmt is called when production foreach_stmt is exited.
func (s *BasepikeListener) ExitForeach_stmt(ctx *Foreach_stmtContext) {}

// EnterBreak_stmt is called when production break_stmt is entered.
func (s *BasepikeListener) EnterBreak_stmt(ctx *Break_stmtContext) {}

// ExitBreak_stmt is called when production break_stmt is exited.
func (s *BasepikeListener) ExitBreak_stmt(ctx *Break_stmtContext) {}

// EnterContinue_stmt is called when production continue_stmt is entered.
func (s *BasepikeListener) EnterContinue_stmt(ctx *Continue_stmtContext) {}

// ExitContinue_stmt is called when production continue_stmt is exited.
func (s *BasepikeListener) ExitContinue_stmt(ctx *Continue_stmtContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasepikeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasepikeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpression2 is called when production expression2 is entered.
func (s *BasepikeListener) EnterExpression2(ctx *Expression2Context) {}

// ExitExpression2 is called when production expression2 is exited.
func (s *BasepikeListener) ExitExpression2(ctx *Expression2Context) {}

// EnterExpression3 is called when production expression3 is entered.
func (s *BasepikeListener) EnterExpression3(ctx *Expression3Context) {}

// ExitExpression3 is called when production expression3 is exited.
func (s *BasepikeListener) ExitExpression3(ctx *Expression3Context) {}

// EnterExpression4 is called when production expression4 is entered.
func (s *BasepikeListener) EnterExpression4(ctx *Expression4Context) {}

// ExitExpression4 is called when production expression4 is exited.
func (s *BasepikeListener) ExitExpression4(ctx *Expression4Context) {}

// EnterExpression5 is called when production expression5 is entered.
func (s *BasepikeListener) EnterExpression5(ctx *Expression5Context) {}

// ExitExpression5 is called when production expression5 is exited.
func (s *BasepikeListener) ExitExpression5(ctx *Expression5Context) {}

// EnterExpression6 is called when production expression6 is entered.
func (s *BasepikeListener) EnterExpression6(ctx *Expression6Context) {}

// ExitExpression6 is called when production expression6 is exited.
func (s *BasepikeListener) ExitExpression6(ctx *Expression6Context) {}

// EnterExtension is called when production extension is entered.
func (s *BasepikeListener) EnterExtension(ctx *ExtensionContext) {}

// ExitExtension is called when production extension is exited.
func (s *BasepikeListener) ExitExtension(ctx *ExtensionContext) {}

// EnterCatch_ is called when production catch_ is entered.
func (s *BasepikeListener) EnterCatch_(ctx *Catch_Context) {}

// ExitCatch_ is called when production catch_ is exited.
func (s *BasepikeListener) ExitCatch_(ctx *Catch_Context) {}

// EnterGauge is called when production gauge is entered.
func (s *BasepikeListener) EnterGauge(ctx *GaugeContext) {}

// ExitGauge is called when production gauge is exited.
func (s *BasepikeListener) ExitGauge(ctx *GaugeContext) {}

// EnterSscanf is called when production sscanf is entered.
func (s *BasepikeListener) EnterSscanf(ctx *SscanfContext) {}

// ExitSscanf is called when production sscanf is exited.
func (s *BasepikeListener) ExitSscanf(ctx *SscanfContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BasepikeListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BasepikeListener) ExitLvalue(ctx *LvalueContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BasepikeListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BasepikeListener) ExitLambda(ctx *LambdaContext) {}

// EnterConstant_identifier is called when production constant_identifier is entered.
func (s *BasepikeListener) EnterConstant_identifier(ctx *Constant_identifierContext) {}

// ExitConstant_identifier is called when production constant_identifier is exited.
func (s *BasepikeListener) ExitConstant_identifier(ctx *Constant_identifierContext) {}

// EnterArray is called when production array is entered.
func (s *BasepikeListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BasepikeListener) ExitArray(ctx *ArrayContext) {}

// EnterMultiset is called when production multiset is entered.
func (s *BasepikeListener) EnterMultiset(ctx *MultisetContext) {}

// ExitMultiset is called when production multiset is exited.
func (s *BasepikeListener) ExitMultiset(ctx *MultisetContext) {}

// EnterMapping is called when production mapping is entered.
func (s *BasepikeListener) EnterMapping(ctx *MappingContext) {}

// ExitMapping is called when production mapping is exited.
func (s *BasepikeListener) ExitMapping(ctx *MappingContext) {}

// EnterProgram_specifier is called when production program_specifier is entered.
func (s *BasepikeListener) EnterProgram_specifier(ctx *Program_specifierContext) {}

// ExitProgram_specifier is called when production program_specifier is exited.
func (s *BasepikeListener) ExitProgram_specifier(ctx *Program_specifierContext) {}

// EnterParenthesis is called when production parenthesis is entered.
func (s *BasepikeListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production parenthesis is exited.
func (s *BasepikeListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BasepikeListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BasepikeListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterSplice_expression is called when production splice_expression is entered.
func (s *BasepikeListener) EnterSplice_expression(ctx *Splice_expressionContext) {}

// ExitSplice_expression is called when production splice_expression is exited.
func (s *BasepikeListener) ExitSplice_expression(ctx *Splice_expressionContext) {}

// EnterArgument is called when production argument is entered.
func (s *BasepikeListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BasepikeListener) ExitArgument(ctx *ArgumentContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BasepikeListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BasepikeListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BasepikeListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BasepikeListener) ExitType_(ctx *Type_Context) {}

// EnterFunction_type is called when production function_type is entered.
func (s *BasepikeListener) EnterFunction_type(ctx *Function_typeContext) {}

// ExitFunction_type is called when production function_type is exited.
func (s *BasepikeListener) ExitFunction_type(ctx *Function_typeContext) {}
