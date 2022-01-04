// Code generated from iso8601.g4 by ANTLR 4.9.3. DO NOT EDIT.

package iso8601 // iso8601
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Baseiso8601Listener is a complete listener for a parse tree produced by iso8601Parser.
type Baseiso8601Listener struct{}

var _ iso8601Listener = &Baseiso8601Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baseiso8601Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baseiso8601Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baseiso8601Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baseiso8601Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInt_ is called when production int_ is entered.
func (s *Baseiso8601Listener) EnterInt_(ctx *Int_Context) {}

// ExitInt_ is called when production int_ is exited.
func (s *Baseiso8601Listener) ExitInt_(ctx *Int_Context) {}

// EnterDec is called when production dec is entered.
func (s *Baseiso8601Listener) EnterDec(ctx *DecContext) {}

// ExitDec is called when production dec is exited.
func (s *Baseiso8601Listener) ExitDec(ctx *DecContext) {}

// EnterInt2 is called when production int2 is entered.
func (s *Baseiso8601Listener) EnterInt2(ctx *Int2Context) {}

// ExitInt2 is called when production int2 is exited.
func (s *Baseiso8601Listener) ExitInt2(ctx *Int2Context) {}

// EnterInt3 is called when production int3 is entered.
func (s *Baseiso8601Listener) EnterInt3(ctx *Int3Context) {}

// ExitInt3 is called when production int3 is exited.
func (s *Baseiso8601Listener) ExitInt3(ctx *Int3Context) {}

// EnterInt4 is called when production int4 is entered.
func (s *Baseiso8601Listener) EnterInt4(ctx *Int4Context) {}

// ExitInt4 is called when production int4 is exited.
func (s *Baseiso8601Listener) ExitInt4(ctx *Int4Context) {}

// EnterSint2p is called when production sint2p is entered.
func (s *Baseiso8601Listener) EnterSint2p(ctx *Sint2pContext) {}

// ExitSint2p is called when production sint2p is exited.
func (s *Baseiso8601Listener) ExitSint2p(ctx *Sint2pContext) {}

// EnterSint4p is called when production sint4p is entered.
func (s *Baseiso8601Listener) EnterSint4p(ctx *Sint4pContext) {}

// ExitSint4p is called when production sint4p is exited.
func (s *Baseiso8601Listener) ExitSint4p(ctx *Sint4pContext) {}

// EnterDec2 is called when production dec2 is entered.
func (s *Baseiso8601Listener) EnterDec2(ctx *Dec2Context) {}

// ExitDec2 is called when production dec2 is exited.
func (s *Baseiso8601Listener) ExitDec2(ctx *Dec2Context) {}

// EnterCompleteCentury is called when production CompleteCentury is entered.
func (s *Baseiso8601Listener) EnterCompleteCentury(ctx *CompleteCenturyContext) {}

// ExitCompleteCentury is called when production CompleteCentury is exited.
func (s *Baseiso8601Listener) ExitCompleteCentury(ctx *CompleteCenturyContext) {}

// EnterExpandedCentury is called when production ExpandedCentury is entered.
func (s *Baseiso8601Listener) EnterExpandedCentury(ctx *ExpandedCenturyContext) {}

// ExitExpandedCentury is called when production ExpandedCentury is exited.
func (s *Baseiso8601Listener) ExitExpandedCentury(ctx *ExpandedCenturyContext) {}

// EnterCompleteYear is called when production CompleteYear is entered.
func (s *Baseiso8601Listener) EnterCompleteYear(ctx *CompleteYearContext) {}

// ExitCompleteYear is called when production CompleteYear is exited.
func (s *Baseiso8601Listener) ExitCompleteYear(ctx *CompleteYearContext) {}

// EnterExpandedYear is called when production ExpandedYear is entered.
func (s *Baseiso8601Listener) EnterExpandedYear(ctx *ExpandedYearContext) {}

// ExitExpandedYear is called when production ExpandedYear is exited.
func (s *Baseiso8601Listener) ExitExpandedYear(ctx *ExpandedYearContext) {}

// EnterMonth is called when production month is entered.
func (s *Baseiso8601Listener) EnterMonth(ctx *MonthContext) {}

// ExitMonth is called when production month is exited.
func (s *Baseiso8601Listener) ExitMonth(ctx *MonthContext) {}

// EnterDay is called when production day is entered.
func (s *Baseiso8601Listener) EnterDay(ctx *DayContext) {}

