// Generated from CSV.g4 by ANTLR 4.7.

package csv // CSV
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CSVParser.
type CSVVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CSVParser#csvFile.
	VisitCsvFile(ctx *CsvFileContext) interface{}

	// Visit a parse tree produced by CSVParser#hdr.
	VisitHdr(ctx *HdrContext) interface{}

	// Visit a parse tree produced by CSVParser#row.
	VisitRow(ctx *RowContext) interface{}

	// Visit a parse tree produced by CSVParser#field.
	VisitField(ctx *FieldContext) interface{}
}
