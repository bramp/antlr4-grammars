// Code generated from Cobol85Preprocessor.g4 by ANTLR 4.7.2. DO NOT EDIT.

package cobol85preprocessor // Cobol85Preprocessor
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCobol85PreprocessorListener is a complete listener for a parse tree produced by Cobol85PreprocessorParser.
type BaseCobol85PreprocessorListener struct{}

var _ Cobol85PreprocessorListener = &BaseCobol85PreprocessorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCobol85PreprocessorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCobol85PreprocessorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCobol85PreprocessorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCobol85PreprocessorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStartRule is called when production startRule is entered.
func (s *BaseCobol85PreprocessorListener) EnterStartRule(ctx *StartRuleContext) {}

// ExitStartRule is called when production startRule is exited.
func (s *BaseCobol85PreprocessorListener) ExitStartRule(ctx *StartRuleContext) {}

// EnterCompilerOptions is called when production compilerOptions is entered.
func (s *BaseCobol85PreprocessorListener) EnterCompilerOptions(ctx *CompilerOptionsContext) {}

// ExitCompilerOptions is called when production compilerOptions is exited.
func (s *BaseCobol85PreprocessorListener) ExitCompilerOptions(ctx *CompilerOptionsContext) {}

// EnterCompilerXOpts is called when production compilerXOpts is entered.
func (s *BaseCobol85PreprocessorListener) EnterCompilerXOpts(ctx *CompilerXOptsContext) {}

// ExitCompilerXOpts is called when production compilerXOpts is exited.
func (s *BaseCobol85PreprocessorListener) ExitCompilerXOpts(ctx *CompilerXOptsContext) {}

// EnterCompilerOption is called when production compilerOption is entered.
func (s *BaseCobol85PreprocessorListener) EnterCompilerOption(ctx *CompilerOptionContext) {}

// ExitCompilerOption is called when production compilerOption is exited.
func (s *BaseCobol85PreprocessorListener) ExitCompilerOption(ctx *CompilerOptionContext) {}

// EnterExecCicsStatement is called when production execCicsStatement is entered.
func (s *BaseCobol85PreprocessorListener) EnterExecCicsStatement(ctx *ExecCicsStatementContext) {}

// ExitExecCicsStatement is called when production execCicsStatement is exited.
func (s *BaseCobol85PreprocessorListener) ExitExecCicsStatement(ctx *ExecCicsStatementContext) {}

// EnterExecSqlStatement is called when production execSqlStatement is entered.
func (s *BaseCobol85PreprocessorListener) EnterExecSqlStatement(ctx *ExecSqlStatementContext) {}

// ExitExecSqlStatement is called when production execSqlStatement is exited.
func (s *BaseCobol85PreprocessorListener) ExitExecSqlStatement(ctx *ExecSqlStatementContext) {}

// EnterExecSqlImsStatement is called when production execSqlImsStatement is entered.
func (s *BaseCobol85PreprocessorListener) EnterExecSqlImsStatement(ctx *ExecSqlImsStatementContext) {}

// ExitExecSqlImsStatement is called when production execSqlImsStatement is exited.
func (s *BaseCobol85PreprocessorListener) ExitExecSqlImsStatement(ctx *ExecSqlImsStatementContext) {}

// EnterCopyStatement is called when production copyStatement is entered.
func (s *BaseCobol85PreprocessorListener) EnterCopyStatement(ctx *CopyStatementContext) {}

// ExitCopyStatement is called when production copyStatement is exited.
func (s *BaseCobol85PreprocessorListener) ExitCopyStatement(ctx *CopyStatementContext) {}

// EnterCopySource is called when production copySource is entered.
func (s *BaseCobol85PreprocessorListener) EnterCopySource(ctx *CopySourceContext) {}

// ExitCopySource is called when production copySource is exited.
func (s *BaseCobol85PreprocessorListener) ExitCopySource(ctx *CopySourceContext) {}

// EnterCopyLibrary is called when production copyLibrary is entered.
func (s *BaseCobol85PreprocessorListener) EnterCopyLibrary(ctx *CopyLibraryContext) {}

// ExitCopyLibrary is called when production copyLibrary is exited.
func (s *BaseCobol85PreprocessorListener) ExitCopyLibrary(ctx *CopyLibraryContext) {}

// EnterReplacingPhrase is called when production replacingPhrase is entered.
func (s *BaseCobol85PreprocessorListener) EnterReplacingPhrase(ctx *ReplacingPhraseContext) {}

// ExitReplacingPhrase is called when production replacingPhrase is exited.
func (s *BaseCobol85PreprocessorListener) ExitReplacingPhrase(ctx *ReplacingPhraseContext) {}

// EnterReplaceArea is called when production replaceArea is entered.
func (s *BaseCobol85PreprocessorListener) EnterReplaceArea(ctx *ReplaceAreaContext) {}

// ExitReplaceArea is called when production replaceArea is exited.
func (s *BaseCobol85PreprocessorListener) ExitReplaceArea(ctx *ReplaceAreaContext) {}

// EnterReplaceByStatement is called when production replaceByStatement is entered.
func (s *BaseCobol85PreprocessorListener) EnterReplaceByStatement(ctx *ReplaceByStatementContext) {}

// ExitReplaceByStatement is called when production replaceByStatement is exited.
func (s *BaseCobol85PreprocessorListener) ExitReplaceByStatement(ctx *ReplaceByStatementContext) {}

// EnterReplaceOffStatement is called when production replaceOffStatement is entered.
func (s *BaseCobol85PreprocessorListener) EnterReplaceOffStatement(ctx *ReplaceOffStatementContext) {}

