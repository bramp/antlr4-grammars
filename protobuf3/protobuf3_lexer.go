// Code generated from Protobuf3.g4 by ANTLR 4.9.3. DO NOT EDIT.

package protobuf3

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 60, 574,
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
	70, 4, 71, 9, 71, 4, 72, 9, 72, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3,
	39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44,
	3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3,
	49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 3, 52, 7, 52, 417, 10, 52, 12, 52,
	14, 52, 420, 11, 52, 3, 52, 3, 52, 3, 52, 7, 52, 425, 10, 52, 12, 52, 14,
	52, 428, 11, 52, 3, 52, 5, 52, 431, 10, 52, 3, 53, 3, 53, 3, 53, 3, 53,
	5, 53, 437, 10, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3,
	55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 57, 5, 57, 461, 10, 57, 3, 58, 3, 58, 3, 58, 5,
	58, 466, 10, 58, 3, 58, 5, 58, 469, 10, 58, 3, 58, 3, 58, 3, 58, 3, 58,
	3, 58, 3, 58, 5, 58, 477, 10, 58, 5, 58, 479, 10, 58, 3, 58, 3, 58, 3,
	58, 3, 58, 3, 58, 3, 58, 5, 58, 487, 10, 58, 3, 59, 3, 59, 3, 59, 5, 59,
	492, 10, 59, 3, 59, 3, 59, 3, 60, 6, 60, 497, 10, 60, 13, 60, 14, 60, 498,
	3, 61, 3, 61, 3, 61, 5, 61, 504, 10, 61, 3, 62, 3, 62, 7, 62, 508, 10,
	62, 12, 62, 14, 62, 511, 11, 62, 3, 63, 3, 63, 7, 63, 515, 10, 63, 12,
	63, 14, 63, 518, 11, 63, 3, 64, 3, 64, 3, 64, 6, 64, 523, 10, 64, 13, 64,
	14, 64, 524, 3, 65, 3, 65, 3, 65, 7, 65, 530, 10, 65, 12, 65, 14, 65, 533,
	11, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70,
	6, 70, 544, 10, 70, 13, 70, 14, 70, 545, 3, 70, 3, 70, 3, 71, 3, 71, 3,
	71, 3, 71, 7, 71, 554, 10, 71, 12, 71, 14, 71, 557, 11, 71, 3, 71, 3, 71,
	3, 72, 3, 72, 3, 72, 3, 72, 7, 72, 565, 10, 72, 12, 72, 14, 72, 568, 11,
	72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 5, 418, 426, 566, 2, 73, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61,
	32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79,
	41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97,
	50, 99, 51, 101, 52, 103, 53, 105, 2, 107, 2, 109, 2, 111, 2, 113, 54,
	115, 55, 117, 2, 119, 2, 121, 56, 123, 2, 125, 2, 127, 2, 129, 57, 131,
	2, 133, 2, 135, 2, 137, 2, 139, 58, 141, 59, 143, 60, 3, 2, 13, 5, 2, 2,
	2, 12, 12, 94, 94, 4, 2, 90, 90, 122, 122, 11, 2, 36, 36, 41, 41, 94, 94,
	99, 100, 104, 104, 112, 112, 116, 116, 118, 118, 120, 120, 4, 2, 71, 71,
	103, 103, 3, 2, 51, 59, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 3,
	2, 50, 57, 5, 2, 50, 59, 67, 72, 99, 104, 5, 2, 11, 12, 14, 15, 34, 34,
	4, 2, 12, 12, 15, 15, 2, 587, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7,
	3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2,
	2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3,
	2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53,
	3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2,
	61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2,
	2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2,
	2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2,
	2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3,
	2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99,
	3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2,
	2, 115, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 139, 3,
	2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 3, 145, 3, 2, 2, 2, 5,
	152, 3, 2, 2, 2, 7, 159, 3, 2, 2, 2, 9, 164, 3, 2, 2, 2, 11, 171, 3, 2,
	2, 2, 13, 179, 3, 2, 2, 2, 15, 186, 3, 2, 2, 2, 17, 195, 3, 2, 2, 2, 19,
	201, 3, 2, 2, 2, 21, 205, 3, 2, 2, 2, 23, 211, 3, 2, 2, 2, 25, 217, 3,
	2, 2, 2, 27, 224, 3, 2, 2, 2, 29, 231, 3, 2, 2, 2, 31, 238, 3, 2, 2, 2,
	33, 245, 3, 2, 2, 2, 35, 253, 3, 2, 2, 2, 37, 261, 3, 2, 2, 2, 39, 270,
	3, 2, 2, 2, 41, 279, 3, 2, 2, 2, 43, 284, 3, 2, 2, 2, 45, 291, 3, 2, 2,
	2, 47, 298, 3, 2, 2, 2, 49, 304, 3, 2, 2, 2, 51, 310, 3, 2, 2, 2, 53, 319,
	3, 2, 2, 2, 55, 322, 3, 2, 2, 2, 57, 326, 3, 2, 2, 2, 59, 331, 3, 2, 2,
	2, 61, 339, 3, 2, 2, 2, 63, 347, 3, 2, 2, 2, 65, 351, 3, 2, 2, 2, 67, 358,
	3, 2, 2, 2, 69, 366, 3, 2, 2, 2, 71, 375, 3, 2, 2, 2, 73, 384, 3, 2, 2,
	2, 75, 386, 3, 2, 2, 2, 77, 388, 3, 2, 2, 2, 79, 390, 3, 2, 2, 2, 81, 392,
	3, 2, 2, 2, 83, 394, 3, 2, 2, 2, 85, 396, 3, 2, 2, 2, 87, 398, 3, 2, 2,
	2, 89, 400, 3, 2, 2, 2, 91, 402, 3, 2, 2, 2, 93, 404, 3, 2, 2, 2, 95, 406,
	3, 2, 2, 2, 97, 408, 3, 2, 2, 2, 99, 410, 3, 2, 2, 2, 101, 412, 3, 2, 2,
	2, 103, 430, 3, 2, 2, 2, 105, 436, 3, 2, 2, 2, 107, 438, 3, 2, 2, 2, 109,
	443, 3, 2, 2, 2, 111, 448, 3, 2, 2, 2, 113, 460, 3, 2, 2, 2, 115, 486,
	3, 2, 2, 2, 117, 488, 3, 2, 2, 2, 119, 496, 3, 2, 2, 2, 121, 503, 3, 2,
	2, 2, 123, 505, 3, 2, 2, 2, 125, 512, 3, 2, 2, 2, 127, 519, 3, 2, 2, 2,
	129, 526, 3, 2, 2, 2, 131, 534, 3, 2, 2, 2, 133, 536, 3, 2, 2, 2, 135,
	538, 3, 2, 2, 2, 137, 540, 3, 2, 2, 2, 139, 543, 3, 2, 2, 2, 141, 549,
	3, 2, 2, 2, 143, 560, 3, 2, 2, 2, 145, 146, 7, 117, 2, 2, 146, 147, 7,
	123, 2, 2, 147, 148, 7, 112, 2, 2, 148, 149, 7, 118, 2, 2, 149, 150, 7,
	99, 2, 2, 150, 151, 7, 122, 2, 2, 151, 4, 3, 2, 2, 2, 152, 153, 7, 107,
	2, 2, 153, 154, 7, 111, 2, 2, 154, 155, 7, 114, 2, 2, 155, 156, 7, 113,
	2, 2, 156, 157, 7, 116, 2, 2, 157, 158, 7, 118, 2, 2, 158, 6, 3, 2, 2,
	2, 159, 160, 7, 121, 2, 2, 160, 161, 7, 103, 2, 2, 161, 162, 7, 99, 2,
	2, 162, 163, 7, 109, 2, 2, 163, 8, 3, 2, 2, 2, 164, 165, 7, 114, 2, 2,
	165, 166, 7, 119, 2, 2, 166, 167, 7, 100, 2, 2, 167, 168, 7, 110, 2, 2,
	168, 169, 7, 107, 2, 2, 169, 170, 7, 101, 2, 2, 170, 10, 3, 2, 2, 2, 171,
	172, 7, 114, 2, 2, 172, 173, 7, 99, 2, 2, 173, 174, 7, 101, 2, 2, 174,
	175, 7, 109, 2, 2, 175, 176, 7, 99, 2, 2, 176, 177, 7, 105, 2, 2, 177,
	178, 7, 103, 2, 2, 178, 12, 3, 2, 2, 2, 179, 180, 7, 113, 2, 2, 180, 181,
	7, 114, 2, 2, 181, 182, 7, 118, 2, 2, 182, 183, 7, 107, 2, 2, 183, 184,
	7, 113, 2, 2, 184, 185, 7, 112, 2, 2, 185, 14, 3, 2, 2, 2, 186, 187, 7,
	116, 2, 2, 187, 188, 7, 103, 2, 2, 188, 189, 7, 114, 2, 2, 189, 190, 7,
	103, 2, 2, 190, 191, 7, 99, 2, 2, 191, 192, 7, 118, 2, 2, 192, 193, 7,
	103, 2, 2, 193, 194, 7, 102, 2, 2, 194, 16, 3, 2, 2, 2, 195, 196, 7, 113,
	2, 2, 196, 197, 7, 112, 2, 2, 197, 198, 7, 103, 2, 2, 198, 199, 7, 113,
	2, 2, 199, 200, 7, 104, 2, 2, 200, 18, 3, 2, 2, 2, 201, 202, 7, 111, 2,
	2, 202, 203, 7, 99, 2, 2, 203, 204, 7, 114, 2, 2, 204, 20, 3, 2, 2, 2,
	205, 206, 7, 107, 2, 2, 206, 207, 7, 112, 2, 2, 207, 208, 7, 118, 2, 2,
	208, 209, 7, 53, 2, 2, 209, 210, 7, 52, 2, 2, 210, 22, 3, 2, 2, 2, 211,
	212, 7, 107, 2, 2, 212, 213, 7, 112, 2, 2, 213, 214, 7, 118, 2, 2, 214,
	215, 7, 56, 2, 2, 215, 216, 7, 54, 2, 2, 216, 24, 3, 2, 2, 2, 217, 218,
	7, 119, 2, 2, 218, 219, 7, 107, 2, 2, 219, 220, 7, 112, 2, 2, 220, 221,
	7, 118, 2, 2, 221, 222, 7, 53, 2, 2, 222, 223, 7, 52, 2, 2, 223, 26, 3,
	2, 2, 2, 224, 225, 7, 119, 2, 2, 225, 226, 7, 107, 2, 2, 226, 227, 7, 112,
	2, 2, 227, 228, 7, 118, 2, 2, 228, 229, 7, 56, 2, 2, 229, 230, 7, 54, 2,
	2, 230, 28, 3, 2, 2, 2, 231, 232, 7, 117, 2, 2, 232, 233, 7, 107, 2, 2,
	233, 234, 7, 112, 2, 2, 234, 235, 7, 118, 2, 2, 235, 236, 7, 53, 2, 2,
	236, 237, 7, 52, 2, 2, 237, 30, 3, 2, 2, 2, 238, 239, 7, 117, 2, 2, 239,
	240, 7, 107, 2, 2, 240, 241, 7, 112, 2, 2, 241, 242, 7, 118, 2, 2, 242,
	243, 7, 56, 2, 2, 243, 244, 7, 54, 2, 2, 244, 32, 3, 2, 2, 2, 245, 246,
	7, 104, 2, 2, 246, 247, 7, 107, 2, 2, 247, 248, 7, 122, 2, 2, 248, 249,
	7, 103, 2, 2, 249, 250, 7, 102, 2, 2, 250, 251, 7, 53, 2, 2, 251, 252,
	7, 52, 2, 2, 252, 34, 3, 2, 2, 2, 253, 254, 7, 104, 2, 2, 254, 255, 7,
	107, 2, 2, 255, 256, 7, 122, 2, 2, 256, 257, 7, 103, 2, 2, 257, 258, 7,
	102, 2, 2, 258, 259, 7, 56, 2, 2, 259, 260, 7, 54, 2, 2, 260, 36, 3, 2,
	2, 2, 261, 262, 7, 117, 2, 2, 262, 263, 7, 104, 2, 2, 263, 264, 7, 107,
	2, 2, 264, 265, 7, 122, 2, 2, 265, 266, 7, 103, 2, 2, 266, 267, 7, 102,
	2, 2, 267, 268, 7, 53, 2, 2, 268, 269, 7, 52, 2, 2, 269, 38, 3, 2, 2, 2,
	270, 271, 7, 117, 2, 2, 271, 272, 7, 104, 2, 2, 272, 273, 7, 107, 2, 2,
	273, 274, 7, 122, 2, 2, 274, 275, 7, 103, 2, 2, 275, 276, 7, 102, 2, 2,
	276, 277, 7, 56, 2, 2, 277, 278, 7, 54, 2, 2, 278, 40, 3, 2, 2, 2, 279,
	280, 7, 100, 2, 2, 280, 281, 7, 113, 2, 2, 281, 282, 7, 113, 2, 2, 282,
	283, 7, 110, 2, 2, 283, 42, 3, 2, 2, 2, 284, 285, 7, 117, 2, 2, 285, 286,
	7, 118, 2, 2, 286, 287, 7, 116, 2, 2, 287, 288, 7, 107, 2, 2, 288, 289,
	7, 112, 2, 2, 289, 290, 7, 105, 2, 2, 290, 44, 3, 2, 2, 2, 291, 292, 7,
	102, 2, 2, 292, 293, 7, 113, 2, 2, 293, 294, 7, 119, 2, 2, 294, 295, 7,
	100, 2, 2, 295, 296, 7, 110, 2, 2, 296, 297, 7, 103, 2, 2, 297, 46, 3,
	2, 2, 2, 298, 299, 7, 104, 2, 2, 299, 300, 7, 110, 2, 2, 300, 301, 7, 113,
	2, 2, 301, 302, 7, 99, 2, 2, 302, 303, 7, 118, 2, 2, 303, 48, 3, 2, 2,
	2, 304, 305, 7, 100, 2, 2, 305, 306, 7, 123, 2, 2, 306, 307, 7, 118, 2,
	2, 307, 308, 7, 103, 2, 2, 308, 309, 7, 117, 2, 2, 309, 50, 3, 2, 2, 2,
	310, 311, 7, 116, 2, 2, 311, 312, 7, 103, 2, 2, 312, 313, 7, 117, 2, 2,
	313, 314, 7, 103, 2, 2, 314, 315, 7, 116, 2, 2, 315, 316, 7, 120, 2, 2,
	316, 317, 7, 103, 2, 2, 317, 318, 7, 102, 2, 2, 318, 52, 3, 2, 2, 2, 319,
	320, 7, 118, 2, 2, 320, 321, 7, 113, 2, 2, 321, 54, 3, 2, 2, 2, 322, 323,
	7, 111, 2, 2, 323, 324, 7, 99, 2, 2, 324, 325, 7, 122, 2, 2, 325, 56, 3,
	2, 2, 2, 326, 327, 7, 103, 2, 2, 327, 328, 7, 112, 2, 2, 328, 329, 7, 119,
	2, 2, 329, 330, 7, 111, 2, 2, 330, 58, 3, 2, 2, 2, 331, 332, 7, 111, 2,
	2, 332, 333, 7, 103, 2, 2, 333, 334, 7, 117, 2, 2, 334, 335, 7, 117, 2,
	2, 335, 336, 7, 99, 2, 2, 336, 337, 7, 105, 2, 2, 337, 338, 7, 103, 2,
	2, 338, 60, 3, 2, 2, 2, 339, 340, 7, 117, 2, 2, 340, 341, 7, 103, 2, 2,
	341, 342, 7, 116, 2, 2, 342, 343, 7, 120, 2, 2, 343, 344, 7, 107, 2, 2,
	344, 345, 7, 101, 2, 2, 345, 346, 7, 103, 2, 2, 346, 62, 3, 2, 2, 2, 347,
	348, 7, 116, 2, 2, 348, 349, 7, 114, 2, 2, 349, 350, 7, 101, 2, 2, 350,
	64, 3, 2, 2, 2, 351, 352, 7, 117, 2, 2, 352, 353, 7, 118, 2, 2, 353, 354,
	7, 116, 2, 2, 354, 355, 7, 103, 2, 2, 355, 356, 7, 99, 2, 2, 356, 357,
	7, 111, 2, 2, 357, 66, 3, 2, 2, 2, 358, 359, 7, 116, 2, 2, 359, 360, 7,
	103, 2, 2, 360, 361, 7, 118, 2, 2, 361, 362, 7, 119, 2, 2, 362, 363, 7,
	116, 2, 2, 363, 364, 7, 112, 2, 2, 364, 365, 7, 117, 2, 2, 365, 68, 3,
	2, 2, 2, 366, 367, 7, 36, 2, 2, 367, 368, 7, 114, 2, 2, 368, 369, 7, 116,
	2, 2, 369, 370, 7, 113, 2, 2, 370, 371, 7, 118, 2, 2, 371, 372, 7, 113,
	2, 2, 372, 373, 7, 53, 2, 2, 373, 374, 7, 36, 2, 2, 374, 70, 3, 2, 2, 2,
	375, 376, 7, 41, 2, 2, 376, 377, 7, 114, 2, 2, 377, 378, 7, 116, 2, 2,
	378, 379, 7, 113, 2, 2, 379, 380, 7, 118, 2, 2, 380, 381, 7, 113, 2, 2,
	381, 382, 7, 53, 2, 2, 382, 383, 7, 41, 2, 2, 383, 72, 3, 2, 2, 2, 384,
	385, 7, 61, 2, 2, 385, 74, 3, 2, 2, 2, 386, 387, 7, 63, 2, 2, 387, 76,
	3, 2, 2, 2, 388, 389, 7, 42, 2, 2, 389, 78, 3, 2, 2, 2, 390, 391, 7, 43,
	2, 2, 391, 80, 3, 2, 2, 2, 392, 393, 7, 93, 2, 2, 393, 82, 3, 2, 2, 2,
	394, 395, 7, 95, 2, 2, 395, 84, 3, 2, 2, 2, 396, 397, 7, 125, 2, 2, 397,
	86, 3, 2, 2, 2, 398, 399, 7, 127, 2, 2, 399, 88, 3, 2, 2, 2, 400, 401,
	7, 62, 2, 2, 401, 90, 3, 2, 2, 2, 402, 403, 7, 64, 2, 2, 403, 92, 3, 2,
	2, 2, 404, 405, 7, 48, 2, 2, 405, 94, 3, 2, 2, 2, 406, 407, 7, 46, 2, 2,
	407, 96, 3, 2, 2, 2, 408, 409, 7, 60, 2, 2, 409, 98, 3, 2, 2, 2, 410, 411,
	7, 45, 2, 2, 411, 100, 3, 2, 2, 2, 412, 413, 7, 47, 2, 2, 413, 102, 3,
	2, 2, 2, 414, 418, 7, 41, 2, 2, 415, 417, 5, 105, 53, 2, 416, 415, 3, 2,
	2, 2, 417, 420, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 418, 416, 3, 2, 2, 2,
	419, 421, 3, 2, 2, 2, 420, 418, 3, 2, 2, 2, 421, 431, 7, 41, 2, 2, 422,
	426, 7, 36, 2, 2, 423, 425, 5, 105, 53, 2, 424, 423, 3, 2, 2, 2, 425, 428,
	3, 2, 2, 2, 426, 427, 3, 2, 2, 2, 426, 424, 3, 2, 2, 2, 427, 429, 3, 2,
	2, 2, 428, 426, 3, 2, 2, 2, 429, 431, 7, 36, 2, 2, 430, 414, 3, 2, 2, 2,
	430, 422, 3, 2, 2, 2, 431, 104, 3, 2, 2, 2, 432, 437, 5, 107, 54, 2, 433,
	437, 5, 109, 55, 2, 434, 437, 5, 111, 56, 2, 435, 437, 10, 2, 2, 2, 436,
	432, 3, 2, 2, 2, 436, 433, 3, 2, 2, 2, 436, 434, 3, 2, 2, 2, 436, 435,
	3, 2, 2, 2, 437, 106, 3, 2, 2, 2, 438, 439, 7, 94, 2, 2, 439, 440, 9, 3,
	2, 2, 440, 441, 5, 137, 69, 2, 441, 442, 5, 137, 69, 2, 442, 108, 3, 2,
	2, 2, 443, 444, 7, 94, 2, 2, 444, 445, 5, 135, 68, 2, 445, 446, 5, 135,
	68, 2, 446, 447, 5, 135, 68, 2, 447, 110, 3, 2, 2, 2, 448, 449, 7, 94,
	2, 2, 449, 450, 9, 4, 2, 2, 450, 112, 3, 2, 2, 2, 451, 452, 7, 118, 2,
	2, 452, 453, 7, 116, 2, 2, 453, 454, 7, 119, 2, 2, 454, 461, 7, 103, 2,
	2, 455, 456, 7, 104, 2, 2, 456, 457, 7, 99, 2, 2, 457, 458, 7, 110, 2,
	2, 458, 459, 7, 117, 2, 2, 459, 461, 7, 103, 2, 2, 460, 451, 3, 2, 2, 2,
	460, 455, 3, 2, 2, 2, 461, 114, 3, 2, 2, 2, 462, 463, 5, 119, 60, 2, 463,
	465, 5, 93, 47, 2, 464, 466, 5, 119, 60, 2, 465, 464, 3, 2, 2, 2, 465,
	466, 3, 2, 2, 2, 466, 468, 3, 2, 2, 2, 467, 469, 5, 117, 59, 2, 468, 467,
	3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 479, 3, 2, 2, 2, 470, 471, 5, 119,
	60, 2, 471, 472, 5, 117, 59, 2, 472, 479, 3, 2, 2, 2, 473, 474, 5, 93,
	47, 2, 474, 476, 5, 119, 60, 2, 475, 477, 5, 117, 59, 2, 476, 475, 3, 2,
	2, 2, 476, 477, 3, 2, 2, 2, 477, 479, 3, 2, 2, 2, 478, 462, 3, 2, 2, 2,
	478, 470, 3, 2, 2, 2, 478, 473, 3, 2, 2, 2, 479, 487, 3, 2, 2, 2, 480,
	481, 7, 107, 2, 2, 481, 482, 7, 112, 2, 2, 482, 487, 7, 104, 2, 2, 483,
	484, 7, 112, 2, 2, 484, 485, 7, 99, 2, 2, 485, 487, 7, 112, 2, 2, 486,
	478, 3, 2, 2, 2, 486, 480, 3, 2, 2, 2, 486, 483, 3, 2, 2, 2, 487, 116,
	3, 2, 2, 2, 488, 491, 9, 5, 2, 2, 489, 492, 5, 99, 50, 2, 490, 492, 5,
	101, 51, 2, 491, 489, 3, 2, 2, 2, 491, 490, 3, 2, 2, 2, 491, 492, 3, 2,
	2, 2, 492, 493, 3, 2, 2, 2, 493, 494, 5, 119, 60, 2, 494, 118, 3, 2, 2,
	2, 495, 497, 5, 133, 67, 2, 496, 495, 3, 2, 2, 2, 497, 498, 3, 2, 2, 2,
	498, 496, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 499, 120, 3, 2, 2, 2, 500,
	504, 5, 123, 62, 2, 501, 504, 5, 125, 63, 2, 502, 504, 5, 127, 64, 2, 503,
	500, 3, 2, 2, 2, 503, 501, 3, 2, 2, 2, 503, 502, 3, 2, 2, 2, 504, 122,
	3, 2, 2, 2, 505, 509, 9, 6, 2, 2, 506, 508, 5, 133, 67, 2, 507, 506, 3,
	2, 2, 2, 508, 511, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 509, 510, 3, 2, 2,
	2, 510, 124, 3, 2, 2, 2, 511, 509, 3, 2, 2, 2, 512, 516, 7, 50, 2, 2, 513,
	515, 5, 135, 68, 2, 514, 513, 3, 2, 2, 2, 515, 518, 3, 2, 2, 2, 516, 514,
	3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517, 126, 3, 2, 2, 2, 518, 516, 3, 2,
	2, 2, 519, 520, 7, 50, 2, 2, 520, 522, 9, 3, 2, 2, 521, 523, 5, 137, 69,
	2, 522, 521, 3, 2, 2, 2, 523, 524, 3, 2, 2, 2, 524, 522, 3, 2, 2, 2, 524,
	525, 3, 2, 2, 2, 525, 128, 3, 2, 2, 2, 526, 531, 5, 131, 66, 2, 527, 530,
	5, 131, 66, 2, 528, 530, 5, 133, 67, 2, 529, 527, 3, 2, 2, 2, 529, 528,
	3, 2, 2, 2, 530, 533, 3, 2, 2, 2, 531, 529, 3, 2, 2, 2, 531, 532, 3, 2,
	2, 2, 532, 130, 3, 2, 2, 2, 533, 531, 3, 2, 2, 2, 534, 535, 9, 7, 2, 2,
	535, 132, 3, 2, 2, 2, 536, 537, 9, 8, 2, 2, 537, 134, 3, 2, 2, 2, 538,
	539, 9, 9, 2, 2, 539, 136, 3, 2, 2, 2, 540, 541, 9, 10, 2, 2, 541, 138,
	3, 2, 2, 2, 542, 544, 9, 11, 2, 2, 543, 542, 3, 2, 2, 2, 544, 545, 3, 2,
	2, 2, 545, 543, 3, 2, 2, 2, 545, 546, 3, 2, 2, 2, 546, 547, 3, 2, 2, 2,
	547, 548, 8, 70, 2, 2, 548, 140, 3, 2, 2, 2, 549, 550, 7, 49, 2, 2, 550,
	551, 7, 49, 2, 2, 551, 555, 3, 2, 2, 2, 552, 554, 10, 12, 2, 2, 553, 552,
	3, 2, 2, 2, 554, 557, 3, 2, 2, 2, 555, 553, 3, 2, 2, 2, 555, 556, 3, 2,
	2, 2, 556, 558, 3, 2, 2, 2, 557, 555, 3, 2, 2, 2, 558, 559, 8, 71, 2, 2,
	559, 142, 3, 2, 2, 2, 560, 561, 7, 49, 2, 2, 561, 562, 7, 44, 2, 2, 562,
	566, 3, 2, 2, 2, 563, 565, 11, 2, 2, 2, 564, 563, 3, 2, 2, 2, 565, 568,
	3, 2, 2, 2, 566, 567, 3, 2, 2, 2, 566, 564, 3, 2, 2, 2, 567, 569, 3, 2,
	2, 2, 568, 566, 3, 2, 2, 2, 569, 570, 7, 44, 2, 2, 570, 571, 7, 49, 2,
	2, 571, 572, 3, 2, 2, 2, 572, 573, 8, 72, 2, 2, 573, 144, 3, 2, 2, 2, 24,
	2, 418, 426, 430, 436, 460, 465, 468, 476, 478, 486, 491, 498, 503, 509,
	516, 524, 529, 531, 545, 555, 566, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'syntax'", "'import'", "'weak'", "'public'", "'package'", "'option'",
	"'repeated'", "'oneof'", "'map'", "'int32'", "'int64'", "'uint32'", "'uint64'",
	"'sint32'", "'sint64'", "'fixed32'", "'fixed64'", "'sfixed32'", "'sfixed64'",
	"'bool'", "'string'", "'double'", "'float'", "'bytes'", "'reserved'", "'to'",
	"'max'", "'enum'", "'message'", "'service'", "'rpc'", "'stream'", "'returns'",
	"'\"proto3\"'", "''proto3''", "';'", "'='", "'('", "')'", "'['", "']'",
	"'{'", "'}'", "'<'", "'>'", "'.'", "','", "':'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "SYNTAX", "IMPORT", "WEAK", "PUBLIC", "PACKAGE", "OPTION", "REPEATED",
	"ONEOF", "MAP", "INT32", "INT64", "UINT32", "UINT64", "SINT32", "SINT64",
	"FIXED32", "FIXED64", "SFIXED32", "SFIXED64", "BOOL", "STRING", "DOUBLE",
	"FLOAT", "BYTES", "RESERVED", "TO", "MAX", "ENUM", "MESSAGE", "SERVICE",
	"RPC", "STREAM", "RETURNS", "PROTO3_LIT_SINGLE", "PROTO3_LIT_DOBULE", "SEMI",
	"EQ", "LP", "RP", "LB", "RB", "LC", "RC", "LT", "GT", "DOT", "COMMA", "COLON",
	"PLUS", "MINUS", "STR_LIT", "BOOL_LIT", "FLOAT_LIT", "INT_LIT", "IDENTIFIER",
	"WS", "LINE_COMMENT", "COMMENT",
}

