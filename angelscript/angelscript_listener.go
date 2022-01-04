// Code generated from angelscript.g4 by ANTLR 4.9.3. DO NOT EDIT.

package angelscript // angelscript
import "github.com/antlr/antlr4/runtime/Go/antlr"

// angelscriptListener is a complete listener for a parse tree produced by angelscriptParser.
type angelscriptListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterClass_ is called when entering the class_ production.
	EnterClass_(c *Class_Context)

	// EnterTypdef is called when entering the typdef production.
	EnterTypdef(c *TypdefContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterFunc_ is called when entering the func_ production.
	EnterFunc_(c *Func_Context)

	// EnterInterface_ is called when entering the interface_ production.
	EnterInterface_(c *Interface_Context)

	// EnterVar_ is called when entering the var_ production.
	EnterVar_(c *Var_Context)

	// EnterImport_ is called when entering the import_ production.
	EnterImport_(c *Import_Context)

	// EnterEnum_ is called when entering the enum_ production.
	EnterEnum_(c *Enum_Context)

	// EnterFuncdef is called when entering the funcdef production.
	EnterFuncdef(c *FuncdefContext)

	// EnterVirtprop is called when entering the virtprop production.
	EnterVirtprop(c *VirtpropContext)

	// EnterMixin_ is called when entering the mixin_ production.
	EnterMixin_(c *Mixin_Context)

	// EnterIntfmthd is called when entering the intfmthd production.
	EnterIntfmthd(c *IntfmthdContext)

	// EnterStatblock is called when entering the statblock production.
	EnterStatblock(c *StatblockContext)

	// EnterParamlist is called when entering the paramlist production.
	EnterParamlist(c *ParamlistContext)

	// EnterTypemod is called when entering the typemod production.
	EnterTypemod(c *TypemodContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterInitlist is called when entering the initlist production.
	EnterInitlist(c *InitlistContext)

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

	// EnterDatatype is called when entering the datatype production.
	EnterDatatype(c *DatatypeContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSwitch_ is called when entering the switch_ production.
	EnterSwitch_(c *Switch_Context)

	// EnterBreak_ is called when entering the break_ production.
	EnterBreak_(c *Break_Context)

	// EnterFor_ is called when entering the for_ production.
	EnterFor_(c *For_Context)

	// EnterWhile_ is called when entering the while_ production.
	EnterWhile_(c *While_Context)

	// EnterDowhile is called when entering the dowhile production.
	EnterDowhile(c *DowhileContext)

	// EnterIf_ is called when entering the if_ production.
	EnterIf_(c *If_Context)

	// EnterContinue_ is called when entering the continue_ production.
	EnterContinue_(c *Continue_Context)

	// EnterExprstat is called when entering the exprstat production.
	EnterExprstat(c *ExprstatContext)

	// EnterTry_ is called when entering the try_ production.
	EnterTry_(c *Try_Context)

	// EnterReturn_ is called when entering the return_ production.
	EnterReturn_(c *Return_Context)

	// EnterCase_ is called when entering the case_ production.
	EnterCase_(c *Case_Context)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExprterm is called when entering the exprterm production.
	EnterExprterm(c *ExprtermContext)

	// EnterExprvalue is called when entering the exprvalue production.
	EnterExprvalue(c *ExprvalueContext)

	// EnterConstructcall is called when entering the constructcall production.
	EnterConstructcall(c *ConstructcallContext)

	// EnterExprpostop is called when entering the exprpostop production.
	EnterExprpostop(c *ExprpostopContext)

	// EnterCast is called when entering the cast production.
	EnterCast(c *CastContext)

	// EnterLambda_ is called when entering the lambda_ production.
	EnterLambda_(c *Lambda_Context)

	// EnterFunccall is called when entering the funccall production.
	EnterFunccall(c *FunccallContext)

	// EnterVaraccess is called when entering the varaccess production.
	EnterVaraccess(c *VaraccessContext)

	// EnterArglist is called when entering the arglist production.
	EnterArglist(c *ArglistContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterExprop is called when entering the exprop production.
	EnterExprop(c *ExpropContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitClass_ is called when exiting the class_ production.
	ExitClass_(c *Class_Context)

	// ExitTypdef is called when exiting the typdef production.
	ExitTypdef(c *TypdefContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitFunc_ is called when exiting the func_ production.
	ExitFunc_(c *Func_Context)

	// ExitInterface_ is called when exiting the interface_ production.
	ExitInterface_(c *Interface_Context)

	// ExitVar_ is called when exiting the var_ production.
	ExitVar_(c *Var_Context)

	// ExitImport_ is called when exiting the import_ production.
	ExitImport_(c *Import_Context)

	// ExitEnum_ is called when exiting the enum_ production.
	ExitEnum_(c *Enum_Context)

	// ExitFuncdef is called when exiting the funcdef production.
	ExitFuncdef(c *FuncdefContext)

	// ExitVirtprop is called when exiting the virtprop production.
	ExitVirtprop(c *VirtpropContext)

	// ExitMixin_ is called when exiting the mixin_ production.
	ExitMixin_(c *Mixin_Context)

	// ExitIntfmthd is called when exiting the intfmthd production.
	ExitIntfmthd(c *IntfmthdContext)

	// ExitStatblock is called when exiting the statblock production.
	ExitStatblock(c *StatblockContext)

	// ExitParamlist is called when exiting the paramlist production.
	ExitParamlist(c *ParamlistContext)

	// ExitTypemod is called when exiting the typemod production.
	ExitTypemod(c *TypemodContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitInitlist is called when exiting the initlist production.
	ExitInitlist(c *InitlistContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)

	// ExitDatatype is called when exiting the datatype production.
	ExitDatatype(c *DatatypeContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSwitch_ is called when exiting the switch_ production.
	ExitSwitch_(c *Switch_Context)

	// ExitBreak_ is called when exiting the break_ production.
	ExitBreak_(c *Break_Context)

	// ExitFor_ is called when exiting the for_ production.
	ExitFor_(c *For_Context)

	// ExitWhile_ is called when exiting the while_ production.
	ExitWhile_(c *While_Context)

	// ExitDowhile is called when exiting the dowhile production.
	ExitDowhile(c *DowhileContext)

	// ExitIf_ is called when exiting the if_ production.
	ExitIf_(c *If_Context)

	// ExitContinue_ is called when exiting the continue_ production.
	ExitContinue_(c *Continue_Context)

	// ExitExprstat is called when exiting the exprstat production.
	ExitExprstat(c *ExprstatContext)

	// ExitTry_ is called when exiting the try_ production.
	ExitTry_(c *Try_Context)

	// ExitReturn_ is called when exiting the return_ production.
	ExitReturn_(c *Return_Context)

	// ExitCase_ is called when exiting the case_ production.
	ExitCase_(c *Case_Context)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExprterm is called when exiting the exprterm production.
	ExitExprterm(c *ExprtermContext)

	// ExitExprvalue is called when exiting the exprvalue production.
	ExitExprvalue(c *ExprvalueContext)

	// ExitConstructcall is called when exiting the constructcall production.
	ExitConstructcall(c *ConstructcallContext)

	// ExitExprpostop is called when exiting the exprpostop production.
	ExitExprpostop(c *ExprpostopContext)

	// ExitCast is called when exiting the cast production.
	ExitCast(c *CastContext)

	// ExitLambda_ is called when exiting the lambda_ production.
	ExitLambda_(c *Lambda_Context)

	// ExitFunccall is called when exiting the funccall production.
	ExitFunccall(c *FunccallContext)

	// ExitVaraccess is called when exiting the varaccess production.
	ExitVaraccess(c *VaraccessContext)

	// ExitArglist is called when exiting the arglist production.
	ExitArglist(c *ArglistContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitExprop is called when exiting the exprop production.
	ExitExprop(c *ExpropContext)
}
