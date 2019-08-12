// Code generated from telephone.g4 by ANTLR 4.7.2. DO NOT EDIT.

package telephone // telephone
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetelephoneListener is a complete listener for a parse tree produced by telephoneParser.
type BasetelephoneListener struct{}

var _ telephoneListener = &BasetelephoneListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetelephoneListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetelephoneListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetelephoneListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetelephoneListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNumber is called when production number is entered.
func (s *BasetelephoneListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasetelephoneListener) ExitNumber(ctx *NumberContext) {}

// EnterVariation is called when production variation is entered.
func (s *BasetelephoneListener) EnterVariation(ctx *VariationContext) {}

// ExitVariation is called when production variation is exited.
func (s *BasetelephoneListener) ExitVariation(ctx *VariationContext) {}

// EnterNanp is called when production nanp is entered.
func (s *BasetelephoneListener) EnterNanp(ctx *NanpContext) {}

// ExitNanp is called when production nanp is exited.
func (s *BasetelephoneListener) ExitNanp(ctx *NanpContext) {}

// EnterAreacode is called when production areacode is entered.
func (s *BasetelephoneListener) EnterAreacode(ctx *AreacodeContext) {}

// ExitAreacode is called when production areacode is exited.
func (s *BasetelephoneListener) ExitAreacode(ctx *AreacodeContext) {}

// EnterExchange is called when production exchange is entered.
func (s *BasetelephoneListener) EnterExchange(ctx *ExchangeContext) {}

// ExitExchange is called when production exchange is exited.
func (s *BasetelephoneListener) ExitExchange(ctx *ExchangeContext) {}

// EnterSubscriber is called when production subscriber is entered.
func (s *BasetelephoneListener) EnterSubscriber(ctx *SubscriberContext) {}

// ExitSubscriber is called when production subscriber is exited.
func (s *BasetelephoneListener) ExitSubscriber(ctx *SubscriberContext) {}

// EnterJapan is called when production japan is entered.
func (s *BasetelephoneListener) EnterJapan(ctx *JapanContext) {}

// ExitJapan is called when production japan is exited.
func (s *BasetelephoneListener) ExitJapan(ctx *JapanContext) {}
