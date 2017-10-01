// Generated from ZOperatorParser.g4 by ANTLR 4.7.

package z // ZOperatorParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 158, 245, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2, 
	3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 
	3, 3, 3, 3, 3, 3, 5, 3, 60, 10, 3, 3, 3, 3, 3, 7, 3, 64, 10, 3, 12, 3, 
	14, 3, 67, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 74, 10, 3, 12, 3, 
	14, 3, 77, 11, 3, 5, 3, 79, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 
	7, 4, 87, 10, 4, 12, 4, 14, 4, 90, 11, 4, 3, 4, 3, 4, 3, 4, 7, 4, 95, 10, 
	4, 12, 4, 14, 4, 98, 11, 4, 3, 4, 3, 4, 3, 4, 7, 4, 103, 10, 4, 12, 4, 
	14, 4, 106, 11, 4, 3, 4, 5, 4, 109, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 114, 
	10, 5, 12, 5, 14, 5, 117, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 
	6, 125, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 134, 10, 
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 144, 10, 
	10, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 150, 10, 11, 3, 11, 3, 11, 3, 12, 
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 
	14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 
	3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 183, 10, 19, 12, 
	19, 14, 19, 186, 11, 19, 3, 19, 3, 19, 5, 19, 190, 10, 19, 3, 19, 3, 19, 
	5, 19, 194, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 202, 
	10, 20, 12, 20, 14, 20, 205, 11, 20, 3, 20, 3, 20, 5, 20, 209, 10, 20, 
	5, 20, 211, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 
	21, 220, 10, 21, 12, 21, 14, 21, 223, 11, 21, 3, 21, 3, 21, 5, 21, 227, 
	10, 21, 3, 21, 3, 21, 5, 21, 231, 10, 21, 3, 22, 3, 22, 3, 22, 7, 22, 236, 
	10, 22, 12, 22, 14, 22, 239, 11, 22, 3, 22, 3, 22, 5, 22, 243, 10, 22, 
	3, 22, 5, 88, 96, 104, 2, 23, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 
	26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 3, 4, 2, 29, 29, 35, 35, 2, 261, 
	2, 49, 3, 2, 2, 2, 4, 78, 3, 2, 2, 2, 6, 108, 3, 2, 2, 2, 8, 110, 3, 2, 
	2, 2, 10, 124, 3, 2, 2, 2, 12, 133, 3, 2, 2, 2, 14, 135, 3, 2, 2, 2, 16, 
	137, 3, 2, 2, 2, 18, 143, 3, 2, 2, 2, 20, 145, 3, 2, 2, 2, 22, 153, 3, 
	2, 2, 2, 24, 157, 3, 2, 2, 2, 26, 161, 3, 2, 2, 2, 28, 165, 3, 2, 2, 2, 
	30, 168, 3, 2, 2, 2, 32, 171, 3, 2, 2, 2, 34, 174, 3, 2, 2, 2, 36, 193, 
	3, 2, 2, 2, 38, 210, 3, 2, 2, 2, 40, 230, 3, 2, 2, 2, 42, 232, 3, 2, 2, 
	2, 44, 48, 7, 6, 2, 2, 45, 48, 5, 4, 3, 2, 46, 48, 5, 6, 4, 2, 47, 44, 
	3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 
	49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 49, 3, 
	2, 2, 2, 52, 53, 7, 2, 2, 3, 53, 3, 3, 2, 2, 2, 54, 55, 7, 3, 2, 2, 55, 
	56, 7, 36, 2, 2, 56, 57, 7, 70, 2, 2, 57, 59, 7, 32, 2, 2, 58, 60, 5, 8, 
	5, 2, 59, 58, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 65, 
	7, 20, 2, 2, 62, 64, 5, 6, 4, 2, 63, 62, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2, 
	65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 79, 3, 2, 2, 2, 67, 65, 3, 
	2, 2, 2, 68, 69, 7, 3, 2, 2, 69, 70, 7, 36, 2, 2, 70, 71, 7, 70, 2, 2, 
	71, 75, 7, 20, 2, 2, 72, 74, 5, 6, 4, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3, 
	2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 
	75, 3, 2, 2, 2, 78, 54, 3, 2, 2, 2, 78, 68, 3, 2, 2, 2, 79, 5, 3, 2, 2, 
	2, 80, 81, 7, 3, 2, 2, 81, 82, 5, 10, 6, 2, 82, 83, 7, 20, 2, 2, 83, 109, 
	3, 2, 2, 2, 84, 88, 7, 5, 2, 2, 85, 87, 11, 2, 2, 2, 86, 85, 3, 2, 2, 2, 
	87, 90, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 91, 3, 
	2, 2, 2, 90, 88, 3, 2, 2, 2, 91, 109, 7, 20, 2, 2, 92, 96, 7, 4, 2, 2, 
	93, 95, 11, 2, 2, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 97, 3, 
	2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 99, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 
	109, 7, 20, 2, 2, 100, 104, 7, 3, 2, 2, 101, 103, 11, 2, 2, 2, 102, 101, 
	3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 104, 102, 3, 2, 
	2, 2, 105, 107, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 107, 109, 7, 20, 2, 2, 
	108, 80, 3, 2, 2, 2, 108, 84, 3, 2, 2, 2, 108, 92, 3, 2, 2, 2, 108, 100, 
	3, 2, 2, 2, 109, 7, 3, 2, 2, 2, 110, 115, 7, 70, 2, 2, 111, 112, 7, 41, 
	2, 2, 112, 114, 7, 70, 2, 2, 113, 111, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 
	115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 9, 3, 2, 2, 2, 117, 115, 
	3, 2, 2, 2, 118, 119, 7, 34, 2, 2, 119, 125, 5, 18, 10, 2, 120, 121, 7, 
	26, 2, 2, 121, 125, 5, 12, 7, 2, 122, 123, 7, 27, 2, 2, 123, 125, 5, 12, 
	7, 2, 124, 118, 3, 2, 2, 2, 124, 120, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 
	125, 11, 3, 2, 2, 2, 126, 134, 5, 20, 11, 2, 127, 134, 5, 22, 12, 2, 128, 
	129, 5, 14, 8, 2, 129, 130, 5, 16, 9, 2, 130, 131, 5, 24, 13, 2, 131, 134, 
	3, 2, 2, 2, 132, 134, 5, 26, 14, 2, 133, 126, 3, 2, 2, 2, 133, 127, 3, 
	2, 2, 2, 133, 128, 3, 2, 2, 2, 133, 132, 3, 2, 2, 2, 134, 13, 3, 2, 2, 
	2, 135, 136, 7, 7, 2, 2, 136, 15, 3, 2, 2, 2, 137, 138, 9, 2, 2, 2, 138, 
	17, 3, 2, 2, 2, 139, 144, 5, 20, 11, 2, 140, 144, 5, 22, 12, 2, 141, 144, 
	5, 24, 13, 2, 142, 144, 5, 26, 14, 2, 143, 139, 3, 2, 2, 2, 143, 140, 3, 
	2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 142, 3, 2, 2, 2, 144, 19, 3, 2, 2, 
	2, 145, 149, 7, 9, 2, 2, 146, 150, 5, 36, 19, 2, 147, 148, 7, 31, 2, 2, 
	148, 150, 7, 49, 2, 2, 149, 146, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 
	151, 3, 2, 2, 2, 151, 152, 7, 10, 2, 2, 152, 21, 3, 2, 2, 2, 153, 154, 
	7, 9, 2, 2, 154, 155, 5, 38, 20, 2, 155, 156, 7, 10, 2, 2, 156, 23, 3, 
	2, 2, 2, 157, 158, 7, 9, 2, 2, 158, 159, 5, 40, 21, 2, 159, 160, 7, 10, 
	2, 2, 160, 25, 3, 2, 2, 2, 161, 162, 7, 9, 2, 2, 162, 163, 5, 42, 22, 2, 
	163, 164, 7, 10, 2, 2, 164, 27, 3, 2, 2, 2, 165, 166, 7, 49, 2, 2, 166, 
	167, 7, 70, 2, 2, 167, 29, 3, 2, 2, 2, 168, 169, 7, 50, 2, 2, 169, 170, 
	7, 70, 2, 2, 170, 31, 3, 2, 2, 2, 171, 172, 7, 49, 2, 2, 172, 173, 7, 70, 
	2, 2, 173, 33, 3, 2, 2, 2, 174, 175, 7, 50, 2, 2, 175, 176, 7, 70, 2, 2, 
	176, 35, 3, 2, 2, 2, 177, 178, 7, 70, 2, 2, 178, 194, 7, 49, 2, 2, 179, 
	184, 7, 70, 2, 2, 180, 183, 5, 28, 15, 2, 181, 183, 5, 30, 16, 2, 182, 
	180, 3, 2, 2, 2, 182, 181, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2, 184, 182, 
	3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 189, 3, 2, 2, 2, 186, 184, 3, 2, 
	2, 2, 187, 190, 5, 32, 17, 2, 188, 190, 5, 34, 18, 2, 189, 187, 3, 2, 2, 
	2, 189, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192, 7, 49, 2, 2, 192, 
	194, 3, 2, 2, 2, 193, 177, 3, 2, 2, 2, 193, 179, 3, 2, 2, 2, 194, 37, 3, 
	2, 2, 2, 195, 196, 7, 49, 2, 2, 196, 211, 7, 70, 2, 2, 197, 198, 7, 49, 
	2, 2, 198, 203, 7, 70, 2, 2, 199, 202, 5, 28, 15, 2, 200, 202, 5, 30, 16, 
	2, 201, 199, 3, 2, 2, 2, 201, 200, 3, 2, 2, 2, 202, 205, 3, 2, 2, 2, 203, 
	201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 208, 3, 2, 2, 2, 205, 203, 
	3, 2, 2, 2, 206, 209, 5, 32, 17, 2, 207, 209, 5, 34, 18, 2, 208, 206, 3, 
	2, 2, 2, 208, 207, 3, 2, 2, 2, 209, 211, 3, 2, 2, 2, 210, 195, 3, 2, 2, 
	2, 210, 197, 3, 2, 2, 2, 211, 39, 3, 2, 2, 2, 212, 213, 7, 49, 2, 2, 213, 
	214, 7, 70, 2, 2, 214, 231, 7, 49, 2, 2, 215, 216, 7, 49, 2, 2, 216, 221, 
	7, 70, 2, 2, 217, 220, 5, 28, 15, 2, 218, 220, 5, 30, 16, 2, 219, 217, 
	3, 2, 2, 2, 219, 218, 3, 2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3, 2, 
	2, 2, 221, 222, 3, 2, 2, 2, 222, 226, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 
	224, 227, 5, 32, 17, 2, 225, 227, 5, 34, 18, 2, 226, 224, 3, 2, 2, 2, 226, 
	225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 229, 7, 49, 2, 2, 229, 231, 
	3, 2, 2, 2, 230, 212, 3, 2, 2, 2, 230, 215, 3, 2, 2, 2, 231, 41, 3, 2, 
	2, 2, 232, 237, 7, 70, 2, 2, 233, 236, 5, 28, 15, 2, 234, 236, 5, 30, 16, 
	2, 235, 233, 3, 2, 2, 2, 235, 234, 3, 2, 2, 2, 236, 239, 3, 2, 2, 2, 237, 
	235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 242, 3, 2, 2, 2, 239, 237, 
	3, 2, 2, 2, 240, 243, 5, 32, 17, 2, 241, 243, 5, 34, 18, 2, 242, 240, 3, 
	2, 2, 2, 242, 241, 3, 2, 2, 2, 243, 43, 3, 2, 2, 2, 32, 47, 49, 59, 65, 
	75, 78, 88, 96, 104, 108, 115, 124, 133, 143, 149, 182, 184, 189, 193, 
	201, 203, 208, 210, 219, 221, 226, 230, 235, 237, 242,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\u2500'", "'\u250C'", "'\u2577'", "", "", "", "'\u0028'", "'\u0029'", 
	"'\u005B'", "'\u005D'", "'\u007B'", "'\u007D'", "'\u2989'", "'\u298A'", 
	"'\u27EA'", "'\u27EB'", "'\u2550'", "'\u2514'", "", "", "", "'else'", "'false'", 
	"'function'", "'generic'", "'if'", "'leftassoc'", "'let'", "'\u2119'", 
	"'parents'", "'pre'", "'relation'", "'rightassoc'", "'section'", "'then'", 
	"'true'", "':'", "'=='", "','", "'::='", "'|'", "'\u0026'", "'\u005C'", 
	"'/'", "'.'", "';'", "'_'", "',,'", "'='", "", "'\u2200'", "'\u2981'", 
	"'\u2203'", "", "'\u21D4'", "'\u21D2'", "'\u2228'", "'\u2227'", "'\u00AC'", 
	"'\u2208'", "'\u2A21'", "'\u00D7'", "'\u03B8'", "'\u03BB'", "'\u03BC'", 
	"'\u2A1F'", "'\u2A20'", "", "", "'\u0394'", "'\u039E'", "'\uD835\uDD38'", 
	"'\u2115'", "'\u2032'", "'\u0021'", "'\u003F'", "'\u2197'", "'\u2199'", 
	"'\u2198'", "'\u2196'", "'\u22A2'", "'\u29F9'", "'\u002B'", "", "'\u2194'", 
	"'\u2192'", "'\u2260'", "'\u2209'", "'\u2205'", "'\u2286'", "'\u2282'", 
	"'\u222A'", "'\u2229'", "'\u2216'", "'\u2296'", "'\u22C3'", "'\u22C2'", 
	"'\uD835\uDD3D'", "'\u21A6'", "'\u2A3E'", "'\u2218'", "'\u25C1'", "'\u25B7'", 
	"'\u2A64'", "'\u2A65'", "'\u223C'", "'\u2987'", "'\u2988'", "'\u2295'", 
	"'\u21F8'", "'\u2914'", "'\u21A3'", "'\u2900'", "'\u21A0'", "'\u2916'", 
	"'\u21FB'", "'\u2915'", "'\u2124'", "'\u002D'", "'\u2212'", "'\u2264'", 
	"'\u003C'", "'\u2265'", "'\u003E'", "'\u002A'", "'\u0023'", "'\u2040'", 
	"'\u21BF'", "'\u21BE'", "'\u27E8'", "'\u27E9'", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"'\u2028'",
}
var symbolicNames = []string{
	"", "ZED", "SCH", "AX", "TEXT", "NUMERAL", "STROKE", "LEFT_PARENTHESIS", 
	"RIGHT_PARENTHESIS", "LEFT_SQUARE_BRACKET", "RIGHT_SQUARE_BRACKET", "LEFT_CURLY_BRACKET", 
	"RIGHT_CURLY_BRACKET", "LEFT_BINDING_BRACKET", "RIGHT_BINDING_BRACKET", 
	"LEFT_DOUBLE_ANGLE_BRACKET", "RIGHT_DOUBLE_ANGLE_BRACKET", "GEN", "END", 
	"WS", "NL", "IGNORE_NL", "ELSE", "FALSE", "FUNCTION", "GENERIC", "IF", 
	"LEFTASSOC", "LET", "POWERSET", "PARENTS", "PRE_KEY", "RELATION", "RIGHTASSOC", 
	"SECTION", "THEN", "TRUE", "COLON", "DEFINE_EQUAL", "COMMA", "FREE_EQUALS", 
	"VERTICAL_LINE", "AMPERSAND", "REVERSE_SOLIDUS", "SOLIDUS", "FULL_STOP", 
	"SEMICOLON", "ARGUMENT", "LIST", "EQUALS_SIGN", "CONJECTURE", "FOR_ALL", 
	"SPOT", "THERE_EXISTS", "UNIQUE_EXISTS", "LEFT_RIGHT_DOUBLE_ARROW", "RIGHTWARDS_DOUBLE_ARROW", 
	"LOGICAL_OR", "LOGICAL_AND", "NOT_SIGN", "ELEMENT_OF", "SCHEMA_PROJECTION", 
	"MULTIPLICATION_SIGN", "GREEK_SMALL_LETTER_THETA", "GREEK_SMALL_LETTER_LAMBDA", 
	"GREEK_SMALL_LETTER_MU", "SCHEMA_COMPOSITION", "SCHEMA_PIPING", "NAME", 
	"WORDPART", "GREEK_CAPITAL_LETTER_DELTA", "GREEK_CAPITAL_LETTER_XI", "MATHEMATICAL_DOUBLE_STRUCK_CAPITAL_A", 
	"DOUBLE_STRUCK_CAPITAL_N", "MODIFIER_LETTER_PRIME", "EXCLAMATION_MARK", 
	"QUESTION_MARK", "NORTH_EAST_ARROW", "SOUTH_WEST_ARROW", "SOUTH_EAST_ARROW", 
	"NORTH_WEST_ARROW", "RIGHT_TACK", "BIG_REVERSE_SOLIDUS", "PLUS_SIGN", "MATHEMATICAL_TOOLKIT_SYMBOLS", 
	"LEFT_RIGHT_ARROW", "RIGHTWARDS_ARROW", "NOT_EQUAL_TO", "NOT_AN_ELEMENT_OF", 
	"EMPTY_SET", "SUBSET_OF_OR_EQUAL_TO", "SUBSET_OF", "UNION", "INTERSECTION", 
	"SET_MINUS", "CIRCLED_MINUS", "N_ARY_UNION", "N_ARY_INTERSECTION", "MATHEMATICAL_DOUBLE_STRUCK_CAPITAL_F", 
	"RIGHTWARDS_ARROW_FROM_BAR", "RELATIONAL_COMPOSITION", "RING_OPERATOR", 
	"WHITE_LEFT_POINTING_TRIANGLE", "WHITE_RIGHT_POINTING_TRIANGLE", "DOMAIN_ANTIRESTRICTION", 
	"RANGE_ANTIRESTRICTION", "TILDE_OPERATOR", "LEFT_IMAGE_BRACKET", "RIGHT_IMAGE_BRACKET", 
	"CIRCLED_PLUS", "RIGHTWARDS_ARROW_WITH_VERTICAL_STROKE", "RIGHTWARDS_ARROW_WITH_TAIL_WITH_VERTICAL_STROKE", 
	"RIGHTWARDS_ARROW_WITH_TAIL", "RIGHTWARDS_TWO_HEADED_ARROW_WITH_VERTICAL_STROKE", 
	"RIGHTWARDS_TWO_HEADED_ARROW", "RIGHTWARDS_TWO_HEADED_ARROW_WITH_TAIL", 
	"RIGHTWARDS_ARROW_WITH_DOUBLE_VERTICAL_STROKE", "RIGHTWARDS_ARROW_WITH_TAIL_WITH_DOUBLE_VERTICAL_STROKE", 
	"DOUBLE_STRUCK_CAPITAL_Z", "HYPHEN_MINUS", "MINUS_SIGN", "LESS_THAN_OR_EQUAL_TO", 
	"LESS_THAN_SIGN", "GREATER_THAN_OR_EQUAL_TO", "GREATER_THAN_SIGN", "ASTERISK", 
	"NUMBER_SIGN", "CHARACTER_TIE", "UPWARDS_HARPOON_WITH_BARB_LEFTWARDS", 
	"UPWARDS_HARPOON_WITH_BARB_RIGHTWARDS", "LEFT_ANGLE_BRACKET", "RIGHT_ANGLE_BRACKET", 
	"ID0", "ID1", "ID", "PREP", "PRE", "POSTP", "POST", "IP", "I", "LP", "L", 
	"ELP", "EL", "ERP", "ER", "SRP", "SR", "EREP", "ERE", "SREP", "SRE", "ES", 
	"SS", "UNKNOWN", "NLCHAR",
}

