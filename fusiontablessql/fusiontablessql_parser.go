// Code generated from FusionTablesSql.g4 by ANTLR 4.7.2. DO NOT EDIT.

package fusiontablessql // FusionTablesSql
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 80, 516,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	3, 2, 7, 2, 90, 10, 2, 12, 2, 14, 2, 93, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 105, 10, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7,
	10, 144, 10, 10, 12, 10, 14, 10, 147, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	5, 10, 153, 10, 10, 3, 10, 7, 10, 156, 10, 10, 12, 10, 14, 10, 159, 11,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 7, 12, 174, 10, 12, 12, 12, 14, 12, 177, 11, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 186, 10, 12, 12, 12,
	14, 12, 189, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 7, 13, 199, 10, 13, 12, 13, 14, 13, 202, 11, 13, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 5, 15, 219, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 5, 17, 228, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 234, 10,
	18, 12, 18, 14, 18, 237, 11, 18, 3, 18, 3, 18, 3, 18, 7, 18, 242, 10, 18,
	12, 18, 14, 18, 245, 11, 18, 3, 18, 3, 18, 5, 18, 249, 10, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 7, 18, 256, 10, 18, 12, 18, 14, 18, 259, 11, 18,
	5, 18, 261, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 268, 10,
	18, 12, 18, 14, 18, 271, 11, 18, 5, 18, 273, 10, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 5, 18, 279, 10, 18, 5, 18, 281, 10, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 291, 10, 19, 3, 19, 5, 19, 294,
	10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 312, 10, 21, 3,
	22, 3, 22, 3, 22, 5, 22, 317, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 332, 10,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 351, 10, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 5, 24, 357, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 7, 24, 365, 10, 24, 12, 24, 14, 24, 368, 11, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 5, 24, 374, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 5, 24, 384, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 395, 10, 24, 5, 24, 397, 10, 24,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 5, 28, 407, 10,
	28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 488, 10, 32, 3, 33, 3,
	33, 3, 34, 3, 34, 5, 34, 494, 10, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37,
	3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3,
	42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 2, 2, 45, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
	2, 6, 4, 2, 11, 11, 22, 22, 6, 2, 12, 12, 19, 19, 41, 42, 55, 55, 3, 2,
	8, 9, 3, 2, 67, 71, 2, 581, 2, 91, 3, 2, 2, 2, 4, 104, 3, 2, 2, 2, 6, 108,
	3, 2, 2, 2, 8, 110, 3, 2, 2, 2, 10, 112, 3, 2, 2, 2, 12, 121, 3, 2, 2,
	2, 14, 124, 3, 2, 2, 2, 16, 127, 3, 2, 2, 2, 18, 134, 3, 2, 2, 2, 20, 162,
	3, 2, 2, 2, 22, 166, 3, 2, 2, 2, 24, 192, 3, 2, 2, 2, 26, 206, 3, 2, 2,
	2, 28, 210, 3, 2, 2, 2, 30, 220, 3, 2, 2, 2, 32, 224, 3, 2, 2, 2, 34, 229,
	3, 2, 2, 2, 36, 290, 3, 2, 2, 2, 38, 295, 3, 2, 2, 2, 40, 311, 3, 2, 2,
	2, 42, 316, 3, 2, 2, 2, 44, 320, 3, 2, 2, 2, 46, 396, 3, 2, 2, 2, 48, 398,
	3, 2, 2, 2, 50, 400, 3, 2, 2, 2, 52, 402, 3, 2, 2, 2, 54, 406, 3, 2, 2,
	2, 56, 408, 3, 2, 2, 2, 58, 415, 3, 2, 2, 2, 60, 422, 3, 2, 2, 2, 62, 487,
	3, 2, 2, 2, 64, 489, 3, 2, 2, 2, 66, 493, 3, 2, 2, 2, 68, 495, 3, 2, 2,
	2, 70, 497, 3, 2, 2, 2, 72, 499, 3, 2, 2, 2, 74, 501, 3, 2, 2, 2, 76, 503,
	3, 2, 2, 2, 78, 505, 3, 2, 2, 2, 80, 507, 3, 2, 2, 2, 82, 509, 3, 2, 2,
	2, 84, 511, 3, 2, 2, 2, 86, 513, 3, 2, 2, 2, 88, 90, 5, 4, 3, 2, 89, 88,
	3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2,
	92, 3, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 105, 5, 16, 9, 2, 95, 105, 5,
	34, 18, 2, 96, 105, 5, 18, 10, 2, 97, 105, 5, 10, 6, 2, 98, 105, 5, 28,
	15, 2, 99, 105, 5, 20, 11, 2, 100, 105, 5, 22, 12, 2, 101, 105, 5, 24,
	13, 2, 102, 105, 5, 12, 7, 2, 103, 105, 5, 14, 8, 2, 104, 94, 3, 2, 2,
	2, 104, 95, 3, 2, 2, 2, 104, 96, 3, 2, 2, 2, 104, 97, 3, 2, 2, 2, 104,
	98, 3, 2, 2, 2, 104, 99, 3, 2, 2, 2, 104, 100, 3, 2, 2, 2, 104, 101, 3,
	2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2,
	2, 106, 107, 7, 3, 2, 2, 107, 5, 3, 2, 2, 2, 108, 109, 5, 74, 38, 2, 109,
	7, 3, 2, 2, 2, 110, 111, 5, 74, 38, 2, 111, 9, 3, 2, 2, 2, 112, 113, 7,
	20, 2, 2, 113, 114, 7, 59, 2, 2, 114, 115, 5, 70, 36, 2, 115, 116, 7, 10,
	2, 2, 116, 117, 7, 53, 2, 2, 117, 118, 7, 4, 2, 2, 118, 119, 7, 28, 2,
	2, 119, 120, 5, 6, 4, 2, 120, 11, 3, 2, 2, 2, 121, 122, 7, 23, 2, 2, 122,
	123, 5, 6, 4, 2, 123, 13, 3, 2, 2, 2, 124, 125, 7, 57, 2, 2, 125, 126,
	7, 60, 2, 2, 126, 15, 3, 2, 2, 2, 127, 128, 7, 7, 2, 2, 128, 129, 7, 59,
	2, 2, 129, 130, 5, 6, 4, 2, 130, 131, 7, 51, 2, 2, 131, 132, 7, 61, 2,
	2, 132, 133, 5, 70, 36, 2, 133, 17, 3, 2, 2, 2, 134, 135, 7, 20, 2, 2,
	135, 136, 7, 64, 2, 2, 136, 137, 5, 80, 41, 2, 137, 138, 7, 10, 2, 2, 138,
	139, 7, 72, 2, 2, 139, 140, 7, 53, 2, 2, 140, 145, 5, 40, 21, 2, 141, 142,
	7, 5, 2, 2, 142, 144, 5, 40, 21, 2, 143, 141, 3, 2, 2, 2, 144, 147, 3,
	2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 148, 3, 2, 2,
	2, 147, 145, 3, 2, 2, 2, 148, 149, 7, 28, 2, 2, 149, 152, 5, 32, 17, 2,
	150, 151, 7, 65, 2, 2, 151, 153, 5, 46, 24, 2, 152, 150, 3, 2, 2, 2, 152,
	153, 3, 2, 2, 2, 153, 157, 3, 2, 2, 2, 154, 156, 5, 38, 20, 2, 155, 154,
	3, 2, 2, 2, 156, 159, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2,
	2, 2, 158, 160, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 160, 161, 7, 73, 2, 2,
	161, 19, 3, 2, 2, 2, 162, 163, 7, 26, 2, 2, 163, 164, 7, 59, 2, 2, 164,
	165, 5, 6, 4, 2, 165, 21, 3, 2, 2, 2, 166, 167, 7, 33, 2, 2, 167, 168,
	7, 34, 2, 2, 168, 169, 5, 8, 5, 2, 169, 170, 7, 72, 2, 2, 170, 175, 5,
	50, 26, 2, 171, 172, 7, 5, 2, 2, 172, 174, 5, 50, 26, 2, 173, 171, 3, 2,
	2, 2, 174, 177, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2,
	176, 178, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 178, 179, 7, 73, 2, 2, 179,
	180, 3, 2, 2, 2, 180, 181, 7, 63, 2, 2, 181, 182, 7, 72, 2, 2, 182, 187,
	5, 66, 34, 2, 183, 184, 7, 5, 2, 2, 184, 186, 5, 66, 34, 2, 185, 183, 3,
	2, 2, 2, 186, 189, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2,
	2, 188, 190, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 191, 7, 73, 2, 2, 191,
	23, 3, 2, 2, 2, 192, 193, 7, 62, 2, 2, 193, 194, 5, 8, 5, 2, 194, 195,
	7, 56, 2, 2, 195, 200, 5, 26, 14, 2, 196, 197, 7, 5, 2, 2, 197, 199, 5,
	26, 14, 2, 198, 196, 3, 2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2,
	2, 2, 200, 201, 3, 2, 2, 2, 201, 203, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2,
	203, 204, 7, 65, 2, 2, 204, 205, 5, 30, 16, 2, 205, 25, 3, 2, 2, 2, 206,
	207, 5, 50, 26, 2, 207, 208, 7, 70, 2, 2, 208, 209, 5, 66, 34, 2, 209,
	27, 3, 2, 2, 2, 210, 211, 7, 21, 2, 2, 211, 212, 7, 28, 2, 2, 212, 218,
	5, 8, 5, 2, 213, 214, 7, 65, 2, 2, 214, 215, 5, 50, 26, 2, 215, 216, 7,
	70, 2, 2, 216, 217, 5, 66, 34, 2, 217, 219, 3, 2, 2, 2, 218, 213, 3, 2,
	2, 2, 218, 219, 3, 2, 2, 2, 219, 29, 3, 2, 2, 2, 220, 221, 5, 70, 36, 2,
	221, 222, 7, 70, 2, 2, 222, 223, 5, 86, 44, 2, 223, 31, 3, 2, 2, 2, 224,
	227, 5, 74, 38, 2, 225, 226, 7, 10, 2, 2, 226, 228, 5, 82, 42, 2, 227,
	225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 33, 3, 2, 2, 2, 229, 230, 7,
	53, 2, 2, 230, 235, 5, 40, 21, 2, 231, 232, 7, 5, 2, 2, 232, 234, 5, 40,
	21, 2, 233, 231, 3, 2, 2, 2, 234, 237, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2,
	235, 236, 3, 2, 2, 2, 236, 238, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 238,
	239, 7, 28, 2, 2, 239, 243, 5, 32, 17, 2, 240, 242, 5, 38, 20, 2, 241,
	240, 3, 2, 2, 2, 242, 245, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 244,
	3, 2, 2, 2, 244, 248, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 246, 247, 7, 65,
	2, 2, 247, 249, 5, 46, 24, 2, 248, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2,
	2, 249, 260, 3, 2, 2, 2, 250, 251, 7, 29, 2, 2, 251, 252, 7, 13, 2, 2,
	252, 257, 5, 42, 22, 2, 253, 254, 7, 5, 2, 2, 254, 256, 5, 42, 22, 2, 255,
	253, 3, 2, 2, 2, 256, 259, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 257, 258,
	3, 2, 2, 2, 258, 261, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 260, 250, 3, 2,
	2, 2, 260, 261, 3, 2, 2, 2, 261, 272, 3, 2, 2, 2, 262, 263, 7, 48, 2, 2,
	263, 264, 7, 13, 2, 2, 264, 269, 5, 36, 19, 2, 265, 266, 7, 5, 2, 2, 266,
	268, 5, 36, 19, 2, 267, 265, 3, 2, 2, 2, 268, 271, 3, 2, 2, 2, 269, 267,
	3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 273, 3, 2, 2, 2, 271, 269, 3, 2,
	2, 2, 272, 262, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 280, 3, 2, 2, 2,
	274, 275, 7, 46, 2, 2, 275, 278, 5, 84, 43, 2, 276, 277, 7, 39, 2, 2, 277,
	279, 5, 84, 43, 2, 278, 276, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 281,
	3, 2, 2, 2, 280, 274, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 35, 3, 2,
	2, 2, 282, 291, 5, 42, 22, 2, 283, 284, 7, 52, 2, 2, 284, 285, 7, 72, 2,
	2, 285, 286, 5, 42, 22, 2, 286, 287, 7, 5, 2, 2, 287, 288, 5, 60, 31, 2,
	288, 289, 7, 73, 2, 2, 289, 291, 3, 2, 2, 2, 290, 282, 3, 2, 2, 2, 290,
	283, 3, 2, 2, 2, 291, 293, 3, 2, 2, 2, 292, 294, 9, 2, 2, 2, 293, 292,
	3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294, 37, 3, 2, 2, 2, 295, 296, 7, 37,
	2, 2, 296, 297, 7, 49, 2, 2, 297, 298, 7, 35, 2, 2, 298, 299, 5, 32, 17,
	2, 299, 300, 7, 47, 2, 2, 300, 301, 5, 42, 22, 2, 301, 302, 7, 70, 2, 2,
	302, 303, 5, 42, 22, 2, 303, 39, 3, 2, 2, 2, 304, 312, 7, 4, 2, 2, 305,
	306, 5, 74, 38, 2, 306, 307, 7, 6, 2, 2, 307, 308, 7, 4, 2, 2, 308, 312,
	3, 2, 2, 2, 309, 312, 5, 42, 22, 2, 310, 312, 5, 44, 23, 2, 311, 304, 3,
	2, 2, 2, 311, 305, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 311, 310, 3, 2, 2,
	2, 312, 41, 3, 2, 2, 2, 313, 314, 5, 74, 38, 2, 314, 315, 7, 6, 2, 2, 315,
	317, 3, 2, 2, 2, 316, 313, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 318,
	3, 2, 2, 2, 318, 319, 5, 76, 39, 2, 319, 43, 3, 2, 2, 2, 320, 321, 9, 3,
	2, 2, 321, 322, 7, 72, 2, 2, 322, 323, 5, 42, 22, 2, 323, 324, 7, 73, 2,
	2, 324, 45, 3, 2, 2, 2, 325, 326, 5, 48, 25, 2, 326, 327, 5, 64, 33, 2,
	327, 331, 5, 66, 34, 2, 328, 329, 5, 52, 27, 2, 329, 330, 5, 46, 24, 2,
	330, 332, 3, 2, 2, 2, 331, 328, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332,
	397, 3, 2, 2, 2, 333, 350, 5, 48, 25, 2, 334, 351, 7, 38, 2, 2, 335, 351,
	7, 40, 2, 2, 336, 337, 7, 58, 2, 2, 337, 351, 7, 66, 2, 2, 338, 339, 7,
	27, 2, 2, 339, 351, 7, 66, 2, 2, 340, 351, 7, 18, 2, 2, 341, 342, 7, 18,
	2, 2, 342, 343, 7, 31, 2, 2, 343, 351, 7, 15, 2, 2, 344, 345, 7, 24, 2,
	2, 345, 346, 7, 43, 2, 2, 346, 351, 7, 25, 2, 2, 347, 348, 7, 43, 2, 2,
	348, 349, 7, 44, 2, 2, 349, 351, 7, 61, 2, 2, 350, 334, 3, 2, 2, 2, 350,
	335, 3, 2, 2, 2, 350, 336, 3, 2, 2, 2, 350, 338, 3, 2, 2, 2, 350, 340,
	3, 2, 2, 2, 350, 341, 3, 2, 2, 2, 350, 344, 3, 2, 2, 2, 350, 347, 3, 2,
	2, 2, 351, 352, 3, 2, 2, 2, 352, 356, 5, 86, 44, 2, 353, 354, 5, 52, 27,
	2, 354, 355, 5, 46, 24, 2, 355, 357, 3, 2, 2, 2, 356, 353, 3, 2, 2, 2,
	356, 357, 3, 2, 2, 2, 357, 397, 3, 2, 2, 2, 358, 359, 5, 48, 25, 2, 359,
	360, 7, 32, 2, 2, 360, 361, 7, 72, 2, 2, 361, 366, 5, 86, 44, 2, 362, 363,
	7, 5, 2, 2, 363, 365, 5, 86, 44, 2, 364, 362, 3, 2, 2, 2, 365, 368, 3,
	2, 2, 2, 366, 364, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 369, 3, 2, 2,
	2, 368, 366, 3, 2, 2, 2, 369, 373, 7, 73, 2, 2, 370, 371, 5, 52, 27, 2,
	371, 372, 5, 46, 24, 2, 372, 374, 3, 2, 2, 2, 373, 370, 3, 2, 2, 2, 373,
	374, 3, 2, 2, 2, 374, 397, 3, 2, 2, 2, 375, 376, 5, 48, 25, 2, 376, 377,
	7, 14, 2, 2, 377, 378, 5, 66, 34, 2, 378, 379, 7, 8, 2, 2, 379, 383, 5,
	66, 34, 2, 380, 381, 5, 52, 27, 2, 381, 382, 5, 46, 24, 2, 382, 384, 3,
	2, 2, 2, 383, 380, 3, 2, 2, 2, 383, 384, 3, 2, 2, 2, 384, 397, 3, 2, 2,
	2, 385, 386, 7, 54, 2, 2, 386, 387, 7, 72, 2, 2, 387, 388, 5, 42, 22, 2,
	388, 389, 7, 5, 2, 2, 389, 390, 5, 54, 28, 2, 390, 394, 7, 73, 2, 2, 391,
	392, 5, 52, 27, 2, 392, 393, 5, 46, 24, 2, 393, 395, 3, 2, 2, 2, 394, 391,
	3, 2, 2, 2, 394, 395, 3, 2, 2, 2, 395, 397, 3, 2, 2, 2, 396, 325, 3, 2,
	2, 2, 396, 333, 3, 2, 2, 2, 396, 358, 3, 2, 2, 2, 396, 375, 3, 2, 2, 2,
	396, 385, 3, 2, 2, 2, 397, 47, 3, 2, 2, 2, 398, 399, 5, 42, 22, 2, 399,
	49, 3, 2, 2, 2, 400, 401, 5, 76, 39, 2, 401, 51, 3, 2, 2, 2, 402, 403,
	9, 4, 2, 2, 403, 53, 3, 2, 2, 2, 404, 407, 5, 56, 29, 2, 405, 407, 5, 58,
	30, 2, 406, 404, 3, 2, 2, 2, 406, 405, 3, 2, 2, 2, 407, 55, 3, 2, 2, 2,
	408, 409, 7, 16, 2, 2, 409, 410, 7, 72, 2, 2, 410, 411, 5, 60, 31, 2, 411,
	412, 7, 5, 2, 2, 412, 413, 5, 84, 43, 2, 413, 414, 7, 73, 2, 2, 414, 57,
	3, 2, 2, 2, 415, 416, 7, 50, 2, 2, 416, 417, 7, 72, 2, 2, 417, 418, 5,
	60, 31, 2, 418, 419, 7, 5, 2, 2, 419, 420, 5, 60, 31, 2, 420, 421, 7, 73,
	2, 2, 421, 59, 3, 2, 2, 2, 422, 423, 7, 36, 2, 2, 423, 424, 7, 72, 2, 2,
	424, 425, 5, 84, 43, 2, 425, 426, 7, 5, 2, 2, 426, 427, 5, 84, 43, 2, 427,
	428, 7, 73, 2, 2, 428, 61, 3, 2, 2, 2, 429, 488, 3, 2, 2, 2, 430, 488,
	7, 7, 2, 2, 431, 488, 7, 8, 2, 2, 432, 488, 7, 9, 2, 2, 433, 488, 7, 10,
	2, 2, 434, 488, 7, 11, 2, 2, 435, 488, 7, 12, 2, 2, 436, 488, 7, 13, 2,
	2, 437, 488, 7, 14, 2, 2, 438, 488, 7, 15, 2, 2, 439, 488, 7, 16, 2, 2,
	440, 488, 7, 17, 2, 2, 441, 488, 7, 25, 2, 2, 442, 488, 7, 18, 2, 2, 443,
	488, 7, 19, 2, 2, 444, 488, 7, 20, 2, 2, 445, 488, 7, 21, 2, 2, 446, 488,
	7, 22, 2, 2, 447, 488, 7, 24, 2, 2, 448, 488, 7, 26, 2, 2, 449, 488, 7,
	27, 2, 2, 450, 488, 7, 44, 2, 2, 451, 488, 7, 28, 2, 2, 452, 488, 7, 29,
	2, 2, 453, 488, 7, 30, 2, 2, 454, 488, 7, 31, 2, 2, 455, 488, 7, 32, 2,
	2, 456, 488, 7, 33, 2, 2, 457, 488, 7, 34, 2, 2, 458, 488, 7, 35, 2, 2,
	459, 488, 7, 36, 2, 2, 460, 488, 7, 37, 2, 2, 461, 488, 7, 38, 2, 2, 462,
	488, 7, 39, 2, 2, 463, 488, 7, 40, 2, 2, 464, 488, 7, 41, 2, 2, 465, 488,
	7, 42, 2, 2, 466, 488, 7, 43, 2, 2, 467, 488, 7, 45, 2, 2, 468, 488, 7,
	46, 2, 2, 469, 488, 7, 47, 2, 2, 470, 488, 7, 48, 2, 2, 471, 488, 7, 49,
	2, 2, 472, 488, 7, 50, 2, 2, 473, 488, 7, 51, 2, 2, 474, 488, 7, 53, 2,
	2, 475, 488, 7, 56, 2, 2, 476, 488, 7, 58, 2, 2, 477, 488, 7, 52, 2, 2,
	478, 488, 7, 54, 2, 2, 479, 488, 7, 55, 2, 2, 480, 488, 7, 59, 2, 2, 481,
	488, 7, 61, 2, 2, 482, 488, 7, 62, 2, 2, 483, 488, 7, 63, 2, 2, 484, 488,
	7, 64, 2, 2, 485, 488, 7, 65, 2, 2, 486, 488, 7, 66, 2, 2, 487, 429, 3,
	2, 2, 2, 487, 430, 3, 2, 2, 2, 487, 431, 3, 2, 2, 2, 487, 432, 3, 2, 2,
	2, 487, 433, 3, 2, 2, 2, 487, 434, 3, 2, 2, 2, 487, 435, 3, 2, 2, 2, 487,
	436, 3, 2, 2, 2, 487, 437, 3, 2, 2, 2, 487, 438, 3, 2, 2, 2, 487, 439,
	3, 2, 2, 2, 487, 440, 3, 2, 2, 2, 487, 441, 3, 2, 2, 2, 487, 442, 3, 2,
	2, 2, 487, 443, 3, 2, 2, 2, 487, 444, 3, 2, 2, 2, 487, 445, 3, 2, 2, 2,
	487, 446, 3, 2, 2, 2, 487, 447, 3, 2, 2, 2, 487, 448, 3, 2, 2, 2, 487,
	449, 3, 2, 2, 2, 487, 450, 3, 2, 2, 2, 487, 451, 3, 2, 2, 2, 487, 452,
	3, 2, 2, 2, 487, 453, 3, 2, 2, 2, 487, 454, 3, 2, 2, 2, 487, 455, 3, 2,
	2, 2, 487, 456, 3, 2, 2, 2, 487, 457, 3, 2, 2, 2, 487, 458, 3, 2, 2, 2,
	487, 459, 3, 2, 2, 2, 487, 460, 3, 2, 2, 2, 487, 461, 3, 2, 2, 2, 487,
	462, 3, 2, 2, 2, 487, 463, 3, 2, 2, 2, 487, 464, 3, 2, 2, 2, 487, 465,
	3, 2, 2, 2, 487, 466, 3, 2, 2, 2, 487, 467, 3, 2, 2, 2, 487, 468, 3, 2,
	2, 2, 487, 469, 3, 2, 2, 2, 487, 470, 3, 2, 2, 2, 487, 471, 3, 2, 2, 2,
	487, 472, 3, 2, 2, 2, 487, 473, 3, 2, 2, 2, 487, 474, 3, 2, 2, 2, 487,
	475, 3, 2, 2, 2, 487, 476, 3, 2, 2, 2, 487, 477, 3, 2, 2, 2, 487, 478,
	3, 2, 2, 2, 487, 479, 3, 2, 2, 2, 487, 480, 3, 2, 2, 2, 487, 481, 3, 2,
	2, 2, 487, 482, 3, 2, 2, 2, 487, 483, 3, 2, 2, 2, 487, 484, 3, 2, 2, 2,
	487, 485, 3, 2, 2, 2, 487, 486, 3, 2, 2, 2, 488, 63, 3, 2, 2, 2, 489, 490,
	9, 5, 2, 2, 490, 65, 3, 2, 2, 2, 491, 494, 5, 84, 43, 2, 492, 494, 5, 86,
	44, 2, 493, 491, 3, 2, 2, 2, 493, 492, 3, 2, 2, 2, 494, 67, 3, 2, 2, 2,
	495, 496, 5, 86, 44, 2, 496, 69, 3, 2, 2, 2, 497, 498, 5, 86, 44, 2, 498,
	71, 3, 2, 2, 2, 499, 500, 5, 70, 36, 2, 500, 73, 3, 2, 2, 2, 501, 502,
	5, 70, 36, 2, 502, 75, 3, 2, 2, 2, 503, 504, 5, 70, 36, 2, 504, 77, 3,
	2, 2, 2, 505, 506, 5, 74, 38, 2, 506, 79, 3, 2, 2, 2, 507, 508, 5, 70,
	36, 2, 508, 81, 3, 2, 2, 2, 509, 510, 5, 70, 36, 2, 510, 83, 3, 2, 2, 2,
	511, 512, 7, 74, 2, 2, 512, 85, 3, 2, 2, 2, 513, 514, 7, 75, 2, 2, 514,
	87, 3, 2, 2, 2, 36, 91, 104, 145, 152, 157, 175, 187, 200, 218, 227, 235,
	243, 248, 257, 260, 269, 272, 278, 280, 290, 293, 311, 316, 331, 350, 356,
	366, 373, 383, 394, 396, 406, 487, 493,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'*'", "','", "'.'", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "'<='", "'>='",
	"'>'", "'='", "'<'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "K_ALTER", "K_AND", "K_OR", "K_AS", "K_ASC", "K_AVERAGE",
	"K_BY", "K_BETWEEN", "K_CASE", "K_CIRCLE", "K_COLUMN", "K_CONTAINS", "K_COUNT",
	"K_CREATE", "K_DELETE", "K_DESC", "K_DESCRIBE", "K_DOES", "K_CONTAIN",
	"K_DROP", "K_ENDS", "K_FROM", "K_GROUP", "K_HAVING", "K_IGNORING", "K_IN",
	"K_INSERT", "K_INTO", "K_JOIN", "K_LATLNG", "K_LEFT", "K_LIKE", "K_LIMIT",
	"K_MATCHES", "K_MAXIMUM", "K_MINIMUM", "K_NOT", "K_EQUAL", "K_OF", "K_OFFSET",
	"K_ON", "K_ORDER", "K_OUTER", "K_RECTANGLE", "K_RENAME", "K_ST_DISTANCE",
	"K_SELECT", "K_ST_INTERSECTS", "K_SUM", "K_SET", "K_SHOW", "K_STARTS",
	"K_TABLE", "K_TABLES", "K_TO", "K_UPDATE", "K_VALUES", "K_VIEW", "K_WHERE",
	"K_WITH", "LT_EQ", "GT_EQ", "GT", "EQ", "LT", "LPAR", "RPAR", "NUMERIC_LITERAL",
	"STRING_LITERAL", "STRING", "QUOTED_STRING", "SINGLELINE_COMMENT", "MULTILINE_COMMENT",
	"WHITESPACE",
}

