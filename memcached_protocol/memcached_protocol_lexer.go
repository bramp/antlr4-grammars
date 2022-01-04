// Code generated from memcached_protocol.g4 by ANTLR 4.9.3. DO NOT EDIT.

package memcached_protocol

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 36, 312,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 6,
	33, 293, 10, 33, 13, 33, 14, 33, 294, 3, 34, 6, 34, 298, 10, 34, 13, 34,
	14, 34, 299, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 6, 37, 307, 10, 37, 13,
	37, 14, 37, 308, 3, 37, 3, 37, 2, 2, 38, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49,
	26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67,
	35, 69, 2, 71, 2, 73, 36, 3, 2, 3, 5, 2, 11, 12, 14, 15, 34, 34, 2, 312,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41,
	3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2,
	49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2,
	2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2,
	2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 3, 75, 3, 2,
	2, 2, 5, 79, 3, 2, 2, 2, 7, 83, 3, 2, 2, 2, 9, 87, 3, 2, 2, 2, 11, 95,
	3, 2, 2, 2, 13, 102, 3, 2, 2, 2, 15, 110, 3, 2, 2, 2, 17, 114, 3, 2, 2,
	2, 19, 119, 3, 2, 2, 2, 21, 126, 3, 2, 2, 2, 23, 131, 3, 2, 2, 2, 25, 136,
	3, 2, 2, 2, 27, 142, 3, 2, 2, 2, 29, 148, 3, 2, 2, 2, 31, 154, 3, 2, 2,
	2, 33, 160, 3, 2, 2, 2, 35, 170, 3, 2, 2, 2, 37, 178, 3, 2, 2, 2, 39, 188,
	3, 2, 2, 2, 41, 193, 3, 2, 2, 2, 43, 200, 3, 2, 2, 2, 45, 211, 3, 2, 2,
	2, 47, 218, 3, 2, 2, 2, 49, 228, 3, 2, 2, 2, 51, 234, 3, 2, 2, 2, 53, 242,
	3, 2, 2, 2, 55, 247, 3, 2, 2, 2, 57, 253, 3, 2, 2, 2, 59, 266, 3, 2, 2,
	2, 61, 279, 3, 2, 2, 2, 63, 283, 3, 2, 2, 2, 65, 292, 3, 2, 2, 2, 67, 297,
	3, 2, 2, 2, 69, 301, 3, 2, 2, 2, 71, 303, 3, 2, 2, 2, 73, 306, 3, 2, 2,
	2, 75, 76, 7, 101, 2, 2, 76, 77, 7, 99, 2, 2, 77, 78, 7, 117, 2, 2, 78,
	4, 3, 2, 2, 2, 79, 80, 7, 117, 2, 2, 80, 81, 7, 103, 2, 2, 81, 82, 7, 118,
	2, 2, 82, 6, 3, 2, 2, 2, 83, 84, 7, 99, 2, 2, 84, 85, 7, 102, 2, 2, 85,
	86, 7, 102, 2, 2, 86, 8, 3, 2, 2, 2, 87, 88, 7, 116, 2, 2, 88, 89, 7, 103,
	2, 2, 89, 90, 7, 114, 2, 2, 90, 91, 7, 110, 2, 2, 91, 92, 7, 99, 2, 2,
	92, 93, 7, 101, 2, 2, 93, 94, 7, 103, 2, 2, 94, 10, 3, 2, 2, 2, 95, 96,
	7, 99, 2, 2, 96, 97, 7, 114, 2, 2, 97, 98, 7, 114, 2, 2, 98, 99, 7, 103,
	2, 2, 99, 100, 7, 112, 2, 2, 100, 101, 7, 102, 2, 2, 101, 12, 3, 2, 2,
	2, 102, 103, 7, 114, 2, 2, 103, 104, 7, 116, 2, 2, 104, 105, 7, 103, 2,
	2, 105, 106, 7, 114, 2, 2, 106, 107, 7, 103, 2, 2, 107, 108, 7, 112, 2,
	2, 108, 109, 7, 102, 2, 2, 109, 14, 3, 2, 2, 2, 110, 111, 7, 105, 2, 2,
	111, 112, 7, 103, 2, 2, 112, 113, 7, 118, 2, 2, 113, 16, 3, 2, 2, 2, 114,
	115, 7, 105, 2, 2, 115, 116, 7, 103, 2, 2, 116, 117, 7, 118, 2, 2, 117,
	118, 7, 117, 2, 2, 118, 18, 3, 2, 2, 2, 119, 120, 7, 102, 2, 2, 120, 121,
	7, 103, 2, 2, 121, 122, 7, 110, 2, 2, 122, 123, 7, 103, 2, 2, 123, 124,
	7, 118, 2, 2, 124, 125, 7, 103, 2, 2, 125, 20, 3, 2, 2, 2, 126, 127, 7,
	107, 2, 2, 127, 128, 7, 112, 2, 2, 128, 129, 7, 101, 2, 2, 129, 130, 7,
	116, 2, 2, 130, 22, 3, 2, 2, 2, 131, 132, 7, 102, 2, 2, 132, 133, 7, 103,
	2, 2, 133, 134, 7, 101, 2, 2, 134, 135, 7, 116, 2, 2, 135, 24, 3, 2, 2,
	2, 136, 137, 7, 117, 2, 2, 137, 138, 7, 118, 2, 2, 138, 139, 7, 99, 2,
	2, 139, 140, 7, 118, 2, 2, 140, 141, 7, 117, 2, 2, 141, 26, 3, 2, 2, 2,
	142, 143, 7, 107, 2, 2, 143, 144, 7, 118, 2, 2, 144, 145, 7, 103, 2, 2,
	145, 146, 7, 111, 2, 2, 146, 147, 7, 117, 2, 2, 147, 28, 3, 2, 2, 2, 148,
	149, 7, 117, 2, 2, 149, 150, 7, 110, 2, 2, 150, 151, 7, 99, 2, 2, 151,
	152, 7, 100, 2, 2, 152, 153, 7, 117, 2, 2, 153, 30, 3, 2, 2, 2, 154, 155,
	7, 117, 2, 2, 155, 156, 7, 107, 2, 2, 156, 157, 7, 124, 2, 2, 157, 158,
	7, 103, 2, 2, 158, 159, 7, 117, 2, 2, 159, 32, 3, 2, 2, 2, 160, 161, 7,
	104, 2, 2, 161, 162, 7, 110, 2, 2, 162, 163, 7, 119, 2, 2, 163, 164, 7,
	117, 2, 2, 164, 165, 7, 106, 2, 2, 165, 166, 7, 97, 2, 2, 166, 167, 7,
	99, 2, 2, 167, 168, 7, 110, 2, 2, 168, 169, 7, 110, 2, 2, 169, 34, 3, 2,
	2, 2, 170, 171, 7, 120, 2, 2, 171, 172, 7, 103, 2, 2, 172, 173, 7, 116,
	2, 2, 173, 174, 7, 117, 2, 2, 174, 175, 7, 107, 2, 2, 175, 176, 7, 113,
	2, 2, 176, 177, 7, 112, 2, 2, 177, 36, 3, 2, 2, 2, 178, 179, 7, 120, 2,
	2, 179, 180, 7, 103, 2, 2, 180, 181, 7, 116, 2, 2, 181, 182, 7, 100, 2,
	2, 182, 183, 7, 113, 2, 2, 183, 184, 7, 117, 2, 2, 184, 185, 7, 107, 2,
	2, 185, 186, 7, 118, 2, 2, 186, 187, 7, 123, 2, 2, 187, 38, 3, 2, 2, 2,
	188, 189, 7, 115, 2, 2, 189, 190, 7, 119, 2, 2, 190, 191, 7, 107, 2, 2,
	191, 192, 7, 118, 2, 2, 192, 40, 3, 2, 2, 2, 193, 194, 7, 85, 2, 2, 194,
	195, 7, 86, 2, 2, 195, 196, 7, 81, 2, 2, 196, 197, 7, 84, 2, 2, 197, 198,
	7, 71, 2, 2, 198, 199, 7, 70, 2, 2, 199, 42, 3, 2, 2, 2, 200, 201, 7, 80,
	2, 2, 201, 202, 7, 81, 2, 2, 202, 203, 7, 86, 2, 2, 203, 204, 7, 97, 2,
	2, 204, 205, 7, 85, 2, 2, 205, 206, 7, 86, 2, 2, 206, 207, 7, 81, 2, 2,
	207, 208, 7, 84, 2, 2, 208, 209, 7, 71, 2, 2, 209, 210, 7, 70, 2, 2, 210,
	44, 3, 2, 2, 2, 211, 212, 7, 71, 2, 2, 212, 213, 7, 90, 2, 2, 213, 214,
	7, 75, 2, 2, 214, 215, 7, 85, 2, 2, 215, 216, 7, 86, 2, 2, 216, 217, 7,
	85, 2, 2, 217, 46, 3, 2, 2, 2, 218, 219, 7, 80, 2, 2, 219, 220, 7, 81,
	2, 2, 220, 221, 7, 86, 2, 2, 221, 222, 7, 97, 2, 2, 222, 223, 7, 72, 2,
	2, 223, 224, 7, 81, 2, 2, 224, 225, 7, 87, 2, 2, 225, 226, 7, 80, 2, 2,
	226, 227, 7, 70, 2, 2, 227, 48, 3, 2, 2, 2, 228, 229, 7, 88, 2, 2, 229,
	230, 7, 67, 2, 2, 230, 231, 7, 78, 2, 2, 231, 232, 7, 87, 2, 2, 232, 233,
	7, 71, 2, 2, 233, 50, 3, 2, 2, 2, 234, 235, 7, 70, 2, 2, 235, 236, 7, 71,
	2, 2, 236, 237, 7, 78, 2, 2, 237, 238, 7, 71, 2, 2, 238, 239, 7, 86, 2,
	2, 239, 240, 7, 71, 2, 2, 240, 241, 7, 70, 2, 2, 241, 52, 3, 2, 2, 2, 242,
	243, 7, 85, 2, 2, 243, 244, 7, 86, 2, 2, 244, 245, 7, 67, 2, 2, 245, 246,
	7, 86, 2, 2, 246, 54, 3, 2, 2, 2, 247, 248, 7, 71, 2, 2, 248, 249, 7, 84,
	2, 2, 249, 250, 7, 84, 2, 2, 250, 251, 7, 81, 2, 2, 251, 252, 7, 84, 2,
	2, 252, 56, 3, 2, 2, 2, 253, 254, 7, 69, 2, 2, 254, 255, 7, 78, 2, 2, 255,
	256, 7, 75, 2, 2, 256, 257, 7, 71, 2, 2, 257, 258, 7, 80, 2, 2, 258, 259,
	7, 86, 2, 2, 259, 260, 7, 97, 2, 2, 260, 261, 7, 71, 2, 2, 261, 262, 7,
	84, 2, 2, 262, 263, 7, 84, 2, 2, 263, 264, 7, 81, 2, 2, 264, 265, 7, 84,
	2, 2, 265, 58, 3, 2, 2, 2, 266, 267, 7, 85, 2, 2, 267, 268, 7, 71, 2, 2,
	268, 269, 7, 84, 2, 2, 269, 270, 7, 88, 2, 2, 270, 271, 7, 71, 2, 2, 271,
	272, 7, 84, 2, 2, 272, 273, 7, 97, 2, 2, 273, 274, 7, 71, 2, 2, 274, 275,
	7, 84, 2, 2, 275, 276, 7, 84, 2, 2, 276, 277, 7, 81, 2, 2, 277, 278, 7,
	84, 2, 2, 278, 60, 3, 2, 2, 2, 279, 280, 7, 71, 2, 2, 280, 281, 7, 80,
	2, 2, 281, 282, 7, 70, 2, 2, 282, 62, 3, 2, 2, 2, 283, 284, 7, 112, 2,
	2, 284, 285, 7, 113, 2, 2, 285, 286, 7, 116, 2, 2, 286, 287, 7, 103, 2,
	2, 287, 288, 7, 114, 2, 2, 288, 289, 7, 110, 2, 2, 289, 290, 7, 123, 2,
	2, 290, 64, 3, 2, 2, 2, 291, 293, 5, 69, 35, 2, 292, 291, 3, 2, 2, 2, 293,
	294, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 66, 3,
	2, 2, 2, 296, 298, 5, 71, 36, 2, 297, 296, 3, 2, 2, 2, 298, 299, 3, 2,
	2, 2, 299, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 68, 3, 2, 2, 2,
	301, 302, 4, 50, 59, 2, 302, 70, 3, 2, 2, 2, 303, 304, 4, 35, 128, 2, 304,
	72, 3, 2, 2, 2, 305, 307, 9, 2, 2, 2, 306, 305, 3, 2, 2, 2, 307, 308, 3,
	2, 2, 2, 308, 306, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 310, 3, 2, 2,
	2, 310, 311, 8, 37, 2, 2, 311, 74, 3, 2, 2, 2, 6, 2, 294, 299, 308, 3,
	8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'cas'", "'set'", "'add'", "'replace'", "'append'", "'prepend'", "'get'",
	"'gets'", "'delete'", "'incr'", "'decr'", "'stats'", "'items'", "'slabs'",
	"'sizes'", "'flush_all'", "'version'", "'verbosity'", "'quit'", "'STORED'",
	"'NOT_STORED'", "'EXISTS'", "'NOT_FOUND'", "'VALUE'", "'DELETED'", "'STAT'",
	"'ERROR'", "'CLIENT_ERROR'", "'SERVER_ERROR'", "'END'", "'noreply'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "INTEGER", "WORD",
	"WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "INTEGER", "WORD",
	"DIGIT", "PRINTABLE_CHAR", "WHITESPACE",
}

