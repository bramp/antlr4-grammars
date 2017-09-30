// Generated from memcached_protocol.g4 by ANTLR 4.7.

package memcached_protocol // memcached_protocol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by memcached_protocolParser.
type memcached_protocolVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by memcached_protocolParser#command_line.
	VisitCommand_line(ctx *Command_lineContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#storage_command.
	VisitStorage_command(ctx *Storage_commandContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#storage_command_name.
	VisitStorage_command_name(ctx *Storage_command_nameContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#retrieval_command.
	VisitRetrieval_command(ctx *Retrieval_commandContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#delete_command.
	VisitDelete_command(ctx *Delete_commandContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#increment_command.
	VisitIncrement_command(ctx *Increment_commandContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#decrement_command.
	VisitDecrement_command(ctx *Decrement_commandContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#statistics_command.
	VisitStatistics_command(ctx *Statistics_commandContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#statistics_option.
	VisitStatistics_option(ctx *Statistics_optionContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#flush_command.
	VisitFlush_command(ctx *Flush_commandContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#version_command.
	VisitVersion_command(ctx *Version_commandContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#verbosity_command.
	VisitVerbosity_command(ctx *Verbosity_commandContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#quit_command.
	VisitQuit_command(ctx *Quit_commandContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#storage_response.
	VisitStorage_response(ctx *Storage_responseContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#retrieval_response.
	VisitRetrieval_response(ctx *Retrieval_responseContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#deletion_response.
	VisitDeletion_response(ctx *Deletion_responseContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#incr_or_decr_response.
	VisitIncr_or_decr_response(ctx *Incr_or_decr_responseContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#statistics_response.
	VisitStatistics_response(ctx *Statistics_responseContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#error_response.
	VisitError_response(ctx *Error_responseContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#general_statistic.
	VisitGeneral_statistic(ctx *General_statisticContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#size_statistic.
	VisitSize_statistic(ctx *Size_statisticContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#general_error.
	VisitGeneral_error(ctx *General_errorContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#client_error_message.
	VisitClient_error_message(ctx *Client_error_messageContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#server_error_message.
	VisitServer_error_message(ctx *Server_error_messageContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#end.
	VisitEnd(ctx *EndContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#noreply.
	VisitNoreply(ctx *NoreplyContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#flags.
	VisitFlags(ctx *FlagsContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#exptime.
	VisitExptime(ctx *ExptimeContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#bytes.
	VisitBytes(ctx *BytesContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#cas_unique.
	VisitCas_unique(ctx *Cas_uniqueContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#time.
	VisitTime(ctx *TimeContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#delay.
	VisitDelay(ctx *DelayContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#verbosity_level.
	VisitVerbosity_level(ctx *Verbosity_levelContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#statistic_name.
	VisitStatistic_name(ctx *Statistic_nameContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#statistic_value.
	VisitStatistic_value(ctx *Statistic_valueContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#size.
	VisitSize(ctx *SizeContext) interface{}

	// Visit a parse tree produced by memcached_protocolParser#count.
	VisitCount(ctx *CountContext) interface{}
}
