// Code generated from Cobol85.g4 by ANTLR 4.7.2. DO NOT EDIT.

package cobol85 // Cobol85
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCobol85Listener is a complete listener for a parse tree produced by Cobol85Parser.
type BaseCobol85Listener struct{}

var _ Cobol85Listener = &BaseCobol85Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCobol85Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCobol85Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCobol85Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCobol85Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStartRule is called when production startRule is entered.
func (s *BaseCobol85Listener) EnterStartRule(ctx *StartRuleContext) {}

// ExitStartRule is called when production startRule is exited.
func (s *BaseCobol85Listener) ExitStartRule(ctx *StartRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseCobol85Listener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseCobol85Listener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterProgramUnit is called when production programUnit is entered.
func (s *BaseCobol85Listener) EnterProgramUnit(ctx *ProgramUnitContext) {}

// ExitProgramUnit is called when production programUnit is exited.
func (s *BaseCobol85Listener) ExitProgramUnit(ctx *ProgramUnitContext) {}

// EnterEndProgramStatement is called when production endProgramStatement is entered.
func (s *BaseCobol85Listener) EnterEndProgramStatement(ctx *EndProgramStatementContext) {}

// ExitEndProgramStatement is called when production endProgramStatement is exited.
func (s *BaseCobol85Listener) ExitEndProgramStatement(ctx *EndProgramStatementContext) {}

// EnterIdentificationDivision is called when production identificationDivision is entered.
func (s *BaseCobol85Listener) EnterIdentificationDivision(ctx *IdentificationDivisionContext) {}

// ExitIdentificationDivision is called when production identificationDivision is exited.
func (s *BaseCobol85Listener) ExitIdentificationDivision(ctx *IdentificationDivisionContext) {}

// EnterIdentificationDivisionBody is called when production identificationDivisionBody is entered.
func (s *BaseCobol85Listener) EnterIdentificationDivisionBody(ctx *IdentificationDivisionBodyContext) {
}

// ExitIdentificationDivisionBody is called when production identificationDivisionBody is exited.
func (s *BaseCobol85Listener) ExitIdentificationDivisionBody(ctx *IdentificationDivisionBodyContext) {}

// EnterProgramIdParagraph is called when production programIdParagraph is entered.
func (s *BaseCobol85Listener) EnterProgramIdParagraph(ctx *ProgramIdParagraphContext) {}

// ExitProgramIdParagraph is called when production programIdParagraph is exited.
func (s *BaseCobol85Listener) ExitProgramIdParagraph(ctx *ProgramIdParagraphContext) {}

// EnterAuthorParagraph is called when production authorParagraph is entered.
func (s *BaseCobol85Listener) EnterAuthorParagraph(ctx *AuthorParagraphContext) {}

// ExitAuthorParagraph is called when production authorParagraph is exited.
func (s *BaseCobol85Listener) ExitAuthorParagraph(ctx *AuthorParagraphContext) {}

// EnterInstallationParagraph is called when production installationParagraph is entered.
func (s *BaseCobol85Listener) EnterInstallationParagraph(ctx *InstallationParagraphContext) {}

// ExitInstallationParagraph is called when production installationParagraph is exited.
func (s *BaseCobol85Listener) ExitInstallationParagraph(ctx *InstallationParagraphContext) {}

// EnterDateWrittenParagraph is called when production dateWrittenParagraph is entered.
func (s *BaseCobol85Listener) EnterDateWrittenParagraph(ctx *DateWrittenParagraphContext) {}

// ExitDateWrittenParagraph is called when production dateWrittenParagraph is exited.
func (s *BaseCobol85Listener) ExitDateWrittenParagraph(ctx *DateWrittenParagraphContext) {}

// EnterDateCompiledParagraph is called when production dateCompiledParagraph is entered.
func (s *BaseCobol85Listener) EnterDateCompiledParagraph(ctx *DateCompiledParagraphContext) {}

// ExitDateCompiledParagraph is called when production dateCompiledParagraph is exited.
func (s *BaseCobol85Listener) ExitDateCompiledParagraph(ctx *DateCompiledParagraphContext) {}

// EnterSecurityParagraph is called when production securityParagraph is entered.
func (s *BaseCobol85Listener) EnterSecurityParagraph(ctx *SecurityParagraphContext) {}

// ExitSecurityParagraph is called when production securityParagraph is exited.
func (s *BaseCobol85Listener) ExitSecurityParagraph(ctx *SecurityParagraphContext) {}

// EnterRemarksParagraph is called when production remarksParagraph is entered.
func (s *BaseCobol85Listener) EnterRemarksParagraph(ctx *RemarksParagraphContext) {}

// ExitRemarksParagraph is called when production remarksParagraph is exited.
func (s *BaseCobol85Listener) ExitRemarksParagraph(ctx *RemarksParagraphContext) {}

// EnterEnvironmentDivision is called when production environmentDivision is entered.
func (s *BaseCobol85Listener) EnterEnvironmentDivision(ctx *EnvironmentDivisionContext) {}

// ExitEnvironmentDivision is called when production environmentDivision is exited.
func (s *BaseCobol85Listener) ExitEnvironmentDivision(ctx *EnvironmentDivisionContext) {}

// EnterEnvironmentDivisionBody is called when production environmentDivisionBody is entered.
func (s *BaseCobol85Listener) EnterEnvironmentDivisionBody(ctx *EnvironmentDivisionBodyContext) {}

// ExitEnvironmentDivisionBody is called when production environmentDivisionBody is exited.
func (s *BaseCobol85Listener) ExitEnvironmentDivisionBody(ctx *EnvironmentDivisionBodyContext) {}

// EnterConfigurationSection is called when production configurationSection is entered.
func (s *BaseCobol85Listener) EnterConfigurationSection(ctx *ConfigurationSectionContext) {}

// ExitConfigurationSection is called when production configurationSection is exited.
func (s *BaseCobol85Listener) ExitConfigurationSection(ctx *ConfigurationSectionContext) {}

// EnterConfigurationSectionParagraph is called when production configurationSectionParagraph is entered.
func (s *BaseCobol85Listener) EnterConfigurationSectionParagraph(ctx *ConfigurationSectionParagraphContext) {
}

// ExitConfigurationSectionParagraph is called when production configurationSectionParagraph is exited.
func (s *BaseCobol85Listener) ExitConfigurationSectionParagraph(ctx *ConfigurationSectionParagraphContext) {
}

// EnterSourceComputerParagraph is called when production sourceComputerParagraph is entered.
func (s *BaseCobol85Listener) EnterSourceComputerParagraph(ctx *SourceComputerParagraphContext) {}

// ExitSourceComputerParagraph is called when production sourceComputerParagraph is exited.
func (s *BaseCobol85Listener) ExitSourceComputerParagraph(ctx *SourceComputerParagraphContext) {}

// EnterObjectComputerParagraph is called when production objectComputerParagraph is entered.
func (s *BaseCobol85Listener) EnterObjectComputerParagraph(ctx *ObjectComputerParagraphContext) {}

// ExitObjectComputerParagraph is called when production objectComputerParagraph is exited.
func (s *BaseCobol85Listener) ExitObjectComputerParagraph(ctx *ObjectComputerParagraphContext) {}

// EnterObjectComputerClause is called when production objectComputerClause is entered.
func (s *BaseCobol85Listener) EnterObjectComputerClause(ctx *ObjectComputerClauseContext) {}

// ExitObjectComputerClause is called when production objectComputerClause is exited.
func (s *BaseCobol85Listener) ExitObjectComputerClause(ctx *ObjectComputerClauseContext) {}

// EnterMemorySizeClause is called when production memorySizeClause is entered.
func (s *BaseCobol85Listener) EnterMemorySizeClause(ctx *MemorySizeClauseContext) {}

// ExitMemorySizeClause is called when production memorySizeClause is exited.
func (s *BaseCobol85Listener) ExitMemorySizeClause(ctx *MemorySizeClauseContext) {}

// EnterDiskSizeClause is called when production diskSizeClause is entered.
func (s *BaseCobol85Listener) EnterDiskSizeClause(ctx *DiskSizeClauseContext) {}

// ExitDiskSizeClause is called when production diskSizeClause is exited.
func (s *BaseCobol85Listener) ExitDiskSizeClause(ctx *DiskSizeClauseContext) {}

// EnterCollatingSequenceClause is called when production collatingSequenceClause is entered.
func (s *BaseCobol85Listener) EnterCollatingSequenceClause(ctx *CollatingSequenceClauseContext) {}

// ExitCollatingSequenceClause is called when production collatingSequenceClause is exited.
func (s *BaseCobol85Listener) ExitCollatingSequenceClause(ctx *CollatingSequenceClauseContext) {}

// EnterCollatingSequenceClauseAlphanumeric is called when production collatingSequenceClauseAlphanumeric is entered.
func (s *BaseCobol85Listener) EnterCollatingSequenceClauseAlphanumeric(ctx *CollatingSequenceClauseAlphanumericContext) {
}

// ExitCollatingSequenceClauseAlphanumeric is called when production collatingSequenceClauseAlphanumeric is exited.
func (s *BaseCobol85Listener) ExitCollatingSequenceClauseAlphanumeric(ctx *CollatingSequenceClauseAlphanumericContext) {
}

// EnterCollatingSequenceClauseNational is called when production collatingSequenceClauseNational is entered.
func (s *BaseCobol85Listener) EnterCollatingSequenceClauseNational(ctx *CollatingSequenceClauseNationalContext) {
}

// ExitCollatingSequenceClauseNational is called when production collatingSequenceClauseNational is exited.
func (s *BaseCobol85Listener) ExitCollatingSequenceClauseNational(ctx *CollatingSequenceClauseNationalContext) {
}

// EnterSegmentLimitClause is called when production segmentLimitClause is entered.
func (s *BaseCobol85Listener) EnterSegmentLimitClause(ctx *SegmentLimitClauseContext) {}

// ExitSegmentLimitClause is called when production segmentLimitClause is exited.
func (s *BaseCobol85Listener) ExitSegmentLimitClause(ctx *SegmentLimitClauseContext) {}

// EnterCharacterSetClause is called when production characterSetClause is entered.
func (s *BaseCobol85Listener) EnterCharacterSetClause(ctx *CharacterSetClauseContext) {}

// ExitCharacterSetClause is called when production characterSetClause is exited.
func (s *BaseCobol85Listener) ExitCharacterSetClause(ctx *CharacterSetClauseContext) {}

// EnterSpecialNamesParagraph is called when production specialNamesParagraph is entered.
func (s *BaseCobol85Listener) EnterSpecialNamesParagraph(ctx *SpecialNamesParagraphContext) {}

// ExitSpecialNamesParagraph is called when production specialNamesParagraph is exited.
func (s *BaseCobol85Listener) ExitSpecialNamesParagraph(ctx *SpecialNamesParagraphContext) {}

// EnterSpecialNameClause is called when production specialNameClause is entered.
func (s *BaseCobol85Listener) EnterSpecialNameClause(ctx *SpecialNameClauseContext) {}

// ExitSpecialNameClause is called when production specialNameClause is exited.
func (s *BaseCobol85Listener) ExitSpecialNameClause(ctx *SpecialNameClauseContext) {}

// EnterAlphabetClause is called when production alphabetClause is entered.
func (s *BaseCobol85Listener) EnterAlphabetClause(ctx *AlphabetClauseContext) {}

// ExitAlphabetClause is called when production alphabetClause is exited.
func (s *BaseCobol85Listener) ExitAlphabetClause(ctx *AlphabetClauseContext) {}

// EnterAlphabetClauseFormat1 is called when production alphabetClauseFormat1 is entered.
func (s *BaseCobol85Listener) EnterAlphabetClauseFormat1(ctx *AlphabetClauseFormat1Context) {}

// ExitAlphabetClauseFormat1 is called when production alphabetClauseFormat1 is exited.
func (s *BaseCobol85Listener) ExitAlphabetClauseFormat1(ctx *AlphabetClauseFormat1Context) {}

// EnterAlphabetLiterals is called when production alphabetLiterals is entered.
func (s *BaseCobol85Listener) EnterAlphabetLiterals(ctx *AlphabetLiteralsContext) {}

// ExitAlphabetLiterals is called when production alphabetLiterals is exited.
func (s *BaseCobol85Listener) ExitAlphabetLiterals(ctx *AlphabetLiteralsContext) {}

// EnterAlphabetThrough is called when production alphabetThrough is entered.
func (s *BaseCobol85Listener) EnterAlphabetThrough(ctx *AlphabetThroughContext) {}

// ExitAlphabetThrough is called when production alphabetThrough is exited.
func (s *BaseCobol85Listener) ExitAlphabetThrough(ctx *AlphabetThroughContext) {}

// EnterAlphabetAlso is called when production alphabetAlso is entered.
func (s *BaseCobol85Listener) EnterAlphabetAlso(ctx *AlphabetAlsoContext) {}

// ExitAlphabetAlso is called when production alphabetAlso is exited.
func (s *BaseCobol85Listener) ExitAlphabetAlso(ctx *AlphabetAlsoContext) {}

// EnterAlphabetClauseFormat2 is called when production alphabetClauseFormat2 is entered.
func (s *BaseCobol85Listener) EnterAlphabetClauseFormat2(ctx *AlphabetClauseFormat2Context) {}

// ExitAlphabetClauseFormat2 is called when production alphabetClauseFormat2 is exited.
func (s *BaseCobol85Listener) ExitAlphabetClauseFormat2(ctx *AlphabetClauseFormat2Context) {}

// EnterChannelClause is called when production channelClause is entered.
func (s *BaseCobol85Listener) EnterChannelClause(ctx *ChannelClauseContext) {}

// ExitChannelClause is called when production channelClause is exited.
func (s *BaseCobol85Listener) ExitChannelClause(ctx *ChannelClauseContext) {}

// EnterClassClause is called when production classClause is entered.
func (s *BaseCobol85Listener) EnterClassClause(ctx *ClassClauseContext) {}

// ExitClassClause is called when production classClause is exited.
func (s *BaseCobol85Listener) ExitClassClause(ctx *ClassClauseContext) {}

// EnterClassClauseThrough is called when production classClauseThrough is entered.
func (s *BaseCobol85Listener) EnterClassClauseThrough(ctx *ClassClauseThroughContext) {}

// ExitClassClauseThrough is called when production classClauseThrough is exited.
func (s *BaseCobol85Listener) ExitClassClauseThrough(ctx *ClassClauseThroughContext) {}

// EnterClassClauseFrom is called when production classClauseFrom is entered.
func (s *BaseCobol85Listener) EnterClassClauseFrom(ctx *ClassClauseFromContext) {}

// ExitClassClauseFrom is called when production classClauseFrom is exited.
func (s *BaseCobol85Listener) ExitClassClauseFrom(ctx *ClassClauseFromContext) {}

// EnterClassClauseTo is called when production classClauseTo is entered.
func (s *BaseCobol85Listener) EnterClassClauseTo(ctx *ClassClauseToContext) {}

// ExitClassClauseTo is called when production classClauseTo is exited.
func (s *BaseCobol85Listener) ExitClassClauseTo(ctx *ClassClauseToContext) {}

// EnterCurrencySignClause is called when production currencySignClause is entered.
func (s *BaseCobol85Listener) EnterCurrencySignClause(ctx *CurrencySignClauseContext) {}

// ExitCurrencySignClause is called when production currencySignClause is exited.
func (s *BaseCobol85Listener) ExitCurrencySignClause(ctx *CurrencySignClauseContext) {}

// EnterDecimalPointClause is called when production decimalPointClause is entered.
func (s *BaseCobol85Listener) EnterDecimalPointClause(ctx *DecimalPointClauseContext) {}

// ExitDecimalPointClause is called when production decimalPointClause is exited.
func (s *BaseCobol85Listener) ExitDecimalPointClause(ctx *DecimalPointClauseContext) {}

// EnterDefaultComputationalSignClause is called when production defaultComputationalSignClause is entered.
func (s *BaseCobol85Listener) EnterDefaultComputationalSignClause(ctx *DefaultComputationalSignClauseContext) {
}

// ExitDefaultComputationalSignClause is called when production defaultComputationalSignClause is exited.
func (s *BaseCobol85Listener) ExitDefaultComputationalSignClause(ctx *DefaultComputationalSignClauseContext) {
}

// EnterDefaultDisplaySignClause is called when production defaultDisplaySignClause is entered.
func (s *BaseCobol85Listener) EnterDefaultDisplaySignClause(ctx *DefaultDisplaySignClauseContext) {}

// ExitDefaultDisplaySignClause is called when production defaultDisplaySignClause is exited.
func (s *BaseCobol85Listener) ExitDefaultDisplaySignClause(ctx *DefaultDisplaySignClauseContext) {}

// EnterEnvironmentSwitchNameClause is called when production environmentSwitchNameClause is entered.
func (s *BaseCobol85Listener) EnterEnvironmentSwitchNameClause(ctx *EnvironmentSwitchNameClauseContext) {
}

// ExitEnvironmentSwitchNameClause is called when production environmentSwitchNameClause is exited.
func (s *BaseCobol85Listener) ExitEnvironmentSwitchNameClause(ctx *EnvironmentSwitchNameClauseContext) {
}

// EnterEnvironmentSwitchNameSpecialNamesStatusPhrase is called when production environmentSwitchNameSpecialNamesStatusPhrase is entered.
func (s *BaseCobol85Listener) EnterEnvironmentSwitchNameSpecialNamesStatusPhrase(ctx *EnvironmentSwitchNameSpecialNamesStatusPhraseContext) {
}

// ExitEnvironmentSwitchNameSpecialNamesStatusPhrase is called when production environmentSwitchNameSpecialNamesStatusPhrase is exited.
func (s *BaseCobol85Listener) ExitEnvironmentSwitchNameSpecialNamesStatusPhrase(ctx *EnvironmentSwitchNameSpecialNamesStatusPhraseContext) {
}

// EnterOdtClause is called when production odtClause is entered.
func (s *BaseCobol85Listener) EnterOdtClause(ctx *OdtClauseContext) {}

// ExitOdtClause is called when production odtClause is exited.
func (s *BaseCobol85Listener) ExitOdtClause(ctx *OdtClauseContext) {}

// EnterReserveNetworkClause is called when production reserveNetworkClause is entered.
func (s *BaseCobol85Listener) EnterReserveNetworkClause(ctx *ReserveNetworkClauseContext) {}

// ExitReserveNetworkClause is called when production reserveNetworkClause is exited.
func (s *BaseCobol85Listener) ExitReserveNetworkClause(ctx *ReserveNetworkClauseContext) {}

// EnterSymbolicCharactersClause is called when production symbolicCharactersClause is entered.
func (s *BaseCobol85Listener) EnterSymbolicCharactersClause(ctx *SymbolicCharactersClauseContext) {}

// ExitSymbolicCharactersClause is called when production symbolicCharactersClause is exited.
func (s *BaseCobol85Listener) ExitSymbolicCharactersClause(ctx *SymbolicCharactersClauseContext) {}

// EnterSymbolicCharacters is called when production symbolicCharacters is entered.
func (s *BaseCobol85Listener) EnterSymbolicCharacters(ctx *SymbolicCharactersContext) {}

// ExitSymbolicCharacters is called when production symbolicCharacters is exited.
func (s *BaseCobol85Listener) ExitSymbolicCharacters(ctx *SymbolicCharactersContext) {}

// EnterInputOutputSection is called when production inputOutputSection is entered.
func (s *BaseCobol85Listener) EnterInputOutputSection(ctx *InputOutputSectionContext) {}

// ExitInputOutputSection is called when production inputOutputSection is exited.
func (s *BaseCobol85Listener) ExitInputOutputSection(ctx *InputOutputSectionContext) {}

// EnterInputOutputSectionParagraph is called when production inputOutputSectionParagraph is entered.
func (s *BaseCobol85Listener) EnterInputOutputSectionParagraph(ctx *InputOutputSectionParagraphContext) {
}

// ExitInputOutputSectionParagraph is called when production inputOutputSectionParagraph is exited.
func (s *BaseCobol85Listener) ExitInputOutputSectionParagraph(ctx *InputOutputSectionParagraphContext) {
}

// EnterFileControlParagraph is called when production fileControlParagraph is entered.
func (s *BaseCobol85Listener) EnterFileControlParagraph(ctx *FileControlParagraphContext) {}

// ExitFileControlParagraph is called when production fileControlParagraph is exited.
func (s *BaseCobol85Listener) ExitFileControlParagraph(ctx *FileControlParagraphContext) {}

// EnterFileControlEntry is called when production fileControlEntry is entered.
func (s *BaseCobol85Listener) EnterFileControlEntry(ctx *FileControlEntryContext) {}

// ExitFileControlEntry is called when production fileControlEntry is exited.
func (s *BaseCobol85Listener) ExitFileControlEntry(ctx *FileControlEntryContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseCobol85Listener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseCobol85Listener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterFileControlClause is called when production fileControlClause is entered.
func (s *BaseCobol85Listener) EnterFileControlClause(ctx *FileControlClauseContext) {}

// ExitFileControlClause is called when production fileControlClause is exited.
func (s *BaseCobol85Listener) ExitFileControlClause(ctx *FileControlClauseContext) {}

// EnterAssignClause is called when production assignClause is entered.
func (s *BaseCobol85Listener) EnterAssignClause(ctx *AssignClauseContext) {}

// ExitAssignClause is called when production assignClause is exited.
func (s *BaseCobol85Listener) ExitAssignClause(ctx *AssignClauseContext) {}

// EnterReserveClause is called when production reserveClause is entered.
func (s *BaseCobol85Listener) EnterReserveClause(ctx *ReserveClauseContext) {}

// ExitReserveClause is called when production reserveClause is exited.
func (s *BaseCobol85Listener) ExitReserveClause(ctx *ReserveClauseContext) {}

// EnterOrganizationClause is called when production organizationClause is entered.
func (s *BaseCobol85Listener) EnterOrganizationClause(ctx *OrganizationClauseContext) {}

// ExitOrganizationClause is called when production organizationClause is exited.
func (s *BaseCobol85Listener) ExitOrganizationClause(ctx *OrganizationClauseContext) {}

// EnterPaddingCharacterClause is called when production paddingCharacterClause is entered.
func (s *BaseCobol85Listener) EnterPaddingCharacterClause(ctx *PaddingCharacterClauseContext) {}

// ExitPaddingCharacterClause is called when production paddingCharacterClause is exited.
func (s *BaseCobol85Listener) ExitPaddingCharacterClause(ctx *PaddingCharacterClauseContext) {}

// EnterRecordDelimiterClause is called when production recordDelimiterClause is entered.
func (s *BaseCobol85Listener) EnterRecordDelimiterClause(ctx *RecordDelimiterClauseContext) {}

// ExitRecordDelimiterClause is called when production recordDelimiterClause is exited.
func (s *BaseCobol85Listener) ExitRecordDelimiterClause(ctx *RecordDelimiterClauseContext) {}

// EnterAccessModeClause is called when production accessModeClause is entered.
func (s *BaseCobol85Listener) EnterAccessModeClause(ctx *AccessModeClauseContext) {}

// ExitAccessModeClause is called when production accessModeClause is exited.
func (s *BaseCobol85Listener) ExitAccessModeClause(ctx *AccessModeClauseContext) {}

// EnterRecordKeyClause is called when production recordKeyClause is entered.
func (s *BaseCobol85Listener) EnterRecordKeyClause(ctx *RecordKeyClauseContext) {}

// ExitRecordKeyClause is called when production recordKeyClause is exited.
func (s *BaseCobol85Listener) ExitRecordKeyClause(ctx *RecordKeyClauseContext) {}

// EnterAlternateRecordKeyClause is called when production alternateRecordKeyClause is entered.
func (s *BaseCobol85Listener) EnterAlternateRecordKeyClause(ctx *AlternateRecordKeyClauseContext) {}

// ExitAlternateRecordKeyClause is called when production alternateRecordKeyClause is exited.
func (s *BaseCobol85Listener) ExitAlternateRecordKeyClause(ctx *AlternateRecordKeyClauseContext) {}

// EnterPasswordClause is called when production passwordClause is entered.
func (s *BaseCobol85Listener) EnterPasswordClause(ctx *PasswordClauseContext) {}

// ExitPasswordClause is called when production passwordClause is exited.
func (s *BaseCobol85Listener) ExitPasswordClause(ctx *PasswordClauseContext) {}

// EnterFileStatusClause is called when production fileStatusClause is entered.
func (s *BaseCobol85Listener) EnterFileStatusClause(ctx *FileStatusClauseContext) {}

// ExitFileStatusClause is called when production fileStatusClause is exited.
func (s *BaseCobol85Listener) ExitFileStatusClause(ctx *FileStatusClauseContext) {}

// EnterRelativeKeyClause is called when production relativeKeyClause is entered.
func (s *BaseCobol85Listener) EnterRelativeKeyClause(ctx *RelativeKeyClauseContext) {}

// ExitRelativeKeyClause is called when production relativeKeyClause is exited.
func (s *BaseCobol85Listener) ExitRelativeKeyClause(ctx *RelativeKeyClauseContext) {}

// EnterIoControlParagraph is called when production ioControlParagraph is entered.
func (s *BaseCobol85Listener) EnterIoControlParagraph(ctx *IoControlParagraphContext) {}

// ExitIoControlParagraph is called when production ioControlParagraph is exited.
func (s *BaseCobol85Listener) ExitIoControlParagraph(ctx *IoControlParagraphContext) {}

// EnterIoControlClause is called when production ioControlClause is entered.
func (s *BaseCobol85Listener) EnterIoControlClause(ctx *IoControlClauseContext) {}

// ExitIoControlClause is called when production ioControlClause is exited.
func (s *BaseCobol85Listener) ExitIoControlClause(ctx *IoControlClauseContext) {}

// EnterRerunClause is called when production rerunClause is entered.
func (s *BaseCobol85Listener) EnterRerunClause(ctx *RerunClauseContext) {}

// ExitRerunClause is called when production rerunClause is exited.
func (s *BaseCobol85Listener) ExitRerunClause(ctx *RerunClauseContext) {}

// EnterRerunEveryRecords is called when production rerunEveryRecords is entered.
func (s *BaseCobol85Listener) EnterRerunEveryRecords(ctx *RerunEveryRecordsContext) {}

// ExitRerunEveryRecords is called when production rerunEveryRecords is exited.
func (s *BaseCobol85Listener) ExitRerunEveryRecords(ctx *RerunEveryRecordsContext) {}

// EnterRerunEveryOf is called when production rerunEveryOf is entered.
func (s *BaseCobol85Listener) EnterRerunEveryOf(ctx *RerunEveryOfContext) {}

// ExitRerunEveryOf is called when production rerunEveryOf is exited.
func (s *BaseCobol85Listener) ExitRerunEveryOf(ctx *RerunEveryOfContext) {}

// EnterRerunEveryClock is called when production rerunEveryClock is entered.
func (s *BaseCobol85Listener) EnterRerunEveryClock(ctx *RerunEveryClockContext) {}

// ExitRerunEveryClock is called when production rerunEveryClock is exited.
func (s *BaseCobol85Listener) ExitRerunEveryClock(ctx *RerunEveryClockContext) {}

// EnterSameClause is called when production sameClause is entered.
func (s *BaseCobol85Listener) EnterSameClause(ctx *SameClauseContext) {}

// ExitSameClause is called when production sameClause is exited.
func (s *BaseCobol85Listener) ExitSameClause(ctx *SameClauseContext) {}

// EnterMultipleFileClause is called when production multipleFileClause is entered.
func (s *BaseCobol85Listener) EnterMultipleFileClause(ctx *MultipleFileClauseContext) {}

// ExitMultipleFileClause is called when production multipleFileClause is exited.
func (s *BaseCobol85Listener) ExitMultipleFileClause(ctx *MultipleFileClauseContext) {}

// EnterMultipleFilePosition is called when production multipleFilePosition is entered.
func (s *BaseCobol85Listener) EnterMultipleFilePosition(ctx *MultipleFilePositionContext) {}

// ExitMultipleFilePosition is called when production multipleFilePosition is exited.
func (s *BaseCobol85Listener) ExitMultipleFilePosition(ctx *MultipleFilePositionContext) {}

// EnterCommitmentControlClause is called when production commitmentControlClause is entered.
func (s *BaseCobol85Listener) EnterCommitmentControlClause(ctx *CommitmentControlClauseContext) {}

// ExitCommitmentControlClause is called when production commitmentControlClause is exited.
func (s *BaseCobol85Listener) ExitCommitmentControlClause(ctx *CommitmentControlClauseContext) {}

// EnterDataDivision is called when production dataDivision is entered.
func (s *BaseCobol85Listener) EnterDataDivision(ctx *DataDivisionContext) {}

// ExitDataDivision is called when production dataDivision is exited.
func (s *BaseCobol85Listener) ExitDataDivision(ctx *DataDivisionContext) {}

// EnterDataDivisionSection is called when production dataDivisionSection is entered.
func (s *BaseCobol85Listener) EnterDataDivisionSection(ctx *DataDivisionSectionContext) {}

// ExitDataDivisionSection is called when production dataDivisionSection is exited.
func (s *BaseCobol85Listener) ExitDataDivisionSection(ctx *DataDivisionSectionContext) {}

// EnterFileSection is called when production fileSection is entered.
func (s *BaseCobol85Listener) EnterFileSection(ctx *FileSectionContext) {}

// ExitFileSection is called when production fileSection is exited.
func (s *BaseCobol85Listener) ExitFileSection(ctx *FileSectionContext) {}

// EnterFileDescriptionEntry is called when production fileDescriptionEntry is entered.
func (s *BaseCobol85Listener) EnterFileDescriptionEntry(ctx *FileDescriptionEntryContext) {}

// ExitFileDescriptionEntry is called when production fileDescriptionEntry is exited.
func (s *BaseCobol85Listener) ExitFileDescriptionEntry(ctx *FileDescriptionEntryContext) {}

// EnterFileDescriptionEntryClause is called when production fileDescriptionEntryClause is entered.
func (s *BaseCobol85Listener) EnterFileDescriptionEntryClause(ctx *FileDescriptionEntryClauseContext) {
}

// ExitFileDescriptionEntryClause is called when production fileDescriptionEntryClause is exited.
func (s *BaseCobol85Listener) ExitFileDescriptionEntryClause(ctx *FileDescriptionEntryClauseContext) {}

// EnterExternalClause is called when production externalClause is entered.
func (s *BaseCobol85Listener) EnterExternalClause(ctx *ExternalClauseContext) {}

// ExitExternalClause is called when production externalClause is exited.
func (s *BaseCobol85Listener) ExitExternalClause(ctx *ExternalClauseContext) {}

// EnterGlobalClause is called when production globalClause is entered.
func (s *BaseCobol85Listener) EnterGlobalClause(ctx *GlobalClauseContext) {}

// ExitGlobalClause is called when production globalClause is exited.
func (s *BaseCobol85Listener) ExitGlobalClause(ctx *GlobalClauseContext) {}

// EnterBlockContainsClause is called when production blockContainsClause is entered.
func (s *BaseCobol85Listener) EnterBlockContainsClause(ctx *BlockContainsClauseContext) {}

// ExitBlockContainsClause is called when production blockContainsClause is exited.
func (s *BaseCobol85Listener) ExitBlockContainsClause(ctx *BlockContainsClauseContext) {}

// EnterBlockContainsTo is called when production blockContainsTo is entered.
func (s *BaseCobol85Listener) EnterBlockContainsTo(ctx *BlockContainsToContext) {}

// ExitBlockContainsTo is called when production blockContainsTo is exited.
func (s *BaseCobol85Listener) ExitBlockContainsTo(ctx *BlockContainsToContext) {}

// EnterRecordContainsClause is called when production recordContainsClause is entered.
func (s *BaseCobol85Listener) EnterRecordContainsClause(ctx *RecordContainsClauseContext) {}

// ExitRecordContainsClause is called when production recordContainsClause is exited.
func (s *BaseCobol85Listener) ExitRecordContainsClause(ctx *RecordContainsClauseContext) {}

// EnterRecordContainsClauseFormat1 is called when production recordContainsClauseFormat1 is entered.
func (s *BaseCobol85Listener) EnterRecordContainsClauseFormat1(ctx *RecordContainsClauseFormat1Context) {
}

// ExitRecordContainsClauseFormat1 is called when production recordContainsClauseFormat1 is exited.
func (s *BaseCobol85Listener) ExitRecordContainsClauseFormat1(ctx *RecordContainsClauseFormat1Context) {
}

// EnterRecordContainsClauseFormat2 is called when production recordContainsClauseFormat2 is entered.
func (s *BaseCobol85Listener) EnterRecordContainsClauseFormat2(ctx *RecordContainsClauseFormat2Context) {
}

// ExitRecordContainsClauseFormat2 is called when production recordContainsClauseFormat2 is exited.
func (s *BaseCobol85Listener) ExitRecordContainsClauseFormat2(ctx *RecordContainsClauseFormat2Context) {
}

// EnterRecordContainsClauseFormat3 is called when production recordContainsClauseFormat3 is entered.
func (s *BaseCobol85Listener) EnterRecordContainsClauseFormat3(ctx *RecordContainsClauseFormat3Context) {
}

// ExitRecordContainsClauseFormat3 is called when production recordContainsClauseFormat3 is exited.
func (s *BaseCobol85Listener) ExitRecordContainsClauseFormat3(ctx *RecordContainsClauseFormat3Context) {
}

// EnterRecordContainsTo is called when production recordContainsTo is entered.
func (s *BaseCobol85Listener) EnterRecordContainsTo(ctx *RecordContainsToContext) {}

// ExitRecordContainsTo is called when production recordContainsTo is exited.
func (s *BaseCobol85Listener) ExitRecordContainsTo(ctx *RecordContainsToContext) {}

// EnterLabelRecordsClause is called when production labelRecordsClause is entered.
func (s *BaseCobol85Listener) EnterLabelRecordsClause(ctx *LabelRecordsClauseContext) {}

// ExitLabelRecordsClause is called when production labelRecordsClause is exited.
func (s *BaseCobol85Listener) ExitLabelRecordsClause(ctx *LabelRecordsClauseContext) {}

// EnterValueOfClause is called when production valueOfClause is entered.
func (s *BaseCobol85Listener) EnterValueOfClause(ctx *ValueOfClauseContext) {}

// ExitValueOfClause is called when production valueOfClause is exited.
func (s *BaseCobol85Listener) ExitValueOfClause(ctx *ValueOfClauseContext) {}

// EnterValuePair is called when production valuePair is entered.
func (s *BaseCobol85Listener) EnterValuePair(ctx *ValuePairContext) {}

// ExitValuePair is called when production valuePair is exited.
func (s *BaseCobol85Listener) ExitValuePair(ctx *ValuePairContext) {}

// EnterDataRecordsClause is called when production dataRecordsClause is entered.
func (s *BaseCobol85Listener) EnterDataRecordsClause(ctx *DataRecordsClauseContext) {}

// ExitDataRecordsClause is called when production dataRecordsClause is exited.
func (s *BaseCobol85Listener) ExitDataRecordsClause(ctx *DataRecordsClauseContext) {}

// EnterLinageClause is called when production linageClause is entered.
func (s *BaseCobol85Listener) EnterLinageClause(ctx *LinageClauseContext) {}

// ExitLinageClause is called when production linageClause is exited.
func (s *BaseCobol85Listener) ExitLinageClause(ctx *LinageClauseContext) {}

// EnterLinageAt is called when production linageAt is entered.
func (s *BaseCobol85Listener) EnterLinageAt(ctx *LinageAtContext) {}

// ExitLinageAt is called when production linageAt is exited.
func (s *BaseCobol85Listener) ExitLinageAt(ctx *LinageAtContext) {}

// EnterLinageFootingAt is called when production linageFootingAt is entered.
func (s *BaseCobol85Listener) EnterLinageFootingAt(ctx *LinageFootingAtContext) {}

// ExitLinageFootingAt is called when production linageFootingAt is exited.
func (s *BaseCobol85Listener) ExitLinageFootingAt(ctx *LinageFootingAtContext) {}

// EnterLinageLinesAtTop is called when production linageLinesAtTop is entered.
func (s *BaseCobol85Listener) EnterLinageLinesAtTop(ctx *LinageLinesAtTopContext) {}

// ExitLinageLinesAtTop is called when production linageLinesAtTop is exited.
func (s *BaseCobol85Listener) ExitLinageLinesAtTop(ctx *LinageLinesAtTopContext) {}

// EnterLinageLinesAtBottom is called when production linageLinesAtBottom is entered.
func (s *BaseCobol85Listener) EnterLinageLinesAtBottom(ctx *LinageLinesAtBottomContext) {}

// ExitLinageLinesAtBottom is called when production linageLinesAtBottom is exited.
func (s *BaseCobol85Listener) ExitLinageLinesAtBottom(ctx *LinageLinesAtBottomContext) {}

// EnterRecordingModeClause is called when production recordingModeClause is entered.
func (s *BaseCobol85Listener) EnterRecordingModeClause(ctx *RecordingModeClauseContext) {}

// ExitRecordingModeClause is called when production recordingModeClause is exited.
func (s *BaseCobol85Listener) ExitRecordingModeClause(ctx *RecordingModeClauseContext) {}

// EnterModeStatement is called when production modeStatement is entered.
func (s *BaseCobol85Listener) EnterModeStatement(ctx *ModeStatementContext) {}

// ExitModeStatement is called when production modeStatement is exited.
func (s *BaseCobol85Listener) ExitModeStatement(ctx *ModeStatementContext) {}

// EnterCodeSetClause is called when production codeSetClause is entered.
func (s *BaseCobol85Listener) EnterCodeSetClause(ctx *CodeSetClauseContext) {}

// ExitCodeSetClause is called when production codeSetClause is exited.
func (s *BaseCobol85Listener) ExitCodeSetClause(ctx *CodeSetClauseContext) {}

// EnterReportClause is called when production reportClause is entered.
func (s *BaseCobol85Listener) EnterReportClause(ctx *ReportClauseContext) {}

// ExitReportClause is called when production reportClause is exited.
func (s *BaseCobol85Listener) ExitReportClause(ctx *ReportClauseContext) {}

// EnterDataBaseSection is called when production dataBaseSection is entered.
func (s *BaseCobol85Listener) EnterDataBaseSection(ctx *DataBaseSectionContext) {}

// ExitDataBaseSection is called when production dataBaseSection is exited.
func (s *BaseCobol85Listener) ExitDataBaseSection(ctx *DataBaseSectionContext) {}

// EnterDataBaseSectionEntry is called when production dataBaseSectionEntry is entered.
func (s *BaseCobol85Listener) EnterDataBaseSectionEntry(ctx *DataBaseSectionEntryContext) {}

// ExitDataBaseSectionEntry is called when production dataBaseSectionEntry is exited.
func (s *BaseCobol85Listener) ExitDataBaseSectionEntry(ctx *DataBaseSectionEntryContext) {}

// EnterWorkingStorageSection is called when production workingStorageSection is entered.
func (s *BaseCobol85Listener) EnterWorkingStorageSection(ctx *WorkingStorageSectionContext) {}

// ExitWorkingStorageSection is called when production workingStorageSection is exited.
func (s *BaseCobol85Listener) ExitWorkingStorageSection(ctx *WorkingStorageSectionContext) {}

// EnterLinkageSection is called when production linkageSection is entered.
func (s *BaseCobol85Listener) EnterLinkageSection(ctx *LinkageSectionContext) {}

// ExitLinkageSection is called when production linkageSection is exited.
func (s *BaseCobol85Listener) ExitLinkageSection(ctx *LinkageSectionContext) {}

// EnterCommunicationSection is called when production communicationSection is entered.
func (s *BaseCobol85Listener) EnterCommunicationSection(ctx *CommunicationSectionContext) {}

// ExitCommunicationSection is called when production communicationSection is exited.
func (s *BaseCobol85Listener) ExitCommunicationSection(ctx *CommunicationSectionContext) {}

// EnterCommunicationDescriptionEntry is called when production communicationDescriptionEntry is entered.
func (s *BaseCobol85Listener) EnterCommunicationDescriptionEntry(ctx *CommunicationDescriptionEntryContext) {
}

// ExitCommunicationDescriptionEntry is called when production communicationDescriptionEntry is exited.
func (s *BaseCobol85Listener) ExitCommunicationDescriptionEntry(ctx *CommunicationDescriptionEntryContext) {
}

// EnterCommunicationDescriptionEntryFormat1 is called when production communicationDescriptionEntryFormat1 is entered.
func (s *BaseCobol85Listener) EnterCommunicationDescriptionEntryFormat1(ctx *CommunicationDescriptionEntryFormat1Context) {
}

// ExitCommunicationDescriptionEntryFormat1 is called when production communicationDescriptionEntryFormat1 is exited.
func (s *BaseCobol85Listener) ExitCommunicationDescriptionEntryFormat1(ctx *CommunicationDescriptionEntryFormat1Context) {
}

// EnterCommunicationDescriptionEntryFormat2 is called when production communicationDescriptionEntryFormat2 is entered.
func (s *BaseCobol85Listener) EnterCommunicationDescriptionEntryFormat2(ctx *CommunicationDescriptionEntryFormat2Context) {
}

// ExitCommunicationDescriptionEntryFormat2 is called when production communicationDescriptionEntryFormat2 is exited.
func (s *BaseCobol85Listener) ExitCommunicationDescriptionEntryFormat2(ctx *CommunicationDescriptionEntryFormat2Context) {
}

// EnterCommunicationDescriptionEntryFormat3 is called when production communicationDescriptionEntryFormat3 is entered.
func (s *BaseCobol85Listener) EnterCommunicationDescriptionEntryFormat3(ctx *CommunicationDescriptionEntryFormat3Context) {
}

// ExitCommunicationDescriptionEntryFormat3 is called when production communicationDescriptionEntryFormat3 is exited.
func (s *BaseCobol85Listener) ExitCommunicationDescriptionEntryFormat3(ctx *CommunicationDescriptionEntryFormat3Context) {
}

// EnterDestinationCountClause is called when production destinationCountClause is entered.
func (s *BaseCobol85Listener) EnterDestinationCountClause(ctx *DestinationCountClauseContext) {}

// ExitDestinationCountClause is called when production destinationCountClause is exited.
func (s *BaseCobol85Listener) ExitDestinationCountClause(ctx *DestinationCountClauseContext) {}

// EnterDestinationTableClause is called when production destinationTableClause is entered.
func (s *BaseCobol85Listener) EnterDestinationTableClause(ctx *DestinationTableClauseContext) {}

// ExitDestinationTableClause is called when production destinationTableClause is exited.
func (s *BaseCobol85Listener) ExitDestinationTableClause(ctx *DestinationTableClauseContext) {}

// EnterEndKeyClause is called when production endKeyClause is entered.
func (s *BaseCobol85Listener) EnterEndKeyClause(ctx *EndKeyClauseContext) {}

// ExitEndKeyClause is called when production endKeyClause is exited.
func (s *BaseCobol85Listener) ExitEndKeyClause(ctx *EndKeyClauseContext) {}

// EnterErrorKeyClause is called when production errorKeyClause is entered.
func (s *BaseCobol85Listener) EnterErrorKeyClause(ctx *ErrorKeyClauseContext) {}

// ExitErrorKeyClause is called when production errorKeyClause is exited.
func (s *BaseCobol85Listener) ExitErrorKeyClause(ctx *ErrorKeyClauseContext) {}

// EnterMessageCountClause is called when production messageCountClause is entered.
func (s *BaseCobol85Listener) EnterMessageCountClause(ctx *MessageCountClauseContext) {}

// ExitMessageCountClause is called when production messageCountClause is exited.
func (s *BaseCobol85Listener) ExitMessageCountClause(ctx *MessageCountClauseContext) {}

// EnterMessageDateClause is called when production messageDateClause is entered.
func (s *BaseCobol85Listener) EnterMessageDateClause(ctx *MessageDateClauseContext) {}

// ExitMessageDateClause is called when production messageDateClause is exited.
func (s *BaseCobol85Listener) ExitMessageDateClause(ctx *MessageDateClauseContext) {}

// EnterMessageTimeClause is called when production messageTimeClause is entered.
func (s *BaseCobol85Listener) EnterMessageTimeClause(ctx *MessageTimeClauseContext) {}

// ExitMessageTimeClause is called when production messageTimeClause is exited.
func (s *BaseCobol85Listener) ExitMessageTimeClause(ctx *MessageTimeClauseContext) {}

// EnterStatusKeyClause is called when production statusKeyClause is entered.
func (s *BaseCobol85Listener) EnterStatusKeyClause(ctx *StatusKeyClauseContext) {}

// ExitStatusKeyClause is called when production statusKeyClause is exited.
func (s *BaseCobol85Listener) ExitStatusKeyClause(ctx *StatusKeyClauseContext) {}

// EnterSymbolicDestinationClause is called when production symbolicDestinationClause is entered.
func (s *BaseCobol85Listener) EnterSymbolicDestinationClause(ctx *SymbolicDestinationClauseContext) {}

// ExitSymbolicDestinationClause is called when production symbolicDestinationClause is exited.
func (s *BaseCobol85Listener) ExitSymbolicDestinationClause(ctx *SymbolicDestinationClauseContext) {}

// EnterSymbolicQueueClause is called when production symbolicQueueClause is entered.
func (s *BaseCobol85Listener) EnterSymbolicQueueClause(ctx *SymbolicQueueClauseContext) {}

// ExitSymbolicQueueClause is called when production symbolicQueueClause is exited.
func (s *BaseCobol85Listener) ExitSymbolicQueueClause(ctx *SymbolicQueueClauseContext) {}

// EnterSymbolicSourceClause is called when production symbolicSourceClause is entered.
func (s *BaseCobol85Listener) EnterSymbolicSourceClause(ctx *SymbolicSourceClauseContext) {}

// ExitSymbolicSourceClause is called when production symbolicSourceClause is exited.
func (s *BaseCobol85Listener) ExitSymbolicSourceClause(ctx *SymbolicSourceClauseContext) {}

// EnterSymbolicTerminalClause is called when production symbolicTerminalClause is entered.
func (s *BaseCobol85Listener) EnterSymbolicTerminalClause(ctx *SymbolicTerminalClauseContext) {}

// ExitSymbolicTerminalClause is called when production symbolicTerminalClause is exited.
func (s *BaseCobol85Listener) ExitSymbolicTerminalClause(ctx *SymbolicTerminalClauseContext) {}

// EnterSymbolicSubQueueClause is called when production symbolicSubQueueClause is entered.
func (s *BaseCobol85Listener) EnterSymbolicSubQueueClause(ctx *SymbolicSubQueueClauseContext) {}

// ExitSymbolicSubQueueClause is called when production symbolicSubQueueClause is exited.
func (s *BaseCobol85Listener) ExitSymbolicSubQueueClause(ctx *SymbolicSubQueueClauseContext) {}

// EnterTextLengthClause is called when production textLengthClause is entered.
func (s *BaseCobol85Listener) EnterTextLengthClause(ctx *TextLengthClauseContext) {}

// ExitTextLengthClause is called when production textLengthClause is exited.
func (s *BaseCobol85Listener) ExitTextLengthClause(ctx *TextLengthClauseContext) {}

// EnterLocalStorageSection is called when production localStorageSection is entered.
func (s *BaseCobol85Listener) EnterLocalStorageSection(ctx *LocalStorageSectionContext) {}

// ExitLocalStorageSection is called when production localStorageSection is exited.
func (s *BaseCobol85Listener) ExitLocalStorageSection(ctx *LocalStorageSectionContext) {}

// EnterScreenSection is called when production screenSection is entered.
func (s *BaseCobol85Listener) EnterScreenSection(ctx *ScreenSectionContext) {}

// ExitScreenSection is called when production screenSection is exited.
func (s *BaseCobol85Listener) ExitScreenSection(ctx *ScreenSectionContext) {}

// EnterScreenDescriptionEntry is called when production screenDescriptionEntry is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionEntry(ctx *ScreenDescriptionEntryContext) {}

// ExitScreenDescriptionEntry is called when production screenDescriptionEntry is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionEntry(ctx *ScreenDescriptionEntryContext) {}

// EnterScreenDescriptionBlankClause is called when production screenDescriptionBlankClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionBlankClause(ctx *ScreenDescriptionBlankClauseContext) {
}

// ExitScreenDescriptionBlankClause is called when production screenDescriptionBlankClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionBlankClause(ctx *ScreenDescriptionBlankClauseContext) {
}

