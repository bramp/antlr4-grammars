// Code generated from R.g4 by ANTLR 4.7.2. DO NOT EDIT.

package r

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 66, 533,
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
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3,
	39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3,
	53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55,
	3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 6, 57, 327, 10, 57, 13, 57, 14, 57,
	328, 3, 57, 5, 57, 332, 10, 57, 3, 58, 6, 58, 335, 10, 58, 13, 58, 14,
	58, 336, 3, 58, 5, 58, 340, 10, 58, 3, 59, 3, 59, 3, 60, 6, 60, 345, 10,
	60, 13, 60, 14, 60, 346, 3, 60, 3, 60, 7, 60, 351, 10, 60, 12, 60, 14,
	60, 354, 11, 60, 3, 60, 5, 60, 357, 10, 60, 3, 60, 5, 60, 360, 10, 60,
	3, 60, 6, 60, 363, 10, 60, 13, 60, 14, 60, 364, 3, 60, 5, 60, 368, 10,
	60, 3, 60, 5, 60, 371, 10, 60, 3, 60, 3, 60, 6, 60, 375, 10, 60, 13, 60,
	14, 60, 376, 3, 60, 5, 60, 380, 10, 60, 3, 60, 5, 60, 383, 10, 60, 5, 60,
	385, 10, 60, 3, 61, 3, 61, 3, 62, 3, 62, 5, 62, 391, 10, 62, 3, 62, 3,
	62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 5, 63, 401, 10, 63, 3, 64,
	3, 64, 3, 64, 7, 64, 406, 10, 64, 12, 64, 14, 64, 409, 11, 64, 3, 64, 3,
	64, 3, 64, 3, 64, 7, 64, 415, 10, 64, 12, 64, 14, 64, 418, 11, 64, 3, 64,
	3, 64, 3, 64, 3, 64, 7, 64, 424, 10, 64, 12, 64, 14, 64, 427, 11, 64, 3,
	64, 5, 64, 430, 10, 64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 5, 65, 437,
	10, 65, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66,
	3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 5, 66, 455, 10, 66, 3,
	67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 5, 67, 466,
	10, 67, 3, 68, 3, 68, 3, 68, 5, 68, 471, 10, 68, 3, 69, 3, 69, 3, 69, 5,
	69, 476, 10, 69, 3, 69, 3, 69, 3, 69, 7, 69, 481, 10, 69, 12, 69, 14, 69,
	484, 11, 69, 3, 69, 3, 69, 3, 69, 3, 69, 7, 69, 490, 10, 69, 12, 69, 14,
	69, 493, 11, 69, 5, 69, 495, 10, 69, 3, 70, 3, 70, 3, 71, 3, 71, 7, 71,
	501, 10, 71, 12, 71, 14, 71, 504, 11, 71, 3, 71, 3, 71, 3, 72, 3, 72, 7,
	72, 510, 10, 72, 12, 72, 14, 72, 513, 11, 72, 3, 72, 5, 72, 516, 10, 72,
	3, 72, 3, 72, 3, 72, 3, 72, 3, 73, 5, 73, 523, 10, 73, 3, 73, 3, 73, 3,
	74, 6, 74, 528, 10, 74, 13, 74, 14, 74, 529, 3, 74, 3, 74, 7, 407, 416,
	425, 502, 511, 2, 75, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35,
	19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53,
	28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71,
	37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89,
	46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54,
	107, 55, 109, 56, 111, 57, 113, 58, 115, 59, 117, 2, 119, 60, 121, 2, 123,
	2, 125, 61, 127, 62, 129, 2, 131, 2, 133, 2, 135, 2, 137, 63, 139, 2, 141,
	64, 143, 2, 145, 65, 147, 66, 3, 2, 15, 4, 2, 90, 90, 122, 122, 4, 2, 78,
	78, 110, 110, 5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 71, 71, 103, 103, 4,
	2, 45, 45, 47, 47, 4, 2, 36, 36, 94, 94, 4, 2, 41, 41, 94, 94, 11, 2, 36,
	36, 41, 41, 94, 94, 99, 100, 104, 104, 112, 112, 116, 116, 118, 118, 120,
	120, 3, 2, 50, 53, 3, 2, 50, 57, 4, 2, 48, 48, 97, 97, 4, 2, 67, 92, 99,
	124, 5, 2, 11, 11, 14, 14, 34, 34, 2, 570, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
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
	2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 125, 3,
	2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2,
	143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 3, 149, 3, 2,
	2, 2, 5, 151, 3, 2, 2, 2, 7, 154, 3, 2, 2, 2, 9, 156, 3, 2, 2, 2, 11, 158,
	3, 2, 2, 2, 13, 161, 3, 2, 2, 2, 15, 165, 3, 2, 2, 2, 17, 167, 3, 2, 2,
	2, 19, 169, 3, 2, 2, 2, 21, 171, 3, 2, 2, 2, 23, 173, 3, 2, 2, 2, 25, 175,
	3, 2, 2, 2, 27, 177, 3, 2, 2, 2, 29, 179, 3, 2, 2, 2, 31, 181, 3, 2, 2,
	2, 33, 183, 3, 2, 2, 2, 35, 186, 3, 2, 2, 2, 37, 188, 3, 2, 2, 2, 39, 191,
	3, 2, 2, 2, 41, 194, 3, 2, 2, 2, 43, 197, 3, 2, 2, 2, 45, 199, 3, 2, 2,
	2, 47, 201, 3, 2, 2, 2, 49, 204, 3, 2, 2, 2, 51, 206, 3, 2, 2, 2, 53, 209,
	3, 2, 2, 2, 55, 211, 3, 2, 2, 2, 57, 214, 3, 2, 2, 2, 59, 218, 3, 2, 2,
	2, 61, 220, 3, 2, 2, 2, 63, 223, 3, 2, 2, 2, 65, 227, 3, 2, 2, 2, 67, 230,
	3, 2, 2, 2, 69, 239, 3, 2, 2, 2, 71, 241, 3, 2, 2, 2, 73, 243, 3, 2, 2,
	2, 75, 245, 3, 2, 2, 2, 77, 247, 3, 2, 2, 2, 79, 250, 3, 2, 2, 2, 81, 255,
	3, 2, 2, 2, 83, 259, 3, 2, 2, 2, 85, 262, 3, 2, 2, 2, 87, 268, 3, 2, 2,
	2, 89, 275, 3, 2, 2, 2, 91, 277, 3, 2, 2, 2, 93, 282, 3, 2, 2, 2, 95, 288,
	3, 2, 2, 2, 97, 293, 3, 2, 2, 2, 99, 296, 3, 2, 2, 2, 101, 300, 3, 2, 2,
	2, 103, 304, 3, 2, 2, 2, 105, 309, 3, 2, 2, 2, 107, 315, 3, 2, 2, 2, 109,
	317, 3, 2, 2, 2, 111, 321, 3, 2, 2, 2, 113, 323, 3, 2, 2, 2, 115, 334,
	3, 2, 2, 2, 117, 341, 3, 2, 2, 2, 119, 384, 3, 2, 2, 2, 121, 386, 3, 2,
	2, 2, 123, 388, 3, 2, 2, 2, 125, 400, 3, 2, 2, 2, 127, 429, 3, 2, 2, 2,
	129, 436, 3, 2, 2, 2, 131, 454, 3, 2, 2, 2, 133, 465, 3, 2, 2, 2, 135,
	467, 3, 2, 2, 2, 137, 494, 3, 2, 2, 2, 139, 496, 3, 2, 2, 2, 141, 498,
	3, 2, 2, 2, 143, 507, 3, 2, 2, 2, 145, 522, 3, 2, 2, 2, 147, 527, 3, 2,
	2, 2, 149, 150, 7, 61, 2, 2, 150, 4, 3, 2, 2, 2, 151, 152, 7, 93, 2, 2,
	152, 153, 7, 93, 2, 2, 153, 6, 3, 2, 2, 2, 154, 155, 7, 95, 2, 2, 155,
	8, 3, 2, 2, 2, 156, 157, 7, 93, 2, 2, 157, 10, 3, 2, 2, 2, 158, 159, 7,
	60, 2, 2, 159, 160, 7, 60, 2, 2, 160, 12, 3, 2, 2, 2, 161, 162, 7, 60,
	2, 2, 162, 163, 7, 60, 2, 2, 163, 164, 7, 60, 2, 2, 164, 14, 3, 2, 2, 2,
	165, 166, 7, 38, 2, 2, 166, 16, 3, 2, 2, 2, 167, 168, 7, 66, 2, 2, 168,
	18, 3, 2, 2, 2, 169, 170, 7, 96, 2, 2, 170, 20, 3, 2, 2, 2, 171, 172, 7,
	47, 2, 2, 172, 22, 3, 2, 2, 2, 173, 174, 7, 45, 2, 2, 174, 24, 3, 2, 2,
	2, 175, 176, 7, 60, 2, 2, 176, 26, 3, 2, 2, 2, 177, 178, 7, 44, 2, 2, 178,
	28, 3, 2, 2, 2, 179, 180, 7, 49, 2, 2, 180, 30, 3, 2, 2, 2, 181, 182, 7,
	64, 2, 2, 182, 32, 3, 2, 2, 2, 183, 184, 7, 64, 2, 2, 184, 185, 7, 63,
	2, 2, 185, 34, 3, 2, 2, 2, 186, 187, 7, 62, 2, 2, 187, 36, 3, 2, 2, 2,
	188, 189, 7, 62, 2, 2, 189, 190, 7, 63, 2, 2, 190, 38, 3, 2, 2, 2, 191,
	192, 7, 63, 2, 2, 192, 193, 7, 63, 2, 2, 193, 40, 3, 2, 2, 2, 194, 195,
	7, 35, 2, 2, 195, 196, 7, 63, 2, 2, 196, 42, 3, 2, 2, 2, 197, 198, 7, 35,
	2, 2, 198, 44, 3, 2, 2, 2, 199, 200, 7, 40, 2, 2, 200, 46, 3, 2, 2, 2,
	201, 202, 7, 40, 2, 2, 202, 203, 7, 40, 2, 2, 203, 48, 3, 2, 2, 2, 204,
	205, 7, 126, 2, 2, 205, 50, 3, 2, 2, 2, 206, 207, 7, 126, 2, 2, 207, 208,
	7, 126, 2, 2, 208, 52, 3, 2, 2, 2, 209, 210, 7, 128, 2, 2, 210, 54, 3,
	2, 2, 2, 211, 212, 7, 62, 2, 2, 212, 213, 7, 47, 2, 2, 213, 56, 3, 2, 2,
	2, 214, 215, 7, 62, 2, 2, 215, 216, 7, 62, 2, 2, 216, 217, 7, 47, 2, 2,
	217, 58, 3, 2, 2, 2, 218, 219, 7, 63, 2, 2, 219, 60, 3, 2, 2, 2, 220, 221,
	7, 47, 2, 2, 221, 222, 7, 64, 2, 2, 222, 62, 3, 2, 2, 2, 223, 224, 7, 47,
	2, 2, 224, 225, 7, 64, 2, 2, 225, 226, 7, 64, 2, 2, 226, 64, 3, 2, 2, 2,
	227, 228, 7, 60, 2, 2, 228, 229, 7, 63, 2, 2, 229, 66, 3, 2, 2, 2, 230,
	231, 7, 104, 2, 2, 231, 232, 7, 119, 2, 2, 232, 233, 7, 112, 2, 2, 233,
	234, 7, 101, 2, 2, 234, 235, 7, 118, 2, 2, 235, 236, 7, 107, 2, 2, 236,
	237, 7, 113, 2, 2, 237, 238, 7, 112, 2, 2, 238, 68, 3, 2, 2, 2, 239, 240,
	7, 42, 2, 2, 240, 70, 3, 2, 2, 2, 241, 242, 7, 43, 2, 2, 242, 72, 3, 2,
	2, 2, 243, 244, 7, 125, 2, 2, 244, 74, 3, 2, 2, 2, 245, 246, 7, 127, 2,
	2, 246, 76, 3, 2, 2, 2, 247, 248, 7, 107, 2, 2, 248, 249, 7, 104, 2, 2,
	249, 78, 3, 2, 2, 2, 250, 251, 7, 103, 2, 2, 251, 252, 7, 110, 2, 2, 252,
	253, 7, 117, 2, 2, 253, 254, 7, 103, 2, 2, 254, 80, 3, 2, 2, 2, 255, 256,
	7, 104, 2, 2, 256, 257, 7, 113, 2, 2, 257, 258, 7, 116, 2, 2, 258, 82,
	3, 2, 2, 2, 259, 260, 7, 107, 2, 2, 260, 261, 7, 112, 2, 2, 261, 84, 3,
	2, 2, 2, 262, 263, 7, 121, 2, 2, 263, 264, 7, 106, 2, 2, 264, 265, 7, 107,
	2, 2, 265, 266, 7, 110, 2, 2, 266, 267, 7, 103, 2, 2, 267, 86, 3, 2, 2,
	2, 268, 269, 7, 116, 2, 2, 269, 270, 7, 103, 2, 2, 270, 271, 7, 114, 2,
	2, 271, 272, 7, 103, 2, 2, 272, 273, 7, 99, 2, 2, 273, 274, 7, 118, 2,
	2, 274, 88, 3, 2, 2, 2, 275, 276, 7, 65, 2, 2, 276, 90, 3, 2, 2, 2, 277,
	278, 7, 112, 2, 2, 278, 279, 7, 103, 2, 2, 279, 280, 7, 122, 2, 2, 280,
	281, 7, 118, 2, 2, 281, 92, 3, 2, 2, 2, 282, 283, 7, 100, 2, 2, 283, 284,
	7, 116, 2, 2, 284, 285, 7, 103, 2, 2, 285, 286, 7, 99, 2, 2, 286, 287,
	7, 109, 2, 2, 287, 94, 3, 2, 2, 2, 288, 289, 7, 80, 2, 2, 289, 290, 7,
	87, 2, 2, 290, 291, 7, 78, 2, 2, 291, 292, 7, 78, 2, 2, 292, 96, 3, 2,
	2, 2, 293, 294, 7, 80, 2, 2, 294, 295, 7, 67, 2, 2, 295, 98, 3, 2, 2, 2,
	296, 297, 7, 75, 2, 2, 297, 298, 7, 112, 2, 2, 298, 299, 7, 104, 2, 2,
	299, 100, 3, 2, 2, 2, 300, 301, 7, 80, 2, 2, 301, 302, 7, 99, 2, 2, 302,
	303, 7, 80, 2, 2, 303, 102, 3, 2, 2, 2, 304, 305, 7, 86, 2, 2, 305, 306,
	7, 84, 2, 2, 306, 307, 7, 87, 2, 2, 307, 308, 7, 71, 2, 2, 308, 104, 3,
	2, 2, 2, 309, 310, 7, 72, 2, 2, 310, 311, 7, 67, 2, 2, 311, 312, 7, 78,
	2, 2, 312, 313, 7, 85, 2, 2, 313, 314, 7, 71, 2, 2, 314, 106, 3, 2, 2,
	2, 315, 316, 7, 46, 2, 2, 316, 108, 3, 2, 2, 2, 317, 318, 7, 48, 2, 2,
	318, 319, 7, 48, 2, 2, 319, 320, 7, 48, 2, 2, 320, 110, 3, 2, 2, 2, 321,
	322, 7, 48, 2, 2, 322, 112, 3, 2, 2, 2, 323, 324, 7, 50, 2, 2, 324, 326,
	9, 2, 2, 2, 325, 327, 5, 117, 59, 2, 326, 325, 3, 2, 2, 2, 327, 328, 3,
	2, 2, 2, 328, 326, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 331, 3, 2, 2,
	2, 330, 332, 9, 3, 2, 2, 331, 330, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332,
	114, 3, 2, 2, 2, 333, 335, 5, 121, 61, 2, 334, 333, 3, 2, 2, 2, 335, 336,
	3, 2, 2, 2, 336, 334, 3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 339, 3, 2,
	2, 2, 338, 340, 9, 3, 2, 2, 339, 338, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2,
	340, 116, 3, 2, 2, 2, 341, 342, 9, 4, 2, 2, 342, 118, 3, 2, 2, 2, 343,
	345, 5, 121, 61, 2, 344, 343, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346, 344,
	3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 352, 7, 48,
	2, 2, 349, 351, 5, 121, 61, 2, 350, 349, 3, 2, 2, 2, 351, 354, 3, 2, 2,
	2, 352, 350, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 356, 3, 2, 2, 2, 354,
	352, 3, 2, 2, 2, 355, 357, 5, 123, 62, 2, 356, 355, 3, 2, 2, 2, 356, 357,
	3, 2, 2, 2, 357, 359, 3, 2, 2, 2, 358, 360, 9, 3, 2, 2, 359, 358, 3, 2,
	2, 2, 359, 360, 3, 2, 2, 2, 360, 385, 3, 2, 2, 2, 361, 363, 5, 121, 61,
	2, 362, 361, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 364,
	365, 3, 2, 2, 2, 365, 367, 3, 2, 2, 2, 366, 368, 5, 123, 62, 2, 367, 366,
	3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 370, 3, 2, 2, 2, 369, 371, 9, 3,
	2, 2, 370, 369, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 385, 3, 2, 2, 2,
	372, 374, 7, 48, 2, 2, 373, 375, 5, 121, 61, 2, 374, 373, 3, 2, 2, 2, 375,
	376, 3, 2, 2, 2, 376, 374, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 379,
	3, 2, 2, 2, 378, 380, 5, 123, 62, 2, 379, 378, 3, 2, 2, 2, 379, 380, 3,
	2, 2, 2, 380, 382, 3, 2, 2, 2, 381, 383, 9, 3, 2, 2, 382, 381, 3, 2, 2,
	2, 382, 383, 3, 2, 2, 2, 383, 385, 3, 2, 2, 2, 384, 344, 3, 2, 2, 2, 384,
	362, 3, 2, 2, 2, 384, 372, 3, 2, 2, 2, 385, 120, 3, 2, 2, 2, 386, 387,
	4, 50, 59, 2, 387, 122, 3, 2, 2, 2, 388, 390, 9, 5, 2, 2, 389, 391, 9,
	6, 2, 2, 390, 389, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 392, 3, 2, 2,
	2, 392, 393, 5, 115, 58, 2, 393, 124, 3, 2, 2, 2, 394, 395, 5, 115, 58,
	2, 395, 396, 7, 107, 2, 2, 396, 401, 3, 2, 2, 2, 397, 398, 5, 119, 60,
	2, 398, 399, 7, 107, 2, 2, 399, 401, 3, 2, 2, 2, 400, 394, 3, 2, 2, 2,
	400, 397, 3, 2, 2, 2, 401, 126, 3, 2, 2, 2, 402, 407, 7, 36, 2, 2, 403,
	406, 5, 129, 65, 2, 404, 406, 10, 7, 2, 2, 405, 403, 3, 2, 2, 2, 405, 404,
	3, 2, 2, 2, 406, 409, 3, 2, 2, 2, 407, 408, 3, 2, 2, 2, 407, 405, 3, 2,
	2, 2, 408, 410, 3, 2, 2, 2, 409, 407, 3, 2, 2, 2, 410, 430, 7, 36, 2, 2,
	411, 416, 7, 41, 2, 2, 412, 415, 5, 129, 65, 2, 413, 415, 10, 8, 2, 2,
	414, 412, 3, 2, 2, 2, 414, 413, 3, 2, 2, 2, 415, 418, 3, 2, 2, 2, 416,
	417, 3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 417, 419, 3, 2, 2, 2, 418, 416,
	3, 2, 2, 2, 419, 430, 7, 41, 2, 2, 420, 425, 7, 98, 2, 2, 421, 424, 5,
	129, 65, 2, 422, 424, 10, 8, 2, 2, 423, 421, 3, 2, 2, 2, 423, 422, 3, 2,
	2, 2, 424, 427, 3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2,
	426, 428, 3, 2, 2, 2, 427, 425, 3, 2, 2, 2, 428, 430, 7, 98, 2, 2, 429,
	402, 3, 2, 2, 2, 429, 411, 3, 2, 2, 2, 429, 420, 3, 2, 2, 2, 430, 128,
	3, 2, 2, 2, 431, 432, 7, 94, 2, 2, 432, 437, 9, 9, 2, 2, 433, 437, 5, 131,
	66, 2, 434, 437, 5, 135, 68, 2, 435, 437, 5, 133, 67, 2, 436, 431, 3, 2,
	2, 2, 436, 433, 3, 2, 2, 2, 436, 434, 3, 2, 2, 2, 436, 435, 3, 2, 2, 2,
	437, 130, 3, 2, 2, 2, 438, 439, 7, 94, 2, 2, 439, 440, 7, 119, 2, 2, 440,
	441, 5, 117, 59, 2, 441, 442, 5, 117, 59, 2, 442, 443, 5, 117, 59, 2, 443,
	444, 5, 117, 59, 2, 444, 455, 3, 2, 2, 2, 445, 446, 7, 94, 2, 2, 446, 447,
	7, 119, 2, 2, 447, 448, 7, 125, 2, 2, 448, 449, 5, 117, 59, 2, 449, 450,
	5, 117, 59, 2, 450, 451, 5, 117, 59, 2, 451, 452, 5, 117, 59, 2, 452, 453,
	7, 127, 2, 2, 453, 455, 3, 2, 2, 2, 454, 438, 3, 2, 2, 2, 454, 445, 3,
	2, 2, 2, 455, 132, 3, 2, 2, 2, 456, 457, 7, 94, 2, 2, 457, 458, 9, 10,
	2, 2, 458, 459, 9, 11, 2, 2, 459, 466, 9, 11, 2, 2, 460, 461, 7, 94, 2,
	2, 461, 462, 9, 11, 2, 2, 462, 466, 9, 11, 2, 2, 463, 464, 7, 94, 2, 2,
	464, 466, 9, 11, 2, 2, 465, 456, 3, 2, 2, 2, 465, 460, 3, 2, 2, 2, 465,
	463, 3, 2, 2, 2, 466, 134, 3, 2, 2, 2, 467, 468, 7, 94, 2, 2, 468, 470,
	5, 117, 59, 2, 469, 471, 5, 117, 59, 2, 470, 469, 3, 2, 2, 2, 470, 471,
	3, 2, 2, 2, 471, 136, 3, 2, 2, 2, 472, 475, 7, 48, 2, 2, 473, 476, 5, 139,
	70, 2, 474, 476, 9, 12, 2, 2, 475, 473, 3, 2, 2, 2, 475, 474, 3, 2, 2,
	2, 476, 482, 3, 2, 2, 2, 477, 481, 5, 139, 70, 2, 478, 481, 5, 121, 61,
	2, 479, 481, 9, 12, 2, 2, 480, 477, 3, 2, 2, 2, 480, 478, 3, 2, 2, 2, 480,
	479, 3, 2, 2, 2, 481, 484, 3, 2, 2, 2, 482, 480, 3, 2, 2, 2, 482, 483,
	3, 2, 2, 2, 483, 495, 3, 2, 2, 2, 484, 482, 3, 2, 2, 2, 485, 491, 5, 139,
	70, 2, 486, 490, 5, 139, 70, 2, 487, 490, 5, 121, 61, 2, 488, 490, 9, 12,
	2, 2, 489, 486, 3, 2, 2, 2, 489, 487, 3, 2, 2, 2, 489, 488, 3, 2, 2, 2,
	490, 493, 3, 2, 2, 2, 491, 489, 3, 2, 2, 2, 491, 492, 3, 2, 2, 2, 492,
	495, 3, 2, 2, 2, 493, 491, 3, 2, 2, 2, 494, 472, 3, 2, 2, 2, 494, 485,
	3, 2, 2, 2, 495, 138, 3, 2, 2, 2, 496, 497, 9, 13, 2, 2, 497, 140, 3, 2,
	2, 2, 498, 502, 7, 39, 2, 2, 499, 501, 11, 2, 2, 2, 500, 499, 3, 2, 2,
	2, 501, 504, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 502, 500, 3, 2, 2, 2, 503,
	505, 3, 2, 2, 2, 504, 502, 3, 2, 2, 2, 505, 506, 7, 39, 2, 2, 506, 142,
	3, 2, 2, 2, 507, 511, 7, 37, 2, 2, 508, 510, 11, 2, 2, 2, 509, 508, 3,
	2, 2, 2, 510, 513, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2, 511, 509, 3, 2, 2,
	2, 512, 515, 3, 2, 2, 2, 513, 511, 3, 2, 2, 2, 514, 516, 7, 15, 2, 2, 515,
	514, 3, 2, 2, 2, 515, 516, 3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517, 518,
	7, 12, 2, 2, 518, 519, 3, 2, 2, 2, 519, 520, 8, 72, 2, 2, 520, 144, 3,
	2, 2, 2, 521, 523, 7, 15, 2, 2, 522, 521, 3, 2, 2, 2, 522, 523, 3, 2, 2,
	2, 523, 524, 3, 2, 2, 2, 524, 525, 7, 12, 2, 2, 525, 146, 3, 2, 2, 2, 526,
	528, 9, 14, 2, 2, 527, 526, 3, 2, 2, 2, 528, 529, 3, 2, 2, 2, 529, 527,
	3, 2, 2, 2, 529, 530, 3, 2, 2, 2, 530, 531, 3, 2, 2, 2, 531, 532, 8, 74,
	3, 2, 532, 148, 3, 2, 2, 2, 42, 2, 328, 331, 336, 339, 346, 352, 356, 359,
	364, 367, 370, 376, 379, 382, 384, 390, 400, 405, 407, 414, 416, 423, 425,
	429, 436, 454, 465, 470, 475, 480, 482, 489, 491, 494, 502, 511, 515, 522,
	529, 4, 9, 65, 2, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'[['", "']'", "'['", "'::'", "':::'", "'$'", "'@'", "'^'",
	"'-'", "'+'", "':'", "'*'", "'/'", "'>'", "'>='", "'<'", "'<='", "'=='",
	"'!='", "'!'", "'&'", "'&&'", "'|'", "'||'", "'~'", "'<-'", "'<<-'", "'='",
	"'->'", "'->>'", "':='", "'function'", "'('", "')'", "'{'", "'}'", "'if'",
	"'else'", "'for'", "'in'", "'while'", "'repeat'", "'?'", "'next'", "'break'",
	"'NULL'", "'NA'", "'Inf'", "'NaN'", "'TRUE'", "'FALSE'", "','", "'...'",
	"'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "HEX", "INT", "FLOAT", "COMPLEX", "STRING", "ID", "USER_OP", "NL",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
	"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
	"T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "HEX", "INT", "HEXDIGIT",
	"FLOAT", "DIGIT", "EXP", "COMPLEX", "STRING", "ESC", "UNICODE_ESCAPE",
	"OCTAL_ESCAPE", "HEX_ESCAPE", "ID", "LETTER", "USER_OP", "COMMENT", "NL",
	"WS",
}

type RLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewRLexer(input antlr.CharStream) *RLexer {

	l := new(RLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "R.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RLexer tokens.
const (
	RLexerT__0    = 1
	RLexerT__1    = 2
	RLexerT__2    = 3
	RLexerT__3    = 4
	RLexerT__4    = 5
	RLexerT__5    = 6
	RLexerT__6    = 7
	RLexerT__7    = 8
	RLexerT__8    = 9
	RLexerT__9    = 10
	RLexerT__10   = 11
	RLexerT__11   = 12
	RLexerT__12   = 13
	RLexerT__13   = 14
	RLexerT__14   = 15
	RLexerT__15   = 16
	RLexerT__16   = 17
	RLexerT__17   = 18
	RLexerT__18   = 19
	RLexerT__19   = 20
	RLexerT__20   = 21
	RLexerT__21   = 22
	RLexerT__22   = 23
	RLexerT__23   = 24
	RLexerT__24   = 25
	RLexerT__25   = 26
	RLexerT__26   = 27
	RLexerT__27   = 28
	RLexerT__28   = 29
	RLexerT__29   = 30
	RLexerT__30   = 31
	RLexerT__31   = 32
	RLexerT__32   = 33
	RLexerT__33   = 34
	RLexerT__34   = 35
	RLexerT__35   = 36
	RLexerT__36   = 37
	RLexerT__37   = 38
	RLexerT__38   = 39
	RLexerT__39   = 40
	RLexerT__40   = 41
	RLexerT__41   = 42
	RLexerT__42   = 43
	RLexerT__43   = 44
	RLexerT__44   = 45
	RLexerT__45   = 46
	RLexerT__46   = 47
	RLexerT__47   = 48
	RLexerT__48   = 49
	RLexerT__49   = 50
	RLexerT__50   = 51
	RLexerT__51   = 52
	RLexerT__52   = 53
	RLexerT__53   = 54
	RLexerT__54   = 55
	RLexerHEX     = 56
	RLexerINT     = 57
	RLexerFLOAT   = 58
	RLexerCOMPLEX = 59
	RLexerSTRING  = 60
	RLexerID      = 61
	RLexerUSER_OP = 62
	RLexerNL      = 63
	RLexerWS      = 64
)
