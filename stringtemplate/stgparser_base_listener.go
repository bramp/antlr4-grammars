// Generated from STGParser.g4 by ANTLR 4.7.

package stringtemplate // STGParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSTGParserListener is a complete listener for a parse tree produced by STGParser.
type BaseSTGParserListener struct{}

var _ STGParserListener = &BaseSTGParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSTGParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSTGParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSTGParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSTGParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseSTGParserListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseSTGParserListener) ExitGroup(ctx *GroupContext) {}

// EnterDelimiters is called when production delimiters is entered.
func (s *BaseSTGParserListener) EnterDelimiters(ctx *DelimitersContext) {}

// ExitDelimiters is called when production delimiters is exited.
func (s *BaseSTGParserListener) ExitDelimiters(ctx *DelimitersContext) {}

// EnterImports is called when production imports is entered.
func (s *BaseSTGParserListener) EnterImports(ctx *ImportsContext) {}

// ExitImports is called when production imports is exited.
func (s *BaseSTGParserListener) ExitImports(ctx *ImportsContext) {}

// EnterTemplate is called when production template is entered.
func (s *BaseSTGParserListener) EnterTemplate(ctx *TemplateContext) {}

// ExitTemplate is called when production template is exited.
func (s *BaseSTGParserListener) ExitTemplate(ctx *TemplateContext) {}

// EnterFormalArgs is called when production formalArgs is entered.
func (s *BaseSTGParserListener) EnterFormalArgs(ctx *FormalArgsContext) {}

// ExitFormalArgs is called when production formalArgs is exited.
func (s *BaseSTGParserListener) ExitFormalArgs(ctx *FormalArgsContext) {}

// EnterFormalArg is called when production formalArg is entered.
func (s *BaseSTGParserListener) EnterFormalArg(ctx *FormalArgContext) {}

// ExitFormalArg is called when production formalArg is exited.
func (s *BaseSTGParserListener) ExitFormalArg(ctx *FormalArgContext) {}

// EnterDict is called when production dict is entered.
func (s *BaseSTGParserListener) EnterDict(ctx *DictContext) {}

// ExitDict is called when production dict is exited.
func (s *BaseSTGParserListener) ExitDict(ctx *DictContext) {}

// EnterDictPairs is called when production dictPairs is entered.
func (s *BaseSTGParserListener) EnterDictPairs(ctx *DictPairsContext) {}

// ExitDictPairs is called when production dictPairs is exited.
func (s *BaseSTGParserListener) ExitDictPairs(ctx *DictPairsContext) {}

// EnterKeyValuePair is called when production keyValuePair is entered.
func (s *BaseSTGParserListener) EnterKeyValuePair(ctx *KeyValuePairContext) {}

// ExitKeyValuePair is called when production keyValuePair is exited.
func (s *BaseSTGParserListener) ExitKeyValuePair(ctx *KeyValuePairContext) {}

// EnterDefaultValuePair is called when production defaultValuePair is entered.
func (s *BaseSTGParserListener) EnterDefaultValuePair(ctx *DefaultValuePairContext) {}

// ExitDefaultValuePair is called when production defaultValuePair is exited.
func (s *BaseSTGParserListener) ExitDefaultValuePair(ctx *DefaultValuePairContext) {}

// EnterKeyValue is called when production keyValue is entered.
func (s *BaseSTGParserListener) EnterKeyValue(ctx *KeyValueContext) {}

// ExitKeyValue is called when production keyValue is exited.
func (s *BaseSTGParserListener) ExitKeyValue(ctx *KeyValueContext) {}
