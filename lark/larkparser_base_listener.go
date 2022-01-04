// Code generated from LarkParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lark // LarkParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLarkParserListener is a complete listener for a parse tree produced by LarkParser.
type BaseLarkParserListener struct{}

var _ LarkParserListener = &BaseLarkParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLarkParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLarkParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLarkParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLarkParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseLarkParserListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseLarkParserListener) ExitStart(ctx *StartContext) {}

// EnterItem is called when production item is entered.
func (s *BaseLarkParserListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BaseLarkParserListener) ExitItem(ctx *ItemContext) {}

// EnterRule_ is called when production rule_ is entered.
func (s *BaseLarkParserListener) EnterRule_(ctx *Rule_Context) {}

// ExitRule_ is called when production rule_ is exited.
func (s *BaseLarkParserListener) ExitRule_(ctx *Rule_Context) {}

// EnterToken is called when production token is entered.
func (s *BaseLarkParserListener) EnterToken(ctx *TokenContext) {}

// ExitToken is called when production token is exited.
func (s *BaseLarkParserListener) ExitToken(ctx *TokenContext) {}

// EnterRule_params is called when production rule_params is entered.
func (s *BaseLarkParserListener) EnterRule_params(ctx *Rule_paramsContext) {}

// ExitRule_params is called when production rule_params is exited.
func (s *BaseLarkParserListener) ExitRule_params(ctx *Rule_paramsContext) {}

// EnterToken_params is called when production token_params is entered.
func (s *BaseLarkParserListener) EnterToken_params(ctx *Token_paramsContext) {}

// ExitToken_params is called when production token_params is exited.
func (s *BaseLarkParserListener) ExitToken_params(ctx *Token_paramsContext) {}

// EnterPriority is called when production priority is entered.
func (s *BaseLarkParserListener) EnterPriority(ctx *PriorityContext) {}

// ExitPriority is called when production priority is exited.
func (s *BaseLarkParserListener) ExitPriority(ctx *PriorityContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseLarkParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseLarkParserListener) ExitStatement(ctx *StatementContext) {}

// EnterImport_path is called when production import_path is entered.
func (s *BaseLarkParserListener) EnterImport_path(ctx *Import_pathContext) {}

// ExitImport_path is called when production import_path is exited.
func (s *BaseLarkParserListener) ExitImport_path(ctx *Import_pathContext) {}

// EnterName_list is called when production name_list is entered.
func (s *BaseLarkParserListener) EnterName_list(ctx *Name_listContext) {}

// ExitName_list is called when production name_list is exited.
func (s *BaseLarkParserListener) ExitName_list(ctx *Name_listContext) {}

// EnterExpansions is called when production expansions is entered.
func (s *BaseLarkParserListener) EnterExpansions(ctx *ExpansionsContext) {}

// ExitExpansions is called when production expansions is exited.
func (s *BaseLarkParserListener) ExitExpansions(ctx *ExpansionsContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseLarkParserListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseLarkParserListener) ExitAlias(ctx *AliasContext) {}

// EnterExpansion is called when production expansion is entered.
func (s *BaseLarkParserListener) EnterExpansion(ctx *ExpansionContext) {}

// ExitExpansion is called when production expansion is exited.
func (s *BaseLarkParserListener) ExitExpansion(ctx *ExpansionContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseLarkParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseLarkParserListener) ExitExpr(ctx *ExprContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseLarkParserListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseLarkParserListener) ExitAtom(ctx *AtomContext) {}

// EnterValue is called when production value is entered.
func (s *BaseLarkParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseLarkParserListener) ExitValue(ctx *ValueContext) {}

// EnterName is called when production name is entered.
func (s *BaseLarkParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseLarkParserListener) ExitName(ctx *NameContext) {}
