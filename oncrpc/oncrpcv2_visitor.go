// Generated from oncrpcv2.g4 by ANTLR 4.7.

package oncrpc // oncrpcv2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by oncrpcv2Parser.
type oncrpcv2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by oncrpcv2Parser#programDef.
	VisitProgramDef(ctx *ProgramDefContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#versionDef.
	VisitVersionDef(ctx *VersionDefContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#procedureDef.
	VisitProcedureDef(ctx *ProcedureDefContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#procReturn.
	VisitProcReturn(ctx *ProcReturnContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#procFirstArg.
	VisitProcFirstArg(ctx *ProcFirstArgContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#oncrpcv2Specification.
	VisitOncrpcv2Specification(ctx *Oncrpcv2SpecificationContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#typeSpecifier.
	VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#enumTypeSpec.
	VisitEnumTypeSpec(ctx *EnumTypeSpecContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#structTypeSpec.
	VisitStructTypeSpec(ctx *StructTypeSpecContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#structBody.
	VisitStructBody(ctx *StructBodyContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#unionTypeSpec.
	VisitUnionTypeSpec(ctx *UnionTypeSpecContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#unionBody.
	VisitUnionBody(ctx *UnionBodyContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#caseSpec.
	VisitCaseSpec(ctx *CaseSpecContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#constantDef.
	VisitConstantDef(ctx *ConstantDefContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#typeDef.
	VisitTypeDef(ctx *TypeDefContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by oncrpcv2Parser#xdrSpecification.
	VisitXdrSpecification(ctx *XdrSpecificationContext) interface{}
}
