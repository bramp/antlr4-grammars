// Code generated from molecule.g4 by ANTLR 4.9.3. DO NOT EDIT.

package molecule // molecule
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasemoleculeListener is a complete listener for a parse tree produced by moleculeParser.
type BasemoleculeListener struct{}

var _ moleculeListener = &BasemoleculeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemoleculeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemoleculeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemoleculeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemoleculeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMolecule is called when production molecule is entered.
func (s *BasemoleculeListener) EnterMolecule(ctx *MoleculeContext) {}

// ExitMolecule is called when production molecule is exited.
func (s *BasemoleculeListener) ExitMolecule(ctx *MoleculeContext) {}

// EnterPart_ is called when production part_ is entered.
func (s *BasemoleculeListener) EnterPart_(ctx *Part_Context) {}

// ExitPart_ is called when production part_ is exited.
func (s *BasemoleculeListener) ExitPart_(ctx *Part_Context) {}

// EnterStructure is called when production structure is entered.
func (s *BasemoleculeListener) EnterStructure(ctx *StructureContext) {}

// ExitStructure is called when production structure is exited.
func (s *BasemoleculeListener) ExitStructure(ctx *StructureContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BasemoleculeListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BasemoleculeListener) ExitSymbol(ctx *SymbolContext) {}

// EnterGroup is called when production group is entered.
func (s *BasemoleculeListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BasemoleculeListener) ExitGroup(ctx *GroupContext) {}

// EnterIon is called when production ion is entered.
func (s *BasemoleculeListener) EnterIon(ctx *IonContext) {}

// ExitIon is called when production ion is exited.
func (s *BasemoleculeListener) ExitIon(ctx *IonContext) {}

// EnterElement is called when production element is entered.
func (s *BasemoleculeListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BasemoleculeListener) ExitElement(ctx *ElementContext) {}

// EnterCount is called when production count is entered.
func (s *BasemoleculeListener) EnterCount(ctx *CountContext) {}

// ExitCount is called when production count is exited.
func (s *BasemoleculeListener) ExitCount(ctx *CountContext) {}
