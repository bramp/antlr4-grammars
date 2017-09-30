// Generated from MASM.g4 by ANTLR 4.7.

package masm // MASM
import "github.com/antlr/antlr4/runtime/Go/antlr"

 
 	 package com.Ostermiller.Syntax;
 	 

// A complete Visitor for a parse tree produced by MASMParser.
type MASMVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MASMParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by MASMParser#segments.
	VisitSegments(ctx *SegmentsContext) interface{}

	// Visit a parse tree produced by MASMParser#proc.
	VisitProc(ctx *ProcContext) interface{}

	// Visit a parse tree produced by MASMParser#code.
	VisitCode(ctx *CodeContext) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp1.
	VisitBinary_exp1(ctx *Binary_exp1Context) interface{}

	// Visit a parse tree produced by MASMParser#unuary_exp1.
	VisitUnuary_exp1(ctx *Unuary_exp1Context) interface{}

	// Visit a parse tree produced by MASMParser#unuary_exp2.
	VisitUnuary_exp2(ctx *Unuary_exp2Context) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp2.
	VisitBinary_exp2(ctx *Binary_exp2Context) interface{}

	// Visit a parse tree produced by MASMParser#notarguments.
	VisitNotarguments(ctx *NotargumentsContext) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp3.
	VisitBinary_exp3(ctx *Binary_exp3Context) interface{}

	// Visit a parse tree produced by MASMParser#unuary_exp3.
	VisitUnuary_exp3(ctx *Unuary_exp3Context) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp4.
	VisitBinary_exp4(ctx *Binary_exp4Context) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp5.
	VisitBinary_exp5(ctx *Binary_exp5Context) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp6.
	VisitBinary_exp6(ctx *Binary_exp6Context) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp7.
	VisitBinary_exp7(ctx *Binary_exp7Context) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp8.
	VisitBinary_exp8(ctx *Binary_exp8Context) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp9.
	VisitBinary_exp9(ctx *Binary_exp9Context) interface{}

	// Visit a parse tree produced by MASMParser#unuary_exp4.
	VisitUnuary_exp4(ctx *Unuary_exp4Context) interface{}

	// Visit a parse tree produced by MASMParser#unuary_exp5.
	VisitUnuary_exp5(ctx *Unuary_exp5Context) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp10.
	VisitBinary_exp10(ctx *Binary_exp10Context) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp11.
	VisitBinary_exp11(ctx *Binary_exp11Context) interface{}

	// Visit a parse tree produced by MASMParser#binary_exp12.
	VisitBinary_exp12(ctx *Binary_exp12Context) interface{}

	// Visit a parse tree produced by MASMParser#directive_exp1.
	VisitDirective_exp1(ctx *Directive_exp1Context) interface{}

	// Visit a parse tree produced by MASMParser#variabledeclaration.
	VisitVariabledeclaration(ctx *VariabledeclarationContext) interface{}

	// Visit a parse tree produced by MASMParser#memory.
	VisitMemory(ctx *MemoryContext) interface{}

	// Visit a parse tree produced by MASMParser#segmentos.
	VisitSegmentos(ctx *SegmentosContext) interface{}

	// Visit a parse tree produced by MASMParser#register.
	VisitRegister(ctx *RegisterContext) interface{}

	// Visit a parse tree produced by MASMParser#o.
	VisitO(ctx *OContext) interface{}

	// Visit a parse tree produced by MASMParser#op.
	VisitOp(ctx *OpContext) interface{}

	// Visit a parse tree produced by MASMParser#ope.
	VisitOpe(ctx *OpeContext) interface{}

	// Visit a parse tree produced by MASMParser#oper.
	VisitOper(ctx *OperContext) interface{}

	// Visit a parse tree produced by MASMParser#opera.
	VisitOpera(ctx *OperaContext) interface{}

	// Visit a parse tree produced by MASMParser#operat.
	VisitOperat(ctx *OperatContext) interface{}

	// Visit a parse tree produced by MASMParser#operato.
	VisitOperato(ctx *OperatoContext) interface{}

	// Visit a parse tree produced by MASMParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}

	// Visit a parse tree produced by MASMParser#l.
	VisitL(ctx *LContext) interface{}

	// Visit a parse tree produced by MASMParser#x.
	VisitX(ctx *XContext) interface{}

	// Visit a parse tree produced by MASMParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by MASMParser#sh.
	VisitSh(ctx *ShContext) interface{}

	// Visit a parse tree produced by MASMParser#b.
	VisitB(ctx *BContext) interface{}

	// Visit a parse tree produced by MASMParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by MASMParser#interruption.
	VisitInterruption(ctx *InterruptionContext) interface{}

	// Visit a parse tree produced by MASMParser#in.
	VisitIn(ctx *InContext) interface{}

	// Visit a parse tree produced by MASMParser#out.
	VisitOut(ctx *OutContext) interface{}

	// Visit a parse tree produced by MASMParser#re.
	VisitRe(ctx *ReContext) interface{}

	// Visit a parse tree produced by MASMParser#directives.
	VisitDirectives(ctx *DirectivesContext) interface{}

	// Visit a parse tree produced by MASMParser#ty.
	VisitTy(ctx *TyContext) interface{}

	// Visit a parse tree produced by MASMParser#question.
	VisitQuestion(ctx *QuestionContext) interface{}

	// Visit a parse tree produced by MASMParser#time.
	VisitTime(ctx *TimeContext) interface{}

}