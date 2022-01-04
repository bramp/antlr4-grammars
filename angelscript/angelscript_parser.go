// Code generated from angelscript.g4 by ANTLR 4.9.3. DO NOT EDIT.

package angelscript // angelscript
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 75, 715,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 105, 10, 2, 13, 2, 14, 2, 106, 3, 3, 7,
	3, 110, 10, 3, 12, 3, 14, 3, 113, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 7, 3, 122, 10, 3, 12, 3, 14, 3, 125, 11, 3, 5, 3, 127, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 134, 10, 3, 12, 3, 14, 3, 137, 11,
	3, 3, 3, 5, 3, 140, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 6, 7, 6, 154, 10, 6, 12, 6, 14, 6, 157, 11, 6,
	3, 6, 5, 6, 160, 10, 6, 3, 6, 3, 6, 5, 6, 164, 10, 6, 3, 6, 5, 6, 167,
	10, 6, 5, 6, 169, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 174, 10, 6, 3, 6, 5, 6,
	177, 10, 6, 3, 6, 3, 6, 5, 6, 181, 10, 6, 3, 7, 7, 7, 184, 10, 7, 12, 7,
	14, 7, 187, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 196,
	10, 7, 12, 7, 14, 7, 199, 11, 7, 5, 7, 201, 10, 7, 3, 7, 3, 7, 3, 7, 7,
	7, 206, 10, 7, 12, 7, 14, 7, 209, 11, 7, 3, 7, 5, 7, 212, 10, 7, 3, 8,
	5, 8, 215, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 222, 10, 8, 3, 8,
	5, 8, 225, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 232, 10, 8, 3, 8,
	5, 8, 235, 10, 8, 7, 8, 237, 10, 8, 12, 8, 14, 8, 240, 11, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 5, 9, 247, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 10, 7, 10, 257, 10, 10, 12, 10, 14, 10, 260, 11, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 269, 10, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 5, 10, 275, 10, 10, 7, 10, 277, 10, 10, 12, 10, 14, 10,
	280, 11, 10, 3, 10, 5, 10, 283, 10, 10, 3, 11, 7, 11, 286, 10, 11, 12,
	11, 14, 11, 289, 11, 11, 3, 11, 3, 11, 3, 11, 5, 11, 294, 10, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 5, 12, 301, 10, 12, 3, 12, 3, 12, 5, 12, 305,
	10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 311, 10, 12, 3, 12, 3, 12, 3,
	12, 5, 12, 316, 10, 12, 7, 12, 318, 10, 12, 12, 12, 14, 12, 321, 11, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 330, 10, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 335, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15,
	7, 15, 342, 10, 15, 12, 15, 14, 15, 345, 11, 15, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 5, 16, 354, 10, 16, 3, 16, 3, 16, 5, 16, 358,
	10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 364, 10, 16, 3, 16, 3, 16, 5,
	16, 368, 10, 16, 7, 16, 370, 10, 16, 12, 16, 14, 16, 373, 11, 16, 5, 16,
	375, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 381, 10, 17, 5, 17, 383,
	10, 17, 3, 18, 5, 18, 386, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 7, 18, 394, 10, 18, 12, 18, 14, 18, 397, 11, 18, 3, 18, 3, 18, 7, 18,
	401, 10, 18, 12, 18, 14, 18, 404, 11, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 410, 10, 18, 7, 18, 412, 10, 18, 12, 18, 14, 18, 415, 11, 18, 3, 19,
	3, 19, 3, 19, 5, 19, 420, 10, 19, 3, 19, 3, 19, 3, 19, 5, 19, 425, 10,
	19, 7, 19, 427, 10, 19, 12, 19, 14, 19, 430, 11, 19, 3, 19, 3, 19, 3, 20,
	5, 20, 435, 10, 20, 3, 20, 3, 20, 7, 20, 439, 10, 20, 12, 20, 14, 20, 442,
	11, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 449, 10, 20, 12, 20,
	14, 20, 452, 11, 20, 3, 20, 3, 20, 5, 20, 456, 10, 20, 3, 20, 5, 20, 459,
	10, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 474, 10, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 7, 23, 482, 10, 23, 12, 23, 14, 23, 485, 11, 23, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 496, 10,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 502, 10, 25, 12, 25, 14, 25, 505,
	11, 25, 5, 25, 507, 10, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 533, 10, 28, 3,
	29, 3, 29, 3, 29, 3, 30, 5, 30, 539, 10, 30, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 5, 32, 550, 10, 32, 3, 32, 3, 32, 3,
	33, 3, 33, 3, 33, 5, 33, 557, 10, 33, 3, 33, 3, 33, 7, 33, 561, 10, 33,
	12, 33, 14, 33, 564, 11, 33, 3, 34, 3, 34, 3, 34, 3, 34, 7, 34, 570, 10,
	34, 12, 34, 14, 34, 573, 11, 34, 3, 35, 3, 35, 3, 35, 5, 35, 578, 10, 35,
	3, 35, 3, 35, 7, 35, 582, 10, 35, 12, 35, 14, 35, 585, 11, 35, 3, 35, 3,
	35, 7, 35, 589, 10, 35, 12, 35, 14, 35, 592, 11, 35, 5, 35, 594, 10, 35,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 5, 36, 607, 10, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 5, 38,
	615, 10, 38, 3, 38, 3, 38, 3, 38, 5, 38, 620, 10, 38, 3, 38, 3, 38, 3,
	38, 5, 38, 625, 10, 38, 3, 38, 3, 38, 7, 38, 629, 10, 38, 12, 38, 14, 38,
	632, 11, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 639, 10, 38, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 40, 5, 40, 654, 10, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 5,
	40, 661, 10, 40, 3, 40, 7, 40, 664, 10, 40, 12, 40, 14, 40, 667, 11, 40,
	5, 40, 669, 10, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 5, 43, 684, 10, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 5, 43, 690, 10, 43, 3, 43, 7, 43, 693, 10, 43, 12, 43, 14,
	43, 696, 11, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 5, 44, 703, 10, 44,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45, 711, 10, 45, 3, 46, 3,
	46, 3, 46, 2, 2, 47, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
	66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 2, 9, 3, 2, 4, 7, 4,
	2, 4, 4, 7, 7, 3, 2, 15, 16, 3, 2, 26, 27, 3, 2, 31, 33, 5, 2, 40, 41,
	65, 65, 70, 70, 3, 2, 60, 63, 2, 801, 2, 104, 3, 2, 2, 2, 4, 111, 3, 2,
	2, 2, 6, 141, 3, 2, 2, 2, 8, 146, 3, 2, 2, 2, 10, 155, 3, 2, 2, 2, 12,
	185, 3, 2, 2, 2, 14, 214, 3, 2, 2, 2, 16, 243, 3, 2, 2, 2, 18, 258, 3,
	2, 2, 2, 20, 287, 3, 2, 2, 2, 22, 300, 3, 2, 2, 2, 24, 324, 3, 2, 2, 2,
	26, 327, 3, 2, 2, 2, 28, 338, 3, 2, 2, 2, 30, 348, 3, 2, 2, 2, 32, 382,
	3, 2, 2, 2, 34, 385, 3, 2, 2, 2, 36, 416, 3, 2, 2, 2, 38, 434, 3, 2, 2,
	2, 40, 460, 3, 2, 2, 2, 42, 473, 3, 2, 2, 2, 44, 475, 3, 2, 2, 2, 46, 488,
	3, 2, 2, 2, 48, 491, 3, 2, 2, 2, 50, 511, 3, 2, 2, 2, 52, 517, 3, 2, 2,
	2, 54, 525, 3, 2, 2, 2, 56, 534, 3, 2, 2, 2, 58, 538, 3, 2, 2, 2, 60, 542,
	3, 2, 2, 2, 62, 547, 3, 2, 2, 2, 64, 556, 3, 2, 2, 2, 66, 565, 3, 2, 2,
	2, 68, 593, 3, 2, 2, 2, 70, 606, 3, 2, 2, 2, 72, 608, 3, 2, 2, 2, 74, 638,
	3, 2, 2, 2, 76, 640, 3, 2, 2, 2, 78, 648, 3, 2, 2, 2, 80, 673, 3, 2, 2,
	2, 82, 677, 3, 2, 2, 2, 84, 680, 3, 2, 2, 2, 86, 699, 3, 2, 2, 2, 88, 704,
	3, 2, 2, 2, 90, 712, 3, 2, 2, 2, 92, 105, 5, 16, 9, 2, 93, 105, 5, 18,
	10, 2, 94, 105, 5, 6, 4, 2, 95, 105, 5, 4, 3, 2, 96, 105, 5, 24, 13, 2,
	97, 105, 5, 12, 7, 2, 98, 105, 5, 20, 11, 2, 99, 105, 5, 22, 12, 2, 100,
	105, 5, 14, 8, 2, 101, 105, 5, 10, 6, 2, 102, 105, 5, 8, 5, 2, 103, 105,
	7, 3, 2, 2, 104, 92, 3, 2, 2, 2, 104, 93, 3, 2, 2, 2, 104, 94, 3, 2, 2,
	2, 104, 95, 3, 2, 2, 2, 104, 96, 3, 2, 2, 2, 104, 97, 3, 2, 2, 2, 104,
	98, 3, 2, 2, 2, 104, 99, 3, 2, 2, 2, 104, 100, 3, 2, 2, 2, 104, 101, 3,
	2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2,
	2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 3, 3, 2, 2, 2, 108,
	110, 9, 2, 2, 2, 109, 108, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109,
	3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 114, 3, 2, 2, 2, 113, 111, 3, 2,
	2, 2, 114, 115, 7, 8, 2, 2, 115, 139, 7, 70, 2, 2, 116, 140, 7, 3, 2, 2,
	117, 118, 7, 9, 2, 2, 118, 123, 7, 70, 2, 2, 119, 120, 7, 10, 2, 2, 120,
	122, 7, 70, 2, 2, 121, 119, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2, 123, 121,
	3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 127, 3, 2, 2, 2, 125, 123, 3, 2,
	2, 2, 126, 117, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2,
	128, 135, 7, 11, 2, 2, 129, 134, 5, 22, 12, 2, 130, 134, 5, 10, 6, 2, 131,
	134, 5, 14, 8, 2, 132, 134, 5, 20, 11, 2, 133, 129, 3, 2, 2, 2, 133, 130,
	3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2, 2, 2, 134, 137, 3, 2,
	2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 138, 3, 2, 2, 2,
	137, 135, 3, 2, 2, 2, 138, 140, 7, 12, 2, 2, 139, 116, 3, 2, 2, 2, 139,
	126, 3, 2, 2, 2, 140, 5, 3, 2, 2, 2, 141, 142, 7, 13, 2, 2, 142, 143, 7,
	65, 2, 2, 143, 144, 7, 70, 2, 2, 144, 145, 7, 3, 2, 2, 145, 7, 3, 2, 2,
	2, 146, 147, 7, 14, 2, 2, 147, 148, 7, 70, 2, 2, 148, 149, 7, 11, 2, 2,
	149, 150, 5, 2, 2, 2, 150, 151, 7, 12, 2, 2, 151, 9, 3, 2, 2, 2, 152, 154,
	9, 3, 2, 2, 153, 152, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2,
	2, 2, 155, 156, 3, 2, 2, 2, 156, 159, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2,
	158, 160, 9, 4, 2, 2, 159, 158, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	168, 3, 2, 2, 2, 161, 163, 5, 34, 18, 2, 162, 164, 7, 17, 2, 2, 163, 162,
	3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 167, 7, 18,
	2, 2, 166, 161, 3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 167, 169, 3, 2, 2, 2,
	168, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170,
	171, 7, 70, 2, 2, 171, 173, 5, 30, 16, 2, 172, 174, 7, 19, 2, 2, 173, 172,
	3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 176, 3, 2, 2, 2, 175, 177, 7, 66,
	2, 2, 176, 175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2,
	178, 181, 7, 3, 2, 2, 179, 181, 5, 28, 15, 2, 180, 178, 3, 2, 2, 2, 180,
	179, 3, 2, 2, 2, 181, 11, 3, 2, 2, 2, 182, 184, 9, 3, 2, 2, 183, 182, 3,
	2, 2, 2, 184, 187, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2,
	2, 186, 188, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 188, 189, 7, 20, 2, 2, 189,
	211, 7, 70, 2, 2, 190, 212, 7, 3, 2, 2, 191, 192, 7, 9, 2, 2, 192, 197,
	7, 70, 2, 2, 193, 194, 7, 10, 2, 2, 194, 196, 7, 70, 2, 2, 195, 193, 3,
	2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2,
	2, 198, 201, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 191, 3, 2, 2, 2, 200,
	201, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 207, 7, 11, 2, 2, 203, 206,
	5, 22, 12, 2, 204, 206, 5, 26, 14, 2, 205, 203, 3, 2, 2, 2, 205, 204, 3,
	2, 2, 2, 206, 209, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2,
	2, 208, 210, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 210, 212, 7, 12, 2, 2, 211,
	190, 3, 2, 2, 2, 211, 200, 3, 2, 2, 2, 212, 13, 3, 2, 2, 2, 213, 215, 9,
	4, 2, 2, 214, 213, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216, 3, 2, 2,
	2, 216, 217, 5, 34, 18, 2, 217, 224, 7, 70, 2, 2, 218, 221, 7, 21, 2, 2,
	219, 222, 5, 36, 19, 2, 220, 222, 5, 66, 34, 2, 221, 219, 3, 2, 2, 2, 221,
	220, 3, 2, 2, 2, 222, 225, 3, 2, 2, 2, 223, 225, 5, 84, 43, 2, 224, 218,
	3, 2, 2, 2, 224, 223, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 238, 3, 2,
	2, 2, 226, 227, 7, 10, 2, 2, 227, 234, 7, 70, 2, 2, 228, 231, 7, 21, 2,
	2, 229, 232, 5, 36, 19, 2, 230, 232, 5, 66, 34, 2, 231, 229, 3, 2, 2, 2,
	231, 230, 3, 2, 2, 2, 232, 235, 3, 2, 2, 2, 233, 235, 5, 84, 43, 2, 234,
	228, 3, 2, 2, 2, 234, 233, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 237,
	3, 2, 2, 2, 236, 226, 3, 2, 2, 2, 237, 240, 3, 2, 2, 2, 238, 236, 3, 2,
	2, 2, 238, 239, 3, 2, 2, 2, 239, 241, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2,
	241, 242, 7, 3, 2, 2, 242, 15, 3, 2, 2, 2, 243, 244, 7, 22, 2, 2, 244,
	246, 5, 34, 18, 2, 245, 247, 7, 17, 2, 2, 246, 245, 3, 2, 2, 2, 246, 247,
	3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 249, 7, 70, 2, 2, 249, 250, 5, 30,
	16, 2, 250, 251, 7, 66, 2, 2, 251, 252, 7, 23, 2, 2, 252, 253, 7, 72, 2,
	2, 253, 254, 7, 3, 2, 2, 254, 17, 3, 2, 2, 2, 255, 257, 9, 3, 2, 2, 256,
	255, 3, 2, 2, 2, 257, 260, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 258, 259,
	3, 2, 2, 2, 259, 261, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 261, 262, 7, 24,
	2, 2, 262, 282, 7, 70, 2, 2, 263, 283, 7, 3, 2, 2, 264, 265, 7, 11, 2,
	2, 265, 268, 7, 70, 2, 2, 266, 267, 7, 21, 2, 2, 267, 269, 5, 66, 34, 2,
	268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 278, 3, 2, 2, 2, 270,
	271, 7, 10, 2, 2, 271, 274, 7, 70, 2, 2, 272, 273, 7, 21, 2, 2, 273, 275,
	5, 66, 34, 2, 274, 272, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 277, 3,
	2, 2, 2, 276, 270, 3, 2, 2, 2, 277, 280, 3, 2, 2, 2, 278, 276, 3, 2, 2,
	2, 278, 279, 3, 2, 2, 2, 279, 281, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 281,
	283, 7, 12, 2, 2, 282, 263, 3, 2, 2, 2, 282, 264, 3, 2, 2, 2, 283, 19,
	3, 2, 2, 2, 284, 286, 9, 3, 2, 2, 285, 284, 3, 2, 2, 2, 286, 289, 3, 2,
	2, 2, 287, 285, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 290, 3, 2, 2, 2,
	289, 287, 3, 2, 2, 2, 290, 291, 7, 25, 2, 2, 291, 293, 5, 34, 18, 2, 292,
	294, 7, 17, 2, 2, 293, 292, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294, 295,
	3, 2, 2, 2, 295, 296, 7, 70, 2, 2, 296, 297, 5, 30, 16, 2, 297, 298, 7,
	3, 2, 2, 298, 21, 3, 2, 2, 2, 299, 301, 9, 4, 2, 2, 300, 299, 3, 2, 2,
	2, 300, 301, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302, 304, 5, 34, 18, 2,
	303, 305, 7, 17, 2, 2, 304, 303, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305,
	306, 3, 2, 2, 2, 306, 307, 7, 70, 2, 2, 307, 319, 7, 11, 2, 2, 308, 310,
	9, 5, 2, 2, 309, 311, 7, 19, 2, 2, 310, 309, 3, 2, 2, 2, 310, 311, 3, 2,
	2, 2, 311, 312, 3, 2, 2, 2, 312, 315, 7, 66, 2, 2, 313, 316, 5, 28, 15,
	2, 314, 316, 7, 3, 2, 2, 315, 313, 3, 2, 2, 2, 315, 314, 3, 2, 2, 2, 316,
	318, 3, 2, 2, 2, 317, 308, 3, 2, 2, 2, 318, 321, 3, 2, 2, 2, 319, 317,
	3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 322, 3, 2, 2, 2, 321, 319, 3, 2,
	2, 2, 322, 323, 7, 12, 2, 2, 323, 23, 3, 2, 2, 2, 324, 325, 7, 28, 2, 2,
	325, 326, 5, 4, 3, 2, 326, 25, 3, 2, 2, 2, 327, 329, 5, 34, 18, 2, 328,
	330, 7, 17, 2, 2, 329, 328, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 331,
	3, 2, 2, 2, 331, 332, 7, 70, 2, 2, 332, 334, 5, 30, 16, 2, 333, 335, 7,
	19, 2, 2, 334, 333, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 336, 3, 2, 2,
	2, 336, 337, 7, 3, 2, 2, 337, 27, 3, 2, 2, 2, 338, 343, 7, 11, 2, 2, 339,
	342, 5, 14, 8, 2, 340, 342, 5, 42, 22, 2, 341, 339, 3, 2, 2, 2, 341, 340,
	3, 2, 2, 2, 342, 345, 3, 2, 2, 2, 343, 341, 3, 2, 2, 2, 343, 344, 3, 2,
	2, 2, 344, 346, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 346, 347, 7, 12, 2, 2,
	347, 29, 3, 2, 2, 2, 348, 374, 7, 29, 2, 2, 349, 375, 7, 68, 2, 2, 350,
	351, 5, 34, 18, 2, 351, 353, 5, 32, 17, 2, 352, 354, 7, 70, 2, 2, 353,
	352, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 357, 3, 2, 2, 2, 355, 356,
	7, 21, 2, 2, 356, 358, 5, 66, 34, 2, 357, 355, 3, 2, 2, 2, 357, 358, 3,
	2, 2, 2, 358, 371, 3, 2, 2, 2, 359, 360, 7, 10, 2, 2, 360, 361, 5, 34,
	18, 2, 361, 363, 5, 32, 17, 2, 362, 364, 7, 70, 2, 2, 363, 362, 3, 2, 2,
	2, 363, 364, 3, 2, 2, 2, 364, 367, 3, 2, 2, 2, 365, 366, 7, 21, 2, 2, 366,
	368, 5, 66, 34, 2, 367, 365, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 370,
	3, 2, 2, 2, 369, 359, 3, 2, 2, 2, 370, 373, 3, 2, 2, 2, 371, 369, 3, 2,
	2, 2, 371, 372, 3, 2, 2, 2, 372, 375, 3, 2, 2, 2, 373, 371, 3, 2, 2, 2,
	374, 349, 3, 2, 2, 2, 374, 350, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375,
	376, 3, 2, 2, 2, 376, 377, 7, 30, 2, 2, 377, 31, 3, 2, 2, 2, 378, 380,
	7, 17, 2, 2, 379, 381, 9, 6, 2, 2, 380, 379, 3, 2, 2, 2, 380, 381, 3, 2,
	2, 2, 381, 383, 3, 2, 2, 2, 382, 378, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2,
	383, 33, 3, 2, 2, 2, 384, 386, 7, 19, 2, 2, 385, 384, 3, 2, 2, 2, 385,
	386, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 388, 5, 38, 20, 2, 388, 402,
	5, 40, 21, 2, 389, 390, 7, 34, 2, 2, 390, 395, 5, 34, 18, 2, 391, 392,
	7, 10, 2, 2, 392, 394, 5, 34, 18, 2, 393, 391, 3, 2, 2, 2, 394, 397, 3,
	2, 2, 2, 395, 393, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 398, 3, 2, 2,
	2, 397, 395, 3, 2, 2, 2, 398, 399, 7, 35, 2, 2, 399, 401, 3, 2, 2, 2, 400,
	389, 3, 2, 2, 2, 401, 404, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2, 402, 403,
	3, 2, 2, 2, 403, 413, 3, 2, 2, 2, 404, 402, 3, 2, 2, 2, 405, 406, 7, 36,
	2, 2, 406, 412, 7, 37, 2, 2, 407, 409, 7, 38, 2, 2, 408, 410, 7, 19, 2,
	2, 409, 408, 3, 2, 2, 2, 409, 410, 3, 2, 2, 2, 410, 412, 3, 2, 2, 2, 411,
	405, 3, 2, 2, 2, 411, 407, 3, 2, 2, 2, 412, 415, 3, 2, 2, 2, 413, 411,
	3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 35, 3, 2, 2, 2, 415, 413, 3, 2,
	2, 2, 416, 419, 7, 11, 2, 2, 417, 420, 5, 86, 44, 2, 418, 420, 5, 36, 19,
	2, 419, 417, 3, 2, 2, 2, 419, 418, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420,
	428, 3, 2, 2, 2, 421, 424, 7, 10, 2, 2, 422, 425, 5, 86, 44, 2, 423, 425,
	5, 36, 19, 2, 424, 422, 3, 2, 2, 2, 424, 423, 3, 2, 2, 2, 424, 425, 3,
	2, 2, 2, 425, 427, 3, 2, 2, 2, 426, 421, 3, 2, 2, 2, 427, 430, 3, 2, 2,
	2, 428, 426, 3, 2, 2, 2, 428, 429, 3, 2, 2, 2, 429, 431, 3, 2, 2, 2, 430,
	428, 3, 2, 2, 2, 431, 432, 7, 12, 2, 2, 432, 37, 3, 2, 2, 2, 433, 435,
	7, 39, 2, 2, 434, 433, 3, 2, 2, 2, 434, 435, 3, 2, 2, 2, 435, 440, 3, 2,
	2, 2, 436, 437, 7, 70, 2, 2, 437, 439, 7, 39, 2, 2, 438, 436, 3, 2, 2,
	2, 439, 442, 3, 2, 2, 2, 440, 438, 3, 2, 2, 2, 440, 441, 3, 2, 2, 2, 441,
	458, 3, 2, 2, 2, 442, 440, 3, 2, 2, 2, 443, 455, 7, 70, 2, 2, 444, 445,
	7, 34, 2, 2, 445, 450, 5, 34, 18, 2, 446, 447, 7, 10, 2, 2, 447, 449, 5,
	34, 18, 2, 448, 446, 3, 2, 2, 2, 449, 452, 3, 2, 2, 2, 450, 448, 3, 2,
	2, 2, 450, 451, 3, 2, 2, 2, 451, 453, 3, 2, 2, 2, 452, 450, 3, 2, 2, 2,
	453, 454, 7, 35, 2, 2, 454, 456, 3, 2, 2, 2, 455, 444, 3, 2, 2, 2, 455,
	456, 3, 2, 2, 2, 456, 457, 3, 2, 2, 2, 457, 459, 7, 39, 2, 2, 458, 443,
	3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 39, 3, 2, 2, 2, 460, 461, 9, 7,
	2, 2, 461, 41, 3, 2, 2, 2, 462, 474, 5, 54, 28, 2, 463, 474, 5, 48, 25,
	2, 464, 474, 5, 50, 26, 2, 465, 474, 5, 62, 32, 2, 466, 474, 5, 28, 15,
	2, 467, 474, 5, 46, 24, 2, 468, 474, 5, 56, 29, 2, 469, 474, 5, 52, 27,
	2, 470, 474, 5, 44, 23, 2, 471, 474, 5, 58, 30, 2, 472, 474, 5, 60, 31,
	2, 473, 462, 3, 2, 2, 2, 473, 463, 3, 2, 2, 2, 473, 464, 3, 2, 2, 2, 473,
	465, 3, 2, 2, 2, 473, 466, 3, 2, 2, 2, 473, 467, 3, 2, 2, 2, 473, 468,
	3, 2, 2, 2, 473, 469, 3, 2, 2, 2, 473, 470, 3, 2, 2, 2, 473, 471, 3, 2,
	2, 2, 473, 472, 3, 2, 2, 2, 474, 43, 3, 2, 2, 2, 475, 476, 7, 42, 2, 2,
	476, 477, 7, 29, 2, 2, 477, 478, 5, 86, 44, 2, 478, 479, 7, 30, 2, 2, 479,
	483, 7, 11, 2, 2, 480, 482, 5, 64, 33, 2, 481, 480, 3, 2, 2, 2, 482, 485,
	3, 2, 2, 2, 483, 481, 3, 2, 2, 2, 483, 484, 3, 2, 2, 2, 484, 486, 3, 2,
	2, 2, 485, 483, 3, 2, 2, 2, 486, 487, 7, 12, 2, 2, 487, 45, 3, 2, 2, 2,
	488, 489, 7, 43, 2, 2, 489, 490, 7, 3, 2, 2, 490, 47, 3, 2, 2, 2, 491,
	492, 7, 44, 2, 2, 492, 495, 7, 29, 2, 2, 493, 496, 5, 14, 8, 2, 494, 496,
	5, 58, 30, 2, 495, 493, 3, 2, 2, 2, 495, 494, 3, 2, 2, 2, 496, 497, 3,
	2, 2, 2, 497, 506, 5, 58, 30, 2, 498, 503, 5, 86, 44, 2, 499, 500, 7, 10,
	2, 2, 500, 502, 5, 86, 44, 2, 501, 499, 3, 2, 2, 2, 502, 505, 3, 2, 2,
	2, 503, 501, 3, 2, 2, 2, 503, 504, 3, 2, 2, 2, 504, 507, 3, 2, 2, 2, 505,
	503, 3, 2, 2, 2, 506, 498, 3, 2, 2, 2, 506, 507, 3, 2, 2, 2, 507, 508,
	3, 2, 2, 2, 508, 509, 7, 30, 2, 2, 509, 510, 5, 42, 22, 2, 510, 49, 3,
	2, 2, 2, 511, 512, 7, 45, 2, 2, 512, 513, 7, 29, 2, 2, 513, 514, 5, 86,
	44, 2, 514, 515, 7, 30, 2, 2, 515, 516, 5, 42, 22, 2, 516, 51, 3, 2, 2,
	2, 517, 518, 7, 46, 2, 2, 518, 519, 5, 42, 22, 2, 519, 520, 7, 45, 2, 2,
	520, 521, 7, 29, 2, 2, 521, 522, 5, 86, 44, 2, 522, 523, 7, 30, 2, 2, 523,
	524, 7, 3, 2, 2, 524, 53, 3, 2, 2, 2, 525, 526, 7, 47, 2, 2, 526, 527,
	7, 29, 2, 2, 527, 528, 5, 86, 44, 2, 528, 529, 7, 30, 2, 2, 529, 532, 5,
	42, 22, 2, 530, 531, 7, 48, 2, 2, 531, 533, 5, 42, 22, 2, 532, 530, 3,
	2, 2, 2, 532, 533, 3, 2, 2, 2, 533, 55, 3, 2, 2, 2, 534, 535, 7, 49, 2,
	2, 535, 536, 7, 3, 2, 2, 536, 57, 3, 2, 2, 2, 537, 539, 5, 86, 44, 2, 538,
	537, 3, 2, 2, 2, 538, 539, 3, 2, 2, 2, 539, 540, 3, 2, 2, 2, 540, 541,
	7, 3, 2, 2, 541, 59, 3, 2, 2, 2, 542, 543, 7, 50, 2, 2, 543, 544, 5, 28,
	15, 2, 544, 545, 7, 51, 2, 2, 545, 546, 5, 28, 15, 2, 546, 61, 3, 2, 2,
	2, 547, 549, 7, 52, 2, 2, 548, 550, 5, 86, 44, 2, 549, 548, 3, 2, 2, 2,
	549, 550, 3, 2, 2, 2, 550, 551, 3, 2, 2, 2, 551, 552, 7, 3, 2, 2, 552,
	63, 3, 2, 2, 2, 553, 554, 7, 53, 2, 2, 554, 557, 5, 66, 34, 2, 555, 557,
	7, 54, 2, 2, 556, 553, 3, 2, 2, 2, 556, 555, 3, 2, 2, 2, 557, 558, 3, 2,
	2, 2, 558, 562, 7, 9, 2, 2, 559, 561, 5, 42, 22, 2, 560, 559, 3, 2, 2,
	2, 561, 564, 3, 2, 2, 2, 562, 560, 3, 2, 2, 2, 562, 563, 3, 2, 2, 2, 563,
	65, 3, 2, 2, 2, 564, 562, 3, 2, 2, 2, 565, 571, 5, 68, 35, 2, 566, 567,
	5, 90, 46, 2, 567, 568, 5, 68, 35, 2, 568, 570, 3, 2, 2, 2, 569, 566, 3,
	2, 2, 2, 570, 573, 3, 2, 2, 2, 571, 569, 3, 2, 2, 2, 571, 572, 3, 2, 2,
	2, 572, 67, 3, 2, 2, 2, 573, 571, 3, 2, 2, 2, 574, 575, 5, 34, 18, 2, 575,
	576, 7, 21, 2, 2, 576, 578, 3, 2, 2, 2, 577, 574, 3, 2, 2, 2, 577, 578,
	3, 2, 2, 2, 578, 579, 3, 2, 2, 2, 579, 594, 5, 36, 19, 2, 580, 582, 7,
	67, 2, 2, 581, 580, 3, 2, 2, 2, 582, 585, 3, 2, 2, 2, 583, 581, 3, 2, 2,
	2, 583, 584, 3, 2, 2, 2, 584, 586, 3, 2, 2, 2, 585, 583, 3, 2, 2, 2, 586,
	590, 5, 70, 36, 2, 587, 589, 5, 74, 38, 2, 588, 587, 3, 2, 2, 2, 589, 592,
	3, 2, 2, 2, 590, 588, 3, 2, 2, 2, 590, 591, 3, 2, 2, 2, 591, 594, 3, 2,
	2, 2, 592, 590, 3, 2, 2, 2, 593, 577, 3, 2, 2, 2, 593, 583, 3, 2, 2, 2,
	594, 69, 3, 2, 2, 2, 595, 607, 7, 68, 2, 2, 596, 607, 5, 72, 37, 2, 597,
	607, 5, 80, 41, 2, 598, 607, 5, 82, 42, 2, 599, 607, 5, 76, 39, 2, 600,
	607, 7, 69, 2, 2, 601, 602, 7, 29, 2, 2, 602, 603, 5, 86, 44, 2, 603, 604,
	7, 30, 2, 2, 604, 607, 3, 2, 2, 2, 605, 607, 5, 78, 40, 2, 606, 595, 3,
	2, 2, 2, 606, 596, 3, 2, 2, 2, 606, 597, 3, 2, 2, 2, 606, 598, 3, 2, 2,
	2, 606, 599, 3, 2, 2, 2, 606, 600, 3, 2, 2, 2, 606, 601, 3, 2, 2, 2, 606,
	605, 3, 2, 2, 2, 607, 71, 3, 2, 2, 2, 608, 609, 5, 34, 18, 2, 609, 610,
	5, 84, 43, 2, 610, 73, 3, 2, 2, 2, 611, 614, 7, 55, 2, 2, 612, 615, 5,
	80, 41, 2, 613, 615, 7, 70, 2, 2, 614, 612, 3, 2, 2, 2, 614, 613, 3, 2,
	2, 2, 615, 639, 3, 2, 2, 2, 616, 619, 7, 36, 2, 2, 617, 618, 7, 70, 2,
	2, 618, 620, 7, 9, 2, 2, 619, 617, 3, 2, 2, 2, 619, 620, 3, 2, 2, 2, 620,
	621, 3, 2, 2, 2, 621, 630, 5, 86, 44, 2, 622, 624, 7, 10, 2, 2, 623, 625,
	7, 70, 2, 2, 624, 623, 3, 2, 2, 2, 624, 625, 3, 2, 2, 2, 625, 626, 3, 2,
	2, 2, 626, 627, 7, 9, 2, 2, 627, 629, 5, 86, 44, 2, 628, 622, 3, 2, 2,
	2, 629, 632, 3, 2, 2, 2, 630, 628, 3, 2, 2, 2, 630, 631, 3, 2, 2, 2, 631,
	633, 3, 2, 2, 2, 632, 630, 3, 2, 2, 2, 633, 634, 7, 37, 2, 2, 634, 639,
	3, 2, 2, 2, 635, 639, 5, 84, 43, 2, 636, 639, 7, 56, 2, 2, 637, 639, 7,
	57, 2, 2, 638, 611, 3, 2, 2, 2, 638, 616, 3, 2, 2, 2, 638, 635, 3, 2, 2,
	2, 638, 636, 3, 2, 2, 2, 638, 637, 3, 2, 2, 2, 639, 75, 3, 2, 2, 2, 640,
	641, 7, 58, 2, 2, 641, 642, 7, 34, 2, 2, 642, 643, 5, 34, 18, 2, 643, 644,
	7, 35, 2, 2, 644, 645, 7, 29, 2, 2, 645, 646, 5, 86, 44, 2, 646, 647, 7,
	30, 2, 2, 647, 77, 3, 2, 2, 2, 648, 649, 7, 59, 2, 2, 649, 668, 7, 29,
	2, 2, 650, 651, 5, 34, 18, 2, 651, 652, 5, 32, 17, 2, 652, 654, 3, 2, 2,
	2, 653, 650, 3, 2, 2, 2, 653, 654, 3, 2, 2, 2, 654, 655, 3, 2, 2, 2, 655,
	665, 7, 70, 2, 2, 656, 660, 7, 10, 2, 2, 657, 658, 5, 34, 18, 2, 658, 659,
	5, 32, 17, 2, 659, 661, 3, 2, 2, 2, 660, 657, 3, 2, 2, 2, 660, 661, 3,
	2, 2, 2, 661, 662, 3, 2, 2, 2, 662, 664, 7, 70, 2, 2, 663, 656, 3, 2, 2,
	2, 664, 667, 3, 2, 2, 2, 665, 663, 3, 2, 2, 2, 665, 666, 3, 2, 2, 2, 666,
	669, 3, 2, 2, 2, 667, 665, 3, 2, 2, 2, 668, 653, 3, 2, 2, 2, 668, 669,
	3, 2, 2, 2, 669, 670, 3, 2, 2, 2, 670, 671, 7, 30, 2, 2, 671, 672, 5, 28,
	15, 2, 672, 79, 3, 2, 2, 2, 673, 674, 5, 38, 20, 2, 674, 675, 7, 70, 2,
	2, 675, 676, 5, 84, 43, 2, 676, 81, 3, 2, 2, 2, 677, 678, 5, 38, 20, 2,
	678, 679, 7, 70, 2, 2, 679, 83, 3, 2, 2, 2, 680, 683, 7, 29, 2, 2, 681,
	682, 7, 70, 2, 2, 682, 684, 7, 9, 2, 2, 683, 681, 3, 2, 2, 2, 683, 684,
	3, 2, 2, 2, 684, 685, 3, 2, 2, 2, 685, 694, 5, 86, 44, 2, 686, 689, 7,
	10, 2, 2, 687, 688, 7, 70, 2, 2, 688, 690, 7, 9, 2, 2, 689, 687, 3, 2,
	2, 2, 689, 690, 3, 2, 2, 2, 690, 691, 3, 2, 2, 2, 691, 693, 5, 86, 44,
	2, 692, 686, 3, 2, 2, 2, 693, 696, 3, 2, 2, 2, 694, 692, 3, 2, 2, 2, 694,
	695, 3, 2, 2, 2, 695, 697, 3, 2, 2, 2, 696, 694, 3, 2, 2, 2, 697, 698,
	7, 30, 2, 2, 698, 85, 3, 2, 2, 2, 699, 702, 5, 88, 45, 2, 700, 701, 7,
	64, 2, 2, 701, 703, 5, 86, 44, 2, 702, 700, 3, 2, 2, 2, 702, 703, 3, 2,
	2, 2, 703, 87, 3, 2, 2, 2, 704, 710, 5, 66, 34, 2, 705, 706, 7, 40, 2,
	2, 706, 707, 5, 86, 44, 2, 707, 708, 7, 9, 2, 2, 708, 709, 5, 86, 44, 2,
	709, 711, 3, 2, 2, 2, 710, 705, 3, 2, 2, 2, 710, 711, 3, 2, 2, 2, 711,
	89, 3, 2, 2, 2, 712, 713, 9, 8, 2, 2, 713, 91, 3, 2, 2, 2, 99, 104, 106,
	111, 123, 126, 133, 135, 139, 155, 159, 163, 166, 168, 173, 176, 180, 185,
	197, 200, 205, 207, 211, 214, 221, 224, 231, 234, 238, 246, 258, 268, 274,
	278, 282, 287, 293, 300, 304, 310, 315, 319, 329, 334, 341, 343, 353, 357,
	363, 367, 371, 374, 380, 382, 385, 395, 402, 409, 411, 413, 419, 424, 428,
	434, 440, 450, 455, 458, 473, 483, 495, 503, 506, 532, 538, 549, 556, 562,
	571, 577, 583, 590, 593, 606, 614, 619, 624, 630, 638, 653, 660, 665, 668,
	683, 689, 694, 702, 710,
}
var literalNames = []string{
	"", "';'", "'shared'", "'abstract'", "'final'", "'external'", "'class'",
	"':'", "','", "'{'", "'}'", "'typedef'", "'namespace'", "'private'", "'protected'",
	"'&'", "'~'", "'const'", "'interface'", "'='", "'import'", "'from'", "'enum'",
	"'funcdef'", "'get'", "'set'", "'mixin'", "'('", "')'", "'in'", "'out'",
	"'inout'", "'<'", "'>'", "'['", "']'", "'@'", "'::'", "'?'", "'auto'",
	"'switch'", "'break'", "'for'", "'while'", "'do'", "'if'", "'else'", "'continue'",
	"'try'", "'catch'", "'return'", "'case'", "'default'", "'.'", "'++'", "'--'",
	"'cast'", "'function'", "", "", "", "", "", "", "", "", "'void'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "BITOP", "MATHOP", "COMPOP", "LOGICOP", "ASSIGNOP", "PRIMTYPE",
	"FUNCATTR", "EXPRPREOP", "VOID", "LITERAL", "IDENTIFIER", "NUMBER", "STRING",
	"BITS", "COMMENT", "WS",
}

