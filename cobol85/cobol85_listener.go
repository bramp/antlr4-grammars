// Code generated from Cobol85.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cobol85 // Cobol85
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Cobol85Listener is a complete listener for a parse tree produced by Cobol85Parser.
type Cobol85Listener interface {
	antlr.ParseTreeListener

	// EnterStartRule is called when entering the startRule production.
	EnterStartRule(c *StartRuleContext)

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterProgramUnit is called when entering the programUnit production.
	EnterProgramUnit(c *ProgramUnitContext)

	// EnterEndProgramStatement is called when entering the endProgramStatement production.
	EnterEndProgramStatement(c *EndProgramStatementContext)

	// EnterIdentificationDivision is called when entering the identificationDivision production.
	EnterIdentificationDivision(c *IdentificationDivisionContext)

	// EnterIdentificationDivisionBody is called when entering the identificationDivisionBody production.
	EnterIdentificationDivisionBody(c *IdentificationDivisionBodyContext)

	// EnterProgramIdParagraph is called when entering the programIdParagraph production.
	EnterProgramIdParagraph(c *ProgramIdParagraphContext)

	// EnterAuthorParagraph is called when entering the authorParagraph production.
	EnterAuthorParagraph(c *AuthorParagraphContext)

	// EnterInstallationParagraph is called when entering the installationParagraph production.
	EnterInstallationParagraph(c *InstallationParagraphContext)

	// EnterDateWrittenParagraph is called when entering the dateWrittenParagraph production.
	EnterDateWrittenParagraph(c *DateWrittenParagraphContext)

	// EnterDateCompiledParagraph is called when entering the dateCompiledParagraph production.
	EnterDateCompiledParagraph(c *DateCompiledParagraphContext)

	// EnterSecurityParagraph is called when entering the securityParagraph production.
	EnterSecurityParagraph(c *SecurityParagraphContext)

	// EnterRemarksParagraph is called when entering the remarksParagraph production.
	EnterRemarksParagraph(c *RemarksParagraphContext)

	// EnterEnvironmentDivision is called when entering the environmentDivision production.
	EnterEnvironmentDivision(c *EnvironmentDivisionContext)

	// EnterEnvironmentDivisionBody is called when entering the environmentDivisionBody production.
	EnterEnvironmentDivisionBody(c *EnvironmentDivisionBodyContext)

	// EnterConfigurationSection is called when entering the configurationSection production.
	EnterConfigurationSection(c *ConfigurationSectionContext)

	// EnterConfigurationSectionParagraph is called when entering the configurationSectionParagraph production.
	EnterConfigurationSectionParagraph(c *ConfigurationSectionParagraphContext)

	// EnterSourceComputerParagraph is called when entering the sourceComputerParagraph production.
	EnterSourceComputerParagraph(c *SourceComputerParagraphContext)

	// EnterObjectComputerParagraph is called when entering the objectComputerParagraph production.
	EnterObjectComputerParagraph(c *ObjectComputerParagraphContext)

	// EnterObjectComputerClause is called when entering the objectComputerClause production.
	EnterObjectComputerClause(c *ObjectComputerClauseContext)

	// EnterMemorySizeClause is called when entering the memorySizeClause production.
	EnterMemorySizeClause(c *MemorySizeClauseContext)

	// EnterDiskSizeClause is called when entering the diskSizeClause production.
	EnterDiskSizeClause(c *DiskSizeClauseContext)

	// EnterCollatingSequenceClause is called when entering the collatingSequenceClause production.
	EnterCollatingSequenceClause(c *CollatingSequenceClauseContext)

	// EnterCollatingSequenceClauseAlphanumeric is called when entering the collatingSequenceClauseAlphanumeric production.
	EnterCollatingSequenceClauseAlphanumeric(c *CollatingSequenceClauseAlphanumericContext)

	// EnterCollatingSequenceClauseNational is called when entering the collatingSequenceClauseNational production.
	EnterCollatingSequenceClauseNational(c *CollatingSequenceClauseNationalContext)

	// EnterSegmentLimitClause is called when entering the segmentLimitClause production.
	EnterSegmentLimitClause(c *SegmentLimitClauseContext)

	// EnterCharacterSetClause is called when entering the characterSetClause production.
	EnterCharacterSetClause(c *CharacterSetClauseContext)

	// EnterSpecialNamesParagraph is called when entering the specialNamesParagraph production.
	EnterSpecialNamesParagraph(c *SpecialNamesParagraphContext)

	// EnterSpecialNameClause is called when entering the specialNameClause production.
	EnterSpecialNameClause(c *SpecialNameClauseContext)

	// EnterAlphabetClause is called when entering the alphabetClause production.
	EnterAlphabetClause(c *AlphabetClauseContext)

	// EnterAlphabetClauseFormat1 is called when entering the alphabetClauseFormat1 production.
	EnterAlphabetClauseFormat1(c *AlphabetClauseFormat1Context)

	// EnterAlphabetLiterals is called when entering the alphabetLiterals production.
	EnterAlphabetLiterals(c *AlphabetLiteralsContext)

	// EnterAlphabetThrough is called when entering the alphabetThrough production.
	EnterAlphabetThrough(c *AlphabetThroughContext)

	// EnterAlphabetAlso is called when entering the alphabetAlso production.
	EnterAlphabetAlso(c *AlphabetAlsoContext)

	// EnterAlphabetClauseFormat2 is called when entering the alphabetClauseFormat2 production.
	EnterAlphabetClauseFormat2(c *AlphabetClauseFormat2Context)

	// EnterChannelClause is called when entering the channelClause production.
	EnterChannelClause(c *ChannelClauseContext)

	// EnterClassClause is called when entering the classClause production.
	EnterClassClause(c *ClassClauseContext)

	// EnterClassClauseThrough is called when entering the classClauseThrough production.
	EnterClassClauseThrough(c *ClassClauseThroughContext)

	// EnterClassClauseFrom is called when entering the classClauseFrom production.
	EnterClassClauseFrom(c *ClassClauseFromContext)

	// EnterClassClauseTo is called when entering the classClauseTo production.
	EnterClassClauseTo(c *ClassClauseToContext)

	// EnterCurrencySignClause is called when entering the currencySignClause production.
	EnterCurrencySignClause(c *CurrencySignClauseContext)

	// EnterDecimalPointClause is called when entering the decimalPointClause production.
	EnterDecimalPointClause(c *DecimalPointClauseContext)

	// EnterDefaultComputationalSignClause is called when entering the defaultComputationalSignClause production.
	EnterDefaultComputationalSignClause(c *DefaultComputationalSignClauseContext)

	// EnterDefaultDisplaySignClause is called when entering the defaultDisplaySignClause production.
	EnterDefaultDisplaySignClause(c *DefaultDisplaySignClauseContext)

	// EnterEnvironmentSwitchNameClause is called when entering the environmentSwitchNameClause production.
	EnterEnvironmentSwitchNameClause(c *EnvironmentSwitchNameClauseContext)

	// EnterEnvironmentSwitchNameSpecialNamesStatusPhrase is called when entering the environmentSwitchNameSpecialNamesStatusPhrase production.
	EnterEnvironmentSwitchNameSpecialNamesStatusPhrase(c *EnvironmentSwitchNameSpecialNamesStatusPhraseContext)

	// EnterOdtClause is called when entering the odtClause production.
	EnterOdtClause(c *OdtClauseContext)

	// EnterReserveNetworkClause is called when entering the reserveNetworkClause production.
	EnterReserveNetworkClause(c *ReserveNetworkClauseContext)

	// EnterSymbolicCharactersClause is called when entering the symbolicCharactersClause production.
	EnterSymbolicCharactersClause(c *SymbolicCharactersClauseContext)

	// EnterSymbolicCharacters is called when entering the symbolicCharacters production.
	EnterSymbolicCharacters(c *SymbolicCharactersContext)

	// EnterInputOutputSection is called when entering the inputOutputSection production.
	EnterInputOutputSection(c *InputOutputSectionContext)

	// EnterInputOutputSectionParagraph is called when entering the inputOutputSectionParagraph production.
	EnterInputOutputSectionParagraph(c *InputOutputSectionParagraphContext)

	// EnterFileControlParagraph is called when entering the fileControlParagraph production.
	EnterFileControlParagraph(c *FileControlParagraphContext)

	// EnterFileControlEntry is called when entering the fileControlEntry production.
	EnterFileControlEntry(c *FileControlEntryContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterFileControlClause is called when entering the fileControlClause production.
	EnterFileControlClause(c *FileControlClauseContext)

	// EnterAssignClause is called when entering the assignClause production.
	EnterAssignClause(c *AssignClauseContext)

	// EnterReserveClause is called when entering the reserveClause production.
	EnterReserveClause(c *ReserveClauseContext)

	// EnterOrganizationClause is called when entering the organizationClause production.
	EnterOrganizationClause(c *OrganizationClauseContext)

	// EnterPaddingCharacterClause is called when entering the paddingCharacterClause production.
	EnterPaddingCharacterClause(c *PaddingCharacterClauseContext)

	// EnterRecordDelimiterClause is called when entering the recordDelimiterClause production.
	EnterRecordDelimiterClause(c *RecordDelimiterClauseContext)

	// EnterAccessModeClause is called when entering the accessModeClause production.
	EnterAccessModeClause(c *AccessModeClauseContext)

	// EnterRecordKeyClause is called when entering the recordKeyClause production.
	EnterRecordKeyClause(c *RecordKeyClauseContext)

	// EnterAlternateRecordKeyClause is called when entering the alternateRecordKeyClause production.
	EnterAlternateRecordKeyClause(c *AlternateRecordKeyClauseContext)

	// EnterPasswordClause is called when entering the passwordClause production.
	EnterPasswordClause(c *PasswordClauseContext)

	// EnterFileStatusClause is called when entering the fileStatusClause production.
	EnterFileStatusClause(c *FileStatusClauseContext)

	// EnterRelativeKeyClause is called when entering the relativeKeyClause production.
	EnterRelativeKeyClause(c *RelativeKeyClauseContext)

	// EnterIoControlParagraph is called when entering the ioControlParagraph production.
	EnterIoControlParagraph(c *IoControlParagraphContext)

	// EnterIoControlClause is called when entering the ioControlClause production.
	EnterIoControlClause(c *IoControlClauseContext)

	// EnterRerunClause is called when entering the rerunClause production.
	EnterRerunClause(c *RerunClauseContext)

	// EnterRerunEveryRecords is called when entering the rerunEveryRecords production.
	EnterRerunEveryRecords(c *RerunEveryRecordsContext)

	// EnterRerunEveryOf is called when entering the rerunEveryOf production.
	EnterRerunEveryOf(c *RerunEveryOfContext)

	// EnterRerunEveryClock is called when entering the rerunEveryClock production.
	EnterRerunEveryClock(c *RerunEveryClockContext)

	// EnterSameClause is called when entering the sameClause production.
	EnterSameClause(c *SameClauseContext)

	// EnterMultipleFileClause is called when entering the multipleFileClause production.
	EnterMultipleFileClause(c *MultipleFileClauseContext)

	// EnterMultipleFilePosition is called when entering the multipleFilePosition production.
	EnterMultipleFilePosition(c *MultipleFilePositionContext)

	// EnterCommitmentControlClause is called when entering the commitmentControlClause production.
	EnterCommitmentControlClause(c *CommitmentControlClauseContext)

	// EnterDataDivision is called when entering the dataDivision production.
	EnterDataDivision(c *DataDivisionContext)

	// EnterDataDivisionSection is called when entering the dataDivisionSection production.
	EnterDataDivisionSection(c *DataDivisionSectionContext)

	// EnterFileSection is called when entering the fileSection production.
	EnterFileSection(c *FileSectionContext)

	// EnterFileDescriptionEntry is called when entering the fileDescriptionEntry production.
	EnterFileDescriptionEntry(c *FileDescriptionEntryContext)

	// EnterFileDescriptionEntryClause is called when entering the fileDescriptionEntryClause production.
	EnterFileDescriptionEntryClause(c *FileDescriptionEntryClauseContext)

	// EnterExternalClause is called when entering the externalClause production.
	EnterExternalClause(c *ExternalClauseContext)

	// EnterGlobalClause is called when entering the globalClause production.
	EnterGlobalClause(c *GlobalClauseContext)

	// EnterBlockContainsClause is called when entering the blockContainsClause production.
	EnterBlockContainsClause(c *BlockContainsClauseContext)

	// EnterBlockContainsTo is called when entering the blockContainsTo production.
	EnterBlockContainsTo(c *BlockContainsToContext)

	// EnterRecordContainsClause is called when entering the recordContainsClause production.
	EnterRecordContainsClause(c *RecordContainsClauseContext)

	// EnterRecordContainsClauseFormat1 is called when entering the recordContainsClauseFormat1 production.
	EnterRecordContainsClauseFormat1(c *RecordContainsClauseFormat1Context)

	// EnterRecordContainsClauseFormat2 is called when entering the recordContainsClauseFormat2 production.
	EnterRecordContainsClauseFormat2(c *RecordContainsClauseFormat2Context)

	// EnterRecordContainsClauseFormat3 is called when entering the recordContainsClauseFormat3 production.
	EnterRecordContainsClauseFormat3(c *RecordContainsClauseFormat3Context)

	// EnterRecordContainsTo is called when entering the recordContainsTo production.
	EnterRecordContainsTo(c *RecordContainsToContext)

	// EnterLabelRecordsClause is called when entering the labelRecordsClause production.
	EnterLabelRecordsClause(c *LabelRecordsClauseContext)

	// EnterValueOfClause is called when entering the valueOfClause production.
	EnterValueOfClause(c *ValueOfClauseContext)

	// EnterValuePair is called when entering the valuePair production.
	EnterValuePair(c *ValuePairContext)

	// EnterDataRecordsClause is called when entering the dataRecordsClause production.
	EnterDataRecordsClause(c *DataRecordsClauseContext)

	// EnterLinageClause is called when entering the linageClause production.
	EnterLinageClause(c *LinageClauseContext)

	// EnterLinageAt is called when entering the linageAt production.
	EnterLinageAt(c *LinageAtContext)

	// EnterLinageFootingAt is called when entering the linageFootingAt production.
	EnterLinageFootingAt(c *LinageFootingAtContext)

	// EnterLinageLinesAtTop is called when entering the linageLinesAtTop production.
	EnterLinageLinesAtTop(c *LinageLinesAtTopContext)

	// EnterLinageLinesAtBottom is called when entering the linageLinesAtBottom production.
	EnterLinageLinesAtBottom(c *LinageLinesAtBottomContext)

	// EnterRecordingModeClause is called when entering the recordingModeClause production.
	EnterRecordingModeClause(c *RecordingModeClauseContext)

	// EnterModeStatement is called when entering the modeStatement production.
	EnterModeStatement(c *ModeStatementContext)

	// EnterCodeSetClause is called when entering the codeSetClause production.
	EnterCodeSetClause(c *CodeSetClauseContext)

	// EnterReportClause is called when entering the reportClause production.
	EnterReportClause(c *ReportClauseContext)

	// EnterDataBaseSection is called when entering the dataBaseSection production.
	EnterDataBaseSection(c *DataBaseSectionContext)

	// EnterDataBaseSectionEntry is called when entering the dataBaseSectionEntry production.
	EnterDataBaseSectionEntry(c *DataBaseSectionEntryContext)

	// EnterWorkingStorageSection is called when entering the workingStorageSection production.
	EnterWorkingStorageSection(c *WorkingStorageSectionContext)

	// EnterLinkageSection is called when entering the linkageSection production.
	EnterLinkageSection(c *LinkageSectionContext)

	// EnterCommunicationSection is called when entering the communicationSection production.
	EnterCommunicationSection(c *CommunicationSectionContext)

	// EnterCommunicationDescriptionEntry is called when entering the communicationDescriptionEntry production.
	EnterCommunicationDescriptionEntry(c *CommunicationDescriptionEntryContext)

	// EnterCommunicationDescriptionEntryFormat1 is called when entering the communicationDescriptionEntryFormat1 production.
	EnterCommunicationDescriptionEntryFormat1(c *CommunicationDescriptionEntryFormat1Context)

	// EnterCommunicationDescriptionEntryFormat2 is called when entering the communicationDescriptionEntryFormat2 production.
	EnterCommunicationDescriptionEntryFormat2(c *CommunicationDescriptionEntryFormat2Context)

	// EnterCommunicationDescriptionEntryFormat3 is called when entering the communicationDescriptionEntryFormat3 production.
	EnterCommunicationDescriptionEntryFormat3(c *CommunicationDescriptionEntryFormat3Context)

	// EnterDestinationCountClause is called when entering the destinationCountClause production.
	EnterDestinationCountClause(c *DestinationCountClauseContext)

	// EnterDestinationTableClause is called when entering the destinationTableClause production.
	EnterDestinationTableClause(c *DestinationTableClauseContext)

	// EnterEndKeyClause is called when entering the endKeyClause production.
	EnterEndKeyClause(c *EndKeyClauseContext)

	// EnterErrorKeyClause is called when entering the errorKeyClause production.
	EnterErrorKeyClause(c *ErrorKeyClauseContext)

	// EnterMessageCountClause is called when entering the messageCountClause production.
	EnterMessageCountClause(c *MessageCountClauseContext)

	// EnterMessageDateClause is called when entering the messageDateClause production.
	EnterMessageDateClause(c *MessageDateClauseContext)

	// EnterMessageTimeClause is called when entering the messageTimeClause production.
	EnterMessageTimeClause(c *MessageTimeClauseContext)

	// EnterStatusKeyClause is called when entering the statusKeyClause production.
	EnterStatusKeyClause(c *StatusKeyClauseContext)

	// EnterSymbolicDestinationClause is called when entering the symbolicDestinationClause production.
	EnterSymbolicDestinationClause(c *SymbolicDestinationClauseContext)

	// EnterSymbolicQueueClause is called when entering the symbolicQueueClause production.
	EnterSymbolicQueueClause(c *SymbolicQueueClauseContext)

	// EnterSymbolicSourceClause is called when entering the symbolicSourceClause production.
	EnterSymbolicSourceClause(c *SymbolicSourceClauseContext)

	// EnterSymbolicTerminalClause is called when entering the symbolicTerminalClause production.
	EnterSymbolicTerminalClause(c *SymbolicTerminalClauseContext)

	// EnterSymbolicSubQueueClause is called when entering the symbolicSubQueueClause production.
	EnterSymbolicSubQueueClause(c *SymbolicSubQueueClauseContext)

	// EnterTextLengthClause is called when entering the textLengthClause production.
	EnterTextLengthClause(c *TextLengthClauseContext)

	// EnterLocalStorageSection is called when entering the localStorageSection production.
	EnterLocalStorageSection(c *LocalStorageSectionContext)

	// EnterScreenSection is called when entering the screenSection production.
	EnterScreenSection(c *ScreenSectionContext)

	// EnterScreenDescriptionEntry is called when entering the screenDescriptionEntry production.
	EnterScreenDescriptionEntry(c *ScreenDescriptionEntryContext)

	// EnterScreenDescriptionBlankClause is called when entering the screenDescriptionBlankClause production.
	EnterScreenDescriptionBlankClause(c *ScreenDescriptionBlankClauseContext)

	// EnterScreenDescriptionBellClause is called when entering the screenDescriptionBellClause production.
	EnterScreenDescriptionBellClause(c *ScreenDescriptionBellClauseContext)

	// EnterScreenDescriptionBlinkClause is called when entering the screenDescriptionBlinkClause production.
	EnterScreenDescriptionBlinkClause(c *ScreenDescriptionBlinkClauseContext)

	// EnterScreenDescriptionEraseClause is called when entering the screenDescriptionEraseClause production.
	EnterScreenDescriptionEraseClause(c *ScreenDescriptionEraseClauseContext)

	// EnterScreenDescriptionLightClause is called when entering the screenDescriptionLightClause production.
	EnterScreenDescriptionLightClause(c *ScreenDescriptionLightClauseContext)

	// EnterScreenDescriptionGridClause is called when entering the screenDescriptionGridClause production.
	EnterScreenDescriptionGridClause(c *ScreenDescriptionGridClauseContext)

	// EnterScreenDescriptionReverseVideoClause is called when entering the screenDescriptionReverseVideoClause production.
	EnterScreenDescriptionReverseVideoClause(c *ScreenDescriptionReverseVideoClauseContext)

	// EnterScreenDescriptionUnderlineClause is called when entering the screenDescriptionUnderlineClause production.
	EnterScreenDescriptionUnderlineClause(c *ScreenDescriptionUnderlineClauseContext)

	// EnterScreenDescriptionSizeClause is called when entering the screenDescriptionSizeClause production.
	EnterScreenDescriptionSizeClause(c *ScreenDescriptionSizeClauseContext)

	// EnterScreenDescriptionLineClause is called when entering the screenDescriptionLineClause production.
	EnterScreenDescriptionLineClause(c *ScreenDescriptionLineClauseContext)

	// EnterScreenDescriptionColumnClause is called when entering the screenDescriptionColumnClause production.
	EnterScreenDescriptionColumnClause(c *ScreenDescriptionColumnClauseContext)

	// EnterScreenDescriptionForegroundColorClause is called when entering the screenDescriptionForegroundColorClause production.
	EnterScreenDescriptionForegroundColorClause(c *ScreenDescriptionForegroundColorClauseContext)

	// EnterScreenDescriptionBackgroundColorClause is called when entering the screenDescriptionBackgroundColorClause production.
	EnterScreenDescriptionBackgroundColorClause(c *ScreenDescriptionBackgroundColorClauseContext)

	// EnterScreenDescriptionControlClause is called when entering the screenDescriptionControlClause production.
	EnterScreenDescriptionControlClause(c *ScreenDescriptionControlClauseContext)

	// EnterScreenDescriptionValueClause is called when entering the screenDescriptionValueClause production.
	EnterScreenDescriptionValueClause(c *ScreenDescriptionValueClauseContext)

	// EnterScreenDescriptionPictureClause is called when entering the screenDescriptionPictureClause production.
	EnterScreenDescriptionPictureClause(c *ScreenDescriptionPictureClauseContext)

	// EnterScreenDescriptionFromClause is called when entering the screenDescriptionFromClause production.
	EnterScreenDescriptionFromClause(c *ScreenDescriptionFromClauseContext)

	// EnterScreenDescriptionToClause is called when entering the screenDescriptionToClause production.
	EnterScreenDescriptionToClause(c *ScreenDescriptionToClauseContext)

	// EnterScreenDescriptionUsingClause is called when entering the screenDescriptionUsingClause production.
	EnterScreenDescriptionUsingClause(c *ScreenDescriptionUsingClauseContext)

	// EnterScreenDescriptionUsageClause is called when entering the screenDescriptionUsageClause production.
	EnterScreenDescriptionUsageClause(c *ScreenDescriptionUsageClauseContext)

	// EnterScreenDescriptionBlankWhenZeroClause is called when entering the screenDescriptionBlankWhenZeroClause production.
	EnterScreenDescriptionBlankWhenZeroClause(c *ScreenDescriptionBlankWhenZeroClauseContext)

	// EnterScreenDescriptionJustifiedClause is called when entering the screenDescriptionJustifiedClause production.
	EnterScreenDescriptionJustifiedClause(c *ScreenDescriptionJustifiedClauseContext)

	// EnterScreenDescriptionSignClause is called when entering the screenDescriptionSignClause production.
	EnterScreenDescriptionSignClause(c *ScreenDescriptionSignClauseContext)

	// EnterScreenDescriptionAutoClause is called when entering the screenDescriptionAutoClause production.
	EnterScreenDescriptionAutoClause(c *ScreenDescriptionAutoClauseContext)

	// EnterScreenDescriptionSecureClause is called when entering the screenDescriptionSecureClause production.
	EnterScreenDescriptionSecureClause(c *ScreenDescriptionSecureClauseContext)

	// EnterScreenDescriptionRequiredClause is called when entering the screenDescriptionRequiredClause production.
	EnterScreenDescriptionRequiredClause(c *ScreenDescriptionRequiredClauseContext)

	// EnterScreenDescriptionPromptClause is called when entering the screenDescriptionPromptClause production.
	EnterScreenDescriptionPromptClause(c *ScreenDescriptionPromptClauseContext)

	// EnterScreenDescriptionPromptOccursClause is called when entering the screenDescriptionPromptOccursClause production.
	EnterScreenDescriptionPromptOccursClause(c *ScreenDescriptionPromptOccursClauseContext)

	// EnterScreenDescriptionFullClause is called when entering the screenDescriptionFullClause production.
	EnterScreenDescriptionFullClause(c *ScreenDescriptionFullClauseContext)

	// EnterScreenDescriptionZeroFillClause is called when entering the screenDescriptionZeroFillClause production.
	EnterScreenDescriptionZeroFillClause(c *ScreenDescriptionZeroFillClauseContext)

	// EnterReportSection is called when entering the reportSection production.
	EnterReportSection(c *ReportSectionContext)

	// EnterReportDescription is called when entering the reportDescription production.
	EnterReportDescription(c *ReportDescriptionContext)

	// EnterReportDescriptionEntry is called when entering the reportDescriptionEntry production.
	EnterReportDescriptionEntry(c *ReportDescriptionEntryContext)

	// EnterReportDescriptionGlobalClause is called when entering the reportDescriptionGlobalClause production.
	EnterReportDescriptionGlobalClause(c *ReportDescriptionGlobalClauseContext)

	// EnterReportDescriptionPageLimitClause is called when entering the reportDescriptionPageLimitClause production.
	EnterReportDescriptionPageLimitClause(c *ReportDescriptionPageLimitClauseContext)

	// EnterReportDescriptionHeadingClause is called when entering the reportDescriptionHeadingClause production.
	EnterReportDescriptionHeadingClause(c *ReportDescriptionHeadingClauseContext)

	// EnterReportDescriptionFirstDetailClause is called when entering the reportDescriptionFirstDetailClause production.
	EnterReportDescriptionFirstDetailClause(c *ReportDescriptionFirstDetailClauseContext)

	// EnterReportDescriptionLastDetailClause is called when entering the reportDescriptionLastDetailClause production.
	EnterReportDescriptionLastDetailClause(c *ReportDescriptionLastDetailClauseContext)

	// EnterReportDescriptionFootingClause is called when entering the reportDescriptionFootingClause production.
	EnterReportDescriptionFootingClause(c *ReportDescriptionFootingClauseContext)

	// EnterReportGroupDescriptionEntry is called when entering the reportGroupDescriptionEntry production.
	EnterReportGroupDescriptionEntry(c *ReportGroupDescriptionEntryContext)

	// EnterReportGroupDescriptionEntryFormat1 is called when entering the reportGroupDescriptionEntryFormat1 production.
	EnterReportGroupDescriptionEntryFormat1(c *ReportGroupDescriptionEntryFormat1Context)

	// EnterReportGroupDescriptionEntryFormat2 is called when entering the reportGroupDescriptionEntryFormat2 production.
	EnterReportGroupDescriptionEntryFormat2(c *ReportGroupDescriptionEntryFormat2Context)

	// EnterReportGroupDescriptionEntryFormat3 is called when entering the reportGroupDescriptionEntryFormat3 production.
	EnterReportGroupDescriptionEntryFormat3(c *ReportGroupDescriptionEntryFormat3Context)

	// EnterReportGroupBlankWhenZeroClause is called when entering the reportGroupBlankWhenZeroClause production.
	EnterReportGroupBlankWhenZeroClause(c *ReportGroupBlankWhenZeroClauseContext)

	// EnterReportGroupColumnNumberClause is called when entering the reportGroupColumnNumberClause production.
	EnterReportGroupColumnNumberClause(c *ReportGroupColumnNumberClauseContext)

	// EnterReportGroupIndicateClause is called when entering the reportGroupIndicateClause production.
	EnterReportGroupIndicateClause(c *ReportGroupIndicateClauseContext)

	// EnterReportGroupJustifiedClause is called when entering the reportGroupJustifiedClause production.
	EnterReportGroupJustifiedClause(c *ReportGroupJustifiedClauseContext)

	// EnterReportGroupLineNumberClause is called when entering the reportGroupLineNumberClause production.
	EnterReportGroupLineNumberClause(c *ReportGroupLineNumberClauseContext)

	// EnterReportGroupLineNumberNextPage is called when entering the reportGroupLineNumberNextPage production.
	EnterReportGroupLineNumberNextPage(c *ReportGroupLineNumberNextPageContext)

	// EnterReportGroupLineNumberPlus is called when entering the reportGroupLineNumberPlus production.
	EnterReportGroupLineNumberPlus(c *ReportGroupLineNumberPlusContext)

	// EnterReportGroupNextGroupClause is called when entering the reportGroupNextGroupClause production.
	EnterReportGroupNextGroupClause(c *ReportGroupNextGroupClauseContext)

	// EnterReportGroupNextGroupPlus is called when entering the reportGroupNextGroupPlus production.
	EnterReportGroupNextGroupPlus(c *ReportGroupNextGroupPlusContext)

	// EnterReportGroupNextGroupNextPage is called when entering the reportGroupNextGroupNextPage production.
	EnterReportGroupNextGroupNextPage(c *ReportGroupNextGroupNextPageContext)

	// EnterReportGroupPictureClause is called when entering the reportGroupPictureClause production.
	EnterReportGroupPictureClause(c *ReportGroupPictureClauseContext)

	// EnterReportGroupResetClause is called when entering the reportGroupResetClause production.
	EnterReportGroupResetClause(c *ReportGroupResetClauseContext)

	// EnterReportGroupSignClause is called when entering the reportGroupSignClause production.
	EnterReportGroupSignClause(c *ReportGroupSignClauseContext)

	// EnterReportGroupSourceClause is called when entering the reportGroupSourceClause production.
	EnterReportGroupSourceClause(c *ReportGroupSourceClauseContext)

	// EnterReportGroupSumClause is called when entering the reportGroupSumClause production.
	EnterReportGroupSumClause(c *ReportGroupSumClauseContext)

	// EnterReportGroupTypeClause is called when entering the reportGroupTypeClause production.
	EnterReportGroupTypeClause(c *ReportGroupTypeClauseContext)

	// EnterReportGroupTypeReportHeading is called when entering the reportGroupTypeReportHeading production.
	EnterReportGroupTypeReportHeading(c *ReportGroupTypeReportHeadingContext)

	// EnterReportGroupTypePageHeading is called when entering the reportGroupTypePageHeading production.
	EnterReportGroupTypePageHeading(c *ReportGroupTypePageHeadingContext)

	// EnterReportGroupTypeControlHeading is called when entering the reportGroupTypeControlHeading production.
	EnterReportGroupTypeControlHeading(c *ReportGroupTypeControlHeadingContext)

	// EnterReportGroupTypeDetail is called when entering the reportGroupTypeDetail production.
	EnterReportGroupTypeDetail(c *ReportGroupTypeDetailContext)

	// EnterReportGroupTypeControlFooting is called when entering the reportGroupTypeControlFooting production.
	EnterReportGroupTypeControlFooting(c *ReportGroupTypeControlFootingContext)

	// EnterReportGroupUsageClause is called when entering the reportGroupUsageClause production.
	EnterReportGroupUsageClause(c *ReportGroupUsageClauseContext)

	// EnterReportGroupTypePageFooting is called when entering the reportGroupTypePageFooting production.
	EnterReportGroupTypePageFooting(c *ReportGroupTypePageFootingContext)

	// EnterReportGroupTypeReportFooting is called when entering the reportGroupTypeReportFooting production.
	EnterReportGroupTypeReportFooting(c *ReportGroupTypeReportFootingContext)

	// EnterReportGroupValueClause is called when entering the reportGroupValueClause production.
	EnterReportGroupValueClause(c *ReportGroupValueClauseContext)

	// EnterProgramLibrarySection is called when entering the programLibrarySection production.
	EnterProgramLibrarySection(c *ProgramLibrarySectionContext)

	// EnterLibraryDescriptionEntry is called when entering the libraryDescriptionEntry production.
	EnterLibraryDescriptionEntry(c *LibraryDescriptionEntryContext)

	// EnterLibraryDescriptionEntryFormat1 is called when entering the libraryDescriptionEntryFormat1 production.
	EnterLibraryDescriptionEntryFormat1(c *LibraryDescriptionEntryFormat1Context)

	// EnterLibraryDescriptionEntryFormat2 is called when entering the libraryDescriptionEntryFormat2 production.
	EnterLibraryDescriptionEntryFormat2(c *LibraryDescriptionEntryFormat2Context)

	// EnterLibraryAttributeClauseFormat1 is called when entering the libraryAttributeClauseFormat1 production.
	EnterLibraryAttributeClauseFormat1(c *LibraryAttributeClauseFormat1Context)

	// EnterLibraryAttributeClauseFormat2 is called when entering the libraryAttributeClauseFormat2 production.
	EnterLibraryAttributeClauseFormat2(c *LibraryAttributeClauseFormat2Context)

	// EnterLibraryAttributeFunction is called when entering the libraryAttributeFunction production.
	EnterLibraryAttributeFunction(c *LibraryAttributeFunctionContext)

	// EnterLibraryAttributeParameter is called when entering the libraryAttributeParameter production.
	EnterLibraryAttributeParameter(c *LibraryAttributeParameterContext)

	// EnterLibraryAttributeTitle is called when entering the libraryAttributeTitle production.
	EnterLibraryAttributeTitle(c *LibraryAttributeTitleContext)

	// EnterLibraryEntryProcedureClauseFormat1 is called when entering the libraryEntryProcedureClauseFormat1 production.
	EnterLibraryEntryProcedureClauseFormat1(c *LibraryEntryProcedureClauseFormat1Context)

	// EnterLibraryEntryProcedureClauseFormat2 is called when entering the libraryEntryProcedureClauseFormat2 production.
	EnterLibraryEntryProcedureClauseFormat2(c *LibraryEntryProcedureClauseFormat2Context)

	// EnterLibraryEntryProcedureForClause is called when entering the libraryEntryProcedureForClause production.
	EnterLibraryEntryProcedureForClause(c *LibraryEntryProcedureForClauseContext)

	// EnterLibraryEntryProcedureGivingClause is called when entering the libraryEntryProcedureGivingClause production.
	EnterLibraryEntryProcedureGivingClause(c *LibraryEntryProcedureGivingClauseContext)

	// EnterLibraryEntryProcedureUsingClause is called when entering the libraryEntryProcedureUsingClause production.
	EnterLibraryEntryProcedureUsingClause(c *LibraryEntryProcedureUsingClauseContext)

	// EnterLibraryEntryProcedureUsingName is called when entering the libraryEntryProcedureUsingName production.
	EnterLibraryEntryProcedureUsingName(c *LibraryEntryProcedureUsingNameContext)

	// EnterLibraryEntryProcedureWithClause is called when entering the libraryEntryProcedureWithClause production.
	EnterLibraryEntryProcedureWithClause(c *LibraryEntryProcedureWithClauseContext)

	// EnterLibraryEntryProcedureWithName is called when entering the libraryEntryProcedureWithName production.
	EnterLibraryEntryProcedureWithName(c *LibraryEntryProcedureWithNameContext)

	// EnterLibraryIsCommonClause is called when entering the libraryIsCommonClause production.
	EnterLibraryIsCommonClause(c *LibraryIsCommonClauseContext)

	// EnterLibraryIsGlobalClause is called when entering the libraryIsGlobalClause production.
	EnterLibraryIsGlobalClause(c *LibraryIsGlobalClauseContext)

	// EnterDataDescriptionEntry is called when entering the dataDescriptionEntry production.
	EnterDataDescriptionEntry(c *DataDescriptionEntryContext)

	// EnterDataDescriptionEntryFormat1 is called when entering the dataDescriptionEntryFormat1 production.
	EnterDataDescriptionEntryFormat1(c *DataDescriptionEntryFormat1Context)

	// EnterDataDescriptionEntryFormat2 is called when entering the dataDescriptionEntryFormat2 production.
	EnterDataDescriptionEntryFormat2(c *DataDescriptionEntryFormat2Context)

	// EnterDataDescriptionEntryFormat3 is called when entering the dataDescriptionEntryFormat3 production.
	EnterDataDescriptionEntryFormat3(c *DataDescriptionEntryFormat3Context)

	// EnterDataDescriptionEntryExecSql is called when entering the dataDescriptionEntryExecSql production.
	EnterDataDescriptionEntryExecSql(c *DataDescriptionEntryExecSqlContext)

	// EnterDataAlignedClause is called when entering the dataAlignedClause production.
	EnterDataAlignedClause(c *DataAlignedClauseContext)

	// EnterDataBlankWhenZeroClause is called when entering the dataBlankWhenZeroClause production.
	EnterDataBlankWhenZeroClause(c *DataBlankWhenZeroClauseContext)

	// EnterDataCommonOwnLocalClause is called when entering the dataCommonOwnLocalClause production.
	EnterDataCommonOwnLocalClause(c *DataCommonOwnLocalClauseContext)

	// EnterDataExternalClause is called when entering the dataExternalClause production.
	EnterDataExternalClause(c *DataExternalClauseContext)

	// EnterDataGlobalClause is called when entering the dataGlobalClause production.
	EnterDataGlobalClause(c *DataGlobalClauseContext)

	// EnterDataIntegerStringClause is called when entering the dataIntegerStringClause production.
	EnterDataIntegerStringClause(c *DataIntegerStringClauseContext)

	// EnterDataJustifiedClause is called when entering the dataJustifiedClause production.
	EnterDataJustifiedClause(c *DataJustifiedClauseContext)

	// EnterDataOccursClause is called when entering the dataOccursClause production.
	EnterDataOccursClause(c *DataOccursClauseContext)

	// EnterDataOccursTo is called when entering the dataOccursTo production.
	EnterDataOccursTo(c *DataOccursToContext)

	// EnterDataOccursSort is called when entering the dataOccursSort production.
	EnterDataOccursSort(c *DataOccursSortContext)

	// EnterDataPictureClause is called when entering the dataPictureClause production.
	EnterDataPictureClause(c *DataPictureClauseContext)

	// EnterPictureString is called when entering the pictureString production.
	EnterPictureString(c *PictureStringContext)

	// EnterPictureChars is called when entering the pictureChars production.
	EnterPictureChars(c *PictureCharsContext)

	// EnterPictureCardinality is called when entering the pictureCardinality production.
	EnterPictureCardinality(c *PictureCardinalityContext)

	// EnterDataReceivedByClause is called when entering the dataReceivedByClause production.
	EnterDataReceivedByClause(c *DataReceivedByClauseContext)

	// EnterDataRecordAreaClause is called when entering the dataRecordAreaClause production.
	EnterDataRecordAreaClause(c *DataRecordAreaClauseContext)

	// EnterDataRedefinesClause is called when entering the dataRedefinesClause production.
	EnterDataRedefinesClause(c *DataRedefinesClauseContext)

	// EnterDataRenamesClause is called when entering the dataRenamesClause production.
	EnterDataRenamesClause(c *DataRenamesClauseContext)

	// EnterDataSignClause is called when entering the dataSignClause production.
	EnterDataSignClause(c *DataSignClauseContext)

	// EnterDataSynchronizedClause is called when entering the dataSynchronizedClause production.
	EnterDataSynchronizedClause(c *DataSynchronizedClauseContext)

	// EnterDataThreadLocalClause is called when entering the dataThreadLocalClause production.
	EnterDataThreadLocalClause(c *DataThreadLocalClauseContext)

	// EnterDataTypeClause is called when entering the dataTypeClause production.
	EnterDataTypeClause(c *DataTypeClauseContext)

	// EnterDataTypeDefClause is called when entering the dataTypeDefClause production.
	EnterDataTypeDefClause(c *DataTypeDefClauseContext)

	// EnterDataUsageClause is called when entering the dataUsageClause production.
	EnterDataUsageClause(c *DataUsageClauseContext)

	// EnterDataUsingClause is called when entering the dataUsingClause production.
	EnterDataUsingClause(c *DataUsingClauseContext)

	// EnterDataValueClause is called when entering the dataValueClause production.
	EnterDataValueClause(c *DataValueClauseContext)

	// EnterDataValueInterval is called when entering the dataValueInterval production.
	EnterDataValueInterval(c *DataValueIntervalContext)

	// EnterDataValueIntervalFrom is called when entering the dataValueIntervalFrom production.
	EnterDataValueIntervalFrom(c *DataValueIntervalFromContext)

	// EnterDataValueIntervalTo is called when entering the dataValueIntervalTo production.
	EnterDataValueIntervalTo(c *DataValueIntervalToContext)

	// EnterDataWithLowerBoundsClause is called when entering the dataWithLowerBoundsClause production.
	EnterDataWithLowerBoundsClause(c *DataWithLowerBoundsClauseContext)

	// EnterProcedureDivision is called when entering the procedureDivision production.
	EnterProcedureDivision(c *ProcedureDivisionContext)

	// EnterProcedureDivisionUsingClause is called when entering the procedureDivisionUsingClause production.
	EnterProcedureDivisionUsingClause(c *ProcedureDivisionUsingClauseContext)

	// EnterProcedureDivisionGivingClause is called when entering the procedureDivisionGivingClause production.
	EnterProcedureDivisionGivingClause(c *ProcedureDivisionGivingClauseContext)

	// EnterProcedureDivisionUsingParameter is called when entering the procedureDivisionUsingParameter production.
	EnterProcedureDivisionUsingParameter(c *ProcedureDivisionUsingParameterContext)

	// EnterProcedureDivisionByReferencePhrase is called when entering the procedureDivisionByReferencePhrase production.
	EnterProcedureDivisionByReferencePhrase(c *ProcedureDivisionByReferencePhraseContext)

	// EnterProcedureDivisionByReference is called when entering the procedureDivisionByReference production.
	EnterProcedureDivisionByReference(c *ProcedureDivisionByReferenceContext)

	// EnterProcedureDivisionByValuePhrase is called when entering the procedureDivisionByValuePhrase production.
	EnterProcedureDivisionByValuePhrase(c *ProcedureDivisionByValuePhraseContext)

	// EnterProcedureDivisionByValue is called when entering the procedureDivisionByValue production.
	EnterProcedureDivisionByValue(c *ProcedureDivisionByValueContext)

	// EnterProcedureDeclaratives is called when entering the procedureDeclaratives production.
	EnterProcedureDeclaratives(c *ProcedureDeclarativesContext)

	// EnterProcedureDeclarative is called when entering the procedureDeclarative production.
	EnterProcedureDeclarative(c *ProcedureDeclarativeContext)

	// EnterProcedureSectionHeader is called when entering the procedureSectionHeader production.
	EnterProcedureSectionHeader(c *ProcedureSectionHeaderContext)

	// EnterProcedureDivisionBody is called when entering the procedureDivisionBody production.
	EnterProcedureDivisionBody(c *ProcedureDivisionBodyContext)

	// EnterProcedureSection is called when entering the procedureSection production.
	EnterProcedureSection(c *ProcedureSectionContext)

	// EnterParagraphs is called when entering the paragraphs production.
	EnterParagraphs(c *ParagraphsContext)

	// EnterParagraph is called when entering the paragraph production.
	EnterParagraph(c *ParagraphContext)

	// EnterSentence is called when entering the sentence production.
	EnterSentence(c *SentenceContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAcceptStatement is called when entering the acceptStatement production.
	EnterAcceptStatement(c *AcceptStatementContext)

	// EnterAcceptFromDateStatement is called when entering the acceptFromDateStatement production.
	EnterAcceptFromDateStatement(c *AcceptFromDateStatementContext)

	// EnterAcceptFromMnemonicStatement is called when entering the acceptFromMnemonicStatement production.
	EnterAcceptFromMnemonicStatement(c *AcceptFromMnemonicStatementContext)

	// EnterAcceptFromEscapeKeyStatement is called when entering the acceptFromEscapeKeyStatement production.
	EnterAcceptFromEscapeKeyStatement(c *AcceptFromEscapeKeyStatementContext)

	// EnterAcceptMessageCountStatement is called when entering the acceptMessageCountStatement production.
	EnterAcceptMessageCountStatement(c *AcceptMessageCountStatementContext)

	// EnterAddStatement is called when entering the addStatement production.
	EnterAddStatement(c *AddStatementContext)

	// EnterAddToStatement is called when entering the addToStatement production.
	EnterAddToStatement(c *AddToStatementContext)

	// EnterAddToGivingStatement is called when entering the addToGivingStatement production.
	EnterAddToGivingStatement(c *AddToGivingStatementContext)

	// EnterAddCorrespondingStatement is called when entering the addCorrespondingStatement production.
	EnterAddCorrespondingStatement(c *AddCorrespondingStatementContext)

	// EnterAddFrom is called when entering the addFrom production.
	EnterAddFrom(c *AddFromContext)

	// EnterAddTo is called when entering the addTo production.
	EnterAddTo(c *AddToContext)

	// EnterAddToGiving is called when entering the addToGiving production.
	EnterAddToGiving(c *AddToGivingContext)

	// EnterAddGiving is called when entering the addGiving production.
	EnterAddGiving(c *AddGivingContext)

	// EnterAlteredGoTo is called when entering the alteredGoTo production.
	EnterAlteredGoTo(c *AlteredGoToContext)

	// EnterAlterStatement is called when entering the alterStatement production.
	EnterAlterStatement(c *AlterStatementContext)

	// EnterAlterProceedTo is called when entering the alterProceedTo production.
	EnterAlterProceedTo(c *AlterProceedToContext)

	// EnterCallStatement is called when entering the callStatement production.
	EnterCallStatement(c *CallStatementContext)

	// EnterCallUsingPhrase is called when entering the callUsingPhrase production.
	EnterCallUsingPhrase(c *CallUsingPhraseContext)

	// EnterCallUsingParameter is called when entering the callUsingParameter production.
	EnterCallUsingParameter(c *CallUsingParameterContext)

	// EnterCallByReferencePhrase is called when entering the callByReferencePhrase production.
	EnterCallByReferencePhrase(c *CallByReferencePhraseContext)

	// EnterCallByReference is called when entering the callByReference production.
	EnterCallByReference(c *CallByReferenceContext)

	// EnterCallByValuePhrase is called when entering the callByValuePhrase production.
	EnterCallByValuePhrase(c *CallByValuePhraseContext)

	// EnterCallByValue is called when entering the callByValue production.
	EnterCallByValue(c *CallByValueContext)

	// EnterCallByContentPhrase is called when entering the callByContentPhrase production.
	EnterCallByContentPhrase(c *CallByContentPhraseContext)

	// EnterCallByContent is called when entering the callByContent production.
	EnterCallByContent(c *CallByContentContext)

	// EnterCallGivingPhrase is called when entering the callGivingPhrase production.
	EnterCallGivingPhrase(c *CallGivingPhraseContext)

	// EnterCancelStatement is called when entering the cancelStatement production.
	EnterCancelStatement(c *CancelStatementContext)

	// EnterCancelCall is called when entering the cancelCall production.
	EnterCancelCall(c *CancelCallContext)

	// EnterCloseStatement is called when entering the closeStatement production.
	EnterCloseStatement(c *CloseStatementContext)

	// EnterCloseFile is called when entering the closeFile production.
	EnterCloseFile(c *CloseFileContext)

	// EnterCloseReelUnitStatement is called when entering the closeReelUnitStatement production.
	EnterCloseReelUnitStatement(c *CloseReelUnitStatementContext)

	// EnterCloseRelativeStatement is called when entering the closeRelativeStatement production.
	EnterCloseRelativeStatement(c *CloseRelativeStatementContext)

	// EnterClosePortFileIOStatement is called when entering the closePortFileIOStatement production.
	EnterClosePortFileIOStatement(c *ClosePortFileIOStatementContext)

	// EnterClosePortFileIOUsing is called when entering the closePortFileIOUsing production.
	EnterClosePortFileIOUsing(c *ClosePortFileIOUsingContext)

	// EnterClosePortFileIOUsingCloseDisposition is called when entering the closePortFileIOUsingCloseDisposition production.
	EnterClosePortFileIOUsingCloseDisposition(c *ClosePortFileIOUsingCloseDispositionContext)

	// EnterClosePortFileIOUsingAssociatedData is called when entering the closePortFileIOUsingAssociatedData production.
	EnterClosePortFileIOUsingAssociatedData(c *ClosePortFileIOUsingAssociatedDataContext)

	// EnterClosePortFileIOUsingAssociatedDataLength is called when entering the closePortFileIOUsingAssociatedDataLength production.
	EnterClosePortFileIOUsingAssociatedDataLength(c *ClosePortFileIOUsingAssociatedDataLengthContext)

	// EnterComputeStatement is called when entering the computeStatement production.
	EnterComputeStatement(c *ComputeStatementContext)

	// EnterComputeStore is called when entering the computeStore production.
	EnterComputeStore(c *ComputeStoreContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterDisableStatement is called when entering the disableStatement production.
	EnterDisableStatement(c *DisableStatementContext)

	// EnterDisplayStatement is called when entering the displayStatement production.
	EnterDisplayStatement(c *DisplayStatementContext)

	// EnterDisplayOperand is called when entering the displayOperand production.
	EnterDisplayOperand(c *DisplayOperandContext)

	// EnterDisplayAt is called when entering the displayAt production.
	EnterDisplayAt(c *DisplayAtContext)

	// EnterDisplayUpon is called when entering the displayUpon production.
	EnterDisplayUpon(c *DisplayUponContext)

	// EnterDisplayWith is called when entering the displayWith production.
	EnterDisplayWith(c *DisplayWithContext)

	// EnterDivideStatement is called when entering the divideStatement production.
	EnterDivideStatement(c *DivideStatementContext)

	// EnterDivideIntoStatement is called when entering the divideIntoStatement production.
	EnterDivideIntoStatement(c *DivideIntoStatementContext)

	// EnterDivideIntoGivingStatement is called when entering the divideIntoGivingStatement production.
	EnterDivideIntoGivingStatement(c *DivideIntoGivingStatementContext)

	// EnterDivideByGivingStatement is called when entering the divideByGivingStatement production.
	EnterDivideByGivingStatement(c *DivideByGivingStatementContext)

	// EnterDivideGivingPhrase is called when entering the divideGivingPhrase production.
	EnterDivideGivingPhrase(c *DivideGivingPhraseContext)

	// EnterDivideInto is called when entering the divideInto production.
	EnterDivideInto(c *DivideIntoContext)

	// EnterDivideGiving is called when entering the divideGiving production.
	EnterDivideGiving(c *DivideGivingContext)

	// EnterDivideRemainder is called when entering the divideRemainder production.
	EnterDivideRemainder(c *DivideRemainderContext)

	// EnterEnableStatement is called when entering the enableStatement production.
	EnterEnableStatement(c *EnableStatementContext)

	// EnterEntryStatement is called when entering the entryStatement production.
	EnterEntryStatement(c *EntryStatementContext)

	// EnterEvaluateStatement is called when entering the evaluateStatement production.
	EnterEvaluateStatement(c *EvaluateStatementContext)

	// EnterEvaluateSelect is called when entering the evaluateSelect production.
	EnterEvaluateSelect(c *EvaluateSelectContext)

	// EnterEvaluateAlsoSelect is called when entering the evaluateAlsoSelect production.
	EnterEvaluateAlsoSelect(c *EvaluateAlsoSelectContext)

	// EnterEvaluateWhenPhrase is called when entering the evaluateWhenPhrase production.
	EnterEvaluateWhenPhrase(c *EvaluateWhenPhraseContext)

	// EnterEvaluateWhen is called when entering the evaluateWhen production.
	EnterEvaluateWhen(c *EvaluateWhenContext)

	// EnterEvaluateCondition is called when entering the evaluateCondition production.
	EnterEvaluateCondition(c *EvaluateConditionContext)

	// EnterEvaluateThrough is called when entering the evaluateThrough production.
	EnterEvaluateThrough(c *EvaluateThroughContext)

	// EnterEvaluateAlsoCondition is called when entering the evaluateAlsoCondition production.
	EnterEvaluateAlsoCondition(c *EvaluateAlsoConditionContext)

	// EnterEvaluateWhenOther is called when entering the evaluateWhenOther production.
	EnterEvaluateWhenOther(c *EvaluateWhenOtherContext)

	// EnterEvaluateValue is called when entering the evaluateValue production.
	EnterEvaluateValue(c *EvaluateValueContext)

	// EnterExecCicsStatement is called when entering the execCicsStatement production.
	EnterExecCicsStatement(c *ExecCicsStatementContext)

	// EnterExecSqlStatement is called when entering the execSqlStatement production.
	EnterExecSqlStatement(c *ExecSqlStatementContext)

	// EnterExecSqlImsStatement is called when entering the execSqlImsStatement production.
	EnterExecSqlImsStatement(c *ExecSqlImsStatementContext)

	// EnterExhibitStatement is called when entering the exhibitStatement production.
	EnterExhibitStatement(c *ExhibitStatementContext)

	// EnterExhibitOperand is called when entering the exhibitOperand production.
	EnterExhibitOperand(c *ExhibitOperandContext)

	// EnterExitStatement is called when entering the exitStatement production.
	EnterExitStatement(c *ExitStatementContext)

	// EnterGenerateStatement is called when entering the generateStatement production.
	EnterGenerateStatement(c *GenerateStatementContext)

	// EnterGobackStatement is called when entering the gobackStatement production.
	EnterGobackStatement(c *GobackStatementContext)

	// EnterGoToStatement is called when entering the goToStatement production.
	EnterGoToStatement(c *GoToStatementContext)

	// EnterGoToStatementSimple is called when entering the goToStatementSimple production.
	EnterGoToStatementSimple(c *GoToStatementSimpleContext)

	// EnterGoToDependingOnStatement is called when entering the goToDependingOnStatement production.
	EnterGoToDependingOnStatement(c *GoToDependingOnStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIfThen is called when entering the ifThen production.
	EnterIfThen(c *IfThenContext)

	// EnterIfElse is called when entering the ifElse production.
	EnterIfElse(c *IfElseContext)

	// EnterInitializeStatement is called when entering the initializeStatement production.
	EnterInitializeStatement(c *InitializeStatementContext)

	// EnterInitializeReplacingPhrase is called when entering the initializeReplacingPhrase production.
	EnterInitializeReplacingPhrase(c *InitializeReplacingPhraseContext)

	// EnterInitializeReplacingBy is called when entering the initializeReplacingBy production.
	EnterInitializeReplacingBy(c *InitializeReplacingByContext)

	// EnterInitiateStatement is called when entering the initiateStatement production.
	EnterInitiateStatement(c *InitiateStatementContext)

	// EnterInspectStatement is called when entering the inspectStatement production.
	EnterInspectStatement(c *InspectStatementContext)

	// EnterInspectTallyingPhrase is called when entering the inspectTallyingPhrase production.
	EnterInspectTallyingPhrase(c *InspectTallyingPhraseContext)

	// EnterInspectReplacingPhrase is called when entering the inspectReplacingPhrase production.
	EnterInspectReplacingPhrase(c *InspectReplacingPhraseContext)

	// EnterInspectTallyingReplacingPhrase is called when entering the inspectTallyingReplacingPhrase production.
	EnterInspectTallyingReplacingPhrase(c *InspectTallyingReplacingPhraseContext)

	// EnterInspectConvertingPhrase is called when entering the inspectConvertingPhrase production.
	EnterInspectConvertingPhrase(c *InspectConvertingPhraseContext)

	// EnterInspectFor is called when entering the inspectFor production.
	EnterInspectFor(c *InspectForContext)

	// EnterInspectCharacters is called when entering the inspectCharacters production.
	EnterInspectCharacters(c *InspectCharactersContext)

	// EnterInspectReplacingCharacters is called when entering the inspectReplacingCharacters production.
	EnterInspectReplacingCharacters(c *InspectReplacingCharactersContext)

	// EnterInspectAllLeadings is called when entering the inspectAllLeadings production.
	EnterInspectAllLeadings(c *InspectAllLeadingsContext)

	// EnterInspectReplacingAllLeadings is called when entering the inspectReplacingAllLeadings production.
	EnterInspectReplacingAllLeadings(c *InspectReplacingAllLeadingsContext)

	// EnterInspectAllLeading is called when entering the inspectAllLeading production.
	EnterInspectAllLeading(c *InspectAllLeadingContext)

	// EnterInspectReplacingAllLeading is called when entering the inspectReplacingAllLeading production.
	EnterInspectReplacingAllLeading(c *InspectReplacingAllLeadingContext)

	// EnterInspectBy is called when entering the inspectBy production.
	EnterInspectBy(c *InspectByContext)

	// EnterInspectTo is called when entering the inspectTo production.
	EnterInspectTo(c *InspectToContext)

	// EnterInspectBeforeAfter is called when entering the inspectBeforeAfter production.
	EnterInspectBeforeAfter(c *InspectBeforeAfterContext)

	// EnterMergeStatement is called when entering the mergeStatement production.
	EnterMergeStatement(c *MergeStatementContext)

	// EnterMergeOnKeyClause is called when entering the mergeOnKeyClause production.
	EnterMergeOnKeyClause(c *MergeOnKeyClauseContext)

	// EnterMergeCollatingSequencePhrase is called when entering the mergeCollatingSequencePhrase production.
	EnterMergeCollatingSequencePhrase(c *MergeCollatingSequencePhraseContext)

	// EnterMergeCollatingAlphanumeric is called when entering the mergeCollatingAlphanumeric production.
	EnterMergeCollatingAlphanumeric(c *MergeCollatingAlphanumericContext)

	// EnterMergeCollatingNational is called when entering the mergeCollatingNational production.
	EnterMergeCollatingNational(c *MergeCollatingNationalContext)

	// EnterMergeUsing is called when entering the mergeUsing production.
	EnterMergeUsing(c *MergeUsingContext)

	// EnterMergeOutputProcedurePhrase is called when entering the mergeOutputProcedurePhrase production.
	EnterMergeOutputProcedurePhrase(c *MergeOutputProcedurePhraseContext)

	// EnterMergeOutputThrough is called when entering the mergeOutputThrough production.
	EnterMergeOutputThrough(c *MergeOutputThroughContext)

	// EnterMergeGivingPhrase is called when entering the mergeGivingPhrase production.
	EnterMergeGivingPhrase(c *MergeGivingPhraseContext)

	// EnterMergeGiving is called when entering the mergeGiving production.
	EnterMergeGiving(c *MergeGivingContext)

	// EnterMoveStatement is called when entering the moveStatement production.
	EnterMoveStatement(c *MoveStatementContext)

	// EnterMoveToStatement is called when entering the moveToStatement production.
	EnterMoveToStatement(c *MoveToStatementContext)

	// EnterMoveToSendingArea is called when entering the moveToSendingArea production.
	EnterMoveToSendingArea(c *MoveToSendingAreaContext)

	// EnterMoveCorrespondingToStatement is called when entering the moveCorrespondingToStatement production.
	EnterMoveCorrespondingToStatement(c *MoveCorrespondingToStatementContext)

	// EnterMoveCorrespondingToSendingArea is called when entering the moveCorrespondingToSendingArea production.
	EnterMoveCorrespondingToSendingArea(c *MoveCorrespondingToSendingAreaContext)

	// EnterMultiplyStatement is called when entering the multiplyStatement production.
	EnterMultiplyStatement(c *MultiplyStatementContext)

	// EnterMultiplyRegular is called when entering the multiplyRegular production.
	EnterMultiplyRegular(c *MultiplyRegularContext)

	// EnterMultiplyRegularOperand is called when entering the multiplyRegularOperand production.
	EnterMultiplyRegularOperand(c *MultiplyRegularOperandContext)

	// EnterMultiplyGiving is called when entering the multiplyGiving production.
	EnterMultiplyGiving(c *MultiplyGivingContext)

	// EnterMultiplyGivingOperand is called when entering the multiplyGivingOperand production.
	EnterMultiplyGivingOperand(c *MultiplyGivingOperandContext)

	// EnterMultiplyGivingResult is called when entering the multiplyGivingResult production.
	EnterMultiplyGivingResult(c *MultiplyGivingResultContext)

	// EnterOpenStatement is called when entering the openStatement production.
	EnterOpenStatement(c *OpenStatementContext)

	// EnterOpenInputStatement is called when entering the openInputStatement production.
	EnterOpenInputStatement(c *OpenInputStatementContext)

	// EnterOpenInput is called when entering the openInput production.
	EnterOpenInput(c *OpenInputContext)

	// EnterOpenOutputStatement is called when entering the openOutputStatement production.
	EnterOpenOutputStatement(c *OpenOutputStatementContext)

	// EnterOpenOutput is called when entering the openOutput production.
	EnterOpenOutput(c *OpenOutputContext)

	// EnterOpenIOStatement is called when entering the openIOStatement production.
	EnterOpenIOStatement(c *OpenIOStatementContext)

	// EnterOpenExtendStatement is called when entering the openExtendStatement production.
	EnterOpenExtendStatement(c *OpenExtendStatementContext)

	// EnterPerformStatement is called when entering the performStatement production.
	EnterPerformStatement(c *PerformStatementContext)

	// EnterPerformInlineStatement is called when entering the performInlineStatement production.
	EnterPerformInlineStatement(c *PerformInlineStatementContext)

	// EnterPerformProcedureStatement is called when entering the performProcedureStatement production.
	EnterPerformProcedureStatement(c *PerformProcedureStatementContext)

	// EnterPerformType is called when entering the performType production.
	EnterPerformType(c *PerformTypeContext)

	// EnterPerformTimes is called when entering the performTimes production.
	EnterPerformTimes(c *PerformTimesContext)

	// EnterPerformUntil is called when entering the performUntil production.
	EnterPerformUntil(c *PerformUntilContext)

	// EnterPerformVarying is called when entering the performVarying production.
	EnterPerformVarying(c *PerformVaryingContext)

	// EnterPerformVaryingClause is called when entering the performVaryingClause production.
	EnterPerformVaryingClause(c *PerformVaryingClauseContext)

	// EnterPerformVaryingPhrase is called when entering the performVaryingPhrase production.
	EnterPerformVaryingPhrase(c *PerformVaryingPhraseContext)

	// EnterPerformAfter is called when entering the performAfter production.
	EnterPerformAfter(c *PerformAfterContext)

	// EnterPerformFrom is called when entering the performFrom production.
	EnterPerformFrom(c *PerformFromContext)

	// EnterPerformBy is called when entering the performBy production.
	EnterPerformBy(c *PerformByContext)

	// EnterPerformTestClause is called when entering the performTestClause production.
	EnterPerformTestClause(c *PerformTestClauseContext)

	// EnterPurgeStatement is called when entering the purgeStatement production.
	EnterPurgeStatement(c *PurgeStatementContext)

	// EnterReadStatement is called when entering the readStatement production.
	EnterReadStatement(c *ReadStatementContext)

	// EnterReadInto is called when entering the readInto production.
	EnterReadInto(c *ReadIntoContext)

	// EnterReadWith is called when entering the readWith production.
	EnterReadWith(c *ReadWithContext)

	// EnterReadKey is called when entering the readKey production.
	EnterReadKey(c *ReadKeyContext)

	// EnterReceiveStatement is called when entering the receiveStatement production.
	EnterReceiveStatement(c *ReceiveStatementContext)

	// EnterReceiveFromStatement is called when entering the receiveFromStatement production.
	EnterReceiveFromStatement(c *ReceiveFromStatementContext)

	// EnterReceiveFrom is called when entering the receiveFrom production.
	EnterReceiveFrom(c *ReceiveFromContext)

	// EnterReceiveIntoStatement is called when entering the receiveIntoStatement production.
	EnterReceiveIntoStatement(c *ReceiveIntoStatementContext)

	// EnterReceiveNoData is called when entering the receiveNoData production.
	EnterReceiveNoData(c *ReceiveNoDataContext)

	// EnterReceiveWithData is called when entering the receiveWithData production.
	EnterReceiveWithData(c *ReceiveWithDataContext)

	// EnterReceiveBefore is called when entering the receiveBefore production.
	EnterReceiveBefore(c *ReceiveBeforeContext)

	// EnterReceiveWith is called when entering the receiveWith production.
	EnterReceiveWith(c *ReceiveWithContext)

	// EnterReceiveThread is called when entering the receiveThread production.
	EnterReceiveThread(c *ReceiveThreadContext)

	// EnterReceiveSize is called when entering the receiveSize production.
	EnterReceiveSize(c *ReceiveSizeContext)

	// EnterReceiveStatus is called when entering the receiveStatus production.
	EnterReceiveStatus(c *ReceiveStatusContext)

	// EnterReleaseStatement is called when entering the releaseStatement production.
	EnterReleaseStatement(c *ReleaseStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterReturnInto is called when entering the returnInto production.
	EnterReturnInto(c *ReturnIntoContext)

	// EnterRewriteStatement is called when entering the rewriteStatement production.
	EnterRewriteStatement(c *RewriteStatementContext)

	// EnterRewriteFrom is called when entering the rewriteFrom production.
	EnterRewriteFrom(c *RewriteFromContext)

	// EnterSearchStatement is called when entering the searchStatement production.
	EnterSearchStatement(c *SearchStatementContext)

	// EnterSearchVarying is called when entering the searchVarying production.
	EnterSearchVarying(c *SearchVaryingContext)

	// EnterSearchWhen is called when entering the searchWhen production.
	EnterSearchWhen(c *SearchWhenContext)

	// EnterSendStatement is called when entering the sendStatement production.
	EnterSendStatement(c *SendStatementContext)

	// EnterSendStatementSync is called when entering the sendStatementSync production.
	EnterSendStatementSync(c *SendStatementSyncContext)

	// EnterSendStatementAsync is called when entering the sendStatementAsync production.
	EnterSendStatementAsync(c *SendStatementAsyncContext)

	// EnterSendFromPhrase is called when entering the sendFromPhrase production.
	EnterSendFromPhrase(c *SendFromPhraseContext)

	// EnterSendWithPhrase is called when entering the sendWithPhrase production.
	EnterSendWithPhrase(c *SendWithPhraseContext)

	// EnterSendReplacingPhrase is called when entering the sendReplacingPhrase production.
	EnterSendReplacingPhrase(c *SendReplacingPhraseContext)

	// EnterSendAdvancingPhrase is called when entering the sendAdvancingPhrase production.
	EnterSendAdvancingPhrase(c *SendAdvancingPhraseContext)

	// EnterSendAdvancingPage is called when entering the sendAdvancingPage production.
	EnterSendAdvancingPage(c *SendAdvancingPageContext)

	// EnterSendAdvancingLines is called when entering the sendAdvancingLines production.
	EnterSendAdvancingLines(c *SendAdvancingLinesContext)

	// EnterSendAdvancingMnemonic is called when entering the sendAdvancingMnemonic production.
	EnterSendAdvancingMnemonic(c *SendAdvancingMnemonicContext)

	// EnterSetStatement is called when entering the setStatement production.
	EnterSetStatement(c *SetStatementContext)

	// EnterSetToStatement is called when entering the setToStatement production.
	EnterSetToStatement(c *SetToStatementContext)

	// EnterSetUpDownByStatement is called when entering the setUpDownByStatement production.
	EnterSetUpDownByStatement(c *SetUpDownByStatementContext)

	// EnterSetTo is called when entering the setTo production.
	EnterSetTo(c *SetToContext)

	// EnterSetToValue is called when entering the setToValue production.
	EnterSetToValue(c *SetToValueContext)

	// EnterSetByValue is called when entering the setByValue production.
	EnterSetByValue(c *SetByValueContext)

	// EnterSortStatement is called when entering the sortStatement production.
	EnterSortStatement(c *SortStatementContext)

	// EnterSortOnKeyClause is called when entering the sortOnKeyClause production.
	EnterSortOnKeyClause(c *SortOnKeyClauseContext)

	// EnterSortDuplicatesPhrase is called when entering the sortDuplicatesPhrase production.
	EnterSortDuplicatesPhrase(c *SortDuplicatesPhraseContext)

	// EnterSortCollatingSequencePhrase is called when entering the sortCollatingSequencePhrase production.
	EnterSortCollatingSequencePhrase(c *SortCollatingSequencePhraseContext)

	// EnterSortCollatingAlphanumeric is called when entering the sortCollatingAlphanumeric production.
	EnterSortCollatingAlphanumeric(c *SortCollatingAlphanumericContext)

	// EnterSortCollatingNational is called when entering the sortCollatingNational production.
	EnterSortCollatingNational(c *SortCollatingNationalContext)

	// EnterSortInputProcedurePhrase is called when entering the sortInputProcedurePhrase production.
	EnterSortInputProcedurePhrase(c *SortInputProcedurePhraseContext)

	// EnterSortInputThrough is called when entering the sortInputThrough production.
	EnterSortInputThrough(c *SortInputThroughContext)

	// EnterSortUsing is called when entering the sortUsing production.
	EnterSortUsing(c *SortUsingContext)

	// EnterSortOutputProcedurePhrase is called when entering the sortOutputProcedurePhrase production.
	EnterSortOutputProcedurePhrase(c *SortOutputProcedurePhraseContext)

	// EnterSortOutputThrough is called when entering the sortOutputThrough production.
	EnterSortOutputThrough(c *SortOutputThroughContext)

	// EnterSortGivingPhrase is called when entering the sortGivingPhrase production.
	EnterSortGivingPhrase(c *SortGivingPhraseContext)

	// EnterSortGiving is called when entering the sortGiving production.
	EnterSortGiving(c *SortGivingContext)

	// EnterStartStatement is called when entering the startStatement production.
	EnterStartStatement(c *StartStatementContext)

	// EnterStartKey is called when entering the startKey production.
	EnterStartKey(c *StartKeyContext)

	// EnterStopStatement is called when entering the stopStatement production.
	EnterStopStatement(c *StopStatementContext)

	// EnterStringStatement is called when entering the stringStatement production.
	EnterStringStatement(c *StringStatementContext)

	// EnterStringSendingPhrase is called when entering the stringSendingPhrase production.
	EnterStringSendingPhrase(c *StringSendingPhraseContext)

	// EnterStringSending is called when entering the stringSending production.
	EnterStringSending(c *StringSendingContext)

	// EnterStringDelimitedByPhrase is called when entering the stringDelimitedByPhrase production.
	EnterStringDelimitedByPhrase(c *StringDelimitedByPhraseContext)

	// EnterStringForPhrase is called when entering the stringForPhrase production.
	EnterStringForPhrase(c *StringForPhraseContext)

	// EnterStringIntoPhrase is called when entering the stringIntoPhrase production.
	EnterStringIntoPhrase(c *StringIntoPhraseContext)

	// EnterStringWithPointerPhrase is called when entering the stringWithPointerPhrase production.
	EnterStringWithPointerPhrase(c *StringWithPointerPhraseContext)

	// EnterSubtractStatement is called when entering the subtractStatement production.
	EnterSubtractStatement(c *SubtractStatementContext)

	// EnterSubtractFromStatement is called when entering the subtractFromStatement production.
	EnterSubtractFromStatement(c *SubtractFromStatementContext)

	// EnterSubtractFromGivingStatement is called when entering the subtractFromGivingStatement production.
	EnterSubtractFromGivingStatement(c *SubtractFromGivingStatementContext)

	// EnterSubtractCorrespondingStatement is called when entering the subtractCorrespondingStatement production.
	EnterSubtractCorrespondingStatement(c *SubtractCorrespondingStatementContext)

	// EnterSubtractSubtrahend is called when entering the subtractSubtrahend production.
	EnterSubtractSubtrahend(c *SubtractSubtrahendContext)

	// EnterSubtractMinuend is called when entering the subtractMinuend production.
	EnterSubtractMinuend(c *SubtractMinuendContext)

	// EnterSubtractMinuendGiving is called when entering the subtractMinuendGiving production.
	EnterSubtractMinuendGiving(c *SubtractMinuendGivingContext)

	// EnterSubtractGiving is called when entering the subtractGiving production.
	EnterSubtractGiving(c *SubtractGivingContext)

	// EnterSubtractMinuendCorresponding is called when entering the subtractMinuendCorresponding production.
	EnterSubtractMinuendCorresponding(c *SubtractMinuendCorrespondingContext)

	// EnterTerminateStatement is called when entering the terminateStatement production.
	EnterTerminateStatement(c *TerminateStatementContext)

	// EnterUnstringStatement is called when entering the unstringStatement production.
	EnterUnstringStatement(c *UnstringStatementContext)

	// EnterUnstringSendingPhrase is called when entering the unstringSendingPhrase production.
	EnterUnstringSendingPhrase(c *UnstringSendingPhraseContext)

	// EnterUnstringDelimitedByPhrase is called when entering the unstringDelimitedByPhrase production.
	EnterUnstringDelimitedByPhrase(c *UnstringDelimitedByPhraseContext)

	// EnterUnstringOrAllPhrase is called when entering the unstringOrAllPhrase production.
	EnterUnstringOrAllPhrase(c *UnstringOrAllPhraseContext)

	// EnterUnstringIntoPhrase is called when entering the unstringIntoPhrase production.
	EnterUnstringIntoPhrase(c *UnstringIntoPhraseContext)

	// EnterUnstringInto is called when entering the unstringInto production.
	EnterUnstringInto(c *UnstringIntoContext)

	// EnterUnstringDelimiterIn is called when entering the unstringDelimiterIn production.
	EnterUnstringDelimiterIn(c *UnstringDelimiterInContext)

	// EnterUnstringCountIn is called when entering the unstringCountIn production.
	EnterUnstringCountIn(c *UnstringCountInContext)

	// EnterUnstringWithPointerPhrase is called when entering the unstringWithPointerPhrase production.
	EnterUnstringWithPointerPhrase(c *UnstringWithPointerPhraseContext)

	// EnterUnstringTallyingPhrase is called when entering the unstringTallyingPhrase production.
	EnterUnstringTallyingPhrase(c *UnstringTallyingPhraseContext)

	// EnterUseStatement is called when entering the useStatement production.
	EnterUseStatement(c *UseStatementContext)

	// EnterUseAfterClause is called when entering the useAfterClause production.
	EnterUseAfterClause(c *UseAfterClauseContext)

	// EnterUseAfterOn is called when entering the useAfterOn production.
	EnterUseAfterOn(c *UseAfterOnContext)

	// EnterUseDebugClause is called when entering the useDebugClause production.
	EnterUseDebugClause(c *UseDebugClauseContext)

	// EnterUseDebugOn is called when entering the useDebugOn production.
	EnterUseDebugOn(c *UseDebugOnContext)

	// EnterWriteStatement is called when entering the writeStatement production.
	EnterWriteStatement(c *WriteStatementContext)

	// EnterWriteFromPhrase is called when entering the writeFromPhrase production.
	EnterWriteFromPhrase(c *WriteFromPhraseContext)

	// EnterWriteAdvancingPhrase is called when entering the writeAdvancingPhrase production.
	EnterWriteAdvancingPhrase(c *WriteAdvancingPhraseContext)

	// EnterWriteAdvancingPage is called when entering the writeAdvancingPage production.
	EnterWriteAdvancingPage(c *WriteAdvancingPageContext)

	// EnterWriteAdvancingLines is called when entering the writeAdvancingLines production.
	EnterWriteAdvancingLines(c *WriteAdvancingLinesContext)

	// EnterWriteAdvancingMnemonic is called when entering the writeAdvancingMnemonic production.
	EnterWriteAdvancingMnemonic(c *WriteAdvancingMnemonicContext)

	// EnterWriteAtEndOfPagePhrase is called when entering the writeAtEndOfPagePhrase production.
	EnterWriteAtEndOfPagePhrase(c *WriteAtEndOfPagePhraseContext)

	// EnterWriteNotAtEndOfPagePhrase is called when entering the writeNotAtEndOfPagePhrase production.
	EnterWriteNotAtEndOfPagePhrase(c *WriteNotAtEndOfPagePhraseContext)

	// EnterAtEndPhrase is called when entering the atEndPhrase production.
	EnterAtEndPhrase(c *AtEndPhraseContext)

	// EnterNotAtEndPhrase is called when entering the notAtEndPhrase production.
	EnterNotAtEndPhrase(c *NotAtEndPhraseContext)

	// EnterInvalidKeyPhrase is called when entering the invalidKeyPhrase production.
	EnterInvalidKeyPhrase(c *InvalidKeyPhraseContext)

	// EnterNotInvalidKeyPhrase is called when entering the notInvalidKeyPhrase production.
	EnterNotInvalidKeyPhrase(c *NotInvalidKeyPhraseContext)

	// EnterOnOverflowPhrase is called when entering the onOverflowPhrase production.
	EnterOnOverflowPhrase(c *OnOverflowPhraseContext)

	// EnterNotOnOverflowPhrase is called when entering the notOnOverflowPhrase production.
	EnterNotOnOverflowPhrase(c *NotOnOverflowPhraseContext)

	// EnterOnSizeErrorPhrase is called when entering the onSizeErrorPhrase production.
	EnterOnSizeErrorPhrase(c *OnSizeErrorPhraseContext)

	// EnterNotOnSizeErrorPhrase is called when entering the notOnSizeErrorPhrase production.
	EnterNotOnSizeErrorPhrase(c *NotOnSizeErrorPhraseContext)

	// EnterOnExceptionClause is called when entering the onExceptionClause production.
	EnterOnExceptionClause(c *OnExceptionClauseContext)

	// EnterNotOnExceptionClause is called when entering the notOnExceptionClause production.
	EnterNotOnExceptionClause(c *NotOnExceptionClauseContext)

	// EnterArithmeticExpression is called when entering the arithmeticExpression production.
	EnterArithmeticExpression(c *ArithmeticExpressionContext)

	// EnterPlusMinus is called when entering the plusMinus production.
	EnterPlusMinus(c *PlusMinusContext)

	// EnterMultDivs is called when entering the multDivs production.
	EnterMultDivs(c *MultDivsContext)

	// EnterMultDiv is called when entering the multDiv production.
	EnterMultDiv(c *MultDivContext)

	// EnterPowers is called when entering the powers production.
	EnterPowers(c *PowersContext)

	// EnterPower is called when entering the power production.
	EnterPower(c *PowerContext)

	// EnterBasis is called when entering the basis production.
	EnterBasis(c *BasisContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterAndOrCondition is called when entering the andOrCondition production.
	EnterAndOrCondition(c *AndOrConditionContext)

	// EnterCombinableCondition is called when entering the combinableCondition production.
	EnterCombinableCondition(c *CombinableConditionContext)

	// EnterSimpleCondition is called when entering the simpleCondition production.
	EnterSimpleCondition(c *SimpleConditionContext)

	// EnterClassCondition is called when entering the classCondition production.
	EnterClassCondition(c *ClassConditionContext)

	// EnterConditionNameReference is called when entering the conditionNameReference production.
	EnterConditionNameReference(c *ConditionNameReferenceContext)

	// EnterConditionNameSubscriptReference is called when entering the conditionNameSubscriptReference production.
	EnterConditionNameSubscriptReference(c *ConditionNameSubscriptReferenceContext)

	// EnterRelationCondition is called when entering the relationCondition production.
	EnterRelationCondition(c *RelationConditionContext)

	// EnterRelationSignCondition is called when entering the relationSignCondition production.
	EnterRelationSignCondition(c *RelationSignConditionContext)

	// EnterRelationArithmeticComparison is called when entering the relationArithmeticComparison production.
	EnterRelationArithmeticComparison(c *RelationArithmeticComparisonContext)

	// EnterRelationCombinedComparison is called when entering the relationCombinedComparison production.
	EnterRelationCombinedComparison(c *RelationCombinedComparisonContext)

	// EnterRelationCombinedCondition is called when entering the relationCombinedCondition production.
	EnterRelationCombinedCondition(c *RelationCombinedConditionContext)

	// EnterRelationalOperator is called when entering the relationalOperator production.
	EnterRelationalOperator(c *RelationalOperatorContext)

	// EnterAbbreviation is called when entering the abbreviation production.
	EnterAbbreviation(c *AbbreviationContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterTableCall is called when entering the tableCall production.
	EnterTableCall(c *TableCallContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterReferenceModifier is called when entering the referenceModifier production.
	EnterReferenceModifier(c *ReferenceModifierContext)

	// EnterCharacterPosition is called when entering the characterPosition production.
	EnterCharacterPosition(c *CharacterPositionContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// EnterSubscript_ is called when entering the subscript_ production.
	EnterSubscript_(c *Subscript_Context)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterQualifiedDataName is called when entering the qualifiedDataName production.
	EnterQualifiedDataName(c *QualifiedDataNameContext)

	// EnterQualifiedDataNameFormat1 is called when entering the qualifiedDataNameFormat1 production.
	EnterQualifiedDataNameFormat1(c *QualifiedDataNameFormat1Context)

	// EnterQualifiedDataNameFormat2 is called when entering the qualifiedDataNameFormat2 production.
	EnterQualifiedDataNameFormat2(c *QualifiedDataNameFormat2Context)

	// EnterQualifiedDataNameFormat3 is called when entering the qualifiedDataNameFormat3 production.
	EnterQualifiedDataNameFormat3(c *QualifiedDataNameFormat3Context)

	// EnterQualifiedDataNameFormat4 is called when entering the qualifiedDataNameFormat4 production.
	EnterQualifiedDataNameFormat4(c *QualifiedDataNameFormat4Context)

	// EnterQualifiedInData is called when entering the qualifiedInData production.
	EnterQualifiedInData(c *QualifiedInDataContext)

	// EnterInData is called when entering the inData production.
	EnterInData(c *InDataContext)

	// EnterInFile is called when entering the inFile production.
	EnterInFile(c *InFileContext)

	// EnterInMnemonic is called when entering the inMnemonic production.
	EnterInMnemonic(c *InMnemonicContext)

	// EnterInSection is called when entering the inSection production.
	EnterInSection(c *InSectionContext)

	// EnterInLibrary is called when entering the inLibrary production.
	EnterInLibrary(c *InLibraryContext)

	// EnterInTable is called when entering the inTable production.
	EnterInTable(c *InTableContext)

	// EnterAlphabetName is called when entering the alphabetName production.
	EnterAlphabetName(c *AlphabetNameContext)

	// EnterAssignmentName is called when entering the assignmentName production.
	EnterAssignmentName(c *AssignmentNameContext)

	// EnterBasisName is called when entering the basisName production.
	EnterBasisName(c *BasisNameContext)

	// EnterCdName is called when entering the cdName production.
	EnterCdName(c *CdNameContext)

	// EnterClassName is called when entering the className production.
	EnterClassName(c *ClassNameContext)

	// EnterComputerName is called when entering the computerName production.
	EnterComputerName(c *ComputerNameContext)

	// EnterConditionName is called when entering the conditionName production.
	EnterConditionName(c *ConditionNameContext)

	// EnterDataName is called when entering the dataName production.
	EnterDataName(c *DataNameContext)

	// EnterDataDescName is called when entering the dataDescName production.
	EnterDataDescName(c *DataDescNameContext)

	// EnterEnvironmentName is called when entering the environmentName production.
	EnterEnvironmentName(c *EnvironmentNameContext)

	// EnterFileName is called when entering the fileName production.
	EnterFileName(c *FileNameContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterLanguageName is called when entering the languageName production.
	EnterLanguageName(c *LanguageNameContext)

	// EnterLibraryName is called when entering the libraryName production.
	EnterLibraryName(c *LibraryNameContext)

	// EnterLocalName is called when entering the localName production.
	EnterLocalName(c *LocalNameContext)

	// EnterMnemonicName is called when entering the mnemonicName production.
	EnterMnemonicName(c *MnemonicNameContext)

	// EnterParagraphName is called when entering the paragraphName production.
	EnterParagraphName(c *ParagraphNameContext)

	// EnterProcedureName is called when entering the procedureName production.
	EnterProcedureName(c *ProcedureNameContext)

	// EnterProgramName is called when entering the programName production.
	EnterProgramName(c *ProgramNameContext)

	// EnterRecordName is called when entering the recordName production.
	EnterRecordName(c *RecordNameContext)

	// EnterReportName is called when entering the reportName production.
	EnterReportName(c *ReportNameContext)

	// EnterRoutineName is called when entering the routineName production.
	EnterRoutineName(c *RoutineNameContext)

	// EnterScreenName is called when entering the screenName production.
	EnterScreenName(c *ScreenNameContext)

	// EnterSectionName is called when entering the sectionName production.
	EnterSectionName(c *SectionNameContext)

	// EnterSystemName is called when entering the systemName production.
	EnterSystemName(c *SystemNameContext)

	// EnterSymbolicCharacter is called when entering the symbolicCharacter production.
	EnterSymbolicCharacter(c *SymbolicCharacterContext)

	// EnterTextName is called when entering the textName production.
	EnterTextName(c *TextNameContext)

	// EnterCobolWord is called when entering the cobolWord production.
	EnterCobolWord(c *CobolWordContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterCicsDfhRespLiteral is called when entering the cicsDfhRespLiteral production.
	EnterCicsDfhRespLiteral(c *CicsDfhRespLiteralContext)

	// EnterCicsDfhValueLiteral is called when entering the cicsDfhValueLiteral production.
	EnterCicsDfhValueLiteral(c *CicsDfhValueLiteralContext)

	// EnterFigurativeConstant is called when entering the figurativeConstant production.
	EnterFigurativeConstant(c *FigurativeConstantContext)

	// EnterSpecialRegister is called when entering the specialRegister production.
	EnterSpecialRegister(c *SpecialRegisterContext)

	// EnterCommentEntry is called when entering the commentEntry production.
	EnterCommentEntry(c *CommentEntryContext)

	// ExitStartRule is called when exiting the startRule production.
	ExitStartRule(c *StartRuleContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitProgramUnit is called when exiting the programUnit production.
	ExitProgramUnit(c *ProgramUnitContext)

	// ExitEndProgramStatement is called when exiting the endProgramStatement production.
	ExitEndProgramStatement(c *EndProgramStatementContext)

	// ExitIdentificationDivision is called when exiting the identificationDivision production.
	ExitIdentificationDivision(c *IdentificationDivisionContext)

	// ExitIdentificationDivisionBody is called when exiting the identificationDivisionBody production.
	ExitIdentificationDivisionBody(c *IdentificationDivisionBodyContext)

	// ExitProgramIdParagraph is called when exiting the programIdParagraph production.
	ExitProgramIdParagraph(c *ProgramIdParagraphContext)

	// ExitAuthorParagraph is called when exiting the authorParagraph production.
	ExitAuthorParagraph(c *AuthorParagraphContext)

	// ExitInstallationParagraph is called when exiting the installationParagraph production.
	ExitInstallationParagraph(c *InstallationParagraphContext)

	// ExitDateWrittenParagraph is called when exiting the dateWrittenParagraph production.
	ExitDateWrittenParagraph(c *DateWrittenParagraphContext)

	// ExitDateCompiledParagraph is called when exiting the dateCompiledParagraph production.
	ExitDateCompiledParagraph(c *DateCompiledParagraphContext)

	// ExitSecurityParagraph is called when exiting the securityParagraph production.
	ExitSecurityParagraph(c *SecurityParagraphContext)

	// ExitRemarksParagraph is called when exiting the remarksParagraph production.
	ExitRemarksParagraph(c *RemarksParagraphContext)

	// ExitEnvironmentDivision is called when exiting the environmentDivision production.
	ExitEnvironmentDivision(c *EnvironmentDivisionContext)

	// ExitEnvironmentDivisionBody is called when exiting the environmentDivisionBody production.
	ExitEnvironmentDivisionBody(c *EnvironmentDivisionBodyContext)

	// ExitConfigurationSection is called when exiting the configurationSection production.
	ExitConfigurationSection(c *ConfigurationSectionContext)

	// ExitConfigurationSectionParagraph is called when exiting the configurationSectionParagraph production.
	ExitConfigurationSectionParagraph(c *ConfigurationSectionParagraphContext)

	// ExitSourceComputerParagraph is called when exiting the sourceComputerParagraph production.
	ExitSourceComputerParagraph(c *SourceComputerParagraphContext)

	// ExitObjectComputerParagraph is called when exiting the objectComputerParagraph production.
	ExitObjectComputerParagraph(c *ObjectComputerParagraphContext)

	// ExitObjectComputerClause is called when exiting the objectComputerClause production.
	ExitObjectComputerClause(c *ObjectComputerClauseContext)

	// ExitMemorySizeClause is called when exiting the memorySizeClause production.
	ExitMemorySizeClause(c *MemorySizeClauseContext)

	// ExitDiskSizeClause is called when exiting the diskSizeClause production.
	ExitDiskSizeClause(c *DiskSizeClauseContext)

	// ExitCollatingSequenceClause is called when exiting the collatingSequenceClause production.
	ExitCollatingSequenceClause(c *CollatingSequenceClauseContext)

	// ExitCollatingSequenceClauseAlphanumeric is called when exiting the collatingSequenceClauseAlphanumeric production.
	ExitCollatingSequenceClauseAlphanumeric(c *CollatingSequenceClauseAlphanumericContext)

	// ExitCollatingSequenceClauseNational is called when exiting the collatingSequenceClauseNational production.
	ExitCollatingSequenceClauseNational(c *CollatingSequenceClauseNationalContext)

	// ExitSegmentLimitClause is called when exiting the segmentLimitClause production.
	ExitSegmentLimitClause(c *SegmentLimitClauseContext)

	// ExitCharacterSetClause is called when exiting the characterSetClause production.
	ExitCharacterSetClause(c *CharacterSetClauseContext)

	// ExitSpecialNamesParagraph is called when exiting the specialNamesParagraph production.
	ExitSpecialNamesParagraph(c *SpecialNamesParagraphContext)

	// ExitSpecialNameClause is called when exiting the specialNameClause production.
	ExitSpecialNameClause(c *SpecialNameClauseContext)

	// ExitAlphabetClause is called when exiting the alphabetClause production.
	ExitAlphabetClause(c *AlphabetClauseContext)

	// ExitAlphabetClauseFormat1 is called when exiting the alphabetClauseFormat1 production.
	ExitAlphabetClauseFormat1(c *AlphabetClauseFormat1Context)

	// ExitAlphabetLiterals is called when exiting the alphabetLiterals production.
	ExitAlphabetLiterals(c *AlphabetLiteralsContext)

	// ExitAlphabetThrough is called when exiting the alphabetThrough production.
	ExitAlphabetThrough(c *AlphabetThroughContext)

	// ExitAlphabetAlso is called when exiting the alphabetAlso production.
	ExitAlphabetAlso(c *AlphabetAlsoContext)

	// ExitAlphabetClauseFormat2 is called when exiting the alphabetClauseFormat2 production.
	ExitAlphabetClauseFormat2(c *AlphabetClauseFormat2Context)

	// ExitChannelClause is called when exiting the channelClause production.
	ExitChannelClause(c *ChannelClauseContext)

	// ExitClassClause is called when exiting the classClause production.
	ExitClassClause(c *ClassClauseContext)

	// ExitClassClauseThrough is called when exiting the classClauseThrough production.
	ExitClassClauseThrough(c *ClassClauseThroughContext)

	// ExitClassClauseFrom is called when exiting the classClauseFrom production.
	ExitClassClauseFrom(c *ClassClauseFromContext)

	// ExitClassClauseTo is called when exiting the classClauseTo production.
	ExitClassClauseTo(c *ClassClauseToContext)

	// ExitCurrencySignClause is called when exiting the currencySignClause production.
	ExitCurrencySignClause(c *CurrencySignClauseContext)

	// ExitDecimalPointClause is called when exiting the decimalPointClause production.
	ExitDecimalPointClause(c *DecimalPointClauseContext)

	// ExitDefaultComputationalSignClause is called when exiting the defaultComputationalSignClause production.
	ExitDefaultComputationalSignClause(c *DefaultComputationalSignClauseContext)

	// ExitDefaultDisplaySignClause is called when exiting the defaultDisplaySignClause production.
	ExitDefaultDisplaySignClause(c *DefaultDisplaySignClauseContext)

	// ExitEnvironmentSwitchNameClause is called when exiting the environmentSwitchNameClause production.
	ExitEnvironmentSwitchNameClause(c *EnvironmentSwitchNameClauseContext)

	// ExitEnvironmentSwitchNameSpecialNamesStatusPhrase is called when exiting the environmentSwitchNameSpecialNamesStatusPhrase production.
	ExitEnvironmentSwitchNameSpecialNamesStatusPhrase(c *EnvironmentSwitchNameSpecialNamesStatusPhraseContext)

	// ExitOdtClause is called when exiting the odtClause production.
	ExitOdtClause(c *OdtClauseContext)

	// ExitReserveNetworkClause is called when exiting the reserveNetworkClause production.
	ExitReserveNetworkClause(c *ReserveNetworkClauseContext)

	// ExitSymbolicCharactersClause is called when exiting the symbolicCharactersClause production.
	ExitSymbolicCharactersClause(c *SymbolicCharactersClauseContext)

	// ExitSymbolicCharacters is called when exiting the symbolicCharacters production.
	ExitSymbolicCharacters(c *SymbolicCharactersContext)

	// ExitInputOutputSection is called when exiting the inputOutputSection production.
	ExitInputOutputSection(c *InputOutputSectionContext)

	// ExitInputOutputSectionParagraph is called when exiting the inputOutputSectionParagraph production.
	ExitInputOutputSectionParagraph(c *InputOutputSectionParagraphContext)

	// ExitFileControlParagraph is called when exiting the fileControlParagraph production.
	ExitFileControlParagraph(c *FileControlParagraphContext)

	// ExitFileControlEntry is called when exiting the fileControlEntry production.
	ExitFileControlEntry(c *FileControlEntryContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitFileControlClause is called when exiting the fileControlClause production.
	ExitFileControlClause(c *FileControlClauseContext)

	// ExitAssignClause is called when exiting the assignClause production.
	ExitAssignClause(c *AssignClauseContext)

	// ExitReserveClause is called when exiting the reserveClause production.
	ExitReserveClause(c *ReserveClauseContext)

	// ExitOrganizationClause is called when exiting the organizationClause production.
	ExitOrganizationClause(c *OrganizationClauseContext)

	// ExitPaddingCharacterClause is called when exiting the paddingCharacterClause production.
	ExitPaddingCharacterClause(c *PaddingCharacterClauseContext)

	// ExitRecordDelimiterClause is called when exiting the recordDelimiterClause production.
	ExitRecordDelimiterClause(c *RecordDelimiterClauseContext)

	// ExitAccessModeClause is called when exiting the accessModeClause production.
	ExitAccessModeClause(c *AccessModeClauseContext)

	// ExitRecordKeyClause is called when exiting the recordKeyClause production.
	ExitRecordKeyClause(c *RecordKeyClauseContext)

	// ExitAlternateRecordKeyClause is called when exiting the alternateRecordKeyClause production.
	ExitAlternateRecordKeyClause(c *AlternateRecordKeyClauseContext)

	// ExitPasswordClause is called when exiting the passwordClause production.
	ExitPasswordClause(c *PasswordClauseContext)

	// ExitFileStatusClause is called when exiting the fileStatusClause production.
	ExitFileStatusClause(c *FileStatusClauseContext)

	// ExitRelativeKeyClause is called when exiting the relativeKeyClause production.
	ExitRelativeKeyClause(c *RelativeKeyClauseContext)

	// ExitIoControlParagraph is called when exiting the ioControlParagraph production.
	ExitIoControlParagraph(c *IoControlParagraphContext)

	// ExitIoControlClause is called when exiting the ioControlClause production.
	ExitIoControlClause(c *IoControlClauseContext)

	// ExitRerunClause is called when exiting the rerunClause production.
	ExitRerunClause(c *RerunClauseContext)

	// ExitRerunEveryRecords is called when exiting the rerunEveryRecords production.
	ExitRerunEveryRecords(c *RerunEveryRecordsContext)

	// ExitRerunEveryOf is called when exiting the rerunEveryOf production.
	ExitRerunEveryOf(c *RerunEveryOfContext)

	// ExitRerunEveryClock is called when exiting the rerunEveryClock production.
	ExitRerunEveryClock(c *RerunEveryClockContext)

	// ExitSameClause is called when exiting the sameClause production.
	ExitSameClause(c *SameClauseContext)

	// ExitMultipleFileClause is called when exiting the multipleFileClause production.
	ExitMultipleFileClause(c *MultipleFileClauseContext)

	// ExitMultipleFilePosition is called when exiting the multipleFilePosition production.
	ExitMultipleFilePosition(c *MultipleFilePositionContext)

	// ExitCommitmentControlClause is called when exiting the commitmentControlClause production.
	ExitCommitmentControlClause(c *CommitmentControlClauseContext)

	// ExitDataDivision is called when exiting the dataDivision production.
	ExitDataDivision(c *DataDivisionContext)

	// ExitDataDivisionSection is called when exiting the dataDivisionSection production.
	ExitDataDivisionSection(c *DataDivisionSectionContext)

	// ExitFileSection is called when exiting the fileSection production.
	ExitFileSection(c *FileSectionContext)

	// ExitFileDescriptionEntry is called when exiting the fileDescriptionEntry production.
	ExitFileDescriptionEntry(c *FileDescriptionEntryContext)

	// ExitFileDescriptionEntryClause is called when exiting the fileDescriptionEntryClause production.
	ExitFileDescriptionEntryClause(c *FileDescriptionEntryClauseContext)

	// ExitExternalClause is called when exiting the externalClause production.
	ExitExternalClause(c *ExternalClauseContext)

	// ExitGlobalClause is called when exiting the globalClause production.
	ExitGlobalClause(c *GlobalClauseContext)

	// ExitBlockContainsClause is called when exiting the blockContainsClause production.
	ExitBlockContainsClause(c *BlockContainsClauseContext)

	// ExitBlockContainsTo is called when exiting the blockContainsTo production.
	ExitBlockContainsTo(c *BlockContainsToContext)

	// ExitRecordContainsClause is called when exiting the recordContainsClause production.
	ExitRecordContainsClause(c *RecordContainsClauseContext)

	// ExitRecordContainsClauseFormat1 is called when exiting the recordContainsClauseFormat1 production.
	ExitRecordContainsClauseFormat1(c *RecordContainsClauseFormat1Context)

	// ExitRecordContainsClauseFormat2 is called when exiting the recordContainsClauseFormat2 production.
	ExitRecordContainsClauseFormat2(c *RecordContainsClauseFormat2Context)

	// ExitRecordContainsClauseFormat3 is called when exiting the recordContainsClauseFormat3 production.
	ExitRecordContainsClauseFormat3(c *RecordContainsClauseFormat3Context)

	// ExitRecordContainsTo is called when exiting the recordContainsTo production.
	ExitRecordContainsTo(c *RecordContainsToContext)

	// ExitLabelRecordsClause is called when exiting the labelRecordsClause production.
	ExitLabelRecordsClause(c *LabelRecordsClauseContext)

	// ExitValueOfClause is called when exiting the valueOfClause production.
	ExitValueOfClause(c *ValueOfClauseContext)

	// ExitValuePair is called when exiting the valuePair production.
	ExitValuePair(c *ValuePairContext)

	// ExitDataRecordsClause is called when exiting the dataRecordsClause production.
	ExitDataRecordsClause(c *DataRecordsClauseContext)

	// ExitLinageClause is called when exiting the linageClause production.
	ExitLinageClause(c *LinageClauseContext)

	// ExitLinageAt is called when exiting the linageAt production.
	ExitLinageAt(c *LinageAtContext)

	// ExitLinageFootingAt is called when exiting the linageFootingAt production.
	ExitLinageFootingAt(c *LinageFootingAtContext)

	// ExitLinageLinesAtTop is called when exiting the linageLinesAtTop production.
	ExitLinageLinesAtTop(c *LinageLinesAtTopContext)

	// ExitLinageLinesAtBottom is called when exiting the linageLinesAtBottom production.
	ExitLinageLinesAtBottom(c *LinageLinesAtBottomContext)

	// ExitRecordingModeClause is called when exiting the recordingModeClause production.
	ExitRecordingModeClause(c *RecordingModeClauseContext)

	// ExitModeStatement is called when exiting the modeStatement production.
	ExitModeStatement(c *ModeStatementContext)

	// ExitCodeSetClause is called when exiting the codeSetClause production.
	ExitCodeSetClause(c *CodeSetClauseContext)

	// ExitReportClause is called when exiting the reportClause production.
	ExitReportClause(c *ReportClauseContext)

	// ExitDataBaseSection is called when exiting the dataBaseSection production.
	ExitDataBaseSection(c *DataBaseSectionContext)

	// ExitDataBaseSectionEntry is called when exiting the dataBaseSectionEntry production.
	ExitDataBaseSectionEntry(c *DataBaseSectionEntryContext)

	// ExitWorkingStorageSection is called when exiting the workingStorageSection production.
	ExitWorkingStorageSection(c *WorkingStorageSectionContext)

	// ExitLinkageSection is called when exiting the linkageSection production.
	ExitLinkageSection(c *LinkageSectionContext)

	// ExitCommunicationSection is called when exiting the communicationSection production.
	ExitCommunicationSection(c *CommunicationSectionContext)

	// ExitCommunicationDescriptionEntry is called when exiting the communicationDescriptionEntry production.
	ExitCommunicationDescriptionEntry(c *CommunicationDescriptionEntryContext)

	// ExitCommunicationDescriptionEntryFormat1 is called when exiting the communicationDescriptionEntryFormat1 production.
	ExitCommunicationDescriptionEntryFormat1(c *CommunicationDescriptionEntryFormat1Context)

	// ExitCommunicationDescriptionEntryFormat2 is called when exiting the communicationDescriptionEntryFormat2 production.
	ExitCommunicationDescriptionEntryFormat2(c *CommunicationDescriptionEntryFormat2Context)

	// ExitCommunicationDescriptionEntryFormat3 is called when exiting the communicationDescriptionEntryFormat3 production.
	ExitCommunicationDescriptionEntryFormat3(c *CommunicationDescriptionEntryFormat3Context)

	// ExitDestinationCountClause is called when exiting the destinationCountClause production.
	ExitDestinationCountClause(c *DestinationCountClauseContext)

	// ExitDestinationTableClause is called when exiting the destinationTableClause production.
	ExitDestinationTableClause(c *DestinationTableClauseContext)

	// ExitEndKeyClause is called when exiting the endKeyClause production.
	ExitEndKeyClause(c *EndKeyClauseContext)

	// ExitErrorKeyClause is called when exiting the errorKeyClause production.
	ExitErrorKeyClause(c *ErrorKeyClauseContext)

	// ExitMessageCountClause is called when exiting the messageCountClause production.
	ExitMessageCountClause(c *MessageCountClauseContext)

	// ExitMessageDateClause is called when exiting the messageDateClause production.
	ExitMessageDateClause(c *MessageDateClauseContext)

	// ExitMessageTimeClause is called when exiting the messageTimeClause production.
	ExitMessageTimeClause(c *MessageTimeClauseContext)

	// ExitStatusKeyClause is called when exiting the statusKeyClause production.
	ExitStatusKeyClause(c *StatusKeyClauseContext)

	// ExitSymbolicDestinationClause is called when exiting the symbolicDestinationClause production.
	ExitSymbolicDestinationClause(c *SymbolicDestinationClauseContext)

	// ExitSymbolicQueueClause is called when exiting the symbolicQueueClause production.
	ExitSymbolicQueueClause(c *SymbolicQueueClauseContext)

	// ExitSymbolicSourceClause is called when exiting the symbolicSourceClause production.
	ExitSymbolicSourceClause(c *SymbolicSourceClauseContext)

	// ExitSymbolicTerminalClause is called when exiting the symbolicTerminalClause production.
	ExitSymbolicTerminalClause(c *SymbolicTerminalClauseContext)

	// ExitSymbolicSubQueueClause is called when exiting the symbolicSubQueueClause production.
	ExitSymbolicSubQueueClause(c *SymbolicSubQueueClauseContext)

	// ExitTextLengthClause is called when exiting the textLengthClause production.
	ExitTextLengthClause(c *TextLengthClauseContext)

	// ExitLocalStorageSection is called when exiting the localStorageSection production.
	ExitLocalStorageSection(c *LocalStorageSectionContext)

	// ExitScreenSection is called when exiting the screenSection production.
	ExitScreenSection(c *ScreenSectionContext)

	// ExitScreenDescriptionEntry is called when exiting the screenDescriptionEntry production.
	ExitScreenDescriptionEntry(c *ScreenDescriptionEntryContext)

	// ExitScreenDescriptionBlankClause is called when exiting the screenDescriptionBlankClause production.
	ExitScreenDescriptionBlankClause(c *ScreenDescriptionBlankClauseContext)

	// ExitScreenDescriptionBellClause is called when exiting the screenDescriptionBellClause production.
	ExitScreenDescriptionBellClause(c *ScreenDescriptionBellClauseContext)

	// ExitScreenDescriptionBlinkClause is called when exiting the screenDescriptionBlinkClause production.
	ExitScreenDescriptionBlinkClause(c *ScreenDescriptionBlinkClauseContext)

	// ExitScreenDescriptionEraseClause is called when exiting the screenDescriptionEraseClause production.
	ExitScreenDescriptionEraseClause(c *ScreenDescriptionEraseClauseContext)

	// ExitScreenDescriptionLightClause is called when exiting the screenDescriptionLightClause production.
	ExitScreenDescriptionLightClause(c *ScreenDescriptionLightClauseContext)

	// ExitScreenDescriptionGridClause is called when exiting the screenDescriptionGridClause production.
	ExitScreenDescriptionGridClause(c *ScreenDescriptionGridClauseContext)

	// ExitScreenDescriptionReverseVideoClause is called when exiting the screenDescriptionReverseVideoClause production.
	ExitScreenDescriptionReverseVideoClause(c *ScreenDescriptionReverseVideoClauseContext)

	// ExitScreenDescriptionUnderlineClause is called when exiting the screenDescriptionUnderlineClause production.
	ExitScreenDescriptionUnderlineClause(c *ScreenDescriptionUnderlineClauseContext)

	// ExitScreenDescriptionSizeClause is called when exiting the screenDescriptionSizeClause production.
	ExitScreenDescriptionSizeClause(c *ScreenDescriptionSizeClauseContext)

	// ExitScreenDescriptionLineClause is called when exiting the screenDescriptionLineClause production.
	ExitScreenDescriptionLineClause(c *ScreenDescriptionLineClauseContext)

	// ExitScreenDescriptionColumnClause is called when exiting the screenDescriptionColumnClause production.
	ExitScreenDescriptionColumnClause(c *ScreenDescriptionColumnClauseContext)

	// ExitScreenDescriptionForegroundColorClause is called when exiting the screenDescriptionForegroundColorClause production.
	ExitScreenDescriptionForegroundColorClause(c *ScreenDescriptionForegroundColorClauseContext)

	// ExitScreenDescriptionBackgroundColorClause is called when exiting the screenDescriptionBackgroundColorClause production.
	ExitScreenDescriptionBackgroundColorClause(c *ScreenDescriptionBackgroundColorClauseContext)

	// ExitScreenDescriptionControlClause is called when exiting the screenDescriptionControlClause production.
	ExitScreenDescriptionControlClause(c *ScreenDescriptionControlClauseContext)

	// ExitScreenDescriptionValueClause is called when exiting the screenDescriptionValueClause production.
	ExitScreenDescriptionValueClause(c *ScreenDescriptionValueClauseContext)

	// ExitScreenDescriptionPictureClause is called when exiting the screenDescriptionPictureClause production.
	ExitScreenDescriptionPictureClause(c *ScreenDescriptionPictureClauseContext)

	// ExitScreenDescriptionFromClause is called when exiting the screenDescriptionFromClause production.
	ExitScreenDescriptionFromClause(c *ScreenDescriptionFromClauseContext)

	// ExitScreenDescriptionToClause is called when exiting the screenDescriptionToClause production.
	ExitScreenDescriptionToClause(c *ScreenDescriptionToClauseContext)

	// ExitScreenDescriptionUsingClause is called when exiting the screenDescriptionUsingClause production.
	ExitScreenDescriptionUsingClause(c *ScreenDescriptionUsingClauseContext)

	// ExitScreenDescriptionUsageClause is called when exiting the screenDescriptionUsageClause production.
	ExitScreenDescriptionUsageClause(c *ScreenDescriptionUsageClauseContext)

	// ExitScreenDescriptionBlankWhenZeroClause is called when exiting the screenDescriptionBlankWhenZeroClause production.
	ExitScreenDescriptionBlankWhenZeroClause(c *ScreenDescriptionBlankWhenZeroClauseContext)

	// ExitScreenDescriptionJustifiedClause is called when exiting the screenDescriptionJustifiedClause production.
	ExitScreenDescriptionJustifiedClause(c *ScreenDescriptionJustifiedClauseContext)

	// ExitScreenDescriptionSignClause is called when exiting the screenDescriptionSignClause production.
	ExitScreenDescriptionSignClause(c *ScreenDescriptionSignClauseContext)

	// ExitScreenDescriptionAutoClause is called when exiting the screenDescriptionAutoClause production.
	ExitScreenDescriptionAutoClause(c *ScreenDescriptionAutoClauseContext)

	// ExitScreenDescriptionSecureClause is called when exiting the screenDescriptionSecureClause production.
	ExitScreenDescriptionSecureClause(c *ScreenDescriptionSecureClauseContext)

	// ExitScreenDescriptionRequiredClause is called when exiting the screenDescriptionRequiredClause production.
	ExitScreenDescriptionRequiredClause(c *ScreenDescriptionRequiredClauseContext)

	// ExitScreenDescriptionPromptClause is called when exiting the screenDescriptionPromptClause production.
	ExitScreenDescriptionPromptClause(c *ScreenDescriptionPromptClauseContext)

	// ExitScreenDescriptionPromptOccursClause is called when exiting the screenDescriptionPromptOccursClause production.
	ExitScreenDescriptionPromptOccursClause(c *ScreenDescriptionPromptOccursClauseContext)

	// ExitScreenDescriptionFullClause is called when exiting the screenDescriptionFullClause production.
	ExitScreenDescriptionFullClause(c *ScreenDescriptionFullClauseContext)

	// ExitScreenDescriptionZeroFillClause is called when exiting the screenDescriptionZeroFillClause production.
	ExitScreenDescriptionZeroFillClause(c *ScreenDescriptionZeroFillClauseContext)

	// ExitReportSection is called when exiting the reportSection production.
	ExitReportSection(c *ReportSectionContext)

	// ExitReportDescription is called when exiting the reportDescription production.
	ExitReportDescription(c *ReportDescriptionContext)

	// ExitReportDescriptionEntry is called when exiting the reportDescriptionEntry production.
	ExitReportDescriptionEntry(c *ReportDescriptionEntryContext)

	// ExitReportDescriptionGlobalClause is called when exiting the reportDescriptionGlobalClause production.
	ExitReportDescriptionGlobalClause(c *ReportDescriptionGlobalClauseContext)

	// ExitReportDescriptionPageLimitClause is called when exiting the reportDescriptionPageLimitClause production.
	ExitReportDescriptionPageLimitClause(c *ReportDescriptionPageLimitClauseContext)

	// ExitReportDescriptionHeadingClause is called when exiting the reportDescriptionHeadingClause production.
	ExitReportDescriptionHeadingClause(c *ReportDescriptionHeadingClauseContext)

	// ExitReportDescriptionFirstDetailClause is called when exiting the reportDescriptionFirstDetailClause production.
	ExitReportDescriptionFirstDetailClause(c *ReportDescriptionFirstDetailClauseContext)

	// ExitReportDescriptionLastDetailClause is called when exiting the reportDescriptionLastDetailClause production.
	ExitReportDescriptionLastDetailClause(c *ReportDescriptionLastDetailClauseContext)

	// ExitReportDescriptionFootingClause is called when exiting the reportDescriptionFootingClause production.
	ExitReportDescriptionFootingClause(c *ReportDescriptionFootingClauseContext)

	// ExitReportGroupDescriptionEntry is called when exiting the reportGroupDescriptionEntry production.
	ExitReportGroupDescriptionEntry(c *ReportGroupDescriptionEntryContext)

	// ExitReportGroupDescriptionEntryFormat1 is called when exiting the reportGroupDescriptionEntryFormat1 production.
	ExitReportGroupDescriptionEntryFormat1(c *ReportGroupDescriptionEntryFormat1Context)

	// ExitReportGroupDescriptionEntryFormat2 is called when exiting the reportGroupDescriptionEntryFormat2 production.
	ExitReportGroupDescriptionEntryFormat2(c *ReportGroupDescriptionEntryFormat2Context)

	// ExitReportGroupDescriptionEntryFormat3 is called when exiting the reportGroupDescriptionEntryFormat3 production.
	ExitReportGroupDescriptionEntryFormat3(c *ReportGroupDescriptionEntryFormat3Context)

	// ExitReportGroupBlankWhenZeroClause is called when exiting the reportGroupBlankWhenZeroClause production.
	ExitReportGroupBlankWhenZeroClause(c *ReportGroupBlankWhenZeroClauseContext)

	// ExitReportGroupColumnNumberClause is called when exiting the reportGroupColumnNumberClause production.
	ExitReportGroupColumnNumberClause(c *ReportGroupColumnNumberClauseContext)

	// ExitReportGroupIndicateClause is called when exiting the reportGroupIndicateClause production.
	ExitReportGroupIndicateClause(c *ReportGroupIndicateClauseContext)

	// ExitReportGroupJustifiedClause is called when exiting the reportGroupJustifiedClause production.
	ExitReportGroupJustifiedClause(c *ReportGroupJustifiedClauseContext)

	// ExitReportGroupLineNumberClause is called when exiting the reportGroupLineNumberClause production.
	ExitReportGroupLineNumberClause(c *ReportGroupLineNumberClauseContext)

	// ExitReportGroupLineNumberNextPage is called when exiting the reportGroupLineNumberNextPage production.
	ExitReportGroupLineNumberNextPage(c *ReportGroupLineNumberNextPageContext)

	// ExitReportGroupLineNumberPlus is called when exiting the reportGroupLineNumberPlus production.
	ExitReportGroupLineNumberPlus(c *ReportGroupLineNumberPlusContext)

	// ExitReportGroupNextGroupClause is called when exiting the reportGroupNextGroupClause production.
	ExitReportGroupNextGroupClause(c *ReportGroupNextGroupClauseContext)

	// ExitReportGroupNextGroupPlus is called when exiting the reportGroupNextGroupPlus production.
	ExitReportGroupNextGroupPlus(c *ReportGroupNextGroupPlusContext)

	// ExitReportGroupNextGroupNextPage is called when exiting the reportGroupNextGroupNextPage production.
	ExitReportGroupNextGroupNextPage(c *ReportGroupNextGroupNextPageContext)

	// ExitReportGroupPictureClause is called when exiting the reportGroupPictureClause production.
	ExitReportGroupPictureClause(c *ReportGroupPictureClauseContext)

	// ExitReportGroupResetClause is called when exiting the reportGroupResetClause production.
	ExitReportGroupResetClause(c *ReportGroupResetClauseContext)

	// ExitReportGroupSignClause is called when exiting the reportGroupSignClause production.
	ExitReportGroupSignClause(c *ReportGroupSignClauseContext)

	// ExitReportGroupSourceClause is called when exiting the reportGroupSourceClause production.
	ExitReportGroupSourceClause(c *ReportGroupSourceClauseContext)

	// ExitReportGroupSumClause is called when exiting the reportGroupSumClause production.
	ExitReportGroupSumClause(c *ReportGroupSumClauseContext)

	// ExitReportGroupTypeClause is called when exiting the reportGroupTypeClause production.
	ExitReportGroupTypeClause(c *ReportGroupTypeClauseContext)

	// ExitReportGroupTypeReportHeading is called when exiting the reportGroupTypeReportHeading production.
	ExitReportGroupTypeReportHeading(c *ReportGroupTypeReportHeadingContext)

	// ExitReportGroupTypePageHeading is called when exiting the reportGroupTypePageHeading production.
	ExitReportGroupTypePageHeading(c *ReportGroupTypePageHeadingContext)

	// ExitReportGroupTypeControlHeading is called when exiting the reportGroupTypeControlHeading production.
	ExitReportGroupTypeControlHeading(c *ReportGroupTypeControlHeadingContext)

	// ExitReportGroupTypeDetail is called when exiting the reportGroupTypeDetail production.
	ExitReportGroupTypeDetail(c *ReportGroupTypeDetailContext)

	// ExitReportGroupTypeControlFooting is called when exiting the reportGroupTypeControlFooting production.
	ExitReportGroupTypeControlFooting(c *ReportGroupTypeControlFootingContext)

	// ExitReportGroupUsageClause is called when exiting the reportGroupUsageClause production.
	ExitReportGroupUsageClause(c *ReportGroupUsageClauseContext)

	// ExitReportGroupTypePageFooting is called when exiting the reportGroupTypePageFooting production.
	ExitReportGroupTypePageFooting(c *ReportGroupTypePageFootingContext)

	// ExitReportGroupTypeReportFooting is called when exiting the reportGroupTypeReportFooting production.
	ExitReportGroupTypeReportFooting(c *ReportGroupTypeReportFootingContext)

	// ExitReportGroupValueClause is called when exiting the reportGroupValueClause production.
	ExitReportGroupValueClause(c *ReportGroupValueClauseContext)

	// ExitProgramLibrarySection is called when exiting the programLibrarySection production.
	ExitProgramLibrarySection(c *ProgramLibrarySectionContext)

	// ExitLibraryDescriptionEntry is called when exiting the libraryDescriptionEntry production.
	ExitLibraryDescriptionEntry(c *LibraryDescriptionEntryContext)

	// ExitLibraryDescriptionEntryFormat1 is called when exiting the libraryDescriptionEntryFormat1 production.
	ExitLibraryDescriptionEntryFormat1(c *LibraryDescriptionEntryFormat1Context)

	// ExitLibraryDescriptionEntryFormat2 is called when exiting the libraryDescriptionEntryFormat2 production.
	ExitLibraryDescriptionEntryFormat2(c *LibraryDescriptionEntryFormat2Context)

	// ExitLibraryAttributeClauseFormat1 is called when exiting the libraryAttributeClauseFormat1 production.
	ExitLibraryAttributeClauseFormat1(c *LibraryAttributeClauseFormat1Context)

	// ExitLibraryAttributeClauseFormat2 is called when exiting the libraryAttributeClauseFormat2 production.
	ExitLibraryAttributeClauseFormat2(c *LibraryAttributeClauseFormat2Context)

	// ExitLibraryAttributeFunction is called when exiting the libraryAttributeFunction production.
	ExitLibraryAttributeFunction(c *LibraryAttributeFunctionContext)

	// ExitLibraryAttributeParameter is called when exiting the libraryAttributeParameter production.
	ExitLibraryAttributeParameter(c *LibraryAttributeParameterContext)

	// ExitLibraryAttributeTitle is called when exiting the libraryAttributeTitle production.
	ExitLibraryAttributeTitle(c *LibraryAttributeTitleContext)

	// ExitLibraryEntryProcedureClauseFormat1 is called when exiting the libraryEntryProcedureClauseFormat1 production.
	ExitLibraryEntryProcedureClauseFormat1(c *LibraryEntryProcedureClauseFormat1Context)

	// ExitLibraryEntryProcedureClauseFormat2 is called when exiting the libraryEntryProcedureClauseFormat2 production.
	ExitLibraryEntryProcedureClauseFormat2(c *LibraryEntryProcedureClauseFormat2Context)

	// ExitLibraryEntryProcedureForClause is called when exiting the libraryEntryProcedureForClause production.
	ExitLibraryEntryProcedureForClause(c *LibraryEntryProcedureForClauseContext)

	// ExitLibraryEntryProcedureGivingClause is called when exiting the libraryEntryProcedureGivingClause production.
	ExitLibraryEntryProcedureGivingClause(c *LibraryEntryProcedureGivingClauseContext)

	// ExitLibraryEntryProcedureUsingClause is called when exiting the libraryEntryProcedureUsingClause production.
	ExitLibraryEntryProcedureUsingClause(c *LibraryEntryProcedureUsingClauseContext)

	// ExitLibraryEntryProcedureUsingName is called when exiting the libraryEntryProcedureUsingName production.
	ExitLibraryEntryProcedureUsingName(c *LibraryEntryProcedureUsingNameContext)

	// ExitLibraryEntryProcedureWithClause is called when exiting the libraryEntryProcedureWithClause production.
	ExitLibraryEntryProcedureWithClause(c *LibraryEntryProcedureWithClauseContext)

	// ExitLibraryEntryProcedureWithName is called when exiting the libraryEntryProcedureWithName production.
	ExitLibraryEntryProcedureWithName(c *LibraryEntryProcedureWithNameContext)

	// ExitLibraryIsCommonClause is called when exiting the libraryIsCommonClause production.
	ExitLibraryIsCommonClause(c *LibraryIsCommonClauseContext)

	// ExitLibraryIsGlobalClause is called when exiting the libraryIsGlobalClause production.
	ExitLibraryIsGlobalClause(c *LibraryIsGlobalClauseContext)

	// ExitDataDescriptionEntry is called when exiting the dataDescriptionEntry production.
	ExitDataDescriptionEntry(c *DataDescriptionEntryContext)

	// ExitDataDescriptionEntryFormat1 is called when exiting the dataDescriptionEntryFormat1 production.
	ExitDataDescriptionEntryFormat1(c *DataDescriptionEntryFormat1Context)

	// ExitDataDescriptionEntryFormat2 is called when exiting the dataDescriptionEntryFormat2 production.
	ExitDataDescriptionEntryFormat2(c *DataDescriptionEntryFormat2Context)

	// ExitDataDescriptionEntryFormat3 is called when exiting the dataDescriptionEntryFormat3 production.
	ExitDataDescriptionEntryFormat3(c *DataDescriptionEntryFormat3Context)

	// ExitDataDescriptionEntryExecSql is called when exiting the dataDescriptionEntryExecSql production.
	ExitDataDescriptionEntryExecSql(c *DataDescriptionEntryExecSqlContext)

	// ExitDataAlignedClause is called when exiting the dataAlignedClause production.
	ExitDataAlignedClause(c *DataAlignedClauseContext)

	// ExitDataBlankWhenZeroClause is called when exiting the dataBlankWhenZeroClause production.
	ExitDataBlankWhenZeroClause(c *DataBlankWhenZeroClauseContext)

	// ExitDataCommonOwnLocalClause is called when exiting the dataCommonOwnLocalClause production.
	ExitDataCommonOwnLocalClause(c *DataCommonOwnLocalClauseContext)

	// ExitDataExternalClause is called when exiting the dataExternalClause production.
	ExitDataExternalClause(c *DataExternalClauseContext)

	// ExitDataGlobalClause is called when exiting the dataGlobalClause production.
	ExitDataGlobalClause(c *DataGlobalClauseContext)

	// ExitDataIntegerStringClause is called when exiting the dataIntegerStringClause production.
	ExitDataIntegerStringClause(c *DataIntegerStringClauseContext)

	// ExitDataJustifiedClause is called when exiting the dataJustifiedClause production.
	ExitDataJustifiedClause(c *DataJustifiedClauseContext)

	// ExitDataOccursClause is called when exiting the dataOccursClause production.
	ExitDataOccursClause(c *DataOccursClauseContext)

	// ExitDataOccursTo is called when exiting the dataOccursTo production.
	ExitDataOccursTo(c *DataOccursToContext)

	// ExitDataOccursSort is called when exiting the dataOccursSort production.
	ExitDataOccursSort(c *DataOccursSortContext)

	// ExitDataPictureClause is called when exiting the dataPictureClause production.
	ExitDataPictureClause(c *DataPictureClauseContext)

	// ExitPictureString is called when exiting the pictureString production.
	ExitPictureString(c *PictureStringContext)

	// ExitPictureChars is called when exiting the pictureChars production.
	ExitPictureChars(c *PictureCharsContext)

	// ExitPictureCardinality is called when exiting the pictureCardinality production.
	ExitPictureCardinality(c *PictureCardinalityContext)

	// ExitDataReceivedByClause is called when exiting the dataReceivedByClause production.
	ExitDataReceivedByClause(c *DataReceivedByClauseContext)

	// ExitDataRecordAreaClause is called when exiting the dataRecordAreaClause production.
	ExitDataRecordAreaClause(c *DataRecordAreaClauseContext)

	// ExitDataRedefinesClause is called when exiting the dataRedefinesClause production.
	ExitDataRedefinesClause(c *DataRedefinesClauseContext)

	// ExitDataRenamesClause is called when exiting the dataRenamesClause production.
	ExitDataRenamesClause(c *DataRenamesClauseContext)

	// ExitDataSignClause is called when exiting the dataSignClause production.
	ExitDataSignClause(c *DataSignClauseContext)

	// ExitDataSynchronizedClause is called when exiting the dataSynchronizedClause production.
	ExitDataSynchronizedClause(c *DataSynchronizedClauseContext)

	// ExitDataThreadLocalClause is called when exiting the dataThreadLocalClause production.
	ExitDataThreadLocalClause(c *DataThreadLocalClauseContext)

	// ExitDataTypeClause is called when exiting the dataTypeClause production.
	ExitDataTypeClause(c *DataTypeClauseContext)

	// ExitDataTypeDefClause is called when exiting the dataTypeDefClause production.
	ExitDataTypeDefClause(c *DataTypeDefClauseContext)

	// ExitDataUsageClause is called when exiting the dataUsageClause production.
	ExitDataUsageClause(c *DataUsageClauseContext)

	// ExitDataUsingClause is called when exiting the dataUsingClause production.
	ExitDataUsingClause(c *DataUsingClauseContext)

	// ExitDataValueClause is called when exiting the dataValueClause production.
	ExitDataValueClause(c *DataValueClauseContext)

	// ExitDataValueInterval is called when exiting the dataValueInterval production.
	ExitDataValueInterval(c *DataValueIntervalContext)

	// ExitDataValueIntervalFrom is called when exiting the dataValueIntervalFrom production.
	ExitDataValueIntervalFrom(c *DataValueIntervalFromContext)

	// ExitDataValueIntervalTo is called when exiting the dataValueIntervalTo production.
	ExitDataValueIntervalTo(c *DataValueIntervalToContext)

	// ExitDataWithLowerBoundsClause is called when exiting the dataWithLowerBoundsClause production.
	ExitDataWithLowerBoundsClause(c *DataWithLowerBoundsClauseContext)

	// ExitProcedureDivision is called when exiting the procedureDivision production.
	ExitProcedureDivision(c *ProcedureDivisionContext)

	// ExitProcedureDivisionUsingClause is called when exiting the procedureDivisionUsingClause production.
	ExitProcedureDivisionUsingClause(c *ProcedureDivisionUsingClauseContext)

	// ExitProcedureDivisionGivingClause is called when exiting the procedureDivisionGivingClause production.
	ExitProcedureDivisionGivingClause(c *ProcedureDivisionGivingClauseContext)

	// ExitProcedureDivisionUsingParameter is called when exiting the procedureDivisionUsingParameter production.
	ExitProcedureDivisionUsingParameter(c *ProcedureDivisionUsingParameterContext)

	// ExitProcedureDivisionByReferencePhrase is called when exiting the procedureDivisionByReferencePhrase production.
	ExitProcedureDivisionByReferencePhrase(c *ProcedureDivisionByReferencePhraseContext)

	// ExitProcedureDivisionByReference is called when exiting the procedureDivisionByReference production.
	ExitProcedureDivisionByReference(c *ProcedureDivisionByReferenceContext)

	// ExitProcedureDivisionByValuePhrase is called when exiting the procedureDivisionByValuePhrase production.
	ExitProcedureDivisionByValuePhrase(c *ProcedureDivisionByValuePhraseContext)

	// ExitProcedureDivisionByValue is called when exiting the procedureDivisionByValue production.
	ExitProcedureDivisionByValue(c *ProcedureDivisionByValueContext)

	// ExitProcedureDeclaratives is called when exiting the procedureDeclaratives production.
	ExitProcedureDeclaratives(c *ProcedureDeclarativesContext)

	// ExitProcedureDeclarative is called when exiting the procedureDeclarative production.
	ExitProcedureDeclarative(c *ProcedureDeclarativeContext)

	// ExitProcedureSectionHeader is called when exiting the procedureSectionHeader production.
	ExitProcedureSectionHeader(c *ProcedureSectionHeaderContext)

	// ExitProcedureDivisionBody is called when exiting the procedureDivisionBody production.
	ExitProcedureDivisionBody(c *ProcedureDivisionBodyContext)

	// ExitProcedureSection is called when exiting the procedureSection production.
	ExitProcedureSection(c *ProcedureSectionContext)

	// ExitParagraphs is called when exiting the paragraphs production.
	ExitParagraphs(c *ParagraphsContext)

	// ExitParagraph is called when exiting the paragraph production.
	ExitParagraph(c *ParagraphContext)

	// ExitSentence is called when exiting the sentence production.
	ExitSentence(c *SentenceContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAcceptStatement is called when exiting the acceptStatement production.
	ExitAcceptStatement(c *AcceptStatementContext)

	// ExitAcceptFromDateStatement is called when exiting the acceptFromDateStatement production.
	ExitAcceptFromDateStatement(c *AcceptFromDateStatementContext)

	// ExitAcceptFromMnemonicStatement is called when exiting the acceptFromMnemonicStatement production.
	ExitAcceptFromMnemonicStatement(c *AcceptFromMnemonicStatementContext)

	// ExitAcceptFromEscapeKeyStatement is called when exiting the acceptFromEscapeKeyStatement production.
	ExitAcceptFromEscapeKeyStatement(c *AcceptFromEscapeKeyStatementContext)

	// ExitAcceptMessageCountStatement is called when exiting the acceptMessageCountStatement production.
	ExitAcceptMessageCountStatement(c *AcceptMessageCountStatementContext)

	// ExitAddStatement is called when exiting the addStatement production.
	ExitAddStatement(c *AddStatementContext)

	// ExitAddToStatement is called when exiting the addToStatement production.
	ExitAddToStatement(c *AddToStatementContext)

	// ExitAddToGivingStatement is called when exiting the addToGivingStatement production.
	ExitAddToGivingStatement(c *AddToGivingStatementContext)

	// ExitAddCorrespondingStatement is called when exiting the addCorrespondingStatement production.
	ExitAddCorrespondingStatement(c *AddCorrespondingStatementContext)

	// ExitAddFrom is called when exiting the addFrom production.
	ExitAddFrom(c *AddFromContext)

	// ExitAddTo is called when exiting the addTo production.
	ExitAddTo(c *AddToContext)

	// ExitAddToGiving is called when exiting the addToGiving production.
	ExitAddToGiving(c *AddToGivingContext)

	// ExitAddGiving is called when exiting the addGiving production.
	ExitAddGiving(c *AddGivingContext)

	// ExitAlteredGoTo is called when exiting the alteredGoTo production.
	ExitAlteredGoTo(c *AlteredGoToContext)

	// ExitAlterStatement is called when exiting the alterStatement production.
	ExitAlterStatement(c *AlterStatementContext)

	// ExitAlterProceedTo is called when exiting the alterProceedTo production.
	ExitAlterProceedTo(c *AlterProceedToContext)

	// ExitCallStatement is called when exiting the callStatement production.
	ExitCallStatement(c *CallStatementContext)

	// ExitCallUsingPhrase is called when exiting the callUsingPhrase production.
	ExitCallUsingPhrase(c *CallUsingPhraseContext)

	// ExitCallUsingParameter is called when exiting the callUsingParameter production.
	ExitCallUsingParameter(c *CallUsingParameterContext)

	// ExitCallByReferencePhrase is called when exiting the callByReferencePhrase production.
	ExitCallByReferencePhrase(c *CallByReferencePhraseContext)

	// ExitCallByReference is called when exiting the callByReference production.
	ExitCallByReference(c *CallByReferenceContext)

	// ExitCallByValuePhrase is called when exiting the callByValuePhrase production.
	ExitCallByValuePhrase(c *CallByValuePhraseContext)

	// ExitCallByValue is called when exiting the callByValue production.
	ExitCallByValue(c *CallByValueContext)

	// ExitCallByContentPhrase is called when exiting the callByContentPhrase production.
	ExitCallByContentPhrase(c *CallByContentPhraseContext)

	// ExitCallByContent is called when exiting the callByContent production.
	ExitCallByContent(c *CallByContentContext)

	// ExitCallGivingPhrase is called when exiting the callGivingPhrase production.
	ExitCallGivingPhrase(c *CallGivingPhraseContext)

	// ExitCancelStatement is called when exiting the cancelStatement production.
	ExitCancelStatement(c *CancelStatementContext)

	// ExitCancelCall is called when exiting the cancelCall production.
	ExitCancelCall(c *CancelCallContext)

	// ExitCloseStatement is called when exiting the closeStatement production.
	ExitCloseStatement(c *CloseStatementContext)

	// ExitCloseFile is called when exiting the closeFile production.
	ExitCloseFile(c *CloseFileContext)

	// ExitCloseReelUnitStatement is called when exiting the closeReelUnitStatement production.
	ExitCloseReelUnitStatement(c *CloseReelUnitStatementContext)

	// ExitCloseRelativeStatement is called when exiting the closeRelativeStatement production.
	ExitCloseRelativeStatement(c *CloseRelativeStatementContext)

	// ExitClosePortFileIOStatement is called when exiting the closePortFileIOStatement production.
	ExitClosePortFileIOStatement(c *ClosePortFileIOStatementContext)

	// ExitClosePortFileIOUsing is called when exiting the closePortFileIOUsing production.
	ExitClosePortFileIOUsing(c *ClosePortFileIOUsingContext)

	// ExitClosePortFileIOUsingCloseDisposition is called when exiting the closePortFileIOUsingCloseDisposition production.
	ExitClosePortFileIOUsingCloseDisposition(c *ClosePortFileIOUsingCloseDispositionContext)

	// ExitClosePortFileIOUsingAssociatedData is called when exiting the closePortFileIOUsingAssociatedData production.
	ExitClosePortFileIOUsingAssociatedData(c *ClosePortFileIOUsingAssociatedDataContext)

	// ExitClosePortFileIOUsingAssociatedDataLength is called when exiting the closePortFileIOUsingAssociatedDataLength production.
	ExitClosePortFileIOUsingAssociatedDataLength(c *ClosePortFileIOUsingAssociatedDataLengthContext)

	// ExitComputeStatement is called when exiting the computeStatement production.
	ExitComputeStatement(c *ComputeStatementContext)

	// ExitComputeStore is called when exiting the computeStore production.
	ExitComputeStore(c *ComputeStoreContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitDisableStatement is called when exiting the disableStatement production.
	ExitDisableStatement(c *DisableStatementContext)

	// ExitDisplayStatement is called when exiting the displayStatement production.
	ExitDisplayStatement(c *DisplayStatementContext)

	// ExitDisplayOperand is called when exiting the displayOperand production.
	ExitDisplayOperand(c *DisplayOperandContext)

	// ExitDisplayAt is called when exiting the displayAt production.
	ExitDisplayAt(c *DisplayAtContext)

	// ExitDisplayUpon is called when exiting the displayUpon production.
	ExitDisplayUpon(c *DisplayUponContext)

	// ExitDisplayWith is called when exiting the displayWith production.
	ExitDisplayWith(c *DisplayWithContext)

	// ExitDivideStatement is called when exiting the divideStatement production.
	ExitDivideStatement(c *DivideStatementContext)

	// ExitDivideIntoStatement is called when exiting the divideIntoStatement production.
	ExitDivideIntoStatement(c *DivideIntoStatementContext)

	// ExitDivideIntoGivingStatement is called when exiting the divideIntoGivingStatement production.
	ExitDivideIntoGivingStatement(c *DivideIntoGivingStatementContext)

	// ExitDivideByGivingStatement is called when exiting the divideByGivingStatement production.
	ExitDivideByGivingStatement(c *DivideByGivingStatementContext)

	// ExitDivideGivingPhrase is called when exiting the divideGivingPhrase production.
	ExitDivideGivingPhrase(c *DivideGivingPhraseContext)

	// ExitDivideInto is called when exiting the divideInto production.
	ExitDivideInto(c *DivideIntoContext)

	// ExitDivideGiving is called when exiting the divideGiving production.
	ExitDivideGiving(c *DivideGivingContext)

	// ExitDivideRemainder is called when exiting the divideRemainder production.
	ExitDivideRemainder(c *DivideRemainderContext)

	// ExitEnableStatement is called when exiting the enableStatement production.
	ExitEnableStatement(c *EnableStatementContext)

	// ExitEntryStatement is called when exiting the entryStatement production.
	ExitEntryStatement(c *EntryStatementContext)

	// ExitEvaluateStatement is called when exiting the evaluateStatement production.
	ExitEvaluateStatement(c *EvaluateStatementContext)

	// ExitEvaluateSelect is called when exiting the evaluateSelect production.
	ExitEvaluateSelect(c *EvaluateSelectContext)

	// ExitEvaluateAlsoSelect is called when exiting the evaluateAlsoSelect production.
	ExitEvaluateAlsoSelect(c *EvaluateAlsoSelectContext)

	// ExitEvaluateWhenPhrase is called when exiting the evaluateWhenPhrase production.
	ExitEvaluateWhenPhrase(c *EvaluateWhenPhraseContext)

	// ExitEvaluateWhen is called when exiting the evaluateWhen production.
	ExitEvaluateWhen(c *EvaluateWhenContext)

	// ExitEvaluateCondition is called when exiting the evaluateCondition production.
	ExitEvaluateCondition(c *EvaluateConditionContext)

	// ExitEvaluateThrough is called when exiting the evaluateThrough production.
	ExitEvaluateThrough(c *EvaluateThroughContext)

	// ExitEvaluateAlsoCondition is called when exiting the evaluateAlsoCondition production.
	ExitEvaluateAlsoCondition(c *EvaluateAlsoConditionContext)

	// ExitEvaluateWhenOther is called when exiting the evaluateWhenOther production.
	ExitEvaluateWhenOther(c *EvaluateWhenOtherContext)

	// ExitEvaluateValue is called when exiting the evaluateValue production.
	ExitEvaluateValue(c *EvaluateValueContext)

	// ExitExecCicsStatement is called when exiting the execCicsStatement production.
	ExitExecCicsStatement(c *ExecCicsStatementContext)

	// ExitExecSqlStatement is called when exiting the execSqlStatement production.
	ExitExecSqlStatement(c *ExecSqlStatementContext)

	// ExitExecSqlImsStatement is called when exiting the execSqlImsStatement production.
	ExitExecSqlImsStatement(c *ExecSqlImsStatementContext)

	// ExitExhibitStatement is called when exiting the exhibitStatement production.
	ExitExhibitStatement(c *ExhibitStatementContext)

	// ExitExhibitOperand is called when exiting the exhibitOperand production.
	ExitExhibitOperand(c *ExhibitOperandContext)

	// ExitExitStatement is called when exiting the exitStatement production.
	ExitExitStatement(c *ExitStatementContext)

	// ExitGenerateStatement is called when exiting the generateStatement production.
	ExitGenerateStatement(c *GenerateStatementContext)

	// ExitGobackStatement is called when exiting the gobackStatement production.
	ExitGobackStatement(c *GobackStatementContext)

	// ExitGoToStatement is called when exiting the goToStatement production.
	ExitGoToStatement(c *GoToStatementContext)

	// ExitGoToStatementSimple is called when exiting the goToStatementSimple production.
	ExitGoToStatementSimple(c *GoToStatementSimpleContext)

	// ExitGoToDependingOnStatement is called when exiting the goToDependingOnStatement production.
	ExitGoToDependingOnStatement(c *GoToDependingOnStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIfThen is called when exiting the ifThen production.
	ExitIfThen(c *IfThenContext)

	// ExitIfElse is called when exiting the ifElse production.
	ExitIfElse(c *IfElseContext)

	// ExitInitializeStatement is called when exiting the initializeStatement production.
	ExitInitializeStatement(c *InitializeStatementContext)

	// ExitInitializeReplacingPhrase is called when exiting the initializeReplacingPhrase production.
	ExitInitializeReplacingPhrase(c *InitializeReplacingPhraseContext)

	// ExitInitializeReplacingBy is called when exiting the initializeReplacingBy production.
	ExitInitializeReplacingBy(c *InitializeReplacingByContext)

	// ExitInitiateStatement is called when exiting the initiateStatement production.
	ExitInitiateStatement(c *InitiateStatementContext)

	// ExitInspectStatement is called when exiting the inspectStatement production.
	ExitInspectStatement(c *InspectStatementContext)

	// ExitInspectTallyingPhrase is called when exiting the inspectTallyingPhrase production.
	ExitInspectTallyingPhrase(c *InspectTallyingPhraseContext)

	// ExitInspectReplacingPhrase is called when exiting the inspectReplacingPhrase production.
	ExitInspectReplacingPhrase(c *InspectReplacingPhraseContext)

	// ExitInspectTallyingReplacingPhrase is called when exiting the inspectTallyingReplacingPhrase production.
	ExitInspectTallyingReplacingPhrase(c *InspectTallyingReplacingPhraseContext)

	// ExitInspectConvertingPhrase is called when exiting the inspectConvertingPhrase production.
	ExitInspectConvertingPhrase(c *InspectConvertingPhraseContext)

	// ExitInspectFor is called when exiting the inspectFor production.
	ExitInspectFor(c *InspectForContext)

	// ExitInspectCharacters is called when exiting the inspectCharacters production.
	ExitInspectCharacters(c *InspectCharactersContext)

	// ExitInspectReplacingCharacters is called when exiting the inspectReplacingCharacters production.
	ExitInspectReplacingCharacters(c *InspectReplacingCharactersContext)

	// ExitInspectAllLeadings is called when exiting the inspectAllLeadings production.
	ExitInspectAllLeadings(c *InspectAllLeadingsContext)

	// ExitInspectReplacingAllLeadings is called when exiting the inspectReplacingAllLeadings production.
	ExitInspectReplacingAllLeadings(c *InspectReplacingAllLeadingsContext)

	// ExitInspectAllLeading is called when exiting the inspectAllLeading production.
	ExitInspectAllLeading(c *InspectAllLeadingContext)

	// ExitInspectReplacingAllLeading is called when exiting the inspectReplacingAllLeading production.
	ExitInspectReplacingAllLeading(c *InspectReplacingAllLeadingContext)

	// ExitInspectBy is called when exiting the inspectBy production.
	ExitInspectBy(c *InspectByContext)

	// ExitInspectTo is called when exiting the inspectTo production.
	ExitInspectTo(c *InspectToContext)

	// ExitInspectBeforeAfter is called when exiting the inspectBeforeAfter production.
	ExitInspectBeforeAfter(c *InspectBeforeAfterContext)

	// ExitMergeStatement is called when exiting the mergeStatement production.
	ExitMergeStatement(c *MergeStatementContext)

	// ExitMergeOnKeyClause is called when exiting the mergeOnKeyClause production.
	ExitMergeOnKeyClause(c *MergeOnKeyClauseContext)

	// ExitMergeCollatingSequencePhrase is called when exiting the mergeCollatingSequencePhrase production.
	ExitMergeCollatingSequencePhrase(c *MergeCollatingSequencePhraseContext)

	// ExitMergeCollatingAlphanumeric is called when exiting the mergeCollatingAlphanumeric production.
	ExitMergeCollatingAlphanumeric(c *MergeCollatingAlphanumericContext)

	// ExitMergeCollatingNational is called when exiting the mergeCollatingNational production.
	ExitMergeCollatingNational(c *MergeCollatingNationalContext)

	// ExitMergeUsing is called when exiting the mergeUsing production.
	ExitMergeUsing(c *MergeUsingContext)

	// ExitMergeOutputProcedurePhrase is called when exiting the mergeOutputProcedurePhrase production.
	ExitMergeOutputProcedurePhrase(c *MergeOutputProcedurePhraseContext)

	// ExitMergeOutputThrough is called when exiting the mergeOutputThrough production.
	ExitMergeOutputThrough(c *MergeOutputThroughContext)

	// ExitMergeGivingPhrase is called when exiting the mergeGivingPhrase production.
	ExitMergeGivingPhrase(c *MergeGivingPhraseContext)

	// ExitMergeGiving is called when exiting the mergeGiving production.
	ExitMergeGiving(c *MergeGivingContext)

	// ExitMoveStatement is called when exiting the moveStatement production.
	ExitMoveStatement(c *MoveStatementContext)

	// ExitMoveToStatement is called when exiting the moveToStatement production.
	ExitMoveToStatement(c *MoveToStatementContext)

	// ExitMoveToSendingArea is called when exiting the moveToSendingArea production.
	ExitMoveToSendingArea(c *MoveToSendingAreaContext)

	// ExitMoveCorrespondingToStatement is called when exiting the moveCorrespondingToStatement production.
	ExitMoveCorrespondingToStatement(c *MoveCorrespondingToStatementContext)

	// ExitMoveCorrespondingToSendingArea is called when exiting the moveCorrespondingToSendingArea production.
	ExitMoveCorrespondingToSendingArea(c *MoveCorrespondingToSendingAreaContext)

	// ExitMultiplyStatement is called when exiting the multiplyStatement production.
	ExitMultiplyStatement(c *MultiplyStatementContext)

	// ExitMultiplyRegular is called when exiting the multiplyRegular production.
	ExitMultiplyRegular(c *MultiplyRegularContext)

	// ExitMultiplyRegularOperand is called when exiting the multiplyRegularOperand production.
	ExitMultiplyRegularOperand(c *MultiplyRegularOperandContext)

	// ExitMultiplyGiving is called when exiting the multiplyGiving production.
	ExitMultiplyGiving(c *MultiplyGivingContext)

	// ExitMultiplyGivingOperand is called when exiting the multiplyGivingOperand production.
	ExitMultiplyGivingOperand(c *MultiplyGivingOperandContext)

	// ExitMultiplyGivingResult is called when exiting the multiplyGivingResult production.
	ExitMultiplyGivingResult(c *MultiplyGivingResultContext)

	// ExitOpenStatement is called when exiting the openStatement production.
	ExitOpenStatement(c *OpenStatementContext)

	// ExitOpenInputStatement is called when exiting the openInputStatement production.
	ExitOpenInputStatement(c *OpenInputStatementContext)

	// ExitOpenInput is called when exiting the openInput production.
	ExitOpenInput(c *OpenInputContext)

	// ExitOpenOutputStatement is called when exiting the openOutputStatement production.
	ExitOpenOutputStatement(c *OpenOutputStatementContext)

	// ExitOpenOutput is called when exiting the openOutput production.
	ExitOpenOutput(c *OpenOutputContext)

	// ExitOpenIOStatement is called when exiting the openIOStatement production.
	ExitOpenIOStatement(c *OpenIOStatementContext)

	// ExitOpenExtendStatement is called when exiting the openExtendStatement production.
	ExitOpenExtendStatement(c *OpenExtendStatementContext)

	// ExitPerformStatement is called when exiting the performStatement production.
	ExitPerformStatement(c *PerformStatementContext)

	// ExitPerformInlineStatement is called when exiting the performInlineStatement production.
	ExitPerformInlineStatement(c *PerformInlineStatementContext)

	// ExitPerformProcedureStatement is called when exiting the performProcedureStatement production.
	ExitPerformProcedureStatement(c *PerformProcedureStatementContext)

	// ExitPerformType is called when exiting the performType production.
	ExitPerformType(c *PerformTypeContext)

	// ExitPerformTimes is called when exiting the performTimes production.
	ExitPerformTimes(c *PerformTimesContext)

	// ExitPerformUntil is called when exiting the performUntil production.
	ExitPerformUntil(c *PerformUntilContext)

	// ExitPerformVarying is called when exiting the performVarying production.
	ExitPerformVarying(c *PerformVaryingContext)

	// ExitPerformVaryingClause is called when exiting the performVaryingClause production.
	ExitPerformVaryingClause(c *PerformVaryingClauseContext)

	// ExitPerformVaryingPhrase is called when exiting the performVaryingPhrase production.
	ExitPerformVaryingPhrase(c *PerformVaryingPhraseContext)

	// ExitPerformAfter is called when exiting the performAfter production.
	ExitPerformAfter(c *PerformAfterContext)

	// ExitPerformFrom is called when exiting the performFrom production.
	ExitPerformFrom(c *PerformFromContext)

	// ExitPerformBy is called when exiting the performBy production.
	ExitPerformBy(c *PerformByContext)

	// ExitPerformTestClause is called when exiting the performTestClause production.
	ExitPerformTestClause(c *PerformTestClauseContext)

	// ExitPurgeStatement is called when exiting the purgeStatement production.
	ExitPurgeStatement(c *PurgeStatementContext)

	// ExitReadStatement is called when exiting the readStatement production.
	ExitReadStatement(c *ReadStatementContext)

	// ExitReadInto is called when exiting the readInto production.
	ExitReadInto(c *ReadIntoContext)

	// ExitReadWith is called when exiting the readWith production.
	ExitReadWith(c *ReadWithContext)

	// ExitReadKey is called when exiting the readKey production.
	ExitReadKey(c *ReadKeyContext)

	// ExitReceiveStatement is called when exiting the receiveStatement production.
	ExitReceiveStatement(c *ReceiveStatementContext)

	// ExitReceiveFromStatement is called when exiting the receiveFromStatement production.
	ExitReceiveFromStatement(c *ReceiveFromStatementContext)

	// ExitReceiveFrom is called when exiting the receiveFrom production.
	ExitReceiveFrom(c *ReceiveFromContext)

	// ExitReceiveIntoStatement is called when exiting the receiveIntoStatement production.
	ExitReceiveIntoStatement(c *ReceiveIntoStatementContext)

	// ExitReceiveNoData is called when exiting the receiveNoData production.
	ExitReceiveNoData(c *ReceiveNoDataContext)

	// ExitReceiveWithData is called when exiting the receiveWithData production.
	ExitReceiveWithData(c *ReceiveWithDataContext)

	// ExitReceiveBefore is called when exiting the receiveBefore production.
	ExitReceiveBefore(c *ReceiveBeforeContext)

	// ExitReceiveWith is called when exiting the receiveWith production.
	ExitReceiveWith(c *ReceiveWithContext)

	// ExitReceiveThread is called when exiting the receiveThread production.
	ExitReceiveThread(c *ReceiveThreadContext)

	// ExitReceiveSize is called when exiting the receiveSize production.
	ExitReceiveSize(c *ReceiveSizeContext)

	// ExitReceiveStatus is called when exiting the receiveStatus production.
	ExitReceiveStatus(c *ReceiveStatusContext)

	// ExitReleaseStatement is called when exiting the releaseStatement production.
	ExitReleaseStatement(c *ReleaseStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitReturnInto is called when exiting the returnInto production.
	ExitReturnInto(c *ReturnIntoContext)

	// ExitRewriteStatement is called when exiting the rewriteStatement production.
	ExitRewriteStatement(c *RewriteStatementContext)

	// ExitRewriteFrom is called when exiting the rewriteFrom production.
	ExitRewriteFrom(c *RewriteFromContext)

	// ExitSearchStatement is called when exiting the searchStatement production.
	ExitSearchStatement(c *SearchStatementContext)

	// ExitSearchVarying is called when exiting the searchVarying production.
	ExitSearchVarying(c *SearchVaryingContext)

	// ExitSearchWhen is called when exiting the searchWhen production.
	ExitSearchWhen(c *SearchWhenContext)

	// ExitSendStatement is called when exiting the sendStatement production.
	ExitSendStatement(c *SendStatementContext)

	// ExitSendStatementSync is called when exiting the sendStatementSync production.
	ExitSendStatementSync(c *SendStatementSyncContext)

	// ExitSendStatementAsync is called when exiting the sendStatementAsync production.
	ExitSendStatementAsync(c *SendStatementAsyncContext)

	// ExitSendFromPhrase is called when exiting the sendFromPhrase production.
	ExitSendFromPhrase(c *SendFromPhraseContext)

	// ExitSendWithPhrase is called when exiting the sendWithPhrase production.
	ExitSendWithPhrase(c *SendWithPhraseContext)

	// ExitSendReplacingPhrase is called when exiting the sendReplacingPhrase production.
	ExitSendReplacingPhrase(c *SendReplacingPhraseContext)

	// ExitSendAdvancingPhrase is called when exiting the sendAdvancingPhrase production.
	ExitSendAdvancingPhrase(c *SendAdvancingPhraseContext)

	// ExitSendAdvancingPage is called when exiting the sendAdvancingPage production.
	ExitSendAdvancingPage(c *SendAdvancingPageContext)

	// ExitSendAdvancingLines is called when exiting the sendAdvancingLines production.
	ExitSendAdvancingLines(c *SendAdvancingLinesContext)

	// ExitSendAdvancingMnemonic is called when exiting the sendAdvancingMnemonic production.
	ExitSendAdvancingMnemonic(c *SendAdvancingMnemonicContext)

	// ExitSetStatement is called when exiting the setStatement production.
	ExitSetStatement(c *SetStatementContext)

	// ExitSetToStatement is called when exiting the setToStatement production.
	ExitSetToStatement(c *SetToStatementContext)

	// ExitSetUpDownByStatement is called when exiting the setUpDownByStatement production.
	ExitSetUpDownByStatement(c *SetUpDownByStatementContext)

	// ExitSetTo is called when exiting the setTo production.
	ExitSetTo(c *SetToContext)

	// ExitSetToValue is called when exiting the setToValue production.
	ExitSetToValue(c *SetToValueContext)

	// ExitSetByValue is called when exiting the setByValue production.
	ExitSetByValue(c *SetByValueContext)

	// ExitSortStatement is called when exiting the sortStatement production.
	ExitSortStatement(c *SortStatementContext)

	// ExitSortOnKeyClause is called when exiting the sortOnKeyClause production.
	ExitSortOnKeyClause(c *SortOnKeyClauseContext)

	// ExitSortDuplicatesPhrase is called when exiting the sortDuplicatesPhrase production.
	ExitSortDuplicatesPhrase(c *SortDuplicatesPhraseContext)

	// ExitSortCollatingSequencePhrase is called when exiting the sortCollatingSequencePhrase production.
	ExitSortCollatingSequencePhrase(c *SortCollatingSequencePhraseContext)

	// ExitSortCollatingAlphanumeric is called when exiting the sortCollatingAlphanumeric production.
	ExitSortCollatingAlphanumeric(c *SortCollatingAlphanumericContext)

	// ExitSortCollatingNational is called when exiting the sortCollatingNational production.
	ExitSortCollatingNational(c *SortCollatingNationalContext)

	// ExitSortInputProcedurePhrase is called when exiting the sortInputProcedurePhrase production.
	ExitSortInputProcedurePhrase(c *SortInputProcedurePhraseContext)

	// ExitSortInputThrough is called when exiting the sortInputThrough production.
	ExitSortInputThrough(c *SortInputThroughContext)

	// ExitSortUsing is called when exiting the sortUsing production.
	ExitSortUsing(c *SortUsingContext)

	// ExitSortOutputProcedurePhrase is called when exiting the sortOutputProcedurePhrase production.
	ExitSortOutputProcedurePhrase(c *SortOutputProcedurePhraseContext)

	// ExitSortOutputThrough is called when exiting the sortOutputThrough production.
	ExitSortOutputThrough(c *SortOutputThroughContext)

	// ExitSortGivingPhrase is called when exiting the sortGivingPhrase production.
	ExitSortGivingPhrase(c *SortGivingPhraseContext)

	// ExitSortGiving is called when exiting the sortGiving production.
	ExitSortGiving(c *SortGivingContext)

	// ExitStartStatement is called when exiting the startStatement production.
	ExitStartStatement(c *StartStatementContext)

	// ExitStartKey is called when exiting the startKey production.
	ExitStartKey(c *StartKeyContext)

	// ExitStopStatement is called when exiting the stopStatement production.
	ExitStopStatement(c *StopStatementContext)

	// ExitStringStatement is called when exiting the stringStatement production.
	ExitStringStatement(c *StringStatementContext)

	// ExitStringSendingPhrase is called when exiting the stringSendingPhrase production.
	ExitStringSendingPhrase(c *StringSendingPhraseContext)

	// ExitStringSending is called when exiting the stringSending production.
	ExitStringSending(c *StringSendingContext)

	// ExitStringDelimitedByPhrase is called when exiting the stringDelimitedByPhrase production.
	ExitStringDelimitedByPhrase(c *StringDelimitedByPhraseContext)

	// ExitStringForPhrase is called when exiting the stringForPhrase production.
	ExitStringForPhrase(c *StringForPhraseContext)

	// ExitStringIntoPhrase is called when exiting the stringIntoPhrase production.
	ExitStringIntoPhrase(c *StringIntoPhraseContext)

	// ExitStringWithPointerPhrase is called when exiting the stringWithPointerPhrase production.
	ExitStringWithPointerPhrase(c *StringWithPointerPhraseContext)

	// ExitSubtractStatement is called when exiting the subtractStatement production.
	ExitSubtractStatement(c *SubtractStatementContext)

	// ExitSubtractFromStatement is called when exiting the subtractFromStatement production.
	ExitSubtractFromStatement(c *SubtractFromStatementContext)

	// ExitSubtractFromGivingStatement is called when exiting the subtractFromGivingStatement production.
	ExitSubtractFromGivingStatement(c *SubtractFromGivingStatementContext)

	// ExitSubtractCorrespondingStatement is called when exiting the subtractCorrespondingStatement production.
	ExitSubtractCorrespondingStatement(c *SubtractCorrespondingStatementContext)

	// ExitSubtractSubtrahend is called when exiting the subtractSubtrahend production.
	ExitSubtractSubtrahend(c *SubtractSubtrahendContext)

	// ExitSubtractMinuend is called when exiting the subtractMinuend production.
	ExitSubtractMinuend(c *SubtractMinuendContext)

	// ExitSubtractMinuendGiving is called when exiting the subtractMinuendGiving production.
	ExitSubtractMinuendGiving(c *SubtractMinuendGivingContext)

	// ExitSubtractGiving is called when exiting the subtractGiving production.
	ExitSubtractGiving(c *SubtractGivingContext)

	// ExitSubtractMinuendCorresponding is called when exiting the subtractMinuendCorresponding production.
	ExitSubtractMinuendCorresponding(c *SubtractMinuendCorrespondingContext)

	// ExitTerminateStatement is called when exiting the terminateStatement production.
	ExitTerminateStatement(c *TerminateStatementContext)

	// ExitUnstringStatement is called when exiting the unstringStatement production.
	ExitUnstringStatement(c *UnstringStatementContext)

	// ExitUnstringSendingPhrase is called when exiting the unstringSendingPhrase production.
	ExitUnstringSendingPhrase(c *UnstringSendingPhraseContext)

	// ExitUnstringDelimitedByPhrase is called when exiting the unstringDelimitedByPhrase production.
	ExitUnstringDelimitedByPhrase(c *UnstringDelimitedByPhraseContext)

	// ExitUnstringOrAllPhrase is called when exiting the unstringOrAllPhrase production.
	ExitUnstringOrAllPhrase(c *UnstringOrAllPhraseContext)

	// ExitUnstringIntoPhrase is called when exiting the unstringIntoPhrase production.
	ExitUnstringIntoPhrase(c *UnstringIntoPhraseContext)

	// ExitUnstringInto is called when exiting the unstringInto production.
	ExitUnstringInto(c *UnstringIntoContext)

	// ExitUnstringDelimiterIn is called when exiting the unstringDelimiterIn production.
	ExitUnstringDelimiterIn(c *UnstringDelimiterInContext)

	// ExitUnstringCountIn is called when exiting the unstringCountIn production.
	ExitUnstringCountIn(c *UnstringCountInContext)

	// ExitUnstringWithPointerPhrase is called when exiting the unstringWithPointerPhrase production.
	ExitUnstringWithPointerPhrase(c *UnstringWithPointerPhraseContext)

	// ExitUnstringTallyingPhrase is called when exiting the unstringTallyingPhrase production.
	ExitUnstringTallyingPhrase(c *UnstringTallyingPhraseContext)

	// ExitUseStatement is called when exiting the useStatement production.
	ExitUseStatement(c *UseStatementContext)

	// ExitUseAfterClause is called when exiting the useAfterClause production.
	ExitUseAfterClause(c *UseAfterClauseContext)

	// ExitUseAfterOn is called when exiting the useAfterOn production.
	ExitUseAfterOn(c *UseAfterOnContext)

	// ExitUseDebugClause is called when exiting the useDebugClause production.
	ExitUseDebugClause(c *UseDebugClauseContext)

	// ExitUseDebugOn is called when exiting the useDebugOn production.
	ExitUseDebugOn(c *UseDebugOnContext)

	// ExitWriteStatement is called when exiting the writeStatement production.
	ExitWriteStatement(c *WriteStatementContext)

	// ExitWriteFromPhrase is called when exiting the writeFromPhrase production.
	ExitWriteFromPhrase(c *WriteFromPhraseContext)

	// ExitWriteAdvancingPhrase is called when exiting the writeAdvancingPhrase production.
	ExitWriteAdvancingPhrase(c *WriteAdvancingPhraseContext)

	// ExitWriteAdvancingPage is called when exiting the writeAdvancingPage production.
	ExitWriteAdvancingPage(c *WriteAdvancingPageContext)

	// ExitWriteAdvancingLines is called when exiting the writeAdvancingLines production.
	ExitWriteAdvancingLines(c *WriteAdvancingLinesContext)

	// ExitWriteAdvancingMnemonic is called when exiting the writeAdvancingMnemonic production.
	ExitWriteAdvancingMnemonic(c *WriteAdvancingMnemonicContext)

	// ExitWriteAtEndOfPagePhrase is called when exiting the writeAtEndOfPagePhrase production.
	ExitWriteAtEndOfPagePhrase(c *WriteAtEndOfPagePhraseContext)

	// ExitWriteNotAtEndOfPagePhrase is called when exiting the writeNotAtEndOfPagePhrase production.
	ExitWriteNotAtEndOfPagePhrase(c *WriteNotAtEndOfPagePhraseContext)

	// ExitAtEndPhrase is called when exiting the atEndPhrase production.
	ExitAtEndPhrase(c *AtEndPhraseContext)

	// ExitNotAtEndPhrase is called when exiting the notAtEndPhrase production.
	ExitNotAtEndPhrase(c *NotAtEndPhraseContext)

	// ExitInvalidKeyPhrase is called when exiting the invalidKeyPhrase production.
	ExitInvalidKeyPhrase(c *InvalidKeyPhraseContext)

	// ExitNotInvalidKeyPhrase is called when exiting the notInvalidKeyPhrase production.
	ExitNotInvalidKeyPhrase(c *NotInvalidKeyPhraseContext)

	// ExitOnOverflowPhrase is called when exiting the onOverflowPhrase production.
	ExitOnOverflowPhrase(c *OnOverflowPhraseContext)

	// ExitNotOnOverflowPhrase is called when exiting the notOnOverflowPhrase production.
	ExitNotOnOverflowPhrase(c *NotOnOverflowPhraseContext)

	// ExitOnSizeErrorPhrase is called when exiting the onSizeErrorPhrase production.
	ExitOnSizeErrorPhrase(c *OnSizeErrorPhraseContext)

	// ExitNotOnSizeErrorPhrase is called when exiting the notOnSizeErrorPhrase production.
	ExitNotOnSizeErrorPhrase(c *NotOnSizeErrorPhraseContext)

	// ExitOnExceptionClause is called when exiting the onExceptionClause production.
	ExitOnExceptionClause(c *OnExceptionClauseContext)

	// ExitNotOnExceptionClause is called when exiting the notOnExceptionClause production.
	ExitNotOnExceptionClause(c *NotOnExceptionClauseContext)

	// ExitArithmeticExpression is called when exiting the arithmeticExpression production.
	ExitArithmeticExpression(c *ArithmeticExpressionContext)

	// ExitPlusMinus is called when exiting the plusMinus production.
	ExitPlusMinus(c *PlusMinusContext)

	// ExitMultDivs is called when exiting the multDivs production.
	ExitMultDivs(c *MultDivsContext)

	// ExitMultDiv is called when exiting the multDiv production.
	ExitMultDiv(c *MultDivContext)

	// ExitPowers is called when exiting the powers production.
	ExitPowers(c *PowersContext)

	// ExitPower is called when exiting the power production.
	ExitPower(c *PowerContext)

	// ExitBasis is called when exiting the basis production.
	ExitBasis(c *BasisContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitAndOrCondition is called when exiting the andOrCondition production.
	ExitAndOrCondition(c *AndOrConditionContext)

	// ExitCombinableCondition is called when exiting the combinableCondition production.
	ExitCombinableCondition(c *CombinableConditionContext)

	// ExitSimpleCondition is called when exiting the simpleCondition production.
	ExitSimpleCondition(c *SimpleConditionContext)

	// ExitClassCondition is called when exiting the classCondition production.
	ExitClassCondition(c *ClassConditionContext)

	// ExitConditionNameReference is called when exiting the conditionNameReference production.
	ExitConditionNameReference(c *ConditionNameReferenceContext)

	// ExitConditionNameSubscriptReference is called when exiting the conditionNameSubscriptReference production.
	ExitConditionNameSubscriptReference(c *ConditionNameSubscriptReferenceContext)

	// ExitRelationCondition is called when exiting the relationCondition production.
	ExitRelationCondition(c *RelationConditionContext)

	// ExitRelationSignCondition is called when exiting the relationSignCondition production.
	ExitRelationSignCondition(c *RelationSignConditionContext)

	// ExitRelationArithmeticComparison is called when exiting the relationArithmeticComparison production.
	ExitRelationArithmeticComparison(c *RelationArithmeticComparisonContext)

	// ExitRelationCombinedComparison is called when exiting the relationCombinedComparison production.
	ExitRelationCombinedComparison(c *RelationCombinedComparisonContext)

	// ExitRelationCombinedCondition is called when exiting the relationCombinedCondition production.
	ExitRelationCombinedCondition(c *RelationCombinedConditionContext)

	// ExitRelationalOperator is called when exiting the relationalOperator production.
	ExitRelationalOperator(c *RelationalOperatorContext)

	// ExitAbbreviation is called when exiting the abbreviation production.
	ExitAbbreviation(c *AbbreviationContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitTableCall is called when exiting the tableCall production.
	ExitTableCall(c *TableCallContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitReferenceModifier is called when exiting the referenceModifier production.
	ExitReferenceModifier(c *ReferenceModifierContext)

	// ExitCharacterPosition is called when exiting the characterPosition production.
	ExitCharacterPosition(c *CharacterPositionContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)

	// ExitSubscript_ is called when exiting the subscript_ production.
	ExitSubscript_(c *Subscript_Context)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitQualifiedDataName is called when exiting the qualifiedDataName production.
	ExitQualifiedDataName(c *QualifiedDataNameContext)

	// ExitQualifiedDataNameFormat1 is called when exiting the qualifiedDataNameFormat1 production.
	ExitQualifiedDataNameFormat1(c *QualifiedDataNameFormat1Context)

	// ExitQualifiedDataNameFormat2 is called when exiting the qualifiedDataNameFormat2 production.
	ExitQualifiedDataNameFormat2(c *QualifiedDataNameFormat2Context)

	// ExitQualifiedDataNameFormat3 is called when exiting the qualifiedDataNameFormat3 production.
	ExitQualifiedDataNameFormat3(c *QualifiedDataNameFormat3Context)

	// ExitQualifiedDataNameFormat4 is called when exiting the qualifiedDataNameFormat4 production.
	ExitQualifiedDataNameFormat4(c *QualifiedDataNameFormat4Context)

	// ExitQualifiedInData is called when exiting the qualifiedInData production.
	ExitQualifiedInData(c *QualifiedInDataContext)

	// ExitInData is called when exiting the inData production.
	ExitInData(c *InDataContext)

	// ExitInFile is called when exiting the inFile production.
	ExitInFile(c *InFileContext)

	// ExitInMnemonic is called when exiting the inMnemonic production.
	ExitInMnemonic(c *InMnemonicContext)

	// ExitInSection is called when exiting the inSection production.
	ExitInSection(c *InSectionContext)

	// ExitInLibrary is called when exiting the inLibrary production.
	ExitInLibrary(c *InLibraryContext)

	// ExitInTable is called when exiting the inTable production.
	ExitInTable(c *InTableContext)

	// ExitAlphabetName is called when exiting the alphabetName production.
	ExitAlphabetName(c *AlphabetNameContext)

	// ExitAssignmentName is called when exiting the assignmentName production.
	ExitAssignmentName(c *AssignmentNameContext)

	// ExitBasisName is called when exiting the basisName production.
	ExitBasisName(c *BasisNameContext)

	// ExitCdName is called when exiting the cdName production.
	ExitCdName(c *CdNameContext)

	// ExitClassName is called when exiting the className production.
	ExitClassName(c *ClassNameContext)

	// ExitComputerName is called when exiting the computerName production.
	ExitComputerName(c *ComputerNameContext)

	// ExitConditionName is called when exiting the conditionName production.
	ExitConditionName(c *ConditionNameContext)

	// ExitDataName is called when exiting the dataName production.
	ExitDataName(c *DataNameContext)

	// ExitDataDescName is called when exiting the dataDescName production.
	ExitDataDescName(c *DataDescNameContext)

	// ExitEnvironmentName is called when exiting the environmentName production.
	ExitEnvironmentName(c *EnvironmentNameContext)

	// ExitFileName is called when exiting the fileName production.
	ExitFileName(c *FileNameContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitLanguageName is called when exiting the languageName production.
	ExitLanguageName(c *LanguageNameContext)

	// ExitLibraryName is called when exiting the libraryName production.
	ExitLibraryName(c *LibraryNameContext)

	// ExitLocalName is called when exiting the localName production.
	ExitLocalName(c *LocalNameContext)

	// ExitMnemonicName is called when exiting the mnemonicName production.
	ExitMnemonicName(c *MnemonicNameContext)

	// ExitParagraphName is called when exiting the paragraphName production.
	ExitParagraphName(c *ParagraphNameContext)

	// ExitProcedureName is called when exiting the procedureName production.
	ExitProcedureName(c *ProcedureNameContext)

	// ExitProgramName is called when exiting the programName production.
	ExitProgramName(c *ProgramNameContext)

	// ExitRecordName is called when exiting the recordName production.
	ExitRecordName(c *RecordNameContext)

	// ExitReportName is called when exiting the reportName production.
	ExitReportName(c *ReportNameContext)

	// ExitRoutineName is called when exiting the routineName production.
	ExitRoutineName(c *RoutineNameContext)

	// ExitScreenName is called when exiting the screenName production.
	ExitScreenName(c *ScreenNameContext)

	// ExitSectionName is called when exiting the sectionName production.
	ExitSectionName(c *SectionNameContext)

	// ExitSystemName is called when exiting the systemName production.
	ExitSystemName(c *SystemNameContext)

	// ExitSymbolicCharacter is called when exiting the symbolicCharacter production.
	ExitSymbolicCharacter(c *SymbolicCharacterContext)

	// ExitTextName is called when exiting the textName production.
	ExitTextName(c *TextNameContext)

	// ExitCobolWord is called when exiting the cobolWord production.
	ExitCobolWord(c *CobolWordContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitCicsDfhRespLiteral is called when exiting the cicsDfhRespLiteral production.
	ExitCicsDfhRespLiteral(c *CicsDfhRespLiteralContext)

	// ExitCicsDfhValueLiteral is called when exiting the cicsDfhValueLiteral production.
	ExitCicsDfhValueLiteral(c *CicsDfhValueLiteralContext)

	// ExitFigurativeConstant is called when exiting the figurativeConstant production.
	ExitFigurativeConstant(c *FigurativeConstantContext)

	// ExitSpecialRegister is called when exiting the specialRegister production.
	ExitSpecialRegister(c *SpecialRegisterContext)

	// ExitCommentEntry is called when exiting the commentEntry production.
	ExitCommentEntry(c *CommentEntryContext)
}
