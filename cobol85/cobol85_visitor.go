// Generated from Cobol85.g4 by ANTLR 4.7.

package cobol85 // Cobol85
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by Cobol85Parser.
type Cobol85Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Cobol85Parser#startRule.
	VisitStartRule(ctx *StartRuleContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#programUnit.
	VisitProgramUnit(ctx *ProgramUnitContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#endProgramStatement.
	VisitEndProgramStatement(ctx *EndProgramStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#compilerOptions.
	VisitCompilerOptions(ctx *CompilerOptionsContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#compilerOption.
	VisitCompilerOption(ctx *CompilerOptionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#compilerSubOption.
	VisitCompilerSubOption(ctx *CompilerSubOptionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#identificationDivision.
	VisitIdentificationDivision(ctx *IdentificationDivisionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#identificationDivisionBody.
	VisitIdentificationDivisionBody(ctx *IdentificationDivisionBodyContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#programIdParagraph.
	VisitProgramIdParagraph(ctx *ProgramIdParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#authorParagraph.
	VisitAuthorParagraph(ctx *AuthorParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#installationParagraph.
	VisitInstallationParagraph(ctx *InstallationParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dateWrittenParagraph.
	VisitDateWrittenParagraph(ctx *DateWrittenParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dateCompiledParagraph.
	VisitDateCompiledParagraph(ctx *DateCompiledParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#securityParagraph.
	VisitSecurityParagraph(ctx *SecurityParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#remarksParagraph.
	VisitRemarksParagraph(ctx *RemarksParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#environmentDivision.
	VisitEnvironmentDivision(ctx *EnvironmentDivisionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#environmentDivisionBody.
	VisitEnvironmentDivisionBody(ctx *EnvironmentDivisionBodyContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#configurationSection.
	VisitConfigurationSection(ctx *ConfigurationSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#configurationSectionParagraph.
	VisitConfigurationSectionParagraph(ctx *ConfigurationSectionParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sourceComputerParagraph.
	VisitSourceComputerParagraph(ctx *SourceComputerParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#objectComputerParagraph.
	VisitObjectComputerParagraph(ctx *ObjectComputerParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#objectComputerClause.
	VisitObjectComputerClause(ctx *ObjectComputerClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#memorySizeClause.
	VisitMemorySizeClause(ctx *MemorySizeClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#diskSizeClause.
	VisitDiskSizeClause(ctx *DiskSizeClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#collatingSequenceClause.
	VisitCollatingSequenceClause(ctx *CollatingSequenceClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#collatingSequenceClauseAlphanumeric.
	VisitCollatingSequenceClauseAlphanumeric(ctx *CollatingSequenceClauseAlphanumericContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#collatingSequenceClauseNational.
	VisitCollatingSequenceClauseNational(ctx *CollatingSequenceClauseNationalContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#segmentLimitClause.
	VisitSegmentLimitClause(ctx *SegmentLimitClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#characterSetClause.
	VisitCharacterSetClause(ctx *CharacterSetClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#specialNamesParagraph.
	VisitSpecialNamesParagraph(ctx *SpecialNamesParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#specialNameClause.
	VisitSpecialNameClause(ctx *SpecialNameClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#alphabetClause.
	VisitAlphabetClause(ctx *AlphabetClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#alphabetClauseFormat1.
	VisitAlphabetClauseFormat1(ctx *AlphabetClauseFormat1Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#alphabetLiterals.
	VisitAlphabetLiterals(ctx *AlphabetLiteralsContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#alphabetThrough.
	VisitAlphabetThrough(ctx *AlphabetThroughContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#alphabetAlso.
	VisitAlphabetAlso(ctx *AlphabetAlsoContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#alphabetClauseFormat2.
	VisitAlphabetClauseFormat2(ctx *AlphabetClauseFormat2Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#channelClause.
	VisitChannelClause(ctx *ChannelClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#classClause.
	VisitClassClause(ctx *ClassClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#classClauseThrough.
	VisitClassClauseThrough(ctx *ClassClauseThroughContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#classClauseFrom.
	VisitClassClauseFrom(ctx *ClassClauseFromContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#classClauseTo.
	VisitClassClauseTo(ctx *ClassClauseToContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#currencySignClause.
	VisitCurrencySignClause(ctx *CurrencySignClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#decimalPointClause.
	VisitDecimalPointClause(ctx *DecimalPointClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#defaultComputationalSignClause.
	VisitDefaultComputationalSignClause(ctx *DefaultComputationalSignClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#defaultDisplaySignClause.
	VisitDefaultDisplaySignClause(ctx *DefaultDisplaySignClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#environmentSwitchNameClause.
	VisitEnvironmentSwitchNameClause(ctx *EnvironmentSwitchNameClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#environmentSwitchNameSpecialNamesStatusPhrase.
	VisitEnvironmentSwitchNameSpecialNamesStatusPhrase(ctx *EnvironmentSwitchNameSpecialNamesStatusPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#odtClause.
	VisitOdtClause(ctx *OdtClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reserveNetworkClause.
	VisitReserveNetworkClause(ctx *ReserveNetworkClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#symbolicCharactersClause.
	VisitSymbolicCharactersClause(ctx *SymbolicCharactersClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#symbolicCharacters.
	VisitSymbolicCharacters(ctx *SymbolicCharactersContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inputOutputSection.
	VisitInputOutputSection(ctx *InputOutputSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inputOutputSectionParagraph.
	VisitInputOutputSectionParagraph(ctx *InputOutputSectionParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#fileControlParagraph.
	VisitFileControlParagraph(ctx *FileControlParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#fileControlEntry.
	VisitFileControlEntry(ctx *FileControlEntryContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#fileControlClause.
	VisitFileControlClause(ctx *FileControlClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#assignClause.
	VisitAssignClause(ctx *AssignClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reserveClause.
	VisitReserveClause(ctx *ReserveClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#organizationClause.
	VisitOrganizationClause(ctx *OrganizationClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#paddingCharacterClause.
	VisitPaddingCharacterClause(ctx *PaddingCharacterClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#recordDelimiterClause.
	VisitRecordDelimiterClause(ctx *RecordDelimiterClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#accessModeClause.
	VisitAccessModeClause(ctx *AccessModeClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#recordKeyClause.
	VisitRecordKeyClause(ctx *RecordKeyClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#alternateRecordKeyClause.
	VisitAlternateRecordKeyClause(ctx *AlternateRecordKeyClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#passwordClause.
	VisitPasswordClause(ctx *PasswordClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#fileStatusClause.
	VisitFileStatusClause(ctx *FileStatusClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#relativeKeyClause.
	VisitRelativeKeyClause(ctx *RelativeKeyClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#ioControlParagraph.
	VisitIoControlParagraph(ctx *IoControlParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#ioControlClause.
	VisitIoControlClause(ctx *IoControlClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#rerunClause.
	VisitRerunClause(ctx *RerunClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#rerunEveryRecords.
	VisitRerunEveryRecords(ctx *RerunEveryRecordsContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#rerunEveryOf.
	VisitRerunEveryOf(ctx *RerunEveryOfContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#rerunEveryClock.
	VisitRerunEveryClock(ctx *RerunEveryClockContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sameClause.
	VisitSameClause(ctx *SameClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#multipleFileClause.
	VisitMultipleFileClause(ctx *MultipleFileClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#multipleFilePosition.
	VisitMultipleFilePosition(ctx *MultipleFilePositionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#commitmentControlClause.
	VisitCommitmentControlClause(ctx *CommitmentControlClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataDivision.
	VisitDataDivision(ctx *DataDivisionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataDivisionBody.
	VisitDataDivisionBody(ctx *DataDivisionBodyContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#fileSection.
	VisitFileSection(ctx *FileSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#fileDescriptionEntry.
	VisitFileDescriptionEntry(ctx *FileDescriptionEntryContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#fileDescriptionEntryClause.
	VisitFileDescriptionEntryClause(ctx *FileDescriptionEntryClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#externalClause.
	VisitExternalClause(ctx *ExternalClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#globalClause.
	VisitGlobalClause(ctx *GlobalClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#blockContainsClause.
	VisitBlockContainsClause(ctx *BlockContainsClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#blockContainsTo.
	VisitBlockContainsTo(ctx *BlockContainsToContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#recordContainsClause.
	VisitRecordContainsClause(ctx *RecordContainsClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#recordContainsClauseFormat1.
	VisitRecordContainsClauseFormat1(ctx *RecordContainsClauseFormat1Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#recordContainsClauseFormat2.
	VisitRecordContainsClauseFormat2(ctx *RecordContainsClauseFormat2Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#recordContainsClauseFormat3.
	VisitRecordContainsClauseFormat3(ctx *RecordContainsClauseFormat3Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#recordContainsTo.
	VisitRecordContainsTo(ctx *RecordContainsToContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#labelRecordsClause.
	VisitLabelRecordsClause(ctx *LabelRecordsClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#valueOfClause.
	VisitValueOfClause(ctx *ValueOfClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#valuePair.
	VisitValuePair(ctx *ValuePairContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataRecordsClause.
	VisitDataRecordsClause(ctx *DataRecordsClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#linageClause.
	VisitLinageClause(ctx *LinageClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#linageAt.
	VisitLinageAt(ctx *LinageAtContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#linageFootingAt.
	VisitLinageFootingAt(ctx *LinageFootingAtContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#linageLinesAtTop.
	VisitLinageLinesAtTop(ctx *LinageLinesAtTopContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#linageLinesAtBottom.
	VisitLinageLinesAtBottom(ctx *LinageLinesAtBottomContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#recordingModeClause.
	VisitRecordingModeClause(ctx *RecordingModeClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#modeStatement.
	VisitModeStatement(ctx *ModeStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#codeSetClause.
	VisitCodeSetClause(ctx *CodeSetClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportClause.
	VisitReportClause(ctx *ReportClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataBaseSection.
	VisitDataBaseSection(ctx *DataBaseSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataBaseSectionEntry.
	VisitDataBaseSectionEntry(ctx *DataBaseSectionEntryContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#workingStorageSection.
	VisitWorkingStorageSection(ctx *WorkingStorageSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#linkageSection.
	VisitLinkageSection(ctx *LinkageSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#communicationSection.
	VisitCommunicationSection(ctx *CommunicationSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#communicationDescriptionEntry.
	VisitCommunicationDescriptionEntry(ctx *CommunicationDescriptionEntryContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#communicationDescriptionEntryFormat1.
	VisitCommunicationDescriptionEntryFormat1(ctx *CommunicationDescriptionEntryFormat1Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#communicationDescriptionEntryFormat2.
	VisitCommunicationDescriptionEntryFormat2(ctx *CommunicationDescriptionEntryFormat2Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#communicationDescriptionEntryFormat3.
	VisitCommunicationDescriptionEntryFormat3(ctx *CommunicationDescriptionEntryFormat3Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#destinationCountClause.
	VisitDestinationCountClause(ctx *DestinationCountClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#destinationTableClause.
	VisitDestinationTableClause(ctx *DestinationTableClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#endKeyClause.
	VisitEndKeyClause(ctx *EndKeyClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#errorKeyClause.
	VisitErrorKeyClause(ctx *ErrorKeyClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#messageCountClause.
	VisitMessageCountClause(ctx *MessageCountClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#messageDateClause.
	VisitMessageDateClause(ctx *MessageDateClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#messageTimeClause.
	VisitMessageTimeClause(ctx *MessageTimeClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#statusKeyClause.
	VisitStatusKeyClause(ctx *StatusKeyClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#symbolicDestinationClause.
	VisitSymbolicDestinationClause(ctx *SymbolicDestinationClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#symbolicQueueClause.
	VisitSymbolicQueueClause(ctx *SymbolicQueueClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#symbolicSourceClause.
	VisitSymbolicSourceClause(ctx *SymbolicSourceClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#symbolicTerminalClause.
	VisitSymbolicTerminalClause(ctx *SymbolicTerminalClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#symbolicSubQueueClause.
	VisitSymbolicSubQueueClause(ctx *SymbolicSubQueueClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#textLengthClause.
	VisitTextLengthClause(ctx *TextLengthClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#localStorageSection.
	VisitLocalStorageSection(ctx *LocalStorageSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenSection.
	VisitScreenSection(ctx *ScreenSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionEntry.
	VisitScreenDescriptionEntry(ctx *ScreenDescriptionEntryContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionBlankClause.
	VisitScreenDescriptionBlankClause(ctx *ScreenDescriptionBlankClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionBellClause.
	VisitScreenDescriptionBellClause(ctx *ScreenDescriptionBellClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionBlinkClause.
	VisitScreenDescriptionBlinkClause(ctx *ScreenDescriptionBlinkClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionEraseClause.
	VisitScreenDescriptionEraseClause(ctx *ScreenDescriptionEraseClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionLightClause.
	VisitScreenDescriptionLightClause(ctx *ScreenDescriptionLightClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionGridClause.
	VisitScreenDescriptionGridClause(ctx *ScreenDescriptionGridClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionReverseVideoClause.
	VisitScreenDescriptionReverseVideoClause(ctx *ScreenDescriptionReverseVideoClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionUnderlineClause.
	VisitScreenDescriptionUnderlineClause(ctx *ScreenDescriptionUnderlineClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionSizeClause.
	VisitScreenDescriptionSizeClause(ctx *ScreenDescriptionSizeClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionLineClause.
	VisitScreenDescriptionLineClause(ctx *ScreenDescriptionLineClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionColumnClause.
	VisitScreenDescriptionColumnClause(ctx *ScreenDescriptionColumnClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionForegroundColorClause.
	VisitScreenDescriptionForegroundColorClause(ctx *ScreenDescriptionForegroundColorClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionBackgroundColorClause.
	VisitScreenDescriptionBackgroundColorClause(ctx *ScreenDescriptionBackgroundColorClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionControlClause.
	VisitScreenDescriptionControlClause(ctx *ScreenDescriptionControlClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionValueClause.
	VisitScreenDescriptionValueClause(ctx *ScreenDescriptionValueClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionPictureClause.
	VisitScreenDescriptionPictureClause(ctx *ScreenDescriptionPictureClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionFromClause.
	VisitScreenDescriptionFromClause(ctx *ScreenDescriptionFromClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionToClause.
	VisitScreenDescriptionToClause(ctx *ScreenDescriptionToClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionUsingClause.
	VisitScreenDescriptionUsingClause(ctx *ScreenDescriptionUsingClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionUsageClause.
	VisitScreenDescriptionUsageClause(ctx *ScreenDescriptionUsageClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionBlankWhenZeroClause.
	VisitScreenDescriptionBlankWhenZeroClause(ctx *ScreenDescriptionBlankWhenZeroClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionJustifiedClause.
	VisitScreenDescriptionJustifiedClause(ctx *ScreenDescriptionJustifiedClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionSignClause.
	VisitScreenDescriptionSignClause(ctx *ScreenDescriptionSignClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionAutoClause.
	VisitScreenDescriptionAutoClause(ctx *ScreenDescriptionAutoClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionSecureClause.
	VisitScreenDescriptionSecureClause(ctx *ScreenDescriptionSecureClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionRequiredClause.
	VisitScreenDescriptionRequiredClause(ctx *ScreenDescriptionRequiredClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionPromptClause.
	VisitScreenDescriptionPromptClause(ctx *ScreenDescriptionPromptClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionPromptOccursClause.
	VisitScreenDescriptionPromptOccursClause(ctx *ScreenDescriptionPromptOccursClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionFullClause.
	VisitScreenDescriptionFullClause(ctx *ScreenDescriptionFullClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenDescriptionZeroFillClause.
	VisitScreenDescriptionZeroFillClause(ctx *ScreenDescriptionZeroFillClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportSection.
	VisitReportSection(ctx *ReportSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportDescription.
	VisitReportDescription(ctx *ReportDescriptionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportDescriptionEntry.
	VisitReportDescriptionEntry(ctx *ReportDescriptionEntryContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportDescriptionGlobalClause.
	VisitReportDescriptionGlobalClause(ctx *ReportDescriptionGlobalClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportDescriptionPageLimitClause.
	VisitReportDescriptionPageLimitClause(ctx *ReportDescriptionPageLimitClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportDescriptionHeadingClause.
	VisitReportDescriptionHeadingClause(ctx *ReportDescriptionHeadingClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportDescriptionFirstDetailClause.
	VisitReportDescriptionFirstDetailClause(ctx *ReportDescriptionFirstDetailClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportDescriptionLastDetailClause.
	VisitReportDescriptionLastDetailClause(ctx *ReportDescriptionLastDetailClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportDescriptionFootingClause.
	VisitReportDescriptionFootingClause(ctx *ReportDescriptionFootingClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupDescriptionEntry.
	VisitReportGroupDescriptionEntry(ctx *ReportGroupDescriptionEntryContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupDescriptionEntryFormat1.
	VisitReportGroupDescriptionEntryFormat1(ctx *ReportGroupDescriptionEntryFormat1Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupDescriptionEntryFormat2.
	VisitReportGroupDescriptionEntryFormat2(ctx *ReportGroupDescriptionEntryFormat2Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupDescriptionEntryFormat3.
	VisitReportGroupDescriptionEntryFormat3(ctx *ReportGroupDescriptionEntryFormat3Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupBlankWhenZeroClause.
	VisitReportGroupBlankWhenZeroClause(ctx *ReportGroupBlankWhenZeroClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupColumnNumberClause.
	VisitReportGroupColumnNumberClause(ctx *ReportGroupColumnNumberClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupIndicateClause.
	VisitReportGroupIndicateClause(ctx *ReportGroupIndicateClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupJustifiedClause.
	VisitReportGroupJustifiedClause(ctx *ReportGroupJustifiedClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupLineNumberClause.
	VisitReportGroupLineNumberClause(ctx *ReportGroupLineNumberClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupLineNumberNextPage.
	VisitReportGroupLineNumberNextPage(ctx *ReportGroupLineNumberNextPageContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupLineNumberPlus.
	VisitReportGroupLineNumberPlus(ctx *ReportGroupLineNumberPlusContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupNextGroupClause.
	VisitReportGroupNextGroupClause(ctx *ReportGroupNextGroupClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupNextGroupPlus.
	VisitReportGroupNextGroupPlus(ctx *ReportGroupNextGroupPlusContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupNextGroupNextPage.
	VisitReportGroupNextGroupNextPage(ctx *ReportGroupNextGroupNextPageContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupPictureClause.
	VisitReportGroupPictureClause(ctx *ReportGroupPictureClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupResetClause.
	VisitReportGroupResetClause(ctx *ReportGroupResetClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupSignClause.
	VisitReportGroupSignClause(ctx *ReportGroupSignClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupSourceClause.
	VisitReportGroupSourceClause(ctx *ReportGroupSourceClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupSumClause.
	VisitReportGroupSumClause(ctx *ReportGroupSumClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupTypeClause.
	VisitReportGroupTypeClause(ctx *ReportGroupTypeClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupTypeReportHeading.
	VisitReportGroupTypeReportHeading(ctx *ReportGroupTypeReportHeadingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupTypePageHeading.
	VisitReportGroupTypePageHeading(ctx *ReportGroupTypePageHeadingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupTypeControlHeading.
	VisitReportGroupTypeControlHeading(ctx *ReportGroupTypeControlHeadingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupTypeDetail.
	VisitReportGroupTypeDetail(ctx *ReportGroupTypeDetailContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupTypeControlFooting.
	VisitReportGroupTypeControlFooting(ctx *ReportGroupTypeControlFootingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupUsageClause.
	VisitReportGroupUsageClause(ctx *ReportGroupUsageClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupTypePageFooting.
	VisitReportGroupTypePageFooting(ctx *ReportGroupTypePageFootingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupTypeReportFooting.
	VisitReportGroupTypeReportFooting(ctx *ReportGroupTypeReportFootingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportGroupValueClause.
	VisitReportGroupValueClause(ctx *ReportGroupValueClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#programLibrarySection.
	VisitProgramLibrarySection(ctx *ProgramLibrarySectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryDescriptionEntry.
	VisitLibraryDescriptionEntry(ctx *LibraryDescriptionEntryContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryDescriptionEntryFormat1.
	VisitLibraryDescriptionEntryFormat1(ctx *LibraryDescriptionEntryFormat1Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryDescriptionEntryFormat2.
	VisitLibraryDescriptionEntryFormat2(ctx *LibraryDescriptionEntryFormat2Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryAttributeClauseFormat1.
	VisitLibraryAttributeClauseFormat1(ctx *LibraryAttributeClauseFormat1Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryAttributeClauseFormat2.
	VisitLibraryAttributeClauseFormat2(ctx *LibraryAttributeClauseFormat2Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryAttributeFunction.
	VisitLibraryAttributeFunction(ctx *LibraryAttributeFunctionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryAttributeParameter.
	VisitLibraryAttributeParameter(ctx *LibraryAttributeParameterContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryAttributeTitle.
	VisitLibraryAttributeTitle(ctx *LibraryAttributeTitleContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryEntryProcedureClauseFormat1.
	VisitLibraryEntryProcedureClauseFormat1(ctx *LibraryEntryProcedureClauseFormat1Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryEntryProcedureClauseFormat2.
	VisitLibraryEntryProcedureClauseFormat2(ctx *LibraryEntryProcedureClauseFormat2Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryEntryProcedureForClause.
	VisitLibraryEntryProcedureForClause(ctx *LibraryEntryProcedureForClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryEntryProcedureGivingClause.
	VisitLibraryEntryProcedureGivingClause(ctx *LibraryEntryProcedureGivingClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryEntryProcedureUsingClause.
	VisitLibraryEntryProcedureUsingClause(ctx *LibraryEntryProcedureUsingClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryEntryProcedureUsingName.
	VisitLibraryEntryProcedureUsingName(ctx *LibraryEntryProcedureUsingNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryEntryProcedureWithClause.
	VisitLibraryEntryProcedureWithClause(ctx *LibraryEntryProcedureWithClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryEntryProcedureWithName.
	VisitLibraryEntryProcedureWithName(ctx *LibraryEntryProcedureWithNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryIsCommonClause.
	VisitLibraryIsCommonClause(ctx *LibraryIsCommonClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryIsGlobalClause.
	VisitLibraryIsGlobalClause(ctx *LibraryIsGlobalClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataDescriptionEntry.
	VisitDataDescriptionEntry(ctx *DataDescriptionEntryContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataDescriptionEntryFormat1.
	VisitDataDescriptionEntryFormat1(ctx *DataDescriptionEntryFormat1Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataDescriptionEntryFormat2.
	VisitDataDescriptionEntryFormat2(ctx *DataDescriptionEntryFormat2Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataDescriptionEntryFormat3.
	VisitDataDescriptionEntryFormat3(ctx *DataDescriptionEntryFormat3Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataDescriptionEntryExecSql.
	VisitDataDescriptionEntryExecSql(ctx *DataDescriptionEntryExecSqlContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataAlignedClause.
	VisitDataAlignedClause(ctx *DataAlignedClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataBlankWhenZeroClause.
	VisitDataBlankWhenZeroClause(ctx *DataBlankWhenZeroClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataCommonOwnLocalClause.
	VisitDataCommonOwnLocalClause(ctx *DataCommonOwnLocalClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataExternalClause.
	VisitDataExternalClause(ctx *DataExternalClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataGlobalClause.
	VisitDataGlobalClause(ctx *DataGlobalClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataIntegerStringClause.
	VisitDataIntegerStringClause(ctx *DataIntegerStringClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataJustifiedClause.
	VisitDataJustifiedClause(ctx *DataJustifiedClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataOccursClause.
	VisitDataOccursClause(ctx *DataOccursClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataOccursTo.
	VisitDataOccursTo(ctx *DataOccursToContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataOccursSort.
	VisitDataOccursSort(ctx *DataOccursSortContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataPictureClause.
	VisitDataPictureClause(ctx *DataPictureClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#pictureString.
	VisitPictureString(ctx *PictureStringContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#pictureChars.
	VisitPictureChars(ctx *PictureCharsContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#pictureCharsKeyword.
	VisitPictureCharsKeyword(ctx *PictureCharsKeywordContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#pictureCardinality.
	VisitPictureCardinality(ctx *PictureCardinalityContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataReceivedByClause.
	VisitDataReceivedByClause(ctx *DataReceivedByClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataRecordAreaClause.
	VisitDataRecordAreaClause(ctx *DataRecordAreaClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataRedefinesClause.
	VisitDataRedefinesClause(ctx *DataRedefinesClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataRenamesClause.
	VisitDataRenamesClause(ctx *DataRenamesClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataSignClause.
	VisitDataSignClause(ctx *DataSignClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataSynchronizedClause.
	VisitDataSynchronizedClause(ctx *DataSynchronizedClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataThreadLocalClause.
	VisitDataThreadLocalClause(ctx *DataThreadLocalClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataTypeClause.
	VisitDataTypeClause(ctx *DataTypeClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataTypeDefClause.
	VisitDataTypeDefClause(ctx *DataTypeDefClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataUsageClause.
	VisitDataUsageClause(ctx *DataUsageClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataUsingClause.
	VisitDataUsingClause(ctx *DataUsingClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataValueClause.
	VisitDataValueClause(ctx *DataValueClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataValueInterval.
	VisitDataValueInterval(ctx *DataValueIntervalContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataValueIntervalFrom.
	VisitDataValueIntervalFrom(ctx *DataValueIntervalFromContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataValueIntervalTo.
	VisitDataValueIntervalTo(ctx *DataValueIntervalToContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataWithLowerBoundsClause.
	VisitDataWithLowerBoundsClause(ctx *DataWithLowerBoundsClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#procedureDivision.
	VisitProcedureDivision(ctx *ProcedureDivisionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#procedureDivisionUsingClause.
	VisitProcedureDivisionUsingClause(ctx *ProcedureDivisionUsingClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#procedureDivisionGivingClause.
	VisitProcedureDivisionGivingClause(ctx *ProcedureDivisionGivingClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#procedureDeclaratives.
	VisitProcedureDeclaratives(ctx *ProcedureDeclarativesContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#procedureDeclarative.
	VisitProcedureDeclarative(ctx *ProcedureDeclarativeContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#procedureSectionHeader.
	VisitProcedureSectionHeader(ctx *ProcedureSectionHeaderContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#procedureDivisionBody.
	VisitProcedureDivisionBody(ctx *ProcedureDivisionBodyContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#procedureSection.
	VisitProcedureSection(ctx *ProcedureSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#paragraphs.
	VisitParagraphs(ctx *ParagraphsContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#paragraph.
	VisitParagraph(ctx *ParagraphContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sentence.
	VisitSentence(ctx *SentenceContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#acceptStatement.
	VisitAcceptStatement(ctx *AcceptStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#acceptFromDateStatement.
	VisitAcceptFromDateStatement(ctx *AcceptFromDateStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#acceptFromMnemonicStatement.
	VisitAcceptFromMnemonicStatement(ctx *AcceptFromMnemonicStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#acceptFromEscapeKeyStatement.
	VisitAcceptFromEscapeKeyStatement(ctx *AcceptFromEscapeKeyStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#acceptMessageCountStatement.
	VisitAcceptMessageCountStatement(ctx *AcceptMessageCountStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#addStatement.
	VisitAddStatement(ctx *AddStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#addToStatement.
	VisitAddToStatement(ctx *AddToStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#addToGivingStatement.
	VisitAddToGivingStatement(ctx *AddToGivingStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#addCorrespondingStatement.
	VisitAddCorrespondingStatement(ctx *AddCorrespondingStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#addFrom.
	VisitAddFrom(ctx *AddFromContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#addTo.
	VisitAddTo(ctx *AddToContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#addGiving.
	VisitAddGiving(ctx *AddGivingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#alteredGoTo.
	VisitAlteredGoTo(ctx *AlteredGoToContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#alterStatement.
	VisitAlterStatement(ctx *AlterStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#alterProceedTo.
	VisitAlterProceedTo(ctx *AlterProceedToContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#callStatement.
	VisitCallStatement(ctx *CallStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#callUsingPhrase.
	VisitCallUsingPhrase(ctx *CallUsingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#callUsingParameter.
	VisitCallUsingParameter(ctx *CallUsingParameterContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#callByReferencePhrase.
	VisitCallByReferencePhrase(ctx *CallByReferencePhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#callByReference.
	VisitCallByReference(ctx *CallByReferenceContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#callByValuePhrase.
	VisitCallByValuePhrase(ctx *CallByValuePhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#callByValue.
	VisitCallByValue(ctx *CallByValueContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#callByContentPhrase.
	VisitCallByContentPhrase(ctx *CallByContentPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#callByContent.
	VisitCallByContent(ctx *CallByContentContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#callGivingPhrase.
	VisitCallGivingPhrase(ctx *CallGivingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#cancelStatement.
	VisitCancelStatement(ctx *CancelStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#cancelCall.
	VisitCancelCall(ctx *CancelCallContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#closeStatement.
	VisitCloseStatement(ctx *CloseStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#closeFile.
	VisitCloseFile(ctx *CloseFileContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#closeReelUnitStatement.
	VisitCloseReelUnitStatement(ctx *CloseReelUnitStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#closeRelativeStatement.
	VisitCloseRelativeStatement(ctx *CloseRelativeStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#closePortFileIOStatement.
	VisitClosePortFileIOStatement(ctx *ClosePortFileIOStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#closePortFileIOUsing.
	VisitClosePortFileIOUsing(ctx *ClosePortFileIOUsingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#closePortFileIOUsingCloseDisposition.
	VisitClosePortFileIOUsingCloseDisposition(ctx *ClosePortFileIOUsingCloseDispositionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#closePortFileIOUsingAssociatedData.
	VisitClosePortFileIOUsingAssociatedData(ctx *ClosePortFileIOUsingAssociatedDataContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#closePortFileIOUsingAssociatedDataLength.
	VisitClosePortFileIOUsingAssociatedDataLength(ctx *ClosePortFileIOUsingAssociatedDataLengthContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#computeStatement.
	VisitComputeStatement(ctx *ComputeStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#computeStore.
	VisitComputeStore(ctx *ComputeStoreContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#disableStatement.
	VisitDisableStatement(ctx *DisableStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#displayStatement.
	VisitDisplayStatement(ctx *DisplayStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#displayOperand.
	VisitDisplayOperand(ctx *DisplayOperandContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#displayAt.
	VisitDisplayAt(ctx *DisplayAtContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#displayUpon.
	VisitDisplayUpon(ctx *DisplayUponContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#displayWith.
	VisitDisplayWith(ctx *DisplayWithContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#divideStatement.
	VisitDivideStatement(ctx *DivideStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#divideIntoStatement.
	VisitDivideIntoStatement(ctx *DivideIntoStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#divideIntoGivingStatement.
	VisitDivideIntoGivingStatement(ctx *DivideIntoGivingStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#divideByGivingStatement.
	VisitDivideByGivingStatement(ctx *DivideByGivingStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#divideGivingPhrase.
	VisitDivideGivingPhrase(ctx *DivideGivingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#divideGiving.
	VisitDivideGiving(ctx *DivideGivingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#divideRemainder.
	VisitDivideRemainder(ctx *DivideRemainderContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#enableStatement.
	VisitEnableStatement(ctx *EnableStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#entryStatement.
	VisitEntryStatement(ctx *EntryStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#evaluateStatement.
	VisitEvaluateStatement(ctx *EvaluateStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#evaluateSelect.
	VisitEvaluateSelect(ctx *EvaluateSelectContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#evaluateAlsoSelect.
	VisitEvaluateAlsoSelect(ctx *EvaluateAlsoSelectContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#evaluateWhenPhrase.
	VisitEvaluateWhenPhrase(ctx *EvaluateWhenPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#evaluateWhen.
	VisitEvaluateWhen(ctx *EvaluateWhenContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#evaluateCondition.
	VisitEvaluateCondition(ctx *EvaluateConditionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#evaluateThrough.
	VisitEvaluateThrough(ctx *EvaluateThroughContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#evaluateAlsoCondition.
	VisitEvaluateAlsoCondition(ctx *EvaluateAlsoConditionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#evaluateWhenOther.
	VisitEvaluateWhenOther(ctx *EvaluateWhenOtherContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#evaluateValue.
	VisitEvaluateValue(ctx *EvaluateValueContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#execCicsStatement.
	VisitExecCicsStatement(ctx *ExecCicsStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#execSqlStatement.
	VisitExecSqlStatement(ctx *ExecSqlStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#execSqlImsStatement.
	VisitExecSqlImsStatement(ctx *ExecSqlImsStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#exhibitStatement.
	VisitExhibitStatement(ctx *ExhibitStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#exhibitOperand.
	VisitExhibitOperand(ctx *ExhibitOperandContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#exitStatement.
	VisitExitStatement(ctx *ExitStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#generateStatement.
	VisitGenerateStatement(ctx *GenerateStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#gobackStatement.
	VisitGobackStatement(ctx *GobackStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#goToStatement.
	VisitGoToStatement(ctx *GoToStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#goToStatementSimple.
	VisitGoToStatementSimple(ctx *GoToStatementSimpleContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#goToDependingOnStatement.
	VisitGoToDependingOnStatement(ctx *GoToDependingOnStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#ifThen.
	VisitIfThen(ctx *IfThenContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#ifElse.
	VisitIfElse(ctx *IfElseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#initializeStatement.
	VisitInitializeStatement(ctx *InitializeStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#initializeReplacingPhrase.
	VisitInitializeReplacingPhrase(ctx *InitializeReplacingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#initializeReplacingBy.
	VisitInitializeReplacingBy(ctx *InitializeReplacingByContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#initiateStatement.
	VisitInitiateStatement(ctx *InitiateStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectStatement.
	VisitInspectStatement(ctx *InspectStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectTallyingPhrase.
	VisitInspectTallyingPhrase(ctx *InspectTallyingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectReplacingPhrase.
	VisitInspectReplacingPhrase(ctx *InspectReplacingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectTallyingReplacingPhrase.
	VisitInspectTallyingReplacingPhrase(ctx *InspectTallyingReplacingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectConvertingPhrase.
	VisitInspectConvertingPhrase(ctx *InspectConvertingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectFor.
	VisitInspectFor(ctx *InspectForContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectCharacters.
	VisitInspectCharacters(ctx *InspectCharactersContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectReplacingCharacters.
	VisitInspectReplacingCharacters(ctx *InspectReplacingCharactersContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectAllLeadings.
	VisitInspectAllLeadings(ctx *InspectAllLeadingsContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectReplacingAllLeadings.
	VisitInspectReplacingAllLeadings(ctx *InspectReplacingAllLeadingsContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectAllLeading.
	VisitInspectAllLeading(ctx *InspectAllLeadingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectReplacingAllLeading.
	VisitInspectReplacingAllLeading(ctx *InspectReplacingAllLeadingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectBy.
	VisitInspectBy(ctx *InspectByContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectTo.
	VisitInspectTo(ctx *InspectToContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inspectBeforeAfter.
	VisitInspectBeforeAfter(ctx *InspectBeforeAfterContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mergeStatement.
	VisitMergeStatement(ctx *MergeStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mergeOnKeyClause.
	VisitMergeOnKeyClause(ctx *MergeOnKeyClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mergeCollatingSequencePhrase.
	VisitMergeCollatingSequencePhrase(ctx *MergeCollatingSequencePhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mergeCollatingAlphanumeric.
	VisitMergeCollatingAlphanumeric(ctx *MergeCollatingAlphanumericContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mergeCollatingNational.
	VisitMergeCollatingNational(ctx *MergeCollatingNationalContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mergeUsing.
	VisitMergeUsing(ctx *MergeUsingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mergeOutputProcedurePhrase.
	VisitMergeOutputProcedurePhrase(ctx *MergeOutputProcedurePhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mergeOutputThrough.
	VisitMergeOutputThrough(ctx *MergeOutputThroughContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mergeGivingPhrase.
	VisitMergeGivingPhrase(ctx *MergeGivingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mergeGiving.
	VisitMergeGiving(ctx *MergeGivingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#moveStatement.
	VisitMoveStatement(ctx *MoveStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#moveToStatement.
	VisitMoveToStatement(ctx *MoveToStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#moveToSendingArea.
	VisitMoveToSendingArea(ctx *MoveToSendingAreaContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#moveCorrespondingToStatement.
	VisitMoveCorrespondingToStatement(ctx *MoveCorrespondingToStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#multiplyStatement.
	VisitMultiplyStatement(ctx *MultiplyStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#multiplyRegular.
	VisitMultiplyRegular(ctx *MultiplyRegularContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#multiplyRegularOperand.
	VisitMultiplyRegularOperand(ctx *MultiplyRegularOperandContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#multiplyGiving.
	VisitMultiplyGiving(ctx *MultiplyGivingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#multiplyGivingOperand.
	VisitMultiplyGivingOperand(ctx *MultiplyGivingOperandContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#multiplyGivingResult.
	VisitMultiplyGivingResult(ctx *MultiplyGivingResultContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#openStatement.
	VisitOpenStatement(ctx *OpenStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#openInputStatement.
	VisitOpenInputStatement(ctx *OpenInputStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#openInput.
	VisitOpenInput(ctx *OpenInputContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#openOutputStatement.
	VisitOpenOutputStatement(ctx *OpenOutputStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#openOutput.
	VisitOpenOutput(ctx *OpenOutputContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#openIOStatement.
	VisitOpenIOStatement(ctx *OpenIOStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#openExtendStatement.
	VisitOpenExtendStatement(ctx *OpenExtendStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performStatement.
	VisitPerformStatement(ctx *PerformStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performInlineStatement.
	VisitPerformInlineStatement(ctx *PerformInlineStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performProcedureStatement.
	VisitPerformProcedureStatement(ctx *PerformProcedureStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performType.
	VisitPerformType(ctx *PerformTypeContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performTimes.
	VisitPerformTimes(ctx *PerformTimesContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performUntil.
	VisitPerformUntil(ctx *PerformUntilContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performVarying.
	VisitPerformVarying(ctx *PerformVaryingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performVaryingClause.
	VisitPerformVaryingClause(ctx *PerformVaryingClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performVaryingPhrase.
	VisitPerformVaryingPhrase(ctx *PerformVaryingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performAfter.
	VisitPerformAfter(ctx *PerformAfterContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performFrom.
	VisitPerformFrom(ctx *PerformFromContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performBy.
	VisitPerformBy(ctx *PerformByContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#performTestClause.
	VisitPerformTestClause(ctx *PerformTestClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#purgeStatement.
	VisitPurgeStatement(ctx *PurgeStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#readStatement.
	VisitReadStatement(ctx *ReadStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#readInto.
	VisitReadInto(ctx *ReadIntoContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#readWith.
	VisitReadWith(ctx *ReadWithContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#readKey.
	VisitReadKey(ctx *ReadKeyContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveStatement.
	VisitReceiveStatement(ctx *ReceiveStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveFromStatement.
	VisitReceiveFromStatement(ctx *ReceiveFromStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveFrom.
	VisitReceiveFrom(ctx *ReceiveFromContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveIntoStatement.
	VisitReceiveIntoStatement(ctx *ReceiveIntoStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveNoData.
	VisitReceiveNoData(ctx *ReceiveNoDataContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveWithData.
	VisitReceiveWithData(ctx *ReceiveWithDataContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveBefore.
	VisitReceiveBefore(ctx *ReceiveBeforeContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveWith.
	VisitReceiveWith(ctx *ReceiveWithContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveThread.
	VisitReceiveThread(ctx *ReceiveThreadContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveSize.
	VisitReceiveSize(ctx *ReceiveSizeContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#receiveStatus.
	VisitReceiveStatus(ctx *ReceiveStatusContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#releaseStatement.
	VisitReleaseStatement(ctx *ReleaseStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#returnInto.
	VisitReturnInto(ctx *ReturnIntoContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#rewriteStatement.
	VisitRewriteStatement(ctx *RewriteStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#rewriteFrom.
	VisitRewriteFrom(ctx *RewriteFromContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#searchStatement.
	VisitSearchStatement(ctx *SearchStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#searchVarying.
	VisitSearchVarying(ctx *SearchVaryingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#searchWhen.
	VisitSearchWhen(ctx *SearchWhenContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sendStatement.
	VisitSendStatement(ctx *SendStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sendStatementSync.
	VisitSendStatementSync(ctx *SendStatementSyncContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sendStatementAsync.
	VisitSendStatementAsync(ctx *SendStatementAsyncContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sendFromPhrase.
	VisitSendFromPhrase(ctx *SendFromPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sendWithPhrase.
	VisitSendWithPhrase(ctx *SendWithPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sendReplacingPhrase.
	VisitSendReplacingPhrase(ctx *SendReplacingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sendAdvancingPhrase.
	VisitSendAdvancingPhrase(ctx *SendAdvancingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sendAdvancingPage.
	VisitSendAdvancingPage(ctx *SendAdvancingPageContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sendAdvancingLines.
	VisitSendAdvancingLines(ctx *SendAdvancingLinesContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sendAdvancingMnemonic.
	VisitSendAdvancingMnemonic(ctx *SendAdvancingMnemonicContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#setStatement.
	VisitSetStatement(ctx *SetStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#setToStatement.
	VisitSetToStatement(ctx *SetToStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#setUpDownByStatement.
	VisitSetUpDownByStatement(ctx *SetUpDownByStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#setTo.
	VisitSetTo(ctx *SetToContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#setToValue.
	VisitSetToValue(ctx *SetToValueContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#setByValue.
	VisitSetByValue(ctx *SetByValueContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortStatement.
	VisitSortStatement(ctx *SortStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortOnKeyClause.
	VisitSortOnKeyClause(ctx *SortOnKeyClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortDuplicatesPhrase.
	VisitSortDuplicatesPhrase(ctx *SortDuplicatesPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortCollatingSequencePhrase.
	VisitSortCollatingSequencePhrase(ctx *SortCollatingSequencePhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortCollatingAlphanumeric.
	VisitSortCollatingAlphanumeric(ctx *SortCollatingAlphanumericContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortCollatingNational.
	VisitSortCollatingNational(ctx *SortCollatingNationalContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortInputProcedurePhrase.
	VisitSortInputProcedurePhrase(ctx *SortInputProcedurePhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortInputThrough.
	VisitSortInputThrough(ctx *SortInputThroughContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortUsing.
	VisitSortUsing(ctx *SortUsingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortOutputProcedurePhrase.
	VisitSortOutputProcedurePhrase(ctx *SortOutputProcedurePhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortOutputThrough.
	VisitSortOutputThrough(ctx *SortOutputThroughContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortGivingPhrase.
	VisitSortGivingPhrase(ctx *SortGivingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sortGiving.
	VisitSortGiving(ctx *SortGivingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#startStatement.
	VisitStartStatement(ctx *StartStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#startKey.
	VisitStartKey(ctx *StartKeyContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#stopStatement.
	VisitStopStatement(ctx *StopStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#stringStatement.
	VisitStringStatement(ctx *StringStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#stringSendingPhrase.
	VisitStringSendingPhrase(ctx *StringSendingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#stringSending.
	VisitStringSending(ctx *StringSendingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#stringDelimitedByPhrase.
	VisitStringDelimitedByPhrase(ctx *StringDelimitedByPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#stringForPhrase.
	VisitStringForPhrase(ctx *StringForPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#stringIntoPhrase.
	VisitStringIntoPhrase(ctx *StringIntoPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#stringWithPointerPhrase.
	VisitStringWithPointerPhrase(ctx *StringWithPointerPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#subtractStatement.
	VisitSubtractStatement(ctx *SubtractStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#subtractFromStatement.
	VisitSubtractFromStatement(ctx *SubtractFromStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#subtractFromGivingStatement.
	VisitSubtractFromGivingStatement(ctx *SubtractFromGivingStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#subtractCorrespondingStatement.
	VisitSubtractCorrespondingStatement(ctx *SubtractCorrespondingStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#subtractSubtrahend.
	VisitSubtractSubtrahend(ctx *SubtractSubtrahendContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#subtractMinuend.
	VisitSubtractMinuend(ctx *SubtractMinuendContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#subtractMinuendGiving.
	VisitSubtractMinuendGiving(ctx *SubtractMinuendGivingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#subtractGiving.
	VisitSubtractGiving(ctx *SubtractGivingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#subtractMinuendCorresponding.
	VisitSubtractMinuendCorresponding(ctx *SubtractMinuendCorrespondingContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#terminateStatement.
	VisitTerminateStatement(ctx *TerminateStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#unstringStatement.
	VisitUnstringStatement(ctx *UnstringStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#unstringSendingPhrase.
	VisitUnstringSendingPhrase(ctx *UnstringSendingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#unstringDelimitedByPhrase.
	VisitUnstringDelimitedByPhrase(ctx *UnstringDelimitedByPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#unstringOrAllPhrase.
	VisitUnstringOrAllPhrase(ctx *UnstringOrAllPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#unstringIntoPhrase.
	VisitUnstringIntoPhrase(ctx *UnstringIntoPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#unstringInto.
	VisitUnstringInto(ctx *UnstringIntoContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#unstringDelimiterIn.
	VisitUnstringDelimiterIn(ctx *UnstringDelimiterInContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#unstringCountIn.
	VisitUnstringCountIn(ctx *UnstringCountInContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#unstringWithPointerPhrase.
	VisitUnstringWithPointerPhrase(ctx *UnstringWithPointerPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#unstringTallyingPhrase.
	VisitUnstringTallyingPhrase(ctx *UnstringTallyingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#useStatement.
	VisitUseStatement(ctx *UseStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#useAfterClause.
	VisitUseAfterClause(ctx *UseAfterClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#useAfterOn.
	VisitUseAfterOn(ctx *UseAfterOnContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#useDebugClause.
	VisitUseDebugClause(ctx *UseDebugClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#useDebugOn.
	VisitUseDebugOn(ctx *UseDebugOnContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#writeStatement.
	VisitWriteStatement(ctx *WriteStatementContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#writeFromPhrase.
	VisitWriteFromPhrase(ctx *WriteFromPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#writeAdvancingPhrase.
	VisitWriteAdvancingPhrase(ctx *WriteAdvancingPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#writeAdvancingPage.
	VisitWriteAdvancingPage(ctx *WriteAdvancingPageContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#writeAdvancingLines.
	VisitWriteAdvancingLines(ctx *WriteAdvancingLinesContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#writeAdvancingMnemonic.
	VisitWriteAdvancingMnemonic(ctx *WriteAdvancingMnemonicContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#writeAtEndOfPagePhrase.
	VisitWriteAtEndOfPagePhrase(ctx *WriteAtEndOfPagePhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#writeNotAtEndOfPagePhrase.
	VisitWriteNotAtEndOfPagePhrase(ctx *WriteNotAtEndOfPagePhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#atEndPhrase.
	VisitAtEndPhrase(ctx *AtEndPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#notAtEndPhrase.
	VisitNotAtEndPhrase(ctx *NotAtEndPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#invalidKeyPhrase.
	VisitInvalidKeyPhrase(ctx *InvalidKeyPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#notInvalidKeyPhrase.
	VisitNotInvalidKeyPhrase(ctx *NotInvalidKeyPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#onOverflowPhrase.
	VisitOnOverflowPhrase(ctx *OnOverflowPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#notOnOverflowPhrase.
	VisitNotOnOverflowPhrase(ctx *NotOnOverflowPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#onSizeErrorPhrase.
	VisitOnSizeErrorPhrase(ctx *OnSizeErrorPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#notOnSizeErrorPhrase.
	VisitNotOnSizeErrorPhrase(ctx *NotOnSizeErrorPhraseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#onExceptionClause.
	VisitOnExceptionClause(ctx *OnExceptionClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#notOnExceptionClause.
	VisitNotOnExceptionClause(ctx *NotOnExceptionClauseContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#arithmeticExpression.
	VisitArithmeticExpression(ctx *ArithmeticExpressionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#plusMinus.
	VisitPlusMinus(ctx *PlusMinusContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#multDivs.
	VisitMultDivs(ctx *MultDivsContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#multDiv.
	VisitMultDiv(ctx *MultDivContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#powers.
	VisitPowers(ctx *PowersContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#basis.
	VisitBasis(ctx *BasisContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#andOrCondition.
	VisitAndOrCondition(ctx *AndOrConditionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#combinableCondition.
	VisitCombinableCondition(ctx *CombinableConditionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#simpleCondition.
	VisitSimpleCondition(ctx *SimpleConditionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#classCondition.
	VisitClassCondition(ctx *ClassConditionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#conditionNameReference.
	VisitConditionNameReference(ctx *ConditionNameReferenceContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#conditionNameSubscriptReference.
	VisitConditionNameSubscriptReference(ctx *ConditionNameSubscriptReferenceContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#relationCondition.
	VisitRelationCondition(ctx *RelationConditionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#relationSignCondition.
	VisitRelationSignCondition(ctx *RelationSignConditionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#relationArithmeticComparison.
	VisitRelationArithmeticComparison(ctx *RelationArithmeticComparisonContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#relationCombinedComparison.
	VisitRelationCombinedComparison(ctx *RelationCombinedComparisonContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#relationCombinedCondition.
	VisitRelationCombinedCondition(ctx *RelationCombinedConditionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#relationalOperator.
	VisitRelationalOperator(ctx *RelationalOperatorContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#abbreviation.
	VisitAbbreviation(ctx *AbbreviationContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#tableCall.
	VisitTableCall(ctx *TableCallContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#referenceModifier.
	VisitReferenceModifier(ctx *ReferenceModifierContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#characterPosition.
	VisitCharacterPosition(ctx *CharacterPositionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#length.
	VisitLength(ctx *LengthContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#subscript.
	VisitSubscript(ctx *SubscriptContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#qualifiedDataName.
	VisitQualifiedDataName(ctx *QualifiedDataNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#qualifiedDataNameFormat1.
	VisitQualifiedDataNameFormat1(ctx *QualifiedDataNameFormat1Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#qualifiedDataNameFormat2.
	VisitQualifiedDataNameFormat2(ctx *QualifiedDataNameFormat2Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#qualifiedDataNameFormat3.
	VisitQualifiedDataNameFormat3(ctx *QualifiedDataNameFormat3Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#qualifiedDataNameFormat4.
	VisitQualifiedDataNameFormat4(ctx *QualifiedDataNameFormat4Context) interface{}

	// Visit a parse tree produced by Cobol85Parser#qualifiedInData.
	VisitQualifiedInData(ctx *QualifiedInDataContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inData.
	VisitInData(ctx *InDataContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inFile.
	VisitInFile(ctx *InFileContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inMnemonic.
	VisitInMnemonic(ctx *InMnemonicContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inSection.
	VisitInSection(ctx *InSectionContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inLibrary.
	VisitInLibrary(ctx *InLibraryContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#inTable.
	VisitInTable(ctx *InTableContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#alphabetName.
	VisitAlphabetName(ctx *AlphabetNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#assignmentName.
	VisitAssignmentName(ctx *AssignmentNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#basisName.
	VisitBasisName(ctx *BasisNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#cdName.
	VisitCdName(ctx *CdNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#className.
	VisitClassName(ctx *ClassNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#computerName.
	VisitComputerName(ctx *ComputerNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#conditionName.
	VisitConditionName(ctx *ConditionNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataName.
	VisitDataName(ctx *DataNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#dataDescName.
	VisitDataDescName(ctx *DataDescNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#environmentName.
	VisitEnvironmentName(ctx *EnvironmentNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#fileName.
	VisitFileName(ctx *FileNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#languageName.
	VisitLanguageName(ctx *LanguageNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#libraryName.
	VisitLibraryName(ctx *LibraryNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#localName.
	VisitLocalName(ctx *LocalNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#mnemonicName.
	VisitMnemonicName(ctx *MnemonicNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#paragraphName.
	VisitParagraphName(ctx *ParagraphNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#procedureName.
	VisitProcedureName(ctx *ProcedureNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#programName.
	VisitProgramName(ctx *ProgramNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#recordName.
	VisitRecordName(ctx *RecordNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#reportName.
	VisitReportName(ctx *ReportNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#routineName.
	VisitRoutineName(ctx *RoutineNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#screenName.
	VisitScreenName(ctx *ScreenNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#sectionName.
	VisitSectionName(ctx *SectionNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#systemName.
	VisitSystemName(ctx *SystemNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#symbolicCharacter.
	VisitSymbolicCharacter(ctx *SymbolicCharacterContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#textName.
	VisitTextName(ctx *TextNameContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#cobolWord.
	VisitCobolWord(ctx *CobolWordContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#cicsDfhRespLiteral.
	VisitCicsDfhRespLiteral(ctx *CicsDfhRespLiteralContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#cicsDfhValueLiteral.
	VisitCicsDfhValueLiteral(ctx *CicsDfhValueLiteralContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#figurativeConstant.
	VisitFigurativeConstant(ctx *FigurativeConstantContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#specialRegister.
	VisitSpecialRegister(ctx *SpecialRegisterContext) interface{}

	// Visit a parse tree produced by Cobol85Parser#commentEntry.
	VisitCommentEntry(ctx *CommentEntryContext) interface{}
}
