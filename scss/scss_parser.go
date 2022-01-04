// Code generated from ScssParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package scss // ScssParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 79, 728,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4,
	71, 9, 71, 4, 72, 9, 72, 3, 2, 7, 2, 146, 10, 2, 12, 2, 14, 2, 149, 11,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 163, 10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 168, 10, 4, 12, 4, 14, 4,
	171, 11, 4, 3, 4, 5, 4, 174, 10, 4, 3, 5, 3, 5, 5, 5, 178, 10, 5, 3, 6,
	5, 6, 181, 10, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 6, 7, 188, 10, 7, 13, 7,
	14, 7, 189, 3, 8, 3, 8, 3, 8, 7, 8, 195, 10, 8, 12, 8, 14, 8, 198, 11,
	8, 3, 8, 5, 8, 201, 10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 206, 10, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 5, 9, 212, 10, 9, 3, 10, 3, 10, 3, 10, 5, 10, 217, 10, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 223, 10, 10, 3, 10, 5, 10, 226, 10,
	10, 5, 10, 228, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 235,
	10, 11, 3, 11, 5, 11, 238, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5,
	12, 245, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 253,
	10, 12, 3, 12, 5, 12, 256, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 262,
	10, 13, 3, 13, 5, 13, 265, 10, 13, 3, 13, 3, 13, 3, 13, 5, 13, 270, 10,
	13, 3, 13, 3, 13, 3, 14, 7, 14, 275, 10, 14, 12, 14, 14, 14, 278, 11, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5,
	16, 290, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 297, 10, 17,
	3, 17, 5, 17, 300, 10, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 315, 10, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 7, 21, 321, 10, 21, 12, 21, 14, 21, 324, 11, 21, 3,
	21, 5, 21, 327, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 340, 10, 24, 3, 24, 5, 24, 343, 10,
	24, 3, 25, 3, 25, 3, 25, 5, 25, 348, 10, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	5, 25, 354, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 361, 10,
	26, 3, 26, 5, 26, 364, 10, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 5, 29, 380, 10,
	29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 390,
	10, 31, 12, 31, 14, 31, 393, 11, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 32, 5, 32, 402, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 5, 33, 411, 10, 33, 3, 33, 5, 33, 414, 10, 33, 3, 33, 3, 33,
	5, 33, 418, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 424, 10, 34, 3,
	35, 3, 35, 3, 35, 5, 35, 429, 10, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	7, 36, 436, 10, 36, 12, 36, 14, 36, 439, 11, 36, 3, 36, 5, 36, 442, 10,
	36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 39, 3, 39, 3, 39, 7, 39, 457, 10, 39, 12, 39, 14, 39, 460, 11, 39, 5,
	39, 462, 10, 39, 3, 40, 5, 40, 465, 10, 40, 3, 40, 3, 40, 3, 40, 7, 40,
	470, 10, 40, 12, 40, 14, 40, 473, 11, 40, 3, 40, 3, 40, 3, 40, 7, 40, 478,
	10, 40, 12, 40, 14, 40, 481, 11, 40, 5, 40, 483, 10, 40, 3, 41, 3, 41,
	3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 491, 10, 42, 3, 42, 3, 42, 3, 43, 3,
	43, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 7, 45, 503, 10, 45, 12, 45,
	14, 45, 506, 11, 45, 3, 45, 5, 45, 509, 10, 45, 3, 45, 3, 45, 3, 46, 3,
	46, 3, 46, 7, 46, 516, 10, 46, 12, 46, 14, 46, 519, 11, 46, 3, 47, 6, 47,
	522, 10, 47, 13, 47, 14, 47, 523, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 48, 3, 48, 5, 48, 536, 10, 48, 3, 49, 3, 49, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 50, 5, 50, 545, 10, 50, 3, 50, 3, 50, 5, 50, 549,
	10, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 5, 51, 556, 10, 51, 3, 51, 3,
	51, 3, 52, 3, 52, 3, 53, 3, 53, 7, 53, 564, 10, 53, 12, 53, 14, 53, 567,
	11, 53, 3, 53, 3, 53, 3, 53, 3, 53, 7, 53, 573, 10, 53, 12, 53, 14, 53,
	576, 11, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 5, 53, 585,
	10, 53, 3, 54, 3, 54, 7, 54, 589, 10, 54, 12, 54, 14, 54, 592, 11, 54,
	3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 599, 10, 55, 3, 56, 3, 56, 3,
	56, 3, 57, 3, 57, 3, 57, 3, 57, 5, 57, 608, 10, 57, 3, 57, 3, 57, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 5, 57, 620, 10, 57, 3,
	57, 3, 57, 5, 57, 624, 10, 57, 3, 58, 3, 58, 3, 58, 3, 58, 5, 58, 630,
	10, 58, 3, 59, 3, 59, 5, 59, 634, 10, 59, 3, 59, 7, 59, 637, 10, 59, 12,
	59, 14, 59, 640, 11, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 5, 61,
	648, 10, 61, 3, 62, 5, 62, 651, 10, 62, 3, 62, 3, 62, 5, 62, 655, 10, 62,
	3, 62, 3, 62, 3, 63, 3, 63, 6, 63, 661, 10, 63, 13, 63, 14, 63, 662, 3,
	64, 3, 64, 3, 64, 5, 64, 668, 10, 64, 3, 65, 3, 65, 3, 65, 6, 65, 673,
	10, 65, 13, 65, 14, 65, 674, 3, 65, 5, 65, 678, 10, 65, 3, 66, 3, 66, 6,
	66, 682, 10, 66, 13, 66, 14, 66, 683, 3, 67, 3, 67, 3, 67, 5, 67, 689,
	10, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 5, 68, 698, 10,
	68, 3, 69, 3, 69, 3, 69, 3, 69, 7, 69, 704, 10, 69, 12, 69, 14, 69, 707,
	11, 69, 3, 69, 5, 69, 710, 10, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 3,
	70, 3, 71, 3, 71, 3, 71, 5, 71, 721, 10, 71, 3, 72, 3, 72, 3, 72, 5, 72,
	726, 10, 72, 3, 72, 2, 2, 73, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96,
	98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126,
	128, 130, 132, 134, 136, 138, 140, 142, 2, 13, 4, 2, 22, 22, 31, 32, 4,
	2, 7, 7, 20, 20, 4, 2, 9, 9, 33, 34, 3, 2, 26, 30, 5, 2, 15, 15, 17, 17,
	36, 37, 3, 2, 56, 57, 3, 2, 60, 61, 4, 2, 15, 16, 26, 26, 4, 2, 66, 66,
	69, 69, 3, 2, 38, 40, 4, 2, 66, 66, 79, 79, 2, 781, 2, 147, 3, 2, 2, 2,
	4, 162, 3, 2, 2, 2, 6, 164, 3, 2, 2, 2, 8, 175, 3, 2, 2, 2, 10, 180, 3,
	2, 2, 2, 12, 185, 3, 2, 2, 2, 14, 191, 3, 2, 2, 2, 16, 205, 3, 2, 2, 2,
	18, 213, 3, 2, 2, 2, 20, 231, 3, 2, 2, 2, 22, 241, 3, 2, 2, 2, 24, 257,
	3, 2, 2, 2, 26, 276, 3, 2, 2, 2, 28, 281, 3, 2, 2, 2, 30, 289, 3, 2, 2,
	2, 32, 296, 3, 2, 2, 2, 34, 301, 3, 2, 2, 2, 36, 303, 3, 2, 2, 2, 38, 314,
	3, 2, 2, 2, 40, 316, 3, 2, 2, 2, 42, 328, 3, 2, 2, 2, 44, 333, 3, 2, 2,
	2, 46, 342, 3, 2, 2, 2, 48, 353, 3, 2, 2, 2, 50, 355, 3, 2, 2, 2, 52, 367,
	3, 2, 2, 2, 54, 375, 3, 2, 2, 2, 56, 379, 3, 2, 2, 2, 58, 381, 3, 2, 2,
	2, 60, 385, 3, 2, 2, 2, 62, 401, 3, 2, 2, 2, 64, 417, 3, 2, 2, 2, 66, 423,
	3, 2, 2, 2, 68, 425, 3, 2, 2, 2, 70, 430, 3, 2, 2, 2, 72, 445, 3, 2, 2,
	2, 74, 449, 3, 2, 2, 2, 76, 461, 3, 2, 2, 2, 78, 482, 3, 2, 2, 2, 80, 484,
	3, 2, 2, 2, 82, 486, 3, 2, 2, 2, 84, 494, 3, 2, 2, 2, 86, 496, 3, 2, 2,
	2, 88, 499, 3, 2, 2, 2, 90, 512, 3, 2, 2, 2, 92, 521, 3, 2, 2, 2, 94, 535,
	3, 2, 2, 2, 96, 537, 3, 2, 2, 2, 98, 548, 3, 2, 2, 2, 100, 550, 3, 2, 2,
	2, 102, 559, 3, 2, 2, 2, 104, 584, 3, 2, 2, 2, 106, 586, 3, 2, 2, 2, 108,
	598, 3, 2, 2, 2, 110, 600, 3, 2, 2, 2, 112, 623, 3, 2, 2, 2, 114, 625,
	3, 2, 2, 2, 116, 631, 3, 2, 2, 2, 118, 641, 3, 2, 2, 2, 120, 645, 3, 2,
	2, 2, 122, 650, 3, 2, 2, 2, 124, 660, 3, 2, 2, 2, 126, 667, 3, 2, 2, 2,
	128, 669, 3, 2, 2, 2, 130, 679, 3, 2, 2, 2, 132, 685, 3, 2, 2, 2, 134,
	697, 3, 2, 2, 2, 136, 699, 3, 2, 2, 2, 138, 713, 3, 2, 2, 2, 140, 720,
	3, 2, 2, 2, 142, 725, 3, 2, 2, 2, 144, 146, 5, 4, 3, 2, 145, 144, 3, 2,
	2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2,
	148, 3, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 163, 5, 64, 33, 2, 151,
	163, 5, 74, 38, 2, 152, 163, 5, 86, 44, 2, 153, 163, 5, 18, 10, 2, 154,
	163, 5, 20, 11, 2, 155, 163, 5, 24, 13, 2, 156, 163, 5, 50, 26, 2, 157,
	163, 5, 22, 12, 2, 158, 163, 5, 40, 21, 2, 159, 163, 5, 52, 27, 2, 160,
	163, 5, 58, 30, 2, 161, 163, 5, 60, 31, 2, 162, 150, 3, 2, 2, 2, 162, 151,
	3, 2, 2, 2, 162, 152, 3, 2, 2, 2, 162, 153, 3, 2, 2, 2, 162, 154, 3, 2,
	2, 2, 162, 155, 3, 2, 2, 2, 162, 156, 3, 2, 2, 2, 162, 157, 3, 2, 2, 2,
	162, 158, 3, 2, 2, 2, 162, 159, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162,
	161, 3, 2, 2, 2, 163, 5, 3, 2, 2, 2, 164, 169, 5, 8, 5, 2, 165, 166, 7,
	20, 2, 2, 166, 168, 5, 8, 5, 2, 167, 165, 3, 2, 2, 2, 168, 171, 3, 2, 2,
	2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2, 171,
	169, 3, 2, 2, 2, 172, 174, 7, 7, 2, 2, 173, 172, 3, 2, 2, 2, 173, 174,
	3, 2, 2, 2, 174, 7, 3, 2, 2, 2, 175, 177, 5, 10, 6, 2, 176, 178, 5, 12,
	7, 2, 177, 176, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 9, 3, 2, 2, 2, 179,
	181, 5, 124, 63, 2, 180, 179, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182,
	3, 2, 2, 2, 182, 183, 9, 2, 2, 2, 183, 184, 7, 66, 2, 2, 184, 11, 3, 2,
	2, 2, 185, 187, 7, 18, 2, 2, 186, 188, 5, 38, 20, 2, 187, 186, 3, 2, 2,
	2, 188, 189, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190,
	13, 3, 2, 2, 2, 191, 196, 5, 16, 9, 2, 192, 193, 7, 20, 2, 2, 193, 195,
	5, 16, 9, 2, 194, 192, 3, 2, 2, 2, 195, 198, 3, 2, 2, 2, 196, 194, 3, 2,
	2, 2, 196, 197, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2,
	199, 201, 9, 3, 2, 2, 200, 199, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201,
	15, 3, 2, 2, 2, 202, 203, 5, 10, 6, 2, 203, 204, 7, 18, 2, 2, 204, 206,
	3, 2, 2, 2, 205, 202, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 211, 3, 2,
	2, 2, 207, 212, 5, 32, 17, 2, 208, 212, 5, 130, 66, 2, 209, 212, 5, 132,
	67, 2, 210, 212, 5, 136, 69, 2, 211, 207, 3, 2, 2, 2, 211, 208, 3, 2, 2,
	2, 211, 209, 3, 2, 2, 2, 211, 210, 3, 2, 2, 2, 212, 17, 3, 2, 2, 2, 213,
	227, 7, 41, 2, 2, 214, 216, 7, 68, 2, 2, 215, 217, 5, 6, 4, 2, 216, 215,
	3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 228, 7, 10,
	2, 2, 219, 225, 7, 66, 2, 2, 220, 222, 7, 9, 2, 2, 221, 223, 5, 6, 4, 2,
	222, 221, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224,
	226, 7, 10, 2, 2, 225, 220, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 228,
	3, 2, 2, 2, 227, 214, 3, 2, 2, 2, 227, 219, 3, 2, 2, 2, 228, 229, 3, 2,
	2, 2, 229, 230, 5, 88, 45, 2, 230, 19, 3, 2, 2, 2, 231, 237, 7, 54, 2,
	2, 232, 234, 7, 9, 2, 2, 233, 235, 5, 14, 8, 2, 234, 233, 3, 2, 2, 2, 234,
	235, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 238, 7, 10, 2, 2, 237, 232,
	3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 240, 7, 19,
	2, 2, 240, 21, 3, 2, 2, 2, 241, 244, 7, 49, 2, 2, 242, 245, 7, 66, 2, 2,
	243, 245, 5, 122, 62, 2, 244, 242, 3, 2, 2, 2, 244, 243, 3, 2, 2, 2, 245,
	255, 3, 2, 2, 2, 246, 256, 7, 19, 2, 2, 247, 248, 7, 63, 2, 2, 248, 249,
	7, 9, 2, 2, 249, 250, 5, 6, 4, 2, 250, 251, 7, 10, 2, 2, 251, 253, 3, 2,
	2, 2, 252, 247, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2,
	254, 256, 5, 88, 45, 2, 255, 246, 3, 2, 2, 2, 255, 252, 3, 2, 2, 2, 255,
	256, 3, 2, 2, 2, 256, 23, 3, 2, 2, 2, 257, 261, 7, 42, 2, 2, 258, 262,
	7, 68, 2, 2, 259, 260, 7, 66, 2, 2, 260, 262, 7, 9, 2, 2, 261, 258, 3,
	2, 2, 2, 261, 259, 3, 2, 2, 2, 262, 264, 3, 2, 2, 2, 263, 265, 5, 6, 4,
	2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266,
	267, 7, 10, 2, 2, 267, 269, 7, 11, 2, 2, 268, 270, 5, 26, 14, 2, 269, 268,
	3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 272, 7, 12,
	2, 2, 272, 25, 3, 2, 2, 2, 273, 275, 5, 30, 16, 2, 274, 273, 3, 2, 2, 2,
	275, 278, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277,
	279, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 279, 280, 5, 28, 15, 2, 280, 27,
	3, 2, 2, 2, 281, 282, 7, 52, 2, 2, 282, 283, 5, 32, 17, 2, 283, 284, 7,
	19, 2, 2, 284, 29, 3, 2, 2, 2, 285, 286, 5, 32, 17, 2, 286, 287, 7, 19,
	2, 2, 287, 290, 3, 2, 2, 2, 288, 290, 5, 4, 3, 2, 289, 285, 3, 2, 2, 2,
	289, 288, 3, 2, 2, 2, 290, 31, 3, 2, 2, 2, 291, 297, 5, 38, 20, 2, 292,
	293, 9, 4, 2, 2, 293, 294, 5, 32, 17, 2, 294, 295, 7, 10, 2, 2, 295, 297,
	3, 2, 2, 2, 296, 291, 3, 2, 2, 2, 296, 292, 3, 2, 2, 2, 297, 299, 3, 2,
	2, 2, 298, 300, 5, 36, 19, 2, 299, 298, 3, 2, 2, 2, 299, 300, 3, 2, 2,
	2, 300, 33, 3, 2, 2, 2, 301, 302, 9, 5, 2, 2, 302, 35, 3, 2, 2, 2, 303,
	304, 5, 34, 18, 2, 304, 305, 5, 32, 17, 2, 305, 37, 3, 2, 2, 2, 306, 315,
	5, 120, 61, 2, 307, 315, 5, 104, 53, 2, 308, 315, 7, 71, 2, 2, 309, 315,
	7, 69, 2, 2, 310, 315, 7, 3, 2, 2, 311, 315, 5, 118, 60, 2, 312, 315, 5,
	10, 6, 2, 313, 315, 5, 122, 62, 2, 314, 306, 3, 2, 2, 2, 314, 307, 3, 2,
	2, 2, 314, 308, 3, 2, 2, 2, 314, 309, 3, 2, 2, 2, 314, 310, 3, 2, 2, 2,
	314, 311, 3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 314, 313, 3, 2, 2, 2, 315,
	39, 3, 2, 2, 2, 316, 317, 7, 45, 2, 2, 317, 318, 5, 46, 24, 2, 318, 322,
	5, 88, 45, 2, 319, 321, 5, 42, 22, 2, 320, 319, 3, 2, 2, 2, 321, 324, 3,
	2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 326, 3, 2, 2,
	2, 324, 322, 3, 2, 2, 2, 325, 327, 5, 44, 23, 2, 326, 325, 3, 2, 2, 2,
	326, 327, 3, 2, 2, 2, 327, 41, 3, 2, 2, 2, 328, 329, 7, 43, 2, 2, 329,
	330, 7, 44, 2, 2, 330, 331, 5, 46, 24, 2, 331, 332, 5, 88, 45, 2, 332,
	43, 3, 2, 2, 2, 333, 334, 7, 43, 2, 2, 334, 335, 5, 88, 45, 2, 335, 45,
	3, 2, 2, 2, 336, 339, 5, 48, 25, 2, 337, 338, 7, 6, 2, 2, 338, 340, 5,
	46, 24, 2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 343, 3, 2,
	2, 2, 341, 343, 7, 3, 2, 2, 342, 336, 3, 2, 2, 2, 342, 341, 3, 2, 2, 2,
	343, 47, 3, 2, 2, 2, 344, 347, 5, 32, 17, 2, 345, 346, 9, 6, 2, 2, 346,
	348, 5, 46, 24, 2, 347, 345, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 354,
	3, 2, 2, 2, 349, 350, 7, 9, 2, 2, 350, 351, 5, 46, 24, 2, 351, 352, 7,
	10, 2, 2, 352, 354, 3, 2, 2, 2, 353, 344, 3, 2, 2, 2, 353, 349, 3, 2, 2,
	2, 354, 49, 3, 2, 2, 2, 355, 356, 5, 10, 6, 2, 356, 360, 7, 18, 2, 2, 357,
	361, 5, 116, 59, 2, 358, 361, 5, 132, 67, 2, 359, 361, 5, 136, 69, 2, 360,
	357, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360, 359, 3, 2, 2, 2, 361, 363,
	3, 2, 2, 2, 362, 364, 7, 58, 2, 2, 363, 362, 3, 2, 2, 2, 363, 364, 3, 2,
	2, 2, 364, 365, 3, 2, 2, 2, 365, 366, 7, 19, 2, 2, 366, 51, 3, 2, 2, 2,
	367, 368, 7, 46, 2, 2, 368, 369, 5, 10, 6, 2, 369, 370, 7, 55, 2, 2, 370,
	371, 5, 54, 28, 2, 371, 372, 9, 7, 2, 2, 372, 373, 5, 56, 29, 2, 373, 374,
	5, 88, 45, 2, 374, 53, 3, 2, 2, 2, 375, 376, 7, 70, 2, 2, 376, 55, 3, 2,
	2, 2, 377, 380, 7, 70, 2, 2, 378, 380, 5, 122, 62, 2, 379, 377, 3, 2, 2,
	2, 379, 378, 3, 2, 2, 2, 380, 57, 3, 2, 2, 2, 381, 382, 7, 47, 2, 2, 382,
	383, 5, 46, 24, 2, 383, 384, 5, 88, 45, 2, 384, 59, 3, 2, 2, 2, 385, 386,
	7, 48, 2, 2, 386, 391, 5, 10, 6, 2, 387, 388, 7, 20, 2, 2, 388, 390, 5,
	10, 6, 2, 389, 387, 3, 2, 2, 2, 390, 393, 3, 2, 2, 2, 391, 389, 3, 2, 2,
	2, 391, 392, 3, 2, 2, 2, 392, 394, 3, 2, 2, 2, 393, 391, 3, 2, 2, 2, 394,
	395, 7, 4, 2, 2, 395, 396, 5, 62, 32, 2, 396, 397, 5, 88, 45, 2, 397, 61,
	3, 2, 2, 2, 398, 402, 5, 32, 17, 2, 399, 402, 5, 126, 64, 2, 400, 402,
	5, 136, 69, 2, 401, 398, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2, 401, 400, 3,
	2, 2, 2, 402, 63, 3, 2, 2, 2, 403, 404, 7, 50, 2, 2, 404, 405, 5, 66, 34,
	2, 405, 406, 7, 19, 2, 2, 406, 418, 3, 2, 2, 2, 407, 408, 7, 51, 2, 2,
	408, 410, 5, 66, 34, 2, 409, 411, 5, 68, 35, 2, 410, 409, 3, 2, 2, 2, 410,
	411, 3, 2, 2, 2, 411, 413, 3, 2, 2, 2, 412, 414, 5, 70, 36, 2, 413, 412,
	3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 416, 7, 19,
	2, 2, 416, 418, 3, 2, 2, 2, 417, 403, 3, 2, 2, 2, 417, 407, 3, 2, 2, 2,
	418, 65, 3, 2, 2, 2, 419, 424, 7, 69, 2, 2, 420, 421, 7, 35, 2, 2, 421,
	422, 7, 76, 2, 2, 422, 424, 7, 75, 2, 2, 423, 419, 3, 2, 2, 2, 423, 420,
	3, 2, 2, 2, 424, 67, 3, 2, 2, 2, 425, 428, 7, 64, 2, 2, 426, 429, 7, 27,
	2, 2, 427, 429, 5, 104, 53, 2, 428, 426, 3, 2, 2, 2, 428, 427, 3, 2, 2,
	2, 429, 69, 3, 2, 2, 2, 430, 431, 7, 65, 2, 2, 431, 432, 7, 9, 2, 2, 432,
	437, 5, 72, 37, 2, 433, 434, 7, 20, 2, 2, 434, 436, 5, 72, 37, 2, 435,
	433, 3, 2, 2, 2, 436, 439, 3, 2, 2, 2, 437, 435, 3, 2, 2, 2, 437, 438,
	3, 2, 2, 2, 438, 441, 3, 2, 2, 2, 439, 437, 3, 2, 2, 2, 440, 442, 7, 20,
	2, 2, 441, 440, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 443, 3, 2, 2, 2,
	443, 444, 7, 10, 2, 2, 444, 71, 3, 2, 2, 2, 445, 446, 5, 110, 56, 2, 446,
	447, 7, 18, 2, 2, 447, 448, 5, 38, 20, 2, 448, 73, 3, 2, 2, 2, 449, 450,
	7, 53, 2, 2, 450, 451, 5, 76, 39, 2, 451, 452, 5, 88, 45, 2, 452, 75, 3,
	2, 2, 2, 453, 458, 5, 78, 40, 2, 454, 455, 7, 20, 2, 2, 455, 457, 5, 78,
	40, 2, 456, 454, 3, 2, 2, 2, 457, 460, 3, 2, 2, 2, 458, 456, 3, 2, 2, 2,
	458, 459, 3, 2, 2, 2, 459, 462, 3, 2, 2, 2, 460, 458, 3, 2, 2, 2, 461,
	453, 3, 2, 2, 2, 461, 462, 3, 2, 2, 2, 462, 77, 3, 2, 2, 2, 463, 465, 9,
	8, 2, 2, 464, 463, 3, 2, 2, 2, 464, 465, 3, 2, 2, 2, 465, 466, 3, 2, 2,
	2, 466, 471, 5, 80, 41, 2, 467, 468, 7, 62, 2, 2, 468, 470, 5, 82, 42,
	2, 469, 467, 3, 2, 2, 2, 470, 473, 3, 2, 2, 2, 471, 469, 3, 2, 2, 2, 471,
	472, 3, 2, 2, 2, 472, 483, 3, 2, 2, 2, 473, 471, 3, 2, 2, 2, 474, 479,
	5, 82, 42, 2, 475, 476, 7, 62, 2, 2, 476, 478, 5, 82, 42, 2, 477, 475,
	3, 2, 2, 2, 478, 481, 3, 2, 2, 2, 479, 477, 3, 2, 2, 2, 479, 480, 3, 2,
	2, 2, 480, 483, 3, 2, 2, 2, 481, 479, 3, 2, 2, 2, 482, 464, 3, 2, 2, 2,
	482, 474, 3, 2, 2, 2, 483, 79, 3, 2, 2, 2, 484, 485, 7, 66, 2, 2, 485,
	81, 3, 2, 2, 2, 486, 487, 7, 9, 2, 2, 487, 490, 5, 84, 43, 2, 488, 489,
	7, 18, 2, 2, 489, 491, 5, 32, 17, 2, 490, 488, 3, 2, 2, 2, 490, 491, 3,
	2, 2, 2, 491, 492, 3, 2, 2, 2, 492, 493, 7, 10, 2, 2, 493, 83, 3, 2, 2,
	2, 494, 495, 7, 66, 2, 2, 495, 85, 3, 2, 2, 2, 496, 497, 5, 90, 46, 2,
	497, 498, 5, 88, 45, 2, 498, 87, 3, 2, 2, 2, 499, 504, 7, 11, 2, 2, 500,
	503, 5, 112, 57, 2, 501, 503, 5, 4, 3, 2, 502, 500, 3, 2, 2, 2, 502, 501,
	3, 2, 2, 2, 503, 506, 3, 2, 2, 2, 504, 502, 3, 2, 2, 2, 504, 505, 3, 2,
	2, 2, 505, 508, 3, 2, 2, 2, 506, 504, 3, 2, 2, 2, 507, 509, 5, 114, 58,
	2, 508, 507, 3, 2, 2, 2, 508, 509, 3, 2, 2, 2, 509, 510, 3, 2, 2, 2, 510,
	511, 7, 12, 2, 2, 511, 89, 3, 2, 2, 2, 512, 517, 5, 92, 47, 2, 513, 514,
	7, 20, 2, 2, 514, 516, 5, 92, 47, 2, 515, 513, 3, 2, 2, 2, 516, 519, 3,
	2, 2, 2, 517, 515, 3, 2, 2, 2, 517, 518, 3, 2, 2, 2, 518, 91, 3, 2, 2,
	2, 519, 517, 3, 2, 2, 2, 520, 522, 5, 94, 48, 2, 521, 520, 3, 2, 2, 2,
	522, 523, 3, 2, 2, 2, 523, 521, 3, 2, 2, 2, 523, 524, 3, 2, 2, 2, 524,
	93, 3, 2, 2, 2, 525, 536, 5, 104, 53, 2, 526, 527, 7, 25, 2, 2, 527, 536,
	5, 104, 53, 2, 528, 529, 7, 21, 2, 2, 529, 536, 5, 104, 53, 2, 530, 536,
	7, 24, 2, 2, 531, 536, 7, 27, 2, 2, 532, 536, 5, 96, 49, 2, 533, 536, 5,
	100, 51, 2, 534, 536, 5, 98, 50, 2, 535, 525, 3, 2, 2, 2, 535, 526, 3,
	2, 2, 2, 535, 528, 3, 2, 2, 2, 535, 530, 3, 2, 2, 2, 535, 531, 3, 2, 2,
	2, 535, 532, 3, 2, 2, 2, 535, 533, 3, 2, 2, 2, 535, 534, 3, 2, 2, 2, 536,
	95, 3, 2, 2, 2, 537, 538, 9, 9, 2, 2, 538, 97, 3, 2, 2, 2, 539, 549, 5,
	106, 54, 2, 540, 541, 5, 106, 54, 2, 541, 544, 7, 9, 2, 2, 542, 545, 5,
	92, 47, 2, 543, 545, 5, 32, 17, 2, 544, 542, 3, 2, 2, 2, 544, 543, 3, 2,
	2, 2, 545, 546, 3, 2, 2, 2, 546, 547, 7, 10, 2, 2, 547, 549, 3, 2, 2, 2,
	548, 539, 3, 2, 2, 2, 548, 540, 3, 2, 2, 2, 549, 99, 3, 2, 2, 2, 550, 551,
	7, 13, 2, 2, 551, 555, 7, 66, 2, 2, 552, 553, 5, 102, 52, 2, 553, 554,
	9, 10, 2, 2, 554, 556, 3, 2, 2, 2, 555, 552, 3, 2, 2, 2, 555, 556, 3, 2,
	2, 2, 556, 557, 3, 2, 2, 2, 557, 558, 7, 14, 2, 2, 558, 101, 3, 2, 2, 2,
	559, 560, 9, 11, 2, 2, 560, 103, 3, 2, 2, 2, 561, 565, 7, 66, 2, 2, 562,
	564, 5, 108, 55, 2, 563, 562, 3, 2, 2, 2, 564, 567, 3, 2, 2, 2, 565, 563,
	3, 2, 2, 2, 565, 566, 3, 2, 2, 2, 566, 585, 3, 2, 2, 2, 567, 565, 3, 2,
	2, 2, 568, 569, 7, 8, 2, 2, 569, 570, 5, 110, 56, 2, 570, 574, 7, 12, 2,
	2, 571, 573, 5, 108, 55, 2, 572, 571, 3, 2, 2, 2, 573, 576, 3, 2, 2, 2,
	574, 572, 3, 2, 2, 2, 574, 575, 3, 2, 2, 2, 575, 585, 3, 2, 2, 2, 576,
	574, 3, 2, 2, 2, 577, 585, 7, 62, 2, 2, 578, 585, 7, 55, 2, 2, 579, 585,
	7, 61, 2, 2, 580, 585, 7, 60, 2, 2, 581, 585, 7, 57, 2, 2, 582, 585, 7,
	56, 2, 2, 583, 585, 7, 63, 2, 2, 584, 561, 3, 2, 2, 2, 584, 568, 3, 2,
	2, 2, 584, 577, 3, 2, 2, 2, 584, 578, 3, 2, 2, 2, 584, 579, 3, 2, 2, 2,
	584, 580, 3, 2, 2, 2, 584, 581, 3, 2, 2, 2, 584, 582, 3, 2, 2, 2, 584,
	583, 3, 2, 2, 2, 585, 105, 3, 2, 2, 2, 586, 590, 7, 67, 2, 2, 587, 589,
	5, 108, 55, 2, 588, 587, 3, 2, 2, 2, 589, 592, 3, 2, 2, 2, 590, 588, 3,
	2, 2, 2, 590, 591, 3, 2, 2, 2, 591, 107, 3, 2, 2, 2, 592, 590, 3, 2, 2,
	2, 593, 594, 7, 78, 2, 2, 594, 595, 5, 110, 56, 2, 595, 596, 7, 12, 2,
	2, 596, 599, 3, 2, 2, 2, 597, 599, 7, 79, 2, 2, 598, 593, 3, 2, 2, 2, 598,
	597, 3, 2, 2, 2, 599, 109, 3, 2, 2, 2, 600, 601, 7, 22, 2, 2, 601, 602,
	9, 12, 2, 2, 602, 111, 3, 2, 2, 2, 603, 604, 5, 104, 53, 2, 604, 605, 7,
	18, 2, 2, 605, 607, 5, 116, 59, 2, 606, 608, 7, 59, 2, 2, 607, 606, 3,
	2, 2, 2, 607, 608, 3, 2, 2, 2, 608, 609, 3, 2, 2, 2, 609, 610, 7, 19, 2,
	2, 610, 624, 3, 2, 2, 2, 611, 612, 5, 104, 53, 2, 612, 613, 7, 18, 2, 2,
	613, 614, 5, 88, 45, 2, 614, 624, 3, 2, 2, 2, 615, 616, 5, 104, 53, 2,
	616, 617, 7, 18, 2, 2, 617, 619, 5, 116, 59, 2, 618, 620, 7, 59, 2, 2,
	619, 618, 3, 2, 2, 2, 619, 620, 3, 2, 2, 2, 620, 621, 3, 2, 2, 2, 621,
	622, 5, 88, 45, 2, 622, 624, 3, 2, 2, 2, 623, 603, 3, 2, 2, 2, 623, 611,
	3, 2, 2, 2, 623, 615, 3, 2, 2, 2, 624, 113, 3, 2, 2, 2, 625, 626, 5, 104,
	53, 2, 626, 627, 7, 18, 2, 2, 627, 629, 5, 116, 59, 2, 628, 630, 7, 59,
	2, 2, 629, 628, 3, 2, 2, 2, 629, 630, 3, 2, 2, 2, 630, 115, 3, 2, 2, 2,
	631, 638, 5, 32, 17, 2, 632, 634, 7, 20, 2, 2, 633, 632, 3, 2, 2, 2, 633,
	634, 3, 2, 2, 2, 634, 635, 3, 2, 2, 2, 635, 637, 5, 32, 17, 2, 636, 633,
	3, 2, 2, 2, 637, 640, 3, 2, 2, 2, 638, 636, 3, 2, 2, 2, 638, 639, 3, 2,
	2, 2, 639, 117, 3, 2, 2, 2, 640, 638, 3, 2, 2, 2, 641, 642, 7, 35, 2, 2,
	642, 643, 7, 76, 2, 2, 643, 644, 7, 75, 2, 2, 644, 119, 3, 2, 2, 2, 645,
	647, 7, 70, 2, 2, 646, 648, 7, 5, 2, 2, 647, 646, 3, 2, 2, 2, 647, 648,
	3, 2, 2, 2, 648, 121, 3, 2, 2, 2, 649, 651, 5, 124, 63, 2, 650, 649, 3,
	2, 2, 2, 650, 651, 3, 2, 2, 2, 651, 652, 3, 2, 2, 2, 652, 654, 7, 68, 2,
	2, 653, 655, 5, 14, 8, 2, 654, 653, 3, 2, 2, 2, 654, 655, 3, 2, 2, 2, 655,
	656, 3, 2, 2, 2, 656, 657, 7, 10, 2, 2, 657, 123, 3, 2, 2, 2, 658, 659,
	7, 66, 2, 2, 659, 661, 7, 21, 2, 2, 660, 658, 3, 2, 2, 2, 661, 662, 3,
	2, 2, 2, 662, 660, 3, 2, 2, 2, 662, 663, 3, 2, 2, 2, 663, 125, 3, 2, 2,
	2, 664, 668, 5, 128, 65, 2, 665, 668, 5, 130, 66, 2, 666, 668, 5, 132,
	67, 2, 667, 664, 3, 2, 2, 2, 667, 665, 3, 2, 2, 2, 667, 666, 3, 2, 2, 2,
	668, 127, 3, 2, 2, 2, 669, 672, 5, 134, 68, 2, 670, 671, 7, 20, 2, 2, 671,
	673, 5, 134, 68, 2, 672, 670, 3, 2, 2, 2, 673, 674, 3, 2, 2, 2, 674, 672,
	3, 2, 2, 2, 674, 675, 3, 2, 2, 2, 675, 677, 3, 2, 2, 2, 676, 678, 7, 20,
	2, 2, 677, 676, 3, 2, 2, 2, 677, 678, 3, 2, 2, 2, 678, 129, 3, 2, 2, 2,
	679, 681, 5, 134, 68, 2, 680, 682, 5, 134, 68, 2, 681, 680, 3, 2, 2, 2,
	682, 683, 3, 2, 2, 2, 683, 681, 3, 2, 2, 2, 683, 684, 3, 2, 2, 2, 684,
	131, 3, 2, 2, 2, 685, 688, 7, 13, 2, 2, 686, 689, 5, 128, 65, 2, 687, 689,
	5, 130, 66, 2, 688, 686, 3, 2, 2, 2, 688, 687, 3, 2, 2, 2, 689, 690, 3,
	2, 2, 2, 690, 691, 7, 14, 2, 2, 691, 133, 3, 2, 2, 2, 692, 698, 5, 32,
	17, 2, 693, 694, 7, 9, 2, 2, 694, 695, 5, 126, 64, 2, 695, 696, 7, 10,
	2, 2, 696, 698, 3, 2, 2, 2, 697, 692, 3, 2, 2, 2, 697, 693, 3, 2, 2, 2,
	698, 135, 3, 2, 2, 2, 699, 700, 7, 9, 2, 2, 700, 705, 5, 138, 70, 2, 701,
	702, 7, 20, 2, 2, 702, 704, 5, 138, 70, 2, 703, 701, 3, 2, 2, 2, 704, 707,
	3, 2, 2, 2, 705, 703, 3, 2, 2, 2, 705, 706, 3, 2, 2, 2, 706, 709, 3, 2,
	2, 2, 707, 705, 3, 2, 2, 2, 708, 710, 7, 20, 2, 2, 709, 708, 3, 2, 2, 2,
	709, 710, 3, 2, 2, 2, 710, 711, 3, 2, 2, 2, 711, 712, 7, 10, 2, 2, 712,
	137, 3, 2, 2, 2, 713, 714, 5, 140, 71, 2, 714, 715, 7, 18, 2, 2, 715, 716,
	5, 142, 72, 2, 716, 139, 3, 2, 2, 2, 717, 721, 5, 32, 17, 2, 718, 721,
	5, 126, 64, 2, 719, 721, 5, 136, 69, 2, 720, 717, 3, 2, 2, 2, 720, 718,
	3, 2, 2, 2, 720, 719, 3, 2, 2, 2, 721, 141, 3, 2, 2, 2, 722, 726, 5, 32,
	17, 2, 723, 726, 5, 126, 64, 2, 724, 726, 5, 136, 69, 2, 725, 722, 3, 2,
	2, 2, 725, 723, 3, 2, 2, 2, 725, 724, 3, 2, 2, 2, 726, 143, 3, 2, 2, 2,
	89, 147, 162, 169, 173, 177, 180, 189, 196, 200, 205, 211, 216, 222, 225,
	227, 234, 237, 244, 252, 255, 261, 264, 269, 276, 289, 296, 299, 314, 322,
	326, 339, 342, 347, 353, 360, 363, 379, 391, 401, 410, 413, 417, 423, 428,
	437, 441, 458, 461, 464, 471, 479, 482, 490, 502, 504, 508, 517, 523, 535,
	544, 548, 555, 565, 574, 584, 590, 598, 607, 619, 623, 629, 633, 638, 647,
	650, 654, 662, 667, 674, 677, 683, 688, 697, 705, 709, 720, 725,
}
var literalNames = []string{
	"", "'null'", "'in'", "", "", "'...'", "", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "'>'", "'~'", "'<'", "':'", "';'", "','", "'.'", "'$'", "'@'",
	"'&'", "'#'", "'+'", "'*'", "'/'", "'-'", "'%'", "", "", "", "", "", "'=='",
	"'!='", "'='", "'|='", "'~='", "'@mixin'", "'@function'", "'@else'", "'if'",
	"'@if'", "'@for'", "'@while'", "'@each'", "'@include'", "'@import'", "'@use'",
	"'@return'", "'@media'", "'@content'", "'from'", "'to'", "'through'", "'!default'",
	"'!important'", "'only'", "'not'", "'and'", "'using'", "'as'", "'with'",
}
var symbolicNames = []string{
	"", "NULL_", "IN", "Unit", "COMBINE_COMPARE", "Ellipsis", "InterpolationStart",
	"LPAREN", "RPAREN", "BlockStart", "BlockEnd", "LBRACK", "RBRACK", "GT",
	"TIL", "LT", "COLON", "SEMI", "COMMA", "DOT", "DOLLAR", "AT", "AND", "HASH",
	"PLUS", "TIMES", "DIV", "MINUS", "PERC", "MINUS_DOLLAR", "PLUS_DOLLAR",
	"MINUS_LPAREN", "PLUS_LPAREN", "UrlStart", "EQEQ", "NOTEQ", "EQ", "PIPE_EQ",
	"TILD_EQ", "MIXIN", "FUNCTION", "AT_ELSE", "IF", "AT_IF", "AT_FOR", "AT_WHILE",
	"AT_EACH", "INCLUDE", "IMPORT", "USE", "RETURN", "MEDIA", "CONTENT", "FROM",
	"TO", "THROUGH", "POUND_DEFAULT", "IMPORTANT", "ONLY", "NOT", "AND_WORD",
	"USING", "AS", "WITH", "Identifier", "PseudoIdentifier", "FunctionIdentifier",
	"StringLiteral", "Number", "Color", "WS", "SL_COMMENT", "COMMENT", "UrlEnd",
	"Url", "SPACE", "InterpolationStartAfter", "IdentifierAfter",
}

