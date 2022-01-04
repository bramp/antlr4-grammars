// Code generated from Cobol85Preprocessor.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cobol85preprocessor // Cobol85Preprocessor
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Cobol85PreprocessorListener is a complete listener for a parse tree produced by Cobol85PreprocessorParser.
type Cobol85PreprocessorListener interface {
	antlr.ParseTreeListener

	// EnterStartRule is called when entering the startRule production.
	EnterStartRule(c *StartRuleContext)

	// EnterCompilerOptions is called when entering the compilerOptions production.
	EnterCompilerOptions(c *CompilerOptionsContext)

	// EnterCompilerXOpts is called when entering the compilerXOpts production.
	EnterCompilerXOpts(c *CompilerXOptsContext)

	// EnterCompilerOption is called when entering the compilerOption production.
	EnterCompilerOption(c *CompilerOptionContext)

	// EnterExecCicsStatement is called when entering the execCicsStatement production.
	EnterExecCicsStatement(c *ExecCicsStatementContext)

	// EnterExecSqlStatement is called when entering the execSqlStatement production.
	EnterExecSqlStatement(c *ExecSqlStatementContext)

	// EnterExecSqlImsStatement is called when entering the execSqlImsStatement production.
	EnterExecSqlImsStatement(c *ExecSqlImsStatementContext)

	// EnterCopyStatement is called when entering the copyStatement production.
	EnterCopyStatement(c *CopyStatementContext)

	// EnterCopySource is called when entering the copySource production.
	EnterCopySource(c *CopySourceContext)

	// EnterCopyLibrary is called when entering the copyLibrary production.
	EnterCopyLibrary(c *CopyLibraryContext)

	// EnterReplacingPhrase is called when entering the replacingPhrase production.
	EnterReplacingPhrase(c *ReplacingPhraseContext)

	// EnterReplaceArea is called when entering the replaceArea production.
	EnterReplaceArea(c *ReplaceAreaContext)

	// EnterReplaceByStatement is called when entering the replaceByStatement production.
	EnterReplaceByStatement(c *ReplaceByStatementContext)

	// EnterReplaceOffStatement is called when entering the replaceOffStatement production.
	EnterReplaceOffStatement(c *ReplaceOffStatementContext)

	// EnterReplaceClause is called when entering the replaceClause production.
	EnterReplaceClause(c *ReplaceClauseContext)

	// EnterDirectoryPhrase is called when entering the directoryPhrase production.
	EnterDirectoryPhrase(c *DirectoryPhraseContext)

	// EnterFamilyPhrase is called when entering the familyPhrase production.
	EnterFamilyPhrase(c *FamilyPhraseContext)

	// EnterReplaceable is called when entering the replaceable production.
	EnterReplaceable(c *ReplaceableContext)

	// EnterReplacement is called when entering the replacement production.
	EnterReplacement(c *ReplacementContext)

	// EnterEjectStatement is called when entering the ejectStatement production.
	EnterEjectStatement(c *EjectStatementContext)

	// EnterSkipStatement is called when entering the skipStatement production.
	EnterSkipStatement(c *SkipStatementContext)

	// EnterTitleStatement is called when entering the titleStatement production.
	EnterTitleStatement(c *TitleStatementContext)

	// EnterPseudoText is called when entering the pseudoText production.
	EnterPseudoText(c *PseudoTextContext)

	// EnterCharData is called when entering the charData production.
	EnterCharData(c *CharDataContext)

	// EnterCharDataSql is called when entering the charDataSql production.
	EnterCharDataSql(c *CharDataSqlContext)

	// EnterCharDataLine is called when entering the charDataLine production.
	EnterCharDataLine(c *CharDataLineContext)

	// EnterCobolWord is called when entering the cobolWord production.
	EnterCobolWord(c *CobolWordContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterFilename is called when entering the filename production.
	EnterFilename(c *FilenameContext)

	// EnterCharDataKeyword is called when entering the charDataKeyword production.
	EnterCharDataKeyword(c *CharDataKeywordContext)

	// ExitStartRule is called when exiting the startRule production.
	ExitStartRule(c *StartRuleContext)

	// ExitCompilerOptions is called when exiting the compilerOptions production.
	ExitCompilerOptions(c *CompilerOptionsContext)

	// ExitCompilerXOpts is called when exiting the compilerXOpts production.
	ExitCompilerXOpts(c *CompilerXOptsContext)

	// ExitCompilerOption is called when exiting the compilerOption production.
	ExitCompilerOption(c *CompilerOptionContext)

	// ExitExecCicsStatement is called when exiting the execCicsStatement production.
	ExitExecCicsStatement(c *ExecCicsStatementContext)

	// ExitExecSqlStatement is called when exiting the execSqlStatement production.
	ExitExecSqlStatement(c *ExecSqlStatementContext)

	// ExitExecSqlImsStatement is called when exiting the execSqlImsStatement production.
	ExitExecSqlImsStatement(c *ExecSqlImsStatementContext)

	// ExitCopyStatement is called when exiting the copyStatement production.
	ExitCopyStatement(c *CopyStatementContext)

	// ExitCopySource is called when exiting the copySource production.
	ExitCopySource(c *CopySourceContext)

	// ExitCopyLibrary is called when exiting the copyLibrary production.
	ExitCopyLibrary(c *CopyLibraryContext)

	// ExitReplacingPhrase is called when exiting the replacingPhrase production.
	ExitReplacingPhrase(c *ReplacingPhraseContext)

	// ExitReplaceArea is called when exiting the replaceArea production.
	ExitReplaceArea(c *ReplaceAreaContext)

	// ExitReplaceByStatement is called when exiting the replaceByStatement production.
	ExitReplaceByStatement(c *ReplaceByStatementContext)

	// ExitReplaceOffStatement is called when exiting the replaceOffStatement production.
	ExitReplaceOffStatement(c *ReplaceOffStatementContext)

	// ExitReplaceClause is called when exiting the replaceClause production.
	ExitReplaceClause(c *ReplaceClauseContext)

	// ExitDirectoryPhrase is called when exiting the directoryPhrase production.
	ExitDirectoryPhrase(c *DirectoryPhraseContext)

	// ExitFamilyPhrase is called when exiting the familyPhrase production.
	ExitFamilyPhrase(c *FamilyPhraseContext)

	// ExitReplaceable is called when exiting the replaceable production.
	ExitReplaceable(c *ReplaceableContext)

	// ExitReplacement is called when exiting the replacement production.
	ExitReplacement(c *ReplacementContext)

	// ExitEjectStatement is called when exiting the ejectStatement production.
	ExitEjectStatement(c *EjectStatementContext)

	// ExitSkipStatement is called when exiting the skipStatement production.
	ExitSkipStatement(c *SkipStatementContext)

	// ExitTitleStatement is called when exiting the titleStatement production.
	ExitTitleStatement(c *TitleStatementContext)

	// ExitPseudoText is called when exiting the pseudoText production.
	ExitPseudoText(c *PseudoTextContext)

	// ExitCharData is called when exiting the charData production.
	ExitCharData(c *CharDataContext)

	// ExitCharDataSql is called when exiting the charDataSql production.
	ExitCharDataSql(c *CharDataSqlContext)

	// ExitCharDataLine is called when exiting the charDataLine production.
	ExitCharDataLine(c *CharDataLineContext)

	// ExitCobolWord is called when exiting the cobolWord production.
	ExitCobolWord(c *CobolWordContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitFilename is called when exiting the filename production.
	ExitFilename(c *FilenameContext)

	// ExitCharDataKeyword is called when exiting the charDataKeyword production.
	ExitCharDataKeyword(c *CharDataKeywordContext)
}
