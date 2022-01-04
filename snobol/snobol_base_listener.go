// Code generated from snobol.g4 by ANTLR 4.9.3. DO NOT EDIT.

package snobol // snobol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesnobolListener is a complete listener for a parse tree produced by snobolParser.
type BasesnobolListener struct{}

var _ snobolListener = &BasesnobolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesnobolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesnobolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesnobolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesnobolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasesnobolListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasesnobolListener) ExitProg(ctx *ProgContext) {}

// EnterLin is called when production lin is entered.
func (s *BasesnobolListener) EnterLin(ctx *LinContext) {}

// ExitLin is called when production lin is exited.
func (s *BasesnobolListener) ExitLin(ctx *LinContext) {}

// EnterLine is called when production line is entered.
func (s *BasesnobolListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BasesnobolListener) ExitLine(ctx *LineContext) {}

// EnterLabel is called when production label is entered.
func (s *BasesnobolListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BasesnobolListener) ExitLabel(ctx *LabelContext) {}

// EnterSubject is called when production subject is entered.
func (s *BasesnobolListener) EnterSubject(ctx *SubjectContext) {}

// ExitSubject is called when production subject is exited.
func (s *BasesnobolListener) ExitSubject(ctx *SubjectContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BasesnobolListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BasesnobolListener) ExitPattern(ctx *PatternContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasesnobolListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasesnobolListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMultiplyingExpression is called when production multiplyingExpression is entered.
func (s *BasesnobolListener) EnterMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// ExitMultiplyingExpression is called when production multiplyingExpression is exited.
func (s *BasesnobolListener) ExitMultiplyingExpression(ctx *MultiplyingExpressionContext) {}

// EnterPowExpression is called when production powExpression is entered.
func (s *BasesnobolListener) EnterPowExpression(ctx *PowExpressionContext) {}

// ExitPowExpression is called when production powExpression is exited.
func (s *BasesnobolListener) ExitPowExpression(ctx *PowExpressionContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasesnobolListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasesnobolListener) ExitAtom(ctx *AtomContext) {}

// EnterCommand is called when production command is entered.
func (s *BasesnobolListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BasesnobolListener) ExitCommand(ctx *CommandContext) {}

// EnterIdent is called when production ident is entered.
func (s *BasesnobolListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BasesnobolListener) ExitIdent(ctx *IdentContext) {}

// EnterDiffer is called when production differ is entered.
func (s *BasesnobolListener) EnterDiffer(ctx *DifferContext) {}

// ExitDiffer is called when production differ is exited.
func (s *BasesnobolListener) ExitDiffer(ctx *DifferContext) {}

// EnterEq is called when production eq is entered.
func (s *BasesnobolListener) EnterEq(ctx *EqContext) {}

// ExitEq is called when production eq is exited.
func (s *BasesnobolListener) ExitEq(ctx *EqContext) {}

// EnterNe is called when production ne is entered.
func (s *BasesnobolListener) EnterNe(ctx *NeContext) {}

// ExitNe is called when production ne is exited.
func (s *BasesnobolListener) ExitNe(ctx *NeContext) {}

// EnterGe is called when production ge is entered.
func (s *BasesnobolListener) EnterGe(ctx *GeContext) {}

// ExitGe is called when production ge is exited.
func (s *BasesnobolListener) ExitGe(ctx *GeContext) {}

// EnterGt is called when production gt is entered.
func (s *BasesnobolListener) EnterGt(ctx *GtContext) {}

// ExitGt is called when production gt is exited.
func (s *BasesnobolListener) ExitGt(ctx *GtContext) {}

// EnterLe is called when production le is entered.
func (s *BasesnobolListener) EnterLe(ctx *LeContext) {}

// ExitLe is called when production le is exited.
func (s *BasesnobolListener) ExitLe(ctx *LeContext) {}

// EnterLt is called when production lt is entered.
func (s *BasesnobolListener) EnterLt(ctx *LtContext) {}

// ExitLt is called when production lt is exited.
func (s *BasesnobolListener) ExitLt(ctx *LtContext) {}

// EnterInteger is called when production integer is entered.
func (s *BasesnobolListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasesnobolListener) ExitInteger(ctx *IntegerContext) {}

// EnterLgt is called when production lgt is entered.
func (s *BasesnobolListener) EnterLgt(ctx *LgtContext) {}

// ExitLgt is called when production lgt is exited.
func (s *BasesnobolListener) ExitLgt(ctx *LgtContext) {}

// EnterAtan is called when production atan is entered.
func (s *BasesnobolListener) EnterAtan(ctx *AtanContext) {}

// ExitAtan is called when production atan is exited.
func (s *BasesnobolListener) ExitAtan(ctx *AtanContext) {}

// EnterChop is called when production chop is entered.
func (s *BasesnobolListener) EnterChop(ctx *ChopContext) {}

// ExitChop is called when production chop is exited.
func (s *BasesnobolListener) ExitChop(ctx *ChopContext) {}

// EnterCos is called when production cos is entered.
func (s *BasesnobolListener) EnterCos(ctx *CosContext) {}

// ExitCos is called when production cos is exited.
func (s *BasesnobolListener) ExitCos(ctx *CosContext) {}

// EnterExp is called when production exp is entered.
func (s *BasesnobolListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BasesnobolListener) ExitExp(ctx *ExpContext) {}

// EnterLn is called when production ln is entered.
func (s *BasesnobolListener) EnterLn(ctx *LnContext) {}

// ExitLn is called when production ln is exited.
func (s *BasesnobolListener) ExitLn(ctx *LnContext) {}

// EnterRemdr is called when production remdr is entered.
func (s *BasesnobolListener) EnterRemdr(ctx *RemdrContext) {}

// ExitRemdr is called when production remdr is exited.
func (s *BasesnobolListener) ExitRemdr(ctx *RemdrContext) {}

// EnterSin is called when production sin is entered.
func (s *BasesnobolListener) EnterSin(ctx *SinContext) {}

// ExitSin is called when production sin is exited.
func (s *BasesnobolListener) ExitSin(ctx *SinContext) {}

// EnterTan is called when production tan is entered.
func (s *BasesnobolListener) EnterTan(ctx *TanContext) {}

// ExitTan is called when production tan is exited.
func (s *BasesnobolListener) ExitTan(ctx *TanContext) {}

// EnterDupl is called when production dupl is entered.
func (s *BasesnobolListener) EnterDupl(ctx *DuplContext) {}

// ExitDupl is called when production dupl is exited.
func (s *BasesnobolListener) ExitDupl(ctx *DuplContext) {}

// EnterReverse is called when production reverse is entered.
func (s *BasesnobolListener) EnterReverse(ctx *ReverseContext) {}

// ExitReverse is called when production reverse is exited.
func (s *BasesnobolListener) ExitReverse(ctx *ReverseContext) {}

// EnterDate is called when production date is entered.
func (s *BasesnobolListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BasesnobolListener) ExitDate(ctx *DateContext) {}

// EnterReplace is called when production replace is entered.
func (s *BasesnobolListener) EnterReplace(ctx *ReplaceContext) {}

// ExitReplace is called when production replace is exited.
func (s *BasesnobolListener) ExitReplace(ctx *ReplaceContext) {}

// EnterSize is called when production size is entered.
func (s *BasesnobolListener) EnterSize(ctx *SizeContext) {}

// ExitSize is called when production size is exited.
func (s *BasesnobolListener) ExitSize(ctx *SizeContext) {}

// EnterTrim is called when production trim is entered.
func (s *BasesnobolListener) EnterTrim(ctx *TrimContext) {}

// ExitTrim is called when production trim is exited.
func (s *BasesnobolListener) ExitTrim(ctx *TrimContext) {}

// EnterArray_ is called when production array_ is entered.
func (s *BasesnobolListener) EnterArray_(ctx *Array_Context) {}

// ExitArray_ is called when production array_ is exited.
func (s *BasesnobolListener) ExitArray_(ctx *Array_Context) {}

// EnterConvert is called when production convert is entered.
func (s *BasesnobolListener) EnterConvert(ctx *ConvertContext) {}

// ExitConvert is called when production convert is exited.
func (s *BasesnobolListener) ExitConvert(ctx *ConvertContext) {}

// EnterTable is called when production table is entered.
func (s *BasesnobolListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BasesnobolListener) ExitTable(ctx *TableContext) {}

// EnterSort is called when production sort is entered.
func (s *BasesnobolListener) EnterSort(ctx *SortContext) {}

// ExitSort is called when production sort is exited.
func (s *BasesnobolListener) ExitSort(ctx *SortContext) {}

// EnterBreak_ is called when production break_ is entered.
func (s *BasesnobolListener) EnterBreak_(ctx *Break_Context) {}

// ExitBreak_ is called when production break_ is exited.
func (s *BasesnobolListener) ExitBreak_(ctx *Break_Context) {}

// EnterTransfer is called when production transfer is entered.
func (s *BasesnobolListener) EnterTransfer(ctx *TransferContext) {}

// ExitTransfer is called when production transfer is exited.
func (s *BasesnobolListener) ExitTransfer(ctx *TransferContext) {}

// EnterTransferpre is called when production transferpre is entered.
func (s *BasesnobolListener) EnterTransferpre(ctx *TransferpreContext) {}

// ExitTransferpre is called when production transferpre is exited.
func (s *BasesnobolListener) ExitTransferpre(ctx *TransferpreContext) {}
