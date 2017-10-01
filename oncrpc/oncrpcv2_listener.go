// Generated from oncrpcv2.g4 by ANTLR 4.7.

package oncrpc // oncrpcv2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// oncrpcv2Listener is a complete listener for a parse tree produced by oncrpcv2Parser.
type oncrpcv2Listener interface {
	antlr.ParseTreeListener

	// EnterProgramDef is called when entering the programDef production.
	EnterProgramDef(c *ProgramDefContext)

	// EnterVersionDef is called when entering the versionDef production.
	EnterVersionDef(c *VersionDefContext)

	// EnterProcedureDef is called when entering the procedureDef production.
	EnterProcedureDef(c *ProcedureDefContext)

	// EnterProcReturn is called when entering the procReturn production.
	EnterProcReturn(c *ProcReturnContext)

	// EnterProcFirstArg is called when entering the procFirstArg production.
	EnterProcFirstArg(c *ProcFirstArgContext)

	// EnterOncrpcv2Specification is called when entering the oncrpcv2Specification production.
	EnterOncrpcv2Specification(c *Oncrpcv2SpecificationContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterTypeSpecifier is called when entering the typeSpecifier production.
	EnterTypeSpecifier(c *TypeSpecifierContext)

	// EnterEnumTypeSpec is called when entering the enumTypeSpec production.
	EnterEnumTypeSpec(c *EnumTypeSpecContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterStructTypeSpec is called when entering the structTypeSpec production.
	EnterStructTypeSpec(c *StructTypeSpecContext)

	// EnterStructBody is called when entering the structBody production.
	EnterStructBody(c *StructBodyContext)

	// EnterUnionTypeSpec is called when entering the unionTypeSpec production.
	EnterUnionTypeSpec(c *UnionTypeSpecContext)

	// EnterUnionBody is called when entering the unionBody production.
	EnterUnionBody(c *UnionBodyContext)

	// EnterCaseSpec is called when entering the caseSpec production.
	EnterCaseSpec(c *CaseSpecContext)

	// EnterConstantDef is called when entering the constantDef production.
	EnterConstantDef(c *ConstantDefContext)

	// EnterTypeDef is called when entering the typeDef production.
	EnterTypeDef(c *TypeDefContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterXdrSpecification is called when entering the xdrSpecification production.
	EnterXdrSpecification(c *XdrSpecificationContext)

	// ExitProgramDef is called when exiting the programDef production.
	ExitProgramDef(c *ProgramDefContext)

	// ExitVersionDef is called when exiting the versionDef production.
	ExitVersionDef(c *VersionDefContext)

	// ExitProcedureDef is called when exiting the procedureDef production.
	ExitProcedureDef(c *ProcedureDefContext)

	// ExitProcReturn is called when exiting the procReturn production.
	ExitProcReturn(c *ProcReturnContext)

	// ExitProcFirstArg is called when exiting the procFirstArg production.
	ExitProcFirstArg(c *ProcFirstArgContext)

	// ExitOncrpcv2Specification is called when exiting the oncrpcv2Specification production.
	ExitOncrpcv2Specification(c *Oncrpcv2SpecificationContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitTypeSpecifier is called when exiting the typeSpecifier production.
	ExitTypeSpecifier(c *TypeSpecifierContext)

	// ExitEnumTypeSpec is called when exiting the enumTypeSpec production.
	ExitEnumTypeSpec(c *EnumTypeSpecContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitStructTypeSpec is called when exiting the structTypeSpec production.
	ExitStructTypeSpec(c *StructTypeSpecContext)

	// ExitStructBody is called when exiting the structBody production.
	ExitStructBody(c *StructBodyContext)

	// ExitUnionTypeSpec is called when exiting the unionTypeSpec production.
	ExitUnionTypeSpec(c *UnionTypeSpecContext)

	// ExitUnionBody is called when exiting the unionBody production.
	ExitUnionBody(c *UnionBodyContext)

	// ExitCaseSpec is called when exiting the caseSpec production.
	ExitCaseSpec(c *CaseSpecContext)

	// ExitConstantDef is called when exiting the constantDef production.
	ExitConstantDef(c *ConstantDefContext)

	// ExitTypeDef is called when exiting the typeDef production.
	ExitTypeDef(c *TypeDefContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitXdrSpecification is called when exiting the xdrSpecification production.
	ExitXdrSpecification(c *XdrSpecificationContext)
}
