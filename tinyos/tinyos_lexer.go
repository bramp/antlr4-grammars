// Code generated from TinyosLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package tinyos

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 101, 724,
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
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96,
	4, 97, 9, 97, 4, 98, 9, 98, 4, 99, 9, 99, 4, 100, 9, 100, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3,
	49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51,
	3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3,
	53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 55,
	3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3,
	59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63,
	3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3,
	68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73,
	3, 73, 3, 74, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77, 3, 77, 3,
	78, 3, 78, 3, 79, 3, 79, 3, 80, 3, 80, 3, 81, 3, 81, 3, 82, 3, 82, 3, 83,
	3, 83, 3, 84, 3, 84, 3, 85, 3, 85, 3, 85, 3, 86, 3, 86, 3, 86, 3, 87, 3,
	87, 3, 87, 3, 88, 3, 88, 3, 89, 3, 89, 3, 90, 3, 90, 3, 91, 3, 91, 3, 92,
	3, 92, 7, 92, 642, 10, 92, 12, 92, 14, 92, 645, 11, 92, 3, 93, 6, 93, 648,
	10, 93, 13, 93, 14, 93, 649, 3, 94, 6, 94, 653, 10, 94, 13, 94, 14, 94,
	654, 3, 94, 3, 94, 7, 94, 659, 10, 94, 12, 94, 14, 94, 662, 11, 94, 3,
	94, 3, 94, 6, 94, 666, 10, 94, 13, 94, 14, 94, 667, 5, 94, 670, 10, 94,
	3, 95, 3, 95, 3, 95, 3, 95, 7, 95, 676, 10, 95, 12, 95, 14, 95, 679, 11,
	95, 3, 95, 3, 95, 3, 96, 3, 96, 3, 96, 3, 96, 7, 96, 687, 10, 96, 12, 96,
	14, 96, 690, 11, 96, 3, 96, 3, 96, 3, 96, 3, 96, 3, 96, 3, 97, 3, 97, 3,
	97, 3, 97, 7, 97, 701, 10, 97, 12, 97, 14, 97, 704, 11, 97, 3, 97, 3, 97,
	3, 98, 6, 98, 709, 10, 98, 13, 98, 14, 98, 710, 3, 98, 3, 98, 3, 99, 3,
	99, 3, 100, 3, 100, 3, 100, 3, 100, 6, 100, 721, 10, 100, 13, 100, 14,
	100, 722, 3, 688, 2, 101, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18,
	35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27,
	53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36,
	71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45,
	89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105,
	54, 107, 55, 109, 56, 111, 57, 113, 58, 115, 59, 117, 60, 119, 61, 121,
	62, 123, 63, 125, 64, 127, 65, 129, 66, 131, 67, 133, 68, 135, 69, 137,
	70, 139, 71, 141, 72, 143, 73, 145, 74, 147, 75, 149, 76, 151, 77, 153,
	78, 155, 79, 157, 80, 159, 81, 161, 82, 163, 83, 165, 84, 167, 85, 169,
	86, 171, 87, 173, 88, 175, 89, 177, 90, 179, 91, 181, 92, 183, 93, 185,
	94, 187, 95, 189, 96, 191, 97, 193, 98, 195, 99, 197, 100, 199, 101, 3,
	2, 9, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99,
	124, 3, 2, 50, 59, 5, 2, 12, 12, 15, 15, 36, 36, 4, 2, 12, 12, 15, 15,
	5, 2, 11, 12, 15, 15, 34, 34, 5, 2, 50, 59, 67, 72, 99, 104, 2, 735, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2,
	2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3,
	2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49,
	3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2,
	57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2,
	2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2,
	2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2,
	2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3,
	2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95,
	3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2,
	103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2,
	2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117,
	3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2,
	2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3,
	2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2,
	139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2,
	2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153,
	3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 2, 159, 3, 2, 2, 2,
	2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2, 165, 3, 2, 2, 2, 2, 167, 3,
	2, 2, 2, 2, 169, 3, 2, 2, 2, 2, 171, 3, 2, 2, 2, 2, 173, 3, 2, 2, 2, 2,
	175, 3, 2, 2, 2, 2, 177, 3, 2, 2, 2, 2, 179, 3, 2, 2, 2, 2, 181, 3, 2,
	2, 2, 2, 183, 3, 2, 2, 2, 2, 185, 3, 2, 2, 2, 2, 187, 3, 2, 2, 2, 2, 189,
	3, 2, 2, 2, 2, 191, 3, 2, 2, 2, 2, 193, 3, 2, 2, 2, 2, 195, 3, 2, 2, 2,
	2, 197, 3, 2, 2, 2, 2, 199, 3, 2, 2, 2, 3, 201, 3, 2, 2, 2, 5, 210, 3,
	2, 2, 2, 7, 213, 3, 2, 2, 2, 9, 219, 3, 2, 2, 2, 11, 226, 3, 2, 2, 2, 13,
	234, 3, 2, 2, 2, 15, 240, 3, 2, 2, 2, 17, 245, 3, 2, 2, 2, 19, 250, 3,
	2, 2, 2, 21, 255, 3, 2, 2, 2, 23, 269, 3, 2, 2, 2, 25, 277, 3, 2, 2, 2,
	27, 287, 3, 2, 2, 2, 29, 298, 3, 2, 2, 2, 31, 307, 3, 2, 2, 2, 33, 310,
	3, 2, 2, 2, 35, 317, 3, 2, 2, 2, 37, 325, 3, 2, 2, 2, 39, 332, 3, 2, 2,
	2, 41, 340, 3, 2, 2, 2, 43, 345, 3, 2, 2, 2, 45, 350, 3, 2, 2, 2, 47, 356,
	3, 2, 2, 2, 49, 364, 3, 2, 2, 2, 51, 369, 3, 2, 2, 2, 53, 375, 3, 2, 2,
	2, 55, 381, 3, 2, 2, 2, 57, 387, 3, 2, 2, 2, 59, 393, 3, 2, 2, 2, 61, 397,
	3, 2, 2, 2, 63, 405, 3, 2, 2, 2, 65, 408, 3, 2, 2, 2, 67, 423, 3, 2, 2,
	2, 69, 431, 3, 2, 2, 2, 71, 440, 3, 2, 2, 2, 73, 450, 3, 2, 2, 2, 75, 454,
	3, 2, 2, 2, 77, 459, 3, 2, 2, 2, 79, 466, 3, 2, 2, 2, 81, 470, 3, 2, 2,
	2, 83, 475, 3, 2, 2, 2, 85, 484, 3, 2, 2, 2, 87, 491, 3, 2, 2, 2, 89, 497,
	3, 2, 2, 2, 91, 504, 3, 2, 2, 2, 93, 511, 3, 2, 2, 2, 95, 518, 3, 2, 2,
	2, 97, 523, 3, 2, 2, 2, 99, 528, 3, 2, 2, 2, 101, 533, 3, 2, 2, 2, 103,
	538, 3, 2, 2, 2, 105, 544, 3, 2, 2, 2, 107, 552, 3, 2, 2, 2, 109, 555,
	3, 2, 2, 2, 111, 558, 3, 2, 2, 2, 113, 561, 3, 2, 2, 2, 115, 564, 3, 2,
	2, 2, 117, 566, 3, 2, 2, 2, 119, 568, 3, 2, 2, 2, 121, 571, 3, 2, 2, 2,
	123, 574, 3, 2, 2, 2, 125, 576, 3, 2, 2, 2, 127, 578, 3, 2, 2, 2, 129,
	580, 3, 2, 2, 2, 131, 582, 3, 2, 2, 2, 133, 584, 3, 2, 2, 2, 135, 586,
	3, 2, 2, 2, 137, 588, 3, 2, 2, 2, 139, 590, 3, 2, 2, 2, 141, 592, 3, 2,
	2, 2, 143, 594, 3, 2, 2, 2, 145, 596, 3, 2, 2, 2, 147, 599, 3, 2, 2, 2,
	149, 602, 3, 2, 2, 2, 151, 604, 3, 2, 2, 2, 153, 606, 3, 2, 2, 2, 155,
	608, 3, 2, 2, 2, 157, 610, 3, 2, 2, 2, 159, 612, 3, 2, 2, 2, 161, 614,
	3, 2, 2, 2, 163, 616, 3, 2, 2, 2, 165, 618, 3, 2, 2, 2, 167, 620, 3, 2,
	2, 2, 169, 622, 3, 2, 2, 2, 171, 625, 3, 2, 2, 2, 173, 628, 3, 2, 2, 2,
	175, 631, 3, 2, 2, 2, 177, 633, 3, 2, 2, 2, 179, 635, 3, 2, 2, 2, 181,
	637, 3, 2, 2, 2, 183, 639, 3, 2, 2, 2, 185, 647, 3, 2, 2, 2, 187, 669,
	3, 2, 2, 2, 189, 671, 3, 2, 2, 2, 191, 682, 3, 2, 2, 2, 193, 696, 3, 2,
	2, 2, 195, 708, 3, 2, 2, 2, 197, 714, 3, 2, 2, 2, 199, 716, 3, 2, 2, 2,
	201, 202, 7, 99, 2, 2, 202, 203, 7, 100, 2, 2, 203, 204, 7, 117, 2, 2,
	204, 205, 7, 118, 2, 2, 205, 206, 7, 116, 2, 2, 206, 207, 7, 99, 2, 2,
	207, 208, 7, 101, 2, 2, 208, 209, 7, 118, 2, 2, 209, 4, 3, 2, 2, 2, 210,
	211, 7, 99, 2, 2, 211, 212, 7, 117, 2, 2, 212, 6, 3, 2, 2, 2, 213, 214,
	7, 99, 2, 2, 214, 215, 7, 117, 2, 2, 215, 216, 7, 123, 2, 2, 216, 217,
	7, 112, 2, 2, 217, 218, 7, 101, 2, 2, 218, 8, 3, 2, 2, 2, 219, 220, 7,
	99, 2, 2, 220, 221, 7, 118, 2, 2, 221, 222, 7, 113, 2, 2, 222, 223, 7,
	111, 2, 2, 223, 224, 7, 107, 2, 2, 224, 225, 7, 101, 2, 2, 225, 10, 3,
	2, 2, 2, 226, 227, 7, 100, 2, 2, 227, 228, 7, 113, 2, 2, 228, 229, 7, 113,
	2, 2, 229, 230, 7, 110, 2, 2, 230, 231, 7, 103, 2, 2, 231, 232, 7, 99,
	2, 2, 232, 233, 7, 112, 2, 2, 233, 12, 3, 2, 2, 2, 234, 235, 7, 100, 2,
	2, 235, 236, 7, 116, 2, 2, 236, 237, 7, 103, 2, 2, 237, 238, 7, 99, 2,
	2, 238, 239, 7, 109, 2, 2, 239, 14, 3, 2, 2, 2, 240, 241, 7, 101, 2, 2,
	241, 242, 7, 99, 2, 2, 242, 243, 7, 110, 2, 2, 243, 244, 7, 110, 2, 2,
	244, 16, 3, 2, 2, 2, 245, 246, 7, 101, 2, 2, 246, 247, 7, 99, 2, 2, 247,
	248, 7, 117, 2, 2, 248, 249, 7, 103, 2, 2, 249, 18, 3, 2, 2, 2, 250, 251,
	7, 101, 2, 2, 251, 252, 7, 106, 2, 2, 252, 253, 7, 99, 2, 2, 253, 254,
	7, 116, 2, 2, 254, 20, 3, 2, 2, 2, 255, 256, 7, 101, 2, 2, 256, 257, 7,
	113, 2, 2, 257, 258, 7, 112, 2, 2, 258, 259, 7, 104, 2, 2, 259, 260, 7,
	107, 2, 2, 260, 261, 7, 105, 2, 2, 261, 262, 7, 119, 2, 2, 262, 263, 7,
	116, 2, 2, 263, 264, 7, 99, 2, 2, 264, 265, 7, 118, 2, 2, 265, 266, 7,
	107, 2, 2, 266, 267, 7, 113, 2, 2, 267, 268, 7, 112, 2, 2, 268, 22, 3,
	2, 2, 2, 269, 270, 7, 101, 2, 2, 270, 271, 7, 113, 2, 2, 271, 272, 7, 111,
	2, 2, 272, 273, 7, 111, 2, 2, 273, 274, 7, 99, 2, 2, 274, 275, 7, 112,
	2, 2, 275, 276, 7, 102, 2, 2, 276, 24, 3, 2, 2, 2, 277, 278, 7, 101, 2,
	2, 278, 279, 7, 113, 2, 2, 279, 280, 7, 111, 2, 2, 280, 281, 7, 114, 2,
	2, 281, 282, 7, 113, 2, 2, 282, 283, 7, 112, 2, 2, 283, 284, 7, 103, 2,
	2, 284, 285, 7, 112, 2, 2, 285, 286, 7, 118, 2, 2, 286, 26, 3, 2, 2, 2,
	287, 288, 7, 101, 2, 2, 288, 289, 7, 113, 2, 2, 289, 290, 7, 111, 2, 2,
	290, 291, 7, 114, 2, 2, 291, 292, 7, 113, 2, 2, 292, 293, 7, 112, 2, 2,
	293, 294, 7, 103, 2, 2, 294, 295, 7, 112, 2, 2, 295, 296, 7, 118, 2, 2,
	296, 297, 7, 117, 2, 2, 297, 28, 3, 2, 2, 2, 298, 299, 7, 101, 2, 2, 299,
	300, 7, 113, 2, 2, 300, 301, 7, 112, 2, 2, 301, 302, 7, 118, 2, 2, 302,
	303, 7, 107, 2, 2, 303, 304, 7, 112, 2, 2, 304, 305, 7, 119, 2, 2, 305,
	306, 7, 103, 2, 2, 306, 30, 3, 2, 2, 2, 307, 308, 7, 102, 2, 2, 308, 309,
	7, 113, 2, 2, 309, 32, 3, 2, 2, 2, 310, 311, 7, 102, 2, 2, 311, 312, 7,
	113, 2, 2, 312, 313, 7, 119, 2, 2, 313, 314, 7, 100, 2, 2, 314, 315, 7,
	110, 2, 2, 315, 316, 7, 103, 2, 2, 316, 34, 3, 2, 2, 2, 317, 318, 7, 102,
	2, 2, 318, 319, 7, 103, 2, 2, 319, 320, 7, 104, 2, 2, 320, 321, 7, 107,
	2, 2, 321, 322, 7, 112, 2, 2, 322, 323, 7, 103, 2, 2, 323, 324, 7, 102,
	2, 2, 324, 36, 3, 2, 2, 2, 325, 326, 7, 102, 2, 2, 326, 327, 7, 103, 2,
	2, 327, 328, 7, 104, 2, 2, 328, 329, 7, 107, 2, 2, 329, 330, 7, 112, 2,
	2, 330, 331, 7, 103, 2, 2, 331, 38, 3, 2, 2, 2, 332, 333, 7, 102, 2, 2,
	333, 334, 7, 103, 2, 2, 334, 335, 7, 104, 2, 2, 335, 336, 7, 99, 2, 2,
	336, 337, 7, 119, 2, 2, 337, 338, 7, 110, 2, 2, 338, 339, 7, 118, 2, 2,
	339, 40, 3, 2, 2, 2, 340, 341, 7, 103, 2, 2, 341, 342, 7, 110, 2, 2, 342,
	343, 7, 117, 2, 2, 343, 344, 7, 103, 2, 2, 344, 42, 3, 2, 2, 2, 345, 346,
	7, 103, 2, 2, 346, 347, 7, 112, 2, 2, 347, 348, 7, 119, 2, 2, 348, 349,
	7, 111, 2, 2, 349, 44, 3, 2, 2, 2, 350, 351, 7, 103, 2, 2, 351, 352, 7,
	120, 2, 2, 352, 353, 7, 103, 2, 2, 353, 354, 7, 112, 2, 2, 354, 355, 7,
	118, 2, 2, 355, 46, 3, 2, 2, 2, 356, 357, 7, 103, 2, 2, 357, 358, 7, 122,
	2, 2, 358, 359, 7, 118, 2, 2, 359, 360, 7, 103, 2, 2, 360, 361, 7, 112,
	2, 2, 361, 362, 7, 102, 2, 2, 362, 363, 7, 117, 2, 2, 363, 48, 3, 2, 2,
	2, 364, 365, 7, 103, 2, 2, 365, 366, 7, 110, 2, 2, 366, 367, 7, 107, 2,
	2, 367, 368, 7, 104, 2, 2, 368, 50, 3, 2, 2, 2, 369, 370, 7, 103, 2, 2,
	370, 371, 7, 112, 2, 2, 371, 372, 7, 102, 2, 2, 372, 373, 7, 107, 2, 2,
	373, 374, 7, 104, 2, 2, 374, 52, 3, 2, 2, 2, 375, 376, 7, 103, 2, 2, 376,
	377, 7, 116, 2, 2, 377, 378, 7, 116, 2, 2, 378, 379, 7, 113, 2, 2, 379,
	380, 7, 116, 2, 2, 380, 54, 3, 2, 2, 2, 381, 382, 7, 104, 2, 2, 382, 383,
	7, 99, 2, 2, 383, 384, 7, 110, 2, 2, 384, 385, 7, 117, 2, 2, 385, 386,
	7, 103, 2, 2, 386, 56, 3, 2, 2, 2, 387, 388, 7, 104, 2, 2, 388, 389, 7,
	107, 2, 2, 389, 390, 7, 112, 2, 2, 390, 391, 7, 99, 2, 2, 391, 392, 7,
	110, 2, 2, 392, 58, 3, 2, 2, 2, 393, 394, 7, 104, 2, 2, 394, 395, 7, 113,
	2, 2, 395, 396, 7, 116, 2, 2, 396, 60, 3, 2, 2, 2, 397, 398, 7, 105, 2,
	2, 398, 399, 7, 103, 2, 2, 399, 400, 7, 112, 2, 2, 400, 401, 7, 103, 2,
	2, 401, 402, 7, 116, 2, 2, 402, 403, 7, 107, 2, 2, 403, 404, 7, 101, 2,
	2, 404, 62, 3, 2, 2, 2, 405, 406, 7, 107, 2, 2, 406, 407, 7, 104, 2, 2,
	407, 64, 3, 2, 2, 2, 408, 409, 7, 107, 2, 2, 409, 410, 7, 111, 2, 2, 410,
	411, 7, 114, 2, 2, 411, 412, 7, 110, 2, 2, 412, 413, 7, 103, 2, 2, 413,
	414, 7, 111, 2, 2, 414, 415, 7, 103, 2, 2, 415, 416, 7, 112, 2, 2, 416,
	417, 7, 118, 2, 2, 417, 418, 7, 99, 2, 2, 418, 419, 7, 118, 2, 2, 419,
	420, 7, 107, 2, 2, 420, 421, 7, 113, 2, 2, 421, 422, 7, 112, 2, 2, 422,
	66, 3, 2, 2, 2, 423, 424, 7, 107, 2, 2, 424, 425, 7, 112, 2, 2, 425, 426,
	7, 101, 2, 2, 426, 427, 7, 110, 2, 2, 427, 428, 7, 119, 2, 2, 428, 429,
	7, 102, 2, 2, 429, 430, 7, 103, 2, 2, 430, 68, 3, 2, 2, 2, 431, 432, 7,
	107, 2, 2, 432, 433, 7, 112, 2, 2, 433, 434, 7, 101, 2, 2, 434, 435, 7,
	110, 2, 2, 435, 436, 7, 119, 2, 2, 436, 437, 7, 102, 2, 2, 437, 438, 7,
	103, 2, 2, 438, 439, 7, 117, 2, 2, 439, 70, 3, 2, 2, 2, 440, 441, 7, 107,
	2, 2, 441, 442, 7, 112, 2, 2, 442, 443, 7, 118, 2, 2, 443, 444, 7, 103,
	2, 2, 444, 445, 7, 116, 2, 2, 445, 446, 7, 104, 2, 2, 446, 447, 7, 99,
	2, 2, 447, 448, 7, 101, 2, 2, 448, 449, 7, 103, 2, 2, 449, 72, 3, 2, 2,
	2, 450, 451, 7, 110, 2, 2, 451, 452, 7, 113, 2, 2, 452, 453, 7, 105, 2,
	2, 453, 74, 3, 2, 2, 2, 454, 455, 7, 110, 2, 2, 455, 456, 7, 113, 2, 2,
	456, 457, 7, 112, 2, 2, 457, 458, 7, 105, 2, 2, 458, 76, 3, 2, 2, 2, 459,
	460, 7, 111, 2, 2, 460, 461, 7, 113, 2, 2, 461, 462, 7, 102, 2, 2, 462,
	463, 7, 119, 2, 2, 463, 464, 7, 110, 2, 2, 464, 465, 7, 103, 2, 2, 465,
	78, 3, 2, 2, 2, 466, 467, 7, 112, 2, 2, 467, 468, 7, 103, 2, 2, 468, 469,
	7, 121, 2, 2, 469, 80, 3, 2, 2, 2, 470, 471, 7, 114, 2, 2, 471, 472, 7,
	113, 2, 2, 472, 473, 7, 117, 2, 2, 473, 474, 7, 118, 2, 2, 474, 82, 3,
	2, 2, 2, 475, 476, 7, 114, 2, 2, 476, 477, 7, 116, 2, 2, 477, 478, 7, 113,
	2, 2, 478, 479, 7, 120, 2, 2, 479, 480, 7, 107, 2, 2, 480, 481, 7, 102,
	2, 2, 481, 482, 7, 103, 2, 2, 482, 483, 7, 117, 2, 2, 483, 84, 3, 2, 2,
	2, 484, 485, 7, 116, 2, 2, 485, 486, 7, 103, 2, 2, 486, 487, 7, 118, 2,
	2, 487, 488, 7, 119, 2, 2, 488, 489, 7, 116, 2, 2, 489, 490, 7, 112, 2,
	2, 490, 86, 3, 2, 2, 2, 491, 492, 7, 117, 2, 2, 492, 493, 7, 106, 2, 2,
	493, 494, 7, 113, 2, 2, 494, 495, 7, 116, 2, 2, 495, 496, 7, 118, 2, 2,
	496, 88, 3, 2, 2, 2, 497, 498, 7, 117, 2, 2, 498, 499, 7, 107, 2, 2, 499,
	500, 7, 105, 2, 2, 500, 501, 7, 112, 2, 2, 501, 502, 7, 99, 2, 2, 502,
	503, 7, 110, 2, 2, 503, 90, 3, 2, 2, 2, 504, 505, 7, 117, 2, 2, 505, 506,
	7, 118, 2, 2, 506, 507, 7, 99, 2, 2, 507, 508, 7, 118, 2, 2, 508, 509,
	7, 107, 2, 2, 509, 510, 7, 101, 2, 2, 510, 92, 3, 2, 2, 2, 511, 512, 7,
	117, 2, 2, 512, 513, 7, 121, 2, 2, 513, 514, 7, 107, 2, 2, 514, 515, 7,
	118, 2, 2, 515, 516, 7, 101, 2, 2, 516, 517, 7, 106, 2, 2, 517, 94, 3,
	2, 2, 2, 518, 519, 7, 118, 2, 2, 519, 520, 7, 99, 2, 2, 520, 521, 7, 117,
	2, 2, 521, 522, 7, 109, 2, 2, 522, 96, 3, 2, 2, 2, 523, 524, 7, 118, 2,
	2, 524, 525, 7, 116, 2, 2, 525, 526, 7, 119, 2, 2, 526, 527, 7, 103, 2,
	2, 527, 98, 3, 2, 2, 2, 528, 529, 7, 119, 2, 2, 529, 530, 7, 117, 2, 2,
	530, 531, 7, 103, 2, 2, 531, 532, 7, 117, 2, 2, 532, 100, 3, 2, 2, 2, 533,
	534, 7, 120, 2, 2, 534, 535, 7, 113, 2, 2, 535, 536, 7, 107, 2, 2, 536,
	537, 7, 102, 2, 2, 537, 102, 3, 2, 2, 2, 538, 539, 7, 121, 2, 2, 539, 540,
	7, 106, 2, 2, 540, 541, 7, 107, 2, 2, 541, 542, 7, 110, 2, 2, 542, 543,
	7, 103, 2, 2, 543, 104, 3, 2, 2, 2, 544, 545, 7, 118, 2, 2, 545, 546, 7,
	123, 2, 2, 546, 547, 7, 114, 2, 2, 547, 548, 7, 103, 2, 2, 548, 549, 7,
	102, 2, 2, 549, 550, 7, 103, 2, 2, 550, 551, 7, 104, 2, 2, 551, 106, 3,
	2, 2, 2, 552, 553, 7, 126, 2, 2, 553, 554, 7, 126, 2, 2, 554, 108, 3, 2,
	2, 2, 555, 556, 7, 40, 2, 2, 556, 557, 7, 40, 2, 2, 557, 110, 3, 2, 2,
	2, 558, 559, 7, 63, 2, 2, 559, 560, 7, 63, 2, 2, 560, 112, 3, 2, 2, 2,
	561, 562, 7, 35, 2, 2, 562, 563, 7, 63, 2, 2, 563, 114, 3, 2, 2, 2, 564,
	565, 7, 64, 2, 2, 565, 116, 3, 2, 2, 2, 566, 567, 7, 62, 2, 2, 567, 118,
	3, 2, 2, 2, 568, 569, 7, 64, 2, 2, 569, 570, 7, 63, 2, 2, 570, 120, 3,
	2, 2, 2, 571, 572, 7, 62, 2, 2, 572, 573, 7, 63, 2, 2, 573, 122, 3, 2,
	2, 2, 574, 575, 7, 45, 2, 2, 575, 124, 3, 2, 2, 2, 576, 577, 7, 47, 2,
	2, 577, 126, 3, 2, 2, 2, 578, 579, 7, 44, 2, 2, 579, 128, 3, 2, 2, 2, 580,
	581, 7, 49, 2, 2, 581, 130, 3, 2, 2, 2, 582, 583, 7, 39, 2, 2, 583, 132,
	3, 2, 2, 2, 584, 585, 7, 96, 2, 2, 585, 134, 3, 2, 2, 2, 586, 587, 7, 35,
	2, 2, 587, 136, 3, 2, 2, 2, 588, 589, 7, 63, 2, 2, 589, 138, 3, 2, 2, 2,
	590, 591, 7, 128, 2, 2, 591, 140, 3, 2, 2, 2, 592, 593, 7, 65, 2, 2, 593,
	142, 3, 2, 2, 2, 594, 595, 7, 60, 2, 2, 595, 144, 3, 2, 2, 2, 596, 597,
	7, 45, 2, 2, 597, 598, 7, 45, 2, 2, 598, 146, 3, 2, 2, 2, 599, 600, 7,
	47, 2, 2, 600, 601, 7, 47, 2, 2, 601, 148, 3, 2, 2, 2, 602, 603, 7, 40,
	2, 2, 603, 150, 3, 2, 2, 2, 604, 605, 7, 126, 2, 2, 605, 152, 3, 2, 2,
	2, 606, 607, 7, 37, 2, 2, 607, 154, 3, 2, 2, 2, 608, 609, 7, 61, 2, 2,
	609, 156, 3, 2, 2, 2, 610, 611, 7, 93, 2, 2, 611, 158, 3, 2, 2, 2, 612,
	613, 7, 95, 2, 2, 613, 160, 3, 2, 2, 2, 614, 615, 7, 42, 2, 2, 615, 162,
	3, 2, 2, 2, 616, 617, 7, 43, 2, 2, 617, 164, 3, 2, 2, 2, 618, 619, 7, 125,
	2, 2, 619, 166, 3, 2, 2, 2, 620, 621, 7, 127, 2, 2, 621, 168, 3, 2, 2,
	2, 622, 623, 7, 47, 2, 2, 623, 624, 7, 64, 2, 2, 624, 170, 3, 2, 2, 2,
	625, 626, 7, 62, 2, 2, 626, 627, 7, 47, 2, 2, 627, 172, 3, 2, 2, 2, 628,
	629, 7, 60, 2, 2, 629, 630, 7, 60, 2, 2, 630, 174, 3, 2, 2, 2, 631, 632,
	7, 66, 2, 2, 632, 176, 3, 2, 2, 2, 633, 634, 7, 46, 2, 2, 634, 178, 3,
	2, 2, 2, 635, 636, 7, 48, 2, 2, 636, 180, 3, 2, 2, 2, 637, 638, 7, 97,
	2, 2, 638, 182, 3, 2, 2, 2, 639, 643, 9, 2, 2, 2, 640, 642, 9, 3, 2, 2,
	641, 640, 3, 2, 2, 2, 642, 645, 3, 2, 2, 2, 643, 641, 3, 2, 2, 2, 643,
	644, 3, 2, 2, 2, 644, 184, 3, 2, 2, 2, 645, 643, 3, 2, 2, 2, 646, 648,
	9, 4, 2, 2, 647, 646, 3, 2, 2, 2, 648, 649, 3, 2, 2, 2, 649, 647, 3, 2,
	2, 2, 649, 650, 3, 2, 2, 2, 650, 186, 3, 2, 2, 2, 651, 653, 9, 4, 2, 2,
	652, 651, 3, 2, 2, 2, 653, 654, 3, 2, 2, 2, 654, 652, 3, 2, 2, 2, 654,
	655, 3, 2, 2, 2, 655, 656, 3, 2, 2, 2, 656, 660, 7, 48, 2, 2, 657, 659,
	9, 4, 2, 2, 658, 657, 3, 2, 2, 2, 659, 662, 3, 2, 2, 2, 660, 658, 3, 2,
	2, 2, 660, 661, 3, 2, 2, 2, 661, 670, 3, 2, 2, 2, 662, 660, 3, 2, 2, 2,
	663, 665, 7, 48, 2, 2, 664, 666, 9, 4, 2, 2, 665, 664, 3, 2, 2, 2, 666,
	667, 3, 2, 2, 2, 667, 665, 3, 2, 2, 2, 667, 668, 3, 2, 2, 2, 668, 670,
	3, 2, 2, 2, 669, 652, 3, 2, 2, 2, 669, 663, 3, 2, 2, 2, 670, 188, 3, 2,
	2, 2, 671, 677, 7, 36, 2, 2, 672, 676, 10, 5, 2, 2, 673, 674, 7, 36, 2,
	2, 674, 676, 7, 36, 2, 2, 675, 672, 3, 2, 2, 2, 675, 673, 3, 2, 2, 2, 676,
	679, 3, 2, 2, 2, 677, 675, 3, 2, 2, 2, 677, 678, 3, 2, 2, 2, 678, 680,
	3, 2, 2, 2, 679, 677, 3, 2, 2, 2, 680, 681, 7, 36, 2, 2, 681, 190, 3, 2,
	2, 2, 682, 683, 7, 49, 2, 2, 683, 684, 7, 44, 2, 2, 684, 688, 3, 2, 2,
	2, 685, 687, 11, 2, 2, 2, 686, 685, 3, 2, 2, 2, 687, 690, 3, 2, 2, 2, 688,
	689, 3, 2, 2, 2, 688, 686, 3, 2, 2, 2, 689, 691, 3, 2, 2, 2, 690, 688,
	3, 2, 2, 2, 691, 692, 7, 44, 2, 2, 692, 693, 7, 49, 2, 2, 693, 694, 3,
	2, 2, 2, 694, 695, 8, 96, 2, 2, 695, 192, 3, 2, 2, 2, 696, 697, 7, 49,
	2, 2, 697, 698, 7, 49, 2, 2, 698, 702, 3, 2, 2, 2, 699, 701, 10, 6, 2,
	2, 700, 699, 3, 2, 2, 2, 701, 704, 3, 2, 2, 2, 702, 700, 3, 2, 2, 2, 702,
	703, 3, 2, 2, 2, 703, 705, 3, 2, 2, 2, 704, 702, 3, 2, 2, 2, 705, 706,
	8, 97, 2, 2, 706, 194, 3, 2, 2, 2, 707, 709, 9, 7, 2, 2, 708, 707, 3, 2,
	2, 2, 709, 710, 3, 2, 2, 2, 710, 708, 3, 2, 2, 2, 710, 711, 3, 2, 2, 2,
	711, 712, 3, 2, 2, 2, 712, 713, 8, 98, 3, 2, 713, 196, 3, 2, 2, 2, 714,
	715, 11, 2, 2, 2, 715, 198, 3, 2, 2, 2, 716, 717, 7, 50, 2, 2, 717, 718,
	7, 122, 2, 2, 718, 720, 3, 2, 2, 2, 719, 721, 9, 8, 2, 2, 720, 719, 3,
	2, 2, 2, 721, 722, 3, 2, 2, 2, 722, 720, 3, 2, 2, 2, 722, 723, 3, 2, 2,
	2, 723, 200, 3, 2, 2, 2, 15, 2, 643, 649, 654, 660, 667, 669, 675, 677,
	688, 702, 710, 722, 4, 2, 3, 2, 2, 4, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN", "WHITESPACE",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'abstract'", "'as'", "'async'", "'atomic'", "'boolean'", "'break'",
	"'call'", "'case'", "'char'", "'configuration'", "'command'", "'component'",
	"'components'", "'continue'", "'do'", "'double'", "'defined'", "'define'",
	"'default'", "'else'", "'enum'", "'event'", "'extends'", "'elif'", "'endif'",
	"'error'", "'false'", "'final'", "'for'", "'generic'", "'if'", "'implementation'",
	"'include'", "'includes'", "'interface'", "'log'", "'long'", "'module'",
	"'new'", "'post'", "'provides'", "'return'", "'short'", "'signal'", "'static'",
	"'switch'", "'task'", "'true'", "'uses'", "'void'", "'while'", "'typedef'",
	"'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'+'", "'-'",
	"'*'", "'/'", "'%'", "'^'", "'!'", "'='", "'~'", "'?'", "':'", "'++'",
	"'--'", "'&'", "'|'", "'#'", "';'", "'['", "']'", "'('", "')'", "'{'",
	"'}'", "'->'", "'<-'", "'::'", "'@'", "','", "'.'", "'_'",
}