type memcached_protocolLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// Newmemcached_protocolLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *memcached_protocolLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newmemcached_protocolLexer(input antlr.CharStream) *memcached_protocolLexer {
	l := new(memcached_protocolLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "memcached_protocol.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// memcached_protocolLexer tokens.
const (
	memcached_protocolLexerT__0       = 1
	memcached_protocolLexerT__1       = 2
	memcached_protocolLexerT__2       = 3
	memcached_protocolLexerT__3       = 4
	memcached_protocolLexerT__4       = 5
	memcached_protocolLexerT__5       = 6
	memcached_protocolLexerT__6       = 7
	memcached_protocolLexerT__7       = 8
	memcached_protocolLexerT__8       = 9
	memcached_protocolLexerT__9       = 10
	memcached_protocolLexerT__10      = 11
	memcached_protocolLexerT__11      = 12
	memcached_protocolLexerT__12      = 13
	memcached_protocolLexerT__13      = 14
	memcached_protocolLexerT__14      = 15
	memcached_protocolLexerT__15      = 16
	memcached_protocolLexerT__16      = 17
	memcached_protocolLexerT__17      = 18
	memcached_protocolLexerT__18      = 19
	memcached_protocolLexerT__19      = 20
	memcached_protocolLexerT__20      = 21
	memcached_protocolLexerT__21      = 22
	memcached_protocolLexerT__22      = 23
	memcached_protocolLexerT__23      = 24
	memcached_protocolLexerT__24      = 25
	memcached_protocolLexerT__25      = 26
	memcached_protocolLexerT__26      = 27
	memcached_protocolLexerT__27      = 28
	memcached_protocolLexerT__28      = 29
	memcached_protocolLexerT__29      = 30
	memcached_protocolLexerT__30      = 31
	memcached_protocolLexerINTEGER    = 32
	memcached_protocolLexerWORD       = 33
	memcached_protocolLexerWHITESPACE = 34
)
