// Code generated from parkingsign.g4 by ANTLR 4.9.3. DO NOT EDIT.

package file. // parkingsign
import "github.com/antlr/antlr4/runtime/Go/antlr"

// parkingsignListener is a complete listener for a parse tree produced by parkingsignParser.
type parkingsignListener interface {
	antlr.ParseTreeListener

	// EnterParkingSigns is called when entering the parkingSigns production.
	EnterParkingSigns(c *ParkingSignsContext)

	// EnterParkingSign is called when entering the parkingSign production.
	EnterParkingSign(c *ParkingSignContext)

	// EnterStreetSweepingSign is called when entering the streetSweepingSign production.
	EnterStreetSweepingSign(c *StreetSweepingSignContext)

	// EnterNoParkingSign is called when entering the noParkingSign production.
	EnterNoParkingSign(c *NoParkingSignContext)

	// EnterNoStoppingSign is called when entering the noStoppingSign production.
	EnterNoStoppingSign(c *NoStoppingSignContext)

	// EnterPassengerLoadingSign is called when entering the passengerLoadingSign production.
	EnterPassengerLoadingSign(c *PassengerLoadingSignContext)

	// EnterTemporaryNoParkingSign is called when entering the temporaryNoParkingSign production.
	EnterTemporaryNoParkingSign(c *TemporaryNoParkingSignContext)

	// EnterSingleTimeLimitSign is called when entering the singleTimeLimitSign production.
	EnterSingleTimeLimitSign(c *SingleTimeLimitSignContext)

	// EnterDoubleTimeLimitSign is called when entering the doubleTimeLimitSign production.
	EnterDoubleTimeLimitSign(c *DoubleTimeLimitSignContext)

	// EnterPermitSign is called when entering the permitSign production.
	EnterPermitSign(c *PermitSignContext)

	// EnterStreetSweeping is called when entering the streetSweeping production.
	EnterStreetSweeping(c *StreetSweepingContext)

	// EnterNoParking is called when entering the noParking production.
	EnterNoParking(c *NoParkingContext)

	// EnterNoStopping is called when entering the noStopping production.
	EnterNoStopping(c *NoStoppingContext)

	// EnterValetOnly is called when entering the valetOnly production.
	EnterValetOnly(c *ValetOnlyContext)

	// EnterLoadingOnly is called when entering the loadingOnly production.
	EnterLoadingOnly(c *LoadingOnlyContext)

	// EnterSchoolDays is called when entering the schoolDays production.
	EnterSchoolDays(c *SchoolDaysContext)

	// EnterTimeRange is called when entering the timeRange production.
	EnterTimeRange(c *TimeRangeContext)

	// EnterEveryDay is called when entering the everyDay production.
	EnterEveryDay(c *EveryDayContext)

	// EnterDayToDay is called when entering the dayToDay production.
	EnterDayToDay(c *DayToDayContext)

	// EnterDayAndDay is called when entering the dayAndDay production.
	EnterDayAndDay(c *DayAndDayContext)

	// EnterDayRange is called when entering the dayRange production.
	EnterDayRange(c *DayRangeContext)

	// EnterDayRangePlus is called when entering the dayRangePlus production.
	EnterDayRangePlus(c *DayRangePlusContext)

	// EnterTo is called when entering the to production.
	EnterTo(c *ToContext)

	// EnterAnd_ is called when entering the and_ production.
	EnterAnd_(c *And_Context)

	// EnterTowAway is called when entering the towAway production.
	EnterTowAway(c *TowAwayContext)

	// EnterMinute is called when entering the minute production.
	EnterMinute(c *MinuteContext)

	// EnterExempt is called when entering the exempt production.
	EnterExempt(c *ExemptContext)

	// EnterAnyTime is called when entering the anyTime production.
	EnterAnyTime(c *AnyTimeContext)

	// EnterDay is called when entering the day production.
	EnterDay(c *DayContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// EnterTwelveNoon is called when entering the twelveNoon production.
	EnterTwelveNoon(c *TwelveNoonContext)

	// EnterTwelveMidnight is called when entering the twelveMidnight production.
	EnterTwelveMidnight(c *TwelveMidnightContext)

	// EnterAm is called when entering the am production.
	EnterAm(c *AmContext)

	// EnterPm is called when entering the pm production.
	EnterPm(c *PmContext)

	// ExitParkingSigns is called when exiting the parkingSigns production.
	ExitParkingSigns(c *ParkingSignsContext)

	// ExitParkingSign is called when exiting the parkingSign production.
	ExitParkingSign(c *ParkingSignContext)

	// ExitStreetSweepingSign is called when exiting the streetSweepingSign production.
	ExitStreetSweepingSign(c *StreetSweepingSignContext)

	// ExitNoParkingSign is called when exiting the noParkingSign production.
	ExitNoParkingSign(c *NoParkingSignContext)

	// ExitNoStoppingSign is called when exiting the noStoppingSign production.
	ExitNoStoppingSign(c *NoStoppingSignContext)

	// ExitPassengerLoadingSign is called when exiting the passengerLoadingSign production.
	ExitPassengerLoadingSign(c *PassengerLoadingSignContext)

	// ExitTemporaryNoParkingSign is called when exiting the temporaryNoParkingSign production.
	ExitTemporaryNoParkingSign(c *TemporaryNoParkingSignContext)

	// ExitSingleTimeLimitSign is called when exiting the singleTimeLimitSign production.
	ExitSingleTimeLimitSign(c *SingleTimeLimitSignContext)

	// ExitDoubleTimeLimitSign is called when exiting the doubleTimeLimitSign production.
	ExitDoubleTimeLimitSign(c *DoubleTimeLimitSignContext)

	// ExitPermitSign is called when exiting the permitSign production.
	ExitPermitSign(c *PermitSignContext)

	// ExitStreetSweeping is called when exiting the streetSweeping production.
	ExitStreetSweeping(c *StreetSweepingContext)

	// ExitNoParking is called when exiting the noParking production.
	ExitNoParking(c *NoParkingContext)

	// ExitNoStopping is called when exiting the noStopping production.
	ExitNoStopping(c *NoStoppingContext)

	// ExitValetOnly is called when exiting the valetOnly production.
	ExitValetOnly(c *ValetOnlyContext)

	// ExitLoadingOnly is called when exiting the loadingOnly production.
	ExitLoadingOnly(c *LoadingOnlyContext)

	// ExitSchoolDays is called when exiting the schoolDays production.
	ExitSchoolDays(c *SchoolDaysContext)

	// ExitTimeRange is called when exiting the timeRange production.
	ExitTimeRange(c *TimeRangeContext)

	// ExitEveryDay is called when exiting the everyDay production.
	ExitEveryDay(c *EveryDayContext)

	// ExitDayToDay is called when exiting the dayToDay production.
	ExitDayToDay(c *DayToDayContext)

	// ExitDayAndDay is called when exiting the dayAndDay production.
	ExitDayAndDay(c *DayAndDayContext)

	// ExitDayRange is called when exiting the dayRange production.
	ExitDayRange(c *DayRangeContext)

	// ExitDayRangePlus is called when exiting the dayRangePlus production.
	ExitDayRangePlus(c *DayRangePlusContext)

	// ExitTo is called when exiting the to production.
	ExitTo(c *ToContext)

	// ExitAnd_ is called when exiting the and_ production.
	ExitAnd_(c *And_Context)

	// ExitTowAway is called when exiting the towAway production.
	ExitTowAway(c *TowAwayContext)

	// ExitMinute is called when exiting the minute production.
	ExitMinute(c *MinuteContext)

	// ExitExempt is called when exiting the exempt production.
	ExitExempt(c *ExemptContext)

	// ExitAnyTime is called when exiting the anyTime production.
	ExitAnyTime(c *AnyTimeContext)

	// ExitDay is called when exiting the day production.
	ExitDay(c *DayContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)

	// ExitTwelveNoon is called when exiting the twelveNoon production.
	ExitTwelveNoon(c *TwelveNoonContext)

	// ExitTwelveMidnight is called when exiting the twelveMidnight production.
	ExitTwelveMidnight(c *TwelveMidnightContext)

	// ExitAm is called when exiting the am production.
	ExitAm(c *AmContext)

	// ExitPm is called when exiting the pm production.
	ExitPm(c *PmContext)
}