// ExitDay is called when production day is exited.
func (s *Baseiso8601Listener) ExitDay(ctx *DayContext) {}

// EnterOrdinalDay is called when production ordinalDay is entered.
func (s *Baseiso8601Listener) EnterOrdinalDay(ctx *OrdinalDayContext) {}

// ExitOrdinalDay is called when production ordinalDay is exited.
func (s *Baseiso8601Listener) ExitOrdinalDay(ctx *OrdinalDayContext) {}

// EnterWeek is called when production week is entered.
func (s *Baseiso8601Listener) EnterWeek(ctx *WeekContext) {}

// ExitWeek is called when production week is exited.
func (s *Baseiso8601Listener) ExitWeek(ctx *WeekContext) {}

// EnterWeekDay is called when production weekDay is entered.
func (s *Baseiso8601Listener) EnterWeekDay(ctx *WeekDayContext) {}

// ExitWeekDay is called when production weekDay is exited.
func (s *Baseiso8601Listener) ExitWeekDay(ctx *WeekDayContext) {}

// EnterHour is called when production hour is entered.
func (s *Baseiso8601Listener) EnterHour(ctx *HourContext) {}

// ExitHour is called when production hour is exited.
func (s *Baseiso8601Listener) ExitHour(ctx *HourContext) {}

// EnterMinute is called when production minute is entered.
func (s *Baseiso8601Listener) EnterMinute(ctx *MinuteContext) {}

// ExitMinute is called when production minute is exited.
func (s *Baseiso8601Listener) ExitMinute(ctx *MinuteContext) {}

// EnterSecond is called when production second is entered.
func (s *Baseiso8601Listener) EnterSecond(ctx *SecondContext) {}

// ExitSecond is called when production second is exited.
func (s *Baseiso8601Listener) ExitSecond(ctx *SecondContext) {}

// EnterHourFraction is called when production hourFraction is entered.
func (s *Baseiso8601Listener) EnterHourFraction(ctx *HourFractionContext) {}

// ExitHourFraction is called when production hourFraction is exited.
func (s *Baseiso8601Listener) ExitHourFraction(ctx *HourFractionContext) {}

// EnterMinuteFraction is called when production minuteFraction is entered.
func (s *Baseiso8601Listener) EnterMinuteFraction(ctx *MinuteFractionContext) {}

// ExitMinuteFraction is called when production minuteFraction is exited.
func (s *Baseiso8601Listener) ExitMinuteFraction(ctx *MinuteFractionContext) {}

// EnterSecondFraction is called when production secondFraction is entered.
func (s *Baseiso8601Listener) EnterSecondFraction(ctx *SecondFractionContext) {}

// ExitSecondFraction is called when production secondFraction is exited.
func (s *Baseiso8601Listener) ExitSecondFraction(ctx *SecondFractionContext) {}

// EnterCalendarDate is called when production calendarDate is entered.
func (s *Baseiso8601Listener) EnterCalendarDate(ctx *CalendarDateContext) {}

// ExitCalendarDate is called when production calendarDate is exited.
func (s *Baseiso8601Listener) ExitCalendarDate(ctx *CalendarDateContext) {}

// EnterCalendarDateBasic is called when production calendarDateBasic is entered.
func (s *Baseiso8601Listener) EnterCalendarDateBasic(ctx *CalendarDateBasicContext) {}

// ExitCalendarDateBasic is called when production calendarDateBasic is exited.
func (s *Baseiso8601Listener) ExitCalendarDateBasic(ctx *CalendarDateBasicContext) {}

// EnterCalendarDateExtended is called when production calendarDateExtended is entered.
func (s *Baseiso8601Listener) EnterCalendarDateExtended(ctx *CalendarDateExtendedContext) {}

// ExitCalendarDateExtended is called when production calendarDateExtended is exited.
func (s *Baseiso8601Listener) ExitCalendarDateExtended(ctx *CalendarDateExtendedContext) {}

// EnterSpecificMonthCalendarDate is called when production specificMonthCalendarDate is entered.
func (s *Baseiso8601Listener) EnterSpecificMonthCalendarDate(ctx *SpecificMonthCalendarDateContext) {}

// ExitSpecificMonthCalendarDate is called when production specificMonthCalendarDate is exited.
func (s *Baseiso8601Listener) ExitSpecificMonthCalendarDate(ctx *SpecificMonthCalendarDateContext) {}

