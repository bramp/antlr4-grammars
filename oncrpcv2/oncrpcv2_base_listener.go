// Code generated from oncrpcv2.g4 by ANTLR 4.9.3. DO NOT EDIT.

package oncrpcv2 // oncrpcv2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Baseoncrpcv2Listener is a complete listener for a parse tree produced by oncrpcv2Parser.
type Baseoncrpcv2Listener struct{}

var _ oncrpcv2Listener = &Baseoncrpcv2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baseoncrpcv2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baseoncrpcv2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baseoncrpcv2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baseoncrpcv2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgramDef is called when production programDef is entered.
func (s *Baseoncrpcv2Listener) EnterProgramDef(ctx *ProgramDefContext) {}

// ExitProgramDef is called when production programDef is exited.
func (s *Baseoncrpcv2Listener) ExitProgramDef(ctx *ProgramDefContext) {}

// EnterVersionDef is called when production versionDef is entered.
func (s *Baseoncrpcv2Listener) EnterVersionDef(ctx *VersionDefContext) {}

// ExitVersionDef is called when production versionDef is exited.
func (s *Baseoncrpcv2Listener) ExitVersionDef(ctx *VersionDefContext) {}

// EnterProcedureDef is called when production procedureDef is entered.
func (s *Baseoncrpcv2Listener) EnterProcedureDef(ctx *ProcedureDefContext) {}

// ExitProcedureDef is called when production procedureDef is exited.
func (s *Baseoncrpcv2Listener) ExitProcedureDef(ctx *ProcedureDefContext) {}

// EnterProcReturn is called when production procReturn is entered.
func (s *Baseoncrpcv2Listener) EnterProcReturn(ctx *ProcReturnContext) {}

// ExitProcReturn is called when production procReturn is exited.
func (s *Baseoncrpcv2Listener) ExitProcReturn(ctx *ProcReturnContext) {}

// EnterProcFirstArg is called when production procFirstArg is entered.
func (s *Baseoncrpcv2Listener) EnterProcFirstArg(ctx *ProcFirstArgContext) {}

// ExitProcFirstArg is called when production procFirstArg is exited.
func (s *Baseoncrpcv2Listener) ExitProcFirstArg(ctx *ProcFirstArgContext) {}

// EnterOncrpcv2Specification is called when production oncrpcv2Specification is entered.
func (s *Baseoncrpcv2Listener) EnterOncrpcv2Specification(ctx *Oncrpcv2SpecificationContext) {}

// ExitOncrpcv2Specification is called when production oncrpcv2Specification is exited.
func (s *Baseoncrpcv2Listener) ExitOncrpcv2Specification(ctx *Oncrpcv2SpecificationContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *Baseoncrpcv2Listener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *Baseoncrpcv2Listener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterValue is called when production value is entered.
func (s *Baseoncrpcv2Listener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *Baseoncrpcv2Listener) ExitValue(ctx *ValueContext) {}

// EnterConstant is called when production constant is entered.
func (s *Baseoncrpcv2Listener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *Baseoncrpcv2Listener) ExitConstant(ctx *ConstantContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *Baseoncrpcv2Listener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *Baseoncrpcv2Listener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterEnumTypeSpec is called when production enumTypeSpec is entered.
func (s *Baseoncrpcv2Listener) EnterEnumTypeSpec(ctx *EnumTypeSpecContext) {}

// ExitEnumTypeSpec is called when production enumTypeSpec is exited.
func (s *Baseoncrpcv2Listener) ExitEnumTypeSpec(ctx *EnumTypeSpecContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *Baseoncrpcv2Listener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *Baseoncrpcv2Listener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterStructTypeSpec is called when production structTypeSpec is entered.
func (s *Baseoncrpcv2Listener) EnterStructTypeSpec(ctx *StructTypeSpecContext) {}

// ExitStructTypeSpec is called when production structTypeSpec is exited.
func (s *Baseoncrpcv2Listener) ExitStructTypeSpec(ctx *StructTypeSpecContext) {}

// EnterStructBody is called when production structBody is entered.
func (s *Baseoncrpcv2Listener) EnterStructBody(ctx *StructBodyContext) {}

// ExitStructBody is called when production structBody is exited.
func (s *Baseoncrpcv2Listener) ExitStructBody(ctx *StructBodyContext) {}

// EnterUnionTypeSpec is called when production unionTypeSpec is entered.
func (s *Baseoncrpcv2Listener) EnterUnionTypeSpec(ctx *UnionTypeSpecContext) {}

// ExitUnionTypeSpec is called when production unionTypeSpec is exited.
func (s *Baseoncrpcv2Listener) ExitUnionTypeSpec(ctx *UnionTypeSpecContext) {}

// EnterUnionBody is called when production unionBody is entered.
func (s *Baseoncrpcv2Listener) EnterUnionBody(ctx *UnionBodyContext) {}

// ExitUnionBody is called when production unionBody is exited.
func (s *Baseoncrpcv2Listener) ExitUnionBody(ctx *UnionBodyContext) {}

// EnterCaseSpec is called when production caseSpec is entered.
func (s *Baseoncrpcv2Listener) EnterCaseSpec(ctx *CaseSpecContext) {}

// ExitCaseSpec is called when production caseSpec is exited.
func (s *Baseoncrpcv2Listener) ExitCaseSpec(ctx *CaseSpecContext) {}

// EnterConstantDef is called when production constantDef is entered.
func (s *Baseoncrpcv2Listener) EnterConstantDef(ctx *ConstantDefContext) {}

// ExitConstantDef is called when production constantDef is exited.
func (s *Baseoncrpcv2Listener) ExitConstantDef(ctx *ConstantDefContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *Baseoncrpcv2Listener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *Baseoncrpcv2Listener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterDefinition is called when production definition is entered.
func (s *Baseoncrpcv2Listener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *Baseoncrpcv2Listener) ExitDefinition(ctx *DefinitionContext) {}

// EnterXdrSpecification is called when production xdrSpecification is entered.
func (s *Baseoncrpcv2Listener) EnterXdrSpecification(ctx *XdrSpecificationContext) {}

// ExitXdrSpecification is called when production xdrSpecification is exited.
func (s *Baseoncrpcv2Listener) ExitXdrSpecification(ctx *XdrSpecificationContext) {}
