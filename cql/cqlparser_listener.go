// Code generated from CqlParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cql // CqlParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CqlParserListener is a complete listener for a parse tree produced by CqlParser.
type CqlParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterCqls is called when entering the cqls production.
	EnterCqls(c *CqlsContext)

	// EnterStatementSeparator is called when entering the statementSeparator production.
	EnterStatementSeparator(c *StatementSeparatorContext)

	// EnterEmpty_ is called when entering the empty_ production.
	EnterEmpty_(c *Empty_Context)

	// EnterCql is called when entering the cql production.
	EnterCql(c *CqlContext)

	// EnterRevoke is called when entering the revoke production.
	EnterRevoke(c *RevokeContext)

	// EnterListUsers is called when entering the listUsers production.
	EnterListUsers(c *ListUsersContext)

	// EnterListRoles is called when entering the listRoles production.
	EnterListRoles(c *ListRolesContext)

	// EnterListPermissions is called when entering the listPermissions production.
	EnterListPermissions(c *ListPermissionsContext)

	// EnterGrant is called when entering the grant production.
	EnterGrant(c *GrantContext)

	// EnterPriviledge is called when entering the priviledge production.
	EnterPriviledge(c *PriviledgeContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterCreateUser is called when entering the createUser production.
	EnterCreateUser(c *CreateUserContext)

	// EnterCreateRole is called when entering the createRole production.
	EnterCreateRole(c *CreateRoleContext)

	// EnterCreateType is called when entering the createType production.
	EnterCreateType(c *CreateTypeContext)

	// EnterTypeMemberColumnList is called when entering the typeMemberColumnList production.
	EnterTypeMemberColumnList(c *TypeMemberColumnListContext)

	// EnterCreateTrigger is called when entering the createTrigger production.
	EnterCreateTrigger(c *CreateTriggerContext)

	// EnterCreateMaterializedView is called when entering the createMaterializedView production.
	EnterCreateMaterializedView(c *CreateMaterializedViewContext)

	// EnterMaterializedViewWhere is called when entering the materializedViewWhere production.
	EnterMaterializedViewWhere(c *MaterializedViewWhereContext)

	// EnterColumnNotNullList is called when entering the columnNotNullList production.
	EnterColumnNotNullList(c *ColumnNotNullListContext)

	// EnterColumnNotNull is called when entering the columnNotNull production.
	EnterColumnNotNull(c *ColumnNotNullContext)

	// EnterMaterializedViewOptions is called when entering the materializedViewOptions production.
	EnterMaterializedViewOptions(c *MaterializedViewOptionsContext)

	// EnterCreateKeyspace is called when entering the createKeyspace production.
	EnterCreateKeyspace(c *CreateKeyspaceContext)

	// EnterCreateFunction is called when entering the createFunction production.
	EnterCreateFunction(c *CreateFunctionContext)

	// EnterCodeBlock is called when entering the codeBlock production.
	EnterCodeBlock(c *CodeBlockContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterReturnMode is called when entering the returnMode production.
	EnterReturnMode(c *ReturnModeContext)

	// EnterCreateAggregate is called when entering the createAggregate production.
	EnterCreateAggregate(c *CreateAggregateContext)

	// EnterInitCondDefinition is called when entering the initCondDefinition production.
	EnterInitCondDefinition(c *InitCondDefinitionContext)

	// EnterInitCondHash is called when entering the initCondHash production.
	EnterInitCondHash(c *InitCondHashContext)

	// EnterInitCondHashItem is called when entering the initCondHashItem production.
	EnterInitCondHashItem(c *InitCondHashItemContext)

	// EnterInitCondListNested is called when entering the initCondListNested production.
	EnterInitCondListNested(c *InitCondListNestedContext)

	// EnterInitCondList is called when entering the initCondList production.
	EnterInitCondList(c *InitCondListContext)

	// EnterOrReplace is called when entering the orReplace production.
	EnterOrReplace(c *OrReplaceContext)

	// EnterAlterUser is called when entering the alterUser production.
	EnterAlterUser(c *AlterUserContext)

	// EnterUserPassword is called when entering the userPassword production.
	EnterUserPassword(c *UserPasswordContext)

	// EnterUserSuperUser is called when entering the userSuperUser production.
	EnterUserSuperUser(c *UserSuperUserContext)

	// EnterAlterType is called when entering the alterType production.
	EnterAlterType(c *AlterTypeContext)

	// EnterAlterTypeOperation is called when entering the alterTypeOperation production.
	EnterAlterTypeOperation(c *AlterTypeOperationContext)

	// EnterAlterTypeRename is called when entering the alterTypeRename production.
	EnterAlterTypeRename(c *AlterTypeRenameContext)

	// EnterAlterTypeRenameList is called when entering the alterTypeRenameList production.
	EnterAlterTypeRenameList(c *AlterTypeRenameListContext)

	// EnterAlterTypeRenameItem is called when entering the alterTypeRenameItem production.
	EnterAlterTypeRenameItem(c *AlterTypeRenameItemContext)

	// EnterAlterTypeAdd is called when entering the alterTypeAdd production.
	EnterAlterTypeAdd(c *AlterTypeAddContext)

	// EnterAlterTypeAlterType is called when entering the alterTypeAlterType production.
	EnterAlterTypeAlterType(c *AlterTypeAlterTypeContext)

	// EnterAlterTable is called when entering the alterTable production.
	EnterAlterTable(c *AlterTableContext)

	// EnterAlterTableOperation is called when entering the alterTableOperation production.
	EnterAlterTableOperation(c *AlterTableOperationContext)

	// EnterAlterTableWith is called when entering the alterTableWith production.
	EnterAlterTableWith(c *AlterTableWithContext)

	// EnterAlterTableRename is called when entering the alterTableRename production.
	EnterAlterTableRename(c *AlterTableRenameContext)

	// EnterAlterTableDropCompactStorage is called when entering the alterTableDropCompactStorage production.
	EnterAlterTableDropCompactStorage(c *AlterTableDropCompactStorageContext)

	// EnterAlterTableDropColumns is called when entering the alterTableDropColumns production.
	EnterAlterTableDropColumns(c *AlterTableDropColumnsContext)

	// EnterAlterTableDropColumnList is called when entering the alterTableDropColumnList production.
	EnterAlterTableDropColumnList(c *AlterTableDropColumnListContext)

	// EnterAlterTableAdd is called when entering the alterTableAdd production.
	EnterAlterTableAdd(c *AlterTableAddContext)

	// EnterAlterTableColumnDefinition is called when entering the alterTableColumnDefinition production.
	EnterAlterTableColumnDefinition(c *AlterTableColumnDefinitionContext)

	// EnterAlterRole is called when entering the alterRole production.
	EnterAlterRole(c *AlterRoleContext)

	// EnterRoleWith is called when entering the roleWith production.
	EnterRoleWith(c *RoleWithContext)

	// EnterRoleWithOptions is called when entering the roleWithOptions production.
	EnterRoleWithOptions(c *RoleWithOptionsContext)

	// EnterAlterMaterializedView is called when entering the alterMaterializedView production.
	EnterAlterMaterializedView(c *AlterMaterializedViewContext)

	// EnterDropUser is called when entering the dropUser production.
	EnterDropUser(c *DropUserContext)

	// EnterDropType is called when entering the dropType production.
	EnterDropType(c *DropTypeContext)

	// EnterDropMaterializedView is called when entering the dropMaterializedView production.
	EnterDropMaterializedView(c *DropMaterializedViewContext)

	// EnterDropAggregate is called when entering the dropAggregate production.
	EnterDropAggregate(c *DropAggregateContext)

	// EnterDropFunction is called when entering the dropFunction production.
	EnterDropFunction(c *DropFunctionContext)

	// EnterDropTrigger is called when entering the dropTrigger production.
	EnterDropTrigger(c *DropTriggerContext)

	// EnterDropRole is called when entering the dropRole production.
	EnterDropRole(c *DropRoleContext)

	// EnterDropTable is called when entering the dropTable production.
	EnterDropTable(c *DropTableContext)

	// EnterDropKeyspace is called when entering the dropKeyspace production.
	EnterDropKeyspace(c *DropKeyspaceContext)

	// EnterDropIndex is called when entering the dropIndex production.
	EnterDropIndex(c *DropIndexContext)

	// EnterCreateTable is called when entering the createTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterWithElement is called when entering the withElement production.
	EnterWithElement(c *WithElementContext)

	// EnterClusteringOrder is called when entering the clusteringOrder production.
	EnterClusteringOrder(c *ClusteringOrderContext)

	// EnterTableOptions is called when entering the tableOptions production.
	EnterTableOptions(c *TableOptionsContext)

	// EnterTableOptionItem is called when entering the tableOptionItem production.
	EnterTableOptionItem(c *TableOptionItemContext)

	// EnterTableOptionName is called when entering the tableOptionName production.
	EnterTableOptionName(c *TableOptionNameContext)

	// EnterTableOptionValue is called when entering the tableOptionValue production.
	EnterTableOptionValue(c *TableOptionValueContext)

	// EnterOptionHash is called when entering the optionHash production.
	EnterOptionHash(c *OptionHashContext)

	// EnterOptionHashItem is called when entering the optionHashItem production.
	EnterOptionHashItem(c *OptionHashItemContext)

	// EnterOptionHashKey is called when entering the optionHashKey production.
	EnterOptionHashKey(c *OptionHashKeyContext)

	// EnterOptionHashValue is called when entering the optionHashValue production.
	EnterOptionHashValue(c *OptionHashValueContext)

	// EnterColumnDefinitionList is called when entering the columnDefinitionList production.
	EnterColumnDefinitionList(c *ColumnDefinitionListContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterPrimaryKeyColumn is called when entering the primaryKeyColumn production.
	EnterPrimaryKeyColumn(c *PrimaryKeyColumnContext)

	// EnterPrimaryKeyElement is called when entering the primaryKeyElement production.
	EnterPrimaryKeyElement(c *PrimaryKeyElementContext)

	// EnterPrimaryKeyDefinition is called when entering the primaryKeyDefinition production.
	EnterPrimaryKeyDefinition(c *PrimaryKeyDefinitionContext)

	// EnterSinglePrimaryKey is called when entering the singlePrimaryKey production.
	EnterSinglePrimaryKey(c *SinglePrimaryKeyContext)

	// EnterCompoundKey is called when entering the compoundKey production.
	EnterCompoundKey(c *CompoundKeyContext)

	// EnterCompositeKey is called when entering the compositeKey production.
	EnterCompositeKey(c *CompositeKeyContext)

	// EnterPartitionKeyList is called when entering the partitionKeyList production.
	EnterPartitionKeyList(c *PartitionKeyListContext)

	// EnterClusteringKeyList is called when entering the clusteringKeyList production.
	EnterClusteringKeyList(c *ClusteringKeyListContext)

	// EnterPartitionKey is called when entering the partitionKey production.
	EnterPartitionKey(c *PartitionKeyContext)

	// EnterClusteringKey is called when entering the clusteringKey production.
	EnterClusteringKey(c *ClusteringKeyContext)

	// EnterApplyBatch is called when entering the applyBatch production.
	EnterApplyBatch(c *ApplyBatchContext)

	// EnterBeginBatch is called when entering the beginBatch production.
	EnterBeginBatch(c *BeginBatchContext)

	// EnterBatchType is called when entering the batchType production.
	EnterBatchType(c *BatchTypeContext)

	// EnterAlterKeyspace is called when entering the alterKeyspace production.
	EnterAlterKeyspace(c *AlterKeyspaceContext)

	// EnterReplicationList is called when entering the replicationList production.
	EnterReplicationList(c *ReplicationListContext)

	// EnterReplicationListItem is called when entering the replicationListItem production.
	EnterReplicationListItem(c *ReplicationListItemContext)

	// EnterDurableWrites is called when entering the durableWrites production.
	EnterDurableWrites(c *DurableWritesContext)

	// EnterUse_ is called when entering the use_ production.
	EnterUse_(c *Use_Context)

	// EnterTruncate is called when entering the truncate production.
	EnterTruncate(c *TruncateContext)

	// EnterCreateIndex is called when entering the createIndex production.
	EnterCreateIndex(c *CreateIndexContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterIndexColumnSpec is called when entering the indexColumnSpec production.
	EnterIndexColumnSpec(c *IndexColumnSpecContext)

	// EnterIndexKeysSpec is called when entering the indexKeysSpec production.
	EnterIndexKeysSpec(c *IndexKeysSpecContext)

	// EnterIndexEntriesSSpec is called when entering the indexEntriesSSpec production.
	EnterIndexEntriesSSpec(c *IndexEntriesSSpecContext)

	// EnterIndexFullSpec is called when entering the indexFullSpec production.
	EnterIndexFullSpec(c *IndexFullSpecContext)

	// EnterDelete_ is called when entering the delete_ production.
	EnterDelete_(c *Delete_Context)

	// EnterDeleteColumnList is called when entering the deleteColumnList production.
	EnterDeleteColumnList(c *DeleteColumnListContext)

	// EnterDeleteColumnItem is called when entering the deleteColumnItem production.
	EnterDeleteColumnItem(c *DeleteColumnItemContext)

	// EnterUpdate is called when entering the update production.
	EnterUpdate(c *UpdateContext)

	// EnterIfSpec is called when entering the ifSpec production.
	EnterIfSpec(c *IfSpecContext)

	// EnterIfConditionList is called when entering the ifConditionList production.
	EnterIfConditionList(c *IfConditionListContext)

	// EnterIfCondition is called when entering the ifCondition production.
	EnterIfCondition(c *IfConditionContext)

	// EnterAssignments is called when entering the assignments production.
	EnterAssignments(c *AssignmentsContext)

	// EnterAssignmentElement is called when entering the assignmentElement production.
	EnterAssignmentElement(c *AssignmentElementContext)

	// EnterAssignmentSet is called when entering the assignmentSet production.
	EnterAssignmentSet(c *AssignmentSetContext)

	// EnterAssignmentMap is called when entering the assignmentMap production.
	EnterAssignmentMap(c *AssignmentMapContext)

	// EnterAssignmentList is called when entering the assignmentList production.
	EnterAssignmentList(c *AssignmentListContext)

	// EnterAssignmentTuple is called when entering the assignmentTuple production.
	EnterAssignmentTuple(c *AssignmentTupleContext)

	// EnterInsert is called when entering the insert production.
	EnterInsert(c *InsertContext)

	// EnterUsingTtlTimestamp is called when entering the usingTtlTimestamp production.
	EnterUsingTtlTimestamp(c *UsingTtlTimestampContext)

	// EnterTimestamp is called when entering the timestamp production.
	EnterTimestamp(c *TimestampContext)

	// EnterTtl is called when entering the ttl production.
	EnterTtl(c *TtlContext)

	// EnterUsingTimestampSpec is called when entering the usingTimestampSpec production.
	EnterUsingTimestampSpec(c *UsingTimestampSpecContext)

	// EnterIfNotExist is called when entering the ifNotExist production.
	EnterIfNotExist(c *IfNotExistContext)

	// EnterIfExist is called when entering the ifExist production.
	EnterIfExist(c *IfExistContext)

	// EnterInsertValuesSpec is called when entering the insertValuesSpec production.
	EnterInsertValuesSpec(c *InsertValuesSpecContext)

	// EnterInsertColumnSpec is called when entering the insertColumnSpec production.
	EnterInsertColumnSpec(c *InsertColumnSpecContext)

	// EnterColumnList is called when entering the columnList production.
	EnterColumnList(c *ColumnListContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSelect_ is called when entering the select_ production.
	EnterSelect_(c *Select_Context)

	// EnterAllowFilteringSpec is called when entering the allowFilteringSpec production.
	EnterAllowFilteringSpec(c *AllowFilteringSpecContext)

	// EnterLimitSpec is called when entering the limitSpec production.
	EnterLimitSpec(c *LimitSpecContext)

	// EnterFromSpec is called when entering the fromSpec production.
	EnterFromSpec(c *FromSpecContext)

	// EnterFromSpecElement is called when entering the fromSpecElement production.
	EnterFromSpecElement(c *FromSpecElementContext)

	// EnterOrderSpec is called when entering the orderSpec production.
	EnterOrderSpec(c *OrderSpecContext)

	// EnterOrderSpecElement is called when entering the orderSpecElement production.
	EnterOrderSpecElement(c *OrderSpecElementContext)

	// EnterWhereSpec is called when entering the whereSpec production.
	EnterWhereSpec(c *WhereSpecContext)

	// EnterDistinctSpec is called when entering the distinctSpec production.
	EnterDistinctSpec(c *DistinctSpecContext)

	// EnterSelectElements is called when entering the selectElements production.
	EnterSelectElements(c *SelectElementsContext)

	// EnterSelectElement is called when entering the selectElement production.
	EnterSelectElement(c *SelectElementContext)

	// EnterRelationElements is called when entering the relationElements production.
	EnterRelationElements(c *RelationElementsContext)

	// EnterRelationElement is called when entering the relationElement production.
	EnterRelationElement(c *RelationElementContext)

	// EnterRelalationContains is called when entering the relalationContains production.
	EnterRelalationContains(c *RelalationContainsContext)

	// EnterRelalationContainsKey is called when entering the relalationContainsKey production.
	EnterRelalationContainsKey(c *RelalationContainsKeyContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterFunctionArgs is called when entering the functionArgs production.
	EnterFunctionArgs(c *FunctionArgsContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterHexadecimalLiteral is called when entering the hexadecimalLiteral production.
	EnterHexadecimalLiteral(c *HexadecimalLiteralContext)

	// EnterKeyspace is called when entering the keyspace production.
	EnterKeyspace(c *KeyspaceContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterColumn is called when entering the column production.
	EnterColumn(c *ColumnContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterDataTypeName is called when entering the dataTypeName production.
	EnterDataTypeName(c *DataTypeNameContext)

	// EnterDataTypeDefinition is called when entering the dataTypeDefinition production.
	EnterDataTypeDefinition(c *DataTypeDefinitionContext)

	// EnterOrderDirection is called when entering the orderDirection production.
	EnterOrderDirection(c *OrderDirectionContext)

	// EnterRole is called when entering the role production.
	EnterRole(c *RoleContext)

	// EnterTrigger is called when entering the trigger production.
	EnterTrigger(c *TriggerContext)

	// EnterTriggerClass is called when entering the triggerClass production.
	EnterTriggerClass(c *TriggerClassContext)

	// EnterMaterializedView is called when entering the materializedView production.
	EnterMaterializedView(c *MaterializedViewContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterAggregate is called when entering the aggregate production.
	EnterAggregate(c *AggregateContext)

	// EnterFunction_ is called when entering the function_ production.
	EnterFunction_(c *Function_Context)

	// EnterLanguage is called when entering the language production.
	EnterLanguage(c *LanguageContext)

	// EnterUser is called when entering the user production.
	EnterUser(c *UserContext)

	// EnterPassword is called when entering the password production.
	EnterPassword(c *PasswordContext)

	// EnterHashKey is called when entering the hashKey production.
	EnterHashKey(c *HashKeyContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterParamName is called when entering the paramName production.
	EnterParamName(c *ParamNameContext)

	// EnterKwAdd is called when entering the kwAdd production.
	EnterKwAdd(c *KwAddContext)

	// EnterKwAggregate is called when entering the kwAggregate production.
	EnterKwAggregate(c *KwAggregateContext)

	// EnterKwAll is called when entering the kwAll production.
	EnterKwAll(c *KwAllContext)

	// EnterKwAllPermissions is called when entering the kwAllPermissions production.
	EnterKwAllPermissions(c *KwAllPermissionsContext)

	// EnterKwAllow is called when entering the kwAllow production.
	EnterKwAllow(c *KwAllowContext)

	// EnterKwAlter is called when entering the kwAlter production.
	EnterKwAlter(c *KwAlterContext)

	// EnterKwAnd is called when entering the kwAnd production.
	EnterKwAnd(c *KwAndContext)

	// EnterKwApply is called when entering the kwApply production.
	EnterKwApply(c *KwApplyContext)

	// EnterKwAs is called when entering the kwAs production.
	EnterKwAs(c *KwAsContext)

	// EnterKwAsc is called when entering the kwAsc production.
	EnterKwAsc(c *KwAscContext)

	// EnterKwAuthorize is called when entering the kwAuthorize production.
	EnterKwAuthorize(c *KwAuthorizeContext)

	// EnterKwBatch is called when entering the kwBatch production.
	EnterKwBatch(c *KwBatchContext)

	// EnterKwBegin is called when entering the kwBegin production.
	EnterKwBegin(c *KwBeginContext)

	// EnterKwBy is called when entering the kwBy production.
	EnterKwBy(c *KwByContext)

	// EnterKwCalled is called when entering the kwCalled production.
	EnterKwCalled(c *KwCalledContext)

	// EnterKwClustering is called when entering the kwClustering production.
	EnterKwClustering(c *KwClusteringContext)

	// EnterKwCompact is called when entering the kwCompact production.
	EnterKwCompact(c *KwCompactContext)

	// EnterKwContains is called when entering the kwContains production.
	EnterKwContains(c *KwContainsContext)

	// EnterKwCreate is called when entering the kwCreate production.
	EnterKwCreate(c *KwCreateContext)

	// EnterKwDelete is called when entering the kwDelete production.
	EnterKwDelete(c *KwDeleteContext)

	// EnterKwDesc is called when entering the kwDesc production.
	EnterKwDesc(c *KwDescContext)

	// EnterKwDescibe is called when entering the kwDescibe production.
	EnterKwDescibe(c *KwDescibeContext)

	// EnterKwDistinct is called when entering the kwDistinct production.
	EnterKwDistinct(c *KwDistinctContext)

	// EnterKwDrop is called when entering the kwDrop production.
	EnterKwDrop(c *KwDropContext)

	// EnterKwDurableWrites is called when entering the kwDurableWrites production.
	EnterKwDurableWrites(c *KwDurableWritesContext)

	// EnterKwEntries is called when entering the kwEntries production.
	EnterKwEntries(c *KwEntriesContext)

	// EnterKwExecute is called when entering the kwExecute production.
	EnterKwExecute(c *KwExecuteContext)

	// EnterKwExists is called when entering the kwExists production.
	EnterKwExists(c *KwExistsContext)

	// EnterKwFiltering is called when entering the kwFiltering production.
	EnterKwFiltering(c *KwFilteringContext)

	// EnterKwFinalfunc is called when entering the kwFinalfunc production.
	EnterKwFinalfunc(c *KwFinalfuncContext)

	// EnterKwFrom is called when entering the kwFrom production.
	EnterKwFrom(c *KwFromContext)

	// EnterKwFull is called when entering the kwFull production.
	EnterKwFull(c *KwFullContext)

	// EnterKwFunction is called when entering the kwFunction production.
	EnterKwFunction(c *KwFunctionContext)

	// EnterKwFunctions is called when entering the kwFunctions production.
	EnterKwFunctions(c *KwFunctionsContext)

	// EnterKwGrant is called when entering the kwGrant production.
	EnterKwGrant(c *KwGrantContext)

	// EnterKwIf is called when entering the kwIf production.
	EnterKwIf(c *KwIfContext)

	// EnterKwIn is called when entering the kwIn production.
	EnterKwIn(c *KwInContext)

	// EnterKwIndex is called when entering the kwIndex production.
	EnterKwIndex(c *KwIndexContext)

	// EnterKwInitcond is called when entering the kwInitcond production.
	EnterKwInitcond(c *KwInitcondContext)

	// EnterKwInput is called when entering the kwInput production.
	EnterKwInput(c *KwInputContext)

	// EnterKwInsert is called when entering the kwInsert production.
	EnterKwInsert(c *KwInsertContext)

	// EnterKwInto is called when entering the kwInto production.
	EnterKwInto(c *KwIntoContext)

	// EnterKwIs is called when entering the kwIs production.
	EnterKwIs(c *KwIsContext)

	// EnterKwJson is called when entering the kwJson production.
	EnterKwJson(c *KwJsonContext)

	// EnterKwKey is called when entering the kwKey production.
	EnterKwKey(c *KwKeyContext)

	// EnterKwKeys is called when entering the kwKeys production.
	EnterKwKeys(c *KwKeysContext)

	// EnterKwKeyspace is called when entering the kwKeyspace production.
	EnterKwKeyspace(c *KwKeyspaceContext)

	// EnterKwKeyspaces is called when entering the kwKeyspaces production.
	EnterKwKeyspaces(c *KwKeyspacesContext)

	// EnterKwLanguage is called when entering the kwLanguage production.
	EnterKwLanguage(c *KwLanguageContext)

	// EnterKwLimit is called when entering the kwLimit production.
	EnterKwLimit(c *KwLimitContext)

	// EnterKwList is called when entering the kwList production.
	EnterKwList(c *KwListContext)

	// EnterKwLogged is called when entering the kwLogged production.
	EnterKwLogged(c *KwLoggedContext)

	// EnterKwLogin is called when entering the kwLogin production.
	EnterKwLogin(c *KwLoginContext)

	// EnterKwMaterialized is called when entering the kwMaterialized production.
	EnterKwMaterialized(c *KwMaterializedContext)

	// EnterKwModify is called when entering the kwModify production.
	EnterKwModify(c *KwModifyContext)

	// EnterKwNosuperuser is called when entering the kwNosuperuser production.
	EnterKwNosuperuser(c *KwNosuperuserContext)

	// EnterKwNorecursive is called when entering the kwNorecursive production.
	EnterKwNorecursive(c *KwNorecursiveContext)

	// EnterKwNot is called when entering the kwNot production.
	EnterKwNot(c *KwNotContext)

	// EnterKwNull is called when entering the kwNull production.
	EnterKwNull(c *KwNullContext)

	// EnterKwOf is called when entering the kwOf production.
	EnterKwOf(c *KwOfContext)

	// EnterKwOn is called when entering the kwOn production.
	EnterKwOn(c *KwOnContext)

	// EnterKwOptions is called when entering the kwOptions production.
	EnterKwOptions(c *KwOptionsContext)

	// EnterKwOr is called when entering the kwOr production.
	EnterKwOr(c *KwOrContext)

	// EnterKwOrder is called when entering the kwOrder production.
	EnterKwOrder(c *KwOrderContext)

	// EnterKwPassword is called when entering the kwPassword production.
	EnterKwPassword(c *KwPasswordContext)

	// EnterKwPrimary is called when entering the kwPrimary production.
	EnterKwPrimary(c *KwPrimaryContext)

	// EnterKwRename is called when entering the kwRename production.
	EnterKwRename(c *KwRenameContext)

	// EnterKwReplace is called when entering the kwReplace production.
	EnterKwReplace(c *KwReplaceContext)

	// EnterKwReplication is called when entering the kwReplication production.
	EnterKwReplication(c *KwReplicationContext)

	// EnterKwReturns is called when entering the kwReturns production.
	EnterKwReturns(c *KwReturnsContext)

	// EnterKwRole is called when entering the kwRole production.
	EnterKwRole(c *KwRoleContext)

	// EnterKwRoles is called when entering the kwRoles production.
	EnterKwRoles(c *KwRolesContext)

	// EnterKwSelect is called when entering the kwSelect production.
	EnterKwSelect(c *KwSelectContext)

	// EnterKwSet is called when entering the kwSet production.
	EnterKwSet(c *KwSetContext)

	// EnterKwSfunc is called when entering the kwSfunc production.
	EnterKwSfunc(c *KwSfuncContext)

	// EnterKwStorage is called when entering the kwStorage production.
	EnterKwStorage(c *KwStorageContext)

	// EnterKwStype is called when entering the kwStype production.
	EnterKwStype(c *KwStypeContext)

	// EnterKwSuperuser is called when entering the kwSuperuser production.
	EnterKwSuperuser(c *KwSuperuserContext)

	// EnterKwTable is called when entering the kwTable production.
	EnterKwTable(c *KwTableContext)

	// EnterKwTimestamp is called when entering the kwTimestamp production.
	EnterKwTimestamp(c *KwTimestampContext)

	// EnterKwTo is called when entering the kwTo production.
	EnterKwTo(c *KwToContext)

	// EnterKwTrigger is called when entering the kwTrigger production.
	EnterKwTrigger(c *KwTriggerContext)

	// EnterKwTruncate is called when entering the kwTruncate production.
	EnterKwTruncate(c *KwTruncateContext)

	// EnterKwTtl is called when entering the kwTtl production.
	EnterKwTtl(c *KwTtlContext)

	// EnterKwType is called when entering the kwType production.
	EnterKwType(c *KwTypeContext)

	// EnterKwUnlogged is called when entering the kwUnlogged production.
	EnterKwUnlogged(c *KwUnloggedContext)

	// EnterKwUpdate is called when entering the kwUpdate production.
	EnterKwUpdate(c *KwUpdateContext)

	// EnterKwUse is called when entering the kwUse production.
	EnterKwUse(c *KwUseContext)

	// EnterKwUser is called when entering the kwUser production.
	EnterKwUser(c *KwUserContext)

	// EnterKwUsers is called when entering the kwUsers production.
	EnterKwUsers(c *KwUsersContext)

	// EnterKwUsing is called when entering the kwUsing production.
	EnterKwUsing(c *KwUsingContext)

	// EnterKwValues is called when entering the kwValues production.
	EnterKwValues(c *KwValuesContext)

	// EnterKwView is called when entering the kwView production.
	EnterKwView(c *KwViewContext)

	// EnterKwWhere is called when entering the kwWhere production.
	EnterKwWhere(c *KwWhereContext)

	// EnterKwWith is called when entering the kwWith production.
	EnterKwWith(c *KwWithContext)

	// EnterKwRevoke is called when entering the kwRevoke production.
	EnterKwRevoke(c *KwRevokeContext)

	// EnterEof is called when entering the eof production.
	EnterEof(c *EofContext)

	// EnterSyntaxBracketLr is called when entering the syntaxBracketLr production.
	EnterSyntaxBracketLr(c *SyntaxBracketLrContext)

	// EnterSyntaxBracketRr is called when entering the syntaxBracketRr production.
	EnterSyntaxBracketRr(c *SyntaxBracketRrContext)

	// EnterSyntaxBracketLc is called when entering the syntaxBracketLc production.
	EnterSyntaxBracketLc(c *SyntaxBracketLcContext)

	// EnterSyntaxBracketRc is called when entering the syntaxBracketRc production.
	EnterSyntaxBracketRc(c *SyntaxBracketRcContext)

	// EnterSyntaxBracketLa is called when entering the syntaxBracketLa production.
	EnterSyntaxBracketLa(c *SyntaxBracketLaContext)

	// EnterSyntaxBracketRa is called when entering the syntaxBracketRa production.
	EnterSyntaxBracketRa(c *SyntaxBracketRaContext)

	// EnterSyntaxBracketLs is called when entering the syntaxBracketLs production.
	EnterSyntaxBracketLs(c *SyntaxBracketLsContext)

	// EnterSyntaxBracketRs is called when entering the syntaxBracketRs production.
	EnterSyntaxBracketRs(c *SyntaxBracketRsContext)

	// EnterSyntaxComma is called when entering the syntaxComma production.
	EnterSyntaxComma(c *SyntaxCommaContext)

	// EnterSyntaxColon is called when entering the syntaxColon production.
	EnterSyntaxColon(c *SyntaxColonContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitCqls is called when exiting the cqls production.
	ExitCqls(c *CqlsContext)

	// ExitStatementSeparator is called when exiting the statementSeparator production.
	ExitStatementSeparator(c *StatementSeparatorContext)

	// ExitEmpty_ is called when exiting the empty_ production.
	ExitEmpty_(c *Empty_Context)

	// ExitCql is called when exiting the cql production.
	ExitCql(c *CqlContext)

	// ExitRevoke is called when exiting the revoke production.
	ExitRevoke(c *RevokeContext)

	// ExitListUsers is called when exiting the listUsers production.
	ExitListUsers(c *ListUsersContext)

	// ExitListRoles is called when exiting the listRoles production.
	ExitListRoles(c *ListRolesContext)

	// ExitListPermissions is called when exiting the listPermissions production.
	ExitListPermissions(c *ListPermissionsContext)

	// ExitGrant is called when exiting the grant production.
	ExitGrant(c *GrantContext)

	// ExitPriviledge is called when exiting the priviledge production.
	ExitPriviledge(c *PriviledgeContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitCreateUser is called when exiting the createUser production.
	ExitCreateUser(c *CreateUserContext)

	// ExitCreateRole is called when exiting the createRole production.
	ExitCreateRole(c *CreateRoleContext)

	// ExitCreateType is called when exiting the createType production.
	ExitCreateType(c *CreateTypeContext)

	// ExitTypeMemberColumnList is called when exiting the typeMemberColumnList production.
	ExitTypeMemberColumnList(c *TypeMemberColumnListContext)

	// ExitCreateTrigger is called when exiting the createTrigger production.
	ExitCreateTrigger(c *CreateTriggerContext)

	// ExitCreateMaterializedView is called when exiting the createMaterializedView production.
	ExitCreateMaterializedView(c *CreateMaterializedViewContext)

	// ExitMaterializedViewWhere is called when exiting the materializedViewWhere production.
	ExitMaterializedViewWhere(c *MaterializedViewWhereContext)

	// ExitColumnNotNullList is called when exiting the columnNotNullList production.
	ExitColumnNotNullList(c *ColumnNotNullListContext)

	// ExitColumnNotNull is called when exiting the columnNotNull production.
	ExitColumnNotNull(c *ColumnNotNullContext)

	// ExitMaterializedViewOptions is called when exiting the materializedViewOptions production.
	ExitMaterializedViewOptions(c *MaterializedViewOptionsContext)

	// ExitCreateKeyspace is called when exiting the createKeyspace production.
	ExitCreateKeyspace(c *CreateKeyspaceContext)

	// ExitCreateFunction is called when exiting the createFunction production.
	ExitCreateFunction(c *CreateFunctionContext)

	// ExitCodeBlock is called when exiting the codeBlock production.
	ExitCodeBlock(c *CodeBlockContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitReturnMode is called when exiting the returnMode production.
	ExitReturnMode(c *ReturnModeContext)

	// ExitCreateAggregate is called when exiting the createAggregate production.
	ExitCreateAggregate(c *CreateAggregateContext)

	// ExitInitCondDefinition is called when exiting the initCondDefinition production.
	ExitInitCondDefinition(c *InitCondDefinitionContext)

	// ExitInitCondHash is called when exiting the initCondHash production.
	ExitInitCondHash(c *InitCondHashContext)

	// ExitInitCondHashItem is called when exiting the initCondHashItem production.
	ExitInitCondHashItem(c *InitCondHashItemContext)

	// ExitInitCondListNested is called when exiting the initCondListNested production.
	ExitInitCondListNested(c *InitCondListNestedContext)

	// ExitInitCondList is called when exiting the initCondList production.
	ExitInitCondList(c *InitCondListContext)

	// ExitOrReplace is called when exiting the orReplace production.
	ExitOrReplace(c *OrReplaceContext)

	// ExitAlterUser is called when exiting the alterUser production.
	ExitAlterUser(c *AlterUserContext)

	// ExitUserPassword is called when exiting the userPassword production.
	ExitUserPassword(c *UserPasswordContext)

	// ExitUserSuperUser is called when exiting the userSuperUser production.
	ExitUserSuperUser(c *UserSuperUserContext)

	// ExitAlterType is called when exiting the alterType production.
	ExitAlterType(c *AlterTypeContext)

	// ExitAlterTypeOperation is called when exiting the alterTypeOperation production.
	ExitAlterTypeOperation(c *AlterTypeOperationContext)

	// ExitAlterTypeRename is called when exiting the alterTypeRename production.
	ExitAlterTypeRename(c *AlterTypeRenameContext)

	// ExitAlterTypeRenameList is called when exiting the alterTypeRenameList production.
	ExitAlterTypeRenameList(c *AlterTypeRenameListContext)

	// ExitAlterTypeRenameItem is called when exiting the alterTypeRenameItem production.
	ExitAlterTypeRenameItem(c *AlterTypeRenameItemContext)

	// ExitAlterTypeAdd is called when exiting the alterTypeAdd production.
	ExitAlterTypeAdd(c *AlterTypeAddContext)

	// ExitAlterTypeAlterType is called when exiting the alterTypeAlterType production.
	ExitAlterTypeAlterType(c *AlterTypeAlterTypeContext)

	// ExitAlterTable is called when exiting the alterTable production.
	ExitAlterTable(c *AlterTableContext)

	// ExitAlterTableOperation is called when exiting the alterTableOperation production.
	ExitAlterTableOperation(c *AlterTableOperationContext)

	// ExitAlterTableWith is called when exiting the alterTableWith production.
	ExitAlterTableWith(c *AlterTableWithContext)

	// ExitAlterTableRename is called when exiting the alterTableRename production.
	ExitAlterTableRename(c *AlterTableRenameContext)

	// ExitAlterTableDropCompactStorage is called when exiting the alterTableDropCompactStorage production.
	ExitAlterTableDropCompactStorage(c *AlterTableDropCompactStorageContext)

	// ExitAlterTableDropColumns is called when exiting the alterTableDropColumns production.
	ExitAlterTableDropColumns(c *AlterTableDropColumnsContext)

	// ExitAlterTableDropColumnList is called when exiting the alterTableDropColumnList production.
	ExitAlterTableDropColumnList(c *AlterTableDropColumnListContext)

	// ExitAlterTableAdd is called when exiting the alterTableAdd production.
	ExitAlterTableAdd(c *AlterTableAddContext)

	// ExitAlterTableColumnDefinition is called when exiting the alterTableColumnDefinition production.
	ExitAlterTableColumnDefinition(c *AlterTableColumnDefinitionContext)

	// ExitAlterRole is called when exiting the alterRole production.
	ExitAlterRole(c *AlterRoleContext)

	// ExitRoleWith is called when exiting the roleWith production.
	ExitRoleWith(c *RoleWithContext)

	// ExitRoleWithOptions is called when exiting the roleWithOptions production.
	ExitRoleWithOptions(c *RoleWithOptionsContext)

	// ExitAlterMaterializedView is called when exiting the alterMaterializedView production.
	ExitAlterMaterializedView(c *AlterMaterializedViewContext)

	// ExitDropUser is called when exiting the dropUser production.
	ExitDropUser(c *DropUserContext)

	// ExitDropType is called when exiting the dropType production.
	ExitDropType(c *DropTypeContext)

	// ExitDropMaterializedView is called when exiting the dropMaterializedView production.
	ExitDropMaterializedView(c *DropMaterializedViewContext)

	// ExitDropAggregate is called when exiting the dropAggregate production.
	ExitDropAggregate(c *DropAggregateContext)

	// ExitDropFunction is called when exiting the dropFunction production.
	ExitDropFunction(c *DropFunctionContext)

	// ExitDropTrigger is called when exiting the dropTrigger production.
	ExitDropTrigger(c *DropTriggerContext)

	// ExitDropRole is called when exiting the dropRole production.
	ExitDropRole(c *DropRoleContext)

	// ExitDropTable is called when exiting the dropTable production.
	ExitDropTable(c *DropTableContext)

	// ExitDropKeyspace is called when exiting the dropKeyspace production.
	ExitDropKeyspace(c *DropKeyspaceContext)

	// ExitDropIndex is called when exiting the dropIndex production.
	ExitDropIndex(c *DropIndexContext)

	// ExitCreateTable is called when exiting the createTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitWithElement is called when exiting the withElement production.
	ExitWithElement(c *WithElementContext)

	// ExitClusteringOrder is called when exiting the clusteringOrder production.
	ExitClusteringOrder(c *ClusteringOrderContext)

	// ExitTableOptions is called when exiting the tableOptions production.
	ExitTableOptions(c *TableOptionsContext)

	// ExitTableOptionItem is called when exiting the tableOptionItem production.
	ExitTableOptionItem(c *TableOptionItemContext)

	// ExitTableOptionName is called when exiting the tableOptionName production.
	ExitTableOptionName(c *TableOptionNameContext)

	// ExitTableOptionValue is called when exiting the tableOptionValue production.
	ExitTableOptionValue(c *TableOptionValueContext)

	// ExitOptionHash is called when exiting the optionHash production.
	ExitOptionHash(c *OptionHashContext)

	// ExitOptionHashItem is called when exiting the optionHashItem production.
	ExitOptionHashItem(c *OptionHashItemContext)

	// ExitOptionHashKey is called when exiting the optionHashKey production.
	ExitOptionHashKey(c *OptionHashKeyContext)

	// ExitOptionHashValue is called when exiting the optionHashValue production.
	ExitOptionHashValue(c *OptionHashValueContext)

	// ExitColumnDefinitionList is called when exiting the columnDefinitionList production.
	ExitColumnDefinitionList(c *ColumnDefinitionListContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitPrimaryKeyColumn is called when exiting the primaryKeyColumn production.
	ExitPrimaryKeyColumn(c *PrimaryKeyColumnContext)

	// ExitPrimaryKeyElement is called when exiting the primaryKeyElement production.
	ExitPrimaryKeyElement(c *PrimaryKeyElementContext)

	// ExitPrimaryKeyDefinition is called when exiting the primaryKeyDefinition production.
	ExitPrimaryKeyDefinition(c *PrimaryKeyDefinitionContext)

	// ExitSinglePrimaryKey is called when exiting the singlePrimaryKey production.
	ExitSinglePrimaryKey(c *SinglePrimaryKeyContext)

	// ExitCompoundKey is called when exiting the compoundKey production.
	ExitCompoundKey(c *CompoundKeyContext)

	// ExitCompositeKey is called when exiting the compositeKey production.
	ExitCompositeKey(c *CompositeKeyContext)

	// ExitPartitionKeyList is called when exiting the partitionKeyList production.
	ExitPartitionKeyList(c *PartitionKeyListContext)

	// ExitClusteringKeyList is called when exiting the clusteringKeyList production.
	ExitClusteringKeyList(c *ClusteringKeyListContext)

	// ExitPartitionKey is called when exiting the partitionKey production.
	ExitPartitionKey(c *PartitionKeyContext)

	// ExitClusteringKey is called when exiting the clusteringKey production.
	ExitClusteringKey(c *ClusteringKeyContext)

	// ExitApplyBatch is called when exiting the applyBatch production.
	ExitApplyBatch(c *ApplyBatchContext)

	// ExitBeginBatch is called when exiting the beginBatch production.
	ExitBeginBatch(c *BeginBatchContext)

	// ExitBatchType is called when exiting the batchType production.
	ExitBatchType(c *BatchTypeContext)

	// ExitAlterKeyspace is called when exiting the alterKeyspace production.
	ExitAlterKeyspace(c *AlterKeyspaceContext)

	// ExitReplicationList is called when exiting the replicationList production.
	ExitReplicationList(c *ReplicationListContext)

	// ExitReplicationListItem is called when exiting the replicationListItem production.
	ExitReplicationListItem(c *ReplicationListItemContext)

	// ExitDurableWrites is called when exiting the durableWrites production.
	ExitDurableWrites(c *DurableWritesContext)

	// ExitUse_ is called when exiting the use_ production.
	ExitUse_(c *Use_Context)

	// ExitTruncate is called when exiting the truncate production.
	ExitTruncate(c *TruncateContext)

	// ExitCreateIndex is called when exiting the createIndex production.
	ExitCreateIndex(c *CreateIndexContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitIndexColumnSpec is called when exiting the indexColumnSpec production.
	ExitIndexColumnSpec(c *IndexColumnSpecContext)

	// ExitIndexKeysSpec is called when exiting the indexKeysSpec production.
	ExitIndexKeysSpec(c *IndexKeysSpecContext)

	// ExitIndexEntriesSSpec is called when exiting the indexEntriesSSpec production.
	ExitIndexEntriesSSpec(c *IndexEntriesSSpecContext)

	// ExitIndexFullSpec is called when exiting the indexFullSpec production.
	ExitIndexFullSpec(c *IndexFullSpecContext)

	// ExitDelete_ is called when exiting the delete_ production.
	ExitDelete_(c *Delete_Context)

	// ExitDeleteColumnList is called when exiting the deleteColumnList production.
	ExitDeleteColumnList(c *DeleteColumnListContext)

	// ExitDeleteColumnItem is called when exiting the deleteColumnItem production.
	ExitDeleteColumnItem(c *DeleteColumnItemContext)

	// ExitUpdate is called when exiting the update production.
	ExitUpdate(c *UpdateContext)

	// ExitIfSpec is called when exiting the ifSpec production.
	ExitIfSpec(c *IfSpecContext)

	// ExitIfConditionList is called when exiting the ifConditionList production.
	ExitIfConditionList(c *IfConditionListContext)

	// ExitIfCondition is called when exiting the ifCondition production.
	ExitIfCondition(c *IfConditionContext)

	// ExitAssignments is called when exiting the assignments production.
	ExitAssignments(c *AssignmentsContext)

	// ExitAssignmentElement is called when exiting the assignmentElement production.
	ExitAssignmentElement(c *AssignmentElementContext)

	// ExitAssignmentSet is called when exiting the assignmentSet production.
	ExitAssignmentSet(c *AssignmentSetContext)

	// ExitAssignmentMap is called when exiting the assignmentMap production.
	ExitAssignmentMap(c *AssignmentMapContext)

	// ExitAssignmentList is called when exiting the assignmentList production.
	ExitAssignmentList(c *AssignmentListContext)

	// ExitAssignmentTuple is called when exiting the assignmentTuple production.
	ExitAssignmentTuple(c *AssignmentTupleContext)

	// ExitInsert is called when exiting the insert production.
	ExitInsert(c *InsertContext)

	// ExitUsingTtlTimestamp is called when exiting the usingTtlTimestamp production.
	ExitUsingTtlTimestamp(c *UsingTtlTimestampContext)

	// ExitTimestamp is called when exiting the timestamp production.
	ExitTimestamp(c *TimestampContext)

	// ExitTtl is called when exiting the ttl production.
	ExitTtl(c *TtlContext)

	// ExitUsingTimestampSpec is called when exiting the usingTimestampSpec production.
	ExitUsingTimestampSpec(c *UsingTimestampSpecContext)

	// ExitIfNotExist is called when exiting the ifNotExist production.
	ExitIfNotExist(c *IfNotExistContext)

	// ExitIfExist is called when exiting the ifExist production.
	ExitIfExist(c *IfExistContext)

	// ExitInsertValuesSpec is called when exiting the insertValuesSpec production.
	ExitInsertValuesSpec(c *InsertValuesSpecContext)

	// ExitInsertColumnSpec is called when exiting the insertColumnSpec production.
	ExitInsertColumnSpec(c *InsertColumnSpecContext)

	// ExitColumnList is called when exiting the columnList production.
	ExitColumnList(c *ColumnListContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSelect_ is called when exiting the select_ production.
	ExitSelect_(c *Select_Context)

	// ExitAllowFilteringSpec is called when exiting the allowFilteringSpec production.
	ExitAllowFilteringSpec(c *AllowFilteringSpecContext)

	// ExitLimitSpec is called when exiting the limitSpec production.
	ExitLimitSpec(c *LimitSpecContext)

	// ExitFromSpec is called when exiting the fromSpec production.
	ExitFromSpec(c *FromSpecContext)

	// ExitFromSpecElement is called when exiting the fromSpecElement production.
	ExitFromSpecElement(c *FromSpecElementContext)

	// ExitOrderSpec is called when exiting the orderSpec production.
	ExitOrderSpec(c *OrderSpecContext)

	// ExitOrderSpecElement is called when exiting the orderSpecElement production.
	ExitOrderSpecElement(c *OrderSpecElementContext)

	// ExitWhereSpec is called when exiting the whereSpec production.
	ExitWhereSpec(c *WhereSpecContext)

	// ExitDistinctSpec is called when exiting the distinctSpec production.
	ExitDistinctSpec(c *DistinctSpecContext)

	// ExitSelectElements is called when exiting the selectElements production.
	ExitSelectElements(c *SelectElementsContext)

	// ExitSelectElement is called when exiting the selectElement production.
	ExitSelectElement(c *SelectElementContext)

	// ExitRelationElements is called when exiting the relationElements production.
	ExitRelationElements(c *RelationElementsContext)

	// ExitRelationElement is called when exiting the relationElement production.
	ExitRelationElement(c *RelationElementContext)

	// ExitRelalationContains is called when exiting the relalationContains production.
	ExitRelalationContains(c *RelalationContainsContext)

	// ExitRelalationContainsKey is called when exiting the relalationContainsKey production.
	ExitRelalationContainsKey(c *RelalationContainsKeyContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitFunctionArgs is called when exiting the functionArgs production.
	ExitFunctionArgs(c *FunctionArgsContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitHexadecimalLiteral is called when exiting the hexadecimalLiteral production.
	ExitHexadecimalLiteral(c *HexadecimalLiteralContext)

	// ExitKeyspace is called when exiting the keyspace production.
	ExitKeyspace(c *KeyspaceContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitColumn is called when exiting the column production.
	ExitColumn(c *ColumnContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitDataTypeName is called when exiting the dataTypeName production.
	ExitDataTypeName(c *DataTypeNameContext)

	// ExitDataTypeDefinition is called when exiting the dataTypeDefinition production.
	ExitDataTypeDefinition(c *DataTypeDefinitionContext)

	// ExitOrderDirection is called when exiting the orderDirection production.
	ExitOrderDirection(c *OrderDirectionContext)

	// ExitRole is called when exiting the role production.
	ExitRole(c *RoleContext)

	// ExitTrigger is called when exiting the trigger production.
	ExitTrigger(c *TriggerContext)

	// ExitTriggerClass is called when exiting the triggerClass production.
	ExitTriggerClass(c *TriggerClassContext)

	// ExitMaterializedView is called when exiting the materializedView production.
	ExitMaterializedView(c *MaterializedViewContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitAggregate is called when exiting the aggregate production.
	ExitAggregate(c *AggregateContext)

	// ExitFunction_ is called when exiting the function_ production.
	ExitFunction_(c *Function_Context)

	// ExitLanguage is called when exiting the language production.
	ExitLanguage(c *LanguageContext)

	// ExitUser is called when exiting the user production.
	ExitUser(c *UserContext)

	// ExitPassword is called when exiting the password production.
	ExitPassword(c *PasswordContext)

	// ExitHashKey is called when exiting the hashKey production.
	ExitHashKey(c *HashKeyContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitParamName is called when exiting the paramName production.
	ExitParamName(c *ParamNameContext)

	// ExitKwAdd is called when exiting the kwAdd production.
	ExitKwAdd(c *KwAddContext)

	// ExitKwAggregate is called when exiting the kwAggregate production.
	ExitKwAggregate(c *KwAggregateContext)

	// ExitKwAll is called when exiting the kwAll production.
	ExitKwAll(c *KwAllContext)

	// ExitKwAllPermissions is called when exiting the kwAllPermissions production.
	ExitKwAllPermissions(c *KwAllPermissionsContext)

	// ExitKwAllow is called when exiting the kwAllow production.
	ExitKwAllow(c *KwAllowContext)

	// ExitKwAlter is called when exiting the kwAlter production.
	ExitKwAlter(c *KwAlterContext)

	// ExitKwAnd is called when exiting the kwAnd production.
	ExitKwAnd(c *KwAndContext)

	// ExitKwApply is called when exiting the kwApply production.
	ExitKwApply(c *KwApplyContext)

	// ExitKwAs is called when exiting the kwAs production.
	ExitKwAs(c *KwAsContext)

	// ExitKwAsc is called when exiting the kwAsc production.
	ExitKwAsc(c *KwAscContext)

	// ExitKwAuthorize is called when exiting the kwAuthorize production.
	ExitKwAuthorize(c *KwAuthorizeContext)

	// ExitKwBatch is called when exiting the kwBatch production.
	ExitKwBatch(c *KwBatchContext)

	// ExitKwBegin is called when exiting the kwBegin production.
	ExitKwBegin(c *KwBeginContext)

	// ExitKwBy is called when exiting the kwBy production.
	ExitKwBy(c *KwByContext)

	// ExitKwCalled is called when exiting the kwCalled production.
	ExitKwCalled(c *KwCalledContext)

	// ExitKwClustering is called when exiting the kwClustering production.
	ExitKwClustering(c *KwClusteringContext)

	// ExitKwCompact is called when exiting the kwCompact production.
	ExitKwCompact(c *KwCompactContext)

	// ExitKwContains is called when exiting the kwContains production.
	ExitKwContains(c *KwContainsContext)

	// ExitKwCreate is called when exiting the kwCreate production.
	ExitKwCreate(c *KwCreateContext)

	// ExitKwDelete is called when exiting the kwDelete production.
	ExitKwDelete(c *KwDeleteContext)

	// ExitKwDesc is called when exiting the kwDesc production.
	ExitKwDesc(c *KwDescContext)

	// ExitKwDescibe is called when exiting the kwDescibe production.
	ExitKwDescibe(c *KwDescibeContext)

	// ExitKwDistinct is called when exiting the kwDistinct production.
	ExitKwDistinct(c *KwDistinctContext)

	// ExitKwDrop is called when exiting the kwDrop production.
	ExitKwDrop(c *KwDropContext)

	// ExitKwDurableWrites is called when exiting the kwDurableWrites production.
	ExitKwDurableWrites(c *KwDurableWritesContext)

	// ExitKwEntries is called when exiting the kwEntries production.
	ExitKwEntries(c *KwEntriesContext)

	// ExitKwExecute is called when exiting the kwExecute production.
	ExitKwExecute(c *KwExecuteContext)

	// ExitKwExists is called when exiting the kwExists production.
	ExitKwExists(c *KwExistsContext)

	// ExitKwFiltering is called when exiting the kwFiltering production.
	ExitKwFiltering(c *KwFilteringContext)

	// ExitKwFinalfunc is called when exiting the kwFinalfunc production.
	ExitKwFinalfunc(c *KwFinalfuncContext)

	// ExitKwFrom is called when exiting the kwFrom production.
	ExitKwFrom(c *KwFromContext)

	// ExitKwFull is called when exiting the kwFull production.
	ExitKwFull(c *KwFullContext)

	// ExitKwFunction is called when exiting the kwFunction production.
	ExitKwFunction(c *KwFunctionContext)

	// ExitKwFunctions is called when exiting the kwFunctions production.
	ExitKwFunctions(c *KwFunctionsContext)

	// ExitKwGrant is called when exiting the kwGrant production.
	ExitKwGrant(c *KwGrantContext)

	// ExitKwIf is called when exiting the kwIf production.
	ExitKwIf(c *KwIfContext)

	// ExitKwIn is called when exiting the kwIn production.
	ExitKwIn(c *KwInContext)

	// ExitKwIndex is called when exiting the kwIndex production.
	ExitKwIndex(c *KwIndexContext)

	// ExitKwInitcond is called when exiting the kwInitcond production.
	ExitKwInitcond(c *KwInitcondContext)

	// ExitKwInput is called when exiting the kwInput production.
	ExitKwInput(c *KwInputContext)

	// ExitKwInsert is called when exiting the kwInsert production.
	ExitKwInsert(c *KwInsertContext)

	// ExitKwInto is called when exiting the kwInto production.
	ExitKwInto(c *KwIntoContext)

	// ExitKwIs is called when exiting the kwIs production.
	ExitKwIs(c *KwIsContext)

	// ExitKwJson is called when exiting the kwJson production.
	ExitKwJson(c *KwJsonContext)

	// ExitKwKey is called when exiting the kwKey production.
	ExitKwKey(c *KwKeyContext)

	// ExitKwKeys is called when exiting the kwKeys production.
	ExitKwKeys(c *KwKeysContext)

	// ExitKwKeyspace is called when exiting the kwKeyspace production.
	ExitKwKeyspace(c *KwKeyspaceContext)

	// ExitKwKeyspaces is called when exiting the kwKeyspaces production.
	ExitKwKeyspaces(c *KwKeyspacesContext)

	// ExitKwLanguage is called when exiting the kwLanguage production.
	ExitKwLanguage(c *KwLanguageContext)

	// ExitKwLimit is called when exiting the kwLimit production.
	ExitKwLimit(c *KwLimitContext)

	// ExitKwList is called when exiting the kwList production.
	ExitKwList(c *KwListContext)

	// ExitKwLogged is called when exiting the kwLogged production.
	ExitKwLogged(c *KwLoggedContext)

	// ExitKwLogin is called when exiting the kwLogin production.
	ExitKwLogin(c *KwLoginContext)

	// ExitKwMaterialized is called when exiting the kwMaterialized production.
	ExitKwMaterialized(c *KwMaterializedContext)

	// ExitKwModify is called when exiting the kwModify production.
	ExitKwModify(c *KwModifyContext)

	// ExitKwNosuperuser is called when exiting the kwNosuperuser production.
	ExitKwNosuperuser(c *KwNosuperuserContext)

	// ExitKwNorecursive is called when exiting the kwNorecursive production.
	ExitKwNorecursive(c *KwNorecursiveContext)

	// ExitKwNot is called when exiting the kwNot production.
	ExitKwNot(c *KwNotContext)

	// ExitKwNull is called when exiting the kwNull production.
	ExitKwNull(c *KwNullContext)

	// ExitKwOf is called when exiting the kwOf production.
	ExitKwOf(c *KwOfContext)

	// ExitKwOn is called when exiting the kwOn production.
	ExitKwOn(c *KwOnContext)

	// ExitKwOptions is called when exiting the kwOptions production.
	ExitKwOptions(c *KwOptionsContext)

	// ExitKwOr is called when exiting the kwOr production.
	ExitKwOr(c *KwOrContext)

	// ExitKwOrder is called when exiting the kwOrder production.
	ExitKwOrder(c *KwOrderContext)

	// ExitKwPassword is called when exiting the kwPassword production.
	ExitKwPassword(c *KwPasswordContext)

	// ExitKwPrimary is called when exiting the kwPrimary production.
	ExitKwPrimary(c *KwPrimaryContext)

	// ExitKwRename is called when exiting the kwRename production.
	ExitKwRename(c *KwRenameContext)

	// ExitKwReplace is called when exiting the kwReplace production.
	ExitKwReplace(c *KwReplaceContext)

	// ExitKwReplication is called when exiting the kwReplication production.
	ExitKwReplication(c *KwReplicationContext)

	// ExitKwReturns is called when exiting the kwReturns production.
	ExitKwReturns(c *KwReturnsContext)

	// ExitKwRole is called when exiting the kwRole production.
	ExitKwRole(c *KwRoleContext)

	// ExitKwRoles is called when exiting the kwRoles production.
	ExitKwRoles(c *KwRolesContext)

	// ExitKwSelect is called when exiting the kwSelect production.
	ExitKwSelect(c *KwSelectContext)

	// ExitKwSet is called when exiting the kwSet production.
	ExitKwSet(c *KwSetContext)

	// ExitKwSfunc is called when exiting the kwSfunc production.
	ExitKwSfunc(c *KwSfuncContext)

	// ExitKwStorage is called when exiting the kwStorage production.
	ExitKwStorage(c *KwStorageContext)

	// ExitKwStype is called when exiting the kwStype production.
	ExitKwStype(c *KwStypeContext)

	// ExitKwSuperuser is called when exiting the kwSuperuser production.
	ExitKwSuperuser(c *KwSuperuserContext)

	// ExitKwTable is called when exiting the kwTable production.
	ExitKwTable(c *KwTableContext)

	// ExitKwTimestamp is called when exiting the kwTimestamp production.
	ExitKwTimestamp(c *KwTimestampContext)

	// ExitKwTo is called when exiting the kwTo production.
	ExitKwTo(c *KwToContext)

	// ExitKwTrigger is called when exiting the kwTrigger production.
	ExitKwTrigger(c *KwTriggerContext)

	// ExitKwTruncate is called when exiting the kwTruncate production.
	ExitKwTruncate(c *KwTruncateContext)

	// ExitKwTtl is called when exiting the kwTtl production.
	ExitKwTtl(c *KwTtlContext)

	// ExitKwType is called when exiting the kwType production.
	ExitKwType(c *KwTypeContext)

	// ExitKwUnlogged is called when exiting the kwUnlogged production.
	ExitKwUnlogged(c *KwUnloggedContext)

	// ExitKwUpdate is called when exiting the kwUpdate production.
	ExitKwUpdate(c *KwUpdateContext)

	// ExitKwUse is called when exiting the kwUse production.
	ExitKwUse(c *KwUseContext)

	// ExitKwUser is called when exiting the kwUser production.
	ExitKwUser(c *KwUserContext)

	// ExitKwUsers is called when exiting the kwUsers production.
	ExitKwUsers(c *KwUsersContext)

	// ExitKwUsing is called when exiting the kwUsing production.
	ExitKwUsing(c *KwUsingContext)

	// ExitKwValues is called when exiting the kwValues production.
	ExitKwValues(c *KwValuesContext)

	// ExitKwView is called when exiting the kwView production.
	ExitKwView(c *KwViewContext)

	// ExitKwWhere is called when exiting the kwWhere production.
	ExitKwWhere(c *KwWhereContext)

	// ExitKwWith is called when exiting the kwWith production.
	ExitKwWith(c *KwWithContext)

	// ExitKwRevoke is called when exiting the kwRevoke production.
	ExitKwRevoke(c *KwRevokeContext)

	// ExitEof is called when exiting the eof production.
	ExitEof(c *EofContext)

	// ExitSyntaxBracketLr is called when exiting the syntaxBracketLr production.
	ExitSyntaxBracketLr(c *SyntaxBracketLrContext)

	// ExitSyntaxBracketRr is called when exiting the syntaxBracketRr production.
	ExitSyntaxBracketRr(c *SyntaxBracketRrContext)

	// ExitSyntaxBracketLc is called when exiting the syntaxBracketLc production.
	ExitSyntaxBracketLc(c *SyntaxBracketLcContext)

	// ExitSyntaxBracketRc is called when exiting the syntaxBracketRc production.
	ExitSyntaxBracketRc(c *SyntaxBracketRcContext)

	// ExitSyntaxBracketLa is called when exiting the syntaxBracketLa production.
	ExitSyntaxBracketLa(c *SyntaxBracketLaContext)

	// ExitSyntaxBracketRa is called when exiting the syntaxBracketRa production.
	ExitSyntaxBracketRa(c *SyntaxBracketRaContext)

	// ExitSyntaxBracketLs is called when exiting the syntaxBracketLs production.
	ExitSyntaxBracketLs(c *SyntaxBracketLsContext)

	// ExitSyntaxBracketRs is called when exiting the syntaxBracketRs production.
	ExitSyntaxBracketRs(c *SyntaxBracketRsContext)

	// ExitSyntaxComma is called when exiting the syntaxComma production.
	ExitSyntaxComma(c *SyntaxCommaContext)

	// ExitSyntaxColon is called when exiting the syntaxColon production.
	ExitSyntaxColon(c *SyntaxColonContext)
}
