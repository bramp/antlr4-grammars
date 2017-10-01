// Generated from tnsnamesParser.g4 by ANTLR 4.7.

package tnsnames // tnsnamesParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by tnsnamesParser.
type tnsnamesParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tnsnamesParser#tnsnames.
	VisitTnsnames(ctx *TnsnamesContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#tns_entry.
	VisitTns_entry(ctx *Tns_entryContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#ifile.
	VisitIfile(ctx *IfileContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#lsnr_entry.
	VisitLsnr_entry(ctx *Lsnr_entryContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#lsnr_description.
	VisitLsnr_description(ctx *Lsnr_descriptionContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#alias_list.
	VisitAlias_list(ctx *Alias_listContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#description_list.
	VisitDescription_list(ctx *Description_listContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#dl_params.
	VisitDl_params(ctx *Dl_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#dl_parameter.
	VisitDl_parameter(ctx *Dl_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#description.
	VisitDescription(ctx *DescriptionContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_params.
	VisitD_params(ctx *D_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_parameter.
	VisitD_parameter(ctx *D_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_enable.
	VisitD_enable(ctx *D_enableContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_sdu.
	VisitD_sdu(ctx *D_sduContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_recv_buf.
	VisitD_recv_buf(ctx *D_recv_bufContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_send_buf.
	VisitD_send_buf(ctx *D_send_bufContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_service_type.
	VisitD_service_type(ctx *D_service_typeContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_security.
	VisitD_security(ctx *D_securityContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_conn_timeout.
	VisitD_conn_timeout(ctx *D_conn_timeoutContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_retry_count.
	VisitD_retry_count(ctx *D_retry_countContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#d_tct.
	VisitD_tct(ctx *D_tctContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#ds_parameter.
	VisitDs_parameter(ctx *Ds_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#address_list.
	VisitAddress_list(ctx *Address_listContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#al_params.
	VisitAl_params(ctx *Al_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#al_parameter.
	VisitAl_parameter(ctx *Al_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#al_failover.
	VisitAl_failover(ctx *Al_failoverContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#al_load_balance.
	VisitAl_load_balance(ctx *Al_load_balanceContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#al_source_route.
	VisitAl_source_route(ctx *Al_source_routeContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#address.
	VisitAddress(ctx *AddressContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#a_params.
	VisitA_params(ctx *A_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#a_parameter.
	VisitA_parameter(ctx *A_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#protocol_info.
	VisitProtocol_info(ctx *Protocol_infoContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#tcp_protocol.
	VisitTcp_protocol(ctx *Tcp_protocolContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#tcp_params.
	VisitTcp_params(ctx *Tcp_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#tcp_parameter.
	VisitTcp_parameter(ctx *Tcp_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#tcp_host.
	VisitTcp_host(ctx *Tcp_hostContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#tcp_port.
	VisitTcp_port(ctx *Tcp_portContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#tcp_tcp.
	VisitTcp_tcp(ctx *Tcp_tcpContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#host.
	VisitHost(ctx *HostContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#port.
	VisitPort(ctx *PortContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#ipc_protocol.
	VisitIpc_protocol(ctx *Ipc_protocolContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#ipc_params.
	VisitIpc_params(ctx *Ipc_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#ipc_parameter.
	VisitIpc_parameter(ctx *Ipc_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#ipc_ipc.
	VisitIpc_ipc(ctx *Ipc_ipcContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#ipc_key.
	VisitIpc_key(ctx *Ipc_keyContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#spx_protocol.
	VisitSpx_protocol(ctx *Spx_protocolContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#spx_params.
	VisitSpx_params(ctx *Spx_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#spx_parameter.
	VisitSpx_parameter(ctx *Spx_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#spx_spx.
	VisitSpx_spx(ctx *Spx_spxContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#spx_service.
	VisitSpx_service(ctx *Spx_serviceContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#nmp_protocol.
	VisitNmp_protocol(ctx *Nmp_protocolContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#nmp_params.
	VisitNmp_params(ctx *Nmp_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#nmp_parameter.
	VisitNmp_parameter(ctx *Nmp_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#nmp_nmp.
	VisitNmp_nmp(ctx *Nmp_nmpContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#nmp_server.
	VisitNmp_server(ctx *Nmp_serverContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#nmp_pipe.
	VisitNmp_pipe(ctx *Nmp_pipeContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#beq_protocol.
	VisitBeq_protocol(ctx *Beq_protocolContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#beq_params.
	VisitBeq_params(ctx *Beq_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#beq_parameter.
	VisitBeq_parameter(ctx *Beq_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#beq_beq.
	VisitBeq_beq(ctx *Beq_beqContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#beq_program.
	VisitBeq_program(ctx *Beq_programContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#beq_argv0.
	VisitBeq_argv0(ctx *Beq_argv0Context) interface{}

	// Visit a parse tree produced by tnsnamesParser#beq_args.
	VisitBeq_args(ctx *Beq_argsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#ba_parameter.
	VisitBa_parameter(ctx *Ba_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#ba_description.
	VisitBa_description(ctx *Ba_descriptionContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#bad_params.
	VisitBad_params(ctx *Bad_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#bad_parameter.
	VisitBad_parameter(ctx *Bad_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#bad_local.
	VisitBad_local(ctx *Bad_localContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#bad_address.
	VisitBad_address(ctx *Bad_addressContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#connect_data.
	VisitConnect_data(ctx *Connect_dataContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_params.
	VisitCd_params(ctx *Cd_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_parameter.
	VisitCd_parameter(ctx *Cd_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_service_name.
	VisitCd_service_name(ctx *Cd_service_nameContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_sid.
	VisitCd_sid(ctx *Cd_sidContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_instance_name.
	VisitCd_instance_name(ctx *Cd_instance_nameContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_failover_mode.
	VisitCd_failover_mode(ctx *Cd_failover_modeContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_global_name.
	VisitCd_global_name(ctx *Cd_global_nameContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_hs.
	VisitCd_hs(ctx *Cd_hsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_rdb_database.
	VisitCd_rdb_database(ctx *Cd_rdb_databaseContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_server.
	VisitCd_server(ctx *Cd_serverContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#cd_ur.
	VisitCd_ur(ctx *Cd_urContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#fo_params.
	VisitFo_params(ctx *Fo_paramsContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#fo_parameter.
	VisitFo_parameter(ctx *Fo_parameterContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#fo_type.
	VisitFo_type(ctx *Fo_typeContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#fo_backup.
	VisitFo_backup(ctx *Fo_backupContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#fo_method.
	VisitFo_method(ctx *Fo_methodContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#fo_retries.
	VisitFo_retries(ctx *Fo_retriesContext) interface{}

	// Visit a parse tree produced by tnsnamesParser#fo_delay.
	VisitFo_delay(ctx *Fo_delayContext) interface{}
}
