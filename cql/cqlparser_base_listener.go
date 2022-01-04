// Code generated from CqlParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cql // CqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCqlParserListener is a complete listener for a parse tree produced by CqlParser.
type BaseCqlParserListener struct{}

var _ CqlParserListener = &BaseCqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseCqlParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseCqlParserListener) ExitRoot(ctx *RootContext) {}

// EnterCqls is called when production cqls is entered.
func (s *BaseCqlParserListener) EnterCqls(ctx *CqlsContext) {}

// ExitCqls is called when production cqls is exited.
func (s *BaseCqlParserListener) ExitCqls(ctx *CqlsContext) {}

// EnterStatementSeparator is called when production statementSeparator is entered.
func (s *BaseCqlParserListener) EnterStatementSeparator(ctx *StatementSeparatorContext) {}

// ExitStatementSeparator is called when production statementSeparator is exited.
func (s *BaseCqlParserListener) ExitStatementSeparator(ctx *StatementSeparatorContext) {}

// EnterEmpty_ is called when production empty_ is entered.
func (s *BaseCqlParserListener) EnterEmpty_(ctx *Empty_Context) {}

// ExitEmpty_ is called when production empty_ is exited.
func (s *BaseCqlParserListener) ExitEmpty_(ctx *Empty_Context) {}

// EnterCql is called when production cql is entered.
func (s *BaseCqlParserListener) EnterCql(ctx *CqlContext) {}

// ExitCql is called when production cql is exited.
func (s *BaseCqlParserListener) ExitCql(ctx *CqlContext) {}

// EnterRevoke is called when production revoke is entered.
func (s *BaseCqlParserListener) EnterRevoke(ctx *RevokeContext) {}

// ExitRevoke is called when production revoke is exited.
func (s *BaseCqlParserListener) ExitRevoke(ctx *RevokeContext) {}

// EnterListUsers is called when production listUsers is entered.
func (s *BaseCqlParserListener) EnterListUsers(ctx *ListUsersContext) {}

// ExitListUsers is called when production listUsers is exited.
func (s *BaseCqlParserListener) ExitListUsers(ctx *ListUsersContext) {}

// EnterListRoles is called when production listRoles is entered.
func (s *BaseCqlParserListener) EnterListRoles(ctx *ListRolesContext) {}

// ExitListRoles is called when production listRoles is exited.
func (s *BaseCqlParserListener) ExitListRoles(ctx *ListRolesContext) {}

// EnterListPermissions is called when production listPermissions is entered.
func (s *BaseCqlParserListener) EnterListPermissions(ctx *ListPermissionsContext) {}

// ExitListPermissions is called when production listPermissions is exited.
func (s *BaseCqlParserListener) ExitListPermissions(ctx *ListPermissionsContext) {}

// EnterGrant is called when production grant is entered.
func (s *BaseCqlParserListener) EnterGrant(ctx *GrantContext) {}

// ExitGrant is called when production grant is exited.
func (s *BaseCqlParserListener) ExitGrant(ctx *GrantContext) {}

// EnterPriviledge is called when production priviledge is entered.
func (s *BaseCqlParserListener) EnterPriviledge(ctx *PriviledgeContext) {}

// ExitPriviledge is called when production priviledge is exited.
func (s *BaseCqlParserListener) ExitPriviledge(ctx *PriviledgeContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseCqlParserListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseCqlParserListener) ExitResource(ctx *ResourceContext) {}

// EnterCreateUser is called when production createUser is entered.
func (s *BaseCqlParserListener) EnterCreateUser(ctx *CreateUserContext) {}

// ExitCreateUser is called when production createUser is exited.
func (s *BaseCqlParserListener) ExitCreateUser(ctx *CreateUserContext) {}

// EnterCreateRole is called when production createRole is entered.
func (s *BaseCqlParserListener) EnterCreateRole(ctx *CreateRoleContext) {}

// ExitCreateRole is called when production createRole is exited.
func (s *BaseCqlParserListener) ExitCreateRole(ctx *CreateRoleContext) {}

// EnterCreateType is called when production createType is entered.
func (s *BaseCqlParserListener) EnterCreateType(ctx *CreateTypeContext) {}

// ExitCreateType is called when production createType is exited.
func (s *BaseCqlParserListener) ExitCreateType(ctx *CreateTypeContext) {}

// EnterTypeMemberColumnList is called when production typeMemberColumnList is entered.
func (s *BaseCqlParserListener) EnterTypeMemberColumnList(ctx *TypeMemberColumnListContext) {}

// ExitTypeMemberColumnList is called when production typeMemberColumnList is exited.
func (s *BaseCqlParserListener) ExitTypeMemberColumnList(ctx *TypeMemberColumnListContext) {}

// EnterCreateTrigger is called when production createTrigger is entered.
func (s *BaseCqlParserListener) EnterCreateTrigger(ctx *CreateTriggerContext) {}

// ExitCreateTrigger is called when production createTrigger is exited.
func (s *BaseCqlParserListener) ExitCreateTrigger(ctx *CreateTriggerContext) {}

// EnterCreateMaterializedView is called when production createMaterializedView is entered.
func (s *BaseCqlParserListener) EnterCreateMaterializedView(ctx *CreateMaterializedViewContext) {}

// ExitCreateMaterializedView is called when production createMaterializedView is exited.
func (s *BaseCqlParserListener) ExitCreateMaterializedView(ctx *CreateMaterializedViewContext) {}

// EnterMaterializedViewWhere is called when production materializedViewWhere is entered.
func (s *BaseCqlParserListener) EnterMaterializedViewWhere(ctx *MaterializedViewWhereContext) {}

// ExitMaterializedViewWhere is called when production materializedViewWhere is exited.
func (s *BaseCqlParserListener) ExitMaterializedViewWhere(ctx *MaterializedViewWhereContext) {}

// EnterColumnNotNullList is called when production columnNotNullList is entered.
func (s *BaseCqlParserListener) EnterColumnNotNullList(ctx *ColumnNotNullListContext) {}

// ExitColumnNotNullList is called when production columnNotNullList is exited.
func (s *BaseCqlParserListener) ExitColumnNotNullList(ctx *ColumnNotNullListContext) {}

// EnterColumnNotNull is called when production columnNotNull is entered.
func (s *BaseCqlParserListener) EnterColumnNotNull(ctx *ColumnNotNullContext) {}

// ExitColumnNotNull is called when production columnNotNull is exited.
func (s *BaseCqlParserListener) ExitColumnNotNull(ctx *ColumnNotNullContext) {}

// EnterMaterializedViewOptions is called when production materializedViewOptions is entered.
func (s *BaseCqlParserListener) EnterMaterializedViewOptions(ctx *MaterializedViewOptionsContext) {}

// ExitMaterializedViewOptions is called when production materializedViewOptions is exited.
func (s *BaseCqlParserListener) ExitMaterializedViewOptions(ctx *MaterializedViewOptionsContext) {}

// EnterCreateKeyspace is called when production createKeyspace is entered.
func (s *BaseCqlParserListener) EnterCreateKeyspace(ctx *CreateKeyspaceContext) {}

// ExitCreateKeyspace is called when production createKeyspace is exited.
func (s *BaseCqlParserListener) ExitCreateKeyspace(ctx *CreateKeyspaceContext) {}

// EnterCreateFunction is called when production createFunction is entered.
func (s *BaseCqlParserListener) EnterCreateFunction(ctx *CreateFunctionContext) {}

// ExitCreateFunction is called when production createFunction is exited.
func (s *BaseCqlParserListener) ExitCreateFunction(ctx *CreateFunctionContext) {}

// EnterCodeBlock is called when production codeBlock is entered.
func (s *BaseCqlParserListener) EnterCodeBlock(ctx *CodeBlockContext) {}

// ExitCodeBlock is called when production codeBlock is exited.
func (s *BaseCqlParserListener) ExitCodeBlock(ctx *CodeBlockContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseCqlParserListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseCqlParserListener) ExitParamList(ctx *ParamListContext) {}

// EnterReturnMode is called when production returnMode is entered.
func (s *BaseCqlParserListener) EnterReturnMode(ctx *ReturnModeContext) {}

// ExitReturnMode is called when production returnMode is exited.
func (s *BaseCqlParserListener) ExitReturnMode(ctx *ReturnModeContext) {}

// EnterCreateAggregate is called when production createAggregate is entered.
func (s *BaseCqlParserListener) EnterCreateAggregate(ctx *CreateAggregateContext) {}

// ExitCreateAggregate is called when production createAggregate is exited.
func (s *BaseCqlParserListener) ExitCreateAggregate(ctx *CreateAggregateContext) {}

// EnterInitCondDefinition is called when production initCondDefinition is entered.
func (s *BaseCqlParserListener) EnterInitCondDefinition(ctx *InitCondDefinitionContext) {}

// ExitInitCondDefinition is called when production initCondDefinition is exited.
func (s *BaseCqlParserListener) ExitInitCondDefinition(ctx *InitCondDefinitionContext) {}

// EnterInitCondHash is called when production initCondHash is entered.
func (s *BaseCqlParserListener) EnterInitCondHash(ctx *InitCondHashContext) {}

// ExitInitCondHash is called when production initCondHash is exited.
func (s *BaseCqlParserListener) ExitInitCondHash(ctx *InitCondHashContext) {}

// EnterInitCondHashItem is called when production initCondHashItem is entered.
func (s *BaseCqlParserListener) EnterInitCondHashItem(ctx *InitCondHashItemContext) {}

// ExitInitCondHashItem is called when production initCondHashItem is exited.
func (s *BaseCqlParserListener) ExitInitCondHashItem(ctx *InitCondHashItemContext) {}

// EnterInitCondListNested is called when production initCondListNested is entered.
func (s *BaseCqlParserListener) EnterInitCondListNested(ctx *InitCondListNestedContext) {}

// ExitInitCondListNested is called when production initCondListNested is exited.
func (s *BaseCqlParserListener) ExitInitCondListNested(ctx *InitCondListNestedContext) {}

// EnterInitCondList is called when production initCondList is entered.
func (s *BaseCqlParserListener) EnterInitCondList(ctx *InitCondListContext) {}

// ExitInitCondList is called when production initCondList is exited.
func (s *BaseCqlParserListener) ExitInitCondList(ctx *InitCondListContext) {}

// EnterOrReplace is called when production orReplace is entered.
func (s *BaseCqlParserListener) EnterOrReplace(ctx *OrReplaceContext) {}

// ExitOrReplace is called when production orReplace is exited.
func (s *BaseCqlParserListener) ExitOrReplace(ctx *OrReplaceContext) {}

// EnterAlterUser is called when production alterUser is entered.
func (s *BaseCqlParserListener) EnterAlterUser(ctx *AlterUserContext) {}

