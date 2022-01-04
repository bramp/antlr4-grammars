// Code generated from edn.g4 by ANTLR 4.9.3. DO NOT EDIT.

package edn // edn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseednListener is a complete listener for a parse tree produced by ednParser.
type BaseednListener struct{}

var _ ednListener = &BaseednListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseednListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseednListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseednListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseednListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseednListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseednListener) ExitProgram(ctx *ProgramContext) {}

// EnterValue is called when production value is entered.
func (s *BaseednListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseednListener) ExitValue(ctx *ValueContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseednListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseednListener) ExitTag(ctx *TagContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseednListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseednListener) ExitKeyword(ctx *KeywordContext) {}

// EnterList_ is called when production list_ is entered.
func (s *BaseednListener) EnterList_(ctx *List_Context) {}

// ExitList_ is called when production list_ is exited.
func (s *BaseednListener) ExitList_(ctx *List_Context) {}

// EnterVector is called when production vector is entered.
func (s *BaseednListener) EnterVector(ctx *VectorContext) {}

// ExitVector is called when production vector is exited.
func (s *BaseednListener) ExitVector(ctx *VectorContext) {}

// EnterMap_ is called when production map_ is entered.
func (s *BaseednListener) EnterMap_(ctx *Map_Context) {}

// ExitMap_ is called when production map_ is exited.
func (s *BaseednListener) ExitMap_(ctx *Map_Context) {}

// EnterSet_ is called when production set_ is entered.
func (s *BaseednListener) EnterSet_(ctx *Set_Context) {}

// ExitSet_ is called when production set_ is exited.
func (s *BaseednListener) ExitSet_(ctx *Set_Context) {}
