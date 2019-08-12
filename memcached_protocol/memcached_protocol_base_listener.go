// Code generated from memcached_protocol.g4 by ANTLR 4.7.2. DO NOT EDIT.

package memcached_protocol // memcached_protocol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Basememcached_protocolListener is a complete listener for a parse tree produced by memcached_protocolParser.
type Basememcached_protocolListener struct{}

var _ memcached_protocolListener = &Basememcached_protocolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basememcached_protocolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basememcached_protocolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basememcached_protocolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basememcached_protocolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCommand_line is called when production command_line is entered.
func (s *Basememcached_protocolListener) EnterCommand_line(ctx *Command_lineContext) {}

// ExitCommand_line is called when production command_line is exited.
func (s *Basememcached_protocolListener) ExitCommand_line(ctx *Command_lineContext) {}

// EnterStorage_command is called when production storage_command is entered.
func (s *Basememcached_protocolListener) EnterStorage_command(ctx *Storage_commandContext) {}

// ExitStorage_command is called when production storage_command is exited.
func (s *Basememcached_protocolListener) ExitStorage_command(ctx *Storage_commandContext) {}

// EnterStorage_command_name is called when production storage_command_name is entered.
func (s *Basememcached_protocolListener) EnterStorage_command_name(ctx *Storage_command_nameContext) {}

// ExitStorage_command_name is called when production storage_command_name is exited.
func (s *Basememcached_protocolListener) ExitStorage_command_name(ctx *Storage_command_nameContext) {}

// EnterRetrieval_command is called when production retrieval_command is entered.
func (s *Basememcached_protocolListener) EnterRetrieval_command(ctx *Retrieval_commandContext) {}

// ExitRetrieval_command is called when production retrieval_command is exited.
func (s *Basememcached_protocolListener) ExitRetrieval_command(ctx *Retrieval_commandContext) {}

// EnterDelete_command is called when production delete_command is entered.
func (s *Basememcached_protocolListener) EnterDelete_command(ctx *Delete_commandContext) {}

// ExitDelete_command is called when production delete_command is exited.
func (s *Basememcached_protocolListener) ExitDelete_command(ctx *Delete_commandContext) {}

// EnterIncrement_command is called when production increment_command is entered.
func (s *Basememcached_protocolListener) EnterIncrement_command(ctx *Increment_commandContext) {}

// ExitIncrement_command is called when production increment_command is exited.
func (s *Basememcached_protocolListener) ExitIncrement_command(ctx *Increment_commandContext) {}

// EnterDecrement_command is called when production decrement_command is entered.
func (s *Basememcached_protocolListener) EnterDecrement_command(ctx *Decrement_commandContext) {}

// ExitDecrement_command is called when production decrement_command is exited.
func (s *Basememcached_protocolListener) ExitDecrement_command(ctx *Decrement_commandContext) {}

// EnterStatistics_command is called when production statistics_command is entered.
func (s *Basememcached_protocolListener) EnterStatistics_command(ctx *Statistics_commandContext) {}

// ExitStatistics_command is called when production statistics_command is exited.
func (s *Basememcached_protocolListener) ExitStatistics_command(ctx *Statistics_commandContext) {}

// EnterStatistics_option is called when production statistics_option is entered.
func (s *Basememcached_protocolListener) EnterStatistics_option(ctx *Statistics_optionContext) {}

// ExitStatistics_option is called when production statistics_option is exited.
func (s *Basememcached_protocolListener) ExitStatistics_option(ctx *Statistics_optionContext) {}

// EnterFlush_command is called when production flush_command is entered.
func (s *Basememcached_protocolListener) EnterFlush_command(ctx *Flush_commandContext) {}

// ExitFlush_command is called when production flush_command is exited.
func (s *Basememcached_protocolListener) ExitFlush_command(ctx *Flush_commandContext) {}

// EnterVersion_command is called when production version_command is entered.
func (s *Basememcached_protocolListener) EnterVersion_command(ctx *Version_commandContext) {}

// ExitVersion_command is called when production version_command is exited.
func (s *Basememcached_protocolListener) ExitVersion_command(ctx *Version_commandContext) {}