// EnterScreenDescriptionBellClause is called when production screenDescriptionBellClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionBellClause(ctx *ScreenDescriptionBellClauseContext) {
}

// ExitScreenDescriptionBellClause is called when production screenDescriptionBellClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionBellClause(ctx *ScreenDescriptionBellClauseContext) {
}

// EnterScreenDescriptionBlinkClause is called when production screenDescriptionBlinkClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionBlinkClause(ctx *ScreenDescriptionBlinkClauseContext) {
}

// ExitScreenDescriptionBlinkClause is called when production screenDescriptionBlinkClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionBlinkClause(ctx *ScreenDescriptionBlinkClauseContext) {
}

// EnterScreenDescriptionEraseClause is called when production screenDescriptionEraseClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionEraseClause(ctx *ScreenDescriptionEraseClauseContext) {
}

// ExitScreenDescriptionEraseClause is called when production screenDescriptionEraseClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionEraseClause(ctx *ScreenDescriptionEraseClauseContext) {
}

// EnterScreenDescriptionLightClause is called when production screenDescriptionLightClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionLightClause(ctx *ScreenDescriptionLightClauseContext) {
}

// ExitScreenDescriptionLightClause is called when production screenDescriptionLightClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionLightClause(ctx *ScreenDescriptionLightClauseContext) {
}

// EnterScreenDescriptionGridClause is called when production screenDescriptionGridClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionGridClause(ctx *ScreenDescriptionGridClauseContext) {
}

// ExitScreenDescriptionGridClause is called when production screenDescriptionGridClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionGridClause(ctx *ScreenDescriptionGridClauseContext) {
}

// EnterScreenDescriptionReverseVideoClause is called when production screenDescriptionReverseVideoClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionReverseVideoClause(ctx *ScreenDescriptionReverseVideoClauseContext) {
}

// ExitScreenDescriptionReverseVideoClause is called when production screenDescriptionReverseVideoClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionReverseVideoClause(ctx *ScreenDescriptionReverseVideoClauseContext) {
}

// EnterScreenDescriptionUnderlineClause is called when production screenDescriptionUnderlineClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionUnderlineClause(ctx *ScreenDescriptionUnderlineClauseContext) {
}

// ExitScreenDescriptionUnderlineClause is called when production screenDescriptionUnderlineClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionUnderlineClause(ctx *ScreenDescriptionUnderlineClauseContext) {
}

// EnterScreenDescriptionSizeClause is called when production screenDescriptionSizeClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionSizeClause(ctx *ScreenDescriptionSizeClauseContext) {
}

// ExitScreenDescriptionSizeClause is called when production screenDescriptionSizeClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionSizeClause(ctx *ScreenDescriptionSizeClauseContext) {
}

// EnterScreenDescriptionLineClause is called when production screenDescriptionLineClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionLineClause(ctx *ScreenDescriptionLineClauseContext) {
}

// ExitScreenDescriptionLineClause is called when production screenDescriptionLineClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionLineClause(ctx *ScreenDescriptionLineClauseContext) {
}

// EnterScreenDescriptionColumnClause is called when production screenDescriptionColumnClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionColumnClause(ctx *ScreenDescriptionColumnClauseContext) {
}

// ExitScreenDescriptionColumnClause is called when production screenDescriptionColumnClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionColumnClause(ctx *ScreenDescriptionColumnClauseContext) {
}

// EnterScreenDescriptionForegroundColorClause is called when production screenDescriptionForegroundColorClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionForegroundColorClause(ctx *ScreenDescriptionForegroundColorClauseContext) {
}

// ExitScreenDescriptionForegroundColorClause is called when production screenDescriptionForegroundColorClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionForegroundColorClause(ctx *ScreenDescriptionForegroundColorClauseContext) {
}

// EnterScreenDescriptionBackgroundColorClause is called when production screenDescriptionBackgroundColorClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionBackgroundColorClause(ctx *ScreenDescriptionBackgroundColorClauseContext) {
}

// ExitScreenDescriptionBackgroundColorClause is called when production screenDescriptionBackgroundColorClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionBackgroundColorClause(ctx *ScreenDescriptionBackgroundColorClauseContext) {
}

// EnterScreenDescriptionControlClause is called when production screenDescriptionControlClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionControlClause(ctx *ScreenDescriptionControlClauseContext) {
}

// ExitScreenDescriptionControlClause is called when production screenDescriptionControlClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionControlClause(ctx *ScreenDescriptionControlClauseContext) {
}

// EnterScreenDescriptionValueClause is called when production screenDescriptionValueClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionValueClause(ctx *ScreenDescriptionValueClauseContext) {
}

// ExitScreenDescriptionValueClause is called when production screenDescriptionValueClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionValueClause(ctx *ScreenDescriptionValueClauseContext) {
}

// EnterScreenDescriptionPictureClause is called when production screenDescriptionPictureClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionPictureClause(ctx *ScreenDescriptionPictureClauseContext) {
}

// ExitScreenDescriptionPictureClause is called when production screenDescriptionPictureClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionPictureClause(ctx *ScreenDescriptionPictureClauseContext) {
}

