// Code generated from memcached_protocol.g4 by ANTLR 4.7.2. DO NOT EDIT.

package memcached_protocol // memcached_protocol
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 36, 251,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 6, 2, 91, 10, 2, 13, 2, 14, 2, 92, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 108, 10, 3, 3, 3,
	5, 3, 111, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 6, 5, 117, 10, 5, 13, 5, 14,
	5, 118, 3, 6, 3, 6, 3, 6, 5, 6, 124, 10, 6, 3, 6, 5, 6, 127, 10, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 5, 7, 133, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 139,
	10, 8, 3, 9, 3, 9, 5, 9, 143, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 5, 11,
	149, 10, 11, 3, 11, 5, 11, 152, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 166, 10, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 174, 10, 16, 3, 16, 5,
	16, 177, 10, 16, 3, 17, 3, 17, 3, 17, 5, 17, 182, 10, 17, 3, 18, 3, 18,
	3, 18, 5, 18, 187, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 193, 10,
	19, 3, 20, 3, 20, 3, 20, 5, 20, 198, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 6, 24, 211, 10, 24, 13,
	24, 14, 24, 212, 3, 25, 3, 25, 6, 25, 217, 10, 25, 13, 25, 14, 25, 218,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 4,
	212, 218, 2, 41, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 70, 72, 74, 76, 78, 2, 5, 3, 2, 4, 8, 3, 2, 9, 10, 3, 2, 15, 17, 2,
	249, 2, 90, 3, 2, 2, 2, 4, 107, 3, 2, 2, 2, 6, 112, 3, 2, 2, 2, 8, 114,
	3, 2, 2, 2, 10, 120, 3, 2, 2, 2, 12, 128, 3, 2, 2, 2, 14, 134, 3, 2, 2,
	2, 16, 140, 3, 2, 2, 2, 18, 144, 3, 2, 2, 2, 20, 146, 3, 2, 2, 2, 22, 153,
	3, 2, 2, 2, 24, 155, 3, 2, 2, 2, 26, 158, 3, 2, 2, 2, 28, 165, 3, 2, 2,
	2, 30, 176, 3, 2, 2, 2, 32, 181, 3, 2, 2, 2, 34, 186, 3, 2, 2, 2, 36, 192,
	3, 2, 2, 2, 38, 197, 3, 2, 2, 2, 40, 199, 3, 2, 2, 2, 42, 203, 3, 2, 2,
	2, 44, 206, 3, 2, 2, 2, 46, 208, 3, 2, 2, 2, 48, 214, 3, 2, 2, 2, 50, 220,
	3, 2, 2, 2, 52, 222, 3, 2, 2, 2, 54, 224, 3, 2, 2, 2, 56, 226, 3, 2, 2,
	2, 58, 228, 3, 2, 2, 2, 60, 230, 3, 2, 2, 2, 62, 232, 3, 2, 2, 2, 64, 234,
	3, 2, 2, 2, 66, 236, 3, 2, 2, 2, 68, 238, 3, 2, 2, 2, 70, 240, 3, 2, 2,
	2, 72, 242, 3, 2, 2, 2, 74, 244, 3, 2, 2, 2, 76, 246, 3, 2, 2, 2, 78, 248,
	3, 2, 2, 2, 80, 91, 5, 4, 3, 2, 81, 91, 5, 8, 5, 2, 82, 91, 5, 10, 6, 2,
	83, 91, 5, 12, 7, 2, 84, 91, 5, 14, 8, 2, 85, 91, 5, 16, 9, 2, 86, 91,
	5, 20, 11, 2, 87, 91, 5, 22, 12, 2, 88, 91, 5, 24, 13, 2, 89, 91, 5, 26,
	14, 2, 90, 80, 3, 2, 2, 2, 90, 81, 3, 2, 2, 2, 90, 82, 3, 2, 2, 2, 90,
	83, 3, 2, 2, 2, 90, 84, 3, 2, 2, 2, 90, 85, 3, 2, 2, 2, 90, 86, 3, 2, 2,
	2, 90, 87, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 89, 3, 2, 2, 2, 91, 92,
	3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 3, 3, 2, 2, 2,
	94, 95, 5, 6, 4, 2, 95, 96, 5, 54, 28, 2, 96, 97, 5, 56, 29, 2, 97, 98,
	5, 58, 30, 2, 98, 99, 5, 60, 31, 2, 99, 108, 3, 2, 2, 2, 100, 101, 7, 3,
	2, 2, 101, 102, 5, 54, 28, 2, 102, 103, 5, 56, 29, 2, 103, 104, 5, 58,
	30, 2, 104, 105, 5, 60, 31, 2, 105, 106, 5, 62, 32, 2, 106, 108, 3, 2,
	2, 2, 107, 94, 3, 2, 2, 2, 107, 100, 3, 2, 2, 2, 108, 110, 3, 2, 2, 2,
	109, 111, 5, 52, 27, 2, 110, 109, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111,
	5, 3, 2, 2, 2, 112, 113, 9, 2, 2, 2, 113, 7, 3, 2, 2, 2, 114, 116, 9, 3,
	2, 2, 115, 117, 5, 54, 28, 2, 116, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2,
	2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 9, 3, 2, 2, 2, 120,
	121, 7, 11, 2, 2, 121, 123, 5, 54, 28, 2, 122, 124, 5, 66, 34, 2, 123,
	122, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 126, 3, 2, 2, 2, 125, 127,
	5, 52, 27, 2, 126, 125, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 11, 3, 2,
	2, 2, 128, 129, 7, 12, 2, 2, 129, 130, 5, 54, 28, 2, 130, 132, 5, 64, 33,
	2, 131, 133, 5, 52, 27, 2, 132, 131, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2,
	133, 13, 3, 2, 2, 2, 134, 135, 7, 13, 2, 2, 135, 136, 5, 54, 28, 2, 136,
	138, 5, 64, 33, 2, 137, 139, 5, 52, 27, 2, 138, 137, 3, 2, 2, 2, 138, 139,
	3, 2, 2, 2, 139, 15, 3, 2, 2, 2, 140, 142, 7, 14, 2, 2, 141, 143, 5, 18,
	10, 2, 142, 141, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 17, 3, 2, 2, 2,
	144, 145, 9, 4, 2, 2, 145, 19, 3, 2, 2, 2, 146, 148, 7, 18, 2, 2, 147,
	149, 5, 68, 35, 2, 148, 147, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 151,
	3, 2, 2, 2, 150, 152, 5, 52, 27, 2, 151, 150, 3, 2, 2, 2, 151, 152, 3,
	2, 2, 2, 152, 21, 3, 2, 2, 2, 153, 154, 7, 19, 2, 2, 154, 23, 3, 2, 2,
	2, 155, 156, 7, 20, 2, 2, 156, 157, 5, 70, 36, 2, 157, 25, 3, 2, 2, 2,
	158, 159, 7, 21, 2, 2, 159, 27, 3, 2, 2, 2, 160, 166, 5, 38, 20, 2, 161,
	166, 7, 22, 2, 2, 162, 166, 7, 23, 2, 2, 163, 166, 7, 24, 2, 2, 164, 166,
	7, 25, 2, 2, 165, 160, 3, 2, 2, 2, 165, 161, 3, 2, 2, 2, 165, 162, 3, 2,
	2, 2, 165, 163, 3, 2, 2, 2, 165, 164, 3, 2, 2, 2, 166, 29, 3, 2, 2, 2,
	167, 177, 5, 38, 20, 2, 168, 169, 7, 26, 2, 2, 169, 170, 5, 54, 28, 2,
	170, 171, 5, 56, 29, 2, 171, 173, 5, 60, 31, 2, 172, 174, 5, 62, 32, 2,
	173, 172, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 177, 3, 2, 2, 2, 175,
	177, 5, 50, 26, 2, 176, 167, 3, 2, 2, 2, 176, 168, 3, 2, 2, 2, 176, 175,
	3, 2, 2, 2, 177, 31, 3, 2, 2, 2, 178, 182, 5, 38, 20, 2, 179, 182, 7, 27,
	2, 2, 180, 182, 7, 25, 2, 2, 181, 178, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2,
	181, 180, 3, 2, 2, 2, 182, 33, 3, 2, 2, 2, 183, 187, 5, 38, 20, 2, 184,
	187, 7, 25, 2, 2, 185, 187, 7, 34, 2, 2, 186, 183, 3, 2, 2, 2, 186, 184,
	3, 2, 2, 2, 186, 185, 3, 2, 2, 2, 187, 35, 3, 2, 2, 2, 188, 193, 5, 38,
	20, 2, 189, 193, 5, 40, 21, 2, 190, 193, 5, 42, 22, 2, 191, 193, 5, 50,
	26, 2, 192, 188, 3, 2, 2, 2, 192, 189, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2,
	192, 191, 3, 2, 2, 2, 193, 37, 3, 2, 2, 2, 194, 198, 5, 44, 23, 2, 195,
	198, 5, 46, 24, 2, 196, 198, 5, 48, 25, 2, 197, 194, 3, 2, 2, 2, 197, 195,
	3, 2, 2, 2, 197, 196, 3, 2, 2, 2, 198, 39, 3, 2, 2, 2, 199, 200, 7, 28,
	2, 2, 200, 201, 5, 72, 37, 2, 201, 202, 5, 74, 38, 2, 202, 41, 3, 2, 2,
	2, 203, 204, 5, 76, 39, 2, 204, 205, 5, 78, 40, 2, 205, 43, 3, 2, 2, 2,
	206, 207, 7, 29, 2, 2, 207, 45, 3, 2, 2, 2, 208, 210, 7, 30, 2, 2, 209,
	211, 11, 2, 2, 2, 210, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213,
	3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 47, 3, 2, 2, 2, 214, 216, 7, 31,
	2, 2, 215, 217, 11, 2, 2, 2, 216, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2,
	218, 219, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 219, 49, 3, 2, 2, 2, 220, 221,
	7, 32, 2, 2, 221, 51, 3, 2, 2, 2, 222, 223, 7, 33, 2, 2, 223, 53, 3, 2,
	2, 2, 224, 225, 11, 2, 2, 2, 225, 55, 3, 2, 2, 2, 226, 227, 7, 34, 2, 2,
	227, 57, 3, 2, 2, 2, 228, 229, 7, 34, 2, 2, 229, 59, 3, 2, 2, 2, 230, 231,
	7, 34, 2, 2, 231, 61, 3, 2, 2, 2, 232, 233, 7, 34, 2, 2, 233, 63, 3, 2,
	2, 2, 234, 235, 7, 34, 2, 2, 235, 65, 3, 2, 2, 2, 236, 237, 7, 34, 2, 2,
	237, 67, 3, 2, 2, 2, 238, 239, 7, 34, 2, 2, 239, 69, 3, 2, 2, 2, 240, 241,
	7, 34, 2, 2, 241, 71, 3, 2, 2, 2, 242, 243, 7, 35, 2, 2, 243, 73, 3, 2,
	2, 2, 244, 245, 11, 2, 2, 2, 245, 75, 3, 2, 2, 2, 246, 247, 7, 34, 2, 2,
	247, 77, 3, 2, 2, 2, 248, 249, 7, 34, 2, 2, 249, 79, 3, 2, 2, 2, 23, 90,
	92, 107, 110, 118, 123, 126, 132, 138, 142, 148, 151, 165, 173, 176, 181,
	186, 192, 197, 212, 218,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'cas'", "'set'", "'add'", "'replace'", "'append'", "'prepend'", "'get'",
	"'gets'", "'delete'", "'incr'", "'decr'", "'stats'", "'items'", "'slabs'",
	"'sizes'", "'flush_all'", "'version'", "'verbosity'", "'quit'", "'STORED'",
	"'NOT_STORED'", "'EXISTS'", "'NOT_FOUND'", "'VALUE'", "'DELETED'", "'STAT'",
	"'ERROR'", "'CLIENT_ERROR'", "'SERVER_ERROR'", "'END'", "'noreply'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "INTEGER", "WORD",
	"WHITESPACE",
}

