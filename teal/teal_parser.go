// Code generated from Teal.g4 by ANTLR 4.9.3. DO NOT EDIT.

package teal // Teal
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 78, 653,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 3,
	2, 3, 2, 3, 2, 3, 3, 7, 3, 103, 10, 3, 12, 3, 14, 3, 106, 11, 3, 3, 3,
	5, 3, 109, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 7, 4, 145, 10, 4, 12, 4, 14, 4, 148, 11, 4, 3, 4, 3, 4, 5, 4, 152,
	10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	164, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 190, 10, 4, 3, 4, 3, 4, 5, 4, 194, 10, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 5, 4, 213, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 5, 4, 222, 10, 4, 3, 5, 3, 5, 5, 5, 226, 10, 5, 3, 5, 3, 5, 3, 5,
	5, 5, 231, 10, 5, 7, 5, 233, 10, 5, 12, 5, 14, 5, 236, 11, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 249, 10,
	7, 12, 7, 14, 7, 252, 11, 7, 5, 7, 254, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 5, 8, 274, 10, 8, 5, 8, 276, 10, 8, 3, 9, 3, 9, 3, 9, 5, 9,
	281, 10, 9, 3, 10, 3, 10, 5, 10, 285, 10, 10, 3, 10, 5, 10, 288, 10, 10,
	3, 10, 3, 10, 3, 10, 5, 10, 293, 10, 10, 5, 10, 295, 10, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 7, 11, 301, 10, 11, 12, 11, 14, 11, 304, 11, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 5, 12, 310, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5,
	12, 316, 10, 12, 3, 12, 3, 12, 3, 12, 7, 12, 321, 10, 12, 12, 12, 14, 12,
	324, 11, 12, 3, 12, 3, 12, 3, 12, 7, 12, 329, 10, 12, 12, 12, 14, 12, 332,
	11, 12, 3, 12, 3, 12, 3, 12, 7, 12, 337, 10, 12, 12, 12, 14, 12, 340, 11,
	12, 3, 12, 3, 12, 3, 12, 5, 12, 345, 10, 12, 3, 13, 5, 13, 348, 10, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 355, 10, 13, 3, 14, 3, 14, 3,
	14, 7, 14, 360, 10, 14, 12, 14, 14, 14, 363, 11, 14, 3, 15, 3, 15, 5, 15,
	367, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 7, 16, 374, 10, 16, 12,
	16, 14, 16, 377, 11, 16, 3, 17, 3, 17, 3, 17, 5, 17, 382, 10, 17, 3, 18,
	3, 18, 5, 18, 386, 10, 18, 3, 18, 5, 18, 389, 10, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 7, 20, 398, 10, 20, 12, 20, 14, 20, 401,
	11, 20, 3, 20, 3, 20, 5, 20, 405, 10, 20, 3, 21, 3, 21, 3, 21, 7, 21, 410,
	10, 21, 12, 21, 14, 21, 413, 11, 21, 3, 22, 3, 22, 3, 22, 7, 22, 418, 10,
	22, 12, 22, 14, 22, 421, 11, 22, 3, 23, 3, 23, 3, 23, 7, 23, 426, 10, 23,
	12, 23, 14, 23, 429, 11, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5,
	24, 447, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 484, 10, 24, 12, 24, 14,
	24, 487, 11, 24, 3, 25, 3, 25, 7, 25, 491, 10, 25, 12, 25, 14, 25, 494,
	11, 25, 3, 26, 3, 26, 6, 26, 498, 10, 26, 13, 26, 14, 26, 499, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 5, 27, 507, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 5, 28, 515, 10, 28, 3, 28, 7, 28, 518, 10, 28, 12, 28, 14,
	28, 521, 11, 28, 3, 29, 7, 29, 524, 10, 29, 12, 29, 14, 29, 527, 11, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 535, 10, 29, 3, 30, 3,
	30, 5, 30, 539, 10, 30, 3, 30, 3, 30, 3, 31, 3, 31, 5, 31, 545, 10, 31,
	3, 31, 3, 31, 3, 31, 5, 31, 550, 10, 31, 3, 32, 3, 32, 3, 32, 3, 33, 5,
	33, 556, 10, 33, 3, 33, 3, 33, 5, 33, 560, 10, 33, 3, 33, 3, 33, 3, 33,
	5, 33, 565, 10, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 5, 34, 573,
	10, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 581, 10, 34, 5,
	34, 583, 10, 34, 3, 34, 3, 34, 3, 34, 5, 34, 588, 10, 34, 5, 34, 590, 10,
	34, 3, 35, 3, 35, 5, 35, 594, 10, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36,
	3, 36, 7, 36, 602, 10, 36, 12, 36, 14, 36, 605, 11, 36, 3, 36, 5, 36, 608,
	10, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	5, 37, 619, 10, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 627,
	10, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42,
	3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3,
	47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 2, 3, 46, 50, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84,
	86, 88, 90, 92, 94, 96, 2, 10, 4, 2, 3, 3, 17, 17, 4, 2, 23, 24, 49, 52,
	3, 2, 54, 55, 3, 2, 56, 59, 4, 2, 27, 27, 60, 63, 5, 2, 55, 55, 61, 61,
	64, 65, 3, 2, 71, 74, 3, 2, 68, 70, 2, 717, 2, 98, 3, 2, 2, 2, 4, 104,
	3, 2, 2, 2, 6, 221, 3, 2, 2, 2, 8, 223, 3, 2, 2, 2, 10, 237, 3, 2, 2, 2,
	12, 253, 3, 2, 2, 2, 14, 275, 3, 2, 2, 2, 16, 277, 3, 2, 2, 2, 18, 294,
	3, 2, 2, 2, 20, 296, 3, 2, 2, 2, 22, 344, 3, 2, 2, 2, 24, 347, 3, 2, 2,
	2, 26, 356, 3, 2, 2, 2, 28, 366, 3, 2, 2, 2, 30, 370, 3, 2, 2, 2, 32, 378,
	3, 2, 2, 2, 34, 383, 3, 2, 2, 2, 36, 390, 3, 2, 2, 2, 38, 394, 3, 2, 2,
	2, 40, 406, 3, 2, 2, 2, 42, 414, 3, 2, 2, 2, 44, 422, 3, 2, 2, 2, 46, 446,
	3, 2, 2, 2, 48, 488, 3, 2, 2, 2, 50, 495, 3, 2, 2, 2, 52, 506, 3, 2, 2,
	2, 54, 514, 3, 2, 2, 2, 56, 525, 3, 2, 2, 2, 58, 538, 3, 2, 2, 2, 60, 549,
	3, 2, 2, 2, 62, 551, 3, 2, 2, 2, 64, 555, 3, 2, 2, 2, 66, 589, 3, 2, 2,
	2, 68, 591, 3, 2, 2, 2, 70, 597, 3, 2, 2, 2, 72, 626, 3, 2, 2, 2, 74, 628,
	3, 2, 2, 2, 76, 630, 3, 2, 2, 2, 78, 632, 3, 2, 2, 2, 80, 634, 3, 2, 2,
	2, 82, 636, 3, 2, 2, 2, 84, 638, 3, 2, 2, 2, 86, 640, 3, 2, 2, 2, 88, 642,
	3, 2, 2, 2, 90, 644, 3, 2, 2, 2, 92, 646, 3, 2, 2, 2, 94, 648, 3, 2, 2,
	2, 96, 650, 3, 2, 2, 2, 98, 99, 5, 4, 3, 2, 99, 100, 7, 2, 2, 3, 100, 3,
	3, 2, 2, 2, 101, 103, 5, 6, 4, 2, 102, 101, 3, 2, 2, 2, 103, 106, 3, 2,
	2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2,
	106, 104, 3, 2, 2, 2, 107, 109, 5, 34, 18, 2, 108, 107, 3, 2, 2, 2, 108,
	109, 3, 2, 2, 2, 109, 5, 3, 2, 2, 2, 110, 222, 7, 3, 2, 2, 111, 112, 5,
	40, 21, 2, 112, 113, 7, 4, 2, 2, 113, 114, 5, 44, 23, 2, 114, 222, 3, 2,
	2, 2, 115, 222, 5, 50, 26, 2, 116, 222, 5, 36, 19, 2, 117, 222, 7, 5, 2,
	2, 118, 119, 7, 6, 2, 2, 119, 222, 7, 67, 2, 2, 120, 121, 7, 7, 2, 2, 121,
	122, 5, 4, 3, 2, 122, 123, 7, 8, 2, 2, 123, 222, 3, 2, 2, 2, 124, 125,
	7, 9, 2, 2, 125, 126, 5, 46, 24, 2, 126, 127, 7, 7, 2, 2, 127, 128, 5,
	4, 3, 2, 128, 129, 7, 8, 2, 2, 129, 222, 3, 2, 2, 2, 130, 131, 7, 10, 2,
	2, 131, 132, 5, 4, 3, 2, 132, 133, 7, 11, 2, 2, 133, 134, 5, 46, 24, 2,
	134, 222, 3, 2, 2, 2, 135, 136, 7, 12, 2, 2, 136, 137, 5, 46, 24, 2, 137,
	138, 7, 13, 2, 2, 138, 146, 5, 4, 3, 2, 139, 140, 7, 14, 2, 2, 140, 141,
	5, 46, 24, 2, 141, 142, 7, 13, 2, 2, 142, 143, 5, 4, 3, 2, 143, 145, 3,
	2, 2, 2, 144, 139, 3, 2, 2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2,
	2, 146, 147, 3, 2, 2, 2, 147, 151, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149,
	150, 7, 15, 2, 2, 150, 152, 5, 4, 3, 2, 151, 149, 3, 2, 2, 2, 151, 152,
	3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 7, 8, 2, 2, 154, 222, 3, 2,
	2, 2, 155, 156, 7, 16, 2, 2, 156, 157, 7, 67, 2, 2, 157, 158, 7, 4, 2,
	2, 158, 159, 5, 46, 24, 2, 159, 160, 7, 17, 2, 2, 160, 163, 5, 46, 24,
	2, 161, 162, 7, 17, 2, 2, 162, 164, 5, 46, 24, 2, 163, 161, 3, 2, 2, 2,
	163, 164, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 7, 7, 2, 2, 166,
	167, 5, 4, 3, 2, 167, 168, 7, 8, 2, 2, 168, 222, 3, 2, 2, 2, 169, 170,
	7, 16, 2, 2, 170, 171, 5, 42, 22, 2, 171, 172, 7, 18, 2, 2, 172, 173, 5,
	44, 23, 2, 173, 174, 7, 7, 2, 2, 174, 175, 5, 4, 3, 2, 175, 176, 7, 8,
	2, 2, 176, 222, 3, 2, 2, 2, 177, 178, 7, 19, 2, 2, 178, 179, 5, 38, 20,
	2, 179, 180, 5, 64, 33, 2, 180, 222, 3, 2, 2, 2, 181, 182, 7, 20, 2, 2,
	182, 183, 7, 19, 2, 2, 183, 184, 7, 67, 2, 2, 184, 222, 5, 64, 33, 2, 185,
	186, 7, 20, 2, 2, 186, 189, 5, 8, 5, 2, 187, 188, 7, 21, 2, 2, 188, 190,
	5, 16, 9, 2, 189, 187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 193, 3, 2,
	2, 2, 191, 192, 7, 4, 2, 2, 192, 194, 5, 44, 23, 2, 193, 191, 3, 2, 2,
	2, 193, 194, 3, 2, 2, 2, 194, 222, 3, 2, 2, 2, 195, 196, 7, 20, 2, 2, 196,
	197, 7, 67, 2, 2, 197, 198, 7, 4, 2, 2, 198, 222, 5, 22, 12, 2, 199, 200,
	7, 22, 2, 2, 200, 201, 7, 19, 2, 2, 201, 202, 7, 67, 2, 2, 202, 222, 5,
	64, 33, 2, 203, 204, 7, 22, 2, 2, 204, 205, 5, 8, 5, 2, 205, 206, 7, 21,
	2, 2, 206, 207, 5, 16, 9, 2, 207, 222, 3, 2, 2, 2, 208, 209, 7, 22, 2,
	2, 209, 212, 5, 8, 5, 2, 210, 211, 7, 21, 2, 2, 211, 213, 5, 16, 9, 2,
	212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214,
	215, 7, 4, 2, 2, 215, 216, 5, 44, 23, 2, 216, 222, 3, 2, 2, 2, 217, 218,
	7, 22, 2, 2, 218, 219, 7, 67, 2, 2, 219, 220, 7, 4, 2, 2, 220, 222, 5,
	22, 12, 2, 221, 110, 3, 2, 2, 2, 221, 111, 3, 2, 2, 2, 221, 115, 3, 2,
	2, 2, 221, 116, 3, 2, 2, 2, 221, 117, 3, 2, 2, 2, 221, 118, 3, 2, 2, 2,
	221, 120, 3, 2, 2, 2, 221, 124, 3, 2, 2, 2, 221, 130, 3, 2, 2, 2, 221,
	135, 3, 2, 2, 2, 221, 155, 3, 2, 2, 2, 221, 169, 3, 2, 2, 2, 221, 177,
	3, 2, 2, 2, 221, 181, 3, 2, 2, 2, 221, 185, 3, 2, 2, 2, 221, 195, 3, 2,
	2, 2, 221, 199, 3, 2, 2, 2, 221, 203, 3, 2, 2, 2, 221, 208, 3, 2, 2, 2,
	221, 217, 3, 2, 2, 2, 222, 7, 3, 2, 2, 2, 223, 225, 7, 67, 2, 2, 224, 226,
	5, 10, 6, 2, 225, 224, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 234, 3, 2,
	2, 2, 227, 228, 7, 17, 2, 2, 228, 230, 7, 67, 2, 2, 229, 231, 5, 10, 6,
	2, 230, 229, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 233, 3, 2, 2, 2, 232,
	227, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235,
	3, 2, 2, 2, 235, 9, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 237, 238, 7, 23,
	2, 2, 238, 239, 7, 67, 2, 2, 239, 240, 7, 24, 2, 2, 240, 11, 3, 2, 2, 2,
	241, 242, 7, 25, 2, 2, 242, 243, 5, 12, 7, 2, 243, 244, 7, 26, 2, 2, 244,
	254, 3, 2, 2, 2, 245, 250, 5, 14, 8, 2, 246, 247, 7, 27, 2, 2, 247, 249,
	5, 14, 8, 2, 248, 246, 3, 2, 2, 2, 249, 252, 3, 2, 2, 2, 250, 248, 3, 2,
	2, 2, 250, 251, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2,
	253, 241, 3, 2, 2, 2, 253, 245, 3, 2, 2, 2, 254, 13, 3, 2, 2, 2, 255, 276,
	7, 28, 2, 2, 256, 276, 7, 29, 2, 2, 257, 276, 7, 30, 2, 2, 258, 276, 7,
	31, 2, 2, 259, 260, 7, 32, 2, 2, 260, 261, 5, 12, 7, 2, 261, 262, 7, 33,
	2, 2, 262, 276, 3, 2, 2, 2, 263, 264, 7, 32, 2, 2, 264, 265, 5, 12, 7,
	2, 265, 266, 7, 21, 2, 2, 266, 267, 5, 12, 7, 2, 267, 268, 7, 33, 2, 2,
	268, 276, 3, 2, 2, 2, 269, 270, 7, 19, 2, 2, 270, 276, 5, 24, 13, 2, 271,
	273, 7, 67, 2, 2, 272, 274, 5, 20, 11, 2, 273, 272, 3, 2, 2, 2, 273, 274,
	3, 2, 2, 2, 274, 276, 3, 2, 2, 2, 275, 255, 3, 2, 2, 2, 275, 256, 3, 2,
	2, 2, 275, 257, 3, 2, 2, 2, 275, 258, 3, 2, 2, 2, 275, 259, 3, 2, 2, 2,
	275, 263, 3, 2, 2, 2, 275, 269, 3, 2, 2, 2, 275, 271, 3, 2, 2, 2, 276,
	15, 3, 2, 2, 2, 277, 280, 5, 12, 7, 2, 278, 279, 7, 17, 2, 2, 279, 281,
	5, 12, 7, 2, 280, 278, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 17, 3, 2,
	2, 2, 282, 284, 7, 25, 2, 2, 283, 285, 5, 16, 9, 2, 284, 283, 3, 2, 2,
	2, 284, 285, 3, 2, 2, 2, 285, 287, 3, 2, 2, 2, 286, 288, 7, 34, 2, 2, 287,
	286, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 295,
	7, 26, 2, 2, 290, 292, 5, 16, 9, 2, 291, 293, 7, 34, 2, 2, 292, 291, 3,
	2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 295, 3, 2, 2, 2, 294, 282, 3, 2, 2,
	2, 294, 290, 3, 2, 2, 2, 295, 19, 3, 2, 2, 2, 296, 297, 7, 23, 2, 2, 297,
	302, 7, 67, 2, 2, 298, 299, 7, 17, 2, 2, 299, 301, 7, 67, 2, 2, 300, 298,
	3, 2, 2, 2, 301, 304, 3, 2, 2, 2, 302, 300, 3, 2, 2, 2, 302, 303, 3, 2,
	2, 2, 303, 305, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 305, 306, 7, 24, 2, 2,
	306, 21, 3, 2, 2, 2, 307, 309, 7, 35, 2, 2, 308, 310, 5, 20, 11, 2, 309,
	308, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 315, 3, 2, 2, 2, 311, 312,
	7, 32, 2, 2, 312, 313, 5, 12, 7, 2, 313, 314, 7, 33, 2, 2, 314, 316, 3,
	2, 2, 2, 315, 311, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 322, 3, 2, 2,
	2, 317, 318, 7, 67, 2, 2, 318, 319, 7, 4, 2, 2, 319, 321, 5, 22, 12, 2,
	320, 317, 3, 2, 2, 2, 321, 324, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322,
	323, 3, 2, 2, 2, 323, 330, 3, 2, 2, 2, 324, 322, 3, 2, 2, 2, 325, 326,
	7, 67, 2, 2, 326, 327, 7, 21, 2, 2, 327, 329, 5, 12, 7, 2, 328, 325, 3,
	2, 2, 2, 329, 332, 3, 2, 2, 2, 330, 328, 3, 2, 2, 2, 330, 331, 3, 2, 2,
	2, 331, 333, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 333, 345, 7, 8, 2, 2, 334,
	338, 7, 36, 2, 2, 335, 337, 5, 96, 49, 2, 336, 335, 3, 2, 2, 2, 337, 340,
	3, 2, 2, 2, 338, 336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 341, 3, 2,
	2, 2, 340, 338, 3, 2, 2, 2, 341, 345, 7, 8, 2, 2, 342, 343, 7, 37, 2, 2,
	343, 345, 5, 24, 13, 2, 344, 307, 3, 2, 2, 2, 344, 334, 3, 2, 2, 2, 344,
	342, 3, 2, 2, 2, 345, 23, 3, 2, 2, 2, 346, 348, 5, 20, 11, 2, 347, 346,
	3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 350, 7, 25,
	2, 2, 350, 351, 5, 26, 14, 2, 351, 354, 7, 26, 2, 2, 352, 353, 7, 21, 2,
	2, 353, 355, 5, 18, 10, 2, 354, 352, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2,
	355, 25, 3, 2, 2, 2, 356, 361, 5, 28, 15, 2, 357, 358, 7, 17, 2, 2, 358,
	360, 5, 28, 15, 2, 359, 357, 3, 2, 2, 2, 360, 363, 3, 2, 2, 2, 361, 359,
	3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 27, 3, 2, 2, 2, 363, 361, 3, 2,
	2, 2, 364, 365, 7, 67, 2, 2, 365, 367, 7, 21, 2, 2, 366, 364, 3, 2, 2,
	2, 366, 367, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 369, 5, 12, 7, 2, 369,
	29, 3, 2, 2, 2, 370, 375, 5, 32, 17, 2, 371, 372, 7, 17, 2, 2, 372, 374,
	5, 32, 17, 2, 373, 371, 3, 2, 2, 2, 374, 377, 3, 2, 2, 2, 375, 373, 3,
	2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 31, 3, 2, 2, 2, 377, 375, 3, 2, 2,
	2, 378, 381, 7, 67, 2, 2, 379, 380, 7, 21, 2, 2, 380, 382, 5, 12, 7, 2,
	381, 379, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 33, 3, 2, 2, 2, 383, 385,
	7, 38, 2, 2, 384, 386, 5, 44, 23, 2, 385, 384, 3, 2, 2, 2, 385, 386, 3,
	2, 2, 2, 386, 388, 3, 2, 2, 2, 387, 389, 7, 3, 2, 2, 388, 387, 3, 2, 2,
	2, 388, 389, 3, 2, 2, 2, 389, 35, 3, 2, 2, 2, 390, 391, 7, 39, 2, 2, 391,
	392, 7, 67, 2, 2, 392, 393, 7, 39, 2, 2, 393, 37, 3, 2, 2, 2, 394, 399,
	7, 67, 2, 2, 395, 396, 7, 40, 2, 2, 396, 398, 7, 67, 2, 2, 397, 395, 3,
	2, 2, 2, 398, 401, 3, 2, 2, 2, 399, 397, 3, 2, 2, 2, 399, 400, 3, 2, 2,
	2, 400, 404, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2, 402, 403, 7, 21, 2, 2, 403,
	405, 7, 67, 2, 2, 404, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 39,
	3, 2, 2, 2, 406, 411, 5, 54, 28, 2, 407, 408, 7, 17, 2, 2, 408, 410, 5,
	54, 28, 2, 409, 407, 3, 2, 2, 2, 410, 413, 3, 2, 2, 2, 411, 409, 3, 2,
	2, 2, 411, 412, 3, 2, 2, 2, 412, 41, 3, 2, 2, 2, 413, 411, 3, 2, 2, 2,
	414, 419, 7, 67, 2, 2, 415, 416, 7, 17, 2, 2, 416, 418, 7, 67, 2, 2, 417,
	415, 3, 2, 2, 2, 418, 421, 3, 2, 2, 2, 419, 417, 3, 2, 2, 2, 419, 420,
	3, 2, 2, 2, 420, 43, 3, 2, 2, 2, 421, 419, 3, 2, 2, 2, 422, 427, 5, 46,
	24, 2, 423, 424, 7, 17, 2, 2, 424, 426, 5, 46, 24, 2, 425, 423, 3, 2, 2,
	2, 426, 429, 3, 2, 2, 2, 427, 425, 3, 2, 2, 2, 427, 428, 3, 2, 2, 2, 428,
	45, 3, 2, 2, 2, 429, 427, 3, 2, 2, 2, 430, 431, 8, 24, 1, 2, 431, 447,
	7, 30, 2, 2, 432, 447, 7, 41, 2, 2, 433, 447, 7, 42, 2, 2, 434, 447, 5,
	94, 48, 2, 435, 447, 5, 96, 49, 2, 436, 447, 7, 34, 2, 2, 437, 447, 5,
	62, 32, 2, 438, 447, 5, 48, 25, 2, 439, 447, 5, 68, 35, 2, 440, 441, 5,
	90, 46, 2, 441, 442, 5, 46, 24, 11, 442, 447, 3, 2, 2, 2, 443, 444, 7,
	67, 2, 2, 444, 445, 7, 44, 2, 2, 445, 447, 5, 12, 7, 2, 446, 430, 3, 2,
	2, 2, 446, 432, 3, 2, 2, 2, 446, 433, 3, 2, 2, 2, 446, 434, 3, 2, 2, 2,
	446, 435, 3, 2, 2, 2, 446, 436, 3, 2, 2, 2, 446, 437, 3, 2, 2, 2, 446,
	438, 3, 2, 2, 2, 446, 439, 3, 2, 2, 2, 446, 440, 3, 2, 2, 2, 446, 443,
	3, 2, 2, 2, 447, 485, 3, 2, 2, 2, 448, 449, 12, 12, 2, 2, 449, 450, 5,
	92, 47, 2, 450, 451, 5, 46, 24, 12, 451, 484, 3, 2, 2, 2, 452, 453, 12,
	10, 2, 2, 453, 454, 5, 86, 44, 2, 454, 455, 5, 46, 24, 11, 455, 484, 3,
	2, 2, 2, 456, 457, 12, 9, 2, 2, 457, 458, 5, 84, 43, 2, 458, 459, 5, 46,
	24, 10, 459, 484, 3, 2, 2, 2, 460, 461, 12, 8, 2, 2, 461, 462, 5, 82, 42,
	2, 462, 463, 5, 46, 24, 8, 463, 484, 3, 2, 2, 2, 464, 465, 12, 7, 2, 2,
	465, 466, 5, 80, 41, 2, 466, 467, 5, 46, 24, 8, 467, 484, 3, 2, 2, 2, 468,
	469, 12, 5, 2, 2, 469, 470, 5, 78, 40, 2, 470, 471, 5, 46, 24, 6, 471,
	484, 3, 2, 2, 2, 472, 473, 12, 4, 2, 2, 473, 474, 5, 76, 39, 2, 474, 475,
	5, 46, 24, 5, 475, 484, 3, 2, 2, 2, 476, 477, 12, 3, 2, 2, 477, 478, 5,
	88, 45, 2, 478, 479, 5, 46, 24, 4, 479, 484, 3, 2, 2, 2, 480, 481, 12,
	13, 2, 2, 481, 482, 7, 43, 2, 2, 482, 484, 5, 12, 7, 2, 483, 448, 3, 2,
	2, 2, 483, 452, 3, 2, 2, 2, 483, 456, 3, 2, 2, 2, 483, 460, 3, 2, 2, 2,
	483, 464, 3, 2, 2, 2, 483, 468, 3, 2, 2, 2, 483, 472, 3, 2, 2, 2, 483,
	476, 3, 2, 2, 2, 483, 480, 3, 2, 2, 2, 484, 487, 3, 2, 2, 2, 485, 483,
	3, 2, 2, 2, 485, 486, 3, 2, 2, 2, 486, 47, 3, 2, 2, 2, 487, 485, 3, 2,
	2, 2, 488, 492, 5, 52, 27, 2, 489, 491, 5, 58, 30, 2, 490, 489, 3, 2, 2,
	2, 491, 494, 3, 2, 2, 2, 492, 490, 3, 2, 2, 2, 492, 493, 3, 2, 2, 2, 493,
	49, 3, 2, 2, 2, 494, 492, 3, 2, 2, 2, 495, 497, 5, 52, 27, 2, 496, 498,
	5, 58, 30, 2, 497, 496, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 499, 497, 3,
	2, 2, 2, 499, 500, 3, 2, 2, 2, 500, 51, 3, 2, 2, 2, 501, 507, 5, 54, 28,
	2, 502, 503, 7, 25, 2, 2, 503, 504, 5, 46, 24, 2, 504, 505, 7, 26, 2, 2,
	505, 507, 3, 2, 2, 2, 506, 501, 3, 2, 2, 2, 506, 502, 3, 2, 2, 2, 507,
	53, 3, 2, 2, 2, 508, 515, 7, 67, 2, 2, 509, 510, 7, 25, 2, 2, 510, 511,
	5, 46, 24, 2, 511, 512, 7, 26, 2, 2, 512, 513, 5, 56, 29, 2, 513, 515,
	3, 2, 2, 2, 514, 508, 3, 2, 2, 2, 514, 509, 3, 2, 2, 2, 515, 519, 3, 2,
	2, 2, 516, 518, 5, 56, 29, 2, 517, 516, 3, 2, 2, 2, 518, 521, 3, 2, 2,
	2, 519, 517, 3, 2, 2, 2, 519, 520, 3, 2, 2, 2, 520, 55, 3, 2, 2, 2, 521,
	519, 3, 2, 2, 2, 522, 524, 5, 58, 30, 2, 523, 522, 3, 2, 2, 2, 524, 527,
	3, 2, 2, 2, 525, 523, 3, 2, 2, 2, 525, 526, 3, 2, 2, 2, 526, 534, 3, 2,
	2, 2, 527, 525, 3, 2, 2, 2, 528, 529, 7, 45, 2, 2, 529, 530, 5, 46, 24,
	2, 530, 531, 7, 46, 2, 2, 531, 535, 3, 2, 2, 2, 532, 533, 7, 40, 2, 2,
	533, 535, 7, 67, 2, 2, 534, 528, 3, 2, 2, 2, 534, 532, 3, 2, 2, 2, 535,
	57, 3, 2, 2, 2, 536, 537, 7, 21, 2, 2, 537, 539, 7, 67, 2, 2, 538, 536,
	3, 2, 2, 2, 538, 539, 3, 2, 2, 2, 539, 540, 3, 2, 2, 2, 540, 541, 5, 60,
	31, 2, 541, 59, 3, 2, 2, 2, 542, 544, 7, 25, 2, 2, 543, 545, 5, 44, 23,
	2, 544, 543, 3, 2, 2, 2, 544, 545, 3, 2, 2, 2, 545, 546, 3, 2, 2, 2, 546,
	550, 7, 26, 2, 2, 547, 550, 5, 68, 35, 2, 548, 550, 5, 96, 49, 2, 549,
	542, 3, 2, 2, 2, 549, 547, 3, 2, 2, 2, 549, 548, 3, 2, 2, 2, 550, 61, 3,
	2, 2, 2, 551, 552, 7, 19, 2, 2, 552, 553, 5, 64, 33, 2, 553, 63, 3, 2,
	2, 2, 554, 556, 5, 20, 11, 2, 555, 554, 3, 2, 2, 2, 555, 556, 3, 2, 2,
	2, 556, 557, 3, 2, 2, 2, 557, 559, 7, 25, 2, 2, 558, 560, 5, 66, 34, 2,
	559, 558, 3, 2, 2, 2, 559, 560, 3, 2, 2, 2, 560, 561, 3, 2, 2, 2, 561,
	564, 7, 26, 2, 2, 562, 563, 7, 21, 2, 2, 563, 565, 5, 18, 10, 2, 564, 562,
	3, 2, 2, 2, 564, 565, 3, 2, 2, 2, 565, 566, 3, 2, 2, 2, 566, 567, 5, 4,
	3, 2, 567, 568, 7, 8, 2, 2, 568, 65, 3, 2, 2, 2, 569, 572, 5, 42, 22, 2,
	570, 571, 7, 17, 2, 2, 571, 573, 7, 34, 2, 2, 572, 570, 3, 2, 2, 2, 572,
	573, 3, 2, 2, 2, 573, 590, 3, 2, 2, 2, 574, 590, 7, 34, 2, 2, 575, 582,
	5, 30, 16, 2, 576, 577, 7, 17, 2, 2, 577, 580, 7, 34, 2, 2, 578, 579, 7,
	21, 2, 2, 579, 581, 5, 12, 7, 2, 580, 578, 3, 2, 2, 2, 580, 581, 3, 2,
	2, 2, 581, 583, 3, 2, 2, 2, 582, 576, 3, 2, 2, 2, 582, 583, 3, 2, 2, 2,
	583, 590, 3, 2, 2, 2, 584, 587, 7, 34, 2, 2, 585, 586, 7, 21, 2, 2, 586,
	588, 5, 12, 7, 2, 587, 585, 3, 2, 2, 2, 587, 588, 3, 2, 2, 2, 588, 590,
	3, 2, 2, 2, 589, 569, 3, 2, 2, 2, 589, 574, 3, 2, 2, 2, 589, 575, 3, 2,
	2, 2, 589, 584, 3, 2, 2, 2, 590, 67, 3, 2, 2, 2, 591, 593, 7, 32, 2, 2,
	592, 594, 5, 70, 36, 2, 593, 592, 3, 2, 2, 2, 593, 594, 3, 2, 2, 2, 594,
	595, 3, 2, 2, 2, 595, 596, 7, 33, 2, 2, 596, 69, 3, 2, 2, 2, 597, 603,
	5, 72, 37, 2, 598, 599, 5, 74, 38, 2, 599, 600, 5, 72, 37, 2, 600, 602,
	3, 2, 2, 2, 601, 598, 3, 2, 2, 2, 602, 605, 3, 2, 2, 2, 603, 601, 3, 2,
	2, 2, 603, 604, 3, 2, 2, 2, 604, 607, 3, 2, 2, 2, 605, 603, 3, 2, 2, 2,
	606, 608, 5, 74, 38, 2, 607, 606, 3, 2, 2, 2, 607, 608, 3, 2, 2, 2, 608,
	71, 3, 2, 2, 2, 609, 610, 7, 45, 2, 2, 610, 611, 5, 46, 24, 2, 611, 612,
	7, 46, 2, 2, 612, 613, 7, 4, 2, 2, 613, 614, 5, 46, 24, 2, 614, 627, 3,
	2, 2, 2, 615, 618, 7, 67, 2, 2, 616, 617, 7, 21, 2, 2, 617, 619, 5, 12,
	7, 2, 618, 616, 3, 2, 2, 2, 618, 619, 3, 2, 2, 2, 619, 620, 3, 2, 2, 2,
	620, 621, 7, 4, 2, 2, 621, 627, 5, 46, 24, 2, 622, 623, 7, 67, 2, 2, 623,
	624, 7, 4, 2, 2, 624, 627, 5, 22, 12, 2, 625, 627, 5, 46, 24, 2, 626, 609,
	3, 2, 2, 2, 626, 615, 3, 2, 2, 2, 626, 622, 3, 2, 2, 2, 626, 625, 3, 2,
	2, 2, 627, 73, 3, 2, 2, 2, 628, 629, 9, 2, 2, 2, 629, 75, 3, 2, 2, 2, 630,
	631, 7, 47, 2, 2, 631, 77, 3, 2, 2, 2, 632, 633, 7, 48, 2, 2, 633, 79,
	3, 2, 2, 2, 634, 635, 9, 3, 2, 2, 635, 81, 3, 2, 2, 2, 636, 637, 7, 53,
	2, 2, 637, 83, 3, 2, 2, 2, 638, 639, 9, 4, 2, 2, 639, 85, 3, 2, 2, 2, 640,
	641, 9, 5, 2, 2, 641, 87, 3, 2, 2, 2, 642, 643, 9, 6, 2, 2, 643, 89, 3,
	2, 2, 2, 644, 645, 9, 7, 2, 2, 645, 91, 3, 2, 2, 2, 646, 647, 7, 66, 2,
	2, 647, 93, 3, 2, 2, 2, 648, 649, 9, 8, 2, 2, 649, 95, 3, 2, 2, 2, 650,
	651, 9, 9, 2, 2, 651, 97, 3, 2, 2, 2, 69, 104, 108, 146, 151, 163, 189,
	193, 212, 221, 225, 230, 234, 250, 253, 273, 275, 280, 284, 287, 292, 294,
	302, 309, 315, 322, 330, 338, 344, 347, 354, 361, 366, 375, 381, 385, 388,
	399, 404, 411, 419, 427, 446, 483, 485, 492, 499, 506, 514, 519, 525, 534,
	538, 544, 549, 555, 559, 564, 572, 580, 582, 587, 589, 593, 603, 607, 618,
	626,
}
var literalNames = []string{
	"", "';'", "'='", "'break'", "'goto'", "'do'", "'end'", "'while'", "'repeat'",
	"'until'", "'if'", "'then'", "'elseif'", "'else'", "'for'", "','", "'in'",
	"'function'", "'local'", "':'", "'global'", "'<'", "'>'", "'('", "')'",
	"'|'", "'string'", "'boolean'", "'nil'", "'number'", "'{'", "'}'", "'...'",
	"'record'", "'enum'", "'functiontype'", "'return'", "'::'", "'.'", "'false'",
	"'true'", "'as'", "'is'", "'['", "']'", "'or'", "'and'", "'<='", "'>='",
	"'~='", "'=='", "'..'", "'+'", "'-'", "'*'", "'/'", "'%'", "'//'", "'&'",
	"'~'", "'<<'", "'>>'", "'not'", "'#'", "'^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "NAME", "NORMALSTRING", "CHARSTRING",
	"LONGSTRING", "INT", "HEX", "FLOAT", "HEX_FLOAT", "COMMENT", "LINE_COMMENT",
	"WS", "SHEBANG",
}

