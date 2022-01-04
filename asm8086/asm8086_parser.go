// Code generated from asm8086.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asm8086 // asm8086
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 292,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3, 2, 3,
	2, 3, 2, 7, 2, 80, 10, 2, 12, 2, 14, 2, 83, 11, 2, 3, 2, 3, 2, 7, 2, 87,
	10, 2, 12, 2, 14, 2, 90, 11, 2, 3, 3, 5, 3, 93, 10, 3, 3, 3, 3, 3, 5, 3,
	97, 10, 3, 3, 3, 5, 3, 100, 10, 3, 3, 4, 5, 4, 103, 10, 4, 3, 4, 3, 4,
	5, 4, 107, 10, 4, 3, 5, 3, 5, 5, 5, 111, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	5, 6, 129, 10, 6, 3, 7, 5, 7, 132, 10, 7, 3, 7, 3, 7, 3, 7, 3, 8, 5, 8,
	138, 10, 8, 3, 8, 3, 8, 3, 8, 3, 9, 5, 9, 144, 10, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 5, 10, 151, 10, 10, 3, 11, 3, 11, 5, 11, 155, 10, 11, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 177,
	10, 17, 12, 17, 14, 17, 180, 11, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17,
	186, 10, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 194, 10,
	19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 7, 25, 212, 10, 25, 12, 25, 14,
	25, 215, 11, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 7, 27, 222, 10, 27,
	12, 27, 14, 27, 225, 11, 27, 3, 28, 3, 28, 3, 28, 7, 28, 230, 10, 28, 12,
	28, 14, 28, 233, 11, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 246, 10, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 266, 10, 29, 3, 30, 5, 30, 269, 10,
	30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34,
	3, 35, 5, 35, 282, 10, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 2, 2, 39, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 2, 5, 3, 2, 8, 9, 3, 2, 11, 14, 3, 2, 17, 19,
	2, 305, 2, 88, 3, 2, 2, 2, 4, 92, 3, 2, 2, 2, 6, 102, 3, 2, 2, 2, 8, 108,
	3, 2, 2, 2, 10, 128, 3, 2, 2, 2, 12, 131, 3, 2, 2, 2, 14, 137, 3, 2, 2,
	2, 16, 143, 3, 2, 2, 2, 18, 148, 3, 2, 2, 2, 20, 152, 3, 2, 2, 2, 22, 156,
	3, 2, 2, 2, 24, 159, 3, 2, 2, 2, 26, 162, 3, 2, 2, 2, 28, 165, 3, 2, 2,
	2, 30, 169, 3, 2, 2, 2, 32, 185, 3, 2, 2, 2, 34, 187, 3, 2, 2, 2, 36, 193,
	3, 2, 2, 2, 38, 195, 3, 2, 2, 2, 40, 197, 3, 2, 2, 2, 42, 199, 3, 2, 2,
	2, 44, 202, 3, 2, 2, 2, 46, 205, 3, 2, 2, 2, 48, 208, 3, 2, 2, 2, 50, 216,
	3, 2, 2, 2, 52, 218, 3, 2, 2, 2, 54, 226, 3, 2, 2, 2, 56, 265, 3, 2, 2,
	2, 58, 268, 3, 2, 2, 2, 60, 272, 3, 2, 2, 2, 62, 274, 3, 2, 2, 2, 64, 276,
	3, 2, 2, 2, 66, 278, 3, 2, 2, 2, 68, 281, 3, 2, 2, 2, 70, 285, 3, 2, 2,
	2, 72, 287, 3, 2, 2, 2, 74, 289, 3, 2, 2, 2, 76, 81, 5, 4, 3, 2, 77, 78,
	7, 3, 2, 2, 78, 80, 5, 4, 3, 2, 79, 77, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2,
	81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2, 83, 81, 3,
	2, 2, 2, 84, 85, 7, 48, 2, 2, 85, 87, 3, 2, 2, 2, 86, 76, 3, 2, 2, 2, 87,
	90, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 3, 3, 2, 2,
	2, 90, 88, 3, 2, 2, 2, 91, 93, 5, 8, 5, 2, 92, 91, 3, 2, 2, 2, 92, 93,
	3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 97, 5, 10, 6, 2, 95, 97, 5, 6, 4, 2,
	96, 94, 3, 2, 2, 2, 96, 95, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 99, 3,
	2, 2, 2, 98, 100, 5, 74, 38, 2, 99, 98, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2,
	100, 5, 3, 2, 2, 2, 101, 103, 5, 72, 37, 2, 102, 101, 3, 2, 2, 2, 102,
	103, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 106, 5, 70, 36, 2, 105, 107,
	5, 48, 25, 2, 106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 7, 3, 2,
	2, 2, 108, 110, 5, 50, 26, 2, 109, 111, 7, 4, 2, 2, 110, 109, 3, 2, 2,
	2, 110, 111, 3, 2, 2, 2, 111, 9, 3, 2, 2, 2, 112, 129, 5, 42, 22, 2, 113,
	129, 5, 40, 21, 2, 114, 129, 5, 30, 16, 2, 115, 129, 5, 38, 20, 2, 116,
	129, 5, 28, 15, 2, 117, 129, 5, 24, 13, 2, 118, 129, 5, 22, 12, 2, 119,
	129, 5, 18, 10, 2, 120, 129, 5, 26, 14, 2, 121, 129, 5, 20, 11, 2, 122,
	129, 5, 44, 23, 2, 123, 129, 5, 46, 24, 2, 124, 129, 5, 12, 7, 2, 125,
	129, 5, 14, 8, 2, 126, 129, 5, 16, 9, 2, 127, 129, 7, 5, 2, 2, 128, 112,
	3, 2, 2, 2, 128, 113, 3, 2, 2, 2, 128, 114, 3, 2, 2, 2, 128, 115, 3, 2,
	2, 2, 128, 116, 3, 2, 2, 2, 128, 117, 3, 2, 2, 2, 128, 118, 3, 2, 2, 2,
	128, 119, 3, 2, 2, 2, 128, 120, 3, 2, 2, 2, 128, 121, 3, 2, 2, 2, 128,
	122, 3, 2, 2, 2, 128, 123, 3, 2, 2, 2, 128, 124, 3, 2, 2, 2, 128, 125,
	3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129, 11, 3, 2,
	2, 2, 130, 132, 5, 66, 34, 2, 131, 130, 3, 2, 2, 2, 131, 132, 3, 2, 2,
	2, 132, 133, 3, 2, 2, 2, 133, 134, 7, 35, 2, 2, 134, 135, 5, 52, 27, 2,
	135, 13, 3, 2, 2, 2, 136, 138, 5, 66, 34, 2, 137, 136, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 7, 36, 2, 2, 140, 141,
	5, 52, 27, 2, 141, 15, 3, 2, 2, 2, 142, 144, 5, 66, 34, 2, 143, 142, 3,
	2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 146, 7, 37, 2,
	2, 146, 147, 5, 52, 27, 2, 147, 17, 3, 2, 2, 2, 148, 150, 7, 21, 2, 2,
	149, 151, 5, 52, 27, 2, 150, 149, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151,
	19, 3, 2, 2, 2, 152, 154, 7, 20, 2, 2, 153, 155, 5, 52, 27, 2, 154, 153,
	3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 21, 3, 2, 2, 2, 156, 157, 7, 29,
	2, 2, 157, 158, 5, 48, 25, 2, 158, 23, 3, 2, 2, 2, 159, 160, 7, 30, 2,
	2, 160, 161, 5, 48, 25, 2, 161, 25, 3, 2, 2, 2, 162, 163, 7, 31, 2, 2,
	163, 164, 5, 48, 25, 2, 164, 27, 3, 2, 2, 2, 165, 166, 5, 66, 34, 2, 166,
	167, 7, 28, 2, 2, 167, 168, 5, 52, 27, 2, 168, 29, 3, 2, 2, 2, 169, 170,
	7, 27, 2, 2, 170, 171, 5, 32, 17, 2, 171, 31, 3, 2, 2, 2, 172, 178, 5,
	36, 19, 2, 173, 174, 5, 34, 18, 2, 174, 175, 5, 36, 19, 2, 175, 177, 3,
	2, 2, 2, 176, 173, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2, 178, 176, 3, 2, 2,
	2, 178, 179, 3, 2, 2, 2, 179, 186, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 181,
	182, 7, 6, 2, 2, 182, 183, 5, 32, 17, 2, 183, 184, 7, 7, 2, 2, 184, 186,
	3, 2, 2, 2, 185, 172, 3, 2, 2, 2, 185, 181, 3, 2, 2, 2, 186, 33, 3, 2,
	2, 2, 187, 188, 9, 2, 2, 2, 188, 35, 3, 2, 2, 2, 189, 194, 5, 66, 34, 2,
	190, 194, 5, 68, 35, 2, 191, 192, 7, 33, 2, 2, 192, 194, 5, 36, 19, 2,
	193, 189, 3, 2, 2, 2, 193, 190, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194,
	37, 3, 2, 2, 2, 195, 196, 7, 26, 2, 2, 196, 39, 3, 2, 2, 2, 197, 198, 7,
	24, 2, 2, 198, 41, 3, 2, 2, 2, 199, 200, 7, 25, 2, 2, 200, 201, 5, 52,
	27, 2, 201, 43, 3, 2, 2, 2, 202, 203, 7, 23, 2, 2, 203, 204, 5, 64, 33,
	2, 204, 45, 3, 2, 2, 2, 205, 206, 7, 22, 2, 2, 206, 207, 5, 66, 34, 2,
	207, 47, 3, 2, 2, 2, 208, 213, 5, 52, 27, 2, 209, 210, 7, 10, 2, 2, 210,
	212, 5, 52, 27, 2, 211, 209, 3, 2, 2, 2, 212, 215, 3, 2, 2, 2, 213, 211,
	3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 49, 3, 2, 2, 2, 215, 213, 3, 2,
	2, 2, 216, 217, 5, 66, 34, 2, 217, 51, 3, 2, 2, 2, 218, 223, 5, 54, 28,
	2, 219, 220, 7, 44, 2, 2, 220, 222, 5, 54, 28, 2, 221, 219, 3, 2, 2, 2,
	222, 225, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224,
	53, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 226, 231, 5, 56, 29, 2, 227, 228,
	9, 3, 2, 2, 228, 230, 5, 56, 29, 2, 229, 227, 3, 2, 2, 2, 230, 233, 3,
	2, 2, 2, 231, 229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 55, 3, 2, 2,
	2, 233, 231, 3, 2, 2, 2, 234, 266, 5, 68, 35, 2, 235, 266, 5, 60, 31, 2,
	236, 266, 5, 62, 32, 2, 237, 266, 5, 66, 34, 2, 238, 266, 5, 64, 33, 2,
	239, 240, 7, 6, 2, 2, 240, 241, 5, 52, 27, 2, 241, 242, 7, 7, 2, 2, 242,
	266, 3, 2, 2, 2, 243, 246, 5, 68, 35, 2, 244, 246, 5, 66, 34, 2, 245, 243,
	3, 2, 2, 2, 245, 244, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 247, 3, 2,
	2, 2, 247, 248, 7, 15, 2, 2, 248, 249, 5, 52, 27, 2, 249, 250, 7, 16, 2,
	2, 250, 266, 3, 2, 2, 2, 251, 252, 5, 58, 30, 2, 252, 253, 5, 52, 27, 2,
	253, 266, 3, 2, 2, 2, 254, 255, 7, 33, 2, 2, 255, 266, 5, 52, 27, 2, 256,
	257, 7, 34, 2, 2, 257, 266, 5, 52, 27, 2, 258, 259, 7, 38, 2, 2, 259, 266,
	5, 52, 27, 2, 260, 261, 5, 62, 32, 2, 261, 262, 7, 4, 2, 2, 262, 263, 3,
	2, 2, 2, 263, 264, 5, 52, 27, 2, 264, 266, 3, 2, 2, 2, 265, 234, 3, 2,
	2, 2, 265, 235, 3, 2, 2, 2, 265, 236, 3, 2, 2, 2, 265, 237, 3, 2, 2, 2,
	265, 238, 3, 2, 2, 2, 265, 239, 3, 2, 2, 2, 265, 245, 3, 2, 2, 2, 265,
	251, 3, 2, 2, 2, 265, 254, 3, 2, 2, 2, 265, 256, 3, 2, 2, 2, 265, 258,
	3, 2, 2, 2, 265, 260, 3, 2, 2, 2, 266, 57, 3, 2, 2, 2, 267, 269, 9, 4,
	2, 2, 268, 267, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2,
	270, 271, 7, 32, 2, 2, 271, 59, 3, 2, 2, 2, 272, 273, 7, 43, 2, 2, 273,
	61, 3, 2, 2, 2, 274, 275, 7, 40, 2, 2, 275, 63, 3, 2, 2, 2, 276, 277, 7,
	47, 2, 2, 277, 65, 3, 2, 2, 2, 278, 279, 7, 45, 2, 2, 279, 67, 3, 2, 2,
	2, 280, 282, 7, 44, 2, 2, 281, 280, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282,
	283, 3, 2, 2, 2, 283, 284, 7, 46, 2, 2, 284, 69, 3, 2, 2, 2, 285, 286,
	7, 41, 2, 2, 286, 71, 3, 2, 2, 2, 287, 288, 7, 42, 2, 2, 288, 73, 3, 2,
	2, 2, 289, 290, 7, 39, 2, 2, 290, 75, 3, 2, 2, 2, 26, 81, 88, 92, 96, 99,
	102, 106, 110, 128, 131, 137, 143, 150, 154, 178, 185, 193, 213, 223, 231,
	245, 265, 268, 281,
}
var literalNames = []string{
	"", "'!'", "':'", "'.'", "'('", "')'", "'eq'", "'ne'", "','", "'*'", "'/'",
	"'mod'", "'and'", "'['", "']'", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'$'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "BYTE", "WORD",
	"DWORD", "DSEG", "CSEG", "INCLUDE", "TITLE", "END", "ORG", "ENDIF", "IF",
	"EQU", "DW", "DB", "DD", "PTR", "NOT", "OFFSET", "RW", "RB", "RS", "LENGTH",
	"COMMENT", "REGISTER", "OPCODE", "REP", "DOLLAR", "SIGN", "NAME", "NUMBER",
	"STRING", "EOL", "WS",
}