var ruleNames = []string{
	"command_line", "storage_command", "storage_command_name", "retrieval_command",
	"delete_command", "increment_command", "decrement_command", "statistics_command",
	"statistics_option", "flush_command", "version_command", "verbosity_command",
	"quit_command", "storage_response", "retrieval_response", "deletion_response",
	"incr_or_decr_response", "statistics_response", "error_response", "general_statistic",
	"size_statistic", "general_error", "client_error_message", "server_error_message",
	"end", "noreply", "key", "flags", "exptime", "bytes", "cas_unique", "value",
	"time", "delay", "verbosity_level", "statistic_name", "statistic_value",
	"size", "count",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type memcached_protocolParser struct {
	*antlr.BaseParser
}

func Newmemcached_protocolParser(input antlr.TokenStream) *memcached_protocolParser {
	this := new(memcached_protocolParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "memcached_protocol.g4"

	return this
}

// memcached_protocolParser tokens.
const (
	memcached_protocolParserEOF        = antlr.TokenEOF
	memcached_protocolParserT__0       = 1
	memcached_protocolParserT__1       = 2
	memcached_protocolParserT__2       = 3
	memcached_protocolParserT__3       = 4
	memcached_protocolParserT__4       = 5
	memcached_protocolParserT__5       = 6
	memcached_protocolParserT__6       = 7
	memcached_protocolParserT__7       = 8
	memcached_protocolParserT__8       = 9
	memcached_protocolParserT__9       = 10
	memcached_protocolParserT__10      = 11
	memcached_protocolParserT__11      = 12
	memcached_protocolParserT__12      = 13
	memcached_protocolParserT__13      = 14
	memcached_protocolParserT__14      = 15
	memcached_protocolParserT__15      = 16
	memcached_protocolParserT__16      = 17
	memcached_protocolParserT__17      = 18
	memcached_protocolParserT__18      = 19
	memcached_protocolParserT__19      = 20
	memcached_protocolParserT__20      = 21
	memcached_protocolParserT__21      = 22
	memcached_protocolParserT__22      = 23
	memcached_protocolParserT__23      = 24
	memcached_protocolParserT__24      = 25
	memcached_protocolParserT__25      = 26
	memcached_protocolParserT__26      = 27
	memcached_protocolParserT__27      = 28
	memcached_protocolParserT__28      = 29
	memcached_protocolParserT__29      = 30
	memcached_protocolParserT__30      = 31
	memcached_protocolParserINTEGER    = 32
	memcached_protocolParserWORD       = 33
	memcached_protocolParserWHITESPACE = 34
)

// memcached_protocolParser rules.
const (
	memcached_protocolParserRULE_command_line          = 0
	memcached_protocolParserRULE_storage_command       = 1
	memcached_protocolParserRULE_storage_command_name  = 2
	memcached_protocolParserRULE_retrieval_command     = 3
	memcached_protocolParserRULE_delete_command        = 4
	memcached_protocolParserRULE_increment_command     = 5
	memcached_protocolParserRULE_decrement_command     = 6
	memcached_protocolParserRULE_statistics_command    = 7
	memcached_protocolParserRULE_statistics_option     = 8
	memcached_protocolParserRULE_flush_command         = 9
	memcached_protocolParserRULE_version_command       = 10
	memcached_protocolParserRULE_verbosity_command     = 11
	memcached_protocolParserRULE_quit_command          = 12
	memcached_protocolParserRULE_storage_response      = 13
	memcached_protocolParserRULE_retrieval_response    = 14
	memcached_protocolParserRULE_deletion_response     = 15
	memcached_protocolParserRULE_incr_or_decr_response = 16
	memcached_protocolParserRULE_statistics_response   = 17
	memcached_protocolParserRULE_error_response        = 18
	memcached_protocolParserRULE_general_statistic     = 19
	memcached_protocolParserRULE_size_statistic        = 20
	memcached_protocolParserRULE_general_error         = 21
	memcached_protocolParserRULE_client_error_message  = 22
	memcached_protocolParserRULE_server_error_message  = 23
	memcached_protocolParserRULE_end                   = 24
	memcached_protocolParserRULE_noreply               = 25
	memcached_protocolParserRULE_key                   = 26
	memcached_protocolParserRULE_flags                 = 27
	memcached_protocolParserRULE_exptime               = 28
	memcached_protocolParserRULE_bytes                 = 29
	memcached_protocolParserRULE_cas_unique            = 30
	memcached_protocolParserRULE_value                 = 31
	memcached_protocolParserRULE_time                  = 32
	memcached_protocolParserRULE_delay                 = 33
	memcached_protocolParserRULE_verbosity_level       = 34
	memcached_protocolParserRULE_statistic_name        = 35
	memcached_protocolParserRULE_statistic_value       = 36
	memcached_protocolParserRULE_size                  = 37
	memcached_protocolParserRULE_count                 = 38
)

// ICommand_lineContext is an interface to support dynamic dispatch.
type ICommand_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommand_lineContext differentiates from other interfaces.
	IsCommand_lineContext()
}

type Command_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommand_lineContext() *Command_lineContext {
	var p = new(Command_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_command_line
	return p
}

func (*Command_lineContext) IsCommand_lineContext() {}

func NewCommand_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Command_lineContext {
	var p = new(Command_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_command_line

	return p
}

func (s *Command_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Command_lineContext) AllStorage_command() []IStorage_commandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStorage_commandContext)(nil)).Elem())
	var tst = make([]IStorage_commandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStorage_commandContext)
		}
	}

	return tst
}