var ruleNames = []string{
	"stylesheet", "statement", "declaredParams", "declaredParam", "variableName",
	"paramOptionalValue", "passedParams", "passedParam", "mixinDeclaration",
	"contentDeclaration", "includeDeclaration", "functionDeclaration", "functionBody",
	"functionReturn", "functionStatement", "commandStatement", "mathCharacter",
	"mathStatement", "expression", "ifDeclaration", "elseIfStatement", "elseStatement",
	"conditions", "condition", "variableDeclaration", "forDeclaration", "fromNumber",
	"throughNumber", "whileDeclaration", "eachDeclaration", "eachValueList",
	"importDeclaration", "referenceUrl", "asClause", "withClause", "keywordArgument",
	"mediaDeclaration", "mediaQueryList", "mediaQuery", "mediaType", "mediaExpression",
	"mediaFeature", "ruleset", "block", "selectors", "selector", "element",
	"combinator", "pseudo", "attrib", "attribRelate", "identifier", "pseudoIdentifier",
	"identifierPart", "identifierVariableName", "property_", "lastProperty",
	"propertyValue", "url", "measurement", "functionCall", "namespace", "list_",
	"listCommaSeparated", "listSpaceSeparated", "listBracketed", "listElement",
	"map_", "mapEntry", "mapKey", "mapValue",
}

