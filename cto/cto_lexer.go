// Code generated from CtoLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cto

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 51, 633,
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
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3,
	39, 3, 39, 3, 39, 5, 39, 351, 10, 39, 3, 39, 6, 39, 354, 10, 39, 13, 39,
	14, 39, 355, 3, 39, 5, 39, 359, 10, 39, 5, 39, 361, 10, 39, 3, 39, 5, 39,
	364, 10, 39, 3, 40, 3, 40, 7, 40, 368, 10, 40, 12, 40, 14, 40, 371, 11,
	40, 3, 40, 3, 40, 7, 40, 375, 10, 40, 12, 40, 14, 40, 378, 11, 40, 3, 40,
	5, 40, 381, 10, 40, 3, 40, 5, 40, 384, 10, 40, 3, 41, 3, 41, 3, 41, 5,
	41, 389, 10, 41, 3, 41, 3, 41, 5, 41, 393, 10, 41, 3, 41, 5, 41, 396, 10,
	41, 3, 41, 5, 41, 399, 10, 41, 3, 41, 3, 41, 3, 41, 5, 41, 404, 10, 41,
	3, 41, 5, 41, 407, 10, 41, 5, 41, 409, 10, 41, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 420, 10, 42, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 6, 44, 429, 10, 44, 13, 44, 14, 44,
	430, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 7, 45, 439, 10, 45, 12,
	45, 14, 45, 442, 11, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 7, 46,
	450, 10, 46, 12, 46, 14, 46, 453, 11, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3,
	46, 3, 47, 3, 47, 7, 47, 462, 10, 47, 12, 47, 14, 47, 465, 11, 47, 3, 47,
	3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3,
	50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 5, 51, 486, 10, 51,
	3, 52, 3, 52, 3, 52, 3, 52, 5, 52, 492, 10, 52, 3, 53, 3, 53, 3, 53, 3,
	54, 3, 54, 5, 54, 499, 10, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55,
	506, 10, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 515,
	10, 55, 5, 55, 517, 10, 55, 5, 55, 519, 10, 55, 3, 56, 3, 56, 3, 56, 3,
	56, 5, 56, 525, 10, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57,
	3, 57, 7, 57, 535, 10, 57, 12, 57, 14, 57, 538, 11, 57, 5, 57, 540, 10,
	57, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 6, 60, 548, 10, 60, 13, 60,
	14, 60, 549, 3, 61, 3, 61, 3, 61, 7, 61, 555, 10, 61, 12, 61, 14, 61, 558,
	11, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 7, 62, 565, 10, 62, 12, 62,
	14, 62, 568, 11, 62, 3, 62, 3, 62, 3, 63, 3, 63, 5, 63, 574, 10, 63, 3,
	63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 5, 64, 582, 10, 64, 3, 64, 5, 64,
	585, 10, 64, 3, 64, 3, 64, 3, 64, 6, 64, 590, 10, 64, 13, 64, 14, 64, 591,
	3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 5, 64, 599, 10, 64, 3, 65, 3, 65, 3,
	65, 7, 65, 604, 10, 65, 12, 65, 14, 65, 607, 11, 65, 3, 65, 5, 65, 610,
	10, 65, 3, 66, 3, 66, 3, 67, 3, 67, 7, 67, 616, 10, 67, 12, 67, 14, 67,
	619, 11, 67, 3, 67, 5, 67, 622, 10, 67, 3, 68, 3, 68, 5, 68, 626, 10, 68,
	3, 69, 3, 69, 3, 69, 3, 69, 5, 69, 632, 10, 69, 4, 451, 463, 2, 70, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31,
	61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40,
	79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 2,
	97, 2, 99, 2, 101, 2, 103, 2, 105, 2, 107, 2, 109, 2, 111, 2, 113, 2, 115,
	2, 117, 2, 119, 49, 121, 50, 123, 51, 125, 2, 127, 2, 129, 2, 131, 2, 133,
	2, 135, 2, 137, 2, 3, 2, 30, 3, 2, 51, 59, 4, 2, 78, 78, 110, 110, 3, 2,
	50, 57, 4, 2, 50, 57, 97, 97, 6, 2, 70, 70, 72, 72, 102, 102, 104, 104,
	5, 2, 11, 12, 14, 15, 34, 34, 4, 2, 12, 12, 15, 15, 4, 2, 36, 36, 41, 41,
	3, 2, 50, 50, 3, 2, 50, 59, 3, 2, 51, 51, 3, 2, 50, 52, 3, 2, 50, 53, 3,
	2, 50, 51, 3, 2, 50, 55, 3, 2, 54, 54, 3, 2, 55, 55, 3, 2, 53, 53, 6, 2,
	12, 12, 15, 15, 36, 36, 94, 94, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47,
	47, 10, 2, 36, 36, 41, 41, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116,
	118, 118, 5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 50, 59, 97, 97, 6, 2, 38,
	38, 67, 92, 97, 97, 99, 124, 4, 2, 2, 129, 55298, 56321, 3, 2, 55298, 56321,
	3, 2, 56322, 57345, 2, 663, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2,
	2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3,
	2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61,
	3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2,
	69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2,
	2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2,
	2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2,
	2, 2, 2, 93, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123,
	3, 2, 2, 2, 3, 139, 3, 2, 2, 2, 5, 148, 3, 2, 2, 2, 7, 154, 3, 2, 2, 2,
	9, 162, 3, 2, 2, 2, 11, 170, 3, 2, 2, 2, 13, 175, 3, 2, 2, 2, 15, 181,
	3, 2, 2, 2, 17, 189, 3, 2, 2, 2, 19, 203, 3, 2, 2, 2, 21, 210, 3, 2, 2,
	2, 23, 220, 3, 2, 2, 2, 25, 229, 3, 2, 2, 2, 27, 241, 3, 2, 2, 2, 29, 247,
	3, 2, 2, 2, 31, 253, 3, 2, 2, 2, 33, 265, 3, 2, 2, 2, 35, 273, 3, 2, 2,
	2, 37, 282, 3, 2, 2, 2, 39, 289, 3, 2, 2, 2, 41, 297, 3, 2, 2, 2, 43, 302,
	3, 2, 2, 2, 45, 309, 3, 2, 2, 2, 47, 311, 3, 2, 2, 2, 49, 313, 3, 2, 2,
	2, 51, 315, 3, 2, 2, 2, 53, 317, 3, 2, 2, 2, 55, 319, 3, 2, 2, 2, 57, 321,
	3, 2, 2, 2, 59, 323, 3, 2, 2, 2, 61, 325, 3, 2, 2, 2, 63, 327, 3, 2, 2,
	2, 65, 329, 3, 2, 2, 2, 67, 331, 3, 2, 2, 2, 69, 333, 3, 2, 2, 2, 71, 335,
	3, 2, 2, 2, 73, 339, 3, 2, 2, 2, 75, 344, 3, 2, 2, 2, 77, 360, 3, 2, 2,
	2, 79, 365, 3, 2, 2, 2, 81, 408, 3, 2, 2, 2, 83, 419, 3, 2, 2, 2, 85, 421,
	3, 2, 2, 2, 87, 428, 3, 2, 2, 2, 89, 434, 3, 2, 2, 2, 91, 445, 3, 2, 2,
	2, 93, 459, 3, 2, 2, 2, 95, 468, 3, 2, 2, 2, 97, 470, 3, 2, 2, 2, 99, 476,
	3, 2, 2, 2, 101, 485, 3, 2, 2, 2, 103, 491, 3, 2, 2, 2, 105, 493, 3, 2,
	2, 2, 107, 498, 3, 2, 2, 2, 109, 518, 3, 2, 2, 2, 111, 524, 3, 2, 2, 2,
	113, 526, 3, 2, 2, 2, 115, 541, 3, 2, 2, 2, 117, 544, 3, 2, 2, 2, 119,
	547, 3, 2, 2, 2, 121, 551, 3, 2, 2, 2, 123, 561, 3, 2, 2, 2, 125, 571,
	3, 2, 2, 2, 127, 598, 3, 2, 2, 2, 129, 600, 3, 2, 2, 2, 131, 611, 3, 2,
	2, 2, 133, 613, 3, 2, 2, 2, 135, 625, 3, 2, 2, 2, 137, 631, 3, 2, 2, 2,
	139, 140, 7, 99, 2, 2, 140, 141, 7, 100, 2, 2, 141, 142, 7, 117, 2, 2,
	142, 143, 7, 118, 2, 2, 143, 144, 7, 116, 2, 2, 144, 145, 7, 99, 2, 2,
	145, 146, 7, 101, 2, 2, 146, 147, 7, 118, 2, 2, 147, 4, 3, 2, 2, 2, 148,
	149, 7, 99, 2, 2, 149, 150, 7, 117, 2, 2, 150, 151, 7, 117, 2, 2, 151,
	152, 7, 103, 2, 2, 152, 153, 7, 118, 2, 2, 153, 6, 3, 2, 2, 2, 154, 155,
	7, 101, 2, 2, 155, 156, 7, 113, 2, 2, 156, 157, 7, 112, 2, 2, 157, 158,
	7, 101, 2, 2, 158, 159, 7, 103, 2, 2, 159, 160, 7, 114, 2, 2, 160, 161,
	7, 118, 2, 2, 161, 8, 3, 2, 2, 2, 162, 163, 7, 102, 2, 2, 163, 164, 7,
	103, 2, 2, 164, 165, 7, 104, 2, 2, 165, 166, 7, 99, 2, 2, 166, 167, 7,
	119, 2, 2, 167, 168, 7, 110, 2, 2, 168, 169, 7, 118, 2, 2, 169, 10, 3,
	2, 2, 2, 170, 171, 7, 103, 2, 2, 171, 172, 7, 112, 2, 2, 172, 173, 7, 119,
	2, 2, 173, 174, 7, 111, 2, 2, 174, 12, 3, 2, 2, 2, 175, 176, 7, 103, 2,
	2, 176, 177, 7, 120, 2, 2, 177, 178, 7, 103, 2, 2, 178, 179, 7, 112, 2,
	2, 179, 180, 7, 118, 2, 2, 180, 14, 3, 2, 2, 2, 181, 182, 7, 103, 2, 2,
	182, 183, 7, 122, 2, 2, 183, 184, 7, 118, 2, 2, 184, 185, 7, 103, 2, 2,
	185, 186, 7, 112, 2, 2, 186, 187, 7, 102, 2, 2, 187, 188, 7, 117, 2, 2,
	188, 16, 3, 2, 2, 2, 189, 190, 7, 107, 2, 2, 190, 191, 7, 102, 2, 2, 191,
	192, 7, 103, 2, 2, 192, 193, 7, 112, 2, 2, 193, 194, 7, 118, 2, 2, 194,
	195, 7, 107, 2, 2, 195, 196, 7, 104, 2, 2, 196, 197, 7, 107, 2, 2, 197,
	198, 7, 103, 2, 2, 198, 199, 7, 102, 2, 2, 199, 200, 7, 34, 2, 2, 200,
	201, 7, 100, 2, 2, 201, 202, 7, 123, 2, 2, 202, 18, 3, 2, 2, 2, 203, 204,
	7, 107, 2, 2, 204, 205, 7, 111, 2, 2, 205, 206, 7, 114, 2, 2, 206, 207,
	7, 113, 2, 2, 207, 208, 7, 116, 2, 2, 208, 209, 7, 118, 2, 2, 209, 20,
	3, 2, 2, 2, 210, 211, 7, 112, 2, 2, 211, 212, 7, 99, 2, 2, 212, 213, 7,
	111, 2, 2, 213, 214, 7, 103, 2, 2, 214, 215, 7, 117, 2, 2, 215, 216, 7,
	114, 2, 2, 216, 217, 7, 99, 2, 2, 217, 218, 7, 101, 2, 2, 218, 219, 7,
	103, 2, 2, 219, 22, 3, 2, 2, 2, 220, 221, 7, 113, 2, 2, 221, 222, 7, 114,
	2, 2, 222, 223, 7, 118, 2, 2, 223, 224, 7, 107, 2, 2, 224, 225, 7, 113,
	2, 2, 225, 226, 7, 112, 2, 2, 226, 227, 7, 99, 2, 2, 227, 228, 7, 110,
	2, 2, 228, 24, 3, 2, 2, 2, 229, 230, 7, 114, 2, 2, 230, 231, 7, 99, 2,
	2, 231, 232, 7, 116, 2, 2, 232, 233, 7, 118, 2, 2, 233, 234, 7, 107, 2,
	2, 234, 235, 7, 101, 2, 2, 235, 236, 7, 107, 2, 2, 236, 237, 7, 114, 2,
	2, 237, 238, 7, 99, 2, 2, 238, 239, 7, 112, 2, 2, 239, 240, 7, 118, 2,
	2, 240, 26, 3, 2, 2, 2, 241, 242, 7, 116, 2, 2, 242, 243, 7, 99, 2, 2,
	243, 244, 7, 112, 2, 2, 244, 245, 7, 105, 2, 2, 245, 246, 7, 103, 2, 2,
	246, 28, 3, 2, 2, 2, 247, 248, 7, 116, 2, 2, 248, 249, 7, 103, 2, 2, 249,
	250, 7, 105, 2, 2, 250, 251, 7, 103, 2, 2, 251, 252, 7, 122, 2, 2, 252,
	30, 3, 2, 2, 2, 253, 254, 7, 118, 2, 2, 254, 255, 7, 116, 2, 2, 255, 256,
	7, 99, 2, 2, 256, 257, 7, 112, 2, 2, 257, 258, 7, 117, 2, 2, 258, 259,
	7, 99, 2, 2, 259, 260, 7, 101, 2, 2, 260, 261, 7, 118, 2, 2, 261, 262,
	7, 107, 2, 2, 262, 263, 7, 113, 2, 2, 263, 264, 7, 112, 2, 2, 264, 32,
	3, 2, 2, 2, 265, 266, 7, 68, 2, 2, 266, 267, 7, 113, 2, 2, 267, 268, 7,
	113, 2, 2, 268, 269, 7, 110, 2, 2, 269, 270, 7, 103, 2, 2, 270, 271, 7,
	99, 2, 2, 271, 272, 7, 112, 2, 2, 272, 34, 3, 2, 2, 2, 273, 274, 7, 70,
	2, 2, 274, 275, 7, 99, 2, 2, 275, 276, 7, 118, 2, 2, 276, 277, 7, 103,
	2, 2, 277, 278, 7, 86, 2, 2, 278, 279, 7, 107, 2, 2, 279, 280, 7, 111,
	2, 2, 280, 281, 7, 103, 2, 2, 281, 36, 3, 2, 2, 2, 282, 283, 7, 70, 2,
	2, 283, 284, 7, 113, 2, 2, 284, 285, 7, 119, 2, 2, 285, 286, 7, 100, 2,
	2, 286, 287, 7, 110, 2, 2, 287, 288, 7, 103, 2, 2, 288, 38, 3, 2, 2, 2,
	289, 290, 7, 75, 2, 2, 290, 291, 7, 112, 2, 2, 291, 292, 7, 118, 2, 2,
	292, 293, 7, 103, 2, 2, 293, 294, 7, 105, 2, 2, 294, 295, 7, 103, 2, 2,
	295, 296, 7, 116, 2, 2, 296, 40, 3, 2, 2, 2, 297, 298, 7, 78, 2, 2, 298,
	299, 7, 113, 2, 2, 299, 300, 7, 112, 2, 2, 300, 301, 7, 105, 2, 2, 301,
	42, 3, 2, 2, 2, 302, 303, 7, 85, 2, 2, 303, 304, 7, 118, 2, 2, 304, 305,
	7, 116, 2, 2, 305, 306, 7, 107, 2, 2, 306, 307, 7, 112, 2, 2, 307, 308,
	7, 105, 2, 2, 308, 44, 3, 2, 2, 2, 309, 310, 7, 42, 2, 2, 310, 46, 3, 2,
	2, 2, 311, 312, 7, 43, 2, 2, 312, 48, 3, 2, 2, 2, 313, 314, 7, 125, 2,
	2, 314, 50, 3, 2, 2, 2, 315, 316, 7, 127, 2, 2, 316, 52, 3, 2, 2, 2, 317,
	318, 7, 93, 2, 2, 318, 54, 3, 2, 2, 2, 319, 320, 7, 95, 2, 2, 320, 56,
	3, 2, 2, 2, 321, 322, 7, 61, 2, 2, 322, 58, 3, 2, 2, 2, 323, 324, 7, 46,
	2, 2, 324, 60, 3, 2, 2, 2, 325, 326, 7, 48, 2, 2, 326, 62, 3, 2, 2, 2,
	327, 328, 7, 60, 2, 2, 328, 64, 3, 2, 2, 2, 329, 330, 7, 63, 2, 2, 330,
	66, 3, 2, 2, 2, 331, 332, 7, 44, 2, 2, 332, 68, 3, 2, 2, 2, 333, 334, 7,
	66, 2, 2, 334, 70, 3, 2, 2, 2, 335, 336, 7, 48, 2, 2, 336, 337, 7, 48,
	2, 2, 337, 338, 7, 48, 2, 2, 338, 72, 3, 2, 2, 2, 339, 340, 7, 47, 2, 2,
	340, 341, 7, 47, 2, 2, 341, 342, 7, 64, 2, 2, 342, 343, 7, 34, 2, 2, 343,
	74, 3, 2, 2, 2, 344, 345, 7, 113, 2, 2, 345, 346, 7, 34, 2, 2, 346, 76,
	3, 2, 2, 2, 347, 361, 7, 50, 2, 2, 348, 358, 9, 2, 2, 2, 349, 351, 5, 133,
	67, 2, 350, 349, 3, 2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 359, 3, 2, 2, 2,
	352, 354, 7, 97, 2, 2, 353, 352, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355,
	353, 3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 359,
	5, 133, 67, 2, 358, 350, 3, 2, 2, 2, 358, 353, 3, 2, 2, 2, 359, 361, 3,
	2, 2, 2, 360, 347, 3, 2, 2, 2, 360, 348, 3, 2, 2, 2, 361, 363, 3, 2, 2,
	2, 362, 364, 9, 3, 2, 2, 363, 362, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364,
	78, 3, 2, 2, 2, 365, 369, 7, 50, 2, 2, 366, 368, 7, 97, 2, 2, 367, 366,
	3, 2, 2, 2, 368, 371, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2,
	2, 2, 370, 372, 3, 2, 2, 2, 371, 369, 3, 2, 2, 2, 372, 380, 9, 4, 2, 2,
	373, 375, 9, 5, 2, 2, 374, 373, 3, 2, 2, 2, 375, 378, 3, 2, 2, 2, 376,
	374, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 379, 3, 2, 2, 2, 378, 376,
	3, 2, 2, 2, 379, 381, 9, 4, 2, 2, 380, 376, 3, 2, 2, 2, 380, 381, 3, 2,
	2, 2, 381, 383, 3, 2, 2, 2, 382, 384, 9, 3, 2, 2, 383, 382, 3, 2, 2, 2,
	383, 384, 3, 2, 2, 2, 384, 80, 3, 2, 2, 2, 385, 386, 5, 133, 67, 2, 386,
	388, 7, 48, 2, 2, 387, 389, 5, 133, 67, 2, 388, 387, 3, 2, 2, 2, 388, 389,
	3, 2, 2, 2, 389, 393, 3, 2, 2, 2, 390, 391, 7, 48, 2, 2, 391, 393, 5, 133,
	67, 2, 392, 385, 3, 2, 2, 2, 392, 390, 3, 2, 2, 2, 393, 395, 3, 2, 2, 2,
	394, 396, 5, 125, 63, 2, 395, 394, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396,
	398, 3, 2, 2, 2, 397, 399, 9, 6, 2, 2, 398, 397, 3, 2, 2, 2, 398, 399,
	3, 2, 2, 2, 399, 409, 3, 2, 2, 2, 400, 406, 5, 133, 67, 2, 401, 403, 5,
	125, 63, 2, 402, 404, 9, 6, 2, 2, 403, 402, 3, 2, 2, 2, 403, 404, 3, 2,
	2, 2, 404, 407, 3, 2, 2, 2, 405, 407, 9, 6, 2, 2, 406, 401, 3, 2, 2, 2,
	406, 405, 3, 2, 2, 2, 407, 409, 3, 2, 2, 2, 408, 392, 3, 2, 2, 2, 408,
	400, 3, 2, 2, 2, 409, 82, 3, 2, 2, 2, 410, 411, 7, 118, 2, 2, 411, 412,
	7, 116, 2, 2, 412, 413, 7, 119, 2, 2, 413, 420, 7, 103, 2, 2, 414, 415,
	7, 104, 2, 2, 415, 416, 7, 99, 2, 2, 416, 417, 7, 110, 2, 2, 417, 418,
	7, 117, 2, 2, 418, 420, 7, 103, 2, 2, 419, 410, 3, 2, 2, 2, 419, 414, 3,
	2, 2, 2, 420, 84, 3, 2, 2, 2, 421, 422, 5, 95, 48, 2, 422, 423, 5, 97,
	49, 2, 423, 424, 7, 86, 2, 2, 424, 425, 5, 105, 53, 2, 425, 426, 5, 95,
	48, 2, 426, 86, 3, 2, 2, 2, 427, 429, 9, 7, 2, 2, 428, 427, 3, 2, 2, 2,
	429, 430, 3, 2, 2, 2, 430, 428, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431,
	432, 3, 2, 2, 2, 432, 433, 8, 44, 2, 2, 433, 88, 3, 2, 2, 2, 434, 435,
	7, 49, 2, 2, 435, 436, 7, 49, 2, 2, 436, 440, 3, 2, 2, 2, 437, 439, 10,
	8, 2, 2, 438, 437, 3, 2, 2, 2, 439, 442, 3, 2, 2, 2, 440, 438, 3, 2, 2,
	2, 440, 441, 3, 2, 2, 2, 441, 443, 3, 2, 2, 2, 442, 440, 3, 2, 2, 2, 443,
	444, 8, 45, 3, 2, 444, 90, 3, 2, 2, 2, 445, 446, 7, 49, 2, 2, 446, 447,
	7, 44, 2, 2, 447, 451, 3, 2, 2, 2, 448, 450, 11, 2, 2, 2, 449, 448, 3,
	2, 2, 2, 450, 453, 3, 2, 2, 2, 451, 452, 3, 2, 2, 2, 451, 449, 3, 2, 2,
	2, 452, 454, 3, 2, 2, 2, 453, 451, 3, 2, 2, 2, 454, 455, 7, 44, 2, 2, 455,
	456, 7, 49, 2, 2, 456, 457, 3, 2, 2, 2, 457, 458, 8, 46, 3, 2, 458, 92,
	3, 2, 2, 2, 459, 463, 7, 49, 2, 2, 460, 462, 11, 2, 2, 2, 461, 460, 3,
	2, 2, 2, 462, 465, 3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 463, 461, 3, 2, 2,
	2, 464, 466, 3, 2, 2, 2, 465, 463, 3, 2, 2, 2, 466, 467, 7, 49, 2, 2, 467,
	94, 3, 2, 2, 2, 468, 469, 9, 9, 2, 2, 469, 96, 3, 2, 2, 2, 470, 471, 5,
	99, 50, 2, 471, 472, 7, 47, 2, 2, 472, 473, 5, 101, 51, 2, 473, 474, 7,
	47, 2, 2, 474, 475, 5, 103, 52, 2, 475, 98, 3, 2, 2, 2, 476, 477, 5, 117,
	59, 2, 477, 478, 5, 117, 59, 2, 478, 479, 5, 117, 59, 2, 479, 480, 5, 117,
	59, 2, 480, 100, 3, 2, 2, 2, 481, 482, 9, 10, 2, 2, 482, 486, 9, 11, 2,
	2, 483, 484, 9, 12, 2, 2, 484, 486, 9, 13, 2, 2, 485, 481, 3, 2, 2, 2,
	485, 483, 3, 2, 2, 2, 486, 102, 3, 2, 2, 2, 487, 488, 9, 13, 2, 2, 488,
	492, 9, 11, 2, 2, 489, 490, 9, 14, 2, 2, 490, 492, 9, 15, 2, 2, 491, 487,
	3, 2, 2, 2, 491, 489, 3, 2, 2, 2, 492, 104, 3, 2, 2, 2, 493, 494, 5, 113,
	57, 2, 494, 495, 5, 107, 54, 2, 495, 106, 3, 2, 2, 2, 496, 499, 7, 92,
	2, 2, 497, 499, 5, 109, 55, 2, 498, 496, 3, 2, 2, 2, 498, 497, 3, 2, 2,
	2, 499, 108, 3, 2, 2, 2, 500, 501, 7, 47, 2, 2, 501, 502, 9, 15, 2, 2,
	502, 505, 9, 13, 2, 2, 503, 504, 7, 60, 2, 2, 504, 506, 5, 111, 56, 2,
	505, 503, 3, 2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 519, 3, 2, 2, 2, 507,
	508, 7, 45, 2, 2, 508, 509, 9, 15, 2, 2, 509, 516, 9, 16, 2, 2, 510, 514,
	7, 60, 2, 2, 511, 515, 5, 111, 56, 2, 512, 513, 9, 17, 2, 2, 513, 515,
	9, 18, 2, 2, 514, 511, 3, 2, 2, 2, 514, 512, 3, 2, 2, 2, 515, 517, 3, 2,
	2, 2, 516, 510, 3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517, 519, 3, 2, 2, 2,
	518, 500, 3, 2, 2, 2, 518, 507, 3, 2, 2, 2, 519, 110, 3, 2, 2, 2, 520,
	521, 9, 10, 2, 2, 521, 525, 9, 10, 2, 2, 522, 523, 9, 19, 2, 2, 523, 525,
	9, 10, 2, 2, 524, 520, 3, 2, 2, 2, 524, 522, 3, 2, 2, 2, 525, 112, 3, 2,
	2, 2, 526, 527, 9, 13, 2, 2, 527, 528, 9, 14, 2, 2, 528, 529, 7, 60, 2,
	2, 529, 530, 5, 115, 58, 2, 530, 531, 7, 60, 2, 2, 531, 539, 5, 115, 58,
	2, 532, 536, 7, 48, 2, 2, 533, 535, 9, 11, 2, 2, 534, 533, 3, 2, 2, 2,
	535, 538, 3, 2, 2, 2, 536, 534, 3, 2, 2, 2, 536, 537, 3, 2, 2, 2, 537,
	540, 3, 2, 2, 2, 538, 536, 3, 2, 2, 2, 539, 532, 3, 2, 2, 2, 539, 540,
	3, 2, 2, 2, 540, 114, 3, 2, 2, 2, 541, 542, 9, 16, 2, 2, 542, 543, 5, 117,
	59, 2, 543, 116, 3, 2, 2, 2, 544, 545, 9, 11, 2, 2, 545, 118, 3, 2, 2,
	2, 546, 548, 5, 135, 68, 2, 547, 546, 3, 2, 2, 2, 548, 549, 3, 2, 2, 2,
	549, 547, 3, 2, 2, 2, 549, 550, 3, 2, 2, 2, 550, 120, 3, 2, 2, 2, 551,
	556, 7, 41, 2, 2, 552, 555, 10, 20, 2, 2, 553, 555, 5, 127, 64, 2, 554,
	552, 3, 2, 2, 2, 554, 553, 3, 2, 2, 2, 555, 558, 3, 2, 2, 2, 556, 554,
	3, 2, 2, 2, 556, 557, 3, 2, 2, 2, 557, 559, 3, 2, 2, 2, 558, 556, 3, 2,
	2, 2, 559, 560, 7, 41, 2, 2, 560, 122, 3, 2, 2, 2, 561, 566, 7, 36, 2,
	2, 562, 565, 10, 20, 2, 2, 563, 565, 5, 127, 64, 2, 564, 562, 3, 2, 2,
	2, 564, 563, 3, 2, 2, 2, 565, 568, 3, 2, 2, 2, 566, 564, 3, 2, 2, 2, 566,
	567, 3, 2, 2, 2, 567, 569, 3, 2, 2, 2, 568, 566, 3, 2, 2, 2, 569, 570,
	7, 36, 2, 2, 570, 124, 3, 2, 2, 2, 571, 573, 9, 21, 2, 2, 572, 574, 9,
	22, 2, 2, 573, 572, 3, 2, 2, 2, 573, 574, 3, 2, 2, 2, 574, 575, 3, 2, 2,
	2, 575, 576, 5, 133, 67, 2, 576, 126, 3, 2, 2, 2, 577, 578, 7, 94, 2, 2,
	578, 599, 9, 23, 2, 2, 579, 584, 7, 94, 2, 2, 580, 582, 9, 14, 2, 2, 581,
	580, 3, 2, 2, 2, 581, 582, 3, 2, 2, 2, 582, 583, 3, 2, 2, 2, 583, 585,
	9, 4, 2, 2, 584, 581, 3, 2, 2, 2, 584, 585, 3, 2, 2, 2, 585, 586, 3, 2,
	2, 2, 586, 599, 9, 4, 2, 2, 587, 589, 7, 94, 2, 2, 588, 590, 7, 119, 2,
	2, 589, 588, 3, 2, 2, 2, 590, 591, 3, 2, 2, 2, 591, 589, 3, 2, 2, 2, 591,
	592, 3, 2, 2, 2, 592, 593, 3, 2, 2, 2, 593, 594, 5, 131, 66, 2, 594, 595,
	5, 131, 66, 2, 595, 596, 5, 131, 66, 2, 596, 597, 5, 131, 66, 2, 597, 599,
	3, 2, 2, 2, 598, 577, 3, 2, 2, 2, 598, 579, 3, 2, 2, 2, 598, 587, 3, 2,
	2, 2, 599, 128, 3, 2, 2, 2, 600, 609, 5, 131, 66, 2, 601, 604, 5, 131,
	66, 2, 602, 604, 7, 97, 2, 2, 603, 601, 3, 2, 2, 2, 603, 602, 3, 2, 2,
	2, 604, 607, 3, 2, 2, 2, 605, 603, 3, 2, 2, 2, 605, 606, 3, 2, 2, 2, 606,
	608, 3, 2, 2, 2, 607, 605, 3, 2, 2, 2, 608, 610, 5, 131, 66, 2, 609, 605,
	3, 2, 2, 2, 609, 610, 3, 2, 2, 2, 610, 130, 3, 2, 2, 2, 611, 612, 9, 24,
	2, 2, 612, 132, 3, 2, 2, 2, 613, 621, 9, 11, 2, 2, 614, 616, 9, 25, 2,
	2, 615, 614, 3, 2, 2, 2, 616, 619, 3, 2, 2, 2, 617, 615, 3, 2, 2, 2, 617,
	618, 3, 2, 2, 2, 618, 620, 3, 2, 2, 2, 619, 617, 3, 2, 2, 2, 620, 622,
	9, 11, 2, 2, 621, 617, 3, 2, 2, 2, 621, 622, 3, 2, 2, 2, 622, 134, 3, 2,
	2, 2, 623, 626, 5, 137, 69, 2, 624, 626, 9, 11, 2, 2, 625, 623, 3, 2, 2,
	2, 625, 624, 3, 2, 2, 2, 626, 136, 3, 2, 2, 2, 627, 632, 9, 26, 2, 2, 628,
	632, 10, 27, 2, 2, 629, 630, 9, 28, 2, 2, 630, 632, 9, 29, 2, 2, 631, 627,
	3, 2, 2, 2, 631, 628, 3, 2, 2, 2, 631, 629, 3, 2, 2, 2, 632, 138, 3, 2,
	2, 2, 51, 2, 350, 355, 358, 360, 363, 369, 376, 380, 383, 388, 392, 395,
	398, 403, 406, 408, 419, 430, 440, 451, 463, 485, 491, 498, 505, 514, 516,
	518, 524, 536, 539, 549, 554, 556, 564, 566, 573, 581, 584, 591, 598, 603,
	605, 609, 617, 621, 625, 631, 4, 8, 2, 2, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'abstract'", "'asset'", "'concept'", "'default'", "'enum'", "'event'",
	"'extends'", "'identified by'", "'import'", "'namespace'", "'optional'",
	"'participant'", "'range'", "'regex'", "'transaction'", "'Boolean'", "'DateTime'",
	"'Double'", "'Integer'", "'Long'", "'String'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "';'", "','", "'.'", "':'", "'='", "'*'", "'@'", "'...'",
	"'--> '", "'o '",
}