func (s *Command_lineContext) Storage_command(i int) IStorage_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStorage_commandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStorage_commandContext)
}

func (s *Command_lineContext) AllRetrieval_command() []IRetrieval_commandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRetrieval_commandContext)(nil)).Elem())
	var tst = make([]IRetrieval_commandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRetrieval_commandContext)
		}
	}

	return tst
}

func (s *Command_lineContext) Retrieval_command(i int) IRetrieval_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetrieval_commandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRetrieval_commandContext)
}

func (s *Command_lineContext) AllDelete_command() []IDelete_commandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDelete_commandContext)(nil)).Elem())
	var tst = make([]IDelete_commandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDelete_commandContext)
		}
	}

	return tst
}

func (s *Command_lineContext) Delete_command(i int) IDelete_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_commandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDelete_commandContext)
}

func (s *Command_lineContext) AllIncrement_command() []IIncrement_commandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIncrement_commandContext)(nil)).Elem())
	var tst = make([]IIncrement_commandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIncrement_commandContext)
		}
	}

	return tst
}

func (s *Command_lineContext) Increment_command(i int) IIncrement_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncrement_commandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIncrement_commandContext)
}

func (s *Command_lineContext) AllDecrement_command() []IDecrement_commandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDecrement_commandContext)(nil)).Elem())
	var tst = make([]IDecrement_commandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDecrement_commandContext)
		}
	}

	return tst
}

func (s *Command_lineContext) Decrement_command(i int) IDecrement_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecrement_commandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDecrement_commandContext)
}

func (s *Command_lineContext) AllStatistics_command() []IStatistics_commandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatistics_commandContext)(nil)).Elem())
	var tst = make([]IStatistics_commandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatistics_commandContext)
		}
	}

	return tst
}

func (s *Command_lineContext) Statistics_command(i int) IStatistics_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatistics_commandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatistics_commandContext)
}

func (s *Command_lineContext) AllFlush_command() []IFlush_commandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFlush_commandContext)(nil)).Elem())
	var tst = make([]IFlush_commandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFlush_commandContext)
		}
	}

	return tst
}

func (s *Command_lineContext) Flush_command(i int) IFlush_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlush_commandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFlush_commandContext)
}

func (s *Command_lineContext) AllVersion_command() []IVersion_commandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVersion_commandContext)(nil)).Elem())
	var tst = make([]IVersion_commandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVersion_commandContext)
		}
	}

	return tst
}

func (s *Command_lineContext) Version_command(i int) IVersion_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVersion_commandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVersion_commandContext)
}

func (s *Command_lineContext) AllVerbosity_command() []IVerbosity_commandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVerbosity_commandContext)(nil)).Elem())
	var tst = make([]IVerbosity_commandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVerbosity_commandContext)
		}
	}

	return tst
}

func (s *Command_lineContext) Verbosity_command(i int) IVerbosity_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVerbosity_commandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVerbosity_commandContext)
}

func (s *Command_lineContext) AllQuit_command() []IQuit_commandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuit_commandContext)(nil)).Elem())
	var tst = make([]IQuit_commandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuit_commandContext)
		}
	}

	return tst
}

func (s *Command_lineContext) Quit_command(i int) IQuit_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuit_commandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuit_commandContext)
}

func (s *Command_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Command_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterCommand_line(s)
	}
}

func (s *Command_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitCommand_line(s)
	}
}

func (p *memcached_protocolParser) Command_line() (localctx ICommand_lineContext) {
	localctx = NewCommand_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, memcached_protocolParserRULE_command_line)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<memcached_protocolParserT__0)|(1<<memcached_protocolParserT__1)|(1<<memcached_protocolParserT__2)|(1<<memcached_protocolParserT__3)|(1<<memcached_protocolParserT__4)|(1<<memcached_protocolParserT__5)|(1<<memcached_protocolParserT__6)|(1<<memcached_protocolParserT__7)|(1<<memcached_protocolParserT__8)|(1<<memcached_protocolParserT__9)|(1<<memcached_protocolParserT__10)|(1<<memcached_protocolParserT__11)|(1<<memcached_protocolParserT__15)|(1<<memcached_protocolParserT__16)|(1<<memcached_protocolParserT__17)|(1<<memcached_protocolParserT__18))) != 0) {
		p.SetState(88)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case memcached_protocolParserT__0, memcached_protocolParserT__1, memcached_protocolParserT__2, memcached_protocolParserT__3, memcached_protocolParserT__4, memcached_protocolParserT__5:
			{
				p.SetState(78)
				p.Storage_command()
			}

		case memcached_protocolParserT__6, memcached_protocolParserT__7:
			{
				p.SetState(79)
				p.Retrieval_command()
			}

		case memcached_protocolParserT__8:
			{
				p.SetState(80)
				p.Delete_command()
			}

		case memcached_protocolParserT__9:
			{
				p.SetState(81)
				p.Increment_command()
			}

		case memcached_protocolParserT__10:
			{
				p.SetState(82)
				p.Decrement_command()
			}

		case memcached_protocolParserT__11:
			{
				p.SetState(83)
				p.Statistics_command()
			}

		case memcached_protocolParserT__15:
			{
				p.SetState(84)
				p.Flush_command()
			}

		case memcached_protocolParserT__16:
			{
				p.SetState(85)
				p.Version_command()
			}

		case memcached_protocolParserT__17:
			{
				p.SetState(86)
				p.Verbosity_command()
			}

		case memcached_protocolParserT__18:
			{
				p.SetState(87)
				p.Quit_command()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStorage_commandContext is an interface to support dynamic dispatch.
type IStorage_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStorage_commandContext differentiates from other interfaces.
	IsStorage_commandContext()
}

type Storage_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStorage_commandContext() *Storage_commandContext {
	var p = new(Storage_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_storage_command
	return p
}

func (*Storage_commandContext) IsStorage_commandContext() {}

func NewStorage_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Storage_commandContext {
	var p = new(Storage_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_storage_command

	return p
}

func (s *Storage_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Storage_commandContext) Noreply() INoreplyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoreplyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoreplyContext)
}

func (s *Storage_commandContext) Storage_command_name() IStorage_command_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStorage_command_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStorage_command_nameContext)
}

func (s *Storage_commandContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Storage_commandContext) Flags() IFlagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlagsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlagsContext)
}

func (s *Storage_commandContext) Exptime() IExptimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExptimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExptimeContext)
}

func (s *Storage_commandContext) Bytes() IBytesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBytesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBytesContext)
}

func (s *Storage_commandContext) Cas_unique() ICas_uniqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICas_uniqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICas_uniqueContext)
}

func (s *Storage_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Storage_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Storage_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterStorage_command(s)
	}
}

func (s *Storage_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitStorage_command(s)
	}
}

