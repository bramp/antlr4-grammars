// Generated from telephone.g4 by ANTLR 4.7.

package telephone // telephone
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasetelephoneVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetelephoneVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetelephoneVisitor) VisitVariation(ctx *VariationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetelephoneVisitor) VisitNanp(ctx *NanpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetelephoneVisitor) VisitAreacode(ctx *AreacodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetelephoneVisitor) VisitExchange(ctx *ExchangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetelephoneVisitor) VisitSubscriber(ctx *SubscriberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetelephoneVisitor) VisitJapan(ctx *JapanContext) interface{} {
	return v.VisitChildren(ctx)
}
