// Generated from STParser.g4 by ANTLR 4.7.

package stringtemplate // STParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 40, 338,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3,
	57, 10, 3, 12, 3, 14, 3, 60, 11, 3, 3, 4, 3, 4, 5, 4, 64, 10, 4, 3, 5,
	3, 5, 6, 5, 68, 10, 5, 13, 5, 14, 5, 69, 5, 5, 72, 10, 5, 3, 6, 3, 6, 5,
	6, 76, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 82, 10, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	7, 9, 99, 10, 9, 12, 9, 14, 9, 102, 11, 9, 3, 9, 5, 9, 105, 10, 9, 3, 9,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 125, 10, 10, 12, 10,
	14, 10, 128, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 134, 10, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 143, 10, 11, 12, 11,
	14, 11, 146, 11, 11, 3, 12, 3, 12, 3, 12, 7, 12, 151, 10, 12, 12, 12, 14,
	12, 154, 11, 12, 3, 13, 3, 13, 3, 13, 5, 13, 159, 10, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 169, 10, 14, 12, 14, 14,
	14, 172, 11, 14, 3, 15, 3, 15, 3, 15, 7, 15, 177, 10, 15, 12, 15, 14, 15,
	180, 11, 15, 3, 16, 3, 16, 3, 16, 5, 16, 185, 10, 16, 3, 17, 3, 17, 3,
	17, 5, 17, 190, 10, 17, 3, 18, 3, 18, 3, 18, 6, 18, 195, 10, 18, 13, 18,
	14, 18, 196, 3, 18, 3, 18, 3, 18, 5, 18, 202, 10, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 7, 18, 208, 10, 18, 12, 18, 14, 18, 211, 11, 18, 7, 18, 213,
	10, 18, 12, 18, 14, 18, 216, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 7, 19, 226, 10, 19, 12, 19, 14, 19, 229, 11, 19, 3,
	20, 3, 20, 3, 20, 5, 20, 234, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 5, 20, 243, 10, 20, 3, 20, 3, 20, 5, 20, 247, 10, 20, 3,
	21, 3, 21, 3, 21, 5, 21, 252, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 5, 21, 260, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 266, 10,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 5, 21, 280, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22,
	297, 10, 22, 3, 22, 5, 22, 300, 10, 22, 5, 22, 302, 10, 22, 3, 23, 3, 23,
	5, 23, 306, 10, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 314,
	10, 24, 12, 24, 14, 24, 317, 11, 24, 3, 24, 3, 24, 5, 24, 321, 10, 24,
	3, 24, 5, 24, 324, 10, 24, 3, 25, 3, 25, 3, 25, 7, 25, 329, 10, 25, 12,
	25, 14, 25, 332, 11, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 2, 2, 27, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 2, 2, 2, 363, 2, 52, 3, 2, 2, 2, 4, 58, 3, 2, 2, 2,
	6, 63, 3, 2, 2, 2, 8, 71, 3, 2, 2, 2, 10, 75, 3, 2, 2, 2, 12, 77, 3, 2,
	2, 2, 14, 85, 3, 2, 2, 2, 16, 94, 3, 2, 2, 2, 18, 109, 3, 2, 2, 2, 20,
	139, 3, 2, 2, 2, 22, 147, 3, 2, 2, 2, 24, 158, 3, 2, 2, 2, 26, 160, 3,
	2, 2, 2, 28, 173, 3, 2, 2, 2, 30, 181, 3, 2, 2, 2, 32, 186, 3, 2, 2, 2,
	34, 191, 3, 2, 2, 2, 36, 217, 3, 2, 2, 2, 38, 246, 3, 2, 2, 2, 40, 279,
	3, 2, 2, 2, 42, 301, 3, 2, 2, 2, 44, 303, 3, 2, 2, 2, 46, 323, 3, 2, 2,
	2, 48, 325, 3, 2, 2, 2, 50, 333, 3, 2, 2, 2, 52, 53, 5, 4, 3, 2, 53, 54,
	7, 2, 2, 3, 54, 3, 3, 2, 2, 2, 55, 57, 5, 6, 4, 2, 56, 55, 3, 2, 2, 2,
	57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 5, 3, 2,
	2, 2, 60, 58, 3, 2, 2, 2, 61, 64, 5, 8, 5, 2, 62, 64, 5, 10, 6, 2, 63,
	61, 3, 2, 2, 2, 63, 62, 3, 2, 2, 2, 64, 7, 3, 2, 2, 2, 65, 72, 5, 12, 7,
	2, 66, 68, 7, 12, 2, 2, 67, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 67,
	3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 72, 3, 2, 2, 2, 71, 65, 3, 2, 2, 2,
	71, 67, 3, 2, 2, 2, 72, 9, 3, 2, 2, 2, 73, 76, 5, 18, 10, 2, 74, 76, 5,
	14, 8, 2, 75, 73, 3, 2, 2, 2, 75, 74, 3, 2, 2, 2, 76, 11, 3, 2, 2, 2, 77,
	78, 7, 10, 2, 2, 78, 81, 5, 34, 18, 2, 79, 80, 7, 29, 2, 2, 80, 82, 5,
	28, 15, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2,
	83, 84, 7, 14, 2, 2, 84, 13, 3, 2, 2, 2, 85, 86, 7, 10, 2, 2, 86, 87, 7,
	24, 2, 2, 87, 88, 7, 39, 2, 2, 88, 89, 7, 14, 2, 2, 89, 90, 5, 4, 3, 2,
	90, 91, 7, 10, 2, 2, 91, 92, 7, 21, 2, 2, 92, 93, 7, 14, 2, 2, 93, 15,
	3, 2, 2, 2, 94, 104, 7, 13, 2, 2, 95, 100, 7, 39, 2, 2, 96, 97, 7, 27,
	2, 2, 97, 99, 7, 39, 2, 2, 98, 96, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100,
	98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 103, 3, 2, 2, 2, 102, 100, 3,
	2, 2, 2, 103, 105, 7, 40, 2, 2, 104, 95, 3, 2, 2, 2, 104, 105, 3, 2, 2,
	2, 105, 106, 3, 2, 2, 2, 106, 107, 5, 4, 3, 2, 107, 108, 7, 11, 2, 2, 108,
	17, 3, 2, 2, 2, 109, 110, 7, 10, 2, 2, 110, 111, 7, 16, 2, 2, 111, 112,
	7, 32, 2, 2, 112, 113, 5, 20, 11, 2, 113, 114, 7, 33, 2, 2, 114, 115, 7,
	14, 2, 2, 115, 126, 5, 4, 3, 2, 116, 117, 7, 10, 2, 2, 117, 118, 7, 17,
	2, 2, 118, 119, 7, 32, 2, 2, 119, 120, 5, 20, 11, 2, 120, 121, 7, 33, 2,
	2, 121, 122, 7, 14, 2, 2, 122, 123, 5, 4, 3, 2, 123, 125, 3, 2, 2, 2, 124,
	116, 3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127,
	3, 2, 2, 2, 127, 133, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 130, 7, 10,
	2, 2, 130, 131, 7, 18, 2, 2, 131, 132, 7, 14, 2, 2, 132, 134, 5, 4, 3,
	2, 133, 129, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135,
	136, 7, 10, 2, 2, 136, 137, 7, 19, 2, 2, 137, 138, 7, 14, 2, 2, 138, 19,
	3, 2, 2, 2, 139, 144, 5, 22, 12, 2, 140, 141, 7, 31, 2, 2, 141, 143, 5,
	22, 12, 2, 142, 140, 3, 2, 2, 2, 143, 146, 3, 2, 2, 2, 144, 142, 3, 2,
	2, 2, 144, 145, 3, 2, 2, 2, 145, 21, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2,
	147, 152, 5, 24, 13, 2, 148, 149, 7, 30, 2, 2, 149, 151, 5, 24, 13, 2,
	150, 148, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152,
	153, 3, 2, 2, 2, 153, 23, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 155, 156, 7,
	37, 2, 2, 156, 159, 5, 24, 13, 2, 157, 159, 5, 36, 19, 2, 158, 155, 3,
	2, 2, 2, 158, 157, 3, 2, 2, 2, 159, 25, 3, 2, 2, 2, 160, 170, 7, 39, 2,
	2, 161, 162, 7, 26, 2, 2, 162, 169, 7, 39, 2, 2, 163, 164, 7, 26, 2, 2,
	164, 165, 7, 32, 2, 2, 165, 166, 5, 34, 18, 2, 166, 167, 7, 33, 2, 2, 167,
	169, 3, 2, 2, 2, 168, 161, 3, 2, 2, 2, 168, 163, 3, 2, 2, 2, 169, 172,
	3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 27, 3, 2,
	2, 2, 172, 170, 3, 2, 2, 2, 173, 178, 5, 30, 16, 2, 174, 175, 7, 27, 2,
	2, 175, 177, 5, 30, 16, 2, 176, 174, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2,
	178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 29, 3, 2, 2, 2, 180, 178,
	3, 2, 2, 2, 181, 184, 7, 39, 2, 2, 182, 183, 7, 36, 2, 2, 183, 185, 5,
	32, 17, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 31, 3, 2, 2,
	2, 186, 189, 5, 36, 19, 2, 187, 188, 7, 28, 2, 2, 188, 190, 5, 38, 20,
	2, 189, 187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 33, 3, 2, 2, 2, 191,
	201, 5, 36, 19, 2, 192, 193, 7, 27, 2, 2, 193, 195, 5, 36, 19, 2, 194,
	192, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 197,
	3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 199, 7, 28, 2, 2, 199, 200, 5, 38,
	20, 2, 200, 202, 3, 2, 2, 2, 201, 194, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2,
	202, 214, 3, 2, 2, 2, 203, 204, 7, 28, 2, 2, 204, 209, 5, 38, 20, 2, 205,
	206, 7, 27, 2, 2, 206, 208, 5, 38, 20, 2, 207, 205, 3, 2, 2, 2, 208, 211,
	3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 213, 3, 2,
	2, 2, 211, 209, 3, 2, 2, 2, 212, 203, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2,
	214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 35, 3, 2, 2, 2, 216, 214,
	3, 2, 2, 2, 217, 227, 5, 40, 21, 2, 218, 219, 7, 26, 2, 2, 219, 226, 7,
	39, 2, 2, 220, 221, 7, 26, 2, 2, 221, 222, 7, 32, 2, 2, 222, 223, 5, 34,
	18, 2, 223, 224, 7, 33, 2, 2, 224, 226, 3, 2, 2, 2, 225, 218, 3, 2, 2,
	2, 225, 220, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227,
	228, 3, 2, 2, 2, 228, 37, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 231, 7,
	39, 2, 2, 231, 233, 7, 32, 2, 2, 232, 234, 5, 46, 24, 2, 233, 232, 3, 2,
	2, 2, 233, 234, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 247, 7, 33, 2, 2,
	236, 247, 5, 16, 9, 2, 237, 238, 7, 32, 2, 2, 238, 239, 5, 34, 18, 2, 239,
	240, 7, 33, 2, 2, 240, 242, 7, 32, 2, 2, 241, 243, 5, 48, 25, 2, 242, 241,
	3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 245, 7, 33,
	2, 2, 245, 247, 3, 2, 2, 2, 246, 230, 3, 2, 2, 2, 246, 236, 3, 2, 2, 2,
	246, 237, 3, 2, 2, 2, 247, 39, 3, 2, 2, 2, 248, 249, 7, 39, 2, 2, 249,
	251, 7, 32, 2, 2, 250, 252, 5, 34, 18, 2, 251, 250, 3, 2, 2, 2, 251, 252,
	3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 280, 7, 33, 2, 2, 254, 255, 7, 20,
	2, 2, 255, 256, 7, 26, 2, 2, 256, 257, 7, 39, 2, 2, 257, 259, 7, 32, 2,
	2, 258, 260, 5, 46, 24, 2, 259, 258, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2,
	260, 261, 3, 2, 2, 2, 261, 280, 7, 33, 2, 2, 262, 263, 7, 39, 2, 2, 263,
	265, 7, 32, 2, 2, 264, 266, 5, 46, 24, 2, 265, 264, 3, 2, 2, 2, 265, 266,
	3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 280, 7, 33, 2, 2, 268, 269, 7, 24,
	2, 2, 269, 270, 7, 20, 2, 2, 270, 271, 7, 26, 2, 2, 271, 272, 7, 39, 2,
	2, 272, 273, 7, 32, 2, 2, 273, 280, 7, 33, 2, 2, 274, 275, 7, 24, 2, 2,
	275, 276, 7, 39, 2, 2, 276, 277, 7, 32, 2, 2, 277, 280, 7, 33, 2, 2, 278,
	280, 5, 42, 22, 2, 279, 248, 3, 2, 2, 2, 279, 254, 3, 2, 2, 2, 279, 262,
	3, 2, 2, 2, 279, 268, 3, 2, 2, 2, 279, 274, 3, 2, 2, 2, 279, 278, 3, 2,
	2, 2, 280, 41, 3, 2, 2, 2, 281, 302, 7, 39, 2, 2, 282, 302, 7, 15, 2, 2,
	283, 302, 7, 22, 2, 2, 284, 302, 7, 23, 2, 2, 285, 302, 5, 16, 9, 2, 286,
	302, 5, 44, 23, 2, 287, 288, 7, 32, 2, 2, 288, 289, 5, 20, 11, 2, 289,
	290, 7, 33, 2, 2, 290, 302, 3, 2, 2, 2, 291, 292, 7, 32, 2, 2, 292, 293,
	5, 34, 18, 2, 293, 299, 7, 33, 2, 2, 294, 296, 7, 32, 2, 2, 295, 297, 5,
	48, 25, 2, 296, 295, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298, 3, 2,
	2, 2, 298, 300, 7, 33, 2, 2, 299, 294, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2,
	300, 302, 3, 2, 2, 2, 301, 281, 3, 2, 2, 2, 301, 282, 3, 2, 2, 2, 301,
	283, 3, 2, 2, 2, 301, 284, 3, 2, 2, 2, 301, 285, 3, 2, 2, 2, 301, 286,
	3, 2, 2, 2, 301, 287, 3, 2, 2, 2, 301, 291, 3, 2, 2, 2, 302, 43, 3, 2,
	2, 2, 303, 305, 7, 34, 2, 2, 304, 306, 5, 48, 25, 2, 305, 304, 3, 2, 2,
	2, 305, 306, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 308, 7, 35, 2, 2, 308,
	45, 3, 2, 2, 2, 309, 324, 5, 48, 25, 2, 310, 315, 5, 50, 26, 2, 311, 312,
	7, 27, 2, 2, 312, 314, 5, 50, 26, 2, 313, 311, 3, 2, 2, 2, 314, 317, 3,
	2, 2, 2, 315, 313, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 320, 3, 2, 2,
	2, 317, 315, 3, 2, 2, 2, 318, 319, 7, 27, 2, 2, 319, 321, 7, 25, 2, 2,
	320, 318, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321, 324, 3, 2, 2, 2, 322,
	324, 7, 25, 2, 2, 323, 309, 3, 2, 2, 2, 323, 310, 3, 2, 2, 2, 323, 322,
	3, 2, 2, 2, 324, 47, 3, 2, 2, 2, 325, 330, 5, 32, 17, 2, 326, 327, 7, 27,
	2, 2, 327, 329, 5, 32, 17, 2, 328, 326, 3, 2, 2, 2, 329, 332, 3, 2, 2,
	2, 330, 328, 3, 2, 2, 2, 330, 331, 3, 2, 2, 2, 331, 49, 3, 2, 2, 2, 332,
	330, 3, 2, 2, 2, 333, 334, 7, 39, 2, 2, 334, 335, 7, 36, 2, 2, 335, 336,
	5, 32, 17, 2, 336, 51, 3, 2, 2, 2, 41, 58, 63, 69, 71, 75, 81, 100, 104,
	126, 133, 144, 152, 158, 168, 170, 178, 184, 189, 196, 201, 209, 214, 225,
	227, 233, 242, 246, 251, 259, 265, 279, 296, 299, 301, 305, 315, 320, 323,
	330,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "'if'", "'elseif'",
	"'else'", "'endif'", "'super'", "'@end'",
}
var symbolicNames = []string{
	"", "DOC_COMMENT", "BLOCK_COMMENT", "LINE_COMMENT", "TMPL_COMMENT", "HORZ_WS",
	"VERT_WS", "ESCAPE", "LDELIM", "RBRACE", "TEXT", "LBRACE", "RDELIM", "STRING",
	"IF", "ELSEIF", "ELSE", "ENDIF", "SUPER", "END", "TRUE", "FALSE", "AT",
	"ELLIPSIS", "DOT", "COMMA", "COLON", "SEMI", "AND", "OR", "LPAREN", "RPAREN",
	"LBRACK", "RBRACK", "EQUALS", "BANG", "ERR_CHAR", "ID", "PIPE",
}

