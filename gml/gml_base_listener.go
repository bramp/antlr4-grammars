// Code generated from gml.g4 by ANTLR 4.9.3. DO NOT EDIT.

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

// EnterList_ is called when production list_ is entered.
func (s *BasegmlListener) EnterList_(ctx *List_Context) {}

// ExitList_ is called when production list_ is exited.
func (s *BasegmlListener) ExitList_(ctx *List_Context) {}

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

// EnterStr_ is called when production str_ is entered.
func (s *BasegmlListener) EnterStr_(ctx *Str_Context) {}

// ExitStr_ is called when production str_ is exited.
func (s *BasegmlListener) ExitStr_(ctx *Str_Context) {}

// EnterStringliteral is called when production stringliteral is entered.
func (s *BasegmlListener) EnterStringliteral(ctx *StringliteralContext) {}

// ExitStringliteral is called when production stringliteral is exited.
func (s *BasegmlListener) ExitStringliteral(ctx *StringliteralContext) {}
