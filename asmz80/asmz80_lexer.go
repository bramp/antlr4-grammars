// Code generated from asmZ80.g4 by ANTLR 4.9.3. DO NOT EDIT.

package asmz80

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 556,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 139, 10, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 179, 10, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 5, 13, 458, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37,
	3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 7, 40, 514, 10, 40, 12, 40, 14,
	40, 517, 11, 40, 3, 41, 5, 41, 520, 10, 41, 3, 41, 6, 41, 523, 10, 41,
	13, 41, 14, 41, 524, 3, 41, 5, 41, 528, 10, 41, 3, 42, 3, 42, 7, 42, 532,
	10, 42, 12, 42, 14, 42, 535, 11, 42, 3, 42, 3, 42, 3, 43, 3, 43, 7, 43,
	541, 10, 43, 12, 43, 14, 43, 544, 11, 43, 3, 43, 3, 43, 3, 44, 6, 44, 549,
	10, 44, 13, 44, 14, 44, 550, 3, 45, 3, 45, 3, 45, 3, 45, 2, 2, 46, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 2, 29, 2, 31, 2, 33, 2, 35, 2, 37, 2, 39, 2, 41, 2, 43, 2,
	45, 2, 47, 2, 49, 2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2, 63, 2, 65,
	2, 67, 2, 69, 2, 71, 2, 73, 2, 75, 2, 77, 2, 79, 15, 81, 16, 83, 17, 85,
	18, 87, 19, 89, 20, 3, 2, 35, 6, 2, 67, 72, 74, 75, 78, 78, 84, 84, 4,
	2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2,
	70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2,
	73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2,
	76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2,
	79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2,
	82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2,
	85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2,
	88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2,
	91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 4, 2, 67, 92, 99, 124, 7, 2,
	36, 36, 48, 48, 50, 59, 67, 92, 99, 124, 5, 2, 50, 59, 67, 72, 99, 104,
	4, 2, 12, 12, 15, 15, 3, 2, 41, 41, 4, 2, 11, 11, 34, 34, 2, 623, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2,
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2,
	2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2,
	2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 3, 91, 3, 2, 2, 2, 5, 93, 3, 2,
	2, 2, 7, 95, 3, 2, 2, 2, 9, 97, 3, 2, 2, 2, 11, 99, 3, 2, 2, 2, 13, 101,
	3, 2, 2, 2, 15, 103, 3, 2, 2, 2, 17, 105, 3, 2, 2, 2, 19, 107, 3, 2, 2,
	2, 21, 138, 3, 2, 2, 2, 23, 178, 3, 2, 2, 2, 25, 457, 3, 2, 2, 2, 27, 459,
	3, 2, 2, 2, 29, 461, 3, 2, 2, 2, 31, 463, 3, 2, 2, 2, 33, 465, 3, 2, 2,
	2, 35, 467, 3, 2, 2, 2, 37, 469, 3, 2, 2, 2, 39, 471, 3, 2, 2, 2, 41, 473,
	3, 2, 2, 2, 43, 475, 3, 2, 2, 2, 45, 477, 3, 2, 2, 2, 47, 479, 3, 2, 2,
	2, 49, 481, 3, 2, 2, 2, 51, 483, 3, 2, 2, 2, 53, 485, 3, 2, 2, 2, 55, 487,
	3, 2, 2, 2, 57, 489, 3, 2, 2, 2, 59, 491, 3, 2, 2, 2, 61, 493, 3, 2, 2,
	2, 63, 495, 3, 2, 2, 2, 65, 497, 3, 2, 2, 2, 67, 499, 3, 2, 2, 2, 69, 501,
	3, 2, 2, 2, 71, 503, 3, 2, 2, 2, 73, 505, 3, 2, 2, 2, 75, 507, 3, 2, 2,
	2, 77, 509, 3, 2, 2, 2, 79, 511, 3, 2, 2, 2, 81, 519, 3, 2, 2, 2, 83, 529,
	3, 2, 2, 2, 85, 538, 3, 2, 2, 2, 87, 548, 3, 2, 2, 2, 89, 552, 3, 2, 2,
	2, 91, 92, 7, 60, 2, 2, 92, 4, 3, 2, 2, 2, 93, 94, 7, 46, 2, 2, 94, 6,
	3, 2, 2, 2, 95, 96, 7, 45, 2, 2, 96, 8, 3, 2, 2, 2, 97, 98, 7, 47, 2, 2,
	98, 10, 3, 2, 2, 2, 99, 100, 7, 44, 2, 2, 100, 12, 3, 2, 2, 2, 101, 102,
	7, 49, 2, 2, 102, 14, 3, 2, 2, 2, 103, 104, 7, 42, 2, 2, 104, 16, 3, 2,
	2, 2, 105, 106, 7, 43, 2, 2, 106, 18, 3, 2, 2, 2, 107, 108, 7, 38, 2, 2,
	108, 20, 3, 2, 2, 2, 109, 139, 9, 2, 2, 2, 110, 111, 7, 75, 2, 2, 111,
	112, 7, 90, 2, 2, 112, 139, 7, 74, 2, 2, 113, 114, 7, 75, 2, 2, 114, 115,
	7, 90, 2, 2, 115, 139, 7, 78, 2, 2, 116, 117, 7, 75, 2, 2, 117, 118, 7,
	91, 2, 2, 118, 139, 7, 74, 2, 2, 119, 120, 7, 75, 2, 2, 120, 121, 7, 91,
	2, 2, 121, 139, 7, 78, 2, 2, 122, 123, 7, 67, 2, 2, 123, 139, 7, 72, 2,
	2, 124, 125, 7, 68, 2, 2, 125, 139, 7, 69, 2, 2, 126, 127, 7, 70, 2, 2,
	127, 139, 7, 71, 2, 2, 128, 129, 7, 74, 2, 2, 129, 139, 7, 78, 2, 2, 130,
	131, 7, 82, 2, 2, 131, 139, 7, 69, 2, 2, 132, 133, 7, 85, 2, 2, 133, 139,
	7, 82, 2, 2, 134, 135, 7, 75, 2, 2, 135, 139, 7, 90, 2, 2, 136, 137, 7,
	75, 2, 2, 137, 139, 7, 91, 2, 2, 138, 109, 3, 2, 2, 2, 138, 110, 3, 2,
	2, 2, 138, 113, 3, 2, 2, 2, 138, 116, 3, 2, 2, 2, 138, 119, 3, 2, 2, 2,
	138, 122, 3, 2, 2, 2, 138, 124, 3, 2, 2, 2, 138, 126, 3, 2, 2, 2, 138,
	128, 3, 2, 2, 2, 138, 130, 3, 2, 2, 2, 138, 132, 3, 2, 2, 2, 138, 134,
	3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 22, 3, 2, 2, 2, 140, 141, 5, 55,
	28, 2, 141, 142, 5, 61, 31, 2, 142, 143, 5, 39, 20, 2, 143, 179, 3, 2,
	2, 2, 144, 145, 5, 35, 18, 2, 145, 146, 5, 53, 27, 2, 146, 147, 5, 33,
	17, 2, 147, 179, 3, 2, 2, 2, 148, 149, 5, 35, 18, 2, 149, 150, 5, 59, 30,
	2, 150, 151, 5, 67, 34, 2, 151, 179, 3, 2, 2, 2, 152, 153, 5, 33, 17, 2,
	153, 154, 5, 35, 18, 2, 154, 155, 5, 37, 19, 2, 155, 156, 5, 29, 15, 2,
	156, 179, 3, 2, 2, 2, 157, 158, 5, 33, 17, 2, 158, 159, 5, 35, 18, 2, 159,
	160, 5, 37, 19, 2, 160, 161, 5, 71, 36, 2, 161, 179, 3, 2, 2, 2, 162, 163,
	5, 33, 17, 2, 163, 164, 5, 63, 32, 2, 164, 179, 3, 2, 2, 2, 165, 166, 5,
	43, 22, 2, 166, 167, 5, 37, 19, 2, 167, 179, 3, 2, 2, 2, 168, 169, 5, 35,
	18, 2, 169, 170, 5, 53, 27, 2, 170, 171, 5, 33, 17, 2, 171, 172, 5, 43,
	22, 2, 172, 173, 5, 37, 19, 2, 173, 179, 3, 2, 2, 2, 174, 175, 5, 63, 32,
	2, 175, 176, 5, 35, 18, 2, 176, 177, 5, 65, 33, 2, 177, 179, 3, 2, 2, 2,
	178, 140, 3, 2, 2, 2, 178, 144, 3, 2, 2, 2, 178, 148, 3, 2, 2, 2, 178,
	152, 3, 2, 2, 2, 178, 157, 3, 2, 2, 2, 178, 162, 3, 2, 2, 2, 178, 165,
	3, 2, 2, 2, 178, 168, 3, 2, 2, 2, 178, 174, 3, 2, 2, 2, 179, 24, 3, 2,
	2, 2, 180, 181, 5, 27, 14, 2, 181, 182, 5, 33, 17, 2, 182, 183, 5, 31,
	16, 2, 183, 458, 3, 2, 2, 2, 184, 185, 5, 27, 14, 2, 185, 186, 5, 33, 17,
	2, 186, 187, 5, 33, 17, 2, 187, 458, 3, 2, 2, 2, 188, 189, 5, 27, 14, 2,
	189, 190, 5, 53, 27, 2, 190, 191, 5, 33, 17, 2, 191, 458, 3, 2, 2, 2, 192,
	193, 5, 29, 15, 2, 193, 194, 5, 43, 22, 2, 194, 195, 5, 65, 33, 2, 195,
	458, 3, 2, 2, 2, 196, 197, 5, 31, 16, 2, 197, 198, 5, 27, 14, 2, 198, 199,
	5, 49, 25, 2, 199, 200, 5, 49, 25, 2, 200, 458, 3, 2, 2, 2, 201, 202, 5,
	31, 16, 2, 202, 203, 5, 31, 16, 2, 203, 204, 5, 37, 19, 2, 204, 458, 3,
	2, 2, 2, 205, 206, 5, 31, 16, 2, 206, 207, 5, 57, 29, 2, 207, 458, 3, 2,
	2, 2, 208, 209, 5, 31, 16, 2, 209, 210, 5, 57, 29, 2, 210, 211, 5, 33,
	17, 2, 211, 458, 3, 2, 2, 2, 212, 213, 5, 31, 16, 2, 213, 214, 5, 57, 29,
	2, 214, 215, 5, 33, 17, 2, 215, 216, 5, 61, 31, 2, 216, 458, 3, 2, 2, 2,
	217, 218, 5, 31, 16, 2, 218, 219, 5, 57, 29, 2, 219, 220, 5, 43, 22, 2,
	220, 458, 3, 2, 2, 2, 221, 222, 5, 31, 16, 2, 222, 223, 5, 57, 29, 2, 223,
	224, 5, 43, 22, 2, 224, 225, 5, 61, 31, 2, 225, 458, 3, 2, 2, 2, 226, 227,
	5, 31, 16, 2, 227, 228, 5, 57, 29, 2, 228, 229, 5, 49, 25, 2, 229, 458,
	3, 2, 2, 2, 230, 231, 5, 33, 17, 2, 231, 232, 5, 27, 14, 2, 232, 233, 5,
	27, 14, 2, 233, 458, 3, 2, 2, 2, 234, 235, 5, 33, 17, 2, 235, 236, 5, 35,
	18, 2, 236, 237, 5, 31, 16, 2, 237, 458, 3, 2, 2, 2, 238, 239, 5, 33, 17,
	2, 239, 240, 5, 43, 22, 2, 240, 458, 3, 2, 2, 2, 241, 242, 5, 33, 17, 2,
	242, 243, 5, 45, 23, 2, 243, 244, 5, 53, 27, 2, 244, 245, 5, 77, 39, 2,
	245, 458, 3, 2, 2, 2, 246, 247, 5, 35, 18, 2, 247, 248, 5, 43, 22, 2, 248,
	458, 3, 2, 2, 2, 249, 250, 5, 35, 18, 2, 250, 251, 5, 73, 37, 2, 251, 458,
	3, 2, 2, 2, 252, 253, 5, 35, 18, 2, 253, 254, 5, 73, 37, 2, 254, 255, 5,
	73, 37, 2, 255, 458, 3, 2, 2, 2, 256, 257, 5, 43, 22, 2, 257, 258, 5, 51,
	26, 2, 258, 458, 3, 2, 2, 2, 259, 260, 5, 43, 22, 2, 260, 261, 5, 53, 27,
	2, 261, 458, 3, 2, 2, 2, 262, 263, 5, 43, 22, 2, 263, 264, 5, 53, 27, 2,
	264, 265, 5, 31, 16, 2, 265, 458, 3, 2, 2, 2, 266, 267, 5, 43, 22, 2, 267,
	268, 5, 53, 27, 2, 268, 269, 5, 33, 17, 2, 269, 458, 3, 2, 2, 2, 270, 271,
	5, 43, 22, 2, 271, 272, 5, 53, 27, 2, 272, 273, 5, 33, 17, 2, 273, 274,
	5, 61, 31, 2, 274, 458, 3, 2, 2, 2, 275, 276, 5, 43, 22, 2, 276, 277, 5,
	53, 27, 2, 277, 278, 5, 43, 22, 2, 278, 458, 3, 2, 2, 2, 279, 280, 5, 43,
	22, 2, 280, 281, 5, 53, 27, 2, 281, 282, 5, 43, 22, 2, 282, 283, 5, 61,
	31, 2, 283, 458, 3, 2, 2, 2, 284, 285, 5, 45, 23, 2, 285, 286, 5, 57, 29,
	2, 286, 458, 3, 2, 2, 2, 287, 288, 5, 45, 23, 2, 288, 289, 5, 61, 31, 2,
	289, 458, 3, 2, 2, 2, 290, 291, 5, 49, 25, 2, 291, 292, 5, 33, 17, 2, 292,
	458, 3, 2, 2, 2, 293, 294, 5, 49, 25, 2, 294, 295, 5, 33, 17, 2, 295, 296,
	5, 33, 17, 2, 296, 458, 3, 2, 2, 2, 297, 298, 5, 49, 25, 2, 298, 299, 5,
	33, 17, 2, 299, 300, 5, 33, 17, 2, 300, 301, 5, 61, 31, 2, 301, 458, 3,
	2, 2, 2, 302, 303, 5, 49, 25, 2, 303, 304, 5, 33, 17, 2, 304, 305, 5, 43,
	22, 2, 305, 458, 3, 2, 2, 2, 306, 307, 5, 49, 25, 2, 307, 308, 5, 33, 17,
	2, 308, 309, 5, 43, 22, 2, 309, 310, 5, 61, 31, 2, 310, 458, 3, 2, 2, 2,
	311, 312, 5, 53, 27, 2, 312, 313, 5, 35, 18, 2, 313, 314, 5, 39, 20, 2,
	314, 458, 3, 2, 2, 2, 315, 316, 5, 53, 27, 2, 316, 317, 5, 55, 28, 2, 317,
	318, 5, 57, 29, 2, 318, 458, 3, 2, 2, 2, 319, 320, 5, 55, 28, 2, 320, 321,
	5, 61, 31, 2, 321, 458, 3, 2, 2, 2, 322, 323, 5, 55, 28, 2, 323, 324, 5,
	65, 33, 2, 324, 325, 5, 33, 17, 2, 325, 326, 5, 61, 31, 2, 326, 458, 3,
	2, 2, 2, 327, 328, 5, 55, 28, 2, 328, 329, 5, 65, 33, 2, 329, 330, 5, 43,
	22, 2, 330, 331, 5, 61, 31, 2, 331, 458, 3, 2, 2, 2, 332, 333, 5, 55, 28,
	2, 333, 334, 5, 67, 34, 2, 334, 335, 5, 65, 33, 2, 335, 458, 3, 2, 2, 2,
	336, 337, 5, 55, 28, 2, 337, 338, 5, 67, 34, 2, 338, 339, 5, 65, 33, 2,
	339, 340, 5, 33, 17, 2, 340, 458, 3, 2, 2, 2, 341, 342, 5, 55, 28, 2, 342,
	343, 5, 67, 34, 2, 343, 344, 5, 65, 33, 2, 344, 345, 5, 43, 22, 2, 345,
	458, 3, 2, 2, 2, 346, 347, 5, 57, 29, 2, 347, 348, 5, 55, 28, 2, 348, 349,
	5, 57, 29, 2, 349, 458, 3, 2, 2, 2, 350, 351, 5, 57, 29, 2, 351, 352, 5,
	67, 34, 2, 352, 353, 5, 63, 32, 2, 353, 354, 5, 41, 21, 2, 354, 458, 3,
	2, 2, 2, 355, 356, 5, 61, 31, 2, 356, 357, 5, 35, 18, 2, 357, 358, 5, 63,
	32, 2, 358, 458, 3, 2, 2, 2, 359, 360, 5, 61, 31, 2, 360, 361, 5, 35, 18,
	2, 361, 362, 5, 65, 33, 2, 362, 458, 3, 2, 2, 2, 363, 364, 5, 61, 31, 2,
	364, 365, 5, 35, 18, 2, 365, 366, 5, 65, 33, 2, 366, 367, 5, 43, 22, 2,
	367, 458, 3, 2, 2, 2, 368, 369, 5, 61, 31, 2, 369, 370, 5, 35, 18, 2, 370,
	371, 5, 65, 33, 2, 371, 372, 5, 53, 27, 2, 372, 458, 3, 2, 2, 2, 373, 374,
	5, 61, 31, 2, 374, 375, 5, 49, 25, 2, 375, 458, 3, 2, 2, 2, 376, 377, 5,
	61, 31, 2, 377, 378, 5, 49, 25, 2, 378, 379, 5, 27, 14, 2, 379, 458, 3,
	2, 2, 2, 380, 381, 5, 61, 31, 2, 381, 382, 5, 49, 25, 2, 382, 383, 5, 31,
	16, 2, 383, 458, 3, 2, 2, 2, 384, 385, 5, 61, 31, 2, 385, 386, 5, 49, 25,
	2, 386, 387, 5, 31, 16, 2, 387, 388, 5, 27, 14, 2, 388, 458, 3, 2, 2, 2,
	389, 390, 5, 61, 31, 2, 390, 391, 5, 49, 25, 2, 391, 392, 5, 33, 17, 2,
	392, 458, 3, 2, 2, 2, 393, 394, 5, 61, 31, 2, 394, 395, 5, 61, 31, 2, 395,
	458, 3, 2, 2, 2, 396, 397, 5, 61, 31, 2, 397, 398, 5, 61, 31, 2, 398, 399,
	5, 27, 14, 2, 399, 458, 3, 2, 2, 2, 400, 401, 5, 61, 31, 2, 401, 402, 5,
	61, 31, 2, 402, 403, 5, 31, 16, 2, 403, 458, 3, 2, 2, 2, 404, 405, 5, 61,
	31, 2, 405, 406, 5, 61, 31, 2, 406, 407, 5, 31, 16, 2, 407, 408, 5, 27,
	14, 2, 408, 458, 3, 2, 2, 2, 409, 410, 5, 61, 31, 2, 410, 411, 5, 61, 31,
	2, 411, 412, 5, 33, 17, 2, 412, 458, 3, 2, 2, 2, 413, 414, 5, 61, 31, 2,
	414, 415, 5, 63, 32, 2, 415, 416, 5, 65, 33, 2, 416, 458, 3, 2, 2, 2, 417,
	418, 5, 63, 32, 2, 418, 419, 5, 29, 15, 2, 419, 420, 5, 31, 16, 2, 420,
	458, 3, 2, 2, 2, 421, 422, 5, 63, 32, 2, 422, 423, 5, 31, 16, 2, 423, 424,
	5, 37, 19, 2, 424, 458, 3, 2, 2, 2, 425, 426, 5, 63, 32, 2, 426, 427, 5,
	35, 18, 2, 427, 428, 5, 65, 33, 2, 428, 458, 3, 2, 2, 2, 429, 430, 5, 63,
	32, 2, 430, 431, 5, 49, 25, 2, 431, 432, 5, 27, 14, 2, 432, 458, 3, 2,
	2, 2, 433, 434, 5, 63, 32, 2, 434, 435, 5, 49, 25, 2, 435, 436, 5, 49,
	25, 2, 436, 458, 3, 2, 2, 2, 437, 438, 5, 63, 32, 2, 438, 439, 5, 49, 25,
	2, 439, 440, 7, 51, 2, 2, 440, 458, 3, 2, 2, 2, 441, 442, 5, 63, 32, 2,
	442, 443, 5, 61, 31, 2, 443, 444, 5, 27, 14, 2, 444, 458, 3, 2, 2, 2, 445,
	446, 5, 63, 32, 2, 446, 447, 5, 61, 31, 2, 447, 448, 5, 49, 25, 2, 448,
	458, 3, 2, 2, 2, 449, 450, 5, 63, 32, 2, 450, 451, 5, 67, 34, 2, 451, 452,
	5, 29, 15, 2, 452, 458, 3, 2, 2, 2, 453, 454, 5, 73, 37, 2, 454, 455, 5,
	55, 28, 2, 455, 456, 5, 61, 31, 2, 456, 458, 3, 2, 2, 2, 457, 180, 3, 2,
	2, 2, 457, 184, 3, 2, 2, 2, 457, 188, 3, 2, 2, 2, 457, 192, 3, 2, 2, 2,
	457, 196, 3, 2, 2, 2, 457, 201, 3, 2, 2, 2, 457, 205, 3, 2, 2, 2, 457,
	208, 3, 2, 2, 2, 457, 212, 3, 2, 2, 2, 457, 217, 3, 2, 2, 2, 457, 221,
	3, 2, 2, 2, 457, 226, 3, 2, 2, 2, 457, 230, 3, 2, 2, 2, 457, 234, 3, 2,
	2, 2, 457, 238, 3, 2, 2, 2, 457, 241, 3, 2, 2, 2, 457, 246, 3, 2, 2, 2,
	457, 249, 3, 2, 2, 2, 457, 252, 3, 2, 2, 2, 457, 256, 3, 2, 2, 2, 457,
	259, 3, 2, 2, 2, 457, 262, 3, 2, 2, 2, 457, 266, 3, 2, 2, 2, 457, 270,
	3, 2, 2, 2, 457, 275, 3, 2, 2, 2, 457, 279, 3, 2, 2, 2, 457, 284, 3, 2,
	2, 2, 457, 287, 3, 2, 2, 2, 457, 290, 3, 2, 2, 2, 457, 293, 3, 2, 2, 2,
	457, 297, 3, 2, 2, 2, 457, 302, 3, 2, 2, 2, 457, 306, 3, 2, 2, 2, 457,
	311, 3, 2, 2, 2, 457, 315, 3, 2, 2, 2, 457, 319, 3, 2, 2, 2, 457, 322,
	3, 2, 2, 2, 457, 327, 3, 2, 2, 2, 457, 332, 3, 2, 2, 2, 457, 336, 3, 2,
	2, 2, 457, 341, 3, 2, 2, 2, 457, 346, 3, 2, 2, 2, 457, 350, 3, 2, 2, 2,
	457, 355, 3, 2, 2, 2, 457, 359, 3, 2, 2, 2, 457, 363, 3, 2, 2, 2, 457,
	368, 3, 2, 2, 2, 457, 373, 3, 2, 2, 2, 457, 376, 3, 2, 2, 2, 457, 380,
	3, 2, 2, 2, 457, 384, 3, 2, 2, 2, 457, 389, 3, 2, 2, 2, 457, 393, 3, 2,
	2, 2, 457, 396, 3, 2, 2, 2, 457, 400, 3, 2, 2, 2, 457, 404, 3, 2, 2, 2,
	457, 409, 3, 2, 2, 2, 457, 413, 3, 2, 2, 2, 457, 417, 3, 2, 2, 2, 457,
	421, 3, 2, 2, 2, 457, 425, 3, 2, 2, 2, 457, 429, 3, 2, 2, 2, 457, 433,
	3, 2, 2, 2, 457, 437, 3, 2, 2, 2, 457, 441, 3, 2, 2, 2, 457, 445, 3, 2,
	2, 2, 457, 449, 3, 2, 2, 2, 457, 453, 3, 2, 2, 2, 458, 26, 3, 2, 2, 2,
	459, 460, 9, 3, 2, 2, 460, 28, 3, 2, 2, 2, 461, 462, 9, 4, 2, 2, 462, 30,
	3, 2, 2, 2, 463, 464, 9, 5, 2, 2, 464, 32, 3, 2, 2, 2, 465, 466, 9, 6,
	2, 2, 466, 34, 3, 2, 2, 2, 467, 468, 9, 7, 2, 2, 468, 36, 3, 2, 2, 2, 469,
	470, 9, 8, 2, 2, 470, 38, 3, 2, 2, 2, 471, 472, 9, 9, 2, 2, 472, 40, 3,
	2, 2, 2, 473, 474, 9, 10, 2, 2, 474, 42, 3, 2, 2, 2, 475, 476, 9, 11, 2,
	2, 476, 44, 3, 2, 2, 2, 477, 478, 9, 12, 2, 2, 478, 46, 3, 2, 2, 2, 479,
	480, 9, 13, 2, 2, 480, 48, 3, 2, 2, 2, 481, 482, 9, 14, 2, 2, 482, 50,
	3, 2, 2, 2, 483, 484, 9, 15, 2, 2, 484, 52, 3, 2, 2, 2, 485, 486, 9, 16,
	2, 2, 486, 54, 3, 2, 2, 2, 487, 488, 9, 17, 2, 2, 488, 56, 3, 2, 2, 2,
	489, 490, 9, 18, 2, 2, 490, 58, 3, 2, 2, 2, 491, 492, 9, 19, 2, 2, 492,
	60, 3, 2, 2, 2, 493, 494, 9, 20, 2, 2, 494, 62, 3, 2, 2, 2, 495, 496, 9,
	21, 2, 2, 496, 64, 3, 2, 2, 2, 497, 498, 9, 22, 2, 2, 498, 66, 3, 2, 2,
	2, 499, 500, 9, 23, 2, 2, 500, 68, 3, 2, 2, 2, 501, 502, 9, 24, 2, 2, 502,
	70, 3, 2, 2, 2, 503, 504, 9, 25, 2, 2, 504, 72, 3, 2, 2, 2, 505, 506, 9,
	26, 2, 2, 506, 74, 3, 2, 2, 2, 507, 508, 9, 27, 2, 2, 508, 76, 3, 2, 2,
	2, 509, 510, 9, 28, 2, 2, 510, 78, 3, 2, 2, 2, 511, 515, 9, 29, 2, 2, 512,
	514, 9, 30, 2, 2, 513, 512, 3, 2, 2, 2, 514, 517, 3, 2, 2, 2, 515, 513,
	3, 2, 2, 2, 515, 516, 3, 2, 2, 2, 516, 80, 3, 2, 2, 2, 517, 515, 3, 2,
	2, 2, 518, 520, 7, 38, 2, 2, 519, 518, 3, 2, 2, 2, 519, 520, 3, 2, 2, 2,
	520, 522, 3, 2, 2, 2, 521, 523, 9, 31, 2, 2, 522, 521, 3, 2, 2, 2, 523,
	524, 3, 2, 2, 2, 524, 522, 3, 2, 2, 2, 524, 525, 3, 2, 2, 2, 525, 527,
	3, 2, 2, 2, 526, 528, 9, 10, 2, 2, 527, 526, 3, 2, 2, 2, 527, 528, 3, 2,
	2, 2, 528, 82, 3, 2, 2, 2, 529, 533, 7, 61, 2, 2, 530, 532, 10, 32, 2,
	2, 531, 530, 3, 2, 2, 2, 532, 535, 3, 2, 2, 2, 533, 531, 3, 2, 2, 2, 533,
	534, 3, 2, 2, 2, 534, 536, 3, 2, 2, 2, 535, 533, 3, 2, 2, 2, 536, 537,
	8, 42, 2, 2, 537, 84, 3, 2, 2, 2, 538, 542, 7, 41, 2, 2, 539, 541, 10,
	33, 2, 2, 540, 539, 3, 2, 2, 2, 541, 544, 3, 2, 2, 2, 542, 540, 3, 2, 2,
	2, 542, 543, 3, 2, 2, 2, 543, 545, 3, 2, 2, 2, 544, 542, 3, 2, 2, 2, 545,
	546, 7, 41, 2, 2, 546, 86, 3, 2, 2, 2, 547, 549, 9, 32, 2, 2, 548, 547,
	3, 2, 2, 2, 549, 550, 3, 2, 2, 2, 550, 548, 3, 2, 2, 2, 550, 551, 3, 2,
	2, 2, 551, 88, 3, 2, 2, 2, 552, 553, 9, 34, 2, 2, 553, 554, 3, 2, 2, 2,
	554, 555, 8, 45, 2, 2, 555, 90, 3, 2, 2, 2, 13, 2, 138, 178, 457, 515,
	519, 524, 527, 533, 542, 550, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "':'", "','", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'$'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "REGISTER", "ASSEMBLER_DIRECTIVE",
	"OPCODE", "NAME", "NUMBER", "COMMENT", "STRING", "EOL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"REGISTER", "ASSEMBLER_DIRECTIVE", "OPCODE", "A", "B", "C", "D", "E", "F",
	"G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
	"V", "W", "X", "Y", "Z", "NAME", "NUMBER", "COMMENT", "STRING", "EOL",
	"WS",
}

type asmZ80Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewasmZ80Lexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *asmZ80Lexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewasmZ80Lexer(input antlr.CharStream) *asmZ80Lexer {
	l := new(asmZ80Lexer)
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
	l.GrammarFileName = "asmZ80.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// asmZ80Lexer tokens.
const (
	asmZ80LexerT__0                = 1
	asmZ80LexerT__1                = 2
	asmZ80LexerT__2                = 3
	asmZ80LexerT__3                = 4
	asmZ80LexerT__4                = 5
	asmZ80LexerT__5                = 6
	asmZ80LexerT__6                = 7
	asmZ80LexerT__7                = 8
	asmZ80LexerT__8                = 9
	asmZ80LexerREGISTER            = 10
	asmZ80LexerASSEMBLER_DIRECTIVE = 11
	asmZ80LexerOPCODE              = 12
	asmZ80LexerNAME                = 13
	asmZ80LexerNUMBER              = 14
	asmZ80LexerCOMMENT             = 15
	asmZ80LexerSTRING              = 16
	asmZ80LexerEOL                 = 17
	asmZ80LexerWS                  = 18
)