// EnterSpecificYearCalendarDate is called when production specificYearCalendarDate is entered.
func (s *Baseiso8601Listener) EnterSpecificYearCalendarDate(ctx *SpecificYearCalendarDateContext) {}

// ExitSpecificYearCalendarDate is called when production specificYearCalendarDate is exited.
func (s *Baseiso8601Listener) ExitSpecificYearCalendarDate(ctx *SpecificYearCalendarDateContext) {}

// EnterSpecificCenturyCalendarDate is called when production specificCenturyCalendarDate is entered.
func (s *Baseiso8601Listener) EnterSpecificCenturyCalendarDate(ctx *SpecificCenturyCalendarDateContext) {
}

// ExitSpecificCenturyCalendarDate is called when production specificCenturyCalendarDate is exited.
func (s *Baseiso8601Listener) ExitSpecificCenturyCalendarDate(ctx *SpecificCenturyCalendarDateContext) {
}

// EnterOrdinalDate is called when production ordinalDate is entered.
func (s *Baseiso8601Listener) EnterOrdinalDate(ctx *OrdinalDateContext) {}

// ExitOrdinalDate is called when production ordinalDate is exited.
func (s *Baseiso8601Listener) ExitOrdinalDate(ctx *OrdinalDateContext) {}

// EnterOrdinalDateBasic is called when production ordinalDateBasic is entered.
func (s *Baseiso8601Listener) EnterOrdinalDateBasic(ctx *OrdinalDateBasicContext) {}

// ExitOrdinalDateBasic is called when production ordinalDateBasic is exited.
func (s *Baseiso8601Listener) ExitOrdinalDateBasic(ctx *OrdinalDateBasicContext) {}

// EnterOrdinalDateExtended is called when production ordinalDateExtended is entered.
func (s *Baseiso8601Listener) EnterOrdinalDateExtended(ctx *OrdinalDateExtendedContext) {}

// ExitOrdinalDateExtended is called when production ordinalDateExtended is exited.
func (s *Baseiso8601Listener) ExitOrdinalDateExtended(ctx *OrdinalDateExtendedContext) {}

// EnterWeekDate is called when production weekDate is entered.
func (s *Baseiso8601Listener) EnterWeekDate(ctx *WeekDateContext) {}

// ExitWeekDate is called when production weekDate is exited.
func (s *Baseiso8601Listener) ExitWeekDate(ctx *WeekDateContext) {}

// EnterWeekDateBasic is called when production weekDateBasic is entered.
func (s *Baseiso8601Listener) EnterWeekDateBasic(ctx *WeekDateBasicContext) {}

// ExitWeekDateBasic is called when production weekDateBasic is exited.
func (s *Baseiso8601Listener) ExitWeekDateBasic(ctx *WeekDateBasicContext) {}

// EnterWeekDateExtended is called when production weekDateExtended is entered.
func (s *Baseiso8601Listener) EnterWeekDateExtended(ctx *WeekDateExtendedContext) {}

// ExitWeekDateExtended is called when production weekDateExtended is exited.
func (s *Baseiso8601Listener) ExitWeekDateExtended(ctx *WeekDateExtendedContext) {}

// EnterSpecificWeekWeekDate is called when production specificWeekWeekDate is entered.
func (s *Baseiso8601Listener) EnterSpecificWeekWeekDate(ctx *SpecificWeekWeekDateContext) {}

// ExitSpecificWeekWeekDate is called when production specificWeekWeekDate is exited.
func (s *Baseiso8601Listener) ExitSpecificWeekWeekDate(ctx *SpecificWeekWeekDateContext) {}

// EnterSpecificWeekWeekDateBasic is called when production specificWeekWeekDateBasic is entered.
func (s *Baseiso8601Listener) EnterSpecificWeekWeekDateBasic(ctx *SpecificWeekWeekDateBasicContext) {}

// ExitSpecificWeekWeekDateBasic is called when production specificWeekWeekDateBasic is exited.
func (s *Baseiso8601Listener) ExitSpecificWeekWeekDateBasic(ctx *SpecificWeekWeekDateBasicContext) {}

// EnterSpecificWeekWeekDateExtended is called when production specificWeekWeekDateExtended is entered.
func (s *Baseiso8601Listener) EnterSpecificWeekWeekDateExtended(ctx *SpecificWeekWeekDateExtendedContext) {
}

// ExitSpecificWeekWeekDateExtended is called when production specificWeekWeekDateExtended is exited.
func (s *Baseiso8601Listener) ExitSpecificWeekWeekDateExtended(ctx *SpecificWeekWeekDateExtendedContext) {
}