var ruleNames = []string{
	"template", "elements", "element", "singleElement", "compoundElement",
	"exprTag", "region", "subtemplate", "ifstat", "conditional", "andConditional",
	"notConditional", "notConditionalExpr", "exprOptions", "option", "expr",
	"mapExpr", "memberExpr", "mapTemplateRef", "includeExpr", "primary", "list",
	"args", "argExprList", "namedArg",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type STParser struct {
	*antlr.BaseParser
}

func NewSTParser(input antlr.TokenStream) *STParser {
	this := new(STParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "STParser.g4"

	return this
}

// STParser tokens.
const (
	STParserEOF           = antlr.TokenEOF
	STParserDOC_COMMENT   = 1
	STParserBLOCK_COMMENT = 2
	STParserLINE_COMMENT  = 3
	STParserTMPL_COMMENT  = 4
	STParserHORZ_WS       = 5
	STParserVERT_WS       = 6
	STParserESCAPE        = 7
	STParserLDELIM        = 8
	STParserRBRACE        = 9
	STParserTEXT          = 10
	STParserLBRACE        = 11
	STParserRDELIM        = 12
	STParserSTRING        = 13
	STParserIF            = 14
	STParserELSEIF        = 15
	STParserELSE          = 16
	STParserENDIF         = 17
	STParserSUPER         = 18
	STParserEND           = 19
	STParserTRUE          = 20
	STParserFALSE         = 21
	STParserAT            = 22
	STParserELLIPSIS      = 23
	STParserDOT           = 24
	STParserCOMMA         = 25
	STParserCOLON         = 26
	STParserSEMI          = 27
	STParserAND           = 28
	STParserOR            = 29
	STParserLPAREN        = 30
	STParserRPAREN        = 31
	STParserLBRACK        = 32
	STParserRBRACK        = 33
	STParserEQUALS        = 34
	STParserBANG          = 35
	STParserERR_CHAR      = 36
	STParserID            = 37
	STParserPIPE          = 38
)

// STParser rules.
const (
	STParserRULE_template           = 0
	STParserRULE_elements           = 1
	STParserRULE_element            = 2
	STParserRULE_singleElement      = 3
	STParserRULE_compoundElement    = 4
	STParserRULE_exprTag            = 5
	STParserRULE_region             = 6
	STParserRULE_subtemplate        = 7
	STParserRULE_ifstat             = 8
	STParserRULE_conditional        = 9
	STParserRULE_andConditional     = 10
	STParserRULE_notConditional     = 11
	STParserRULE_notConditionalExpr = 12
	STParserRULE_exprOptions        = 13
	STParserRULE_option             = 14
	STParserRULE_expr               = 15
	STParserRULE_mapExpr            = 16
	STParserRULE_memberExpr         = 17
	STParserRULE_mapTemplateRef     = 18
	STParserRULE_includeExpr        = 19
	STParserRULE_primary            = 20
	STParserRULE_list               = 21
	STParserRULE_args               = 22
	STParserRULE_argExprList        = 23
	STParserRULE_namedArg           = 24
)

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
	p.RuleIndex = STParserRULE_template
	return p
}

func (*TemplateContext) IsTemplateContext() {}

func NewTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateContext {
	var p = new(TemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_template

	return p
}

func (s *TemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateContext) Elements() IElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementsContext)
}