// EnterScreenDescriptionFromClause is called when production screenDescriptionFromClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionFromClause(ctx *ScreenDescriptionFromClauseContext) {
}

// ExitScreenDescriptionFromClause is called when production screenDescriptionFromClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionFromClause(ctx *ScreenDescriptionFromClauseContext) {
}

// EnterScreenDescriptionToClause is called when production screenDescriptionToClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionToClause(ctx *ScreenDescriptionToClauseContext) {}

// ExitScreenDescriptionToClause is called when production screenDescriptionToClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionToClause(ctx *ScreenDescriptionToClauseContext) {}

// EnterScreenDescriptionUsingClause is called when production screenDescriptionUsingClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionUsingClause(ctx *ScreenDescriptionUsingClauseContext) {
}

// ExitScreenDescriptionUsingClause is called when production screenDescriptionUsingClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionUsingClause(ctx *ScreenDescriptionUsingClauseContext) {
}

// EnterScreenDescriptionUsageClause is called when production screenDescriptionUsageClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionUsageClause(ctx *ScreenDescriptionUsageClauseContext) {
}

// ExitScreenDescriptionUsageClause is called when production screenDescriptionUsageClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionUsageClause(ctx *ScreenDescriptionUsageClauseContext) {
}

// EnterScreenDescriptionBlankWhenZeroClause is called when production screenDescriptionBlankWhenZeroClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionBlankWhenZeroClause(ctx *ScreenDescriptionBlankWhenZeroClauseContext) {
}

// ExitScreenDescriptionBlankWhenZeroClause is called when production screenDescriptionBlankWhenZeroClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionBlankWhenZeroClause(ctx *ScreenDescriptionBlankWhenZeroClauseContext) {
}

// EnterScreenDescriptionJustifiedClause is called when production screenDescriptionJustifiedClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionJustifiedClause(ctx *ScreenDescriptionJustifiedClauseContext) {
}

// ExitScreenDescriptionJustifiedClause is called when production screenDescriptionJustifiedClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionJustifiedClause(ctx *ScreenDescriptionJustifiedClauseContext) {
}

// EnterScreenDescriptionSignClause is called when production screenDescriptionSignClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionSignClause(ctx *ScreenDescriptionSignClauseContext) {
}

// ExitScreenDescriptionSignClause is called when production screenDescriptionSignClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionSignClause(ctx *ScreenDescriptionSignClauseContext) {
}

// EnterScreenDescriptionAutoClause is called when production screenDescriptionAutoClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionAutoClause(ctx *ScreenDescriptionAutoClauseContext) {
}

// ExitScreenDescriptionAutoClause is called when production screenDescriptionAutoClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionAutoClause(ctx *ScreenDescriptionAutoClauseContext) {
}

// EnterScreenDescriptionSecureClause is called when production screenDescriptionSecureClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionSecureClause(ctx *ScreenDescriptionSecureClauseContext) {
}

// ExitScreenDescriptionSecureClause is called when production screenDescriptionSecureClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionSecureClause(ctx *ScreenDescriptionSecureClauseContext) {
}

// EnterScreenDescriptionRequiredClause is called when production screenDescriptionRequiredClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionRequiredClause(ctx *ScreenDescriptionRequiredClauseContext) {
}

// ExitScreenDescriptionRequiredClause is called when production screenDescriptionRequiredClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionRequiredClause(ctx *ScreenDescriptionRequiredClauseContext) {
}

// EnterScreenDescriptionPromptClause is called when production screenDescriptionPromptClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionPromptClause(ctx *ScreenDescriptionPromptClauseContext) {
}

// ExitScreenDescriptionPromptClause is called when production screenDescriptionPromptClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionPromptClause(ctx *ScreenDescriptionPromptClauseContext) {
}

// EnterScreenDescriptionPromptOccursClause is called when production screenDescriptionPromptOccursClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionPromptOccursClause(ctx *ScreenDescriptionPromptOccursClauseContext) {
}

// ExitScreenDescriptionPromptOccursClause is called when production screenDescriptionPromptOccursClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionPromptOccursClause(ctx *ScreenDescriptionPromptOccursClauseContext) {
}

// EnterScreenDescriptionFullClause is called when production screenDescriptionFullClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionFullClause(ctx *ScreenDescriptionFullClauseContext) {
}

// ExitScreenDescriptionFullClause is called when production screenDescriptionFullClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionFullClause(ctx *ScreenDescriptionFullClauseContext) {
}

// EnterScreenDescriptionZeroFillClause is called when production screenDescriptionZeroFillClause is entered.
func (s *BaseCobol85Listener) EnterScreenDescriptionZeroFillClause(ctx *ScreenDescriptionZeroFillClauseContext) {
}

// ExitScreenDescriptionZeroFillClause is called when production screenDescriptionZeroFillClause is exited.
func (s *BaseCobol85Listener) ExitScreenDescriptionZeroFillClause(ctx *ScreenDescriptionZeroFillClauseContext) {
}

// EnterReportSection is called when production reportSection is entered.
func (s *BaseCobol85Listener) EnterReportSection(ctx *ReportSectionContext) {}

// ExitReportSection is called when production reportSection is exited.
func (s *BaseCobol85Listener) ExitReportSection(ctx *ReportSectionContext) {}

// EnterReportDescription is called when production reportDescription is entered.
func (s *BaseCobol85Listener) EnterReportDescription(ctx *ReportDescriptionContext) {}

// ExitReportDescription is called when production reportDescription is exited.
func (s *BaseCobol85Listener) ExitReportDescription(ctx *ReportDescriptionContext) {}

// EnterReportDescriptionEntry is called when production reportDescriptionEntry is entered.
func (s *BaseCobol85Listener) EnterReportDescriptionEntry(ctx *ReportDescriptionEntryContext) {}

// ExitReportDescriptionEntry is called when production reportDescriptionEntry is exited.
func (s *BaseCobol85Listener) ExitReportDescriptionEntry(ctx *ReportDescriptionEntryContext) {}

// EnterReportDescriptionGlobalClause is called when production reportDescriptionGlobalClause is entered.
func (s *BaseCobol85Listener) EnterReportDescriptionGlobalClause(ctx *ReportDescriptionGlobalClauseContext) {
}

// ExitReportDescriptionGlobalClause is called when production reportDescriptionGlobalClause is exited.
func (s *BaseCobol85Listener) ExitReportDescriptionGlobalClause(ctx *ReportDescriptionGlobalClauseContext) {
}

// EnterReportDescriptionPageLimitClause is called when production reportDescriptionPageLimitClause is entered.
func (s *BaseCobol85Listener) EnterReportDescriptionPageLimitClause(ctx *ReportDescriptionPageLimitClauseContext) {
}

// ExitReportDescriptionPageLimitClause is called when production reportDescriptionPageLimitClause is exited.
func (s *BaseCobol85Listener) ExitReportDescriptionPageLimitClause(ctx *ReportDescriptionPageLimitClauseContext) {
}

// EnterReportDescriptionHeadingClause is called when production reportDescriptionHeadingClause is entered.
func (s *BaseCobol85Listener) EnterReportDescriptionHeadingClause(ctx *ReportDescriptionHeadingClauseContext) {
}

// ExitReportDescriptionHeadingClause is called when production reportDescriptionHeadingClause is exited.
func (s *BaseCobol85Listener) ExitReportDescriptionHeadingClause(ctx *ReportDescriptionHeadingClauseContext) {
}

// EnterReportDescriptionFirstDetailClause is called when production reportDescriptionFirstDetailClause is entered.
func (s *BaseCobol85Listener) EnterReportDescriptionFirstDetailClause(ctx *ReportDescriptionFirstDetailClauseContext) {
}

// ExitReportDescriptionFirstDetailClause is called when production reportDescriptionFirstDetailClause is exited.
func (s *BaseCobol85Listener) ExitReportDescriptionFirstDetailClause(ctx *ReportDescriptionFirstDetailClauseContext) {
}

// EnterReportDescriptionLastDetailClause is called when production reportDescriptionLastDetailClause is entered.
func (s *BaseCobol85Listener) EnterReportDescriptionLastDetailClause(ctx *ReportDescriptionLastDetailClauseContext) {
}

// ExitReportDescriptionLastDetailClause is called when production reportDescriptionLastDetailClause is exited.
func (s *BaseCobol85Listener) ExitReportDescriptionLastDetailClause(ctx *ReportDescriptionLastDetailClauseContext) {
}

// EnterReportDescriptionFootingClause is called when production reportDescriptionFootingClause is entered.
func (s *BaseCobol85Listener) EnterReportDescriptionFootingClause(ctx *ReportDescriptionFootingClauseContext) {
}

// ExitReportDescriptionFootingClause is called when production reportDescriptionFootingClause is exited.
func (s *BaseCobol85Listener) ExitReportDescriptionFootingClause(ctx *ReportDescriptionFootingClauseContext) {
}

// EnterReportGroupDescriptionEntry is called when production reportGroupDescriptionEntry is entered.
func (s *BaseCobol85Listener) EnterReportGroupDescriptionEntry(ctx *ReportGroupDescriptionEntryContext) {
}

// ExitReportGroupDescriptionEntry is called when production reportGroupDescriptionEntry is exited.
func (s *BaseCobol85Listener) ExitReportGroupDescriptionEntry(ctx *ReportGroupDescriptionEntryContext) {
}

// EnterReportGroupDescriptionEntryFormat1 is called when production reportGroupDescriptionEntryFormat1 is entered.
func (s *BaseCobol85Listener) EnterReportGroupDescriptionEntryFormat1(ctx *ReportGroupDescriptionEntryFormat1Context) {
}

// ExitReportGroupDescriptionEntryFormat1 is called when production reportGroupDescriptionEntryFormat1 is exited.
func (s *BaseCobol85Listener) ExitReportGroupDescriptionEntryFormat1(ctx *ReportGroupDescriptionEntryFormat1Context) {
}

// EnterReportGroupDescriptionEntryFormat2 is called when production reportGroupDescriptionEntryFormat2 is entered.
func (s *BaseCobol85Listener) EnterReportGroupDescriptionEntryFormat2(ctx *ReportGroupDescriptionEntryFormat2Context) {
}

// ExitReportGroupDescriptionEntryFormat2 is called when production reportGroupDescriptionEntryFormat2 is exited.
func (s *BaseCobol85Listener) ExitReportGroupDescriptionEntryFormat2(ctx *ReportGroupDescriptionEntryFormat2Context) {
}

// EnterReportGroupDescriptionEntryFormat3 is called when production reportGroupDescriptionEntryFormat3 is entered.
func (s *BaseCobol85Listener) EnterReportGroupDescriptionEntryFormat3(ctx *ReportGroupDescriptionEntryFormat3Context) {
}

// ExitReportGroupDescriptionEntryFormat3 is called when production reportGroupDescriptionEntryFormat3 is exited.
func (s *BaseCobol85Listener) ExitReportGroupDescriptionEntryFormat3(ctx *ReportGroupDescriptionEntryFormat3Context) {
}

// EnterReportGroupBlankWhenZeroClause is called when production reportGroupBlankWhenZeroClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupBlankWhenZeroClause(ctx *ReportGroupBlankWhenZeroClauseContext) {
}

// ExitReportGroupBlankWhenZeroClause is called when production reportGroupBlankWhenZeroClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupBlankWhenZeroClause(ctx *ReportGroupBlankWhenZeroClauseContext) {
}

// EnterReportGroupColumnNumberClause is called when production reportGroupColumnNumberClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupColumnNumberClause(ctx *ReportGroupColumnNumberClauseContext) {
}

// ExitReportGroupColumnNumberClause is called when production reportGroupColumnNumberClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupColumnNumberClause(ctx *ReportGroupColumnNumberClauseContext) {
}

// EnterReportGroupIndicateClause is called when production reportGroupIndicateClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupIndicateClause(ctx *ReportGroupIndicateClauseContext) {}

// ExitReportGroupIndicateClause is called when production reportGroupIndicateClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupIndicateClause(ctx *ReportGroupIndicateClauseContext) {}

// EnterReportGroupJustifiedClause is called when production reportGroupJustifiedClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupJustifiedClause(ctx *ReportGroupJustifiedClauseContext) {
}

// ExitReportGroupJustifiedClause is called when production reportGroupJustifiedClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupJustifiedClause(ctx *ReportGroupJustifiedClauseContext) {}

// EnterReportGroupLineNumberClause is called when production reportGroupLineNumberClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupLineNumberClause(ctx *ReportGroupLineNumberClauseContext) {
}

// ExitReportGroupLineNumberClause is called when production reportGroupLineNumberClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupLineNumberClause(ctx *ReportGroupLineNumberClauseContext) {
}

// EnterReportGroupLineNumberNextPage is called when production reportGroupLineNumberNextPage is entered.
func (s *BaseCobol85Listener) EnterReportGroupLineNumberNextPage(ctx *ReportGroupLineNumberNextPageContext) {
}

// ExitReportGroupLineNumberNextPage is called when production reportGroupLineNumberNextPage is exited.
func (s *BaseCobol85Listener) ExitReportGroupLineNumberNextPage(ctx *ReportGroupLineNumberNextPageContext) {
}

// EnterReportGroupLineNumberPlus is called when production reportGroupLineNumberPlus is entered.
func (s *BaseCobol85Listener) EnterReportGroupLineNumberPlus(ctx *ReportGroupLineNumberPlusContext) {}

// ExitReportGroupLineNumberPlus is called when production reportGroupLineNumberPlus is exited.
func (s *BaseCobol85Listener) ExitReportGroupLineNumberPlus(ctx *ReportGroupLineNumberPlusContext) {}

// EnterReportGroupNextGroupClause is called when production reportGroupNextGroupClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupNextGroupClause(ctx *ReportGroupNextGroupClauseContext) {
}

// ExitReportGroupNextGroupClause is called when production reportGroupNextGroupClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupNextGroupClause(ctx *ReportGroupNextGroupClauseContext) {}

// EnterReportGroupNextGroupPlus is called when production reportGroupNextGroupPlus is entered.
func (s *BaseCobol85Listener) EnterReportGroupNextGroupPlus(ctx *ReportGroupNextGroupPlusContext) {}

// ExitReportGroupNextGroupPlus is called when production reportGroupNextGroupPlus is exited.
func (s *BaseCobol85Listener) ExitReportGroupNextGroupPlus(ctx *ReportGroupNextGroupPlusContext) {}

// EnterReportGroupNextGroupNextPage is called when production reportGroupNextGroupNextPage is entered.
func (s *BaseCobol85Listener) EnterReportGroupNextGroupNextPage(ctx *ReportGroupNextGroupNextPageContext) {
}

// ExitReportGroupNextGroupNextPage is called when production reportGroupNextGroupNextPage is exited.
func (s *BaseCobol85Listener) ExitReportGroupNextGroupNextPage(ctx *ReportGroupNextGroupNextPageContext) {
}

// EnterReportGroupPictureClause is called when production reportGroupPictureClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupPictureClause(ctx *ReportGroupPictureClauseContext) {}

// ExitReportGroupPictureClause is called when production reportGroupPictureClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupPictureClause(ctx *ReportGroupPictureClauseContext) {}

// EnterReportGroupResetClause is called when production reportGroupResetClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupResetClause(ctx *ReportGroupResetClauseContext) {}

// ExitReportGroupResetClause is called when production reportGroupResetClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupResetClause(ctx *ReportGroupResetClauseContext) {}

// EnterReportGroupSignClause is called when production reportGroupSignClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupSignClause(ctx *ReportGroupSignClauseContext) {}

// ExitReportGroupSignClause is called when production reportGroupSignClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupSignClause(ctx *ReportGroupSignClauseContext) {}

// EnterReportGroupSourceClause is called when production reportGroupSourceClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupSourceClause(ctx *ReportGroupSourceClauseContext) {}

// ExitReportGroupSourceClause is called when production reportGroupSourceClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupSourceClause(ctx *ReportGroupSourceClauseContext) {}

// EnterReportGroupSumClause is called when production reportGroupSumClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupSumClause(ctx *ReportGroupSumClauseContext) {}

// ExitReportGroupSumClause is called when production reportGroupSumClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupSumClause(ctx *ReportGroupSumClauseContext) {}

// EnterReportGroupTypeClause is called when production reportGroupTypeClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupTypeClause(ctx *ReportGroupTypeClauseContext) {}

// ExitReportGroupTypeClause is called when production reportGroupTypeClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupTypeClause(ctx *ReportGroupTypeClauseContext) {}

// EnterReportGroupTypeReportHeading is called when production reportGroupTypeReportHeading is entered.
func (s *BaseCobol85Listener) EnterReportGroupTypeReportHeading(ctx *ReportGroupTypeReportHeadingContext) {
}

// ExitReportGroupTypeReportHeading is called when production reportGroupTypeReportHeading is exited.
func (s *BaseCobol85Listener) ExitReportGroupTypeReportHeading(ctx *ReportGroupTypeReportHeadingContext) {
}

// EnterReportGroupTypePageHeading is called when production reportGroupTypePageHeading is entered.
func (s *BaseCobol85Listener) EnterReportGroupTypePageHeading(ctx *ReportGroupTypePageHeadingContext) {
}

// ExitReportGroupTypePageHeading is called when production reportGroupTypePageHeading is exited.
func (s *BaseCobol85Listener) ExitReportGroupTypePageHeading(ctx *ReportGroupTypePageHeadingContext) {}

// EnterReportGroupTypeControlHeading is called when production reportGroupTypeControlHeading is entered.
func (s *BaseCobol85Listener) EnterReportGroupTypeControlHeading(ctx *ReportGroupTypeControlHeadingContext) {
}

// ExitReportGroupTypeControlHeading is called when production reportGroupTypeControlHeading is exited.
func (s *BaseCobol85Listener) ExitReportGroupTypeControlHeading(ctx *ReportGroupTypeControlHeadingContext) {
}

// EnterReportGroupTypeDetail is called when production reportGroupTypeDetail is entered.
func (s *BaseCobol85Listener) EnterReportGroupTypeDetail(ctx *ReportGroupTypeDetailContext) {}

// ExitReportGroupTypeDetail is called when production reportGroupTypeDetail is exited.
func (s *BaseCobol85Listener) ExitReportGroupTypeDetail(ctx *ReportGroupTypeDetailContext) {}

// EnterReportGroupTypeControlFooting is called when production reportGroupTypeControlFooting is entered.
func (s *BaseCobol85Listener) EnterReportGroupTypeControlFooting(ctx *ReportGroupTypeControlFootingContext) {
}

// ExitReportGroupTypeControlFooting is called when production reportGroupTypeControlFooting is exited.
func (s *BaseCobol85Listener) ExitReportGroupTypeControlFooting(ctx *ReportGroupTypeControlFootingContext) {
}

// EnterReportGroupUsageClause is called when production reportGroupUsageClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupUsageClause(ctx *ReportGroupUsageClauseContext) {}

// ExitReportGroupUsageClause is called when production reportGroupUsageClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupUsageClause(ctx *ReportGroupUsageClauseContext) {}

// EnterReportGroupTypePageFooting is called when production reportGroupTypePageFooting is entered.
func (s *BaseCobol85Listener) EnterReportGroupTypePageFooting(ctx *ReportGroupTypePageFootingContext) {
}

// ExitReportGroupTypePageFooting is called when production reportGroupTypePageFooting is exited.
func (s *BaseCobol85Listener) ExitReportGroupTypePageFooting(ctx *ReportGroupTypePageFootingContext) {}

// EnterReportGroupTypeReportFooting is called when production reportGroupTypeReportFooting is entered.
func (s *BaseCobol85Listener) EnterReportGroupTypeReportFooting(ctx *ReportGroupTypeReportFootingContext) {
}

// ExitReportGroupTypeReportFooting is called when production reportGroupTypeReportFooting is exited.
func (s *BaseCobol85Listener) ExitReportGroupTypeReportFooting(ctx *ReportGroupTypeReportFootingContext) {
}

// EnterReportGroupValueClause is called when production reportGroupValueClause is entered.
func (s *BaseCobol85Listener) EnterReportGroupValueClause(ctx *ReportGroupValueClauseContext) {}

// ExitReportGroupValueClause is called when production reportGroupValueClause is exited.
func (s *BaseCobol85Listener) ExitReportGroupValueClause(ctx *ReportGroupValueClauseContext) {}

// EnterProgramLibrarySection is called when production programLibrarySection is entered.
func (s *BaseCobol85Listener) EnterProgramLibrarySection(ctx *ProgramLibrarySectionContext) {}

// ExitProgramLibrarySection is called when production programLibrarySection is exited.
func (s *BaseCobol85Listener) ExitProgramLibrarySection(ctx *ProgramLibrarySectionContext) {}

// EnterLibraryDescriptionEntry is called when production libraryDescriptionEntry is entered.
func (s *BaseCobol85Listener) EnterLibraryDescriptionEntry(ctx *LibraryDescriptionEntryContext) {}

// ExitLibraryDescriptionEntry is called when production libraryDescriptionEntry is exited.
func (s *BaseCobol85Listener) ExitLibraryDescriptionEntry(ctx *LibraryDescriptionEntryContext) {}

// EnterLibraryDescriptionEntryFormat1 is called when production libraryDescriptionEntryFormat1 is entered.
func (s *BaseCobol85Listener) EnterLibraryDescriptionEntryFormat1(ctx *LibraryDescriptionEntryFormat1Context) {
}

// ExitLibraryDescriptionEntryFormat1 is called when production libraryDescriptionEntryFormat1 is exited.
func (s *BaseCobol85Listener) ExitLibraryDescriptionEntryFormat1(ctx *LibraryDescriptionEntryFormat1Context) {
}

// EnterLibraryDescriptionEntryFormat2 is called when production libraryDescriptionEntryFormat2 is entered.
func (s *BaseCobol85Listener) EnterLibraryDescriptionEntryFormat2(ctx *LibraryDescriptionEntryFormat2Context) {
}

// ExitLibraryDescriptionEntryFormat2 is called when production libraryDescriptionEntryFormat2 is exited.
func (s *BaseCobol85Listener) ExitLibraryDescriptionEntryFormat2(ctx *LibraryDescriptionEntryFormat2Context) {
}

// EnterLibraryAttributeClauseFormat1 is called when production libraryAttributeClauseFormat1 is entered.
func (s *BaseCobol85Listener) EnterLibraryAttributeClauseFormat1(ctx *LibraryAttributeClauseFormat1Context) {
}

// ExitLibraryAttributeClauseFormat1 is called when production libraryAttributeClauseFormat1 is exited.
func (s *BaseCobol85Listener) ExitLibraryAttributeClauseFormat1(ctx *LibraryAttributeClauseFormat1Context) {
}

// EnterLibraryAttributeClauseFormat2 is called when production libraryAttributeClauseFormat2 is entered.
func (s *BaseCobol85Listener) EnterLibraryAttributeClauseFormat2(ctx *LibraryAttributeClauseFormat2Context) {
}

// ExitLibraryAttributeClauseFormat2 is called when production libraryAttributeClauseFormat2 is exited.
func (s *BaseCobol85Listener) ExitLibraryAttributeClauseFormat2(ctx *LibraryAttributeClauseFormat2Context) {
}

// EnterLibraryAttributeFunction is called when production libraryAttributeFunction is entered.
func (s *BaseCobol85Listener) EnterLibraryAttributeFunction(ctx *LibraryAttributeFunctionContext) {}

// ExitLibraryAttributeFunction is called when production libraryAttributeFunction is exited.
func (s *BaseCobol85Listener) ExitLibraryAttributeFunction(ctx *LibraryAttributeFunctionContext) {}

// EnterLibraryAttributeParameter is called when production libraryAttributeParameter is entered.
func (s *BaseCobol85Listener) EnterLibraryAttributeParameter(ctx *LibraryAttributeParameterContext) {}

// ExitLibraryAttributeParameter is called when production libraryAttributeParameter is exited.
func (s *BaseCobol85Listener) ExitLibraryAttributeParameter(ctx *LibraryAttributeParameterContext) {}

// EnterLibraryAttributeTitle is called when production libraryAttributeTitle is entered.
func (s *BaseCobol85Listener) EnterLibraryAttributeTitle(ctx *LibraryAttributeTitleContext) {}

// ExitLibraryAttributeTitle is called when production libraryAttributeTitle is exited.
func (s *BaseCobol85Listener) ExitLibraryAttributeTitle(ctx *LibraryAttributeTitleContext) {}

// EnterLibraryEntryProcedureClauseFormat1 is called when production libraryEntryProcedureClauseFormat1 is entered.
func (s *BaseCobol85Listener) EnterLibraryEntryProcedureClauseFormat1(ctx *LibraryEntryProcedureClauseFormat1Context) {
}

// ExitLibraryEntryProcedureClauseFormat1 is called when production libraryEntryProcedureClauseFormat1 is exited.
func (s *BaseCobol85Listener) ExitLibraryEntryProcedureClauseFormat1(ctx *LibraryEntryProcedureClauseFormat1Context) {
}

// EnterLibraryEntryProcedureClauseFormat2 is called when production libraryEntryProcedureClauseFormat2 is entered.
func (s *BaseCobol85Listener) EnterLibraryEntryProcedureClauseFormat2(ctx *LibraryEntryProcedureClauseFormat2Context) {
}

// ExitLibraryEntryProcedureClauseFormat2 is called when production libraryEntryProcedureClauseFormat2 is exited.
func (s *BaseCobol85Listener) ExitLibraryEntryProcedureClauseFormat2(ctx *LibraryEntryProcedureClauseFormat2Context) {
}

// EnterLibraryEntryProcedureForClause is called when production libraryEntryProcedureForClause is entered.
func (s *BaseCobol85Listener) EnterLibraryEntryProcedureForClause(ctx *LibraryEntryProcedureForClauseContext) {
}

// ExitLibraryEntryProcedureForClause is called when production libraryEntryProcedureForClause is exited.
func (s *BaseCobol85Listener) ExitLibraryEntryProcedureForClause(ctx *LibraryEntryProcedureForClauseContext) {
}

// EnterLibraryEntryProcedureGivingClause is called when production libraryEntryProcedureGivingClause is entered.
func (s *BaseCobol85Listener) EnterLibraryEntryProcedureGivingClause(ctx *LibraryEntryProcedureGivingClauseContext) {
}

// ExitLibraryEntryProcedureGivingClause is called when production libraryEntryProcedureGivingClause is exited.
func (s *BaseCobol85Listener) ExitLibraryEntryProcedureGivingClause(ctx *LibraryEntryProcedureGivingClauseContext) {
}

// EnterLibraryEntryProcedureUsingClause is called when production libraryEntryProcedureUsingClause is entered.
func (s *BaseCobol85Listener) EnterLibraryEntryProcedureUsingClause(ctx *LibraryEntryProcedureUsingClauseContext) {
}

