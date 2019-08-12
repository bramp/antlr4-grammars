// Code generated from tnsnamesParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package tnsnames // tnsnamesParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetnsnamesParserListener is a complete listener for a parse tree produced by tnsnamesParser.
type BasetnsnamesParserListener struct{}

var _ tnsnamesParserListener = &BasetnsnamesParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetnsnamesParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetnsnamesParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetnsnamesParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetnsnamesParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTnsnames is called when production tnsnames is entered.
func (s *BasetnsnamesParserListener) EnterTnsnames(ctx *TnsnamesContext) {}

// ExitTnsnames is called when production tnsnames is exited.
func (s *BasetnsnamesParserListener) ExitTnsnames(ctx *TnsnamesContext) {}

// EnterTns_entry is called when production tns_entry is entered.
func (s *BasetnsnamesParserListener) EnterTns_entry(ctx *Tns_entryContext) {}

// ExitTns_entry is called when production tns_entry is exited.
func (s *BasetnsnamesParserListener) ExitTns_entry(ctx *Tns_entryContext) {}

// EnterIfile is called when production ifile is entered.
func (s *BasetnsnamesParserListener) EnterIfile(ctx *IfileContext) {}

// ExitIfile is called when production ifile is exited.
func (s *BasetnsnamesParserListener) ExitIfile(ctx *IfileContext) {}

// EnterLsnr_entry is called when production lsnr_entry is entered.
func (s *BasetnsnamesParserListener) EnterLsnr_entry(ctx *Lsnr_entryContext) {}

// ExitLsnr_entry is called when production lsnr_entry is exited.
func (s *BasetnsnamesParserListener) ExitLsnr_entry(ctx *Lsnr_entryContext) {}

// EnterLsnr_description is called when production lsnr_description is entered.
func (s *BasetnsnamesParserListener) EnterLsnr_description(ctx *Lsnr_descriptionContext) {}

// ExitLsnr_description is called when production lsnr_description is exited.
func (s *BasetnsnamesParserListener) ExitLsnr_description(ctx *Lsnr_descriptionContext) {}

// EnterAlias_list is called when production alias_list is entered.
func (s *BasetnsnamesParserListener) EnterAlias_list(ctx *Alias_listContext) {}

// ExitAlias_list is called when production alias_list is exited.
func (s *BasetnsnamesParserListener) ExitAlias_list(ctx *Alias_listContext) {}

// EnterAlias is called when production alias is entered.
func (s *BasetnsnamesParserListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BasetnsnamesParserListener) ExitAlias(ctx *AliasContext) {}

// EnterDescription_list is called when production description_list is entered.
func (s *BasetnsnamesParserListener) EnterDescription_list(ctx *Description_listContext) {}

// ExitDescription_list is called when production description_list is exited.
func (s *BasetnsnamesParserListener) ExitDescription_list(ctx *Description_listContext) {}

// EnterDl_params is called when production dl_params is entered.
func (s *BasetnsnamesParserListener) EnterDl_params(ctx *Dl_paramsContext) {}

// ExitDl_params is called when production dl_params is exited.
func (s *BasetnsnamesParserListener) ExitDl_params(ctx *Dl_paramsContext) {}

// EnterDl_parameter is called when production dl_parameter is entered.
func (s *BasetnsnamesParserListener) EnterDl_parameter(ctx *Dl_parameterContext) {}

// ExitDl_parameter is called when production dl_parameter is exited.
func (s *BasetnsnamesParserListener) ExitDl_parameter(ctx *Dl_parameterContext) {}

// EnterDescription is called when production description is entered.
func (s *BasetnsnamesParserListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BasetnsnamesParserListener) ExitDescription(ctx *DescriptionContext) {}

// EnterD_params is called when production d_params is entered.
func (s *BasetnsnamesParserListener) EnterD_params(ctx *D_paramsContext) {}

// ExitD_params is called when production d_params is exited.
func (s *BasetnsnamesParserListener) ExitD_params(ctx *D_paramsContext) {}

// EnterD_parameter is called when production d_parameter is entered.
func (s *BasetnsnamesParserListener) EnterD_parameter(ctx *D_parameterContext) {}

