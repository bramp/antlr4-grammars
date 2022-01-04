// Code generated from JSON5.g4 by ANTLR 4.9.3. DO NOT EDIT.

package json5 // JSON5
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJSON5Listener is a complete listener for a parse tree produced by JSON5Parser.
type BaseJSON5Listener struct{}

var _ JSON5Listener = &BaseJSON5Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJSON5Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJSON5Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJSON5Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJSON5Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJson5 is called when production json5 is entered.
func (s *BaseJSON5Listener) EnterJson5(ctx *Json5Context) {}

// ExitJson5 is called when production json5 is exited.
func (s *BaseJSON5Listener) ExitJson5(ctx *Json5Context) {}

// EnterObj is called when production obj is entered.
func (s *BaseJSON5Listener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseJSON5Listener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseJSON5Listener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseJSON5Listener) ExitPair(ctx *PairContext) {}

// EnterKey is called when production key is entered.
func (s *BaseJSON5Listener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseJSON5Listener) ExitKey(ctx *KeyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseJSON5Listener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseJSON5Listener) ExitValue(ctx *ValueContext) {}

// EnterArr is called when production arr is entered.
func (s *BaseJSON5Listener) EnterArr(ctx *ArrContext) {}

// ExitArr is called when production arr is exited.
func (s *BaseJSON5Listener) ExitArr(ctx *ArrContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseJSON5Listener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseJSON5Listener) ExitNumber(ctx *NumberContext) {}
