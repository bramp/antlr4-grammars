// Generated from istc.g4 by ANTLR 4.7.

package istc // istc
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseistcVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseistcVisitor) VisitIstc(ctx *IstcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseistcVisitor) VisitRegistration(ctx *RegistrationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseistcVisitor) VisitYear(ctx *YearContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseistcVisitor) VisitWork(ctx *WorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseistcVisitor) VisitCheck(ctx *CheckContext) interface{} {
	return v.VisitChildren(ctx)
}
