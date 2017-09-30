// Generated from CSV.g4 by ANTLR 4.7.

package csv // CSV
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCSVVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCSVVisitor) VisitCsvFile(ctx *CsvFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCSVVisitor) VisitHdr(ctx *HdrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCSVVisitor) VisitRow(ctx *RowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCSVVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}
