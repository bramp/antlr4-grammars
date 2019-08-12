// Code generated from smiles.g4 by ANTLR 4.7.2. DO NOT EDIT.

package smiles // smiles
import "github.com/antlr/antlr4/runtime/Go/antlr"

// smilesListener is a complete listener for a parse tree produced by smilesParser.
type smilesListener interface {
	antlr.ParseTreeListener

	// EnterSmiles is called when entering the smiles production.
	EnterSmiles(c *SmilesContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterAliphatic_organic is called when entering the aliphatic_organic production.
	EnterAliphatic_organic(c *Aliphatic_organicContext)

	// EnterAromatic_organic is called when entering the aromatic_organic production.
	EnterAromatic_organic(c *Aromatic_organicContext)

	// EnterBracket_atom is called when entering the bracket_atom production.
	EnterBracket_atom(c *Bracket_atomContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterIsotope is called when entering the isotope production.
	EnterIsotope(c *IsotopeContext)

	// EnterElement_symbols is called when entering the element_symbols production.
	EnterElement_symbols(c *Element_symbolsContext)

	// EnterAromatic_symbols is called when entering the aromatic_symbols production.
	EnterAromatic_symbols(c *Aromatic_symbolsContext)

	// EnterChiral is called when entering the chiral production.
	EnterChiral(c *ChiralContext)

	// EnterHcount is called when entering the hcount production.
	EnterHcount(c *HcountContext)

	// EnterCharge is called when entering the charge production.
	EnterCharge(c *ChargeContext)

	// EnterClass_ is called when entering the class_ production.
	EnterClass_(c *Class_Context)

	// EnterBond is called when entering the bond production.
	EnterBond(c *BondContext)

	// EnterRingbond is called when entering the ringbond production.
	EnterRingbond(c *RingbondContext)

	// EnterBranched_atom is called when entering the branched_atom production.
	EnterBranched_atom(c *Branched_atomContext)

	// EnterBranch is called when entering the branch production.
	EnterBranch(c *BranchContext)

	// EnterChain is called when entering the chain production.
	EnterChain(c *ChainContext)

	// EnterTerminator is called when entering the terminator production.
	EnterTerminator(c *TerminatorContext)

	// ExitSmiles is called when exiting the smiles production.
	ExitSmiles(c *SmilesContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitAliphatic_organic is called when exiting the aliphatic_organic production.
	ExitAliphatic_organic(c *Aliphatic_organicContext)

	// ExitAromatic_organic is called when exiting the aromatic_organic production.
	ExitAromatic_organic(c *Aromatic_organicContext)

	// ExitBracket_atom is called when exiting the bracket_atom production.
	ExitBracket_atom(c *Bracket_atomContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitIsotope is called when exiting the isotope production.
	ExitIsotope(c *IsotopeContext)

	// ExitElement_symbols is called when exiting the element_symbols production.
	ExitElement_symbols(c *Element_symbolsContext)

	// ExitAromatic_symbols is called when exiting the aromatic_symbols production.
	ExitAromatic_symbols(c *Aromatic_symbolsContext)

	// ExitChiral is called when exiting the chiral production.
	ExitChiral(c *ChiralContext)

	// ExitHcount is called when exiting the hcount production.
	ExitHcount(c *HcountContext)

	// ExitCharge is called when exiting the charge production.
	ExitCharge(c *ChargeContext)

	// ExitClass_ is called when exiting the class_ production.
	ExitClass_(c *Class_Context)

	// ExitBond is called when exiting the bond production.
	ExitBond(c *BondContext)

	// ExitRingbond is called when exiting the ringbond production.
	ExitRingbond(c *RingbondContext)

	// ExitBranched_atom is called when exiting the branched_atom production.
	ExitBranched_atom(c *Branched_atomContext)

	// ExitBranch is called when exiting the branch production.
	ExitBranch(c *BranchContext)

	// ExitChain is called when exiting the chain production.
	ExitChain(c *ChainContext)

	// ExitTerminator is called when exiting the terminator production.
	ExitTerminator(c *TerminatorContext)
}
