// Code generated from asmZ80.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asmz80 // asmZ80
import "github.com/antlr/antlr4/runtime/Go/antlr"

// asmZ80Listener is a complete listener for a parse tree produced by asmZ80Parser.
type asmZ80Listener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// EnterRegister_ is called when entering the register_ production.
	EnterRegister_(c *Register_Context)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterAssemblerdirective is called when entering the assemblerdirective production.
	EnterAssemblerdirective(c *AssemblerdirectiveContext)

	// EnterLbl is called when entering the lbl production.
	EnterLbl(c *LblContext)

	// EnterExpressionlist is called when entering the expressionlist production.
	EnterExpressionlist(c *ExpressionlistContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMultiplyingExpression is called when entering the multiplyingExpression production.
	EnterMultiplyingExpression(c *MultiplyingExpressionContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterDollar is called when entering the dollar production.
	EnterDollar(c *DollarContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)

	// ExitRegister_ is called when exiting the register_ production.
	ExitRegister_(c *Register_Context)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitAssemblerdirective is called when exiting the assemblerdirective production.
	ExitAssemblerdirective(c *AssemblerdirectiveContext)

	// ExitLbl is called when exiting the lbl production.
	ExitLbl(c *LblContext)

	// ExitExpressionlist is called when exiting the expressionlist production.
	ExitExpressionlist(c *ExpressionlistContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMultiplyingExpression is called when exiting the multiplyingExpression production.
	ExitMultiplyingExpression(c *MultiplyingExpressionContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitDollar is called when exiting the dollar production.
	ExitDollar(c *DollarContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