var ruleNames = []string{
	"script", "class_", "typdef", "namespace", "func_", "interface_", "var_",
	"import_", "enum_", "funcdef", "virtprop", "mixin_", "intfmthd", "statblock",
	"paramlist", "typemod", "type_", "initlist", "scope", "datatype", "statement",
	"switch_", "break_", "for_", "while_", "dowhile", "if_", "continue_", "exprstat",
	"try_", "return_", "case_", "expr", "exprterm", "exprvalue", "constructcall",
	"exprpostop", "cast", "lambda_", "funccall", "varaccess", "arglist", "assign",
	"condition", "exprop",
}

type angelscriptParser struct {
	*antlr.BaseParser
}

// NewangelscriptParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *angelscriptParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewangelscriptParser(input antlr.TokenStream) *angelscriptParser {
	this := new(angelscriptParser)
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
	this.GrammarFileName = "angelscript.g4"

	return this
}

// angelscriptParser tokens.
const (
	angelscriptParserEOF        = antlr.TokenEOF
	angelscriptParserT__0       = 1
	angelscriptParserT__1       = 2
	angelscriptParserT__2       = 3
	angelscriptParserT__3       = 4
	angelscriptParserT__4       = 5
	angelscriptParserT__5       = 6
	angelscriptParserT__6       = 7
	angelscriptParserT__7       = 8
	angelscriptParserT__8       = 9
	angelscriptParserT__9       = 10
	angelscriptParserT__10      = 11
	angelscriptParserT__11      = 12
	angelscriptParserT__12      = 13
	angelscriptParserT__13      = 14
	angelscriptParserT__14      = 15
	angelscriptParserT__15      = 16
	angelscriptParserT__16      = 17
	angelscriptParserT__17      = 18
	angelscriptParserT__18      = 19
	angelscriptParserT__19      = 20
	angelscriptParserT__20      = 21
	angelscriptParserT__21      = 22
	angelscriptParserT__22      = 23
	angelscriptParserT__23      = 24
	angelscriptParserT__24      = 25
	angelscriptParserT__25      = 26
	angelscriptParserT__26      = 27
	angelscriptParserT__27      = 28
	angelscriptParserT__28      = 29
	angelscriptParserT__29      = 30
	angelscriptParserT__30      = 31
	angelscriptParserT__31      = 32
	angelscriptParserT__32      = 33
	angelscriptParserT__33      = 34
	angelscriptParserT__34      = 35
	angelscriptParserT__35      = 36
	angelscriptParserT__36      = 37
	angelscriptParserT__37      = 38
	angelscriptParserT__38      = 39
	angelscriptParserT__39      = 40
	angelscriptParserT__40      = 41
	angelscriptParserT__41      = 42
	angelscriptParserT__42      = 43
	angelscriptParserT__43      = 44
	angelscriptParserT__44      = 45
	angelscriptParserT__45      = 46
	angelscriptParserT__46      = 47
	angelscriptParserT__47      = 48
	angelscriptParserT__48      = 49
	angelscriptParserT__49      = 50
	angelscriptParserT__50      = 51
	angelscriptParserT__51      = 52
	angelscriptParserT__52      = 53
	angelscriptParserT__53      = 54
	angelscriptParserT__54      = 55
	angelscriptParserT__55      = 56
	angelscriptParserT__56      = 57
	angelscriptParserBITOP      = 58
	angelscriptParserMATHOP     = 59
	angelscriptParserCOMPOP     = 60
	angelscriptParserLOGICOP    = 61
	angelscriptParserASSIGNOP   = 62
	angelscriptParserPRIMTYPE   = 63
	angelscriptParserFUNCATTR   = 64
	angelscriptParserEXPRPREOP  = 65
	angelscriptParserVOID       = 66
	angelscriptParserLITERAL    = 67
	angelscriptParserIDENTIFIER = 68
	angelscriptParserNUMBER     = 69
	angelscriptParserSTRING     = 70
	angelscriptParserBITS       = 71
	angelscriptParserCOMMENT    = 72
	angelscriptParserWS         = 73
)

