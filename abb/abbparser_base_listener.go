// Code generated from abbParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package abb // abbParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseabbParserListener is a complete listener for a parse tree produced by abbParser.
type BaseabbParserListener struct{}

var _ abbParserListener = &BaseabbParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseabbParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseabbParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseabbParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseabbParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterModule is called when production module is entered.
func (s *BaseabbParserListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseabbParserListener) ExitModule(ctx *ModuleContext) {}

// EnterModuleData is called when production moduleData is entered.
func (s *BaseabbParserListener) EnterModuleData(ctx *ModuleDataContext) {}

// ExitModuleData is called when production moduleData is exited.
func (s *BaseabbParserListener) ExitModuleData(ctx *ModuleDataContext) {}

// EnterModuleName is called when production moduleName is entered.
func (s *BaseabbParserListener) EnterModuleName(ctx *ModuleNameContext) {}

// ExitModuleName is called when production moduleName is exited.
func (s *BaseabbParserListener) ExitModuleName(ctx *ModuleNameContext) {}

// EnterDataList is called when production dataList is entered.
func (s *BaseabbParserListener) EnterDataList(ctx *DataListContext) {}

// ExitDataList is called when production dataList is exited.
func (s *BaseabbParserListener) ExitDataList(ctx *DataListContext) {}

// EnterProcedure is called when production procedure is entered.
func (s *BaseabbParserListener) EnterProcedure(ctx *ProcedureContext) {}

// ExitProcedure is called when production procedure is exited.
func (s *BaseabbParserListener) ExitProcedure(ctx *ProcedureContext) {}

// EnterProcCall is called when production procCall is entered.
func (s *BaseabbParserListener) EnterProcCall(ctx *ProcCallContext) {}

// ExitProcCall is called when production procCall is exited.
func (s *BaseabbParserListener) ExitProcCall(ctx *ProcCallContext) {}

// EnterProcName is called when production procName is entered.
func (s *BaseabbParserListener) EnterProcName(ctx *ProcNameContext) {}

// ExitProcName is called when production procName is exited.
func (s *BaseabbParserListener) ExitProcName(ctx *ProcNameContext) {}

// EnterProcParameter is called when production procParameter is entered.
func (s *BaseabbParserListener) EnterProcParameter(ctx *ProcParameterContext) {}

// ExitProcParameter is called when production procParameter is exited.
func (s *BaseabbParserListener) ExitProcParameter(ctx *ProcParameterContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseabbParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseabbParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionParameter is called when production functionParameter is entered.
func (s *BaseabbParserListener) EnterFunctionParameter(ctx *FunctionParameterContext) {}

// ExitFunctionParameter is called when production functionParameter is exited.
func (s *BaseabbParserListener) ExitFunctionParameter(ctx *FunctionParameterContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseabbParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseabbParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseabbParserListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseabbParserListener) ExitType_(ctx *Type_Context) {}

// EnterInit_ is called when production init_ is entered.
func (s *BaseabbParserListener) EnterInit_(ctx *Init_Context) {}

// ExitInit_ is called when production init_ is exited.
func (s *BaseabbParserListener) ExitInit_(ctx *Init_Context) {}

// EnterExpression is called when production expression is entered.
func (s *BaseabbParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseabbParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArray_ is called when production array_ is entered.
func (s *BaseabbParserListener) EnterArray_(ctx *Array_Context) {}

// ExitArray_ is called when production array_ is exited.
func (s *BaseabbParserListener) ExitArray_(ctx *Array_Context) {}

// EnterPrimitive is called when production primitive is entered.
func (s *BaseabbParserListener) EnterPrimitive(ctx *PrimitiveContext) {}

// ExitPrimitive is called when production primitive is exited.
func (s *BaseabbParserListener) ExitPrimitive(ctx *PrimitiveContext) {}
