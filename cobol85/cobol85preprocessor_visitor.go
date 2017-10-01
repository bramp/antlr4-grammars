// Generated from Cobol85Preprocessor.g4 by ANTLR 4.7.

package cobol85 // Cobol85Preprocessor
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by Cobol85PreprocessorParser.
type Cobol85PreprocessorVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Cobol85PreprocessorParser#startRule.
	VisitStartRule(ctx *StartRuleContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#execCicsStatement.
	VisitExecCicsStatement(ctx *ExecCicsStatementContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#execSqlStatement.
	VisitExecSqlStatement(ctx *ExecSqlStatementContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#execSqlImsStatement.
	VisitExecSqlImsStatement(ctx *ExecSqlImsStatementContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#copyStatement.
	VisitCopyStatement(ctx *CopyStatementContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#copySource.
	VisitCopySource(ctx *CopySourceContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#replacingPhrase.
	VisitReplacingPhrase(ctx *ReplacingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#replaceArea.
	VisitReplaceArea(ctx *ReplaceAreaContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#replaceByStatement.
	VisitReplaceByStatement(ctx *ReplaceByStatementContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#replaceOffStatement.
	VisitReplaceOffStatement(ctx *ReplaceOffStatementContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#replaceClause.
	VisitReplaceClause(ctx *ReplaceClauseContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#directoryPhrase.
	VisitDirectoryPhrase(ctx *DirectoryPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#familyPhrase.
	VisitFamilyPhrase(ctx *FamilyPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#replaceable.
	VisitReplaceable(ctx *ReplaceableContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#replacement.
	VisitReplacement(ctx *ReplacementContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#controlSpacingStatement.
	VisitControlSpacingStatement(ctx *ControlSpacingStatementContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#cobolWord.
	VisitCobolWord(ctx *CobolWordContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#pseudoText.
	VisitPseudoText(ctx *PseudoTextContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#charData.
	VisitCharData(ctx *CharDataContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#charDataLine.
	VisitCharDataLine(ctx *CharDataLineContext) interface{}

	// Visit a parse tree produced by Cobol85PreprocessorParser#charDataKeyword.
	VisitCharDataKeyword(ctx *CharDataKeywordContext) interface{}
}