// angelscriptParser rules.
const (
	angelscriptParserRULE_script        = 0
	angelscriptParserRULE_class_        = 1
	angelscriptParserRULE_typdef        = 2
	angelscriptParserRULE_namespace     = 3
	angelscriptParserRULE_func_         = 4
	angelscriptParserRULE_interface_    = 5
	angelscriptParserRULE_var_          = 6
	angelscriptParserRULE_import_       = 7
	angelscriptParserRULE_enum_         = 8
	angelscriptParserRULE_funcdef       = 9
	angelscriptParserRULE_virtprop      = 10
	angelscriptParserRULE_mixin_        = 11
	angelscriptParserRULE_intfmthd      = 12
	angelscriptParserRULE_statblock     = 13
	angelscriptParserRULE_paramlist     = 14
	angelscriptParserRULE_typemod       = 15
	angelscriptParserRULE_type_         = 16
	angelscriptParserRULE_initlist      = 17
	angelscriptParserRULE_scope         = 18
	angelscriptParserRULE_datatype      = 19
	angelscriptParserRULE_statement     = 20
	angelscriptParserRULE_switch_       = 21
	angelscriptParserRULE_break_        = 22
	angelscriptParserRULE_for_          = 23
	angelscriptParserRULE_while_        = 24
	angelscriptParserRULE_dowhile       = 25
	angelscriptParserRULE_if_           = 26
	angelscriptParserRULE_continue_     = 27
	angelscriptParserRULE_exprstat      = 28
	angelscriptParserRULE_try_          = 29
	angelscriptParserRULE_return_       = 30
	angelscriptParserRULE_case_         = 31
	angelscriptParserRULE_expr          = 32
	angelscriptParserRULE_exprterm      = 33
	angelscriptParserRULE_exprvalue     = 34
	angelscriptParserRULE_constructcall = 35
	angelscriptParserRULE_exprpostop    = 36
	angelscriptParserRULE_cast          = 37
	angelscriptParserRULE_lambda_       = 38
	angelscriptParserRULE_funccall      = 39
	angelscriptParserRULE_varaccess     = 40
	angelscriptParserRULE_arglist       = 41
	angelscriptParserRULE_assign        = 42
	angelscriptParserRULE_condition     = 43
	angelscriptParserRULE_exprop        = 44
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) AllImport_() []IImport_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImport_Context)(nil)).Elem())
	var tst = make([]IImport_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImport_Context)
		}
	}

	return tst
}

func (s *ScriptContext) Import_(i int) IImport_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImport_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImport_Context)
}

func (s *ScriptContext) AllEnum_() []IEnum_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnum_Context)(nil)).Elem())
	var tst = make([]IEnum_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnum_Context)
		}
	}

	return tst
}

func (s *ScriptContext) Enum_(i int) IEnum_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnum_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnum_Context)
}

func (s *ScriptContext) AllTypdef() []ITypdefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypdefContext)(nil)).Elem())
	var tst = make([]ITypdefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypdefContext)
		}
	}

	return tst
}

func (s *ScriptContext) Typdef(i int) ITypdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypdefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypdefContext)
}

func (s *ScriptContext) AllClass_() []IClass_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClass_Context)(nil)).Elem())
	var tst = make([]IClass_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClass_Context)
		}
	}

	return tst
}

func (s *ScriptContext) Class_(i int) IClass_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClass_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClass_Context)
}

func (s *ScriptContext) AllMixin_() []IMixin_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMixin_Context)(nil)).Elem())
	var tst = make([]IMixin_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMixin_Context)
		}
	}

	return tst
}

func (s *ScriptContext) Mixin_(i int) IMixin_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMixin_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMixin_Context)
}

func (s *ScriptContext) AllInterface_() []IInterface_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInterface_Context)(nil)).Elem())
	var tst = make([]IInterface_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInterface_Context)
		}
	}

	return tst
}

func (s *ScriptContext) Interface_(i int) IInterface_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterface_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInterface_Context)
}