type ScssParser struct {
	*antlr.BaseParser
}

// NewScssParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ScssParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewScssParser(input antlr.TokenStream) *ScssParser {
	this := new(ScssParser)
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
	this.GrammarFileName = "ScssParser.g4"

	return this
}

// ScssParser tokens.
const (
	ScssParserEOF                     = antlr.TokenEOF
	ScssParserNULL_                   = 1
	ScssParserIN                      = 2
	ScssParserUnit                    = 3
	ScssParserCOMBINE_COMPARE         = 4
	ScssParserEllipsis                = 5
	ScssParserInterpolationStart      = 6
	ScssParserLPAREN                  = 7
	ScssParserRPAREN                  = 8
	ScssParserBlockStart              = 9
	ScssParserBlockEnd                = 10
	ScssParserLBRACK                  = 11
	ScssParserRBRACK                  = 12
	ScssParserGT                      = 13
	ScssParserTIL                     = 14
	ScssParserLT                      = 15
	ScssParserCOLON                   = 16
	ScssParserSEMI                    = 17
	ScssParserCOMMA                   = 18
	ScssParserDOT                     = 19
	ScssParserDOLLAR                  = 20
	ScssParserAT                      = 21
	ScssParserAND                     = 22
	ScssParserHASH                    = 23
	ScssParserPLUS                    = 24
	ScssParserTIMES                   = 25
	ScssParserDIV                     = 26
	ScssParserMINUS                   = 27
	ScssParserPERC                    = 28
	ScssParserMINUS_DOLLAR            = 29
	ScssParserPLUS_DOLLAR             = 30
	ScssParserMINUS_LPAREN            = 31
	ScssParserPLUS_LPAREN             = 32
	ScssParserUrlStart                = 33
	ScssParserEQEQ                    = 34
	ScssParserNOTEQ                   = 35
	ScssParserEQ                      = 36
	ScssParserPIPE_EQ                 = 37
	ScssParserTILD_EQ                 = 38
	ScssParserMIXIN                   = 39
	ScssParserFUNCTION                = 40
	ScssParserAT_ELSE                 = 41
	ScssParserIF                      = 42
	ScssParserAT_IF                   = 43
	ScssParserAT_FOR                  = 44
	ScssParserAT_WHILE                = 45
	ScssParserAT_EACH                 = 46
	ScssParserINCLUDE                 = 47
	ScssParserIMPORT                  = 48
	ScssParserUSE                     = 49
	ScssParserRETURN                  = 50
	ScssParserMEDIA                   = 51
	ScssParserCONTENT                 = 52
	ScssParserFROM                    = 53
	ScssParserTO                      = 54
	ScssParserTHROUGH                 = 55
	ScssParserPOUND_DEFAULT           = 56
	ScssParserIMPORTANT               = 57
	ScssParserONLY                    = 58
	ScssParserNOT                     = 59
	ScssParserAND_WORD                = 60
	ScssParserUSING                   = 61
	ScssParserAS                      = 62
	ScssParserWITH                    = 63
	ScssParserIdentifier              = 64
	ScssParserPseudoIdentifier        = 65
	ScssParserFunctionIdentifier      = 66
	ScssParserStringLiteral           = 67
	ScssParserNumber                  = 68
	ScssParserColor                   = 69
	ScssParserWS                      = 70
	ScssParserSL_COMMENT              = 71
	ScssParserCOMMENT                 = 72
	ScssParserUrlEnd                  = 73
	ScssParserUrl                     = 74
	ScssParserSPACE                   = 75
	ScssParserInterpolationStartAfter = 76
	ScssParserIdentifierAfter         = 77
)

// ScssParser rules.
const (
	ScssParserRULE_stylesheet             = 0
	ScssParserRULE_statement              = 1
	ScssParserRULE_declaredParams         = 2
	ScssParserRULE_declaredParam          = 3
	ScssParserRULE_variableName           = 4
	ScssParserRULE_paramOptionalValue     = 5
	ScssParserRULE_passedParams           = 6
	ScssParserRULE_passedParam            = 7
	ScssParserRULE_mixinDeclaration       = 8
	ScssParserRULE_contentDeclaration     = 9
	ScssParserRULE_includeDeclaration     = 10
	ScssParserRULE_functionDeclaration    = 11
	ScssParserRULE_functionBody           = 12
	ScssParserRULE_functionReturn         = 13
	ScssParserRULE_functionStatement      = 14
	ScssParserRULE_commandStatement       = 15
	ScssParserRULE_mathCharacter          = 16
	ScssParserRULE_mathStatement          = 17
	ScssParserRULE_expression             = 18
	ScssParserRULE_ifDeclaration          = 19
	ScssParserRULE_elseIfStatement        = 20
	ScssParserRULE_elseStatement          = 21
	ScssParserRULE_conditions             = 22
	ScssParserRULE_condition              = 23
	ScssParserRULE_variableDeclaration    = 24
	ScssParserRULE_forDeclaration         = 25
	ScssParserRULE_fromNumber             = 26
	ScssParserRULE_throughNumber          = 27
	ScssParserRULE_whileDeclaration       = 28
	ScssParserRULE_eachDeclaration        = 29
	ScssParserRULE_eachValueList          = 30
	ScssParserRULE_importDeclaration      = 31
	ScssParserRULE_referenceUrl           = 32
	ScssParserRULE_asClause               = 33
	ScssParserRULE_withClause             = 34
	ScssParserRULE_keywordArgument        = 35
	ScssParserRULE_mediaDeclaration       = 36
	ScssParserRULE_mediaQueryList         = 37
	ScssParserRULE_mediaQuery             = 38
	ScssParserRULE_mediaType              = 39
	ScssParserRULE_mediaExpression        = 40
	ScssParserRULE_mediaFeature           = 41
	ScssParserRULE_ruleset                = 42
	ScssParserRULE_block                  = 43
	ScssParserRULE_selectors              = 44
	ScssParserRULE_selector               = 45
	ScssParserRULE_element                = 46
	ScssParserRULE_combinator             = 47
	ScssParserRULE_pseudo                 = 48
	ScssParserRULE_attrib                 = 49
	ScssParserRULE_attribRelate           = 50
	ScssParserRULE_identifier             = 51
	ScssParserRULE_pseudoIdentifier       = 52
	ScssParserRULE_identifierPart         = 53
	ScssParserRULE_identifierVariableName = 54
	ScssParserRULE_property_              = 55
	ScssParserRULE_lastProperty           = 56
	ScssParserRULE_propertyValue          = 57
	ScssParserRULE_url                    = 58
	ScssParserRULE_measurement            = 59
	ScssParserRULE_functionCall           = 60
	ScssParserRULE_namespace              = 61
	ScssParserRULE_list_                  = 62
	ScssParserRULE_listCommaSeparated     = 63
	ScssParserRULE_listSpaceSeparated     = 64
	ScssParserRULE_listBracketed          = 65
	ScssParserRULE_listElement            = 66
	ScssParserRULE_map_                   = 67
	ScssParserRULE_mapEntry               = 68
	ScssParserRULE_mapKey                 = 69
	ScssParserRULE_mapValue               = 70
)

// IStylesheetContext is an interface to support dynamic dispatch.
type IStylesheetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStylesheetContext differentiates from other interfaces.
	IsStylesheetContext()
}

type StylesheetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStylesheetContext() *StylesheetContext {
	var p = new(StylesheetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_stylesheet
	return p
}

func (*StylesheetContext) IsStylesheetContext() {}

func NewStylesheetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StylesheetContext {
	var p = new(StylesheetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_stylesheet

	return p
}

func (s *StylesheetContext) GetParser() antlr.Parser { return s.parser }

func (s *StylesheetContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StylesheetContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StylesheetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StylesheetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StylesheetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterStylesheet(s)
	}
}

func (s *StylesheetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitStylesheet(s)
	}
}

func (p *ScssParser) Stylesheet() (localctx IStylesheetContext) {
	this := p
	_ = this

	localctx = NewStylesheetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ScssParserRULE_stylesheet)
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
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserInterpolationStart)|(1<<ScssParserLBRACK)|(1<<ScssParserGT)|(1<<ScssParserTIL)|(1<<ScssParserDOT)|(1<<ScssParserDOLLAR)|(1<<ScssParserAND)|(1<<ScssParserHASH)|(1<<ScssParserPLUS)|(1<<ScssParserTIMES)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(ScssParserMIXIN-39))|(1<<(ScssParserFUNCTION-39))|(1<<(ScssParserAT_IF-39))|(1<<(ScssParserAT_FOR-39))|(1<<(ScssParserAT_WHILE-39))|(1<<(ScssParserAT_EACH-39))|(1<<(ScssParserINCLUDE-39))|(1<<(ScssParserIMPORT-39))|(1<<(ScssParserUSE-39))|(1<<(ScssParserMEDIA-39))|(1<<(ScssParserCONTENT-39))|(1<<(ScssParserFROM-39))|(1<<(ScssParserTO-39))|(1<<(ScssParserTHROUGH-39))|(1<<(ScssParserONLY-39))|(1<<(ScssParserNOT-39))|(1<<(ScssParserAND_WORD-39))|(1<<(ScssParserUSING-39))|(1<<(ScssParserIdentifier-39))|(1<<(ScssParserPseudoIdentifier-39)))) != 0) {
		{
			p.SetState(142)
			p.Statement()
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ImportDeclaration() IImportDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *StatementContext) MediaDeclaration() IMediaDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMediaDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMediaDeclarationContext)
}

func (s *StatementContext) Ruleset() IRulesetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRulesetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRulesetContext)
}

func (s *StatementContext) MixinDeclaration() IMixinDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMixinDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMixinDeclarationContext)
}

func (s *StatementContext) ContentDeclaration() IContentDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentDeclarationContext)
}

func (s *StatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) IncludeDeclaration() IIncludeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncludeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncludeDeclarationContext)
}

func (s *StatementContext) IfDeclaration() IIfDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfDeclarationContext)
}

func (s *StatementContext) ForDeclaration() IForDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForDeclarationContext)
}

func (s *StatementContext) WhileDeclaration() IWhileDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileDeclarationContext)
}

func (s *StatementContext) EachDeclaration() IEachDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEachDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEachDeclarationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *ScssParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ScssParserRULE_statement)

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

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.ImportDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.MediaDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)
			p.Ruleset()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(151)
			p.MixinDeclaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(152)
			p.ContentDeclaration()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(153)
			p.FunctionDeclaration()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(154)
			p.VariableDeclaration()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(155)
			p.IncludeDeclaration()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(156)
			p.IfDeclaration()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(157)
			p.ForDeclaration()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(158)
			p.WhileDeclaration()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(159)
			p.EachDeclaration()
		}

	}

	return localctx
}

// IDeclaredParamsContext is an interface to support dynamic dispatch.
type IDeclaredParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclaredParamsContext differentiates from other interfaces.
	IsDeclaredParamsContext()
}

type DeclaredParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaredParamsContext() *DeclaredParamsContext {
	var p = new(DeclaredParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_declaredParams
	return p
}

func (*DeclaredParamsContext) IsDeclaredParamsContext() {}

func NewDeclaredParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaredParamsContext {
	var p = new(DeclaredParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_declaredParams

	return p
}

func (s *DeclaredParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaredParamsContext) AllDeclaredParam() []IDeclaredParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclaredParamContext)(nil)).Elem())
	var tst = make([]IDeclaredParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclaredParamContext)
		}
	}

	return tst
}

func (s *DeclaredParamsContext) DeclaredParam(i int) IDeclaredParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaredParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclaredParamContext)
}

func (s *DeclaredParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *DeclaredParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *DeclaredParamsContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(ScssParserEllipsis, 0)
}

func (s *DeclaredParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaredParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaredParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterDeclaredParams(s)
	}
}

func (s *DeclaredParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitDeclaredParams(s)
	}
}

func (p *ScssParser) DeclaredParams() (localctx IDeclaredParamsContext) {
	this := p
	_ = this

	localctx = NewDeclaredParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ScssParserRULE_declaredParams)
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
		p.SetState(162)
		p.DeclaredParam()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserCOMMA {
		{
			p.SetState(163)
			p.Match(ScssParserCOMMA)
		}
		{
			p.SetState(164)
			p.DeclaredParam()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserEllipsis {
		{
			p.SetState(170)
			p.Match(ScssParserEllipsis)
		}

	}

	return localctx
}

// IDeclaredParamContext is an interface to support dynamic dispatch.
type IDeclaredParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclaredParamContext differentiates from other interfaces.
	IsDeclaredParamContext()
}

type DeclaredParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaredParamContext() *DeclaredParamContext {
	var p = new(DeclaredParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_declaredParam
	return p
}

func (*DeclaredParamContext) IsDeclaredParamContext() {}

func NewDeclaredParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaredParamContext {
	var p = new(DeclaredParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_declaredParam

	return p
}

func (s *DeclaredParamContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaredParamContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *DeclaredParamContext) ParamOptionalValue() IParamOptionalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamOptionalValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamOptionalValueContext)
}

func (s *DeclaredParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaredParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaredParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterDeclaredParam(s)
	}
}

func (s *DeclaredParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitDeclaredParam(s)
	}
}

func (p *ScssParser) DeclaredParam() (localctx IDeclaredParamContext) {
	this := p
	_ = this

	localctx = NewDeclaredParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ScssParserRULE_declaredParam)
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
		p.SetState(173)
		p.VariableName()
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserCOLON {
		{
			p.SetState(174)
			p.ParamOptionalValue()
		}

	}

	return localctx
}

// IVariableNameContext is an interface to support dynamic dispatch.
type IVariableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableNameContext differentiates from other interfaces.
	IsVariableNameContext()
}

type VariableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableNameContext() *VariableNameContext {
	var p = new(VariableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_variableName
	return p
}

func (*VariableNameContext) IsVariableNameContext() {}

func NewVariableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableNameContext {
	var p = new(VariableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_variableName

	return p
}

func (s *VariableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *VariableNameContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(ScssParserDOLLAR, 0)
}

func (s *VariableNameContext) MINUS_DOLLAR() antlr.TerminalNode {
	return s.GetToken(ScssParserMINUS_DOLLAR, 0)
}

func (s *VariableNameContext) PLUS_DOLLAR() antlr.TerminalNode {
	return s.GetToken(ScssParserPLUS_DOLLAR, 0)
}

func (s *VariableNameContext) Namespace() INamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceContext)
}

func (s *VariableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterVariableName(s)
	}
}

func (s *VariableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitVariableName(s)
	}
}

func (p *ScssParser) VariableName() (localctx IVariableNameContext) {
	this := p
	_ = this

	localctx = NewVariableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ScssParserRULE_variableName)
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
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserIdentifier {
		{
			p.SetState(177)
			p.Namespace()
		}

	}
	{
		p.SetState(180)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserDOLLAR)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(181)
		p.Match(ScssParserIdentifier)
	}

	return localctx
}

// IParamOptionalValueContext is an interface to support dynamic dispatch.
type IParamOptionalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamOptionalValueContext differentiates from other interfaces.
	IsParamOptionalValueContext()
}

type ParamOptionalValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamOptionalValueContext() *ParamOptionalValueContext {
	var p = new(ParamOptionalValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_paramOptionalValue
	return p
}

func (*ParamOptionalValueContext) IsParamOptionalValueContext() {}

func NewParamOptionalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamOptionalValueContext {
	var p = new(ParamOptionalValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_paramOptionalValue

	return p
}

func (s *ParamOptionalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamOptionalValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *ParamOptionalValueContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ParamOptionalValueContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParamOptionalValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamOptionalValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamOptionalValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterParamOptionalValue(s)
	}
}

func (s *ParamOptionalValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitParamOptionalValue(s)
	}
}

func (p *ScssParser) ParamOptionalValue() (localctx IParamOptionalValueContext) {
	this := p
	_ = this

	localctx = NewParamOptionalValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ScssParserRULE_paramOptionalValue)
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
		p.SetState(183)
		p.Match(ScssParserCOLON)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL_)|(1<<ScssParserInterpolationStart)|(1<<ScssParserDOLLAR)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ScssParserUrlStart-33))|(1<<(ScssParserFROM-33))|(1<<(ScssParserTO-33))|(1<<(ScssParserTHROUGH-33))|(1<<(ScssParserONLY-33))|(1<<(ScssParserNOT-33))|(1<<(ScssParserAND_WORD-33))|(1<<(ScssParserUSING-33))|(1<<(ScssParserIdentifier-33)))) != 0) || (((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(ScssParserFunctionIdentifier-66))|(1<<(ScssParserStringLiteral-66))|(1<<(ScssParserNumber-66))|(1<<(ScssParserColor-66)))) != 0) {
		{
			p.SetState(184)
			p.Expression()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPassedParamsContext is an interface to support dynamic dispatch.
type IPassedParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassedParamsContext differentiates from other interfaces.
	IsPassedParamsContext()
}

type PassedParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassedParamsContext() *PassedParamsContext {
	var p = new(PassedParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_passedParams
	return p
}

func (*PassedParamsContext) IsPassedParamsContext() {}

func NewPassedParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassedParamsContext {
	var p = new(PassedParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_passedParams

	return p
}

func (s *PassedParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *PassedParamsContext) AllPassedParam() []IPassedParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPassedParamContext)(nil)).Elem())
	var tst = make([]IPassedParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPassedParamContext)
		}
	}

	return tst
}

