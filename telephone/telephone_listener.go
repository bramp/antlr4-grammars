// Code generated from telephone.g4 by ANTLR 4.9.3. DO NOT EDIT.

package telephone // telephone
import "github.com/antlr/antlr4/runtime/Go/antlr"

// telephoneListener is a complete listener for a parse tree produced by telephoneParser.
type telephoneListener interface {
	antlr.ParseTreeListener

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterVariation is called when entering the variation production.
	EnterVariation(c *VariationContext)

	// EnterNanp is called when entering the nanp production.
	EnterNanp(c *NanpContext)

	// EnterAreacode is called when entering the areacode production.
	EnterAreacode(c *AreacodeContext)

	// EnterExchange is called when entering the exchange production.
	EnterExchange(c *ExchangeContext)

	// EnterSubscriber is called when entering the subscriber production.
	EnterSubscriber(c *SubscriberContext)

	// EnterJapan is called when entering the japan production.
	EnterJapan(c *JapanContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitVariation is called when exiting the variation production.
	ExitVariation(c *VariationContext)

	// ExitNanp is called when exiting the nanp production.
	ExitNanp(c *NanpContext)

	// ExitAreacode is called when exiting the areacode production.
	ExitAreacode(c *AreacodeContext)

	// ExitExchange is called when exiting the exchange production.
	ExitExchange(c *ExchangeContext)

	// ExitSubscriber is called when exiting the subscriber production.
	ExitSubscriber(c *SubscriberContext)

	// ExitJapan is called when exiting the japan production.
	ExitJapan(c *JapanContext)
}
