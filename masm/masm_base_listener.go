// Generated from MASM.g4 by ANTLR 4.7.

package masm // MASM
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMASMListener is a complete listener for a parse tree produced by MASMParser.
type BaseMASMListener struct{}

var _ MASMListener = &BaseMASMListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMASMListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMASMListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMASMListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMASMListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseMASMListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseMASMListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterSegments is called when production segments is entered.
func (s *BaseMASMListener) EnterSegments(ctx *SegmentsContext) {}

// ExitSegments is called when production segments is exited.
func (s *BaseMASMListener) ExitSegments(ctx *SegmentsContext) {}

// EnterProc is called when production proc is entered.
func (s *BaseMASMListener) EnterProc(ctx *ProcContext) {}

// ExitProc is called when production proc is exited.
func (s *BaseMASMListener) ExitProc(ctx *ProcContext) {}

// EnterCode is called when production code is entered.
func (s *BaseMASMListener) EnterCode(ctx *CodeContext) {}

// ExitCode is called when production code is exited.
func (s *BaseMASMListener) ExitCode(ctx *CodeContext) {}

// EnterBinary_exp1 is called when production binary_exp1 is entered.
func (s *BaseMASMListener) EnterBinary_exp1(ctx *Binary_exp1Context) {}

// ExitBinary_exp1 is called when production binary_exp1 is exited.
func (s *BaseMASMListener) ExitBinary_exp1(ctx *Binary_exp1Context) {}

// EnterUnuary_exp1 is called when production unuary_exp1 is entered.
func (s *BaseMASMListener) EnterUnuary_exp1(ctx *Unuary_exp1Context) {}

// ExitUnuary_exp1 is called when production unuary_exp1 is exited.
func (s *BaseMASMListener) ExitUnuary_exp1(ctx *Unuary_exp1Context) {}

// EnterUnuary_exp2 is called when production unuary_exp2 is entered.
func (s *BaseMASMListener) EnterUnuary_exp2(ctx *Unuary_exp2Context) {}

// ExitUnuary_exp2 is called when production unuary_exp2 is exited.
func (s *BaseMASMListener) ExitUnuary_exp2(ctx *Unuary_exp2Context) {}

// EnterBinary_exp2 is called when production binary_exp2 is entered.
func (s *BaseMASMListener) EnterBinary_exp2(ctx *Binary_exp2Context) {}

// ExitBinary_exp2 is called when production binary_exp2 is exited.
func (s *BaseMASMListener) ExitBinary_exp2(ctx *Binary_exp2Context) {}

// EnterNotarguments is called when production notarguments is entered.
func (s *BaseMASMListener) EnterNotarguments(ctx *NotargumentsContext) {}

// ExitNotarguments is called when production notarguments is exited.
func (s *BaseMASMListener) ExitNotarguments(ctx *NotargumentsContext) {}

// EnterBinary_exp3 is called when production binary_exp3 is entered.
func (s *BaseMASMListener) EnterBinary_exp3(ctx *Binary_exp3Context) {}

// ExitBinary_exp3 is called when production binary_exp3 is exited.
func (s *BaseMASMListener) ExitBinary_exp3(ctx *Binary_exp3Context) {}

// EnterUnuary_exp3 is called when production unuary_exp3 is entered.
func (s *BaseMASMListener) EnterUnuary_exp3(ctx *Unuary_exp3Context) {}

// ExitUnuary_exp3 is called when production unuary_exp3 is exited.
func (s *BaseMASMListener) ExitUnuary_exp3(ctx *Unuary_exp3Context) {}

// EnterBinary_exp4 is called when production binary_exp4 is entered.
func (s *BaseMASMListener) EnterBinary_exp4(ctx *Binary_exp4Context) {}

// ExitBinary_exp4 is called when production binary_exp4 is exited.
func (s *BaseMASMListener) ExitBinary_exp4(ctx *Binary_exp4Context) {}

// EnterBinary_exp5 is called when production binary_exp5 is entered.
func (s *BaseMASMListener) EnterBinary_exp5(ctx *Binary_exp5Context) {}

// ExitBinary_exp5 is called when production binary_exp5 is exited.
func (s *BaseMASMListener) ExitBinary_exp5(ctx *Binary_exp5Context) {}

// EnterBinary_exp6 is called when production binary_exp6 is entered.
func (s *BaseMASMListener) EnterBinary_exp6(ctx *Binary_exp6Context) {}

// ExitBinary_exp6 is called when production binary_exp6 is exited.
func (s *BaseMASMListener) ExitBinary_exp6(ctx *Binary_exp6Context) {}