func (s *ScriptContext) AllFuncdef() []IFuncdefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncdefContext)(nil)).Elem())
	var tst = make([]IFuncdefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncdefContext)
		}
	}

	return tst
}

func (s *ScriptContext) Funcdef(i int) IFuncdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncdefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncdefContext)
}

func (s *ScriptContext) AllVirtprop() []IVirtpropContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVirtpropContext)(nil)).Elem())
	var tst = make([]IVirtpropContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVirtpropContext)
		}
	}

	return tst
}

func (s *ScriptContext) Virtprop(i int) IVirtpropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVirtpropContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVirtpropContext)
}

func (s *ScriptContext) AllVar_() []IVar_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVar_Context)(nil)).Elem())
	var tst = make([]IVar_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVar_Context)
		}
	}

	return tst
}

func (s *ScriptContext) Var_(i int) IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *ScriptContext) AllFunc_() []IFunc_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunc_Context)(nil)).Elem())
	var tst = make([]IFunc_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunc_Context)
		}
	}

	return tst
}

func (s *ScriptContext) Func_(i int) IFunc_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunc_Context)
}

func (s *ScriptContext) AllNamespace() []INamespaceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INamespaceContext)(nil)).Elem())
	var tst = make([]INamespaceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INamespaceContext)
		}
	}

	return tst
}

func (s *ScriptContext) Namespace(i int) INamespaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INamespaceContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *angelscriptParser) Script() (localctx IScriptContext) {
	this := p
	_ = this

	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, angelscriptParserRULE_script)
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

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__0)|(1<<angelscriptParserT__1)|(1<<angelscriptParserT__2)|(1<<angelscriptParserT__3)|(1<<angelscriptParserT__4)|(1<<angelscriptParserT__5)|(1<<angelscriptParserT__10)|(1<<angelscriptParserT__11)|(1<<angelscriptParserT__12)|(1<<angelscriptParserT__13)|(1<<angelscriptParserT__15)|(1<<angelscriptParserT__16)|(1<<angelscriptParserT__17)|(1<<angelscriptParserT__19)|(1<<angelscriptParserT__21)|(1<<angelscriptParserT__22)|(1<<angelscriptParserT__25))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(angelscriptParserT__36-37))|(1<<(angelscriptParserT__37-37))|(1<<(angelscriptParserT__38-37))|(1<<(angelscriptParserPRIMTYPE-37))|(1<<(angelscriptParserIDENTIFIER-37)))) != 0) {
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(90)
				p.Import_()
			}

		case 2:
			{
				p.SetState(91)
				p.Enum_()
			}

		case 3:
			{
				p.SetState(92)
				p.Typdef()
			}

		case 4:
			{
				p.SetState(93)
				p.Class_()
			}

		case 5:
			{
				p.SetState(94)
				p.Mixin_()
			}

		case 6:
			{
				p.SetState(95)
				p.Interface_()
			}

		case 7:
			{
				p.SetState(96)
				p.Funcdef()
			}

		case 8:
			{
				p.SetState(97)
				p.Virtprop()
			}

		case 9:
			{
				p.SetState(98)
				p.Var_()
			}

		case 10:
			{
				p.SetState(99)
				p.Func_()
			}

		case 11:
			{
				p.SetState(100)
				p.Namespace()
			}

		case 12:
			{
				p.SetState(101)
				p.Match(angelscriptParserT__0)
			}

		}

		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IClass_Context is an interface to support dynamic dispatch.
type IClass_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClass_Context differentiates from other interfaces.
	IsClass_Context()
}

type Class_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClass_Context() *Class_Context {
	var p = new(Class_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_class_
	return p
}

func (*Class_Context) IsClass_Context() {}

func NewClass_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Class_Context {
	var p = new(Class_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_class_

	return p
}

func (s *Class_Context) GetParser() antlr.Parser { return s.parser }

func (s *Class_Context) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserIDENTIFIER)
}

func (s *Class_Context) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, i)
}

func (s *Class_Context) AllVirtprop() []IVirtpropContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVirtpropContext)(nil)).Elem())
	var tst = make([]IVirtpropContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVirtpropContext)
		}
	}

	return tst
}

func (s *Class_Context) Virtprop(i int) IVirtpropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVirtpropContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVirtpropContext)
}

func (s *Class_Context) AllFunc_() []IFunc_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunc_Context)(nil)).Elem())
	var tst = make([]IFunc_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunc_Context)
		}
	}

	return tst
}

func (s *Class_Context) Func_(i int) IFunc_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunc_Context)
}

func (s *Class_Context) AllVar_() []IVar_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVar_Context)(nil)).Elem())
	var tst = make([]IVar_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVar_Context)
		}
	}

	return tst
}

func (s *Class_Context) Var_(i int) IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *Class_Context) AllFuncdef() []IFuncdefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncdefContext)(nil)).Elem())
	var tst = make([]IFuncdefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncdefContext)
		}
	}

	return tst
}

func (s *Class_Context) Funcdef(i int) IFuncdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncdefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncdefContext)
}

func (s *Class_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Class_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Class_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterClass_(s)
	}
}

func (s *Class_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitClass_(s)
	}
}

func (p *angelscriptParser) Class_() (localctx IClass_Context) {
	this := p
	_ = this

	localctx = NewClass_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, angelscriptParserRULE_class_)
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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__1)|(1<<angelscriptParserT__2)|(1<<angelscriptParserT__3)|(1<<angelscriptParserT__4))) != 0 {
		{
			p.SetState(106)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__1)|(1<<angelscriptParserT__2)|(1<<angelscriptParserT__3)|(1<<angelscriptParserT__4))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(112)
		p.Match(angelscriptParserT__5)
	}
	{
		p.SetState(113)
		p.Match(angelscriptParserIDENTIFIER)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case angelscriptParserT__0:
		{
			p.SetState(114)
			p.Match(angelscriptParserT__0)
		}

	case angelscriptParserT__6, angelscriptParserT__8:
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == angelscriptParserT__6 {
			{
				p.SetState(115)
				p.Match(angelscriptParserT__6)
			}
			{
				p.SetState(116)
				p.Match(angelscriptParserIDENTIFIER)
			}
			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == angelscriptParserT__7 {
				{
					p.SetState(117)
					p.Match(angelscriptParserT__7)
				}
				{
					p.SetState(118)
					p.Match(angelscriptParserIDENTIFIER)
				}

				p.SetState(123)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(126)
			p.Match(angelscriptParserT__8)
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__1)|(1<<angelscriptParserT__4)|(1<<angelscriptParserT__12)|(1<<angelscriptParserT__13)|(1<<angelscriptParserT__15)|(1<<angelscriptParserT__16)|(1<<angelscriptParserT__22))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(angelscriptParserT__36-37))|(1<<(angelscriptParserT__37-37))|(1<<(angelscriptParserT__38-37))|(1<<(angelscriptParserPRIMTYPE-37))|(1<<(angelscriptParserIDENTIFIER-37)))) != 0) {
			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(127)
					p.Virtprop()
				}

			case 2:
				{
					p.SetState(128)
					p.Func_()
				}

			case 3:
				{
					p.SetState(129)
					p.Var_()
				}

			case 4:
				{
					p.SetState(130)
					p.Funcdef()
				}

			}

			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(136)
			p.Match(angelscriptParserT__9)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypdefContext is an interface to support dynamic dispatch.
type ITypdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypdefContext differentiates from other interfaces.
	IsTypdefContext()
}

type TypdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypdefContext() *TypdefContext {
	var p = new(TypdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_typdef
	return p
}

func (*TypdefContext) IsTypdefContext() {}

func NewTypdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypdefContext {
	var p = new(TypdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_typdef

	return p
}

func (s *TypdefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypdefContext) PRIMTYPE() antlr.TerminalNode {
	return s.GetToken(angelscriptParserPRIMTYPE, 0)
}

func (s *TypdefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, 0)
}

func (s *TypdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterTypdef(s)
	}
}

func (s *TypdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitTypdef(s)
	}
}

func (p *angelscriptParser) Typdef() (localctx ITypdefContext) {
	this := p
	_ = this

	localctx = NewTypdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, angelscriptParserRULE_typdef)

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
		p.SetState(139)
		p.Match(angelscriptParserT__10)
	}
	{
		p.SetState(140)
		p.Match(angelscriptParserPRIMTYPE)
	}
	{
		p.SetState(141)
		p.Match(angelscriptParserIDENTIFIER)
	}
	{
		p.SetState(142)
		p.Match(angelscriptParserT__0)
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
	p.RuleIndex = angelscriptParserRULE_namespace
	return p
}

func (*NamespaceContext) IsNamespaceContext() {}

func NewNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceContext {
	var p = new(NamespaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_namespace

	return p
}

func (s *NamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, 0)
}

func (s *NamespaceContext) Script() IScriptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScriptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScriptContext)
}

func (s *NamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterNamespace(s)
	}
}

func (s *NamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitNamespace(s)
	}
}

func (p *angelscriptParser) Namespace() (localctx INamespaceContext) {
	this := p
	_ = this

	localctx = NewNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, angelscriptParserRULE_namespace)

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
		p.SetState(144)
		p.Match(angelscriptParserT__11)
	}
	{
		p.SetState(145)
		p.Match(angelscriptParserIDENTIFIER)
	}
	{
		p.SetState(146)
		p.Match(angelscriptParserT__8)
	}
	{
		p.SetState(147)
		p.Script()
	}
	{
		p.SetState(148)
		p.Match(angelscriptParserT__9)
	}

	return localctx
}

// IFunc_Context is an interface to support dynamic dispatch.
type IFunc_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_Context differentiates from other interfaces.
	IsFunc_Context()
}

type Func_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_Context() *Func_Context {
	var p = new(Func_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_func_
	return p
}

func (*Func_Context) IsFunc_Context() {}

func NewFunc_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_Context {
	var p = new(Func_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_func_

	return p
}

func (s *Func_Context) GetParser() antlr.Parser { return s.parser }

func (s *Func_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, 0)
}

func (s *Func_Context) Paramlist() IParamlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamlistContext)
}

func (s *Func_Context) Statblock() IStatblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatblockContext)
}

func (s *Func_Context) FUNCATTR() antlr.TerminalNode {
	return s.GetToken(angelscriptParserFUNCATTR, 0)
}

func (s *Func_Context) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Func_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterFunc_(s)
	}
}

func (s *Func_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitFunc_(s)
	}
}

func (p *angelscriptParser) Func_() (localctx IFunc_Context) {
	this := p
	_ = this

	localctx = NewFunc_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, angelscriptParserRULE_func_)
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
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__1 || _la == angelscriptParserT__4 {
		{
			p.SetState(150)
			_la = p.GetTokenStream().LA(1)

			if !(_la == angelscriptParserT__1 || _la == angelscriptParserT__4) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__12 || _la == angelscriptParserT__13 {
		{
			p.SetState(156)
			_la = p.GetTokenStream().LA(1)

			if !(_la == angelscriptParserT__12 || _la == angelscriptParserT__13) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		p.SetState(164)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case angelscriptParserT__16, angelscriptParserT__36, angelscriptParserT__37, angelscriptParserT__38, angelscriptParserPRIMTYPE, angelscriptParserIDENTIFIER:
			{
				p.SetState(159)
				p.Type_()
			}
			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == angelscriptParserT__14 {
				{
					p.SetState(160)
					p.Match(angelscriptParserT__14)
				}

			}

		case angelscriptParserT__15:
			{
				p.SetState(163)
				p.Match(angelscriptParserT__15)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}
	{
		p.SetState(168)
		p.Match(angelscriptParserIDENTIFIER)
	}
	{
		p.SetState(169)
		p.Paramlist()
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__16 {
		{
			p.SetState(170)
			p.Match(angelscriptParserT__16)
		}

	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserFUNCATTR {
		{
			p.SetState(173)
			p.Match(angelscriptParserFUNCATTR)
		}

	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case angelscriptParserT__0:
		{
			p.SetState(176)
			p.Match(angelscriptParserT__0)
		}

	case angelscriptParserT__8:
		{
			p.SetState(177)
			p.Statblock()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInterface_Context is an interface to support dynamic dispatch.
type IInterface_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterface_Context differentiates from other interfaces.
	IsInterface_Context()
}

type Interface_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterface_Context() *Interface_Context {
	var p = new(Interface_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_interface_
	return p
}

func (*Interface_Context) IsInterface_Context() {}

func NewInterface_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Interface_Context {
	var p = new(Interface_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_interface_

	return p
}

func (s *Interface_Context) GetParser() antlr.Parser { return s.parser }

func (s *Interface_Context) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserIDENTIFIER)
}

func (s *Interface_Context) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, i)
}

func (s *Interface_Context) AllVirtprop() []IVirtpropContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVirtpropContext)(nil)).Elem())
	var tst = make([]IVirtpropContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVirtpropContext)
		}
	}

	return tst
}

func (s *Interface_Context) Virtprop(i int) IVirtpropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVirtpropContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVirtpropContext)
}

func (s *Interface_Context) AllIntfmthd() []IIntfmthdContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntfmthdContext)(nil)).Elem())
	var tst = make([]IIntfmthdContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntfmthdContext)
		}
	}

	return tst
}

func (s *Interface_Context) Intfmthd(i int) IIntfmthdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntfmthdContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntfmthdContext)
}

func (s *Interface_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Interface_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Interface_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterInterface_(s)
	}
}

func (s *Interface_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitInterface_(s)
	}
}

func (p *angelscriptParser) Interface_() (localctx IInterface_Context) {
	this := p
	_ = this

	localctx = NewInterface_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, angelscriptParserRULE_interface_)
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
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__1 || _la == angelscriptParserT__4 {
		{
			p.SetState(180)
			_la = p.GetTokenStream().LA(1)

			if !(_la == angelscriptParserT__1 || _la == angelscriptParserT__4) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(186)
		p.Match(angelscriptParserT__17)
	}
	{
		p.SetState(187)
		p.Match(angelscriptParserIDENTIFIER)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case angelscriptParserT__0:
		{
			p.SetState(188)
			p.Match(angelscriptParserT__0)
		}

	case angelscriptParserT__6, angelscriptParserT__8:
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == angelscriptParserT__6 {
			{
				p.SetState(189)
				p.Match(angelscriptParserT__6)
			}
			{
				p.SetState(190)
				p.Match(angelscriptParserIDENTIFIER)
			}
			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == angelscriptParserT__7 {
				{
					p.SetState(191)
					p.Match(angelscriptParserT__7)
				}
				{
					p.SetState(192)
					p.Match(angelscriptParserIDENTIFIER)
				}

				p.SetState(197)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(200)
			p.Match(angelscriptParserT__8)
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__12)|(1<<angelscriptParserT__13)|(1<<angelscriptParserT__16))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(angelscriptParserT__36-37))|(1<<(angelscriptParserT__37-37))|(1<<(angelscriptParserT__38-37))|(1<<(angelscriptParserPRIMTYPE-37))|(1<<(angelscriptParserIDENTIFIER-37)))) != 0) {
			p.SetState(203)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(201)
					p.Virtprop()
				}

			case 2:
				{
					p.SetState(202)
					p.Intfmthd()
				}

			}

			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(208)
			p.Match(angelscriptParserT__9)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVar_Context is an interface to support dynamic dispatch.
type IVar_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_Context differentiates from other interfaces.
	IsVar_Context()
}

type Var_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_Context() *Var_Context {
	var p = new(Var_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_var_
	return p
}

func (*Var_Context) IsVar_Context() {}

func NewVar_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_Context {
	var p = new(Var_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_var_

	return p
}

func (s *Var_Context) GetParser() antlr.Parser { return s.parser }

func (s *Var_Context) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Var_Context) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserIDENTIFIER)
}

func (s *Var_Context) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, i)
}