func (s *PassedParamsContext) PassedParam(i int) IPassedParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassedParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPassedParamContext)
}

func (s *PassedParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *PassedParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *PassedParamsContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(ScssParserEllipsis, 0)
}

func (s *PassedParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassedParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassedParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterPassedParams(s)
	}
}

func (s *PassedParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitPassedParams(s)
	}
}

func (p *ScssParser) PassedParams() (localctx IPassedParamsContext) {
	this := p
	_ = this

	localctx = NewPassedParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ScssParserRULE_passedParams)
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
		p.SetState(189)
		p.PassedParam()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(190)
				p.Match(ScssParserCOMMA)
			}
			{
				p.SetState(191)
				p.PassedParam()
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserEllipsis || _la == ScssParserCOMMA {
		{
			p.SetState(197)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ScssParserEllipsis || _la == ScssParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IPassedParamContext is an interface to support dynamic dispatch.
type IPassedParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassedParamContext differentiates from other interfaces.
	IsPassedParamContext()
}

type PassedParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassedParamContext() *PassedParamContext {
	var p = new(PassedParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_passedParam
	return p
}

func (*PassedParamContext) IsPassedParamContext() {}

func NewPassedParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassedParamContext {
	var p = new(PassedParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_passedParam

	return p
}

func (s *PassedParamContext) GetParser() antlr.Parser { return s.parser }

func (s *PassedParamContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *PassedParamContext) ListSpaceSeparated() IListSpaceSeparatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListSpaceSeparatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListSpaceSeparatedContext)
}

func (s *PassedParamContext) ListBracketed() IListBracketedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListBracketedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListBracketedContext)
}

func (s *PassedParamContext) Map_() IMap_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMap_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMap_Context)
}

func (s *PassedParamContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *PassedParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *PassedParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassedParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassedParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterPassedParam(s)
	}
}

func (s *PassedParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitPassedParam(s)
	}
}

func (p *ScssParser) PassedParam() (localctx IPassedParamContext) {
	this := p
	_ = this

	localctx = NewPassedParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ScssParserRULE_passedParam)

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
	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(200)
			p.VariableName()
		}
		{
			p.SetState(201)
			p.Match(ScssParserCOLON)
		}

	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(205)
			p.CommandStatement()
		}

	case 2:
		{
			p.SetState(206)
			p.ListSpaceSeparated()
		}

	case 3:
		{
			p.SetState(207)
			p.ListBracketed()
		}

	case 4:
		{
			p.SetState(208)
			p.Map_()
		}

	}

	return localctx
}

// IMixinDeclarationContext is an interface to support dynamic dispatch.
type IMixinDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixinDeclarationContext differentiates from other interfaces.
	IsMixinDeclarationContext()
}

type MixinDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixinDeclarationContext() *MixinDeclarationContext {
	var p = new(MixinDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mixinDeclaration
	return p
}

func (*MixinDeclarationContext) IsMixinDeclarationContext() {}

func NewMixinDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixinDeclarationContext {
	var p = new(MixinDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mixinDeclaration

	return p
}

func (s *MixinDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MixinDeclarationContext) MIXIN() antlr.TerminalNode {
	return s.GetToken(ScssParserMIXIN, 0)
}

func (s *MixinDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *MixinDeclarationContext) FunctionIdentifier() antlr.TerminalNode {
	return s.GetToken(ScssParserFunctionIdentifier, 0)
}

func (s *MixinDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *MixinDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *MixinDeclarationContext) DeclaredParams() IDeclaredParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaredParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaredParamsContext)
}

func (s *MixinDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *MixinDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixinDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixinDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMixinDeclaration(s)
	}
}

func (s *MixinDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMixinDeclaration(s)
	}
}

func (p *ScssParser) MixinDeclaration() (localctx IMixinDeclarationContext) {
	this := p
	_ = this

	localctx = NewMixinDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ScssParserRULE_mixinDeclaration)
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
		p.SetState(211)
		p.Match(ScssParserMIXIN)
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserFunctionIdentifier:
		{
			p.SetState(212)
			p.Match(ScssParserFunctionIdentifier)
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserDOLLAR)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR))) != 0) || _la == ScssParserIdentifier {
			{
				p.SetState(213)
				p.DeclaredParams()
			}

		}
		{
			p.SetState(216)
			p.Match(ScssParserRPAREN)
		}

	case ScssParserIdentifier:
		{
			p.SetState(217)
			p.Match(ScssParserIdentifier)
		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScssParserLPAREN {
			{
				p.SetState(218)
				p.Match(ScssParserLPAREN)
			}
			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserDOLLAR)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR))) != 0) || _la == ScssParserIdentifier {
				{
					p.SetState(219)
					p.DeclaredParams()
				}

			}
			{
				p.SetState(222)
				p.Match(ScssParserRPAREN)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(227)
		p.Block()
	}

	return localctx
}

// IContentDeclarationContext is an interface to support dynamic dispatch.
type IContentDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentDeclarationContext differentiates from other interfaces.
	IsContentDeclarationContext()
}

type ContentDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentDeclarationContext() *ContentDeclarationContext {
	var p = new(ContentDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_contentDeclaration
	return p
}

func (*ContentDeclarationContext) IsContentDeclarationContext() {}

func NewContentDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentDeclarationContext {
	var p = new(ContentDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_contentDeclaration

	return p
}

func (s *ContentDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentDeclarationContext) CONTENT() antlr.TerminalNode {
	return s.GetToken(ScssParserCONTENT, 0)
}

func (s *ContentDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ScssParserSEMI, 0)
}

func (s *ContentDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *ContentDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *ContentDeclarationContext) PassedParams() IPassedParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassedParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassedParamsContext)
}

func (s *ContentDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterContentDeclaration(s)
	}
}

func (s *ContentDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitContentDeclaration(s)
	}
}

func (p *ScssParser) ContentDeclaration() (localctx IContentDeclarationContext) {
	this := p
	_ = this

	localctx = NewContentDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ScssParserRULE_contentDeclaration)
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
		p.SetState(229)
		p.Match(ScssParserCONTENT)
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserLPAREN {
		{
			p.SetState(230)
			p.Match(ScssParserLPAREN)
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL_)|(1<<ScssParserInterpolationStart)|(1<<ScssParserLPAREN)|(1<<ScssParserLBRACK)|(1<<ScssParserDOLLAR)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR)|(1<<ScssParserMINUS_LPAREN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ScssParserPLUS_LPAREN-32))|(1<<(ScssParserUrlStart-32))|(1<<(ScssParserFROM-32))|(1<<(ScssParserTO-32))|(1<<(ScssParserTHROUGH-32))|(1<<(ScssParserONLY-32))|(1<<(ScssParserNOT-32))|(1<<(ScssParserAND_WORD-32))|(1<<(ScssParserUSING-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(ScssParserIdentifier-64))|(1<<(ScssParserFunctionIdentifier-64))|(1<<(ScssParserStringLiteral-64))|(1<<(ScssParserNumber-64))|(1<<(ScssParserColor-64)))) != 0) {
			{
				p.SetState(231)
				p.PassedParams()
			}

		}
		{
			p.SetState(234)
			p.Match(ScssParserRPAREN)
		}

	}
	{
		p.SetState(237)
		p.Match(ScssParserSEMI)
	}

	return localctx
}

// IIncludeDeclarationContext is an interface to support dynamic dispatch.
type IIncludeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncludeDeclarationContext differentiates from other interfaces.
	IsIncludeDeclarationContext()
}

type IncludeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeDeclarationContext() *IncludeDeclarationContext {
	var p = new(IncludeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_includeDeclaration
	return p
}

func (*IncludeDeclarationContext) IsIncludeDeclarationContext() {}

func NewIncludeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeDeclarationContext {
	var p = new(IncludeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_includeDeclaration

	return p
}

func (s *IncludeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeDeclarationContext) INCLUDE() antlr.TerminalNode {
	return s.GetToken(ScssParserINCLUDE, 0)
}

func (s *IncludeDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *IncludeDeclarationContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *IncludeDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ScssParserSEMI, 0)
}

func (s *IncludeDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IncludeDeclarationContext) USING() antlr.TerminalNode {
	return s.GetToken(ScssParserUSING, 0)
}

func (s *IncludeDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *IncludeDeclarationContext) DeclaredParams() IDeclaredParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaredParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaredParamsContext)
}

func (s *IncludeDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *IncludeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterIncludeDeclaration(s)
	}
}

func (s *IncludeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitIncludeDeclaration(s)
	}
}

func (p *ScssParser) IncludeDeclaration() (localctx IIncludeDeclarationContext) {
	this := p
	_ = this

	localctx = NewIncludeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ScssParserRULE_includeDeclaration)
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
		p.SetState(239)
		p.Match(ScssParserINCLUDE)
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(240)
			p.Match(ScssParserIdentifier)
		}

	case 2:
		{
			p.SetState(241)
			p.FunctionCall()
		}

	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(244)
			p.Match(ScssParserSEMI)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 2 {
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScssParserUSING {
			{
				p.SetState(245)
				p.Match(ScssParserUSING)
			}
			{
				p.SetState(246)
				p.Match(ScssParserLPAREN)
			}
			{
				p.SetState(247)
				p.DeclaredParams()
			}
			{
				p.SetState(248)
				p.Match(ScssParserRPAREN)
			}

		}
		{
			p.SetState(252)
			p.Block()
		}

	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(ScssParserFUNCTION, 0)
}

func (s *FunctionDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *FunctionDeclarationContext) BlockStart() antlr.TerminalNode {
	return s.GetToken(ScssParserBlockStart, 0)
}

func (s *FunctionDeclarationContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(ScssParserBlockEnd, 0)
}

func (s *FunctionDeclarationContext) FunctionIdentifier() antlr.TerminalNode {
	return s.GetToken(ScssParserFunctionIdentifier, 0)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *FunctionDeclarationContext) DeclaredParams() IDeclaredParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaredParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaredParamsContext)
}

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *ScssParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	this := p
	_ = this

	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ScssParserRULE_functionDeclaration)
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
		p.SetState(255)
		p.Match(ScssParserFUNCTION)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserFunctionIdentifier:
		{
			p.SetState(256)
			p.Match(ScssParserFunctionIdentifier)
		}

	case ScssParserIdentifier:
		{
			p.SetState(257)
			p.Match(ScssParserIdentifier)
		}
		{
			p.SetState(258)
			p.Match(ScssParserLPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserDOLLAR)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR))) != 0) || _la == ScssParserIdentifier {
		{
			p.SetState(261)
			p.DeclaredParams()
		}

	}
	{
		p.SetState(264)
		p.Match(ScssParserRPAREN)
	}
	{
		p.SetState(265)
		p.Match(ScssParserBlockStart)
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL_)|(1<<ScssParserInterpolationStart)|(1<<ScssParserLPAREN)|(1<<ScssParserLBRACK)|(1<<ScssParserGT)|(1<<ScssParserTIL)|(1<<ScssParserDOT)|(1<<ScssParserDOLLAR)|(1<<ScssParserAND)|(1<<ScssParserHASH)|(1<<ScssParserPLUS)|(1<<ScssParserTIMES)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR)|(1<<ScssParserMINUS_LPAREN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ScssParserPLUS_LPAREN-32))|(1<<(ScssParserUrlStart-32))|(1<<(ScssParserMIXIN-32))|(1<<(ScssParserFUNCTION-32))|(1<<(ScssParserAT_IF-32))|(1<<(ScssParserAT_FOR-32))|(1<<(ScssParserAT_WHILE-32))|(1<<(ScssParserAT_EACH-32))|(1<<(ScssParserINCLUDE-32))|(1<<(ScssParserIMPORT-32))|(1<<(ScssParserUSE-32))|(1<<(ScssParserRETURN-32))|(1<<(ScssParserMEDIA-32))|(1<<(ScssParserCONTENT-32))|(1<<(ScssParserFROM-32))|(1<<(ScssParserTO-32))|(1<<(ScssParserTHROUGH-32))|(1<<(ScssParserONLY-32))|(1<<(ScssParserNOT-32))|(1<<(ScssParserAND_WORD-32))|(1<<(ScssParserUSING-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(ScssParserIdentifier-64))|(1<<(ScssParserPseudoIdentifier-64))|(1<<(ScssParserFunctionIdentifier-64))|(1<<(ScssParserStringLiteral-64))|(1<<(ScssParserNumber-64))|(1<<(ScssParserColor-64)))) != 0) {
		{
			p.SetState(266)
			p.FunctionBody()
		}

	}
	{
		p.SetState(269)
		p.Match(ScssParserBlockEnd)
	}

	return localctx
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_functionBody
	return p
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) FunctionReturn() IFunctionReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionReturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionReturnContext)
}

func (s *FunctionBodyContext) AllFunctionStatement() []IFunctionStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionStatementContext)(nil)).Elem())
	var tst = make([]IFunctionStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionStatementContext)
		}
	}

	return tst
}

func (s *FunctionBodyContext) FunctionStatement(i int) IFunctionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionStatementContext)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (p *ScssParser) FunctionBody() (localctx IFunctionBodyContext) {
	this := p
	_ = this

	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ScssParserRULE_functionBody)
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
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL_)|(1<<ScssParserInterpolationStart)|(1<<ScssParserLPAREN)|(1<<ScssParserLBRACK)|(1<<ScssParserGT)|(1<<ScssParserTIL)|(1<<ScssParserDOT)|(1<<ScssParserDOLLAR)|(1<<ScssParserAND)|(1<<ScssParserHASH)|(1<<ScssParserPLUS)|(1<<ScssParserTIMES)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR)|(1<<ScssParserMINUS_LPAREN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ScssParserPLUS_LPAREN-32))|(1<<(ScssParserUrlStart-32))|(1<<(ScssParserMIXIN-32))|(1<<(ScssParserFUNCTION-32))|(1<<(ScssParserAT_IF-32))|(1<<(ScssParserAT_FOR-32))|(1<<(ScssParserAT_WHILE-32))|(1<<(ScssParserAT_EACH-32))|(1<<(ScssParserINCLUDE-32))|(1<<(ScssParserIMPORT-32))|(1<<(ScssParserUSE-32))|(1<<(ScssParserMEDIA-32))|(1<<(ScssParserCONTENT-32))|(1<<(ScssParserFROM-32))|(1<<(ScssParserTO-32))|(1<<(ScssParserTHROUGH-32))|(1<<(ScssParserONLY-32))|(1<<(ScssParserNOT-32))|(1<<(ScssParserAND_WORD-32))|(1<<(ScssParserUSING-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(ScssParserIdentifier-64))|(1<<(ScssParserPseudoIdentifier-64))|(1<<(ScssParserFunctionIdentifier-64))|(1<<(ScssParserStringLiteral-64))|(1<<(ScssParserNumber-64))|(1<<(ScssParserColor-64)))) != 0) {
		{
			p.SetState(271)
			p.FunctionStatement()
		}

		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(277)
		p.FunctionReturn()
	}

	return localctx
}

// IFunctionReturnContext is an interface to support dynamic dispatch.
type IFunctionReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionReturnContext differentiates from other interfaces.
	IsFunctionReturnContext()
}

type FunctionReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionReturnContext() *FunctionReturnContext {
	var p = new(FunctionReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_functionReturn
	return p
}

func (*FunctionReturnContext) IsFunctionReturnContext() {}

func NewFunctionReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionReturnContext {
	var p = new(FunctionReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_functionReturn

	return p
}

func (s *FunctionReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(ScssParserRETURN, 0)
}

func (s *FunctionReturnContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *FunctionReturnContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ScssParserSEMI, 0)
}

func (s *FunctionReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterFunctionReturn(s)
	}
}

func (s *FunctionReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitFunctionReturn(s)
	}
}

func (p *ScssParser) FunctionReturn() (localctx IFunctionReturnContext) {
	this := p
	_ = this

	localctx = NewFunctionReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ScssParserRULE_functionReturn)

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
		p.SetState(279)
		p.Match(ScssParserRETURN)
	}
	{
		p.SetState(280)
		p.CommandStatement()
	}
	{
		p.SetState(281)
		p.Match(ScssParserSEMI)
	}

	return localctx
}

// IFunctionStatementContext is an interface to support dynamic dispatch.
type IFunctionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionStatementContext differentiates from other interfaces.
	IsFunctionStatementContext()
}

type FunctionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionStatementContext() *FunctionStatementContext {
	var p = new(FunctionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_functionStatement
	return p
}

func (*FunctionStatementContext) IsFunctionStatementContext() {}

func NewFunctionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionStatementContext {
	var p = new(FunctionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_functionStatement

	return p
}

func (s *FunctionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionStatementContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *FunctionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ScssParserSEMI, 0)
}

func (s *FunctionStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FunctionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterFunctionStatement(s)
	}
}

func (s *FunctionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitFunctionStatement(s)
	}
}

func (p *ScssParser) FunctionStatement() (localctx IFunctionStatementContext) {
	this := p
	_ = this

	localctx = NewFunctionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ScssParserRULE_functionStatement)

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

	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(283)
			p.CommandStatement()
		}
		{
			p.SetState(284)
			p.Match(ScssParserSEMI)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Statement()
		}

	}

	return localctx
}

// ICommandStatementContext is an interface to support dynamic dispatch.
type ICommandStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandStatementContext differentiates from other interfaces.
	IsCommandStatementContext()
}

type CommandStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandStatementContext() *CommandStatementContext {
	var p = new(CommandStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_commandStatement
	return p
}

func (*CommandStatementContext) IsCommandStatementContext() {}

func NewCommandStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandStatementContext {
	var p = new(CommandStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_commandStatement

	return p
}

func (s *CommandStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CommandStatementContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *CommandStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *CommandStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *CommandStatementContext) MINUS_LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserMINUS_LPAREN, 0)
}

func (s *CommandStatementContext) PLUS_LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserPLUS_LPAREN, 0)
}

func (s *CommandStatementContext) MathStatement() IMathStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathStatementContext)
}

func (s *CommandStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterCommandStatement(s)
	}
}

func (s *CommandStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitCommandStatement(s)
	}
}

func (p *ScssParser) CommandStatement() (localctx ICommandStatementContext) {
	this := p
	_ = this

	localctx = NewCommandStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ScssParserRULE_commandStatement)
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
	p.SetState(294)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserNULL_, ScssParserInterpolationStart, ScssParserDOLLAR, ScssParserMINUS_DOLLAR, ScssParserPLUS_DOLLAR, ScssParserUrlStart, ScssParserFROM, ScssParserTO, ScssParserTHROUGH, ScssParserONLY, ScssParserNOT, ScssParserAND_WORD, ScssParserUSING, ScssParserIdentifier, ScssParserFunctionIdentifier, ScssParserStringLiteral, ScssParserNumber, ScssParserColor:
		{
			p.SetState(289)
			p.Expression()
		}

	case ScssParserLPAREN, ScssParserMINUS_LPAREN, ScssParserPLUS_LPAREN:
		{
			p.SetState(290)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-7)&-(0x1f+1)) == 0 && ((1<<uint((_la-7)))&((1<<(ScssParserLPAREN-7))|(1<<(ScssParserMINUS_LPAREN-7))|(1<<(ScssParserPLUS_LPAREN-7)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(291)
			p.CommandStatement()
		}
		{
			p.SetState(292)
			p.Match(ScssParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserPLUS)|(1<<ScssParserTIMES)|(1<<ScssParserDIV)|(1<<ScssParserMINUS)|(1<<ScssParserPERC))) != 0 {
		{
			p.SetState(296)
			p.MathStatement()
		}

	}

	return localctx
}

