// Code generated from iso8601.g4 by ANTLR 4.9.3. DO NOT EDIT.

package iso8601 // iso8601
import "github.com/antlr/antlr4/runtime/Go/antlr"

// iso8601Listener is a complete listener for a parse tree produced by iso8601Parser.
type iso8601Listener interface {
	antlr.ParseTreeListener

	// EnterInt_ is called when entering the int_ production.
	EnterInt_(c *Int_Context)

	// EnterDec is called when entering the dec production.
	EnterDec(c *DecContext)

	// EnterInt2 is called when entering the int2 production.
	EnterInt2(c *Int2Context)

	// EnterInt3 is called when entering the int3 production.
	EnterInt3(c *Int3Context)

	// EnterInt4 is called when entering the int4 production.
	EnterInt4(c *Int4Context)

	// EnterSint2p is called when entering the sint2p production.
	EnterSint2p(c *Sint2pContext)

	// EnterSint4p is called when entering the sint4p production.
	EnterSint4p(c *Sint4pContext)

	// EnterDec2 is called when entering the dec2 production.
	EnterDec2(c *Dec2Context)

	// EnterCompleteCentury is called when entering the CompleteCentury production.
	EnterCompleteCentury(c *CompleteCenturyContext)

	// EnterExpandedCentury is called when entering the ExpandedCentury production.
	EnterExpandedCentury(c *ExpandedCenturyContext)

	// EnterCompleteYear is called when entering the CompleteYear production.
	EnterCompleteYear(c *CompleteYearContext)

	// EnterExpandedYear is called when entering the ExpandedYear production.
	EnterExpandedYear(c *ExpandedYearContext)

	// EnterMonth is called when entering the month production.
	EnterMonth(c *MonthContext)

	// EnterDay is called when entering the day production.
	EnterDay(c *DayContext)

	// EnterOrdinalDay is called when entering the ordinalDay production.
	EnterOrdinalDay(c *OrdinalDayContext)

	// EnterWeek is called when entering the week production.
	EnterWeek(c *WeekContext)

	// EnterWeekDay is called when entering the weekDay production.
	EnterWeekDay(c *WeekDayContext)

	// EnterHour is called when entering the hour production.
	EnterHour(c *HourContext)

	// EnterMinute is called when entering the minute production.
	EnterMinute(c *MinuteContext)

	// EnterSecond is called when entering the second production.
	EnterSecond(c *SecondContext)

	// EnterHourFraction is called when entering the hourFraction production.
	EnterHourFraction(c *HourFractionContext)

	// EnterMinuteFraction is called when entering the minuteFraction production.
	EnterMinuteFraction(c *MinuteFractionContext)

	// EnterSecondFraction is called when entering the secondFraction production.
	EnterSecondFraction(c *SecondFractionContext)

	// EnterCalendarDate is called when entering the calendarDate production.
	EnterCalendarDate(c *CalendarDateContext)

	// EnterCalendarDateBasic is called when entering the calendarDateBasic production.
	EnterCalendarDateBasic(c *CalendarDateBasicContext)

	// EnterCalendarDateExtended is called when entering the calendarDateExtended production.
	EnterCalendarDateExtended(c *CalendarDateExtendedContext)

	// EnterSpecificMonthCalendarDate is called when entering the specificMonthCalendarDate production.
	EnterSpecificMonthCalendarDate(c *SpecificMonthCalendarDateContext)

	// EnterSpecificYearCalendarDate is called when entering the specificYearCalendarDate production.
	EnterSpecificYearCalendarDate(c *SpecificYearCalendarDateContext)

	// EnterSpecificCenturyCalendarDate is called when entering the specificCenturyCalendarDate production.
	EnterSpecificCenturyCalendarDate(c *SpecificCenturyCalendarDateContext)

	// EnterOrdinalDate is called when entering the ordinalDate production.
	EnterOrdinalDate(c *OrdinalDateContext)

	// EnterOrdinalDateBasic is called when entering the ordinalDateBasic production.
	EnterOrdinalDateBasic(c *OrdinalDateBasicContext)

	// EnterOrdinalDateExtended is called when entering the ordinalDateExtended production.
	EnterOrdinalDateExtended(c *OrdinalDateExtendedContext)

	// EnterWeekDate is called when entering the weekDate production.
	EnterWeekDate(c *WeekDateContext)

	// EnterWeekDateBasic is called when entering the weekDateBasic production.
	EnterWeekDateBasic(c *WeekDateBasicContext)

	// EnterWeekDateExtended is called when entering the weekDateExtended production.
	EnterWeekDateExtended(c *WeekDateExtendedContext)

	// EnterSpecificWeekWeekDate is called when entering the specificWeekWeekDate production.
	EnterSpecificWeekWeekDate(c *SpecificWeekWeekDateContext)

	// EnterSpecificWeekWeekDateBasic is called when entering the specificWeekWeekDateBasic production.
	EnterSpecificWeekWeekDateBasic(c *SpecificWeekWeekDateBasicContext)

	// EnterSpecificWeekWeekDateExtended is called when entering the specificWeekWeekDateExtended production.
	EnterSpecificWeekWeekDateExtended(c *SpecificWeekWeekDateExtendedContext)

	// EnterDatePrecise is called when entering the datePrecise production.
	EnterDatePrecise(c *DatePreciseContext)

	// EnterDatePreciseBasic is called when entering the datePreciseBasic production.
	EnterDatePreciseBasic(c *DatePreciseBasicContext)

	// EnterDatePreciseExtended is called when entering the datePreciseExtended production.
	EnterDatePreciseExtended(c *DatePreciseExtendedContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterDateBasic is called when entering the dateBasic production.
	EnterDateBasic(c *DateBasicContext)

	// EnterDateExtended is called when entering the dateExtended production.
	EnterDateExtended(c *DateExtendedContext)

	// EnterLocalTimePrecise is called when entering the localTimePrecise production.
	EnterLocalTimePrecise(c *LocalTimePreciseContext)

	// EnterLocalTimePreciseBasic is called when entering the localTimePreciseBasic production.
	EnterLocalTimePreciseBasic(c *LocalTimePreciseBasicContext)

	// EnterLocalTimePreciseExtended is called when entering the localTimePreciseExtended production.
	EnterLocalTimePreciseExtended(c *LocalTimePreciseExtendedContext)

	// EnterSpecificMinuteLocalTime is called when entering the specificMinuteLocalTime production.
	EnterSpecificMinuteLocalTime(c *SpecificMinuteLocalTimeContext)

	// EnterSpecificMinuteLocalTimeBasic is called when entering the specificMinuteLocalTimeBasic production.
	EnterSpecificMinuteLocalTimeBasic(c *SpecificMinuteLocalTimeBasicContext)

	// EnterSpecificMinuteLocalTimeExtended is called when entering the specificMinuteLocalTimeExtended production.
	EnterSpecificMinuteLocalTimeExtended(c *SpecificMinuteLocalTimeExtendedContext)

	// EnterSpecificHourLocalTime is called when entering the specificHourLocalTime production.
	EnterSpecificHourLocalTime(c *SpecificHourLocalTimeContext)

	// EnterLocalTime is called when entering the localTime production.
	EnterLocalTime(c *LocalTimeContext)

	// EnterLocalTimeBasic is called when entering the localTimeBasic production.
	EnterLocalTimeBasic(c *LocalTimeBasicContext)

	// EnterLocalTimeExtended is called when entering the localTimeExtended production.
	EnterLocalTimeExtended(c *LocalTimeExtendedContext)

	// EnterTimeZoneUtc is called when entering the timeZoneUtc production.
	EnterTimeZoneUtc(c *TimeZoneUtcContext)

	// EnterTimeZone is called when entering the timeZone production.
	EnterTimeZone(c *TimeZoneContext)

	// EnterTimeZoneBasic is called when entering the timeZoneBasic production.
	EnterTimeZoneBasic(c *TimeZoneBasicContext)

	// EnterTimeZoneExtended is called when entering the timeZoneExtended production.
	EnterTimeZoneExtended(c *TimeZoneExtendedContext)

	// EnterLocalTimeAndTimeZone is called when entering the localTimeAndTimeZone production.
	EnterLocalTimeAndTimeZone(c *LocalTimeAndTimeZoneContext)

	// EnterLocalTimeAndTimeZoneBasic is called when entering the localTimeAndTimeZoneBasic production.
	EnterLocalTimeAndTimeZoneBasic(c *LocalTimeAndTimeZoneBasicContext)

	// EnterLocalTimeAndTimeZoneExtended is called when entering the localTimeAndTimeZoneExtended production.
	EnterLocalTimeAndTimeZoneExtended(c *LocalTimeAndTimeZoneExtendedContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// EnterDatetimePrecise is called when entering the datetimePrecise production.
	EnterDatetimePrecise(c *DatetimePreciseContext)

	// EnterDatetimePreciseBasic is called when entering the datetimePreciseBasic production.
	EnterDatetimePreciseBasic(c *DatetimePreciseBasicContext)

	// EnterDatetimePreciseExtended is called when entering the datetimePreciseExtended production.
	EnterDatetimePreciseExtended(c *DatetimePreciseExtendedContext)

	// EnterDatetime is called when entering the datetime production.
	EnterDatetime(c *DatetimeContext)

	// EnterDatetimeBasic is called when entering the datetimeBasic production.
	EnterDatetimeBasic(c *DatetimeBasicContext)

	// EnterDatetimeExtended is called when entering the datetimeExtended production.
	EnterDatetimeExtended(c *DatetimeExtendedContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterIntervalBasic is called when entering the intervalBasic production.
	EnterIntervalBasic(c *IntervalBasicContext)

	// EnterIntervalExtended is called when entering the intervalExtended production.
	EnterIntervalExtended(c *IntervalExtendedContext)

	// EnterIntervalT is called when entering the intervalT production.
	EnterIntervalT(c *IntervalTContext)

	// EnterMonthDateBasic is called when entering the monthDateBasic production.
	EnterMonthDateBasic(c *MonthDateBasicContext)

	// EnterMonthDateExtended is called when entering the monthDateExtended production.
	EnterMonthDateExtended(c *MonthDateExtendedContext)

	// EnterIntervalYMDTimeBasic is called when entering the intervalYMDTimeBasic production.
	EnterIntervalYMDTimeBasic(c *IntervalYMDTimeBasicContext)

	// EnterIntervalYMDTimeExtended is called when entering the intervalYMDTimeExtended production.
	EnterIntervalYMDTimeExtended(c *IntervalYMDTimeExtendedContext)

	// EnterIntervalYMD is called when entering the intervalYMD production.
	EnterIntervalYMD(c *IntervalYMDContext)

	// EnterIntervalYMDBasic is called when entering the intervalYMDBasic production.
	EnterIntervalYMDBasic(c *IntervalYMDBasicContext)

	// EnterIntervalYMDExtended is called when entering the intervalYMDExtended production.
	EnterIntervalYMDExtended(c *IntervalYMDExtendedContext)

	// EnterTimeBeginEnd is called when entering the timeBeginEnd production.
	EnterTimeBeginEnd(c *TimeBeginEndContext)

	// EnterTimeBeginEndBasic is called when entering the timeBeginEndBasic production.
	EnterTimeBeginEndBasic(c *TimeBeginEndBasicContext)

	// EnterTimeBeginEndExtended is called when entering the timeBeginEndExtended production.
	EnterTimeBeginEndExtended(c *TimeBeginEndExtendedContext)

	// EnterTimeBeginInterval is called when entering the timeBeginInterval production.
	EnterTimeBeginInterval(c *TimeBeginIntervalContext)

	// EnterTimeBeginIntervalBasic is called when entering the timeBeginIntervalBasic production.
	EnterTimeBeginIntervalBasic(c *TimeBeginIntervalBasicContext)

	// EnterTimeBeginIntervalExtended is called when entering the timeBeginIntervalExtended production.
	EnterTimeBeginIntervalExtended(c *TimeBeginIntervalExtendedContext)

	// EnterTimeIntervalEnd is called when entering the timeIntervalEnd production.
	EnterTimeIntervalEnd(c *TimeIntervalEndContext)

	// EnterTimeIntervalEndBasic is called when entering the timeIntervalEndBasic production.
	EnterTimeIntervalEndBasic(c *TimeIntervalEndBasicContext)

	// EnterTimeIntervalEndExtended is called when entering the timeIntervalEndExtended production.
	EnterTimeIntervalEndExtended(c *TimeIntervalEndExtendedContext)

	// EnterDuration is called when entering the duration production.
	EnterDuration(c *DurationContext)

	// EnterDurationBasic is called when entering the durationBasic production.
	EnterDurationBasic(c *DurationBasicContext)

	// EnterDurationExtended is called when entering the durationExtended production.
	EnterDurationExtended(c *DurationExtendedContext)

	// EnterRecurringCount is called when entering the recurringCount production.
	EnterRecurringCount(c *RecurringCountContext)

	// EnterRecurring is called when entering the recurring production.
	EnterRecurring(c *RecurringContext)

	// EnterRecurringBasic is called when entering the recurringBasic production.
	EnterRecurringBasic(c *RecurringBasicContext)

	// EnterRecurringExtended is called when entering the recurringExtended production.
	EnterRecurringExtended(c *RecurringExtendedContext)

	// EnterIso is called when entering the iso production.
	EnterIso(c *IsoContext)

	// ExitInt_ is called when exiting the int_ production.
	ExitInt_(c *Int_Context)

	// ExitDec is called when exiting the dec production.
	ExitDec(c *DecContext)

	// ExitInt2 is called when exiting the int2 production.
	ExitInt2(c *Int2Context)

	// ExitInt3 is called when exiting the int3 production.
	ExitInt3(c *Int3Context)

	// ExitInt4 is called when exiting the int4 production.
	ExitInt4(c *Int4Context)

	// ExitSint2p is called when exiting the sint2p production.
	ExitSint2p(c *Sint2pContext)

	// ExitSint4p is called when exiting the sint4p production.
	ExitSint4p(c *Sint4pContext)

	// ExitDec2 is called when exiting the dec2 production.
	ExitDec2(c *Dec2Context)

	// ExitCompleteCentury is called when exiting the CompleteCentury production.
	ExitCompleteCentury(c *CompleteCenturyContext)

	// ExitExpandedCentury is called when exiting the ExpandedCentury production.
	ExitExpandedCentury(c *ExpandedCenturyContext)

	// ExitCompleteYear is called when exiting the CompleteYear production.
	ExitCompleteYear(c *CompleteYearContext)

	// ExitExpandedYear is called when exiting the ExpandedYear production.
	ExitExpandedYear(c *ExpandedYearContext)

	// ExitMonth is called when exiting the month production.
	ExitMonth(c *MonthContext)

	// ExitDay is called when exiting the day production.
	ExitDay(c *DayContext)

	// ExitOrdinalDay is called when exiting the ordinalDay production.
	ExitOrdinalDay(c *OrdinalDayContext)

	// ExitWeek is called when exiting the week production.
	ExitWeek(c *WeekContext)

	// ExitWeekDay is called when exiting the weekDay production.
	ExitWeekDay(c *WeekDayContext)

	// ExitHour is called when exiting the hour production.
	ExitHour(c *HourContext)

	// ExitMinute is called when exiting the minute production.
	ExitMinute(c *MinuteContext)

	// ExitSecond is called when exiting the second production.
	ExitSecond(c *SecondContext)

	// ExitHourFraction is called when exiting the hourFraction production.
	ExitHourFraction(c *HourFractionContext)

	// ExitMinuteFraction is called when exiting the minuteFraction production.
	ExitMinuteFraction(c *MinuteFractionContext)

	// ExitSecondFraction is called when exiting the secondFraction production.
	ExitSecondFraction(c *SecondFractionContext)

	// ExitCalendarDate is called when exiting the calendarDate production.
	ExitCalendarDate(c *CalendarDateContext)

	// ExitCalendarDateBasic is called when exiting the calendarDateBasic production.
	ExitCalendarDateBasic(c *CalendarDateBasicContext)

	// ExitCalendarDateExtended is called when exiting the calendarDateExtended production.
	ExitCalendarDateExtended(c *CalendarDateExtendedContext)

	// ExitSpecificMonthCalendarDate is called when exiting the specificMonthCalendarDate production.
	ExitSpecificMonthCalendarDate(c *SpecificMonthCalendarDateContext)

	// ExitSpecificYearCalendarDate is called when exiting the specificYearCalendarDate production.
	ExitSpecificYearCalendarDate(c *SpecificYearCalendarDateContext)

	// ExitSpecificCenturyCalendarDate is called when exiting the specificCenturyCalendarDate production.
	ExitSpecificCenturyCalendarDate(c *SpecificCenturyCalendarDateContext)

	// ExitOrdinalDate is called when exiting the ordinalDate production.
	ExitOrdinalDate(c *OrdinalDateContext)

	// ExitOrdinalDateBasic is called when exiting the ordinalDateBasic production.
	ExitOrdinalDateBasic(c *OrdinalDateBasicContext)

	// ExitOrdinalDateExtended is called when exiting the ordinalDateExtended production.
	ExitOrdinalDateExtended(c *OrdinalDateExtendedContext)

	// ExitWeekDate is called when exiting the weekDate production.
	ExitWeekDate(c *WeekDateContext)

	// ExitWeekDateBasic is called when exiting the weekDateBasic production.
	ExitWeekDateBasic(c *WeekDateBasicContext)

	// ExitWeekDateExtended is called when exiting the weekDateExtended production.
	ExitWeekDateExtended(c *WeekDateExtendedContext)

	// ExitSpecificWeekWeekDate is called when exiting the specificWeekWeekDate production.
	ExitSpecificWeekWeekDate(c *SpecificWeekWeekDateContext)

	// ExitSpecificWeekWeekDateBasic is called when exiting the specificWeekWeekDateBasic production.
	ExitSpecificWeekWeekDateBasic(c *SpecificWeekWeekDateBasicContext)

	// ExitSpecificWeekWeekDateExtended is called when exiting the specificWeekWeekDateExtended production.
	ExitSpecificWeekWeekDateExtended(c *SpecificWeekWeekDateExtendedContext)

	// ExitDatePrecise is called when exiting the datePrecise production.
	ExitDatePrecise(c *DatePreciseContext)

	// ExitDatePreciseBasic is called when exiting the datePreciseBasic production.
	ExitDatePreciseBasic(c *DatePreciseBasicContext)

	// ExitDatePreciseExtended is called when exiting the datePreciseExtended production.
	ExitDatePreciseExtended(c *DatePreciseExtendedContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitDateBasic is called when exiting the dateBasic production.
	ExitDateBasic(c *DateBasicContext)

	// ExitDateExtended is called when exiting the dateExtended production.
	ExitDateExtended(c *DateExtendedContext)

	// ExitLocalTimePrecise is called when exiting the localTimePrecise production.
	ExitLocalTimePrecise(c *LocalTimePreciseContext)

	// ExitLocalTimePreciseBasic is called when exiting the localTimePreciseBasic production.
	ExitLocalTimePreciseBasic(c *LocalTimePreciseBasicContext)

	// ExitLocalTimePreciseExtended is called when exiting the localTimePreciseExtended production.
	ExitLocalTimePreciseExtended(c *LocalTimePreciseExtendedContext)

	// ExitSpecificMinuteLocalTime is called when exiting the specificMinuteLocalTime production.
	ExitSpecificMinuteLocalTime(c *SpecificMinuteLocalTimeContext)

	// ExitSpecificMinuteLocalTimeBasic is called when exiting the specificMinuteLocalTimeBasic production.
	ExitSpecificMinuteLocalTimeBasic(c *SpecificMinuteLocalTimeBasicContext)

	// ExitSpecificMinuteLocalTimeExtended is called when exiting the specificMinuteLocalTimeExtended production.
	ExitSpecificMinuteLocalTimeExtended(c *SpecificMinuteLocalTimeExtendedContext)

	// ExitSpecificHourLocalTime is called when exiting the specificHourLocalTime production.
	ExitSpecificHourLocalTime(c *SpecificHourLocalTimeContext)

	// ExitLocalTime is called when exiting the localTime production.
	ExitLocalTime(c *LocalTimeContext)

	// ExitLocalTimeBasic is called when exiting the localTimeBasic production.
	ExitLocalTimeBasic(c *LocalTimeBasicContext)

	// ExitLocalTimeExtended is called when exiting the localTimeExtended production.
	ExitLocalTimeExtended(c *LocalTimeExtendedContext)

	// ExitTimeZoneUtc is called when exiting the timeZoneUtc production.
	ExitTimeZoneUtc(c *TimeZoneUtcContext)

	// ExitTimeZone is called when exiting the timeZone production.
	ExitTimeZone(c *TimeZoneContext)

	// ExitTimeZoneBasic is called when exiting the timeZoneBasic production.
	ExitTimeZoneBasic(c *TimeZoneBasicContext)

	// ExitTimeZoneExtended is called when exiting the timeZoneExtended production.
	ExitTimeZoneExtended(c *TimeZoneExtendedContext)

	// ExitLocalTimeAndTimeZone is called when exiting the localTimeAndTimeZone production.
	ExitLocalTimeAndTimeZone(c *LocalTimeAndTimeZoneContext)

	// ExitLocalTimeAndTimeZoneBasic is called when exiting the localTimeAndTimeZoneBasic production.
	ExitLocalTimeAndTimeZoneBasic(c *LocalTimeAndTimeZoneBasicContext)

	// ExitLocalTimeAndTimeZoneExtended is called when exiting the localTimeAndTimeZoneExtended production.
	ExitLocalTimeAndTimeZoneExtended(c *LocalTimeAndTimeZoneExtendedContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)

	// ExitDatetimePrecise is called when exiting the datetimePrecise production.
	ExitDatetimePrecise(c *DatetimePreciseContext)

	// ExitDatetimePreciseBasic is called when exiting the datetimePreciseBasic production.
	ExitDatetimePreciseBasic(c *DatetimePreciseBasicContext)

	// ExitDatetimePreciseExtended is called when exiting the datetimePreciseExtended production.
	ExitDatetimePreciseExtended(c *DatetimePreciseExtendedContext)

	// ExitDatetime is called when exiting the datetime production.
	ExitDatetime(c *DatetimeContext)

	// ExitDatetimeBasic is called when exiting the datetimeBasic production.
	ExitDatetimeBasic(c *DatetimeBasicContext)

	// ExitDatetimeExtended is called when exiting the datetimeExtended production.
	ExitDatetimeExtended(c *DatetimeExtendedContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitIntervalBasic is called when exiting the intervalBasic production.
	ExitIntervalBasic(c *IntervalBasicContext)

	// ExitIntervalExtended is called when exiting the intervalExtended production.
	ExitIntervalExtended(c *IntervalExtendedContext)

	// ExitIntervalT is called when exiting the intervalT production.
	ExitIntervalT(c *IntervalTContext)

	// ExitMonthDateBasic is called when exiting the monthDateBasic production.
	ExitMonthDateBasic(c *MonthDateBasicContext)

	// ExitMonthDateExtended is called when exiting the monthDateExtended production.
	ExitMonthDateExtended(c *MonthDateExtendedContext)

	// ExitIntervalYMDTimeBasic is called when exiting the intervalYMDTimeBasic production.
	ExitIntervalYMDTimeBasic(c *IntervalYMDTimeBasicContext)

	// ExitIntervalYMDTimeExtended is called when exiting the intervalYMDTimeExtended production.
	ExitIntervalYMDTimeExtended(c *IntervalYMDTimeExtendedContext)

	// ExitIntervalYMD is called when exiting the intervalYMD production.
	ExitIntervalYMD(c *IntervalYMDContext)

	// ExitIntervalYMDBasic is called when exiting the intervalYMDBasic production.
	ExitIntervalYMDBasic(c *IntervalYMDBasicContext)

	// ExitIntervalYMDExtended is called when exiting the intervalYMDExtended production.
	ExitIntervalYMDExtended(c *IntervalYMDExtendedContext)

	// ExitTimeBeginEnd is called when exiting the timeBeginEnd production.
	ExitTimeBeginEnd(c *TimeBeginEndContext)

	// ExitTimeBeginEndBasic is called when exiting the timeBeginEndBasic production.
	ExitTimeBeginEndBasic(c *TimeBeginEndBasicContext)

	// ExitTimeBeginEndExtended is called when exiting the timeBeginEndExtended production.
	ExitTimeBeginEndExtended(c *TimeBeginEndExtendedContext)

	// ExitTimeBeginInterval is called when exiting the timeBeginInterval production.
	ExitTimeBeginInterval(c *TimeBeginIntervalContext)

	// ExitTimeBeginIntervalBasic is called when exiting the timeBeginIntervalBasic production.
	ExitTimeBeginIntervalBasic(c *TimeBeginIntervalBasicContext)

	// ExitTimeBeginIntervalExtended is called when exiting the timeBeginIntervalExtended production.
	ExitTimeBeginIntervalExtended(c *TimeBeginIntervalExtendedContext)

	// ExitTimeIntervalEnd is called when exiting the timeIntervalEnd production.
	ExitTimeIntervalEnd(c *TimeIntervalEndContext)

	// ExitTimeIntervalEndBasic is called when exiting the timeIntervalEndBasic production.
	ExitTimeIntervalEndBasic(c *TimeIntervalEndBasicContext)

	// ExitTimeIntervalEndExtended is called when exiting the timeIntervalEndExtended production.
	ExitTimeIntervalEndExtended(c *TimeIntervalEndExtendedContext)

	// ExitDuration is called when exiting the duration production.
	ExitDuration(c *DurationContext)

	// ExitDurationBasic is called when exiting the durationBasic production.
	ExitDurationBasic(c *DurationBasicContext)

	// ExitDurationExtended is called when exiting the durationExtended production.
	ExitDurationExtended(c *DurationExtendedContext)

	// ExitRecurringCount is called when exiting the recurringCount production.
	ExitRecurringCount(c *RecurringCountContext)

	// ExitRecurring is called when exiting the recurring production.
	ExitRecurring(c *RecurringContext)

	// ExitRecurringBasic is called when exiting the recurringBasic production.
	ExitRecurringBasic(c *RecurringBasicContext)

	// ExitRecurringExtended is called when exiting the recurringExtended production.
	ExitRecurringExtended(c *RecurringExtendedContext)

	// ExitIso is called when exiting the iso production.
	ExitIso(c *IsoContext)
}