func (s *Var_Context) AllArglist() []IArglistContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArglistContext)(nil)).Elem())
	var tst = make([]IArglistContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArglistContext)
		}
	}

	return tst
}

func (s *Var_Context) Arglist(i int) IArglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArglistContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *Var_Context) AllInitlist() []IInitlistContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInitlistContext)(nil)).Elem())
	var tst = make([]IInitlistContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInitlistContext)
		}
	}

	return tst
}

func (s *Var_Context) Initlist(i int) IInitlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitlistContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInitlistContext)
}

func (s *Var_Context) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Var_Context) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Var_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterVar_(s)
	}
}

func (s *Var_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitVar_(s)
	}
}

func (p *angelscriptParser) Var_() (localctx IVar_Context) {
	this := p
	_ = this

	localctx = NewVar_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, angelscriptParserRULE_var_)
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
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__12 || _la == angelscriptParserT__13 {
		{
			p.SetState(211)
			_la = p.GetTokenStream().LA(1)

			if !(_la == angelscriptParserT__12 || _la == angelscriptParserT__13) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(214)
		p.Type_()
	}
	{
		p.SetState(215)
		p.Match(angelscriptParserIDENTIFIER)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case angelscriptParserT__18:
		{
			p.SetState(216)
			p.Match(angelscriptParserT__18)
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(217)
				p.Initlist()
			}

		case 2:
			{
				p.SetState(218)
				p.Expr()
			}

		}

	case angelscriptParserT__26:
		{
			p.SetState(221)
			p.Arglist()
		}

	case angelscriptParserT__0, angelscriptParserT__7:

	default:
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__7 {
		{
			p.SetState(224)
			p.Match(angelscriptParserT__7)
		}
		{
			p.SetState(225)
			p.Match(angelscriptParserIDENTIFIER)
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case angelscriptParserT__18:
			{
				p.SetState(226)
				p.Match(angelscriptParserT__18)
			}
			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(227)
					p.Initlist()
				}

			case 2:
				{
					p.SetState(228)
					p.Expr()
				}

			}

		case angelscriptParserT__26:
			{
				p.SetState(231)
				p.Arglist()
			}

		case angelscriptParserT__0, angelscriptParserT__7:

		default:
		}

		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(239)
		p.Match(angelscriptParserT__0)
	}

	return localctx
}

// IImport_Context is an interface to support dynamic dispatch.
type IImport_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImport_Context differentiates from other interfaces.
	IsImport_Context()
}

type Import_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_Context() *Import_Context {
	var p = new(Import_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_import_
	return p
}

func (*Import_Context) IsImport_Context() {}

func NewImport_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_Context {
	var p = new(Import_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_import_

	return p
}

func (s *Import_Context) GetParser() antlr.Parser { return s.parser }

func (s *Import_Context) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Import_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, 0)
}

func (s *Import_Context) Paramlist() IParamlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamlistContext)
}

func (s *Import_Context) FUNCATTR() antlr.TerminalNode {
	return s.GetToken(angelscriptParserFUNCATTR, 0)
}

func (s *Import_Context) STRING() antlr.TerminalNode {
	return s.GetToken(angelscriptParserSTRING, 0)
}

func (s *Import_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterImport_(s)
	}
}

func (s *Import_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitImport_(s)
	}
}

func (p *angelscriptParser) Import_() (localctx IImport_Context) {
	this := p
	_ = this

	localctx = NewImport_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, angelscriptParserRULE_import_)
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
		p.SetState(241)
		p.Match(angelscriptParserT__19)
	}
	{
		p.SetState(242)
		p.Type_()
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__14 {
		{
			p.SetState(243)
			p.Match(angelscriptParserT__14)
		}

	}
	{
		p.SetState(246)
		p.Match(angelscriptParserIDENTIFIER)
	}
	{
		p.SetState(247)
		p.Paramlist()
	}
	{
		p.SetState(248)
		p.Match(angelscriptParserFUNCATTR)
	}
	{
		p.SetState(249)
		p.Match(angelscriptParserT__20)
	}
	{
		p.SetState(250)
		p.Match(angelscriptParserSTRING)
	}
	{
		p.SetState(251)
		p.Match(angelscriptParserT__0)
	}

	return localctx
}

// IEnum_Context is an interface to support dynamic dispatch.
type IEnum_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnum_Context differentiates from other interfaces.
	IsEnum_Context()
}

type Enum_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnum_Context() *Enum_Context {
	var p = new(Enum_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_enum_
	return p
}

func (*Enum_Context) IsEnum_Context() {}

func NewEnum_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_Context {
	var p = new(Enum_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_enum_

	return p
}

func (s *Enum_Context) GetParser() antlr.Parser { return s.parser }

func (s *Enum_Context) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserIDENTIFIER)
}

func (s *Enum_Context) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, i)
}

func (s *Enum_Context) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Enum_Context) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Enum_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterEnum_(s)
	}
}

func (s *Enum_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitEnum_(s)
	}
}

func (p *angelscriptParser) Enum_() (localctx IEnum_Context) {
	this := p
	_ = this

	localctx = NewEnum_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, angelscriptParserRULE_enum_)
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
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__1 || _la == angelscriptParserT__4 {
		{
			p.SetState(253)
			_la = p.GetTokenStream().LA(1)

			if !(_la == angelscriptParserT__1 || _la == angelscriptParserT__4) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(259)
		p.Match(angelscriptParserT__21)
	}
	{
		p.SetState(260)
		p.Match(angelscriptParserIDENTIFIER)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case angelscriptParserT__0:
		{
			p.SetState(261)
			p.Match(angelscriptParserT__0)
		}

	case angelscriptParserT__8:
		{
			p.SetState(262)
			p.Match(angelscriptParserT__8)
		}
		{
			p.SetState(263)
			p.Match(angelscriptParserIDENTIFIER)
		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == angelscriptParserT__18 {
			{
				p.SetState(264)
				p.Match(angelscriptParserT__18)
			}
			{
				p.SetState(265)
				p.Expr()
			}

		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == angelscriptParserT__7 {
			{
				p.SetState(268)
				p.Match(angelscriptParserT__7)
			}
			{
				p.SetState(269)
				p.Match(angelscriptParserIDENTIFIER)
			}
			p.SetState(272)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == angelscriptParserT__18 {
				{
					p.SetState(270)
					p.Match(angelscriptParserT__18)
				}
				{
					p.SetState(271)
					p.Expr()
				}

			}

			p.SetState(278)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(279)
			p.Match(angelscriptParserT__9)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncdefContext is an interface to support dynamic dispatch.
type IFuncdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncdefContext differentiates from other interfaces.
	IsFuncdefContext()
}

type FuncdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncdefContext() *FuncdefContext {
	var p = new(FuncdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_funcdef
	return p
}

func (*FuncdefContext) IsFuncdefContext() {}

func NewFuncdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncdefContext {
	var p = new(FuncdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_funcdef

	return p
}

func (s *FuncdefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncdefContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FuncdefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, 0)
}

func (s *FuncdefContext) Paramlist() IParamlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamlistContext)
}

func (s *FuncdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterFuncdef(s)
	}
}

func (s *FuncdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitFuncdef(s)
	}
}

func (p *angelscriptParser) Funcdef() (localctx IFuncdefContext) {
	this := p
	_ = this

	localctx = NewFuncdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, angelscriptParserRULE_funcdef)
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
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__1 || _la == angelscriptParserT__4 {
		{
			p.SetState(282)
			_la = p.GetTokenStream().LA(1)

			if !(_la == angelscriptParserT__1 || _la == angelscriptParserT__4) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(288)
		p.Match(angelscriptParserT__22)
	}
	{
		p.SetState(289)
		p.Type_()
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__14 {
		{
			p.SetState(290)
			p.Match(angelscriptParserT__14)
		}

	}
	{
		p.SetState(293)
		p.Match(angelscriptParserIDENTIFIER)
	}
	{
		p.SetState(294)
		p.Paramlist()
	}
	{
		p.SetState(295)
		p.Match(angelscriptParserT__0)
	}

	return localctx
}

// IVirtpropContext is an interface to support dynamic dispatch.
type IVirtpropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVirtpropContext differentiates from other interfaces.
	IsVirtpropContext()
}

type VirtpropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVirtpropContext() *VirtpropContext {
	var p = new(VirtpropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_virtprop
	return p
}

func (*VirtpropContext) IsVirtpropContext() {}

func NewVirtpropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VirtpropContext {
	var p = new(VirtpropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_virtprop

	return p
}

func (s *VirtpropContext) GetParser() antlr.Parser { return s.parser }

func (s *VirtpropContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VirtpropContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, 0)
}

func (s *VirtpropContext) AllFUNCATTR() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserFUNCATTR)
}

func (s *VirtpropContext) FUNCATTR(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserFUNCATTR, i)
}

func (s *VirtpropContext) AllStatblock() []IStatblockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatblockContext)(nil)).Elem())
	var tst = make([]IStatblockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatblockContext)
		}
	}

	return tst
}

func (s *VirtpropContext) Statblock(i int) IStatblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatblockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatblockContext)
}

func (s *VirtpropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VirtpropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VirtpropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterVirtprop(s)
	}
}

func (s *VirtpropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitVirtprop(s)
	}
}

func (p *angelscriptParser) Virtprop() (localctx IVirtpropContext) {
	this := p
	_ = this

	localctx = NewVirtpropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, angelscriptParserRULE_virtprop)
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
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__12 || _la == angelscriptParserT__13 {
		{
			p.SetState(297)
			_la = p.GetTokenStream().LA(1)

			if !(_la == angelscriptParserT__12 || _la == angelscriptParserT__13) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(300)
		p.Type_()
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__14 {
		{
			p.SetState(301)
			p.Match(angelscriptParserT__14)
		}

	}
	{
		p.SetState(304)
		p.Match(angelscriptParserIDENTIFIER)
	}
	{
		p.SetState(305)
		p.Match(angelscriptParserT__8)
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__23 || _la == angelscriptParserT__24 {
		{
			p.SetState(306)
			_la = p.GetTokenStream().LA(1)

			if !(_la == angelscriptParserT__23 || _la == angelscriptParserT__24) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == angelscriptParserT__16 {
			{
				p.SetState(307)
				p.Match(angelscriptParserT__16)
			}

		}
		{
			p.SetState(310)
			p.Match(angelscriptParserFUNCATTR)
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case angelscriptParserT__8:
			{
				p.SetState(311)
				p.Statblock()
			}

		case angelscriptParserT__0:
			{
				p.SetState(312)
				p.Match(angelscriptParserT__0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(320)
		p.Match(angelscriptParserT__9)
	}

	return localctx
}

// IMixin_Context is an interface to support dynamic dispatch.
type IMixin_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixin_Context differentiates from other interfaces.
	IsMixin_Context()
}

type Mixin_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixin_Context() *Mixin_Context {
	var p = new(Mixin_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_mixin_
	return p
}

func (*Mixin_Context) IsMixin_Context() {}

func NewMixin_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mixin_Context {
	var p = new(Mixin_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_mixin_

	return p
}

func (s *Mixin_Context) GetParser() antlr.Parser { return s.parser }

func (s *Mixin_Context) Class_() IClass_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClass_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClass_Context)
}

func (s *Mixin_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mixin_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mixin_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterMixin_(s)
	}
}

func (s *Mixin_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitMixin_(s)
	}
}

func (p *angelscriptParser) Mixin_() (localctx IMixin_Context) {
	this := p
	_ = this

	localctx = NewMixin_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, angelscriptParserRULE_mixin_)

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
		p.SetState(322)
		p.Match(angelscriptParserT__25)
	}
	{
		p.SetState(323)
		p.Class_()
	}

	return localctx
}

// IIntfmthdContext is an interface to support dynamic dispatch.
type IIntfmthdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntfmthdContext differentiates from other interfaces.
	IsIntfmthdContext()
}

type IntfmthdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntfmthdContext() *IntfmthdContext {
	var p = new(IntfmthdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_intfmthd
	return p
}

func (*IntfmthdContext) IsIntfmthdContext() {}

func NewIntfmthdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntfmthdContext {
	var p = new(IntfmthdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_intfmthd

	return p
}

func (s *IntfmthdContext) GetParser() antlr.Parser { return s.parser }

func (s *IntfmthdContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *IntfmthdContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, 0)
}

func (s *IntfmthdContext) Paramlist() IParamlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamlistContext)
}

func (s *IntfmthdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntfmthdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntfmthdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterIntfmthd(s)
	}
}

func (s *IntfmthdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitIntfmthd(s)
	}
}

func (p *angelscriptParser) Intfmthd() (localctx IIntfmthdContext) {
	this := p
	_ = this

	localctx = NewIntfmthdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, angelscriptParserRULE_intfmthd)
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
		p.SetState(325)
		p.Type_()
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__14 {
		{
			p.SetState(326)
			p.Match(angelscriptParserT__14)
		}

	}
	{
		p.SetState(329)
		p.Match(angelscriptParserIDENTIFIER)
	}
	{
		p.SetState(330)
		p.Paramlist()
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__16 {
		{
			p.SetState(331)
			p.Match(angelscriptParserT__16)
		}

	}
	{
		p.SetState(334)
		p.Match(angelscriptParserT__0)
	}

	return localctx
}

// IStatblockContext is an interface to support dynamic dispatch.
type IStatblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatblockContext differentiates from other interfaces.
	IsStatblockContext()
}

type StatblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatblockContext() *StatblockContext {
	var p = new(StatblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_statblock
	return p
}

func (*StatblockContext) IsStatblockContext() {}

func NewStatblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatblockContext {
	var p = new(StatblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_statblock

	return p
}

func (s *StatblockContext) GetParser() antlr.Parser { return s.parser }

func (s *StatblockContext) AllVar_() []IVar_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVar_Context)(nil)).Elem())
	var tst = make([]IVar_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVar_Context)
		}
	}

	return tst
}

func (s *StatblockContext) Var_(i int) IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *StatblockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatblockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterStatblock(s)
	}
}

func (s *StatblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitStatblock(s)
	}
}

func (p *angelscriptParser) Statblock() (localctx IStatblockContext) {
	this := p
	_ = this

	localctx = NewStatblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, angelscriptParserRULE_statblock)
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
		p.SetState(336)
		p.Match(angelscriptParserT__8)
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__0)|(1<<angelscriptParserT__8)|(1<<angelscriptParserT__12)|(1<<angelscriptParserT__13)|(1<<angelscriptParserT__16)|(1<<angelscriptParserT__26))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(angelscriptParserT__36-37))|(1<<(angelscriptParserT__37-37))|(1<<(angelscriptParserT__38-37))|(1<<(angelscriptParserT__39-37))|(1<<(angelscriptParserT__40-37))|(1<<(angelscriptParserT__41-37))|(1<<(angelscriptParserT__42-37))|(1<<(angelscriptParserT__43-37))|(1<<(angelscriptParserT__44-37))|(1<<(angelscriptParserT__46-37))|(1<<(angelscriptParserT__47-37))|(1<<(angelscriptParserT__49-37))|(1<<(angelscriptParserT__55-37))|(1<<(angelscriptParserT__56-37))|(1<<(angelscriptParserPRIMTYPE-37))|(1<<(angelscriptParserEXPRPREOP-37))|(1<<(angelscriptParserVOID-37))|(1<<(angelscriptParserLITERAL-37))|(1<<(angelscriptParserIDENTIFIER-37)))) != 0) {
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(337)
				p.Var_()
			}

		case 2:
			{
				p.SetState(338)
				p.Statement()
			}

		}

		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(344)
		p.Match(angelscriptParserT__9)
	}

	return localctx
}

// IParamlistContext is an interface to support dynamic dispatch.
type IParamlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamlistContext differentiates from other interfaces.
	IsParamlistContext()
}

type ParamlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamlistContext() *ParamlistContext {
	var p = new(ParamlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_paramlist
	return p
}

func (*ParamlistContext) IsParamlistContext() {}

func NewParamlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamlistContext {
	var p = new(ParamlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_paramlist

	return p
}

func (s *ParamlistContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamlistContext) VOID() antlr.TerminalNode {
	return s.GetToken(angelscriptParserVOID, 0)
}

func (s *ParamlistContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *ParamlistContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ParamlistContext) AllTypemod() []ITypemodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypemodContext)(nil)).Elem())
	var tst = make([]ITypemodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypemodContext)
		}
	}

	return tst
}

func (s *ParamlistContext) Typemod(i int) ITypemodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypemodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypemodContext)
}

func (s *ParamlistContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserIDENTIFIER)
}

func (s *ParamlistContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, i)
}

func (s *ParamlistContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ParamlistContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParamlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterParamlist(s)
	}
}

func (s *ParamlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitParamlist(s)
	}
}

func (p *angelscriptParser) Paramlist() (localctx IParamlistContext) {
	this := p
	_ = this

	localctx = NewParamlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, angelscriptParserRULE_paramlist)
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
		p.SetState(346)
		p.Match(angelscriptParserT__26)
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case angelscriptParserVOID:
		{
			p.SetState(347)
			p.Match(angelscriptParserVOID)
		}

	case angelscriptParserT__16, angelscriptParserT__36, angelscriptParserT__37, angelscriptParserT__38, angelscriptParserPRIMTYPE, angelscriptParserIDENTIFIER:
		{
			p.SetState(348)
			p.Type_()
		}
		{
			p.SetState(349)
			p.Typemod()
		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == angelscriptParserIDENTIFIER {
			{
				p.SetState(350)
				p.Match(angelscriptParserIDENTIFIER)
			}

		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == angelscriptParserT__18 {
			{
				p.SetState(353)
				p.Match(angelscriptParserT__18)
			}
			{
				p.SetState(354)
				p.Expr()
			}

		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == angelscriptParserT__7 {
			{
				p.SetState(357)
				p.Match(angelscriptParserT__7)
			}
			{
				p.SetState(358)
				p.Type_()
			}
			{
				p.SetState(359)
				p.Typemod()
			}
			p.SetState(361)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == angelscriptParserIDENTIFIER {
				{
					p.SetState(360)
					p.Match(angelscriptParserIDENTIFIER)
				}

			}
			p.SetState(365)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == angelscriptParserT__18 {
				{
					p.SetState(363)
					p.Match(angelscriptParserT__18)
				}
				{
					p.SetState(364)
					p.Expr()
				}

			}

			p.SetState(371)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case angelscriptParserT__27:

	default:
	}
	{
		p.SetState(374)
		p.Match(angelscriptParserT__27)
	}

	return localctx
}

// ITypemodContext is an interface to support dynamic dispatch.
type ITypemodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypemodContext differentiates from other interfaces.
	IsTypemodContext()
}

type TypemodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypemodContext() *TypemodContext {
	var p = new(TypemodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_typemod
	return p
}

func (*TypemodContext) IsTypemodContext() {}

func NewTypemodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypemodContext {
	var p = new(TypemodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_typemod

	return p
}

func (s *TypemodContext) GetParser() antlr.Parser { return s.parser }
func (s *TypemodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypemodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypemodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterTypemod(s)
	}
}

func (s *TypemodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitTypemod(s)
	}
}

func (p *angelscriptParser) Typemod() (localctx ITypemodContext) {
	this := p
	_ = this

	localctx = NewTypemodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, angelscriptParserRULE_typemod)
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
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__14 {
		{
			p.SetState(376)
			p.Match(angelscriptParserT__14)
		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__28)|(1<<angelscriptParserT__29)|(1<<angelscriptParserT__30))) != 0 {
			{
				p.SetState(377)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__28)|(1<<angelscriptParserT__29)|(1<<angelscriptParserT__30))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) Scope() IScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *Type_Context) Datatype() IDatatypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatatypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *Type_Context) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *Type_Context) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *angelscriptParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, angelscriptParserRULE_type_)
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
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__16 {
		{
			p.SetState(382)
			p.Match(angelscriptParserT__16)
		}

	}
	{
		p.SetState(385)
		p.Scope()
	}
	{
		p.SetState(386)
		p.Datatype()
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__31 {
		{
			p.SetState(387)
			p.Match(angelscriptParserT__31)
		}
		{
			p.SetState(388)
			p.Type_()
		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == angelscriptParserT__7 {
			{
				p.SetState(389)
				p.Match(angelscriptParserT__7)
			}
			{
				p.SetState(390)
				p.Type_()
			}

			p.SetState(395)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(396)
			p.Match(angelscriptParserT__32)
		}

		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__33 || _la == angelscriptParserT__35 {
		p.SetState(409)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case angelscriptParserT__33:
			{
				p.SetState(403)
				p.Match(angelscriptParserT__33)
			}
			{
				p.SetState(404)
				p.Match(angelscriptParserT__34)
			}

		case angelscriptParserT__35:
			{
				p.SetState(405)
				p.Match(angelscriptParserT__35)
			}
			p.SetState(407)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == angelscriptParserT__16 {
				{
					p.SetState(406)
					p.Match(angelscriptParserT__16)
				}

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInitlistContext is an interface to support dynamic dispatch.
type IInitlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitlistContext differentiates from other interfaces.
	IsInitlistContext()
}

type InitlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitlistContext() *InitlistContext {
	var p = new(InitlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_initlist
	return p
}

func (*InitlistContext) IsInitlistContext() {}

func NewInitlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitlistContext {
	var p = new(InitlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_initlist

	return p
}

func (s *InitlistContext) GetParser() antlr.Parser { return s.parser }

func (s *InitlistContext) AllAssign() []IAssignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignContext)(nil)).Elem())
	var tst = make([]IAssignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignContext)
		}
	}

	return tst
}

func (s *InitlistContext) Assign(i int) IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *InitlistContext) AllInitlist() []IInitlistContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInitlistContext)(nil)).Elem())
	var tst = make([]IInitlistContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInitlistContext)
		}
	}

	return tst
}

func (s *InitlistContext) Initlist(i int) IInitlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitlistContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInitlistContext)
}

func (s *InitlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterInitlist(s)
	}
}

func (s *InitlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitInitlist(s)
	}
}

func (p *angelscriptParser) Initlist() (localctx IInitlistContext) {
	this := p
	_ = this

	localctx = NewInitlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, angelscriptParserRULE_initlist)
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
		p.SetState(414)
		p.Match(angelscriptParserT__8)
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(415)
			p.Assign()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(416)
			p.Initlist()
		}

	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__7 {
		{
			p.SetState(419)
			p.Match(angelscriptParserT__7)
		}
		p.SetState(422)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(420)
				p.Assign()
			}

		} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(421)
				p.Initlist()
			}

		}

		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(429)
		p.Match(angelscriptParserT__9)
	}

	return localctx
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_scope
	return p
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserIDENTIFIER)
}

func (s *ScopeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, i)
}

func (s *ScopeContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *ScopeContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterScope(s)
	}
}

func (s *ScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitScope(s)
	}
}

func (p *angelscriptParser) Scope() (localctx IScopeContext) {
	this := p
	_ = this

	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, angelscriptParserRULE_scope)
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
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__36 {
		{
			p.SetState(431)
			p.Match(angelscriptParserT__36)
		}

	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(434)
				p.Match(angelscriptParserIDENTIFIER)
			}
			{
				p.SetState(435)
				p.Match(angelscriptParserT__36)
			}

		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(441)
			p.Match(angelscriptParserIDENTIFIER)
		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == angelscriptParserT__31 {
			{
				p.SetState(442)
				p.Match(angelscriptParserT__31)
			}
			{
				p.SetState(443)
				p.Type_()
			}
			p.SetState(448)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == angelscriptParserT__7 {
				{
					p.SetState(444)
					p.Match(angelscriptParserT__7)
				}
				{
					p.SetState(445)
					p.Type_()
				}

				p.SetState(450)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(451)
				p.Match(angelscriptParserT__32)
			}

		}
		{
			p.SetState(455)
			p.Match(angelscriptParserT__36)
		}

	}

	return localctx
}

// IDatatypeContext is an interface to support dynamic dispatch.
type IDatatypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatatypeContext differentiates from other interfaces.
	IsDatatypeContext()
}

type DatatypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatatypeContext() *DatatypeContext {
	var p = new(DatatypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_datatype
	return p
}

func (*DatatypeContext) IsDatatypeContext() {}

func NewDatatypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatatypeContext {
	var p = new(DatatypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_datatype

	return p
}

func (s *DatatypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DatatypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, 0)
}

func (s *DatatypeContext) PRIMTYPE() antlr.TerminalNode {
	return s.GetToken(angelscriptParserPRIMTYPE, 0)
}

func (s *DatatypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatatypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatatypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterDatatype(s)
	}
}

func (s *DatatypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitDatatype(s)
	}
}

func (p *angelscriptParser) Datatype() (localctx IDatatypeContext) {
	this := p
	_ = this

	localctx = NewDatatypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, angelscriptParserRULE_datatype)
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
		p.SetState(458)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(angelscriptParserT__37-38))|(1<<(angelscriptParserT__38-38))|(1<<(angelscriptParserPRIMTYPE-38))|(1<<(angelscriptParserIDENTIFIER-38)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = angelscriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) If_() IIf_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_Context)
}

func (s *StatementContext) For_() IFor_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_Context)
}

func (s *StatementContext) While_() IWhile_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_Context)
}

func (s *StatementContext) Return_() IReturn_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_Context)
}

func (s *StatementContext) Statblock() IStatblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatblockContext)
}

func (s *StatementContext) Break_() IBreak_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreak_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreak_Context)
}

func (s *StatementContext) Continue_() IContinue_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinue_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinue_Context)
}

func (s *StatementContext) Dowhile() IDowhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDowhileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDowhileContext)
}

func (s *StatementContext) Switch_() ISwitch_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitch_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitch_Context)
}

func (s *StatementContext) Exprstat() IExprstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprstatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprstatContext)
}

func (s *StatementContext) Try_() ITry_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITry_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITry_Context)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *angelscriptParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, angelscriptParserRULE_statement)

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
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(460)
			p.If_()
		}

	case 2:
		{
			p.SetState(461)
			p.For_()
		}

	case 3:
		{
			p.SetState(462)
			p.While_()
		}

	case 4:
		{
			p.SetState(463)
			p.Return_()
		}

	case 5:
		{
			p.SetState(464)
			p.Statblock()
		}

	case 6:
		{
			p.SetState(465)
			p.Break_()
		}

	case 7:
		{
			p.SetState(466)
			p.Continue_()
		}

	case 8:
		{
			p.SetState(467)
			p.Dowhile()
		}

	case 9:
		{
			p.SetState(468)
			p.Switch_()
		}

	case 10:
		{
			p.SetState(469)
			p.Exprstat()
		}

	case 11:
		{
			p.SetState(470)
			p.Try_()
		}

	}

	return localctx
}

// ISwitch_Context is an interface to support dynamic dispatch.
type ISwitch_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitch_Context differentiates from other interfaces.
	IsSwitch_Context()
}

type Switch_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_Context() *Switch_Context {
	var p = new(Switch_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_switch_
	return p
}

func (*Switch_Context) IsSwitch_Context() {}

func NewSwitch_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_Context {
	var p = new(Switch_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_switch_

	return p
}

func (s *Switch_Context) GetParser() antlr.Parser { return s.parser }

func (s *Switch_Context) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *Switch_Context) AllCase_() []ICase_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICase_Context)(nil)).Elem())
	var tst = make([]ICase_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICase_Context)
		}
	}

	return tst
}

func (s *Switch_Context) Case_(i int) ICase_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICase_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICase_Context)
}

func (s *Switch_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterSwitch_(s)
	}
}

func (s *Switch_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitSwitch_(s)
	}
}

func (p *angelscriptParser) Switch_() (localctx ISwitch_Context) {
	this := p
	_ = this

	localctx = NewSwitch_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, angelscriptParserRULE_switch_)
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
		p.SetState(473)
		p.Match(angelscriptParserT__39)
	}
	{
		p.SetState(474)
		p.Match(angelscriptParserT__26)
	}
	{
		p.SetState(475)
		p.Assign()
	}
	{
		p.SetState(476)
		p.Match(angelscriptParserT__27)
	}
	{
		p.SetState(477)
		p.Match(angelscriptParserT__8)
	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__50 || _la == angelscriptParserT__51 {
		{
			p.SetState(478)
			p.Case_()
		}

		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(484)
		p.Match(angelscriptParserT__9)
	}

	return localctx
}

// IBreak_Context is an interface to support dynamic dispatch.
type IBreak_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreak_Context differentiates from other interfaces.
	IsBreak_Context()
}

type Break_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_Context() *Break_Context {
	var p = new(Break_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_break_
	return p
}

func (*Break_Context) IsBreak_Context() {}

func NewBreak_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_Context {
	var p = new(Break_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_break_

	return p
}

func (s *Break_Context) GetParser() antlr.Parser { return s.parser }
func (s *Break_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterBreak_(s)
	}
}

func (s *Break_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitBreak_(s)
	}
}

func (p *angelscriptParser) Break_() (localctx IBreak_Context) {
	this := p
	_ = this

	localctx = NewBreak_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, angelscriptParserRULE_break_)

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
		p.SetState(486)
		p.Match(angelscriptParserT__40)
	}
	{
		p.SetState(487)
		p.Match(angelscriptParserT__0)
	}

	return localctx
}

// IFor_Context is an interface to support dynamic dispatch.
type IFor_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_Context differentiates from other interfaces.
	IsFor_Context()
}

type For_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_Context() *For_Context {
	var p = new(For_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_for_
	return p
}

func (*For_Context) IsFor_Context() {}

func NewFor_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_Context {
	var p = new(For_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_for_

	return p
}

func (s *For_Context) GetParser() antlr.Parser { return s.parser }

func (s *For_Context) AllExprstat() []IExprstatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprstatContext)(nil)).Elem())
	var tst = make([]IExprstatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprstatContext)
		}
	}

	return tst
}

func (s *For_Context) Exprstat(i int) IExprstatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprstatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprstatContext)
}

func (s *For_Context) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *For_Context) Var_() IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *For_Context) AllAssign() []IAssignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignContext)(nil)).Elem())
	var tst = make([]IAssignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignContext)
		}
	}

	return tst
}

func (s *For_Context) Assign(i int) IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *For_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterFor_(s)
	}
}

func (s *For_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitFor_(s)
	}
}

func (p *angelscriptParser) For_() (localctx IFor_Context) {
	this := p
	_ = this

	localctx = NewFor_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, angelscriptParserRULE_for_)
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
		p.SetState(489)
		p.Match(angelscriptParserT__41)
	}
	{
		p.SetState(490)
		p.Match(angelscriptParserT__26)
	}
	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(491)
			p.Var_()
		}

	case 2:
		{
			p.SetState(492)
			p.Exprstat()
		}

	}
	{
		p.SetState(495)
		p.Exprstat()
	}
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__8)|(1<<angelscriptParserT__16)|(1<<angelscriptParserT__26))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(angelscriptParserT__36-37))|(1<<(angelscriptParserT__37-37))|(1<<(angelscriptParserT__38-37))|(1<<(angelscriptParserT__55-37))|(1<<(angelscriptParserT__56-37))|(1<<(angelscriptParserPRIMTYPE-37))|(1<<(angelscriptParserEXPRPREOP-37))|(1<<(angelscriptParserVOID-37))|(1<<(angelscriptParserLITERAL-37))|(1<<(angelscriptParserIDENTIFIER-37)))) != 0) {
		{
			p.SetState(496)
			p.Assign()
		}
		p.SetState(501)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == angelscriptParserT__7 {
			{
				p.SetState(497)
				p.Match(angelscriptParserT__7)
			}
			{
				p.SetState(498)
				p.Assign()
			}

			p.SetState(503)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(506)
		p.Match(angelscriptParserT__27)
	}
	{
		p.SetState(507)
		p.Statement()
	}

	return localctx
}

