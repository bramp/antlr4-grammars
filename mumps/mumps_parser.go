// Code generated from mumps.g4 by ANTLR 4.9.3. DO NOT EDIT.

package mumps // mumps
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 63, 500,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 6, 2, 74, 10, 2, 13, 2, 14, 2,
	75, 3, 2, 3, 2, 3, 3, 7, 3, 81, 10, 3, 12, 3, 14, 3, 84, 11, 3, 3, 3, 5,
	3, 87, 10, 3, 3, 4, 3, 4, 3, 4, 5, 4, 92, 10, 4, 3, 5, 3, 5, 6, 5, 96,
	10, 5, 13, 5, 14, 5, 97, 3, 5, 6, 5, 101, 10, 5, 13, 5, 14, 5, 102, 3,
	5, 6, 5, 106, 10, 5, 13, 5, 14, 5, 107, 5, 5, 110, 10, 5, 3, 5, 6, 5, 113,
	10, 5, 13, 5, 14, 5, 114, 3, 5, 3, 5, 5, 5, 119, 10, 5, 3, 5, 7, 5, 122,
	10, 5, 12, 5, 14, 5, 125, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 6, 6, 131, 10,
	6, 13, 6, 14, 6, 132, 3, 7, 5, 7, 136, 10, 7, 3, 7, 3, 7, 3, 7, 5, 7, 141,
	10, 7, 3, 7, 5, 7, 144, 10, 7, 3, 7, 7, 7, 147, 10, 7, 12, 7, 14, 7, 150,
	11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 157, 10, 8, 12, 8, 14, 8, 160,
	11, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 5, 10, 167, 10, 10, 3, 10, 5, 10,
	170, 10, 10, 3, 10, 7, 10, 173, 10, 10, 12, 10, 14, 10, 176, 11, 10, 3,
	10, 6, 10, 179, 10, 10, 13, 10, 14, 10, 180, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 5, 11, 198, 10, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 7, 13, 205,
	10, 13, 12, 13, 14, 13, 208, 11, 13, 3, 13, 3, 13, 7, 13, 212, 10, 13,
	12, 13, 14, 13, 215, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 5, 14, 224, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 231,
	10, 15, 3, 16, 3, 16, 3, 17, 7, 17, 236, 10, 17, 12, 17, 14, 17, 239, 11,
	17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 249,
	10, 18, 3, 19, 3, 19, 5, 19, 253, 10, 19, 3, 20, 3, 20, 5, 20, 257, 10,
	20, 3, 20, 6, 20, 260, 10, 20, 13, 20, 14, 20, 261, 3, 20, 3, 20, 3, 20,
	5, 20, 267, 10, 20, 3, 20, 5, 20, 270, 10, 20, 3, 21, 3, 21, 6, 21, 274,
	10, 21, 13, 21, 14, 21, 275, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 5, 21, 285, 10, 21, 3, 21, 3, 21, 6, 21, 289, 10, 21, 13, 21, 14,
	21, 290, 3, 21, 3, 21, 5, 21, 295, 10, 21, 7, 21, 297, 10, 21, 12, 21,
	14, 21, 300, 11, 21, 3, 21, 3, 21, 7, 21, 304, 10, 21, 12, 21, 14, 21,
	307, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 5, 22, 313, 10, 22, 3, 23, 3,
	23, 5, 23, 317, 10, 23, 3, 23, 6, 23, 320, 10, 23, 13, 23, 14, 23, 321,
	3, 23, 3, 23, 3, 24, 3, 24, 6, 24, 328, 10, 24, 13, 24, 14, 24, 329, 3,
	24, 3, 24, 7, 24, 334, 10, 24, 12, 24, 14, 24, 337, 11, 24, 3, 24, 3, 24,
	3, 25, 3, 25, 5, 25, 343, 10, 25, 3, 25, 6, 25, 346, 10, 25, 13, 25, 14,
	25, 347, 3, 25, 3, 25, 3, 26, 3, 26, 5, 26, 354, 10, 26, 3, 26, 6, 26,
	357, 10, 26, 13, 26, 14, 26, 358, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 7, 26, 369, 10, 26, 12, 26, 14, 26, 372, 11, 26, 3, 27,
	3, 27, 5, 27, 376, 10, 27, 3, 27, 6, 27, 379, 10, 27, 13, 27, 14, 27, 380,
	3, 27, 3, 27, 3, 28, 3, 28, 5, 28, 387, 10, 28, 3, 28, 6, 28, 390, 10,
	28, 13, 28, 14, 28, 391, 3, 28, 5, 28, 395, 10, 28, 3, 29, 3, 29, 5, 29,
	399, 10, 29, 3, 29, 6, 29, 402, 10, 29, 13, 29, 14, 29, 403, 3, 29, 3,
	29, 3, 30, 3, 30, 5, 30, 410, 10, 30, 3, 30, 6, 30, 413, 10, 30, 13, 30,
	14, 30, 414, 3, 30, 3, 30, 3, 30, 7, 30, 420, 10, 30, 12, 30, 14, 30, 423,
	11, 30, 3, 31, 3, 31, 5, 31, 427, 10, 31, 3, 31, 6, 31, 430, 10, 31, 13,
	31, 14, 31, 431, 3, 31, 3, 31, 3, 32, 3, 32, 5, 32, 438, 10, 32, 3, 32,
	6, 32, 441, 10, 32, 13, 32, 14, 32, 442, 3, 32, 3, 32, 3, 33, 3, 33, 5,
	33, 449, 10, 33, 3, 33, 6, 33, 452, 10, 33, 13, 33, 14, 33, 453, 3, 33,
	3, 33, 3, 34, 5, 34, 459, 10, 34, 3, 34, 3, 34, 5, 34, 463, 10, 34, 5,
	34, 465, 10, 34, 3, 34, 7, 34, 468, 10, 34, 12, 34, 14, 34, 471, 11, 34,
	3, 34, 3, 34, 7, 34, 475, 10, 34, 12, 34, 14, 34, 478, 11, 34, 3, 34, 3,
	34, 3, 35, 3, 35, 7, 35, 484, 10, 35, 12, 35, 14, 35, 487, 11, 35, 3, 35,
	3, 35, 7, 35, 491, 10, 35, 12, 35, 14, 35, 494, 11, 35, 3, 36, 3, 36, 5,
	36, 498, 10, 36, 3, 36, 2, 2, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 68, 70, 2, 7, 9, 2, 4, 4, 6, 6, 8, 8, 12, 12, 14, 14, 17,
	17, 21, 25, 3, 2, 45, 51, 4, 2, 41, 44, 52, 52, 4, 2, 33, 33, 35, 35, 4,
	2, 36, 36, 57, 57, 2, 556, 2, 73, 3, 2, 2, 2, 4, 82, 3, 2, 2, 2, 6, 91,
	3, 2, 2, 2, 8, 109, 3, 2, 2, 2, 10, 128, 3, 2, 2, 2, 12, 135, 3, 2, 2,
	2, 14, 153, 3, 2, 2, 2, 16, 161, 3, 2, 2, 2, 18, 163, 3, 2, 2, 2, 20, 197,
	3, 2, 2, 2, 22, 199, 3, 2, 2, 2, 24, 202, 3, 2, 2, 2, 26, 223, 3, 2, 2,
	2, 28, 230, 3, 2, 2, 2, 30, 232, 3, 2, 2, 2, 32, 237, 3, 2, 2, 2, 34, 242,
	3, 2, 2, 2, 36, 250, 3, 2, 2, 2, 38, 254, 3, 2, 2, 2, 40, 271, 3, 2, 2,
	2, 42, 310, 3, 2, 2, 2, 44, 314, 3, 2, 2, 2, 46, 325, 3, 2, 2, 2, 48, 340,
	3, 2, 2, 2, 50, 351, 3, 2, 2, 2, 52, 373, 3, 2, 2, 2, 54, 384, 3, 2, 2,
	2, 56, 396, 3, 2, 2, 2, 58, 407, 3, 2, 2, 2, 60, 424, 3, 2, 2, 2, 62, 435,
	3, 2, 2, 2, 64, 446, 3, 2, 2, 2, 66, 464, 3, 2, 2, 2, 68, 481, 3, 2, 2,
	2, 70, 497, 3, 2, 2, 2, 72, 74, 5, 6, 4, 2, 73, 72, 3, 2, 2, 2, 74, 75,
	3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2,
	77, 78, 5, 4, 3, 2, 78, 3, 3, 2, 2, 2, 79, 81, 7, 61, 2, 2, 80, 79, 3,
	2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83,
	86, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 87, 7, 62, 2, 2, 86, 85, 3, 2,
	2, 2, 86, 87, 3, 2, 2, 2, 87, 5, 3, 2, 2, 2, 88, 92, 5, 8, 5, 2, 89, 92,
	5, 12, 7, 2, 90, 92, 7, 62, 2, 2, 91, 88, 3, 2, 2, 2, 91, 89, 3, 2, 2,
	2, 91, 90, 3, 2, 2, 2, 92, 7, 3, 2, 2, 2, 93, 110, 5, 10, 6, 2, 94, 96,
	7, 61, 2, 2, 95, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2,
	97, 98, 3, 2, 2, 2, 98, 100, 3, 2, 2, 2, 99, 101, 7, 54, 2, 2, 100, 99,
	3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2,
	2, 2, 103, 110, 3, 2, 2, 2, 104, 106, 7, 61, 2, 2, 105, 104, 3, 2, 2, 2,
	106, 107, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	110, 3, 2, 2, 2, 109, 93, 3, 2, 2, 2, 109, 95, 3, 2, 2, 2, 109, 105, 3,
	2, 2, 2, 110, 118, 3, 2, 2, 2, 111, 113, 5, 20, 11, 2, 112, 111, 3, 2,
	2, 2, 113, 114, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2,
	115, 119, 3, 2, 2, 2, 116, 119, 5, 46, 24, 2, 117, 119, 5, 18, 10, 2, 118,
	112, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 117, 3, 2, 2, 2, 118, 119,
	3, 2, 2, 2, 119, 123, 3, 2, 2, 2, 120, 122, 7, 61, 2, 2, 121, 120, 3, 2,
	2, 2, 122, 125, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2,
	124, 126, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 126, 127, 7, 62, 2, 2, 127,
	9, 3, 2, 2, 2, 128, 130, 5, 30, 16, 2, 129, 131, 7, 61, 2, 2, 130, 129,
	3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2,
	2, 2, 133, 11, 3, 2, 2, 2, 134, 136, 7, 32, 2, 2, 135, 134, 3, 2, 2, 2,
	135, 136, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 143, 5, 30, 16, 2, 138,
	140, 7, 37, 2, 2, 139, 141, 5, 14, 8, 2, 140, 139, 3, 2, 2, 2, 140, 141,
	3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 144, 7, 38, 2, 2, 143, 138, 3, 2,
	2, 2, 143, 144, 3, 2, 2, 2, 144, 148, 3, 2, 2, 2, 145, 147, 7, 61, 2, 2,
	146, 145, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148,
	149, 3, 2, 2, 2, 149, 151, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 152,
	7, 62, 2, 2, 152, 13, 3, 2, 2, 2, 153, 158, 5, 16, 9, 2, 154, 155, 7, 30,
	2, 2, 155, 157, 5, 16, 9, 2, 156, 154, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2,
	158, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 15, 3, 2, 2, 2, 160, 158,
	3, 2, 2, 2, 161, 162, 5, 32, 17, 2, 162, 17, 3, 2, 2, 2, 163, 169, 5, 30,
	16, 2, 164, 166, 7, 37, 2, 2, 165, 167, 5, 14, 8, 2, 166, 165, 3, 2, 2,
	2, 166, 167, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 170, 7, 38, 2, 2, 169,
	164, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 178, 3, 2, 2, 2, 171, 173,
	7, 61, 2, 2, 172, 171, 3, 2, 2, 2, 173, 176, 3, 2, 2, 2, 174, 172, 3, 2,
	2, 2, 174, 175, 3, 2, 2, 2, 175, 177, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2,
	177, 179, 5, 20, 11, 2, 178, 174, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180,
	178, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 19, 3, 2, 2, 2, 182, 198, 5,
	58, 30, 2, 183, 198, 5, 40, 21, 2, 184, 198, 5, 62, 32, 2, 185, 198, 5,
	56, 29, 2, 186, 198, 5, 54, 28, 2, 187, 198, 5, 42, 22, 2, 188, 198, 5,
	44, 23, 2, 189, 198, 5, 52, 27, 2, 190, 198, 5, 36, 19, 2, 191, 198, 5,
	38, 20, 2, 192, 198, 5, 48, 25, 2, 193, 198, 5, 60, 31, 2, 194, 198, 5,
	50, 26, 2, 195, 198, 5, 64, 33, 2, 196, 198, 9, 2, 2, 2, 197, 182, 3, 2,
	2, 2, 197, 183, 3, 2, 2, 2, 197, 184, 3, 2, 2, 2, 197, 185, 3, 2, 2, 2,
	197, 186, 3, 2, 2, 2, 197, 187, 3, 2, 2, 2, 197, 188, 3, 2, 2, 2, 197,
	189, 3, 2, 2, 2, 197, 190, 3, 2, 2, 2, 197, 191, 3, 2, 2, 2, 197, 192,
	3, 2, 2, 2, 197, 193, 3, 2, 2, 2, 197, 194, 3, 2, 2, 2, 197, 195, 3, 2,
	2, 2, 197, 196, 3, 2, 2, 2, 198, 21, 3, 2, 2, 2, 199, 200, 7, 29, 2, 2,
	200, 201, 5, 24, 13, 2, 201, 23, 3, 2, 2, 2, 202, 213, 5, 26, 14, 2, 203,
	205, 7, 61, 2, 2, 204, 203, 3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204,
	3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 209, 3, 2, 2, 2, 208, 206, 3, 2,
	2, 2, 209, 210, 9, 3, 2, 2, 210, 212, 5, 24, 13, 2, 211, 206, 3, 2, 2,
	2, 212, 215, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214,
	25, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 216, 224, 5, 32, 17, 2, 217, 224,
	5, 34, 18, 2, 218, 224, 7, 58, 2, 2, 219, 220, 7, 37, 2, 2, 220, 221, 5,
	24, 13, 2, 221, 222, 7, 38, 2, 2, 222, 224, 3, 2, 2, 2, 223, 216, 3, 2,
	2, 2, 223, 217, 3, 2, 2, 2, 223, 218, 3, 2, 2, 2, 223, 219, 3, 2, 2, 2,
	224, 27, 3, 2, 2, 2, 225, 231, 5, 26, 14, 2, 226, 227, 5, 26, 14, 2, 227,
	228, 9, 4, 2, 2, 228, 229, 5, 26, 14, 2, 229, 231, 3, 2, 2, 2, 230, 225,
	3, 2, 2, 2, 230, 226, 3, 2, 2, 2, 231, 29, 3, 2, 2, 2, 232, 233, 7, 56,
	2, 2, 233, 31, 3, 2, 2, 2, 234, 236, 9, 5, 2, 2, 235, 234, 3, 2, 2, 2,
	236, 239, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238,
	240, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 240, 241, 5, 30, 16, 2, 241, 33,
	3, 2, 2, 2, 242, 243, 7, 31, 2, 2, 243, 248, 5, 30, 16, 2, 244, 245, 7,
	37, 2, 2, 245, 246, 5, 68, 35, 2, 246, 247, 7, 38, 2, 2, 247, 249, 3, 2,
	2, 2, 248, 244, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 35, 3, 2, 2, 2,
	250, 252, 7, 3, 2, 2, 251, 253, 5, 22, 12, 2, 252, 251, 3, 2, 2, 2, 252,
	253, 3, 2, 2, 2, 253, 37, 3, 2, 2, 2, 254, 256, 7, 5, 2, 2, 255, 257, 5,
	22, 12, 2, 256, 255, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 259, 3, 2,
	2, 2, 258, 260, 7, 61, 2, 2, 259, 258, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2,
	261, 259, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263,
	269, 5, 30, 16, 2, 264, 266, 7, 37, 2, 2, 265, 267, 5, 14, 8, 2, 266, 265,
	3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 270, 7, 38,
	2, 2, 269, 264, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 39, 3, 2, 2, 2,
	271, 273, 7, 7, 2, 2, 272, 274, 7, 61, 2, 2, 273, 272, 3, 2, 2, 2, 274,
	275, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 277,
	3, 2, 2, 2, 277, 278, 5, 26, 14, 2, 278, 279, 7, 52, 2, 2, 279, 280, 5,
	26, 14, 2, 280, 284, 7, 29, 2, 2, 281, 282, 5, 26, 14, 2, 282, 283, 7,
	29, 2, 2, 283, 285, 3, 2, 2, 2, 284, 281, 3, 2, 2, 2, 284, 285, 3, 2, 2,
	2, 285, 286, 3, 2, 2, 2, 286, 288, 5, 26, 14, 2, 287, 289, 7, 61, 2, 2,
	288, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290,
	291, 3, 2, 2, 2, 291, 298, 3, 2, 2, 2, 292, 294, 5, 20, 11, 2, 293, 295,
	7, 61, 2, 2, 294, 293, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 297, 3, 2,
	2, 2, 296, 292, 3, 2, 2, 2, 297, 300, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2,
	298, 299, 3, 2, 2, 2, 299, 301, 3, 2, 2, 2, 300, 298, 3, 2, 2, 2, 301,
	305, 7, 29, 2, 2, 302, 304, 7, 61, 2, 2, 303, 302, 3, 2, 2, 2, 304, 307,
	3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 308, 3, 2,
	2, 2, 307, 305, 3, 2, 2, 2, 308, 309, 5, 28, 15, 2, 309, 41, 3, 2, 2, 2,
	310, 312, 7, 9, 2, 2, 311, 313, 5, 22, 12, 2, 312, 311, 3, 2, 2, 2, 312,
	313, 3, 2, 2, 2, 313, 43, 3, 2, 2, 2, 314, 316, 7, 10, 2, 2, 315, 317,
	5, 22, 12, 2, 316, 315, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 319, 3,
	2, 2, 2, 318, 320, 7, 61, 2, 2, 319, 318, 3, 2, 2, 2, 320, 321, 3, 2, 2,
	2, 321, 319, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323,
	324, 5, 26, 14, 2, 324, 45, 3, 2, 2, 2, 325, 327, 7, 11, 2, 2, 326, 328,
	7, 61, 2, 2, 327, 326, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 327, 3, 2,
	2, 2, 329, 330, 3, 2, 2, 2, 330, 331, 3, 2, 2, 2, 331, 335, 5, 28, 15,
	2, 332, 334, 7, 61, 2, 2, 333, 332, 3, 2, 2, 2, 334, 337, 3, 2, 2, 2, 335,
	333, 3, 2, 2, 2, 335, 336, 3, 2, 2, 2, 336, 338, 3, 2, 2, 2, 337, 335,
	3, 2, 2, 2, 338, 339, 5, 20, 11, 2, 339, 47, 3, 2, 2, 2, 340, 342, 7, 13,
	2, 2, 341, 343, 5, 22, 12, 2, 342, 341, 3, 2, 2, 2, 342, 343, 3, 2, 2,
	2, 343, 345, 3, 2, 2, 2, 344, 346, 7, 61, 2, 2, 345, 344, 3, 2, 2, 2, 346,
	347, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 349,
	3, 2, 2, 2, 349, 350, 5, 68, 35, 2, 350, 49, 3, 2, 2, 2, 351, 353, 7, 15,
	2, 2, 352, 354, 5, 22, 12, 2, 353, 352, 3, 2, 2, 2, 353, 354, 3, 2, 2,
	2, 354, 356, 3, 2, 2, 2, 355, 357, 7, 61, 2, 2, 356, 355, 3, 2, 2, 2, 357,
	358, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 360,
	3, 2, 2, 2, 360, 361, 5, 26, 14, 2, 361, 362, 7, 52, 2, 2, 362, 370, 5,
	26, 14, 2, 363, 364, 7, 30, 2, 2, 364, 365, 5, 26, 14, 2, 365, 366, 7,
	52, 2, 2, 366, 367, 5, 26, 14, 2, 367, 369, 3, 2, 2, 2, 368, 363, 3, 2,
	2, 2, 369, 372, 3, 2, 2, 2, 370, 368, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2,
	371, 51, 3, 2, 2, 2, 372, 370, 3, 2, 2, 2, 373, 375, 7, 16, 2, 2, 374,
	376, 5, 22, 12, 2, 375, 374, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 378,
	3, 2, 2, 2, 377, 379, 7, 61, 2, 2, 378, 377, 3, 2, 2, 2, 379, 380, 3, 2,
	2, 2, 380, 378, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2,
	382, 383, 5, 68, 35, 2, 383, 53, 3, 2, 2, 2, 384, 386, 7, 18, 2, 2, 385,
	387, 5, 22, 12, 2, 386, 385, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 394,
	3, 2, 2, 2, 388, 390, 7, 61, 2, 2, 389, 388, 3, 2, 2, 2, 390, 391, 3, 2,
	2, 2, 391, 389, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2,
	393, 395, 5, 26, 14, 2, 394, 389, 3, 2, 2, 2, 394, 395, 3, 2, 2, 2, 395,
	55, 3, 2, 2, 2, 396, 398, 7, 19, 2, 2, 397, 399, 5, 22, 12, 2, 398, 397,
	3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 401, 3, 2, 2, 2, 400, 402, 7, 61,
	2, 2, 401, 400, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 401, 3, 2, 2, 2,
	403, 404, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 406, 5, 68, 35, 2, 406,
	57, 3, 2, 2, 2, 407, 409, 7, 20, 2, 2, 408, 410, 5, 22, 12, 2, 409, 408,
	3, 2, 2, 2, 409, 410, 3, 2, 2, 2, 410, 412, 3, 2, 2, 2, 411, 413, 7, 61,
	2, 2, 412, 411, 3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 412, 3, 2, 2, 2,
	414, 415, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 421, 5, 66, 34, 2, 417,
	418, 7, 30, 2, 2, 418, 420, 5, 66, 34, 2, 419, 417, 3, 2, 2, 2, 420, 423,
	3, 2, 2, 2, 421, 419, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 59, 3, 2,
	2, 2, 423, 421, 3, 2, 2, 2, 424, 426, 7, 26, 2, 2, 425, 427, 5, 22, 12,
	2, 426, 425, 3, 2, 2, 2, 426, 427, 3, 2, 2, 2, 427, 429, 3, 2, 2, 2, 428,
	430, 7, 61, 2, 2, 429, 428, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431, 429,
	3, 2, 2, 2, 431, 432, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 434, 7, 56,
	2, 2, 434, 61, 3, 2, 2, 2, 435, 437, 7, 27, 2, 2, 436, 438, 5, 22, 12,
	2, 437, 436, 3, 2, 2, 2, 437, 438, 3, 2, 2, 2, 438, 440, 3, 2, 2, 2, 439,
	441, 7, 61, 2, 2, 440, 439, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 440,
	3, 2, 2, 2, 442, 443, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 445, 5, 68,
	35, 2, 445, 63, 3, 2, 2, 2, 446, 448, 7, 28, 2, 2, 447, 449, 5, 22, 12,
	2, 448, 447, 3, 2, 2, 2, 448, 449, 3, 2, 2, 2, 449, 451, 3, 2, 2, 2, 450,
	452, 7, 61, 2, 2, 451, 450, 3, 2, 2, 2, 452, 453, 3, 2, 2, 2, 453, 451,
	3, 2, 2, 2, 453, 454, 3, 2, 2, 2, 454, 455, 3, 2, 2, 2, 455, 456, 7, 57,
	2, 2, 456, 65, 3, 2, 2, 2, 457, 459, 7, 37, 2, 2, 458, 457, 3, 2, 2, 2,
	458, 459, 3, 2, 2, 2, 459, 460, 3, 2, 2, 2, 460, 462, 5, 68, 35, 2, 461,
	463, 7, 38, 2, 2, 462, 461, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2, 463, 465,
	3, 2, 2, 2, 464, 458, 3, 2, 2, 2, 464, 465, 3, 2, 2, 2, 465, 469, 3, 2,
	2, 2, 466, 468, 7, 61, 2, 2, 467, 466, 3, 2, 2, 2, 468, 471, 3, 2, 2, 2,
	469, 467, 3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470, 472, 3, 2, 2, 2, 471,
	469, 3, 2, 2, 2, 472, 476, 7, 52, 2, 2, 473, 475, 7, 61, 2, 2, 474, 473,
	3, 2, 2, 2, 475, 478, 3, 2, 2, 2, 476, 474, 3, 2, 2, 2, 476, 477, 3, 2,
	2, 2, 477, 479, 3, 2, 2, 2, 478, 476, 3, 2, 2, 2, 479, 480, 5, 70, 36,
	2, 480, 67, 3, 2, 2, 2, 481, 492, 5, 70, 36, 2, 482, 484, 7, 61, 2, 2,
	483, 482, 3, 2, 2, 2, 484, 487, 3, 2, 2, 2, 485, 483, 3, 2, 2, 2, 485,
	486, 3, 2, 2, 2, 486, 488, 3, 2, 2, 2, 487, 485, 3, 2, 2, 2, 488, 489,
	7, 30, 2, 2, 489, 491, 5, 70, 36, 2, 490, 485, 3, 2, 2, 2, 491, 494, 3,
	2, 2, 2, 492, 490, 3, 2, 2, 2, 492, 493, 3, 2, 2, 2, 493, 69, 3, 2, 2,
	2, 494, 492, 3, 2, 2, 2, 495, 498, 5, 24, 13, 2, 496, 498, 9, 6, 2, 2,
	497, 495, 3, 2, 2, 2, 497, 496, 3, 2, 2, 2, 498, 71, 3, 2, 2, 2, 75, 75,
	82, 86, 91, 97, 102, 107, 109, 114, 118, 123, 132, 135, 140, 143, 148,
	158, 166, 169, 174, 180, 197, 206, 213, 223, 230, 237, 248, 252, 256, 261,
	266, 269, 275, 284, 290, 294, 298, 305, 312, 316, 321, 329, 335, 342, 347,
	353, 358, 370, 375, 380, 386, 391, 394, 398, 403, 409, 414, 421, 426, 431,
	437, 442, 448, 453, 458, 462, 464, 469, 476, 485, 492, 497,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "':'", "','", "'$'", "'%'", "'&'",
	"'@'", "'^'", "'!'", "'('", "')'", "'{'", "'}'", "''>'", "''<'", "'>'",
	"'<'", "'+'", "'-'", "'*'", "'/'", "'\\'", "'#'", "'**'", "'='", "'?'",
	"'.'", "'_'", "", "", "", "'''", "", "' '",
}
var symbolicNames = []string{
	"", "BREAK", "CLOSE", "DO", "ELSE", "FOR", "GOTO", "HALT", "HANG", "IF",
	"JOB", "KILL", "LOCK", "MERGE", "NEW", "OPEN", "QUIT", "READ", "SET", "TCOMMIT",
	"TRESTART", "TROLLBACK", "TSTART", "USE", "VIEW", "WRITE", "XECUTE", "COLON",
	"COMMA", "DOLLAR", "PERCENT", "AMPERSAND", "INDIRECT", "CARAT", "BANG",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "NGT", "NLT", "GT", "LT", "ADD",
	"SUBTRACT", "MULTIPLY", "DIVIDE", "INTDIVIDE", "MODULO", "EXPONENT", "EQUALS",
	"QUESTION", "DOT", "CONCAT", "IDENTIFIER", "STRING_LITERAL", "NUMBER",
	"NOT", "COMMENT", "SPACE", "CR", "WS",
}