func (s *TemplateContext) EOF() antlr.TerminalNode {
	return s.GetToken(STParserEOF, 0)
}

func (s *TemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterTemplate(s)
	}
}

func (s *TemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitTemplate(s)
	}
}

func (s *TemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Template() (localctx ITemplateContext) {
	localctx = NewTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, STParserRULE_template)

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
		p.SetState(50)
		p.Elements()
	}
	{
		p.SetState(51)
		p.Match(STParserEOF)
	}

	return localctx
}

// IElementsContext is an interface to support dynamic dispatch.
type IElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementsContext differentiates from other interfaces.
	IsElementsContext()
}

type ElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementsContext() *ElementsContext {
	var p = new(ElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_elements
	return p
}

func (*ElementsContext) IsElementsContext() {}

func NewElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementsContext {
	var p = new(ElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_elements

	return p
}

func (s *ElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementsContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *ElementsContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterElements(s)
	}
}

func (s *ElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitElements(s)
	}
}

func (s *ElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Elements() (localctx IElementsContext) {
	localctx = NewElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, STParserRULE_elements)

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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(53)
				p.Element()
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) SingleElement() ISingleElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleElementContext)
}

func (s *ElementContext) CompoundElement() ICompoundElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundElementContext)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitElement(s)
	}
}

func (s *ElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, STParserRULE_element)

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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.SingleElement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.CompoundElement()
		}

	}

	return localctx
}

// ISingleElementContext is an interface to support dynamic dispatch.
type ISingleElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleElementContext differentiates from other interfaces.
	IsSingleElementContext()
}

type SingleElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleElementContext() *SingleElementContext {
	var p = new(SingleElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_singleElement
	return p
}

func (*SingleElementContext) IsSingleElementContext() {}

func NewSingleElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleElementContext {
	var p = new(SingleElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_singleElement

	return p
}

func (s *SingleElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleElementContext) ExprTag() IExprTagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprTagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprTagContext)
}

