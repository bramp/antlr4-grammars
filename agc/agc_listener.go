// Generated from agc.g4 by ANTLR 4.7.

package agc // agc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// agcListener is a complete listener for a parse tree produced by agcParser.
type agcListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterBlank_line is called when entering the blank_line production.
	EnterBlank_line(c *Blank_lineContext)

	// EnterComment_line is called when entering the comment_line production.
	EnterComment_line(c *Comment_lineContext)

	// EnterInstruction_line is called when entering the instruction_line production.
	EnterInstruction_line(c *Instruction_lineContext)

	// EnterErase_line is called when entering the erase_line production.
	EnterErase_line(c *Erase_lineContext)

	// EnterAssignment_line is called when entering the assignment_line production.
	EnterAssignment_line(c *Assignment_lineContext)

	// EnterOpcodes is called when entering the opcodes production.
	EnterOpcodes(c *OpcodesContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterWs is called when entering the ws production.
	EnterWs(c *WsContext)

	// EnterEol is called when entering the eol production.
	EnterEol(c *EolContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMultiplyingExpression is called when entering the multiplyingExpression production.
	EnterMultiplyingExpression(c *MultiplyingExpressionContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterInte is called when entering the inte production.
	EnterInte(c *InteContext)

	// EnterDecimal is called when entering the decimal production.
	EnterDecimal(c *DecimalContext)

	// EnterRegister is called when entering the register production.
	EnterRegister(c *RegisterContext)

	// EnterOpcode is called when entering the opcode production.
	EnterOpcode(c *OpcodeContext)

	// EnterAxt_opcode is called when entering the axt_opcode production.
	EnterAxt_opcode(c *Axt_opcodeContext)

	// EnterPseudo_opcode is called when entering the pseudo_opcode production.
	EnterPseudo_opcode(c *Pseudo_opcodeContext)

	// EnterStandard_opcode is called when entering the standard_opcode production.
	EnterStandard_opcode(c *Standard_opcodeContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitBlank_line is called when exiting the blank_line production.
	ExitBlank_line(c *Blank_lineContext)

	// ExitComment_line is called when exiting the comment_line production.
	ExitComment_line(c *Comment_lineContext)

	// ExitInstruction_line is called when exiting the instruction_line production.
	ExitInstruction_line(c *Instruction_lineContext)

	// ExitErase_line is called when exiting the erase_line production.
	ExitErase_line(c *Erase_lineContext)

	// ExitAssignment_line is called when exiting the assignment_line production.
	ExitAssignment_line(c *Assignment_lineContext)

	// ExitOpcodes is called when exiting the opcodes production.
	ExitOpcodes(c *OpcodesContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitWs is called when exiting the ws production.
	ExitWs(c *WsContext)

	// ExitEol is called when exiting the eol production.
	ExitEol(c *EolContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMultiplyingExpression is called when exiting the multiplyingExpression production.
	ExitMultiplyingExpression(c *MultiplyingExpressionContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitInte is called when exiting the inte production.
	ExitInte(c *InteContext)

	// ExitDecimal is called when exiting the decimal production.
	ExitDecimal(c *DecimalContext)

	// ExitRegister is called when exiting the register production.
	ExitRegister(c *RegisterContext)

	// ExitOpcode is called when exiting the opcode production.
	ExitOpcode(c *OpcodeContext)

	// ExitAxt_opcode is called when exiting the axt_opcode production.
	ExitAxt_opcode(c *Axt_opcodeContext)

	// ExitPseudo_opcode is called when exiting the pseudo_opcode production.
	ExitPseudo_opcode(c *Pseudo_opcodeContext)

	// ExitStandard_opcode is called when exiting the standard_opcode production.
	ExitStandard_opcode(c *Standard_opcodeContext)
}