var ruleNames = []string{
	"specification", "section", "paragraph", "formals", "operatorTemplate", 
	"categoryTemplate", "prec", "assoc", "template", "prefixTemplate", "postfixTemplate", 
	"infixTemplate", "nofixTemplate", "optArgName", "optListName", "argName", 
	"listName", "prefixName", "postfixName", "infixName", "nofixName",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ZOperatorParser struct {
	*antlr.BaseParser
}

func NewZOperatorParser(input antlr.TokenStream) *ZOperatorParser {
	this := new(ZOperatorParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ZOperatorParser.g4"

	return this
}

// ZOperatorParser tokens.
const (
	ZOperatorParserEOF = antlr.TokenEOF
	ZOperatorParserZED = 1
	ZOperatorParserSCH = 2
	ZOperatorParserAX = 3
	ZOperatorParserTEXT = 4
	ZOperatorParserNUMERAL = 5
	ZOperatorParserSTROKE = 6
	ZOperatorParserLEFT_PARENTHESIS = 7
	ZOperatorParserRIGHT_PARENTHESIS = 8
	ZOperatorParserLEFT_SQUARE_BRACKET = 9
	ZOperatorParserRIGHT_SQUARE_BRACKET = 10
	ZOperatorParserLEFT_CURLY_BRACKET = 11
	ZOperatorParserRIGHT_CURLY_BRACKET = 12
	ZOperatorParserLEFT_BINDING_BRACKET = 13
	ZOperatorParserRIGHT_BINDING_BRACKET = 14
	ZOperatorParserLEFT_DOUBLE_ANGLE_BRACKET = 15
	ZOperatorParserRIGHT_DOUBLE_ANGLE_BRACKET = 16
	ZOperatorParserGEN = 17
	ZOperatorParserEND = 18
	ZOperatorParserWS = 19
	ZOperatorParserNL = 20
	ZOperatorParserIGNORE_NL = 21
	ZOperatorParserELSE = 22
	ZOperatorParserFALSE = 23
	ZOperatorParserFUNCTION = 24
	ZOperatorParserGENERIC = 25
	ZOperatorParserIF = 26
	ZOperatorParserLEFTASSOC = 27
	ZOperatorParserLET = 28
	ZOperatorParserPOWERSET = 29
	ZOperatorParserPARENTS = 30
	ZOperatorParserPRE_KEY = 31
	ZOperatorParserRELATION = 32
	ZOperatorParserRIGHTASSOC = 33
	ZOperatorParserSECTION = 34
	ZOperatorParserTHEN = 35
	ZOperatorParserTRUE = 36
	ZOperatorParserCOLON = 37
	ZOperatorParserDEFINE_EQUAL = 38
	ZOperatorParserCOMMA = 39
	ZOperatorParserFREE_EQUALS = 40
	ZOperatorParserVERTICAL_LINE = 41
	ZOperatorParserAMPERSAND = 42
	ZOperatorParserREVERSE_SOLIDUS = 43
	ZOperatorParserSOLIDUS = 44
	ZOperatorParserFULL_STOP = 45
	ZOperatorParserSEMICOLON = 46
	ZOperatorParserARGUMENT = 47
	ZOperatorParserLIST = 48
	ZOperatorParserEQUALS_SIGN = 49
	ZOperatorParserCONJECTURE = 50
	ZOperatorParserFOR_ALL = 51
	ZOperatorParserSPOT = 52
	ZOperatorParserTHERE_EXISTS = 53
	ZOperatorParserUNIQUE_EXISTS = 54
	ZOperatorParserLEFT_RIGHT_DOUBLE_ARROW = 55
	ZOperatorParserRIGHTWARDS_DOUBLE_ARROW = 56
	ZOperatorParserLOGICAL_OR = 57
	ZOperatorParserLOGICAL_AND = 58
	ZOperatorParserNOT_SIGN = 59
	ZOperatorParserELEMENT_OF = 60
	ZOperatorParserSCHEMA_PROJECTION = 61
	ZOperatorParserMULTIPLICATION_SIGN = 62
	ZOperatorParserGREEK_SMALL_LETTER_THETA = 63
	ZOperatorParserGREEK_SMALL_LETTER_LAMBDA = 64
	ZOperatorParserGREEK_SMALL_LETTER_MU = 65
	ZOperatorParserSCHEMA_COMPOSITION = 66
	ZOperatorParserSCHEMA_PIPING = 67
	ZOperatorParserNAME = 68
	ZOperatorParserWORDPART = 69
	ZOperatorParserGREEK_CAPITAL_LETTER_DELTA = 70
	ZOperatorParserGREEK_CAPITAL_LETTER_XI = 71
	ZOperatorParserMATHEMATICAL_DOUBLE_STRUCK_CAPITAL_A = 72
	ZOperatorParserDOUBLE_STRUCK_CAPITAL_N = 73
	ZOperatorParserMODIFIER_LETTER_PRIME = 74
	ZOperatorParserEXCLAMATION_MARK = 75
	ZOperatorParserQUESTION_MARK = 76
	ZOperatorParserNORTH_EAST_ARROW = 77
	ZOperatorParserSOUTH_WEST_ARROW = 78
	ZOperatorParserSOUTH_EAST_ARROW = 79
	ZOperatorParserNORTH_WEST_ARROW = 80
	ZOperatorParserRIGHT_TACK = 81
	ZOperatorParserBIG_REVERSE_SOLIDUS = 82
	ZOperatorParserPLUS_SIGN = 83
	ZOperatorParserMATHEMATICAL_TOOLKIT_SYMBOLS = 84
	ZOperatorParserLEFT_RIGHT_ARROW = 85
	ZOperatorParserRIGHTWARDS_ARROW = 86
	ZOperatorParserNOT_EQUAL_TO = 87
	ZOperatorParserNOT_AN_ELEMENT_OF = 88
	ZOperatorParserEMPTY_SET = 89
	ZOperatorParserSUBSET_OF_OR_EQUAL_TO = 90
	ZOperatorParserSUBSET_OF = 91
	ZOperatorParserUNION = 92
	ZOperatorParserINTERSECTION = 93
	ZOperatorParserSET_MINUS = 94
	ZOperatorParserCIRCLED_MINUS = 95
	ZOperatorParserN_ARY_UNION = 96
	ZOperatorParserN_ARY_INTERSECTION = 97
	ZOperatorParserMATHEMATICAL_DOUBLE_STRUCK_CAPITAL_F = 98
	ZOperatorParserRIGHTWARDS_ARROW_FROM_BAR = 99
	ZOperatorParserRELATIONAL_COMPOSITION = 100
	ZOperatorParserRING_OPERATOR = 101
	ZOperatorParserWHITE_LEFT_POINTING_TRIANGLE = 102
	ZOperatorParserWHITE_RIGHT_POINTING_TRIANGLE = 103
	ZOperatorParserDOMAIN_ANTIRESTRICTION = 104
	ZOperatorParserRANGE_ANTIRESTRICTION = 105
	ZOperatorParserTILDE_OPERATOR = 106
	ZOperatorParserLEFT_IMAGE_BRACKET = 107
	ZOperatorParserRIGHT_IMAGE_BRACKET = 108
	ZOperatorParserCIRCLED_PLUS = 109
	ZOperatorParserRIGHTWARDS_ARROW_WITH_VERTICAL_STROKE = 110
	ZOperatorParserRIGHTWARDS_ARROW_WITH_TAIL_WITH_VERTICAL_STROKE = 111
	ZOperatorParserRIGHTWARDS_ARROW_WITH_TAIL = 112
	ZOperatorParserRIGHTWARDS_TWO_HEADED_ARROW_WITH_VERTICAL_STROKE = 113
	ZOperatorParserRIGHTWARDS_TWO_HEADED_ARROW = 114
	ZOperatorParserRIGHTWARDS_TWO_HEADED_ARROW_WITH_TAIL = 115
	ZOperatorParserRIGHTWARDS_ARROW_WITH_DOUBLE_VERTICAL_STROKE = 116
	ZOperatorParserRIGHTWARDS_ARROW_WITH_TAIL_WITH_DOUBLE_VERTICAL_STROKE = 117
	ZOperatorParserDOUBLE_STRUCK_CAPITAL_Z = 118
	ZOperatorParserHYPHEN_MINUS = 119
	ZOperatorParserMINUS_SIGN = 120
	ZOperatorParserLESS_THAN_OR_EQUAL_TO = 121
	ZOperatorParserLESS_THAN_SIGN = 122
	ZOperatorParserGREATER_THAN_OR_EQUAL_TO = 123
	ZOperatorParserGREATER_THAN_SIGN = 124
	ZOperatorParserASTERISK = 125
	ZOperatorParserNUMBER_SIGN = 126
	ZOperatorParserCHARACTER_TIE = 127
	ZOperatorParserUPWARDS_HARPOON_WITH_BARB_LEFTWARDS = 128
	ZOperatorParserUPWARDS_HARPOON_WITH_BARB_RIGHTWARDS = 129
	ZOperatorParserLEFT_ANGLE_BRACKET = 130
	ZOperatorParserRIGHT_ANGLE_BRACKET = 131
	ZOperatorParserID0 = 132
	ZOperatorParserID1 = 133
	ZOperatorParserID = 134
	ZOperatorParserPREP = 135
	ZOperatorParserPRE = 136
	ZOperatorParserPOSTP = 137
	ZOperatorParserPOST = 138
	ZOperatorParserIP = 139
	ZOperatorParserI = 140
	ZOperatorParserLP = 141
	ZOperatorParserL = 142
	ZOperatorParserELP = 143
	ZOperatorParserEL = 144
	ZOperatorParserERP = 145
	ZOperatorParserER = 146
	ZOperatorParserSRP = 147
	ZOperatorParserSR = 148
	ZOperatorParserEREP = 149
	ZOperatorParserERE = 150
	ZOperatorParserSREP = 151
	ZOperatorParserSRE = 152
	ZOperatorParserES = 153
	ZOperatorParserSS = 154
	ZOperatorParserUNKNOWN = 155
	ZOperatorParserNLCHAR = 156
)

// ZOperatorParser rules.
const (
	ZOperatorParserRULE_specification = 0
	ZOperatorParserRULE_section = 1
	ZOperatorParserRULE_paragraph = 2
	ZOperatorParserRULE_formals = 3
	ZOperatorParserRULE_operatorTemplate = 4
	ZOperatorParserRULE_categoryTemplate = 5
	ZOperatorParserRULE_prec = 6
	ZOperatorParserRULE_assoc = 7
	ZOperatorParserRULE_template = 8
	ZOperatorParserRULE_prefixTemplate = 9
	ZOperatorParserRULE_postfixTemplate = 10
	ZOperatorParserRULE_infixTemplate = 11
	ZOperatorParserRULE_nofixTemplate = 12
	ZOperatorParserRULE_optArgName = 13
	ZOperatorParserRULE_optListName = 14
	ZOperatorParserRULE_argName = 15
	ZOperatorParserRULE_listName = 16
	ZOperatorParserRULE_prefixName = 17
	ZOperatorParserRULE_postfixName = 18
	ZOperatorParserRULE_infixName = 19
	ZOperatorParserRULE_nofixName = 20
)

// ISpecificationContext is an interface to support dynamic dispatch.
type ISpecificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificationContext differentiates from other interfaces.
	IsSpecificationContext()
}

type SpecificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificationContext() *SpecificationContext {
	var p = new(SpecificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_specification
	return p
}

func (*SpecificationContext) IsSpecificationContext() {}

func NewSpecificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificationContext {
	var p = new(SpecificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_specification

	return p
}

func (s *SpecificationContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificationContext) EOF() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserEOF, 0)
}

func (s *SpecificationContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(ZOperatorParserTEXT)
}

func (s *SpecificationContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(ZOperatorParserTEXT, i)
}

func (s *SpecificationContext) AllSection() []ISectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISectionContext)(nil)).Elem())
	var tst = make([]ISectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISectionContext)
		}
	}

	return tst
}

