// Code generated from mumps.g4 by ANTLR 4.9.3. DO NOT EDIT.

package mumps // mumps
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasemumpsListener is a complete listener for a parse tree produced by mumpsParser.
type BasemumpsListener struct{}

var _ mumpsListener = &BasemumpsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemumpsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemumpsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemumpsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemumpsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasemumpsListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasemumpsListener) ExitProgram(ctx *ProgramContext) {}

// EnterEof is called when production eof is entered.
func (s *BasemumpsListener) EnterEof(ctx *EofContext) {}

// ExitEof is called when production eof is exited.
func (s *BasemumpsListener) ExitEof(ctx *EofContext) {}

// EnterLine is called when production line is entered.
func (s *BasemumpsListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BasemumpsListener) ExitLine(ctx *LineContext) {}

// EnterCode is called when production code is entered.
func (s *BasemumpsListener) EnterCode(ctx *CodeContext) {}

// ExitCode is called when production code is exited.
func (s *BasemumpsListener) ExitCode(ctx *CodeContext) {}

// EnterLabel is called when production label is entered.
func (s *BasemumpsListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BasemumpsListener) ExitLabel(ctx *LabelContext) {}

// EnterRoutinedecl is called when production routinedecl is entered.
func (s *BasemumpsListener) EnterRoutinedecl(ctx *RoutinedeclContext) {}

// ExitRoutinedecl is called when production routinedecl is exited.
func (s *BasemumpsListener) ExitRoutinedecl(ctx *RoutinedeclContext) {}

// EnterParamlist is called when production paramlist is entered.
func (s *BasemumpsListener) EnterParamlist(ctx *ParamlistContext) {}

// ExitParamlist is called when production paramlist is exited.
func (s *BasemumpsListener) ExitParamlist(ctx *ParamlistContext) {}

// EnterParam is called when production param is entered.
func (s *BasemumpsListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasemumpsListener) ExitParam(ctx *ParamContext) {}

// EnterSubproc is called when production subproc is entered.
func (s *BasemumpsListener) EnterSubproc(ctx *SubprocContext) {}

// ExitSubproc is called when production subproc is exited.
func (s *BasemumpsListener) ExitSubproc(ctx *SubprocContext) {}

// EnterCommand is called when production command is entered.
func (s *BasemumpsListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BasemumpsListener) ExitCommand(ctx *CommandContext) {}

// EnterPostcondition is called when production postcondition is entered.
func (s *BasemumpsListener) EnterPostcondition(ctx *PostconditionContext) {}