var ruleNames = []string{
	"fusionTablesSql", "sql_stmt", "table_name_in_ddl", "table_name_in_dml",
	"create_table_as_select_stmt", "describe_stmt", "show_tables_stmt", "alter_table_stmt",
	"create_view_stmt", "drop_table_stmt", "insert_stmt", "update_stmt", "column_assignment",
	"delete_stmt", "eq_comparison", "table_name_with_alias", "select_stmt",
	"ordering_term", "join_clause", "result_column", "qualified_column_name",
	"aggregate_exp", "expr", "column_name_beginning_expr", "column_name_in_dml",
	"and_or_or", "geometry", "circle", "rectangle", "coordinate", "keyword",
	"operator", "literal", "error_message", "identifier", "column_alias", "table_name",
	"column_name", "new_table_name", "view_name", "table_alias", "numeric_literal",
	"string_literal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FusionTablesSqlParser struct {
	*antlr.BaseParser
}

func NewFusionTablesSqlParser(input antlr.TokenStream) *FusionTablesSqlParser {
	this := new(FusionTablesSqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FusionTablesSql.g4"

	return this
}

// FusionTablesSqlParser tokens.
const (
	FusionTablesSqlParserEOF                = antlr.TokenEOF
	FusionTablesSqlParserT__0               = 1
	FusionTablesSqlParserT__1               = 2
	FusionTablesSqlParserT__2               = 3
	FusionTablesSqlParserT__3               = 4
	FusionTablesSqlParserK_ALTER            = 5
	FusionTablesSqlParserK_AND              = 6
	FusionTablesSqlParserK_OR               = 7
	FusionTablesSqlParserK_AS               = 8
	FusionTablesSqlParserK_ASC              = 9
	FusionTablesSqlParserK_AVERAGE          = 10
	FusionTablesSqlParserK_BY               = 11
	FusionTablesSqlParserK_BETWEEN          = 12
	FusionTablesSqlParserK_CASE             = 13
	FusionTablesSqlParserK_CIRCLE           = 14
	FusionTablesSqlParserK_COLUMN           = 15
	FusionTablesSqlParserK_CONTAINS         = 16
	FusionTablesSqlParserK_COUNT            = 17
	FusionTablesSqlParserK_CREATE           = 18
	FusionTablesSqlParserK_DELETE           = 19
	FusionTablesSqlParserK_DESC             = 20
	FusionTablesSqlParserK_DESCRIBE         = 21
	FusionTablesSqlParserK_DOES             = 22
	FusionTablesSqlParserK_CONTAIN          = 23
	FusionTablesSqlParserK_DROP             = 24
	FusionTablesSqlParserK_ENDS             = 25
	FusionTablesSqlParserK_FROM             = 26
	FusionTablesSqlParserK_GROUP            = 27
	FusionTablesSqlParserK_HAVING           = 28
	FusionTablesSqlParserK_IGNORING         = 29
	FusionTablesSqlParserK_IN               = 30
	FusionTablesSqlParserK_INSERT           = 31
	FusionTablesSqlParserK_INTO             = 32
	FusionTablesSqlParserK_JOIN             = 33
	FusionTablesSqlParserK_LATLNG           = 34
	FusionTablesSqlParserK_LEFT             = 35
	FusionTablesSqlParserK_LIKE             = 36
	FusionTablesSqlParserK_LIMIT            = 37
	FusionTablesSqlParserK_MATCHES          = 38
	FusionTablesSqlParserK_MAXIMUM          = 39
	FusionTablesSqlParserK_MINIMUM          = 40
	FusionTablesSqlParserK_NOT              = 41
	FusionTablesSqlParserK_EQUAL            = 42
	FusionTablesSqlParserK_OF               = 43
	FusionTablesSqlParserK_OFFSET           = 44
	FusionTablesSqlParserK_ON               = 45
	FusionTablesSqlParserK_ORDER            = 46
	FusionTablesSqlParserK_OUTER            = 47
	FusionTablesSqlParserK_RECTANGLE        = 48
	FusionTablesSqlParserK_RENAME           = 49
	FusionTablesSqlParserK_ST_DISTANCE      = 50
	FusionTablesSqlParserK_SELECT           = 51
	FusionTablesSqlParserK_ST_INTERSECTS    = 52
	FusionTablesSqlParserK_SUM              = 53
	FusionTablesSqlParserK_SET              = 54
	FusionTablesSqlParserK_SHOW             = 55
	FusionTablesSqlParserK_STARTS           = 56
	FusionTablesSqlParserK_TABLE            = 57
	FusionTablesSqlParserK_TABLES           = 58
	FusionTablesSqlParserK_TO               = 59
	FusionTablesSqlParserK_UPDATE           = 60
	FusionTablesSqlParserK_VALUES           = 61
	FusionTablesSqlParserK_VIEW             = 62
	FusionTablesSqlParserK_WHERE            = 63
	FusionTablesSqlParserK_WITH             = 64
	FusionTablesSqlParserLT_EQ              = 65
	FusionTablesSqlParserGT_EQ              = 66
	FusionTablesSqlParserGT                 = 67
	FusionTablesSqlParserEQ                 = 68
	FusionTablesSqlParserLT                 = 69
	FusionTablesSqlParserLPAR               = 70
	FusionTablesSqlParserRPAR               = 71
	FusionTablesSqlParserNUMERIC_LITERAL    = 72
	FusionTablesSqlParserSTRING_LITERAL     = 73
	FusionTablesSqlParserSTRING             = 74
	FusionTablesSqlParserQUOTED_STRING      = 75
	FusionTablesSqlParserSINGLELINE_COMMENT = 76
	FusionTablesSqlParserMULTILINE_COMMENT  = 77
	FusionTablesSqlParserWHITESPACE         = 78
)

// FusionTablesSqlParser rules.
const (
	FusionTablesSqlParserRULE_fusionTablesSql             = 0
	FusionTablesSqlParserRULE_sql_stmt                    = 1
	FusionTablesSqlParserRULE_table_name_in_ddl           = 2
	FusionTablesSqlParserRULE_table_name_in_dml           = 3
	FusionTablesSqlParserRULE_create_table_as_select_stmt = 4
	FusionTablesSqlParserRULE_describe_stmt               = 5
	FusionTablesSqlParserRULE_show_tables_stmt            = 6
	FusionTablesSqlParserRULE_alter_table_stmt            = 7
	FusionTablesSqlParserRULE_create_view_stmt            = 8
	FusionTablesSqlParserRULE_drop_table_stmt             = 9
	FusionTablesSqlParserRULE_insert_stmt                 = 10
	FusionTablesSqlParserRULE_update_stmt                 = 11
	FusionTablesSqlParserRULE_column_assignment           = 12
	FusionTablesSqlParserRULE_delete_stmt                 = 13
	FusionTablesSqlParserRULE_eq_comparison               = 14
	FusionTablesSqlParserRULE_table_name_with_alias       = 15
	FusionTablesSqlParserRULE_select_stmt                 = 16
	FusionTablesSqlParserRULE_ordering_term               = 17
	FusionTablesSqlParserRULE_join_clause                 = 18
	FusionTablesSqlParserRULE_result_column               = 19
	FusionTablesSqlParserRULE_qualified_column_name       = 20
	FusionTablesSqlParserRULE_aggregate_exp               = 21
	FusionTablesSqlParserRULE_expr                        = 22
	FusionTablesSqlParserRULE_column_name_beginning_expr  = 23
	FusionTablesSqlParserRULE_column_name_in_dml          = 24
	FusionTablesSqlParserRULE_and_or_or                   = 25
	FusionTablesSqlParserRULE_geometry                    = 26
	FusionTablesSqlParserRULE_circle                      = 27
	FusionTablesSqlParserRULE_rectangle                   = 28
	FusionTablesSqlParserRULE_coordinate                  = 29
	FusionTablesSqlParserRULE_keyword                     = 30
	FusionTablesSqlParserRULE_operator                    = 31
	FusionTablesSqlParserRULE_literal                     = 32
	FusionTablesSqlParserRULE_error_message               = 33
	FusionTablesSqlParserRULE_identifier                  = 34
	FusionTablesSqlParserRULE_column_alias                = 35
	FusionTablesSqlParserRULE_table_name                  = 36
	FusionTablesSqlParserRULE_column_name                 = 37
	FusionTablesSqlParserRULE_new_table_name              = 38
	FusionTablesSqlParserRULE_view_name                   = 39
	FusionTablesSqlParserRULE_table_alias                 = 40
	FusionTablesSqlParserRULE_numeric_literal             = 41
	FusionTablesSqlParserRULE_string_literal              = 42
)

// IFusionTablesSqlContext is an interface to support dynamic dispatch.
type IFusionTablesSqlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFusionTablesSqlContext differentiates from other interfaces.
	IsFusionTablesSqlContext()
}

type FusionTablesSqlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFusionTablesSqlContext() *FusionTablesSqlContext {
	var p = new(FusionTablesSqlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_fusionTablesSql
	return p
}

func (*FusionTablesSqlContext) IsFusionTablesSqlContext() {}

func NewFusionTablesSqlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FusionTablesSqlContext {
	var p = new(FusionTablesSqlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_fusionTablesSql

	return p
}

func (s *FusionTablesSqlContext) GetParser() antlr.Parser { return s.parser }

func (s *FusionTablesSqlContext) AllSql_stmt() []ISql_stmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISql_stmtContext)(nil)).Elem())
	var tst = make([]ISql_stmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISql_stmtContext)
		}
	}

	return tst
}

func (s *FusionTablesSqlContext) Sql_stmt(i int) ISql_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISql_stmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISql_stmtContext)
}

func (s *FusionTablesSqlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FusionTablesSqlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FusionTablesSqlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterFusionTablesSql(s)
	}
}

func (s *FusionTablesSqlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitFusionTablesSql(s)
	}
}

func (p *FusionTablesSqlParser) FusionTablesSql() (localctx IFusionTablesSqlContext) {
	localctx = NewFusionTablesSqlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FusionTablesSqlParserRULE_fusionTablesSql)
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
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FusionTablesSqlParserK_ALTER)|(1<<FusionTablesSqlParserK_CREATE)|(1<<FusionTablesSqlParserK_DELETE)|(1<<FusionTablesSqlParserK_DESCRIBE)|(1<<FusionTablesSqlParserK_DROP)|(1<<FusionTablesSqlParserK_INSERT))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(FusionTablesSqlParserK_SELECT-51))|(1<<(FusionTablesSqlParserK_SHOW-51))|(1<<(FusionTablesSqlParserK_UPDATE-51)))) != 0) {
		{
			p.SetState(86)
			p.Sql_stmt()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISql_stmtContext is an interface to support dynamic dispatch.
type ISql_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSql_stmtContext differentiates from other interfaces.
	IsSql_stmtContext()
}

type Sql_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySql_stmtContext() *Sql_stmtContext {
	var p = new(Sql_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_sql_stmt
	return p
}

func (*Sql_stmtContext) IsSql_stmtContext() {}

func NewSql_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sql_stmtContext {
	var p = new(Sql_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_sql_stmt

	return p
}

func (s *Sql_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Sql_stmtContext) Alter_table_stmt() IAlter_table_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlter_table_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlter_table_stmtContext)
}

func (s *Sql_stmtContext) Select_stmt() ISelect_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelect_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelect_stmtContext)
}

func (s *Sql_stmtContext) Create_view_stmt() ICreate_view_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreate_view_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreate_view_stmtContext)
}

func (s *Sql_stmtContext) Create_table_as_select_stmt() ICreate_table_as_select_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreate_table_as_select_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreate_table_as_select_stmtContext)
}

func (s *Sql_stmtContext) Delete_stmt() IDelete_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDelete_stmtContext)
}

func (s *Sql_stmtContext) Drop_table_stmt() IDrop_table_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDrop_table_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDrop_table_stmtContext)
}

func (s *Sql_stmtContext) Insert_stmt() IInsert_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsert_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsert_stmtContext)
}

func (s *Sql_stmtContext) Update_stmt() IUpdate_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdate_stmtContext)
}

func (s *Sql_stmtContext) Describe_stmt() IDescribe_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescribe_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescribe_stmtContext)
}

func (s *Sql_stmtContext) Show_tables_stmt() IShow_tables_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShow_tables_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShow_tables_stmtContext)
}

func (s *Sql_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sql_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sql_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterSql_stmt(s)
	}
}

func (s *Sql_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitSql_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Sql_stmt() (localctx ISql_stmtContext) {
	localctx = NewSql_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FusionTablesSqlParserRULE_sql_stmt)

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
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(92)
			p.Alter_table_stmt()
		}

	case 2:
		{
			p.SetState(93)
			p.Select_stmt()
		}

	case 3:
		{
			p.SetState(94)
			p.Create_view_stmt()
		}

	case 4:
		{
			p.SetState(95)
			p.Create_table_as_select_stmt()
		}

	case 5:
		{
			p.SetState(96)
			p.Delete_stmt()
		}

	case 6:
		{
			p.SetState(97)
			p.Drop_table_stmt()
		}

	case 7:
		{
			p.SetState(98)
			p.Insert_stmt()
		}

	case 8:
		{
			p.SetState(99)
			p.Update_stmt()
		}

	case 9:
		{
			p.SetState(100)
			p.Describe_stmt()
		}

	case 10:
		{
			p.SetState(101)
			p.Show_tables_stmt()
		}

	}
	{
		p.SetState(104)
		p.Match(FusionTablesSqlParserT__0)
	}

	return localctx
}

// ITable_name_in_ddlContext is an interface to support dynamic dispatch.
type ITable_name_in_ddlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_name_in_ddlContext differentiates from other interfaces.
	IsTable_name_in_ddlContext()
}

type Table_name_in_ddlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_name_in_ddlContext() *Table_name_in_ddlContext {
	var p = new(Table_name_in_ddlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_table_name_in_ddl
	return p
}

func (*Table_name_in_ddlContext) IsTable_name_in_ddlContext() {}

func NewTable_name_in_ddlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_name_in_ddlContext {
	var p = new(Table_name_in_ddlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_table_name_in_ddl

	return p
}

func (s *Table_name_in_ddlContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_name_in_ddlContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Table_name_in_ddlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_name_in_ddlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_name_in_ddlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterTable_name_in_ddl(s)
	}
}

func (s *Table_name_in_ddlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitTable_name_in_ddl(s)
	}
}

func (p *FusionTablesSqlParser) Table_name_in_ddl() (localctx ITable_name_in_ddlContext) {
	localctx = NewTable_name_in_ddlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FusionTablesSqlParserRULE_table_name_in_ddl)

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
		p.Table_name()
	}

	return localctx
}

// ITable_name_in_dmlContext is an interface to support dynamic dispatch.
type ITable_name_in_dmlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_name_in_dmlContext differentiates from other interfaces.
	IsTable_name_in_dmlContext()
}

type Table_name_in_dmlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_name_in_dmlContext() *Table_name_in_dmlContext {
	var p = new(Table_name_in_dmlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_table_name_in_dml
	return p
}

func (*Table_name_in_dmlContext) IsTable_name_in_dmlContext() {}

func NewTable_name_in_dmlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_name_in_dmlContext {
	var p = new(Table_name_in_dmlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_table_name_in_dml

	return p
}

func (s *Table_name_in_dmlContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_name_in_dmlContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Table_name_in_dmlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_name_in_dmlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_name_in_dmlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterTable_name_in_dml(s)
	}
}

func (s *Table_name_in_dmlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitTable_name_in_dml(s)
	}
}

func (p *FusionTablesSqlParser) Table_name_in_dml() (localctx ITable_name_in_dmlContext) {
	localctx = NewTable_name_in_dmlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FusionTablesSqlParserRULE_table_name_in_dml)

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
		p.Table_name()
	}

	return localctx
}

// ICreate_table_as_select_stmtContext is an interface to support dynamic dispatch.
type ICreate_table_as_select_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreate_table_as_select_stmtContext differentiates from other interfaces.
	IsCreate_table_as_select_stmtContext()
}

type Create_table_as_select_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_table_as_select_stmtContext() *Create_table_as_select_stmtContext {
	var p = new(Create_table_as_select_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_create_table_as_select_stmt
	return p
}

func (*Create_table_as_select_stmtContext) IsCreate_table_as_select_stmtContext() {}

func NewCreate_table_as_select_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_table_as_select_stmtContext {
	var p = new(Create_table_as_select_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_create_table_as_select_stmt

	return p
}

func (s *Create_table_as_select_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_table_as_select_stmtContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CREATE, 0)
}

func (s *Create_table_as_select_stmtContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_TABLE, 0)
}

func (s *Create_table_as_select_stmtContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Create_table_as_select_stmtContext) K_AS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_AS, 0)
}

func (s *Create_table_as_select_stmtContext) K_SELECT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_SELECT, 0)
}

func (s *Create_table_as_select_stmtContext) K_FROM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_FROM, 0)
}

func (s *Create_table_as_select_stmtContext) Table_name_in_ddl() ITable_name_in_ddlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_name_in_ddlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_name_in_ddlContext)
}

func (s *Create_table_as_select_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_table_as_select_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_table_as_select_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterCreate_table_as_select_stmt(s)
	}
}

func (s *Create_table_as_select_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitCreate_table_as_select_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Create_table_as_select_stmt() (localctx ICreate_table_as_select_stmtContext) {
	localctx = NewCreate_table_as_select_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FusionTablesSqlParserRULE_create_table_as_select_stmt)

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
		p.Match(FusionTablesSqlParserK_CREATE)
	}
	{
		p.SetState(111)
		p.Match(FusionTablesSqlParserK_TABLE)
	}
	{
		p.SetState(112)
		p.Identifier()
	}
	{
		p.SetState(113)
		p.Match(FusionTablesSqlParserK_AS)
	}
	{
		p.SetState(114)
		p.Match(FusionTablesSqlParserK_SELECT)
	}
	{
		p.SetState(115)
		p.Match(FusionTablesSqlParserT__1)
	}
	{
		p.SetState(116)
		p.Match(FusionTablesSqlParserK_FROM)
	}
	{
		p.SetState(117)
		p.Table_name_in_ddl()
	}

	return localctx
}