// EnterBinary_exp7 is called when production binary_exp7 is entered.
func (s *BaseMASMListener) EnterBinary_exp7(ctx *Binary_exp7Context) {}

// ExitBinary_exp7 is called when production binary_exp7 is exited.
func (s *BaseMASMListener) ExitBinary_exp7(ctx *Binary_exp7Context) {}

// EnterBinary_exp8 is called when production binary_exp8 is entered.
func (s *BaseMASMListener) EnterBinary_exp8(ctx *Binary_exp8Context) {}

// ExitBinary_exp8 is called when production binary_exp8 is exited.
func (s *BaseMASMListener) ExitBinary_exp8(ctx *Binary_exp8Context) {}

// EnterBinary_exp9 is called when production binary_exp9 is entered.
func (s *BaseMASMListener) EnterBinary_exp9(ctx *Binary_exp9Context) {}

// ExitBinary_exp9 is called when production binary_exp9 is exited.
func (s *BaseMASMListener) ExitBinary_exp9(ctx *Binary_exp9Context) {}

// EnterUnuary_exp4 is called when production unuary_exp4 is entered.
func (s *BaseMASMListener) EnterUnuary_exp4(ctx *Unuary_exp4Context) {}

// ExitUnuary_exp4 is called when production unuary_exp4 is exited.
func (s *BaseMASMListener) ExitUnuary_exp4(ctx *Unuary_exp4Context) {}

// EnterUnuary_exp5 is called when production unuary_exp5 is entered.
func (s *BaseMASMListener) EnterUnuary_exp5(ctx *Unuary_exp5Context) {}

// ExitUnuary_exp5 is called when production unuary_exp5 is exited.
func (s *BaseMASMListener) ExitUnuary_exp5(ctx *Unuary_exp5Context) {}

// EnterBinary_exp10 is called when production binary_exp10 is entered.
func (s *BaseMASMListener) EnterBinary_exp10(ctx *Binary_exp10Context) {}

// ExitBinary_exp10 is called when production binary_exp10 is exited.
func (s *BaseMASMListener) ExitBinary_exp10(ctx *Binary_exp10Context) {}

// EnterBinary_exp11 is called when production binary_exp11 is entered.
func (s *BaseMASMListener) EnterBinary_exp11(ctx *Binary_exp11Context) {}

// ExitBinary_exp11 is called when production binary_exp11 is exited.
func (s *BaseMASMListener) ExitBinary_exp11(ctx *Binary_exp11Context) {}

// EnterBinary_exp12 is called when production binary_exp12 is entered.
func (s *BaseMASMListener) EnterBinary_exp12(ctx *Binary_exp12Context) {}

// ExitBinary_exp12 is called when production binary_exp12 is exited.
func (s *BaseMASMListener) ExitBinary_exp12(ctx *Binary_exp12Context) {}

// EnterDirective_exp1 is called when production directive_exp1 is entered.
func (s *BaseMASMListener) EnterDirective_exp1(ctx *Directive_exp1Context) {}

// ExitDirective_exp1 is called when production directive_exp1 is exited.
func (s *BaseMASMListener) ExitDirective_exp1(ctx *Directive_exp1Context) {}

// EnterVariabledeclaration is called when production variabledeclaration is entered.
func (s *BaseMASMListener) EnterVariabledeclaration(ctx *VariabledeclarationContext) {}

// ExitVariabledeclaration is called when production variabledeclaration is exited.
func (s *BaseMASMListener) ExitVariabledeclaration(ctx *VariabledeclarationContext) {}

// EnterMemory is called when production memory is entered.
func (s *BaseMASMListener) EnterMemory(ctx *MemoryContext) {}

// ExitMemory is called when production memory is exited.
func (s *BaseMASMListener) ExitMemory(ctx *MemoryContext) {}

// EnterSegmentos is called when production segmentos is entered.
func (s *BaseMASMListener) EnterSegmentos(ctx *SegmentosContext) {}

// ExitSegmentos is called when production segmentos is exited.
func (s *BaseMASMListener) ExitSegmentos(ctx *SegmentosContext) {}

// EnterRegister is called when production register is entered.
func (s *BaseMASMListener) EnterRegister(ctx *RegisterContext) {}

// ExitRegister is called when production register is exited.
func (s *BaseMASMListener) ExitRegister(ctx *RegisterContext) {}

// EnterO is called when production o is entered.
func (s *BaseMASMListener) EnterO(ctx *OContext) {}

// ExitO is called when production o is exited.
func (s *BaseMASMListener) ExitO(ctx *OContext) {}

