// Code generated from DiceNotationParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dicenotation // DiceNotationParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDiceNotationParserListener is a complete listener for a parse tree produced by DiceNotationParser.
type BaseDiceNotationParserListener struct{}

var _ DiceNotationParserListener = &BaseDiceNotationParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiceNotationParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiceNotationParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiceNotationParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiceNotationParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNotation is called when production notation is entered.
func (s *BaseDiceNotationParserListener) EnterNotation(ctx *NotationContext) {}

// ExitNotation is called when production notation is exited.
func (s *BaseDiceNotationParserListener) ExitNotation(ctx *NotationContext) {}

// EnterAddOp is called when production addOp is entered.
func (s *BaseDiceNotationParserListener) EnterAddOp(ctx *AddOpContext) {}

// ExitAddOp is called when production addOp is exited.
func (s *BaseDiceNotationParserListener) ExitAddOp(ctx *AddOpContext) {}

// EnterMultOp is called when production multOp is entered.
func (s *BaseDiceNotationParserListener) EnterMultOp(ctx *MultOpContext) {}

// ExitMultOp is called when production multOp is exited.
func (s *BaseDiceNotationParserListener) ExitMultOp(ctx *MultOpContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseDiceNotationParserListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseDiceNotationParserListener) ExitOperand(ctx *OperandContext) {}

// EnterDice is called when production dice is entered.
func (s *BaseDiceNotationParserListener) EnterDice(ctx *DiceContext) {}

// ExitDice is called when production dice is exited.
func (s *BaseDiceNotationParserListener) ExitDice(ctx *DiceContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseDiceNotationParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseDiceNotationParserListener) ExitNumber(ctx *NumberContext) {}