// ExitReplaceOffStatement is called when production replaceOffStatement is exited.
func (s *BaseCobol85PreprocessorListener) ExitReplaceOffStatement(ctx *ReplaceOffStatementContext) {}

// EnterReplaceClause is called when production replaceClause is entered.
func (s *BaseCobol85PreprocessorListener) EnterReplaceClause(ctx *ReplaceClauseContext) {}

// ExitReplaceClause is called when production replaceClause is exited.
func (s *BaseCobol85PreprocessorListener) ExitReplaceClause(ctx *ReplaceClauseContext) {}

// EnterDirectoryPhrase is called when production directoryPhrase is entered.
func (s *BaseCobol85PreprocessorListener) EnterDirectoryPhrase(ctx *DirectoryPhraseContext) {}

// ExitDirectoryPhrase is called when production directoryPhrase is exited.
func (s *BaseCobol85PreprocessorListener) ExitDirectoryPhrase(ctx *DirectoryPhraseContext) {}

// EnterFamilyPhrase is called when production familyPhrase is entered.
func (s *BaseCobol85PreprocessorListener) EnterFamilyPhrase(ctx *FamilyPhraseContext) {}

// ExitFamilyPhrase is called when production familyPhrase is exited.
func (s *BaseCobol85PreprocessorListener) ExitFamilyPhrase(ctx *FamilyPhraseContext) {}

// EnterReplaceable is called when production replaceable is entered.
func (s *BaseCobol85PreprocessorListener) EnterReplaceable(ctx *ReplaceableContext) {}

// ExitReplaceable is called when production replaceable is exited.
func (s *BaseCobol85PreprocessorListener) ExitReplaceable(ctx *ReplaceableContext) {}

// EnterReplacement is called when production replacement is entered.
func (s *BaseCobol85PreprocessorListener) EnterReplacement(ctx *ReplacementContext) {}

// ExitReplacement is called when production replacement is exited.
func (s *BaseCobol85PreprocessorListener) ExitReplacement(ctx *ReplacementContext) {}

// EnterEjectStatement is called when production ejectStatement is entered.
func (s *BaseCobol85PreprocessorListener) EnterEjectStatement(ctx *EjectStatementContext) {}

// ExitEjectStatement is called when production ejectStatement is exited.
func (s *BaseCobol85PreprocessorListener) ExitEjectStatement(ctx *EjectStatementContext) {}

// EnterSkipStatement is called when production skipStatement is entered.
func (s *BaseCobol85PreprocessorListener) EnterSkipStatement(ctx *SkipStatementContext) {}

// ExitSkipStatement is called when production skipStatement is exited.
func (s *BaseCobol85PreprocessorListener) ExitSkipStatement(ctx *SkipStatementContext) {}

// EnterTitleStatement is called when production titleStatement is entered.
func (s *BaseCobol85PreprocessorListener) EnterTitleStatement(ctx *TitleStatementContext) {}

// ExitTitleStatement is called when production titleStatement is exited.
func (s *BaseCobol85PreprocessorListener) ExitTitleStatement(ctx *TitleStatementContext) {}

// EnterPseudoText is called when production pseudoText is entered.
func (s *BaseCobol85PreprocessorListener) EnterPseudoText(ctx *PseudoTextContext) {}

// ExitPseudoText is called when production pseudoText is exited.
func (s *BaseCobol85PreprocessorListener) ExitPseudoText(ctx *PseudoTextContext) {}

// EnterCharData is called when production charData is entered.
func (s *BaseCobol85PreprocessorListener) EnterCharData(ctx *CharDataContext) {}

// ExitCharData is called when production charData is exited.
func (s *BaseCobol85PreprocessorListener) ExitCharData(ctx *CharDataContext) {}

// EnterCharDataSql is called when production charDataSql is entered.
func (s *BaseCobol85PreprocessorListener) EnterCharDataSql(ctx *CharDataSqlContext) {}

// ExitCharDataSql is called when production charDataSql is exited.
func (s *BaseCobol85PreprocessorListener) ExitCharDataSql(ctx *CharDataSqlContext) {}

// EnterCharDataLine is called when production charDataLine is entered.
func (s *BaseCobol85PreprocessorListener) EnterCharDataLine(ctx *CharDataLineContext) {}

// ExitCharDataLine is called when production charDataLine is exited.
func (s *BaseCobol85PreprocessorListener) ExitCharDataLine(ctx *CharDataLineContext) {}

// EnterCobolWord is called when production cobolWord is entered.
func (s *BaseCobol85PreprocessorListener) EnterCobolWord(ctx *CobolWordContext) {}

// ExitCobolWord is called when production cobolWord is exited.
func (s *BaseCobol85PreprocessorListener) ExitCobolWord(ctx *CobolWordContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCobol85PreprocessorListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCobol85PreprocessorListener) ExitLiteral(ctx *LiteralContext) {}

// EnterFilename is called when production filename is entered.
func (s *BaseCobol85PreprocessorListener) EnterFilename(ctx *FilenameContext) {}

// ExitFilename is called when production filename is exited.
func (s *BaseCobol85PreprocessorListener) ExitFilename(ctx *FilenameContext) {}

// EnterCharDataKeyword is called when production charDataKeyword is entered.
func (s *BaseCobol85PreprocessorListener) EnterCharDataKeyword(ctx *CharDataKeywordContext) {}

// ExitCharDataKeyword is called when production charDataKeyword is exited.
func (s *BaseCobol85PreprocessorListener) ExitCharDataKeyword(ctx *CharDataKeywordContext) {}
