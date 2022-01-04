// Code generated from asm8080.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asm8080 // asm8080
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Baseasm8080Listener is a complete listener for a parse tree produced by asm8080Parser.
type Baseasm8080Listener struct{}

var _ asm8080Listener = &Baseasm8080Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baseasm8080Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baseasm8080Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baseasm8080Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baseasm8080Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *Baseasm8080Listener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *Baseasm8080Listener) ExitProg(ctx *ProgContext) {}

// EnterLine is called when production line is entered.
func (s *Baseasm8080Listener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *Baseasm8080Listener) ExitLine(ctx *LineContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *Baseasm8080Listener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *Baseasm8080Listener) ExitInstruction(ctx *InstructionContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *Baseasm8080Listener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *Baseasm8080Listener) ExitOpcode(ctx *OpcodeContext) {}

// EnterRegister_ is called when production register_ is entered.
func (s *Baseasm8080Listener) EnterRegister_(ctx *Register_Context) {}

// ExitRegister_ is called when production register_ is exited.
func (s *Baseasm8080Listener) ExitRegister_(ctx *Register_Context) {}

// EnterDirective is called when production directive is entered.
func (s *Baseasm8080Listener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *Baseasm8080Listener) ExitDirective(ctx *DirectiveContext) {}

// EnterAssemblerdirective is called when production assemblerdirective is entered.
func (s *Baseasm8080Listener) EnterAssemblerdirective(ctx *AssemblerdirectiveContext) {}

// ExitAssemblerdirective is called when production assemblerdirective is exited.
func (s *Baseasm8080Listener) ExitAssemblerdirective(ctx *AssemblerdirectiveContext) {}

// EnterLbl is called when production lbl is entered.
func (s *Baseasm8080Listener) EnterLbl(ctx *LblContext) {}

// ExitLbl is called when production lbl is exited.
func (s *Baseasm8080Listener) ExitLbl(ctx *LblContext) {}

// EnterExpressionlist is called when production expressionlist is entered.
func (s *Baseasm8080Listener) EnterExpressionlist(ctx *ExpressionlistContext) {}

// ExitExpressionlist is called when production expressionlist is exited.
func (s *Baseasm8080Listener) ExitExpressionlist(ctx *ExpressionlistContext) {}

// EnterLabel is called when production label is entered.
func (s *Baseasm8080Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *Baseasm8080Listener) ExitLabel(ctx *LabelContext) {}

// EnterExpression is called when production expression is entered.
func (s *Baseasm8080Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Baseasm8080Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterMultiplyingExpression is called when production multiplyingExpression is entered.
func (s *Baseasm8080Listener) EnterMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// ExitMultiplyingExpression is called when production multiplyingExpression is exited.
func (s *Baseasm8080Listener) ExitMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// EnterArgument is called when production argument is entered.
func (s *Baseasm8080Listener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *Baseasm8080Listener) ExitArgument(ctx *ArgumentContext) {}

// EnterDollar is called when production dollar is entered.
func (s *Baseasm8080Listener) EnterDollar(ctx *DollarContext) {}

// ExitDollar is called when production dollar is exited.
func (s *Baseasm8080Listener) ExitDollar(ctx *DollarContext) {}

// EnterString_ is called when production string_ is entered.
func (s *Baseasm8080Listener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *Baseasm8080Listener) ExitString_(ctx *String_Context) {}

// EnterName is called when production name is entered.
func (s *Baseasm8080Listener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *Baseasm8080Listener) ExitName(ctx *NameContext) {}

// EnterNumber is called when production number is entered.
func (s *Baseasm8080Listener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *Baseasm8080Listener) ExitNumber(ctx *NumberContext) {}

// EnterComment is called when production comment is entered.
func (s *Baseasm8080Listener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *Baseasm8080Listener) ExitComment(ctx *CommentContext) {}