// EnterDatePrecise is called when production datePrecise is entered.
func (s *Baseiso8601Listener) EnterDatePrecise(ctx *DatePreciseContext) {}

// ExitDatePrecise is called when production datePrecise is exited.
func (s *Baseiso8601Listener) ExitDatePrecise(ctx *DatePreciseContext) {}

// EnterDatePreciseBasic is called when production datePreciseBasic is entered.
func (s *Baseiso8601Listener) EnterDatePreciseBasic(ctx *DatePreciseBasicContext) {}

// ExitDatePreciseBasic is called when production datePreciseBasic is exited.
func (s *Baseiso8601Listener) ExitDatePreciseBasic(ctx *DatePreciseBasicContext) {}

// EnterDatePreciseExtended is called when production datePreciseExtended is entered.
func (s *Baseiso8601Listener) EnterDatePreciseExtended(ctx *DatePreciseExtendedContext) {}

// ExitDatePreciseExtended is called when production datePreciseExtended is exited.
func (s *Baseiso8601Listener) ExitDatePreciseExtended(ctx *DatePreciseExtendedContext) {}

// EnterDate is called when production date is entered.
func (s *Baseiso8601Listener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *Baseiso8601Listener) ExitDate(ctx *DateContext) {}

// EnterDateBasic is called when production dateBasic is entered.
func (s *Baseiso8601Listener) EnterDateBasic(ctx *DateBasicContext) {}

// ExitDateBasic is called when production dateBasic is exited.
func (s *Baseiso8601Listener) ExitDateBasic(ctx *DateBasicContext) {}

// EnterDateExtended is called when production dateExtended is entered.
func (s *Baseiso8601Listener) EnterDateExtended(ctx *DateExtendedContext) {}

// ExitDateExtended is called when production dateExtended is exited.
func (s *Baseiso8601Listener) ExitDateExtended(ctx *DateExtendedContext) {}

// EnterLocalTimePrecise is called when production localTimePrecise is entered.
func (s *Baseiso8601Listener) EnterLocalTimePrecise(ctx *LocalTimePreciseContext) {}

// ExitLocalTimePrecise is called when production localTimePrecise is exited.
func (s *Baseiso8601Listener) ExitLocalTimePrecise(ctx *LocalTimePreciseContext) {}

// EnterLocalTimePreciseBasic is called when production localTimePreciseBasic is entered.
func (s *Baseiso8601Listener) EnterLocalTimePreciseBasic(ctx *LocalTimePreciseBasicContext) {}

// ExitLocalTimePreciseBasic is called when production localTimePreciseBasic is exited.
func (s *Baseiso8601Listener) ExitLocalTimePreciseBasic(ctx *LocalTimePreciseBasicContext) {}

// EnterLocalTimePreciseExtended is called when production localTimePreciseExtended is entered.
func (s *Baseiso8601Listener) EnterLocalTimePreciseExtended(ctx *LocalTimePreciseExtendedContext) {}

// ExitLocalTimePreciseExtended is called when production localTimePreciseExtended is exited.
func (s *Baseiso8601Listener) ExitLocalTimePreciseExtended(ctx *LocalTimePreciseExtendedContext) {}

// EnterSpecificMinuteLocalTime is called when production specificMinuteLocalTime is entered.
func (s *Baseiso8601Listener) EnterSpecificMinuteLocalTime(ctx *SpecificMinuteLocalTimeContext) {}

// ExitSpecificMinuteLocalTime is called when production specificMinuteLocalTime is exited.
func (s *Baseiso8601Listener) ExitSpecificMinuteLocalTime(ctx *SpecificMinuteLocalTimeContext) {}

// EnterSpecificMinuteLocalTimeBasic is called when production specificMinuteLocalTimeBasic is entered.
func (s *Baseiso8601Listener) EnterSpecificMinuteLocalTimeBasic(ctx *SpecificMinuteLocalTimeBasicContext) {
}

// ExitSpecificMinuteLocalTimeBasic is called when production specificMinuteLocalTimeBasic is exited.
func (s *Baseiso8601Listener) ExitSpecificMinuteLocalTimeBasic(ctx *SpecificMinuteLocalTimeBasicContext) {
}

// EnterSpecificMinuteLocalTimeExtended is called when production specificMinuteLocalTimeExtended is entered.
func (s *Baseiso8601Listener) EnterSpecificMinuteLocalTimeExtended(ctx *SpecificMinuteLocalTimeExtendedContext) {
}