var ruleNames = []string{
	"program", "eof", "line", "code", "label", "routinedecl", "paramlist",
	"param", "subproc", "command", "postcondition", "expression", "term", "condition",
	"identifier", "variable", "function_", "break_", "do_", "for_", "halt_",
	"hang_", "if_", "kill_", "merge_", "new_", "quit_", "read_", "set_", "view_",
	"write_", "xecute_", "assign", "arglist", "arg",
}

type mumpsParser struct {
	*antlr.BaseParser
}

// NewmumpsParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *mumpsParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewmumpsParser(input antlr.TokenStream) *mumpsParser {
	this := new(mumpsParser)
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
	this.GrammarFileName = "mumps.g4"

	return this
}

// mumpsParser tokens.
const (
	mumpsParserEOF            = antlr.TokenEOF
	mumpsParserBREAK          = 1
	mumpsParserCLOSE          = 2
	mumpsParserDO             = 3
	mumpsParserELSE           = 4
	mumpsParserFOR            = 5
	mumpsParserGOTO           = 6
	mumpsParserHALT           = 7
	mumpsParserHANG           = 8
	mumpsParserIF             = 9
	mumpsParserJOB            = 10
	mumpsParserKILL           = 11
	mumpsParserLOCK           = 12
	mumpsParserMERGE          = 13
	mumpsParserNEW            = 14
	mumpsParserOPEN           = 15
	mumpsParserQUIT           = 16
	mumpsParserREAD           = 17
	mumpsParserSET            = 18
	mumpsParserTCOMMIT        = 19
	mumpsParserTRESTART       = 20
	mumpsParserTROLLBACK      = 21
	mumpsParserTSTART         = 22
	mumpsParserUSE            = 23
	mumpsParserVIEW           = 24
	mumpsParserWRITE          = 25
	mumpsParserXECUTE         = 26
	mumpsParserCOLON          = 27
	mumpsParserCOMMA          = 28
	mumpsParserDOLLAR         = 29
	mumpsParserPERCENT        = 30
	mumpsParserAMPERSAND      = 31
	mumpsParserINDIRECT       = 32
	mumpsParserCARAT          = 33
	mumpsParserBANG           = 34
	mumpsParserLPAREN         = 35
	mumpsParserRPAREN         = 36
	mumpsParserLBRACE         = 37
	mumpsParserRBRACE         = 38
	mumpsParserNGT            = 39
	mumpsParserNLT            = 40
	mumpsParserGT             = 41
	mumpsParserLT             = 42
	mumpsParserADD            = 43
	mumpsParserSUBTRACT       = 44
	mumpsParserMULTIPLY       = 45
	mumpsParserDIVIDE         = 46
	mumpsParserINTDIVIDE      = 47
	mumpsParserMODULO         = 48
	mumpsParserEXPONENT       = 49
	mumpsParserEQUALS         = 50
	mumpsParserQUESTION       = 51
	mumpsParserDOT            = 52
	mumpsParserCONCAT         = 53
	mumpsParserIDENTIFIER     = 54
	mumpsParserSTRING_LITERAL = 55
	mumpsParserNUMBER         = 56
	mumpsParserNOT            = 57
	mumpsParserCOMMENT        = 58
	mumpsParserSPACE          = 59
	mumpsParserCR             = 60
	mumpsParserWS             = 61
)

