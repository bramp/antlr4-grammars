// Generated from datetime.g4 by ANTLR 4.7.

package datetime // datetime
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasedatetimeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasedatetimeVisitor) VisitDate_time(ctx *Date_timeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedatetimeVisitor) VisitDay(ctx *DayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedatetimeVisitor) VisitDate(ctx *DateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedatetimeVisitor) VisitMonth(ctx *MonthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedatetimeVisitor) VisitTime(ctx *TimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedatetimeVisitor) VisitHour(ctx *HourContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedatetimeVisitor) VisitZone(ctx *ZoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedatetimeVisitor) VisitTwo_digit(ctx *Two_digitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedatetimeVisitor) VisitFour_digit(ctx *Four_digitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedatetimeVisitor) VisitAlphanumeric(ctx *AlphanumericContext) interface{} {
	return v.VisitChildren(ctx)
}