// ExitSpecificMinuteLocalTimeExtended is called when production specificMinuteLocalTimeExtended is exited.
func (s *Baseiso8601Listener) ExitSpecificMinuteLocalTimeExtended(ctx *SpecificMinuteLocalTimeExtendedContext) {
}

// EnterSpecificHourLocalTime is called when production specificHourLocalTime is entered.
func (s *Baseiso8601Listener) EnterSpecificHourLocalTime(ctx *SpecificHourLocalTimeContext) {}

// ExitSpecificHourLocalTime is called when production specificHourLocalTime is exited.
func (s *Baseiso8601Listener) ExitSpecificHourLocalTime(ctx *SpecificHourLocalTimeContext) {}

// EnterLocalTime is called when production localTime is entered.
func (s *Baseiso8601Listener) EnterLocalTime(ctx *LocalTimeContext) {}

// ExitLocalTime is called when production localTime is exited.
func (s *Baseiso8601Listener) ExitLocalTime(ctx *LocalTimeContext) {}

// EnterLocalTimeBasic is called when production localTimeBasic is entered.
func (s *Baseiso8601Listener) EnterLocalTimeBasic(ctx *LocalTimeBasicContext) {}

// ExitLocalTimeBasic is called when production localTimeBasic is exited.
func (s *Baseiso8601Listener) ExitLocalTimeBasic(ctx *LocalTimeBasicContext) {}

// EnterLocalTimeExtended is called when production localTimeExtended is entered.
func (s *Baseiso8601Listener) EnterLocalTimeExtended(ctx *LocalTimeExtendedContext) {}

// ExitLocalTimeExtended is called when production localTimeExtended is exited.
func (s *Baseiso8601Listener) ExitLocalTimeExtended(ctx *LocalTimeExtendedContext) {}

// EnterTimeZoneUtc is called when production timeZoneUtc is entered.
func (s *Baseiso8601Listener) EnterTimeZoneUtc(ctx *TimeZoneUtcContext) {}

// ExitTimeZoneUtc is called when production timeZoneUtc is exited.
func (s *Baseiso8601Listener) ExitTimeZoneUtc(ctx *TimeZoneUtcContext) {}

// EnterTimeZone is called when production timeZone is entered.
func (s *Baseiso8601Listener) EnterTimeZone(ctx *TimeZoneContext) {}

// ExitTimeZone is called when production timeZone is exited.
func (s *Baseiso8601Listener) ExitTimeZone(ctx *TimeZoneContext) {}

// EnterTimeZoneBasic is called when production timeZoneBasic is entered.
func (s *Baseiso8601Listener) EnterTimeZoneBasic(ctx *TimeZoneBasicContext) {}

// ExitTimeZoneBasic is called when production timeZoneBasic is exited.
func (s *Baseiso8601Listener) ExitTimeZoneBasic(ctx *TimeZoneBasicContext) {}

// EnterTimeZoneExtended is called when production timeZoneExtended is entered.
func (s *Baseiso8601Listener) EnterTimeZoneExtended(ctx *TimeZoneExtendedContext) {}

// ExitTimeZoneExtended is called when production timeZoneExtended is exited.
func (s *Baseiso8601Listener) ExitTimeZoneExtended(ctx *TimeZoneExtendedContext) {}

// EnterLocalTimeAndTimeZone is called when production localTimeAndTimeZone is entered.
func (s *Baseiso8601Listener) EnterLocalTimeAndTimeZone(ctx *LocalTimeAndTimeZoneContext) {}

// ExitLocalTimeAndTimeZone is called when production localTimeAndTimeZone is exited.
func (s *Baseiso8601Listener) ExitLocalTimeAndTimeZone(ctx *LocalTimeAndTimeZoneContext) {}

// EnterLocalTimeAndTimeZoneBasic is called when production localTimeAndTimeZoneBasic is entered.
func (s *Baseiso8601Listener) EnterLocalTimeAndTimeZoneBasic(ctx *LocalTimeAndTimeZoneBasicContext) {}

// ExitLocalTimeAndTimeZoneBasic is called when production localTimeAndTimeZoneBasic is exited.
func (s *Baseiso8601Listener) ExitLocalTimeAndTimeZoneBasic(ctx *LocalTimeAndTimeZoneBasicContext) {}

