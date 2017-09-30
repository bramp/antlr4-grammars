// Generated from COOL.g4 by ANTLR 4.7.

package cool // COOL
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCOOLListener is a complete listener for a parse tree produced by COOLParser.
type BaseCOOLListener struct{}

var _ COOLListener = &BaseCOOLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCOOLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCOOLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCOOLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCOOLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseCOOLListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseCOOLListener) ExitProgram(ctx *ProgramContext) {}

// EnterClasses is called when production classes is entered.
func (s *BaseCOOLListener) EnterClasses(ctx *ClassesContext) {}

// ExitClasses is called when production classes is exited.
func (s *BaseCOOLListener) ExitClasses(ctx *ClassesContext) {}

// EnterEof is called when production eof is entered.
func (s *BaseCOOLListener) EnterEof(ctx *EofContext) {}

// ExitEof is called when production eof is exited.
func (s *BaseCOOLListener) ExitEof(ctx *EofContext) {}

// EnterClassDefine is called when production classDefine is entered.
func (s *BaseCOOLListener) EnterClassDefine(ctx *ClassDefineContext) {}

// ExitClassDefine is called when production classDefine is exited.
func (s *BaseCOOLListener) ExitClassDefine(ctx *ClassDefineContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseCOOLListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseCOOLListener) ExitMethod(ctx *MethodContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseCOOLListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseCOOLListener) ExitProperty(ctx *PropertyContext) {}

// EnterFormal is called when production formal is entered.
func (s *BaseCOOLListener) EnterFormal(ctx *FormalContext) {}

// ExitFormal is called when production formal is exited.
func (s *BaseCOOLListener) ExitFormal(ctx *FormalContext) {}

// EnterLetIn is called when production letIn is entered.
func (s *BaseCOOLListener) EnterLetIn(ctx *LetInContext) {}

// ExitLetIn is called when production letIn is exited.
func (s *BaseCOOLListener) ExitLetIn(ctx *LetInContext) {}

// EnterMinus is called when production minus is entered.
func (s *BaseCOOLListener) EnterMinus(ctx *MinusContext) {}

// ExitMinus is called when production minus is exited.
func (s *BaseCOOLListener) ExitMinus(ctx *MinusContext) {}

// EnterString is called when production string is entered.
func (s *BaseCOOLListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseCOOLListener) ExitString(ctx *StringContext) {}

// EnterIsvoid is called when production isvoid is entered.
func (s *BaseCOOLListener) EnterIsvoid(ctx *IsvoidContext) {}

// ExitIsvoid is called when production isvoid is exited.
func (s *BaseCOOLListener) ExitIsvoid(ctx *IsvoidContext) {}

// EnterWhile is called when production while is entered.
func (s *BaseCOOLListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *BaseCOOLListener) ExitWhile(ctx *WhileContext) {}

// EnterDivision is called when production division is entered.
func (s *BaseCOOLListener) EnterDivision(ctx *DivisionContext) {}

// ExitDivision is called when production division is exited.
func (s *BaseCOOLListener) ExitDivision(ctx *DivisionContext) {}

// EnterNegative is called when production negative is entered.
func (s *BaseCOOLListener) EnterNegative(ctx *NegativeContext) {}

// ExitNegative is called when production negative is exited.
func (s *BaseCOOLListener) ExitNegative(ctx *NegativeContext) {}

// EnterBoolNot is called when production boolNot is entered.
func (s *BaseCOOLListener) EnterBoolNot(ctx *BoolNotContext) {}

// ExitBoolNot is called when production boolNot is exited.
func (s *BaseCOOLListener) ExitBoolNot(ctx *BoolNotContext) {}

// EnterLessThan is called when production lessThan is entered.
func (s *BaseCOOLListener) EnterLessThan(ctx *LessThanContext) {}

// ExitLessThan is called when production lessThan is exited.
func (s *BaseCOOLListener) ExitLessThan(ctx *LessThanContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseCOOLListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseCOOLListener) ExitBlock(ctx *BlockContext) {}

// EnterId is called when production id is entered.
func (s *BaseCOOLListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseCOOLListener) ExitId(ctx *IdContext) {}

// EnterMultiply is called when production multiply is entered.
func (s *BaseCOOLListener) EnterMultiply(ctx *MultiplyContext) {}

// ExitMultiply is called when production multiply is exited.
func (s *BaseCOOLListener) ExitMultiply(ctx *MultiplyContext) {}

// EnterIf is called when production if is entered.
func (s *BaseCOOLListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BaseCOOLListener) ExitIf(ctx *IfContext) {}

// EnterCase is called when production case is entered.
func (s *BaseCOOLListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production case is exited.
func (s *BaseCOOLListener) ExitCase(ctx *CaseContext) {}

// EnterOwnMethodCall is called when production ownMethodCall is entered.
func (s *BaseCOOLListener) EnterOwnMethodCall(ctx *OwnMethodCallContext) {}

// ExitOwnMethodCall is called when production ownMethodCall is exited.
func (s *BaseCOOLListener) ExitOwnMethodCall(ctx *OwnMethodCallContext) {}

// EnterAdd is called when production add is entered.
func (s *BaseCOOLListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production add is exited.
func (s *BaseCOOLListener) ExitAdd(ctx *AddContext) {}

// EnterNew is called when production new is entered.
func (s *BaseCOOLListener) EnterNew(ctx *NewContext) {}

// ExitNew is called when production new is exited.
func (s *BaseCOOLListener) ExitNew(ctx *NewContext) {}

// EnterParentheses is called when production parentheses is entered.
func (s *BaseCOOLListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production parentheses is exited.
func (s *BaseCOOLListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseCOOLListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseCOOLListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterFalse is called when production false is entered.
func (s *BaseCOOLListener) EnterFalse(ctx *FalseContext) {}

// ExitFalse is called when production false is exited.
func (s *BaseCOOLListener) ExitFalse(ctx *FalseContext) {}

// EnterInt is called when production int is entered.
func (s *BaseCOOLListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseCOOLListener) ExitInt(ctx *IntContext) {}

// EnterEqual is called when production equal is entered.
func (s *BaseCOOLListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BaseCOOLListener) ExitEqual(ctx *EqualContext) {}

// EnterTrue is called when production true is entered.
func (s *BaseCOOLListener) EnterTrue(ctx *TrueContext) {}

// ExitTrue is called when production true is exited.
func (s *BaseCOOLListener) ExitTrue(ctx *TrueContext) {}

// EnterLessEqual is called when production lessEqual is entered.
func (s *BaseCOOLListener) EnterLessEqual(ctx *LessEqualContext) {}

// ExitLessEqual is called when production lessEqual is exited.
func (s *BaseCOOLListener) ExitLessEqual(ctx *LessEqualContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseCOOLListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseCOOLListener) ExitMethodCall(ctx *MethodCallContext) {}
