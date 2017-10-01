// Generated from Cobol85Preprocessor.g4 by ANTLR 4.7.

package cobol85 // Cobol85Preprocessor
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCobol85PreprocessorVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCobol85PreprocessorVisitor) VisitStartRule(ctx *StartRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitExecCicsStatement(ctx *ExecCicsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitExecSqlStatement(ctx *ExecSqlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitExecSqlImsStatement(ctx *ExecSqlImsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitCopyStatement(ctx *CopyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitCopySource(ctx *CopySourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitReplacingPhrase(ctx *ReplacingPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitReplaceArea(ctx *ReplaceAreaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitReplaceByStatement(ctx *ReplaceByStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitReplaceOffStatement(ctx *ReplaceOffStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitReplaceClause(ctx *ReplaceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitDirectoryPhrase(ctx *DirectoryPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitFamilyPhrase(ctx *FamilyPhraseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitReplaceable(ctx *ReplaceableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitReplacement(ctx *ReplacementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitControlSpacingStatement(ctx *ControlSpacingStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitCobolWord(ctx *CobolWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitPseudoText(ctx *PseudoTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitCharData(ctx *CharDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitCharDataLine(ctx *CharDataLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCobol85PreprocessorVisitor) VisitCharDataKeyword(ctx *CharDataKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}
