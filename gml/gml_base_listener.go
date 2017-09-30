// Generated from gml.g4 by ANTLR 4.7.

package gml // gml
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasegmlListener is a complete listener for a parse tree produced by gmlParser.
type BasegmlListener struct{}

var _ gmlListener = &BasegmlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegmlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegmlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegmlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegmlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGraph is called when production graph is entered.
func (s *BasegmlListener) EnterGraph(ctx *GraphContext) {}

// ExitGraph is called when production graph is exited.
func (s *BasegmlListener) ExitGraph(ctx *GraphContext) {}

// EnterList is called when production list is entered.
func (s *BasegmlListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BasegmlListener) ExitList(ctx *ListContext) {}

// EnterKv is called when production kv is entered.
func (s *BasegmlListener) EnterKv(ctx *KvContext) {}

// ExitKv is called when production kv is exited.
func (s *BasegmlListener) ExitKv(ctx *KvContext) {}

// EnterValue is called when production value is entered.
func (s *BasegmlListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasegmlListener) ExitValue(ctx *ValueContext) {}

// EnterKey is called when production key is entered.
func (s *BasegmlListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BasegmlListener) ExitKey(ctx *KeyContext) {}

// EnterInteger is called when production integer is entered.
func (s *BasegmlListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasegmlListener) ExitInteger(ctx *IntegerContext) {}

// EnterRealnum is called when production realnum is entered.
func (s *BasegmlListener) EnterRealnum(ctx *RealnumContext) {}

// ExitRealnum is called when production realnum is exited.
func (s *BasegmlListener) ExitRealnum(ctx *RealnumContext) {}

// EnterStr is called when production str is entered.
func (s *BasegmlListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BasegmlListener) ExitStr(ctx *StrContext) {}

// EnterStringliteral is called when production stringliteral is entered.
func (s *BasegmlListener) EnterStringliteral(ctx *StringliteralContext) {}

// ExitStringliteral is called when production stringliteral is exited.
func (s *BasegmlListener) ExitStringliteral(ctx *StringliteralContext) {}