var ruleNames = []string{
	"prog", "line", "instruction", "lbl", "assemblerdirective", "rw", "rb",
	"rs", "cseg", "dseg", "dw", "db", "dd", "equ", "if_", "assemblerexpression",
	"assemblerlogical", "assemblerterm", "endif_", "end", "org", "title", "include_",
	"expressionlist", "label", "expression", "multiplyingExpression", "argument",
	"ptr", "dollar", "register_", "string_", "name", "number", "opcode", "rep",
	"comment",
}

type asm8086Parser struct {
	*antlr.BaseParser
}

// Newasm8086Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *asm8086Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newasm8086Parser(input antlr.TokenStream) *asm8086Parser {
	this := new(asm8086Parser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "asm8086.g4"

	return this
}

// asm8086Parser tokens.
const (
	asm8086ParserEOF      = antlr.TokenEOF
	asm8086ParserT__0     = 1
	asm8086ParserT__1     = 2
	asm8086ParserT__2     = 3
	asm8086ParserT__3     = 4
	asm8086ParserT__4     = 5
	asm8086ParserT__5     = 6
	asm8086ParserT__6     = 7
	asm8086ParserT__7     = 8
	asm8086ParserT__8     = 9
	asm8086ParserT__9     = 10
	asm8086ParserT__10    = 11
	asm8086ParserT__11    = 12
	asm8086ParserT__12    = 13
	asm8086ParserT__13    = 14
	asm8086ParserBYTE     = 15
	asm8086ParserWORD     = 16
	asm8086ParserDWORD    = 17
	asm8086ParserDSEG     = 18
	asm8086ParserCSEG     = 19
	asm8086ParserINCLUDE  = 20
	asm8086ParserTITLE    = 21
	asm8086ParserEND      = 22
	asm8086ParserORG      = 23
	asm8086ParserENDIF    = 24
	asm8086ParserIF       = 25
	asm8086ParserEQU      = 26
	asm8086ParserDW       = 27
	asm8086ParserDB       = 28
	asm8086ParserDD       = 29
	asm8086ParserPTR      = 30
	asm8086ParserNOT      = 31
	asm8086ParserOFFSET   = 32
	asm8086ParserRW       = 33
	asm8086ParserRB       = 34
	asm8086ParserRS       = 35
	asm8086ParserLENGTH   = 36
	asm8086ParserCOMMENT  = 37
	asm8086ParserREGISTER = 38
	asm8086ParserOPCODE   = 39
	asm8086ParserREP      = 40
	asm8086ParserDOLLAR   = 41
	asm8086ParserSIGN     = 42
	asm8086ParserNAME     = 43
	asm8086ParserNUMBER   = 44
	asm8086ParserSTRING   = 45
	asm8086ParserEOL      = 46
	asm8086ParserWS       = 47
)

