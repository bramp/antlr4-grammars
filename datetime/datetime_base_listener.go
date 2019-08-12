// Code generated from datetime.g4 by ANTLR 4.7.2. DO NOT EDIT.

package datetime // datetime
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasedatetimeListener is a complete listener for a parse tree produced by datetimeParser.
type BasedatetimeListener struct{}

var _ datetimeListener = &BasedatetimeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasedatetimeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasedatetimeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasedatetimeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasedatetimeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDate_time is called when production date_time is entered.
func (s *BasedatetimeListener) EnterDate_time(ctx *Date_timeContext) {}

// ExitDate_time is called when production date_time is exited.
func (s *BasedatetimeListener) ExitDate_time(ctx *Date_timeContext) {}

// EnterDay is called when production day is entered.
func (s *BasedatetimeListener) EnterDay(ctx *DayContext) {}

// ExitDay is called when production day is exited.
func (s *BasedatetimeListener) ExitDay(ctx *DayContext) {}

// EnterDate is called when production date is entered.
func (s *BasedatetimeListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BasedatetimeListener) ExitDate(ctx *DateContext) {}

// EnterMonth is called when production month is entered.
func (s *BasedatetimeListener) EnterMonth(ctx *MonthContext) {}

// ExitMonth is called when production month is exited.
func (s *BasedatetimeListener) ExitMonth(ctx *MonthContext) {}

// EnterTime is called when production time is entered.
func (s *BasedatetimeListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BasedatetimeListener) ExitTime(ctx *TimeContext) {}

// EnterHour is called when production hour is entered.
func (s *BasedatetimeListener) EnterHour(ctx *HourContext) {}

// ExitHour is called when production hour is exited.
func (s *BasedatetimeListener) ExitHour(ctx *HourContext) {}

// EnterZone is called when production zone is entered.
func (s *BasedatetimeListener) EnterZone(ctx *ZoneContext) {}

// ExitZone is called when production zone is exited.
func (s *BasedatetimeListener) ExitZone(ctx *ZoneContext) {}

// EnterTwo_digit is called when production two_digit is entered.
func (s *BasedatetimeListener) EnterTwo_digit(ctx *Two_digitContext) {}

// ExitTwo_digit is called when production two_digit is exited.
func (s *BasedatetimeListener) ExitTwo_digit(ctx *Two_digitContext) {}

// EnterFour_digit is called when production four_digit is entered.
func (s *BasedatetimeListener) EnterFour_digit(ctx *Four_digitContext) {}

// ExitFour_digit is called when production four_digit is exited.
func (s *BasedatetimeListener) ExitFour_digit(ctx *Four_digitContext) {}

// EnterAlphanumeric is called when production alphanumeric is entered.
func (s *BasedatetimeListener) EnterAlphanumeric(ctx *AlphanumericContext) {}

// ExitAlphanumeric is called when production alphanumeric is exited.
func (s *BasedatetimeListener) ExitAlphanumeric(ctx *AlphanumericContext) {}
