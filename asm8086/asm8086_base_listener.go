// Code generated from asm8086.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asm8086 // asm8086
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Baseasm8086Listener is a complete listener for a parse tree produced by asm8086Parser.
type Baseasm8086Listener struct{}

var _ asm8086Listener = &Baseasm8086Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baseasm8086Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baseasm8086Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baseasm8086Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baseasm8086Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *Baseasm8086Listener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *Baseasm8086Listener) ExitProg(ctx *ProgContext) {}

// EnterLine is called when production line is entered.
func (s *Baseasm8086Listener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *Baseasm8086Listener) ExitLine(ctx *LineContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *Baseasm8086Listener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *Baseasm8086Listener) ExitInstruction(ctx *InstructionContext) {}

// EnterLbl is called when production lbl is entered.
func (s *Baseasm8086Listener) EnterLbl(ctx *LblContext) {}

// ExitLbl is called when production lbl is exited.
func (s *Baseasm8086Listener) ExitLbl(ctx *LblContext) {}

// EnterAssemblerdirective is called when production assemblerdirective is entered.
func (s *Baseasm8086Listener) EnterAssemblerdirective(ctx *AssemblerdirectiveContext) {}

// ExitAssemblerdirective is called when production assemblerdirective is exited.
func (s *Baseasm8086Listener) ExitAssemblerdirective(ctx *AssemblerdirectiveContext) {}

// EnterRw is called when production rw is entered.
func (s *Baseasm8086Listener) EnterRw(ctx *RwContext) {}

// ExitRw is called when production rw is exited.
func (s *Baseasm8086Listener) ExitRw(ctx *RwContext) {}

// EnterRb is called when production rb is entered.
func (s *Baseasm8086Listener) EnterRb(ctx *RbContext) {}

// ExitRb is called when production rb is exited.
func (s *Baseasm8086Listener) ExitRb(ctx *RbContext) {}

// EnterRs is called when production rs is entered.
func (s *Baseasm8086Listener) EnterRs(ctx *RsContext) {}

// ExitRs is called when production rs is exited.
func (s *Baseasm8086Listener) ExitRs(ctx *RsContext) {}

// EnterCseg is called when production cseg is entered.
func (s *Baseasm8086Listener) EnterCseg(ctx *CsegContext) {}

// ExitCseg is called when production cseg is exited.
func (s *Baseasm8086Listener) ExitCseg(ctx *CsegContext) {}

// EnterDseg is called when production dseg is entered.
func (s *Baseasm8086Listener) EnterDseg(ctx *DsegContext) {}

// ExitDseg is called when production dseg is exited.
func (s *Baseasm8086Listener) ExitDseg(ctx *DsegContext) {}

// EnterDw is called when production dw is entered.
func (s *Baseasm8086Listener) EnterDw(ctx *DwContext) {}

// ExitDw is called when production dw is exited.
func (s *Baseasm8086Listener) ExitDw(ctx *DwContext) {}

// EnterDb is called when production db is entered.
func (s *Baseasm8086Listener) EnterDb(ctx *DbContext) {}

// ExitDb is called when production db is exited.
func (s *Baseasm8086Listener) ExitDb(ctx *DbContext) {}

// EnterDd is called when production dd is entered.
func (s *Baseasm8086Listener) EnterDd(ctx *DdContext) {}

// ExitDd is called when production dd is exited.
func (s *Baseasm8086Listener) ExitDd(ctx *DdContext) {}

// EnterEqu is called when production equ is entered.
func (s *Baseasm8086Listener) EnterEqu(ctx *EquContext) {}

// ExitEqu is called when production equ is exited.
func (s *Baseasm8086Listener) ExitEqu(ctx *EquContext) {}

// EnterIf_ is called when production if_ is entered.
func (s *Baseasm8086Listener) EnterIf_(ctx *If_Context) {}

// ExitIf_ is called when production if_ is exited.
func (s *Baseasm8086Listener) ExitIf_(ctx *If_Context) {}

// EnterAssemblerexpression is called when production assemblerexpression is entered.
func (s *Baseasm8086Listener) EnterAssemblerexpression(ctx *AssemblerexpressionContext) {}

// ExitAssemblerexpression is called when production assemblerexpression is exited.
func (s *Baseasm8086Listener) ExitAssemblerexpression(ctx *AssemblerexpressionContext) {}

// EnterAssemblerlogical is called when production assemblerlogical is entered.
func (s *Baseasm8086Listener) EnterAssemblerlogical(ctx *AssemblerlogicalContext) {}

// ExitAssemblerlogical is called when production assemblerlogical is exited.
func (s *Baseasm8086Listener) ExitAssemblerlogical(ctx *AssemblerlogicalContext) {}