// EnterLocalTimeAndTimeZoneExtended is called when production localTimeAndTimeZoneExtended is entered.
func (s *Baseiso8601Listener) EnterLocalTimeAndTimeZoneExtended(ctx *LocalTimeAndTimeZoneExtendedContext) {
}

// ExitLocalTimeAndTimeZoneExtended is called when production localTimeAndTimeZoneExtended is exited.
func (s *Baseiso8601Listener) ExitLocalTimeAndTimeZoneExtended(ctx *LocalTimeAndTimeZoneExtendedContext) {
}

// EnterTime is called when production time is entered.
func (s *Baseiso8601Listener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *Baseiso8601Listener) ExitTime(ctx *TimeContext) {}

// EnterDatetimePrecise is called when production datetimePrecise is entered.
func (s *Baseiso8601Listener) EnterDatetimePrecise(ctx *DatetimePreciseContext) {}

// ExitDatetimePrecise is called when production datetimePrecise is exited.
func (s *Baseiso8601Listener) ExitDatetimePrecise(ctx *DatetimePreciseContext) {}

// EnterDatetimePreciseBasic is called when production datetimePreciseBasic is entered.
func (s *Baseiso8601Listener) EnterDatetimePreciseBasic(ctx *DatetimePreciseBasicContext) {}

// ExitDatetimePreciseBasic is called when production datetimePreciseBasic is exited.
func (s *Baseiso8601Listener) ExitDatetimePreciseBasic(ctx *DatetimePreciseBasicContext) {}

// EnterDatetimePreciseExtended is called when production datetimePreciseExtended is entered.
func (s *Baseiso8601Listener) EnterDatetimePreciseExtended(ctx *DatetimePreciseExtendedContext) {}

// ExitDatetimePreciseExtended is called when production datetimePreciseExtended is exited.
func (s *Baseiso8601Listener) ExitDatetimePreciseExtended(ctx *DatetimePreciseExtendedContext) {}

// EnterDatetime is called when production datetime is entered.
func (s *Baseiso8601Listener) EnterDatetime(ctx *DatetimeContext) {}

// ExitDatetime is called when production datetime is exited.
func (s *Baseiso8601Listener) ExitDatetime(ctx *DatetimeContext) {}

// EnterDatetimeBasic is called when production datetimeBasic is entered.
func (s *Baseiso8601Listener) EnterDatetimeBasic(ctx *DatetimeBasicContext) {}

// ExitDatetimeBasic is called when production datetimeBasic is exited.
func (s *Baseiso8601Listener) ExitDatetimeBasic(ctx *DatetimeBasicContext) {}

// EnterDatetimeExtended is called when production datetimeExtended is entered.
func (s *Baseiso8601Listener) EnterDatetimeExtended(ctx *DatetimeExtendedContext) {}

// ExitDatetimeExtended is called when production datetimeExtended is exited.
func (s *Baseiso8601Listener) ExitDatetimeExtended(ctx *DatetimeExtendedContext) {}

// EnterInterval is called when production interval is entered.
func (s *Baseiso8601Listener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *Baseiso8601Listener) ExitInterval(ctx *IntervalContext) {}

// EnterIntervalBasic is called when production intervalBasic is entered.
func (s *Baseiso8601Listener) EnterIntervalBasic(ctx *IntervalBasicContext) {}

// ExitIntervalBasic is called when production intervalBasic is exited.
func (s *Baseiso8601Listener) ExitIntervalBasic(ctx *IntervalBasicContext) {}

// EnterIntervalExtended is called when production intervalExtended is entered.
func (s *Baseiso8601Listener) EnterIntervalExtended(ctx *IntervalExtendedContext) {}

// ExitIntervalExtended is called when production intervalExtended is exited.
func (s *Baseiso8601Listener) ExitIntervalExtended(ctx *IntervalExtendedContext) {}

// EnterIntervalT is called when production intervalT is entered.
func (s *Baseiso8601Listener) EnterIntervalT(ctx *IntervalTContext) {}

// ExitIntervalT is called when production intervalT is exited.
func (s *Baseiso8601Listener) ExitIntervalT(ctx *IntervalTContext) {}

// EnterMonthDateBasic is called when production monthDateBasic is entered.
func (s *Baseiso8601Listener) EnterMonthDateBasic(ctx *MonthDateBasicContext) {}

// ExitMonthDateBasic is called when production monthDateBasic is exited.
func (s *Baseiso8601Listener) ExitMonthDateBasic(ctx *MonthDateBasicContext) {}

