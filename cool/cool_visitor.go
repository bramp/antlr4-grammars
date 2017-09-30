// Generated from COOL.g4 by ANTLR 4.7.

package cool // COOL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by COOLParser.
type COOLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by COOLParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by COOLParser#classes.
	VisitClasses(ctx *ClassesContext) interface{}

	// Visit a parse tree produced by COOLParser#eof.
	VisitEof(ctx *EofContext) interface{}

	// Visit a parse tree produced by COOLParser#classDefine.
	VisitClassDefine(ctx *ClassDefineContext) interface{}

	// Visit a parse tree produced by COOLParser#method.
	VisitMethod(ctx *MethodContext) interface{}

	// Visit a parse tree produced by COOLParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by COOLParser#formal.
	VisitFormal(ctx *FormalContext) interface{}

	// Visit a parse tree produced by COOLParser#letIn.
	VisitLetIn(ctx *LetInContext) interface{}

	// Visit a parse tree produced by COOLParser#minus.
	VisitMinus(ctx *MinusContext) interface{}

	// Visit a parse tree produced by COOLParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by COOLParser#isvoid.
	VisitIsvoid(ctx *IsvoidContext) interface{}

	// Visit a parse tree produced by COOLParser#while.
	VisitWhile(ctx *WhileContext) interface{}

	// Visit a parse tree produced by COOLParser#division.
	VisitDivision(ctx *DivisionContext) interface{}

	// Visit a parse tree produced by COOLParser#negative.
	VisitNegative(ctx *NegativeContext) interface{}

	// Visit a parse tree produced by COOLParser#boolNot.
	VisitBoolNot(ctx *BoolNotContext) interface{}

	// Visit a parse tree produced by COOLParser#lessThan.
	VisitLessThan(ctx *LessThanContext) interface{}

	// Visit a parse tree produced by COOLParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by COOLParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by COOLParser#multiply.
	VisitMultiply(ctx *MultiplyContext) interface{}

	// Visit a parse tree produced by COOLParser#if.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by COOLParser#case.
	VisitCase(ctx *CaseContext) interface{}

	// Visit a parse tree produced by COOLParser#ownMethodCall.
	VisitOwnMethodCall(ctx *OwnMethodCallContext) interface{}

	// Visit a parse tree produced by COOLParser#add.
	VisitAdd(ctx *AddContext) interface{}

	// Visit a parse tree produced by COOLParser#new.
	VisitNew(ctx *NewContext) interface{}

	// Visit a parse tree produced by COOLParser#parentheses.
	VisitParentheses(ctx *ParenthesesContext) interface{}

	// Visit a parse tree produced by COOLParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by COOLParser#false.
	VisitFalse(ctx *FalseContext) interface{}

	// Visit a parse tree produced by COOLParser#int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by COOLParser#equal.
	VisitEqual(ctx *EqualContext) interface{}

	// Visit a parse tree produced by COOLParser#true.
	VisitTrue(ctx *TrueContext) interface{}

	// Visit a parse tree produced by COOLParser#lessEqual.
	VisitLessEqual(ctx *LessEqualContext) interface{}

	// Visit a parse tree produced by COOLParser#methodCall.
	VisitMethodCall(ctx *MethodCallContext) interface{}
}
