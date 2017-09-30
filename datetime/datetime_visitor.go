// Generated from datetime.g4 by ANTLR 4.7.

package datetime // datetime
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by datetimeParser.
type datetimeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by datetimeParser#date_time.
	VisitDate_time(ctx *Date_timeContext) interface{}

	// Visit a parse tree produced by datetimeParser#day.
	VisitDay(ctx *DayContext) interface{}

	// Visit a parse tree produced by datetimeParser#date.
	VisitDate(ctx *DateContext) interface{}

	// Visit a parse tree produced by datetimeParser#month.
	VisitMonth(ctx *MonthContext) interface{}

	// Visit a parse tree produced by datetimeParser#time.
	VisitTime(ctx *TimeContext) interface{}

	// Visit a parse tree produced by datetimeParser#hour.
	VisitHour(ctx *HourContext) interface{}

	// Visit a parse tree produced by datetimeParser#zone.
	VisitZone(ctx *ZoneContext) interface{}

	// Visit a parse tree produced by datetimeParser#two_digit.
	VisitTwo_digit(ctx *Two_digitContext) interface{}

	// Visit a parse tree produced by datetimeParser#four_digit.
	VisitFour_digit(ctx *Four_digitContext) interface{}

	// Visit a parse tree produced by datetimeParser#alphanumeric.
	VisitAlphanumeric(ctx *AlphanumericContext) interface{}
}