// ExitD_parameter is called when production d_parameter is exited.
func (s *BasetnsnamesParserListener) ExitD_parameter(ctx *D_parameterContext) {}

// EnterD_enable is called when production d_enable is entered.
func (s *BasetnsnamesParserListener) EnterD_enable(ctx *D_enableContext) {}

// ExitD_enable is called when production d_enable is exited.
func (s *BasetnsnamesParserListener) ExitD_enable(ctx *D_enableContext) {}

// EnterD_sdu is called when production d_sdu is entered.
func (s *BasetnsnamesParserListener) EnterD_sdu(ctx *D_sduContext) {}

// ExitD_sdu is called when production d_sdu is exited.
func (s *BasetnsnamesParserListener) ExitD_sdu(ctx *D_sduContext) {}

// EnterD_recv_buf is called when production d_recv_buf is entered.
func (s *BasetnsnamesParserListener) EnterD_recv_buf(ctx *D_recv_bufContext) {}

// ExitD_recv_buf is called when production d_recv_buf is exited.
func (s *BasetnsnamesParserListener) ExitD_recv_buf(ctx *D_recv_bufContext) {}

// EnterD_send_buf is called when production d_send_buf is entered.
func (s *BasetnsnamesParserListener) EnterD_send_buf(ctx *D_send_bufContext) {}

// ExitD_send_buf is called when production d_send_buf is exited.
func (s *BasetnsnamesParserListener) ExitD_send_buf(ctx *D_send_bufContext) {}

// EnterD_service_type is called when production d_service_type is entered.
func (s *BasetnsnamesParserListener) EnterD_service_type(ctx *D_service_typeContext) {}

// ExitD_service_type is called when production d_service_type is exited.
func (s *BasetnsnamesParserListener) ExitD_service_type(ctx *D_service_typeContext) {}

// EnterD_security is called when production d_security is entered.
func (s *BasetnsnamesParserListener) EnterD_security(ctx *D_securityContext) {}

// ExitD_security is called when production d_security is exited.
func (s *BasetnsnamesParserListener) ExitD_security(ctx *D_securityContext) {}

// EnterD_conn_timeout is called when production d_conn_timeout is entered.
func (s *BasetnsnamesParserListener) EnterD_conn_timeout(ctx *D_conn_timeoutContext) {}

// ExitD_conn_timeout is called when production d_conn_timeout is exited.
func (s *BasetnsnamesParserListener) ExitD_conn_timeout(ctx *D_conn_timeoutContext) {}

// EnterD_retry_count is called when production d_retry_count is entered.
func (s *BasetnsnamesParserListener) EnterD_retry_count(ctx *D_retry_countContext) {}

// ExitD_retry_count is called when production d_retry_count is exited.
func (s *BasetnsnamesParserListener) ExitD_retry_count(ctx *D_retry_countContext) {}

// EnterD_tct is called when production d_tct is entered.
func (s *BasetnsnamesParserListener) EnterD_tct(ctx *D_tctContext) {}

// ExitD_tct is called when production d_tct is exited.
func (s *BasetnsnamesParserListener) ExitD_tct(ctx *D_tctContext) {}

// EnterDs_parameter is called when production ds_parameter is entered.
func (s *BasetnsnamesParserListener) EnterDs_parameter(ctx *Ds_parameterContext) {}

// ExitDs_parameter is called when production ds_parameter is exited.
func (s *BasetnsnamesParserListener) ExitDs_parameter(ctx *Ds_parameterContext) {}

// EnterAddress_list is called when production address_list is entered.
func (s *BasetnsnamesParserListener) EnterAddress_list(ctx *Address_listContext) {}

// ExitAddress_list is called when production address_list is exited.
func (s *BasetnsnamesParserListener) ExitAddress_list(ctx *Address_listContext) {}

// EnterAl_params is called when production al_params is entered.
func (s *BasetnsnamesParserListener) EnterAl_params(ctx *Al_paramsContext) {}

// ExitAl_params is called when production al_params is exited.
func (s *BasetnsnamesParserListener) ExitAl_params(ctx *Al_paramsContext) {}

// EnterAl_parameter is called when production al_parameter is entered.
func (s *BasetnsnamesParserListener) EnterAl_parameter(ctx *Al_parameterContext) {}