// EnterMonthDateExtended is called when production monthDateExtended is entered.
func (s *Baseiso8601Listener) EnterMonthDateExtended(ctx *MonthDateExtendedContext) {}

// ExitMonthDateExtended is called when production monthDateExtended is exited.
func (s *Baseiso8601Listener) ExitMonthDateExtended(ctx *MonthDateExtendedContext) {}

// EnterIntervalYMDTimeBasic is called when production intervalYMDTimeBasic is entered.
func (s *Baseiso8601Listener) EnterIntervalYMDTimeBasic(ctx *IntervalYMDTimeBasicContext) {}

// ExitIntervalYMDTimeBasic is called when production intervalYMDTimeBasic is exited.
func (s *Baseiso8601Listener) ExitIntervalYMDTimeBasic(ctx *IntervalYMDTimeBasicContext) {}

// EnterIntervalYMDTimeExtended is called when production intervalYMDTimeExtended is entered.
func (s *Baseiso8601Listener) EnterIntervalYMDTimeExtended(ctx *IntervalYMDTimeExtendedContext) {}

// ExitIntervalYMDTimeExtended is called when production intervalYMDTimeExtended is exited.
func (s *Baseiso8601Listener) ExitIntervalYMDTimeExtended(ctx *IntervalYMDTimeExtendedContext) {}

// EnterIntervalYMD is called when production intervalYMD is entered.
func (s *Baseiso8601Listener) EnterIntervalYMD(ctx *IntervalYMDContext) {}

// ExitIntervalYMD is called when production intervalYMD is exited.
func (s *Baseiso8601Listener) ExitIntervalYMD(ctx *IntervalYMDContext) {}

// EnterIntervalYMDBasic is called when production intervalYMDBasic is entered.
func (s *Baseiso8601Listener) EnterIntervalYMDBasic(ctx *IntervalYMDBasicContext) {}

// ExitIntervalYMDBasic is called when production intervalYMDBasic is exited.
func (s *Baseiso8601Listener) ExitIntervalYMDBasic(ctx *IntervalYMDBasicContext) {}

// EnterIntervalYMDExtended is called when production intervalYMDExtended is entered.
func (s *Baseiso8601Listener) EnterIntervalYMDExtended(ctx *IntervalYMDExtendedContext) {}

// ExitIntervalYMDExtended is called when production intervalYMDExtended is exited.
func (s *Baseiso8601Listener) ExitIntervalYMDExtended(ctx *IntervalYMDExtendedContext) {}

// EnterTimeBeginEnd is called when production timeBeginEnd is entered.
func (s *Baseiso8601Listener) EnterTimeBeginEnd(ctx *TimeBeginEndContext) {}

// ExitTimeBeginEnd is called when production timeBeginEnd is exited.
func (s *Baseiso8601Listener) ExitTimeBeginEnd(ctx *TimeBeginEndContext) {}

// EnterTimeBeginEndBasic is called when production timeBeginEndBasic is entered.
func (s *Baseiso8601Listener) EnterTimeBeginEndBasic(ctx *TimeBeginEndBasicContext) {}

// ExitTimeBeginEndBasic is called when production timeBeginEndBasic is exited.
func (s *Baseiso8601Listener) ExitTimeBeginEndBasic(ctx *TimeBeginEndBasicContext) {}

// EnterTimeBeginEndExtended is called when production timeBeginEndExtended is entered.
func (s *Baseiso8601Listener) EnterTimeBeginEndExtended(ctx *TimeBeginEndExtendedContext) {}

// ExitTimeBeginEndExtended is called when production timeBeginEndExtended is exited.
func (s *Baseiso8601Listener) ExitTimeBeginEndExtended(ctx *TimeBeginEndExtendedContext) {}

// EnterTimeBeginInterval is called when production timeBeginInterval is entered.
func (s *Baseiso8601Listener) EnterTimeBeginInterval(ctx *TimeBeginIntervalContext) {}

// ExitTimeBeginInterval is called when production timeBeginInterval is exited.
func (s *Baseiso8601Listener) ExitTimeBeginInterval(ctx *TimeBeginIntervalContext) {}

// EnterTimeBeginIntervalBasic is called when production timeBeginIntervalBasic is entered.
func (s *Baseiso8601Listener) EnterTimeBeginIntervalBasic(ctx *TimeBeginIntervalBasicContext) {}

