// Code generated from apt.g4 by ANTLR 4.9.3. DO NOT EDIT.

package apt // apt
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseaptListener is a complete listener for a parse tree produced by aptParser.
type BaseaptListener struct{}

var _ aptListener = &BaseaptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseaptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseaptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseaptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseaptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRecord is called when production record is entered.
func (s *BaseaptListener) EnterRecord(ctx *RecordContext) {}

// ExitRecord is called when production record is exited.
func (s *BaseaptListener) ExitRecord(ctx *RecordContext) {}

// EnterWordWithDashSegment is called when production wordWithDashSegment is entered.
func (s *BaseaptListener) EnterWordWithDashSegment(ctx *WordWithDashSegmentContext) {}

// ExitWordWithDashSegment is called when production wordWithDashSegment is exited.
func (s *BaseaptListener) ExitWordWithDashSegment(ctx *WordWithDashSegmentContext) {}

// EnterWordWithDash is called when production wordWithDash is entered.
func (s *BaseaptListener) EnterWordWithDash(ctx *WordWithDashContext) {}

// ExitWordWithDash is called when production wordWithDash is exited.
func (s *BaseaptListener) ExitWordWithDash(ctx *WordWithDashContext) {}

// EnterComponent is called when production component is entered.
func (s *BaseaptListener) EnterComponent(ctx *ComponentContext) {}

// ExitComponent is called when production component is exited.
func (s *BaseaptListener) ExitComponent(ctx *ComponentContext) {}

// EnterComponentsR is called when production componentsR is entered.
func (s *BaseaptListener) EnterComponentsR(ctx *ComponentsRContext) {}

// ExitComponentsR is called when production componentsR is exited.
func (s *BaseaptListener) ExitComponentsR(ctx *ComponentsRContext) {}

// EnterOptionsR is called when production optionsR is entered.
func (s *BaseaptListener) EnterOptionsR(ctx *OptionsRContext) {}

// ExitOptionsR is called when production optionsR is exited.
func (s *BaseaptListener) ExitOptionsR(ctx *OptionsRContext) {}

// EnterOptionsList is called when production optionsList is entered.
func (s *BaseaptListener) EnterOptionsList(ctx *OptionsListContext) {}

// ExitOptionsList is called when production optionsList is exited.
func (s *BaseaptListener) ExitOptionsList(ctx *OptionsListContext) {}

// EnterAdditionalOptions is called when production additionalOptions is entered.
func (s *BaseaptListener) EnterAdditionalOptions(ctx *AdditionalOptionsContext) {}

// ExitAdditionalOptions is called when production additionalOptions is exited.
func (s *BaseaptListener) ExitAdditionalOptions(ctx *AdditionalOptionsContext) {}

// EnterAdditionalOption is called when production additionalOption is entered.
func (s *BaseaptListener) EnterAdditionalOption(ctx *AdditionalOptionContext) {}

// ExitAdditionalOption is called when production additionalOption is exited.
func (s *BaseaptListener) ExitAdditionalOption(ctx *AdditionalOptionContext) {}

// EnterOptionR is called when production optionR is entered.
func (s *BaseaptListener) EnterOptionR(ctx *OptionRContext) {}

// ExitOptionR is called when production optionR is exited.
func (s *BaseaptListener) ExitOptionR(ctx *OptionRContext) {}

// EnterWordWithPlus is called when production wordWithPlus is entered.
func (s *BaseaptListener) EnterWordWithPlus(ctx *WordWithPlusContext) {}

// ExitWordWithPlus is called when production wordWithPlus is exited.
func (s *BaseaptListener) ExitWordWithPlus(ctx *WordWithPlusContext) {}

// EnterUriSchema is called when production uriSchema is entered.
func (s *BaseaptListener) EnterUriSchema(ctx *UriSchemaContext) {}

// ExitUriSchema is called when production uriSchema is exited.
func (s *BaseaptListener) ExitUriSchema(ctx *UriSchemaContext) {}

// EnterRestSchemaWords is called when production restSchemaWords is entered.
func (s *BaseaptListener) EnterRestSchemaWords(ctx *RestSchemaWordsContext) {}

// ExitRestSchemaWords is called when production restSchemaWords is exited.
func (s *BaseaptListener) ExitRestSchemaWords(ctx *RestSchemaWordsContext) {}

// EnterGenericURI is called when production genericURI is entered.
func (s *BaseaptListener) EnterGenericURI(ctx *GenericURIContext) {}

// ExitGenericURI is called when production genericURI is exited.
func (s *BaseaptListener) ExitGenericURI(ctx *GenericURIContext) {}