// ExitAl_parameter is called when production al_parameter is exited.
func (s *BasetnsnamesParserListener) ExitAl_parameter(ctx *Al_parameterContext) {}

// EnterAl_failover is called when production al_failover is entered.
func (s *BasetnsnamesParserListener) EnterAl_failover(ctx *Al_failoverContext) {}

// ExitAl_failover is called when production al_failover is exited.
func (s *BasetnsnamesParserListener) ExitAl_failover(ctx *Al_failoverContext) {}

// EnterAl_load_balance is called when production al_load_balance is entered.
func (s *BasetnsnamesParserListener) EnterAl_load_balance(ctx *Al_load_balanceContext) {}

// ExitAl_load_balance is called when production al_load_balance is exited.
func (s *BasetnsnamesParserListener) ExitAl_load_balance(ctx *Al_load_balanceContext) {}

// EnterAl_source_route is called when production al_source_route is entered.
func (s *BasetnsnamesParserListener) EnterAl_source_route(ctx *Al_source_routeContext) {}

// ExitAl_source_route is called when production al_source_route is exited.
func (s *BasetnsnamesParserListener) ExitAl_source_route(ctx *Al_source_routeContext) {}

// EnterAddress is called when production address is entered.
func (s *BasetnsnamesParserListener) EnterAddress(ctx *AddressContext) {}

// ExitAddress is called when production address is exited.
func (s *BasetnsnamesParserListener) ExitAddress(ctx *AddressContext) {}

// EnterA_params is called when production a_params is entered.
func (s *BasetnsnamesParserListener) EnterA_params(ctx *A_paramsContext) {}

// ExitA_params is called when production a_params is exited.
func (s *BasetnsnamesParserListener) ExitA_params(ctx *A_paramsContext) {}

// EnterA_parameter is called when production a_parameter is entered.
func (s *BasetnsnamesParserListener) EnterA_parameter(ctx *A_parameterContext) {}

// ExitA_parameter is called when production a_parameter is exited.
func (s *BasetnsnamesParserListener) ExitA_parameter(ctx *A_parameterContext) {}

// EnterProtocol_info is called when production protocol_info is entered.
func (s *BasetnsnamesParserListener) EnterProtocol_info(ctx *Protocol_infoContext) {}

// ExitProtocol_info is called when production protocol_info is exited.
func (s *BasetnsnamesParserListener) ExitProtocol_info(ctx *Protocol_infoContext) {}

// EnterTcp_protocol is called when production tcp_protocol is entered.
func (s *BasetnsnamesParserListener) EnterTcp_protocol(ctx *Tcp_protocolContext) {}

// ExitTcp_protocol is called when production tcp_protocol is exited.
func (s *BasetnsnamesParserListener) ExitTcp_protocol(ctx *Tcp_protocolContext) {}

// EnterTcp_params is called when production tcp_params is entered.
func (s *BasetnsnamesParserListener) EnterTcp_params(ctx *Tcp_paramsContext) {}

// ExitTcp_params is called when production tcp_params is exited.
func (s *BasetnsnamesParserListener) ExitTcp_params(ctx *Tcp_paramsContext) {}

// EnterTcp_parameter is called when production tcp_parameter is entered.
func (s *BasetnsnamesParserListener) EnterTcp_parameter(ctx *Tcp_parameterContext) {}

// ExitTcp_parameter is called when production tcp_parameter is exited.
func (s *BasetnsnamesParserListener) ExitTcp_parameter(ctx *Tcp_parameterContext) {}

// EnterTcp_host is called when production tcp_host is entered.
func (s *BasetnsnamesParserListener) EnterTcp_host(ctx *Tcp_hostContext) {}

// ExitTcp_host is called when production tcp_host is exited.
func (s *BasetnsnamesParserListener) ExitTcp_host(ctx *Tcp_hostContext) {}

// EnterTcp_port is called when production tcp_port is entered.
func (s *BasetnsnamesParserListener) EnterTcp_port(ctx *Tcp_portContext) {}

// ExitTcp_port is called when production tcp_port is exited.
func (s *BasetnsnamesParserListener) ExitTcp_port(ctx *Tcp_portContext) {}

// EnterTcp_tcp is called when production tcp_tcp is entered.
func (s *BasetnsnamesParserListener) EnterTcp_tcp(ctx *Tcp_tcpContext) {}