func (s *SingleElementContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(STParserTEXT)
}

func (s *SingleElementContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(STParserTEXT, i)
}

func (s *SingleElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterSingleElement(s)
	}
}

func (s *SingleElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitSingleElement(s)
	}
}

func (s *SingleElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitSingleElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) SingleElement() (localctx ISingleElementContext) {
	localctx = NewSingleElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, STParserRULE_singleElement)

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

	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case STParserLDELIM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.ExprTag()
		}

	case STParserTEXT:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(64)
					p.Match(STParserTEXT)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICompoundElementContext is an interface to support dynamic dispatch.
type ICompoundElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompoundElementContext differentiates from other interfaces.
	IsCompoundElementContext()
}

type CompoundElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundElementContext() *CompoundElementContext {
	var p = new(CompoundElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_compoundElement
	return p
}

func (*CompoundElementContext) IsCompoundElementContext() {}

func NewCompoundElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundElementContext {
	var p = new(CompoundElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_compoundElement

	return p
}

func (s *CompoundElementContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundElementContext) Ifstat() IIfstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfstatContext)
}

func (s *CompoundElementContext) Region() IRegionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegionContext)
}

func (s *CompoundElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterCompoundElement(s)
	}
}

func (s *CompoundElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitCompoundElement(s)
	}
}

func (s *CompoundElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitCompoundElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) CompoundElement() (localctx ICompoundElementContext) {
	localctx = NewCompoundElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, STParserRULE_compoundElement)

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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Ifstat()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Region()
		}

	}

	return localctx
}

// IExprTagContext is an interface to support dynamic dispatch.
type IExprTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprTagContext differentiates from other interfaces.
	IsExprTagContext()
}

type ExprTagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprTagContext() *ExprTagContext {
	var p = new(ExprTagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_exprTag
	return p
}

func (*ExprTagContext) IsExprTagContext() {}

func NewExprTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprTagContext {
	var p = new(ExprTagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_exprTag

	return p
}

func (s *ExprTagContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprTagContext) LDELIM() antlr.TerminalNode {
	return s.GetToken(STParserLDELIM, 0)
}

func (s *ExprTagContext) MapExpr() IMapExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapExprContext)
}

func (s *ExprTagContext) RDELIM() antlr.TerminalNode {
	return s.GetToken(STParserRDELIM, 0)
}

func (s *ExprTagContext) SEMI() antlr.TerminalNode {
	return s.GetToken(STParserSEMI, 0)
}

func (s *ExprTagContext) ExprOptions() IExprOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprOptionsContext)
}

func (s *ExprTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterExprTag(s)
	}
}

func (s *ExprTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitExprTag(s)
	}
}

func (s *ExprTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitExprTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) ExprTag() (localctx IExprTagContext) {
	localctx = NewExprTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, STParserRULE_exprTag)
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
		p.SetState(75)
		p.Match(STParserLDELIM)
	}
	{
		p.SetState(76)
		p.MapExpr()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == STParserSEMI {
		{
			p.SetState(77)
			p.Match(STParserSEMI)
		}
		{
			p.SetState(78)
			p.ExprOptions()
		}

	}
	{
		p.SetState(81)
		p.Match(STParserRDELIM)
	}

	return localctx
}

// IRegionContext is an interface to support dynamic dispatch.
type IRegionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegionContext differentiates from other interfaces.
	IsRegionContext()
}

type RegionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegionContext() *RegionContext {
	var p = new(RegionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_region
	return p
}

func (*RegionContext) IsRegionContext() {}

func NewRegionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegionContext {
	var p = new(RegionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_region

	return p
}

func (s *RegionContext) GetParser() antlr.Parser { return s.parser }

func (s *RegionContext) AllLDELIM() []antlr.TerminalNode {
	return s.GetTokens(STParserLDELIM)
}

func (s *RegionContext) LDELIM(i int) antlr.TerminalNode {
	return s.GetToken(STParserLDELIM, i)
}

func (s *RegionContext) AT() antlr.TerminalNode {
	return s.GetToken(STParserAT, 0)
}

func (s *RegionContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
}

func (s *RegionContext) AllRDELIM() []antlr.TerminalNode {
	return s.GetTokens(STParserRDELIM)
}

func (s *RegionContext) RDELIM(i int) antlr.TerminalNode {
	return s.GetToken(STParserRDELIM, i)
}

func (s *RegionContext) Elements() IElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementsContext)
}

func (s *RegionContext) END() antlr.TerminalNode {
	return s.GetToken(STParserEND, 0)
}

func (s *RegionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterRegion(s)
	}
}

func (s *RegionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitRegion(s)
	}
}

func (s *RegionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitRegion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Region() (localctx IRegionContext) {
	localctx = NewRegionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, STParserRULE_region)

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
		p.SetState(83)
		p.Match(STParserLDELIM)
	}
	{
		p.SetState(84)
		p.Match(STParserAT)
	}
	{
		p.SetState(85)
		p.Match(STParserID)
	}
	{
		p.SetState(86)
		p.Match(STParserRDELIM)
	}
	{
		p.SetState(87)
		p.Elements()
	}
	{
		p.SetState(88)
		p.Match(STParserLDELIM)
	}
	{
		p.SetState(89)
		p.Match(STParserEND)
	}
	{
		p.SetState(90)
		p.Match(STParserRDELIM)
	}

	return localctx
}

// ISubtemplateContext is an interface to support dynamic dispatch.
type ISubtemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubtemplateContext differentiates from other interfaces.
	IsSubtemplateContext()
}

type SubtemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubtemplateContext() *SubtemplateContext {
	var p = new(SubtemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_subtemplate
	return p
}

func (*SubtemplateContext) IsSubtemplateContext() {}

func NewSubtemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubtemplateContext {
	var p = new(SubtemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_subtemplate

	return p
}

func (s *SubtemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *SubtemplateContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(STParserLBRACE, 0)
}

func (s *SubtemplateContext) Elements() IElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementsContext)
}

func (s *SubtemplateContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(STParserRBRACE, 0)
}

func (s *SubtemplateContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(STParserID)
}

func (s *SubtemplateContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(STParserID, i)
}

func (s *SubtemplateContext) PIPE() antlr.TerminalNode {
	return s.GetToken(STParserPIPE, 0)
}

func (s *SubtemplateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *SubtemplateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *SubtemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubtemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterSubtemplate(s)
	}
}

func (s *SubtemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitSubtemplate(s)
	}
}

func (s *SubtemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitSubtemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Subtemplate() (localctx ISubtemplateContext) {
	localctx = NewSubtemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, STParserRULE_subtemplate)
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
		p.SetState(92)
		p.Match(STParserLBRACE)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == STParserID {
		{
			p.SetState(93)
			p.Match(STParserID)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == STParserCOMMA {
			{
				p.SetState(94)
				p.Match(STParserCOMMA)
			}
			{
				p.SetState(95)
				p.Match(STParserID)
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(101)
			p.Match(STParserPIPE)
		}

	}
	{
		p.SetState(104)
		p.Elements()
	}
	{
		p.SetState(105)
		p.Match(STParserRBRACE)
	}

	return localctx
}

// IIfstatContext is an interface to support dynamic dispatch.
type IIfstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfstatContext differentiates from other interfaces.
	IsIfstatContext()
}

type IfstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstatContext() *IfstatContext {
	var p = new(IfstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_ifstat
	return p
}

func (*IfstatContext) IsIfstatContext() {}

func NewIfstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstatContext {
	var p = new(IfstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_ifstat

	return p
}

func (s *IfstatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstatContext) AllLDELIM() []antlr.TerminalNode {
	return s.GetTokens(STParserLDELIM)
}

func (s *IfstatContext) LDELIM(i int) antlr.TerminalNode {
	return s.GetToken(STParserLDELIM, i)
}

func (s *IfstatContext) IF() antlr.TerminalNode {
	return s.GetToken(STParserIF, 0)
}

func (s *IfstatContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(STParserLPAREN)
}

func (s *IfstatContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(STParserLPAREN, i)
}

func (s *IfstatContext) AllConditional() []IConditionalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionalContext)(nil)).Elem())
	var tst = make([]IConditionalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionalContext)
		}
	}

	return tst
}