// IDescribe_stmtContext is an interface to support dynamic dispatch.
type IDescribe_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescribe_stmtContext differentiates from other interfaces.
	IsDescribe_stmtContext()
}

type Describe_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescribe_stmtContext() *Describe_stmtContext {
	var p = new(Describe_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_describe_stmt
	return p
}

func (*Describe_stmtContext) IsDescribe_stmtContext() {}

func NewDescribe_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Describe_stmtContext {
	var p = new(Describe_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_describe_stmt

	return p
}

func (s *Describe_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Describe_stmtContext) K_DESCRIBE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_DESCRIBE, 0)
}

func (s *Describe_stmtContext) Table_name_in_ddl() ITable_name_in_ddlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_name_in_ddlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_name_in_ddlContext)
}

func (s *Describe_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Describe_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Describe_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterDescribe_stmt(s)
	}
}

func (s *Describe_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitDescribe_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Describe_stmt() (localctx IDescribe_stmtContext) {
	localctx = NewDescribe_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FusionTablesSqlParserRULE_describe_stmt)

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
		p.SetState(119)
		p.Match(FusionTablesSqlParserK_DESCRIBE)
	}
	{
		p.SetState(120)
		p.Table_name_in_ddl()
	}

	return localctx
}

// IShow_tables_stmtContext is an interface to support dynamic dispatch.
type IShow_tables_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShow_tables_stmtContext differentiates from other interfaces.
	IsShow_tables_stmtContext()
}

type Show_tables_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShow_tables_stmtContext() *Show_tables_stmtContext {
	var p = new(Show_tables_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_show_tables_stmt
	return p
}

func (*Show_tables_stmtContext) IsShow_tables_stmtContext() {}

func NewShow_tables_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Show_tables_stmtContext {
	var p = new(Show_tables_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_show_tables_stmt

	return p
}

func (s *Show_tables_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Show_tables_stmtContext) K_SHOW() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_SHOW, 0)
}

func (s *Show_tables_stmtContext) K_TABLES() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_TABLES, 0)
}

func (s *Show_tables_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Show_tables_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Show_tables_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterShow_tables_stmt(s)
	}
}

func (s *Show_tables_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitShow_tables_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Show_tables_stmt() (localctx IShow_tables_stmtContext) {
	localctx = NewShow_tables_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FusionTablesSqlParserRULE_show_tables_stmt)

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
		p.SetState(122)
		p.Match(FusionTablesSqlParserK_SHOW)
	}
	{
		p.SetState(123)
		p.Match(FusionTablesSqlParserK_TABLES)
	}

	return localctx
}

// IAlter_table_stmtContext is an interface to support dynamic dispatch.
type IAlter_table_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlter_table_stmtContext differentiates from other interfaces.
	IsAlter_table_stmtContext()
}

type Alter_table_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlter_table_stmtContext() *Alter_table_stmtContext {
	var p = new(Alter_table_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_alter_table_stmt
	return p
}

func (*Alter_table_stmtContext) IsAlter_table_stmtContext() {}

func NewAlter_table_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Alter_table_stmtContext {
	var p = new(Alter_table_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_alter_table_stmt

	return p
}

func (s *Alter_table_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Alter_table_stmtContext) K_ALTER() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ALTER, 0)
}

func (s *Alter_table_stmtContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_TABLE, 0)
}

func (s *Alter_table_stmtContext) Table_name_in_ddl() ITable_name_in_ddlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_name_in_ddlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_name_in_ddlContext)
}

func (s *Alter_table_stmtContext) K_RENAME() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_RENAME, 0)
}

func (s *Alter_table_stmtContext) K_TO() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_TO, 0)
}

func (s *Alter_table_stmtContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Alter_table_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Alter_table_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Alter_table_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterAlter_table_stmt(s)
	}
}

func (s *Alter_table_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitAlter_table_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Alter_table_stmt() (localctx IAlter_table_stmtContext) {
	localctx = NewAlter_table_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FusionTablesSqlParserRULE_alter_table_stmt)

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
		p.SetState(125)
		p.Match(FusionTablesSqlParserK_ALTER)
	}
	{
		p.SetState(126)
		p.Match(FusionTablesSqlParserK_TABLE)
	}
	{
		p.SetState(127)
		p.Table_name_in_ddl()
	}

	{
		p.SetState(128)
		p.Match(FusionTablesSqlParserK_RENAME)
	}
	{
		p.SetState(129)
		p.Match(FusionTablesSqlParserK_TO)
	}
	{
		p.SetState(130)
		p.Identifier()
	}

	return localctx
}

// ICreate_view_stmtContext is an interface to support dynamic dispatch.
type ICreate_view_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreate_view_stmtContext differentiates from other interfaces.
	IsCreate_view_stmtContext()
}

type Create_view_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_view_stmtContext() *Create_view_stmtContext {
	var p = new(Create_view_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_create_view_stmt
	return p
}

func (*Create_view_stmtContext) IsCreate_view_stmtContext() {}

func NewCreate_view_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_view_stmtContext {
	var p = new(Create_view_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_create_view_stmt

	return p
}

func (s *Create_view_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_view_stmtContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CREATE, 0)
}

func (s *Create_view_stmtContext) K_VIEW() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_VIEW, 0)
}

func (s *Create_view_stmtContext) View_name() IView_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IView_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IView_nameContext)
}

func (s *Create_view_stmtContext) K_AS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_AS, 0)
}

func (s *Create_view_stmtContext) LPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserLPAR, 0)
}

func (s *Create_view_stmtContext) K_SELECT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_SELECT, 0)
}

func (s *Create_view_stmtContext) AllResult_column() []IResult_columnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IResult_columnContext)(nil)).Elem())
	var tst = make([]IResult_columnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IResult_columnContext)
		}
	}

	return tst
}

func (s *Create_view_stmtContext) Result_column(i int) IResult_columnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResult_columnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IResult_columnContext)
}

func (s *Create_view_stmtContext) K_FROM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_FROM, 0)
}

func (s *Create_view_stmtContext) RPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserRPAR, 0)
}

func (s *Create_view_stmtContext) Table_name_with_alias() ITable_name_with_aliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_name_with_aliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_name_with_aliasContext)
}

func (s *Create_view_stmtContext) AllJoin_clause() []IJoin_clauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJoin_clauseContext)(nil)).Elem())
	var tst = make([]IJoin_clauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJoin_clauseContext)
		}
	}

	return tst
}

func (s *Create_view_stmtContext) Join_clause(i int) IJoin_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJoin_clauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJoin_clauseContext)
}

func (s *Create_view_stmtContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_WHERE, 0)
}

func (s *Create_view_stmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Create_view_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_view_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_view_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterCreate_view_stmt(s)
	}
}

func (s *Create_view_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitCreate_view_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Create_view_stmt() (localctx ICreate_view_stmtContext) {
	localctx = NewCreate_view_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FusionTablesSqlParserRULE_create_view_stmt)
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
		p.Match(FusionTablesSqlParserK_CREATE)
	}
	{
		p.SetState(133)
		p.Match(FusionTablesSqlParserK_VIEW)
	}
	{
		p.SetState(134)
		p.View_name()
	}
	{
		p.SetState(135)
		p.Match(FusionTablesSqlParserK_AS)
	}
	{
		p.SetState(136)
		p.Match(FusionTablesSqlParserLPAR)
	}
	{
		p.SetState(137)
		p.Match(FusionTablesSqlParserK_SELECT)
	}
	{
		p.SetState(138)
		p.Result_column()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FusionTablesSqlParserT__2 {
		{
			p.SetState(139)
			p.Match(FusionTablesSqlParserT__2)
		}
		{
			p.SetState(140)
			p.Result_column()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
		p.Match(FusionTablesSqlParserK_FROM)
	}

	{
		p.SetState(147)
		p.Table_name_with_alias()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FusionTablesSqlParserK_WHERE {
		{
			p.SetState(148)
			p.Match(FusionTablesSqlParserK_WHERE)
		}
		{
			p.SetState(149)
			p.Expr()
		}

	}

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FusionTablesSqlParserK_LEFT {
		{
			p.SetState(152)
			p.Join_clause()
		}

		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	{
		p.SetState(158)
		p.Match(FusionTablesSqlParserRPAR)
	}

	return localctx
}

// IDrop_table_stmtContext is an interface to support dynamic dispatch.
type IDrop_table_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDrop_table_stmtContext differentiates from other interfaces.
	IsDrop_table_stmtContext()
}

type Drop_table_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDrop_table_stmtContext() *Drop_table_stmtContext {
	var p = new(Drop_table_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_drop_table_stmt
	return p
}

func (*Drop_table_stmtContext) IsDrop_table_stmtContext() {}

func NewDrop_table_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Drop_table_stmtContext {
	var p = new(Drop_table_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_drop_table_stmt

	return p
}

func (s *Drop_table_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Drop_table_stmtContext) K_DROP() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_DROP, 0)
}

func (s *Drop_table_stmtContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_TABLE, 0)
}

func (s *Drop_table_stmtContext) Table_name_in_ddl() ITable_name_in_ddlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_name_in_ddlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_name_in_ddlContext)
}

func (s *Drop_table_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Drop_table_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Drop_table_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterDrop_table_stmt(s)
	}
}

func (s *Drop_table_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitDrop_table_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Drop_table_stmt() (localctx IDrop_table_stmtContext) {
	localctx = NewDrop_table_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FusionTablesSqlParserRULE_drop_table_stmt)

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
		p.Match(FusionTablesSqlParserK_DROP)
	}
	{
		p.SetState(161)
		p.Match(FusionTablesSqlParserK_TABLE)
	}
	{
		p.SetState(162)
		p.Table_name_in_ddl()
	}

	return localctx
}

// IInsert_stmtContext is an interface to support dynamic dispatch.
type IInsert_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsert_stmtContext differentiates from other interfaces.
	IsInsert_stmtContext()
}

type Insert_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsert_stmtContext() *Insert_stmtContext {
	var p = new(Insert_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_insert_stmt
	return p
}

func (*Insert_stmtContext) IsInsert_stmtContext() {}

func NewInsert_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Insert_stmtContext {
	var p = new(Insert_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_insert_stmt

	return p
}

func (s *Insert_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Insert_stmtContext) K_INSERT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_INSERT, 0)
}

func (s *Insert_stmtContext) K_INTO() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_INTO, 0)
}

func (s *Insert_stmtContext) Table_name_in_dml() ITable_name_in_dmlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_name_in_dmlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_name_in_dmlContext)
}

func (s *Insert_stmtContext) AllLPAR() []antlr.TerminalNode {
	return s.GetTokens(FusionTablesSqlParserLPAR)
}

func (s *Insert_stmtContext) LPAR(i int) antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserLPAR, i)
}

func (s *Insert_stmtContext) AllColumn_name_in_dml() []IColumn_name_in_dmlContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumn_name_in_dmlContext)(nil)).Elem())
	var tst = make([]IColumn_name_in_dmlContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumn_name_in_dmlContext)
		}
	}

	return tst
}

func (s *Insert_stmtContext) Column_name_in_dml(i int) IColumn_name_in_dmlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_name_in_dmlContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumn_name_in_dmlContext)
}

func (s *Insert_stmtContext) AllRPAR() []antlr.TerminalNode {
	return s.GetTokens(FusionTablesSqlParserRPAR)
}

func (s *Insert_stmtContext) RPAR(i int) antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserRPAR, i)
}

func (s *Insert_stmtContext) K_VALUES() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_VALUES, 0)
}

func (s *Insert_stmtContext) AllLiteral() []ILiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralContext)(nil)).Elem())
	var tst = make([]ILiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralContext)
		}
	}

	return tst
}

func (s *Insert_stmtContext) Literal(i int) ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Insert_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Insert_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Insert_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterInsert_stmt(s)
	}
}

func (s *Insert_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitInsert_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Insert_stmt() (localctx IInsert_stmtContext) {
	localctx = NewInsert_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FusionTablesSqlParserRULE_insert_stmt)
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
		p.SetState(164)
		p.Match(FusionTablesSqlParserK_INSERT)
	}
	{
		p.SetState(165)
		p.Match(FusionTablesSqlParserK_INTO)
	}
	{
		p.SetState(166)
		p.Table_name_in_dml()
	}

	{
		p.SetState(167)
		p.Match(FusionTablesSqlParserLPAR)
	}
	{
		p.SetState(168)
		p.Column_name_in_dml()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FusionTablesSqlParserT__2 {
		{
			p.SetState(169)
			p.Match(FusionTablesSqlParserT__2)
		}
		{
			p.SetState(170)
			p.Column_name_in_dml()
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(FusionTablesSqlParserRPAR)
	}

	{
		p.SetState(178)
		p.Match(FusionTablesSqlParserK_VALUES)
	}
	{
		p.SetState(179)
		p.Match(FusionTablesSqlParserLPAR)
	}
	{
		p.SetState(180)
		p.Literal()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FusionTablesSqlParserT__2 {
		{
			p.SetState(181)
			p.Match(FusionTablesSqlParserT__2)
		}
		{
			p.SetState(182)
			p.Literal()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(188)
		p.Match(FusionTablesSqlParserRPAR)
	}

	return localctx
}

// IUpdate_stmtContext is an interface to support dynamic dispatch.
type IUpdate_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_stmtContext differentiates from other interfaces.
	IsUpdate_stmtContext()
}

type Update_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_stmtContext() *Update_stmtContext {
	var p = new(Update_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_update_stmt
	return p
}

func (*Update_stmtContext) IsUpdate_stmtContext() {}

func NewUpdate_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_stmtContext {
	var p = new(Update_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_update_stmt

	return p
}

func (s *Update_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_stmtContext) K_UPDATE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_UPDATE, 0)
}

func (s *Update_stmtContext) Table_name_in_dml() ITable_name_in_dmlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_name_in_dmlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_name_in_dmlContext)
}

func (s *Update_stmtContext) K_SET() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_SET, 0)
}

func (s *Update_stmtContext) AllColumn_assignment() []IColumn_assignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumn_assignmentContext)(nil)).Elem())
	var tst = make([]IColumn_assignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumn_assignmentContext)
		}
	}

	return tst
}

func (s *Update_stmtContext) Column_assignment(i int) IColumn_assignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_assignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumn_assignmentContext)
}

func (s *Update_stmtContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_WHERE, 0)
}

func (s *Update_stmtContext) Eq_comparison() IEq_comparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEq_comparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEq_comparisonContext)
}

func (s *Update_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterUpdate_stmt(s)
	}
}