// asm8086Parser rules.
const (
	asm8086ParserRULE_prog                  = 0
	asm8086ParserRULE_line                  = 1
	asm8086ParserRULE_instruction           = 2
	asm8086ParserRULE_lbl                   = 3
	asm8086ParserRULE_assemblerdirective    = 4
	asm8086ParserRULE_rw                    = 5
	asm8086ParserRULE_rb                    = 6
	asm8086ParserRULE_rs                    = 7
	asm8086ParserRULE_cseg                  = 8
	asm8086ParserRULE_dseg                  = 9
	asm8086ParserRULE_dw                    = 10
	asm8086ParserRULE_db                    = 11
	asm8086ParserRULE_dd                    = 12
	asm8086ParserRULE_equ                   = 13
	asm8086ParserRULE_if_                   = 14
	asm8086ParserRULE_assemblerexpression   = 15
	asm8086ParserRULE_assemblerlogical      = 16
	asm8086ParserRULE_assemblerterm         = 17
	asm8086ParserRULE_endif_                = 18
	asm8086ParserRULE_end                   = 19
	asm8086ParserRULE_org                   = 20
	asm8086ParserRULE_title                 = 21
	asm8086ParserRULE_include_              = 22
	asm8086ParserRULE_expressionlist        = 23
	asm8086ParserRULE_label                 = 24
	asm8086ParserRULE_expression            = 25
	asm8086ParserRULE_multiplyingExpression = 26
	asm8086ParserRULE_argument              = 27
	asm8086ParserRULE_ptr                   = 28
	asm8086ParserRULE_dollar                = 29
	asm8086ParserRULE_register_             = 30
	asm8086ParserRULE_string_               = 31
	asm8086ParserRULE_name                  = 32
	asm8086ParserRULE_number                = 33
	asm8086ParserRULE_opcode                = 34
	asm8086ParserRULE_rep                   = 35
	asm8086ParserRULE_comment               = 36
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *ProgContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ProgContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(asm8086ParserEOL)
}

func (s *ProgContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(asm8086ParserEOL, i)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *asm8086Parser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, asm8086ParserRULE_prog)
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
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm8086ParserT__0)|(1<<asm8086ParserT__2)|(1<<asm8086ParserDSEG)|(1<<asm8086ParserCSEG)|(1<<asm8086ParserINCLUDE)|(1<<asm8086ParserTITLE)|(1<<asm8086ParserEND)|(1<<asm8086ParserORG)|(1<<asm8086ParserENDIF)|(1<<asm8086ParserIF)|(1<<asm8086ParserDW)|(1<<asm8086ParserDB)|(1<<asm8086ParserDD))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(asm8086ParserRW-33))|(1<<(asm8086ParserRB-33))|(1<<(asm8086ParserRS-33))|(1<<(asm8086ParserCOMMENT-33))|(1<<(asm8086ParserOPCODE-33))|(1<<(asm8086ParserREP-33))|(1<<(asm8086ParserNAME-33))|(1<<(asm8086ParserEOL-33)))) != 0) {
		{
			p.SetState(74)
			p.Line()
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == asm8086ParserT__0 {
			{
				p.SetState(75)
				p.Match(asm8086ParserT__0)
			}
			{
				p.SetState(76)
				p.Line()
			}

			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(82)
			p.Match(asm8086ParserEOL)
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Lbl() ILblContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILblContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILblContext)
}

func (s *LineContext) Assemblerdirective() IAssemblerdirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssemblerdirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssemblerdirectiveContext)
}

func (s *LineContext) Instruction() IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *LineContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *asm8086Parser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, asm8086ParserRULE_line)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(89)
			p.Lbl()
		}

	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case asm8086ParserT__2, asm8086ParserDSEG, asm8086ParserCSEG, asm8086ParserINCLUDE, asm8086ParserTITLE, asm8086ParserEND, asm8086ParserORG, asm8086ParserENDIF, asm8086ParserIF, asm8086ParserDW, asm8086ParserDB, asm8086ParserDD, asm8086ParserRW, asm8086ParserRB, asm8086ParserRS, asm8086ParserNAME:
		{
			p.SetState(92)
			p.Assemblerdirective()
		}

	case asm8086ParserOPCODE, asm8086ParserREP:
		{
			p.SetState(93)
			p.Instruction()
		}

	case asm8086ParserT__0, asm8086ParserCOMMENT, asm8086ParserEOL:

	default:
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm8086ParserCOMMENT {
		{
			p.SetState(96)
			p.Comment()
		}

	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Opcode() IOpcodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpcodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpcodeContext)
}

func (s *InstructionContext) Rep() IRepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRepContext)
}

func (s *InstructionContext) Expressionlist() IExpressionlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionlistContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *asm8086Parser) Instruction() (localctx IInstructionContext) {
	this := p
	_ = this

	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, asm8086ParserRULE_instruction)
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm8086ParserREP {
		{
			p.SetState(99)
			p.Rep()
		}

	}
	{
		p.SetState(102)
		p.Opcode()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm8086ParserT__3)|(1<<asm8086ParserT__12)|(1<<asm8086ParserBYTE)|(1<<asm8086ParserWORD)|(1<<asm8086ParserDWORD)|(1<<asm8086ParserPTR)|(1<<asm8086ParserNOT))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(asm8086ParserOFFSET-32))|(1<<(asm8086ParserLENGTH-32))|(1<<(asm8086ParserREGISTER-32))|(1<<(asm8086ParserDOLLAR-32))|(1<<(asm8086ParserSIGN-32))|(1<<(asm8086ParserNAME-32))|(1<<(asm8086ParserNUMBER-32))|(1<<(asm8086ParserSTRING-32)))) != 0) {
		{
			p.SetState(103)
			p.Expressionlist()
		}

	}

	return localctx
}