var ruleNames = []string{
	"chunk", "block", "stat", "attnamelist", "attrib", "typ", "basetype", "typelist",
	"retlist", "typeargs", "newtype", "functiontype", "partypelist", "partype",
	"parnamelist", "parname", "retstat", "label", "funcname", "varlist", "namelist",
	"explist", "exp", "prefixexp", "functioncall", "varOrExp", "variable",
	"varSuffix", "nameAndArgs", "args", "functiondef", "funcbody", "parlist",
	"tableconstructor", "fieldlist", "field", "fieldsep", "operatorOr", "operatorAnd",
	"operatorComparison", "operatorStrcat", "operatorAddSub", "operatorMulDivMod",
	"operatorBitwise", "operatorUnary", "operatorPower", "number", "str",
}

type TealParser struct {
	*antlr.BaseParser
}

// NewTealParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *TealParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewTealParser(input antlr.TokenStream) *TealParser {
	this := new(TealParser)
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
	this.GrammarFileName = "Teal.g4"

	return this
}

// TealParser tokens.
const (
	TealParserEOF          = antlr.TokenEOF
	TealParserT__0         = 1
	TealParserT__1         = 2
	TealParserT__2         = 3
	TealParserT__3         = 4
	TealParserT__4         = 5
	TealParserT__5         = 6
	TealParserT__6         = 7
	TealParserT__7         = 8
	TealParserT__8         = 9
	TealParserT__9         = 10
	TealParserT__10        = 11
	TealParserT__11        = 12
	TealParserT__12        = 13
	TealParserT__13        = 14
	TealParserT__14        = 15
	TealParserT__15        = 16
	TealParserT__16        = 17
	TealParserT__17        = 18
	TealParserT__18        = 19
	TealParserT__19        = 20
	TealParserT__20        = 21
	TealParserT__21        = 22
	TealParserT__22        = 23
	TealParserT__23        = 24
	TealParserT__24        = 25
	TealParserT__25        = 26
	TealParserT__26        = 27
	TealParserT__27        = 28
	TealParserT__28        = 29
	TealParserT__29        = 30
	TealParserT__30        = 31
	TealParserT__31        = 32
	TealParserT__32        = 33
	TealParserT__33        = 34
	TealParserT__34        = 35
	TealParserT__35        = 36
	TealParserT__36        = 37
	TealParserT__37        = 38
	TealParserT__38        = 39
	TealParserT__39        = 40
	TealParserT__40        = 41
	TealParserT__41        = 42
	TealParserT__42        = 43
	TealParserT__43        = 44
	TealParserT__44        = 45
	TealParserT__45        = 46
	TealParserT__46        = 47
	TealParserT__47        = 48
	TealParserT__48        = 49
	TealParserT__49        = 50
	TealParserT__50        = 51
	TealParserT__51        = 52
	TealParserT__52        = 53
	TealParserT__53        = 54
	TealParserT__54        = 55
	TealParserT__55        = 56
	TealParserT__56        = 57
	TealParserT__57        = 58
	TealParserT__58        = 59
	TealParserT__59        = 60
	TealParserT__60        = 61
	TealParserT__61        = 62
	TealParserT__62        = 63
	TealParserT__63        = 64
	TealParserNAME         = 65
	TealParserNORMALSTRING = 66
	TealParserCHARSTRING   = 67
	TealParserLONGSTRING   = 68
	TealParserINT          = 69
	TealParserHEX          = 70
	TealParserFLOAT        = 71
	TealParserHEX_FLOAT    = 72
	TealParserCOMMENT      = 73
	TealParserLINE_COMMENT = 74
	TealParserWS           = 75
	TealParserSHEBANG      = 76
)