// IWhile_Context is an interface to support dynamic dispatch.
type IWhile_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhile_Context differentiates from other interfaces.
	IsWhile_Context()
}

type While_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_Context() *While_Context {
	var p = new(While_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_while_
	return p
}

func (*While_Context) IsWhile_Context() {}

func NewWhile_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_Context {
	var p = new(While_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_while_

	return p
}

func (s *While_Context) GetParser() antlr.Parser { return s.parser }

func (s *While_Context) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *While_Context) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *While_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterWhile_(s)
	}
}

func (s *While_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitWhile_(s)
	}
}

func (p *angelscriptParser) While_() (localctx IWhile_Context) {
	this := p
	_ = this

	localctx = NewWhile_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, angelscriptParserRULE_while_)

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
		p.SetState(509)
		p.Match(angelscriptParserT__42)
	}
	{
		p.SetState(510)
		p.Match(angelscriptParserT__26)
	}
	{
		p.SetState(511)
		p.Assign()
	}
	{
		p.SetState(512)
		p.Match(angelscriptParserT__27)
	}
	{
		p.SetState(513)
		p.Statement()
	}

	return localctx
}

// IDowhileContext is an interface to support dynamic dispatch.
type IDowhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDowhileContext differentiates from other interfaces.
	IsDowhileContext()
}

type DowhileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDowhileContext() *DowhileContext {
	var p = new(DowhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_dowhile
	return p
}

func (*DowhileContext) IsDowhileContext() {}

func NewDowhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DowhileContext {
	var p = new(DowhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_dowhile

	return p
}

func (s *DowhileContext) GetParser() antlr.Parser { return s.parser }

func (s *DowhileContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DowhileContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *DowhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DowhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DowhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterDowhile(s)
	}
}

func (s *DowhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitDowhile(s)
	}
}

func (p *angelscriptParser) Dowhile() (localctx IDowhileContext) {
	this := p
	_ = this

	localctx = NewDowhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, angelscriptParserRULE_dowhile)

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
		p.SetState(515)
		p.Match(angelscriptParserT__43)
	}
	{
		p.SetState(516)
		p.Statement()
	}
	{
		p.SetState(517)
		p.Match(angelscriptParserT__42)
	}
	{
		p.SetState(518)
		p.Match(angelscriptParserT__26)
	}
	{
		p.SetState(519)
		p.Assign()
	}
	{
		p.SetState(520)
		p.Match(angelscriptParserT__27)
	}
	{
		p.SetState(521)
		p.Match(angelscriptParserT__0)
	}

	return localctx
}

// IIf_Context is an interface to support dynamic dispatch.
type IIf_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_Context differentiates from other interfaces.
	IsIf_Context()
}

type If_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_Context() *If_Context {
	var p = new(If_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_if_
	return p
}

func (*If_Context) IsIf_Context() {}

func NewIf_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_Context {
	var p = new(If_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_if_

	return p
}

func (s *If_Context) GetParser() antlr.Parser { return s.parser }

func (s *If_Context) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *If_Context) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *If_Context) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *If_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterIf_(s)
	}
}

func (s *If_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitIf_(s)
	}
}

func (p *angelscriptParser) If_() (localctx IIf_Context) {
	this := p
	_ = this

	localctx = NewIf_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, angelscriptParserRULE_if_)

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
		p.SetState(523)
		p.Match(angelscriptParserT__44)
	}
	{
		p.SetState(524)
		p.Match(angelscriptParserT__26)
	}
	{
		p.SetState(525)
		p.Assign()
	}
	{
		p.SetState(526)
		p.Match(angelscriptParserT__27)
	}
	{
		p.SetState(527)
		p.Statement()
	}
	p.SetState(530)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(528)
			p.Match(angelscriptParserT__45)
		}
		{
			p.SetState(529)
			p.Statement()
		}

	}

	return localctx
}

// IContinue_Context is an interface to support dynamic dispatch.
type IContinue_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinue_Context differentiates from other interfaces.
	IsContinue_Context()
}

type Continue_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_Context() *Continue_Context {
	var p = new(Continue_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_continue_
	return p
}

func (*Continue_Context) IsContinue_Context() {}

func NewContinue_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_Context {
	var p = new(Continue_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_continue_

	return p
}

func (s *Continue_Context) GetParser() antlr.Parser { return s.parser }
func (s *Continue_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterContinue_(s)
	}
}

func (s *Continue_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitContinue_(s)
	}
}

func (p *angelscriptParser) Continue_() (localctx IContinue_Context) {
	this := p
	_ = this

	localctx = NewContinue_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, angelscriptParserRULE_continue_)

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
		p.SetState(532)
		p.Match(angelscriptParserT__46)
	}
	{
		p.SetState(533)
		p.Match(angelscriptParserT__0)
	}

	return localctx
}

// IExprstatContext is an interface to support dynamic dispatch.
type IExprstatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprstatContext differentiates from other interfaces.
	IsExprstatContext()
}

type ExprstatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprstatContext() *ExprstatContext {
	var p = new(ExprstatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_exprstat
	return p
}

func (*ExprstatContext) IsExprstatContext() {}

func NewExprstatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprstatContext {
	var p = new(ExprstatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_exprstat

	return p
}

func (s *ExprstatContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprstatContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ExprstatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprstatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprstatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterExprstat(s)
	}
}

func (s *ExprstatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitExprstat(s)
	}
}

func (p *angelscriptParser) Exprstat() (localctx IExprstatContext) {
	this := p
	_ = this

	localctx = NewExprstatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, angelscriptParserRULE_exprstat)
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

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__8)|(1<<angelscriptParserT__16)|(1<<angelscriptParserT__26))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(angelscriptParserT__36-37))|(1<<(angelscriptParserT__37-37))|(1<<(angelscriptParserT__38-37))|(1<<(angelscriptParserT__55-37))|(1<<(angelscriptParserT__56-37))|(1<<(angelscriptParserPRIMTYPE-37))|(1<<(angelscriptParserEXPRPREOP-37))|(1<<(angelscriptParserVOID-37))|(1<<(angelscriptParserLITERAL-37))|(1<<(angelscriptParserIDENTIFIER-37)))) != 0) {
		{
			p.SetState(535)
			p.Assign()
		}

	}
	{
		p.SetState(538)
		p.Match(angelscriptParserT__0)
	}

	return localctx
}

// ITry_Context is an interface to support dynamic dispatch.
type ITry_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTry_Context differentiates from other interfaces.
	IsTry_Context()
}

type Try_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTry_Context() *Try_Context {
	var p = new(Try_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_try_
	return p
}

func (*Try_Context) IsTry_Context() {}

func NewTry_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Try_Context {
	var p = new(Try_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_try_

	return p
}

func (s *Try_Context) GetParser() antlr.Parser { return s.parser }

func (s *Try_Context) AllStatblock() []IStatblockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatblockContext)(nil)).Elem())
	var tst = make([]IStatblockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatblockContext)
		}
	}

	return tst
}

func (s *Try_Context) Statblock(i int) IStatblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatblockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatblockContext)
}

func (s *Try_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Try_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Try_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterTry_(s)
	}
}

func (s *Try_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitTry_(s)
	}
}

func (p *angelscriptParser) Try_() (localctx ITry_Context) {
	this := p
	_ = this

	localctx = NewTry_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, angelscriptParserRULE_try_)

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
		p.SetState(540)
		p.Match(angelscriptParserT__47)
	}
	{
		p.SetState(541)
		p.Statblock()
	}
	{
		p.SetState(542)
		p.Match(angelscriptParserT__48)
	}
	{
		p.SetState(543)
		p.Statblock()
	}

	return localctx
}

// IReturn_Context is an interface to support dynamic dispatch.
type IReturn_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturn_Context differentiates from other interfaces.
	IsReturn_Context()
}

type Return_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_Context() *Return_Context {
	var p = new(Return_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_return_
	return p
}

func (*Return_Context) IsReturn_Context() {}

func NewReturn_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_Context {
	var p = new(Return_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_return_

	return p
}

func (s *Return_Context) GetParser() antlr.Parser { return s.parser }

func (s *Return_Context) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *Return_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterReturn_(s)
	}
}

func (s *Return_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitReturn_(s)
	}
}

func (p *angelscriptParser) Return_() (localctx IReturn_Context) {
	this := p
	_ = this

	localctx = NewReturn_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, angelscriptParserRULE_return_)
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
		p.SetState(545)
		p.Match(angelscriptParserT__49)
	}
	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__8)|(1<<angelscriptParserT__16)|(1<<angelscriptParserT__26))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(angelscriptParserT__36-37))|(1<<(angelscriptParserT__37-37))|(1<<(angelscriptParserT__38-37))|(1<<(angelscriptParserT__55-37))|(1<<(angelscriptParserT__56-37))|(1<<(angelscriptParserPRIMTYPE-37))|(1<<(angelscriptParserEXPRPREOP-37))|(1<<(angelscriptParserVOID-37))|(1<<(angelscriptParserLITERAL-37))|(1<<(angelscriptParserIDENTIFIER-37)))) != 0) {
		{
			p.SetState(546)
			p.Assign()
		}

	}
	{
		p.SetState(549)
		p.Match(angelscriptParserT__0)
	}

	return localctx
}

// ICase_Context is an interface to support dynamic dispatch.
type ICase_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCase_Context differentiates from other interfaces.
	IsCase_Context()
}

type Case_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_Context() *Case_Context {
	var p = new(Case_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_case_
	return p
}

func (*Case_Context) IsCase_Context() {}

func NewCase_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_Context {
	var p = new(Case_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_case_

	return p
}

func (s *Case_Context) GetParser() antlr.Parser { return s.parser }

func (s *Case_Context) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *Case_Context) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Case_Context) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Case_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterCase_(s)
	}
}

func (s *Case_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitCase_(s)
	}
}

func (p *angelscriptParser) Case_() (localctx ICase_Context) {
	this := p
	_ = this

	localctx = NewCase_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, angelscriptParserRULE_case_)
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
	p.SetState(554)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case angelscriptParserT__50:
		{
			p.SetState(551)
			p.Match(angelscriptParserT__50)
		}
		{
			p.SetState(552)
			p.Expr()
		}

	case angelscriptParserT__51:
		{
			p.SetState(553)
			p.Match(angelscriptParserT__51)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(556)
		p.Match(angelscriptParserT__6)
	}
	p.SetState(560)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<angelscriptParserT__0)|(1<<angelscriptParserT__8)|(1<<angelscriptParserT__16)|(1<<angelscriptParserT__26))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(angelscriptParserT__36-37))|(1<<(angelscriptParserT__37-37))|(1<<(angelscriptParserT__38-37))|(1<<(angelscriptParserT__39-37))|(1<<(angelscriptParserT__40-37))|(1<<(angelscriptParserT__41-37))|(1<<(angelscriptParserT__42-37))|(1<<(angelscriptParserT__43-37))|(1<<(angelscriptParserT__44-37))|(1<<(angelscriptParserT__46-37))|(1<<(angelscriptParserT__47-37))|(1<<(angelscriptParserT__49-37))|(1<<(angelscriptParserT__55-37))|(1<<(angelscriptParserT__56-37))|(1<<(angelscriptParserPRIMTYPE-37))|(1<<(angelscriptParserEXPRPREOP-37))|(1<<(angelscriptParserVOID-37))|(1<<(angelscriptParserLITERAL-37))|(1<<(angelscriptParserIDENTIFIER-37)))) != 0) {
		{
			p.SetState(557)
			p.Statement()
		}

		p.SetState(562)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllExprterm() []IExprtermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprtermContext)(nil)).Elem())
	var tst = make([]IExprtermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprtermContext)
		}
	}

	return tst
}

func (s *ExprContext) Exprterm(i int) IExprtermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprtermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprtermContext)
}

func (s *ExprContext) AllExprop() []IExpropContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpropContext)(nil)).Elem())
	var tst = make([]IExpropContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpropContext)
		}
	}

	return tst
}

func (s *ExprContext) Exprop(i int) IExpropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpropContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpropContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *angelscriptParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, angelscriptParserRULE_expr)
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
		p.SetState(563)
		p.Exprterm()
	}
	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(angelscriptParserBITOP-58))|(1<<(angelscriptParserMATHOP-58))|(1<<(angelscriptParserCOMPOP-58))|(1<<(angelscriptParserLOGICOP-58)))) != 0 {
		{
			p.SetState(564)
			p.Exprop()
		}
		{
			p.SetState(565)
			p.Exprterm()
		}

		p.SetState(571)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprtermContext is an interface to support dynamic dispatch.
type IExprtermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprtermContext differentiates from other interfaces.
	IsExprtermContext()
}

type ExprtermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprtermContext() *ExprtermContext {
	var p = new(ExprtermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_exprterm
	return p
}

func (*ExprtermContext) IsExprtermContext() {}

func NewExprtermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprtermContext {
	var p = new(ExprtermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_exprterm

	return p
}

func (s *ExprtermContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprtermContext) Initlist() IInitlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitlistContext)
}

func (s *ExprtermContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ExprtermContext) Exprvalue() IExprvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprvalueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprvalueContext)
}

func (s *ExprtermContext) AllEXPRPREOP() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserEXPRPREOP)
}

func (s *ExprtermContext) EXPRPREOP(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserEXPRPREOP, i)
}

func (s *ExprtermContext) AllExprpostop() []IExprpostopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprpostopContext)(nil)).Elem())
	var tst = make([]IExprpostopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprpostopContext)
		}
	}

	return tst
}

func (s *ExprtermContext) Exprpostop(i int) IExprpostopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprpostopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprpostopContext)
}

func (s *ExprtermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprtermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprtermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterExprterm(s)
	}
}

func (s *ExprtermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitExprterm(s)
	}
}

func (p *angelscriptParser) Exprterm() (localctx IExprtermContext) {
	this := p
	_ = this

	localctx = NewExprtermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, angelscriptParserRULE_exprterm)
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

	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 81, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == angelscriptParserT__16 || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(angelscriptParserT__36-37))|(1<<(angelscriptParserT__37-37))|(1<<(angelscriptParserT__38-37))|(1<<(angelscriptParserPRIMTYPE-37))|(1<<(angelscriptParserIDENTIFIER-37)))) != 0) {
			{
				p.SetState(572)
				p.Type_()
			}
			{
				p.SetState(573)
				p.Match(angelscriptParserT__18)
			}

		}
		{
			p.SetState(577)
			p.Initlist()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(581)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == angelscriptParserEXPRPREOP {
			{
				p.SetState(578)
				p.Match(angelscriptParserEXPRPREOP)
			}

			p.SetState(583)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(584)
			p.Exprvalue()
		}
		p.SetState(588)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(angelscriptParserT__26-27))|(1<<(angelscriptParserT__33-27))|(1<<(angelscriptParserT__52-27))|(1<<(angelscriptParserT__53-27))|(1<<(angelscriptParserT__54-27)))) != 0 {
			{
				p.SetState(585)
				p.Exprpostop()
			}

			p.SetState(590)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IExprvalueContext is an interface to support dynamic dispatch.
type IExprvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprvalueContext differentiates from other interfaces.
	IsExprvalueContext()
}

type ExprvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprvalueContext() *ExprvalueContext {
	var p = new(ExprvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_exprvalue
	return p
}

func (*ExprvalueContext) IsExprvalueContext() {}

func NewExprvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprvalueContext {
	var p = new(ExprvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_exprvalue

	return p
}

func (s *ExprvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprvalueContext) VOID() antlr.TerminalNode {
	return s.GetToken(angelscriptParserVOID, 0)
}

func (s *ExprvalueContext) Constructcall() IConstructcallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstructcallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstructcallContext)
}