var lexerRuleNames = []string{
	"SYNTAX", "IMPORT", "WEAK", "PUBLIC", "PACKAGE", "OPTION", "REPEATED",
	"ONEOF", "MAP", "INT32", "INT64", "UINT32", "UINT64", "SINT32", "SINT64",
	"FIXED32", "FIXED64", "SFIXED32", "SFIXED64", "BOOL", "STRING", "DOUBLE",
	"FLOAT", "BYTES", "RESERVED", "TO", "MAX", "ENUM", "MESSAGE", "SERVICE",
	"RPC", "STREAM", "RETURNS", "PROTO3_LIT_SINGLE", "PROTO3_LIT_DOBULE", "SEMI",
	"EQ", "LP", "RP", "LB", "RB", "LC", "RC", "LT", "GT", "DOT", "COMMA", "COLON",
	"PLUS", "MINUS", "STR_LIT", "CHAR_VALUE", "HEX_ESCAPE", "OCT_ESCAPE", "CHAR_ESCAPE",
	"BOOL_LIT", "FLOAT_LIT", "EXPONENT", "DECIMALS", "INT_LIT", "DECIMAL_LIT",
	"OCTAL_LIT", "HEX_LIT", "IDENTIFIER", "LETTER", "DECIMAL_DIGIT", "OCTAL_DIGIT",
	"HEX_DIGIT", "WS", "LINE_COMMENT", "COMMENT",
}