// TealParser rules.
const (
	TealParserRULE_chunk              = 0
	TealParserRULE_block              = 1
	TealParserRULE_stat               = 2
	TealParserRULE_attnamelist        = 3
	TealParserRULE_attrib             = 4
	TealParserRULE_typ                = 5
	TealParserRULE_basetype           = 6
	TealParserRULE_typelist           = 7
	TealParserRULE_retlist            = 8
	TealParserRULE_typeargs           = 9
	TealParserRULE_newtype            = 10
	TealParserRULE_functiontype       = 11
	TealParserRULE_partypelist        = 12
	TealParserRULE_partype            = 13
	TealParserRULE_parnamelist        = 14
	TealParserRULE_parname            = 15
	TealParserRULE_retstat            = 16
	TealParserRULE_label              = 17
	TealParserRULE_funcname           = 18
	TealParserRULE_varlist            = 19
	TealParserRULE_namelist           = 20
	TealParserRULE_explist            = 21
	TealParserRULE_exp                = 22
	TealParserRULE_prefixexp          = 23
	TealParserRULE_functioncall       = 24
	TealParserRULE_varOrExp           = 25
	TealParserRULE_variable           = 26
	TealParserRULE_varSuffix          = 27
	TealParserRULE_nameAndArgs        = 28
	TealParserRULE_args               = 29
	TealParserRULE_functiondef        = 30
	TealParserRULE_funcbody           = 31
	TealParserRULE_parlist            = 32
	TealParserRULE_tableconstructor   = 33
	TealParserRULE_fieldlist          = 34
	TealParserRULE_field              = 35
	TealParserRULE_fieldsep           = 36
	TealParserRULE_operatorOr         = 37
	TealParserRULE_operatorAnd        = 38
	TealParserRULE_operatorComparison = 39
	TealParserRULE_operatorStrcat     = 40
	TealParserRULE_operatorAddSub     = 41
	TealParserRULE_operatorMulDivMod  = 42
	TealParserRULE_operatorBitwise    = 43
	TealParserRULE_operatorUnary      = 44
	TealParserRULE_operatorPower      = 45
	TealParserRULE_number             = 46
	TealParserRULE_str                = 47
)

// IChunkContext is an interface to support dynamic dispatch.
type IChunkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChunkContext differentiates from other interfaces.
	IsChunkContext()
}

type ChunkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChunkContext() *ChunkContext {
	var p = new(ChunkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_chunk
	return p
}

func (*ChunkContext) IsChunkContext() {}

func NewChunkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChunkContext {
	var p = new(ChunkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_chunk

	return p
}

func (s *ChunkContext) GetParser() antlr.Parser { return s.parser }

func (s *ChunkContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ChunkContext) EOF() antlr.TerminalNode {
	return s.GetToken(TealParserEOF, 0)
}

func (s *ChunkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChunkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChunkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterChunk(s)
	}
}

func (s *ChunkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitChunk(s)
	}
}

func (p *TealParser) Chunk() (localctx IChunkContext) {
	this := p
	_ = this

	localctx = NewChunkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TealParserRULE_chunk)

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
		p.SetState(96)
		p.Block()
	}
	{
		p.SetState(97)
		p.Match(TealParserEOF)
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
	p.RuleIndex = TealParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) Retstat() IRetstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetstatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *TealParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TealParserRULE_block)
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
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TealParserT__0)|(1<<TealParserT__2)|(1<<TealParserT__3)|(1<<TealParserT__4)|(1<<TealParserT__6)|(1<<TealParserT__7)|(1<<TealParserT__9)|(1<<TealParserT__13)|(1<<TealParserT__16)|(1<<TealParserT__17)|(1<<TealParserT__19)|(1<<TealParserT__22))) != 0) || _la == TealParserT__36 || _la == TealParserNAME {
		{
			p.SetState(99)
			p.Stat()
		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__35 {
		{
			p.SetState(105)
			p.Retstat()
		}

	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyFrom(ctx *StatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RepeatStatContext struct {
	*StatContext
}

func NewRepeatStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepeatStatContext {
	var p = new(RepeatStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *RepeatStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RepeatStatContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *RepeatStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterRepeatStat(s)
	}
}

func (s *RepeatStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitRepeatStat(s)
	}
}

type BreakStatContext struct {
	*StatContext
}

func NewBreakStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatContext {
	var p = new(BreakStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *BreakStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterBreakStat(s)
	}
}

func (s *BreakStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitBreakStat(s)
	}
}

type FuncStatContext struct {
	*StatContext
}

func NewFuncStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncStatContext {
	var p = new(FuncStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *FuncStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncStatContext) Funcname() IFuncnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncnameContext)
}

func (s *FuncStatContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *FuncStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterFuncStat(s)
	}
}

func (s *FuncStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitFuncStat(s)
	}
}

type SemiStatContext struct {
	*StatContext
}

func NewSemiStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SemiStatContext {
	var p = new(SemiStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *SemiStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SemiStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterSemiStat(s)
	}
}

func (s *SemiStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitSemiStat(s)
	}
}

type ForInStatContext struct {
	*StatContext
}

func NewForInStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInStatContext {
	var p = new(ForInStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *ForInStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInStatContext) Namelist() INamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *ForInStatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *ForInStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForInStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterForInStat(s)
	}
}

func (s *ForInStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitForInStat(s)
	}
}

type GlobalAttrStatContext struct {
	*StatContext
}

func NewGlobalAttrStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GlobalAttrStatContext {
	var p = new(GlobalAttrStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *GlobalAttrStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalAttrStatContext) Attnamelist() IAttnamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttnamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttnamelistContext)
}

func (s *GlobalAttrStatContext) Typelist() ITypelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypelistContext)
}

func (s *GlobalAttrStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterGlobalAttrStat(s)
	}
}

func (s *GlobalAttrStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitGlobalAttrStat(s)
	}
}

type IfStatContext struct {
	*StatContext
}

func NewIfStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatContext {
	var p = new(IfStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *IfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *IfStatContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IfStatContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfStatContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterIfStat(s)
	}
}

func (s *IfStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitIfStat(s)
	}
}

type LocalFuncStatContext struct {
	*StatContext
}

func NewLocalFuncStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalFuncStatContext {
	var p = new(LocalFuncStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *LocalFuncStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalFuncStatContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *LocalFuncStatContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *LocalFuncStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterLocalFuncStat(s)
	}
}

func (s *LocalFuncStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitLocalFuncStat(s)
	}
}

type WhileStatContext struct {
	*StatContext
}

func NewWhileStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatContext {
	var p = new(WhileStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *WhileStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *WhileStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterWhileStat(s)
	}
}

func (s *WhileStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitWhileStat(s)
	}
}

type GotoStatContext struct {
	*StatContext
}

func NewGotoStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GotoStatContext {
	var p = new(GotoStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *GotoStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotoStatContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *GotoStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterGotoStat(s)
	}
}

func (s *GotoStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitGotoStat(s)
	}
}

type LocalAttrAssignStatContext struct {
	*StatContext
}

func NewLocalAttrAssignStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalAttrAssignStatContext {
	var p = new(LocalAttrAssignStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *LocalAttrAssignStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalAttrAssignStatContext) Attnamelist() IAttnamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttnamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttnamelistContext)
}

func (s *LocalAttrAssignStatContext) Typelist() ITypelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypelistContext)
}

func (s *LocalAttrAssignStatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *LocalAttrAssignStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterLocalAttrAssignStat(s)
	}
}

func (s *LocalAttrAssignStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitLocalAttrAssignStat(s)
	}
}

type LocalNewTypeStatContext struct {
	*StatContext
}

func NewLocalNewTypeStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalNewTypeStatContext {
	var p = new(LocalNewTypeStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *LocalNewTypeStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalNewTypeStatContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *LocalNewTypeStatContext) Newtype() INewtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewtypeContext)
}

func (s *LocalNewTypeStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterLocalNewTypeStat(s)
	}
}

func (s *LocalNewTypeStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitLocalNewTypeStat(s)
	}
}

type GlobalAttrAssignStatContext struct {
	*StatContext
}

func NewGlobalAttrAssignStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GlobalAttrAssignStatContext {
	var p = new(GlobalAttrAssignStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *GlobalAttrAssignStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalAttrAssignStatContext) Attnamelist() IAttnamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttnamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttnamelistContext)
}

func (s *GlobalAttrAssignStatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *GlobalAttrAssignStatContext) Typelist() ITypelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypelistContext)
}

func (s *GlobalAttrAssignStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterGlobalAttrAssignStat(s)
	}
}

func (s *GlobalAttrAssignStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitGlobalAttrAssignStat(s)
	}
}

type LabelStatContext struct {
	*StatContext
}

func NewLabelStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LabelStatContext {
	var p = new(LabelStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *LabelStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelStatContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LabelStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterLabelStat(s)
	}
}

func (s *LabelStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitLabelStat(s)
	}
}

type ForStatContext struct {
	*StatContext
}

func NewForStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStatContext {
	var p = new(ForStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *ForStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *ForStatContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ForStatContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ForStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterForStat(s)
	}
}

func (s *ForStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitForStat(s)
	}
}

type GlobalFuncStatContext struct {
	*StatContext
}

func NewGlobalFuncStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GlobalFuncStatContext {
	var p = new(GlobalFuncStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *GlobalFuncStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalFuncStatContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *GlobalFuncStatContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *GlobalFuncStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterGlobalFuncStat(s)
	}
}

func (s *GlobalFuncStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitGlobalFuncStat(s)
	}
}

type AssignStatContext struct {
	*StatContext
}

func NewAssignStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStatContext {
	var p = new(AssignStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *AssignStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatContext) Varlist() IVarlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarlistContext)
}

func (s *AssignStatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *AssignStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterAssignStat(s)
	}
}

func (s *AssignStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitAssignStat(s)
	}
}

type DoStatContext struct {
	*StatContext
}

func NewDoStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoStatContext {
	var p = new(DoStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *DoStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoStatContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DoStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterDoStat(s)
	}
}

func (s *DoStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitDoStat(s)
	}
}

type GlobalAssignStatContext struct {
	*StatContext
}

func NewGlobalAssignStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GlobalAssignStatContext {
	var p = new(GlobalAssignStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *GlobalAssignStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalAssignStatContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *GlobalAssignStatContext) Newtype() INewtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewtypeContext)
}

func (s *GlobalAssignStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterGlobalAssignStat(s)
	}
}

func (s *GlobalAssignStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitGlobalAssignStat(s)
	}
}

type FuncCallStatContext struct {
	*StatContext
}

func NewFuncCallStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallStatContext {
	var p = new(FuncCallStatContext)

	p.StatContext = NewEmptyStatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatContext))

	return p
}

func (s *FuncCallStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallStatContext) Functioncall() IFunctioncallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctioncallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctioncallContext)
}

func (s *FuncCallStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterFuncCallStat(s)
	}
}

func (s *FuncCallStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitFuncCallStat(s)
	}
}

func (p *TealParser) Stat() (localctx IStatContext) {
	this := p
	_ = this

	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TealParserRULE_stat)
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

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSemiStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(TealParserT__0)
		}

	case 2:
		localctx = NewAssignStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Varlist()
		}
		{
			p.SetState(110)
			p.Match(TealParserT__1)
		}
		{
			p.SetState(111)
			p.Explist()
		}

	case 3:
		localctx = NewFuncCallStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.Functioncall()
		}

	case 4:
		localctx = NewLabelStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.Label()
		}

	case 5:
		localctx = NewBreakStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(115)
			p.Match(TealParserT__2)
		}

	case 6:
		localctx = NewGotoStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(116)
			p.Match(TealParserT__3)
		}
		{
			p.SetState(117)
			p.Match(TealParserNAME)
		}

	case 7:
		localctx = NewDoStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(118)
			p.Match(TealParserT__4)
		}
		{
			p.SetState(119)
			p.Block()
		}
		{
			p.SetState(120)
			p.Match(TealParserT__5)
		}

	case 8:
		localctx = NewWhileStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(122)
			p.Match(TealParserT__6)
		}
		{
			p.SetState(123)
			p.exp(0)
		}
		{
			p.SetState(124)
			p.Match(TealParserT__4)
		}
		{
			p.SetState(125)
			p.Block()
		}
		{
			p.SetState(126)
			p.Match(TealParserT__5)
		}

	case 9:
		localctx = NewRepeatStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(128)
			p.Match(TealParserT__7)
		}
		{
			p.SetState(129)
			p.Block()
		}
		{
			p.SetState(130)
			p.Match(TealParserT__8)
		}
		{
			p.SetState(131)
			p.exp(0)
		}

	case 10:
		localctx = NewIfStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(133)
			p.Match(TealParserT__9)
		}
		{
			p.SetState(134)
			p.exp(0)
		}
		{
			p.SetState(135)
			p.Match(TealParserT__10)
		}
		{
			p.SetState(136)
			p.Block()
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TealParserT__11 {
			{
				p.SetState(137)
				p.Match(TealParserT__11)
			}
			{
				p.SetState(138)
				p.exp(0)
			}
			{
				p.SetState(139)
				p.Match(TealParserT__10)
			}
			{
				p.SetState(140)
				p.Block()
			}

			p.SetState(146)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__12 {
			{
				p.SetState(147)
				p.Match(TealParserT__12)
			}
			{
				p.SetState(148)
				p.Block()
			}

		}
		{
			p.SetState(151)
			p.Match(TealParserT__5)
		}

	case 11:
		localctx = NewForStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(153)
			p.Match(TealParserT__13)
		}
		{
			p.SetState(154)
			p.Match(TealParserNAME)
		}
		{
			p.SetState(155)
			p.Match(TealParserT__1)
		}
		{
			p.SetState(156)
			p.exp(0)
		}
		{
			p.SetState(157)
			p.Match(TealParserT__14)
		}
		{
			p.SetState(158)
			p.exp(0)
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__14 {
			{
				p.SetState(159)
				p.Match(TealParserT__14)
			}
			{
				p.SetState(160)
				p.exp(0)
			}

		}
		{
			p.SetState(163)
			p.Match(TealParserT__4)
		}
		{
			p.SetState(164)
			p.Block()
		}
		{
			p.SetState(165)
			p.Match(TealParserT__5)
		}

	case 12:
		localctx = NewForInStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(167)
			p.Match(TealParserT__13)
		}
		{
			p.SetState(168)
			p.Namelist()
		}
		{
			p.SetState(169)
			p.Match(TealParserT__15)
		}
		{
			p.SetState(170)
			p.Explist()
		}
		{
			p.SetState(171)
			p.Match(TealParserT__4)
		}
		{
			p.SetState(172)
			p.Block()
		}
		{
			p.SetState(173)
			p.Match(TealParserT__5)
		}

	case 13:
		localctx = NewFuncStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(175)
			p.Match(TealParserT__16)
		}
		{
			p.SetState(176)
			p.Funcname()
		}
		{
			p.SetState(177)
			p.Funcbody()
		}

	case 14:
		localctx = NewLocalFuncStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(179)
			p.Match(TealParserT__17)
		}
		{
			p.SetState(180)
			p.Match(TealParserT__16)
		}
		{
			p.SetState(181)
			p.Match(TealParserNAME)
		}
		{
			p.SetState(182)
			p.Funcbody()
		}

	case 15:
		localctx = NewLocalAttrAssignStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(183)
			p.Match(TealParserT__17)
		}
		{
			p.SetState(184)
			p.Attnamelist()
		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__18 {
			{
				p.SetState(185)
				p.Match(TealParserT__18)
			}
			{
				p.SetState(186)
				p.Typelist()
			}

		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__1 {
			{
				p.SetState(189)
				p.Match(TealParserT__1)
			}
			{
				p.SetState(190)
				p.Explist()
			}

		}

	case 16:
		localctx = NewLocalNewTypeStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(193)
			p.Match(TealParserT__17)
		}
		{
			p.SetState(194)
			p.Match(TealParserNAME)
		}
		{
			p.SetState(195)
			p.Match(TealParserT__1)
		}
		{
			p.SetState(196)
			p.Newtype()
		}

	case 17:
		localctx = NewGlobalFuncStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(197)
			p.Match(TealParserT__19)
		}
		{
			p.SetState(198)
			p.Match(TealParserT__16)
		}
		{
			p.SetState(199)
			p.Match(TealParserNAME)
		}
		{
			p.SetState(200)
			p.Funcbody()
		}

	case 18:
		localctx = NewGlobalAttrStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(201)
			p.Match(TealParserT__19)
		}
		{
			p.SetState(202)
			p.Attnamelist()
		}
		{
			p.SetState(203)
			p.Match(TealParserT__18)
		}
		{
			p.SetState(204)
			p.Typelist()
		}

	case 19:
		localctx = NewGlobalAttrAssignStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(206)
			p.Match(TealParserT__19)
		}
		{
			p.SetState(207)
			p.Attnamelist()
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__18 {
			{
				p.SetState(208)
				p.Match(TealParserT__18)
			}
			{
				p.SetState(209)
				p.Typelist()
			}

		}
		{
			p.SetState(212)
			p.Match(TealParserT__1)
		}
		{
			p.SetState(213)
			p.Explist()
		}

	case 20:
		localctx = NewGlobalAssignStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(215)
			p.Match(TealParserT__19)
		}
		{
			p.SetState(216)
			p.Match(TealParserNAME)
		}
		{
			p.SetState(217)
			p.Match(TealParserT__1)
		}
		{
			p.SetState(218)
			p.Newtype()
		}

	}

	return localctx
}

// IAttnamelistContext is an interface to support dynamic dispatch.
type IAttnamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttnamelistContext differentiates from other interfaces.
	IsAttnamelistContext()
}

type AttnamelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttnamelistContext() *AttnamelistContext {
	var p = new(AttnamelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_attnamelist
	return p
}

func (*AttnamelistContext) IsAttnamelistContext() {}

func NewAttnamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttnamelistContext {
	var p = new(AttnamelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_attnamelist

	return p
}

func (s *AttnamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *AttnamelistContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(TealParserNAME)
}

func (s *AttnamelistContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(TealParserNAME, i)
}

func (s *AttnamelistContext) AllAttrib() []IAttribContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttribContext)(nil)).Elem())
	var tst = make([]IAttribContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttribContext)
		}
	}

	return tst
}

func (s *AttnamelistContext) Attrib(i int) IAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *AttnamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttnamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttnamelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterAttnamelist(s)
	}
}

func (s *AttnamelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitAttnamelist(s)
	}
}

func (p *TealParser) Attnamelist() (localctx IAttnamelistContext) {
	this := p
	_ = this

	localctx = NewAttnamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TealParserRULE_attnamelist)
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
		p.SetState(221)
		p.Match(TealParserNAME)
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__20 {
		{
			p.SetState(222)
			p.Attrib()
		}

	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TealParserT__14 {
		{
			p.SetState(225)
			p.Match(TealParserT__14)
		}
		{
			p.SetState(226)
			p.Match(TealParserNAME)
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__20 {
			{
				p.SetState(227)
				p.Attrib()
			}

		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = TealParserRULE_attrib
	return p
}

func (*AttribContext) IsAttribContext() {}

func NewAttribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribContext {
	var p = new(AttribContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_attrib

	return p
}

func (s *AttribContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *AttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterAttrib(s)
	}
}

func (s *AttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitAttrib(s)
	}
}

func (p *TealParser) Attrib() (localctx IAttribContext) {
	this := p
	_ = this

	localctx = NewAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TealParserRULE_attrib)

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
		p.SetState(235)
		p.Match(TealParserT__20)
	}
	{
		p.SetState(236)
		p.Match(TealParserNAME)
	}
	{
		p.SetState(237)
		p.Match(TealParserT__21)
	}

	return localctx
}

