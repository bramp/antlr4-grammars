// Code generated from Protobuf3.g4 by ANTLR 4.9.3. DO NOT EDIT.

package protobuf3 // Protobuf3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProtobuf3Listener is a complete listener for a parse tree produced by Protobuf3Parser.
type BaseProtobuf3Listener struct{}

var _ Protobuf3Listener = &BaseProtobuf3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProtobuf3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProtobuf3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProtobuf3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProtobuf3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProto is called when production proto is entered.
func (s *BaseProtobuf3Listener) EnterProto(ctx *ProtoContext) {}

// ExitProto is called when production proto is exited.
func (s *BaseProtobuf3Listener) ExitProto(ctx *ProtoContext) {}

// EnterSyntax is called when production syntax is entered.
func (s *BaseProtobuf3Listener) EnterSyntax(ctx *SyntaxContext) {}

// ExitSyntax is called when production syntax is exited.
func (s *BaseProtobuf3Listener) ExitSyntax(ctx *SyntaxContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseProtobuf3Listener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseProtobuf3Listener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterPackageStatement is called when production packageStatement is entered.
func (s *BaseProtobuf3Listener) EnterPackageStatement(ctx *PackageStatementContext) {}

// ExitPackageStatement is called when production packageStatement is exited.
func (s *BaseProtobuf3Listener) ExitPackageStatement(ctx *PackageStatementContext) {}

// EnterOptionStatement is called when production optionStatement is entered.
func (s *BaseProtobuf3Listener) EnterOptionStatement(ctx *OptionStatementContext) {}

// ExitOptionStatement is called when production optionStatement is exited.
func (s *BaseProtobuf3Listener) ExitOptionStatement(ctx *OptionStatementContext) {}

// EnterOptionName is called when production optionName is entered.
func (s *BaseProtobuf3Listener) EnterOptionName(ctx *OptionNameContext) {}

// ExitOptionName is called when production optionName is exited.
func (s *BaseProtobuf3Listener) ExitOptionName(ctx *OptionNameContext) {}

// EnterField is called when production field is entered.
func (s *BaseProtobuf3Listener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseProtobuf3Listener) ExitField(ctx *FieldContext) {}

// EnterFieldOptions is called when production fieldOptions is entered.
func (s *BaseProtobuf3Listener) EnterFieldOptions(ctx *FieldOptionsContext) {}

// ExitFieldOptions is called when production fieldOptions is exited.
func (s *BaseProtobuf3Listener) ExitFieldOptions(ctx *FieldOptionsContext) {}

// EnterFieldOption is called when production fieldOption is entered.
func (s *BaseProtobuf3Listener) EnterFieldOption(ctx *FieldOptionContext) {}

// ExitFieldOption is called when production fieldOption is exited.
func (s *BaseProtobuf3Listener) ExitFieldOption(ctx *FieldOptionContext) {}

// EnterFieldNumber is called when production fieldNumber is entered.
func (s *BaseProtobuf3Listener) EnterFieldNumber(ctx *FieldNumberContext) {}

// ExitFieldNumber is called when production fieldNumber is exited.
func (s *BaseProtobuf3Listener) ExitFieldNumber(ctx *FieldNumberContext) {}

// EnterOneof is called when production oneof is entered.
func (s *BaseProtobuf3Listener) EnterOneof(ctx *OneofContext) {}

// ExitOneof is called when production oneof is exited.
func (s *BaseProtobuf3Listener) ExitOneof(ctx *OneofContext) {}

// EnterOneofField is called when production oneofField is entered.
func (s *BaseProtobuf3Listener) EnterOneofField(ctx *OneofFieldContext) {}

// ExitOneofField is called when production oneofField is exited.
func (s *BaseProtobuf3Listener) ExitOneofField(ctx *OneofFieldContext) {}

// EnterMapField is called when production mapField is entered.
func (s *BaseProtobuf3Listener) EnterMapField(ctx *MapFieldContext) {}

// ExitMapField is called when production mapField is exited.
func (s *BaseProtobuf3Listener) ExitMapField(ctx *MapFieldContext) {}

// EnterKeyType is called when production keyType is entered.
func (s *BaseProtobuf3Listener) EnterKeyType(ctx *KeyTypeContext) {}

// ExitKeyType is called when production keyType is exited.
func (s *BaseProtobuf3Listener) ExitKeyType(ctx *KeyTypeContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseProtobuf3Listener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseProtobuf3Listener) ExitType_(ctx *Type_Context) {}

// EnterReserved is called when production reserved is entered.
func (s *BaseProtobuf3Listener) EnterReserved(ctx *ReservedContext) {}

// ExitReserved is called when production reserved is exited.
func (s *BaseProtobuf3Listener) ExitReserved(ctx *ReservedContext) {}

// EnterRanges is called when production ranges is entered.
func (s *BaseProtobuf3Listener) EnterRanges(ctx *RangesContext) {}

// ExitRanges is called when production ranges is exited.
func (s *BaseProtobuf3Listener) ExitRanges(ctx *RangesContext) {}

// EnterRange_ is called when production range_ is entered.
func (s *BaseProtobuf3Listener) EnterRange_(ctx *Range_Context) {}

// ExitRange_ is called when production range_ is exited.
func (s *BaseProtobuf3Listener) ExitRange_(ctx *Range_Context) {}

// EnterReservedFieldNames is called when production reservedFieldNames is entered.
func (s *BaseProtobuf3Listener) EnterReservedFieldNames(ctx *ReservedFieldNamesContext) {}

// ExitReservedFieldNames is called when production reservedFieldNames is exited.
func (s *BaseProtobuf3Listener) ExitReservedFieldNames(ctx *ReservedFieldNamesContext) {}

// EnterTopLevelDef is called when production topLevelDef is entered.
func (s *BaseProtobuf3Listener) EnterTopLevelDef(ctx *TopLevelDefContext) {}

// ExitTopLevelDef is called when production topLevelDef is exited.
func (s *BaseProtobuf3Listener) ExitTopLevelDef(ctx *TopLevelDefContext) {}

// EnterEnumDef is called when production enumDef is entered.
func (s *BaseProtobuf3Listener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production enumDef is exited.
func (s *BaseProtobuf3Listener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseProtobuf3Listener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseProtobuf3Listener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumElement is called when production enumElement is entered.
func (s *BaseProtobuf3Listener) EnterEnumElement(ctx *EnumElementContext) {}

// ExitEnumElement is called when production enumElement is exited.
func (s *BaseProtobuf3Listener) ExitEnumElement(ctx *EnumElementContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *BaseProtobuf3Listener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *BaseProtobuf3Listener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterEnumValueOptions is called when production enumValueOptions is entered.
func (s *BaseProtobuf3Listener) EnterEnumValueOptions(ctx *EnumValueOptionsContext) {}

// ExitEnumValueOptions is called when production enumValueOptions is exited.
func (s *BaseProtobuf3Listener) ExitEnumValueOptions(ctx *EnumValueOptionsContext) {}

// EnterEnumValueOption is called when production enumValueOption is entered.
func (s *BaseProtobuf3Listener) EnterEnumValueOption(ctx *EnumValueOptionContext) {}

// ExitEnumValueOption is called when production enumValueOption is exited.
func (s *BaseProtobuf3Listener) ExitEnumValueOption(ctx *EnumValueOptionContext) {}

// EnterMessageDef is called when production messageDef is entered.
func (s *BaseProtobuf3Listener) EnterMessageDef(ctx *MessageDefContext) {}

// ExitMessageDef is called when production messageDef is exited.
func (s *BaseProtobuf3Listener) ExitMessageDef(ctx *MessageDefContext) {}

// EnterMessageBody is called when production messageBody is entered.
func (s *BaseProtobuf3Listener) EnterMessageBody(ctx *MessageBodyContext) {}

// ExitMessageBody is called when production messageBody is exited.
func (s *BaseProtobuf3Listener) ExitMessageBody(ctx *MessageBodyContext) {}

// EnterMessageElement is called when production messageElement is entered.
func (s *BaseProtobuf3Listener) EnterMessageElement(ctx *MessageElementContext) {}

// ExitMessageElement is called when production messageElement is exited.
func (s *BaseProtobuf3Listener) ExitMessageElement(ctx *MessageElementContext) {}

// EnterServiceDef is called when production serviceDef is entered.
func (s *BaseProtobuf3Listener) EnterServiceDef(ctx *ServiceDefContext) {}

// ExitServiceDef is called when production serviceDef is exited.
func (s *BaseProtobuf3Listener) ExitServiceDef(ctx *ServiceDefContext) {}

// EnterServiceElement is called when production serviceElement is entered.
func (s *BaseProtobuf3Listener) EnterServiceElement(ctx *ServiceElementContext) {}

// ExitServiceElement is called when production serviceElement is exited.
func (s *BaseProtobuf3Listener) ExitServiceElement(ctx *ServiceElementContext) {}

// EnterRpc is called when production rpc is entered.
func (s *BaseProtobuf3Listener) EnterRpc(ctx *RpcContext) {}

// ExitRpc is called when production rpc is exited.
func (s *BaseProtobuf3Listener) ExitRpc(ctx *RpcContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseProtobuf3Listener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseProtobuf3Listener) ExitConstant(ctx *ConstantContext) {}

// EnterBlockLit is called when production blockLit is entered.
func (s *BaseProtobuf3Listener) EnterBlockLit(ctx *BlockLitContext) {}

// ExitBlockLit is called when production blockLit is exited.
func (s *BaseProtobuf3Listener) ExitBlockLit(ctx *BlockLitContext) {}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BaseProtobuf3Listener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BaseProtobuf3Listener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}

// EnterIdent is called when production ident is entered.
func (s *BaseProtobuf3Listener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseProtobuf3Listener) ExitIdent(ctx *IdentContext) {}

// EnterFullIdent is called when production fullIdent is entered.
func (s *BaseProtobuf3Listener) EnterFullIdent(ctx *FullIdentContext) {}

// ExitFullIdent is called when production fullIdent is exited.
func (s *BaseProtobuf3Listener) ExitFullIdent(ctx *FullIdentContext) {}

// EnterMessageName is called when production messageName is entered.
func (s *BaseProtobuf3Listener) EnterMessageName(ctx *MessageNameContext) {}

// ExitMessageName is called when production messageName is exited.
func (s *BaseProtobuf3Listener) ExitMessageName(ctx *MessageNameContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *BaseProtobuf3Listener) EnterEnumName(ctx *EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *BaseProtobuf3Listener) ExitEnumName(ctx *EnumNameContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseProtobuf3Listener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseProtobuf3Listener) ExitFieldName(ctx *FieldNameContext) {}

// EnterOneofName is called when production oneofName is entered.
func (s *BaseProtobuf3Listener) EnterOneofName(ctx *OneofNameContext) {}

// ExitOneofName is called when production oneofName is exited.
func (s *BaseProtobuf3Listener) ExitOneofName(ctx *OneofNameContext) {}

// EnterMapName is called when production mapName is entered.
func (s *BaseProtobuf3Listener) EnterMapName(ctx *MapNameContext) {}

// ExitMapName is called when production mapName is exited.
func (s *BaseProtobuf3Listener) ExitMapName(ctx *MapNameContext) {}

// EnterServiceName is called when production serviceName is entered.
func (s *BaseProtobuf3Listener) EnterServiceName(ctx *ServiceNameContext) {}

// ExitServiceName is called when production serviceName is exited.
func (s *BaseProtobuf3Listener) ExitServiceName(ctx *ServiceNameContext) {}

// EnterRpcName is called when production rpcName is entered.
func (s *BaseProtobuf3Listener) EnterRpcName(ctx *RpcNameContext) {}

// ExitRpcName is called when production rpcName is exited.
func (s *BaseProtobuf3Listener) ExitRpcName(ctx *RpcNameContext) {}

// EnterMessageType is called when production messageType is entered.
func (s *BaseProtobuf3Listener) EnterMessageType(ctx *MessageTypeContext) {}

// ExitMessageType is called when production messageType is exited.
func (s *BaseProtobuf3Listener) ExitMessageType(ctx *MessageTypeContext) {}

// EnterEnumType is called when production enumType is entered.
func (s *BaseProtobuf3Listener) EnterEnumType(ctx *EnumTypeContext) {}

// ExitEnumType is called when production enumType is exited.
func (s *BaseProtobuf3Listener) ExitEnumType(ctx *EnumTypeContext) {}

// EnterIntLit is called when production intLit is entered.
func (s *BaseProtobuf3Listener) EnterIntLit(ctx *IntLitContext) {}

// ExitIntLit is called when production intLit is exited.
func (s *BaseProtobuf3Listener) ExitIntLit(ctx *IntLitContext) {}

// EnterStrLit is called when production strLit is entered.
func (s *BaseProtobuf3Listener) EnterStrLit(ctx *StrLitContext) {}

// ExitStrLit is called when production strLit is exited.
func (s *BaseProtobuf3Listener) ExitStrLit(ctx *StrLitContext) {}

// EnterBoolLit is called when production boolLit is entered.
func (s *BaseProtobuf3Listener) EnterBoolLit(ctx *BoolLitContext) {}

// ExitBoolLit is called when production boolLit is exited.
func (s *BaseProtobuf3Listener) ExitBoolLit(ctx *BoolLitContext) {}

// EnterFloatLit is called when production floatLit is entered.
func (s *BaseProtobuf3Listener) EnterFloatLit(ctx *FloatLitContext) {}

// ExitFloatLit is called when production floatLit is exited.
func (s *BaseProtobuf3Listener) ExitFloatLit(ctx *FloatLitContext) {}

// EnterKeywords is called when production keywords is entered.
func (s *BaseProtobuf3Listener) EnterKeywords(ctx *KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (s *BaseProtobuf3Listener) ExitKeywords(ctx *KeywordsContext) {}