// IMathCharacterContext is an interface to support dynamic dispatch.
type IMathCharacterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathCharacterContext differentiates from other interfaces.
	IsMathCharacterContext()
}

type MathCharacterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathCharacterContext() *MathCharacterContext {
	var p = new(MathCharacterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mathCharacter
	return p
}

func (*MathCharacterContext) IsMathCharacterContext() {}

func NewMathCharacterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathCharacterContext {
	var p = new(MathCharacterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mathCharacter

	return p
}

func (s *MathCharacterContext) GetParser() antlr.Parser { return s.parser }

func (s *MathCharacterContext) TIMES() antlr.TerminalNode {
	return s.GetToken(ScssParserTIMES, 0)
}

func (s *MathCharacterContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ScssParserPLUS, 0)
}

func (s *MathCharacterContext) DIV() antlr.TerminalNode {
	return s.GetToken(ScssParserDIV, 0)
}

func (s *MathCharacterContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ScssParserMINUS, 0)
}

func (s *MathCharacterContext) PERC() antlr.TerminalNode {
	return s.GetToken(ScssParserPERC, 0)
}

func (s *MathCharacterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathCharacterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathCharacterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMathCharacter(s)
	}
}

func (s *MathCharacterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMathCharacter(s)
	}
}

func (p *ScssParser) MathCharacter() (localctx IMathCharacterContext) {
	this := p
	_ = this

	localctx = NewMathCharacterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ScssParserRULE_mathCharacter)
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
		p.SetState(299)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserPLUS)|(1<<ScssParserTIMES)|(1<<ScssParserDIV)|(1<<ScssParserMINUS)|(1<<ScssParserPERC))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMathStatementContext is an interface to support dynamic dispatch.
type IMathStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathStatementContext differentiates from other interfaces.
	IsMathStatementContext()
}

type MathStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathStatementContext() *MathStatementContext {
	var p = new(MathStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mathStatement
	return p
}

func (*MathStatementContext) IsMathStatementContext() {}

func NewMathStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathStatementContext {
	var p = new(MathStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mathStatement

	return p
}

func (s *MathStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MathStatementContext) MathCharacter() IMathCharacterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathCharacterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathCharacterContext)
}

func (s *MathStatementContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *MathStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMathStatement(s)
	}
}

func (s *MathStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMathStatement(s)
	}
}

func (p *ScssParser) MathStatement() (localctx IMathStatementContext) {
	this := p
	_ = this

	localctx = NewMathStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ScssParserRULE_mathStatement)

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
		p.MathCharacter()
	}
	{
		p.SetState(302)
		p.CommandStatement()
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
	p.RuleIndex = ScssParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Measurement() IMeasurementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeasurementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeasurementContext)
}

func (s *ExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExpressionContext) Color() antlr.TerminalNode {
	return s.GetToken(ScssParserColor, 0)
}

func (s *ExpressionContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ScssParserStringLiteral, 0)
}

func (s *ExpressionContext) NULL_() antlr.TerminalNode {
	return s.GetToken(ScssParserNULL_, 0)
}

func (s *ExpressionContext) Url() IUrlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUrlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUrlContext)
}

func (s *ExpressionContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ScssParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ScssParserRULE_expression)

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

	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.Measurement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Identifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(306)
			p.Match(ScssParserColor)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(307)
			p.Match(ScssParserStringLiteral)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(308)
			p.Match(ScssParserNULL_)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(309)
			p.Url()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(310)
			p.VariableName()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(311)
			p.FunctionCall()
		}

	}

	return localctx
}

// IIfDeclarationContext is an interface to support dynamic dispatch.
type IIfDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfDeclarationContext differentiates from other interfaces.
	IsIfDeclarationContext()
}

type IfDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfDeclarationContext() *IfDeclarationContext {
	var p = new(IfDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_ifDeclaration
	return p
}

func (*IfDeclarationContext) IsIfDeclarationContext() {}

func NewIfDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfDeclarationContext {
	var p = new(IfDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_ifDeclaration

	return p
}

func (s *IfDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *IfDeclarationContext) AT_IF() antlr.TerminalNode {
	return s.GetToken(ScssParserAT_IF, 0)
}

func (s *IfDeclarationContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *IfDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfDeclarationContext) AllElseIfStatement() []IElseIfStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfStatementContext)(nil)).Elem())
	var tst = make([]IElseIfStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfStatementContext)
		}
	}

	return tst
}

func (s *IfDeclarationContext) ElseIfStatement(i int) IElseIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfStatementContext)
}

func (s *IfDeclarationContext) ElseStatement() IElseStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStatementContext)
}

func (s *IfDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterIfDeclaration(s)
	}
}

func (s *IfDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitIfDeclaration(s)
	}
}

func (p *ScssParser) IfDeclaration() (localctx IIfDeclarationContext) {
	this := p
	_ = this

	localctx = NewIfDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ScssParserRULE_ifDeclaration)
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
		p.SetState(314)
		p.Match(ScssParserAT_IF)
	}
	{
		p.SetState(315)
		p.Conditions()
	}
	{
		p.SetState(316)
		p.Block()
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(317)
				p.ElseIfStatement()
			}

		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserAT_ELSE {
		{
			p.SetState(323)
			p.ElseStatement()
		}

	}

	return localctx
}

// IElseIfStatementContext is an interface to support dynamic dispatch.
type IElseIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfStatementContext differentiates from other interfaces.
	IsElseIfStatementContext()
}

type ElseIfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStatementContext() *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_elseIfStatement
	return p
}

func (*ElseIfStatementContext) IsElseIfStatementContext() {}

func NewElseIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_elseIfStatement

	return p
}

func (s *ElseIfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStatementContext) AT_ELSE() antlr.TerminalNode {
	return s.GetToken(ScssParserAT_ELSE, 0)
}

func (s *ElseIfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(ScssParserIF, 0)
}

func (s *ElseIfStatementContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *ElseIfStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseIfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterElseIfStatement(s)
	}
}

func (s *ElseIfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitElseIfStatement(s)
	}
}

func (p *ScssParser) ElseIfStatement() (localctx IElseIfStatementContext) {
	this := p
	_ = this

	localctx = NewElseIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ScssParserRULE_elseIfStatement)

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
		p.SetState(326)
		p.Match(ScssParserAT_ELSE)
	}
	{
		p.SetState(327)
		p.Match(ScssParserIF)
	}
	{
		p.SetState(328)
		p.Conditions()
	}
	{
		p.SetState(329)
		p.Block()
	}

	return localctx
}

// IElseStatementContext is an interface to support dynamic dispatch.
type IElseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatementContext differentiates from other interfaces.
	IsElseStatementContext()
}

type ElseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatementContext() *ElseStatementContext {
	var p = new(ElseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_elseStatement
	return p
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) AT_ELSE() antlr.TerminalNode {
	return s.GetToken(ScssParserAT_ELSE, 0)
}

func (s *ElseStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterElseStatement(s)
	}
}

func (s *ElseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitElseStatement(s)
	}
}

func (p *ScssParser) ElseStatement() (localctx IElseStatementContext) {
	this := p
	_ = this

	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ScssParserRULE_elseStatement)

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
		p.Match(ScssParserAT_ELSE)
	}
	{
		p.SetState(332)
		p.Block()
	}

	return localctx
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_conditions
	return p
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionsContext) COMBINE_COMPARE() antlr.TerminalNode {
	return s.GetToken(ScssParserCOMBINE_COMPARE, 0)
}

func (s *ConditionsContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *ConditionsContext) NULL_() antlr.TerminalNode {
	return s.GetToken(ScssParserNULL_, 0)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (p *ScssParser) Conditions() (localctx IConditionsContext) {
	this := p
	_ = this

	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ScssParserRULE_conditions)

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

	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(334)
			p.Condition()
		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(335)
				p.Match(ScssParserCOMBINE_COMPARE)
			}
			{
				p.SetState(336)
				p.Conditions()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(339)
			p.Match(ScssParserNULL_)
		}

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
	p.RuleIndex = ScssParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *ConditionContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *ConditionContext) EQEQ() antlr.TerminalNode {
	return s.GetToken(ScssParserEQEQ, 0)
}

func (s *ConditionContext) LT() antlr.TerminalNode {
	return s.GetToken(ScssParserLT, 0)
}

func (s *ConditionContext) GT() antlr.TerminalNode {
	return s.GetToken(ScssParserGT, 0)
}

func (s *ConditionContext) NOTEQ() antlr.TerminalNode {
	return s.GetToken(ScssParserNOTEQ, 0)
}

func (s *ConditionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *ConditionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *ScssParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ScssParserRULE_condition)
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

	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(342)
			p.CommandStatement()
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(ScssParserGT-13))|(1<<(ScssParserLT-13))|(1<<(ScssParserEQEQ-13))|(1<<(ScssParserNOTEQ-13)))) != 0 {
			{
				p.SetState(343)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-13)&-(0x1f+1)) == 0 && ((1<<uint((_la-13)))&((1<<(ScssParserGT-13))|(1<<(ScssParserLT-13))|(1<<(ScssParserEQEQ-13))|(1<<(ScssParserNOTEQ-13)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(344)
				p.Conditions()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(347)
			p.Match(ScssParserLPAREN)
		}
		{
			p.SetState(348)
			p.Conditions()
		}
		{
			p.SetState(349)
			p.Match(ScssParserRPAREN)
		}

	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *VariableDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *VariableDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ScssParserSEMI, 0)
}

func (s *VariableDeclarationContext) PropertyValue() IPropertyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyValueContext)
}

func (s *VariableDeclarationContext) ListBracketed() IListBracketedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListBracketedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListBracketedContext)
}

func (s *VariableDeclarationContext) Map_() IMap_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMap_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMap_Context)
}

func (s *VariableDeclarationContext) POUND_DEFAULT() antlr.TerminalNode {
	return s.GetToken(ScssParserPOUND_DEFAULT, 0)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *ScssParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	this := p
	_ = this

	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ScssParserRULE_variableDeclaration)
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
		p.SetState(353)
		p.VariableName()
	}
	{
		p.SetState(354)
		p.Match(ScssParserCOLON)
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(355)
			p.PropertyValue()
		}

	case 2:
		{
			p.SetState(356)
			p.ListBracketed()
		}

	case 3:
		{
			p.SetState(357)
			p.Map_()
		}

	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserPOUND_DEFAULT {
		{
			p.SetState(360)
			p.Match(ScssParserPOUND_DEFAULT)
		}

	}
	{
		p.SetState(363)
		p.Match(ScssParserSEMI)
	}

	return localctx
}

// IForDeclarationContext is an interface to support dynamic dispatch.
type IForDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForDeclarationContext differentiates from other interfaces.
	IsForDeclarationContext()
}

type ForDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForDeclarationContext() *ForDeclarationContext {
	var p = new(ForDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_forDeclaration
	return p
}

func (*ForDeclarationContext) IsForDeclarationContext() {}

func NewForDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForDeclarationContext {
	var p = new(ForDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_forDeclaration

	return p
}

func (s *ForDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ForDeclarationContext) AT_FOR() antlr.TerminalNode {
	return s.GetToken(ScssParserAT_FOR, 0)
}

func (s *ForDeclarationContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *ForDeclarationContext) FROM() antlr.TerminalNode {
	return s.GetToken(ScssParserFROM, 0)
}

func (s *ForDeclarationContext) FromNumber() IFromNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFromNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFromNumberContext)
}

func (s *ForDeclarationContext) ThroughNumber() IThroughNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThroughNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThroughNumberContext)
}

func (s *ForDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForDeclarationContext) TO() antlr.TerminalNode {
	return s.GetToken(ScssParserTO, 0)
}

func (s *ForDeclarationContext) THROUGH() antlr.TerminalNode {
	return s.GetToken(ScssParserTHROUGH, 0)
}

func (s *ForDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterForDeclaration(s)
	}
}

func (s *ForDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitForDeclaration(s)
	}
}

func (p *ScssParser) ForDeclaration() (localctx IForDeclarationContext) {
	this := p
	_ = this

	localctx = NewForDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ScssParserRULE_forDeclaration)
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
		p.SetState(365)
		p.Match(ScssParserAT_FOR)
	}
	{
		p.SetState(366)
		p.VariableName()
	}
	{
		p.SetState(367)
		p.Match(ScssParserFROM)
	}
	{
		p.SetState(368)
		p.FromNumber()
	}
	{
		p.SetState(369)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ScssParserTO || _la == ScssParserTHROUGH) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(370)
		p.ThroughNumber()
	}
	{
		p.SetState(371)
		p.Block()
	}

	return localctx
}

// IFromNumberContext is an interface to support dynamic dispatch.
type IFromNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFromNumberContext differentiates from other interfaces.
	IsFromNumberContext()
}

type FromNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromNumberContext() *FromNumberContext {
	var p = new(FromNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_fromNumber
	return p
}

func (*FromNumberContext) IsFromNumberContext() {}

func NewFromNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromNumberContext {
	var p = new(FromNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_fromNumber

	return p
}

func (s *FromNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *FromNumberContext) Number() antlr.TerminalNode {
	return s.GetToken(ScssParserNumber, 0)
}

func (s *FromNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterFromNumber(s)
	}
}

func (s *FromNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitFromNumber(s)
	}
}

func (p *ScssParser) FromNumber() (localctx IFromNumberContext) {
	this := p
	_ = this

	localctx = NewFromNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ScssParserRULE_fromNumber)

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
		p.SetState(373)
		p.Match(ScssParserNumber)
	}

	return localctx
}

// IThroughNumberContext is an interface to support dynamic dispatch.
type IThroughNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThroughNumberContext differentiates from other interfaces.
	IsThroughNumberContext()
}

type ThroughNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThroughNumberContext() *ThroughNumberContext {
	var p = new(ThroughNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_throughNumber
	return p
}

func (*ThroughNumberContext) IsThroughNumberContext() {}

func NewThroughNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThroughNumberContext {
	var p = new(ThroughNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_throughNumber

	return p
}

func (s *ThroughNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *ThroughNumberContext) Number() antlr.TerminalNode {
	return s.GetToken(ScssParserNumber, 0)
}

func (s *ThroughNumberContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ThroughNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThroughNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThroughNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterThroughNumber(s)
	}
}

func (s *ThroughNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitThroughNumber(s)
	}
}

func (p *ScssParser) ThroughNumber() (localctx IThroughNumberContext) {
	this := p
	_ = this

	localctx = NewThroughNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ScssParserRULE_throughNumber)

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

	p.SetState(377)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserNumber:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(375)
			p.Match(ScssParserNumber)
		}

	case ScssParserIdentifier, ScssParserFunctionIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(376)
			p.FunctionCall()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWhileDeclarationContext is an interface to support dynamic dispatch.
type IWhileDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileDeclarationContext differentiates from other interfaces.
	IsWhileDeclarationContext()
}

type WhileDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileDeclarationContext() *WhileDeclarationContext {
	var p = new(WhileDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_whileDeclaration
	return p
}

func (*WhileDeclarationContext) IsWhileDeclarationContext() {}

func NewWhileDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileDeclarationContext {
	var p = new(WhileDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_whileDeclaration

	return p
}

func (s *WhileDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileDeclarationContext) AT_WHILE() antlr.TerminalNode {
	return s.GetToken(ScssParserAT_WHILE, 0)
}

func (s *WhileDeclarationContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *WhileDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterWhileDeclaration(s)
	}
}

func (s *WhileDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitWhileDeclaration(s)
	}
}

func (p *ScssParser) WhileDeclaration() (localctx IWhileDeclarationContext) {
	this := p
	_ = this

	localctx = NewWhileDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ScssParserRULE_whileDeclaration)

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
		p.SetState(379)
		p.Match(ScssParserAT_WHILE)
	}
	{
		p.SetState(380)
		p.Conditions()
	}
	{
		p.SetState(381)
		p.Block()
	}

	return localctx
}

// IEachDeclarationContext is an interface to support dynamic dispatch.
type IEachDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEachDeclarationContext differentiates from other interfaces.
	IsEachDeclarationContext()
}

type EachDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEachDeclarationContext() *EachDeclarationContext {
	var p = new(EachDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_eachDeclaration
	return p
}

func (*EachDeclarationContext) IsEachDeclarationContext() {}

func NewEachDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EachDeclarationContext {
	var p = new(EachDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_eachDeclaration

	return p
}

func (s *EachDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *EachDeclarationContext) AT_EACH() antlr.TerminalNode {
	return s.GetToken(ScssParserAT_EACH, 0)
}

func (s *EachDeclarationContext) AllVariableName() []IVariableNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableNameContext)(nil)).Elem())
	var tst = make([]IVariableNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableNameContext)
		}
	}

	return tst
}

func (s *EachDeclarationContext) VariableName(i int) IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *EachDeclarationContext) IN() antlr.TerminalNode {
	return s.GetToken(ScssParserIN, 0)
}

func (s *EachDeclarationContext) EachValueList() IEachValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEachValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEachValueListContext)
}

func (s *EachDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *EachDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *EachDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *EachDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EachDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EachDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterEachDeclaration(s)
	}
}

func (s *EachDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitEachDeclaration(s)
	}
}

func (p *ScssParser) EachDeclaration() (localctx IEachDeclarationContext) {
	this := p
	_ = this

	localctx = NewEachDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ScssParserRULE_eachDeclaration)
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
		p.SetState(383)
		p.Match(ScssParserAT_EACH)
	}
	{
		p.SetState(384)
		p.VariableName()
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserCOMMA {
		{
			p.SetState(385)
			p.Match(ScssParserCOMMA)
		}
		{
			p.SetState(386)
			p.VariableName()
		}

		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(392)
		p.Match(ScssParserIN)
	}
	{
		p.SetState(393)
		p.EachValueList()
	}
	{
		p.SetState(394)
		p.Block()
	}

	return localctx
}

// IEachValueListContext is an interface to support dynamic dispatch.
type IEachValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEachValueListContext differentiates from other interfaces.
	IsEachValueListContext()
}

type EachValueListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEachValueListContext() *EachValueListContext {
	var p = new(EachValueListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_eachValueList
	return p
}

func (*EachValueListContext) IsEachValueListContext() {}

func NewEachValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EachValueListContext {
	var p = new(EachValueListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_eachValueList

	return p
}

func (s *EachValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *EachValueListContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *EachValueListContext) List_() IList_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_Context)
}

func (s *EachValueListContext) Map_() IMap_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMap_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMap_Context)
}

func (s *EachValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EachValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EachValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterEachValueList(s)
	}
}

func (s *EachValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitEachValueList(s)
	}
}

func (p *ScssParser) EachValueList() (localctx IEachValueListContext) {
	this := p
	_ = this

	localctx = NewEachValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ScssParserRULE_eachValueList)

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

	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(396)
			p.CommandStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(397)
			p.List_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(398)
			p.Map_()
		}

	}

	return localctx
}

// IImportDeclarationContext is an interface to support dynamic dispatch.
type IImportDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclarationContext differentiates from other interfaces.
	IsImportDeclarationContext()
}

type ImportDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclarationContext() *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_importDeclaration
	return p
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ScssParserIMPORT, 0)
}

func (s *ImportDeclarationContext) ReferenceUrl() IReferenceUrlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceUrlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferenceUrlContext)
}

func (s *ImportDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ScssParserSEMI, 0)
}

func (s *ImportDeclarationContext) USE() antlr.TerminalNode {
	return s.GetToken(ScssParserUSE, 0)
}

func (s *ImportDeclarationContext) AsClause() IAsClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsClauseContext)
}