// EnterUriR is called when production uriR is entered.
func (s *BaseaptListener) EnterUriR(ctx *UriRContext) {}

// ExitUriR is called when production uriR is exited.
func (s *BaseaptListener) ExitUriR(ctx *UriRContext) {}

// EnterCommenterR is called when production commenterR is entered.
func (s *BaseaptListener) EnterCommenterR(ctx *CommenterRContext) {}

// ExitCommenterR is called when production commenterR is exited.
func (s *BaseaptListener) ExitCommenterR(ctx *CommenterRContext) {}

// EnterOptionValueSegment is called when production optionValueSegment is entered.
func (s *BaseaptListener) EnterOptionValueSegment(ctx *OptionValueSegmentContext) {}

// ExitOptionValueSegment is called when production optionValueSegment is exited.
func (s *BaseaptListener) ExitOptionValueSegment(ctx *OptionValueSegmentContext) {}

// EnterOptionValue is called when production optionValue is entered.
func (s *BaseaptListener) EnterOptionValue(ctx *OptionValueContext) {}

// ExitOptionValue is called when production optionValue is exited.
func (s *BaseaptListener) ExitOptionValue(ctx *OptionValueContext) {}

// EnterNonSquareBracketStringSegment is called when production nonSquareBracketStringSegment is entered.
func (s *BaseaptListener) EnterNonSquareBracketStringSegment(ctx *NonSquareBracketStringSegmentContext) {
}

// ExitNonSquareBracketStringSegment is called when production nonSquareBracketStringSegment is exited.
func (s *BaseaptListener) ExitNonSquareBracketStringSegment(ctx *NonSquareBracketStringSegmentContext) {
}

// EnterNonSquareBracketString is called when production nonSquareBracketString is entered.
func (s *BaseaptListener) EnterNonSquareBracketString(ctx *NonSquareBracketStringContext) {}

// ExitNonSquareBracketString is called when production nonSquareBracketString is exited.
func (s *BaseaptListener) ExitNonSquareBracketString(ctx *NonSquareBracketStringContext) {}

// EnterNonSpaceStringSegment is called when production nonSpaceStringSegment is entered.
func (s *BaseaptListener) EnterNonSpaceStringSegment(ctx *NonSpaceStringSegmentContext) {}

// ExitNonSpaceStringSegment is called when production nonSpaceStringSegment is exited.
func (s *BaseaptListener) ExitNonSpaceStringSegment(ctx *NonSpaceStringSegmentContext) {}

// EnterNonSpaceString is called when production nonSpaceString is entered.
func (s *BaseaptListener) EnterNonSpaceString(ctx *NonSpaceStringContext) {}

// ExitNonSpaceString is called when production nonSpaceString is exited.
func (s *BaseaptListener) ExitNonSpaceString(ctx *NonSpaceStringContext) {}

// EnterSingleTickEnclosedString is called when production singleTickEnclosedString is entered.
func (s *BaseaptListener) EnterSingleTickEnclosedString(ctx *SingleTickEnclosedStringContext) {}

// ExitSingleTickEnclosedString is called when production singleTickEnclosedString is exited.
func (s *BaseaptListener) ExitSingleTickEnclosedString(ctx *SingleTickEnclosedStringContext) {}

// EnterDoubleTickEnclosedString is called when production doubleTickEnclosedString is entered.
func (s *BaseaptListener) EnterDoubleTickEnclosedString(ctx *DoubleTickEnclosedStringContext) {}

// ExitDoubleTickEnclosedString is called when production doubleTickEnclosedString is exited.
func (s *BaseaptListener) ExitDoubleTickEnclosedString(ctx *DoubleTickEnclosedStringContext) {}

// EnterTickEnclosedString is called when production tickEnclosedString is entered.
func (s *BaseaptListener) EnterTickEnclosedString(ctx *TickEnclosedStringContext) {}

// ExitTickEnclosedString is called when production tickEnclosedString is exited.
func (s *BaseaptListener) ExitTickEnclosedString(ctx *TickEnclosedStringContext) {}

// EnterEnclosedString is called when production enclosedString is entered.
func (s *BaseaptListener) EnterEnclosedString(ctx *EnclosedStringContext) {}

// ExitEnclosedString is called when production enclosedString is exited.
func (s *BaseaptListener) ExitEnclosedString(ctx *EnclosedStringContext) {}

// EnterCdromURI is called when production cdromURI is entered.
func (s *BaseaptListener) EnterCdromURI(ctx *CdromURIContext) {}

// ExitCdromURI is called when production cdromURI is exited.
func (s *BaseaptListener) ExitCdromURI(ctx *CdromURIContext) {}