// ExitTcp_tcp is called when production tcp_tcp is exited.
func (s *BasetnsnamesParserListener) ExitTcp_tcp(ctx *Tcp_tcpContext) {}

// EnterHost is called when production host is entered.
func (s *BasetnsnamesParserListener) EnterHost(ctx *HostContext) {}

// ExitHost is called when production host is exited.
func (s *BasetnsnamesParserListener) ExitHost(ctx *HostContext) {}

// EnterPort is called when production port is entered.
func (s *BasetnsnamesParserListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BasetnsnamesParserListener) ExitPort(ctx *PortContext) {}

// EnterIpc_protocol is called when production ipc_protocol is entered.
func (s *BasetnsnamesParserListener) EnterIpc_protocol(ctx *Ipc_protocolContext) {}

// ExitIpc_protocol is called when production ipc_protocol is exited.
func (s *BasetnsnamesParserListener) ExitIpc_protocol(ctx *Ipc_protocolContext) {}

// EnterIpc_params is called when production ipc_params is entered.
func (s *BasetnsnamesParserListener) EnterIpc_params(ctx *Ipc_paramsContext) {}

// ExitIpc_params is called when production ipc_params is exited.
func (s *BasetnsnamesParserListener) ExitIpc_params(ctx *Ipc_paramsContext) {}

// EnterIpc_parameter is called when production ipc_parameter is entered.
func (s *BasetnsnamesParserListener) EnterIpc_parameter(ctx *Ipc_parameterContext) {}

// ExitIpc_parameter is called when production ipc_parameter is exited.
func (s *BasetnsnamesParserListener) ExitIpc_parameter(ctx *Ipc_parameterContext) {}

// EnterIpc_ipc is called when production ipc_ipc is entered.
func (s *BasetnsnamesParserListener) EnterIpc_ipc(ctx *Ipc_ipcContext) {}

// ExitIpc_ipc is called when production ipc_ipc is exited.
func (s *BasetnsnamesParserListener) ExitIpc_ipc(ctx *Ipc_ipcContext) {}

// EnterIpc_key is called when production ipc_key is entered.
func (s *BasetnsnamesParserListener) EnterIpc_key(ctx *Ipc_keyContext) {}

// ExitIpc_key is called when production ipc_key is exited.
func (s *BasetnsnamesParserListener) ExitIpc_key(ctx *Ipc_keyContext) {}

// EnterSpx_protocol is called when production spx_protocol is entered.
func (s *BasetnsnamesParserListener) EnterSpx_protocol(ctx *Spx_protocolContext) {}

// ExitSpx_protocol is called when production spx_protocol is exited.
func (s *BasetnsnamesParserListener) ExitSpx_protocol(ctx *Spx_protocolContext) {}

// EnterSpx_params is called when production spx_params is entered.
func (s *BasetnsnamesParserListener) EnterSpx_params(ctx *Spx_paramsContext) {}

// ExitSpx_params is called when production spx_params is exited.
func (s *BasetnsnamesParserListener) ExitSpx_params(ctx *Spx_paramsContext) {}

// EnterSpx_parameter is called when production spx_parameter is entered.
func (s *BasetnsnamesParserListener) EnterSpx_parameter(ctx *Spx_parameterContext) {}

// ExitSpx_parameter is called when production spx_parameter is exited.
func (s *BasetnsnamesParserListener) ExitSpx_parameter(ctx *Spx_parameterContext) {}

// EnterSpx_spx is called when production spx_spx is entered.
func (s *BasetnsnamesParserListener) EnterSpx_spx(ctx *Spx_spxContext) {}

// ExitSpx_spx is called when production spx_spx is exited.
func (s *BasetnsnamesParserListener) ExitSpx_spx(ctx *Spx_spxContext) {}

// EnterSpx_service is called when production spx_service is entered.
func (s *BasetnsnamesParserListener) EnterSpx_service(ctx *Spx_serviceContext) {}

// ExitSpx_service is called when production spx_service is exited.
func (s *BasetnsnamesParserListener) ExitSpx_service(ctx *Spx_serviceContext) {}

// EnterNmp_protocol is called when production nmp_protocol is entered.
func (s *BasetnsnamesParserListener) EnterNmp_protocol(ctx *Nmp_protocolContext) {}