// EnterOp is called when production op is entered.
func (s *BaseMASMListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseMASMListener) ExitOp(ctx *OpContext) {}

// EnterOpe is called when production ope is entered.
func (s *BaseMASMListener) EnterOpe(ctx *OpeContext) {}

// ExitOpe is called when production ope is exited.
func (s *BaseMASMListener) ExitOpe(ctx *OpeContext) {}

// EnterOper is called when production oper is entered.
func (s *BaseMASMListener) EnterOper(ctx *OperContext) {}

// ExitOper is called when production oper is exited.
func (s *BaseMASMListener) ExitOper(ctx *OperContext) {}

// EnterOpera is called when production opera is entered.
func (s *BaseMASMListener) EnterOpera(ctx *OperaContext) {}

// ExitOpera is called when production opera is exited.
func (s *BaseMASMListener) ExitOpera(ctx *OperaContext) {}

// EnterOperat is called when production operat is entered.
func (s *BaseMASMListener) EnterOperat(ctx *OperatContext) {}

// ExitOperat is called when production operat is exited.
func (s *BaseMASMListener) ExitOperat(ctx *OperatContext) {}

// EnterOperato is called when production operato is entered.
func (s *BaseMASMListener) EnterOperato(ctx *OperatoContext) {}

// ExitOperato is called when production operato is exited.
func (s *BaseMASMListener) ExitOperato(ctx *OperatoContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseMASMListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseMASMListener) ExitOperator(ctx *OperatorContext) {}

// EnterL is called when production l is entered.
func (s *BaseMASMListener) EnterL(ctx *LContext) {}

// ExitL is called when production l is exited.
func (s *BaseMASMListener) ExitL(ctx *LContext) {}

// EnterX is called when production x is entered.
func (s *BaseMASMListener) EnterX(ctx *XContext) {}

// ExitX is called when production x is exited.
func (s *BaseMASMListener) ExitX(ctx *XContext) {}

// EnterS is called when production s is entered.
func (s *BaseMASMListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseMASMListener) ExitS(ctx *SContext) {}

// EnterSh is called when production sh is entered.
func (s *BaseMASMListener) EnterSh(ctx *ShContext) {}

// ExitSh is called when production sh is exited.
func (s *BaseMASMListener) ExitSh(ctx *ShContext) {}

// EnterB is called when production b is entered.
func (s *BaseMASMListener) EnterB(ctx *BContext) {}

// ExitB is called when production b is exited.
func (s *BaseMASMListener) ExitB(ctx *BContext) {}

// EnterCall is called when production call is entered.
func (s *BaseMASMListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseMASMListener) ExitCall(ctx *CallContext) {}

// EnterInterruption is called when production interruption is entered.
func (s *BaseMASMListener) EnterInterruption(ctx *InterruptionContext) {}

// ExitInterruption is called when production interruption is exited.
func (s *BaseMASMListener) ExitInterruption(ctx *InterruptionContext) {}

// EnterIn is called when production in is entered.
func (s *BaseMASMListener) EnterIn(ctx *InContext) {}

// ExitIn is called when production in is exited.
func (s *BaseMASMListener) ExitIn(ctx *InContext) {}

// EnterOut is called when production out is entered.
func (s *BaseMASMListener) EnterOut(ctx *OutContext) {}

// ExitOut is called when production out is exited.
func (s *BaseMASMListener) ExitOut(ctx *OutContext) {}

// EnterRe is called when production re is entered.
func (s *BaseMASMListener) EnterRe(ctx *ReContext) {}

// ExitRe is called when production re is exited.
func (s *BaseMASMListener) ExitRe(ctx *ReContext) {}

// EnterDirectives is called when production directives is entered.
func (s *BaseMASMListener) EnterDirectives(ctx *DirectivesContext) {}

// ExitDirectives is called when production directives is exited.
func (s *BaseMASMListener) ExitDirectives(ctx *DirectivesContext) {}

// EnterTy is called when production ty is entered.
func (s *BaseMASMListener) EnterTy(ctx *TyContext) {}

// ExitTy is called when production ty is exited.
func (s *BaseMASMListener) ExitTy(ctx *TyContext) {}

// EnterQuestion is called when production question is entered.
func (s *BaseMASMListener) EnterQuestion(ctx *QuestionContext) {}

// ExitQuestion is called when production question is exited.
func (s *BaseMASMListener) ExitQuestion(ctx *QuestionContext) {}

// EnterTime is called when production time is entered.
func (s *BaseMASMListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BaseMASMListener) ExitTime(ctx *TimeContext) {}
