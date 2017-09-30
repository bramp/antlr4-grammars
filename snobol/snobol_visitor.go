// Generated from snobol.g4 by ANTLR 4.7.

package snobol // snobol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by snobolParser.
type snobolVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by snobolParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by snobolParser#lin.
	VisitLin(ctx *LinContext) interface{}

	// Visit a parse tree produced by snobolParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by snobolParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by snobolParser#subject.
	VisitSubject(ctx *SubjectContext) interface{}

	// Visit a parse tree produced by snobolParser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by snobolParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by snobolParser#multiplyingExpression.
	VisitMultiplyingExpression(ctx *MultiplyingExpressionContext) interface{}

	// Visit a parse tree produced by snobolParser#powExpression.
	VisitPowExpression(ctx *PowExpressionContext) interface{}

	// Visit a parse tree produced by snobolParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by snobolParser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by snobolParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by snobolParser#differ.
	VisitDiffer(ctx *DifferContext) interface{}

	// Visit a parse tree produced by snobolParser#eq.
	VisitEq(ctx *EqContext) interface{}

	// Visit a parse tree produced by snobolParser#ne.
	VisitNe(ctx *NeContext) interface{}

	// Visit a parse tree produced by snobolParser#ge.
	VisitGe(ctx *GeContext) interface{}

	// Visit a parse tree produced by snobolParser#gt.
	VisitGt(ctx *GtContext) interface{}

	// Visit a parse tree produced by snobolParser#le.
	VisitLe(ctx *LeContext) interface{}

	// Visit a parse tree produced by snobolParser#lt.
	VisitLt(ctx *LtContext) interface{}

	// Visit a parse tree produced by snobolParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by snobolParser#lgt.
	VisitLgt(ctx *LgtContext) interface{}

	// Visit a parse tree produced by snobolParser#atan.
	VisitAtan(ctx *AtanContext) interface{}

	// Visit a parse tree produced by snobolParser#chop.
	VisitChop(ctx *ChopContext) interface{}

	// Visit a parse tree produced by snobolParser#cos.
	VisitCos(ctx *CosContext) interface{}

	// Visit a parse tree produced by snobolParser#exp.
	VisitExp(ctx *ExpContext) interface{}

	// Visit a parse tree produced by snobolParser#ln.
	VisitLn(ctx *LnContext) interface{}

	// Visit a parse tree produced by snobolParser#remdr.
	VisitRemdr(ctx *RemdrContext) interface{}

	// Visit a parse tree produced by snobolParser#sin.
	VisitSin(ctx *SinContext) interface{}

	// Visit a parse tree produced by snobolParser#tan.
	VisitTan(ctx *TanContext) interface{}

	// Visit a parse tree produced by snobolParser#dupl.
	VisitDupl(ctx *DuplContext) interface{}

	// Visit a parse tree produced by snobolParser#reverse.
	VisitReverse(ctx *ReverseContext) interface{}

	// Visit a parse tree produced by snobolParser#date.
	VisitDate(ctx *DateContext) interface{}

	// Visit a parse tree produced by snobolParser#replace.
	VisitReplace(ctx *ReplaceContext) interface{}

	// Visit a parse tree produced by snobolParser#size.
	VisitSize(ctx *SizeContext) interface{}

	// Visit a parse tree produced by snobolParser#trim.
	VisitTrim(ctx *TrimContext) interface{}

	// Visit a parse tree produced by snobolParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by snobolParser#convert.
	VisitConvert(ctx *ConvertContext) interface{}

	// Visit a parse tree produced by snobolParser#table.
	VisitTable(ctx *TableContext) interface{}

	// Visit a parse tree produced by snobolParser#sort.
	VisitSort(ctx *SortContext) interface{}

	// Visit a parse tree produced by snobolParser#break_.
	VisitBreak_(ctx *Break_Context) interface{}

	// Visit a parse tree produced by snobolParser#transfer.
	VisitTransfer(ctx *TransferContext) interface{}

	// Visit a parse tree produced by snobolParser#transferpre.
	VisitTransferpre(ctx *TransferpreContext) interface{}
}