func (p *memcached_protocolParser) Storage_command() (localctx IStorage_commandContext) {
	localctx = NewStorage_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, memcached_protocolParserRULE_storage_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case memcached_protocolParserT__1, memcached_protocolParserT__2, memcached_protocolParserT__3, memcached_protocolParserT__4, memcached_protocolParserT__5:
		{
			p.SetState(92)
			p.Storage_command_name()
		}
		{
			p.SetState(93)
			p.Key()
		}
		{
			p.SetState(94)
			p.Flags()
		}
		{
			p.SetState(95)
			p.Exptime()
		}
		{
			p.SetState(96)
			p.Bytes()
		}

	case memcached_protocolParserT__0:
		{
			p.SetState(98)
			p.Match(memcached_protocolParserT__0)
		}
		{
			p.SetState(99)
			p.Key()
		}
		{
			p.SetState(100)
			p.Flags()
		}
		{
			p.SetState(101)
			p.Exptime()
		}
		{
			p.SetState(102)
			p.Bytes()
		}
		{
			p.SetState(103)
			p.Cas_unique()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == memcached_protocolParserT__30 {
		{
			p.SetState(107)
			p.Noreply()
		}

	}

	return localctx
}

// IStorage_command_nameContext is an interface to support dynamic dispatch.
type IStorage_command_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStorage_command_nameContext differentiates from other interfaces.
	IsStorage_command_nameContext()
}

type Storage_command_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStorage_command_nameContext() *Storage_command_nameContext {
	var p = new(Storage_command_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_storage_command_name
	return p
}

func (*Storage_command_nameContext) IsStorage_command_nameContext() {}

func NewStorage_command_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Storage_command_nameContext {
	var p = new(Storage_command_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_storage_command_name

	return p
}

func (s *Storage_command_nameContext) GetParser() antlr.Parser { return s.parser }
func (s *Storage_command_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Storage_command_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Storage_command_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterStorage_command_name(s)
	}
}

func (s *Storage_command_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitStorage_command_name(s)
	}
}

func (p *memcached_protocolParser) Storage_command_name() (localctx IStorage_command_nameContext) {
	localctx = NewStorage_command_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, memcached_protocolParserRULE_storage_command_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<memcached_protocolParserT__1)|(1<<memcached_protocolParserT__2)|(1<<memcached_protocolParserT__3)|(1<<memcached_protocolParserT__4)|(1<<memcached_protocolParserT__5))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRetrieval_commandContext is an interface to support dynamic dispatch.
type IRetrieval_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetrieval_commandContext differentiates from other interfaces.
	IsRetrieval_commandContext()
}

type Retrieval_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetrieval_commandContext() *Retrieval_commandContext {
	var p = new(Retrieval_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_retrieval_command
	return p
}

func (*Retrieval_commandContext) IsRetrieval_commandContext() {}

func NewRetrieval_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Retrieval_commandContext {
	var p = new(Retrieval_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_retrieval_command

	return p
}

func (s *Retrieval_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Retrieval_commandContext) AllKey() []IKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyContext)(nil)).Elem())
	var tst = make([]IKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyContext)
		}
	}

	return tst
}

func (s *Retrieval_commandContext) Key(i int) IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Retrieval_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Retrieval_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Retrieval_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterRetrieval_command(s)
	}
}

func (s *Retrieval_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitRetrieval_command(s)
	}
}

func (p *memcached_protocolParser) Retrieval_command() (localctx IRetrieval_commandContext) {
	localctx = NewRetrieval_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, memcached_protocolParserRULE_retrieval_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !(_la == memcached_protocolParserT__6 || _la == memcached_protocolParserT__7) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(113)
				p.Key()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IDelete_commandContext is an interface to support dynamic dispatch.
type IDelete_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelete_commandContext differentiates from other interfaces.
	IsDelete_commandContext()
}

type Delete_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_commandContext() *Delete_commandContext {
	var p = new(Delete_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_delete_command
	return p
}

func (*Delete_commandContext) IsDelete_commandContext() {}

func NewDelete_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_commandContext {
	var p = new(Delete_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_delete_command

	return p
}

func (s *Delete_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_commandContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Delete_commandContext) Time() ITimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeContext)
}

func (s *Delete_commandContext) Noreply() INoreplyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoreplyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoreplyContext)
}

func (s *Delete_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterDelete_command(s)
	}
}

func (s *Delete_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitDelete_command(s)
	}
}

func (p *memcached_protocolParser) Delete_command() (localctx IDelete_commandContext) {
	localctx = NewDelete_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, memcached_protocolParserRULE_delete_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(memcached_protocolParserT__8)
	}
	{
		p.SetState(119)
		p.Key()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == memcached_protocolParserINTEGER {
		{
			p.SetState(120)
			p.Time()
		}

	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == memcached_protocolParserT__30 {
		{
			p.SetState(123)
			p.Noreply()
		}

	}

	return localctx
}

// IIncrement_commandContext is an interface to support dynamic dispatch.
type IIncrement_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncrement_commandContext differentiates from other interfaces.
	IsIncrement_commandContext()
}

type Increment_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncrement_commandContext() *Increment_commandContext {
	var p = new(Increment_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_increment_command
	return p
}

func (*Increment_commandContext) IsIncrement_commandContext() {}

func NewIncrement_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Increment_commandContext {
	var p = new(Increment_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_increment_command

	return p
}

func (s *Increment_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Increment_commandContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Increment_commandContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Increment_commandContext) Noreply() INoreplyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoreplyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoreplyContext)
}

func (s *Increment_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Increment_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Increment_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterIncrement_command(s)
	}
}

func (s *Increment_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitIncrement_command(s)
	}
}

func (p *memcached_protocolParser) Increment_command() (localctx IIncrement_commandContext) {
	localctx = NewIncrement_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, memcached_protocolParserRULE_increment_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(memcached_protocolParserT__9)
	}
	{
		p.SetState(127)
		p.Key()
	}
	{
		p.SetState(128)
		p.Value()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == memcached_protocolParserT__30 {
		{
			p.SetState(129)
			p.Noreply()
		}

	}

	return localctx
}

// IDecrement_commandContext is an interface to support dynamic dispatch.
type IDecrement_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecrement_commandContext differentiates from other interfaces.
	IsDecrement_commandContext()
}

type Decrement_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecrement_commandContext() *Decrement_commandContext {
	var p = new(Decrement_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_decrement_command
	return p
}

func (*Decrement_commandContext) IsDecrement_commandContext() {}

func NewDecrement_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decrement_commandContext {
	var p = new(Decrement_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_decrement_command

	return p
}

func (s *Decrement_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Decrement_commandContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Decrement_commandContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Decrement_commandContext) Noreply() INoreplyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoreplyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoreplyContext)
}

func (s *Decrement_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decrement_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Decrement_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterDecrement_command(s)
	}
}

func (s *Decrement_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitDecrement_command(s)
	}
}

func (p *memcached_protocolParser) Decrement_command() (localctx IDecrement_commandContext) {
	localctx = NewDecrement_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, memcached_protocolParserRULE_decrement_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(memcached_protocolParserT__10)
	}
	{
		p.SetState(133)
		p.Key()
	}
	{
		p.SetState(134)
		p.Value()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == memcached_protocolParserT__30 {
		{
			p.SetState(135)
			p.Noreply()
		}

	}

	return localctx
}

// IStatistics_commandContext is an interface to support dynamic dispatch.
type IStatistics_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatistics_commandContext differentiates from other interfaces.
	IsStatistics_commandContext()
}

type Statistics_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatistics_commandContext() *Statistics_commandContext {
	var p = new(Statistics_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_statistics_command
	return p
}

func (*Statistics_commandContext) IsStatistics_commandContext() {}

func NewStatistics_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statistics_commandContext {
	var p = new(Statistics_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_statistics_command

	return p
}

func (s *Statistics_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Statistics_commandContext) Statistics_option() IStatistics_optionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatistics_optionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatistics_optionContext)
}

func (s *Statistics_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statistics_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statistics_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterStatistics_command(s)
	}
}