func (s *IfstatContext) Conditional(i int) IConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *IfstatContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(STParserRPAREN)
}

func (s *IfstatContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(STParserRPAREN, i)
}

func (s *IfstatContext) AllRDELIM() []antlr.TerminalNode {
	return s.GetTokens(STParserRDELIM)
}

func (s *IfstatContext) RDELIM(i int) antlr.TerminalNode {
	return s.GetToken(STParserRDELIM, i)
}

func (s *IfstatContext) AllElements() []IElementsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementsContext)(nil)).Elem())
	var tst = make([]IElementsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementsContext)
		}
	}

	return tst
}

func (s *IfstatContext) Elements(i int) IElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementsContext)
}

func (s *IfstatContext) ENDIF() antlr.TerminalNode {
	return s.GetToken(STParserENDIF, 0)
}

func (s *IfstatContext) AllELSEIF() []antlr.TerminalNode {
	return s.GetTokens(STParserELSEIF)
}

func (s *IfstatContext) ELSEIF(i int) antlr.TerminalNode {
	return s.GetToken(STParserELSEIF, i)
}

func (s *IfstatContext) ELSE() antlr.TerminalNode {
	return s.GetToken(STParserELSE, 0)
}

func (s *IfstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterIfstat(s)
	}
}

func (s *IfstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitIfstat(s)
	}
}

func (s *IfstatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitIfstat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Ifstat() (localctx IIfstatContext) {
	localctx = NewIfstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, STParserRULE_ifstat)

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
		p.SetState(107)
		p.Match(STParserLDELIM)
	}
	{
		p.SetState(108)
		p.Match(STParserIF)
	}
	{
		p.SetState(109)
		p.Match(STParserLPAREN)
	}
	{
		p.SetState(110)
		p.Conditional()
	}
	{
		p.SetState(111)
		p.Match(STParserRPAREN)
	}
	{
		p.SetState(112)
		p.Match(STParserRDELIM)
	}
	{
		p.SetState(113)
		p.Elements()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(114)
				p.Match(STParserLDELIM)
			}
			{
				p.SetState(115)
				p.Match(STParserELSEIF)
			}
			{
				p.SetState(116)
				p.Match(STParserLPAREN)
			}
			{
				p.SetState(117)
				p.Conditional()
			}
			{
				p.SetState(118)
				p.Match(STParserRPAREN)
			}
			{
				p.SetState(119)
				p.Match(STParserRDELIM)
			}
			{
				p.SetState(120)
				p.Elements()
			}

		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(127)
			p.Match(STParserLDELIM)
		}
		{
			p.SetState(128)
			p.Match(STParserELSE)
		}
		{
			p.SetState(129)
			p.Match(STParserRDELIM)
		}
		{
			p.SetState(130)
			p.Elements()
		}

	}
	{
		p.SetState(133)
		p.Match(STParserLDELIM)
	}
	{
		p.SetState(134)
		p.Match(STParserENDIF)
	}
	{
		p.SetState(135)
		p.Match(STParserRDELIM)
	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) AllAndConditional() []IAndConditionalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAndConditionalContext)(nil)).Elem())
	var tst = make([]IAndConditionalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAndConditionalContext)
		}
	}

	return tst
}

func (s *ConditionalContext) AndConditional(i int) IAndConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndConditionalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAndConditionalContext)
}

func (s *ConditionalContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(STParserOR)
}

func (s *ConditionalContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(STParserOR, i)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (s *ConditionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitConditional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Conditional() (localctx IConditionalContext) {
	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, STParserRULE_conditional)
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
		p.SetState(137)
		p.AndConditional()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == STParserOR {
		{
			p.SetState(138)
			p.Match(STParserOR)
		}
		{
			p.SetState(139)
			p.AndConditional()
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAndConditionalContext is an interface to support dynamic dispatch.
type IAndConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndConditionalContext differentiates from other interfaces.
	IsAndConditionalContext()
}

type AndConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndConditionalContext() *AndConditionalContext {
	var p = new(AndConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_andConditional
	return p
}

func (*AndConditionalContext) IsAndConditionalContext() {}

func NewAndConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndConditionalContext {
	var p = new(AndConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_andConditional

	return p
}

func (s *AndConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *AndConditionalContext) AllNotConditional() []INotConditionalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INotConditionalContext)(nil)).Elem())
	var tst = make([]INotConditionalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INotConditionalContext)
		}
	}

	return tst
}

func (s *AndConditionalContext) NotConditional(i int) INotConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotConditionalContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INotConditionalContext)
}

func (s *AndConditionalContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(STParserAND)
}

func (s *AndConditionalContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(STParserAND, i)
}

func (s *AndConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterAndConditional(s)
	}
}

func (s *AndConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitAndConditional(s)
	}
}

func (s *AndConditionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitAndConditional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) AndConditional() (localctx IAndConditionalContext) {
	localctx = NewAndConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, STParserRULE_andConditional)
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
		p.SetState(145)
		p.NotConditional()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == STParserAND {
		{
			p.SetState(146)
			p.Match(STParserAND)
		}
		{
			p.SetState(147)
			p.NotConditional()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INotConditionalContext is an interface to support dynamic dispatch.
type INotConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotConditionalContext differentiates from other interfaces.
	IsNotConditionalContext()
}

type NotConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotConditionalContext() *NotConditionalContext {
	var p = new(NotConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_notConditional
	return p
}

func (*NotConditionalContext) IsNotConditionalContext() {}

func NewNotConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotConditionalContext {
	var p = new(NotConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_notConditional

	return p
}

func (s *NotConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *NotConditionalContext) BANG() antlr.TerminalNode {
	return s.GetToken(STParserBANG, 0)
}

func (s *NotConditionalContext) NotConditional() INotConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotConditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotConditionalContext)
}

func (s *NotConditionalContext) MemberExpr() IMemberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberExprContext)
}

func (s *NotConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterNotConditional(s)
	}
}

func (s *NotConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitNotConditional(s)
	}
}

func (s *NotConditionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitNotConditional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) NotConditional() (localctx INotConditionalContext) {
	localctx = NewNotConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, STParserRULE_notConditional)

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

	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case STParserBANG:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(STParserBANG)
		}
		{
			p.SetState(154)
			p.NotConditional()
		}

	case STParserLBRACE, STParserSTRING, STParserSUPER, STParserTRUE, STParserFALSE, STParserAT, STParserLPAREN, STParserLBRACK, STParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.MemberExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INotConditionalExprContext is an interface to support dynamic dispatch.
type INotConditionalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotConditionalExprContext differentiates from other interfaces.
	IsNotConditionalExprContext()
}

type NotConditionalExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotConditionalExprContext() *NotConditionalExprContext {
	var p = new(NotConditionalExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_notConditionalExpr
	return p
}

func (*NotConditionalExprContext) IsNotConditionalExprContext() {}

func NewNotConditionalExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotConditionalExprContext {
	var p = new(NotConditionalExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_notConditionalExpr

	return p
}

func (s *NotConditionalExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NotConditionalExprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(STParserID)
}

func (s *NotConditionalExprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(STParserID, i)
}

func (s *NotConditionalExprContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(STParserDOT)
}

func (s *NotConditionalExprContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(STParserDOT, i)
}

func (s *NotConditionalExprContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(STParserLPAREN)
}

func (s *NotConditionalExprContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(STParserLPAREN, i)
}

func (s *NotConditionalExprContext) AllMapExpr() []IMapExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMapExprContext)(nil)).Elem())
	var tst = make([]IMapExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMapExprContext)
		}
	}

	return tst
}

func (s *NotConditionalExprContext) MapExpr(i int) IMapExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMapExprContext)
}

func (s *NotConditionalExprContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(STParserRPAREN)
}

func (s *NotConditionalExprContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(STParserRPAREN, i)
}

func (s *NotConditionalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotConditionalExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotConditionalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterNotConditionalExpr(s)
	}
}

func (s *NotConditionalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitNotConditionalExpr(s)
	}
}

func (s *NotConditionalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitNotConditionalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) NotConditionalExpr() (localctx INotConditionalExprContext) {
	localctx = NewNotConditionalExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, STParserRULE_notConditionalExpr)
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
		p.SetState(158)
		p.Match(STParserID)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == STParserDOT {
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(159)
				p.Match(STParserDOT)
			}
			{
				p.SetState(160)
				p.Match(STParserID)
			}

		case 2:
			{
				p.SetState(161)
				p.Match(STParserDOT)
			}
			{
				p.SetState(162)
				p.Match(STParserLPAREN)
			}
			{
				p.SetState(163)
				p.MapExpr()
			}
			{
				p.SetState(164)
				p.Match(STParserRPAREN)
			}

		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprOptionsContext is an interface to support dynamic dispatch.
type IExprOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprOptionsContext differentiates from other interfaces.
	IsExprOptionsContext()
}

type ExprOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprOptionsContext() *ExprOptionsContext {
	var p = new(ExprOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_exprOptions
	return p
}

func (*ExprOptionsContext) IsExprOptionsContext() {}

func NewExprOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprOptionsContext {
	var p = new(ExprOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_exprOptions

	return p
}

func (s *ExprOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprOptionsContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *ExprOptionsContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *ExprOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *ExprOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *ExprOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterExprOptions(s)
	}
}

func (s *ExprOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitExprOptions(s)
	}
}

func (s *ExprOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitExprOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) ExprOptions() (localctx IExprOptionsContext) {
	localctx = NewExprOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, STParserRULE_exprOptions)
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
		p.SetState(171)
		p.Option()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(172)
			p.Match(STParserCOMMA)
		}
		{
			p.SetState(173)
			p.Option()
		}

		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_option
	return p
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
}

func (s *OptionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(STParserEQUALS, 0)
}

func (s *OptionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitOption(s)
	}
}

func (s *OptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, STParserRULE_option)
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
		p.SetState(179)
		p.Match(STParserID)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == STParserEQUALS {
		{
			p.SetState(180)
			p.Match(STParserEQUALS)
		}
		{
			p.SetState(181)
			p.Expr()
		}

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
	p.RuleIndex = STParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) MemberExpr() IMemberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberExprContext)
}

func (s *ExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(STParserCOLON, 0)
}

func (s *ExprContext) MapTemplateRef() IMapTemplateRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTemplateRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapTemplateRefContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, STParserRULE_expr)
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
		p.SetState(184)
		p.MemberExpr()
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == STParserCOLON {
		{
			p.SetState(185)
			p.Match(STParserCOLON)
		}
		{
			p.SetState(186)
			p.MapTemplateRef()
		}

	}

	return localctx
}

// IMapExprContext is an interface to support dynamic dispatch.
type IMapExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapExprContext differentiates from other interfaces.
	IsMapExprContext()
}

type MapExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapExprContext() *MapExprContext {
	var p = new(MapExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_mapExpr
	return p
}

func (*MapExprContext) IsMapExprContext() {}

func NewMapExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapExprContext {
	var p = new(MapExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_mapExpr

	return p
}

func (s *MapExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MapExprContext) AllMemberExpr() []IMemberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemberExprContext)(nil)).Elem())
	var tst = make([]IMemberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemberExprContext)
		}
	}

	return tst
}

func (s *MapExprContext) MemberExpr(i int) IMemberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemberExprContext)
}

func (s *MapExprContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(STParserCOLON)
}

func (s *MapExprContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOLON, i)
}

func (s *MapExprContext) AllMapTemplateRef() []IMapTemplateRefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMapTemplateRefContext)(nil)).Elem())
	var tst = make([]IMapTemplateRefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMapTemplateRefContext)
		}
	}

	return tst
}

func (s *MapExprContext) MapTemplateRef(i int) IMapTemplateRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTemplateRefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMapTemplateRefContext)
}

func (s *MapExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *MapExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *MapExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterMapExpr(s)
	}
}

func (s *MapExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitMapExpr(s)
	}
}

func (s *MapExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitMapExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) MapExpr() (localctx IMapExprContext) {
	localctx = NewMapExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, STParserRULE_mapExpr)
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
		p.SetState(189)
		p.MemberExpr()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == STParserCOMMA {
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == STParserCOMMA {
			{
				p.SetState(190)
				p.Match(STParserCOMMA)
			}
			{
				p.SetState(191)
				p.MemberExpr()
			}

			p.SetState(194)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(196)
			p.Match(STParserCOLON)
		}
		{
			p.SetState(197)
			p.MapTemplateRef()
		}

	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOLON {
		{
			p.SetState(201)
			p.Match(STParserCOLON)
		}
		{
			p.SetState(202)
			p.MapTemplateRef()
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == STParserCOMMA {
			{
				p.SetState(203)
				p.Match(STParserCOMMA)
			}
			{
				p.SetState(204)
				p.MapTemplateRef()
			}

			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMemberExprContext is an interface to support dynamic dispatch.
type IMemberExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberExprContext differentiates from other interfaces.
	IsMemberExprContext()
}

type MemberExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberExprContext() *MemberExprContext {
	var p = new(MemberExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_memberExpr
	return p
}

func (*MemberExprContext) IsMemberExprContext() {}

func NewMemberExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberExprContext {
	var p = new(MemberExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_memberExpr

	return p
}

func (s *MemberExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberExprContext) IncludeExpr() IIncludeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncludeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncludeExprContext)
}

func (s *MemberExprContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(STParserDOT)
}

func (s *MemberExprContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(STParserDOT, i)
}

func (s *MemberExprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(STParserID)
}

func (s *MemberExprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(STParserID, i)
}

func (s *MemberExprContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(STParserLPAREN)
}

func (s *MemberExprContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(STParserLPAREN, i)
}

func (s *MemberExprContext) AllMapExpr() []IMapExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMapExprContext)(nil)).Elem())
	var tst = make([]IMapExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMapExprContext)
		}
	}

	return tst
}