func (s *Update_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitUpdate_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Update_stmt() (localctx IUpdate_stmtContext) {
	localctx = NewUpdate_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FusionTablesSqlParserRULE_update_stmt)
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
		p.SetState(190)
		p.Match(FusionTablesSqlParserK_UPDATE)
	}
	{
		p.SetState(191)
		p.Table_name_in_dml()
	}
	{
		p.SetState(192)
		p.Match(FusionTablesSqlParserK_SET)
	}
	{
		p.SetState(193)
		p.Column_assignment()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FusionTablesSqlParserT__2 {
		{
			p.SetState(194)
			p.Match(FusionTablesSqlParserT__2)
		}
		{
			p.SetState(195)
			p.Column_assignment()
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(201)
		p.Match(FusionTablesSqlParserK_WHERE)
	}
	{
		p.SetState(202)
		p.Eq_comparison()
	}

	return localctx
}

// IColumn_assignmentContext is an interface to support dynamic dispatch.
type IColumn_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_assignmentContext differentiates from other interfaces.
	IsColumn_assignmentContext()
}

type Column_assignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_assignmentContext() *Column_assignmentContext {
	var p = new(Column_assignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_column_assignment
	return p
}

func (*Column_assignmentContext) IsColumn_assignmentContext() {}

func NewColumn_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_assignmentContext {
	var p = new(Column_assignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_column_assignment

	return p
}

func (s *Column_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_assignmentContext) Column_name_in_dml() IColumn_name_in_dmlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_name_in_dmlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_name_in_dmlContext)
}

func (s *Column_assignmentContext) EQ() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserEQ, 0)
}

func (s *Column_assignmentContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Column_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterColumn_assignment(s)
	}
}

func (s *Column_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitColumn_assignment(s)
	}
}

func (p *FusionTablesSqlParser) Column_assignment() (localctx IColumn_assignmentContext) {
	localctx = NewColumn_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FusionTablesSqlParserRULE_column_assignment)

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
		p.Column_name_in_dml()
	}
	{
		p.SetState(205)
		p.Match(FusionTablesSqlParserEQ)
	}
	{
		p.SetState(206)
		p.Literal()
	}

	return localctx
}

// IDelete_stmtContext is an interface to support dynamic dispatch.
type IDelete_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelete_stmtContext differentiates from other interfaces.
	IsDelete_stmtContext()
}

type Delete_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_stmtContext() *Delete_stmtContext {
	var p = new(Delete_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_delete_stmt
	return p
}

func (*Delete_stmtContext) IsDelete_stmtContext() {}

func NewDelete_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_stmtContext {
	var p = new(Delete_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_delete_stmt

	return p
}

func (s *Delete_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_stmtContext) K_DELETE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_DELETE, 0)
}

func (s *Delete_stmtContext) K_FROM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_FROM, 0)
}

func (s *Delete_stmtContext) Table_name_in_dml() ITable_name_in_dmlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_name_in_dmlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_name_in_dmlContext)
}

func (s *Delete_stmtContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_WHERE, 0)
}

func (s *Delete_stmtContext) Column_name_in_dml() IColumn_name_in_dmlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_name_in_dmlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_name_in_dmlContext)
}

func (s *Delete_stmtContext) EQ() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserEQ, 0)
}

func (s *Delete_stmtContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Delete_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterDelete_stmt(s)
	}
}

func (s *Delete_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitDelete_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Delete_stmt() (localctx IDelete_stmtContext) {
	localctx = NewDelete_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FusionTablesSqlParserRULE_delete_stmt)
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
		p.SetState(208)
		p.Match(FusionTablesSqlParserK_DELETE)
	}
	{
		p.SetState(209)
		p.Match(FusionTablesSqlParserK_FROM)
	}
	{
		p.SetState(210)
		p.Table_name_in_dml()
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FusionTablesSqlParserK_WHERE {
		{
			p.SetState(211)
			p.Match(FusionTablesSqlParserK_WHERE)
		}
		{
			p.SetState(212)
			p.Column_name_in_dml()
		}
		{
			p.SetState(213)
			p.Match(FusionTablesSqlParserEQ)
		}
		{
			p.SetState(214)
			p.Literal()
		}

	}

	return localctx
}

// IEq_comparisonContext is an interface to support dynamic dispatch.
type IEq_comparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEq_comparisonContext differentiates from other interfaces.
	IsEq_comparisonContext()
}

type Eq_comparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEq_comparisonContext() *Eq_comparisonContext {
	var p = new(Eq_comparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_eq_comparison
	return p
}

func (*Eq_comparisonContext) IsEq_comparisonContext() {}

func NewEq_comparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Eq_comparisonContext {
	var p = new(Eq_comparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_eq_comparison

	return p
}

func (s *Eq_comparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *Eq_comparisonContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Eq_comparisonContext) EQ() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserEQ, 0)
}

func (s *Eq_comparisonContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Eq_comparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Eq_comparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Eq_comparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterEq_comparison(s)
	}
}

func (s *Eq_comparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitEq_comparison(s)
	}
}

func (p *FusionTablesSqlParser) Eq_comparison() (localctx IEq_comparisonContext) {
	localctx = NewEq_comparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FusionTablesSqlParserRULE_eq_comparison)

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
		p.Identifier()
	}
	{
		p.SetState(219)
		p.Match(FusionTablesSqlParserEQ)
	}
	{
		p.SetState(220)
		p.String_literal()
	}

	return localctx
}

// ITable_name_with_aliasContext is an interface to support dynamic dispatch.
type ITable_name_with_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_name_with_aliasContext differentiates from other interfaces.
	IsTable_name_with_aliasContext()
}

type Table_name_with_aliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_name_with_aliasContext() *Table_name_with_aliasContext {
	var p = new(Table_name_with_aliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_table_name_with_alias
	return p
}

func (*Table_name_with_aliasContext) IsTable_name_with_aliasContext() {}

func NewTable_name_with_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_name_with_aliasContext {
	var p = new(Table_name_with_aliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_table_name_with_alias

	return p
}

func (s *Table_name_with_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_name_with_aliasContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Table_name_with_aliasContext) K_AS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_AS, 0)
}

func (s *Table_name_with_aliasContext) Table_alias() ITable_aliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_aliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_aliasContext)
}

func (s *Table_name_with_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_name_with_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_name_with_aliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterTable_name_with_alias(s)
	}
}

func (s *Table_name_with_aliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitTable_name_with_alias(s)
	}
}

func (p *FusionTablesSqlParser) Table_name_with_alias() (localctx ITable_name_with_aliasContext) {
	localctx = NewTable_name_with_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FusionTablesSqlParserRULE_table_name_with_alias)
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
		p.SetState(222)
		p.Table_name()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FusionTablesSqlParserK_AS {
		{
			p.SetState(223)
			p.Match(FusionTablesSqlParserK_AS)
		}
		{
			p.SetState(224)
			p.Table_alias()
		}

	}

	return localctx
}

// ISelect_stmtContext is an interface to support dynamic dispatch.
type ISelect_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelect_stmtContext differentiates from other interfaces.
	IsSelect_stmtContext()
}

type Select_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_stmtContext() *Select_stmtContext {
	var p = new(Select_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_select_stmt
	return p
}

func (*Select_stmtContext) IsSelect_stmtContext() {}

func NewSelect_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_stmtContext {
	var p = new(Select_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_select_stmt

	return p
}

func (s *Select_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_stmtContext) K_SELECT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_SELECT, 0)
}

func (s *Select_stmtContext) AllResult_column() []IResult_columnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IResult_columnContext)(nil)).Elem())
	var tst = make([]IResult_columnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IResult_columnContext)
		}
	}

	return tst
}

func (s *Select_stmtContext) Result_column(i int) IResult_columnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResult_columnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IResult_columnContext)
}

func (s *Select_stmtContext) K_FROM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_FROM, 0)
}

func (s *Select_stmtContext) Table_name_with_alias() ITable_name_with_aliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_name_with_aliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_name_with_aliasContext)
}

func (s *Select_stmtContext) AllJoin_clause() []IJoin_clauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJoin_clauseContext)(nil)).Elem())
	var tst = make([]IJoin_clauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJoin_clauseContext)
		}
	}

	return tst
}

func (s *Select_stmtContext) Join_clause(i int) IJoin_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJoin_clauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJoin_clauseContext)
}

func (s *Select_stmtContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_WHERE, 0)
}

func (s *Select_stmtContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Select_stmtContext) K_GROUP() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_GROUP, 0)
}

func (s *Select_stmtContext) AllK_BY() []antlr.TerminalNode {
	return s.GetTokens(FusionTablesSqlParserK_BY)
}

func (s *Select_stmtContext) K_BY(i int) antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_BY, i)
}

func (s *Select_stmtContext) AllQualified_column_name() []IQualified_column_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQualified_column_nameContext)(nil)).Elem())
	var tst = make([]IQualified_column_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQualified_column_nameContext)
		}
	}

	return tst
}

func (s *Select_stmtContext) Qualified_column_name(i int) IQualified_column_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualified_column_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQualified_column_nameContext)
}

func (s *Select_stmtContext) K_ORDER() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ORDER, 0)
}

func (s *Select_stmtContext) AllOrdering_term() []IOrdering_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrdering_termContext)(nil)).Elem())
	var tst = make([]IOrdering_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrdering_termContext)
		}
	}

	return tst
}

func (s *Select_stmtContext) Ordering_term(i int) IOrdering_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrdering_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrdering_termContext)
}

func (s *Select_stmtContext) K_OFFSET() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_OFFSET, 0)
}

func (s *Select_stmtContext) AllNumeric_literal() []INumeric_literalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumeric_literalContext)(nil)).Elem())
	var tst = make([]INumeric_literalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumeric_literalContext)
		}
	}

	return tst
}

func (s *Select_stmtContext) Numeric_literal(i int) INumeric_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_literalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumeric_literalContext)
}

func (s *Select_stmtContext) K_LIMIT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_LIMIT, 0)
}

func (s *Select_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterSelect_stmt(s)
	}
}

func (s *Select_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitSelect_stmt(s)
	}
}

func (p *FusionTablesSqlParser) Select_stmt() (localctx ISelect_stmtContext) {
	localctx = NewSelect_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FusionTablesSqlParserRULE_select_stmt)
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
		p.SetState(227)
		p.Match(FusionTablesSqlParserK_SELECT)
	}
	{
		p.SetState(228)
		p.Result_column()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FusionTablesSqlParserT__2 {
		{
			p.SetState(229)
			p.Match(FusionTablesSqlParserT__2)
		}
		{
			p.SetState(230)
			p.Result_column()
		}

		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(236)
		p.Match(FusionTablesSqlParserK_FROM)
	}
	{
		p.SetState(237)
		p.Table_name_with_alias()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FusionTablesSqlParserK_LEFT {
		{
			p.SetState(238)
			p.Join_clause()
		}

		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FusionTablesSqlParserK_WHERE {
		{
			p.SetState(244)
			p.Match(FusionTablesSqlParserK_WHERE)
		}
		{
			p.SetState(245)
			p.Expr()
		}

	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FusionTablesSqlParserK_GROUP {
		{
			p.SetState(248)
			p.Match(FusionTablesSqlParserK_GROUP)
		}
		{
			p.SetState(249)
			p.Match(FusionTablesSqlParserK_BY)
		}
		{
			p.SetState(250)
			p.Qualified_column_name()
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FusionTablesSqlParserT__2 {
			{
				p.SetState(251)
				p.Match(FusionTablesSqlParserT__2)
			}
			{
				p.SetState(252)
				p.Qualified_column_name()
			}

			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FusionTablesSqlParserK_ORDER {
		{
			p.SetState(260)
			p.Match(FusionTablesSqlParserK_ORDER)
		}
		{
			p.SetState(261)
			p.Match(FusionTablesSqlParserK_BY)
		}
		{
			p.SetState(262)
			p.Ordering_term()
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FusionTablesSqlParserT__2 {
			{
				p.SetState(263)
				p.Match(FusionTablesSqlParserT__2)
			}
			{
				p.SetState(264)
				p.Ordering_term()
			}

			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FusionTablesSqlParserK_OFFSET {
		{
			p.SetState(272)
			p.Match(FusionTablesSqlParserK_OFFSET)
		}
		{
			p.SetState(273)
			p.Numeric_literal()
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FusionTablesSqlParserK_LIMIT {
			{
				p.SetState(274)
				p.Match(FusionTablesSqlParserK_LIMIT)
			}
			{
				p.SetState(275)
				p.Numeric_literal()
			}

		}

	}

	return localctx
}

// IOrdering_termContext is an interface to support dynamic dispatch.
type IOrdering_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrdering_termContext differentiates from other interfaces.
	IsOrdering_termContext()
}

type Ordering_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrdering_termContext() *Ordering_termContext {
	var p = new(Ordering_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_ordering_term
	return p
}

func (*Ordering_termContext) IsOrdering_termContext() {}

func NewOrdering_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ordering_termContext {
	var p = new(Ordering_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_ordering_term

	return p
}

func (s *Ordering_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Ordering_termContext) Qualified_column_name() IQualified_column_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualified_column_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualified_column_nameContext)
}

func (s *Ordering_termContext) K_ST_DISTANCE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ST_DISTANCE, 0)
}

func (s *Ordering_termContext) LPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserLPAR, 0)
}

func (s *Ordering_termContext) Coordinate() ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *Ordering_termContext) RPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserRPAR, 0)
}

func (s *Ordering_termContext) K_ASC() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ASC, 0)
}

func (s *Ordering_termContext) K_DESC() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_DESC, 0)
}

func (s *Ordering_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ordering_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ordering_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterOrdering_term(s)
	}
}

func (s *Ordering_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitOrdering_term(s)
	}
}

func (p *FusionTablesSqlParser) Ordering_term() (localctx IOrdering_termContext) {
	localctx = NewOrdering_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FusionTablesSqlParserRULE_ordering_term)
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
	p.SetState(288)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FusionTablesSqlParserSTRING_LITERAL:
		{
			p.SetState(280)
			p.Qualified_column_name()
		}

	case FusionTablesSqlParserK_ST_DISTANCE:
		{
			p.SetState(281)
			p.Match(FusionTablesSqlParserK_ST_DISTANCE)
		}
		{
			p.SetState(282)
			p.Match(FusionTablesSqlParserLPAR)
		}
		{
			p.SetState(283)
			p.Qualified_column_name()
		}
		{
			p.SetState(284)
			p.Match(FusionTablesSqlParserT__2)
		}
		{
			p.SetState(285)
			p.Coordinate()
		}
		{
			p.SetState(286)
			p.Match(FusionTablesSqlParserRPAR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FusionTablesSqlParserK_ASC || _la == FusionTablesSqlParserK_DESC {
		{
			p.SetState(290)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FusionTablesSqlParserK_ASC || _la == FusionTablesSqlParserK_DESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IJoin_clauseContext is an interface to support dynamic dispatch.
type IJoin_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJoin_clauseContext differentiates from other interfaces.
	IsJoin_clauseContext()
}

type Join_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoin_clauseContext() *Join_clauseContext {
	var p = new(Join_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_join_clause
	return p
}

func (*Join_clauseContext) IsJoin_clauseContext() {}

func NewJoin_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Join_clauseContext {
	var p = new(Join_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_join_clause

	return p
}

func (s *Join_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Join_clauseContext) K_LEFT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_LEFT, 0)
}

func (s *Join_clauseContext) K_OUTER() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_OUTER, 0)
}

func (s *Join_clauseContext) K_JOIN() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_JOIN, 0)
}

func (s *Join_clauseContext) Table_name_with_alias() ITable_name_with_aliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_name_with_aliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_name_with_aliasContext)
}

func (s *Join_clauseContext) K_ON() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ON, 0)
}

