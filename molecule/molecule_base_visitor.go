// Generated from molecule.g4 by ANTLR 4.7.

package molecule // molecule
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasemoleculeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasemoleculeVisitor) VisitMolecule(ctx *MoleculeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemoleculeVisitor) VisitPart(ctx *PartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemoleculeVisitor) VisitStructure(ctx *StructureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemoleculeVisitor) VisitSymbol(ctx *SymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemoleculeVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemoleculeVisitor) VisitIon(ctx *IonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemoleculeVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemoleculeVisitor) VisitCount(ctx *CountContext) interface{} {
	return v.VisitChildren(ctx)
}