// mumpsParser rules.
const (
	mumpsParserRULE_program       = 0
	mumpsParserRULE_eof           = 1
	mumpsParserRULE_line          = 2
	mumpsParserRULE_code          = 3
	mumpsParserRULE_label         = 4
	mumpsParserRULE_routinedecl   = 5
	mumpsParserRULE_paramlist     = 6
	mumpsParserRULE_param         = 7
	mumpsParserRULE_subproc       = 8
	mumpsParserRULE_command       = 9
	mumpsParserRULE_postcondition = 10
	mumpsParserRULE_expression    = 11
	mumpsParserRULE_term          = 12
	mumpsParserRULE_condition     = 13
	mumpsParserRULE_identifier    = 14
	mumpsParserRULE_variable      = 15
	mumpsParserRULE_function_     = 16
	mumpsParserRULE_break_        = 17
	mumpsParserRULE_do_           = 18
	mumpsParserRULE_for_          = 19
	mumpsParserRULE_halt_         = 20
	mumpsParserRULE_hang_         = 21
	mumpsParserRULE_if_           = 22
	mumpsParserRULE_kill_         = 23
	mumpsParserRULE_merge_        = 24
	mumpsParserRULE_new_          = 25
	mumpsParserRULE_quit_         = 26
	mumpsParserRULE_read_         = 27
	mumpsParserRULE_set_          = 28
	mumpsParserRULE_view_         = 29
	mumpsParserRULE_write_        = 30
	mumpsParserRULE_xecute_       = 31
	mumpsParserRULE_assign        = 32
	mumpsParserRULE_arglist       = 33
	mumpsParserRULE_arg           = 34
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Eof() IEofContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEofContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEofContext)
}

