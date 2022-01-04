// Code generated from angelscript.g4 by ANTLR 4.9.3. DO NOT EDIT.

package angelscript // angelscript
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseangelscriptListener is a complete listener for a parse tree produced by angelscriptParser.
type BaseangelscriptListener struct{}

var _ angelscriptListener = &BaseangelscriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseangelscriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseangelscriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseangelscriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseangelscriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseangelscriptListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseangelscriptListener) ExitScript(ctx *ScriptContext) {}

// EnterClass_ is called when production class_ is entered.
func (s *BaseangelscriptListener) EnterClass_(ctx *Class_Context) {}

// ExitClass_ is called when production class_ is exited.
func (s *BaseangelscriptListener) ExitClass_(ctx *Class_Context) {}

// EnterTypdef is called when production typdef is entered.
func (s *BaseangelscriptListener) EnterTypdef(ctx *TypdefContext) {}

// ExitTypdef is called when production typdef is exited.
func (s *BaseangelscriptListener) ExitTypdef(ctx *TypdefContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseangelscriptListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseangelscriptListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterFunc_ is called when production func_ is entered.
func (s *BaseangelscriptListener) EnterFunc_(ctx *Func_Context) {}

// ExitFunc_ is called when production func_ is exited.
func (s *BaseangelscriptListener) ExitFunc_(ctx *Func_Context) {}

// EnterInterface_ is called when production interface_ is entered.
func (s *BaseangelscriptListener) EnterInterface_(ctx *Interface_Context) {}

// ExitInterface_ is called when production interface_ is exited.
func (s *BaseangelscriptListener) ExitInterface_(ctx *Interface_Context) {}

// EnterVar_ is called when production var_ is entered.
func (s *BaseangelscriptListener) EnterVar_(ctx *Var_Context) {}

// ExitVar_ is called when production var_ is exited.
func (s *BaseangelscriptListener) ExitVar_(ctx *Var_Context) {}

// EnterImport_ is called when production import_ is entered.
func (s *BaseangelscriptListener) EnterImport_(ctx *Import_Context) {}

// ExitImport_ is called when production import_ is exited.
func (s *BaseangelscriptListener) ExitImport_(ctx *Import_Context) {}

// EnterEnum_ is called when production enum_ is entered.
func (s *BaseangelscriptListener) EnterEnum_(ctx *Enum_Context) {}

// ExitEnum_ is called when production enum_ is exited.
func (s *BaseangelscriptListener) ExitEnum_(ctx *Enum_Context) {}

// EnterFuncdef is called when production funcdef is entered.
func (s *BaseangelscriptListener) EnterFuncdef(ctx *FuncdefContext) {}

// ExitFuncdef is called when production funcdef is exited.
func (s *BaseangelscriptListener) ExitFuncdef(ctx *FuncdefContext) {}

// EnterVirtprop is called when production virtprop is entered.
func (s *BaseangelscriptListener) EnterVirtprop(ctx *VirtpropContext) {}

// ExitVirtprop is called when production virtprop is exited.
func (s *BaseangelscriptListener) ExitVirtprop(ctx *VirtpropContext) {}

// EnterMixin_ is called when production mixin_ is entered.
func (s *BaseangelscriptListener) EnterMixin_(ctx *Mixin_Context) {}

// ExitMixin_ is called when production mixin_ is exited.
func (s *BaseangelscriptListener) ExitMixin_(ctx *Mixin_Context) {}

// EnterIntfmthd is called when production intfmthd is entered.
func (s *BaseangelscriptListener) EnterIntfmthd(ctx *IntfmthdContext) {}

// ExitIntfmthd is called when production intfmthd is exited.
func (s *BaseangelscriptListener) ExitIntfmthd(ctx *IntfmthdContext) {}

// EnterStatblock is called when production statblock is entered.
func (s *BaseangelscriptListener) EnterStatblock(ctx *StatblockContext) {}

// ExitStatblock is called when production statblock is exited.
func (s *BaseangelscriptListener) ExitStatblock(ctx *StatblockContext) {}

// EnterParamlist is called when production paramlist is entered.
func (s *BaseangelscriptListener) EnterParamlist(ctx *ParamlistContext) {}

// ExitParamlist is called when production paramlist is exited.
func (s *BaseangelscriptListener) ExitParamlist(ctx *ParamlistContext) {}

// EnterTypemod is called when production typemod is entered.
func (s *BaseangelscriptListener) EnterTypemod(ctx *TypemodContext) {}

// ExitTypemod is called when production typemod is exited.
func (s *BaseangelscriptListener) ExitTypemod(ctx *TypemodContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseangelscriptListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseangelscriptListener) ExitType_(ctx *Type_Context) {}

// EnterInitlist is called when production initlist is entered.
func (s *BaseangelscriptListener) EnterInitlist(ctx *InitlistContext) {}

// ExitInitlist is called when production initlist is exited.
func (s *BaseangelscriptListener) ExitInitlist(ctx *InitlistContext) {}

// EnterScope is called when production scope is entered.
func (s *BaseangelscriptListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BaseangelscriptListener) ExitScope(ctx *ScopeContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BaseangelscriptListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BaseangelscriptListener) ExitDatatype(ctx *DatatypeContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseangelscriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseangelscriptListener) ExitStatement(ctx *StatementContext) {}

// EnterSwitch_ is called when production switch_ is entered.
func (s *BaseangelscriptListener) EnterSwitch_(ctx *Switch_Context) {}

// ExitSwitch_ is called when production switch_ is exited.
func (s *BaseangelscriptListener) ExitSwitch_(ctx *Switch_Context) {}

// EnterBreak_ is called when production break_ is entered.
func (s *BaseangelscriptListener) EnterBreak_(ctx *Break_Context) {}

// ExitBreak_ is called when production break_ is exited.
func (s *BaseangelscriptListener) ExitBreak_(ctx *Break_Context) {}

// EnterFor_ is called when production for_ is entered.
func (s *BaseangelscriptListener) EnterFor_(ctx *For_Context) {}

// ExitFor_ is called when production for_ is exited.
func (s *BaseangelscriptListener) ExitFor_(ctx *For_Context) {}

// EnterWhile_ is called when production while_ is entered.
func (s *BaseangelscriptListener) EnterWhile_(ctx *While_Context) {}

// ExitWhile_ is called when production while_ is exited.
func (s *BaseangelscriptListener) ExitWhile_(ctx *While_Context) {}

// EnterDowhile is called when production dowhile is entered.
func (s *BaseangelscriptListener) EnterDowhile(ctx *DowhileContext) {}

// ExitDowhile is called when production dowhile is exited.
func (s *BaseangelscriptListener) ExitDowhile(ctx *DowhileContext) {}

// EnterIf_ is called when production if_ is entered.
func (s *BaseangelscriptListener) EnterIf_(ctx *If_Context) {}

// ExitIf_ is called when production if_ is exited.
func (s *BaseangelscriptListener) ExitIf_(ctx *If_Context) {}

// EnterContinue_ is called when production continue_ is entered.
func (s *BaseangelscriptListener) EnterContinue_(ctx *Continue_Context) {}

// ExitContinue_ is called when production continue_ is exited.
func (s *BaseangelscriptListener) ExitContinue_(ctx *Continue_Context) {}

// EnterExprstat is called when production exprstat is entered.
func (s *BaseangelscriptListener) EnterExprstat(ctx *ExprstatContext) {}

// ExitExprstat is called when production exprstat is exited.
func (s *BaseangelscriptListener) ExitExprstat(ctx *ExprstatContext) {}

// EnterTry_ is called when production try_ is entered.
func (s *BaseangelscriptListener) EnterTry_(ctx *Try_Context) {}

// ExitTry_ is called when production try_ is exited.
func (s *BaseangelscriptListener) ExitTry_(ctx *Try_Context) {}

// EnterReturn_ is called when production return_ is entered.
func (s *BaseangelscriptListener) EnterReturn_(ctx *Return_Context) {}

// ExitReturn_ is called when production return_ is exited.
func (s *BaseangelscriptListener) ExitReturn_(ctx *Return_Context) {}

// EnterCase_ is called when production case_ is entered.
func (s *BaseangelscriptListener) EnterCase_(ctx *Case_Context) {}

// ExitCase_ is called when production case_ is exited.
func (s *BaseangelscriptListener) ExitCase_(ctx *Case_Context) {}

// EnterExpr is called when production expr is entered.
func (s *BaseangelscriptListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseangelscriptListener) ExitExpr(ctx *ExprContext) {}

// EnterExprterm is called when production exprterm is entered.
func (s *BaseangelscriptListener) EnterExprterm(ctx *ExprtermContext) {}

// ExitExprterm is called when production exprterm is exited.
func (s *BaseangelscriptListener) ExitExprterm(ctx *ExprtermContext) {}

// EnterExprvalue is called when production exprvalue is entered.
func (s *BaseangelscriptListener) EnterExprvalue(ctx *ExprvalueContext) {}

// ExitExprvalue is called when production exprvalue is exited.
func (s *BaseangelscriptListener) ExitExprvalue(ctx *ExprvalueContext) {}

// EnterConstructcall is called when production constructcall is entered.
func (s *BaseangelscriptListener) EnterConstructcall(ctx *ConstructcallContext) {}

// ExitConstructcall is called when production constructcall is exited.
func (s *BaseangelscriptListener) ExitConstructcall(ctx *ConstructcallContext) {}

// EnterExprpostop is called when production exprpostop is entered.
func (s *BaseangelscriptListener) EnterExprpostop(ctx *ExprpostopContext) {}

// ExitExprpostop is called when production exprpostop is exited.
func (s *BaseangelscriptListener) ExitExprpostop(ctx *ExprpostopContext) {}

// EnterCast is called when production cast is entered.
func (s *BaseangelscriptListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaseangelscriptListener) ExitCast(ctx *CastContext) {}

// EnterLambda_ is called when production lambda_ is entered.
func (s *BaseangelscriptListener) EnterLambda_(ctx *Lambda_Context) {}

// ExitLambda_ is called when production lambda_ is exited.
func (s *BaseangelscriptListener) ExitLambda_(ctx *Lambda_Context) {}

// EnterFunccall is called when production funccall is entered.
func (s *BaseangelscriptListener) EnterFunccall(ctx *FunccallContext) {}

// ExitFunccall is called when production funccall is exited.
func (s *BaseangelscriptListener) ExitFunccall(ctx *FunccallContext) {}

// EnterVaraccess is called when production varaccess is entered.
func (s *BaseangelscriptListener) EnterVaraccess(ctx *VaraccessContext) {}

// ExitVaraccess is called when production varaccess is exited.
func (s *BaseangelscriptListener) ExitVaraccess(ctx *VaraccessContext) {}

// EnterArglist is called when production arglist is entered.
func (s *BaseangelscriptListener) EnterArglist(ctx *ArglistContext) {}

// ExitArglist is called when production arglist is exited.
func (s *BaseangelscriptListener) ExitArglist(ctx *ArglistContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseangelscriptListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseangelscriptListener) ExitAssign(ctx *AssignContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseangelscriptListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseangelscriptListener) ExitCondition(ctx *ConditionContext) {}

// EnterExprop is called when production exprop is entered.
func (s *BaseangelscriptListener) EnterExprop(ctx *ExpropContext) {}

// ExitExprop is called when production exprop is exited.
func (s *BaseangelscriptListener) ExitExprop(ctx *ExpropContext) {}
