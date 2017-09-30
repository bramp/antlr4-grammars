// Generated from telephone.g4 by ANTLR 4.7.

package telephone // telephone
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by telephoneParser.
type telephoneVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by telephoneParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by telephoneParser#variation.
	VisitVariation(ctx *VariationContext) interface{}

	// Visit a parse tree produced by telephoneParser#nanp.
	VisitNanp(ctx *NanpContext) interface{}

	// Visit a parse tree produced by telephoneParser#areacode.
	VisitAreacode(ctx *AreacodeContext) interface{}

	// Visit a parse tree produced by telephoneParser#exchange.
	VisitExchange(ctx *ExchangeContext) interface{}

	// Visit a parse tree produced by telephoneParser#subscriber.
	VisitSubscriber(ctx *SubscriberContext) interface{}

	// Visit a parse tree produced by telephoneParser#japan.
	VisitJapan(ctx *JapanContext) interface{}
}