// EnterVerbosity_command is called when production verbosity_command is entered.
func (s *Basememcached_protocolListener) EnterVerbosity_command(ctx *Verbosity_commandContext) {}

// ExitVerbosity_command is called when production verbosity_command is exited.
func (s *Basememcached_protocolListener) ExitVerbosity_command(ctx *Verbosity_commandContext) {}

// EnterQuit_command is called when production quit_command is entered.
func (s *Basememcached_protocolListener) EnterQuit_command(ctx *Quit_commandContext) {}

// ExitQuit_command is called when production quit_command is exited.
func (s *Basememcached_protocolListener) ExitQuit_command(ctx *Quit_commandContext) {}

// EnterStorage_response is called when production storage_response is entered.
func (s *Basememcached_protocolListener) EnterStorage_response(ctx *Storage_responseContext) {}

// ExitStorage_response is called when production storage_response is exited.
func (s *Basememcached_protocolListener) ExitStorage_response(ctx *Storage_responseContext) {}

// EnterRetrieval_response is called when production retrieval_response is entered.
func (s *Basememcached_protocolListener) EnterRetrieval_response(ctx *Retrieval_responseContext) {}

// ExitRetrieval_response is called when production retrieval_response is exited.
func (s *Basememcached_protocolListener) ExitRetrieval_response(ctx *Retrieval_responseContext) {}

// EnterDeletion_response is called when production deletion_response is entered.
func (s *Basememcached_protocolListener) EnterDeletion_response(ctx *Deletion_responseContext) {}

// ExitDeletion_response is called when production deletion_response is exited.
func (s *Basememcached_protocolListener) ExitDeletion_response(ctx *Deletion_responseContext) {}

// EnterIncr_or_decr_response is called when production incr_or_decr_response is entered.
func (s *Basememcached_protocolListener) EnterIncr_or_decr_response(ctx *Incr_or_decr_responseContext) {
}

// ExitIncr_or_decr_response is called when production incr_or_decr_response is exited.
func (s *Basememcached_protocolListener) ExitIncr_or_decr_response(ctx *Incr_or_decr_responseContext) {
}

// EnterStatistics_response is called when production statistics_response is entered.
func (s *Basememcached_protocolListener) EnterStatistics_response(ctx *Statistics_responseContext) {}

// ExitStatistics_response is called when production statistics_response is exited.
func (s *Basememcached_protocolListener) ExitStatistics_response(ctx *Statistics_responseContext) {}

// EnterError_response is called when production error_response is entered.
func (s *Basememcached_protocolListener) EnterError_response(ctx *Error_responseContext) {}

// ExitError_response is called when production error_response is exited.
func (s *Basememcached_protocolListener) ExitError_response(ctx *Error_responseContext) {}

// EnterGeneral_statistic is called when production general_statistic is entered.
func (s *Basememcached_protocolListener) EnterGeneral_statistic(ctx *General_statisticContext) {}

// ExitGeneral_statistic is called when production general_statistic is exited.
func (s *Basememcached_protocolListener) ExitGeneral_statistic(ctx *General_statisticContext) {}

// EnterSize_statistic is called when production size_statistic is entered.
func (s *Basememcached_protocolListener) EnterSize_statistic(ctx *Size_statisticContext) {}

// ExitSize_statistic is called when production size_statistic is exited.
func (s *Basememcached_protocolListener) ExitSize_statistic(ctx *Size_statisticContext) {}

// EnterGeneral_error is called when production general_error is entered.
func (s *Basememcached_protocolListener) EnterGeneral_error(ctx *General_errorContext) {}

// ExitGeneral_error is called when production general_error is exited.
func (s *Basememcached_protocolListener) ExitGeneral_error(ctx *General_errorContext) {}

// EnterClient_error_message is called when production client_error_message is entered.
func (s *Basememcached_protocolListener) EnterClient_error_message(ctx *Client_error_messageContext) {}

// ExitClient_error_message is called when production client_error_message is exited.
func (s *Basememcached_protocolListener) ExitClient_error_message(ctx *Client_error_messageContext) {}

// EnterServer_error_message is called when production server_error_message is entered.
func (s *Basememcached_protocolListener) EnterServer_error_message(ctx *Server_error_messageContext) {}

