// Code generated from snobol.g4 by ANTLR 4.7.2. DO NOT EDIT.

package snobol // snobol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// snobolListener is a complete listener for a parse tree produced by snobolParser.
type snobolListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLin is called when entering the lin production.
	EnterLin(c *LinContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterSubject is called when entering the subject production.
	EnterSubject(c *SubjectContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMultiplyingExpression is called when entering the multiplyingExpression production.
	EnterMultiplyingExpression(c *MultiplyingExpressionContext)

	// EnterPowExpression is called when entering the powExpression production.
	EnterPowExpression(c *PowExpressionContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterDiffer is called when entering the differ production.
	EnterDiffer(c *DifferContext)

	// EnterEq is called when entering the eq production.
	EnterEq(c *EqContext)

	// EnterNe is called when entering the ne production.
	EnterNe(c *NeContext)

	// EnterGe is called when entering the ge production.
	EnterGe(c *GeContext)

	// EnterGt is called when entering the gt production.
	EnterGt(c *GtContext)

	// EnterLe is called when entering the le production.
	EnterLe(c *LeContext)

	// EnterLt is called when entering the lt production.
	EnterLt(c *LtContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterLgt is called when entering the lgt production.
	EnterLgt(c *LgtContext)

	// EnterAtan is called when entering the atan production.
	EnterAtan(c *AtanContext)

	// EnterChop is called when entering the chop production.
	EnterChop(c *ChopContext)

	// EnterCos is called when entering the cos production.
	EnterCos(c *CosContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterLn is called when entering the ln production.
	EnterLn(c *LnContext)

	// EnterRemdr is called when entering the remdr production.
	EnterRemdr(c *RemdrContext)

	// EnterSin is called when entering the sin production.
	EnterSin(c *SinContext)

	// EnterTan is called when entering the tan production.
	EnterTan(c *TanContext)

	// EnterDupl is called when entering the dupl production.
	EnterDupl(c *DuplContext)

	// EnterReverse is called when entering the reverse production.
	EnterReverse(c *ReverseContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterReplace is called when entering the replace production.
	EnterReplace(c *ReplaceContext)

	// EnterSize is called when entering the size production.
	EnterSize(c *SizeContext)

	// EnterTrim is called when entering the trim production.
	EnterTrim(c *TrimContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterConvert is called when entering the convert production.
	EnterConvert(c *ConvertContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterSort is called when entering the sort production.
	EnterSort(c *SortContext)

	// EnterBreak_ is called when entering the break_ production.
	EnterBreak_(c *Break_Context)

	// EnterTransfer is called when entering the transfer production.
	EnterTransfer(c *TransferContext)

	// EnterTransferpre is called when entering the transferpre production.
	EnterTransferpre(c *TransferpreContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLin is called when exiting the lin production.
	ExitLin(c *LinContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitSubject is called when exiting the subject production.
	ExitSubject(c *SubjectContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMultiplyingExpression is called when exiting the multiplyingExpression production.
	ExitMultiplyingExpression(c *MultiplyingExpressionContext)

	// ExitPowExpression is called when exiting the powExpression production.
	ExitPowExpression(c *PowExpressionContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitDiffer is called when exiting the differ production.
	ExitDiffer(c *DifferContext)

	// ExitEq is called when exiting the eq production.
	ExitEq(c *EqContext)

	// ExitNe is called when exiting the ne production.
	ExitNe(c *NeContext)

	// ExitGe is called when exiting the ge production.
	ExitGe(c *GeContext)

	// ExitGt is called when exiting the gt production.
	ExitGt(c *GtContext)

	// ExitLe is called when exiting the le production.
	ExitLe(c *LeContext)

	// ExitLt is called when exiting the lt production.
	ExitLt(c *LtContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitLgt is called when exiting the lgt production.
	ExitLgt(c *LgtContext)

	// ExitAtan is called when exiting the atan production.
	ExitAtan(c *AtanContext)

	// ExitChop is called when exiting the chop production.
	ExitChop(c *ChopContext)

	// ExitCos is called when exiting the cos production.
	ExitCos(c *CosContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitLn is called when exiting the ln production.
	ExitLn(c *LnContext)

	// ExitRemdr is called when exiting the remdr production.
	ExitRemdr(c *RemdrContext)

	// ExitSin is called when exiting the sin production.
	ExitSin(c *SinContext)

	// ExitTan is called when exiting the tan production.
	ExitTan(c *TanContext)

	// ExitDupl is called when exiting the dupl production.
	ExitDupl(c *DuplContext)

	// ExitReverse is called when exiting the reverse production.
	ExitReverse(c *ReverseContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitReplace is called when exiting the replace production.
	ExitReplace(c *ReplaceContext)

	// ExitSize is called when exiting the size production.
	ExitSize(c *SizeContext)

	// ExitTrim is called when exiting the trim production.
	ExitTrim(c *TrimContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitConvert is called when exiting the convert production.
	ExitConvert(c *ConvertContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitSort is called when exiting the sort production.
	ExitSort(c *SortContext)

	// ExitBreak_ is called when exiting the break_ production.
	ExitBreak_(c *Break_Context)

	// ExitTransfer is called when exiting the transfer production.
	ExitTransfer(c *TransferContext)

	// ExitTransferpre is called when exiting the transferpre production.
	ExitTransferpre(c *TransferpreContext)
}
