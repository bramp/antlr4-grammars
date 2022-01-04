// Code generated from mumps.g4 by ANTLR 4.9.3. DO NOT EDIT.

package mumps // mumps
import "github.com/antlr/antlr4/runtime/Go/antlr"

// mumpsListener is a complete listener for a parse tree produced by mumpsParser.
type mumpsListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterEof is called when entering the eof production.
	EnterEof(c *EofContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterCode is called when entering the code production.
	EnterCode(c *CodeContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterRoutinedecl is called when entering the routinedecl production.
	EnterRoutinedecl(c *RoutinedeclContext)

	// EnterParamlist is called when entering the paramlist production.
	EnterParamlist(c *ParamlistContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterSubproc is called when entering the subproc production.
	EnterSubproc(c *SubprocContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterPostcondition is called when entering the postcondition production.
	EnterPostcondition(c *PostconditionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterFunction_ is called when entering the function_ production.
	EnterFunction_(c *Function_Context)

	// EnterBreak_ is called when entering the break_ production.
	EnterBreak_(c *Break_Context)

	// EnterDo_ is called when entering the do_ production.
	EnterDo_(c *Do_Context)

	// EnterFor_ is called when entering the for_ production.
	EnterFor_(c *For_Context)

	// EnterHalt_ is called when entering the halt_ production.
	EnterHalt_(c *Halt_Context)

	// EnterHang_ is called when entering the hang_ production.
	EnterHang_(c *Hang_Context)

	// EnterIf_ is called when entering the if_ production.
	EnterIf_(c *If_Context)

	// EnterKill_ is called when entering the kill_ production.
	EnterKill_(c *Kill_Context)

	// EnterMerge_ is called when entering the merge_ production.
	EnterMerge_(c *Merge_Context)

	// EnterNew_ is called when entering the new_ production.
	EnterNew_(c *New_Context)

	// EnterQuit_ is called when entering the quit_ production.
	EnterQuit_(c *Quit_Context)

	// EnterRead_ is called when entering the read_ production.
	EnterRead_(c *Read_Context)

	// EnterSet_ is called when entering the set_ production.
	EnterSet_(c *Set_Context)

	// EnterView_ is called when entering the view_ production.
	EnterView_(c *View_Context)

	// EnterWrite_ is called when entering the write_ production.
	EnterWrite_(c *Write_Context)

	// EnterXecute_ is called when entering the xecute_ production.
	EnterXecute_(c *Xecute_Context)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterArglist is called when entering the arglist production.
	EnterArglist(c *ArglistContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitEof is called when exiting the eof production.
	ExitEof(c *EofContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitCode is called when exiting the code production.
	ExitCode(c *CodeContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitRoutinedecl is called when exiting the routinedecl production.
	ExitRoutinedecl(c *RoutinedeclContext)

	// ExitParamlist is called when exiting the paramlist production.
	ExitParamlist(c *ParamlistContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitSubproc is called when exiting the subproc production.
	ExitSubproc(c *SubprocContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitPostcondition is called when exiting the postcondition production.
	ExitPostcondition(c *PostconditionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitFunction_ is called when exiting the function_ production.
	ExitFunction_(c *Function_Context)

	// ExitBreak_ is called when exiting the break_ production.
	ExitBreak_(c *Break_Context)

	// ExitDo_ is called when exiting the do_ production.
	ExitDo_(c *Do_Context)

	// ExitFor_ is called when exiting the for_ production.
	ExitFor_(c *For_Context)

	// ExitHalt_ is called when exiting the halt_ production.
	ExitHalt_(c *Halt_Context)

	// ExitHang_ is called when exiting the hang_ production.
	ExitHang_(c *Hang_Context)

	// ExitIf_ is called when exiting the if_ production.
	ExitIf_(c *If_Context)

	// ExitKill_ is called when exiting the kill_ production.
	ExitKill_(c *Kill_Context)

	// ExitMerge_ is called when exiting the merge_ production.
	ExitMerge_(c *Merge_Context)

	// ExitNew_ is called when exiting the new_ production.
	ExitNew_(c *New_Context)

	// ExitQuit_ is called when exiting the quit_ production.
	ExitQuit_(c *Quit_Context)

	// ExitRead_ is called when exiting the read_ production.
	ExitRead_(c *Read_Context)

	// ExitSet_ is called when exiting the set_ production.
	ExitSet_(c *Set_Context)

	// ExitView_ is called when exiting the view_ production.
	ExitView_(c *View_Context)

	// ExitWrite_ is called when exiting the write_ production.
	ExitWrite_(c *Write_Context)

	// ExitXecute_ is called when exiting the xecute_ production.
	ExitXecute_(c *Xecute_Context)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitArglist is called when exiting the arglist production.
	ExitArglist(c *ArglistContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)
}
