// Code generated from WavefrontOBJ.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wavefrontobj

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 63, 462,
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
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 5, 53, 394, 10, 53,
	3, 53, 3, 53, 3, 53, 3, 54, 5, 54, 400, 10, 54, 3, 54, 6, 54, 403, 10,
	54, 13, 54, 14, 54, 404, 3, 55, 3, 55, 3, 55, 7, 55, 410, 10, 55, 12, 55,
	14, 55, 413, 11, 55, 5, 55, 415, 10, 55, 3, 56, 3, 56, 3, 57, 3, 57, 7,
	57, 421, 10, 57, 12, 57, 14, 57, 424, 11, 57, 3, 57, 3, 57, 3, 57, 3, 57,
	3, 58, 6, 58, 431, 10, 58, 13, 58, 14, 58, 432, 3, 59, 6, 59, 436, 10,
	59, 13, 59, 14, 59, 437, 3, 60, 3, 60, 3, 60, 6, 60, 443, 10, 60, 13, 60,
	14, 60, 444, 3, 60, 3, 60, 3, 61, 3, 61, 5, 61, 451, 10, 61, 3, 61, 5,
	61, 454, 10, 61, 3, 62, 3, 62, 3, 63, 6, 63, 459, 10, 63, 13, 63, 14, 63,
	460, 2, 2, 64, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73,
	38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91,
	47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55,
	109, 56, 111, 2, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62,
	125, 63, 3, 2, 7, 9, 2, 42, 43, 47, 47, 50, 59, 67, 92, 97, 97, 99, 124,
	128, 128, 6, 2, 11, 12, 15, 15, 34, 34, 49, 49, 4, 2, 11, 11, 34, 34, 4,
	2, 12, 12, 15, 15, 5, 2, 11, 12, 15, 15, 34, 34, 2, 473, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35,
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2,
	43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2,
	2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2,
	2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2,
	2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3,
	2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81,
	3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2,
	89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2,
	2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2,
	2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 113,
	3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2,
	2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 3, 127, 3,
	2, 2, 2, 5, 129, 3, 2, 2, 2, 7, 131, 3, 2, 2, 2, 9, 134, 3, 2, 2, 2, 11,
	137, 3, 2, 2, 2, 13, 140, 3, 2, 2, 2, 15, 147, 3, 2, 2, 2, 17, 151, 3,
	2, 2, 2, 19, 159, 3, 2, 2, 2, 21, 166, 3, 2, 2, 2, 23, 174, 3, 2, 2, 2,
	25, 183, 3, 2, 2, 2, 27, 190, 3, 2, 2, 2, 29, 194, 3, 2, 2, 2, 31, 199,
	3, 2, 2, 2, 33, 201, 3, 2, 2, 2, 35, 203, 3, 2, 2, 2, 37, 205, 3, 2, 2,
	2, 39, 207, 3, 2, 2, 2, 41, 212, 3, 2, 2, 2, 43, 218, 3, 2, 2, 2, 45, 223,
	3, 2, 2, 2, 47, 228, 3, 2, 2, 2, 49, 233, 3, 2, 2, 2, 51, 238, 3, 2, 2,
	2, 53, 243, 3, 2, 2, 2, 55, 246, 3, 2, 2, 2, 57, 250, 3, 2, 2, 2, 59, 254,
	3, 2, 2, 2, 61, 256, 3, 2, 2, 2, 63, 258, 3, 2, 2, 2, 65, 262, 3, 2, 2,
	2, 67, 265, 3, 2, 2, 2, 69, 267, 3, 2, 2, 2, 71, 273, 3, 2, 2, 2, 73, 276,
	3, 2, 2, 2, 75, 285, 3, 2, 2, 2, 77, 294, 3, 2, 2, 2, 79, 298, 3, 2, 2,
	2, 81, 305, 3, 2, 2, 2, 83, 312, 3, 2, 2, 2, 85, 319, 3, 2, 2, 2, 87, 326,
	3, 2, 2, 2, 89, 337, 3, 2, 2, 2, 91, 347, 3, 2, 2, 2, 93, 353, 3, 2, 2,
	2, 95, 359, 3, 2, 2, 2, 97, 366, 3, 2, 2, 2, 99, 372, 3, 2, 2, 2, 101,
	379, 3, 2, 2, 2, 103, 386, 3, 2, 2, 2, 105, 390, 3, 2, 2, 2, 107, 399,
	3, 2, 2, 2, 109, 406, 3, 2, 2, 2, 111, 416, 3, 2, 2, 2, 113, 418, 3, 2,
	2, 2, 115, 430, 3, 2, 2, 2, 117, 435, 3, 2, 2, 2, 119, 442, 3, 2, 2, 2,
	121, 453, 3, 2, 2, 2, 123, 455, 3, 2, 2, 2, 125, 458, 3, 2, 2, 2, 127,
	128, 7, 47, 2, 2, 128, 4, 3, 2, 2, 2, 129, 130, 7, 120, 2, 2, 130, 6, 3,
	2, 2, 2, 131, 132, 7, 120, 2, 2, 132, 133, 7, 114, 2, 2, 133, 8, 3, 2,
	2, 2, 134, 135, 7, 120, 2, 2, 135, 136, 7, 112, 2, 2, 136, 10, 3, 2, 2,
	2, 137, 138, 7, 120, 2, 2, 138, 139, 7, 118, 2, 2, 139, 12, 3, 2, 2, 2,
	140, 141, 7, 101, 2, 2, 141, 142, 7, 117, 2, 2, 142, 143, 7, 118, 2, 2,
	143, 144, 7, 123, 2, 2, 144, 145, 7, 114, 2, 2, 145, 146, 7, 103, 2, 2,
	146, 14, 3, 2, 2, 2, 147, 148, 7, 116, 2, 2, 148, 149, 7, 99, 2, 2, 149,
	150, 7, 118, 2, 2, 150, 16, 3, 2, 2, 2, 151, 152, 7, 100, 2, 2, 152, 153,
	7, 111, 2, 2, 153, 154, 7, 99, 2, 2, 154, 155, 7, 118, 2, 2, 155, 156,
	7, 116, 2, 2, 156, 157, 7, 107, 2, 2, 157, 158, 7, 122, 2, 2, 158, 18,
	3, 2, 2, 2, 159, 160, 7, 100, 2, 2, 160, 161, 7, 103, 2, 2, 161, 162, 7,
	124, 2, 2, 162, 163, 7, 107, 2, 2, 163, 164, 7, 103, 2, 2, 164, 165, 7,
	116, 2, 2, 165, 20, 3, 2, 2, 2, 166, 167, 7, 100, 2, 2, 167, 168, 7, 117,
	2, 2, 168, 169, 7, 114, 2, 2, 169, 170, 7, 110, 2, 2, 170, 171, 7, 107,
	2, 2, 171, 172, 7, 112, 2, 2, 172, 173, 7, 103, 2, 2, 173, 22, 3, 2, 2,
	2, 174, 175, 7, 101, 2, 2, 175, 176, 7, 99, 2, 2, 176, 177, 7, 116, 2,
	2, 177, 178, 7, 102, 2, 2, 178, 179, 7, 107, 2, 2, 179, 180, 7, 112, 2,
	2, 180, 181, 7, 99, 2, 2, 181, 182, 7, 110, 2, 2, 182, 24, 3, 2, 2, 2,
	183, 184, 7, 118, 2, 2, 184, 185, 7, 99, 2, 2, 185, 186, 7, 123, 2, 2,
	186, 187, 7, 110, 2, 2, 187, 188, 7, 113, 2, 2, 188, 189, 7, 116, 2, 2,
	189, 26, 3, 2, 2, 2, 190, 191, 7, 102, 2, 2, 191, 192, 7, 103, 2, 2, 192,
	193, 7, 105, 2, 2, 193, 28, 3, 2, 2, 2, 194, 195, 7, 100, 2, 2, 195, 196,
	7, 111, 2, 2, 196, 197, 7, 99, 2, 2, 197, 198, 7, 118, 2, 2, 198, 30, 3,
	2, 2, 2, 199, 200, 7, 119, 2, 2, 200, 32, 3, 2, 2, 2, 201, 202, 7, 114,
	2, 2, 202, 34, 3, 2, 2, 2, 203, 204, 7, 110, 2, 2, 204, 36, 3, 2, 2, 2,
	205, 206, 7, 104, 2, 2, 206, 38, 3, 2, 2, 2, 207, 208, 7, 101, 2, 2, 208,
	209, 7, 119, 2, 2, 209, 210, 7, 116, 2, 2, 210, 211, 7, 120, 2, 2, 211,
	40, 3, 2, 2, 2, 212, 213, 7, 101, 2, 2, 213, 214, 7, 119, 2, 2, 214, 215,
	7, 116, 2, 2, 215, 216, 7, 120, 2, 2, 216, 217, 7, 52, 2, 2, 217, 42, 3,
	2, 2, 2, 218, 219, 7, 117, 2, 2, 219, 220, 7, 119, 2, 2, 220, 221, 7, 116,
	2, 2, 221, 222, 7, 104, 2, 2, 222, 44, 3, 2, 2, 2, 223, 224, 7, 114, 2,
	2, 224, 225, 7, 99, 2, 2, 225, 226, 7, 116, 2, 2, 226, 227, 7, 111, 2,
	2, 227, 46, 3, 2, 2, 2, 228, 229, 7, 118, 2, 2, 229, 230, 7, 116, 2, 2,
	230, 231, 7, 107, 2, 2, 231, 232, 7, 111, 2, 2, 232, 48, 3, 2, 2, 2, 233,
	234, 7, 106, 2, 2, 234, 235, 7, 113, 2, 2, 235, 236, 7, 110, 2, 2, 236,
	237, 7, 103, 2, 2, 237, 50, 3, 2, 2, 2, 238, 239, 7, 117, 2, 2, 239, 240,
	7, 101, 2, 2, 240, 241, 7, 116, 2, 2, 241, 242, 7, 120, 2, 2, 242, 52,
	3, 2, 2, 2, 243, 244, 7, 117, 2, 2, 244, 245, 7, 114, 2, 2, 245, 54, 3,
	2, 2, 2, 246, 247, 7, 103, 2, 2, 247, 248, 7, 112, 2, 2, 248, 249, 7, 102,
	2, 2, 249, 56, 3, 2, 2, 2, 250, 251, 7, 101, 2, 2, 251, 252, 7, 113, 2,
	2, 252, 253, 7, 112, 2, 2, 253, 58, 3, 2, 2, 2, 254, 255, 7, 105, 2, 2,
	255, 60, 3, 2, 2, 2, 256, 257, 7, 117, 2, 2, 257, 62, 3, 2, 2, 2, 258,
	259, 7, 113, 2, 2, 259, 260, 7, 104, 2, 2, 260, 261, 7, 104, 2, 2, 261,
	64, 3, 2, 2, 2, 262, 263, 7, 111, 2, 2, 263, 264, 7, 105, 2, 2, 264, 66,
	3, 2, 2, 2, 265, 266, 7, 113, 2, 2, 266, 68, 3, 2, 2, 2, 267, 268, 7, 100,
	2, 2, 268, 269, 7, 103, 2, 2, 269, 270, 7, 120, 2, 2, 270, 271, 7, 103,
	2, 2, 271, 272, 7, 110, 2, 2, 272, 70, 3, 2, 2, 2, 273, 274, 7, 113, 2,
	2, 274, 275, 7, 112, 2, 2, 275, 72, 3, 2, 2, 2, 276, 277, 7, 101, 2, 2,
	277, 278, 7, 97, 2, 2, 278, 279, 7, 107, 2, 2, 279, 280, 7, 112, 2, 2,
	280, 281, 7, 118, 2, 2, 281, 282, 7, 103, 2, 2, 282, 283, 7, 116, 2, 2,
	283, 284, 7, 114, 2, 2, 284, 74, 3, 2, 2, 2, 285, 286, 7, 102, 2, 2, 286,
	287, 7, 97, 2, 2, 287, 288, 7, 107, 2, 2, 288, 289, 7, 112, 2, 2, 289,
	290, 7, 118, 2, 2, 290, 291, 7, 103, 2, 2, 291, 292, 7, 116, 2, 2, 292,
	293, 7, 114, 2, 2, 293, 76, 3, 2, 2, 2, 294, 295, 7, 110, 2, 2, 295, 296,
	7, 113, 2, 2, 296, 297, 7, 102, 2, 2, 297, 78, 3, 2, 2, 2, 298, 299, 7,
	111, 2, 2, 299, 300, 7, 99, 2, 2, 300, 301, 7, 114, 2, 2, 301, 302, 7,
	110, 2, 2, 302, 303, 7, 107, 2, 2, 303, 304, 7, 100, 2, 2, 304, 80, 3,
	2, 2, 2, 305, 306, 7, 119, 2, 2, 306, 307, 7, 117, 2, 2, 307, 308, 7, 103,
	2, 2, 308, 309, 7, 111, 2, 2, 309, 310, 7, 99, 2, 2, 310, 311, 7, 114,
	2, 2, 311, 82, 3, 2, 2, 2, 312, 313, 7, 111, 2, 2, 313, 314, 7, 118, 2,
	2, 314, 315, 7, 110, 2, 2, 315, 316, 7, 110, 2, 2, 316, 317, 7, 107, 2,
	2, 317, 318, 7, 100, 2, 2, 318, 84, 3, 2, 2, 2, 319, 320, 7, 119, 2, 2,
	320, 321, 7, 117, 2, 2, 321, 322, 7, 103, 2, 2, 322, 323, 7, 111, 2, 2,
	323, 324, 7, 118, 2, 2, 324, 325, 7, 110, 2, 2, 325, 86, 3, 2, 2, 2, 326,
	327, 7, 117, 2, 2, 327, 328, 7, 106, 2, 2, 328, 329, 7, 99, 2, 2, 329,
	330, 7, 102, 2, 2, 330, 331, 7, 113, 2, 2, 331, 332, 7, 121, 2, 2, 332,
	333, 7, 97, 2, 2, 333, 334, 7, 113, 2, 2, 334, 335, 7, 100, 2, 2, 335,
	336, 7, 108, 2, 2, 336, 88, 3, 2, 2, 2, 337, 338, 7, 118, 2, 2, 338, 339,
	7, 116, 2, 2, 339, 340, 7, 99, 2, 2, 340, 341, 7, 101, 2, 2, 341, 342,
	7, 103, 2, 2, 342, 343, 7, 97, 2, 2, 343, 344, 7, 113, 2, 2, 344, 345,
	7, 100, 2, 2, 345, 346, 7, 108, 2, 2, 346, 90, 3, 2, 2, 2, 347, 348, 7,
	101, 2, 2, 348, 349, 7, 118, 2, 2, 349, 350, 7, 103, 2, 2, 350, 351, 7,
	101, 2, 2, 351, 352, 7, 106, 2, 2, 352, 92, 3, 2, 2, 2, 353, 354, 7, 101,
	2, 2, 354, 355, 7, 114, 2, 2, 355, 356, 7, 99, 2, 2, 356, 357, 7, 116,
	2, 2, 357, 358, 7, 111, 2, 2, 358, 94, 3, 2, 2, 2, 359, 360, 7, 101, 2,
	2, 360, 361, 7, 117, 2, 2, 361, 362, 7, 114, 2, 2, 362, 363, 7, 99, 2,
	2, 363, 364, 7, 101, 2, 2, 364, 365, 7, 103, 2, 2, 365, 96, 3, 2, 2, 2,
	366, 367, 7, 117, 2, 2, 367, 368, 7, 118, 2, 2, 368, 369, 7, 103, 2, 2,
	369, 370, 7, 101, 2, 2, 370, 371, 7, 106, 2, 2, 371, 98, 3, 2, 2, 2, 372,
	373, 7, 101, 2, 2, 373, 374, 7, 114, 2, 2, 374, 375, 7, 99, 2, 2, 375,
	376, 7, 116, 2, 2, 376, 377, 7, 111, 2, 2, 377, 378, 7, 99, 2, 2, 378,
	100, 3, 2, 2, 2, 379, 380, 7, 101, 2, 2, 380, 381, 7, 114, 2, 2, 381, 382,
	7, 99, 2, 2, 382, 383, 7, 116, 2, 2, 383, 384, 7, 111, 2, 2, 384, 385,
	7, 100, 2, 2, 385, 102, 3, 2, 2, 2, 386, 387, 5, 107, 54, 2, 387, 388,
	7, 49, 2, 2, 388, 389, 5, 107, 54, 2, 389, 104, 3, 2, 2, 2, 390, 391, 5,
	107, 54, 2, 391, 393, 7, 49, 2, 2, 392, 394, 5, 107, 54, 2, 393, 392, 3,
	2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 395, 3, 2, 2, 2, 395, 396, 7, 49, 2,
	2, 396, 397, 5, 107, 54, 2, 397, 106, 3, 2, 2, 2, 398, 400, 7, 47, 2, 2,
	399, 398, 3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 402, 3, 2, 2, 2, 401,
	403, 5, 111, 56, 2, 402, 401, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 402,
	3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 108, 3, 2, 2, 2, 406, 414, 5, 107,
	54, 2, 407, 411, 7, 48, 2, 2, 408, 410, 5, 111, 56, 2, 409, 408, 3, 2,
	2, 2, 410, 413, 3, 2, 2, 2, 411, 409, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2,
	412, 415, 3, 2, 2, 2, 413, 411, 3, 2, 2, 2, 414, 407, 3, 2, 2, 2, 414,
	415, 3, 2, 2, 2, 415, 110, 3, 2, 2, 2, 416, 417, 4, 50, 59, 2, 417, 112,
	3, 2, 2, 2, 418, 422, 7, 37, 2, 2, 419, 421, 5, 123, 62, 2, 420, 419, 3,
	2, 2, 2, 421, 424, 3, 2, 2, 2, 422, 420, 3, 2, 2, 2, 422, 423, 3, 2, 2,
	2, 423, 425, 3, 2, 2, 2, 424, 422, 3, 2, 2, 2, 425, 426, 5, 121, 61, 2,
	426, 427, 3, 2, 2, 2, 427, 428, 8, 57, 2, 2, 428, 114, 3, 2, 2, 2, 429,
	431, 9, 2, 2, 2, 430, 429, 3, 2, 2, 2, 431, 432, 3, 2, 2, 2, 432, 430,
	3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 116, 3, 2, 2, 2, 434, 436, 10, 3,
	2, 2, 435, 434, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437, 435, 3, 2, 2, 2,
	437, 438, 3, 2, 2, 2, 438, 118, 3, 2, 2, 2, 439, 443, 9, 4, 2, 2, 440,
	441, 7, 94, 2, 2, 441, 443, 5, 121, 61, 2, 442, 439, 3, 2, 2, 2, 442, 440,
	3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 442, 3, 2, 2, 2, 444, 445, 3, 2,
	2, 2, 445, 446, 3, 2, 2, 2, 446, 447, 8, 60, 2, 2, 447, 120, 3, 2, 2, 2,
	448, 450, 7, 15, 2, 2, 449, 451, 7, 12, 2, 2, 450, 449, 3, 2, 2, 2, 450,
	451, 3, 2, 2, 2, 451, 454, 3, 2, 2, 2, 452, 454, 7, 12, 2, 2, 453, 448,
	3, 2, 2, 2, 453, 452, 3, 2, 2, 2, 454, 122, 3, 2, 2, 2, 455, 456, 10, 5,
	2, 2, 456, 124, 3, 2, 2, 2, 457, 459, 10, 6, 2, 2, 458, 457, 3, 2, 2, 2,
	459, 460, 3, 2, 2, 2, 460, 458, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461,
	126, 3, 2, 2, 2, 16, 2, 393, 399, 404, 411, 414, 422, 432, 437, 442, 444,
	450, 453, 460, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'-'", "'v'", "'vp'", "'vn'", "'vt'", "'cstype'", "'rat'", "'bmatrix'",
	"'bezier'", "'bspline'", "'cardinal'", "'taylor'", "'deg'", "'bmat'", "'u'",
	"'p'", "'l'", "'f'", "'curv'", "'curv2'", "'surf'", "'parm'", "'trim'",
	"'hole'", "'scrv'", "'sp'", "'end'", "'con'", "'g'", "'s'", "'off'", "'mg'",
	"'o'", "'bevel'", "'on'", "'c_interp'", "'d_interp'", "'lod'", "'maplib'",
	"'usemap'", "'mtllib'", "'usemtl'", "'shadow_obj'", "'trace_obj'", "'ctech'",
	"'cparm'", "'cspace'", "'stech'", "'cparma'", "'cparmb'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "INTEGER_PAIR",
	"INTEGER_TRIPLET", "INTEGER", "DECIMAL", "COMMENT", "NAME", "FILENAME",
	"WS", "NL", "NON_NL", "NON_WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
	"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
	"T__49", "INTEGER_PAIR", "INTEGER_TRIPLET", "INTEGER", "DECIMAL", "DIGIT",
	"COMMENT", "NAME", "FILENAME", "WS", "NL", "NON_NL", "NON_WS",
}

type WavefrontOBJLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewWavefrontOBJLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *WavefrontOBJLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewWavefrontOBJLexer(input antlr.CharStream) *WavefrontOBJLexer {
	l := new(WavefrontOBJLexer)
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
	l.GrammarFileName = "WavefrontOBJ.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// WavefrontOBJLexer tokens.
const (
	WavefrontOBJLexerT__0            = 1
	WavefrontOBJLexerT__1            = 2
	WavefrontOBJLexerT__2            = 3
	WavefrontOBJLexerT__3            = 4
	WavefrontOBJLexerT__4            = 5
	WavefrontOBJLexerT__5            = 6
	WavefrontOBJLexerT__6            = 7
	WavefrontOBJLexerT__7            = 8
	WavefrontOBJLexerT__8            = 9
	WavefrontOBJLexerT__9            = 10
	WavefrontOBJLexerT__10           = 11
	WavefrontOBJLexerT__11           = 12
	WavefrontOBJLexerT__12           = 13
	WavefrontOBJLexerT__13           = 14
	WavefrontOBJLexerT__14           = 15
	WavefrontOBJLexerT__15           = 16
	WavefrontOBJLexerT__16           = 17
	WavefrontOBJLexerT__17           = 18
	WavefrontOBJLexerT__18           = 19
	WavefrontOBJLexerT__19           = 20
	WavefrontOBJLexerT__20           = 21
	WavefrontOBJLexerT__21           = 22
	WavefrontOBJLexerT__22           = 23
	WavefrontOBJLexerT__23           = 24
	WavefrontOBJLexerT__24           = 25
	WavefrontOBJLexerT__25           = 26
	WavefrontOBJLexerT__26           = 27
	WavefrontOBJLexerT__27           = 28
	WavefrontOBJLexerT__28           = 29
	WavefrontOBJLexerT__29           = 30
	WavefrontOBJLexerT__30           = 31
	WavefrontOBJLexerT__31           = 32
	WavefrontOBJLexerT__32           = 33
	WavefrontOBJLexerT__33           = 34
	WavefrontOBJLexerT__34           = 35
	WavefrontOBJLexerT__35           = 36
	WavefrontOBJLexerT__36           = 37
	WavefrontOBJLexerT__37           = 38
	WavefrontOBJLexerT__38           = 39
	WavefrontOBJLexerT__39           = 40
	WavefrontOBJLexerT__40           = 41
	WavefrontOBJLexerT__41           = 42
	WavefrontOBJLexerT__42           = 43
	WavefrontOBJLexerT__43           = 44
	WavefrontOBJLexerT__44           = 45
	WavefrontOBJLexerT__45           = 46
	WavefrontOBJLexerT__46           = 47
	WavefrontOBJLexerT__47           = 48
	WavefrontOBJLexerT__48           = 49
	WavefrontOBJLexerT__49           = 50
	WavefrontOBJLexerINTEGER_PAIR    = 51
	WavefrontOBJLexerINTEGER_TRIPLET = 52
	WavefrontOBJLexerINTEGER         = 53
	WavefrontOBJLexerDECIMAL         = 54
	WavefrontOBJLexerCOMMENT         = 55
	WavefrontOBJLexerNAME            = 56
	WavefrontOBJLexerFILENAME        = 57
	WavefrontOBJLexerWS              = 58
	WavefrontOBJLexerNL              = 59
	WavefrontOBJLexerNON_NL          = 60
	WavefrontOBJLexerNON_WS          = 61
)