// ILblContext is an interface to support dynamic dispatch.
type ILblContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLblContext differentiates from other interfaces.
	IsLblContext()
}

type LblContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLblContext() *LblContext {
	var p = new(LblContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_lbl
	return p
}

func (*LblContext) IsLblContext() {}

func NewLblContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LblContext {
	var p = new(LblContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_lbl

	return p
}

func (s *LblContext) GetParser() antlr.Parser { return s.parser }

func (s *LblContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LblContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LblContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LblContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterLbl(s)
	}
}

func (s *LblContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitLbl(s)
	}
}

func (p *asm8086Parser) Lbl() (localctx ILblContext) {
	this := p
	_ = this

	localctx = NewLblContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, asm8086ParserRULE_lbl)
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
		p.SetState(106)
		p.Label()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm8086ParserT__1 {
		{
			p.SetState(107)
			p.Match(asm8086ParserT__1)
		}

	}

	return localctx
}

// IAssemblerdirectiveContext is an interface to support dynamic dispatch.
type IAssemblerdirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssemblerdirectiveContext differentiates from other interfaces.
	IsAssemblerdirectiveContext()
}

type AssemblerdirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssemblerdirectiveContext() *AssemblerdirectiveContext {
	var p = new(AssemblerdirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_assemblerdirective
	return p
}

func (*AssemblerdirectiveContext) IsAssemblerdirectiveContext() {}

func NewAssemblerdirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssemblerdirectiveContext {
	var p = new(AssemblerdirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_assemblerdirective

	return p
}

func (s *AssemblerdirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *AssemblerdirectiveContext) Org() IOrgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrgContext)
}

func (s *AssemblerdirectiveContext) End() IEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndContext)
}

func (s *AssemblerdirectiveContext) If_() IIf_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_Context)
}

func (s *AssemblerdirectiveContext) Endif_() IEndif_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndif_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndif_Context)
}

func (s *AssemblerdirectiveContext) Equ() IEquContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEquContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEquContext)
}

func (s *AssemblerdirectiveContext) Db() IDbContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDbContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDbContext)
}

func (s *AssemblerdirectiveContext) Dw() IDwContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDwContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDwContext)
}

func (s *AssemblerdirectiveContext) Cseg() ICsegContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICsegContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICsegContext)
}

func (s *AssemblerdirectiveContext) Dd() IDdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDdContext)
}

func (s *AssemblerdirectiveContext) Dseg() IDsegContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDsegContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDsegContext)
}

func (s *AssemblerdirectiveContext) Title() ITitleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITitleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITitleContext)
}

func (s *AssemblerdirectiveContext) Include_() IInclude_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInclude_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInclude_Context)
}

func (s *AssemblerdirectiveContext) Rw() IRwContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRwContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRwContext)
}

func (s *AssemblerdirectiveContext) Rb() IRbContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRbContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRbContext)
}

func (s *AssemblerdirectiveContext) Rs() IRsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRsContext)
}

func (s *AssemblerdirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssemblerdirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssemblerdirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterAssemblerdirective(s)
	}
}

func (s *AssemblerdirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitAssemblerdirective(s)
	}
}

func (p *asm8086Parser) Assemblerdirective() (localctx IAssemblerdirectiveContext) {
	this := p
	_ = this

	localctx = NewAssemblerdirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, asm8086ParserRULE_assemblerdirective)

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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Org()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.End()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.If_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Endif_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.Equ()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(115)
			p.Db()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(116)
			p.Dw()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(117)
			p.Cseg()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(118)
			p.Dd()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(119)
			p.Dseg()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(120)
			p.Title()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(121)
			p.Include_()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(122)
			p.Rw()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(123)
			p.Rb()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(124)
			p.Rs()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(125)
			p.Match(asm8086ParserT__2)
		}

	}

	return localctx
}

// IRwContext is an interface to support dynamic dispatch.
type IRwContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRwContext differentiates from other interfaces.
	IsRwContext()
}

type RwContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRwContext() *RwContext {
	var p = new(RwContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_rw
	return p
}

func (*RwContext) IsRwContext() {}

func NewRwContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RwContext {
	var p = new(RwContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_rw

	return p
}

func (s *RwContext) GetParser() antlr.Parser { return s.parser }

func (s *RwContext) RW() antlr.TerminalNode {
	return s.GetToken(asm8086ParserRW, 0)
}

func (s *RwContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RwContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RwContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RwContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RwContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterRw(s)
	}
}

func (s *RwContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitRw(s)
	}
}

func (p *asm8086Parser) Rw() (localctx IRwContext) {
	this := p
	_ = this

	localctx = NewRwContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, asm8086ParserRULE_rw)
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
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm8086ParserNAME {
		{
			p.SetState(128)
			p.Name()
		}

	}
	{
		p.SetState(131)
		p.Match(asm8086ParserRW)
	}
	{
		p.SetState(132)
		p.Expression()
	}

	return localctx
}

// IRbContext is an interface to support dynamic dispatch.
type IRbContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRbContext differentiates from other interfaces.
	IsRbContext()
}

type RbContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRbContext() *RbContext {
	var p = new(RbContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_rb
	return p
}

func (*RbContext) IsRbContext() {}

func NewRbContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RbContext {
	var p = new(RbContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_rb

	return p
}

func (s *RbContext) GetParser() antlr.Parser { return s.parser }

func (s *RbContext) RB() antlr.TerminalNode {
	return s.GetToken(asm8086ParserRB, 0)
}

func (s *RbContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RbContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RbContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RbContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RbContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterRb(s)
	}
}

func (s *RbContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitRb(s)
	}
}

func (p *asm8086Parser) Rb() (localctx IRbContext) {
	this := p
	_ = this

	localctx = NewRbContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, asm8086ParserRULE_rb)
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
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm8086ParserNAME {
		{
			p.SetState(134)
			p.Name()
		}

	}
	{
		p.SetState(137)
		p.Match(asm8086ParserRB)
	}
	{
		p.SetState(138)
		p.Expression()
	}

	return localctx
}

// IRsContext is an interface to support dynamic dispatch.
type IRsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRsContext differentiates from other interfaces.
	IsRsContext()
}

type RsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRsContext() *RsContext {
	var p = new(RsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_rs
	return p
}

func (*RsContext) IsRsContext() {}

func NewRsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RsContext {
	var p = new(RsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_rs

	return p
}

func (s *RsContext) GetParser() antlr.Parser { return s.parser }

func (s *RsContext) RS() antlr.TerminalNode {
	return s.GetToken(asm8086ParserRS, 0)
}

func (s *RsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RsContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterRs(s)
	}
}

func (s *RsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitRs(s)
	}
}

func (p *asm8086Parser) Rs() (localctx IRsContext) {
	this := p
	_ = this

	localctx = NewRsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, asm8086ParserRULE_rs)
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
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm8086ParserNAME {
		{
			p.SetState(140)
			p.Name()
		}

	}
	{
		p.SetState(143)
		p.Match(asm8086ParserRS)
	}
	{
		p.SetState(144)
		p.Expression()
	}

	return localctx
}

// ICsegContext is an interface to support dynamic dispatch.
type ICsegContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCsegContext differentiates from other interfaces.
	IsCsegContext()
}

type CsegContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCsegContext() *CsegContext {
	var p = new(CsegContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_cseg
	return p
}

func (*CsegContext) IsCsegContext() {}

func NewCsegContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CsegContext {
	var p = new(CsegContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_cseg

	return p
}

func (s *CsegContext) GetParser() antlr.Parser { return s.parser }

func (s *CsegContext) CSEG() antlr.TerminalNode {
	return s.GetToken(asm8086ParserCSEG, 0)
}

func (s *CsegContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CsegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CsegContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CsegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterCseg(s)
	}
}

func (s *CsegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitCseg(s)
	}
}

