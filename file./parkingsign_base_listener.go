// Code generated from parkingsign.g4 by ANTLR 4.9.3. DO NOT EDIT.

package file. // parkingsign
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseparkingsignListener is a complete listener for a parse tree produced by parkingsignParser.
type BaseparkingsignListener struct{}

var _ parkingsignListener = &BaseparkingsignListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseparkingsignListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseparkingsignListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseparkingsignListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseparkingsignListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParkingSigns is called when production parkingSigns is entered.
func (s *BaseparkingsignListener) EnterParkingSigns(ctx *ParkingSignsContext) {}

// ExitParkingSigns is called when production parkingSigns is exited.
func (s *BaseparkingsignListener) ExitParkingSigns(ctx *ParkingSignsContext) {}

// EnterParkingSign is called when production parkingSign is entered.
func (s *BaseparkingsignListener) EnterParkingSign(ctx *ParkingSignContext) {}

// ExitParkingSign is called when production parkingSign is exited.
func (s *BaseparkingsignListener) ExitParkingSign(ctx *ParkingSignContext) {}

// EnterStreetSweepingSign is called when production streetSweepingSign is entered.
func (s *BaseparkingsignListener) EnterStreetSweepingSign(ctx *StreetSweepingSignContext) {}

// ExitStreetSweepingSign is called when production streetSweepingSign is exited.
func (s *BaseparkingsignListener) ExitStreetSweepingSign(ctx *StreetSweepingSignContext) {}

// EnterNoParkingSign is called when production noParkingSign is entered.
func (s *BaseparkingsignListener) EnterNoParkingSign(ctx *NoParkingSignContext) {}

// ExitNoParkingSign is called when production noParkingSign is exited.
func (s *BaseparkingsignListener) ExitNoParkingSign(ctx *NoParkingSignContext) {}

// EnterNoStoppingSign is called when production noStoppingSign is entered.
func (s *BaseparkingsignListener) EnterNoStoppingSign(ctx *NoStoppingSignContext) {}

// ExitNoStoppingSign is called when production noStoppingSign is exited.
func (s *BaseparkingsignListener) ExitNoStoppingSign(ctx *NoStoppingSignContext) {}

// EnterPassengerLoadingSign is called when production passengerLoadingSign is entered.
func (s *BaseparkingsignListener) EnterPassengerLoadingSign(ctx *PassengerLoadingSignContext) {}

// ExitPassengerLoadingSign is called when production passengerLoadingSign is exited.
func (s *BaseparkingsignListener) ExitPassengerLoadingSign(ctx *PassengerLoadingSignContext) {}

// EnterTemporaryNoParkingSign is called when production temporaryNoParkingSign is entered.
func (s *BaseparkingsignListener) EnterTemporaryNoParkingSign(ctx *TemporaryNoParkingSignContext) {}

// ExitTemporaryNoParkingSign is called when production temporaryNoParkingSign is exited.
func (s *BaseparkingsignListener) ExitTemporaryNoParkingSign(ctx *TemporaryNoParkingSignContext) {}

// EnterSingleTimeLimitSign is called when production singleTimeLimitSign is entered.
func (s *BaseparkingsignListener) EnterSingleTimeLimitSign(ctx *SingleTimeLimitSignContext) {}

// ExitSingleTimeLimitSign is called when production singleTimeLimitSign is exited.
func (s *BaseparkingsignListener) ExitSingleTimeLimitSign(ctx *SingleTimeLimitSignContext) {}

// EnterDoubleTimeLimitSign is called when production doubleTimeLimitSign is entered.
func (s *BaseparkingsignListener) EnterDoubleTimeLimitSign(ctx *DoubleTimeLimitSignContext) {}

// ExitDoubleTimeLimitSign is called when production doubleTimeLimitSign is exited.
func (s *BaseparkingsignListener) ExitDoubleTimeLimitSign(ctx *DoubleTimeLimitSignContext) {}

// EnterPermitSign is called when production permitSign is entered.
func (s *BaseparkingsignListener) EnterPermitSign(ctx *PermitSignContext) {}

// ExitPermitSign is called when production permitSign is exited.
func (s *BaseparkingsignListener) ExitPermitSign(ctx *PermitSignContext) {}

// EnterStreetSweeping is called when production streetSweeping is entered.
func (s *BaseparkingsignListener) EnterStreetSweeping(ctx *StreetSweepingContext) {}

// ExitStreetSweeping is called when production streetSweeping is exited.
func (s *BaseparkingsignListener) ExitStreetSweeping(ctx *StreetSweepingContext) {}

// EnterNoParking is called when production noParking is entered.
func (s *BaseparkingsignListener) EnterNoParking(ctx *NoParkingContext) {}

// ExitNoParking is called when production noParking is exited.
func (s *BaseparkingsignListener) ExitNoParking(ctx *NoParkingContext) {}

// EnterNoStopping is called when production noStopping is entered.
func (s *BaseparkingsignListener) EnterNoStopping(ctx *NoStoppingContext) {}

// ExitNoStopping is called when production noStopping is exited.
func (s *BaseparkingsignListener) ExitNoStopping(ctx *NoStoppingContext) {}

// EnterValetOnly is called when production valetOnly is entered.
func (s *BaseparkingsignListener) EnterValetOnly(ctx *ValetOnlyContext) {}

// ExitValetOnly is called when production valetOnly is exited.
func (s *BaseparkingsignListener) ExitValetOnly(ctx *ValetOnlyContext) {}

// EnterLoadingOnly is called when production loadingOnly is entered.
func (s *BaseparkingsignListener) EnterLoadingOnly(ctx *LoadingOnlyContext) {}

