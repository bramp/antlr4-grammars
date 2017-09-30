// Generated from molecule.g4 by ANTLR 4.7.

package molecule // molecule
import "github.com/antlr/antlr4/runtime/Go/antlr"

// moleculeListener is a complete listener for a parse tree produced by moleculeParser.
type moleculeListener interface {
	antlr.ParseTreeListener

	// EnterMolecule is called when entering the molecule production.
	EnterMolecule(c *MoleculeContext)

	// EnterPart is called when entering the part production.
	EnterPart(c *PartContext)

	// EnterStructure is called when entering the structure production.
	EnterStructure(c *StructureContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterIon is called when entering the ion production.
	EnterIon(c *IonContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterCount is called when entering the count production.
	EnterCount(c *CountContext)

	// ExitMolecule is called when exiting the molecule production.
	ExitMolecule(c *MoleculeContext)

	// ExitPart is called when exiting the part production.
	ExitPart(c *PartContext)

	// ExitStructure is called when exiting the structure production.
	ExitStructure(c *StructureContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitIon is called when exiting the ion production.
	ExitIon(c *IonContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitCount is called when exiting the count production.
	ExitCount(c *CountContext)
}