func (p *asm8086Parser) Cseg() (localctx ICsegContext) {
	this := p
	_ = this

	localctx = NewCsegContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, asm8086ParserRULE_cseg)
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
		p.SetState(146)
		p.Match(asm8086ParserCSEG)
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm8086ParserT__3)|(1<<asm8086ParserT__12)|(1<<asm8086ParserBYTE)|(1<<asm8086ParserWORD)|(1<<asm8086ParserDWORD)|(1<<asm8086ParserPTR)|(1<<asm8086ParserNOT))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(asm8086ParserOFFSET-32))|(1<<(asm8086ParserLENGTH-32))|(1<<(asm8086ParserREGISTER-32))|(1<<(asm8086ParserDOLLAR-32))|(1<<(asm8086ParserSIGN-32))|(1<<(asm8086ParserNAME-32))|(1<<(asm8086ParserNUMBER-32))|(1<<(asm8086ParserSTRING-32)))) != 0) {
		{
			p.SetState(147)
			p.Expression()
		}

	}

	return localctx
}

// IDsegContext is an interface to support dynamic dispatch.
type IDsegContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDsegContext differentiates from other interfaces.
	IsDsegContext()
}

type DsegContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDsegContext() *DsegContext {
	var p = new(DsegContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_dseg
	return p
}

func (*DsegContext) IsDsegContext() {}

func NewDsegContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DsegContext {
	var p = new(DsegContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_dseg

	return p
}

func (s *DsegContext) GetParser() antlr.Parser { return s.parser }

func (s *DsegContext) DSEG() antlr.TerminalNode {
	return s.GetToken(asm8086ParserDSEG, 0)
}

func (s *DsegContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DsegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DsegContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DsegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterDseg(s)
	}
}

func (s *DsegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitDseg(s)
	}
}

func (p *asm8086Parser) Dseg() (localctx IDsegContext) {
	this := p
	_ = this

	localctx = NewDsegContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, asm8086ParserRULE_dseg)
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
		p.SetState(150)
		p.Match(asm8086ParserDSEG)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm8086ParserT__3)|(1<<asm8086ParserT__12)|(1<<asm8086ParserBYTE)|(1<<asm8086ParserWORD)|(1<<asm8086ParserDWORD)|(1<<asm8086ParserPTR)|(1<<asm8086ParserNOT))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(asm8086ParserOFFSET-32))|(1<<(asm8086ParserLENGTH-32))|(1<<(asm8086ParserREGISTER-32))|(1<<(asm8086ParserDOLLAR-32))|(1<<(asm8086ParserSIGN-32))|(1<<(asm8086ParserNAME-32))|(1<<(asm8086ParserNUMBER-32))|(1<<(asm8086ParserSTRING-32)))) != 0) {
		{
			p.SetState(151)
			p.Expression()
		}

	}

	return localctx
}

// IDwContext is an interface to support dynamic dispatch.
type IDwContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDwContext differentiates from other interfaces.
	IsDwContext()
}

type DwContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDwContext() *DwContext {
	var p = new(DwContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_dw
	return p
}

func (*DwContext) IsDwContext() {}

func NewDwContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DwContext {
	var p = new(DwContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_dw

	return p
}

func (s *DwContext) GetParser() antlr.Parser { return s.parser }

func (s *DwContext) DW() antlr.TerminalNode {
	return s.GetToken(asm8086ParserDW, 0)
}

func (s *DwContext) Expressionlist() IExpressionlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionlistContext)
}

func (s *DwContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DwContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DwContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterDw(s)
	}
}

func (s *DwContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitDw(s)
	}
}

func (p *asm8086Parser) Dw() (localctx IDwContext) {
	this := p
	_ = this

	localctx = NewDwContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, asm8086ParserRULE_dw)

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
		p.SetState(154)
		p.Match(asm8086ParserDW)
	}
	{
		p.SetState(155)
		p.Expressionlist()
	}

	return localctx
}

// IDbContext is an interface to support dynamic dispatch.
type IDbContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDbContext differentiates from other interfaces.
	IsDbContext()
}

type DbContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDbContext() *DbContext {
	var p = new(DbContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_db
	return p
}

func (*DbContext) IsDbContext() {}

func NewDbContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DbContext {
	var p = new(DbContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_db

	return p
}

func (s *DbContext) GetParser() antlr.Parser { return s.parser }

func (s *DbContext) DB() antlr.TerminalNode {
	return s.GetToken(asm8086ParserDB, 0)
}

func (s *DbContext) Expressionlist() IExpressionlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionlistContext)
}

func (s *DbContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DbContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DbContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterDb(s)
	}
}

func (s *DbContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitDb(s)
	}
}

func (p *asm8086Parser) Db() (localctx IDbContext) {
	this := p
	_ = this

	localctx = NewDbContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, asm8086ParserRULE_db)

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
		p.SetState(157)
		p.Match(asm8086ParserDB)
	}
	{
		p.SetState(158)
		p.Expressionlist()
	}

	return localctx
}

// IDdContext is an interface to support dynamic dispatch.
type IDdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDdContext differentiates from other interfaces.
	IsDdContext()
}

type DdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDdContext() *DdContext {
	var p = new(DdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_dd
	return p
}

func (*DdContext) IsDdContext() {}

func NewDdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DdContext {
	var p = new(DdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_dd

	return p
}

func (s *DdContext) GetParser() antlr.Parser { return s.parser }

func (s *DdContext) DD() antlr.TerminalNode {
	return s.GetToken(asm8086ParserDD, 0)
}

func (s *DdContext) Expressionlist() IExpressionlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionlistContext)
}

func (s *DdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterDd(s)
	}
}

func (s *DdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitDd(s)
	}
}

func (p *asm8086Parser) Dd() (localctx IDdContext) {
	this := p
	_ = this

	localctx = NewDdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, asm8086ParserRULE_dd)

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
		p.SetState(160)
		p.Match(asm8086ParserDD)
	}
	{
		p.SetState(161)
		p.Expressionlist()
	}

	return localctx
}

// IEquContext is an interface to support dynamic dispatch.
type IEquContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEquContext differentiates from other interfaces.
	IsEquContext()
}

type EquContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquContext() *EquContext {
	var p = new(EquContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_equ
	return p
}

func (*EquContext) IsEquContext() {}

func NewEquContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquContext {
	var p = new(EquContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_equ

	return p
}

func (s *EquContext) GetParser() antlr.Parser { return s.parser }

func (s *EquContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EquContext) EQU() antlr.TerminalNode {
	return s.GetToken(asm8086ParserEQU, 0)
}

func (s *EquContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EquContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EquContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterEqu(s)
	}
}

func (s *EquContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitEqu(s)
	}
}

func (p *asm8086Parser) Equ() (localctx IEquContext) {
	this := p
	_ = this

	localctx = NewEquContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, asm8086ParserRULE_equ)

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
		p.Name()
	}
	{
		p.SetState(164)
		p.Match(asm8086ParserEQU)
	}
	{
		p.SetState(165)
		p.Expression()
	}

	return localctx
}

// IIf_Context is an interface to support dynamic dispatch.
type IIf_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_Context differentiates from other interfaces.
	IsIf_Context()
}