func (s *ImportDeclarationContext) WithClause() IWithClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWithClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWithClauseContext)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (p *ScssParser) ImportDeclaration() (localctx IImportDeclarationContext) {
	this := p
	_ = this

	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ScssParserRULE_importDeclaration)
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

	p.SetState(415)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserIMPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(401)
			p.Match(ScssParserIMPORT)
		}
		{
			p.SetState(402)
			p.ReferenceUrl()
		}
		{
			p.SetState(403)
			p.Match(ScssParserSEMI)
		}

	case ScssParserUSE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(405)
			p.Match(ScssParserUSE)
		}
		{
			p.SetState(406)
			p.ReferenceUrl()
		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScssParserAS {
			{
				p.SetState(407)
				p.AsClause()
			}

		}
		p.SetState(411)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScssParserWITH {
			{
				p.SetState(410)
				p.WithClause()
			}

		}
		{
			p.SetState(413)
			p.Match(ScssParserSEMI)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReferenceUrlContext is an interface to support dynamic dispatch.
type IReferenceUrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceUrlContext differentiates from other interfaces.
	IsReferenceUrlContext()
}

type ReferenceUrlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceUrlContext() *ReferenceUrlContext {
	var p = new(ReferenceUrlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_referenceUrl
	return p
}

func (*ReferenceUrlContext) IsReferenceUrlContext() {}

func NewReferenceUrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceUrlContext {
	var p = new(ReferenceUrlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_referenceUrl

	return p
}

func (s *ReferenceUrlContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceUrlContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ScssParserStringLiteral, 0)
}

func (s *ReferenceUrlContext) UrlStart() antlr.TerminalNode {
	return s.GetToken(ScssParserUrlStart, 0)
}

func (s *ReferenceUrlContext) Url() antlr.TerminalNode {
	return s.GetToken(ScssParserUrl, 0)
}

func (s *ReferenceUrlContext) UrlEnd() antlr.TerminalNode {
	return s.GetToken(ScssParserUrlEnd, 0)
}

func (s *ReferenceUrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceUrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceUrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterReferenceUrl(s)
	}
}

func (s *ReferenceUrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitReferenceUrl(s)
	}
}

func (p *ScssParser) ReferenceUrl() (localctx IReferenceUrlContext) {
	this := p
	_ = this

	localctx = NewReferenceUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ScssParserRULE_referenceUrl)

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

	p.SetState(421)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(417)
			p.Match(ScssParserStringLiteral)
		}

	case ScssParserUrlStart:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(418)
			p.Match(ScssParserUrlStart)
		}
		{
			p.SetState(419)
			p.Match(ScssParserUrl)
		}
		{
			p.SetState(420)
			p.Match(ScssParserUrlEnd)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAsClauseContext is an interface to support dynamic dispatch.
type IAsClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAsClauseContext differentiates from other interfaces.
	IsAsClauseContext()
}

type AsClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsClauseContext() *AsClauseContext {
	var p = new(AsClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_asClause
	return p
}

func (*AsClauseContext) IsAsClauseContext() {}

func NewAsClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsClauseContext {
	var p = new(AsClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_asClause

	return p
}

func (s *AsClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *AsClauseContext) AS() antlr.TerminalNode {
	return s.GetToken(ScssParserAS, 0)
}

func (s *AsClauseContext) TIMES() antlr.TerminalNode {
	return s.GetToken(ScssParserTIMES, 0)
}

func (s *AsClauseContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AsClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterAsClause(s)
	}
}

func (s *AsClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitAsClause(s)
	}
}

func (p *ScssParser) AsClause() (localctx IAsClauseContext) {
	this := p
	_ = this

	localctx = NewAsClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ScssParserRULE_asClause)

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
		p.SetState(423)
		p.Match(ScssParserAS)
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserTIMES:
		{
			p.SetState(424)
			p.Match(ScssParserTIMES)
		}

	case ScssParserInterpolationStart, ScssParserFROM, ScssParserTO, ScssParserTHROUGH, ScssParserONLY, ScssParserNOT, ScssParserAND_WORD, ScssParserUSING, ScssParserIdentifier:
		{
			p.SetState(425)
			p.Identifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWithClauseContext is an interface to support dynamic dispatch.
type IWithClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWithClauseContext differentiates from other interfaces.
	IsWithClauseContext()
}

type WithClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithClauseContext() *WithClauseContext {
	var p = new(WithClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_withClause
	return p
}

func (*WithClauseContext) IsWithClauseContext() {}

func NewWithClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithClauseContext {
	var p = new(WithClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_withClause

	return p
}

func (s *WithClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WithClauseContext) WITH() antlr.TerminalNode {
	return s.GetToken(ScssParserWITH, 0)
}

func (s *WithClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *WithClauseContext) AllKeywordArgument() []IKeywordArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeywordArgumentContext)(nil)).Elem())
	var tst = make([]IKeywordArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeywordArgumentContext)
		}
	}

	return tst
}

func (s *WithClauseContext) KeywordArgument(i int) IKeywordArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeywordArgumentContext)
}

func (s *WithClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *WithClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *WithClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *WithClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterWithClause(s)
	}
}

func (s *WithClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitWithClause(s)
	}
}

func (p *ScssParser) WithClause() (localctx IWithClauseContext) {
	this := p
	_ = this

	localctx = NewWithClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ScssParserRULE_withClause)
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
		p.SetState(428)
		p.Match(ScssParserWITH)
	}
	{
		p.SetState(429)
		p.Match(ScssParserLPAREN)
	}
	{
		p.SetState(430)
		p.KeywordArgument()
	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(431)
				p.Match(ScssParserCOMMA)
			}
			{
				p.SetState(432)
				p.KeywordArgument()
			}

		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserCOMMA {
		{
			p.SetState(438)
			p.Match(ScssParserCOMMA)
		}

	}
	{
		p.SetState(441)
		p.Match(ScssParserRPAREN)
	}

	return localctx
}

// IKeywordArgumentContext is an interface to support dynamic dispatch.
type IKeywordArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordArgumentContext differentiates from other interfaces.
	IsKeywordArgumentContext()
}

type KeywordArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordArgumentContext() *KeywordArgumentContext {
	var p = new(KeywordArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_keywordArgument
	return p
}

func (*KeywordArgumentContext) IsKeywordArgumentContext() {}

func NewKeywordArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordArgumentContext {
	var p = new(KeywordArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_keywordArgument

	return p
}

func (s *KeywordArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordArgumentContext) IdentifierVariableName() IIdentifierVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierVariableNameContext)
}

func (s *KeywordArgumentContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *KeywordArgumentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *KeywordArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterKeywordArgument(s)
	}
}

func (s *KeywordArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitKeywordArgument(s)
	}
}

func (p *ScssParser) KeywordArgument() (localctx IKeywordArgumentContext) {
	this := p
	_ = this

	localctx = NewKeywordArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ScssParserRULE_keywordArgument)

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
		p.SetState(443)
		p.IdentifierVariableName()
	}
	{
		p.SetState(444)
		p.Match(ScssParserCOLON)
	}
	{
		p.SetState(445)
		p.Expression()
	}

	return localctx
}

// IMediaDeclarationContext is an interface to support dynamic dispatch.
type IMediaDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediaDeclarationContext differentiates from other interfaces.
	IsMediaDeclarationContext()
}

type MediaDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediaDeclarationContext() *MediaDeclarationContext {
	var p = new(MediaDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mediaDeclaration
	return p
}

func (*MediaDeclarationContext) IsMediaDeclarationContext() {}

func NewMediaDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediaDeclarationContext {
	var p = new(MediaDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mediaDeclaration

	return p
}

func (s *MediaDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MediaDeclarationContext) MEDIA() antlr.TerminalNode {
	return s.GetToken(ScssParserMEDIA, 0)
}

func (s *MediaDeclarationContext) MediaQueryList() IMediaQueryListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMediaQueryListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMediaQueryListContext)
}

func (s *MediaDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *MediaDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediaDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediaDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMediaDeclaration(s)
	}
}

func (s *MediaDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMediaDeclaration(s)
	}
}

func (p *ScssParser) MediaDeclaration() (localctx IMediaDeclarationContext) {
	this := p
	_ = this

	localctx = NewMediaDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ScssParserRULE_mediaDeclaration)

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
		p.SetState(447)
		p.Match(ScssParserMEDIA)
	}
	{
		p.SetState(448)
		p.MediaQueryList()
	}
	{
		p.SetState(449)
		p.Block()
	}

	return localctx
}

// IMediaQueryListContext is an interface to support dynamic dispatch.
type IMediaQueryListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediaQueryListContext differentiates from other interfaces.
	IsMediaQueryListContext()
}

type MediaQueryListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediaQueryListContext() *MediaQueryListContext {
	var p = new(MediaQueryListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mediaQueryList
	return p
}

func (*MediaQueryListContext) IsMediaQueryListContext() {}

func NewMediaQueryListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediaQueryListContext {
	var p = new(MediaQueryListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mediaQueryList

	return p
}

func (s *MediaQueryListContext) GetParser() antlr.Parser { return s.parser }

func (s *MediaQueryListContext) AllMediaQuery() []IMediaQueryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMediaQueryContext)(nil)).Elem())
	var tst = make([]IMediaQueryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMediaQueryContext)
		}
	}

	return tst
}

func (s *MediaQueryListContext) MediaQuery(i int) IMediaQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMediaQueryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMediaQueryContext)
}

func (s *MediaQueryListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *MediaQueryListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *MediaQueryListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediaQueryListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediaQueryListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMediaQueryList(s)
	}
}

func (s *MediaQueryListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMediaQueryList(s)
	}
}

func (p *ScssParser) MediaQueryList() (localctx IMediaQueryListContext) {
	this := p
	_ = this

	localctx = NewMediaQueryListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ScssParserRULE_mediaQueryList)
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
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserLPAREN || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(ScssParserONLY-58))|(1<<(ScssParserNOT-58))|(1<<(ScssParserIdentifier-58)))) != 0) {
		{
			p.SetState(451)
			p.MediaQuery()
		}
		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScssParserCOMMA {
			{
				p.SetState(452)
				p.Match(ScssParserCOMMA)
			}
			{
				p.SetState(453)
				p.MediaQuery()
			}

			p.SetState(458)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IMediaQueryContext is an interface to support dynamic dispatch.
type IMediaQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediaQueryContext differentiates from other interfaces.
	IsMediaQueryContext()
}

type MediaQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediaQueryContext() *MediaQueryContext {
	var p = new(MediaQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mediaQuery
	return p
}

func (*MediaQueryContext) IsMediaQueryContext() {}

func NewMediaQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediaQueryContext {
	var p = new(MediaQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mediaQuery

	return p
}

func (s *MediaQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *MediaQueryContext) MediaType() IMediaTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMediaTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMediaTypeContext)
}

func (s *MediaQueryContext) AllAND_WORD() []antlr.TerminalNode {
	return s.GetTokens(ScssParserAND_WORD)
}

func (s *MediaQueryContext) AND_WORD(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserAND_WORD, i)
}

func (s *MediaQueryContext) AllMediaExpression() []IMediaExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMediaExpressionContext)(nil)).Elem())
	var tst = make([]IMediaExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMediaExpressionContext)
		}
	}

	return tst
}

func (s *MediaQueryContext) MediaExpression(i int) IMediaExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMediaExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMediaExpressionContext)
}

func (s *MediaQueryContext) ONLY() antlr.TerminalNode {
	return s.GetToken(ScssParserONLY, 0)
}

func (s *MediaQueryContext) NOT() antlr.TerminalNode {
	return s.GetToken(ScssParserNOT, 0)
}

func (s *MediaQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediaQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediaQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMediaQuery(s)
	}
}

func (s *MediaQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMediaQuery(s)
	}
}

func (p *ScssParser) MediaQuery() (localctx IMediaQueryContext) {
	this := p
	_ = this

	localctx = NewMediaQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ScssParserRULE_mediaQuery)
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

	p.SetState(480)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserONLY, ScssParserNOT, ScssParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(462)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScssParserONLY || _la == ScssParserNOT {
			{
				p.SetState(461)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ScssParserONLY || _la == ScssParserNOT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(464)
			p.MediaType()
		}
		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScssParserAND_WORD {
			{
				p.SetState(465)
				p.Match(ScssParserAND_WORD)
			}
			{
				p.SetState(466)
				p.MediaExpression()
			}

			p.SetState(471)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ScssParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(472)
			p.MediaExpression()
		}
		p.SetState(477)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScssParserAND_WORD {
			{
				p.SetState(473)
				p.Match(ScssParserAND_WORD)
			}
			{
				p.SetState(474)
				p.MediaExpression()
			}

			p.SetState(479)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMediaTypeContext is an interface to support dynamic dispatch.
type IMediaTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediaTypeContext differentiates from other interfaces.
	IsMediaTypeContext()
}

type MediaTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediaTypeContext() *MediaTypeContext {
	var p = new(MediaTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mediaType
	return p
}

func (*MediaTypeContext) IsMediaTypeContext() {}

func NewMediaTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediaTypeContext {
	var p = new(MediaTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mediaType

	return p
}

func (s *MediaTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MediaTypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *MediaTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediaTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediaTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMediaType(s)
	}
}

func (s *MediaTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMediaType(s)
	}
}

func (p *ScssParser) MediaType() (localctx IMediaTypeContext) {
	this := p
	_ = this

	localctx = NewMediaTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ScssParserRULE_mediaType)

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
		p.SetState(482)
		p.Match(ScssParserIdentifier)
	}

	return localctx
}

// IMediaExpressionContext is an interface to support dynamic dispatch.
type IMediaExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediaExpressionContext differentiates from other interfaces.
	IsMediaExpressionContext()
}

type MediaExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediaExpressionContext() *MediaExpressionContext {
	var p = new(MediaExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mediaExpression
	return p
}

func (*MediaExpressionContext) IsMediaExpressionContext() {}

func NewMediaExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediaExpressionContext {
	var p = new(MediaExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mediaExpression

	return p
}

func (s *MediaExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MediaExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *MediaExpressionContext) MediaFeature() IMediaFeatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMediaFeatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMediaFeatureContext)
}

func (s *MediaExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *MediaExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *MediaExpressionContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *MediaExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediaExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediaExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMediaExpression(s)
	}
}

func (s *MediaExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMediaExpression(s)
	}
}

func (p *ScssParser) MediaExpression() (localctx IMediaExpressionContext) {
	this := p
	_ = this

	localctx = NewMediaExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ScssParserRULE_mediaExpression)
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
		p.SetState(484)
		p.Match(ScssParserLPAREN)
	}
	{
		p.SetState(485)
		p.MediaFeature()
	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserCOLON {
		{
			p.SetState(486)
			p.Match(ScssParserCOLON)
		}
		{
			p.SetState(487)
			p.CommandStatement()
		}

	}
	{
		p.SetState(490)
		p.Match(ScssParserRPAREN)
	}

	return localctx
}

// IMediaFeatureContext is an interface to support dynamic dispatch.
type IMediaFeatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediaFeatureContext differentiates from other interfaces.
	IsMediaFeatureContext()
}

type MediaFeatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediaFeatureContext() *MediaFeatureContext {
	var p = new(MediaFeatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mediaFeature
	return p
}

func (*MediaFeatureContext) IsMediaFeatureContext() {}

func NewMediaFeatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediaFeatureContext {
	var p = new(MediaFeatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mediaFeature

	return p
}

func (s *MediaFeatureContext) GetParser() antlr.Parser { return s.parser }

func (s *MediaFeatureContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *MediaFeatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediaFeatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediaFeatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMediaFeature(s)
	}
}

func (s *MediaFeatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMediaFeature(s)
	}
}

func (p *ScssParser) MediaFeature() (localctx IMediaFeatureContext) {
	this := p
	_ = this

	localctx = NewMediaFeatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ScssParserRULE_mediaFeature)

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
		p.SetState(492)
		p.Match(ScssParserIdentifier)
	}

	return localctx
}

// IRulesetContext is an interface to support dynamic dispatch.
type IRulesetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulesetContext differentiates from other interfaces.
	IsRulesetContext()
}

type RulesetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesetContext() *RulesetContext {
	var p = new(RulesetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_ruleset
	return p
}

func (*RulesetContext) IsRulesetContext() {}

func NewRulesetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesetContext {
	var p = new(RulesetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_ruleset

	return p
}

func (s *RulesetContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesetContext) Selectors() ISelectorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorsContext)
}

func (s *RulesetContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RulesetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterRuleset(s)
	}
}

func (s *RulesetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitRuleset(s)
	}
}

func (p *ScssParser) Ruleset() (localctx IRulesetContext) {
	this := p
	_ = this

	localctx = NewRulesetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ScssParserRULE_ruleset)

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
		p.SetState(494)
		p.Selectors()
	}
	{
		p.SetState(495)
		p.Block()
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) BlockStart() antlr.TerminalNode {
	return s.GetToken(ScssParserBlockStart, 0)
}

func (s *BlockContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(ScssParserBlockEnd, 0)
}

func (s *BlockContext) AllProperty_() []IProperty_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProperty_Context)(nil)).Elem())
	var tst = make([]IProperty_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProperty_Context)
		}
	}

	return tst
}

func (s *BlockContext) Property_(i int) IProperty_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProperty_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProperty_Context)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) LastProperty() ILastPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILastPropertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILastPropertyContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *ScssParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ScssParserRULE_block)
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
		p.SetState(497)
		p.Match(ScssParserBlockStart)
	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(500)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(498)
					p.Property_()
				}

			case 2:
				{
					p.SetState(499)
					p.Statement()
				}

			}

		}
		p.SetState(504)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext())
	}
	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserInterpolationStart || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(ScssParserFROM-53))|(1<<(ScssParserTO-53))|(1<<(ScssParserTHROUGH-53))|(1<<(ScssParserONLY-53))|(1<<(ScssParserNOT-53))|(1<<(ScssParserAND_WORD-53))|(1<<(ScssParserUSING-53))|(1<<(ScssParserIdentifier-53)))) != 0) {
		{
			p.SetState(505)
			p.LastProperty()
		}

	}
	{
		p.SetState(508)
		p.Match(ScssParserBlockEnd)
	}

	return localctx
}

// ISelectorsContext is an interface to support dynamic dispatch.
type ISelectorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorsContext differentiates from other interfaces.
	IsSelectorsContext()
}

type SelectorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorsContext() *SelectorsContext {
	var p = new(SelectorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_selectors
	return p
}

func (*SelectorsContext) IsSelectorsContext() {}

func NewSelectorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorsContext {
	var p = new(SelectorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_selectors

	return p
}

func (s *SelectorsContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorsContext) AllSelector() []ISelectorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectorContext)(nil)).Elem())
	var tst = make([]ISelectorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectorContext)
		}
	}

	return tst
}

func (s *SelectorsContext) Selector(i int) ISelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *SelectorsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *SelectorsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *SelectorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterSelectors(s)
	}
}

func (s *SelectorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitSelectors(s)
	}
}

func (p *ScssParser) Selectors() (localctx ISelectorsContext) {
	this := p
	_ = this

	localctx = NewSelectorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ScssParserRULE_selectors)
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
		p.SetState(510)
		p.Selector()
	}
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserCOMMA {
		{
			p.SetState(511)
			p.Match(ScssParserCOMMA)
		}
		{
			p.SetState(512)
			p.Selector()
		}

		p.SetState(517)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_selector
	return p
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *SelectorContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (p *ScssParser) Selector() (localctx ISelectorContext) {
	this := p
	_ = this

	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ScssParserRULE_selector)
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
	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserInterpolationStart)|(1<<ScssParserLBRACK)|(1<<ScssParserGT)|(1<<ScssParserTIL)|(1<<ScssParserDOT)|(1<<ScssParserAND)|(1<<ScssParserHASH)|(1<<ScssParserPLUS)|(1<<ScssParserTIMES))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(ScssParserFROM-53))|(1<<(ScssParserTO-53))|(1<<(ScssParserTHROUGH-53))|(1<<(ScssParserONLY-53))|(1<<(ScssParserNOT-53))|(1<<(ScssParserAND_WORD-53))|(1<<(ScssParserUSING-53))|(1<<(ScssParserIdentifier-53))|(1<<(ScssParserPseudoIdentifier-53)))) != 0) {
		{
			p.SetState(518)
			p.Element()
		}

		p.SetState(521)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = ScssParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ElementContext) HASH() antlr.TerminalNode {
	return s.GetToken(ScssParserHASH, 0)
}

