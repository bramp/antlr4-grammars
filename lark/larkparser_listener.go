// Code generated from LarkParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lark // LarkParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// LarkParserListener is a complete listener for a parse tree produced by LarkParser.
type LarkParserListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterRule_ is called when entering the rule_ production.
	EnterRule_(c *Rule_Context)

	// EnterToken is called when entering the token production.
	EnterToken(c *TokenContext)

	// EnterRule_params is called when entering the rule_params production.
	EnterRule_params(c *Rule_paramsContext)

	// EnterToken_params is called when entering the token_params production.
	EnterToken_params(c *Token_paramsContext)

	// EnterPriority is called when entering the priority production.
	EnterPriority(c *PriorityContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterImport_path is called when entering the import_path production.
	EnterImport_path(c *Import_pathContext)

	// EnterName_list is called when entering the name_list production.
	EnterName_list(c *Name_listContext)

	// EnterExpansions is called when entering the expansions production.
	EnterExpansions(c *ExpansionsContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterExpansion is called when entering the expansion production.
	EnterExpansion(c *ExpansionContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitRule_ is called when exiting the rule_ production.
	ExitRule_(c *Rule_Context)

	// ExitToken is called when exiting the token production.
	ExitToken(c *TokenContext)

	// ExitRule_params is called when exiting the rule_params production.
	ExitRule_params(c *Rule_paramsContext)

	// ExitToken_params is called when exiting the token_params production.
	ExitToken_params(c *Token_paramsContext)

	// ExitPriority is called when exiting the priority production.
	ExitPriority(c *PriorityContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitImport_path is called when exiting the import_path production.
	ExitImport_path(c *Import_pathContext)

	// ExitName_list is called when exiting the name_list production.
	ExitName_list(c *Name_listContext)

	// ExitExpansions is called when exiting the expansions production.
	ExitExpansions(c *ExpansionsContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitExpansion is called when exiting the expansion production.
	ExitExpansion(c *ExpansionContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)
}