// ExitLibraryEntryProcedureUsingClause is called when production libraryEntryProcedureUsingClause is exited.
func (s *BaseCobol85Listener) ExitLibraryEntryProcedureUsingClause(ctx *LibraryEntryProcedureUsingClauseContext) {
}

// EnterLibraryEntryProcedureUsingName is called when production libraryEntryProcedureUsingName is entered.
func (s *BaseCobol85Listener) EnterLibraryEntryProcedureUsingName(ctx *LibraryEntryProcedureUsingNameContext) {
}

// ExitLibraryEntryProcedureUsingName is called when production libraryEntryProcedureUsingName is exited.
func (s *BaseCobol85Listener) ExitLibraryEntryProcedureUsingName(ctx *LibraryEntryProcedureUsingNameContext) {
}

// EnterLibraryEntryProcedureWithClause is called when production libraryEntryProcedureWithClause is entered.
func (s *BaseCobol85Listener) EnterLibraryEntryProcedureWithClause(ctx *LibraryEntryProcedureWithClauseContext) {
}

// ExitLibraryEntryProcedureWithClause is called when production libraryEntryProcedureWithClause is exited.
func (s *BaseCobol85Listener) ExitLibraryEntryProcedureWithClause(ctx *LibraryEntryProcedureWithClauseContext) {
}

// EnterLibraryEntryProcedureWithName is called when production libraryEntryProcedureWithName is entered.
func (s *BaseCobol85Listener) EnterLibraryEntryProcedureWithName(ctx *LibraryEntryProcedureWithNameContext) {
}

// ExitLibraryEntryProcedureWithName is called when production libraryEntryProcedureWithName is exited.
func (s *BaseCobol85Listener) ExitLibraryEntryProcedureWithName(ctx *LibraryEntryProcedureWithNameContext) {
}

// EnterLibraryIsCommonClause is called when production libraryIsCommonClause is entered.
func (s *BaseCobol85Listener) EnterLibraryIsCommonClause(ctx *LibraryIsCommonClauseContext) {}

// ExitLibraryIsCommonClause is called when production libraryIsCommonClause is exited.
func (s *BaseCobol85Listener) ExitLibraryIsCommonClause(ctx *LibraryIsCommonClauseContext) {}

// EnterLibraryIsGlobalClause is called when production libraryIsGlobalClause is entered.
func (s *BaseCobol85Listener) EnterLibraryIsGlobalClause(ctx *LibraryIsGlobalClauseContext) {}

// ExitLibraryIsGlobalClause is called when production libraryIsGlobalClause is exited.
func (s *BaseCobol85Listener) ExitLibraryIsGlobalClause(ctx *LibraryIsGlobalClauseContext) {}

// EnterDataDescriptionEntry is called when production dataDescriptionEntry is entered.
func (s *BaseCobol85Listener) EnterDataDescriptionEntry(ctx *DataDescriptionEntryContext) {}

// ExitDataDescriptionEntry is called when production dataDescriptionEntry is exited.
func (s *BaseCobol85Listener) ExitDataDescriptionEntry(ctx *DataDescriptionEntryContext) {}

// EnterDataDescriptionEntryFormat1 is called when production dataDescriptionEntryFormat1 is entered.
func (s *BaseCobol85Listener) EnterDataDescriptionEntryFormat1(ctx *DataDescriptionEntryFormat1Context) {
}

// ExitDataDescriptionEntryFormat1 is called when production dataDescriptionEntryFormat1 is exited.
func (s *BaseCobol85Listener) ExitDataDescriptionEntryFormat1(ctx *DataDescriptionEntryFormat1Context) {
}

// EnterDataDescriptionEntryFormat2 is called when production dataDescriptionEntryFormat2 is entered.
func (s *BaseCobol85Listener) EnterDataDescriptionEntryFormat2(ctx *DataDescriptionEntryFormat2Context) {
}

// ExitDataDescriptionEntryFormat2 is called when production dataDescriptionEntryFormat2 is exited.
func (s *BaseCobol85Listener) ExitDataDescriptionEntryFormat2(ctx *DataDescriptionEntryFormat2Context) {
}

// EnterDataDescriptionEntryFormat3 is called when production dataDescriptionEntryFormat3 is entered.
func (s *BaseCobol85Listener) EnterDataDescriptionEntryFormat3(ctx *DataDescriptionEntryFormat3Context) {
}

// ExitDataDescriptionEntryFormat3 is called when production dataDescriptionEntryFormat3 is exited.
func (s *BaseCobol85Listener) ExitDataDescriptionEntryFormat3(ctx *DataDescriptionEntryFormat3Context) {
}

// EnterDataDescriptionEntryExecSql is called when production dataDescriptionEntryExecSql is entered.
func (s *BaseCobol85Listener) EnterDataDescriptionEntryExecSql(ctx *DataDescriptionEntryExecSqlContext) {
}

// ExitDataDescriptionEntryExecSql is called when production dataDescriptionEntryExecSql is exited.
func (s *BaseCobol85Listener) ExitDataDescriptionEntryExecSql(ctx *DataDescriptionEntryExecSqlContext) {
}

// EnterDataAlignedClause is called when production dataAlignedClause is entered.
func (s *BaseCobol85Listener) EnterDataAlignedClause(ctx *DataAlignedClauseContext) {}

// ExitDataAlignedClause is called when production dataAlignedClause is exited.
func (s *BaseCobol85Listener) ExitDataAlignedClause(ctx *DataAlignedClauseContext) {}

// EnterDataBlankWhenZeroClause is called when production dataBlankWhenZeroClause is entered.
func (s *BaseCobol85Listener) EnterDataBlankWhenZeroClause(ctx *DataBlankWhenZeroClauseContext) {}

// ExitDataBlankWhenZeroClause is called when production dataBlankWhenZeroClause is exited.
func (s *BaseCobol85Listener) ExitDataBlankWhenZeroClause(ctx *DataBlankWhenZeroClauseContext) {}

// EnterDataCommonOwnLocalClause is called when production dataCommonOwnLocalClause is entered.
func (s *BaseCobol85Listener) EnterDataCommonOwnLocalClause(ctx *DataCommonOwnLocalClauseContext) {}

// ExitDataCommonOwnLocalClause is called when production dataCommonOwnLocalClause is exited.
func (s *BaseCobol85Listener) ExitDataCommonOwnLocalClause(ctx *DataCommonOwnLocalClauseContext) {}

// EnterDataExternalClause is called when production dataExternalClause is entered.
func (s *BaseCobol85Listener) EnterDataExternalClause(ctx *DataExternalClauseContext) {}

// ExitDataExternalClause is called when production dataExternalClause is exited.
func (s *BaseCobol85Listener) ExitDataExternalClause(ctx *DataExternalClauseContext) {}

// EnterDataGlobalClause is called when production dataGlobalClause is entered.
func (s *BaseCobol85Listener) EnterDataGlobalClause(ctx *DataGlobalClauseContext) {}

// ExitDataGlobalClause is called when production dataGlobalClause is exited.
func (s *BaseCobol85Listener) ExitDataGlobalClause(ctx *DataGlobalClauseContext) {}

// EnterDataIntegerStringClause is called when production dataIntegerStringClause is entered.
func (s *BaseCobol85Listener) EnterDataIntegerStringClause(ctx *DataIntegerStringClauseContext) {}

// ExitDataIntegerStringClause is called when production dataIntegerStringClause is exited.
func (s *BaseCobol85Listener) ExitDataIntegerStringClause(ctx *DataIntegerStringClauseContext) {}

// EnterDataJustifiedClause is called when production dataJustifiedClause is entered.
func (s *BaseCobol85Listener) EnterDataJustifiedClause(ctx *DataJustifiedClauseContext) {}

// ExitDataJustifiedClause is called when production dataJustifiedClause is exited.
func (s *BaseCobol85Listener) ExitDataJustifiedClause(ctx *DataJustifiedClauseContext) {}

// EnterDataOccursClause is called when production dataOccursClause is entered.
func (s *BaseCobol85Listener) EnterDataOccursClause(ctx *DataOccursClauseContext) {}

// ExitDataOccursClause is called when production dataOccursClause is exited.
func (s *BaseCobol85Listener) ExitDataOccursClause(ctx *DataOccursClauseContext) {}

// EnterDataOccursTo is called when production dataOccursTo is entered.
func (s *BaseCobol85Listener) EnterDataOccursTo(ctx *DataOccursToContext) {}

// ExitDataOccursTo is called when production dataOccursTo is exited.
func (s *BaseCobol85Listener) ExitDataOccursTo(ctx *DataOccursToContext) {}

// EnterDataOccursSort is called when production dataOccursSort is entered.
func (s *BaseCobol85Listener) EnterDataOccursSort(ctx *DataOccursSortContext) {}

// ExitDataOccursSort is called when production dataOccursSort is exited.
func (s *BaseCobol85Listener) ExitDataOccursSort(ctx *DataOccursSortContext) {}

// EnterDataPictureClause is called when production dataPictureClause is entered.
func (s *BaseCobol85Listener) EnterDataPictureClause(ctx *DataPictureClauseContext) {}

// ExitDataPictureClause is called when production dataPictureClause is exited.
func (s *BaseCobol85Listener) ExitDataPictureClause(ctx *DataPictureClauseContext) {}

// EnterPictureString is called when production pictureString is entered.
func (s *BaseCobol85Listener) EnterPictureString(ctx *PictureStringContext) {}

// ExitPictureString is called when production pictureString is exited.
func (s *BaseCobol85Listener) ExitPictureString(ctx *PictureStringContext) {}

// EnterPictureChars is called when production pictureChars is entered.
func (s *BaseCobol85Listener) EnterPictureChars(ctx *PictureCharsContext) {}

// ExitPictureChars is called when production pictureChars is exited.
func (s *BaseCobol85Listener) ExitPictureChars(ctx *PictureCharsContext) {}

// EnterPictureCardinality is called when production pictureCardinality is entered.
func (s *BaseCobol85Listener) EnterPictureCardinality(ctx *PictureCardinalityContext) {}

// ExitPictureCardinality is called when production pictureCardinality is exited.
func (s *BaseCobol85Listener) ExitPictureCardinality(ctx *PictureCardinalityContext) {}

// EnterDataReceivedByClause is called when production dataReceivedByClause is entered.
func (s *BaseCobol85Listener) EnterDataReceivedByClause(ctx *DataReceivedByClauseContext) {}

// ExitDataReceivedByClause is called when production dataReceivedByClause is exited.
func (s *BaseCobol85Listener) ExitDataReceivedByClause(ctx *DataReceivedByClauseContext) {}

// EnterDataRecordAreaClause is called when production dataRecordAreaClause is entered.
func (s *BaseCobol85Listener) EnterDataRecordAreaClause(ctx *DataRecordAreaClauseContext) {}

// ExitDataRecordAreaClause is called when production dataRecordAreaClause is exited.
func (s *BaseCobol85Listener) ExitDataRecordAreaClause(ctx *DataRecordAreaClauseContext) {}

// EnterDataRedefinesClause is called when production dataRedefinesClause is entered.
func (s *BaseCobol85Listener) EnterDataRedefinesClause(ctx *DataRedefinesClauseContext) {}

// ExitDataRedefinesClause is called when production dataRedefinesClause is exited.
func (s *BaseCobol85Listener) ExitDataRedefinesClause(ctx *DataRedefinesClauseContext) {}

// EnterDataRenamesClause is called when production dataRenamesClause is entered.
func (s *BaseCobol85Listener) EnterDataRenamesClause(ctx *DataRenamesClauseContext) {}

// ExitDataRenamesClause is called when production dataRenamesClause is exited.
func (s *BaseCobol85Listener) ExitDataRenamesClause(ctx *DataRenamesClauseContext) {}

// EnterDataSignClause is called when production dataSignClause is entered.
func (s *BaseCobol85Listener) EnterDataSignClause(ctx *DataSignClauseContext) {}

// ExitDataSignClause is called when production dataSignClause is exited.
func (s *BaseCobol85Listener) ExitDataSignClause(ctx *DataSignClauseContext) {}

// EnterDataSynchronizedClause is called when production dataSynchronizedClause is entered.
func (s *BaseCobol85Listener) EnterDataSynchronizedClause(ctx *DataSynchronizedClauseContext) {}

// ExitDataSynchronizedClause is called when production dataSynchronizedClause is exited.
func (s *BaseCobol85Listener) ExitDataSynchronizedClause(ctx *DataSynchronizedClauseContext) {}

// EnterDataThreadLocalClause is called when production dataThreadLocalClause is entered.
func (s *BaseCobol85Listener) EnterDataThreadLocalClause(ctx *DataThreadLocalClauseContext) {}

// ExitDataThreadLocalClause is called when production dataThreadLocalClause is exited.
func (s *BaseCobol85Listener) ExitDataThreadLocalClause(ctx *DataThreadLocalClauseContext) {}

// EnterDataTypeClause is called when production dataTypeClause is entered.
func (s *BaseCobol85Listener) EnterDataTypeClause(ctx *DataTypeClauseContext) {}

// ExitDataTypeClause is called when production dataTypeClause is exited.
func (s *BaseCobol85Listener) ExitDataTypeClause(ctx *DataTypeClauseContext) {}

// EnterDataTypeDefClause is called when production dataTypeDefClause is entered.
func (s *BaseCobol85Listener) EnterDataTypeDefClause(ctx *DataTypeDefClauseContext) {}

// ExitDataTypeDefClause is called when production dataTypeDefClause is exited.
func (s *BaseCobol85Listener) ExitDataTypeDefClause(ctx *DataTypeDefClauseContext) {}

// EnterDataUsageClause is called when production dataUsageClause is entered.
func (s *BaseCobol85Listener) EnterDataUsageClause(ctx *DataUsageClauseContext) {}

// ExitDataUsageClause is called when production dataUsageClause is exited.
func (s *BaseCobol85Listener) ExitDataUsageClause(ctx *DataUsageClauseContext) {}

// EnterDataUsingClause is called when production dataUsingClause is entered.
func (s *BaseCobol85Listener) EnterDataUsingClause(ctx *DataUsingClauseContext) {}

// ExitDataUsingClause is called when production dataUsingClause is exited.
func (s *BaseCobol85Listener) ExitDataUsingClause(ctx *DataUsingClauseContext) {}

// EnterDataValueClause is called when production dataValueClause is entered.
func (s *BaseCobol85Listener) EnterDataValueClause(ctx *DataValueClauseContext) {}

// ExitDataValueClause is called when production dataValueClause is exited.
func (s *BaseCobol85Listener) ExitDataValueClause(ctx *DataValueClauseContext) {}

// EnterDataValueInterval is called when production dataValueInterval is entered.
func (s *BaseCobol85Listener) EnterDataValueInterval(ctx *DataValueIntervalContext) {}

// ExitDataValueInterval is called when production dataValueInterval is exited.
func (s *BaseCobol85Listener) ExitDataValueInterval(ctx *DataValueIntervalContext) {}

// EnterDataValueIntervalFrom is called when production dataValueIntervalFrom is entered.
func (s *BaseCobol85Listener) EnterDataValueIntervalFrom(ctx *DataValueIntervalFromContext) {}

// ExitDataValueIntervalFrom is called when production dataValueIntervalFrom is exited.
func (s *BaseCobol85Listener) ExitDataValueIntervalFrom(ctx *DataValueIntervalFromContext) {}

// EnterDataValueIntervalTo is called when production dataValueIntervalTo is entered.
func (s *BaseCobol85Listener) EnterDataValueIntervalTo(ctx *DataValueIntervalToContext) {}

// ExitDataValueIntervalTo is called when production dataValueIntervalTo is exited.
func (s *BaseCobol85Listener) ExitDataValueIntervalTo(ctx *DataValueIntervalToContext) {}

// EnterDataWithLowerBoundsClause is called when production dataWithLowerBoundsClause is entered.
func (s *BaseCobol85Listener) EnterDataWithLowerBoundsClause(ctx *DataWithLowerBoundsClauseContext) {}

// ExitDataWithLowerBoundsClause is called when production dataWithLowerBoundsClause is exited.
func (s *BaseCobol85Listener) ExitDataWithLowerBoundsClause(ctx *DataWithLowerBoundsClauseContext) {}

// EnterProcedureDivision is called when production procedureDivision is entered.
func (s *BaseCobol85Listener) EnterProcedureDivision(ctx *ProcedureDivisionContext) {}

// ExitProcedureDivision is called when production procedureDivision is exited.
func (s *BaseCobol85Listener) ExitProcedureDivision(ctx *ProcedureDivisionContext) {}

// EnterProcedureDivisionUsingClause is called when production procedureDivisionUsingClause is entered.
func (s *BaseCobol85Listener) EnterProcedureDivisionUsingClause(ctx *ProcedureDivisionUsingClauseContext) {
}

// ExitProcedureDivisionUsingClause is called when production procedureDivisionUsingClause is exited.
func (s *BaseCobol85Listener) ExitProcedureDivisionUsingClause(ctx *ProcedureDivisionUsingClauseContext) {
}

// EnterProcedureDivisionGivingClause is called when production procedureDivisionGivingClause is entered.
func (s *BaseCobol85Listener) EnterProcedureDivisionGivingClause(ctx *ProcedureDivisionGivingClauseContext) {
}

// ExitProcedureDivisionGivingClause is called when production procedureDivisionGivingClause is exited.
func (s *BaseCobol85Listener) ExitProcedureDivisionGivingClause(ctx *ProcedureDivisionGivingClauseContext) {
}

// EnterProcedureDivisionUsingParameter is called when production procedureDivisionUsingParameter is entered.
func (s *BaseCobol85Listener) EnterProcedureDivisionUsingParameter(ctx *ProcedureDivisionUsingParameterContext) {
}

// ExitProcedureDivisionUsingParameter is called when production procedureDivisionUsingParameter is exited.
func (s *BaseCobol85Listener) ExitProcedureDivisionUsingParameter(ctx *ProcedureDivisionUsingParameterContext) {
}

// EnterProcedureDivisionByReferencePhrase is called when production procedureDivisionByReferencePhrase is entered.
func (s *BaseCobol85Listener) EnterProcedureDivisionByReferencePhrase(ctx *ProcedureDivisionByReferencePhraseContext) {
}

// ExitProcedureDivisionByReferencePhrase is called when production procedureDivisionByReferencePhrase is exited.
func (s *BaseCobol85Listener) ExitProcedureDivisionByReferencePhrase(ctx *ProcedureDivisionByReferencePhraseContext) {
}

// EnterProcedureDivisionByReference is called when production procedureDivisionByReference is entered.
func (s *BaseCobol85Listener) EnterProcedureDivisionByReference(ctx *ProcedureDivisionByReferenceContext) {
}

// ExitProcedureDivisionByReference is called when production procedureDivisionByReference is exited.
func (s *BaseCobol85Listener) ExitProcedureDivisionByReference(ctx *ProcedureDivisionByReferenceContext) {
}

// EnterProcedureDivisionByValuePhrase is called when production procedureDivisionByValuePhrase is entered.
func (s *BaseCobol85Listener) EnterProcedureDivisionByValuePhrase(ctx *ProcedureDivisionByValuePhraseContext) {
}

// ExitProcedureDivisionByValuePhrase is called when production procedureDivisionByValuePhrase is exited.
func (s *BaseCobol85Listener) ExitProcedureDivisionByValuePhrase(ctx *ProcedureDivisionByValuePhraseContext) {
}

// EnterProcedureDivisionByValue is called when production procedureDivisionByValue is entered.
func (s *BaseCobol85Listener) EnterProcedureDivisionByValue(ctx *ProcedureDivisionByValueContext) {}

// ExitProcedureDivisionByValue is called when production procedureDivisionByValue is exited.
func (s *BaseCobol85Listener) ExitProcedureDivisionByValue(ctx *ProcedureDivisionByValueContext) {}

// EnterProcedureDeclaratives is called when production procedureDeclaratives is entered.
func (s *BaseCobol85Listener) EnterProcedureDeclaratives(ctx *ProcedureDeclarativesContext) {}

// ExitProcedureDeclaratives is called when production procedureDeclaratives is exited.
func (s *BaseCobol85Listener) ExitProcedureDeclaratives(ctx *ProcedureDeclarativesContext) {}

// EnterProcedureDeclarative is called when production procedureDeclarative is entered.
func (s *BaseCobol85Listener) EnterProcedureDeclarative(ctx *ProcedureDeclarativeContext) {}

// ExitProcedureDeclarative is called when production procedureDeclarative is exited.
func (s *BaseCobol85Listener) ExitProcedureDeclarative(ctx *ProcedureDeclarativeContext) {}

// EnterProcedureSectionHeader is called when production procedureSectionHeader is entered.
func (s *BaseCobol85Listener) EnterProcedureSectionHeader(ctx *ProcedureSectionHeaderContext) {}

// ExitProcedureSectionHeader is called when production procedureSectionHeader is exited.
func (s *BaseCobol85Listener) ExitProcedureSectionHeader(ctx *ProcedureSectionHeaderContext) {}

// EnterProcedureDivisionBody is called when production procedureDivisionBody is entered.
func (s *BaseCobol85Listener) EnterProcedureDivisionBody(ctx *ProcedureDivisionBodyContext) {}

// ExitProcedureDivisionBody is called when production procedureDivisionBody is exited.
func (s *BaseCobol85Listener) ExitProcedureDivisionBody(ctx *ProcedureDivisionBodyContext) {}

// EnterProcedureSection is called when production procedureSection is entered.
func (s *BaseCobol85Listener) EnterProcedureSection(ctx *ProcedureSectionContext) {}

// ExitProcedureSection is called when production procedureSection is exited.
func (s *BaseCobol85Listener) ExitProcedureSection(ctx *ProcedureSectionContext) {}

// EnterParagraphs is called when production paragraphs is entered.
func (s *BaseCobol85Listener) EnterParagraphs(ctx *ParagraphsContext) {}

// ExitParagraphs is called when production paragraphs is exited.
func (s *BaseCobol85Listener) ExitParagraphs(ctx *ParagraphsContext) {}

// EnterParagraph is called when production paragraph is entered.
func (s *BaseCobol85Listener) EnterParagraph(ctx *ParagraphContext) {}

// ExitParagraph is called when production paragraph is exited.
func (s *BaseCobol85Listener) ExitParagraph(ctx *ParagraphContext) {}

// EnterSentence is called when production sentence is entered.
func (s *BaseCobol85Listener) EnterSentence(ctx *SentenceContext) {}

// ExitSentence is called when production sentence is exited.
func (s *BaseCobol85Listener) ExitSentence(ctx *SentenceContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCobol85Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCobol85Listener) ExitStatement(ctx *StatementContext) {}

// EnterAcceptStatement is called when production acceptStatement is entered.
func (s *BaseCobol85Listener) EnterAcceptStatement(ctx *AcceptStatementContext) {}

// ExitAcceptStatement is called when production acceptStatement is exited.
func (s *BaseCobol85Listener) ExitAcceptStatement(ctx *AcceptStatementContext) {}

// EnterAcceptFromDateStatement is called when production acceptFromDateStatement is entered.
func (s *BaseCobol85Listener) EnterAcceptFromDateStatement(ctx *AcceptFromDateStatementContext) {}

// ExitAcceptFromDateStatement is called when production acceptFromDateStatement is exited.
func (s *BaseCobol85Listener) ExitAcceptFromDateStatement(ctx *AcceptFromDateStatementContext) {}

// EnterAcceptFromMnemonicStatement is called when production acceptFromMnemonicStatement is entered.
func (s *BaseCobol85Listener) EnterAcceptFromMnemonicStatement(ctx *AcceptFromMnemonicStatementContext) {
}

// ExitAcceptFromMnemonicStatement is called when production acceptFromMnemonicStatement is exited.
func (s *BaseCobol85Listener) ExitAcceptFromMnemonicStatement(ctx *AcceptFromMnemonicStatementContext) {
}

// EnterAcceptFromEscapeKeyStatement is called when production acceptFromEscapeKeyStatement is entered.
func (s *BaseCobol85Listener) EnterAcceptFromEscapeKeyStatement(ctx *AcceptFromEscapeKeyStatementContext) {
}

// ExitAcceptFromEscapeKeyStatement is called when production acceptFromEscapeKeyStatement is exited.
func (s *BaseCobol85Listener) ExitAcceptFromEscapeKeyStatement(ctx *AcceptFromEscapeKeyStatementContext) {
}

// EnterAcceptMessageCountStatement is called when production acceptMessageCountStatement is entered.
func (s *BaseCobol85Listener) EnterAcceptMessageCountStatement(ctx *AcceptMessageCountStatementContext) {
}

// ExitAcceptMessageCountStatement is called when production acceptMessageCountStatement is exited.
func (s *BaseCobol85Listener) ExitAcceptMessageCountStatement(ctx *AcceptMessageCountStatementContext) {
}

// EnterAddStatement is called when production addStatement is entered.
func (s *BaseCobol85Listener) EnterAddStatement(ctx *AddStatementContext) {}

// ExitAddStatement is called when production addStatement is exited.
func (s *BaseCobol85Listener) ExitAddStatement(ctx *AddStatementContext) {}

// EnterAddToStatement is called when production addToStatement is entered.
func (s *BaseCobol85Listener) EnterAddToStatement(ctx *AddToStatementContext) {}

// ExitAddToStatement is called when production addToStatement is exited.
func (s *BaseCobol85Listener) ExitAddToStatement(ctx *AddToStatementContext) {}

// EnterAddToGivingStatement is called when production addToGivingStatement is entered.
func (s *BaseCobol85Listener) EnterAddToGivingStatement(ctx *AddToGivingStatementContext) {}

// ExitAddToGivingStatement is called when production addToGivingStatement is exited.
func (s *BaseCobol85Listener) ExitAddToGivingStatement(ctx *AddToGivingStatementContext) {}

// EnterAddCorrespondingStatement is called when production addCorrespondingStatement is entered.
func (s *BaseCobol85Listener) EnterAddCorrespondingStatement(ctx *AddCorrespondingStatementContext) {}

// ExitAddCorrespondingStatement is called when production addCorrespondingStatement is exited.
func (s *BaseCobol85Listener) ExitAddCorrespondingStatement(ctx *AddCorrespondingStatementContext) {}

// EnterAddFrom is called when production addFrom is entered.
func (s *BaseCobol85Listener) EnterAddFrom(ctx *AddFromContext) {}

// ExitAddFrom is called when production addFrom is exited.
func (s *BaseCobol85Listener) ExitAddFrom(ctx *AddFromContext) {}

// EnterAddTo is called when production addTo is entered.
func (s *BaseCobol85Listener) EnterAddTo(ctx *AddToContext) {}

// ExitAddTo is called when production addTo is exited.
func (s *BaseCobol85Listener) ExitAddTo(ctx *AddToContext) {}

