// Code generated from mdx.g4 by ANTLR 4.9.3. DO NOT EDIT.

package mdx

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 57, 416,
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
	4, 55, 9, 55, 4, 56, 9, 56, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 6, 50, 331,
	10, 50, 13, 50, 14, 50, 332, 3, 51, 6, 51, 336, 10, 51, 13, 51, 14, 51,
	337, 3, 51, 3, 51, 7, 51, 342, 10, 51, 12, 51, 14, 51, 345, 11, 51, 3,
	52, 3, 52, 7, 52, 349, 10, 52, 12, 52, 14, 52, 352, 11, 52, 3, 53, 3, 53,
	3, 53, 3, 53, 3, 53, 6, 53, 359, 10, 53, 13, 53, 14, 53, 360, 3, 53, 7,
	53, 364, 10, 53, 12, 53, 14, 53, 367, 11, 53, 3, 53, 5, 53, 370, 10, 53,
	3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 6, 54, 377, 10, 54, 13, 54, 14, 54,
	378, 3, 54, 7, 54, 382, 10, 54, 12, 54, 14, 54, 385, 11, 54, 3, 54, 5,
	54, 388, 10, 54, 3, 54, 3, 54, 3, 55, 3, 55, 7, 55, 394, 10, 55, 12, 55,
	14, 55, 397, 11, 55, 3, 55, 3, 55, 3, 55, 7, 55, 402, 10, 55, 12, 55, 14,
	55, 405, 11, 55, 3, 55, 5, 55, 408, 10, 55, 3, 56, 6, 56, 411, 10, 56,
	13, 56, 14, 56, 412, 3, 56, 3, 56, 2, 2, 57, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16,
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25,
	49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34,
	67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43,
	85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52,
	103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 3, 2, 8, 6, 2, 38, 38, 67,
	92, 97, 97, 99, 124, 7, 2, 38, 38, 50, 59, 67, 92, 97, 97, 99, 124, 4,
	2, 11, 11, 34, 34, 3, 2, 36, 36, 3, 2, 41, 41, 5, 2, 11, 12, 14, 15, 34,
	34, 2, 429, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2,
	63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2,
	2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2,
	2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2,
	2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3,
	2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101,
	3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2,
	2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 3, 113, 3, 2, 2, 2, 5, 115, 3,
	2, 2, 2, 7, 117, 3, 2, 2, 2, 9, 119, 3, 2, 2, 2, 11, 121, 3, 2, 2, 2, 13,
	123, 3, 2, 2, 2, 15, 126, 3, 2, 2, 2, 17, 128, 3, 2, 2, 2, 19, 130, 3,
	2, 2, 2, 21, 133, 3, 2, 2, 2, 23, 135, 3, 2, 2, 2, 25, 137, 3, 2, 2, 2,
	27, 140, 3, 2, 2, 2, 29, 142, 3, 2, 2, 2, 31, 144, 3, 2, 2, 2, 33, 146,
	3, 2, 2, 2, 35, 149, 3, 2, 2, 2, 37, 151, 3, 2, 2, 2, 39, 153, 3, 2, 2,
	2, 41, 155, 3, 2, 2, 2, 43, 157, 3, 2, 2, 2, 45, 161, 3, 2, 2, 2, 47, 164,
	3, 2, 2, 2, 49, 169, 3, 2, 2, 2, 51, 174, 3, 2, 2, 2, 53, 187, 3, 2, 2,
	2, 55, 194, 3, 2, 2, 2, 57, 204, 3, 2, 2, 2, 59, 209, 3, 2, 2, 2, 61, 215,
	3, 2, 2, 2, 63, 219, 3, 2, 2, 2, 65, 235, 3, 2, 2, 2, 67, 240, 3, 2, 2,
	2, 69, 247, 3, 2, 2, 2, 71, 254, 3, 2, 2, 2, 73, 258, 3, 2, 2, 2, 75, 262,
	3, 2, 2, 2, 77, 265, 3, 2, 2, 2, 79, 268, 3, 2, 2, 2, 81, 279, 3, 2, 2,
	2, 83, 286, 3, 2, 2, 2, 85, 294, 3, 2, 2, 2, 87, 298, 3, 2, 2, 2, 89, 303,
	3, 2, 2, 2, 91, 309, 3, 2, 2, 2, 93, 314, 3, 2, 2, 2, 95, 320, 3, 2, 2,
	2, 97, 324, 3, 2, 2, 2, 99, 330, 3, 2, 2, 2, 101, 335, 3, 2, 2, 2, 103,
	346, 3, 2, 2, 2, 105, 353, 3, 2, 2, 2, 107, 373, 3, 2, 2, 2, 109, 407,
	3, 2, 2, 2, 111, 410, 3, 2, 2, 2, 113, 114, 7, 41, 2, 2, 114, 4, 3, 2,
	2, 2, 115, 116, 7, 44, 2, 2, 116, 6, 3, 2, 2, 2, 117, 118, 7, 60, 2, 2,
	118, 8, 3, 2, 2, 2, 119, 120, 7, 61, 2, 2, 120, 10, 3, 2, 2, 2, 121, 122,
	7, 46, 2, 2, 122, 12, 3, 2, 2, 2, 123, 124, 7, 126, 2, 2, 124, 125, 7,
	126, 2, 2, 125, 14, 3, 2, 2, 2, 126, 127, 7, 48, 2, 2, 127, 16, 3, 2, 2,
	2, 128, 129, 7, 63, 2, 2, 129, 18, 3, 2, 2, 2, 130, 131, 7, 64, 2, 2, 131,
	132, 7, 63, 2, 2, 132, 20, 3, 2, 2, 2, 133, 134, 7, 64, 2, 2, 134, 22,
	3, 2, 2, 2, 135, 136, 7, 125, 2, 2, 136, 24, 3, 2, 2, 2, 137, 138, 7, 62,
	2, 2, 138, 139, 7, 63, 2, 2, 139, 26, 3, 2, 2, 2, 140, 141, 7, 42, 2, 2,
	141, 28, 3, 2, 2, 2, 142, 143, 7, 62, 2, 2, 143, 30, 3, 2, 2, 2, 144, 145,
	7, 47, 2, 2, 145, 32, 3, 2, 2, 2, 146, 147, 7, 62, 2, 2, 147, 148, 7, 64,
	2, 2, 148, 34, 3, 2, 2, 2, 149, 150, 7, 45, 2, 2, 150, 36, 3, 2, 2, 2,
	151, 152, 7, 127, 2, 2, 152, 38, 3, 2, 2, 2, 153, 154, 7, 43, 2, 2, 154,
	40, 3, 2, 2, 2, 155, 156, 7, 49, 2, 2, 156, 42, 3, 2, 2, 2, 157, 158, 7,
	67, 2, 2, 158, 159, 7, 80, 2, 2, 159, 160, 7, 70, 2, 2, 160, 44, 3, 2,
	2, 2, 161, 162, 7, 67, 2, 2, 162, 163, 7, 85, 2, 2, 163, 46, 3, 2, 2, 2,
	164, 165, 7, 69, 2, 2, 165, 166, 7, 67, 2, 2, 166, 167, 7, 85, 2, 2, 167,
	168, 7, 71, 2, 2, 168, 48, 3, 2, 2, 2, 169, 170, 7, 69, 2, 2, 170, 171,
	7, 71, 2, 2, 171, 172, 7, 78, 2, 2, 172, 173, 7, 78, 2, 2, 173, 50, 3,
	2, 2, 2, 174, 175, 7, 69, 2, 2, 175, 176, 7, 71, 2, 2, 176, 177, 7, 78,
	2, 2, 177, 178, 7, 78, 2, 2, 178, 179, 7, 97, 2, 2, 179, 180, 7, 81, 2,
	2, 180, 181, 7, 84, 2, 2, 181, 182, 7, 70, 2, 2, 182, 183, 7, 75, 2, 2,
	183, 184, 7, 80, 2, 2, 184, 185, 7, 67, 2, 2, 185, 186, 7, 78, 2, 2, 186,
	52, 3, 2, 2, 2, 187, 188, 7, 69, 2, 2, 188, 189, 7, 84, 2, 2, 189, 190,
	7, 71, 2, 2, 190, 191, 7, 67, 2, 2, 191, 192, 7, 86, 2, 2, 192, 193, 7,
	71, 2, 2, 193, 54, 3, 2, 2, 2, 194, 195, 7, 70, 2, 2, 195, 196, 7, 75,
	2, 2, 196, 197, 7, 79, 2, 2, 197, 198, 7, 71, 2, 2, 198, 199, 7, 80, 2,
	2, 199, 200, 7, 85, 2, 2, 200, 201, 7, 75, 2, 2, 201, 202, 7, 81, 2, 2,
	202, 203, 7, 80, 2, 2, 203, 56, 3, 2, 2, 2, 204, 205, 7, 71, 2, 2, 205,
	206, 7, 78, 2, 2, 206, 207, 7, 85, 2, 2, 207, 208, 7, 71, 2, 2, 208, 58,
	3, 2, 2, 2, 209, 210, 7, 71, 2, 2, 210, 211, 7, 79, 2, 2, 211, 212, 7,
	82, 2, 2, 212, 213, 7, 86, 2, 2, 213, 214, 7, 91, 2, 2, 214, 60, 3, 2,
	2, 2, 215, 216, 7, 71, 2, 2, 216, 217, 7, 80, 2, 2, 217, 218, 7, 70, 2,
	2, 218, 62, 3, 2, 2, 2, 219, 220, 7, 72, 2, 2, 220, 221, 7, 81, 2, 2, 221,
	222, 7, 84, 2, 2, 222, 223, 7, 79, 2, 2, 223, 224, 7, 67, 2, 2, 224, 225,
	7, 86, 2, 2, 225, 226, 7, 86, 2, 2, 226, 227, 7, 71, 2, 2, 227, 228, 7,
	70, 2, 2, 228, 229, 7, 97, 2, 2, 229, 230, 7, 88, 2, 2, 230, 231, 7, 67,
	2, 2, 231, 232, 7, 78, 2, 2, 232, 233, 7, 87, 2, 2, 233, 234, 7, 71, 2,
	2, 234, 64, 3, 2, 2, 2, 235, 236, 7, 72, 2, 2, 236, 237, 7, 84, 2, 2, 237,
	238, 7, 81, 2, 2, 238, 239, 7, 79, 2, 2, 239, 66, 3, 2, 2, 2, 240, 241,
	7, 73, 2, 2, 241, 242, 7, 78, 2, 2, 242, 243, 7, 81, 2, 2, 243, 244, 7,
	68, 2, 2, 244, 245, 7, 67, 2, 2, 245, 246, 7, 78, 2, 2, 246, 68, 3, 2,
	2, 2, 247, 248, 7, 79, 2, 2, 248, 249, 7, 71, 2, 2, 249, 250, 7, 79, 2,
	2, 250, 251, 7, 68, 2, 2, 251, 252, 7, 71, 2, 2, 252, 253, 7, 84, 2, 2,
	253, 70, 3, 2, 2, 2, 254, 255, 7, 80, 2, 2, 255, 256, 7, 81, 2, 2, 256,
	257, 7, 80, 2, 2, 257, 72, 3, 2, 2, 2, 258, 259, 7, 80, 2, 2, 259, 260,
	7, 81, 2, 2, 260, 261, 7, 86, 2, 2, 261, 74, 3, 2, 2, 2, 262, 263, 7, 81,
	2, 2, 263, 264, 7, 80, 2, 2, 264, 76, 3, 2, 2, 2, 265, 266, 7, 81, 2, 2,
	266, 267, 7, 84, 2, 2, 267, 78, 3, 2, 2, 2, 268, 269, 7, 82, 2, 2, 269,
	270, 7, 84, 2, 2, 270, 271, 7, 81, 2, 2, 271, 272, 7, 82, 2, 2, 272, 273,
	7, 71, 2, 2, 273, 274, 7, 84, 2, 2, 274, 275, 7, 86, 2, 2, 275, 276, 7,
	75, 2, 2, 276, 277, 7, 71, 2, 2, 277, 278, 7, 85, 2, 2, 278, 80, 3, 2,
	2, 2, 279, 280, 7, 85, 2, 2, 280, 281, 7, 71, 2, 2, 281, 282, 7, 78, 2,
	2, 282, 283, 7, 71, 2, 2, 283, 284, 7, 69, 2, 2, 284, 285, 7, 86, 2, 2,
	285, 82, 3, 2, 2, 2, 286, 287, 7, 85, 2, 2, 287, 288, 7, 71, 2, 2, 288,
	289, 7, 85, 2, 2, 289, 290, 7, 85, 2, 2, 290, 291, 7, 75, 2, 2, 291, 292,
	7, 81, 2, 2, 292, 293, 7, 80, 2, 2, 293, 84, 3, 2, 2, 2, 294, 295, 7, 85,
	2, 2, 295, 296, 7, 71, 2, 2, 296, 297, 7, 86, 2, 2, 297, 86, 3, 2, 2, 2,
	298, 299, 7, 86, 2, 2, 299, 300, 7, 74, 2, 2, 300, 301, 7, 71, 2, 2, 301,
	302, 7, 80, 2, 2, 302, 88, 3, 2, 2, 2, 303, 304, 7, 88, 2, 2, 304, 305,
	7, 67, 2, 2, 305, 306, 7, 78, 2, 2, 306, 307, 7, 87, 2, 2, 307, 308, 7,
	71, 2, 2, 308, 90, 3, 2, 2, 2, 309, 310, 7, 89, 2, 2, 310, 311, 7, 74,
	2, 2, 311, 312, 7, 71, 2, 2, 312, 313, 7, 80, 2, 2, 313, 92, 3, 2, 2, 2,
	314, 315, 7, 89, 2, 2, 315, 316, 7, 74, 2, 2, 316, 317, 7, 71, 2, 2, 317,
	318, 7, 84, 2, 2, 318, 319, 7, 71, 2, 2, 319, 94, 3, 2, 2, 2, 320, 321,
	7, 90, 2, 2, 321, 322, 7, 81, 2, 2, 322, 323, 7, 84, 2, 2, 323, 96, 3,
	2, 2, 2, 324, 325, 7, 89, 2, 2, 325, 326, 7, 75, 2, 2, 326, 327, 7, 86,
	2, 2, 327, 328, 7, 74, 2, 2, 328, 98, 3, 2, 2, 2, 329, 331, 4, 50, 59,
	2, 330, 329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 332,
	333, 3, 2, 2, 2, 333, 100, 3, 2, 2, 2, 334, 336, 4, 50, 59, 2, 335, 334,
	3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 335, 3, 2, 2, 2, 337, 338, 3, 2,
	2, 2, 338, 339, 3, 2, 2, 2, 339, 343, 7, 48, 2, 2, 340, 342, 4, 50, 59,
	2, 341, 340, 3, 2, 2, 2, 342, 345, 3, 2, 2, 2, 343, 341, 3, 2, 2, 2, 343,
	344, 3, 2, 2, 2, 344, 102, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 346, 350,
	9, 2, 2, 2, 347, 349, 9, 3, 2, 2, 348, 347, 3, 2, 2, 2, 349, 352, 3, 2,
	2, 2, 350, 348, 3, 2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 104, 3, 2, 2, 2,
	352, 350, 3, 2, 2, 2, 353, 354, 7, 93, 2, 2, 354, 355, 7, 40, 2, 2, 355,
	369, 3, 2, 2, 2, 356, 365, 5, 103, 52, 2, 357, 359, 9, 4, 2, 2, 358, 357,
	3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360, 361, 3, 2,
	2, 2, 361, 362, 3, 2, 2, 2, 362, 364, 5, 103, 52, 2, 363, 358, 3, 2, 2,
	2, 364, 367, 3, 2, 2, 2, 365, 363, 3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366,
	370, 3, 2, 2, 2, 367, 365, 3, 2, 2, 2, 368, 370, 5, 99, 50, 2, 369, 356,
	3, 2, 2, 2, 369, 368, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 372, 7, 95,
	2, 2, 372, 106, 3, 2, 2, 2, 373, 387, 7, 93, 2, 2, 374, 383, 5, 103, 52,
	2, 375, 377, 9, 4, 2, 2, 376, 375, 3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378,
	376, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2, 380, 382,
	5, 103, 52, 2, 381, 376, 3, 2, 2, 2, 382, 385, 3, 2, 2, 2, 383, 381, 3,
	2, 2, 2, 383, 384, 3, 2, 2, 2, 384, 388, 3, 2, 2, 2, 385, 383, 3, 2, 2,
	2, 386, 388, 5, 99, 50, 2, 387, 374, 3, 2, 2, 2, 387, 386, 3, 2, 2, 2,
	388, 389, 3, 2, 2, 2, 389, 390, 7, 95, 2, 2, 390, 108, 3, 2, 2, 2, 391,
	395, 7, 36, 2, 2, 392, 394, 10, 5, 2, 2, 393, 392, 3, 2, 2, 2, 394, 397,
	3, 2, 2, 2, 395, 393, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 398, 3, 2,
	2, 2, 397, 395, 3, 2, 2, 2, 398, 408, 7, 36, 2, 2, 399, 403, 7, 41, 2,
	2, 400, 402, 10, 6, 2, 2, 401, 400, 3, 2, 2, 2, 402, 405, 3, 2, 2, 2, 403,
	401, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 406, 3, 2, 2, 2, 405, 403,
	3, 2, 2, 2, 406, 408, 7, 41, 2, 2, 407, 391, 3, 2, 2, 2, 407, 399, 3, 2,
	2, 2, 408, 110, 3, 2, 2, 2, 409, 411, 9, 7, 2, 2, 410, 409, 3, 2, 2, 2,
	411, 412, 3, 2, 2, 2, 412, 410, 3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413,
	414, 3, 2, 2, 2, 414, 415, 8, 56, 2, 2, 415, 112, 3, 2, 2, 2, 17, 2, 332,
	337, 343, 350, 360, 365, 369, 378, 383, 387, 395, 403, 407, 412, 3, 8,
	2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'''", "'*'", "':'", "';'", "','", "'||'", "'.'", "'='", "'>='", "'>'",
	"'{'", "'<='", "'('", "'<'", "'-'", "'<>'", "'+'", "'}'", "')'", "'/'",
	"'AND'", "'AS'", "'CASE'", "'CELL'", "'CELL_ORDINAL'", "'CREATE'", "'DIMENSION'",
	"'ELSE'", "'EMPTY'", "'END'", "'FORMATTED_VALUE'", "'FROM'", "'GLOBAL'",
	"'MEMBER'", "'NON'", "'NOT'", "'ON'", "'OR'", "'PROPERTIES'", "'SELECT'",
	"'SESSION'", "'SET'", "'THEN'", "'VALUE'", "'WHEN'", "'WHERE'", "'XOR'",
	"'WITH'",
}

var lexerSymbolicNames = []string{
	"", "QUOTE", "ASTERISK", "COLON", "SEMICOLON", "COMMA", "CONCAT", "DOT",
	"EQ", "GE", "GT", "LBRACE", "LE", "LPAREN", "LT", "MINUS", "NE", "PLUS",
	"RBRACE", "RPAREN", "SOLIDUS", "AND", "AS", "CASE", "CELL", "CELL_ORDINAL",
	"CREATE", "DIMENSION", "ELSE", "EMPTY", "END", "FORMATTED_VALUE", "FROM",
	"GLOBAL", "MEMBER", "NON", "NOT", "ON", "OR", "PROPERTIES", "SELECT", "SESSION",
	"SET", "THEN", "VALUE", "WHEN", "WHERE", "XOR", "WITH", "NUMBER", "F",
	"ID", "AMP_QUOTED_ID", "QUOTED_ID", "STRING", "WS",
}

var lexerRuleNames = []string{
	"QUOTE", "ASTERISK", "COLON", "SEMICOLON", "COMMA", "CONCAT", "DOT", "EQ",
	"GE", "GT", "LBRACE", "LE", "LPAREN", "LT", "MINUS", "NE", "PLUS", "RBRACE",
	"RPAREN", "SOLIDUS", "AND", "AS", "CASE", "CELL", "CELL_ORDINAL", "CREATE",
	"DIMENSION", "ELSE", "EMPTY", "END", "FORMATTED_VALUE", "FROM", "GLOBAL",
	"MEMBER", "NON", "NOT", "ON", "OR", "PROPERTIES", "SELECT", "SESSION",
	"SET", "THEN", "VALUE", "WHEN", "WHERE", "XOR", "WITH", "NUMBER", "F",
	"ID", "AMP_QUOTED_ID", "QUOTED_ID", "STRING", "WS",
}

type mdxLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewmdxLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *mdxLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewmdxLexer(input antlr.CharStream) *mdxLexer {
	l := new(mdxLexer)
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
	l.GrammarFileName = "mdx.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// mdxLexer tokens.
const (
	mdxLexerQUOTE           = 1
	mdxLexerASTERISK        = 2
	mdxLexerCOLON           = 3
	mdxLexerSEMICOLON       = 4
	mdxLexerCOMMA           = 5
	mdxLexerCONCAT          = 6
	mdxLexerDOT             = 7
	mdxLexerEQ              = 8
	mdxLexerGE              = 9
	mdxLexerGT              = 10
	mdxLexerLBRACE          = 11
	mdxLexerLE              = 12
	mdxLexerLPAREN          = 13
	mdxLexerLT              = 14
	mdxLexerMINUS           = 15
	mdxLexerNE              = 16
	mdxLexerPLUS            = 17
	mdxLexerRBRACE          = 18
	mdxLexerRPAREN          = 19
	mdxLexerSOLIDUS         = 20
	mdxLexerAND             = 21
	mdxLexerAS              = 22
	mdxLexerCASE            = 23
	mdxLexerCELL            = 24
	mdxLexerCELL_ORDINAL    = 25
	mdxLexerCREATE          = 26
	mdxLexerDIMENSION       = 27
	mdxLexerELSE            = 28
	mdxLexerEMPTY           = 29
	mdxLexerEND             = 30
	mdxLexerFORMATTED_VALUE = 31
	mdxLexerFROM            = 32
	mdxLexerGLOBAL          = 33
	mdxLexerMEMBER          = 34
	mdxLexerNON             = 35
	mdxLexerNOT             = 36
	mdxLexerON              = 37
	mdxLexerOR              = 38
	mdxLexerPROPERTIES      = 39
	mdxLexerSELECT          = 40
	mdxLexerSESSION         = 41
	mdxLexerSET             = 42
	mdxLexerTHEN            = 43
	mdxLexerVALUE           = 44
	mdxLexerWHEN            = 45
	mdxLexerWHERE           = 46
	mdxLexerXOR             = 47
	mdxLexerWITH            = 48
	mdxLexerNUMBER          = 49
	mdxLexerF               = 50
	mdxLexerID              = 51
	mdxLexerAMP_QUOTED_ID   = 52
	mdxLexerQUOTED_ID       = 53
	mdxLexerSTRING          = 54
	mdxLexerWS              = 55
)
