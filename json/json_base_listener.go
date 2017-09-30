// Generated from JSON.g4 by ANTLR 4.7.

package json // JSON
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJSONListener is a complete listener for a parse tree produced by JSONParser.
type BaseJSONListener struct{}

var _ JSONListener = &BaseJSONListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJSONListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJSONListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJSONListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJSONListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJson is called when production json is entered.
func (s *BaseJSONListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseJSONListener) ExitJson(ctx *JsonContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseJSONListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseJSONListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseJSONListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseJSONListener) ExitPair(ctx *PairContext) {}

// EnterArray is called when production array is entered.
func (s *BaseJSONListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseJSONListener) ExitArray(ctx *ArrayContext) {}

// EnterValue is called when production value is entered.
func (s *BaseJSONListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseJSONListener) ExitValue(ctx *ValueContext) {}
