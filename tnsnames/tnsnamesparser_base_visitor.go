// Generated from tnsnamesParser.g4 by ANTLR 4.7.

package tnsnames // tnsnamesParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasetnsnamesParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetnsnamesParserVisitor) VisitTnsnames(ctx *TnsnamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitTns_entry(ctx *Tns_entryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitIfile(ctx *IfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitLsnr_entry(ctx *Lsnr_entryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitLsnr_description(ctx *Lsnr_descriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitAlias_list(ctx *Alias_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitDescription_list(ctx *Description_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitDl_params(ctx *Dl_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitDl_parameter(ctx *Dl_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitDescription(ctx *DescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_params(ctx *D_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_parameter(ctx *D_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_enable(ctx *D_enableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_sdu(ctx *D_sduContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_recv_buf(ctx *D_recv_bufContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_send_buf(ctx *D_send_bufContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_service_type(ctx *D_service_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_security(ctx *D_securityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_conn_timeout(ctx *D_conn_timeoutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_retry_count(ctx *D_retry_countContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitD_tct(ctx *D_tctContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitDs_parameter(ctx *Ds_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitAddress_list(ctx *Address_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitAl_params(ctx *Al_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitAl_parameter(ctx *Al_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitAl_failover(ctx *Al_failoverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitAl_load_balance(ctx *Al_load_balanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitAl_source_route(ctx *Al_source_routeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitAddress(ctx *AddressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitA_params(ctx *A_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitA_parameter(ctx *A_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitProtocol_info(ctx *Protocol_infoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitTcp_protocol(ctx *Tcp_protocolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitTcp_params(ctx *Tcp_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitTcp_parameter(ctx *Tcp_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitTcp_host(ctx *Tcp_hostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitTcp_port(ctx *Tcp_portContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitTcp_tcp(ctx *Tcp_tcpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitHost(ctx *HostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitPort(ctx *PortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitIpc_protocol(ctx *Ipc_protocolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitIpc_params(ctx *Ipc_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitIpc_parameter(ctx *Ipc_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitIpc_ipc(ctx *Ipc_ipcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitIpc_key(ctx *Ipc_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitSpx_protocol(ctx *Spx_protocolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitSpx_params(ctx *Spx_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitSpx_parameter(ctx *Spx_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitSpx_spx(ctx *Spx_spxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitSpx_service(ctx *Spx_serviceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitNmp_protocol(ctx *Nmp_protocolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitNmp_params(ctx *Nmp_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitNmp_parameter(ctx *Nmp_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitNmp_nmp(ctx *Nmp_nmpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitNmp_server(ctx *Nmp_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitNmp_pipe(ctx *Nmp_pipeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBeq_protocol(ctx *Beq_protocolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBeq_params(ctx *Beq_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBeq_parameter(ctx *Beq_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBeq_beq(ctx *Beq_beqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBeq_program(ctx *Beq_programContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBeq_argv0(ctx *Beq_argv0Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBeq_args(ctx *Beq_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBa_parameter(ctx *Ba_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBa_description(ctx *Ba_descriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBad_params(ctx *Bad_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBad_parameter(ctx *Bad_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBad_local(ctx *Bad_localContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitBad_address(ctx *Bad_addressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitConnect_data(ctx *Connect_dataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_params(ctx *Cd_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_parameter(ctx *Cd_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_service_name(ctx *Cd_service_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_sid(ctx *Cd_sidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_instance_name(ctx *Cd_instance_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_failover_mode(ctx *Cd_failover_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_global_name(ctx *Cd_global_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_hs(ctx *Cd_hsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_rdb_database(ctx *Cd_rdb_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_server(ctx *Cd_serverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitCd_ur(ctx *Cd_urContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitFo_params(ctx *Fo_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitFo_parameter(ctx *Fo_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitFo_type(ctx *Fo_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitFo_backup(ctx *Fo_backupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitFo_method(ctx *Fo_methodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitFo_retries(ctx *Fo_retriesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetnsnamesParserVisitor) VisitFo_delay(ctx *Fo_delayContext) interface{} {
	return v.VisitChildren(ctx)
}