func (s *MemberExprContext) MapExpr(i int) IMapExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMapExprContext)
}

func (s *MemberExprContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(STParserRPAREN)
}

func (s *MemberExprContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(STParserRPAREN, i)
}

func (s *MemberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterMemberExpr(s)
	}
}

func (s *MemberExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitMemberExpr(s)
	}
}

func (s *MemberExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitMemberExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) MemberExpr() (localctx IMemberExprContext) {
	localctx = NewMemberExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, STParserRULE_memberExpr)
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
		p.SetState(215)
		p.IncludeExpr()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == STParserDOT {
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(216)
				p.Match(STParserDOT)
			}
			{
				p.SetState(217)
				p.Match(STParserID)
			}

		case 2:
			{
				p.SetState(218)
				p.Match(STParserDOT)
			}
			{
				p.SetState(219)
				p.Match(STParserLPAREN)
			}
			{
				p.SetState(220)
				p.MapExpr()
			}
			{
				p.SetState(221)
				p.Match(STParserRPAREN)
			}

		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMapTemplateRefContext is an interface to support dynamic dispatch.
type IMapTemplateRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapTemplateRefContext differentiates from other interfaces.
	IsMapTemplateRefContext()
}

type MapTemplateRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTemplateRefContext() *MapTemplateRefContext {
	var p = new(MapTemplateRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_mapTemplateRef
	return p
}

func (*MapTemplateRefContext) IsMapTemplateRefContext() {}

func NewMapTemplateRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTemplateRefContext {
	var p = new(MapTemplateRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_mapTemplateRef

	return p
}

func (s *MapTemplateRefContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTemplateRefContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
}

func (s *MapTemplateRefContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(STParserLPAREN)
}

func (s *MapTemplateRefContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(STParserLPAREN, i)
}

func (s *MapTemplateRefContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(STParserRPAREN)
}

func (s *MapTemplateRefContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(STParserRPAREN, i)
}

func (s *MapTemplateRefContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *MapTemplateRefContext) Subtemplate() ISubtemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubtemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubtemplateContext)
}

func (s *MapTemplateRefContext) MapExpr() IMapExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapExprContext)
}

func (s *MapTemplateRefContext) ArgExprList() IArgExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgExprListContext)
}

func (s *MapTemplateRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTemplateRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTemplateRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterMapTemplateRef(s)
	}
}

func (s *MapTemplateRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitMapTemplateRef(s)
	}
}

func (s *MapTemplateRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitMapTemplateRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) MapTemplateRef() (localctx IMapTemplateRefContext) {
	localctx = NewMapTemplateRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, STParserRULE_mapTemplateRef)
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

	p.SetState(244)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case STParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Match(STParserID)
		}
		{
			p.SetState(229)
			p.Match(STParserLPAREN)
		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(STParserLBRACE-11))|(1<<(STParserSTRING-11))|(1<<(STParserSUPER-11))|(1<<(STParserTRUE-11))|(1<<(STParserFALSE-11))|(1<<(STParserAT-11))|(1<<(STParserELLIPSIS-11))|(1<<(STParserLPAREN-11))|(1<<(STParserLBRACK-11))|(1<<(STParserID-11)))) != 0 {
			{
				p.SetState(230)
				p.Args()
			}

		}
		{
			p.SetState(233)
			p.Match(STParserRPAREN)
		}

	case STParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Subtemplate()
		}

	case STParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(235)
			p.Match(STParserLPAREN)
		}
		{
			p.SetState(236)
			p.MapExpr()
		}
		{
			p.SetState(237)
			p.Match(STParserRPAREN)
		}
		{
			p.SetState(238)
			p.Match(STParserLPAREN)
		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(STParserLBRACE-11))|(1<<(STParserSTRING-11))|(1<<(STParserSUPER-11))|(1<<(STParserTRUE-11))|(1<<(STParserFALSE-11))|(1<<(STParserAT-11))|(1<<(STParserLPAREN-11))|(1<<(STParserLBRACK-11))|(1<<(STParserID-11)))) != 0 {
			{
				p.SetState(239)
				p.ArgExprList()
			}

		}
		{
			p.SetState(242)
			p.Match(STParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIncludeExprContext is an interface to support dynamic dispatch.
type IIncludeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncludeExprContext differentiates from other interfaces.
	IsIncludeExprContext()
}

type IncludeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeExprContext() *IncludeExprContext {
	var p = new(IncludeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_includeExpr
	return p
}

func (*IncludeExprContext) IsIncludeExprContext() {}

func NewIncludeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeExprContext {
	var p = new(IncludeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_includeExpr

	return p
}

func (s *IncludeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeExprContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
}

func (s *IncludeExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(STParserLPAREN, 0)
}

func (s *IncludeExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(STParserRPAREN, 0)
}

func (s *IncludeExprContext) MapExpr() IMapExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapExprContext)
}

func (s *IncludeExprContext) SUPER() antlr.TerminalNode {
	return s.GetToken(STParserSUPER, 0)
}

func (s *IncludeExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(STParserDOT, 0)
}

func (s *IncludeExprContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *IncludeExprContext) AT() antlr.TerminalNode {
	return s.GetToken(STParserAT, 0)
}

func (s *IncludeExprContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *IncludeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterIncludeExpr(s)
	}
}

func (s *IncludeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitIncludeExpr(s)
	}
}

func (s *IncludeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitIncludeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) IncludeExpr() (localctx IIncludeExprContext) {
	localctx = NewIncludeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, STParserRULE_includeExpr)
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

	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			p.Match(STParserID)
		}
		{
			p.SetState(247)
			p.Match(STParserLPAREN)
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(STParserLBRACE-11))|(1<<(STParserSTRING-11))|(1<<(STParserSUPER-11))|(1<<(STParserTRUE-11))|(1<<(STParserFALSE-11))|(1<<(STParserAT-11))|(1<<(STParserLPAREN-11))|(1<<(STParserLBRACK-11))|(1<<(STParserID-11)))) != 0 {
			{
				p.SetState(248)
				p.MapExpr()
			}

		}
		{
			p.SetState(251)
			p.Match(STParserRPAREN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)
			p.Match(STParserSUPER)
		}
		{
			p.SetState(253)
			p.Match(STParserDOT)
		}
		{
			p.SetState(254)
			p.Match(STParserID)
		}
		{
			p.SetState(255)
			p.Match(STParserLPAREN)
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(STParserLBRACE-11))|(1<<(STParserSTRING-11))|(1<<(STParserSUPER-11))|(1<<(STParserTRUE-11))|(1<<(STParserFALSE-11))|(1<<(STParserAT-11))|(1<<(STParserELLIPSIS-11))|(1<<(STParserLPAREN-11))|(1<<(STParserLBRACK-11))|(1<<(STParserID-11)))) != 0 {
			{
				p.SetState(256)
				p.Args()
			}

		}
		{
			p.SetState(259)
			p.Match(STParserRPAREN)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.Match(STParserID)
		}
		{
			p.SetState(261)
			p.Match(STParserLPAREN)
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(STParserLBRACE-11))|(1<<(STParserSTRING-11))|(1<<(STParserSUPER-11))|(1<<(STParserTRUE-11))|(1<<(STParserFALSE-11))|(1<<(STParserAT-11))|(1<<(STParserELLIPSIS-11))|(1<<(STParserLPAREN-11))|(1<<(STParserLBRACK-11))|(1<<(STParserID-11)))) != 0 {
			{
				p.SetState(262)
				p.Args()
			}

		}
		{
			p.SetState(265)
			p.Match(STParserRPAREN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(266)
			p.Match(STParserAT)
		}
		{
			p.SetState(267)
			p.Match(STParserSUPER)
		}
		{
			p.SetState(268)
			p.Match(STParserDOT)
		}
		{
			p.SetState(269)
			p.Match(STParserID)
		}
		{
			p.SetState(270)
			p.Match(STParserLPAREN)
		}
		{
			p.SetState(271)
			p.Match(STParserRPAREN)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(272)
			p.Match(STParserAT)
		}
		{
			p.SetState(273)
			p.Match(STParserID)
		}
		{
			p.SetState(274)
			p.Match(STParserLPAREN)
		}
		{
			p.SetState(275)
			p.Match(STParserRPAREN)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(276)
			p.Primary()
		}

	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
}