var lexerSymbolicNames = []string{
	"", "ABSTRACT", "AS", "ASYNC", "ATOMIC", "BOOLEAN", "BREAK", "CALL", "CASE",
	"CHAR", "COFIGURATION", "COMMAND", "COMPONENT", "COMPONENTS", "CONTINUE",
	"DO", "DOUBLE", "DEFINED", "DEFINE", "DEFAULT", "ELSE", "ENUM", "EVENT",
	"EXTENDS", "ELIF", "ENDIF", "ERROR", "FALSE", "FINAL", "FOR", "GENERIC",
	"IF", "IMPLEMENTATION", "INCLUDE", "INCLUDES", "INTERFACE", "LOG", "LONG",
	"MODULE", "NEW", "POST", "PROVIDES", "RETURN", "SHORT", "SIGNAL", "STATIC",
	"SWITCH", "TASK", "TRUE", "USES", "VOID", "WHILE", "TYPEDEF", "OR", "AND",
	"EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS", "MULT", "DIV",
	"MOD", "POW", "NOT", "ASSIGN", "TILDE", "QUESTION", "COLON", "INC", "DEC",
	"BITAND", "BITOR", "HASHTAG", "SCOL", "OBRACK", "CBRACK", "OPAR", "CPAR",
	"OBRACE", "CBRACE", "FORWARDARROW", "BACKARROW", "COLONCOLON", "AT", "COMMA",
	"DOT", "UNDERSCORE", "ID", "INT", "FLOAT", "STRING", "COMMENT", "LINE_COMMENT",
	"SPACE", "OTHER", "HEX",
}