// ExitServer_error_message is called when production server_error_message is exited.
func (s *Basememcached_protocolListener) ExitServer_error_message(ctx *Server_error_messageContext) {}

// EnterEnd is called when production end is entered.
func (s *Basememcached_protocolListener) EnterEnd(ctx *EndContext) {}

// ExitEnd is called when production end is exited.
func (s *Basememcached_protocolListener) ExitEnd(ctx *EndContext) {}

// EnterNoreply is called when production noreply is entered.
func (s *Basememcached_protocolListener) EnterNoreply(ctx *NoreplyContext) {}

// ExitNoreply is called when production noreply is exited.
func (s *Basememcached_protocolListener) ExitNoreply(ctx *NoreplyContext) {}

// EnterKey is called when production key is entered.
func (s *Basememcached_protocolListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *Basememcached_protocolListener) ExitKey(ctx *KeyContext) {}

// EnterFlags is called when production flags is entered.
func (s *Basememcached_protocolListener) EnterFlags(ctx *FlagsContext) {}

// ExitFlags is called when production flags is exited.
func (s *Basememcached_protocolListener) ExitFlags(ctx *FlagsContext) {}

// EnterExptime is called when production exptime is entered.
func (s *Basememcached_protocolListener) EnterExptime(ctx *ExptimeContext) {}

// ExitExptime is called when production exptime is exited.
func (s *Basememcached_protocolListener) ExitExptime(ctx *ExptimeContext) {}

// EnterBytes is called when production bytes is entered.
func (s *Basememcached_protocolListener) EnterBytes(ctx *BytesContext) {}

// ExitBytes is called when production bytes is exited.
func (s *Basememcached_protocolListener) ExitBytes(ctx *BytesContext) {}

// EnterCas_unique is called when production cas_unique is entered.
func (s *Basememcached_protocolListener) EnterCas_unique(ctx *Cas_uniqueContext) {}

// ExitCas_unique is called when production cas_unique is exited.
func (s *Basememcached_protocolListener) ExitCas_unique(ctx *Cas_uniqueContext) {}

// EnterValue is called when production value is entered.
func (s *Basememcached_protocolListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *Basememcached_protocolListener) ExitValue(ctx *ValueContext) {}

// EnterTime is called when production time is entered.
func (s *Basememcached_protocolListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *Basememcached_protocolListener) ExitTime(ctx *TimeContext) {}

// EnterDelay is called when production delay is entered.
func (s *Basememcached_protocolListener) EnterDelay(ctx *DelayContext) {}

// ExitDelay is called when production delay is exited.
func (s *Basememcached_protocolListener) ExitDelay(ctx *DelayContext) {}

// EnterVerbosity_level is called when production verbosity_level is entered.
func (s *Basememcached_protocolListener) EnterVerbosity_level(ctx *Verbosity_levelContext) {}

// ExitVerbosity_level is called when production verbosity_level is exited.
func (s *Basememcached_protocolListener) ExitVerbosity_level(ctx *Verbosity_levelContext) {}

// EnterStatistic_name is called when production statistic_name is entered.
func (s *Basememcached_protocolListener) EnterStatistic_name(ctx *Statistic_nameContext) {}

// ExitStatistic_name is called when production statistic_name is exited.
func (s *Basememcached_protocolListener) ExitStatistic_name(ctx *Statistic_nameContext) {}

// EnterStatistic_value is called when production statistic_value is entered.
func (s *Basememcached_protocolListener) EnterStatistic_value(ctx *Statistic_valueContext) {}

// ExitStatistic_value is called when production statistic_value is exited.
func (s *Basememcached_protocolListener) ExitStatistic_value(ctx *Statistic_valueContext) {}

// EnterSize is called when production size is entered.
func (s *Basememcached_protocolListener) EnterSize(ctx *SizeContext) {}

// ExitSize is called when production size is exited.
func (s *Basememcached_protocolListener) ExitSize(ctx *SizeContext) {}

// EnterCount is called when production count is entered.
func (s *Basememcached_protocolListener) EnterCount(ctx *CountContext) {}

// ExitCount is called when production count is exited.
func (s *Basememcached_protocolListener) ExitCount(ctx *CountContext) {}