// ITypContext is an interface to support dynamic dispatch.
type ITypContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypContext differentiates from other interfaces.
	IsTypContext()
}

type TypContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypContext() *TypContext {
	var p = new(TypContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_typ
	return p
}

func (*TypContext) IsTypContext() {}

func NewTypContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypContext {
	var p = new(TypContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_typ

	return p
}

func (s *TypContext) GetParser() antlr.Parser { return s.parser }

func (s *TypContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *TypContext) AllBasetype() []IBasetypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBasetypeContext)(nil)).Elem())
	var tst = make([]IBasetypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBasetypeContext)
		}
	}

	return tst
}

func (s *TypContext) Basetype(i int) IBasetypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasetypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBasetypeContext)
}

func (s *TypContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterTyp(s)
	}
}

func (s *TypContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitTyp(s)
	}
}

func (p *TealParser) Typ() (localctx ITypContext) {
	this := p
	_ = this

	localctx = NewTypContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TealParserRULE_typ)

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

	p.SetState(251)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TealParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			p.Match(TealParserT__22)
		}
		{
			p.SetState(240)
			p.Typ()
		}
		{
			p.SetState(241)
			p.Match(TealParserT__23)
		}

	case TealParserT__16, TealParserT__25, TealParserT__26, TealParserT__27, TealParserT__28, TealParserT__29, TealParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.Basetype()
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(244)
					p.Match(TealParserT__24)
				}
				{
					p.SetState(245)
					p.Basetype()
				}

			}
			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBasetypeContext is an interface to support dynamic dispatch.
type IBasetypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasetypeContext differentiates from other interfaces.
	IsBasetypeContext()
}

type BasetypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasetypeContext() *BasetypeContext {
	var p = new(BasetypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_basetype
	return p
}

func (*BasetypeContext) IsBasetypeContext() {}

func NewBasetypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasetypeContext {
	var p = new(BasetypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_basetype

	return p
}

func (s *BasetypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BasetypeContext) AllTyp() []ITypContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypContext)(nil)).Elem())
	var tst = make([]ITypContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypContext)
		}
	}

	return tst
}

func (s *BasetypeContext) Typ(i int) ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *BasetypeContext) Functiontype() IFunctiontypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctiontypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctiontypeContext)
}

func (s *BasetypeContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *BasetypeContext) Typeargs() ITypeargsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeargsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeargsContext)
}

func (s *BasetypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasetypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasetypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterBasetype(s)
	}
}

func (s *BasetypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitBasetype(s)
	}
}

func (p *TealParser) Basetype() (localctx IBasetypeContext) {
	this := p
	_ = this

	localctx = NewBasetypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TealParserRULE_basetype)

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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.Match(TealParserT__25)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Match(TealParserT__26)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(255)
			p.Match(TealParserT__27)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(256)
			p.Match(TealParserT__28)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(257)
			p.Match(TealParserT__29)
		}
		{
			p.SetState(258)
			p.Typ()
		}
		{
			p.SetState(259)
			p.Match(TealParserT__30)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(261)
			p.Match(TealParserT__29)
		}
		{
			p.SetState(262)
			p.Typ()
		}
		{
			p.SetState(263)
			p.Match(TealParserT__18)
		}
		{
			p.SetState(264)
			p.Typ()
		}
		{
			p.SetState(265)
			p.Match(TealParserT__30)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(267)
			p.Match(TealParserT__16)
		}
		{
			p.SetState(268)
			p.Functiontype()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(269)
			p.Match(TealParserNAME)
		}
		p.SetState(271)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(270)
				p.Typeargs()
			}

		}

	}

	return localctx
}

// ITypelistContext is an interface to support dynamic dispatch.
type ITypelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypelistContext differentiates from other interfaces.
	IsTypelistContext()
}

type TypelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypelistContext() *TypelistContext {
	var p = new(TypelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_typelist
	return p
}

func (*TypelistContext) IsTypelistContext() {}

func NewTypelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypelistContext {
	var p = new(TypelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_typelist

	return p
}

func (s *TypelistContext) GetParser() antlr.Parser { return s.parser }

func (s *TypelistContext) AllTyp() []ITypContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypContext)(nil)).Elem())
	var tst = make([]ITypContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypContext)
		}
	}

	return tst
}

func (s *TypelistContext) Typ(i int) ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *TypelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterTypelist(s)
	}
}

func (s *TypelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitTypelist(s)
	}
}

func (p *TealParser) Typelist() (localctx ITypelistContext) {
	this := p
	_ = this

	localctx = NewTypelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TealParserRULE_typelist)

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
		p.SetState(275)
		p.Typ()
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(276)
			p.Match(TealParserT__14)
		}
		{
			p.SetState(277)
			p.Typ()
		}

	}

	return localctx
}

// IRetlistContext is an interface to support dynamic dispatch.
type IRetlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetlistContext differentiates from other interfaces.
	IsRetlistContext()
}

type RetlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetlistContext() *RetlistContext {
	var p = new(RetlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_retlist
	return p
}

func (*RetlistContext) IsRetlistContext() {}

func NewRetlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetlistContext {
	var p = new(RetlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_retlist

	return p
}

func (s *RetlistContext) GetParser() antlr.Parser { return s.parser }

func (s *RetlistContext) Typelist() ITypelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypelistContext)
}

func (s *RetlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterRetlist(s)
	}
}

func (s *RetlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitRetlist(s)
	}
}

func (p *TealParser) Retlist() (localctx IRetlistContext) {
	this := p
	_ = this

	localctx = NewRetlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TealParserRULE_retlist)
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

	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.Match(TealParserT__22)
		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TealParserT__16)|(1<<TealParserT__22)|(1<<TealParserT__25)|(1<<TealParserT__26)|(1<<TealParserT__27)|(1<<TealParserT__28)|(1<<TealParserT__29))) != 0) || _la == TealParserNAME {
			{
				p.SetState(281)
				p.Typelist()
			}

		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__31 {
			{
				p.SetState(284)
				p.Match(TealParserT__31)
			}

		}
		{
			p.SetState(287)
			p.Match(TealParserT__23)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(288)
			p.Typelist()
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(289)
				p.Match(TealParserT__31)
			}

		}

	}

	return localctx
}

// ITypeargsContext is an interface to support dynamic dispatch.
type ITypeargsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeargsContext differentiates from other interfaces.
	IsTypeargsContext()
}

type TypeargsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeargsContext() *TypeargsContext {
	var p = new(TypeargsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_typeargs
	return p
}

func (*TypeargsContext) IsTypeargsContext() {}

func NewTypeargsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeargsContext {
	var p = new(TypeargsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_typeargs

	return p
}

func (s *TypeargsContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeargsContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(TealParserNAME)
}

func (s *TypeargsContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(TealParserNAME, i)
}

func (s *TypeargsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeargsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeargsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterTypeargs(s)
	}
}

func (s *TypeargsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitTypeargs(s)
	}
}

func (p *TealParser) Typeargs() (localctx ITypeargsContext) {
	this := p
	_ = this

	localctx = NewTypeargsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TealParserRULE_typeargs)
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
		p.SetState(294)
		p.Match(TealParserT__20)
	}
	{
		p.SetState(295)
		p.Match(TealParserNAME)
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TealParserT__14 {
		{
			p.SetState(296)
			p.Match(TealParserT__14)
		}
		{
			p.SetState(297)
			p.Match(TealParserNAME)
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(303)
		p.Match(TealParserT__21)
	}

	return localctx
}

// INewtypeContext is an interface to support dynamic dispatch.
type INewtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNewtypeContext differentiates from other interfaces.
	IsNewtypeContext()
}

type NewtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewtypeContext() *NewtypeContext {
	var p = new(NewtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_newtype
	return p
}

func (*NewtypeContext) IsNewtypeContext() {}

func NewNewtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewtypeContext {
	var p = new(NewtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_newtype

	return p
}

func (s *NewtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NewtypeContext) CopyFrom(ctx *NewtypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NewtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RecordNewTypeContext struct {
	*NewtypeContext
}

func NewRecordNewTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecordNewTypeContext {
	var p = new(RecordNewTypeContext)

	p.NewtypeContext = NewEmptyNewtypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NewtypeContext))

	return p
}

func (s *RecordNewTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordNewTypeContext) Typeargs() ITypeargsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeargsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeargsContext)
}

func (s *RecordNewTypeContext) AllTyp() []ITypContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypContext)(nil)).Elem())
	var tst = make([]ITypContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypContext)
		}
	}

	return tst
}

func (s *RecordNewTypeContext) Typ(i int) ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *RecordNewTypeContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(TealParserNAME)
}

func (s *RecordNewTypeContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(TealParserNAME, i)
}

func (s *RecordNewTypeContext) AllNewtype() []INewtypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewtypeContext)(nil)).Elem())
	var tst = make([]INewtypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewtypeContext)
		}
	}

	return tst
}

func (s *RecordNewTypeContext) Newtype(i int) INewtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewtypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewtypeContext)
}

func (s *RecordNewTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterRecordNewType(s)
	}
}

func (s *RecordNewTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitRecordNewType(s)
	}
}

type EnumNewTypeContext struct {
	*NewtypeContext
}

func NewEnumNewTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumNewTypeContext {
	var p = new(EnumNewTypeContext)

	p.NewtypeContext = NewEmptyNewtypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NewtypeContext))

	return p
}

func (s *EnumNewTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNewTypeContext) AllStr() []IStrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStrContext)(nil)).Elem())
	var tst = make([]IStrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStrContext)
		}
	}

	return tst
}

func (s *EnumNewTypeContext) Str(i int) IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *EnumNewTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterEnumNewType(s)
	}
}

func (s *EnumNewTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitEnumNewType(s)
	}
}

type FuncNewTypeContext struct {
	*NewtypeContext
}

func NewFuncNewTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncNewTypeContext {
	var p = new(FuncNewTypeContext)

	p.NewtypeContext = NewEmptyNewtypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NewtypeContext))

	return p
}

func (s *FuncNewTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNewTypeContext) Functiontype() IFunctiontypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctiontypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctiontypeContext)
}

func (s *FuncNewTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterFuncNewType(s)
	}
}

func (s *FuncNewTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitFuncNewType(s)
	}
}

func (p *TealParser) Newtype() (localctx INewtypeContext) {
	this := p
	_ = this

	localctx = NewNewtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TealParserRULE_newtype)
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

	p.SetState(342)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TealParserT__32:
		localctx = NewRecordNewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)
			p.Match(TealParserT__32)
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__20 {
			{
				p.SetState(306)
				p.Typeargs()
			}

		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__29 {
			{
				p.SetState(309)
				p.Match(TealParserT__29)
			}
			{
				p.SetState(310)
				p.Typ()
			}
			{
				p.SetState(311)
				p.Match(TealParserT__30)
			}

		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(315)
					p.Match(TealParserNAME)
				}
				{
					p.SetState(316)
					p.Match(TealParserT__1)
				}
				{
					p.SetState(317)
					p.Newtype()
				}

			}
			p.SetState(322)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
		}
		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TealParserNAME {
			{
				p.SetState(323)
				p.Match(TealParserNAME)
			}
			{
				p.SetState(324)
				p.Match(TealParserT__18)
			}
			{
				p.SetState(325)
				p.Typ()
			}

			p.SetState(330)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(331)
			p.Match(TealParserT__5)
		}

	case TealParserT__33:
		localctx = NewEnumNewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.Match(TealParserT__33)
		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(TealParserNORMALSTRING-66))|(1<<(TealParserCHARSTRING-66))|(1<<(TealParserLONGSTRING-66)))) != 0 {
			{
				p.SetState(333)
				p.Str()
			}

			p.SetState(338)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(339)
			p.Match(TealParserT__5)
		}

	case TealParserT__34:
		localctx = NewFuncNewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(340)
			p.Match(TealParserT__34)
		}
		{
			p.SetState(341)
			p.Functiontype()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctiontypeContext is an interface to support dynamic dispatch.
type IFunctiontypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctiontypeContext differentiates from other interfaces.
	IsFunctiontypeContext()
}

type FunctiontypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctiontypeContext() *FunctiontypeContext {
	var p = new(FunctiontypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_functiontype
	return p
}

func (*FunctiontypeContext) IsFunctiontypeContext() {}

func NewFunctiontypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctiontypeContext {
	var p = new(FunctiontypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_functiontype

	return p
}

func (s *FunctiontypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctiontypeContext) Partypelist() IPartypelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartypelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPartypelistContext)
}

func (s *FunctiontypeContext) Typeargs() ITypeargsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeargsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeargsContext)
}

func (s *FunctiontypeContext) Retlist() IRetlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetlistContext)
}

func (s *FunctiontypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctiontypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctiontypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterFunctiontype(s)
	}
}

func (s *FunctiontypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitFunctiontype(s)
	}
}

func (p *TealParser) Functiontype() (localctx IFunctiontypeContext) {
	this := p
	_ = this

	localctx = NewFunctiontypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TealParserRULE_functiontype)
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
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__20 {
		{
			p.SetState(344)
			p.Typeargs()
		}

	}
	{
		p.SetState(347)
		p.Match(TealParserT__22)
	}
	{
		p.SetState(348)
		p.Partypelist()
	}
	{
		p.SetState(349)
		p.Match(TealParserT__23)
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(350)
			p.Match(TealParserT__18)
		}
		{
			p.SetState(351)
			p.Retlist()
		}

	}

	return localctx
}

// IPartypelistContext is an interface to support dynamic dispatch.
type IPartypelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPartypelistContext differentiates from other interfaces.
	IsPartypelistContext()
}

type PartypelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartypelistContext() *PartypelistContext {
	var p = new(PartypelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_partypelist
	return p
}

func (*PartypelistContext) IsPartypelistContext() {}

func NewPartypelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartypelistContext {
	var p = new(PartypelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_partypelist

	return p
}

func (s *PartypelistContext) GetParser() antlr.Parser { return s.parser }

func (s *PartypelistContext) AllPartype() []IPartypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPartypeContext)(nil)).Elem())
	var tst = make([]IPartypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPartypeContext)
		}
	}

	return tst
}

func (s *PartypelistContext) Partype(i int) IPartypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPartypeContext)
}

func (s *PartypelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartypelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PartypelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterPartypelist(s)
	}
}

func (s *PartypelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitPartypelist(s)
	}
}

func (p *TealParser) Partypelist() (localctx IPartypelistContext) {
	this := p
	_ = this

	localctx = NewPartypelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TealParserRULE_partypelist)
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
		p.SetState(354)
		p.Partype()
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TealParserT__14 {
		{
			p.SetState(355)
			p.Match(TealParserT__14)
		}
		{
			p.SetState(356)
			p.Partype()
		}

		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPartypeContext is an interface to support dynamic dispatch.
type IPartypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPartypeContext differentiates from other interfaces.
	IsPartypeContext()
}

type PartypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartypeContext() *PartypeContext {
	var p = new(PartypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_partype
	return p
}

func (*PartypeContext) IsPartypeContext() {}

func NewPartypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartypeContext {
	var p = new(PartypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_partype

	return p
}

func (s *PartypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PartypeContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *PartypeContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *PartypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PartypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterPartype(s)
	}
}

func (s *PartypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitPartype(s)
	}
}

func (p *TealParser) Partype() (localctx IPartypeContext) {
	this := p
	_ = this

	localctx = NewPartypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TealParserRULE_partype)

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
	p.SetState(364)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(362)
			p.Match(TealParserNAME)
		}
		{
			p.SetState(363)
			p.Match(TealParserT__18)
		}

	}
	{
		p.SetState(366)
		p.Typ()
	}

	return localctx
}

// IParnamelistContext is an interface to support dynamic dispatch.
type IParnamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParnamelistContext differentiates from other interfaces.
	IsParnamelistContext()
}

type ParnamelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParnamelistContext() *ParnamelistContext {
	var p = new(ParnamelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_parnamelist
	return p
}

func (*ParnamelistContext) IsParnamelistContext() {}

func NewParnamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParnamelistContext {
	var p = new(ParnamelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_parnamelist

	return p
}

func (s *ParnamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParnamelistContext) AllParname() []IParnameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParnameContext)(nil)).Elem())
	var tst = make([]IParnameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParnameContext)
		}
	}

	return tst
}

func (s *ParnamelistContext) Parname(i int) IParnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParnameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParnameContext)
}

func (s *ParnamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParnamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParnamelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterParnamelist(s)
	}
}

func (s *ParnamelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitParnamelist(s)
	}
}

func (p *TealParser) Parnamelist() (localctx IParnamelistContext) {
	this := p
	_ = this

	localctx = NewParnamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TealParserRULE_parnamelist)

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
		p.SetState(368)
		p.Parname()
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(369)
				p.Match(TealParserT__14)
			}
			{
				p.SetState(370)
				p.Parname()
			}

		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
	}

	return localctx
}

// IParnameContext is an interface to support dynamic dispatch.
type IParnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParnameContext differentiates from other interfaces.
	IsParnameContext()
}

type ParnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParnameContext() *ParnameContext {
	var p = new(ParnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_parname
	return p
}

func (*ParnameContext) IsParnameContext() {}

func NewParnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParnameContext {
	var p = new(ParnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_parname

	return p
}

func (s *ParnameContext) GetParser() antlr.Parser { return s.parser }

func (s *ParnameContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *ParnameContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *ParnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterParname(s)
	}
}

func (s *ParnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitParname(s)
	}
}

func (p *TealParser) Parname() (localctx IParnameContext) {
	this := p
	_ = this

	localctx = NewParnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TealParserRULE_parname)
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
		p.SetState(376)
		p.Match(TealParserNAME)
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__18 {
		{
			p.SetState(377)
			p.Match(TealParserT__18)
		}
		{
			p.SetState(378)
			p.Typ()
		}

	}

	return localctx
}

// IRetstatContext is an interface to support dynamic dispatch.
type IRetstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetstatContext differentiates from other interfaces.
	IsRetstatContext()
}

type RetstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetstatContext() *RetstatContext {
	var p = new(RetstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_retstat
	return p
}

func (*RetstatContext) IsRetstatContext() {}

func NewRetstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetstatContext {
	var p = new(RetstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_retstat

	return p
}

func (s *RetstatContext) GetParser() antlr.Parser { return s.parser }

func (s *RetstatContext) CopyFrom(ctx *RetstatContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RetstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ReturnStatContext struct {
	*RetstatContext
}

func NewReturnStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatContext {
	var p = new(ReturnStatContext)

	p.RetstatContext = NewEmptyRetstatContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RetstatContext))

	return p
}

func (s *ReturnStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *ReturnStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterReturnStat(s)
	}
}

func (s *ReturnStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitReturnStat(s)
	}
}

func (p *TealParser) Retstat() (localctx IRetstatContext) {
	this := p
	_ = this

	localctx = NewRetstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TealParserRULE_retstat)
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

	localctx = NewReturnStatContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(TealParserT__35)
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(TealParserT__16-17))|(1<<(TealParserT__22-17))|(1<<(TealParserT__27-17))|(1<<(TealParserT__29-17))|(1<<(TealParserT__31-17))|(1<<(TealParserT__38-17))|(1<<(TealParserT__39-17)))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(TealParserT__52-53))|(1<<(TealParserT__58-53))|(1<<(TealParserT__61-53))|(1<<(TealParserT__62-53))|(1<<(TealParserNAME-53))|(1<<(TealParserNORMALSTRING-53))|(1<<(TealParserCHARSTRING-53))|(1<<(TealParserLONGSTRING-53))|(1<<(TealParserINT-53))|(1<<(TealParserHEX-53))|(1<<(TealParserFLOAT-53))|(1<<(TealParserHEX_FLOAT-53)))) != 0) {
		{
			p.SetState(382)
			p.Explist()
		}

	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__0 {
		{
			p.SetState(385)
			p.Match(TealParserT__0)
		}

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
	p.RuleIndex = TealParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *TealParser) Label() (localctx ILabelContext) {
	this := p
	_ = this

	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TealParserRULE_label)

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
		p.SetState(388)
		p.Match(TealParserT__36)
	}
	{
		p.SetState(389)
		p.Match(TealParserNAME)
	}
	{
		p.SetState(390)
		p.Match(TealParserT__36)
	}

	return localctx
}

// IFuncnameContext is an interface to support dynamic dispatch.
type IFuncnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncnameContext differentiates from other interfaces.
	IsFuncnameContext()
}

type FuncnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncnameContext() *FuncnameContext {
	var p = new(FuncnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_funcname
	return p
}

func (*FuncnameContext) IsFuncnameContext() {}

func NewFuncnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncnameContext {
	var p = new(FuncnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_funcname

	return p
}

func (s *FuncnameContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncnameContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(TealParserNAME)
}

func (s *FuncnameContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(TealParserNAME, i)
}

func (s *FuncnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterFuncname(s)
	}
}

func (s *FuncnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitFuncname(s)
	}
}

func (p *TealParser) Funcname() (localctx IFuncnameContext) {
	this := p
	_ = this

	localctx = NewFuncnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TealParserRULE_funcname)
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
		p.SetState(392)
		p.Match(TealParserNAME)
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TealParserT__37 {
		{
			p.SetState(393)
			p.Match(TealParserT__37)
		}
		{
			p.SetState(394)
			p.Match(TealParserNAME)
		}

		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__18 {
		{
			p.SetState(400)
			p.Match(TealParserT__18)
		}
		{
			p.SetState(401)
			p.Match(TealParserNAME)
		}

	}

	return localctx
}

// IVarlistContext is an interface to support dynamic dispatch.
type IVarlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarlistContext differentiates from other interfaces.
	IsVarlistContext()
}

type VarlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarlistContext() *VarlistContext {
	var p = new(VarlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_varlist
	return p
}

func (*VarlistContext) IsVarlistContext() {}

func NewVarlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarlistContext {
	var p = new(VarlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_varlist

	return p
}

func (s *VarlistContext) GetParser() antlr.Parser { return s.parser }

func (s *VarlistContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *VarlistContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VarlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterVarlist(s)
	}
}

func (s *VarlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitVarlist(s)
	}
}

func (p *TealParser) Varlist() (localctx IVarlistContext) {
	this := p
	_ = this

	localctx = NewVarlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TealParserRULE_varlist)
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
		p.SetState(404)
		p.Variable()
	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TealParserT__14 {
		{
			p.SetState(405)
			p.Match(TealParserT__14)
		}
		{
			p.SetState(406)
			p.Variable()
		}

		p.SetState(411)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INamelistContext is an interface to support dynamic dispatch.
type INamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamelistContext differentiates from other interfaces.
	IsNamelistContext()
}

type NamelistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamelistContext() *NamelistContext {
	var p = new(NamelistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_namelist
	return p
}

func (*NamelistContext) IsNamelistContext() {}

func NewNamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamelistContext {
	var p = new(NamelistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_namelist

	return p
}

func (s *NamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *NamelistContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(TealParserNAME)
}

func (s *NamelistContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(TealParserNAME, i)
}

func (s *NamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterNamelist(s)
	}
}

func (s *NamelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitNamelist(s)
	}
}

func (p *TealParser) Namelist() (localctx INamelistContext) {
	this := p
	_ = this

	localctx = NewNamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TealParserRULE_namelist)

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
		p.SetState(412)
		p.Match(TealParserNAME)
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(413)
				p.Match(TealParserT__14)
			}
			{
				p.SetState(414)
				p.Match(TealParserNAME)
			}

		}
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

// IExplistContext is an interface to support dynamic dispatch.
type IExplistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExplistContext differentiates from other interfaces.
	IsExplistContext()
}

type ExplistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExplistContext() *ExplistContext {
	var p = new(ExplistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_explist
	return p
}

func (*ExplistContext) IsExplistContext() {}

func NewExplistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExplistContext {
	var p = new(ExplistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_explist

	return p
}

func (s *ExplistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExplistContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExplistContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExplistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExplistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExplistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterExplist(s)
	}
}

func (s *ExplistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitExplist(s)
	}
}

func (p *TealParser) Explist() (localctx IExplistContext) {
	this := p
	_ = this

	localctx = NewExplistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TealParserRULE_explist)
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
		p.SetState(420)
		p.exp(0)
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TealParserT__14 {
		{
			p.SetState(421)
			p.Match(TealParserT__14)
		}
		{
			p.SetState(422)
			p.exp(0)
		}

		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ExpContext) Str() IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *ExpContext) Functiondef() IFunctiondefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctiondefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctiondefContext)
}

func (s *ExpContext) Prefixexp() IPrefixexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixexpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixexpContext)
}

func (s *ExpContext) Tableconstructor() ITableconstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableconstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableconstructorContext)
}

func (s *ExpContext) OperatorUnary() IOperatorUnaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorUnaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorUnaryContext)
}

func (s *ExpContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *ExpContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *ExpContext) OperatorPower() IOperatorPowerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorPowerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorPowerContext)
}

func (s *ExpContext) OperatorMulDivMod() IOperatorMulDivModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorMulDivModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorMulDivModContext)
}

func (s *ExpContext) OperatorAddSub() IOperatorAddSubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorAddSubContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorAddSubContext)
}

func (s *ExpContext) OperatorStrcat() IOperatorStrcatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorStrcatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorStrcatContext)
}

func (s *ExpContext) OperatorComparison() IOperatorComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorComparisonContext)
}

func (s *ExpContext) OperatorAnd() IOperatorAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorAndContext)
}

func (s *ExpContext) OperatorOr() IOperatorOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorOrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorOrContext)
}

func (s *ExpContext) OperatorBitwise() IOperatorBitwiseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperatorBitwiseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperatorBitwiseContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *TealParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *TealParser) exp(_p int) (localctx IExpContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, TealParserRULE_exp, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(429)
			p.Match(TealParserT__27)
		}

	case 2:
		{
			p.SetState(430)
			p.Match(TealParserT__38)
		}

	case 3:
		{
			p.SetState(431)
			p.Match(TealParserT__39)
		}

	case 4:
		{
			p.SetState(432)
			p.Number()
		}

	case 5:
		{
			p.SetState(433)
			p.Str()
		}

	case 6:
		{
			p.SetState(434)
			p.Match(TealParserT__31)
		}

	case 7:
		{
			p.SetState(435)
			p.Functiondef()
		}

	case 8:
		{
			p.SetState(436)
			p.Prefixexp()
		}

	case 9:
		{
			p.SetState(437)
			p.Tableconstructor()
		}

	case 10:
		{
			p.SetState(438)
			p.OperatorUnary()
		}
		{
			p.SetState(439)
			p.exp(9)
		}

	case 11:
		{
			p.SetState(441)
			p.Match(TealParserNAME)
		}
		{
			p.SetState(442)
			p.Match(TealParserT__41)
		}
		{
			p.SetState(443)
			p.Typ()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(481)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TealParserRULE_exp)
				p.SetState(446)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(447)
					p.OperatorPower()
				}
				{
					p.SetState(448)
					p.exp(10)
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TealParserRULE_exp)
				p.SetState(450)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(451)
					p.OperatorMulDivMod()
				}
				{
					p.SetState(452)
					p.exp(9)
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TealParserRULE_exp)
				p.SetState(454)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(455)
					p.OperatorAddSub()
				}
				{
					p.SetState(456)
					p.exp(8)
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TealParserRULE_exp)
				p.SetState(458)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(459)
					p.OperatorStrcat()
				}
				{
					p.SetState(460)
					p.exp(6)
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TealParserRULE_exp)
				p.SetState(462)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(463)
					p.OperatorComparison()
				}
				{
					p.SetState(464)
					p.exp(6)
				}

			case 6:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TealParserRULE_exp)
				p.SetState(466)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(467)
					p.OperatorAnd()
				}
				{
					p.SetState(468)
					p.exp(4)
				}

			case 7:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TealParserRULE_exp)
				p.SetState(470)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(471)
					p.OperatorOr()
				}
				{
					p.SetState(472)
					p.exp(3)
				}

			case 8:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TealParserRULE_exp)
				p.SetState(474)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(475)
					p.OperatorBitwise()
				}
				{
					p.SetState(476)
					p.exp(2)
				}

			case 9:
				localctx = NewExpContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TealParserRULE_exp)
				p.SetState(478)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(479)
					p.Match(TealParserT__40)
				}
				{
					p.SetState(480)
					p.Typ()
				}

			}

		}
		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
	}

	return localctx
}

// IPrefixexpContext is an interface to support dynamic dispatch.
type IPrefixexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixexpContext differentiates from other interfaces.
	IsPrefixexpContext()
}

type PrefixexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixexpContext() *PrefixexpContext {
	var p = new(PrefixexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_prefixexp
	return p
}

func (*PrefixexpContext) IsPrefixexpContext() {}

func NewPrefixexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixexpContext {
	var p = new(PrefixexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_prefixexp

	return p
}

func (s *PrefixexpContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixexpContext) VarOrExp() IVarOrExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarOrExpContext)
}

func (s *PrefixexpContext) AllNameAndArgs() []INameAndArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem())
	var tst = make([]INameAndArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameAndArgsContext)
		}
	}

	return tst
}

func (s *PrefixexpContext) NameAndArgs(i int) INameAndArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *PrefixexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterPrefixexp(s)
	}
}

func (s *PrefixexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitPrefixexp(s)
	}
}

func (p *TealParser) Prefixexp() (localctx IPrefixexpContext) {
	this := p
	_ = this

	localctx = NewPrefixexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TealParserRULE_prefixexp)

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
		p.SetState(486)
		p.VarOrExp()
	}
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(487)
				p.NameAndArgs()
			}

		}
		p.SetState(492)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctioncallContext is an interface to support dynamic dispatch.
type IFunctioncallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctioncallContext differentiates from other interfaces.
	IsFunctioncallContext()
}

type FunctioncallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctioncallContext() *FunctioncallContext {
	var p = new(FunctioncallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_functioncall
	return p
}

func (*FunctioncallContext) IsFunctioncallContext() {}

func NewFunctioncallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctioncallContext {
	var p = new(FunctioncallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_functioncall

	return p
}

func (s *FunctioncallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctioncallContext) VarOrExp() IVarOrExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarOrExpContext)
}

func (s *FunctioncallContext) AllNameAndArgs() []INameAndArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem())
	var tst = make([]INameAndArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameAndArgsContext)
		}
	}

	return tst
}

func (s *FunctioncallContext) NameAndArgs(i int) INameAndArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *FunctioncallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctioncallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctioncallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterFunctioncall(s)
	}
}

func (s *FunctioncallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitFunctioncall(s)
	}
}

func (p *TealParser) Functioncall() (localctx IFunctioncallContext) {
	this := p
	_ = this

	localctx = NewFunctioncallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TealParserRULE_functioncall)

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
		p.SetState(493)
		p.VarOrExp()
	}
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(494)
				p.NameAndArgs()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
	}

	return localctx
}

// IVarOrExpContext is an interface to support dynamic dispatch.
type IVarOrExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarOrExpContext differentiates from other interfaces.
	IsVarOrExpContext()
}

type VarOrExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarOrExpContext() *VarOrExpContext {
	var p = new(VarOrExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_varOrExp
	return p
}

func (*VarOrExpContext) IsVarOrExpContext() {}

func NewVarOrExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarOrExpContext {
	var p = new(VarOrExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_varOrExp

	return p
}

func (s *VarOrExpContext) GetParser() antlr.Parser { return s.parser }

func (s *VarOrExpContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VarOrExpContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarOrExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarOrExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarOrExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterVarOrExp(s)
	}
}

func (s *VarOrExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitVarOrExp(s)
	}
}

func (p *TealParser) VarOrExp() (localctx IVarOrExpContext) {
	this := p
	_ = this

	localctx = NewVarOrExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, TealParserRULE_varOrExp)

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

	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(500)
			p.Match(TealParserT__22)
		}
		{
			p.SetState(501)
			p.exp(0)
		}
		{
			p.SetState(502)
			p.Match(TealParserT__23)
		}

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
	p.RuleIndex = TealParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *VariableContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VariableContext) AllVarSuffix() []IVarSuffixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarSuffixContext)(nil)).Elem())
	var tst = make([]IVarSuffixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarSuffixContext)
		}
	}

	return tst
}

func (s *VariableContext) VarSuffix(i int) IVarSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarSuffixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarSuffixContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *TealParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, TealParserRULE_variable)

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
	p.SetState(512)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TealParserNAME:
		{
			p.SetState(506)
			p.Match(TealParserNAME)
		}

	case TealParserT__22:
		{
			p.SetState(507)
			p.Match(TealParserT__22)
		}
		{
			p.SetState(508)
			p.exp(0)
		}
		{
			p.SetState(509)
			p.Match(TealParserT__23)
		}
		{
			p.SetState(510)
			p.VarSuffix()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(514)
				p.VarSuffix()
			}

		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}

	return localctx
}

// IVarSuffixContext is an interface to support dynamic dispatch.
type IVarSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarSuffixContext differentiates from other interfaces.
	IsVarSuffixContext()
}

type VarSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarSuffixContext() *VarSuffixContext {
	var p = new(VarSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_varSuffix
	return p
}

func (*VarSuffixContext) IsVarSuffixContext() {}

func NewVarSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarSuffixContext {
	var p = new(VarSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_varSuffix

	return p
}

func (s *VarSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *VarSuffixContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *VarSuffixContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *VarSuffixContext) AllNameAndArgs() []INameAndArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem())
	var tst = make([]INameAndArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameAndArgsContext)
		}
	}

	return tst
}

func (s *VarSuffixContext) NameAndArgs(i int) INameAndArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameAndArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameAndArgsContext)
}

func (s *VarSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterVarSuffix(s)
	}
}

func (s *VarSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitVarSuffix(s)
	}
}

func (p *TealParser) VarSuffix() (localctx IVarSuffixContext) {
	this := p
	_ = this

	localctx = NewVarSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, TealParserRULE_varSuffix)
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
	p.SetState(523)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TealParserT__18)|(1<<TealParserT__22)|(1<<TealParserT__29))) != 0) || (((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(TealParserNORMALSTRING-66))|(1<<(TealParserCHARSTRING-66))|(1<<(TealParserLONGSTRING-66)))) != 0) {
		{
			p.SetState(520)
			p.NameAndArgs()
		}

		p.SetState(525)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(532)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TealParserT__42:
		{
			p.SetState(526)
			p.Match(TealParserT__42)
		}
		{
			p.SetState(527)
			p.exp(0)
		}
		{
			p.SetState(528)
			p.Match(TealParserT__43)
		}

	case TealParserT__37:
		{
			p.SetState(530)
			p.Match(TealParserT__37)
		}
		{
			p.SetState(531)
			p.Match(TealParserNAME)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INameAndArgsContext is an interface to support dynamic dispatch.
type INameAndArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameAndArgsContext differentiates from other interfaces.
	IsNameAndArgsContext()
}

type NameAndArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameAndArgsContext() *NameAndArgsContext {
	var p = new(NameAndArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_nameAndArgs
	return p
}

func (*NameAndArgsContext) IsNameAndArgsContext() {}

func NewNameAndArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameAndArgsContext {
	var p = new(NameAndArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_nameAndArgs

	return p
}

func (s *NameAndArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *NameAndArgsContext) Args() IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *NameAndArgsContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *NameAndArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameAndArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameAndArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterNameAndArgs(s)
	}
}

func (s *NameAndArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitNameAndArgs(s)
	}
}

func (p *TealParser) NameAndArgs() (localctx INameAndArgsContext) {
	this := p
	_ = this

	localctx = NewNameAndArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, TealParserRULE_nameAndArgs)
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
	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__18 {
		{
			p.SetState(534)
			p.Match(TealParserT__18)
		}
		{
			p.SetState(535)
			p.Match(TealParserNAME)
		}

	}
	{
		p.SetState(538)
		p.Args()
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) Explist() IExplistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExplistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExplistContext)
}

func (s *ArgsContext) Tableconstructor() ITableconstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableconstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableconstructorContext)
}

func (s *ArgsContext) Str() IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *TealParser) Args() (localctx IArgsContext) {
	this := p
	_ = this

	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, TealParserRULE_args)
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

	p.SetState(547)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TealParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(540)
			p.Match(TealParserT__22)
		}
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(TealParserT__16-17))|(1<<(TealParserT__22-17))|(1<<(TealParserT__27-17))|(1<<(TealParserT__29-17))|(1<<(TealParserT__31-17))|(1<<(TealParserT__38-17))|(1<<(TealParserT__39-17)))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(TealParserT__52-53))|(1<<(TealParserT__58-53))|(1<<(TealParserT__61-53))|(1<<(TealParserT__62-53))|(1<<(TealParserNAME-53))|(1<<(TealParserNORMALSTRING-53))|(1<<(TealParserCHARSTRING-53))|(1<<(TealParserLONGSTRING-53))|(1<<(TealParserINT-53))|(1<<(TealParserHEX-53))|(1<<(TealParserFLOAT-53))|(1<<(TealParserHEX_FLOAT-53)))) != 0) {
			{
				p.SetState(541)
				p.Explist()
			}

		}
		{
			p.SetState(544)
			p.Match(TealParserT__23)
		}

	case TealParserT__29:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(545)
			p.Tableconstructor()
		}

	case TealParserNORMALSTRING, TealParserCHARSTRING, TealParserLONGSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(546)
			p.Str()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctiondefContext is an interface to support dynamic dispatch.
type IFunctiondefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctiondefContext differentiates from other interfaces.
	IsFunctiondefContext()
}

type FunctiondefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctiondefContext() *FunctiondefContext {
	var p = new(FunctiondefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_functiondef
	return p
}

func (*FunctiondefContext) IsFunctiondefContext() {}

func NewFunctiondefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctiondefContext {
	var p = new(FunctiondefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_functiondef

	return p
}

func (s *FunctiondefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctiondefContext) Funcbody() IFuncbodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncbodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncbodyContext)
}

func (s *FunctiondefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctiondefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctiondefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterFunctiondef(s)
	}
}

func (s *FunctiondefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitFunctiondef(s)
	}
}

func (p *TealParser) Functiondef() (localctx IFunctiondefContext) {
	this := p
	_ = this

	localctx = NewFunctiondefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, TealParserRULE_functiondef)

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
		p.SetState(549)
		p.Match(TealParserT__16)
	}
	{
		p.SetState(550)
		p.Funcbody()
	}

	return localctx
}

// IFuncbodyContext is an interface to support dynamic dispatch.
type IFuncbodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncbodyContext differentiates from other interfaces.
	IsFuncbodyContext()
}

type FuncbodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncbodyContext() *FuncbodyContext {
	var p = new(FuncbodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_funcbody
	return p
}

func (*FuncbodyContext) IsFuncbodyContext() {}

func NewFuncbodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncbodyContext {
	var p = new(FuncbodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_funcbody

	return p
}

func (s *FuncbodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncbodyContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncbodyContext) Typeargs() ITypeargsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeargsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeargsContext)
}

func (s *FuncbodyContext) Parlist() IParlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParlistContext)
}

func (s *FuncbodyContext) Retlist() IRetlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetlistContext)
}

func (s *FuncbodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncbodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncbodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterFuncbody(s)
	}
}

func (s *FuncbodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitFuncbody(s)
	}
}

func (p *TealParser) Funcbody() (localctx IFuncbodyContext) {
	this := p
	_ = this

	localctx = NewFuncbodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, TealParserRULE_funcbody)
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
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__20 {
		{
			p.SetState(552)
			p.Typeargs()
		}

	}
	{
		p.SetState(555)
		p.Match(TealParserT__22)
	}
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__31 || _la == TealParserNAME {
		{
			p.SetState(556)
			p.Parlist()
		}

	}
	{
		p.SetState(559)
		p.Match(TealParserT__23)
	}
	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__18 {
		{
			p.SetState(560)
			p.Match(TealParserT__18)
		}
		{
			p.SetState(561)
			p.Retlist()
		}

	}
	{
		p.SetState(564)
		p.Block()
	}
	{
		p.SetState(565)
		p.Match(TealParserT__5)
	}

	return localctx
}

// IParlistContext is an interface to support dynamic dispatch.
type IParlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParlistContext differentiates from other interfaces.
	IsParlistContext()
}

type ParlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParlistContext() *ParlistContext {
	var p = new(ParlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_parlist
	return p
}

func (*ParlistContext) IsParlistContext() {}

func NewParlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParlistContext {
	var p = new(ParlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_parlist

	return p
}

func (s *ParlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParlistContext) Namelist() INamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *ParlistContext) Parnamelist() IParnamelistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParnamelistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParnamelistContext)
}

func (s *ParlistContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *ParlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterParlist(s)
	}
}

func (s *ParlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitParlist(s)
	}
}

func (p *TealParser) Parlist() (localctx IParlistContext) {
	this := p
	_ = this

	localctx = NewParlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, TealParserRULE_parlist)
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

	p.SetState(587)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(567)
			p.Namelist()
		}
		p.SetState(570)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__14 {
			{
				p.SetState(568)
				p.Match(TealParserT__14)
			}
			{
				p.SetState(569)
				p.Match(TealParserT__31)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(572)
			p.Match(TealParserT__31)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(573)
			p.Parnamelist()
		}
		p.SetState(580)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__14 {
			{
				p.SetState(574)
				p.Match(TealParserT__14)
			}
			{
				p.SetState(575)
				p.Match(TealParserT__31)
			}
			p.SetState(578)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == TealParserT__18 {
				{
					p.SetState(576)
					p.Match(TealParserT__18)
				}
				{
					p.SetState(577)
					p.Typ()
				}

			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(582)
			p.Match(TealParserT__31)
		}
		p.SetState(585)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__18 {
			{
				p.SetState(583)
				p.Match(TealParserT__18)
			}
			{
				p.SetState(584)
				p.Typ()
			}

		}

	}

	return localctx
}

// ITableconstructorContext is an interface to support dynamic dispatch.
type ITableconstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableconstructorContext differentiates from other interfaces.
	IsTableconstructorContext()
}

type TableconstructorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableconstructorContext() *TableconstructorContext {
	var p = new(TableconstructorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_tableconstructor
	return p
}

func (*TableconstructorContext) IsTableconstructorContext() {}

func NewTableconstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableconstructorContext {
	var p = new(TableconstructorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_tableconstructor

	return p
}

func (s *TableconstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *TableconstructorContext) Fieldlist() IFieldlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldlistContext)
}

func (s *TableconstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableconstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableconstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterTableconstructor(s)
	}
}

func (s *TableconstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitTableconstructor(s)
	}
}

func (p *TealParser) Tableconstructor() (localctx ITableconstructorContext) {
	this := p
	_ = this

	localctx = NewTableconstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, TealParserRULE_tableconstructor)
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
		p.SetState(589)
		p.Match(TealParserT__29)
	}
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(TealParserT__16-17))|(1<<(TealParserT__22-17))|(1<<(TealParserT__27-17))|(1<<(TealParserT__29-17))|(1<<(TealParserT__31-17))|(1<<(TealParserT__38-17))|(1<<(TealParserT__39-17))|(1<<(TealParserT__42-17)))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(TealParserT__52-53))|(1<<(TealParserT__58-53))|(1<<(TealParserT__61-53))|(1<<(TealParserT__62-53))|(1<<(TealParserNAME-53))|(1<<(TealParserNORMALSTRING-53))|(1<<(TealParserCHARSTRING-53))|(1<<(TealParserLONGSTRING-53))|(1<<(TealParserINT-53))|(1<<(TealParserHEX-53))|(1<<(TealParserFLOAT-53))|(1<<(TealParserHEX_FLOAT-53)))) != 0) {
		{
			p.SetState(590)
			p.Fieldlist()
		}

	}
	{
		p.SetState(593)
		p.Match(TealParserT__30)
	}

	return localctx
}

// IFieldlistContext is an interface to support dynamic dispatch.
type IFieldlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldlistContext differentiates from other interfaces.
	IsFieldlistContext()
}

type FieldlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldlistContext() *FieldlistContext {
	var p = new(FieldlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_fieldlist
	return p
}

func (*FieldlistContext) IsFieldlistContext() {}

func NewFieldlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldlistContext {
	var p = new(FieldlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_fieldlist

	return p
}

func (s *FieldlistContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldlistContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *FieldlistContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldlistContext) AllFieldsep() []IFieldsepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldsepContext)(nil)).Elem())
	var tst = make([]IFieldsepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldsepContext)
		}
	}

	return tst
}