func (s *Statistics_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitStatistics_command(s)
	}
}

func (p *memcached_protocolParser) Statistics_command() (localctx IStatistics_commandContext) {
	localctx = NewStatistics_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, memcached_protocolParserRULE_statistics_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(memcached_protocolParserT__11)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<memcached_protocolParserT__12)|(1<<memcached_protocolParserT__13)|(1<<memcached_protocolParserT__14))) != 0 {
		{
			p.SetState(139)
			p.Statistics_option()
		}

	}

	return localctx
}

// IStatistics_optionContext is an interface to support dynamic dispatch.
type IStatistics_optionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatistics_optionContext differentiates from other interfaces.
	IsStatistics_optionContext()
}

type Statistics_optionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatistics_optionContext() *Statistics_optionContext {
	var p = new(Statistics_optionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_statistics_option
	return p
}

func (*Statistics_optionContext) IsStatistics_optionContext() {}

func NewStatistics_optionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statistics_optionContext {
	var p = new(Statistics_optionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_statistics_option

	return p
}

func (s *Statistics_optionContext) GetParser() antlr.Parser { return s.parser }
func (s *Statistics_optionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statistics_optionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statistics_optionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterStatistics_option(s)
	}
}

func (s *Statistics_optionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitStatistics_option(s)
	}
}

func (p *memcached_protocolParser) Statistics_option() (localctx IStatistics_optionContext) {
	localctx = NewStatistics_optionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, memcached_protocolParserRULE_statistics_option)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<memcached_protocolParserT__12)|(1<<memcached_protocolParserT__13)|(1<<memcached_protocolParserT__14))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFlush_commandContext is an interface to support dynamic dispatch.
type IFlush_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlush_commandContext differentiates from other interfaces.
	IsFlush_commandContext()
}

type Flush_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlush_commandContext() *Flush_commandContext {
	var p = new(Flush_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_flush_command
	return p
}

func (*Flush_commandContext) IsFlush_commandContext() {}

func NewFlush_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Flush_commandContext {
	var p = new(Flush_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_flush_command

	return p
}

func (s *Flush_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Flush_commandContext) Delay() IDelayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelayContext)
}

func (s *Flush_commandContext) Noreply() INoreplyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoreplyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoreplyContext)
}

func (s *Flush_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Flush_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Flush_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterFlush_command(s)
	}
}

func (s *Flush_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitFlush_command(s)
	}
}

func (p *memcached_protocolParser) Flush_command() (localctx IFlush_commandContext) {
	localctx = NewFlush_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, memcached_protocolParserRULE_flush_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(memcached_protocolParserT__15)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == memcached_protocolParserINTEGER {
		{
			p.SetState(145)
			p.Delay()
		}

	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == memcached_protocolParserT__30 {
		{
			p.SetState(148)
			p.Noreply()
		}

	}

	return localctx
}

// IVersion_commandContext is an interface to support dynamic dispatch.
type IVersion_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVersion_commandContext differentiates from other interfaces.
	IsVersion_commandContext()
}

type Version_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersion_commandContext() *Version_commandContext {
	var p = new(Version_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_version_command
	return p
}

func (*Version_commandContext) IsVersion_commandContext() {}

func NewVersion_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Version_commandContext {
	var p = new(Version_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_version_command

	return p
}

func (s *Version_commandContext) GetParser() antlr.Parser { return s.parser }
func (s *Version_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Version_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Version_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterVersion_command(s)
	}
}

func (s *Version_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitVersion_command(s)
	}
}

func (p *memcached_protocolParser) Version_command() (localctx IVersion_commandContext) {
	localctx = NewVersion_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, memcached_protocolParserRULE_version_command)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(memcached_protocolParserT__16)
	}

	return localctx
}

// IVerbosity_commandContext is an interface to support dynamic dispatch.
type IVerbosity_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVerbosity_commandContext differentiates from other interfaces.
	IsVerbosity_commandContext()
}

type Verbosity_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerbosity_commandContext() *Verbosity_commandContext {
	var p = new(Verbosity_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_verbosity_command
	return p
}

func (*Verbosity_commandContext) IsVerbosity_commandContext() {}

func NewVerbosity_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Verbosity_commandContext {
	var p = new(Verbosity_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_verbosity_command

	return p
}

func (s *Verbosity_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Verbosity_commandContext) Verbosity_level() IVerbosity_levelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVerbosity_levelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVerbosity_levelContext)
}

func (s *Verbosity_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Verbosity_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Verbosity_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterVerbosity_command(s)
	}
}

func (s *Verbosity_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitVerbosity_command(s)
	}
}

func (p *memcached_protocolParser) Verbosity_command() (localctx IVerbosity_commandContext) {
	localctx = NewVerbosity_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, memcached_protocolParserRULE_verbosity_command)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(memcached_protocolParserT__17)
	}
	{
		p.SetState(154)
		p.Verbosity_level()
	}

	return localctx
}

// IQuit_commandContext is an interface to support dynamic dispatch.
type IQuit_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuit_commandContext differentiates from other interfaces.
	IsQuit_commandContext()
}

type Quit_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuit_commandContext() *Quit_commandContext {
	var p = new(Quit_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_quit_command
	return p
}

func (*Quit_commandContext) IsQuit_commandContext() {}

func NewQuit_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Quit_commandContext {
	var p = new(Quit_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_quit_command

	return p
}

func (s *Quit_commandContext) GetParser() antlr.Parser { return s.parser }
func (s *Quit_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Quit_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Quit_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterQuit_command(s)
	}
}

func (s *Quit_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitQuit_command(s)
	}
}

func (p *memcached_protocolParser) Quit_command() (localctx IQuit_commandContext) {
	localctx = NewQuit_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, memcached_protocolParserRULE_quit_command)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(memcached_protocolParserT__18)
	}

	return localctx
}

// IStorage_responseContext is an interface to support dynamic dispatch.
type IStorage_responseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStorage_responseContext differentiates from other interfaces.
	IsStorage_responseContext()
}

type Storage_responseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStorage_responseContext() *Storage_responseContext {
	var p = new(Storage_responseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_storage_response
	return p
}

func (*Storage_responseContext) IsStorage_responseContext() {}

func NewStorage_responseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Storage_responseContext {
	var p = new(Storage_responseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_storage_response

	return p
}

func (s *Storage_responseContext) GetParser() antlr.Parser { return s.parser }

func (s *Storage_responseContext) Error_response() IError_responseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IError_responseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IError_responseContext)
}

func (s *Storage_responseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Storage_responseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Storage_responseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterStorage_response(s)
	}
}

func (s *Storage_responseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitStorage_response(s)
	}
}

func (p *memcached_protocolParser) Storage_response() (localctx IStorage_responseContext) {
	localctx = NewStorage_responseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, memcached_protocolParserRULE_storage_response)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case memcached_protocolParserT__26, memcached_protocolParserT__27, memcached_protocolParserT__28:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Error_response()
		}

	case memcached_protocolParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Match(memcached_protocolParserT__19)
		}

	case memcached_protocolParserT__20:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Match(memcached_protocolParserT__20)
		}

	case memcached_protocolParserT__21:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(161)
			p.Match(memcached_protocolParserT__21)
		}

	case memcached_protocolParserT__22:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(162)
			p.Match(memcached_protocolParserT__22)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRetrieval_responseContext is an interface to support dynamic dispatch.
type IRetrieval_responseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetrieval_responseContext differentiates from other interfaces.
	IsRetrieval_responseContext()
}

type Retrieval_responseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetrieval_responseContext() *Retrieval_responseContext {
	var p = new(Retrieval_responseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_retrieval_response
	return p
}

func (*Retrieval_responseContext) IsRetrieval_responseContext() {}

func NewRetrieval_responseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Retrieval_responseContext {
	var p = new(Retrieval_responseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_retrieval_response

	return p
}

func (s *Retrieval_responseContext) GetParser() antlr.Parser { return s.parser }

func (s *Retrieval_responseContext) Error_response() IError_responseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IError_responseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IError_responseContext)
}

func (s *Retrieval_responseContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Retrieval_responseContext) Flags() IFlagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFlagsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFlagsContext)
}