func (s *ProgramContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *ProgramContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *mumpsParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, mumpsParserRULE_program)

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
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(70)
				p.Line()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(75)
		p.Eof()
	}

	return localctx
}

// IEofContext is an interface to support dynamic dispatch.
type IEofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEofContext differentiates from other interfaces.
	IsEofContext()
}

type EofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEofContext() *EofContext {
	var p = new(EofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_eof
	return p
}

func (*EofContext) IsEofContext() {}

func NewEofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EofContext {
	var p = new(EofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_eof

	return p
}

func (s *EofContext) GetParser() antlr.Parser { return s.parser }

func (s *EofContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *EofContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *EofContext) CR() antlr.TerminalNode {
	return s.GetToken(mumpsParserCR, 0)
}

func (s *EofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterEof(s)
	}
}

func (s *EofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitEof(s)
	}
}

func (p *mumpsParser) Eof() (localctx IEofContext) {
	this := p
	_ = this

	localctx = NewEofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, mumpsParserRULE_eof)
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
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserSPACE {
		{
			p.SetState(77)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCR {
		{
			p.SetState(83)
			p.Match(mumpsParserCR)
		}

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
	p.RuleIndex = mumpsParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Code() ICodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeContext)
}

func (s *LineContext) Routinedecl() IRoutinedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRoutinedeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRoutinedeclContext)
}

func (s *LineContext) CR() antlr.TerminalNode {
	return s.GetToken(mumpsParserCR, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *mumpsParser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, mumpsParserRULE_line)

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

	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Code()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Routinedecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Match(mumpsParserCR)
		}

	}

	return localctx
}

// ICodeContext is an interface to support dynamic dispatch.
type ICodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCodeContext differentiates from other interfaces.
	IsCodeContext()
}

type CodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeContext() *CodeContext {
	var p = new(CodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_code
	return p
}

func (*CodeContext) IsCodeContext() {}

func NewCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeContext {
	var p = new(CodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_code

	return p
}

func (s *CodeContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeContext) CR() antlr.TerminalNode {
	return s.GetToken(mumpsParserCR, 0)
}

func (s *CodeContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *CodeContext) If_() IIf_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_Context)
}

func (s *CodeContext) Subproc() ISubprocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubprocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubprocContext)
}

func (s *CodeContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *CodeContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *CodeContext) AllCommand() []ICommandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandContext)(nil)).Elem())
	var tst = make([]ICommandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandContext)
		}
	}

	return tst
}

func (s *CodeContext) Command(i int) ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CodeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserDOT)
}

func (s *CodeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserDOT, i)
}

func (s *CodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterCode(s)
	}
}

func (s *CodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitCode(s)
	}
}

func (p *mumpsParser) Code() (localctx ICodeContext) {
	this := p
	_ = this

	localctx = NewCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, mumpsParserRULE_code)
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
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(91)
			p.Label()
		}

	case 2:
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == mumpsParserSPACE {
			{
				p.SetState(92)
				p.Match(mumpsParserSPACE)
			}

			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == mumpsParserDOT {
			{
				p.SetState(97)
				p.Match(mumpsParserDOT)
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(102)
					p.Match(mumpsParserSPACE)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}

	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mumpsParserBREAK, mumpsParserCLOSE, mumpsParserDO, mumpsParserELSE, mumpsParserFOR, mumpsParserGOTO, mumpsParserHALT, mumpsParserHANG, mumpsParserJOB, mumpsParserKILL, mumpsParserLOCK, mumpsParserMERGE, mumpsParserNEW, mumpsParserOPEN, mumpsParserQUIT, mumpsParserREAD, mumpsParserSET, mumpsParserTCOMMIT, mumpsParserTRESTART, mumpsParserTROLLBACK, mumpsParserTSTART, mumpsParserUSE, mumpsParserVIEW, mumpsParserWRITE, mumpsParserXECUTE:
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mumpsParserBREAK)|(1<<mumpsParserCLOSE)|(1<<mumpsParserDO)|(1<<mumpsParserELSE)|(1<<mumpsParserFOR)|(1<<mumpsParserGOTO)|(1<<mumpsParserHALT)|(1<<mumpsParserHANG)|(1<<mumpsParserJOB)|(1<<mumpsParserKILL)|(1<<mumpsParserLOCK)|(1<<mumpsParserMERGE)|(1<<mumpsParserNEW)|(1<<mumpsParserOPEN)|(1<<mumpsParserQUIT)|(1<<mumpsParserREAD)|(1<<mumpsParserSET)|(1<<mumpsParserTCOMMIT)|(1<<mumpsParserTRESTART)|(1<<mumpsParserTROLLBACK)|(1<<mumpsParserTSTART)|(1<<mumpsParserUSE)|(1<<mumpsParserVIEW)|(1<<mumpsParserWRITE)|(1<<mumpsParserXECUTE))) != 0) {
			{
				p.SetState(109)
				p.Command()
			}

			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case mumpsParserIF:
		{
			p.SetState(114)
			p.If_()
		}

	case mumpsParserIDENTIFIER:
		{
			p.SetState(115)
			p.Subproc()
		}

	case mumpsParserSPACE, mumpsParserCR:

	default:
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserSPACE {
		{
			p.SetState(118)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(124)
		p.Match(mumpsParserCR)
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
	p.RuleIndex = mumpsParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *LabelContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *LabelContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *mumpsParser) Label() (localctx ILabelContext) {
	this := p
	_ = this

	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, mumpsParserRULE_label)

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
		p.SetState(126)
		p.Identifier()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(127)
				p.Match(mumpsParserSPACE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IRoutinedeclContext is an interface to support dynamic dispatch.
type IRoutinedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoutinedeclContext differentiates from other interfaces.
	IsRoutinedeclContext()
}

type RoutinedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoutinedeclContext() *RoutinedeclContext {
	var p = new(RoutinedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_routinedecl
	return p
}

func (*RoutinedeclContext) IsRoutinedeclContext() {}

func NewRoutinedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoutinedeclContext {
	var p = new(RoutinedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_routinedecl

	return p
}

func (s *RoutinedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *RoutinedeclContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *RoutinedeclContext) CR() antlr.TerminalNode {
	return s.GetToken(mumpsParserCR, 0)
}

func (s *RoutinedeclContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(mumpsParserPERCENT, 0)
}

func (s *RoutinedeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserLPAREN, 0)
}

func (s *RoutinedeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserRPAREN, 0)
}

func (s *RoutinedeclContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *RoutinedeclContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *RoutinedeclContext) Paramlist() IParamlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamlistContext)
}

func (s *RoutinedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoutinedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoutinedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterRoutinedecl(s)
	}
}

func (s *RoutinedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitRoutinedecl(s)
	}
}

func (p *mumpsParser) Routinedecl() (localctx IRoutinedeclContext) {
	this := p
	_ = this

	localctx = NewRoutinedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, mumpsParserRULE_routinedecl)
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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserPERCENT {
		{
			p.SetState(132)
			p.Match(mumpsParserPERCENT)
		}

	}
	{
		p.SetState(135)
		p.Identifier()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserLPAREN {
		{
			p.SetState(136)
			p.Match(mumpsParserLPAREN)
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(mumpsParserAMPERSAND-31))|(1<<(mumpsParserCARAT-31))|(1<<(mumpsParserIDENTIFIER-31)))) != 0 {
			{
				p.SetState(137)
				p.Paramlist()
			}

		}
		{
			p.SetState(140)
			p.Match(mumpsParserRPAREN)
		}

	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserSPACE {
		{
			p.SetState(143)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
		p.Match(mumpsParserCR)
	}

	return localctx
}

