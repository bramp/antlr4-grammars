// Code generated from Erlang.g4 by ANTLR 4.9.3. DO NOT EDIT.

package erlang

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 75, 524,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3,
	28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 35,
	3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3,
	47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3,
	52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55,
	3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3,
	58, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3, 62,
	3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3,
	65, 3, 66, 3, 66, 3, 67, 5, 67, 378, 10, 67, 3, 68, 5, 68, 381, 10, 68,
	3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 7, 69, 388, 10, 69, 12, 69, 14, 69,
	391, 11, 69, 3, 69, 3, 69, 3, 69, 3, 69, 5, 69, 397, 10, 69, 3, 69, 7,
	69, 400, 10, 69, 12, 69, 14, 69, 403, 11, 69, 3, 69, 5, 69, 406, 10, 69,
	3, 70, 3, 70, 5, 70, 410, 10, 70, 3, 70, 3, 70, 3, 70, 3, 70, 7, 70, 416,
	10, 70, 12, 70, 14, 70, 419, 11, 70, 3, 71, 5, 71, 422, 10, 71, 3, 71,
	6, 71, 425, 10, 71, 13, 71, 14, 71, 426, 3, 71, 3, 71, 6, 71, 431, 10,
	71, 13, 71, 14, 71, 432, 3, 71, 3, 71, 5, 71, 437, 10, 71, 3, 71, 6, 71,
	440, 10, 71, 13, 71, 14, 71, 441, 5, 71, 444, 10, 71, 3, 72, 5, 72, 447,
	10, 72, 3, 72, 6, 72, 450, 10, 72, 13, 72, 14, 72, 451, 3, 72, 3, 72, 3,
	72, 6, 72, 457, 10, 72, 13, 72, 14, 72, 458, 5, 72, 461, 10, 72, 3, 73,
	3, 73, 5, 73, 465, 10, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 5,
	73, 473, 10, 73, 3, 74, 3, 74, 3, 74, 3, 74, 5, 74, 479, 10, 74, 3, 74,
	7, 74, 482, 10, 74, 12, 74, 14, 74, 485, 11, 74, 3, 74, 3, 74, 3, 75, 3,
	75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75,
	3, 75, 5, 75, 502, 10, 75, 3, 76, 3, 76, 7, 76, 506, 10, 76, 12, 76, 14,
	76, 509, 11, 76, 3, 76, 5, 76, 512, 10, 76, 3, 76, 3, 76, 3, 76, 3, 76,
	3, 77, 6, 77, 519, 10, 77, 13, 77, 14, 77, 520, 3, 77, 3, 77, 2, 2, 78,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41,
	22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59,
	31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77,
	40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95,
	49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57,
	113, 58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64, 127, 65,
	129, 66, 131, 2, 133, 2, 135, 2, 137, 67, 139, 68, 141, 69, 143, 70, 145,
	71, 147, 72, 149, 73, 151, 74, 153, 75, 3, 2, 14, 3, 2, 50, 59, 5, 2, 99,
	124, 225, 248, 250, 257, 5, 2, 67, 92, 194, 216, 218, 224, 4, 2, 66, 66,
	97, 97, 3, 2, 94, 94, 4, 2, 41, 41, 94, 94, 4, 2, 71, 71, 103, 103, 4,
	2, 45, 45, 47, 47, 4, 2, 67, 92, 99, 124, 4, 2, 12, 12, 15, 15, 4, 2, 36,
	36, 94, 94, 4, 2, 2, 34, 130, 162, 2, 553, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2,
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2,
	2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2,
	2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3,
	2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67,
	3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2,
	75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2,
	2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2,
	2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2,
	2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105,
	3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2,
	2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3,
	2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2,
	127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2,
	2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147,
	3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2,
	3, 155, 3, 2, 2, 2, 5, 157, 3, 2, 2, 2, 7, 159, 3, 2, 2, 2, 9, 161, 3,
	2, 2, 2, 11, 163, 3, 2, 2, 2, 13, 165, 3, 2, 2, 2, 15, 167, 3, 2, 2, 2,
	17, 170, 3, 2, 2, 2, 19, 172, 3, 2, 2, 2, 21, 174, 3, 2, 2, 2, 23, 176,
	3, 2, 2, 2, 25, 181, 3, 2, 2, 2, 27, 183, 3, 2, 2, 2, 29, 186, 3, 2, 2,
	2, 31, 188, 3, 2, 2, 2, 33, 190, 3, 2, 2, 2, 35, 194, 3, 2, 2, 2, 37, 196,
	3, 2, 2, 2, 39, 200, 3, 2, 2, 2, 41, 203, 3, 2, 2, 2, 43, 206, 3, 2, 2,
	2, 45, 209, 3, 2, 2, 2, 47, 212, 3, 2, 2, 2, 49, 215, 3, 2, 2, 2, 51, 217,
	3, 2, 2, 2, 53, 223, 3, 2, 2, 2, 55, 225, 3, 2, 2, 2, 57, 227, 3, 2, 2,
	2, 59, 234, 3, 2, 2, 2, 61, 242, 3, 2, 2, 2, 63, 248, 3, 2, 2, 2, 65, 252,
	3, 2, 2, 2, 67, 254, 3, 2, 2, 2, 69, 257, 3, 2, 2, 2, 71, 260, 3, 2, 2,
	2, 73, 263, 3, 2, 2, 2, 75, 266, 3, 2, 2, 2, 77, 271, 3, 2, 2, 2, 79, 274,
	3, 2, 2, 2, 81, 282, 3, 2, 2, 2, 83, 288, 3, 2, 2, 2, 85, 292, 3, 2, 2,
	2, 87, 294, 3, 2, 2, 2, 89, 299, 3, 2, 2, 2, 91, 303, 3, 2, 2, 2, 93, 307,
	3, 2, 2, 2, 95, 311, 3, 2, 2, 2, 97, 316, 3, 2, 2, 2, 99, 320, 3, 2, 2,
	2, 101, 324, 3, 2, 2, 2, 103, 329, 3, 2, 2, 2, 105, 333, 3, 2, 2, 2, 107,
	337, 3, 2, 2, 2, 109, 340, 3, 2, 2, 2, 111, 344, 3, 2, 2, 2, 113, 347,
	3, 2, 2, 2, 115, 350, 3, 2, 2, 2, 117, 353, 3, 2, 2, 2, 119, 356, 3, 2,
	2, 2, 121, 359, 3, 2, 2, 2, 123, 361, 3, 2, 2, 2, 125, 364, 3, 2, 2, 2,
	127, 366, 3, 2, 2, 2, 129, 370, 3, 2, 2, 2, 131, 374, 3, 2, 2, 2, 133,
	377, 3, 2, 2, 2, 135, 380, 3, 2, 2, 2, 137, 405, 3, 2, 2, 2, 139, 409,
	3, 2, 2, 2, 141, 421, 3, 2, 2, 2, 143, 446, 3, 2, 2, 2, 145, 462, 3, 2,
	2, 2, 147, 474, 3, 2, 2, 2, 149, 488, 3, 2, 2, 2, 151, 503, 3, 2, 2, 2,
	153, 518, 3, 2, 2, 2, 155, 156, 7, 48, 2, 2, 156, 4, 3, 2, 2, 2, 157, 158,
	7, 47, 2, 2, 158, 6, 3, 2, 2, 2, 159, 160, 7, 42, 2, 2, 160, 8, 3, 2, 2,
	2, 161, 162, 7, 43, 2, 2, 162, 10, 3, 2, 2, 2, 163, 164, 7, 60, 2, 2, 164,
	12, 3, 2, 2, 2, 165, 166, 7, 46, 2, 2, 166, 14, 3, 2, 2, 2, 167, 168, 7,
	60, 2, 2, 168, 169, 7, 60, 2, 2, 169, 16, 3, 2, 2, 2, 170, 171, 7, 125,
	2, 2, 171, 18, 3, 2, 2, 2, 172, 173, 7, 127, 2, 2, 173, 20, 3, 2, 2, 2,
	174, 175, 7, 61, 2, 2, 175, 22, 3, 2, 2, 2, 176, 177, 7, 121, 2, 2, 177,
	178, 7, 106, 2, 2, 178, 179, 7, 103, 2, 2, 179, 180, 7, 112, 2, 2, 180,
	24, 3, 2, 2, 2, 181, 182, 7, 126, 2, 2, 182, 26, 3, 2, 2, 2, 183, 184,
	7, 48, 2, 2, 184, 185, 7, 48, 2, 2, 185, 28, 3, 2, 2, 2, 186, 187, 7, 93,
	2, 2, 187, 30, 3, 2, 2, 2, 188, 189, 7, 95, 2, 2, 189, 32, 3, 2, 2, 2,
	190, 191, 7, 48, 2, 2, 191, 192, 7, 48, 2, 2, 192, 193, 7, 48, 2, 2, 193,
	34, 3, 2, 2, 2, 194, 195, 7, 37, 2, 2, 195, 36, 3, 2, 2, 2, 196, 197, 7,
	104, 2, 2, 197, 198, 7, 119, 2, 2, 198, 199, 7, 112, 2, 2, 199, 38, 3,
	2, 2, 2, 200, 201, 7, 47, 2, 2, 201, 202, 7, 64, 2, 2, 202, 40, 3, 2, 2,
	2, 203, 204, 7, 63, 2, 2, 204, 205, 7, 64, 2, 2, 205, 42, 3, 2, 2, 2, 206,
	207, 7, 60, 2, 2, 207, 208, 7, 63, 2, 2, 208, 44, 3, 2, 2, 2, 209, 210,
	7, 62, 2, 2, 210, 211, 7, 62, 2, 2, 211, 46, 3, 2, 2, 2, 212, 213, 7, 64,
	2, 2, 213, 214, 7, 64, 2, 2, 214, 48, 3, 2, 2, 2, 215, 216, 7, 44, 2, 2,
	216, 50, 3, 2, 2, 2, 217, 218, 7, 101, 2, 2, 218, 219, 7, 99, 2, 2, 219,
	220, 7, 118, 2, 2, 220, 221, 7, 101, 2, 2, 221, 222, 7, 106, 2, 2, 222,
	52, 3, 2, 2, 2, 223, 224, 7, 63, 2, 2, 224, 54, 3, 2, 2, 2, 225, 226, 7,
	35, 2, 2, 226, 56, 3, 2, 2, 2, 227, 228, 7, 113, 2, 2, 228, 229, 7, 116,
	2, 2, 229, 230, 7, 103, 2, 2, 230, 231, 7, 110, 2, 2, 231, 232, 7, 117,
	2, 2, 232, 233, 7, 103, 2, 2, 233, 58, 3, 2, 2, 2, 234, 235, 7, 99, 2,
	2, 235, 236, 7, 112, 2, 2, 236, 237, 7, 102, 2, 2, 237, 238, 7, 99, 2,
	2, 238, 239, 7, 110, 2, 2, 239, 240, 7, 117, 2, 2, 240, 241, 7, 113, 2,
	2, 241, 60, 3, 2, 2, 2, 242, 243, 7, 100, 2, 2, 243, 244, 7, 103, 2, 2,
	244, 245, 7, 105, 2, 2, 245, 246, 7, 107, 2, 2, 246, 247, 7, 112, 2, 2,
	247, 62, 3, 2, 2, 2, 248, 249, 7, 103, 2, 2, 249, 250, 7, 112, 2, 2, 250,
	251, 7, 102, 2, 2, 251, 64, 3, 2, 2, 2, 252, 253, 7, 49, 2, 2, 253, 66,
	3, 2, 2, 2, 254, 255, 7, 126, 2, 2, 255, 256, 7, 126, 2, 2, 256, 68, 3,
	2, 2, 2, 257, 258, 7, 62, 2, 2, 258, 259, 7, 47, 2, 2, 259, 70, 3, 2, 2,
	2, 260, 261, 7, 62, 2, 2, 261, 262, 7, 63, 2, 2, 262, 72, 3, 2, 2, 2, 263,
	264, 7, 107, 2, 2, 264, 265, 7, 104, 2, 2, 265, 74, 3, 2, 2, 2, 266, 267,
	7, 101, 2, 2, 267, 268, 7, 99, 2, 2, 268, 269, 7, 117, 2, 2, 269, 270,
	7, 103, 2, 2, 270, 76, 3, 2, 2, 2, 271, 272, 7, 113, 2, 2, 272, 273, 7,
	104, 2, 2, 273, 78, 3, 2, 2, 2, 274, 275, 7, 116, 2, 2, 275, 276, 7, 103,
	2, 2, 276, 277, 7, 101, 2, 2, 277, 278, 7, 103, 2, 2, 278, 279, 7, 107,
	2, 2, 279, 280, 7, 120, 2, 2, 280, 281, 7, 103, 2, 2, 281, 80, 3, 2, 2,
	2, 282, 283, 7, 99, 2, 2, 283, 284, 7, 104, 2, 2, 284, 285, 7, 118, 2,
	2, 285, 286, 7, 103, 2, 2, 286, 287, 7, 116, 2, 2, 287, 82, 3, 2, 2, 2,
	288, 289, 7, 118, 2, 2, 289, 290, 7, 116, 2, 2, 290, 291, 7, 123, 2, 2,
	291, 84, 3, 2, 2, 2, 292, 293, 7, 45, 2, 2, 293, 86, 3, 2, 2, 2, 294, 295,
	7, 100, 2, 2, 295, 296, 7, 112, 2, 2, 296, 297, 7, 113, 2, 2, 297, 298,
	7, 118, 2, 2, 298, 88, 3, 2, 2, 2, 299, 300, 7, 112, 2, 2, 300, 301, 7,
	113, 2, 2, 301, 302, 7, 118, 2, 2, 302, 90, 3, 2, 2, 2, 303, 304, 7, 102,
	2, 2, 304, 305, 7, 107, 2, 2, 305, 306, 7, 120, 2, 2, 306, 92, 3, 2, 2,
	2, 307, 308, 7, 116, 2, 2, 308, 309, 7, 103, 2, 2, 309, 310, 7, 111, 2,
	2, 310, 94, 3, 2, 2, 2, 311, 312, 7, 100, 2, 2, 312, 313, 7, 99, 2, 2,
	313, 314, 7, 112, 2, 2, 314, 315, 7, 102, 2, 2, 315, 96, 3, 2, 2, 2, 316,
	317, 7, 99, 2, 2, 317, 318, 7, 112, 2, 2, 318, 319, 7, 102, 2, 2, 319,
	98, 3, 2, 2, 2, 320, 321, 7, 100, 2, 2, 321, 322, 7, 113, 2, 2, 322, 323,
	7, 116, 2, 2, 323, 100, 3, 2, 2, 2, 324, 325, 7, 100, 2, 2, 325, 326, 7,
	122, 2, 2, 326, 327, 7, 113, 2, 2, 327, 328, 7, 116, 2, 2, 328, 102, 3,
	2, 2, 2, 329, 330, 7, 100, 2, 2, 330, 331, 7, 117, 2, 2, 331, 332, 7, 110,
	2, 2, 332, 104, 3, 2, 2, 2, 333, 334, 7, 100, 2, 2, 334, 335, 7, 117, 2,
	2, 335, 336, 7, 116, 2, 2, 336, 106, 3, 2, 2, 2, 337, 338, 7, 113, 2, 2,
	338, 339, 7, 116, 2, 2, 339, 108, 3, 2, 2, 2, 340, 341, 7, 122, 2, 2, 341,
	342, 7, 113, 2, 2, 342, 343, 7, 116, 2, 2, 343, 110, 3, 2, 2, 2, 344, 345,
	7, 45, 2, 2, 345, 346, 7, 45, 2, 2, 346, 112, 3, 2, 2, 2, 347, 348, 7,
	47, 2, 2, 348, 349, 7, 47, 2, 2, 349, 114, 3, 2, 2, 2, 350, 351, 7, 63,
	2, 2, 351, 352, 7, 63, 2, 2, 352, 116, 3, 2, 2, 2, 353, 354, 7, 49, 2,
	2, 354, 355, 7, 63, 2, 2, 355, 118, 3, 2, 2, 2, 356, 357, 7, 63, 2, 2,
	357, 358, 7, 62, 2, 2, 358, 120, 3, 2, 2, 2, 359, 360, 7, 62, 2, 2, 360,
	122, 3, 2, 2, 2, 361, 362, 7, 64, 2, 2, 362, 363, 7, 63, 2, 2, 363, 124,
	3, 2, 2, 2, 364, 365, 7, 64, 2, 2, 365, 126, 3, 2, 2, 2, 366, 367, 7, 63,
	2, 2, 367, 368, 7, 60, 2, 2, 368, 369, 7, 63, 2, 2, 369, 128, 3, 2, 2,
	2, 370, 371, 7, 63, 2, 2, 371, 372, 7, 49, 2, 2, 372, 373, 7, 63, 2, 2,
	373, 130, 3, 2, 2, 2, 374, 375, 9, 2, 2, 2, 375, 132, 3, 2, 2, 2, 376,
	378, 9, 3, 2, 2, 377, 376, 3, 2, 2, 2, 378, 134, 3, 2, 2, 2, 379, 381,
	9, 4, 2, 2, 380, 379, 3, 2, 2, 2, 381, 136, 3, 2, 2, 2, 382, 389, 5, 133,
	67, 2, 383, 388, 5, 131, 66, 2, 384, 388, 5, 133, 67, 2, 385, 388, 5, 135,
	68, 2, 386, 388, 9, 5, 2, 2, 387, 383, 3, 2, 2, 2, 387, 384, 3, 2, 2, 2,
	387, 385, 3, 2, 2, 2, 387, 386, 3, 2, 2, 2, 388, 391, 3, 2, 2, 2, 389,
	387, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 406, 3, 2, 2, 2, 391, 389,
	3, 2, 2, 2, 392, 401, 7, 41, 2, 2, 393, 396, 7, 94, 2, 2, 394, 397, 10,
	6, 2, 2, 395, 397, 7, 94, 2, 2, 396, 394, 3, 2, 2, 2, 396, 395, 3, 2, 2,
	2, 397, 400, 3, 2, 2, 2, 398, 400, 10, 7, 2, 2, 399, 393, 3, 2, 2, 2, 399,
	398, 3, 2, 2, 2, 400, 403, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2, 401, 402,
	3, 2, 2, 2, 402, 404, 3, 2, 2, 2, 403, 401, 3, 2, 2, 2, 404, 406, 7, 41,
	2, 2, 405, 382, 3, 2, 2, 2, 405, 392, 3, 2, 2, 2, 406, 138, 3, 2, 2, 2,
	407, 410, 5, 135, 68, 2, 408, 410, 7, 97, 2, 2, 409, 407, 3, 2, 2, 2, 409,
	408, 3, 2, 2, 2, 410, 417, 3, 2, 2, 2, 411, 416, 5, 131, 66, 2, 412, 416,
	5, 133, 67, 2, 413, 416, 5, 135, 68, 2, 414, 416, 9, 5, 2, 2, 415, 411,
	3, 2, 2, 2, 415, 412, 3, 2, 2, 2, 415, 413, 3, 2, 2, 2, 415, 414, 3, 2,
	2, 2, 416, 419, 3, 2, 2, 2, 417, 415, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2,
	418, 140, 3, 2, 2, 2, 419, 417, 3, 2, 2, 2, 420, 422, 7, 47, 2, 2, 421,
	420, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 424, 3, 2, 2, 2, 423, 425,
	5, 131, 66, 2, 424, 423, 3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 426, 424, 3,
	2, 2, 2, 426, 427, 3, 2, 2, 2, 427, 428, 3, 2, 2, 2, 428, 430, 7, 48, 2,
	2, 429, 431, 5, 131, 66, 2, 430, 429, 3, 2, 2, 2, 431, 432, 3, 2, 2, 2,
	432, 430, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 443, 3, 2, 2, 2, 434,
	436, 9, 8, 2, 2, 435, 437, 9, 9, 2, 2, 436, 435, 3, 2, 2, 2, 436, 437,
	3, 2, 2, 2, 437, 439, 3, 2, 2, 2, 438, 440, 5, 131, 66, 2, 439, 438, 3,
	2, 2, 2, 440, 441, 3, 2, 2, 2, 441, 439, 3, 2, 2, 2, 441, 442, 3, 2, 2,
	2, 442, 444, 3, 2, 2, 2, 443, 434, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444,
	142, 3, 2, 2, 2, 445, 447, 7, 47, 2, 2, 446, 445, 3, 2, 2, 2, 446, 447,
	3, 2, 2, 2, 447, 449, 3, 2, 2, 2, 448, 450, 5, 131, 66, 2, 449, 448, 3,
	2, 2, 2, 450, 451, 3, 2, 2, 2, 451, 449, 3, 2, 2, 2, 451, 452, 3, 2, 2,
	2, 452, 460, 3, 2, 2, 2, 453, 456, 7, 37, 2, 2, 454, 457, 5, 131, 66, 2,
	455, 457, 9, 10, 2, 2, 456, 454, 3, 2, 2, 2, 456, 455, 3, 2, 2, 2, 457,
	458, 3, 2, 2, 2, 458, 456, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 461,
	3, 2, 2, 2, 460, 453, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461, 144, 3, 2,
	2, 2, 462, 472, 7, 38, 2, 2, 463, 465, 7, 94, 2, 2, 464, 463, 3, 2, 2,
	2, 464, 465, 3, 2, 2, 2, 465, 466, 3, 2, 2, 2, 466, 473, 10, 11, 2, 2,
	467, 468, 7, 94, 2, 2, 468, 469, 5, 131, 66, 2, 469, 470, 5, 131, 66, 2,
	470, 471, 5, 131, 66, 2, 471, 473, 3, 2, 2, 2, 472, 464, 3, 2, 2, 2, 472,
	467, 3, 2, 2, 2, 473, 146, 3, 2, 2, 2, 474, 483, 7, 36, 2, 2, 475, 478,
	7, 94, 2, 2, 476, 479, 10, 6, 2, 2, 477, 479, 7, 94, 2, 2, 478, 476, 3,
	2, 2, 2, 478, 477, 3, 2, 2, 2, 479, 482, 3, 2, 2, 2, 480, 482, 10, 12,
	2, 2, 481, 475, 3, 2, 2, 2, 481, 480, 3, 2, 2, 2, 482, 485, 3, 2, 2, 2,
	483, 481, 3, 2, 2, 2, 483, 484, 3, 2, 2, 2, 484, 486, 3, 2, 2, 2, 485,
	483, 3, 2, 2, 2, 486, 487, 7, 36, 2, 2, 487, 148, 3, 2, 2, 2, 488, 501,
	7, 47, 2, 2, 489, 490, 7, 117, 2, 2, 490, 491, 7, 114, 2, 2, 491, 492,
	7, 103, 2, 2, 492, 502, 7, 101, 2, 2, 493, 494, 7, 101, 2, 2, 494, 495,
	7, 99, 2, 2, 495, 496, 7, 110, 2, 2, 496, 497, 7, 110, 2, 2, 497, 498,
	7, 100, 2, 2, 498, 499, 7, 99, 2, 2, 499, 500, 7, 101, 2, 2, 500, 502,
	7, 109, 2, 2, 501, 489, 3, 2, 2, 2, 501, 493, 3, 2, 2, 2, 502, 150, 3,
	2, 2, 2, 503, 507, 7, 39, 2, 2, 504, 506, 10, 11, 2, 2, 505, 504, 3, 2,
	2, 2, 506, 509, 3, 2, 2, 2, 507, 505, 3, 2, 2, 2, 507, 508, 3, 2, 2, 2,
	508, 511, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 510, 512, 7, 15, 2, 2, 511,
	510, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2, 513, 514,
	7, 12, 2, 2, 514, 515, 3, 2, 2, 2, 515, 516, 8, 76, 2, 2, 516, 152, 3,
	2, 2, 2, 517, 519, 9, 13, 2, 2, 518, 517, 3, 2, 2, 2, 519, 520, 3, 2, 2,
	2, 520, 518, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521, 522, 3, 2, 2, 2, 522,
	523, 8, 77, 2, 2, 523, 154, 3, 2, 2, 2, 34, 2, 377, 380, 387, 389, 396,
	399, 401, 405, 409, 415, 417, 421, 426, 432, 436, 441, 443, 446, 451, 456,
	458, 460, 464, 472, 478, 481, 483, 501, 507, 511, 520, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'.'", "'-'", "'('", "')'", "':'", "','", "'::'", "'{'", "'}'", "';'",
	"'when'", "'|'", "'..'", "'['", "']'", "'...'", "'#'", "'fun'", "'->'",
	"'=>'", "':='", "'<<'", "'>>'", "'*'", "'catch'", "'='", "'!'", "'orelse'",
	"'andalso'", "'begin'", "'end'", "'/'", "'||'", "'<-'", "'<='", "'if'",
	"'case'", "'of'", "'receive'", "'after'", "'try'", "'+'", "'bnot'", "'not'",
	"'div'", "'rem'", "'band'", "'and'", "'bor'", "'bxor'", "'bsl'", "'bsr'",
	"'or'", "'xor'", "'++'", "'--'", "'=='", "'/='", "'=<'", "'<'", "'>='",
	"'>'", "'=:='", "'=/='",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "TokAtom", "TokVar", "TokFloat",
	"TokInteger", "TokChar", "TokString", "AttrName", "Comment", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
	"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
	"T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "T__55", "T__56",
	"T__57", "T__58", "T__59", "T__60", "T__61", "T__62", "T__63", "DIGIT",
	"LOWERCASE", "UPPERCASE", "TokAtom", "TokVar", "TokFloat", "TokInteger",
	"TokChar", "TokString", "AttrName", "Comment", "WS",
}

type ErlangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewErlangLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ErlangLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewErlangLexer(input antlr.CharStream) *ErlangLexer {
	l := new(ErlangLexer)
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
	l.GrammarFileName = "Erlang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ErlangLexer tokens.
const (
	ErlangLexerT__0       = 1
	ErlangLexerT__1       = 2
	ErlangLexerT__2       = 3
	ErlangLexerT__3       = 4
	ErlangLexerT__4       = 5
	ErlangLexerT__5       = 6
	ErlangLexerT__6       = 7
	ErlangLexerT__7       = 8
	ErlangLexerT__8       = 9
	ErlangLexerT__9       = 10
	ErlangLexerT__10      = 11
	ErlangLexerT__11      = 12
	ErlangLexerT__12      = 13
	ErlangLexerT__13      = 14
	ErlangLexerT__14      = 15
	ErlangLexerT__15      = 16
	ErlangLexerT__16      = 17
	ErlangLexerT__17      = 18
	ErlangLexerT__18      = 19
	ErlangLexerT__19      = 20
	ErlangLexerT__20      = 21
	ErlangLexerT__21      = 22
	ErlangLexerT__22      = 23
	ErlangLexerT__23      = 24
	ErlangLexerT__24      = 25
	ErlangLexerT__25      = 26
	ErlangLexerT__26      = 27
	ErlangLexerT__27      = 28
	ErlangLexerT__28      = 29
	ErlangLexerT__29      = 30
	ErlangLexerT__30      = 31
	ErlangLexerT__31      = 32
	ErlangLexerT__32      = 33
	ErlangLexerT__33      = 34
	ErlangLexerT__34      = 35
	ErlangLexerT__35      = 36
	ErlangLexerT__36      = 37
	ErlangLexerT__37      = 38
	ErlangLexerT__38      = 39
	ErlangLexerT__39      = 40
	ErlangLexerT__40      = 41
	ErlangLexerT__41      = 42
	ErlangLexerT__42      = 43
	ErlangLexerT__43      = 44
	ErlangLexerT__44      = 45
	ErlangLexerT__45      = 46
	ErlangLexerT__46      = 47
	ErlangLexerT__47      = 48
	ErlangLexerT__48      = 49
	ErlangLexerT__49      = 50
	ErlangLexerT__50      = 51
	ErlangLexerT__51      = 52
	ErlangLexerT__52      = 53
	ErlangLexerT__53      = 54
	ErlangLexerT__54      = 55
	ErlangLexerT__55      = 56
	ErlangLexerT__56      = 57
	ErlangLexerT__57      = 58
	ErlangLexerT__58      = 59
	ErlangLexerT__59      = 60
	ErlangLexerT__60      = 61
	ErlangLexerT__61      = 62
	ErlangLexerT__62      = 63
	ErlangLexerT__63      = 64
	ErlangLexerTokAtom    = 65
	ErlangLexerTokVar     = 66
	ErlangLexerTokFloat   = 67
	ErlangLexerTokInteger = 68
	ErlangLexerTokChar    = 69
	ErlangLexerTokString  = 70
	ErlangLexerAttrName   = 71
	ErlangLexerComment    = 72
	ErlangLexerWS         = 73
)
