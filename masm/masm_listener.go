// Generated from MASM.g4 by ANTLR 4.7.

package masm // MASM
import "github.com/antlr/antlr4/runtime/Go/antlr"

// MASMListener is a complete listener for a parse tree produced by MASMParser.
type MASMListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterSegments is called when entering the segments production.
	EnterSegments(c *SegmentsContext)

	// EnterProc is called when entering the proc production.
	EnterProc(c *ProcContext)

	// EnterCode is called when entering the code production.
	EnterCode(c *CodeContext)

	// EnterBinary_exp1 is called when entering the binary_exp1 production.
	EnterBinary_exp1(c *Binary_exp1Context)

	// EnterUnuary_exp1 is called when entering the unuary_exp1 production.
	EnterUnuary_exp1(c *Unuary_exp1Context)

	// EnterUnuary_exp2 is called when entering the unuary_exp2 production.
	EnterUnuary_exp2(c *Unuary_exp2Context)

	// EnterBinary_exp2 is called when entering the binary_exp2 production.
	EnterBinary_exp2(c *Binary_exp2Context)

	// EnterNotarguments is called when entering the notarguments production.
	EnterNotarguments(c *NotargumentsContext)

	// EnterBinary_exp3 is called when entering the binary_exp3 production.
	EnterBinary_exp3(c *Binary_exp3Context)

	// EnterUnuary_exp3 is called when entering the unuary_exp3 production.
	EnterUnuary_exp3(c *Unuary_exp3Context)

	// EnterBinary_exp4 is called when entering the binary_exp4 production.
	EnterBinary_exp4(c *Binary_exp4Context)

	// EnterBinary_exp5 is called when entering the binary_exp5 production.
	EnterBinary_exp5(c *Binary_exp5Context)

	// EnterBinary_exp6 is called when entering the binary_exp6 production.
	EnterBinary_exp6(c *Binary_exp6Context)

	// EnterBinary_exp7 is called when entering the binary_exp7 production.
	EnterBinary_exp7(c *Binary_exp7Context)

	// EnterBinary_exp8 is called when entering the binary_exp8 production.
	EnterBinary_exp8(c *Binary_exp8Context)

	// EnterBinary_exp9 is called when entering the binary_exp9 production.
	EnterBinary_exp9(c *Binary_exp9Context)

	// EnterUnuary_exp4 is called when entering the unuary_exp4 production.
	EnterUnuary_exp4(c *Unuary_exp4Context)

	// EnterUnuary_exp5 is called when entering the unuary_exp5 production.
	EnterUnuary_exp5(c *Unuary_exp5Context)

	// EnterBinary_exp10 is called when entering the binary_exp10 production.
	EnterBinary_exp10(c *Binary_exp10Context)

	// EnterBinary_exp11 is called when entering the binary_exp11 production.
	EnterBinary_exp11(c *Binary_exp11Context)

	// EnterBinary_exp12 is called when entering the binary_exp12 production.
	EnterBinary_exp12(c *Binary_exp12Context)

	// EnterDirective_exp1 is called when entering the directive_exp1 production.
	EnterDirective_exp1(c *Directive_exp1Context)

	// EnterVariabledeclaration is called when entering the variabledeclaration production.
	EnterVariabledeclaration(c *VariabledeclarationContext)

	// EnterMemory is called when entering the memory production.
	EnterMemory(c *MemoryContext)

	// EnterSegmentos is called when entering the segmentos production.
	EnterSegmentos(c *SegmentosContext)

	// EnterRegister is called when entering the register production.
	EnterRegister(c *RegisterContext)

	// EnterO is called when entering the o production.
	EnterO(c *OContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// EnterOpe is called when entering the ope production.
	EnterOpe(c *OpeContext)

	// EnterOper is called when entering the oper production.
	EnterOper(c *OperContext)

	// EnterOpera is called when entering the opera production.
	EnterOpera(c *OperaContext)

	// EnterOperat is called when entering the operat production.
	EnterOperat(c *OperatContext)

	// EnterOperato is called when entering the operato production.
	EnterOperato(c *OperatoContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterL is called when entering the l production.
	EnterL(c *LContext)

	// EnterX is called when entering the x production.
	EnterX(c *XContext)

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterSh is called when entering the sh production.
	EnterSh(c *ShContext)

	// EnterB is called when entering the b production.
	EnterB(c *BContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterInterruption is called when entering the interruption production.
	EnterInterruption(c *InterruptionContext)

	// EnterIn is called when entering the in production.
	EnterIn(c *InContext)

	// EnterOut is called when entering the out production.
	EnterOut(c *OutContext)

	// EnterRe is called when entering the re production.
	EnterRe(c *ReContext)

	// EnterDirectives is called when entering the directives production.
	EnterDirectives(c *DirectivesContext)

	// EnterTy is called when entering the ty production.
	EnterTy(c *TyContext)

	// EnterQuestion is called when entering the question production.
	EnterQuestion(c *QuestionContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitSegments is called when exiting the segments production.
	ExitSegments(c *SegmentsContext)

	// ExitProc is called when exiting the proc production.
	ExitProc(c *ProcContext)

	// ExitCode is called when exiting the code production.
	ExitCode(c *CodeContext)

	// ExitBinary_exp1 is called when exiting the binary_exp1 production.
	ExitBinary_exp1(c *Binary_exp1Context)

	// ExitUnuary_exp1 is called when exiting the unuary_exp1 production.
	ExitUnuary_exp1(c *Unuary_exp1Context)

	// ExitUnuary_exp2 is called when exiting the unuary_exp2 production.
	ExitUnuary_exp2(c *Unuary_exp2Context)

	// ExitBinary_exp2 is called when exiting the binary_exp2 production.
	ExitBinary_exp2(c *Binary_exp2Context)

	// ExitNotarguments is called when exiting the notarguments production.
	ExitNotarguments(c *NotargumentsContext)

	// ExitBinary_exp3 is called when exiting the binary_exp3 production.
	ExitBinary_exp3(c *Binary_exp3Context)

	// ExitUnuary_exp3 is called when exiting the unuary_exp3 production.
	ExitUnuary_exp3(c *Unuary_exp3Context)

	// ExitBinary_exp4 is called when exiting the binary_exp4 production.
	ExitBinary_exp4(c *Binary_exp4Context)

	// ExitBinary_exp5 is called when exiting the binary_exp5 production.
	ExitBinary_exp5(c *Binary_exp5Context)

	// ExitBinary_exp6 is called when exiting the binary_exp6 production.
	ExitBinary_exp6(c *Binary_exp6Context)

	// ExitBinary_exp7 is called when exiting the binary_exp7 production.
	ExitBinary_exp7(c *Binary_exp7Context)

	// ExitBinary_exp8 is called when exiting the binary_exp8 production.
	ExitBinary_exp8(c *Binary_exp8Context)

	// ExitBinary_exp9 is called when exiting the binary_exp9 production.
	ExitBinary_exp9(c *Binary_exp9Context)

	// ExitUnuary_exp4 is called when exiting the unuary_exp4 production.
	ExitUnuary_exp4(c *Unuary_exp4Context)

	// ExitUnuary_exp5 is called when exiting the unuary_exp5 production.
	ExitUnuary_exp5(c *Unuary_exp5Context)

	// ExitBinary_exp10 is called when exiting the binary_exp10 production.
	ExitBinary_exp10(c *Binary_exp10Context)

	// ExitBinary_exp11 is called when exiting the binary_exp11 production.
	ExitBinary_exp11(c *Binary_exp11Context)

	// ExitBinary_exp12 is called when exiting the binary_exp12 production.
	ExitBinary_exp12(c *Binary_exp12Context)

	// ExitDirective_exp1 is called when exiting the directive_exp1 production.
	ExitDirective_exp1(c *Directive_exp1Context)

	// ExitVariabledeclaration is called when exiting the variabledeclaration production.
	ExitVariabledeclaration(c *VariabledeclarationContext)

	// ExitMemory is called when exiting the memory production.
	ExitMemory(c *MemoryContext)

	// ExitSegmentos is called when exiting the segmentos production.
	ExitSegmentos(c *SegmentosContext)

	// ExitRegister is called when exiting the register production.
	ExitRegister(c *RegisterContext)

	// ExitO is called when exiting the o production.
	ExitO(c *OContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)

	// ExitOpe is called when exiting the ope production.
	ExitOpe(c *OpeContext)

	// ExitOper is called when exiting the oper production.
	ExitOper(c *OperContext)

	// ExitOpera is called when exiting the opera production.
	ExitOpera(c *OperaContext)

	// ExitOperat is called when exiting the operat production.
	ExitOperat(c *OperatContext)

	// ExitOperato is called when exiting the operato production.
	ExitOperato(c *OperatoContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitL is called when exiting the l production.
	ExitL(c *LContext)

	// ExitX is called when exiting the x production.
	ExitX(c *XContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitSh is called when exiting the sh production.
	ExitSh(c *ShContext)

	// ExitB is called when exiting the b production.
	ExitB(c *BContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitInterruption is called when exiting the interruption production.
	ExitInterruption(c *InterruptionContext)

	// ExitIn is called when exiting the in production.
	ExitIn(c *InContext)

	// ExitOut is called when exiting the out production.
	ExitOut(c *OutContext)

	// ExitRe is called when exiting the re production.
	ExitRe(c *ReContext)

	// ExitDirectives is called when exiting the directives production.
	ExitDirectives(c *DirectivesContext)

	// ExitTy is called when exiting the ty production.
	ExitTy(c *TyContext)

	// ExitQuestion is called when exiting the question production.
	ExitQuestion(c *QuestionContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)
}