// IParamlistContext is an interface to support dynamic dispatch.
type IParamlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamlistContext differentiates from other interfaces.
	IsParamlistContext()
}

type ParamlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamlistContext() *ParamlistContext {
	var p = new(ParamlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_paramlist
	return p
}

func (*ParamlistContext) IsParamlistContext() {}

func NewParamlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamlistContext {
	var p = new(ParamlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_paramlist

	return p
}

func (s *ParamlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamlistContext) AllParam() []IParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamContext)(nil)).Elem())
	var tst = make([]IParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamContext)
		}
	}

	return tst
}

func (s *ParamlistContext) Param(i int) IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamlistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserCOMMA)
}

func (s *ParamlistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserCOMMA, i)
}

func (s *ParamlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterParamlist(s)
	}
}

func (s *ParamlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitParamlist(s)
	}
}

func (p *mumpsParser) Paramlist() (localctx IParamlistContext) {
	this := p
	_ = this

	localctx = NewParamlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, mumpsParserRULE_paramlist)
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
		p.SetState(151)
		p.Param()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserCOMMA {
		{
			p.SetState(152)
			p.Match(mumpsParserCOMMA)
		}
		{
			p.SetState(153)
			p.Param()
		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *mumpsParser) Param() (localctx IParamContext) {
	this := p
	_ = this

	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, mumpsParserRULE_param)

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
		p.Variable()
	}

	return localctx
}

// ISubprocContext is an interface to support dynamic dispatch.
type ISubprocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubprocContext differentiates from other interfaces.
	IsSubprocContext()
}

type SubprocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubprocContext() *SubprocContext {
	var p = new(SubprocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_subproc
	return p
}

func (*SubprocContext) IsSubprocContext() {}

func NewSubprocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubprocContext {
	var p = new(SubprocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_subproc

	return p
}

func (s *SubprocContext) GetParser() antlr.Parser { return s.parser }

func (s *SubprocContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SubprocContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserLPAREN, 0)
}

func (s *SubprocContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserRPAREN, 0)
}

func (s *SubprocContext) AllCommand() []ICommandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandContext)(nil)).Elem())
	var tst = make([]ICommandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandContext)
		}
	}

	return tst
}

func (s *SubprocContext) Command(i int) ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *SubprocContext) Paramlist() IParamlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamlistContext)
}

func (s *SubprocContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *SubprocContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *SubprocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubprocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubprocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterSubproc(s)
	}
}

func (s *SubprocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitSubproc(s)
	}
}

func (p *mumpsParser) Subproc() (localctx ISubprocContext) {
	this := p
	_ = this

	localctx = NewSubprocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, mumpsParserRULE_subproc)
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
		p.SetState(161)
		p.Identifier()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserLPAREN {
		{
			p.SetState(162)
			p.Match(mumpsParserLPAREN)
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(mumpsParserAMPERSAND-31))|(1<<(mumpsParserCARAT-31))|(1<<(mumpsParserIDENTIFIER-31)))) != 0 {
			{
				p.SetState(163)
				p.Paramlist()
			}

		}
		{
			p.SetState(166)
			p.Match(mumpsParserRPAREN)
		}

	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(172)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == mumpsParserSPACE {
				{
					p.SetState(169)
					p.Match(mumpsParserSPACE)
				}

				p.SetState(174)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(175)
				p.Command()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Set_() ISet_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_Context)
}

func (s *CommandContext) For_() IFor_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_Context)
}

func (s *CommandContext) Write_() IWrite_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWrite_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWrite_Context)
}

func (s *CommandContext) Read_() IRead_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRead_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRead_Context)
}

func (s *CommandContext) Quit_() IQuit_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuit_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuit_Context)
}

func (s *CommandContext) Halt_() IHalt_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHalt_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHalt_Context)
}

func (s *CommandContext) Hang_() IHang_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHang_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHang_Context)
}

func (s *CommandContext) New_() INew_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INew_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INew_Context)
}

func (s *CommandContext) Break_() IBreak_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreak_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreak_Context)
}

func (s *CommandContext) Do_() IDo_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDo_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDo_Context)
}

func (s *CommandContext) Kill_() IKill_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKill_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKill_Context)
}

func (s *CommandContext) View_() IView_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IView_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IView_Context)
}

func (s *CommandContext) Merge_() IMerge_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMerge_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMerge_Context)
}

func (s *CommandContext) Xecute_() IXecute_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXecute_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXecute_Context)
}

func (s *CommandContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(mumpsParserCLOSE, 0)
}

func (s *CommandContext) ELSE() antlr.TerminalNode {
	return s.GetToken(mumpsParserELSE, 0)
}

func (s *CommandContext) GOTO() antlr.TerminalNode {
	return s.GetToken(mumpsParserGOTO, 0)
}

func (s *CommandContext) JOB() antlr.TerminalNode {
	return s.GetToken(mumpsParserJOB, 0)
}

func (s *CommandContext) LOCK() antlr.TerminalNode {
	return s.GetToken(mumpsParserLOCK, 0)
}

func (s *CommandContext) OPEN() antlr.TerminalNode {
	return s.GetToken(mumpsParserOPEN, 0)
}

func (s *CommandContext) TCOMMIT() antlr.TerminalNode {
	return s.GetToken(mumpsParserTCOMMIT, 0)
}

func (s *CommandContext) TRESTART() antlr.TerminalNode {
	return s.GetToken(mumpsParserTRESTART, 0)
}

func (s *CommandContext) TROLLBACK() antlr.TerminalNode {
	return s.GetToken(mumpsParserTROLLBACK, 0)
}

func (s *CommandContext) TSTART() antlr.TerminalNode {
	return s.GetToken(mumpsParserTSTART, 0)
}

func (s *CommandContext) USE() antlr.TerminalNode {
	return s.GetToken(mumpsParserUSE, 0)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *mumpsParser) Command() (localctx ICommandContext) {
	this := p
	_ = this

	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, mumpsParserRULE_command)
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

	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mumpsParserSET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Set_()
		}

	case mumpsParserFOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.For_()
		}

	case mumpsParserWRITE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Write_()
		}

	case mumpsParserREAD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.Read_()
		}

	case mumpsParserQUIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(184)
			p.Quit_()
		}

	case mumpsParserHALT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(185)
			p.Halt_()
		}

	case mumpsParserHANG:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(186)
			p.Hang_()
		}

	case mumpsParserNEW:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(187)
			p.New_()
		}

	case mumpsParserBREAK:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(188)
			p.Break_()
		}

	case mumpsParserDO:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(189)
			p.Do_()
		}

	case mumpsParserKILL:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(190)
			p.Kill_()
		}

	case mumpsParserVIEW:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(191)
			p.View_()
		}

	case mumpsParserMERGE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(192)
			p.Merge_()
		}

	case mumpsParserXECUTE:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(193)
			p.Xecute_()
		}

	case mumpsParserCLOSE, mumpsParserELSE, mumpsParserGOTO, mumpsParserJOB, mumpsParserLOCK, mumpsParserOPEN, mumpsParserTCOMMIT, mumpsParserTRESTART, mumpsParserTROLLBACK, mumpsParserTSTART, mumpsParserUSE:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(194)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mumpsParserCLOSE)|(1<<mumpsParserELSE)|(1<<mumpsParserGOTO)|(1<<mumpsParserJOB)|(1<<mumpsParserLOCK)|(1<<mumpsParserOPEN)|(1<<mumpsParserTCOMMIT)|(1<<mumpsParserTRESTART)|(1<<mumpsParserTROLLBACK)|(1<<mumpsParserTSTART)|(1<<mumpsParserUSE))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPostconditionContext is an interface to support dynamic dispatch.
type IPostconditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostconditionContext differentiates from other interfaces.
	IsPostconditionContext()
}

type PostconditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostconditionContext() *PostconditionContext {
	var p = new(PostconditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_postcondition
	return p
}

func (*PostconditionContext) IsPostconditionContext() {}

func NewPostconditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostconditionContext {
	var p = new(PostconditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_postcondition

	return p
}

func (s *PostconditionContext) GetParser() antlr.Parser { return s.parser }

func (s *PostconditionContext) COLON() antlr.TerminalNode {
	return s.GetToken(mumpsParserCOLON, 0)
}

func (s *PostconditionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PostconditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostconditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostconditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterPostcondition(s)
	}
}

func (s *PostconditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitPostcondition(s)
	}
}

func (p *mumpsParser) Postcondition() (localctx IPostconditionContext) {
	this := p
	_ = this

	localctx = NewPostconditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, mumpsParserRULE_postcondition)

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
		p.Match(mumpsParserCOLON)
	}
	{
		p.SetState(198)
		p.Expression()
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
	p.RuleIndex = mumpsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) AllADD() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserADD)
}

func (s *ExpressionContext) ADD(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserADD, i)
}

func (s *ExpressionContext) AllMULTIPLY() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserMULTIPLY)
}

func (s *ExpressionContext) MULTIPLY(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserMULTIPLY, i)
}

func (s *ExpressionContext) AllSUBTRACT() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSUBTRACT)
}

func (s *ExpressionContext) SUBTRACT(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSUBTRACT, i)
}

func (s *ExpressionContext) AllDIVIDE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserDIVIDE)
}

func (s *ExpressionContext) DIVIDE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserDIVIDE, i)
}

func (s *ExpressionContext) AllINTDIVIDE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserINTDIVIDE)
}

func (s *ExpressionContext) INTDIVIDE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserINTDIVIDE, i)
}

func (s *ExpressionContext) AllMODULO() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserMODULO)
}

