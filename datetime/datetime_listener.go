// Code generated from datetime.g4 by ANTLR 4.7.2. DO NOT EDIT.

package datetime // datetime
import "github.com/antlr/antlr4/runtime/Go/antlr"

// datetimeListener is a complete listener for a parse tree produced by datetimeParser.
type datetimeListener interface {
	antlr.ParseTreeListener

	// EnterDate_time is called when entering the date_time production.
	EnterDate_time(c *Date_timeContext)

	// EnterDay is called when entering the day production.
	EnterDay(c *DayContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterMonth is called when entering the month production.
	EnterMonth(c *MonthContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// EnterHour is called when entering the hour production.
	EnterHour(c *HourContext)

	// EnterZone is called when entering the zone production.
	EnterZone(c *ZoneContext)

	// EnterTwo_digit is called when entering the two_digit production.
	EnterTwo_digit(c *Two_digitContext)

	// EnterFour_digit is called when entering the four_digit production.
	EnterFour_digit(c *Four_digitContext)

	// EnterAlphanumeric is called when entering the alphanumeric production.
	EnterAlphanumeric(c *AlphanumericContext)

	// ExitDate_time is called when exiting the date_time production.
	ExitDate_time(c *Date_timeContext)

	// ExitDay is called when exiting the day production.
	ExitDay(c *DayContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitMonth is called when exiting the month production.
	ExitMonth(c *MonthContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)

	// ExitHour is called when exiting the hour production.
	ExitHour(c *HourContext)

	// ExitZone is called when exiting the zone production.
	ExitZone(c *ZoneContext)

	// ExitTwo_digit is called when exiting the two_digit production.
	ExitTwo_digit(c *Two_digitContext)

	// ExitFour_digit is called when exiting the four_digit production.
	ExitFour_digit(c *Four_digitContext)

	// ExitAlphanumeric is called when exiting the alphanumeric production.
	ExitAlphanumeric(c *AlphanumericContext)
}
