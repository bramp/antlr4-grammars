// Generated from lambda.g4 by ANTLR 4.7.

package lambda // lambda
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaselambdaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaselambdaVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselambdaVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselambdaVisitor) VisitApplication(ctx *ApplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaselambdaVisitor) VisitScope(ctx *ScopeContext) interface{} {
	return v.VisitChildren(ctx)
}