type If_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_Context() *If_Context {
	var p = new(If_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_if_
	return p
}

func (*If_Context) IsIf_Context() {}

func NewIf_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_Context {
	var p = new(If_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_if_

	return p
}

func (s *If_Context) GetParser() antlr.Parser { return s.parser }

func (s *If_Context) IF() antlr.TerminalNode {
	return s.GetToken(asm8086ParserIF, 0)
}

func (s *If_Context) Assemblerexpression() IAssemblerexpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssemblerexpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssemblerexpressionContext)
}

func (s *If_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterIf_(s)
	}
}

func (s *If_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitIf_(s)
	}
}

func (p *asm8086Parser) If_() (localctx IIf_Context) {
	this := p
	_ = this

	localctx = NewIf_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, asm8086ParserRULE_if_)

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
		p.SetState(167)
		p.Match(asm8086ParserIF)
	}
	{
		p.SetState(168)
		p.Assemblerexpression()
	}

	return localctx
}

// IAssemblerexpressionContext is an interface to support dynamic dispatch.
type IAssemblerexpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssemblerexpressionContext differentiates from other interfaces.
	IsAssemblerexpressionContext()
}

type AssemblerexpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssemblerexpressionContext() *AssemblerexpressionContext {
	var p = new(AssemblerexpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_assemblerexpression
	return p
}

func (*AssemblerexpressionContext) IsAssemblerexpressionContext() {}

func NewAssemblerexpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssemblerexpressionContext {
	var p = new(AssemblerexpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_assemblerexpression

	return p
}

func (s *AssemblerexpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssemblerexpressionContext) AllAssemblerterm() []IAssemblertermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssemblertermContext)(nil)).Elem())
	var tst = make([]IAssemblertermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssemblertermContext)
		}
	}

	return tst
}

func (s *AssemblerexpressionContext) Assemblerterm(i int) IAssemblertermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssemblertermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssemblertermContext)
}

func (s *AssemblerexpressionContext) AllAssemblerlogical() []IAssemblerlogicalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssemblerlogicalContext)(nil)).Elem())
	var tst = make([]IAssemblerlogicalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssemblerlogicalContext)
		}
	}

	return tst
}

func (s *AssemblerexpressionContext) Assemblerlogical(i int) IAssemblerlogicalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssemblerlogicalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssemblerlogicalContext)
}

func (s *AssemblerexpressionContext) Assemblerexpression() IAssemblerexpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssemblerexpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssemblerexpressionContext)
}

func (s *AssemblerexpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssemblerexpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssemblerexpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterAssemblerexpression(s)
	}
}

func (s *AssemblerexpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitAssemblerexpression(s)
	}
}

func (p *asm8086Parser) Assemblerexpression() (localctx IAssemblerexpressionContext) {
	this := p
	_ = this

	localctx = NewAssemblerexpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, asm8086ParserRULE_assemblerexpression)
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

	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case asm8086ParserNOT, asm8086ParserSIGN, asm8086ParserNAME, asm8086ParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Assemblerterm()
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == asm8086ParserT__5 || _la == asm8086ParserT__6 {
			{
				p.SetState(171)
				p.Assemblerlogical()
			}
			{
				p.SetState(172)
				p.Assemblerterm()
			}

			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case asm8086ParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Match(asm8086ParserT__3)
		}
		{
			p.SetState(180)
			p.Assemblerexpression()
		}
		{
			p.SetState(181)
			p.Match(asm8086ParserT__4)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssemblerlogicalContext is an interface to support dynamic dispatch.
type IAssemblerlogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssemblerlogicalContext differentiates from other interfaces.
	IsAssemblerlogicalContext()
}

type AssemblerlogicalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssemblerlogicalContext() *AssemblerlogicalContext {
	var p = new(AssemblerlogicalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_assemblerlogical
	return p
}

func (*AssemblerlogicalContext) IsAssemblerlogicalContext() {}

func NewAssemblerlogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssemblerlogicalContext {
	var p = new(AssemblerlogicalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_assemblerlogical

	return p
}

func (s *AssemblerlogicalContext) GetParser() antlr.Parser { return s.parser }
func (s *AssemblerlogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssemblerlogicalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssemblerlogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterAssemblerlogical(s)
	}
}

func (s *AssemblerlogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitAssemblerlogical(s)
	}
}

func (p *asm8086Parser) Assemblerlogical() (localctx IAssemblerlogicalContext) {
	this := p
	_ = this

	localctx = NewAssemblerlogicalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, asm8086ParserRULE_assemblerlogical)
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
		p.SetState(185)
		_la = p.GetTokenStream().LA(1)

		if !(_la == asm8086ParserT__5 || _la == asm8086ParserT__6) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssemblertermContext is an interface to support dynamic dispatch.
type IAssemblertermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssemblertermContext differentiates from other interfaces.
	IsAssemblertermContext()
}

type AssemblertermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssemblertermContext() *AssemblertermContext {
	var p = new(AssemblertermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_assemblerterm
	return p
}

func (*AssemblertermContext) IsAssemblertermContext() {}

func NewAssemblertermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssemblertermContext {
	var p = new(AssemblertermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_assemblerterm

	return p
}

func (s *AssemblertermContext) GetParser() antlr.Parser { return s.parser }

func (s *AssemblertermContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AssemblertermContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *AssemblertermContext) NOT() antlr.TerminalNode {
	return s.GetToken(asm8086ParserNOT, 0)
}

func (s *AssemblertermContext) Assemblerterm() IAssemblertermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssemblertermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssemblertermContext)
}

func (s *AssemblertermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssemblertermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssemblertermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterAssemblerterm(s)
	}
}

func (s *AssemblertermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitAssemblerterm(s)
	}
}

func (p *asm8086Parser) Assemblerterm() (localctx IAssemblertermContext) {
	this := p
	_ = this

	localctx = NewAssemblertermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, asm8086ParserRULE_assemblerterm)

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

	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case asm8086ParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Name()
		}

	case asm8086ParserSIGN, asm8086ParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Number()
		}

	case asm8086ParserNOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(189)
			p.Match(asm8086ParserNOT)
		}
		{
			p.SetState(190)
			p.Assemblerterm()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEndif_Context is an interface to support dynamic dispatch.
type IEndif_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndif_Context differentiates from other interfaces.
	IsEndif_Context()
}

type Endif_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndif_Context() *Endif_Context {
	var p = new(Endif_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_endif_
	return p
}

func (*Endif_Context) IsEndif_Context() {}

func NewEndif_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Endif_Context {
	var p = new(Endif_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_endif_

	return p
}

func (s *Endif_Context) GetParser() antlr.Parser { return s.parser }

func (s *Endif_Context) ENDIF() antlr.TerminalNode {
	return s.GetToken(asm8086ParserENDIF, 0)
}

func (s *Endif_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Endif_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Endif_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterEndif_(s)
	}
}

func (s *Endif_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitEndif_(s)
	}
}

func (p *asm8086Parser) Endif_() (localctx IEndif_Context) {
	this := p
	_ = this

	localctx = NewEndif_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, asm8086ParserRULE_endif_)

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
		p.SetState(193)
		p.Match(asm8086ParserENDIF)
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
	p.RuleIndex = asm8086ParserRULE_end
	return p
}

func (*EndContext) IsEndContext() {}

func NewEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndContext {
	var p = new(EndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_end

	return p
}

func (s *EndContext) GetParser() antlr.Parser { return s.parser }

func (s *EndContext) END() antlr.TerminalNode {
	return s.GetToken(asm8086ParserEND, 0)
}

func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterEnd(s)
	}
}

func (s *EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitEnd(s)
	}
}