// ExitLoadingOnly is called when production loadingOnly is exited.
func (s *BaseparkingsignListener) ExitLoadingOnly(ctx *LoadingOnlyContext) {}

// EnterSchoolDays is called when production schoolDays is entered.
func (s *BaseparkingsignListener) EnterSchoolDays(ctx *SchoolDaysContext) {}

// ExitSchoolDays is called when production schoolDays is exited.
func (s *BaseparkingsignListener) ExitSchoolDays(ctx *SchoolDaysContext) {}

// EnterTimeRange is called when production timeRange is entered.
func (s *BaseparkingsignListener) EnterTimeRange(ctx *TimeRangeContext) {}

// ExitTimeRange is called when production timeRange is exited.
func (s *BaseparkingsignListener) ExitTimeRange(ctx *TimeRangeContext) {}

// EnterEveryDay is called when production everyDay is entered.
func (s *BaseparkingsignListener) EnterEveryDay(ctx *EveryDayContext) {}

// ExitEveryDay is called when production everyDay is exited.
func (s *BaseparkingsignListener) ExitEveryDay(ctx *EveryDayContext) {}

// EnterDayToDay is called when production dayToDay is entered.
func (s *BaseparkingsignListener) EnterDayToDay(ctx *DayToDayContext) {}

// ExitDayToDay is called when production dayToDay is exited.
func (s *BaseparkingsignListener) ExitDayToDay(ctx *DayToDayContext) {}

// EnterDayAndDay is called when production dayAndDay is entered.
func (s *BaseparkingsignListener) EnterDayAndDay(ctx *DayAndDayContext) {}

// ExitDayAndDay is called when production dayAndDay is exited.
func (s *BaseparkingsignListener) ExitDayAndDay(ctx *DayAndDayContext) {}

// EnterDayRange is called when production dayRange is entered.
func (s *BaseparkingsignListener) EnterDayRange(ctx *DayRangeContext) {}

// ExitDayRange is called when production dayRange is exited.
func (s *BaseparkingsignListener) ExitDayRange(ctx *DayRangeContext) {}

// EnterDayRangePlus is called when production dayRangePlus is entered.
func (s *BaseparkingsignListener) EnterDayRangePlus(ctx *DayRangePlusContext) {}

// ExitDayRangePlus is called when production dayRangePlus is exited.
func (s *BaseparkingsignListener) ExitDayRangePlus(ctx *DayRangePlusContext) {}

// EnterTo is called when production to is entered.
func (s *BaseparkingsignListener) EnterTo(ctx *ToContext) {}

// ExitTo is called when production to is exited.
func (s *BaseparkingsignListener) ExitTo(ctx *ToContext) {}

// EnterAnd_ is called when production and_ is entered.
func (s *BaseparkingsignListener) EnterAnd_(ctx *And_Context) {}

// ExitAnd_ is called when production and_ is exited.
func (s *BaseparkingsignListener) ExitAnd_(ctx *And_Context) {}

// EnterTowAway is called when production towAway is entered.
func (s *BaseparkingsignListener) EnterTowAway(ctx *TowAwayContext) {}

// ExitTowAway is called when production towAway is exited.
func (s *BaseparkingsignListener) ExitTowAway(ctx *TowAwayContext) {}

// EnterMinute is called when production minute is entered.
func (s *BaseparkingsignListener) EnterMinute(ctx *MinuteContext) {}

// ExitMinute is called when production minute is exited.
func (s *BaseparkingsignListener) ExitMinute(ctx *MinuteContext) {}

// EnterExempt is called when production exempt is entered.
func (s *BaseparkingsignListener) EnterExempt(ctx *ExemptContext) {}

// ExitExempt is called when production exempt is exited.
func (s *BaseparkingsignListener) ExitExempt(ctx *ExemptContext) {}

// EnterAnyTime is called when production anyTime is entered.
func (s *BaseparkingsignListener) EnterAnyTime(ctx *AnyTimeContext) {}

// ExitAnyTime is called when production anyTime is exited.
func (s *BaseparkingsignListener) ExitAnyTime(ctx *AnyTimeContext) {}

// EnterDay is called when production day is entered.
func (s *BaseparkingsignListener) EnterDay(ctx *DayContext) {}

// ExitDay is called when production day is exited.
func (s *BaseparkingsignListener) ExitDay(ctx *DayContext) {}

// EnterTime is called when production time is entered.
func (s *BaseparkingsignListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BaseparkingsignListener) ExitTime(ctx *TimeContext) {}

// EnterTwelveNoon is called when production twelveNoon is entered.
func (s *BaseparkingsignListener) EnterTwelveNoon(ctx *TwelveNoonContext) {}

// ExitTwelveNoon is called when production twelveNoon is exited.
func (s *BaseparkingsignListener) ExitTwelveNoon(ctx *TwelveNoonContext) {}

// EnterTwelveMidnight is called when production twelveMidnight is entered.
func (s *BaseparkingsignListener) EnterTwelveMidnight(ctx *TwelveMidnightContext) {}

// ExitTwelveMidnight is called when production twelveMidnight is exited.
func (s *BaseparkingsignListener) ExitTwelveMidnight(ctx *TwelveMidnightContext) {}

// EnterAm is called when production am is entered.
func (s *BaseparkingsignListener) EnterAm(ctx *AmContext) {}

// ExitAm is called when production am is exited.
func (s *BaseparkingsignListener) ExitAm(ctx *AmContext) {}

// EnterPm is called when production pm is entered.
func (s *BaseparkingsignListener) EnterPm(ctx *PmContext) {}

// ExitPm is called when production pm is exited.
func (s *BaseparkingsignListener) ExitPm(ctx *PmContext) {}
