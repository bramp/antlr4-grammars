// Generated from tnsnamesParser.g4 by ANTLR 4.7.

package tnsnames // tnsnamesParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// tnsnamesParserListener is a complete listener for a parse tree produced by tnsnamesParser.
type tnsnamesParserListener interface {
	antlr.ParseTreeListener

	// EnterTnsnames is called when entering the tnsnames production.
	EnterTnsnames(c *TnsnamesContext)

	// EnterTns_entry is called when entering the tns_entry production.
	EnterTns_entry(c *Tns_entryContext)

	// EnterIfile is called when entering the ifile production.
	EnterIfile(c *IfileContext)

	// EnterLsnr_entry is called when entering the lsnr_entry production.
	EnterLsnr_entry(c *Lsnr_entryContext)

	// EnterLsnr_description is called when entering the lsnr_description production.
	EnterLsnr_description(c *Lsnr_descriptionContext)

	// EnterAlias_list is called when entering the alias_list production.
	EnterAlias_list(c *Alias_listContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterDescription_list is called when entering the description_list production.
	EnterDescription_list(c *Description_listContext)

	// EnterDl_params is called when entering the dl_params production.
	EnterDl_params(c *Dl_paramsContext)

	// EnterDl_parameter is called when entering the dl_parameter production.
	EnterDl_parameter(c *Dl_parameterContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterD_params is called when entering the d_params production.
	EnterD_params(c *D_paramsContext)

	// EnterD_parameter is called when entering the d_parameter production.
	EnterD_parameter(c *D_parameterContext)

	// EnterD_enable is called when entering the d_enable production.
	EnterD_enable(c *D_enableContext)

	// EnterD_sdu is called when entering the d_sdu production.
	EnterD_sdu(c *D_sduContext)

	// EnterD_recv_buf is called when entering the d_recv_buf production.
	EnterD_recv_buf(c *D_recv_bufContext)

	// EnterD_send_buf is called when entering the d_send_buf production.
	EnterD_send_buf(c *D_send_bufContext)

	// EnterD_service_type is called when entering the d_service_type production.
	EnterD_service_type(c *D_service_typeContext)

	// EnterD_security is called when entering the d_security production.
	EnterD_security(c *D_securityContext)

	// EnterD_conn_timeout is called when entering the d_conn_timeout production.
	EnterD_conn_timeout(c *D_conn_timeoutContext)

	// EnterD_retry_count is called when entering the d_retry_count production.
	EnterD_retry_count(c *D_retry_countContext)

	// EnterD_tct is called when entering the d_tct production.
	EnterD_tct(c *D_tctContext)

	// EnterDs_parameter is called when entering the ds_parameter production.
	EnterDs_parameter(c *Ds_parameterContext)

	// EnterAddress_list is called when entering the address_list production.
	EnterAddress_list(c *Address_listContext)

	// EnterAl_params is called when entering the al_params production.
	EnterAl_params(c *Al_paramsContext)

	// EnterAl_parameter is called when entering the al_parameter production.
	EnterAl_parameter(c *Al_parameterContext)

	// EnterAl_failover is called when entering the al_failover production.
	EnterAl_failover(c *Al_failoverContext)

	// EnterAl_load_balance is called when entering the al_load_balance production.
	EnterAl_load_balance(c *Al_load_balanceContext)

	// EnterAl_source_route is called when entering the al_source_route production.
	EnterAl_source_route(c *Al_source_routeContext)

	// EnterAddress is called when entering the address production.
	EnterAddress(c *AddressContext)

	// EnterA_params is called when entering the a_params production.
	EnterA_params(c *A_paramsContext)

	// EnterA_parameter is called when entering the a_parameter production.
	EnterA_parameter(c *A_parameterContext)

	// EnterProtocol_info is called when entering the protocol_info production.
	EnterProtocol_info(c *Protocol_infoContext)

	// EnterTcp_protocol is called when entering the tcp_protocol production.
	EnterTcp_protocol(c *Tcp_protocolContext)

	// EnterTcp_params is called when entering the tcp_params production.
	EnterTcp_params(c *Tcp_paramsContext)

	// EnterTcp_parameter is called when entering the tcp_parameter production.
	EnterTcp_parameter(c *Tcp_parameterContext)

	// EnterTcp_host is called when entering the tcp_host production.
	EnterTcp_host(c *Tcp_hostContext)

	// EnterTcp_port is called when entering the tcp_port production.
	EnterTcp_port(c *Tcp_portContext)

	// EnterTcp_tcp is called when entering the tcp_tcp production.
	EnterTcp_tcp(c *Tcp_tcpContext)

	// EnterHost is called when entering the host production.
	EnterHost(c *HostContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterIpc_protocol is called when entering the ipc_protocol production.
	EnterIpc_protocol(c *Ipc_protocolContext)

	// EnterIpc_params is called when entering the ipc_params production.
	EnterIpc_params(c *Ipc_paramsContext)

	// EnterIpc_parameter is called when entering the ipc_parameter production.
	EnterIpc_parameter(c *Ipc_parameterContext)

	// EnterIpc_ipc is called when entering the ipc_ipc production.
	EnterIpc_ipc(c *Ipc_ipcContext)

	// EnterIpc_key is called when entering the ipc_key production.
	EnterIpc_key(c *Ipc_keyContext)

	// EnterSpx_protocol is called when entering the spx_protocol production.
	EnterSpx_protocol(c *Spx_protocolContext)

	// EnterSpx_params is called when entering the spx_params production.
	EnterSpx_params(c *Spx_paramsContext)

	// EnterSpx_parameter is called when entering the spx_parameter production.
	EnterSpx_parameter(c *Spx_parameterContext)

	// EnterSpx_spx is called when entering the spx_spx production.
	EnterSpx_spx(c *Spx_spxContext)

	// EnterSpx_service is called when entering the spx_service production.
	EnterSpx_service(c *Spx_serviceContext)

	// EnterNmp_protocol is called when entering the nmp_protocol production.
	EnterNmp_protocol(c *Nmp_protocolContext)

	// EnterNmp_params is called when entering the nmp_params production.
	EnterNmp_params(c *Nmp_paramsContext)

	// EnterNmp_parameter is called when entering the nmp_parameter production.
	EnterNmp_parameter(c *Nmp_parameterContext)

	// EnterNmp_nmp is called when entering the nmp_nmp production.
	EnterNmp_nmp(c *Nmp_nmpContext)

	// EnterNmp_server is called when entering the nmp_server production.
	EnterNmp_server(c *Nmp_serverContext)

	// EnterNmp_pipe is called when entering the nmp_pipe production.
	EnterNmp_pipe(c *Nmp_pipeContext)

	// EnterBeq_protocol is called when entering the beq_protocol production.
	EnterBeq_protocol(c *Beq_protocolContext)

	// EnterBeq_params is called when entering the beq_params production.
	EnterBeq_params(c *Beq_paramsContext)

	// EnterBeq_parameter is called when entering the beq_parameter production.
	EnterBeq_parameter(c *Beq_parameterContext)

	// EnterBeq_beq is called when entering the beq_beq production.
	EnterBeq_beq(c *Beq_beqContext)

	// EnterBeq_program is called when entering the beq_program production.
	EnterBeq_program(c *Beq_programContext)

	// EnterBeq_argv0 is called when entering the beq_argv0 production.
	EnterBeq_argv0(c *Beq_argv0Context)

	// EnterBeq_args is called when entering the beq_args production.
	EnterBeq_args(c *Beq_argsContext)

	// EnterBa_parameter is called when entering the ba_parameter production.
	EnterBa_parameter(c *Ba_parameterContext)

	// EnterBa_description is called when entering the ba_description production.
	EnterBa_description(c *Ba_descriptionContext)

	// EnterBad_params is called when entering the bad_params production.
	EnterBad_params(c *Bad_paramsContext)

	// EnterBad_parameter is called when entering the bad_parameter production.
	EnterBad_parameter(c *Bad_parameterContext)

	// EnterBad_local is called when entering the bad_local production.
	EnterBad_local(c *Bad_localContext)

	// EnterBad_address is called when entering the bad_address production.
	EnterBad_address(c *Bad_addressContext)

	// EnterConnect_data is called when entering the connect_data production.
	EnterConnect_data(c *Connect_dataContext)

	// EnterCd_params is called when entering the cd_params production.
	EnterCd_params(c *Cd_paramsContext)

	// EnterCd_parameter is called when entering the cd_parameter production.
	EnterCd_parameter(c *Cd_parameterContext)

	// EnterCd_service_name is called when entering the cd_service_name production.
	EnterCd_service_name(c *Cd_service_nameContext)

	// EnterCd_sid is called when entering the cd_sid production.
	EnterCd_sid(c *Cd_sidContext)

	// EnterCd_instance_name is called when entering the cd_instance_name production.
	EnterCd_instance_name(c *Cd_instance_nameContext)

	// EnterCd_failover_mode is called when entering the cd_failover_mode production.
	EnterCd_failover_mode(c *Cd_failover_modeContext)

	// EnterCd_global_name is called when entering the cd_global_name production.
	EnterCd_global_name(c *Cd_global_nameContext)

	// EnterCd_hs is called when entering the cd_hs production.
	EnterCd_hs(c *Cd_hsContext)

	// EnterCd_rdb_database is called when entering the cd_rdb_database production.
	EnterCd_rdb_database(c *Cd_rdb_databaseContext)

	// EnterCd_server is called when entering the cd_server production.
	EnterCd_server(c *Cd_serverContext)

	// EnterCd_ur is called when entering the cd_ur production.
	EnterCd_ur(c *Cd_urContext)

	// EnterFo_params is called when entering the fo_params production.
	EnterFo_params(c *Fo_paramsContext)

	// EnterFo_parameter is called when entering the fo_parameter production.
	EnterFo_parameter(c *Fo_parameterContext)

	// EnterFo_type is called when entering the fo_type production.
	EnterFo_type(c *Fo_typeContext)

	// EnterFo_backup is called when entering the fo_backup production.
	EnterFo_backup(c *Fo_backupContext)

	// EnterFo_method is called when entering the fo_method production.
	EnterFo_method(c *Fo_methodContext)

	// EnterFo_retries is called when entering the fo_retries production.
	EnterFo_retries(c *Fo_retriesContext)

	// EnterFo_delay is called when entering the fo_delay production.
	EnterFo_delay(c *Fo_delayContext)

	// ExitTnsnames is called when exiting the tnsnames production.
	ExitTnsnames(c *TnsnamesContext)

	// ExitTns_entry is called when exiting the tns_entry production.
	ExitTns_entry(c *Tns_entryContext)

	// ExitIfile is called when exiting the ifile production.
	ExitIfile(c *IfileContext)

	// ExitLsnr_entry is called when exiting the lsnr_entry production.
	ExitLsnr_entry(c *Lsnr_entryContext)

	// ExitLsnr_description is called when exiting the lsnr_description production.
	ExitLsnr_description(c *Lsnr_descriptionContext)

	// ExitAlias_list is called when exiting the alias_list production.
	ExitAlias_list(c *Alias_listContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitDescription_list is called when exiting the description_list production.
	ExitDescription_list(c *Description_listContext)

	// ExitDl_params is called when exiting the dl_params production.
	ExitDl_params(c *Dl_paramsContext)

	// ExitDl_parameter is called when exiting the dl_parameter production.
	ExitDl_parameter(c *Dl_parameterContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitD_params is called when exiting the d_params production.
	ExitD_params(c *D_paramsContext)

	// ExitD_parameter is called when exiting the d_parameter production.
	ExitD_parameter(c *D_parameterContext)

	// ExitD_enable is called when exiting the d_enable production.
	ExitD_enable(c *D_enableContext)

	// ExitD_sdu is called when exiting the d_sdu production.
	ExitD_sdu(c *D_sduContext)

	// ExitD_recv_buf is called when exiting the d_recv_buf production.
	ExitD_recv_buf(c *D_recv_bufContext)

	// ExitD_send_buf is called when exiting the d_send_buf production.
	ExitD_send_buf(c *D_send_bufContext)

	// ExitD_service_type is called when exiting the d_service_type production.
	ExitD_service_type(c *D_service_typeContext)

	// ExitD_security is called when exiting the d_security production.
	ExitD_security(c *D_securityContext)

	// ExitD_conn_timeout is called when exiting the d_conn_timeout production.
	ExitD_conn_timeout(c *D_conn_timeoutContext)

	// ExitD_retry_count is called when exiting the d_retry_count production.
	ExitD_retry_count(c *D_retry_countContext)

	// ExitD_tct is called when exiting the d_tct production.
	ExitD_tct(c *D_tctContext)

	// ExitDs_parameter is called when exiting the ds_parameter production.
	ExitDs_parameter(c *Ds_parameterContext)

	// ExitAddress_list is called when exiting the address_list production.
	ExitAddress_list(c *Address_listContext)

	// ExitAl_params is called when exiting the al_params production.
	ExitAl_params(c *Al_paramsContext)

	// ExitAl_parameter is called when exiting the al_parameter production.
	ExitAl_parameter(c *Al_parameterContext)

	// ExitAl_failover is called when exiting the al_failover production.
	ExitAl_failover(c *Al_failoverContext)

	// ExitAl_load_balance is called when exiting the al_load_balance production.
	ExitAl_load_balance(c *Al_load_balanceContext)

	// ExitAl_source_route is called when exiting the al_source_route production.
	ExitAl_source_route(c *Al_source_routeContext)

	// ExitAddress is called when exiting the address production.
	ExitAddress(c *AddressContext)

	// ExitA_params is called when exiting the a_params production.
	ExitA_params(c *A_paramsContext)

	// ExitA_parameter is called when exiting the a_parameter production.
	ExitA_parameter(c *A_parameterContext)

	// ExitProtocol_info is called when exiting the protocol_info production.
	ExitProtocol_info(c *Protocol_infoContext)

	// ExitTcp_protocol is called when exiting the tcp_protocol production.
	ExitTcp_protocol(c *Tcp_protocolContext)

	// ExitTcp_params is called when exiting the tcp_params production.
	ExitTcp_params(c *Tcp_paramsContext)

	// ExitTcp_parameter is called when exiting the tcp_parameter production.
	ExitTcp_parameter(c *Tcp_parameterContext)

	// ExitTcp_host is called when exiting the tcp_host production.
	ExitTcp_host(c *Tcp_hostContext)

	// ExitTcp_port is called when exiting the tcp_port production.
	ExitTcp_port(c *Tcp_portContext)

	// ExitTcp_tcp is called when exiting the tcp_tcp production.
	ExitTcp_tcp(c *Tcp_tcpContext)

	// ExitHost is called when exiting the host production.
	ExitHost(c *HostContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitIpc_protocol is called when exiting the ipc_protocol production.
	ExitIpc_protocol(c *Ipc_protocolContext)

	// ExitIpc_params is called when exiting the ipc_params production.
	ExitIpc_params(c *Ipc_paramsContext)

	// ExitIpc_parameter is called when exiting the ipc_parameter production.
	ExitIpc_parameter(c *Ipc_parameterContext)

	// ExitIpc_ipc is called when exiting the ipc_ipc production.
	ExitIpc_ipc(c *Ipc_ipcContext)

	// ExitIpc_key is called when exiting the ipc_key production.
	ExitIpc_key(c *Ipc_keyContext)

	// ExitSpx_protocol is called when exiting the spx_protocol production.
	ExitSpx_protocol(c *Spx_protocolContext)

	// ExitSpx_params is called when exiting the spx_params production.
	ExitSpx_params(c *Spx_paramsContext)

	// ExitSpx_parameter is called when exiting the spx_parameter production.
	ExitSpx_parameter(c *Spx_parameterContext)

	// ExitSpx_spx is called when exiting the spx_spx production.
	ExitSpx_spx(c *Spx_spxContext)

	// ExitSpx_service is called when exiting the spx_service production.
	ExitSpx_service(c *Spx_serviceContext)

	// ExitNmp_protocol is called when exiting the nmp_protocol production.
	ExitNmp_protocol(c *Nmp_protocolContext)

	// ExitNmp_params is called when exiting the nmp_params production.
	ExitNmp_params(c *Nmp_paramsContext)

	// ExitNmp_parameter is called when exiting the nmp_parameter production.
	ExitNmp_parameter(c *Nmp_parameterContext)

	// ExitNmp_nmp is called when exiting the nmp_nmp production.
	ExitNmp_nmp(c *Nmp_nmpContext)

	// ExitNmp_server is called when exiting the nmp_server production.
	ExitNmp_server(c *Nmp_serverContext)

	// ExitNmp_pipe is called when exiting the nmp_pipe production.
	ExitNmp_pipe(c *Nmp_pipeContext)

	// ExitBeq_protocol is called when exiting the beq_protocol production.
	ExitBeq_protocol(c *Beq_protocolContext)

	// ExitBeq_params is called when exiting the beq_params production.
	ExitBeq_params(c *Beq_paramsContext)

	// ExitBeq_parameter is called when exiting the beq_parameter production.
	ExitBeq_parameter(c *Beq_parameterContext)

	// ExitBeq_beq is called when exiting the beq_beq production.
	ExitBeq_beq(c *Beq_beqContext)

	// ExitBeq_program is called when exiting the beq_program production.
	ExitBeq_program(c *Beq_programContext)

	// ExitBeq_argv0 is called when exiting the beq_argv0 production.
	ExitBeq_argv0(c *Beq_argv0Context)

	// ExitBeq_args is called when exiting the beq_args production.
	ExitBeq_args(c *Beq_argsContext)

	// ExitBa_parameter is called when exiting the ba_parameter production.
	ExitBa_parameter(c *Ba_parameterContext)

	// ExitBa_description is called when exiting the ba_description production.
	ExitBa_description(c *Ba_descriptionContext)

	// ExitBad_params is called when exiting the bad_params production.
	ExitBad_params(c *Bad_paramsContext)

	// ExitBad_parameter is called when exiting the bad_parameter production.
	ExitBad_parameter(c *Bad_parameterContext)

	// ExitBad_local is called when exiting the bad_local production.
	ExitBad_local(c *Bad_localContext)

	// ExitBad_address is called when exiting the bad_address production.
	ExitBad_address(c *Bad_addressContext)

	// ExitConnect_data is called when exiting the connect_data production.
	ExitConnect_data(c *Connect_dataContext)

	// ExitCd_params is called when exiting the cd_params production.
	ExitCd_params(c *Cd_paramsContext)

	// ExitCd_parameter is called when exiting the cd_parameter production.
	ExitCd_parameter(c *Cd_parameterContext)

	// ExitCd_service_name is called when exiting the cd_service_name production.
	ExitCd_service_name(c *Cd_service_nameContext)

	// ExitCd_sid is called when exiting the cd_sid production.
	ExitCd_sid(c *Cd_sidContext)

	// ExitCd_instance_name is called when exiting the cd_instance_name production.
	ExitCd_instance_name(c *Cd_instance_nameContext)

	// ExitCd_failover_mode is called when exiting the cd_failover_mode production.
	ExitCd_failover_mode(c *Cd_failover_modeContext)

	// ExitCd_global_name is called when exiting the cd_global_name production.
	ExitCd_global_name(c *Cd_global_nameContext)

	// ExitCd_hs is called when exiting the cd_hs production.
	ExitCd_hs(c *Cd_hsContext)

	// ExitCd_rdb_database is called when exiting the cd_rdb_database production.
	ExitCd_rdb_database(c *Cd_rdb_databaseContext)

	// ExitCd_server is called when exiting the cd_server production.
	ExitCd_server(c *Cd_serverContext)

	// ExitCd_ur is called when exiting the cd_ur production.
	ExitCd_ur(c *Cd_urContext)

	// ExitFo_params is called when exiting the fo_params production.
	ExitFo_params(c *Fo_paramsContext)

	// ExitFo_parameter is called when exiting the fo_parameter production.
	ExitFo_parameter(c *Fo_parameterContext)

	// ExitFo_type is called when exiting the fo_type production.
	ExitFo_type(c *Fo_typeContext)

	// ExitFo_backup is called when exiting the fo_backup production.
	ExitFo_backup(c *Fo_backupContext)

	// ExitFo_method is called when exiting the fo_method production.
	ExitFo_method(c *Fo_methodContext)

	// ExitFo_retries is called when exiting the fo_retries production.
	ExitFo_retries(c *Fo_retriesContext)

	// ExitFo_delay is called when exiting the fo_delay production.
	ExitFo_delay(c *Fo_delayContext)
}
