// Generated from STGParser.g4 by ANTLR 4.7.

package stringtemplate // STGParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// STGParserListener is a complete listener for a parse tree produced by STGParser.
type STGParserListener interface {
	antlr.ParseTreeListener

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterDelimiters is called when entering the delimiters production.
	EnterDelimiters(c *DelimitersContext)

	// EnterImports is called when entering the imports production.
	EnterImports(c *ImportsContext)

	// EnterTemplate is called when entering the template production.
	EnterTemplate(c *TemplateContext)

	// EnterFormalArgs is called when entering the formalArgs production.
	EnterFormalArgs(c *FormalArgsContext)

	// EnterFormalArg is called when entering the formalArg production.
	EnterFormalArg(c *FormalArgContext)

	// EnterDict is called when entering the dict production.
	EnterDict(c *DictContext)

	// EnterDictPairs is called when entering the dictPairs production.
	EnterDictPairs(c *DictPairsContext)

	// EnterKeyValuePair is called when entering the keyValuePair production.
	EnterKeyValuePair(c *KeyValuePairContext)

	// EnterDefaultValuePair is called when entering the defaultValuePair production.
	EnterDefaultValuePair(c *DefaultValuePairContext)

	// EnterKeyValue is called when entering the keyValue production.
	EnterKeyValue(c *KeyValueContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitDelimiters is called when exiting the delimiters production.
	ExitDelimiters(c *DelimitersContext)

	// ExitImports is called when exiting the imports production.
	ExitImports(c *ImportsContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitFormalArgs is called when exiting the formalArgs production.
	ExitFormalArgs(c *FormalArgsContext)

	// ExitFormalArg is called when exiting the formalArg production.
	ExitFormalArg(c *FormalArgContext)

	// ExitDict is called when exiting the dict production.
	ExitDict(c *DictContext)

	// ExitDictPairs is called when exiting the dictPairs production.
	ExitDictPairs(c *DictPairsContext)

	// ExitKeyValuePair is called when exiting the keyValuePair production.
	ExitKeyValuePair(c *KeyValuePairContext)

	// ExitDefaultValuePair is called when exiting the defaultValuePair production.
	ExitDefaultValuePair(c *DefaultValuePairContext)

	// ExitKeyValue is called when exiting the keyValue production.
	ExitKeyValue(c *KeyValueContext)
}