func (s *ExpressionContext) MODULO(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserMODULO, i)
}

func (s *ExpressionContext) AllEXPONENT() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserEXPONENT)
}

func (s *ExpressionContext) EXPONENT(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserEXPONENT, i)
}

func (s *ExpressionContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *ExpressionContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *mumpsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, mumpsParserRULE_expression)
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
		p.SetState(200)
		p.Term()
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(204)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == mumpsParserSPACE {
				{
					p.SetState(201)
					p.Match(mumpsParserSPACE)
				}

				p.SetState(206)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(207)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(mumpsParserADD-43))|(1<<(mumpsParserSUBTRACT-43))|(1<<(mumpsParserMULTIPLY-43))|(1<<(mumpsParserDIVIDE-43))|(1<<(mumpsParserINTDIVIDE-43))|(1<<(mumpsParserMODULO-43))|(1<<(mumpsParserEXPONENT-43)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(208)
				p.Expression()
			}

		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *TermContext) Function_() IFunction_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_Context)
}

func (s *TermContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(mumpsParserNUMBER, 0)
}

func (s *TermContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserLPAREN, 0)
}

func (s *TermContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TermContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserRPAREN, 0)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *mumpsParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, mumpsParserRULE_term)

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

	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mumpsParserAMPERSAND, mumpsParserCARAT, mumpsParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Variable()
		}

	case mumpsParserDOLLAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Function_()
		}

	case mumpsParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Match(mumpsParserNUMBER)
		}

	case mumpsParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
			p.Match(mumpsParserLPAREN)
		}
		{
			p.SetState(218)
			p.Expression()
		}
		{
			p.SetState(219)
			p.Match(mumpsParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *ConditionContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ConditionContext) NGT() antlr.TerminalNode {
	return s.GetToken(mumpsParserNGT, 0)
}

func (s *ConditionContext) NLT() antlr.TerminalNode {
	return s.GetToken(mumpsParserNLT, 0)
}

func (s *ConditionContext) LT() antlr.TerminalNode {
	return s.GetToken(mumpsParserLT, 0)
}

func (s *ConditionContext) GT() antlr.TerminalNode {
	return s.GetToken(mumpsParserGT, 0)
}

func (s *ConditionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(mumpsParserEQUALS, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *mumpsParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, mumpsParserRULE_condition)
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

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Term()
		}
		{
			p.SetState(225)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(mumpsParserNGT-39))|(1<<(mumpsParserNLT-39))|(1<<(mumpsParserGT-39))|(1<<(mumpsParserLT-39))|(1<<(mumpsParserEQUALS-39)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(226)
			p.Term()
		}

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
	p.RuleIndex = mumpsParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(mumpsParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *mumpsParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, mumpsParserRULE_identifier)

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
		p.Match(mumpsParserIDENTIFIER)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *VariableContext) AllCARAT() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserCARAT)
}

func (s *VariableContext) CARAT(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserCARAT, i)
}

func (s *VariableContext) AllAMPERSAND() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserAMPERSAND)
}

func (s *VariableContext) AMPERSAND(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserAMPERSAND, i)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *mumpsParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, mumpsParserRULE_variable)
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
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserAMPERSAND || _la == mumpsParserCARAT {
		{
			p.SetState(232)
			_la = p.GetTokenStream().LA(1)

			if !(_la == mumpsParserAMPERSAND || _la == mumpsParserCARAT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(238)
		p.Identifier()
	}

	return localctx
}

// IFunction_Context is an interface to support dynamic dispatch.
type IFunction_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_Context differentiates from other interfaces.
	IsFunction_Context()
}

type Function_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_Context() *Function_Context {
	var p = new(Function_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_function_
	return p
}

func (*Function_Context) IsFunction_Context() {}

func NewFunction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_Context {
	var p = new(Function_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_function_

	return p
}

func (s *Function_Context) GetParser() antlr.Parser { return s.parser }

func (s *Function_Context) DOLLAR() antlr.TerminalNode {
	return s.GetToken(mumpsParserDOLLAR, 0)
}

func (s *Function_Context) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Function_Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserLPAREN, 0)
}

func (s *Function_Context) Arglist() IArglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *Function_Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserRPAREN, 0)
}

func (s *Function_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterFunction_(s)
	}
}

func (s *Function_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitFunction_(s)
	}
}

func (p *mumpsParser) Function_() (localctx IFunction_Context) {
	this := p
	_ = this

	localctx = NewFunction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, mumpsParserRULE_function_)
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
		p.SetState(240)
		p.Match(mumpsParserDOLLAR)
	}
	{
		p.SetState(241)
		p.Identifier()
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserLPAREN {
		{
			p.SetState(242)
			p.Match(mumpsParserLPAREN)
		}
		{
			p.SetState(243)
			p.Arglist()
		}
		{
			p.SetState(244)
			p.Match(mumpsParserRPAREN)
		}

	}

	return localctx
}

// IBreak_Context is an interface to support dynamic dispatch.
type IBreak_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreak_Context differentiates from other interfaces.
	IsBreak_Context()
}

type Break_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_Context() *Break_Context {
	var p = new(Break_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_break_
	return p
}

func (*Break_Context) IsBreak_Context() {}

func NewBreak_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_Context {
	var p = new(Break_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_break_

	return p
}

func (s *Break_Context) GetParser() antlr.Parser { return s.parser }

func (s *Break_Context) BREAK() antlr.TerminalNode {
	return s.GetToken(mumpsParserBREAK, 0)
}

func (s *Break_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Break_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterBreak_(s)
	}
}

func (s *Break_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitBreak_(s)
	}
}

func (p *mumpsParser) Break_() (localctx IBreak_Context) {
	this := p
	_ = this

	localctx = NewBreak_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, mumpsParserRULE_break_)

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
		p.SetState(248)
		p.Match(mumpsParserBREAK)
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(249)
			p.Postcondition()
		}

	}

	return localctx
}

// IDo_Context is an interface to support dynamic dispatch.
type IDo_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDo_Context differentiates from other interfaces.
	IsDo_Context()
}

type Do_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDo_Context() *Do_Context {
	var p = new(Do_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_do_
	return p
}

func (*Do_Context) IsDo_Context() {}

func NewDo_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Do_Context {
	var p = new(Do_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_do_

	return p
}

func (s *Do_Context) GetParser() antlr.Parser { return s.parser }

func (s *Do_Context) DO() antlr.TerminalNode {
	return s.GetToken(mumpsParserDO, 0)
}

func (s *Do_Context) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Do_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Do_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *Do_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *Do_Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserLPAREN, 0)
}

func (s *Do_Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserRPAREN, 0)
}

func (s *Do_Context) Paramlist() IParamlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamlistContext)
}

func (s *Do_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Do_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Do_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterDo_(s)
	}
}

func (s *Do_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitDo_(s)
	}
}

func (p *mumpsParser) Do_() (localctx IDo_Context) {
	this := p
	_ = this

	localctx = NewDo_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, mumpsParserRULE_do_)
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
		p.SetState(252)
		p.Match(mumpsParserDO)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCOLON {
		{
			p.SetState(253)
			p.Postcondition()
		}

	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(256)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(261)
		p.Identifier()
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserLPAREN {
		{
			p.SetState(262)
			p.Match(mumpsParserLPAREN)
		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(mumpsParserAMPERSAND-31))|(1<<(mumpsParserCARAT-31))|(1<<(mumpsParserIDENTIFIER-31)))) != 0 {
			{
				p.SetState(263)
				p.Paramlist()
			}

		}
		{
			p.SetState(266)
			p.Match(mumpsParserRPAREN)
		}

	}

	return localctx
}

// IFor_Context is an interface to support dynamic dispatch.
type IFor_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_Context differentiates from other interfaces.
	IsFor_Context()
}

type For_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_Context() *For_Context {
	var p = new(For_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_for_
	return p
}

func (*For_Context) IsFor_Context() {}

func NewFor_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_Context {
	var p = new(For_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_for_

	return p
}

func (s *For_Context) GetParser() antlr.Parser { return s.parser }

func (s *For_Context) FOR() antlr.TerminalNode {
	return s.GetToken(mumpsParserFOR, 0)
}

func (s *For_Context) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *For_Context) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *For_Context) EQUALS() antlr.TerminalNode {
	return s.GetToken(mumpsParserEQUALS, 0)
}

func (s *For_Context) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserCOLON)
}

func (s *For_Context) COLON(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserCOLON, i)
}

func (s *For_Context) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *For_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *For_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *For_Context) AllCommand() []ICommandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandContext)(nil)).Elem())
	var tst = make([]ICommandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandContext)
		}
	}

	return tst
}

func (s *For_Context) Command(i int) ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *For_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterFor_(s)
	}
}

func (s *For_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitFor_(s)
	}
}