func (s *PrimaryContext) STRING() antlr.TerminalNode {
	return s.GetToken(STParserSTRING, 0)
}

func (s *PrimaryContext) TRUE() antlr.TerminalNode {
	return s.GetToken(STParserTRUE, 0)
}

func (s *PrimaryContext) FALSE() antlr.TerminalNode {
	return s.GetToken(STParserFALSE, 0)
}

func (s *PrimaryContext) Subtemplate() ISubtemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubtemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubtemplateContext)
}

func (s *PrimaryContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *PrimaryContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(STParserLPAREN)
}

func (s *PrimaryContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(STParserLPAREN, i)
}

func (s *PrimaryContext) Conditional() IConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *PrimaryContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(STParserRPAREN)
}

func (s *PrimaryContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(STParserRPAREN, i)
}

func (s *PrimaryContext) MapExpr() IMapExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapExprContext)
}

func (s *PrimaryContext) ArgExprList() IArgExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgExprListContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, STParserRULE_primary)
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

	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(279)
			p.Match(STParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(280)
			p.Match(STParserSTRING)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(281)
			p.Match(STParserTRUE)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(282)
			p.Match(STParserFALSE)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(283)
			p.Subtemplate()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(284)
			p.List()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(285)
			p.Match(STParserLPAREN)
		}
		{
			p.SetState(286)
			p.Conditional()
		}
		{
			p.SetState(287)
			p.Match(STParserRPAREN)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(289)
			p.Match(STParserLPAREN)
		}
		{
			p.SetState(290)
			p.MapExpr()
		}
		{
			p.SetState(291)
			p.Match(STParserRPAREN)
		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == STParserLPAREN {
			{
				p.SetState(292)
				p.Match(STParserLPAREN)
			}
			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(STParserLBRACE-11))|(1<<(STParserSTRING-11))|(1<<(STParserSUPER-11))|(1<<(STParserTRUE-11))|(1<<(STParserFALSE-11))|(1<<(STParserAT-11))|(1<<(STParserLPAREN-11))|(1<<(STParserLBRACK-11))|(1<<(STParserID-11)))) != 0 {
				{
					p.SetState(293)
					p.ArgExprList()
				}

			}
			{
				p.SetState(296)
				p.Match(STParserRPAREN)
			}

		}

	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(STParserLBRACK, 0)
}

func (s *ListContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(STParserRBRACK, 0)
}

func (s *ListContext) ArgExprList() IArgExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgExprListContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, STParserRULE_list)
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
		p.SetState(301)
		p.Match(STParserLBRACK)
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-11)&-(0x1f+1)) == 0 && ((1<<uint((_la-11)))&((1<<(STParserLBRACE-11))|(1<<(STParserSTRING-11))|(1<<(STParserSUPER-11))|(1<<(STParserTRUE-11))|(1<<(STParserFALSE-11))|(1<<(STParserAT-11))|(1<<(STParserLPAREN-11))|(1<<(STParserLBRACK-11))|(1<<(STParserID-11)))) != 0 {
		{
			p.SetState(302)
			p.ArgExprList()
		}

	}
	{
		p.SetState(305)
		p.Match(STParserRBRACK)
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) ArgExprList() IArgExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgExprListContext)
}

func (s *ArgsContext) AllNamedArg() []INamedArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INamedArgContext)(nil)).Elem())
	var tst = make([]INamedArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INamedArgContext)
		}
	}

	return tst
}

func (s *ArgsContext) NamedArg(i int) INamedArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INamedArgContext)
}

func (s *ArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *ArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *ArgsContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(STParserELLIPSIS, 0)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, STParserRULE_args)
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

	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)
			p.ArgExprList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)
			p.NamedArg()
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(309)
					p.Match(STParserCOMMA)
				}
				{
					p.SetState(310)
					p.NamedArg()
				}

			}
			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == STParserCOMMA {
			{
				p.SetState(316)
				p.Match(STParserCOMMA)
			}
			{
				p.SetState(317)
				p.Match(STParserELLIPSIS)
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(320)
			p.Match(STParserELLIPSIS)
		}

	}

	return localctx
}

// IArgExprListContext is an interface to support dynamic dispatch.
type IArgExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgExprListContext differentiates from other interfaces.
	IsArgExprListContext()
}

type ArgExprListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgExprListContext() *ArgExprListContext {
	var p = new(ArgExprListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_argExprList
	return p
}

func (*ArgExprListContext) IsArgExprListContext() {}

func NewArgExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgExprListContext {
	var p = new(ArgExprListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_argExprList

	return p
}

func (s *ArgExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgExprListContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArgExprListContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(STParserCOMMA)
}

func (s *ArgExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(STParserCOMMA, i)
}

func (s *ArgExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterArgExprList(s)
	}
}

func (s *ArgExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitArgExprList(s)
	}
}

func (s *ArgExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitArgExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) ArgExprList() (localctx IArgExprListContext) {
	localctx = NewArgExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, STParserRULE_argExprList)
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
		p.SetState(323)
		p.Expr()
	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == STParserCOMMA {
		{
			p.SetState(324)
			p.Match(STParserCOMMA)
		}
		{
			p.SetState(325)
			p.Expr()
		}

		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INamedArgContext is an interface to support dynamic dispatch.
type INamedArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedArgContext differentiates from other interfaces.
	IsNamedArgContext()
}

type NamedArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedArgContext() *NamedArgContext {
	var p = new(NamedArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = STParserRULE_namedArg
	return p
}

func (*NamedArgContext) IsNamedArgContext() {}

func NewNamedArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedArgContext {
	var p = new(NamedArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = STParserRULE_namedArg

	return p
}

func (s *NamedArgContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedArgContext) ID() antlr.TerminalNode {
	return s.GetToken(STParserID, 0)
}

func (s *NamedArgContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(STParserEQUALS, 0)
}

func (s *NamedArgContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NamedArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.EnterNamedArg(s)
	}
}

func (s *NamedArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(STParserListener); ok {
		listenerT.ExitNamedArg(s)
	}
}

func (s *NamedArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case STParserVisitor:
		return t.VisitNamedArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *STParser) NamedArg() (localctx INamedArgContext) {
	localctx = NewNamedArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, STParserRULE_namedArg)

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
		p.SetState(331)
		p.Match(STParserID)
	}
	{
		p.SetState(332)
		p.Match(STParserEQUALS)
	}
	{
		p.SetState(333)
		p.Expr()
	}

	return localctx
}
