// Code generated from ISL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package isl

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 49, 439,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 39, 3, 39, 3, 39, 3, 40, 6, 40, 330, 10, 40, 13, 40, 14, 40, 331, 3,
	41, 3, 41, 3, 41, 3, 41, 7, 41, 338, 10, 41, 12, 41, 14, 41, 341, 11, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 349, 10, 41, 3, 42, 3,
	42, 7, 42, 353, 10, 42, 12, 42, 14, 42, 356, 11, 42, 3, 42, 5, 42, 359,
	10, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5,
	43, 380, 10, 43, 3, 44, 3, 44, 7, 44, 384, 10, 44, 12, 44, 14, 44, 387,
	11, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 5, 45, 401, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 7, 46, 410, 10, 46, 12, 46, 14, 46, 413, 11, 46, 3, 46,
	5, 46, 416, 10, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 7, 47, 424,
	10, 47, 12, 47, 14, 47, 427, 11, 47, 3, 47, 5, 47, 430, 10, 47, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 2, 2, 49, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14,
	27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23,
	45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32,
	63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41,
	81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 3, 2, 9,
	11, 2, 35, 35, 38, 40, 44, 45, 47, 60, 62, 92, 94, 94, 96, 97, 99, 124,
	128, 128, 3, 2, 50, 59, 3, 2, 51, 59, 3, 2, 34, 128, 5, 2, 50, 59, 67,
	92, 99, 124, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 14, 15, 34, 34, 2, 455,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41,
	3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2,
	49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2,
	2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2,
	2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2,
	2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3,
	2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87,
	3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2,
	95, 3, 2, 2, 2, 3, 97, 3, 2, 2, 2, 5, 99, 3, 2, 2, 2, 7, 106, 3, 2, 2,
	2, 9, 108, 3, 2, 2, 2, 11, 122, 3, 2, 2, 2, 13, 128, 3, 2, 2, 2, 15, 135,
	3, 2, 2, 2, 17, 140, 3, 2, 2, 2, 19, 146, 3, 2, 2, 2, 21, 153, 3, 2, 2,
	2, 23, 155, 3, 2, 2, 2, 25, 161, 3, 2, 2, 2, 27, 163, 3, 2, 2, 2, 29, 165,
	3, 2, 2, 2, 31, 172, 3, 2, 2, 2, 33, 179, 3, 2, 2, 2, 35, 183, 3, 2, 2,
	2, 37, 188, 3, 2, 2, 2, 39, 194, 3, 2, 2, 2, 41, 199, 3, 2, 2, 2, 43, 205,
	3, 2, 2, 2, 45, 208, 3, 2, 2, 2, 47, 212, 3, 2, 2, 2, 49, 215, 3, 2, 2,
	2, 51, 220, 3, 2, 2, 2, 53, 222, 3, 2, 2, 2, 55, 224, 3, 2, 2, 2, 57, 226,
	3, 2, 2, 2, 59, 229, 3, 2, 2, 2, 61, 231, 3, 2, 2, 2, 63, 234, 3, 2, 2,
	2, 65, 247, 3, 2, 2, 2, 67, 260, 3, 2, 2, 2, 69, 273, 3, 2, 2, 2, 71, 289,
	3, 2, 2, 2, 73, 305, 3, 2, 2, 2, 75, 317, 3, 2, 2, 2, 77, 325, 3, 2, 2,
	2, 79, 329, 3, 2, 2, 2, 81, 348, 3, 2, 2, 2, 83, 358, 3, 2, 2, 2, 85, 379,
	3, 2, 2, 2, 87, 381, 3, 2, 2, 2, 89, 400, 3, 2, 2, 2, 91, 402, 3, 2, 2,
	2, 93, 421, 3, 2, 2, 2, 95, 435, 3, 2, 2, 2, 97, 98, 7, 42, 2, 2, 98, 4,
	3, 2, 2, 2, 99, 100, 7, 102, 2, 2, 100, 101, 7, 103, 2, 2, 101, 102, 7,
	104, 2, 2, 102, 103, 7, 107, 2, 2, 103, 104, 7, 112, 2, 2, 104, 105, 7,
	103, 2, 2, 105, 6, 3, 2, 2, 2, 106, 107, 7, 43, 2, 2, 107, 8, 3, 2, 2,
	2, 108, 109, 7, 102, 2, 2, 109, 110, 7, 103, 2, 2, 110, 111, 7, 104, 2,
	2, 111, 112, 7, 107, 2, 2, 112, 113, 7, 112, 2, 2, 113, 114, 7, 103, 2,
	2, 114, 115, 7, 47, 2, 2, 115, 116, 7, 117, 2, 2, 116, 117, 7, 118, 2,
	2, 117, 118, 7, 116, 2, 2, 118, 119, 7, 119, 2, 2, 119, 120, 7, 101, 2,
	2, 120, 121, 7, 118, 2, 2, 121, 10, 3, 2, 2, 2, 122, 123, 7, 100, 2, 2,
	123, 124, 7, 103, 2, 2, 124, 125, 7, 105, 2, 2, 125, 126, 7, 107, 2, 2,
	126, 127, 7, 112, 2, 2, 127, 12, 3, 2, 2, 2, 128, 129, 7, 100, 2, 2, 129,
	130, 7, 103, 2, 2, 130, 131, 7, 105, 2, 2, 131, 132, 7, 107, 2, 2, 132,
	133, 7, 112, 2, 2, 133, 134, 7, 50, 2, 2, 134, 14, 3, 2, 2, 2, 135, 136,
	7, 117, 2, 2, 136, 137, 7, 103, 2, 2, 137, 138, 7, 118, 2, 2, 138, 139,
	7, 35, 2, 2, 139, 16, 3, 2, 2, 2, 140, 141, 7, 102, 2, 2, 141, 142, 7,
	103, 2, 2, 142, 143, 7, 110, 2, 2, 143, 144, 7, 99, 2, 2, 144, 145, 7,
	123, 2, 2, 145, 18, 3, 2, 2, 2, 146, 147, 7, 110, 2, 2, 147, 148, 7, 99,
	2, 2, 148, 149, 7, 111, 2, 2, 149, 150, 7, 100, 2, 2, 150, 151, 7, 102,
	2, 2, 151, 152, 7, 99, 2, 2, 152, 20, 3, 2, 2, 2, 153, 154, 7, 957, 2,
	2, 154, 22, 3, 2, 2, 2, 155, 156, 7, 110, 2, 2, 156, 157, 7, 113, 2, 2,
	157, 158, 7, 101, 2, 2, 158, 159, 7, 99, 2, 2, 159, 160, 7, 110, 2, 2,
	160, 24, 3, 2, 2, 2, 161, 162, 7, 93, 2, 2, 162, 26, 3, 2, 2, 2, 163, 164,
	7, 95, 2, 2, 164, 28, 3, 2, 2, 2, 165, 166, 7, 110, 2, 2, 166, 167, 7,
	103, 2, 2, 167, 168, 7, 118, 2, 2, 168, 169, 7, 116, 2, 2, 169, 170, 7,
	103, 2, 2, 170, 171, 7, 101, 2, 2, 171, 30, 3, 2, 2, 2, 172, 173, 7, 117,
	2, 2, 173, 174, 7, 106, 2, 2, 174, 175, 7, 99, 2, 2, 175, 176, 7, 116,
	2, 2, 176, 177, 7, 103, 2, 2, 177, 178, 7, 102, 2, 2, 178, 32, 3, 2, 2,
	2, 179, 180, 7, 110, 2, 2, 180, 181, 7, 103, 2, 2, 181, 182, 7, 118, 2,
	2, 182, 34, 3, 2, 2, 2, 183, 184, 7, 110, 2, 2, 184, 185, 7, 103, 2, 2,
	185, 186, 7, 118, 2, 2, 186, 187, 7, 44, 2, 2, 187, 36, 3, 2, 2, 2, 188,
	189, 7, 116, 2, 2, 189, 190, 7, 103, 2, 2, 190, 191, 7, 101, 2, 2, 191,
	192, 7, 119, 2, 2, 192, 193, 7, 116, 2, 2, 193, 38, 3, 2, 2, 2, 194, 195,
	7, 101, 2, 2, 195, 196, 7, 113, 2, 2, 196, 197, 7, 112, 2, 2, 197, 198,
	7, 102, 2, 2, 198, 40, 3, 2, 2, 2, 199, 200, 7, 103, 2, 2, 200, 201, 7,
	110, 2, 2, 201, 202, 7, 117, 2, 2, 202, 203, 7, 103, 2, 2, 203, 204, 7,
	34, 2, 2, 204, 42, 3, 2, 2, 2, 205, 206, 7, 107, 2, 2, 206, 207, 7, 104,
	2, 2, 207, 44, 3, 2, 2, 2, 208, 209, 7, 99, 2, 2, 209, 210, 7, 112, 2,
	2, 210, 211, 7, 102, 2, 2, 211, 46, 3, 2, 2, 2, 212, 213, 7, 113, 2, 2,
	213, 214, 7, 116, 2, 2, 214, 48, 3, 2, 2, 2, 215, 216, 7, 118, 2, 2, 216,
	217, 7, 107, 2, 2, 217, 218, 7, 111, 2, 2, 218, 219, 7, 103, 2, 2, 219,
	50, 3, 2, 2, 2, 220, 221, 7, 8219, 2, 2, 221, 52, 3, 2, 2, 2, 222, 223,
	7, 8218, 2, 2, 223, 54, 3, 2, 2, 2, 224, 225, 7, 41, 2, 2, 225, 56, 3,
	2, 2, 2, 226, 227, 7, 42, 2, 2, 227, 228, 7, 43, 2, 2, 228, 58, 3, 2, 2,
	2, 229, 230, 7, 46, 2, 2, 230, 60, 3, 2, 2, 2, 231, 232, 7, 46, 2, 2, 232,
	233, 7, 66, 2, 2, 233, 62, 3, 2, 2, 2, 234, 235, 7, 101, 2, 2, 235, 236,
	7, 106, 2, 2, 236, 237, 7, 103, 2, 2, 237, 238, 7, 101, 2, 2, 238, 239,
	7, 109, 2, 2, 239, 240, 7, 47, 2, 2, 240, 241, 7, 103, 2, 2, 241, 242,
	7, 122, 2, 2, 242, 243, 7, 114, 2, 2, 243, 244, 7, 103, 2, 2, 244, 245,
	7, 101, 2, 2, 245, 246, 7, 118, 2, 2, 246, 64, 3, 2, 2, 2, 247, 248, 7,
	101, 2, 2, 248, 249, 7, 106, 2, 2, 249, 250, 7, 103, 2, 2, 250, 251, 7,
	101, 2, 2, 251, 252, 7, 109, 2, 2, 252, 253, 7, 47, 2, 2, 253, 254, 7,
	116, 2, 2, 254, 255, 7, 99, 2, 2, 255, 256, 7, 112, 2, 2, 256, 257, 7,
	102, 2, 2, 257, 258, 7, 113, 2, 2, 258, 259, 7, 111, 2, 2, 259, 66, 3,
	2, 2, 2, 260, 261, 7, 101, 2, 2, 261, 262, 7, 106, 2, 2, 262, 263, 7, 103,
	2, 2, 263, 264, 7, 101, 2, 2, 264, 265, 7, 109, 2, 2, 265, 266, 7, 47,
	2, 2, 266, 267, 7, 121, 2, 2, 267, 268, 7, 107, 2, 2, 268, 269, 7, 118,
	2, 2, 269, 270, 7, 106, 2, 2, 270, 271, 7, 107, 2, 2, 271, 272, 7, 112,
	2, 2, 272, 68, 3, 2, 2, 2, 273, 274, 7, 101, 2, 2, 274, 275, 7, 106, 2,
	2, 275, 276, 7, 103, 2, 2, 276, 277, 7, 101, 2, 2, 277, 278, 7, 109, 2,
	2, 278, 279, 7, 47, 2, 2, 279, 280, 7, 111, 2, 2, 280, 281, 7, 103, 2,
	2, 281, 282, 7, 111, 2, 2, 282, 283, 7, 100, 2, 2, 283, 284, 7, 103, 2,
	2, 284, 285, 7, 116, 2, 2, 285, 286, 7, 47, 2, 2, 286, 287, 7, 113, 2,
	2, 287, 288, 7, 104, 2, 2, 288, 70, 3, 2, 2, 2, 289, 290, 7, 101, 2, 2,
	290, 291, 7, 106, 2, 2, 291, 292, 7, 103, 2, 2, 292, 293, 7, 101, 2, 2,
	293, 294, 7, 109, 2, 2, 294, 295, 7, 47, 2, 2, 295, 296, 7, 117, 2, 2,
	296, 297, 7, 99, 2, 2, 297, 298, 7, 118, 2, 2, 298, 299, 7, 107, 2, 2,
	299, 300, 7, 117, 2, 2, 300, 301, 7, 104, 2, 2, 301, 302, 7, 107, 2, 2,
	302, 303, 7, 103, 2, 2, 303, 304, 7, 102, 2, 2, 304, 72, 3, 2, 2, 2, 305,
	306, 7, 101, 2, 2, 306, 307, 7, 106, 2, 2, 307, 308, 7, 103, 2, 2, 308,
	309, 7, 101, 2, 2, 309, 310, 7, 109, 2, 2, 310, 311, 7, 47, 2, 2, 311,
	312, 7, 103, 2, 2, 312, 313, 7, 116, 2, 2, 313, 314, 7, 116, 2, 2, 314,
	315, 7, 113, 2, 2, 315, 316, 7, 116, 2, 2, 316, 74, 3, 2, 2, 2, 317, 318,
	7, 116, 2, 2, 318, 319, 7, 103, 2, 2, 319, 320, 7, 115, 2, 2, 320, 321,
	7, 119, 2, 2, 321, 322, 7, 107, 2, 2, 322, 323, 7, 116, 2, 2, 323, 324,
	7, 103, 2, 2, 324, 76, 3, 2, 2, 2, 325, 326, 7, 8219, 2, 2, 326, 327, 5,
	79, 40, 2, 327, 78, 3, 2, 2, 2, 328, 330, 9, 2, 2, 2, 329, 328, 3, 2, 2,
	2, 330, 331, 3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332,
	80, 3, 2, 2, 2, 333, 349, 5, 83, 42, 2, 334, 335, 5, 83, 42, 2, 335, 339,
	7, 48, 2, 2, 336, 338, 9, 3, 2, 2, 337, 336, 3, 2, 2, 2, 338, 341, 3, 2,
	2, 2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 342, 3, 2, 2, 2,
	341, 339, 3, 2, 2, 2, 342, 343, 9, 4, 2, 2, 343, 349, 3, 2, 2, 2, 344,
	345, 5, 83, 42, 2, 345, 346, 7, 49, 2, 2, 346, 347, 5, 83, 42, 2, 347,
	349, 3, 2, 2, 2, 348, 333, 3, 2, 2, 2, 348, 334, 3, 2, 2, 2, 348, 344,
	3, 2, 2, 2, 349, 82, 3, 2, 2, 2, 350, 354, 9, 4, 2, 2, 351, 353, 9, 3,
	2, 2, 352, 351, 3, 2, 2, 2, 353, 356, 3, 2, 2, 2, 354, 352, 3, 2, 2, 2,
	354, 355, 3, 2, 2, 2, 355, 359, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 357,
	359, 7, 50, 2, 2, 358, 350, 3, 2, 2, 2, 358, 357, 3, 2, 2, 2, 359, 84,
	3, 2, 2, 2, 360, 361, 7, 37, 2, 2, 361, 362, 7, 118, 2, 2, 362, 363, 7,
	116, 2, 2, 363, 364, 7, 119, 2, 2, 364, 380, 7, 103, 2, 2, 365, 366, 7,
	37, 2, 2, 366, 380, 7, 86, 2, 2, 367, 368, 7, 37, 2, 2, 368, 380, 7, 118,
	2, 2, 369, 370, 7, 37, 2, 2, 370, 371, 7, 104, 2, 2, 371, 372, 7, 99, 2,
	2, 372, 373, 7, 110, 2, 2, 373, 374, 7, 117, 2, 2, 374, 380, 7, 103, 2,
	2, 375, 376, 7, 37, 2, 2, 376, 380, 7, 72, 2, 2, 377, 378, 7, 37, 2, 2,
	378, 380, 7, 104, 2, 2, 379, 360, 3, 2, 2, 2, 379, 365, 3, 2, 2, 2, 379,
	367, 3, 2, 2, 2, 379, 369, 3, 2, 2, 2, 379, 375, 3, 2, 2, 2, 379, 377,
	3, 2, 2, 2, 380, 86, 3, 2, 2, 2, 381, 385, 7, 36, 2, 2, 382, 384, 9, 5,
	2, 2, 383, 382, 3, 2, 2, 2, 384, 387, 3, 2, 2, 2, 385, 383, 3, 2, 2, 2,
	385, 386, 3, 2, 2, 2, 386, 388, 3, 2, 2, 2, 387, 385, 3, 2, 2, 2, 388,
	389, 7, 36, 2, 2, 389, 88, 3, 2, 2, 2, 390, 391, 7, 37, 2, 2, 391, 392,
	7, 94, 2, 2, 392, 401, 9, 6, 2, 2, 393, 394, 7, 37, 2, 2, 394, 395, 7,
	94, 2, 2, 395, 396, 7, 117, 2, 2, 396, 397, 7, 114, 2, 2, 397, 398, 7,
	99, 2, 2, 398, 399, 7, 101, 2, 2, 399, 401, 7, 103, 2, 2, 400, 390, 3,
	2, 2, 2, 400, 393, 3, 2, 2, 2, 401, 90, 3, 2, 2, 2, 402, 403, 7, 37, 2,
	2, 403, 404, 7, 110, 2, 2, 404, 405, 7, 99, 2, 2, 405, 406, 7, 112, 2,
	2, 406, 407, 7, 105, 2, 2, 407, 411, 3, 2, 2, 2, 408, 410, 10, 7, 2, 2,
	409, 408, 3, 2, 2, 2, 410, 413, 3, 2, 2, 2, 411, 409, 3, 2, 2, 2, 411,
	412, 3, 2, 2, 2, 412, 415, 3, 2, 2, 2, 413, 411, 3, 2, 2, 2, 414, 416,
	7, 15, 2, 2, 415, 414, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 417, 3, 2,
	2, 2, 417, 418, 7, 12, 2, 2, 418, 419, 3, 2, 2, 2, 419, 420, 8, 46, 2,
	2, 420, 92, 3, 2, 2, 2, 421, 425, 7, 61, 2, 2, 422, 424, 10, 7, 2, 2, 423,
	422, 3, 2, 2, 2, 424, 427, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 425, 426,
	3, 2, 2, 2, 426, 429, 3, 2, 2, 2, 427, 425, 3, 2, 2, 2, 428, 430, 7, 15,
	2, 2, 429, 428, 3, 2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2,
	431, 432, 7, 12, 2, 2, 432, 433, 3, 2, 2, 2, 433, 434, 8, 47, 2, 2, 434,
	94, 3, 2, 2, 2, 435, 436, 9, 8, 2, 2, 436, 437, 3, 2, 2, 2, 437, 438, 8,
	48, 2, 2, 438, 96, 3, 2, 2, 2, 16, 2, 329, 331, 339, 348, 354, 358, 379,
	385, 400, 411, 415, 425, 429, 3, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "'define'", "')'", "'define-struct'", "'begin'", "'begin0'",
	"'set!'", "'delay'", "'lambda'", "'\u03BB'", "'local'", "'['", "']'", "'letrec'",
	"'shared'", "'let'", "'let*'", "'recur'", "'cond'", "'else '", "'if'",
	"'and'", "'or'", "'time'", "'\u2019'", "'\u2018'", "'\u0027'", "'()'",
	"','", "',@'", "'check-expect'", "'check-random'", "'check-within'", "'check-member-of'",
	"'check-satisfied'", "'check-error'", "'require'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "SYMBOL", "NAME", "NUMBER", "INT", "BOOLEAN", "STRING", "CHARACTER",
	"LANG", "COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "SYMBOL", "NAME", "NUMBER", "INT",
	"BOOLEAN", "STRING", "CHARACTER", "LANG", "COMMENT", "WS",
}

type ISLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewISLLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *ISLLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewISLLexer(input antlr.CharStream) *ISLLexer {
	l := new(ISLLexer)
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
	l.GrammarFileName = "ISL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ISLLexer tokens.
const (
	ISLLexerT__0      = 1
	ISLLexerT__1      = 2
	ISLLexerT__2      = 3
	ISLLexerT__3      = 4
	ISLLexerT__4      = 5
	ISLLexerT__5      = 6
	ISLLexerT__6      = 7
	ISLLexerT__7      = 8
	ISLLexerT__8      = 9
	ISLLexerT__9      = 10
	ISLLexerT__10     = 11
	ISLLexerT__11     = 12
	ISLLexerT__12     = 13
	ISLLexerT__13     = 14
	ISLLexerT__14     = 15
	ISLLexerT__15     = 16
	ISLLexerT__16     = 17
	ISLLexerT__17     = 18
	ISLLexerT__18     = 19
	ISLLexerT__19     = 20
	ISLLexerT__20     = 21
	ISLLexerT__21     = 22
	ISLLexerT__22     = 23
	ISLLexerT__23     = 24
	ISLLexerT__24     = 25
	ISLLexerT__25     = 26
	ISLLexerT__26     = 27
	ISLLexerT__27     = 28
	ISLLexerT__28     = 29
	ISLLexerT__29     = 30
	ISLLexerT__30     = 31
	ISLLexerT__31     = 32
	ISLLexerT__32     = 33
	ISLLexerT__33     = 34
	ISLLexerT__34     = 35
	ISLLexerT__35     = 36
	ISLLexerT__36     = 37
	ISLLexerSYMBOL    = 38
	ISLLexerNAME      = 39
	ISLLexerNUMBER    = 40
	ISLLexerINT       = 41
	ISLLexerBOOLEAN   = 42
	ISLLexerSTRING    = 43
	ISLLexerCHARACTER = 44
	ISLLexerLANG      = 45
	ISLLexerCOMMENT   = 46
	ISLLexerWS        = 47
)
