// Code generated from smiles.g4 by ANTLR 4.9.3. DO NOT EDIT.

package smiles // smiles
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesmilesListener is a complete listener for a parse tree produced by smilesParser.
type BasesmilesListener struct{}

var _ smilesListener = &BasesmilesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesmilesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesmilesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesmilesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesmilesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSmiles is called when production smiles is entered.
func (s *BasesmilesListener) EnterSmiles(ctx *SmilesContext) {}

// ExitSmiles is called when production smiles is exited.
func (s *BasesmilesListener) ExitSmiles(ctx *SmilesContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasesmilesListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasesmilesListener) ExitAtom(ctx *AtomContext) {}

// EnterAliphatic_organic is called when production aliphatic_organic is entered.
func (s *BasesmilesListener) EnterAliphatic_organic(ctx *Aliphatic_organicContext) {}

// ExitAliphatic_organic is called when production aliphatic_organic is exited.
func (s *BasesmilesListener) ExitAliphatic_organic(ctx *Aliphatic_organicContext) {}

// EnterAromatic_organic is called when production aromatic_organic is entered.
func (s *BasesmilesListener) EnterAromatic_organic(ctx *Aromatic_organicContext) {}

// ExitAromatic_organic is called when production aromatic_organic is exited.
func (s *BasesmilesListener) ExitAromatic_organic(ctx *Aromatic_organicContext) {}

// EnterBracket_atom is called when production bracket_atom is entered.
func (s *BasesmilesListener) EnterBracket_atom(ctx *Bracket_atomContext) {}

// ExitBracket_atom is called when production bracket_atom is exited.
func (s *BasesmilesListener) ExitBracket_atom(ctx *Bracket_atomContext) {}

// EnterIsotope is called when production isotope is entered.
func (s *BasesmilesListener) EnterIsotope(ctx *IsotopeContext) {}

// ExitIsotope is called when production isotope is exited.
func (s *BasesmilesListener) ExitIsotope(ctx *IsotopeContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BasesmilesListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BasesmilesListener) ExitSymbol(ctx *SymbolContext) {}

// EnterElement_symbols is called when production element_symbols is entered.
func (s *BasesmilesListener) EnterElement_symbols(ctx *Element_symbolsContext) {}

// ExitElement_symbols is called when production element_symbols is exited.
func (s *BasesmilesListener) ExitElement_symbols(ctx *Element_symbolsContext) {}

// EnterAromatic_symbol is called when production aromatic_symbol is entered.
func (s *BasesmilesListener) EnterAromatic_symbol(ctx *Aromatic_symbolContext) {}

// ExitAromatic_symbol is called when production aromatic_symbol is exited.
func (s *BasesmilesListener) ExitAromatic_symbol(ctx *Aromatic_symbolContext) {}

// EnterChiral is called when production chiral is entered.
func (s *BasesmilesListener) EnterChiral(ctx *ChiralContext) {}

// ExitChiral is called when production chiral is exited.
func (s *BasesmilesListener) ExitChiral(ctx *ChiralContext) {}

// EnterHcount is called when production hcount is entered.
func (s *BasesmilesListener) EnterHcount(ctx *HcountContext) {}

// ExitHcount is called when production hcount is exited.
func (s *BasesmilesListener) ExitHcount(ctx *HcountContext) {}

// EnterCharge is called when production charge is entered.
func (s *BasesmilesListener) EnterCharge(ctx *ChargeContext) {}

// ExitCharge is called when production charge is exited.
func (s *BasesmilesListener) ExitCharge(ctx *ChargeContext) {}

// EnterClass_ is called when production class_ is entered.
func (s *BasesmilesListener) EnterClass_(ctx *Class_Context) {}

// ExitClass_ is called when production class_ is exited.
func (s *BasesmilesListener) ExitClass_(ctx *Class_Context) {}

// EnterBond is called when production bond is entered.
func (s *BasesmilesListener) EnterBond(ctx *BondContext) {}

// ExitBond is called when production bond is exited.
func (s *BasesmilesListener) ExitBond(ctx *BondContext) {}

// EnterRingbond is called when production ringbond is entered.
func (s *BasesmilesListener) EnterRingbond(ctx *RingbondContext) {}

// ExitRingbond is called when production ringbond is exited.
func (s *BasesmilesListener) ExitRingbond(ctx *RingbondContext) {}

// EnterBranched_atom is called when production branched_atom is entered.
func (s *BasesmilesListener) EnterBranched_atom(ctx *Branched_atomContext) {}

// ExitBranched_atom is called when production branched_atom is exited.
func (s *BasesmilesListener) ExitBranched_atom(ctx *Branched_atomContext) {}

// EnterBranch is called when production branch is entered.
func (s *BasesmilesListener) EnterBranch(ctx *BranchContext) {}

// ExitBranch is called when production branch is exited.
func (s *BasesmilesListener) ExitBranch(ctx *BranchContext) {}

// EnterChain is called when production chain is entered.
func (s *BasesmilesListener) EnterChain(ctx *ChainContext) {}

// ExitChain is called when production chain is exited.
func (s *BasesmilesListener) ExitChain(ctx *ChainContext) {}

// EnterTerminator is called when production terminator is entered.
func (s *BasesmilesListener) EnterTerminator(ctx *TerminatorContext) {}

// ExitTerminator is called when production terminator is exited.
func (s *BasesmilesListener) ExitTerminator(ctx *TerminatorContext) {}
