// Code generated from asmZ80.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asmz80 // asmZ80
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseasmZ80Listener is a complete listener for a parse tree produced by asmZ80Parser.
type BaseasmZ80Listener struct{}

var _ asmZ80Listener = &BaseasmZ80Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseasmZ80Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseasmZ80Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseasmZ80Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseasmZ80Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseasmZ80Listener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseasmZ80Listener) ExitProg(ctx *ProgContext) {}

// EnterLine is called when production line is entered.
func (s *BaseasmZ80Listener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseasmZ80Listener) ExitLine(ctx *LineContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseasmZ80Listener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseasmZ80Listener) ExitInstruction(ctx *InstructionContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *BaseasmZ80Listener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *BaseasmZ80Listener) ExitOpcode(ctx *OpcodeContext) {}

// EnterRegister_ is called when production register_ is entered.
func (s *BaseasmZ80Listener) EnterRegister_(ctx *Register_Context) {}

// ExitRegister_ is called when production register_ is exited.
func (s *BaseasmZ80Listener) ExitRegister_(ctx *Register_Context) {}

// EnterDirective is called when production directive is entered.
func (s *BaseasmZ80Listener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseasmZ80Listener) ExitDirective(ctx *DirectiveContext) {}

// EnterAssemblerdirective is called when production assemblerdirective is entered.
func (s *BaseasmZ80Listener) EnterAssemblerdirective(ctx *AssemblerdirectiveContext) {}

// ExitAssemblerdirective is called when production assemblerdirective is exited.
func (s *BaseasmZ80Listener) ExitAssemblerdirective(ctx *AssemblerdirectiveContext) {}

// EnterLbl is called when production lbl is entered.
func (s *BaseasmZ80Listener) EnterLbl(ctx *LblContext) {}

// ExitLbl is called when production lbl is exited.
func (s *BaseasmZ80Listener) ExitLbl(ctx *LblContext) {}

// EnterExpressionlist is called when production expressionlist is entered.
func (s *BaseasmZ80Listener) EnterExpressionlist(ctx *ExpressionlistContext) {}

// ExitExpressionlist is called when production expressionlist is exited.
func (s *BaseasmZ80Listener) ExitExpressionlist(ctx *ExpressionlistContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseasmZ80Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseasmZ80Listener) ExitLabel(ctx *LabelContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseasmZ80Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseasmZ80Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterMultiplyingExpression is called when production multiplyingExpression is entered.
func (s *BaseasmZ80Listener) EnterMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// ExitMultiplyingExpression is called when production multiplyingExpression is exited.
func (s *BaseasmZ80Listener) ExitMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseasmZ80Listener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseasmZ80Listener) ExitArgument(ctx *ArgumentContext) {}

// EnterDollar is called when production dollar is entered.
func (s *BaseasmZ80Listener) EnterDollar(ctx *DollarContext) {}

// ExitDollar is called when production dollar is exited.
func (s *BaseasmZ80Listener) ExitDollar(ctx *DollarContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BaseasmZ80Listener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BaseasmZ80Listener) ExitString_(ctx *String_Context) {}

// EnterName is called when production name is entered.
func (s *BaseasmZ80Listener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseasmZ80Listener) ExitName(ctx *NameContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseasmZ80Listener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseasmZ80Listener) ExitNumber(ctx *NumberContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseasmZ80Listener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseasmZ80Listener) ExitComment(ctx *CommentContext) {}