var lexerRuleNames = []string{
	"ABSTRACT", "AS", "ASYNC", "ATOMIC", "BOOLEAN", "BREAK", "CALL", "CASE",
	"CHAR", "COFIGURATION", "COMMAND", "COMPONENT", "COMPONENTS", "CONTINUE",
	"DO", "DOUBLE", "DEFINED", "DEFINE", "DEFAULT", "ELSE", "ENUM", "EVENT",
	"EXTENDS", "ELIF", "ENDIF", "ERROR", "FALSE", "FINAL", "FOR", "GENERIC",
	"IF", "IMPLEMENTATION", "INCLUDE", "INCLUDES", "INTERFACE", "LOG", "LONG",
	"MODULE", "NEW", "POST", "PROVIDES", "RETURN", "SHORT", "SIGNAL", "STATIC",
	"SWITCH", "TASK", "TRUE", "USES", "VOID", "WHILE", "TYPEDEF", "OR", "AND",
	"EQ", "NEQ", "GT", "LT", "GTEQ", "LTEQ", "PLUS", "MINUS", "MULT", "DIV",
	"MOD", "POW", "NOT", "ASSIGN", "TILDE", "QUESTION", "COLON", "INC", "DEC",
	"BITAND", "BITOR", "HASHTAG", "SCOL", "OBRACK", "CBRACK", "OPAR", "CPAR",
	"OBRACE", "CBRACE", "FORWARDARROW", "BACKARROW", "COLONCOLON", "AT", "COMMA",
	"DOT", "UNDERSCORE", "ID", "INT", "FLOAT", "STRING", "COMMENT", "LINE_COMMENT",
	"SPACE", "OTHER", "HEX",
}

type TinyosLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewTinyosLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *TinyosLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewTinyosLexer(input antlr.CharStream) *TinyosLexer {
	l := new(TinyosLexer)
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
	l.GrammarFileName = "TinyosLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TinyosLexer tokens.
const (
	TinyosLexerABSTRACT       = 1
	TinyosLexerAS             = 2
	TinyosLexerASYNC          = 3
	TinyosLexerATOMIC         = 4
	TinyosLexerBOOLEAN        = 5
	TinyosLexerBREAK          = 6
	TinyosLexerCALL           = 7
	TinyosLexerCASE           = 8
	TinyosLexerCHAR           = 9
	TinyosLexerCOFIGURATION   = 10
	TinyosLexerCOMMAND        = 11
	TinyosLexerCOMPONENT      = 12
	TinyosLexerCOMPONENTS     = 13
	TinyosLexerCONTINUE       = 14
	TinyosLexerDO             = 15
	TinyosLexerDOUBLE         = 16
	TinyosLexerDEFINED        = 17
	TinyosLexerDEFINE         = 18
	TinyosLexerDEFAULT        = 19
	TinyosLexerELSE           = 20
	TinyosLexerENUM           = 21
	TinyosLexerEVENT          = 22
	TinyosLexerEXTENDS        = 23
	TinyosLexerELIF           = 24
	TinyosLexerENDIF          = 25
	TinyosLexerERROR          = 26
	TinyosLexerFALSE          = 27
	TinyosLexerFINAL          = 28
	TinyosLexerFOR            = 29
	TinyosLexerGENERIC        = 30
	TinyosLexerIF             = 31
	TinyosLexerIMPLEMENTATION = 32
	TinyosLexerINCLUDE        = 33
	TinyosLexerINCLUDES       = 34
	TinyosLexerINTERFACE      = 35
	TinyosLexerLOG            = 36
	TinyosLexerLONG           = 37
	TinyosLexerMODULE         = 38
	TinyosLexerNEW            = 39
	TinyosLexerPOST           = 40
	TinyosLexerPROVIDES       = 41
	TinyosLexerRETURN         = 42
	TinyosLexerSHORT          = 43
	TinyosLexerSIGNAL         = 44
	TinyosLexerSTATIC         = 45
	TinyosLexerSWITCH         = 46
	TinyosLexerTASK           = 47
	TinyosLexerTRUE           = 48
	TinyosLexerUSES           = 49
	TinyosLexerVOID           = 50
	TinyosLexerWHILE          = 51
	TinyosLexerTYPEDEF        = 52
	TinyosLexerOR             = 53
	TinyosLexerAND            = 54
	TinyosLexerEQ             = 55
	TinyosLexerNEQ            = 56
	TinyosLexerGT             = 57
	TinyosLexerLT             = 58
	TinyosLexerGTEQ           = 59
	TinyosLexerLTEQ           = 60
	TinyosLexerPLUS           = 61
	TinyosLexerMINUS          = 62
	TinyosLexerMULT           = 63
	TinyosLexerDIV            = 64
	TinyosLexerMOD            = 65
	TinyosLexerPOW            = 66
	TinyosLexerNOT            = 67
	TinyosLexerASSIGN         = 68
	TinyosLexerTILDE          = 69
	TinyosLexerQUESTION       = 70
	TinyosLexerCOLON          = 71
	TinyosLexerINC            = 72
	TinyosLexerDEC            = 73
	TinyosLexerBITAND         = 74
	TinyosLexerBITOR          = 75
	TinyosLexerHASHTAG        = 76
	TinyosLexerSCOL           = 77
	TinyosLexerOBRACK         = 78
	TinyosLexerCBRACK         = 79
	TinyosLexerOPAR           = 80
	TinyosLexerCPAR           = 81
	TinyosLexerOBRACE         = 82
	TinyosLexerCBRACE         = 83
	TinyosLexerFORWARDARROW   = 84
	TinyosLexerBACKARROW      = 85
	TinyosLexerCOLONCOLON     = 86
	TinyosLexerAT             = 87
	TinyosLexerCOMMA          = 88
	TinyosLexerDOT            = 89
	TinyosLexerUNDERSCORE     = 90
	TinyosLexerID             = 91
	TinyosLexerINT            = 92
	TinyosLexerFLOAT          = 93
	TinyosLexerSTRING         = 94
	TinyosLexerCOMMENT        = 95
	TinyosLexerLINE_COMMENT   = 96
	TinyosLexerSPACE          = 97
	TinyosLexerOTHER          = 98
	TinyosLexerHEX            = 99
)

// TinyosLexerWHITESPACE is the TinyosLexer channel.
const TinyosLexerWHITESPACE = 2