// ExitNmp_protocol is called when production nmp_protocol is exited.
func (s *BasetnsnamesParserListener) ExitNmp_protocol(ctx *Nmp_protocolContext) {}

// EnterNmp_params is called when production nmp_params is entered.
func (s *BasetnsnamesParserListener) EnterNmp_params(ctx *Nmp_paramsContext) {}

// ExitNmp_params is called when production nmp_params is exited.
func (s *BasetnsnamesParserListener) ExitNmp_params(ctx *Nmp_paramsContext) {}

// EnterNmp_parameter is called when production nmp_parameter is entered.
func (s *BasetnsnamesParserListener) EnterNmp_parameter(ctx *Nmp_parameterContext) {}

// ExitNmp_parameter is called when production nmp_parameter is exited.
func (s *BasetnsnamesParserListener) ExitNmp_parameter(ctx *Nmp_parameterContext) {}

// EnterNmp_nmp is called when production nmp_nmp is entered.
func (s *BasetnsnamesParserListener) EnterNmp_nmp(ctx *Nmp_nmpContext) {}

// ExitNmp_nmp is called when production nmp_nmp is exited.
func (s *BasetnsnamesParserListener) ExitNmp_nmp(ctx *Nmp_nmpContext) {}

// EnterNmp_server is called when production nmp_server is entered.
func (s *BasetnsnamesParserListener) EnterNmp_server(ctx *Nmp_serverContext) {}

// ExitNmp_server is called when production nmp_server is exited.
func (s *BasetnsnamesParserListener) ExitNmp_server(ctx *Nmp_serverContext) {}

// EnterNmp_pipe is called when production nmp_pipe is entered.
func (s *BasetnsnamesParserListener) EnterNmp_pipe(ctx *Nmp_pipeContext) {}

// ExitNmp_pipe is called when production nmp_pipe is exited.
func (s *BasetnsnamesParserListener) ExitNmp_pipe(ctx *Nmp_pipeContext) {}

// EnterBeq_protocol is called when production beq_protocol is entered.
func (s *BasetnsnamesParserListener) EnterBeq_protocol(ctx *Beq_protocolContext) {}

// ExitBeq_protocol is called when production beq_protocol is exited.
func (s *BasetnsnamesParserListener) ExitBeq_protocol(ctx *Beq_protocolContext) {}

// EnterBeq_params is called when production beq_params is entered.
func (s *BasetnsnamesParserListener) EnterBeq_params(ctx *Beq_paramsContext) {}

// ExitBeq_params is called when production beq_params is exited.
func (s *BasetnsnamesParserListener) ExitBeq_params(ctx *Beq_paramsContext) {}

// EnterBeq_parameter is called when production beq_parameter is entered.
func (s *BasetnsnamesParserListener) EnterBeq_parameter(ctx *Beq_parameterContext) {}

// ExitBeq_parameter is called when production beq_parameter is exited.
func (s *BasetnsnamesParserListener) ExitBeq_parameter(ctx *Beq_parameterContext) {}

// EnterBeq_beq is called when production beq_beq is entered.
func (s *BasetnsnamesParserListener) EnterBeq_beq(ctx *Beq_beqContext) {}

// ExitBeq_beq is called when production beq_beq is exited.
func (s *BasetnsnamesParserListener) ExitBeq_beq(ctx *Beq_beqContext) {}

// EnterBeq_program is called when production beq_program is entered.
func (s *BasetnsnamesParserListener) EnterBeq_program(ctx *Beq_programContext) {}

// ExitBeq_program is called when production beq_program is exited.
func (s *BasetnsnamesParserListener) ExitBeq_program(ctx *Beq_programContext) {}

// EnterBeq_argv0 is called when production beq_argv0 is entered.
func (s *BasetnsnamesParserListener) EnterBeq_argv0(ctx *Beq_argv0Context) {}

// ExitBeq_argv0 is called when production beq_argv0 is exited.
func (s *BasetnsnamesParserListener) ExitBeq_argv0(ctx *Beq_argv0Context) {}

// EnterBeq_args is called when production beq_args is entered.
func (s *BasetnsnamesParserListener) EnterBeq_args(ctx *Beq_argsContext) {}