func (s *Retrieval_responseContext) Bytes() IBytesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBytesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBytesContext)
}

func (s *Retrieval_responseContext) Cas_unique() ICas_uniqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICas_uniqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICas_uniqueContext)
}

func (s *Retrieval_responseContext) End() IEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndContext)
}

func (s *Retrieval_responseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Retrieval_responseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Retrieval_responseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterRetrieval_response(s)
	}
}

func (s *Retrieval_responseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitRetrieval_response(s)
	}
}

func (p *memcached_protocolParser) Retrieval_response() (localctx IRetrieval_responseContext) {
	localctx = NewRetrieval_responseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, memcached_protocolParserRULE_retrieval_response)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case memcached_protocolParserT__26, memcached_protocolParserT__27, memcached_protocolParserT__28:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Error_response()
		}

	case memcached_protocolParserT__23:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Match(memcached_protocolParserT__23)
		}
		{
			p.SetState(167)
			p.Key()
		}
		{
			p.SetState(168)
			p.Flags()
		}
		{
			p.SetState(169)
			p.Bytes()
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == memcached_protocolParserINTEGER {
			{
				p.SetState(170)
				p.Cas_unique()
			}

		}

	case memcached_protocolParserT__29:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(173)
			p.End()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeletion_responseContext is an interface to support dynamic dispatch.
type IDeletion_responseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeletion_responseContext differentiates from other interfaces.
	IsDeletion_responseContext()
}

type Deletion_responseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeletion_responseContext() *Deletion_responseContext {
	var p = new(Deletion_responseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_deletion_response
	return p
}

func (*Deletion_responseContext) IsDeletion_responseContext() {}

func NewDeletion_responseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Deletion_responseContext {
	var p = new(Deletion_responseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_deletion_response

	return p
}

func (s *Deletion_responseContext) GetParser() antlr.Parser { return s.parser }

func (s *Deletion_responseContext) Error_response() IError_responseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IError_responseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IError_responseContext)
}

func (s *Deletion_responseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Deletion_responseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Deletion_responseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterDeletion_response(s)
	}
}

func (s *Deletion_responseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitDeletion_response(s)
	}
}

func (p *memcached_protocolParser) Deletion_response() (localctx IDeletion_responseContext) {
	localctx = NewDeletion_responseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, memcached_protocolParserRULE_deletion_response)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(179)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case memcached_protocolParserT__26, memcached_protocolParserT__27, memcached_protocolParserT__28:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Error_response()
		}

	case memcached_protocolParserT__24:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Match(memcached_protocolParserT__24)
		}

	case memcached_protocolParserT__22:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(178)
			p.Match(memcached_protocolParserT__22)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIncr_or_decr_responseContext is an interface to support dynamic dispatch.
type IIncr_or_decr_responseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncr_or_decr_responseContext differentiates from other interfaces.
	IsIncr_or_decr_responseContext()
}

type Incr_or_decr_responseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncr_or_decr_responseContext() *Incr_or_decr_responseContext {
	var p = new(Incr_or_decr_responseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_incr_or_decr_response
	return p
}

func (*Incr_or_decr_responseContext) IsIncr_or_decr_responseContext() {}

func NewIncr_or_decr_responseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Incr_or_decr_responseContext {
	var p = new(Incr_or_decr_responseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_incr_or_decr_response

	return p
}

func (s *Incr_or_decr_responseContext) GetParser() antlr.Parser { return s.parser }

func (s *Incr_or_decr_responseContext) Error_response() IError_responseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IError_responseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IError_responseContext)
}

func (s *Incr_or_decr_responseContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *Incr_or_decr_responseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Incr_or_decr_responseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Incr_or_decr_responseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterIncr_or_decr_response(s)
	}
}

func (s *Incr_or_decr_responseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitIncr_or_decr_response(s)
	}
}

func (p *memcached_protocolParser) Incr_or_decr_response() (localctx IIncr_or_decr_responseContext) {
	localctx = NewIncr_or_decr_responseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, memcached_protocolParserRULE_incr_or_decr_response)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case memcached_protocolParserT__26, memcached_protocolParserT__27, memcached_protocolParserT__28:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)
			p.Error_response()
		}

	case memcached_protocolParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.Match(memcached_protocolParserT__22)
		}

	case memcached_protocolParserINTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(183)
			p.Match(memcached_protocolParserINTEGER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatistics_responseContext is an interface to support dynamic dispatch.
type IStatistics_responseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatistics_responseContext differentiates from other interfaces.
	IsStatistics_responseContext()
}

type Statistics_responseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatistics_responseContext() *Statistics_responseContext {
	var p = new(Statistics_responseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_statistics_response
	return p
}

func (*Statistics_responseContext) IsStatistics_responseContext() {}

func NewStatistics_responseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statistics_responseContext {
	var p = new(Statistics_responseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_statistics_response

	return p
}

func (s *Statistics_responseContext) GetParser() antlr.Parser { return s.parser }

func (s *Statistics_responseContext) Error_response() IError_responseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IError_responseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IError_responseContext)
}

func (s *Statistics_responseContext) General_statistic() IGeneral_statisticContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeneral_statisticContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeneral_statisticContext)
}

func (s *Statistics_responseContext) Size_statistic() ISize_statisticContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISize_statisticContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISize_statisticContext)
}

func (s *Statistics_responseContext) End() IEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndContext)
}

func (s *Statistics_responseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statistics_responseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statistics_responseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterStatistics_response(s)
	}
}

func (s *Statistics_responseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitStatistics_response(s)
	}
}

func (p *memcached_protocolParser) Statistics_response() (localctx IStatistics_responseContext) {
	localctx = NewStatistics_responseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, memcached_protocolParserRULE_statistics_response)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case memcached_protocolParserT__26, memcached_protocolParserT__27, memcached_protocolParserT__28:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Error_response()
		}

	case memcached_protocolParserT__25:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.General_statistic()
		}

	case memcached_protocolParserINTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.Size_statistic()
		}

	case memcached_protocolParserT__29:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(189)
			p.End()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IError_responseContext is an interface to support dynamic dispatch.
type IError_responseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsError_responseContext differentiates from other interfaces.
	IsError_responseContext()
}

type Error_responseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyError_responseContext() *Error_responseContext {
	var p = new(Error_responseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_error_response
	return p
}

func (*Error_responseContext) IsError_responseContext() {}

func NewError_responseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Error_responseContext {
	var p = new(Error_responseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_error_response

	return p
}

func (s *Error_responseContext) GetParser() antlr.Parser { return s.parser }

func (s *Error_responseContext) General_error() IGeneral_errorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeneral_errorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeneral_errorContext)
}

func (s *Error_responseContext) Client_error_message() IClient_error_messageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClient_error_messageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClient_error_messageContext)
}

func (s *Error_responseContext) Server_error_message() IServer_error_messageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServer_error_messageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServer_error_messageContext)
}

func (s *Error_responseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Error_responseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Error_responseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterError_response(s)
	}
}

func (s *Error_responseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitError_response(s)
	}
}

func (p *memcached_protocolParser) Error_response() (localctx IError_responseContext) {
	localctx = NewError_responseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, memcached_protocolParserRULE_error_response)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case memcached_protocolParserT__26:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.General_error()
		}

	case memcached_protocolParserT__27:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Client_error_message()
		}

	case memcached_protocolParserT__28:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.Server_error_message()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGeneral_statisticContext is an interface to support dynamic dispatch.
type IGeneral_statisticContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeneral_statisticContext differentiates from other interfaces.
	IsGeneral_statisticContext()
}

type General_statisticContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeneral_statisticContext() *General_statisticContext {
	var p = new(General_statisticContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_general_statistic
	return p
}

func (*General_statisticContext) IsGeneral_statisticContext() {}

func NewGeneral_statisticContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *General_statisticContext {
	var p = new(General_statisticContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_general_statistic

	return p
}

func (s *General_statisticContext) GetParser() antlr.Parser { return s.parser }

func (s *General_statisticContext) Statistic_name() IStatistic_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatistic_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatistic_nameContext)
}

func (s *General_statisticContext) Statistic_value() IStatistic_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatistic_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatistic_valueContext)
}

func (s *General_statisticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *General_statisticContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *General_statisticContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterGeneral_statistic(s)
	}
}