type Protobuf3Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewProtobuf3Lexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *Protobuf3Lexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewProtobuf3Lexer(input antlr.CharStream) *Protobuf3Lexer {
	l := new(Protobuf3Lexer)
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
	l.GrammarFileName = "Protobuf3.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Protobuf3Lexer tokens.
const (
	Protobuf3LexerSYNTAX            = 1
	Protobuf3LexerIMPORT            = 2
	Protobuf3LexerWEAK              = 3
	Protobuf3LexerPUBLIC            = 4
	Protobuf3LexerPACKAGE           = 5
	Protobuf3LexerOPTION            = 6
	Protobuf3LexerREPEATED          = 7
	Protobuf3LexerONEOF             = 8
	Protobuf3LexerMAP               = 9
	Protobuf3LexerINT32             = 10
	Protobuf3LexerINT64             = 11
	Protobuf3LexerUINT32            = 12
	Protobuf3LexerUINT64            = 13
	Protobuf3LexerSINT32            = 14
	Protobuf3LexerSINT64            = 15
	Protobuf3LexerFIXED32           = 16
	Protobuf3LexerFIXED64           = 17
	Protobuf3LexerSFIXED32          = 18
	Protobuf3LexerSFIXED64          = 19
	Protobuf3LexerBOOL              = 20
	Protobuf3LexerSTRING            = 21
	Protobuf3LexerDOUBLE            = 22
	Protobuf3LexerFLOAT             = 23
	Protobuf3LexerBYTES             = 24
	Protobuf3LexerRESERVED          = 25
	Protobuf3LexerTO                = 26
	Protobuf3LexerMAX               = 27
	Protobuf3LexerENUM              = 28
	Protobuf3LexerMESSAGE           = 29
	Protobuf3LexerSERVICE           = 30
	Protobuf3LexerRPC               = 31
	Protobuf3LexerSTREAM            = 32
	Protobuf3LexerRETURNS           = 33
	Protobuf3LexerPROTO3_LIT_SINGLE = 34
	Protobuf3LexerPROTO3_LIT_DOBULE = 35
	Protobuf3LexerSEMI              = 36
	Protobuf3LexerEQ                = 37
	Protobuf3LexerLP                = 38
	Protobuf3LexerRP                = 39
	Protobuf3LexerLB                = 40
	Protobuf3LexerRB                = 41
	Protobuf3LexerLC                = 42
	Protobuf3LexerRC                = 43
	Protobuf3LexerLT                = 44
	Protobuf3LexerGT                = 45
	Protobuf3LexerDOT               = 46
	Protobuf3LexerCOMMA             = 47
	Protobuf3LexerCOLON             = 48
	Protobuf3LexerPLUS              = 49
	Protobuf3LexerMINUS             = 50
	Protobuf3LexerSTR_LIT           = 51
	Protobuf3LexerBOOL_LIT          = 52
	Protobuf3LexerFLOAT_LIT         = 53
	Protobuf3LexerINT_LIT           = 54
	Protobuf3LexerIDENTIFIER        = 55
	Protobuf3LexerWS                = 56
	Protobuf3LexerLINE_COMMENT      = 57
	Protobuf3LexerCOMMENT           = 58
)