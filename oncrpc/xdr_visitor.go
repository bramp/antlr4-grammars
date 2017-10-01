// Generated from xdr.g4 by ANTLR 4.7.

package oncrpc // xdr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by xdrParser.
type xdrVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by xdrParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by xdrParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by xdrParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by xdrParser#typeSpecifier.
	VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{}

	// Visit a parse tree produced by xdrParser#enumTypeSpec.
	VisitEnumTypeSpec(ctx *EnumTypeSpecContext) interface{}

	// Visit a parse tree produced by xdrParser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by xdrParser#structTypeSpec.
	VisitStructTypeSpec(ctx *StructTypeSpecContext) interface{}

	// Visit a parse tree produced by xdrParser#structBody.
	VisitStructBody(ctx *StructBodyContext) interface{}

	// Visit a parse tree produced by xdrParser#unionTypeSpec.
	VisitUnionTypeSpec(ctx *UnionTypeSpecContext) interface{}

	// Visit a parse tree produced by xdrParser#unionBody.
	VisitUnionBody(ctx *UnionBodyContext) interface{}

	// Visit a parse tree produced by xdrParser#caseSpec.
	VisitCaseSpec(ctx *CaseSpecContext) interface{}

	// Visit a parse tree produced by xdrParser#constantDef.
	VisitConstantDef(ctx *ConstantDefContext) interface{}

	// Visit a parse tree produced by xdrParser#typeDef.
	VisitTypeDef(ctx *TypeDefContext) interface{}

	// Visit a parse tree produced by xdrParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by xdrParser#xdrSpecification.
	VisitXdrSpecification(ctx *XdrSpecificationContext) interface{}
}