func (s *ExprvalueContext) Funccall() IFunccallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunccallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunccallContext)
}

func (s *ExprvalueContext) Varaccess() IVaraccessContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVaraccessContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVaraccessContext)
}

func (s *ExprvalueContext) Cast() ICastContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICastContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICastContext)
}

func (s *ExprvalueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(angelscriptParserLITERAL, 0)
}

func (s *ExprvalueContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ExprvalueContext) Lambda_() ILambda_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambda_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILambda_Context)
}

func (s *ExprvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterExprvalue(s)
	}
}

func (s *ExprvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitExprvalue(s)
	}
}

func (p *angelscriptParser) Exprvalue() (localctx IExprvalueContext) {
	this := p
	_ = this

	localctx = NewExprvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, angelscriptParserRULE_exprvalue)

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

	p.SetState(604)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 82, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(593)
			p.Match(angelscriptParserVOID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(594)
			p.Constructcall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(595)
			p.Funccall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(596)
			p.Varaccess()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(597)
			p.Cast()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(598)
			p.Match(angelscriptParserLITERAL)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(599)
			p.Match(angelscriptParserT__26)
		}
		{
			p.SetState(600)
			p.Assign()
		}
		{
			p.SetState(601)
			p.Match(angelscriptParserT__27)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(603)
			p.Lambda_()
		}

	}

	return localctx
}

// IConstructcallContext is an interface to support dynamic dispatch.
type IConstructcallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstructcallContext differentiates from other interfaces.
	IsConstructcallContext()
}

type ConstructcallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructcallContext() *ConstructcallContext {
	var p = new(ConstructcallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_constructcall
	return p
}

func (*ConstructcallContext) IsConstructcallContext() {}

func NewConstructcallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructcallContext {
	var p = new(ConstructcallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_constructcall

	return p
}

func (s *ConstructcallContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructcallContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ConstructcallContext) Arglist() IArglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *ConstructcallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructcallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructcallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterConstructcall(s)
	}
}

func (s *ConstructcallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitConstructcall(s)
	}
}

func (p *angelscriptParser) Constructcall() (localctx IConstructcallContext) {
	this := p
	_ = this

	localctx = NewConstructcallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, angelscriptParserRULE_constructcall)

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
		p.SetState(606)
		p.Type_()
	}
	{
		p.SetState(607)
		p.Arglist()
	}

	return localctx
}

// IExprpostopContext is an interface to support dynamic dispatch.
type IExprpostopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprpostopContext differentiates from other interfaces.
	IsExprpostopContext()
}

type ExprpostopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprpostopContext() *ExprpostopContext {
	var p = new(ExprpostopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_exprpostop
	return p
}

func (*ExprpostopContext) IsExprpostopContext() {}

func NewExprpostopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprpostopContext {
	var p = new(ExprpostopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_exprpostop

	return p
}

func (s *ExprpostopContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprpostopContext) Funccall() IFunccallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunccallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunccallContext)
}

func (s *ExprpostopContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserIDENTIFIER)
}

func (s *ExprpostopContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, i)
}

func (s *ExprpostopContext) AllAssign() []IAssignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignContext)(nil)).Elem())
	var tst = make([]IAssignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignContext)
		}
	}

	return tst
}

func (s *ExprpostopContext) Assign(i int) IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ExprpostopContext) Arglist() IArglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *ExprpostopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprpostopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprpostopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterExprpostop(s)
	}
}

func (s *ExprpostopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitExprpostop(s)
	}
}

func (p *angelscriptParser) Exprpostop() (localctx IExprpostopContext) {
	this := p
	_ = this

	localctx = NewExprpostopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, angelscriptParserRULE_exprpostop)
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

	p.SetState(636)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case angelscriptParserT__52:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(609)
			p.Match(angelscriptParserT__52)
		}
		p.SetState(612)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(610)
				p.Funccall()
			}

		case 2:
			{
				p.SetState(611)
				p.Match(angelscriptParserIDENTIFIER)
			}

		}

	case angelscriptParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(614)
			p.Match(angelscriptParserT__33)
		}
		p.SetState(617)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 84, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(615)
				p.Match(angelscriptParserIDENTIFIER)
			}
			{
				p.SetState(616)
				p.Match(angelscriptParserT__6)
			}

		}
		{
			p.SetState(619)
			p.Assign()
		}
		p.SetState(628)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == angelscriptParserT__7 {
			{
				p.SetState(620)
				p.Match(angelscriptParserT__7)
			}
			p.SetState(622)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == angelscriptParserIDENTIFIER {
				{
					p.SetState(621)
					p.Match(angelscriptParserIDENTIFIER)
				}

			}
			{
				p.SetState(624)
				p.Match(angelscriptParserT__6)
			}
			{
				p.SetState(625)
				p.Assign()
			}

			p.SetState(630)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(631)
			p.Match(angelscriptParserT__34)
		}

	case angelscriptParserT__26:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(633)
			p.Arglist()
		}

	case angelscriptParserT__53:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(634)
			p.Match(angelscriptParserT__53)
		}

	case angelscriptParserT__54:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(635)
			p.Match(angelscriptParserT__54)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICastContext is an interface to support dynamic dispatch.
type ICastContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCastContext differentiates from other interfaces.
	IsCastContext()
}

type CastContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCastContext() *CastContext {
	var p = new(CastContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_cast
	return p
}

func (*CastContext) IsCastContext() {}

func NewCastContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CastContext {
	var p = new(CastContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_cast

	return p
}

func (s *CastContext) GetParser() antlr.Parser { return s.parser }

func (s *CastContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *CastContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *CastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterCast(s)
	}
}

func (s *CastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitCast(s)
	}
}

func (p *angelscriptParser) Cast() (localctx ICastContext) {
	this := p
	_ = this

	localctx = NewCastContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, angelscriptParserRULE_cast)

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
		p.Match(angelscriptParserT__55)
	}
	{
		p.SetState(639)
		p.Match(angelscriptParserT__31)
	}
	{
		p.SetState(640)
		p.Type_()
	}
	{
		p.SetState(641)
		p.Match(angelscriptParserT__32)
	}
	{
		p.SetState(642)
		p.Match(angelscriptParserT__26)
	}
	{
		p.SetState(643)
		p.Assign()
	}
	{
		p.SetState(644)
		p.Match(angelscriptParserT__27)
	}

	return localctx
}

// ILambda_Context is an interface to support dynamic dispatch.
type ILambda_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambda_Context differentiates from other interfaces.
	IsLambda_Context()
}

type Lambda_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambda_Context() *Lambda_Context {
	var p = new(Lambda_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_lambda_
	return p
}

func (*Lambda_Context) IsLambda_Context() {}

func NewLambda_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lambda_Context {
	var p = new(Lambda_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_lambda_

	return p
}

func (s *Lambda_Context) GetParser() antlr.Parser { return s.parser }

func (s *Lambda_Context) Statblock() IStatblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatblockContext)
}

func (s *Lambda_Context) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserIDENTIFIER)
}

func (s *Lambda_Context) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, i)
}

func (s *Lambda_Context) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *Lambda_Context) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Lambda_Context) AllTypemod() []ITypemodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypemodContext)(nil)).Elem())
	var tst = make([]ITypemodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypemodContext)
		}
	}

	return tst
}

func (s *Lambda_Context) Typemod(i int) ITypemodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypemodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypemodContext)
}

func (s *Lambda_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lambda_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lambda_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterLambda_(s)
	}
}

func (s *Lambda_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitLambda_(s)
	}
}

func (p *angelscriptParser) Lambda_() (localctx ILambda_Context) {
	this := p
	_ = this

	localctx = NewLambda_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, angelscriptParserRULE_lambda_)
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
		p.Match(angelscriptParserT__56)
	}
	{
		p.SetState(647)
		p.Match(angelscriptParserT__26)
	}
	p.SetState(666)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__16 || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(angelscriptParserT__36-37))|(1<<(angelscriptParserT__37-37))|(1<<(angelscriptParserT__38-37))|(1<<(angelscriptParserPRIMTYPE-37))|(1<<(angelscriptParserIDENTIFIER-37)))) != 0) {
		p.SetState(651)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 88, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(648)
				p.Type_()
			}
			{
				p.SetState(649)
				p.Typemod()
			}

		}
		{
			p.SetState(653)
			p.Match(angelscriptParserIDENTIFIER)
		}
		p.SetState(663)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == angelscriptParserT__7 {
			{
				p.SetState(654)
				p.Match(angelscriptParserT__7)
			}
			p.SetState(658)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 89, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(655)
					p.Type_()
				}
				{
					p.SetState(656)
					p.Typemod()
				}

			}
			{
				p.SetState(660)
				p.Match(angelscriptParserIDENTIFIER)
			}

			p.SetState(665)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(668)
		p.Match(angelscriptParserT__27)
	}
	{
		p.SetState(669)
		p.Statblock()
	}

	return localctx
}

// IFunccallContext is an interface to support dynamic dispatch.
type IFunccallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunccallContext differentiates from other interfaces.
	IsFunccallContext()
}

type FunccallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunccallContext() *FunccallContext {
	var p = new(FunccallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_funccall
	return p
}

func (*FunccallContext) IsFunccallContext() {}

func NewFunccallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunccallContext {
	var p = new(FunccallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_funccall

	return p
}

func (s *FunccallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunccallContext) Scope() IScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *FunccallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, 0)
}

func (s *FunccallContext) Arglist() IArglistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArglistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *FunccallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunccallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunccallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterFunccall(s)
	}
}

func (s *FunccallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitFunccall(s)
	}
}

func (p *angelscriptParser) Funccall() (localctx IFunccallContext) {
	this := p
	_ = this

	localctx = NewFunccallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, angelscriptParserRULE_funccall)

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
		p.SetState(671)
		p.Scope()
	}
	{
		p.SetState(672)
		p.Match(angelscriptParserIDENTIFIER)
	}
	{
		p.SetState(673)
		p.Arglist()
	}

	return localctx
}

// IVaraccessContext is an interface to support dynamic dispatch.
type IVaraccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVaraccessContext differentiates from other interfaces.
	IsVaraccessContext()
}

type VaraccessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVaraccessContext() *VaraccessContext {
	var p = new(VaraccessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_varaccess
	return p
}

func (*VaraccessContext) IsVaraccessContext() {}

func NewVaraccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VaraccessContext {
	var p = new(VaraccessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_varaccess

	return p
}

func (s *VaraccessContext) GetParser() antlr.Parser { return s.parser }

func (s *VaraccessContext) Scope() IScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *VaraccessContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, 0)
}

func (s *VaraccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VaraccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VaraccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterVaraccess(s)
	}
}

func (s *VaraccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitVaraccess(s)
	}
}

func (p *angelscriptParser) Varaccess() (localctx IVaraccessContext) {
	this := p
	_ = this

	localctx = NewVaraccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, angelscriptParserRULE_varaccess)

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
		p.SetState(675)
		p.Scope()
	}
	{
		p.SetState(676)
		p.Match(angelscriptParserIDENTIFIER)
	}

	return localctx
}

// IArglistContext is an interface to support dynamic dispatch.
type IArglistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArglistContext differentiates from other interfaces.
	IsArglistContext()
}

type ArglistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArglistContext() *ArglistContext {
	var p = new(ArglistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_arglist
	return p
}

func (*ArglistContext) IsArglistContext() {}

func NewArglistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArglistContext {
	var p = new(ArglistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_arglist

	return p
}

func (s *ArglistContext) GetParser() antlr.Parser { return s.parser }

func (s *ArglistContext) AllAssign() []IAssignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignContext)(nil)).Elem())
	var tst = make([]IAssignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignContext)
		}
	}

	return tst
}

func (s *ArglistContext) Assign(i int) IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ArglistContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(angelscriptParserIDENTIFIER)
}

func (s *ArglistContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(angelscriptParserIDENTIFIER, i)
}

func (s *ArglistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArglistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArglistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterArglist(s)
	}
}

func (s *ArglistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitArglist(s)
	}
}

func (p *angelscriptParser) Arglist() (localctx IArglistContext) {
	this := p
	_ = this

	localctx = NewArglistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, angelscriptParserRULE_arglist)
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
		p.SetState(678)
		p.Match(angelscriptParserT__26)
	}
	p.SetState(681)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 92, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(679)
			p.Match(angelscriptParserIDENTIFIER)
		}
		{
			p.SetState(680)
			p.Match(angelscriptParserT__6)
		}

	}
	{
		p.SetState(683)
		p.Assign()
	}
	p.SetState(692)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == angelscriptParserT__7 {
		{
			p.SetState(684)
			p.Match(angelscriptParserT__7)
		}
		p.SetState(687)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 93, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(685)
				p.Match(angelscriptParserIDENTIFIER)
			}
			{
				p.SetState(686)
				p.Match(angelscriptParserT__6)
			}

		}
		{
			p.SetState(689)
			p.Assign()
		}

		p.SetState(694)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(695)
		p.Match(angelscriptParserT__27)
	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *AssignContext) ASSIGNOP() antlr.TerminalNode {
	return s.GetToken(angelscriptParserASSIGNOP, 0)
}

func (s *AssignContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *angelscriptParser) Assign() (localctx IAssignContext) {
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, angelscriptParserRULE_assign)

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
		p.SetState(697)
		p.Condition()
	}
	p.SetState(700)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 95, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(698)
			p.Match(angelscriptParserASSIGNOP)
		}
		{
			p.SetState(699)
			p.Assign()
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
	p.RuleIndex = angelscriptParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionContext) AllAssign() []IAssignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignContext)(nil)).Elem())
	var tst = make([]IAssignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignContext)
		}
	}

	return tst
}

func (s *ConditionContext) Assign(i int) IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *angelscriptParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, angelscriptParserRULE_condition)
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
		p.SetState(702)
		p.Expr()
	}
	p.SetState(708)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == angelscriptParserT__37 {
		{
			p.SetState(703)
			p.Match(angelscriptParserT__37)
		}
		{
			p.SetState(704)
			p.Assign()
		}
		{
			p.SetState(705)
			p.Match(angelscriptParserT__6)
		}
		{
			p.SetState(706)
			p.Assign()
		}

	}

	return localctx
}

// IExpropContext is an interface to support dynamic dispatch.
type IExpropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpropContext differentiates from other interfaces.
	IsExpropContext()
}

type ExpropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpropContext() *ExpropContext {
	var p = new(ExpropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = angelscriptParserRULE_exprop
	return p
}

func (*ExpropContext) IsExpropContext() {}

func NewExpropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpropContext {
	var p = new(ExpropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = angelscriptParserRULE_exprop

	return p
}

func (s *ExpropContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpropContext) MATHOP() antlr.TerminalNode {
	return s.GetToken(angelscriptParserMATHOP, 0)
}

func (s *ExpropContext) COMPOP() antlr.TerminalNode {
	return s.GetToken(angelscriptParserCOMPOP, 0)
}

func (s *ExpropContext) LOGICOP() antlr.TerminalNode {
	return s.GetToken(angelscriptParserLOGICOP, 0)
}

func (s *ExpropContext) BITOP() antlr.TerminalNode {
	return s.GetToken(angelscriptParserBITOP, 0)
}

func (s *ExpropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.EnterExprop(s)
	}
}

func (s *ExpropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(angelscriptListener); ok {
		listenerT.ExitExprop(s)
	}
}

func (p *angelscriptParser) Exprop() (localctx IExpropContext) {
	this := p
	_ = this

	localctx = NewExpropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, angelscriptParserRULE_exprop)
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
		p.SetState(710)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(angelscriptParserBITOP-58))|(1<<(angelscriptParserMATHOP-58))|(1<<(angelscriptParserCOMPOP-58))|(1<<(angelscriptParserLOGICOP-58)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
