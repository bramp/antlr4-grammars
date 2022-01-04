// Code generated from asm8086.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asm8086 // asm8086
import "github.com/antlr/antlr4/runtime/Go/antlr"

// asm8086Listener is a complete listener for a parse tree produced by asm8086Parser.
type asm8086Listener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterLbl is called when entering the lbl production.
	EnterLbl(c *LblContext)

	// EnterAssemblerdirective is called when entering the assemblerdirective production.
	EnterAssemblerdirective(c *AssemblerdirectiveContext)

	// EnterRw is called when entering the rw production.
	EnterRw(c *RwContext)

	// EnterRb is called when entering the rb production.
	EnterRb(c *RbContext)

	// EnterRs is called when entering the rs production.
	EnterRs(c *RsContext)

	// EnterCseg is called when entering the cseg production.
	EnterCseg(c *CsegContext)

	// EnterDseg is called when entering the dseg production.
	EnterDseg(c *DsegContext)

	// EnterDw is called when entering the dw production.
	EnterDw(c *DwContext)

	// EnterDb is called when entering the db production.
	EnterDb(c *DbContext)

	// EnterDd is called when entering the dd production.
	EnterDd(c *DdContext)

	// EnterEqu is called when entering the equ production.
	EnterEqu(c *EquContext)

	// EnterIf_ is called when entering the if_ production.
	EnterIf_(c *If_Context)

	// EnterAssemblerexpression is called when entering the assemblerexpression production.
	EnterAssemblerexpression(c *AssemblerexpressionContext)

	// EnterAssemblerlogical is called when entering the assemblerlogical production.
	EnterAssemblerlogical(c *AssemblerlogicalContext)

	// EnterAssemblerterm is called when entering the assemblerterm production.
	EnterAssemblerterm(c *AssemblertermContext)

	// EnterEndif_ is called when entering the endif_ production.
	EnterEndif_(c *Endif_Context)

	// EnterEnd is called when entering the end production.
	EnterEnd(c *EndContext)

	// EnterOrg is called when entering the org production.
	EnterOrg(c *OrgContext)

	// EnterTitle is called when entering the title production.
	EnterTitle(c *TitleContext)

	// EnterInclude_ is called when entering the include_ production.
	EnterInclude_(c *Include_Context)

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

	// EnterPtr is called when entering the ptr production.
	EnterPtr(c *PtrContext)

	// EnterDollar is called when entering the dollar production.
	EnterDollar(c *DollarContext)

	// EnterRegister_ is called when entering the register_ production.
	EnterRegister_(c *Register_Context)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// EnterRep is called when entering the rep production.
	EnterRep(c *RepContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitLbl is called when exiting the lbl production.
	ExitLbl(c *LblContext)

	// ExitAssemblerdirective is called when exiting the assemblerdirective production.
	ExitAssemblerdirective(c *AssemblerdirectiveContext)

	// ExitRw is called when exiting the rw production.
	ExitRw(c *RwContext)

	// ExitRb is called when exiting the rb production.
	ExitRb(c *RbContext)

	// ExitRs is called when exiting the rs production.
	ExitRs(c *RsContext)

	// ExitCseg is called when exiting the cseg production.
	ExitCseg(c *CsegContext)

	// ExitDseg is called when exiting the dseg production.
	ExitDseg(c *DsegContext)

	// ExitDw is called when exiting the dw production.
	ExitDw(c *DwContext)

	// ExitDb is called when exiting the db production.
	ExitDb(c *DbContext)

	// ExitDd is called when exiting the dd production.
	ExitDd(c *DdContext)

	// ExitEqu is called when exiting the equ production.
	ExitEqu(c *EquContext)

	// ExitIf_ is called when exiting the if_ production.
	ExitIf_(c *If_Context)

	// ExitAssemblerexpression is called when exiting the assemblerexpression production.
	ExitAssemblerexpression(c *AssemblerexpressionContext)

	// ExitAssemblerlogical is called when exiting the assemblerlogical production.
	ExitAssemblerlogical(c *AssemblerlogicalContext)

	// ExitAssemblerterm is called when exiting the assemblerterm production.
	ExitAssemblerterm(c *AssemblertermContext)

	// ExitEndif_ is called when exiting the endif_ production.
	ExitEndif_(c *Endif_Context)

	// ExitEnd is called when exiting the end production.
	ExitEnd(c *EndContext)

	// ExitOrg is called when exiting the org production.
	ExitOrg(c *OrgContext)

	// ExitTitle is called when exiting the title production.
	ExitTitle(c *TitleContext)

	// ExitInclude_ is called when exiting the include_ production.
	ExitInclude_(c *Include_Context)

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

	// ExitPtr is called when exiting the ptr production.
	ExitPtr(c *PtrContext)

	// ExitDollar is called when exiting the dollar production.
	ExitDollar(c *DollarContext)

	// ExitRegister_ is called when exiting the register_ production.
	ExitRegister_(c *Register_Context)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)

	// ExitRep is called when exiting the rep production.
	ExitRep(c *RepContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