func (p *asm8086Parser) End() (localctx IEndContext) {
	this := p
	_ = this

	localctx = NewEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, asm8086ParserRULE_end)

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
		p.SetState(195)
		p.Match(asm8086ParserEND)
	}

	return localctx
}

// IOrgContext is an interface to support dynamic dispatch.
type IOrgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrgContext differentiates from other interfaces.
	IsOrgContext()
}

type OrgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrgContext() *OrgContext {
	var p = new(OrgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_org
	return p
}

func (*OrgContext) IsOrgContext() {}

func NewOrgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrgContext {
	var p = new(OrgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_org

	return p
}

func (s *OrgContext) GetParser() antlr.Parser { return s.parser }

func (s *OrgContext) ORG() antlr.TerminalNode {
	return s.GetToken(asm8086ParserORG, 0)
}

func (s *OrgContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterOrg(s)
	}
}

func (s *OrgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitOrg(s)
	}
}

func (p *asm8086Parser) Org() (localctx IOrgContext) {
	this := p
	_ = this

	localctx = NewOrgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, asm8086ParserRULE_org)

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
		p.Match(asm8086ParserORG)
	}
	{
		p.SetState(198)
		p.Expression()
	}

	return localctx
}

// ITitleContext is an interface to support dynamic dispatch.
type ITitleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTitleContext differentiates from other interfaces.
	IsTitleContext()
}

type TitleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTitleContext() *TitleContext {
	var p = new(TitleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_title
	return p
}

func (*TitleContext) IsTitleContext() {}

func NewTitleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TitleContext {
	var p = new(TitleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_title

	return p
}

func (s *TitleContext) GetParser() antlr.Parser { return s.parser }

func (s *TitleContext) TITLE() antlr.TerminalNode {
	return s.GetToken(asm8086ParserTITLE, 0)
}

func (s *TitleContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *TitleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TitleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TitleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterTitle(s)
	}
}

func (s *TitleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitTitle(s)
	}
}

func (p *asm8086Parser) Title() (localctx ITitleContext) {
	this := p
	_ = this

	localctx = NewTitleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, asm8086ParserRULE_title)

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
		p.SetState(200)
		p.Match(asm8086ParserTITLE)
	}
	{
		p.SetState(201)
		p.String_()
	}

	return localctx
}

// IInclude_Context is an interface to support dynamic dispatch.
type IInclude_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInclude_Context differentiates from other interfaces.
	IsInclude_Context()
}

type Include_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInclude_Context() *Include_Context {
	var p = new(Include_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_include_
	return p
}

func (*Include_Context) IsInclude_Context() {}

func NewInclude_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Include_Context {
	var p = new(Include_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_include_

	return p
}

func (s *Include_Context) GetParser() antlr.Parser { return s.parser }

func (s *Include_Context) INCLUDE() antlr.TerminalNode {
	return s.GetToken(asm8086ParserINCLUDE, 0)
}

func (s *Include_Context) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Include_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Include_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Include_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterInclude_(s)
	}
}

func (s *Include_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitInclude_(s)
	}
}

func (p *asm8086Parser) Include_() (localctx IInclude_Context) {
	this := p
	_ = this

	localctx = NewInclude_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, asm8086ParserRULE_include_)

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
		p.SetState(203)
		p.Match(asm8086ParserINCLUDE)
	}
	{
		p.SetState(204)
		p.Name()
	}

	return localctx
}

// IExpressionlistContext is an interface to support dynamic dispatch.
type IExpressionlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionlistContext differentiates from other interfaces.
	IsExpressionlistContext()
}

type ExpressionlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionlistContext() *ExpressionlistContext {
	var p = new(ExpressionlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_expressionlist
	return p
}

func (*ExpressionlistContext) IsExpressionlistContext() {}

func NewExpressionlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionlistContext {
	var p = new(ExpressionlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_expressionlist

	return p
}

func (s *ExpressionlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionlistContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionlistContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterExpressionlist(s)
	}
}

func (s *ExpressionlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitExpressionlist(s)
	}
}

func (p *asm8086Parser) Expressionlist() (localctx IExpressionlistContext) {
	this := p
	_ = this

	localctx = NewExpressionlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, asm8086ParserRULE_expressionlist)
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
		p.SetState(206)
		p.Expression()
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == asm8086ParserT__7 {
		{
			p.SetState(207)
			p.Match(asm8086ParserT__7)
		}
		{
			p.SetState(208)
			p.Expression()
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *asm8086Parser) Label() (localctx ILabelContext) {
	this := p
	_ = this

	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, asm8086ParserRULE_label)

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
		p.SetState(214)
		p.Name()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllMultiplyingExpression() []IMultiplyingExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplyingExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplyingExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) MultiplyingExpression(i int) IMultiplyingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplyingExpressionContext)
}

func (s *ExpressionContext) AllSIGN() []antlr.TerminalNode {
	return s.GetTokens(asm8086ParserSIGN)
}

func (s *ExpressionContext) SIGN(i int) antlr.TerminalNode {
	return s.GetToken(asm8086ParserSIGN, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *asm8086Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, asm8086ParserRULE_expression)

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
		p.SetState(216)
		p.MultiplyingExpression()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(217)
				p.Match(asm8086ParserSIGN)
			}
			{
				p.SetState(218)
				p.MultiplyingExpression()
			}

		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplyingExpressionContext is an interface to support dynamic dispatch.
type IMultiplyingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplyingExpressionContext differentiates from other interfaces.
	IsMultiplyingExpressionContext()
}

type MultiplyingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyingExpressionContext() *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_multiplyingExpression
	return p
}

func (*MultiplyingExpressionContext) IsMultiplyingExpressionContext() {}

func NewMultiplyingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_multiplyingExpression

	return p
}

func (s *MultiplyingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyingExpressionContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *MultiplyingExpressionContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *MultiplyingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterMultiplyingExpression(s)
	}
}

func (s *MultiplyingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitMultiplyingExpression(s)
	}
}

func (p *asm8086Parser) MultiplyingExpression() (localctx IMultiplyingExpressionContext) {
	this := p
	_ = this

	localctx = NewMultiplyingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, asm8086ParserRULE_multiplyingExpression)
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
		p.SetState(224)
		p.Argument()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(225)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm8086ParserT__8)|(1<<asm8086ParserT__9)|(1<<asm8086ParserT__10)|(1<<asm8086ParserT__11))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(226)
				p.Argument()
			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ArgumentContext) Dollar() IDollarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDollarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDollarContext)
}

func (s *ArgumentContext) Register_() IRegister_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegister_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegister_Context)
}

func (s *ArgumentContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ArgumentContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *ArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentContext) Ptr() IPtrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPtrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPtrContext)
}

func (s *ArgumentContext) NOT() antlr.TerminalNode {
	return s.GetToken(asm8086ParserNOT, 0)
}

func (s *ArgumentContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(asm8086ParserOFFSET, 0)
}