func (s *Join_clauseContext) AllQualified_column_name() []IQualified_column_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQualified_column_nameContext)(nil)).Elem())
	var tst = make([]IQualified_column_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQualified_column_nameContext)
		}
	}

	return tst
}

func (s *Join_clauseContext) Qualified_column_name(i int) IQualified_column_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualified_column_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQualified_column_nameContext)
}

func (s *Join_clauseContext) EQ() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserEQ, 0)
}

func (s *Join_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Join_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Join_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterJoin_clause(s)
	}
}

func (s *Join_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitJoin_clause(s)
	}
}

func (p *FusionTablesSqlParser) Join_clause() (localctx IJoin_clauseContext) {
	localctx = NewJoin_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FusionTablesSqlParserRULE_join_clause)

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
		p.SetState(293)
		p.Match(FusionTablesSqlParserK_LEFT)
	}
	{
		p.SetState(294)
		p.Match(FusionTablesSqlParserK_OUTER)
	}
	{
		p.SetState(295)
		p.Match(FusionTablesSqlParserK_JOIN)
	}
	{
		p.SetState(296)
		p.Table_name_with_alias()
	}
	{
		p.SetState(297)
		p.Match(FusionTablesSqlParserK_ON)
	}
	{
		p.SetState(298)
		p.Qualified_column_name()
	}
	{
		p.SetState(299)
		p.Match(FusionTablesSqlParserEQ)
	}
	{
		p.SetState(300)
		p.Qualified_column_name()
	}

	return localctx
}

// IResult_columnContext is an interface to support dynamic dispatch.
type IResult_columnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResult_columnContext differentiates from other interfaces.
	IsResult_columnContext()
}

type Result_columnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResult_columnContext() *Result_columnContext {
	var p = new(Result_columnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_result_column
	return p
}

func (*Result_columnContext) IsResult_columnContext() {}

func NewResult_columnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Result_columnContext {
	var p = new(Result_columnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_result_column

	return p
}

func (s *Result_columnContext) GetParser() antlr.Parser { return s.parser }

func (s *Result_columnContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Result_columnContext) Qualified_column_name() IQualified_column_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualified_column_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualified_column_nameContext)
}

func (s *Result_columnContext) Aggregate_exp() IAggregate_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAggregate_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAggregate_expContext)
}

func (s *Result_columnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Result_columnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Result_columnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterResult_column(s)
	}
}

func (s *Result_columnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitResult_column(s)
	}
}

func (p *FusionTablesSqlParser) Result_column() (localctx IResult_columnContext) {
	localctx = NewResult_columnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FusionTablesSqlParserRULE_result_column)

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

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(302)
			p.Match(FusionTablesSqlParserT__1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(303)
			p.Table_name()
		}
		{
			p.SetState(304)
			p.Match(FusionTablesSqlParserT__3)
		}
		{
			p.SetState(305)
			p.Match(FusionTablesSqlParserT__1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(307)
			p.Qualified_column_name()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(308)
			p.Aggregate_exp()
		}

	}

	return localctx
}

// IQualified_column_nameContext is an interface to support dynamic dispatch.
type IQualified_column_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualified_column_nameContext differentiates from other interfaces.
	IsQualified_column_nameContext()
}

type Qualified_column_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualified_column_nameContext() *Qualified_column_nameContext {
	var p = new(Qualified_column_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_qualified_column_name
	return p
}

func (*Qualified_column_nameContext) IsQualified_column_nameContext() {}

func NewQualified_column_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Qualified_column_nameContext {
	var p = new(Qualified_column_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_qualified_column_name

	return p
}

func (s *Qualified_column_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Qualified_column_nameContext) Column_name() IColumn_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Qualified_column_nameContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Qualified_column_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Qualified_column_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Qualified_column_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterQualified_column_name(s)
	}
}

func (s *Qualified_column_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitQualified_column_name(s)
	}
}

func (p *FusionTablesSqlParser) Qualified_column_name() (localctx IQualified_column_nameContext) {
	localctx = NewQualified_column_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FusionTablesSqlParserRULE_qualified_column_name)

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
	p.SetState(314)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(311)
			p.Table_name()
		}
		{
			p.SetState(312)
			p.Match(FusionTablesSqlParserT__3)
		}

	}
	{
		p.SetState(316)
		p.Column_name()
	}

	return localctx
}

// IAggregate_expContext is an interface to support dynamic dispatch.
type IAggregate_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAggregate_expContext differentiates from other interfaces.
	IsAggregate_expContext()
}

type Aggregate_expContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregate_expContext() *Aggregate_expContext {
	var p = new(Aggregate_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_aggregate_exp
	return p
}

func (*Aggregate_expContext) IsAggregate_expContext() {}

func NewAggregate_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aggregate_expContext {
	var p = new(Aggregate_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_aggregate_exp

	return p
}

func (s *Aggregate_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Aggregate_expContext) LPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserLPAR, 0)
}

func (s *Aggregate_expContext) Qualified_column_name() IQualified_column_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualified_column_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualified_column_nameContext)
}

func (s *Aggregate_expContext) RPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserRPAR, 0)
}

func (s *Aggregate_expContext) K_SUM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_SUM, 0)
}

func (s *Aggregate_expContext) K_COUNT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_COUNT, 0)
}

func (s *Aggregate_expContext) K_AVERAGE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_AVERAGE, 0)
}

func (s *Aggregate_expContext) K_MAXIMUM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_MAXIMUM, 0)
}

func (s *Aggregate_expContext) K_MINIMUM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_MINIMUM, 0)
}

func (s *Aggregate_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aggregate_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Aggregate_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterAggregate_exp(s)
	}
}

func (s *Aggregate_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitAggregate_exp(s)
	}
}

func (p *FusionTablesSqlParser) Aggregate_exp() (localctx IAggregate_expContext) {
	localctx = NewAggregate_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FusionTablesSqlParserRULE_aggregate_exp)
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
		p.SetState(318)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FusionTablesSqlParserK_AVERAGE || _la == FusionTablesSqlParserK_COUNT || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(FusionTablesSqlParserK_MAXIMUM-39))|(1<<(FusionTablesSqlParserK_MINIMUM-39))|(1<<(FusionTablesSqlParserK_SUM-39)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(319)
		p.Match(FusionTablesSqlParserLPAR)
	}
	{
		p.SetState(320)
		p.Qualified_column_name()
	}
	{
		p.SetState(321)
		p.Match(FusionTablesSqlParserRPAR)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Column_name_beginning_expr() IColumn_name_beginning_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_name_beginning_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_name_beginning_exprContext)
}

func (s *ExprContext) AllLiteral() []ILiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralContext)(nil)).Elem())
	var tst = make([]ILiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralContext)
		}
	}

	return tst
}

func (s *ExprContext) Literal(i int) ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprContext) Operator() IOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ExprContext) And_or_or() IAnd_or_orContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_or_orContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnd_or_orContext)
}

func (s *ExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) AllString_literal() []IString_literalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IString_literalContext)(nil)).Elem())
	var tst = make([]IString_literalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IString_literalContext)
		}
	}

	return tst
}

func (s *ExprContext) String_literal(i int) IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *ExprContext) K_LIKE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_LIKE, 0)
}

func (s *ExprContext) K_MATCHES() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_MATCHES, 0)
}

func (s *ExprContext) K_STARTS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_STARTS, 0)
}

func (s *ExprContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_WITH, 0)
}

func (s *ExprContext) K_ENDS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ENDS, 0)
}

func (s *ExprContext) K_CONTAINS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CONTAINS, 0)
}

func (s *ExprContext) K_IGNORING() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_IGNORING, 0)
}

func (s *ExprContext) K_CASE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CASE, 0)
}

func (s *ExprContext) K_DOES() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_DOES, 0)
}

func (s *ExprContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_NOT, 0)
}

func (s *ExprContext) K_CONTAIN() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CONTAIN, 0)
}

func (s *ExprContext) K_EQUAL() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_EQUAL, 0)
}

func (s *ExprContext) K_TO() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_TO, 0)
}

func (s *ExprContext) K_IN() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_IN, 0)
}

func (s *ExprContext) LPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserLPAR, 0)
}

func (s *ExprContext) RPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserRPAR, 0)
}

func (s *ExprContext) K_BETWEEN() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_BETWEEN, 0)
}

func (s *ExprContext) K_AND() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_AND, 0)
}

func (s *ExprContext) K_ST_INTERSECTS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ST_INTERSECTS, 0)
}

func (s *ExprContext) Qualified_column_name() IQualified_column_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualified_column_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualified_column_nameContext)
}

func (s *ExprContext) Geometry() IGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeometryContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *FusionTablesSqlParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FusionTablesSqlParserRULE_expr)
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

	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)
			p.Column_name_beginning_expr()
		}

		{
			p.SetState(324)
			p.Operator()
		}

		{
			p.SetState(325)
			p.Literal()
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FusionTablesSqlParserK_AND || _la == FusionTablesSqlParserK_OR {
			{
				p.SetState(326)
				p.And_or_or()
			}
			{
				p.SetState(327)
				p.Expr()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.Column_name_beginning_expr()
		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(332)
				p.Match(FusionTablesSqlParserK_LIKE)
			}

		case 2:
			{
				p.SetState(333)
				p.Match(FusionTablesSqlParserK_MATCHES)
			}

		case 3:
			{
				p.SetState(334)
				p.Match(FusionTablesSqlParserK_STARTS)
			}
			{
				p.SetState(335)
				p.Match(FusionTablesSqlParserK_WITH)
			}

		case 4:
			{
				p.SetState(336)
				p.Match(FusionTablesSqlParserK_ENDS)
			}
			{
				p.SetState(337)
				p.Match(FusionTablesSqlParserK_WITH)
			}

		case 5:
			{
				p.SetState(338)
				p.Match(FusionTablesSqlParserK_CONTAINS)
			}

		case 6:
			{
				p.SetState(339)
				p.Match(FusionTablesSqlParserK_CONTAINS)
			}
			{
				p.SetState(340)
				p.Match(FusionTablesSqlParserK_IGNORING)
			}
			{
				p.SetState(341)
				p.Match(FusionTablesSqlParserK_CASE)
			}

		case 7:
			{
				p.SetState(342)
				p.Match(FusionTablesSqlParserK_DOES)
			}
			{
				p.SetState(343)
				p.Match(FusionTablesSqlParserK_NOT)
			}
			{
				p.SetState(344)
				p.Match(FusionTablesSqlParserK_CONTAIN)
			}

		case 8:
			{
				p.SetState(345)
				p.Match(FusionTablesSqlParserK_NOT)
			}
			{
				p.SetState(346)
				p.Match(FusionTablesSqlParserK_EQUAL)
			}
			{
				p.SetState(347)
				p.Match(FusionTablesSqlParserK_TO)
			}

		}
		{
			p.SetState(350)
			p.String_literal()
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FusionTablesSqlParserK_AND || _la == FusionTablesSqlParserK_OR {
			{
				p.SetState(351)
				p.And_or_or()
			}
			{
				p.SetState(352)
				p.Expr()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(356)
			p.Column_name_beginning_expr()
		}
		{
			p.SetState(357)
			p.Match(FusionTablesSqlParserK_IN)
		}
		{
			p.SetState(358)
			p.Match(FusionTablesSqlParserLPAR)
		}
		{
			p.SetState(359)
			p.String_literal()
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FusionTablesSqlParserT__2 {
			{
				p.SetState(360)
				p.Match(FusionTablesSqlParserT__2)
			}
			{
				p.SetState(361)
				p.String_literal()
			}

			p.SetState(366)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(367)
			p.Match(FusionTablesSqlParserRPAR)
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FusionTablesSqlParserK_AND || _la == FusionTablesSqlParserK_OR {
			{
				p.SetState(368)
				p.And_or_or()
			}
			{
				p.SetState(369)
				p.Expr()
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(373)
			p.Column_name_beginning_expr()
		}
		{
			p.SetState(374)
			p.Match(FusionTablesSqlParserK_BETWEEN)
		}
		{
			p.SetState(375)
			p.Literal()
		}
		{
			p.SetState(376)
			p.Match(FusionTablesSqlParserK_AND)
		}
		{
			p.SetState(377)
			p.Literal()
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FusionTablesSqlParserK_AND || _la == FusionTablesSqlParserK_OR {
			{
				p.SetState(378)
				p.And_or_or()
			}
			{
				p.SetState(379)
				p.Expr()
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(383)
			p.Match(FusionTablesSqlParserK_ST_INTERSECTS)
		}
		{
			p.SetState(384)
			p.Match(FusionTablesSqlParserLPAR)
		}
		{
			p.SetState(385)
			p.Qualified_column_name()
		}
		{
			p.SetState(386)
			p.Match(FusionTablesSqlParserT__2)
		}
		{
			p.SetState(387)
			p.Geometry()
		}
		{
			p.SetState(388)
			p.Match(FusionTablesSqlParserRPAR)
		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FusionTablesSqlParserK_AND || _la == FusionTablesSqlParserK_OR {
			{
				p.SetState(389)
				p.And_or_or()
			}
			{
				p.SetState(390)
				p.Expr()
			}

		}

	}

	return localctx
}

// IColumn_name_beginning_exprContext is an interface to support dynamic dispatch.
type IColumn_name_beginning_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_name_beginning_exprContext differentiates from other interfaces.
	IsColumn_name_beginning_exprContext()
}

type Column_name_beginning_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_name_beginning_exprContext() *Column_name_beginning_exprContext {
	var p = new(Column_name_beginning_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_column_name_beginning_expr
	return p
}

func (*Column_name_beginning_exprContext) IsColumn_name_beginning_exprContext() {}

func NewColumn_name_beginning_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_name_beginning_exprContext {
	var p = new(Column_name_beginning_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_column_name_beginning_expr

	return p
}

func (s *Column_name_beginning_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_name_beginning_exprContext) Qualified_column_name() IQualified_column_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualified_column_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualified_column_nameContext)
}

func (s *Column_name_beginning_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_name_beginning_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_name_beginning_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterColumn_name_beginning_expr(s)
	}
}

func (s *Column_name_beginning_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitColumn_name_beginning_expr(s)
	}
}

func (p *FusionTablesSqlParser) Column_name_beginning_expr() (localctx IColumn_name_beginning_exprContext) {
	localctx = NewColumn_name_beginning_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FusionTablesSqlParserRULE_column_name_beginning_expr)

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
		p.SetState(396)
		p.Qualified_column_name()
	}

	return localctx
}

// IColumn_name_in_dmlContext is an interface to support dynamic dispatch.
type IColumn_name_in_dmlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_name_in_dmlContext differentiates from other interfaces.
	IsColumn_name_in_dmlContext()
}

type Column_name_in_dmlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_name_in_dmlContext() *Column_name_in_dmlContext {
	var p = new(Column_name_in_dmlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_column_name_in_dml
	return p
}

func (*Column_name_in_dmlContext) IsColumn_name_in_dmlContext() {}

func NewColumn_name_in_dmlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_name_in_dmlContext {
	var p = new(Column_name_in_dmlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_column_name_in_dml

	return p
}

func (s *Column_name_in_dmlContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_name_in_dmlContext) Column_name() IColumn_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Column_name_in_dmlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_name_in_dmlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_name_in_dmlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterColumn_name_in_dml(s)
	}
}

func (s *Column_name_in_dmlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitColumn_name_in_dml(s)
	}
}