func (s *FieldlistContext) Fieldsep(i int) IFieldsepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldsepContext)
}

func (s *FieldlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterFieldlist(s)
	}
}

func (s *FieldlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitFieldlist(s)
	}
}

func (p *TealParser) Fieldlist() (localctx IFieldlistContext) {
	this := p
	_ = this

	localctx = NewFieldlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, TealParserRULE_fieldlist)
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
		p.SetState(595)
		p.Field()
	}
	p.SetState(601)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(596)
				p.Fieldsep()
			}
			{
				p.SetState(597)
				p.Field()
			}

		}
		p.SetState(603)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())
	}
	p.SetState(605)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TealParserT__0 || _la == TealParserT__14 {
		{
			p.SetState(604)
			p.Fieldsep()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) CopyFrom(ctx *FieldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignFieldContext struct {
	*FieldContext
}

func NewAssignFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignFieldContext {
	var p = new(AssignFieldContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *AssignFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignFieldContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *AssignFieldContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AssignFieldContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *AssignFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterAssignField(s)
	}
}

func (s *AssignFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitAssignField(s)
	}
}

type AssignNewTypeFieldContext struct {
	*FieldContext
}

func NewAssignNewTypeFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignNewTypeFieldContext {
	var p = new(AssignNewTypeFieldContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *AssignNewTypeFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignNewTypeFieldContext) NAME() antlr.TerminalNode {
	return s.GetToken(TealParserNAME, 0)
}

func (s *AssignNewTypeFieldContext) Newtype() INewtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewtypeContext)
}

func (s *AssignNewTypeFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterAssignNewTypeField(s)
	}
}

func (s *AssignNewTypeFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitAssignNewTypeField(s)
	}
}

type ExprFieldContext struct {
	*FieldContext
}

func NewExprFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFieldContext {
	var p = new(ExprFieldContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *ExprFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFieldContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExprFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterExprField(s)
	}
}

func (s *ExprFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitExprField(s)
	}
}

type BracketAssginFieldContext struct {
	*FieldContext
}

func NewBracketAssginFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketAssginFieldContext {
	var p = new(BracketAssginFieldContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *BracketAssginFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketAssginFieldContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *BracketAssginFieldContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *BracketAssginFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterBracketAssginField(s)
	}
}

func (s *BracketAssginFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitBracketAssginField(s)
	}
}

func (p *TealParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, TealParserRULE_field)
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

	p.SetState(624)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBracketAssginFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(607)
			p.Match(TealParserT__42)
		}
		{
			p.SetState(608)
			p.exp(0)
		}
		{
			p.SetState(609)
			p.Match(TealParserT__43)
		}
		{
			p.SetState(610)
			p.Match(TealParserT__1)
		}
		{
			p.SetState(611)
			p.exp(0)
		}

	case 2:
		localctx = NewAssignFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(613)
			p.Match(TealParserNAME)
		}
		p.SetState(616)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TealParserT__18 {
			{
				p.SetState(614)
				p.Match(TealParserT__18)
			}
			{
				p.SetState(615)
				p.Typ()
			}

		}
		{
			p.SetState(618)
			p.Match(TealParserT__1)
		}
		{
			p.SetState(619)
			p.exp(0)
		}

	case 3:
		localctx = NewAssignNewTypeFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(620)
			p.Match(TealParserNAME)
		}
		{
			p.SetState(621)
			p.Match(TealParserT__1)
		}
		{
			p.SetState(622)
			p.Newtype()
		}

	case 4:
		localctx = NewExprFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(623)
			p.exp(0)
		}

	}

	return localctx
}

// IFieldsepContext is an interface to support dynamic dispatch.
type IFieldsepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsepContext differentiates from other interfaces.
	IsFieldsepContext()
}

type FieldsepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsepContext() *FieldsepContext {
	var p = new(FieldsepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_fieldsep
	return p
}

func (*FieldsepContext) IsFieldsepContext() {}

func NewFieldsepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsepContext {
	var p = new(FieldsepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_fieldsep

	return p
}

func (s *FieldsepContext) GetParser() antlr.Parser { return s.parser }
func (s *FieldsepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterFieldsep(s)
	}
}

func (s *FieldsepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitFieldsep(s)
	}
}

func (p *TealParser) Fieldsep() (localctx IFieldsepContext) {
	this := p
	_ = this

	localctx = NewFieldsepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, TealParserRULE_fieldsep)
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
		p.SetState(626)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TealParserT__0 || _la == TealParserT__14) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorOrContext is an interface to support dynamic dispatch.
type IOperatorOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorOrContext differentiates from other interfaces.
	IsOperatorOrContext()
}

type OperatorOrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorOrContext() *OperatorOrContext {
	var p = new(OperatorOrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_operatorOr
	return p
}

func (*OperatorOrContext) IsOperatorOrContext() {}

func NewOperatorOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorOrContext {
	var p = new(OperatorOrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_operatorOr

	return p
}

func (s *OperatorOrContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterOperatorOr(s)
	}
}

func (s *OperatorOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitOperatorOr(s)
	}
}

func (p *TealParser) OperatorOr() (localctx IOperatorOrContext) {
	this := p
	_ = this

	localctx = NewOperatorOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, TealParserRULE_operatorOr)

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
		p.SetState(628)
		p.Match(TealParserT__44)
	}

	return localctx
}

// IOperatorAndContext is an interface to support dynamic dispatch.
type IOperatorAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorAndContext differentiates from other interfaces.
	IsOperatorAndContext()
}

type OperatorAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorAndContext() *OperatorAndContext {
	var p = new(OperatorAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_operatorAnd
	return p
}

func (*OperatorAndContext) IsOperatorAndContext() {}

func NewOperatorAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorAndContext {
	var p = new(OperatorAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_operatorAnd

	return p
}

func (s *OperatorAndContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterOperatorAnd(s)
	}
}

func (s *OperatorAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitOperatorAnd(s)
	}
}

func (p *TealParser) OperatorAnd() (localctx IOperatorAndContext) {
	this := p
	_ = this

	localctx = NewOperatorAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, TealParserRULE_operatorAnd)

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
		p.SetState(630)
		p.Match(TealParserT__45)
	}

	return localctx
}

// IOperatorComparisonContext is an interface to support dynamic dispatch.
type IOperatorComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorComparisonContext differentiates from other interfaces.
	IsOperatorComparisonContext()
}

type OperatorComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorComparisonContext() *OperatorComparisonContext {
	var p = new(OperatorComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_operatorComparison
	return p
}

func (*OperatorComparisonContext) IsOperatorComparisonContext() {}

func NewOperatorComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorComparisonContext {
	var p = new(OperatorComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_operatorComparison

	return p
}

func (s *OperatorComparisonContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterOperatorComparison(s)
	}
}

func (s *OperatorComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitOperatorComparison(s)
	}
}

func (p *TealParser) OperatorComparison() (localctx IOperatorComparisonContext) {
	this := p
	_ = this

	localctx = NewOperatorComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, TealParserRULE_operatorComparison)
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
		p.SetState(632)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(TealParserT__20-21))|(1<<(TealParserT__21-21))|(1<<(TealParserT__46-21))|(1<<(TealParserT__47-21))|(1<<(TealParserT__48-21))|(1<<(TealParserT__49-21)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorStrcatContext is an interface to support dynamic dispatch.
type IOperatorStrcatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorStrcatContext differentiates from other interfaces.
	IsOperatorStrcatContext()
}

type OperatorStrcatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorStrcatContext() *OperatorStrcatContext {
	var p = new(OperatorStrcatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_operatorStrcat
	return p
}

func (*OperatorStrcatContext) IsOperatorStrcatContext() {}

func NewOperatorStrcatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorStrcatContext {
	var p = new(OperatorStrcatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_operatorStrcat

	return p
}

func (s *OperatorStrcatContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorStrcatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorStrcatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorStrcatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterOperatorStrcat(s)
	}
}

func (s *OperatorStrcatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitOperatorStrcat(s)
	}
}

func (p *TealParser) OperatorStrcat() (localctx IOperatorStrcatContext) {
	this := p
	_ = this

	localctx = NewOperatorStrcatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, TealParserRULE_operatorStrcat)

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
		p.SetState(634)
		p.Match(TealParserT__50)
	}

	return localctx
}

// IOperatorAddSubContext is an interface to support dynamic dispatch.
type IOperatorAddSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorAddSubContext differentiates from other interfaces.
	IsOperatorAddSubContext()
}

type OperatorAddSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorAddSubContext() *OperatorAddSubContext {
	var p = new(OperatorAddSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_operatorAddSub
	return p
}

func (*OperatorAddSubContext) IsOperatorAddSubContext() {}

func NewOperatorAddSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorAddSubContext {
	var p = new(OperatorAddSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_operatorAddSub

	return p
}

func (s *OperatorAddSubContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAddSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterOperatorAddSub(s)
	}
}

func (s *OperatorAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitOperatorAddSub(s)
	}
}

func (p *TealParser) OperatorAddSub() (localctx IOperatorAddSubContext) {
	this := p
	_ = this

	localctx = NewOperatorAddSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, TealParserRULE_operatorAddSub)
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
		p.SetState(636)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TealParserT__51 || _la == TealParserT__52) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorMulDivModContext is an interface to support dynamic dispatch.
type IOperatorMulDivModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorMulDivModContext differentiates from other interfaces.
	IsOperatorMulDivModContext()
}

type OperatorMulDivModContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorMulDivModContext() *OperatorMulDivModContext {
	var p = new(OperatorMulDivModContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_operatorMulDivMod
	return p
}

func (*OperatorMulDivModContext) IsOperatorMulDivModContext() {}

func NewOperatorMulDivModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorMulDivModContext {
	var p = new(OperatorMulDivModContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_operatorMulDivMod

	return p
}

func (s *OperatorMulDivModContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorMulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorMulDivModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorMulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterOperatorMulDivMod(s)
	}
}

func (s *OperatorMulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitOperatorMulDivMod(s)
	}
}

func (p *TealParser) OperatorMulDivMod() (localctx IOperatorMulDivModContext) {
	this := p
	_ = this

	localctx = NewOperatorMulDivModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, TealParserRULE_operatorMulDivMod)
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
		p.SetState(638)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(TealParserT__53-54))|(1<<(TealParserT__54-54))|(1<<(TealParserT__55-54))|(1<<(TealParserT__56-54)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorBitwiseContext is an interface to support dynamic dispatch.
type IOperatorBitwiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorBitwiseContext differentiates from other interfaces.
	IsOperatorBitwiseContext()
}

type OperatorBitwiseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorBitwiseContext() *OperatorBitwiseContext {
	var p = new(OperatorBitwiseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_operatorBitwise
	return p
}

func (*OperatorBitwiseContext) IsOperatorBitwiseContext() {}

func NewOperatorBitwiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorBitwiseContext {
	var p = new(OperatorBitwiseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_operatorBitwise

	return p
}

func (s *OperatorBitwiseContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorBitwiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorBitwiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorBitwiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterOperatorBitwise(s)
	}
}

func (s *OperatorBitwiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitOperatorBitwise(s)
	}
}

func (p *TealParser) OperatorBitwise() (localctx IOperatorBitwiseContext) {
	this := p
	_ = this

	localctx = NewOperatorBitwiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, TealParserRULE_operatorBitwise)
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
		p.SetState(640)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TealParserT__24 || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(TealParserT__57-58))|(1<<(TealParserT__58-58))|(1<<(TealParserT__59-58))|(1<<(TealParserT__60-58)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorUnaryContext is an interface to support dynamic dispatch.
type IOperatorUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorUnaryContext differentiates from other interfaces.
	IsOperatorUnaryContext()
}

type OperatorUnaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorUnaryContext() *OperatorUnaryContext {
	var p = new(OperatorUnaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_operatorUnary
	return p
}

func (*OperatorUnaryContext) IsOperatorUnaryContext() {}

func NewOperatorUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorUnaryContext {
	var p = new(OperatorUnaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_operatorUnary

	return p
}

func (s *OperatorUnaryContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorUnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterOperatorUnary(s)
	}
}

func (s *OperatorUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitOperatorUnary(s)
	}
}

func (p *TealParser) OperatorUnary() (localctx IOperatorUnaryContext) {
	this := p
	_ = this

	localctx = NewOperatorUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, TealParserRULE_operatorUnary)
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
		p.SetState(642)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(TealParserT__52-53))|(1<<(TealParserT__58-53))|(1<<(TealParserT__61-53))|(1<<(TealParserT__62-53)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorPowerContext is an interface to support dynamic dispatch.
type IOperatorPowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorPowerContext differentiates from other interfaces.
	IsOperatorPowerContext()
}

type OperatorPowerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorPowerContext() *OperatorPowerContext {
	var p = new(OperatorPowerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_operatorPower
	return p
}

func (*OperatorPowerContext) IsOperatorPowerContext() {}

func NewOperatorPowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorPowerContext {
	var p = new(OperatorPowerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_operatorPower

	return p
}

func (s *OperatorPowerContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorPowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorPowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorPowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterOperatorPower(s)
	}
}

func (s *OperatorPowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitOperatorPower(s)
	}
}

func (p *TealParser) OperatorPower() (localctx IOperatorPowerContext) {
	this := p
	_ = this

	localctx = NewOperatorPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, TealParserRULE_operatorPower)

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
		p.SetState(644)
		p.Match(TealParserT__63)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) INT() antlr.TerminalNode {
	return s.GetToken(TealParserINT, 0)
}

func (s *NumberContext) HEX() antlr.TerminalNode {
	return s.GetToken(TealParserHEX, 0)
}

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(TealParserFLOAT, 0)
}

func (s *NumberContext) HEX_FLOAT() antlr.TerminalNode {
	return s.GetToken(TealParserHEX_FLOAT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *TealParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, TealParserRULE_number)
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
		p.SetState(646)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(TealParserINT-69))|(1<<(TealParserHEX-69))|(1<<(TealParserFLOAT-69))|(1<<(TealParserHEX_FLOAT-69)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TealParserRULE_str
	return p
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TealParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) NORMALSTRING() antlr.TerminalNode {
	return s.GetToken(TealParserNORMALSTRING, 0)
}

func (s *StrContext) CHARSTRING() antlr.TerminalNode {
	return s.GetToken(TealParserCHARSTRING, 0)
}

func (s *StrContext) LONGSTRING() antlr.TerminalNode {
	return s.GetToken(TealParserLONGSTRING, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TealListener); ok {
		listenerT.ExitStr(s)
	}
}

func (p *TealParser) Str() (localctx IStrContext) {
	this := p
	_ = this

	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, TealParserRULE_str)
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
		p.SetState(648)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(TealParserNORMALSTRING-66))|(1<<(TealParserCHARSTRING-66))|(1<<(TealParserLONGSTRING-66)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *TealParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TealParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