func (s *SpecificationContext) Section(i int) ISectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *SpecificationContext) AllParagraph() []IParagraphContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParagraphContext)(nil)).Elem())
	var tst = make([]IParagraphContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParagraphContext)
		}
	}

	return tst
}

func (s *SpecificationContext) Paragraph(i int) IParagraphContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParagraphContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParagraphContext)
}

func (s *SpecificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SpecificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterSpecification(s)
	}
}

func (s *SpecificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitSpecification(s)
	}
}

func (s *SpecificationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitSpecification(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) Specification() (localctx ISpecificationContext) {
	localctx = NewSpecificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZOperatorParserRULE_specification)
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
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ZOperatorParserZED) | (1 << ZOperatorParserSCH) | (1 << ZOperatorParserAX) | (1 << ZOperatorParserTEXT))) != 0) {
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(42)
				p.Match(ZOperatorParserTEXT)
			}


		case 2:
			{
				p.SetState(43)
				p.Section()
			}


		case 3:
			{
				p.SetState(44)
				p.Paragraph()
			}

		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(ZOperatorParserEOF)
	}



	return localctx
}


// ISectionContext is an interface to support dynamic dispatch.
type ISectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionContext differentiates from other interfaces.
	IsSectionContext()
}

type SectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionContext() *SectionContext {
	var p = new(SectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) CopyFrom(ctx *SectionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type BaseSectionContext struct {
	*SectionContext
}

func NewBaseSectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BaseSectionContext {
	var p = new(BaseSectionContext)

	p.SectionContext = NewEmptySectionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SectionContext))

	return p
}