func (p *FusionTablesSqlParser) Column_name_in_dml() (localctx IColumn_name_in_dmlContext) {
	localctx = NewColumn_name_in_dmlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FusionTablesSqlParserRULE_column_name_in_dml)

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
		p.SetState(398)
		p.Column_name()
	}

	return localctx
}

// IAnd_or_orContext is an interface to support dynamic dispatch.
type IAnd_or_orContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_or_orContext differentiates from other interfaces.
	IsAnd_or_orContext()
}

type And_or_orContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_or_orContext() *And_or_orContext {
	var p = new(And_or_orContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_and_or_or
	return p
}

func (*And_or_orContext) IsAnd_or_orContext() {}

func NewAnd_or_orContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_or_orContext {
	var p = new(And_or_orContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_and_or_or

	return p
}

func (s *And_or_orContext) GetParser() antlr.Parser { return s.parser }

func (s *And_or_orContext) K_AND() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_AND, 0)
}

func (s *And_or_orContext) K_OR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_OR, 0)
}

func (s *And_or_orContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_or_orContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *And_or_orContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterAnd_or_or(s)
	}
}

func (s *And_or_orContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitAnd_or_or(s)
	}
}

func (p *FusionTablesSqlParser) And_or_or() (localctx IAnd_or_orContext) {
	localctx = NewAnd_or_orContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FusionTablesSqlParserRULE_and_or_or)
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
		p.SetState(400)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FusionTablesSqlParserK_AND || _la == FusionTablesSqlParserK_OR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGeometryContext is an interface to support dynamic dispatch.
type IGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeometryContext differentiates from other interfaces.
	IsGeometryContext()
}

type GeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeometryContext() *GeometryContext {
	var p = new(GeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_geometry
	return p
}

func (*GeometryContext) IsGeometryContext() {}

func NewGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeometryContext {
	var p = new(GeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_geometry

	return p
}

func (s *GeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *GeometryContext) Circle() ICircleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICircleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICircleContext)
}

func (s *GeometryContext) Rectangle() IRectangleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRectangleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRectangleContext)
}

func (s *GeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterGeometry(s)
	}
}

func (s *GeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitGeometry(s)
	}
}

func (p *FusionTablesSqlParser) Geometry() (localctx IGeometryContext) {
	localctx = NewGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FusionTablesSqlParserRULE_geometry)

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

	p.SetState(404)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FusionTablesSqlParserK_CIRCLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(402)
			p.Circle()
		}

	case FusionTablesSqlParserK_RECTANGLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(403)
			p.Rectangle()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICircleContext is an interface to support dynamic dispatch.
type ICircleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCircleContext differentiates from other interfaces.
	IsCircleContext()
}

type CircleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCircleContext() *CircleContext {
	var p = new(CircleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_circle
	return p
}

func (*CircleContext) IsCircleContext() {}

func NewCircleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CircleContext {
	var p = new(CircleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_circle

	return p
}

func (s *CircleContext) GetParser() antlr.Parser { return s.parser }

func (s *CircleContext) K_CIRCLE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CIRCLE, 0)
}

func (s *CircleContext) LPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserLPAR, 0)
}

func (s *CircleContext) Coordinate() ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *CircleContext) Numeric_literal() INumeric_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumeric_literalContext)
}

func (s *CircleContext) RPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserRPAR, 0)
}

func (s *CircleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CircleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CircleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterCircle(s)
	}
}

func (s *CircleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitCircle(s)
	}
}

func (p *FusionTablesSqlParser) Circle() (localctx ICircleContext) {
	localctx = NewCircleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FusionTablesSqlParserRULE_circle)

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
		p.SetState(406)
		p.Match(FusionTablesSqlParserK_CIRCLE)
	}
	{
		p.SetState(407)
		p.Match(FusionTablesSqlParserLPAR)
	}
	{
		p.SetState(408)
		p.Coordinate()
	}
	{
		p.SetState(409)
		p.Match(FusionTablesSqlParserT__2)
	}
	{
		p.SetState(410)
		p.Numeric_literal()
	}
	{
		p.SetState(411)
		p.Match(FusionTablesSqlParserRPAR)
	}

	return localctx
}

// IRectangleContext is an interface to support dynamic dispatch.
type IRectangleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRectangleContext differentiates from other interfaces.
	IsRectangleContext()
}

type RectangleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRectangleContext() *RectangleContext {
	var p = new(RectangleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_rectangle
	return p
}

func (*RectangleContext) IsRectangleContext() {}

func NewRectangleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RectangleContext {
	var p = new(RectangleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_rectangle

	return p
}

func (s *RectangleContext) GetParser() antlr.Parser { return s.parser }

func (s *RectangleContext) K_RECTANGLE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_RECTANGLE, 0)
}

func (s *RectangleContext) LPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserLPAR, 0)
}

func (s *RectangleContext) AllCoordinate() []ICoordinateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICoordinateContext)(nil)).Elem())
	var tst = make([]ICoordinateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICoordinateContext)
		}
	}

	return tst
}

func (s *RectangleContext) Coordinate(i int) ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *RectangleContext) RPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserRPAR, 0)
}

func (s *RectangleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RectangleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RectangleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterRectangle(s)
	}
}

func (s *RectangleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitRectangle(s)
	}
}

func (p *FusionTablesSqlParser) Rectangle() (localctx IRectangleContext) {
	localctx = NewRectangleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FusionTablesSqlParserRULE_rectangle)

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
		p.SetState(413)
		p.Match(FusionTablesSqlParserK_RECTANGLE)
	}
	{
		p.SetState(414)
		p.Match(FusionTablesSqlParserLPAR)
	}
	{
		p.SetState(415)
		p.Coordinate()
	}
	{
		p.SetState(416)
		p.Match(FusionTablesSqlParserT__2)
	}
	{
		p.SetState(417)
		p.Coordinate()
	}
	{
		p.SetState(418)
		p.Match(FusionTablesSqlParserRPAR)
	}

	return localctx
}

// ICoordinateContext is an interface to support dynamic dispatch.
type ICoordinateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCoordinateContext differentiates from other interfaces.
	IsCoordinateContext()
}

type CoordinateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoordinateContext() *CoordinateContext {
	var p = new(CoordinateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_coordinate
	return p
}

func (*CoordinateContext) IsCoordinateContext() {}

func NewCoordinateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordinateContext {
	var p = new(CoordinateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_coordinate

	return p
}

func (s *CoordinateContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordinateContext) K_LATLNG() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_LATLNG, 0)
}

func (s *CoordinateContext) LPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserLPAR, 0)
}

func (s *CoordinateContext) AllNumeric_literal() []INumeric_literalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumeric_literalContext)(nil)).Elem())
	var tst = make([]INumeric_literalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumeric_literalContext)
		}
	}

	return tst
}

func (s *CoordinateContext) Numeric_literal(i int) INumeric_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_literalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumeric_literalContext)
}

func (s *CoordinateContext) RPAR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserRPAR, 0)
}

func (s *CoordinateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoordinateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterCoordinate(s)
	}
}

func (s *CoordinateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitCoordinate(s)
	}
}

func (p *FusionTablesSqlParser) Coordinate() (localctx ICoordinateContext) {
	localctx = NewCoordinateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FusionTablesSqlParserRULE_coordinate)

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
		p.SetState(420)
		p.Match(FusionTablesSqlParserK_LATLNG)
	}
	{
		p.SetState(421)
		p.Match(FusionTablesSqlParserLPAR)
	}
	{
		p.SetState(422)
		p.Numeric_literal()
	}
	{
		p.SetState(423)
		p.Match(FusionTablesSqlParserT__2)
	}
	{
		p.SetState(424)
		p.Numeric_literal()
	}
	{
		p.SetState(425)
		p.Match(FusionTablesSqlParserRPAR)
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) K_ALTER() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ALTER, 0)
}

func (s *KeywordContext) K_AND() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_AND, 0)
}

func (s *KeywordContext) K_OR() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_OR, 0)
}

func (s *KeywordContext) K_AS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_AS, 0)
}

func (s *KeywordContext) K_ASC() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ASC, 0)
}

func (s *KeywordContext) K_AVERAGE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_AVERAGE, 0)
}

func (s *KeywordContext) K_BY() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_BY, 0)
}

func (s *KeywordContext) K_BETWEEN() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_BETWEEN, 0)
}

func (s *KeywordContext) K_CASE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CASE, 0)
}

func (s *KeywordContext) K_CIRCLE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CIRCLE, 0)
}

func (s *KeywordContext) K_COLUMN() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_COLUMN, 0)
}

func (s *KeywordContext) K_CONTAIN() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CONTAIN, 0)
}

func (s *KeywordContext) K_CONTAINS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CONTAINS, 0)
}

func (s *KeywordContext) K_COUNT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_COUNT, 0)
}

func (s *KeywordContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_CREATE, 0)
}

func (s *KeywordContext) K_DELETE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_DELETE, 0)
}

func (s *KeywordContext) K_DESC() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_DESC, 0)
}

func (s *KeywordContext) K_DOES() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_DOES, 0)
}

func (s *KeywordContext) K_DROP() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_DROP, 0)
}

func (s *KeywordContext) K_ENDS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ENDS, 0)
}

func (s *KeywordContext) K_EQUAL() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_EQUAL, 0)
}

func (s *KeywordContext) K_FROM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_FROM, 0)
}

func (s *KeywordContext) K_GROUP() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_GROUP, 0)
}

func (s *KeywordContext) K_HAVING() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_HAVING, 0)
}

func (s *KeywordContext) K_IGNORING() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_IGNORING, 0)
}

func (s *KeywordContext) K_IN() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_IN, 0)
}

func (s *KeywordContext) K_INSERT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_INSERT, 0)
}

func (s *KeywordContext) K_INTO() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_INTO, 0)
}

func (s *KeywordContext) K_JOIN() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_JOIN, 0)
}

func (s *KeywordContext) K_LATLNG() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_LATLNG, 0)
}

func (s *KeywordContext) K_LEFT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_LEFT, 0)
}

func (s *KeywordContext) K_LIKE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_LIKE, 0)
}

func (s *KeywordContext) K_LIMIT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_LIMIT, 0)
}

func (s *KeywordContext) K_MATCHES() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_MATCHES, 0)
}

func (s *KeywordContext) K_MAXIMUM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_MAXIMUM, 0)
}

func (s *KeywordContext) K_MINIMUM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_MINIMUM, 0)
}

func (s *KeywordContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_NOT, 0)
}

func (s *KeywordContext) K_OF() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_OF, 0)
}

func (s *KeywordContext) K_OFFSET() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_OFFSET, 0)
}

func (s *KeywordContext) K_ON() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ON, 0)
}

func (s *KeywordContext) K_ORDER() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ORDER, 0)
}

func (s *KeywordContext) K_OUTER() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_OUTER, 0)
}

func (s *KeywordContext) K_RECTANGLE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_RECTANGLE, 0)
}

func (s *KeywordContext) K_RENAME() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_RENAME, 0)
}

func (s *KeywordContext) K_SELECT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_SELECT, 0)
}

func (s *KeywordContext) K_SET() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_SET, 0)
}

func (s *KeywordContext) K_STARTS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_STARTS, 0)
}

func (s *KeywordContext) K_ST_DISTANCE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ST_DISTANCE, 0)
}

func (s *KeywordContext) K_ST_INTERSECTS() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_ST_INTERSECTS, 0)
}

func (s *KeywordContext) K_SUM() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_SUM, 0)
}

func (s *KeywordContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_TABLE, 0)
}

func (s *KeywordContext) K_TO() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_TO, 0)
}

func (s *KeywordContext) K_UPDATE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_UPDATE, 0)
}

func (s *KeywordContext) K_VALUES() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_VALUES, 0)
}

func (s *KeywordContext) K_VIEW() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_VIEW, 0)
}

func (s *KeywordContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_WHERE, 0)
}