// ExitAlterUser is called when production alterUser is exited.
func (s *BaseCqlParserListener) ExitAlterUser(ctx *AlterUserContext) {}

// EnterUserPassword is called when production userPassword is entered.
func (s *BaseCqlParserListener) EnterUserPassword(ctx *UserPasswordContext) {}

// ExitUserPassword is called when production userPassword is exited.
func (s *BaseCqlParserListener) ExitUserPassword(ctx *UserPasswordContext) {}

// EnterUserSuperUser is called when production userSuperUser is entered.
func (s *BaseCqlParserListener) EnterUserSuperUser(ctx *UserSuperUserContext) {}

// ExitUserSuperUser is called when production userSuperUser is exited.
func (s *BaseCqlParserListener) ExitUserSuperUser(ctx *UserSuperUserContext) {}

// EnterAlterType is called when production alterType is entered.
func (s *BaseCqlParserListener) EnterAlterType(ctx *AlterTypeContext) {}

// ExitAlterType is called when production alterType is exited.
func (s *BaseCqlParserListener) ExitAlterType(ctx *AlterTypeContext) {}

// EnterAlterTypeOperation is called when production alterTypeOperation is entered.
func (s *BaseCqlParserListener) EnterAlterTypeOperation(ctx *AlterTypeOperationContext) {}

// ExitAlterTypeOperation is called when production alterTypeOperation is exited.
func (s *BaseCqlParserListener) ExitAlterTypeOperation(ctx *AlterTypeOperationContext) {}

// EnterAlterTypeRename is called when production alterTypeRename is entered.
func (s *BaseCqlParserListener) EnterAlterTypeRename(ctx *AlterTypeRenameContext) {}

// ExitAlterTypeRename is called when production alterTypeRename is exited.
func (s *BaseCqlParserListener) ExitAlterTypeRename(ctx *AlterTypeRenameContext) {}

// EnterAlterTypeRenameList is called when production alterTypeRenameList is entered.
func (s *BaseCqlParserListener) EnterAlterTypeRenameList(ctx *AlterTypeRenameListContext) {}

// ExitAlterTypeRenameList is called when production alterTypeRenameList is exited.
func (s *BaseCqlParserListener) ExitAlterTypeRenameList(ctx *AlterTypeRenameListContext) {}

// EnterAlterTypeRenameItem is called when production alterTypeRenameItem is entered.
func (s *BaseCqlParserListener) EnterAlterTypeRenameItem(ctx *AlterTypeRenameItemContext) {}

// ExitAlterTypeRenameItem is called when production alterTypeRenameItem is exited.
func (s *BaseCqlParserListener) ExitAlterTypeRenameItem(ctx *AlterTypeRenameItemContext) {}

// EnterAlterTypeAdd is called when production alterTypeAdd is entered.
func (s *BaseCqlParserListener) EnterAlterTypeAdd(ctx *AlterTypeAddContext) {}

// ExitAlterTypeAdd is called when production alterTypeAdd is exited.
func (s *BaseCqlParserListener) ExitAlterTypeAdd(ctx *AlterTypeAddContext) {}

// EnterAlterTypeAlterType is called when production alterTypeAlterType is entered.
func (s *BaseCqlParserListener) EnterAlterTypeAlterType(ctx *AlterTypeAlterTypeContext) {}

// ExitAlterTypeAlterType is called when production alterTypeAlterType is exited.
func (s *BaseCqlParserListener) ExitAlterTypeAlterType(ctx *AlterTypeAlterTypeContext) {}

// EnterAlterTable is called when production alterTable is entered.
func (s *BaseCqlParserListener) EnterAlterTable(ctx *AlterTableContext) {}

// ExitAlterTable is called when production alterTable is exited.
func (s *BaseCqlParserListener) ExitAlterTable(ctx *AlterTableContext) {}

// EnterAlterTableOperation is called when production alterTableOperation is entered.
func (s *BaseCqlParserListener) EnterAlterTableOperation(ctx *AlterTableOperationContext) {}

// ExitAlterTableOperation is called when production alterTableOperation is exited.
func (s *BaseCqlParserListener) ExitAlterTableOperation(ctx *AlterTableOperationContext) {}

// EnterAlterTableWith is called when production alterTableWith is entered.
func (s *BaseCqlParserListener) EnterAlterTableWith(ctx *AlterTableWithContext) {}

// ExitAlterTableWith is called when production alterTableWith is exited.
func (s *BaseCqlParserListener) ExitAlterTableWith(ctx *AlterTableWithContext) {}

// EnterAlterTableRename is called when production alterTableRename is entered.
func (s *BaseCqlParserListener) EnterAlterTableRename(ctx *AlterTableRenameContext) {}

// ExitAlterTableRename is called when production alterTableRename is exited.
func (s *BaseCqlParserListener) ExitAlterTableRename(ctx *AlterTableRenameContext) {}

// EnterAlterTableDropCompactStorage is called when production alterTableDropCompactStorage is entered.
func (s *BaseCqlParserListener) EnterAlterTableDropCompactStorage(ctx *AlterTableDropCompactStorageContext) {
}

// ExitAlterTableDropCompactStorage is called when production alterTableDropCompactStorage is exited.
func (s *BaseCqlParserListener) ExitAlterTableDropCompactStorage(ctx *AlterTableDropCompactStorageContext) {
}

// EnterAlterTableDropColumns is called when production alterTableDropColumns is entered.
func (s *BaseCqlParserListener) EnterAlterTableDropColumns(ctx *AlterTableDropColumnsContext) {}

// ExitAlterTableDropColumns is called when production alterTableDropColumns is exited.
func (s *BaseCqlParserListener) ExitAlterTableDropColumns(ctx *AlterTableDropColumnsContext) {}

// EnterAlterTableDropColumnList is called when production alterTableDropColumnList is entered.
func (s *BaseCqlParserListener) EnterAlterTableDropColumnList(ctx *AlterTableDropColumnListContext) {}

// ExitAlterTableDropColumnList is called when production alterTableDropColumnList is exited.
func (s *BaseCqlParserListener) ExitAlterTableDropColumnList(ctx *AlterTableDropColumnListContext) {}

// EnterAlterTableAdd is called when production alterTableAdd is entered.
func (s *BaseCqlParserListener) EnterAlterTableAdd(ctx *AlterTableAddContext) {}

// ExitAlterTableAdd is called when production alterTableAdd is exited.
func (s *BaseCqlParserListener) ExitAlterTableAdd(ctx *AlterTableAddContext) {}

// EnterAlterTableColumnDefinition is called when production alterTableColumnDefinition is entered.
func (s *BaseCqlParserListener) EnterAlterTableColumnDefinition(ctx *AlterTableColumnDefinitionContext) {
}

// ExitAlterTableColumnDefinition is called when production alterTableColumnDefinition is exited.
func (s *BaseCqlParserListener) ExitAlterTableColumnDefinition(ctx *AlterTableColumnDefinitionContext) {
}

// EnterAlterRole is called when production alterRole is entered.
func (s *BaseCqlParserListener) EnterAlterRole(ctx *AlterRoleContext) {}

// ExitAlterRole is called when production alterRole is exited.
func (s *BaseCqlParserListener) ExitAlterRole(ctx *AlterRoleContext) {}

// EnterRoleWith is called when production roleWith is entered.
func (s *BaseCqlParserListener) EnterRoleWith(ctx *RoleWithContext) {}

// ExitRoleWith is called when production roleWith is exited.
func (s *BaseCqlParserListener) ExitRoleWith(ctx *RoleWithContext) {}

// EnterRoleWithOptions is called when production roleWithOptions is entered.
func (s *BaseCqlParserListener) EnterRoleWithOptions(ctx *RoleWithOptionsContext) {}

// ExitRoleWithOptions is called when production roleWithOptions is exited.
func (s *BaseCqlParserListener) ExitRoleWithOptions(ctx *RoleWithOptionsContext) {}

// EnterAlterMaterializedView is called when production alterMaterializedView is entered.
func (s *BaseCqlParserListener) EnterAlterMaterializedView(ctx *AlterMaterializedViewContext) {}

// ExitAlterMaterializedView is called when production alterMaterializedView is exited.
func (s *BaseCqlParserListener) ExitAlterMaterializedView(ctx *AlterMaterializedViewContext) {}

// EnterDropUser is called when production dropUser is entered.
func (s *BaseCqlParserListener) EnterDropUser(ctx *DropUserContext) {}

// ExitDropUser is called when production dropUser is exited.
func (s *BaseCqlParserListener) ExitDropUser(ctx *DropUserContext) {}

// EnterDropType is called when production dropType is entered.
func (s *BaseCqlParserListener) EnterDropType(ctx *DropTypeContext) {}

// ExitDropType is called when production dropType is exited.
func (s *BaseCqlParserListener) ExitDropType(ctx *DropTypeContext) {}

// EnterDropMaterializedView is called when production dropMaterializedView is entered.
func (s *BaseCqlParserListener) EnterDropMaterializedView(ctx *DropMaterializedViewContext) {}

// ExitDropMaterializedView is called when production dropMaterializedView is exited.
func (s *BaseCqlParserListener) ExitDropMaterializedView(ctx *DropMaterializedViewContext) {}

// EnterDropAggregate is called when production dropAggregate is entered.
func (s *BaseCqlParserListener) EnterDropAggregate(ctx *DropAggregateContext) {}

// ExitDropAggregate is called when production dropAggregate is exited.
func (s *BaseCqlParserListener) ExitDropAggregate(ctx *DropAggregateContext) {}

// EnterDropFunction is called when production dropFunction is entered.
func (s *BaseCqlParserListener) EnterDropFunction(ctx *DropFunctionContext) {}

// ExitDropFunction is called when production dropFunction is exited.
func (s *BaseCqlParserListener) ExitDropFunction(ctx *DropFunctionContext) {}

// EnterDropTrigger is called when production dropTrigger is entered.
func (s *BaseCqlParserListener) EnterDropTrigger(ctx *DropTriggerContext) {}

// ExitDropTrigger is called when production dropTrigger is exited.
func (s *BaseCqlParserListener) ExitDropTrigger(ctx *DropTriggerContext) {}

// EnterDropRole is called when production dropRole is entered.
func (s *BaseCqlParserListener) EnterDropRole(ctx *DropRoleContext) {}

