// Code generated from metric.g4 by ANTLR 4.7.2. DO NOT EDIT.

package metric // metric
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasemetricListener is a complete listener for a parse tree produced by metricParser.
type BasemetricListener struct{}

var _ metricListener = &BasemetricListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemetricListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemetricListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemetricListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemetricListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterUom is called when production uom is entered.
func (s *BasemetricListener) EnterUom(ctx *UomContext) {}

// ExitUom is called when production uom is exited.
func (s *BasemetricListener) ExitUom(ctx *UomContext) {}

// EnterMeasure is called when production measure is entered.
func (s *BasemetricListener) EnterMeasure(ctx *MeasureContext) {}

// ExitMeasure is called when production measure is exited.
func (s *BasemetricListener) ExitMeasure(ctx *MeasureContext) {}

// EnterExponent is called when production exponent is entered.
func (s *BasemetricListener) EnterExponent(ctx *ExponentContext) {}

// ExitExponent is called when production exponent is exited.
func (s *BasemetricListener) ExitExponent(ctx *ExponentContext) {}

// EnterPrefix is called when production prefix is entered.
func (s *BasemetricListener) EnterPrefix(ctx *PrefixContext) {}

// ExitPrefix is called when production prefix is exited.
func (s *BasemetricListener) ExitPrefix(ctx *PrefixContext) {}

// EnterUnit is called when production unit is entered.
func (s *BasemetricListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BasemetricListener) ExitUnit(ctx *UnitContext) {}

// EnterBaseunit is called when production baseunit is entered.
func (s *BasemetricListener) EnterBaseunit(ctx *BaseunitContext) {}

// ExitBaseunit is called when production baseunit is exited.
func (s *BasemetricListener) ExitBaseunit(ctx *BaseunitContext) {}

// EnterDerivedunit is called when production derivedunit is entered.
func (s *BasemetricListener) EnterDerivedunit(ctx *DerivedunitContext) {}

// ExitDerivedunit is called when production derivedunit is exited.
func (s *BasemetricListener) ExitDerivedunit(ctx *DerivedunitContext) {}