func (s *General_statisticContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitGeneral_statistic(s)
	}
}

func (p *memcached_protocolParser) General_statistic() (localctx IGeneral_statisticContext) {
	localctx = NewGeneral_statisticContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, memcached_protocolParserRULE_general_statistic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(memcached_protocolParserT__25)
	}
	{
		p.SetState(198)
		p.Statistic_name()
	}
	{
		p.SetState(199)
		p.Statistic_value()
	}

	return localctx
}

// ISize_statisticContext is an interface to support dynamic dispatch.
type ISize_statisticContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSize_statisticContext differentiates from other interfaces.
	IsSize_statisticContext()
}

type Size_statisticContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySize_statisticContext() *Size_statisticContext {
	var p = new(Size_statisticContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_size_statistic
	return p
}

func (*Size_statisticContext) IsSize_statisticContext() {}

func NewSize_statisticContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Size_statisticContext {
	var p = new(Size_statisticContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_size_statistic

	return p
}

func (s *Size_statisticContext) GetParser() antlr.Parser { return s.parser }

func (s *Size_statisticContext) Size() ISizeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISizeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISizeContext)
}

func (s *Size_statisticContext) Count() ICountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICountContext)
}

func (s *Size_statisticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Size_statisticContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Size_statisticContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterSize_statistic(s)
	}
}

func (s *Size_statisticContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitSize_statistic(s)
	}
}

func (p *memcached_protocolParser) Size_statistic() (localctx ISize_statisticContext) {
	localctx = NewSize_statisticContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, memcached_protocolParserRULE_size_statistic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Size()
	}
	{
		p.SetState(202)
		p.Count()
	}

	return localctx
}

// IGeneral_errorContext is an interface to support dynamic dispatch.
type IGeneral_errorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeneral_errorContext differentiates from other interfaces.
	IsGeneral_errorContext()
}

type General_errorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeneral_errorContext() *General_errorContext {
	var p = new(General_errorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_general_error
	return p
}

func (*General_errorContext) IsGeneral_errorContext() {}

func NewGeneral_errorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *General_errorContext {
	var p = new(General_errorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_general_error

	return p
}

func (s *General_errorContext) GetParser() antlr.Parser { return s.parser }
func (s *General_errorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *General_errorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *General_errorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterGeneral_error(s)
	}
}

func (s *General_errorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitGeneral_error(s)
	}
}

func (p *memcached_protocolParser) General_error() (localctx IGeneral_errorContext) {
	localctx = NewGeneral_errorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, memcached_protocolParserRULE_general_error)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(memcached_protocolParserT__26)
	}

	return localctx
}

// IClient_error_messageContext is an interface to support dynamic dispatch.
type IClient_error_messageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClient_error_messageContext differentiates from other interfaces.
	IsClient_error_messageContext()
}

type Client_error_messageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClient_error_messageContext() *Client_error_messageContext {
	var p = new(Client_error_messageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_client_error_message
	return p
}

func (*Client_error_messageContext) IsClient_error_messageContext() {}

func NewClient_error_messageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Client_error_messageContext {
	var p = new(Client_error_messageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_client_error_message

	return p
}

func (s *Client_error_messageContext) GetParser() antlr.Parser { return s.parser }
func (s *Client_error_messageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Client_error_messageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Client_error_messageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterClient_error_message(s)
	}
}

func (s *Client_error_messageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitClient_error_message(s)
	}
}

func (p *memcached_protocolParser) Client_error_message() (localctx IClient_error_messageContext) {
	localctx = NewClient_error_messageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, memcached_protocolParserRULE_client_error_message)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(memcached_protocolParserT__27)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			p.SetState(207)
			p.MatchWildcard()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IServer_error_messageContext is an interface to support dynamic dispatch.
type IServer_error_messageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServer_error_messageContext differentiates from other interfaces.
	IsServer_error_messageContext()
}

type Server_error_messageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServer_error_messageContext() *Server_error_messageContext {
	var p = new(Server_error_messageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_server_error_message
	return p
}

func (*Server_error_messageContext) IsServer_error_messageContext() {}

func NewServer_error_messageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Server_error_messageContext {
	var p = new(Server_error_messageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_server_error_message

	return p
}

func (s *Server_error_messageContext) GetParser() antlr.Parser { return s.parser }
func (s *Server_error_messageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Server_error_messageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Server_error_messageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterServer_error_message(s)
	}
}

func (s *Server_error_messageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitServer_error_message(s)
	}
}

func (p *memcached_protocolParser) Server_error_message() (localctx IServer_error_messageContext) {
	localctx = NewServer_error_messageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, memcached_protocolParserRULE_server_error_message)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(memcached_protocolParserT__28)
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			p.SetState(213)
			p.MatchWildcard()

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IEndContext is an interface to support dynamic dispatch.
type IEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndContext differentiates from other interfaces.
	IsEndContext()
}

type EndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndContext() *EndContext {
	var p = new(EndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_end
	return p
}

func (*EndContext) IsEndContext() {}

func NewEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndContext {
	var p = new(EndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_end

	return p
}

func (s *EndContext) GetParser() antlr.Parser { return s.parser }
func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterEnd(s)
	}
}

func (s *EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitEnd(s)
	}
}

func (p *memcached_protocolParser) End() (localctx IEndContext) {
	localctx = NewEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, memcached_protocolParserRULE_end)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(memcached_protocolParserT__29)
	}

	return localctx
}

// INoreplyContext is an interface to support dynamic dispatch.
type INoreplyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoreplyContext differentiates from other interfaces.
	IsNoreplyContext()
}

type NoreplyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoreplyContext() *NoreplyContext {
	var p = new(NoreplyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_noreply
	return p
}

func (*NoreplyContext) IsNoreplyContext() {}

func NewNoreplyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoreplyContext {
	var p = new(NoreplyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_noreply

	return p
}

func (s *NoreplyContext) GetParser() antlr.Parser { return s.parser }
func (s *NoreplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoreplyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoreplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterNoreply(s)
	}
}

func (s *NoreplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitNoreply(s)
	}
}

func (p *memcached_protocolParser) Noreply() (localctx INoreplyContext) {
	localctx = NewNoreplyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, memcached_protocolParserRULE_noreply)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(memcached_protocolParserT__30)
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }
func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *memcached_protocolParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, memcached_protocolParserRULE_key)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(222)
	p.MatchWildcard()

	return localctx
}

// IFlagsContext is an interface to support dynamic dispatch.
type IFlagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFlagsContext differentiates from other interfaces.
	IsFlagsContext()
}

type FlagsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlagsContext() *FlagsContext {
	var p = new(FlagsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_flags
	return p
}

func (*FlagsContext) IsFlagsContext() {}

func NewFlagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlagsContext {
	var p = new(FlagsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_flags

	return p
}

func (s *FlagsContext) GetParser() antlr.Parser { return s.parser }

func (s *FlagsContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *FlagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterFlags(s)
	}
}

func (s *FlagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitFlags(s)
	}
}

func (p *memcached_protocolParser) Flags() (localctx IFlagsContext) {
	localctx = NewFlagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, memcached_protocolParserRULE_flags)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(memcached_protocolParserINTEGER)
	}

	return localctx
}

// IExptimeContext is an interface to support dynamic dispatch.
type IExptimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExptimeContext differentiates from other interfaces.
	IsExptimeContext()
}

type ExptimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExptimeContext() *ExptimeContext {
	var p = new(ExptimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_exptime
	return p
}

func (*ExptimeContext) IsExptimeContext() {}

func NewExptimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExptimeContext {
	var p = new(ExptimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_exptime

	return p
}

func (s *ExptimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExptimeContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *ExptimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExptimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExptimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterExptime(s)
	}
}

func (s *ExptimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitExptime(s)
	}
}

func (p *memcached_protocolParser) Exptime() (localctx IExptimeContext) {
	localctx = NewExptimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, memcached_protocolParserRULE_exptime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(memcached_protocolParserINTEGER)
	}

	return localctx
}

// IBytesContext is an interface to support dynamic dispatch.
type IBytesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBytesContext differentiates from other interfaces.
	IsBytesContext()
}

type BytesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBytesContext() *BytesContext {
	var p = new(BytesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_bytes
	return p
}

func (*BytesContext) IsBytesContext() {}

func NewBytesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BytesContext {
	var p = new(BytesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_bytes

	return p
}

func (s *BytesContext) GetParser() antlr.Parser { return s.parser }

func (s *BytesContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *BytesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BytesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BytesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterBytes(s)
	}
}

func (s *BytesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitBytes(s)
	}
}

func (p *memcached_protocolParser) Bytes() (localctx IBytesContext) {
	localctx = NewBytesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, memcached_protocolParserRULE_bytes)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(memcached_protocolParserINTEGER)
	}

	return localctx
}

// ICas_uniqueContext is an interface to support dynamic dispatch.
type ICas_uniqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCas_uniqueContext differentiates from other interfaces.
	IsCas_uniqueContext()
}

type Cas_uniqueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCas_uniqueContext() *Cas_uniqueContext {
	var p = new(Cas_uniqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_cas_unique
	return p
}

func (*Cas_uniqueContext) IsCas_uniqueContext() {}

func NewCas_uniqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cas_uniqueContext {
	var p = new(Cas_uniqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_cas_unique

	return p
}

func (s *Cas_uniqueContext) GetParser() antlr.Parser { return s.parser }

func (s *Cas_uniqueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *Cas_uniqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cas_uniqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cas_uniqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterCas_unique(s)
	}
}

func (s *Cas_uniqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitCas_unique(s)
	}
}

func (p *memcached_protocolParser) Cas_unique() (localctx ICas_uniqueContext) {
	localctx = NewCas_uniqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, memcached_protocolParserRULE_cas_unique)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(memcached_protocolParserINTEGER)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *memcached_protocolParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, memcached_protocolParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(memcached_protocolParserINTEGER)
	}

	return localctx
}

// ITimeContext is an interface to support dynamic dispatch.
type ITimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeContext differentiates from other interfaces.
	IsTimeContext()
}

type TimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeContext() *TimeContext {
	var p = new(TimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_time
	return p
}

func (*TimeContext) IsTimeContext() {}

func NewTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeContext {
	var p = new(TimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_time

	return p
}

func (s *TimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *TimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterTime(s)
	}
}

func (s *TimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitTime(s)
	}
}

func (p *memcached_protocolParser) Time() (localctx ITimeContext) {
	localctx = NewTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, memcached_protocolParserRULE_time)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(memcached_protocolParserINTEGER)
	}

	return localctx
}

// IDelayContext is an interface to support dynamic dispatch.
type IDelayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelayContext differentiates from other interfaces.
	IsDelayContext()
}

type DelayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelayContext() *DelayContext {
	var p = new(DelayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_delay
	return p
}

func (*DelayContext) IsDelayContext() {}

func NewDelayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DelayContext {
	var p = new(DelayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_delay

	return p
}

func (s *DelayContext) GetParser() antlr.Parser { return s.parser }

func (s *DelayContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *DelayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DelayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterDelay(s)
	}
}

func (s *DelayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitDelay(s)
	}
}

func (p *memcached_protocolParser) Delay() (localctx IDelayContext) {
	localctx = NewDelayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, memcached_protocolParserRULE_delay)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(memcached_protocolParserINTEGER)
	}

	return localctx
}

// IVerbosity_levelContext is an interface to support dynamic dispatch.
type IVerbosity_levelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVerbosity_levelContext differentiates from other interfaces.
	IsVerbosity_levelContext()
}

type Verbosity_levelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerbosity_levelContext() *Verbosity_levelContext {
	var p = new(Verbosity_levelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_verbosity_level
	return p
}

func (*Verbosity_levelContext) IsVerbosity_levelContext() {}

func NewVerbosity_levelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Verbosity_levelContext {
	var p = new(Verbosity_levelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_verbosity_level

	return p
}

func (s *Verbosity_levelContext) GetParser() antlr.Parser { return s.parser }

func (s *Verbosity_levelContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *Verbosity_levelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Verbosity_levelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Verbosity_levelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterVerbosity_level(s)
	}
}

func (s *Verbosity_levelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitVerbosity_level(s)
	}
}

func (p *memcached_protocolParser) Verbosity_level() (localctx IVerbosity_levelContext) {
	localctx = NewVerbosity_levelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, memcached_protocolParserRULE_verbosity_level)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(memcached_protocolParserINTEGER)
	}

	return localctx
}

// IStatistic_nameContext is an interface to support dynamic dispatch.
type IStatistic_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatistic_nameContext differentiates from other interfaces.
	IsStatistic_nameContext()
}

type Statistic_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatistic_nameContext() *Statistic_nameContext {
	var p = new(Statistic_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_statistic_name
	return p
}

func (*Statistic_nameContext) IsStatistic_nameContext() {}

func NewStatistic_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statistic_nameContext {
	var p = new(Statistic_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_statistic_name

	return p
}

func (s *Statistic_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Statistic_nameContext) WORD() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserWORD, 0)
}

func (s *Statistic_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statistic_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statistic_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterStatistic_name(s)
	}
}

func (s *Statistic_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitStatistic_name(s)
	}
}

func (p *memcached_protocolParser) Statistic_name() (localctx IStatistic_nameContext) {
	localctx = NewStatistic_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, memcached_protocolParserRULE_statistic_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(memcached_protocolParserWORD)
	}

	return localctx
}

// IStatistic_valueContext is an interface to support dynamic dispatch.
type IStatistic_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatistic_valueContext differentiates from other interfaces.
	IsStatistic_valueContext()
}

type Statistic_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatistic_valueContext() *Statistic_valueContext {
	var p = new(Statistic_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_statistic_value
	return p
}

func (*Statistic_valueContext) IsStatistic_valueContext() {}

func NewStatistic_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statistic_valueContext {
	var p = new(Statistic_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_statistic_value

	return p
}

func (s *Statistic_valueContext) GetParser() antlr.Parser { return s.parser }
func (s *Statistic_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statistic_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statistic_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterStatistic_value(s)
	}
}

func (s *Statistic_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitStatistic_value(s)
	}
}

func (p *memcached_protocolParser) Statistic_value() (localctx IStatistic_valueContext) {
	localctx = NewStatistic_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, memcached_protocolParserRULE_statistic_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(242)
	p.MatchWildcard()

	return localctx
}

// ISizeContext is an interface to support dynamic dispatch.
type ISizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSizeContext differentiates from other interfaces.
	IsSizeContext()
}

type SizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySizeContext() *SizeContext {
	var p = new(SizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_size
	return p
}

func (*SizeContext) IsSizeContext() {}

func NewSizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SizeContext {
	var p = new(SizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_size

	return p
}

func (s *SizeContext) GetParser() antlr.Parser { return s.parser }

func (s *SizeContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *SizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterSize(s)
	}
}

func (s *SizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitSize(s)
	}
}

func (p *memcached_protocolParser) Size() (localctx ISizeContext) {
	localctx = NewSizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, memcached_protocolParserRULE_size)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Match(memcached_protocolParserINTEGER)
	}

	return localctx
}

// ICountContext is an interface to support dynamic dispatch.
type ICountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCountContext differentiates from other interfaces.
	IsCountContext()
}

type CountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountContext() *CountContext {
	var p = new(CountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = memcached_protocolParserRULE_count
	return p
}

func (*CountContext) IsCountContext() {}

func NewCountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountContext {
	var p = new(CountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = memcached_protocolParserRULE_count

	return p
}

func (s *CountContext) GetParser() antlr.Parser { return s.parser }

func (s *CountContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(memcached_protocolParserINTEGER, 0)
}

func (s *CountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.EnterCount(s)
	}
}

func (s *CountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(memcached_protocolListener); ok {
		listenerT.ExitCount(s)
	}
}

func (p *memcached_protocolParser) Count() (localctx ICountContext) {
	localctx = NewCountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, memcached_protocolParserRULE_count)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(memcached_protocolParserINTEGER)
	}

	return localctx
}