func (p *mumpsParser) For_() (localctx IFor_Context) {
	this := p
	_ = this

	localctx = NewFor_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, mumpsParserRULE_for_)
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
		p.SetState(269)
		p.Match(mumpsParserFOR)
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(270)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(275)
		p.Term()
	}
	{
		p.SetState(276)
		p.Match(mumpsParserEQUALS)
	}
	{
		p.SetState(277)
		p.Term()
	}
	{
		p.SetState(278)
		p.Match(mumpsParserCOLON)
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(279)
			p.Term()
		}
		{
			p.SetState(280)
			p.Match(mumpsParserCOLON)
		}

	}
	{
		p.SetState(284)
		p.Term()
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(285)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mumpsParserBREAK)|(1<<mumpsParserCLOSE)|(1<<mumpsParserDO)|(1<<mumpsParserELSE)|(1<<mumpsParserFOR)|(1<<mumpsParserGOTO)|(1<<mumpsParserHALT)|(1<<mumpsParserHANG)|(1<<mumpsParserJOB)|(1<<mumpsParserKILL)|(1<<mumpsParserLOCK)|(1<<mumpsParserMERGE)|(1<<mumpsParserNEW)|(1<<mumpsParserOPEN)|(1<<mumpsParserQUIT)|(1<<mumpsParserREAD)|(1<<mumpsParserSET)|(1<<mumpsParserTCOMMIT)|(1<<mumpsParserTRESTART)|(1<<mumpsParserTROLLBACK)|(1<<mumpsParserTSTART)|(1<<mumpsParserUSE)|(1<<mumpsParserVIEW)|(1<<mumpsParserWRITE)|(1<<mumpsParserXECUTE))) != 0 {
		{
			p.SetState(290)
			p.Command()
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == mumpsParserSPACE {
			{
				p.SetState(291)
				p.Match(mumpsParserSPACE)
			}

		}

		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(299)
		p.Match(mumpsParserCOLON)
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserSPACE {
		{
			p.SetState(300)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(306)
		p.Condition()
	}

	return localctx
}

// IHalt_Context is an interface to support dynamic dispatch.
type IHalt_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHalt_Context differentiates from other interfaces.
	IsHalt_Context()
}

type Halt_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHalt_Context() *Halt_Context {
	var p = new(Halt_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_halt_
	return p
}

func (*Halt_Context) IsHalt_Context() {}

func NewHalt_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Halt_Context {
	var p = new(Halt_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_halt_

	return p
}

func (s *Halt_Context) GetParser() antlr.Parser { return s.parser }

func (s *Halt_Context) HALT() antlr.TerminalNode {
	return s.GetToken(mumpsParserHALT, 0)
}

func (s *Halt_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Halt_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Halt_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Halt_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterHalt_(s)
	}
}

func (s *Halt_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitHalt_(s)
	}
}

func (p *mumpsParser) Halt_() (localctx IHalt_Context) {
	this := p
	_ = this

	localctx = NewHalt_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, mumpsParserRULE_halt_)

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
		p.SetState(308)
		p.Match(mumpsParserHALT)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(309)
			p.Postcondition()
		}

	}

	return localctx
}

// IHang_Context is an interface to support dynamic dispatch.
type IHang_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHang_Context differentiates from other interfaces.
	IsHang_Context()
}

type Hang_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHang_Context() *Hang_Context {
	var p = new(Hang_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_hang_
	return p
}

func (*Hang_Context) IsHang_Context() {}

func NewHang_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Hang_Context {
	var p = new(Hang_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_hang_

	return p
}

func (s *Hang_Context) GetParser() antlr.Parser { return s.parser }

func (s *Hang_Context) HANG() antlr.TerminalNode {
	return s.GetToken(mumpsParserHANG, 0)
}

func (s *Hang_Context) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Hang_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Hang_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *Hang_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *Hang_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Hang_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Hang_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterHang_(s)
	}
}

func (s *Hang_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitHang_(s)
	}
}

func (p *mumpsParser) Hang_() (localctx IHang_Context) {
	this := p
	_ = this

	localctx = NewHang_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, mumpsParserRULE_hang_)
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
		p.SetState(312)
		p.Match(mumpsParserHANG)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCOLON {
		{
			p.SetState(313)
			p.Postcondition()
		}

	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(316)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(321)
		p.Term()
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
	p.RuleIndex = mumpsParserRULE_if_
	return p
}

func (*If_Context) IsIf_Context() {}

func NewIf_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_Context {
	var p = new(If_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_if_

	return p
}

func (s *If_Context) GetParser() antlr.Parser { return s.parser }

func (s *If_Context) IF() antlr.TerminalNode {
	return s.GetToken(mumpsParserIF, 0)
}

func (s *If_Context) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *If_Context) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *If_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *If_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *If_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterIf_(s)
	}
}

func (s *If_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitIf_(s)
	}
}

func (p *mumpsParser) If_() (localctx IIf_Context) {
	this := p
	_ = this

	localctx = NewIf_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, mumpsParserRULE_if_)
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
		p.Match(mumpsParserIF)
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(324)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(329)
		p.Condition()
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserSPACE {
		{
			p.SetState(330)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(336)
		p.Command()
	}

	return localctx
}

// IKill_Context is an interface to support dynamic dispatch.
type IKill_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKill_Context differentiates from other interfaces.
	IsKill_Context()
}

type Kill_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKill_Context() *Kill_Context {
	var p = new(Kill_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_kill_
	return p
}

func (*Kill_Context) IsKill_Context() {}

func NewKill_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kill_Context {
	var p = new(Kill_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_kill_

	return p
}

func (s *Kill_Context) GetParser() antlr.Parser { return s.parser }

func (s *Kill_Context) KILL() antlr.TerminalNode {
	return s.GetToken(mumpsParserKILL, 0)
}

func (s *Kill_Context) Arglist() IArglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *Kill_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Kill_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *Kill_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *Kill_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kill_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kill_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterKill_(s)
	}
}

func (s *Kill_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitKill_(s)
	}
}

func (p *mumpsParser) Kill_() (localctx IKill_Context) {
	this := p
	_ = this

	localctx = NewKill_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, mumpsParserRULE_kill_)
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
		p.SetState(338)
		p.Match(mumpsParserKILL)
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCOLON {
		{
			p.SetState(339)
			p.Postcondition()
		}

	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(342)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(347)
		p.Arglist()
	}

	return localctx
}

// IMerge_Context is an interface to support dynamic dispatch.
type IMerge_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMerge_Context differentiates from other interfaces.
	IsMerge_Context()
}

type Merge_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMerge_Context() *Merge_Context {
	var p = new(Merge_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_merge_
	return p
}

func (*Merge_Context) IsMerge_Context() {}

func NewMerge_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Merge_Context {
	var p = new(Merge_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_merge_

	return p
}

func (s *Merge_Context) GetParser() antlr.Parser { return s.parser }

func (s *Merge_Context) MERGE() antlr.TerminalNode {
	return s.GetToken(mumpsParserMERGE, 0)
}

func (s *Merge_Context) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Merge_Context) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Merge_Context) AllEQUALS() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserEQUALS)
}

func (s *Merge_Context) EQUALS(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserEQUALS, i)
}

func (s *Merge_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Merge_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *Merge_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *Merge_Context) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserCOMMA)
}

func (s *Merge_Context) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserCOMMA, i)
}

func (s *Merge_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Merge_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Merge_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterMerge_(s)
	}
}

func (s *Merge_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitMerge_(s)
	}
}

func (p *mumpsParser) Merge_() (localctx IMerge_Context) {
	this := p
	_ = this

	localctx = NewMerge_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, mumpsParserRULE_merge_)
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
		p.SetState(349)
		p.Match(mumpsParserMERGE)
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCOLON {
		{
			p.SetState(350)
			p.Postcondition()
		}

	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(353)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(358)
		p.Term()
	}
	{
		p.SetState(359)
		p.Match(mumpsParserEQUALS)
	}
	{
		p.SetState(360)
		p.Term()
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserCOMMA {
		{
			p.SetState(361)
			p.Match(mumpsParserCOMMA)
		}
		{
			p.SetState(362)
			p.Term()
		}
		{
			p.SetState(363)
			p.Match(mumpsParserEQUALS)
		}
		{
			p.SetState(364)
			p.Term()
		}

		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INew_Context is an interface to support dynamic dispatch.
type INew_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNew_Context differentiates from other interfaces.
	IsNew_Context()
}

type New_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNew_Context() *New_Context {
	var p = new(New_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_new_
	return p
}

func (*New_Context) IsNew_Context() {}

func NewNew_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *New_Context {
	var p = new(New_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_new_

	return p
}

func (s *New_Context) GetParser() antlr.Parser { return s.parser }

func (s *New_Context) NEW() antlr.TerminalNode {
	return s.GetToken(mumpsParserNEW, 0)
}

func (s *New_Context) Arglist() IArglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *New_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *New_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *New_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *New_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *New_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *New_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterNew_(s)
	}
}

func (s *New_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitNew_(s)
	}
}

func (p *mumpsParser) New_() (localctx INew_Context) {
	this := p
	_ = this

	localctx = NewNew_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, mumpsParserRULE_new_)
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
		p.SetState(371)
		p.Match(mumpsParserNEW)
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCOLON {
		{
			p.SetState(372)
			p.Postcondition()
		}

	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(375)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(380)
		p.Arglist()
	}

	return localctx
}

// IQuit_Context is an interface to support dynamic dispatch.
type IQuit_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuit_Context differentiates from other interfaces.
	IsQuit_Context()
}

type Quit_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuit_Context() *Quit_Context {
	var p = new(Quit_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_quit_
	return p
}

func (*Quit_Context) IsQuit_Context() {}

func NewQuit_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Quit_Context {
	var p = new(Quit_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_quit_

	return p
}

func (s *Quit_Context) GetParser() antlr.Parser { return s.parser }

func (s *Quit_Context) QUIT() antlr.TerminalNode {
	return s.GetToken(mumpsParserQUIT, 0)
}

func (s *Quit_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Quit_Context) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Quit_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *Quit_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *Quit_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Quit_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Quit_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterQuit_(s)
	}
}

func (s *Quit_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitQuit_(s)
	}
}

func (p *mumpsParser) Quit_() (localctx IQuit_Context) {
	this := p
	_ = this

	localctx = NewQuit_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, mumpsParserRULE_quit_)
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
		p.SetState(382)
		p.Match(mumpsParserQUIT)
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(383)
			p.Postcondition()
		}

	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == mumpsParserSPACE {
			{
				p.SetState(386)
				p.Match(mumpsParserSPACE)
			}

			p.SetState(389)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(391)
			p.Term()
		}

	}

	return localctx
}

// IRead_Context is an interface to support dynamic dispatch.
type IRead_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRead_Context differentiates from other interfaces.
	IsRead_Context()
}

type Read_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRead_Context() *Read_Context {
	var p = new(Read_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_read_
	return p
}

func (*Read_Context) IsRead_Context() {}

func NewRead_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Read_Context {
	var p = new(Read_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_read_

	return p
}

func (s *Read_Context) GetParser() antlr.Parser { return s.parser }

func (s *Read_Context) READ() antlr.TerminalNode {
	return s.GetToken(mumpsParserREAD, 0)
}

func (s *Read_Context) Arglist() IArglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *Read_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Read_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *Read_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *Read_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Read_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Read_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterRead_(s)
	}
}

