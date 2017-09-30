// Generated from memcached_protocol.g4 by ANTLR 4.7.

package memcached_protocol // memcached_protocol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// memcached_protocolListener is a complete listener for a parse tree produced by memcached_protocolParser.
type memcached_protocolListener interface {
	antlr.ParseTreeListener

	// EnterCommand_line is called when entering the command_line production.
	EnterCommand_line(c *Command_lineContext)

	// EnterStorage_command is called when entering the storage_command production.
	EnterStorage_command(c *Storage_commandContext)

	// EnterStorage_command_name is called when entering the storage_command_name production.
	EnterStorage_command_name(c *Storage_command_nameContext)

	// EnterRetrieval_command is called when entering the retrieval_command production.
	EnterRetrieval_command(c *Retrieval_commandContext)

	// EnterDelete_command is called when entering the delete_command production.
	EnterDelete_command(c *Delete_commandContext)

	// EnterIncrement_command is called when entering the increment_command production.
	EnterIncrement_command(c *Increment_commandContext)

	// EnterDecrement_command is called when entering the decrement_command production.
	EnterDecrement_command(c *Decrement_commandContext)

	// EnterStatistics_command is called when entering the statistics_command production.
	EnterStatistics_command(c *Statistics_commandContext)

	// EnterStatistics_option is called when entering the statistics_option production.
	EnterStatistics_option(c *Statistics_optionContext)

	// EnterFlush_command is called when entering the flush_command production.
	EnterFlush_command(c *Flush_commandContext)

	// EnterVersion_command is called when entering the version_command production.
	EnterVersion_command(c *Version_commandContext)

	// EnterVerbosity_command is called when entering the verbosity_command production.
	EnterVerbosity_command(c *Verbosity_commandContext)

	// EnterQuit_command is called when entering the quit_command production.
	EnterQuit_command(c *Quit_commandContext)

	// EnterStorage_response is called when entering the storage_response production.
	EnterStorage_response(c *Storage_responseContext)

	// EnterRetrieval_response is called when entering the retrieval_response production.
	EnterRetrieval_response(c *Retrieval_responseContext)

	// EnterDeletion_response is called when entering the deletion_response production.
	EnterDeletion_response(c *Deletion_responseContext)

	// EnterIncr_or_decr_response is called when entering the incr_or_decr_response production.
	EnterIncr_or_decr_response(c *Incr_or_decr_responseContext)

	// EnterStatistics_response is called when entering the statistics_response production.
	EnterStatistics_response(c *Statistics_responseContext)

	// EnterError_response is called when entering the error_response production.
	EnterError_response(c *Error_responseContext)

	// EnterGeneral_statistic is called when entering the general_statistic production.
	EnterGeneral_statistic(c *General_statisticContext)

	// EnterSize_statistic is called when entering the size_statistic production.
	EnterSize_statistic(c *Size_statisticContext)

	// EnterGeneral_error is called when entering the general_error production.
	EnterGeneral_error(c *General_errorContext)

	// EnterClient_error_message is called when entering the client_error_message production.
	EnterClient_error_message(c *Client_error_messageContext)

	// EnterServer_error_message is called when entering the server_error_message production.
	EnterServer_error_message(c *Server_error_messageContext)

	// EnterEnd is called when entering the end production.
	EnterEnd(c *EndContext)

	// EnterNoreply is called when entering the noreply production.
	EnterNoreply(c *NoreplyContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterFlags is called when entering the flags production.
	EnterFlags(c *FlagsContext)

	// EnterExptime is called when entering the exptime production.
	EnterExptime(c *ExptimeContext)

	// EnterBytes is called when entering the bytes production.
	EnterBytes(c *BytesContext)

	// EnterCas_unique is called when entering the cas_unique production.
	EnterCas_unique(c *Cas_uniqueContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// EnterDelay is called when entering the delay production.
	EnterDelay(c *DelayContext)

	// EnterVerbosity_level is called when entering the verbosity_level production.
	EnterVerbosity_level(c *Verbosity_levelContext)

	// EnterStatistic_name is called when entering the statistic_name production.
	EnterStatistic_name(c *Statistic_nameContext)

	// EnterStatistic_value is called when entering the statistic_value production.
	EnterStatistic_value(c *Statistic_valueContext)

	// EnterSize is called when entering the size production.
	EnterSize(c *SizeContext)

	// EnterCount is called when entering the count production.
	EnterCount(c *CountContext)

	// ExitCommand_line is called when exiting the command_line production.
	ExitCommand_line(c *Command_lineContext)

	// ExitStorage_command is called when exiting the storage_command production.
	ExitStorage_command(c *Storage_commandContext)

	// ExitStorage_command_name is called when exiting the storage_command_name production.
	ExitStorage_command_name(c *Storage_command_nameContext)

	// ExitRetrieval_command is called when exiting the retrieval_command production.
	ExitRetrieval_command(c *Retrieval_commandContext)

	// ExitDelete_command is called when exiting the delete_command production.
	ExitDelete_command(c *Delete_commandContext)

	// ExitIncrement_command is called when exiting the increment_command production.
	ExitIncrement_command(c *Increment_commandContext)

	// ExitDecrement_command is called when exiting the decrement_command production.
	ExitDecrement_command(c *Decrement_commandContext)

	// ExitStatistics_command is called when exiting the statistics_command production.
	ExitStatistics_command(c *Statistics_commandContext)

	// ExitStatistics_option is called when exiting the statistics_option production.
	ExitStatistics_option(c *Statistics_optionContext)

	// ExitFlush_command is called when exiting the flush_command production.
	ExitFlush_command(c *Flush_commandContext)

	// ExitVersion_command is called when exiting the version_command production.
	ExitVersion_command(c *Version_commandContext)

	// ExitVerbosity_command is called when exiting the verbosity_command production.
	ExitVerbosity_command(c *Verbosity_commandContext)

	// ExitQuit_command is called when exiting the quit_command production.
	ExitQuit_command(c *Quit_commandContext)

	// ExitStorage_response is called when exiting the storage_response production.
	ExitStorage_response(c *Storage_responseContext)

	// ExitRetrieval_response is called when exiting the retrieval_response production.
	ExitRetrieval_response(c *Retrieval_responseContext)

	// ExitDeletion_response is called when exiting the deletion_response production.
	ExitDeletion_response(c *Deletion_responseContext)

	// ExitIncr_or_decr_response is called when exiting the incr_or_decr_response production.
	ExitIncr_or_decr_response(c *Incr_or_decr_responseContext)

	// ExitStatistics_response is called when exiting the statistics_response production.
	ExitStatistics_response(c *Statistics_responseContext)

	// ExitError_response is called when exiting the error_response production.
	ExitError_response(c *Error_responseContext)

	// ExitGeneral_statistic is called when exiting the general_statistic production.
	ExitGeneral_statistic(c *General_statisticContext)

	// ExitSize_statistic is called when exiting the size_statistic production.
	ExitSize_statistic(c *Size_statisticContext)

	// ExitGeneral_error is called when exiting the general_error production.
	ExitGeneral_error(c *General_errorContext)

	// ExitClient_error_message is called when exiting the client_error_message production.
	ExitClient_error_message(c *Client_error_messageContext)

	// ExitServer_error_message is called when exiting the server_error_message production.
	ExitServer_error_message(c *Server_error_messageContext)

	// ExitEnd is called when exiting the end production.
	ExitEnd(c *EndContext)

	// ExitNoreply is called when exiting the noreply production.
	ExitNoreply(c *NoreplyContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitFlags is called when exiting the flags production.
	ExitFlags(c *FlagsContext)

	// ExitExptime is called when exiting the exptime production.
	ExitExptime(c *ExptimeContext)

	// ExitBytes is called when exiting the bytes production.
	ExitBytes(c *BytesContext)

	// ExitCas_unique is called when exiting the cas_unique production.
	ExitCas_unique(c *Cas_uniqueContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)

	// ExitDelay is called when exiting the delay production.
	ExitDelay(c *DelayContext)

	// ExitVerbosity_level is called when exiting the verbosity_level production.
	ExitVerbosity_level(c *Verbosity_levelContext)

	// ExitStatistic_name is called when exiting the statistic_name production.
	ExitStatistic_name(c *Statistic_nameContext)

	// ExitStatistic_value is called when exiting the statistic_value production.
	ExitStatistic_value(c *Statistic_valueContext)

	// ExitSize is called when exiting the size production.
	ExitSize(c *SizeContext)

	// ExitCount is called when exiting the count production.
	ExitCount(c *CountContext)
}