func (s *KeywordContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserK_WITH, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *FusionTablesSqlParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FusionTablesSqlParserRULE_keyword)

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

	p.SetState(485)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FusionTablesSqlParserEOF:
		p.EnterOuterAlt(localctx, 1)

	case FusionTablesSqlParserK_ALTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(428)
			p.Match(FusionTablesSqlParserK_ALTER)
		}

	case FusionTablesSqlParserK_AND:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(429)
			p.Match(FusionTablesSqlParserK_AND)
		}

	case FusionTablesSqlParserK_OR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(430)
			p.Match(FusionTablesSqlParserK_OR)
		}

	case FusionTablesSqlParserK_AS:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(431)
			p.Match(FusionTablesSqlParserK_AS)
		}

	case FusionTablesSqlParserK_ASC:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(432)
			p.Match(FusionTablesSqlParserK_ASC)
		}

	case FusionTablesSqlParserK_AVERAGE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(433)
			p.Match(FusionTablesSqlParserK_AVERAGE)
		}

	case FusionTablesSqlParserK_BY:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(434)
			p.Match(FusionTablesSqlParserK_BY)
		}

	case FusionTablesSqlParserK_BETWEEN:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(435)
			p.Match(FusionTablesSqlParserK_BETWEEN)
		}

	case FusionTablesSqlParserK_CASE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(436)
			p.Match(FusionTablesSqlParserK_CASE)
		}

	case FusionTablesSqlParserK_CIRCLE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(437)
			p.Match(FusionTablesSqlParserK_CIRCLE)
		}

	case FusionTablesSqlParserK_COLUMN:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(438)
			p.Match(FusionTablesSqlParserK_COLUMN)
		}

	case FusionTablesSqlParserK_CONTAIN:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(439)
			p.Match(FusionTablesSqlParserK_CONTAIN)
		}

	case FusionTablesSqlParserK_CONTAINS:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(440)
			p.Match(FusionTablesSqlParserK_CONTAINS)
		}

	case FusionTablesSqlParserK_COUNT:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(441)
			p.Match(FusionTablesSqlParserK_COUNT)
		}

	case FusionTablesSqlParserK_CREATE:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(442)
			p.Match(FusionTablesSqlParserK_CREATE)
		}

	case FusionTablesSqlParserK_DELETE:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(443)
			p.Match(FusionTablesSqlParserK_DELETE)
		}

	case FusionTablesSqlParserK_DESC:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(444)
			p.Match(FusionTablesSqlParserK_DESC)
		}

	case FusionTablesSqlParserK_DOES:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(445)
			p.Match(FusionTablesSqlParserK_DOES)
		}

	case FusionTablesSqlParserK_DROP:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(446)
			p.Match(FusionTablesSqlParserK_DROP)
		}

	case FusionTablesSqlParserK_ENDS:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(447)
			p.Match(FusionTablesSqlParserK_ENDS)
		}

	case FusionTablesSqlParserK_EQUAL:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(448)
			p.Match(FusionTablesSqlParserK_EQUAL)
		}

	case FusionTablesSqlParserK_FROM:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(449)
			p.Match(FusionTablesSqlParserK_FROM)
		}

	case FusionTablesSqlParserK_GROUP:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(450)
			p.Match(FusionTablesSqlParserK_GROUP)
		}

	case FusionTablesSqlParserK_HAVING:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(451)
			p.Match(FusionTablesSqlParserK_HAVING)
		}

	case FusionTablesSqlParserK_IGNORING:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(452)
			p.Match(FusionTablesSqlParserK_IGNORING)
		}

	case FusionTablesSqlParserK_IN:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(453)
			p.Match(FusionTablesSqlParserK_IN)
		}

	case FusionTablesSqlParserK_INSERT:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(454)
			p.Match(FusionTablesSqlParserK_INSERT)
		}

	case FusionTablesSqlParserK_INTO:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(455)
			p.Match(FusionTablesSqlParserK_INTO)
		}

	case FusionTablesSqlParserK_JOIN:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(456)
			p.Match(FusionTablesSqlParserK_JOIN)
		}

	case FusionTablesSqlParserK_LATLNG:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(457)
			p.Match(FusionTablesSqlParserK_LATLNG)
		}

	case FusionTablesSqlParserK_LEFT:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(458)
			p.Match(FusionTablesSqlParserK_LEFT)
		}

	case FusionTablesSqlParserK_LIKE:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(459)
			p.Match(FusionTablesSqlParserK_LIKE)
		}

	case FusionTablesSqlParserK_LIMIT:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(460)
			p.Match(FusionTablesSqlParserK_LIMIT)
		}

	case FusionTablesSqlParserK_MATCHES:
		p.EnterOuterAlt(localctx, 35)
		{
			p.SetState(461)
			p.Match(FusionTablesSqlParserK_MATCHES)
		}

	case FusionTablesSqlParserK_MAXIMUM:
		p.EnterOuterAlt(localctx, 36)
		{
			p.SetState(462)
			p.Match(FusionTablesSqlParserK_MAXIMUM)
		}

	case FusionTablesSqlParserK_MINIMUM:
		p.EnterOuterAlt(localctx, 37)
		{
			p.SetState(463)
			p.Match(FusionTablesSqlParserK_MINIMUM)
		}

	case FusionTablesSqlParserK_NOT:
		p.EnterOuterAlt(localctx, 38)
		{
			p.SetState(464)
			p.Match(FusionTablesSqlParserK_NOT)
		}

	case FusionTablesSqlParserK_OF:
		p.EnterOuterAlt(localctx, 39)
		{
			p.SetState(465)
			p.Match(FusionTablesSqlParserK_OF)
		}

	case FusionTablesSqlParserK_OFFSET:
		p.EnterOuterAlt(localctx, 40)
		{
			p.SetState(466)
			p.Match(FusionTablesSqlParserK_OFFSET)
		}

	case FusionTablesSqlParserK_ON:
		p.EnterOuterAlt(localctx, 41)
		{
			p.SetState(467)
			p.Match(FusionTablesSqlParserK_ON)
		}

	case FusionTablesSqlParserK_ORDER:
		p.EnterOuterAlt(localctx, 42)
		{
			p.SetState(468)
			p.Match(FusionTablesSqlParserK_ORDER)
		}

	case FusionTablesSqlParserK_OUTER:
		p.EnterOuterAlt(localctx, 43)
		{
			p.SetState(469)
			p.Match(FusionTablesSqlParserK_OUTER)
		}

	case FusionTablesSqlParserK_RECTANGLE:
		p.EnterOuterAlt(localctx, 44)
		{
			p.SetState(470)
			p.Match(FusionTablesSqlParserK_RECTANGLE)
		}

	case FusionTablesSqlParserK_RENAME:
		p.EnterOuterAlt(localctx, 45)
		{
			p.SetState(471)
			p.Match(FusionTablesSqlParserK_RENAME)
		}

	case FusionTablesSqlParserK_SELECT:
		p.EnterOuterAlt(localctx, 46)
		{
			p.SetState(472)
			p.Match(FusionTablesSqlParserK_SELECT)
		}

	case FusionTablesSqlParserK_SET:
		p.EnterOuterAlt(localctx, 47)
		{
			p.SetState(473)
			p.Match(FusionTablesSqlParserK_SET)
		}

	case FusionTablesSqlParserK_STARTS:
		p.EnterOuterAlt(localctx, 48)
		{
			p.SetState(474)
			p.Match(FusionTablesSqlParserK_STARTS)
		}

	case FusionTablesSqlParserK_ST_DISTANCE:
		p.EnterOuterAlt(localctx, 49)
		{
			p.SetState(475)
			p.Match(FusionTablesSqlParserK_ST_DISTANCE)
		}

	case FusionTablesSqlParserK_ST_INTERSECTS:
		p.EnterOuterAlt(localctx, 50)
		{
			p.SetState(476)
			p.Match(FusionTablesSqlParserK_ST_INTERSECTS)
		}

	case FusionTablesSqlParserK_SUM:
		p.EnterOuterAlt(localctx, 51)
		{
			p.SetState(477)
			p.Match(FusionTablesSqlParserK_SUM)
		}

	case FusionTablesSqlParserK_TABLE:
		p.EnterOuterAlt(localctx, 52)
		{
			p.SetState(478)
			p.Match(FusionTablesSqlParserK_TABLE)
		}

	case FusionTablesSqlParserK_TO:
		p.EnterOuterAlt(localctx, 53)
		{
			p.SetState(479)
			p.Match(FusionTablesSqlParserK_TO)
		}

	case FusionTablesSqlParserK_UPDATE:
		p.EnterOuterAlt(localctx, 54)
		{
			p.SetState(480)
			p.Match(FusionTablesSqlParserK_UPDATE)
		}

	case FusionTablesSqlParserK_VALUES:
		p.EnterOuterAlt(localctx, 55)
		{
			p.SetState(481)
			p.Match(FusionTablesSqlParserK_VALUES)
		}

	case FusionTablesSqlParserK_VIEW:
		p.EnterOuterAlt(localctx, 56)
		{
			p.SetState(482)
			p.Match(FusionTablesSqlParserK_VIEW)
		}

	case FusionTablesSqlParserK_WHERE:
		p.EnterOuterAlt(localctx, 57)
		{
			p.SetState(483)
			p.Match(FusionTablesSqlParserK_WHERE)
		}

	case FusionTablesSqlParserK_WITH:
		p.EnterOuterAlt(localctx, 58)
		{
			p.SetState(484)
			p.Match(FusionTablesSqlParserK_WITH)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserLT, 0)
}

func (s *OperatorContext) LT_EQ() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserLT_EQ, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserGT, 0)
}

func (s *OperatorContext) GT_EQ() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserGT_EQ, 0)
}

func (s *OperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserEQ, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *FusionTablesSqlParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, FusionTablesSqlParserRULE_operator)
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
		p.SetState(487)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-65)&-(0x1f+1)) == 0 && ((1<<uint((_la-65)))&((1<<(FusionTablesSqlParserLT_EQ-65))|(1<<(FusionTablesSqlParserGT_EQ-65))|(1<<(FusionTablesSqlParserGT-65))|(1<<(FusionTablesSqlParserEQ-65))|(1<<(FusionTablesSqlParserLT-65)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) Numeric_literal() INumeric_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumeric_literalContext)
}

func (s *LiteralContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *FusionTablesSqlParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, FusionTablesSqlParserRULE_literal)

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

	p.SetState(491)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FusionTablesSqlParserNUMERIC_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(489)
			p.Numeric_literal()
		}

	case FusionTablesSqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(490)
			p.String_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IError_messageContext is an interface to support dynamic dispatch.
type IError_messageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsError_messageContext differentiates from other interfaces.
	IsError_messageContext()
}

type Error_messageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyError_messageContext() *Error_messageContext {
	var p = new(Error_messageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_error_message
	return p
}

func (*Error_messageContext) IsError_messageContext() {}

func NewError_messageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Error_messageContext {
	var p = new(Error_messageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_error_message

	return p
}

func (s *Error_messageContext) GetParser() antlr.Parser { return s.parser }

func (s *Error_messageContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Error_messageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Error_messageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Error_messageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterError_message(s)
	}
}

func (s *Error_messageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitError_message(s)
	}
}

func (p *FusionTablesSqlParser) Error_message() (localctx IError_messageContext) {
	localctx = NewError_messageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, FusionTablesSqlParserRULE_error_message)

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
		p.SetState(493)
		p.String_literal()
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *FusionTablesSqlParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, FusionTablesSqlParserRULE_identifier)

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
		p.SetState(495)
		p.String_literal()
	}

	return localctx
}

// IColumn_aliasContext is an interface to support dynamic dispatch.
type IColumn_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_aliasContext differentiates from other interfaces.
	IsColumn_aliasContext()
}

type Column_aliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_aliasContext() *Column_aliasContext {
	var p = new(Column_aliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_column_alias
	return p
}

func (*Column_aliasContext) IsColumn_aliasContext() {}

func NewColumn_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_aliasContext {
	var p = new(Column_aliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_column_alias

	return p
}

func (s *Column_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_aliasContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Column_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_aliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterColumn_alias(s)
	}
}

func (s *Column_aliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitColumn_alias(s)
	}
}

func (p *FusionTablesSqlParser) Column_alias() (localctx IColumn_aliasContext) {
	localctx = NewColumn_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, FusionTablesSqlParserRULE_column_alias)

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
		p.SetState(497)
		p.Identifier()
	}

	return localctx
}

// ITable_nameContext is an interface to support dynamic dispatch.
type ITable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_nameContext differentiates from other interfaces.
	IsTable_nameContext()
}

type Table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_table_name
	return p
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterTable_name(s)
	}
}

func (s *Table_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitTable_name(s)
	}
}

func (p *FusionTablesSqlParser) Table_name() (localctx ITable_nameContext) {
	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, FusionTablesSqlParserRULE_table_name)

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
		p.SetState(499)
		p.Identifier()
	}

	return localctx
}

// IColumn_nameContext is an interface to support dynamic dispatch.
type IColumn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_nameContext differentiates from other interfaces.
	IsColumn_nameContext()
}

type Column_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_nameContext() *Column_nameContext {
	var p = new(Column_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_column_name
	return p
}

func (*Column_nameContext) IsColumn_nameContext() {}

func NewColumn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_nameContext {
	var p = new(Column_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_column_name

	return p
}

func (s *Column_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_nameContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Column_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterColumn_name(s)
	}
}

func (s *Column_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitColumn_name(s)
	}
}

func (p *FusionTablesSqlParser) Column_name() (localctx IColumn_nameContext) {
	localctx = NewColumn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, FusionTablesSqlParserRULE_column_name)

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
		p.SetState(501)
		p.Identifier()
	}

	return localctx
}

// INew_table_nameContext is an interface to support dynamic dispatch.
type INew_table_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNew_table_nameContext differentiates from other interfaces.
	IsNew_table_nameContext()
}

type New_table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNew_table_nameContext() *New_table_nameContext {
	var p = new(New_table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_new_table_name
	return p
}

func (*New_table_nameContext) IsNew_table_nameContext() {}

func NewNew_table_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *New_table_nameContext {
	var p = new(New_table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_new_table_name

	return p
}

func (s *New_table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *New_table_nameContext) Table_name() ITable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *New_table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *New_table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *New_table_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterNew_table_name(s)
	}
}

func (s *New_table_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitNew_table_name(s)
	}
}

func (p *FusionTablesSqlParser) New_table_name() (localctx INew_table_nameContext) {
	localctx = NewNew_table_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, FusionTablesSqlParserRULE_new_table_name)

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
		p.SetState(503)
		p.Table_name()
	}

	return localctx
}

// IView_nameContext is an interface to support dynamic dispatch.
type IView_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsView_nameContext differentiates from other interfaces.
	IsView_nameContext()
}

type View_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyView_nameContext() *View_nameContext {
	var p = new(View_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_view_name
	return p
}

func (*View_nameContext) IsView_nameContext() {}

func NewView_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *View_nameContext {
	var p = new(View_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_view_name

	return p
}

func (s *View_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *View_nameContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *View_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *View_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *View_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterView_name(s)
	}
}

func (s *View_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitView_name(s)
	}
}

func (p *FusionTablesSqlParser) View_name() (localctx IView_nameContext) {
	localctx = NewView_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, FusionTablesSqlParserRULE_view_name)

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
		p.SetState(505)
		p.Identifier()
	}

	return localctx
}

// ITable_aliasContext is an interface to support dynamic dispatch.
type ITable_aliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_aliasContext differentiates from other interfaces.
	IsTable_aliasContext()
}

type Table_aliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_aliasContext() *Table_aliasContext {
	var p = new(Table_aliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_table_alias
	return p
}

func (*Table_aliasContext) IsTable_aliasContext() {}

func NewTable_aliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_aliasContext {
	var p = new(Table_aliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_table_alias

	return p
}

func (s *Table_aliasContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_aliasContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Table_aliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_aliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_aliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterTable_alias(s)
	}
}

func (s *Table_aliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitTable_alias(s)
	}
}

func (p *FusionTablesSqlParser) Table_alias() (localctx ITable_aliasContext) {
	localctx = NewTable_aliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, FusionTablesSqlParserRULE_table_alias)

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
		p.SetState(507)
		p.Identifier()
	}

	return localctx
}

// INumeric_literalContext is an interface to support dynamic dispatch.
type INumeric_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumeric_literalContext differentiates from other interfaces.
	IsNumeric_literalContext()
}

type Numeric_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumeric_literalContext() *Numeric_literalContext {
	var p = new(Numeric_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_numeric_literal
	return p
}

func (*Numeric_literalContext) IsNumeric_literalContext() {}

func NewNumeric_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numeric_literalContext {
	var p = new(Numeric_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_numeric_literal

	return p
}

func (s *Numeric_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Numeric_literalContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserNUMERIC_LITERAL, 0)
}

func (s *Numeric_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numeric_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numeric_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterNumeric_literal(s)
	}
}

func (s *Numeric_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitNumeric_literal(s)
	}
}

func (p *FusionTablesSqlParser) Numeric_literal() (localctx INumeric_literalContext) {
	localctx = NewNumeric_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, FusionTablesSqlParserRULE_numeric_literal)

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
		p.SetState(509)
		p.Match(FusionTablesSqlParserNUMERIC_LITERAL)
	}

	return localctx
}

// IString_literalContext is an interface to support dynamic dispatch.
type IString_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_literalContext differentiates from other interfaces.
	IsString_literalContext()
}

type String_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_literalContext() *String_literalContext {
	var p = new(String_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FusionTablesSqlParserRULE_string_literal
	return p
}

func (*String_literalContext) IsString_literalContext() {}

func NewString_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_literalContext {
	var p = new(String_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FusionTablesSqlParserRULE_string_literal

	return p
}

func (s *String_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *String_literalContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(FusionTablesSqlParserSTRING_LITERAL, 0)
}

func (s *String_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.EnterString_literal(s)
	}
}

func (s *String_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FusionTablesSqlListener); ok {
		listenerT.ExitString_literal(s)
	}
}

func (p *FusionTablesSqlParser) String_literal() (localctx IString_literalContext) {
	localctx = NewString_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, FusionTablesSqlParserRULE_string_literal)

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
		p.SetState(511)
		p.Match(FusionTablesSqlParserSTRING_LITERAL)
	}

	return localctx
}