func (s *Read_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitRead_(s)
	}
}

func (p *mumpsParser) Read_() (localctx IRead_Context) {
	this := p
	_ = this

	localctx = NewRead_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, mumpsParserRULE_read_)
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
		p.SetState(394)
		p.Match(mumpsParserREAD)
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCOLON {
		{
			p.SetState(395)
			p.Postcondition()
		}

	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(398)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(403)
		p.Arglist()
	}

	return localctx
}

// ISet_Context is an interface to support dynamic dispatch.
type ISet_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_Context differentiates from other interfaces.
	IsSet_Context()
}

type Set_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_Context() *Set_Context {
	var p = new(Set_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_set_
	return p
}

func (*Set_Context) IsSet_Context() {}

func NewSet_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_Context {
	var p = new(Set_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_set_

	return p
}

func (s *Set_Context) GetParser() antlr.Parser { return s.parser }

func (s *Set_Context) SET() antlr.TerminalNode {
	return s.GetToken(mumpsParserSET, 0)
}

func (s *Set_Context) AllAssign() []IAssignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignContext)(nil)).Elem())
	var tst = make([]IAssignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignContext)
		}
	}

	return tst
}

func (s *Set_Context) Assign(i int) IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *Set_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Set_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *Set_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *Set_Context) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserCOMMA)
}

func (s *Set_Context) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserCOMMA, i)
}

func (s *Set_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterSet_(s)
	}
}

func (s *Set_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitSet_(s)
	}
}

func (p *mumpsParser) Set_() (localctx ISet_Context) {
	this := p
	_ = this

	localctx = NewSet_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, mumpsParserRULE_set_)
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
		p.SetState(405)
		p.Match(mumpsParserSET)
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCOLON {
		{
			p.SetState(406)
			p.Postcondition()
		}

	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(409)
				p.Match(mumpsParserSPACE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext())
	}
	{
		p.SetState(414)
		p.Assign()
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserCOMMA {
		{
			p.SetState(415)
			p.Match(mumpsParserCOMMA)
		}
		{
			p.SetState(416)
			p.Assign()
		}

		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IView_Context is an interface to support dynamic dispatch.
type IView_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsView_Context differentiates from other interfaces.
	IsView_Context()
}

type View_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyView_Context() *View_Context {
	var p = new(View_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_view_
	return p
}

func (*View_Context) IsView_Context() {}

func NewView_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *View_Context {
	var p = new(View_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_view_

	return p
}

func (s *View_Context) GetParser() antlr.Parser { return s.parser }

func (s *View_Context) VIEW() antlr.TerminalNode {
	return s.GetToken(mumpsParserVIEW, 0)
}

func (s *View_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(mumpsParserIDENTIFIER, 0)
}

func (s *View_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *View_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *View_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *View_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *View_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *View_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterView_(s)
	}
}

func (s *View_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitView_(s)
	}
}

func (p *mumpsParser) View_() (localctx IView_Context) {
	this := p
	_ = this

	localctx = NewView_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, mumpsParserRULE_view_)
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
		p.SetState(422)
		p.Match(mumpsParserVIEW)
	}
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCOLON {
		{
			p.SetState(423)
			p.Postcondition()
		}

	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(426)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(431)
		p.Match(mumpsParserIDENTIFIER)
	}

	return localctx
}

// IWrite_Context is an interface to support dynamic dispatch.
type IWrite_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWrite_Context differentiates from other interfaces.
	IsWrite_Context()
}

type Write_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWrite_Context() *Write_Context {
	var p = new(Write_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_write_
	return p
}

func (*Write_Context) IsWrite_Context() {}

func NewWrite_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Write_Context {
	var p = new(Write_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_write_

	return p
}

func (s *Write_Context) GetParser() antlr.Parser { return s.parser }

func (s *Write_Context) WRITE() antlr.TerminalNode {
	return s.GetToken(mumpsParserWRITE, 0)
}

func (s *Write_Context) Arglist() IArglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *Write_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Write_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *Write_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *Write_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Write_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Write_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterWrite_(s)
	}
}

func (s *Write_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitWrite_(s)
	}
}

func (p *mumpsParser) Write_() (localctx IWrite_Context) {
	this := p
	_ = this

	localctx = NewWrite_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, mumpsParserRULE_write_)
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
		p.SetState(433)
		p.Match(mumpsParserWRITE)
	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCOLON {
		{
			p.SetState(434)
			p.Postcondition()
		}

	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(437)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(442)
		p.Arglist()
	}

	return localctx
}

// IXecute_Context is an interface to support dynamic dispatch.
type IXecute_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXecute_Context differentiates from other interfaces.
	IsXecute_Context()
}

type Xecute_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXecute_Context() *Xecute_Context {
	var p = new(Xecute_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_xecute_
	return p
}

func (*Xecute_Context) IsXecute_Context() {}

func NewXecute_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Xecute_Context {
	var p = new(Xecute_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_xecute_

	return p
}

func (s *Xecute_Context) GetParser() antlr.Parser { return s.parser }

func (s *Xecute_Context) XECUTE() antlr.TerminalNode {
	return s.GetToken(mumpsParserXECUTE, 0)
}

func (s *Xecute_Context) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(mumpsParserSTRING_LITERAL, 0)
}

func (s *Xecute_Context) Postcondition() IPostconditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostconditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostconditionContext)
}

func (s *Xecute_Context) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *Xecute_Context) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *Xecute_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Xecute_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Xecute_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterXecute_(s)
	}
}

func (s *Xecute_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitXecute_(s)
	}
}

func (p *mumpsParser) Xecute_() (localctx IXecute_Context) {
	this := p
	_ = this

	localctx = NewXecute_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, mumpsParserRULE_xecute_)
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
		p.SetState(444)
		p.Match(mumpsParserXECUTE)
	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mumpsParserCOLON {
		{
			p.SetState(445)
			p.Postcondition()
		}

	}
	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mumpsParserSPACE {
		{
			p.SetState(448)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(451)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(453)
		p.Match(mumpsParserSTRING_LITERAL)
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(mumpsParserEQUALS, 0)
}

func (s *AssignContext) Arg() IArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *AssignContext) Arglist() IArglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *AssignContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *AssignContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *AssignContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserLPAREN, 0)
}

func (s *AssignContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mumpsParserRPAREN, 0)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *mumpsParser) Assign() (localctx IAssignContext) {
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, mumpsParserRULE_assign)
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
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(mumpsParserDOLLAR-29))|(1<<(mumpsParserAMPERSAND-29))|(1<<(mumpsParserCARAT-29))|(1<<(mumpsParserBANG-29))|(1<<(mumpsParserLPAREN-29))|(1<<(mumpsParserIDENTIFIER-29))|(1<<(mumpsParserSTRING_LITERAL-29))|(1<<(mumpsParserNUMBER-29)))) != 0 {
		p.SetState(456)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(455)
				p.Match(mumpsParserLPAREN)
			}

		}
		{
			p.SetState(458)
			p.Arglist()
		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == mumpsParserRPAREN {
			{
				p.SetState(459)
				p.Match(mumpsParserRPAREN)
			}

		}

	}
	p.SetState(467)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserSPACE {
		{
			p.SetState(464)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(470)
		p.Match(mumpsParserEQUALS)
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mumpsParserSPACE {
		{
			p.SetState(471)
			p.Match(mumpsParserSPACE)
		}

		p.SetState(476)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(477)
		p.Arg()
	}

	return localctx
}

// IArglistContext is an interface to support dynamic dispatch.
type IArglistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArglistContext differentiates from other interfaces.
	IsArglistContext()
}

type ArglistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArglistContext() *ArglistContext {
	var p = new(ArglistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_arglist
	return p
}

func (*ArglistContext) IsArglistContext() {}

func NewArglistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArglistContext {
	var p = new(ArglistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_arglist

	return p
}

func (s *ArglistContext) GetParser() antlr.Parser { return s.parser }

func (s *ArglistContext) AllArg() []IArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgContext)(nil)).Elem())
	var tst = make([]IArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgContext)
		}
	}

	return tst
}

func (s *ArglistContext) Arg(i int) IArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *ArglistContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserCOMMA)
}

func (s *ArglistContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserCOMMA, i)
}

func (s *ArglistContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(mumpsParserSPACE)
}

func (s *ArglistContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(mumpsParserSPACE, i)
}

func (s *ArglistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArglistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArglistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterArglist(s)
	}
}

func (s *ArglistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitArglist(s)
	}
}

func (p *mumpsParser) Arglist() (localctx IArglistContext) {
	this := p
	_ = this

	localctx = NewArglistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, mumpsParserRULE_arglist)
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
		p.SetState(479)
		p.Arg()
	}
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(483)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == mumpsParserSPACE {
				{
					p.SetState(480)
					p.Match(mumpsParserSPACE)
				}

				p.SetState(485)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(486)
				p.Match(mumpsParserCOMMA)
			}
			{
				p.SetState(487)
				p.Arg()
			}

		}
		p.SetState(492)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext())
	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mumpsParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mumpsParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgContext) BANG() antlr.TerminalNode {
	return s.GetToken(mumpsParserBANG, 0)
}

func (s *ArgContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(mumpsParserSTRING_LITERAL, 0)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mumpsListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *mumpsParser) Arg() (localctx IArgContext) {
	this := p
	_ = this

	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, mumpsParserRULE_arg)
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

	p.SetState(495)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mumpsParserDOLLAR, mumpsParserAMPERSAND, mumpsParserCARAT, mumpsParserLPAREN, mumpsParserIDENTIFIER, mumpsParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(493)
			p.Expression()
		}

	case mumpsParserBANG, mumpsParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(494)
			_la = p.GetTokenStream().LA(1)

			if !(_la == mumpsParserBANG || _la == mumpsParserSTRING_LITERAL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