func (s *BaseSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseSectionContext) ZED() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserZED, 0)
}

func (s *BaseSectionContext) SECTION() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserSECTION, 0)
}

func (s *BaseSectionContext) NAME() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, 0)
}

func (s *BaseSectionContext) END() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserEND, 0)
}

func (s *BaseSectionContext) AllParagraph() []IParagraphContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParagraphContext)(nil)).Elem())
	var tst = make([]IParagraphContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParagraphContext)
		}
	}

	return tst
}

func (s *BaseSectionContext) Paragraph(i int) IParagraphContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParagraphContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParagraphContext)
}


func (s *BaseSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterBaseSection(s)
	}
}

func (s *BaseSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitBaseSection(s)
	}
}

func (s *BaseSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitBaseSection(s)

	default:
		return t.VisitChildren(s)
	}
}


type InheritingSectionContext struct {
	*SectionContext
}

func NewInheritingSectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InheritingSectionContext {
	var p = new(InheritingSectionContext)

	p.SectionContext = NewEmptySectionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SectionContext))

	return p
}

func (s *InheritingSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InheritingSectionContext) ZED() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserZED, 0)
}

func (s *InheritingSectionContext) SECTION() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserSECTION, 0)
}

func (s *InheritingSectionContext) NAME() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, 0)
}

func (s *InheritingSectionContext) PARENTS() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserPARENTS, 0)
}

func (s *InheritingSectionContext) END() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserEND, 0)
}

func (s *InheritingSectionContext) Formals() IFormalsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalsContext)
}

func (s *InheritingSectionContext) AllParagraph() []IParagraphContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParagraphContext)(nil)).Elem())
	var tst = make([]IParagraphContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParagraphContext)
		}
	}

	return tst
}

func (s *InheritingSectionContext) Paragraph(i int) IParagraphContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParagraphContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParagraphContext)
}


func (s *InheritingSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterInheritingSection(s)
	}
}

func (s *InheritingSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitInheritingSection(s)
	}
}

func (s *InheritingSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitInheritingSection(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *ZOperatorParser) Section() (localctx ISectionContext) {
	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ZOperatorParserRULE_section)
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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInheritingSectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(ZOperatorParserZED)
		}
		{
			p.SetState(53)
			p.Match(ZOperatorParserSECTION)
		}
		{
			p.SetState(54)
			p.Match(ZOperatorParserNAME)
		}
		{
			p.SetState(55)
			p.Match(ZOperatorParserPARENTS)
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == ZOperatorParserNAME {
			{
				p.SetState(56)
				p.Formals()
			}

		}
		{
			p.SetState(59)
			p.Match(ZOperatorParserEND)
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(60)
					p.Paragraph()
				}


			}
			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}


	case 2:
		localctx = NewBaseSectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(ZOperatorParserZED)
		}
		{
			p.SetState(67)
			p.Match(ZOperatorParserSECTION)
		}
		{
			p.SetState(68)
			p.Match(ZOperatorParserNAME)
		}
		{
			p.SetState(69)
			p.Match(ZOperatorParserEND)
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(70)
					p.Paragraph()
				}


			}
			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}

	}


	return localctx
}