// ExitTimeBeginIntervalBasic is called when production timeBeginIntervalBasic is exited.
func (s *Baseiso8601Listener) ExitTimeBeginIntervalBasic(ctx *TimeBeginIntervalBasicContext) {}

// EnterTimeBeginIntervalExtended is called when production timeBeginIntervalExtended is entered.
func (s *Baseiso8601Listener) EnterTimeBeginIntervalExtended(ctx *TimeBeginIntervalExtendedContext) {}

// ExitTimeBeginIntervalExtended is called when production timeBeginIntervalExtended is exited.
func (s *Baseiso8601Listener) ExitTimeBeginIntervalExtended(ctx *TimeBeginIntervalExtendedContext) {}

// EnterTimeIntervalEnd is called when production timeIntervalEnd is entered.
func (s *Baseiso8601Listener) EnterTimeIntervalEnd(ctx *TimeIntervalEndContext) {}

// ExitTimeIntervalEnd is called when production timeIntervalEnd is exited.
func (s *Baseiso8601Listener) ExitTimeIntervalEnd(ctx *TimeIntervalEndContext) {}

// EnterTimeIntervalEndBasic is called when production timeIntervalEndBasic is entered.
func (s *Baseiso8601Listener) EnterTimeIntervalEndBasic(ctx *TimeIntervalEndBasicContext) {}

// ExitTimeIntervalEndBasic is called when production timeIntervalEndBasic is exited.
func (s *Baseiso8601Listener) ExitTimeIntervalEndBasic(ctx *TimeIntervalEndBasicContext) {}

// EnterTimeIntervalEndExtended is called when production timeIntervalEndExtended is entered.
func (s *Baseiso8601Listener) EnterTimeIntervalEndExtended(ctx *TimeIntervalEndExtendedContext) {}

// ExitTimeIntervalEndExtended is called when production timeIntervalEndExtended is exited.
func (s *Baseiso8601Listener) ExitTimeIntervalEndExtended(ctx *TimeIntervalEndExtendedContext) {}

// EnterDuration is called when production duration is entered.
func (s *Baseiso8601Listener) EnterDuration(ctx *DurationContext) {}

// ExitDuration is called when production duration is exited.
func (s *Baseiso8601Listener) ExitDuration(ctx *DurationContext) {}

// EnterDurationBasic is called when production durationBasic is entered.
func (s *Baseiso8601Listener) EnterDurationBasic(ctx *DurationBasicContext) {}

// ExitDurationBasic is called when production durationBasic is exited.
func (s *Baseiso8601Listener) ExitDurationBasic(ctx *DurationBasicContext) {}

// EnterDurationExtended is called when production durationExtended is entered.
func (s *Baseiso8601Listener) EnterDurationExtended(ctx *DurationExtendedContext) {}

// ExitDurationExtended is called when production durationExtended is exited.
func (s *Baseiso8601Listener) ExitDurationExtended(ctx *DurationExtendedContext) {}

// EnterRecurringCount is called when production recurringCount is entered.
func (s *Baseiso8601Listener) EnterRecurringCount(ctx *RecurringCountContext) {}

// ExitRecurringCount is called when production recurringCount is exited.
func (s *Baseiso8601Listener) ExitRecurringCount(ctx *RecurringCountContext) {}

// EnterRecurring is called when production recurring is entered.
func (s *Baseiso8601Listener) EnterRecurring(ctx *RecurringContext) {}

// ExitRecurring is called when production recurring is exited.
func (s *Baseiso8601Listener) ExitRecurring(ctx *RecurringContext) {}

// EnterRecurringBasic is called when production recurringBasic is entered.
func (s *Baseiso8601Listener) EnterRecurringBasic(ctx *RecurringBasicContext) {}

// ExitRecurringBasic is called when production recurringBasic is exited.
func (s *Baseiso8601Listener) ExitRecurringBasic(ctx *RecurringBasicContext) {}

// EnterRecurringExtended is called when production recurringExtended is entered.
func (s *Baseiso8601Listener) EnterRecurringExtended(ctx *RecurringExtendedContext) {}

// ExitRecurringExtended is called when production recurringExtended is exited.
func (s *Baseiso8601Listener) ExitRecurringExtended(ctx *RecurringExtendedContext) {}

// EnterIso is called when production iso is entered.
func (s *Baseiso8601Listener) EnterIso(ctx *IsoContext) {}

// ExitIso is called when production iso is exited.
func (s *Baseiso8601Listener) ExitIso(ctx *IsoContext) {}