func (s *ArgumentContext) LENGTH() antlr.TerminalNode {
	return s.GetToken(asm8086ParserLENGTH, 0)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *asm8086Parser) Argument() (localctx IArgumentContext) {
	this := p
	_ = this

	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, asm8086ParserRULE_argument)

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

	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Dollar()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(234)
			p.Register_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(235)
			p.Name()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(236)
			p.String_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(237)
			p.Match(asm8086ParserT__3)
		}
		{
			p.SetState(238)
			p.Expression()
		}
		{
			p.SetState(239)
			p.Match(asm8086ParserT__4)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		p.SetState(243)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case asm8086ParserSIGN, asm8086ParserNUMBER:
			{
				p.SetState(241)
				p.Number()
			}

		case asm8086ParserNAME:
			{
				p.SetState(242)
				p.Name()
			}

		case asm8086ParserT__12:

		default:
		}
		{
			p.SetState(245)
			p.Match(asm8086ParserT__12)
		}
		{
			p.SetState(246)
			p.Expression()
		}
		{
			p.SetState(247)
			p.Match(asm8086ParserT__13)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(249)
			p.Ptr()
		}
		{
			p.SetState(250)
			p.Expression()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(252)
			p.Match(asm8086ParserNOT)
		}
		{
			p.SetState(253)
			p.Expression()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(254)
			p.Match(asm8086ParserOFFSET)
		}
		{
			p.SetState(255)
			p.Expression()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(256)
			p.Match(asm8086ParserLENGTH)
		}
		{
			p.SetState(257)
			p.Expression()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(258)
			p.Register_()
		}
		{
			p.SetState(259)
			p.Match(asm8086ParserT__1)
		}

		{
			p.SetState(261)
			p.Expression()
		}

	}

	return localctx
}

// IPtrContext is an interface to support dynamic dispatch.
type IPtrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPtrContext differentiates from other interfaces.
	IsPtrContext()
}

type PtrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPtrContext() *PtrContext {
	var p = new(PtrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_ptr
	return p
}

func (*PtrContext) IsPtrContext() {}

func NewPtrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PtrContext {
	var p = new(PtrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_ptr

	return p
}

func (s *PtrContext) GetParser() antlr.Parser { return s.parser }

func (s *PtrContext) PTR() antlr.TerminalNode {
	return s.GetToken(asm8086ParserPTR, 0)
}

func (s *PtrContext) BYTE() antlr.TerminalNode {
	return s.GetToken(asm8086ParserBYTE, 0)
}

func (s *PtrContext) WORD() antlr.TerminalNode {
	return s.GetToken(asm8086ParserWORD, 0)
}

func (s *PtrContext) DWORD() antlr.TerminalNode {
	return s.GetToken(asm8086ParserDWORD, 0)
}

func (s *PtrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PtrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PtrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterPtr(s)
	}
}

func (s *PtrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitPtr(s)
	}
}

func (p *asm8086Parser) Ptr() (localctx IPtrContext) {
	this := p
	_ = this

	localctx = NewPtrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, asm8086ParserRULE_ptr)
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
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm8086ParserBYTE)|(1<<asm8086ParserWORD)|(1<<asm8086ParserDWORD))) != 0 {
		{
			p.SetState(265)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<asm8086ParserBYTE)|(1<<asm8086ParserWORD)|(1<<asm8086ParserDWORD))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(268)
		p.Match(asm8086ParserPTR)
	}

	return localctx
}

// IDollarContext is an interface to support dynamic dispatch.
type IDollarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDollarContext differentiates from other interfaces.
	IsDollarContext()
}

type DollarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDollarContext() *DollarContext {
	var p = new(DollarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_dollar
	return p
}

func (*DollarContext) IsDollarContext() {}

func NewDollarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DollarContext {
	var p = new(DollarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_dollar

	return p
}

func (s *DollarContext) GetParser() antlr.Parser { return s.parser }

func (s *DollarContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(asm8086ParserDOLLAR, 0)
}

func (s *DollarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DollarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DollarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterDollar(s)
	}
}

func (s *DollarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitDollar(s)
	}
}

func (p *asm8086Parser) Dollar() (localctx IDollarContext) {
	this := p
	_ = this

	localctx = NewDollarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, asm8086ParserRULE_dollar)

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
		p.SetState(270)
		p.Match(asm8086ParserDOLLAR)
	}

	return localctx
}

// IRegister_Context is an interface to support dynamic dispatch.
type IRegister_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegister_Context differentiates from other interfaces.
	IsRegister_Context()
}

type Register_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegister_Context() *Register_Context {
	var p = new(Register_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_register_
	return p
}

func (*Register_Context) IsRegister_Context() {}

func NewRegister_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Register_Context {
	var p = new(Register_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_register_

	return p
}

func (s *Register_Context) GetParser() antlr.Parser { return s.parser }

func (s *Register_Context) REGISTER() antlr.TerminalNode {
	return s.GetToken(asm8086ParserREGISTER, 0)
}

func (s *Register_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Register_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Register_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterRegister_(s)
	}
}

func (s *Register_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitRegister_(s)
	}
}

func (p *asm8086Parser) Register_() (localctx IRegister_Context) {
	this := p
	_ = this

	localctx = NewRegister_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, asm8086ParserRULE_register_)

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
		p.SetState(272)
		p.Match(asm8086ParserREGISTER)
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) STRING() antlr.TerminalNode {
	return s.GetToken(asm8086ParserSTRING, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitString_(s)
	}
}

func (p *asm8086Parser) String_() (localctx IString_Context) {
	this := p
	_ = this

	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, asm8086ParserRULE_string_)

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
		p.SetState(274)
		p.Match(asm8086ParserSTRING)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(asm8086ParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitName(s)
	}
}

func (p *asm8086Parser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, asm8086ParserRULE_name)

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
		p.SetState(276)
		p.Match(asm8086ParserNAME)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(asm8086ParserNUMBER, 0)
}

func (s *NumberContext) SIGN() antlr.TerminalNode {
	return s.GetToken(asm8086ParserSIGN, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *asm8086Parser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, asm8086ParserRULE_number)
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
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == asm8086ParserSIGN {
		{
			p.SetState(278)
			p.Match(asm8086ParserSIGN)
		}

	}
	{
		p.SetState(281)
		p.Match(asm8086ParserNUMBER)
	}

	return localctx
}

// IOpcodeContext is an interface to support dynamic dispatch.
type IOpcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpcodeContext differentiates from other interfaces.
	IsOpcodeContext()
}

type OpcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcodeContext() *OpcodeContext {
	var p = new(OpcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_opcode
	return p
}

func (*OpcodeContext) IsOpcodeContext() {}

func NewOpcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcodeContext {
	var p = new(OpcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_opcode

	return p
}

func (s *OpcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcodeContext) OPCODE() antlr.TerminalNode {
	return s.GetToken(asm8086ParserOPCODE, 0)
}

func (s *OpcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterOpcode(s)
	}
}

func (s *OpcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitOpcode(s)
	}
}

func (p *asm8086Parser) Opcode() (localctx IOpcodeContext) {
	this := p
	_ = this

	localctx = NewOpcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, asm8086ParserRULE_opcode)

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
		p.SetState(283)
		p.Match(asm8086ParserOPCODE)
	}

	return localctx
}

// IRepContext is an interface to support dynamic dispatch.
type IRepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRepContext differentiates from other interfaces.
	IsRepContext()
}

type RepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepContext() *RepContext {
	var p = new(RepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_rep
	return p
}

func (*RepContext) IsRepContext() {}

func NewRepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepContext {
	var p = new(RepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_rep

	return p
}

func (s *RepContext) GetParser() antlr.Parser { return s.parser }

func (s *RepContext) REP() antlr.TerminalNode {
	return s.GetToken(asm8086ParserREP, 0)
}

func (s *RepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterRep(s)
	}
}

func (s *RepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitRep(s)
	}
}

func (p *asm8086Parser) Rep() (localctx IRepContext) {
	this := p
	_ = this

	localctx = NewRepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, asm8086ParserRULE_rep)

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
		p.SetState(285)
		p.Match(asm8086ParserREP)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = asm8086ParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = asm8086ParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(asm8086ParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(asm8086Listener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *asm8086Parser) Comment() (localctx ICommentContext) {
	this := p
	_ = this

	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, asm8086ParserRULE_comment)

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
		p.SetState(287)
		p.Match(asm8086ParserCOMMENT)
	}

	return localctx
}