// IParagraphContext is an interface to support dynamic dispatch.
type IParagraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParagraphContext differentiates from other interfaces.
	IsParagraphContext()
}

type ParagraphContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParagraphContext() *ParagraphContext {
	var p = new(ParagraphContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_paragraph
	return p
}

func (*ParagraphContext) IsParagraphContext() {}

func NewParagraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParagraphContext {
	var p = new(ParagraphContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_paragraph

	return p
}

func (s *ParagraphContext) GetParser() antlr.Parser { return s.parser }

func (s *ParagraphContext) CopyFrom(ctx *ParagraphContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ParagraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParagraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type NONOperatorTemplateParagraphContext struct {
	*ParagraphContext
}

func NewNONOperatorTemplateParagraphContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NONOperatorTemplateParagraphContext {
	var p = new(NONOperatorTemplateParagraphContext)

	p.ParagraphContext = NewEmptyParagraphContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParagraphContext))

	return p
}

func (s *NONOperatorTemplateParagraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NONOperatorTemplateParagraphContext) ZED() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserZED, 0)
}

func (s *NONOperatorTemplateParagraphContext) END() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserEND, 0)
}


func (s *NONOperatorTemplateParagraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterNONOperatorTemplateParagraph(s)
	}
}

func (s *NONOperatorTemplateParagraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitNONOperatorTemplateParagraph(s)
	}
}

func (s *NONOperatorTemplateParagraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitNONOperatorTemplateParagraph(s)

	default:
		return t.VisitChildren(s)
	}
}


type OperatorTemplateParagraphContext struct {
	*ParagraphContext
}

func NewOperatorTemplateParagraphContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperatorTemplateParagraphContext {
	var p = new(OperatorTemplateParagraphContext)

	p.ParagraphContext = NewEmptyParagraphContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParagraphContext))

	return p
}

func (s *OperatorTemplateParagraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorTemplateParagraphContext) ZED() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserZED, 0)
}

func (s *OperatorTemplateParagraphContext) OperatorTemplate() IOperatorTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorTemplateContext)
}

func (s *OperatorTemplateParagraphContext) END() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserEND, 0)
}


func (s *OperatorTemplateParagraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterOperatorTemplateParagraph(s)
	}
}

func (s *OperatorTemplateParagraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitOperatorTemplateParagraph(s)
	}
}

func (s *OperatorTemplateParagraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitOperatorTemplateParagraph(s)

	default:
		return t.VisitChildren(s)
	}
}


type SchemaDefinitionParagraphContext struct {
	*ParagraphContext
}

func NewSchemaDefinitionParagraphContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SchemaDefinitionParagraphContext {
	var p = new(SchemaDefinitionParagraphContext)

	p.ParagraphContext = NewEmptyParagraphContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParagraphContext))

	return p
}

func (s *SchemaDefinitionParagraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaDefinitionParagraphContext) SCH() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserSCH, 0)
}

func (s *SchemaDefinitionParagraphContext) END() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserEND, 0)
}


func (s *SchemaDefinitionParagraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterSchemaDefinitionParagraph(s)
	}
}

func (s *SchemaDefinitionParagraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitSchemaDefinitionParagraph(s)
	}
}

func (s *SchemaDefinitionParagraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitSchemaDefinitionParagraph(s)

	default:
		return t.VisitChildren(s)
	}
}


type AxiomaticDescriptionParagraphContext struct {
	*ParagraphContext
}

func NewAxiomaticDescriptionParagraphContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AxiomaticDescriptionParagraphContext {
	var p = new(AxiomaticDescriptionParagraphContext)

	p.ParagraphContext = NewEmptyParagraphContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParagraphContext))

	return p
}

func (s *AxiomaticDescriptionParagraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AxiomaticDescriptionParagraphContext) AX() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserAX, 0)
}

func (s *AxiomaticDescriptionParagraphContext) END() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserEND, 0)
}


func (s *AxiomaticDescriptionParagraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterAxiomaticDescriptionParagraph(s)
	}
}

func (s *AxiomaticDescriptionParagraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitAxiomaticDescriptionParagraph(s)
	}
}

func (s *AxiomaticDescriptionParagraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitAxiomaticDescriptionParagraph(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *ZOperatorParser) Paragraph() (localctx IParagraphContext) {
	localctx = NewParagraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ZOperatorParserRULE_paragraph)

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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOperatorTemplateParagraphContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Match(ZOperatorParserZED)
		}
		{
			p.SetState(79)
			p.OperatorTemplate()
		}
		{
			p.SetState(80)
			p.Match(ZOperatorParserEND)
		}


	case 2:
		localctx = NewAxiomaticDescriptionParagraphContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Match(ZOperatorParserAX)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				p.SetState(83)
				p.MatchWildcard()



			}
			p.SetState(88)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}
		{
			p.SetState(89)
			p.Match(ZOperatorParserEND)
		}


	case 3:
		localctx = NewSchemaDefinitionParagraphContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Match(ZOperatorParserSCH)
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				p.SetState(91)
				p.MatchWildcard()



			}
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
		}
		{
			p.SetState(97)
			p.Match(ZOperatorParserEND)
		}


	case 4:
		localctx = NewNONOperatorTemplateParagraphContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.Match(ZOperatorParserZED)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

		for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1+1 {
				p.SetState(99)
				p.MatchWildcard()



			}
			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
		}
		{
			p.SetState(105)
			p.Match(ZOperatorParserEND)
		}

	}


	return localctx
}


// IFormalsContext is an interface to support dynamic dispatch.
type IFormalsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalsContext differentiates from other interfaces.
	IsFormalsContext()
}

type FormalsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalsContext() *FormalsContext {
	var p = new(FormalsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_formals
	return p
}

func (*FormalsContext) IsFormalsContext() {}

func NewFormalsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalsContext {
	var p = new(FormalsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_formals

	return p
}

func (s *FormalsContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalsContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(ZOperatorParserNAME)
}

func (s *FormalsContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, i)
}

func (s *FormalsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ZOperatorParserCOMMA)
}

func (s *FormalsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ZOperatorParserCOMMA, i)
}

func (s *FormalsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FormalsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterFormals(s)
	}
}

func (s *FormalsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitFormals(s)
	}
}

func (s *FormalsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitFormals(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) Formals() (localctx IFormalsContext) {
	localctx = NewFormalsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ZOperatorParserRULE_formals)
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
		p.SetState(108)
		p.Match(ZOperatorParserNAME)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ZOperatorParserCOMMA {
		{
			p.SetState(109)
			p.Match(ZOperatorParserCOMMA)
		}
		{
			p.SetState(110)
			p.Match(ZOperatorParserNAME)
		}


		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IOperatorTemplateContext is an interface to support dynamic dispatch.
type IOperatorTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorTemplateContext differentiates from other interfaces.
	IsOperatorTemplateContext()
}

type OperatorTemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorTemplateContext() *OperatorTemplateContext {
	var p = new(OperatorTemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_operatorTemplate
	return p
}

func (*OperatorTemplateContext) IsOperatorTemplateContext() {}

func NewOperatorTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorTemplateContext {
	var p = new(OperatorTemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_operatorTemplate

	return p
}

func (s *OperatorTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorTemplateContext) CopyFrom(ctx *OperatorTemplateContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OperatorTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorTemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type RelationOperatorTemplateContext struct {
	*OperatorTemplateContext
}

func NewRelationOperatorTemplateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationOperatorTemplateContext {
	var p = new(RelationOperatorTemplateContext)

	p.OperatorTemplateContext = NewEmptyOperatorTemplateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperatorTemplateContext))

	return p
}

func (s *RelationOperatorTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationOperatorTemplateContext) RELATION() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserRELATION, 0)
}

func (s *RelationOperatorTemplateContext) Template() ITemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateContext)
}


func (s *RelationOperatorTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterRelationOperatorTemplate(s)
	}
}

func (s *RelationOperatorTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitRelationOperatorTemplate(s)
	}
}

func (s *RelationOperatorTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitRelationOperatorTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}


type GenericOperatorTemplateContext struct {
	*OperatorTemplateContext
}

func NewGenericOperatorTemplateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GenericOperatorTemplateContext {
	var p = new(GenericOperatorTemplateContext)

	p.OperatorTemplateContext = NewEmptyOperatorTemplateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperatorTemplateContext))

	return p
}

func (s *GenericOperatorTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericOperatorTemplateContext) GENERIC() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserGENERIC, 0)
}

func (s *GenericOperatorTemplateContext) CategoryTemplate() ICategoryTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICategoryTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICategoryTemplateContext)
}