// EnterAddToGiving is called when production addToGiving is entered.
func (s *BaseCobol85Listener) EnterAddToGiving(ctx *AddToGivingContext) {}

// ExitAddToGiving is called when production addToGiving is exited.
func (s *BaseCobol85Listener) ExitAddToGiving(ctx *AddToGivingContext) {}

// EnterAddGiving is called when production addGiving is entered.
func (s *BaseCobol85Listener) EnterAddGiving(ctx *AddGivingContext) {}

// ExitAddGiving is called when production addGiving is exited.
func (s *BaseCobol85Listener) ExitAddGiving(ctx *AddGivingContext) {}

// EnterAlteredGoTo is called when production alteredGoTo is entered.
func (s *BaseCobol85Listener) EnterAlteredGoTo(ctx *AlteredGoToContext) {}

// ExitAlteredGoTo is called when production alteredGoTo is exited.
func (s *BaseCobol85Listener) ExitAlteredGoTo(ctx *AlteredGoToContext) {}

// EnterAlterStatement is called when production alterStatement is entered.
func (s *BaseCobol85Listener) EnterAlterStatement(ctx *AlterStatementContext) {}

// ExitAlterStatement is called when production alterStatement is exited.
func (s *BaseCobol85Listener) ExitAlterStatement(ctx *AlterStatementContext) {}

// EnterAlterProceedTo is called when production alterProceedTo is entered.
func (s *BaseCobol85Listener) EnterAlterProceedTo(ctx *AlterProceedToContext) {}

// ExitAlterProceedTo is called when production alterProceedTo is exited.
func (s *BaseCobol85Listener) ExitAlterProceedTo(ctx *AlterProceedToContext) {}

// EnterCallStatement is called when production callStatement is entered.
func (s *BaseCobol85Listener) EnterCallStatement(ctx *CallStatementContext) {}

// ExitCallStatement is called when production callStatement is exited.
func (s *BaseCobol85Listener) ExitCallStatement(ctx *CallStatementContext) {}

// EnterCallUsingPhrase is called when production callUsingPhrase is entered.
func (s *BaseCobol85Listener) EnterCallUsingPhrase(ctx *CallUsingPhraseContext) {}

// ExitCallUsingPhrase is called when production callUsingPhrase is exited.
func (s *BaseCobol85Listener) ExitCallUsingPhrase(ctx *CallUsingPhraseContext) {}

// EnterCallUsingParameter is called when production callUsingParameter is entered.
func (s *BaseCobol85Listener) EnterCallUsingParameter(ctx *CallUsingParameterContext) {}

// ExitCallUsingParameter is called when production callUsingParameter is exited.
func (s *BaseCobol85Listener) ExitCallUsingParameter(ctx *CallUsingParameterContext) {}

// EnterCallByReferencePhrase is called when production callByReferencePhrase is entered.
func (s *BaseCobol85Listener) EnterCallByReferencePhrase(ctx *CallByReferencePhraseContext) {}

// ExitCallByReferencePhrase is called when production callByReferencePhrase is exited.
func (s *BaseCobol85Listener) ExitCallByReferencePhrase(ctx *CallByReferencePhraseContext) {}

// EnterCallByReference is called when production callByReference is entered.
func (s *BaseCobol85Listener) EnterCallByReference(ctx *CallByReferenceContext) {}

// ExitCallByReference is called when production callByReference is exited.
func (s *BaseCobol85Listener) ExitCallByReference(ctx *CallByReferenceContext) {}

// EnterCallByValuePhrase is called when production callByValuePhrase is entered.
func (s *BaseCobol85Listener) EnterCallByValuePhrase(ctx *CallByValuePhraseContext) {}

// ExitCallByValuePhrase is called when production callByValuePhrase is exited.
func (s *BaseCobol85Listener) ExitCallByValuePhrase(ctx *CallByValuePhraseContext) {}

// EnterCallByValue is called when production callByValue is entered.
func (s *BaseCobol85Listener) EnterCallByValue(ctx *CallByValueContext) {}

// ExitCallByValue is called when production callByValue is exited.
func (s *BaseCobol85Listener) ExitCallByValue(ctx *CallByValueContext) {}

// EnterCallByContentPhrase is called when production callByContentPhrase is entered.
func (s *BaseCobol85Listener) EnterCallByContentPhrase(ctx *CallByContentPhraseContext) {}

// ExitCallByContentPhrase is called when production callByContentPhrase is exited.
func (s *BaseCobol85Listener) ExitCallByContentPhrase(ctx *CallByContentPhraseContext) {}

// EnterCallByContent is called when production callByContent is entered.
func (s *BaseCobol85Listener) EnterCallByContent(ctx *CallByContentContext) {}

// ExitCallByContent is called when production callByContent is exited.
func (s *BaseCobol85Listener) ExitCallByContent(ctx *CallByContentContext) {}

// EnterCallGivingPhrase is called when production callGivingPhrase is entered.
func (s *BaseCobol85Listener) EnterCallGivingPhrase(ctx *CallGivingPhraseContext) {}

// ExitCallGivingPhrase is called when production callGivingPhrase is exited.
func (s *BaseCobol85Listener) ExitCallGivingPhrase(ctx *CallGivingPhraseContext) {}

// EnterCancelStatement is called when production cancelStatement is entered.
func (s *BaseCobol85Listener) EnterCancelStatement(ctx *CancelStatementContext) {}

// ExitCancelStatement is called when production cancelStatement is exited.
func (s *BaseCobol85Listener) ExitCancelStatement(ctx *CancelStatementContext) {}

// EnterCancelCall is called when production cancelCall is entered.
func (s *BaseCobol85Listener) EnterCancelCall(ctx *CancelCallContext) {}

// ExitCancelCall is called when production cancelCall is exited.
func (s *BaseCobol85Listener) ExitCancelCall(ctx *CancelCallContext) {}

// EnterCloseStatement is called when production closeStatement is entered.
func (s *BaseCobol85Listener) EnterCloseStatement(ctx *CloseStatementContext) {}

// ExitCloseStatement is called when production closeStatement is exited.
func (s *BaseCobol85Listener) ExitCloseStatement(ctx *CloseStatementContext) {}

// EnterCloseFile is called when production closeFile is entered.
func (s *BaseCobol85Listener) EnterCloseFile(ctx *CloseFileContext) {}

// ExitCloseFile is called when production closeFile is exited.
func (s *BaseCobol85Listener) ExitCloseFile(ctx *CloseFileContext) {}

// EnterCloseReelUnitStatement is called when production closeReelUnitStatement is entered.
func (s *BaseCobol85Listener) EnterCloseReelUnitStatement(ctx *CloseReelUnitStatementContext) {}

// ExitCloseReelUnitStatement is called when production closeReelUnitStatement is exited.
func (s *BaseCobol85Listener) ExitCloseReelUnitStatement(ctx *CloseReelUnitStatementContext) {}

// EnterCloseRelativeStatement is called when production closeRelativeStatement is entered.
func (s *BaseCobol85Listener) EnterCloseRelativeStatement(ctx *CloseRelativeStatementContext) {}

// ExitCloseRelativeStatement is called when production closeRelativeStatement is exited.
func (s *BaseCobol85Listener) ExitCloseRelativeStatement(ctx *CloseRelativeStatementContext) {}

// EnterClosePortFileIOStatement is called when production closePortFileIOStatement is entered.
func (s *BaseCobol85Listener) EnterClosePortFileIOStatement(ctx *ClosePortFileIOStatementContext) {}

// ExitClosePortFileIOStatement is called when production closePortFileIOStatement is exited.
func (s *BaseCobol85Listener) ExitClosePortFileIOStatement(ctx *ClosePortFileIOStatementContext) {}

// EnterClosePortFileIOUsing is called when production closePortFileIOUsing is entered.
func (s *BaseCobol85Listener) EnterClosePortFileIOUsing(ctx *ClosePortFileIOUsingContext) {}

// ExitClosePortFileIOUsing is called when production closePortFileIOUsing is exited.
func (s *BaseCobol85Listener) ExitClosePortFileIOUsing(ctx *ClosePortFileIOUsingContext) {}

// EnterClosePortFileIOUsingCloseDisposition is called when production closePortFileIOUsingCloseDisposition is entered.
func (s *BaseCobol85Listener) EnterClosePortFileIOUsingCloseDisposition(ctx *ClosePortFileIOUsingCloseDispositionContext) {
}

// ExitClosePortFileIOUsingCloseDisposition is called when production closePortFileIOUsingCloseDisposition is exited.
func (s *BaseCobol85Listener) ExitClosePortFileIOUsingCloseDisposition(ctx *ClosePortFileIOUsingCloseDispositionContext) {
}

// EnterClosePortFileIOUsingAssociatedData is called when production closePortFileIOUsingAssociatedData is entered.
func (s *BaseCobol85Listener) EnterClosePortFileIOUsingAssociatedData(ctx *ClosePortFileIOUsingAssociatedDataContext) {
}

// ExitClosePortFileIOUsingAssociatedData is called when production closePortFileIOUsingAssociatedData is exited.
func (s *BaseCobol85Listener) ExitClosePortFileIOUsingAssociatedData(ctx *ClosePortFileIOUsingAssociatedDataContext) {
}

// EnterClosePortFileIOUsingAssociatedDataLength is called when production closePortFileIOUsingAssociatedDataLength is entered.
func (s *BaseCobol85Listener) EnterClosePortFileIOUsingAssociatedDataLength(ctx *ClosePortFileIOUsingAssociatedDataLengthContext) {
}

// ExitClosePortFileIOUsingAssociatedDataLength is called when production closePortFileIOUsingAssociatedDataLength is exited.
func (s *BaseCobol85Listener) ExitClosePortFileIOUsingAssociatedDataLength(ctx *ClosePortFileIOUsingAssociatedDataLengthContext) {
}

// EnterComputeStatement is called when production computeStatement is entered.
func (s *BaseCobol85Listener) EnterComputeStatement(ctx *ComputeStatementContext) {}

// ExitComputeStatement is called when production computeStatement is exited.
func (s *BaseCobol85Listener) ExitComputeStatement(ctx *ComputeStatementContext) {}

// EnterComputeStore is called when production computeStore is entered.
func (s *BaseCobol85Listener) EnterComputeStore(ctx *ComputeStoreContext) {}