var lexerSymbolicNames = []string{
	"", "ABSTRACT", "ASSET", "CONCEPT", "DEFAULT", "ENUM", "EVENT", "EXTENDS",
	"IDENTIFIED", "IMPORT", "NAMESPACE", "OPTIONAL", "PARTICIPANT", "RANGE",
	"REGEX", "TRANSACTION", "BOOLEAN", "DATE_TIME", "DOUBLE", "INTEGER", "LONG",
	"STRING", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI",
	"COMMA", "DOT", "COLON", "ASSIGN", "MUL", "AT", "ELLIPSIS", "REF", "VAR",
	"DECIMAL_LITERAL", "OCT_LITERAL", "FLOAT_LITERAL", "BOOL_LITERAL", "DATE_TIME_LITERAL",
	"WS", "LINE_COMMENT", "COMMENT", "REGEX_EXPR", "IDENTIFIER", "CHAR_LITERAL",
	"STRING_LITERAL",
}

var lexerRuleNames = []string{
	"ABSTRACT", "ASSET", "CONCEPT", "DEFAULT", "ENUM", "EVENT", "EXTENDS",
	"IDENTIFIED", "IMPORT", "NAMESPACE", "OPTIONAL", "PARTICIPANT", "RANGE",
	"REGEX", "TRANSACTION", "BOOLEAN", "DATE_TIME", "DOUBLE", "INTEGER", "LONG",
	"STRING", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI",
	"COMMA", "DOT", "COLON", "ASSIGN", "MUL", "AT", "ELLIPSIS", "REF", "VAR",
	"DECIMAL_LITERAL", "OCT_LITERAL", "FLOAT_LITERAL", "BOOL_LITERAL", "DATE_TIME_LITERAL",
	"WS", "LINE_COMMENT", "COMMENT", "REGEX_EXPR", "Bound", "FullDate", "Year",
	"Month", "Day", "FullTime", "TimeOffset", "TimeNumOffset", "HalfHour",
	"PartialTime", "Sixty", "Digit", "IDENTIFIER", "CHAR_LITERAL", "STRING_LITERAL",
	"ExponentPart", "EscapeSequence", "HexDigits", "HexDigit", "Digits", "LetterOrDigit",
	"Letter",
}

type CtoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCtoLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CtoLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCtoLexer(input antlr.CharStream) *CtoLexer {
	l := new(CtoLexer)
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
	l.GrammarFileName = "CtoLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CtoLexer tokens.
const (
	CtoLexerABSTRACT          = 1
	CtoLexerASSET             = 2
	CtoLexerCONCEPT           = 3
	CtoLexerDEFAULT           = 4
	CtoLexerENUM              = 5
	CtoLexerEVENT             = 6
	CtoLexerEXTENDS           = 7
	CtoLexerIDENTIFIED        = 8
	CtoLexerIMPORT            = 9
	CtoLexerNAMESPACE         = 10
	CtoLexerOPTIONAL          = 11
	CtoLexerPARTICIPANT       = 12
	CtoLexerRANGE             = 13
	CtoLexerREGEX             = 14
	CtoLexerTRANSACTION       = 15
	CtoLexerBOOLEAN           = 16
	CtoLexerDATE_TIME         = 17
	CtoLexerDOUBLE            = 18
	CtoLexerINTEGER           = 19
	CtoLexerLONG              = 20
	CtoLexerSTRING            = 21
	CtoLexerLPAREN            = 22
	CtoLexerRPAREN            = 23
	CtoLexerLBRACE            = 24
	CtoLexerRBRACE            = 25
	CtoLexerLBRACK            = 26
	CtoLexerRBRACK            = 27
	CtoLexerSEMI              = 28
	CtoLexerCOMMA             = 29
	CtoLexerDOT               = 30
	CtoLexerCOLON             = 31
	CtoLexerASSIGN            = 32
	CtoLexerMUL               = 33
	CtoLexerAT                = 34
	CtoLexerELLIPSIS          = 35
	CtoLexerREF               = 36
	CtoLexerVAR               = 37
	CtoLexerDECIMAL_LITERAL   = 38
	CtoLexerOCT_LITERAL       = 39
	CtoLexerFLOAT_LITERAL     = 40
	CtoLexerBOOL_LITERAL      = 41
	CtoLexerDATE_TIME_LITERAL = 42
	CtoLexerWS                = 43
	CtoLexerLINE_COMMENT      = 44
	CtoLexerCOMMENT           = 45
	CtoLexerREGEX_EXPR        = 46
	CtoLexerIDENTIFIER        = 47
	CtoLexerCHAR_LITERAL      = 48
	CtoLexerSTRING_LITERAL    = 49
)