func (s *GenericOperatorTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterGenericOperatorTemplate(s)
	}
}

func (s *GenericOperatorTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitGenericOperatorTemplate(s)
	}
}

func (s *GenericOperatorTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitGenericOperatorTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}


type FunctionOperatorTemplateContext struct {
	*OperatorTemplateContext
}

func NewFunctionOperatorTemplateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionOperatorTemplateContext {
	var p = new(FunctionOperatorTemplateContext)

	p.OperatorTemplateContext = NewEmptyOperatorTemplateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperatorTemplateContext))

	return p
}

func (s *FunctionOperatorTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionOperatorTemplateContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserFUNCTION, 0)
}

func (s *FunctionOperatorTemplateContext) CategoryTemplate() ICategoryTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICategoryTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICategoryTemplateContext)
}


func (s *FunctionOperatorTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterFunctionOperatorTemplate(s)
	}
}

func (s *FunctionOperatorTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitFunctionOperatorTemplate(s)
	}
}

func (s *FunctionOperatorTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitFunctionOperatorTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *ZOperatorParser) OperatorTemplate() (localctx IOperatorTemplateContext) {
	localctx = NewOperatorTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ZOperatorParserRULE_operatorTemplate)

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

	p.SetState(122)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZOperatorParserRELATION:
		localctx = NewRelationOperatorTemplateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(ZOperatorParserRELATION)
		}
		{
			p.SetState(117)
			p.Template()
		}


	case ZOperatorParserFUNCTION:
		localctx = NewFunctionOperatorTemplateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(ZOperatorParserFUNCTION)
		}
		{
			p.SetState(119)
			p.CategoryTemplate()
		}


	case ZOperatorParserGENERIC:
		localctx = NewGenericOperatorTemplateContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.Match(ZOperatorParserGENERIC)
		}
		{
			p.SetState(121)
			p.CategoryTemplate()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ICategoryTemplateContext is an interface to support dynamic dispatch.
type ICategoryTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCategoryTemplateContext differentiates from other interfaces.
	IsCategoryTemplateContext()
}

type CategoryTemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCategoryTemplateContext() *CategoryTemplateContext {
	var p = new(CategoryTemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_categoryTemplate
	return p
}

func (*CategoryTemplateContext) IsCategoryTemplateContext() {}

func NewCategoryTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CategoryTemplateContext {
	var p = new(CategoryTemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_categoryTemplate

	return p
}

func (s *CategoryTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *CategoryTemplateContext) PrefixTemplate() IPrefixTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixTemplateContext)
}

func (s *CategoryTemplateContext) PostfixTemplate() IPostfixTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfixTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfixTemplateContext)
}

func (s *CategoryTemplateContext) Prec() IPrecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrecContext)
}

func (s *CategoryTemplateContext) Assoc() IAssocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssocContext)
}

func (s *CategoryTemplateContext) InfixTemplate() IInfixTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixTemplateContext)
}

func (s *CategoryTemplateContext) NofixTemplate() INofixTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INofixTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INofixTemplateContext)
}

func (s *CategoryTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CategoryTemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CategoryTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterCategoryTemplate(s)
	}
}

func (s *CategoryTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitCategoryTemplate(s)
	}
}

func (s *CategoryTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitCategoryTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) CategoryTemplate() (localctx ICategoryTemplateContext) {
	localctx = NewCategoryTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ZOperatorParserRULE_categoryTemplate)

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

	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.PrefixTemplate()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.PostfixTemplate()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.Prec()
		}
		{
			p.SetState(127)
			p.Assoc()
		}
		{
			p.SetState(128)
			p.InfixTemplate()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.NofixTemplate()
		}

	}


	return localctx
}


// IPrecContext is an interface to support dynamic dispatch.
type IPrecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrecContext differentiates from other interfaces.
	IsPrecContext()
}

type PrecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrecContext() *PrecContext {
	var p = new(PrecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_prec
	return p
}

func (*PrecContext) IsPrecContext() {}

func NewPrecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrecContext {
	var p = new(PrecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_prec

	return p
}

func (s *PrecContext) GetParser() antlr.Parser { return s.parser }

func (s *PrecContext) NUMERAL() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNUMERAL, 0)
}

func (s *PrecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PrecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterPrec(s)
	}
}

func (s *PrecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitPrec(s)
	}
}

func (s *PrecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitPrec(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) Prec() (localctx IPrecContext) {
	localctx = NewPrecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ZOperatorParserRULE_prec)

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
		p.SetState(133)
		p.Match(ZOperatorParserNUMERAL)
	}



	return localctx
}


// IAssocContext is an interface to support dynamic dispatch.
type IAssocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssocContext differentiates from other interfaces.
	IsAssocContext()
}

type AssocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssocContext() *AssocContext {
	var p = new(AssocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_assoc
	return p
}

func (*AssocContext) IsAssocContext() {}

func NewAssocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssocContext {
	var p = new(AssocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_assoc

	return p
}

func (s *AssocContext) GetParser() antlr.Parser { return s.parser }

func (s *AssocContext) LEFTASSOC() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserLEFTASSOC, 0)
}

func (s *AssocContext) RIGHTASSOC() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserRIGHTASSOC, 0)
}

func (s *AssocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AssocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterAssoc(s)
	}
}

func (s *AssocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitAssoc(s)
	}
}

func (s *AssocContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitAssoc(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) Assoc() (localctx IAssocContext) {
	localctx = NewAssocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ZOperatorParserRULE_assoc)
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
	p.SetState(135)
	_la = p.GetTokenStream().LA(1)

	if !(_la == ZOperatorParserLEFTASSOC || _la == ZOperatorParserRIGHTASSOC) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}



	return localctx
}


// ITemplateContext is an interface to support dynamic dispatch.
type ITemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateContext differentiates from other interfaces.
	IsTemplateContext()
}

type TemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateContext() *TemplateContext {
	var p = new(TemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_template
	return p
}

func (*TemplateContext) IsTemplateContext() {}

func NewTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateContext {
	var p = new(TemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_template

	return p
}

func (s *TemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateContext) PrefixTemplate() IPrefixTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixTemplateContext)
}

func (s *TemplateContext) PostfixTemplate() IPostfixTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfixTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfixTemplateContext)
}

func (s *TemplateContext) InfixTemplate() IInfixTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixTemplateContext)
}

func (s *TemplateContext) NofixTemplate() INofixTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INofixTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INofixTemplateContext)
}

func (s *TemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterTemplate(s)
	}
}

func (s *TemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitTemplate(s)
	}
}

func (s *TemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) Template() (localctx ITemplateContext) {
	localctx = NewTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ZOperatorParserRULE_template)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.PrefixTemplate()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.PostfixTemplate()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
			p.InfixTemplate()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(140)
			p.NofixTemplate()
		}

	}


	return localctx
}


// IPrefixTemplateContext is an interface to support dynamic dispatch.
type IPrefixTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixTemplateContext differentiates from other interfaces.
	IsPrefixTemplateContext()
}

type PrefixTemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixTemplateContext() *PrefixTemplateContext {
	var p = new(PrefixTemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_prefixTemplate
	return p
}

func (*PrefixTemplateContext) IsPrefixTemplateContext() {}

func NewPrefixTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixTemplateContext {
	var p = new(PrefixTemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_prefixTemplate

	return p
}

func (s *PrefixTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixTemplateContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserLEFT_PARENTHESIS, 0)
}

func (s *PrefixTemplateContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserRIGHT_PARENTHESIS, 0)
}

func (s *PrefixTemplateContext) PrefixName() IPrefixNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixNameContext)
}

func (s *PrefixTemplateContext) POWERSET() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserPOWERSET, 0)
}

func (s *PrefixTemplateContext) ARGUMENT() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserARGUMENT, 0)
}

func (s *PrefixTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixTemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PrefixTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterPrefixTemplate(s)
	}
}

func (s *PrefixTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitPrefixTemplate(s)
	}
}

func (s *PrefixTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitPrefixTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) PrefixTemplate() (localctx IPrefixTemplateContext) {
	localctx = NewPrefixTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ZOperatorParserRULE_prefixTemplate)

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
		p.SetState(143)
		p.Match(ZOperatorParserLEFT_PARENTHESIS)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZOperatorParserNAME:
		{
			p.SetState(144)
			p.PrefixName()
		}


	case ZOperatorParserPOWERSET:
		{
			p.SetState(145)
			p.Match(ZOperatorParserPOWERSET)
		}
		{
			p.SetState(146)
			p.Match(ZOperatorParserARGUMENT)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(149)
		p.Match(ZOperatorParserRIGHT_PARENTHESIS)
	}



	return localctx
}


// IPostfixTemplateContext is an interface to support dynamic dispatch.
type IPostfixTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostfixTemplateContext differentiates from other interfaces.
	IsPostfixTemplateContext()
}

type PostfixTemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixTemplateContext() *PostfixTemplateContext {
	var p = new(PostfixTemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_postfixTemplate
	return p
}

func (*PostfixTemplateContext) IsPostfixTemplateContext() {}

func NewPostfixTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixTemplateContext {
	var p = new(PostfixTemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_postfixTemplate

	return p
}

func (s *PostfixTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixTemplateContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserLEFT_PARENTHESIS, 0)
}

func (s *PostfixTemplateContext) PostfixName() IPostfixNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfixNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfixNameContext)
}

func (s *PostfixTemplateContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserRIGHT_PARENTHESIS, 0)
}

func (s *PostfixTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixTemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PostfixTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterPostfixTemplate(s)
	}
}

func (s *PostfixTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitPostfixTemplate(s)
	}
}

func (s *PostfixTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitPostfixTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) PostfixTemplate() (localctx IPostfixTemplateContext) {
	localctx = NewPostfixTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ZOperatorParserRULE_postfixTemplate)

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
		p.Match(ZOperatorParserLEFT_PARENTHESIS)
	}
	{
		p.SetState(152)
		p.PostfixName()
	}
	{
		p.SetState(153)
		p.Match(ZOperatorParserRIGHT_PARENTHESIS)
	}



	return localctx
}


// IInfixTemplateContext is an interface to support dynamic dispatch.
type IInfixTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfixTemplateContext differentiates from other interfaces.
	IsInfixTemplateContext()
}

type InfixTemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfixTemplateContext() *InfixTemplateContext {
	var p = new(InfixTemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_infixTemplate
	return p
}

func (*InfixTemplateContext) IsInfixTemplateContext() {}

func NewInfixTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfixTemplateContext {
	var p = new(InfixTemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_infixTemplate

	return p
}

func (s *InfixTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *InfixTemplateContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserLEFT_PARENTHESIS, 0)
}

func (s *InfixTemplateContext) InfixName() IInfixNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixNameContext)
}

func (s *InfixTemplateContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserRIGHT_PARENTHESIS, 0)
}

func (s *InfixTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfixTemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InfixTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterInfixTemplate(s)
	}
}

func (s *InfixTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitInfixTemplate(s)
	}
}

func (s *InfixTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitInfixTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) InfixTemplate() (localctx IInfixTemplateContext) {
	localctx = NewInfixTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ZOperatorParserRULE_infixTemplate)

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
		p.SetState(155)
		p.Match(ZOperatorParserLEFT_PARENTHESIS)
	}
	{
		p.SetState(156)
		p.InfixName()
	}
	{
		p.SetState(157)
		p.Match(ZOperatorParserRIGHT_PARENTHESIS)
	}



	return localctx
}


// INofixTemplateContext is an interface to support dynamic dispatch.
type INofixTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNofixTemplateContext differentiates from other interfaces.
	IsNofixTemplateContext()
}

type NofixTemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNofixTemplateContext() *NofixTemplateContext {
	var p = new(NofixTemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_nofixTemplate
	return p
}

func (*NofixTemplateContext) IsNofixTemplateContext() {}

func NewNofixTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NofixTemplateContext {
	var p = new(NofixTemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_nofixTemplate

	return p
}

func (s *NofixTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *NofixTemplateContext) LEFT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserLEFT_PARENTHESIS, 0)
}

func (s *NofixTemplateContext) NofixName() INofixNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INofixNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INofixNameContext)
}

func (s *NofixTemplateContext) RIGHT_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserRIGHT_PARENTHESIS, 0)
}

func (s *NofixTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NofixTemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NofixTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterNofixTemplate(s)
	}
}

func (s *NofixTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitNofixTemplate(s)
	}
}

func (s *NofixTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitNofixTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) NofixTemplate() (localctx INofixTemplateContext) {
	localctx = NewNofixTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ZOperatorParserRULE_nofixTemplate)

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
		p.SetState(159)
		p.Match(ZOperatorParserLEFT_PARENTHESIS)
	}
	{
		p.SetState(160)
		p.NofixName()
	}
	{
		p.SetState(161)
		p.Match(ZOperatorParserRIGHT_PARENTHESIS)
	}



	return localctx
}


// IOptArgNameContext is an interface to support dynamic dispatch.
type IOptArgNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptArgNameContext differentiates from other interfaces.
	IsOptArgNameContext()
}

type OptArgNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptArgNameContext() *OptArgNameContext {
	var p = new(OptArgNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_optArgName
	return p
}

func (*OptArgNameContext) IsOptArgNameContext() {}

func NewOptArgNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptArgNameContext {
	var p = new(OptArgNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_optArgName

	return p
}

func (s *OptArgNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptArgNameContext) ARGUMENT() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserARGUMENT, 0)
}

func (s *OptArgNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, 0)
}

func (s *OptArgNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptArgNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OptArgNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterOptArgName(s)
	}
}

func (s *OptArgNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitOptArgName(s)
	}
}

func (s *OptArgNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitOptArgName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) OptArgName() (localctx IOptArgNameContext) {
	localctx = NewOptArgNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ZOperatorParserRULE_optArgName)

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
		p.SetState(163)
		p.Match(ZOperatorParserARGUMENT)
	}
	{
		p.SetState(164)
		p.Match(ZOperatorParserNAME)
	}



	return localctx
}


// IOptListNameContext is an interface to support dynamic dispatch.
type IOptListNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptListNameContext differentiates from other interfaces.
	IsOptListNameContext()
}

type OptListNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptListNameContext() *OptListNameContext {
	var p = new(OptListNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_optListName
	return p
}

func (*OptListNameContext) IsOptListNameContext() {}

func NewOptListNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptListNameContext {
	var p = new(OptListNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_optListName

	return p
}

func (s *OptListNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptListNameContext) LIST() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserLIST, 0)
}

func (s *OptListNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, 0)
}

func (s *OptListNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptListNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OptListNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterOptListName(s)
	}
}

func (s *OptListNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitOptListName(s)
	}
}

func (s *OptListNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitOptListName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) OptListName() (localctx IOptListNameContext) {
	localctx = NewOptListNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ZOperatorParserRULE_optListName)

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
		p.SetState(166)
		p.Match(ZOperatorParserLIST)
	}
	{
		p.SetState(167)
		p.Match(ZOperatorParserNAME)
	}



	return localctx
}


// IArgNameContext is an interface to support dynamic dispatch.
type IArgNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgNameContext differentiates from other interfaces.
	IsArgNameContext()
}

type ArgNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgNameContext() *ArgNameContext {
	var p = new(ArgNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_argName
	return p
}

func (*ArgNameContext) IsArgNameContext() {}

func NewArgNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgNameContext {
	var p = new(ArgNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_argName

	return p
}

func (s *ArgNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgNameContext) ARGUMENT() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserARGUMENT, 0)
}

func (s *ArgNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, 0)
}

func (s *ArgNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArgNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterArgName(s)
	}
}

func (s *ArgNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitArgName(s)
	}
}

func (s *ArgNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitArgName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) ArgName() (localctx IArgNameContext) {
	localctx = NewArgNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ZOperatorParserRULE_argName)

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
		p.SetState(169)
		p.Match(ZOperatorParserARGUMENT)
	}
	{
		p.SetState(170)
		p.Match(ZOperatorParserNAME)
	}



	return localctx
}


// IListNameContext is an interface to support dynamic dispatch.
type IListNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListNameContext differentiates from other interfaces.
	IsListNameContext()
}

type ListNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListNameContext() *ListNameContext {
	var p = new(ListNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_listName
	return p
}

func (*ListNameContext) IsListNameContext() {}

func NewListNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListNameContext {
	var p = new(ListNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_listName

	return p
}

func (s *ListNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ListNameContext) LIST() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserLIST, 0)
}

func (s *ListNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, 0)
}

func (s *ListNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ListNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterListName(s)
	}
}

func (s *ListNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitListName(s)
	}
}

func (s *ListNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitListName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) ListName() (localctx IListNameContext) {
	localctx = NewListNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ZOperatorParserRULE_listName)

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
		p.SetState(172)
		p.Match(ZOperatorParserLIST)
	}
	{
		p.SetState(173)
		p.Match(ZOperatorParserNAME)
	}



	return localctx
}


// IPrefixNameContext is an interface to support dynamic dispatch.
type IPrefixNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixNameContext differentiates from other interfaces.
	IsPrefixNameContext()
}

type PrefixNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixNameContext() *PrefixNameContext {
	var p = new(PrefixNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_prefixName
	return p
}

func (*PrefixNameContext) IsPrefixNameContext() {}

func NewPrefixNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixNameContext {
	var p = new(PrefixNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_prefixName

	return p
}

func (s *PrefixNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, 0)
}

func (s *PrefixNameContext) ARGUMENT() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserARGUMENT, 0)
}

func (s *PrefixNameContext) ArgName() IArgNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgNameContext)
}

func (s *PrefixNameContext) ListName() IListNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListNameContext)
}

func (s *PrefixNameContext) AllOptArgName() []IOptArgNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptArgNameContext)(nil)).Elem())
	var tst = make([]IOptArgNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptArgNameContext)
		}
	}

	return tst
}

