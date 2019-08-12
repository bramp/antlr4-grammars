// Code generated from COOL.g4 by ANTLR 4.7.2. DO NOT EDIT.

package cool // COOL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// COOLListener is a complete listener for a parse tree produced by COOLParser.
type COOLListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterClasses is called when entering the classes production.
	EnterClasses(c *ClassesContext)

	// EnterEof is called when entering the eof production.
	EnterEof(c *EofContext)

	// EnterClassDefine is called when entering the classDefine production.
	EnterClassDefine(c *ClassDefineContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterFormal is called when entering the formal production.
	EnterFormal(c *FormalContext)

	// EnterLetIn is called when entering the letIn production.
	EnterLetIn(c *LetInContext)

	// EnterMinus is called when entering the minus production.
	EnterMinus(c *MinusContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterIsvoid is called when entering the isvoid production.
	EnterIsvoid(c *IsvoidContext)

	// EnterWhile is called when entering the while production.
	EnterWhile(c *WhileContext)

	// EnterDivision is called when entering the division production.
	EnterDivision(c *DivisionContext)

	// EnterNegative is called when entering the negative production.
	EnterNegative(c *NegativeContext)

	// EnterBoolNot is called when entering the boolNot production.
	EnterBoolNot(c *BoolNotContext)

	// EnterLessThan is called when entering the lessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterMultiply is called when entering the multiply production.
	EnterMultiply(c *MultiplyContext)

	// EnterIf is called when entering the if production.
	EnterIf(c *IfContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterOwnMethodCall is called when entering the ownMethodCall production.
	EnterOwnMethodCall(c *OwnMethodCallContext)

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

	// EnterNew is called when entering the new production.
	EnterNew(c *NewContext)

	// EnterParentheses is called when entering the parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterFalse is called when entering the false production.
	EnterFalse(c *FalseContext)

	// EnterInt is called when entering the int production.
	EnterInt(c *IntContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterTrue is called when entering the true production.
	EnterTrue(c *TrueContext)

	// EnterLessEqual is called when entering the lessEqual production.
	EnterLessEqual(c *LessEqualContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitClasses is called when exiting the classes production.
	ExitClasses(c *ClassesContext)

	// ExitEof is called when exiting the eof production.
	ExitEof(c *EofContext)

	// ExitClassDefine is called when exiting the classDefine production.
	ExitClassDefine(c *ClassDefineContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitFormal is called when exiting the formal production.
	ExitFormal(c *FormalContext)

	// ExitLetIn is called when exiting the letIn production.
	ExitLetIn(c *LetInContext)

	// ExitMinus is called when exiting the minus production.
	ExitMinus(c *MinusContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitIsvoid is called when exiting the isvoid production.
	ExitIsvoid(c *IsvoidContext)

	// ExitWhile is called when exiting the while production.
	ExitWhile(c *WhileContext)

	// ExitDivision is called when exiting the division production.
	ExitDivision(c *DivisionContext)

	// ExitNegative is called when exiting the negative production.
	ExitNegative(c *NegativeContext)

	// ExitBoolNot is called when exiting the boolNot production.
	ExitBoolNot(c *BoolNotContext)

	// ExitLessThan is called when exiting the lessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitMultiply is called when exiting the multiply production.
	ExitMultiply(c *MultiplyContext)

	// ExitIf is called when exiting the if production.
	ExitIf(c *IfContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitOwnMethodCall is called when exiting the ownMethodCall production.
	ExitOwnMethodCall(c *OwnMethodCallContext)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

	// ExitNew is called when exiting the new production.
	ExitNew(c *NewContext)

	// ExitParentheses is called when exiting the parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitFalse is called when exiting the false production.
	ExitFalse(c *FalseContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitTrue is called when exiting the true production.
	ExitTrue(c *TrueContext)

	// ExitLessEqual is called when exiting the lessEqual production.
	ExitLessEqual(c *LessEqualContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)
}