func (s *ElementContext) DOT() antlr.TerminalNode {
	return s.GetToken(ScssParserDOT, 0)
}

func (s *ElementContext) AND() antlr.TerminalNode {
	return s.GetToken(ScssParserAND, 0)
}

func (s *ElementContext) TIMES() antlr.TerminalNode {
	return s.GetToken(ScssParserTIMES, 0)
}

func (s *ElementContext) Combinator() ICombinatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICombinatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICombinatorContext)
}

func (s *ElementContext) Attrib() IAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *ElementContext) Pseudo() IPseudoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPseudoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPseudoContext)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitElement(s)
	}
}

func (p *ScssParser) Element() (localctx IElementContext) {
	this := p
	_ = this

	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ScssParserRULE_element)

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

	p.SetState(533)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserInterpolationStart, ScssParserFROM, ScssParserTO, ScssParserTHROUGH, ScssParserONLY, ScssParserNOT, ScssParserAND_WORD, ScssParserUSING, ScssParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(523)
			p.Identifier()
		}

	case ScssParserHASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(524)
			p.Match(ScssParserHASH)
		}
		{
			p.SetState(525)
			p.Identifier()
		}

	case ScssParserDOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(526)
			p.Match(ScssParserDOT)
		}
		{
			p.SetState(527)
			p.Identifier()
		}

	case ScssParserAND:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(528)
			p.Match(ScssParserAND)
		}

	case ScssParserTIMES:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(529)
			p.Match(ScssParserTIMES)
		}

	case ScssParserGT, ScssParserTIL, ScssParserPLUS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(530)
			p.Combinator()
		}

	case ScssParserLBRACK:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(531)
			p.Attrib()
		}

	case ScssParserPseudoIdentifier:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(532)
			p.Pseudo()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICombinatorContext is an interface to support dynamic dispatch.
type ICombinatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCombinatorContext differentiates from other interfaces.
	IsCombinatorContext()
}

type CombinatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCombinatorContext() *CombinatorContext {
	var p = new(CombinatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_combinator
	return p
}

func (*CombinatorContext) IsCombinatorContext() {}

func NewCombinatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CombinatorContext {
	var p = new(CombinatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_combinator

	return p
}

func (s *CombinatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CombinatorContext) GT() antlr.TerminalNode {
	return s.GetToken(ScssParserGT, 0)
}

func (s *CombinatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ScssParserPLUS, 0)
}

func (s *CombinatorContext) TIL() antlr.TerminalNode {
	return s.GetToken(ScssParserTIL, 0)
}

func (s *CombinatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CombinatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CombinatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterCombinator(s)
	}
}

func (s *CombinatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitCombinator(s)
	}
}

func (p *ScssParser) Combinator() (localctx ICombinatorContext) {
	this := p
	_ = this

	localctx = NewCombinatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ScssParserRULE_combinator)
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
		p.SetState(535)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserGT)|(1<<ScssParserTIL)|(1<<ScssParserPLUS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPseudoContext is an interface to support dynamic dispatch.
type IPseudoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPseudoContext differentiates from other interfaces.
	IsPseudoContext()
}

type PseudoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPseudoContext() *PseudoContext {
	var p = new(PseudoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_pseudo
	return p
}

func (*PseudoContext) IsPseudoContext() {}

func NewPseudoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PseudoContext {
	var p = new(PseudoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_pseudo

	return p
}

func (s *PseudoContext) GetParser() antlr.Parser { return s.parser }

func (s *PseudoContext) PseudoIdentifier() IPseudoIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPseudoIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPseudoIdentifierContext)
}

func (s *PseudoContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *PseudoContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *PseudoContext) Selector() ISelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *PseudoContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *PseudoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PseudoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PseudoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterPseudo(s)
	}
}

func (s *PseudoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitPseudo(s)
	}
}

func (p *ScssParser) Pseudo() (localctx IPseudoContext) {
	this := p
	_ = this

	localctx = NewPseudoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ScssParserRULE_pseudo)

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

	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(537)
			p.PseudoIdentifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(538)
			p.PseudoIdentifier()
		}
		{
			p.SetState(539)
			p.Match(ScssParserLPAREN)
		}
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(540)
				p.Selector()
			}

		case 2:
			{
				p.SetState(541)
				p.CommandStatement()
			}

		}
		{
			p.SetState(544)
			p.Match(ScssParserRPAREN)
		}

	}

	return localctx
}

// IAttribContext is an interface to support dynamic dispatch.
type IAttribContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribContext differentiates from other interfaces.
	IsAttribContext()
}

type AttribContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribContext() *AttribContext {
	var p = new(AttribContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_attrib
	return p
}

func (*AttribContext) IsAttribContext() {}

func NewAttribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribContext {
	var p = new(AttribContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_attrib

	return p
}

func (s *AttribContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ScssParserLBRACK, 0)
}

func (s *AttribContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(ScssParserIdentifier)
}

func (s *AttribContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, i)
}

func (s *AttribContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ScssParserRBRACK, 0)
}

func (s *AttribContext) AttribRelate() IAttribRelateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribRelateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribRelateContext)
}

func (s *AttribContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ScssParserStringLiteral, 0)
}

func (s *AttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterAttrib(s)
	}
}

func (s *AttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitAttrib(s)
	}
}

func (p *ScssParser) Attrib() (localctx IAttribContext) {
	this := p
	_ = this

	localctx = NewAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ScssParserRULE_attrib)
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
		p.SetState(548)
		p.Match(ScssParserLBRACK)
	}
	{
		p.SetState(549)
		p.Match(ScssParserIdentifier)
	}
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ScssParserEQ-36))|(1<<(ScssParserPIPE_EQ-36))|(1<<(ScssParserTILD_EQ-36)))) != 0 {
		{
			p.SetState(550)
			p.AttribRelate()
		}
		{
			p.SetState(551)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ScssParserIdentifier || _la == ScssParserStringLiteral) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(555)
		p.Match(ScssParserRBRACK)
	}

	return localctx
}

// IAttribRelateContext is an interface to support dynamic dispatch.
type IAttribRelateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribRelateContext differentiates from other interfaces.
	IsAttribRelateContext()
}

type AttribRelateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribRelateContext() *AttribRelateContext {
	var p = new(AttribRelateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_attribRelate
	return p
}

func (*AttribRelateContext) IsAttribRelateContext() {}

func NewAttribRelateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribRelateContext {
	var p = new(AttribRelateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_attribRelate

	return p
}

func (s *AttribRelateContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribRelateContext) EQ() antlr.TerminalNode {
	return s.GetToken(ScssParserEQ, 0)
}

func (s *AttribRelateContext) PIPE_EQ() antlr.TerminalNode {
	return s.GetToken(ScssParserPIPE_EQ, 0)
}

func (s *AttribRelateContext) TILD_EQ() antlr.TerminalNode {
	return s.GetToken(ScssParserTILD_EQ, 0)
}

func (s *AttribRelateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribRelateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribRelateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterAttribRelate(s)
	}
}

func (s *AttribRelateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitAttribRelate(s)
	}
}

func (p *ScssParser) AttribRelate() (localctx IAttribRelateContext) {
	this := p
	_ = this

	localctx = NewAttribRelateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ScssParserRULE_attribRelate)
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
		p.SetState(557)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(ScssParserEQ-36))|(1<<(ScssParserPIPE_EQ-36))|(1<<(ScssParserTILD_EQ-36)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.RuleIndex = ScssParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *IdentifierContext) AllIdentifierPart() []IIdentifierPartContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierPartContext)(nil)).Elem())
	var tst = make([]IIdentifierPartContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierPartContext)
		}
	}

	return tst
}

func (s *IdentifierContext) IdentifierPart(i int) IIdentifierPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierPartContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierPartContext)
}

func (s *IdentifierContext) InterpolationStart() antlr.TerminalNode {
	return s.GetToken(ScssParserInterpolationStart, 0)
}

func (s *IdentifierContext) IdentifierVariableName() IIdentifierVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierVariableNameContext)
}

func (s *IdentifierContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(ScssParserBlockEnd, 0)
}

func (s *IdentifierContext) AND_WORD() antlr.TerminalNode {
	return s.GetToken(ScssParserAND_WORD, 0)
}

func (s *IdentifierContext) FROM() antlr.TerminalNode {
	return s.GetToken(ScssParserFROM, 0)
}

func (s *IdentifierContext) NOT() antlr.TerminalNode {
	return s.GetToken(ScssParserNOT, 0)
}

func (s *IdentifierContext) ONLY() antlr.TerminalNode {
	return s.GetToken(ScssParserONLY, 0)
}

func (s *IdentifierContext) THROUGH() antlr.TerminalNode {
	return s.GetToken(ScssParserTHROUGH, 0)
}

func (s *IdentifierContext) TO() antlr.TerminalNode {
	return s.GetToken(ScssParserTO, 0)
}

func (s *IdentifierContext) USING() antlr.TerminalNode {
	return s.GetToken(ScssParserUSING, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *ScssParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ScssParserRULE_identifier)
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

	p.SetState(582)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(559)
			p.Match(ScssParserIdentifier)
		}
		p.SetState(563)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScssParserInterpolationStartAfter || _la == ScssParserIdentifierAfter {
			{
				p.SetState(560)
				p.IdentifierPart()
			}

			p.SetState(565)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ScssParserInterpolationStart:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(566)
			p.Match(ScssParserInterpolationStart)
		}
		{
			p.SetState(567)
			p.IdentifierVariableName()
		}
		{
			p.SetState(568)
			p.Match(ScssParserBlockEnd)
		}
		p.SetState(572)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ScssParserInterpolationStartAfter || _la == ScssParserIdentifierAfter {
			{
				p.SetState(569)
				p.IdentifierPart()
			}

			p.SetState(574)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ScssParserAND_WORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(575)
			p.Match(ScssParserAND_WORD)
		}

	case ScssParserFROM:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(576)
			p.Match(ScssParserFROM)
		}

	case ScssParserNOT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(577)
			p.Match(ScssParserNOT)
		}

	case ScssParserONLY:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(578)
			p.Match(ScssParserONLY)
		}

	case ScssParserTHROUGH:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(579)
			p.Match(ScssParserTHROUGH)
		}

	case ScssParserTO:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(580)
			p.Match(ScssParserTO)
		}

	case ScssParserUSING:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(581)
			p.Match(ScssParserUSING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPseudoIdentifierContext is an interface to support dynamic dispatch.
type IPseudoIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPseudoIdentifierContext differentiates from other interfaces.
	IsPseudoIdentifierContext()
}

type PseudoIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPseudoIdentifierContext() *PseudoIdentifierContext {
	var p = new(PseudoIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_pseudoIdentifier
	return p
}

func (*PseudoIdentifierContext) IsPseudoIdentifierContext() {}

func NewPseudoIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PseudoIdentifierContext {
	var p = new(PseudoIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_pseudoIdentifier

	return p
}

func (s *PseudoIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *PseudoIdentifierContext) PseudoIdentifier() antlr.TerminalNode {
	return s.GetToken(ScssParserPseudoIdentifier, 0)
}

func (s *PseudoIdentifierContext) AllIdentifierPart() []IIdentifierPartContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierPartContext)(nil)).Elem())
	var tst = make([]IIdentifierPartContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierPartContext)
		}
	}

	return tst
}

func (s *PseudoIdentifierContext) IdentifierPart(i int) IIdentifierPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierPartContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierPartContext)
}

func (s *PseudoIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PseudoIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PseudoIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterPseudoIdentifier(s)
	}
}

func (s *PseudoIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitPseudoIdentifier(s)
	}
}

func (p *ScssParser) PseudoIdentifier() (localctx IPseudoIdentifierContext) {
	this := p
	_ = this

	localctx = NewPseudoIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, ScssParserRULE_pseudoIdentifier)
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
		p.SetState(584)
		p.Match(ScssParserPseudoIdentifier)
	}
	p.SetState(588)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ScssParserInterpolationStartAfter || _la == ScssParserIdentifierAfter {
		{
			p.SetState(585)
			p.IdentifierPart()
		}

		p.SetState(590)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIdentifierPartContext is an interface to support dynamic dispatch.
type IIdentifierPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierPartContext differentiates from other interfaces.
	IsIdentifierPartContext()
}

type IdentifierPartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierPartContext() *IdentifierPartContext {
	var p = new(IdentifierPartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_identifierPart
	return p
}

func (*IdentifierPartContext) IsIdentifierPartContext() {}

func NewIdentifierPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierPartContext {
	var p = new(IdentifierPartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_identifierPart

	return p
}

func (s *IdentifierPartContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierPartContext) InterpolationStartAfter() antlr.TerminalNode {
	return s.GetToken(ScssParserInterpolationStartAfter, 0)
}

func (s *IdentifierPartContext) IdentifierVariableName() IIdentifierVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierVariableNameContext)
}

func (s *IdentifierPartContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(ScssParserBlockEnd, 0)
}

func (s *IdentifierPartContext) IdentifierAfter() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifierAfter, 0)
}

func (s *IdentifierPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterIdentifierPart(s)
	}
}

func (s *IdentifierPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitIdentifierPart(s)
	}
}

func (p *ScssParser) IdentifierPart() (localctx IIdentifierPartContext) {
	this := p
	_ = this

	localctx = NewIdentifierPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, ScssParserRULE_identifierPart)

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

	p.SetState(596)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScssParserInterpolationStartAfter:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(591)
			p.Match(ScssParserInterpolationStartAfter)
		}
		{
			p.SetState(592)
			p.IdentifierVariableName()
		}
		{
			p.SetState(593)
			p.Match(ScssParserBlockEnd)
		}

	case ScssParserIdentifierAfter:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(595)
			p.Match(ScssParserIdentifierAfter)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierVariableNameContext is an interface to support dynamic dispatch.
type IIdentifierVariableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierVariableNameContext differentiates from other interfaces.
	IsIdentifierVariableNameContext()
}

type IdentifierVariableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierVariableNameContext() *IdentifierVariableNameContext {
	var p = new(IdentifierVariableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_identifierVariableName
	return p
}

func (*IdentifierVariableNameContext) IsIdentifierVariableNameContext() {}

func NewIdentifierVariableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierVariableNameContext {
	var p = new(IdentifierVariableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_identifierVariableName

	return p
}

func (s *IdentifierVariableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierVariableNameContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(ScssParserDOLLAR, 0)
}

func (s *IdentifierVariableNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, 0)
}

func (s *IdentifierVariableNameContext) IdentifierAfter() antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifierAfter, 0)
}

func (s *IdentifierVariableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierVariableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierVariableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterIdentifierVariableName(s)
	}
}

func (s *IdentifierVariableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitIdentifierVariableName(s)
	}
}

func (p *ScssParser) IdentifierVariableName() (localctx IIdentifierVariableNameContext) {
	this := p
	_ = this

	localctx = NewIdentifierVariableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, ScssParserRULE_identifierVariableName)
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
		p.SetState(598)
		p.Match(ScssParserDOLLAR)
	}
	{
		p.SetState(599)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ScssParserIdentifier || _la == ScssParserIdentifierAfter) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IProperty_Context is an interface to support dynamic dispatch.
type IProperty_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProperty_Context differentiates from other interfaces.
	IsProperty_Context()
}

type Property_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_Context() *Property_Context {
	var p = new(Property_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_property_
	return p
}

func (*Property_Context) IsProperty_Context() {}

func NewProperty_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_Context {
	var p = new(Property_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_property_

	return p
}

func (s *Property_Context) GetParser() antlr.Parser { return s.parser }

func (s *Property_Context) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Property_Context) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *Property_Context) PropertyValue() IPropertyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyValueContext)
}

func (s *Property_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(ScssParserSEMI, 0)
}

func (s *Property_Context) IMPORTANT() antlr.TerminalNode {
	return s.GetToken(ScssParserIMPORTANT, 0)
}

func (s *Property_Context) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Property_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterProperty_(s)
	}
}

func (s *Property_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitProperty_(s)
	}
}

func (p *ScssParser) Property_() (localctx IProperty_Context) {
	this := p
	_ = this

	localctx = NewProperty_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, ScssParserRULE_property_)
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

	p.SetState(621)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(601)
			p.Identifier()
		}
		{
			p.SetState(602)
			p.Match(ScssParserCOLON)
		}
		{
			p.SetState(603)
			p.PropertyValue()
		}
		p.SetState(605)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScssParserIMPORTANT {
			{
				p.SetState(604)
				p.Match(ScssParserIMPORTANT)
			}

		}
		{
			p.SetState(607)
			p.Match(ScssParserSEMI)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(609)
			p.Identifier()
		}
		{
			p.SetState(610)
			p.Match(ScssParserCOLON)
		}
		{
			p.SetState(611)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(613)
			p.Identifier()
		}
		{
			p.SetState(614)
			p.Match(ScssParserCOLON)
		}
		{
			p.SetState(615)
			p.PropertyValue()
		}
		p.SetState(617)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScssParserIMPORTANT {
			{
				p.SetState(616)
				p.Match(ScssParserIMPORTANT)
			}

		}
		{
			p.SetState(619)
			p.Block()
		}

	}

	return localctx
}

// ILastPropertyContext is an interface to support dynamic dispatch.
type ILastPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLastPropertyContext differentiates from other interfaces.
	IsLastPropertyContext()
}

type LastPropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLastPropertyContext() *LastPropertyContext {
	var p = new(LastPropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_lastProperty
	return p
}

func (*LastPropertyContext) IsLastPropertyContext() {}

func NewLastPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LastPropertyContext {
	var p = new(LastPropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_lastProperty

	return p
}

func (s *LastPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *LastPropertyContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *LastPropertyContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *LastPropertyContext) PropertyValue() IPropertyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyValueContext)
}

func (s *LastPropertyContext) IMPORTANT() antlr.TerminalNode {
	return s.GetToken(ScssParserIMPORTANT, 0)
}

func (s *LastPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LastPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LastPropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterLastProperty(s)
	}
}

func (s *LastPropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitLastProperty(s)
	}
}

func (p *ScssParser) LastProperty() (localctx ILastPropertyContext) {
	this := p
	_ = this

	localctx = NewLastPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, ScssParserRULE_lastProperty)
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
		p.SetState(623)
		p.Identifier()
	}
	{
		p.SetState(624)
		p.Match(ScssParserCOLON)
	}
	{
		p.SetState(625)
		p.PropertyValue()
	}
	p.SetState(627)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserIMPORTANT {
		{
			p.SetState(626)
			p.Match(ScssParserIMPORTANT)
		}

	}

	return localctx
}

// IPropertyValueContext is an interface to support dynamic dispatch.
type IPropertyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyValueContext differentiates from other interfaces.
	IsPropertyValueContext()
}

type PropertyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyValueContext() *PropertyValueContext {
	var p = new(PropertyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_propertyValue
	return p
}

func (*PropertyValueContext) IsPropertyValueContext() {}

func NewPropertyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyValueContext {
	var p = new(PropertyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_propertyValue

	return p
}

func (s *PropertyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyValueContext) AllCommandStatement() []ICommandStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem())
	var tst = make([]ICommandStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandStatementContext)
		}
	}

	return tst
}

func (s *PropertyValueContext) CommandStatement(i int) ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *PropertyValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *PropertyValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *PropertyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterPropertyValue(s)
	}
}

func (s *PropertyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitPropertyValue(s)
	}
}