func (s *PrefixNameContext) OptArgName(i int) IOptArgNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptArgNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptArgNameContext)
}

func (s *PrefixNameContext) AllOptListName() []IOptListNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptListNameContext)(nil)).Elem())
	var tst = make([]IOptListNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptListNameContext)
		}
	}

	return tst
}

func (s *PrefixNameContext) OptListName(i int) IOptListNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptListNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptListNameContext)
}

func (s *PrefixNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PrefixNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterPrefixName(s)
	}
}

func (s *PrefixNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitPrefixName(s)
	}
}

func (s *PrefixNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitPrefixName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) PrefixName() (localctx IPrefixNameContext) {
	localctx = NewPrefixNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ZOperatorParserRULE_prefixName)

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

	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.Match(ZOperatorParserNAME)
		}
		{
			p.SetState(176)
			p.Match(ZOperatorParserARGUMENT)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Match(ZOperatorParserNAME)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(180)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case ZOperatorParserARGUMENT:
					{
						p.SetState(178)
						p.OptArgName()
					}


				case ZOperatorParserLIST:
					{
						p.SetState(179)
						p.OptListName()
					}



				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			}
			p.SetState(184)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ZOperatorParserARGUMENT:
			{
				p.SetState(185)
				p.ArgName()
			}


		case ZOperatorParserLIST:
			{
				p.SetState(186)
				p.ListName()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(189)
			p.Match(ZOperatorParserARGUMENT)
		}

	}


	return localctx
}


// IPostfixNameContext is an interface to support dynamic dispatch.
type IPostfixNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostfixNameContext differentiates from other interfaces.
	IsPostfixNameContext()
}

type PostfixNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixNameContext() *PostfixNameContext {
	var p = new(PostfixNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_postfixName
	return p
}

func (*PostfixNameContext) IsPostfixNameContext() {}

func NewPostfixNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixNameContext {
	var p = new(PostfixNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_postfixName

	return p
}

func (s *PostfixNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixNameContext) ARGUMENT() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserARGUMENT, 0)
}

func (s *PostfixNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, 0)
}

func (s *PostfixNameContext) ArgName() IArgNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgNameContext)
}

func (s *PostfixNameContext) ListName() IListNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListNameContext)
}

func (s *PostfixNameContext) AllOptArgName() []IOptArgNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptArgNameContext)(nil)).Elem())
	var tst = make([]IOptArgNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptArgNameContext)
		}
	}

	return tst
}

func (s *PostfixNameContext) OptArgName(i int) IOptArgNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptArgNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptArgNameContext)
}

func (s *PostfixNameContext) AllOptListName() []IOptListNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptListNameContext)(nil)).Elem())
	var tst = make([]IOptListNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptListNameContext)
		}
	}

	return tst
}

func (s *PostfixNameContext) OptListName(i int) IOptListNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptListNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptListNameContext)
}

func (s *PostfixNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PostfixNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterPostfixName(s)
	}
}

func (s *PostfixNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitPostfixName(s)
	}
}

func (s *PostfixNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitPostfixName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) PostfixName() (localctx IPostfixNameContext) {
	localctx = NewPostfixNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ZOperatorParserRULE_postfixName)

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

	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Match(ZOperatorParserARGUMENT)
		}
		{
			p.SetState(194)
			p.Match(ZOperatorParserNAME)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.Match(ZOperatorParserARGUMENT)
		}
		{
			p.SetState(196)
			p.Match(ZOperatorParserNAME)
		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(199)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case ZOperatorParserARGUMENT:
					{
						p.SetState(197)
						p.OptArgName()
					}


				case ZOperatorParserLIST:
					{
						p.SetState(198)
						p.OptListName()
					}



				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			}
			p.SetState(203)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ZOperatorParserARGUMENT:
			{
				p.SetState(204)
				p.ArgName()
			}


		case ZOperatorParserLIST:
			{
				p.SetState(205)
				p.ListName()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}


	return localctx
}


// IInfixNameContext is an interface to support dynamic dispatch.
type IInfixNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfixNameContext differentiates from other interfaces.
	IsInfixNameContext()
}

type InfixNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfixNameContext() *InfixNameContext {
	var p = new(InfixNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_infixName
	return p
}

func (*InfixNameContext) IsInfixNameContext() {}

func NewInfixNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfixNameContext {
	var p = new(InfixNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_infixName

	return p
}

func (s *InfixNameContext) GetParser() antlr.Parser { return s.parser }

func (s *InfixNameContext) AllARGUMENT() []antlr.TerminalNode {
	return s.GetTokens(ZOperatorParserARGUMENT)
}

func (s *InfixNameContext) ARGUMENT(i int) antlr.TerminalNode {
	return s.GetToken(ZOperatorParserARGUMENT, i)
}

func (s *InfixNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, 0)
}

func (s *InfixNameContext) ArgName() IArgNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgNameContext)
}

func (s *InfixNameContext) ListName() IListNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListNameContext)
}

func (s *InfixNameContext) AllOptArgName() []IOptArgNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptArgNameContext)(nil)).Elem())
	var tst = make([]IOptArgNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptArgNameContext)
		}
	}

	return tst
}

func (s *InfixNameContext) OptArgName(i int) IOptArgNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptArgNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptArgNameContext)
}

func (s *InfixNameContext) AllOptListName() []IOptListNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptListNameContext)(nil)).Elem())
	var tst = make([]IOptListNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptListNameContext)
		}
	}

	return tst
}

func (s *InfixNameContext) OptListName(i int) IOptListNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptListNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptListNameContext)
}

func (s *InfixNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfixNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InfixNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterInfixName(s)
	}
}

func (s *InfixNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitInfixName(s)
	}
}

func (s *InfixNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitInfixName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) InfixName() (localctx IInfixNameContext) {
	localctx = NewInfixNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ZOperatorParserRULE_infixName)

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

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(ZOperatorParserARGUMENT)
		}
		{
			p.SetState(211)
			p.Match(ZOperatorParserNAME)
		}
		{
			p.SetState(212)
			p.Match(ZOperatorParserARGUMENT)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)
			p.Match(ZOperatorParserARGUMENT)
		}
		{
			p.SetState(214)
			p.Match(ZOperatorParserNAME)
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(217)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case ZOperatorParserARGUMENT:
					{
						p.SetState(215)
						p.OptArgName()
					}


				case ZOperatorParserLIST:
					{
						p.SetState(216)
						p.OptListName()
					}



				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			}
			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ZOperatorParserARGUMENT:
			{
				p.SetState(222)
				p.ArgName()
			}


		case ZOperatorParserLIST:
			{
				p.SetState(223)
				p.ListName()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(226)
			p.Match(ZOperatorParserARGUMENT)
		}

	}


	return localctx
}


// INofixNameContext is an interface to support dynamic dispatch.
type INofixNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNofixNameContext differentiates from other interfaces.
	IsNofixNameContext()
}

type NofixNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNofixNameContext() *NofixNameContext {
	var p = new(NofixNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZOperatorParserRULE_nofixName
	return p
}

func (*NofixNameContext) IsNofixNameContext() {}

func NewNofixNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NofixNameContext {
	var p = new(NofixNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZOperatorParserRULE_nofixName

	return p
}

func (s *NofixNameContext) GetParser() antlr.Parser { return s.parser }

func (s *NofixNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(ZOperatorParserNAME, 0)
}

func (s *NofixNameContext) ArgName() IArgNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgNameContext)
}

func (s *NofixNameContext) ListName() IListNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListNameContext)
}

func (s *NofixNameContext) AllOptArgName() []IOptArgNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptArgNameContext)(nil)).Elem())
	var tst = make([]IOptArgNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptArgNameContext)
		}
	}

	return tst
}

func (s *NofixNameContext) OptArgName(i int) IOptArgNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptArgNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptArgNameContext)
}

func (s *NofixNameContext) AllOptListName() []IOptListNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptListNameContext)(nil)).Elem())
	var tst = make([]IOptListNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptListNameContext)
		}
	}

	return tst
}

func (s *NofixNameContext) OptListName(i int) IOptListNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptListNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptListNameContext)
}

func (s *NofixNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NofixNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NofixNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.EnterNofixName(s)
	}
}

func (s *NofixNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZOperatorParserListener); ok {
		listenerT.ExitNofixName(s)
	}
}

func (s *NofixNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZOperatorParserVisitor:
		return t.VisitNofixName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ZOperatorParser) NofixName() (localctx INofixNameContext) {
	localctx = NewNofixNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ZOperatorParserRULE_nofixName)

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
		p.SetState(230)
		p.Match(ZOperatorParserNAME)
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(233)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case ZOperatorParserARGUMENT:
				{
					p.SetState(231)
					p.OptArgName()
				}


			case ZOperatorParserLIST:
				{
					p.SetState(232)
					p.OptListName()
				}



			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZOperatorParserARGUMENT:
		{
			p.SetState(238)
			p.ArgName()
		}


	case ZOperatorParserLIST:
		{
			p.SetState(239)
			p.ListName()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


