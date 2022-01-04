// Code generated from abbParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package abb // abbParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// abbParserListener is a complete listener for a parse tree produced by abbParser.
type abbParserListener interface {
	antlr.ParseTreeListener

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterModuleData is called when entering the moduleData production.
	EnterModuleData(c *ModuleDataContext)

	// EnterModuleName is called when entering the moduleName production.
	EnterModuleName(c *ModuleNameContext)

	// EnterDataList is called when entering the dataList production.
	EnterDataList(c *DataListContext)

	// EnterProcedure is called when entering the procedure production.
	EnterProcedure(c *ProcedureContext)

	// EnterProcCall is called when entering the procCall production.
	EnterProcCall(c *ProcCallContext)

	// EnterProcName is called when entering the procName production.
	EnterProcName(c *ProcNameContext)

	// EnterProcParameter is called when entering the procParameter production.
	EnterProcParameter(c *ProcParameterContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterFunctionParameter is called when entering the functionParameter production.
	EnterFunctionParameter(c *FunctionParameterContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterInit_ is called when entering the init_ production.
	EnterInit_(c *Init_Context)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterArray_ is called when entering the array_ production.
	EnterArray_(c *Array_Context)

	// EnterPrimitive is called when entering the primitive production.
	EnterPrimitive(c *PrimitiveContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitModuleData is called when exiting the moduleData production.
	ExitModuleData(c *ModuleDataContext)

	// ExitModuleName is called when exiting the moduleName production.
	ExitModuleName(c *ModuleNameContext)

	// ExitDataList is called when exiting the dataList production.
	ExitDataList(c *DataListContext)

	// ExitProcedure is called when exiting the procedure production.
	ExitProcedure(c *ProcedureContext)

	// ExitProcCall is called when exiting the procCall production.
	ExitProcCall(c *ProcCallContext)

	// ExitProcName is called when exiting the procName production.
	ExitProcName(c *ProcNameContext)

	// ExitProcParameter is called when exiting the procParameter production.
	ExitProcParameter(c *ProcParameterContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitFunctionParameter is called when exiting the functionParameter production.
	ExitFunctionParameter(c *FunctionParameterContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitInit_ is called when exiting the init_ production.
	ExitInit_(c *Init_Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitArray_ is called when exiting the array_ production.
	ExitArray_(c *Array_Context)

	// ExitPrimitive is called when exiting the primitive production.
	ExitPrimitive(c *PrimitiveContext)
}