// EnterAssemblerterm is called when production assemblerterm is entered.
func (s *Baseasm8086Listener) EnterAssemblerterm(ctx *AssemblertermContext) {}

// ExitAssemblerterm is called when production assemblerterm is exited.
func (s *Baseasm8086Listener) ExitAssemblerterm(ctx *AssemblertermContext) {}

// EnterEndif_ is called when production endif_ is entered.
func (s *Baseasm8086Listener) EnterEndif_(ctx *Endif_Context) {}

// ExitEndif_ is called when production endif_ is exited.
func (s *Baseasm8086Listener) ExitEndif_(ctx *Endif_Context) {}

// EnterEnd is called when production end is entered.
func (s *Baseasm8086Listener) EnterEnd(ctx *EndContext) {}

// ExitEnd is called when production end is exited.
func (s *Baseasm8086Listener) ExitEnd(ctx *EndContext) {}

// EnterOrg is called when production org is entered.
func (s *Baseasm8086Listener) EnterOrg(ctx *OrgContext) {}

// ExitOrg is called when production org is exited.
func (s *Baseasm8086Listener) ExitOrg(ctx *OrgContext) {}

// EnterTitle is called when production title is entered.
func (s *Baseasm8086Listener) EnterTitle(ctx *TitleContext) {}

// ExitTitle is called when production title is exited.
func (s *Baseasm8086Listener) ExitTitle(ctx *TitleContext) {}

// EnterInclude_ is called when production include_ is entered.
func (s *Baseasm8086Listener) EnterInclude_(ctx *Include_Context) {}

// ExitInclude_ is called when production include_ is exited.
func (s *Baseasm8086Listener) ExitInclude_(ctx *Include_Context) {}

// EnterExpressionlist is called when production expressionlist is entered.
func (s *Baseasm8086Listener) EnterExpressionlist(ctx *ExpressionlistContext) {}

// ExitExpressionlist is called when production expressionlist is exited.
func (s *Baseasm8086Listener) ExitExpressionlist(ctx *ExpressionlistContext) {}

// EnterLabel is called when production label is entered.
func (s *Baseasm8086Listener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *Baseasm8086Listener) ExitLabel(ctx *LabelContext) {}

// EnterExpression is called when production expression is entered.
func (s *Baseasm8086Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Baseasm8086Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterMultiplyingExpression is called when production multiplyingExpression is entered.
func (s *Baseasm8086Listener) EnterMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// ExitMultiplyingExpression is called when production multiplyingExpression is exited.
func (s *Baseasm8086Listener) ExitMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// EnterArgument is called when production argument is entered.
func (s *Baseasm8086Listener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *Baseasm8086Listener) ExitArgument(ctx *ArgumentContext) {}

// EnterPtr is called when production ptr is entered.
func (s *Baseasm8086Listener) EnterPtr(ctx *PtrContext) {}

// ExitPtr is called when production ptr is exited.
func (s *Baseasm8086Listener) ExitPtr(ctx *PtrContext) {}

// EnterDollar is called when production dollar is entered.
func (s *Baseasm8086Listener) EnterDollar(ctx *DollarContext) {}

// ExitDollar is called when production dollar is exited.
func (s *Baseasm8086Listener) ExitDollar(ctx *DollarContext) {}

// EnterRegister_ is called when production register_ is entered.
func (s *Baseasm8086Listener) EnterRegister_(ctx *Register_Context) {}

// ExitRegister_ is called when production register_ is exited.
func (s *Baseasm8086Listener) ExitRegister_(ctx *Register_Context) {}

// EnterString_ is called when production string_ is entered.
func (s *Baseasm8086Listener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *Baseasm8086Listener) ExitString_(ctx *String_Context) {}

// EnterName is called when production name is entered.
func (s *Baseasm8086Listener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *Baseasm8086Listener) ExitName(ctx *NameContext) {}

// EnterNumber is called when production number is entered.
func (s *Baseasm8086Listener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *Baseasm8086Listener) ExitNumber(ctx *NumberContext) {}

// EnterOpcode is called when production opcode is entered.
func (s *Baseasm8086Listener) EnterOpcode(ctx *OpcodeContext) {}

// ExitOpcode is called when production opcode is exited.
func (s *Baseasm8086Listener) ExitOpcode(ctx *OpcodeContext) {}

// EnterRep is called when production rep is entered.
func (s *Baseasm8086Listener) EnterRep(ctx *RepContext) {}

// ExitRep is called when production rep is exited.
func (s *Baseasm8086Listener) ExitRep(ctx *RepContext) {}

// EnterComment is called when production comment is entered.
func (s *Baseasm8086Listener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *Baseasm8086Listener) ExitComment(ctx *CommentContext) {}