func (p *ScssParser) PropertyValue() (localctx IPropertyValueContext) {
	this := p
	_ = this

	localctx = NewPropertyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, ScssParserRULE_propertyValue)
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
		p.SetState(629)
		p.CommandStatement()
	}
	p.SetState(636)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL_)|(1<<ScssParserInterpolationStart)|(1<<ScssParserLPAREN)|(1<<ScssParserCOMMA)|(1<<ScssParserDOLLAR)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR)|(1<<ScssParserMINUS_LPAREN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ScssParserPLUS_LPAREN-32))|(1<<(ScssParserUrlStart-32))|(1<<(ScssParserFROM-32))|(1<<(ScssParserTO-32))|(1<<(ScssParserTHROUGH-32))|(1<<(ScssParserONLY-32))|(1<<(ScssParserNOT-32))|(1<<(ScssParserAND_WORD-32))|(1<<(ScssParserUSING-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(ScssParserIdentifier-64))|(1<<(ScssParserFunctionIdentifier-64))|(1<<(ScssParserStringLiteral-64))|(1<<(ScssParserNumber-64))|(1<<(ScssParserColor-64)))) != 0) {
		p.SetState(631)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ScssParserCOMMA {
			{
				p.SetState(630)
				p.Match(ScssParserCOMMA)
			}

		}
		{
			p.SetState(633)
			p.CommandStatement()
		}

		p.SetState(638)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUrlContext is an interface to support dynamic dispatch.
type IUrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUrlContext differentiates from other interfaces.
	IsUrlContext()
}

type UrlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrlContext() *UrlContext {
	var p = new(UrlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_url
	return p
}

func (*UrlContext) IsUrlContext() {}

func NewUrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UrlContext {
	var p = new(UrlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_url

	return p
}

func (s *UrlContext) GetParser() antlr.Parser { return s.parser }

func (s *UrlContext) UrlStart() antlr.TerminalNode {
	return s.GetToken(ScssParserUrlStart, 0)
}

func (s *UrlContext) Url() antlr.TerminalNode {
	return s.GetToken(ScssParserUrl, 0)
}

func (s *UrlContext) UrlEnd() antlr.TerminalNode {
	return s.GetToken(ScssParserUrlEnd, 0)
}

func (s *UrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterUrl(s)
	}
}

func (s *UrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitUrl(s)
	}
}

func (p *ScssParser) Url() (localctx IUrlContext) {
	this := p
	_ = this

	localctx = NewUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, ScssParserRULE_url)

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
		p.SetState(639)
		p.Match(ScssParserUrlStart)
	}
	{
		p.SetState(640)
		p.Match(ScssParserUrl)
	}
	{
		p.SetState(641)
		p.Match(ScssParserUrlEnd)
	}

	return localctx
}

// IMeasurementContext is an interface to support dynamic dispatch.
type IMeasurementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeasurementContext differentiates from other interfaces.
	IsMeasurementContext()
}

type MeasurementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeasurementContext() *MeasurementContext {
	var p = new(MeasurementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_measurement
	return p
}

func (*MeasurementContext) IsMeasurementContext() {}

func NewMeasurementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MeasurementContext {
	var p = new(MeasurementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_measurement

	return p
}

func (s *MeasurementContext) GetParser() antlr.Parser { return s.parser }

func (s *MeasurementContext) Number() antlr.TerminalNode {
	return s.GetToken(ScssParserNumber, 0)
}

func (s *MeasurementContext) Unit() antlr.TerminalNode {
	return s.GetToken(ScssParserUnit, 0)
}

func (s *MeasurementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeasurementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MeasurementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMeasurement(s)
	}
}

func (s *MeasurementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMeasurement(s)
	}
}

func (p *ScssParser) Measurement() (localctx IMeasurementContext) {
	this := p
	_ = this

	localctx = NewMeasurementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, ScssParserRULE_measurement)
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
		p.SetState(643)
		p.Match(ScssParserNumber)
	}
	p.SetState(645)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserUnit {
		{
			p.SetState(644)
			p.Match(ScssParserUnit)
		}

	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) FunctionIdentifier() antlr.TerminalNode {
	return s.GetToken(ScssParserFunctionIdentifier, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *FunctionCallContext) Namespace() INamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceContext)
}

func (s *FunctionCallContext) PassedParams() IPassedParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassedParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassedParamsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *ScssParser) FunctionCall() (localctx IFunctionCallContext) {
	this := p
	_ = this

	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, ScssParserRULE_functionCall)
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
	p.SetState(648)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserIdentifier {
		{
			p.SetState(647)
			p.Namespace()
		}

	}
	{
		p.SetState(650)
		p.Match(ScssParserFunctionIdentifier)
	}
	p.SetState(652)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL_)|(1<<ScssParserInterpolationStart)|(1<<ScssParserLPAREN)|(1<<ScssParserLBRACK)|(1<<ScssParserDOLLAR)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR)|(1<<ScssParserMINUS_LPAREN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ScssParserPLUS_LPAREN-32))|(1<<(ScssParserUrlStart-32))|(1<<(ScssParserFROM-32))|(1<<(ScssParserTO-32))|(1<<(ScssParserTHROUGH-32))|(1<<(ScssParserONLY-32))|(1<<(ScssParserNOT-32))|(1<<(ScssParserAND_WORD-32))|(1<<(ScssParserUSING-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(ScssParserIdentifier-64))|(1<<(ScssParserFunctionIdentifier-64))|(1<<(ScssParserStringLiteral-64))|(1<<(ScssParserNumber-64))|(1<<(ScssParserColor-64)))) != 0) {
		{
			p.SetState(651)
			p.PassedParams()
		}

	}
	{
		p.SetState(654)
		p.Match(ScssParserRPAREN)
	}

	return localctx
}

// INamespaceContext is an interface to support dynamic dispatch.
type INamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceContext differentiates from other interfaces.
	IsNamespaceContext()
}

type NamespaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceContext() *NamespaceContext {
	var p = new(NamespaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_namespace
	return p
}

func (*NamespaceContext) IsNamespaceContext() {}

func NewNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceContext {
	var p = new(NamespaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_namespace

	return p
}

func (s *NamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(ScssParserIdentifier)
}

func (s *NamespaceContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserIdentifier, i)
}

func (s *NamespaceContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(ScssParserDOT)
}

func (s *NamespaceContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserDOT, i)
}

func (s *NamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterNamespace(s)
	}
}

func (s *NamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitNamespace(s)
	}
}

func (p *ScssParser) Namespace() (localctx INamespaceContext) {
	this := p
	_ = this

	localctx = NewNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, ScssParserRULE_namespace)
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
	p.SetState(658)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ScssParserIdentifier {
		{
			p.SetState(656)
			p.Match(ScssParserIdentifier)
		}
		{
			p.SetState(657)
			p.Match(ScssParserDOT)
		}

		p.SetState(660)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IList_Context is an interface to support dynamic dispatch.
type IList_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_Context differentiates from other interfaces.
	IsList_Context()
}

type List_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_Context() *List_Context {
	var p = new(List_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_list_
	return p
}

func (*List_Context) IsList_Context() {}

func NewList_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_Context {
	var p = new(List_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_list_

	return p
}

func (s *List_Context) GetParser() antlr.Parser { return s.parser }

func (s *List_Context) ListCommaSeparated() IListCommaSeparatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListCommaSeparatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListCommaSeparatedContext)
}

func (s *List_Context) ListSpaceSeparated() IListSpaceSeparatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListSpaceSeparatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListSpaceSeparatedContext)
}

func (s *List_Context) ListBracketed() IListBracketedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListBracketedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListBracketedContext)
}

func (s *List_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterList_(s)
	}
}

func (s *List_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitList_(s)
	}
}

func (p *ScssParser) List_() (localctx IList_Context) {
	this := p
	_ = this

	localctx = NewList_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, ScssParserRULE_list_)

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

	p.SetState(665)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 77, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(662)
			p.ListCommaSeparated()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(663)
			p.ListSpaceSeparated()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(664)
			p.ListBracketed()
		}

	}

	return localctx
}

// IListCommaSeparatedContext is an interface to support dynamic dispatch.
type IListCommaSeparatedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListCommaSeparatedContext differentiates from other interfaces.
	IsListCommaSeparatedContext()
}

type ListCommaSeparatedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListCommaSeparatedContext() *ListCommaSeparatedContext {
	var p = new(ListCommaSeparatedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_listCommaSeparated
	return p
}

func (*ListCommaSeparatedContext) IsListCommaSeparatedContext() {}

func NewListCommaSeparatedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListCommaSeparatedContext {
	var p = new(ListCommaSeparatedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_listCommaSeparated

	return p
}

func (s *ListCommaSeparatedContext) GetParser() antlr.Parser { return s.parser }

func (s *ListCommaSeparatedContext) AllListElement() []IListElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IListElementContext)(nil)).Elem())
	var tst = make([]IListElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IListElementContext)
		}
	}

	return tst
}

func (s *ListCommaSeparatedContext) ListElement(i int) IListElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IListElementContext)
}

func (s *ListCommaSeparatedContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *ListCommaSeparatedContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *ListCommaSeparatedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListCommaSeparatedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListCommaSeparatedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterListCommaSeparated(s)
	}
}

func (s *ListCommaSeparatedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitListCommaSeparated(s)
	}
}

func (p *ScssParser) ListCommaSeparated() (localctx IListCommaSeparatedContext) {
	this := p
	_ = this

	localctx = NewListCommaSeparatedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, ScssParserRULE_listCommaSeparated)

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
		p.SetState(667)
		p.ListElement()
	}
	p.SetState(670)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(668)
				p.Match(ScssParserCOMMA)
			}
			{
				p.SetState(669)
				p.ListElement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(672)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 78, p.GetParserRuleContext())
	}
	p.SetState(675)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 79, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(674)
			p.Match(ScssParserCOMMA)
		}

	}

	return localctx
}

// IListSpaceSeparatedContext is an interface to support dynamic dispatch.
type IListSpaceSeparatedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListSpaceSeparatedContext differentiates from other interfaces.
	IsListSpaceSeparatedContext()
}

type ListSpaceSeparatedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListSpaceSeparatedContext() *ListSpaceSeparatedContext {
	var p = new(ListSpaceSeparatedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_listSpaceSeparated
	return p
}

func (*ListSpaceSeparatedContext) IsListSpaceSeparatedContext() {}

func NewListSpaceSeparatedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListSpaceSeparatedContext {
	var p = new(ListSpaceSeparatedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_listSpaceSeparated

	return p
}

func (s *ListSpaceSeparatedContext) GetParser() antlr.Parser { return s.parser }

func (s *ListSpaceSeparatedContext) AllListElement() []IListElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IListElementContext)(nil)).Elem())
	var tst = make([]IListElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IListElementContext)
		}
	}

	return tst
}

func (s *ListSpaceSeparatedContext) ListElement(i int) IListElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IListElementContext)
}

func (s *ListSpaceSeparatedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListSpaceSeparatedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListSpaceSeparatedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterListSpaceSeparated(s)
	}
}

func (s *ListSpaceSeparatedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitListSpaceSeparated(s)
	}
}

func (p *ScssParser) ListSpaceSeparated() (localctx IListSpaceSeparatedContext) {
	this := p
	_ = this

	localctx = NewListSpaceSeparatedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, ScssParserRULE_listSpaceSeparated)
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
		p.SetState(677)
		p.ListElement()
	}
	p.SetState(679)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ScssParserNULL_)|(1<<ScssParserInterpolationStart)|(1<<ScssParserLPAREN)|(1<<ScssParserDOLLAR)|(1<<ScssParserMINUS_DOLLAR)|(1<<ScssParserPLUS_DOLLAR)|(1<<ScssParserMINUS_LPAREN))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(ScssParserPLUS_LPAREN-32))|(1<<(ScssParserUrlStart-32))|(1<<(ScssParserFROM-32))|(1<<(ScssParserTO-32))|(1<<(ScssParserTHROUGH-32))|(1<<(ScssParserONLY-32))|(1<<(ScssParserNOT-32))|(1<<(ScssParserAND_WORD-32))|(1<<(ScssParserUSING-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(ScssParserIdentifier-64))|(1<<(ScssParserFunctionIdentifier-64))|(1<<(ScssParserStringLiteral-64))|(1<<(ScssParserNumber-64))|(1<<(ScssParserColor-64)))) != 0) {
		{
			p.SetState(678)
			p.ListElement()
		}

		p.SetState(681)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IListBracketedContext is an interface to support dynamic dispatch.
type IListBracketedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListBracketedContext differentiates from other interfaces.
	IsListBracketedContext()
}

type ListBracketedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListBracketedContext() *ListBracketedContext {
	var p = new(ListBracketedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_listBracketed
	return p
}

func (*ListBracketedContext) IsListBracketedContext() {}

func NewListBracketedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListBracketedContext {
	var p = new(ListBracketedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_listBracketed

	return p
}

func (s *ListBracketedContext) GetParser() antlr.Parser { return s.parser }

func (s *ListBracketedContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ScssParserLBRACK, 0)
}

func (s *ListBracketedContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ScssParserRBRACK, 0)
}

func (s *ListBracketedContext) ListCommaSeparated() IListCommaSeparatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListCommaSeparatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListCommaSeparatedContext)
}

func (s *ListBracketedContext) ListSpaceSeparated() IListSpaceSeparatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListSpaceSeparatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListSpaceSeparatedContext)
}

func (s *ListBracketedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListBracketedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListBracketedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterListBracketed(s)
	}
}

func (s *ListBracketedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitListBracketed(s)
	}
}

func (p *ScssParser) ListBracketed() (localctx IListBracketedContext) {
	this := p
	_ = this

	localctx = NewListBracketedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, ScssParserRULE_listBracketed)

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
		p.SetState(683)
		p.Match(ScssParserLBRACK)
	}
	p.SetState(686)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 81, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(684)
			p.ListCommaSeparated()
		}

	case 2:
		{
			p.SetState(685)
			p.ListSpaceSeparated()
		}

	}
	{
		p.SetState(688)
		p.Match(ScssParserRBRACK)
	}

	return localctx
}

// IListElementContext is an interface to support dynamic dispatch.
type IListElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListElementContext differentiates from other interfaces.
	IsListElementContext()
}

type ListElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListElementContext() *ListElementContext {
	var p = new(ListElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_listElement
	return p
}

func (*ListElementContext) IsListElementContext() {}

func NewListElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListElementContext {
	var p = new(ListElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_listElement

	return p
}

func (s *ListElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ListElementContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *ListElementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *ListElementContext) List_() IList_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_Context)
}

func (s *ListElementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *ListElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterListElement(s)
	}
}

func (s *ListElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitListElement(s)
	}
}

func (p *ScssParser) ListElement() (localctx IListElementContext) {
	this := p
	_ = this

	localctx = NewListElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, ScssParserRULE_listElement)

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

	p.SetState(695)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 82, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(690)
			p.CommandStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(691)
			p.Match(ScssParserLPAREN)
		}
		{
			p.SetState(692)
			p.List_()
		}
		{
			p.SetState(693)
			p.Match(ScssParserRPAREN)
		}

	}

	return localctx
}

// IMap_Context is an interface to support dynamic dispatch.
type IMap_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMap_Context differentiates from other interfaces.
	IsMap_Context()
}

type Map_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMap_Context() *Map_Context {
	var p = new(Map_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_map_
	return p
}

func (*Map_Context) IsMap_Context() {}

func NewMap_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_Context {
	var p = new(Map_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_map_

	return p
}

func (s *Map_Context) GetParser() antlr.Parser { return s.parser }

func (s *Map_Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserLPAREN, 0)
}

func (s *Map_Context) AllMapEntry() []IMapEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMapEntryContext)(nil)).Elem())
	var tst = make([]IMapEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMapEntryContext)
		}
	}

	return tst
}

func (s *Map_Context) MapEntry(i int) IMapEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMapEntryContext)
}

func (s *Map_Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(ScssParserRPAREN, 0)
}

func (s *Map_Context) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ScssParserCOMMA)
}

func (s *Map_Context) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ScssParserCOMMA, i)
}

func (s *Map_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Map_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMap_(s)
	}
}

func (s *Map_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMap_(s)
	}
}

func (p *ScssParser) Map_() (localctx IMap_Context) {
	this := p
	_ = this

	localctx = NewMap_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, ScssParserRULE_map_)
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
		p.SetState(697)
		p.Match(ScssParserLPAREN)
	}
	{
		p.SetState(698)
		p.MapEntry()
	}
	p.SetState(703)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(699)
				p.Match(ScssParserCOMMA)
			}
			{
				p.SetState(700)
				p.MapEntry()
			}

		}
		p.SetState(705)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext())
	}
	p.SetState(707)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ScssParserCOMMA {
		{
			p.SetState(706)
			p.Match(ScssParserCOMMA)
		}

	}
	{
		p.SetState(709)
		p.Match(ScssParserRPAREN)
	}

	return localctx
}

// IMapEntryContext is an interface to support dynamic dispatch.
type IMapEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapEntryContext differentiates from other interfaces.
	IsMapEntryContext()
}

type MapEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapEntryContext() *MapEntryContext {
	var p = new(MapEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mapEntry
	return p
}

func (*MapEntryContext) IsMapEntryContext() {}

func NewMapEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapEntryContext {
	var p = new(MapEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mapEntry

	return p
}

func (s *MapEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *MapEntryContext) MapKey() IMapKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapKeyContext)
}

func (s *MapEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(ScssParserCOLON, 0)
}

func (s *MapEntryContext) MapValue() IMapValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapValueContext)
}

func (s *MapEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMapEntry(s)
	}
}

func (s *MapEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMapEntry(s)
	}
}

func (p *ScssParser) MapEntry() (localctx IMapEntryContext) {
	this := p
	_ = this

	localctx = NewMapEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, ScssParserRULE_mapEntry)

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
		p.SetState(711)
		p.MapKey()
	}
	{
		p.SetState(712)
		p.Match(ScssParserCOLON)
	}
	{
		p.SetState(713)
		p.MapValue()
	}

	return localctx
}

// IMapKeyContext is an interface to support dynamic dispatch.
type IMapKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapKeyContext differentiates from other interfaces.
	IsMapKeyContext()
}

type MapKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapKeyContext() *MapKeyContext {
	var p = new(MapKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mapKey
	return p
}

func (*MapKeyContext) IsMapKeyContext() {}

func NewMapKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapKeyContext {
	var p = new(MapKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mapKey

	return p
}

func (s *MapKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *MapKeyContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *MapKeyContext) List_() IList_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_Context)
}

func (s *MapKeyContext) Map_() IMap_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMap_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMap_Context)
}

func (s *MapKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMapKey(s)
	}
}

func (s *MapKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMapKey(s)
	}
}

func (p *ScssParser) MapKey() (localctx IMapKeyContext) {
	this := p
	_ = this

	localctx = NewMapKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, ScssParserRULE_mapKey)

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

	p.SetState(718)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 85, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(715)
			p.CommandStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(716)
			p.List_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(717)
			p.Map_()
		}

	}

	return localctx
}

// IMapValueContext is an interface to support dynamic dispatch.
type IMapValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapValueContext differentiates from other interfaces.
	IsMapValueContext()
}

type MapValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapValueContext() *MapValueContext {
	var p = new(MapValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScssParserRULE_mapValue
	return p
}

func (*MapValueContext) IsMapValueContext() {}

func NewMapValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapValueContext {
	var p = new(MapValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScssParserRULE_mapValue

	return p
}

func (s *MapValueContext) GetParser() antlr.Parser { return s.parser }

func (s *MapValueContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *MapValueContext) List_() IList_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_Context)
}

func (s *MapValueContext) Map_() IMap_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMap_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMap_Context)
}

func (s *MapValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.EnterMapValue(s)
	}
}

func (s *MapValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScssParserListener); ok {
		listenerT.ExitMapValue(s)
	}
}

func (p *ScssParser) MapValue() (localctx IMapValueContext) {
	this := p
	_ = this

	localctx = NewMapValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, ScssParserRULE_mapValue)

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

	p.SetState(723)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 86, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(720)
			p.CommandStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(721)
			p.List_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(722)
			p.Map_()
		}

	}

	return localctx
}
