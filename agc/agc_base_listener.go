// Generated from agc.g4 by ANTLR 4.7.

package agc // agc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseagcListener is a complete listener for a parse tree produced by agcParser.
type BaseagcListener struct{}

var _ agcListener = &BaseagcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseagcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseagcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseagcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseagcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseagcListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseagcListener) ExitProg(ctx *ProgContext) {}

// EnterLine is called when production line is entered.
func (s *BaseagcListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseagcListener) ExitLine(ctx *LineContext) {}

// EnterBlank_line is called when production blank_line is entered.
func (s *BaseagcListener) EnterBlank_line(ctx *Blank_lineContext) {}

// ExitBlank_line is called when production blank_line is exited.
func (s *BaseagcListener) ExitBlank_line(ctx *Blank_lineContext) {}

// EnterComment_line is called when production comment_line is entered.
func (s *BaseagcListener) EnterComment_line(ctx *Comment_lineContext) {}

// ExitComment_line is called when production comment_line is exited.
func (s *BaseagcListener) ExitComment_line(ctx *Comment_lineContext) {}

// EnterInstruction_line is called when production instruction_line is entered.
func (s *BaseagcListener) EnterInstruction_line(ctx *Instruction_lineContext) {}

// ExitInstruction_line is called when production instruction_line is exited.
func (s *BaseagcListener) ExitInstruction_line(ctx *Instruction_lineContext) {}

// EnterErase_line is called when production erase_line is entered.
func (s *BaseagcListener) EnterErase_line(ctx *Erase_lineContext) {}

// ExitErase_line is called when production erase_line is exited.
func (s *BaseagcListener) ExitErase_line(ctx *Erase_lineContext) {}

// EnterAssignment_line is called when production assignment_line is entered.
func (s *BaseagcListener) EnterAssignment_line(ctx *Assignment_lineContext) {}

// ExitAssignment_line is called when production assignment_line is exited.
func (s *BaseagcListener) ExitAssignment_line(ctx *Assignment_lineContext) {}

// EnterOpcodes is called when production opcodes is entered.
func (s *BaseagcListener) EnterOpcodes(ctx *OpcodesContext) {}

// ExitOpcodes is called when production opcodes is exited.
func (s *BaseagcListener) ExitOpcodes(ctx *OpcodesContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseagcListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseagcListener) ExitArgument(ctx *ArgumentContext) {}

// EnterWs is called when production ws is entered.
func (s *BaseagcListener) EnterWs(ctx *WsContext) {}

// ExitWs is called when production ws is exited.
func (s *BaseagcListener) ExitWs(ctx *WsContext) {}

// EnterEol is called when production eol is entered.
func (s *BaseagcListener) EnterEol(ctx *EolContext) {}

// ExitEol is called when production eol is exited.
func (s *BaseagcListener) ExitEol(ctx *EolContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseagcListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseagcListener) ExitComment(ctx *CommentContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseagcListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseagcListener) ExitLabel(ctx *LabelContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseagcListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseagcListener) ExitVariable(ctx *VariableContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseagcListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseagcListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMultiplyingExpression is called when production multiplyingExpression is entered.
func (s *BaseagcListener) EnterMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// ExitMultiplyingExpression is called when production multiplyingExpression is exited.
func (s *BaseagcListener) ExitMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseagcListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseagcListener) ExitAtom(ctx *AtomContext) {}

// EnterInte is called when production inte is entered.
func (s *BaseagcListener) EnterInte(ctx *InteContext) {}

// ExitInte is called when production inte is exited.
func (s *BaseagcListener) ExitInte(ctx *InteContext) {}

// EnterDecimal is called when production decimal is entered.
func (s *BaseagcListener) EnterDecimal(ctx *DecimalContext) {}

// ExitDecimal is called when production decimal is exited.
func (s *BaseagcListener) ExitDecimal(ctx *DecimalContext) {}

// EnterRegister is called when production register is entered.
func (s *BaseagcListener) EnterRegister(ctx *RegisterContext) {}

// ExitRegister is called when production register is exited.
func (s *BaseagcListener) ExitRegister(ctx *RegisterContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *BaseagcListener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *BaseagcListener) ExitOpcode(ctx *OpcodeContext) {}

// EnterAxt_opcode is called when production axt_opcode is entered.
func (s *BaseagcListener) EnterAxt_opcode(ctx *Axt_opcodeContext) {}

// ExitAxt_opcode is called when production axt_opcode is exited.
func (s *BaseagcListener) ExitAxt_opcode(ctx *Axt_opcodeContext) {}

// EnterPseudo_opcode is called when production pseudo_opcode is entered.
func (s *BaseagcListener) EnterPseudo_opcode(ctx *Pseudo_opcodeContext) {}

// ExitPseudo_opcode is called when production pseudo_opcode is exited.
func (s *BaseagcListener) ExitPseudo_opcode(ctx *Pseudo_opcodeContext) {}

// EnterStandard_opcode is called when production standard_opcode is entered.
func (s *BaseagcListener) EnterStandard_opcode(ctx *Standard_opcodeContext) {}

// ExitStandard_opcode is called when production standard_opcode is exited.
func (s *BaseagcListener) ExitStandard_opcode(ctx *Standard_opcodeContext) {}
