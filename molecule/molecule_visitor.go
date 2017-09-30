// Generated from molecule.g4 by ANTLR 4.7.

package molecule // molecule
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by moleculeParser.
type moleculeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by moleculeParser#molecule.
	VisitMolecule(ctx *MoleculeContext) interface{}

	// Visit a parse tree produced by moleculeParser#part.
	VisitPart(ctx *PartContext) interface{}

	// Visit a parse tree produced by moleculeParser#structure.
	VisitStructure(ctx *StructureContext) interface{}

	// Visit a parse tree produced by moleculeParser#symbol.
	VisitSymbol(ctx *SymbolContext) interface{}

	// Visit a parse tree produced by moleculeParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by moleculeParser#ion.
	VisitIon(ctx *IonContext) interface{}

	// Visit a parse tree produced by moleculeParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by moleculeParser#count.
	VisitCount(ctx *CountContext) interface{}
}