// ExitComputeStore is called when production computeStore is exited.
func (s *BaseCobol85Listener) ExitComputeStore(ctx *ComputeStoreContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseCobol85Listener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseCobol85Listener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseCobol85Listener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseCobol85Listener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterDisableStatement is called when production disableStatement is entered.
func (s *BaseCobol85Listener) EnterDisableStatement(ctx *DisableStatementContext) {}

// ExitDisableStatement is called when production disableStatement is exited.
func (s *BaseCobol85Listener) ExitDisableStatement(ctx *DisableStatementContext) {}

// EnterDisplayStatement is called when production displayStatement is entered.
func (s *BaseCobol85Listener) EnterDisplayStatement(ctx *DisplayStatementContext) {}

// ExitDisplayStatement is called when production displayStatement is exited.
func (s *BaseCobol85Listener) ExitDisplayStatement(ctx *DisplayStatementContext) {}

// EnterDisplayOperand is called when production displayOperand is entered.
func (s *BaseCobol85Listener) EnterDisplayOperand(ctx *DisplayOperandContext) {}

// ExitDisplayOperand is called when production displayOperand is exited.
func (s *BaseCobol85Listener) ExitDisplayOperand(ctx *DisplayOperandContext) {}

// EnterDisplayAt is called when production displayAt is entered.
func (s *BaseCobol85Listener) EnterDisplayAt(ctx *DisplayAtContext) {}

// ExitDisplayAt is called when production displayAt is exited.
func (s *BaseCobol85Listener) ExitDisplayAt(ctx *DisplayAtContext) {}

// EnterDisplayUpon is called when production displayUpon is entered.
func (s *BaseCobol85Listener) EnterDisplayUpon(ctx *DisplayUponContext) {}

// ExitDisplayUpon is called when production displayUpon is exited.
func (s *BaseCobol85Listener) ExitDisplayUpon(ctx *DisplayUponContext) {}

// EnterDisplayWith is called when production displayWith is entered.
func (s *BaseCobol85Listener) EnterDisplayWith(ctx *DisplayWithContext) {}

// ExitDisplayWith is called when production displayWith is exited.
func (s *BaseCobol85Listener) ExitDisplayWith(ctx *DisplayWithContext) {}

// EnterDivideStatement is called when production divideStatement is entered.
func (s *BaseCobol85Listener) EnterDivideStatement(ctx *DivideStatementContext) {}

// ExitDivideStatement is called when production divideStatement is exited.
func (s *BaseCobol85Listener) ExitDivideStatement(ctx *DivideStatementContext) {}

// EnterDivideIntoStatement is called when production divideIntoStatement is entered.
func (s *BaseCobol85Listener) EnterDivideIntoStatement(ctx *DivideIntoStatementContext) {}

// ExitDivideIntoStatement is called when production divideIntoStatement is exited.
func (s *BaseCobol85Listener) ExitDivideIntoStatement(ctx *DivideIntoStatementContext) {}

// EnterDivideIntoGivingStatement is called when production divideIntoGivingStatement is entered.
func (s *BaseCobol85Listener) EnterDivideIntoGivingStatement(ctx *DivideIntoGivingStatementContext) {}

// ExitDivideIntoGivingStatement is called when production divideIntoGivingStatement is exited.
func (s *BaseCobol85Listener) ExitDivideIntoGivingStatement(ctx *DivideIntoGivingStatementContext) {}

// EnterDivideByGivingStatement is called when production divideByGivingStatement is entered.
func (s *BaseCobol85Listener) EnterDivideByGivingStatement(ctx *DivideByGivingStatementContext) {}

// ExitDivideByGivingStatement is called when production divideByGivingStatement is exited.
func (s *BaseCobol85Listener) ExitDivideByGivingStatement(ctx *DivideByGivingStatementContext) {}

// EnterDivideGivingPhrase is called when production divideGivingPhrase is entered.
func (s *BaseCobol85Listener) EnterDivideGivingPhrase(ctx *DivideGivingPhraseContext) {}

// ExitDivideGivingPhrase is called when production divideGivingPhrase is exited.
func (s *BaseCobol85Listener) ExitDivideGivingPhrase(ctx *DivideGivingPhraseContext) {}

// EnterDivideInto is called when production divideInto is entered.
func (s *BaseCobol85Listener) EnterDivideInto(ctx *DivideIntoContext) {}

// ExitDivideInto is called when production divideInto is exited.
func (s *BaseCobol85Listener) ExitDivideInto(ctx *DivideIntoContext) {}

// EnterDivideGiving is called when production divideGiving is entered.
func (s *BaseCobol85Listener) EnterDivideGiving(ctx *DivideGivingContext) {}

// ExitDivideGiving is called when production divideGiving is exited.
func (s *BaseCobol85Listener) ExitDivideGiving(ctx *DivideGivingContext) {}

// EnterDivideRemainder is called when production divideRemainder is entered.
func (s *BaseCobol85Listener) EnterDivideRemainder(ctx *DivideRemainderContext) {}

// ExitDivideRemainder is called when production divideRemainder is exited.
func (s *BaseCobol85Listener) ExitDivideRemainder(ctx *DivideRemainderContext) {}

// EnterEnableStatement is called when production enableStatement is entered.
func (s *BaseCobol85Listener) EnterEnableStatement(ctx *EnableStatementContext) {}

// ExitEnableStatement is called when production enableStatement is exited.
func (s *BaseCobol85Listener) ExitEnableStatement(ctx *EnableStatementContext) {}

// EnterEntryStatement is called when production entryStatement is entered.
func (s *BaseCobol85Listener) EnterEntryStatement(ctx *EntryStatementContext) {}

// ExitEntryStatement is called when production entryStatement is exited.
func (s *BaseCobol85Listener) ExitEntryStatement(ctx *EntryStatementContext) {}

// EnterEvaluateStatement is called when production evaluateStatement is entered.
func (s *BaseCobol85Listener) EnterEvaluateStatement(ctx *EvaluateStatementContext) {}

// ExitEvaluateStatement is called when production evaluateStatement is exited.
func (s *BaseCobol85Listener) ExitEvaluateStatement(ctx *EvaluateStatementContext) {}

// EnterEvaluateSelect is called when production evaluateSelect is entered.
func (s *BaseCobol85Listener) EnterEvaluateSelect(ctx *EvaluateSelectContext) {}

// ExitEvaluateSelect is called when production evaluateSelect is exited.
func (s *BaseCobol85Listener) ExitEvaluateSelect(ctx *EvaluateSelectContext) {}

// EnterEvaluateAlsoSelect is called when production evaluateAlsoSelect is entered.
func (s *BaseCobol85Listener) EnterEvaluateAlsoSelect(ctx *EvaluateAlsoSelectContext) {}

// ExitEvaluateAlsoSelect is called when production evaluateAlsoSelect is exited.
func (s *BaseCobol85Listener) ExitEvaluateAlsoSelect(ctx *EvaluateAlsoSelectContext) {}

// EnterEvaluateWhenPhrase is called when production evaluateWhenPhrase is entered.
func (s *BaseCobol85Listener) EnterEvaluateWhenPhrase(ctx *EvaluateWhenPhraseContext) {}

// ExitEvaluateWhenPhrase is called when production evaluateWhenPhrase is exited.
func (s *BaseCobol85Listener) ExitEvaluateWhenPhrase(ctx *EvaluateWhenPhraseContext) {}

// EnterEvaluateWhen is called when production evaluateWhen is entered.
func (s *BaseCobol85Listener) EnterEvaluateWhen(ctx *EvaluateWhenContext) {}

// ExitEvaluateWhen is called when production evaluateWhen is exited.
func (s *BaseCobol85Listener) ExitEvaluateWhen(ctx *EvaluateWhenContext) {}

// EnterEvaluateCondition is called when production evaluateCondition is entered.
func (s *BaseCobol85Listener) EnterEvaluateCondition(ctx *EvaluateConditionContext) {}

// ExitEvaluateCondition is called when production evaluateCondition is exited.
func (s *BaseCobol85Listener) ExitEvaluateCondition(ctx *EvaluateConditionContext) {}

// EnterEvaluateThrough is called when production evaluateThrough is entered.
func (s *BaseCobol85Listener) EnterEvaluateThrough(ctx *EvaluateThroughContext) {}

// ExitEvaluateThrough is called when production evaluateThrough is exited.
func (s *BaseCobol85Listener) ExitEvaluateThrough(ctx *EvaluateThroughContext) {}

// EnterEvaluateAlsoCondition is called when production evaluateAlsoCondition is entered.
func (s *BaseCobol85Listener) EnterEvaluateAlsoCondition(ctx *EvaluateAlsoConditionContext) {}

// ExitEvaluateAlsoCondition is called when production evaluateAlsoCondition is exited.
func (s *BaseCobol85Listener) ExitEvaluateAlsoCondition(ctx *EvaluateAlsoConditionContext) {}

// EnterEvaluateWhenOther is called when production evaluateWhenOther is entered.
func (s *BaseCobol85Listener) EnterEvaluateWhenOther(ctx *EvaluateWhenOtherContext) {}

// ExitEvaluateWhenOther is called when production evaluateWhenOther is exited.
func (s *BaseCobol85Listener) ExitEvaluateWhenOther(ctx *EvaluateWhenOtherContext) {}

// EnterEvaluateValue is called when production evaluateValue is entered.
func (s *BaseCobol85Listener) EnterEvaluateValue(ctx *EvaluateValueContext) {}

// ExitEvaluateValue is called when production evaluateValue is exited.
func (s *BaseCobol85Listener) ExitEvaluateValue(ctx *EvaluateValueContext) {}

// EnterExecCicsStatement is called when production execCicsStatement is entered.
func (s *BaseCobol85Listener) EnterExecCicsStatement(ctx *ExecCicsStatementContext) {}

// ExitExecCicsStatement is called when production execCicsStatement is exited.
func (s *BaseCobol85Listener) ExitExecCicsStatement(ctx *ExecCicsStatementContext) {}

// EnterExecSqlStatement is called when production execSqlStatement is entered.
func (s *BaseCobol85Listener) EnterExecSqlStatement(ctx *ExecSqlStatementContext) {}

// ExitExecSqlStatement is called when production execSqlStatement is exited.
func (s *BaseCobol85Listener) ExitExecSqlStatement(ctx *ExecSqlStatementContext) {}

// EnterExecSqlImsStatement is called when production execSqlImsStatement is entered.
func (s *BaseCobol85Listener) EnterExecSqlImsStatement(ctx *ExecSqlImsStatementContext) {}

// ExitExecSqlImsStatement is called when production execSqlImsStatement is exited.
func (s *BaseCobol85Listener) ExitExecSqlImsStatement(ctx *ExecSqlImsStatementContext) {}

// EnterExhibitStatement is called when production exhibitStatement is entered.
func (s *BaseCobol85Listener) EnterExhibitStatement(ctx *ExhibitStatementContext) {}

// ExitExhibitStatement is called when production exhibitStatement is exited.
func (s *BaseCobol85Listener) ExitExhibitStatement(ctx *ExhibitStatementContext) {}

// EnterExhibitOperand is called when production exhibitOperand is entered.
func (s *BaseCobol85Listener) EnterExhibitOperand(ctx *ExhibitOperandContext) {}

// ExitExhibitOperand is called when production exhibitOperand is exited.
func (s *BaseCobol85Listener) ExitExhibitOperand(ctx *ExhibitOperandContext) {}

// EnterExitStatement is called when production exitStatement is entered.
func (s *BaseCobol85Listener) EnterExitStatement(ctx *ExitStatementContext) {}

// ExitExitStatement is called when production exitStatement is exited.
func (s *BaseCobol85Listener) ExitExitStatement(ctx *ExitStatementContext) {}

// EnterGenerateStatement is called when production generateStatement is entered.
func (s *BaseCobol85Listener) EnterGenerateStatement(ctx *GenerateStatementContext) {}

// ExitGenerateStatement is called when production generateStatement is exited.
func (s *BaseCobol85Listener) ExitGenerateStatement(ctx *GenerateStatementContext) {}

// EnterGobackStatement is called when production gobackStatement is entered.
func (s *BaseCobol85Listener) EnterGobackStatement(ctx *GobackStatementContext) {}

// ExitGobackStatement is called when production gobackStatement is exited.
func (s *BaseCobol85Listener) ExitGobackStatement(ctx *GobackStatementContext) {}

// EnterGoToStatement is called when production goToStatement is entered.
func (s *BaseCobol85Listener) EnterGoToStatement(ctx *GoToStatementContext) {}

// ExitGoToStatement is called when production goToStatement is exited.
func (s *BaseCobol85Listener) ExitGoToStatement(ctx *GoToStatementContext) {}

// EnterGoToStatementSimple is called when production goToStatementSimple is entered.
func (s *BaseCobol85Listener) EnterGoToStatementSimple(ctx *GoToStatementSimpleContext) {}

// ExitGoToStatementSimple is called when production goToStatementSimple is exited.
func (s *BaseCobol85Listener) ExitGoToStatementSimple(ctx *GoToStatementSimpleContext) {}

// EnterGoToDependingOnStatement is called when production goToDependingOnStatement is entered.
func (s *BaseCobol85Listener) EnterGoToDependingOnStatement(ctx *GoToDependingOnStatementContext) {}

// ExitGoToDependingOnStatement is called when production goToDependingOnStatement is exited.
func (s *BaseCobol85Listener) ExitGoToDependingOnStatement(ctx *GoToDependingOnStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseCobol85Listener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseCobol85Listener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterIfThen is called when production ifThen is entered.
func (s *BaseCobol85Listener) EnterIfThen(ctx *IfThenContext) {}

// ExitIfThen is called when production ifThen is exited.
func (s *BaseCobol85Listener) ExitIfThen(ctx *IfThenContext) {}

// EnterIfElse is called when production ifElse is entered.
func (s *BaseCobol85Listener) EnterIfElse(ctx *IfElseContext) {}

// ExitIfElse is called when production ifElse is exited.
func (s *BaseCobol85Listener) ExitIfElse(ctx *IfElseContext) {}

// EnterInitializeStatement is called when production initializeStatement is entered.
func (s *BaseCobol85Listener) EnterInitializeStatement(ctx *InitializeStatementContext) {}

// ExitInitializeStatement is called when production initializeStatement is exited.
func (s *BaseCobol85Listener) ExitInitializeStatement(ctx *InitializeStatementContext) {}

// EnterInitializeReplacingPhrase is called when production initializeReplacingPhrase is entered.
func (s *BaseCobol85Listener) EnterInitializeReplacingPhrase(ctx *InitializeReplacingPhraseContext) {}

// ExitInitializeReplacingPhrase is called when production initializeReplacingPhrase is exited.
func (s *BaseCobol85Listener) ExitInitializeReplacingPhrase(ctx *InitializeReplacingPhraseContext) {}

// EnterInitializeReplacingBy is called when production initializeReplacingBy is entered.
func (s *BaseCobol85Listener) EnterInitializeReplacingBy(ctx *InitializeReplacingByContext) {}

// ExitInitializeReplacingBy is called when production initializeReplacingBy is exited.
func (s *BaseCobol85Listener) ExitInitializeReplacingBy(ctx *InitializeReplacingByContext) {}

// EnterInitiateStatement is called when production initiateStatement is entered.
func (s *BaseCobol85Listener) EnterInitiateStatement(ctx *InitiateStatementContext) {}

// ExitInitiateStatement is called when production initiateStatement is exited.
func (s *BaseCobol85Listener) ExitInitiateStatement(ctx *InitiateStatementContext) {}

// EnterInspectStatement is called when production inspectStatement is entered.
func (s *BaseCobol85Listener) EnterInspectStatement(ctx *InspectStatementContext) {}

// ExitInspectStatement is called when production inspectStatement is exited.
func (s *BaseCobol85Listener) ExitInspectStatement(ctx *InspectStatementContext) {}

// EnterInspectTallyingPhrase is called when production inspectTallyingPhrase is entered.
func (s *BaseCobol85Listener) EnterInspectTallyingPhrase(ctx *InspectTallyingPhraseContext) {}

// ExitInspectTallyingPhrase is called when production inspectTallyingPhrase is exited.
func (s *BaseCobol85Listener) ExitInspectTallyingPhrase(ctx *InspectTallyingPhraseContext) {}

// EnterInspectReplacingPhrase is called when production inspectReplacingPhrase is entered.
func (s *BaseCobol85Listener) EnterInspectReplacingPhrase(ctx *InspectReplacingPhraseContext) {}

// ExitInspectReplacingPhrase is called when production inspectReplacingPhrase is exited.
func (s *BaseCobol85Listener) ExitInspectReplacingPhrase(ctx *InspectReplacingPhraseContext) {}

// EnterInspectTallyingReplacingPhrase is called when production inspectTallyingReplacingPhrase is entered.
func (s *BaseCobol85Listener) EnterInspectTallyingReplacingPhrase(ctx *InspectTallyingReplacingPhraseContext) {
}

// ExitInspectTallyingReplacingPhrase is called when production inspectTallyingReplacingPhrase is exited.
func (s *BaseCobol85Listener) ExitInspectTallyingReplacingPhrase(ctx *InspectTallyingReplacingPhraseContext) {
}

// EnterInspectConvertingPhrase is called when production inspectConvertingPhrase is entered.
func (s *BaseCobol85Listener) EnterInspectConvertingPhrase(ctx *InspectConvertingPhraseContext) {}

// ExitInspectConvertingPhrase is called when production inspectConvertingPhrase is exited.
func (s *BaseCobol85Listener) ExitInspectConvertingPhrase(ctx *InspectConvertingPhraseContext) {}

// EnterInspectFor is called when production inspectFor is entered.
func (s *BaseCobol85Listener) EnterInspectFor(ctx *InspectForContext) {}

// ExitInspectFor is called when production inspectFor is exited.
func (s *BaseCobol85Listener) ExitInspectFor(ctx *InspectForContext) {}

// EnterInspectCharacters is called when production inspectCharacters is entered.
func (s *BaseCobol85Listener) EnterInspectCharacters(ctx *InspectCharactersContext) {}

// ExitInspectCharacters is called when production inspectCharacters is exited.
func (s *BaseCobol85Listener) ExitInspectCharacters(ctx *InspectCharactersContext) {}

// EnterInspectReplacingCharacters is called when production inspectReplacingCharacters is entered.
func (s *BaseCobol85Listener) EnterInspectReplacingCharacters(ctx *InspectReplacingCharactersContext) {
}

// ExitInspectReplacingCharacters is called when production inspectReplacingCharacters is exited.
func (s *BaseCobol85Listener) ExitInspectReplacingCharacters(ctx *InspectReplacingCharactersContext) {}

// EnterInspectAllLeadings is called when production inspectAllLeadings is entered.
func (s *BaseCobol85Listener) EnterInspectAllLeadings(ctx *InspectAllLeadingsContext) {}

// ExitInspectAllLeadings is called when production inspectAllLeadings is exited.
func (s *BaseCobol85Listener) ExitInspectAllLeadings(ctx *InspectAllLeadingsContext) {}

// EnterInspectReplacingAllLeadings is called when production inspectReplacingAllLeadings is entered.
func (s *BaseCobol85Listener) EnterInspectReplacingAllLeadings(ctx *InspectReplacingAllLeadingsContext) {
}

// ExitInspectReplacingAllLeadings is called when production inspectReplacingAllLeadings is exited.
func (s *BaseCobol85Listener) ExitInspectReplacingAllLeadings(ctx *InspectReplacingAllLeadingsContext) {
}

// EnterInspectAllLeading is called when production inspectAllLeading is entered.
func (s *BaseCobol85Listener) EnterInspectAllLeading(ctx *InspectAllLeadingContext) {}

// ExitInspectAllLeading is called when production inspectAllLeading is exited.
func (s *BaseCobol85Listener) ExitInspectAllLeading(ctx *InspectAllLeadingContext) {}

// EnterInspectReplacingAllLeading is called when production inspectReplacingAllLeading is entered.
func (s *BaseCobol85Listener) EnterInspectReplacingAllLeading(ctx *InspectReplacingAllLeadingContext) {
}

// ExitInspectReplacingAllLeading is called when production inspectReplacingAllLeading is exited.
func (s *BaseCobol85Listener) ExitInspectReplacingAllLeading(ctx *InspectReplacingAllLeadingContext) {}

// EnterInspectBy is called when production inspectBy is entered.
func (s *BaseCobol85Listener) EnterInspectBy(ctx *InspectByContext) {}

// ExitInspectBy is called when production inspectBy is exited.
func (s *BaseCobol85Listener) ExitInspectBy(ctx *InspectByContext) {}

// EnterInspectTo is called when production inspectTo is entered.
func (s *BaseCobol85Listener) EnterInspectTo(ctx *InspectToContext) {}

// ExitInspectTo is called when production inspectTo is exited.
func (s *BaseCobol85Listener) ExitInspectTo(ctx *InspectToContext) {}

// EnterInspectBeforeAfter is called when production inspectBeforeAfter is entered.
func (s *BaseCobol85Listener) EnterInspectBeforeAfter(ctx *InspectBeforeAfterContext) {}

// ExitInspectBeforeAfter is called when production inspectBeforeAfter is exited.
func (s *BaseCobol85Listener) ExitInspectBeforeAfter(ctx *InspectBeforeAfterContext) {}

// EnterMergeStatement is called when production mergeStatement is entered.
func (s *BaseCobol85Listener) EnterMergeStatement(ctx *MergeStatementContext) {}

// ExitMergeStatement is called when production mergeStatement is exited.
func (s *BaseCobol85Listener) ExitMergeStatement(ctx *MergeStatementContext) {}

// EnterMergeOnKeyClause is called when production mergeOnKeyClause is entered.
func (s *BaseCobol85Listener) EnterMergeOnKeyClause(ctx *MergeOnKeyClauseContext) {}

// ExitMergeOnKeyClause is called when production mergeOnKeyClause is exited.
func (s *BaseCobol85Listener) ExitMergeOnKeyClause(ctx *MergeOnKeyClauseContext) {}

// EnterMergeCollatingSequencePhrase is called when production mergeCollatingSequencePhrase is entered.
func (s *BaseCobol85Listener) EnterMergeCollatingSequencePhrase(ctx *MergeCollatingSequencePhraseContext) {
}

// ExitMergeCollatingSequencePhrase is called when production mergeCollatingSequencePhrase is exited.
func (s *BaseCobol85Listener) ExitMergeCollatingSequencePhrase(ctx *MergeCollatingSequencePhraseContext) {
}

// EnterMergeCollatingAlphanumeric is called when production mergeCollatingAlphanumeric is entered.
func (s *BaseCobol85Listener) EnterMergeCollatingAlphanumeric(ctx *MergeCollatingAlphanumericContext) {
}

// ExitMergeCollatingAlphanumeric is called when production mergeCollatingAlphanumeric is exited.
func (s *BaseCobol85Listener) ExitMergeCollatingAlphanumeric(ctx *MergeCollatingAlphanumericContext) {}

// EnterMergeCollatingNational is called when production mergeCollatingNational is entered.
func (s *BaseCobol85Listener) EnterMergeCollatingNational(ctx *MergeCollatingNationalContext) {}

// ExitMergeCollatingNational is called when production mergeCollatingNational is exited.
func (s *BaseCobol85Listener) ExitMergeCollatingNational(ctx *MergeCollatingNationalContext) {}

// EnterMergeUsing is called when production mergeUsing is entered.
func (s *BaseCobol85Listener) EnterMergeUsing(ctx *MergeUsingContext) {}

// ExitMergeUsing is called when production mergeUsing is exited.
func (s *BaseCobol85Listener) ExitMergeUsing(ctx *MergeUsingContext) {}

// EnterMergeOutputProcedurePhrase is called when production mergeOutputProcedurePhrase is entered.
func (s *BaseCobol85Listener) EnterMergeOutputProcedurePhrase(ctx *MergeOutputProcedurePhraseContext) {
}

// ExitMergeOutputProcedurePhrase is called when production mergeOutputProcedurePhrase is exited.
func (s *BaseCobol85Listener) ExitMergeOutputProcedurePhrase(ctx *MergeOutputProcedurePhraseContext) {}

// EnterMergeOutputThrough is called when production mergeOutputThrough is entered.
func (s *BaseCobol85Listener) EnterMergeOutputThrough(ctx *MergeOutputThroughContext) {}

// ExitMergeOutputThrough is called when production mergeOutputThrough is exited.
func (s *BaseCobol85Listener) ExitMergeOutputThrough(ctx *MergeOutputThroughContext) {}

// EnterMergeGivingPhrase is called when production mergeGivingPhrase is entered.
func (s *BaseCobol85Listener) EnterMergeGivingPhrase(ctx *MergeGivingPhraseContext) {}

// ExitMergeGivingPhrase is called when production mergeGivingPhrase is exited.
func (s *BaseCobol85Listener) ExitMergeGivingPhrase(ctx *MergeGivingPhraseContext) {}

// EnterMergeGiving is called when production mergeGiving is entered.
func (s *BaseCobol85Listener) EnterMergeGiving(ctx *MergeGivingContext) {}

// ExitMergeGiving is called when production mergeGiving is exited.
func (s *BaseCobol85Listener) ExitMergeGiving(ctx *MergeGivingContext) {}

// EnterMoveStatement is called when production moveStatement is entered.
func (s *BaseCobol85Listener) EnterMoveStatement(ctx *MoveStatementContext) {}

// ExitMoveStatement is called when production moveStatement is exited.
func (s *BaseCobol85Listener) ExitMoveStatement(ctx *MoveStatementContext) {}

// EnterMoveToStatement is called when production moveToStatement is entered.
func (s *BaseCobol85Listener) EnterMoveToStatement(ctx *MoveToStatementContext) {}

// ExitMoveToStatement is called when production moveToStatement is exited.
func (s *BaseCobol85Listener) ExitMoveToStatement(ctx *MoveToStatementContext) {}

// EnterMoveToSendingArea is called when production moveToSendingArea is entered.
func (s *BaseCobol85Listener) EnterMoveToSendingArea(ctx *MoveToSendingAreaContext) {}

// ExitMoveToSendingArea is called when production moveToSendingArea is exited.
func (s *BaseCobol85Listener) ExitMoveToSendingArea(ctx *MoveToSendingAreaContext) {}

// EnterMoveCorrespondingToStatement is called when production moveCorrespondingToStatement is entered.
func (s *BaseCobol85Listener) EnterMoveCorrespondingToStatement(ctx *MoveCorrespondingToStatementContext) {
}

// ExitMoveCorrespondingToStatement is called when production moveCorrespondingToStatement is exited.
func (s *BaseCobol85Listener) ExitMoveCorrespondingToStatement(ctx *MoveCorrespondingToStatementContext) {
}

// EnterMoveCorrespondingToSendingArea is called when production moveCorrespondingToSendingArea is entered.
func (s *BaseCobol85Listener) EnterMoveCorrespondingToSendingArea(ctx *MoveCorrespondingToSendingAreaContext) {
}

// ExitMoveCorrespondingToSendingArea is called when production moveCorrespondingToSendingArea is exited.
func (s *BaseCobol85Listener) ExitMoveCorrespondingToSendingArea(ctx *MoveCorrespondingToSendingAreaContext) {
}

// EnterMultiplyStatement is called when production multiplyStatement is entered.
func (s *BaseCobol85Listener) EnterMultiplyStatement(ctx *MultiplyStatementContext) {}

// ExitMultiplyStatement is called when production multiplyStatement is exited.
func (s *BaseCobol85Listener) ExitMultiplyStatement(ctx *MultiplyStatementContext) {}

// EnterMultiplyRegular is called when production multiplyRegular is entered.
func (s *BaseCobol85Listener) EnterMultiplyRegular(ctx *MultiplyRegularContext) {}

// ExitMultiplyRegular is called when production multiplyRegular is exited.
func (s *BaseCobol85Listener) ExitMultiplyRegular(ctx *MultiplyRegularContext) {}

// EnterMultiplyRegularOperand is called when production multiplyRegularOperand is entered.
func (s *BaseCobol85Listener) EnterMultiplyRegularOperand(ctx *MultiplyRegularOperandContext) {}

// ExitMultiplyRegularOperand is called when production multiplyRegularOperand is exited.
func (s *BaseCobol85Listener) ExitMultiplyRegularOperand(ctx *MultiplyRegularOperandContext) {}

// EnterMultiplyGiving is called when production multiplyGiving is entered.
func (s *BaseCobol85Listener) EnterMultiplyGiving(ctx *MultiplyGivingContext) {}

// ExitMultiplyGiving is called when production multiplyGiving is exited.
func (s *BaseCobol85Listener) ExitMultiplyGiving(ctx *MultiplyGivingContext) {}

// EnterMultiplyGivingOperand is called when production multiplyGivingOperand is entered.
func (s *BaseCobol85Listener) EnterMultiplyGivingOperand(ctx *MultiplyGivingOperandContext) {}

// ExitMultiplyGivingOperand is called when production multiplyGivingOperand is exited.
func (s *BaseCobol85Listener) ExitMultiplyGivingOperand(ctx *MultiplyGivingOperandContext) {}

// EnterMultiplyGivingResult is called when production multiplyGivingResult is entered.
func (s *BaseCobol85Listener) EnterMultiplyGivingResult(ctx *MultiplyGivingResultContext) {}

// ExitMultiplyGivingResult is called when production multiplyGivingResult is exited.
func (s *BaseCobol85Listener) ExitMultiplyGivingResult(ctx *MultiplyGivingResultContext) {}

// EnterOpenStatement is called when production openStatement is entered.
func (s *BaseCobol85Listener) EnterOpenStatement(ctx *OpenStatementContext) {}

// ExitOpenStatement is called when production openStatement is exited.
func (s *BaseCobol85Listener) ExitOpenStatement(ctx *OpenStatementContext) {}

// EnterOpenInputStatement is called when production openInputStatement is entered.
func (s *BaseCobol85Listener) EnterOpenInputStatement(ctx *OpenInputStatementContext) {}

// ExitOpenInputStatement is called when production openInputStatement is exited.
func (s *BaseCobol85Listener) ExitOpenInputStatement(ctx *OpenInputStatementContext) {}

// EnterOpenInput is called when production openInput is entered.
func (s *BaseCobol85Listener) EnterOpenInput(ctx *OpenInputContext) {}

// ExitOpenInput is called when production openInput is exited.
func (s *BaseCobol85Listener) ExitOpenInput(ctx *OpenInputContext) {}

// EnterOpenOutputStatement is called when production openOutputStatement is entered.
func (s *BaseCobol85Listener) EnterOpenOutputStatement(ctx *OpenOutputStatementContext) {}

// ExitOpenOutputStatement is called when production openOutputStatement is exited.
func (s *BaseCobol85Listener) ExitOpenOutputStatement(ctx *OpenOutputStatementContext) {}

// EnterOpenOutput is called when production openOutput is entered.
func (s *BaseCobol85Listener) EnterOpenOutput(ctx *OpenOutputContext) {}

// ExitOpenOutput is called when production openOutput is exited.
func (s *BaseCobol85Listener) ExitOpenOutput(ctx *OpenOutputContext) {}

// EnterOpenIOStatement is called when production openIOStatement is entered.
func (s *BaseCobol85Listener) EnterOpenIOStatement(ctx *OpenIOStatementContext) {}

// ExitOpenIOStatement is called when production openIOStatement is exited.
func (s *BaseCobol85Listener) ExitOpenIOStatement(ctx *OpenIOStatementContext) {}

// EnterOpenExtendStatement is called when production openExtendStatement is entered.
func (s *BaseCobol85Listener) EnterOpenExtendStatement(ctx *OpenExtendStatementContext) {}

// ExitOpenExtendStatement is called when production openExtendStatement is exited.
func (s *BaseCobol85Listener) ExitOpenExtendStatement(ctx *OpenExtendStatementContext) {}

// EnterPerformStatement is called when production performStatement is entered.
func (s *BaseCobol85Listener) EnterPerformStatement(ctx *PerformStatementContext) {}

// ExitPerformStatement is called when production performStatement is exited.
func (s *BaseCobol85Listener) ExitPerformStatement(ctx *PerformStatementContext) {}

// EnterPerformInlineStatement is called when production performInlineStatement is entered.
func (s *BaseCobol85Listener) EnterPerformInlineStatement(ctx *PerformInlineStatementContext) {}

// ExitPerformInlineStatement is called when production performInlineStatement is exited.
func (s *BaseCobol85Listener) ExitPerformInlineStatement(ctx *PerformInlineStatementContext) {}

// EnterPerformProcedureStatement is called when production performProcedureStatement is entered.
func (s *BaseCobol85Listener) EnterPerformProcedureStatement(ctx *PerformProcedureStatementContext) {}

// ExitPerformProcedureStatement is called when production performProcedureStatement is exited.
func (s *BaseCobol85Listener) ExitPerformProcedureStatement(ctx *PerformProcedureStatementContext) {}

// EnterPerformType is called when production performType is entered.
func (s *BaseCobol85Listener) EnterPerformType(ctx *PerformTypeContext) {}

// ExitPerformType is called when production performType is exited.
func (s *BaseCobol85Listener) ExitPerformType(ctx *PerformTypeContext) {}

// EnterPerformTimes is called when production performTimes is entered.
func (s *BaseCobol85Listener) EnterPerformTimes(ctx *PerformTimesContext) {}

// ExitPerformTimes is called when production performTimes is exited.
func (s *BaseCobol85Listener) ExitPerformTimes(ctx *PerformTimesContext) {}

// EnterPerformUntil is called when production performUntil is entered.
func (s *BaseCobol85Listener) EnterPerformUntil(ctx *PerformUntilContext) {}

// ExitPerformUntil is called when production performUntil is exited.
func (s *BaseCobol85Listener) ExitPerformUntil(ctx *PerformUntilContext) {}

// EnterPerformVarying is called when production performVarying is entered.
func (s *BaseCobol85Listener) EnterPerformVarying(ctx *PerformVaryingContext) {}

// ExitPerformVarying is called when production performVarying is exited.
func (s *BaseCobol85Listener) ExitPerformVarying(ctx *PerformVaryingContext) {}

// EnterPerformVaryingClause is called when production performVaryingClause is entered.
func (s *BaseCobol85Listener) EnterPerformVaryingClause(ctx *PerformVaryingClauseContext) {}

// ExitPerformVaryingClause is called when production performVaryingClause is exited.
func (s *BaseCobol85Listener) ExitPerformVaryingClause(ctx *PerformVaryingClauseContext) {}

// EnterPerformVaryingPhrase is called when production performVaryingPhrase is entered.
func (s *BaseCobol85Listener) EnterPerformVaryingPhrase(ctx *PerformVaryingPhraseContext) {}

// ExitPerformVaryingPhrase is called when production performVaryingPhrase is exited.
func (s *BaseCobol85Listener) ExitPerformVaryingPhrase(ctx *PerformVaryingPhraseContext) {}

// EnterPerformAfter is called when production performAfter is entered.
func (s *BaseCobol85Listener) EnterPerformAfter(ctx *PerformAfterContext) {}

// ExitPerformAfter is called when production performAfter is exited.
func (s *BaseCobol85Listener) ExitPerformAfter(ctx *PerformAfterContext) {}

// EnterPerformFrom is called when production performFrom is entered.
func (s *BaseCobol85Listener) EnterPerformFrom(ctx *PerformFromContext) {}

// ExitPerformFrom is called when production performFrom is exited.
func (s *BaseCobol85Listener) ExitPerformFrom(ctx *PerformFromContext) {}

// EnterPerformBy is called when production performBy is entered.
func (s *BaseCobol85Listener) EnterPerformBy(ctx *PerformByContext) {}

// ExitPerformBy is called when production performBy is exited.
func (s *BaseCobol85Listener) ExitPerformBy(ctx *PerformByContext) {}

// EnterPerformTestClause is called when production performTestClause is entered.
func (s *BaseCobol85Listener) EnterPerformTestClause(ctx *PerformTestClauseContext) {}

// ExitPerformTestClause is called when production performTestClause is exited.
func (s *BaseCobol85Listener) ExitPerformTestClause(ctx *PerformTestClauseContext) {}

// EnterPurgeStatement is called when production purgeStatement is entered.
func (s *BaseCobol85Listener) EnterPurgeStatement(ctx *PurgeStatementContext) {}

// ExitPurgeStatement is called when production purgeStatement is exited.
func (s *BaseCobol85Listener) ExitPurgeStatement(ctx *PurgeStatementContext) {}

// EnterReadStatement is called when production readStatement is entered.
func (s *BaseCobol85Listener) EnterReadStatement(ctx *ReadStatementContext) {}

// ExitReadStatement is called when production readStatement is exited.
func (s *BaseCobol85Listener) ExitReadStatement(ctx *ReadStatementContext) {}

// EnterReadInto is called when production readInto is entered.
func (s *BaseCobol85Listener) EnterReadInto(ctx *ReadIntoContext) {}

// ExitReadInto is called when production readInto is exited.
func (s *BaseCobol85Listener) ExitReadInto(ctx *ReadIntoContext) {}

// EnterReadWith is called when production readWith is entered.
func (s *BaseCobol85Listener) EnterReadWith(ctx *ReadWithContext) {}

// ExitReadWith is called when production readWith is exited.
func (s *BaseCobol85Listener) ExitReadWith(ctx *ReadWithContext) {}

// EnterReadKey is called when production readKey is entered.
func (s *BaseCobol85Listener) EnterReadKey(ctx *ReadKeyContext) {}

// ExitReadKey is called when production readKey is exited.
func (s *BaseCobol85Listener) ExitReadKey(ctx *ReadKeyContext) {}

// EnterReceiveStatement is called when production receiveStatement is entered.
func (s *BaseCobol85Listener) EnterReceiveStatement(ctx *ReceiveStatementContext) {}

// ExitReceiveStatement is called when production receiveStatement is exited.
func (s *BaseCobol85Listener) ExitReceiveStatement(ctx *ReceiveStatementContext) {}

// EnterReceiveFromStatement is called when production receiveFromStatement is entered.
func (s *BaseCobol85Listener) EnterReceiveFromStatement(ctx *ReceiveFromStatementContext) {}

// ExitReceiveFromStatement is called when production receiveFromStatement is exited.
func (s *BaseCobol85Listener) ExitReceiveFromStatement(ctx *ReceiveFromStatementContext) {}

// EnterReceiveFrom is called when production receiveFrom is entered.
func (s *BaseCobol85Listener) EnterReceiveFrom(ctx *ReceiveFromContext) {}

// ExitReceiveFrom is called when production receiveFrom is exited.
func (s *BaseCobol85Listener) ExitReceiveFrom(ctx *ReceiveFromContext) {}

// EnterReceiveIntoStatement is called when production receiveIntoStatement is entered.
func (s *BaseCobol85Listener) EnterReceiveIntoStatement(ctx *ReceiveIntoStatementContext) {}

// ExitReceiveIntoStatement is called when production receiveIntoStatement is exited.
func (s *BaseCobol85Listener) ExitReceiveIntoStatement(ctx *ReceiveIntoStatementContext) {}

// EnterReceiveNoData is called when production receiveNoData is entered.
func (s *BaseCobol85Listener) EnterReceiveNoData(ctx *ReceiveNoDataContext) {}

// ExitReceiveNoData is called when production receiveNoData is exited.
func (s *BaseCobol85Listener) ExitReceiveNoData(ctx *ReceiveNoDataContext) {}

// EnterReceiveWithData is called when production receiveWithData is entered.
func (s *BaseCobol85Listener) EnterReceiveWithData(ctx *ReceiveWithDataContext) {}

// ExitReceiveWithData is called when production receiveWithData is exited.
func (s *BaseCobol85Listener) ExitReceiveWithData(ctx *ReceiveWithDataContext) {}

// EnterReceiveBefore is called when production receiveBefore is entered.
func (s *BaseCobol85Listener) EnterReceiveBefore(ctx *ReceiveBeforeContext) {}

// ExitReceiveBefore is called when production receiveBefore is exited.
func (s *BaseCobol85Listener) ExitReceiveBefore(ctx *ReceiveBeforeContext) {}

// EnterReceiveWith is called when production receiveWith is entered.
func (s *BaseCobol85Listener) EnterReceiveWith(ctx *ReceiveWithContext) {}

// ExitReceiveWith is called when production receiveWith is exited.
func (s *BaseCobol85Listener) ExitReceiveWith(ctx *ReceiveWithContext) {}

// EnterReceiveThread is called when production receiveThread is entered.
func (s *BaseCobol85Listener) EnterReceiveThread(ctx *ReceiveThreadContext) {}

// ExitReceiveThread is called when production receiveThread is exited.
func (s *BaseCobol85Listener) ExitReceiveThread(ctx *ReceiveThreadContext) {}

// EnterReceiveSize is called when production receiveSize is entered.
func (s *BaseCobol85Listener) EnterReceiveSize(ctx *ReceiveSizeContext) {}

// ExitReceiveSize is called when production receiveSize is exited.
func (s *BaseCobol85Listener) ExitReceiveSize(ctx *ReceiveSizeContext) {}

// EnterReceiveStatus is called when production receiveStatus is entered.
func (s *BaseCobol85Listener) EnterReceiveStatus(ctx *ReceiveStatusContext) {}

// ExitReceiveStatus is called when production receiveStatus is exited.
func (s *BaseCobol85Listener) ExitReceiveStatus(ctx *ReceiveStatusContext) {}

// EnterReleaseStatement is called when production releaseStatement is entered.
func (s *BaseCobol85Listener) EnterReleaseStatement(ctx *ReleaseStatementContext) {}

// ExitReleaseStatement is called when production releaseStatement is exited.
func (s *BaseCobol85Listener) ExitReleaseStatement(ctx *ReleaseStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseCobol85Listener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseCobol85Listener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterReturnInto is called when production returnInto is entered.
func (s *BaseCobol85Listener) EnterReturnInto(ctx *ReturnIntoContext) {}

// ExitReturnInto is called when production returnInto is exited.
func (s *BaseCobol85Listener) ExitReturnInto(ctx *ReturnIntoContext) {}

// EnterRewriteStatement is called when production rewriteStatement is entered.
func (s *BaseCobol85Listener) EnterRewriteStatement(ctx *RewriteStatementContext) {}

// ExitRewriteStatement is called when production rewriteStatement is exited.
func (s *BaseCobol85Listener) ExitRewriteStatement(ctx *RewriteStatementContext) {}

// EnterRewriteFrom is called when production rewriteFrom is entered.
func (s *BaseCobol85Listener) EnterRewriteFrom(ctx *RewriteFromContext) {}

// ExitRewriteFrom is called when production rewriteFrom is exited.
func (s *BaseCobol85Listener) ExitRewriteFrom(ctx *RewriteFromContext) {}

// EnterSearchStatement is called when production searchStatement is entered.
func (s *BaseCobol85Listener) EnterSearchStatement(ctx *SearchStatementContext) {}

// ExitSearchStatement is called when production searchStatement is exited.
func (s *BaseCobol85Listener) ExitSearchStatement(ctx *SearchStatementContext) {}

// EnterSearchVarying is called when production searchVarying is entered.
func (s *BaseCobol85Listener) EnterSearchVarying(ctx *SearchVaryingContext) {}

// ExitSearchVarying is called when production searchVarying is exited.
func (s *BaseCobol85Listener) ExitSearchVarying(ctx *SearchVaryingContext) {}

// EnterSearchWhen is called when production searchWhen is entered.
func (s *BaseCobol85Listener) EnterSearchWhen(ctx *SearchWhenContext) {}

// ExitSearchWhen is called when production searchWhen is exited.
func (s *BaseCobol85Listener) ExitSearchWhen(ctx *SearchWhenContext) {}

// EnterSendStatement is called when production sendStatement is entered.
func (s *BaseCobol85Listener) EnterSendStatement(ctx *SendStatementContext) {}

// ExitSendStatement is called when production sendStatement is exited.
func (s *BaseCobol85Listener) ExitSendStatement(ctx *SendStatementContext) {}

// EnterSendStatementSync is called when production sendStatementSync is entered.
func (s *BaseCobol85Listener) EnterSendStatementSync(ctx *SendStatementSyncContext) {}

// ExitSendStatementSync is called when production sendStatementSync is exited.
func (s *BaseCobol85Listener) ExitSendStatementSync(ctx *SendStatementSyncContext) {}

// EnterSendStatementAsync is called when production sendStatementAsync is entered.
func (s *BaseCobol85Listener) EnterSendStatementAsync(ctx *SendStatementAsyncContext) {}

// ExitSendStatementAsync is called when production sendStatementAsync is exited.
func (s *BaseCobol85Listener) ExitSendStatementAsync(ctx *SendStatementAsyncContext) {}

// EnterSendFromPhrase is called when production sendFromPhrase is entered.
func (s *BaseCobol85Listener) EnterSendFromPhrase(ctx *SendFromPhraseContext) {}

// ExitSendFromPhrase is called when production sendFromPhrase is exited.
func (s *BaseCobol85Listener) ExitSendFromPhrase(ctx *SendFromPhraseContext) {}

// EnterSendWithPhrase is called when production sendWithPhrase is entered.
func (s *BaseCobol85Listener) EnterSendWithPhrase(ctx *SendWithPhraseContext) {}

// ExitSendWithPhrase is called when production sendWithPhrase is exited.
func (s *BaseCobol85Listener) ExitSendWithPhrase(ctx *SendWithPhraseContext) {}

// EnterSendReplacingPhrase is called when production sendReplacingPhrase is entered.
func (s *BaseCobol85Listener) EnterSendReplacingPhrase(ctx *SendReplacingPhraseContext) {}

// ExitSendReplacingPhrase is called when production sendReplacingPhrase is exited.
func (s *BaseCobol85Listener) ExitSendReplacingPhrase(ctx *SendReplacingPhraseContext) {}

// EnterSendAdvancingPhrase is called when production sendAdvancingPhrase is entered.
func (s *BaseCobol85Listener) EnterSendAdvancingPhrase(ctx *SendAdvancingPhraseContext) {}

// ExitSendAdvancingPhrase is called when production sendAdvancingPhrase is exited.
func (s *BaseCobol85Listener) ExitSendAdvancingPhrase(ctx *SendAdvancingPhraseContext) {}

// EnterSendAdvancingPage is called when production sendAdvancingPage is entered.
func (s *BaseCobol85Listener) EnterSendAdvancingPage(ctx *SendAdvancingPageContext) {}

// ExitSendAdvancingPage is called when production sendAdvancingPage is exited.
func (s *BaseCobol85Listener) ExitSendAdvancingPage(ctx *SendAdvancingPageContext) {}

// EnterSendAdvancingLines is called when production sendAdvancingLines is entered.
func (s *BaseCobol85Listener) EnterSendAdvancingLines(ctx *SendAdvancingLinesContext) {}

// ExitSendAdvancingLines is called when production sendAdvancingLines is exited.
func (s *BaseCobol85Listener) ExitSendAdvancingLines(ctx *SendAdvancingLinesContext) {}

// EnterSendAdvancingMnemonic is called when production sendAdvancingMnemonic is entered.
func (s *BaseCobol85Listener) EnterSendAdvancingMnemonic(ctx *SendAdvancingMnemonicContext) {}

// ExitSendAdvancingMnemonic is called when production sendAdvancingMnemonic is exited.
func (s *BaseCobol85Listener) ExitSendAdvancingMnemonic(ctx *SendAdvancingMnemonicContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseCobol85Listener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseCobol85Listener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterSetToStatement is called when production setToStatement is entered.
func (s *BaseCobol85Listener) EnterSetToStatement(ctx *SetToStatementContext) {}

// ExitSetToStatement is called when production setToStatement is exited.
func (s *BaseCobol85Listener) ExitSetToStatement(ctx *SetToStatementContext) {}

// EnterSetUpDownByStatement is called when production setUpDownByStatement is entered.
func (s *BaseCobol85Listener) EnterSetUpDownByStatement(ctx *SetUpDownByStatementContext) {}

// ExitSetUpDownByStatement is called when production setUpDownByStatement is exited.
func (s *BaseCobol85Listener) ExitSetUpDownByStatement(ctx *SetUpDownByStatementContext) {}

// EnterSetTo is called when production setTo is entered.
func (s *BaseCobol85Listener) EnterSetTo(ctx *SetToContext) {}

// ExitSetTo is called when production setTo is exited.
func (s *BaseCobol85Listener) ExitSetTo(ctx *SetToContext) {}

// EnterSetToValue is called when production setToValue is entered.
func (s *BaseCobol85Listener) EnterSetToValue(ctx *SetToValueContext) {}

// ExitSetToValue is called when production setToValue is exited.
func (s *BaseCobol85Listener) ExitSetToValue(ctx *SetToValueContext) {}

// EnterSetByValue is called when production setByValue is entered.
func (s *BaseCobol85Listener) EnterSetByValue(ctx *SetByValueContext) {}

// ExitSetByValue is called when production setByValue is exited.
func (s *BaseCobol85Listener) ExitSetByValue(ctx *SetByValueContext) {}

// EnterSortStatement is called when production sortStatement is entered.
func (s *BaseCobol85Listener) EnterSortStatement(ctx *SortStatementContext) {}

// ExitSortStatement is called when production sortStatement is exited.
func (s *BaseCobol85Listener) ExitSortStatement(ctx *SortStatementContext) {}

// EnterSortOnKeyClause is called when production sortOnKeyClause is entered.
func (s *BaseCobol85Listener) EnterSortOnKeyClause(ctx *SortOnKeyClauseContext) {}

// ExitSortOnKeyClause is called when production sortOnKeyClause is exited.
func (s *BaseCobol85Listener) ExitSortOnKeyClause(ctx *SortOnKeyClauseContext) {}

// EnterSortDuplicatesPhrase is called when production sortDuplicatesPhrase is entered.
func (s *BaseCobol85Listener) EnterSortDuplicatesPhrase(ctx *SortDuplicatesPhraseContext) {}

// ExitSortDuplicatesPhrase is called when production sortDuplicatesPhrase is exited.
func (s *BaseCobol85Listener) ExitSortDuplicatesPhrase(ctx *SortDuplicatesPhraseContext) {}

// EnterSortCollatingSequencePhrase is called when production sortCollatingSequencePhrase is entered.
func (s *BaseCobol85Listener) EnterSortCollatingSequencePhrase(ctx *SortCollatingSequencePhraseContext) {
}

// ExitSortCollatingSequencePhrase is called when production sortCollatingSequencePhrase is exited.
func (s *BaseCobol85Listener) ExitSortCollatingSequencePhrase(ctx *SortCollatingSequencePhraseContext) {
}

// EnterSortCollatingAlphanumeric is called when production sortCollatingAlphanumeric is entered.
func (s *BaseCobol85Listener) EnterSortCollatingAlphanumeric(ctx *SortCollatingAlphanumericContext) {}

// ExitSortCollatingAlphanumeric is called when production sortCollatingAlphanumeric is exited.
func (s *BaseCobol85Listener) ExitSortCollatingAlphanumeric(ctx *SortCollatingAlphanumericContext) {}

// EnterSortCollatingNational is called when production sortCollatingNational is entered.
func (s *BaseCobol85Listener) EnterSortCollatingNational(ctx *SortCollatingNationalContext) {}

// ExitSortCollatingNational is called when production sortCollatingNational is exited.
func (s *BaseCobol85Listener) ExitSortCollatingNational(ctx *SortCollatingNationalContext) {}

// EnterSortInputProcedurePhrase is called when production sortInputProcedurePhrase is entered.
func (s *BaseCobol85Listener) EnterSortInputProcedurePhrase(ctx *SortInputProcedurePhraseContext) {}

// ExitSortInputProcedurePhrase is called when production sortInputProcedurePhrase is exited.
func (s *BaseCobol85Listener) ExitSortInputProcedurePhrase(ctx *SortInputProcedurePhraseContext) {}

// EnterSortInputThrough is called when production sortInputThrough is entered.
func (s *BaseCobol85Listener) EnterSortInputThrough(ctx *SortInputThroughContext) {}

// ExitSortInputThrough is called when production sortInputThrough is exited.
func (s *BaseCobol85Listener) ExitSortInputThrough(ctx *SortInputThroughContext) {}

// EnterSortUsing is called when production sortUsing is entered.
func (s *BaseCobol85Listener) EnterSortUsing(ctx *SortUsingContext) {}

// ExitSortUsing is called when production sortUsing is exited.
func (s *BaseCobol85Listener) ExitSortUsing(ctx *SortUsingContext) {}

// EnterSortOutputProcedurePhrase is called when production sortOutputProcedurePhrase is entered.
func (s *BaseCobol85Listener) EnterSortOutputProcedurePhrase(ctx *SortOutputProcedurePhraseContext) {}

// ExitSortOutputProcedurePhrase is called when production sortOutputProcedurePhrase is exited.
func (s *BaseCobol85Listener) ExitSortOutputProcedurePhrase(ctx *SortOutputProcedurePhraseContext) {}

// EnterSortOutputThrough is called when production sortOutputThrough is entered.
func (s *BaseCobol85Listener) EnterSortOutputThrough(ctx *SortOutputThroughContext) {}

// ExitSortOutputThrough is called when production sortOutputThrough is exited.
func (s *BaseCobol85Listener) ExitSortOutputThrough(ctx *SortOutputThroughContext) {}

// EnterSortGivingPhrase is called when production sortGivingPhrase is entered.
func (s *BaseCobol85Listener) EnterSortGivingPhrase(ctx *SortGivingPhraseContext) {}

// ExitSortGivingPhrase is called when production sortGivingPhrase is exited.
func (s *BaseCobol85Listener) ExitSortGivingPhrase(ctx *SortGivingPhraseContext) {}

// EnterSortGiving is called when production sortGiving is entered.
func (s *BaseCobol85Listener) EnterSortGiving(ctx *SortGivingContext) {}

// ExitSortGiving is called when production sortGiving is exited.
func (s *BaseCobol85Listener) ExitSortGiving(ctx *SortGivingContext) {}

// EnterStartStatement is called when production startStatement is entered.
func (s *BaseCobol85Listener) EnterStartStatement(ctx *StartStatementContext) {}

// ExitStartStatement is called when production startStatement is exited.
func (s *BaseCobol85Listener) ExitStartStatement(ctx *StartStatementContext) {}

// EnterStartKey is called when production startKey is entered.
func (s *BaseCobol85Listener) EnterStartKey(ctx *StartKeyContext) {}

// ExitStartKey is called when production startKey is exited.
func (s *BaseCobol85Listener) ExitStartKey(ctx *StartKeyContext) {}

// EnterStopStatement is called when production stopStatement is entered.
func (s *BaseCobol85Listener) EnterStopStatement(ctx *StopStatementContext) {}

// ExitStopStatement is called when production stopStatement is exited.
func (s *BaseCobol85Listener) ExitStopStatement(ctx *StopStatementContext) {}

// EnterStringStatement is called when production stringStatement is entered.
func (s *BaseCobol85Listener) EnterStringStatement(ctx *StringStatementContext) {}

// ExitStringStatement is called when production stringStatement is exited.
func (s *BaseCobol85Listener) ExitStringStatement(ctx *StringStatementContext) {}

// EnterStringSendingPhrase is called when production stringSendingPhrase is entered.
func (s *BaseCobol85Listener) EnterStringSendingPhrase(ctx *StringSendingPhraseContext) {}

// ExitStringSendingPhrase is called when production stringSendingPhrase is exited.
func (s *BaseCobol85Listener) ExitStringSendingPhrase(ctx *StringSendingPhraseContext) {}

// EnterStringSending is called when production stringSending is entered.
func (s *BaseCobol85Listener) EnterStringSending(ctx *StringSendingContext) {}

// ExitStringSending is called when production stringSending is exited.
func (s *BaseCobol85Listener) ExitStringSending(ctx *StringSendingContext) {}

// EnterStringDelimitedByPhrase is called when production stringDelimitedByPhrase is entered.
func (s *BaseCobol85Listener) EnterStringDelimitedByPhrase(ctx *StringDelimitedByPhraseContext) {}

// ExitStringDelimitedByPhrase is called when production stringDelimitedByPhrase is exited.
func (s *BaseCobol85Listener) ExitStringDelimitedByPhrase(ctx *StringDelimitedByPhraseContext) {}

// EnterStringForPhrase is called when production stringForPhrase is entered.
func (s *BaseCobol85Listener) EnterStringForPhrase(ctx *StringForPhraseContext) {}

// ExitStringForPhrase is called when production stringForPhrase is exited.
func (s *BaseCobol85Listener) ExitStringForPhrase(ctx *StringForPhraseContext) {}

// EnterStringIntoPhrase is called when production stringIntoPhrase is entered.
func (s *BaseCobol85Listener) EnterStringIntoPhrase(ctx *StringIntoPhraseContext) {}

// ExitStringIntoPhrase is called when production stringIntoPhrase is exited.
func (s *BaseCobol85Listener) ExitStringIntoPhrase(ctx *StringIntoPhraseContext) {}

// EnterStringWithPointerPhrase is called when production stringWithPointerPhrase is entered.
func (s *BaseCobol85Listener) EnterStringWithPointerPhrase(ctx *StringWithPointerPhraseContext) {}

// ExitStringWithPointerPhrase is called when production stringWithPointerPhrase is exited.
func (s *BaseCobol85Listener) ExitStringWithPointerPhrase(ctx *StringWithPointerPhraseContext) {}

// EnterSubtractStatement is called when production subtractStatement is entered.
func (s *BaseCobol85Listener) EnterSubtractStatement(ctx *SubtractStatementContext) {}

// ExitSubtractStatement is called when production subtractStatement is exited.
func (s *BaseCobol85Listener) ExitSubtractStatement(ctx *SubtractStatementContext) {}

// EnterSubtractFromStatement is called when production subtractFromStatement is entered.
func (s *BaseCobol85Listener) EnterSubtractFromStatement(ctx *SubtractFromStatementContext) {}

// ExitSubtractFromStatement is called when production subtractFromStatement is exited.
func (s *BaseCobol85Listener) ExitSubtractFromStatement(ctx *SubtractFromStatementContext) {}

// EnterSubtractFromGivingStatement is called when production subtractFromGivingStatement is entered.
func (s *BaseCobol85Listener) EnterSubtractFromGivingStatement(ctx *SubtractFromGivingStatementContext) {
}

// ExitSubtractFromGivingStatement is called when production subtractFromGivingStatement is exited.
func (s *BaseCobol85Listener) ExitSubtractFromGivingStatement(ctx *SubtractFromGivingStatementContext) {
}

// EnterSubtractCorrespondingStatement is called when production subtractCorrespondingStatement is entered.
func (s *BaseCobol85Listener) EnterSubtractCorrespondingStatement(ctx *SubtractCorrespondingStatementContext) {
}

// ExitSubtractCorrespondingStatement is called when production subtractCorrespondingStatement is exited.
func (s *BaseCobol85Listener) ExitSubtractCorrespondingStatement(ctx *SubtractCorrespondingStatementContext) {
}

// EnterSubtractSubtrahend is called when production subtractSubtrahend is entered.
func (s *BaseCobol85Listener) EnterSubtractSubtrahend(ctx *SubtractSubtrahendContext) {}

// ExitSubtractSubtrahend is called when production subtractSubtrahend is exited.
func (s *BaseCobol85Listener) ExitSubtractSubtrahend(ctx *SubtractSubtrahendContext) {}

// EnterSubtractMinuend is called when production subtractMinuend is entered.
func (s *BaseCobol85Listener) EnterSubtractMinuend(ctx *SubtractMinuendContext) {}

// ExitSubtractMinuend is called when production subtractMinuend is exited.
func (s *BaseCobol85Listener) ExitSubtractMinuend(ctx *SubtractMinuendContext) {}

// EnterSubtractMinuendGiving is called when production subtractMinuendGiving is entered.
func (s *BaseCobol85Listener) EnterSubtractMinuendGiving(ctx *SubtractMinuendGivingContext) {}

// ExitSubtractMinuendGiving is called when production subtractMinuendGiving is exited.
func (s *BaseCobol85Listener) ExitSubtractMinuendGiving(ctx *SubtractMinuendGivingContext) {}

// EnterSubtractGiving is called when production subtractGiving is entered.
func (s *BaseCobol85Listener) EnterSubtractGiving(ctx *SubtractGivingContext) {}

// ExitSubtractGiving is called when production subtractGiving is exited.
func (s *BaseCobol85Listener) ExitSubtractGiving(ctx *SubtractGivingContext) {}

// EnterSubtractMinuendCorresponding is called when production subtractMinuendCorresponding is entered.
func (s *BaseCobol85Listener) EnterSubtractMinuendCorresponding(ctx *SubtractMinuendCorrespondingContext) {
}

// ExitSubtractMinuendCorresponding is called when production subtractMinuendCorresponding is exited.
func (s *BaseCobol85Listener) ExitSubtractMinuendCorresponding(ctx *SubtractMinuendCorrespondingContext) {
}

// EnterTerminateStatement is called when production terminateStatement is entered.
func (s *BaseCobol85Listener) EnterTerminateStatement(ctx *TerminateStatementContext) {}

// ExitTerminateStatement is called when production terminateStatement is exited.
func (s *BaseCobol85Listener) ExitTerminateStatement(ctx *TerminateStatementContext) {}

// EnterUnstringStatement is called when production unstringStatement is entered.
func (s *BaseCobol85Listener) EnterUnstringStatement(ctx *UnstringStatementContext) {}

// ExitUnstringStatement is called when production unstringStatement is exited.
func (s *BaseCobol85Listener) ExitUnstringStatement(ctx *UnstringStatementContext) {}

// EnterUnstringSendingPhrase is called when production unstringSendingPhrase is entered.
func (s *BaseCobol85Listener) EnterUnstringSendingPhrase(ctx *UnstringSendingPhraseContext) {}

// ExitUnstringSendingPhrase is called when production unstringSendingPhrase is exited.
func (s *BaseCobol85Listener) ExitUnstringSendingPhrase(ctx *UnstringSendingPhraseContext) {}

// EnterUnstringDelimitedByPhrase is called when production unstringDelimitedByPhrase is entered.
func (s *BaseCobol85Listener) EnterUnstringDelimitedByPhrase(ctx *UnstringDelimitedByPhraseContext) {}

// ExitUnstringDelimitedByPhrase is called when production unstringDelimitedByPhrase is exited.
func (s *BaseCobol85Listener) ExitUnstringDelimitedByPhrase(ctx *UnstringDelimitedByPhraseContext) {}

// EnterUnstringOrAllPhrase is called when production unstringOrAllPhrase is entered.
func (s *BaseCobol85Listener) EnterUnstringOrAllPhrase(ctx *UnstringOrAllPhraseContext) {}

// ExitUnstringOrAllPhrase is called when production unstringOrAllPhrase is exited.
func (s *BaseCobol85Listener) ExitUnstringOrAllPhrase(ctx *UnstringOrAllPhraseContext) {}

// EnterUnstringIntoPhrase is called when production unstringIntoPhrase is entered.
func (s *BaseCobol85Listener) EnterUnstringIntoPhrase(ctx *UnstringIntoPhraseContext) {}

// ExitUnstringIntoPhrase is called when production unstringIntoPhrase is exited.
func (s *BaseCobol85Listener) ExitUnstringIntoPhrase(ctx *UnstringIntoPhraseContext) {}

// EnterUnstringInto is called when production unstringInto is entered.
func (s *BaseCobol85Listener) EnterUnstringInto(ctx *UnstringIntoContext) {}

// ExitUnstringInto is called when production unstringInto is exited.
func (s *BaseCobol85Listener) ExitUnstringInto(ctx *UnstringIntoContext) {}

// EnterUnstringDelimiterIn is called when production unstringDelimiterIn is entered.
func (s *BaseCobol85Listener) EnterUnstringDelimiterIn(ctx *UnstringDelimiterInContext) {}

// ExitUnstringDelimiterIn is called when production unstringDelimiterIn is exited.
func (s *BaseCobol85Listener) ExitUnstringDelimiterIn(ctx *UnstringDelimiterInContext) {}

// EnterUnstringCountIn is called when production unstringCountIn is entered.
func (s *BaseCobol85Listener) EnterUnstringCountIn(ctx *UnstringCountInContext) {}

// ExitUnstringCountIn is called when production unstringCountIn is exited.
func (s *BaseCobol85Listener) ExitUnstringCountIn(ctx *UnstringCountInContext) {}

// EnterUnstringWithPointerPhrase is called when production unstringWithPointerPhrase is entered.
func (s *BaseCobol85Listener) EnterUnstringWithPointerPhrase(ctx *UnstringWithPointerPhraseContext) {}

// ExitUnstringWithPointerPhrase is called when production unstringWithPointerPhrase is exited.
func (s *BaseCobol85Listener) ExitUnstringWithPointerPhrase(ctx *UnstringWithPointerPhraseContext) {}

// EnterUnstringTallyingPhrase is called when production unstringTallyingPhrase is entered.
func (s *BaseCobol85Listener) EnterUnstringTallyingPhrase(ctx *UnstringTallyingPhraseContext) {}

// ExitUnstringTallyingPhrase is called when production unstringTallyingPhrase is exited.
func (s *BaseCobol85Listener) ExitUnstringTallyingPhrase(ctx *UnstringTallyingPhraseContext) {}

// EnterUseStatement is called when production useStatement is entered.
func (s *BaseCobol85Listener) EnterUseStatement(ctx *UseStatementContext) {}

// ExitUseStatement is called when production useStatement is exited.
func (s *BaseCobol85Listener) ExitUseStatement(ctx *UseStatementContext) {}

// EnterUseAfterClause is called when production useAfterClause is entered.
func (s *BaseCobol85Listener) EnterUseAfterClause(ctx *UseAfterClauseContext) {}

// ExitUseAfterClause is called when production useAfterClause is exited.
func (s *BaseCobol85Listener) ExitUseAfterClause(ctx *UseAfterClauseContext) {}

// EnterUseAfterOn is called when production useAfterOn is entered.
func (s *BaseCobol85Listener) EnterUseAfterOn(ctx *UseAfterOnContext) {}

// ExitUseAfterOn is called when production useAfterOn is exited.
func (s *BaseCobol85Listener) ExitUseAfterOn(ctx *UseAfterOnContext) {}

// EnterUseDebugClause is called when production useDebugClause is entered.
func (s *BaseCobol85Listener) EnterUseDebugClause(ctx *UseDebugClauseContext) {}

// ExitUseDebugClause is called when production useDebugClause is exited.
func (s *BaseCobol85Listener) ExitUseDebugClause(ctx *UseDebugClauseContext) {}

// EnterUseDebugOn is called when production useDebugOn is entered.
func (s *BaseCobol85Listener) EnterUseDebugOn(ctx *UseDebugOnContext) {}

// ExitUseDebugOn is called when production useDebugOn is exited.
func (s *BaseCobol85Listener) ExitUseDebugOn(ctx *UseDebugOnContext) {}

// EnterWriteStatement is called when production writeStatement is entered.
func (s *BaseCobol85Listener) EnterWriteStatement(ctx *WriteStatementContext) {}

// ExitWriteStatement is called when production writeStatement is exited.
func (s *BaseCobol85Listener) ExitWriteStatement(ctx *WriteStatementContext) {}

// EnterWriteFromPhrase is called when production writeFromPhrase is entered.
func (s *BaseCobol85Listener) EnterWriteFromPhrase(ctx *WriteFromPhraseContext) {}

// ExitWriteFromPhrase is called when production writeFromPhrase is exited.
func (s *BaseCobol85Listener) ExitWriteFromPhrase(ctx *WriteFromPhraseContext) {}

// EnterWriteAdvancingPhrase is called when production writeAdvancingPhrase is entered.
func (s *BaseCobol85Listener) EnterWriteAdvancingPhrase(ctx *WriteAdvancingPhraseContext) {}

// ExitWriteAdvancingPhrase is called when production writeAdvancingPhrase is exited.
func (s *BaseCobol85Listener) ExitWriteAdvancingPhrase(ctx *WriteAdvancingPhraseContext) {}

// EnterWriteAdvancingPage is called when production writeAdvancingPage is entered.
func (s *BaseCobol85Listener) EnterWriteAdvancingPage(ctx *WriteAdvancingPageContext) {}

// ExitWriteAdvancingPage is called when production writeAdvancingPage is exited.
func (s *BaseCobol85Listener) ExitWriteAdvancingPage(ctx *WriteAdvancingPageContext) {}

// EnterWriteAdvancingLines is called when production writeAdvancingLines is entered.
func (s *BaseCobol85Listener) EnterWriteAdvancingLines(ctx *WriteAdvancingLinesContext) {}

// ExitWriteAdvancingLines is called when production writeAdvancingLines is exited.
func (s *BaseCobol85Listener) ExitWriteAdvancingLines(ctx *WriteAdvancingLinesContext) {}

// EnterWriteAdvancingMnemonic is called when production writeAdvancingMnemonic is entered.
func (s *BaseCobol85Listener) EnterWriteAdvancingMnemonic(ctx *WriteAdvancingMnemonicContext) {}

// ExitWriteAdvancingMnemonic is called when production writeAdvancingMnemonic is exited.
func (s *BaseCobol85Listener) ExitWriteAdvancingMnemonic(ctx *WriteAdvancingMnemonicContext) {}

// EnterWriteAtEndOfPagePhrase is called when production writeAtEndOfPagePhrase is entered.
func (s *BaseCobol85Listener) EnterWriteAtEndOfPagePhrase(ctx *WriteAtEndOfPagePhraseContext) {}

// ExitWriteAtEndOfPagePhrase is called when production writeAtEndOfPagePhrase is exited.
func (s *BaseCobol85Listener) ExitWriteAtEndOfPagePhrase(ctx *WriteAtEndOfPagePhraseContext) {}

// EnterWriteNotAtEndOfPagePhrase is called when production writeNotAtEndOfPagePhrase is entered.
func (s *BaseCobol85Listener) EnterWriteNotAtEndOfPagePhrase(ctx *WriteNotAtEndOfPagePhraseContext) {}

// ExitWriteNotAtEndOfPagePhrase is called when production writeNotAtEndOfPagePhrase is exited.
func (s *BaseCobol85Listener) ExitWriteNotAtEndOfPagePhrase(ctx *WriteNotAtEndOfPagePhraseContext) {}

// EnterAtEndPhrase is called when production atEndPhrase is entered.
func (s *BaseCobol85Listener) EnterAtEndPhrase(ctx *AtEndPhraseContext) {}

// ExitAtEndPhrase is called when production atEndPhrase is exited.
func (s *BaseCobol85Listener) ExitAtEndPhrase(ctx *AtEndPhraseContext) {}

// EnterNotAtEndPhrase is called when production notAtEndPhrase is entered.
func (s *BaseCobol85Listener) EnterNotAtEndPhrase(ctx *NotAtEndPhraseContext) {}

// ExitNotAtEndPhrase is called when production notAtEndPhrase is exited.
func (s *BaseCobol85Listener) ExitNotAtEndPhrase(ctx *NotAtEndPhraseContext) {}

// EnterInvalidKeyPhrase is called when production invalidKeyPhrase is entered.
func (s *BaseCobol85Listener) EnterInvalidKeyPhrase(ctx *InvalidKeyPhraseContext) {}

// ExitInvalidKeyPhrase is called when production invalidKeyPhrase is exited.
func (s *BaseCobol85Listener) ExitInvalidKeyPhrase(ctx *InvalidKeyPhraseContext) {}

// EnterNotInvalidKeyPhrase is called when production notInvalidKeyPhrase is entered.
func (s *BaseCobol85Listener) EnterNotInvalidKeyPhrase(ctx *NotInvalidKeyPhraseContext) {}

// ExitNotInvalidKeyPhrase is called when production notInvalidKeyPhrase is exited.
func (s *BaseCobol85Listener) ExitNotInvalidKeyPhrase(ctx *NotInvalidKeyPhraseContext) {}

// EnterOnOverflowPhrase is called when production onOverflowPhrase is entered.
func (s *BaseCobol85Listener) EnterOnOverflowPhrase(ctx *OnOverflowPhraseContext) {}

// ExitOnOverflowPhrase is called when production onOverflowPhrase is exited.
func (s *BaseCobol85Listener) ExitOnOverflowPhrase(ctx *OnOverflowPhraseContext) {}

// EnterNotOnOverflowPhrase is called when production notOnOverflowPhrase is entered.
func (s *BaseCobol85Listener) EnterNotOnOverflowPhrase(ctx *NotOnOverflowPhraseContext) {}

// ExitNotOnOverflowPhrase is called when production notOnOverflowPhrase is exited.
func (s *BaseCobol85Listener) ExitNotOnOverflowPhrase(ctx *NotOnOverflowPhraseContext) {}

// EnterOnSizeErrorPhrase is called when production onSizeErrorPhrase is entered.
func (s *BaseCobol85Listener) EnterOnSizeErrorPhrase(ctx *OnSizeErrorPhraseContext) {}

// ExitOnSizeErrorPhrase is called when production onSizeErrorPhrase is exited.
func (s *BaseCobol85Listener) ExitOnSizeErrorPhrase(ctx *OnSizeErrorPhraseContext) {}

// EnterNotOnSizeErrorPhrase is called when production notOnSizeErrorPhrase is entered.
func (s *BaseCobol85Listener) EnterNotOnSizeErrorPhrase(ctx *NotOnSizeErrorPhraseContext) {}

// ExitNotOnSizeErrorPhrase is called when production notOnSizeErrorPhrase is exited.
func (s *BaseCobol85Listener) ExitNotOnSizeErrorPhrase(ctx *NotOnSizeErrorPhraseContext) {}

// EnterOnExceptionClause is called when production onExceptionClause is entered.
func (s *BaseCobol85Listener) EnterOnExceptionClause(ctx *OnExceptionClauseContext) {}

// ExitOnExceptionClause is called when production onExceptionClause is exited.
func (s *BaseCobol85Listener) ExitOnExceptionClause(ctx *OnExceptionClauseContext) {}

// EnterNotOnExceptionClause is called when production notOnExceptionClause is entered.
func (s *BaseCobol85Listener) EnterNotOnExceptionClause(ctx *NotOnExceptionClauseContext) {}

// ExitNotOnExceptionClause is called when production notOnExceptionClause is exited.
func (s *BaseCobol85Listener) ExitNotOnExceptionClause(ctx *NotOnExceptionClauseContext) {}

// EnterArithmeticExpression is called when production arithmeticExpression is entered.
func (s *BaseCobol85Listener) EnterArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// ExitArithmeticExpression is called when production arithmeticExpression is exited.
func (s *BaseCobol85Listener) ExitArithmeticExpression(ctx *ArithmeticExpressionContext) {}

// EnterPlusMinus is called when production plusMinus is entered.
func (s *BaseCobol85Listener) EnterPlusMinus(ctx *PlusMinusContext) {}

// ExitPlusMinus is called when production plusMinus is exited.
func (s *BaseCobol85Listener) ExitPlusMinus(ctx *PlusMinusContext) {}

// EnterMultDivs is called when production multDivs is entered.
func (s *BaseCobol85Listener) EnterMultDivs(ctx *MultDivsContext) {}

// ExitMultDivs is called when production multDivs is exited.
func (s *BaseCobol85Listener) ExitMultDivs(ctx *MultDivsContext) {}

// EnterMultDiv is called when production multDiv is entered.
func (s *BaseCobol85Listener) EnterMultDiv(ctx *MultDivContext) {}

// ExitMultDiv is called when production multDiv is exited.
func (s *BaseCobol85Listener) ExitMultDiv(ctx *MultDivContext) {}

// EnterPowers is called when production powers is entered.
func (s *BaseCobol85Listener) EnterPowers(ctx *PowersContext) {}

// ExitPowers is called when production powers is exited.
func (s *BaseCobol85Listener) ExitPowers(ctx *PowersContext) {}

// EnterPower is called when production power is entered.
func (s *BaseCobol85Listener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production power is exited.
func (s *BaseCobol85Listener) ExitPower(ctx *PowerContext) {}

// EnterBasis is called when production basis is entered.
func (s *BaseCobol85Listener) EnterBasis(ctx *BasisContext) {}

// ExitBasis is called when production basis is exited.
func (s *BaseCobol85Listener) ExitBasis(ctx *BasisContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseCobol85Listener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseCobol85Listener) ExitCondition(ctx *ConditionContext) {}

// EnterAndOrCondition is called when production andOrCondition is entered.
func (s *BaseCobol85Listener) EnterAndOrCondition(ctx *AndOrConditionContext) {}

// ExitAndOrCondition is called when production andOrCondition is exited.
func (s *BaseCobol85Listener) ExitAndOrCondition(ctx *AndOrConditionContext) {}

// EnterCombinableCondition is called when production combinableCondition is entered.
func (s *BaseCobol85Listener) EnterCombinableCondition(ctx *CombinableConditionContext) {}

// ExitCombinableCondition is called when production combinableCondition is exited.
func (s *BaseCobol85Listener) ExitCombinableCondition(ctx *CombinableConditionContext) {}

// EnterSimpleCondition is called when production simpleCondition is entered.
func (s *BaseCobol85Listener) EnterSimpleCondition(ctx *SimpleConditionContext) {}

// ExitSimpleCondition is called when production simpleCondition is exited.
func (s *BaseCobol85Listener) ExitSimpleCondition(ctx *SimpleConditionContext) {}

// EnterClassCondition is called when production classCondition is entered.
func (s *BaseCobol85Listener) EnterClassCondition(ctx *ClassConditionContext) {}

// ExitClassCondition is called when production classCondition is exited.
func (s *BaseCobol85Listener) ExitClassCondition(ctx *ClassConditionContext) {}

// EnterConditionNameReference is called when production conditionNameReference is entered.
func (s *BaseCobol85Listener) EnterConditionNameReference(ctx *ConditionNameReferenceContext) {}

// ExitConditionNameReference is called when production conditionNameReference is exited.
func (s *BaseCobol85Listener) ExitConditionNameReference(ctx *ConditionNameReferenceContext) {}

// EnterConditionNameSubscriptReference is called when production conditionNameSubscriptReference is entered.
func (s *BaseCobol85Listener) EnterConditionNameSubscriptReference(ctx *ConditionNameSubscriptReferenceContext) {
}

// ExitConditionNameSubscriptReference is called when production conditionNameSubscriptReference is exited.
func (s *BaseCobol85Listener) ExitConditionNameSubscriptReference(ctx *ConditionNameSubscriptReferenceContext) {
}

// EnterRelationCondition is called when production relationCondition is entered.
func (s *BaseCobol85Listener) EnterRelationCondition(ctx *RelationConditionContext) {}

// ExitRelationCondition is called when production relationCondition is exited.
func (s *BaseCobol85Listener) ExitRelationCondition(ctx *RelationConditionContext) {}

// EnterRelationSignCondition is called when production relationSignCondition is entered.
func (s *BaseCobol85Listener) EnterRelationSignCondition(ctx *RelationSignConditionContext) {}

// ExitRelationSignCondition is called when production relationSignCondition is exited.
func (s *BaseCobol85Listener) ExitRelationSignCondition(ctx *RelationSignConditionContext) {}

// EnterRelationArithmeticComparison is called when production relationArithmeticComparison is entered.
func (s *BaseCobol85Listener) EnterRelationArithmeticComparison(ctx *RelationArithmeticComparisonContext) {
}

// ExitRelationArithmeticComparison is called when production relationArithmeticComparison is exited.
func (s *BaseCobol85Listener) ExitRelationArithmeticComparison(ctx *RelationArithmeticComparisonContext) {
}

// EnterRelationCombinedComparison is called when production relationCombinedComparison is entered.
func (s *BaseCobol85Listener) EnterRelationCombinedComparison(ctx *RelationCombinedComparisonContext) {
}

// ExitRelationCombinedComparison is called when production relationCombinedComparison is exited.
func (s *BaseCobol85Listener) ExitRelationCombinedComparison(ctx *RelationCombinedComparisonContext) {}

// EnterRelationCombinedCondition is called when production relationCombinedCondition is entered.
func (s *BaseCobol85Listener) EnterRelationCombinedCondition(ctx *RelationCombinedConditionContext) {}

// ExitRelationCombinedCondition is called when production relationCombinedCondition is exited.
func (s *BaseCobol85Listener) ExitRelationCombinedCondition(ctx *RelationCombinedConditionContext) {}

// EnterRelationalOperator is called when production relationalOperator is entered.
func (s *BaseCobol85Listener) EnterRelationalOperator(ctx *RelationalOperatorContext) {}

// ExitRelationalOperator is called when production relationalOperator is exited.
func (s *BaseCobol85Listener) ExitRelationalOperator(ctx *RelationalOperatorContext) {}

// EnterAbbreviation is called when production abbreviation is entered.
func (s *BaseCobol85Listener) EnterAbbreviation(ctx *AbbreviationContext) {}

// ExitAbbreviation is called when production abbreviation is exited.
func (s *BaseCobol85Listener) ExitAbbreviation(ctx *AbbreviationContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseCobol85Listener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseCobol85Listener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterTableCall is called when production tableCall is entered.
func (s *BaseCobol85Listener) EnterTableCall(ctx *TableCallContext) {}

// ExitTableCall is called when production tableCall is exited.
func (s *BaseCobol85Listener) ExitTableCall(ctx *TableCallContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseCobol85Listener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseCobol85Listener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterReferenceModifier is called when production referenceModifier is entered.
func (s *BaseCobol85Listener) EnterReferenceModifier(ctx *ReferenceModifierContext) {}

// ExitReferenceModifier is called when production referenceModifier is exited.
func (s *BaseCobol85Listener) ExitReferenceModifier(ctx *ReferenceModifierContext) {}

// EnterCharacterPosition is called when production characterPosition is entered.
func (s *BaseCobol85Listener) EnterCharacterPosition(ctx *CharacterPositionContext) {}

// ExitCharacterPosition is called when production characterPosition is exited.
func (s *BaseCobol85Listener) ExitCharacterPosition(ctx *CharacterPositionContext) {}

// EnterLength is called when production length is entered.
func (s *BaseCobol85Listener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BaseCobol85Listener) ExitLength(ctx *LengthContext) {}

// EnterSubscript is called when production subscript is entered.
func (s *BaseCobol85Listener) EnterSubscript(ctx *SubscriptContext) {}

// ExitSubscript is called when production subscript is exited.
func (s *BaseCobol85Listener) ExitSubscript(ctx *SubscriptContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseCobol85Listener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseCobol85Listener) ExitArgument(ctx *ArgumentContext) {}

// EnterQualifiedDataName is called when production qualifiedDataName is entered.
func (s *BaseCobol85Listener) EnterQualifiedDataName(ctx *QualifiedDataNameContext) {}

// ExitQualifiedDataName is called when production qualifiedDataName is exited.
func (s *BaseCobol85Listener) ExitQualifiedDataName(ctx *QualifiedDataNameContext) {}

// EnterQualifiedDataNameFormat1 is called when production qualifiedDataNameFormat1 is entered.
func (s *BaseCobol85Listener) EnterQualifiedDataNameFormat1(ctx *QualifiedDataNameFormat1Context) {}

// ExitQualifiedDataNameFormat1 is called when production qualifiedDataNameFormat1 is exited.
func (s *BaseCobol85Listener) ExitQualifiedDataNameFormat1(ctx *QualifiedDataNameFormat1Context) {}

// EnterQualifiedDataNameFormat2 is called when production qualifiedDataNameFormat2 is entered.
func (s *BaseCobol85Listener) EnterQualifiedDataNameFormat2(ctx *QualifiedDataNameFormat2Context) {}

// ExitQualifiedDataNameFormat2 is called when production qualifiedDataNameFormat2 is exited.
func (s *BaseCobol85Listener) ExitQualifiedDataNameFormat2(ctx *QualifiedDataNameFormat2Context) {}

// EnterQualifiedDataNameFormat3 is called when production qualifiedDataNameFormat3 is entered.
func (s *BaseCobol85Listener) EnterQualifiedDataNameFormat3(ctx *QualifiedDataNameFormat3Context) {}

// ExitQualifiedDataNameFormat3 is called when production qualifiedDataNameFormat3 is exited.
func (s *BaseCobol85Listener) ExitQualifiedDataNameFormat3(ctx *QualifiedDataNameFormat3Context) {}

// EnterQualifiedDataNameFormat4 is called when production qualifiedDataNameFormat4 is entered.
func (s *BaseCobol85Listener) EnterQualifiedDataNameFormat4(ctx *QualifiedDataNameFormat4Context) {}

// ExitQualifiedDataNameFormat4 is called when production qualifiedDataNameFormat4 is exited.
func (s *BaseCobol85Listener) ExitQualifiedDataNameFormat4(ctx *QualifiedDataNameFormat4Context) {}

// EnterQualifiedInData is called when production qualifiedInData is entered.
func (s *BaseCobol85Listener) EnterQualifiedInData(ctx *QualifiedInDataContext) {}

// ExitQualifiedInData is called when production qualifiedInData is exited.
func (s *BaseCobol85Listener) ExitQualifiedInData(ctx *QualifiedInDataContext) {}

// EnterInData is called when production inData is entered.
func (s *BaseCobol85Listener) EnterInData(ctx *InDataContext) {}

// ExitInData is called when production inData is exited.
func (s *BaseCobol85Listener) ExitInData(ctx *InDataContext) {}

// EnterInFile is called when production inFile is entered.
func (s *BaseCobol85Listener) EnterInFile(ctx *InFileContext) {}

// ExitInFile is called when production inFile is exited.
func (s *BaseCobol85Listener) ExitInFile(ctx *InFileContext) {}

// EnterInMnemonic is called when production inMnemonic is entered.
func (s *BaseCobol85Listener) EnterInMnemonic(ctx *InMnemonicContext) {}

// ExitInMnemonic is called when production inMnemonic is exited.
func (s *BaseCobol85Listener) ExitInMnemonic(ctx *InMnemonicContext) {}

// EnterInSection is called when production inSection is entered.
func (s *BaseCobol85Listener) EnterInSection(ctx *InSectionContext) {}

// ExitInSection is called when production inSection is exited.
func (s *BaseCobol85Listener) ExitInSection(ctx *InSectionContext) {}

// EnterInLibrary is called when production inLibrary is entered.
func (s *BaseCobol85Listener) EnterInLibrary(ctx *InLibraryContext) {}

// ExitInLibrary is called when production inLibrary is exited.
func (s *BaseCobol85Listener) ExitInLibrary(ctx *InLibraryContext) {}

// EnterInTable is called when production inTable is entered.
func (s *BaseCobol85Listener) EnterInTable(ctx *InTableContext) {}

// ExitInTable is called when production inTable is exited.
func (s *BaseCobol85Listener) ExitInTable(ctx *InTableContext) {}

// EnterAlphabetName is called when production alphabetName is entered.
func (s *BaseCobol85Listener) EnterAlphabetName(ctx *AlphabetNameContext) {}

// ExitAlphabetName is called when production alphabetName is exited.
func (s *BaseCobol85Listener) ExitAlphabetName(ctx *AlphabetNameContext) {}

// EnterAssignmentName is called when production assignmentName is entered.
func (s *BaseCobol85Listener) EnterAssignmentName(ctx *AssignmentNameContext) {}

// ExitAssignmentName is called when production assignmentName is exited.
func (s *BaseCobol85Listener) ExitAssignmentName(ctx *AssignmentNameContext) {}

// EnterBasisName is called when production basisName is entered.
func (s *BaseCobol85Listener) EnterBasisName(ctx *BasisNameContext) {}

// ExitBasisName is called when production basisName is exited.
func (s *BaseCobol85Listener) ExitBasisName(ctx *BasisNameContext) {}

// EnterCdName is called when production cdName is entered.
func (s *BaseCobol85Listener) EnterCdName(ctx *CdNameContext) {}

// ExitCdName is called when production cdName is exited.
func (s *BaseCobol85Listener) ExitCdName(ctx *CdNameContext) {}

// EnterClassName is called when production className is entered.
func (s *BaseCobol85Listener) EnterClassName(ctx *ClassNameContext) {}

// ExitClassName is called when production className is exited.
func (s *BaseCobol85Listener) ExitClassName(ctx *ClassNameContext) {}

// EnterComputerName is called when production computerName is entered.
func (s *BaseCobol85Listener) EnterComputerName(ctx *ComputerNameContext) {}

// ExitComputerName is called when production computerName is exited.
func (s *BaseCobol85Listener) ExitComputerName(ctx *ComputerNameContext) {}

// EnterConditionName is called when production conditionName is entered.
func (s *BaseCobol85Listener) EnterConditionName(ctx *ConditionNameContext) {}

// ExitConditionName is called when production conditionName is exited.
func (s *BaseCobol85Listener) ExitConditionName(ctx *ConditionNameContext) {}

// EnterDataName is called when production dataName is entered.
func (s *BaseCobol85Listener) EnterDataName(ctx *DataNameContext) {}

// ExitDataName is called when production dataName is exited.
func (s *BaseCobol85Listener) ExitDataName(ctx *DataNameContext) {}

// EnterDataDescName is called when production dataDescName is entered.
func (s *BaseCobol85Listener) EnterDataDescName(ctx *DataDescNameContext) {}

// ExitDataDescName is called when production dataDescName is exited.
func (s *BaseCobol85Listener) ExitDataDescName(ctx *DataDescNameContext) {}

// EnterEnvironmentName is called when production environmentName is entered.
func (s *BaseCobol85Listener) EnterEnvironmentName(ctx *EnvironmentNameContext) {}

// ExitEnvironmentName is called when production environmentName is exited.
func (s *BaseCobol85Listener) ExitEnvironmentName(ctx *EnvironmentNameContext) {}

// EnterFileName is called when production fileName is entered.
func (s *BaseCobol85Listener) EnterFileName(ctx *FileNameContext) {}

// ExitFileName is called when production fileName is exited.
func (s *BaseCobol85Listener) ExitFileName(ctx *FileNameContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseCobol85Listener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseCobol85Listener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseCobol85Listener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseCobol85Listener) ExitIndexName(ctx *IndexNameContext) {}

// EnterLanguageName is called when production languageName is entered.
func (s *BaseCobol85Listener) EnterLanguageName(ctx *LanguageNameContext) {}

// ExitLanguageName is called when production languageName is exited.
func (s *BaseCobol85Listener) ExitLanguageName(ctx *LanguageNameContext) {}

// EnterLibraryName is called when production libraryName is entered.
func (s *BaseCobol85Listener) EnterLibraryName(ctx *LibraryNameContext) {}

// ExitLibraryName is called when production libraryName is exited.
func (s *BaseCobol85Listener) ExitLibraryName(ctx *LibraryNameContext) {}

// EnterLocalName is called when production localName is entered.
func (s *BaseCobol85Listener) EnterLocalName(ctx *LocalNameContext) {}

// ExitLocalName is called when production localName is exited.
func (s *BaseCobol85Listener) ExitLocalName(ctx *LocalNameContext) {}

// EnterMnemonicName is called when production mnemonicName is entered.
func (s *BaseCobol85Listener) EnterMnemonicName(ctx *MnemonicNameContext) {}

// ExitMnemonicName is called when production mnemonicName is exited.
func (s *BaseCobol85Listener) ExitMnemonicName(ctx *MnemonicNameContext) {}

// EnterParagraphName is called when production paragraphName is entered.
func (s *BaseCobol85Listener) EnterParagraphName(ctx *ParagraphNameContext) {}

// ExitParagraphName is called when production paragraphName is exited.
func (s *BaseCobol85Listener) ExitParagraphName(ctx *ParagraphNameContext) {}

// EnterProcedureName is called when production procedureName is entered.
func (s *BaseCobol85Listener) EnterProcedureName(ctx *ProcedureNameContext) {}

// ExitProcedureName is called when production procedureName is exited.
func (s *BaseCobol85Listener) ExitProcedureName(ctx *ProcedureNameContext) {}

// EnterProgramName is called when production programName is entered.
func (s *BaseCobol85Listener) EnterProgramName(ctx *ProgramNameContext) {}

// ExitProgramName is called when production programName is exited.
func (s *BaseCobol85Listener) ExitProgramName(ctx *ProgramNameContext) {}

// EnterRecordName is called when production recordName is entered.
func (s *BaseCobol85Listener) EnterRecordName(ctx *RecordNameContext) {}

// ExitRecordName is called when production recordName is exited.
func (s *BaseCobol85Listener) ExitRecordName(ctx *RecordNameContext) {}

// EnterReportName is called when production reportName is entered.
func (s *BaseCobol85Listener) EnterReportName(ctx *ReportNameContext) {}

// ExitReportName is called when production reportName is exited.
func (s *BaseCobol85Listener) ExitReportName(ctx *ReportNameContext) {}

// EnterRoutineName is called when production routineName is entered.
func (s *BaseCobol85Listener) EnterRoutineName(ctx *RoutineNameContext) {}

// ExitRoutineName is called when production routineName is exited.
func (s *BaseCobol85Listener) ExitRoutineName(ctx *RoutineNameContext) {}

// EnterScreenName is called when production screenName is entered.
func (s *BaseCobol85Listener) EnterScreenName(ctx *ScreenNameContext) {}

// ExitScreenName is called when production screenName is exited.
func (s *BaseCobol85Listener) ExitScreenName(ctx *ScreenNameContext) {}

// EnterSectionName is called when production sectionName is entered.
func (s *BaseCobol85Listener) EnterSectionName(ctx *SectionNameContext) {}

// ExitSectionName is called when production sectionName is exited.
func (s *BaseCobol85Listener) ExitSectionName(ctx *SectionNameContext) {}

// EnterSystemName is called when production systemName is entered.
func (s *BaseCobol85Listener) EnterSystemName(ctx *SystemNameContext) {}

// ExitSystemName is called when production systemName is exited.
func (s *BaseCobol85Listener) ExitSystemName(ctx *SystemNameContext) {}

// EnterSymbolicCharacter is called when production symbolicCharacter is entered.
func (s *BaseCobol85Listener) EnterSymbolicCharacter(ctx *SymbolicCharacterContext) {}

// ExitSymbolicCharacter is called when production symbolicCharacter is exited.
func (s *BaseCobol85Listener) ExitSymbolicCharacter(ctx *SymbolicCharacterContext) {}

// EnterTextName is called when production textName is entered.
func (s *BaseCobol85Listener) EnterTextName(ctx *TextNameContext) {}

// ExitTextName is called when production textName is exited.
func (s *BaseCobol85Listener) ExitTextName(ctx *TextNameContext) {}

// EnterCobolWord is called when production cobolWord is entered.
func (s *BaseCobol85Listener) EnterCobolWord(ctx *CobolWordContext) {}

// ExitCobolWord is called when production cobolWord is exited.
func (s *BaseCobol85Listener) ExitCobolWord(ctx *CobolWordContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCobol85Listener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCobol85Listener) ExitLiteral(ctx *LiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseCobol85Listener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseCobol85Listener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseCobol85Listener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseCobol85Listener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseCobol85Listener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseCobol85Listener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterCicsDfhRespLiteral is called when production cicsDfhRespLiteral is entered.
func (s *BaseCobol85Listener) EnterCicsDfhRespLiteral(ctx *CicsDfhRespLiteralContext) {}

// ExitCicsDfhRespLiteral is called when production cicsDfhRespLiteral is exited.
func (s *BaseCobol85Listener) ExitCicsDfhRespLiteral(ctx *CicsDfhRespLiteralContext) {}

// EnterCicsDfhValueLiteral is called when production cicsDfhValueLiteral is entered.
func (s *BaseCobol85Listener) EnterCicsDfhValueLiteral(ctx *CicsDfhValueLiteralContext) {}

// ExitCicsDfhValueLiteral is called when production cicsDfhValueLiteral is exited.
func (s *BaseCobol85Listener) ExitCicsDfhValueLiteral(ctx *CicsDfhValueLiteralContext) {}

// EnterFigurativeConstant is called when production figurativeConstant is entered.
func (s *BaseCobol85Listener) EnterFigurativeConstant(ctx *FigurativeConstantContext) {}

// ExitFigurativeConstant is called when production figurativeConstant is exited.
func (s *BaseCobol85Listener) ExitFigurativeConstant(ctx *FigurativeConstantContext) {}

// EnterSpecialRegister is called when production specialRegister is entered.
func (s *BaseCobol85Listener) EnterSpecialRegister(ctx *SpecialRegisterContext) {}

// ExitSpecialRegister is called when production specialRegister is exited.
func (s *BaseCobol85Listener) ExitSpecialRegister(ctx *SpecialRegisterContext) {}

// EnterCommentEntry is called when production commentEntry is entered.
func (s *BaseCobol85Listener) EnterCommentEntry(ctx *CommentEntryContext) {}

// ExitCommentEntry is called when production commentEntry is exited.
func (s *BaseCobol85Listener) ExitCommentEntry(ctx *CommentEntryContext) {}