// ExitPostcondition is called when production postcondition is exited.
func (s *BasemumpsListener) ExitPostcondition(ctx *PostconditionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasemumpsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasemumpsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BasemumpsListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasemumpsListener) ExitTerm(ctx *TermContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasemumpsListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasemumpsListener) ExitCondition(ctx *ConditionContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasemumpsListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasemumpsListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterVariable is called when production variable is entered.
func (s *BasemumpsListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BasemumpsListener) ExitVariable(ctx *VariableContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BasemumpsListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BasemumpsListener) ExitFunction_(ctx *Function_Context) {}

// EnterBreak_ is called when production break_ is entered.
func (s *BasemumpsListener) EnterBreak_(ctx *Break_Context) {}

// ExitBreak_ is called when production break_ is exited.
func (s *BasemumpsListener) ExitBreak_(ctx *Break_Context) {}

// EnterDo_ is called when production do_ is entered.
func (s *BasemumpsListener) EnterDo_(ctx *Do_Context) {}

// ExitDo_ is called when production do_ is exited.
func (s *BasemumpsListener) ExitDo_(ctx *Do_Context) {}

// EnterFor_ is called when production for_ is entered.
func (s *BasemumpsListener) EnterFor_(ctx *For_Context) {}

// ExitFor_ is called when production for_ is exited.
func (s *BasemumpsListener) ExitFor_(ctx *For_Context) {}

// EnterHalt_ is called when production halt_ is entered.
func (s *BasemumpsListener) EnterHalt_(ctx *Halt_Context) {}

// ExitHalt_ is called when production halt_ is exited.
func (s *BasemumpsListener) ExitHalt_(ctx *Halt_Context) {}

// EnterHang_ is called when production hang_ is entered.
func (s *BasemumpsListener) EnterHang_(ctx *Hang_Context) {}

// ExitHang_ is called when production hang_ is exited.
func (s *BasemumpsListener) ExitHang_(ctx *Hang_Context) {}

// EnterIf_ is called when production if_ is entered.
func (s *BasemumpsListener) EnterIf_(ctx *If_Context) {}

// ExitIf_ is called when production if_ is exited.
func (s *BasemumpsListener) ExitIf_(ctx *If_Context) {}

// EnterKill_ is called when production kill_ is entered.
func (s *BasemumpsListener) EnterKill_(ctx *Kill_Context) {}

// ExitKill_ is called when production kill_ is exited.
func (s *BasemumpsListener) ExitKill_(ctx *Kill_Context) {}

// EnterMerge_ is called when production merge_ is entered.
func (s *BasemumpsListener) EnterMerge_(ctx *Merge_Context) {}

// ExitMerge_ is called when production merge_ is exited.
func (s *BasemumpsListener) ExitMerge_(ctx *Merge_Context) {}

// EnterNew_ is called when production new_ is entered.
func (s *BasemumpsListener) EnterNew_(ctx *New_Context) {}

// ExitNew_ is called when production new_ is exited.
func (s *BasemumpsListener) ExitNew_(ctx *New_Context) {}

// EnterQuit_ is called when production quit_ is entered.
func (s *BasemumpsListener) EnterQuit_(ctx *Quit_Context) {}

// ExitQuit_ is called when production quit_ is exited.
func (s *BasemumpsListener) ExitQuit_(ctx *Quit_Context) {}

// EnterRead_ is called when production read_ is entered.
func (s *BasemumpsListener) EnterRead_(ctx *Read_Context) {}

// ExitRead_ is called when production read_ is exited.
func (s *BasemumpsListener) ExitRead_(ctx *Read_Context) {}

// EnterSet_ is called when production set_ is entered.
func (s *BasemumpsListener) EnterSet_(ctx *Set_Context) {}

// ExitSet_ is called when production set_ is exited.
func (s *BasemumpsListener) ExitSet_(ctx *Set_Context) {}

// EnterView_ is called when production view_ is entered.
func (s *BasemumpsListener) EnterView_(ctx *View_Context) {}

// ExitView_ is called when production view_ is exited.
func (s *BasemumpsListener) ExitView_(ctx *View_Context) {}

// EnterWrite_ is called when production write_ is entered.
func (s *BasemumpsListener) EnterWrite_(ctx *Write_Context) {}

// ExitWrite_ is called when production write_ is exited.
func (s *BasemumpsListener) ExitWrite_(ctx *Write_Context) {}

// EnterXecute_ is called when production xecute_ is entered.
func (s *BasemumpsListener) EnterXecute_(ctx *Xecute_Context) {}

// ExitXecute_ is called when production xecute_ is exited.
func (s *BasemumpsListener) ExitXecute_(ctx *Xecute_Context) {}

// EnterAssign is called when production assign is entered.
func (s *BasemumpsListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasemumpsListener) ExitAssign(ctx *AssignContext) {}

// EnterArglist is called when production arglist is entered.
func (s *BasemumpsListener) EnterArglist(ctx *ArglistContext) {}

// ExitArglist is called when production arglist is exited.
func (s *BasemumpsListener) ExitArglist(ctx *ArglistContext) {}

// EnterArg is called when production arg is entered.
func (s *BasemumpsListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BasemumpsListener) ExitArg(ctx *ArgContext) {}