// ExitDropRole is called when production dropRole is exited.
func (s *BaseCqlParserListener) ExitDropRole(ctx *DropRoleContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseCqlParserListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseCqlParserListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropKeyspace is called when production dropKeyspace is entered.
func (s *BaseCqlParserListener) EnterDropKeyspace(ctx *DropKeyspaceContext) {}

// ExitDropKeyspace is called when production dropKeyspace is exited.
func (s *BaseCqlParserListener) ExitDropKeyspace(ctx *DropKeyspaceContext) {}

// EnterDropIndex is called when production dropIndex is entered.
func (s *BaseCqlParserListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production dropIndex is exited.
func (s *BaseCqlParserListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseCqlParserListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseCqlParserListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterWithElement is called when production withElement is entered.
func (s *BaseCqlParserListener) EnterWithElement(ctx *WithElementContext) {}

// ExitWithElement is called when production withElement is exited.
func (s *BaseCqlParserListener) ExitWithElement(ctx *WithElementContext) {}

// EnterClusteringOrder is called when production clusteringOrder is entered.
func (s *BaseCqlParserListener) EnterClusteringOrder(ctx *ClusteringOrderContext) {}

// ExitClusteringOrder is called when production clusteringOrder is exited.
func (s *BaseCqlParserListener) ExitClusteringOrder(ctx *ClusteringOrderContext) {}

// EnterTableOptions is called when production tableOptions is entered.
func (s *BaseCqlParserListener) EnterTableOptions(ctx *TableOptionsContext) {}

// ExitTableOptions is called when production tableOptions is exited.
func (s *BaseCqlParserListener) ExitTableOptions(ctx *TableOptionsContext) {}

// EnterTableOptionItem is called when production tableOptionItem is entered.
func (s *BaseCqlParserListener) EnterTableOptionItem(ctx *TableOptionItemContext) {}

// ExitTableOptionItem is called when production tableOptionItem is exited.
func (s *BaseCqlParserListener) ExitTableOptionItem(ctx *TableOptionItemContext) {}

// EnterTableOptionName is called when production tableOptionName is entered.
func (s *BaseCqlParserListener) EnterTableOptionName(ctx *TableOptionNameContext) {}

// ExitTableOptionName is called when production tableOptionName is exited.
func (s *BaseCqlParserListener) ExitTableOptionName(ctx *TableOptionNameContext) {}

// EnterTableOptionValue is called when production tableOptionValue is entered.
func (s *BaseCqlParserListener) EnterTableOptionValue(ctx *TableOptionValueContext) {}

// ExitTableOptionValue is called when production tableOptionValue is exited.
func (s *BaseCqlParserListener) ExitTableOptionValue(ctx *TableOptionValueContext) {}

// EnterOptionHash is called when production optionHash is entered.
func (s *BaseCqlParserListener) EnterOptionHash(ctx *OptionHashContext) {}

// ExitOptionHash is called when production optionHash is exited.
func (s *BaseCqlParserListener) ExitOptionHash(ctx *OptionHashContext) {}

// EnterOptionHashItem is called when production optionHashItem is entered.
func (s *BaseCqlParserListener) EnterOptionHashItem(ctx *OptionHashItemContext) {}

// ExitOptionHashItem is called when production optionHashItem is exited.
func (s *BaseCqlParserListener) ExitOptionHashItem(ctx *OptionHashItemContext) {}

// EnterOptionHashKey is called when production optionHashKey is entered.
func (s *BaseCqlParserListener) EnterOptionHashKey(ctx *OptionHashKeyContext) {}

// ExitOptionHashKey is called when production optionHashKey is exited.
func (s *BaseCqlParserListener) ExitOptionHashKey(ctx *OptionHashKeyContext) {}

// EnterOptionHashValue is called when production optionHashValue is entered.
func (s *BaseCqlParserListener) EnterOptionHashValue(ctx *OptionHashValueContext) {}

// ExitOptionHashValue is called when production optionHashValue is exited.
func (s *BaseCqlParserListener) ExitOptionHashValue(ctx *OptionHashValueContext) {}

// EnterColumnDefinitionList is called when production columnDefinitionList is entered.
func (s *BaseCqlParserListener) EnterColumnDefinitionList(ctx *ColumnDefinitionListContext) {}

// ExitColumnDefinitionList is called when production columnDefinitionList is exited.
func (s *BaseCqlParserListener) ExitColumnDefinitionList(ctx *ColumnDefinitionListContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseCqlParserListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseCqlParserListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterPrimaryKeyColumn is called when production primaryKeyColumn is entered.
func (s *BaseCqlParserListener) EnterPrimaryKeyColumn(ctx *PrimaryKeyColumnContext) {}

// ExitPrimaryKeyColumn is called when production primaryKeyColumn is exited.
func (s *BaseCqlParserListener) ExitPrimaryKeyColumn(ctx *PrimaryKeyColumnContext) {}

// EnterPrimaryKeyElement is called when production primaryKeyElement is entered.
func (s *BaseCqlParserListener) EnterPrimaryKeyElement(ctx *PrimaryKeyElementContext) {}

// ExitPrimaryKeyElement is called when production primaryKeyElement is exited.
func (s *BaseCqlParserListener) ExitPrimaryKeyElement(ctx *PrimaryKeyElementContext) {}

// EnterPrimaryKeyDefinition is called when production primaryKeyDefinition is entered.
func (s *BaseCqlParserListener) EnterPrimaryKeyDefinition(ctx *PrimaryKeyDefinitionContext) {}

// ExitPrimaryKeyDefinition is called when production primaryKeyDefinition is exited.
func (s *BaseCqlParserListener) ExitPrimaryKeyDefinition(ctx *PrimaryKeyDefinitionContext) {}

// EnterSinglePrimaryKey is called when production singlePrimaryKey is entered.
func (s *BaseCqlParserListener) EnterSinglePrimaryKey(ctx *SinglePrimaryKeyContext) {}

// ExitSinglePrimaryKey is called when production singlePrimaryKey is exited.
func (s *BaseCqlParserListener) ExitSinglePrimaryKey(ctx *SinglePrimaryKeyContext) {}

// EnterCompoundKey is called when production compoundKey is entered.
func (s *BaseCqlParserListener) EnterCompoundKey(ctx *CompoundKeyContext) {}

// ExitCompoundKey is called when production compoundKey is exited.
func (s *BaseCqlParserListener) ExitCompoundKey(ctx *CompoundKeyContext) {}

// EnterCompositeKey is called when production compositeKey is entered.
func (s *BaseCqlParserListener) EnterCompositeKey(ctx *CompositeKeyContext) {}

// ExitCompositeKey is called when production compositeKey is exited.
func (s *BaseCqlParserListener) ExitCompositeKey(ctx *CompositeKeyContext) {}

// EnterPartitionKeyList is called when production partitionKeyList is entered.
func (s *BaseCqlParserListener) EnterPartitionKeyList(ctx *PartitionKeyListContext) {}

// ExitPartitionKeyList is called when production partitionKeyList is exited.
func (s *BaseCqlParserListener) ExitPartitionKeyList(ctx *PartitionKeyListContext) {}

// EnterClusteringKeyList is called when production clusteringKeyList is entered.
func (s *BaseCqlParserListener) EnterClusteringKeyList(ctx *ClusteringKeyListContext) {}

// ExitClusteringKeyList is called when production clusteringKeyList is exited.
func (s *BaseCqlParserListener) ExitClusteringKeyList(ctx *ClusteringKeyListContext) {}

// EnterPartitionKey is called when production partitionKey is entered.
func (s *BaseCqlParserListener) EnterPartitionKey(ctx *PartitionKeyContext) {}

// ExitPartitionKey is called when production partitionKey is exited.
func (s *BaseCqlParserListener) ExitPartitionKey(ctx *PartitionKeyContext) {}

// EnterClusteringKey is called when production clusteringKey is entered.
func (s *BaseCqlParserListener) EnterClusteringKey(ctx *ClusteringKeyContext) {}

// ExitClusteringKey is called when production clusteringKey is exited.
func (s *BaseCqlParserListener) ExitClusteringKey(ctx *ClusteringKeyContext) {}

// EnterApplyBatch is called when production applyBatch is entered.
func (s *BaseCqlParserListener) EnterApplyBatch(ctx *ApplyBatchContext) {}

// ExitApplyBatch is called when production applyBatch is exited.
func (s *BaseCqlParserListener) ExitApplyBatch(ctx *ApplyBatchContext) {}

// EnterBeginBatch is called when production beginBatch is entered.
func (s *BaseCqlParserListener) EnterBeginBatch(ctx *BeginBatchContext) {}

// ExitBeginBatch is called when production beginBatch is exited.
func (s *BaseCqlParserListener) ExitBeginBatch(ctx *BeginBatchContext) {}

// EnterBatchType is called when production batchType is entered.
func (s *BaseCqlParserListener) EnterBatchType(ctx *BatchTypeContext) {}

// ExitBatchType is called when production batchType is exited.
func (s *BaseCqlParserListener) ExitBatchType(ctx *BatchTypeContext) {}

// EnterAlterKeyspace is called when production alterKeyspace is entered.
func (s *BaseCqlParserListener) EnterAlterKeyspace(ctx *AlterKeyspaceContext) {}

// ExitAlterKeyspace is called when production alterKeyspace is exited.
func (s *BaseCqlParserListener) ExitAlterKeyspace(ctx *AlterKeyspaceContext) {}

// EnterReplicationList is called when production replicationList is entered.
func (s *BaseCqlParserListener) EnterReplicationList(ctx *ReplicationListContext) {}

// ExitReplicationList is called when production replicationList is exited.
func (s *BaseCqlParserListener) ExitReplicationList(ctx *ReplicationListContext) {}

// EnterReplicationListItem is called when production replicationListItem is entered.
func (s *BaseCqlParserListener) EnterReplicationListItem(ctx *ReplicationListItemContext) {}

// ExitReplicationListItem is called when production replicationListItem is exited.
func (s *BaseCqlParserListener) ExitReplicationListItem(ctx *ReplicationListItemContext) {}

// EnterDurableWrites is called when production durableWrites is entered.
func (s *BaseCqlParserListener) EnterDurableWrites(ctx *DurableWritesContext) {}

// ExitDurableWrites is called when production durableWrites is exited.
func (s *BaseCqlParserListener) ExitDurableWrites(ctx *DurableWritesContext) {}

// EnterUse_ is called when production use_ is entered.
func (s *BaseCqlParserListener) EnterUse_(ctx *Use_Context) {}

// ExitUse_ is called when production use_ is exited.
func (s *BaseCqlParserListener) ExitUse_(ctx *Use_Context) {}

// EnterTruncate is called when production truncate is entered.
func (s *BaseCqlParserListener) EnterTruncate(ctx *TruncateContext) {}

// ExitTruncate is called when production truncate is exited.
func (s *BaseCqlParserListener) ExitTruncate(ctx *TruncateContext) {}

// EnterCreateIndex is called when production createIndex is entered.
func (s *BaseCqlParserListener) EnterCreateIndex(ctx *CreateIndexContext) {}

// ExitCreateIndex is called when production createIndex is exited.
func (s *BaseCqlParserListener) ExitCreateIndex(ctx *CreateIndexContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseCqlParserListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseCqlParserListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterIndexColumnSpec is called when production indexColumnSpec is entered.
func (s *BaseCqlParserListener) EnterIndexColumnSpec(ctx *IndexColumnSpecContext) {}

// ExitIndexColumnSpec is called when production indexColumnSpec is exited.
func (s *BaseCqlParserListener) ExitIndexColumnSpec(ctx *IndexColumnSpecContext) {}

// EnterIndexKeysSpec is called when production indexKeysSpec is entered.
func (s *BaseCqlParserListener) EnterIndexKeysSpec(ctx *IndexKeysSpecContext) {}

// ExitIndexKeysSpec is called when production indexKeysSpec is exited.
func (s *BaseCqlParserListener) ExitIndexKeysSpec(ctx *IndexKeysSpecContext) {}

// EnterIndexEntriesSSpec is called when production indexEntriesSSpec is entered.
func (s *BaseCqlParserListener) EnterIndexEntriesSSpec(ctx *IndexEntriesSSpecContext) {}

// ExitIndexEntriesSSpec is called when production indexEntriesSSpec is exited.
func (s *BaseCqlParserListener) ExitIndexEntriesSSpec(ctx *IndexEntriesSSpecContext) {}

// EnterIndexFullSpec is called when production indexFullSpec is entered.
func (s *BaseCqlParserListener) EnterIndexFullSpec(ctx *IndexFullSpecContext) {}

// ExitIndexFullSpec is called when production indexFullSpec is exited.
func (s *BaseCqlParserListener) ExitIndexFullSpec(ctx *IndexFullSpecContext) {}

// EnterDelete_ is called when production delete_ is entered.
func (s *BaseCqlParserListener) EnterDelete_(ctx *Delete_Context) {}

// ExitDelete_ is called when production delete_ is exited.
func (s *BaseCqlParserListener) ExitDelete_(ctx *Delete_Context) {}

// EnterDeleteColumnList is called when production deleteColumnList is entered.
func (s *BaseCqlParserListener) EnterDeleteColumnList(ctx *DeleteColumnListContext) {}

// ExitDeleteColumnList is called when production deleteColumnList is exited.
func (s *BaseCqlParserListener) ExitDeleteColumnList(ctx *DeleteColumnListContext) {}

// EnterDeleteColumnItem is called when production deleteColumnItem is entered.
func (s *BaseCqlParserListener) EnterDeleteColumnItem(ctx *DeleteColumnItemContext) {}

// ExitDeleteColumnItem is called when production deleteColumnItem is exited.
func (s *BaseCqlParserListener) ExitDeleteColumnItem(ctx *DeleteColumnItemContext) {}

// EnterUpdate is called when production update is entered.
func (s *BaseCqlParserListener) EnterUpdate(ctx *UpdateContext) {}

// ExitUpdate is called when production update is exited.
func (s *BaseCqlParserListener) ExitUpdate(ctx *UpdateContext) {}

// EnterIfSpec is called when production ifSpec is entered.
func (s *BaseCqlParserListener) EnterIfSpec(ctx *IfSpecContext) {}

// ExitIfSpec is called when production ifSpec is exited.
func (s *BaseCqlParserListener) ExitIfSpec(ctx *IfSpecContext) {}

// EnterIfConditionList is called when production ifConditionList is entered.
func (s *BaseCqlParserListener) EnterIfConditionList(ctx *IfConditionListContext) {}

// ExitIfConditionList is called when production ifConditionList is exited.
func (s *BaseCqlParserListener) ExitIfConditionList(ctx *IfConditionListContext) {}

// EnterIfCondition is called when production ifCondition is entered.
func (s *BaseCqlParserListener) EnterIfCondition(ctx *IfConditionContext) {}

// ExitIfCondition is called when production ifCondition is exited.
func (s *BaseCqlParserListener) ExitIfCondition(ctx *IfConditionContext) {}

// EnterAssignments is called when production assignments is entered.
func (s *BaseCqlParserListener) EnterAssignments(ctx *AssignmentsContext) {}

// ExitAssignments is called when production assignments is exited.
func (s *BaseCqlParserListener) ExitAssignments(ctx *AssignmentsContext) {}

// EnterAssignmentElement is called when production assignmentElement is entered.
func (s *BaseCqlParserListener) EnterAssignmentElement(ctx *AssignmentElementContext) {}

// ExitAssignmentElement is called when production assignmentElement is exited.
func (s *BaseCqlParserListener) ExitAssignmentElement(ctx *AssignmentElementContext) {}

// EnterAssignmentSet is called when production assignmentSet is entered.
func (s *BaseCqlParserListener) EnterAssignmentSet(ctx *AssignmentSetContext) {}

// ExitAssignmentSet is called when production assignmentSet is exited.
func (s *BaseCqlParserListener) ExitAssignmentSet(ctx *AssignmentSetContext) {}

// EnterAssignmentMap is called when production assignmentMap is entered.
func (s *BaseCqlParserListener) EnterAssignmentMap(ctx *AssignmentMapContext) {}

// ExitAssignmentMap is called when production assignmentMap is exited.
func (s *BaseCqlParserListener) ExitAssignmentMap(ctx *AssignmentMapContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseCqlParserListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseCqlParserListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterAssignmentTuple is called when production assignmentTuple is entered.
func (s *BaseCqlParserListener) EnterAssignmentTuple(ctx *AssignmentTupleContext) {}

// ExitAssignmentTuple is called when production assignmentTuple is exited.
func (s *BaseCqlParserListener) ExitAssignmentTuple(ctx *AssignmentTupleContext) {}

// EnterInsert is called when production insert is entered.
func (s *BaseCqlParserListener) EnterInsert(ctx *InsertContext) {}

// ExitInsert is called when production insert is exited.
func (s *BaseCqlParserListener) ExitInsert(ctx *InsertContext) {}

// EnterUsingTtlTimestamp is called when production usingTtlTimestamp is entered.
func (s *BaseCqlParserListener) EnterUsingTtlTimestamp(ctx *UsingTtlTimestampContext) {}

// ExitUsingTtlTimestamp is called when production usingTtlTimestamp is exited.
func (s *BaseCqlParserListener) ExitUsingTtlTimestamp(ctx *UsingTtlTimestampContext) {}

// EnterTimestamp is called when production timestamp is entered.
func (s *BaseCqlParserListener) EnterTimestamp(ctx *TimestampContext) {}

// ExitTimestamp is called when production timestamp is exited.
func (s *BaseCqlParserListener) ExitTimestamp(ctx *TimestampContext) {}

// EnterTtl is called when production ttl is entered.
func (s *BaseCqlParserListener) EnterTtl(ctx *TtlContext) {}

// ExitTtl is called when production ttl is exited.
func (s *BaseCqlParserListener) ExitTtl(ctx *TtlContext) {}

// EnterUsingTimestampSpec is called when production usingTimestampSpec is entered.
func (s *BaseCqlParserListener) EnterUsingTimestampSpec(ctx *UsingTimestampSpecContext) {}

// ExitUsingTimestampSpec is called when production usingTimestampSpec is exited.
func (s *BaseCqlParserListener) ExitUsingTimestampSpec(ctx *UsingTimestampSpecContext) {}

// EnterIfNotExist is called when production ifNotExist is entered.
func (s *BaseCqlParserListener) EnterIfNotExist(ctx *IfNotExistContext) {}

// ExitIfNotExist is called when production ifNotExist is exited.
func (s *BaseCqlParserListener) ExitIfNotExist(ctx *IfNotExistContext) {}

// EnterIfExist is called when production ifExist is entered.
func (s *BaseCqlParserListener) EnterIfExist(ctx *IfExistContext) {}

// ExitIfExist is called when production ifExist is exited.
func (s *BaseCqlParserListener) ExitIfExist(ctx *IfExistContext) {}

// EnterInsertValuesSpec is called when production insertValuesSpec is entered.
func (s *BaseCqlParserListener) EnterInsertValuesSpec(ctx *InsertValuesSpecContext) {}

// ExitInsertValuesSpec is called when production insertValuesSpec is exited.
func (s *BaseCqlParserListener) ExitInsertValuesSpec(ctx *InsertValuesSpecContext) {}

// EnterInsertColumnSpec is called when production insertColumnSpec is entered.
func (s *BaseCqlParserListener) EnterInsertColumnSpec(ctx *InsertColumnSpecContext) {}

// ExitInsertColumnSpec is called when production insertColumnSpec is exited.
func (s *BaseCqlParserListener) ExitInsertColumnSpec(ctx *InsertColumnSpecContext) {}

// EnterColumnList is called when production columnList is entered.
func (s *BaseCqlParserListener) EnterColumnList(ctx *ColumnListContext) {}

// ExitColumnList is called when production columnList is exited.
func (s *BaseCqlParserListener) ExitColumnList(ctx *ColumnListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseCqlParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseCqlParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCqlParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCqlParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSelect_ is called when production select_ is entered.
func (s *BaseCqlParserListener) EnterSelect_(ctx *Select_Context) {}

// ExitSelect_ is called when production select_ is exited.
func (s *BaseCqlParserListener) ExitSelect_(ctx *Select_Context) {}

// EnterAllowFilteringSpec is called when production allowFilteringSpec is entered.
func (s *BaseCqlParserListener) EnterAllowFilteringSpec(ctx *AllowFilteringSpecContext) {}

// ExitAllowFilteringSpec is called when production allowFilteringSpec is exited.
func (s *BaseCqlParserListener) ExitAllowFilteringSpec(ctx *AllowFilteringSpecContext) {}

// EnterLimitSpec is called when production limitSpec is entered.
func (s *BaseCqlParserListener) EnterLimitSpec(ctx *LimitSpecContext) {}

// ExitLimitSpec is called when production limitSpec is exited.
func (s *BaseCqlParserListener) ExitLimitSpec(ctx *LimitSpecContext) {}

// EnterFromSpec is called when production fromSpec is entered.
func (s *BaseCqlParserListener) EnterFromSpec(ctx *FromSpecContext) {}

// ExitFromSpec is called when production fromSpec is exited.
func (s *BaseCqlParserListener) ExitFromSpec(ctx *FromSpecContext) {}

// EnterFromSpecElement is called when production fromSpecElement is entered.
func (s *BaseCqlParserListener) EnterFromSpecElement(ctx *FromSpecElementContext) {}

// ExitFromSpecElement is called when production fromSpecElement is exited.
func (s *BaseCqlParserListener) ExitFromSpecElement(ctx *FromSpecElementContext) {}

// EnterOrderSpec is called when production orderSpec is entered.
func (s *BaseCqlParserListener) EnterOrderSpec(ctx *OrderSpecContext) {}

// ExitOrderSpec is called when production orderSpec is exited.
func (s *BaseCqlParserListener) ExitOrderSpec(ctx *OrderSpecContext) {}

// EnterOrderSpecElement is called when production orderSpecElement is entered.
func (s *BaseCqlParserListener) EnterOrderSpecElement(ctx *OrderSpecElementContext) {}

// ExitOrderSpecElement is called when production orderSpecElement is exited.
func (s *BaseCqlParserListener) ExitOrderSpecElement(ctx *OrderSpecElementContext) {}

// EnterWhereSpec is called when production whereSpec is entered.
func (s *BaseCqlParserListener) EnterWhereSpec(ctx *WhereSpecContext) {}

// ExitWhereSpec is called when production whereSpec is exited.
func (s *BaseCqlParserListener) ExitWhereSpec(ctx *WhereSpecContext) {}

// EnterDistinctSpec is called when production distinctSpec is entered.
func (s *BaseCqlParserListener) EnterDistinctSpec(ctx *DistinctSpecContext) {}

// ExitDistinctSpec is called when production distinctSpec is exited.
func (s *BaseCqlParserListener) ExitDistinctSpec(ctx *DistinctSpecContext) {}

// EnterSelectElements is called when production selectElements is entered.
func (s *BaseCqlParserListener) EnterSelectElements(ctx *SelectElementsContext) {}

// ExitSelectElements is called when production selectElements is exited.
func (s *BaseCqlParserListener) ExitSelectElements(ctx *SelectElementsContext) {}

// EnterSelectElement is called when production selectElement is entered.
func (s *BaseCqlParserListener) EnterSelectElement(ctx *SelectElementContext) {}

// ExitSelectElement is called when production selectElement is exited.
func (s *BaseCqlParserListener) ExitSelectElement(ctx *SelectElementContext) {}

// EnterRelationElements is called when production relationElements is entered.
func (s *BaseCqlParserListener) EnterRelationElements(ctx *RelationElementsContext) {}

// ExitRelationElements is called when production relationElements is exited.
func (s *BaseCqlParserListener) ExitRelationElements(ctx *RelationElementsContext) {}

// EnterRelationElement is called when production relationElement is entered.
func (s *BaseCqlParserListener) EnterRelationElement(ctx *RelationElementContext) {}

// ExitRelationElement is called when production relationElement is exited.
func (s *BaseCqlParserListener) ExitRelationElement(ctx *RelationElementContext) {}

// EnterRelalationContains is called when production relalationContains is entered.
func (s *BaseCqlParserListener) EnterRelalationContains(ctx *RelalationContainsContext) {}

// ExitRelalationContains is called when production relalationContains is exited.
func (s *BaseCqlParserListener) ExitRelalationContains(ctx *RelalationContainsContext) {}

// EnterRelalationContainsKey is called when production relalationContainsKey is entered.
func (s *BaseCqlParserListener) EnterRelalationContainsKey(ctx *RelalationContainsKeyContext) {}

// ExitRelalationContainsKey is called when production relalationContainsKey is exited.
func (s *BaseCqlParserListener) ExitRelalationContainsKey(ctx *RelalationContainsKeyContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseCqlParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseCqlParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *BaseCqlParserListener) EnterFunctionArgs(ctx *FunctionArgsContext) {}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *BaseCqlParserListener) ExitFunctionArgs(ctx *FunctionArgsContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCqlParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCqlParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseCqlParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseCqlParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseCqlParserListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseCqlParserListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseCqlParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseCqlParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseCqlParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseCqlParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterHexadecimalLiteral is called when production hexadecimalLiteral is entered.
func (s *BaseCqlParserListener) EnterHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// ExitHexadecimalLiteral is called when production hexadecimalLiteral is exited.
func (s *BaseCqlParserListener) ExitHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// EnterKeyspace is called when production keyspace is entered.
func (s *BaseCqlParserListener) EnterKeyspace(ctx *KeyspaceContext) {}

// ExitKeyspace is called when production keyspace is exited.
func (s *BaseCqlParserListener) ExitKeyspace(ctx *KeyspaceContext) {}

// EnterTable is called when production table is entered.
func (s *BaseCqlParserListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseCqlParserListener) ExitTable(ctx *TableContext) {}

// EnterColumn is called when production column is entered.
func (s *BaseCqlParserListener) EnterColumn(ctx *ColumnContext) {}

// ExitColumn is called when production column is exited.
func (s *BaseCqlParserListener) ExitColumn(ctx *ColumnContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseCqlParserListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseCqlParserListener) ExitDataType(ctx *DataTypeContext) {}

// EnterDataTypeName is called when production dataTypeName is entered.
func (s *BaseCqlParserListener) EnterDataTypeName(ctx *DataTypeNameContext) {}

// ExitDataTypeName is called when production dataTypeName is exited.
func (s *BaseCqlParserListener) ExitDataTypeName(ctx *DataTypeNameContext) {}

// EnterDataTypeDefinition is called when production dataTypeDefinition is entered.
func (s *BaseCqlParserListener) EnterDataTypeDefinition(ctx *DataTypeDefinitionContext) {}

// ExitDataTypeDefinition is called when production dataTypeDefinition is exited.
func (s *BaseCqlParserListener) ExitDataTypeDefinition(ctx *DataTypeDefinitionContext) {}

// EnterOrderDirection is called when production orderDirection is entered.
func (s *BaseCqlParserListener) EnterOrderDirection(ctx *OrderDirectionContext) {}

// ExitOrderDirection is called when production orderDirection is exited.
func (s *BaseCqlParserListener) ExitOrderDirection(ctx *OrderDirectionContext) {}

// EnterRole is called when production role is entered.
func (s *BaseCqlParserListener) EnterRole(ctx *RoleContext) {}

// ExitRole is called when production role is exited.
func (s *BaseCqlParserListener) ExitRole(ctx *RoleContext) {}

// EnterTrigger is called when production trigger is entered.
func (s *BaseCqlParserListener) EnterTrigger(ctx *TriggerContext) {}

// ExitTrigger is called when production trigger is exited.
func (s *BaseCqlParserListener) ExitTrigger(ctx *TriggerContext) {}

// EnterTriggerClass is called when production triggerClass is entered.
func (s *BaseCqlParserListener) EnterTriggerClass(ctx *TriggerClassContext) {}

// ExitTriggerClass is called when production triggerClass is exited.
func (s *BaseCqlParserListener) ExitTriggerClass(ctx *TriggerClassContext) {}

// EnterMaterializedView is called when production materializedView is entered.
func (s *BaseCqlParserListener) EnterMaterializedView(ctx *MaterializedViewContext) {}

// ExitMaterializedView is called when production materializedView is exited.
func (s *BaseCqlParserListener) ExitMaterializedView(ctx *MaterializedViewContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseCqlParserListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseCqlParserListener) ExitType_(ctx *Type_Context) {}

// EnterAggregate is called when production aggregate is entered.
func (s *BaseCqlParserListener) EnterAggregate(ctx *AggregateContext) {}

// ExitAggregate is called when production aggregate is exited.
func (s *BaseCqlParserListener) ExitAggregate(ctx *AggregateContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BaseCqlParserListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BaseCqlParserListener) ExitFunction_(ctx *Function_Context) {}

// EnterLanguage is called when production language is entered.
func (s *BaseCqlParserListener) EnterLanguage(ctx *LanguageContext) {}

// ExitLanguage is called when production language is exited.
func (s *BaseCqlParserListener) ExitLanguage(ctx *LanguageContext) {}

// EnterUser is called when production user is entered.
func (s *BaseCqlParserListener) EnterUser(ctx *UserContext) {}

// ExitUser is called when production user is exited.
func (s *BaseCqlParserListener) ExitUser(ctx *UserContext) {}

// EnterPassword is called when production password is entered.
func (s *BaseCqlParserListener) EnterPassword(ctx *PasswordContext) {}

// ExitPassword is called when production password is exited.
func (s *BaseCqlParserListener) ExitPassword(ctx *PasswordContext) {}

// EnterHashKey is called when production hashKey is entered.
func (s *BaseCqlParserListener) EnterHashKey(ctx *HashKeyContext) {}

// ExitHashKey is called when production hashKey is exited.
func (s *BaseCqlParserListener) ExitHashKey(ctx *HashKeyContext) {}

// EnterParam is called when production param is entered.
func (s *BaseCqlParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseCqlParserListener) ExitParam(ctx *ParamContext) {}

// EnterParamName is called when production paramName is entered.
func (s *BaseCqlParserListener) EnterParamName(ctx *ParamNameContext) {}

// ExitParamName is called when production paramName is exited.
func (s *BaseCqlParserListener) ExitParamName(ctx *ParamNameContext) {}

// EnterKwAdd is called when production kwAdd is entered.
func (s *BaseCqlParserListener) EnterKwAdd(ctx *KwAddContext) {}

// ExitKwAdd is called when production kwAdd is exited.
func (s *BaseCqlParserListener) ExitKwAdd(ctx *KwAddContext) {}

// EnterKwAggregate is called when production kwAggregate is entered.
func (s *BaseCqlParserListener) EnterKwAggregate(ctx *KwAggregateContext) {}

// ExitKwAggregate is called when production kwAggregate is exited.
func (s *BaseCqlParserListener) ExitKwAggregate(ctx *KwAggregateContext) {}

// EnterKwAll is called when production kwAll is entered.
func (s *BaseCqlParserListener) EnterKwAll(ctx *KwAllContext) {}

// ExitKwAll is called when production kwAll is exited.
func (s *BaseCqlParserListener) ExitKwAll(ctx *KwAllContext) {}

// EnterKwAllPermissions is called when production kwAllPermissions is entered.
func (s *BaseCqlParserListener) EnterKwAllPermissions(ctx *KwAllPermissionsContext) {}

// ExitKwAllPermissions is called when production kwAllPermissions is exited.
func (s *BaseCqlParserListener) ExitKwAllPermissions(ctx *KwAllPermissionsContext) {}

// EnterKwAllow is called when production kwAllow is entered.
func (s *BaseCqlParserListener) EnterKwAllow(ctx *KwAllowContext) {}

// ExitKwAllow is called when production kwAllow is exited.
func (s *BaseCqlParserListener) ExitKwAllow(ctx *KwAllowContext) {}

// EnterKwAlter is called when production kwAlter is entered.
func (s *BaseCqlParserListener) EnterKwAlter(ctx *KwAlterContext) {}

// ExitKwAlter is called when production kwAlter is exited.
func (s *BaseCqlParserListener) ExitKwAlter(ctx *KwAlterContext) {}

// EnterKwAnd is called when production kwAnd is entered.
func (s *BaseCqlParserListener) EnterKwAnd(ctx *KwAndContext) {}

// ExitKwAnd is called when production kwAnd is exited.
func (s *BaseCqlParserListener) ExitKwAnd(ctx *KwAndContext) {}

// EnterKwApply is called when production kwApply is entered.
func (s *BaseCqlParserListener) EnterKwApply(ctx *KwApplyContext) {}

// ExitKwApply is called when production kwApply is exited.
func (s *BaseCqlParserListener) ExitKwApply(ctx *KwApplyContext) {}

// EnterKwAs is called when production kwAs is entered.
func (s *BaseCqlParserListener) EnterKwAs(ctx *KwAsContext) {}

// ExitKwAs is called when production kwAs is exited.
func (s *BaseCqlParserListener) ExitKwAs(ctx *KwAsContext) {}

// EnterKwAsc is called when production kwAsc is entered.
func (s *BaseCqlParserListener) EnterKwAsc(ctx *KwAscContext) {}

// ExitKwAsc is called when production kwAsc is exited.
func (s *BaseCqlParserListener) ExitKwAsc(ctx *KwAscContext) {}

// EnterKwAuthorize is called when production kwAuthorize is entered.
func (s *BaseCqlParserListener) EnterKwAuthorize(ctx *KwAuthorizeContext) {}

// ExitKwAuthorize is called when production kwAuthorize is exited.
func (s *BaseCqlParserListener) ExitKwAuthorize(ctx *KwAuthorizeContext) {}

// EnterKwBatch is called when production kwBatch is entered.
func (s *BaseCqlParserListener) EnterKwBatch(ctx *KwBatchContext) {}

// ExitKwBatch is called when production kwBatch is exited.
func (s *BaseCqlParserListener) ExitKwBatch(ctx *KwBatchContext) {}

// EnterKwBegin is called when production kwBegin is entered.
func (s *BaseCqlParserListener) EnterKwBegin(ctx *KwBeginContext) {}

// ExitKwBegin is called when production kwBegin is exited.
func (s *BaseCqlParserListener) ExitKwBegin(ctx *KwBeginContext) {}

// EnterKwBy is called when production kwBy is entered.
func (s *BaseCqlParserListener) EnterKwBy(ctx *KwByContext) {}

// ExitKwBy is called when production kwBy is exited.
func (s *BaseCqlParserListener) ExitKwBy(ctx *KwByContext) {}

// EnterKwCalled is called when production kwCalled is entered.
func (s *BaseCqlParserListener) EnterKwCalled(ctx *KwCalledContext) {}

// ExitKwCalled is called when production kwCalled is exited.
func (s *BaseCqlParserListener) ExitKwCalled(ctx *KwCalledContext) {}

// EnterKwClustering is called when production kwClustering is entered.
func (s *BaseCqlParserListener) EnterKwClustering(ctx *KwClusteringContext) {}

// ExitKwClustering is called when production kwClustering is exited.
func (s *BaseCqlParserListener) ExitKwClustering(ctx *KwClusteringContext) {}

// EnterKwCompact is called when production kwCompact is entered.
func (s *BaseCqlParserListener) EnterKwCompact(ctx *KwCompactContext) {}

// ExitKwCompact is called when production kwCompact is exited.
func (s *BaseCqlParserListener) ExitKwCompact(ctx *KwCompactContext) {}

// EnterKwContains is called when production kwContains is entered.
func (s *BaseCqlParserListener) EnterKwContains(ctx *KwContainsContext) {}

// ExitKwContains is called when production kwContains is exited.
func (s *BaseCqlParserListener) ExitKwContains(ctx *KwContainsContext) {}

// EnterKwCreate is called when production kwCreate is entered.
func (s *BaseCqlParserListener) EnterKwCreate(ctx *KwCreateContext) {}

// ExitKwCreate is called when production kwCreate is exited.
func (s *BaseCqlParserListener) ExitKwCreate(ctx *KwCreateContext) {}

// EnterKwDelete is called when production kwDelete is entered.
func (s *BaseCqlParserListener) EnterKwDelete(ctx *KwDeleteContext) {}

// ExitKwDelete is called when production kwDelete is exited.
func (s *BaseCqlParserListener) ExitKwDelete(ctx *KwDeleteContext) {}

// EnterKwDesc is called when production kwDesc is entered.
func (s *BaseCqlParserListener) EnterKwDesc(ctx *KwDescContext) {}

// ExitKwDesc is called when production kwDesc is exited.
func (s *BaseCqlParserListener) ExitKwDesc(ctx *KwDescContext) {}

// EnterKwDescibe is called when production kwDescibe is entered.
func (s *BaseCqlParserListener) EnterKwDescibe(ctx *KwDescibeContext) {}

// ExitKwDescibe is called when production kwDescibe is exited.
func (s *BaseCqlParserListener) ExitKwDescibe(ctx *KwDescibeContext) {}

// EnterKwDistinct is called when production kwDistinct is entered.
func (s *BaseCqlParserListener) EnterKwDistinct(ctx *KwDistinctContext) {}

// ExitKwDistinct is called when production kwDistinct is exited.
func (s *BaseCqlParserListener) ExitKwDistinct(ctx *KwDistinctContext) {}

// EnterKwDrop is called when production kwDrop is entered.
func (s *BaseCqlParserListener) EnterKwDrop(ctx *KwDropContext) {}

// ExitKwDrop is called when production kwDrop is exited.
func (s *BaseCqlParserListener) ExitKwDrop(ctx *KwDropContext) {}

// EnterKwDurableWrites is called when production kwDurableWrites is entered.
func (s *BaseCqlParserListener) EnterKwDurableWrites(ctx *KwDurableWritesContext) {}

// ExitKwDurableWrites is called when production kwDurableWrites is exited.
func (s *BaseCqlParserListener) ExitKwDurableWrites(ctx *KwDurableWritesContext) {}

// EnterKwEntries is called when production kwEntries is entered.
func (s *BaseCqlParserListener) EnterKwEntries(ctx *KwEntriesContext) {}

// ExitKwEntries is called when production kwEntries is exited.
func (s *BaseCqlParserListener) ExitKwEntries(ctx *KwEntriesContext) {}

// EnterKwExecute is called when production kwExecute is entered.
func (s *BaseCqlParserListener) EnterKwExecute(ctx *KwExecuteContext) {}

// ExitKwExecute is called when production kwExecute is exited.
func (s *BaseCqlParserListener) ExitKwExecute(ctx *KwExecuteContext) {}

// EnterKwExists is called when production kwExists is entered.
func (s *BaseCqlParserListener) EnterKwExists(ctx *KwExistsContext) {}

// ExitKwExists is called when production kwExists is exited.
func (s *BaseCqlParserListener) ExitKwExists(ctx *KwExistsContext) {}

// EnterKwFiltering is called when production kwFiltering is entered.
func (s *BaseCqlParserListener) EnterKwFiltering(ctx *KwFilteringContext) {}

// ExitKwFiltering is called when production kwFiltering is exited.
func (s *BaseCqlParserListener) ExitKwFiltering(ctx *KwFilteringContext) {}

// EnterKwFinalfunc is called when production kwFinalfunc is entered.
func (s *BaseCqlParserListener) EnterKwFinalfunc(ctx *KwFinalfuncContext) {}

// ExitKwFinalfunc is called when production kwFinalfunc is exited.
func (s *BaseCqlParserListener) ExitKwFinalfunc(ctx *KwFinalfuncContext) {}

// EnterKwFrom is called when production kwFrom is entered.
func (s *BaseCqlParserListener) EnterKwFrom(ctx *KwFromContext) {}

// ExitKwFrom is called when production kwFrom is exited.
func (s *BaseCqlParserListener) ExitKwFrom(ctx *KwFromContext) {}

// EnterKwFull is called when production kwFull is entered.
func (s *BaseCqlParserListener) EnterKwFull(ctx *KwFullContext) {}

// ExitKwFull is called when production kwFull is exited.
func (s *BaseCqlParserListener) ExitKwFull(ctx *KwFullContext) {}

// EnterKwFunction is called when production kwFunction is entered.
func (s *BaseCqlParserListener) EnterKwFunction(ctx *KwFunctionContext) {}

// ExitKwFunction is called when production kwFunction is exited.
func (s *BaseCqlParserListener) ExitKwFunction(ctx *KwFunctionContext) {}

// EnterKwFunctions is called when production kwFunctions is entered.
func (s *BaseCqlParserListener) EnterKwFunctions(ctx *KwFunctionsContext) {}

// ExitKwFunctions is called when production kwFunctions is exited.
func (s *BaseCqlParserListener) ExitKwFunctions(ctx *KwFunctionsContext) {}

// EnterKwGrant is called when production kwGrant is entered.
func (s *BaseCqlParserListener) EnterKwGrant(ctx *KwGrantContext) {}

// ExitKwGrant is called when production kwGrant is exited.
func (s *BaseCqlParserListener) ExitKwGrant(ctx *KwGrantContext) {}

// EnterKwIf is called when production kwIf is entered.
func (s *BaseCqlParserListener) EnterKwIf(ctx *KwIfContext) {}

// ExitKwIf is called when production kwIf is exited.
func (s *BaseCqlParserListener) ExitKwIf(ctx *KwIfContext) {}

// EnterKwIn is called when production kwIn is entered.
func (s *BaseCqlParserListener) EnterKwIn(ctx *KwInContext) {}

// ExitKwIn is called when production kwIn is exited.
func (s *BaseCqlParserListener) ExitKwIn(ctx *KwInContext) {}

// EnterKwIndex is called when production kwIndex is entered.
func (s *BaseCqlParserListener) EnterKwIndex(ctx *KwIndexContext) {}

// ExitKwIndex is called when production kwIndex is exited.
func (s *BaseCqlParserListener) ExitKwIndex(ctx *KwIndexContext) {}

// EnterKwInitcond is called when production kwInitcond is entered.
func (s *BaseCqlParserListener) EnterKwInitcond(ctx *KwInitcondContext) {}

// ExitKwInitcond is called when production kwInitcond is exited.
func (s *BaseCqlParserListener) ExitKwInitcond(ctx *KwInitcondContext) {}

// EnterKwInput is called when production kwInput is entered.
func (s *BaseCqlParserListener) EnterKwInput(ctx *KwInputContext) {}

// ExitKwInput is called when production kwInput is exited.
func (s *BaseCqlParserListener) ExitKwInput(ctx *KwInputContext) {}

// EnterKwInsert is called when production kwInsert is entered.
func (s *BaseCqlParserListener) EnterKwInsert(ctx *KwInsertContext) {}

// ExitKwInsert is called when production kwInsert is exited.
func (s *BaseCqlParserListener) ExitKwInsert(ctx *KwInsertContext) {}

// EnterKwInto is called when production kwInto is entered.
func (s *BaseCqlParserListener) EnterKwInto(ctx *KwIntoContext) {}

// ExitKwInto is called when production kwInto is exited.
func (s *BaseCqlParserListener) ExitKwInto(ctx *KwIntoContext) {}

// EnterKwIs is called when production kwIs is entered.
func (s *BaseCqlParserListener) EnterKwIs(ctx *KwIsContext) {}

// ExitKwIs is called when production kwIs is exited.
func (s *BaseCqlParserListener) ExitKwIs(ctx *KwIsContext) {}

// EnterKwJson is called when production kwJson is entered.
func (s *BaseCqlParserListener) EnterKwJson(ctx *KwJsonContext) {}

// ExitKwJson is called when production kwJson is exited.
func (s *BaseCqlParserListener) ExitKwJson(ctx *KwJsonContext) {}

// EnterKwKey is called when production kwKey is entered.
func (s *BaseCqlParserListener) EnterKwKey(ctx *KwKeyContext) {}

// ExitKwKey is called when production kwKey is exited.
func (s *BaseCqlParserListener) ExitKwKey(ctx *KwKeyContext) {}

// EnterKwKeys is called when production kwKeys is entered.
func (s *BaseCqlParserListener) EnterKwKeys(ctx *KwKeysContext) {}

// ExitKwKeys is called when production kwKeys is exited.
func (s *BaseCqlParserListener) ExitKwKeys(ctx *KwKeysContext) {}

// EnterKwKeyspace is called when production kwKeyspace is entered.
func (s *BaseCqlParserListener) EnterKwKeyspace(ctx *KwKeyspaceContext) {}

// ExitKwKeyspace is called when production kwKeyspace is exited.
func (s *BaseCqlParserListener) ExitKwKeyspace(ctx *KwKeyspaceContext) {}

// EnterKwKeyspaces is called when production kwKeyspaces is entered.
func (s *BaseCqlParserListener) EnterKwKeyspaces(ctx *KwKeyspacesContext) {}

// ExitKwKeyspaces is called when production kwKeyspaces is exited.
func (s *BaseCqlParserListener) ExitKwKeyspaces(ctx *KwKeyspacesContext) {}

// EnterKwLanguage is called when production kwLanguage is entered.
func (s *BaseCqlParserListener) EnterKwLanguage(ctx *KwLanguageContext) {}

// ExitKwLanguage is called when production kwLanguage is exited.
func (s *BaseCqlParserListener) ExitKwLanguage(ctx *KwLanguageContext) {}

// EnterKwLimit is called when production kwLimit is entered.
func (s *BaseCqlParserListener) EnterKwLimit(ctx *KwLimitContext) {}

// ExitKwLimit is called when production kwLimit is exited.
func (s *BaseCqlParserListener) ExitKwLimit(ctx *KwLimitContext) {}

// EnterKwList is called when production kwList is entered.
func (s *BaseCqlParserListener) EnterKwList(ctx *KwListContext) {}

// ExitKwList is called when production kwList is exited.
func (s *BaseCqlParserListener) ExitKwList(ctx *KwListContext) {}

// EnterKwLogged is called when production kwLogged is entered.
func (s *BaseCqlParserListener) EnterKwLogged(ctx *KwLoggedContext) {}

// ExitKwLogged is called when production kwLogged is exited.
func (s *BaseCqlParserListener) ExitKwLogged(ctx *KwLoggedContext) {}

// EnterKwLogin is called when production kwLogin is entered.
func (s *BaseCqlParserListener) EnterKwLogin(ctx *KwLoginContext) {}

// ExitKwLogin is called when production kwLogin is exited.
func (s *BaseCqlParserListener) ExitKwLogin(ctx *KwLoginContext) {}

// EnterKwMaterialized is called when production kwMaterialized is entered.
func (s *BaseCqlParserListener) EnterKwMaterialized(ctx *KwMaterializedContext) {}

// ExitKwMaterialized is called when production kwMaterialized is exited.
func (s *BaseCqlParserListener) ExitKwMaterialized(ctx *KwMaterializedContext) {}

// EnterKwModify is called when production kwModify is entered.
func (s *BaseCqlParserListener) EnterKwModify(ctx *KwModifyContext) {}

// ExitKwModify is called when production kwModify is exited.
func (s *BaseCqlParserListener) ExitKwModify(ctx *KwModifyContext) {}

// EnterKwNosuperuser is called when production kwNosuperuser is entered.
func (s *BaseCqlParserListener) EnterKwNosuperuser(ctx *KwNosuperuserContext) {}

// ExitKwNosuperuser is called when production kwNosuperuser is exited.
func (s *BaseCqlParserListener) ExitKwNosuperuser(ctx *KwNosuperuserContext) {}

// EnterKwNorecursive is called when production kwNorecursive is entered.
func (s *BaseCqlParserListener) EnterKwNorecursive(ctx *KwNorecursiveContext) {}

// ExitKwNorecursive is called when production kwNorecursive is exited.
func (s *BaseCqlParserListener) ExitKwNorecursive(ctx *KwNorecursiveContext) {}

// EnterKwNot is called when production kwNot is entered.
func (s *BaseCqlParserListener) EnterKwNot(ctx *KwNotContext) {}

// ExitKwNot is called when production kwNot is exited.
func (s *BaseCqlParserListener) ExitKwNot(ctx *KwNotContext) {}

// EnterKwNull is called when production kwNull is entered.
func (s *BaseCqlParserListener) EnterKwNull(ctx *KwNullContext) {}

// ExitKwNull is called when production kwNull is exited.
func (s *BaseCqlParserListener) ExitKwNull(ctx *KwNullContext) {}

// EnterKwOf is called when production kwOf is entered.
func (s *BaseCqlParserListener) EnterKwOf(ctx *KwOfContext) {}

// ExitKwOf is called when production kwOf is exited.
func (s *BaseCqlParserListener) ExitKwOf(ctx *KwOfContext) {}

// EnterKwOn is called when production kwOn is entered.
func (s *BaseCqlParserListener) EnterKwOn(ctx *KwOnContext) {}

// ExitKwOn is called when production kwOn is exited.
func (s *BaseCqlParserListener) ExitKwOn(ctx *KwOnContext) {}

// EnterKwOptions is called when production kwOptions is entered.
func (s *BaseCqlParserListener) EnterKwOptions(ctx *KwOptionsContext) {}

// ExitKwOptions is called when production kwOptions is exited.
func (s *BaseCqlParserListener) ExitKwOptions(ctx *KwOptionsContext) {}

// EnterKwOr is called when production kwOr is entered.
func (s *BaseCqlParserListener) EnterKwOr(ctx *KwOrContext) {}

// ExitKwOr is called when production kwOr is exited.
func (s *BaseCqlParserListener) ExitKwOr(ctx *KwOrContext) {}

// EnterKwOrder is called when production kwOrder is entered.
func (s *BaseCqlParserListener) EnterKwOrder(ctx *KwOrderContext) {}

// ExitKwOrder is called when production kwOrder is exited.
func (s *BaseCqlParserListener) ExitKwOrder(ctx *KwOrderContext) {}

// EnterKwPassword is called when production kwPassword is entered.
func (s *BaseCqlParserListener) EnterKwPassword(ctx *KwPasswordContext) {}

// ExitKwPassword is called when production kwPassword is exited.
func (s *BaseCqlParserListener) ExitKwPassword(ctx *KwPasswordContext) {}

// EnterKwPrimary is called when production kwPrimary is entered.
func (s *BaseCqlParserListener) EnterKwPrimary(ctx *KwPrimaryContext) {}

// ExitKwPrimary is called when production kwPrimary is exited.
func (s *BaseCqlParserListener) ExitKwPrimary(ctx *KwPrimaryContext) {}

// EnterKwRename is called when production kwRename is entered.
func (s *BaseCqlParserListener) EnterKwRename(ctx *KwRenameContext) {}

// ExitKwRename is called when production kwRename is exited.
func (s *BaseCqlParserListener) ExitKwRename(ctx *KwRenameContext) {}

// EnterKwReplace is called when production kwReplace is entered.
func (s *BaseCqlParserListener) EnterKwReplace(ctx *KwReplaceContext) {}

// ExitKwReplace is called when production kwReplace is exited.
func (s *BaseCqlParserListener) ExitKwReplace(ctx *KwReplaceContext) {}

// EnterKwReplication is called when production kwReplication is entered.
func (s *BaseCqlParserListener) EnterKwReplication(ctx *KwReplicationContext) {}

// ExitKwReplication is called when production kwReplication is exited.
func (s *BaseCqlParserListener) ExitKwReplication(ctx *KwReplicationContext) {}

// EnterKwReturns is called when production kwReturns is entered.
func (s *BaseCqlParserListener) EnterKwReturns(ctx *KwReturnsContext) {}

// ExitKwReturns is called when production kwReturns is exited.
func (s *BaseCqlParserListener) ExitKwReturns(ctx *KwReturnsContext) {}

// EnterKwRole is called when production kwRole is entered.
func (s *BaseCqlParserListener) EnterKwRole(ctx *KwRoleContext) {}

// ExitKwRole is called when production kwRole is exited.
func (s *BaseCqlParserListener) ExitKwRole(ctx *KwRoleContext) {}

// EnterKwRoles is called when production kwRoles is entered.
func (s *BaseCqlParserListener) EnterKwRoles(ctx *KwRolesContext) {}

// ExitKwRoles is called when production kwRoles is exited.
func (s *BaseCqlParserListener) ExitKwRoles(ctx *KwRolesContext) {}

// EnterKwSelect is called when production kwSelect is entered.
func (s *BaseCqlParserListener) EnterKwSelect(ctx *KwSelectContext) {}

// ExitKwSelect is called when production kwSelect is exited.
func (s *BaseCqlParserListener) ExitKwSelect(ctx *KwSelectContext) {}

// EnterKwSet is called when production kwSet is entered.
func (s *BaseCqlParserListener) EnterKwSet(ctx *KwSetContext) {}

// ExitKwSet is called when production kwSet is exited.
func (s *BaseCqlParserListener) ExitKwSet(ctx *KwSetContext) {}

// EnterKwSfunc is called when production kwSfunc is entered.
func (s *BaseCqlParserListener) EnterKwSfunc(ctx *KwSfuncContext) {}

// ExitKwSfunc is called when production kwSfunc is exited.
func (s *BaseCqlParserListener) ExitKwSfunc(ctx *KwSfuncContext) {}

// EnterKwStorage is called when production kwStorage is entered.
func (s *BaseCqlParserListener) EnterKwStorage(ctx *KwStorageContext) {}

// ExitKwStorage is called when production kwStorage is exited.
func (s *BaseCqlParserListener) ExitKwStorage(ctx *KwStorageContext) {}

// EnterKwStype is called when production kwStype is entered.
func (s *BaseCqlParserListener) EnterKwStype(ctx *KwStypeContext) {}

// ExitKwStype is called when production kwStype is exited.
func (s *BaseCqlParserListener) ExitKwStype(ctx *KwStypeContext) {}

// EnterKwSuperuser is called when production kwSuperuser is entered.
func (s *BaseCqlParserListener) EnterKwSuperuser(ctx *KwSuperuserContext) {}

// ExitKwSuperuser is called when production kwSuperuser is exited.
func (s *BaseCqlParserListener) ExitKwSuperuser(ctx *KwSuperuserContext) {}

// EnterKwTable is called when production kwTable is entered.
func (s *BaseCqlParserListener) EnterKwTable(ctx *KwTableContext) {}

// ExitKwTable is called when production kwTable is exited.
func (s *BaseCqlParserListener) ExitKwTable(ctx *KwTableContext) {}

// EnterKwTimestamp is called when production kwTimestamp is entered.
func (s *BaseCqlParserListener) EnterKwTimestamp(ctx *KwTimestampContext) {}

// ExitKwTimestamp is called when production kwTimestamp is exited.
func (s *BaseCqlParserListener) ExitKwTimestamp(ctx *KwTimestampContext) {}

// EnterKwTo is called when production kwTo is entered.
func (s *BaseCqlParserListener) EnterKwTo(ctx *KwToContext) {}

// ExitKwTo is called when production kwTo is exited.
func (s *BaseCqlParserListener) ExitKwTo(ctx *KwToContext) {}

// EnterKwTrigger is called when production kwTrigger is entered.
func (s *BaseCqlParserListener) EnterKwTrigger(ctx *KwTriggerContext) {}

// ExitKwTrigger is called when production kwTrigger is exited.
func (s *BaseCqlParserListener) ExitKwTrigger(ctx *KwTriggerContext) {}

// EnterKwTruncate is called when production kwTruncate is entered.
func (s *BaseCqlParserListener) EnterKwTruncate(ctx *KwTruncateContext) {}

// ExitKwTruncate is called when production kwTruncate is exited.
func (s *BaseCqlParserListener) ExitKwTruncate(ctx *KwTruncateContext) {}

// EnterKwTtl is called when production kwTtl is entered.
func (s *BaseCqlParserListener) EnterKwTtl(ctx *KwTtlContext) {}

// ExitKwTtl is called when production kwTtl is exited.
func (s *BaseCqlParserListener) ExitKwTtl(ctx *KwTtlContext) {}

// EnterKwType is called when production kwType is entered.
func (s *BaseCqlParserListener) EnterKwType(ctx *KwTypeContext) {}

// ExitKwType is called when production kwType is exited.
func (s *BaseCqlParserListener) ExitKwType(ctx *KwTypeContext) {}

// EnterKwUnlogged is called when production kwUnlogged is entered.
func (s *BaseCqlParserListener) EnterKwUnlogged(ctx *KwUnloggedContext) {}

// ExitKwUnlogged is called when production kwUnlogged is exited.
func (s *BaseCqlParserListener) ExitKwUnlogged(ctx *KwUnloggedContext) {}

// EnterKwUpdate is called when production kwUpdate is entered.
func (s *BaseCqlParserListener) EnterKwUpdate(ctx *KwUpdateContext) {}

// ExitKwUpdate is called when production kwUpdate is exited.
func (s *BaseCqlParserListener) ExitKwUpdate(ctx *KwUpdateContext) {}

// EnterKwUse is called when production kwUse is entered.
func (s *BaseCqlParserListener) EnterKwUse(ctx *KwUseContext) {}

// ExitKwUse is called when production kwUse is exited.
func (s *BaseCqlParserListener) ExitKwUse(ctx *KwUseContext) {}

// EnterKwUser is called when production kwUser is entered.
func (s *BaseCqlParserListener) EnterKwUser(ctx *KwUserContext) {}

// ExitKwUser is called when production kwUser is exited.
func (s *BaseCqlParserListener) ExitKwUser(ctx *KwUserContext) {}

// EnterKwUsers is called when production kwUsers is entered.
func (s *BaseCqlParserListener) EnterKwUsers(ctx *KwUsersContext) {}

// ExitKwUsers is called when production kwUsers is exited.
func (s *BaseCqlParserListener) ExitKwUsers(ctx *KwUsersContext) {}

// EnterKwUsing is called when production kwUsing is entered.
func (s *BaseCqlParserListener) EnterKwUsing(ctx *KwUsingContext) {}

// ExitKwUsing is called when production kwUsing is exited.
func (s *BaseCqlParserListener) ExitKwUsing(ctx *KwUsingContext) {}

// EnterKwValues is called when production kwValues is entered.
func (s *BaseCqlParserListener) EnterKwValues(ctx *KwValuesContext) {}

// ExitKwValues is called when production kwValues is exited.
func (s *BaseCqlParserListener) ExitKwValues(ctx *KwValuesContext) {}

// EnterKwView is called when production kwView is entered.
func (s *BaseCqlParserListener) EnterKwView(ctx *KwViewContext) {}

// ExitKwView is called when production kwView is exited.
func (s *BaseCqlParserListener) ExitKwView(ctx *KwViewContext) {}

// EnterKwWhere is called when production kwWhere is entered.
func (s *BaseCqlParserListener) EnterKwWhere(ctx *KwWhereContext) {}

// ExitKwWhere is called when production kwWhere is exited.
func (s *BaseCqlParserListener) ExitKwWhere(ctx *KwWhereContext) {}

// EnterKwWith is called when production kwWith is entered.
func (s *BaseCqlParserListener) EnterKwWith(ctx *KwWithContext) {}

// ExitKwWith is called when production kwWith is exited.
func (s *BaseCqlParserListener) ExitKwWith(ctx *KwWithContext) {}

// EnterKwRevoke is called when production kwRevoke is entered.
func (s *BaseCqlParserListener) EnterKwRevoke(ctx *KwRevokeContext) {}

// ExitKwRevoke is called when production kwRevoke is exited.
func (s *BaseCqlParserListener) ExitKwRevoke(ctx *KwRevokeContext) {}

// EnterEof is called when production eof is entered.
func (s *BaseCqlParserListener) EnterEof(ctx *EofContext) {}

// ExitEof is called when production eof is exited.
func (s *BaseCqlParserListener) ExitEof(ctx *EofContext) {}

// EnterSyntaxBracketLr is called when production syntaxBracketLr is entered.
func (s *BaseCqlParserListener) EnterSyntaxBracketLr(ctx *SyntaxBracketLrContext) {}

// ExitSyntaxBracketLr is called when production syntaxBracketLr is exited.
func (s *BaseCqlParserListener) ExitSyntaxBracketLr(ctx *SyntaxBracketLrContext) {}

// EnterSyntaxBracketRr is called when production syntaxBracketRr is entered.
func (s *BaseCqlParserListener) EnterSyntaxBracketRr(ctx *SyntaxBracketRrContext) {}

// ExitSyntaxBracketRr is called when production syntaxBracketRr is exited.
func (s *BaseCqlParserListener) ExitSyntaxBracketRr(ctx *SyntaxBracketRrContext) {}

// EnterSyntaxBracketLc is called when production syntaxBracketLc is entered.
func (s *BaseCqlParserListener) EnterSyntaxBracketLc(ctx *SyntaxBracketLcContext) {}

// ExitSyntaxBracketLc is called when production syntaxBracketLc is exited.
func (s *BaseCqlParserListener) ExitSyntaxBracketLc(ctx *SyntaxBracketLcContext) {}

// EnterSyntaxBracketRc is called when production syntaxBracketRc is entered.
func (s *BaseCqlParserListener) EnterSyntaxBracketRc(ctx *SyntaxBracketRcContext) {}

// ExitSyntaxBracketRc is called when production syntaxBracketRc is exited.
func (s *BaseCqlParserListener) ExitSyntaxBracketRc(ctx *SyntaxBracketRcContext) {}

// EnterSyntaxBracketLa is called when production syntaxBracketLa is entered.
func (s *BaseCqlParserListener) EnterSyntaxBracketLa(ctx *SyntaxBracketLaContext) {}

// ExitSyntaxBracketLa is called when production syntaxBracketLa is exited.
func (s *BaseCqlParserListener) ExitSyntaxBracketLa(ctx *SyntaxBracketLaContext) {}

// EnterSyntaxBracketRa is called when production syntaxBracketRa is entered.
func (s *BaseCqlParserListener) EnterSyntaxBracketRa(ctx *SyntaxBracketRaContext) {}

// ExitSyntaxBracketRa is called when production syntaxBracketRa is exited.
func (s *BaseCqlParserListener) ExitSyntaxBracketRa(ctx *SyntaxBracketRaContext) {}

// EnterSyntaxBracketLs is called when production syntaxBracketLs is entered.
func (s *BaseCqlParserListener) EnterSyntaxBracketLs(ctx *SyntaxBracketLsContext) {}

// ExitSyntaxBracketLs is called when production syntaxBracketLs is exited.
func (s *BaseCqlParserListener) ExitSyntaxBracketLs(ctx *SyntaxBracketLsContext) {}

// EnterSyntaxBracketRs is called when production syntaxBracketRs is entered.
func (s *BaseCqlParserListener) EnterSyntaxBracketRs(ctx *SyntaxBracketRsContext) {}

// ExitSyntaxBracketRs is called when production syntaxBracketRs is exited.
func (s *BaseCqlParserListener) ExitSyntaxBracketRs(ctx *SyntaxBracketRsContext) {}

// EnterSyntaxComma is called when production syntaxComma is entered.
func (s *BaseCqlParserListener) EnterSyntaxComma(ctx *SyntaxCommaContext) {}

// ExitSyntaxComma is called when production syntaxComma is exited.
func (s *BaseCqlParserListener) ExitSyntaxComma(ctx *SyntaxCommaContext) {}

// EnterSyntaxColon is called when production syntaxColon is entered.
func (s *BaseCqlParserListener) EnterSyntaxColon(ctx *SyntaxColonContext) {}

// ExitSyntaxColon is called when production syntaxColon is exited.
func (s *BaseCqlParserListener) ExitSyntaxColon(ctx *SyntaxColonContext) {}