// ExitBeq_args is called when production beq_args is exited.
func (s *BasetnsnamesParserListener) ExitBeq_args(ctx *Beq_argsContext) {}

// EnterBa_parameter is called when production ba_parameter is entered.
func (s *BasetnsnamesParserListener) EnterBa_parameter(ctx *Ba_parameterContext) {}

// ExitBa_parameter is called when production ba_parameter is exited.
func (s *BasetnsnamesParserListener) ExitBa_parameter(ctx *Ba_parameterContext) {}

// EnterBa_description is called when production ba_description is entered.
func (s *BasetnsnamesParserListener) EnterBa_description(ctx *Ba_descriptionContext) {}

// ExitBa_description is called when production ba_description is exited.
func (s *BasetnsnamesParserListener) ExitBa_description(ctx *Ba_descriptionContext) {}

// EnterBad_params is called when production bad_params is entered.
func (s *BasetnsnamesParserListener) EnterBad_params(ctx *Bad_paramsContext) {}

// ExitBad_params is called when production bad_params is exited.
func (s *BasetnsnamesParserListener) ExitBad_params(ctx *Bad_paramsContext) {}

// EnterBad_parameter is called when production bad_parameter is entered.
func (s *BasetnsnamesParserListener) EnterBad_parameter(ctx *Bad_parameterContext) {}

// ExitBad_parameter is called when production bad_parameter is exited.
func (s *BasetnsnamesParserListener) ExitBad_parameter(ctx *Bad_parameterContext) {}

// EnterBad_local is called when production bad_local is entered.
func (s *BasetnsnamesParserListener) EnterBad_local(ctx *Bad_localContext) {}

// ExitBad_local is called when production bad_local is exited.
func (s *BasetnsnamesParserListener) ExitBad_local(ctx *Bad_localContext) {}

// EnterBad_address is called when production bad_address is entered.
func (s *BasetnsnamesParserListener) EnterBad_address(ctx *Bad_addressContext) {}

// ExitBad_address is called when production bad_address is exited.
func (s *BasetnsnamesParserListener) ExitBad_address(ctx *Bad_addressContext) {}

// EnterConnect_data is called when production connect_data is entered.
func (s *BasetnsnamesParserListener) EnterConnect_data(ctx *Connect_dataContext) {}

// ExitConnect_data is called when production connect_data is exited.
func (s *BasetnsnamesParserListener) ExitConnect_data(ctx *Connect_dataContext) {}

// EnterCd_params is called when production cd_params is entered.
func (s *BasetnsnamesParserListener) EnterCd_params(ctx *Cd_paramsContext) {}

// ExitCd_params is called when production cd_params is exited.
func (s *BasetnsnamesParserListener) ExitCd_params(ctx *Cd_paramsContext) {}

// EnterCd_parameter is called when production cd_parameter is entered.
func (s *BasetnsnamesParserListener) EnterCd_parameter(ctx *Cd_parameterContext) {}

// ExitCd_parameter is called when production cd_parameter is exited.
func (s *BasetnsnamesParserListener) ExitCd_parameter(ctx *Cd_parameterContext) {}

// EnterCd_service_name is called when production cd_service_name is entered.
func (s *BasetnsnamesParserListener) EnterCd_service_name(ctx *Cd_service_nameContext) {}

// ExitCd_service_name is called when production cd_service_name is exited.
func (s *BasetnsnamesParserListener) ExitCd_service_name(ctx *Cd_service_nameContext) {}

// EnterCd_sid is called when production cd_sid is entered.
func (s *BasetnsnamesParserListener) EnterCd_sid(ctx *Cd_sidContext) {}

// ExitCd_sid is called when production cd_sid is exited.
func (s *BasetnsnamesParserListener) ExitCd_sid(ctx *Cd_sidContext) {}

// EnterCd_instance_name is called when production cd_instance_name is entered.
func (s *BasetnsnamesParserListener) EnterCd_instance_name(ctx *Cd_instance_nameContext) {}

// ExitCd_instance_name is called when production cd_instance_name is exited.
func (s *BasetnsnamesParserListener) ExitCd_instance_name(ctx *Cd_instance_nameContext) {}

// EnterCd_failover_mode is called when production cd_failover_mode is entered.
func (s *BasetnsnamesParserListener) EnterCd_failover_mode(ctx *Cd_failover_modeContext) {}

