// Generated from xdr.g4 by ANTLR 4.7.

package xdr // xdr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasexdrListener is a complete listener for a parse tree produced by xdrParser.
type BasexdrListener struct{}

var _ xdrListener = &BasexdrListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasexdrListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasexdrListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasexdrListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasexdrListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasexdrListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasexdrListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterValue is called when production value is entered.
func (s *BasexdrListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasexdrListener) ExitValue(ctx *ValueContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasexdrListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasexdrListener) ExitConstant(ctx *ConstantContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BasexdrListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BasexdrListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterEnumTypeSpec is called when production enumTypeSpec is entered.
func (s *BasexdrListener) EnterEnumTypeSpec(ctx *EnumTypeSpecContext) {}

// ExitEnumTypeSpec is called when production enumTypeSpec is exited.
func (s *BasexdrListener) ExitEnumTypeSpec(ctx *EnumTypeSpecContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BasexdrListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BasexdrListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterStructTypeSpec is called when production structTypeSpec is entered.
func (s *BasexdrListener) EnterStructTypeSpec(ctx *StructTypeSpecContext) {}

// ExitStructTypeSpec is called when production structTypeSpec is exited.
func (s *BasexdrListener) ExitStructTypeSpec(ctx *StructTypeSpecContext) {}

// EnterStructBody is called when production structBody is entered.
func (s *BasexdrListener) EnterStructBody(ctx *StructBodyContext) {}

// ExitStructBody is called when production structBody is exited.
func (s *BasexdrListener) ExitStructBody(ctx *StructBodyContext) {}

// EnterUnionTypeSpec is called when production unionTypeSpec is entered.
func (s *BasexdrListener) EnterUnionTypeSpec(ctx *UnionTypeSpecContext) {}

// ExitUnionTypeSpec is called when production unionTypeSpec is exited.
func (s *BasexdrListener) ExitUnionTypeSpec(ctx *UnionTypeSpecContext) {}

// EnterUnionBody is called when production unionBody is entered.
func (s *BasexdrListener) EnterUnionBody(ctx *UnionBodyContext) {}

// ExitUnionBody is called when production unionBody is exited.
func (s *BasexdrListener) ExitUnionBody(ctx *UnionBodyContext) {}

// EnterCaseSpec is called when production caseSpec is entered.
func (s *BasexdrListener) EnterCaseSpec(ctx *CaseSpecContext) {}

// ExitCaseSpec is called when production caseSpec is exited.
func (s *BasexdrListener) ExitCaseSpec(ctx *CaseSpecContext) {}

// EnterConstantDef is called when production constantDef is entered.
func (s *BasexdrListener) EnterConstantDef(ctx *ConstantDefContext) {}

// ExitConstantDef is called when production constantDef is exited.
func (s *BasexdrListener) ExitConstantDef(ctx *ConstantDefContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BasexdrListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BasexdrListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BasexdrListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BasexdrListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterXdrSpecification is called when production xdrSpecification is entered.
func (s *BasexdrListener) EnterXdrSpecification(ctx *XdrSpecificationContext) {}

// ExitXdrSpecification is called when production xdrSpecification is exited.
func (s *BasexdrListener) ExitXdrSpecification(ctx *XdrSpecificationContext) {}
