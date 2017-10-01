// Generated from xdr.g4 by ANTLR 4.7.

package oncrpc // xdr
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasexdrVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasexdrVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitEnumTypeSpec(ctx *EnumTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitStructTypeSpec(ctx *StructTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitStructBody(ctx *StructBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitUnionTypeSpec(ctx *UnionTypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitUnionBody(ctx *UnionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitCaseSpec(ctx *CaseSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitConstantDef(ctx *ConstantDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitTypeDef(ctx *TypeDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasexdrVisitor) VisitXdrSpecification(ctx *XdrSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}