// ExitCd_failover_mode is called when production cd_failover_mode is exited.
func (s *BasetnsnamesParserListener) ExitCd_failover_mode(ctx *Cd_failover_modeContext) {}

// EnterCd_global_name is called when production cd_global_name is entered.
func (s *BasetnsnamesParserListener) EnterCd_global_name(ctx *Cd_global_nameContext) {}

// ExitCd_global_name is called when production cd_global_name is exited.
func (s *BasetnsnamesParserListener) ExitCd_global_name(ctx *Cd_global_nameContext) {}

// EnterCd_hs is called when production cd_hs is entered.
func (s *BasetnsnamesParserListener) EnterCd_hs(ctx *Cd_hsContext) {}

// ExitCd_hs is called when production cd_hs is exited.
func (s *BasetnsnamesParserListener) ExitCd_hs(ctx *Cd_hsContext) {}

// EnterCd_rdb_database is called when production cd_rdb_database is entered.
func (s *BasetnsnamesParserListener) EnterCd_rdb_database(ctx *Cd_rdb_databaseContext) {}

// ExitCd_rdb_database is called when production cd_rdb_database is exited.
func (s *BasetnsnamesParserListener) ExitCd_rdb_database(ctx *Cd_rdb_databaseContext) {}

// EnterCd_server is called when production cd_server is entered.
func (s *BasetnsnamesParserListener) EnterCd_server(ctx *Cd_serverContext) {}

// ExitCd_server is called when production cd_server is exited.
func (s *BasetnsnamesParserListener) ExitCd_server(ctx *Cd_serverContext) {}

// EnterCd_ur is called when production cd_ur is entered.
func (s *BasetnsnamesParserListener) EnterCd_ur(ctx *Cd_urContext) {}

// ExitCd_ur is called when production cd_ur is exited.
func (s *BasetnsnamesParserListener) ExitCd_ur(ctx *Cd_urContext) {}

// EnterFo_params is called when production fo_params is entered.
func (s *BasetnsnamesParserListener) EnterFo_params(ctx *Fo_paramsContext) {}

// ExitFo_params is called when production fo_params is exited.
func (s *BasetnsnamesParserListener) ExitFo_params(ctx *Fo_paramsContext) {}

// EnterFo_parameter is called when production fo_parameter is entered.
func (s *BasetnsnamesParserListener) EnterFo_parameter(ctx *Fo_parameterContext) {}

// ExitFo_parameter is called when production fo_parameter is exited.
func (s *BasetnsnamesParserListener) ExitFo_parameter(ctx *Fo_parameterContext) {}

// EnterFo_type is called when production fo_type is entered.
func (s *BasetnsnamesParserListener) EnterFo_type(ctx *Fo_typeContext) {}

// ExitFo_type is called when production fo_type is exited.
func (s *BasetnsnamesParserListener) ExitFo_type(ctx *Fo_typeContext) {}

// EnterFo_backup is called when production fo_backup is entered.
func (s *BasetnsnamesParserListener) EnterFo_backup(ctx *Fo_backupContext) {}

// ExitFo_backup is called when production fo_backup is exited.
func (s *BasetnsnamesParserListener) ExitFo_backup(ctx *Fo_backupContext) {}

// EnterFo_method is called when production fo_method is entered.
func (s *BasetnsnamesParserListener) EnterFo_method(ctx *Fo_methodContext) {}

// ExitFo_method is called when production fo_method is exited.
func (s *BasetnsnamesParserListener) ExitFo_method(ctx *Fo_methodContext) {}

// EnterFo_retries is called when production fo_retries is entered.
func (s *BasetnsnamesParserListener) EnterFo_retries(ctx *Fo_retriesContext) {}

// ExitFo_retries is called when production fo_retries is exited.
func (s *BasetnsnamesParserListener) ExitFo_retries(ctx *Fo_retriesContext) {}

// EnterFo_delay is called when production fo_delay is entered.
func (s *BasetnsnamesParserListener) EnterFo_delay(ctx *Fo_delayContext) {}

// ExitFo_delay is called when production fo_delay is exited.
func (s *BasetnsnamesParserListener) ExitFo_delay(ctx *Fo_delayContext) {}
