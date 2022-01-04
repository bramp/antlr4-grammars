// Code generated from apt.g4 by ANTLR 4.9.3. DO NOT EDIT.

package apt // apt
import "github.com/antlr/antlr4/runtime/Go/antlr"

// aptListener is a complete listener for a parse tree produced by aptParser.
type aptListener interface {
	antlr.ParseTreeListener

	// EnterRecord is called when entering the record production.
	EnterRecord(c *RecordContext)

	// EnterWordWithDashSegment is called when entering the wordWithDashSegment production.
	EnterWordWithDashSegment(c *WordWithDashSegmentContext)

	// EnterWordWithDash is called when entering the wordWithDash production.
	EnterWordWithDash(c *WordWithDashContext)

	// EnterComponent is called when entering the component production.
	EnterComponent(c *ComponentContext)

	// EnterComponentsR is called when entering the componentsR production.
	EnterComponentsR(c *ComponentsRContext)

	// EnterOptionsR is called when entering the optionsR production.
	EnterOptionsR(c *OptionsRContext)

	// EnterOptionsList is called when entering the optionsList production.
	EnterOptionsList(c *OptionsListContext)

	// EnterAdditionalOptions is called when entering the additionalOptions production.
	EnterAdditionalOptions(c *AdditionalOptionsContext)

	// EnterAdditionalOption is called when entering the additionalOption production.
	EnterAdditionalOption(c *AdditionalOptionContext)

	// EnterOptionR is called when entering the optionR production.
	EnterOptionR(c *OptionRContext)

	// EnterWordWithPlus is called when entering the wordWithPlus production.
	EnterWordWithPlus(c *WordWithPlusContext)

	// EnterUriSchema is called when entering the uriSchema production.
	EnterUriSchema(c *UriSchemaContext)

	// EnterRestSchemaWords is called when entering the restSchemaWords production.
	EnterRestSchemaWords(c *RestSchemaWordsContext)

	// EnterGenericURI is called when entering the genericURI production.
	EnterGenericURI(c *GenericURIContext)

	// EnterUriR is called when entering the uriR production.
	EnterUriR(c *UriRContext)

	// EnterCommenterR is called when entering the commenterR production.
	EnterCommenterR(c *CommenterRContext)

	// EnterOptionValueSegment is called when entering the optionValueSegment production.
	EnterOptionValueSegment(c *OptionValueSegmentContext)

	// EnterOptionValue is called when entering the optionValue production.
	EnterOptionValue(c *OptionValueContext)

	// EnterNonSquareBracketStringSegment is called when entering the nonSquareBracketStringSegment production.
	EnterNonSquareBracketStringSegment(c *NonSquareBracketStringSegmentContext)

	// EnterNonSquareBracketString is called when entering the nonSquareBracketString production.
	EnterNonSquareBracketString(c *NonSquareBracketStringContext)

	// EnterNonSpaceStringSegment is called when entering the nonSpaceStringSegment production.
	EnterNonSpaceStringSegment(c *NonSpaceStringSegmentContext)

	// EnterNonSpaceString is called when entering the nonSpaceString production.
	EnterNonSpaceString(c *NonSpaceStringContext)

	// EnterSingleTickEnclosedString is called when entering the singleTickEnclosedString production.
	EnterSingleTickEnclosedString(c *SingleTickEnclosedStringContext)

	// EnterDoubleTickEnclosedString is called when entering the doubleTickEnclosedString production.
	EnterDoubleTickEnclosedString(c *DoubleTickEnclosedStringContext)

	// EnterTickEnclosedString is called when entering the tickEnclosedString production.
	EnterTickEnclosedString(c *TickEnclosedStringContext)

	// EnterEnclosedString is called when entering the enclosedString production.
	EnterEnclosedString(c *EnclosedStringContext)

	// EnterCdromURI is called when entering the cdromURI production.
	EnterCdromURI(c *CdromURIContext)

	// ExitRecord is called when exiting the record production.
	ExitRecord(c *RecordContext)

	// ExitWordWithDashSegment is called when exiting the wordWithDashSegment production.
	ExitWordWithDashSegment(c *WordWithDashSegmentContext)

	// ExitWordWithDash is called when exiting the wordWithDash production.
	ExitWordWithDash(c *WordWithDashContext)

	// ExitComponent is called when exiting the component production.
	ExitComponent(c *ComponentContext)

	// ExitComponentsR is called when exiting the componentsR production.
	ExitComponentsR(c *ComponentsRContext)

	// ExitOptionsR is called when exiting the optionsR production.
	ExitOptionsR(c *OptionsRContext)

	// ExitOptionsList is called when exiting the optionsList production.
	ExitOptionsList(c *OptionsListContext)

	// ExitAdditionalOptions is called when exiting the additionalOptions production.
	ExitAdditionalOptions(c *AdditionalOptionsContext)

	// ExitAdditionalOption is called when exiting the additionalOption production.
	ExitAdditionalOption(c *AdditionalOptionContext)

	// ExitOptionR is called when exiting the optionR production.
	ExitOptionR(c *OptionRContext)

	// ExitWordWithPlus is called when exiting the wordWithPlus production.
	ExitWordWithPlus(c *WordWithPlusContext)

	// ExitUriSchema is called when exiting the uriSchema production.
	ExitUriSchema(c *UriSchemaContext)

	// ExitRestSchemaWords is called when exiting the restSchemaWords production.
	ExitRestSchemaWords(c *RestSchemaWordsContext)

	// ExitGenericURI is called when exiting the genericURI production.
	ExitGenericURI(c *GenericURIContext)

	// ExitUriR is called when exiting the uriR production.
	ExitUriR(c *UriRContext)

	// ExitCommenterR is called when exiting the commenterR production.
	ExitCommenterR(c *CommenterRContext)

	// ExitOptionValueSegment is called when exiting the optionValueSegment production.
	ExitOptionValueSegment(c *OptionValueSegmentContext)

	// ExitOptionValue is called when exiting the optionValue production.
	ExitOptionValue(c *OptionValueContext)

	// ExitNonSquareBracketStringSegment is called when exiting the nonSquareBracketStringSegment production.
	ExitNonSquareBracketStringSegment(c *NonSquareBracketStringSegmentContext)

	// ExitNonSquareBracketString is called when exiting the nonSquareBracketString production.
	ExitNonSquareBracketString(c *NonSquareBracketStringContext)

	// ExitNonSpaceStringSegment is called when exiting the nonSpaceStringSegment production.
	ExitNonSpaceStringSegment(c *NonSpaceStringSegmentContext)

	// ExitNonSpaceString is called when exiting the nonSpaceString production.
	ExitNonSpaceString(c *NonSpaceStringContext)

	// ExitSingleTickEnclosedString is called when exiting the singleTickEnclosedString production.
	ExitSingleTickEnclosedString(c *SingleTickEnclosedStringContext)

	// ExitDoubleTickEnclosedString is called when exiting the doubleTickEnclosedString production.
	ExitDoubleTickEnclosedString(c *DoubleTickEnclosedStringContext)

	// ExitTickEnclosedString is called when exiting the tickEnclosedString production.
	ExitTickEnclosedString(c *TickEnclosedStringContext)

	// ExitEnclosedString is called when exiting the enclosedString production.
	ExitEnclosedString(c *EnclosedStringContext)

	// ExitCdromURI is called when exiting the cdromURI production.
	ExitCdromURI(c *CdromURIContext)
}
